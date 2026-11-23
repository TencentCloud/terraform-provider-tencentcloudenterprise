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

type ConstraintRange struct {

	// 约束类型为section时的最大值

	Max *string `json:"Max,omitempty" name:"Max"`
	// 约束类型为section时的最小值

	Min *string `json:"Min,omitempty" name:"Min"`
}

type DBParamValue struct {

	// 参数名称

	Param *string `json:"Param,omitempty" name:"Param"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceConfig struct {

	// 实例配置ID，形如：1

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 修改时间

	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`
	// 参数名称

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 修改后参数值

	ParamNewValue *string `json:"ParamNewValue,omitempty" name:"ParamNewValue"`
	// 修改前参数值

	ParamOldValue *string `json:"ParamOldValue,omitempty" name:"ParamOldValue"`
	// 修改状态, 0:成功, -1:失败, -2:值非法

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 新的账号备注，长度 0~256。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否启用加密，1-已开启；0-未开启。

		EncryptStatus *int64 `json:"EncryptStatus,omitempty" name:"EncryptStatus"`
		// DEK密钥

		CipherText *string `json:"CipherText,omitempty" name:"CipherText"`
		// DEK密钥过期日期。

		ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEncryptAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// DB节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// DB节点角色，取值为master或者slave

	Role *string `json:"Role,omitempty" name:"Role"`
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：tdsql-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBRollbackInstancesRequest struct {
	*tchttp.BaseRequest

	// 回档源实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 偏移量，默认为 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 10，最大值为 150。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBRollbackInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBRollbackInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数模板Id

		ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBEncryptAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEncryptAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceHAInfoRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceHAInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceHAInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceSpecsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按机型分类的可售卖规格列表

		Specs []*InstanceSpec `json:"Specs,omitempty" name:"Specs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceSpecsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 会话ID列表

	SessionId []*int64 `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *KillSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoTypeRequest struct {
	*tchttp.BaseRequest

	// 只读实例 ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 只读实例变更后类型，取值 NORMAL（普通只读实例）、DELAY_REPLICATION（延迟只读实例）。当从延迟复制变更为普通只读实例时，延迟时间设置为0

	DstRoInstType *string `json:"DstRoInstType,omitempty" name:"DstRoInstType"`
	// 延迟时间（单位：秒），将实例修改为延迟只读实例时必传。最小值1，最大值259200。

	ReplicationDelay *int64 `json:"ReplicationDelay,omitempty" name:"ReplicationDelay"`
}

func (r *ModifyRoTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSmartDBAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，调用DescribeFlow查询任务状态

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSmartDBAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSmartDBAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设置的状态，0 表示成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Flow struct {

	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 任务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 输入数据

	InputData *string `json:"InputData,omitempty" name:"InputData"`
	// 实例 ID，形如：tdsql-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 用户任务类型

	UserTaskType *int64 `json:"UserTaskType,omitempty" name:"UserTaskType"`
}

type IsolateHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *IsolateHourDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateHourDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 节点个数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 购买实例数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 统一网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 统一子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内存大小，单位：GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 储存大小，单位：GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 数据库引擎版本，调用 DescribeDBEngines 查询可用的版本。

	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`
	// 购买的CPU大小，单位为核

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 独享集群id

	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
	// 自定义实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 安全组ID，不传表示不绑定安全组

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 是否支持IPv6

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 标签键值对数组

	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
	// DCN源地域

	DcnRegion *string `json:"DcnRegion,omitempty" name:"DcnRegion"`
	// DCN源实例ID

	DcnInstanceId *string `json:"DcnInstanceId,omitempty" name:"DcnInstanceId"`
	// 平台项目ID

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// cpu架构

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。

	InitParams []*DBParamValue `json:"InitParams,omitempty" name:"InitParams"`
	// RS访问策略, 0-无策略, 1-可用区就近访问

	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
	// 需要回档的源实例ID

	RollbackInstanceId *string `json:"RollbackInstanceId,omitempty" name:"RollbackInstanceId"`
	// 回档时间

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
}

func (r *CreateHourDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHourDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {

	// 标签键key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值value

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例的SerialId

	SerialIds []*string `json:"SerialIds,omitempty" name:"SerialIds"`
	// 分页返回，每页返回的数目，默认为20。取值范围为1到100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 从偏移量Offset开始返回，默认取值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 取值为createtime或者instancename，表示按照创建时间或者实例名进行排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 取值为desc或者asc。desc-表示降序；asc-表示升序

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 集群名称数组

	ClusterNames []*string `json:"ClusterNames,omitempty" name:"ClusterNames"`
	// Zk名称数组

	ZkNames []*string `json:"ZkNames,omitempty" name:"ZkNames"`
	// Vip数组

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// 实例名称数组

	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`
	// 实例ID的数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 按照项目ID进行过滤

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 私有网络ID，形如 vpc-nobbvxkf

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 子网ID，形如 subnet-dnn20os8

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除任务ID。删除任务是异步任务，用FlowId做为参数，调用DescribeFlow接口，查询任务执行状态。

		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`
		// 实例Id，形如tdsql-o0q206pq

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyHourDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseView struct {

	// 视图名称

	View *string `json:"View,omitempty" name:"View"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBRollbackInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回档实例列表总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 回档实例列表。

		Instances []*DBRollbackInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBRollbackInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBRollbackInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DCNReplicaStatus struct {

	// 该只读实例当前的复制状态，枚举值：RUNNING（正在复制中）、PAUSED（已暂停复制）、REACHED_DUE_TIME（已复制到指定时间点）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 该只读实例相对于主实例的延迟情况，单位为秒

	Delay *int64 `json:"Delay,omitempty" name:"Delay"`
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigHistoriesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例配置列表

		InstanceConfigSet []*InstanceConfig `json:"InstanceConfigSet,omitempty" name:"InstanceConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 原价，单位：分

		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
		// 实际价格，单位：分。受折扣等影响，可能和原价不同。计量模式下返回值恒为0。

		Price *int64 `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名称，通过 DescribeDatabases 接口获取。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestCloudDBAReportRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLatestCloudDBAReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestCloudDBAReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeHourDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeHourDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateParamConstraint struct {

	// 约束枚举

	Enum *string `json:"Enum,omitempty" name:"Enum"`
	// 约束范围

	Range *TemplateConstraintRange `json:"Range,omitempty" name:"Range"`
	// 约束字符串

	Str *string `json:"Str,omitempty" name:"Str"`
	// 约束类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DatabaseProcedure struct {

	// 存储过程名称

	Proc *string `json:"Proc,omitempty" name:"Proc"`
}

type CreateHourDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 购买实例异步任务流程ID，以FlowId作为参数调用DescribeFlow API接口，查询创建实例任务执行状态。

		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`
		// 长订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHourDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceHAInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前主所在可用区

		MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
		// 当前切换状态，0-未切换，1-切换中

		SwitchStatus *int64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
		// 当前是否允许切换

		SwitchAllowed *bool `json:"SwitchAllowed,omitempty" name:"SwitchAllowed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceHAInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceHAInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceProxyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关配置参数

		ProxyConfigParamSet []*ProxyConfigParam `json:"ProxyConfigParamSet,omitempty" name:"ProxyConfigParamSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceProxyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceProxyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTmpInstancesRequest struct {
	*tchttp.BaseRequest

	// 回档实例的ID列表，形如：tdsql-ow728lmc。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 回档时间点

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
}

func (r *CreateTmpInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTmpInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyInstanceRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 权限列表。

		Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
		// 数据库账号用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 数据库账号Host

		Host *string `json:"Host,omitempty" name:"Host"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 请求DB的当前参数值

		Params []*ParamDesc `json:"Params,omitempty" name:"Params"`
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

type DBBackupTimeConfig struct {

	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:00

	EndBackupTime *string `json:"EndBackupTime,omitempty" name:"EndBackupTime"`
	// 实例 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00

	StartBackupTime *string `json:"StartBackupTime,omitempty" name:"StartBackupTime"`
}

type DescribeDBSlowLogAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 慢查询SQL出现的开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 慢查询SQL出现的结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 返回的慢查询详情数据。每个点代表一个时间段内慢查询SQL出现的记录次数，一个时间段的长度由请求的开始时间和结束时间的差值决定，小于1天是5分钟一段，大于1天小于7天是30分钟一段，大于7天是2个小时一段。

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 采样周期

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DCNReplicaConfig struct {

	// 只读实例的复制模式，枚举值为：START、STOP

	RoReplicationMode *string `json:"RoReplicationMode,omitempty" name:"RoReplicationMode"`
	// 延迟复制类型。当RoReplicationMode为 START 时有意义，枚举值为 DEFAULT（按照延迟复制时间进行复制，即界面上显示的正常启动）、DUE_TIME（回放到指定时间点）。

	DelayReplicationType *string `json:"DelayReplicationType,omitempty" name:"DelayReplicationType"`
	// 延迟复制类型为 DUE_TIME 有意义，回放的时间点

	DueTime *string `json:"DueTime,omitempty" name:"DueTime"`
	// 复制延迟，单位为秒

	ReplicationDelay *int64 `json:"ReplicationDelay,omitempty" name:"ReplicationDelay"`
}

type MonitorFloatData struct {

	// 监控数据

	Data *float64 `json:"Data,omitempty" name:"Data"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 参数修改结果

		Result []*ParamModifyResult `json:"Result,omitempty" name:"Result"`
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

type DcnDetailItem struct {

	// 实例DCN标志，1-主，2-备

	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
	// 实例DCN状态，0-无，1-创建中，2-同步中，3-已断开

	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 实例IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例IPv6地址

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 实例端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 只读实例复制配置信息，仅当实例为DCN备时有效

	ReplicaConfig *DCNReplicaConfig `json:"ReplicaConfig,omitempty" name:"ReplicaConfig"`
	// 只读实例复制状态，仅当实例为DCN备时有效

	ReplicaStatus *DCNReplicaStatus `json:"ReplicaStatus,omitempty" name:"ReplicaStatus"`
	// KMS开启状态，0-未开启，1-已开启

	EncryptStatus *int64 `json:"EncryptStatus,omitempty" name:"EncryptStatus"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBEnginesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DCN同步详情

		DcnDetails []*DcnDetailItem `json:"DcnDetails,omitempty" name:"DcnDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesArchInfoRequest struct {
	*tchttp.BaseRequest

	// 可用区信息

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

func (r *DescribeZonesArchInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesArchInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecConfigInfo struct {

	// Cpu核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 设备型号

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 数据盘规格最大值，单位 GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 数据盘规格最小值，单位 GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 节点个数，2 表示一主一从，3 表示一主二从

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 产品类型 Id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 最大 Qps 值

	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
	// 推荐的使用场景

	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`
}

type AuthenticateCAMResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前请求账号是否为根账号

		IsRoot *int64 `json:"IsRoot,omitempty" name:"IsRoot"`
		// 当前请求账号是否有全部权限

		IsAllPass *int64 `json:"IsAllPass,omitempty" name:"IsAllPass"`
		// 主账号的appId

		OwnerAppId *int64 `json:"OwnerAppId,omitempty" name:"OwnerAppId"`
		// 权限详情

		Permissions []*Permission `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthenticateCAMResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateCAMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否启用数据加密，开启后暂不支持关闭。本接口的可选值为：1-开启数据加密。

	EncryptEnabled *int64 `json:"EncryptEnabled,omitempty" name:"EncryptEnabled"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 参数模板列表

		TemplateSet []*ConfTemplate `json:"TemplateSet,omitempty" name:"TemplateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可通过 DescribeFlow 查询任务状态。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 透传入参。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否修改

		IsModify *int64 `json:"IsModify,omitempty" name:"IsModify"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceSpec struct {

	// 设备型号

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 该机型对应的可售卖规格列表

	SpecInfos []*SpecConfigInfo `json:"SpecInfos,omitempty" name:"SpecInfos"`
}

type SlowLogData struct {

	// 语句校验和，用于查询详情

	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`
	// 数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 样例Sql

	ExampleSql *string `json:"ExampleSql,omitempty" name:"ExampleSql"`
	// 抽象的SQL语句

	FingerPrint *string `json:"FingerPrint,omitempty" name:"FingerPrint"`
	// 平均的锁时间

	LockTimeAvg *string `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`
	// 最大锁时间

	LockTimeMax *string `json:"LockTimeMax,omitempty" name:"LockTimeMax"`
	// 最小锁时间

	LockTimeMin *string `json:"LockTimeMin,omitempty" name:"LockTimeMin"`
	// 锁时间总和

	LockTimeSum *string `json:"LockTimeSum,omitempty" name:"LockTimeSum"`
	// 查询次数

	QueryCount *string `json:"QueryCount,omitempty" name:"QueryCount"`
	// 平均查询时间

	QueryTimeAvg *string `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`
	// 最大查询时间

	QueryTimeMax *string `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`
	// 最小查询时间

	QueryTimeMin *string `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`
	// 查询时间总和

	QueryTimeSum *string `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`
	// 扫描行数

	RowsExaminedSum *string `json:"RowsExaminedSum,omitempty" name:"RowsExaminedSum"`
	// 发送行数

	RowsSentSum *string `json:"RowsSentSum,omitempty" name:"RowsSentSum"`
	// 最后执行时间

	TsMax *string `json:"TsMax,omitempty" name:"TsMax"`
	// 首次执行时间

	TsMin *string `json:"TsMin,omitempty" name:"TsMin"`
	// 帐号

	User *string `json:"User,omitempty" name:"User"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：mariadb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 项目Id，TCE上恒传0

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 拉取数量限制。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索条件，支持安全组id或者安全组名称。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamDesc struct {

	// 参数限制

	Constraint *ParamConstraint `json:"Constraint,omitempty" name:"Constraint"`
	// 系统默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 是否有设置过值，false:没有设置过值，true:有设置过值。

	HaveSetValue *bool `json:"HaveSetValue,omitempty" name:"HaveSetValue"`
	// 参数名字

	Param *string `json:"Param,omitempty" name:"Param"`
	// 设置过的值，参数生效后，该值和value一样。

	SetValue *string `json:"SetValue,omitempty" name:"SetValue"`
	// 当前参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type SaleZone struct {

	// 可用区英文名

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区数字ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区中文名

	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`
}

type DeleteConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板Id，可以通过 DescribeConfigTemplates 查询参数模板列表获得。

	ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DeleteConfigTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 从结果的第几条数据开始返回

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的结果条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的起始时间，形如2016-07-23 14:55:20

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间，形如2016-08-22 14:55:20

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 要查询的具体数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 排序指标，取值为query_time_sum或者query_count

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型，desc或者asc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 是否查询从机的慢查询，0-主机; 1-从机

	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBTmpInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBTmpInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTmpInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoReplicationModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoReplicationModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoReplicationModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 透传入参

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// Vip可用状态：0 - 可用，300 - 此ip已被子网内其他用户占用， 301 - 此ip不在子网内， 302 - ip格式错误， 303 - 该IP为实例当前IP。

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckIpStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardZoneChooseInfo struct {

	// 主可用区

	MasterZone *ZonesInfo `json:"MasterZone,omitempty" name:"MasterZone"`
	// 可选的从可用区

	SlaveZones []*ZonesInfo `json:"SlaveZones,omitempty" name:"SlaveZones"`
}

type ModifyBackupTimeRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00

	StartBackupTime *string `json:"StartBackupTime,omitempty" name:"StartBackupTime"`
	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:59

	EndBackupTime *string `json:"EndBackupTime,omitempty" name:"EndBackupTime"`
}

func (r *ModifyBackupTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBackupTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchRollbackInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchRollbackInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例名称。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 数据库名称。

		DbName *string `json:"DbName,omitempty" name:"DbName"`
		// 表名称。

		Table *string `json:"Table,omitempty" name:"Table"`
		// 列信息。

		Cols []*TableColumn `json:"Cols,omitempty" name:"Cols"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 参数列表，每一个元素是Param和Value的组合，可以通过 DescribeDBParameters 查询可配置的参数和值限制

	Params []*DBParamValue `json:"Params,omitempty" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
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

type ZonesInfo struct {

	// 可用区英文ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区数字ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区中文名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例详细信息列表

		Instances []*DBInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceSpecsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBInstanceSpecsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceSpecsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例 ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 新的实例名称。允许的字符为字母、数字、下划线、连字符和中文。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 该字段对只读帐号有意义，表示选择主备延迟小于该值的备机

	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
	// 用户备注信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）

	Host *string `json:"Host,omitempty" name:"Host"`
	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DBCharsetItem struct {

	// DB版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// DB默认字符集

	DefaultCharset *string `json:"DefaultCharset,omitempty" name:"DefaultCharset"`
	// DB可选字符集

	OptionalCharsets *string `json:"OptionalCharsets,omitempty" name:"OptionalCharsets"`
}

type CreateTmpInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTmpInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTmpInstancesResponse) FromJsonString(s string) error {
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

type LogFileInfo struct {

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件长度

	Length *uint64 `json:"Length,omitempty" name:"Length"`
	// Log最后修改时间

	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`
	// 下载Log时用到的统一资源标识符

	Uri *string `json:"Uri,omitempty" name:"Uri"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名，由字母、数字、下划线和连字符组成，长度为1~32位。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 账号密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障建议该参数值大于10。当ReadOnly选择1、2时该参数生效。

	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksRequest struct {
	*tchttp.BaseRequest

	// 地域ID列表

	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`
	// 状态列表

	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`
	// 实例ID列表，形如：tdsql-ow728lmc。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 任务类型列表

	FlowTypes []*int64 `json:"FlowTypes,omitempty" name:"FlowTypes"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户任务ID列表

	UTaskIds []*int64 `json:"UTaskIds,omitempty" name:"UTaskIds"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 默认模板名称

	DefaultTemplateName *string `json:"DefaultTemplateName,omitempty" name:"DefaultTemplateName"`
}

func (r *DescribeDefaultConfigTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultConfigTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateHourDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恢复成功的实例id列表

		SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitempty" name:"SuccessInstanceIds"`
		// 恢复失败的实例id列表

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateHourDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSSLAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启实例的SSL认证。0-关闭；1-开启

	SSLEnabled *int64 `json:"SSLEnabled,omitempty" name:"SSLEnabled"`
}

func (r *ModifyInstanceSSLAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSSLAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExclusiveGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableExclusiveGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExclusiveGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogTimeRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBinlogTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceProxyConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceProxyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceProxyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstExtraInfo struct {

	// 如果目标实例通过专线接入的话，表示接入的专线ID

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 目标实例的IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 目标实例的端口地址

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 私用网络ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 目标实例登录用户名

	User *string `json:"User,omitempty" name:"User"`
	// 子网ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {

	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者Id

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 描述信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 是否默认项目，1 是，0 不是

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主账号Id

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 来源AppId

	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`
	// 来源平台

	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`
	// 项目状态,0正常，-1关闭。默认项目为3

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeDBTmpInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时实例

		TmpInstances []*TmpInstance `json:"TmpInstances,omitempty" name:"TmpInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBTmpInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBTmpInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ActivateHourDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateHourDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：mariadb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdsql-o0q206pq

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyHourDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyHourDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest

	// 异步请求接口返回的任务流程号。

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点总个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点信息

		NodesInfo []*NodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceNodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePriceRequest struct {
	*tchttp.BaseRequest

	// 欲新购实例的可用区ID。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例节点个数，可以通过 DescribeDBInstanceSpecs
	//  查询实例规格获得。

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 内存大小，单位：GB，可以通过 DescribeDBInstanceSpecs
	//  查询实例规格获得。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 存储空间大小，单位：GB，可以通过 DescribeDBInstanceSpecs
	//  查询实例规格获得不同内存大小对应的磁盘规格下限和上限。

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 欲购买的时长，单位：月。不传时默认查询1个月的价格。付费模式为 postpaid 时该参数无效，固定查询第一计费阶梯使用1小时的费用。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 欲购买的数量，默认查询购买1个实例的价格。

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 付费类型。postpaid：按量付费；prepaid：预付费。不传时默认为预付费（当前TCE暂不支持预付费功能）

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
}

func (r *DescribePriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：mariadb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 安全组Id。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 实例ID列表，一个或者多个实例Id组成的数组。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组规则。

		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`
		// 符合条件的安全组总数量。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHARequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如 tdsql-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 切换的目标区域，会自动选择该可用区中延迟最低的节点

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDBInstanceHARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamModifyResult struct {

	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 修改参数名字

	Param *string `json:"Param,omitempty" name:"Param"`
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 日志备份天数

		Days *uint64 `json:"Days,omitempty" name:"Days"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyConfigParam struct {

	// 配置参数名称

	Param *string `json:"Param,omitempty" name:"Param"`
	// 配置参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 保存的天数,不能超过30

	Days *uint64 `json:"Days,omitempty" name:"Days"`
}

func (r *ModifyLogFileRetentionPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 慢查询日志数据

		Data []*SlowLogData `json:"Data,omitempty" name:"Data"`
		// 所有语句锁时间总和

		LockTimeSum *float64 `json:"LockTimeSum,omitempty" name:"LockTimeSum"`
		// 所有语句查询总次数

		QueryCount *int64 `json:"QueryCount,omitempty" name:"QueryCount"`
		// 总记录数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 所有语句查询时间总和

		QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {

	// 约束类型为enum时的可选值列表

	Enum *string `json:"Enum,omitempty" name:"Enum"`
	// 约束类型为section时的范围

	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`
	// 约束类型为string时的可选值列表

	String *string `json:"String,omitempty" name:"String"`
	// 约束类型,如枚举enum，区间section

	Type *string `json:"Type,omitempty" name:"Type"`
}

type TemplateParam struct {

	// 参数约束条件

	Constraint *TemplateParamConstraint `json:"Constraint,omitempty" name:"Constraint"`
	// 参数默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 参数名

	Param *string `json:"Param,omitempty" name:"Param"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type MonitorIntData struct {

	// 监控数据

	Data *int64 `json:"Data,omitempty" name:"Data"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例Id过滤列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例名称过滤列表

	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`
	// 计费类型过滤列表，1-包年包月，0-按量计费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// VPC网络Id信息过滤列表

	VpcIds []*int64 `json:"VpcIds,omitempty" name:"VpcIds"`
	// 子网Id信息过滤列表

	SubnetIds []*int64 `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// VPC网络统一Id信息过滤列表

	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`
	// 子网统一Id信息过滤列表

	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`
	// 实例状态过滤列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 实例Vip信息列表

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// 排序字段，支持ProjectId, InstanceName, CreateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持asc，desc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 数量限制，默认推荐100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCommonDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchRollbackInstanceRequest struct {
	*tchttp.BaseRequest

	// 源/旧实例ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标实例ID，形如：tdsql-ow728lmc。

	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
}

func (r *SwitchRollbackInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchRollbackInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源用户名

	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`
	// 源用户允许的访问 host

	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`
	// 源账号的 ReadOnly 属性

	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`
	// 目的用户名

	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`
	// 目的用户允许的访问 host

	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`
	// 目的账号的 ReadOnly 属性

	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceVportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库名称，通过 DescribeDatabases 接口获取。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 表名称，通过 DescribeDatabaseObjects 接口获取。

	Table *string `json:"Table,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type McLog struct {

	// 用户请求action(用户操作)

	Action *string `json:"Action,omitempty" name:"Action"`
	// 应用Id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 参数数据

	ArgData *string `json:"ArgData,omitempty" name:"ArgData"`
	// 浏览器类型

	BrowserType *string `json:"BrowserType,omitempty" name:"BrowserType"`
	// 客户端Ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 错误码

	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例 ID，形如：tdsql-ow728lmc

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 操作时间

	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`
	// 主账户Uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 地域Id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 账户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数

	Type *string `json:"Type,omitempty" name:"Type"`
	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空

	Object *string `json:"Object,omitempty" name:"Object"`
	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限

	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详细信息列表

		InstanceDetails []*CommonDBInstance `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
		// 符合条件的实例数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCommonDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *DescribeDBInstanceDetailResponseParams `json:"Response"`
}

type DescribeDBInstanceDetailResponseParams struct {
	// 实例Id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 实例状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 实例目前运行状态描述

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 内网 IP 地址

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 内网端口

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 是否临时实例，0为否，非0为是

		IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`
		// 节点数，2为一主一从，3为一主二从

		NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
		// 实例所在地域名称，如 ap-shanghai

		Region *string `json:"Region,omitempty" name:"Region"`
		// 实例所在可用区名称，如 ap-shanghai-1

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 字符串型的私有网络Id

		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
		// 字符串型的私有网络子网Id

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中

		WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`
		// 外网访问的域名，公网可解析

		WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`
		// 外网 IP 地址，公网可访问

		WanVip *string `json:"WanVip,omitempty" name:"WanVip"`
		// 外网端口

		WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`
		// 实例所属项目 Id

		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
		// TDSQL 版本信息

		TdsqlVersion *string `json:"TdsqlVersion,omitempty" name:"TdsqlVersion"`
		// 实例内存大小，单位 GB

		Memory *int64 `json:"Memory,omitempty" name:"Memory"`
		// 实例存储大小，单位 GB

		Storage *int64 `json:"Storage,omitempty" name:"Storage"`
		// 主可用区，如 ap-shanghai-1

		MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
		// 从可用区列表，如 [ap-shanghai-2]

		SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`
		// 自动续费标志：0 否，1 是

		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
		// 独享集群Id，普通实例为空

		ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
		// 付费模式：prepaid 表示预付费

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 实例创建时间，格式为 2006-01-02 15:04:05

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 实例是否支持审计

		IsAuditSupported *bool `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`
		// 实例到期时间，格式为 2006-01-02 15:04:05

		PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
		// 机型信息

		Machine *string `json:"Machine,omitempty" name:"Machine"`
		// 存储空间使用率

		StorageUsage *string `json:"StorageUsage,omitempty" name:"StorageUsage"`
		// 日志存储空间大小，单位 GB

		LogStorage *int64 `json:"LogStorage,omitempty" name:"LogStorage"`
		// 是否支持数据加密。1-支持；0-不支持

		IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`
		// 内网IPv6

		Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
		// 实例Cpu核数

		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
		// 产品类型ID

		Pid *int64 `json:"Pid,omitempty" name:"Pid"`
		// 最大QPS

		Qps *int64 `json:"Qps,omitempty" name:"Qps"`
		// 是否支持IPv6

		Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
		// 外网IPv6地址，公网可访问

		WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`
		// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中

		WanStatusIpv6 *int64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`
		// 外网IPv6端口

		WanPortIpv6 *int64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`
		// 数据库引擎

		DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`
	// 数据库版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// 数据库版本 ID

	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`
	// 标签信息

		ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
		// DCN标志，0-无，1-主实例，2-灾备实例

		DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
		// DCN状态，0-无，1-创建中，2-同步中，3-已断开

		DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
		// DCN灾备实例数

		DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`
		// 平台项目ID

		PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
		// 平台项目名称

		PlatformProjectName *string `json:"PlatformProjectName,omitempty" name:"PlatformProjectName"`
		// 只读实例复制配置信息，仅当实例为DCN备时有效

		ReplicaConfig *DCNReplicaConfig `json:"ReplicaConfig,omitempty" name:"ReplicaConfig"`
		// 只读实例复制状态，仅当实例为DCN备时有效

		ReplicaStatus *DCNReplicaStatus `json:"ReplicaStatus,omitempty" name:"ReplicaStatus"`
		// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）

		InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
		// 实例的各个DB节点信息

		NodesInfo []*NodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo"`
		// cpu架构

		CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 就近访问策略，0-无策略, 1-可用区就近访问

	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

func (r *DescribeDBInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogAnalysisRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 要查询的慢查询语句对应的数据库名称

	Db *string `json:"Db,omitempty" name:"Db"`
	// 要查询的慢查询语句对应的用户名称

	User *string `json:"User,omitempty" name:"User"`
	// 要查询的慢查询语句的校验和，可以通过查询慢查询日志列表获得

	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`
	// 是否查询从机的慢查询。0-主机；1-从机

	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSlowLogAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// 实例对应的AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 数据库版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// 实例数字ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例短ID，形如tdsql-hdbirf67

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例是一主一从还是一主两从。2-一主一从；3-一主两从

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 实例最初的SerialId，实例扩容后SerialId会变，OriginSerialId表示实例最初的SerialId，不会改变

	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例SerialId

	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`
	// 实例状态信息

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例是VPC网络的话，表示实例所属子网ID，数字ID

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例是VPC网络的话，表示实例子网ID，英文短ID

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 实例是VPC网络的话，表示实例VPC网络ID，英文短ID

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 实例Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例是VPC网络的话，表示实例VPC网络ID，数字ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 实例端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例对应的ZK名称

	ZkName *string `json:"ZkName,omitempty" name:"ZkName"`
	// 实例可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 表示这个实例监控页面如何展示

	MonitorType *int64 `json:"MonitorType,omitempty" name:"MonitorType"`
}

type DeleteTmpInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务流程ID。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTmpInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTmpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// 请求日志总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 包含uri、length、mtime（修改时间）等信息

		Files []*LogFileInfo `json:"Files,omitempty" name:"Files"`
		// 如果是VPC网络的实例，做用本前缀加上URI为下载地址

		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`
		// 如果是普通网络的实例，做用本前缀加上URI为下载地址

		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务Id，可通过 DescribeFlow 查询任务状态。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonDBInstance struct {

	// 应用Id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 0-基础网络，1-VPC网络

	NetType *int64 `json:"NetType,omitempty" name:"NetType"`
	// 计费类型，1-包年包月，0-按量计费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 项目Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例状态信息，0-创建中,1-运行中，2-隔离中,下线后无法拉取实例信息

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网统一Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// VPC网络IP

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// VPC网络统一Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC网络端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DBEngineInfo struct {

	// 引擎描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 引擎名称，eg. 基于MariaDB 10.1.9 分支

	Name *string `json:"Name,omitempty" name:"Name"`
	// 引擎类型，MariaDB|Percona

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎版本，eg. 10.1.9

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 希望转到的VPC网络的VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 希望转到的VPC网络的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 如果需要指定VIP，填上该字段

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 如果需要指定VIPv6，填上该字段

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSSLAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceSSLAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSSLAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FenceInfoItem struct {

	// 独享资源池ID

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
}

type ZoneChooseInfo struct {

	// 主可用区

	MasterZone *ZonesInfo `json:"MasterZone,omitempty" name:"MasterZone"`
	// 可选的从可用区

	SlaveZones []*ZonesInfo `json:"SlaveZones,omitempty" name:"SlaveZones"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例总的个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
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

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 入站规则。

	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`
	// 出站规则。

	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组ID

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	// 安全组备注

	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
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

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：mariadb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要绑定的安全组ID，类似sg-efil73jd。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 被绑定的实例ID，类似tdsql-lesecurk，支持指定多个实例。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlushBinlogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlushBinlogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FlushBinlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例用户列表。

		Users []*DBAccount `json:"Users,omitempty" name:"Users"`
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

type DescribeLatestCloudDBAReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次检测的有效分数

		Score *int64 `json:"Score,omitempty" name:"Score"`
		// 本次检测报告的名称

		ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
		// 本次检测报告的详情

		ReportDetail *string `json:"ReportDetail,omitempty" name:"ReportDetail"`
		// 本次检测报告的完成时间

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestCloudDBAReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestCloudDBAReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceVportRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例VPORT

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *ModifyInstanceVportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceVportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Permission struct {

	// 是否有权限，1 - 有权限， 0 - 无权限

	IsPermitted *int64 `json:"IsPermitted,omitempty" name:"IsPermitted"`
	// 资源对象

	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type DescribeConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板Id，可以通过 DescribeConfigTemplates 查询参数模板列表获得。

	ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DescribeConfigTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可售卖地域信息列表

		RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBInstance struct {

	// 实例所属应用 ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 自动续费标志：0 否，1 是

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 实例CPU核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 实例创建时间，格式为 2006-01-02 15:04:05

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 数据库引擎

	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`
	// 数据库版本

	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`
	// DCN灾备实例数

	DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`
	// DCN标志，0-无，1-主实例，2-灾备实例

	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`
	// DCN状态，0-无，1-创建中，2-同步中，3-已断开

	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`
	// 独享集群ID，为空表示为普通实例

	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`
	// 数字实例ID（过时字段，请勿依赖该值）

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例 ID，唯一标识一个 TDSQL 实例

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称，用户可修改

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 该实例是否支持审计。1-支持；0-不支持

	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`
	// 是否支持数据加密。1-支持；0-不支持

	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`
	// 是否临时实例，0为否，非0为是

	IsTmp *uint64 `json:"IsTmp,omitempty" name:"IsTmp"`
	// 实例处于异步任务时的异步任务流程ID

	Locker *int64 `json:"Locker,omitempty" name:"Locker"`
	// 机器型号

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 实例内存大小，单位 GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点数，2为一主一从，3为一主二从

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 原始实例ID（过时字段，请勿依赖该值）

	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`
	// 付费模式

	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`
	// 实例到期时间，格式为 2006-01-02 15:04:05

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 产品类型 ID

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 实例所属项目 ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 最大 Qps 值

	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
	// 实例所在地域名称，如 ap-shanghai

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例状态：0 创建中，1 流程处理中， 2 运行中，3 实例未初始化，-1 实例已隔离，-2 实例已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例目前运行状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 实例存储大小，单位 GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 子网 ID，基础网络时为 0

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// TDSQL 版本信息

	TdsqlVersion *string `json:"TdsqlVersion,omitempty" name:"TdsqlVersion"`
	// 实例所属账号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 字符串型的私有网络子网ID

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 字符串型的私有网络ID

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 实例最后更新时间，格式为 2006-01-02 15:04:05

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 内网 IP 地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 内网IPv6

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 私有网络 ID，基础网络时为 0

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 内网端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 外网访问的域名，公网可解析

	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`
	// 外网端口

	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`
	// 外网IPv6端口

	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`
	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中

	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`
	// 外网IPv6状态

	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`
	// 外网 IP 地址，公网可访问

	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`
	// 外网IPv6

	WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`
	// 实例所在可用区名称，如 ap-shanghai-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）

	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// cpu架构

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 数据库版本 ID

	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`
	// 标签信息

	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type Rsip struct {

	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type AuthenticateCAMRequest struct {
	*tchttp.BaseRequest

	// 待鉴权的接口名，格式为产品名:API名称，如 mariadb:DescribeDBInstances

	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
	// 待鉴权资源信息，cam 六段式资源格式或 *

	Resources []*string `json:"Resources,omitempty" name:"Resources"`
}

func (r *AuthenticateCAMRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthenticateCAMRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 登录用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。

	Host *string `json:"Host,omitempty" name:"Host"`
	// 数据库名。如果为 \*，表示设置全局权限（即 \*.\*），此时忽略 Type 和 Object 参数。当DbName不为\*时，需要传入参 Type。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数

	Type *string `json:"Type,omitempty" name:"Type"`
	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空

	Object *string `json:"Object,omitempty" name:"Object"`
	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权

	ColName *string `json:"ColName,omitempty" name:"ColName"`
	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER
	// 表/视图权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER
	// 存储过程/函数权限： ALTER ROUTINE，EXECUTE
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE

	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest

	// 灾备实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIpStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 目标子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 目标虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 是否IPv6，默认0

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *CheckIpStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerAccessStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// RS就近模式, 0-无策略, 1-可用区就近访问

	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitempty" name:"RsAccessStrategy"`
}

func (r *ModifyRealServerAccessStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRealServerAccessStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaleSpecSet struct {

	// 地域名称，如ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：dcdbt-ow7t8lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableColumn struct {

	// 列名称

	Col *string `json:"Col,omitempty" name:"Col"`
	// 列类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：mariadb。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSmartDBARequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartSmartDBARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSmartDBARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBRollbackInstance struct {

	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态：0 创建中，1 流程处理中， 2 运行中，3 实例未初始化，-1 实例已隔离，4 实例初始化中，5 实例删除中，6 实例重启中，7 数据迁移中。

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 状态中文描述。

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 地域描述。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例所在可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 节点数，2 为一主一从， 3 为一主二从。

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// Cpu核数。

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小，单位 GB。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 存储大小，单位 GB。

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 字符串型的私有网络ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 字符串型的私有网络子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内网IP。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 内网端口。

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// Ipv6标志。

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 内网IPv6。

	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
	// 付费模式。

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 到期时间。

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 实例回档时间。

	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
	// 实例创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 1：临时实例，0:非临时实例。

	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`
	// 实例标签信息。

	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse

	Response *DescribeFlowResponseParams `json:"Response"`
}

type DescribeFlowResponseParams struct {
	// 流程状态，0：成功，1：失败，2：运行中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigParam struct {

	// 参数名

	Param *string `json:"Param,omitempty" name:"Param"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源用户账户名

	SrcUser *string `json:"SrcUser,omitempty" name:"SrcUser"`
	// 源用户HOST

	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`
	// 目的用户账户名

	DstUser *string `json:"DstUser,omitempty" name:"DstUser"`
	// 目的用户HOST

	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`
	// 目的用户账户描述

	DstDesc *string `json:"DstDesc,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloneAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户允许的访问 host

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBCharsetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBCharsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBCharsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRsipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 后端RS信息

		Rsips []*Rsip `json:"Rsips,omitempty" name:"Rsips"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRsipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRsipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSSLAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例SSL认证功能当前状态。1-开启中；2-已开启；3-已关闭；4-关闭中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSSLAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSSLAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		FlowSet []*Flow `json:"FlowSet,omitempty" name:"FlowSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateHourDBInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 销毁成功的实例id列表

		SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitempty" name:"SuccessInstanceIds"`
		// 销毁失败的实例id列表

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateHourDBInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例 ID，形如：tdsql-ow728lmc。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogFileRetentionPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 可选择的主可用区和从可用区

	AvailableChoice []*ZoneChooseInfo `json:"AvailableChoice,omitempty" name:"AvailableChoice"`
	// 地域英文ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域数字ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区列表

	ZoneList []*ZonesInfo `json:"ZoneList,omitempty" name:"ZoneList"`
}

type TemplateConstraintRange struct {

	// 约束范围最大值

	Max *string `json:"Max,omitempty" name:"Max"`
	// 约束范围最小值

	Min *string `json:"Min,omitempty" name:"Min"`
}

type DescribeDBEnginesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// items总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DB引擎信息

		Items []*DBEngineInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEnginesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRealServerAccessStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRealServerAccessStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRealServerAccessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，根据此FlowId通过DescribeFlow接口查询任务进行状态

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlushBinlogRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushBinlogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FlushBinlogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseTable struct {

	// 表名

	Table *string `json:"Table,omitempty" name:"Table"`
}

type DescribeConfigTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数模板Id

		ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
		// 应用Id

		AppId *int64 `json:"AppId,omitempty" name:"AppId"`
		// 参数模板名称

		TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
		// 参数模板描述

		TemplateDesc *string `json:"TemplateDesc,omitempty" name:"TemplateDesc"`
		// 默认模板

		TemplateDefault *string `json:"TemplateDefault,omitempty" name:"TemplateDefault"`
		// 模板参数列表

		TemplateParamSet []*TemplateParam `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSSLAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceSSLAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceSSLAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseFunction struct {

	// 函数名称

	Func *string `json:"Func,omitempty" name:"Func"`
}

type TmpInstance struct {

	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例备注

	InstanceRemark *string `json:"InstanceRemark,omitempty" name:"InstanceRemark"`
	// 实例IPv6标志

	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 有效期结束时间

	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 源实例 ID，形如：tdsql-ow728lmc。

	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`
	// 实例状态,0:待初始化,1:流程处理中,2:有效状态,-1:已隔离，-2：已下线

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 0:非临时实例 ,1:无效临时实例, 2:回档成功的有效临时实例

	TempType *int64 `json:"TempType,omitempty" name:"TempType"`
	// 实例虚IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例虚IPv6

	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`
	// 实例虚端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 实例所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeDBInstanceRsipRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceRsipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceRsipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 实例Id形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdsql-6ltok4u9

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 一次最多返回多少条数据。默认为无穷大，返回符合要求的所有数据

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数据的偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceNodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceNodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDbInfo struct {

	// 同步目标信息

	DstInfo *string `json:"DstInfo,omitempty" name:"DstInfo"`
	// 同步源信息

	OrgInfo *string `json:"OrgInfo,omitempty" name:"OrgInfo"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该实例上的数据库列表。

		Databases []*Database `json:"Databases,omitempty" name:"Databases"`
		// 透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBCharsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DBCharsets总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DB字符集信息

		DBCharsets *DBCharsetItem `json:"DBCharsets,omitempty" name:"DBCharsets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBCharsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBCharsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

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

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// 待修改同步模式的实例ID。形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 同步模式：0 异步，1 强同步， 2 强同步可退化

	SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTmpInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTmpInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTmpInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 待初始化的实例ID列表，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。

	Params []*DBParamValue `json:"Params,omitempty" name:"Params"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableExclusiveGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享资源池数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 独享资源池信息

		Items []*FenceInfoItem `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableExclusiveGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableExclusiveGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例 ID 查询。实例 ID 形如：tdsql-ow728lmc。每次请求的实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。

	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`
	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 按项目 ID 查询

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 是否根据 VPC 网络来搜索

	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`
	// 私有网络 ID， IsFilterVpc 为 1 时有效

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 排序字段， projectId， createtime， instancename 三者之一

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型， desc 或者 asc

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 偏移量，默认为 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为 20，最大值为 100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按 OriginSerialId 查询

	OriginSerialIds []*string `json:"OriginSerialIds,omitempty" name:"OriginSerialIds"`
	// 标识是否使用ExclusterType字段, false不使用，true使用

	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`
	// 实例所属独享集群类型。取值范围：1-非独享集群，2-独享集群， 0-全部

	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`
	// 按独享集群ID过滤实例，独享集群ID形如dbdc-4ih6uct9

	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds"`
	// 按标签key查询

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
	// 实例类型过滤，1-独享实例，2-主实例，3-灾备实例，多个按逗号分隔

	FilterInstanceType *string `json:"FilterInstanceType,omitempty" name:"FilterInstanceType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultConfigTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数列表

		TemplateParamSet []*TemplateParam `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultConfigTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesArchInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用Cpu架构信息

		CpuArch []*string `json:"CpuArch,omitempty" name:"CpuArch"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesArchInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesArchInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 默认模板，默认值default

	TemplateDefault *string `json:"TemplateDefault,omitempty" name:"TemplateDefault"`
	// 配置模板参数

	ConfigParams []*ConfigParam `json:"ConfigParams,omitempty" name:"ConfigParams"`
}

func (r *CreateConfigTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的配置数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例备份时间配置信息

		Items []*DBBackupTimeConfig `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDBInstanceHAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchDBInstanceHAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchDBInstanceHAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigTemplatesRequest struct {
	*tchttp.BaseRequest

	// 模板名称模糊搜索

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 是否返回默认模板，1:返回,0:不返回

	WithDefaultTemplate *int64 `json:"WithDefaultTemplate,omitempty" name:"WithDefaultTemplate"`
	// 分页，页大小，默认150

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页，页偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConfigTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 数据库名称。

		DbName *string `json:"DbName,omitempty" name:"DbName"`
		// 表列表。

		Tables []*DatabaseTable `json:"Tables,omitempty" name:"Tables"`
		// 视图列表。

		Views []*DatabaseView `json:"Views,omitempty" name:"Views"`
		// 存储过程列表。

		Procs []*DatabaseProcedure `json:"Procs,omitempty" name:"Procs"`
		// 函数列表。

		Funcs []*DatabaseFunction `json:"Funcs,omitempty" name:"Funcs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步模式：0 异步，1 强同步， 2 强同步可退化

		SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
		// 是否有修改流程在执行中：1 是， 0 否。

		IsModifying *int64 `json:"IsModifying,omitempty" name:"IsModifying"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ModifyRoReplicationModeRequest struct {
	*tchttp.BaseRequest

	// 只读实例 ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 对只读实例的复制模式进行变更，枚举值为：START和STOP。

	RoReplicationMode *string `json:"RoReplicationMode,omitempty" name:"RoReplicationMode"`
	// 延迟复制类型。当RoReplicationMode为“START”时必选，枚举值为 DEFAULT（按照延迟复制时间进行复制，即界面上显示的正常启动）、DUE_TIME（回放到指定时间点）。

	DelayReplicationType *string `json:"DelayReplicationType,omitempty" name:"DelayReplicationType"`
	// 指定时间点，最大值不能超过当前时间。

	DueTime *string `json:"DueTime,omitempty" name:"DueTime"`
}

func (r *ModifyRoReplicationModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoReplicationModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSaleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSaleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigTemplateRequest struct {
	*tchttp.BaseRequest

	// 参数模板Id，可以通过 DescribeConfigTemplates 查询参数模板列表获得。

	ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
	// 配置模板参数列表

	ConfigParams []*ConfigParam `json:"ConfigParams,omitempty" name:"ConfigParams"`
}

func (r *ModifyConfigTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 内存大小，单位：GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 存储大小，单位：GB

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 升级后实例的cpu大小， 单位核，其值不能小于当前实例的cpu

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。

	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`
	// 切换结束时间,  格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。

	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
	// 是否自动重试。 0：不自动重试  1：自动重试

	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitempty" name:"SwitchAutoRetry"`
}

func (r *UpgradeHourDBInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeHourDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID，透传入参。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 用户名，透传入参。

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 允许访问的 host，透传入参。

		Host *string `json:"Host,omitempty" name:"Host"`
		// 透传入参。

		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTimeRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeBackupTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID，形如：tdsql-ow728lmc。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfTemplate struct {

	// 应用id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 参数模板Id

	ConfigTemplateId *int64 `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
	// 默认模板

	TemplateDefault *string `json:"TemplateDefault,omitempty" name:"TemplateDefault"`
	// 参数模板描述

	TemplateDesc *string `json:"TemplateDesc,omitempty" name:"TemplateDesc"`
	// 参数模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

type Database struct {

	// 数据库名称

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例ID列表。实例 ID 形如：tdsql-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

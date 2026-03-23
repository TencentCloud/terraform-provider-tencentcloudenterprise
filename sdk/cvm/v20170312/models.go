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

type LiveMigrateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 目的实例外网IP。

		Ip *string `json:"Ip,omitempty" name:"Ip"`
		// 目的子机密码。

		Password *string `json:"Password,omitempty" name:"Password"`
		// 目的实例内网IP。

		LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisks struct {

	// 数据盘镜像cos url

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 数据盘大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 数据盘对应设备名，目前是DiskId

	Device *string `json:"Device,omitempty" name:"Device"`
}

type DescribeUserMigrateTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，可选instance-id，job-name，job-id。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询的迁移任务类型，可选`ColdMigrateInstance`，`ImportCbs`和`OfflineMigrate`，其中`OfflineMigrate`代表两种离线迁移任务一起查询。

	MigrateTaskType *string `json:"MigrateTaskType,omitempty" name:"MigrateTaskType"`
	// 任务数量限制，用于分页。默认20，最大50

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 任务起始数，用于分页。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserMigrateTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserMigrateTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 [DescribeInstances](DescribeInstances) API返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指定有效的[镜像](/tcloud/Compute/CVM/292128/835305/mirr_overview)ID，格式形如`img-xxx`。镜像类型分为三种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](//console.{{conf.main_domain}}/cvm/image/list?imageType=PUBLIC_IMAGE&pageIndex=1&pageSize=20)查询；</li><li>通过调用接口 [DescribeImages](../镜像相关接口/DescribeImages) ，取返回信息中的`ImageId`字段。</li>

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。如："SystemDisk":{"DiskSize":xxx}

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *InquiryPriceResetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMarketImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMarketImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMarketImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 是否跳过实际执行逻辑。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *RenewHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatisticsRequest struct {
	*tchttp.BaseRequest

	// region地域参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的VNC连接串

		InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyQuota struct {

	// 实例所属的可用区ID，按照可用区过滤。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例计费模式。支持类型:PREPAID、POSTPAID_BY_HOUR、SPOTPAID. 最大限制为10，value最大限制为5

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 当前数量

	QuotaCurrent *uint64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`
	// 配额数量

	QuotaLimit *uint64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 1

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例的CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例<br><li>EXPIRED：表示过期的实例<br><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。

	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例系统盘信息。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘信息。只包含随实例购买的数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 实例主网卡的内网`IP`列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例主网卡的公网`IP`列表。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 实例带宽信息。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 实例所属虚拟私有网络信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 生产实例所使用的镜像`ID`。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费<br><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费<br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 实例状态。取值范围：<br><li>PENDING：表示创建中<br></li><li>LAUNCH_FAILED：表示创建失败<br></li><li>RUNNING：表示运行中<br></li><li>STOPPED：表示关机<br></li><li>STARTING：表示开机中<br></li><li>STOPPING：表示关机中<br></li><li>REBOOTING：表示重启中<br></li><li>SHUTDOWN：表示停止待销毁<br></li><li>TERMINATING：表示销毁中。<br></li>

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 实例的最新操作。例：StopInstances、ResetInstance。

	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`
	// 实例的最新操作状态。取值范围：<br><li>SUCCESS：表示操作成功<br><li>OPERATING：表示操作执行中<br><li>FAILED：表示操作失败

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 实例最新操作的唯一请求 ID。

	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](/tcloud/api/NetWork/VPC/APIs/私有网络（vpc）/版本（2017-03-12）/安全组相关接口/DescribeSecurityGroups) 的返回值中的sgId字段来获取。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例登录设置。目前只返回实例所关联的密钥。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 分散置放群组ID。

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 实例关联的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 实例的uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 实例的os名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 实例的ipv6地址

	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
	// 实例的关机计费模式。
	// 取值范围：<br><li>KEEP_CHARGING：关机继续收费<br><li>STOP_CHARGING：关机停止收费<li>NOT_APPLICABLE：实例处于非关机状态或者不适用关机停止计费的条件<br>

	StopChargingMode *string `json:"StopChargingMode,omitempty" name:"StopChargingMode"`
	// CAM角色名。

	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
	// IsolatedSource

	IsolatedSource *string `json:"IsolatedSource,omitempty" name:"IsolatedSource"`
	// 资源所属项目Id

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 持续时间

	RemainTime *string `json:"RemainTime,omitempty" name:"RemainTime"`
	// 高性能计算集群`ID`。

	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`
	// 默认登录用户

	DefaultLoginUser *string `json:"DefaultLoginUser,omitempty" name:"DefaultLoginUser"`
	// 默认登录端口

	DefaultLoginPort *uint64 `json:"DefaultLoginPort,omitempty" name:"DefaultLoginPort"`
	// 实例所在的专用集群`ID`。

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 操作者Uin

	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
	// 高性能计算集群`IP`列表

	RdmaIpAddresses []*string `json:"RdmaIpAddresses,omitempty" name:"RdmaIpAddresses"`
	// 实例id

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 实例类型

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 隔离时间

	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`
	// 错误的key

	ErrorKey *string `json:"ErrorKey,omitempty" name:"ErrorKey"`
	// GPU

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
	// 新创建识别

	NewCreationIdentify *bool `json:"NewCreationIdentify,omitempty" name:"NewCreationIdentify"`
	// 操作掩码

	OperationMask *uint64 `json:"OperationMask,omitempty" name:"OperationMask"`
	// 镜像类型

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 密钥id,&nbsp;数据来源为MC、BATCH、AS、tke、VPC_CGW，或者【请求里即没有Token又没有CamContext】时，会返回KeyPairIds字段，其他数据来源不会返回KeyPairIds字段

	KeyPairIds []*string `json:"KeyPairIds,omitempty" name:"KeyPairIds"`
	// 运行标志

	RunFlag *int64 `json:"RunFlag,omitempty" name:"RunFlag"`
	// 内部vpc&nbsp;Id

	InnerVpcId *uint64 `json:"InnerVpcId,omitempty" name:"InnerVpcId"`
	// 实例的Family

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 安全隔离信息

	SafeIsolatedInfo *string `json:"SafeIsolatedInfo,omitempty" name:"SafeIsolatedInfo"`
	// 是否安全隔离

	IsSafeIsolated *bool `json:"IsSafeIsolated,omitempty" name:"IsSafeIsolated"`
	// 虚拟化

	Hypervisor *int64 `json:"Hypervisor,omitempty" name:"Hypervisor"`
	// 实例的最新操作错误信息。

	LatestOperationErrorMsg *string `json:"LatestOperationErrorMsg,omitempty" name:"LatestOperationErrorMsg"`
	// 包销实例到期时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。注意：后付费模式本项为null

	UnderwriteExpiredTime *string `json:"UnderwriteExpiredTime,omitempty" name:"UnderwriteExpiredTime"`
	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：&nbsp;&nbsp;true：表示开启实例保护，不允许通过api接口删除实例&nbsp;false：表示关闭实例保护，允许通过api接口删除实例&nbsp;&nbsp;默认取值：false。

	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
	// 实例是否配置了休眠

	HibernationOptions *HibernationOptions `json:"HibernationOptions,omitempty" name:"HibernationOptions"`
	// 网卡Trunking状态

	VmEniTrunking *string `json:"VmEniTrunking,omitempty" name:"VmEniTrunking"`
	// 实例的启动模式

	BootMode *string `json:"BootMode,omitempty" name:"BootMode"`
	// CHC实例的类型

	ChcInstanceType *string `json:"ChcInstanceType,omitempty" name:"ChcInstanceType"`
	// 实例的交换盘信息列表

	SwapDisks *SwapDisks `json:"SwapDisks,omitempty" name:"SwapDisks"`
	// 最新操作的错误码，表示操作失败的原因。

	LatestOperationCode *string `json:"LatestOperationCode,omitempty" name:"LatestOperationCode"`
	// 置放群组列表

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 实例的操作系统许可类型

	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
	// 是否启用Grid许可证

	EnableGridLicence *string `json:"EnableGridLicence,omitempty" name:"EnableGridLicence"`
}

type HibernationOptions struct {

	// 是否配置了休眠选项

	Configured *bool `json:"Configured,omitempty" name:"Configured"`
}

type SwapDisks struct {

	// 系统盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘大小（GB）

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘唯一标识

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 是否加密

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 加密密钥ID

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
	// 吞吐性能（MB/s）

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 独享集群ID

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分散置放群组id。

		DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
		// 分散置放群组类型，与入参一致。

		Type *string `json:"Type,omitempty" name:"Type"`
		// 分散置放群组名称，长度1-60个字符，支持中、英文。

		Name *string `json:"Name,omitempty" name:"Name"`
		// 置放群组内可容纳的云服务器数量。

		CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`
		// 置放群组内已有的云服务器数量。

		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`
		// 置放群组创建时间。

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 创建置放群组的appId

		Owner *string `json:"Owner,omitempty" name:"Owner"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例详细信息列表。

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
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

type ExitLiveMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 需要退出服务热迁移的实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 服务迁移是否成功。

	MigrateResult *string `json:"MigrateResult,omitempty" name:"MigrateResult"`
}

func (r *ExitLiveMigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitLiveMigrateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组ID，可使用[DescribeDisasterRecoverGroups](DescribeDisasterRecoverGroups)接口获取。

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 分散置放群组名称，长度1-60个字符，支持中、英文。

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTypeConfigSetPrice struct {

	// 后续合计费用的原价，后付费模式使用，单位：元。如返回了其他时间区间项，如UnitPriceSecondStep，则本项代表时间区间在(0, 96)小时；若未返回其他时间区间项，则本项代表全时段，即(0, ∞)小时注意：此字段可能返回 null，表示取不到有效值。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后续计价单元，后付费模式使用，可取值范围：HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。注意：此字段可能返回 null，表示取不到有效值。

	ChargeUnit *float64 `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预支合计费用的原价，预付费模式使用，单位：元。注意：此字段可能返回 null，表示取不到有效值。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支合计费用的折扣价，预付费模式使用，单位：元。注意：此字段可能返回 null，表示取不到有效值。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 预支一年合计费用的原价，预付费模式使用，单位：元。注意：此字段可能返回 null，表示取不到有效值。注意：此字段可能返回 null，表示取不到有效值。

	OriginalPriceOneYear *float64 `json:"OriginalPriceOneYear,omitempty" name:"OriginalPriceOneYear"`
	// 预支一年合计费用的折扣价，预付费模式使用，单位：元。注意：此字段可能返回 null，表示取不到有效值。注意：此字段可能返回 null，表示取不到有效值。

	DiscountPriceOneYear *float64 `json:"DiscountPriceOneYear,omitempty" name:"DiscountPriceOneYear"`
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 定时任务ID列表

	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
	// 实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 定时任务操作

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// 根据时间范围查询,结束时间

	EndActionTime *string `json:"EndActionTime,omitempty" name:"EndActionTime"`
	// 根据时间范围查询， 开始时间

	StartActionTime *string `json:"StartActionTime,omitempty" name:"StartActionTime"`
	// 状态列表， 可选状态“UNDO”，“DONE”

	StatusList []*string `json:"StatusList,omitempty" name:"StatusList"`
}

func (r *QueryInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps"`
}

func (r *ReturnNormalAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNormalAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchTemplateInfo struct {

	// 实例启动模版本号。 注意：此字段可能返回 null，表示取不到有效值。

	LatestVersionNumber *int64 `json:"LatestVersionNumber,omitempty" name:"LatestVersionNumber"`
	// 实例启动模板ID。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
	// 实例启动模板名。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`
	// 实例启动模板默认版本号。 注意：此字段可能返回 null，表示取不到有效值。

	DefaultVersionNumber *int64 `json:"DefaultVersionNumber,omitempty" name:"DefaultVersionNumber"`
	// 实例启动模板包含的版本总数量。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateVersionCount *int64 `json:"LaunchTemplateVersionCount,omitempty" name:"LaunchTemplateVersionCount"`
	// 创建该模板的用户UIN。 注意：此字段可能返回 null，表示取不到有效值。

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 创建该模板的时间。 注意：此字段可能返回 null，表示取不到有效值。

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 1

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportSnapshotTaskRequest struct {
	*tchttp.BaseRequest

	// 1

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeImportSnapshotTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportSnapshotTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserZoneStatusItem struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 售卖状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Placement struct {

	// 实例所属的[可用区](#zoneinfo)ID。该参数也可以通过调用  [DescribeZones](地域相关接口/DescribeZones) 的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [可用区](#zoneinfo) 的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	//实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。当前暂不支持。

	//PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 示例所属平台ProjectID

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 实例所属项目
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 实例所属项目

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type AuditMarketImageRequest struct {
	*tchttp.BaseRequest

	// 1

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 1

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *AuditMarketImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditMarketImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当通过本接口来创建实例启动模板时会返回该参数，表示创建成功的实例启动模板ID。

		LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLaunchTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnterRescueModeRequest struct {
	*tchttp.BaseRequest

	// 需要进入救援模式的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 救援模式下系统密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 救援模式下系统用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 是否强制关机

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 救援镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *EnterRescueModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnterRescueModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargeTypeConfig struct {

	// 实例计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例计费类型描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥对信息。

		KeyPair *CreateKeyPair `json:"KeyPair,omitempty" name:"KeyPair"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-xxxxxxxx`。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *ReturnAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTypeQuotaSet struct {

	// CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// CPU型号。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 磁盘大小，单位：GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 磁盘类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 专用宿主机机型系列。

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 专用宿主机类型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 内存大小，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 售价。

	Price *Price `json:"Price,omitempty" name:"Price"`
	// 售卖状态，“ SOLD_OUT”：售罄，“SELL”：售卖中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceTypeNameConfig struct {

	// 是否显示到实例类型列表

	ShowInMenu *bool `json:"ShowInMenu,omitempty" name:"ShowInMenu"`
	// 实例类型中文名

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 实例类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest

	// 镜像ID，形如`img-gvbnzy6f`。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](DescribeImages)接口返回的`ImageId`获取。<br><li>通过[DescribeImages](../镜像相关接口/DescribeImages)获取。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 设置新的镜像名称；必须满足下列限制：<br> <li> 不得超过20个字符。<br> <li> 镜像名称不能与已有镜像重复。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 设置新的镜像描述；必须满足下列限制：<br> <li> 不得超过60个字符。

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterAllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 用于保证请求幂等性的字符串。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// CDH实例机型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 购买CDH实例数量。

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否跳过实际执行逻辑

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 购买来源

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *SwitchParameterAllocateHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterAllocateHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID，用于查询任务状态、进度。

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 客户端版本。

	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
	// 迁移目的实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 迁移源实例系统信息。

	SysInfo *ClientSysInfo `json:"SysInfo,omitempty" name:"SysInfo"`
	// 是否忽略网络连通性检查

	IgnoreCheckNetworkConnectivity *bool `json:"IgnoreCheckNetworkConnectivity,omitempty" name:"IgnoreCheckNetworkConnectivity"`
	// 是否灰度版本。

	GrayscaleVersion *bool `json:"GrayscaleVersion,omitempty" name:"GrayscaleVersion"`
}

func (r *LiveMigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemPrice struct {

	// 后续单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后续计价单元，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：<br><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预支费用的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 折扣，如20.0代表2折

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
	// 后续合计费用的折扣价，后付费模式使用，单位：元<br><li>如返回了其他时间区间项，如UnitPriceDiscountSecondStep，则本项代表时间区间在(0, 96)小时；若未返回其他时间区间项，则本项代表全时段，即(0, ∞)小时

	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`
	// 使用时间区间在(96, 360)小时的后续合计费用的原价，后付费模式使用，单位：元。

	UnitPriceSecondStep *float64 `json:"UnitPriceSecondStep,omitempty" name:"UnitPriceSecondStep"`
	// 使用时间区间在(96, 360)小时的后续合计费用的折扣价，后付费模式使用，单位：元

	UnitPriceDiscountSecondStep *float64 `json:"UnitPriceDiscountSecondStep,omitempty" name:"UnitPriceDiscountSecondStep"`
	// 使用时间区间在(360, ∞)小时的后续合计费用的原价，后付费模式使用，单位：元。

	UnitPriceThirdStep *float64 `json:"UnitPriceThirdStep,omitempty" name:"UnitPriceThirdStep"`
	// 使用时间区间在(360, ∞)小时的后续合计费用的折扣价，后付费模式使用，单位：元

	UnitPriceDiscountThirdStep *float64 `json:"UnitPriceDiscountThirdStep,omitempty" name:"UnitPriceDiscountThirdStep"`
	// 详细价格

	DetailPrices *DetailPrices `json:"DetailPrices,omitempty" name:"DetailPrices"`
	// 使用时间区间在(360, ∞)小时的后续合计费用的折扣价，后付费模式使用，单位：元<br>注意：此字段可能返回 null，表示取不到有效值。

	DiscountThirdStep *float64 `json:"DiscountThirdStep,omitempty" name:"DiscountThirdStep"`
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSnapshotRequest struct {
	*tchttp.BaseRequest

	// 制作的快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 制作的快照描述

	SnapshotDescription *string `json:"SnapshotDescription,omitempty" name:"SnapshotDescription"`
	// 数据盘镜像COS链接

	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`
	// 制作的快照大小

	SnapshotSize *int64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// true为仅检查参数，默认为false

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ImportSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CbsMigrateTask struct {

	// 任务id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 云平台应用ID，一般来说与Uin存在一一对应的关系。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 快照url

	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`
	// 磁盘id

	DiskId *uint64 `json:"DiskId,omitempty" name:"DiskId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 数据大小

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 迁移任务的进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeKeyPairsAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 密钥对数组

		KeyPairAttributeSet *KeyPairAttributeSet `json:"KeyPairAttributeSet,omitempty" name:"KeyPairAttributeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecomendedZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按照实例可购买配额，从高到低排序的按量计费可用区列表。

		PostPaidZoneSet []*string `json:"PostPaidZoneSet,omitempty" name:"PostPaidZoneSet"`
		// 按照实例可购买配额，从高到低排序的包年包月可用区列表。

		PrePaidZoneSet []*string `json:"PrePaidZoneSet,omitempty" name:"PrePaidZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecomendedZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecomendedZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshInternalUserEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshInternalUserEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshInternalUserEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeNameConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypeNameConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeNameConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用宿主机机型配置信息列表

		HostTypeQuotaSet []*HostTypeConfigSet `json:"HostTypeQuotaSet,omitempty" name:"HostTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneHostConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneHostConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 私有网络相关信息配置。通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 是否对运行中的实例选择强制关机。默认为TRUE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *UpdateInstanceVpcConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceVpcConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可创建置放群组数量的上限。

		GroupQuota *int64 `json:"GroupQuota,omitempty" name:"GroupQuota"`
		// 当前用户已经创建的置放群组数量。

		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`
		// 物理机类型容灾组内实例的配额数。

		CvmInHostGroupQuota *int64 `json:"CvmInHostGroupQuota,omitempty" name:"CvmInHostGroupQuota"`
		// 交换机类型容灾组内实例的配额数。

		CvmInSwGroupQuota *int64 `json:"CvmInSwGroupQuota,omitempty" name:"CvmInSwGroupQuota"`
		// 机架类型容灾组内实例的配额数。

		CvmInRackGroupQuota *int64 `json:"CvmInRackGroupQuota,omitempty" name:"CvmInRackGroupQuota"`
		// 置放群组内实例的配额数。

		CvmInGroupQuota *int64 `json:"CvmInGroupQuota,omitempty" name:"CvmInGroupQuota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 定时器id列表

		ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>TRUE：表示在正常关闭失败后进行强制关闭<br><li>FALSE：表示在正常关闭失败后不进行强制关闭<br><br>默认取值：FALSE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭<br><li>HARD：直接强制关闭<br><li>SOFT：仅软关机<br>默认取值：SOFT。

	StoppedMode *string `json:"StoppedMode,omitempty" name:"StoppedMode"`
	// 实例的关闭模式。取值范围：SOFT_FIRST：表示在正常关闭失败后进行强制关闭HARD：直接强制关闭SOFT：仅软关机默认取值：SOFT。

	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *StopInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransformAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsRequest struct {
	*tchttp.BaseRequest

	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceOperationLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceOperationLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMarketImagesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的市场镜像的镜像Id

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 过滤条件。只支持按镜像Id过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移值，用于分页

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回镜像的数量。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 内部用户

	IsInner *string `json:"IsInner,omitempty" name:"IsInner"`
}

func (r *DescribeMarketImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMarketImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetAccessibleModifyChargeType struct {

	// 网络付费模式

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 外网出带宽值

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type KeyPairAttributeSet struct {

	// 密钥ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 内部密钥ID

	InnerKeyId *string `json:"InnerKeyId,omitempty" name:"InnerKeyId"`
}

type InquiryPriceTerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](../实例相关接口/DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 内部参数，释放弹性IP。释放eip传true要生效必须要先调用vpc的接口ModifyAddressesAttribute 设置eip的"CascadeRelease" 为true

	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
	// 试运行。对传入的数据只做一些参数校验，方便用户自测，默认是false

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceTerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceTerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResizeInstanceDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRunInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RequiredEnhancedService struct {

	// 监控服务

	MonitorService *MonitorServiceItem `json:"MonitorService,omitempty" name:"MonitorService"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表 。镜像ID如：`img-gvbnzy6f`。array型参数的格式可以参考[API简介](/document/api/213/11646)。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](DescribeImages)接口返回的`ImageId`获取。<br><li>通过[DescribeImages](../镜像相关接口/DescribeImages)获取。

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 过滤条件。其中Filters的上限为10，Filters.Values的上限为5。 注意：不可以同时指定ImageIds和Filters。可选值包括 `image-id`, `image-type`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于Offset详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 数量限制，默认为20，最大值为100。关于Limit详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 将该实例机型规格相关的镜像过滤掉，比如S1.SMALL1、S1.SMALL2等。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 宿主机Ip, 用于创建CDH子机时过滤镜像

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的导入镜像的操作系统类型

		ImportImageOsListSupported *ImportImageOsListSupported `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported"`
		// 支持导入镜像的操作系统信息

		ImportImageOsInfoSupported []*ImportImageOsVersion `json:"ImportImageOsInfoSupported,omitempty" name:"ImportImageOsInfoSupported"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCreateImageAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesCreateImageAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesCreateImageAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsAttributeRequest struct {
	*tchttp.BaseRequest

	// 密钥ID数组，密钥形式为skey-xxxxxxxx。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 偏移量，用于分页。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，用于分页，限制为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCbsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入数据盘的任务ID，用于查询任务状态、进度。

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCbsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCbsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 待扩容的系统盘配置信息。只支持扩容随实例购买的系统盘，且[系统盘类型](../数据结构#systemdisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。系统盘容量单位：GB。最小扩容步长：10G。关于系统盘类型的选择请参考硬盘产品简介。可选系统盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因系统盘类型的不同而有所差异。该接口调用时只传DiskSize参数就行。如:{DiskSize: xxx}

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 是否对运行中的实例选择强制关机，默认为False。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 是否在线升级配置

	Online *bool `json:"Online,omitempty" name:"Online"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceUsbInfoRequest struct {
	*tchttp.BaseRequest

	// 虚拟机uuid

	VmUuids []*string `json:"VmUuids,omitempty" name:"VmUuids"`
}

func (r *DescribeInstanceUsbInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 定时器数组

		ActionTimers []*ActionTimer `json:"ActionTimers,omitempty" name:"ActionTimers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例属性列表，标识实例id，实例订单Id。

		InstanceAttributeSet []*KeyBill `json:"InstanceAttributeSet,omitempty" name:"InstanceAttributeSet"`
		// 拉取实例的个数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ActionTimers

		ActionTimers []*ActionTimers `json:"ActionTimers,omitempty" name:"ActionTimers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {

	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。

	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`
	// 当前数量

	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`
	// 配额数量

	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分散置放群组信息列表

		DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet,omitempty" name:"DisasterRecoverGroupSet"`
		// 用户置放群组总量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDiagnosticReportsRequest struct {
	*tchttp.BaseRequest

	// 实例id数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否深度检测

	DeepDiagnose *bool `json:"DeepDiagnose,omitempty" name:"DeepDiagnose"`
	// 是否严格

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *CreateDiagnosticReportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiagnosticReportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageAttribute struct {

	// unImgId

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// deviceImageId

	InnerImageId *string `json:"InnerImageId,omitempty" name:"InnerImageId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
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

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// des或者allinone的taskid

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 内部参数，公网带宽相关信息设置。

	InternetAccessible []*InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 试运行。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 内部参数，续费弹性数据盘。

	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitempty" name:"RenewPortableDataDisk"`
}

func (r *InquiryPriceRenewInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示磁盘扩容成对应配置的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceTerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述退款详情。

		InstanceRefundsSet *InstanceRefundsSet `json:"InstanceRefundsSet,omitempty" name:"InstanceRefundsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceTerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceTerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyInstanceDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyInstanceDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的cdh实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// cdh实例详细信息列表

		HostSet []*HostItem `json:"HostSet,omitempty" name:"HostSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceForHostTypeQuota struct {

	// 不打折价格。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type Snapshot struct {

	// 快照Id。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 创建此快照的云硬盘类型。取值范围：
	// SYSTEM_DISK：系统盘
	// DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 创建此快照的云硬盘大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type SwitchParameterRenewHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cdh订单详细信息

		InstanceOrder *HostOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceCreateImageAttributeSet struct {

	// 云主机id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否支持实例镜像

	SupportInstanceImage *bool `json:"SupportInstanceImage,omitempty" name:"SupportInstanceImage"`
	// 是否支持在线创建镜像

	SupportOnlineCreateImage *bool `json:"SupportOnlineCreateImage,omitempty" name:"SupportOnlineCreateImage"`
	// 是否支持云初始化

	SupportCloudinit *bool `json:"SupportCloudinit,omitempty" name:"SupportCloudinit"`
	// 是否需要关机

	NeedPowerOff *bool `json:"NeedPowerOff,omitempty" name:"NeedPowerOff"`
}

type InstanceTypeQuota struct {

	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例机型配额。

	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
	// 实例计费模式。取值范围： <br>*`PREPAID`：表示预付费，即包年包月 <br>* `POSTPAID_BY_HOUR`：表示后付费，即按量计费 * `CDHPAID`：[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)付费，即只对[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)计费，不对[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例类型。

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 实例的CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例的GPU核数，单位：核。

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
}

type InstanceConfigInfoItem struct {

	// 实例规格。

	// 实例规格名称。

	// 优先级。

	// 实例族信息列表。

}

type ResourceForInstanceType struct {

	// 内存大小，单位GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// CPU核数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例机型规格，比如S1.SMALL1、S1.SMALL2等。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例规格状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ColdMigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 迁移的系统盘镜像COS链接

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 迁移目的实例ID。迁移完成后，该实例将运行迁入的操作系统。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用于测试参数合法性。默认为`false`

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 任务名称，用于在控制台展示。

	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *ColdMigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnNormalAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNormalAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionTimer struct {

	// 定时器

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// 执行时间

	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
	// 扩展数据

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
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

type DiagnosticReportDataSet struct {

	// 开始时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 报告id

	DiagnosticReportId *string `json:"DiagnosticReportId,omitempty" name:"DiagnosticReportId"`
	// 结束时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// ReportDetailSet

	ReportDetailSet *ReportDetailSet `json:"ReportDetailSet,omitempty" name:"ReportDetailSet"`
	// 检测状态

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 1

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的 EIP 数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// EIP 详细信息列表。

		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// instance-family
	//
	// 按照【实例机型系列】进行过滤。实例机型系列形如：S1、I1、M1等。
	//
	// 类型：String
	//
	// 必选：否
	// instance-type
	//
	// 按照【实例机型】进行过滤。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 DescribeInstanceTypeConfigs 来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。
	//
	// 类型：String
	//
	// 必选：否
	//
	// type
	//
	// 按照【实例族】进行过滤。实例族形如：S、I、M等。
	//
	// 类型：String
	//
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceVpcConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstanceVpcConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineMigrateUserTaskData struct {

	// 任务id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 云平台应用ID，一般来说与Uin存在一一对应的关系。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 实例uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 系统盘镜像cos url

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 辅助数据盘大小

	DataSize *uint64 `json:"DataSize,omitempty" name:"DataSize"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 迁移任务的进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 数据盘镜像url

	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`
	// 数据盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘大小

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例状态数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// [实例状态](../数据结构#instancestatus) 列表。

		InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAvailableZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区详情

		AvailableZoneSet *AvailableZoneSet `json:"AvailableZoneSet,omitempty" name:"AvailableZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAvailableZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAvailableZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroup struct {

	// 分散置放群组id。

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 分散置放群组名称，长度1-60个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分散置放群组类型，取值范围：<br><li>HOST：物理机<br><li>SW：交换机<br><li>RACK：机架

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分散置放群组内最大容纳云服务器数量。

	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`
	// 分散置放群组内云服务器当前数量。

	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`
	// 分散置放群组内，云服务器id列表。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 分散置放群组创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 置放群组策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

type HostOrder struct {

	// 订单所属帐号的应用id

	// 创建订单的帐号uin

	// 订单所属帐号的所有者uin

	// 订单发货的资源信息列表

}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户的镜像配额

		ImageNumQuota *int64 `json:"ImageNumQuota,omitempty" name:"ImageNumQuota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeItem struct {

	// 实例类型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU核数。

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// GPU核数。

	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`
	// FPGA核数。

	Fpga *uint64 `json:"Fpga,omitempty" name:"Fpga"`
	// 存储块数。

	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
	// 网卡数。

	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 最大带宽。

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 主频。

	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`
	// CPU型号名称。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 包转发率。

	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`
	// 外部信息。

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
	// 备注信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ImportInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 操作时间

	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
}

func (r *ImportInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRecentFailedOperationRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRecentFailedOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRecentFailedOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyBill struct {

	// 实例ID形如：ins-11112222。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 订单id

	BillId *string `json:"BillId,omitempty" name:"BillId"`
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请到的 EIP 的唯一 ID 列表。

		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCpuQuotaRequest struct {
	*tchttp.BaseRequest

	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneCpuQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneCpuQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// 用户全局配置列表，目前配置仅支持StoppedMode一个key值，表示用户默认关机不收费配置，KEEP_CHARGING为关机收费，STOP_CHARGING为关机不收费，默认为KEEP_CHARGING。

	Configs []*KvType `json:"Configs,omitempty" name:"Configs"`
}

func (r *ModifyUserGlobalConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGlobalConfigsRequest) FromJsonString(s string) error {
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

type CheckInstancesConnectivityRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *CheckInstancesConnectivityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckInstancesConnectivityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像ID

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按照实例可购买配额，从高到低排序的按量计费可用区列表。

		PostPaidZoneSet []*string `json:"PostPaidZoneSet,omitempty" name:"PostPaidZoneSet"`
		// 按照实例可购买配额，从高到低排序的包年包月可用区列表。

		PrePaidZoneSet []*string `json:"PrePaidZoneSet,omitempty" name:"PrePaidZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecommendedZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecommendedZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneCpuQuota struct {

	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例计费模式。PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)付费，即只对[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)计费，不对[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 可用CPU配额。

}

type HostForSellZoneStatus struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可售卖情况, (SELL: 正常售卖，SOLD_OUT: 售罄)。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type HostGoodsDetailItem struct {

	// 请求事务id

	// 操作名称

	// 自动续费标记

	// pid

	// 购买或续费时长

	// 时间单位

	// 资源id

	// 当前到期时间

	// 数字签名

	// 产品信息项列表

}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkSharingGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNetworkSharingGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkSharingGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// 1

	IspId *int64 `json:"IspId,omitempty" name:"IspId"`
	// 1

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 1

	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`
	// 1

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 1

	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`
	// 1

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 1

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *InquiryPriceAllocateAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当通过本接口来创建实例启动模板时会返回该参数，表示创建成功的实例启动模板ID。

		LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLaunchTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostItem struct {

	// cdh实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// cdh实例id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// cdh实例类型

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// cdh实例名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// cdh实例付费模式

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// cdh实例自动续费标记

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// cdh实例创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// cdh实例过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// cdh实例上已创建云子机的实例id列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// cdh实例状态

	HostState *string `json:"HostState,omitempty" name:"HostState"`
	// cdh实例ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// cdh实例资源信息

	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`
	// cage id

	CageId *string `json:"CageId,omitempty" name:"CageId"`
	// 操作掩码

	OperationMask *int64 `json:"OperationMask,omitempty" name:"OperationMask"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type DeleteLaunchTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLaunchTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCreateImageAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例创建镜像列表

		InstanceCreateImageAttributeSet []*InstanceCreateImageAttributeSet `json:"InstanceCreateImageAttributeSet,omitempty" name:"InstanceCreateImageAttributeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesCreateImageAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesCreateImageAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchUserInstanceRequest struct {
	*tchttp.BaseRequest

	// 模糊搜索关键字。包括实例的Id，实例的名称，实例的公网和内网ip 关键字。

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 起始数，用于分页。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量，用于分页，默认20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchUserInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchUserInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchTemplatesRequest struct {
	*tchttp.BaseRequest

	// 启动模板ID。

	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitempty" name:"LaunchTemplateIds"`
}

func (r *DeleteLaunchTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLaunchTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户 EIP 配额信息。

		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttributeSet struct {

	// 属性类型

	AttributeClass *string `json:"AttributeClass,omitempty" name:"AttributeClass"`
	// 属性值

	AttributeValueSet []*AccountAttribute `json:"AttributeValueSet,omitempty" name:"AttributeValueSet"`
}

type TagSpecification struct {

	// 标签绑定的资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CopyInstanceDiskRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	SourceDiskId *string `json:"SourceDiskId,omitempty" name:"SourceDiskId"`
	// 1

	DestinationDiskId *string `json:"DestinationDiskId,omitempty" name:"DestinationDiskId"`
}

func (r *CopyInstanceDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyInstanceDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API[简介](/document/api/213/11646)的`id.N`一节）。每次请求的实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示对应配置实例的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功的列表

		SuccActionTimerIds []*string `json:"SuccActionTimerIds,omitempty" name:"SuccActionTimerIds"`
		// 失败的列表

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageBlock struct {

	// HDD本地存储类型，值为：LOCAL_PRO. <br>注意：此字段可能返回 null，表示取不到有效值。

	Type *string `json:"Type,omitempty" name:"Type"`
	// HDD本地存储的最小容量 <br>注意：此字段可能返回 null，表示取不到有效值。

	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`
	// HDD本地存储的最大容量 <br>注意：此字段可能返回 null，表示取不到有效值。

	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type DescribeImageSnapshotStatusRequest struct {
	*tchttp.BaseRequest

	// 需要查看的镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSnapshotStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSnapshotStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTypeQuota struct {

	// CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// CPU型号。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 磁盘大小，单位：GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 磁盘类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 专用宿主机机型系列。

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 专用宿主机类型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 内存大小，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 售价。

	Price *PriceForHostTypeQuota `json:"Price,omitempty" name:"Price"`
	// 售卖状态，“ SOLD_OUT”：售罄，“SELL”：售卖中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceConfigInfoItemArchitecture struct {

	// 实例规格。

	// 实例规格名称。

	// 优先级。

	// 实例族信息列表。

}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 试运行。默认False

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例机型配置列表。

		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥对ID。

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientSysInfo struct {

	// 操作系统

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 操作系统版本

	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`
	// 需要导入的系统盘数据盘信息

	DiskSize []*uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 额外信息

	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

type ReportDetailSet struct {

	// 检测类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 检测类型名

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
	// 检测编码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 编码名称

	CodeName *string `json:"CodeName,omitempty" name:"CodeName"`
	// 已授权

	IsAuthorized *bool `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
	// 信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 检测状态

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：ins-11112222。每次请求的实例的上限为100

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecomendedZonesRequest struct {
	*tchttp.BaseRequest

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](DescribeInstanceTypeConfigs)来获得最新的规格表或参见[实例类型](/tcloud/Compute/CVM/292128/484318/specification)描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribeRecomendedZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecomendedZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 1

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 1

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例操作。INSTANCE_DEGRADE：实例降配操作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeInstancesOperationLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesOperationLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateRequest struct {
	*tchttp.BaseRequest

	// 需要迁移的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// lossy参数

	LossyLocal *bool `json:"LossyLocal,omitempty" name:"LossyLocal"`
	// 可指定母机迁移

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
}

func (r *FailoverMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailoverMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateRequest struct {
	*tchttp.BaseRequest

	// 待迁移实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指定母机迁移

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要迁入的售卖池， PLAIN, OVERSOLD等

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 跨可用区迁移，要传入可用区表

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 最大迁移带宽

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 最大迁移超时

	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
}

func (r *LiveMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesAttributeRequest struct {
	*tchttp.BaseRequest

	// unImgId数组。

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

func (r *DescribeImagesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttachedDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设备id

		DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
		// 供应商id

		VendorId *string `json:"VendorId,omitempty" name:"VendorId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAttachedDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAttachedDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceChargeTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceChargeTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启[云监控](/tcloud/omtool/BARAD)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type SharePermission struct {

	// 镜像分享时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 镜像分享的账户ID

	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组id列表。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 分散置放群组名称，支持模糊匹配。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 置放群组策略列表

	Strategies []*string `json:"Strategies,omitempty" name:"Strategies"`
}

func (r *DescribeDisasterRecoverGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkSharingGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkSharingGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkSharingGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetBandwidthConfig struct {

	// 开始时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例带宽信息。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
}

type InstanceStatisticsSet struct {

	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 机器总量

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 红点主机总数（刚新建机器）

	NewInstanceCount *uint64 `json:"NewInstanceCount,omitempty" name:"NewInstanceCount"`
	// 过期主机总数

	ExpiredInstanceCount *uint64 `json:"ExpiredInstanceCount,omitempty" name:"ExpiredInstanceCount"`
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组名称，长度1-60个字符，支持中、英文。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分散置放群组类型，取值范围：<br><li>HOST：物理机<br><li>SW：交换机<br><li>RACK：机架

	Type *string `json:"Type,omitempty" name:"Type"`
	// 置放群组类型：分散置放群组；分区置放群组；强亲和置放群组

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 密钥对ID，密钥对ID形如：`skey-11112222`（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API [简介](/document/api/213/11646)的 `id.N` 一节）。参数不支持同时指定 `KeyIds` 和 `Filters`。<br> 密钥对ID可以通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 过滤条件，详见密钥对过滤条件表。参数不支持同时指定 `KeyIds` 和 `Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](/doc/api/229/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataDisk struct {

	// 数据盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘类型：LOCAL_BASIC、CLOUD_BASIC、LOCAL_SSD、CLOUD_SSD、CLOUD_PREMIUM、CLOUD_ENHANCEDSSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

type QueryPlacement struct {

	// 可用区名称

	Zone *int64 `json:"Zone,omitempty" name:"Zone"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// Vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](DescribeImages)接口返回的`ImageId`获取。<br><li>通过[DescribeImages](../镜像相关接口/DescribeImages)获取。<br>镜像ID必须满足限制：<br><li>镜像ID对应的镜像状态必须为`NORMAL`。<br><li>镜像大小小于50GB。<br>镜像状态请参考[镜像数据表](../数据结构#image)。

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 目的同步地域列表；必须满足限制：<br><li>不能为源地域，<br><li>必须是一个合法的Region。<br><li>暂不支持部分地域同步。<br>具体地域参数请参考[Region](/tcloud/Compute/CVM/292128/zone)。

	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`
}

func (r *SyncImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {

	// 付费类型：后付费（PostPaidInstanceCountLimit）,预付费（PrePaidInstanceCountLimit）

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 单次最大购买数量

	AttributeValues *string `json:"AttributeValues,omitempty" name:"AttributeValues"`
}

type DiagnosticReportSet struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 检测报告id

	DiagnosticReportId *string `json:"DiagnosticReportId,omitempty" name:"DiagnosticReportId"`
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// zone按照【可用区】进行过滤。可用区形如：ap-guangzhou-1。类型：String必选：是可选项：可用区列表instance-family按照【实例机型系列】进行过滤。实例机型系列形如：S1、I1、M1等。类型：Integer必选：否instance-type按照【实例机型】进行过滤。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 DescribeInstanceTypeConfigs 来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。类型：String必选：否instance-charge-type按照【实例计费模式】进行过滤。(PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。)类型：String必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CDH实例创建价格信息

		Price *HostPrice `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 1

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组ID列表，可通过[DescribeDisasterRecoverGroups](DescribeDisasterRecoverGroups)接口获取。每次请求允许操作的分散置放群组数量上限是100。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *DeleteDisasterRecoverGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeQuotaItem struct {

	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例计费模式。取值范围： <br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费<br><li>CDHPAID：表示[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)付费，即只对CDH计费，不对CDH上的实例计费。<br><li>`SPOTPAID`：表示竞价实例付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 网卡类型，例如：25代表25G网卡

	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 扩展属性。

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
	// 实例的CPU核数，单位：核。

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例机型系列。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 机型名称。

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 本地磁盘规格列表。当该参数返回为空值时，表示当前情况下无法创建本地盘。

	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`
	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br><li>SOLD_OUT：表示实例已售罄。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例的售卖价格。

	Price *ItemPrice `json:"Price,omitempty" name:"Price"`
	// 售罄原因。

	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`
	// 内网带宽，单位Gbps。

	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`
	// 网络收发包能力，单位万PPS。

	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`
	// 本地存储块数量。

	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`
	// 处理器型号。

	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`
	// 实例的GPU数量。

	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`
	// 实例的FPGA数量。

	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`
	// 实例备注信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeInstanceUsbInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceUsbInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了单台实例操作次数限制

		InstanceOperationLimitSet []*InstanceOperationLimitSet `json:"InstanceOperationLimitSet,omitempty" name:"InstanceOperationLimitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesOperationLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesOperationLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的 `InstanceId` 获取。 每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持 `InternetMaxBandwidthOut` 参数。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 带宽生效的终止时间。格式： `YYYY-MM-DD` ，例如：`2016-10-30` 。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [DescribeInstances](DescribeInstances)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KvType struct {

	// 键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 键所对应的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区机型配置列表。

		InstanceTypeQuotaSet []*InstanceTypeQuotaItemArchitecture `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type ImportImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像id

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 用于保证请求幂等性的字符串

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）, POSTPAID_BY_HOUR(后付费：按小时后付费)。不传该参数默认后付费

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// CDH实例机型。不传该参数 默认HS20机型

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 购买CDH实例数量

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否跳过实际执行逻辑

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 购买来源.API 和MC(mc表示前端调用)

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *AllocateHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInternetBandwidthConfigsRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {

	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](//console.{{conf.main_domain}}/vpc/vpc?rid=1)查询。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](//console.{{conf.main_domain}}/vpc/subnet?rid=1)查询；也可以调用接口  [控制台](//console.{{conf.main_domain}}/vpc/a需更新描述，TCE 无此“DescribeSubnetEx”接口?rid=1) ，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。

	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 为弹性网卡指定随机生成的 IPv6 地址数量。

	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
	// vpc名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 子网名

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest

	// 用户账号名称

	AttributeName []*string `json:"AttributeName,omitempty" name:"AttributeName"`
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGlobalConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户全局配置列表。

		ConfigSet []*KvType `json:"ConfigSet,omitempty" name:"ConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGlobalConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGlobalConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportFullCvmImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID，可利用该ID查询任务状态。

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportFullCvmImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportFullCvmImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示对应配置实例的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGlobalConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserGlobalConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGlobalConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstancesRecentFailedOperationSet struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 操作事件发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 1

	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
}

func (r *DeleteInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyInstancesChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterModifyInstancesChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 1

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 1

	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitempty" name:"RenewPortableDataDisk"`
}

func (r *SwitchParameterRenewInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例信息

		InstanceStatisticsSet []*InstanceStatisticsSet `json:"InstanceStatisticsSet,omitempty" name:"InstanceStatisticsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示重装成对应配置实例的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSecurityServiceEnabled struct {

	// 是否开启[云安全](/tcloud/Security/CWP)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type FailoverMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailoverMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailoverMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDiagnosticReportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例检测报告

		DiagnosticReportSet []*DiagnosticReportSet `json:"DiagnosticReportSet,omitempty" name:"DiagnosticReportSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDiagnosticReportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiagnosticReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostGoodsItem struct {

	// goodsCategoryId

	// 实例付费模式

	// 发货的实例个数

	// 地域id

	// 发起发货的用户uin

	// 发起发货帐号的所有者uin

	// 发起发货帐号对应的appId

	// 项目id

	// 可用区id

	// cdh实例详细信息

}

type InstanceTypeQuotaItemArchitecture struct {

	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例计费模式。取值范围： <br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费<br><li>CDHPAID：表示[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)付费，即只对CDH计费，不对CDH上的实例计费。<br><li>`SPOTPAID`：表示竞价实例付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 网卡类型，例如：25代表25G网卡

	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 扩展属性。

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
	// 实例的CPU核数，单位：核。

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例机型系列。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 机型名称。

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 本地磁盘规格列表。当该参数返回为空值时，表示当前情况下无法创建本地盘。

	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`
	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br><li>SOLD_OUT：表示实例已售罄。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例的售卖价格。

	Price *ItemPrice `json:"Price,omitempty" name:"Price"`
	// 售罄原因。

	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`
	// 内网带宽，单位Gbps。

	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`
	// 网络收发包能力，单位万PPS。

	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`
	// 本地存储块数量。

	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`
	// 处理器型号。

	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`
	// 实例的GPU数量。

	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`
	// 实例的FPGA数量。

	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`
	// 实例备注信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 配置Id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 设备类型

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 禁用

	Disable *string `json:"Disable,omitempty" name:"Disable"`
	// 本地存储块

	StorageBlock *int64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
}

type DeleteDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPair struct {

	// 密钥对的`ID`，是密钥对的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥对名称。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 密钥对所属的项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 密钥对的纯文本公钥。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// 密钥对的纯文本私钥。

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type ModifyDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组id，可使用DescribeDisasterRecoverGroups接口查询。

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 分散置放群组名称，最长128个字符。

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableZoneSet struct {

	// 可用区id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例类型

	InstanceType []*int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type Image struct {

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 镜像类型

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 镜像创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像描述

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 镜像架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 镜像状态

	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`
	// 镜像来源平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 镜像创建者

	ImageCreator *string `json:"ImageCreator,omitempty" name:"ImageCreator"`
	// 镜像来源

	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`
	// 同步百分比

	SyncPercent *int64 `json:"SyncPercent,omitempty" name:"SyncPercent"`
	// 镜像是否支持cloud-init

	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitempty" name:"IsSupportCloudinit"`
	// 镜像关联的快照信息

	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
	// 镜像是否支持自动安装GPU驱动

	IsSupportAutoInstallGPUDriver *bool `json:"IsSupportAutoInstallGPUDriver,omitempty" name:"IsSupportAutoInstallGPUDriver"`
	// 镜像名称

	OsKey *string `json:"OsKey,omitempty" name:"OsKey"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// GPU驱动列表

	GPUDriverList []*string `json:"GPUDriverList,omitempty" name:"GPUDriverList"`
	// 镜像格式

	ImageFormat *string `json:"ImageFormat,omitempty" name:"ImageFormat"`
	// 镜像操作掩码

	OperationMask *int64 `json:"OperationMask,omitempty" name:"OperationMask"`
	// 镜像标记

	Flags *string `json:"Flags,omitempty" name:"Flags"`
	// 自定义镜像id

	DeviceImageId *int64 `json:"DeviceImageId,omitempty" name:"DeviceImageId"`
	// TAT支持镜像情况

	IsSupportTat *bool `json:"IsSupportTat,omitempty" name:"IsSupportTat"`
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例机型组配置的列表信息

		InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesModificationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 跨机型调整配置返回的机型列表

		InstanceTypeConfigStatusSet []*InstanceTypeConfigStatusSet `json:"InstanceTypeConfigStatusSet,omitempty" name:"InstanceTypeConfigStatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesModificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesModificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeConfig struct {

	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例机型系列。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// GPU核数，单位：核。

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
	// CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// FPGA核数，单位：核。

	FPGA *uint64 `json:"FPGA,omitempty" name:"FPGA"`
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 1

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 1

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 1

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCpuQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区 CPU 配额信息。

		ZoneCpuQuotaSet []*ZoneCpuQuota `json:"ZoneCpuQuotaSet,omitempty" name:"ZoneCpuQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneCpuQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneCpuQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateHostsRequest struct {
	*tchttp.BaseRequest

	// 用于保证请求幂等性的字符串。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 实例计费类型。目前支持´：PREPAID（预付费，即包年包月模式），POSTPAID_BY_HOUR：按小时后付费。不传该参数默认后付费。

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// CDH实例机型，不传该参数 默认HS20机型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 购买CDH实例数量。

	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
	// 是否跳过实际执行逻辑

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 购买来源,API 和MC(mc表示前端调用)

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *InquiryPriceAllocateHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterAllocateHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cdh订单详细信息

		InstanceOrder *HostOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterAllocateHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterAllocateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 镜像ID，形如`img-gvbnzy6f`。镜像Id可以通过如下方式获取：<br><li>通过[DescribeImages](DescribeImages)接口返回的`ImageId`获取。<br><li>通过[DescribeImages](../镜像相关接口/DescribeImages)获取。 <br>镜像ID必须指定为状态为`NORMAL`的镜像。镜像状态请参考[镜像数据表](../数据结构#image)。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 接收分享镜像的账号Id列表，array型参数的格式可以参考[API简介](/document/api/213/568)。帐号ID不同于QQ号，查询用户帐号ID请查看[帐号信息](//console.{{conf.main_domain}}/developer)中的帐号ID栏。

	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
	// 操作，包括 `SHARE`，`CANCEL`。其中`SHARE`代表分享操作，`CANCEL`代表取消分享操作。

	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitRescueModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExitRescueModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitRescueModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchTemplateVersionInfo struct {

	// 实例启动模板版本号。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" name:"LaunchTemplateVersion"`
	// 实例启动模板版本数据详情。

	LaunchTemplateVersionData *LaunchTemplateVersionData `json:"LaunchTemplateVersionData,omitempty" name:"LaunchTemplateVersionData"`
	// 实例启动模板版本创建时间。

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 实例启动模板ID。

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
	// 是否为默认启动模板版本。

	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`
	// 实例启动模板版本描述信息。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`
	// 创建者。

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
}

type DeleteLaunchTemplateRequest struct {
	*tchttp.BaseRequest

	// 启动模板ID。

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
}

func (r *DeleteLaunchTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLaunchTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeNameConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例类型名称列表。

		InstanceTypeNameConfigSet []*InstanceTypeNameConfig `json:"InstanceTypeNameConfigSet,omitempty" name:"InstanceTypeNameConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeNameConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeNameConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTypeConfigSet struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 机型注意：此字段可能返回 null，表示取不到有效值。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// cdh实例付费模式

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// cdh类型

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 实例的CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 创建此快照的云硬盘大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// cdh配额

	HostQuota *uint64 `json:"HostQuota,omitempty" name:"HostQuota"`
	// 价格信息

	Price *HostTypeConfigSetPrice `json:"Price,omitempty" name:"Price"`
	// CPU型号名称。

	CpuModeLName *string `json:"CpuModeLName,omitempty" name:"CpuModeLName"`
	// 实例是否售卖。取值范围：SELL：表示实例可购买SOLD_OUT：表示实例已售罄。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费类型配置

		InstanceChargeTypeConfigSet []*InstanceChargeTypeConfig `json:"InstanceChargeTypeConfigSet,omitempty" name:"InstanceChargeTypeConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceChargeTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceChargeTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 项目ID。后续使用[DescribeInstances](DescribeInstances)接口查询实例时，项目ID可用于过滤结果。

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesVpcAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesVpcAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)&nbsp;API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例显示名称。可任意命名，但不得超过60个字符。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内部参数，用户数据。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// 内部参数，安全组Id列表。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
	// 内部参数，未知。

	ResetNewCreationIdentify *bool `json:"ResetNewCreationIdentify,omitempty" name:"ResetNewCreationIdentify"`
	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>&nbsp;TRUE：表示开启实例保护，不允许通过api接口删除实例<br><li>&nbsp;FALSE：表示关闭实例保护，允许通过api接口删除实例&nbsp;<br>&nbsp;默认取值：FALSE。

	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。

	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。

	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type InstanceReturnable struct {

	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例是否可退还。

	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`
	// 实例退还失败错误码。

	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`
	// 实例退还失败错误信息。

	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type EnterRescueModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnterRescueModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnterRescueModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecycleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRecycleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecycleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecycleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收站设置时间

		RecycleTime *string `json:"RecycleTime,omitempty" name:"RecycleTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecycleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecycleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建cdh的实例id列表。

		HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourcesOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyPairInstancesinternetaccessible struct {

	// 密钥对的`ID`，是密钥对的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥对关联的实例`ID`列表。

	AssociatedInstanceIdSet []*string `json:"AssociatedInstanceIdSet,omitempty" name:"AssociatedInstanceIdSet"`
}

type DeleteDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 容灾组id列表，可通过DescribeDisasterRecoverGroups接口查询。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *DeleteDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchTemplateNameRequest struct {
	*tchttp.BaseRequest

	// 启动模板ID。

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
	// 启动模板名。

	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`
}

func (r *ModifyLaunchTemplateNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLaunchTemplateNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransformAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstanceQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用于标识用户在可用区下包年包月、按量计费配额。

		UserInstanceQuotaSet []*KeyQuota `json:"UserInstanceQuotaSet,omitempty" name:"UserInstanceQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInstanceQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstanceQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// 系统盘类型。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_PREMIUM：高性能云盘<br><br>默认取值：LOCAL_BASIC。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC&nbsp;和&nbsp;LOCAL_SSD&nbsp;类型没有ID。暂时不支持该参数。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 系统盘大小，单位：GB。默认值为&nbsp;50

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘指定的存储池。

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云盘的自动备份策略id

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 是否加密

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 加密密钥ID

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
	// 吞吐性能（MB/s）

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 独享集群ID

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的密钥对数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 密钥对详细信息列表。

		KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCdhInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneCdhInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneCdhInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostForSellStatusRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneHostForSellStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneHostForSellStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportImageOsVersion struct {

	// CPU架构

	Architecture []*string `json:"Architecture,omitempty" name:"Architecture"`
	// OS名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// OS版本号

	OsVersions []*string `json:"OsVersions,omitempty" name:"OsVersions"`
}

type ExportImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// COS/CSP Bucket的路径

		CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryResourceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例类型的资源

		InquiryResourceResetInstancesType []*ResourceForInstanceType `json:"InquiryResourceResetInstancesType,omitempty" name:"InquiryResourceResetInstancesType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryResourceResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryResourceResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例计费类型。<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费），该付费模式下必须填写placement.hostid参数<br>默认值：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](DescribeInstanceTypeConfigs)来获得最新的规格表或参见[实例类型](/tcloud/Compute/CVM/292128/484318/specification)描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 指定有效的[镜像](/tcloud/Compute/CVM/292128/835305/mirr_overview)ID，格式形如`img-xxx`。镜像类型分为三种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](//console.{{conf.main_domain}}/cvm/image/list?imageType=PUBLIC_IMAGE&pageIndex=1&pageSize=20)查询；</li><li>通过调用接口&nbsp;[DescribeImages](../镜像相关接口/DescribeImages)&nbsp;，取返回信息中的`ImageId`字段。</li>

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 购买源

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，支持购买时指定多个数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，默认使用vpc网络。若在此参数中指定了私有网络ip，那么InstanceCount参数可以填1或2。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。如果不指定则默认显示.&nbsp;最多只支持60个字符，点后面的名字都会过滤掉。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。若不指定该参数，则绑定默认安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，默认关闭云监控和云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 用于指定价格生产，当前主要用于竞价实例

	SpotPrice *string `json:"SpotPrice,omitempty" name:"SpotPrice"`
	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为&nbsp;HostName&nbsp;的首尾字符，不能连续使用。<br><li>Windows&nbsp;实例：名字符长度为[2,&nbsp;15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux&nbsp;等）实例：字符长度为[2,&nbsp;31]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成，不支持全数字;不支持.-(点和短横线放在一起)。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 提供给实例使用的用户数据，需要以&nbsp;base64&nbsp;方式编码，支持的最大数据大小为&nbsp;16KB。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 指定的项目id，仅能指定一个

	ProjectSpecification *ProjectSpecification `json:"ProjectSpecification,omitempty" name:"ProjectSpecification"`
	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。

	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`
	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>&nbsp;TRUE：表示开启实例保护，不允许通过api接口删除实例<br><li>&nbsp;FALSE：表示关闭实例保护，允许通过api接口删除实例&nbsp;<br>&nbsp;默认取值：FALSE。

	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

func (r *RunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest

	// 需要制作镜像的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像描述

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 软关机失败时是否执行强制关机以制作镜像

	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`
	// 创建Windows镜像时是否启用Sysprep

	Sysprep *string `json:"Sysprep,omitempty" name:"Sysprep"`
	// 实例处于运行中时，是否允许关机执行制作镜像任务。

	Reboot *string `json:"Reboot,omitempty" name:"Reboot"`
	// 基于实例创建整机镜像时，指定包含在镜像里的数据盘Id

	DataDiskIds []*string `json:"DataDiskIds,omitempty" name:"DataDiskIds"`
	// 基于快照创建镜像，指定快照ID，必须包含一个系统盘快照。不可与InstanceId同时传入。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 检测本次请求的是否成功，但不会对操作的资源产生任何影响

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 是否执行软关机以制作镜像。

	SoftPoweroff []*string `json:"SoftPoweroff,omitempty" name:"SoftPoweroff"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressBandwidthConfigsRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAddressBandwidthConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressBandwidthConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID数组。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 私有网络相关信息配置，通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。<br><li>当指定私有网络ID和子网ID（子网必须在实例所在的可用区）与指定实例所在私有网络不一致时，会将实例迁移至指定的私有网络的子网下。<br><li>可通过`PrivateIpAddresses`指定私有网络子网IP，若需指定则所有已指定的实例均需要指定子网IP，此时`InstanceIds`与`PrivateIpAddresses`一一对应。<br><li>不指定`PrivateIpAddresses`时随机分配私有网络子网IP。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 是否对运行中的实例选择强制关机。默认为TRUE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 是否保留主机名。默认为FALSE。

	ReserveHostName *bool `json:"ReserveHostName,omitempty" name:"ReserveHostName"`
}

func (r *ModifyInstancesVpcAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesVpcAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelAuditMarketImageRequest struct {
	*tchttp.BaseRequest

	// 1

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 1

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *CancelAuditMarketImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelAuditMarketImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchTemplatesInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例模板数量。 注意：此字段可能返回 null，表示取不到有效值。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例详细信息列表。 注意：此字段可能返回 null，表示取不到有效值。

		LaunchTemplateInfoSet []*LaunchTemplatesInfo `json:"LaunchTemplateInfoSet,omitempty" name:"LaunchTemplateInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLaunchTemplatesInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLaunchTemplatesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportImageRequest struct {
	*tchttp.BaseRequest

	// 导入镜像的操作系统架构，`x86_64` 或 `arm_64`

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 导入镜像的操作系统类型，通过`DescribeImportImageOs`获取

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 导入镜像的操作系统版本，通过`DescribeImportImageOs`获取

	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`
	// 导入镜像存放的cos地址

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像名称。镜像名称不能与已有的自定义镜像名称重复,长度不超过60

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像描述

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 只检查参数，不执行任务

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 是否强制导入，参考[强制导入镜像](/tcloud/Compute/CVM/221286/642931/948877/forceimgimport)

	Force *bool `json:"Force,omitempty" name:"Force"`
	// 是否导入公共镜像，导入公共镜像为`true`，走imagestage流程，默认为`false`；

	PublicImage *bool `json:"PublicImage,omitempty" name:"PublicImage"`
}

func (r *ImportImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserMigrateTaskData struct {

	// 任务id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 云平台应用ID，一般来说与Uin存在一一对应的关系。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 实例uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据盘镜像cos url

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 数据大小

	DataSize *uint64 `json:"DataSize,omitempty" name:"DataSize"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 迁移任务的进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽配置信息列表。

		InternetBandwidthConfigSet []*InternetBandwidthConfig `json:"InternetBandwidthConfigSet,omitempty" name:"InternetBandwidthConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API[简介](/document/api/213/11646)的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。zone按照【可用区】进行过滤。可用区形如：ap-region1-1。类型：String必选：否可选项：可用区列表project-id按照【项目ID】进行过滤，可通过调用DescribeProject查询已创建的项目列表或登录控制台进行查看；也可以调用AddProject创建新的项目。项目ID形如：1002189。类型：Integer必选：否host-id按照【CDH&nbsp;ID】进行过滤。CDH&nbsp;ID形如：host-xxxxxxxx。类型：String必选：否vpc-id按照【VPC&nbsp;ID】进行过滤。VPC&nbsp;ID形如：vpc-xxxxxxxx。类型：String必选：否subnet-id按照【子网ID】进行过滤。子网ID形如：subnet-xxxxxxxx。类型：String必选：否instance-id按照【实例ID】进行过滤。实例ID形如：ins-61dyt49hxxxx。类型：String必选：否security-group-id按照【安全组ID】进行过滤。安全组ID形如:&nbsp;sg-8jlk3f3r。类型：String必选：否instance-name按照【实例名称】进行过滤。类型：String必选：否instance-charge-type按照【实例计费模式】进行过滤。(POSTPAID_BY_HOUR：表示后付费，即按量计费&nbsp;|&nbsp;CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。)类型：String必选：否instance-state按照【实例状态】进行过滤。PENDING：表示创建中LAUNCH_FAILED：表示创建失败RUNNING：表示运行中STOPPED：表示关机STARTING：表示开机中STOPPING：表示关机中REBOOTING：表示重启中SHUTDOWN：表示停止待销毁TERMINATING：表示销毁中。类型：String必选：否private-ip-address按照【实例主网卡的内网IP】进行过滤。类型：String必选：否public-ip-address按照【实例主网卡的公网IP】进行过滤，包含实例创建时自动分配的IP和实例创建后手动绑定的弹性IP。类型：String必选：否ipv6-address按照【实例的IPv6地址】进行过滤。类型：String必选：否tag-key按照【标签键】进行过滤。类型：String必选：否tag-value按照【标签值】进行过滤。类型：String必选：否tag:tag-key按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。类型：String必选：否。architectures&nbsp;&nbsp;按照【架构】进行过滤。&nbsp;&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;instance-family&nbsp;&nbsp;按照【主机类型】进行过滤。&nbsp;&nbsp;类型：String&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考&nbsp;API&nbsp;[简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考&nbsp;API&nbsp;[简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 内部参数，数字型vpcId列表&nbsp;。

	InnerVpcIds []*int64 `json:"InnerVpcIds,omitempty" name:"InnerVpcIds"`
	// 内部参数，数字型subnetId列表。

	InnerSubnetIds []*int64 `json:"InnerSubnetIds,omitempty" name:"InnerSubnetIds"`
	// 内部参数，精确内外网IP列表。

	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses"`
	// 内部参数，模糊内外网IP。

	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`
	// 内部参数，模糊实例别名。

	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitLiveMigrateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExitLiveMigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitLiveMigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchTemplateVersionData struct {

	// 实例所在的位置。 注意：此字段可能返回 null，表示取不到有效值。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例机型。 注意：此字段可能返回 null，表示取不到有效值。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例名称。 注意：此字段可能返回 null，表示取不到有效值。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例付费类型。取值范围：<br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费<br><li>CDHPAID：专用宿主机付费，即只对专用宿主机计费，不对专用宿主机上的实例计费。<br><li>SPOTPAID：表示竞价实例付费。<br>注意：此字段可能返回 null，表示取不到有效值。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例系统盘信息。 注意：此字段可能返回 null，表示取不到有效值。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘信息。只包含随实例购买的数据盘。 注意：此字段可能返回 null，表示取不到有效值。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 实例带宽信息。<br/>注意：此字段可能返回 null，表示取不到有效值。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 实例所属虚拟私有网络信息。<br/>注意：此字段可能返回 null，表示取不到有效值。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 生产实例所使用的镜像ID。<br/>注意：此字段可能返回 null，表示取不到有效值。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。 <br/>注意：此字段可能返回 null，表示取不到有效值。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例登录设置。目前只返回实例所关联的密钥。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// CAM角色名。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
	// 高性能计算集群ID。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`
	// 购买实例数量。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 增强服务。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// 置放群组id，仅支持指定一个。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
	// 云服务器的主机名。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 用于保证请求幂等性的字符串。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的云服务器、云硬盘实例。<br/> 注意：此字段可能返回 null，表示取不到有效值。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type SwitchParameterRunInstancesRequest struct {
	*tchttp.BaseRequest

	// 1

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 1

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 1

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 1

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 1

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 1

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 1

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 1

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 1

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 1

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 1

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 1

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 1

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 1

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 1

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 1

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// 1

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 1

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
	// 1

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 1

	DesAction *string `json:"DesAction,omitempty" name:"DesAction"`
	// 1

	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
}

func (r *SwitchParameterRunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRunInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageAttributeSet struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 内部镜像id

	InnerImageId *uint64 `json:"InnerImageId,omitempty" name:"InnerImageId"`
}

type InstanceDeniedActions struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作限制列表

	DeniedActions []*DeniedActions `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type InstanceOperationLimitSet struct {

	// 实例操作。取值范围：INSTANCE_DEGRADE：降配操作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 当前已使用次数，如果返回值为-1表示该操作无次数限制。

	CurrentCount *uint64 `json:"CurrentCount,omitempty" name:"CurrentCount"`
	// 操作次数最高额度，如果返回值为-1表示该操作无次数限制，如果返回值为0表示不支持调整配置。

	LimitCount *uint64 `json:"LimitCount,omitempty" name:"LimitCount"`
}

type DescribeInstancesRecentFailedOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最近失败操作详情列表。

		InstancesRecentFailedOperationSet []*InstancesRecentFailedOperationSet `json:"InstancesRecentFailedOperationSet,omitempty" name:"InstancesRecentFailedOperationSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesRecentFailedOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRecentFailedOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络计费类型配置。

		InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet,omitempty" name:"InternetChargeTypeConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchUserInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户appId。

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// 查询的用户实例信息。

		ResourceSet []*ResourceInfo `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 查询到的总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 标识下一页是否还有数据。

		Remain *bool `json:"Remain,omitempty" name:"Remain"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchUserInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchUserInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetChargeTypeConfig struct {

	// 网络计费模式。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 网络计费模式描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux机器密码需10到30位，至少包括三项([a-z],[A-Z],[0-9]和[()`~!@#$%^&*-+=_|{}[]:;'<>,.?/]的特殊符号)。<br><li>Windows机器密码需12到30位，至少包括三项([a-z],[A-Z],[0-9]和[()`~!@#$%^&*-+=_|{}[]:;'<>,.?/]的特殊符号),密码不允许包含用户名密码不允许以`/`符号开头。<br><li>如果实例即包含`Linux`实例又包含`Windows`实例，则密码复杂度限制按照`Windows`实例的限制

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br><br>可以通过以下方式获取可用的实例ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询实例ID。<br><li>通过调用接口 [DescribeInstances](../实例相关接口/DescribeInstances) ，取返回信息中的`InstanceId`获取实例ID。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](DescribeKeyPairs) ，取返回信息中的`KeyId`获取密钥对ID。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br><br>默认取值：FALSE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例机型配额列表。

		InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {

	// 释放地址<br>注意：此字段可能返回 null，表示取不到有效值。

	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
	// 不支持的网络类型，取值范围：<br> <li>BASIC：基础网络<br> <li>VPC1.0：私有网络VPC1.0<br> 注意：此字段可能返回 null，表示取不到有效值。

	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks"`
	// HDD本地存储属性 <br>注意：此字段可能返回 null，表示取不到有效值。

	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
	// 查询整个服务器

	InquiryWithEntireServer *string `json:"InquiryWithEntireServer,omitempty" name:"InquiryWithEntireServer"`
	// 需要的网络功能<br>注意：此字段可能返回 null，表示取不到有效值。

	RequireNetworkFeatures []*string `json:"RequireNetworkFeatures,omitempty" name:"RequireNetworkFeatures"`
	// 需要增强的服务<br>注意：此字段可能返回 null，表示取不到有效值。

	RequiredEnhancedService *RequiredEnhancedService `json:"RequiredEnhancedService,omitempty" name:"RequiredEnhancedService"`
	// GPU参数<br>注意：此字段可能返回 null，表示取不到有效值。

	GpuAttr *GpuAttr `json:"GpuAttr,omitempty" name:"GpuAttr"`
}

type DescribeImportSnapshotTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportSnapshotTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportSnapshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为1。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneHostConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 可用区过滤条件，支持zone和host-charge-type两张过滤条件。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneHostConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneHostConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 待扩容的系统盘配置信息。只支持扩容随实例购买的系统盘，且[系统盘类型](../数据结构#systemdisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。系统盘容量单位：GB。最小扩容步长：10G。关于系统盘类型的选择请参考硬盘产品简介。可选系统盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因系统盘类型的不同而有所差异。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceStatus struct {

	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// [实例状态](#instancestatus)。

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// EIP数量。默认值：1。

	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// 1

	IspId *int64 `json:"IspId,omitempty" name:"IspId"`
	// 1

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// [Deprecated] 该字段无实际功能，废弃

	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`
	// 内部参数，用于指定集群申请EIP。

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 内部参数。

	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`
	// EIP计费方式。<ul style="margin:0"><li>已开通带宽上移白名单的用户，可选值：<ul><li>付费（需额外开通共享带宽包白名单）</li><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li></ul>默认值：TRAFFIC_POSTPAID_BY_HOUR。</li><li>未开通带宽上移白名单的用户，EIP计费方式与其绑定的实例的计费方式一致，无需传递此参数。</li></ul>

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// EIP出带宽上限，单位：Mbps。<ul style="margin:0"><li>已开通带宽上移白名单的用户，可选值范围取决于EIP计费方式：<ul><li>BANDWIDTH_PACKAGE：1 Mbps 至 1000 Mbps</li><li>BANDWIDTH_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li><li>TRAFFIC_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li></ul>默认值：1 Mbps。</li><li>未开通带宽上移白名单的用户，EIP出带宽上限取决于与其绑定的实例的公网出带宽上限，无需传递此参数。</li></ul>

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// EIP出带宽上限，单位：Mbps。<ul style="margin:0"><li>已开通带宽上移白名单的用户，可选值范围取决于EIP计费方式：<ul><li>BANDWIDTH_PACKAGE：1 Mbps 至 1000 Mbps</li><li>BANDWIDTH_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li><li>TRAFFIC_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li></ul>默认值：1 Mbps。</li><li>未开通带宽上移白名单的用户，EIP出带宽上限取决于与其绑定的实例的公网出带宽上限，无需传递此参数。</li></ul>

	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例操作限制列表

		InstanceDeniedActionSet []*InstanceDeniedActions `json:"InstanceDeniedActionSet,omitempty" name:"InstanceDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分散置放群组信息

		DisasterRecoverGroups *DisasterRecoverGroups `json:"DisasterRecoverGroups,omitempty" name:"DisasterRecoverGroups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstanceInternetChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	InternetAccessible *InternetAccessibleModifyChargeType `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesModificationRequest struct {
	*tchttp.BaseRequest

	// 子机id

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstancesModificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesModificationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateRequest struct {
	*tchttp.BaseRequest

	// 待迁移的实例Id如："ins-38d920cd"

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 目标母机数组

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要嵌入的售卖池，默认不变更

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 需要跨可用迁移，可传入可用区列表

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 最大迁移带宽，默认40Mbps

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 最大迁移超时,默认86400秒

	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
	// 是否保留hostname

	HostNameReserved *bool `json:"HostNameReserved,omitempty" name:"HostNameReserved"`
}

func (r *ColdMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitRescueModeRequest struct {
	*tchttp.BaseRequest

	// 退出救援模式的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ExitRescueModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExitRescueModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionTimers struct {

	// ActionTimerId

	ActionTimerId *string `json:"ActionTimerId,omitempty" name:"ActionTimerId"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// TimerAction

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// ActionTime

	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Externals

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
}

type ResourceInfo struct {

	// 用户资源实例id。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 用户资源关键字。

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 用户资源实例关联的cvm实例id。

	RelationInstanceId *string `json:"RelationInstanceId,omitempty" name:"RelationInstanceId"`
	// 地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 该条数据搜索使用的类型，值为实例Id(instance),  实例名称(alias)， 实例内网ip(lan)，实例公网ip(wan)。

	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`
}

type DescribeUserZoneStatusRequest struct {
	*tchttp.BaseRequest

	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUserZoneStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserZoneStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否在正常重启失败后选择强制重启实例。取值范围：<br><li>TRUE：表示在正常重启失败后进行强制重启<br><li>FALSE：表示在正常重启失败后不进行强制重启<br><br>默认取值：FALSE。不支持与StopType参数同时制定

	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`
	// 关机类型。取值范围：SOFT：表示软关机HARD：表示硬关机SOFT_FIRST：表示优先软关机，失败再执行硬关机默认取值：SOFT。

	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Address struct {

	// `EIP`的`ID`，是`EIP`的唯一标识。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// `EIP`名称。

	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`
	// `EIP`状态。

	AddressState *string `json:"AddressState,omitempty" name:"AddressState"`
	// 弹性外网IP

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`，或是弹性网卡。

	BindedResourceId *string `json:"BindedResourceId,omitempty" name:"BindedResourceId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type InstanceTypeConfigStatusSet struct {

	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 描述信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 机型列表

	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`
}

type DescribeUserGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// 全局配置名称列表。

	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeUserGlobalConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGlobalConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeInstancesRequest struct {
	*tchttp.BaseRequest

	// 无

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ResumeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 [DescribeInstancesStatus](DescribeInstancesStatus) 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiagnosticReportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DiagnosticReportDataSet

		DiagnosticReportDataSet []*DiagnosticReportDataSet `json:"DiagnosticReportDataSet,omitempty" name:"DiagnosticReportDataSet"`
		// 报告总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiagnosticReportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiagnosticReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostResource struct {

	// cdh实例总cpu核数

	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// cdh实例可用cpu核数

	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`
	// cdh实例总内存大小（单位为:GiB）

	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// cdh实例可用内存大小（单位为:GiB）

	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`
	// cdh实例总磁盘大小（单位为:GiB）

	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// cdh实例可用磁盘大小（单位为:GiB）

	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
	// cdh实例磁盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

type RegionInfo struct {

	// 地域名称，例如，ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域描述，例如，华南地区(广州)

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域是否可用状态

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type DescribeImagesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// unImgId到deviceImageId的映射的数组

		ImageAttributeSet *ImageAttributeSet `json:"ImageAttributeSet,omitempty" name:"ImageAttributeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserMigrateTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据总数

		TotalCount *bool `json:"TotalCount,omitempty" name:"TotalCount"`
		// 冷迁移任务

		ColdMigrateInstanceTasks *UserMigrateTaskData `json:"ColdMigrateInstanceTasks,omitempty" name:"ColdMigrateInstanceTasks"`
		// 导入cbs任务

		ImportCbsTasks []*CbsMigrateTask `json:"ImportCbsTasks,omitempty" name:"ImportCbsTasks"`
		// 离线迁移任务（冷迁移任务 和导入cbs任务一起查询）

		OfflineMigrateTasks []*OfflineMigrateUserTaskData `json:"OfflineMigrateTasks,omitempty" name:"OfflineMigrateTasks"`
		// 导入完整cvm镜像。ImportFullCvmImage是离线实例迁移勾选了同时导入数据盘的镜像时记录的任务，目前tce暂时不支持ImportFullCvmImage。

		ImportFullCvmImageTasks []*UserMigrateTaskData `json:"ImportFullCvmImageTasks,omitempty" name:"ImportFullCvmImageTasks"`
		// 导入完整cvm镜像。ImportFullCvmImage是离线实例迁移勾选了同时导入数据盘的镜像时记录的任务，目前tce暂时不支持ImportFullCvmImage。

		MigrateInstanceTasks []*UserMigrateTaskData `json:"MigrateInstanceTasks,omitempty" name:"MigrateInstanceTasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserMigrateTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserMigrateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。可通过 [DescribeInstances](DescribeInstances) API返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指定有效的[镜像](/tcloud/Compute/CVM/292128/835305/mirr_overview)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](//console.{{conf.main_domain}}/cvm/image/list?imageType=PUBLIC_IMAGE&pageIndex=1&pageSize=20)查询。</li><li>通过调用接口 [DescribeImages](../镜像相关接口/DescribeImages) ，取返回信息中的`ImageId`字段。</li>

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 是否跳过实际执行逻辑

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceFamilyConfig struct {

	// 机型族名称的中文全称。

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 机型族名称的英文简称。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type DescribeInstanceAttachedDevicesRequest struct {
	*tchttp.BaseRequest

	// 实例ID,形如：ins-11112222

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttachedDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAttachedDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAvailableZonesRequest struct {
	*tchttp.BaseRequest

	// available-zone-set 可用区例如：Name: available-zone-setValues:["ap-guangzhou-1", "ap-guangzhou-2"]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUserAvailableZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAvailableZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds *string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 1

	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *InquiryPriceRenewAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示带宽调整为对应大小之后的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfoItem struct {

	// 信息项名称

	// 信息项对应的值

}

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesReturnableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesReturnableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserZoneStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区计费类型状态列表。

		UserZoneStatus []*UserZoneStatusItem `json:"UserZoneStatus,omitempty" name:"UserZoneStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserZoneStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserZoneStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 指定具体的定时器信息

	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
	// 定时器id列表，可以通过DescribeInstancesActionTimer接口查询。

	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
}

func (r *ModifyInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 分散置放群组ID列表

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *QueryDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 分散置放群组ID，可使用DescribeDisasterRecoverGroups接口获取

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 是否强制更换实例宿主机。取值范围：<br><li>TRUE：表示允许实例更换宿主机，允许重启实例<br><li>FALSE：不允许实例更换宿主机，只在当前宿主机上加入置放群组。这可能导致更换置放群组失败<br><br>默认取值：FALSE

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *ModifyInstancesDisasterRecoverGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesDisasterRecoverGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// 可用区名称，例如，ap-guangzhou-3

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区描述，例如，广州三区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区状态

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type SwitchParameterRenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 是否跳过实际执行逻辑。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterRenewHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 定时任务ID列表

	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
	// 操作时间

	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`
}

func (r *UpdateInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstance struct {

	// Placement结构

	Placement *QueryPlacement `json:"Placement,omitempty" name:"Placement"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// cpu核数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例拥有者AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例Id

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 所在宿主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// Ipv6地址

	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
	// 当前状态

	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	// 当前状态

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 所有者ownerUin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 子机机型规格

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 子机机型类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// Quota使用情况

	NodeQuota []*int64 `json:"NodeQuota,omitempty" name:"NodeQuota"`
	// 系统盘详情

	SystemDisk *QuerySystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 数据盘详情

	DataDisks []*QueryDataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type DescribeInstanceTypeQuotaRequest struct {
	*tchttp.BaseRequest

	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceTypeQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GpuAttr struct {

	// 类型 <br>注意：此字段可能返回 null，表示取不到有效值。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type InstanceFamilyItem struct {

	// 实例族。

	// 优先级。

	// 实例类型信息列表。

	// 实例类型名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
}

type ProjectSpecification struct {

	// 资源类型,默认instance

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 项目Id

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 1

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 1

	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *SwitchParameterModifyInstancesChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 是否跳过实际执行逻辑。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *TerminateHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetAccessible struct {

	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致。

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>公网带宽大于0时必须设置为True,默认开通公网IP；当公网带宽为0，则不允许分配公网IP。

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
	// 网络模式:  移动:"CMCC" 、联通:"CTCC" 、电信:"CUCC"、外网CAP: "BGP"。在三网模式下（移动、联通、电信），必须为带宽包计费模式。即：必须携带InternetChargeType参数, 且值必须为 BANDWIDTH_PACKAGE。   BGP模式下无此限制。该接口不支持多运营商模式，即参数InternetServiceProvider参数不能是Multi-operator。

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
}

type ModifyInstancesActionTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesActionTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest

	// 查询的实例id, 如ins-38dk3j

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedZonesRequest struct {
	*tchttp.BaseRequest

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](DescribeInstanceTypeConfigs)来获得最新的规格表或参见[实例类型](/tcloud/Compute/CVM/292128/484318/specification)描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribeRecommendedZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecommendedZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryResourceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否在线升级配置

	Online *bool `json:"Online,omitempty" name:"Online"`
}

func (r *InquiryResourceResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryResourceResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroups struct {

	// 当前用户已经创建的置放群组数量。

	// 置放群组id。

	// 置放群组内最大容纳云服务器数量。

	// uuid

	// 帐号的所有者uin

	// 标签

	// 类型

	// 创建时间

	// 置放群组名称

	// 分区数

	// 策略

	// 亲和度，匹配度

}

type CancelAuditMarketImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelAuditMarketImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelAuditMarketImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。

		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet"`
		// 符合要求的镜像数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type SwitchParameterResetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例订单详情信息。

		InstanceOrder *InstanceOrder `json:"InstanceOrder,omitempty" name:"InstanceOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 1

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 1

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LaunchTemplatesInfo struct {

	// 实例启动模板ID。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
	// 实例启动模板名。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`
	// 实例启动模版本信息。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateVersionData *LaunchTemplateVersionData `json:"LaunchTemplateVersionData,omitempty" name:"LaunchTemplateVersionData"`
	// 实例启动模板描述。 注意：此字段可能返回 null，表示取不到有效值。

	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`
	// 创建该模板的用户UIN。 注意：此字段可能返回 null，表示取不到有效值。

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 创建该模板的时间。 注意：此字段可能返回 null，表示取不到有效值。

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字符串状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 任务状态

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiagnosticReportsRequest struct {
	*tchttp.BaseRequest

	// 报告id

	ReportIds []*string `json:"ReportIds,omitempty" name:"ReportIds"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiagnosticReportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiagnosticReportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceOrder struct {
}

type DescribeImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 需要共享的镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInstanceQuotaRequest struct {
	*tchttp.BaseRequest

	// zone - String - 是否必填：否 - 按照可用区过滤。instance-charge-type - String - 是否必填：否 - 实例计费模式instance-charge-type支持类型：PREPAID、POSTPAID_BY_HOUR、SPOTPAID最大限制为10，value最大限制为5

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUserInstanceQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInstanceQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalDiskType struct {

	// 本地磁盘类型。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 本地磁盘属性。

	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`
	// 本地磁盘最小值。

	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`
	// 本地磁盘最大值。

	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 购买时本地盘是否为必选。取值范围：<br><li>REQUIRED：表示必选<br><li>OPTIONAL：表示可选。

	Required *string `json:"Required,omitempty" name:"Required"`
}

type ModifyLaunchTemplateRequest struct {
	*tchttp.BaseRequest

	// 实例启动模板名称。长度为2~128个英文或中文字符。

	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 指定有效的镜像ID，格式形如img-xxx。镜像类型分为四种：：<br><li>公共镜像<br><li>自定义镜像<br><li>共享镜像<br><li>服务市场镜像<br>公共镜像、自定义镜像、共享镜像的镜像ID可通过登录控制台查询；服务镜像市场的镜像ID可通过云市场查询。通过调用接口 DescribeImages ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的ImageId字段。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例启动模板版本描述。长度为2~256个英文或中文字符。

	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`
	// 实例机型。不同实例机型指定了不同的资源规格。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// Array of DataDisk	实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 购买实例数量。包年包月实例取值范围：[1，300]，按量计费实例取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见CVM实例购买限制。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 云服务器的主机名<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。

	ActionTimer *string `json:"ActionTimer,omitempty" name:"ActionTimer"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。

	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。

	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`
	// 实例计费类型。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例<br><li>FALSE：表示关闭实例保护，允许通过api接口删除实例<br>默认取值：FALSE。

	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
	// 实例模版id

	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
}

func (r *ModifyLaunchTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLaunchTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageSharedAccount struct {

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 账户ID

	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
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

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailabilityZone struct {

	// 地域ID。

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区状态。

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type MonitorServiceItem struct {

	// 是否启用

	Enabled *string `json:"Enabled,omitempty" name:"Enabled"`
}

type ImportCbsRequest struct {
	*tchttp.BaseRequest

	// 用于测试参数合法性。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 用户需要导入的数据盘镜像存放的COS链接。

	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`
	// 导入目的云盘的ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 任务名称，用于控制台展示。

	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *ImportCbsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCbsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstanceInternetChargeTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyInstanceInternetChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {

	// 描述了实例价格。

	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
	// 描述了网络价格。

	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type RenewAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceFamilyItemArchitecture struct {

	// 实例族。

	// 优先级。

	// 实例类型信息列表。

	// 实例类型名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// CPU架构信息

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户账号属性对象

		AccountAttributeSet []*AccountAttributeSet `json:"AccountAttributeSet,omitempty" name:"AccountAttributeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>`Linux`实例密码必须8到16位，至少包括两项`[a-z，A-Z]、[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>`Windows`实例密码必须12到16位，至少包括三项`[a-z]，[A-Z]，[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>如果实例即包含`Linux`实例又包含`Windows`实例，则密码复杂度限制按照`Windows`实例的限制。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 待重置密码的实例操作系统用户名。不得超过64个字符。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportImageOsListSupported struct {

	// linux系统列表

	Linux []*string `json:"Linux,omitempty" name:"Linux"`
	// windows系统列表

	Windows []*string `json:"Windows,omitempty" name:"Windows"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest

	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 只能为0

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyPairRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportImageRequest struct {
	*tchttp.BaseRequest

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// COS/CSP Bucket名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 1

	Edge *bool `json:"Edge,omitempty" name:"Edge"`
	// 导出文件前缀

	FileNamePrefix *string `json:"FileNamePrefix,omitempty" name:"FileNamePrefix"`
	// 1

	ExportFormat *string `json:"ExportFormat,omitempty" name:"ExportFormat"`
}

func (r *ExportImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesReturnableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可退还实例详细信息列表。

		InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitempty" name:"InstanceReturnableSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesReturnableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](DescribeKeyPairs) ，取返回信息中的 `KeyId` 获取密钥对ID。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区列表信息

		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest

	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 密钥对创建后所属的项目ID。<br><br>可以通过以下方式获取项目ID：<br><li>通过[项目列表](/resource-management/project-manage)查询项目ID。<br><li>通过调用接口 [DescribeProjects](/tcloud/api/TCenter_CAT/TCENTER_SUBCAT/APIs/组织与项目（tpo）/版本（2020-09-20）/Project相关接口/DescribeProjects)，取返回信息中的 `projectId ` 获取项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 密钥对的公钥内容，`OpenSSH RSA` 格式。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyPairRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口[DescribeInstanceTypeConfigs](DescribeInstanceTypeConfigs)来获得最新的规格表。若不指定该参数，则默认机型为S1.SMALL1。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 指定有效的[镜像](/tcloud/Compute/CVM/292128/835305/mirr_overview)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](//console.{{conf.main_domain}}/cvm/image/list?imageType=PUBLIC_IMAGE&pageIndex=1&pageSize=20)查询。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，购买时可指定多个数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，那么InstanceCount参数只能为1。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。如果不指定则默认显示

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。若不指定该参数，则默认不绑定安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 内部参数，购买来源。前端调用的来源是MC

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
}

func (r *InquiryPriceRunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRunInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshInternalUserEnvironmentRequest struct {
	*tchttp.BaseRequest
}

func (r *RefreshInternalUserEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshInternalUserEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 1

	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *RenewAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 服务迁移任务taskId

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeMigrateTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID，可以用于查询任务状态

		TaskId []*string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘类型。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_HSSD：增强型SSD云硬盘<br><br>默认取值：LOCAL_BASIC。<br><br>该参数对`ResizeInstanceDisk`接口无效。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同。默认值为0，表示不购买数据盘。更多限制详见产品文档。

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘是否随子机销毁。取值范围：
	// <li>TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘
	// <li>FALSE：子机销毁时，保留数据盘<br>
	// 默认取值：TRUE<br>
	// 该参数目前仅用于 `RunInstances` 接口。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 数据盘指定的存储池。

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云盘的自动备份策略

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 云盘的快照id

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type HostPrice struct {

	// 描述了cdh实例相关的价格信息

	HostPrice *ItemPrice `json:"HostPrice,omitempty" name:"HostPrice"`
}

type ModifyLaunchTemplateNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLaunchTemplateNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLaunchTemplateNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneCdhInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用宿主机机型配置信息列表

		HostTypeQuotaSet *string `json:"HostTypeQuotaSet,omitempty" name:"HostTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneCdhInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneCdhInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 1

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedActions struct {

	// 操作名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 返回值

	Code *string `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeImageSnapshotStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照状态，存在以下三种状态：NORMAL(正常)、CREATING（创建中）、NULL（不存在）

		CbsStatus *string `json:"CbsStatus,omitempty" name:"CbsStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSnapshotStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSnapshotStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesActionTimerRequest struct {
	*tchttp.BaseRequest

	// 定时器id数组

	ActionTimerIds []*string `json:"ActionTimerIds,omitempty" name:"ActionTimerIds"`
	// 实例id数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 定时任务执行之间，格式如：2018-05-01 19:00:00，必须大于当前时间5分钟。

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// 执行时间的结束范围，用于条件筛选，格式如2018-05-01 19:00:00。

	EndActionTime *string `json:"EndActionTime,omitempty" name:"EndActionTime"`
	// 执行时间的开始范围，用于条件筛选，格式如2018-05-01 19:00:00。

	StartActionTime *string `json:"StartActionTime,omitempty" name:"StartActionTime"`
	// 定时器状态列表数组，可选值为：UNDO（未执行）， DOING（正在执行i），DONE（执行完成）。

	StatusList []*string `json:"StatusList,omitempty" name:"StatusList"`
}

func (r *DescribeInstancesActionTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesActionTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表信息

		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`
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

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br><br>可以通过以下方式获取可用的实例ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询实例ID。<br><li>通过调用接口 [DescribeInstances](../实例相关接口/DescribeInstances) ，取返回信息中的 `InstanceId` 获取密钥对ID。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 密钥对ID列表，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](DescribeKeyPairs) ，取返回信息中的 `KeyId` 获取密钥对ID。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br><br>默认取值：FALSE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditMarketImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditMarketImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditMarketImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// DryRun

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 修改数据盘

	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *InquiryPriceModifyInstancesChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 1

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 1

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 1

	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *ModifyInstancesChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceConnectivity struct {

	// 实例`ID`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 默认远程登录端口连通性状态

	DefaultLoginPortConnectivity *bool `json:"DefaultLoginPortConnectivity,omitempty" name:"DefaultLoginPortConnectivity"`
	// ping包是否可达

	ICMPConnectivity *bool `json:"ICMPConnectivity,omitempty" name:"ICMPConnectivity"`
}

type DescribeAddressBandwidthConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressBandwidthConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressBandwidthConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 1

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 1

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SwitchParameterResizeInstanceDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResizeInstanceDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像共享信息

		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`
		// 最大共享权限

		MaxSharePermission *int64 `json:"MaxSharePermission,omitempty" name:"MaxSharePermission"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchTemplatesInfoRequest struct {
	*tchttp.BaseRequest

	// 启动模板ID，一个或者多个启动模板ID。若未指定，则显示用户所有模板。

	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitempty" name:"LaunchTemplateIds"`
	// 按照【LaunchTemplateName】进行过滤。

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLaunchTemplatesInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLaunchTemplatesInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SharePermissionSet struct {

	// 镜像分享时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 镜像分享的账户ID

	Account *string `json:"Account,omitempty" name:"Account"`
}

type CheckInstancesConnectivityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例连通性结果集合。

		InstanceConnectivitySet []*InstanceConnectivity `json:"InstanceConnectivitySet,omitempty" name:"InstanceConnectivitySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckInstancesConnectivityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckInstancesConnectivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CDH实例续费价格信息

		Price *HostPrice `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest

	// 密钥对ID，密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥 ID：<br><li>通过登录[控制台](//console.{{conf.main_domain}}/cvm/index)查询密钥 ID。<br><li>通过调用接口 [DescribeKeyPairs](DescribeKeyPairs) ，取返回信息中的 `KeyId` 获取密钥对 ID。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyKeyPairAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持`InternetMaxBandwidthOut`参数。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 带宽生效的终止时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过[DescribeInstances](DescribeInstances)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 1

	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`
	// 1

	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示调整成对应机型实例的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// CDH实例显示名称。可任意命名，但不得超过60个字符。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 自动续费标识。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchTemplateRequest struct {
	*tchttp.BaseRequest

	// 实例启动模板名称。长度为2~128个英文或中文字符。

	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 指定有效的镜像ID，格式形如img-xxx。镜像类型分为四种：：<br><li>公共镜像<br><li>自定义镜像<br><li>共享镜像<br><li>服务市场镜像<br>公共镜像、自定义镜像、共享镜像的镜像ID可通过登录控制台查询；服务镜像市场的镜像ID可通过云市场查询。通过调用接口 DescribeImages ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的ImageId字段。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例启动模板版本描述。长度为2~256个英文或中文字符。

	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`
	// 实例机型。不同实例机型指定了不同的资源规格。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// Array of DataDisk	实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 购买实例数量。包年包月实例取值范围：[1，300]，按量计费实例取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见CVM实例购买限制。

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 云服务器的主机名<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。

	ActionTimer *string `json:"ActionTimer,omitempty" name:"ActionTimer"`
	// 置放群组id，仅支持指定一个。

	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。

	UserData *string `json:"UserData,omitempty" name:"UserData"`
	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。

	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。

	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`
	// 实例计费类型。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例<br><li>FALSE：表示关闭实例保护，允许通过api接口删除实例<br>默认取值：FALSE。

	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

func (r *CreateLaunchTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLaunchTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLaunchTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLaunchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewHostsRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的CDH实例ID。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`
	// 是否跳过实际执行逻辑。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InquiryPriceRenewHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务流转ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportFullCvmImageRequest struct {
	*tchttp.BaseRequest

	// 导入任务名称，不超过60个字符。

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// 导入镜像的目标子机实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 系统盘镜像cos url，cos地域需要跟子机所在地域保持一致。

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像名称，IsCreate为`TRUE`则为必填，不超过20个字符。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像描述，不超过60个字符。

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 数据盘镜像数组，最多4个数据盘镜像，传入的数据盘数量必须等于子机上挂载的盘的数量。

	DataDisks []*DataDisks `json:"DataDisks,omitempty" name:"DataDisks"`
	// 用于测试参数合法性，默认为`FALSE`。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 是否制作整机镜像，默认为`FALSE`。

	IsCreate *bool `json:"IsCreate,omitempty" name:"IsCreate"`
}

func (r *ImportFullCvmImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportFullCvmImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[DescribeInstances](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 内部参数，释放弹性IP。

	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
	// 试运行。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 持续时间

	RemainTime *int64 `json:"RemainTime,omitempty" name:"RemainTime"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailPrices struct {

	// 镜像价格

	ImagePrice *ItemPrice `json:"ImagePrice,omitempty" name:"ImagePrice"`
	// 数据盘价格

	DataDisksPrice []*ItemPrice `json:"DataDisksPrice,omitempty" name:"DataDisksPrice"`
	// cpu内存价格

	CpuMemPrice *ItemPrice `json:"CpuMemPrice,omitempty" name:"CpuMemPrice"`
	// 系统盘价格

	SystemDiskPrice *ItemPrice `json:"SystemDiskPrice,omitempty" name:"SystemDiskPrice"`
}

type KeyPair struct {

	// 密钥对的`ID`，是密钥对的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥对名称。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 密钥对描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 密钥对的纯文本公钥。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// 密钥关联的镜像数。

	AssociatedImageCount *int64 `json:"AssociatedImageCount,omitempty" name:"AssociatedImageCount"`
	// 密钥关联的实例数。

	AssociatedInstanceCount *int64 `json:"AssociatedInstanceCount,omitempty" name:"AssociatedInstanceCount"`
	// 密钥关联的实例`ID`列表。

	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`
	// 密钥对创建日期

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type InstanceRefundsSet struct {

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 退款数额。

	Refunds *float64 `json:"Refunds,omitempty" name:"Refunds"`
}

type DescribeZoneHostForSellStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用宿主机可用区售罄状态

		HostForSellZoneStatus *string `json:"HostForSellZoneStatus,omitempty" name:"HostForSellZoneStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneHostForSellStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneHostForSellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest

	// 准备删除的镜像Id列表

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 迁移进度。

		Progress *float64 `json:"Progress,omitempty" name:"Progress"`
		// 任务状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 任务名称。

		JobName *string `json:"JobName,omitempty" name:"JobName"`
		// 任务ID。

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 用户AppId。

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// 账户ID。

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// CVM子机uuid。

		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
		// 实例ID。

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 镜像COS链接。

		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
		// 镜像大小。

		DataSize *uint64 `json:"DataSize,omitempty" name:"DataSize"`
		// 地域信息。

		Region *string `json:"Region,omitempty" name:"Region"`
		// 任务创建时间。

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 任务结束时间。

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 任务类型。

		Action *string `json:"Action,omitempty" name:"Action"`
		// 任务失败的错误原因

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrateTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMigrateTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressChargePrepaid struct {

	// 购买实例的时长

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标志

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type QuerySystemDisk struct {

	// 系统盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘类型：LOCAL_BASIC、CLOUD_BASIC、LOCAL_SSD、CLOUD_SSD、CLOUD_PREMIUM、CLOUD_ENHANCEDSSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

type DescribeInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例静态配置信息列表。

		InstanceConfigInfos []*InstanceConfigInfoItemArchitecture `json:"InstanceConfigInfos,omitempty" name:"InstanceConfigInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 1

	KeepAddressIdBindWithEniPip *bool `json:"KeepAddressIdBindWithEniPip,omitempty" name:"KeepAddressIdBindWithEniPip"`
	// 1

	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDisasterRecoverGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

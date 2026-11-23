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

type SwitchParameterRenewDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费订单参数。

		DiskOrder []*DiskOrder `json:"DiskOrder,omitempty" name:"DiskOrder"`
		// 到期时间对齐到实例时，是否有云盘需要续费。该参数只有在对齐实例到期时间时才返回。

		NeedAlignWithInstance *bool `json:"NeedAlignWithInstance,omitempty" name:"NeedAlignWithInstance"`
		// 云盘续费后的到期时间。

		NewDeadlineSet []*string `json:"NewDeadlineSet,omitempty" name:"NewDeadlineSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSnapOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户快照总大小

		TotalSize *float64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 用户快照总大小（用于计费）

		RealTradeSize *float64 `json:"RealTradeSize,omitempty" name:"RealTradeSize"`
		// 快照免费额度

		FreeQuota *float64 `json:"FreeQuota,omitempty" name:"FreeQuota"`
		// 快照总个数

		TotalNums *int64 `json:"TotalNums,omitempty" name:"TotalNums"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSnapOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSnapOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {

	// 实例所属的[可用区](/document/api/213/9452#zone)ID。该参数也可以通过调用  [DescribeZones](/document/product/213/15707) 的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 实例所属的独享集群ID。作为入参时，表示对指定的CdcId独享集群的资源进行操作，可为空。 作为出参时，表示资源所属的独享集群的ID，可为空。

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 围笼Id。作为入参时，表示对指定的CageId的资源进行操作，可为空。 作为出参时，表示资源所属围笼ID，可为空。

	CageId *string `json:"CageId,omitempty" name:"CageId"`
	// 独享集群名字。作为入参时，忽略。作为出参时，表示云硬盘所属的独享集群名，可为空。

	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// cvm独享集群id。作为入参时，表示对指定独享集群的资源进行操作。作为出参时，表示资源所属的独享集群的ID，可为空。

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

type SwitchParameterModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更云盘类型计费订单参数。

		DiskOrder *ResizeDiskOrder `json:"DiskOrder,omitempty" name:"DiskOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterModifyDiskAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 过滤条件。参数不支持同时指定`SnapshotIds`和`Filters`。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：`snap-11112222`。<br><li>snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。<br><li>snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/api/213/9452#zone)过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据快照的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDiskResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述第个地域的资源详情。

		ResourcesDetailSet []*ResourcesDetail `json:"ResourcesDetailSet,omitempty" name:"ResourcesDetailSet"`
		// 描述指定地域云硬盘总体使用情况。

		DiskOverview []*DiskOverview `json:"DiskOverview,omitempty" name:"DiskOverview"`
		// 描述指定地域快照总体使用情况。

		SnapOverview *SnapOverview `json:"SnapOverview,omitempty" name:"SnapOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDiskResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDiskResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDisksRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长。如果在该参数中指定CurInstanceDeadline，则会按对齐到子机到期时间来续费。如果是批量续费询价，该参数与Disks参数一一对应，元素数量需保持一致。

	DiskChargePrepaids []*DiskChargePrepaid `json:"DiskChargePrepaids,omitempty" name:"DiskChargePrepaids"`
	// 指定云盘新的到期时间，形式如：2017-12-17 00:00:00。参数`NewDeadline`和`DiskChargePrepaids`是两种指定询价时长的方式，两者必传一个。

	NewDeadline *string `json:"NewDeadline,omitempty" name:"NewDeadline"`
	// 云盘所属项目ID。 如传入则仅用于鉴权。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceRenewDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoSnapshotPolicyAttributeRequest struct {
	*tchttp.BaseRequest

	// 定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 通过该定期快照策略创建的快照保留天数，该参数不可与`IsPermanent`参数冲突，即若定期快照策略设置为永久保留，`RetentionDays`应置0。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskDeniedAction struct {

	// 云硬盘ID，比如"disk-7rf24jvb"

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 具体的云硬盘的操作掩码信息

	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type ApplySnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplySnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br><li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br><li>portable - Array of Bool - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br><li>project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)<br><li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/api/213/9452#zone)过滤。<br><li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br><li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。

	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`
	// 内部参数，用于支持搜索框搜索。

	InnerSearch *string `json:"InnerSearch,omitempty" name:"InnerSearch"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {

	// 快照ID。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 创建此快照的云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 创建此快照的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 创建此快照的云硬盘大小。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制快照拷贝中。

	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`
	// 快照名称，用户自定义的快照别名。调用[ModifySnapshotAttribute](/document/product/362/15650)可修改此字段。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照创建进度百分比，快照创建成功后此字段恒为100。

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 快照的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 快照到期时间。如果快照为永久保留，此字段为空。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 是否为加密盘创建的快照。取值范围：<br><li>true：该快照为加密盘创建的<br><li>false:非加密盘创建的快照。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 是否为永久快照。取值范围：<br><li>true：永久快照<br><li>false：非永久快照。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 快照正在跨地域复制的目的地域，默认取值为[]。

	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	// 是否为跨地域复制的快照。取值范围：<br><li>true：表示为跨地域复制的快照。<br><li>false:本地域的快照。

	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	// 快照关联的镜像列表。

	Images []*Image `json:"Images,omitempty" name:"Images"`
	// 快照关联的镜像个数。

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 快照类型，目前该项取值可以为PRIVATE_SNAPSHOT或者SHARED_SNAPSHOT

	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`
	// 快照当前被共享数

	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`
	// 快照的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 快照的TCE项目ID。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 加密算法。

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
}

type PoolCustomAttributes struct {

	// cpu架构。

	CpuArchitecture *string `json:"CpuArchitecture,omitempty" name:"CpuArchitecture"`
	// cpu型号。

	CpuModel *string `json:"CpuModel,omitempty" name:"CpuModel"`
}

type InquirePriceModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了调整云盘额外性能时对应的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceModifyDiskExtraPerformanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewDisksRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的云硬盘ID。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的续费时长。在云盘与挂载的实例一起续费的场景下，可以指定参数CurInstanceDeadline，此时云盘会按对齐到实例续费后的到期时间来续费。

	DiskChargePrepaids []*DiskChargePrepaid `json:"DiskChargePrepaids,omitempty" name:"DiskChargePrepaids"`
}

func (r *SwitchParameterRenewDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要解绑定期快照策略的云盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 要解绑的定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *UnbindAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopySnapshotCrossRegionsRequest struct {
	*tchttp.BaseRequest

	// 需要跨地域复制的源快照ID，可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照需要复制到的目标地域，各地域的标准取值可通过接口[DescribeRegions](https://cloud.tencent.com/document/product/213/9456)查询，且只能传入支持快照的地域。

	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`
	// 新复制快照的名称，如果不传，则默认取值为“Copied snap-11112222 from 地域名”。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CopySnapshotCrossRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySnapshotCrossRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 要查询快照的ID。可通过DescribeSnapshots查询获取。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DescribeSnapshotSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotSharePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateDisksRequest struct {
	*tchttp.BaseRequest

	// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB。步长均为10GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 购买数量。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。默认取值为1。

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 付费模式，目前只有预付费，即只能取值为PREPAID。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 指定创建系统盘或数据盘。取值范围：<br><li>SYSTEM_DISK：表示系统盘<br><li>DATA_DISK：表示数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长、是否设置自动续费等属性。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 云硬盘所在的位置。通过该参数可以指定云硬盘所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘绑定的标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 定期快照策略ID。传入此参数时，当云硬盘创建成功后将会自动绑定该定期快照策略

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。

	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`
	// 资源池组

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 项目ID

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
}

func (r *SwitchParameterCreateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDetail struct {

	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例已挂载数据盘的数量。

	AttachedDiskCount *uint64 `json:"AttachedDiskCount,omitempty" name:"AttachedDiskCount"`
	// 实例最大可挂载数据盘的数量。

	MaxAttachCount *uint64 `json:"MaxAttachCount,omitempty" name:"MaxAttachCount"`
}

type SnapshotGroupDeniedAction struct {

	// 快照组ID

	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
	// 具体的快照的操作掩码信息

	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type DescribeDiskSupportFeaturesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDiskSupportFeaturesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskSupportFeaturesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 快照ID列表

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *DescribeSnapshotsDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费订单参数。

		DiskOrder []*DiskOrder `json:"DiskOrder,omitempty" name:"DiskOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterCreateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ApplySnapshotGroupRequest struct {
	*tchttp.BaseRequest

	// 回滚的快照组ID

	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
	// 回滚的快照ID、云盘ID列表

	ApplyDisks []*ApplyDisk `json:"ApplyDisks,omitempty" name:"ApplyDisks"`
}

func (r *ApplySnapshotGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘扩容后的大小，单位为GB，必须大于当前云硬盘大小。取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB，步长均为10GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskOrder struct {

	// 产品类别。

	GoodsCategoryId *uint64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 产品数量。

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品详情。

	GoodsDetail *ResizeGoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 地域ID。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 付费模式。

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 云盘类别。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 子产品ID。

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
}

type Tag struct {

	// 标签健。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ApplySnapshotRequest struct {
	*tchttp.BaseRequest

	// 快照ID, 可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照原云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *ApplySnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotGroupRequest struct {
	*tchttp.BaseRequest

	// 快照组ID

	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
	// 是否同时删除快照组关联的镜像；取值为false，表示不删除快照组绑定的镜像，此时，如果快照组有绑定的镜像，删除会失败；取值为true，表示同时删除快照组绑定的镜像；默认值为false

	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
	// 解散快照组。取值为true，表示仅解散快照组，不删除快照组关联的快照；取值为false，则会同时删除快照组关联的快照。默认取值为false

	OnlyDismiss *bool `json:"OnlyDismiss,omitempty" name:"OnlyDismiss"`
}

func (r *DeleteSnapshotGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的云硬盘ID列表。

		DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest

	// 需迁移的云盘实例ID列表，**当前仅支持一次传入一块云盘**。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云盘变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。<br>**当前不支持类型降级**

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘所属项目ID。 如传入则仅用于鉴权。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceModifyDiskAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyDiskAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskFeature struct {

	// 云硬盘使用的计费系统。取值范围：<br><li>bmppro：计费系统<br><li>bmp：计量系统

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
}

type CopySnapshotCrossRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照跨地域复制的结果，如果请求下发成功，则返回相应地地域的新快照ID，否则返回Error。

		SnapshotCopyResultSet []*SnapshotCopyResult `json:"SnapshotCopyResultSet,omitempty" name:"SnapshotCopyResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopySnapshotCrossRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySnapshotCrossRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘的详细信息列表。

		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 快照组列表详情

		SnapshotGroupSet []*SnapshotGroup `json:"SnapshotGroupSet,omitempty" name:"SnapshotGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest

	// 需迁移的云盘实例ID列表，**当前仅支持一次传入一块云盘**。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云盘变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。<br>**当前不支持类型降级**

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

func (r *SwitchParameterModifyDiskAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterModifyDiskAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要绑定的定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 要绑定的云硬盘ID列表，一次请求最多绑定80块云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *BindAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDiskNumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各个云服务器已挂载和可挂载弹性云盘的数量。

		AttachDetail []*AttachDetail `json:"AttachDetail,omitempty" name:"AttachDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDiskNumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDiskNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 有效的定期快照策略数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 定期快照策略列表。

		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘绑定的定期快照数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云盘绑定的定期快照列表。

		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskConfigQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘配置列表。

		DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitempty" name:"DiskConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskConfigQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStoragesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为100，最大值为1000。关于Limit的更进一步介绍请参考 API 简介中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 按照一个或者多个云硬盘ID查询。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 内部参数，用于支持搜索框搜索。

	InnerSearch *string `json:"InnerSearch,omitempty" name:"InnerSearch"`
}

func (r *DescribeBlockStoragesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockStoragesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要查询的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskAssociatedSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 快照大小，单位GB

	SnapSize *float64 `json:"SnapSize,omitempty" name:"SnapSize"`
}

func (r *InquiryPriceSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {

	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest

	// 额外购买的云硬盘性能值，单位MB/s。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 需要创建快照的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *ModifyDiskExtraPerformanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 通过该定期快照策略创建的快照保留天数，默认保留7天，该参数不可与`IsPermanent`参数冲突，即若定期快照策略设置为永久保留，本参数应置0。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 是否创建定期快照的执行策略。TRUE表示只需获取首次开始备份的时间，不实际创建定期快照策略，FALSE表示创建，默认为FALSE。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskSupportFeaturesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云硬盘产品特性详情。

		DiskFeature *DiskFeature `json:"DiskFeature,omitempty" name:"DiskFeature"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskSupportFeaturesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskSupportFeaturesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcesDetail struct {

	// 资源所在地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 快照用户数量配额。

	SnapshotNumberTotal *uint64 `json:"SnapshotNumberTotal,omitempty" name:"SnapshotNumberTotal"`
	// 已使用的快照数量。

	SnapshotNumberUsed *uint64 `json:"SnapshotNumberUsed,omitempty" name:"SnapshotNumberUsed"`
	// 快照容量配额。

	SnapshotCapacityTotal *uint64 `json:"SnapshotCapacityTotal,omitempty" name:"SnapshotCapacityTotal"`
	// 已使用的快照容量，单位GB。

	SnapshotCapacityUsed *uint64 `json:"SnapshotCapacityUsed,omitempty" name:"SnapshotCapacityUsed"`
	// 云盘数量。

	DiskNumber *int64 `json:"DiskNumber,omitempty" name:"DiskNumber"`
	// 当前地域是否包含新购买的云盘。

	NewDiskFlag *bool `json:"NewDiskFlag,omitempty" name:"NewDiskFlag"`
	// 当前地域是否支持跨地域复制

	SnapshotSupportCrossCopy *bool `json:"SnapshotSupportCrossCopy,omitempty" name:"SnapshotSupportCrossCopy"`
	// 回收站内的云盘数量。

	DiskNumberIsolated *uint64 `json:"DiskNumberIsolated,omitempty" name:"DiskNumberIsolated"`
}

type DescribeDiskConfigQuotaRequest struct {
	*tchttp.BaseRequest

	// 查询类别，取值范围。<br><li>INQUIRY_CBS_CONFIG：查询云盘配置列表<br><li>INQUIRY_CVM_CONFIG：查询云盘与实例搭配的配置列表。

	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`
	// 查询一个或多个[可用区](/document/api/213/9452#zone)下的配置。

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费<br><li>POSTPAID_BY_HOUR：后付费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskTypes []*string `json:"DiskTypes,omitempty" name:"DiskTypes"`
	// 系统盘或数据盘。取值范围：<br><li>SYSTEM_DISK：表示系统盘<br><li>DATA_DISK：表示数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 实例机型。

	DeviceClasses []*string `json:"DeviceClasses,omitempty" name:"DeviceClasses"`
	// 按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)

	InstanceFamilies []*string `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`
	// 实例CPU核数。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存大小。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// INQUIRY_RESIZE、INQUIRY_CREATE

	InnerInquiryType *string `json:"InnerInquiryType,omitempty" name:"InnerInquiryType"`
}

func (r *DescribeDiskConfigQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了续费云盘的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Policy struct {

	// 选定周一到周日中需要创建快照的日期，取值范围：[0, 6]。0表示周一触发，依此类推。

	DayOfWeek []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	// 指定定期快照策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。

	Hour []*uint64 `json:"Hour,omitempty" name:"Hour"`
}

type InquiryPriceSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照商品单价，每GB每小时多少钱

		Price *float64 `json:"Price,omitempty" name:"Price"`
		// 总收费

		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskOrder struct {

	// 产品类别。

	GoodsCategoryId *uint64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 产品数量。

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品详情。

	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 地域ID。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 付费模式。

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 云盘类别。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type SnapshotGroup struct {

	// 快照组ID

	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
	// 快照组类型。NORMAL: 普通快照组，非一致性快照

	SnapshotGroupType *string `json:"SnapshotGroupType,omitempty" name:"SnapshotGroupType"`
	// 快照组是否包含系统盘快照

	ContainRootSnapshot *bool `json:"ContainRootSnapshot,omitempty" name:"ContainRootSnapshot"`
	// 快照组包含的快照ID列表

	SnapshotIdSet []*string `json:"SnapshotIdSet,omitempty" name:"SnapshotIdSet"`
	// 快照组状态。<br><li>NORMAL: 正常<br><li>CREATING:创建中<br><li>DELETED:已删除<br><li>FAILED:创建失败<br><li>DISMISS:已解散<br><li>ROLLBACKING:回滚中

	SnapshotGroupState *string `json:"SnapshotGroupState,omitempty" name:"SnapshotGroupState"`
	// 快照组创建进度

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 快照组关联的镜像列表

	Images []*Image `json:"Images,omitempty" name:"Images"`
	// 快照组名称

	SnapshotGroupName *string `json:"SnapshotGroupName,omitempty" name:"SnapshotGroupName"`
	// 快照组关联的镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 最后一次执行的异步任务执行结果

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 最后一次执行的异步任务对应的API RequestId

	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	// 最后一次执行的异步任务操作

	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`
	// 最后一次异步任务操作失败的资源列表

	OperationFailedResourceIdSet []*string `json:"OperationFailedResourceIdSet,omitempty" name:"OperationFailedResourceIdSet"`
	// 最后一次异步操作成功的资源列表

	OperationSuccessResourceIdSet []*string `json:"OperationSuccessResourceIdSet,omitempty" name:"OperationSuccessResourceIdSet"`
	// 最后一次异步操作结果的提示码

	ErrorPrompt *string `json:"ErrorPrompt,omitempty" name:"ErrorPrompt"`
}

type ApplySnapshotGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplySnapshotGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest

	// 要删除的定期快照策略ID列表。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
}

func (r *DeleteAutoSnapshotPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要查询的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest

	// 快照ID, 可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 新的快照名称。最长为60个字符。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照的保留时间，FALSE表示非永久保留，TRUE表示永久保留。仅支持将非永久快照修改为永久快照。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 新的快照项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifySnapshotAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotCopyResult struct {

	// 跨地复制的目标地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 复制到目标地域的新快照ID。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type DescribeDiskAssociatedSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘关联的快照列表

		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
		// 云盘关联的快照总大小

		TotalSnapSize *uint64 `json:"TotalSnapSize,omitempty" name:"TotalSnapSize"`
		// 云盘关联的快照总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskAssociatedSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的云硬盘ID。如果传入多个云盘ID，仅支持所有云盘修改为同一属性。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 新的云硬盘项目ID，只支持修改弹性云盘的项目ID。通过[DescribeProject](/document/api/378/4400)接口查询可用项目及其ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 新的云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 是否为弹性云盘，FALSE表示非弹性云盘，TRUE表示弹性云盘。仅支持非弹性云盘修改为弹性云盘。

	Portable *bool `json:"Portable,omitempty" name:"Portable"`
	// 成功挂载到云主机后该云硬盘是否随云主机销毁，TRUE表示随云主机销毁，FALSE表示不随云主机销毁。仅支持按量计费云硬盘数据盘。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 变更云盘类型时，可传入该参数，表示变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。<br>当前不支持批量变更类型，即传入DiskType时，DiskIds仅支持传入一块云盘；<br>变更云盘类型时不支持同时变更其他属性。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

func (r *ModifyDiskAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskChargePrepaid struct {

	// 购买云盘的时长，默认单位为月，此时，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 需要将云盘的到期时间与挂载的子机对齐时，可传入该参数。该参数表示子机当前的到期时间，此时Period如果传入，则表示子机需要续费的时长，云盘会自动按对齐到子机续费后的到期时间续费。

	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest

	// 需要创建快照的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 快照名称，不传则新快照名称默认为“未命名”。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照的项目ID，不传则默认与云盘项目ID保持一致。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 快照绑定的标签列表，不传则默认与云盘标签保持一致。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// TCE环境的项目ID，与公有云的ProjectId不相关，不传则与云盘项目保持一致。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStoragesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的块存储数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 块存储详细信息列表。

		DiskSet []*BlockStorage `json:"DiskSet,omitempty" name:"DiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockStoragesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockStoragesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedAction struct {

	// 不能操作接口名，比如查询云盘列表"DescribeDisks"，对于多用途的接口，会以"接口.入参"的形式返回。比如"ModifyDiskAttributes.DiskType"

	Action *string `json:"Action,omitempty" name:"Action"`
	// 接口不能操作的原因

	Message *string `json:"Message,omitempty" name:"Message"`
	// 接口不能操作对应提示的错误码

	Code *string `json:"Code,omitempty" name:"Code"`
}

type Disk struct {

	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。

	Portable *bool `json:"Portable,omitempty" name:"Portable"`
	// 云硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云盘是否具备创建快照的能力。取值范围：<br><li>false表示不具备<br><li>true表示具备。

	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`
	// 云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云硬盘大小。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载<br><li>ATTACHING：挂载中<br><li>ATTACHED：已挂载<br><li>DETACHING：解挂中<br><li>EXPANDING：扩容中<br><li>ROLLBACKING：回滚中。

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘是否挂载到云主机上。取值范围：<br><li>false:表示未挂载<br><li>true:表示已挂载。

	Attached *bool `json:"Attached,omitempty" name:"Attached"`
	// 云硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云硬盘的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云硬盘的到期时间。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云盘是否处于快照回滚状态。取值范围：<br><li>false:表示不处于快照回滚状态<br><li>true:表示处于快照回滚状态。

	Rollbacking *bool `json:"Rollbacking,omitempty" name:"Rollbacking"`
	// 云盘快照回滚的进度。

	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`
	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br><li>true:表示加密盘。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置<br><li>false：云盘自动续费标识正常。

	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。<br><li>false：云盘到期时间晚于实例。

	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`
	// 判断预付费的云盘是否支持主动退还。<br><li>true:支持主动退还<br><li>false:不支持主动退还。

	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`
	// 预付费云盘在不支持主动退还的情况下，该参数表明不支持主动退还的具体原因。取值范围：<br><li>1：云硬盘已经退还<br><li>2：云硬盘已过期<br><li>3：云盘不支持退还<br><li>8：超过可退还数量的限制。

	ReturnFailCode *uint64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`
	// 云盘关联的定期快照ID。只有在调用DescribeDisks接口时，入参ReturnBindAutoSnapshotPolicy取值为TRUE才会返回该参数。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 与云盘绑定的标签，云盘未绑定标签则取值为空。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 云盘是否与挂载的实例一起销毁。<br><li>true:销毁实例时会同时销毁云盘，只支持按小时后付费云盘。<br><li>false：销毁实例时不销毁云盘。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 当前时间距离盘到期的天数（仅对预付费盘有意义）。

	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`
	// 云盘是否处于类型变更中。取值范围：<br><li>false:表示云盘不处于类型变更中<br><li>true:表示云盘已发起类型变更，正处于迁移中。

	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`
	// 云盘类型变更的迁移进度，取值0到100。

	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`
	// 云盘是否为共享型云盘。

	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`
	// 对于非共享型云盘，该参数为空数组。对于共享型云盘，则表示该云盘当前被挂载到的CVM实例InstanceId

	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
	// 云盘拥有的快照总数。

	SnapshotCount *int64 `json:"SnapshotCount,omitempty" name:"SnapshotCount"`
	// 云盘拥有的快照总容量，单位为MB。

	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// 项目ID

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 云硬盘因欠费销毁或者到期销毁时， 是否使用快照备份数据的标识。true表示销毁时创建快照进行数据备份。false表示直接销毁，不进行数据备份。

	BackupDisk *bool `json:"BackupDisk,omitempty" name:"BackupDisk"`
	// 盘挂载的子机机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 盘所属的资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 最后一次执行的异步任务对应的API RequestId

	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	// 云硬盘额外性能值，单位MB/s。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 最后一次执行的异步任务操作。

	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`
	// 云盘最后挂载的云服务器ID

	LastAttachInsId *string `json:"LastAttachInsId,omitempty" name:"LastAttachInsId"`
	// 最后一次执行的异步任务执行结果。

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 销毁云盘时删除关联的非永久保留快照。0 表示非永久快照不随云盘销毁而销毁，1表示非永久快照随云盘销毁而销毁，默认取0。快照是否永久保留可以通过DescribeSnapshots接口返回的快照详情的IsPermanent字段来判断，true表示永久快照，false表示非永久快照。

	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`
	// 出错时返回给用户的操作提示

	ErrorPrompt *string `json:"ErrorPrompt,omitempty" name:"ErrorPrompt"`
	// 加密算法类型。SM4:国密；AES256:非国密

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
}

type ModifySnapshotsSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 接收分享快照的账号Id列表，array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/568)。帐号ID不同于QQ号，查询用户帐号ID请查看[帐号信息](https://console.cloud.tencent.com/developer)中的帐号ID栏。

	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
	// 操作，包括 SHARE，CANCEL。其中SHARE代表分享操作，CANCEL代表取消分享操作。

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 快照ID, 可通过[DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647)查询获取。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *ModifySnapshotsSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotsSharePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoSnapshotPolicy struct {

	// 定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 定期快照策略名称。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 定期快照策略的状态。取值范围：<br><li>NORMAL：正常<br><li>ISOLATED：已隔离。

	AutoSnapshotPolicyState *string `json:"AutoSnapshotPolicyState,omitempty" name:"AutoSnapshotPolicyState"`
	// 定期快照策略是否激活。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 使用该定期快照策略创建出来的快照是否永久保留。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 使用该定期快照策略创建出来的快照保留天数。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 定期快照策略的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 定期快照下次触发的时间。

	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 已绑定当前定期快照策略的云盘ID列表。

	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
	// 是否跨账号复制快照

	IsCopyToRemote *uint64 `json:"IsCopyToRemote,omitempty" name:"IsCopyToRemote"`
	// 源账号uin

	CopyFromAccountUin *string `json:"CopyFromAccountUin,omitempty" name:"CopyFromAccountUin"`
	// 目的账号uin

	CopyToAccountUin *string `json:"CopyToAccountUin,omitempty" name:"CopyToAccountUin"`
}

type DiskStoragePoolGroupInfo struct {

	// 存储资源组名称

	DiskStoragePoolGroupName *string `json:"DiskStoragePoolGroupName,omitempty" name:"DiskStoragePoolGroupName"`
	// 绑定类型

	PoolBoundType *string `json:"PoolBoundType,omitempty" name:"PoolBoundType"`
	// 存储资源池组

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 云硬盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 是否售罄

	Available *bool `json:"Available,omitempty" name:"Available"`
	// 售罄信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 描述资源池的cpu信息。

	PoolCustomAttributes *PoolCustomAttributes `json:"PoolCustomAttributes,omitempty" name:"PoolCustomAttributes"`
}

type DescribeSnapshotsDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询快照的操作掩码集合

		SnapshotDeniedActionSet []*SnapshotDeniedAction `json:"SnapshotDeniedActionSet,omitempty" name:"SnapshotDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskConfigSet struct {

	// 系统盘或数据盘

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// CLOUD_BASIC、CLOUD_PREMIUM、CLOUD_SSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type SnapshotDeniedAction struct {

	// 根据快照ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 具体的快照的操作掩码信息

	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的快照ID。

		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeDiskRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘扩容后的大小，单位为GB，不得小于当前云硬盘大小。取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB，步长均为10GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘所属项目ID。 如传入则仅用于鉴权。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetail struct {

	// 购买或续费云盘的时长。

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 产品ID。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 云盘类型，系统盘或数据盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 产品信息描述。

	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 云盘介质类型。

	MediumType *string `json:"MediumType,omitempty" name:"MediumType"`
	// 时长“TimeSpan”的单位。

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 云盘大小。

	CbsSize *uint64 `json:"CbsSize,omitempty" name:"CbsSize"`
}

type Image struct {

	// 镜像实例ID。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

type ResizeGoodsDetail struct {

	// 产品ID。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 云盘ID。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 云盘当前到期时间。

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 产品信息描述。

	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 云盘扩容后的配置。

	NewConfig *DiskResizeConfig `json:"NewConfig,omitempty" name:"NewConfig"`
	// 云盘扩容前的配置。

	OldConfig *DiskResizeConfig `json:"OldConfig,omitempty" name:"OldConfig"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 付费模式，目前只有预付费，即只能取值为PREPAID。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长、是否设置自动续费等属性，创建预付费云盘该参数必传。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB。步长均为10GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘绑定的标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。

	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`
	// 定期快照策略ID。传入该参数时，云硬盘创建成功后将会自动绑定该定期快照策略

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 存储资源池组

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 项目ID

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *CreateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 快照组ID列表

	SnapshotGroupIds []*string `json:"SnapshotGroupIds,omitempty" name:"SnapshotGroupIds"`
}

func (r *DescribeSnapshotGroupsDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewDiskRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的续费时长。在云盘与挂载的实例一起续费的场景下，可以指定参数CurInstanceDeadline，此时云盘会按对齐到实例续费后的到期时间来续费。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *RenewDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskConfig struct {

	// 配置是否可用。

	Available *bool `json:"Available,omitempty" name:"Available"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘类型。取值范围：<br><li>SYSTEM_DISK：表示系统盘<br><li>DATA_DISK：表示数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 最大可配置云盘大小。

	MaxDiskSize *uint64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`
	// 最小可配置云盘大小。

	MinDiskSize *uint64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`
	// 所在[可用区](/document/api/213/9452#zone)。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 实例机型系列。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 步长

	StepSize *uint64 `json:"StepSize,omitempty" name:"StepSize"`
	// 额外的性能区间。

	ExtraPerformanceRange []*int64 `json:"ExtraPerformanceRange,omitempty" name:"ExtraPerformanceRange"`
}

type ApplyDisk struct {

	// 快照ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照对应的原云硬盘ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type Price struct {

	// 预支费用的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 后付费云盘折扣单价，单位：元。

	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`
	// 后付费云盘原单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 高精度后付费云盘原单价, 单位：元

	UnitPriceHigh *string `json:"UnitPriceHigh,omitempty" name:"UnitPriceHigh"`
	// 高精度预付费云盘预支费用的原价, 单位：元	。

	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitempty" name:"OriginalPriceHigh"`
	// 高精度预付费云盘预支费用的折扣价, 单位：元

	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitempty" name:"DiscountPriceHigh"`
	// 高精度后付费云盘折扣单价, 单位：元

	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitempty" name:"UnitPriceDiscountHigh"`
	// 后付费云盘的计价单元，取值范围：<br><li>HOUR：表示后付费云盘的计价单元是按小时计算。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

type SharePermission struct {

	// 快照分享的时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 分享的账号Id

	AccountId *int64 `json:"AccountId,omitempty" name:"AccountId"`
}

type DescribeDiskStoragePoolGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 限制单次查询的最大数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskStoragePoolGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStoragePoolGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeDiskRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 扩容后的磁盘大小。必须大于当前值，最大值为4000G，步长为10G。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *SwitchParameterResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest

	// 要查询的定期快照策略ID列表。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 过滤条件。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。<br><li>auto-snapshot-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期快照策略ID进行过滤。定期快照策略ID形如：`asp-11112222`。<br><li>auto-snapshot-policy-state - Array of String - 是否必填：否 -（过滤条件）按定期快照策略的状态进行过滤。定期快照策略ID形如：`asp-11112222`。(NORMAL：正常 | ISOLATED：已隔离。)<br><li>auto-snapshot-policy-name - Array of String - 是否必填：否 -（过滤条件）按定期快照策略名称进行过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介]中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介]中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 输出定期快照列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 定期快照列表排序的依据字段。取值范围：<br><li>CREATETIME：依据定期快照的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 传入此参数时则会尝试创建默认的定期快照策略。如果用户已有定期快照策略，则不会创建

	TryCreateDefaultAutoSnapshotPolicy *bool `json:"TryCreateDefaultAutoSnapshotPolicy,omitempty" name:"TryCreateDefaultAutoSnapshotPolicy"`
}

func (r *DescribeAutoSnapshotPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest

	// 需退还的云盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoSnapshotPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotsSharePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotsSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotsSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoSnapshotPolicyAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的定期快照策略ID。

		AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
		// 首次开始备份的时间。

		NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotGroupRequest struct {
	*tchttp.BaseRequest

	// 需要创建快照组的云盘ID列表

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 快照组名称

	SnapshotGroupName *string `json:"SnapshotGroupName,omitempty" name:"SnapshotGroupName"`
}

func (r *CreateSnapshotGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDiskResourcesRequest struct {
	*tchttp.BaseRequest

	// 指定要查询的地域，支持同时传入多个地域。不传该参数，则默认查询所有地域。

	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeUserDiskResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDiskResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>snapshot-group-id - Array of String - 是否必填：否 -（过滤条件）按快照组ID过滤 <br><li>snapshot-group-state - Array of String - 是否必填：否 -（过滤条件）按快照组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) <br><li>snapshot-group-name - Array of String - 是否必填：否 -（过滤条件）按快照组名称过滤 <br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按快照组内的快照ID过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotSharePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照的分享信息的集合

		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 最大可共享的数量

		MaxShareAccount *uint64 `json:"MaxShareAccount,omitempty" name:"MaxShareAccount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisksRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的云硬盘ID。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云盘的续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyDisksRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisksRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest

	// 将要解挂的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可解挂10块弹性云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 对于非共享型云盘，会忽略该参数；对于共享型云盘，该参数表示要从哪个CVM实例上解挂云盘。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest

	// 额外购买的云硬盘性能值，单位MB/s。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *InquirePriceModifyDiskExtraPerformanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 创建快照的云硬盘大小，单位为GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 指定快照创建的时长。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 创建快照的数量。默认取值为1。

	SnapshotCount *uint64 `json:"SnapshotCount,omitempty" name:"SnapshotCount"`
}

func (r *InquiryPriceCreateSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterResizeDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费订单参数。

		DiskOrder []*DiskOrder `json:"DiskOrder,omitempty" name:"DiskOrder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest

	// 云服务器实例ID，通过[DescribeInstances](/document/product/213/15728)接口查询。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDiskNumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDiskNumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneDiskConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可创建云盘数量

		DiskConfigSetQuota *int64 `json:"DiskConfigSetQuota,omitempty" name:"DiskConfigSetQuota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneDiskConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneDiskConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSnapOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSnapOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSnapOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapOverview struct {

	// 用户已创建快照总容量。

	SnapshotCapacityUsedTotal *float64 `json:"SnapshotCapacityUsedTotal,omitempty" name:"SnapshotCapacityUsedTotal"`
	// 用户已创建快照总个数。

	SnapshotNumberUsedTotal *int64 `json:"SnapshotNumberUsedTotal,omitempty" name:"SnapshotNumberUsedTotal"`
}

type DescribeDiskStoragePoolGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源池列表详情。

		DiskStoragePoolGroupSet []*DiskStoragePoolGroupInfo `json:"DiskStoragePoolGroupSet,omitempty" name:"DiskStoragePoolGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStoragePoolGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStoragePoolGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询快照组的操作掩码集合

		SnapshotGroupDeniedActionSet []*SnapshotGroupDeniedAction `json:"SnapshotGroupDeniedActionSet,omitempty" name:"SnapshotGroupDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotGroupsDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了新购云盘的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了变更云盘类型的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyDiskAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建快照的价格。

		SnapshotPrice *Price `json:"SnapshotPrice,omitempty" name:"SnapshotPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了扩容云盘的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisksRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisksRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisksRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 是否同时删除快照关联的镜像；取值为false，表示不删除快照绑定的镜像，此时，如果快照有绑定的镜像，删除会失败；取值为true，表示同时删除快照绑定的镜像；默认值为false

	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {

	// 属性名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InquiryPriceCreateDisksRequest struct {
	*tchttp.BaseRequest

	// 云硬盘类型。取值范围：<br><li>普通云硬盘：CLOUD_BASIC<br><li>高性能云硬盘：CLOUD_PREMIUM<br><li>SSD云硬盘：CLOUD_SSD。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘大小，取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB，步长均为10GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 付费模式，目前只有预付费，即只能取值为PREPAID。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 预付费相关参数设置，通过该参数可以指定包年包月云盘的购买时长，预付费云盘该参数必传。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 购买云盘的数量。不填则默认为1。

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 云盘所属项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 额外购买的性能大小。

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *InquiryPriceCreateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskExtraPerformanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照组ID

		SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneDiskConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 付费类型。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 云盘配置列表。

	DiskConfigs []*DiskConfigSet `json:"DiskConfigs,omitempty" name:"DiskConfigs"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeZoneDiskConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneDiskConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockStorage struct {

	// 块存储ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 块存储名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 块存储大小，单位GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 块存储介质类型。取值范围：<br><li>CLOUD_BASIC<br><li>CLOUD_PREMIUM<br><li>CLOUD_SSD<br><li>CLOUD_ENHANCEDSSD<br><li>LOCAL_BASIC<br><li>LOCAL_SSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 块存储类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 块存储挂载的实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 块存储挂载的实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type DescribeDisksDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *DescribeDisksDeniedActionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksDeniedActionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksDeniedActionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询的云硬盘的操作掩码的集合

		DiskDeniedActionSet []*DiskDeniedAction `json:"DiskDeniedActionSet,omitempty" name:"DiskDeniedActionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksDeniedActionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照的数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 快照的详情列表。

		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskResizeConfig struct {

	// 产品ID。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 云盘大小。

	CbsSize *uint64 `json:"CbsSize,omitempty" name:"CbsSize"`
	// 云盘介质类型。

	MediumType *string `json:"MediumType,omitempty" name:"MediumType"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可选参数，不传该参数则仅执行挂载操作。<br>传入该参数，会在挂载时将云盘的生命周期与待挂载主机对齐。取值范围：<br><li>AUTO_RENEW：云盘未设置自动续费时，可传该值，将云盘设置为自动续费。<br><li> DEADLINE_ALIGN：当云盘的到期时间早于待挂载主机，可传该值，将云盘的到期时间与主机对齐。

	AlignType *string `json:"AlignType,omitempty" name:"AlignType"`
	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 可选参数，用于控制台批量挂载共享盘到多个CVM时指定实例ID。如果传入多个InstanceId，那么DisksId中指定的云盘必须全部为共享云盘，否则返回错误。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AttachDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskOverview struct {

	// 用户云硬盘总数。

	DiskNumberTotal *int64 `json:"DiskNumberTotal,omitempty" name:"DiskNumberTotal"`
	// 用户已过期云硬盘总数。

	ExpiredNumberTotal *int64 `json:"ExpiredNumberTotal,omitempty" name:"ExpiredNumberTotal"`
	// 用户7天内将到期云硬盘总数

	DiskNumberExpireIn7days *int64 `json:"DiskNumberExpireIn7days,omitempty" name:"DiskNumberExpireIn7days"`
}

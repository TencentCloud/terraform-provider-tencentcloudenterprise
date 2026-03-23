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

package v20201214

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal


type InquirePriceCreateDedicatedClusterUserDefinedOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价信息

		PriceInfoSet []*PriceInfo `json:"PriceInfoSet,omitempty" name:"PriceInfoSet"`
		// 总的原价

		OriginalTotalPrice *float64 `json:"OriginalTotalPrice,omitempty" name:"OriginalTotalPrice"`
		// 总的折扣价

		DiscountTotalPrice *float64 `json:"DiscountTotalPrice,omitempty" name:"DiscountTotalPrice"`
		// 总的折扣百分比

		DiscountTotal *float64 `json:"DiscountTotal,omitempty" name:"DiscountTotal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDedicatedClusterUserDefinedOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDedicatedClusterUserDefinedOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupSsdInfo struct {

	// 硬盘大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type SwitchAggregateInfo struct {

	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type CreateDedicatedClusterOrderRequest struct {
	*tchttp.BaseRequest

	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// order关联的专用集群类型数组

	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitempty" name:"DedicatedClusterTypes"`
	// order关联的cos存储信息

	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
	// order关联的cbs存储信息

	CbsInfo *CbsInfo `json:"CbsInfo,omitempty" name:"CbsInfo"`
	// 购买来源，默认为cloudApi

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
	// 当调用API接口提交订单时，需要提交DedicatedClusterOrderId

	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitempty" name:"DedicatedClusterOrderId"`
}

func (r *CreateDedicatedClusterOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataChannelInfo struct {

	// 数据通道类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DedicatedClusterOrderItem struct {

	// 专用集群类型id

	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitempty" name:"DedicatedClusterTypeId"`
	// 支持的存储类型列表

	SupportedStorageType []*string `json:"SupportedStorageType,omitempty" name:"SupportedStorageType"`
	// 支持的上连交换机的链路传输速率(GiB)

	SupportedUplinkSpeed []*int64 `json:"SupportedUplinkSpeed,omitempty" name:"SupportedUplinkSpeed"`
	// 支持的实例族列表

	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitempty" name:"SupportedInstanceFamily"`
	// 地板承重要求(KG)

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 功率要求(KW)

	PowerDraw *float64 `json:"PowerDraw,omitempty" name:"PowerDraw"`
	// 订单状态

	SubOrderStatus *string `json:"SubOrderStatus,omitempty" name:"SubOrderStatus"`
	// 订单创建的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 子订单ID

	SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
	// 关联的集群规格数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 规格简单描述

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规格详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CPU数

	TotalCpu *int64 `json:"TotalCpu,omitempty" name:"TotalCpu"`
	// 内存数

	TotalMem *int64 `json:"TotalMem,omitempty" name:"TotalMem"`
	// GPU数

	TotalGpu *int64 `json:"TotalGpu,omitempty" name:"TotalGpu"`
	// 规格英文名

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 规格展示

	ComputeFormat *string `json:"ComputeFormat,omitempty" name:"ComputeFormat"`
	// 规格类型

	TypeFamily *string `json:"TypeFamily,omitempty" name:"TypeFamily"`
	// 0未支付，1已支付

	SubOrderPayStatus *int64 `json:"SubOrderPayStatus,omitempty" name:"SubOrderPayStatus"`
}

type DescribeDedicatedClusterCosCapacityRequest struct {
	*tchttp.BaseRequest

	// 查询的专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterCosCapacityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterCosCapacityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterOrdersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用集群订单列表

		DedicatedClusterOrderSet []*DedicatedClusterOrder `json:"DedicatedClusterOrderSet,omitempty" name:"DedicatedClusterOrderSet"`
		// 符合条件的专用集群订单总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterOrdersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DedicatedClusterOrder struct {

	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 地板承重要求(KG)

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 功率要求(KW)

	PowerDraw *float64 `json:"PowerDraw,omitempty" name:"PowerDraw"`
	// 订单状态

	OrderStatus *string `json:"OrderStatus,omitempty" name:"OrderStatus"`
	// 订单创建的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 大订单ID

	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitempty" name:"DedicatedClusterOrderId"`
	// 订单类型，创建CREATE或扩容EXTEND

	Action *string `json:"Action,omitempty" name:"Action"`
	// 子订单详情列表

	DedicatedClusterOrderItems []*DedicatedClusterOrderItem `json:"DedicatedClusterOrderItems,omitempty" name:"DedicatedClusterOrderItems"`
	// cpu值

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// mem值

	Mem *int64 `json:"Mem,omitempty" name:"Mem"`
	// gpu值

	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`
	// 0代表未支付，1代表已支付

	PayStatus *int64 `json:"PayStatus,omitempty" name:"PayStatus"`
	// 支付方式，一次性、按月、按年

	PayType *string `json:"PayType,omitempty" name:"PayType"`
	// 购买时长的单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 购买时长

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 订单类型

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
	// 专用集群类型ID

	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitempty" name:"DedicatedClusterTypeId"`
	// 支持的存储类型

	SupportedStorageType []*string `json:"SupportedStorageType,omitempty" name:"SupportedStorageType"`
	// 支持的上行速率

	SupportedUplinkSpeed *int64 `json:"SupportedUplinkSpeed,omitempty" name:"SupportedUplinkSpeed"`
	// 支持的实例系列

	SupportedInstanceFamily *string `json:"SupportedInstanceFamily,omitempty" name:"SupportedInstanceFamily"`
	// 检查状态

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 预计交货时间

	DeliverExpectTime *string `json:"DeliverExpectTime,omitempty" name:"DeliverExpectTime"`
	// 交付完成时间

	DeliverFinishTime *string `json:"DeliverFinishTime,omitempty" name:"DeliverFinishTime"`
	// 检查预计时间

	CheckExpectTime *string `json:"CheckExpectTime,omitempty" name:"CheckExpectTime"`
	// 检查完成时间

	CheckFinishTime *string `json:"CheckFinishTime,omitempty" name:"CheckFinishTime"`
	// 订单SLA

	OrderSLA *string `json:"OrderSLA,omitempty" name:"OrderSLA"`
	// 订单支付计划

	OrderPayPlan *string `json:"OrderPayPlan,omitempty" name:"OrderPayPlan"`
}

type SiteDetail struct {

	// 站点id

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 站点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 站点创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 光纤类型

	FiberType *string `json:"FiberType,omitempty" name:"FiberType"`
	// 网络到云平台Region区域的上行链路速度

	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitempty" name:"UplinkSpeedGbps"`
	// 将CDC连接到网络时，每台CDC网络设备(每个机架&nbsp;2&nbsp;台设备)使用的上行链路数量。

	UplinkCount *int64 `json:"UplinkCount,omitempty" name:"UplinkCount"`
	// 将CDC连接到网络时采用的光学标准

	OpticalStandard *string `json:"OpticalStandard,omitempty" name:"OpticalStandard"`
	// 是否提供冗余的上游设备(交换机或路由器)，以便两台&nbsp;网络设备都能连接到网络设备。

	RedundantNetworking *bool `json:"RedundantNetworking,omitempty" name:"RedundantNetworking"`
	// 电源连接器类型

	PowerConnectors *string `json:"PowerConnectors,omitempty" name:"PowerConnectors"`
	// 从机架上方还是下方供电。

	PowerFeedDrop *string `json:"PowerFeedDrop,omitempty" name:"PowerFeedDrop"`
	// 功耗(KW)

	PowerDrawKva *float64 `json:"PowerDrawKva,omitempty" name:"PowerDrawKva"`
	// 是否满足下面环境条件：&nbsp;1、场地没有材料要求或验收标准会影响&nbsp;CDC&nbsp;设备配送和安装。&nbsp;2、确定的机架位置包含:&nbsp;温度范围为&nbsp;41&nbsp;到&nbsp;104°F&nbsp;(5&nbsp;到&nbsp;40°C)。&nbsp;湿度范围为&nbsp;10°F&nbsp;(-12°C)和&nbsp;8%&nbsp;RH&nbsp;(相对湿度)到&nbsp;70°F(21°C)和&nbsp;80%&nbsp;RH。&nbsp;机架位置的气流方向为从前向后，且应具有足够的&nbsp;CFM&nbsp;(每分钟立方英尺)。CFM&nbsp;必须是&nbsp;CDC&nbsp;配置的&nbsp;kVA&nbsp;功耗值的&nbsp;145.8&nbsp;倍。

	ConditionRequirement *bool `json:"ConditionRequirement,omitempty" name:"ConditionRequirement"`
	// 是否满足下面的尺寸条件：&nbsp;您的装货站台可以容纳一个机架箱(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;94"&nbsp;x&nbsp;54"&nbsp;x&nbsp;48")。&nbsp;您可以提供从机架(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;80"&nbsp;x&nbsp;24"&nbsp;x&nbsp;48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。&nbsp;在最终的&nbsp;CDC安置位置，前部间隙可以为&nbsp;48"&nbsp;或更大，后部间隙可以为&nbsp;24"&nbsp;或更大。

	DimensionRequirement *bool `json:"DimensionRequirement,omitempty" name:"DimensionRequirement"`
	// 最大承重(KG)

	MaxWeight *int64 `json:"MaxWeight,omitempty" name:"MaxWeight"`
	// 站点地址

	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`
	// 站点所在地区的详细地址信息（补充）

	OptionalAddressLine *string `json:"OptionalAddressLine,omitempty" name:"OptionalAddressLine"`
	// 是否需要云平台团队协助完成机架支撑工作

	NeedHelp *bool `json:"NeedHelp,omitempty" name:"NeedHelp"`
	// 上游断路器是否具备

	BreakerRequirement *bool `json:"BreakerRequirement,omitempty" name:"BreakerRequirement"`
	// 是否电源冗余

	RedundantPower *bool `json:"RedundantPower,omitempty" name:"RedundantPower"`
	// 站点所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 站点所在省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 站点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 站点所在地区的邮编

	PostalCode *int64 `json:"PostalCode,omitempty" name:"PostalCode"`
}

type DescribeDedicatedClusterHostStatisticsRequest struct {
	*tchttp.BaseRequest

	// 查询的专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterHostStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterOverviewRequest struct {
	*tchttp.BaseRequest

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosCapacity struct {

	// 已购cos的总容量大小，单位GB

	TotalCapacity *float64 `json:"TotalCapacity,omitempty" name:"TotalCapacity"`
	// 剩余可用cos的容量大小，单位GB

	TotalFreeCapacity *float64 `json:"TotalFreeCapacity,omitempty" name:"TotalFreeCapacity"`
	// 已用cos的容量大小，单位GB

	TotalUsedCapacity *float64 `json:"TotalUsedCapacity,omitempty" name:"TotalUsedCapacity"`
}

type CreateDedicatedClusterUserDefinedOrderRequest struct {
	*tchttp.BaseRequest

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 基础设施信息

	BaseEquipmentInfo *BaseEquipmentInfo `json:"BaseEquipmentInfo,omitempty" name:"BaseEquipmentInfo"`
	// 服务器信息

	CvmInfoSet []*CvmInfo `json:"CvmInfoSet,omitempty" name:"CvmInfoSet"`
	// 存储信息

	StorageInfo *StorageInfo `json:"StorageInfo,omitempty" name:"StorageInfo"`
	// 网络信息

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
	// 购买时间单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 购买时间数量

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 中间件信息

	MiddleWareInfo *MiddleWareInfo `json:"MiddleWareInfo,omitempty" name:"MiddleWareInfo"`
	// 数据库信息

	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// 支付方式，一次性、按月、按年

	PayType *string `json:"PayType,omitempty" name:"PayType"`
	// 预付费信息

	PrepaidInfo *PrepaidInfo `json:"PrepaidInfo,omitempty" name:"PrepaidInfo"`
}

func (r *CreateDedicatedClusterUserDefinedOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterUserDefinedOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterAlarmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警信息列表

		HostInfoSet []*AlarmInfo `json:"HostInfoSet,omitempty" name:"HostInfoSet"`
		// 告警总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterAlarmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TdsqlInfo struct {

	// 类型规格

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type CreateDedicatedClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的专用集群id

		DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDedicatedClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterInstanceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的实例规格列表

		DedicatedClusterInstanceTypeSet []*DedicatedClusterInstanceType `json:"DedicatedClusterInstanceTypeSet,omitempty" name:"DedicatedClusterInstanceTypeSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterInstanceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {

	// 订单原价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 订单折扣后价格

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 订单折扣

	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`
}

type DescribeDedicatedClusterHostsInfoRequest struct {
	*tchttp.BaseRequest

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClusterHostsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDedicatedClusterInfoRequest struct {
	*tchttp.BaseRequest

	// 本地专用集群ID

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 集群的新名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群的新可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 集群的新描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 集群所在站点

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
}

func (r *ModifyDedicatedClusterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDedicatedClusterInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostStatistic struct {

	// 宿主机规格

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 宿主机机型系列

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 宿主机的CPU核数，单位：核

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 宿主机内存大小，单位：GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 该规格宿主机的数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type Site struct {

	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 站点id

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 站点描述&nbsp;注意：此字段可能返回&nbsp;null，表示取不到有效值。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 站点创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateSiteRequest struct {
	*tchttp.BaseRequest

	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 站点所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 站点所在省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 站点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 站点所在地区的详细地址信息

	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`
	// 站点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 注意事项

	Note *string `json:"Note,omitempty" name:"Note"`
	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项

	FiberType *string `json:"FiberType,omitempty" name:"FiberType"`
	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离

	OpticalStandard *string `json:"OpticalStandard,omitempty" name:"OpticalStandard"`
	// 电源连接器类型

	PowerConnectors *string `json:"PowerConnectors,omitempty" name:"PowerConnectors"`
	// 从机架上方还是下方供电。

	PowerFeedDrop *string `json:"PowerFeedDrop,omitempty" name:"PowerFeedDrop"`
	// 最大承重(KG)

	MaxWeight *int64 `json:"MaxWeight,omitempty" name:"MaxWeight"`
	// 功耗(KW)

	PowerDrawKva *int64 `json:"PowerDrawKva,omitempty" name:"PowerDrawKva"`
	// 网络到云平台Region区域的上行链路速度

	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitempty" name:"UplinkSpeedGbps"`
	// 将CDC连接到网络时，每台CDC网络设备(每个机架&nbsp;2&nbsp;台设备)使用的上行链路数量。

	UplinkCount *int64 `json:"UplinkCount,omitempty" name:"UplinkCount"`
	// 是否满足下面环境条件：&nbsp;1、场地没有材料要求或验收标准会影响&nbsp;CDC&nbsp;设备配送和安装。&nbsp;2、确定的机架位置包含:&nbsp;温度范围为&nbsp;41&nbsp;到&nbsp;104°F&nbsp;(5&nbsp;到&nbsp;40°C)。&nbsp;湿度范围为&nbsp;10°F&nbsp;(-12°C)和&nbsp;8%&nbsp;RH&nbsp;(相对湿度)到&nbsp;70°F(21°C)和&nbsp;80%&nbsp;RH。&nbsp;机架位置的气流方向为从前向后，且应具有足够的&nbsp;CFM&nbsp;(每分钟立方英尺)。CFM&nbsp;必须是&nbsp;CDC&nbsp;配置的&nbsp;kVA&nbsp;功耗值的&nbsp;145.8&nbsp;倍。

	ConditionRequirement *bool `json:"ConditionRequirement,omitempty" name:"ConditionRequirement"`
	// 是否满足下面的尺寸条件：&nbsp;您的装货站台可以容纳一个机架箱(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;94"&nbsp;x&nbsp;54"&nbsp;x&nbsp;48")。&nbsp;您可以提供从机架(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;80"&nbsp;x&nbsp;24"&nbsp;x&nbsp;48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。&nbsp;在最终的&nbsp;CDC安置位置，前部间隙可以为&nbsp;48"&nbsp;或更大，后部间隙可以为&nbsp;24"&nbsp;或更大。

	DimensionRequirement *bool `json:"DimensionRequirement,omitempty" name:"DimensionRequirement"`
	// 是否提供冗余的上游设备(交换机或路由器)，以便两台&nbsp;网络设备都能连接到网络设备。

	RedundantNetworking *bool `json:"RedundantNetworking,omitempty" name:"RedundantNetworking"`
	// 站点所在地区的邮编

	PostalCode *int64 `json:"PostalCode,omitempty" name:"PostalCode"`
	// 站点所在地区的详细地址信息（补充）

	OptionalAddressLine *string `json:"OptionalAddressLine,omitempty" name:"OptionalAddressLine"`
	// 是否需要云平台团队协助完成机架支撑工作

	NeedHelp *bool `json:"NeedHelp,omitempty" name:"NeedHelp"`
	// 是否电源冗余

	RedundantPower *bool `json:"RedundantPower,omitempty" name:"RedundantPower"`
	// 上游断路器是否具备

	BreakerRequirement *bool `json:"BreakerRequirement,omitempty" name:"BreakerRequirement"`
}

func (r *CreateSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDedicatedClustersRequest struct {
	*tchttp.BaseRequest

	// 要删除的专用集群id

	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitempty" name:"DedicatedClusterIds"`
}

func (r *DeleteDedicatedClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterHostsRequest struct {
	*tchttp.BaseRequest

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClusterHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySiteDeviceInfoRequest struct {
	*tchttp.BaseRequest

	// 机房ID

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 您将使用光纤类型将CDC设备连接到网络。有单模和多模两种选项。

	FiberType *string `json:"FiberType,omitempty" name:"FiberType"`
	// 您将CDC连接到网络时采用的光学标准。此字段取决于上行链路速度、光纤类型和到上游设备的距离。

	OpticalStandard *string `json:"OpticalStandard,omitempty" name:"OpticalStandard"`
	// 电源连接器类型

	PowerConnectors *string `json:"PowerConnectors,omitempty" name:"PowerConnectors"`
	// 从机架上方还是下方供电。

	PowerFeedDrop *string `json:"PowerFeedDrop,omitempty" name:"PowerFeedDrop"`
	// 最大承重(KG)

	MaxWeight *int64 `json:"MaxWeight,omitempty" name:"MaxWeight"`
	// 功耗(KW)

	PowerDrawKva *int64 `json:"PowerDrawKva,omitempty" name:"PowerDrawKva"`
	// 网络到云平台Region区域的上行链路速度

	UplinkSpeedGbps *int64 `json:"UplinkSpeedGbps,omitempty" name:"UplinkSpeedGbps"`
	// 将CDC连接到网络时，每台CDC网络设备(每个机架&nbsp;2&nbsp;台设备)使用的上行链路数量。

	UplinkCount *int64 `json:"UplinkCount,omitempty" name:"UplinkCount"`
	// 是否满足下面环境条件：&nbsp;1、场地没有材料要求或验收标准会影响&nbsp;CDC&nbsp;设备配送和安装。&nbsp;2、确定的机架位置包含:&nbsp;温度范围为&nbsp;41&nbsp;到&nbsp;104°F&nbsp;(5&nbsp;到&nbsp;40°C)。&nbsp;湿度范围为&nbsp;10°F&nbsp;(-12°C)和&nbsp;8%&nbsp;RH&nbsp;(相对湿度)到&nbsp;70°F(21°C)和&nbsp;80%&nbsp;RH。&nbsp;机架位置的气流方向为从前向后，且应具有足够的&nbsp;CFM&nbsp;(每分钟立方英尺)。CFM&nbsp;必须是&nbsp;CDC&nbsp;配置的&nbsp;kVA&nbsp;功耗值的&nbsp;145.8&nbsp;倍。

	ConditionRequirement *bool `json:"ConditionRequirement,omitempty" name:"ConditionRequirement"`
	// 是否满足下面的尺寸条件：&nbsp;您的装货站台可以容纳一个机架箱(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;94"&nbsp;x&nbsp;54"&nbsp;x&nbsp;48")。&nbsp;您可以提供从机架(高&nbsp;x&nbsp;宽&nbsp;x&nbsp;深&nbsp;=&nbsp;80"&nbsp;x&nbsp;24"&nbsp;x&nbsp;48")交货地点到机架最终安置位置的明确通道。测量深度时，应包括站台、走廊通道、门、转弯、坡道、货梯，并将其他通道限制考虑在内。&nbsp;在最终的&nbsp;CDC安置位置，前部间隙可以为&nbsp;48"&nbsp;或更大，后部间隙可以为&nbsp;24"&nbsp;或更大。

	DimensionRequirement *bool `json:"DimensionRequirement,omitempty" name:"DimensionRequirement"`
	// 是否提供冗余的上游设备(交换机或路由器)，以便两台&nbsp;网络设备都能连接到网络设备。

	RedundantNetworking *bool `json:"RedundantNetworking,omitempty" name:"RedundantNetworking"`
	// 是否需要云平台团队协助完成机架支撑工作

	NeedHelp *bool `json:"NeedHelp,omitempty" name:"NeedHelp"`
	// 是否电源冗余

	RedundantPower *bool `json:"RedundantPower,omitempty" name:"RedundantPower"`
	// 上游断路器是否具备

	BreakerRequirement *bool `json:"BreakerRequirement,omitempty" name:"BreakerRequirement"`
}

func (r *ModifySiteDeviceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteDeviceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CvmInfo struct {

	// 服务器类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 服务器数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type HostSet struct {

	// 宿主机类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 业务列表

	Business []*string `json:"Business,omitempty" name:"Business"`
	// 内网IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// cpu总数

	// cpu可用数

	// 内存总数

	// 内存可用数

}

type NetworkInfo struct {

	// 负载均衡信息

	ClbInfo *ClbInfo `json:"ClbInfo,omitempty" name:"ClbInfo"`
}

type DescribeDedicatedSupportedZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的可用区列表

		ZoneSet []*RegionZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedSupportedZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedSupportedZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDedicatedClusterUserDefinedOrderRequest struct {
	*tchttp.BaseRequest

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 基础设施信息

	BaseEquipmentInfo *BaseEquipmentInfo `json:"BaseEquipmentInfo,omitempty" name:"BaseEquipmentInfo"`
	// 服务器信息

	CvmInfoSet []*CvmInfo `json:"CvmInfoSet,omitempty" name:"CvmInfoSet"`
	// 存储信息

	StorageInfo *StorageInfo `json:"StorageInfo,omitempty" name:"StorageInfo"`
	// 网络信息

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
	// 购买时间单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 购买时间数量

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 中间件信息

	MiddleWareInfo *MiddleWareInfo `json:"MiddleWareInfo,omitempty" name:"MiddleWareInfo"`
	// 数据库信息

	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
	// 支付方式，一次性、按月、按年

	PayType *string `json:"PayType,omitempty" name:"PayType"`
	// 预付费信息

	PrepaidInfo *PrepaidInfo `json:"PrepaidInfo,omitempty" name:"PrepaidInfo"`
	// 子订单id

	DedicatedClusterSubOrderId *string `json:"DedicatedClusterSubOrderId,omitempty" name:"DedicatedClusterSubOrderId"`
}

func (r *InquirePriceCreateDedicatedClusterUserDefinedOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDedicatedClusterUserDefinedOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterHostsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 宿主机信息

		HostSet []*HostSet `json:"HostSet,omitempty" name:"HostSet"`
		// 宿主机总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterHostsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的站点列表

		SiteSet []*Site `json:"SiteSet,omitempty" name:"SiteSet"`
		// 符合条件的站点数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSitesDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的专用集群列表

		SiteDetailSet []*SiteDetail `json:"SiteDetailSet,omitempty" name:"SiteDetailSet"`
		// 符合条件的站点数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSitesDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDedicatedClusterInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDedicatedClusterInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDedicatedClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInfo struct {

	// 集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 集群名

	DedicatedClusterName *string `json:"DedicatedClusterName,omitempty" name:"DedicatedClusterName"`
	// 告警级别

	AlarmLevel *string `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 告警状态

	AlarmState *string `json:"AlarmState,omitempty" name:"AlarmState"`
	// 告警时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 告警描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 告警所属服务

	AlarmService *string `json:"AlarmService,omitempty" name:"AlarmService"`
}

type BaseEquipmentInfo struct {

	// 基础底座信息

	BaseFoundationInfo *BaseFoundationInfo `json:"BaseFoundationInfo,omitempty" name:"BaseFoundationInfo"`
	// 扩展机柜信息

	ExtendRackInfo *ExtendRackInfo `json:"ExtendRackInfo,omitempty" name:"ExtendRackInfo"`
	// 管控接入信息

	ControlAccessInfo *ControlAccessInfo `json:"ControlAccessInfo,omitempty" name:"ControlAccessInfo"`
	// 管控通道带宽信息

	ControlBandwidthInfo *ControlBandwidthInfo `json:"ControlBandwidthInfo,omitempty" name:"ControlBandwidthInfo"`
	// 管控VPN网关信息

	ControlVPNInfo *ControlVPNInfo `json:"ControlVPNInfo,omitempty" name:"ControlVPNInfo"`
	// 数据通道信息

	DataChannelInfo *DataChannelInfo `json:"DataChannelInfo,omitempty" name:"DataChannelInfo"`
	// 网络汇聚设备

	SwitchAggregateInfo *SwitchAggregateInfo `json:"SwitchAggregateInfo,omitempty" name:"SwitchAggregateInfo"`
}

type DescribeSitesDetailRequest struct {
	*tchttp.BaseRequest

	// 按照站点id过滤

	SiteIds []*string `json:"SiteIds,omitempty" name:"SiteIds"`
	// 按照站定名称模糊匹配

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSitesDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySiteDeviceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySiteDeviceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteDeviceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbInfo struct {

	// 负载均衡类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ControlBandwidthInfo struct {

	// 带宽值

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type PriceDetail struct {

	// 价格项名称

	PriceName *string `json:"PriceName,omitempty" name:"PriceName"`
	// 原价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 折扣后的价格

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 折扣百分比

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
}

type ZoneInfo struct {

	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区描述

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区状态，包含AVAILABLE和UNAVAILABLE。AVAILABLE代表可用，UNAVAILABLE代表不可用。

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type DedicatedClusterType struct {

	// 配置id

	DedicatedClusterTypeId *string `json:"DedicatedClusterTypeId,omitempty" name:"DedicatedClusterTypeId"`
	// 配置描述，对应描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 配置名称，对应计算资源类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建配置的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 支持的存储类型列表

	SupportedStorageType []*string `json:"SupportedStorageType,omitempty" name:"SupportedStorageType"`
	// 支持的上连交换机的链路传输速率

	SupportedUplinkGiB []*int64 `json:"SupportedUplinkGiB,omitempty" name:"SupportedUplinkGiB"`
	// 支持的实例族列表

	SupportedInstanceFamily []*string `json:"SupportedInstanceFamily,omitempty" name:"SupportedInstanceFamily"`
	// 地板承重要求(KG)

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 功率要求(KW)

	PowerDrawKva *float64 `json:"PowerDrawKva,omitempty" name:"PowerDrawKva"`
	// 显示计算资源规格详情，存储等资源不显示；对应规格

	ComputeFormatDesc *string `json:"ComputeFormatDesc,omitempty" name:"ComputeFormatDesc"`
	// 配置的内部名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
}

type CreateDedicatedClusterRequest struct {
	*tchttp.BaseRequest

	// 按照站点id过滤

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 专用集群的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 专用集群所属的可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 专用集群的描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDedicatedClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建Site生成的id

		SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlAccessInfo struct {

	// 接入方式

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RedisInfo struct {

	// redis规格类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CosInfo struct {

	// COS存储大小，单位TB

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// COS存储类型，默认为cos

	Type *string `json:"Type,omitempty" name:"Type"`
	// 子订单ID，用于计费侧转换参数时使用

	SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
}

type HostInfo struct {

	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 云服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 宿主机运行状态

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 宿主机类型

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// cpu可用数

	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`
	// cpu总数

	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// 内存可用数

	MemAvailable *uint64 `json:"MemAvailable,omitempty" name:"MemAvailable"`
	// 内存总数

	MemTotal *uint64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 运行时间

	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`
	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 宿主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type RegionZoneInfo struct {

	// Region&nbsp;id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// ZoneInfo数组

	Zones []*ZoneInfo `json:"Zones,omitempty" name:"Zones"`
}

type CreateDedicatedClusterUserDefinedOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单id

		DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitempty" name:"DedicatedClusterOrderId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDedicatedClusterUserDefinedOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterUserDefinedOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterInstanceTypesRequest struct {
	*tchttp.BaseRequest

	// 查询的专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeDedicatedClusterInstanceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterInstanceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDedicatedClusterOrderRequest struct {
	*tchttp.BaseRequest

	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// order关联的专用集群类型数组

	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitempty" name:"DedicatedClusterTypes"`
	// order关联的cos存储信息

	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
	// order关联的cbs存储信息

	CbsInfo *CbsInfo `json:"CbsInfo,omitempty" name:"CbsInfo"`
}

func (r *InquirePriceCreateDedicatedClusterOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDedicatedClusterOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DedicatedClusterTypeInfo struct {

	// 集群类型Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 集群类型个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 子订单ID，用于计费侧转换参数时使用

	SubOrderId *string `json:"SubOrderId,omitempty" name:"SubOrderId"`
}

type LocalNetInfo struct {

	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 路由信息

	BGPRoute *string `json:"BGPRoute,omitempty" name:"BGPRoute"`
	// 本地IP

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// 上次停止时间

	LastStopTime *string `json:"LastStopTime,omitempty" name:"LastStopTime"`
}

type MiddleWareInfo struct {

	// 中间件CKafka信息

	CkafkaInfo *CkafkaInfo `json:"CkafkaInfo,omitempty" name:"CkafkaInfo"`
	// 中间件TDMQ信息

	TDMQInfo *TDMQInfo `json:"TDMQInfo,omitempty" name:"TDMQInfo"`
}

type CbsInfo struct {

	// cbs存储大小，单位TB

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// cbs存储类型，默认为SSD

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ModifySiteInfoRequest struct {
	*tchttp.BaseRequest

	// 机房ID

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 站点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 注意事项

	Note *string `json:"Note,omitempty" name:"Note"`
	// 站点所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 站点所在省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 站点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 站点所在地区的邮编

	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`
	// 站点所在地区的详细地址信息

	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`
}

func (r *ModifySiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceInfo struct {

	// 询价单项详细信息

	PriceDetailSet []*PriceDetail `json:"PriceDetailSet,omitempty" name:"PriceDetailSet"`
	// 价格项的名称

	PriceName *string `json:"PriceName,omitempty" name:"PriceName"`
	// 价格项模块名

	PriceProperty *string `json:"PriceProperty,omitempty" name:"PriceProperty"`
	// 原价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 折扣价

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 折扣百分比

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
}

type StorageInfo struct {

	// 高性能云硬盘信息

	PremiumInfo *PremiumInfo `json:"PremiumInfo,omitempty" name:"PremiumInfo"`
	// 高性能云硬盘备份节点信息

	BackupPremiumInfo *BackupPremiumInfo `json:"BackupPremiumInfo,omitempty" name:"BackupPremiumInfo"`
	// ssd云硬盘信息

	SsdInfo *SsdInfo `json:"SsdInfo,omitempty" name:"SsdInfo"`
	// ssd云硬盘备份节点信息

	BackupSsdInfo *BackupSsdInfo `json:"BackupSsdInfo,omitempty" name:"BackupSsdInfo"`
	// 对象存储信息

	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type DescribeDedicatedClusterCosCapacityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本集群内cos容量信息，单位：‘GB’

		CosCapacity *CosCapacity `json:"CosCapacity,omitempty" name:"CosCapacity"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterCosCapacityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterCosCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterTypesRequest struct {
	*tchttp.BaseRequest

	// 模糊匹配专用集群配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 待查询的专用集群配置id列表

	DedicatedClusterTypeIds []*string `json:"DedicatedClusterTypeIds,omitempty" name:"DedicatedClusterTypeIds"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否只查询计算规格类型

	IsCompute *bool `json:"IsCompute,omitempty" name:"IsCompute"`
	// 订单配置类型族

	TypeFamily *string `json:"TypeFamily,omitempty" name:"TypeFamily"`
}

func (r *DescribeDedicatedClusterTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Property struct {

	// 属性名称

	PropertyName *string `json:"PropertyName,omitempty" name:"PropertyName"`
	// 某些属性有类型值

	PropertyType *string `json:"PropertyType,omitempty" name:"PropertyType"`
	// 属性是否必选

	PropertyRequired *string `json:"PropertyRequired,omitempty" name:"PropertyRequired"`
	// 属性初始值

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 属性最小值

	MinValue *int64 `json:"MinValue,omitempty" name:"MinValue"`
	// 属性最小起步值

	StartLeastValue *int64 `json:"StartLeastValue,omitempty" name:"StartLeastValue"`
	// 属性最大值

	MaxValue *int64 `json:"MaxValue,omitempty" name:"MaxValue"`
	// 间隔值

	StepValue *int64 `json:"StepValue,omitempty" name:"StepValue"`
	// 枚举值

	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`
}

type DeleteSitesRequest struct {
	*tchttp.BaseRequest

	// 要删除的站点id列表

	SiteIds []*string `json:"SiteIds,omitempty" name:"SiteIds"`
}

func (r *DeleteSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDedicatedClusterOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该集群内宿主机的统计信息列表

		Price *Price `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDedicatedClusterOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDedicatedClusterOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaInfo struct {

	// Ckafka类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ExtendRackInfo struct {

	// 扩展机柜数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type OutBandwidth struct {

	// 时间戳

	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 对应时间的值

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type SsdInfo struct {

	// 硬盘大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type BackupPremiumInfo struct {

	// 硬盘大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type BaseFoundationInfo struct {

	// 底座数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DedicatedCluster struct {

	// 专用集群id。如"cluster-xxxxx"。

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 专用集群所属可用区名称。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 专用集群的描述。&nbsp;注意：此字段可能返回&nbsp;null，表示取不到有效值。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 专用集群的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 专用集群的生命周期。如"PENDING"。

	LifecycleStatus *string `json:"LifecycleStatus,omitempty" name:"LifecycleStatus"`
	// 专用集群的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 专用集群所属的站点id。

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
}

type DedicatedClusterInstanceType struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 规格名称

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 网卡类型，例如：25代表25G网卡

	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 实例的CPU核数，单位：核

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例机型系列。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 机型名称。

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 本地存储块数量。

	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`
	// 内网带宽，单位Gbps。

	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`
	// 网络收发包能力，单位万PPS。

	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`
	// 处理器型号。

	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`
	// 实例的GPU数量。

	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`
	// 实例的FPGA数量。

	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`
	// 机型描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 实例是否售卖。取值范围：SELL：表示实例可购买、SOLD_OUT：表示实例已售罄

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例配额

	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
}

type InBandwidth struct {

	// 时间戳

	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 对应时间的值

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type TDMQInfo struct {

	// 规格类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateDedicatedClusterOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该集群内宿主机的统计信息列表

		HostStatisticSet []*HostStatistic `json:"HostStatisticSet,omitempty" name:"HostStatisticSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDedicatedClusterOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDedicatedClusterOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 宿主机信息

		HostInfoSet []*HostInfo `json:"HostInfoSet,omitempty" name:"HostInfoSet"`
		// 宿主机总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用集群配置列表

		DedicatedClusterTypeSet []*DedicatedClusterType `json:"DedicatedClusterTypeSet,omitempty" name:"DedicatedClusterTypeSet"`
		// 符合条件的个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClustersRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如：cluster-xxxxxxxx

	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitempty" name:"DedicatedClusterIds"`
	// 按照可用区名称过滤

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 按照站点id过滤

	SiteIds []*string `json:"SiteIds,omitempty" name:"SiteIds"`
	// 按照专用集群生命周期过滤

	LifecycleStatuses []*string `json:"LifecycleStatuses,omitempty" name:"LifecycleStatuses"`
	// 模糊匹配专用集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDedicatedClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateDedicatedClusterOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单参数转换结果

		ClusterOrder *string `json:"ClusterOrder,omitempty" name:"ClusterOrder"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterCreateDedicatedClusterOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateDedicatedClusterOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlVPNInfo struct {

	// vpn网关数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DatabaseInfo struct {

	// redis信息

	RedisInfo *RedisInfo `json:"RedisInfo,omitempty" name:"RedisInfo"`
	// tdsql信息

	TdsqlInfoSet []*TdsqlInfo `json:"TdsqlInfoSet,omitempty" name:"TdsqlInfoSet"`
}

type DeleteDedicatedClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDedicatedClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云服务器数量

		CvmCount *uint64 `json:"CvmCount,omitempty" name:"CvmCount"`
		// 宿主机数量

		HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`
		// vpn通道状态

		VpnConnectionState *string `json:"VpnConnectionState,omitempty" name:"VpnConnectionState"`
		// vpn网关监控数据

		VpngwBandwidthData *VpngwBandwidthData `json:"VpngwBandwidthData,omitempty" name:"VpngwBandwidthData"`
		// 本地网关信息

		LocalNetInfo *LocalNetInfo `json:"LocalNetInfo,omitempty" name:"LocalNetInfo"`
		// vpn网关通道监控数据

		VpnConnectionBandwidthData []*VpngwBandwidthData `json:"VpnConnectionBandwidthData,omitempty" name:"VpnConnectionBandwidthData"`
		// 宿主机详细信息

		HostDetailInfo []*string `json:"HostDetailInfo,omitempty" name:"HostDetailInfo"`
		// 宿主机待机数

		HostStandbyCount *uint64 `json:"HostStandbyCount,omitempty" name:"HostStandbyCount"`
		// 宿主机正常数

		HostNormalCount *uint64 `json:"HostNormalCount,omitempty" name:"HostNormalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterAlarmsRequest struct {
	*tchttp.BaseRequest

	// 集群id列表

	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitempty" name:"DedicatedClusterIds"`
	// 查询地域列表

	SearchRegions []*string `json:"SearchRegions,omitempty" name:"SearchRegions"`
	// 查询告警状态

	AlarmStates []*string `json:"AlarmStates,omitempty" name:"AlarmStates"`
	// 查询告警开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询告警结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询最近多长时间

	Period *string `json:"Period,omitempty" name:"Period"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序的字段名称

	OrderField []*string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序字段的升序或降序

	Order []*string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeDedicatedClusterAlarmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterAlarmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClusterHostStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该集群内宿主机的统计信息列表

		HostStatisticSet []*HostStatistic `json:"HostStatisticSet,omitempty" name:"HostStatisticSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClusterHostStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterHostStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateDedicatedClusterOrderRequest struct {
	*tchttp.BaseRequest

	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// order关联的专用集群类型数组

	DedicatedClusterTypes []*DedicatedClusterTypeInfo `json:"DedicatedClusterTypes,omitempty" name:"DedicatedClusterTypes"`
	// order关联的cos存储信息

	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
	// order关联的cbs存储信息

	CbsInfo *CbsInfo `json:"CbsInfo,omitempty" name:"CbsInfo"`
	// 总订单ID

	DedicatedClusterOrderId *string `json:"DedicatedClusterOrderId,omitempty" name:"DedicatedClusterOrderId"`
	// 是否转换为按一个月计费的订单

	NeedFinishOrderByOneMonth *bool `json:"NeedFinishOrderByOneMonth,omitempty" name:"NeedFinishOrderByOneMonth"`
}

func (r *SwitchParameterCreateDedicatedClusterOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateDedicatedClusterOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PremiumInfo struct {

	// 硬盘大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type DescribeDedicatedClusterOrdersRequest struct {
	*tchttp.BaseRequest

	// 专用集群id

	DedicatedClusterIds []*string `json:"DedicatedClusterIds,omitempty" name:"DedicatedClusterIds"`
	// 按照专用集群订单id过滤

	DedicatedClusterOrderIds *string `json:"DedicatedClusterOrderIds,omitempty" name:"DedicatedClusterOrderIds"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 订单状态为过滤条件：PENDING&nbsp;INCONSTRUCTION&nbsp;DELIVERING&nbsp;DELIVERED&nbsp;EXPIRED&nbsp;CANCELLED&nbsp;OFFLINE

	Status *string `json:"Status,omitempty" name:"Status"`
	// 订单类型为过滤条件：CREATE&nbsp;EXTEND

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeDedicatedClusterOrdersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClusterOrdersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDedicatedClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的专用集群列表

		DedicatedClusterSet []*DedicatedCluster `json:"DedicatedClusterSet,omitempty" name:"DedicatedClusterSet"`
		// 符合条件的专用集群数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDedicatedClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpngwBandwidthData struct {

	// 出带宽流量

	OutBandwidth *OutBandwidth `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
	// 入带宽流量

	InBandwidth *InBandwidth `json:"InBandwidth,omitempty" name:"InBandwidth"`
}

type DescribeSitesRequest struct {
	*tchttp.BaseRequest

	// 按照站点id过滤

	SiteIds []*string `json:"SiteIds,omitempty" name:"SiteIds"`
	// 模糊匹配站点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrepaidInfo struct {

	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeDedicatedSupportedZonesRequest struct {
	*tchttp.BaseRequest

	// 传入region列表

	Regions []*int64 `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeDedicatedSupportedZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDedicatedSupportedZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDefinedOrderDefaultValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各项配置的详细信息列表

		PropertySet []*Property `json:"PropertySet,omitempty" name:"PropertySet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDefinedOrderDefaultValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDefinedOrderDefaultValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDefinedOrderDefaultValuesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserDefinedOrderDefaultValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDefinedOrderDefaultValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

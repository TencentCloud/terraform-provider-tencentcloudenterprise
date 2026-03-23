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

package v20180410

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateVifAssociatedRequest struct {
	*tchttp.BaseRequest

	// 专线通道ID, 比如dcx-nsurd338

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
	// 通道负载均衡模式：
	// None 非冗余模式
	// LoadBalance：负载均衡
	// MasterSlave：主备

	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`
	// 关联专线通道ID，比如dcx-nsurd338

	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
}

func (r *UpdateVifAssociatedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVifAssociatedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入点信息。

		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet"`
		// 接入点数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用通道列表

		DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet"`
		// 符合条件的通道总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用通道ID

		DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryCloudAttach struct {

	// 高速上云实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 询价购买月数

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 询价类型：新购create，续费renew

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RouteFilterPrefix struct {

	// 用户侧网段地址

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type DirectConnectTunnel struct {

	// 专线通道ID

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
	// 物理专线ID

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// "专线通道状态
	// AVAILABLE：就绪
	// APPLYING：申请中
	// ALLOCATING：配置中
	// ALLOCATED：配置完成
	// ALTERING：修改中
	// DELETING：删除中
	// DELETED：删除完成
	// PENDING：待接受
	// REJECTED：拒绝

	State *string `json:"State,omitempty" name:"State"`
	// 专线的拥有者，开发商账号 ID

	DirectConnectOwnerAppId *int64 `json:"DirectConnectOwnerAppId,omitempty" name:"DirectConnectOwnerAppId"`
	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 私有网络统一 ID 或者黑石网络统一 ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 专线网关 ID

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// BGP ：BGP路由 STATIC：静态

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 用户侧网段地址

	IdcRoutes []*string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`
	// 专线通道的Vlan

	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`
	// 云侧互联IP

	CloudAddress *string `json:"CloudAddress,omitempty" name:"CloudAddress"`
	// 用户侧互联 IP

	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`
	// 专线通道名称

	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`
	// 专线通道创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 专线通道带宽值

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 关联的网络自定义探测ID

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// 是否为Nat通道

	NatType *bool `json:"NatType,omitempty" name:"NatType"`
	// 通道连接的VPC所在地域

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// BFD状态：
	// DISABLED: 关闭
	// ENABLE： 开启
	// UP：健康
	// DOWN：异常

	BfdState *string `json:"BfdState,omitempty" name:"BfdState"`
	// 专线网关名称

	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// True: 开启, False: 不开启

	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// 通道下支持的组播组

	MulticastGroups *string `json:"MulticastGroups,omitempty" name:"MulticastGroups"`
	// 专线通道的拥有者，开发上账号ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 通道负载均衡模式：
	// 0：None 1: LoadBalance 2: MasterSlave

	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`
	// 互联地址掩码

	ConnectSubnetMask *uint64 `json:"ConnectSubnetMask,omitempty" name:"ConnectSubnetMask"`
	// 专线所有者账号UIN

	DirectConnectOwnerUin *int64 `json:"DirectConnectOwnerUin,omitempty" name:"DirectConnectOwnerUin"`
	// bfd协议中的interval字段值

	BfdInterval *uint64 `json:"BfdInterval,omitempty" name:"BfdInterval"`
	// 关联的通道ID

	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
	// 通道备注

	Comments *string `json:"Comments,omitempty" name:"Comments"`
	// 在主备模式下，通道是否为主

	MasterStatus *bool `json:"MasterStatus,omitempty" name:"MasterStatus"`
	// 用于判断同地域、境内跨地域、跨境

	RegionStatus *int64 `json:"RegionStatus,omitempty" name:"RegionStatus"`
	// 是否支持bfd

	SupportBfd *bool `json:"SupportBfd,omitempty" name:"SupportBfd"`
	// 专线通道所有者账号UIN

	DirectConnectTunnelOwnerUin *int64 `json:"DirectConnectTunnelOwnerUin,omitempty" name:"DirectConnectTunnelOwnerUin"`
	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey

	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`
	// 通道IP协议类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 是否为共享通道

	ShareMode *bool `json:"ShareMode,omitempty" name:"ShareMode"`
	// 专线接入点ID

	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 物理专线名称

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
}

type ApproveDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专线通道ID

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
	// 专线通道名称

	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`
	// 审批结果, True: 同意， False: 不同意

	Approved *bool `json:"Approved,omitempty" name:"Approved"`
	// 备注

	Comments *string `json:"Comments,omitempty" name:"Comments"`
}

func (r *ApproveDirectConnectTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveDirectConnectTunnelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachInfo struct {

	// 高速上云实例id
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参必填
	// CancelCloudAttachService：入参必填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参必填
	// ConfirmCloudAttachService：入参必填
	// DeliveryCloudAttachService：入参必填

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 高速上云名称
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参选填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合作伙伴的AppId
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IapId *string `json:"IapId,omitempty" name:"IapId"`
	// 需要接入高速上云的IDC的地址
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IdcAddress *string `json:"IdcAddress,omitempty" name:"IdcAddress"`
	// 需要接入高速上云的IDC的互联网服务提供商类型
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IdcType *string `json:"IdcType,omitempty" name:"IdcType"`
	// 高速上云的带宽，单位为MB
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 联系电话
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`
	// 高速上云的状态
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参必填
	// DeliveryCloudAttachService：入参不填
	// 出参状态：
	// available：就绪状态
	// applying：申请，待审核状态
	// pendingpay：代付款状态
	// building：建设中状态
	// confirming：待确认状态
	// isolate: 隔离状态
	// stoped：终止状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 高速上云申请的时间。只作为出参

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 高速上云建设完成的时间。只作为出参

	ReadyTime *string `json:"ReadyTime,omitempty" name:"ReadyTime"`
	// 高速上云过期时间。只作为出参

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 备注信息
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 高速上云的地域状态。只作为出参
	// 出参状态：
	// same-region：同地域
	// cross-region：跨地域

	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`
	// 用户的AppId。只作为出参

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type TcapPriceDetail struct {

	// 铜牌流量价格信息

	BwBronze *TcapPartPriceDetail `json:"BwBronze,omitempty" name:"BwBronze"`
	// 银牌流量价格信息

	BwSilver *TcapPartPriceDetail `json:"BwSilver,omitempty" name:"BwSilver"`
	// 金牌流量价格信息

	BwGolden *TcapPartPriceDetail `json:"BwGolden,omitempty" name:"BwGolden"`
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理专线列表。

		DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet"`
		// 符合物理专线列表数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 物理专线的名称。

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
	// 物理专线申请者联系邮箱。默认从账户体系获取。

	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`
	// 物理专线申请者姓名。默认从账户体系获取。

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 物理专线申请者联系号码。默认从账户体系获取。

	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`
	// 物理的专线带宽。

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 是否共享
	IsShare *bool `json:"IsShare,omitempty" name:"IsShare"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcapVif struct {

	// tcap通道id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// tcap通道实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// tcap通道名称

	VifName *string `json:"VifName,omitempty" name:"VifName"`
	// tcap通道带宽

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 用户侧互联IP

	ConnLocalIp *string `json:"ConnLocalIp,omitempty" name:"ConnLocalIp"`
	// 腾讯侧互联IP

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 路由同步协议，默认bgp

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// bgp 秘钥

	BgpKey *string `json:"BgpKey,omitempty" name:"BgpKey"`
	// bgp Asn 号

	BgpAsn *string `json:"BgpAsn,omitempty" name:"BgpAsn"`
	// 接入点ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// tcap通道状态; Enum[0, 1, 2, 5] 分别表示[ready, applying, deploying, deleted]

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 申请时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 是否计费

	ChargeOrNot *bool `json:"ChargeOrNot,omitempty" name:"ChargeOrNot"`
	// 就绪时间

	ReadyTime *string `json:"ReadyTime,omitempty" name:"ReadyTime"`
	// 物理接口类型

	PhysicalPortType *string `json:"PhysicalPortType,omitempty" name:"PhysicalPortType"`
	// 接入点名称：$city-$idc

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
}

type TcapPartPriceDetail struct {

	// 折扣

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
	// 折后价

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 原价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type DescribeDirectConnectPortPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口类型

		PortType *string `json:"PortType,omitempty" name:"PortType"`
		// 端口月租费价格

		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// 端口月租费价格，优惠后的价格

		RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectPortPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectPortPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DcNumberOfAR struct {

	// 接入点数字ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 接入点下的专线数目

	NumberOfDirectConnects *uint64 `json:"NumberOfDirectConnects,omitempty" name:"NumberOfDirectConnects"`
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AccessPoint struct {

	// 接入点的名称。

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点唯一ID。

	AccessPointId *int64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 接入点的状态。AVAILABLE, UNAVAILABLE

	State *string `json:"State,omitempty" name:"State"`
	// 接入点所在机房地址

	Location *string `json:"Location,omitempty" name:"Location"`
	// 接入点支持的运营商列表。

	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 接入点可用带宽百分比

	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`
	// 当前已使用带宽, 单位MBytes

	CurrentBandwidth *uint64 `json:"CurrentBandwidth,omitempty" name:"CurrentBandwidth"`
	// 带宽告警比例

	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`
	// 接入点总带宽, 单位 MBytes

	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
	// IDC所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 物理专线用户侧接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	CloudPortType []*string `json:"CloudPortType,omitempty" name:"CloudPortType"`
	// 物理专线IDC侧接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	IdcPortType []*string `json:"IdcPortType,omitempty" name:"IdcPortType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DirectConnect struct {

	// 物理专线ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 物理专线的名称。

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
	// 物理专线的接入点ID。

	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 物理专线的状态。
	// 申请中：PENDING
	// 申请驳回：REJECTED
	// 建设中：ALLOCATED
	// 已开通：AVAILABLE
	// 删除中 ：DELETING
	// 已删除：DELETED 。

	State *string `json:"State,omitempty" name:"State"`
	// 物理专线创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 物理专线的开通时间。

	EnabledTime *string `json:"EnabledTime,omitempty" name:"EnabledTime"`
	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 本地数据中心的地理位置。

	Location *string `json:"Location,omitempty" name:"Location"`
	// 物理专线接入接口带宽，单位为Mbps。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 冗余物理专线的ID。

	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`
	// 物理专线申请者姓名。默认从账户体系获取。

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 物理专线申请者联系邮箱。默认从账户体系获取。

	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`
	// 物理专线申请者联系号码。默认从账户体系获取。

	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`
	// 报障联系人。

	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`
	// 报障联系电话。

	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
	// IDC所在城市

	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`
	// 物理专线接入IDC侧端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	IdcPortType *string `json:"IdcPortType,omitempty" name:"IdcPortType"`
	// 云侧端口类型

	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`
	// 申请ID

	ApplyId *uint64 `json:"ApplyId,omitempty" name:"ApplyId"`

	// 是否共享
	IsShare *bool `json:"IsShare,omitempty" name:"IsShare"`
}

type TcapAccessPoint struct {

	// 接入点 ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 接入点所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 接入点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 接入点所在机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 接入点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 接入点当前是否可用

	Valid *bool `json:"Valid,omitempty" name:"Valid"`
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，支持通过专线通道ID或名称进行模糊过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcapIpCertificate struct {

	// 证书ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 证书存储路径

	StoragePath *string `json:"StoragePath,omitempty" name:"StoragePath"`
	// 证书上的ipv4地址

	Ipv4Addresses *string `json:"Ipv4Addresses,omitempty" name:"Ipv4Addresses"`
	// 证书上的ipv6地址

	Ipv6Addresses *string `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
	// 证书上的asn号

	Asn *string `json:"Asn,omitempty" name:"Asn"`
	// 证书上的ID

	CertificateNumber *string `json:"CertificateNumber,omitempty" name:"CertificateNumber"`
	// 证书生效时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 证书过期时间

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 证书是否有效

	Valid *bool `json:"Valid,omitempty" name:"Valid"`
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *DeleteDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理专线的ID。

		DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitempty" name:"DirectConnectIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproveDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApproveDirectConnectTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专线 ID，例如：dc-kd7d06of

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 专用通道名称

	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`
	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID

	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`
	// 私有网络统一 ID 或者黑石网络统一 ID

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 专线网关 ID，例如 dcg-d545ddf

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 专线带宽，单位：Mbps
	// 默认是物理专线带宽值

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// BGP ：BGP路由
	// STATIC：静态
	// 默认为 BGP 路由

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey

	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`
	// 静态路由，用户IDC的网段地址

	IdcRoutes *string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`
	// Vlan，范围：11 ~ 4000，需要和用户侧的vlan id一致

	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`
	// TencentAddress，腾讯侧互联 IP

	CloudAddress *string `json:"CloudAddress,omitempty" name:"CloudAddress"`
	// CustomerAddress，用户侧互联 IP

	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`
	// 关联的冗余通道ID

	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
	// 通道负载均衡模式：
	// None 非冗余模式
	// LoadBalance：负载均衡
	// MasterSlave：主备

	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`
	// vpc数字ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 是否开启BFD

	EnableBfd *bool `json:"EnableBfd,omitempty" name:"EnableBfd"`
	// BFD协议interval配置

	BfdInterval *int64 `json:"BfdInterval,omitempty" name:"BfdInterval"`
	// 互联地址掩码

	ConnectSubnetMask *uint64 `json:"ConnectSubnetMask,omitempty" name:"ConnectSubnetMask"`
	// vpc所在地域

	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`
	// 通道IP协议类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数，支持基于接入点ID(AccessPointId)数组过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，支持State过滤，也支持通过专线ID或专线名称进行模糊过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBgpAsnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BGP类型通道云上使用的BGP ASN

		CloudBgpAsn *int64 `json:"CloudBgpAsn,omitempty" name:"CloudBgpAsn"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBgpAsnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudBgpAsnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectPortPriceRequest struct {
	*tchttp.BaseRequest

	// 要询价的端口类型

	PortType *string `json:"PortType,omitempty" name:"PortType"`
}

func (r *DescribeDirectConnectPortPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectPortPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachServiceGatewayPingTaskInfo struct {

	// 高速上云服务网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// Ping探测IP协议类型

	AddressFamily *string `json:"AddressFamily,omitempty" name:"AddressFamily"`
	// 探测的对端IP地址

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 探测次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 探测包大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 探测间隔，单位为毫秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 探测状态
	// running：任务正在进行
	// done：任务执行完毕
	// error：任务执行错误

	Status *string `json:"Status,omitempty" name:"Status"`
	// 探测任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// Ping探测成功率

	SuccessRatio *uint64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`
	// 最小延时

	LatencyMin *string `json:"LatencyMin,omitempty" name:"LatencyMin"`
	// 平均延时

	LatencyAvg *string `json:"LatencyAvg,omitempty" name:"LatencyAvg"`
	// 最大延时

	LatencyMax *string `json:"LatencyMax,omitempty" name:"LatencyMax"`
	// 任务开始时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type LocalIpTranslationNatRule struct {

	// 原始IP

	OriginalIp *string `json:"OriginalIp,omitempty" name:"OriginalIp"`
	// 映射IP

	TranslationIp *string `json:"TranslationIp,omitempty" name:"TranslationIp"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachPriceInfo struct {

	// 折扣

	Discount *string `json:"Discount,omitempty" name:"Discount"`
	// 折后价

	DiscountPrice *string `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 原价

	OriginalPrice *string `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type CloudAttachServiceGatewayPingInfo struct {

	// 高速上云服务网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// Ping探测次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// Ping探测包的大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// Ping探测间隔，单位毫秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 需要探测的对端IP地址

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
	// 专用通道名称

	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`
	// 用户侧网段地址

	IdcRoutes *string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`
	// 专用通道带宽值，单位为M。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 通道备注

	Comments *string `json:"Comments,omitempty" name:"Comments"`
	// 是否开启BFD

	EnableBfd *bool `json:"EnableBfd,omitempty" name:"EnableBfd"`
	// BFD报文interval时间

	BfdInterval *int64 `json:"BfdInterval,omitempty" name:"BfdInterval"`
	// 是否开启组播

	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// 通道组播组地址

	MulticastGroups *string `json:"MulticastGroups,omitempty" name:"MulticastGroups"`
	// BGP ASN和PASSWORD

	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`
}

func (r *ModifyDirectConnectTunnelAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachGatewayInfo struct {

	// 高速上云服务网关实例信息
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参必填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参必填

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 高速上云服务实例ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	UniqCasId *string `json:"UniqCasId,omitempty" name:"UniqCasId"`
	// 高速上云服务网关名称
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合作伙伴的AppId。只作为出参使用

	IapId *string `json:"IapId,omitempty" name:"IapId"`
	// 专线ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	DcId *string `json:"DcId,omitempty" name:"DcId"`
	// vlan ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	VlanId *uint64 `json:"VlanId,omitempty" name:"VlanId"`
	// 带宽信息，单位为MB
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参选填

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 腾讯侧互联IP
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnLocalIp *string `json:"ConnLocalIp,omitempty" name:"ConnLocalIp"`
	// 用户侧互联IP
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 互联IP掩码
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnIpMask *string `json:"ConnIpMask,omitempty" name:"ConnIpMask"`
	// 电路编码
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	CircuitNumber *string `json:"CircuitNumber,omitempty" name:"CircuitNumber"`
	// 网关路由类型，静态路由或者BGP路由
	// 取值范围：static，bgp
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 用户侧IP段信息
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	PeerIp *string `json:"PeerIp,omitempty" name:"PeerIp"`
	// BGP路由的ASN信息
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	BgpAsn *string `json:"BgpAsn,omitempty" name:"BgpAsn"`
	// BGP路由的KEY信息
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	BgpKey *string `json:"BgpKey,omitempty" name:"BgpKey"`
	// 高速上云专线通道实例ID。只作为出参使用

	UniqVifId *string `json:"UniqVifId,omitempty" name:"UniqVifId"`
	// 高速上云专线网关实例ID。只作为出参使用

	UniqVpgId *string `json:"UniqVpgId,omitempty" name:"UniqVpgId"`
	// 高速上云网关所属地域。
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Region *string `json:"Region,omitempty" name:"Region"`
	// BGP状态。只作为出参使用

	BgpStatus *string `json:"BgpStatus,omitempty" name:"BgpStatus"`
	// 网关类型：
	// cas：高速上云网关类型
	// dcg：普通网关类型（VXLAN迁移用）
	// CreateCloudAttachServiceGateway：入参选填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网关所在region的zone名称
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type DescribeCloudBgpAsnRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCloudBgpAsnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudBgpAsnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 专用通道ID

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *DeleteDirectConnectTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DcSwitch struct {

	// 专线接入点ID，通过DescribeAccessPoints接口可获取接入点列表

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 用户自定义交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口

	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 专线接入交换机vxlan隧道的VTEP ip

	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`
	// 专线接入交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
}

type SwitchPort struct {

	// 端口数字ID

	PortId *uint64 `json:"PortId,omitempty" name:"PortId"`
	// 端口名称

	PortName *string `json:"PortName,omitempty" name:"PortName"`
}

type TcapIps struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 对应的tcap 通道 ID

	TcapId *uint64 `json:"TcapId,omitempty" name:"TcapId"`
	// 4（ipv4）或者6（ipv6）

	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
	// ip网段地址，多个用','分隔

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 网段审核状态; 可能取值为: [auditing, valid, invalid]

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 包含此网段的用户证书ID

	IpCertificateId *int64 `json:"IpCertificateId,omitempty" name:"IpCertificateId"`
	// ip 审核时间

	AuditTime *string `json:"AuditTime,omitempty" name:"AuditTime"`
	// 对应的tcap 通道 instance id

	TcapInstanceId *string `json:"TcapInstanceId,omitempty" name:"TcapInstanceId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateVifAssociatedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateVifAssociatedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateVifAssociatedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachIapInfo struct {

	// 高速上云服务提供商的AppId信息

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type ISP struct {

	// ISP数字ID

	IspId *uint64 `json:"IspId,omitempty" name:"IspId"`
	// ISP名称，国际站AppId返回英文

	IspName *string `json:"IspName,omitempty" name:"IspName"`
}

type AccessPointListEntry struct {

	// 接入点数字ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 地域名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// vpc所在地域代号，比如"sh"

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// 是否对客户可见，接入点已下线时返回false

	Visible *bool `json:"Visible,omitempty" name:"Visible"`
	// 接入点名称，调用AppId在国际站时，返回英文

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点支持的运营商列表

	Isp []*ISP `json:"Isp,omitempty" name:"Isp"`
	// 接入点云端支持的端口列表

	PhysicalPort []*SwitchPort `json:"PhysicalPort,omitempty" name:"PhysicalPort"`
	// 接入点云端支持的端口ID列表

	PhysicalPortSort []*uint64 `json:"PhysicalPortSort,omitempty" name:"PhysicalPortSort"`
	// 接入点用户IDC侧支持的端口列表

	IdcPhysicalPortSort []*uint64 `json:"IdcPhysicalPortSort,omitempty" name:"IdcPhysicalPortSort"`
	// 接入点用户IDC侧支持的端口ID列表

	IdcPhysicalPort []*SwitchPort `json:"IdcPhysicalPort,omitempty" name:"IdcPhysicalPort"`
}

type CloudAttachGetPingTask struct {

	// 高速上云网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type SwitchNetdevice struct {

	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口，为空时，业务使用默认值

	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
}

type TcapUnchargeWhiteList struct {

	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type BgpPeer struct {

	// 用户侧，BGP Asn

	Asn *int64 `json:"Asn,omitempty" name:"Asn"`
	// 用户侧BGP密钥

	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的名称。

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
	// 物理专线所在的接入点。
	// 您可以通过调用 DescribeAccessPoints接口获取地域ID。所选择的接入点必须存在且处于可接入的状态。

	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 本地数据中心的地理位置。

	Location *string `json:"Location,omitempty" name:"Location"`
	// 物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`
	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 冗余物理专线的ID。

	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`
	// 物理专线申请者姓名。默认从账户体系获取。

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 物理专线申请者联系邮箱。默认从账户体系获取。

	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`
	// 物理专线申请者联系号码。默认从账户体系获取。

	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`
	// 物理专线接入IDC侧端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	IdcPortType *string `json:"IdcPortType,omitempty" name:"IdcPortType"`
	// 本地数据中心所在城市

	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`
}

func (r *CreateDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

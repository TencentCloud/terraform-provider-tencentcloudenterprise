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

type DescribeCcnBandwidthConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽上限最大值。

		MaxBand *int64 `json:"MaxBand,omitempty" name:"MaxBand"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnBandwidthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnBandwidthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeTrafficQosPolicyRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 本端地域。

	LocalRegion *string `json:"LocalRegion,omitempty" name:"LocalRegion"`
	// 远端地域。

	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
	// 服务等级信息；

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
}

func (r *DescribeTrafficQosPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficQosPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRouteMatchRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnRouteMatchRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRouteMatchRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTableSelectPolicy struct {

	// 实例类型：
	// 私有网络:&nbsp;`VPC`
	// 专线网关:&nbsp;`DIRECTCONNECT`
	// 黑石私有网络:&nbsp;`BMVPC`
	// EDGE设备:&nbsp;`EDGE`
	// EDGE隧道:&nbsp;`EDGE_TUNNEL`
	// EDGE网关:&nbsp;`EDGE_VPNGW`
	// VPN网关：`VPNGW`

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源端CIDR。

	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" name:"SourceCidrBlock"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由表备注。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type RouteSelectionPolicy struct {

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例类型。如VPC

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 源端cidr。

	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" name:"SourceCidrBlock"`
	// 路由表描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 关联实例所属UIN（根账号）。

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AuditCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// 服务商,&nbsp;可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// 表单唯一`ID`。可通过[DescribeCrossBorderCompliance](https://{{conf.main_domain}}/document/product/215/47838)接口查询ComplianceId信息

	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
	// 通过：`APPROVED&nbsp;`，拒绝：`DENY`。

	AuditBehavior *string `json:"AuditBehavior,omitempty" name:"AuditBehavior"`
}

func (r *AuditCrossBorderComplianceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditCrossBorderComplianceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCrossBorderCosTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 过期时间

		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// BucketName

		BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
		// BucketRegion

		BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
		// 临时SecretId

		TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
		// 临时SecretKey

		TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
		// XCosSecurityToken

		XCosSecurityToken *string `json:"XCosSecurityToken,omitempty" name:"XCosSecurityToken"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCrossBorderCosTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCrossBorderCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockCcnBandwidthsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LockCcnBandwidthsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockCcnBandwidthsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderFlowMonitorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网跨境带宽监控数据

		CrossBorderFlowMonitorData []*CrossBorderFlowMonitorData `json:"CrossBorderFlowMonitorData,omitempty" name:"CrossBorderFlowMonitorData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderFlowMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）地域带宽详情。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *InquiryPriceCreateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest struct {
	*tchttp.BaseRequest

	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 带宽，单位Gbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectPortMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDirectConnectPortMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDirectConnectPortMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 关联实例列表。

		InstanceSet []*CcnAttachedInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SourceDestinationRegion struct {

	// 本端地域，例如：region1

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 对端地域，例如：region1

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
}

type UnlockCcnsRequest struct {
	*tchttp.BaseRequest

	// 云联网实例数组。

	Instances []*CcnFlowLock `json:"Instances,omitempty" name:"Instances"`
}

func (r *UnlockCcnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockCcnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDealStatusByNameRequest struct {
	*tchttp.BaseRequest

	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *GetDealStatusByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDealStatusByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficQosPolicyRequest struct {
	*tchttp.BaseRequest

	// ccn-gjug0kul	CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 流量调度策略唯一ID。

	QosPolicyId *string `json:"QosPolicyId,omitempty" name:"QosPolicyId"`
	// 描述。

	QosPolicyDescription *string `json:"QosPolicyDescription,omitempty" name:"QosPolicyDescription"`
	// 名称。

	QosPolicyName *string `json:"QosPolicyName,omitempty" name:"QosPolicyName"`
	// 带宽。

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *ModifyTrafficQosPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficQosPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云联网地域间限速带宽实例的信息。

		CcnBandwidthSet []*CcnBandwidth `json:"CcnBandwidthSet,omitempty" name:"CcnBandwidthSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpdateCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpdateCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateCcnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateCcnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateCcnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号。

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnPolicyBasedRoutingNextHopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网策略路由下一跳信息

		CcnPolicyBasedRoutingNextHopSet []*CcnPolicyBasedRoutingNextHop `json:"CcnPolicyBasedRoutingNextHopSet,omitempty" name:"CcnPolicyBasedRoutingNextHopSet"`
		// 符合条件的对象数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnPolicyBasedRoutingNextHopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnPolicyBasedRoutingNextHopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// 发证机关。

	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`
	// 经办人身份证地址。

	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`
	// 服务受理单。

	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitempty" name:"ServiceHandlingForm"`
	// 授权函。

	AuthorizationLetter *string `json:"AuthorizationLetter,omitempty" name:"AuthorizationLetter"`
	// 期望服务截止时间。

	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`
	// 法人身份证

	LegalPersonIdCard *string `json:"LegalPersonIdCard,omitempty" name:"LegalPersonIdCard"`
	// 营业执照附件。

	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`
	// 营业执照住所。

	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`
	// 经办人身份证号。

	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
	// 经办人身份证附件。

	ManagerIdCard *string `json:"ManagerIdCard,omitempty" name:"ManagerIdCard"`
	// 电子邮箱。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 信息安全承诺书。

	SafetyCommitment *string `json:"SafetyCommitment,omitempty" name:"SafetyCommitment"`
	// 服务商，可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// 经办人。

	Manager *string `json:"Manager,omitempty" name:"Manager"`
	// 经办人联系电话。

	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`
	// 合规化审批单`ID`。

	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
	// 统一社会信用代码。

	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`
	// 法定代表人。

	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`
	// 邮编。

	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`
	// 期望服务开始时间。

	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`
	// 法人身份证号。

	LegalPersonId *string `json:"LegalPersonId,omitempty" name:"LegalPersonId"`
	// 公司全称。

	Company *string `json:"Company,omitempty" name:"Company"`
}

func (r *ModifyCrossBorderComplianceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderComplianceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnPolicyBasedRoutingRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网策略路由下一跳信息

		CcnPolicyBasedRoutingRuleSet []*CcnPolicyBasedRoutingRule `json:"CcnPolicyBasedRoutingRuleSet,omitempty" name:"CcnPolicyBasedRoutingRuleSet"`
		// 符合条件的对象数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnPolicyBasedRoutingRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnPolicyBasedRoutingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeRequest struct {
	*tchttp.BaseRequest

	// 云联网实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网限速类型，INTER_REGION_LIMIT：地域间限速，OUTER_REGION_LIMIT：地域出口限速。默认值：OUTER_REGION_LIMIT。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRouteMatchRuleRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 路由匹配规则唯一ID。

	CcnRouteMatchRuleIds []*string `json:"CcnRouteMatchRuleIds,omitempty" name:"CcnRouteMatchRuleIds"`
}

func (r *DeleteCcnRouteMatchRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRouteMatchRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRenewDirectConnectAccelerateChannelBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *GetRenewDirectConnectAccelerateChannelBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewDirectConnectAccelerateChannelBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 云联网（CCN）地域带宽上限。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *InquiryPriceUpdateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceUpdateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockCcnBandwidthsRequest struct {
	*tchttp.BaseRequest

	// 带宽实例的唯一ID数组。

	Instances []*CcnFlowLock `json:"Instances,omitempty" name:"Instances"`
}

func (r *LockCcnBandwidthsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockCcnBandwidthsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBorderFlowMonitorData struct {

	// 入包，单位：`pps`。

	InPkg []*int64 `json:"InPkg,omitempty" name:"InPkg"`
	// 出包，单位：`pps`。

	OutPkg []*int64 `json:"OutPkg,omitempty" name:"OutPkg"`
	// 入带宽，单位：`bps`。

	InBandwidth []*int64 `json:"InBandwidth,omitempty" name:"InBandwidth"`
	// 出带宽，单位：`bps`。

	OutBandwidth []*int64 `json:"OutBandwidth,omitempty" name:"OutBandwidth"`
}

type TrafficQosPolicySet struct {

	// 名称。

	QosPolicyName *string `json:"QosPolicyName,omitempty" name:"QosPolicyName"`
	// 带宽。

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 流量调度策略ID。

	QosPolicyId *string `json:"QosPolicyId,omitempty" name:"QosPolicyId"`
	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// qos&nbsp;id。

	QosId *uint64 `json:"QosId,omitempty" name:"QosId"`
	// 描述。

	QosPolicyDescription *string `json:"QosPolicyDescription,omitempty" name:"QosPolicyDescription"`
}

type CcnRouteTableAggregatePolicy struct {

	// 路由条件。

	RouteConditions []*CcnRouteBroadcastPolicyRouteCondition `json:"RouteConditions,omitempty" name:"RouteConditions"`
	// 路由行为，`accept`&nbsp;允许，`drop`&nbsp;拒绝。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 聚合范围

	AggregationScope []*CcnRouteBroadcastPolicyRouteCondition `json:"AggregationScope,omitempty" name:"AggregationScope"`
	// 策略描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 聚合路由cidr

	AggregationRoutes []*CcnAggregationRoute `json:"AggregationRoutes,omitempty" name:"AggregationRoutes"`
}

type DirectConnectMapSett struct {

	// 源地域。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 源专线ID。

	SourceDirectConnectId *string `json:"SourceDirectConnectId,omitempty" name:"SourceDirectConnectId"`
	// 目的专线ID。

	DestinationDirectConnectId *string `json:"DestinationDirectConnectId,omitempty" name:"DestinationDirectConnectId"`
	// 备注。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeCrossBorderSettlementRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCrossBorderSettlementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderSettlementRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 拒绝关联实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *RejectAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCreateDirectConnectAccelerateChannelBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 带宽，单位Gbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *GetCreateDirectConnectAccelerateChannelBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateDirectConnectAccelerateChannelBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnBandwidthRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnBandwidthRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnBandwidthRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。可通过DescribeCcnRoutes接口获取。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 过滤条件。
	// <li>sregion&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）源地域，形如：region1。</li><li>dregion&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）目的地域，形如：region1-bm</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序条件，目前支持带宽（`BandwidthLimit`）和过期时间（`ExpireTime`），默认按&nbsp;`ExpireTime`&nbsp;排序。

	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方式，'ASC':升序,'DESC':降序。默认按'ASC'排序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *GetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTableAggregatePolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由接收策略版本号。

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DescribeCcnRouteTableAggregatePolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableAggregatePolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesToCcnRouteTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateInstancesToCcnRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateInstancesToCcnRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantCcnsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单页返回数据量，可选值0到100之间的整数，默认20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，目前`value`值个数只支持一个，允许可支持的字段有：
	// <li>`ccn-ids`&nbsp;云联网ID数组，值形如：`["ccn-12345678"]`</li>&nbsp;
	// <li>`user-account-id`&nbsp;用户账号ID，值形如：`["12345678"]`</li><li>`is-security-lock`&nbsp;是否锁定，值形如：`["true"]`</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTenantCcnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantCcnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionDefaultQosBandwidthLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnRegionDefaultQosBandwidthLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionDefaultQosBandwidthLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）对象。

		Ccn *CCN `json:"Ccn,omitempty" name:"Ccn"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnPolicyBasedRoutingRule struct {

	// 策略路由下一跳ID

	PolicyBasedRoutingNextHopId *string `json:"PolicyBasedRoutingNextHopId,omitempty" name:"PolicyBasedRoutingNextHopId"`
	// 实例类型[VPC,DIRECTCONNECT,VPNGW]

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 源地址CIDR

	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" name:"SourceCidrBlock"`
	// 目的地址CIDR

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略路由匹配策略ID

	PolicyBasedRoutingRuleId *string `json:"PolicyBasedRoutingRuleId,omitempty" name:"PolicyBasedRoutingRuleId"`
}

type CreateDirectConnectAccelerateChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加速通道唯一ID。

		DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectAccelerateChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectAccelerateChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectPortMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口映射信息。

		DirectConnectPortMapSet []*DirectConnectMapSett `json:"DirectConnectPortMapSet,omitempty" name:"DirectConnectPortMapSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectPortMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectPortMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）各地域带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *CreateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号。

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstanceInfo struct {

	// 云联网唯一ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网描述信息。

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// `true`表示封禁，所有经过该云联网的流量都不通`false`解禁，流量正常

	IsSecurityLock *bool `json:"IsSecurityLock,omitempty" name:"IsSecurityLock"`
	// 付费类型，PREPAID为预付费，POSTPAID为后付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 限速类型，INTER_REGION_LIMIT为地域间限速；OUTER_REGION_LIMIT为地域出口限速。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
	// 是否跨境，`true`表示跨境，反之不跨境。

	IsCrossBorder *bool `json:"IsCrossBorder,omitempty" name:"IsCrossBorder"`
	// 公司名称。

	Company *string `json:"Company,omitempty" name:"Company"`
	// 云联网名称。

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// 关联实例数量。

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例状态，&nbsp;'ISOLATED':&nbsp;隔离中（欠费停服），'AVAILABLE'：运行中。

	State *string `json:"State,omitempty" name:"State"`
	// 实例服务质量，’PT’：白金，'AU'：金，'AG'：银。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 实例所属用户主账号ID。

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
}

type CcnRouteTableInputPolicy struct {

	// 路由条件。

	RouteConditions []*CcnRouteBroadcastPolicyRouteCondition `json:"RouteConditions,omitempty" name:"RouteConditions"`
	// 路由行为，`accept`&nbsp;允许，`drop`&nbsp;拒绝。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 策略描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// as-path操作

	OperateAsPath *string `json:"OperateAsPath,omitempty" name:"OperateAsPath"`
	// as-path操作模式

	AsPathOperateMode *string `json:"AsPathOperateMode,omitempty" name:"AsPathOperateMode"`
}

type BandwidthLimitForCcnAlarmOnlyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽上限

		CcnRegionBandwidthLimitSet []*BandwidthLimitForCcnAlarmOnly `json:"CcnRegionBandwidthLimitSet,omitempty" name:"CcnRegionBandwidthLimitSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BandwidthLimitForCcnAlarmOnlyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BandwidthLimitForCcnAlarmOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttachedInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 关联网络实例列表

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *ModifyCcnAttachedInstancesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttachedInstancesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCreateCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCreateCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTableAggregatePolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由表接收策略。

		PolicySet []*CcnRouteTableAggregatePolicys `json:"PolicySet,omitempty" name:"PolicySet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRouteTableAggregatePolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableAggregatePolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnPolicyBasedRoutingRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略路由匹配规则ID

		PolicyBasedRoutingRuleIdSet []*string `json:"PolicyBasedRoutingRuleIdSet,omitempty" name:"PolicyBasedRoutingRuleIdSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnPolicyBasedRoutingRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnPolicyBasedRoutingRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstanceNonDirectFlagRequest struct {
	*tchttp.BaseRequest

	// 业务类型

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// VPC所在地域

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeCcnInstanceNonDirectFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnInstanceNonDirectFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnPolicyBasedRoutingRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnPolicyBasedRoutingRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnPolicyBasedRoutingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// 公司全称。

	Company *string `json:"Company,omitempty" name:"Company"`
	// 营业执照住所。

	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`
	// 经办人。

	Manager *string `json:"Manager,omitempty" name:"Manager"`
	// 授权函。

	AuthorizationLetter *string `json:"AuthorizationLetter,omitempty" name:"AuthorizationLetter"`
	// 统一社会信用代码。

	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`
	// 邮编。

	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`
	// 经办人联系电话。

	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`
	// 电子邮箱。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 服务受理单。

	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitempty" name:"ServiceHandlingForm"`
	// 信息安全承诺书。

	SafetyCommitment *string `json:"SafetyCommitment,omitempty" name:"SafetyCommitment"`
	// 法人身份证号。

	LegalPersonId *string `json:"LegalPersonId,omitempty" name:"LegalPersonId"`
	// 服务商，可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// 经办人身份证地址。

	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`
	// 期望服务截止时间。

	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`
	// 法定代表人。

	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`
	// 发证机关。

	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`
	// 营业执照。

	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`
	// 经办人身份证号。

	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
	// 经办人身份证。

	ManagerIdCard *string `json:"ManagerIdCard,omitempty" name:"ManagerIdCard"`
	// 期望服务开始时间。

	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`
	// 法人身份证。

	LegalPersonIdCard *string `json:"LegalPersonIdCard,omitempty" name:"LegalPersonIdCard"`
}

func (r *CreateCrossBorderComplianceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCrossBorderComplianceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRefundCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 退费云联网地域间带宽获取订单参数

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRefundCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefundCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnFlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启峰值带宽采集。

		PeakBpsSwitch *bool `json:"PeakBpsSwitch,omitempty" name:"PeakBpsSwitch"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnPolicyBasedRoutingRuleAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnPolicyBasedRoutingRuleAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnPolicyBasedRoutingRuleAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockCcnsRequest struct {
	*tchttp.BaseRequest

	// 云联网实例数组&nbsp;。

	Instances []*CcnFlowLock `json:"Instances,omitempty" name:"Instances"`
}

func (r *LockCcnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockCcnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnIpv6RoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CCN路由策略对象。

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnIpv6RoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnIpv6RoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTableAggregatePolicys struct {

	// 策略列表。

	Policys []*CcnRouteTableAggregatePolicy `json:"Policys,omitempty" name:"Policys"`
	// 版本号。

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateCcnRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表信息列表。

		CcnRouteTableSet []*CcnRouteTable `json:"CcnRouteTableSet,omitempty" name:"CcnRouteTableSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnPolicyBasedRoutingRuleRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 策略路由匹配规则ID

	PolicyBasedRoutingRuleIds []*string `json:"PolicyBasedRoutingRuleIds,omitempty" name:"PolicyBasedRoutingRuleIds"`
}

func (r *DeleteCcnPolicyBasedRoutingRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnPolicyBasedRoutingRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteReceivingPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由接收策略信息集合。

		RouteReceivingPolicySet []*RouteReceivingPolicy `json:"RouteReceivingPolicySet,omitempty" name:"RouteReceivingPolicySet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteReceivingPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteReceivingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRoute struct {

	// 实例类型

	AliasType *string `json:"AliasType,omitempty" name:"AliasType"`
	// 目的端

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型（关联实例类型），所有类型：VPC、DIRECTCONNECT

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 路由的扩展状态

	ExtraState *string `json:"ExtraState,omitempty" name:"ExtraState"`
	// 路由优先级

	RoutePriority *uint64 `json:"RoutePriority,omitempty" name:"RoutePriority"`
	// 下一跳扩展名称（关联实例的扩展名称）

	InstanceExtraName *string `json:"InstanceExtraName,omitempty" name:"InstanceExtraName"`
	// 下一跳（关联实例）

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否动态路由

	IsBgp *bool `json:"IsBgp,omitempty" name:"IsBgp"`
	// 实例id

	AliasInstanceId *string `json:"AliasInstanceId,omitempty" name:"AliasInstanceId"`
	// 路由策略ID

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 下一跳名称（关联实例名称）

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 下一跳所属地域（关联实例所属地域）

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 路由是否启用

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 关联实例所属UIN（根账号）

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
}

type GetRenewCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 带宽配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 续费变配参数，可以在续费的同时降配云联网地域间带宽，暂不支持升配。

	MaxBandwidthLimit *uint64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *GetRenewCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID，形如：`ccn-gree226l`。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID，形如：`ccnr-f49l6u0z`。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
	// 过滤条件，参数不支持同时指定RouteIds和Filters。
	// <li>route-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由策略ID。</li><li>cidr-block&nbsp;-&nbsp;String&nbsp;-（过滤条件）目的端。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳类型。</li><li>instance-region&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳所属地域。</li><li>instance-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳实例ID。</li><li>route-table-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表ID列表，形如ccntr-1234edfr，可以根据路由表ID&nbsp;过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnPolicyBasedRoutingNextHopRequest struct {
	*tchttp.BaseRequest

	// 云联网&nbsp;ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnPolicyBasedRoutingNextHopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnPolicyBasedRoutingNextHopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnRouteMatchRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由匹配规则唯一ID。

		CcnRouteMatchRuleId *string `json:"CcnRouteMatchRuleId,omitempty" name:"CcnRouteMatchRuleId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnRouteMatchRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRouteMatchRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTrafficQosPolicyRequest struct {
	*tchttp.BaseRequest

	// 本端地域。

	LocalRegion *string `json:"LocalRegion,omitempty" name:"LocalRegion"`
	// 远端地域。

	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
	// 限制带宽值，单位MB。

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 描述。

	QosPolicyDescription *string `json:"QosPolicyDescription,omitempty" name:"QosPolicyDescription"`
	// 服务等级。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// qos&nbsp;id，范围1-64。

	QosId *uint64 `json:"QosId,omitempty" name:"QosId"`
	// 名称。

	QosPolicyName *string `json:"QosPolicyName,omitempty" name:"QosPolicyName"`
}

func (r *CreateTrafficQosPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficQosPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 带宽配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 带宽上限，单位：Mbps。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *UpdateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnInstanceNonDirectFlagRequest struct {
	*tchttp.BaseRequest

	// 业务类型

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// VPC所在地域

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DisableCcnInstanceNonDirectFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnInstanceNonDirectFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBorderRegionConfig struct {

	// 本端地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 本端地域是否黑石地域。

	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`
	// 对端地域。

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 对端地域是否黑石地域。

	DstIsBm *bool `json:"DstIsBm,omitempty" name:"DstIsBm"`
}

type ModifyCrossBorderSettlementAttributeRequest struct {
	*tchttp.BaseRequest

	// 账单ID

	SettlementId *string `json:"SettlementId,omitempty" name:"SettlementId"`
	// 折扣价格

	DiscountPrice *string `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyCrossBorderSettlementAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderSettlementAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN所属UIN（根账号）。

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
	// 重新申请关联网络实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *ResetAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTableBroadcastPolicys struct {

	// 策略列表

	Policys []*CcnRouteTableBroadcastPolicy `json:"Policys,omitempty" name:"Policys"`
	// 版本号

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteRouteReceivingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 路由表id。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 实例信息。

	Instances []*CcnInstanceWithoutRegion `json:"Instances,omitempty" name:"Instances"`
}

func (r *DeleteRouteReceivingPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteReceivingPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCrossBorderCosTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCrossBorderCosTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCrossBorderCosTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnTrafficPrice struct {

	// 总价

	TotalPrice *float64 `json:"TotalPrice,omitempty" name:"TotalPrice"`
	// 原始单价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 折扣单价

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 数量

	Number *string `json:"Number,omitempty" name:"Number"`
	// 所属地区

	Location *string `json:"Location,omitempty" name:"Location"`
	// 单位

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// UnitPrice

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
}

type DeleteCcnRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DeleteCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttachedInstancesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnAttachedInstancesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttachedInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// 限速实例的ID。

	RegionFlowControlIds []*string `json:"RegionFlowControlIds,omitempty" name:"RegionFlowControlIds"`
}

func (r *DeleteCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearRouteTableSelectionPoliciesRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *ClearRouteTableSelectionPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearRouteTableSelectionPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnPolicyBasedRoutingNextHopAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnPolicyBasedRoutingNextHopAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnPolicyBasedRoutingNextHopAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCcnRouteTableInputPolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCcnRouteTableInputPolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableInputPolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnIpv6RoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCcnIpv6RoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnIpv6RoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectAccelerateChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectAccelerateChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectAccelerateChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesToCcnRouteTableRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 实例列表。

	Instances []*CcnInstanceWithoutRegion `json:"Instances,omitempty" name:"Instances"`
}

func (r *AssociateInstancesToCcnRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateInstancesToCcnRouteTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderSettlementResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 跨境账单对象

		CrossBorderSettlementSet []*CrossBorderSettlement `json:"CrossBorderSettlementSet,omitempty" name:"CrossBorderSettlementSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderSettlementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderSettlementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnlockCcnBandwidthsRequest struct {
	*tchttp.BaseRequest

	// 带宽实例对象数组。

	Instances []*CcnFlowLock `json:"Instances,omitempty" name:"Instances"`
}

func (r *UnlockCcnBandwidthsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockCcnBandwidthsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderRegionConfigRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：`ccn-f49l6u0z`。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DescribeCrossBorderRegionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderRegionConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTableBroadcastPolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表传播策略。

		PolicySet []*CcnRouteTableBroadcastPolicys `json:"PolicySet,omitempty" name:"PolicySet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRouteTableBroadcastPolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableBroadcastPolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽上限

		CcnRegionBandwidthLimitSet []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimitSet,omitempty" name:"CcnRegionBandwidthLimitSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnBandwidthRegion struct {

	// 原地域。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
}

type ModifyCcnRouteMatchRuleRequest struct {
	*tchttp.BaseRequest

	// 路由匹配规则唯一ID。

	CcnRouteMatchRuleId *string `json:"CcnRouteMatchRuleId,omitempty" name:"CcnRouteMatchRuleId"`
	// 备注。

	CcnRouteMatchRuleDescription *string `json:"CcnRouteMatchRuleDescription,omitempty" name:"CcnRouteMatchRuleDescription"`
	// 优先级。

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *ModifyCcnRouteMatchRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRouteMatchRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrafficQosPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTrafficQosPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficQosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CCN路由策略对象。

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DescribeCcnPolicyBasedRoutingRuleRequest struct {
	*tchttp.BaseRequest

	// 云联网&nbsp;ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnPolicyBasedRoutingRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnPolicyBasedRoutingRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnInstanceNonDirectFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCcnInstanceNonDirectFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnInstanceNonDirectFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDealStatusByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号对应的订单状态，UNPAID(未支付)，PAID(已支付)，DELIVERING(发货中)，DELIVER_SUCC(发货成功)，DELIVER_FAIL(发货失败)，REFUNDED(已退款)，CANCELED(已取消)，DENY_PAIDBYOTHERS(代付拒绝)，UNKOWN(未知状态)

		DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`
		// 订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDealStatusByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDealStatusByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCN struct {

	// 云联网名称

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// 实例服务质量，’PT’：白金，'AU'：金，'AG'：银。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 是否开启等价路由功能。`False`&nbsp;未开启，`True`&nbsp;开启。

	RouteECMPFlag *bool `json:"RouteECMPFlag,omitempty" name:"RouteECMPFlag"`
	// 是否开启路由重叠功能。`False`&nbsp;未开启，`True`&nbsp;开启。

	RouteOverlapFlag *bool `json:"RouteOverlapFlag,omitempty" name:"RouteOverlapFlag"`
	// 是否开启路由表选择策略。

	RouteSelectPolicyFlag *bool `json:"RouteSelectPolicyFlag,omitempty" name:"RouteSelectPolicyFlag"`
	// 是否支持策略路由

	PolicyBasedRoutingFlag *bool `json:"PolicyBasedRoutingFlag,omitempty" name:"PolicyBasedRoutingFlag"`
	// 云联网描述信息

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 付费类型，PREPAID为预付费，POSTPAID为后付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 计量类型

	InstanceMeteringType *string `json:"InstanceMeteringType,omitempty" name:"InstanceMeteringType"`
	// 是否开启云联网路由传播策略。`False`&nbsp;未开启，`True`&nbsp;开启。

	RouteBroadcastPolicyFlag *bool `json:"RouteBroadcastPolicyFlag,omitempty" name:"RouteBroadcastPolicyFlag"`
	// 是否支持ipv6路由表

	Ipv6Flag *bool `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// 是否支持AsPath策略值

	MrtbPolicyValueFlag *bool `json:"MrtbPolicyValueFlag,omitempty" name:"MrtbPolicyValueFlag"`
	// 实例状态，&nbsp;'ISOLATED':&nbsp;隔离中（欠费停服），'AVAILABLE'：运行中。

	State *string `json:"State,omitempty" name:"State"`
	// 是否开启QOS。

	TrafficMarkingPolicyFlag *bool `json:"TrafficMarkingPolicyFlag,omitempty" name:"TrafficMarkingPolicyFlag"`
	// 是否开启二层云联网通道。

	DirectConnectAccelerateChannelFlag *bool `json:"DirectConnectAccelerateChannelFlag,omitempty" name:"DirectConnectAccelerateChannelFlag"`
	// 是否支持Community策略值

	RouteTablePolicyValueCommunityFlag *bool `json:"RouteTablePolicyValueCommunityFlag,omitempty" name:"RouteTablePolicyValueCommunityFlag"`
	// 关联实例数量

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 限速类型，`INTER_REGION_LIMIT`&nbsp;为地域间限速；`OUTER_REGION_LIMIT`&nbsp;为地域出口限速。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
	// 标签键值对。

	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
	// 是否支持云联网路由优先级的功能。`False`：不支持，`True`：支持。

	RoutePriorityFlag *bool `json:"RoutePriorityFlag,omitempty" name:"RoutePriorityFlag"`
	// 实例关联的路由表个数。

	RouteTableCount *uint64 `json:"RouteTableCount,omitempty" name:"RouteTableCount"`
	// 是否开启云联网多路由表特性。`False`：未开启，`True`：开启。

	RouteTableFlag *bool `json:"RouteTableFlag,omitempty" name:"RouteTableFlag"`
	// `true`：实例已被封禁，流量不通，`false`:解封禁。

	IsSecurityLock *bool `json:"IsSecurityLock,omitempty" name:"IsSecurityLock"`
	// 是否支持路由表聚合策略

	MrtbAggregatePolicyFlag *bool `json:"MrtbAggregatePolicyFlag,omitempty" name:"MrtbAggregatePolicyFlag"`
	// 云联网唯一ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type CreateDirectConnectAccelerateChannelRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 地域A。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 地域B。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 指定绑定的标签列表，例如：[{"Key":&nbsp;"city",&nbsp;"Value":&nbsp;"region1"}]

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateDirectConnectAccelerateChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectAccelerateChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。可通过DescribeCcnRoutes获取。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DisableCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCcnRouteTableAggregatePolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 新的路由接收策略。

	Policys []*CcnRouteTableAggregatePolicy `json:"Policys" name:"Policys"`
}

func (r *ReplaceCcnRouteTableAggregatePolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableAggregatePolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTrafficQosPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量调度策略唯一ID。

		QosPolicyId *string `json:"QosPolicyId,omitempty" name:"QosPolicyId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTrafficQosPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficQosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceBind struct {

	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 关联实例状态：
	// <li>`PENDING`：申请中</li><li>`ACTIVE`：已连接</li><li>`EXPIRED`：已过期</li><li>`REJECTED`：已拒绝</li><li>`DELETED`：已删除</li><li>`FAILED`：失败的（2小时后将异步强制解关联）</li><li>`ATTACHING`：关联中</li><li>`DETACHING`：解关联中</li><li>`DETACHFAILED`：解关联失败（2小时后将异步强制解关联）</li>

	State *string `json:"State,omitempty" name:"State"`
	// 实例绑定路由表的时间。

	InstanceBindTime *string `json:"InstanceBindTime,omitempty" name:"InstanceBindTime"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在地域。

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 实例所属的账户uin。

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例类型：VPC，DIRECTCONNECT，BMVPC，EDGE，EDGE_TUNNEL，EDGE_VPNGW，VPNGW。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type CreateCcnPolicyBasedRoutingNextHopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略路由下一跳ID

		PolicyBasedRoutingNextHopId *string `json:"PolicyBasedRoutingNextHopId,omitempty" name:"PolicyBasedRoutingNextHopId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnPolicyBasedRoutingNextHopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnPolicyBasedRoutingNextHopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnInstanceNonDirectFlagRequest struct {
	*tchttp.BaseRequest

	// VPC所在地域

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 业务类型

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *EnableCcnInstanceNonDirectFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnInstanceNonDirectFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableSelectionPoliciesRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 选择策略信息集合，表示需要按照当前的策略来修改。

	SelectionPolicies []*CcnRouteTableSelectPolicy `json:"SelectionPolicies,omitempty" name:"SelectionPolicies"`
}

func (r *ModifyRouteTableSelectionPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableSelectionPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeRequest struct {
	*tchttp.BaseRequest

	// 是否开启qos功能。`False`&nbsp;不开启，`True`&nbsp;开启。

	TrafficMarkingPolicyFlag *bool `json:"TrafficMarkingPolicyFlag,omitempty" name:"TrafficMarkingPolicyFlag"`
	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN名称，最大长度不能超过60个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// CCN描述信息，最大长度不能超过100个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// 是否开启等价路由功能。`False`&nbsp;不开启，`True`&nbsp;开启。

	RouteECMPFlag *bool `json:"RouteECMPFlag,omitempty" name:"RouteECMPFlag"`
	// 是否开启路由重叠功能。`False`&nbsp;不开启，`True`&nbsp;开启。

	RouteOverlapFlag *bool `json:"RouteOverlapFlag,omitempty" name:"RouteOverlapFlag"`
}

func (r *ModifyCcnAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CNN）地域带宽详情。

		CcnBandwidth *CcnBandwidthInfo `json:"CcnBandwidth,omitempty" name:"CcnBandwidth"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnRouteMatchRuleRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 优先级，范围1到100，默认为1。

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 源网段，默认为ALL。

	SrcCidr *string `json:"SrcCidr,omitempty" name:"SrcCidr"`
	// 源端口，默认为ALL。

	SrcPort *string `json:"SrcPort,omitempty" name:"SrcPort"`
	// 目的网段，默认为ALL。

	DstCidr *string `json:"DstCidr,omitempty" name:"DstCidr"`
	// 协议，支持'TCP',&nbsp;'UDP',&nbsp;'ICMP',&nbsp;'ALL'，默认为ALL。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 当前地域。

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// qos&nbsp;id。

	QosId *uint64 `json:"QosId,omitempty" name:"QosId"`
	// 目的端口，默认为ALL。

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 备注。

	CcnRouteMatchRuleDescription *string `json:"CcnRouteMatchRuleDescription,omitempty" name:"CcnRouteMatchRuleDescription"`
}

func (r *CreateCcnRouteMatchRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRouteMatchRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsRequest struct {
	*tchttp.BaseRequest

	// 排序字段。支持：`CcnId`&nbsp;`CcnName`&nbsp;`CreateTime`&nbsp;`State`&nbsp;`QosLevel`。默认值:&nbsp;`CreateTime`

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// CCN实例ID。形如：ccn-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定CcnIds和Filters。

	CcnIds []*string `json:"CcnIds,omitempty" name:"CcnIds"`
	// 过滤条件，参数不支持同时指定CcnIds和Filters。
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）CCN唯一ID，形如：`ccn-f49l6u0z`。</li><li>ccn-name&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）CCN名称。</li><li>ccn-description&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）CCN描述。</li><li>state&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）实例状态，&nbsp;'ISOLATED':&nbsp;隔离中（欠费停服），'AVAILABLE'：运行中。</li><li>tag-key&nbsp;-&nbsp;String&nbsp;-是否必填：否-&nbsp;（过滤条件）按照标签键进行过滤。</li><li>tag:tag-key&nbsp;-&nbsp;String&nbsp;-&nbsp;是否必填：否&nbsp;-&nbsp;（过滤条件）按照标签键值对进行过滤。&nbsp;tag-key使用具体的标签键进行替换。使用请参考示例：查询绑定了标签的CCN列表。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirectConnectAccelerateChannelSet struct {

	// 状态。0:&nbsp;'专线接入中',&nbsp;1:&nbsp;'可购买带宽',&nbsp;2:&nbsp;'可配置端口映射',&nbsp;3:&nbsp;'运行中',&nbsp;4:&nbsp;'隔离'。

	State *uint64 `json:"State,omitempty" name:"State"`
	// 到期时间。

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 自动续费类型。NOTIFY_AND_AUTO_RENEW：自动续费；NOTIFY_AND_MANUAL_RENEW：不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 带宽上限，单位Gbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 标签信息。

	TagSet []*TagSet `json:"TagSet,omitempty" name:"TagSet"`
	// 可配置带宽列表。

	ConfigurableBandwidthSet []*uint64 `json:"ConfigurableBandwidthSet,omitempty" name:"ConfigurableBandwidthSet"`
	// 源地域。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 加速通道唯一ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 供应商。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
}

type DescribeRouteTableAssociatedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的绑定路由表的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 绑定信息。

		InstanceBindSet []*InstanceBind `json:"InstanceBindSet,omitempty" name:"InstanceBindSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTableAssociatedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableAssociatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteReceivingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 路由表id。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 实例信息。

	Instances []*CcnInstanceWithoutRegion `json:"Instances,omitempty" name:"Instances"`
}

func (r *CreateRouteReceivingPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteReceivingPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tags struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeDirectConnectAccelerateChannelRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道唯一ID列表。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
}

func (r *DescribeDirectConnectAccelerateChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectAccelerateChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCrossBorderSettlementAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCrossBorderSettlementAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderSettlementAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCrossBorderWhitelistFlagRequest struct {
	*tchttp.BaseRequest

	// 发起审批的用户主账号ID。

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
	// 是否开启白名单

	WhiteListFlag *bool `json:"WhiteListFlag,omitempty" name:"WhiteListFlag"`
	// 服务商，可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// 合规化审批单`ID`。

	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *ModifyCrossBorderWhitelistFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderWhitelistFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteReceivingPolicy struct {

	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在地域。

	InstanceRegion []*string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 实例所属的账户uin。

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例类型。如VPC

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type CcnFlowLock struct {

	// 带宽所属的云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例所属用户主账号ID。

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
	// 带宽实例的唯一ID。作为`UnlockCcnBandwidths`接口和`LockCcnBandwidths`接口的入参时，该字段必传。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

type MigrateCcnGatewayRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 源云联网集群ID,&nbsp;形如`vbcgw-12345678`。

	SrcCcnGatewayId *string `json:"SrcCcnGatewayId,omitempty" name:"SrcCcnGatewayId"`
	// 目的云联网集群ID,&nbsp;形如`vbcgw-87654321`。

	DstCcnGatewayId *string `json:"DstCcnGatewayId,omitempty" name:"DstCcnGatewayId"`
}

func (r *MigrateCcnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateCcnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRenewCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRenewCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnFlowLogAttributeRequest struct {
	*tchttp.BaseRequest

	// 是否开启峰值带宽采集。

	PeakBpsSwitch *bool `json:"PeakBpsSwitch,omitempty" name:"PeakBpsSwitch"`
	// 流日志ID。

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *ModifyCcnFlowLogAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnFlowLogAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteBroadcastPolicyRouteCondition struct {

	// 条件类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 条件值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 匹配模式，`1`&nbsp;精确匹配，`0`&nbsp;模糊匹配

	MatchPattern *uint64 `json:"MatchPattern,omitempty" name:"MatchPattern"`
}

type InquiryPriceRenewCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 带宽，单位Gbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 续费标识；NOTIFY_AND_AUTO_RENEW：自动续费；NOTIFY_AND_MANUAL_RENEW：不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
}

func (r *InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnPolicyBasedRoutingRuleAttributeRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 策略路由匹配规则ID

	PolicyBasedRoutingRuleId *string `json:"PolicyBasedRoutingRuleId,omitempty" name:"PolicyBasedRoutingRuleId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCcnPolicyBasedRoutingRuleAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnPolicyBasedRoutingRuleAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件：
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）CCN实例ID。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联实例类型。</li><li>instance-region&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联实例所属地域。</li><li>instance-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 云联网实例ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 排序字段。支持：`CcnId`&nbsp;`InstanceType`&nbsp;`InstanceId`&nbsp;`InstanceName`&nbsp;`InstanceRegion`&nbsp;`AttachedTime`&nbsp;`State`。默认值：`AttachedTime`

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnAggregationRoute struct {

	// 路由CIDR

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// BGP的AS-Path属性

	AsPath *string `json:"AsPath,omitempty" name:"AsPath"`
}

type GetCreateCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）各地域带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
	// 要绑定的资源标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *GetCreateCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnIpv6RoutesRequest struct {
	*tchttp.BaseRequest

	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DisableCcnIpv6RoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnIpv6RoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnPolicyBasedRoutingNextHopRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 关联实例类型[VPC,DIRECTCONNECT,VPNGW]

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 下一跳资源类型[HAVIP,&nbsp;GWLB_ENDPOINT]]

	NextHopResourceType *string `json:"NextHopResourceType,omitempty" name:"NextHopResourceType"`
	// 下一跳地域

	NextHopRegion *string `json:"NextHopRegion,omitempty" name:"NextHopRegion"`
	// 关联实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 下一跳资源ID

	NextHopResourceId *string `json:"NextHopResourceId,omitempty" name:"NextHopResourceId"`
}

func (r *CreateCcnPolicyBasedRoutingNextHopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnPolicyBasedRoutingNextHopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableAssociatedInstancesRequest struct {
	*tchttp.BaseRequest

	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 一次查询最大返回的数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件：
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）CCN实例ID。</li><li>ccn-route-table-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表ID。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）实例类型：
	// 私有网络:&nbsp;`VPC`
	// 专线网关:&nbsp;`DIRECTCONNECT`
	// 黑石私有网络:&nbsp;`BMVPC`
	// EDGE设备:&nbsp;`EDGE`
	// EDGE隧道:&nbsp;`EDGE_TUNNEL`
	// EDGE网关:&nbsp;`EDGE_VPNGW`
	// VPN网关：`VPNGW`</li><li>instance-id-&nbsp;String&nbsp;-（过滤条件）实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRouteTableAssociatedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableAssociatedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderRegionConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线地域配置列表。

		CrossBorderRegionConfigSet []*CrossBorderRegionConfig `json:"CrossBorderRegionConfigSet,omitempty" name:"CrossBorderRegionConfigSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderRegionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderRegionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnFlowLogRequest struct {
	*tchttp.BaseRequest

	// 流日志ID。

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DescribeCcnFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnFlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstance struct {

	// 关联实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 关联实例ID所属大区，例如：region1。

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 关联实例类型，可选值：
	// <li>`VPC`：私有网络</li><li>`DIRECTCONNECT`：专线网关</li><li>`BMVPC`：黑石私有网络</li><li>`VPNGW`：VPNGW类型</li>

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 实例关联的路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 实例付费方式

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

type DirectConnectMap struct {

	// 源地域专线ID。

	SourceDirectConnectId *string `json:"SourceDirectConnectId,omitempty" name:"SourceDirectConnectId"`
	// 目的地域专线ID。

	DestinationDirectConnectId *string `json:"DestinationDirectConnectId,omitempty" name:"DestinationDirectConnectId"`
	// 描述。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type InstanceChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1,&nbsp;2,&nbsp;3,&nbsp;4,&nbsp;5,&nbsp;6,&nbsp;7,&nbsp;8,&nbsp;9,&nbsp;12,&nbsp;24,&nbsp;36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：&nbsp;NOTIFY_AND_AUTO_RENEW：通知过期且自动续费，&nbsp;NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认：NOTIFY_AND_AUTO_RENEW

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type GetTradeBillingAuth struct {

	// 版本号。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 密钥KEY。

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 时间戳。

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 签名。

	Signature *string `json:"Signature,omitempty" name:"Signature"`
	// 每次签名随机字符串。

	Nonce *string `json:"Nonce,omitempty" name:"Nonce"`
}

type CcnInstanceWithoutRegion struct {

	// 云联网支持的实例类型：
	// `VPC`
	// `DIRECTCONNECT`
	// `BMVPC`&nbsp;
	// `EDGE`
	// `EDGE_TUNNEL`
	// `EDGE_VPNGW`
	// `VPNGW`

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CrossBorderCompliance struct {

	// 法定代表人。

	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`
	// 经办人。

	Manager *string `json:"Manager,omitempty" name:"Manager"`
	// 经办人身份证。

	ManagerIdCard *string `json:"ManagerIdCard,omitempty" name:"ManagerIdCard"`
	// 经办人联系电话。

	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`
	// 法定代表人身份证号。

	LegalPersonId *string `json:"LegalPersonId,omitempty" name:"LegalPersonId"`
	// 法定代表人身份证。

	LegalPersonIdCard *string `json:"LegalPersonIdCard,omitempty" name:"LegalPersonIdCard"`
	// 公司全称。

	Company *string `json:"Company,omitempty" name:"Company"`
	// 发证机关。

	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`
	// 营业执照。

	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`
	// 营业执照住所。

	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`
	// 邮编。

	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`
	// 经办人身份证号。

	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
	// 经办人身份证地址。

	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`
	// 电子邮箱。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 服务受理单。

	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitempty" name:"ServiceHandlingForm"`
	// 信息安全承诺书。

	SafetyCommitment *string `json:"SafetyCommitment,omitempty" name:"SafetyCommitment"`
	// 服务开始时间。

	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`
	// 服务截止时间。

	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`
	// 状态。待审批：`PENDING`，已通过：`APPROVED`，已拒绝：`DENY`。

	State *string `json:"State,omitempty" name:"State"`
	// 审批单创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 发起审批的用户主账号ID。

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
	// 合规化审批单`ID`。

	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
	// 授权函。

	AuthorizationLetter *string `json:"AuthorizationLetter,omitempty" name:"AuthorizationLetter"`
	// 服务商，可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// 统一社会信用代码。

	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 关联网络实例列表

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
	// CCN所属UIN（根账号），默认当前账号所属UIN

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCrossBorderWhitelistFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCrossBorderWhitelistFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderWhitelistFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {

	// 路由策略描述。

	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`
	// 路由策略是否发布到云联网。该字段仅做出参使用，作为入参字段时此参数不生效。

	PublishedToVbc *bool `json:"PublishedToVbc,omitempty" name:"PublishedToVbc"`
	// 路由策略创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// CDC&nbsp;集群唯一&nbsp;ID。

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 支持的&nbsp;ECMP算法有：ECMP_QUINTUPLE_HASH：五元组hash，ECMP_SOURCE_DESTINATION_IP_HASH：源和目的IP&nbsp;hash，ECMP_DESTINATION_IP_HASH：目的IP&nbsp;hash，ECMP_SOURCE_IP_HASH：源IP&nbsp;hash。

	SubnetRouteAlgorithm *string `json:"SubnetRouteAlgorithm,omitempty" name:"SubnetRouteAlgorithm"`
	// 是否为&nbsp;Cdc路由

	IsCdc *bool `json:"IsCdc,omitempty" name:"IsCdc"`
	// 创建IPv4目的网段，取值不能在私有网络网段内，例如：112.20.51.0/24。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型，目前我们支持的类型有：
	// CVM：公网网关类型的云服务器；
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// HAVIP：高可用虚拟IP；
	// NAT：公网NAT网关;&nbsp;
	// NORMAL_CVM：普通云服务器；
	// EIP：云服务器的公网IP；
	// LOCAL_GATEWAY：CDC本地网关；
	// INTRANAT：私网NAT网关；
	// USER_CCN：云联网（自定义路由）；
	// GWLB_ENDPOINT：网关负载均衡终端节点。

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 路由策略ID。IPv4路由策略ID是有意义的值，IPv6路由策略是无意义的值0。后续建议完全使用字符串唯一ID&nbsp;`RouteItemId`操作路由策略。
	// 该字段在删除时必填，其他字段无需填写。

	RouteId *uint64 `json:"RouteId,omitempty" name:"RouteId"`
	// 路由类型，目前我们支持的类型有：
	// USER：用户路由；
	// NETD：网络探测路由，创建网络探测实例时，系统默认下发，不可编辑与删除；
	// CCN：云联网路由，系统默认下发，不可编辑与删除。
	// 用户只能添加和操作&nbsp;USER&nbsp;类型的路由。

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 创建IPv6目的网段，取值不能在私有网络网段内，例如：2402:4e00:1000:810b::/64。

	DestinationIpv6CidrBlock *string `json:"DestinationIpv6CidrBlock,omitempty" name:"DestinationIpv6CidrBlock"`
	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由唯一策略ID。

	RouteItemId *string `json:"RouteItemId,omitempty" name:"RouteItemId"`
	// 下一跳地址，这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址。
	// 特殊说明：
	// GatewayType为NORMAL_CVM时，GatewayId填写实例的内网IP。
	// GatewayType为EIP时，GatewayId填写0。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 是否启用

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// local&nbsp;cidr类型：NORMAL/DOCKER。

	AssistantType *string `json:"AssistantType,omitempty" name:"AssistantType"`
	// 高优路由条目标记。

	HighPriorityRouteFlag *bool `json:"HighPriorityRouteFlag,omitempty" name:"HighPriorityRouteFlag"`
}

type LockCcnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LockCcnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnPolicyBasedRoutingRulesRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网策略路由匹配规则

	CcnPolicyBasedRoutingRuleSet []*CcnPolicyBasedRoutingRule `json:"CcnPolicyBasedRoutingRuleSet,omitempty" name:"CcnPolicyBasedRoutingRuleSet"`
}

func (r *CreateCcnPolicyBasedRoutingRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnPolicyBasedRoutingRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeBillingAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AUTH对象。

		Auth *GetTradeBillingAuth `json:"Auth,omitempty" name:"Auth"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTradeBillingAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeBillingAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnAttachedInstance struct {

	// 关联实例类型：
	// <li>`VPC`：私有网络</li><li>`DIRECTCONNECT`：专线网关</li><li>`BMVPC`：黑石私有网络</li>

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 关联时间。

	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
	// 关联实例所属的大地域，如:&nbsp;CHINA_MAINLAND

	InstanceArea *string `json:"InstanceArea,omitempty" name:"InstanceArea"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 路由表名称

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 关联实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 关联实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 关联实例所属大区，例如：region1。

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 关联实例所属UIN（根账号）。

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 关联实例CIDR。

	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 关联实例状态：
	// <li>`PENDING`：申请中</li><li>`ACTIVE`：已连接</li><li>`EXPIRED`：已过期</li><li>`REJECTED`：已拒绝</li><li>`DELETED`：已删除</li><li>`FAILED`：失败的（2小时后将异步强制解关联）</li><li>`ATTACHING`：关联中</li><li>`DETACHING`：解关联中</li><li>`DETACHFAILED`：解关联失败（2小时后将异步强制解关联）</li>

	State *string `json:"State,omitempty" name:"State"`
	// 云联网所属UIN（根账号）。

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
	// 路由表ID

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 云联网实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单信息。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionDefaultQosBandwidthLimitRequest struct {
	*tchttp.BaseRequest

	// 服务等级；

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// CCN实例ID，形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 原地域；

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域；

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
}

func (r *SetCcnRegionDefaultQosBandwidthLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionDefaultQosBandwidthLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CCN对象。

		CcnSet []*CCN `json:"CcnSet,omitempty" name:"CcnSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTable struct {

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID。

	CcnRouteTableId *string `json:"CcnRouteTableId,omitempty" name:"CcnRouteTableId"`
	// 云联网路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 云联网路由表描述。

	RouteTableDescription *string `json:"RouteTableDescription,omitempty" name:"RouteTableDescription"`
	// True：是默认路由表&nbsp;False：非默认路由表。

	IsDefaultTable *bool `json:"IsDefaultTable,omitempty" name:"IsDefaultTable"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type InquirePriceCreateCcnInstanceRequest struct {
	*tchttp.BaseRequest

	// ccn实例信息

	CcnInstanceChargeInfo []*CcnInstance `json:"CcnInstanceChargeInfo,omitempty" name:"CcnInstanceChargeInfo"`
}

func (r *InquirePriceCreateCcnInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateCcnInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnLimit struct {

	// 云联网配额类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 云联网配额数值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DescribeCcnRouteTableInputPolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由接收策略版本号。

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DescribeCcnRouteTableInputPolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableInputPolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnFlowLogAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnFlowLogAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnFlowLogAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest

	// （模糊查询）经办人。

	Manager *string `json:"Manager,omitempty" name:"Manager"`
	// （精确匹配）服务开始日期，如：`2020-07-28`。

	ServiceStartDate *string `json:"ServiceStartDate,omitempty" name:"ServiceStartDate"`
	// （精确匹配）状态。待审批：`PENDING`，通过：`APPROVED&nbsp;`，拒绝：`DENY`。

	State *string `json:"State,omitempty" name:"State"`
	// （精确匹配）统一社会信用代码。

	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`
	// （精确匹配）经办人联系电话。

	ManagerTelephone *string `json:"ManagerTelephone,omitempty" name:"ManagerTelephone"`
	// （精确匹配）合规化审批单`ID`。

	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
	// （模糊查询）公司名称。

	Company *string `json:"Company,omitempty" name:"Company"`
	// （模糊查询）营业执照住所。

	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`
	// （精确匹配）电子邮箱。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// （精确匹配）服务商，可选值：`UNICOM`。

	ServiceProvider *string `json:"ServiceProvider,omitempty" name:"ServiceProvider"`
	// （模糊查询）发证机关。

	IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`
	// （精确匹配）邮编。

	PostCode *uint64 `json:"PostCode,omitempty" name:"PostCode"`
	// （精确查询）经办人身份证号。

	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
	// （模糊查询）经办人身份证地址。

	ManagerAddress *string `json:"ManagerAddress,omitempty" name:"ManagerAddress"`
	// （精确匹配）服务结束日期，如：`2021-07-28`。

	ServiceEndDate *string `json:"ServiceEndDate,omitempty" name:"ServiceEndDate"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// （模糊查询）法定代表人。

	LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`
	// （精确查询）法人身份证号。

	LegalPersonId *string `json:"LegalPersonId,omitempty" name:"LegalPersonId"`
}

func (r *DescribeCrossBorderComplianceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderComplianceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditCrossBorderComplianceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstancePrice struct {

	// 折扣单价

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 数量

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 所属地区

	Location *string `json:"Location,omitempty" name:"Location"`
	// UnitPrice

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 单位

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 总价

	TotalPrice *float64 `json:"TotalPrice,omitempty" name:"TotalPrice"`
	// 原始单价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type SetCcnDefaultBillingModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnDefaultBillingModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnDefaultBillingModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteMatchRuleRequest struct {
	*tchttp.BaseRequest

	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 当前地域。

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 过滤条件。&nbsp;src-cidr-&nbsp;String&nbsp;-&nbsp;（过滤条件）源&nbsp;IP，形如：`10.0.0.0/8`。&nbsp;dst-cidr&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）目的&nbsp;IP，形如：`11.0.0.0/8`。&nbsp;src-port&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）源端口。&nbsp;dst-port&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）目的端口。&nbsp;protocol&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）协议。&nbsp;QoS&nbsp;ID&nbsp;-&nbsp;String&nbsp;-&nbsp;（过滤条件）QoS&nbsp;ID。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCcnRouteMatchRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteMatchRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectAccelerateChannelRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道唯一ID，形如dcac-dy8akq6c。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
}

func (r *DeleteDirectConnectAccelerateChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectAccelerateChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantCcnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）对象。

		CcnSet []*CcnInstanceInfo `json:"CcnSet,omitempty" name:"CcnSet"`
		// 符合条件的对象总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantCcnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCcnRouteTableBroadcastPolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 新的路由传播策略

	Policys []*CcnRouteTableBroadcastPolicy `json:"Policys" name:"Policys"`
}

func (r *ReplaceCcnRouteTableBroadcastPolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableBroadcastPolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderFlowMonitorRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网所属账号。

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
	// 时间粒度。单位为:秒，如60为60s的时间粒度

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 源地域。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
}

func (r *DescribeCrossBorderFlowMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderFlowMonitorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnBandwidthRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 带宽配置ID。

	RegionFlowControlIds []*string `json:"RegionFlowControlIds,omitempty" name:"RegionFlowControlIds"`
	// 是否自动续费标识，1：自动续费；2：不自动续费

	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 是否自动续费标识，NOTIFY_AND_AUTO_RENEW：自动续费；NOTIFY_AND_MANUAL_RENEW：&nbsp;不自动续费

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *SetCcnBandwidthRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnBandwidthRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRegionBandwidthLimitInfo struct {

	// 源地域，例如：region1

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 目的地域，&nbsp;例如：region1

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 出带宽上限，单位：Mbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
}

type DescribeTrafficQosPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量调度规则。

		TrafficQosPolicySet []*TrafficQosPolicySet `json:"TrafficQosPolicySet,omitempty" name:"TrafficQosPolicySet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficQosPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficQosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectPortMapRequest struct {
	*tchttp.BaseRequest

	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
}

func (r *DescribeDirectConnectPortMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectPortMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTableInputPolicys struct {

	// 策略列表。

	Policys []*CcnRouteTableInputPolicy `json:"Policys,omitempty" name:"Policys"`
	// 版本号。

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CcnRouteMatchRuleSet struct {

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 路由匹配规则唯一ID。

	CcnRouteMatchRuleId *string `json:"CcnRouteMatchRuleId,omitempty" name:"CcnRouteMatchRuleId"`
	// 源网段。

	SrcCidr *string `json:"SrcCidr,omitempty" name:"SrcCidr"`
	// 源端口。

	SrcPort *string `json:"SrcPort,omitempty" name:"SrcPort"`
	// 协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// qos&nbsp;id。

	QosId *uint64 `json:"QosId,omitempty" name:"QosId"`
	// 优先级。

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 当前地域。

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 目的网段。

	DstCidr *string `json:"DstCidr,omitempty" name:"DstCidr"`
	// 目的端口。

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 备注。

	CcnRouteMatchRuleDescription *string `json:"CcnRouteMatchRuleDescription,omitempty" name:"CcnRouteMatchRuleDescription"`
}

type InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 目标升级带宽，单位Gbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
}

func (r *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnPolicyBasedRoutingNextHopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnPolicyBasedRoutingNextHopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnPolicyBasedRoutingNextHopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批单。

		CrossBorderCompliance *CrossBorderCompliance `json:"CrossBorderCompliance,omitempty" name:"CrossBorderCompliance"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCrossBorderComplianceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteMatchRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例总个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由匹配规则

		CcnRouteMatchRuleSet []*CcnRouteMatchRuleSet `json:"CcnRouteMatchRuleSet,omitempty" name:"CcnRouteMatchRuleSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRouteMatchRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteMatchRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnPolicyBasedRoutingNextHopAttributeRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 策略路由下一跳ID

	PolicyBasedRoutingNextHopId *string `json:"PolicyBasedRoutingNextHopId,omitempty" name:"PolicyBasedRoutingNextHopId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 下一跳地域

	NextHopRegion *string `json:"NextHopRegion,omitempty" name:"NextHopRegion"`
	// 下一跳资源类型[HAVIP,GWLB_ENDPOINT]

	NextHopResourceType *string `json:"NextHopResourceType,omitempty" name:"NextHopResourceType"`
	// 下一跳资源ID

	NextHopResourceId *string `json:"NextHopResourceId,omitempty" name:"NextHopResourceId"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例类型[VPC,DIRECTCONNECT,VPNGW]

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 关联实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyCcnPolicyBasedRoutingNextHopAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnPolicyBasedRoutingNextHopAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteReceivingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件：
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）CCN实例ID。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联实例类型:
	// 私有网络:&nbsp;`VPC`
	// 专线网关:&nbsp;`DIRECTCONNECT`
	// 黑石私有网络:&nbsp;`BMVPC`
	// EDGE设备:&nbsp;`EDGE`
	// EDGE隧道:&nbsp;`EDGE_TUNNEL`
	// EDGE网关:&nbsp;`EDGE_VPNGW`
	// VPN网关：`VPNGW`</li><li>ccn-route-table-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表ID。</li><li>instance-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联的实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRouteReceivingPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteReceivingPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网配额列表

		CcnLimitSet []*CcnLimit `json:"CcnLimitSet,omitempty" name:"CcnLimitSet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 接受关联实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *AcceptAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteReceivingPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteReceivingPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteReceivingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnBatchRouteTable struct {

	// 云联网路由表名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云联网路由表描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type CrossBorderSettlement struct {

	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 目的地域

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 跨境账单ID

	SettlementId *string `json:"SettlementId,omitempty" name:"SettlementId"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 销售经理

	SalesManager *string `json:"SalesManager,omitempty" name:"SalesManager"`
	// CCN&nbsp;ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 源地域

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 用量

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 折后总价

	TotalAmount *string `json:"TotalAmount,omitempty" name:"TotalAmount"`
	// 用户UIN

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
	// CCN名称

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// 刊例价

	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 折扣单价

	DiscountPrice *string `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 账单月份

	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
	// AU

	ServiceLevel *string `json:"ServiceLevel,omitempty" name:"ServiceLevel"`
}

type CcnPolicyBasedRoutingNextHop struct {

	// 0

	PolicyBasedRoutingRulesCount *int64 `json:"PolicyBasedRoutingRulesCount,omitempty" name:"PolicyBasedRoutingRulesCount"`
	// 下一跳地域

	NextHopRegion *string `json:"NextHopRegion,omitempty" name:"NextHopRegion"`
	// 关联实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 关联实例类型[VPC,DIRECTCONNECT,VPNGW]

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 下一跳资源ip

	NextHopIp *string `json:"NextHopIp,omitempty" name:"NextHopIp"`
	// 策略路由下一跳ID

	PolicyBasedRoutingNextHopId *string `json:"PolicyBasedRoutingNextHopId,omitempty" name:"PolicyBasedRoutingNextHopId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态(ENABLE/DISABLE)

	State *string `json:"State,omitempty" name:"State"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 下一跳资源类型[HAVIP,GWLB_ENDPOINT]

	NextHopResourceType *string `json:"NextHopResourceType,omitempty" name:"NextHopResourceType"`
	// 下一跳资源ID

	NextHopResourceId *string `json:"NextHopResourceId,omitempty" name:"NextHopResourceId"`
}

type GetRenewDirectConnectAccelerateChannelBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单信息。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRenewDirectConnectAccelerateChannelBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewDirectConnectAccelerateChannelBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTableInputPolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表接收策略。

		PolicySet []*CcnRouteTableInputPolicys `json:"PolicySet,omitempty" name:"PolicySet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRouteTableInputPolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableInputPolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *InquiryPriceRenewCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRouteTableBroadcastPolicy struct {

	// 传播条件

	BroadcastConditions []*CcnRouteBroadcastPolicyRouteCondition `json:"BroadcastConditions,omitempty" name:"BroadcastConditions"`
	// 路由行为，`accept`&nbsp;允许，`drop`&nbsp;拒绝

	Action *string `json:"Action,omitempty" name:"Action"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// as-path操作

	OperateAsPath *string `json:"OperateAsPath,omitempty" name:"OperateAsPath"`
	// as-path操作模式

	AsPathOperateMode *string `json:"AsPathOperateMode,omitempty" name:"AsPathOperateMode"`
	// community操作

	OperateCommunitySet []*string `json:"OperateCommunitySet,omitempty" name:"OperateCommunitySet"`
	// community操作模式

	CommunityOperateMode *string `json:"CommunityOperateMode,omitempty" name:"CommunityOperateMode"`
	// 路由条件

	RouteConditions []*CcnRouteBroadcastPolicyRouteCondition `json:"RouteConditions,omitempty" name:"RouteConditions"`
}

type DescribeCrossBorderCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，目前`value`值个数只支持一个，可支持的字段有：
	// <li>`source-region`&nbsp;源地域，值形如：`["region1"]`</li>&nbsp;<li>`destination-region`&nbsp;目的地域，值形如：`["region1"]`</li>&nbsp;<li>`ccn-ids`&nbsp;云联网ID数组，值形如：`["ccn-12345678"]`</li>&nbsp;<li>`user-account-id`&nbsp;用户账号ID，值形如`["12345678"]`</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单页返回数据量可选值0到100之间的整数，默认20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnPolicyBasedRoutingNextHopRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 策略路由下一跳ID

	PolicyBasedRoutingNextHopIds []*string `json:"PolicyBasedRoutingNextHopIds,omitempty" name:"PolicyBasedRoutingNextHopIds"`
}

func (r *DeleteCcnPolicyBasedRoutingNextHopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnPolicyBasedRoutingNextHopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽详情。

		CcnBandwidthSet []*CcnBandwidthInfo `json:"CcnBandwidthSet,omitempty" name:"CcnBandwidthSet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCcnRouteTableInputPolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 新的路由接收策略。

	Policys []*CcnRouteTableInputPolicy `json:"Policys" name:"Policys"`
}

func (r *ReplaceCcnRouteTableInputPolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableInputPolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableSelectionPoliciesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件：
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）CCN实例ID。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联实例类型:
	// 私有网络:&nbsp;`VPC`
	// 专线网关:&nbsp;`DIRECTCONNECT`
	// 黑石私有网络:&nbsp;`BMVPC`
	// EDGE设备:&nbsp;`EDGE`
	// EDGE隧道:&nbsp;`EDGE_TUNNEL`
	// EDGE网关:&nbsp;`EDGE_VPNGW`
	// VPN网关：`VPNGW`</li><li>ccn-route-table-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表ID。</li><li>instance-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）关联的实例ID。</li><li>route-table-name&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表名称。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRouteTableSelectionPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableSelectionPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批单。

		CrossBorderCompliance *CrossBorderCompliance `json:"CrossBorderCompliance,omitempty" name:"CrossBorderCompliance"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCrossBorderComplianceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 属性名称,&nbsp;若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值,&nbsp;若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。当值类型为布尔类型时，可直接取值为字符串"TRUE"或&nbsp;"FALSE"。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteTrafficQosPolicyRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 流量调度策略唯一ID。

	QosPolicyId *string `json:"QosPolicyId,omitempty" name:"QosPolicyId"`
}

func (r *DeleteTrafficQosPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficQosPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnBandwidth struct {

	// 带宽是否自动续费的标记。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 云市场实例ID。

	MarketId *string `json:"MarketId,omitempty" name:"MarketId"`
	// 实例所属用户主账号ID。

	UserAccountID *string `json:"UserAccountID,omitempty" name:"UserAccountID"`
	// 是否跨境，`true`表示跨境，反之不跨境。

	IsCrossBorder *bool `json:"IsCrossBorder,omitempty" name:"IsCrossBorder"`
	// `true`表示封禁，地域间流量不通，`false`解禁，地域间流量正常

	IsSecurityLock *bool `json:"IsSecurityLock,omitempty" name:"IsSecurityLock"`
	// `POSTPAID`表示后付费，`PREPAID`表示预付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 带宽所属的云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例的过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 描述带宽的地域和限速上限信息。在地域间限速的情况下才会返回参数，出口限速模式不返回。

	CcnRegionBandwidthLimit *CcnRegionBandwidthLimitInfo `json:"CcnRegionBandwidthLimit,omitempty" name:"CcnRegionBandwidthLimit"`
	// 实例更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 公司名称。

	Company *string `json:"Company,omitempty" name:"Company"`
	// 实例的创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 带宽实例的唯一ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

type ModifyCcnRoutePriorityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRoutePriorityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRoutePriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID，形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网（CCN）各地域出带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
	// 是否恢复云联网地域出口/地域间带宽限速为默认值（1Gbps）。false表示不恢复；true表示恢复。恢复默认值后，限速实例将不在控制台展示。该参数默认为&nbsp;false，不恢复。

	SetDefaultLimitFlag *bool `json:"SetDefaultLimitFlag,omitempty" name:"SetDefaultLimitFlag"`
}

func (r *SetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRegionFixedBandwidthLimit struct {

	// 地域，例如：region1。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 出带宽上限，单位：Mbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 地域，例如：region1。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

type SetCcnRegionFixedBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID，形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网（CCN）各地域出带宽上限。

	CcnRegionFixedBandwidthLimit []*CcnRegionFixedBandwidthLimit `json:"CcnRegionFixedBandwidthLimit,omitempty" name:"CcnRegionFixedBandwidthLimit"`
}

func (r *SetCcnRegionFixedBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionFixedBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficQosPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrafficQosPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficQosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableInfo struct {

	// 云联网路由表id。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 云联网路由表名称。Name和Description&nbsp;两者必传一个。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云联网路由表描述。Name和Description&nbsp;两者必传一个。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type ReplaceCcnRouteTableBroadcastPolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCcnRouteTableBroadcastPolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableBroadcastPolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 加速通道实例ID。

	DirectConnectAccelerateChannelIds []*string `json:"DirectConnectAccelerateChannelIds,omitempty" name:"DirectConnectAccelerateChannelIds"`
	// NOTIFY_AND_AUTO_RENEW：自动续费；NOTIFY_AND_MANUAL_RENEW：不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表信息列表。

		CcnRouteTableSet []*CcnRouteTable `json:"CcnRouteTableSet,omitempty" name:"CcnRouteTableSet"`
		// 查询到的路由表数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 需要修改的路由表列表。

	RouteTableInfo []*ModifyRouteTableInfo `json:"RouteTableInfo,omitempty" name:"RouteTableInfo"`
}

func (r *ModifyCcnRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnLimitsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// type-&nbsp;Int&nbsp;-（过滤条件）云联网配额类型。
	// 1：每个开发商可创建的云联网数，2：每个云联网绑定的实例个数，3：单个云联网支持的路由条目。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCcnLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemPrice struct {

	// 按量计费后付费单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 按量计费后付费计价单元，可取值范围：&nbsp;HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：&nbsp;GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预付费商品的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预付费商品的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type InquiryPriceUpdateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpdateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceUpdateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnlockCcnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnlockCcnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRegionBandwidthLimit struct {

	// 地域，例如：region1

	Region *string `json:"Region,omitempty" name:"Region"`
	// 出带宽上限，单位：Mbps

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 是否黑石地域，默认`false`。

	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`
	// 目的地域，例如：region1

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 目的地域是否为黑石地域，默认`false`。

	DstIsBm *bool `json:"DstIsBm,omitempty" name:"DstIsBm"`
	// 服务等级

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
}

type GetCcnTrafficRemainAmountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量剩余额度（单位：MB）

		TrafficRemainAmount *float64 `json:"TrafficRemainAmount,omitempty" name:"TrafficRemainAmount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCcnTrafficRemainAmountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnTrafficRemainAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRouteMatchRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRouteMatchRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRouteMatchRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteReceivingPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteReceivingPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteReceivingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableSelectionPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由表选择策略信息集合。

		RouteSelectionPolicySet []*RouteSelectionPolicy `json:"RouteSelectionPolicySet,omitempty" name:"RouteSelectionPolicySet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTableSelectionPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableSelectionPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 合规化审批单列表。

		CrossBorderComplianceSet []*CrossBorderCompliance `json:"CrossBorderComplianceSet,omitempty" name:"CrossBorderComplianceSet"`
		// 合规化审批单总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCrossBorderComplianceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 需要删除的路由表列表。

	RouteTableId []*string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteCcnRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCreateDirectConnectAccelerateChannelBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单信息。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCreateDirectConnectAccelerateChannelBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateDirectConnectAccelerateChannelBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRoutePriorityRequest struct {
	*tchttp.BaseRequest

	// CCN路由策略唯一ID。形如：`18292`。

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 路由优先级值。取值范围0-100。

	RoutePriority *int64 `json:"RoutePriority,omitempty" name:"RoutePriority"`
	// CCN实例ID。形如：`ccn-f49l6u0z`。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *ModifyCcnRoutePriorityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRoutePriorityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID，形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthLimitForCcnAlarmOnly struct {

	// 地域，例如：region1。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 出带宽上限，单位：Mbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 整型ID。

	VbcId *uint64 `json:"VbcId,omitempty" name:"VbcId"`
	// 实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type ReplaceCcnRouteTableAggregatePolicysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCcnRouteTableAggregatePolicysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCcnRouteTableAggregatePolicysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnIpv6RoutesRequest struct {
	*tchttp.BaseRequest

	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// CCN实例ID，形如：`ccn-gree226l。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID，形如：`ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
	// 过滤条件，参数不支持同时指定RouteIds和Filters。
	// <li>route-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由策略ID。</li><li>cidr-block&nbsp;-&nbsp;String&nbsp;-（过滤条件）目的端。</li><li>instance-type&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳类型。</li><li>instance-region&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳所属地域。</li><li>instance-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）下一跳实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCcnIpv6RoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnIpv6RoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthLimitForCcnAlarmOnlyRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *BandwidthLimitForCcnAlarmOnlyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BandwidthLimitForCcnAlarmOnlyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstanceNonDirectFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 非直通标记

		NonDirectFlag *bool `json:"NonDirectFlag,omitempty" name:"NonDirectFlag"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnInstanceNonDirectFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnInstanceNonDirectFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnRequest struct {
	*tchttp.BaseRequest

	// 指定绑定的标签列表，例如：[{"Key":&nbsp;"city",&nbsp;"Value":&nbsp;"region1"}]

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// CCN名称，最大长度不能超过60个字节。

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// CCN描述信息，最大长度不能超过100个字节。

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// CCN服务质量，`PT`：白金，`AU`：金，`AG`：银，默认为`AU`。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 计费模式，`PREPAID`：表示预付费，即包年包月，`POSTPAID`：表示后付费，即按量计费。默认：`POSTPAID`。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 计量模式

	InstanceMeteringType *string `json:"InstanceMeteringType,omitempty" name:"InstanceMeteringType"`
	// 限速类型，`OUTER_REGION_LIMIT`表示地域出口限速，`INTER_REGION_LIMIT`为地域间限速，默认为`OUTER_REGION_LIMIT`。预付费模式仅支持地域间限速，后付费模式支持地域间限速和地域出口限速。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

func (r *CreateCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数配置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 带宽配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *RenewCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnBandwidthConfigRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：`ccn-f49l6u0z`。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 需要配置的地域带宽的地域对信息

	SourceDestinationRegions []*SourceDestinationRegion `json:"SourceDestinationRegions,omitempty" name:"SourceDestinationRegions"`
}

func (r *DescribeCcnBandwidthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnBandwidthConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectAccelerateChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按照条件查询出来的加速通道数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 加速通道信息。

		DirectConnectAccelerateChannelSet []*DirectConnectAccelerateChannelSet `json:"DirectConnectAccelerateChannelSet,omitempty" name:"DirectConnectAccelerateChannelSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectAccelerateChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectAccelerateChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateCcnInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总价

		TotalPrice *float64 `json:"TotalPrice,omitempty" name:"TotalPrice"`
		// 流量价格

		TrafficPrice []*CcnTrafficPrice `json:"TrafficPrice,omitempty" name:"TrafficPrice"`
		// 实例价格

		InstancePrice []*CcnInstancePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateCcnInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceCreateCcnInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnDefaultBillingModeRequest struct {
	*tchttp.BaseRequest

	// 地域信息。

	CcnBandwidthRegion []*CcnBandwidthRegion `json:"CcnBandwidthRegion,omitempty" name:"CcnBandwidthRegion"`
	// CCN实例ID，形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *SetCcnDefaultBillingModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnDefaultBillingModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnIpv6RoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCcnIpv6RoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnIpv6RoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// <li>ccn-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）CCN实例ID。</li><li>route-table-id&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表ID。</li><li>route-table-name&nbsp;-&nbsp;String&nbsp;-（过滤条件）路由表名称。</li><li>route-table-description-&nbsp;String&nbsp;-（过滤条件）路由表备注。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 一次查询最大返回的数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRefundCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// 云联网带宽唯一ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *GetRefundCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefundCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableSelectionPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteTableSelectionPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableSelectionPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionFixedBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnRegionFixedBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionFixedBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeBillingAuthRequest struct {
	*tchttp.BaseRequest

	// 预付费商品信息，需为JSON字符串。

	Goods *string `json:"Goods,omitempty" name:"Goods"`
}

func (r *GetTradeBillingAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeBillingAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 需要创建的路由表列表。

	RouteTable []*CcnBatchRouteTable `json:"RouteTable,omitempty" name:"RouteTable"`
}

func (r *CreateCcnRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearRouteTableSelectionPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearRouteTableSelectionPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearRouteTableSelectionPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCcnTrafficRemainAmountRequest struct {
	*tchttp.BaseRequest

	// 所属月份

	CurrentMonth *string `json:"CurrentMonth,omitempty" name:"CurrentMonth"`
}

func (r *GetCcnTrafficRemainAmountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnTrafficRemainAmountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectPortMapRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-2gx0nr9r。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 加速通道实例ID。

	DirectConnectAccelerateChannelId *string `json:"DirectConnectAccelerateChannelId,omitempty" name:"DirectConnectAccelerateChannelId"`
	// 端口信息

	DirectConnectMap []*DirectConnectMap `json:"DirectConnectMap,omitempty" name:"DirectConnectMap"`
}

func (r *ReplaceDirectConnectPortMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDirectConnectPortMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnInstanceNonDirectFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCcnInstanceNonDirectFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnInstanceNonDirectFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnlockCcnBandwidthsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnlockCcnBandwidthsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockCcnBandwidthsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagSet struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetUpdateCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 地域间设置带宽，单位：Mbps。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *GetUpdateCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRouteTableBroadcastPolicysRequest struct {
	*tchttp.BaseRequest

	// 云联网ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由表ID

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由传播策略版本号

	PolicyVersion *uint64 `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DescribeCcnRouteTableBroadcastPolicysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRouteTableBroadcastPolicysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnIpv6RoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableCcnIpv6RoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnIpv6RoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 要解关联网络实例列表

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnBandwidthInfo struct {

	// 实例的过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 带宽是否自动续费的标记。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 资源绑定的标签列表

	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
	// `true表示`Qos默认带宽；`false`表示非Qos默认带宽；

	DefaultQosBandwidthFlag *bool `json:"DefaultQosBandwidthFlag,omitempty" name:"DefaultQosBandwidthFlag"`
	// 服务等级信息。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 带宽实例的唯一ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 描述带宽的地域和限速上限信息。在地域间限速的情况下才会返回参数，出口限速模式不返回。

	CcnRegionBandwidthLimit *CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimit,omitempty" name:"CcnRegionBandwidthLimit"`
	// 云市场实例ID。

	MarketId *string `json:"MarketId,omitempty" name:"MarketId"`
	// `true`表示封禁，地域间流量不通，`false`解禁，地域间流量正常

	IsSecurityLock *bool `json:"IsSecurityLock,omitempty" name:"IsSecurityLock"`
	// `ISOLATED`表示隔离，`AVAILABLE`表示可用。

	State *string `json:"State,omitempty" name:"State"`
	// 带宽所属的云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例的创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DeleteCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

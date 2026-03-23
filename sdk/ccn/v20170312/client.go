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
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2017-03-12"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewModifyCcnPolicyBasedRoutingNextHopAttributeRequest() (request *ModifyCcnPolicyBasedRoutingNextHopAttributeRequest) {
	request = &ModifyCcnPolicyBasedRoutingNextHopAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnPolicyBasedRoutingNextHopAttribute")
	return
}

func NewModifyCcnPolicyBasedRoutingNextHopAttributeResponse() (response *ModifyCcnPolicyBasedRoutingNextHopAttributeResponse) {
	response = &ModifyCcnPolicyBasedRoutingNextHopAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新云联网策略路由下一跳参数
func (c *Client) ModifyCcnPolicyBasedRoutingNextHopAttribute(request *ModifyCcnPolicyBasedRoutingNextHopAttributeRequest) (response *ModifyCcnPolicyBasedRoutingNextHopAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCcnPolicyBasedRoutingNextHopAttributeRequest()
	}
	response = NewModifyCcnPolicyBasedRoutingNextHopAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteTableAssociatedInstancesRequest() (request *DescribeRouteTableAssociatedInstancesRequest) {
	request = &DescribeRouteTableAssociatedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeRouteTableAssociatedInstances")
	return
}

func NewDescribeRouteTableAssociatedInstancesResponse() (response *DescribeRouteTableAssociatedInstancesResponse) {
	response = &DescribeRouteTableAssociatedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRouteTableAssociatedInstances）用于查询指定的云联网关联的实例所绑定的路由表信息。
func (c *Client) DescribeRouteTableAssociatedInstances(request *DescribeRouteTableAssociatedInstancesRequest) (response *DescribeRouteTableAssociatedInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeRouteTableAssociatedInstancesRequest()
	}
	response = NewDescribeRouteTableAssociatedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnPolicyBasedRoutingRuleRequest() (request *DeleteCcnPolicyBasedRoutingRuleRequest) {
	request = &DeleteCcnPolicyBasedRoutingRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcnPolicyBasedRoutingRule")
	return
}

func NewDeleteCcnPolicyBasedRoutingRuleResponse() (response *DeleteCcnPolicyBasedRoutingRuleResponse) {
	response = &DeleteCcnPolicyBasedRoutingRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云联网策略路由匹配规则
func (c *Client) DeleteCcnPolicyBasedRoutingRule(request *DeleteCcnPolicyBasedRoutingRuleRequest) (response *DeleteCcnPolicyBasedRoutingRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCcnPolicyBasedRoutingRuleRequest()
	}
	response = NewDeleteCcnPolicyBasedRoutingRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetCreateCcnBandwidthDealRequest() (request *GetCreateCcnBandwidthDealRequest) {
	request = &GetCreateCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetCreateCcnBandwidthDeal")
	return
}

func NewGetCreateCcnBandwidthDealResponse() (response *GetCreateCcnBandwidthDealResponse) {
	response = &GetCreateCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetCreateCcnBandwidthDeal）用于获取创建云联网带宽的商品信息。
func (c *Client) GetCreateCcnBandwidthDeal(request *GetCreateCcnBandwidthDealRequest) (response *GetCreateCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetCreateCcnBandwidthDealRequest()
	}
	response = NewGetCreateCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewGetCreateDirectConnectAccelerateChannelBandwidthDealRequest() (request *GetCreateDirectConnectAccelerateChannelBandwidthDealRequest) {
	request = &GetCreateDirectConnectAccelerateChannelBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetCreateDirectConnectAccelerateChannelBandwidthDeal")
	return
}

func NewGetCreateDirectConnectAccelerateChannelBandwidthDealResponse() (response *GetCreateDirectConnectAccelerateChannelBandwidthDealResponse) {
	response = &GetCreateDirectConnectAccelerateChannelBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 二层CCN加速通道购买带宽-前端获取订单参数
func (c *Client) GetCreateDirectConnectAccelerateChannelBandwidthDeal(request *GetCreateDirectConnectAccelerateChannelBandwidthDealRequest) (response *GetCreateDirectConnectAccelerateChannelBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetCreateDirectConnectAccelerateChannelBandwidthDealRequest()
	}
	response = NewGetCreateDirectConnectAccelerateChannelBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewGetCcnRegionBandwidthLimitsRequest() (request *GetCcnRegionBandwidthLimitsRequest) {
	request = &GetCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetCcnRegionBandwidthLimits")
	return
}

func NewGetCcnRegionBandwidthLimitsResponse() (response *GetCcnRegionBandwidthLimitsResponse) {
	response = &GetCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetCcnRegionBandwidthLimits）用于查询云联网相关地域带宽信息，其中预付费模式的云联网仅支持地域间限速，后付费模式的云联网支持地域间限速和地域出口限速。
func (c *Client) GetCcnRegionBandwidthLimits(request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewGetCcnRegionBandwidthLimitsRequest()
	}
	response = NewGetCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewGetRefundCcnBandwidthDealRequest() (request *GetRefundCcnBandwidthDealRequest) {
	request = &GetRefundCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetRefundCcnBandwidthDeal")
	return
}

func NewGetRefundCcnBandwidthDealResponse() (response *GetRefundCcnBandwidthDealResponse) {
	response = &GetRefundCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退费云联网地域间带宽获取订单参数
func (c *Client) GetRefundCcnBandwidthDeal(request *GetRefundCcnBandwidthDealRequest) (response *GetRefundCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetRefundCcnBandwidthDealRequest()
	}
	response = NewGetRefundCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewResetAttachCcnInstancesRequest() (request *ResetAttachCcnInstancesRequest) {
	request = &ResetAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ResetAttachCcnInstances")
	return
}

func NewResetAttachCcnInstancesResponse() (response *ResetAttachCcnInstancesResponse) {
	response = &ResetAttachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetAttachCcnInstances）用于跨账号关联实例申请过期时，重新申请关联操作。
func (c *Client) ResetAttachCcnInstances(request *ResetAttachCcnInstancesRequest) (response *ResetAttachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewResetAttachCcnInstancesRequest()
	}
	response = NewResetAttachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnAttributeRequest() (request *ModifyCcnAttributeRequest) {
	request = &ModifyCcnAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnAttribute")
	return
}

func NewModifyCcnAttributeResponse() (response *ModifyCcnAttributeResponse) {
	response = &ModifyCcnAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCcnAttribute）用于修改云联网（CCN）的相关属性。
func (c *Client) ModifyCcnAttribute(request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCcnAttributeRequest()
	}
	response = NewModifyCcnAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewAttachCcnInstancesRequest() (request *AttachCcnInstancesRequest) {
	request = &AttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "AttachCcnInstances")
	return
}

func NewAttachCcnInstancesResponse() (response *AttachCcnInstancesResponse) {
	response = &AttachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachCcnInstances）用于将网络实例加载到云联网实例中，网络实例包括VPC和专线网关。<br />
// 每个云联网能够关联的网络实例个数是有限的，详情请参考产品文档。如果需要扩充请联系在线客服。
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewAttachCcnInstancesRequest()
	}
	response = NewAttachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCrossBorderWhitelistFlagRequest() (request *ModifyCrossBorderWhitelistFlagRequest) {
	request = &ModifyCrossBorderWhitelistFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCrossBorderWhitelistFlag")
	return
}

func NewModifyCrossBorderWhitelistFlagResponse() (response *ModifyCrossBorderWhitelistFlagResponse) {
	response = &ModifyCrossBorderWhitelistFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改跨境白名单开关
func (c *Client) ModifyCrossBorderWhitelistFlag(request *ModifyCrossBorderWhitelistFlagRequest) (response *ModifyCrossBorderWhitelistFlagResponse, err error) {
	if request == nil {
		request = NewModifyCrossBorderWhitelistFlagRequest()
	}
	response = NewModifyCrossBorderWhitelistFlagResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceUpdateCcnBandwidthRequest() (request *InquiryPriceUpdateCcnBandwidthRequest) {
	request = &InquiryPriceUpdateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquiryPriceUpdateCcnBandwidth")
	return
}

func NewInquiryPriceUpdateCcnBandwidthResponse() (response *InquiryPriceUpdateCcnBandwidthResponse) {
	response = &InquiryPriceUpdateCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceUpdateCcnBandwidth）用于修改预付费云联网实例地域间带宽时的询价
func (c *Client) InquiryPriceUpdateCcnBandwidth(request *InquiryPriceUpdateCcnBandwidthRequest) (response *InquiryPriceUpdateCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceUpdateCcnBandwidthRequest()
	}
	response = NewInquiryPriceUpdateCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCrossBorderSettlementAttributeRequest() (request *ModifyCrossBorderSettlementAttributeRequest) {
	request = &ModifyCrossBorderSettlementAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCrossBorderSettlementAttribute")
	return
}

func NewModifyCrossBorderSettlementAttributeResponse() (response *ModifyCrossBorderSettlementAttributeResponse) {
	response = &ModifyCrossBorderSettlementAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改跨境账单参数-联通专用
func (c *Client) ModifyCrossBorderSettlementAttribute(request *ModifyCrossBorderSettlementAttributeRequest) (response *ModifyCrossBorderSettlementAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCrossBorderSettlementAttributeRequest()
	}
	response = NewModifyCrossBorderSettlementAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRenewCcnBandwidthRequest() (request *RenewCcnBandwidthRequest) {
	request = &RenewCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "RenewCcnBandwidth")
	return
}

func NewRenewCcnBandwidthResponse() (response *RenewCcnBandwidthResponse) {
	response = &RenewCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RenewCcnBandwidth）用于续费预付费模式下云联网实例的地域间带宽。
func (c *Client) RenewCcnBandwidth(request *RenewCcnBandwidthRequest) (response *RenewCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewRenewCcnBandwidthRequest()
	}
	response = NewRenewCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteTableSelectionPoliciesRequest() (request *DescribeRouteTableSelectionPoliciesRequest) {
	request = &DescribeRouteTableSelectionPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeRouteTableSelectionPolicies")
	return
}

func NewDescribeRouteTableSelectionPoliciesResponse() (response *DescribeRouteTableSelectionPoliciesResponse) {
	response = &DescribeRouteTableSelectionPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRouteTableSelectionPolicies）用于查询云联网路由表选择策略。
func (c *Client) DescribeRouteTableSelectionPolicies(request *DescribeRouteTableSelectionPoliciesRequest) (response *DescribeRouteTableSelectionPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeRouteTableSelectionPoliciesRequest()
	}
	response = NewDescribeRouteTableSelectionPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRouteMatchRuleRequest() (request *DeleteCcnRouteMatchRuleRequest) {
	request = &DeleteCcnRouteMatchRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcnRouteMatchRule")
	return
}

func NewDeleteCcnRouteMatchRuleResponse() (response *DeleteCcnRouteMatchRuleResponse) {
	response = &DeleteCcnRouteMatchRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由匹配规则
func (c *Client) DeleteCcnRouteMatchRule(request *DeleteCcnRouteMatchRuleRequest) (response *DeleteCcnRouteMatchRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCcnRouteMatchRuleRequest()
	}
	response = NewDeleteCcnRouteMatchRuleResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceCcnRouteTableInputPolicysRequest() (request *ReplaceCcnRouteTableInputPolicysRequest) {
	request = &ReplaceCcnRouteTableInputPolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ReplaceCcnRouteTableInputPolicys")
	return
}

func NewReplaceCcnRouteTableInputPolicysResponse() (response *ReplaceCcnRouteTableInputPolicysResponse) {
	response = &ReplaceCcnRouteTableInputPolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ReplaceRouteTableInputPolicys)用于替换云联网路由表路由接收策略。
// > 特别注意：是全量覆盖，非增量添加
//
// **路由条件支持以下四种：**
//
// - 实例类型: `instance-type`，可选值：私有网络 `VPC`、专线网关 `DIRECTCONNECT`、VPN网关 `VPNGW`
// - 实例ID: `instance-id`，例如：`dcg-8zljkrft`、`vpc-jdevjrup`，暂不支持 `Edge` 实例
// - 实例地域: `instance-region`，例如：`region1`<br />产品支持的所有地域列表可通过接口 [DescribeRegions](https://{{conf.main_domain}}/document/product/1596/77930) 查询，其中参数 `Product` 设置为 `ccn`
// - 路由前缀: `cidr-block`，例如：`10.1.0.0/16`
//
// **使用限制：**
// - 一条策略内的单个条件类型，最大支持设置 `25` 个条件值
// - 一张路由表，最大支持 `100` 条路由接收策略
// - 路由条件类型中，只有 `cidr-block` 类型支持模糊匹配和精确匹配两种，其它类型只支持精确匹配一种模式
func (c *Client) ReplaceCcnRouteTableInputPolicys(request *ReplaceCcnRouteTableInputPolicysRequest) (response *ReplaceCcnRouteTableInputPolicysResponse, err error) {
	if request == nil {
		request = NewReplaceCcnRouteTableInputPolicysRequest()
	}
	response = NewReplaceCcnRouteTableInputPolicysResponse()
	err = c.Send(request, response)
	return
}

func NewGetCcnTrafficRemainAmountRequest() (request *GetCcnTrafficRemainAmountRequest) {
	request = &GetCcnTrafficRemainAmountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetCcnTrafficRemainAmount")
	return
}

func NewGetCcnTrafficRemainAmountResponse() (response *GetCcnTrafficRemainAmountResponse) {
	response = &GetCcnTrafficRemainAmountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ccn实例免费额度用量
func (c *Client) GetCcnTrafficRemainAmount(request *GetCcnTrafficRemainAmountRequest) (response *GetCcnTrafficRemainAmountResponse, err error) {
	if request == nil {
		request = NewGetCcnTrafficRemainAmountRequest()
	}
	response = NewGetCcnTrafficRemainAmountResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnBandwidthRenewFlagRequest() (request *SetCcnBandwidthRenewFlagRequest) {
	request = &SetCcnBandwidthRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetCcnBandwidthRenewFlag")
	return
}

func NewSetCcnBandwidthRenewFlagResponse() (response *SetCcnBandwidthRenewFlagResponse) {
	response = &SetCcnBandwidthRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SetCcnBandwidthRenewFlag）用于设置预付费云联网实例地域间限速的自动续费标记。
func (c *Client) SetCcnBandwidthRenewFlag(request *SetCcnBandwidthRenewFlagRequest) (response *SetCcnBandwidthRenewFlagResponse, err error) {
	if request == nil {
		request = NewSetCcnBandwidthRenewFlagRequest()
	}
	response = NewSetCcnBandwidthRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateCcnBandwidthRequest() (request *InquiryPriceCreateCcnBandwidthRequest) {
	request = &InquiryPriceCreateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquiryPriceCreateCcnBandwidth")
	return
}

func NewInquiryPriceCreateCcnBandwidthResponse() (response *InquiryPriceCreateCcnBandwidthResponse) {
	response = &InquiryPriceCreateCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceCreateCcnBandwidth）用于创建预付费云联网实例地域间带宽时的询价
func (c *Client) InquiryPriceCreateCcnBandwidth(request *InquiryPriceCreateCcnBandwidthRequest) (response *InquiryPriceCreateCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateCcnBandwidthRequest()
	}
	response = NewInquiryPriceCreateCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRouteTablesRequest() (request *DeleteCcnRouteTablesRequest) {
	request = &DeleteCcnRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcnRouteTables")
	return
}

func NewDeleteCcnRouteTablesResponse() (response *DeleteCcnRouteTablesResponse) {
	response = &DeleteCcnRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCcnRouteTables）用于删除云联网路由表。
func (c *Client) DeleteCcnRouteTables(request *DeleteCcnRouteTablesRequest) (response *DeleteCcnRouteTablesResponse, err error) {
	if request == nil {
		request = NewDeleteCcnRouteTablesRequest()
	}
	response = NewDeleteCcnRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnPolicyBasedRoutingNextHopRequest() (request *DeleteCcnPolicyBasedRoutingNextHopRequest) {
	request = &DeleteCcnPolicyBasedRoutingNextHopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcnPolicyBasedRoutingNextHop")
	return
}

func NewDeleteCcnPolicyBasedRoutingNextHopResponse() (response *DeleteCcnPolicyBasedRoutingNextHopResponse) {
	response = &DeleteCcnPolicyBasedRoutingNextHopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除云联网策略路由下一跳
func (c *Client) DeleteCcnPolicyBasedRoutingNextHop(request *DeleteCcnPolicyBasedRoutingNextHopRequest) (response *DeleteCcnPolicyBasedRoutingNextHopResponse, err error) {
	if request == nil {
		request = NewDeleteCcnPolicyBasedRoutingNextHopRequest()
	}
	response = NewDeleteCcnPolicyBasedRoutingNextHopResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnPolicyBasedRoutingNextHopRequest() (request *CreateCcnPolicyBasedRoutingNextHopRequest) {
	request = &CreateCcnPolicyBasedRoutingNextHopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcnPolicyBasedRoutingNextHop")
	return
}

func NewCreateCcnPolicyBasedRoutingNextHopResponse() (response *CreateCcnPolicyBasedRoutingNextHopResponse) {
	response = &CreateCcnPolicyBasedRoutingNextHopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云联网策略路由下一跳
func (c *Client) CreateCcnPolicyBasedRoutingNextHop(request *CreateCcnPolicyBasedRoutingNextHopRequest) (response *CreateCcnPolicyBasedRoutingNextHopResponse, err error) {
	if request == nil {
		request = NewCreateCcnPolicyBasedRoutingNextHopRequest()
	}
	response = NewCreateCcnPolicyBasedRoutingNextHopResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCrossBorderComplianceRequest() (request *ModifyCrossBorderComplianceRequest) {
	request = &ModifyCrossBorderComplianceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCrossBorderCompliance")
	return
}

func NewModifyCrossBorderComplianceResponse() (response *ModifyCrossBorderComplianceResponse) {
	response = &ModifyCrossBorderComplianceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCrossBorderCompliance）用于用户修改合规化资质审批单。
// 用户修改任意信息，审批单都将转入待审批状态。未审批通过的用户，将影响带宽购买。请谨慎修改。
func (c *Client) ModifyCrossBorderCompliance(request *ModifyCrossBorderComplianceRequest) (response *ModifyCrossBorderComplianceResponse, err error) {
	if request == nil {
		request = NewModifyCrossBorderComplianceRequest()
	}
	response = NewModifyCrossBorderComplianceResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpdateDirectConnectAccelerateChannelBandwidthDealRequest() (request *GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest) {
	request = &GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetUpdateDirectConnectAccelerateChannelBandwidthDeal")
	return
}

func NewGetUpdateDirectConnectAccelerateChannelBandwidthDealResponse() (response *GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse) {
	response = &GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更加速通道带宽-前端获取订单参数
func (c *Client) GetUpdateDirectConnectAccelerateChannelBandwidthDeal(request *GetUpdateDirectConnectAccelerateChannelBandwidthDealRequest) (response *GetUpdateDirectConnectAccelerateChannelBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetUpdateDirectConnectAccelerateChannelBandwidthDealRequest()
	}
	response = NewGetUpdateDirectConnectAccelerateChannelBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRequest() (request *DeleteCcnRequest) {
	request = &DeleteCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcn")
	return
}

func NewDeleteCcnResponse() (response *DeleteCcnResponse) {
	response = &DeleteCcnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCcn）用于删除云联网。
// * 删除后，云联网关联的所有实例间路由将被删除，网络将会中断，请务必确认
// * 删除云联网是不可逆的操作，请谨慎处理。
func (c *Client) DeleteCcn(request *DeleteCcnRequest) (response *DeleteCcnResponse, err error) {
	if request == nil {
		request = NewDeleteCcnRequest()
	}
	response = NewDeleteCcnResponse()
	err = c.Send(request, response)
	return
}

func NewUnlockCcnsRequest() (request *UnlockCcnsRequest) {
	request = &UnlockCcnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "UnlockCcns")
	return
}

func NewUnlockCcnsResponse() (response *UnlockCcnsResponse) {
	response = &UnlockCcnsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnlockCcns）用于解锁云联网实例
//
// 该接口一般用来解封禁出口限速的云联网实例, 目前联通内部运营系统通过云API调用, 因为出口限速无法按地域间解封禁, 只能按更粗的云联网实例粒度解封禁, 如果是地域间限速, 一般可以通过更细的限速实例粒度解封禁（UnlockCcnBandwidths）
//
// 如有需要, 可以封禁任意限速实例, 可接入到内部运营系统
func (c *Client) UnlockCcns(request *UnlockCcnsRequest) (response *UnlockCcnsResponse, err error) {
	if request == nil {
		request = NewUnlockCcnsRequest()
	}
	response = NewUnlockCcnsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnBandwidthConfigRequest() (request *DescribeCcnBandwidthConfigRequest) {
	request = &DescribeCcnBandwidthConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnBandwidthConfig")
	return
}

func NewDescribeCcnBandwidthConfigResponse() (response *DescribeCcnBandwidthConfigResponse) {
	response = &DescribeCcnBandwidthConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnBandwidthConfig）用于查询某个云联网某一对地域的带宽变配最大值。
func (c *Client) DescribeCcnBandwidthConfig(request *DescribeCcnBandwidthConfigRequest) (response *DescribeCcnBandwidthConfigResponse, err error) {
	if request == nil {
		request = NewDescribeCcnBandwidthConfigRequest()
	}
	response = NewDescribeCcnBandwidthConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCcnBandwidthRequest() (request *UpdateCcnBandwidthRequest) {
	request = &UpdateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "UpdateCcnBandwidth")
	return
}

func NewUpdateCcnBandwidthResponse() (response *UpdateCcnBandwidthResponse) {
	response = &UpdateCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCcnBandwidth）用于变更预付费模式下云联网实例的地域间带宽。
func (c *Client) UpdateCcnBandwidth(request *UpdateCcnBandwidthRequest) (response *UpdateCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewUpdateCcnBandwidthRequest()
	}
	response = NewUpdateCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnRoutePriorityRequest() (request *ModifyCcnRoutePriorityRequest) {
	request = &ModifyCcnRoutePriorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnRoutePriority")
	return
}

func NewModifyCcnRoutePriorityResponse() (response *ModifyCcnRoutePriorityResponse) {
	response = &ModifyCcnRoutePriorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCcnRoutePriority）用于调整云联网路由表中单个路由的优先级字段。
func (c *Client) ModifyCcnRoutePriority(request *ModifyCcnRoutePriorityRequest) (response *ModifyCcnRoutePriorityResponse, err error) {
	if request == nil {
		request = NewModifyCcnRoutePriorityRequest()
	}
	response = NewModifyCcnRoutePriorityResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnRouteTablesRequest() (request *CreateCcnRouteTablesRequest) {
	request = &CreateCcnRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcnRouteTables")
	return
}

func NewCreateCcnRouteTablesResponse() (response *CreateCcnRouteTablesResponse) {
	response = &CreateCcnRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCcnRouteTables）用于给指定的云联网实例新建路由表。
func (c *Client) CreateCcnRouteTables(request *CreateCcnRouteTablesRequest) (response *CreateCcnRouteTablesResponse, err error) {
	if request == nil {
		request = NewCreateCcnRouteTablesRequest()
	}
	response = NewCreateCcnRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRouteTableSelectionPoliciesRequest() (request *ModifyRouteTableSelectionPoliciesRequest) {
	request = &ModifyRouteTableSelectionPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyRouteTableSelectionPolicies")
	return
}

func NewModifyRouteTableSelectionPoliciesResponse() (response *ModifyRouteTableSelectionPoliciesResponse) {
	response = &ModifyRouteTableSelectionPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于编辑云联网路由表选择策略
func (c *Client) ModifyRouteTableSelectionPolicies(request *ModifyRouteTableSelectionPoliciesRequest) (response *ModifyRouteTableSelectionPoliciesResponse, err error) {
	if request == nil {
		request = NewModifyRouteTableSelectionPoliciesRequest()
	}
	response = NewModifyRouteTableSelectionPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteReceivingPoliciesRequest() (request *CreateRouteReceivingPoliciesRequest) {
	request = &CreateRouteReceivingPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateRouteReceivingPolicies")
	return
}

func NewCreateRouteReceivingPoliciesResponse() (response *CreateRouteReceivingPoliciesResponse) {
	response = &CreateRouteReceivingPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于创建路由接收策略
func (c *Client) CreateRouteReceivingPolicies(request *CreateRouteReceivingPoliciesRequest) (response *CreateRouteReceivingPoliciesResponse, err error) {
	if request == nil {
		request = NewCreateRouteReceivingPoliciesRequest()
	}
	response = NewCreateRouteReceivingPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectPortMapRequest() (request *DescribeDirectConnectPortMapRequest) {
	request = &DescribeDirectConnectPortMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeDirectConnectPortMap")
	return
}

func NewDescribeDirectConnectPortMapResponse() (response *DescribeDirectConnectPortMapResponse) {
	response = &DescribeDirectConnectPortMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询端口映射
func (c *Client) DescribeDirectConnectPortMap(request *DescribeDirectConnectPortMapRequest) (response *DescribeDirectConnectPortMapResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectPortMapRequest()
	}
	response = NewDescribeDirectConnectPortMapResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnPolicyBasedRoutingRuleAttributeRequest() (request *ModifyCcnPolicyBasedRoutingRuleAttributeRequest) {
	request = &ModifyCcnPolicyBasedRoutingRuleAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnPolicyBasedRoutingRuleAttribute")
	return
}

func NewModifyCcnPolicyBasedRoutingRuleAttributeResponse() (response *ModifyCcnPolicyBasedRoutingRuleAttributeResponse) {
	response = &ModifyCcnPolicyBasedRoutingRuleAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新云联网策略路由匹配规则参数
func (c *Client) ModifyCcnPolicyBasedRoutingRuleAttribute(request *ModifyCcnPolicyBasedRoutingRuleAttributeRequest) (response *ModifyCcnPolicyBasedRoutingRuleAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCcnPolicyBasedRoutingRuleAttributeRequest()
	}
	response = NewModifyCcnPolicyBasedRoutingRuleAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrafficQosPolicyRequest() (request *ModifyTrafficQosPolicyRequest) {
	request = &ModifyTrafficQosPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyTrafficQosPolicy")
	return
}

func NewModifyTrafficQosPolicyResponse() (response *ModifyTrafficQosPolicyResponse) {
	response = &ModifyTrafficQosPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改流量调度策略
func (c *Client) ModifyTrafficQosPolicy(request *ModifyTrafficQosPolicyRequest) (response *ModifyTrafficQosPolicyResponse, err error) {
	if request == nil {
		request = NewModifyTrafficQosPolicyRequest()
	}
	response = NewModifyTrafficQosPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceCreateCcnInstanceRequest() (request *InquirePriceCreateCcnInstanceRequest) {
	request = &InquirePriceCreateCcnInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquirePriceCreateCcnInstance")
	return
}

func NewInquirePriceCreateCcnInstanceResponse() (response *InquirePriceCreateCcnInstanceResponse) {
	response = &InquirePriceCreateCcnInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云联网实例费和流量费询价
func (c *Client) InquirePriceCreateCcnInstance(request *InquirePriceCreateCcnInstanceRequest) (response *InquirePriceCreateCcnInstanceResponse, err error) {
	if request == nil {
		request = NewInquirePriceCreateCcnInstanceRequest()
	}
	response = NewInquirePriceCreateCcnInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateCcnGatewayRequest() (request *MigrateCcnGatewayRequest) {
	request = &MigrateCcnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "MigrateCcnGateway")
	return
}

func NewMigrateCcnGatewayResponse() (response *MigrateCcnGatewayResponse) {
	response = &MigrateCcnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（MigrateCcnGateway）用于迁移云联网集群，其中云联网网关要在同一个地域。
func (c *Client) MigrateCcnGateway(request *MigrateCcnGatewayRequest) (response *MigrateCcnGatewayResponse, err error) {
	if request == nil {
		request = NewMigrateCcnGatewayRequest()
	}
	response = NewMigrateCcnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantCcnsRequest() (request *DescribeTenantCcnsRequest) {
	request = &DescribeTenantCcnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeTenantCcns")
	return
}

func NewDescribeTenantCcnsResponse() (response *DescribeTenantCcnsResponse) {
	response = &DescribeTenantCcnsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeTenantCcns）用于获取要锁定的云联网实例列表。
// 该接口一般用来封禁出口限速的云联网实例, 目前联通内部运营系统通过云API调用, 因为出口限速无法按地域间封禁, 只能按更粗的云联网实例粒度封禁, 如果是地域间限速, 一般可以通过更细的限速实例粒度封禁（DescribeCrossBorderCcnRegionBandwidthLimits）
// 如有需要, 可以封禁任意云联网实例, 可接入到内部运营系统
func (c *Client) DescribeTenantCcns(request *DescribeTenantCcnsRequest) (response *DescribeTenantCcnsResponse, err error) {
	if request == nil {
		request = NewDescribeTenantCcnsRequest()
	}
	response = NewDescribeTenantCcnsResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCcnInstanceNonDirectFlagRequest() (request *DisableCcnInstanceNonDirectFlagRequest) {
	request = &DisableCcnInstanceNonDirectFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DisableCcnInstanceNonDirectFlag")
	return
}

func NewDisableCcnInstanceNonDirectFlagResponse() (response *DisableCcnInstanceNonDirectFlagResponse) {
	response = &DisableCcnInstanceNonDirectFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭云联网实例非直通标记
func (c *Client) DisableCcnInstanceNonDirectFlag(request *DisableCcnInstanceNonDirectFlagRequest) (response *DisableCcnInstanceNonDirectFlagResponse, err error) {
	if request == nil {
		request = NewDisableCcnInstanceNonDirectFlagRequest()
	}
	response = NewDisableCcnInstanceNonDirectFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnInstanceNonDirectFlagRequest() (request *DescribeCcnInstanceNonDirectFlagRequest) {
	request = &DescribeCcnInstanceNonDirectFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnInstanceNonDirectFlag")
	return
}

func NewDescribeCcnInstanceNonDirectFlagResponse() (response *DescribeCcnInstanceNonDirectFlagResponse) {
	response = &DescribeCcnInstanceNonDirectFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看云联网实例非直通标记
func (c *Client) DescribeCcnInstanceNonDirectFlag(request *DescribeCcnInstanceNonDirectFlagRequest) (response *DescribeCcnInstanceNonDirectFlagResponse, err error) {
	if request == nil {
		request = NewDescribeCcnInstanceNonDirectFlagRequest()
	}
	response = NewDescribeCcnInstanceNonDirectFlagResponse()
	err = c.Send(request, response)
	return
}

func NewGetDealStatusByNameRequest() (request *GetDealStatusByNameRequest) {
	request = &GetDealStatusByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetDealStatusByName")
	return
}

func NewGetDealStatusByNameResponse() (response *GetDealStatusByNameResponse) {
	response = &GetDealStatusByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetDealStatusByName）用于查询云联网预付费订单的状态信息
func (c *Client) GetDealStatusByName(request *GetDealStatusByNameRequest) (response *GetDealStatusByNameResponse, err error) {
	if request == nil {
		request = NewGetDealStatusByNameRequest()
	}
	response = NewGetDealStatusByNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectAccelerateChannelRequest() (request *DescribeDirectConnectAccelerateChannelRequest) {
	request = &DescribeDirectConnectAccelerateChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeDirectConnectAccelerateChannel")
	return
}

func NewDescribeDirectConnectAccelerateChannelResponse() (response *DescribeDirectConnectAccelerateChannelResponse) {
	response = &DescribeDirectConnectAccelerateChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询加速通道列表
func (c *Client) DescribeDirectConnectAccelerateChannel(request *DescribeDirectConnectAccelerateChannelRequest) (response *DescribeDirectConnectAccelerateChannelResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectAccelerateChannelRequest()
	}
	response = NewDescribeDirectConnectAccelerateChannelResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCcnRoutesRequest() (request *DisableCcnRoutesRequest) {
	request = &DisableCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DisableCcnRoutes")
	return
}

func NewDisableCcnRoutesResponse() (response *DisableCcnRoutesResponse) {
	response = &DisableCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableCcnRoutes）用于禁用已经启用的云联网（CCN）路由。
func (c *Client) DisableCcnRoutes(request *DisableCcnRoutesRequest) (response *DisableCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDisableCcnRoutesRequest()
	}
	response = NewDisableCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCrossBorderCcnRegionBandwidthLimitsRequest() (request *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) {
	request = &DescribeCrossBorderCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCrossBorderCcnRegionBandwidthLimits")
	return
}

func NewDescribeCrossBorderCcnRegionBandwidthLimitsResponse() (response *DescribeCrossBorderCcnRegionBandwidthLimitsResponse) {
	response = &DescribeCrossBorderCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCrossBorderCcnRegionBandwidthLimits）用于获取要锁定的限速实例列表。
// 该接口一般用来封禁地域间限速的云联网实例下的限速实例, 目前联通内部运营系统通过云API调用, 如果是出口限速, 一般使用更粗的云联网实例粒度封禁（DescribeTenantCcns）
// 如有需要, 可以封禁任意限速实例, 可接入到内部运营系统
func (c *Client) DescribeCrossBorderCcnRegionBandwidthLimits(request *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) (response *DescribeCrossBorderCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeCrossBorderCcnRegionBandwidthLimitsRequest()
	}
	response = NewDescribeCrossBorderCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDetachCcnInstancesRequest() (request *DetachCcnInstancesRequest) {
	request = &DetachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DetachCcnInstances")
	return
}

func NewDetachCcnInstancesResponse() (response *DetachCcnInstancesResponse) {
	response = &DetachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DetachCcnInstances）用于从云联网实例中解关联指定的网络实例。<br />
// 解关联网络实例后，相应的路由策略会一并删除。
func (c *Client) DetachCcnInstances(request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewDetachCcnInstancesRequest()
	}
	response = NewDetachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetTradeBillingAuthRequest() (request *GetTradeBillingAuthRequest) {
	request = &GetTradeBillingAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetTradeBillingAuth")
	return
}

func NewGetTradeBillingAuthResponse() (response *GetTradeBillingAuthResponse) {
	response = &GetTradeBillingAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费下单屏蔽余额支付场景-获取AUTH
func (c *Client) GetTradeBillingAuth(request *GetTradeBillingAuthRequest) (response *GetTradeBillingAuthResponse, err error) {
	if request == nil {
		request = NewGetTradeBillingAuthRequest()
	}
	response = NewGetTradeBillingAuthResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnRegionFixedBandwidthLimitsRequest() (request *SetCcnRegionFixedBandwidthLimitsRequest) {
	request = &SetCcnRegionFixedBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetCcnRegionFixedBandwidthLimits")
	return
}

func NewSetCcnRegionFixedBandwidthLimitsResponse() (response *SetCcnRegionFixedBandwidthLimitsResponse) {
	response = &SetCcnRegionFixedBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置固定带宽-联通专属接口
func (c *Client) SetCcnRegionFixedBandwidthLimits(request *SetCcnRegionFixedBandwidthLimitsRequest) (response *SetCcnRegionFixedBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewSetCcnRegionFixedBandwidthLimitsRequest()
	}
	response = NewSetCcnRegionFixedBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCrossBorderSettlementRequest() (request *DescribeCrossBorderSettlementRequest) {
	request = &DescribeCrossBorderSettlementRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCrossBorderSettlement")
	return
}

func NewDescribeCrossBorderSettlementResponse() (response *DescribeCrossBorderSettlementResponse) {
	response = &DescribeCrossBorderSettlementResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询跨境账单-联通专用
func (c *Client) DescribeCrossBorderSettlement(request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
	if request == nil {
		request = NewDescribeCrossBorderSettlementRequest()
	}
	response = NewDescribeCrossBorderSettlementResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewCcnBandwidthRequest() (request *InquiryPriceRenewCcnBandwidthRequest) {
	request = &InquiryPriceRenewCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquiryPriceRenewCcnBandwidth")
	return
}

func NewInquiryPriceRenewCcnBandwidthResponse() (response *InquiryPriceRenewCcnBandwidthResponse) {
	response = &InquiryPriceRenewCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceRenewCcnBandwidth）用于续费云联网实例地域间带宽时的询价
func (c *Client) InquiryPriceRenewCcnBandwidth(request *InquiryPriceRenewCcnBandwidthRequest) (response *InquiryPriceRenewCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewCcnBandwidthRequest()
	}
	response = NewInquiryPriceRenewCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRouteTableAggregatePolicysRequest() (request *DescribeCcnRouteTableAggregatePolicysRequest) {
	request = &DescribeCcnRouteTableAggregatePolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRouteTableAggregatePolicys")
	return
}

func NewDescribeCcnRouteTableAggregatePolicysResponse() (response *DescribeCcnRouteTableAggregatePolicysResponse) {
	response = &DescribeCcnRouteTableAggregatePolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeCcnRouteTableAggregatePolicys)用于查询指定云联网路由表的路由聚合策略。
func (c *Client) DescribeCcnRouteTableAggregatePolicys(request *DescribeCcnRouteTableAggregatePolicysRequest) (response *DescribeCcnRouteTableAggregatePolicysResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRouteTableAggregatePolicysRequest()
	}
	response = NewDescribeCcnRouteTableAggregatePolicysResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnFlowLogAttributeRequest() (request *ModifyCcnFlowLogAttributeRequest) {
	request = &ModifyCcnFlowLogAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnFlowLogAttribute")
	return
}

func NewModifyCcnFlowLogAttributeResponse() (response *ModifyCcnFlowLogAttributeResponse) {
	response = &ModifyCcnFlowLogAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCcnFlowLogAttribute）用于修改云联网流日志属性
func (c *Client) ModifyCcnFlowLogAttribute(request *ModifyCcnFlowLogAttributeRequest) (response *ModifyCcnFlowLogAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCcnFlowLogAttributeRequest()
	}
	response = NewModifyCcnFlowLogAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCcnRoutesRequest() (request *EnableCcnRoutesRequest) {
	request = &EnableCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "EnableCcnRoutes")
	return
}

func NewEnableCcnRoutesResponse() (response *EnableCcnRoutesResponse) {
	response = &EnableCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（EnableCcnRoutes）用于启用已经加入云联网（CCN）的路由。<br />
// 本接口会校验启用后，是否与已有路由冲突，如果冲突，则无法启用，失败处理。路由冲突时，需要先禁用与之冲突的路由，才能启用该路由。
func (c *Client) EnableCcnRoutes(request *EnableCcnRoutesRequest) (response *EnableCcnRoutesResponse, err error) {
	if request == nil {
		request = NewEnableCcnRoutesRequest()
	}
	response = NewEnableCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewUnlockCcnBandwidthsRequest() (request *UnlockCcnBandwidthsRequest) {
	request = &UnlockCcnBandwidthsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "UnlockCcnBandwidths")
	return
}

func NewUnlockCcnBandwidthsResponse() (response *UnlockCcnBandwidthsResponse) {
	response = &UnlockCcnBandwidthsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnlockCcnBandwidths）用户解锁云联网限速实例。
// 该接口一般用来封禁地域间限速的云联网实例下的限速实例, 目前联通内部运营系统通过云API调用, 如果是出口限速, 一般使用更粗的云联网实例粒度封禁（SecurityUnlockCcns）。
// 如有需要, 可以封禁任意限速实例, 可接入到内部运营系统。
func (c *Client) UnlockCcnBandwidths(request *UnlockCcnBandwidthsRequest) (response *UnlockCcnBandwidthsResponse, err error) {
	if request == nil {
		request = NewUnlockCcnBandwidthsRequest()
	}
	response = NewUnlockCcnBandwidthsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnAttachedInstancesAttributeRequest() (request *ModifyCcnAttachedInstancesAttributeRequest) {
	request = &ModifyCcnAttachedInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnAttachedInstancesAttribute")
	return
}

func NewModifyCcnAttachedInstancesAttributeResponse() (response *ModifyCcnAttachedInstancesAttributeResponse) {
	response = &ModifyCcnAttachedInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改CCN关联实例属性，目前仅修改备注description
func (c *Client) ModifyCcnAttachedInstancesAttribute(request *ModifyCcnAttachedInstancesAttributeRequest) (response *ModifyCcnAttachedInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCcnAttachedInstancesAttributeRequest()
	}
	response = NewModifyCcnAttachedInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewLockCcnBandwidthsRequest() (request *LockCcnBandwidthsRequest) {
	request = &LockCcnBandwidthsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "LockCcnBandwidths")
	return
}

func NewLockCcnBandwidthsResponse() (response *LockCcnBandwidthsResponse) {
	response = &LockCcnBandwidthsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（LockCcnBandwidths）用户锁定云联网限速实例。
// 该接口一般用来封禁地域间限速的云联网实例下的限速实例, 目前联通内部运营系统通过云API调用, 如果是出口限速, 一般使用更粗的云联网实例粒度封禁（LockCcns）。
// 如有需要, 可以封禁任意限速实例, 可接入到内部运营系统。
func (c *Client) LockCcnBandwidths(request *LockCcnBandwidthsRequest) (response *LockCcnBandwidthsResponse, err error) {
	if request == nil {
		request = NewLockCcnBandwidthsRequest()
	}
	response = NewLockCcnBandwidthsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCrossBorderCosTokenRequest() (request *GetCrossBorderCosTokenRequest) {
	request = &GetCrossBorderCosTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetCrossBorderCosToken")
	return
}

func NewGetCrossBorderCosTokenResponse() (response *GetCrossBorderCosTokenResponse) {
	response = &GetCrossBorderCosTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询跨境Cos临时Token GetCrossBorderCosToken
func (c *Client) GetCrossBorderCosToken(request *GetCrossBorderCosTokenRequest) (response *GetCrossBorderCosTokenResponse, err error) {
	if request == nil {
		request = NewGetCrossBorderCosTokenRequest()
	}
	response = NewGetCrossBorderCosTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnRouteMatchRuleRequest() (request *CreateCcnRouteMatchRuleRequest) {
	request = &CreateCcnRouteMatchRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcnRouteMatchRule")
	return
}

func NewCreateCcnRouteMatchRuleResponse() (response *CreateCcnRouteMatchRuleResponse) {
	response = &CreateCcnRouteMatchRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建路由匹配规则
func (c *Client) CreateCcnRouteMatchRule(request *CreateCcnRouteMatchRuleRequest) (response *CreateCcnRouteMatchRuleResponse, err error) {
	if request == nil {
		request = NewCreateCcnRouteMatchRuleRequest()
	}
	response = NewCreateCcnRouteMatchRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnsRequest() (request *DescribeCcnsRequest) {
	request = &DescribeCcnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcns")
	return
}

func NewDescribeCcnsResponse() (response *DescribeCcnsResponse) {
	response = &DescribeCcnsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcns）用于查询云联网（CCN）列表。
func (c *Client) DescribeCcns(request *DescribeCcnsRequest) (response *DescribeCcnsResponse, err error) {
	if request == nil {
		request = NewDescribeCcnsRequest()
	}
	response = NewDescribeCcnsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnRegionBandwidthLimitsTypeRequest() (request *ModifyCcnRegionBandwidthLimitsTypeRequest) {
	request = &ModifyCcnRegionBandwidthLimitsTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnRegionBandwidthLimitsType")
	return
}

func NewModifyCcnRegionBandwidthLimitsTypeResponse() (response *ModifyCcnRegionBandwidthLimitsTypeResponse) {
	response = &ModifyCcnRegionBandwidthLimitsTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCcnRegionBandwidthLimitsType）用于修改后付费云联网实例修改带宽限速策略。
func (c *Client) ModifyCcnRegionBandwidthLimitsType(request *ModifyCcnRegionBandwidthLimitsTypeRequest) (response *ModifyCcnRegionBandwidthLimitsTypeResponse, err error) {
	if request == nil {
		request = NewModifyCcnRegionBandwidthLimitsTypeRequest()
	}
	response = NewModifyCcnRegionBandwidthLimitsTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnLimitsRequest() (request *DescribeCcnLimitsRequest) {
	request = &DescribeCcnLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnLimits")
	return
}

func NewDescribeCcnLimitsResponse() (response *DescribeCcnLimitsResponse) {
	response = &DescribeCcnLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnLimits）用于查询云联网配额。
func (c *Client) DescribeCcnLimits(request *DescribeCcnLimitsRequest) (response *DescribeCcnLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeCcnLimitsRequest()
	}
	response = NewDescribeCcnLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCrossBorderComplianceRequest() (request *CreateCrossBorderComplianceRequest) {
	request = &CreateCrossBorderComplianceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCrossBorderCompliance")
	return
}

func NewCreateCrossBorderComplianceResponse() (response *CreateCrossBorderComplianceResponse) {
	response = &CreateCrossBorderComplianceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCrossBorderCompliance）用于用户创建合规化资质审批单。提交后，由服务商审批。
func (c *Client) CreateCrossBorderCompliance(request *CreateCrossBorderComplianceRequest) (response *CreateCrossBorderComplianceResponse, err error) {
	if request == nil {
		request = NewCreateCrossBorderComplianceRequest()
	}
	response = NewCreateCrossBorderComplianceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRouteTableInputPolicysRequest() (request *DescribeCcnRouteTableInputPolicysRequest) {
	request = &DescribeCcnRouteTableInputPolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRouteTableInputPolicys")
	return
}

func NewDescribeCcnRouteTableInputPolicysResponse() (response *DescribeCcnRouteTableInputPolicysResponse) {
	response = &DescribeCcnRouteTableInputPolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeCcnRouteTableInputPolicys)用于查询指定云联网路由表的路由接收策略。
func (c *Client) DescribeCcnRouteTableInputPolicys(request *DescribeCcnRouteTableInputPolicysRequest) (response *DescribeCcnRouteTableInputPolicysResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRouteTableInputPolicysRequest()
	}
	response = NewDescribeCcnRouteTableInputPolicysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficQosPolicyRequest() (request *DescribeTrafficQosPolicyRequest) {
	request = &DescribeTrafficQosPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeTrafficQosPolicy")
	return
}

func NewDescribeTrafficQosPolicyResponse() (response *DescribeTrafficQosPolicyResponse) {
	response = &DescribeTrafficQosPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询流量调度规则
func (c *Client) DescribeTrafficQosPolicy(request *DescribeTrafficQosPolicyRequest) (response *DescribeTrafficQosPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficQosPolicyRequest()
	}
	response = NewDescribeTrafficQosPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRoutesRequest() (request *DescribeCcnRoutesRequest) {
	request = &DescribeCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRoutes")
	return
}

func NewDescribeCcnRoutesResponse() (response *DescribeCcnRoutesResponse) {
	response = &DescribeCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnRoutes）用于查询已加入云联网（CCN）的路由。
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRoutesRequest()
	}
	response = NewDescribeCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteReceivingPoliciesRequest() (request *DeleteRouteReceivingPoliciesRequest) {
	request = &DeleteRouteReceivingPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteRouteReceivingPolicies")
	return
}

func NewDeleteRouteReceivingPoliciesResponse() (response *DeleteRouteReceivingPoliciesResponse) {
	response = &DeleteRouteReceivingPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteRouteReceivingPolicies）用于删除云联网路由接收策略。
func (c *Client) DeleteRouteReceivingPolicies(request *DeleteRouteReceivingPoliciesRequest) (response *DeleteRouteReceivingPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteRouteReceivingPoliciesRequest()
	}
	response = NewDeleteRouteReceivingPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewSetDirectConnectAccelerateChannelBandwidthRenewFlagRequest() (request *SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest) {
	request = &SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetDirectConnectAccelerateChannelBandwidthRenewFlag")
	return
}

func NewSetDirectConnectAccelerateChannelBandwidthRenewFlagResponse() (response *SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse) {
	response = &SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置云联网带宽自动续费标志
func (c *Client) SetDirectConnectAccelerateChannelBandwidthRenewFlag(request *SetDirectConnectAccelerateChannelBandwidthRenewFlagRequest) (response *SetDirectConnectAccelerateChannelBandwidthRenewFlagResponse, err error) {
	if request == nil {
		request = NewSetDirectConnectAccelerateChannelBandwidthRenewFlagRequest()
	}
	response = NewSetDirectConnectAccelerateChannelBandwidthRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptAttachCcnInstancesRequest() (request *AcceptAttachCcnInstancesRequest) {
	request = &AcceptAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "AcceptAttachCcnInstances")
	return
}

func NewAcceptAttachCcnInstancesResponse() (response *AcceptAttachCcnInstancesResponse) {
	response = &AcceptAttachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AcceptAttachCcnInstances）用于跨账号关联实例时，云联网所有者接受并同意关联操作。
func (c *Client) AcceptAttachCcnInstances(request *AcceptAttachCcnInstancesRequest) (response *AcceptAttachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewAcceptAttachCcnInstancesRequest()
	}
	response = NewAcceptAttachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCcnIpv6RoutesRequest() (request *DisableCcnIpv6RoutesRequest) {
	request = &DisableCcnIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DisableCcnIpv6Routes")
	return
}

func NewDisableCcnIpv6RoutesResponse() (response *DisableCcnIpv6RoutesResponse) {
	response = &DisableCcnIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableCcnRoutes）用于禁用已经启用的云联网（CCN）Ipv6路由。
func (c *Client) DisableCcnIpv6Routes(request *DisableCcnIpv6RoutesRequest) (response *DisableCcnIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewDisableCcnIpv6RoutesRequest()
	}
	response = NewDisableCcnIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewGetRenewDirectConnectAccelerateChannelBandwidthDealRequest() (request *GetRenewDirectConnectAccelerateChannelBandwidthDealRequest) {
	request = &GetRenewDirectConnectAccelerateChannelBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetRenewDirectConnectAccelerateChannelBandwidthDeal")
	return
}

func NewGetRenewDirectConnectAccelerateChannelBandwidthDealResponse() (response *GetRenewDirectConnectAccelerateChannelBandwidthDealResponse) {
	response = &GetRenewDirectConnectAccelerateChannelBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 二层CCN加速通道续费带宽-前端获取订单参数
func (c *Client) GetRenewDirectConnectAccelerateChannelBandwidthDeal(request *GetRenewDirectConnectAccelerateChannelBandwidthDealRequest) (response *GetRenewDirectConnectAccelerateChannelBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetRenewDirectConnectAccelerateChannelBandwidthDealRequest()
	}
	response = NewGetRenewDirectConnectAccelerateChannelBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnRequest() (request *CreateCcnRequest) {
	request = &CreateCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcn")
	return
}

func NewCreateCcnResponse() (response *CreateCcnResponse) {
	response = &CreateCcnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCcn）用于创建云联网（CCN）。<br />
// * 创建云联网同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
// * 每个账号能创建的云联网实例个数是有限的，详请参考产品文档。如果需要扩充请联系在线客服。
func (c *Client) CreateCcn(request *CreateCcnRequest) (response *CreateCcnResponse, err error) {
	if request == nil {
		request = NewCreateCcnRequest()
	}
	response = NewCreateCcnResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRegionBandwidthLimitsRequest() (request *DescribeCcnRegionBandwidthLimitsRequest) {
	request = &DescribeCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRegionBandwidthLimits")
	return
}

func NewDescribeCcnRegionBandwidthLimitsResponse() (response *DescribeCcnRegionBandwidthLimitsResponse) {
	response = &DescribeCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnRegionBandwidthLimits）用于查询云联网各地域出带宽上限，该接口只返回已关联网络实例包含的地域。
func (c *Client) DescribeCcnRegionBandwidthLimits(request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRegionBandwidthLimitsRequest()
	}
	response = NewDescribeCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCcnInstanceNonDirectFlagRequest() (request *EnableCcnInstanceNonDirectFlagRequest) {
	request = &EnableCcnInstanceNonDirectFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "EnableCcnInstanceNonDirectFlag")
	return
}

func NewEnableCcnInstanceNonDirectFlagResponse() (response *EnableCcnInstanceNonDirectFlagResponse) {
	response = &EnableCcnInstanceNonDirectFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启云联网实例非直通标记
func (c *Client) EnableCcnInstanceNonDirectFlag(request *EnableCcnInstanceNonDirectFlagRequest) (response *EnableCcnInstanceNonDirectFlagResponse, err error) {
	if request == nil {
		request = NewEnableCcnInstanceNonDirectFlagRequest()
	}
	response = NewEnableCcnInstanceNonDirectFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRouteTableBroadcastPolicysRequest() (request *DescribeCcnRouteTableBroadcastPolicysRequest) {
	request = &DescribeCcnRouteTableBroadcastPolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRouteTableBroadcastPolicys")
	return
}

func NewDescribeCcnRouteTableBroadcastPolicysResponse() (response *DescribeCcnRouteTableBroadcastPolicysResponse) {
	response = &DescribeCcnRouteTableBroadcastPolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeCcnRouteTableBroadcastPolicys)用于查询指定云联网路由表的路由传播策略。
func (c *Client) DescribeCcnRouteTableBroadcastPolicys(request *DescribeCcnRouteTableBroadcastPolicysRequest) (response *DescribeCcnRouteTableBroadcastPolicysResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRouteTableBroadcastPolicysRequest()
	}
	response = NewDescribeCcnRouteTableBroadcastPolicysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCrossBorderComplianceRequest() (request *DescribeCrossBorderComplianceRequest) {
	request = &DescribeCrossBorderComplianceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCrossBorderCompliance")
	return
}

func NewDescribeCrossBorderComplianceResponse() (response *DescribeCrossBorderComplianceResponse) {
	response = &DescribeCrossBorderComplianceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCrossBorderCompliance）用于查询用户创建的合规化资质审批单。
// 服务商可以查询服务名下的任意 `APPID` 创建的审批单；非服务商，只能查询自己审批单。
func (c *Client) DescribeCrossBorderCompliance(request *DescribeCrossBorderComplianceRequest) (response *DescribeCrossBorderComplianceResponse, err error) {
	if request == nil {
		request = NewDescribeCrossBorderComplianceRequest()
	}
	response = NewDescribeCrossBorderComplianceResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest() (request *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest) {
	request = &InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquirePriceUpdateDirectConnectAccelerateChannelBandwidth")
	return
}

func NewInquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse() (response *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse) {
	response = &InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更加速通道带宽询价
func (c *Client) InquirePriceUpdateDirectConnectAccelerateChannelBandwidth(request *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest) (response *InquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse, err error) {
	if request == nil {
		request = NewInquirePriceUpdateDirectConnectAccelerateChannelBandwidthRequest()
	}
	response = NewInquirePriceUpdateDirectConnectAccelerateChannelBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewGetRenewCcnBandwidthDealRequest() (request *GetRenewCcnBandwidthDealRequest) {
	request = &GetRenewCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetRenewCcnBandwidthDeal")
	return
}

func NewGetRenewCcnBandwidthDealResponse() (response *GetRenewCcnBandwidthDealResponse) {
	response = &GetRenewCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetRenewCcnBandwidthDeal）用于获取续费云联网带宽订单信息。
func (c *Client) GetRenewCcnBandwidthDeal(request *GetRenewCcnBandwidthDealRequest) (response *GetRenewCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetRenewCcnBandwidthDealRequest()
	}
	response = NewGetRenewCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnBandwidthRequest() (request *CreateCcnBandwidthRequest) {
	request = &CreateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcnBandwidth")
	return
}

func NewCreateCcnBandwidthResponse() (response *CreateCcnBandwidthResponse) {
	response = &CreateCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCcnBandwidth）用于创建预付费模式下云联网实例的地域间带宽
func (c *Client) CreateCcnBandwidth(request *CreateCcnBandwidthRequest) (response *CreateCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewCreateCcnBandwidthRequest()
	}
	response = NewCreateCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCrossBorderRegionConfigRequest() (request *DescribeCrossBorderRegionConfigRequest) {
	request = &DescribeCrossBorderRegionConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCrossBorderRegionConfig")
	return
}

func NewDescribeCrossBorderRegionConfigResponse() (response *DescribeCrossBorderRegionConfigResponse) {
	response = &DescribeCrossBorderRegionConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询可购买的地域间专线带宽。
func (c *Client) DescribeCrossBorderRegionConfig(request *DescribeCrossBorderRegionConfigRequest) (response *DescribeCrossBorderRegionConfigResponse, err error) {
	if request == nil {
		request = NewDescribeCrossBorderRegionConfigRequest()
	}
	response = NewDescribeCrossBorderRegionConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnFlowLogRequest() (request *DescribeCcnFlowLogRequest) {
	request = &DescribeCcnFlowLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnFlowLog")
	return
}

func NewDescribeCcnFlowLogResponse() (response *DescribeCcnFlowLogResponse) {
	response = &DescribeCcnFlowLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnFlowLog）用于查询云联网流日志信息。
func (c *Client) DescribeCcnFlowLog(request *DescribeCcnFlowLogRequest) (response *DescribeCcnFlowLogResponse, err error) {
	if request == nil {
		request = NewDescribeCcnFlowLogRequest()
	}
	response = NewDescribeCcnFlowLogResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceCcnRouteTableAggregatePolicysRequest() (request *ReplaceCcnRouteTableAggregatePolicysRequest) {
	request = &ReplaceCcnRouteTableAggregatePolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ReplaceCcnRouteTableAggregatePolicys")
	return
}

func NewReplaceCcnRouteTableAggregatePolicysResponse() (response *ReplaceCcnRouteTableAggregatePolicysResponse) {
	response = &ReplaceCcnRouteTableAggregatePolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ReplaceRouteTableInputPolicys)用于替换云联网路由表路由聚合策略。
// > 特别注意：是全量覆盖，非增量添加
//
// **路由条件支持以下四种：**
//
// - 实例类型: `instance-type`，可选值：专线网关 `DIRECTCONNECT`、专线网关 `VPNGW`
// - 实例ID: `instance-id`，例如：`dcg-8zljkrft`、
// - 实例地域: `instance-region`，例如：`region1`<br />产品支持的所有地域列表可通过接口 [DescribeRegions](https://{{conf.main_domain}}/document/product/1596/77930) 查询，其中参数 `Product` 设置为 `ccn`
// - 路由前缀: `cidr-block`，例如：`10.1.0.0/16`
//
// **使用限制：**
// - 聚合条件 条件只支持 VPC、VPG、VPN 路由，不支持 SDWAN 和 黑石路由
// - VPC CIDR聚合只支持 聚合条件 为 VPC 实例的路由。
// - 不聚合时按照目前的行为根据专线侧开关决定下发 VPC CIDR还是子网到专线、默认下发 VPC CIDR 到NFV-VPN、默认下发子网到 CVM-VPN。
// - 自定义静态聚合路由之间不支持相互聚合。VPC CIDR 可以被 自定义静态聚合路由 聚合。
func (c *Client) ReplaceCcnRouteTableAggregatePolicys(request *ReplaceCcnRouteTableAggregatePolicysRequest) (response *ReplaceCcnRouteTableAggregatePolicysResponse, err error) {
	if request == nil {
		request = NewReplaceCcnRouteTableAggregatePolicysRequest()
	}
	response = NewReplaceCcnRouteTableAggregatePolicysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCrossBorderFlowMonitorRequest() (request *DescribeCrossBorderFlowMonitorRequest) {
	request = &DescribeCrossBorderFlowMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCrossBorderFlowMonitor")
	return
}

func NewDescribeCrossBorderFlowMonitorResponse() (response *DescribeCrossBorderFlowMonitorResponse) {
	response = &DescribeCrossBorderFlowMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCrossBorderFlowMonitor）用于查询跨境带宽监控数据，该接口目前只提供给服务商联通使用。
func (c *Client) DescribeCrossBorderFlowMonitor(request *DescribeCrossBorderFlowMonitorRequest) (response *DescribeCrossBorderFlowMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeCrossBorderFlowMonitorRequest()
	}
	response = NewDescribeCrossBorderFlowMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnPolicyBasedRoutingRulesRequest() (request *CreateCcnPolicyBasedRoutingRulesRequest) {
	request = &CreateCcnPolicyBasedRoutingRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateCcnPolicyBasedRoutingRules")
	return
}

func NewCreateCcnPolicyBasedRoutingRulesResponse() (response *CreateCcnPolicyBasedRoutingRulesResponse) {
	response = &CreateCcnPolicyBasedRoutingRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建云联网策略路由匹配规则
func (c *Client) CreateCcnPolicyBasedRoutingRules(request *CreateCcnPolicyBasedRoutingRulesRequest) (response *CreateCcnPolicyBasedRoutingRulesResponse, err error) {
	if request == nil {
		request = NewCreateCcnPolicyBasedRoutingRulesRequest()
	}
	response = NewCreateCcnPolicyBasedRoutingRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRouteTablesRequest() (request *DescribeCcnRouteTablesRequest) {
	request = &DescribeCcnRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRouteTables")
	return
}

func NewDescribeCcnRouteTablesResponse() (response *DescribeCcnRouteTablesResponse) {
	response = &DescribeCcnRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询指定的云联网实例的路由表信息。
func (c *Client) DescribeCcnRouteTables(request *DescribeCcnRouteTablesRequest) (response *DescribeCcnRouteTablesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRouteTablesRequest()
	}
	response = NewDescribeCcnRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTrafficQosPolicyRequest() (request *CreateTrafficQosPolicyRequest) {
	request = &CreateTrafficQosPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateTrafficQosPolicy")
	return
}

func NewCreateTrafficQosPolicyResponse() (response *CreateTrafficQosPolicyResponse) {
	response = &CreateTrafficQosPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateTrafficQosPolicy
func (c *Client) CreateTrafficQosPolicy(request *CreateTrafficQosPolicyRequest) (response *CreateTrafficQosPolicyResponse, err error) {
	if request == nil {
		request = NewCreateTrafficQosPolicyRequest()
	}
	response = NewCreateTrafficQosPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnPolicyBasedRoutingNextHopRequest() (request *DescribeCcnPolicyBasedRoutingNextHopRequest) {
	request = &DescribeCcnPolicyBasedRoutingNextHopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnPolicyBasedRoutingNextHop")
	return
}

func NewDescribeCcnPolicyBasedRoutingNextHopResponse() (response *DescribeCcnPolicyBasedRoutingNextHopResponse) {
	response = &DescribeCcnPolicyBasedRoutingNextHopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云联网策略路由下一跳
func (c *Client) DescribeCcnPolicyBasedRoutingNextHop(request *DescribeCcnPolicyBasedRoutingNextHopRequest) (response *DescribeCcnPolicyBasedRoutingNextHopResponse, err error) {
	if request == nil {
		request = NewDescribeCcnPolicyBasedRoutingNextHopRequest()
	}
	response = NewDescribeCcnPolicyBasedRoutingNextHopResponse()
	err = c.Send(request, response)
	return
}

func NewRejectAttachCcnInstancesRequest() (request *RejectAttachCcnInstancesRequest) {
	request = &RejectAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "RejectAttachCcnInstances")
	return
}

func NewRejectAttachCcnInstancesResponse() (response *RejectAttachCcnInstancesResponse) {
	response = &RejectAttachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RejectAttachCcnInstances）用于跨账号关联实例时，云联网所有者拒绝关联操作。
func (c *Client) RejectAttachCcnInstances(request *RejectAttachCcnInstancesRequest) (response *RejectAttachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewRejectAttachCcnInstancesRequest()
	}
	response = NewRejectAttachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest() (request *InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest) {
	request = &InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquirePriceRenewDirectConnectAccelerateChannelBandwidth")
	return
}

func NewInquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse() (response *InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse) {
	response = &InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费加速通道带宽询价
func (c *Client) InquirePriceRenewDirectConnectAccelerateChannelBandwidth(request *InquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest) (response *InquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse, err error) {
	if request == nil {
		request = NewInquirePriceRenewDirectConnectAccelerateChannelBandwidthRequest()
	}
	response = NewInquirePriceRenewDirectConnectAccelerateChannelBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
	request = &DescribeCcnAttachedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnAttachedInstances")
	return
}

func NewDescribeCcnAttachedInstancesResponse() (response *DescribeCcnAttachedInstancesResponse) {
	response = &DescribeCcnAttachedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnAttachedInstances）用于查询云联网实例下已关联的网络实例。
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnAttachedInstancesRequest()
	}
	response = NewDescribeCcnAttachedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpdateCcnBandwidthDealRequest() (request *GetUpdateCcnBandwidthDealRequest) {
	request = &GetUpdateCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "GetUpdateCcnBandwidthDeal")
	return
}

func NewGetUpdateCcnBandwidthDealResponse() (response *GetUpdateCcnBandwidthDealResponse) {
	response = &GetUpdateCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetUpdateCcnBandwidthDeal）用于获取变更云联网带宽的商品信息。
func (c *Client) GetUpdateCcnBandwidthDeal(request *GetUpdateCcnBandwidthDealRequest) (response *GetUpdateCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetUpdateCcnBandwidthDealRequest()
	}
	response = NewGetUpdateCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnIpv6RoutesRequest() (request *DescribeCcnIpv6RoutesRequest) {
	request = &DescribeCcnIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnIpv6Routes")
	return
}

func NewDescribeCcnIpv6RoutesResponse() (response *DescribeCcnIpv6RoutesResponse) {
	response = &DescribeCcnIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnIpv6Routes）用于查询已加入云联网（CCN）的Ipv6路由。
func (c *Client) DescribeCcnIpv6Routes(request *DescribeCcnIpv6RoutesRequest) (response *DescribeCcnIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnIpv6RoutesRequest()
	}
	response = NewDescribeCcnIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewBandwidthLimitForCcnAlarmOnlyRequest() (request *BandwidthLimitForCcnAlarmOnlyRequest) {
	request = &BandwidthLimitForCcnAlarmOnlyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "BandwidthLimitForCcnAlarmOnly")
	return
}

func NewBandwidthLimitForCcnAlarmOnlyResponse() (response *BandwidthLimitForCcnAlarmOnlyResponse) {
	response = &BandwidthLimitForCcnAlarmOnlyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云联网查看带宽上限监控告警接入专用
func (c *Client) BandwidthLimitForCcnAlarmOnly(request *BandwidthLimitForCcnAlarmOnlyRequest) (response *BandwidthLimitForCcnAlarmOnlyResponse, err error) {
	if request == nil {
		request = NewBandwidthLimitForCcnAlarmOnlyRequest()
	}
	response = NewBandwidthLimitForCcnAlarmOnlyResponse()
	err = c.Send(request, response)
	return
}

func NewAuditCrossBorderComplianceRequest() (request *AuditCrossBorderComplianceRequest) {
	request = &AuditCrossBorderComplianceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "AuditCrossBorderCompliance")
	return
}

func NewAuditCrossBorderComplianceResponse() (response *AuditCrossBorderComplianceResponse) {
	response = &AuditCrossBorderComplianceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AuditCrossBorderCompliance）用于服务商操作合规化资质审批。
// * 服务商只能操作提交到本服务商的审批单，后台会校验身份。即只授权给服务商的`APPID` 调用本接口。
// * `APPROVED` 状态的审批单，可以再次操作为 `DENY`；`DENY` 状态的审批单，也可以再次操作为 `APPROVED`。
func (c *Client) AuditCrossBorderCompliance(request *AuditCrossBorderComplianceRequest) (response *AuditCrossBorderComplianceResponse, err error) {
	if request == nil {
		request = NewAuditCrossBorderComplianceRequest()
	}
	response = NewAuditCrossBorderComplianceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTrafficQosPolicyRequest() (request *DeleteTrafficQosPolicyRequest) {
	request = &DeleteTrafficQosPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteTrafficQosPolicy")
	return
}

func NewDeleteTrafficQosPolicyResponse() (response *DeleteTrafficQosPolicyResponse) {
	response = &DeleteTrafficQosPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除流量调度策略
func (c *Client) DeleteTrafficQosPolicy(request *DeleteTrafficQosPolicyRequest) (response *DeleteTrafficQosPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteTrafficQosPolicyRequest()
	}
	response = NewDeleteTrafficQosPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnRouteTablesRequest() (request *ModifyCcnRouteTablesRequest) {
	request = &ModifyCcnRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnRouteTables")
	return
}

func NewModifyCcnRouteTablesResponse() (response *ModifyCcnRouteTablesResponse) {
	response = &ModifyCcnRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改云联网路由表名称和备注。
func (c *Client) ModifyCcnRouteTables(request *ModifyCcnRouteTablesRequest) (response *ModifyCcnRouteTablesResponse, err error) {
	if request == nil {
		request = NewModifyCcnRouteTablesRequest()
	}
	response = NewModifyCcnRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnRouteMatchRuleRequest() (request *ModifyCcnRouteMatchRuleRequest) {
	request = &ModifyCcnRouteMatchRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ModifyCcnRouteMatchRule")
	return
}

func NewModifyCcnRouteMatchRuleResponse() (response *ModifyCcnRouteMatchRuleResponse) {
	response = &ModifyCcnRouteMatchRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改路由匹配规则
func (c *Client) ModifyCcnRouteMatchRule(request *ModifyCcnRouteMatchRuleRequest) (response *ModifyCcnRouteMatchRuleResponse, err error) {
	if request == nil {
		request = NewModifyCcnRouteMatchRuleRequest()
	}
	response = NewModifyCcnRouteMatchRuleResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDirectConnectPortMapRequest() (request *ReplaceDirectConnectPortMapRequest) {
	request = &ReplaceDirectConnectPortMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ReplaceDirectConnectPortMap")
	return
}

func NewReplaceDirectConnectPortMapResponse() (response *ReplaceDirectConnectPortMapResponse) {
	response = &ReplaceDirectConnectPortMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置端口映射接口
func (c *Client) ReplaceDirectConnectPortMap(request *ReplaceDirectConnectPortMapRequest) (response *ReplaceDirectConnectPortMapResponse, err error) {
	if request == nil {
		request = NewReplaceDirectConnectPortMapRequest()
	}
	response = NewReplaceDirectConnectPortMapResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRegionBandwidthLimitsRequest() (request *DeleteCcnRegionBandwidthLimitsRequest) {
	request = &DeleteCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteCcnRegionBandwidthLimits")
	return
}

func NewDeleteCcnRegionBandwidthLimitsResponse() (response *DeleteCcnRegionBandwidthLimitsResponse) {
	response = &DeleteCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCcnRegionBandwidthLimits）用于删除云联网实例的带宽配置。
func (c *Client) DeleteCcnRegionBandwidthLimits(request *DeleteCcnRegionBandwidthLimitsRequest) (response *DeleteCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewDeleteCcnRegionBandwidthLimitsRequest()
	}
	response = NewDeleteCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewClearRouteTableSelectionPoliciesRequest() (request *ClearRouteTableSelectionPoliciesRequest) {
	request = &ClearRouteTableSelectionPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ClearRouteTableSelectionPolicies")
	return
}

func NewClearRouteTableSelectionPoliciesResponse() (response *ClearRouteTableSelectionPoliciesResponse) {
	response = &ClearRouteTableSelectionPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ClearRouteTableSelectionPolicies）用于清空指定云联网的路由表选择策略。
func (c *Client) ClearRouteTableSelectionPolicies(request *ClearRouteTableSelectionPoliciesRequest) (response *ClearRouteTableSelectionPoliciesResponse, err error) {
	if request == nil {
		request = NewClearRouteTableSelectionPoliciesRequest()
	}
	response = NewClearRouteTableSelectionPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDirectConnectAccelerateChannelRequest() (request *DeleteDirectConnectAccelerateChannelRequest) {
	request = &DeleteDirectConnectAccelerateChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DeleteDirectConnectAccelerateChannel")
	return
}

func NewDeleteDirectConnectAccelerateChannelResponse() (response *DeleteDirectConnectAccelerateChannelResponse) {
	response = &DeleteDirectConnectAccelerateChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除加速通道
func (c *Client) DeleteDirectConnectAccelerateChannel(request *DeleteDirectConnectAccelerateChannelRequest) (response *DeleteDirectConnectAccelerateChannelResponse, err error) {
	if request == nil {
		request = NewDeleteDirectConnectAccelerateChannelRequest()
	}
	response = NewDeleteDirectConnectAccelerateChannelResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateInstancesToCcnRouteTableRequest() (request *AssociateInstancesToCcnRouteTableRequest) {
	request = &AssociateInstancesToCcnRouteTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "AssociateInstancesToCcnRouteTable")
	return
}

func NewAssociateInstancesToCcnRouteTableResponse() (response *AssociateInstancesToCcnRouteTableResponse) {
	response = &AssociateInstancesToCcnRouteTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssociateInstancesToCcnRouteTable）用于将指定的云联网实例关联到指定的云联网路由表。
func (c *Client) AssociateInstancesToCcnRouteTable(request *AssociateInstancesToCcnRouteTableRequest) (response *AssociateInstancesToCcnRouteTableResponse, err error) {
	if request == nil {
		request = NewAssociateInstancesToCcnRouteTableRequest()
	}
	response = NewAssociateInstancesToCcnRouteTableResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnPolicyBasedRoutingRuleRequest() (request *DescribeCcnPolicyBasedRoutingRuleRequest) {
	request = &DescribeCcnPolicyBasedRoutingRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnPolicyBasedRoutingRule")
	return
}

func NewDescribeCcnPolicyBasedRoutingRuleResponse() (response *DescribeCcnPolicyBasedRoutingRuleResponse) {
	response = &DescribeCcnPolicyBasedRoutingRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云联网策略路由匹配规则
func (c *Client) DescribeCcnPolicyBasedRoutingRule(request *DescribeCcnPolicyBasedRoutingRuleRequest) (response *DescribeCcnPolicyBasedRoutingRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCcnPolicyBasedRoutingRuleRequest()
	}
	response = NewDescribeCcnPolicyBasedRoutingRuleResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCcnIpv6RoutesRequest() (request *EnableCcnIpv6RoutesRequest) {
	request = &EnableCcnIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "EnableCcnIpv6Routes")
	return
}

func NewEnableCcnIpv6RoutesResponse() (response *EnableCcnIpv6RoutesResponse) {
	response = &EnableCcnIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（EnableCcnIpv6Routes）用于启用已经加入云联网（CCN）的Ipv6路由。<br />
// 本接口会校验启用后，是否与已有路由冲突，如果冲突，则无法启用，失败处理。路由冲突时，需要先禁用与之冲突的路由，才能启用该路由。
func (c *Client) EnableCcnIpv6Routes(request *EnableCcnIpv6RoutesRequest) (response *EnableCcnIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewEnableCcnIpv6RoutesRequest()
	}
	response = NewEnableCcnIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDirectConnectAccelerateChannelRequest() (request *CreateDirectConnectAccelerateChannelRequest) {
	request = &CreateDirectConnectAccelerateChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "CreateDirectConnectAccelerateChannel")
	return
}

func NewCreateDirectConnectAccelerateChannelResponse() (response *CreateDirectConnectAccelerateChannelResponse) {
	response = &CreateDirectConnectAccelerateChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建加速通道
func (c *Client) CreateDirectConnectAccelerateChannel(request *CreateDirectConnectAccelerateChannelRequest) (response *CreateDirectConnectAccelerateChannelResponse, err error) {
	if request == nil {
		request = NewCreateDirectConnectAccelerateChannelRequest()
	}
	response = NewCreateDirectConnectAccelerateChannelResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceCcnRouteTableBroadcastPolicysRequest() (request *ReplaceCcnRouteTableBroadcastPolicysRequest) {
	request = &ReplaceCcnRouteTableBroadcastPolicysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "ReplaceCcnRouteTableBroadcastPolicys")
	return
}

func NewReplaceCcnRouteTableBroadcastPolicysResponse() (response *ReplaceCcnRouteTableBroadcastPolicysResponse) {
	response = &ReplaceCcnRouteTableBroadcastPolicysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ReplaceCcnRouteTableBroadcastPolicys)用于替换云联网路由表路由传播策略。
// > 特别注意：是全量覆盖，非增量添加
//
// **路由条件支持以下四种：**
//
// - 实例类型: `instance-type`，可选值：私有网络 `VPC`、专线网关 `DIRECTCONNECT`、VPN网关 `VPNGW`
// - 实例ID: `instance-id`，例如：`dcg-8zljkrft`、`vpc-jdevjrup`，暂不支持 `Edge` 实例
// - 实例地域: `instance-region`，例如：`region1`<br />产品支持的所有地域列表可通过接口 [DescribeRegions](https://{{conf.main_domain}}/document/product/1596/77930) 查询，其中参数 `Product` 设置为 `ccn`
// - 路由前缀: `cidr-block`，例如：`10.1.0.0/16`
//
// **传播条件支持以下三种：**
//
// - 实例类型: `instance-type`，格式同路由条件
// - 实例ID: `instance-id`，格式同路由条件
// - 实例地域: `instance-region`，格式同路由条件
//
// **使用限制：**
// - 一条策略内的单个条件类型，最大支持设置 `25` 个条件值
// - 一张路由表，最大支持 `100` 条路由传播策略
// - 路由条件类型中，只有 `cidr-block` 类型支持模糊匹配和精确匹配两种，其它类型只支持精确匹配一种模式
func (c *Client) ReplaceCcnRouteTableBroadcastPolicys(request *ReplaceCcnRouteTableBroadcastPolicysRequest) (response *ReplaceCcnRouteTableBroadcastPolicysResponse, err error) {
	if request == nil {
		request = NewReplaceCcnRouteTableBroadcastPolicysRequest()
	}
	response = NewReplaceCcnRouteTableBroadcastPolicysResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnRegionBandwidthLimitsRequest() (request *SetCcnRegionBandwidthLimitsRequest) {
	request = &SetCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetCcnRegionBandwidthLimits")
	return
}

func NewSetCcnRegionBandwidthLimitsResponse() (response *SetCcnRegionBandwidthLimitsResponse) {
	response = &SetCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SetCcnRegionBandwidthLimits）用于设置云联网（CCN）各地域出带宽上限，或者地域间带宽上限。
func (c *Client) SetCcnRegionBandwidthLimits(request *SetCcnRegionBandwidthLimitsRequest) (response *SetCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewSetCcnRegionBandwidthLimitsRequest()
	}
	response = NewSetCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnRegionDefaultQosBandwidthLimitRequest() (request *SetCcnRegionDefaultQosBandwidthLimitRequest) {
	request = &SetCcnRegionDefaultQosBandwidthLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetCcnRegionDefaultQosBandwidthLimit")
	return
}

func NewSetCcnRegionDefaultQosBandwidthLimitResponse() (response *SetCcnRegionDefaultQosBandwidthLimitResponse) {
	response = &SetCcnRegionDefaultQosBandwidthLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置默认QOS带宽
func (c *Client) SetCcnRegionDefaultQosBandwidthLimit(request *SetCcnRegionDefaultQosBandwidthLimitRequest) (response *SetCcnRegionDefaultQosBandwidthLimitResponse, err error) {
	if request == nil {
		request = NewSetCcnRegionDefaultQosBandwidthLimitRequest()
	}
	response = NewSetCcnRegionDefaultQosBandwidthLimitResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest() (request *InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest) {
	request = &InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "InquirePriceCreateDirectConnectAccelerateChannelBandwidth")
	return
}

func NewInquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse() (response *InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse) {
	response = &InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 购买加速通道带宽询价。
func (c *Client) InquirePriceCreateDirectConnectAccelerateChannelBandwidth(request *InquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest) (response *InquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse, err error) {
	if request == nil {
		request = NewInquirePriceCreateDirectConnectAccelerateChannelBandwidthRequest()
	}
	response = NewInquirePriceCreateDirectConnectAccelerateChannelBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewLockCcnsRequest() (request *LockCcnsRequest) {
	request = &LockCcnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "LockCcns")
	return
}

func NewLockCcnsResponse() (response *LockCcnsResponse) {
	response = &LockCcnsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（LockCcns）用于锁定云联网实例
//
// 该接口一般用来封禁出口限速的云联网实例, 目前联通内部运营系统通过云API调用, 因为出口限速无法按地域间封禁, 只能按更粗的云联网实例粒度封禁, 如果是地域间限速, 一般可以通过更细的限速实例粒度封禁（LockCcnBandwidths）
//
// 如有需要, 可以封禁任意限速实例, 可接入到内部运营系统
func (c *Client) LockCcns(request *LockCcnsRequest) (response *LockCcnsResponse, err error) {
	if request == nil {
		request = NewLockCcnsRequest()
	}
	response = NewLockCcnsResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnDefaultBillingModeRequest() (request *SetCcnDefaultBillingModeRequest) {
	request = &SetCcnDefaultBillingModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "SetCcnDefaultBillingMode")
	return
}

func NewSetCcnDefaultBillingModeResponse() (response *SetCcnDefaultBillingModeResponse) {
	response = &SetCcnDefaultBillingModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复默认计费模式-联通专属接口。
func (c *Client) SetCcnDefaultBillingMode(request *SetCcnDefaultBillingModeRequest) (response *SetCcnDefaultBillingModeResponse, err error) {
	if request == nil {
		request = NewSetCcnDefaultBillingModeRequest()
	}
	response = NewSetCcnDefaultBillingModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRouteMatchRuleRequest() (request *DescribeCcnRouteMatchRuleRequest) {
	request = &DescribeCcnRouteMatchRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeCcnRouteMatchRule")
	return
}

func NewDescribeCcnRouteMatchRuleResponse() (response *DescribeCcnRouteMatchRuleResponse) {
	response = &DescribeCcnRouteMatchRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由匹配规则
func (c *Client) DescribeCcnRouteMatchRule(request *DescribeCcnRouteMatchRuleRequest) (response *DescribeCcnRouteMatchRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRouteMatchRuleRequest()
	}
	response = NewDescribeCcnRouteMatchRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteReceivingPoliciesRequest() (request *DescribeRouteReceivingPoliciesRequest) {
	request = &DescribeRouteReceivingPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccn", APIVersion, "DescribeRouteReceivingPolicies")
	return
}

func NewDescribeRouteReceivingPoliciesResponse() (response *DescribeRouteReceivingPoliciesResponse) {
	response = &DescribeRouteReceivingPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云联网路由表接收策略。
func (c *Client) DescribeRouteReceivingPolicies(request *DescribeRouteReceivingPoliciesRequest) (response *DescribeRouteReceivingPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeRouteReceivingPoliciesRequest()
	}
	response = NewDescribeRouteReceivingPoliciesResponse()
	err = c.Send(request, response)
	return
}

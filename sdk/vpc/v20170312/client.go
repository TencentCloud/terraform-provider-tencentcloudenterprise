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

func NewEnableVpcPeeringConnectionExRequest() (request *EnableVpcPeeringConnectionExRequest) {
	request = &EnableVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableVpcPeeringConnectionEx")
	return
}

func NewEnableVpcPeeringConnectionExResponse() (response *EnableVpcPeeringConnectionExResponse) {
	response = &EnableVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用过期跨地域对等连接
func (c *Client) EnableVpcPeeringConnectionEx(request *EnableVpcPeeringConnectionExRequest) (response *EnableVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewEnableVpcPeeringConnectionExRequest()
	}
	response = NewEnableVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
	request = &CloneSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CloneSecurityGroup")
	return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
	response = &CloneSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CloneSecurityGroup）用于根据存量的安全组，克隆创建出同样规则配置的安全组。仅克隆安全组及其规则信息，不会克隆安全组标签信息。
func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) (response *CloneSecurityGroupResponse, err error) {
	if request == nil {
		request = NewCloneSecurityGroupRequest()
	}
	response = NewCloneSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnRequest() (request *CreateCcnRequest) {
	request = &CreateCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateCcn")
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
// 每个账号能创建的云联网实例个数是有限的，详请参考产品文档。如果需要扩充请联系在线客服。
func (c *Client) CreateCcn(request *CreateCcnRequest) (response *CreateCcnResponse, err error) {
	if request == nil {
		request = NewCreateCcnRequest()
	}
	response = NewCreateCcnResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpv6RoutesRequest() (request *DeleteIpv6RoutesRequest) {
	request = &DeleteIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteIpv6Routes")
	return
}

func NewDeleteIpv6RoutesResponse() (response *DeleteIpv6RoutesResponse) {
	response = &DeleteIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由表的IPv6路由规则。
func (c *Client) DeleteIpv6Routes(request *DeleteIpv6RoutesRequest) (response *DeleteIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewDeleteIpv6RoutesRequest()
	}
	response = NewDeleteIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest() (request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) {
	request = &DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewayDestinationIpPortTranslationNatRules")
	return
}

func NewDescribeNatGatewayDestinationIpPortTranslationNatRulesResponse() (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) {
	response = &DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNatGatewayDestinationIpPortTranslationNatRules）用于查询NAT网关端口转发规则对象数组。
func (c *Client) DescribeNatGatewayDestinationIpPortTranslationNatRules(request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
	}
	response = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateVpcEndPointSecurityGroupsRequest() (request *DisassociateVpcEndPointSecurityGroupsRequest) {
	request = &DisassociateVpcEndPointSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateVpcEndPointSecurityGroups")
	return
}

func NewDisassociateVpcEndPointSecurityGroupsResponse() (response *DisassociateVpcEndPointSecurityGroupsResponse) {
	response = &DisassociateVpcEndPointSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 终端节点解绑安全组。
func (c *Client) DisassociateVpcEndPointSecurityGroups(request *DisassociateVpcEndPointSecurityGroupsRequest) (response *DisassociateVpcEndPointSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDisassociateVpcEndPointSecurityGroupsRequest()
	}
	response = NewDisassociateVpcEndPointSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
	request = &ModifyVpcAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcAttribute")
	return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
	response = &ModifyVpcAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyVpcAttribute）用于修改私有网络（VPC）的相关属性。
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
	if request == nil {
		request = NewModifyVpcAttributeRequest()
	}
	response = NewModifyVpcAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressesHistoryRequest() (request *DescribeAddressesHistoryRequest) {
	request = &DescribeAddressesHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressesHistory")
	return
}

func NewDescribeAddressesHistoryResponse() (response *DescribeAddressesHistoryResponse) {
	response = &DescribeAddressesHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询账户以往使用过但是当前已释放的eip记录
func (c *Client) DescribeAddressesHistory(request *DescribeAddressesHistoryRequest) (response *DescribeAddressesHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeAddressesHistoryRequest()
	}
	response = NewDescribeAddressesHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyAddressesBandwidthRequest() (request *InquiryPriceModifyAddressesBandwidthRequest) {
	request = &InquiryPriceModifyAddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceModifyAddressesBandwidth")
	return
}

func NewInquiryPriceModifyAddressesBandwidthResponse() (response *InquiryPriceModifyAddressesBandwidthResponse) {
	response = &InquiryPriceModifyAddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改带宽询价
func (c *Client) InquiryPriceModifyAddressesBandwidth(request *InquiryPriceModifyAddressesBandwidthRequest) (response *InquiryPriceModifyAddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyAddressesBandwidthRequest()
	}
	response = NewInquiryPriceModifyAddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewayQuotaRequest() (request *DescribeVpnGatewayQuotaRequest) {
	request = &DescribeVpnGatewayQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewayQuota")
	return
}

func NewDescribeVpnGatewayQuotaResponse() (response *DescribeVpnGatewayQuotaResponse) {
	response = &DescribeVpnGatewayQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPN网关配额
func (c *Client) DescribeVpnGatewayQuota(request *DescribeVpnGatewayQuotaRequest) (response *DescribeVpnGatewayQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewayQuotaRequest()
	}
	response = NewDescribeVpnGatewayQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrafficMirrorAttributeRequest() (request *ModifyTrafficMirrorAttributeRequest) {
	request = &ModifyTrafficMirrorAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyTrafficMirrorAttribute")
	return
}

func NewModifyTrafficMirrorAttributeResponse() (response *ModifyTrafficMirrorAttributeResponse) {
	response = &ModifyTrafficMirrorAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyTrafficMirrorAttribute）用于修改流量镜像实例属性。
// 注意：只支持修改名字和描述信息
func (c *Client) ModifyTrafficMirrorAttribute(request *ModifyTrafficMirrorAttributeRequest) (response *ModifyTrafficMirrorAttributeResponse, err error) {
	if request == nil {
		request = NewModifyTrafficMirrorAttributeRequest()
	}
	response = NewModifyTrafficMirrorAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkAclRequest() (request *DeleteNetworkAclRequest) {
	request = &DeleteNetworkAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkAcl")
	return
}

func NewDeleteNetworkAclResponse() (response *DeleteNetworkAclResponse) {
	response = &DeleteNetworkAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNetworkAcl）用于删除网络ACL。
func (c *Client) DeleteNetworkAcl(request *DeleteNetworkAclRequest) (response *DeleteNetworkAclResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkAclRequest()
	}
	response = NewDeleteNetworkAclResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssistantCidrRequest() (request *DescribeAssistantCidrRequest) {
	request = &DescribeAssistantCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAssistantCidr")
	return
}

func NewDescribeAssistantCidrResponse() (response *DescribeAssistantCidrResponse) {
	response = &DescribeAssistantCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAssistantCidr）用于查询辅助CIDR列表。（接口灰度中，如需使用请提工单。）
func (c *Client) DescribeAssistantCidr(request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
	if request == nil {
		request = NewDescribeAssistantCidrRequest()
	}
	response = NewDescribeAssistantCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDisableRoutesRequest() (request *DisableRoutesRequest) {
	request = &DisableRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisableRoutes")
	return
}

func NewDisableRoutesResponse() (response *DisableRoutesResponse) {
	response = &DisableRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableRoutes）用于禁用已启用的子网路由
func (c *Client) DisableRoutes(request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
	if request == nil {
		request = NewDisableRoutesRequest()
	}
	response = NewDisableRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcEndPointServiceWhiteListRequest() (request *DeleteVpcEndPointServiceWhiteListRequest) {
	request = &DeleteVpcEndPointServiceWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPointServiceWhiteList")
	return
}

func NewDeleteVpcEndPointServiceWhiteListResponse() (response *DeleteVpcEndPointServiceWhiteListResponse) {
	response = &DeleteVpcEndPointServiceWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除终端节点服务白名单。
func (c *Client) DeleteVpcEndPointServiceWhiteList(request *DeleteVpcEndPointServiceWhiteListRequest) (response *DeleteVpcEndPointServiceWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteVpcEndPointServiceWhiteListRequest()
	}
	response = NewDeleteVpcEndPointServiceWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewayQuotaRequest() (request *DescribeNatGatewayQuotaRequest) {
	request = &DescribeNatGatewayQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewayQuota")
	return
}

func NewDescribeNatGatewayQuotaResponse() (response *DescribeNatGatewayQuotaResponse) {
	response = &DescribeNatGatewayQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNatGatewayQuota）用于查询 NAT 网关配额。
func (c *Client) DescribeNatGatewayQuota(request *DescribeNatGatewayQuotaRequest) (response *DescribeNatGatewayQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewayQuotaRequest()
	}
	response = NewDescribeNatGatewayQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewaySSLServersRequest() (request *DescribeVpnGatewaySSLServersRequest) {
	request = &DescribeVpnGatewaySSLServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewaySSLServers")
	return
}

func NewDescribeVpnGatewaySSLServersResponse() (response *DescribeVpnGatewaySSLServersResponse) {
	response = &DescribeVpnGatewaySSLServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeVpnGatewaySSLServers
func (c *Client) DescribeVpnGatewaySSLServers(request *DescribeVpnGatewaySSLServersRequest) (response *DescribeVpnGatewaySSLServersResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewaySSLServersRequest()
	}
	response = NewDescribeVpnGatewaySSLServersResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadVpnGatewaySSLClientCertRequest() (request *DownloadVpnGatewaySSLClientCertRequest) {
	request = &DownloadVpnGatewaySSLClientCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DownloadVpnGatewaySSLClientCert")
	return
}

func NewDownloadVpnGatewaySSLClientCertResponse() (response *DownloadVpnGatewaySSLClientCertResponse) {
	response = &DownloadVpnGatewaySSLClientCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DownloadVpnGatewaySSLClientCert
func (c *Client) DownloadVpnGatewaySSLClientCert(request *DownloadVpnGatewaySSLClientCertRequest) (response *DownloadVpnGatewaySSLClientCertResponse, err error) {
	if request == nil {
		request = NewDownloadVpnGatewaySSLClientCertRequest()
	}
	response = NewDownloadVpnGatewaySSLClientCertResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
	request = &CreateNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkInterface")
	return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
	response = &CreateNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateNetworkInterface）用于创建弹性网卡。
// * 创建弹性网卡时可以指定内网IP，并且可以指定一个主IP，指定的内网IP必须在弹性网卡所在子网内，而且不能被占用。
// * 创建弹性网卡时可以指定需要申请的内网IP数量，系统会随机生成内网IP地址。
// * 一个弹性网卡支持绑定的IP地址是有限制的，更多资源限制信息详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
// * 创建弹性网卡同时可以绑定已有安全组。
// * 创建弹性网卡同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewCreateNetworkInterfaceRequest()
	}
	response = NewCreateNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) {
	request = &DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGatewayDestinationIpPortTranslationNatRule")
	return
}

func NewDeleteNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) {
	response = &DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNatGatewayDestinationIpPortTranslationNatRule）用于删除NAT网关端口转发规则。
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRule(request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerGatewayVendorsRequest() (request *DescribeCustomerGatewayVendorsRequest) {
	request = &DescribeCustomerGatewayVendorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCustomerGatewayVendors")
	return
}

func NewDescribeCustomerGatewayVendorsResponse() (response *DescribeCustomerGatewayVendorsResponse) {
	response = &DescribeCustomerGatewayVendorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCustomerGatewayVendors）用于查询可支持的对端网关厂商信息。
func (c *Client) DescribeCustomerGatewayVendors(request *DescribeCustomerGatewayVendorsRequest) (response *DescribeCustomerGatewayVendorsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerGatewayVendorsRequest()
	}
	response = NewDescribeCustomerGatewayVendorsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIp6IdcInfoRequest() (request *DescribeIp6IdcInfoRequest) {
	request = &DescribeIp6IdcInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIp6IdcInfo")
	return
}

func NewDescribeIp6IdcInfoResponse() (response *DescribeIp6IdcInfoResponse) {
	response = &DescribeIp6IdcInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IPV6创建时的IDC信息
func (c *Client) DescribeIp6IdcInfo(request *DescribeIp6IdcInfoRequest) (response *DescribeIp6IdcInfoResponse, err error) {
	if request == nil {
		request = NewDescribeIp6IdcInfoRequest()
	}
	response = NewDescribeIp6IdcInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDisableSSLVpnClientCertRequest() (request *DisableSSLVpnClientCertRequest) {
	request = &DisableSSLVpnClientCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisableSSLVpnClientCert")
	return
}

func NewDisableSSLVpnClientCertResponse() (response *DisableSSLVpnClientCertResponse) {
	response = &DisableSSLVpnClientCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DisableSSLVpnClientCert
func (c *Client) DisableSSLVpnClientCert(request *DisableSSLVpnClientCertRequest) (response *DisableSSLVpnClientCertResponse, err error) {
	if request == nil {
		request = NewDisableSSLVpnClientCertRequest()
	}
	response = NewDisableSSLVpnClientCertResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomerGatewayRequest() (request *DeleteCustomerGatewayRequest) {
	request = &DeleteCustomerGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteCustomerGateway")
	return
}

func NewDeleteCustomerGatewayResponse() (response *DeleteCustomerGatewayResponse) {
	response = &DeleteCustomerGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCustomerGateway）用于删除对端网关。
func (c *Client) DeleteCustomerGateway(request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteCustomerGatewayRequest()
	}
	response = NewDeleteCustomerGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpnConnectionRequest() (request *DeleteVpnConnectionRequest) {
	request = &DeleteVpnConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnConnection")
	return
}

func NewDeleteVpnConnectionResponse() (response *DeleteVpnConnectionResponse) {
	response = &DeleteVpnConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteVpnConnection)用于删除VPN通道。
func (c *Client) DeleteVpnConnection(request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
	if request == nil {
		request = NewDeleteVpnConnectionRequest()
	}
	response = NewDeleteVpnConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLocalIpTranslationAclRuleRequest() (request *CreateLocalIpTranslationAclRuleRequest) {
	request = &CreateLocalIpTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalIpTranslationAclRule")
	return
}

func NewCreateLocalIpTranslationAclRuleResponse() (response *CreateLocalIpTranslationAclRuleResponse) {
	response = &CreateLocalIpTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关本端IP转换ACL规则
func (c *Client) CreateLocalIpTranslationAclRule(request *CreateLocalIpTranslationAclRuleRequest) (response *CreateLocalIpTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewCreateLocalIpTranslationAclRuleRequest()
	}
	response = NewCreateLocalIpTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpnGatewaySSLServerRequest() (request *CreateVpnGatewaySSLServerRequest) {
	request = &CreateVpnGatewaySSLServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGatewaySSLServer")
	return
}

func NewCreateVpnGatewaySSLServerResponse() (response *CreateVpnGatewaySSLServerResponse) {
	response = &CreateVpnGatewaySSLServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建sslvpn服务
func (c *Client) CreateVpnGatewaySSLServer(request *CreateVpnGatewaySSLServerRequest) (response *CreateVpnGatewaySSLServerResponse, err error) {
	if request == nil {
		request = NewCreateVpnGatewaySSLServerRequest()
	}
	response = NewCreateVpnGatewaySSLServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
	request = &DescribeFlowLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeFlowLogs")
	return
}

func NewDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
	response = &DescribeFlowLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeFlowLogs）用于查询获取流日志集合
func (c *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowLogsRequest()
	}
	response = NewDescribeFlowLogsResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateNetworkInterfaceRequest() (request *MigrateNetworkInterfaceRequest) {
	request = &MigrateNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "MigrateNetworkInterface")
	return
}

func NewMigrateNetworkInterfaceResponse() (response *MigrateNetworkInterfaceResponse) {
	response = &MigrateNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（MigrateNetworkInterface）用于弹性网卡迁移。
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewMigrateNetworkInterfaceRequest()
	}
	response = NewMigrateNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewMigratePrivateIpAddressRequest() (request *MigratePrivateIpAddressRequest) {
	request = &MigratePrivateIpAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "MigratePrivateIpAddress")
	return
}

func NewMigratePrivateIpAddressResponse() (response *MigratePrivateIpAddressResponse) {
	response = &MigratePrivateIpAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//	本接口（MigratePrivateIpAddress）用于弹性网卡内网IP迁移。
//
// * 该接口用于将一个内网IP从一个弹性网卡上迁移到另外一个弹性网卡，主IP地址不支持迁移。
// * 迁移前后的弹性网卡必须在同一个子网内。
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
	if request == nil {
		request = NewMigratePrivateIpAddressRequest()
	}
	response = NewMigratePrivateIpAddressResponse()
	err = c.Send(request, response)
	return
}

func NewCheckGatewayFlowMonitorRequest() (request *CheckGatewayFlowMonitorRequest) {
	request = &CheckGatewayFlowMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckGatewayFlowMonitor")
	return
}

func NewCheckGatewayFlowMonitorResponse() (response *CheckGatewayFlowMonitorResponse) {
	response = &CheckGatewayFlowMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckGatewayFlowMonitor）用于查询网关是否启用流量监控。
func (c *Client) CheckGatewayFlowMonitor(request *CheckGatewayFlowMonitorRequest) (response *CheckGatewayFlowMonitorResponse, err error) {
	if request == nil {
		request = NewCheckGatewayFlowMonitorRequest()
	}
	response = NewCheckGatewayFlowMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
	request = &CreateVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpc")
	return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
	response = &CreateVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateVpc)用于创建私有网络(VPC)。
// * 用户可以创建的最小网段子网掩码为28（有16个IP地址），最大网段子网掩码为16（65,536个IP地址）,如果规划VPC网段请参见VPC网段规划说明。
// * 同一个地域能创建的VPC资源个数也是有限制的，详见 <a href="https://cloud.tencent.com/doc/product/215/537" title="VPC使用限制">VPC使用限制</a>,如果需要扩充请联系在线客服。
// * 创建VPC同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
	if request == nil {
		request = NewCreateVpcRequest()
	}
	response = NewCreateVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcInstancesRequest() (request *DescribeVpcInstancesRequest) {
	request = &DescribeVpcInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcInstances")
	return
}

func NewDescribeVpcInstancesResponse() (response *DescribeVpcInstancesResponse) {
	response = &DescribeVpcInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcInstances）用于查询VPC下的云主机实例列表。
func (c *Client) DescribeVpcInstances(request *DescribeVpcInstancesRequest) (response *DescribeVpcInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeVpcInstancesRequest()
	}
	response = NewDescribeVpcInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewayZonesRequest() (request *DescribeNatGatewayZonesRequest) {
	request = &DescribeNatGatewayZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewayZones")
	return
}

func NewDescribeNatGatewayZonesResponse() (response *DescribeNatGatewayZonesResponse) {
	response = &DescribeNatGatewayZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询NAT可以使用的可用区
func (c *Client) DescribeNatGatewayZones(request *DescribeNatGatewayZonesRequest) (response *DescribeNatGatewayZonesResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewayZonesRequest()
	}
	response = NewDescribeNatGatewayZonesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGatewayFlowQosRequest() (request *ModifyGatewayFlowQosRequest) {
	request = &ModifyGatewayFlowQosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyGatewayFlowQos")
	return
}

func NewModifyGatewayFlowQosResponse() (response *ModifyGatewayFlowQosResponse) {
	response = &ModifyGatewayFlowQosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyGatewayFlowQos）用于调整网关流控带宽。
func (c *Client) ModifyGatewayFlowQos(request *ModifyGatewayFlowQosRequest) (response *ModifyGatewayFlowQosResponse, err error) {
	if request == nil {
		request = NewModifyGatewayFlowQosRequest()
	}
	response = NewModifyGatewayFlowQosResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLocalSourceIpPortTranslationNatRuleRequest() (request *DeleteLocalSourceIpPortTranslationNatRuleRequest) {
	request = &DeleteLocalSourceIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalSourceIpPortTranslationNatRule")
	return
}

func NewDeleteLocalSourceIpPortTranslationNatRuleResponse() (response *DeleteLocalSourceIpPortTranslationNatRuleResponse) {
	response = &DeleteLocalSourceIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关本端源IP端口转换
func (c *Client) DeleteLocalSourceIpPortTranslationNatRule(request *DeleteLocalSourceIpPortTranslationNatRuleRequest) (response *DeleteLocalSourceIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLocalSourceIpPortTranslationNatRuleRequest()
	}
	response = NewDeleteLocalSourceIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewVpnGatewayRequest() (request *InquiryPriceRenewVpnGatewayRequest) {
	request = &InquiryPriceRenewVpnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewVpnGateway")
	return
}

func NewInquiryPriceRenewVpnGatewayResponse() (response *InquiryPriceRenewVpnGatewayResponse) {
	response = &InquiryPriceRenewVpnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceRenewVpnGateway）用于续费VPN网关询价。目前仅支持IPSEC类型网关的询价。
func (c *Client) InquiryPriceRenewVpnGateway(request *InquiryPriceRenewVpnGatewayRequest) (response *InquiryPriceRenewVpnGatewayResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewVpnGatewayRequest()
	}
	response = NewInquiryPriceRenewVpnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResetVpnGatewayInternetMaxBandwidthRequest() (request *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) {
	request = &InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceResetVpnGatewayInternetMaxBandwidth")
	return
}

func NewInquiryPriceResetVpnGatewayInternetMaxBandwidthResponse() (response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) {
	response = &InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceResetVpnGatewayInternetMaxBandwidth）调整VPN网关带宽上限询价。
func (c *Client) InquiryPriceResetVpnGatewayInternetMaxBandwidth(request *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) (response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResetVpnGatewayInternetMaxBandwidthRequest()
	}
	response = NewInquiryPriceResetVpnGatewayInternetMaxBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
	request = &ModifySubnetAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifySubnetAttribute")
	return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
	response = &ModifySubnetAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySubnetAttribute）用于修改子网属性。
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
	if request == nil {
		request = NewModifySubnetAttributeRequest()
	}
	response = NewModifySubnetAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewStopTrafficMirrorRequest() (request *StopTrafficMirrorRequest) {
	request = &StopTrafficMirrorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "StopTrafficMirror")
	return
}

func NewStopTrafficMirrorResponse() (response *StopTrafficMirrorResponse) {
	response = &StopTrafficMirrorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StopTrafficMirror）用于关闭流量镜像实例。
func (c *Client) StopTrafficMirror(request *StopTrafficMirrorRequest) (response *StopTrafficMirrorResponse, err error) {
	if request == nil {
		request = NewStopTrafficMirrorRequest()
	}
	response = NewStopTrafficMirrorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBandwidthAttributeRequest() (request *DescribeBandwidthAttributeRequest) {
	request = &DescribeBandwidthAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthAttribute")
	return
}

func NewDescribeBandwidthAttributeResponse() (response *DescribeBandwidthAttributeResponse) {
	response = &DescribeBandwidthAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询转化的带宽属性。
func (c *Client) DescribeBandwidthAttribute(request *DescribeBandwidthAttributeRequest) (response *DescribeBandwidthAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeBandwidthAttributeRequest()
	}
	response = NewDescribeBandwidthAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceAllocateAddressesRequest() (request *InquiryPriceAllocateAddressesRequest) {
	request = &InquiryPriceAllocateAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceAllocateAddresses")
	return
}

func NewInquiryPriceAllocateAddressesResponse() (response *InquiryPriceAllocateAddressesResponse) {
	response = &InquiryPriceAllocateAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceAllocateAddresses(request *InquiryPriceAllocateAddressesRequest) (response *InquiryPriceAllocateAddressesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceAllocateAddressesRequest()
	}
	response = NewInquiryPriceAllocateAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnBandwidthRenewFlagRequest() (request *SetCcnBandwidthRenewFlagRequest) {
	request = &SetCcnBandwidthRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "SetCcnBandwidthRenewFlag")
	return
}

func NewSetCcnBandwidthRenewFlagResponse() (response *SetCcnBandwidthRenewFlagResponse) {
	response = &SetCcnBandwidthRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SetCcnBandwidthRenewFlag）用于设置预付费云联网实例地域间限速的自动续费标记
func (c *Client) SetCcnBandwidthRenewFlag(request *SetCcnBandwidthRenewFlagRequest) (response *SetCcnBandwidthRenewFlagResponse, err error) {
	if request == nil {
		request = NewSetCcnBandwidthRenewFlagRequest()
	}
	response = NewSetCcnBandwidthRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpnConnectionRequest() (request *CreateVpnConnectionRequest) {
	request = &CreateVpnConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnConnection")
	return
}

func NewCreateVpnConnectionResponse() (response *CreateVpnConnectionResponse) {
	response = &CreateVpnConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateVpnConnection）用于创建VPN通道。
func (c *Client) CreateVpnConnection(request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
	if request == nil {
		request = NewCreateVpnConnectionRequest()
	}
	response = NewCreateVpnConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressTemplateGroupAttributeRequest() (request *ModifyAddressTemplateGroupAttributeRequest) {
	request = &ModifyAddressTemplateGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateGroupAttribute")
	return
}

func NewModifyAddressTemplateGroupAttributeResponse() (response *ModifyAddressTemplateGroupAttributeResponse) {
	response = &ModifyAddressTemplateGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAddressTemplateGroupAttribute）用于修改IP地址模板集合
func (c *Client) ModifyAddressTemplateGroupAttribute(request *ModifyAddressTemplateGroupAttributeRequest) (response *ModifyAddressTemplateGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAddressTemplateGroupAttributeRequest()
	}
	response = NewModifyAddressTemplateGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewResetAttachCcnInstancesRequest() (request *ResetAttachCcnInstancesRequest) {
	request = &ResetAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetAttachCcnInstances")
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

func NewCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGateway")
	return
}

func NewCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateNatGateway)用于创建NAT网关。
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	if request == nil {
		request = NewCreateNatGatewayRequest()
	}
	response = NewCreateNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceAllocateIp6AddressesBandwidthRequest() (request *InquiryPriceAllocateIp6AddressesBandwidthRequest) {
	request = &InquiryPriceAllocateIp6AddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceAllocateIp6AddressesBandwidth")
	return
}

func NewInquiryPriceAllocateIp6AddressesBandwidthResponse() (response *InquiryPriceAllocateIp6AddressesBandwidthResponse) {
	response = &InquiryPriceAllocateIp6AddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询IPV6地址分配公网带宽价格
func (c *Client) InquiryPriceAllocateIp6AddressesBandwidth(request *InquiryPriceAllocateIp6AddressesBandwidthRequest) (response *InquiryPriceAllocateIp6AddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceAllocateIp6AddressesBandwidthRequest()
	}
	response = NewInquiryPriceAllocateIp6AddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayFlowMonitorDetailRequest() (request *DescribeGatewayFlowMonitorDetailRequest) {
	request = &DescribeGatewayFlowMonitorDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeGatewayFlowMonitorDetail")
	return
}

func NewDescribeGatewayFlowMonitorDetailResponse() (response *DescribeGatewayFlowMonitorDetailResponse) {
	response = &DescribeGatewayFlowMonitorDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeGatewayFlowMonitorDetail）用于查询网关流量监控明细。
// * 只支持单个网关实例查询。即入参 `VpnId` `DirectConnectGatewayId` `PeeringConnectionId` `NatId` 最多只支持传一个，且必须传一个。
// * 如果网关有流量，但调用本接口没有返回数据，请在控制台对应网关详情页确认是否开启网关流量监控。
func (c *Client) DescribeGatewayFlowMonitorDetail(request *DescribeGatewayFlowMonitorDetailRequest) (response *DescribeGatewayFlowMonitorDetailResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayFlowMonitorDetailRequest()
	}
	response = NewDescribeGatewayFlowMonitorDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAddressTemplateRequest() (request *DeleteAddressTemplateRequest) {
	request = &DeleteAddressTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplate")
	return
}

func NewDeleteAddressTemplateResponse() (response *DeleteAddressTemplateResponse) {
	response = &DeleteAddressTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAddressTemplate）用于删除IP地址模板
func (c *Client) DeleteAddressTemplate(request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteAddressTemplateRequest()
	}
	response = NewDeleteAddressTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
	request = &DescribeAddressQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressQuota")
	return
}

func NewDescribeAddressQuotaResponse() (response *DescribeAddressQuotaResponse) {
	response = &DescribeAddressQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAddressQuota) 用于查询您账户的[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）在当前地域的配额信息。配额详情可参见 [EIP 产品简介](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)。
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAddressQuotaRequest()
	}
	response = NewDescribeAddressQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroup")
	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSecurityGroup）用于创建新的安全组（SecurityGroup）。
// * 每个账户下每个地域的每个项目的<a href="https://cloud.tencent.com/document/product/213/12453">安全组数量限制</a>。
// * 新建的安全组的入站和出站规则默认都是全部拒绝，在创建后通常您需要再调用CreateSecurityGroupPolicies将安全组的规则设置为需要的规则。
// * 创建安全组同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
	if request == nil {
		request = NewCreateSecurityGroupRequest()
	}
	response = NewCreateSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcPeeringConnectionRequest() (request *CreateVpcPeeringConnectionRequest) {
	request = &CreateVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcPeeringConnection")
	return
}

func NewCreateVpcPeeringConnectionResponse() (response *CreateVpcPeeringConnectionResponse) {
	response = &CreateVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建同地域私有网络对等连接
func (c *Client) CreateVpcPeeringConnection(request *CreateVpcPeeringConnectionRequest) (response *CreateVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewCreateVpcPeeringConnectionRequest()
	}
	response = NewCreateVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficMirrorsRequest() (request *DescribeTrafficMirrorsRequest) {
	request = &DescribeTrafficMirrorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTrafficMirrors")
	return
}

func NewDescribeTrafficMirrorsResponse() (response *DescribeTrafficMirrorsResponse) {
	response = &DescribeTrafficMirrorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeTrafficMirrors）用于查询流量镜像实例信息。
func (c *Client) DescribeTrafficMirrors(request *DescribeTrafficMirrorsRequest) (response *DescribeTrafficMirrorsResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficMirrorsRequest()
	}
	response = NewDescribeTrafficMirrorsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCcnRegionBandwidthLimitsRequest() (request *GetCcnRegionBandwidthLimitsRequest) {
	request = &GetCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "GetCcnRegionBandwidthLimits")
	return
}

func NewGetCcnRegionBandwidthLimitsResponse() (response *GetCcnRegionBandwidthLimitsResponse) {
	response = &GetCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetCcnRegionBandwidthLimits）用于查询云联网相关地域带宽信息，该接口只返回已关联网络实例包含的地域
func (c *Client) GetCcnRegionBandwidthLimits(request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewGetCcnRegionBandwidthLimitsRequest()
	}
	response = NewGetCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
	request = &DeleteNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkInterface")
	return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
	response = &DeleteNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNetworkInterface）用于删除弹性网卡。
// * 弹性网卡上绑定了云服务器时，不能被删除。
// * 删除指定弹性网卡，弹性网卡必须先和子机解绑才能删除。删除之后弹性网卡上所有内网IP都将被退还。
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkInterfaceRequest()
	}
	response = NewDeleteNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectGatewayCcnRoutesRequest() (request *DescribeDirectConnectGatewayCcnRoutesRequest) {
	request = &DescribeDirectConnectGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGatewayCcnRoutes")
	return
}

func NewDescribeDirectConnectGatewayCcnRoutesResponse() (response *DescribeDirectConnectGatewayCcnRoutesResponse) {
	response = &DescribeDirectConnectGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDirectConnectGatewayCcnRoutes）用于查询专线网关的云联网路由（IDC网段）
func (c *Client) DescribeDirectConnectGatewayCcnRoutes(request *DescribeDirectConnectGatewayCcnRoutesRequest) (response *DescribeDirectConnectGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectGatewayCcnRoutesRequest()
	}
	response = NewDescribeDirectConnectGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPeerIpTranslationNatRuleRequest() (request *ModifyPeerIpTranslationNatRuleRequest) {
	request = &ModifyPeerIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyPeerIpTranslationNatRule")
	return
}

func NewModifyPeerIpTranslationNatRuleResponse() (response *ModifyPeerIpTranslationNatRuleResponse) {
	response = &ModifyPeerIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关对端IP转换
func (c *Client) ModifyPeerIpTranslationNatRule(request *ModifyPeerIpTranslationNatRuleRequest) (response *ModifyPeerIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyPeerIpTranslationNatRuleRequest()
	}
	response = NewModifyPeerIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
	request = &DeleteSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSubnet")
	return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
	response = &DeleteSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSubnet）用于用于删除子网(Subnet)。
// * 删除子网前，请清理该子网下所有资源，包括云服务器、负载均衡、云数据、noSql、弹性网卡等资源。
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
	if request == nil {
		request = NewDeleteSubnetRequest()
	}
	response = NewDeleteSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLocalSourceIpPortTranslationNatRuleRequest() (request *ModifyLocalSourceIpPortTranslationNatRuleRequest) {
	request = &ModifyLocalSourceIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalSourceIpPortTranslationNatRule")
	return
}

func NewModifyLocalSourceIpPortTranslationNatRuleResponse() (response *ModifyLocalSourceIpPortTranslationNatRuleResponse) {
	response = &ModifyLocalSourceIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关本端源IP端口转换
func (c *Client) ModifyLocalSourceIpPortTranslationNatRule(request *ModifyLocalSourceIpPortTranslationNatRuleRequest) (response *ModifyLocalSourceIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyLocalSourceIpPortTranslationNatRuleRequest()
	}
	response = NewModifyLocalSourceIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIp6RuleRequest() (request *ModifyIp6RuleRequest) {
	request = &ModifyIp6RuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIp6Rule")
	return
}

func NewModifyIp6RuleResponse() (response *ModifyIp6RuleResponse) {
	response = &ModifyIp6RuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改IPV6转换规则，当前仅支持修改转换规则名称，IPV4地址和IPV4端口号
func (c *Client) ModifyIp6Rule(request *ModifyIp6RuleRequest) (response *ModifyIp6RuleResponse, err error) {
	if request == nil {
		request = NewModifyIp6RuleRequest()
	}
	response = NewModifyIp6RuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceTemplateGroupRequest() (request *CreateServiceTemplateGroupRequest) {
	request = &CreateServiceTemplateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplateGroup")
	return
}

func NewCreateServiceTemplateGroupResponse() (response *CreateServiceTemplateGroupResponse) {
	response = &CreateServiceTemplateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateServiceTemplateGroup）用于创建协议端口模板集合
func (c *Client) CreateServiceTemplateGroup(request *CreateServiceTemplateGroupRequest) (response *CreateServiceTemplateGroupResponse, err error) {
	if request == nil {
		request = NewCreateServiceTemplateGroupRequest()
	}
	response = NewCreateServiceTemplateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressesAttributeRequest() (request *ModifyAddressesAttributeRequest) {
	request = &ModifyAddressesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressesAttribute")
	return
}

func NewModifyAddressesAttributeResponse() (response *ModifyAddressesAttributeResponse) {
	response = &ModifyAddressesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyAddressAttribute) 用于批量修改[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）的属性。
func (c *Client) ModifyAddressesAttribute(request *ModifyAddressesAttributeRequest) (response *ModifyAddressesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAddressesAttributeRequest()
	}
	response = NewModifyAddressesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecurityGroupPoliciesRequest() (request *ModifySecurityGroupPoliciesRequest) {
	request = &ModifySecurityGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupPolicies")
	return
}

func NewModifySecurityGroupPoliciesResponse() (response *ModifySecurityGroupPoliciesResponse) {
	response = &ModifySecurityGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySecurityGroupPolicies）用于重置安全组出站和入站规则（SecurityGroupPolicy）。
//
// * 接口是先删除当前所有的出入站规则，然后再添加 Egress 和 Ingress 规则，不支持自定义索引 PolicyIndex 。
// * 如果指定 SecurityGroupPolicySet.Version 为0, 表示清空所有规则，并忽略Egress和Ingress。
// * Protocol字段支持输入TCP, UDP, ICMP, ICMPV6, GRE, ALL。
// * CidrBlock字段允许输入符合cidr格式标准的任意字符串。(展开)在基础网络中，如果CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * Ipv6CidrBlock字段允许输入符合IPv6 cidr格式标准的任意字符串。(展开)在基础网络中，如果Ipv6CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IPv6，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * SecurityGroupId字段允许输入与待修改的安全组位于相同项目中的安全组ID，包括这个安全组ID本身，代表安全组下所有云服务器的内网IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个ID所关联的云服务器变化而变化，不需要重新修改。
// * Port字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当Protocol字段是TCP或UDP时，Port字段才被接受。
// * Action字段只允许输入ACCEPT或DROP。
// * CidrBlock, Ipv6CidrBlock, SecurityGroupId, AddressTemplate四者是排他关系，不允许同时输入，Protocol + Port和ServiceTemplate二者是排他关系，不允许同时输入。
func (c *Client) ModifySecurityGroupPolicies(request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewModifySecurityGroupPoliciesRequest()
	}
	response = NewModifySecurityGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceTemplateRequest() (request *CreateServiceTemplateRequest) {
	request = &CreateServiceTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplate")
	return
}

func NewCreateServiceTemplateResponse() (response *CreateServiceTemplateResponse) {
	response = &CreateServiceTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateServiceTemplate）用于创建协议端口模板
func (c *Client) CreateServiceTemplate(request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
	if request == nil {
		request = NewCreateServiceTemplateRequest()
	}
	response = NewCreateServiceTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcPrivateIpAddressesRequest() (request *DescribeVpcPrivateIpAddressesRequest) {
	request = &DescribeVpcPrivateIpAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcPrivateIpAddresses")
	return
}

func NewDescribeVpcPrivateIpAddressesResponse() (response *DescribeVpcPrivateIpAddressesResponse) {
	response = &DescribeVpcPrivateIpAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcPrivateIpAddresses）用于查询VPC内网IP信息。<br />
// 只能查询已使用的IP信息，当查询未使用的IP时，本接口不会报错，但不会出现在返回结果里。
func (c *Client) DescribeVpcPrivateIpAddresses(request *DescribeVpcPrivateIpAddressesRequest) (response *DescribeVpcPrivateIpAddressesResponse, err error) {
	if request == nil {
		request = NewDescribeVpcPrivateIpAddressesRequest()
	}
	response = NewDescribeVpcPrivateIpAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBandwidthPackageAttributeRequest() (request *ModifyBandwidthPackageAttributeRequest) {
	request = &ModifyBandwidthPackageAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyBandwidthPackageAttribute")
	return
}

func NewModifyBandwidthPackageAttributeResponse() (response *ModifyBandwidthPackageAttributeResponse) {
	response = &ModifyBandwidthPackageAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于修改带宽包属性，包括带宽包名字等
func (c *Client) ModifyBandwidthPackageAttribute(request *ModifyBandwidthPackageAttributeRequest) (response *ModifyBandwidthPackageAttributeResponse, err error) {
	if request == nil {
		request = NewModifyBandwidthPackageAttributeRequest()
	}
	response = NewModifyBandwidthPackageAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceTemplateGroupAttributeRequest() (request *ModifyServiceTemplateGroupAttributeRequest) {
	request = &ModifyServiceTemplateGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateGroupAttribute")
	return
}

func NewModifyServiceTemplateGroupAttributeResponse() (response *ModifyServiceTemplateGroupAttributeResponse) {
	response = &ModifyServiceTemplateGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyServiceTemplateGroupAttribute）用于修改协议端口模板集合。
func (c *Client) ModifyServiceTemplateGroupAttribute(request *ModifyServiceTemplateGroupAttributeRequest) (response *ModifyServiceTemplateGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyServiceTemplateGroupAttributeRequest()
	}
	response = NewModifyServiceTemplateGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceRoutesRequest() (request *ReplaceRoutesRequest) {
	request = &ReplaceRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRoutes")
	return
}

func NewReplaceRoutesResponse() (response *ReplaceRoutesResponse) {
	response = &ReplaceRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceRoutes）根据路由策略ID（RouteId）修改指定的路由策略（Route），支持批量修改。
func (c *Client) ReplaceRoutes(request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
	if request == nil {
		request = NewReplaceRoutesRequest()
	}
	response = NewReplaceRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTemplateGroupInstancesRequest() (request *DescribeServiceTemplateGroupInstancesRequest) {
	request = &DescribeServiceTemplateGroupInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplateGroupInstances")
	return
}

func NewDescribeServiceTemplateGroupInstancesResponse() (response *DescribeServiceTemplateGroupInstancesResponse) {
	response = &DescribeServiceTemplateGroupInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceTemplateGroupInstances）用于查询参数模板协议端口组关联的实例列表。本接口不会返回查询的结果，需要根据返回的RequestId调用DescribeVpcTaskResult接口获取结果。
func (c *Client) DescribeServiceTemplateGroupInstances(request *DescribeServiceTemplateGroupInstancesRequest) (response *DescribeServiceTemplateGroupInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceTemplateGroupInstancesRequest()
	}
	response = NewDescribeServiceTemplateGroupInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkInterfaceExtendIpRequest() (request *ModifyNetworkInterfaceExtendIpRequest) {
	request = &ModifyNetworkInterfaceExtendIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkInterfaceExtendIp")
	return
}

func NewModifyNetworkInterfaceExtendIpResponse() (response *ModifyNetworkInterfaceExtendIpResponse) {
	response = &ModifyNetworkInterfaceExtendIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改弹性网卡扩展ip
func (c *Client) ModifyNetworkInterfaceExtendIp(request *ModifyNetworkInterfaceExtendIpRequest) (response *ModifyNetworkInterfaceExtendIpResponse, err error) {
	if request == nil {
		request = NewModifyNetworkInterfaceExtendIpRequest()
	}
	response = NewModifyNetworkInterfaceExtendIpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoutesRequest() (request *CreateRoutesRequest) {
	request = &CreateRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutes")
	return
}

func NewCreateRoutesResponse() (response *CreateRoutesResponse) {
	response = &CreateRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateRoutes)用于创建路由策略。
// * 向指定路由表批量新增路由策略。
func (c *Client) CreateRoutes(request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
	if request == nil {
		request = NewCreateRoutesRequest()
	}
	response = NewCreateRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectGatewaysRequest() (request *DescribeDirectConnectGatewaysRequest) {
	request = &DescribeDirectConnectGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGateways")
	return
}

func NewDescribeDirectConnectGatewaysResponse() (response *DescribeDirectConnectGatewaysResponse) {
	response = &DescribeDirectConnectGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDirectConnectGateways）用于查询专线网关。
func (c *Client) DescribeDirectConnectGateways(request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectGatewaysRequest()
	}
	response = NewDescribeDirectConnectGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewResetVpnConnectionRequest() (request *ResetVpnConnectionRequest) {
	request = &ResetVpnConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetVpnConnection")
	return
}

func NewResetVpnConnectionResponse() (response *ResetVpnConnectionResponse) {
	response = &ResetVpnConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ResetVpnConnection)用于重置VPN通道。
func (c *Client) ResetVpnConnection(request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
	if request == nil {
		request = NewResetVpnConnectionRequest()
	}
	response = NewResetVpnConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowLogRequest() (request *DescribeFlowLogRequest) {
	request = &DescribeFlowLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeFlowLog")
	return
}

func NewDescribeFlowLogResponse() (response *DescribeFlowLogResponse) {
	response = &DescribeFlowLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeFlowLog）用于查询流日志实例信息
func (c *Client) DescribeFlowLog(request *DescribeFlowLogRequest) (response *DescribeFlowLogResponse, err error) {
	if request == nil {
		request = NewDescribeFlowLogRequest()
	}
	response = NewDescribeFlowLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIp6AddressesRequest() (request *DescribeIp6AddressesRequest) {
	request = &DescribeIp6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIp6Addresses")
	return
}

func NewDescribeIp6AddressesResponse() (response *DescribeIp6AddressesResponse) {
	response = &DescribeIp6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询IPV6地址信息
func (c *Client) DescribeIp6Addresses(request *DescribeIp6AddressesRequest) (response *DescribeIp6AddressesResponse, err error) {
	if request == nil {
		request = NewDescribeIp6AddressesRequest()
	}
	response = NewDescribeIp6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseIp6AddressesBandwidthRequest() (request *ReleaseIp6AddressesBandwidthRequest) {
	request = &ReleaseIp6AddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReleaseIp6AddressesBandwidth")
	return
}

func NewReleaseIp6AddressesBandwidthResponse() (response *ReleaseIp6AddressesBandwidthResponse) {
	response = &ReleaseIp6AddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于给弹性公网IPv6地址释放带宽。
func (c *Client) ReleaseIp6AddressesBandwidth(request *ReleaseIp6AddressesBandwidthRequest) (response *ReleaseIp6AddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewReleaseIp6AddressesBandwidthRequest()
	}
	response = NewReleaseIp6AddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLocalIpTranslationNatRuleRequest() (request *ModifyLocalIpTranslationNatRuleRequest) {
	request = &ModifyLocalIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalIpTranslationNatRule")
	return
}

func NewModifyLocalIpTranslationNatRuleResponse() (response *ModifyLocalIpTranslationNatRuleResponse) {
	response = &ModifyLocalIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关本端IP转换
func (c *Client) ModifyLocalIpTranslationNatRule(request *ModifyLocalIpTranslationNatRuleRequest) (response *ModifyLocalIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyLocalIpTranslationNatRuleRequest()
	}
	response = NewModifyLocalIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkAclEntriesRequest() (request *ModifyNetworkAclEntriesRequest) {
	request = &ModifyNetworkAclEntriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAclEntries")
	return
}

func NewModifyNetworkAclEntriesResponse() (response *ModifyNetworkAclEntriesResponse) {
	response = &ModifyNetworkAclEntriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNetworkAclEntries）用于修改（包括添加和删除）网络ACL的入站规则和出站规则。
func (c *Client) ModifyNetworkAclEntries(request *ModifyNetworkAclEntriesRequest) (response *ModifyNetworkAclEntriesResponse, err error) {
	if request == nil {
		request = NewModifyNetworkAclEntriesRequest()
	}
	response = NewModifyNetworkAclEntriesResponse()
	err = c.Send(request, response)
	return
}

func NewResetTrafficMirrorSrcsRequest() (request *ResetTrafficMirrorSrcsRequest) {
	request = &ResetTrafficMirrorSrcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetTrafficMirrorSrcs")
	return
}

func NewResetTrafficMirrorSrcsResponse() (response *ResetTrafficMirrorSrcsResponse) {
	response = &ResetTrafficMirrorSrcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetTrafficMirrorSrcs）用于重置流量镜像实例采集对象。
func (c *Client) ResetTrafficMirrorSrcs(request *ResetTrafficMirrorSrcsRequest) (response *ResetTrafficMirrorSrcsResponse, err error) {
	if request == nil {
		request = NewResetTrafficMirrorSrcsRequest()
	}
	response = NewResetTrafficMirrorSrcsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDirectConnectGatewayCcnRoutesRequest() (request *CreateDirectConnectGatewayCcnRoutesRequest) {
	request = &CreateDirectConnectGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGatewayCcnRoutes")
	return
}

func NewCreateDirectConnectGatewayCcnRoutesResponse() (response *CreateDirectConnectGatewayCcnRoutesResponse) {
	response = &CreateDirectConnectGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDirectConnectGatewayCcnRoutes）用于创建专线网关的云联网路由（IDC网段）
func (c *Client) CreateDirectConnectGatewayCcnRoutes(request *CreateDirectConnectGatewayCcnRoutesRequest) (response *CreateDirectConnectGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewCreateDirectConnectGatewayCcnRoutesRequest()
	}
	response = NewCreateDirectConnectGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcGatewaysRequest() (request *DescribeVpcGatewaysRequest) {
	request = &DescribeVpcGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcGateways")
	return
}

func NewDescribeVpcGatewaysResponse() (response *DescribeVpcGatewaysResponse) {
	response = &DescribeVpcGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcGateways）用于查询VPC下的网关列表。
func (c *Client) DescribeVpcGateways(request *DescribeVpcGatewaysRequest) (response *DescribeVpcGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeVpcGatewaysRequest()
	}
	response = NewDescribeVpcGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDisableVpnGatewaySSLClientCertRequest() (request *DisableVpnGatewaySSLClientCertRequest) {
	request = &DisableVpnGatewaySSLClientCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisableVpnGatewaySSLClientCert")
	return
}

func NewDisableVpnGatewaySSLClientCertResponse() (response *DisableVpnGatewaySSLClientCertResponse) {
	response = &DisableVpnGatewaySSLClientCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DisableVpnGatewaySSLClientCert 1
func (c *Client) DisableVpnGatewaySSLClientCert(request *DisableVpnGatewaySSLClientCertRequest) (response *DisableVpnGatewaySSLClientCertResponse, err error) {
	if request == nil {
		request = NewDisableVpnGatewaySSLClientCertRequest()
	}
	response = NewDisableVpnGatewaySSLClientCertResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpv6RoutesRequest() (request *ModifyIpv6RoutesRequest) {
	request = &ModifyIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIpv6Routes")
	return
}

func NewModifyIpv6RoutesResponse() (response *ModifyIpv6RoutesResponse) {
	response = &ModifyIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IPv6路由规则。
func (c *Client) ModifyIpv6Routes(request *ModifyIpv6RoutesRequest) (response *ModifyIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewModifyIpv6RoutesRequest()
	}
	response = NewModifyIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLocalDestinationIpPortTranslationNatRuleRequest() (request *DeleteLocalDestinationIpPortTranslationNatRuleRequest) {
	request = &DeleteLocalDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalDestinationIpPortTranslationNatRule")
	return
}

func NewDeleteLocalDestinationIpPortTranslationNatRuleResponse() (response *DeleteLocalDestinationIpPortTranslationNatRuleResponse) {
	response = &DeleteLocalDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关本端目的IP端口转换
func (c *Client) DeleteLocalDestinationIpPortTranslationNatRule(request *DeleteLocalDestinationIpPortTranslationNatRuleRequest) (response *DeleteLocalDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLocalDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewDeleteLocalDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTrafficMirrorRequest() (request *DeleteTrafficMirrorRequest) {
	request = &DeleteTrafficMirrorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteTrafficMirror")
	return
}

func NewDeleteTrafficMirrorResponse() (response *DeleteTrafficMirrorResponse) {
	response = &DeleteTrafficMirrorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteTrafficMirror）用于删除流量镜像实例。
func (c *Client) DeleteTrafficMirror(request *DeleteTrafficMirrorRequest) (response *DeleteTrafficMirrorResponse, err error) {
	if request == nil {
		request = NewDeleteTrafficMirrorRequest()
	}
	response = NewDeleteTrafficMirrorResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIp6TranslatorRequest() (request *ModifyIp6TranslatorRequest) {
	request = &ModifyIp6TranslatorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIp6Translator")
	return
}

func NewModifyIp6TranslatorResponse() (response *ModifyIp6TranslatorResponse) {
	response = &ModifyIp6TranslatorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改IP6转换实例属性，当前仅支持修改实例名称。
func (c *Client) ModifyIp6Translator(request *ModifyIp6TranslatorRequest) (response *ModifyIp6TranslatorResponse, err error) {
	if request == nil {
		request = NewModifyIp6TranslatorRequest()
	}
	response = NewModifyIp6TranslatorResponse()
	err = c.Send(request, response)
	return
}

func NewCheckNetDetectStateRequest() (request *CheckNetDetectStateRequest) {
	request = &CheckNetDetectStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckNetDetectState")
	return
}

func NewCheckNetDetectStateResponse() (response *CheckNetDetectStateResponse) {
	response = &CheckNetDetectStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CheckNetDetectState)用于验证网络探测。
func (c *Client) CheckNetDetectState(request *CheckNetDetectStateRequest) (response *CheckNetDetectStateResponse, err error) {
	if request == nil {
		request = NewCheckNetDetectStateRequest()
	}
	response = NewCheckNetDetectStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
	request = &DescribeHaVipsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeHaVips")
	return
}

func NewDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
	response = &DescribeHaVipsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeHaVips）用于查询高可用虚拟IP（HAVIP）列表。
func (c *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
	if request == nil {
		request = NewDescribeHaVipsRequest()
	}
	response = NewDescribeHaVipsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcEndPointServiceWhiteListRequest() (request *ModifyVpcEndPointServiceWhiteListRequest) {
	request = &ModifyVpcEndPointServiceWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointServiceWhiteList")
	return
}

func NewModifyVpcEndPointServiceWhiteListResponse() (response *ModifyVpcEndPointServiceWhiteListResponse) {
	response = &ModifyVpcEndPointServiceWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改终端节点服务白名单。
func (c *Client) ModifyVpcEndPointServiceWhiteList(request *ModifyVpcEndPointServiceWhiteListRequest) (response *ModifyVpcEndPointServiceWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyVpcEndPointServiceWhiteListRequest()
	}
	response = NewModifyVpcEndPointServiceWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewRejectAttachCcnInstancesRequest() (request *RejectAttachCcnInstancesRequest) {
	request = &RejectAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RejectAttachCcnInstances")
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

func NewDescribeSecurityGroupReferencesRequest() (request *DescribeSecurityGroupReferencesRequest) {
	request = &DescribeSecurityGroupReferencesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupReferences")
	return
}

func NewDescribeSecurityGroupReferencesResponse() (response *DescribeSecurityGroupReferencesResponse) {
	response = &DescribeSecurityGroupReferencesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSecurityGroupReferences）用于查询安全组被引用信息。
func (c *Client) DescribeSecurityGroupReferences(request *DescribeSecurityGroupReferencesRequest) (response *DescribeSecurityGroupReferencesResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupReferencesRequest()
	}
	response = NewDescribeSecurityGroupReferencesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIp6TranslatorsRequest() (request *DeleteIp6TranslatorsRequest) {
	request = &DeleteIp6TranslatorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteIp6Translators")
	return
}

func NewDeleteIp6TranslatorsResponse() (response *DeleteIp6TranslatorsResponse) {
	response = &DeleteIp6TranslatorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 1. 该接口用于释放IPV6转换实例，支持批量。
// 2.  如果IPV6转换实例建立有转换规则，会一并删除。
func (c *Client) DeleteIp6Translators(request *DeleteIp6TranslatorsRequest) (response *DeleteIp6TranslatorsResponse, err error) {
	if request == nil {
		request = NewDeleteIp6TranslatorsRequest()
	}
	response = NewDeleteIp6TranslatorsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalSourceIpPortTranslationNatRulesRequest() (request *DescribeLocalSourceIpPortTranslationNatRulesRequest) {
	request = &DescribeLocalSourceIpPortTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalSourceIpPortTranslationNatRules")
	return
}

func NewDescribeLocalSourceIpPortTranslationNatRulesResponse() (response *DescribeLocalSourceIpPortTranslationNatRulesResponse) {
	response = &DescribeLocalSourceIpPortTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关本端源IP端口转换
func (c *Client) DescribeLocalSourceIpPortTranslationNatRules(request *DescribeLocalSourceIpPortTranslationNatRulesRequest) (response *DescribeLocalSourceIpPortTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLocalSourceIpPortTranslationNatRulesRequest()
	}
	response = NewDescribeLocalSourceIpPortTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateTrafficPackagesRequest() (request *InquiryPriceCreateTrafficPackagesRequest) {
	request = &InquiryPriceCreateTrafficPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateTrafficPackages")
	return
}

func NewInquiryPriceCreateTrafficPackagesResponse() (response *InquiryPriceCreateTrafficPackagesResponse) {
	response = &InquiryPriceCreateTrafficPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceCreateTrafficPackages(request *InquiryPriceCreateTrafficPackagesRequest) (response *InquiryPriceCreateTrafficPackagesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateTrafficPackagesRequest()
	}
	response = NewInquiryPriceCreateTrafficPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssistantCidrRequest() (request *ModifyAssistantCidrRequest) {
	request = &ModifyAssistantCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAssistantCidr")
	return
}

func NewModifyAssistantCidrResponse() (response *ModifyAssistantCidrResponse) {
	response = &ModifyAssistantCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyAssistantCidr)用于批量修改辅助CIDR，支持新增和删除。（接口灰度中，如需使用请提工单。）
func (c *Client) ModifyAssistantCidr(request *ModifyAssistantCidrRequest) (response *ModifyAssistantCidrResponse, err error) {
	if request == nil {
		request = NewModifyAssistantCidrRequest()
	}
	response = NewModifyAssistantCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGateways")
	return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNatGateways）用于查询 NAT 网关。
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewaysRequest()
	}
	response = NewDescribeNatGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateIp6TranslatorsRequest() (request *InquiryPriceCreateIp6TranslatorsRequest) {
	request = &InquiryPriceCreateIp6TranslatorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateIp6Translators")
	return
}

func NewInquiryPriceCreateIp6TranslatorsResponse() (response *InquiryPriceCreateIp6TranslatorsResponse) {
	response = &InquiryPriceCreateIp6TranslatorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询创建IPV6转换IPV4实例的的价格
func (c *Client) InquiryPriceCreateIp6Translators(request *InquiryPriceCreateIp6TranslatorsRequest) (response *InquiryPriceCreateIp6TranslatorsResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateIp6TranslatorsRequest()
	}
	response = NewInquiryPriceCreateIp6TranslatorsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceUpdateCcnBandwidthRequest() (request *InquiryPriceUpdateCcnBandwidthRequest) {
	request = &InquiryPriceUpdateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceUpdateCcnBandwidth")
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

func NewInquiryPriceCreateCcnBandwidthRequest() (request *InquiryPriceCreateCcnBandwidthRequest) {
	request = &InquiryPriceCreateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateCcnBandwidth")
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

func NewCreateAddressTemplateGroupRequest() (request *CreateAddressTemplateGroupRequest) {
	request = &CreateAddressTemplateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplateGroup")
	return
}

func NewCreateAddressTemplateGroupResponse() (response *CreateAddressTemplateGroupResponse) {
	response = &CreateAddressTemplateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAddressTemplateGroup）用于创建IP地址模版集合
func (c *Client) CreateAddressTemplateGroup(request *CreateAddressTemplateGroupRequest) (response *CreateAddressTemplateGroupResponse, err error) {
	if request == nil {
		request = NewCreateAddressTemplateGroupRequest()
	}
	response = NewCreateAddressTemplateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTemplateGroupsRequest() (request *DescribeServiceTemplateGroupsRequest) {
	request = &DescribeServiceTemplateGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplateGroups")
	return
}

func NewDescribeServiceTemplateGroupsResponse() (response *DescribeServiceTemplateGroupsResponse) {
	response = &DescribeServiceTemplateGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceTemplateGroups）用于查询协议端口模板集合
func (c *Client) DescribeServiceTemplateGroups(request *DescribeServiceTemplateGroupsRequest) (response *DescribeServiceTemplateGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeServiceTemplateGroupsRequest()
	}
	response = NewDescribeServiceTemplateGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableGatewayFlowMonitorRequest() (request *EnableGatewayFlowMonitorRequest) {
	request = &EnableGatewayFlowMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableGatewayFlowMonitor")
	return
}

func NewEnableGatewayFlowMonitorResponse() (response *EnableGatewayFlowMonitorResponse) {
	response = &EnableGatewayFlowMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（EnableGatewayFlowMonitor）用于开启网关流量监控。
func (c *Client) EnableGatewayFlowMonitor(request *EnableGatewayFlowMonitorRequest) (response *EnableGatewayFlowMonitorResponse, err error) {
	if request == nil {
		request = NewEnableGatewayFlowMonitorRequest()
	}
	response = NewEnableGatewayFlowMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewUnassignIpv6AddressesRequest() (request *UnassignIpv6AddressesRequest) {
	request = &UnassignIpv6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6Addresses")
	return
}

func NewUnassignIpv6AddressesResponse() (response *UnassignIpv6AddressesResponse) {
	response = &UnassignIpv6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnassignIpv6Addresses）用于释放弹性网卡`IPv6`地址。<br />
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`RequestId`轮询`QueryTask`接口。
func (c *Client) UnassignIpv6Addresses(request *UnassignIpv6AddressesRequest) (response *UnassignIpv6AddressesResponse, err error) {
	if request == nil {
		request = NewUnassignIpv6AddressesRequest()
	}
	response = NewUnassignIpv6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewUnassignIpv6CidrBlockRequest() (request *UnassignIpv6CidrBlockRequest) {
	request = &UnassignIpv6CidrBlockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6CidrBlock")
	return
}

func NewUnassignIpv6CidrBlockResponse() (response *UnassignIpv6CidrBlockResponse) {
	response = &UnassignIpv6CidrBlockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnassignIpv6CidrBlock）用于释放IPv6网段。<br />
// 网段如果还有IP占用且未回收，则网段无法释放。
func (c *Client) UnassignIpv6CidrBlock(request *UnassignIpv6CidrBlockRequest) (response *UnassignIpv6CidrBlockResponse, err error) {
	if request == nil {
		request = NewUnassignIpv6CidrBlockRequest()
	}
	response = NewUnassignIpv6CidrBlockResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDetectsRequest() (request *DescribeNetDetectsRequest) {
	request = &DescribeNetDetectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetDetects")
	return
}

func NewDescribeNetDetectsResponse() (response *DescribeNetDetectsResponse) {
	response = &DescribeNetDetectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNetDetects）用于查询网络探测列表。
func (c *Client) DescribeNetDetects(request *DescribeNetDetectsRequest) (response *DescribeNetDetectsResponse, err error) {
	if request == nil {
		request = NewDescribeNetDetectsRequest()
	}
	response = NewDescribeNetDetectsResponse()
	err = c.Send(request, response)
	return
}

func NewUnassignPrivateIpAddressesRequest() (request *UnassignPrivateIpAddressesRequest) {
	request = &UnassignPrivateIpAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnassignPrivateIpAddresses")
	return
}

func NewUnassignPrivateIpAddressesResponse() (response *UnassignPrivateIpAddressesResponse) {
	response = &UnassignPrivateIpAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnassignPrivateIpAddresses）用于弹性网卡退还内网 IP。
// * 退还弹性网卡上的辅助内网IP，接口自动解关联弹性公网 IP。不能退还弹性网卡的主内网IP。
func (c *Client) UnassignPrivateIpAddresses(request *UnassignPrivateIpAddressesRequest) (response *UnassignPrivateIpAddressesResponse, err error) {
	if request == nil {
		request = NewUnassignPrivateIpAddressesRequest()
	}
	response = NewUnassignPrivateIpAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTrafficMirrorAllFilterRequest() (request *UpdateTrafficMirrorAllFilterRequest) {
	request = &UpdateTrafficMirrorAllFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UpdateTrafficMirrorAllFilter")
	return
}

func NewUpdateTrafficMirrorAllFilterResponse() (response *UpdateTrafficMirrorAllFilterResponse) {
	response = &UpdateTrafficMirrorAllFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateTrafficMirrorAllFilter）用于更新流量镜像实例的过滤规则或者采集对象。
func (c *Client) UpdateTrafficMirrorAllFilter(request *UpdateTrafficMirrorAllFilterRequest) (response *UpdateTrafficMirrorAllFilterResponse, err error) {
	if request == nil {
		request = NewUpdateTrafficMirrorAllFilterRequest()
	}
	response = NewUpdateTrafficMirrorAllFilterResponse()
	err = c.Send(request, response)
	return
}

func NewBandwidthLimitForCcnAlarmOnlyRequest() (request *BandwidthLimitForCcnAlarmOnlyRequest) {
	request = &BandwidthLimitForCcnAlarmOnlyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "BandwidthLimitForCcnAlarmOnly")
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

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroup")
	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSecurityGroup）用于删除安全组（SecurityGroup）。
// * 只有当前账号下的安全组允许被删除。
// * 安全组实例ID如果在其他安全组的规则中被引用，则无法直接删除。这种情况下，需要先进行规则修改，再删除安全组。
// * 删除的安全组无法再找回，请谨慎调用。
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	response = NewDeleteSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSecurityGroupPoliciesRequest() (request *DeleteSecurityGroupPoliciesRequest) {
	request = &DeleteSecurityGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroupPolicies")
	return
}

func NewDeleteSecurityGroupPoliciesResponse() (response *DeleteSecurityGroupPoliciesResponse) {
	response = &DeleteSecurityGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSecurityGroupPolicies）用于用于删除安全组规则（SecurityGroupPolicy）。
// * SecurityGroupPolicySet.Version 用于指定要操作的安全组的版本。传入 Version 版本号若不等于当前安全组的最新版本，将返回失败；若不传 Version 则直接删除指定PolicyIndex的规则。
func (c *Client) DeleteSecurityGroupPolicies(request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupPoliciesRequest()
	}
	response = NewDeleteSecurityGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceTemplateGroupRequest() (request *DeleteServiceTemplateGroupRequest) {
	request = &DeleteServiceTemplateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplateGroup")
	return
}

func NewDeleteServiceTemplateGroupResponse() (response *DeleteServiceTemplateGroupResponse) {
	response = &DeleteServiceTemplateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteServiceTemplateGroup）用于删除协议端口模板集合
func (c *Client) DeleteServiceTemplateGroup(request *DeleteServiceTemplateGroupRequest) (response *DeleteServiceTemplateGroupResponse, err error) {
	if request == nil {
		request = NewDeleteServiceTemplateGroupRequest()
	}
	response = NewDeleteServiceTemplateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceRouteTableAssociationRequest() (request *ReplaceRouteTableAssociationRequest) {
	request = &ReplaceRouteTableAssociationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRouteTableAssociation")
	return
}

func NewReplaceRouteTableAssociationResponse() (response *ReplaceRouteTableAssociationResponse) {
	response = &ReplaceRouteTableAssociationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceRouteTableAssociation)用于修改子网（Subnet）关联的路由表（RouteTable）。
// * 一个子网只能关联一个路由表。
func (c *Client) ReplaceRouteTableAssociation(request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
	if request == nil {
		request = NewReplaceRouteTableAssociationRequest()
	}
	response = NewReplaceRouteTableAssociationResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateNatGatewayAddressRequest() (request *AssociateNatGatewayAddressRequest) {
	request = &AssociateNatGatewayAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateNatGatewayAddress")
	return
}

func NewAssociateNatGatewayAddressResponse() (response *AssociateNatGatewayAddressResponse) {
	response = &AssociateNatGatewayAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(AssociateNatGatewayAddress)用于NAT网关绑定弹性IP（EIP）。
func (c *Client) AssociateNatGatewayAddress(request *AssociateNatGatewayAddressRequest) (response *AssociateNatGatewayAddressResponse, err error) {
	if request == nil {
		request = NewAssociateNatGatewayAddressRequest()
	}
	response = NewAssociateNatGatewayAddressResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
	request = &ModifyNetworkInterfaceAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkInterfaceAttribute")
	return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
	response = &ModifyNetworkInterfaceAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNetworkInterfaceAttribute）用于修改弹性网卡属性。
func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
	if request == nil {
		request = NewModifyNetworkInterfaceAttributeRequest()
	}
	response = NewModifyNetworkInterfaceAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewResetVpnGatewayInternetMaxBandwidthRequest() (request *ResetVpnGatewayInternetMaxBandwidthRequest) {
	request = &ResetVpnGatewayInternetMaxBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetVpnGatewayInternetMaxBandwidth")
	return
}

func NewResetVpnGatewayInternetMaxBandwidthResponse() (response *ResetVpnGatewayInternetMaxBandwidthResponse) {
	response = &ResetVpnGatewayInternetMaxBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetVpnGatewayInternetMaxBandwidth）调整VPN网关带宽上限。目前支持升级配置，如果是包年包月VPN网关需要在有效期内。
func (c *Client) ResetVpnGatewayInternetMaxBandwidth(request *ResetVpnGatewayInternetMaxBandwidthRequest) (response *ResetVpnGatewayInternetMaxBandwidthResponse, err error) {
	if request == nil {
		request = NewResetVpnGatewayInternetMaxBandwidthRequest()
	}
	response = NewResetVpnGatewayInternetMaxBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcExtendCidrRequest() (request *ModifyVpcExtendCidrRequest) {
	request = &ModifyVpcExtendCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcExtendCidr")
	return
}

func NewModifyVpcExtendCidrResponse() (response *ModifyVpcExtendCidrResponse) {
	response = &ModifyVpcExtendCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPC扩展CIDR属性
func (c *Client) ModifyVpcExtendCidr(request *ModifyVpcExtendCidrRequest) (response *ModifyVpcExtendCidrResponse, err error) {
	if request == nil {
		request = NewModifyVpcExtendCidrRequest()
	}
	response = NewModifyVpcExtendCidrResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePeerIpTranslationNatRuleRequest() (request *CreatePeerIpTranslationNatRuleRequest) {
	request = &CreatePeerIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreatePeerIpTranslationNatRule")
	return
}

func NewCreatePeerIpTranslationNatRuleResponse() (response *CreatePeerIpTranslationNatRuleResponse) {
	response = &CreatePeerIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关对端IP转换
func (c *Client) CreatePeerIpTranslationNatRule(request *CreatePeerIpTranslationNatRuleRequest) (response *CreatePeerIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreatePeerIpTranslationNatRuleRequest()
	}
	response = NewCreatePeerIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomerGatewayRequest() (request *CreateCustomerGatewayRequest) {
	request = &CreateCustomerGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateCustomerGateway")
	return
}

func NewCreateCustomerGatewayResponse() (response *CreateCustomerGatewayResponse) {
	response = &CreateCustomerGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCustomerGateway）用于创建对端网关。
func (c *Client) CreateCustomerGateway(request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
	if request == nil {
		request = NewCreateCustomerGatewayRequest()
	}
	response = NewCreateCustomerGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcEndPointAndServiceAllRequest() (request *DeleteVpcEndPointAndServiceAllRequest) {
	request = &DeleteVpcEndPointAndServiceAllRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPointAndServiceAll")
	return
}

func NewDeleteVpcEndPointAndServiceAllResponse() (response *DeleteVpcEndPointAndServiceAllResponse) {
	response = &DeleteVpcEndPointAndServiceAllResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当终端节点服务绑定的后端服务欠费时，需要删除关联的终端节点和终端节点服务。
func (c *Client) DeleteVpcEndPointAndServiceAll(request *DeleteVpcEndPointAndServiceAllRequest) (response *DeleteVpcEndPointAndServiceAllResponse, err error) {
	if request == nil {
		request = NewDeleteVpcEndPointAndServiceAllRequest()
	}
	response = NewDeleteVpcEndPointAndServiceAllResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateAddressesRequest() (request *MigrateAddressesRequest) {
	request = &MigrateAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "MigrateAddresses")
	return
}

func NewMigrateAddressesResponse() (response *MigrateAddressesResponse) {
	response = &MigrateAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (MigrateAddresses) 用于跨账号迁移一个或多个弹性公网IP（简称 EIP），且仅限于未绑定任何资源的后付费EIP
func (c *Client) MigrateAddresses(request *MigrateAddressesRequest) (response *MigrateAddressesResponse, err error) {
	if request == nil {
		request = NewMigrateAddressesRequest()
	}
	response = NewMigrateAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceTemplateAttributeRequest() (request *ModifyServiceTemplateAttributeRequest) {
	request = &ModifyServiceTemplateAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateAttribute")
	return
}

func NewModifyServiceTemplateAttributeResponse() (response *ModifyServiceTemplateAttributeResponse) {
	response = &ModifyServiceTemplateAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyServiceTemplateAttribute）用于修改协议端口模板
func (c *Client) ModifyServiceTemplateAttribute(request *ModifyServiceTemplateAttributeRequest) (response *ModifyServiceTemplateAttributeResponse, err error) {
	if request == nil {
		request = NewModifyServiceTemplateAttributeRequest()
	}
	response = NewModifyServiceTemplateAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDirectConnectGatewayRequest() (request *DeleteDirectConnectGatewayRequest) {
	request = &DeleteDirectConnectGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGateway")
	return
}

func NewDeleteDirectConnectGatewayResponse() (response *DeleteDirectConnectGatewayResponse) {
	response = &DeleteDirectConnectGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteDirectConnectGateway）用于删除专线网关。
// <li>如果是 NAT 网关，删除专线网关后，NAT 规则以及 ACL 策略都被清理了。</li>
// <li>删除专线网关后，系统会删除路由表中跟该专线网关相关的路由策略。</li>
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`TaskId `轮询`DescribeVpcTaskResult `接口
func (c *Client) DeleteDirectConnectGateway(request *DeleteDirectConnectGatewayRequest) (response *DeleteDirectConnectGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteDirectConnectGatewayRequest()
	}
	response = NewDeleteDirectConnectGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcGlobalExtendCidrsRequest() (request *DescribeVpcGlobalExtendCidrsRequest) {
	request = &DescribeVpcGlobalExtendCidrsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcGlobalExtendCidrs")
	return
}

func NewDescribeVpcGlobalExtendCidrsResponse() (response *DescribeVpcGlobalExtendCidrsResponse) {
	response = &DescribeVpcGlobalExtendCidrsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPC全局扩展CIDR列表
func (c *Client) DescribeVpcGlobalExtendCidrs(request *DescribeVpcGlobalExtendCidrsRequest) (response *DescribeVpcGlobalExtendCidrsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcGlobalExtendCidrsRequest()
	}
	response = NewDescribeVpcGlobalExtendCidrsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBandwidthPackageRequest() (request *CreateBandwidthPackageRequest) {
	request = &CreateBandwidthPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateBandwidthPackage")
	return
}

func NewCreateBandwidthPackageResponse() (response *CreateBandwidthPackageResponse) {
	response = &CreateBandwidthPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口支持创建设备带宽包和IP带宽包
func (c *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
	if request == nil {
		request = NewCreateBandwidthPackageRequest()
	}
	response = NewCreateBandwidthPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHaVipRequest() (request *DeleteHaVipRequest) {
	request = &DeleteHaVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteHaVip")
	return
}

func NewDeleteHaVipResponse() (response *DeleteHaVipResponse) {
	response = &DeleteHaVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteHaVip）用于删除高可用虚拟IP（HAVIP）<br />
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`RequestId`轮询`QueryTask`接口
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
	if request == nil {
		request = NewDeleteHaVipRequest()
	}
	response = NewDeleteHaVipResponse()
	err = c.Send(request, response)
	return
}

func NewTransformAddressRequest() (request *TransformAddressRequest) {
	request = &TransformAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "TransformAddress")
	return
}

func NewTransformAddressResponse() (response *TransformAddressResponse) {
	response = &TransformAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (TransformAddress) 用于将实例的普通公网 IP 转换为[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * 平台对用户每地域每日解绑 EIP 重新分配普通公网 IP 次数有所限制（可参见[EIP 产品简介](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)）。上述配额可通过 [DescribeAddressQuota](../弹性公网IP相关接口/DescribeAddressQuota) 接口获取。
func (c *Client) TransformAddress(request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
	if request == nil {
		request = NewTransformAddressRequest()
	}
	response = NewTransformAddressResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateIp6AddressesBandwidthRequest() (request *AllocateIp6AddressesBandwidthRequest) {
	request = &AllocateIp6AddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AllocateIp6AddressesBandwidth")
	return
}

func NewAllocateIp6AddressesBandwidthResponse() (response *AllocateIp6AddressesBandwidthResponse) {
	response = &AllocateIp6AddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于给IPv6地址初次分配公网带宽
func (c *Client) AllocateIp6AddressesBandwidth(request *AllocateIp6AddressesBandwidthRequest) (response *AllocateIp6AddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewAllocateIp6AddressesBandwidthRequest()
	}
	response = NewAllocateIp6AddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewAssignIpv6CidrBlockRequest() (request *AssignIpv6CidrBlockRequest) {
	request = &AssignIpv6CidrBlockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6CidrBlock")
	return
}

func NewAssignIpv6CidrBlockResponse() (response *AssignIpv6CidrBlockResponse) {
	response = &AssignIpv6CidrBlockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssignIpv6CidrBlock）用于分配IPv6网段。
// * 使用本接口前，您需要已有VPC实例，如果没有可通过接口<a href="https://cloud.tencent.com/document/api/215/15774" title="CreateVpc" target="_blank">CreateVpc</a>创建。
// * 每个VPC只能申请一个IPv6网段
func (c *Client) AssignIpv6CidrBlock(request *AssignIpv6CidrBlockRequest) (response *AssignIpv6CidrBlockResponse, err error) {
	if request == nil {
		request = NewAssignIpv6CidrBlockRequest()
	}
	response = NewAssignIpv6CidrBlockResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnAttributeRequest() (request *ModifyCcnAttributeRequest) {
	request = &ModifyCcnAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyCcnAttribute")
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

func NewCreateIpv6RoutesRequest() (request *CreateIpv6RoutesRequest) {
	request = &CreateIpv6RoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateIpv6Routes")
	return
}

func NewCreateIpv6RoutesResponse() (response *CreateIpv6RoutesResponse) {
	response = &CreateIpv6RoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户创建IPv6路由。
func (c *Client) CreateIpv6Routes(request *CreateIpv6RoutesRequest) (response *CreateIpv6RoutesResponse, err error) {
	if request == nil {
		request = NewCreateIpv6RoutesRequest()
	}
	response = NewCreateIpv6RoutesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckDefaultSubnetRequest() (request *CheckDefaultSubnetRequest) {
	request = &CheckDefaultSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckDefaultSubnet")
	return
}

func NewCheckDefaultSubnetResponse() (response *CheckDefaultSubnetResponse) {
	response = &CheckDefaultSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckDefaultSubnet）用于预判是否可建默认子网。
func (c *Client) CheckDefaultSubnet(request *CheckDefaultSubnetRequest) (response *CheckDefaultSubnetResponse, err error) {
	if request == nil {
		request = NewCheckDefaultSubnetRequest()
	}
	response = NewCheckDefaultSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIp6TranslatorsRequest() (request *DescribeIp6TranslatorsRequest) {
	request = &DescribeIp6TranslatorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIp6Translators")
	return
}

func NewDescribeIp6TranslatorsResponse() (response *DescribeIp6TranslatorsResponse) {
	response = &DescribeIp6TranslatorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 1. 该接口用于查询账户下的IPV6转换实例及其绑定的转换规则信息
// 2. 支持过滤查询
func (c *Client) DescribeIp6Translators(request *DescribeIp6TranslatorsRequest) (response *DescribeIp6TranslatorsResponse, err error) {
	if request == nil {
		request = NewDescribeIp6TranslatorsRequest()
	}
	response = NewDescribeIp6TranslatorsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressInternetChargeTypeRequest() (request *ModifyAddressInternetChargeTypeRequest) {
	request = &ModifyAddressInternetChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressInternetChargeType")
	return
}

func NewModifyAddressInternetChargeTypeResponse() (response *ModifyAddressInternetChargeTypeResponse) {
	response = &ModifyAddressInternetChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于调整具有带宽属性弹性公网IP的网络计费模式
// * 支持BANDWIDTH_PREPAID_BY_MONTH和TRAFFIC_POSTPAID_BY_HOUR两种网络计费模式之间的切换。
// * 每个弹性公网IP支持调整两次，次数超出则无法调整。
func (c *Client) ModifyAddressInternetChargeType(request *ModifyAddressInternetChargeTypeRequest) (response *ModifyAddressInternetChargeTypeResponse, err error) {
	if request == nil {
		request = NewModifyAddressInternetChargeTypeRequest()
	}
	response = NewModifyAddressInternetChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewResetNatGatewayConnectionRequest() (request *ResetNatGatewayConnectionRequest) {
	request = &ResetNatGatewayConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetNatGatewayConnection")
	return
}

func NewResetNatGatewayConnectionResponse() (response *ResetNatGatewayConnectionResponse) {
	response = &ResetNatGatewayConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetNatGatewayConnection）用来NAT网关并发连接上限。
func (c *Client) ResetNatGatewayConnection(request *ResetNatGatewayConnectionRequest) (response *ResetNatGatewayConnectionResponse, err error) {
	if request == nil {
		request = NewResetNatGatewayConnectionRequest()
	}
	response = NewResetNatGatewayConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateNetworkAclSubnetsRequest() (request *DisassociateNetworkAclSubnetsRequest) {
	request = &DisassociateNetworkAclSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNetworkAclSubnets")
	return
}

func NewDisassociateNetworkAclSubnetsResponse() (response *DisassociateNetworkAclSubnetsResponse) {
	response = &DisassociateNetworkAclSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisassociateNetworkAclSubnets）用于网络ACL解关联vpc下的子网。
func (c *Client) DisassociateNetworkAclSubnets(request *DisassociateNetworkAclSubnetsRequest) (response *DisassociateNetworkAclSubnetsResponse, err error) {
	if request == nil {
		request = NewDisassociateNetworkAclSubnetsRequest()
	}
	response = NewDisassociateNetworkAclSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckVpcEndPointServiceExistRequest() (request *CheckVpcEndPointServiceExistRequest) {
	request = &CheckVpcEndPointServiceExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckVpcEndPointServiceExist")
	return
}

func NewCheckVpcEndPointServiceExistResponse() (response *CheckVpcEndPointServiceExistResponse) {
	response = &CheckVpcEndPointServiceExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查终端节点服务是否存在。
func (c *Client) CheckVpcEndPointServiceExist(request *CheckVpcEndPointServiceExistRequest) (response *CheckVpcEndPointServiceExistResponse, err error) {
	if request == nil {
		request = NewCheckVpcEndPointServiceExistRequest()
	}
	response = NewCheckVpcEndPointServiceExistResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcEndPointServiceAttributeRequest() (request *ModifyVpcEndPointServiceAttributeRequest) {
	request = &ModifyVpcEndPointServiceAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointServiceAttribute")
	return
}

func NewModifyVpcEndPointServiceAttributeResponse() (response *ModifyVpcEndPointServiceAttributeResponse) {
	response = &ModifyVpcEndPointServiceAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改终端节点服务相关信息。
func (c *Client) ModifyVpcEndPointServiceAttribute(request *ModifyVpcEndPointServiceAttributeRequest) (response *ModifyVpcEndPointServiceAttributeResponse, err error) {
	if request == nil {
		request = NewModifyVpcEndPointServiceAttributeRequest()
	}
	response = NewModifyVpcEndPointServiceAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveIp6RulesRequest() (request *RemoveIp6RulesRequest) {
	request = &RemoveIp6RulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RemoveIp6Rules")
	return
}

func NewRemoveIp6RulesResponse() (response *RemoveIp6RulesResponse) {
	response = &RemoveIp6RulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 1. 该接口用于删除IPV6转换规则
// 2. 支持批量删除同一个转换实例下的多个转换规则
func (c *Client) RemoveIp6Rules(request *RemoveIp6RulesRequest) (response *RemoveIp6RulesResponse, err error) {
	if request == nil {
		request = NewRemoveIp6RulesRequest()
	}
	response = NewRemoveIp6RulesResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustPublicAddressRequest() (request *AdjustPublicAddressRequest) {
	request = &AdjustPublicAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AdjustPublicAddress")
	return
}

func NewAdjustPublicAddressResponse() (response *AdjustPublicAddressResponse) {
	response = &AdjustPublicAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更换公网IP
func (c *Client) AdjustPublicAddress(request *AdjustPublicAddressRequest) (response *AdjustPublicAddressResponse, err error) {
	if request == nil {
		request = NewAdjustPublicAddressRequest()
	}
	response = NewAdjustPublicAddressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecurityGroupPoliciesRequest() (request *CreateSecurityGroupPoliciesRequest) {
	request = &CreateSecurityGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroupPolicies")
	return
}

func NewCreateSecurityGroupPoliciesResponse() (response *CreateSecurityGroupPoliciesResponse) {
	response = &CreateSecurityGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSecurityGroupPolicies）用于创建安全组规则（SecurityGroupPolicy）。
//
// * Version安全组规则版本号，用户每次更新安全规则版本会自动加1，防止您更新的路由规则已过期，不填不考虑冲突。
// * Protocol字段支持输入TCP, UDP, ICMP, ICMPV6, GRE, ALL。
// * CidrBlock字段允许输入符合cidr格式标准的任意字符串。(展开)在基础网络中，如果CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * Ipv6CidrBlock字段允许输入符合IPv6 cidr格式标准的任意字符串。(展开)在基础网络中，如果Ipv6CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IPv6，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * SecurityGroupId字段允许输入与待修改的安全组位于相同项目中的安全组ID，包括这个安全组ID本身，代表安全组下所有云服务器的内网IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个ID所关联的云服务器变化而变化，不需要重新修改。
// * Port字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当Protocol字段是TCP或UDP时，Port字段才被接受，即Protocol字段不是TCP或UDP时，Protocol和Port排他关系，不允许同时输入，否则会接口报错。
// * Action字段只允许输入ACCEPT或DROP。
// * CidrBlock, Ipv6CidrBlock, SecurityGroupId, AddressTemplate四者是排他关系，不允许同时输入，Protocol + Port和ServiceTemplate二者是排他关系，不允许同时输入。
// * 一次请求中只能创建单个方向的规则, 如果需要指定索引（PolicyIndex）参数, 多条规则的索引必须一致。
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewCreateSecurityGroupPoliciesRequest()
	}
	response = NewCreateSecurityGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDetachClassicLinkVpcRequest() (request *DetachClassicLinkVpcRequest) {
	request = &DetachClassicLinkVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DetachClassicLinkVpc")
	return
}

func NewDetachClassicLinkVpcResponse() (response *DetachClassicLinkVpcResponse) {
	response = &DetachClassicLinkVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DetachClassicLinkVpc)用于删除私有网络和基础网络设备互通。
func (c *Client) DetachClassicLinkVpc(request *DetachClassicLinkVpcRequest) (response *DetachClassicLinkVpcResponse, err error) {
	if request == nil {
		request = NewDetachClassicLinkVpcRequest()
	}
	response = NewDetachClassicLinkVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetIdsRequest() (request *DescribeSubnetIdsRequest) {
	request = &DescribeSubnetIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnetIds")
	return
}

func NewDescribeSubnetIdsResponse() (response *DescribeSubnetIdsResponse) {
	response = &DescribeSubnetIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubnetIds）用于查询子网数字ID。
func (c *Client) DescribeSubnetIds(request *DescribeSubnetIdsRequest) (response *DescribeSubnetIdsResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetIdsRequest()
	}
	response = NewDescribeSubnetIdsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateBandwidthPackageRequest() (request *InquiryPriceCreateBandwidthPackageRequest) {
	request = &InquiryPriceCreateBandwidthPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateBandwidthPackage")
	return
}

func NewInquiryPriceCreateBandwidthPackageResponse() (response *InquiryPriceCreateBandwidthPackageResponse) {
	response = &InquiryPriceCreateBandwidthPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceCreateBandwidthPackage(request *InquiryPriceCreateBandwidthPackageRequest) (response *InquiryPriceCreateBandwidthPackageResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateBandwidthPackageRequest()
	}
	response = NewInquiryPriceCreateBandwidthPackageResponse()
	err = c.Send(request, response)
	return
}

func NewCheckIpInSubnetRequest() (request *CheckIpInSubnetRequest) {
	request = &CheckIpInSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckIpInSubnet")
	return
}

func NewCheckIpInSubnetResponse() (response *CheckIpInSubnetResponse) {
	response = &CheckIpInSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 判断Ip在子网内是有可用
func (c *Client) CheckIpInSubnet(request *CheckIpInSubnetRequest) (response *CheckIpInSubnetResponse, err error) {
	if request == nil {
		request = NewCheckIpInSubnetRequest()
	}
	response = NewCheckIpInSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointServiceRequest() (request *DescribeVpcEndPointServiceRequest) {
	request = &DescribeVpcEndPointServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPointService")
	return
}

func NewDescribeVpcEndPointServiceResponse() (response *DescribeVpcEndPointServiceResponse) {
	response = &DescribeVpcEndPointServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询终端节点服务列表。
func (c *Client) DescribeVpcEndPointService(request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointServiceRequest()
	}
	response = NewDescribeVpcEndPointServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcTaskResultRequest() (request *DescribeVpcTaskResultRequest) {
	request = &DescribeVpcTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcTaskResult")
	return
}

func NewDescribeVpcTaskResultResponse() (response *DescribeVpcTaskResultResponse) {
	response = &DescribeVpcTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcTaskResult）用于查询VPC任务执行结果。
func (c *Client) DescribeVpcTaskResult(request *DescribeVpcTaskResultRequest) (response *DescribeVpcTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeVpcTaskResultRequest()
	}
	response = NewDescribeVpcTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpnGatewayAttributeRequest() (request *ModifyVpnGatewayAttributeRequest) {
	request = &ModifyVpnGatewayAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayAttribute")
	return
}

func NewModifyVpnGatewayAttributeResponse() (response *ModifyVpnGatewayAttributeResponse) {
	response = &ModifyVpnGatewayAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyVpnGatewayAttribute）用于修改VPN网关属性。
func (c *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
	if request == nil {
		request = NewModifyVpnGatewayAttributeRequest()
	}
	response = NewModifyVpnGatewayAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewGetCreateCcnBandwidthDealRequest() (request *GetCreateCcnBandwidthDealRequest) {
	request = &GetCreateCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "GetCreateCcnBandwidthDeal")
	return
}

func NewGetCreateCcnBandwidthDealResponse() (response *GetCreateCcnBandwidthDealResponse) {
	response = &GetCreateCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取创建云联网带宽的商品信息
func (c *Client) GetCreateCcnBandwidthDeal(request *GetCreateCcnBandwidthDealRequest) (response *GetCreateCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetCreateCcnBandwidthDealRequest()
	}
	response = NewGetCreateCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteTableRequest() (request *CreateRouteTableRequest) {
	request = &CreateRouteTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateRouteTable")
	return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
	response = &CreateRouteTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateRouteTable)用于创建路由表。
// * 创建了VPC后，系统会创建一个默认路由表，所有新建的子网都会关联到默认路由表。默认情况下您可以直接使用默认路由表来管理您的路由策略。当您的路由策略较多时，您可以调用创建路由表接口创建更多路由表管理您的路由策略。
// * 创建路由表同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
	if request == nil {
		request = NewCreateRouteTableRequest()
	}
	response = NewCreateRouteTableResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpnGatewaySSLClientCertRequest() (request *EnableVpnGatewaySSLClientCertRequest) {
	request = &EnableVpnGatewaySSLClientCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableVpnGatewaySSLClientCert")
	return
}

func NewEnableVpnGatewaySSLClientCertResponse() (response *EnableVpnGatewaySSLClientCertResponse) {
	response = &EnableVpnGatewaySSLClientCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EnableVpnGatewaySSLClientCert
func (c *Client) EnableVpnGatewaySSLClientCert(request *EnableVpnGatewaySSLClientCertRequest) (response *EnableVpnGatewaySSLClientCertResponse, err error) {
	if request == nil {
		request = NewEnableVpnGatewaySSLClientCertRequest()
	}
	response = NewEnableVpnGatewaySSLClientCertResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClassicLinkInstancesRequest() (request *DescribeClassicLinkInstancesRequest) {
	request = &DescribeClassicLinkInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeClassicLinkInstances")
	return
}

func NewDescribeClassicLinkInstancesResponse() (response *DescribeClassicLinkInstancesResponse) {
	response = &DescribeClassicLinkInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeClassicLinkInstances)用于查询私有网络和基础网络设备互通列表。
func (c *Client) DescribeClassicLinkInstances(request *DescribeClassicLinkInstancesRequest) (response *DescribeClassicLinkInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeClassicLinkInstancesRequest()
	}
	response = NewDescribeClassicLinkInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadSpecificTrafficPackageUsedDetailsRequest() (request *DownloadSpecificTrafficPackageUsedDetailsRequest) {
	request = &DownloadSpecificTrafficPackageUsedDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DownloadSpecificTrafficPackageUsedDetails")
	return
}

func NewDownloadSpecificTrafficPackageUsedDetailsResponse() (response *DownloadSpecificTrafficPackageUsedDetailsResponse) {
	response = &DownloadSpecificTrafficPackageUsedDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DownloadSpecificTrafficPackageUsedDetails)用于生成指定流量包的用量明细文件.
func (c *Client) DownloadSpecificTrafficPackageUsedDetails(request *DownloadSpecificTrafficPackageUsedDetailsRequest) (response *DownloadSpecificTrafficPackageUsedDetailsResponse, err error) {
	if request == nil {
		request = NewDownloadSpecificTrafficPackageUsedDetailsRequest()
	}
	response = NewDownloadSpecificTrafficPackageUsedDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewGetDealStatusByNameRequest() (request *GetDealStatusByNameRequest) {
	request = &GetDealStatusByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "GetDealStatusByName")
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

func NewInquiryPricePublicIp6AddressesRequest() (request *InquiryPricePublicIp6AddressesRequest) {
	request = &InquiryPricePublicIp6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPricePublicIp6Addresses")
	return
}

func NewInquiryPricePublicIp6AddressesResponse() (response *InquiryPricePublicIp6AddressesResponse) {
	response = &InquiryPricePublicIp6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询IPV6地址访问internet的价格
func (c *Client) InquiryPricePublicIp6Addresses(request *InquiryPricePublicIp6AddressesRequest) (response *InquiryPricePublicIp6AddressesResponse, err error) {
	if request == nil {
		request = NewInquiryPricePublicIp6AddressesRequest()
	}
	response = NewInquiryPricePublicIp6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatGatewayRequest() (request *DeleteNatGatewayRequest) {
	request = &DeleteNatGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGateway")
	return
}

func NewDeleteNatGatewayResponse() (response *DeleteNatGatewayResponse) {
	response = &DeleteNatGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNatGateway）用于删除NAT网关。
// 删除 NAT 网关后，系统会自动删除路由表中包含此 NAT 网关的路由项，同时也会解绑弹性公网IP（EIP）。
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteNatGatewayRequest()
	}
	response = NewDeleteNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTemplateInstancesRequest() (request *DescribeServiceTemplateInstancesRequest) {
	request = &DescribeServiceTemplateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplateInstances")
	return
}

func NewDescribeServiceTemplateInstancesResponse() (response *DescribeServiceTemplateInstancesResponse) {
	response = &DescribeServiceTemplateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceTemplateInstances）用于查询参数模板协议端口关联的实例列表。本接口不会返回查询的结果，需要根据返回的RequestId调用DescribeVpcTaskResult接口获取结果。
func (c *Client) DescribeServiceTemplateInstances(request *DescribeServiceTemplateInstancesRequest) (response *DescribeServiceTemplateInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceTemplateInstancesRequest()
	}
	response = NewDescribeServiceTemplateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptAttachCcnInstancesRequest() (request *AcceptAttachCcnInstancesRequest) {
	request = &AcceptAttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AcceptAttachCcnInstances")
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

func NewDescribeAddressBandwidthConfigsRequest() (request *DescribeAddressBandwidthConfigsRequest) {
	request = &DescribeAddressBandwidthConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressBandwidthConfigs")
	return
}

func NewDescribeAddressBandwidthConfigsResponse() (response *DescribeAddressBandwidthConfigsResponse) {
	response = &DescribeAddressBandwidthConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeAddressBandwidthConfigs(request *DescribeAddressBandwidthConfigsRequest) (response *DescribeAddressBandwidthConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeAddressBandwidthConfigsRequest()
	}
	response = NewDescribeAddressBandwidthConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewAddIp6RulesRequest() (request *AddIp6RulesRequest) {
	request = &AddIp6RulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AddIp6Rules")
	return
}

func NewAddIp6RulesResponse() (response *AddIp6RulesResponse) {
	response = &AddIp6RulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 1. 该接口用于在转换实例下添加IPV6转换规则。
// 2. 支持在同一个转换实例下批量添加转换规则，一个账户在一个地域最多50个。
// 3. 一个完整的转换规则包括vip6:vport6:protocol:vip:vport，其中vip6:vport6:protocol必须是唯一。
func (c *Client) AddIp6Rules(request *AddIp6RulesRequest) (response *AddIp6RulesResponse, err error) {
	if request == nil {
		request = NewAddIp6RulesRequest()
	}
	response = NewAddIp6RulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkAclsRequest() (request *DescribeNetworkAclsRequest) {
	request = &DescribeNetworkAclsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkAcls")
	return
}

func NewDescribeNetworkAclsResponse() (response *DescribeNetworkAclsResponse) {
	response = &DescribeNetworkAclsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNetworkAcls）用于查询网络ACL列表。
func (c *Client) DescribeNetworkAcls(request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkAclsRequest()
	}
	response = NewDescribeNetworkAclsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) {
	request = &ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewayDestinationIpPortTranslationNatRule")
	return
}

func NewModifyNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) {
	response = &ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNatGatewayDestinationIpPortTranslationNatRule）用于修改NAT网关端口转发规则。
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRule(request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewModifyNatGatewayDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHaVipAttributeRequest() (request *ModifyHaVipAttributeRequest) {
	request = &ModifyHaVipAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyHaVipAttribute")
	return
}

func NewModifyHaVipAttributeResponse() (response *ModifyHaVipAttributeResponse) {
	response = &ModifyHaVipAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyHaVipAttribute）用于修改高可用虚拟IP（HAVIP）属性
func (c *Client) ModifyHaVipAttribute(request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
	if request == nil {
		request = NewModifyHaVipAttributeRequest()
	}
	response = NewModifyHaVipAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
	request = &AttachNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AttachNetworkInterface")
	return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
	response = &AttachNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachNetworkInterface）用于弹性网卡绑定云主机。
// * 一个云主机可以绑定多个弹性网卡，但只能绑定一个主网卡。更多限制信息详见<a href="https://cloud.tencent.com/document/product/576/18527">弹性网卡使用限制</a>。
// * 一个弹性网卡只能同时绑定一个云主机。
// * 只有运行中或者已关机状态的云主机才能绑定弹性网卡，查看云主机状态详见<a href="https://cloud.tencent.com/document/api/213/9452#InstanceStatus">腾讯云主机信息</a>。
// * 弹性网卡绑定的云主机必须是私有网络的，而且云主机所在可用区必须和弹性网卡子网的可用区相同。
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewAttachNetworkInterfaceRequest()
	}
	response = NewAttachNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpnGatewaySSLServerRequest() (request *DeleteVpnGatewaySSLServerRequest) {
	request = &DeleteVpnGatewaySSLServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGatewaySSLServer")
	return
}

func NewDeleteVpnGatewaySSLServerResponse() (response *DeleteVpnGatewaySSLServerResponse) {
	response = &DeleteVpnGatewaySSLServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteVpnGatewaySSLServer
func (c *Client) DeleteVpnGatewaySSLServer(request *DeleteVpnGatewaySSLServerRequest) (response *DeleteVpnGatewaySSLServerResponse, err error) {
	if request == nil {
		request = NewDeleteVpnGatewaySSLServerRequest()
	}
	response = NewDeleteVpnGatewaySSLServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewaySSLClientsRequest() (request *DescribeVpnGatewaySSLClientsRequest) {
	request = &DescribeVpnGatewaySSLClientsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewaySSLClients")
	return
}

func NewDescribeVpnGatewaySSLClientsResponse() (response *DescribeVpnGatewaySSLClientsResponse) {
	response = &DescribeVpnGatewaySSLClientsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeVpnGatewaySSLClients
func (c *Client) DescribeVpnGatewaySSLClients(request *DescribeVpnGatewaySSLClientsRequest) (response *DescribeVpnGatewaySSLClientsResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewaySSLClientsRequest()
	}
	response = NewDescribeVpnGatewaySSLClientsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSameCityRequest() (request *CheckSameCityRequest) {
	request = &CheckSameCityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckSameCity")
	return
}

func NewCheckSameCityResponse() (response *CheckSameCityResponse) {
	response = &CheckSameCityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckSameCity）用于用于检查指定的两个Region是否为同城。
// * 参数取值参考<a href="https://cloud.tencent.com/document/api/215/15758#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8">地域列表</a>。
func (c *Client) CheckSameCity(request *CheckSameCityRequest) (response *CheckSameCityResponse, err error) {
	if request == nil {
		request = NewCheckSameCityRequest()
	}
	response = NewCheckSameCityResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDirectConnectGatewayCcnRoutesRequest() (request *ReplaceDirectConnectGatewayCcnRoutesRequest) {
	request = &ReplaceDirectConnectGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReplaceDirectConnectGatewayCcnRoutes")
	return
}

func NewReplaceDirectConnectGatewayCcnRoutesResponse() (response *ReplaceDirectConnectGatewayCcnRoutesResponse) {
	response = &ReplaceDirectConnectGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceDirectConnectGatewayCcnRoutes）根据路由ID（RouteId）修改指定的路由（Route），支持批量修改。
func (c *Client) ReplaceDirectConnectGatewayCcnRoutes(request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewReplaceDirectConnectGatewayCcnRoutesRequest()
	}
	response = NewReplaceDirectConnectGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDirectConnectGatewayRequest() (request *CreateDirectConnectGatewayRequest) {
	request = &CreateDirectConnectGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGateway")
	return
}

func NewCreateDirectConnectGatewayResponse() (response *CreateDirectConnectGatewayResponse) {
	response = &CreateDirectConnectGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDirectConnectGateway）用于创建专线网关。
func (c *Client) CreateDirectConnectGateway(request *CreateDirectConnectGatewayRequest) (response *CreateDirectConnectGatewayResponse, err error) {
	if request == nil {
		request = NewCreateDirectConnectGatewayRequest()
	}
	response = NewCreateDirectConnectGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
	request = &CreateSubnetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnet")
	return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
	response = &CreateSubnetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateSubnet)用于创建子网。
// * 创建子网前必须创建好 VPC。
// * 子网创建成功后，子网网段不能修改。子网网段必须在VPC网段内，可以和VPC网段相同（VPC有且只有一个子网时），建议子网网段在VPC网段内，预留网段给其他子网使用。
// * 您可以创建的最小网段子网掩码为28（有16个IP地址），最大网段子网掩码为16（65,536个IP地址）。
// * 同一个VPC内，多个子网的网段不能重叠。
// * 子网创建后会自动关联到默认路由表。
// * 创建子网同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
	if request == nil {
		request = NewCreateSubnetRequest()
	}
	response = NewCreateSubnetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrivateIpAddressesAttributeRequest() (request *ModifyPrivateIpAddressesAttributeRequest) {
	request = &ModifyPrivateIpAddressesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyPrivateIpAddressesAttribute")
	return
}

func NewModifyPrivateIpAddressesAttributeResponse() (response *ModifyPrivateIpAddressesAttributeResponse) {
	response = &ModifyPrivateIpAddressesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyPrivateIpAddressesAttribute）用于修改弹性网卡内网IP属性。
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyPrivateIpAddressesAttributeRequest()
	}
	response = NewModifyPrivateIpAddressesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTrafficPackagesRequest() (request *CreateTrafficPackagesRequest) {
	request = &CreateTrafficPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateTrafficPackages")
	return
}

func NewCreateTrafficPackagesResponse() (response *CreateTrafficPackagesResponse) {
	response = &CreateTrafficPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建共享流量包
func (c *Client) CreateTrafficPackages(request *CreateTrafficPackagesRequest) (response *CreateTrafficPackagesResponse, err error) {
	if request == nil {
		request = NewCreateTrafficPackagesRequest()
	}
	response = NewCreateTrafficPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpnGatewayRequest() (request *DeleteVpnGatewayRequest) {
	request = &DeleteVpnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGateway")
	return
}

func NewDeleteVpnGatewayResponse() (response *DeleteVpnGatewayResponse) {
	response = &DeleteVpnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteVpnGateway）用于删除VPN网关。目前只支持删除运行中的按量计费的IPSEC网关实例。
func (c *Client) DeleteVpnGateway(request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteVpnGatewayRequest()
	}
	response = NewDeleteVpnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewCcnBandwidthRequest() (request *InquiryPriceRenewCcnBandwidthRequest) {
	request = &InquiryPriceRenewCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewCcnBandwidth")
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

func NewModifyDirectConnectGatewayAttributeRequest() (request *ModifyDirectConnectGatewayAttributeRequest) {
	request = &ModifyDirectConnectGatewayAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyDirectConnectGatewayAttribute")
	return
}

func NewModifyDirectConnectGatewayAttributeResponse() (response *ModifyDirectConnectGatewayAttributeResponse) {
	response = &ModifyDirectConnectGatewayAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDirectConnectGatewayAttribute）用于修改专线网关属性
func (c *Client) ModifyDirectConnectGatewayAttribute(request *ModifyDirectConnectGatewayAttributeRequest) (response *ModifyDirectConnectGatewayAttributeResponse, err error) {
	if request == nil {
		request = NewModifyDirectConnectGatewayAttributeRequest()
	}
	response = NewModifyDirectConnectGatewayAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIp6AddressesBandwidthRequest() (request *ModifyIp6AddressesBandwidthRequest) {
	request = &ModifyIp6AddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIp6AddressesBandwidth")
	return
}

func NewModifyIp6AddressesBandwidthResponse() (response *ModifyIp6AddressesBandwidthResponse) {
	response = &ModifyIp6AddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改IPV6地址访问internet的带宽
func (c *Client) ModifyIp6AddressesBandwidth(request *ModifyIp6AddressesBandwidthRequest) (response *ModifyIp6AddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewModifyIp6AddressesBandwidthRequest()
	}
	response = NewModifyIp6AddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcEndPointServiceRequest() (request *DeleteVpcEndPointServiceRequest) {
	request = &DeleteVpcEndPointServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPointService")
	return
}

func NewDeleteVpcEndPointServiceResponse() (response *DeleteVpcEndPointServiceResponse) {
	response = &DeleteVpcEndPointServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除终端节点服务。
func (c *Client) DeleteVpcEndPointService(request *DeleteVpcEndPointServiceRequest) (response *DeleteVpcEndPointServiceResponse, err error) {
	if request == nil {
		request = NewDeleteVpcEndPointServiceRequest()
	}
	response = NewDeleteVpcEndPointServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBInstancesWithSecurityGroupRequest() (request *DescribeDBInstancesWithSecurityGroupRequest) {
	request = &DescribeDBInstancesWithSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDBInstancesWithSecurityGroup")
	return
}

func NewDescribeDBInstancesWithSecurityGroupResponse() (response *DescribeDBInstancesWithSecurityGroupResponse) {
	response = &DescribeDBInstancesWithSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDBInstancesWithSecurityGroup）用于安全组获取绑定的数据库实例信息。
func (c *Client) DescribeDBInstancesWithSecurityGroup(request *DescribeDBInstancesWithSecurityGroupRequest) (response *DescribeDBInstancesWithSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeDBInstancesWithSecurityGroupRequest()
	}
	response = NewDescribeDBInstancesWithSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAssistantCidrRequest() (request *CheckAssistantCidrRequest) {
	request = &CheckAssistantCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckAssistantCidr")
	return
}

func NewCheckAssistantCidrResponse() (response *CheckAssistantCidrResponse) {
	response = &CheckAssistantCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CheckAssistantCidr)用于检查辅助CIDR是否与存量路由、对等连接（对端VPC的CIDR）等资源存在冲突。如果存在重叠，则返回重叠的资源。（接口灰度中，如需使用请提工单。）
// * 检测辅助CIDR是否与当前VPC的主CIDR和辅助CIDR存在重叠。
// * 检测辅助CIDR是否与当前VPC的路由的目的端存在重叠。
// * 检测辅助CIDR是否与当前VPC的对等连接，对端VPC下的主CIDR或辅助CIDR存在重叠。
func (c *Client) CheckAssistantCidr(request *CheckAssistantCidrRequest) (response *CheckAssistantCidrResponse, err error) {
	if request == nil {
		request = NewCheckAssistantCidrRequest()
	}
	response = NewCheckAssistantCidrResponse()
	err = c.Send(request, response)
	return
}

func NewReturnNormalAddressesRequest() (request *ReturnNormalAddressesRequest) {
	request = &ReturnNormalAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReturnNormalAddresses")
	return
}

func NewReturnNormalAddressesResponse() (response *ReturnNormalAddressesResponse) {
	response = &ReturnNormalAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) ReturnNormalAddresses(request *ReturnNormalAddressesRequest) (response *ReturnNormalAddressesResponse, err error) {
	if request == nil {
		request = NewReturnNormalAddressesRequest()
	}
	response = NewReturnNormalAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDetachCcnInstancesRequest() (request *DetachCcnInstancesRequest) {
	request = &DetachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DetachCcnInstances")
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

func NewUpdateTrafficMirrorDirectionRequest() (request *UpdateTrafficMirrorDirectionRequest) {
	request = &UpdateTrafficMirrorDirectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UpdateTrafficMirrorDirection")
	return
}

func NewUpdateTrafficMirrorDirectionResponse() (response *UpdateTrafficMirrorDirectionResponse) {
	response = &UpdateTrafficMirrorDirectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateTrafficMirrorDirection）用于更新流量镜像实例的采集方向。
func (c *Client) UpdateTrafficMirrorDirection(request *UpdateTrafficMirrorDirectionRequest) (response *UpdateTrafficMirrorDirectionResponse, err error) {
	if request == nil {
		request = NewUpdateTrafficMirrorDirectionRequest()
	}
	response = NewUpdateTrafficMirrorDirectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressActionQuotaRequest() (request *DescribeAddressActionQuotaRequest) {
	request = &DescribeAddressActionQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressActionQuota")
	return
}

func NewDescribeAddressActionQuotaResponse() (response *DescribeAddressActionQuotaResponse) {
	response = &DescribeAddressActionQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询上移账户EIP操作配额
func (c *Client) DescribeAddressActionQuota(request *DescribeAddressActionQuotaRequest) (response *DescribeAddressActionQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAddressActionQuotaRequest()
	}
	response = NewDescribeAddressActionQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBandwidthPackageQuotaRequest() (request *DescribeBandwidthPackageQuotaRequest) {
	request = &DescribeBandwidthPackageQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackageQuota")
	return
}

func NewDescribeBandwidthPackageQuotaResponse() (response *DescribeBandwidthPackageQuotaResponse) {
	response = &DescribeBandwidthPackageQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询账户在当前地域的带宽包上限数量以及使用数量
func (c *Client) DescribeBandwidthPackageQuota(request *DescribeBandwidthPackageQuotaRequest) (response *DescribeBandwidthPackageQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeBandwidthPackageQuotaRequest()
	}
	response = NewDescribeBandwidthPackageQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnLimitsRequest() (request *DescribeCcnLimitsRequest) {
	request = &DescribeCcnLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnLimits")
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

func NewGetRenewCcnBandwidthDealRequest() (request *GetRenewCcnBandwidthDealRequest) {
	request = &GetRenewCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "GetRenewCcnBandwidthDeal")
	return
}

func NewGetRenewCcnBandwidthDealResponse() (response *GetRenewCcnBandwidthDealResponse) {
	response = &GetRenewCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取续费云联网带宽订单信息。
func (c *Client) GetRenewCcnBandwidthDeal(request *GetRenewCcnBandwidthDealRequest) (response *GetRenewCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetRenewCcnBandwidthDealRequest()
	}
	response = NewGetRenewCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewRenewCcnBandwidthRequest() (request *RenewCcnBandwidthRequest) {
	request = &RenewCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RenewCcnBandwidth")
	return
}

func NewRenewCcnBandwidthResponse() (response *RenewCcnBandwidthResponse) {
	response = &RenewCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCcnBandwidth）用于续费预付费模式下云联网实例的地域间带宽
func (c *Client) RenewCcnBandwidth(request *RenewCcnBandwidthRequest) (response *RenewCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewRenewCcnBandwidthRequest()
	}
	response = NewRenewCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewFlushHaVipRequest() (request *FlushHaVipRequest) {
	request = &FlushHaVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "FlushHaVip")
	return
}

func NewFlushHaVipResponse() (response *FlushHaVipResponse) {
	response = &FlushHaVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（FlushHaVip）用于刷新HAVIP配置。
func (c *Client) FlushHaVip(request *FlushHaVipRequest) (response *FlushHaVipResponse, err error) {
	if request == nil {
		request = NewFlushHaVipRequest()
	}
	response = NewFlushHaVipResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFlowLogRequest() (request *CreateFlowLogRequest) {
	request = &CreateFlowLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateFlowLog")
	return
}

func NewCreateFlowLogResponse() (response *CreateFlowLogResponse) {
	response = &CreateFlowLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateFlowLog）用于创建流日志
func (c *Client) CreateFlowLog(request *CreateFlowLogRequest) (response *CreateFlowLogResponse, err error) {
	if request == nil {
		request = NewCreateFlowLogRequest()
	}
	response = NewCreateFlowLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcEndPointRequest() (request *DeleteVpcEndPointRequest) {
	request = &DeleteVpcEndPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPoint")
	return
}

func NewDeleteVpcEndPointResponse() (response *DeleteVpcEndPointResponse) {
	response = &DeleteVpcEndPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除终端节点。
func (c *Client) DeleteVpcEndPoint(request *DeleteVpcEndPointRequest) (response *DeleteVpcEndPointResponse, err error) {
	if request == nil {
		request = NewDeleteVpcEndPointRequest()
	}
	response = NewDeleteVpcEndPointResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableZoneRequest() (request *DescribeAvailableZoneRequest) {
	request = &DescribeAvailableZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAvailableZone")
	return
}

func NewDescribeAvailableZoneResponse() (response *DescribeAvailableZoneResponse) {
	response = &DescribeAvailableZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取EIP可用区列表信息
func (c *Client) DescribeAvailableZone(request *DescribeAvailableZoneRequest) (response *DescribeAvailableZoneResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableZoneRequest()
	}
	response = NewDescribeAvailableZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalIpTranslationAclRulesRequest() (request *DescribeLocalIpTranslationAclRulesRequest) {
	request = &DescribeLocalIpTranslationAclRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalIpTranslationAclRules")
	return
}

func NewDescribeLocalIpTranslationAclRulesResponse() (response *DescribeLocalIpTranslationAclRulesResponse) {
	response = &DescribeLocalIpTranslationAclRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关本端IP转换ACL规则
func (c *Client) DescribeLocalIpTranslationAclRules(request *DescribeLocalIpTranslationAclRulesRequest) (response *DescribeLocalIpTranslationAclRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLocalIpTranslationAclRulesRequest()
	}
	response = NewDescribeLocalIpTranslationAclRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressTemplatesRequest() (request *DescribeAddressTemplatesRequest) {
	request = &DescribeAddressTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplates")
	return
}

func NewDescribeAddressTemplatesResponse() (response *DescribeAddressTemplatesResponse) {
	response = &DescribeAddressTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAddressTemplates）用于查询IP地址模板
func (c *Client) DescribeAddressTemplates(request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeAddressTemplatesRequest()
	}
	response = NewDescribeAddressTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
	request = &DescribeNetworkInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaces")
	return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
	response = &DescribeNetworkInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNetworkInterfaces）用于查询弹性网卡列表。
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkInterfacesRequest()
	}
	response = NewDescribeNetworkInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewHaVipDisassociateAddressIpRequest() (request *HaVipDisassociateAddressIpRequest) {
	request = &HaVipDisassociateAddressIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "HaVipDisassociateAddressIp")
	return
}

func NewHaVipDisassociateAddressIpResponse() (response *HaVipDisassociateAddressIpResponse) {
	response = &HaVipDisassociateAddressIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（HaVipDisassociateAddressIp）用于将高可用虚拟IP（HAVIP）已绑定的弹性公网IP（EIP）解除绑定<br />
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`RequestId`轮询`QueryTask`接口
func (c *Client) HaVipDisassociateAddressIp(request *HaVipDisassociateAddressIpRequest) (response *HaVipDisassociateAddressIpResponse, err error) {
	if request == nil {
		request = NewHaVipDisassociateAddressIpRequest()
	}
	response = NewHaVipDisassociateAddressIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficPackageStatisticsRequest() (request *DescribeTrafficPackageStatisticsRequest) {
	request = &DescribeTrafficPackageStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTrafficPackageStatistics")
	return
}

func NewDescribeTrafficPackageStatisticsResponse() (response *DescribeTrafficPackageStatisticsResponse) {
	response = &DescribeTrafficPackageStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询各地域共享流量包数量统计信息，包括全地域共享流量包总数、各地域的地域名称及其共享流量包总数等。
func (c *Client) DescribeTrafficPackageStatistics(request *DescribeTrafficPackageStatisticsRequest) (response *DescribeTrafficPackageStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficPackageStatisticsRequest()
	}
	response = NewDescribeTrafficPackageStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceVacancyAddressesRequest() (request *InquiryPriceVacancyAddressesRequest) {
	request = &InquiryPriceVacancyAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceVacancyAddresses")
	return
}

func NewInquiryPriceVacancyAddressesResponse() (response *InquiryPriceVacancyAddressesResponse) {
	response = &InquiryPriceVacancyAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 闲置EIP询价
func (c *Client) InquiryPriceVacancyAddresses(request *InquiryPriceVacancyAddressesRequest) (response *InquiryPriceVacancyAddressesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceVacancyAddressesRequest()
	}
	response = NewInquiryPriceVacancyAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLocalIpTranslationAclRuleRequest() (request *DeleteLocalIpTranslationAclRuleRequest) {
	request = &DeleteLocalIpTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalIpTranslationAclRule")
	return
}

func NewDeleteLocalIpTranslationAclRuleResponse() (response *DeleteLocalIpTranslationAclRuleResponse) {
	response = &DeleteLocalIpTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关本端IP转换ACL规则
func (c *Client) DeleteLocalIpTranslationAclRule(request *DeleteLocalIpTranslationAclRuleRequest) (response *DeleteLocalIpTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLocalIpTranslationAclRuleRequest()
	}
	response = NewDeleteLocalIpTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcEndPointRequest() (request *CreateVpcEndPointRequest) {
	request = &CreateVpcEndPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPoint")
	return
}

func NewCreateVpcEndPointResponse() (response *CreateVpcEndPointResponse) {
	response = &CreateVpcEndPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建终端节点。
func (c *Client) CreateVpcEndPoint(request *CreateVpcEndPointRequest) (response *CreateVpcEndPointResponse, err error) {
	if request == nil {
		request = NewCreateVpcEndPointRequest()
	}
	response = NewCreateVpcEndPointResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetDetectRequest() (request *CreateNetDetectRequest) {
	request = &CreateNetDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetDetect")
	return
}

func NewCreateNetDetectResponse() (response *CreateNetDetectResponse) {
	response = &CreateNetDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateNetDetect)用于创建网络探测。
func (c *Client) CreateNetDetect(request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
	if request == nil {
		request = NewCreateNetDetectRequest()
	}
	response = NewCreateNetDetectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTrafficMirrorRequest() (request *CreateTrafficMirrorRequest) {
	request = &CreateTrafficMirrorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateTrafficMirror")
	return
}

func NewCreateTrafficMirrorResponse() (response *CreateTrafficMirrorResponse) {
	response = &CreateTrafficMirrorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateTrafficMirror）用于创建流量镜像实例。
func (c *Client) CreateTrafficMirror(request *CreateTrafficMirrorRequest) (response *CreateTrafficMirrorResponse, err error) {
	if request == nil {
		request = NewCreateTrafficMirrorRequest()
	}
	response = NewCreateTrafficMirrorResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceTemplateRequest() (request *DeleteServiceTemplateRequest) {
	request = &DeleteServiceTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplate")
	return
}

func NewDeleteServiceTemplateResponse() (response *DeleteServiceTemplateResponse) {
	response = &DeleteServiceTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteServiceTemplate）用于删除协议端口模板
func (c *Client) DeleteServiceTemplate(request *DeleteServiceTemplateRequest) (response *DeleteServiceTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteServiceTemplateRequest()
	}
	response = NewDeleteServiceTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressAttributeRequest() (request *ModifyAddressAttributeRequest) {
	request = &ModifyAddressAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressAttribute")
	return
}

func NewModifyAddressAttributeResponse() (response *ModifyAddressAttributeResponse) {
	response = &ModifyAddressAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyAddressAttribute) 用于修改[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）的名称。
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAddressAttributeRequest()
	}
	response = NewModifyAddressAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCcnRoutesRequest() (request *EnableCcnRoutesRequest) {
	request = &EnableCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableCcnRoutes")
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

func NewDescribeAddressSetRequest() (request *DescribeAddressSetRequest) {
	request = &DescribeAddressSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressSet")
	return
}

func NewDescribeAddressSetResponse() (response *DescribeAddressSetResponse) {
	response = &DescribeAddressSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口支持查询弹性公网IP集群上的IP状态
func (c *Client) DescribeAddressSet(request *DescribeAddressSetRequest) (response *DescribeAddressSetResponse, err error) {
	if request == nil {
		request = NewDescribeAddressSetRequest()
	}
	response = NewDescribeAddressSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcPeeringConnectionsRequest() (request *DescribeVpcPeeringConnectionsRequest) {
	request = &DescribeVpcPeeringConnectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcPeeringConnections")
	return
}

func NewDescribeVpcPeeringConnectionsResponse() (response *DescribeVpcPeeringConnectionsResponse) {
	response = &DescribeVpcPeeringConnectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询私有网络对等连接
func (c *Client) DescribeVpcPeeringConnections(request *DescribeVpcPeeringConnectionsRequest) (response *DescribeVpcPeeringConnectionsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcPeeringConnectionsRequest()
	}
	response = NewDescribeVpcPeeringConnectionsResponse()
	err = c.Send(request, response)
	return
}

func NewAssignIpv6SubnetCidrBlockRequest() (request *AssignIpv6SubnetCidrBlockRequest) {
	request = &AssignIpv6SubnetCidrBlockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6SubnetCidrBlock")
	return
}

func NewAssignIpv6SubnetCidrBlockResponse() (response *AssignIpv6SubnetCidrBlockResponse) {
	response = &AssignIpv6SubnetCidrBlockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssignIpv6SubnetCidrBlock）用于分配IPv6子网段。
// * 给子网分配 `IPv6` 网段，要求子网所属 `VPC` 已获得 `IPv6` 网段。如果尚未分配，请先通过接口 `AssignIpv6CidrBlock` 给子网所属 `VPC` 分配一个 `IPv6` 网段。否则无法分配 `IPv6` 子网段。
// * 每个子网只能分配一个IPv6网段。
func (c *Client) AssignIpv6SubnetCidrBlock(request *AssignIpv6SubnetCidrBlockRequest) (response *AssignIpv6SubnetCidrBlockResponse, err error) {
	if request == nil {
		request = NewAssignIpv6SubnetCidrBlockRequest()
	}
	response = NewAssignIpv6SubnetCidrBlockResponse()
	err = c.Send(request, response)
	return
}

func NewRejectVpcPeeringConnectionRequest() (request *RejectVpcPeeringConnectionRequest) {
	request = &RejectVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RejectVpcPeeringConnection")
	return
}

func NewRejectVpcPeeringConnectionResponse() (response *RejectVpcPeeringConnectionResponse) {
	response = &RejectVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驳回同地域对等连接
func (c *Client) RejectVpcPeeringConnection(request *RejectVpcPeeringConnectionRequest) (response *RejectVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewRejectVpcPeeringConnectionRequest()
	}
	response = NewRejectVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupLimitsRequest() (request *DescribeSecurityGroupLimitsRequest) {
	request = &DescribeSecurityGroupLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupLimits")
	return
}

func NewDescribeSecurityGroupLimitsResponse() (response *DescribeSecurityGroupLimitsResponse) {
	response = &DescribeSecurityGroupLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeSecurityGroupLimits)用于查询用户安全组配额。
func (c *Client) DescribeSecurityGroupLimits(request *DescribeSecurityGroupLimitsRequest) (response *DescribeSecurityGroupLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupLimitsRequest()
	}
	response = NewDescribeSecurityGroupLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLocalIpTranslationAclRuleRequest() (request *ModifyLocalIpTranslationAclRuleRequest) {
	request = &ModifyLocalIpTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalIpTranslationAclRule")
	return
}

func NewModifyLocalIpTranslationAclRuleResponse() (response *ModifyLocalIpTranslationAclRuleResponse) {
	response = &ModifyLocalIpTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关本端IP转换ACL规则
func (c *Client) ModifyLocalIpTranslationAclRule(request *ModifyLocalIpTranslationAclRuleRequest) (response *ModifyLocalIpTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewModifyLocalIpTranslationAclRuleRequest()
	}
	response = NewModifyLocalIpTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLocalDestinationIpPortTranslationNatRuleRequest() (request *CreateLocalDestinationIpPortTranslationNatRuleRequest) {
	request = &CreateLocalDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalDestinationIpPortTranslationNatRule")
	return
}

func NewCreateLocalDestinationIpPortTranslationNatRuleResponse() (response *CreateLocalDestinationIpPortTranslationNatRuleResponse) {
	response = &CreateLocalDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关本端目的IP端口转换
func (c *Client) CreateLocalDestinationIpPortTranslationNatRule(request *CreateLocalDestinationIpPortTranslationNatRuleRequest) (response *CreateLocalDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreateLocalDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewCreateLocalDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
	request = &DescribeVpnGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGateways")
	return
}

func NewDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
	response = &DescribeVpnGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpnGateways）用于查询VPN网关列表。
func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewaysRequest()
	}
	response = NewDescribeVpnGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewBandwidthPackageRequest() (request *InquiryPriceRenewBandwidthPackageRequest) {
	request = &InquiryPriceRenewBandwidthPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewBandwidthPackage")
	return
}

func NewInquiryPriceRenewBandwidthPackageResponse() (response *InquiryPriceRenewBandwidthPackageResponse) {
	response = &InquiryPriceRenewBandwidthPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceRenewBandwidthPackage(request *InquiryPriceRenewBandwidthPackageRequest) (response *InquiryPriceRenewBandwidthPackageResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewBandwidthPackageRequest()
	}
	response = NewInquiryPriceRenewBandwidthPackageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpv6AddressesAttributeRequest() (request *ModifyIpv6AddressesAttributeRequest) {
	request = &ModifyIpv6AddressesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyIpv6AddressesAttribute")
	return
}

func NewModifyIpv6AddressesAttributeResponse() (response *ModifyIpv6AddressesAttributeResponse) {
	response = &ModifyIpv6AddressesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyIpv6AddressesAttribute）用于修改弹性网卡内网IPv6地址属性。
func (c *Client) ModifyIpv6AddressesAttribute(request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyIpv6AddressesAttributeRequest()
	}
	response = NewModifyIpv6AddressesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkInterfaceExRequest() (request *DeleteNetworkInterfaceExRequest) {
	request = &DeleteNetworkInterfaceExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkInterfaceEx")
	return
}

func NewDeleteNetworkInterfaceExResponse() (response *DeleteNetworkInterfaceExResponse) {
	response = &DeleteNetworkInterfaceExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNetworkInterfaceEx）用于删除跨租户弹性网卡。
func (c *Client) DeleteNetworkInterfaceEx(request *DeleteNetworkInterfaceExRequest) (response *DeleteNetworkInterfaceExResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkInterfaceExRequest()
	}
	response = NewDeleteNetworkInterfaceExResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCcnRoutesRequest() (request *DisableCcnRoutesRequest) {
	request = &DisableCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisableCcnRoutes")
	return
}

func NewDisableCcnRoutesResponse() (response *DisableCcnRoutesResponse) {
	response = &DisableCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableCcnRoutes）用于禁用已经启用的云联网（CCN）路由
func (c *Client) DisableCcnRoutes(request *DisableCcnRoutesRequest) (response *DisableCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDisableCcnRoutesRequest()
	}
	response = NewDisableCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalIpTranslationNatRulesRequest() (request *DescribeLocalIpTranslationNatRulesRequest) {
	request = &DescribeLocalIpTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalIpTranslationNatRules")
	return
}

func NewDescribeLocalIpTranslationNatRulesResponse() (response *DescribeLocalIpTranslationNatRulesResponse) {
	response = &DescribeLocalIpTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关本端IP转换
func (c *Client) DescribeLocalIpTranslationNatRules(request *DescribeLocalIpTranslationNatRulesRequest) (response *DescribeLocalIpTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLocalIpTranslationNatRulesRequest()
	}
	response = NewDescribeLocalIpTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadCustomerGatewayConfigurationRequest() (request *DownloadCustomerGatewayConfigurationRequest) {
	request = &DownloadCustomerGatewayConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DownloadCustomerGatewayConfiguration")
	return
}

func NewDownloadCustomerGatewayConfigurationResponse() (response *DownloadCustomerGatewayConfigurationResponse) {
	response = &DownloadCustomerGatewayConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DownloadCustomerGatewayConfiguration)用于下载VPN通道配置。
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
	if request == nil {
		request = NewDownloadCustomerGatewayConfigurationRequest()
	}
	response = NewDownloadCustomerGatewayConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRouteTableAttributeRequest() (request *ModifyRouteTableAttributeRequest) {
	request = &ModifyRouteTableAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyRouteTableAttribute")
	return
}

func NewModifyRouteTableAttributeResponse() (response *ModifyRouteTableAttributeResponse) {
	response = &ModifyRouteTableAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyRouteTableAttribute）用于修改路由表（RouteTable）属性。
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
	if request == nil {
		request = NewModifyRouteTableAttributeRequest()
	}
	response = NewModifyRouteTableAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePeerIpTranslationNatRulesRequest() (request *DescribePeerIpTranslationNatRulesRequest) {
	request = &DescribePeerIpTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribePeerIpTranslationNatRules")
	return
}

func NewDescribePeerIpTranslationNatRulesResponse() (response *DescribePeerIpTranslationNatRulesResponse) {
	response = &DescribePeerIpTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关对端IP转换
func (c *Client) DescribePeerIpTranslationNatRules(request *DescribePeerIpTranslationNatRulesRequest) (response *DescribePeerIpTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribePeerIpTranslationNatRulesRequest()
	}
	response = NewDescribePeerIpTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupPolicyTemplatesRequest() (request *DescribeSecurityGroupPolicyTemplatesRequest) {
	request = &DescribeSecurityGroupPolicyTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupPolicyTemplates")
	return
}

func NewDescribeSecurityGroupPolicyTemplatesResponse() (response *DescribeSecurityGroupPolicyTemplatesResponse) {
	response = &DescribeSecurityGroupPolicyTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSecurityGroupPolicyTemplates）用于查询安全组规则模板列表。
func (c *Client) DescribeSecurityGroupPolicyTemplates(request *DescribeSecurityGroupPolicyTemplatesRequest) (response *DescribeSecurityGroupPolicyTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupPolicyTemplatesRequest()
	}
	response = NewDescribeSecurityGroupPolicyTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptVpcPeeringConnectionExRequest() (request *AcceptVpcPeeringConnectionExRequest) {
	request = &AcceptVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AcceptVpcPeeringConnectionEx")
	return
}

func NewAcceptVpcPeeringConnectionExResponse() (response *AcceptVpcPeeringConnectionExResponse) {
	response = &AcceptVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接受跨地域对等连接
func (c *Client) AcceptVpcPeeringConnectionEx(request *AcceptVpcPeeringConnectionExRequest) (response *AcceptVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewAcceptVpcPeeringConnectionExRequest()
	}
	response = NewAcceptVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAnycastRegionRequest() (request *DescribeAnycastRegionRequest) {
	request = &DescribeAnycastRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAnycastRegion")
	return
}

func NewDescribeAnycastRegionResponse() (response *DescribeAnycastRegionResponse) {
	response = &DescribeAnycastRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeAnycastRegion)用于查询当前支持AnycastEIP的地域信息.
func (c *Client) DescribeAnycastRegion(request *DescribeAnycastRegionRequest) (response *DescribeAnycastRegionResponse, err error) {
	if request == nil {
		request = NewDescribeAnycastRegionRequest()
	}
	response = NewDescribeAnycastRegionResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSecurityGroupPoliciesRequest() (request *CheckSecurityGroupPoliciesRequest) {
	request = &CheckSecurityGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckSecurityGroupPolicies")
	return
}

func NewCheckSecurityGroupPoliciesResponse() (response *CheckSecurityGroupPoliciesResponse) {
	response = &CheckSecurityGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckSecurityGroupPolicies）用于查询安全组策略中常用端口的放开策略。
func (c *Client) CheckSecurityGroupPolicies(request *CheckSecurityGroupPoliciesRequest) (response *CheckSecurityGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewCheckSecurityGroupPoliciesRequest()
	}
	response = NewCheckSecurityGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceNetworkInterfaceRequest() (request *DeleteInstanceNetworkInterfaceRequest) {
	request = &DeleteInstanceNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteInstanceNetworkInterface")
	return
}

func NewDeleteInstanceNetworkInterfaceResponse() (response *DeleteInstanceNetworkInterfaceResponse) {
	response = &DeleteInstanceNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteInstanceNetworkInterface）用于删除云主机例弹性网卡。
// * 本接口在CVM云主机销毁时调用。
func (c *Client) DeleteInstanceNetworkInterface(request *DeleteInstanceNetworkInterfaceRequest) (response *DeleteInstanceNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceNetworkInterfaceRequest()
	}
	response = NewDeleteInstanceNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRequest() (request *DeleteCcnRequest) {
	request = &DeleteCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteCcn")
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

func NewDescribeTemplateLimitsRequest() (request *DescribeTemplateLimitsRequest) {
	request = &DescribeTemplateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTemplateLimits")
	return
}

func NewDescribeTemplateLimitsResponse() (response *DescribeTemplateLimitsResponse) {
	response = &DescribeTemplateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeTemplateLimits）用于查询参数模板配额列表。
func (c *Client) DescribeTemplateLimits(request *DescribeTemplateLimitsRequest) (response *DescribeTemplateLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateLimitsRequest()
	}
	response = NewDescribeTemplateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBandwidthPackageRequest() (request *DeleteBandwidthPackageRequest) {
	request = &DeleteBandwidthPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteBandwidthPackage")
	return
}

func NewDeleteBandwidthPackageResponse() (response *DeleteBandwidthPackageResponse) {
	response = &DeleteBandwidthPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口支持删除共享带宽包，包括设备带宽包和IP带宽包
func (c *Client) DeleteBandwidthPackage(request *DeleteBandwidthPackageRequest) (response *DeleteBandwidthPackageResponse, err error) {
	if request == nil {
		request = NewDeleteBandwidthPackageRequest()
	}
	response = NewDeleteBandwidthPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecificTrafficPackageUsedDetailsRequest() (request *DescribeSpecificTrafficPackageUsedDetailsRequest) {
	request = &DescribeSpecificTrafficPackageUsedDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSpecificTrafficPackageUsedDetails")
	return
}

func NewDescribeSpecificTrafficPackageUsedDetailsResponse() (response *DescribeSpecificTrafficPackageUsedDetailsResponse) {
	response = &DescribeSpecificTrafficPackageUsedDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeSpecificTrafficPackageUsedDetails) 用于查询指定 共享流量包 的用量明细。
func (c *Client) DescribeSpecificTrafficPackageUsedDetails(request *DescribeSpecificTrafficPackageUsedDetailsRequest) (response *DescribeSpecificTrafficPackageUsedDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeSpecificTrafficPackageUsedDetailsRequest()
	}
	response = NewDescribeSpecificTrafficPackageUsedDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressTemplateAttributeRequest() (request *ModifyAddressTemplateAttributeRequest) {
	request = &ModifyAddressTemplateAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateAttribute")
	return
}

func NewModifyAddressTemplateAttributeResponse() (response *ModifyAddressTemplateAttributeResponse) {
	response = &ModifyAddressTemplateAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAddressTemplateAttribute）用于修改IP地址模板
func (c *Client) ModifyAddressTemplateAttribute(request *ModifyAddressTemplateAttributeRequest) (response *ModifyAddressTemplateAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAddressTemplateAttributeRequest()
	}
	response = NewModifyAddressTemplateAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAssistantCidrRequest() (request *DeleteAssistantCidrRequest) {
	request = &DeleteAssistantCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteAssistantCidr")
	return
}

func NewDeleteAssistantCidrResponse() (response *DeleteAssistantCidrResponse) {
	response = &DeleteAssistantCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteAssistantCidr)用于删除辅助CIDR。（接口灰度中，如需使用请提工单。）
func (c *Client) DeleteAssistantCidr(request *DeleteAssistantCidrRequest) (response *DeleteAssistantCidrResponse, err error) {
	if request == nil {
		request = NewDeleteAssistantCidrRequest()
	}
	response = NewDeleteAssistantCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBandwidthPackagesRequest() (request *DescribeBandwidthPackagesRequest) {
	request = &DescribeBandwidthPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackages")
	return
}

func NewDescribeBandwidthPackagesResponse() (response *DescribeBandwidthPackagesResponse) {
	response = &DescribeBandwidthPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询带宽包详细信息，包括带宽包唯一标识ID，类型，计费模式，名称，资源信息等
func (c *Client) DescribeBandwidthPackages(request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
	if request == nil {
		request = NewDescribeBandwidthPackagesRequest()
	}
	response = NewDescribeBandwidthPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLocalSourceIpPortTranslationAclRuleRequest() (request *ModifyLocalSourceIpPortTranslationAclRuleRequest) {
	request = &ModifyLocalSourceIpPortTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalSourceIpPortTranslationAclRule")
	return
}

func NewModifyLocalSourceIpPortTranslationAclRuleResponse() (response *ModifyLocalSourceIpPortTranslationAclRuleResponse) {
	response = &ModifyLocalSourceIpPortTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关本端源IP端口转换ACL规则
func (c *Client) ModifyLocalSourceIpPortTranslationAclRule(request *ModifyLocalSourceIpPortTranslationAclRuleRequest) (response *ModifyLocalSourceIpPortTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewModifyLocalSourceIpPortTranslationAclRuleRequest()
	}
	response = NewModifyLocalSourceIpPortTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAttachCcnInstancesRequest() (request *AttachCcnInstancesRequest) {
	request = &AttachCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AttachCcnInstances")
	return
}

func NewAttachCcnInstancesResponse() (response *AttachCcnInstancesResponse) {
	response = &AttachCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachCcnInstances）用于将网络实例加载到云联网实例中，网络实例包括VPC和专线网关。<br />
// 每个云联网能够关联的网络实例个数是有限的，详请参考产品文档。如果需要扩充请联系在线客服。
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
	if request == nil {
		request = NewAttachCcnInstancesRequest()
	}
	response = NewAttachCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountAttributesRequest() (request *DescribeAccountAttributesRequest) {
	request = &DescribeAccountAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAccountAttributes")
	return
}

func NewDescribeAccountAttributesResponse() (response *DescribeAccountAttributesResponse) {
	response = &DescribeAccountAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAccountAttributes）用于查询用户账号私有属性。
func (c *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeAccountAttributesRequest()
	}
	response = NewDescribeAccountAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressTemplateInstancesRequest() (request *DescribeAddressTemplateInstancesRequest) {
	request = &DescribeAddressTemplateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplateInstances")
	return
}

func NewDescribeAddressTemplateInstancesResponse() (response *DescribeAddressTemplateInstancesResponse) {
	response = &DescribeAddressTemplateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAddressTemplateInstances）用于查询参数模板IP地址关联的实例列表。本接口不会返回查询的结果，需要根据返回的RequestId调用DescribeVpcTaskResult接口获取结果。
func (c *Client) DescribeAddressTemplateInstances(request *DescribeAddressTemplateInstancesRequest) (response *DescribeAddressTemplateInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeAddressTemplateInstancesRequest()
	}
	response = NewDescribeAddressTemplateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
	request = &DeleteVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpc")
	return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
	response = &DeleteVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteVpc）用于删除私有网络。
// * 删除前请确保 VPC 内已经没有相关资源，例如云服务器、云数据库、NoSQL、VPN网关、专线网关、负载均衡、对等连接、与之互通的基础网络设备等。
// * 删除私有网络是不可逆的操作，请谨慎处理。
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
	if request == nil {
		request = NewDeleteVpcRequest()
	}
	response = NewDeleteVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcLimitsRequest() (request *DescribeVpcLimitsRequest) {
	request = &DescribeVpcLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcLimits")
	return
}

func NewDescribeVpcLimitsResponse() (response *DescribeVpcLimitsResponse) {
	response = &DescribeVpcLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取私有网络配额，部分私有网络的配额有地域属性。
// LimitTypes取值范围：
// * appid-max-vpcs （每个开发商每个地域可创建的VPC数）
// * vpc-max-subnets（每个VPC可创建的子网数）
// * vpc-max-route-tables（每个VPC可创建的路由表数）
// * route-table-max-policies（每个路由表可添加的策略数）
// * vpc-max-vpn-gateways（每个VPC可创建的VPN网关数）
// * appid-max-custom-gateways（每个开发商可创建的对端网关数）
// * appid-max-vpn-connections（每个开发商可创建的VPN通道数）
// * custom-gateway-max-vpn-connections（每个对端网关可创建的VPN通道数）
// * vpn-gateway-max-custom-gateways（每个VPNGW可以创建的通道数）
// * vpc-max-network-acls（每个VPC可创建的网络ACL数）
// * network-acl-max-inbound-policies（每个网络ACL可添加的入站规则数）
// * network-acl-max-outbound-policies（每个网络ACL可添加的出站规则数）
// * vpc-max-vpcpeers（每个VPC可创建的对等连接数）
// * vpc-max-available-vpcpeers（每个VPC可创建的有效对等连接数）
// * vpc-max-basic-network-interconnections（每个VPC可创建的基础网络云主机与VPC互通数）
// * direct-connection-max-snats（每个专线网关可创建的SNAT数）
// * direct-connection-max-dnats（每个专线网关可创建的DNAT数）
// * direct-connection-max-snapts（每个专线网关可创建的SNAPT数）
// * direct-connection-max-dnapts（每个专线网关可创建的DNAPT数）
// * vpc-max-nat-gateways（每个VPC可创建的NAT网关数）
// * nat-gateway-max-eips（每个NAT可以购买的外网IP数量）
// * vpc-max-enis（每个VPC可创建弹性网卡数）
// * vpc-max-havips（每个VPC可创建HAVIP数）
// * eni-max-private-ips（每个ENI可以绑定的内网IP数（ENI未绑定子机））
// * nat-gateway-max-dnapts（每个NAT网关可创建的DNAPT数）
// * vpc-max-ipv6s（每个VPC可分配的IPv6地址数）
// * eni-max-ipv6s（每个ENI可分配的IPv6地址数）
// * vpc-max-assistant_cidrs（每个VPC可分配的辅助CIDR数）
func (c *Client) DescribeVpcLimits(request *DescribeVpcLimitsRequest) (response *DescribeVpcLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcLimitsRequest()
	}
	response = NewDescribeVpcLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpnGatewaySSLClientRequest() (request *ModifyVpnGatewaySSLClientRequest) {
	request = &ModifyVpnGatewaySSLClientRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewaySSLClient")
	return
}

func NewModifyVpnGatewaySSLClientResponse() (response *ModifyVpnGatewaySSLClientResponse) {
	response = &ModifyVpnGatewaySSLClientResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyVpnGatewaySSLClient
func (c *Client) ModifyVpnGatewaySSLClient(request *ModifyVpnGatewaySSLClientRequest) (response *ModifyVpnGatewaySSLClientResponse, err error) {
	if request == nil {
		request = NewModifyVpnGatewaySSLClientRequest()
	}
	response = NewModifyVpnGatewaySSLClientResponse()
	err = c.Send(request, response)
	return
}

func NewUnassignIpv6SubnetCidrBlockRequest() (request *UnassignIpv6SubnetCidrBlockRequest) {
	request = &UnassignIpv6SubnetCidrBlockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6SubnetCidrBlock")
	return
}

func NewUnassignIpv6SubnetCidrBlockResponse() (response *UnassignIpv6SubnetCidrBlockResponse) {
	response = &UnassignIpv6SubnetCidrBlockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnassignIpv6SubnetCidrBlock）用于释放IPv6子网段。<br />
// 子网段如果还有IP占用且未回收，则子网段无法释放。
func (c *Client) UnassignIpv6SubnetCidrBlock(request *UnassignIpv6SubnetCidrBlockRequest) (response *UnassignIpv6SubnetCidrBlockResponse, err error) {
	if request == nil {
		request = NewUnassignIpv6SubnetCidrBlockRequest()
	}
	response = NewUnassignIpv6SubnetCidrBlockResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkAclAttributeRequest() (request *ModifyNetworkAclAttributeRequest) {
	request = &ModifyNetworkAclAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAclAttribute")
	return
}

func NewModifyNetworkAclAttributeResponse() (response *ModifyNetworkAclAttributeResponse) {
	response = &ModifyNetworkAclAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNetworkAclAttribute）用于修改网络ACL属性。
func (c *Client) ModifyNetworkAclAttribute(request *ModifyNetworkAclAttributeRequest) (response *ModifyNetworkAclAttributeResponse, err error) {
	if request == nil {
		request = NewModifyNetworkAclAttributeRequest()
	}
	response = NewModifyNetworkAclAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewAddressesRequest() (request *InquiryPriceRenewAddressesRequest) {
	request = &InquiryPriceRenewAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewAddresses")
	return
}

func NewInquiryPriceRenewAddressesResponse() (response *InquiryPriceRenewAddressesResponse) {
	response = &InquiryPriceRenewAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceRenewAddresses(request *InquiryPriceRenewAddressesRequest) (response *InquiryPriceRenewAddressesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewAddressesRequest()
	}
	response = NewInquiryPriceRenewAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpLocationRequest() (request *DescribeIpLocationRequest) {
	request = &DescribeIpLocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpLocation")
	return
}

func NewDescribeIpLocationResponse() (response *DescribeIpLocationResponse) {
	response = &DescribeIpLocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口仅面向腾讯内部用户账号，用于查询IP地址信息，包括地理位置信息和网络信息。
func (c *Client) DescribeIpLocation(request *DescribeIpLocationRequest) (response *DescribeIpLocationResponse, err error) {
	if request == nil {
		request = NewDescribeIpLocationRequest()
	}
	response = NewDescribeIpLocationResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceNatGatewayRequest() (request *InquiryPriceNatGatewayRequest) {
	request = &InquiryPriceNatGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceNatGateway")
	return
}

func NewInquiryPriceNatGatewayResponse() (response *InquiryPriceNatGatewayResponse) {
	response = &InquiryPriceNatGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceNatGateway）用于Nat网关的询价。
func (c *Client) InquiryPriceNatGateway(request *InquiryPriceNatGatewayRequest) (response *InquiryPriceNatGatewayResponse, err error) {
	if request == nil {
		request = NewInquiryPriceNatGatewayRequest()
	}
	response = NewInquiryPriceNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssistantCidrRequest() (request *CreateAssistantCidrRequest) {
	request = &CreateAssistantCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateAssistantCidr")
	return
}

func NewCreateAssistantCidrResponse() (response *CreateAssistantCidrResponse) {
	response = &CreateAssistantCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateAssistantCidr)用于批量创建辅助CIDR。（接口灰度中，如需使用请提工单。）
func (c *Client) CreateAssistantCidr(request *CreateAssistantCidrRequest) (response *CreateAssistantCidrResponse, err error) {
	if request == nil {
		request = NewCreateAssistantCidrRequest()
	}
	response = NewCreateAssistantCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointServiceWhiteListRequest() (request *DescribeVpcEndPointServiceWhiteListRequest) {
	request = &DescribeVpcEndPointServiceWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPointServiceWhiteList")
	return
}

func NewDescribeVpcEndPointServiceWhiteListResponse() (response *DescribeVpcEndPointServiceWhiteListResponse) {
	response = &DescribeVpcEndPointServiceWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询终端节点服务的服务白名单。
func (c *Client) DescribeVpcEndPointServiceWhiteList(request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointServiceWhiteListRequest()
	}
	response = NewDescribeVpcEndPointServiceWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLocalDestinationIpPortTranslationNatRuleRequest() (request *ModifyLocalDestinationIpPortTranslationNatRuleRequest) {
	request = &ModifyLocalDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalDestinationIpPortTranslationNatRule")
	return
}

func NewModifyLocalDestinationIpPortTranslationNatRuleResponse() (response *ModifyLocalDestinationIpPortTranslationNatRuleResponse) {
	response = &ModifyLocalDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线网关本端目的IP端口转换
func (c *Client) ModifyLocalDestinationIpPortTranslationNatRule(request *ModifyLocalDestinationIpPortTranslationNatRuleRequest) (response *ModifyLocalDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyLocalDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewModifyLocalDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressesBandwidthRequest() (request *ModifyAddressesBandwidthRequest) {
	request = &ModifyAddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressesBandwidth")
	return
}

func NewModifyAddressesBandwidthResponse() (response *ModifyAddressesBandwidthResponse) {
	response = &ModifyAddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAddressesBandwidth）用于调整[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)(简称EIP)带宽，包括后付费EIP, 预付费EIP和带宽包EIP
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewModifyAddressesBandwidthRequest()
	}
	response = NewModifyAddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
	request = &AssociateAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateAddress")
	return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
	response = &AssociateAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AssociateAddress) 用于将[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）绑定到实例或弹性网卡的指定内网 IP 上。
// * 将 EIP 绑定到实例（CVM）上，其本质是将 EIP 绑定到实例上主网卡的主内网 IP 上。
// * 将 EIP 绑定到主网卡的主内网IP上，绑定过程会把其上绑定的普通公网 IP 自动解绑并释放。
// * 将 EIP 绑定到指定网卡的内网 IP上（非主网卡的主内网IP），则必须先解绑该 EIP，才能再绑定新的。
// * 将 EIP 绑定到NAT网关，请使用接口[AssociateNatGatewayAddress](../NAT网关相关接口/AssociateNatGatewayAddress)
// * EIP 如果欠费或被封堵，则不能被绑定。
// * 只有状态为 UNBIND 的 EIP 才能够被绑定。
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
	if request == nil {
		request = NewAssociateAddressRequest()
	}
	response = NewAssociateAddressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadSpecificTrafficPackageUsedDetailsQuotaRequest() (request *DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaRequest) {
	request = &DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDownloadSpecificTrafficPackageUsedDetailsQuota")
	return
}

func NewDescribeDownloadSpecificTrafficPackageUsedDetailsQuotaResponse() (response *DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaResponse) {
	response = &DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDownloadSpecificTrafficPackageUsedDetailsQuota)用户查询指定共享流量包在当前时间的明细文件生成配额.
func (c *Client) DescribeDownloadSpecificTrafficPackageUsedDetailsQuota(request *DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaRequest) (response *DescribeDownloadSpecificTrafficPackageUsedDetailsQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadSpecificTrafficPackageUsedDetailsQuotaRequest()
	}
	response = NewDescribeDownloadSpecificTrafficPackageUsedDetailsQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
	request = &DescribeRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRouteTables")
	return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
	response = &DescribeRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRouteTables）用于查询路由表。
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
	if request == nil {
		request = NewDescribeRouteTablesRequest()
	}
	response = NewDescribeRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
	request = &DescribeRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRoutes")
	return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
	response = &DescribeRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRoutes）用于查询路由列表。
func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) (response *DescribeRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeRoutesRequest()
	}
	response = NewDescribeRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewAssignPrivateIpAddressesRequest() (request *AssignPrivateIpAddressesRequest) {
	request = &AssignPrivateIpAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssignPrivateIpAddresses")
	return
}

func NewAssignPrivateIpAddressesResponse() (response *AssignPrivateIpAddressesResponse) {
	response = &AssignPrivateIpAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssignPrivateIpAddresses）用于弹性网卡申请内网 IP。
// * 一个弹性网卡支持绑定的IP地址是有限制的，更多资源限制信息详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
// * 可以指定内网IP地址申请，内网IP地址类型不能为主IP，主IP已存在，不能修改，内网IP必须要弹性网卡所在子网内，而且不能被占用。
// * 在弹性网卡上申请一个到多个辅助内网IP，接口会在弹性网卡所在子网网段内返回指定数量的辅助内网IP。
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
	if request == nil {
		request = NewAssignPrivateIpAddressesRequest()
	}
	response = NewAssignPrivateIpAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLocalIpTranslationNatRuleRequest() (request *CreateLocalIpTranslationNatRuleRequest) {
	request = &CreateLocalIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalIpTranslationNatRule")
	return
}

func NewCreateLocalIpTranslationNatRuleResponse() (response *CreateLocalIpTranslationNatRuleResponse) {
	response = &CreateLocalIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关本端IP转换
func (c *Client) CreateLocalIpTranslationNatRule(request *CreateLocalIpTranslationNatRuleRequest) (response *CreateLocalIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreateLocalIpTranslationNatRuleRequest()
	}
	response = NewCreateLocalIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpnGatewayCcnRoutesRequest() (request *ModifyVpnGatewayCcnRoutesRequest) {
	request = &ModifyVpnGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayCcnRoutes")
	return
}

func NewModifyVpnGatewayCcnRoutesResponse() (response *ModifyVpnGatewayCcnRoutesResponse) {
	response = &ModifyVpnGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyVpnGatewayCcnRoutes）用于修改VPN网关云联网路由
func (c *Client) ModifyVpnGatewayCcnRoutes(request *ModifyVpnGatewayCcnRoutesRequest) (response *ModifyVpnGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewModifyVpnGatewayCcnRoutesRequest()
	}
	response = NewModifyVpnGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnsRequest() (request *DescribeCcnsRequest) {
	request = &DescribeCcnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcns")
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

func NewModifyTrafficPackageAttributeRequest() (request *ModifyTrafficPackageAttributeRequest) {
	request = &ModifyTrafficPackageAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyTrafficPackageAttribute")
	return
}

func NewModifyTrafficPackageAttributeResponse() (response *ModifyTrafficPackageAttributeResponse) {
	response = &ModifyTrafficPackageAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于修改共享流量包属性，包括共享流量包名称等
func (c *Client) ModifyTrafficPackageAttribute(request *ModifyTrafficPackageAttributeRequest) (response *ModifyTrafficPackageAttributeResponse, err error) {
	if request == nil {
		request = NewModifyTrafficPackageAttributeRequest()
	}
	response = NewModifyTrafficPackageAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpnGatewaySSLServerRequest() (request *ModifyVpnGatewaySSLServerRequest) {
	request = &ModifyVpnGatewaySSLServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewaySSLServer")
	return
}

func NewModifyVpnGatewaySSLServerResponse() (response *ModifyVpnGatewaySSLServerResponse) {
	response = &ModifyVpnGatewaySSLServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyVpnGatewaySSLServer
func (c *Client) ModifyVpnGatewaySSLServer(request *ModifyVpnGatewaySSLServerRequest) (response *ModifyVpnGatewaySSLServerResponse, err error) {
	if request == nil {
		request = NewModifyVpnGatewaySSLServerRequest()
	}
	response = NewModifyVpnGatewaySSLServerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcPeeringConnectionRequest() (request *DeleteVpcPeeringConnectionRequest) {
	request = &DeleteVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcPeeringConnection")
	return
}

func NewDeleteVpcPeeringConnectionResponse() (response *DeleteVpcPeeringConnectionResponse) {
	response = &DeleteVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除同地域私有网络对等连接
func (c *Client) DeleteVpcPeeringConnection(request *DeleteVpcPeeringConnectionRequest) (response *DeleteVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewDeleteVpcPeeringConnectionRequest()
	}
	response = NewDeleteVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRoutesRequest() (request *EnableRoutesRequest) {
	request = &EnableRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableRoutes")
	return
}

func NewEnableRoutesResponse() (response *EnableRoutesResponse) {
	response = &EnableRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（EnableRoutes）用于启用已禁用的子网路由。<br />
// 本接口会校验启用后，是否与已有路由冲突，如果冲突，则无法启用，失败处理。路由冲突时，需要先禁用与之冲突的路由，才能启用该路由。
func (c *Client) EnableRoutes(request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
	if request == nil {
		request = NewEnableRoutesRequest()
	}
	response = NewEnableRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEipStatisticsRequest() (request *DescribeEipStatisticsRequest) {
	request = &DescribeEipStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeEipStatistics")
	return
}

func NewDescribeEipStatisticsResponse() (response *DescribeEipStatisticsResponse) {
	response = &DescribeEipStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询各地域EIP数量统计信息，包括全地域EIP总数，和各地域的地域名称及其EIP总数等。
func (c *Client) DescribeEipStatistics(request *DescribeEipStatisticsRequest) (response *DescribeEipStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeEipStatisticsRequest()
	}
	response = NewDescribeEipStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressAvailabilityRequest() (request *DescribeAddressAvailabilityRequest) {
	request = &DescribeAddressAvailabilityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressAvailability")
	return
}

func NewDescribeAddressAvailabilityResponse() (response *DescribeAddressAvailabilityResponse) {
	response = &DescribeAddressAvailabilityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性公网IP（EIP）库存状态
func (c *Client) DescribeAddressAvailability(request *DescribeAddressAvailabilityRequest) (response *DescribeAddressAvailabilityResponse, err error) {
	if request == nil {
		request = NewDescribeAddressAvailabilityRequest()
	}
	response = NewDescribeAddressAvailabilityResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcPeeringConnectionExRequest() (request *CreateVpcPeeringConnectionExRequest) {
	request = &CreateVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcPeeringConnectionEx")
	return
}

func NewCreateVpcPeeringConnectionExResponse() (response *CreateVpcPeeringConnectionExResponse) {
	response = &CreateVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建跨地域私有网络对等连接
func (c *Client) CreateVpcPeeringConnectionEx(request *CreateVpcPeeringConnectionExRequest) (response *CreateVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewCreateVpcPeeringConnectionExRequest()
	}
	response = NewCreateVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateVpcPeerConnectionRequest() (request *InquiryPriceCreateVpcPeerConnectionRequest) {
	request = &InquiryPriceCreateVpcPeerConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateVpcPeerConnection")
	return
}

func NewInquiryPriceCreateVpcPeerConnectionResponse() (response *InquiryPriceCreateVpcPeerConnectionResponse) {
	response = &InquiryPriceCreateVpcPeerConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 跨地域对等连接询价。
func (c *Client) InquiryPriceCreateVpcPeerConnection(request *InquiryPriceCreateVpcPeerConnectionRequest) (response *InquiryPriceCreateVpcPeerConnectionResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateVpcPeerConnectionRequest()
	}
	response = NewInquiryPriceCreateVpcPeerConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateNetworkAclSubnetsRequest() (request *AssociateNetworkAclSubnetsRequest) {
	request = &AssociateNetworkAclSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateNetworkAclSubnets")
	return
}

func NewAssociateNetworkAclSubnetsResponse() (response *AssociateNetworkAclSubnetsResponse) {
	response = &AssociateNetworkAclSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssociateNetworkAclSubnets）用于网络ACL关联vpc下的子网。
func (c *Client) AssociateNetworkAclSubnets(request *AssociateNetworkAclSubnetsRequest) (response *AssociateNetworkAclSubnetsResponse, err error) {
	if request == nil {
		request = NewAssociateNetworkAclSubnetsRequest()
	}
	response = NewAssociateNetworkAclSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcExtendCidrsRequest() (request *DescribeVpcExtendCidrsRequest) {
	request = &DescribeVpcExtendCidrsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcExtendCidrs")
	return
}

func NewDescribeVpcExtendCidrsResponse() (response *DescribeVpcExtendCidrsResponse) {
	response = &DescribeVpcExtendCidrsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询vpc扩展CIDR段
func (c *Client) DescribeVpcExtendCidrs(request *DescribeVpcExtendCidrsRequest) (response *DescribeVpcExtendCidrsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcExtendCidrsRequest()
	}
	response = NewDescribeVpcExtendCidrsResponse()
	err = c.Send(request, response)
	return
}

func NewResetRoutesRequest() (request *ResetRoutesRequest) {
	request = &ResetRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetRoutes")
	return
}

func NewResetRoutesResponse() (response *ResetRoutesResponse) {
	response = &ResetRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetRoutes）用于对某个路由表名称和所有路由策略（Route）进行重新设置。<br />
// 注意: 调用本接口是先删除当前路由表中所有路由策略, 再保存新提交的路由策略内容, 会引起网络中断。
func (c *Client) ResetRoutes(request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
	if request == nil {
		request = NewResetRoutesRequest()
	}
	response = NewResetRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewSetVpnGatewaysRenewFlagRequest() (request *SetVpnGatewaysRenewFlagRequest) {
	request = &SetVpnGatewaysRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "SetVpnGatewaysRenewFlag")
	return
}

func NewSetVpnGatewaysRenewFlagResponse() (response *SetVpnGatewaysRenewFlagResponse) {
	response = &SetVpnGatewaysRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置VPNGW续费标记
func (c *Client) SetVpnGatewaysRenewFlag(request *SetVpnGatewaysRenewFlagRequest) (response *SetVpnGatewaysRenewFlagResponse, err error) {
	if request == nil {
		request = NewSetVpnGatewaysRenewFlagRequest()
	}
	response = NewSetVpnGatewaysRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveBandwidthPackageResourcesRequest() (request *RemoveBandwidthPackageResourcesRequest) {
	request = &RemoveBandwidthPackageResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RemoveBandwidthPackageResources")
	return
}

func NewRemoveBandwidthPackageResourcesResponse() (response *RemoveBandwidthPackageResourcesResponse) {
	response = &RemoveBandwidthPackageResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于删除带宽包资源，包括[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)和[负载均衡](/tcloud/NetWork/CLB/218052/productdescription)等
func (c *Client) RemoveBandwidthPackageResources(request *RemoveBandwidthPackageResourcesRequest) (response *RemoveBandwidthPackageResourcesResponse, err error) {
	if request == nil {
		request = NewRemoveBandwidthPackageResourcesRequest()
	}
	response = NewRemoveBandwidthPackageResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoutesWithLocalRequest() (request *DescribeRoutesWithLocalRequest) {
	request = &DescribeRoutesWithLocalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRoutesWithLocal")
	return
}

func NewDescribeRoutesWithLocalResponse() (response *DescribeRoutesWithLocalResponse) {
	response = &DescribeRoutesWithLocalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRoutesWithLocal）用于用户查询路由表策略列表。
func (c *Client) DescribeRoutesWithLocal(request *DescribeRoutesWithLocalRequest) (response *DescribeRoutesWithLocalResponse, err error) {
	if request == nil {
		request = NewDescribeRoutesWithLocalRequest()
	}
	response = NewDescribeRoutesWithLocalResponse()
	err = c.Send(request, response)
	return
}

func NewStartTrafficMirrorRequest() (request *StartTrafficMirrorRequest) {
	request = &StartTrafficMirrorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "StartTrafficMirror")
	return
}

func NewStartTrafficMirrorResponse() (response *StartTrafficMirrorResponse) {
	response = &StartTrafficMirrorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StartTrafficMirror）用于开启流量镜像实例。
func (c *Client) StartTrafficMirror(request *StartTrafficMirrorRequest) (response *StartTrafficMirrorResponse, err error) {
	if request == nil {
		request = NewStartTrafficMirrorRequest()
	}
	response = NewStartTrafficMirrorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLocalSourceIpPortTranslationAclRuleRequest() (request *CreateLocalSourceIpPortTranslationAclRuleRequest) {
	request = &CreateLocalSourceIpPortTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalSourceIpPortTranslationAclRule")
	return
}

func NewCreateLocalSourceIpPortTranslationAclRuleResponse() (response *CreateLocalSourceIpPortTranslationAclRuleResponse) {
	response = &CreateLocalSourceIpPortTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关本端源IP端口转换ACL规则
func (c *Client) CreateLocalSourceIpPortTranslationAclRule(request *CreateLocalSourceIpPortTranslationAclRuleRequest) (response *CreateLocalSourceIpPortTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewCreateLocalSourceIpPortTranslationAclRuleRequest()
	}
	response = NewCreateLocalSourceIpPortTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCcnRegionBandwidthLimitsRequest() (request *DeleteCcnRegionBandwidthLimitsRequest) {
	request = &DeleteCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteCcnRegionBandwidthLimits")
	return
}

func NewDeleteCcnRegionBandwidthLimitsResponse() (response *DeleteCcnRegionBandwidthLimitsResponse) {
	response = &DeleteCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCcnRegionBandwidthLimits）用于删除云联网实例的限速规则。
func (c *Client) DeleteCcnRegionBandwidthLimits(request *DeleteCcnRegionBandwidthLimitsRequest) (response *DeleteCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewDeleteCcnRegionBandwidthLimitsRequest()
	}
	response = NewDeleteCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadSpecificTrafficPackageUsedDetailsRequest() (request *DescribeDownloadSpecificTrafficPackageUsedDetailsRequest) {
	request = &DescribeDownloadSpecificTrafficPackageUsedDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeDownloadSpecificTrafficPackageUsedDetails")
	return
}

func NewDescribeDownloadSpecificTrafficPackageUsedDetailsResponse() (response *DescribeDownloadSpecificTrafficPackageUsedDetailsResponse) {
	response = &DescribeDownloadSpecificTrafficPackageUsedDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDownloadSpecificTrafficPackageUsedDetails)用于查询指定共享流量包的用量明细生成历史
func (c *Client) DescribeDownloadSpecificTrafficPackageUsedDetails(request *DescribeDownloadSpecificTrafficPackageUsedDetailsRequest) (response *DescribeDownloadSpecificTrafficPackageUsedDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadSpecificTrafficPackageUsedDetailsRequest()
	}
	response = NewDescribeDownloadSpecificTrafficPackageUsedDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTemplatesRequest() (request *DescribeServiceTemplatesRequest) {
	request = &DescribeServiceTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplates")
	return
}

func NewDescribeServiceTemplatesResponse() (response *DescribeServiceTemplatesResponse) {
	response = &DescribeServiceTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeServiceTemplates）用于查询协议端口模板
func (c *Client) DescribeServiceTemplates(request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceTemplatesRequest()
	}
	response = NewDescribeServiceTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayFlowQosRequest() (request *DescribeGatewayFlowQosRequest) {
	request = &DescribeGatewayFlowQosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeGatewayFlowQos")
	return
}

func NewDescribeGatewayFlowQosResponse() (response *DescribeGatewayFlowQosResponse) {
	response = &DescribeGatewayFlowQosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeGatewayFlowQos）用于查询网关来访IP流控带宽。
func (c *Client) DescribeGatewayFlowQos(request *DescribeGatewayFlowQosRequest) (response *DescribeGatewayFlowQosResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayFlowQosRequest()
	}
	response = NewDescribeGatewayFlowQosResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecurityGroupWithPoliciesRequest() (request *CreateSecurityGroupWithPoliciesRequest) {
	request = &CreateSecurityGroupWithPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroupWithPolicies")
	return
}

func NewCreateSecurityGroupWithPoliciesResponse() (response *CreateSecurityGroupWithPoliciesResponse) {
	response = &CreateSecurityGroupWithPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSecurityGroupWithPolicies）用于创建新的安全组（SecurityGroup），并且可以同时添加安全组规则（SecurityGroupPolicy）。
// * 每个账户下每个地域的每个项目的<a href="https://cloud.tencent.com/document/product/213/12453">安全组数量限制</a>。
// * 新建的安全组的入站和出站规则默认都是全部拒绝，在创建后通常您需要再调用CreateSecurityGroupPolicies将安全组的规则设置为需要的规则。
//
// 安全组规则说明：
// * Version安全组规则版本号，用户每次更新安全规则版本会自动加1，防止您更新的路由规则已过期，不填不考虑冲突。
// * Protocol字段支持输入TCP, UDP, ICMP, ICMPV6, GRE, ALL。
// * CidrBlock字段允许输入符合cidr格式标准的任意字符串。(展开)在基础网络中，如果CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * Ipv6CidrBlock字段允许输入符合IPv6 cidr格式标准的任意字符串。(展开)在基础网络中，如果Ipv6CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IPv6，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * SecurityGroupId字段允许输入与待修改的安全组位于相同项目中的安全组ID，包括这个安全组ID本身，代表安全组下所有云服务器的内网IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个ID所关联的云服务器变化而变化，不需要重新修改。
// * Port字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当Protocol字段是TCP或UDP时，Port字段才被接受，即Protocol字段不是TCP或UDP时，Protocol和Port排他关系，不允许同时输入，否则会接口报错。
// * Action字段只允许输入ACCEPT或DROP。
// * CidrBlock, Ipv6CidrBlock, SecurityGroupId, AddressTemplate四者是排他关系，不允许同时输入，Protocol + Port和ServiceTemplate二者是排他关系，不允许同时输入。
// * 一次请求中只能创建单个方向的规则, 如果需要指定索引（PolicyIndex）参数, 多条规则的索引必须一致。
func (c *Client) CreateSecurityGroupWithPolicies(request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
	if request == nil {
		request = NewCreateSecurityGroupWithPoliciesRequest()
	}
	response = NewCreateSecurityGroupWithPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePeerIpTranslationNatRuleRequest() (request *DeletePeerIpTranslationNatRuleRequest) {
	request = &DeletePeerIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeletePeerIpTranslationNatRule")
	return
}

func NewDeletePeerIpTranslationNatRuleResponse() (response *DeletePeerIpTranslationNatRuleResponse) {
	response = &DeletePeerIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关对端IP转换
func (c *Client) DeletePeerIpTranslationNatRule(request *DeletePeerIpTranslationNatRuleRequest) (response *DeletePeerIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeletePeerIpTranslationNatRuleRequest()
	}
	response = NewDeletePeerIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcResourceDashboardRequest() (request *DescribeVpcResourceDashboardRequest) {
	request = &DescribeVpcResourceDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcResourceDashboard")
	return
}

func NewDescribeVpcResourceDashboardResponse() (response *DescribeVpcResourceDashboardResponse) {
	response = &DescribeVpcResourceDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeVpcResourceDashboard)用于查看VPC资源信息。
func (c *Client) DescribeVpcResourceDashboard(request *DescribeVpcResourceDashboardRequest) (response *DescribeVpcResourceDashboardResponse, err error) {
	if request == nil {
		request = NewDescribeVpcResourceDashboardRequest()
	}
	response = NewDescribeVpcResourceDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateNetworkInterfaceSecurityGroupsRequest() (request *AssociateNetworkInterfaceSecurityGroupsRequest) {
	request = &AssociateNetworkInterfaceSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssociateNetworkInterfaceSecurityGroups")
	return
}

func NewAssociateNetworkInterfaceSecurityGroupsResponse() (response *AssociateNetworkInterfaceSecurityGroupsResponse) {
	response = &AssociateNetworkInterfaceSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssociateNetworkInterfaceSecurityGroups）用于弹性网卡绑定安全组（SecurityGroup）。
func (c *Client) AssociateNetworkInterfaceSecurityGroups(request *AssociateNetworkInterfaceSecurityGroupsRequest) (response *AssociateNetworkInterfaceSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewAssociateNetworkInterfaceSecurityGroupsRequest()
	}
	response = NewAssociateNetworkInterfaceSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSingleIspRegionRequest() (request *DescribeSingleIspRegionRequest) {
	request = &DescribeSingleIspRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSingleIspRegion")
	return
}

func NewDescribeSingleIspRegionResponse() (response *DescribeSingleIspRegionResponse) {
	response = &DescribeSingleIspRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用区对应的EIP运营商信息，包括CMCC/CTCC/CUCC
func (c *Client) DescribeSingleIspRegion(request *DescribeSingleIspRegionRequest) (response *DescribeSingleIspRegionResponse, err error) {
	if request == nil {
		request = NewDescribeSingleIspRegionRequest()
	}
	response = NewDescribeSingleIspRegionResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpcPeeringConnectionRequest() (request *EnableVpcPeeringConnectionRequest) {
	request = &EnableVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableVpcPeeringConnection")
	return
}

func NewEnableVpcPeeringConnectionResponse() (response *EnableVpcPeeringConnectionResponse) {
	response = &EnableVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用过期同地域对等连接
func (c *Client) EnableVpcPeeringConnection(request *EnableVpcPeeringConnectionRequest) (response *EnableVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewEnableVpcPeeringConnectionRequest()
	}
	response = NewEnableVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCcnRegionBandwidthLimitsTypeRequest() (request *ModifyCcnRegionBandwidthLimitsTypeRequest) {
	request = &ModifyCcnRegionBandwidthLimitsTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyCcnRegionBandwidthLimitsType")
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

func NewDescribeLocalDestinationIpPortTranslationNatRulesRequest() (request *DescribeLocalDestinationIpPortTranslationNatRulesRequest) {
	request = &DescribeLocalDestinationIpPortTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalDestinationIpPortTranslationNatRules")
	return
}

func NewDescribeLocalDestinationIpPortTranslationNatRulesResponse() (response *DescribeLocalDestinationIpPortTranslationNatRulesResponse) {
	response = &DescribeLocalDestinationIpPortTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关本端目的IP端口转换
func (c *Client) DescribeLocalDestinationIpPortTranslationNatRules(request *DescribeLocalDestinationIpPortTranslationNatRulesRequest) (response *DescribeLocalDestinationIpPortTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLocalDestinationIpPortTranslationNatRulesRequest()
	}
	response = NewDescribeLocalDestinationIpPortTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointRequest() (request *DescribeVpcEndPointRequest) {
	request = &DescribeVpcEndPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPoint")
	return
}

func NewDescribeVpcEndPointResponse() (response *DescribeVpcEndPointResponse) {
	response = &DescribeVpcEndPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询终端节点详情。
func (c *Client) DescribeVpcEndPoint(request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointRequest()
	}
	response = NewDescribeVpcEndPointResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIp6TranslatorsRequest() (request *CreateIp6TranslatorsRequest) {
	request = &CreateIp6TranslatorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateIp6Translators")
	return
}

func NewCreateIp6TranslatorsResponse() (response *CreateIp6TranslatorsResponse) {
	response = &CreateIp6TranslatorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 1. 该接口用于创建IPV6转换IPV4实例，支持批量
// 2. 同一个账户在一个地域最多允许创建10个转换实例
func (c *Client) CreateIp6Translators(request *CreateIp6TranslatorsRequest) (response *CreateIp6TranslatorsResponse, err error) {
	if request == nil {
		request = NewCreateIp6TranslatorsRequest()
	}
	response = NewCreateIp6TranslatorsResponse()
	err = c.Send(request, response)
	return
}

func NewRenewVpnGatewayRequest() (request *RenewVpnGatewayRequest) {
	request = &RenewVpnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RenewVpnGateway")
	return
}

func NewRenewVpnGatewayResponse() (response *RenewVpnGatewayResponse) {
	response = &RenewVpnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RenewVpnGateway）用于预付费（包年包月）VPN网关续费。目前只支持IPSEC网关。
func (c *Client) RenewVpnGateway(request *RenewVpnGatewayRequest) (response *RenewVpnGatewayResponse, err error) {
	if request == nil {
		request = NewRenewVpnGatewayRequest()
	}
	response = NewRenewVpnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRegions）用于查询腾讯云地域（大区）列表，本接口没有入参，固定返回已开服的所有大区列表。
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response = NewDescribeRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcEndPointAttributeRequest() (request *ModifyVpcEndPointAttributeRequest) {
	request = &ModifyVpcEndPointAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointAttribute")
	return
}

func NewModifyVpcEndPointAttributeResponse() (response *ModifyVpcEndPointAttributeResponse) {
	response = &ModifyVpcEndPointAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改终端节点。
func (c *Client) ModifyVpcEndPointAttribute(request *ModifyVpcEndPointAttributeRequest) (response *ModifyVpcEndPointAttributeResponse, err error) {
	if request == nil {
		request = NewModifyVpcEndPointAttributeRequest()
	}
	response = NewModifyVpcEndPointAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAddressTemplateRequest() (request *CreateAddressTemplateRequest) {
	request = &CreateAddressTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplate")
	return
}

func NewCreateAddressTemplateResponse() (response *CreateAddressTemplateResponse) {
	response = &CreateAddressTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAddressTemplate）用于创建IP地址模版
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
	if request == nil {
		request = NewCreateAddressTemplateRequest()
	}
	response = NewCreateAddressTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNmsCidrsRequest() (request *DescribeNmsCidrsRequest) {
	request = &DescribeNmsCidrsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNmsCidrs")
	return
}

func NewDescribeNmsCidrsResponse() (response *DescribeNmsCidrsResponse) {
	response = &DescribeNmsCidrsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询非标网段
func (c *Client) DescribeNmsCidrs(request *DescribeNmsCidrsRequest) (response *DescribeNmsCidrsResponse, err error) {
	if request == nil {
		request = NewDescribeNmsCidrsRequest()
	}
	response = NewDescribeNmsCidrsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHaVipRequest() (request *CreateHaVipRequest) {
	request = &CreateHaVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateHaVip")
	return
}

func NewCreateHaVipResponse() (response *CreateHaVipResponse) {
	response = &CreateHaVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateHaVip）用于创建高可用虚拟IP（HAVIP）
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
	if request == nil {
		request = NewCreateHaVipRequest()
	}
	response = NewCreateHaVipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEIPIspInfoRequest() (request *DescribeEIPIspInfoRequest) {
	request = &DescribeEIPIspInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeEIPIspInfo")
	return
}

func NewDescribeEIPIspInfoResponse() (response *DescribeEIPIspInfoResponse) {
	response = &DescribeEIPIspInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeEIPIspInfo）用于查询EIP运营商信息
func (c *Client) DescribeEIPIspInfo(request *DescribeEIPIspInfoRequest) (response *DescribeEIPIspInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEIPIspInfoRequest()
	}
	response = NewDescribeEIPIspInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIp6TranslatorQuotaRequest() (request *DescribeIp6TranslatorQuotaRequest) {
	request = &DescribeIp6TranslatorQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIp6TranslatorQuota")
	return
}

func NewDescribeIp6TranslatorQuotaResponse() (response *DescribeIp6TranslatorQuotaResponse) {
	response = &DescribeIp6TranslatorQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户在指定地域IPV6转换实例和规则的配额
func (c *Client) DescribeIp6TranslatorQuota(request *DescribeIp6TranslatorQuotaRequest) (response *DescribeIp6TranslatorQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeIp6TranslatorQuotaRequest()
	}
	response = NewDescribeIp6TranslatorQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressChangeQuotaRequest() (request *DescribeAddressChangeQuotaRequest) {
	request = &DescribeAddressChangeQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressChangeQuota")
	return
}

func NewDescribeAddressChangeQuotaResponse() (response *DescribeAddressChangeQuotaResponse) {
	response = &DescribeAddressChangeQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询更换公网IP配额
func (c *Client) DescribeAddressChangeQuota(request *DescribeAddressChangeQuotaRequest) (response *DescribeAddressChangeQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAddressChangeQuotaRequest()
	}
	response = NewDescribeAddressChangeQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewayCcnRoutesRequest() (request *DescribeVpnGatewayCcnRoutesRequest) {
	request = &DescribeVpnGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewayCcnRoutes")
	return
}

func NewDescribeVpnGatewayCcnRoutesResponse() (response *DescribeVpnGatewayCcnRoutesResponse) {
	response = &DescribeVpnGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpnGatewayCcnRoutes）用于查询VPN网关云联网路由
func (c *Client) DescribeVpnGatewayCcnRoutes(request *DescribeVpnGatewayCcnRoutesRequest) (response *DescribeVpnGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewayCcnRoutesRequest()
	}
	response = NewDescribeVpnGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcPeeringConnectionRequest() (request *ModifyVpcPeeringConnectionRequest) {
	request = &ModifyVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcPeeringConnection")
	return
}

func NewModifyVpcPeeringConnectionResponse() (response *ModifyVpcPeeringConnectionResponse) {
	response = &ModifyVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改同地域对等连接属性
func (c *Client) ModifyVpcPeeringConnection(request *ModifyVpcPeeringConnectionRequest) (response *ModifyVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewModifyVpcPeeringConnectionRequest()
	}
	response = NewModifyVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBandwidthPackageResourcesRequest() (request *DescribeBandwidthPackageResourcesRequest) {
	request = &DescribeBandwidthPackageResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackageResources")
	return
}

func NewDescribeBandwidthPackageResourcesResponse() (response *DescribeBandwidthPackageResourcesResponse) {
	response = &DescribeBandwidthPackageResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeBandwidthPackageResources) 用于根据共享带宽包唯一ID查询共享带宽包内的资源列表，支持按条件过滤查询结果和分页查询。
func (c *Client) DescribeBandwidthPackageResources(request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeBandwidthPackageResourcesRequest()
	}
	response = NewDescribeBandwidthPackageResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyAddressInternetChargeTypeRequest() (request *InquiryPriceModifyAddressInternetChargeTypeRequest) {
	request = &InquiryPriceModifyAddressInternetChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceModifyAddressInternetChargeType")
	return
}

func NewInquiryPriceModifyAddressInternetChargeTypeResponse() (response *InquiryPriceModifyAddressInternetChargeTypeResponse) {
	response = &InquiryPriceModifyAddressInternetChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于调整上移账户弹性公网IP网络计费模式的差价查询
// * 支持BANDWIDTH_PREPAID_BY_MONTH和TRAFFIC_POSTPAID_BY_HOUR两种网络计费模式相互调整时的差价查询
func (c *Client) InquiryPriceModifyAddressInternetChargeType(request *InquiryPriceModifyAddressInternetChargeTypeRequest) (response *InquiryPriceModifyAddressInternetChargeTypeResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyAddressInternetChargeTypeRequest()
	}
	response = NewInquiryPriceModifyAddressInternetChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewRenewAddressesRequest() (request *RenewAddressesRequest) {
	request = &RenewAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RenewAddresses")
	return
}

func NewRenewAddressesResponse() (response *RenewAddressesResponse) {
	response = &RenewAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于续费包月带宽计费模式的弹性公网IP
func (c *Client) RenewAddresses(request *RenewAddressesRequest) (response *RenewAddressesResponse, err error) {
	if request == nil {
		request = NewRenewAddressesRequest()
	}
	response = NewRenewAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpLocationDownloadLinkRequest() (request *DescribeIpLocationDownloadLinkRequest) {
	request = &DescribeIpLocationDownloadLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpLocationDownloadLink")
	return
}

func NewDescribeIpLocationDownloadLinkResponse() (response *DescribeIpLocationDownloadLinkResponse) {
	response = &DescribeIpLocationDownloadLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口仅面向腾讯内部用户账号，用于获取ip地址库下载链接
func (c *Client) DescribeIpLocationDownloadLink(request *DescribeIpLocationDownloadLinkRequest) (response *DescribeIpLocationDownloadLinkResponse, err error) {
	if request == nil {
		request = NewDescribeIpLocationDownloadLinkRequest()
	}
	response = NewDescribeIpLocationDownloadLinkResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcIpv6AddressesRequest() (request *DescribeVpcIpv6AddressesRequest) {
	request = &DescribeVpcIpv6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcIpv6Addresses")
	return
}

func NewDescribeVpcIpv6AddressesResponse() (response *DescribeVpcIpv6AddressesResponse) {
	response = &DescribeVpcIpv6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcIpv6Addresses）用于查询 `VPC` `IPv6` 信息。
// 只能查询已使用的`IPv6`信息，当查询未使用的IP时，本接口不会报错，但不会出现在返回结果里。
func (c *Client) DescribeVpcIpv6Addresses(request *DescribeVpcIpv6AddressesRequest) (response *DescribeVpcIpv6AddressesResponse, err error) {
	if request == nil {
		request = NewDescribeVpcIpv6AddressesRequest()
	}
	response = NewDescribeVpcIpv6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpnGatewaySSLClientRequest() (request *DeleteVpnGatewaySSLClientRequest) {
	request = &DeleteVpnGatewaySSLClientRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGatewaySSLClient")
	return
}

func NewDeleteVpnGatewaySSLClientResponse() (response *DeleteVpnGatewaySSLClientResponse) {
	response = &DeleteVpnGatewaySSLClientResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteVpnGatewaySSLClient
func (c *Client) DeleteVpnGatewaySSLClient(request *DeleteVpnGatewaySSLClientRequest) (response *DeleteVpnGatewaySSLClientResponse, err error) {
	if request == nil {
		request = NewDeleteVpnGatewaySSLClientRequest()
	}
	response = NewDeleteVpnGatewaySSLClientResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCcnBandwidthRequest() (request *UpdateCcnBandwidthRequest) {
	request = &UpdateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "UpdateCcnBandwidth")
	return
}

func NewUpdateCcnBandwidthResponse() (response *UpdateCcnBandwidthResponse) {
	response = &UpdateCcnBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCcnBandwidth）用于变更预付费模式下云联网实例的地域间带宽
func (c *Client) UpdateCcnBandwidth(request *UpdateCcnBandwidthRequest) (response *UpdateCcnBandwidthResponse, err error) {
	if request == nil {
		request = NewUpdateCcnBandwidthRequest()
	}
	response = NewUpdateCcnBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewResetTrafficMirrorFilterRequest() (request *ResetTrafficMirrorFilterRequest) {
	request = &ResetTrafficMirrorFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetTrafficMirrorFilter")
	return
}

func NewResetTrafficMirrorFilterResponse() (response *ResetTrafficMirrorFilterResponse) {
	response = &ResetTrafficMirrorFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetTrafficMirrorFilter）用于更新流量镜像实例过滤规则。
// 注意：每一个流量镜像实例，不能同时支持按nat网关和五元组两种规则过滤
func (c *Client) ResetTrafficMirrorFilter(request *ResetTrafficMirrorFilterRequest) (response *ResetTrafficMirrorFilterResponse, err error) {
	if request == nil {
		request = NewResetTrafficMirrorFilterRequest()
	}
	response = NewResetTrafficMirrorFilterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDefaultVpcRequest() (request *CreateDefaultVpcRequest) {
	request = &CreateDefaultVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDefaultVpc")
	return
}

func NewCreateDefaultVpcResponse() (response *CreateDefaultVpcResponse) {
	response = &CreateDefaultVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDefaultVpc）用于创建默认私有网络(VPC）。
//
// 默认VPC适用于快速入门和启动公共实例，您可以像使用任何其他VPC一样使用默认VPC。如果您想创建标准VPC，即指定VPC名称、VPC网段、子网网段、子网可用区，请使用常规创建VPC接口（CreateVpc）
//
// 正常情况，本接口并不一定生产默认VPC，而是根据用户账号的网络属性（DescribeAccountAttributes）来决定的
// * 支持基础网络、VPC，返回VpcId为0
// * 只支持VPC，返回默认VPC信息
//
// 您也可以通过 Force 参数，强制返回默认VPC
func (c *Client) CreateDefaultVpc(request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
	if request == nil {
		request = NewCreateDefaultVpcRequest()
	}
	response = NewCreateDefaultVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRegionBandwidthLimitsRequest() (request *DescribeCcnRegionBandwidthLimitsRequest) {
	request = &DescribeCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnRegionBandwidthLimits")
	return
}

func NewDescribeCcnRegionBandwidthLimitsResponse() (response *DescribeCcnRegionBandwidthLimitsResponse) {
	response = &DescribeCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnRegionBandwidthLimits）用于查询云联网各地域出带宽上限，该接口只返回已关联网络实例包含的地域
func (c *Client) DescribeCcnRegionBandwidthLimits(request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRegionBandwidthLimitsRequest()
	}
	response = NewDescribeCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
	request = &DeleteRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutes")
	return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
	response = &DeleteRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteRoutes)用于对某个路由表批量删除路由策略（Route）。
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
	if request == nil {
		request = NewDeleteRoutesRequest()
	}
	response = NewDeleteRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewAttachClassicLinkVpcRequest() (request *AttachClassicLinkVpcRequest) {
	request = &AttachClassicLinkVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AttachClassicLinkVpc")
	return
}

func NewAttachClassicLinkVpcResponse() (response *AttachClassicLinkVpcResponse) {
	response = &AttachClassicLinkVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(AttachClassicLinkVpc)用于创建私有网络和基础网络设备互通。
// * 私有网络和基础网络设备必须在同一个地域。
// * 私有网络和基础网络的区别详见vpc产品文档-<a href="https://cloud.tencent.com/document/product/215/30720">私有网络与基础网络</a>。
func (c *Client) AttachClassicLinkVpc(request *AttachClassicLinkVpcRequest) (response *AttachClassicLinkVpcResponse, err error) {
	if request == nil {
		request = NewAttachClassicLinkVpcRequest()
	}
	response = NewAttachClassicLinkVpcResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcEndPointServiceWhiteListRequest() (request *CreateVpcEndPointServiceWhiteListRequest) {
	request = &CreateVpcEndPointServiceWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPointServiceWhiteList")
	return
}

func NewCreateVpcEndPointServiceWhiteListResponse() (response *CreateVpcEndPointServiceWhiteListResponse) {
	response = &CreateVpcEndPointServiceWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建终端服务白名单。
func (c *Client) CreateVpcEndPointServiceWhiteList(request *CreateVpcEndPointServiceWhiteListRequest) (response *CreateVpcEndPointServiceWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateVpcEndPointServiceWhiteListRequest()
	}
	response = NewCreateVpcEndPointServiceWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewAddBandwidthPackageResourcesRequest() (request *AddBandwidthPackageResourcesRequest) {
	request = &AddBandwidthPackageResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AddBandwidthPackageResources")
	return
}

func NewAddBandwidthPackageResourcesResponse() (response *AddBandwidthPackageResourcesResponse) {
	response = &AddBandwidthPackageResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于添加带宽包资源，包括[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)和[负载均衡](/tcloud/NetWork/CLB/218052/productdescription)等
func (c *Client) AddBandwidthPackageResources(request *AddBandwidthPackageResourcesRequest) (response *AddBandwidthPackageResourcesResponse, err error) {
	if request == nil {
		request = NewAddBandwidthPackageResourcesRequest()
	}
	response = NewAddBandwidthPackageResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetDetectRequest() (request *ModifyNetDetectRequest) {
	request = &ModifyNetDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetDetect")
	return
}

func NewModifyNetDetectResponse() (response *ModifyNetDetectResponse) {
	response = &ModifyNetDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyNetDetect)用于修改网络探测参数。
func (c *Client) ModifyNetDetect(request *ModifyNetDetectRequest) (response *ModifyNetDetectResponse, err error) {
	if request == nil {
		request = NewModifyNetDetectRequest()
	}
	response = NewModifyNetDetectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDetectStatesRequest() (request *DescribeNetDetectStatesRequest) {
	request = &DescribeNetDetectStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetDetectStates")
	return
}

func NewDescribeNetDetectStatesResponse() (response *DescribeNetDetectStatesResponse) {
	response = &DescribeNetDetectStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeNetDetectStates)用于查询网络探测验证结果列表。
func (c *Client) DescribeNetDetectStates(request *DescribeNetDetectStatesRequest) (response *DescribeNetDetectStatesResponse, err error) {
	if request == nil {
		request = NewDescribeNetDetectStatesRequest()
	}
	response = NewDescribeNetDetectStatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroups")
	return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSecurityGroups）用于查询安全组。
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupsRequest()
	}
	response = NewDescribeSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
	request = &AllocateAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AllocateAddresses")
	return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
	response = &AllocateAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AllocateAddresses) 用于申请一个或多个[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * EIP 是专为动态云计算设计的静态 IP 地址。借助 EIP，您可以快速将 EIP 重新映射到您的另一个实例上，从而屏蔽实例故障。
// * 您的 EIP 与云平台账户相关联，而不是与某个实例相关联。在您选择显式释放该地址，或欠费超过24小时之前，它会一直与您的云平台账户保持关联。
// * 一个云平台账户在每个地域能申请的 EIP 最大配额有所限制，可参见 [EIP 产品简介](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)，上述配额可通过 DescribeAddressQuota 接口获取。
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
	if request == nil {
		request = NewAllocateAddressesRequest()
	}
	response = NewAllocateAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkAclRequest() (request *CreateNetworkAclRequest) {
	request = &CreateNetworkAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkAcl")
	return
}

func NewCreateNetworkAclResponse() (response *CreateNetworkAclResponse) {
	response = &CreateNetworkAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateNetworkAcl）用于创建新的<a href="https://cloud.tencent.com/document/product/215/20088">网络ACL</a>。
// * 新建的网络ACL的入站和出站规则默认都是全部拒绝，在创建后通常您需要再调用ModifyNetworkAclEntries将网络ACL的规则设置为需要的规则。
func (c *Client) CreateNetworkAcl(request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
	if request == nil {
		request = NewCreateNetworkAclRequest()
	}
	response = NewCreateNetworkAclResponse()
	err = c.Send(request, response)
	return
}

func NewDisableGatewayFlowMonitorRequest() (request *DisableGatewayFlowMonitorRequest) {
	request = &DisableGatewayFlowMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisableGatewayFlowMonitor")
	return
}

func NewDisableGatewayFlowMonitorResponse() (response *DisableGatewayFlowMonitorResponse) {
	response = &DisableGatewayFlowMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisableGatewayFlowMonitor）用于关闭网关流量监控。
func (c *Client) DisableGatewayFlowMonitor(request *DisableGatewayFlowMonitorRequest) (response *DisableGatewayFlowMonitorResponse, err error) {
	if request == nil {
		request = NewDisableGatewayFlowMonitorRequest()
	}
	response = NewDisableGatewayFlowMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskRequest() (request *QueryTaskRequest) {
	request = &QueryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "QueryTask")
	return
}

func NewQueryTaskResponse() (response *QueryTaskResponse) {
	response = &QueryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询异步任务执行结果
func (c *Client) QueryTask(request *QueryTaskRequest) (response *QueryTaskResponse, err error) {
	if request == nil {
		request = NewQueryTaskRequest()
	}
	response = NewQueryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateVpnGatewayRequest() (request *InquiryPriceCreateVpnGatewayRequest) {
	request = &InquiryPriceCreateVpnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateVpnGateway")
	return
}

func NewInquiryPriceCreateVpnGatewayResponse() (response *InquiryPriceCreateVpnGatewayResponse) {
	response = &InquiryPriceCreateVpnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceCreateVpnGateway）用于创建VPN网关询价。
func (c *Client) InquiryPriceCreateVpnGateway(request *InquiryPriceCreateVpnGatewayRequest) (response *InquiryPriceCreateVpnGatewayResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateVpnGatewayRequest()
	}
	response = NewInquiryPriceCreateVpnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyIp6AddressesBandwidthRequest() (request *InquiryPriceModifyIp6AddressesBandwidthRequest) {
	request = &InquiryPriceModifyIp6AddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceModifyIp6AddressesBandwidth")
	return
}

func NewInquiryPriceModifyIp6AddressesBandwidthResponse() (response *InquiryPriceModifyIp6AddressesBandwidthResponse) {
	response = &InquiryPriceModifyIp6AddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询修改IPV6带宽的价格
func (c *Client) InquiryPriceModifyIp6AddressesBandwidth(request *InquiryPriceModifyIp6AddressesBandwidthRequest) (response *InquiryPriceModifyIp6AddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyIp6AddressesBandwidthRequest()
	}
	response = NewInquiryPriceModifyIp6AddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpnGatewayRequest() (request *CreateVpnGatewayRequest) {
	request = &CreateVpnGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGateway")
	return
}

func NewCreateVpnGatewayResponse() (response *CreateVpnGatewayResponse) {
	response = &CreateVpnGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateVpnGateway）用于创建VPN网关。
func (c *Client) CreateVpnGateway(request *CreateVpnGatewayRequest) (response *CreateVpnGatewayResponse, err error) {
	if request == nil {
		request = NewCreateVpnGatewayRequest()
	}
	response = NewCreateVpnGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyBandwidthPackageBandwidthRequest() (request *InquiryPriceModifyBandwidthPackageBandwidthRequest) {
	request = &InquiryPriceModifyBandwidthPackageBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceModifyBandwidthPackageBandwidth")
	return
}

func NewInquiryPriceModifyBandwidthPackageBandwidthResponse() (response *InquiryPriceModifyBandwidthPackageBandwidthResponse) {
	response = &InquiryPriceModifyBandwidthPackageBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) InquiryPriceModifyBandwidthPackageBandwidth(request *InquiryPriceModifyBandwidthPackageBandwidthRequest) (response *InquiryPriceModifyBandwidthPackageBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyBandwidthPackageBandwidthRequest()
	}
	response = NewInquiryPriceModifyBandwidthPackageBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkInterfaceExRequest() (request *CreateNetworkInterfaceExRequest) {
	request = &CreateNetworkInterfaceExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkInterfaceEx")
	return
}

func NewCreateNetworkInterfaceExResponse() (response *CreateNetworkInterfaceExResponse) {
	response = &CreateNetworkInterfaceExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateNetworkInterfaceEx）用于创建跨租户的弹性网卡。
// * ReservedAddress=0时，从保留网段（169.254.128/17）分配一个未使用的IP。
// * 创建弹性网卡时可以指定内网IP（ReservedAddress=0且指定了用户SubnetId时，才可使用），并且可以指定一个主IP，指定的内网IP必须在弹性网卡所在子网内，而且不能被占用。
// * 不允许网卡VPC和云服务器在同一个VPC内，调用接口创建跨租户的弹性网卡；
func (c *Client) CreateNetworkInterfaceEx(request *CreateNetworkInterfaceExRequest) (response *CreateNetworkInterfaceExResponse, err error) {
	if request == nil {
		request = NewCreateNetworkInterfaceExRequest()
	}
	response = NewCreateNetworkInterfaceExResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
	request = &DeleteRouteTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteRouteTable")
	return
}

func NewDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
	response = &DeleteRouteTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由表
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
	if request == nil {
		request = NewDeleteRouteTableRequest()
	}
	response = NewDeleteRouteTableResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomerGatewayAttributeRequest() (request *ModifyCustomerGatewayAttributeRequest) {
	request = &ModifyCustomerGatewayAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyCustomerGatewayAttribute")
	return
}

func NewModifyCustomerGatewayAttributeResponse() (response *ModifyCustomerGatewayAttributeResponse) {
	response = &ModifyCustomerGatewayAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyCustomerGatewayAttribute）用于修改对端网关信息。
func (c *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCustomerGatewayAttributeRequest()
	}
	response = NewModifyCustomerGatewayAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteConflictsRequest() (request *DescribeRouteConflictsRequest) {
	request = &DescribeRouteConflictsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeRouteConflicts")
	return
}

func NewDescribeRouteConflictsResponse() (response *DescribeRouteConflictsResponse) {
	response = &DescribeRouteConflictsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRouteConflicts）用于查询自定义路由策略与云联网路由策略冲突列表
func (c *Client) DescribeRouteConflicts(request *DescribeRouteConflictsRequest) (response *DescribeRouteConflictsResponse, err error) {
	if request == nil {
		request = NewDescribeRouteConflictsRequest()
	}
	response = NewDescribeRouteConflictsResponse()
	err = c.Send(request, response)
	return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
	request = &DetachNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DetachNetworkInterface")
	return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
	response = &DetachNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DetachNetworkInterface）用于弹性网卡解绑云主机。
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewDetachNetworkInterfaceRequest()
	}
	response = NewDetachNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateNetworkInterfaceSecurityGroupsRequest() (request *DisassociateNetworkInterfaceSecurityGroupsRequest) {
	request = &DisassociateNetworkInterfaceSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNetworkInterfaceSecurityGroups")
	return
}

func NewDisassociateNetworkInterfaceSecurityGroupsResponse() (response *DisassociateNetworkInterfaceSecurityGroupsResponse) {
	response = &DisassociateNetworkInterfaceSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisassociateNetworkInterfaceSecurityGroups）用于弹性网卡解绑安全组。支持弹性网卡完全解绑安全组。
func (c *Client) DisassociateNetworkInterfaceSecurityGroups(request *DisassociateNetworkInterfaceSecurityGroupsRequest) (response *DisassociateNetworkInterfaceSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDisassociateNetworkInterfaceSecurityGroupsRequest()
	}
	response = NewDisassociateNetworkInterfaceSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcPeeringConnectionExRequest() (request *DeleteVpcPeeringConnectionExRequest) {
	request = &DeleteVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcPeeringConnectionEx")
	return
}

func NewDeleteVpcPeeringConnectionExResponse() (response *DeleteVpcPeeringConnectionExResponse) {
	response = &DeleteVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除跨地域私有网络对等连接
func (c *Client) DeleteVpcPeeringConnectionEx(request *DeleteVpcPeeringConnectionExRequest) (response *DeleteVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewDeleteVpcPeeringConnectionExRequest()
	}
	response = NewDeleteVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecurityGroupAttributeRequest() (request *ModifySecurityGroupAttributeRequest) {
	request = &ModifySecurityGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupAttribute")
	return
}

func NewModifySecurityGroupAttributeResponse() (response *ModifySecurityGroupAttributeResponse) {
	response = &ModifySecurityGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySecurityGroupAttribute）用于修改安全组（SecurityGroupPolicy）属性。
func (c *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifySecurityGroupAttributeRequest()
	}
	response = NewModifySecurityGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkInterfaceExtendIpsRequest() (request *DescribeNetworkInterfaceExtendIpsRequest) {
	request = &DescribeNetworkInterfaceExtendIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaceExtendIps")
	return
}

func NewDescribeNetworkInterfaceExtendIpsResponse() (response *DescribeNetworkInterfaceExtendIpsResponse) {
	response = &DescribeNetworkInterfaceExtendIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性网卡的扩展ip
func (c *Client) DescribeNetworkInterfaceExtendIps(request *DescribeNetworkInterfaceExtendIpsRequest) (response *DescribeNetworkInterfaceExtendIpsResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkInterfaceExtendIpsRequest()
	}
	response = NewDescribeNetworkInterfaceExtendIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
	request = &DisassociateAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateAddress")
	return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
	response = &DisassociateAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DisassociateAddress) 用于解绑[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * 支持CVM实例，弹性网卡上的EIP解绑
// * 不支持NAT上的EIP解绑。NAT上的EIP解绑请参考[DisassociateNatGatewayAddress](../NAT网关相关接口/DisassociateNatGatewayAddress)
// * 只有状态为 BIND 和 BIND_ENI 的 EIP 才能进行解绑定操作。
// * EIP 如果被封堵，则不能进行解绑定操作。
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
	if request == nil {
		request = NewDisassociateAddressRequest()
	}
	response = NewDisassociateAddressResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatGatewayAttributeRequest() (request *ModifyNatGatewayAttributeRequest) {
	request = &ModifyNatGatewayAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewayAttribute")
	return
}

func NewModifyNatGatewayAttributeResponse() (response *ModifyNatGatewayAttributeResponse) {
	response = &ModifyNatGatewayAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNatGatewayAttribute）用于修改NAT网关的属性。
func (c *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
	if request == nil {
		request = NewModifyNatGatewayAttributeRequest()
	}
	response = NewModifyNatGatewayAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficPackageQuotaRequest() (request *DescribeTrafficPackageQuotaRequest) {
	request = &DescribeTrafficPackageQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTrafficPackageQuota")
	return
}

func NewDescribeTrafficPackageQuotaResponse() (response *DescribeTrafficPackageQuotaResponse) {
	response = &DescribeTrafficPackageQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询账户在当前地域的共享流量包上限数量以及使用数量
func (c *Client) DescribeTrafficPackageQuota(request *DescribeTrafficPackageQuotaRequest) (response *DescribeTrafficPackageQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficPackageQuotaRequest()
	}
	response = NewDescribeTrafficPackageQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFlowLogAttributeRequest() (request *ModifyFlowLogAttributeRequest) {
	request = &ModifyFlowLogAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyFlowLogAttribute")
	return
}

func NewModifyFlowLogAttributeResponse() (response *ModifyFlowLogAttributeResponse) {
	response = &ModifyFlowLogAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyFlowLogAttribute）用于修改流日志属性
func (c *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (response *ModifyFlowLogAttributeResponse, err error) {
	if request == nil {
		request = NewModifyFlowLogAttributeRequest()
	}
	response = NewModifyFlowLogAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLocalIpTranslationNatRuleRequest() (request *DeleteLocalIpTranslationNatRuleRequest) {
	request = &DeleteLocalIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalIpTranslationNatRule")
	return
}

func NewDeleteLocalIpTranslationNatRuleResponse() (response *DeleteLocalIpTranslationNatRuleResponse) {
	response = &DeleteLocalIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关本端IP转换
func (c *Client) DeleteLocalIpTranslationNatRule(request *DeleteLocalIpTranslationNatRuleRequest) (response *DeleteLocalIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLocalIpTranslationNatRuleRequest()
	}
	response = NewDeleteLocalIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewHaVipAssociateAddressIpRequest() (request *HaVipAssociateAddressIpRequest) {
	request = &HaVipAssociateAddressIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "HaVipAssociateAddressIp")
	return
}

func NewHaVipAssociateAddressIpResponse() (response *HaVipAssociateAddressIpResponse) {
	response = &HaVipAssociateAddressIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（HaVipAssociateAddressIp）用于高可用虚拟IP（HAVIP）绑定弹性公网IP（EIP）<br />
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`RequestId`轮询`QueryTask`接口
func (c *Client) HaVipAssociateAddressIp(request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
	if request == nil {
		request = NewHaVipAssociateAddressIpRequest()
	}
	response = NewHaVipAssociateAddressIpResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
	request = &ReleaseAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReleaseAddresses")
	return
}

func NewReleaseAddressesResponse() (response *ReleaseAddressesResponse) {
	response = &ReleaseAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ReleaseAddresses) 用于释放一个或多个[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * 该操作不可逆，释放后 EIP 关联的 IP 地址将不再属于您的名下。
// * 只有状态为 UNBIND 的 EIP 才能进行释放操作。
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
	if request == nil {
		request = NewReleaseAddressesRequest()
	}
	response = NewReleaseAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpcEndPointConnectionRequest() (request *EnableVpcEndPointConnectionRequest) {
	request = &EnableVpcEndPointConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableVpcEndPointConnection")
	return
}

func NewEnableVpcEndPointConnectionResponse() (response *EnableVpcEndPointConnectionResponse) {
	response = &EnableVpcEndPointConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 是否接受终端节点连接请求。
func (c *Client) EnableVpcEndPointConnection(request *EnableVpcEndPointConnectionRequest) (response *EnableVpcEndPointConnectionResponse, err error) {
	if request == nil {
		request = NewEnableVpcEndPointConnectionRequest()
	}
	response = NewEnableVpcEndPointConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkInterfaceLimitRequest() (request *DescribeNetworkInterfaceLimitRequest) {
	request = &DescribeNetworkInterfaceLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaceLimit")
	return
}

func NewDescribeNetworkInterfaceLimitResponse() (response *DescribeNetworkInterfaceLimitResponse) {
	response = &DescribeNetworkInterfaceLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNetworkInterfaceLimit）根据CVM实例ID或弹性网卡ID查询弹性网卡配额，返回该CVM实例或弹性网卡能绑定的弹性网卡配额，以及弹性网卡可以分配的IP配额
func (c *Client) DescribeNetworkInterfaceLimit(request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkInterfaceLimitRequest()
	}
	response = NewDescribeNetworkInterfaceLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficPackagesRequest() (request *DescribeTrafficPackagesRequest) {
	request = &DescribeTrafficPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTrafficPackages")
	return
}

func NewDescribeTrafficPackagesResponse() (response *DescribeTrafficPackagesResponse) {
	response = &DescribeTrafficPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询共享流量包详细信息，包括共享流量包唯一标识ID，名称，流量使用信息等
func (c *Client) DescribeTrafficPackages(request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficPackagesRequest()
	}
	response = NewDescribeTrafficPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDirectConnectGatewayCcnRoutesRequest() (request *DeleteDirectConnectGatewayCcnRoutesRequest) {
	request = &DeleteDirectConnectGatewayCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGatewayCcnRoutes")
	return
}

func NewDeleteDirectConnectGatewayCcnRoutesResponse() (response *DeleteDirectConnectGatewayCcnRoutesResponse) {
	response = &DeleteDirectConnectGatewayCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteDirectConnectGatewayCcnRoutes）用于删除专线网关的云联网路由（IDC网段）
func (c *Client) DeleteDirectConnectGatewayCcnRoutes(request *DeleteDirectConnectGatewayCcnRoutesRequest) (response *DeleteDirectConnectGatewayCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDeleteDirectConnectGatewayCcnRoutesRequest()
	}
	response = NewDeleteDirectConnectGatewayCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerGatewaysRequest() (request *DescribeCustomerGatewaysRequest) {
	request = &DescribeCustomerGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCustomerGateways")
	return
}

func NewDescribeCustomerGatewaysResponse() (response *DescribeCustomerGatewaysResponse) {
	response = &DescribeCustomerGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCustomerGateways）用于查询对端网关列表。
func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerGatewaysRequest()
	}
	response = NewDescribeCustomerGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpnConnectionAttributeRequest() (request *ModifyVpnConnectionAttributeRequest) {
	request = &ModifyVpnConnectionAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnConnectionAttribute")
	return
}

func NewModifyVpnConnectionAttributeResponse() (response *ModifyVpnConnectionAttributeResponse) {
	response = &ModifyVpnConnectionAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyVpnConnectionAttribute）用于修改VPN通道。
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
	if request == nil {
		request = NewModifyVpnConnectionAttributeRequest()
	}
	response = NewModifyVpnConnectionAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcPeeringConnectionExRequest() (request *ModifyVpcPeeringConnectionExRequest) {
	request = &ModifyVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcPeeringConnectionEx")
	return
}

func NewModifyVpcPeeringConnectionExResponse() (response *ModifyVpcPeeringConnectionExResponse) {
	response = &ModifyVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改跨地域对等连接属性
func (c *Client) ModifyVpcPeeringConnectionEx(request *ModifyVpcPeeringConnectionExRequest) (response *ModifyVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewModifyVpcPeeringConnectionExRequest()
	}
	response = NewModifyVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewResetTrafficMirrorTargetRequest() (request *ResetTrafficMirrorTargetRequest) {
	request = &ResetTrafficMirrorTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ResetTrafficMirrorTarget")
	return
}

func NewResetTrafficMirrorTargetResponse() (response *ResetTrafficMirrorTargetResponse) {
	response = &ResetTrafficMirrorTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetTrafficMirrorTarget）用于更新流量镜像实例的接收目的信息。
func (c *Client) ResetTrafficMirrorTarget(request *ResetTrafficMirrorTargetRequest) (response *ResetTrafficMirrorTargetResponse, err error) {
	if request == nil {
		request = NewResetTrafficMirrorTargetRequest()
	}
	response = NewResetTrafficMirrorTargetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubnetsRequest() (request *CreateSubnetsRequest) {
	request = &CreateSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnets")
	return
}

func NewCreateSubnetsResponse() (response *CreateSubnetsResponse) {
	response = &CreateSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateSubnets)用于批量创建子网。
// * 创建子网前必须创建好 VPC。
// * 子网创建成功后，子网网段不能修改。子网网段必须在VPC网段内，可以和VPC网段相同（VPC有且只有一个子网时），建议子网网段在VPC网段内，预留网段给其他子网使用。
// * 您可以创建的最小网段子网掩码为28（有16个IP地址），最大网段子网掩码为16（65,536个IP地址）。
// * 同一个VPC内，多个子网的网段不能重叠。
// * 子网创建后会自动关联到默认路由表。
// * 创建子网同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateSubnets(request *CreateSubnetsRequest) (response *CreateSubnetsResponse, err error) {
	if request == nil {
		request = NewCreateSubnetsRequest()
	}
	response = NewCreateSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewRejectVpcPeeringConnectionExRequest() (request *RejectVpcPeeringConnectionExRequest) {
	request = &RejectVpcPeeringConnectionExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "RejectVpcPeeringConnectionEx")
	return
}

func NewRejectVpcPeeringConnectionExResponse() (response *RejectVpcPeeringConnectionExResponse) {
	response = &RejectVpcPeeringConnectionExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驳回跨地域对等连接
func (c *Client) RejectVpcPeeringConnectionEx(request *RejectVpcPeeringConnectionExRequest) (response *RejectVpcPeeringConnectionExResponse, err error) {
	if request == nil {
		request = NewRejectVpcPeeringConnectionExRequest()
	}
	response = NewRejectVpcPeeringConnectionExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) {
	request = &CreateNatGatewayDestinationIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGatewayDestinationIpPortTranslationNatRule")
	return
}

func NewCreateNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) {
	response = &CreateNatGatewayDestinationIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateNatGatewayDestinationIpPortTranslationNatRule)用于创建NAT网关端口转发规则。
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRule(request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest()
	}
	response = NewCreateNatGatewayDestinationIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayVpnSSLClientRequest() (request *CreateGatewayVpnSSLClientRequest) {
	request = &CreateGatewayVpnSSLClientRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateGatewayVpnSSLClient")
	return
}

func NewCreateGatewayVpnSSLClientResponse() (response *CreateGatewayVpnSSLClientResponse) {
	response = &CreateGatewayVpnSSLClientResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateGatewayVpnSSLClient
func (c *Client) CreateGatewayVpnSSLClient(request *CreateGatewayVpnSSLClientRequest) (response *CreateGatewayVpnSSLClientResponse, err error) {
	if request == nil {
		request = NewCreateGatewayVpnSSLClientRequest()
	}
	response = NewCreateGatewayVpnSSLClientResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcEndPointServiceRequest() (request *CreateVpcEndPointServiceRequest) {
	request = &CreateVpcEndPointServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPointService")
	return
}

func NewCreateVpcEndPointServiceResponse() (response *CreateVpcEndPointServiceResponse) {
	response = &CreateVpcEndPointServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建终端节点服务。
func (c *Client) CreateVpcEndPointService(request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
	if request == nil {
		request = NewCreateVpcEndPointServiceRequest()
	}
	response = NewCreateVpcEndPointServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLocalSourceIpPortTranslationAclRulesRequest() (request *DescribeLocalSourceIpPortTranslationAclRulesRequest) {
	request = &DescribeLocalSourceIpPortTranslationAclRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalSourceIpPortTranslationAclRules")
	return
}

func NewDescribeLocalSourceIpPortTranslationAclRulesResponse() (response *DescribeLocalSourceIpPortTranslationAclRulesResponse) {
	response = &DescribeLocalSourceIpPortTranslationAclRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关本端源IP端口转换ACL规则
func (c *Client) DescribeLocalSourceIpPortTranslationAclRules(request *DescribeLocalSourceIpPortTranslationAclRulesRequest) (response *DescribeLocalSourceIpPortTranslationAclRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLocalSourceIpPortTranslationAclRulesRequest()
	}
	response = NewDescribeLocalSourceIpPortTranslationAclRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcs")
	return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcs）用于查询私有网络列表。
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcsRequest()
	}
	response = NewDescribeVpcsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcIdsRequest() (request *DescribeVpcIdsRequest) {
	request = &DescribeVpcIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcIds")
	return
}

func NewDescribeVpcIdsResponse() (response *DescribeVpcIdsResponse) {
	response = &DescribeVpcIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpcIds）用于查询VPC数字ID。
func (c *Client) DescribeVpcIds(request *DescribeVpcIdsRequest) (response *DescribeVpcIdsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcIdsRequest()
	}
	response = NewDescribeVpcIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAddressTemplateGroupRequest() (request *DeleteAddressTemplateGroupRequest) {
	request = &DeleteAddressTemplateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplateGroup")
	return
}

func NewDeleteAddressTemplateGroupResponse() (response *DeleteAddressTemplateGroupResponse) {
	response = &DeleteAddressTemplateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAddressTemplateGroup）用于删除IP地址模板集合
func (c *Client) DeleteAddressTemplateGroup(request *DeleteAddressTemplateGroupRequest) (response *DeleteAddressTemplateGroupResponse, err error) {
	if request == nil {
		request = NewDeleteAddressTemplateGroupRequest()
	}
	response = NewDeleteAddressTemplateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressTemplateGroupsRequest() (request *DescribeAddressTemplateGroupsRequest) {
	request = &DescribeAddressTemplateGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplateGroups")
	return
}

func NewDescribeAddressTemplateGroupsResponse() (response *DescribeAddressTemplateGroupsResponse) {
	response = &DescribeAddressTemplateGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAddressTemplateGroups）用于查询IP地址模板集合
func (c *Client) DescribeAddressTemplateGroups(request *DescribeAddressTemplateGroupsRequest) (response *DescribeAddressTemplateGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeAddressTemplateGroupsRequest()
	}
	response = NewDescribeAddressTemplateGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnConnectionsRequest() (request *DescribeVpnConnectionsRequest) {
	request = &DescribeVpnConnectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnConnections")
	return
}

func NewDescribeVpnConnectionsResponse() (response *DescribeVpnConnectionsResponse) {
	response = &DescribeVpnConnectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeVpnConnections）查询VPN通道列表。
func (c *Client) DescribeVpnConnections(request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
	if request == nil {
		request = NewDescribeVpnConnectionsRequest()
	}
	response = NewDescribeVpnConnectionsResponse()
	err = c.Send(request, response)
	return
}

func NewPublicIp6AddressesRequest() (request *PublicIp6AddressesRequest) {
	request = &PublicIp6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "PublicIp6Addresses")
	return
}

func NewPublicIp6AddressesResponse() (response *PublicIp6AddressesResponse) {
	response = &PublicIp6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于给指定的ipv6地址开通internet访问能力。
func (c *Client) PublicIp6Addresses(request *PublicIp6AddressesRequest) (response *PublicIp6AddressesResponse, err error) {
	if request == nil {
		request = NewPublicIp6AddressesRequest()
	}
	response = NewPublicIp6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDefaultSecurityGroupRequest() (request *CreateDefaultSecurityGroupRequest) {
	request = &CreateDefaultSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateDefaultSecurityGroup")
	return
}

func NewCreateDefaultSecurityGroupResponse() (response *CreateDefaultSecurityGroupResponse) {
	response = &CreateDefaultSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDefaultSecurityGroup）用于创建（如果项目下未存在默认安全组，则创建；已存在则获取。）默认安全组（SecurityGroup）。
// * 每个账户下每个地域的每个项目的<a href="https://cloud.tencent.com/document/product/213/12453">安全组数量限制</a>。
// * 新建的安全组的入站和出站规则默认都是全部拒绝，在创建后通常您需要再调用CreateSecurityGroupPolicies将安全组的规则设置为需要的规则。
// * 创建安全组同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateDefaultSecurityGroup(request *CreateDefaultSecurityGroupRequest) (response *CreateDefaultSecurityGroupResponse, err error) {
	if request == nil {
		request = NewCreateDefaultSecurityGroupRequest()
	}
	response = NewCreateDefaultSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFlowLogRequest() (request *DeleteFlowLogRequest) {
	request = &DeleteFlowLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteFlowLog")
	return
}

func NewDeleteFlowLogResponse() (response *DeleteFlowLogResponse) {
	response = &DeleteFlowLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteFlowLog）用于删除流日志
func (c *Client) DeleteFlowLog(request *DeleteFlowLogRequest) (response *DeleteFlowLogResponse, err error) {
	if request == nil {
		request = NewDeleteFlowLogRequest()
	}
	response = NewDeleteFlowLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
	request = &DescribeAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddresses")
	return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
	response = &DescribeAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAddresses) 用于查询一个或多个[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）的详细信息。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的 EIP。
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
	if request == nil {
		request = NewDescribeAddressesRequest()
	}
	response = NewDescribeAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptVpcPeeringConnectionRequest() (request *AcceptVpcPeeringConnectionRequest) {
	request = &AcceptVpcPeeringConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AcceptVpcPeeringConnection")
	return
}

func NewAcceptVpcPeeringConnectionResponse() (response *AcceptVpcPeeringConnectionResponse) {
	response = &AcceptVpcPeeringConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接受同地域对等连接
func (c *Client) AcceptVpcPeeringConnection(request *AcceptVpcPeeringConnectionRequest) (response *AcceptVpcPeeringConnectionResponse, err error) {
	if request == nil {
		request = NewAcceptVpcPeeringConnectionRequest()
	}
	response = NewAcceptVpcPeeringConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewPrivateIp6AddressesRequest() (request *PrivateIp6AddressesRequest) {
	request = &PrivateIp6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "PrivateIp6Addresses")
	return
}

func NewPrivateIp6AddressesResponse() (response *PrivateIp6AddressesResponse) {
	response = &PrivateIp6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于取消IPV6地址访问internet的能力
func (c *Client) PrivateIp6Addresses(request *PrivateIp6AddressesRequest) (response *PrivateIp6AddressesResponse, err error) {
	if request == nil {
		request = NewPrivateIp6AddressesRequest()
	}
	response = NewPrivateIp6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetDetectRequest() (request *DeleteNetDetectRequest) {
	request = &DeleteNetDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetDetect")
	return
}

func NewDeleteNetDetectResponse() (response *DeleteNetDetectResponse) {
	response = &DeleteNetDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DeleteNetDetect)用于删除网络探测实例。
func (c *Client) DeleteNetDetect(request *DeleteNetDetectRequest) (response *DeleteNetDetectResponse, err error) {
	if request == nil {
		request = NewDeleteNetDetectRequest()
	}
	response = NewDeleteNetDetectResponse()
	err = c.Send(request, response)
	return
}

func NewSetCcnRegionBandwidthLimitsRequest() (request *SetCcnRegionBandwidthLimitsRequest) {
	request = &SetCcnRegionBandwidthLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "SetCcnRegionBandwidthLimits")
	return
}

func NewSetCcnRegionBandwidthLimitsResponse() (response *SetCcnRegionBandwidthLimitsResponse) {
	response = &SetCcnRegionBandwidthLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SetCcnRegionBandwidthLimits）用于设置云联网（CCN）各地域出带宽上限，该接口只能设置已关联网络实例包含的地域的出带宽上限
func (c *Client) SetCcnRegionBandwidthLimits(request *SetCcnRegionBandwidthLimitsRequest) (response *SetCcnRegionBandwidthLimitsResponse, err error) {
	if request == nil {
		request = NewSetCcnRegionBandwidthLimitsRequest()
	}
	response = NewSetCcnRegionBandwidthLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressAssociationQuotaRequest() (request *DescribeAddressAssociationQuotaRequest) {
	request = &DescribeAddressAssociationQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressAssociationQuota")
	return
}

func NewDescribeAddressAssociationQuotaResponse() (response *DescribeAddressAssociationQuotaResponse) {
	response = &DescribeAddressAssociationQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性公网IP（EIP）绑定配额
func (c *Client) DescribeAddressAssociationQuota(request *DescribeAddressAssociationQuotaRequest) (response *DescribeAddressAssociationQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAddressAssociationQuotaRequest()
	}
	response = NewDescribeAddressAssociationQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupAssociationStatisticsRequest() (request *DescribeSecurityGroupAssociationStatisticsRequest) {
	request = &DescribeSecurityGroupAssociationStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupAssociationStatistics")
	return
}

func NewDescribeSecurityGroupAssociationStatisticsResponse() (response *DescribeSecurityGroupAssociationStatisticsResponse) {
	response = &DescribeSecurityGroupAssociationStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSecurityGroupAssociationStatistics）用于查询安全组关联的实例统计。
func (c *Client) DescribeSecurityGroupAssociationStatistics(request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupAssociationStatisticsRequest()
	}
	response = NewDescribeSecurityGroupAssociationStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCcnBandwidthRequest() (request *CreateCcnBandwidthRequest) {
	request = &CreateCcnBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateCcnBandwidth")
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

func NewDeleteLocalSourceIpPortTranslationAclRuleRequest() (request *DeleteLocalSourceIpPortTranslationAclRuleRequest) {
	request = &DeleteLocalSourceIpPortTranslationAclRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalSourceIpPortTranslationAclRule")
	return
}

func NewDeleteLocalSourceIpPortTranslationAclRuleResponse() (response *DeleteLocalSourceIpPortTranslationAclRuleResponse) {
	response = &DeleteLocalSourceIpPortTranslationAclRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线网关本端源IP端口转换ACL规则
func (c *Client) DeleteLocalSourceIpPortTranslationAclRule(request *DeleteLocalSourceIpPortTranslationAclRuleRequest) (response *DeleteLocalSourceIpPortTranslationAclRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLocalSourceIpPortTranslationAclRuleRequest()
	}
	response = NewDeleteLocalSourceIpPortTranslationAclRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressTemplateGroupInstancesRequest() (request *DescribeAddressTemplateGroupInstancesRequest) {
	request = &DescribeAddressTemplateGroupInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplateGroupInstances")
	return
}

func NewDescribeAddressTemplateGroupInstancesResponse() (response *DescribeAddressTemplateGroupInstancesResponse) {
	response = &DescribeAddressTemplateGroupInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAddressTemplateGroupInstances）用于查询参数模板IP地址组口关联的实例列表。本接口不会返回查询的结果，需要根据返回的RequestId调用DescribeVpcTaskResult接口获取结果。
func (c *Client) DescribeAddressTemplateGroupInstances(request *DescribeAddressTemplateGroupInstancesRequest) (response *DescribeAddressTemplateGroupInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeAddressTemplateGroupInstancesRequest()
	}
	response = NewDescribeAddressTemplateGroupInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceSecurityGroupPolicyRequest() (request *ReplaceSecurityGroupPolicyRequest) {
	request = &ReplaceSecurityGroupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ReplaceSecurityGroupPolicy")
	return
}

func NewReplaceSecurityGroupPolicyResponse() (response *ReplaceSecurityGroupPolicyResponse) {
	response = &ReplaceSecurityGroupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceSecurityGroupPolicy）用于替换单条安全组规则（SecurityGroupPolicy）。
// 单个请求中只能替换单个方向的一条规则, 必须要指定索引（PolicyIndex）。
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
	if request == nil {
		request = NewReplaceSecurityGroupPolicyRequest()
	}
	response = NewReplaceSecurityGroupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAndAttachNetworkInterfaceRequest() (request *CreateAndAttachNetworkInterfaceRequest) {
	request = &CreateAndAttachNetworkInterfaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateAndAttachNetworkInterface")
	return
}

func NewCreateAndAttachNetworkInterfaceResponse() (response *CreateAndAttachNetworkInterfaceResponse) {
	response = &CreateAndAttachNetworkInterfaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAndAttachNetworkInterface）用于创建弹性网卡并绑定云主机。
// * 创建弹性网卡时可以指定内网IP，并且可以指定一个主IP，指定的内网IP必须在弹性网卡所在子网内，而且不能被占用。
// * 创建弹性网卡时可以指定需要申请的内网IP数量，系统会随机生成内网IP地址。
// * 一个弹性网卡支持绑定的IP地址是有限制的，更多资源限制信息详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
// * 创建弹性网卡同时可以绑定已有安全组。
// * 创建弹性网卡同时可以绑定标签, 应答里的标签列表代表添加成功的标签。
func (c *Client) CreateAndAttachNetworkInterface(request *CreateAndAttachNetworkInterfaceRequest) (response *CreateAndAttachNetworkInterfaceResponse, err error) {
	if request == nil {
		request = NewCreateAndAttachNetworkInterfaceRequest()
	}
	response = NewCreateAndAttachNetworkInterfaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
	request = &DescribeCcnAttachedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnAttachedInstances")
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

func NewDescribeIpOnlineRequest() (request *DescribeIpOnlineRequest) {
	request = &DescribeIpOnlineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpOnline")
	return
}

func NewDescribeIpOnlineResponse() (response *DescribeIpOnlineResponse) {
	response = &DescribeIpOnlineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【废弃，采用DescribeIpLocation替代】查询IP地址信息，包括地理位置信息和网络信息
func (c *Client) DescribeIpOnline(request *DescribeIpOnlineRequest) (response *DescribeIpOnlineResponse, err error) {
	if request == nil {
		request = NewDescribeIpOnlineRequest()
	}
	response = NewDescribeIpOnlineResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateNatGatewayAddressRequest() (request *DisassociateNatGatewayAddressRequest) {
	request = &DisassociateNatGatewayAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNatGatewayAddress")
	return
}

func NewDisassociateNatGatewayAddressResponse() (response *DisassociateNatGatewayAddressResponse) {
	response = &DisassociateNatGatewayAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DisassociateNatGatewayAddress）用于NAT网关解绑弹性IP。
func (c *Client) DisassociateNatGatewayAddress(request *DisassociateNatGatewayAddressRequest) (response *DisassociateNatGatewayAddressResponse, err error) {
	if request == nil {
		request = NewDisassociateNatGatewayAddressRequest()
	}
	response = NewDisassociateNatGatewayAddressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLocalSourceIpPortTranslationNatRuleRequest() (request *CreateLocalSourceIpPortTranslationNatRuleRequest) {
	request = &CreateLocalSourceIpPortTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalSourceIpPortTranslationNatRule")
	return
}

func NewCreateLocalSourceIpPortTranslationNatRuleResponse() (response *CreateLocalSourceIpPortTranslationNatRuleResponse) {
	response = &CreateLocalSourceIpPortTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线网关本端源IP端口转换
func (c *Client) CreateLocalSourceIpPortTranslationNatRule(request *CreateLocalSourceIpPortTranslationNatRuleRequest) (response *CreateLocalSourceIpPortTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreateLocalSourceIpPortTranslationNatRuleRequest()
	}
	response = NewCreateLocalSourceIpPortTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTrafficPackagesRequest() (request *DeleteTrafficPackagesRequest) {
	request = &DeleteTrafficPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteTrafficPackages")
	return
}

func NewDeleteTrafficPackagesResponse() (response *DeleteTrafficPackagesResponse) {
	response = &DeleteTrafficPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除共享带宽包（仅非活动状态的流量包可删除）。
func (c *Client) DeleteTrafficPackages(request *DeleteTrafficPackagesRequest) (response *DeleteTrafficPackagesResponse, err error) {
	if request == nil {
		request = NewDeleteTrafficPackagesRequest()
	}
	response = NewDeleteTrafficPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecificTrafficPackageResourcesUsedStatisticsRequest() (request *DescribeSpecificTrafficPackageResourcesUsedStatisticsRequest) {
	request = &DescribeSpecificTrafficPackageResourcesUsedStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSpecificTrafficPackageResourcesUsedStatistics")
	return
}

func NewDescribeSpecificTrafficPackageResourcesUsedStatisticsResponse() (response *DescribeSpecificTrafficPackageResourcesUsedStatisticsResponse) {
	response = &DescribeSpecificTrafficPackageResourcesUsedStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeSpecificTrafficPackageResourcesUsedStatistics)用于查询指定共享流量包内的资源在某一时间区间内的抵扣量信息
func (c *Client) DescribeSpecificTrafficPackageResourcesUsedStatistics(request *DescribeSpecificTrafficPackageResourcesUsedStatisticsRequest) (response *DescribeSpecificTrafficPackageResourcesUsedStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeSpecificTrafficPackageResourcesUsedStatisticsRequest()
	}
	response = NewDescribeSpecificTrafficPackageResourcesUsedStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpcEndPointConnectRequest() (request *EnableVpcEndPointConnectRequest) {
	request = &EnableVpcEndPointConnectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EnableVpcEndPointConnect")
	return
}

func NewEnableVpcEndPointConnectResponse() (response *EnableVpcEndPointConnectResponse) {
	response = &EnableVpcEndPointConnectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 是否接受终端节点连接请求。
func (c *Client) EnableVpcEndPointConnect(request *EnableVpcEndPointConnectRequest) (response *EnableVpcEndPointConnectResponse, err error) {
	if request == nil {
		request = NewEnableVpcEndPointConnectRequest()
	}
	response = NewEnableVpcEndPointConnectResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBandwidthPackageRequest() (request *CheckBandwidthPackageRequest) {
	request = &CheckBandwidthPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CheckBandwidthPackage")
	return
}

func NewCheckBandwidthPackageResponse() (response *CheckBandwidthPackageResponse) {
	response = &CheckBandwidthPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查账户带宽包属性。
func (c *Client) CheckBandwidthPackage(request *CheckBandwidthPackageRequest) (response *CheckBandwidthPackageResponse, err error) {
	if request == nil {
		request = NewCheckBandwidthPackageRequest()
	}
	response = NewCheckBandwidthPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkInterfacesExtraRequest() (request *DescribeNetworkInterfacesExtraRequest) {
	request = &DescribeNetworkInterfacesExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfacesExtra")
	return
}

func NewDescribeNetworkInterfacesExtraResponse() (response *DescribeNetworkInterfacesExtraResponse) {
	response = &DescribeNetworkInterfacesExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNetworkInterfacesExtra）用于查询弹性网卡列表（IP维度），支持根据内网IP模糊查询。
func (c *Client) DescribeNetworkInterfacesExtra(request *DescribeNetworkInterfacesExtraRequest) (response *DescribeNetworkInterfacesExtraResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkInterfacesExtraRequest()
	}
	response = NewDescribeNetworkInterfacesExtraResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
	request = &DescribeSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnets")
	return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
	response = &DescribeSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSubnets）用于查询子网列表。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetsRequest()
	}
	response = NewDescribeSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBandwidthPackageBandwidthRequest() (request *ModifyBandwidthPackageBandwidthRequest) {
	request = &ModifyBandwidthPackageBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyBandwidthPackageBandwidth")
	return
}

func NewModifyBandwidthPackageBandwidthResponse() (response *ModifyBandwidthPackageBandwidthResponse) {
	response = &ModifyBandwidthPackageBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于调整共享带宽包带宽
func (c *Client) ModifyBandwidthPackageBandwidth(request *ModifyBandwidthPackageBandwidthRequest) (response *ModifyBandwidthPackageBandwidthResponse, err error) {
	if request == nil {
		request = NewModifyBandwidthPackageBandwidthRequest()
	}
	response = NewModifyBandwidthPackageBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupPoliciesRequest() (request *DescribeSecurityGroupPoliciesRequest) {
	request = &DescribeSecurityGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupPolicies")
	return
}

func NewDescribeSecurityGroupPoliciesResponse() (response *DescribeSecurityGroupPoliciesResponse) {
	response = &DescribeSecurityGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSecurityGroupPolicies）用于查询安全组规则。
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupPoliciesRequest()
	}
	response = NewDescribeSecurityGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewAssignIpv6AddressesRequest() (request *AssignIpv6AddressesRequest) {
	request = &AssignIpv6AddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6Addresses")
	return
}

func NewAssignIpv6AddressesResponse() (response *AssignIpv6AddressesResponse) {
	response = &AssignIpv6AddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AssignIpv6Addresses）用于弹性网卡申请`IPv6`地址。<br />
// 本接口是异步完成，如需查询异步任务执行结果，请使用本接口返回的`RequestId`轮询`QueryTask`接口。
// * 一个弹性网卡支持绑定的IP地址是有限制的，更多资源限制信息详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
// * 可以指定`IPv6`地址申请，地址类型不能为主`IP`，`IPv6`地址暂时只支持作为辅助`IP`。
// * 地址必须要在弹性网卡所在子网内，而且不能被占用。
// * 在弹性网卡上申请一个到多个辅助`IPv6`地址，接口会在弹性网卡所在子网段内返回指定数量的辅助`IPv6`地址。
func (c *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
	if request == nil {
		request = NewAssignIpv6AddressesRequest()
	}
	response = NewAssignIpv6AddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRoutesRequest() (request *DescribeCcnRoutesRequest) {
	request = &DescribeCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnRoutes")
	return
}

func NewDescribeCcnRoutesResponse() (response *DescribeCcnRoutesResponse) {
	response = &DescribeCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCcnRoutes）用于查询已加入云联网（CCN）的路由
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRoutesRequest()
	}
	response = NewDescribeCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
	request = &DescribeTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeTaskResult")
	return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
	response = &DescribeTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询EIP异步任务执行结果
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeTaskResultRequest()
	}
	response = NewDescribeTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpdateCcnBandwidthDealRequest() (request *GetUpdateCcnBandwidthDealRequest) {
	request = &GetUpdateCcnBandwidthDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "GetUpdateCcnBandwidthDeal")
	return
}

func NewGetUpdateCcnBandwidthDealResponse() (response *GetUpdateCcnBandwidthDealResponse) {
	response = &GetUpdateCcnBandwidthDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取变更云联网带宽的商品信息
func (c *Client) GetUpdateCcnBandwidthDeal(request *GetUpdateCcnBandwidthDealRequest) (response *GetUpdateCcnBandwidthDealResponse, err error) {
	if request == nil {
		request = NewGetUpdateCcnBandwidthDealRequest()
	}
	response = NewGetUpdateCcnBandwidthDealResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNatGatewaySourceIpTranslationNatRuleRequest() (request *CreateNatGatewaySourceIpTranslationNatRuleRequest) {
	request = &CreateNatGatewaySourceIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGatewaySourceIpTranslationNatRule")
	return
}

func NewCreateNatGatewaySourceIpTranslationNatRuleResponse() (response *CreateNatGatewaySourceIpTranslationNatRuleResponse) {
	response = &CreateNatGatewaySourceIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateNatGatewaySourceIpTranslationNatRule)用于创建NAT网关SNAT规则
func (c *Client) CreateNatGatewaySourceIpTranslationNatRule(request *CreateNatGatewaySourceIpTranslationNatRuleRequest) (response *CreateNatGatewaySourceIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewCreateNatGatewaySourceIpTranslationNatRuleRequest()
	}
	response = NewCreateNatGatewaySourceIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatGatewaySourceIpTranslationNatRuleRequest() (request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) {
	request = &ModifyNatGatewaySourceIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewaySourceIpTranslationNatRule")
	return
}

func NewModifyNatGatewaySourceIpTranslationNatRuleResponse() (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse) {
	response = &ModifyNatGatewaySourceIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyNatGatewaySourceIpTranslationNatRule）用于修改NAT网关SNAT转发规则。
func (c *Client) ModifyNatGatewaySourceIpTranslationNatRule(request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewModifyNatGatewaySourceIpTranslationNatRuleRequest()
	}
	response = NewModifyNatGatewaySourceIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewaySourceIpTranslationNatRulesRequest() (request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) {
	request = &DescribeNatGatewaySourceIpTranslationNatRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewaySourceIpTranslationNatRules")
	return
}

func NewDescribeNatGatewaySourceIpTranslationNatRulesResponse() (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse) {
	response = &DescribeNatGatewaySourceIpTranslationNatRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeNatGatewaySourceIpTranslationNatRules）用于查询NAT网关SNAT转发规则对象数组。
func (c *Client) DescribeNatGatewaySourceIpTranslationNatRules(request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewaySourceIpTranslationNatRulesRequest()
	}
	response = NewDescribeNatGatewaySourceIpTranslationNatRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatGatewaySourceIpTranslationNatRuleRequest() (request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) {
	request = &DeleteNatGatewaySourceIpTranslationNatRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGatewaySourceIpTranslationNatRule")
	return
}

func NewDeleteNatGatewaySourceIpTranslationNatRuleResponse() (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse) {
	response = &DeleteNatGatewaySourceIpTranslationNatRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteNatGatewaySourceIpTranslationNatRule）用于删除NAT网关端口SNAT转发规则。
func (c *Client) DeleteNatGatewaySourceIpTranslationNatRule(request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse, err error) {
	if request == nil {
		request = NewDeleteNatGatewaySourceIpTranslationNatRuleRequest()
	}
	response = NewDeleteNatGatewaySourceIpTranslationNatRuleResponse()
	err = c.Send(request, response)
	return
}

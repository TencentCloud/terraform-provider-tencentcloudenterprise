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

func NewDescribeZoneHostForSellStatusRequest() (request *DescribeZoneHostForSellStatusRequest) {
	request = &DescribeZoneHostForSellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneHostForSellStatus")
	return
}

func NewDescribeZoneHostForSellStatusResponse() (response *DescribeZoneHostForSellStatusResponse) {
	response = &DescribeZoneHostForSellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取专用宿主机的大区售罄情况
func (c *Client) DescribeZoneHostForSellStatus(request *DescribeZoneHostForSellStatusRequest) (response *DescribeZoneHostForSellStatusResponse, err error) {
	if request == nil {
		request = NewDescribeZoneHostForSellStatusRequest()
	}
	response = NewDescribeZoneHostForSellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewResumeInstancesRequest() (request *ResumeInstancesRequest) {
	request = &ResumeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResumeInstances")
	return
}

func NewResumeInstancesResponse() (response *ResumeInstancesResponse) {
	response = &ResumeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResumeInstances) 用于恢复一个或多个实例。
func (c *Client) ResumeInstances(request *ResumeInstancesRequest) (response *ResumeInstancesResponse, err error) {
	if request == nil {
		request = NewResumeInstancesRequest()
	}
	response = NewResumeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
	request = &ResetInstancesPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesPassword")
	return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
	response = &ResetInstancesPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstancesPassword) 用于将实例操作系统的密码重置为用户指定的密码。
//
// * 只修改管理员帐号的密码。实例的操作系统不同，管理员帐号也会不一样(`Windows`为`Administrator`，`Ubuntu`为`ubuntu`，其它系统为`root`)。
// * 重置处于运行中状态的实例，需要显式指定强制关机参数`ForceStop`。如果没有显式指定强制关机参数，则只有处于关机状态的实例才允许执行重置密码操作。
// * 支持批量操作。将多个实例操作系统的密码重置为相同的密码。每次请求批量实例的上限为100。
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
	if request == nil {
		request = NewResetInstancesPasswordRequest()
	}
	response = NewResetInstancesPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterRunInstancesRequest() (request *SwitchParameterRunInstancesRequest) {
	request = &SwitchParameterRunInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterRunInstances")
	return
}

func NewSwitchParameterRunInstancesResponse() (response *SwitchParameterRunInstancesResponse) {
	response = &SwitchParameterRunInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterRunInstances) 用于创建一个或多个指定配置的实例参数转换。
//
// * 实例创建成功后将自动开机启动，[实例状态](../数据结构#instancestatus)变为“运行中”。
// * 预付费实例的购买会预先扣除本次实例购买所需金额，按小时后付费实例购买会预先冻结本次实例购买一小时内所需金额，在调用本接口前请确保账户余额充足。
// * 本接口为异步接口，当创建请求下发成功后会返回一个实例`ID`列表，此时实例的创建并立即未完成。在此期间实例的状态将会处于“准备中”，可以通过调用 [DescribeInstancesStatus](DescribeInstancesStatus) 接口查询对应实例的状态，来判断生产有没有最终成功。如果实例的状态由“准备中”变为“运行中”，则为创建成功。
func (c *Client) SwitchParameterRunInstances(request *SwitchParameterRunInstancesRequest) (response *SwitchParameterRunInstancesResponse, err error) {
	if request == nil {
		request = NewSwitchParameterRunInstancesRequest()
	}
	response = NewSwitchParameterRunInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
	request = &DisassociateSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DisassociateSecurityGroups")
	return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
	response = &DisassociateSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DisassociateSecurityGroups) 用于解绑实例的指定安全组。
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDisassociateSecurityGroupsRequest()
	}
	response = NewDisassociateSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeImages) 用于查看镜像列表。
//
// * 可以通过指定镜像ID来查询指定镜像的详细信息，或通过设定过滤器来查询满足过滤条件的镜像的详细信息。( 该接口不支持公有镜像id的查询 )
// * 指定偏移(Offset)和限制(Limit)来选择结果中的一部分，默认返回满足条件的前20个镜像信息。
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	response = NewDescribeImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRecentFailedOperationRequest() (request *DescribeInstancesRecentFailedOperationRequest) {
	request = &DescribeInstancesRecentFailedOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesRecentFailedOperation")
	return
}

func NewDescribeInstancesRecentFailedOperationResponse() (response *DescribeInstancesRecentFailedOperationResponse) {
	response = &DescribeInstancesRecentFailedOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeInstancesRecentFailedOperation) 获取实例最近失败操作记录。
func (c *Client) DescribeInstancesRecentFailedOperation(request *DescribeInstancesRecentFailedOperationRequest) (response *DescribeInstancesRecentFailedOperationResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRecentFailedOperationRequest()
	}
	response = NewDescribeInstancesRecentFailedOperationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternetChargeTypeConfigsRequest() (request *DescribeInternetChargeTypeConfigsRequest) {
	request = &DescribeInternetChargeTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInternetChargeTypeConfigs")
	return
}

func NewDescribeInternetChargeTypeConfigsResponse() (response *DescribeInternetChargeTypeConfigsResponse) {
	response = &DescribeInternetChargeTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络计费的类型
func (c *Client) DescribeInternetChargeTypeConfigs(request *DescribeInternetChargeTypeConfigsRequest) (response *DescribeInternetChargeTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInternetChargeTypeConfigsRequest()
	}
	response = NewDescribeInternetChargeTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserGlobalConfigsRequest() (request *ModifyUserGlobalConfigsRequest) {
	request = &ModifyUserGlobalConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyUserGlobalConfigs")
	return
}

func NewModifyUserGlobalConfigsResponse() (response *ModifyUserGlobalConfigsResponse) {
	response = &ModifyUserGlobalConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户全局配置
func (c *Client) ModifyUserGlobalConfigs(request *ModifyUserGlobalConfigsRequest) (response *ModifyUserGlobalConfigsResponse, err error) {
	if request == nil {
		request = NewModifyUserGlobalConfigsRequest()
	}
	response = NewModifyUserGlobalConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterAllocateHostsRequest() (request *SwitchParameterAllocateHostsRequest) {
	request = &SwitchParameterAllocateHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterAllocateHosts")
	return
}

func NewSwitchParameterAllocateHostsResponse() (response *SwitchParameterAllocateHostsResponse) {
	response = &SwitchParameterAllocateHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口自2022年4月14日起停止维护。
// 创建CDH实例参数转换（当HostChargeType为PREPAID时，必须指定HostChargePrepaid参数）
func (c *Client) SwitchParameterAllocateHosts(request *SwitchParameterAllocateHostsRequest) (response *SwitchParameterAllocateHostsResponse, err error) {
	if request == nil {
		request = NewSwitchParameterAllocateHostsRequest()
	}
	response = NewSwitchParameterAllocateHostsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDisasterRecoverGroupRequest() (request *ModifyDisasterRecoverGroupRequest) {
	request = &ModifyDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyDisasterRecoverGroup")
	return
}

func NewModifyDisasterRecoverGroupResponse() (response *ModifyDisasterRecoverGroupResponse) {
	response = &ModifyDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改分散置放群组信息
func (c *Client) ModifyDisasterRecoverGroup(request *ModifyDisasterRecoverGroupRequest) (response *ModifyDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewModifyDisasterRecoverGroupRequest()
	}
	response = NewModifyDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewExportImageRequest() (request *ExportImageRequest) {
	request = &ExportImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ExportImage")
	return
}

func NewExportImageResponse() (response *ExportImageResponse) {
	response = &ExportImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供一个COS/CSP Bucket，并授权，我们将会导出这个镜像到指定的 Bucket 中。
func (c *Client) ExportImage(request *ExportImageRequest) (response *ExportImageResponse, err error) {
	if request == nil {
		request = NewExportImageRequest()
	}
	response = NewExportImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressesBandwidthRequest() (request *ModifyAddressesBandwidthRequest) {
	request = &ModifyAddressesBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyAddressesBandwidth")
	return
}

func NewModifyAddressesBandwidthResponse() (response *ModifyAddressesBandwidthResponse) {
	response = &ModifyAddressesBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
	if request == nil {
		request = NewModifyAddressesBandwidthRequest()
	}
	response = NewModifyAddressesBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterRenewInstancesRequest() (request *SwitchParameterRenewInstancesRequest) {
	request = &SwitchParameterRenewInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterRenewInstances")
	return
}

func NewSwitchParameterRenewInstancesResponse() (response *SwitchParameterRenewInstancesResponse) {
	response = &SwitchParameterRenewInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterRenewInstances) 用于续费包年包月实例参数转换。
//
// * 只支持操作包年包月实例。
func (c *Client) SwitchParameterRenewInstances(request *SwitchParameterRenewInstancesRequest) (response *SwitchParameterRenewInstancesResponse, err error) {
	if request == nil {
		request = NewSwitchParameterRenewInstancesRequest()
	}
	response = NewSwitchParameterRenewInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncImagesRequest() (request *SyncImagesRequest) {
	request = &SyncImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SyncImages")
	return
}

func NewSyncImagesResponse() (response *SyncImagesResponse) {
	response = &SyncImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SyncImages）用于将自定义镜像同步到其它地区。
//
// * 该接口每次调用只支持同步一个镜像。
// * 该接口支持多个同步地域。
// * 单个帐号在每个地域最多支持存在10个自定义镜像。
func (c *Client) SyncImages(request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
	if request == nil {
		request = NewSyncImagesRequest()
	}
	response = NewSyncImagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesActionTimerRequest() (request *ModifyInstancesActionTimerRequest) {
	request = &ModifyInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesActionTimer")
	return
}

func NewModifyInstancesActionTimerResponse() (response *ModifyInstancesActionTimerResponse) {
	response = &ModifyInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改定时任务信息
func (c *Client) ModifyInstancesActionTimer(request *ModifyInstancesActionTimerRequest) (response *ModifyInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewModifyInstancesActionTimerRequest()
	}
	response = NewModifyInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKeyPairsAttributeRequest() (request *DescribeKeyPairsAttributeRequest) {
	request = &DescribeKeyPairsAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeKeyPairsAttribute")
	return
}

func NewDescribeKeyPairsAttributeResponse() (response *DescribeKeyPairsAttributeResponse) {
	response = &DescribeKeyPairsAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内部接口，查询密钥ID对应的内部密钥ID。
func (c *Client) DescribeKeyPairsAttribute(request *DescribeKeyPairsAttributeRequest) (response *DescribeKeyPairsAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeKeyPairsAttributeRequest()
	}
	response = NewDescribeKeyPairsAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneHostConfigInfosRequest() (request *DescribeZoneHostConfigInfosRequest) {
	request = &DescribeZoneHostConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneHostConfigInfos")
	return
}

func NewDescribeZoneHostConfigInfosResponse() (response *DescribeZoneHostConfigInfosResponse) {
	response = &DescribeZoneHostConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取专用宿主机的机型配置信息,以及售罄状态信息列表。
func (c *Client) DescribeZoneHostConfigInfos(request *DescribeZoneHostConfigInfosRequest) (response *DescribeZoneHostConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneHostConfigInfosRequest()
	}
	response = NewDescribeZoneHostConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostsAttributeRequest() (request *ModifyHostsAttributeRequest) {
	request = &ModifyHostsAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyHostsAttribute")
	return
}

func NewModifyHostsAttributeResponse() (response *ModifyHostsAttributeResponse) {
	response = &ModifyHostsAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyHostsAttribute）用于修改CDH实例的属性，如实例名称和续费标记等。参数HostName和RenewFlag必须设置其中一个，但不能同时设置。
func (c *Client) ModifyHostsAttribute(request *ModifyHostsAttributeRequest) (response *ModifyHostsAttributeResponse, err error) {
	if request == nil {
		request = NewModifyHostsAttributeRequest()
	}
	response = NewModifyHostsAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLaunchTemplateRequest() (request *ModifyLaunchTemplateRequest) {
	request = &ModifyLaunchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyLaunchTemplate")
	return
}

func NewModifyLaunchTemplateResponse() (response *ModifyLaunchTemplateResponse) {
	response = &ModifyLaunchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyLaunchTemplate）用于修改实例启动模板。
//
// 实例启动模板是一种配置数据并可用于创建实例，其内容包含创建实例所需的配置，比如实例类型，数据盘和系统盘的类型和大小，以及安全组等信息。
func (c *Client) ModifyLaunchTemplate(request *ModifyLaunchTemplateRequest) (response *ModifyLaunchTemplateResponse, err error) {
	if request == nil {
		request = NewModifyLaunchTemplateRequest()
	}
	response = NewModifyLaunchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyInstanceInternetChargeTypeRequest() (request *InquiryPriceModifyInstanceInternetChargeTypeRequest) {
	request = &InquiryPriceModifyInstanceInternetChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceModifyInstanceInternetChargeType")
	return
}

func NewInquiryPriceModifyInstanceInternetChargeTypeResponse() (response *InquiryPriceModifyInstanceInternetChargeTypeResponse) {
	response = &InquiryPriceModifyInstanceInternetChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (InquiryPriceModifyInstanceInternetChargeType) 用于切换实例公网网络的计费模式询价。
//
// * 只支持从 `BANDWIDTH_PREPAID` 计费模式切换为`TRAFFIC_POSTPAID_BY_HOUR`计费模式，或者从 `TRAFFIC_POSTPAID_BY_HOUR` 计费模式切换为`BANDWIDTH_PREPAID`计费模式。
func (c *Client) InquiryPriceModifyInstanceInternetChargeType(request *InquiryPriceModifyInstanceInternetChargeTypeRequest) (response *InquiryPriceModifyInstanceInternetChargeTypeResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyInstanceInternetChargeTypeRequest()
	}
	response = NewInquiryPriceModifyInstanceInternetChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesChargeTypeRequest() (request *ModifyInstancesChargeTypeRequest) {
	request = &ModifyInstancesChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesChargeType")
	return
}

func NewModifyInstancesChargeTypeResponse() (response *ModifyInstancesChargeTypeResponse) {
	response = &ModifyInstancesChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesChargeType) 用于切换实例的计费模式。
//
// * 只支持从 `POSTPAID_BY_HOUR` 计费模式切换为`PREPAID`计费模式。
// * 关机不收费的实例、`BC1`和`BS1`机型族的实例、设置定时销毁的实例不支持该操作。
func (c *Client) ModifyInstancesChargeType(request *ModifyInstancesChargeTypeRequest) (response *ModifyInstancesChargeTypeResponse, err error) {
	if request == nil {
		request = NewModifyInstancesChargeTypeRequest()
	}
	response = NewModifyInstancesChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewImportSnapshotRequest() (request *ImportSnapshotRequest) {
	request = &ImportSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportSnapshot")
	return
}

func NewImportSnapshotResponse() (response *ImportSnapshotResponse) {
	response = &ImportSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入数据盘镜像，并制作为CBS快照。制作的快照可用于创建CBS数据盘。
func (c *Client) ImportSnapshot(request *ImportSnapshotRequest) (response *ImportSnapshotResponse, err error) {
	if request == nil {
		request = NewImportSnapshotRequest()
	}
	response = NewImportSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
	request = &DisassociateAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DisassociateAddress")
	return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
	response = &DisassociateAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DisassociateAddress) 用于解绑[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
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

func NewReturnNormalAddressesRequest() (request *ReturnNormalAddressesRequest) {
	request = &ReturnNormalAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ReturnNormalAddresses")
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

func NewAuditMarketImageRequest() (request *AuditMarketImageRequest) {
	request = &AuditMarketImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AuditMarketImage")
	return
}

func NewAuditMarketImageResponse() (response *AuditMarketImageResponse) {
	response = &AuditMarketImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// None
func (c *Client) AuditMarketImage(request *AuditMarketImageRequest) (response *AuditMarketImageResponse, err error) {
	if request == nil {
		request = NewAuditMarketImageRequest()
	}
	response = NewAuditMarketImageResponse()
	err = c.Send(request, response)
	return
}

func NewLiveMigrateInstanceRequest() (request *LiveMigrateInstanceRequest) {
	request = &LiveMigrateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "LiveMigrateInstance")
	return
}

func NewLiveMigrateInstanceResponse() (response *LiveMigrateInstanceResponse) {
	response = &LiveMigrateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在线迁移实例。用于将一台运行中的实例导入到腾讯云。该接口不可由用户直接调用
func (c *Client) LiveMigrateInstance(request *LiveMigrateInstanceRequest) (response *LiveMigrateInstanceResponse, err error) {
	if request == nil {
		request = NewLiveMigrateInstanceRequest()
	}
	response = NewLiveMigrateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAddressAttributeRequest() (request *ModifyAddressAttributeRequest) {
	request = &ModifyAddressAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyAddressAttribute")
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

func NewModifyInstanceInternetChargeTypeRequest() (request *ModifyInstanceInternetChargeTypeRequest) {
	request = &ModifyInstanceInternetChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstanceInternetChargeType")
	return
}

func NewModifyInstanceInternetChargeTypeResponse() (response *ModifyInstanceInternetChargeTypeResponse) {
	response = &ModifyInstanceInternetChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstanceInternetChargeType) 用于切换实例公网网络的计费模式。
//
// * 只支持从 `BANDWIDTH_PREPAID` 计费模式切换为`TRAFFIC_POSTPAID_BY_HOUR`计费模式，或者从 `TRAFFIC_POSTPAID_BY_HOUR` 计费模式切换为`BANDWIDTH_PREPAID`计费模式。
// * 切换实例公网网络的计费模式有操作次数限制，可通过 [`DescribeInstancesOperationLimit`](DescribeInstancesOperationLimit) 接口查询剩余操作次数。
func (c *Client) ModifyInstanceInternetChargeType(request *ModifyInstanceInternetChargeTypeRequest) (response *ModifyInstanceInternetChargeTypeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceInternetChargeTypeRequest()
	}
	response = NewModifyInstanceInternetChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
	request = &ImportKeyPairRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportKeyPair")
	return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
	response = &ImportKeyPairResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ImportKeyPair) 用于导入密钥对。
//
// * 本接口的功能是将密钥对导入到用户账户，并不会自动绑定到实例。如需绑定可以使用[AssociateInstancesKeyPairs](AssociateInstancesKeyPairs)接口。
// * 需指定密钥对名称以及该密钥对的公钥文本。
// * 如果用户只有私钥，可以通过 `SSL` 工具将私钥转换成公钥后再导入。
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
	if request == nil {
		request = NewImportKeyPairRequest()
	}
	response = NewImportKeyPairResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewInstancesRequest() (request *InquiryPriceRenewInstancesRequest) {
	request = &InquiryPriceRenewInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRenewInstances")
	return
}

func NewInquiryPriceRenewInstancesResponse() (response *InquiryPriceRenewInstancesResponse) {
	response = &InquiryPriceRenewInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (InquiryPriceRenewInstances) 用于续费包年包月实例询价。
//
// * 只支持查询包年包月实例的续费价格。
func (c *Client) InquiryPriceRenewInstances(request *InquiryPriceRenewInstancesRequest) (response *InquiryPriceRenewInstancesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewInstancesRequest()
	}
	response = NewInquiryPriceRenewInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesDisasterRecoverGroupRequest() (request *ModifyInstancesDisasterRecoverGroupRequest) {
	request = &ModifyInstancesDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesDisasterRecoverGroup")
	return
}

func NewModifyInstancesDisasterRecoverGroupResponse() (response *ModifyInstancesDisasterRecoverGroupResponse) {
	response = &ModifyInstancesDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesDisasterRecoverGroup) 用于调整实例所在置放群组。
// * 目前只支持基础网络或私有网络实例，同时系统盘和数据盘类型均为CLOUD_BASIC、CLOUD_PREMIUM、CLOUD_SSD类型的实例使用该接口进行置放群组修改。
func (c *Client) ModifyInstancesDisasterRecoverGroup(request *ModifyInstancesDisasterRecoverGroupRequest) (response *ModifyInstancesDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewModifyInstancesDisasterRecoverGroupRequest()
	}
	response = NewModifyInstancesDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMigrateTaskStatusRequest() (request *DescribeMigrateTaskStatusRequest) {
	request = &DescribeMigrateTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeMigrateTaskStatus")
	return
}

func NewDescribeMigrateTaskStatusResponse() (response *DescribeMigrateTaskStatusResponse) {
	response = &DescribeMigrateTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据服务迁移(离线实例迁移，离线数据迁移)接口返回的JobId查询任务状态及进度。
func (c *Client) DescribeMigrateTaskStatus(request *DescribeMigrateTaskStatusRequest) (response *DescribeMigrateTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeMigrateTaskStatusRequest()
	}
	response = NewDescribeMigrateTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRenewInstancesRequest() (request *RenewInstancesRequest) {
	request = &RenewInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RenewInstances")
	return
}

func NewRenewInstancesResponse() (response *RenewInstancesResponse) {
	response = &RenewInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RenewInstances) 用于续费实例。
func (c *Client) RenewInstances(request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
	if request == nil {
		request = NewRenewInstancesRequest()
	}
	response = NewRenewInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
	request = &AssociateAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AssociateAddress")
	return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
	response = &AssociateAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AssociateAddress) 用于将[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）绑定到实例或弹性网卡的指定内网 IP 上。
// * 将 EIP 绑定到实例上，其本质是将 EIP 绑定到实例上主网卡的主内网 IP 上。
// * 将 EIP 绑定到主网卡的主内网IP上，绑定过程会把其上绑定的普通公网 IP 自动解绑并释放。
// * 如果指定网卡的内网 IP 已经绑定了 EIP，则必须先解绑该 EIP，才能再绑定新的。
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

func NewModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
	request = &ModifyImageSharePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageSharePermission")
	return
}

func NewModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
	response = &ModifyImageSharePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyImageSharePermission）用于修改镜像分享信息。
//
// * 分享镜像后，被分享账户可以通过该镜像创建实例。
// * 每个自定义镜像最多可共享给50个账户。
// * 分享镜像无法更改名称，描述，仅可用于创建实例。
// * 只支持分享到对方账户相同地域。
func (c *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
	if request == nil {
		request = NewModifyImageSharePermissionRequest()
	}
	response = NewModifyImageSharePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewRenewHostsRequest() (request *RenewHostsRequest) {
	request = &RenewHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RenewHosts")
	return
}

func NewRenewHostsResponse() (response *RenewHostsResponse) {
	response = &RenewHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RenewHosts) 用于续费包年包月CDH实例。
//
// * 只支持操作包年包月实例，否则操作会以特定错误码返回。
// * 续费时请确保账户余额充足。
func (c *Client) RenewHosts(request *RenewHostsRequest) (response *RenewHostsResponse, err error) {
	if request == nil {
		request = NewRenewHostsRequest()
	}
	response = NewRenewHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
	request = &DescribeKeyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeKeyPairs")
	return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
	response = &DescribeKeyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeKeyPairs) 用于查询密钥对信息。
//
// * 密钥对是通过一种算法生成的一对密钥，在生成的密钥对中，一个向外界公开，称为公钥；另一个用户自己保留，称为私钥。密钥对的公钥内容可以通过这个接口查询，但私钥内容系统不保留。
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
	if request == nil {
		request = NewDescribeKeyPairsRequest()
	}
	response = NewDescribeKeyPairsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterResizeInstanceDisksRequest() (request *SwitchParameterResizeInstanceDisksRequest) {
	request = &SwitchParameterResizeInstanceDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterResizeInstanceDisks")
	return
}

func NewSwitchParameterResizeInstanceDisksResponse() (response *SwitchParameterResizeInstanceDisksResponse) {
	response = &SwitchParameterResizeInstanceDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterResizeInstanceDisks) 用于扩容实例的数据盘参数转换。
//
// * 目前只支持扩容随实例购买的数据盘，且[系统盘类型](../数据结构#datadisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。
// * 目前不支持[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)实例使用该接口扩容数据盘。
// * 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
// * 目前只支持扩容一块数据盘。
func (c *Client) SwitchParameterResizeInstanceDisks(request *SwitchParameterResizeInstanceDisksRequest) (response *SwitchParameterResizeInstanceDisksResponse, err error) {
	if request == nil {
		request = NewSwitchParameterResizeInstanceDisksRequest()
	}
	response = NewSwitchParameterResizeInstanceDisksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDisasterRecoverGroupRequest() (request *CreateDisasterRecoverGroupRequest) {
	request = &CreateDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateDisasterRecoverGroup")
	return
}

func NewCreateDisasterRecoverGroupResponse() (response *CreateDisasterRecoverGroupResponse) {
	response = &CreateDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建分散置放群组，该功能对资源挑战较大，请谨慎使用。创建好的分散置放群组，可在创建实例的时指定。
func (c *Client) CreateDisasterRecoverGroup(request *CreateDisasterRecoverGroupRequest) (response *CreateDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewCreateDisasterRecoverGroupRequest()
	}
	response = NewCreateDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterModifyInstanceInternetChargeTypeRequest() (request *SwitchParameterModifyInstanceInternetChargeTypeRequest) {
	request = &SwitchParameterModifyInstanceInternetChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterModifyInstanceInternetChargeType")
	return
}

func NewSwitchParameterModifyInstanceInternetChargeTypeResponse() (response *SwitchParameterModifyInstanceInternetChargeTypeResponse) {
	response = &SwitchParameterModifyInstanceInternetChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterModifyInstanceInternetChargeType) 用于切换实例公网网络的计费模式参数转换。
//
// * 只支持从 `BANDWIDTH_PREPAID` 计费模式切换为`TRAFFIC_POSTPAID_BY_HOUR`计费模式，或者从 `TRAFFIC_POSTPAID_BY_HOUR` 计费模式切换为`BANDWIDTH_PREPAID`计费模式。
// * 切换实例公网网络的计费模式有操作次数限制，可通过 [`DescribeInstancesOperationLimit`](DescribeInstancesOperationLimit) 接口查询剩余操作次数。
func (c *Client) SwitchParameterModifyInstanceInternetChargeType(request *SwitchParameterModifyInstanceInternetChargeTypeRequest) (response *SwitchParameterModifyInstanceInternetChargeTypeResponse, err error) {
	if request == nil {
		request = NewSwitchParameterModifyInstanceInternetChargeTypeRequest()
	}
	response = NewSwitchParameterModifyInstanceInternetChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImportSnapshotTaskRequest() (request *DescribeImportSnapshotTaskRequest) {
	request = &DescribeImportSnapshotTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImportSnapshotTask")
	return
}

func NewDescribeImportSnapshotTaskResponse() (response *DescribeImportSnapshotTaskResponse) {
	response = &DescribeImportSnapshotTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeImportSnapshotTask(request *DescribeImportSnapshotTaskRequest) (response *DescribeImportSnapshotTaskResponse, err error) {
	if request == nil {
		request = NewDescribeImportSnapshotTaskRequest()
	}
	response = NewDescribeImportSnapshotTaskResponse()
	err = c.Send(request, response)
	return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RunInstances")
	return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
//
// * 实例创建成功后将自动开机启动，[实例状态](../数据结构#instancestatus)变为“运行中”。
// * 预付费实例的购买会预先扣除本次实例购买所需金额，按小时后付费实例购买会预先冻结本次实例购买一小时内所需金额，在调用本接口前请确保账户余额充足。
// * 本接口为异步接口，当创建请求下发成功后会返回一个实例`ID`列表，此时实例的创建并立即未完成。在此期间实例的状态将会处于“准备中”，可以通过调用 [DescribeInstancesStatus](DescribeInstancesStatus) 接口查询对应实例的状态，来判断生产有没有最终成功。如果实例的状态由“准备中”变为“运行中”，则为创建成功。
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	response = NewRunInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
	request = &DeleteKeyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteKeyPairs")
	return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
	response = &DeleteKeyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteKeyPairs) 用于删除已在腾讯云托管的密钥对。
//
// * 可以同时删除多个密钥对。
// * 不能删除已被实例或镜像引用的密钥对，所以需要独立判断是否所有密钥对都被成功删除。
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
	if request == nil {
		request = NewDeleteKeyPairsRequest()
	}
	response = NewDeleteKeyPairsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiagnosticReportsRequest() (request *DescribeDiagnosticReportsRequest) {
	request = &DescribeDiagnosticReportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeDiagnosticReports")
	return
}

func NewDescribeDiagnosticReportsResponse() (response *DescribeDiagnosticReportsResponse) {
	response = &DescribeDiagnosticReportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例检测报告详细信息。
func (c *Client) DescribeDiagnosticReports(request *DescribeDiagnosticReportsRequest) (response *DescribeDiagnosticReportsResponse, err error) {
	if request == nil {
		request = NewDescribeDiagnosticReportsRequest()
	}
	response = NewDescribeDiagnosticReportsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateInstanceVpcConfigRequest() (request *UpdateInstanceVpcConfigRequest) {
	request = &UpdateInstanceVpcConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "UpdateInstanceVpcConfig")
	return
}

func NewUpdateInstanceVpcConfigResponse() (response *UpdateInstanceVpcConfigResponse) {
	response = &UpdateInstanceVpcConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(UpdateInstanceVpcConfig)用于修改实例vpc属性，如私有网络ip。
// * 此操作默认会关闭实例，完成后再启动。
// * 不支持跨VpcId操作。
func (c *Client) UpdateInstanceVpcConfig(request *UpdateInstanceVpcConfigRequest) (response *UpdateInstanceVpcConfigResponse, err error) {
	if request == nil {
		request = NewUpdateInstanceVpcConfigRequest()
	}
	response = NewUpdateInstanceVpcConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
	request = &AllocateAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AllocateAddresses")
	return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
	response = &AllocateAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (AllocateAddresses) 用于申请一个或多个[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * EIP 是专为动态云计算设计的静态 IP 地址。借助 EIP，您可以快速将 EIP 重新映射到您的另一个实例上，从而屏蔽实例故障。
// * 您的 EIP 与腾讯云账户相关联，而不是与某个实例相关联。在您选择显式释放该地址，或欠费超过七天之前，它会一直与您的腾讯云账户保持关联。
// * 平台对用户每地域能申请的 EIP 最大配额有所限制，可参见 [EIP 产品简介](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)，上述配额可通过 [DescribeAddressQuota](DescribeAddressQuota) 接口获取。
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
	if request == nil {
		request = NewAllocateAddressesRequest()
	}
	response = NewAllocateAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceAllocateHostsRequest() (request *InquiryPriceAllocateHostsRequest) {
	request = &InquiryPriceAllocateHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceAllocateHosts")
	return
}

func NewInquiryPriceAllocateHostsResponse() (response *InquiryPriceAllocateHostsResponse) {
	response = &InquiryPriceAllocateHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CDH实例询价（当HostChargeType为PREPAID时，必须指定HostChargePrepaid参数）
// 3100及以后版本只支持后付费，不支持预付费.传预付费参数会报错。
func (c *Client) InquiryPriceAllocateHosts(request *InquiryPriceAllocateHostsRequest) (response *InquiryPriceAllocateHostsResponse, err error) {
	if request == nil {
		request = NewInquiryPriceAllocateHostsRequest()
	}
	response = NewInquiryPriceAllocateHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstanceQuotaRequest() (request *DescribeUserInstanceQuotaRequest) {
	request = &DescribeUserInstanceQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserInstanceQuota")
	return
}

func NewDescribeUserInstanceQuotaResponse() (response *DescribeUserInstanceQuotaResponse) {
	response = &DescribeUserInstanceQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询用户可用区下配额。
func (c *Client) DescribeUserInstanceQuota(request *DescribeUserInstanceQuotaRequest) (response *DescribeUserInstanceQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstanceQuotaRequest()
	}
	response = NewDescribeUserInstanceQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesTypeRequest() (request *ResetInstancesTypeRequest) {
	request = &ResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesType")
	return
}

func NewResetInstancesTypeResponse() (response *ResetInstancesTypeResponse) {
	response = &ResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstancesType) 用于调整实例的机型。
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口进行机型调整。
// * 目前不支持[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)实例使用该接口调整机型。* 目前不支持跨机型系统来调整机型，即使用该接口时指定的`InstanceType`和实例原来的机型需要属于同一系列。* 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
func (c *Client) ResetInstancesType(request *ResetInstancesTypeRequest) (response *ResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewResetInstancesTypeRequest()
	}
	response = NewResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDiagnosticReportsRequest() (request *CreateDiagnosticReportsRequest) {
	request = &CreateDiagnosticReportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateDiagnosticReports")
	return
}

func NewCreateDiagnosticReportsResponse() (response *CreateDiagnosticReportsResponse) {
	response = &CreateDiagnosticReportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例检测报告
func (c *Client) CreateDiagnosticReports(request *CreateDiagnosticReportsRequest) (response *CreateDiagnosticReportsResponse, err error) {
	if request == nil {
		request = NewCreateDiagnosticReportsRequest()
	}
	response = NewCreateDiagnosticReportsResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesInternetMaxBandwidthRequest() (request *ResetInstancesInternetMaxBandwidthRequest) {
	request = &ResetInstancesInternetMaxBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesInternetMaxBandwidth")
	return
}

func NewResetInstancesInternetMaxBandwidthResponse() (response *ResetInstancesInternetMaxBandwidthResponse) {
	response = &ResetInstancesInternetMaxBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstancesInternetMaxBandwidth) 用于调整实例公网带宽上限。
//
// * 不同机型带宽上限范围不一致。
// * 对于 `BANDWIDTH_PREPAID` 计费方式的带宽，需要输入参数 `StartTime` 和 `EndTime` ，指定调整后的带宽的生效时间段。在这种场景下目前不支持调小带宽，会涉及扣费，请确保账户余额充足。
// * 对于 `TRAFFIC_POSTPAID_BY_HOUR` 、 `BANDWIDTH_POSTPAID_BY_HOUR` 和 `BANDWIDTH_PACKAGE` 计费方式的带宽，使用该接口调整带宽上限是实时生效的，可以在带宽允许的范围内调大或者调小带宽，不支持输入参数 `StartTime` 和 `EndTime` 。
// * 接口不支持调整 `BANDWIDTH_POSTPAID_BY_MONTH` 计费方式的带宽。
// * 接口不支持批量调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽。
// * 接口不支持批量调整混合计费方式的带宽。例如不支持同时调整 `TRAFFIC_POSTPAID_BY_HOUR` 和 `BANDWIDTH_PACKAGE` 计费方式的带宽。
func (c *Client) ResetInstancesInternetMaxBandwidth(request *ResetInstancesInternetMaxBandwidthRequest) (response *ResetInstancesInternetMaxBandwidthResponse, err error) {
	if request == nil {
		request = NewResetInstancesInternetMaxBandwidthRequest()
	}
	response = NewResetInstancesInternetMaxBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
	request = &DisassociateInstancesKeyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DisassociateInstancesKeyPairs")
	return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
	response = &DisassociateInstancesKeyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DisassociateInstancesKeyPairs) 用于解除实例的密钥绑定关系。
//
// * 只支持[STOPPED](../数据结构#instancestatus)状态的`Linux`操作系统的实例。
// * 解绑密钥后，实例可以通过原来设置的密码登录。
// * 如果原来没有设置密码，解绑后将无法使用 `SSH` 登录。可以调用 [ResetInstancesPassword](../实例相关接口/ResetInstancesPassword) 接口来设置登陆密码。
// * 支持批量操作。每次请求批量实例的上限为100。如果批量实例存在不允许操作的实例，操作会以特定错误码返回。
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
	if request == nil {
		request = NewDisassociateInstancesKeyPairsRequest()
	}
	response = NewDisassociateInstancesKeyPairsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLaunchTemplateNameRequest() (request *ModifyLaunchTemplateNameRequest) {
	request = &ModifyLaunchTemplateNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyLaunchTemplateName")
	return
}

func NewModifyLaunchTemplateNameResponse() (response *ModifyLaunchTemplateNameResponse) {
	response = &ModifyLaunchTemplateNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyLaunchTemplateName）用于修改一个实例启动模板的名称。
func (c *Client) ModifyLaunchTemplateName(request *ModifyLaunchTemplateNameRequest) (response *ModifyLaunchTemplateNameResponse, err error) {
	if request == nil {
		request = NewModifyLaunchTemplateNameRequest()
	}
	response = NewModifyLaunchTemplateNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
	request = &DescribeHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeHosts")
	return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
	response = &DescribeHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeHosts) 用于获取一个或多个CDH实例的详细信息。
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
	if request == nil {
		request = NewDescribeHostsRequest()
	}
	response = NewDescribeHostsResponse()
	err = c.Send(request, response)
	return
}

func NewColdMigrateInstanceRequest() (request *ColdMigrateInstanceRequest) {
	request = &ColdMigrateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ColdMigrateInstance")
	return
}

func NewColdMigrateInstanceResponse() (response *ColdMigrateInstanceResponse) {
	response = &ColdMigrateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 离线实例迁移。用于将一个系统盘镜像迁移到一个实例中。用户需要首先在腾讯云创建一个实例，然后通过本接口将需要运行的操作系统系统盘镜像迁移到该实例，实例的配置信息保持不变。
// 注意：为了防止兼容性错误，建议用户在创建实例的时候尽量选取与需要迁移的系统盘相似的操作系统。
func (c *Client) ColdMigrateInstance(request *ColdMigrateInstanceRequest) (response *ColdMigrateInstanceResponse, err error) {
	if request == nil {
		request = NewColdMigrateInstanceRequest()
	}
	response = NewColdMigrateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserMigrateTasksRequest() (request *DescribeUserMigrateTasksRequest) {
	request = &DescribeUserMigrateTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserMigrateTasks")
	return
}

func NewDescribeUserMigrateTasksResponse() (response *DescribeUserMigrateTasksResponse) {
	response = &DescribeUserMigrateTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUserMigrateTasks）用于查询用户的服务迁移任务记录。
func (c *Client) DescribeUserMigrateTasks(request *DescribeUserMigrateTasksRequest) (response *DescribeUserMigrateTasksResponse, err error) {
	if request == nil {
		request = NewDescribeUserMigrateTasksRequest()
	}
	response = NewDescribeUserMigrateTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesRenewFlagRequest() (request *ModifyInstancesRenewFlagRequest) {
	request = &ModifyInstancesRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesRenewFlag")
	return
}

func NewModifyInstancesRenewFlagResponse() (response *ModifyInstancesRenewFlagResponse) {
	response = &ModifyInstancesRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesRenewFlag) 用于修改包年包月实例续费标识。
//
// * 实例被标识为自动续费后，每次在实例到期时，会自动续费一个月。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) ModifyInstancesRenewFlag(request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyInstancesRenewFlagRequest()
	}
	response = NewModifyInstancesRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstancesActionTimerRequest() (request *DeleteInstancesActionTimerRequest) {
	request = &DeleteInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteInstancesActionTimer")
	return
}

func NewDeleteInstancesActionTimerResponse() (response *DeleteInstancesActionTimerResponse) {
	response = &DeleteInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除定时任务
func (c *Client) DeleteInstancesActionTimer(request *DeleteInstancesActionTimerRequest) (response *DeleteInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewDeleteInstancesActionTimerRequest()
	}
	response = NewDeleteInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMarketImagesRequest() (request *DescribeMarketImagesRequest) {
	request = &DescribeMarketImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeMarketImages")
	return
}

func NewDescribeMarketImagesResponse() (response *DescribeMarketImagesResponse) {
	response = &DescribeMarketImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护.
// 用于查询市场镜像。暂不暴露给用户
func (c *Client) DescribeMarketImages(request *DescribeMarketImagesRequest) (response *DescribeMarketImagesResponse, err error) {
	if request == nil {
		request = NewDescribeMarketImagesRequest()
	}
	response = NewDescribeMarketImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeNameConfigsRequest() (request *DescribeInstanceTypeNameConfigsRequest) {
	request = &DescribeInstanceTypeNameConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeNameConfigs")
	return
}

func NewDescribeInstanceTypeNameConfigsResponse() (response *DescribeInstanceTypeNameConfigsResponse) {
	response = &DescribeInstanceTypeNameConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceTypeNameConfigs) 用于查询实例类型的名称。
func (c *Client) DescribeInstanceTypeNameConfigs(request *DescribeInstanceTypeNameConfigsRequest) (response *DescribeInstanceTypeNameConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeNameConfigsRequest()
	}
	response = NewDescribeInstanceTypeNameConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCancelAuditMarketImageRequest() (request *CancelAuditMarketImageRequest) {
	request = &CancelAuditMarketImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CancelAuditMarketImage")
	return
}

func NewCancelAuditMarketImageResponse() (response *CancelAuditMarketImageResponse) {
	response = &CancelAuditMarketImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// None
func (c *Client) CancelAuditMarketImage(request *CancelAuditMarketImageRequest) (response *CancelAuditMarketImageResponse, err error) {
	if request == nil {
		request = NewCancelAuditMarketImageRequest()
	}
	response = NewCancelAuditMarketImageResponse()
	err = c.Send(request, response)
	return
}

func NewImportInstancesActionTimerRequest() (request *ImportInstancesActionTimerRequest) {
	request = &ImportInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportInstancesActionTimer")
	return
}

func NewImportInstancesActionTimerResponse() (response *ImportInstancesActionTimerResponse) {
	response = &ImportInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入定时销毁实例任务接口
func (c *Client) ImportInstancesActionTimer(request *ImportInstancesActionTimerRequest) (response *ImportInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewImportInstancesActionTimerRequest()
	}
	response = NewImportInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeQuotaRequest() (request *DescribeInstanceTypeQuotaRequest) {
	request = &DescribeInstanceTypeQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeQuota")
	return
}

func NewDescribeInstanceTypeQuotaResponse() (response *DescribeInstanceTypeQuotaResponse) {
	response = &DescribeInstanceTypeQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceTypeQuota) 用于查询实例的机型配额。
func (c *Client) DescribeInstanceTypeQuota(request *DescribeInstanceTypeQuotaRequest) (response *DescribeInstanceTypeQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeQuotaRequest()
	}
	response = NewDescribeInstanceTypeQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
	request = &DescribeZoneInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneInstanceConfigInfos")
	return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
	response = &DescribeZoneInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeZoneInstanceConfigInfos) 获取可用区的机型信息。
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneInstanceConfigInfosRequest()
	}
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewGetRecycleInfoRequest() (request *GetRecycleInfoRequest) {
	request = &GetRecycleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "GetRecycleInfo")
	return
}

func NewGetRecycleInfoResponse() (response *GetRecycleInfoResponse) {
	response = &GetRecycleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取回收站设置时间。
func (c *Client) GetRecycleInfo(request *GetRecycleInfoRequest) (response *GetRecycleInfoResponse, err error) {
	if request == nil {
		request = NewGetRecycleInfoRequest()
	}
	response = NewGetRecycleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneCpuQuotaRequest() (request *DescribeZoneCpuQuotaRequest) {
	request = &DescribeZoneCpuQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneCpuQuota")
	return
}

func NewDescribeZoneCpuQuotaResponse() (response *DescribeZoneCpuQuotaResponse) {
	response = &DescribeZoneCpuQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeZoneCpuQuota）用于查询可用区的CPU配额信息。
func (c *Client) DescribeZoneCpuQuota(request *DescribeZoneCpuQuotaRequest) (response *DescribeZoneCpuQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeZoneCpuQuotaRequest()
	}
	response = NewDescribeZoneCpuQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneCdhInstanceConfigInfosRequest() (request *DescribeZoneCdhInstanceConfigInfosRequest) {
	request = &DescribeZoneCdhInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneCdhInstanceConfigInfos")
	return
}

func NewDescribeZoneCdhInstanceConfigInfosResponse() (response *DescribeZoneCdhInstanceConfigInfosResponse) {
	response = &DescribeZoneCdhInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口自2022年4月14日起停止维护。
// 获取专用宿主机的机型配置信息,以及售罄状态信息列表。
func (c *Client) DescribeZoneCdhInstanceConfigInfos(request *DescribeZoneCdhInstanceConfigInfosRequest) (response *DescribeZoneCdhInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneCdhInstanceConfigInfosRequest()
	}
	response = NewDescribeZoneCdhInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesProjectRequest() (request *ModifyInstancesProjectRequest) {
	request = &ModifyInstancesProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesProject")
	return
}

func NewModifyInstancesProjectResponse() (response *ModifyInstancesProjectResponse) {
	response = &ModifyInstancesProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesProject) 用于修改实例所属项目。
//
// * 项目为一个虚拟概念，用户可以在一个账户下面建立多个项目，每个项目中管理不同的资源；将多个不同实例分属到不同项目中，后续使用 [DescribeInstances](DescribeInstances)接口查询实例，项目ID可用于过滤结果。
// * 绑定负载均衡的实例不支持修改实例所属项目。
// * 修改实例所属项目会自动解关联实例原来关联的安全组。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) ModifyInstancesProject(request *ModifyInstancesProjectRequest) (response *ModifyInstancesProjectResponse, err error) {
	if request == nil {
		request = NewModifyInstancesProjectRequest()
	}
	response = NewModifyInstancesProjectResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "StopInstances")
	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StopInstances) 用于关闭一个或多个实例。
//
// * 只有状态为`RUNNING`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`STOPPING`状态；关闭实例成功时，实例会进入`STOPPED`状态。
// * 支持强制关闭。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	response = NewStopInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImportImageOsRequest() (request *DescribeImportImageOsRequest) {
	request = &DescribeImportImageOsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImportImageOs")
	return
}

func NewDescribeImportImageOsResponse() (response *DescribeImportImageOsResponse) {
	response = &DescribeImportImageOsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看可以导入的镜像操作系统信息。
func (c *Client) DescribeImportImageOs(request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
	if request == nil {
		request = NewDescribeImportImageOsRequest()
	}
	response = NewDescribeImportImageOsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceChargeTypeConfigsRequest() (request *DescribeInstanceChargeTypeConfigsRequest) {
	request = &DescribeInstanceChargeTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceChargeTypeConfigs")
	return
}

func NewDescribeInstanceChargeTypeConfigsResponse() (response *DescribeInstanceChargeTypeConfigsResponse) {
	response = &DescribeInstanceChargeTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeInstanceChargeTypeConfigs) 查询 CVM 实例可支持的各种计费模式。
func (c *Client) DescribeInstanceChargeTypeConfigs(request *DescribeInstanceChargeTypeConfigsRequest) (response *DescribeInstanceChargeTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceChargeTypeConfigsRequest()
	}
	response = NewDescribeInstanceChargeTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCopyInstanceDiskRequest() (request *CopyInstanceDiskRequest) {
	request = &CopyInstanceDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CopyInstanceDisk")
	return
}

func NewCopyInstanceDiskResponse() (response *CopyInstanceDiskResponse) {
	response = &CopyInstanceDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CopyInstanceDisk) 用于拷贝实例的系统盘到弹性云硬盘。
//
// * 实例需要处于`关机`状态。
// * 需要指定未挂载、非加密弹性云硬盘，并且容量大于等于系统盘容量。
// * 实例与待挂载的磁盘需要在同一个可用区。
func (c *Client) CopyInstanceDisk(request *CopyInstanceDiskRequest) (response *CopyInstanceDiskResponse, err error) {
	if request == nil {
		request = NewCopyInstanceDiskRequest()
	}
	response = NewCopyInstanceDiskResponse()
	err = c.Send(request, response)
	return
}

func NewLiveMigrateRequest() (request *LiveMigrateRequest) {
	request = &LiveMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "LiveMigrate")
	return
}

func NewLiveMigrateResponse() (response *LiveMigrateResponse) {
	response = &LiveMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 热迁移实例
func (c *Client) LiveMigrate(request *LiveMigrateRequest) (response *LiveMigrateResponse, err error) {
	if request == nil {
		request = NewLiveMigrateRequest()
	}
	response = NewLiveMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterResetInstancesTypeRequest() (request *SwitchParameterResetInstancesTypeRequest) {
	request = &SwitchParameterResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterResetInstancesType")
	return
}

func NewSwitchParameterResetInstancesTypeResponse() (response *SwitchParameterResetInstancesTypeResponse) {
	response = &SwitchParameterResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterResetInstancesType) 用于调整实例的机型参数转换。
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口进行机型调整。
// * 目前不支持[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)实例使用该接口调整机型。* 目前不支持跨机型系统来调整机型，即使用该接口时指定的`InstanceType`和实例原来的机型需要属于同一系列。* 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
func (c *Client) SwitchParameterResetInstancesType(request *SwitchParameterResetInstancesTypeRequest) (response *SwitchParameterResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewSwitchParameterResetInstancesTypeRequest()
	}
	response = NewSwitchParameterResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesCreateImageAttributesRequest() (request *DescribeInstancesCreateImageAttributesRequest) {
	request = &DescribeInstancesCreateImageAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesCreateImageAttributes")
	return
}

func NewDescribeInstancesCreateImageAttributesResponse() (response *DescribeInstancesCreateImageAttributesResponse) {
	response = &DescribeInstancesCreateImageAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否支持在线制作整机镜像的属性接口
func (c *Client) DescribeInstancesCreateImageAttributes(request *DescribeInstancesCreateImageAttributesRequest) (response *DescribeInstancesCreateImageAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesCreateImageAttributesRequest()
	}
	response = NewDescribeInstancesCreateImageAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDisasterRecoverGroupRequest() (request *UpdateDisasterRecoverGroupRequest) {
	request = &UpdateDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "UpdateDisasterRecoverGroup")
	return
}

func NewUpdateDisasterRecoverGroupResponse() (response *UpdateDisasterRecoverGroupResponse) {
	response = &UpdateDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) UpdateDisasterRecoverGroup(request *UpdateDisasterRecoverGroupRequest) (response *UpdateDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewUpdateDisasterRecoverGroupRequest()
	}
	response = NewUpdateDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResizeInstanceDisksRequest() (request *InquiryPriceResizeInstanceDisksRequest) {
	request = &InquiryPriceResizeInstanceDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResizeInstanceDisks")
	return
}

func NewInquiryPriceResizeInstanceDisksResponse() (response *InquiryPriceResizeInstanceDisksResponse) {
	response = &InquiryPriceResizeInstanceDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (InquiryPriceResizeInstanceDisks) 用于扩容实例的系统盘询价。
//
// * 目前只支持扩容随实例购买的系统盘询价，且[系统盘类型](../数据结构#systemdisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。
func (c *Client) InquiryPriceResizeInstanceDisks(request *InquiryPriceResizeInstanceDisksRequest) (response *InquiryPriceResizeInstanceDisksResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResizeInstanceDisksRequest()
	}
	response = NewInquiryPriceResizeInstanceDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
	request = &DescribeAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAddresses")
	return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
	response = &DescribeAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
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

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (TerminateInstances) 用于主动退还实例。
//
// * 不再使用的实例，可通过本接口主动退还。
// * 按量计费的实例通过本接口可直接退还；包年包月实例如符合退还规则，也可通过本接口主动退还。
// * 支持批量操作，每次请求批量实例的上限为100。
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGlobalConfigsRequest() (request *DescribeUserGlobalConfigsRequest) {
	request = &DescribeUserGlobalConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserGlobalConfigs")
	return
}

func NewDescribeUserGlobalConfigsResponse() (response *DescribeUserGlobalConfigsResponse) {
	response = &DescribeUserGlobalConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户全局配置
func (c *Client) DescribeUserGlobalConfigs(request *DescribeUserGlobalConfigsRequest) (response *DescribeUserGlobalConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeUserGlobalConfigsRequest()
	}
	response = NewDescribeUserGlobalConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserZoneStatusRequest() (request *DescribeUserZoneStatusRequest) {
	request = &DescribeUserZoneStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserZoneStatus")
	return
}

func NewDescribeUserZoneStatusResponse() (response *DescribeUserZoneStatusResponse) {
	response = &DescribeUserZoneStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeUserZoneStatus) 查询可用区实例计费类型状态。
func (c *Client) DescribeUserZoneStatus(request *DescribeUserZoneStatusRequest) (response *DescribeUserZoneStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUserZoneStatusRequest()
	}
	response = NewDescribeUserZoneStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDisasterRecoverGroupAttributeRequest() (request *ModifyDisasterRecoverGroupAttributeRequest) {
	request = &ModifyDisasterRecoverGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyDisasterRecoverGroupAttribute")
	return
}

func NewModifyDisasterRecoverGroupAttributeResponse() (response *ModifyDisasterRecoverGroupAttributeResponse) {
	response = &ModifyDisasterRecoverGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyDisasterRecoverGroupAttribute)用于修改[分散置放群组](/tcloud/Compute/CVM/292128/661327/cvm_group)属性。
func (c *Client) ModifyDisasterRecoverGroupAttribute(request *ModifyDisasterRecoverGroupAttributeRequest) (response *ModifyDisasterRecoverGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyDisasterRecoverGroupAttributeRequest()
	}
	response = NewModifyDisasterRecoverGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewEnterRescueModeRequest() (request *EnterRescueModeRequest) {
	request = &EnterRescueModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "EnterRescueMode")
	return
}

func NewEnterRescueModeResponse() (response *EnterRescueModeResponse) {
	response = &EnterRescueModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 进入救援模式。
func (c *Client) EnterRescueMode(request *EnterRescueModeRequest) (response *EnterRescueModeResponse, err error) {
	if request == nil {
		request = NewEnterRescueModeRequest()
	}
	response = NewEnterRescueModeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateInstancesActionTimerRequest() (request *UpdateInstancesActionTimerRequest) {
	request = &UpdateInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "UpdateInstancesActionTimer")
	return
}

func NewUpdateInstancesActionTimerResponse() (response *UpdateInstancesActionTimerResponse) {
	response = &UpdateInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改销毁实例定时任务接口
func (c *Client) UpdateInstancesActionTimer(request *UpdateInstancesActionTimerRequest) (response *UpdateInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewUpdateInstancesActionTimerRequest()
	}
	response = NewUpdateInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewTransformAddressRequest() (request *TransformAddressRequest) {
	request = &TransformAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "TransformAddress")
	return
}

func NewTransformAddressResponse() (response *TransformAddressResponse) {
	response = &TransformAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (TransformAddress) 用于将实例的普通公网 IP 转换为[弹性公网IP](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)（简称 EIP）。
// * 平台对用户每地域每日解绑 EIP 重新分配普通公网 IP 次数有所限制（可参见 [EIP 产品简介](/tcloud/Compute/CVM/292128/661327/nasinfo_flexpnip)）。上述配额可通过 [DescribeAddressQuota](DescribeAddressQuota) 接口获取。
func (c *Client) TransformAddress(request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
	if request == nil {
		request = NewTransformAddressRequest()
	}
	response = NewTransformAddressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务信息, des或者allinone的taskid的信息
// 只支持前端调用不支持api接口调用
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	response = NewDescribeTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceInternetBandwidthConfigsRequest() (request *DescribeInstanceInternetBandwidthConfigsRequest) {
	request = &DescribeInstanceInternetBandwidthConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceInternetBandwidthConfigs")
	return
}

func NewDescribeInstanceInternetBandwidthConfigsResponse() (response *DescribeInstanceInternetBandwidthConfigsResponse) {
	response = &DescribeInstanceInternetBandwidthConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceInternetBandwidthConfigs) 用于查询实例带宽配置。
//
// * 只支持查询`BANDWIDTH_PREPAID`计费模式的带宽配置。
// * 接口返回实例的所有带宽配置信息（包含历史的带宽配置信息）。
func (c *Client) DescribeInstanceInternetBandwidthConfigs(request *DescribeInstanceInternetBandwidthConfigsRequest) (response *DescribeInstanceInternetBandwidthConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceInternetBandwidthConfigsRequest()
	}
	response = NewDescribeInstanceInternetBandwidthConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeRegions)用于查询地域信息。
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response = NewDescribeRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAvailableZonesRequest() (request *DescribeUserAvailableZonesRequest) {
	request = &DescribeUserAvailableZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserAvailableZones")
	return
}

func NewDescribeUserAvailableZonesResponse() (response *DescribeUserAvailableZonesResponse) {
	response = &DescribeUserAvailableZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户可用分区
func (c *Client) DescribeUserAvailableZones(request *DescribeUserAvailableZonesRequest) (response *DescribeUserAvailableZonesResponse, err error) {
	if request == nil {
		request = NewDescribeUserAvailableZonesRequest()
	}
	response = NewDescribeUserAvailableZonesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceAllocateAddressesRequest() (request *InquiryPriceAllocateAddressesRequest) {
	request = &InquiryPriceAllocateAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceAllocateAddresses")
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

func NewDescribeImagesAttributeRequest() (request *DescribeImagesAttributeRequest) {
	request = &DescribeImagesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImagesAttribute")
	return
}

func NewDescribeImagesAttributeResponse() (response *DescribeImagesAttributeResponse) {
	response = &DescribeImagesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据unImgId查询DeviceImageId
// 该接口不支持公有镜像id的查询
func (c *Client) DescribeImagesAttribute(request *DescribeImagesAttributeRequest) (response *DescribeImagesAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeImagesAttributeRequest()
	}
	response = NewDescribeImagesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterRenewHostsRequest() (request *SwitchParameterRenewHostsRequest) {
	request = &SwitchParameterRenewHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterRenewHosts")
	return
}

func NewSwitchParameterRenewHostsResponse() (response *SwitchParameterRenewHostsResponse) {
	response = &SwitchParameterRenewHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费CDH实例参数转换
func (c *Client) SwitchParameterRenewHosts(request *SwitchParameterRenewHostsRequest) (response *SwitchParameterRenewHostsResponse, err error) {
	if request == nil {
		request = NewSwitchParameterRenewHostsRequest()
	}
	response = NewSwitchParameterRenewHostsResponse()
	err = c.Send(request, response)
	return
}

func NewFailoverMigrateRequest() (request *FailoverMigrateRequest) {
	request = &FailoverMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "FailoverMigrate")
	return
}

func NewFailoverMigrateResponse() (response *FailoverMigrateResponse) {
	response = &FailoverMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 故障迁移
func (c *Client) FailoverMigrate(request *FailoverMigrateRequest) (response *FailoverMigrateResponse, err error) {
	if request == nil {
		request = NewFailoverMigrateRequest()
	}
	response = NewFailoverMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewHostsRequest() (request *InquiryPriceRenewHostsRequest) {
	request = &InquiryPriceRenewHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRenewHosts")
	return
}

func NewInquiryPriceRenewHostsResponse() (response *InquiryPriceRenewHostsResponse) {
	response = &InquiryPriceRenewHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费CDH实例询价
func (c *Client) InquiryPriceRenewHosts(request *InquiryPriceRenewHostsRequest) (response *InquiryPriceRenewHostsResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewHostsRequest()
	}
	response = NewInquiryPriceRenewHostsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
	request = &ModifyImageAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageAttribute")
	return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
	response = &ModifyImageAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyImageAttribute）用于修改镜像属性。
//
// * 已分享的镜像无法修改属性。
func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
	if request == nil {
		request = NewModifyImageAttributeRequest()
	}
	response = NewModifyImageAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaunchTemplateRequest() (request *DeleteLaunchTemplateRequest) {
	request = &DeleteLaunchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteLaunchTemplate")
	return
}

func NewDeleteLaunchTemplateResponse() (response *DeleteLaunchTemplateResponse) {
	response = &DeleteLaunchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteLaunchTemplate）用于删除一个实例启动模板。
func (c *Client) DeleteLaunchTemplate(request *DeleteLaunchTemplateRequest) (response *DeleteLaunchTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteLaunchTemplateRequest()
	}
	response = NewDeleteLaunchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportFullCvmImageRequest() (request *ImportFullCvmImageRequest) {
	request = &ImportFullCvmImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportFullCvmImage")
	return
}

func NewImportFullCvmImageResponse() (response *ImportFullCvmImageResponse) {
	response = &ImportFullCvmImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口自2022年4月14月起停止维护。
// 可以同时导入系统盘镜像和数据盘，并制作整机镜像。
func (c *Client) ImportFullCvmImage(request *ImportFullCvmImageRequest) (response *ImportFullCvmImageResponse, err error) {
	if request == nil {
		request = NewImportFullCvmImageRequest()
	}
	response = NewImportFullCvmImageResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthRequest() (request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) {
	request = &InquiryPriceResetInstancesInternetMaxBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesInternetMaxBandwidth")
	return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthResponse() (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse) {
	response = &InquiryPriceResetInstancesInternetMaxBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年5月11日起停止维护。
// 本接口 (InquiryPriceResetInstancesInternetMaxBandwidth) 用于调整实例公网带宽上限询价。
//
// * 不同机型带宽上限范围不一致。
// * 对于`BANDWIDTH_PREPAID`计费方式的带宽，需要输入参数`StartTime`和`EndTime`，指定调整后的带宽的生效时间段。在这种场景下目前不支持调小带宽，会涉及扣费，请确保账户余额充足。
// * 对于 `TRAFFIC_POSTPAID_BY_HOUR`、 `BANDWIDTH_POSTPAID_BY_HOUR` 和 `BANDWIDTH_PACKAGE` 计费方式的带宽，使用该接口调整带宽上限是实时生效的，可以在带宽允许的范围内调大或者调小带宽，不支持输入参数 `StartTime` 和 `EndTime` 。
// * 接口不支持调整`BANDWIDTH_POSTPAID_BY_MONTH`计费方式的带宽。
// * 接口不支持批量调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽。
// * 接口不支持批量调整混合计费方式的带宽。例如不支持同时调整`TRAFFIC_POSTPAID_BY_HOUR`和`BANDWIDTH_PACKAGE`计费方式的带宽。
// 接口仅支持实例带宽计费方式为 `BANDWIDTH_PREPAID` 和 `TRAFFIC_POSTPAID_BY_HOUR` 的询价， 其余均返回0。
func (c *Client) InquiryPriceResetInstancesInternetMaxBandwidth(request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResetInstancesInternetMaxBandwidthRequest()
	}
	response = NewInquiryPriceResetInstancesInternetMaxBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterModifyInstancesChargeTypeRequest() (request *SwitchParameterModifyInstancesChargeTypeRequest) {
	request = &SwitchParameterModifyInstancesChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterModifyInstancesChargeType")
	return
}

func NewSwitchParameterModifyInstancesChargeTypeResponse() (response *SwitchParameterModifyInstancesChargeTypeResponse) {
	response = &SwitchParameterModifyInstancesChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterModifyInstancesChargeType) 用于切换实例的计费模式参数转换。
//
// * 只支持从 `POSTPAID_BY_HOUR` 计费模式切换为`PREPAID`计费模式。
// * 关机不收费的实例、`BC1`和`BS1`机型的实例、设置定时销毁的实例不支持该操作。
func (c *Client) SwitchParameterModifyInstancesChargeType(request *SwitchParameterModifyInstancesChargeTypeRequest) (response *SwitchParameterModifyInstancesChargeTypeResponse, err error) {
	if request == nil {
		request = NewSwitchParameterModifyInstancesChargeTypeRequest()
	}
	response = NewSwitchParameterModifyInstancesChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteImages")
	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteImages）用于删除一个或多个镜像。
//
// * 当[镜像状态](../数据结构#image)为`创建中`和`使用中`时, 不允许删除。镜像状态可以通过[DescribeImages](DescribeImages)获取。
// * 每个地域最多只支持创建10个自定义镜像，删除镜像可以释放账户的配额。
// * 当镜像正在被其它账户分享时，不允许删除。
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	if request == nil {
		request = NewDeleteImagesRequest()
	}
	response = NewDeleteImagesResponse()
	err = c.Send(request, response)
	return
}

func NewImportCbsRequest() (request *ImportCbsRequest) {
	request = &ImportCbsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportCbs")
	return
}

func NewImportCbsResponse() (response *ImportCbsResponse) {
	response = &ImportCbsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ImportCbs)用于导入数据盘至一块云盘(CBS)中。
//
// * 用户需要确认需要导入的数据盘大小小于云盘的容量。
// * 用户需要传入未挂载的云盘作为参数以进行导入。
// * 导入过程中请不要操作云盘。
func (c *Client) ImportCbs(request *ImportCbsRequest) (response *ImportCbsResponse, err error) {
	if request == nil {
		request = NewImportCbsRequest()
	}
	response = NewImportCbsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDisasterRecoverGroupRequest() (request *DeleteDisasterRecoverGroupRequest) {
	request = &DeleteDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteDisasterRecoverGroup")
	return
}

func NewDeleteDisasterRecoverGroupResponse() (response *DeleteDisasterRecoverGroupResponse) {
	response = &DeleteDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容灾组,仅空容灾组才能被删除。
func (c *Client) DeleteDisasterRecoverGroup(request *DeleteDisasterRecoverGroupRequest) (response *DeleteDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewDeleteDisasterRecoverGroupRequest()
	}
	response = NewDeleteDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceFamilyConfigsRequest() (request *DescribeInstanceFamilyConfigsRequest) {
	request = &DescribeInstanceFamilyConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceFamilyConfigs")
	return
}

func NewDescribeInstanceFamilyConfigsResponse() (response *DescribeInstanceFamilyConfigsResponse) {
	response = &DescribeInstanceFamilyConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstanceFamilyConfigs）查询当前用户和地域所支持的机型族列表信息。
func (c *Client) DescribeInstanceFamilyConfigs(request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceFamilyConfigsRequest()
	}
	response = NewDescribeInstanceFamilyConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDisasterRecoverGroupRequest() (request *QueryDisasterRecoverGroupRequest) {
	request = &QueryDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "QueryDisasterRecoverGroup")
	return
}

func NewQueryDisasterRecoverGroupResponse() (response *QueryDisasterRecoverGroupResponse) {
	response = &QueryDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分散置放群组信息
func (c *Client) QueryDisasterRecoverGroup(request *QueryDisasterRecoverGroupRequest) (response *QueryDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewQueryDisasterRecoverGroupRequest()
	}
	response = NewQueryDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
	request = &DescribeInstancesStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesStatus")
	return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
	response = &DescribeInstancesStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstancesStatus) 用于查询一个或多个实例的状态。
//
// * 可以根据实例`ID`来查询实例的状态。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的实例状态。
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesStatusRequest()
	}
	response = NewDescribeInstancesStatusResponse()
	err = c.Send(request, response)
	return
}

func NewResizeInstanceDisksRequest() (request *ResizeInstanceDisksRequest) {
	request = &ResizeInstanceDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResizeInstanceDisks")
	return
}

func NewResizeInstanceDisksResponse() (response *ResizeInstanceDisksResponse) {
	response = &ResizeInstanceDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResizeInstanceDisks) 用于扩容实例的系统盘。
//
// * 目前只支持扩容随实例购买的系统盘，且[系统盘类型](../数据结构#systemdisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。
func (c *Client) ResizeInstanceDisks(request *ResizeInstanceDisksRequest) (response *ResizeInstanceDisksResponse, err error) {
	if request == nil {
		request = NewResizeInstanceDisksRequest()
	}
	response = NewResizeInstanceDisksResponse()
	err = c.Send(request, response)
	return
}

func NewCheckInstancesConnectivityRequest() (request *CheckInstancesConnectivityRequest) {
	request = &CheckInstancesConnectivityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CheckInstancesConnectivity")
	return
}

func NewCheckInstancesConnectivityResponse() (response *CheckInstancesConnectivityResponse) {
	response = &CheckInstancesConnectivityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstancesConnectivity) 用于查询一个或多个实例的网络连通性，包括操作系统默认的远程登录服务端口和ICMP。
// - 如果是Linux操作系统，会检查默认远程登录端口号 22
// - 如果是Windows操作系统，会检查默认远程登录端口号3389
func (c *Client) CheckInstancesConnectivity(request *CheckInstancesConnectivityRequest) (response *CheckInstancesConnectivityResponse, err error) {
	if request == nil {
		request = NewCheckInstancesConnectivityRequest()
	}
	response = NewCheckInstancesConnectivityResponse()
	err = c.Send(request, response)
	return
}

func NewExitLiveMigrateInstanceRequest() (request *ExitLiveMigrateInstanceRequest) {
	request = &ExitLiveMigrateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ExitLiveMigrateInstance")
	return
}

func NewExitLiveMigrateInstanceResponse() (response *ExitLiveMigrateInstanceResponse) {
	response = &ExitLiveMigrateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退出在线服务迁移。用于使实例退出在线服务迁移模式，一般使用在在线服务迁移成功或者失败后。该接口不可直接被用户调用。
func (c *Client) ExitLiveMigrateInstance(request *ExitLiveMigrateInstanceRequest) (response *ExitLiveMigrateInstanceResponse, err error) {
	if request == nil {
		request = NewExitLiveMigrateInstanceRequest()
	}
	response = NewExitLiveMigrateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RebootInstances")
	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RebootInstances) 用于重启实例。
//
// * 只有状态为`RUNNING`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`REBOOTING`状态；重启实例成功时，实例会进入`RUNNING`状态。
// * 支持强制重启。强制重启的效果等同于关闭物理计算机的电源开关再重新启动。强制重启可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常重启时使用。
// * 支持批量操作，每次请求批量实例的上限为100。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	response = NewRebootInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshInternalUserEnvironmentRequest() (request *RefreshInternalUserEnvironmentRequest) {
	request = &RefreshInternalUserEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RefreshInternalUserEnvironment")
	return
}

func NewRefreshInternalUserEnvironmentResponse() (response *RefreshInternalUserEnvironmentResponse) {
	response = &RefreshInternalUserEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// None
func (c *Client) RefreshInternalUserEnvironment(request *RefreshInternalUserEnvironmentRequest) (response *RefreshInternalUserEnvironmentResponse, err error) {
	if request == nil {
		request = NewRefreshInternalUserEnvironmentRequest()
	}
	response = NewRefreshInternalUserEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesAttribute")
	return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
	response = &ModifyInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。
//
// * “实例名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行实例管理操作的依据。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyInstancesAttributeRequest()
	}
	response = NewModifyInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
	request = &ReleaseAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ReleaseAddresses")
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

func NewDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
	request = &DescribeImageSharePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageSharePermission")
	return
}

func NewDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
	response = &DescribeImageSharePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescrbieImageSharePermission）用于查看镜像分享信息。
func (c *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (response *DescribeImageSharePermissionResponse, err error) {
	if request == nil {
		request = NewDescribeImageSharePermissionRequest()
	}
	response = NewDescribeImageSharePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
	request = &DescribeInstanceVncUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceVncUrl")
	return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
	response = &DescribeInstanceVncUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询实例 VNC 地址
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceVncUrlRequest()
	}
	response = NewDescribeInstanceVncUrlResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewAddressesRequest() (request *InquiryPriceRenewAddressesRequest) {
	request = &InquiryPriceRenewAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRenewAddresses")
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

func NewDescribeImageQuotaRequest() (request *DescribeImageQuotaRequest) {
	request = &DescribeImageQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageQuota")
	return
}

func NewDescribeImageQuotaResponse() (response *DescribeImageQuotaResponse) {
	response = &DescribeImageQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeImageQuota)用于查询用户帐号的镜像配额。
func (c *Client) DescribeImageQuota(request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeImageQuotaRequest()
	}
	response = NewDescribeImageQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewRenewAddressesRequest() (request *RenewAddressesRequest) {
	request = &RenewAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RenewAddresses")
	return
}

func NewRenewAddressesResponse() (response *RenewAddressesResponse) {
	response = &RenewAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) RenewAddresses(request *RenewAddressesRequest) (response *RenewAddressesResponse, err error) {
	if request == nil {
		request = NewRenewAddressesRequest()
	}
	response = NewRenewAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateHostsRequest() (request *AllocateHostsRequest) {
	request = &AllocateHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AllocateHosts")
	return
}

func NewAllocateHostsResponse() (response *AllocateHostsResponse) {
	response = &AllocateHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AllocateHosts) 用于创建一个或多个指定配置的CDH实例。
// * 当HostChargeType为PREPAID时，必须指定HostChargePrepaid参数。
//
// 3100及以后版本只支持后付费，不支持预付费.传预付费参数会报错。
func (c *Client) AllocateHosts(request *AllocateHostsRequest) (response *AllocateHostsResponse, err error) {
	if request == nil {
		request = NewAllocateHostsRequest()
	}
	response = NewAllocateHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisasterRecoverGroupQuotaRequest() (request *DescribeDisasterRecoverGroupQuotaRequest) {
	request = &DescribeDisasterRecoverGroupQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroupQuota")
	return
}

func NewDescribeDisasterRecoverGroupQuotaResponse() (response *DescribeDisasterRecoverGroupQuotaResponse) {
	response = &DescribeDisasterRecoverGroupQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询置放群组配额
func (c *Client) DescribeDisasterRecoverGroupQuota(request *DescribeDisasterRecoverGroupQuotaRequest) (response *DescribeDisasterRecoverGroupQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoverGroupQuotaRequest()
	}
	response = NewDescribeDisasterRecoverGroupQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceStatisticsRequest() (request *DescribeInstanceStatisticsRequest) {
	request = &DescribeInstanceStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceStatistics")
	return
}

func NewDescribeInstanceStatisticsResponse() (response *DescribeInstanceStatisticsResponse) {
	response = &DescribeInstanceStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceStatistics) 用于查询各个地域资源概览。
func (c *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (response *DescribeInstanceStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceStatisticsRequest()
	}
	response = NewDescribeInstanceStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceAttachedDevicesRequest() (request *DescribeInstanceAttachedDevicesRequest) {
	request = &DescribeInstanceAttachedDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceAttachedDevices")
	return
}

func NewDescribeInstanceAttachedDevicesResponse() (response *DescribeInstanceAttachedDevicesResponse) {
	response = &DescribeInstanceAttachedDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年5月11日起停止维护。
// 用于查询实例的设备信息
// 只支持GPU实例,非GPU实例查询不会返回
func (c *Client) DescribeInstanceAttachedDevices(request *DescribeInstanceAttachedDevicesRequest) (response *DescribeInstanceAttachedDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceAttachedDevicesRequest()
	}
	response = NewDescribeInstanceAttachedDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRunInstancesRequest() (request *InquiryPriceRunInstancesRequest) {
	request = &InquiryPriceRunInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRunInstances")
	return
}

func NewInquiryPriceRunInstancesResponse() (response *InquiryPriceRunInstancesResponse) {
	response = &InquiryPriceRunInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(InquiryPriceRunInstances)用于创建实例询价。本接口仅允许针对购买限制范围内的实例配置进行询价, 详见: [RunInstances](RunInstances)。
func (c *Client) InquiryPriceRunInstances(request *InquiryPriceRunInstancesRequest) (response *InquiryPriceRunInstancesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRunInstancesRequest()
	}
	response = NewInquiryPriceRunInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceConfigInfosRequest() (request *DescribeInstanceConfigInfosRequest) {
	request = &DescribeInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceConfigInfos")
	return
}

func NewDescribeInstanceConfigInfosResponse() (response *DescribeInstanceConfigInfosResponse) {
	response = &DescribeInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeInstanceConfigInfos) 获取实例静态配置信息，包含CPU核数、CPU型号、内存大小和带宽信息等。
func (c *Client) DescribeInstanceConfigInfos(request *DescribeInstanceConfigInfosRequest) (response *DescribeInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceConfigInfosRequest()
	}
	response = NewDescribeInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
	request = &DescribeAddressQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAddressQuota")
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

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZones")
	return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeZones)用于查询可用区信息。
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	if request == nil {
		request = NewDescribeZonesRequest()
	}
	response = NewDescribeZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesModificationRequest() (request *DescribeInstancesModificationRequest) {
	request = &DescribeInstancesModificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesModification")
	return
}

func NewDescribeInstancesModificationResponse() (response *DescribeInstancesModificationResponse) {
	response = &DescribeInstancesModificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeInstancesModification
func (c *Client) DescribeInstancesModification(request *DescribeInstancesModificationRequest) (response *DescribeInstancesModificationResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesModificationRequest()
	}
	response = NewDescribeInstancesModificationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDisasterRecoverGroupsRequest() (request *DeleteDisasterRecoverGroupsRequest) {
	request = &DeleteDisasterRecoverGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteDisasterRecoverGroups")
	return
}

func NewDeleteDisasterRecoverGroupsResponse() (response *DeleteDisasterRecoverGroupsResponse) {
	response = &DeleteDisasterRecoverGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteDisasterRecoverGroups)用于删除[分散置放群组](/tcloud/Compute/CVM/292128/661327/cvm_group)。只有空的置放群组才能被删除，非空的群组需要先销毁组内所有云服务器，才能执行删除操作，不然会产生删除置放群组失败的错误。
func (c *Client) DeleteDisasterRecoverGroups(request *DeleteDisasterRecoverGroupsRequest) (response *DeleteDisasterRecoverGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteDisasterRecoverGroupsRequest()
	}
	response = NewDeleteDisasterRecoverGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesReturnableRequest() (request *DescribeInstancesReturnableRequest) {
	request = &DescribeInstancesReturnableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesReturnable")
	return
}

func NewDescribeInstancesReturnableResponse() (response *DescribeInstancesReturnableResponse) {
	response = &DescribeInstancesReturnableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstancesReturnable) 用于检查子机是否可退还。
func (c *Client) DescribeInstancesReturnable(request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesReturnableRequest()
	}
	response = NewDescribeInstancesReturnableResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceUsbInfoRequest() (request *DescribeInstanceUsbInfoRequest) {
	request = &DescribeInstanceUsbInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceUsbInfo")
	return
}

func NewDescribeInstanceUsbInfoResponse() (response *DescribeInstanceUsbInfoResponse) {
	response = &DescribeInstanceUsbInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例USB信息
func (c *Client) DescribeInstanceUsbInfo(request *DescribeInstanceUsbInfoRequest) (response *DescribeInstanceUsbInfoResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceUsbInfoRequest()
	}
	response = NewDescribeInstanceUsbInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLaunchTemplateRequest() (request *CreateLaunchTemplateRequest) {
	request = &CreateLaunchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateLaunchTemplate")
	return
}

func NewCreateLaunchTemplateResponse() (response *CreateLaunchTemplateResponse) {
	response = &CreateLaunchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateLaunchTemplate）用于创建实例启动模板。
//
// 实例启动模板是一种配置数据并可用于创建实例，其内容包含创建实例所需的配置，比如实例类型，数据盘和系统盘的类型和大小，以及安全组等信息。
func (c *Client) CreateLaunchTemplate(request *CreateLaunchTemplateRequest) (response *CreateLaunchTemplateResponse, err error) {
	if request == nil {
		request = NewCreateLaunchTemplateRequest()
	}
	response = NewCreateLaunchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecommendedZonesRequest() (request *DescribeRecommendedZonesRequest) {
	request = &DescribeRecommendedZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeRecommendedZones")
	return
}

func NewDescribeRecommendedZonesResponse() (response *DescribeRecommendedZonesResponse) {
	response = &DescribeRecommendedZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用区规格配额排序列表
func (c *Client) DescribeRecommendedZones(request *DescribeRecommendedZonesRequest) (response *DescribeRecommendedZonesResponse, err error) {
	if request == nil {
		request = NewDescribeRecommendedZonesRequest()
	}
	response = NewDescribeRecommendedZonesResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
	request = &AssociateInstancesKeyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AssociateInstancesKeyPairs")
	return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
	response = &AssociateInstancesKeyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AssociateInstancesKeyPairs) 用于将密钥绑定到实例上。
//
// * 将密钥的公钥写入到实例的`SSH`配置当中，用户就可以通过该密钥的私钥来登录实例。
// * 如果实例原来绑定过密钥，那么原来的密钥将失效。
// * 如果实例原来是通过密码登录，绑定密钥后无法使用密码登录。
// * 支持批量操作。每次请求批量实例的上限为100。如果批量实例存在不允许操作的实例，操作会以特定错误码返回。
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
	if request == nil {
		request = NewAssociateInstancesKeyPairsRequest()
	}
	response = NewAssociateInstancesKeyPairsResponse()
	err = c.Send(request, response)
	return
}

func NewColdMigrateRequest() (request *ColdMigrateRequest) {
	request = &ColdMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ColdMigrate")
	return
}

func NewColdMigrateResponse() (response *ColdMigrateResponse) {
	response = &ColdMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起不再维护。
// 冷迁移实例
func (c *Client) ColdMigrate(request *ColdMigrateRequest) (response *ColdMigrateResponse, err error) {
	if request == nil {
		request = NewColdMigrateRequest()
	}
	response = NewColdMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceOperationLogsRequest() (request *DescribeInstanceOperationLogsRequest) {
	request = &DescribeInstanceOperationLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceOperationLogs")
	return
}

func NewDescribeInstanceOperationLogsResponse() (response *DescribeInstanceOperationLogsResponse) {
	response = &DescribeInstanceOperationLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstanceOperationLogs）查询指定实例操作记录。
func (c *Client) DescribeInstanceOperationLogs(request *DescribeInstanceOperationLogsRequest) (response *DescribeInstanceOperationLogsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceOperationLogsRequest()
	}
	response = NewDescribeInstanceOperationLogsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchUserInstanceRequest() (request *SearchUserInstanceRequest) {
	request = &SearchUserInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SearchUserInstance")
	return
}

func NewSearchUserInstanceResponse() (response *SearchUserInstanceResponse) {
	response = &SearchUserInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年5月19日起停止维护。
// 全地域搜索用户实例
func (c *Client) SearchUserInstance(request *SearchUserInstanceRequest) (response *SearchUserInstanceResponse, err error) {
	if request == nil {
		request = NewSearchUserInstanceRequest()
	}
	response = NewSearchUserInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyInstancesChargeTypeRequest() (request *InquiryPriceModifyInstancesChargeTypeRequest) {
	request = &InquiryPriceModifyInstancesChargeTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceModifyInstancesChargeType")
	return
}

func NewInquiryPriceModifyInstancesChargeTypeResponse() (response *InquiryPriceModifyInstancesChargeTypeResponse) {
	response = &InquiryPriceModifyInstancesChargeTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (InquiryPriceModifyInstancesChargeType) 用于切换实例的计费模式询价。
//
// * 只支持从 `POSTPAID_BY_HOUR` 计费模式切换为`PREPAID`计费模式。
// * 关机不收费的实例、`BC1`和`BS1`机型族的实例、设置定时销毁的实例不支持该操作。
func (c *Client) InquiryPriceModifyInstancesChargeType(request *InquiryPriceModifyInstancesChargeTypeRequest) (response *InquiryPriceModifyInstancesChargeTypeResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyInstancesChargeTypeRequest()
	}
	response = NewInquiryPriceModifyInstancesChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResetInstancesTypeRequest() (request *InquiryPriceResetInstancesTypeRequest) {
	request = &InquiryPriceResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesType")
	return
}

func NewInquiryPriceResetInstancesTypeResponse() (response *InquiryPriceResetInstancesTypeResponse) {
	response = &InquiryPriceResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (InquiryPriceResetInstancesType) 用于调整实例的机型询价。
//
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口进行调整机型询价。
// * 目前不支持[CDH](/tcloud/Compute/CVM/292128/347831/cdh_overview)实例使用该接口调整机型询价。
// * 目前不支持跨机型系统来调整机型，即使用该接口时指定的`InstanceType`和实例原来的机型需要属于同一系列。
// * 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
func (c *Client) InquiryPriceResetInstancesType(request *InquiryPriceResetInstancesTypeRequest) (response *InquiryPriceResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResetInstancesTypeRequest()
	}
	response = NewInquiryPriceResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesVpcAttributeRequest() (request *ModifyInstancesVpcAttributeRequest) {
	request = &ModifyInstancesVpcAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesVpcAttribute")
	return
}

func NewModifyInstancesVpcAttributeResponse() (response *ModifyInstancesVpcAttributeResponse) {
	response = &ModifyInstancesVpcAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ModifyInstancesVpcAttribute)用于修改实例vpc属性，如私有网络ip。
// * 此操作默认会关闭实例，完成后再启动。
// * 当指定私有网络ID和子网ID（子网必须在实例所在的可用区）与指定实例所在私有网络不一致时，会将实例迁移至指定的私有网络的子网下。执行此操作前请确保指定的实例上没有绑定[弹性网卡](/tcloud/NetWork/ENI)和[负载均衡](/tcloud/NetWork/CLB)。
// * 实例操作结果可以通过调用 [DescribeInstances](DescribeInstances) 接口查询，如果实例的最新操作状态(LatestOperationState)为“SUCCESS”，则代表操作成功。
func (c *Client) ModifyInstancesVpcAttribute(request *ModifyInstancesVpcAttributeRequest) (response *ModifyInstancesVpcAttributeResponse, err error) {
	if request == nil {
		request = NewModifyInstancesVpcAttributeRequest()
	}
	response = NewModifyInstancesVpcAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddressBandwidthConfigsRequest() (request *DescribeAddressBandwidthConfigsRequest) {
	request = &DescribeAddressBandwidthConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAddressBandwidthConfigs")
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

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
	request = &CreateKeyPairRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateKeyPair")
	return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
	response = &CreateKeyPairResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CreateKeyPair) 用于创建一个 `OpenSSH RSA` 密钥对，可以用于登录 `Linux` 实例。
//
// * 开发者只需指定密钥对名称，即可由系统自动创建密钥对，并返回所生成的密钥对的 `ID` 及其公钥、私钥的内容。
// * 密钥对名称不能和已经存在的密钥对的名称重复。
// * 私钥的内容可以保存到文件中作为 `SSH` 的一种认证方式。
// * 腾讯云不会保存用户的私钥，请妥善保管。
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
	if request == nil {
		request = NewCreateKeyPairRequest()
	}
	response = NewCreateKeyPairResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "StartInstances")
	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StartInstances) 用于启动一个或多个实例。
//
// * 只有状态为`STOPPED`的实例才可以进行此操作。
// * 接口调用成功时，实例会进入`STARTING`状态；启动实例成功时，实例会进入`RUNNING`状态。
// * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	response = NewStartInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
	request = &DescribeInstancesDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesDeniedActions")
	return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
	response = &DescribeInstancesDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstancesDeniedActions) 用于查询一个或多个实例的操作限制列表信息。
//
// * 可以根据实例`ID`来查询实例的详细信息。
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesDeniedActionsRequest()
	}
	response = NewDescribeInstancesDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResetInstanceRequest() (request *InquiryPriceResetInstanceRequest) {
	request = &InquiryPriceResetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstance")
	return
}

func NewInquiryPriceResetInstanceResponse() (response *InquiryPriceResetInstanceResponse) {
	response = &InquiryPriceResetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (InquiryPriceResetInstance) 用于重装实例询价。* 如果指定了`ImageId`参数，则使用指定的镜像进行重装询价；否则按照当前实例使用的镜像进行重装询价。* 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口实现`Linux`和`Windows`操作系统切换的重装询价。* 目前不支持海外地域的实例使用该接口实现`Linux`和`Windows`操作系统切换的重装询价。
//
// 指定有效的[镜像](/tcloud/Compute/CVM/292128/835305/mirr_overview)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li></li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](//console.{{conf.main_domain}}/cvm/image/list?imageType=PUBLIC_IMAGE&pageIndex=1&pageSize=20)查询；</li><li>通过调用接口 [DescribeImages](../镜像相关接口/DescribeImages) ，取返回信息中的`ImageId`字段。</li>
func (c *Client) InquiryPriceResetInstance(request *InquiryPriceResetInstanceRequest) (response *InquiryPriceResetInstanceResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResetInstanceRequest()
	}
	response = NewInquiryPriceResetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageSnapshotStatusRequest() (request *DescribeImageSnapshotStatusRequest) {
	request = &DescribeImageSnapshotStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageSnapshotStatus")
	return
}

func NewDescribeImageSnapshotStatusResponse() (response *DescribeImageSnapshotStatusResponse) {
	response = &DescribeImageSnapshotStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年5月11日起停止维护。
// 本接口(DescribeImageSnapShotStatus)用于查询镜像是否存在快照
func (c *Client) DescribeImageSnapshotStatus(request *DescribeImageSnapshotStatusRequest) (response *DescribeImageSnapshotStatusResponse, err error) {
	if request == nil {
		request = NewDescribeImageSnapshotStatusRequest()
	}
	response = NewDescribeImageSnapshotStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountAttributesRequest() (request *DescribeAccountAttributesRequest) {
	request = &DescribeAccountAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAccountAttributes")
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

func NewQueryInstancesActionTimerRequest() (request *QueryInstancesActionTimerRequest) {
	request = &QueryInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "QueryInstancesActionTimer")
	return
}

func NewQueryInstancesActionTimerResponse() (response *QueryInstancesActionTimerResponse) {
	response = &QueryInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询定时销毁实例任务接口
func (c *Client) QueryInstancesActionTimer(request *QueryInstancesActionTimerRequest) (response *QueryInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewQueryInstancesActionTimerRequest()
	}
	response = NewQueryInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesActionTimerRequest() (request *DescribeInstancesActionTimerRequest) {
	request = &DescribeInstancesActionTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesActionTimer")
	return
}

func NewDescribeInstancesActionTimerResponse() (response *DescribeInstancesActionTimerResponse) {
	response = &DescribeInstancesActionTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询定时任务信息
func (c *Client) DescribeInstancesActionTimer(request *DescribeInstancesActionTimerRequest) (response *DescribeInstancesActionTimerResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesActionTimerRequest()
	}
	response = NewDescribeInstancesActionTimerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkSharingGroupsRequest() (request *DescribeNetworkSharingGroupsRequest) {
	request = &DescribeNetworkSharingGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeNetworkSharingGroups")
	return
}

func NewDescribeNetworkSharingGroupsResponse() (response *DescribeNetworkSharingGroupsResponse) {
	response = &DescribeNetworkSharingGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// None
func (c *Client) DescribeNetworkSharingGroups(request *DescribeNetworkSharingGroupsRequest) (response *DescribeNetworkSharingGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkSharingGroupsRequest()
	}
	response = NewDescribeNetworkSharingGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecomendedZonesRequest() (request *DescribeRecomendedZonesRequest) {
	request = &DescribeRecomendedZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeRecomendedZones")
	return
}

func NewDescribeRecomendedZonesResponse() (response *DescribeRecomendedZonesResponse) {
	response = &DescribeRecomendedZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用区规格配额排序列表
func (c *Client) DescribeRecomendedZones(request *DescribeRecomendedZonesRequest) (response *DescribeRecomendedZonesResponse, err error) {
	if request == nil {
		request = NewDescribeRecomendedZonesRequest()
	}
	response = NewDescribeRecomendedZonesResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterResetInstancesInternetMaxBandwidthRequest() (request *SwitchParameterResetInstancesInternetMaxBandwidthRequest) {
	request = &SwitchParameterResetInstancesInternetMaxBandwidthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterResetInstancesInternetMaxBandwidth")
	return
}

func NewSwitchParameterResetInstancesInternetMaxBandwidthResponse() (response *SwitchParameterResetInstancesInternetMaxBandwidthResponse) {
	response = &SwitchParameterResetInstancesInternetMaxBandwidthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterResetInstancesInternetMaxBandwidth) 用于调整实例公网带宽上限参数转换。
//
// * 不同机型带宽上限范围不一致。
// * 对于 `BANDWIDTH_PREPAID` 计费方式的带宽，需要输入参数 `StartTime` 和 `EndTime` ，指定调整后的带宽的生效时间段。在这种场景下目前不支持调小带宽，会涉及扣费，请确保账户余额充足。
// * 对于 `TRAFFIC_POSTPAID_BY_HOUR` 、 `BANDWIDTH_POSTPAID_BY_HOUR` 和 `BANDWIDTH_PACKAGE` 计费方式的带宽，使用该接口调整带宽上限是实时生效的，可以在带宽允许的范围内调大或者调小带宽，不支持输入参数 `StartTime` 和 `EndTime` 。
// * 接口不支持调整 `BANDWIDTH_POSTPAID_BY_MONTH` 计费方式的带宽。
// * 接口不支持批量调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽。
// * 接口不支持批量调整混合计费方式的带宽。例如不支持同时调整 `TRAFFIC_POSTPAID_BY_HOUR` 和 `BANDWIDTH_PACKAGE` 计费方式的带宽。
func (c *Client) SwitchParameterResetInstancesInternetMaxBandwidth(request *SwitchParameterResetInstancesInternetMaxBandwidthRequest) (response *SwitchParameterResetInstancesInternetMaxBandwidthResponse, err error) {
	if request == nil {
		request = NewSwitchParameterResetInstancesInternetMaxBandwidthRequest()
	}
	response = NewSwitchParameterResetInstancesInternetMaxBandwidthResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceTerminateInstancesRequest() (request *InquiryPriceTerminateInstancesRequest) {
	request = &InquiryPriceTerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceTerminateInstances")
	return
}

func NewInquiryPriceTerminateInstancesResponse() (response *InquiryPriceTerminateInstancesResponse) {
	response = &InquiryPriceTerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (InquiryPriceTerminateInstances) 用于退还实例询价。
//
// * 查询退还实例可以返还的费用。
// * 支持批量操作，每次请求批量实例的上限为100。如果批量实例存在不允许操作的实例，操作会以特定错误码返回。
func (c *Client) InquiryPriceTerminateInstances(request *InquiryPriceTerminateInstancesRequest) (response *InquiryPriceTerminateInstancesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceTerminateInstancesRequest()
	}
	response = NewInquiryPriceTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewImportImageRequest() (request *ImportImageRequest) {
	request = &ImportImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ImportImage")
	return
}

func NewImportImageResponse() (response *ImportImageResponse) {
	response = &ImportImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ImportImage)用于导入镜像，导入后的镜像可用于创建实例。
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
	if request == nil {
		request = NewImportImageRequest()
	}
	response = NewImportImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyKeyPairAttributeRequest() (request *ModifyKeyPairAttributeRequest) {
	request = &ModifyKeyPairAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyKeyPairAttribute")
	return
}

func NewModifyKeyPairAttributeResponse() (response *ModifyKeyPairAttributeResponse) {
	response = &ModifyKeyPairAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyKeyPairAttribute) 用于修改密钥对属性。
//
// * 修改密钥对ID所指定的密钥对的名称和描述信息。
// * 密钥对名称不能和已经存在的密钥对的名称重复。
// * 密钥对ID是密钥对的唯一标识，不可修改。
func (c *Client) ModifyKeyPairAttribute(request *ModifyKeyPairAttributeRequest) (response *ModifyKeyPairAttributeResponse, err error) {
	if request == nil {
		request = NewModifyKeyPairAttributeRequest()
	}
	response = NewModifyKeyPairAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。
//
// * 可以根据实例`ID`、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的实例。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaunchTemplatesInfoRequest() (request *DescribeLaunchTemplatesInfoRequest) {
	request = &DescribeLaunchTemplatesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeLaunchTemplatesInfo")
	return
}

func NewDescribeLaunchTemplatesInfoResponse() (response *DescribeLaunchTemplatesInfoResponse) {
	response = &DescribeLaunchTemplatesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeLaunchTemplatesInfo）用于查询一个或者多个实例启动模板详细信息。
func (c *Client) DescribeLaunchTemplatesInfo(request *DescribeLaunchTemplatesInfoRequest) (response *DescribeLaunchTemplatesInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLaunchTemplatesInfoRequest()
	}
	response = NewDescribeLaunchTemplatesInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateImage)用于将实例的系统盘制作为新镜像，创建后的镜像可以用于创建实例。
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	if request == nil {
		request = NewCreateImageRequest()
	}
	response = NewCreateImageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesAttributeRequest() (request *DescribeInstancesAttributeRequest) {
	request = &DescribeInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesAttribute")
	return
}

func NewDescribeInstancesAttributeResponse() (response *DescribeInstancesAttributeResponse) {
	response = &DescribeInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取实例订单Id。
func (c *Client) DescribeInstancesAttribute(request *DescribeInstancesAttributeRequest) (response *DescribeInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesAttributeRequest()
	}
	response = NewDescribeInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
	request = &ResetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResetInstance")
	return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
	response = &ResetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstance) 用于重装指定实例上的操作系统。
//
// * 如果指定了`ImageId`参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装。
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
// * `Linux`和`Windows`系统互相切换时，该实例系统盘`ID`将发生变化，系统盘关联快照将无法回滚、恢复数据。
// * 密码不指定将会通过站内信下发随机密码。
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口实现`Linux`和`Windows`操作系统切换。
// * 目前不支持海外地域的实例使用该接口实现`Linux`和`Windows`操作系统切换。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	if request == nil {
		request = NewResetInstanceRequest()
	}
	response = NewResetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateHostsRequest() (request *TerminateHostsRequest) {
	request = &TerminateHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "TerminateHosts")
	return
}

func NewTerminateHostsResponse() (response *TerminateHostsResponse) {
	response = &TerminateHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于主动退还CDH
func (c *Client) TerminateHosts(request *TerminateHostsRequest) (response *TerminateHostsResponse, err error) {
	if request == nil {
		request = NewTerminateHostsRequest()
	}
	response = NewTerminateHostsResponse()
	err = c.Send(request, response)
	return
}

func NewExitRescueModeRequest() (request *ExitRescueModeRequest) {
	request = &ExitRescueModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ExitRescueMode")
	return
}

func NewExitRescueModeResponse() (response *ExitRescueModeResponse) {
	response = &ExitRescueModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退出救援模式。
func (c *Client) ExitRescueMode(request *ExitRescueModeRequest) (response *ExitRescueModeResponse, err error) {
	if request == nil {
		request = NewExitRescueModeRequest()
	}
	response = NewExitRescueModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisasterRecoverGroupsRequest() (request *DescribeDisasterRecoverGroupsRequest) {
	request = &DescribeDisasterRecoverGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroups")
	return
}

func NewDescribeDisasterRecoverGroupsResponse() (response *DescribeDisasterRecoverGroupsResponse) {
	response = &DescribeDisasterRecoverGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分散置放群组信息
func (c *Client) DescribeDisasterRecoverGroups(request *DescribeDisasterRecoverGroupsRequest) (response *DescribeDisasterRecoverGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoverGroupsRequest()
	}
	response = NewDescribeDisasterRecoverGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesOverviewRequest() (request *DescribeResourcesOverviewRequest) {
	request = &DescribeResourcesOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeResourcesOverview")
	return
}

func NewDescribeResourcesOverviewResponse() (response *DescribeResourcesOverviewResponse) {
	response = &DescribeResourcesOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询概览页详情
func (c *Client) DescribeResourcesOverview(request *DescribeResourcesOverviewRequest) (response *DescribeResourcesOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesOverviewRequest()
	}
	response = NewDescribeResourcesOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaunchTemplatesRequest() (request *DeleteLaunchTemplatesRequest) {
	request = &DeleteLaunchTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteLaunchTemplates")
	return
}

func NewDeleteLaunchTemplatesResponse() (response *DeleteLaunchTemplatesResponse) {
	response = &DeleteLaunchTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteLaunchTemplates）用于删除一个或多个实例启动模板。
func (c *Client) DeleteLaunchTemplates(request *DeleteLaunchTemplatesRequest) (response *DeleteLaunchTemplatesResponse, err error) {
	if request == nil {
		request = NewDeleteLaunchTemplatesRequest()
	}
	response = NewDeleteLaunchTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesOperationLimitRequest() (request *DescribeInstancesOperationLimitRequest) {
	request = &DescribeInstancesOperationLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesOperationLimit")
	return
}

func NewDescribeInstancesOperationLimitResponse() (response *DescribeInstancesOperationLimitResponse) {
	response = &DescribeInstancesOperationLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例一些受限操作的可执行次数
func (c *Client) DescribeInstancesOperationLimit(request *DescribeInstancesOperationLimitRequest) (response *DescribeInstancesOperationLimitResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesOperationLimitRequest()
	}
	response = NewDescribeInstancesOperationLimitResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
	request = &AssociateSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "AssociateSecurityGroups")
	return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
	response = &AssociateSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (AssociateSecurityGroups) 用于绑定安全组到指定实例。
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewAssociateSecurityGroupsRequest()
	}
	response = NewAssociateSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryResourceResetInstancesTypeRequest() (request *InquiryResourceResetInstancesTypeRequest) {
	request = &InquiryResourceResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryResourceResetInstancesType")
	return
}

func NewInquiryResourceResetInstancesTypeResponse() (response *InquiryResourceResetInstancesTypeResponse) {
	response = &InquiryResourceResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryResourceResetInstancesType）用于调整实例的资源查询
func (c *Client) InquiryResourceResetInstancesType(request *InquiryResourceResetInstancesTypeRequest) (response *InquiryResourceResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewInquiryResourceResetInstancesTypeRequest()
	}
	response = NewInquiryResourceResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
	request = &DescribeInstanceTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeConfigs")
	return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
	response = &DescribeInstanceTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceTypeConfigs) 用于查询实例机型配置。
//
// * 可以根据`zone`、`instance-family`来查询实例机型配置。过滤条件详见过滤器`Filter`。
// * 如果参数为空，返回指定地域的所有实例机型配置。
func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeConfigsRequest()
	}
	response = NewDescribeInstanceTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterResetInstanceRequest() (request *SwitchParameterResetInstanceRequest) {
	request = &SwitchParameterResetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "SwitchParameterResetInstance")
	return
}

func NewSwitchParameterResetInstanceResponse() (response *SwitchParameterResetInstanceResponse) {
	response = &SwitchParameterResetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口自2022年4月14日起停止维护。
// 本接口 (SwitchParameterResetInstance) 用于重装指定实例上的操作系统参数转换。
//
// * 如果指定了`ImageId`参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装。
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
// * `Linux`和`Windows`系统互相切换时，该实例系统盘`ID`将发生变化，系统盘关联快照将无法回滚、恢复数据。
// * 密码不指定将会通过站内信下发随机密码。
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口实现`Linux`和`Windows`操作系统切换。
// * 目前不支持海外地域的实例使用该接口实现`Linux`和`Windows`操作系统切换。
func (c *Client) SwitchParameterResetInstance(request *SwitchParameterResetInstanceRequest) (response *SwitchParameterResetInstanceResponse, err error) {
	if request == nil {
		request = NewSwitchParameterResetInstanceRequest()
	}
	response = NewSwitchParameterResetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewReturnAddressesRequest() (request *ReturnAddressesRequest) {
	request = &ReturnAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ReturnAddresses")
	return
}

func NewReturnAddressesResponse() (response *ReturnAddressesResponse) {
	response = &ReturnAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) ReturnAddresses(request *ReturnAddressesRequest) (response *ReturnAddressesResponse, err error) {
	if request == nil {
		request = NewReturnAddressesRequest()
	}
	response = NewReturnAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewCloneInstanceRequest() (request *CloneInstanceRequest) {
	request = &CloneInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CloneInstance")
	return
}

func NewCloneInstanceResponse() (response *CloneInstanceResponse) {
	response = &CloneInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CloneInstance) 用于克隆指定配置的实例。 * 实例创建成功后将自动开机启动，[实例状态](../数据结构#instancestatus)变为“运行中”。 * 预付费实例的购买会预先扣除本次实例购买所需金额，按小时后付费实例购买会预先冻结本次实例购买一小时内所需金额，在调用本接口前请确保账户余额充足。 * 本接口为异步接口，当创建请求下发成功后会返回一个实例`ID`列表，此时实例的创建并立即未完成。在此期间实例的状态将会处于”准备中“，可以通过调用 [DescribeInstancesStatus](DescribeInstancesStatus) 接口查询对应实例的状态，来判断生产有没有最终成功。如果实例的状态由'准备中'变为'运行中'，则为创建成功。
func (c *Client) CloneInstance(request *CloneInstanceRequest) (response *CloneInstanceResponse, err error) {
	if request == nil {
		request = NewCloneInstanceRequest()
	}
	response = NewCloneInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewConvertOperatingSystemsRequest() (request *ConvertOperatingSystemsRequest) {
	request = &ConvertOperatingSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ConvertOperatingSystems")
	return
}

func NewConvertOperatingSystemsResponse() (response *ConvertOperatingSystemsResponse) {
	response = &ConvertOperatingSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(ConvertOperatingSystems)用于转换实例的操作系统，仅支持源操作系统为 CentOS 7、CentOS 8 的实例。
func (c *Client) ConvertOperatingSystems(request *ConvertOperatingSystemsRequest) (response *ConvertOperatingSystemsResponse, err error) {
	if request == nil {
		request = NewConvertOperatingSystemsRequest()
	}
	response = NewConvertOperatingSystemsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHpcClusterRequest() (request *CreateHpcClusterRequest) {
	request = &CreateHpcClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateHpcCluster")
	return
}

func NewCreateHpcClusterResponse() (response *CreateHpcClusterResponse) {
	response = &CreateHpcClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建高性能计算集群
func (c *Client) CreateHpcCluster(request *CreateHpcClusterRequest) (response *CreateHpcClusterResponse, err error) {
	if request == nil {
		request = NewCreateHpcClusterRequest()
	}
	response = NewCreateHpcClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLaunchTemplateVersionRequest() (request *CreateLaunchTemplateVersionRequest) {
	request = &CreateLaunchTemplateVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "CreateLaunchTemplateVersion")
	return
}

func NewCreateLaunchTemplateVersionResponse() (response *CreateLaunchTemplateVersionResponse) {
	response = &CreateLaunchTemplateVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateLaunchTemplateVersion）根据指定的实例模板ID以及对应的模板版本号创建新的实例启动模板，若未指定模板版本号则使用默认版本号。每个实例启动模板最多创建30个版本。 * 新实例模板中未显式指定的参数值，使用指定版本号对应参数值覆盖。
func (c *Client) CreateLaunchTemplateVersion(request *CreateLaunchTemplateVersionRequest) (response *CreateLaunchTemplateVersionResponse, err error) {
	if request == nil {
		request = NewCreateLaunchTemplateVersionRequest()
	}
	response = NewCreateLaunchTemplateVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHpcClustersRequest() (request *DeleteHpcClustersRequest) {
	request = &DeleteHpcClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteHpcClusters")
	return
}

func NewDeleteHpcClustersResponse() (response *DeleteHpcClustersResponse) {
	response = &DeleteHpcClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高性能计算集群
func (c *Client) DeleteHpcClusters(request *DeleteHpcClustersRequest) (response *DeleteHpcClustersResponse, err error) {
	if request == nil {
		request = NewDeleteHpcClustersRequest()
	}
	response = NewDeleteHpcClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaunchTemplateVersionsRequest() (request *DeleteLaunchTemplateVersionsRequest) {
	request = &DeleteLaunchTemplateVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DeleteLaunchTemplateVersions")
	return
}

func NewDeleteLaunchTemplateVersionsResponse() (response *DeleteLaunchTemplateVersionsResponse) {
	response = &DeleteLaunchTemplateVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteLaunchTemplateVersions）用于删除一个或者多个实例启动模板版本。
func (c *Client) DeleteLaunchTemplateVersions(request *DeleteLaunchTemplateVersionsRequest) (response *DeleteLaunchTemplateVersionsResponse, err error) {
	if request == nil {
		request = NewDeleteLaunchTemplateVersionsRequest()
	}
	response = NewDeleteLaunchTemplateVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountQuotaRequest() (request *DescribeAccountQuotaRequest) {
	request = &DescribeAccountQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAccountQuota")
	return
}

func NewDescribeAccountQuotaResponse() (response *DescribeAccountQuotaResponse) {
	response = &DescribeAccountQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeAccountQuota)用于查询用户配额详情。可以查询后付费、预付费、竞价实例、镜像、置放群组等各种类型的配额信息。
func (c *Client) DescribeAccountQuota(request *DescribeAccountQuotaRequest) (response *DescribeAccountQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAccountQuotaRequest()
	}
	response = NewDescribeAccountQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableFeaturesRequest() (request *DescribeAvailableFeaturesRequest) {
	request = &DescribeAvailableFeaturesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeAvailableFeatures")
	return
}

func NewDescribeAvailableFeaturesResponse() (response *DescribeAvailableFeaturesResponse) {
	response = &DescribeAvailableFeaturesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置的功能特性是否可用
func (c *Client) DescribeAvailableFeatures(request *DescribeAvailableFeaturesRequest) (response *DescribeAvailableFeaturesResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableFeaturesRequest()
	}
	response = NewDescribeAvailableFeaturesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostCapacityRequest() (request *DescribeHostCapacityRequest) {
	request = &DescribeHostCapacityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeHostCapacity")
	return
}

func NewDescribeHostCapacityResponse() (response *DescribeHostCapacityResponse) {
	response = &DescribeHostCapacityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口是用于查询最大一台或者多台cdh宿主机的创建指定规格的子机数量。如果是系统盘的话，必须传image id
func (c *Client) DescribeHostCapacity(request *DescribeHostCapacityRequest) (response *DescribeHostCapacityResponse, err error) {
	if request == nil {
		request = NewDescribeHostCapacityRequest()
	}
	response = NewDescribeHostCapacityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHpcClustersRequest() (request *DescribeHpcClustersRequest) {
	request = &DescribeHpcClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeHpcClusters")
	return
}

func NewDescribeHpcClustersResponse() (response *DescribeHpcClustersResponse) {
	response = &DescribeHpcClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询高性能集群信息
func (c *Client) DescribeHpcClusters(request *DescribeHpcClustersRequest) (response *DescribeHpcClustersResponse, err error) {
	if request == nil {
		request = NewDescribeHpcClustersRequest()
	}
	response = NewDescribeHpcClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageStatisticsRequest() (request *DescribeImageStatisticsRequest) {
	request = &DescribeImageStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageStatistics")
	return
}

func NewDescribeImageStatisticsResponse() (response *DescribeImageStatisticsResponse) {
	response = &DescribeImageStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeImageStatistics) 用于查询各个地域镜像数量概览。 本接口仅支持查询自定义镜像和共享镜像。
func (c *Client) DescribeImageStatistics(request *DescribeImageStatisticsRequest) (response *DescribeImageStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeImageStatisticsRequest()
	}
	response = NewDescribeImageStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceAttributesRequest() (request *DescribeInstanceAttributesRequest) {
	request = &DescribeInstanceAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceAttributes")
	return
}

func NewDescribeInstanceAttributesResponse() (response *DescribeInstanceAttributesResponse) {
	response = &DescribeInstanceAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定实例的属性，目前支持查询实例自定义数据User-Data。
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceAttributesRequest()
	}
	response = NewDescribeInstanceAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeDisasterGroupBlackListRequest() (request *DescribeInstanceTypeDisasterGroupBlackListRequest) {
	request = &DescribeInstanceTypeDisasterGroupBlackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeDisasterGroupBlackList")
	return
}

func NewDescribeInstanceTypeDisasterGroupBlackListResponse() (response *DescribeInstanceTypeDisasterGroupBlackListResponse) {
	response = &DescribeInstanceTypeDisasterGroupBlackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询不支持置放群组的机型列表
func (c *Client) DescribeInstanceTypeDisasterGroupBlackList(request *DescribeInstanceTypeDisasterGroupBlackListRequest) (response *DescribeInstanceTypeDisasterGroupBlackListResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeDisasterGroupBlackListRequest()
	}
	response = NewDescribeInstanceTypeDisasterGroupBlackListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeZoneStatusRequest() (request *DescribeInstanceTypeZoneStatusRequest) {
	request = &DescribeInstanceTypeZoneStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeZoneStatus")
	return
}

func NewDescribeInstanceTypeZoneStatusResponse() (response *DescribeInstanceTypeZoneStatusResponse) {
	response = &DescribeInstanceTypeZoneStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机型相关可用区售卖状态
func (c *Client) DescribeInstanceTypeZoneStatus(request *DescribeInstanceTypeZoneStatusRequest) (response *DescribeInstanceTypeZoneStatusResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeZoneStatusRequest()
	}
	response = NewDescribeInstanceTypeZoneStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesConvertOSAttributesRequest() (request *DescribeInstancesConvertOSAttributesRequest) {
	request = &DescribeInstancesConvertOSAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesConvertOSAttributes")
	return
}

func NewDescribeInstancesConvertOSAttributesResponse() (response *DescribeInstancesConvertOSAttributesResponse) {
	response = &DescribeInstancesConvertOSAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作系统转换的属性接口。
func (c *Client) DescribeInstancesConvertOSAttributes(request *DescribeInstancesConvertOSAttributesRequest) (response *DescribeInstancesConvertOSAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesConvertOSAttributesRequest()
	}
	response = NewDescribeInstancesConvertOSAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaunchTemplateVersionsRequest() (request *DescribeLaunchTemplateVersionsRequest) {
	request = &DescribeLaunchTemplateVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeLaunchTemplateVersions")
	return
}

func NewDescribeLaunchTemplateVersionsResponse() (response *DescribeLaunchTemplateVersionsResponse) {
	response = &DescribeLaunchTemplateVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeLaunchTemplateVersions）用于查询实例模板版本信息。
func (c *Client) DescribeLaunchTemplateVersions(request *DescribeLaunchTemplateVersionsRequest) (response *DescribeLaunchTemplateVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeLaunchTemplateVersionsRequest()
	}
	response = NewDescribeLaunchTemplateVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaunchTemplatesRequest() (request *DescribeLaunchTemplatesRequest) {
	request = &DescribeLaunchTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeLaunchTemplates")
	return
}

func NewDescribeLaunchTemplatesResponse() (response *DescribeLaunchTemplatesResponse) {
	response = &DescribeLaunchTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeLaunchTemplates）用于查询一个或者多个实例启动模板。
func (c *Client) DescribeLaunchTemplates(request *DescribeLaunchTemplatesRequest) (response *DescribeLaunchTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeLaunchTemplatesRequest()
	}
	response = NewDescribeLaunchTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAvailableInstanceTypesRequest() (request *DescribeUserAvailableInstanceTypesRequest) {
	request = &DescribeUserAvailableInstanceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserAvailableInstanceTypes")
	return
}

func NewDescribeUserAvailableInstanceTypesResponse() (response *DescribeUserAvailableInstanceTypesResponse) {
	response = &DescribeUserAvailableInstanceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口仅针对购买页使用，可以展示入参用户的可用机型
func (c *Client) DescribeUserAvailableInstanceTypes(request *DescribeUserAvailableInstanceTypesRequest) (response *DescribeUserAvailableInstanceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeUserAvailableInstanceTypesRequest()
	}
	response = NewDescribeUserAvailableInstanceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAvailableRegionAndZonesRequest() (request *DescribeUserAvailableRegionAndZonesRequest) {
	request = &DescribeUserAvailableRegionAndZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserAvailableRegionAndZones")
	return
}

func NewDescribeUserAvailableRegionAndZonesResponse() (response *DescribeUserAvailableRegionAndZonesResponse) {
	response = &DescribeUserAvailableRegionAndZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// MC 控制台查询用户可用的地域售罄状况等信息展示
func (c *Client) DescribeUserAvailableRegionAndZones(request *DescribeUserAvailableRegionAndZonesRequest) (response *DescribeUserAvailableRegionAndZonesResponse, err error) {
	if request == nil {
		request = NewDescribeUserAvailableRegionAndZonesRequest()
	}
	response = NewDescribeUserAvailableRegionAndZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstancesDiscountInfoRequest() (request *DescribeUserInstancesDiscountInfoRequest) {
	request = &DescribeUserInstancesDiscountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserInstancesDiscountInfo")
	return
}

func NewDescribeUserInstancesDiscountInfoResponse() (response *DescribeUserInstancesDiscountInfoResponse) {
	response = &DescribeUserInstancesDiscountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户当前实例折扣，可以查询到如续费等折扣详情和活动时间范围等。
func (c *Client) DescribeUserInstancesDiscountInfo(request *DescribeUserInstancesDiscountInfoRequest) (response *DescribeUserInstancesDiscountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstancesDiscountInfoRequest()
	}
	response = NewDescribeUserInstancesDiscountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLoginAttributeRequest() (request *DescribeUserLoginAttributeRequest) {
	request = &DescribeUserLoginAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeUserLoginAttribute")
	return
}

func NewDescribeUserLoginAttributeResponse() (response *DescribeUserLoginAttributeResponse) {
	response = &DescribeUserLoginAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询用户登录属性
func (c *Client) DescribeUserLoginAttribute(request *DescribeUserLoginAttributeRequest) (response *DescribeUserLoginAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeUserLoginAttributeRequest()
	}
	response = NewDescribeUserLoginAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneInstanceSoldDiscountRequest() (request *DescribeZoneInstanceSoldDiscountRequest) {
	request = &DescribeZoneInstanceSoldDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneInstanceSoldDiscount")
	return
}

func NewDescribeZoneInstanceSoldDiscountResponse() (response *DescribeZoneInstanceSoldDiscountResponse) {
	response = &DescribeZoneInstanceSoldDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据计费模式、机型、用户uin、可用区查询机型折扣。同一实例族的实例类型的折扣相同，比如S2机型族的所有机型的折扣都是相同的。
func (c *Client) DescribeZoneInstanceSoldDiscount(request *DescribeZoneInstanceSoldDiscountRequest) (response *DescribeZoneInstanceSoldDiscountResponse, err error) {
	if request == nil {
		request = NewDescribeZoneInstanceSoldDiscountRequest()
	}
	response = NewDescribeZoneInstanceSoldDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewExportImagesRequest() (request *ExportImagesRequest) {
	request = &ExportImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ExportImages")
	return
}

func NewExportImagesResponse() (response *ExportImagesResponse) {
	response = &ExportImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供导出自定义镜像到指定COS存储桶的能力
func (c *Client) ExportImages(request *ExportImagesRequest) (response *ExportImagesResponse, err error) {
	if request == nil {
		request = NewExportImagesRequest()
	}
	response = NewExportImagesResponse()
	err = c.Send(request, response)
	return
}

func NewInstanceAttachISORequest() (request *InstanceAttachISORequest) {
	request = &InstanceAttachISORequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InstanceAttachISO")
	return
}

func NewInstanceAttachISOResponse() (response *InstanceAttachISOResponse) {
	response = &InstanceAttachISOResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例挂载ISO镜像
func (c *Client) InstanceAttachISO(request *InstanceAttachISORequest) (response *InstanceAttachISOResponse, err error) {
	if request == nil {
		request = NewInstanceAttachISORequest()
	}
	response = NewInstanceAttachISOResponse()
	err = c.Send(request, response)
	return
}

func NewInstanceDetachISORequest() (request *InstanceDetachISORequest) {
	request = &InstanceDetachISORequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InstanceDetachISO")
	return
}

func NewInstanceDetachISOResponse() (response *InstanceDetachISOResponse) {
	response = &InstanceDetachISOResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例解挂ISO镜像盘
func (c *Client) InstanceDetachISO(request *InstanceDetachISORequest) (response *InstanceDetachISOResponse, err error) {
	if request == nil {
		request = NewInstanceDetachISORequest()
	}
	response = NewInstanceDetachISOResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHpcClusterAttributeRequest() (request *ModifyHpcClusterAttributeRequest) {
	request = &ModifyHpcClusterAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyHpcClusterAttribute")
	return
}

func NewModifyHpcClusterAttributeResponse() (response *ModifyHpcClusterAttributeResponse) {
	response = &ModifyHpcClusterAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改高性能计算集群属性
func (c *Client) ModifyHpcClusterAttribute(request *ModifyHpcClusterAttributeRequest) (response *ModifyHpcClusterAttributeResponse, err error) {
	if request == nil {
		request = NewModifyHpcClusterAttributeRequest()
	}
	response = NewModifyHpcClusterAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLaunchTemplateDefaultVersionRequest() (request *ModifyLaunchTemplateDefaultVersionRequest) {
	request = &ModifyLaunchTemplateDefaultVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyLaunchTemplateDefaultVersion")
	return
}

func NewModifyLaunchTemplateDefaultVersionResponse() (response *ModifyLaunchTemplateDefaultVersionResponse) {
	response = &ModifyLaunchTemplateDefaultVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyLaunchTemplateDefaultVersion）用于修改实例启动模板默认版本。
func (c *Client) ModifyLaunchTemplateDefaultVersion(request *ModifyLaunchTemplateDefaultVersionRequest) (response *ModifyLaunchTemplateDefaultVersionResponse, err error) {
	if request == nil {
		request = NewModifyLaunchTemplateDefaultVersionRequest()
	}
	response = NewModifyLaunchTemplateDefaultVersionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserLoginAttributeRequest() (request *ModifyUserLoginAttributeRequest) {
	request = &ModifyUserLoginAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ModifyUserLoginAttribute")
	return
}

func NewModifyUserLoginAttributeResponse() (response *ModifyUserLoginAttributeResponse) {
	response = &ModifyUserLoginAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于用户登录属性
func (c *Client) ModifyUserLoginAttribute(request *ModifyUserLoginAttributeRequest) (response *ModifyUserLoginAttributeResponse, err error) {
	if request == nil {
		request = NewModifyUserLoginAttributeRequest()
	}
	response = NewModifyUserLoginAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveInstancesDisasterRecoverGroupRequest() (request *RemoveInstancesDisasterRecoverGroupRequest) {
	request = &RemoveInstancesDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "RemoveInstancesDisasterRecoverGroup")
	return
}

func NewRemoveInstancesDisasterRecoverGroupResponse() (response *RemoveInstancesDisasterRecoverGroupResponse) {
	response = &RemoveInstancesDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例解绑置放群组
func (c *Client) RemoveInstancesDisasterRecoverGroup(request *RemoveInstancesDisasterRecoverGroupRequest) (response *RemoveInstancesDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewRemoveInstancesDisasterRecoverGroupRequest()
	}
	response = NewRemoveInstancesDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewViewModifyInstancesAttributeRequest() (request *ViewModifyInstancesAttributeRequest) {
	request = &ViewModifyInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ViewModifyInstancesAttribute")
	return
}

func NewViewModifyInstancesAttributeResponse() (response *ViewModifyInstancesAttributeResponse) {
	response = &ViewModifyInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。 * “实例名称”仅为方便用户自己管理之用，云平台并不以此名称作为提交工单或是进行实例管理操作的依据。 * 支持批量操作。每次请求批量实例的上限为100。
func (c *Client) ViewModifyInstancesAttribute(request *ViewModifyInstancesAttributeRequest) (response *ViewModifyInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewViewModifyInstancesAttributeRequest()
	}
	response = NewViewModifyInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

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

func NewInquiryPriceSnapshotsRequest() (request *InquiryPriceSnapshotsRequest) {
	request = &InquiryPriceSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceSnapshots")
	return
}

func NewInquiryPriceSnapshotsResponse() (response *InquiryPriceSnapshotsResponse) {
	response = &InquiryPriceSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 快照询价接口，指定快照大小，查询一个小时的快照费用。
func (c *Client) InquiryPriceSnapshots(request *InquiryPriceSnapshotsRequest) (response *InquiryPriceSnapshotsResponse, err error) {
	if request == nil {
		request = NewInquiryPriceSnapshotsRequest()
	}
	response = NewInquiryPriceSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceResizeDiskRequest() (request *InquiryPriceResizeDiskRequest) {
	request = &InquiryPriceResizeDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceResizeDisk")
	return
}

func NewInquiryPriceResizeDiskResponse() (response *InquiryPriceResizeDiskResponse) {
	response = &InquiryPriceResizeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceResizeDisk）用于扩容云硬盘询价。
//
// * 只支持预付费模式的云硬盘扩容询价。
func (c *Client) InquiryPriceResizeDisk(request *InquiryPriceResizeDiskRequest) (response *InquiryPriceResizeDiskResponse, err error) {
	if request == nil {
		request = NewInquiryPriceResizeDiskRequest()
	}
	response = NewInquiryPriceResizeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskAssociatedSnapshotsRequest() (request *DescribeDiskAssociatedSnapshotsRequest) {
	request = &DescribeDiskAssociatedSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedSnapshots")
	return
}

func NewDescribeDiskAssociatedSnapshotsResponse() (response *DescribeDiskAssociatedSnapshotsResponse) {
	response = &DescribeDiskAssociatedSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskAssociatedSnapshots）用于查询云盘关联的快照。
func (c *Client) DescribeDiskAssociatedSnapshots(request *DescribeDiskAssociatedSnapshotsRequest) (response *DescribeDiskAssociatedSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskAssociatedSnapshotsRequest()
	}
	response = NewDescribeDiskAssociatedSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ResizeDisk")
	return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResizeDisk）用于扩容云硬盘。
//
// * 只支持扩容弹性云盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。随云主机创建的云硬盘需通过[ResizeInstanceDisks](/document/product/213/15731)接口扩容。
// * 本接口为异步接口，接口成功返回时，云盘并未立即扩容到指定大小，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态为“EXPANDING”，表示正在扩容中，当状态变为“UNATTACHED”，表示扩容完成。
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	if request == nil {
		request = NewResizeDiskRequest()
	}
	response = NewResizeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskAttributesRequest() (request *ModifyDiskAttributesRequest) {
	request = &ModifyDiskAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskAttributes")
	return
}

func NewModifyDiskAttributesResponse() (response *ModifyDiskAttributesResponse) {
	response = &ModifyDiskAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDiskAttributes）用于修改云硬盘属性。
//
// * 只支持修改弹性云盘的项目ID。随云主机创建的云硬盘项目ID与云主机联动。可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。
// * “云硬盘名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行云盘管理操作的依据。
// * 支持批量操作，如果传入多个云盘ID，则所有云盘修改为同一属性。如果存在不允许操作的云盘，则操作不执行，以特定错误码返回。
func (c *Client) ModifyDiskAttributes(request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
	if request == nil {
		request = NewModifyDiskAttributesRequest()
	}
	response = NewModifyDiskAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
	request = &CreateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDisks")
	return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
	response = &CreateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDisks）用于创建云硬盘。
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
	if request == nil {
		request = NewCreateDisksRequest()
	}
	response = NewCreateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshots")
	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotsRequest()
	}
	response = NewDescribeSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
	request = &UnbindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UnbindAutoSnapshotPolicy")
	return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
	response = &UnbindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindAutoSnapshotPolicy）用于解除云硬盘绑定的定期快照策略。
//
// * 支持批量操作，可一次解除多个云盘与同一定期快照策略的绑定。
// * 如果传入的云盘未绑定到当前定期快照策略，接口将自动跳过，仅解绑与当前定期快照策略绑定的云盘。
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewUnbindAutoSnapshotPolicyRequest()
	}
	response = NewUnbindAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDiskResourcesRequest() (request *DescribeUserDiskResourcesRequest) {
	request = &DescribeUserDiskResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeUserDiskResources")
	return
}

func NewDescribeUserDiskResourcesResponse() (response *DescribeUserDiskResourcesResponse) {
	response = &DescribeUserDiskResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeUserTotalResources）用于查询用户在所有地域的云盘数量、快照数量及配额。
func (c *Client) DescribeUserDiskResources(request *DescribeUserDiskResourcesRequest) (response *DescribeUserDiskResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeUserDiskResourcesRequest()
	}
	response = NewDescribeUserDiskResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAutoSnapshotPoliciesRequest() (request *DeleteAutoSnapshotPoliciesRequest) {
	request = &DeleteAutoSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteAutoSnapshotPolicies")
	return
}

func NewDeleteAutoSnapshotPoliciesResponse() (response *DeleteAutoSnapshotPoliciesResponse) {
	response = &DeleteAutoSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAutoSnapshotPolicies）用于删除定期快照策略。
//
// *  支持批量操作。如果多个定期快照策略存在无法删除的，则操作不执行，以特定错误码返回。
func (c *Client) DeleteAutoSnapshotPolicies(request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteAutoSnapshotPoliciesRequest()
	}
	response = NewDeleteAutoSnapshotPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesDiskNumRequest() (request *DescribeInstancesDiskNumRequest) {
	request = &DescribeInstancesDiskNumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeInstancesDiskNum")
	return
}

func NewDescribeInstancesDiskNumResponse() (response *DescribeInstancesDiskNumResponse) {
	response = &DescribeInstancesDiskNumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstancesDiskNum）用于查询实例已挂载云硬盘数量。
//
// * 支持批量操作，当传入多个云服务器实例ID，返回结果会分别列出每个云服务器挂载的云硬盘数量。
func (c *Client) DescribeInstancesDiskNum(request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesDiskNumRequest()
	}
	response = NewDescribeInstancesDiskNumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyRequest() (request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) {
	request = &DescribeDiskAssociatedAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedAutoSnapshotPolicy")
	return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyResponse() (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse) {
	response = &DescribeDiskAssociatedAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskAssociatedAutoSnapshotPolicy）用于查询云盘绑定的定期快照策略。
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicy(request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeDiskAssociatedAutoSnapshotPolicyRequest()
	}
	response = NewDescribeDiskAssociatedAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotGroupRequest() (request *CreateSnapshotGroupRequest) {
	request = &CreateSnapshotGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshotGroup")
	return
}

func NewCreateSnapshotGroupResponse() (response *CreateSnapshotGroupResponse) {
	response = &CreateSnapshotGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建快照组：
// * 创建快照组的盘列表必须挂载在同一实例上，可选择挂载在实例上的全部或部分盘创建快照组
//
//
func (c *Client) CreateSnapshotGroup(request *CreateSnapshotGroupRequest) (response *CreateSnapshotGroupResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotGroupRequest()
	}
	response = NewCreateSnapshotGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoSnapshotPolicyAttributeRequest() (request *ModifyAutoSnapshotPolicyAttributeRequest) {
	request = &ModifyAutoSnapshotPolicyAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyAutoSnapshotPolicyAttribute")
	return
}

func NewModifyAutoSnapshotPolicyAttributeResponse() (response *ModifyAutoSnapshotPolicyAttributeResponse) {
	response = &ModifyAutoSnapshotPolicyAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAutoSnapshotPolicyAttribute）用于修改定期快照策略属性。
//
// * 可通过该接口修改定期快照策略的执行策略、名称、是否激活等属性。
// * 修改保留天数时必须保证不与是否永久保留属性冲突，否则整个操作失败，以特定的错误码返回。
func (c *Client) ModifyAutoSnapshotPolicyAttribute(request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAutoSnapshotPolicyAttributeRequest()
	}
	response = NewModifyAutoSnapshotPolicyAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewCopySnapshotCrossRegionsRequest() (request *CopySnapshotCrossRegionsRequest) {
	request = &CopySnapshotCrossRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CopySnapshotCrossRegions")
	return
}

func NewCopySnapshotCrossRegionsResponse() (response *CopySnapshotCrossRegionsResponse) {
	response = &CopySnapshotCrossRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CopySnapshotCrossRegion）用于快照跨地域复制。
//
// * 本接口为异步接口，当跨地域复制的请求下发成功后会返回一个新的快照ID，此时快照未立即复制到目标地域，可请求目标地域的[DescribeSnapshots](/document/product/362/15647)接口新快照的状态，判断是否复制完成。如果快照的状态为“NORMAL”，表示快照复制完成。
func (c *Client) CopySnapshotCrossRegions(request *CopySnapshotCrossRegionsRequest) (response *CopySnapshotCrossRegionsResponse, err error) {
	if request == nil {
		request = NewCopySnapshotCrossRegionsRequest()
	}
	response = NewCopySnapshotCrossRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceRenewDisksRequest() (request *InquiryPriceRenewDisksRequest) {
	request = &InquiryPriceRenewDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceRenewDisks")
	return
}

func NewInquiryPriceRenewDisksResponse() (response *InquiryPriceRenewDisksResponse) {
	response = &InquiryPriceRenewDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceRenewDisks）用于续费云硬盘询价。
//
// * 只支持查询预付费模式的弹性云盘续费价格。
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到实例续费后的到期时间来续费询价。
// * 支持为多块云盘指定不同的续费时长，此时返回的价格为多块云盘续费的总价格。
func (c *Client) InquiryPriceRenewDisks(request *InquiryPriceRenewDisksRequest) (response *InquiryPriceRenewDisksResponse, err error) {
	if request == nil {
		request = NewInquiryPriceRenewDisksRequest()
	}
	response = NewInquiryPriceRenewDisksResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceModifyDiskExtraPerformanceRequest() (request *InquirePriceModifyDiskExtraPerformanceRequest) {
	request = &InquirePriceModifyDiskExtraPerformanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquirePriceModifyDiskExtraPerformance")
	return
}

func NewInquirePriceModifyDiskExtraPerformanceResponse() (response *InquirePriceModifyDiskExtraPerformanceResponse) {
	response = &InquirePriceModifyDiskExtraPerformanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquirePriceModifyDiskExtraPerformance）用于调整云硬盘额外性能询价。
func (c *Client) InquirePriceModifyDiskExtraPerformance(request *InquirePriceModifyDiskExtraPerformanceRequest) (response *InquirePriceModifyDiskExtraPerformanceResponse, err error) {
	if request == nil {
		request = NewInquirePriceModifyDiskExtraPerformanceRequest()
	}
	response = NewInquirePriceModifyDiskExtraPerformanceResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterCreateDisksRequest() (request *SwitchParameterCreateDisksRequest) {
	request = &SwitchParameterCreateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "SwitchParameterCreateDisks")
	return
}

func NewSwitchParameterCreateDisksResponse() (response *SwitchParameterCreateDisksResponse) {
	response = &SwitchParameterCreateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SwitchParameterCreateDisks）用于得到创建云硬盘的订单参数。
func (c *Client) SwitchParameterCreateDisks(request *SwitchParameterCreateDisksRequest) (response *SwitchParameterCreateDisksResponse, err error) {
	if request == nil {
		request = NewSwitchParameterCreateDisksRequest()
	}
	response = NewSwitchParameterCreateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewApplySnapshotGroupRequest() (request *ApplySnapshotGroupRequest) {
	request = &ApplySnapshotGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ApplySnapshotGroup")
	return
}

func NewApplySnapshotGroupResponse() (response *ApplySnapshotGroupResponse) {
	response = &ApplySnapshotGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于回滚快照组
// * 1.可选择快照组全部或部分盘进行回滚；
// * 2.如果回滚的盘中包含已挂载的盘，要求这些盘必须挂载在同一实例上，且要求该实例已关机才能回滚；
// * 3.回滚为异步操作，接口返回成功不代表回滚成功，可通过调DescribeSnapshotGroups接口查询异步回滚任务的结果
func (c *Client) ApplySnapshotGroup(request *ApplySnapshotGroupRequest) (response *ApplySnapshotGroupResponse, err error) {
	if request == nil {
		request = NewApplySnapshotGroupRequest()
	}
	response = NewApplySnapshotGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifySnapshotsSharePermissionRequest() (request *ModifySnapshotsSharePermissionRequest) {
	request = &ModifySnapshotsSharePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotsSharePermission")
	return
}

func NewModifySnapshotsSharePermissionResponse() (response *ModifySnapshotsSharePermissionResponse) {
	response = &ModifySnapshotsSharePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySnapshotsSharePermission）用于修改快照分享信息。
//
// 分享快照后，被分享账户可以通过该快照创建云硬盘。
// * 每个快照最多可分享给50个账户。
// * 分享快照无法更改名称，描述，仅可用于创建云硬盘。
// * 只支持分享到对方账户相同地域。
// * 仅支持分享数据盘快照。
func (c *Client) ModifySnapshotsSharePermission(request *ModifySnapshotsSharePermissionRequest) (response *ModifySnapshotsSharePermissionResponse, err error) {
	if request == nil {
		request = NewModifySnapshotsSharePermissionRequest()
	}
	response = NewModifySnapshotsSharePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDisksRenewFlagRequest() (request *ModifyDisksRenewFlagRequest) {
	request = &ModifyDisksRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDisksRenewFlag")
	return
}

func NewModifyDisksRenewFlagResponse() (response *ModifyDisksRenewFlagResponse) {
	response = &ModifyDisksRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDisksRenewFlag）用于修改云硬盘续费标识，支持批量修改。
func (c *Client) ModifyDisksRenewFlag(request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyDisksRenewFlagRequest()
	}
	response = NewModifyDisksRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewApplySnapshotRequest() (request *ApplySnapshotRequest) {
	request = &ApplySnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ApplySnapshot")
	return
}

func NewApplySnapshotResponse() (response *ApplySnapshotResponse) {
	response = &ApplySnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ApplySnapshot）用于回滚快照到原云硬盘。
//
// * 仅支持回滚到原云硬盘上。对于数据盘快照，如果您需要复制快照数据到其它云硬盘上，请使用[CreateDisks](/document/product/362/16312)接口创建新的弹性云盘，将快照数据复制到新购云盘上。
// * 用于回滚的快照必须处于NORMAL状态。快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
// * 如果是弹性云盘，则云盘必须处于未挂载状态，云硬盘挂载状态可以通过[DescribeDisks](/document/product/362/16315)接口查询，见Attached字段解释；如果是随云主机一起购买的非弹性云盘，则云主机必须处于关机状态，云主机状态可以通过[DescribeInstancesStatus](/document/product/213/15738)接口查询。
func (c *Client) ApplySnapshot(request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
	if request == nil {
		request = NewApplySnapshotRequest()
	}
	response = NewApplySnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
	request = &BindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "BindAutoSnapshotPolicy")
	return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
	response = &BindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindAutoSnapshotPolicy）用于绑定云硬盘到指定的定期快照策略。
//
// * 每个地域最多可创建10个定期快照策略, 每个定期快照策略最多能绑定80个云硬盘。
// * 当已绑定定期快照策略的云硬盘处于未使用状态（即弹性云盘未挂载或非弹性云盘的主机处于关机状态）将不会创建定期快照。
//
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewBindAutoSnapshotPolicyRequest()
	}
	response = NewBindAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisks")
	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	if request == nil {
		request = NewDescribeDisksRequest()
	}
	response = NewDescribeDisksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshot")
	return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSnapshot）用于对指定云盘创建快照。
//
// * 只有具有快照能力的云硬盘才能创建快照。云硬盘是否具有快照能力可由[DescribeDisks](/document/product/362/16315)接口查询，见SnapshotAbility字段。
// * 可创建快照数量限制见[产品使用限制](https://cloud.tencent.com/doc/product/362/5145)。
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	response = NewCreateSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewGetSnapOverviewRequest() (request *GetSnapOverviewRequest) {
	request = &GetSnapOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "GetSnapOverview")
	return
}

func NewGetSnapOverviewResponse() (response *GetSnapOverviewResponse) {
	response = &GetSnapOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取快照概览信息
func (c *Client) GetSnapOverview(request *GetSnapOverviewRequest) (response *GetSnapOverviewResponse, err error) {
	if request == nil {
		request = NewGetSnapOverviewRequest()
	}
	response = NewGetSnapOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskConfigQuotaRequest() (request *DescribeDiskConfigQuotaRequest) {
	request = &DescribeDiskConfigQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskConfigQuota")
	return
}

func NewDescribeDiskConfigQuotaResponse() (response *DescribeDiskConfigQuotaResponse) {
	response = &DescribeDiskConfigQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskConfigQuota）用于查询云硬盘配额。
func (c *Client) DescribeDiskConfigQuota(request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeDiskConfigQuotaRequest()
	}
	response = NewDescribeDiskConfigQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterModifyDiskAttributesRequest() (request *SwitchParameterModifyDiskAttributesRequest) {
	request = &SwitchParameterModifyDiskAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "SwitchParameterModifyDiskAttributes")
	return
}

func NewSwitchParameterModifyDiskAttributesResponse() (response *SwitchParameterModifyDiskAttributesResponse) {
	response = &SwitchParameterModifyDiskAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SwitchParameterModifyDiskAttributes）用于获取修改云盘类型的订单参数。
// * 仅支持预付费弹性云盘；
// * 当前仅支持一次传入一块云盘。
func (c *Client) SwitchParameterModifyDiskAttributes(request *SwitchParameterModifyDiskAttributesRequest) (response *SwitchParameterModifyDiskAttributesResponse, err error) {
	if request == nil {
		request = NewSwitchParameterModifyDiskAttributesRequest()
	}
	response = NewSwitchParameterModifyDiskAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockStoragesRequest() (request *DescribeBlockStoragesRequest) {
	request = &DescribeBlockStoragesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeBlockStorages")
	return
}

func NewDescribeBlockStoragesResponse() (response *DescribeBlockStoragesResponse) {
	response = &DescribeBlockStoragesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBlockStorages）用于查询块存储列表。可以同时查询本地盘和云盘。
func (c *Client) DescribeBlockStorages(request *DescribeBlockStoragesRequest) (response *DescribeBlockStoragesResponse, err error) {
	if request == nil {
		request = NewDescribeBlockStoragesRequest()
	}
	response = NewDescribeBlockStoragesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
	request = &DescribeAutoSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeAutoSnapshotPolicies")
	return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
	response = &DescribeAutoSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAutoSnapshotPolicies）用于查询定期快照策略。
//
// * 可以根据定期快照策略ID、名称或者状态等信息来查询定期快照策略的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期快照策略表。
//
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoSnapshotPoliciesRequest()
	}
	response = NewDescribeAutoSnapshotPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStoragePoolGroupsRequest() (request *DescribeDiskStoragePoolGroupsRequest) {
	request = &DescribeDiskStoragePoolGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStoragePoolGroups")
	return
}

func NewDescribeDiskStoragePoolGroupsResponse() (response *DescribeDiskStoragePoolGroupsResponse) {
	response = &DescribeDiskStoragePoolGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云硬盘存储资源池组信息
func (c *Client) DescribeDiskStoragePoolGroups(request *DescribeDiskStoragePoolGroupsRequest) (response *DescribeDiskStoragePoolGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStoragePoolGroupsRequest()
	}
	response = NewDescribeDiskStoragePoolGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotSharePermissionRequest() (request *DescribeSnapshotSharePermissionRequest) {
	request = &DescribeSnapshotSharePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotSharePermission")
	return
}

func NewDescribeSnapshotSharePermissionResponse() (response *DescribeSnapshotSharePermissionResponse) {
	response = &DescribeSnapshotSharePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshotSharePermission）用于查询快照的分享信息。
func (c *Client) DescribeSnapshotSharePermission(request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotSharePermissionRequest()
	}
	response = NewDescribeSnapshotSharePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskSupportFeaturesRequest() (request *DescribeDiskSupportFeaturesRequest) {
	request = &DescribeDiskSupportFeaturesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskSupportFeatures")
	return
}

func NewDescribeDiskSupportFeaturesResponse() (response *DescribeDiskSupportFeaturesResponse) {
	response = &DescribeDiskSupportFeaturesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云硬盘产品支持的特性。
func (c *Client) DescribeDiskSupportFeatures(request *DescribeDiskSupportFeaturesRequest) (response *DescribeDiskSupportFeaturesResponse, err error) {
	if request == nil {
		request = NewDescribeDiskSupportFeaturesRequest()
	}
	response = NewDescribeDiskSupportFeaturesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotGroupsDeniedActionsRequest() (request *DescribeSnapshotGroupsDeniedActionsRequest) {
	request = &DescribeSnapshotGroupsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotGroupsDeniedActions")
	return
}

func NewDescribeSnapshotGroupsDeniedActionsResponse() (response *DescribeSnapshotGroupsDeniedActionsResponse) {
	response = &DescribeSnapshotGroupsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshotGroupsDeniedActions）用于查询快照组的操作掩码。
//
// * 根据入参SnapshotGroupIds提供的快照组ID，返回对应快照组的操作掩码。
func (c *Client) DescribeSnapshotGroupsDeniedActions(request *DescribeSnapshotGroupsDeniedActionsRequest) (response *DescribeSnapshotGroupsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotGroupsDeniedActionsRequest()
	}
	response = NewDescribeSnapshotGroupsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotsDeniedActionsRequest() (request *DescribeSnapshotsDeniedActionsRequest) {
	request = &DescribeSnapshotsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotsDeniedActions")
	return
}

func NewDescribeSnapshotsDeniedActionsResponse() (response *DescribeSnapshotsDeniedActionsResponse) {
	response = &DescribeSnapshotsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshotsDeniedActions）用于查询快照的操作掩码。
//
// * 根据入参SnapshotIds提供的快照ID，返回对应快照的操作掩码。
func (c *Client) DescribeSnapshotsDeniedActions(request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotsDeniedActionsRequest()
	}
	response = NewDescribeSnapshotsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
	request = &TerminateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "TerminateDisks")
	return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
	response = &TerminateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（TerminateDisks）用于退还云硬盘。
//
// * 当前仅支持退还包年包月云盘。
// * 支持批量操作，每次请求批量云硬盘的上限为50。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
	if request == nil {
		request = NewTerminateDisksRequest()
	}
	response = NewTerminateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskExtraPerformanceRequest() (request *ModifyDiskExtraPerformanceRequest) {
	request = &ModifyDiskExtraPerformanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskExtraPerformance")
	return
}

func NewModifyDiskExtraPerformanceResponse() (response *ModifyDiskExtraPerformanceResponse) {
	response = &ModifyDiskExtraPerformanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDiskExtraPerformance）用于调整云硬盘额外的性能。
//
// * 目前仅支持极速型SSD云硬盘（CLOUD_TSSD）和高性能SSD云硬盘(CLOUD_HSSD)。
func (c *Client) ModifyDiskExtraPerformance(request *ModifyDiskExtraPerformanceRequest) (response *ModifyDiskExtraPerformanceResponse, err error) {
	if request == nil {
		request = NewModifyDiskExtraPerformanceRequest()
	}
	response = NewModifyDiskExtraPerformanceResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateSnapshotsRequest() (request *InquiryPriceCreateSnapshotsRequest) {
	request = &InquiryPriceCreateSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceCreateSnapshots")
	return
}

func NewInquiryPriceCreateSnapshotsResponse() (response *InquiryPriceCreateSnapshotsResponse) {
	response = &InquiryPriceCreateSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceCreateSnapshots）用于创建快照询价。
// * 当前快照不计费，询价价格均为0。
func (c *Client) InquiryPriceCreateSnapshots(request *InquiryPriceCreateSnapshotsRequest) (response *InquiryPriceCreateSnapshotsResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateSnapshotsRequest()
	}
	response = NewInquiryPriceCreateSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterResizeDiskRequest() (request *SwitchParameterResizeDiskRequest) {
	request = &SwitchParameterResizeDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "SwitchParameterResizeDisk")
	return
}

func NewSwitchParameterResizeDiskResponse() (response *SwitchParameterResizeDiskResponse) {
	response = &SwitchParameterResizeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SwitchParameterResizeDisk）用于得到扩容云硬盘的订单参数。
func (c *Client) SwitchParameterResizeDisk(request *SwitchParameterResizeDiskRequest) (response *SwitchParameterResizeDiskResponse, err error) {
	if request == nil {
		request = NewSwitchParameterResizeDiskRequest()
	}
	response = NewSwitchParameterResizeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
	request = &DeleteSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshots")
	return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
	response = &DeleteSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSnapshots）用于删除快照。
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以返回特定的错误码返回。
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotsRequest()
	}
	response = NewDeleteSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotGroupsRequest() (request *DescribeSnapshotGroupsRequest) {
	request = &DescribeSnapshotGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotGroups")
	return
}

func NewDescribeSnapshotGroupsResponse() (response *DescribeSnapshotGroupsResponse) {
	response = &DescribeSnapshotGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照组列表，返回状态正常、创建中、回滚中的快照组列表
func (c *Client) DescribeSnapshotGroups(request *DescribeSnapshotGroupsRequest) (response *DescribeSnapshotGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotGroupsRequest()
	}
	response = NewDescribeSnapshotGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterRenewDisksRequest() (request *SwitchParameterRenewDisksRequest) {
	request = &SwitchParameterRenewDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "SwitchParameterRenewDisks")
	return
}

func NewSwitchParameterRenewDisksResponse() (response *SwitchParameterRenewDisksResponse) {
	response = &SwitchParameterRenewDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SwitchParameterRenewDisks）用于得到续费云硬盘的订单参数。
// * 支持与挂载云主机一起续费的场景，需要在DiskChargePrepaid参数中指定CurInstanceDeadline，此时会按对齐到子机到期时间来续费。
func (c *Client) SwitchParameterRenewDisks(request *SwitchParameterRenewDisksRequest) (response *SwitchParameterRenewDisksResponse, err error) {
	if request == nil {
		request = NewSwitchParameterRenewDisksRequest()
	}
	response = NewSwitchParameterRenewDisksResponse()
	err = c.Send(request, response)
	return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
	request = &AttachDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "AttachDisks")
	return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
	response = &AttachDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachDisks）用于挂载云硬盘。
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘存在不允许挂载的云盘，则操作不执行，以返回特定的错误码返回。
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
	if request == nil {
		request = NewAttachDisksRequest()
	}
	response = NewAttachDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksDeniedActionsRequest() (request *DescribeDisksDeniedActionsRequest) {
	request = &DescribeDisksDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisksDeniedActions")
	return
}

func NewDescribeDisksDeniedActionsResponse() (response *DescribeDisksDeniedActionsResponse) {
	response = &DescribeDisksDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisksDeniedActions）用于查询云硬盘的操作掩码。
//
// * 根据入参DiskIds提供的云盘ID，返回对应云硬盘的操作掩码。
func (c *Client) DescribeDisksDeniedActions(request *DescribeDisksDeniedActionsRequest) (response *DescribeDisksDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeDisksDeniedActionsRequest()
	}
	response = NewDescribeDisksDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyDiskAttributesRequest() (request *InquiryPriceModifyDiskAttributesRequest) {
	request = &InquiryPriceModifyDiskAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceModifyDiskAttributes")
	return
}

func NewInquiryPriceModifyDiskAttributesResponse() (response *InquiryPriceModifyDiskAttributesResponse) {
	response = &InquiryPriceModifyDiskAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceModifyDiskAttributes）用于修改云盘类型询价。
//
// * 当前仅支持弹性云盘修改类型（[DescribeDisks](/document/product/362/16315)接口的返回字段Portable为true表示弹性云盘）。
// * 当前仅支持云盘类型升级，不支持降级，具体如下:
//     * CLOUD_BASIC变更为CLOUD_PREMIUM；
//     * CLOUD_BASIC变更为CLOUD_SSD；
//     * CLOUD_PREMIUM变更为CLOUD_SSD。
func (c *Client) InquiryPriceModifyDiskAttributes(request *InquiryPriceModifyDiskAttributesRequest) (response *InquiryPriceModifyDiskAttributesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyDiskAttributesRequest()
	}
	response = NewInquiryPriceModifyDiskAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
	request = &CreateAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateAutoSnapshotPolicy")
	return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
	response = &CreateAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAutoSnapshotPolicy）用于创建定期快照策略。
//
// * 每个地域最多创建10个定期快照策略。
// * 每个地域可创建的快照有数量和容量的限制，具体请见腾讯云控制台快照页面提示。
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAutoSnapshotPolicyRequest()
	}
	response = NewCreateAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneDiskConfigInfosRequest() (request *DescribeZoneDiskConfigInfosRequest) {
	request = &DescribeZoneDiskConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeZoneDiskConfigInfos")
	return
}

func NewDescribeZoneDiskConfigInfosResponse() (response *DescribeZoneDiskConfigInfosResponse) {
	response = &DescribeZoneDiskConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对应的云盘资源是否可以创建
func (c *Client) DescribeZoneDiskConfigInfos(request *DescribeZoneDiskConfigInfosRequest) (response *DescribeZoneDiskConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneDiskConfigInfosRequest()
	}
	response = NewDescribeZoneDiskConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
	request = &DetachDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DetachDisks")
	return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
	response = &DetachDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DetachDisks）用于解挂云硬盘。
//
// * 支持批量操作，解挂挂载在同一主机上的多块云盘。如果多块云盘存在不允许解挂载的云盘，则操作不执行，以返回特定的错误码返回。
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机解挂载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为解挂载成功。
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
	if request == nil {
		request = NewDetachDisksRequest()
	}
	response = NewDetachDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotGroupRequest() (request *DeleteSnapshotGroupRequest) {
	request = &DeleteSnapshotGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshotGroup")
	return
}

func NewDeleteSnapshotGroupResponse() (response *DeleteSnapshotGroupResponse) {
	response = &DeleteSnapshotGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除快照组，一次调用仅支持删除一个快照组
// * 默认会删除快照组内的所有快照
// * 如果快照组内的快照有关联镜像，则删除失败，所有快照均不会删除；如果需要同时删除快照绑定的镜像，可传入参数DeleteBindImages=true
func (c *Client) DeleteSnapshotGroup(request *DeleteSnapshotGroupRequest) (response *DeleteSnapshotGroupResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotGroupRequest()
	}
	response = NewDeleteSnapshotGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
	request = &ModifySnapshotAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotAttribute")
	return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
	response = &ModifySnapshotAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifySnapshotAttribute）用于修改指定快照的属性。
//
// * 当前仅支持修改快照名称及将非永久快照修改为永久快照。
// * “快照名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行快照管理操作的依据。
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
	if request == nil {
		request = NewModifySnapshotAttributeRequest()
	}
	response = NewModifySnapshotAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRenewDiskRequest() (request *RenewDiskRequest) {
	request = &RenewDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RenewDisk")
	return
}

func NewRenewDiskResponse() (response *RenewDiskResponse) {
	response = &RenewDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RenewDisk）用于续费云硬盘。
//
// * 只支持预付费的云硬盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中DiskChargeType字段解释。
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到子机续费后的到期时间来续费。
// * 续费时请确保账户余额充足。可通过[DescribeAccountBalance](/document/product/378/4397)接口查询账户余额。
func (c *Client) RenewDisk(request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
	if request == nil {
		request = NewRenewDiskRequest()
	}
	response = NewRenewDiskResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateDisksRequest() (request *InquiryPriceCreateDisksRequest) {
	request = &InquiryPriceCreateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceCreateDisks")
	return
}

func NewInquiryPriceCreateDisksResponse() (response *InquiryPriceCreateDisksResponse) {
	response = &InquiryPriceCreateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceCreateDisks）用于创建云硬盘询价。
//
// * 支持查询创建多块云硬盘的价格，此时返回结果为总价格。
func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateDisksRequest()
	}
	response = NewInquiryPriceCreateDisksResponse()
	err = c.Send(request, response)
	return
}

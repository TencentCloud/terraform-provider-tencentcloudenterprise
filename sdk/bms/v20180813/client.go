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

package v20180813

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-08-13"

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

func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户自定义镜像
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	response = NewDeleteImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHeartbeatRequest() (request *ModifyHeartbeatRequest) {
	request = &ModifyHeartbeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ModifyHeartbeat")
	return
}

func NewModifyHeartbeatResponse() (response *ModifyHeartbeatResponse) {
	response = &ModifyHeartbeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于设置一个或多个bms心跳网络组
func (c *Client) ModifyHeartbeat(request *ModifyHeartbeatRequest) (response *ModifyHeartbeatResponse, err error) {
	if request == nil {
		request = NewModifyHeartbeatRequest()
	}
	response = NewModifyHeartbeatResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupIPRequest() (request *CreateBackupIPRequest) {
	request = &CreateBackupIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "CreateBackupIP")
	return
}

func NewCreateBackupIPResponse() (response *CreateBackupIPResponse) {
	response = &CreateBackupIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建备用IP
func (c *Client) CreateBackupIP(request *CreateBackupIPRequest) (response *CreateBackupIPResponse, err error) {
	if request == nil {
		request = NewCreateBackupIPRequest()
	}
	response = NewCreateBackupIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeInstances")
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

func NewDescribeOutbandInfoRequest() (request *DescribeOutbandInfoRequest) {
	request = &DescribeOutbandInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeOutbandInfo")
	return
}

func NewDescribeOutbandInfoResponse() (response *DescribeOutbandInfoResponse) {
	response = &DescribeOutbandInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询BMS带外账号密码
func (c *Client) DescribeOutbandInfo(request *DescribeOutbandInfoRequest) (response *DescribeOutbandInfoResponse, err error) {
	if request == nil {
		request = NewDescribeOutbandInfoRequest()
	}
	response = NewDescribeOutbandInfoResponse()
	err = c.Send(request, response)
	return
}

func NewReturnBackupIPRequest() (request *ReturnBackupIPRequest) {
	request = &ReturnBackupIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ReturnBackupIP")
	return
}

func NewReturnBackupIPResponse() (response *ReturnBackupIPResponse) {
	response = &ReturnBackupIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退还备用IP
func (c *Client) ReturnBackupIP(request *ReturnBackupIPRequest) (response *ReturnBackupIPResponse, err error) {
	if request == nil {
		request = NewReturnBackupIPRequest()
	}
	response = NewReturnBackupIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisasterRecoverGroupsRequest() (request *DescribeDisasterRecoverGroupsRequest) {
	request = &DescribeDisasterRecoverGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeDisasterRecoverGroups")
	return
}

func NewDescribeDisasterRecoverGroupsResponse() (response *DescribeDisasterRecoverGroupsResponse) {
	response = &DescribeDisasterRecoverGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询置放群组信息列表
func (c *Client) DescribeDisasterRecoverGroups(request *DescribeDisasterRecoverGroupsRequest) (response *DescribeDisasterRecoverGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoverGroupsRequest()
	}
	response = NewDescribeDisasterRecoverGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "StartInstances")
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
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	response = NewStartInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "TerminateInstances")
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
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOutbandIPRequest() (request *DeleteOutbandIPRequest) {
	request = &DeleteOutbandIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteOutbandIP")
	return
}

func NewDeleteOutbandIPResponse() (response *DeleteOutbandIPResponse) {
	response = &DeleteOutbandIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除带外映射
func (c *Client) DeleteOutbandIP(request *DeleteOutbandIPRequest) (response *DeleteOutbandIPResponse, err error) {
	if request == nil {
		request = NewDeleteOutbandIPRequest()
	}
	response = NewDeleteOutbandIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeDisks")
	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisks）用于查询BMS硬盘列表。
//
// * 可以根据BMS硬盘ID等信息来查询BMS硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的BMS硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	if request == nil {
		request = NewDescribeDisksRequest()
	}
	response = NewDescribeDisksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ModifyInstancesAttribute")
	return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
	response = &ModifyInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyInstancesAttributeRequest()
	}
	response = NewModifyInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateHeartbeatRequest() (request *UpdateHeartbeatRequest) {
	request = &UpdateHeartbeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateHeartbeat")
	return
}

func NewUpdateHeartbeatResponse() (response *UpdateHeartbeatResponse) {
	response = &UpdateHeartbeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新心跳网络组名称
func (c *Client) UpdateHeartbeat(request *UpdateHeartbeatRequest) (response *UpdateHeartbeatResponse, err error) {
	if request == nil {
		request = NewUpdateHeartbeatRequest()
	}
	response = NewUpdateHeartbeatResponse()
	err = c.Send(request, response)
	return
}

func NewAddTenantImageRequest() (request *AddTenantImageRequest) {
	request = &AddTenantImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddTenantImage")
	return
}

func NewAddTenantImageResponse() (response *AddTenantImageResponse) {
	response = &AddTenantImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加用户自定义镜像
func (c *Client) AddTenantImage(request *AddTenantImageRequest) (response *AddTenantImageResponse, err error) {
	if request == nil {
		request = NewAddTenantImageRequest()
	}
	response = NewAddTenantImageResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
	request = &ResetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ResetInstance")
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
// * 如果指定了`OperatingSystem`参数，则使用指定的系统重装；否则按照当前实例使用的系统进行重装。
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	if request == nil {
		request = NewResetInstanceRequest()
	}
	response = NewResetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDisasterRecoverGroupRequest() (request *UpdateDisasterRecoverGroupRequest) {
	request = &UpdateDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateDisasterRecoverGroup")
	return
}

func NewUpdateDisasterRecoverGroupResponse() (response *UpdateDisasterRecoverGroupResponse) {
	response = &UpdateDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改置放群组信息名称
func (c *Client) UpdateDisasterRecoverGroup(request *UpdateDisasterRecoverGroupRequest) (response *UpdateDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewUpdateDisasterRecoverGroupRequest()
	}
	response = NewUpdateDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHeartbeatsRequest() (request *DescribeHeartbeatsRequest) {
	request = &DescribeHeartbeatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeHeartbeats")
	return
}

func NewDescribeHeartbeatsResponse() (response *DescribeHeartbeatsResponse) {
	response = &DescribeHeartbeatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询心跳网络组信息
func (c *Client) DescribeHeartbeats(request *DescribeHeartbeatsRequest) (response *DescribeHeartbeatsResponse, err error) {
	if request == nil {
		request = NewDescribeHeartbeatsRequest()
	}
	response = NewDescribeHeartbeatsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageListRequest() (request *DescribeImageListRequest) {
	request = &DescribeImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeImageList")
	return
}

func NewDescribeImageListResponse() (response *DescribeImageListResponse) {
	response = &DescribeImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像列表展示
func (c *Client) DescribeImageList(request *DescribeImageListRequest) (response *DescribeImageListResponse, err error) {
	if request == nil {
		request = NewDescribeImageListRequest()
	}
	response = NewDescribeImageListResponse()
	err = c.Send(request, response)
	return
}

func NewResetOutbandInfoRequest() (request *ResetOutbandInfoRequest) {
	request = &ResetOutbandInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ResetOutbandInfo")
	return
}

func NewResetOutbandInfoResponse() (response *ResetOutbandInfoResponse) {
	response = &ResetOutbandInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置BMS带外账号密码
func (c *Client) ResetOutbandInfo(request *ResetOutbandInfoRequest) (response *ResetOutbandInfoResponse, err error) {
	if request == nil {
		request = NewResetOutbandInfoRequest()
	}
	response = NewResetOutbandInfoResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "StopInstances")
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
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	response = NewStopInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDisasterRecoverGroupRequest() (request *CreateDisasterRecoverGroupRequest) {
	request = &CreateDisasterRecoverGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "CreateDisasterRecoverGroup")
	return
}

func NewCreateDisasterRecoverGroupResponse() (response *CreateDisasterRecoverGroupResponse) {
	response = &CreateDisasterRecoverGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建分散置放群组，该功能对资源挑战较大，请谨慎使用。创建好的容灾组，可在创建实例的时指定。
func (c *Client) CreateDisasterRecoverGroup(request *CreateDisasterRecoverGroupRequest) (response *CreateDisasterRecoverGroupResponse, err error) {
	if request == nil {
		request = NewCreateDisasterRecoverGroupRequest()
	}
	response = NewCreateDisasterRecoverGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDisasterRecoverGroupsRequest() (request *DeleteDisasterRecoverGroupsRequest) {
	request = &DeleteDisasterRecoverGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteDisasterRecoverGroups")
	return
}

func NewDeleteDisasterRecoverGroupsResponse() (response *DeleteDisasterRecoverGroupsResponse) {
	response = &DeleteDisasterRecoverGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteDisasterRecoverGroups)用于删除[分散置放群组]。只有空的置放群组才能被删除，非空的群组需要先销毁组内所有BMS服务器，才能执行删除操作，不然会产生删除置放群组失败的错误。
func (c *Client) DeleteDisasterRecoverGroups(request *DeleteDisasterRecoverGroupsRequest) (response *DeleteDisasterRecoverGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteDisasterRecoverGroupsRequest()
	}
	response = NewDeleteDisasterRecoverGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlavorsRequest() (request *DescribeFlavorsRequest) {
	request = &DescribeFlavorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeFlavors")
	return
}

func NewDescribeFlavorsResponse() (response *DescribeFlavorsResponse) {
	response = &DescribeFlavorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示套餐列表详情。
func (c *Client) DescribeFlavors(request *DescribeFlavorsRequest) (response *DescribeFlavorsResponse, err error) {
	if request == nil {
		request = NewDescribeFlavorsRequest()
	}
	response = NewDescribeFlavorsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOutbandIPRequest() (request *CreateOutbandIPRequest) {
	request = &CreateOutbandIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "CreateOutbandIP")
	return
}

func NewCreateOutbandIPResponse() (response *CreateOutbandIPResponse) {
	response = &CreateOutbandIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建带外映射
func (c *Client) CreateOutbandIP(request *CreateOutbandIPRequest) (response *CreateOutbandIPResponse, err error) {
	if request == nil {
		request = NewCreateOutbandIPRequest()
	}
	response = NewCreateOutbandIPResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHeartbeatRequest() (request *CreateHeartbeatRequest) {
	request = &CreateHeartbeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "CreateHeartbeat")
	return
}

func NewCreateHeartbeatResponse() (response *CreateHeartbeatResponse) {
	response = &CreateHeartbeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建心跳网络组
func (c *Client) CreateHeartbeat(request *CreateHeartbeatRequest) (response *CreateHeartbeatResponse, err error) {
	if request == nil {
		request = NewCreateHeartbeatRequest()
	}
	response = NewCreateHeartbeatResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHeartbeatsRequest() (request *DeleteHeartbeatsRequest) {
	request = &DeleteHeartbeatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteHeartbeats")
	return
}

func NewDeleteHeartbeatsResponse() (response *DeleteHeartbeatsResponse) {
	response = &DeleteHeartbeatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除心跳网络组
func (c *Client) DeleteHeartbeats(request *DeleteHeartbeatsRequest) (response *DeleteHeartbeatsResponse, err error) {
	if request == nil {
		request = NewDeleteHeartbeatsRequest()
	}
	response = NewDeleteHeartbeatsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskRequest() (request *QueryTaskRequest) {
	request = &QueryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "QueryTask")
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

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "RebootInstances")
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
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	response = NewRebootInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceBmsInstanceForTradeRequest() (request *InquiryPriceBmsInstanceForTradeRequest) {
	request = &InquiryPriceBmsInstanceForTradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "InquiryPriceBmsInstanceForTrade")
	return
}

func NewInquiryPriceBmsInstanceForTradeResponse() (response *InquiryPriceBmsInstanceForTradeResponse) {
	response = &InquiryPriceBmsInstanceForTradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bms询价
func (c *Client) InquiryPriceBmsInstanceForTrade(request *InquiryPriceBmsInstanceForTradeRequest) (response *InquiryPriceBmsInstanceForTradeResponse, err error) {
	if request == nil {
		request = NewInquiryPriceBmsInstanceForTradeRequest()
	}
	response = NewInquiryPriceBmsInstanceForTradeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesGroupIdRequest() (request *ModifyInstancesGroupIdRequest) {
	request = &ModifyInstancesGroupIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ModifyInstancesGroupId")
	return
}

func NewModifyInstancesGroupIdResponse() (response *ModifyInstancesGroupIdResponse) {
	response = &ModifyInstancesGroupIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesGroupId) 用于修改实例置放群组
func (c *Client) ModifyInstancesGroupId(request *ModifyInstancesGroupIdRequest) (response *ModifyInstancesGroupIdResponse, err error) {
	if request == nil {
		request = NewModifyInstancesGroupIdRequest()
	}
	response = NewModifyInstancesGroupIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTenantImageRequest() (request *ModifyTenantImageRequest) {
	request = &ModifyTenantImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ModifyTenantImage")
	return
}

func NewModifyTenantImageResponse() (response *ModifyTenantImageResponse) {
	response = &ModifyTenantImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户自定义镜像
func (c *Client) ModifyTenantImage(request *ModifyTenantImageRequest) (response *ModifyTenantImageResponse, err error) {
	if request == nil {
		request = NewModifyTenantImageRequest()
	}
	response = NewModifyTenantImageResponse()
	err = c.Send(request, response)
	return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "RunInstances")
	return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	response = NewRunInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOutbandIPRequest() (request *DescribeOutbandIPRequest) {
	request = &DescribeOutbandIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeOutbandIP")
	return
}

func NewDescribeOutbandIPResponse() (response *DescribeOutbandIPResponse) {
	response = &DescribeOutbandIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示带外映射
func (c *Client) DescribeOutbandIP(request *DescribeOutbandIPRequest) (response *DescribeOutbandIPResponse, err error) {
	if request == nil {
		request = NewDescribeOutbandIPRequest()
	}
	response = NewDescribeOutbandIPResponse()
	err = c.Send(request, response)
	return
}

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

package v20180416

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-04-16"

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

func NewDeleteCfsRuleRequest() (request *DeleteCfsRuleRequest) {
	request = &DeleteCfsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsRule")
	return
}

func NewDeleteCfsRuleResponse() (response *DeleteCfsRuleResponse) {
	response = &DeleteCfsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCfsRule）用于删除权限组规则。
func (c *Client) DeleteCfsRule(request *DeleteCfsRuleRequest) (response *DeleteCfsRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCfsRuleRequest()
	}
	response = NewDeleteCfsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMountTargetsWithRegionRequest() (request *QueryMountTargetsWithRegionRequest) {
	request = &QueryMountTargetsWithRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryMountTargetsWithRegion")
	return
}

func NewQueryMountTargetsWithRegionResponse() (response *QueryMountTargetsWithRegionResponse) {
	response = &QueryMountTargetsWithRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询区域挂载点情况，只用于控制台
func (c *Client) QueryMountTargetsWithRegion(request *QueryMountTargetsWithRegionRequest) (response *QueryMountTargetsWithRegionResponse, err error) {
	if request == nil {
		request = NewQueryMountTargetsWithRegionRequest()
	}
	response = NewQueryMountTargetsWithRegionResponse()
	err = c.Send(request, response)
	return
}

func NewSignUpCfsServiceRequest() (request *SignUpCfsServiceRequest) {
	request = &SignUpCfsServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "SignUpCfsService")
	return
}

func NewSignUpCfsServiceResponse() (response *SignUpCfsServiceResponse) {
	response = &SignUpCfsServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SignUpCfsService）用于开通CFS服务。
func (c *Client) SignUpCfsService(request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
	if request == nil {
		request = NewSignUpCfsServiceRequest()
	}
	response = NewSignUpCfsServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCfsFileSystemRequest() (request *DeleteCfsFileSystemRequest) {
	request = &DeleteCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsFileSystem")
	return
}

func NewDeleteCfsFileSystemResponse() (response *DeleteCfsFileSystemResponse) {
	response = &DeleteCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除文件系统
func (c *Client) DeleteCfsFileSystem(request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewDeleteCfsFileSystemRequest()
	}
	response = NewDeleteCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsFileSystemSizeLimitRequest() (request *UpdateCfsFileSystemSizeLimitRequest) {
	request = &UpdateCfsFileSystemSizeLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemSizeLimit")
	return
}

func NewUpdateCfsFileSystemSizeLimitResponse() (response *UpdateCfsFileSystemSizeLimitResponse) {
	response = &UpdateCfsFileSystemSizeLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsFileSystemSizeLimit）用于更新文件系统存储容量限制。
func (c *Client) UpdateCfsFileSystemSizeLimit(request *UpdateCfsFileSystemSizeLimitRequest) (response *UpdateCfsFileSystemSizeLimitResponse, err error) {
	if request == nil {
		request = NewUpdateCfsFileSystemSizeLimitRequest()
	}
	response = NewUpdateCfsFileSystemSizeLimitResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsPGroupRequest() (request *UpdateCfsPGroupRequest) {
	request = &UpdateCfsPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsPGroup")
	return
}

func NewUpdateCfsPGroupResponse() (response *UpdateCfsPGroupResponse) {
	response = &UpdateCfsPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsPGroup）更新权限组信息。
func (c *Client) UpdateCfsPGroup(request *UpdateCfsPGroupRequest) (response *UpdateCfsPGroupResponse, err error) {
	if request == nil {
		request = NewUpdateCfsPGroupRequest()
	}
	response = NewUpdateCfsPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCfsFileSystemRequest() (request *QueryCfsFileSystemRequest) {
	request = &QueryCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryCfsFileSystem")
	return
}

func NewQueryCfsFileSystemResponse() (response *QueryCfsFileSystemResponse) {
	response = &QueryCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询文件系统
func (c *Client) QueryCfsFileSystem(request *QueryCfsFileSystemRequest) (response *QueryCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewQueryCfsFileSystemRequest()
	}
	response = NewQueryCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCfsPGroupRequest() (request *QueryCfsPGroupRequest) {
	request = &QueryCfsPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryCfsPGroup")
	return
}

func NewQueryCfsPGroupResponse() (response *QueryCfsPGroupResponse) {
	response = &QueryCfsPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryCfsPGroup）用于查询权限组列表。
func (c *Client) QueryCfsPGroup(request *QueryCfsPGroupRequest) (response *QueryCfsPGroupResponse, err error) {
	if request == nil {
		request = NewQueryCfsPGroupRequest()
	}
	response = NewQueryCfsPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsRuleRequest() (request *UpdateCfsRuleRequest) {
	request = &UpdateCfsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsRule")
	return
}

func NewUpdateCfsRuleResponse() (response *UpdateCfsRuleResponse) {
	response = &UpdateCfsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsRule）用于更新权限规则。
func (c *Client) UpdateCfsRule(request *UpdateCfsRuleRequest) (response *UpdateCfsRuleResponse, err error) {
	if request == nil {
		request = NewUpdateCfsRuleRequest()
	}
	response = NewUpdateCfsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCfsServiceStatusRequest() (request *QueryCfsServiceStatusRequest) {
	request = &QueryCfsServiceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryCfsServiceStatus")
	return
}

func NewQueryCfsServiceStatusResponse() (response *QueryCfsServiceStatusResponse) {
	response = &QueryCfsServiceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryCfsServiceStatus）用于查询用户使用CFS的服务状态。
func (c *Client) QueryCfsServiceStatus(request *QueryCfsServiceStatusRequest) (response *QueryCfsServiceStatusResponse, err error) {
	if request == nil {
		request = NewQueryCfsServiceStatusRequest()
	}
	response = NewQueryCfsServiceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCfsPGroupRequest() (request *DeleteCfsPGroupRequest) {
	request = &DeleteCfsPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsPGroup")
	return
}

func NewDeleteCfsPGroupResponse() (response *DeleteCfsPGroupResponse) {
	response = &DeleteCfsPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteCfsPGroup）用于删除权限组。
func (c *Client) DeleteCfsPGroup(request *DeleteCfsPGroupRequest) (response *DeleteCfsPGroupResponse, err error) {
	if request == nil {
		request = NewDeleteCfsPGroupRequest()
	}
	response = NewDeleteCfsPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsPGroupRequest() (request *CreateCfsPGroupRequest) {
	request = &CreateCfsPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsPGroup")
	return
}

func NewCreateCfsPGroupResponse() (response *CreateCfsPGroupResponse) {
	response = &CreateCfsPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsPGroup）用于创建权限组
func (c *Client) CreateCfsPGroup(request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
	if request == nil {
		request = NewCreateCfsPGroupRequest()
	}
	response = NewCreateCfsPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMountTargetRequest() (request *QueryMountTargetRequest) {
	request = &QueryMountTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryMountTarget")
	return
}

func NewQueryMountTargetResponse() (response *QueryMountTargetResponse) {
	response = &QueryMountTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryMountTarget）用于查询文件系统挂载点信息
func (c *Client) QueryMountTarget(request *QueryMountTargetRequest) (response *QueryMountTargetResponse, err error) {
	if request == nil {
		request = NewQueryMountTargetRequest()
	}
	response = NewQueryMountTargetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsFileSystemRequest() (request *CreateCfsFileSystemRequest) {
	request = &CreateCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsFileSystem")
	return
}

func NewCreateCfsFileSystemResponse() (response *CreateCfsFileSystemResponse) {
	response = &CreateCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于添加新文件系统
func (c *Client) CreateCfsFileSystem(request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewCreateCfsFileSystemRequest()
	}
	response = NewCreateCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsFileSystemPGroupRequest() (request *UpdateCfsFileSystemPGroupRequest) {
	request = &UpdateCfsFileSystemPGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemPGroup")
	return
}

func NewUpdateCfsFileSystemPGroupResponse() (response *UpdateCfsFileSystemPGroupResponse) {
	response = &UpdateCfsFileSystemPGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsFileSystemPGroup）用于更新文件系统所使用的权限组
func (c *Client) UpdateCfsFileSystemPGroup(request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
	if request == nil {
		request = NewUpdateCfsFileSystemPGroupRequest()
	}
	response = NewUpdateCfsFileSystemPGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsFileSystemNameRequest() (request *UpdateCfsFileSystemNameRequest) {
	request = &UpdateCfsFileSystemNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemName")
	return
}

func NewUpdateCfsFileSystemNameResponse() (response *UpdateCfsFileSystemNameResponse) {
	response = &UpdateCfsFileSystemNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateCfsFileSystemName）用于更新文件系统名
func (c *Client) UpdateCfsFileSystemName(request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
	if request == nil {
		request = NewUpdateCfsFileSystemNameRequest()
	}
	response = NewUpdateCfsFileSystemNameResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsRuleRequest() (request *CreateCfsRuleRequest) {
	request = &CreateCfsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsRule")
	return
}

func NewCreateCfsRuleResponse() (response *CreateCfsRuleResponse) {
	response = &CreateCfsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCfsRule）用于创建权限组规则。
func (c *Client) CreateCfsRule(request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
	if request == nil {
		request = NewCreateCfsRuleRequest()
	}
	response = NewCreateCfsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCfsKmsKeysRequest() (request *QueryCfsKmsKeysRequest) {
	request = &QueryCfsKmsKeysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryCfsKmsKeys")
	return
}

func NewQueryCfsKmsKeysResponse() (response *QueryCfsKmsKeysResponse) {
	response = &QueryCfsKmsKeysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryCfsKmsKeys）用于查询KMS服务的密钥。
func (c *Client) QueryCfsKmsKeys(request *QueryCfsKmsKeysRequest) (response *QueryCfsKmsKeysResponse, err error) {
	if request == nil {
		request = NewQueryCfsKmsKeysRequest()
	}
	response = NewQueryCfsKmsKeysResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAvailableZoneInfoRequest() (request *QueryAvailableZoneInfoRequest) {
	request = &QueryAvailableZoneInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryAvailableZoneInfo")
	return
}

func NewQueryAvailableZoneInfoResponse() (response *QueryAvailableZoneInfoResponse) {
	response = &QueryAvailableZoneInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryAvailableZoneInfo）用于查询区域的可用情况。
//
func (c *Client) QueryAvailableZoneInfo(request *QueryAvailableZoneInfoRequest) (response *QueryAvailableZoneInfoResponse, err error) {
	if request == nil {
		request = NewQueryAvailableZoneInfoRequest()
	}
	response = NewQueryAvailableZoneInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCfsRuleRequest() (request *QueryCfsRuleRequest) {
	request = &QueryCfsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "QueryCfsRule")
	return
}

func NewQueryCfsRuleResponse() (response *QueryCfsRuleResponse) {
	response = &QueryCfsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（QueryCfsRule）用于查询权限组规则列表。
func (c *Client) QueryCfsRule(request *QueryCfsRuleRequest) (response *QueryCfsRuleResponse, err error) {
	if request == nil {
		request = NewQueryCfsRuleRequest()
	}
	response = NewQueryCfsRuleResponse()
	err = c.Send(request, response)
	return
}

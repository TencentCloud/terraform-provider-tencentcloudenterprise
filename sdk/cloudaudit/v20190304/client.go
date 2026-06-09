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

package v20190304

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-03-04"

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

func NewDeleteExporJobRequest() (request *DeleteExporJobRequest) {
	request = &DeleteExporJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteExporJob")
	return
}

func NewDeleteExporJobResponse() (response *DeleteExporJobResponse) {
	response = &DeleteExporJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除任务
func (c *Client) DeleteExporJob(request *DeleteExporJobRequest) (response *DeleteExporJobResponse, err error) {
	if request == nil {
		request = NewDeleteExporJobRequest()
	}
	response = NewDeleteExporJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetExporLoggingListRequest() (request *GetExporLoggingListRequest) {
	request = &GetExporLoggingListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "GetExporLoggingList")
	return
}

func NewGetExporLoggingListResponse() (response *GetExporLoggingListResponse) {
	response = &GetExporLoggingListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务列表
func (c *Client) GetExporLoggingList(request *GetExporLoggingListRequest) (response *GetExporLoggingListResponse, err error) {
	if request == nil {
		request = NewGetExporLoggingListRequest()
	}
	response = NewGetExporLoggingListResponse()
	err = c.Send(request, response)
	return
}

func NewLookupEventsRequest() (request *LookupEventsRequest) {
	request = &LookupEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "LookupEvents")
	return
}

func NewLookupEventsResponse() (response *LookupEventsResponse) {
	response = &LookupEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索日志接口
func (c *Client) LookupEvents(request *LookupEventsRequest) (response *LookupEventsResponse, err error) {
	if request == nil {
		request = NewLookupEventsRequest()
	}
	response = NewLookupEventsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatExporJobRequest() (request *CreatExporJobRequest) {
	request = &CreatExporJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "CreatExporJob")
	return
}

func NewCreatExporJobResponse() (response *CreatExporJobResponse) {
	response = &CreatExporJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建导出任务
func (c *Client) CreatExporJob(request *CreatExporJobRequest) (response *CreatExporJobResponse, err error) {
	if request == nil {
		request = NewCreatExporJobRequest()
	}
	response = NewCreatExporJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBucketRequest() (request *CreateBucketRequest) {
	request = &CreateBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateBucket")
	return
}

func NewCreateBucketResponse() (response *CreateBucketResponse) {
	response = &CreateBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建存储桶
func (c *Client) CreateBucket(request *CreateBucketRequest) (response *CreateBucketResponse, err error) {
	if request == nil {
		request = NewCreateBucketRequest()
	}
	response = NewCreateBucketResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAuditTrackRequest() (request *CreateAuditTrackRequest) {
	request = &CreateAuditTrackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAuditTrack")
	return
}

func NewCreateAuditTrackResponse() (response *CreateAuditTrackResponse) {
	response = &CreateAuditTrackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建跟踪集。
func (c *Client) CreateAuditTrack(request *CreateAuditTrackRequest) (response *CreateAuditTrackResponse, err error) {
	if request == nil {
		request = NewCreateAuditTrackRequest()
	}
	response = NewCreateAuditTrackResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAuditTrackRequest() (request *DeleteAuditTrackRequest) {
	request = &DeleteAuditTrackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAuditTrack")
	return
}

func NewDeleteAuditTrackResponse() (response *DeleteAuditTrackResponse) {
	response = &DeleteAuditTrackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除跟踪集。
func (c *Client) DeleteAuditTrack(request *DeleteAuditTrackRequest) (response *DeleteAuditTrackResponse, err error) {
	if request == nil {
		request = NewDeleteAuditTrackRequest()
	}
	response = NewDeleteAuditTrackResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuditTrackRequest() (request *DescribeAuditTrackRequest) {
	request = &DescribeAuditTrackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTrack")
	return
}

func NewDescribeAuditTrackResponse() (response *DescribeAuditTrackResponse) {
	response = &DescribeAuditTrackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询跟踪集详情。
func (c *Client) DescribeAuditTrack(request *DescribeAuditTrackRequest) (response *DescribeAuditTrackResponse, err error) {
	if request == nil {
		request = NewDescribeAuditTrackRequest()
	}
	response = NewDescribeAuditTrackResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuditTracksRequest() (request *DescribeAuditTracksRequest) {
	request = &DescribeAuditTracksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTracks")
	return
}

func NewDescribeAuditTracksResponse() (response *DescribeAuditTracksResponse) {
	response = &DescribeAuditTracksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询跟踪集列表。
func (c *Client) DescribeAuditTracks(request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
	if request == nil {
		request = NewDescribeAuditTracksRequest()
	}
	response = NewDescribeAuditTracksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
	request = &DescribeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeEvents")
	return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
	response = &DescribeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索日志接口
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
	if request == nil {
		request = NewDescribeEventsRequest()
	}
	response = NewDescribeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewGetExportTypeRequest() (request *GetExportTypeRequest) {
	request = &GetExportTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "GetExportType")
	return
}

func NewGetExportTypeResponse() (response *GetExportTypeResponse) {
	response = &GetExportTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署存储桶类型
func (c *Client) GetExportType(request *GetExportTypeRequest) (response *GetExportTypeResponse, err error) {
	if request == nil {
		request = NewGetExportTypeRequest()
	}
	response = NewGetExportTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuditTrackRequest() (request *ModifyAuditTrackRequest) {
	request = &ModifyAuditTrackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cloudaudit", APIVersion, "ModifyAuditTrack")
	return
}

func NewModifyAuditTrackResponse() (response *ModifyAuditTrackResponse) {
	response = &ModifyAuditTrackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改跟踪集。
func (c *Client) ModifyAuditTrack(request *ModifyAuditTrackRequest) (response *ModifyAuditTrackResponse, err error) {
	if request == nil {
		request = NewModifyAuditTrackRequest()
	}
	response = NewModifyAuditTrackResponse()
	err = c.Send(request, response)
	return
}

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

package v20230616

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2023-06-16"

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

func NewDescribeDefaultNoticeTmplRequest() (request *DescribeDefaultNoticeTmplRequest) {
	request = &DescribeDefaultNoticeTmplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeDefaultNoticeTmpl")
	return
}

func NewDescribeDefaultNoticeTmplResponse() (response *DescribeDefaultNoticeTmplResponse) {
	response = &DescribeDefaultNoticeTmplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取默认自定义通知模板
func (c *Client) DescribeDefaultNoticeTmpl(request *DescribeDefaultNoticeTmplRequest) (response *DescribeDefaultNoticeTmplResponse, err error) {
	if request == nil {
		request = NewDescribeDefaultNoticeTmplRequest()
	}
	response = NewDescribeDefaultNoticeTmplResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNoticeContentTmplRequest() (request *DescribeNoticeContentTmplRequest) {
	request = &DescribeNoticeContentTmplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeNoticeContentTmpl")
	return
}

func NewDescribeNoticeContentTmplResponse() (response *DescribeNoticeContentTmplResponse) {
	response = &DescribeNoticeContentTmplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义的通知内容模版
func (c *Client) DescribeNoticeContentTmpl(request *DescribeNoticeContentTmplRequest) (response *DescribeNoticeContentTmplResponse, err error) {
	if request == nil {
		request = NewDescribeNoticeContentTmplRequest()
	}
	response = NewDescribeNoticeContentTmplResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNoticeContentTmplRequest() (request *CreateNoticeContentTmplRequest) {
	request = &CreateNoticeContentTmplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "CreateNoticeContentTmpl")
	return
}

func NewCreateNoticeContentTmplResponse() (response *CreateNoticeContentTmplResponse) {
	response = &CreateNoticeContentTmplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义的通知内容模版
func (c *Client) CreateNoticeContentTmpl(request *CreateNoticeContentTmplRequest) (response *CreateNoticeContentTmplResponse, err error) {
	if request == nil {
		request = NewCreateNoticeContentTmplRequest()
	}
	response = NewCreateNoticeContentTmplResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNoticeContentTmplsRequest() (request *DeleteNoticeContentTmplsRequest) {
	request = &DeleteNoticeContentTmplsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DeleteNoticeContentTmpls")
	return
}

func NewDeleteNoticeContentTmplsResponse() (response *DeleteNoticeContentTmplsResponse) {
	response = &DeleteNoticeContentTmplsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义的通知内容模版
func (c *Client) DeleteNoticeContentTmpls(request *DeleteNoticeContentTmplsRequest) (response *DeleteNoticeContentTmplsResponse, err error) {
	if request == nil {
		request = NewDeleteNoticeContentTmplsRequest()
	}
	response = NewDeleteNoticeContentTmplsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNoticeContentTmplRequest() (request *ModifyNoticeContentTmplRequest) {
	request = &ModifyNoticeContentTmplRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "ModifyNoticeContentTmpl")
	return
}

func NewModifyNoticeContentTmplResponse() (response *ModifyNoticeContentTmplResponse) {
	response = &ModifyNoticeContentTmplResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自定义的通知内容模版
func (c *Client) ModifyNoticeContentTmpl(request *ModifyNoticeContentTmplRequest) (response *ModifyNoticeContentTmplResponse, err error) {
	if request == nil {
		request = NewModifyNoticeContentTmplRequest()
	}
	response = NewModifyNoticeContentTmplResponse()
	err = c.Send(request, response)
	return
}

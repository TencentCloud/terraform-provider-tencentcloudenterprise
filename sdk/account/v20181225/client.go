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

package v20181225

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-12-25"

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

func NewDescribeUserGroupRequest() (request *DescribeUserGroupRequest) {
	request = &DescribeUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DescribeUserGroup")
	return
}

func NewDescribeUserGroupResponse() (response *DescribeUserGroupResponse) {
	response = &DescribeUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户组信息
func (c *Client) DescribeUserGroup(request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
	if request == nil {
		request = NewDescribeUserGroupRequest()
	}
	response = NewDescribeUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetAreaByUinRequest() (request *GetAreaByUinRequest) {
	request = &GetAreaByUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAreaByUin")
	return
}

func NewGetAreaByUinResponse() (response *GetAreaByUinResponse) {
	response = &GetAreaByUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetAreaByUin
func (c *Client) GetAreaByUin(request *GetAreaByUinRequest) (response *GetAreaByUinResponse, err error) {
	if request == nil {
		request = NewGetAreaByUinRequest()
	}
	response = NewGetAreaByUinResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppIdByUinRequest() (request *GetAppIdByUinRequest) {
	request = &GetAppIdByUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAppIdByUin")
	return
}

func NewGetAppIdByUinResponse() (response *GetAppIdByUinResponse) {
	response = &GetAppIdByUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetAppIdByUin
func (c *Client) GetAppIdByUin(request *GetAppIdByUinRequest) (response *GetAppIdByUinResponse, err error) {
	if request == nil {
		request = NewGetAppIdByUinRequest()
	}
	response = NewGetAppIdByUinResponse()
	err = c.Send(request, response)
	return
}

func NewBatchCheckWhitelistRequest() (request *BatchCheckWhitelistRequest) {
	request = &BatchCheckWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "BatchCheckWhitelist")
	return
}

func NewBatchCheckWhitelistResponse() (response *BatchCheckWhitelistResponse) {
	response = &BatchCheckWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BatchCheckWhitelist
func (c *Client) BatchCheckWhitelist(request *BatchCheckWhitelistRequest) (response *BatchCheckWhitelistResponse, err error) {
	if request == nil {
		request = NewBatchCheckWhitelistRequest()
	}
	response = NewBatchCheckWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNicknameRequest() (request *DescribeNicknameRequest) {
	request = &DescribeNicknameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DescribeNickname")
	return
}

func NewDescribeNicknameResponse() (response *DescribeNicknameResponse) {
	response = &DescribeNicknameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询昵称
func (c *Client) DescribeNickname(request *DescribeNicknameRequest) (response *DescribeNicknameResponse, err error) {
	if request == nil {
		request = NewDescribeNicknameRequest()
	}
	response = NewDescribeNicknameResponse()
	err = c.Send(request, response)
	return
}

func NewTestDescRequest() (request *TestDescRequest) {
	request = &TestDescRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "TestDesc")
	return
}

func NewTestDescResponse() (response *TestDescResponse) {
	response = &TestDescResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// test
func (c *Client) TestDesc(request *TestDescRequest) (response *TestDescResponse, err error) {
	if request == nil {
		request = NewTestDescRequest()
	}
	response = NewTestDescResponse()
	err = c.Send(request, response)
	return
}

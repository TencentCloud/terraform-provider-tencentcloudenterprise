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

package v20200102

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-01-02"

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

func NewVerifyMenuRequest() (request *VerifyMenuRequest) {
	request = &VerifyMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "VerifyMenu")
	return
}

func NewVerifyMenuResponse() (response *VerifyMenuResponse) {
	response = &VerifyMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 菜单鉴权
func (c *Client) VerifyMenu(request *VerifyMenuRequest) (response *VerifyMenuResponse, err error) {
	if request == nil {
		request = NewVerifyMenuRequest()
	}
	response = NewVerifyMenuResponse()
	err = c.Send(request, response)
	return
}

func NewGetUidByUinArrRequest() (request *GetUidByUinArrRequest) {
	request = &GetUidByUinArrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUidByUinArr")
	return
}

func NewGetUidByUinArrResponse() (response *GetUidByUinArrResponse) {
	response = &GetUidByUinArrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin列表返回对应的uid信息
func (c *Client) GetUidByUinArr(request *GetUidByUinArrRequest) (response *GetUidByUinArrResponse, err error) {
	if request == nil {
		request = NewGetUidByUinArrRequest()
	}
	response = NewGetUidByUinArrResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRelatedRolePolicyRequest() (request *UpdateRelatedRolePolicyRequest) {
	request = &UpdateRelatedRolePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateRelatedRolePolicy")
	return
}

func NewUpdateRelatedRolePolicyResponse() (response *UpdateRelatedRolePolicyResponse) {
	response = &UpdateRelatedRolePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改可以扮演角色的uin列表
func (c *Client) UpdateRelatedRolePolicy(request *UpdateRelatedRolePolicyRequest) (response *UpdateRelatedRolePolicyResponse, err error) {
	if request == nil {
		request = NewUpdateRelatedRolePolicyRequest()
	}
	response = NewUpdateRelatedRolePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewBatchCheckUserHavePoliciesRequest() (request *BatchCheckUserHavePoliciesRequest) {
	request = &BatchCheckUserHavePoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "BatchCheckUserHavePolicies")
	return
}

func NewBatchCheckUserHavePoliciesResponse() (response *BatchCheckUserHavePoliciesResponse) {
	response = &BatchCheckUserHavePoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量检查用户是否拥有对应策略
func (c *Client) BatchCheckUserHavePolicies(request *BatchCheckUserHavePoliciesRequest) (response *BatchCheckUserHavePoliciesResponse, err error) {
	if request == nil {
		request = NewBatchCheckUserHavePoliciesRequest()
	}
	response = NewBatchCheckUserHavePoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewGetUinListByRolePolicyRequest() (request *GetUinListByRolePolicyRequest) {
	request = &GetUinListByRolePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUinListByRolePolicy")
	return
}

func NewGetUinListByRolePolicyResponse() (response *GetUinListByRolePolicyResponse) {
	response = &GetUinListByRolePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可以扮演角色的uin列表
func (c *Client) GetUinListByRolePolicy(request *GetUinListByRolePolicyRequest) (response *GetUinListByRolePolicyResponse, err error) {
	if request == nil {
		request = NewGetUinListByRolePolicyRequest()
	}
	response = NewGetUinListByRolePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetRoleListByOwnerUinRequest() (request *GetRoleListByOwnerUinRequest) {
	request = &GetRoleListByOwnerUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetRoleListByOwnerUin")
	return
}

func NewGetRoleListByOwnerUinResponse() (response *GetRoleListByOwnerUinResponse) {
	response = &GetRoleListByOwnerUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主账号可扮演角色列表
func (c *Client) GetRoleListByOwnerUin(request *GetRoleListByOwnerUinRequest) (response *GetRoleListByOwnerUinResponse, err error) {
	if request == nil {
		request = NewGetRoleListByOwnerUinRequest()
	}
	response = NewGetRoleListByOwnerUinResponse()
	err = c.Send(request, response)
	return
}

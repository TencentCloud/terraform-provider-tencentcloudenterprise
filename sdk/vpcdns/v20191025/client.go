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

package v20191025

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-10-25"

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

func NewCreateVpcDnsRecordRequest() (request *CreateVpcDnsRecordRequest) {
	request = &CreateVpcDnsRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateVpcDnsRecord")
	return
}

func NewCreateVpcDnsRecordResponse() (response *CreateVpcDnsRecordResponse) {
	response = &CreateVpcDnsRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建vpcdns记录
func (c *Client) CreateVpcDnsRecord(request *CreateVpcDnsRecordRequest) (response *CreateVpcDnsRecordResponse, err error) {
	if request == nil {
		request = NewCreateVpcDnsRecordRequest()
	}
	response = NewCreateVpcDnsRecordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcDnsDomainRequest() (request *CreateVpcDnsDomainRequest) {
	request = &CreateVpcDnsDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateVpcDnsDomain")
	return
}

func NewCreateVpcDnsDomainResponse() (response *CreateVpcDnsDomainResponse) {
	response = &CreateVpcDnsDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建vpcdns域名
func (c *Client) CreateVpcDnsDomain(request *CreateVpcDnsDomainRequest) (response *CreateVpcDnsDomainResponse, err error) {
	if request == nil {
		request = NewCreateVpcDnsDomainRequest()
	}
	response = NewCreateVpcDnsDomainResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcDnsDomainRemarkRequest() (request *CreateVpcDnsDomainRemarkRequest) {
	request = &CreateVpcDnsDomainRemarkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateVpcDnsDomainRemark")
	return
}

func NewCreateVpcDnsDomainRemarkResponse() (response *CreateVpcDnsDomainRemarkResponse) {
	response = &CreateVpcDnsDomainRemarkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建vpcdns域名备注
func (c *Client) CreateVpcDnsDomainRemark(request *CreateVpcDnsDomainRemarkRequest) (response *CreateVpcDnsDomainRemarkResponse, err error) {
	if request == nil {
		request = NewCreateVpcDnsDomainRemarkRequest()
	}
	response = NewCreateVpcDnsDomainRemarkResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcDnsRecordRequest() (request *DeleteVpcDnsRecordRequest) {
	request = &DeleteVpcDnsRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DeleteVpcDnsRecord")
	return
}

func NewDeleteVpcDnsRecordResponse() (response *DeleteVpcDnsRecordResponse) {
	response = &DeleteVpcDnsRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除vpcdns记录
func (c *Client) DeleteVpcDnsRecord(request *DeleteVpcDnsRecordRequest) (response *DeleteVpcDnsRecordResponse, err error) {
	if request == nil {
		request = NewDeleteVpcDnsRecordRequest()
	}
	response = NewDeleteVpcDnsRecordResponse()
	err = c.Send(request, response)
	return
}

func NewBindVpcDnsDomainRequest() (request *BindVpcDnsDomainRequest) {
	request = &BindVpcDnsDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "BindVpcDnsDomain")
	return
}

func NewBindVpcDnsDomainResponse() (response *BindVpcDnsDomainResponse) {
	response = &BindVpcDnsDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联VpcId
func (c *Client) BindVpcDnsDomain(request *BindVpcDnsDomainRequest) (response *BindVpcDnsDomainResponse, err error) {
	if request == nil {
		request = NewBindVpcDnsDomainRequest()
	}
	response = NewBindVpcDnsDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDnsDomainListRequest() (request *DescribeVpcDnsDomainListRequest) {
	request = &DescribeVpcDnsDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeVpcDnsDomainList")
	return
}

func NewDescribeVpcDnsDomainListResponse() (response *DescribeVpcDnsDomainListResponse) {
	response = &DescribeVpcDnsDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取vpcdns域名列表
func (c *Client) DescribeVpcDnsDomainList(request *DescribeVpcDnsDomainListRequest) (response *DescribeVpcDnsDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDnsDomainListRequest()
	}
	response = NewDescribeVpcDnsDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcDnsDomainRequest() (request *ModifyVpcDnsDomainRequest) {
	request = &ModifyVpcDnsDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyVpcDnsDomain")
	return
}

func NewModifyVpcDnsDomainResponse() (response *ModifyVpcDnsDomainResponse) {
	response = &ModifyVpcDnsDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改vpcdns域名
func (c *Client) ModifyVpcDnsDomain(request *ModifyVpcDnsDomainRequest) (response *ModifyVpcDnsDomainResponse, err error) {
	if request == nil {
		request = NewModifyVpcDnsDomainRequest()
	}
	response = NewModifyVpcDnsDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcDnsDomainRequest() (request *DeleteVpcDnsDomainRequest) {
	request = &DeleteVpcDnsDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DeleteVpcDnsDomain")
	return
}

func NewDeleteVpcDnsDomainResponse() (response *DeleteVpcDnsDomainResponse) {
	response = &DeleteVpcDnsDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除vpcdns域名
func (c *Client) DeleteVpcDnsDomain(request *DeleteVpcDnsDomainRequest) (response *DeleteVpcDnsDomainResponse, err error) {
	if request == nil {
		request = NewDeleteVpcDnsDomainRequest()
	}
	response = NewDeleteVpcDnsDomainResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcDnsRecordRequest() (request *ModifyVpcDnsRecordRequest) {
	request = &ModifyVpcDnsRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyVpcDnsRecord")
	return
}

func NewModifyVpcDnsRecordResponse() (response *ModifyVpcDnsRecordResponse) {
	response = &ModifyVpcDnsRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改vpcdns记录
func (c *Client) ModifyVpcDnsRecord(request *ModifyVpcDnsRecordRequest) (response *ModifyVpcDnsRecordResponse, err error) {
	if request == nil {
		request = NewModifyVpcDnsRecordRequest()
	}
	response = NewModifyVpcDnsRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDnsRecordListRequest() (request *DescribeVpcDnsRecordListRequest) {
	request = &DescribeVpcDnsRecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeVpcDnsRecordList")
	return
}

func NewDescribeVpcDnsRecordListResponse() (response *DescribeVpcDnsRecordListResponse) {
	response = &DescribeVpcDnsRecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取vpcdns记录列表
func (c *Client) DescribeVpcDnsRecordList(request *DescribeVpcDnsRecordListRequest) (response *DescribeVpcDnsRecordListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDnsRecordListRequest()
	}
	response = NewDescribeVpcDnsRecordListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcDnsForwardRuleRequest() (request *CreateVpcDnsForwardRuleRequest) {
	request = &CreateVpcDnsForwardRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateForwardRule")
	return
}

func NewCreateVpcDnsForwardRuleResponse() (response *CreateVpcDnsForwardRuleResponse) {
	response = &CreateVpcDnsForwardRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpcDnsForwardRule(request *CreateVpcDnsForwardRuleRequest) (response *CreateVpcDnsForwardRuleResponse, err error) {
	if request == nil {
		request = NewCreateVpcDnsForwardRuleRequest()
	}
	response = NewCreateVpcDnsForwardRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDnsForwardRuleRequest() (request *DescribeVpcDnsForwardRuleRequest) {
	request = &DescribeVpcDnsForwardRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeForwardRuleList")
	return
}

func NewDescribeVpcDnsForwardRuleResponse() (response *DescribeVpcDnsForwardRuleResponse) {
	response = &DescribeVpcDnsForwardRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcDnsForwardRule(request *DescribeVpcDnsForwardRuleRequest) (response *DescribeVpcDnsForwardRuleResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDnsForwardRuleRequest()
	}
	response = NewDescribeVpcDnsForwardRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcDnsForwardRuleRequest() (request *ModifyVpcDnsForwardRuleRequest) {
	request = &ModifyVpcDnsForwardRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyForwardRule")
	return
}

func NewModifyVpcDnsForwardRuleResponse() (response *ModifyVpcDnsForwardRuleResponse) {
	response = &ModifyVpcDnsForwardRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpcDnsForwardRule(request *ModifyVpcDnsForwardRuleRequest) (response *ModifyVpcDnsForwardRuleResponse, err error) {
	if request == nil {
		request = NewModifyVpcDnsForwardRuleRequest()
	}
	response = NewModifyVpcDnsForwardRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcDnsForwardRuleRequest() (request *DeleteVpcDnsForwardRuleRequest) {
	request = &DeleteVpcDnsForwardRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DeleteForwardRule")
	return
}

func NewDeleteVpcDnsForwardRuleResponse() (response *DeleteVpcDnsForwardRuleResponse) {
	response = &DeleteVpcDnsForwardRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpcDnsForwardRule(request *DeleteVpcDnsForwardRuleRequest) (response *DeleteVpcDnsForwardRuleResponse, err error) {
	if request == nil {
		request = NewDeleteVpcDnsForwardRuleRequest()
	}
	response = NewDeleteVpcDnsForwardRuleResponse()
	err = c.Send(request, response)
	return
}
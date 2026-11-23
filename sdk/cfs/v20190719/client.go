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

package v20190719

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-07-19"

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

func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
	request = &DescribeCfsFileSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystems")
	return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
	response = &DescribeCfsFileSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsFileSystems）用于查询文件系统
func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemsRequest()
	}
	response = NewDescribeCfsFileSystemsResponse()
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

func NewDeleteAutoSnapshotPolicyRequest() (request *DeleteAutoSnapshotPolicyRequest) {
	request = &DeleteAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DeleteAutoSnapshotPolicy")
	return
}

func NewDeleteAutoSnapshotPolicyResponse() (response *DeleteAutoSnapshotPolicyResponse) {
	response = &DeleteAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除快照策略
func (c *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (response *DeleteAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteAutoSnapshotPolicyRequest()
	}
	response = NewDeleteAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableZoneInfoRequest() (request *DescribeAvailableZoneInfoRequest) {
	request = &DescribeAvailableZoneInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeAvailableZoneInfo")
	return
}

func NewDescribeAvailableZoneInfoResponse() (response *DescribeAvailableZoneInfoResponse) {
	response = &DescribeAvailableZoneInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAvailableZoneInfo）用于查询区域的可用情况。
func (c *Client) DescribeAvailableZoneInfo(request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableZoneInfoRequest()
	}
	response = NewDescribeAvailableZoneInfoResponse()
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

func NewDeleteCfsSnapRequest() (request *DeleteCfsSnapRequest) {
	request = &DeleteCfsSnapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsSnap")
	return
}

func NewDeleteCfsSnapResponse() (response *DeleteCfsSnapResponse) {
	response = &DeleteCfsSnapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cfs快照
func (c *Client) DeleteCfsSnap(request *DeleteCfsSnapRequest) (response *DeleteCfsSnapResponse, err error) {
	if request == nil {
		request = NewDeleteCfsSnapRequest()
	}
	response = NewDeleteCfsSnapResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCfsSnapNameRequest() (request *UpdateCfsSnapNameRequest) {
	request = &UpdateCfsSnapNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsSnapName")
	return
}

func NewUpdateCfsSnapNameResponse() (response *UpdateCfsSnapNameResponse) {
	response = &UpdateCfsSnapNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改快照名称
func (c *Client) UpdateCfsSnapName(request *UpdateCfsSnapNameRequest) (response *UpdateCfsSnapNameResponse, err error) {
	if request == nil {
		request = NewUpdateCfsSnapNameRequest()
	}
	response = NewUpdateCfsSnapNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
	request = &DescribeMountTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeMountTargets")
	return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
	response = &DescribeMountTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeMountTargets）用于查询文件系统挂载点信息
func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
	if request == nil {
		request = NewDescribeMountTargetsRequest()
	}
	response = NewDescribeMountTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewCopyFileSystemSnapshotCrossRegionRequest() (request *CopyFileSystemSnapshotCrossRegionRequest) {
	request = &CopyFileSystemSnapshotCrossRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CopyFileSystemSnapshotCrossRegion")
	return
}

func NewCopyFileSystemSnapshotCrossRegionResponse() (response *CopyFileSystemSnapshotCrossRegionResponse) {
	response = &CopyFileSystemSnapshotCrossRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 跨地域复制文件系统快照
func (c *Client) CopyFileSystemSnapshotCrossRegion(request *CopyFileSystemSnapshotCrossRegionRequest) (response *CopyFileSystemSnapshotCrossRegionResponse, err error) {
	if request == nil {
		request = NewCopyFileSystemSnapshotCrossRegionRequest()
	}
	response = NewCopyFileSystemSnapshotCrossRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsFileSystemsByVpcRequest() (request *DescribeCfsFileSystemsByVpcRequest) {
	request = &DescribeCfsFileSystemsByVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystemsByVpc")
	return
}

func NewDescribeCfsFileSystemsByVpcResponse() (response *DescribeCfsFileSystemsByVpcResponse) {
	response = &DescribeCfsFileSystemsByVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用VPC信息查询文件系统实例
func (c *Client) DescribeCfsFileSystemsByVpc(request *DescribeCfsFileSystemsByVpcRequest) (response *DescribeCfsFileSystemsByVpcResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemsByVpcRequest()
	}
	response = NewDescribeCfsFileSystemsByVpcResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
	request = &CreateAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CreateAutoSnapshotPolicy")
	return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
	response = &CreateAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建定期快照策略
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAutoSnapshotPolicyRequest()
	}
	response = NewCreateAutoSnapshotPolicyResponse()
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

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
	request = &BindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "BindAutoSnapshotPolicy")
	return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
	response = &BindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件系统绑定快照策略
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewBindAutoSnapshotPolicyRequest()
	}
	response = NewBindAutoSnapshotPolicyResponse()
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

func NewDescribeCfsRulesRequest() (request *DescribeCfsRulesRequest) {
	request = &DescribeCfsRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsRules")
	return
}

func NewDescribeCfsRulesResponse() (response *DescribeCfsRulesResponse) {
	response = &DescribeCfsRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsRules）用于查询权限组规则列表。
func (c *Client) DescribeCfsRules(request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
	if request == nil {
		request = NewDescribeCfsRulesRequest()
	}
	response = NewDescribeCfsRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsTagsRequest() (request *DescribeCfsTagsRequest) {
	request = &DescribeCfsTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsTags")
	return
}

func NewDescribeCfsTagsResponse() (response *DescribeCfsTagsResponse) {
	response = &DescribeCfsTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cfs资源池列表
func (c *Client) DescribeCfsTags(request *DescribeCfsTagsRequest) (response *DescribeCfsTagsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsTagsRequest()
	}
	response = NewDescribeCfsTagsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCfsSnapRequest() (request *CreateCfsSnapRequest) {
	request = &CreateCfsSnapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsSnap")
	return
}

func NewCreateCfsSnapResponse() (response *CreateCfsSnapResponse) {
	response = &CreateCfsSnapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建cfs快照
func (c *Client) CreateCfsSnap(request *CreateCfsSnapRequest) (response *CreateCfsSnapResponse, err error) {
	if request == nil {
		request = NewCreateCfsSnapRequest()
	}
	response = NewCreateCfsSnapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVipStatusRequest() (request *DescribeVipStatusRequest) {
	request = &DescribeVipStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeVipStatus")
	return
}

func NewDescribeVipStatusResponse() (response *DescribeVipStatusResponse) {
	response = &DescribeVipStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内部接口，只供控制台使用
func (c *Client) DescribeVipStatus(request *DescribeVipStatusRequest) (response *DescribeVipStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVipStatusRequest()
	}
	response = NewDescribeVipStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsSnapsRequest() (request *DescribeCfsSnapsRequest) {
	request = &DescribeCfsSnapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsSnaps")
	return
}

func NewDescribeCfsSnapsResponse() (response *DescribeCfsSnapsResponse) {
	response = &DescribeCfsSnapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统快照信息
func (c *Client) DescribeCfsSnaps(request *DescribeCfsSnapsRequest) (response *DescribeCfsSnapsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsSnapsRequest()
	}
	response = NewDescribeCfsSnapsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfsServiceStatusRequest() (request *DescribeCfsServiceStatusRequest) {
	request = &DescribeCfsServiceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsServiceStatus")
	return
}

func NewDescribeCfsServiceStatusResponse() (response *DescribeCfsServiceStatusResponse) {
	response = &DescribeCfsServiceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsServiceStatus）用于查询用户使用CFS的服务状态。
func (c *Client) DescribeCfsServiceStatus(request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCfsServiceStatusRequest()
	}
	response = NewDescribeCfsServiceStatusResponse()
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

func NewUpdateAutoSnapshotPolicyRequest() (request *UpdateAutoSnapshotPolicyRequest) {
	request = &UpdateAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UpdateAutoSnapshotPolicy")
	return
}

func NewUpdateAutoSnapshotPolicyResponse() (response *UpdateAutoSnapshotPolicyResponse) {
	response = &UpdateAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新定期自动快照策略
func (c *Client) UpdateAutoSnapshotPolicy(request *UpdateAutoSnapshotPolicyRequest) (response *UpdateAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewUpdateAutoSnapshotPolicyRequest()
	}
	response = NewUpdateAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
	request = &DescribeAutoSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeAutoSnapshotPolicies")
	return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
	response = &DescribeAutoSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统快照d定期策略列表信息
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoSnapshotPoliciesRequest()
	}
	response = NewDescribeAutoSnapshotPoliciesResponse()
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

func NewDescribeCfsPGroupsRequest() (request *DescribeCfsPGroupsRequest) {
	request = &DescribeCfsPGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsPGroups")
	return
}

func NewDescribeCfsPGroupsResponse() (response *DescribeCfsPGroupsResponse) {
	response = &DescribeCfsPGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCfsPGroups）用于查询权限组列表。
func (c *Client) DescribeCfsPGroups(request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsPGroupsRequest()
	}
	response = NewDescribeCfsPGroupsResponse()
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

func NewDescribeCfsFileSystemClientsRequest() (request *DescribeCfsFileSystemClientsRequest) {
	request = &DescribeCfsFileSystemClientsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystemClients")
	return
}

func NewDescribeCfsFileSystemClientsResponse() (response *DescribeCfsFileSystemClientsResponse) {
	response = &DescribeCfsFileSystemClientsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询挂载该文件系统的客户端。此功能需要客户端安装CFS监控插件。
func (c *Client) DescribeCfsFileSystemClients(request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemClientsRequest()
	}
	response = NewDescribeCfsFileSystemClientsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticsRequest() (request *DescribeStatisticsRequest) {
	request = &DescribeStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeStatistics")
	return
}

func NewDescribeStatisticsResponse() (response *DescribeStatisticsResponse) {
	response = &DescribeStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 不对外，仅供控制台使用
func (c *Client) DescribeStatistics(request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticsRequest()
	}
	response = NewDescribeStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
	request = &UnbindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "UnbindAutoSnapshotPolicy")
	return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
	response = &UnbindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑快照策略
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewUnbindAutoSnapshotPolicyRequest()
	}
	response = NewUnbindAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMountTargetsWithRegionRequest() (request *DescribeMountTargetsWithRegionRequest) {
	request = &DescribeMountTargetsWithRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "DescribeMountTargetsWithRegion")
	return
}

func NewDescribeMountTargetsWithRegionResponse() (response *DescribeMountTargetsWithRegionResponse) {
	response = &DescribeMountTargetsWithRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询区域挂载点情况，只用于控制台。
func (c *Client) DescribeMountTargetsWithRegion(request *DescribeMountTargetsWithRegionRequest) (response *DescribeMountTargetsWithRegionResponse, err error) {
	if request == nil {
		request = NewDescribeMountTargetsWithRegionRequest()
	}
	response = NewDescribeMountTargetsWithRegionResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackCfsSnapRequest() (request *RollbackCfsSnapRequest) {
	request = &RollbackCfsSnapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "RollbackCfsSnap")
	return
}

func NewRollbackCfsSnapResponse() (response *RollbackCfsSnapResponse) {
	response = &RollbackCfsSnapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚快照到fs
func (c *Client) RollbackCfsSnap(request *RollbackCfsSnapRequest) (response *RollbackCfsSnapResponse, err error) {
	if request == nil {
		request = NewRollbackCfsSnapRequest()
	}
	response = NewRollbackCfsSnapResponse()
	err = c.Send(request, response)
	return
}

func NewOverrideCfsRulesRequest() (request *OverrideCfsRulesRequest) {
	request = &OverrideCfsRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfs", APIVersion, "OverrideCfsRules")
	return
}

func NewOverrideCfsRulesResponse() (response *OverrideCfsRulesResponse) {
	response = &OverrideCfsRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（OverrideCfsRules）用于批量覆盖式创建权限组规则。
func (c *Client) OverrideCfsRules(request *OverrideCfsRulesRequest) (response *OverrideCfsRulesResponse, err error) {
	if request == nil {
		request = NewOverrideCfsRulesRequest()
	}
	response = NewOverrideCfsRulesResponse()
	err = c.Send(request, response)
	return
}

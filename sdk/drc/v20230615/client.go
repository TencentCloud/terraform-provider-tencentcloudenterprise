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

package v20230615

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2023-06-15"

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

// CreateDisasterRecoverySitePair
func NewCreateDisasterRecoverySitePairRequest() (request *CreateDisasterRecoverySitePairRequest) {
	request = &CreateDisasterRecoverySitePairRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateDisasterRecoverySitePair")
	return
}

func NewCreateDisasterRecoverySitePairResponse() (response *CreateDisasterRecoverySitePairResponse) {
	response = &CreateDisasterRecoverySitePairResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDisasterRecoverySitePair(request *CreateDisasterRecoverySitePairRequest) (response *CreateDisasterRecoverySitePairResponse, err error) {
	if request == nil {
		request = NewCreateDisasterRecoverySitePairRequest()
	}
	response = NewCreateDisasterRecoverySitePairResponse()
	err = c.Send(request, response)
	return
}

// CreateDrcProtectGroup
func NewCreateDrcProtectGroupRequest() (request *CreateDrcProtectGroupRequest) {
	request = &CreateDrcProtectGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateDrcProtectGroup")
	return
}

func NewCreateDrcProtectGroupResponse() (response *CreateDrcProtectGroupResponse) {
	response = &CreateDrcProtectGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDrcProtectGroup(request *CreateDrcProtectGroupRequest) (response *CreateDrcProtectGroupResponse, err error) {
	if request == nil {
		request = NewCreateDrcProtectGroupRequest()
	}
	response = NewCreateDrcProtectGroupResponse()
	err = c.Send(request, response)
	return
}

// CreateDrcSecurityGroupMapping
func NewCreateDrcSecurityGroupMappingRequest() (request *CreateDrcSecurityGroupMappingRequest) {
	request = &CreateDrcSecurityGroupMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateDrcSecurityGroupMapping")
	return
}

func NewCreateDrcSecurityGroupMappingResponse() (response *CreateDrcSecurityGroupMappingResponse) {
	response = &CreateDrcSecurityGroupMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDrcSecurityGroupMapping(request *CreateDrcSecurityGroupMappingRequest) (response *CreateDrcSecurityGroupMappingResponse, err error) {
	if request == nil {
		request = NewCreateDrcSecurityGroupMappingRequest()
	}
	response = NewCreateDrcSecurityGroupMappingResponse()
	err = c.Send(request, response)
	return
}

// CreateDrcVpcMapping
func NewCreateDrcVpcMappingRequest() (request *CreateDrcVpcMappingRequest) {
	request = &CreateDrcVpcMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateDrcVpcMapping")
	return
}

func NewCreateDrcVpcMappingResponse() (response *CreateDrcVpcMappingResponse) {
	response = &CreateDrcVpcMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateDrcVpcMapping(request *CreateDrcVpcMappingRequest) (response *CreateDrcVpcMappingResponse, err error) {
	if request == nil {
		request = NewCreateDrcVpcMappingRequest()
	}
	response = NewCreateDrcVpcMappingResponse()
	err = c.Send(request, response)
	return
}

// CreateFileSystemCopyPair
func NewCreateFileSystemCopyPairRequest() (request *CreateFileSystemCopyPairRequest) {
	request = &CreateFileSystemCopyPairRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateFileSystemCopyPair")
	return
}

func NewCreateFileSystemCopyPairResponse() (response *CreateFileSystemCopyPairResponse) {
	response = &CreateFileSystemCopyPairResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileSystemCopyPair(request *CreateFileSystemCopyPairRequest) (response *CreateFileSystemCopyPairResponse, err error) {
	if request == nil {
		request = NewCreateFileSystemCopyPairRequest()
	}
	response = NewCreateFileSystemCopyPairResponse()
	err = c.Send(request, response)
	return
}

// CreateFileSystemDrillPairs
func NewCreateFileSystemDrillPairsRequest() (request *CreateFileSystemDrillPairsRequest) {
	request = &CreateFileSystemDrillPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "CreateFileSystemDrillPairs")
	return
}

func NewCreateFileSystemDrillPairsResponse() (response *CreateFileSystemDrillPairsResponse) {
	response = &CreateFileSystemDrillPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateFileSystemDrillPairs(request *CreateFileSystemDrillPairsRequest) (response *CreateFileSystemDrillPairsResponse, err error) {
	if request == nil {
		request = NewCreateFileSystemDrillPairsRequest()
	}
	response = NewCreateFileSystemDrillPairsResponse()
	err = c.Send(request, response)
	return
}

// DeleteCopyPairs
func NewDeleteCopyPairsRequest() (request *DeleteCopyPairsRequest) {
	request = &DeleteCopyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteCopyPairs")
	return
}

func NewDeleteCopyPairsResponse() (response *DeleteCopyPairsResponse) {
	response = &DeleteCopyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCopyPairs(request *DeleteCopyPairsRequest) (response *DeleteCopyPairsResponse, err error) {
	if request == nil {
		request = NewDeleteCopyPairsRequest()
	}
	response = NewDeleteCopyPairsResponse()
	err = c.Send(request, response)
	return
}

// DeleteDisasterRecoverySitePairs
func NewDeleteDisasterRecoverySitePairsRequest() (request *DeleteDisasterRecoverySitePairsRequest) {
	request = &DeleteDisasterRecoverySitePairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteDisasterRecoverySitePairs")
	return
}

func NewDeleteDisasterRecoverySitePairsResponse() (response *DeleteDisasterRecoverySitePairsResponse) {
	response = &DeleteDisasterRecoverySitePairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDisasterRecoverySitePairs(request *DeleteDisasterRecoverySitePairsRequest) (response *DeleteDisasterRecoverySitePairsResponse, err error) {
	if request == nil {
		request = NewDeleteDisasterRecoverySitePairsRequest()
	}
	response = NewDeleteDisasterRecoverySitePairsResponse()
	err = c.Send(request, response)
	return
}

// DeleteDrcProtectGroups
func NewDeleteDrcProtectGroupsRequest() (request *DeleteDrcProtectGroupsRequest) {
	request = &DeleteDrcProtectGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteDrcProtectGroups")
	return
}

func NewDeleteDrcProtectGroupsResponse() (response *DeleteDrcProtectGroupsResponse) {
	response = &DeleteDrcProtectGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDrcProtectGroups(request *DeleteDrcProtectGroupsRequest) (response *DeleteDrcProtectGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteDrcProtectGroupsRequest()
	}
	response = NewDeleteDrcProtectGroupsResponse()
	err = c.Send(request, response)
	return
}

// DeleteDrcSecurityGroupMapping
func NewDeleteDrcSecurityGroupMappingRequest() (request *DeleteDrcSecurityGroupMappingRequest) {
	request = &DeleteDrcSecurityGroupMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteDrcSecurityGroupMapping")
	return
}

func NewDeleteDrcSecurityGroupMappingResponse() (response *DeleteDrcSecurityGroupMappingResponse) {
	response = &DeleteDrcSecurityGroupMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDrcSecurityGroupMapping(request *DeleteDrcSecurityGroupMappingRequest) (response *DeleteDrcSecurityGroupMappingResponse, err error) {
	if request == nil {
		request = NewDeleteDrcSecurityGroupMappingRequest()
	}
	response = NewDeleteDrcSecurityGroupMappingResponse()
	err = c.Send(request, response)
	return
}

// DeleteDrcVpcMapping
func NewDeleteDrcVpcMappingRequest() (request *DeleteDrcVpcMappingRequest) {
	request = &DeleteDrcVpcMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteDrcVpcMapping")
	return
}

func NewDeleteDrcVpcMappingResponse() (response *DeleteDrcVpcMappingResponse) {
	response = &DeleteDrcVpcMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDrcVpcMapping(request *DeleteDrcVpcMappingRequest) (response *DeleteDrcVpcMappingResponse, err error) {
	if request == nil {
		request = NewDeleteDrcVpcMappingRequest()
	}
	response = NewDeleteDrcVpcMappingResponse()
	err = c.Send(request, response)
	return
}

// DeleteDrillPairs
func NewDeleteDrillPairsRequest() (request *DeleteDrillPairsRequest) {
	request = &DeleteDrillPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DeleteDrillPairs")
	return
}

func NewDeleteDrillPairsResponse() (response *DeleteDrillPairsResponse) {
	response = &DeleteDrillPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DeleteDrillPairs(request *DeleteDrillPairsRequest) (response *DeleteDrillPairsResponse, err error) {
	if request == nil {
		request = NewDeleteDrillPairsRequest()
	}
	response = NewDeleteDrillPairsResponse()
	err = c.Send(request, response)
	return
}

// DescribeCfsFileSystems
func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
	request = &DescribeCfsFileSystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeCfsFileSystems")
	return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
	response = &DescribeCfsFileSystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsFileSystemsRequest()
	}
	response = NewDescribeCfsFileSystemsResponse()
	err = c.Send(request, response)
	return
}

// DescribeCfsPGroups
func NewDescribeCfsPGroupsRequest() (request *DescribeCfsPGroupsRequest) {
	request = &DescribeCfsPGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeCfsPGroups")
	return
}

func NewDescribeCfsPGroupsResponse() (response *DescribeCfsPGroupsResponse) {
	response = &DescribeCfsPGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfsPGroups(request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsPGroupsRequest()
	}
	response = NewDescribeCfsPGroupsResponse()
	err = c.Send(request, response)
	return
}

// DescribeCfsTags
func NewDescribeCfsTagsRequest() (request *DescribeCfsTagsRequest) {
	request = &DescribeCfsTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeCfsTags")
	return
}

func NewDescribeCfsTagsResponse() (response *DescribeCfsTagsResponse) {
	response = &DescribeCfsTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCfsTags(request *DescribeCfsTagsRequest) (response *DescribeCfsTagsResponse, err error) {
	if request == nil {
		request = NewDescribeCfsTagsRequest()
	}
	response = NewDescribeCfsTagsResponse()
	err = c.Send(request, response)
	return
}

// DescribeCopyPairs
func NewDescribeCopyPairsRequest() (request *DescribeCopyPairsRequest) {
	request = &DescribeCopyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeCopyPairs")
	return
}

func NewDescribeCopyPairsResponse() (response *DescribeCopyPairsResponse) {
	response = &DescribeCopyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCopyPairs(request *DescribeCopyPairsRequest) (response *DescribeCopyPairsResponse, err error) {
	if request == nil {
		request = NewDescribeCopyPairsRequest()
	}
	response = NewDescribeCopyPairsResponse()
	err = c.Send(request, response)
	return
}

// DescribeCopyPairsDeniedActions
func NewDescribeCopyPairsDeniedActionsRequest() (request *DescribeCopyPairsDeniedActionsRequest) {
	request = &DescribeCopyPairsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeCopyPairsDeniedActions")
	return
}

func NewDescribeCopyPairsDeniedActionsResponse() (response *DescribeCopyPairsDeniedActionsResponse) {
	response = &DescribeCopyPairsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCopyPairsDeniedActions(request *DescribeCopyPairsDeniedActionsRequest) (response *DescribeCopyPairsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeCopyPairsDeniedActionsRequest()
	}
	response = NewDescribeCopyPairsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDisasterRecoveryOperations
func NewDescribeDisasterRecoveryOperationsRequest() (request *DescribeDisasterRecoveryOperationsRequest) {
	request = &DescribeDisasterRecoveryOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDisasterRecoveryOperations")
	return
}

func NewDescribeDisasterRecoveryOperationsResponse() (response *DescribeDisasterRecoveryOperationsResponse) {
	response = &DescribeDisasterRecoveryOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisasterRecoveryOperations(request *DescribeDisasterRecoveryOperationsRequest) (response *DescribeDisasterRecoveryOperationsResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoveryOperationsRequest()
	}
	response = NewDescribeDisasterRecoveryOperationsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDisasterRecoverySitePairs
func NewDescribeDisasterRecoverySitePairsRequest() (request *DescribeDisasterRecoverySitePairsRequest) {
	request = &DescribeDisasterRecoverySitePairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDisasterRecoverySitePairs")
	return
}

func NewDescribeDisasterRecoverySitePairsResponse() (response *DescribeDisasterRecoverySitePairsResponse) {
	response = &DescribeDisasterRecoverySitePairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisasterRecoverySitePairs(request *DescribeDisasterRecoverySitePairsRequest) (response *DescribeDisasterRecoverySitePairsResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoverySitePairsRequest()
	}
	response = NewDescribeDisasterRecoverySitePairsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDisasterRecoverySitePairsDeniedActions
func NewDescribeDisasterRecoverySitePairsDeniedActionsRequest() (request *DescribeDisasterRecoverySitePairsDeniedActionsRequest) {
	request = &DescribeDisasterRecoverySitePairsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDisasterRecoverySitePairsDeniedActions")
	return
}

func NewDescribeDisasterRecoverySitePairsDeniedActionsResponse() (response *DescribeDisasterRecoverySitePairsDeniedActionsResponse) {
	response = &DescribeDisasterRecoverySitePairsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisasterRecoverySitePairsDeniedActions(request *DescribeDisasterRecoverySitePairsDeniedActionsRequest) (response *DescribeDisasterRecoverySitePairsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoverySitePairsDeniedActionsRequest()
	}
	response = NewDescribeDisasterRecoverySitePairsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDrcDrillGroups
func NewDescribeDrcDrillGroupsRequest() (request *DescribeDrcDrillGroupsRequest) {
	request = &DescribeDrcDrillGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDrcDrillGroups")
	return
}

func NewDescribeDrcDrillGroupsResponse() (response *DescribeDrcDrillGroupsResponse) {
	response = &DescribeDrcDrillGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDrcDrillGroups(request *DescribeDrcDrillGroupsRequest) (response *DescribeDrcDrillGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDrcDrillGroupsRequest()
	}
	response = NewDescribeDrcDrillGroupsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDrcProtectGroups
func NewDescribeDrcProtectGroupsRequest() (request *DescribeDrcProtectGroupsRequest) {
	request = &DescribeDrcProtectGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDrcProtectGroups")
	return
}

func NewDescribeDrcProtectGroupsResponse() (response *DescribeDrcProtectGroupsResponse) {
	response = &DescribeDrcProtectGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDrcProtectGroups(request *DescribeDrcProtectGroupsRequest) (response *DescribeDrcProtectGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDrcProtectGroupsRequest()
	}
	response = NewDescribeDrcProtectGroupsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDrillPairs
func NewDescribeDrillPairsRequest() (request *DescribeDrillPairsRequest) {
	request = &DescribeDrillPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDrillPairs")
	return
}

func NewDescribeDrillPairsResponse() (response *DescribeDrillPairsResponse) {
	response = &DescribeDrillPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDrillPairs(request *DescribeDrillPairsRequest) (response *DescribeDrillPairsResponse, err error) {
	if request == nil {
		request = NewDescribeDrillPairsRequest()
	}
	response = NewDescribeDrillPairsResponse()
	err = c.Send(request, response)
	return
}

// DescribeDrillPairsDeniedActions
func NewDescribeDrillPairsDeniedActionsRequest() (request *DescribeDrillPairsDeniedActionsRequest) {
	request = &DescribeDrillPairsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDrillPairsDeniedActions")
	return
}

func NewDescribeDrillPairsDeniedActionsResponse() (response *DescribeDrillPairsDeniedActionsResponse) {
	response = &DescribeDrillPairsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDrillPairsDeniedActions(request *DescribeDrillPairsDeniedActionsRequest) (response *DescribeDrillPairsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeDrillPairsDeniedActionsRequest()
	}
	response = NewDescribeDrillPairsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

// DescribeProtectGroupsDeniedActions
func NewDescribeProtectGroupsDeniedActionsRequest() (request *DescribeProtectGroupsDeniedActionsRequest) {
	request = &DescribeProtectGroupsDeniedActionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeProtectGroupsDeniedActions")
	return
}

func NewDescribeProtectGroupsDeniedActionsResponse() (response *DescribeProtectGroupsDeniedActionsResponse) {
	response = &DescribeProtectGroupsDeniedActionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeProtectGroupsDeniedActions(request *DescribeProtectGroupsDeniedActionsRequest) (response *DescribeProtectGroupsDeniedActionsResponse, err error) {
	if request == nil {
		request = NewDescribeProtectGroupsDeniedActionsRequest()
	}
	response = NewDescribeProtectGroupsDeniedActionsResponse()
	err = c.Send(request, response)
	return
}

// DescribeSecurityGroupMappings
func NewDescribeSecurityGroupMappingsRequest() (request *DescribeSecurityGroupMappingsRequest) {
	request = &DescribeSecurityGroupMappingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeSecurityGroupMappings")
	return
}

func NewDescribeSecurityGroupMappingsResponse() (response *DescribeSecurityGroupMappingsResponse) {
	response = &DescribeSecurityGroupMappingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroupMappings(request *DescribeSecurityGroupMappingsRequest) (response *DescribeSecurityGroupMappingsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupMappingsRequest()
	}
	response = NewDescribeSecurityGroupMappingsResponse()
	err = c.Send(request, response)
	return
}

// DescribeVipStatus
func NewDescribeVipStatusRequest() (request *DescribeVipStatusRequest) {
	request = &DescribeVipStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeVipStatus")
	return
}

func NewDescribeVipStatusResponse() (response *DescribeVipStatusResponse) {
	response = &DescribeVipStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVipStatus(request *DescribeVipStatusRequest) (response *DescribeVipStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVipStatusRequest()
	}
	response = NewDescribeVipStatusResponse()
	err = c.Send(request, response)
	return
}

// DescribeVpcMappings
func NewDescribeVpcMappingsRequest() (request *DescribeVpcMappingsRequest) {
	request = &DescribeVpcMappingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeVpcMappings")
	return
}

func NewDescribeVpcMappingsResponse() (response *DescribeVpcMappingsResponse) {
	response = &DescribeVpcMappingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcMappings(request *DescribeVpcMappingsRequest) (response *DescribeVpcMappingsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcMappingsRequest()
	}
	response = NewDescribeVpcMappingsResponse()
	err = c.Send(request, response)
	return
}

// FailoverCopyPairs
func NewFailoverCopyPairsRequest() (request *FailoverCopyPairsRequest) {
	request = &FailoverCopyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "FailoverCopyPairs")
	return
}

func NewFailoverCopyPairsResponse() (response *FailoverCopyPairsResponse) {
	response = &FailoverCopyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) FailoverCopyPairs(request *FailoverCopyPairsRequest) (response *FailoverCopyPairsResponse, err error) {
	if request == nil {
		request = NewFailoverCopyPairsRequest()
	}
	response = NewFailoverCopyPairsResponse()
	err = c.Send(request, response)
	return
}

// ModifyCopyPairAttribute
func NewModifyCopyPairAttributeRequest() (request *ModifyCopyPairAttributeRequest) {
	request = &ModifyCopyPairAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "ModifyCopyPairAttribute")
	return
}

func NewModifyCopyPairAttributeResponse() (response *ModifyCopyPairAttributeResponse) {
	response = &ModifyCopyPairAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyCopyPairAttribute(request *ModifyCopyPairAttributeRequest) (response *ModifyCopyPairAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCopyPairAttributeRequest()
	}
	response = NewModifyCopyPairAttributeResponse()
	err = c.Send(request, response)
	return
}

// ModifyDrillPairAttribute
func NewModifyDrillPairAttributeRequest() (request *ModifyDrillPairAttributeRequest) {
	request = &ModifyDrillPairAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "ModifyDrillPairAttribute")
	return
}

func NewModifyDrillPairAttributeResponse() (response *ModifyDrillPairAttributeResponse) {
	response = &ModifyDrillPairAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDrillPairAttribute(request *ModifyDrillPairAttributeRequest) (response *ModifyDrillPairAttributeResponse, err error) {
	if request == nil {
		request = NewModifyDrillPairAttributeRequest()
	}
	response = NewModifyDrillPairAttributeResponse()
	err = c.Send(request, response)
	return
}

// ModifyProtectGroupAttribute
func NewModifyProtectGroupAttributeRequest() (request *ModifyProtectGroupAttributeRequest) {
	request = &ModifyProtectGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "ModifyProtectGroupAttribute")
	return
}

func NewModifyProtectGroupAttributeResponse() (response *ModifyProtectGroupAttributeResponse) {
	response = &ModifyProtectGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifyProtectGroupAttribute(request *ModifyProtectGroupAttributeRequest) (response *ModifyProtectGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyProtectGroupAttributeRequest()
	}
	response = NewModifyProtectGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

// ModifySitePairAttribute
func NewModifySitePairAttributeRequest() (request *ModifySitePairAttributeRequest) {
	request = &ModifySitePairAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "ModifySitePairAttribute")
	return
}

func NewModifySitePairAttributeResponse() (response *ModifySitePairAttributeResponse) {
	response = &ModifySitePairAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ModifySitePairAttribute(request *ModifySitePairAttributeRequest) (response *ModifySitePairAttributeResponse, err error) {
	if request == nil {
		request = NewModifySitePairAttributeRequest()
	}
	response = NewModifySitePairAttributeResponse()
	err = c.Send(request, response)
	return
}

// RunCopyPairTasks
func NewRunCopyPairTasksRequest() (request *RunCopyPairTasksRequest) {
	request = &RunCopyPairTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "RunCopyPairTasks")
	return
}

func NewRunCopyPairTasksResponse() (response *RunCopyPairTasksResponse) {
	response = &RunCopyPairTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) RunCopyPairTasks(request *RunCopyPairTasksRequest) (response *RunCopyPairTasksResponse, err error) {
	if request == nil {
		request = NewRunCopyPairTasksRequest()
	}
	response = NewRunCopyPairTasksResponse()
	err = c.Send(request, response)
	return
}

// StopCopyPairTasks
func NewStopCopyPairTasksRequest() (request *StopCopyPairTasksRequest) {
	request = &StopCopyPairTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "StopCopyPairTasks")
	return
}

func NewStopCopyPairTasksResponse() (response *StopCopyPairTasksResponse) {
	response = &StopCopyPairTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) StopCopyPairTasks(request *StopCopyPairTasksRequest) (response *StopCopyPairTasksResponse, err error) {
	if request == nil {
		request = NewStopCopyPairTasksRequest()
	}
	response = NewStopCopyPairTasksResponse()
	err = c.Send(request, response)
	return
}

// DescribeDisasterRecoveryOverview
func NewDescribeDisasterRecoveryOverviewRequest() (request *DescribeDisasterRecoveryOverviewRequest) {
	request = &DescribeDisasterRecoveryOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drc", APIVersion, "DescribeDisasterRecoveryOverview")
	return
}

func NewDescribeDisasterRecoveryOverviewResponse() (response *DescribeDisasterRecoveryOverviewResponse) {
	response = &DescribeDisasterRecoveryOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisasterRecoveryOverview(request *DescribeDisasterRecoveryOverviewRequest) (response *DescribeDisasterRecoveryOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDisasterRecoveryOverviewRequest()
	}
	response = NewDescribeDisasterRecoveryOverviewResponse()
	err = c.Send(request, response)
	return
}
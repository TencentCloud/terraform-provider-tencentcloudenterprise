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
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// CreateDisasterRecoverySitePairRequest is the request struct for api CreateDisasterRecoverySitePair
type CreateDisasterRecoverySitePairRequest struct {
	*tchttp.BaseRequest
	SitePairName         *string `json:"SitePairName,omitempty" name:"SitePairName"`
	SourceZone           *string `json:"SourceZone,omitempty" name:"SourceZone"`
	TargetZone           *string `json:"TargetZone,omitempty" name:"TargetZone"`
	SecretId             *string `json:"SecretId,omitempty" name:"SecretId"`
	SecretKey            *string `json:"SecretKey,omitempty" name:"SecretKey"`
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitempty" name:"DisasterRecoveryType"`
	CopyType             *string `json:"CopyType,omitempty" name:"CopyType"`
	SourceVpc            *string `json:"SourceVpc,omitempty" name:"SourceVpc"`
	TargetVpc            *string `json:"TargetVpc,omitempty" name:"TargetVpc"`
	SitePairProductType  *string `json:"SitePairProductType,omitempty" name:"SitePairProductType"`
	SourceRegion         *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	PeerCloudName        *string `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	LocalCloudName       *string `json:"LocalCloudName,omitempty" name:"LocalCloudName"`
	TargetRegion         *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
}

// CreateDisasterRecoverySitePairResponse is the response struct for api CreateDisasterRecoverySitePair
type CreateDisasterRecoverySitePairResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		DisasterRecoverySitePairId *string `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
		RequestId                  *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateDrcProtectGroupRequest is the request struct for api CreateDrcProtectGroup
type CreateDrcProtectGroupRequest struct {
	*tchttp.BaseRequest
	SitePairId       *string `json:"SitePairId,omitempty" name:"SitePairId"`
	ProtectGroupName *string `json:"ProtectGroupName,omitempty" name:"ProtectGroupName"`
	ProtectGroupType *string `json:"ProtectGroupType,omitempty" name:"ProtectGroupType"`
	Rpo              *string `json:"Rpo,omitempty" name:"Rpo"`
	DataDirection    *string `json:"DataDirection,omitempty" name:"DataDirection"`
}

// CreateDrcProtectGroupResponse is the response struct for api CreateDrcProtectGroup
type CreateDrcProtectGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateDrcSecurityGroupMappingRequest is the request struct for api CreateDrcSecurityGroupMapping
type CreateDrcSecurityGroupMappingRequest struct {
	*tchttp.BaseRequest
	SitePairId            *string `json:"SitePairId,omitempty" name:"SitePairId"`
	IsAutoSyn             *bool   `json:"IsAutoSyn,omitempty" name:"IsAutoSyn"`
	TargetSecurityGroupId *string `json:"TargetSecurityGroupId,omitempty" name:"TargetSecurityGroupId"`
	SrcSecurityGroupId    *string `json:"SrcSecurityGroupId,omitempty" name:"SrcSecurityGroupId"`
}

// CreateDrcSecurityGroupMappingResponse is the response struct for api CreateDrcSecurityGroupMapping
type CreateDrcSecurityGroupMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateDrcVpcMappingRequest is the request struct for api CreateDrcVpcMapping
type CreateDrcVpcMappingRequest struct {
	*tchttp.BaseRequest
	SitePairId     *string `json:"SitePairId,omitempty" name:"SitePairId"`
	SourceVpcId    *string `json:"SourceVpcId,omitempty" name:"SourceVpcId"`
	SourceSubnetId *string `json:"SourceSubnetId,omitempty" name:"SourceSubnetId"`
	TargetVpcId    *string `json:"TargetVpcId,omitempty" name:"TargetVpcId"`
	TargetSubnetId *string `json:"TargetSubnetId,omitempty" name:"TargetSubnetId"`
}

// CreateDrcVpcMappingResponse is the response struct for api CreateDrcVpcMapping
type CreateDrcVpcMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateFileSystemCopyPairRequest is the request struct for api CreateFileSystemCopyPair
type CreateFileSystemCopyPairRequest struct {
	*tchttp.BaseRequest
	FilesystemCopyPairName    *string                      `json:"FilesystemCopyPairName,omitempty" name:"FilesystemCopyPairName"`
	TagId                     *uint64                      `json:"TagId,omitempty" name:"TagId"`
	ProjectId                 *string                      `json:"ProjectId,omitempty" name:"ProjectId"`
	ProtectGroupId            *string                      `json:"ProtectGroupId,omitempty" name:"ProtectGroupId"`
	Rpo                       *uint64                      `json:"Rpo,omitempty" name:"Rpo"`
	CreationToken             *string                      `json:"CreationToken,omitempty" name:"CreationToken"`
	CreateTargetCfsParameters []*CreateTargetCfsParameters `json:"CreateTargetCfsParameters,omitempty" name:"CreateTargetCfsParameters"`
	IsCrossCloud              *bool                        `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
	InternalDirectCreate      *bool                        `json:"InternalDirectCreate,omitempty" name:"InternalDirectCreate"`
}

// CreateFileSystemCopyPairResponse is the response struct for api CreateFileSystemCopyPair
type CreateFileSystemCopyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		CopyPairId *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
		RequestId  *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateFileSystemDrillPairsRequest is the request struct for api CreateFileSystemDrillPairs
type CreateFileSystemDrillPairsRequest struct {
	*tchttp.BaseRequest
	DrillPairGroupName        *string                      `json:"DrillPairGroupName,omitempty" name:"DrillPairGroupName"`
	TagId                     *uint64                      `json:"TagId,omitempty" name:"TagId"`
	ProjectId                 *string                      `json:"ProjectId,omitempty" name:"ProjectId"`
	ProtectGroupId            *string                      `json:"ProtectGroupId,omitempty" name:"ProtectGroupId"`
	CreationToken             *string                      `json:"CreationToken,omitempty" name:"CreationToken"`
	CreateTargetCfsParameters []*CreateTargetCfsParameters `json:"CreateTargetCfsParameters,omitempty" name:"CreateTargetCfsParameters"`
	IsCrossCloud              *bool                        `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
	DrillPairGroupVpc         *string                      `json:"DrillPairGroupVpc,omitempty" name:"DrillPairGroupVpc"`
}

// CreateFileSystemDrillPairsResponse is the response struct for api CreateFileSystemDrillPairs
type CreateFileSystemDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		DrillPairIds []*string `json:"DrillPairIds,omitempty" name:"DrillPairIds"`
		RequestId    *string   `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteCopyPairsRequest is the request struct for api DeleteCopyPairs
type DeleteCopyPairsRequest struct {
	*tchttp.BaseRequest
	CopyPairIds  []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
	// a typo in yunapi TODO: sync
	IsCrossCloudd *bool     `json:"IsCrossCloudd,omitempty" name:"IsCrossCloudd"`
}

// DeleteCopyPairsResponse is the response struct for api DeleteCopyPairs
type DeleteCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteDisasterRecoverySitePairsRequest is the request struct for api DeleteDisasterRecoverySitePairs
type DeleteDisasterRecoverySitePairsRequest struct {
	*tchttp.BaseRequest
	DisasterRecoverySitePairIds []*string `json:"DisasterRecoverySitePairIds,omitempty" name:"DisasterRecoverySitePairIds"`
	IsCrossCloud                *bool     `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
}

// DeleteDisasterRecoverySitePairsResponse is the response struct for api DeleteDisasterRecoverySitePairs
type DeleteDisasterRecoverySitePairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteDrcProtectGroupsRequest is the request struct for api DeleteDrcProtectGroups
type DeleteDrcProtectGroupsRequest struct {
	*tchttp.BaseRequest
	ProtectGroups []*string `json:"ProtectGroups,omitempty" name:"ProtectGroups"`
}

// DeleteDrcProtectGroupsResponse is the response struct for api DeleteDrcProtectGroups
type DeleteDrcProtectGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteDrcSecurityGroupMappingRequest is the request struct for api DeleteDrcSecurityGroupMapping
type DeleteDrcSecurityGroupMappingRequest struct {
	*tchttp.BaseRequest
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

// DeleteDrcSecurityGroupMappingResponse is the response struct for api DeleteDrcSecurityGroupMapping
type DeleteDrcSecurityGroupMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteDrcVpcMappingRequest is the request struct for api DeleteDrcVpcMapping
type DeleteDrcVpcMappingRequest struct {
	*tchttp.BaseRequest
	VpcMappingIds []*string `json:"VpcMappingIds,omitempty" name:"VpcMappingIds"`
}

// DeleteDrcVpcMappingResponse is the response struct for api DeleteDrcVpcMapping
type DeleteDrcVpcMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DeleteDrillPairsRequest is the request struct for api DeleteDrillPairs
type DeleteDrillPairsRequest struct {
	*tchttp.BaseRequest
	DrillPairIds       []*string `json:"DrillPairIds,omitempty" name:"DrillPairIds"`
	DrillPairType      *string   `json:"DrillPairType,omitempty" name:"DrillPairType"`
	IsDeleteDrillGroup *bool     `json:"IsDeleteDrillGroup,omitempty" name:"IsDeleteDrillGroup"`
}

// DeleteDrillPairsResponse is the response struct for api DeleteDrillPairs
type DeleteDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeCfsFileSystemsRequest is the request struct for api DescribeCfsFileSystems
type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest
	FileSystemId      *string   `json:"FileSystemId,omitempty" name:"FileSystemId"`
	InternationalFlag *uint64   `json:"InternationalFlag,omitempty" name:"InternationalFlag"`
	VpcId             *string   `json:"VpcId,omitempty" name:"VpcId"`
	SubnetId          *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	Offset            *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit             *uint64   `json:"Limit,omitempty" name:"Limit"`
	CreationToken     *string   `json:"CreationToken,omitempty" name:"CreationToken"`
	FileSystemIds     []*string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
	SourceRegion      *string   `json:"SourceRegion,omitempty" name:"SourceRegion"`
	PeerCloudName     *string   `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	ApiModule         *string   `json:"ApiModule,omitempty" name:"ApiModule"`
}

// DescribeCfsFileSystemsResponse is the response struct for api DescribeCfsFileSystems
type DescribeCfsFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeCfsPGroupsRequest is the request struct for api DescribeCfsPGroups
type DescribeCfsPGroupsRequest struct {
	*tchttp.BaseRequest
	InternationalFlag *int64    `json:"InternationalFlag,omitempty" name:"InternationalFlag"`
	PGroupId          *string   `json:"PGroupId,omitempty" name:"PGroupId"`
	Name              *string   `json:"Name,omitempty" name:"Name"`
	DescInfo          *string   `json:"DescInfo,omitempty" name:"DescInfo"`
	Offset            *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit             *uint64   `json:"Limit,omitempty" name:"Limit"`
	OrderBy           *string   `json:"OrderBy,omitempty" name:"OrderBy"`
	OrderByType       *string   `json:"OrderByType,omitempty" name:"OrderByType"`
	PeerCloudName     *string   `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	ApiModule         *string   `json:"ApiModule,omitempty" name:"ApiModule"`
	SourceRegion      []*string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

// DescribeCfsPGroupsResponse is the response struct for api DescribeCfsPGroups
type DescribeCfsPGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeCfsTagsRequest is the request struct for api DescribeCfsTags
type DescribeCfsTagsRequest struct {
	*tchttp.BaseRequest
	InternationalFlag *int64  `json:"InternationalFlag,omitempty" name:"InternationalFlag"`
	PeerCloudName     *string `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	ApiModule         *string `json:"ApiModule,omitempty" name:"ApiModule"`
	SourceRegion      *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	ZoneId            *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

// DescribeCfsTagsResponse is the response struct for api DescribeCfsTags
type DescribeCfsTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeCopyPairsRequest is the request struct for api DescribeCopyPairs
type DescribeCopyPairsRequest struct {
	*tchttp.BaseRequest
	Filters             []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset              *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit               *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order               *string   `json:"Order,omitempty" name:"Order"`
	OrderField          *string   `json:"OrderField,omitempty" name:"OrderField"`
	CopyPairIds         []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType        *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
	QueryProtectionTime *bool     `json:"QueryProtectionTime,omitempty" name:"QueryProtectionTime"`
	IsCrossCloud        *bool     `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
}

// DescribeCopyPairsResponse is the response struct for api DescribeCopyPairs
type DescribeCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount  *uint64     `json:"TotalCount,omitempty" name:"TotalCount"`
		CopyPairSet []*CopyPair `json:"CopyPairSet,omitempty" name:"CopyPairSet"`
		RequestId   *string     `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeCopyPairsDeniedActionsRequest is the request struct for api DescribeCopyPairsDeniedActions
type DescribeCopyPairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	CopyPairIds  []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
}

// DescribeCopyPairsDeniedActionsResponse is the response struct for api DescribeCopyPairsDeniedActions
type DescribeCopyPairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		CopyPairDeniedActionSet []*CopyPairDeniedAction `json:"CopyPairDeniedActionSet,omitempty" name:"CopyPairDeniedActionSet"`
		RequestId               *string                 `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDisasterRecoveryOperationsRequest is the request struct for api DescribeDisasterRecoveryOperations
type DescribeDisasterRecoveryOperationsRequest struct {
	*tchttp.BaseRequest
	Filters      []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset       *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit        *uint64   `json:"Limit,omitempty" name:"Limit"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
}

// DescribeDisasterRecoveryOperationsResponse is the response struct for api DescribeDisasterRecoveryOperations
type DescribeDisasterRecoveryOperationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount      *uint64         `json:"TotalCount,omitempty" name:"TotalCount"`
		OperationLogSet []*OperationLog `json:"OperationLogSet,omitempty" name:"OperationLogSet"`
		RequestId       *string         `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDisasterRecoverySitePairsRequest is the request struct for api DescribeDisasterRecoverySitePairs
type DescribeDisasterRecoverySitePairsRequest struct {
	*tchttp.BaseRequest
	Filters      []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset       *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit        *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order        *string   `json:"Order,omitempty" name:"Order"`
	OrderField   *string   `json:"OrderField,omitempty" name:"OrderField"`
	SitePairIds  []*string `json:"SitePairIds,omitempty" name:"SitePairIds"`
	IsCrossCloud *bool     `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
	SitePairType *string   `json:"SitePairType,omitempty" name:"SitePairType"`
}

// DescribeDisasterRecoverySitePairsResponse is the response struct for api DescribeDisasterRecoverySitePairs
type DescribeDisasterRecoverySitePairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		SitePairSet []*DisasterRecoverySitePair `json:"SitePairSet,omitempty" name:"SitePairSet"`
		TotalCount  *uint64                     `json:"TotalCount,omitempty" name:"TotalCount"`
		RequestId   *string                     `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDisasterRecoverySitePairsDeniedActionsRequest is the request struct for api DescribeDisasterRecoverySitePairsDeniedActions
type DescribeDisasterRecoverySitePairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	SitePairIds  []*string `json:"SitePairIds,omitempty" name:"SitePairIds"`
	IsCrossCloud *bool     `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
}

// DescribeDisasterRecoverySitePairsDeniedActionsResponse is the response struct for api DescribeDisasterRecoverySitePairsDeniedActions
type DescribeDisasterRecoverySitePairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		DisasterRecoverySitePairDeniedActionSet []*DisasterRecoverySitePairDeniedAction `json:"DisasterRecoverySitePairDeniedActionSet,omitempty" name:"DisasterRecoverySitePairDeniedActionSet"`
		RequestId                               *string                                 `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDrcDrillGroupsRequest is the request struct for api DescribeDrcDrillGroups
type DescribeDrcDrillGroupsRequest struct {
	*tchttp.BaseRequest
	DrillGroupIds  []*string `json:"DrillGroupIds,omitempty" name:"DrillGroupIds"`
	DrillGroupType *string   `json:"DrillGroupType,omitempty" name:"DrillGroupType"`
	Filters        []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset         *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit          *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order          *string   `json:"Order,omitempty" name:"Order"`
	OrderField     *string   `json:"OrderField,omitempty" name:"OrderField"`
}

// DescribeDrcDrillGroupsResponse is the response struct for api DescribeDrcDrillGroups
type DescribeDrcDrillGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDrcProtectGroupsRequest is the request struct for api DescribeDrcProtectGroups
type DescribeDrcProtectGroupsRequest struct {
	*tchttp.BaseRequest
	ProtectGroupIds  []*string `json:"ProtectGroupIds,omitempty" name:"ProtectGroupIds"`
	ProtectGroupType *string   `json:"ProtectGroupType,omitempty" name:"ProtectGroupType"`
	Filters          []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset           *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit            *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order            *string   `json:"Order,omitempty" name:"Order"`
	OrderField       *string   `json:"OrderField,omitempty" name:"OrderField"`
	IsCrossCloud     *bool     `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
}

// DescribeDrcProtectGroupsResponse is the response struct for api DescribeDrcProtectGroups
type DescribeDrcProtectGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDrillPairsRequest is the request struct for api DescribeDrillPairs
type DescribeDrillPairsRequest struct {
	*tchttp.BaseRequest
	Filters          []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset           *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit            *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order            *string   `json:"Order,omitempty" name:"Order"`
	OrderField       *string   `json:"OrderField,omitempty" name:"OrderField"`
	DrillPairIds     []*string `json:"DrillPairIds,omitempty" name:"DrillPairIds"`
	DrillPairType    *string   `json:"DrillPairType,omitempty" name:"DrillPairType"`
	DrillPairGroupId *string   `json:"DrillPairGroupId,omitempty" name:"DrillPairGroupId"`
}

// DescribeDrillPairsResponse is the response struct for api DescribeDrillPairs
type DescribeDrillPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TotalCount   *uint64      `json:"TotalCount,omitempty" name:"TotalCount"`
		DrillPairSet []*DrillPair `json:"DrillPairSet,omitempty" name:"DrillPairSet"`
		RequestId    *string      `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDrillPairsDeniedActionsRequest is the request struct for api DescribeDrillPairsDeniedActions
type DescribeDrillPairsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	DrillPairIds  []*string `json:"DrillPairIds,omitempty" name:"DrillPairIds"`
	DrillPairType *string   `json:"DrillPairType,omitempty" name:"DrillPairType"`
}

// DescribeDrillPairsDeniedActionsResponse is the response struct for api DescribeDrillPairsDeniedActions
type DescribeDrillPairsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		DrillPairDeniedActionSet []*DrillPairDeniedAction `json:"DrillPairDeniedActionSet,omitempty" name:"DrillPairDeniedActionSet"`
		RequestId                *string                  `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeProtectGroupsDeniedActionsRequest is the request struct for api DescribeProtectGroupsDeniedActions
type DescribeProtectGroupsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	ProtectGroupIds []*string `json:"ProtectGroupIds,omitempty" name:"ProtectGroupIds"`
}

// DescribeProtectGroupsDeniedActionsResponse is the response struct for api DescribeProtectGroupsDeniedActions
type DescribeProtectGroupsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeSecurityGroupMappingsRequest is the request struct for api DescribeSecurityGroupMappings
type DescribeSecurityGroupMappingsRequest struct {
	*tchttp.BaseRequest
	SitePairId *string   `json:"SitePairId,omitempty" name:"SitePairId"`
	Filters    []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset     *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit      *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order      *string   `json:"Order,omitempty" name:"Order"`
	OrderField *string   `json:"OrderField,omitempty" name:"OrderField"`
}

// DescribeSecurityGroupMappingsResponse is the response struct for api DescribeSecurityGroupMappings
type DescribeSecurityGroupMappingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeVipStatusRequest is the request struct for api DescribeVipStatus
type DescribeVipStatusRequest struct {
	*tchttp.BaseRequest
	PeerCloudName *string `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	ApiModule     *string `json:"ApiModule,omitempty" name:"ApiModule"`
	SourceRegion  *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	Zone          *string `json:"Zone,omitempty" name:"Zone"`
	SubnetId      *string `json:"SubnetId,omitempty" name:"SubnetId"`
	Vip           *string `json:"Vip,omitempty" name:"Vip"`
}

// DescribeVipStatusResponse is the response struct for api DescribeVipStatus
type DescribeVipStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeVpcMappingsRequest is the request struct for api DescribeVpcMappings
type DescribeVpcMappingsRequest struct {
	*tchttp.BaseRequest
	SitePairId *string   `json:"SitePairId,omitempty" name:"SitePairId"`
	Filters    []*Filter `json:"Filters,omitempty" name:"Filters"`
	Offset     *uint64   `json:"Offset,omitempty" name:"Offset"`
	Limit      *uint64   `json:"Limit,omitempty" name:"Limit"`
	Order      *string   `json:"Order,omitempty" name:"Order"`
	OrderField *string   `json:"OrderField,omitempty" name:"OrderField"`
}

// DescribeVpcMappingsResponse is the response struct for api DescribeVpcMappings
type DescribeVpcMappingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// FailoverCopyPairsRequest is the request struct for api FailoverCopyPairs
type FailoverCopyPairsRequest struct {
	*tchttp.BaseRequest
	CopyPairIds  []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
	FailoverType *string   `json:"FailoverType,omitempty" name:"FailoverType"`
}

// FailoverCopyPairsResponse is the response struct for api FailoverCopyPairs
type FailoverCopyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		TaskId    []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		RequestId *string   `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// ModifyCopyPairAttributeRequest is the request struct for api ModifyCopyPairAttribute
type ModifyCopyPairAttributeRequest struct {
	*tchttp.BaseRequest
	CopyPairId   *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
	CopyPairName *string `json:"CopyPairName,omitempty" name:"CopyPairName"`
}

// ModifyCopyPairAttributeResponse is the response struct for api ModifyCopyPairAttribute
type ModifyCopyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// ModifyDrillPairAttributeRequest is the request struct for api ModifyDrillPairAttribute
type ModifyDrillPairAttributeRequest struct {
	*tchttp.BaseRequest
	DrillPairId   *string `json:"DrillPairId,omitempty" name:"DrillPairId"`
	DrillPairName *string `json:"DrillPairName,omitempty" name:"DrillPairName"`
}

// ModifyDrillPairAttributeResponse is the response struct for api ModifyDrillPairAttribute
type ModifyDrillPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// ModifyProtectGroupAttributeRequest is the request struct for api ModifyProtectGroupAttribute
type ModifyProtectGroupAttributeRequest struct {
	*tchttp.BaseRequest
	ProtectGroupId   *string `json:"ProtectGroupId,omitempty" name:"ProtectGroupId"`
	ProtectGroupName *string `json:"ProtectGroupName,omitempty" name:"ProtectGroupName"`
}

// ModifyProtectGroupAttributeResponse is the response struct for api ModifyProtectGroupAttribute
type ModifyProtectGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// ModifySitePairAttributeRequest is the request struct for api ModifySitePairAttribute
type ModifySitePairAttributeRequest struct {
	*tchttp.BaseRequest
	DisasterRecoverySitePairId   *string `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	DisasterRecoverySitePairName *string `json:"DisasterRecoverySitePairName,omitempty" name:"DisasterRecoverySitePairName"`
	IsCrossCloud                 *bool   `json:"IsCrossCloud,omitempty" name:"IsCrossCloud"`
}

// ModifySitePairAttributeResponse is the response struct for api ModifySitePairAttribute
type ModifySitePairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// RunCopyPairTasksRequest is the request struct for api RunCopyPairTasks
type RunCopyPairTasksRequest struct {
	*tchttp.BaseRequest
	CopyPairIds  []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
}

// RunCopyPairTasksResponse is the response struct for api RunCopyPairTasks
type RunCopyPairTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RunCopyTaskResultSet *CopyTaskResult `json:"RunCopyTaskResultSet,omitempty" name:"RunCopyTaskResultSet"`
		RequestId            *string         `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// StopCopyPairTasksRequest is the request struct for api StopCopyPairTasks
type StopCopyPairTasksRequest struct {
	*tchttp.BaseRequest
	CopyPairIds  []*string `json:"CopyPairIds,omitempty" name:"CopyPairIds"`
	CopyPairType *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
}

// StopCopyPairTasksResponse is the response struct for api StopCopyPairTasks
type StopCopyPairTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		StopCopyTaskResultSet []*CopyTaskResult `json:"StopCopyTaskResultSet,omitempty" name:"StopCopyTaskResultSet"`
		RequestId             *string           `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeDisasterRecoveryOverviewRequest is the request struct for api DescribeDisasterRecoveryOverview
type DescribeDisasterRecoveryOverviewRequest struct {
	*tchttp.BaseRequest
}

// DescribeDisasterRecoveryOverviewResponse is the response struct for api DescribeDisasterRecoveryOverview
type DescribeDisasterRecoveryOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		DrcOverview         *DrcOverview        `json:"DrcOverview,omitempty" name:"DrcOverview"`
		OverviewInRegionSet []*OverviewInRegion `json:"OverviewInRegionSet,omitempty" name:"OverviewInRegionSet"`
		RequestId           *string             `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateTargetCfsParameters 批量创建cfs的参数
type CreateTargetCfsParameters struct {
	SourceFilesystemId       *string `json:"SourceFilesystemId,omitempty" name:"SourceFilesystemId"`
	Protocol                 *string `json:"Protocol,omitempty" name:"Protocol"`
	StorageType              *string `json:"StorageType,omitempty" name:"StorageType"`
	NetInterface             *string `json:"NetInterface,omitempty" name:"NetInterface"`
	PGroupId                 *string `json:"PGroupId,omitempty" name:"PGroupId"`
	VpcId                    *string `json:"VpcId,omitempty" name:"VpcId"`
	UnVpcId                  *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	SubnetId                 *string `json:"SubnetId,omitempty" name:"SubnetId"`
	UnSubnetId               *string `json:"UnSubnetId,omitempty" name:"UnSubnetId"`
	MountIP                  *string `json:"MountIP,omitempty" name:"MountIP"`
	StorageResourcePkgId     *string `json:"StorageResourcePkgId,omitempty" name:"StorageResourcePkgId"`
	BandwidthResourcePkgId   *string `json:"BandwidthResourcePkgId,omitempty" name:"BandwidthResourcePkgId"`
	FsName                   *string `json:"FsName,omitempty" name:"FsName"`
	Zone                     *string `json:"Zone,omitempty" name:"Zone"`
	CopyPairId               *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
	RecoveryTime             *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
}

type CopyPair struct {
	AppId                      *uint64   `json:"AppId,omitempty" name:"AppId"`
	CopyPairId                 *string   `json:"CopyPairId,omitempty" name:"CopyPairId"`
	CopyPairName               *string   `json:"CopyPairName,omitempty" name:"CopyPairName"`
	DisasterRecoverySitePairId *string   `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	CopyPairState              *string   `json:"CopyPairState,omitempty" name:"CopyPairState"`
	SourceRegion               *string   `json:"SourceRegion,omitempty" name:"SourceRegion"`
	SourceZone                 *string   `json:"SourceZone,omitempty" name:"SourceZone"`
	TargetRegion               *string   `json:"TargetRegion,omitempty" name:"TargetRegion"`
	TargetZone                 *string   `json:"TargetZone,omitempty" name:"TargetZone"`
	SourceDiskId               *string   `json:"SourceDiskId,omitempty" name:"SourceDiskId"`
	TargetDiskId               *string   `json:"TargetDiskId,omitempty" name:"TargetDiskId"`
	InstanceId                 *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Percent                    *uint64   `json:"Percent,omitempty" name:"Percent"`
	CopyPairType               *string   `json:"CopyPairType,omitempty" name:"CopyPairType"`
	CreateTime                 *string   `json:"CreateTime,omitempty" name:"CreateTime"`
	Rpo                        *uint64   `json:"Rpo,omitempty" name:"Rpo"`
	AccountUin                 *string   `json:"AccountUin,omitempty" name:"AccountUin"`
	SubAccountUin              *string   `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	AccountName                *string   `json:"AccountName,omitempty" name:"AccountName"`
	LatestProtectionTime       *string   `json:"LatestProtectionTime,omitempty" name:"LatestProtectionTime"`
	ProtectionTimeSet          []*string `json:"ProtectionTimeSet,omitempty" name:"ProtectionTimeSet"`
	ProtectGroupId             *string   `json:"ProtectGroupId,omitempty" name:"ProtectGroupId"`
	SourceVpc                  *string   `json:"SourceVpc,omitempty" name:"SourceVpc"`
	TargetVpc                  *string   `json:"TargetVpc,omitempty" name:"TargetVpc"`
	DataDirection              *string   `json:"DataDirection,omitempty" name:"DataDirection"`
	CreateFrom                 *string   `json:"CreateFrom,omitempty" name:"CreateFrom"`
	PeerCloudName              *string   `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	InstanceCopyPairId         *string   `json:"InstanceCopyPairId,omitempty" name:"InstanceCopyPairId"`
	SourceResourceId           *string   `json:"SourceResourceId,omitempty" name:"SourceResourceId"`
	TargetResourceId           *string   `json:"TargetResourceId,omitempty" name:"TargetResourceId"`
}

type CopyPairDeniedAction struct {
	CopyPairId    *string         `json:"CopyPairId,omitempty" name:"CopyPairId"`
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type CopyTaskResult struct {
	CopyPairId *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
	Code       *string `json:"Code,omitempty" name:"Code"`
	Message    *string `json:"Message,omitempty" name:"Message"`
}

type DeniedAction struct {
	Action  *string `json:"Action,omitempty" name:"Action"`
	Message *string `json:"Message,omitempty" name:"Message"`
	Code    *string `json:"Code,omitempty" name:"Code"`
}

type DisasterRecoverySitePair struct {
	AppId                         *uint64              `json:"AppId,omitempty" name:"AppId"`
	DisasterRecoverySitePairId    *string              `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	DisasterRecoverySitePairName  *string              `json:"DisasterRecoverySitePairName,omitempty" name:"DisasterRecoverySitePairName"`
	DisasterRecoverySitePairType  *string              `json:"DisasterRecoverySitePairType,omitempty" name:"DisasterRecoverySitePairType"`
	DisasterRecoverySitePairState *string              `json:"DisasterRecoverySitePairState,omitempty" name:"DisasterRecoverySitePairState"`
	SourceRegion                  *string              `json:"SourceRegion,omitempty" name:"SourceRegion"`
	SourceZone                    *string              `json:"SourceZone,omitempty" name:"SourceZone"`
	TargetRegion                  *string              `json:"TargetRegion,omitempty" name:"TargetRegion"`
	TargetZone                    *string              `json:"TargetZone,omitempty" name:"TargetZone"`
	CreateTime                    *string              `json:"CreateTime,omitempty" name:"CreateTime"`
	AccountUin                    *string              `json:"AccountUin,omitempty" name:"AccountUin"`
	SubAccountUin                 *string              `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	ProtectedResourceSet          []*ProtectedResource `json:"ProtectedResourceSet,omitempty" name:"ProtectedResourceSet"`
	AccountName                   *string              `json:"AccountName,omitempty" name:"AccountName"`
	BindProtectGroupCount         *uint64              `json:"BindProtectGroupCount,omitempty" name:"BindProtectGroupCount"`
	SourceVpc                     *string              `json:"SourceVpc,omitempty" name:"SourceVpc"`
	TargetVpc                     *string              `json:"TargetVpc,omitempty" name:"TargetVpc"`
	CopyType                      *string              `json:"CopyType,omitempty" name:"CopyType"`
	CreateFrom                    *string              `json:"CreateFrom,omitempty" name:"CreateFrom"`
	PeerCloudName                 *string              `json:"PeerCloudName,omitempty" name:"PeerCloudName"`
	LocalCloudName                *string              `json:"LocalCloudName,omitempty" name:"LocalCloudName"`
}

type DisasterRecoverySitePairDeniedAction struct {
	DisasterRecoverySitePairId *string         `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	DeniedActions              []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type DrcOverview struct {
	SitePairCount            *uint64 `json:"SitePairCount,omitempty" name:"SitePairCount"`
	FilesystemCopyPairCount  *uint64 `json:"FilesystemCopyPairCount,omitempty" name:"FilesystemCopyPairCount"`
	FilesystemDrillPairCount *uint64 `json:"FilesystemDrillPairCount,omitempty" name:"FilesystemDrillPairCount"`
}

type DrillPair struct {
	AppId                      *uint64 `json:"AppId,omitempty" name:"AppId"`
	DrillPairId                *string `json:"DrillPairId,omitempty" name:"DrillPairId"`
	DrillPairName              *string `json:"DrillPairName,omitempty" name:"DrillPairName"`
	DrillPairState             *string `json:"DrillPairState,omitempty" name:"DrillPairState"`
	CopyPairId                 *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
	SourceRegion               *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	SourceZone                 *string `json:"SourceZone,omitempty" name:"SourceZone"`
	TargetRegion               *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	TargetZone                 *string `json:"TargetZone,omitempty" name:"TargetZone"`
	SourceResourceId           *string `json:"SourceResourceId,omitempty" name:"SourceResourceId"`
	TargetResourceId           *string `json:"TargetResourceId,omitempty" name:"TargetResourceId"`
	DrillPairType              *string `json:"DrillPairType,omitempty" name:"DrillPairType"`
	Size                       *uint64 `json:"Size,omitempty" name:"Size"`
	RecoveryTime               *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
	CreateTime                 *string `json:"CreateTime,omitempty" name:"CreateTime"`
	EndTimd                    *string `json:"EndTimd,omitempty" name:"EndTimd"`
	Rollbacking                *uint64 `json:"Rollbacking,omitempty" name:"Rollbacking"`
	RollbackPercent            *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`
	AccountUin                 *string `json:"AccountUin,omitempty" name:"AccountUin"`
	SubAccountUin              *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	AccountName                *string `json:"AccountName,omitempty" name:"AccountName"`
	DisasterRecoverySitePairId *string `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	ProtectGroupId             *string `json:"ProtectGroupId,omitempty" name:"ProtectGroupId"`
	DrillGroupId               *string `json:"DrillGroupId,omitempty" name:"DrillGroupId"`
}

type DrillPairDeniedAction struct {
	DrillPairId   *string         `json:"DrillPairId,omitempty" name:"DrillPairId"`
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type Filter struct {
	Name   *string   `json:"Name,omitempty" name:"Name"`
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type OperationLog struct {
	TaskId                     *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	TaskName                   *string `json:"TaskName,omitempty" name:"TaskName"`
	TaskDescription            *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	AppId                      *uint64 `json:"AppId,omitempty" name:"AppId"`
	DisasterRecoverySitePairId *string `json:"DisasterRecoverySitePairId,omitempty" name:"DisasterRecoverySitePairId"`
	CopyPairId                 *string `json:"CopyPairId,omitempty" name:"CopyPairId"`
	DrillPairId                *string `json:"DrillPairId,omitempty" name:"DrillPairId"`
	SourceResourceId           *string `json:"SourceResourceId,omitempty" name:"SourceResourceId"`
	TargetResourceId           *string `json:"TargetResourceId,omitempty" name:"TargetResourceId"`
	ResourceType               *string `json:"ResourceType,omitempty" name:"ResourceType"`
	LogTaskId                  *string `json:"LogTaskId,omitempty" name:"LogTaskId"`
	StartTime                  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime                    *string `json:"EndTime,omitempty" name:"EndTime"`
	TaskState                  *string `json:"TaskState,omitempty" name:"TaskState"`
	AccountName                *string `json:"AccountName,omitempty" name:"AccountName"`
	SubAccountUin              *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type OverviewInRegion struct {
	Region                   *string `json:"Region,omitempty" name:"Region"`
	SitePairCount            *uint64 `json:"SitePairCount,omitempty" name:"SitePairCount"`
	FilesystemCopyPairCount  *uint64 `json:"FilesystemCopyPairCount,omitempty" name:"FilesystemCopyPairCount"`
	FilesystemDrillPairCount *uint64 `json:"FilesystemDrillPairCount,omitempty" name:"FilesystemDrillPairCount"`
}

type ProtectedResource struct {
	ResourceType  *string   `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`
}

type ResourceTag struct {
	TagKey   *string `json:"TagKey,omitempty" name:"TagKey"`
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type SubTaskStatistic struct {
	SucceedCount *uint64 `json:"SucceedCount,omitempty" name:"SucceedCount"`
	FailedCount  *uint64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

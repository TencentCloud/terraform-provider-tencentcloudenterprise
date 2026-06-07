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

package v20220508

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 英文：en，中文：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员列表

		Items []*NodeMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMemberExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的成员账号uin列表

		ExistMemberUin []*uint64 `json:"ExistMemberUin,omitempty" name:"ExistMemberUin"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMemberExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMemberExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 父节点id

		ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略名称， tce后端自动生成

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 身份id， 固定传1

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员账号uin

		MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 成员类型

		MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
		// 集团策略类型

		OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
		// 集团策略名称

		OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
		// 集团权限

		OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
		// 代付uin

		PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
		// 代付用户名称

		PayName *string `json:"PayName,omitempty" name:"PayName"`
		// 节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否允许退出

		IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
		// 授权身份列表

		OrgIdentity []*Identity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数据量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 授权子账号列表

		Items []*OrgMemberAuthAccount `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数据量大小

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 节点列表

		Items []*OrgNode `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 新的父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 被移动的节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *MoveOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Permission struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest

	// 节点id列表

	NodeId []*int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesByParentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织节点列表

		Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodesByParentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesByParentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Identity struct {

	// id

	IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 别名

	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
}

type CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 成员账号列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgNode struct {

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 标签

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin， 传0则不使用成员账号进行过滤

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id，传0则不指定策略

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小，[0,50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 子账号uin， 传0则不根据子账号进行过滤

	TargetSubAccountUin *int64 `json:"TargetSubAccountUin,omitempty" name:"TargetSubAccountUin"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodeMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小[0,50]

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字，支持账号名称和uin

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 语言版本。zh：中文版本；en：英文版本。

	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// 主体名称（实体名称）

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`

	// 可信服务简称，查询可信服务管理员时使用。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 是否允许成员退出。允许：Allow，不允许：Denied。

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 成员类型。邀请：Invite， 创建：Create

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 成员uin列表。最大10个

	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeMember struct {

	// 成员账号uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 成员类型，Create,Invite,Admin等

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 组织策略类型，Finance

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 权限列表

	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 代付账户

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 代付用户名

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 成员访问权限

	OrgIdentities []*MemberIdentity `json:"OrgIdentities,omitempty" name:"OrgIdentities"`
	// 成员标签列表

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略id

		PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest

	// 新的节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 成员账号uin列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 授权子账号uin列表。 <=5

	OrgSubAccountUins []*uint64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// 企业组织ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostUin *int64 `json:"HostUin,omitnil,omitempty" name:"HostUin"`

	// 创建者昵称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 企业组织类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgType *int64 `json:"OrgType,omitnil,omitempty" name:"OrgType"`

	// 是否组织管理员。是：true ，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsManager *bool `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// 策略类型。财务管理：Financial
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// 策略名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// 成员财务权限列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPermission []*Permission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// 组织根节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootNodeId *int64 `json:"RootNodeId,omitnil,omitempty" name:"RootNodeId"`

	// 组织创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 成员加入时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// 成员是否允许退出。允许：Allow，不允许：Denied
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 代付者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// 是否可信服务管理员。是：true，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// 是否实名主体管理员。是：true，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *DescribeOrganizationResponseParams `json:"Response"`

	//Response *struct {
	//	// 集团组织id
	//
	//	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
	//	// 组织类型
	//
	//	OrgType *int64 `json:"OrgType,omitempty" name:"OrgType"`
	//	// 管理员账号uin
	//
	//	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	//	// 管理员账号昵称
	//
	//	NickName *string `json:"NickName,omitempty" name:"NickName"`
	//	// 请求账号是否为管理员
	//
	//	IsManager *bool `json:"IsManager,omitempty" name:"IsManager"`
	//	// 请求账号是否允许退出组织
	//
	//	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	//	// 根几点id
	//
	//	RootNodeId *int64 `json:"RootNodeId,omitempty" name:"RootNodeId"`
	//	// 组织关系策略类型，如Financial
	//
	//	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	//	// 关系策略名称，如财务管理
	//
	//	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	//	// 关联策略权限
	//
	//	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	//	// 代付账号uin
	//
	//	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	//	// 代付账号用户名
	//
	//	PayName *string `json:"PayName,omitempty" name:"PayName"`
	//	// 组织创建时间
	//
	//	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	//	// 请求账号加入集团组织时间
	//
	//	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
	//	// 管理员用户名
	//
	//	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`
	//	// 集团组织名称
	//
	//	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	//	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	//	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	//} `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的节点id

		NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgMember struct {

	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 成员类型

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 集团策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 集团策略名称

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 集团权限

	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 代付uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 代付用户名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 节点id

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否允许退出

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 权限状态

	PermissionStatus *string `json:"PermissionStatus,omitempty" name:"PermissionStatus"`
	// 授权身份列表

	OrgIdentity []*Identity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否是管理员

	IsManager *int64 `json:"IsManager,omitempty" name:"IsManager"`
	// 实名名称

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`
	// 成员主体状态 1:正常 2:管理员实名变更，待变更成员实名 3:管理员实名变更，待重新账号互 4:成员实名变更，待重新账号互信 5：成员实名变更，待关联主体

	AuthStatus *uint64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
	// 安全信息绑定状态 未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed

	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`
	// 成员标签列表

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type OrgMemberAuthAccount struct {

	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
	// 策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 身份id

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份角色名称

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 身份角色别名

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 被授权的成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 子账号名称

	OrgSubAccountName *string `json:"OrgSubAccountName,omitempty" name:"OrgSubAccountName"`
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据量大小

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员账号列表

		Items []*OrgMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest

	// 语言，支持cn,en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 可信服务产品简称。查询是否该可信服务管理员时必须指定

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量[0-50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	OrgSubAccountUin *uint64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMemberExistRequest struct {
	*tchttp.BaseRequest

	// 待检测的账号uin列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckMemberExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMemberExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeRequest struct {
	*tchttp.BaseRequest

	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DescribeOrganizationNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodesByParentRequest struct {
	*tchttp.BaseRequest

	// 父节点id

	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
}

func (r *DescribeOrganizationNodesByParentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodesByParentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationRequestParams struct {
}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationResponseParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationResponseParams `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationRequestParams struct {
}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// 被删除成员的Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 被删除成员的Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdentityPolicy struct {
	// CAM预设策略ID。PolicyType 为预设策略时有效且必选
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// CAM预设策略名称。PolicyType 为预设策略时有效且必选
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型。取值 1-自定义策略  2-预设策略；默认值2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *uint64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 自定义策略内容，遵循CAM策略语法。PolicyType 为自定义策略时有效且必选
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`
}

// Predefined struct for user
type CreateOrganizationIdentityRequestParams struct {
	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest

	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationIdentityResponseParams struct {
	// 身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgIdentity struct {
	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份类型。 1-预设、 2-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
	// 配置状态。取值：1-配置完成 2-需重新配置

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ListOrganizationIdentityRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID。可以通过身份ID搜索
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

type ListOrganizationIdentityRequest struct {
	*tchttp.BaseRequest

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID。可以通过身份ID搜索
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

func (r *ListOrganizationIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityResponseParams struct {
	// 总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 条目详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationIdentityResponseParams `json:"Response"`
}

func (r *ListOrganizationIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationIdentityRequestParams struct {
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

type UpdateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

func (r *UpdateOrganizationIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *UpdateOrganizationIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityRequestParams struct {
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationIdentityRequest struct {
	*tchttp.BaseRequest

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyRequestParams struct {
	// 策略名。
	// 长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略名。
	// 长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyResponseParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyResponseParams `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyResponseParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型。1-自定义 2-预设策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type ListPoliciesRequestParams struct {
	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询范围。取值范围： All-获取所有策略、QCS-只获取预设策略、Local-只获取自定义策略，默认值：All
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 搜索关键字。按照策略名搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询范围。取值范围： All-获取所有策略、QCS-只获取预设策略、Local-只获取自定义策略，默认值：All
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 搜索关键字。按照策略名搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesResponseParams struct {
	// 策略总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 策略列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ListPolicyNode `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesResponseParams `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPolicyNode struct {
	// 策略创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略绑定次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedTimes *uint64 `json:"AttachedTimes,omitnil,omitempty" name:"AttachedTimes"`

	// 策略描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略Id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略类型 1-自定义 2-预设
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type DescribePolicyRequestParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *DescribePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyResponseParams `json:"Response"`
}

func (r *DescribePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyRequestParams struct {
	// 需要编辑的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略名。长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 需要编辑的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略名。长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePolicyResponseParams `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequestParams struct {
	// 需要删除的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// 需要删除的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyResponseParams `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnablePolicyTypeRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type EnablePolicyTypeRequest struct {
	*tchttp.BaseRequest

	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *EnablePolicyTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnablePolicyTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnablePolicyTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnablePolicyTypeResponse struct {
	*tchttp.BaseResponse
	Response *EnablePolicyTypeResponseParams `json:"Response"`
}

func (r *EnablePolicyTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnablePolicyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyConfigRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值0，取值范围：0-服务控制策略、1-标签策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribePolicyConfigRequest struct {
	*tchttp.BaseRequest

	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值0，取值范围：0-服务控制策略、1-标签策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribePolicyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyConfigResponseParams struct {
	// 开启状态。0-未开启、1-开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 策略类型。SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyConfigResponseParams `json:"Response"`
}

func (r *DescribePolicyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisablePolicyTypeRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type DisablePolicyTypeRequest struct {
	*tchttp.BaseRequest

	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *DisablePolicyTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisablePolicyTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisablePolicyTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisablePolicyTypeResponse struct {
	*tchttp.BaseResponse
	Response *DisablePolicyTypeResponseParams `json:"Response"`
}

func (r *DisablePolicyTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisablePolicyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachPolicyRequestParams struct {
	// 绑定策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AttachPolicyRequest struct {
	*tchttp.BaseRequest

	// 绑定策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *AttachPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachPolicyResponseParams `json:"Response"`
}

func (r *AttachPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTargetsForPolicyNode struct {
	// scp账号uin或节点Id
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 关联类型 1-节点关联 2-用户关联
	RelatedType *uint64 `json:"RelatedType,omitnil,omitempty" name:"RelatedType"`

	// 账号或者节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 绑定时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`
}

// Predefined struct for user
type ListTargetsForPolicyRequestParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。取值范围：All-全部、User-用户、Node-节点
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 按照多个策略id搜索，空格隔开。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListTargetsForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。取值范围：All-全部、User-用户、Node-节点
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 按照多个策略id搜索，空格隔开。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListTargetsForPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsForPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsForPolicyResponseParams struct {
	// 总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 指定SCP策略关联目标列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ListTargetsForPolicyNode `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTargetsForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListTargetsForPolicyResponseParams `json:"Response"`
}

func (r *ListTargetsForPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPolicyRequestParams struct {
	// 解绑策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DetachPolicyRequest struct {
	*tchttp.BaseRequest

	// 解绑策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DetachPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachPolicyResponseParams `json:"Response"`
}

func (r *DetachPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type CreateOrganizationMemberRequestParams struct {
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberResponseParams struct {
	// 成员Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberAuthIdentityRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个，可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

type CreateOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest

	// 成员Uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个，可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

func (r *CreateOrganizationMemberAuthIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberAuthIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgMemberAuthIdentity struct {
	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份的角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份的角色别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 身份描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 首次配置成功的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次配置成功的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 身份类型。取值： 1-预设身份  2-自定义身份
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 配置状态。取值：1-配置完成 2-需重新配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 成员Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个, 可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DescribeOrganizationMemberAuthIdentitiesRequest struct {
	*tchttp.BaseRequest

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个, 可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesResponseParams struct {
	// 授权身份列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMemberAuthIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthIdentitiesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityRequestParams struct {
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationMemberAuthIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMemberAuthIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationMemberEmailRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type AddOrganizationMemberEmailRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *AddOrganizationMemberEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationMemberEmailResponseParams struct {
	// 绑定Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationMemberEmailResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationMemberEmailResponseParams `json:"Response"`
}

func (r *AddOrganizationMemberEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DescribeOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberEmailBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindResponseParams struct {
	// 绑定ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 申请时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 邮箱地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 安全手机号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 绑定状态。    未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 绑定时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindTime *string `json:"BindTime,omitnil,omitempty" name:"BindTime"`

	// 失败说明。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 安全手机绑定状态 。 未绑定：0，已绑定：1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneBind *uint64 `json:"PhoneBind,omitnil,omitempty" name:"PhoneBind"`

	// 国际区号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberEmailBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type UpdateOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *UpdateOrganizationMemberEmailBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberEmailBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMembersPolicyRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest

	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationMembersPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMembersPolicyResponseParams struct {
	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMembersPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyRequestParams struct {
	// 访问策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DeleteOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest

	// 访问策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DeleteOrganizationMembersPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationRequestParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

type QuitOrganizationRequest struct {
	*tchttp.BaseRequest

	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

func (r *QuitOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QuitOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *QuitOrganizationResponseParams `json:"Response"`
}

func (r *QuitOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 是否允许成员退出组织。取值：Allow-允许、Denied-不允许

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 代付者Uin。成员财务权限有代付费时需要，取值为成员对应主体的主体管理员Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 成员Uin。

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注。最大长度为40个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 关系策略类型。PolicyType不为空，PermissionIds不能为空。取值：Financial

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 成员财务权限ID列表。PermissionIds不为空，PolicyType不能为空。
	// 取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析，如果有值，1、2 默认必须

	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountStatus struct {

	// 主账号Uin。

	AccountUin *int64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 是否经销类型账号。true-是、false-否

	DistributionAccount *bool `json:"DistributionAccount,omitempty" name:"DistributionAccount"`
	// 是否和当前账号不同组织。true-是、false-否

	OtherOrganization *bool `json:"OtherOrganization,omitempty" name:"OtherOrganization"`
	// 是否和当前账号同实名主体。true-是、false-否

	SameAuth *bool `json:"SameAuth,omitempty" name:"SameAuth"`
	// 是否和当前账号同组织。true-是、false-否

	SameOrganization *bool `json:"SameOrganization,omitempty" name:"SameOrganization"`
}

type AddOrgMember struct {

	// 成员名

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 所属节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 关系策略权限Id列表

	PermissionIds []*int64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
	// 关系策略类型

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type AuthGroup struct {

	// 用户组Id。

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称。

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type AuthNode struct {

	// 互信主体名称

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`
	// 主体管理员

	Manager *MemberMainInfo `json:"Manager,omitempty" name:"Manager"`
	// 互信主体关系ID

	RelationId *int64 `json:"RelationId,omitempty" name:"RelationId"`
}

type AuthRelation struct {

	// 主体名称

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`
	// 主体互信状态 1:正常 2:待重新认证 3:待变更主体下成员实名

	AuthStatus *uint64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
	// 主体类型 manager：管理、member：成员

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 添加日期

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 主体管理账号

	Manager *MemberMainInfo `json:"Manager,omitempty" name:"Manager"`
	// 主体管理员状态 1:正常 2:实名不一致,待替换

	ManagerStatus *uint64 `json:"ManagerStatus,omitempty" name:"ManagerStatus"`
	// 主体下成员数

	MemberCount *string `json:"MemberCount,omitempty" name:"MemberCount"`
	// 主体下成员Uin

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 待审批创建成员数

	PendingCount *string `json:"PendingCount,omitempty" name:"PendingCount"`
	// 主体关系ID

	RelationId *int64 `json:"RelationId,omitempty" name:"RelationId"`
}

type AuthRelationApply struct {

	// 审批人。

	Approver []*string `json:"Approver,omitempty" name:"Approver"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 管理员实名名称。

	HostAuthName *string `json:"HostAuthName,omitempty" name:"HostAuthName"`
	// 管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 申请ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 成员实名名称。

	MemberAuthName *string `json:"MemberAuthName,omitempty" name:"MemberAuthName"`
	// 成员Uin。

	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 申请类型。控制台：Console， 野鹤：YeHe。

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 状态。可用：Valid。拒绝：Denied。接受：Approve。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AuthRelationFile struct {

	// 文件名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件路径。

	Url *string `json:"Url,omitempty" name:"Url"`
}

type CheckImportMemberResult struct {

	// 错误码。

	Code *string `json:"Code,omitempty" name:"Code"`
	// 成员名。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 错误信息。

	Message *string `json:"Message,omitempty" name:"Message"`
	// 节点ID。

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 策略权限ID。

	PermissionIds []*int64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
	// 策略类型。

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

type CreateMemberApply struct {

	// 审批人。

	Approver []*string `json:"Approver,omitempty" name:"Approver"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 申请ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 成员实名名称。

	MemberAuthName *string `json:"MemberAuthName,omitempty" name:"MemberAuthName"`
	// 成员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 节点ID。

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 策略权限。

	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 策略名称。

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 策略类型。

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 状态。可用：Valid。拒绝：Denied。接受：Approve。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type EffectivePolicy struct {

	// 有效策略更新时间。

	LastUpdatedTimestamp *uint64 `json:"LastUpdatedTimestamp,omitempty" name:"LastUpdatedTimestamp"`
	// 有效策略内容。

	PolicyContent *string `json:"PolicyContent,omitempty" name:"PolicyContent"`
	// 目标ID。

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
}

type EventInfo struct {

	// 事件ID

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 关联产品

	RelationProduct *string `json:"RelationProduct,omitempty" name:"RelationProduct"`
}

type EventProductInfo struct {

	// 产品标识

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type EventRelationParam struct {

	// 参数名

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 不需要审批的参数值列表

	ParamValues []*string `json:"ParamValues,omitempty" name:"ParamValues"`
}

type FinancialMemberNum struct {

	// 主体名称

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`
	// 主体类型 manager：管理、member：成员

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 代付费成员数

	BehalfPay *int64 `json:"BehalfPay,omitempty" name:"BehalfPay"`
	// 代付费账号

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 自付费优惠继承成员数

	SelfPayInherit *int64 `json:"SelfPayInherit,omitempty" name:"SelfPayInherit"`
	// 自付费非优惠继承成员数

	SelfPayNoInherit *int64 `json:"SelfPayNoInherit,omitempty" name:"SelfPayNoInherit"`
}

type FinancialProductInfo struct {

	// 产品简称。

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称。

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type ImportMemberApply struct {

	// 审批人。

	Approver []*string `json:"Approver,omitempty" name:"Approver"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 导入文件名。

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 导入文件路径。

	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
	// 申请ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 状态。可用：Valid。拒绝：Denied。接受：Approve。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ListPoliciesForTarget struct {

	// 策略创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略绑定时间

	AttachTime *string `json:"AttachTime,omitempty" name:"AttachTime"`
	// 部门名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略Id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 关联类型 1-节点 2-用户

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 关联的账号或节点

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 策略更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ManagerShareMember struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 共享单元资源数

	SharedResourceNum *int64 `json:"SharedResourceNum,omitempty" name:"SharedResourceNum"`
	// 共享成员Uin

	ShareMemberUin *int64 `json:"ShareMemberUin,omitempty" name:"ShareMemberUin"`
	// 共享成员ID

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
	// 共享单元名

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type ManagerShareResource struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 产品资源ID

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 共享资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 共享单元成员数

	SharedMemberNum *int64 `json:"SharedMemberNum,omitempty" name:"SharedMemberNum"`
	// 使用中共享单元成员数

	SharedMemberUseNum *int64 `json:"SharedMemberUseNum,omitempty" name:"SharedMemberUseNum"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 共享单元ID

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
	// 共享单元名

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type ManagerShareResourceByType struct {

	// 资源地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 业务产品资源ID

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 共享资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ManagerShareUnit struct {

	// 共享单元地域。

	Area *string `json:"Area,omitempty" name:"Area"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 共享单元名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 共享单元管理员OwnerUin。

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 共享单元成员数。

	ShareMemberNum *int64 `json:"ShareMemberNum,omitempty" name:"ShareMemberNum"`
	// 共享单元资源数。

	ShareResourceNum *int64 `json:"ShareResourceNum,omitempty" name:"ShareResourceNum"`
	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号

	ShareScope *uint64 `json:"ShareScope,omitempty" name:"ShareScope"`
	// 共享单元管理员Uin。

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 共享单元ID。

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
}

type MemberBaseInfo struct {

	// 组织管理员Uin

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员Uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 组织Id

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 成员描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type MemberCloudApplication struct {

	// 应用描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 应用链接。

	Link *string `json:"Link,omitempty" name:"Link"`
	// 应用名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 二级应用列表。

	SubApplications []*SubCloudApplication `json:"SubApplications,omitempty" name:"SubApplications"`
}

type MemberEvent struct {

	// API接口

	ActionName []*string `json:"ActionName,omitempty" name:"ActionName"`
	// 事件英文名称

	EventEnName *string `json:"EventEnName,omitempty" name:"EventEnName"`
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 产品标识

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 关联参数

	RelationParams []*EventRelationParam `json:"RelationParams,omitempty" name:"RelationParams"`
	// 关联接口

	RelationProduct *string `json:"RelationProduct,omitempty" name:"RelationProduct"`
}

type MemberIdentity struct {

	// 身份名称。

	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
	// 身份ID。

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
}

type MemberMainInfo struct {

	// 成员名称

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type MemberNodes struct {

	// 成员uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 节点列表。

	Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes"`
}

type MemberOperateAccount struct {

	// 成员账号列表

	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 部门ID列表

	NodeId []*int64 `json:"NodeId,omitempty" name:"NodeId"`
}

type MemberOperateAccountDetail struct {

	// 关联成员

	Members []*MemberMainInfo `json:"Members,omitempty" name:"Members"`
	// 关联部门

	Nodes []*NodeMainInfo `json:"Nodes,omitempty" name:"Nodes"`
}

type MemberOperateProcess struct {

	// 审批流程名称

	BpaasName *string `json:"BpaasName,omitempty" name:"BpaasName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 操作审批ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 关联账号

	OperateAccount *MemberOperateAccountDetail `json:"OperateAccount,omitempty" name:"OperateAccount"`
	// 操作事件

	OperateEvent *EventInfo `json:"OperateEvent,omitempty" name:"OperateEvent"`
	// 产品key

	Product *string `json:"Product,omitempty" name:"Product"`
	// 状态 on-开启、off-关闭

	Status *string `json:"Status,omitempty" name:"Status"`
	// 审批流程UniqueId

	UniqueId *int64 `json:"UniqueId,omitempty" name:"UniqueId"`
}

type MemberPermissionChangeRecord struct {

	// 提交时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 管理员Uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 记录Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 成员Uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 成员名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更后权限

	NewOrgPermission []*OrgPermission `json:"NewOrgPermission,omitempty" name:"NewOrgPermission"`
	// 变更后代付人名称

	NewPayName *string `json:"NewPayName,omitempty" name:"NewPayName"`
	// 变更后代付人

	NewPayUin *string `json:"NewPayUin,omitempty" name:"NewPayUin"`
	// 变更前权限

	OldOrgPermission []*OrgPermission `json:"OldOrgPermission,omitempty" name:"OldOrgPermission"`
	// 变更前代付人名称

	OldPayName *string `json:"OldPayName,omitempty" name:"OldPayName"`
	// 变更前代付人

	OldPayUin *string `json:"OldPayUin,omitempty" name:"OldPayUin"`
	// 变更操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 变更状态 Valid-待确认 Accept-已生效 Denied-已拒绝 Canceled-已取消 InValid-已失效

	Status *string `json:"Status,omitempty" name:"Status"`
}

type MemberTagCompliance struct {

	// 成员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 部门ID。

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// 部门名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 不合规数。

	NonCompliant *string `json:"NonCompliant,omitempty" name:"NonCompliant"`
}

type NodeMainInfo struct {

	// 部门ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 部门名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type NodeTag struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type NotAllowReason struct {

	// 是否允许解除成员财务权限。true-是、false-否；成员不能解除财务权限时不允许删除

	BillingPermission *bool `json:"BillingPermission,omitempty" name:"BillingPermission"`
	// 成员删除许可。true-开启、false-关闭；成员删除许可关闭时不允许删除

	DeletionPermission *bool `json:"DeletionPermission,omitempty" name:"DeletionPermission"`
	// 检测失败的资源列表。账号有资源检测失败时不允许删除。

	DetectFailedResources []*string `json:"DetectFailedResources,omitempty" name:"DetectFailedResources"`
	// 存在的资源列表。账号存在资源时不允许删除

	ExistResources []*string `json:"ExistResources,omitempty" name:"ExistResources"`
	// 是否可信服务委派管理员。true-是、false-否；成员是可信服务委派管理员不允许删除

	IsAssignManager *bool `json:"IsAssignManager,omitempty" name:"IsAssignManager"`
	// 是否主体管理员。true-是、false-否；成员是主体管理员不允许删除

	IsAuthManager *bool `json:"IsAuthManager,omitempty" name:"IsAuthManager"`
	// 是否创建的成员。true-是、false-否；成员不是创建的成员不允许删除

	IsCreateMember *bool `json:"IsCreateMember,omitempty" name:"IsCreateMember"`
	// 是否共享资源管理员。true-是、false-否；成员是共享资源管理员不允许删除

	IsShareManager *bool `json:"IsShareManager,omitempty" name:"IsShareManager"`
	// 成员是否设置了操作审批。true-是、false-否；成员设置了操作审批时不允许删除

	OperateProcess *bool `json:"OperateProcess,omitempty" name:"OperateProcess"`
}

type OrgApply struct {

	// 审批人。

	Approver []*string `json:"Approver,omitempty" name:"Approver"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 管理员实名名称。

	HostAuthName *string `json:"HostAuthName,omitempty" name:"HostAuthName"`
	// 管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 申请ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 状态。可用：Valid。拒绝：Denied。接受：Approve。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgAuthPolicy struct {

	// 关联用户组。

	AuthGroups []*AuthGroup `json:"AuthGroups,omitempty" name:"AuthGroups"`
	// 关联子用户。

	AuthSubAccounts []*OrgAuthSubAccount `json:"AuthSubAccounts,omitempty" name:"AuthSubAccounts"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 身份ID。

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份的角色别名。

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 身份的角色名。

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 授权成员。

	Members []*MemberMainInfo `json:"Members,omitempty" name:"Members"`
	// 策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgAuthRelationApply struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 拒绝或同意的原因。

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 申请原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 被关联的实名主体。

	RelationAuthName *string `json:"RelationAuthName,omitempty" name:"RelationAuthName"`
	// 状态。可用：Valid。拒绝：Denied。接受：Approve。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgAuthSubAccount struct {

	// 子用户名称。

	SubAccountName *string `json:"SubAccountName,omitempty" name:"SubAccountName"`
	// 子用户uin。

	SubAccountUin *uint64 `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type OrgBeInviteRecord struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 管理员昵称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 管理员uin

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否允许成员退出。允许：Allow，不允许：Denied。

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 邀请成员UIn

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 邀请成员名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 组织ID

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 策略权限

	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 关系策略名

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 关系策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 组织类型

	OrgType *int64 `json:"OrgType,omitempty" name:"OrgType"`
	// 代付者名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 代付者Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 邀请状态，有效：Valid， 接收：Accepted，拒绝：Denied，取消：Canceled，无效：InValid

	Status *string `json:"Status,omitempty" name:"Status"`
}

type OrgCloudApplication struct {

	// 应用Id。

	ApplicationId *uint64 `json:"ApplicationId,omitempty" name:"ApplicationId"`
	// 应用描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 应用链接。

	Link *string `json:"Link,omitempty" name:"Link"`
	// 应用关联成员列表。

	Members []*MemberMainInfo `json:"Members,omitempty" name:"Members"`
	// 应用名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用状态。Enabled-开启，Disabled-关闭

	Status *string `json:"Status,omitempty" name:"Status"`
	// 二级应用列表。

	SubApplications []*SubCloudApplication `json:"SubApplications,omitempty" name:"SubApplications"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgCollPolicy struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 身份ID。

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份角色别名。

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 身份角色名。

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgCreateRecord struct {

	// 账户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 申请时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 审核时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// 审核账号

	ApprovalUin *int64 `json:"ApprovalUin,omitempty" name:"ApprovalUin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 管理员uin

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 邀请成员UIn

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 邀请成员名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 管理身份

	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
	// 策略权限

	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 关系策略名

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 关系策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 组织类型

	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`
	// 代付者名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 代付者Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 失败原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 邀请状态，有效：Valid， 接收：Accepted，拒绝：Denied，取消：Canceled，无效：InValid

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgFinancialByMonth struct {

	// 比上月增长率%。正数增长，负数下降，空值无法统计。

	GrowthRate *string `json:"GrowthRate,omitempty" name:"GrowthRate"`
	// 记录ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 月份，格式：yyyy-mm，示例：2021-01。

	Month *string `json:"Month,omitempty" name:"Month"`
	// 消耗金额，单元：元。

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type OrgInYehe struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 类型。控制台：Console；野鹤管理端：YeHe。

	CreateType *string `json:"CreateType,omitempty" name:"CreateType"`
	// 管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员数。

	MemberNum *int64 `json:"MemberNum,omitempty" name:"MemberNum"`
	// 组织ID。

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgInvitation struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 管理员uin

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否允许成员退出。允许：Allow，不允许：Denied

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 加入时间

	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
	// 邀请成员UIn

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 邀请成员名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 策略权限

	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 关系策略名

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 关系策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 组织类型

	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`
	// 代付者名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 代付者Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 邀请状态，有效：Valid， 接收：Accepted，拒绝：Denied，取消：Canceled，无效：InValid，待审核：Pending

	Status *string `json:"Status,omitempty" name:"Status"`
}

type OrgMemberFinancial struct {

	// 成员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 占比%。

	Ratio *string `json:"Ratio,omitempty" name:"Ratio"`
	// 消耗金额，单位：元。

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type OrgMemberInYeHe struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员实名名称。

	MemberAuthName *string `json:"MemberAuthName,omitempty" name:"MemberAuthName"`
	// 成员名。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员类类型。控制台创建：Create，控制台邀请：Invite，野鹤添加：YeHeCreate，野鹤导入：YeHeImport，计费添加：BillingCreate，集团管理员：Admin

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 节点ID。

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 组织ID。

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 访问权限

	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
	// 策略权限。

	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 策略名称。

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 策略类型。

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 代付者名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 代付者Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgMemberPolicy struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 身份ID。

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份角色别名。

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 身份角色名。

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgMembersAuthAccount struct {

	// 成员访问策略列表。

	AuthPolicies []*OrgMemberPolicy `json:"AuthPolicies,omitempty" name:"AuthPolicies"`
	// 成员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 集团管理员子账号名称。

	OrgSubAccountName *string `json:"OrgSubAccountName,omitempty" name:"OrgSubAccountName"`
	// 集团管理员子账号uin。

	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

type OrgMembersAuthPolicy struct {

	// 绑定类型。1-子账号、2-用户组

	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 身份Id。

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 身份的角色别名。

	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`
	// 身份的角色名。

	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`
	// 成员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员信息。

	Members []*MemberMainInfo `json:"Members,omitempty" name:"Members"`
	// 成员uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 子账号名称或者用户组名称。

	OrgSubAccountName *string `json:"OrgSubAccountName,omitempty" name:"OrgSubAccountName"`
	// 子账号uin或者用户组Id。

	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
	// 成员访问策略Id。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 成员访问策略名称。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type OrgNodeMemberRecord struct {

	// 移动前节点ID。

	BeforeNodeId *int64 `json:"BeforeNodeId,omitempty" name:"BeforeNodeId"`
	// 移动前节点名。

	BeforeNodeName *string `json:"BeforeNodeName,omitempty" name:"BeforeNodeName"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 当前节点ID。

	CurrentNodeId *int64 `json:"CurrentNodeId,omitempty" name:"CurrentNodeId"`
	// 当前节点名。

	CurrentNodeName *string `json:"CurrentNodeName,omitempty" name:"CurrentNodeName"`
	// 记录ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 成员名。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 组织ID。

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgNodeRecord struct {

	// 之前的节点ID。不存在为0。

	BeforeNodeId *int64 `json:"BeforeNodeId,omitempty" name:"BeforeNodeId"`
	// 之前的节点名。不存在为空。

	BeforeNodeName *string `json:"BeforeNodeName,omitempty" name:"BeforeNodeName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 当前节点ID。不存在为0。

	CurrentNodeId *int64 `json:"CurrentNodeId,omitempty" name:"CurrentNodeId"`
	// 当前节点名。不存在为空。

	CurrentNodeName *string `json:"CurrentNodeName,omitempty" name:"CurrentNodeName"`
	// 记录ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 组织ID

	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`
	// 节点动作。Create：创建，Delete：删除，Move：移动。

	RecordAction *string `json:"RecordAction,omitempty" name:"RecordAction"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgPermission struct {

	// 权限Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 权限名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type OrgPolicy struct {

	// 默认策略权限

	DefaultPermission []*OrgPermission `json:"DefaultPermission,omitempty" name:"DefaultPermission"`
	// 关系策略名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 权限列表

	Permission []*OrgPermission `json:"Permission,omitempty" name:"Permission"`
	// 关系策略类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type OrgProductFinancial struct {

	// 产品名。

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品Code。

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 占比%。

	Ratio *string `json:"Ratio,omitempty" name:"Ratio"`
	// 产品消耗，单位：元。

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type OrgRecord struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 邀请成员UIn

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 邀请成员名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 记录类型，邀请：Invite，创建：Create，退出：Quit，删除：Delete

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 邀请状态，有效：Valid， 接收：Accepted，拒绝：Denied，取消：Canceled，无效：InValid，成功：Success，失败：Failed

	Status *string `json:"Status,omitempty" name:"Status"`
	// 过期时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgSubAccountByDay struct {

	// 日期。格式：yyyy-mm-dd，示例：2021-06-09。

	Day *string `json:"Day,omitempty" name:"Day"`
	// 记录ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 当天成员总数。

	MemberNum *int64 `json:"MemberNum,omitempty" name:"MemberNum"`
	// 当天子账号总数。

	SubAccountNum *int64 `json:"SubAccountNum,omitempty" name:"SubAccountNum"`
}

type OrgSubAccountByMonth struct {

	// 记录ID。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 当月成员总数。

	MemberNum *int64 `json:"MemberNum,omitempty" name:"MemberNum"`
	// 月份。格式：yyyy-mm，示例：2021-06。

	Month *string `json:"Month,omitempty" name:"Month"`
	// 当月子账号总数。

	SubAccountNum *int64 `json:"SubAccountNum,omitempty" name:"SubAccountNum"`
}

type OrganizationService struct {

	// 委派管理员数量

	AssignCount *uint64 `json:"AssignCount,omitempty" name:"AssignCount"`
	// 产品控制台路径

	ConsoleUrl *string `json:"ConsoleUrl,omitempty" name:"ConsoleUrl"`
	// 是否支持创建服务角色 1-支持、2-不支持

	CreateRole *uint64 `json:"CreateRole,omitempty" name:"CreateRole"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 英文描述

	DescriptionEn *string `json:"DescriptionEn,omitempty" name:"DescriptionEn"`
	// 帮助文档

	Document *string `json:"Document,omitempty" name:"Document"`
	// 可信服务授权接口

	GrantAction *OrganizationServiceGrantAction `json:"GrantAction,omitempty" name:"GrantAction"`
	// 集团服务ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否支持委派 1-是、2-否

	IsAssign *uint64 `json:"IsAssign,omitempty" name:"IsAssign"`
	// 是否开放 1-是、2-否

	IsOpen *uint64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 是否支持设置管理范围 1-是、2-否

	IsSetManagementScope *uint64 `json:"IsSetManagementScope,omitempty" name:"IsSetManagementScope"`
	// 是否接入使用状态 1-是、2-否

	IsUsageStatus *uint64 `json:"IsUsageStatus,omitempty" name:"IsUsageStatus"`
	// 产品简称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名称

	ProductNameEn *string `json:"ProductNameEn,omitempty" name:"ProductNameEn"`
	// 产品委派管理员调用organization:DescribeOrganizationMembers 返回的字段

	ReturnColumn []*string `json:"ReturnColumn,omitempty" name:"ReturnColumn"`
	// 角色信息

	RoleInfo []*ServiceRoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`
	// 是否使用可信服务授权 1-是、2-否

	ServiceGrant *uint64 `json:"ServiceGrant,omitempty" name:"ServiceGrant"`
	// 查询成员使用状态接口url

	UsageStatusUrl *string `json:"UsageStatusUrl,omitempty" name:"UsageStatusUrl"`
}

type OrganizationServiceAssign struct {

	// 委派管理员数量限制。

	CanAssignCount *uint64 `json:"CanAssignCount,omitempty" name:"CanAssignCount"`
	// 集团服务产品控制台路径。

	ConsoleUrl *string `json:"ConsoleUrl,omitempty" name:"ConsoleUrl"`
	// 集团服务描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 帮助文档。

	Document *string `json:"Document,omitempty" name:"Document"`
	// 集团服务授权启用状态。ServiceGrant值为1时该字段有效 ，取值：Enabled-开启 Disabled-关闭

	GrantStatus *string `json:"GrantStatus,omitempty" name:"GrantStatus"`
	// 是否支持委派。取值: 1-是 2-否

	IsAssign *uint64 `json:"IsAssign,omitempty" name:"IsAssign"`
	// 是否支持设置委派管理范围。取值: 1-是 2-否

	IsSetManagementScope *uint64 `json:"IsSetManagementScope,omitempty" name:"IsSetManagementScope"`
	// 是否接入使用状态。取值: 1-是 2-否

	IsUsageStatus *uint64 `json:"IsUsageStatus,omitempty" name:"IsUsageStatus"`
	// 当前委派管理员数。

	MemberNum *string `json:"MemberNum,omitempty" name:"MemberNum"`
	// 集团服务产品标识。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集团服务产品名称。

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 是否支持集团服务授权。取值 1-是、2-否

	ServiceGrant *uint64 `json:"ServiceGrant,omitempty" name:"ServiceGrant"`
	// 集团服务ID。

	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

type OrganizationServiceAssignMember struct {

	// 委派时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 委派管理员管理范围。取值: 1-全部成员 2-部分成员

	ManagementScope *uint64 `json:"ManagementScope,omitempty" name:"ManagementScope"`
	// 管理的成员Uin列表。ManagementScope值为2时该参数有效

	ManagementScopeMembers []*MemberMainInfo `json:"ManagementScopeMembers,omitempty" name:"ManagementScopeMembers"`
	// 管理的部门ID列表。ManagementScope值为2时该参数有效

	ManagementScopeNodes []*NodeMainInfo `json:"ManagementScopeNodes,omitempty" name:"ManagementScopeNodes"`
	// 委派管理员名称。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 委派管理员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 集团服务产品名称。

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 集团服务ID。

	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 启用状态 。取值：0-服务无启用状态 1-已启用 2-未启用

	UsageStatus *uint64 `json:"UsageStatus,omitempty" name:"UsageStatus"`
}

type OrganizationServiceGrantAction struct {

	// 授权可信服务管理员接口列表

	AssignGrantActions []*string `json:"AssignGrantActions,omitempty" name:"AssignGrantActions"`
	// 授权集团管理员接口列表

	HostGrantActions []*string `json:"HostGrantActions,omitempty" name:"HostGrantActions"`
}

type PlatformAccessContent struct {

	// 配置内容。

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
	// 配置说明。

	ConfigRemark *string `json:"ConfigRemark,omitempty" name:"ConfigRemark"`
	// 配置状态。接入中、已接入

	ConfigStatus *string `json:"ConfigStatus,omitempty" name:"ConfigStatus"`
}

type ProductResource struct {

	// 产品资源ID。

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 资源六段式最后一节

	ResourceGrantLast *string `json:"ResourceGrantLast,omitempty" name:"ResourceGrantLast"`
}

type RelationParam struct {

	// 参数名

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 参数值

	ParamValue *string `json:"ParamValue,omitempty" name:"ParamValue"`
}

type ResourceTagMapping struct {

	// 合规详情。

	ComplianceDetails *TagComplianceDetails `json:"ComplianceDetails,omitempty" name:"ComplianceDetails"`
	// 资源六段式。使用资源六段式描述一个资源。 例如：qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。

	Resource *string `json:"Resource,omitempty" name:"Resource"`
	// 资源标签。

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
}

type ResourceType struct {

	// 资源列表API入参转换

	ApiTransformInFunc *string `json:"ApiTransformInFunc,omitempty" name:"ApiTransformInFunc"`
	// 资源列表API出参转换

	ApiTransformOutFunc *string `json:"ApiTransformOutFunc,omitempty" name:"ApiTransformOutFunc"`
	// 资源是否区分地域 1-区分、2-不区分

	AreaType *int64 `json:"AreaType,omitempty" name:"AreaType"`
	// 是否允许自行取消共享 1-允许、2-不允许

	CanCancel *int64 `json:"CanCancel,omitempty" name:"CanCancel"`
	// 共享描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 共享资源名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品简称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 资源列表API

	ResourceApi *string `json:"ResourceApi,omitempty" name:"ResourceApi"`
	// 资源授权格式

	ResourceGrantModel *string `json:"ResourceGrantModel,omitempty" name:"ResourceGrantModel"`
	// 共享级别 1-集团管理员、2-成员、3-组织外

	ShareLevel *int64 `json:"ShareLevel,omitempty" name:"ShareLevel"`
	// 是否显示资源使用状态 1-显示、2-不显示

	ShowUsageStatus *int64 `json:"ShowUsageStatus,omitempty" name:"ShowUsageStatus"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 白名单key

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

type ResourceTypeInfo struct {

	// 是否允许主动退出共享 1-允许、2-不允许

	AllowExist *uint64 `json:"AllowExist,omitempty" name:"AllowExist"`
	// 资源列表API入参转换

	ApiTransformInFunc *string `json:"ApiTransformInFunc,omitempty" name:"ApiTransformInFunc"`
	// 资源列表API出参转换

	ApiTransformOutFunc *string `json:"ApiTransformOutFunc,omitempty" name:"ApiTransformOutFunc"`
	// 资源是否区分地域 1-区分、2-不区分

	AreaType *int64 `json:"AreaType,omitempty" name:"AreaType"`
	// 是否允许自行取消共享 1-允许、2-不允许

	CanCancel *int64 `json:"CanCancel,omitempty" name:"CanCancel"`
	// 共享描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 可共享资源类型ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否开启 1-开启、2-不开启

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源英文名称

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 授权接口

	PolicyActions []*string `json:"PolicyActions,omitempty" name:"PolicyActions"`
	// 产品简称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 资源列表API

	ResourceApi *string `json:"ResourceApi,omitempty" name:"ResourceApi"`
	// 资源授权格式

	ResourceGrantModel *string `json:"ResourceGrantModel,omitempty" name:"ResourceGrantModel"`
	// 共享场景描述

	SharedScene *string `json:"SharedScene,omitempty" name:"SharedScene"`
	// 允许成员间共享 1-允许、2-不允许

	ShareLevel *int64 `json:"ShareLevel,omitempty" name:"ShareLevel"`
	// 共享资源授权 1-是、2-否

	ShareResourceGrant *uint64 `json:"ShareResourceGrant,omitempty" name:"ShareResourceGrant"`
	// 共享资源展示 1-是、2-否

	ShareResourceShow *uint64 `json:"ShareResourceShow,omitempty" name:"ShareResourceShow"`
	// 是否显示资源使用状态 1-显示、2-不显示

	ShowUsageStatus *int64 `json:"ShowUsageStatus,omitempty" name:"ShowUsageStatus"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 资源使用状态接口地址

	UsageStatusUrl *string `json:"UsageStatusUrl,omitempty" name:"UsageStatusUrl"`
	// 白名单key

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

type ServiceRoleInfo struct {

	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 关联策略名称

	RolePolicyName []*string `json:"RolePolicyName,omitempty" name:"RolePolicyName"`
	// 角色类型 ServiceRole：服务角色 、ServiceLinkedRole：服务相关角色

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
	// 服务载体

	ServiceName []*string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ShareArea struct {

	// 地域标识。

	Area *string `json:"Area,omitempty" name:"Area"`
	// 地域ID。

	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`
	// 地域名称。

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ShareMember struct {

	// 共享成员Uin。

	ShareMemberUin *int64 `json:"ShareMemberUin,omitempty" name:"ShareMemberUin"`
}

type ShareResource struct {

	// 产品资源ID。

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 共享资源ID。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type ShareResourceToMember struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 业务资源ID

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 共享管理员uin

	ShareManagerUin *int64 `json:"ShareManagerUin,omitempty" name:"ShareManagerUin"`
	// 资源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 共享单元ID

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
	// 共享单元名

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type ShareResourceUsageRecordInfo struct {

	// 使用时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 关联资源ID

	RelationResourceId *string `json:"RelationResourceId,omitempty" name:"RelationResourceId"`
	// 关联资源类型

	RelationResourceType *string `json:"RelationResourceType,omitempty" name:"RelationResourceType"`
	// 共享账号ID

	ShareMemberUin *int64 `json:"ShareMemberUin,omitempty" name:"ShareMemberUin"`
}

type ShareUnitMainInfo struct {

	// 共享单元地域。

	Area *string `json:"Area,omitempty" name:"Area"`
	// 共享单元管理员Uin。

	ShareManagerUin *int64 `json:"ShareManagerUin,omitempty" name:"ShareManagerUin"`
	// 共享单元ID。

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
}

type ShareUnitMember struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 共享成员Uin。

	ShareMemberUin *int64 `json:"ShareMemberUin,omitempty" name:"ShareMemberUin"`
}

type ShareUnitMemberRecord struct {

	// 退出共享时间。

	ExitTime *string `json:"ExitTime,omitempty" name:"ExitTime"`
	// 共享成员Uin。

	ShareMemberUin *int64 `json:"ShareMemberUin,omitempty" name:"ShareMemberUin"`
	// 共享成员状态 1-共享中 2-待确认 3-拒绝共享 4-退出共享

	ShareStatus *uint64 `json:"ShareStatus,omitempty" name:"ShareStatus"`
	// 共享时间。

	ShareTime *string `json:"ShareTime,omitempty" name:"ShareTime"`
}

type ShareUnitResource struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 产品资源ID。

	ProductResourceId *string `json:"ProductResourceId,omitempty" name:"ProductResourceId"`
	// 共享资源ID。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 共享单元成员数。

	SharedMemberNum *uint64 `json:"SharedMemberNum,omitempty" name:"SharedMemberNum"`
	// 使用中共享单元成员数。

	SharedMemberUseNum *uint64 `json:"SharedMemberUseNum,omitempty" name:"SharedMemberUseNum"`
	// 共享管理员OwnerUin。

	ShareManagerUin *int64 `json:"ShareManagerUin,omitempty" name:"ShareManagerUin"`
	// 共享资源类型。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ShareUnitToMember struct {

	// 是否允许主动退出共享。取值：1-允许 2-不允许

	AllowExist *uint64 `json:"AllowExist,omitempty" name:"AllowExist"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 退出共享时间

	ExitTime *string `json:"ExitTime,omitempty" name:"ExitTime"`
	// 共享单元资源数

	SharedResourceNum *int64 `json:"SharedResourceNum,omitempty" name:"SharedResourceNum"`
	// 共享单元管理员Uin

	ShareManagerUin *int64 `json:"ShareManagerUin,omitempty" name:"ShareManagerUin"`
	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号

	ShareScope *uint64 `json:"ShareScope,omitempty" name:"ShareScope"`
	// 共享成员状态 1-共享中 2-待确认 3-拒绝共享 4-退出共享

	ShareStatus *uint64 `json:"ShareStatus,omitempty" name:"ShareStatus"`
	// 共享时间

	ShareTime *string `json:"ShareTime,omitempty" name:"ShareTime"`
	// 共享单元ID

	UnitId *string `json:"UnitId,omitempty" name:"UnitId"`
	// 共享单元名

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type SortInfo struct {

	// 排序字段。

	SortFilter *string `json:"SortFilter,omitempty" name:"SortFilter"`
	// 排序类型。正序：ASC，逆序：DESC

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type StatusInfo struct {

	// 状态码 成功:Success

	Code *string `json:"Code,omitempty" name:"Code"`
	// 成员Uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 说明

	Message *string `json:"Message,omitempty" name:"Message"`
}

type SubAccountInfo struct {

	// 是否登陆控制台 1-是、0-否

	CanLogin *uint64 `json:"CanLogin,omitempty" name:"CanLogin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 所属部门

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// 邮箱地址

	Email *string `json:"Email,omitempty" name:"Email"`
	// 成员账号名称

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员账号uin

	MemberUin *string `json:"MemberUin,omitempty" name:"MemberUin"`
	// 手机号码

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 成员账号描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 子账号名称

	SubName *string `json:"SubName,omitempty" name:"SubName"`
	// 子账号uin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
}

type SubCloudApplication struct {

	// 应用描述。最大200个字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 应用链接。最大200个字符

	Link *string `json:"Link,omitempty" name:"Link"`
	// 应用名称。最大50个字符

	Name *string `json:"Name,omitempty" name:"Name"`
}

type TagComplianceDetails struct {

	// 合规状态。true-合规，false-不合规

	ComplianceStatus *bool `json:"ComplianceStatus,omitempty" name:"ComplianceStatus"`
	// 值不合规的标签键列表。

	KeysWithNonCompliantValues []*string `json:"KeysWithNonCompliantValues,omitempty" name:"KeysWithNonCompliantValues"`
	// 键不合规的标签键列表。

	NonCompliantKeys []*string `json:"NonCompliantKeys,omitempty" name:"NonCompliantKeys"`
}

type Tags struct {

	// 标签键。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值。

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type YeheAuthRelationsApply struct {

	// 申请ID。

	ApplyId *int64 `json:"ApplyId,omitempty" name:"ApplyId"`
	// 磐石审批流ID。

	ApproveId *string `json:"ApproveId,omitempty" name:"ApproveId"`
	// 审批人。

	Approver []*string `json:"Approver,omitempty" name:"Approver"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 材料审核详情。

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 管理员实名名称。

	HostAuthName *string `json:"HostAuthName,omitempty" name:"HostAuthName"`
	// 管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 关联的邀请成员列表。

	MemberList []*YeheRelationAuthMember `json:"MemberList,omitempty" name:"MemberList"`
	// 申请原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 互信主体列表。

	RelationAuthList []*YeheRelationAuthInfo `json:"RelationAuthList,omitempty" name:"RelationAuthList"`
	// 状态。 编辑中：Edit, 申请中：Valid, 审批中：Pending, 通过：Approve, 拒绝：Denied, 撤回：Withdraw

	Status *string `json:"Status,omitempty" name:"Status"`
	// 审批时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type YeheRelationAuthInfo struct {

	// 互信证明文件列表。

	AuthFile []*AuthRelationFile `json:"AuthFile,omitempty" name:"AuthFile"`
	// 企业实名名称。

	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`
}

type YeheRelationAuthMember struct {

	// 成员实名名称。

	MemberAuthName *string `json:"MemberAuthName,omitempty" name:"MemberAuthName"`
	// 成员名称

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 所属节点ID。

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 财务权限ID。取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、8-成本分析，1、2 默认必须

	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
}

type AcceptMemberChangePermissionRequest struct {
	*tchttp.BaseRequest

	// 变更记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *AcceptMemberChangePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptMemberChangePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptMemberChangePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptMemberChangePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptMemberChangePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *AcceptOrganizationInvitationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptOrganizationInvitationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptOrganizationInvitationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptOrganizationInvitationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeTagsRequest struct {
	*tchttp.BaseRequest

	// 组织节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 标签列表

	Tags []*NodeTag `json:"Tags,omitempty" name:"Tags"`
}

func (r *AddOrganizationNodeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOrganizationNodeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOrganizationNodeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOrganizationNodeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberEmailRequest struct {
	*tchttp.BaseRequest

	// 授权码

	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 手机验证码

	PhoneCode *string `json:"PhoneCode,omitempty" name:"PhoneCode"`
}

func (r *BindOrganizationMemberEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationMemberEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationMemberEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 组织管理员用户组Id列表。最大50个

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 成员访问策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *BindOrganizationPolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationPolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationPolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationPolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationPolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationPolicySubAccountRequest struct {
	*tchttp.BaseRequest

	// 组织管理员子账号Uin列表。最大5个

	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
	// 策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *BindOrganizationPolicySubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationPolicySubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationPolicySubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationPolicySubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindOrganizationPolicySubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelMemberChangePermissionRequest struct {
	*tchttp.BaseRequest

	// 变更记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CancelMemberChangePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelMemberChangePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelMemberChangePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelMemberChangePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelMemberChangePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CancelOrganizationInvitationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationInvitationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationInvitationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationInvitationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 组织管理员用户组Id列表。最大50个

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 成员访问策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CancelOrganizationPolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationPolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationPolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationPolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationPolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationPolicySubAccountRequest struct {
	*tchttp.BaseRequest

	// 组织管理员子账号Uin列表。最大5个

	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
	// 策略ID。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CancelOrganizationPolicySubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationPolicySubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrganizationPolicySubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelOrganizationPolicySubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrganizationPolicySubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountIsSubClientRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckAccountIsSubClientRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountIsSubClientRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountIsSubClientResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否代理商。 true-是、false-否

		IsAgent *bool `json:"IsAgent,omitempty" name:"IsAgent"`
		// 是否代理商子客。 true-是、false-否

		IsAgentClient *bool `json:"IsAgentClient,omitempty" name:"IsAgentClient"`
		// 是否经销商。 true-是、false-否

		IsDistribution *bool `json:"IsDistribution,omitempty" name:"IsDistribution"`
		// 是否经销商子客。 true-是、false-否

		IsSubClient *bool `json:"IsSubClient,omitempty" name:"IsSubClient"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAccountIsSubClientResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountIsSubClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountStatusRequest struct {
	*tchttp.BaseRequest

	// 账号Uin列表。列表最大长度10

	AccountUins []*int64 `json:"AccountUins,omitempty" name:"AccountUins"`
}

func (r *CheckAccountStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号状态列表

		AccountStatusList []*AccountStatus `json:"AccountStatusList,omitempty" name:"AccountStatusList"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAccountStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAddOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 组织管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员名。

	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckAddOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAddOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAddOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织根节点ID

		RootNodeId *int64 `json:"RootNodeId,omitempty" name:"RootNodeId"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAddOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAddOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBindOrganizationMemberEmailRequest struct {
	*tchttp.BaseRequest

	// 授权码

	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
}

func (r *CheckBindOrganizationMemberEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBindOrganizationMemberEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBindOrganizationMemberEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全手机

		Phone *string `json:"Phone,omitempty" name:"Phone"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBindOrganizationMemberEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBindOrganizationMemberEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationAuthManageUinRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckOrganizationAuthManageUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationAuthManageUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationAuthManageUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否企业组织主体管理账号。 true:是、false:否

		IsManage *bool `json:"IsManage,omitempty" name:"IsManage"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOrganizationAuthManageUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationAuthManageUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberAuthRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckOrganizationMemberAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOrganizationMemberAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberAuthRelationRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *CheckOrganizationMemberAuthRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberAuthRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberAuthRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOrganizationMemberAuthRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberAuthRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberPermissionRequest struct {
	*tchttp.BaseRequest

	// 认证主体关系ID

	AuthRelationId *int64 `json:"AuthRelationId,omitempty" name:"AuthRelationId"`
	// 成员Uin。 邀请、编辑成员必填

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 操作类型。 创建成员-Create 、邀请成员-Invite、 编辑成员-Update

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	// 代付者Uin。

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 策略权限。

	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
	// 策略类型。

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

func (r *CheckOrganizationMemberPermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberPermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOrganizationMemberPermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOrganizationMemberPermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOrganizationMemberPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrgMemberProductServiceRoleRequest struct {
	*tchttp.BaseRequest

	// 成员Uin

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 产品简称

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateOrgMemberProductServiceRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrgMemberProductServiceRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrgMemberProductServiceRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrgMemberProductServiceRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrgMemberProductServiceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrgServiceAssignRequest struct {
	*tchttp.BaseRequest

	// 委派管理员管理范围。 取值：1-全部成员 2-部分成员，默认值1

	ManagementScope *uint64 `json:"ManagementScope,omitempty" name:"ManagementScope"`
	// 管理的部门ID列表。ManagementScope为2时该参数有效

	ManagementScopeNodeIds []*int64 `json:"ManagementScopeNodeIds,omitempty" name:"ManagementScopeNodeIds"`
	// 管理的成员Uin列表。ManagementScope为2时该参数有效

	ManagementScopeUins []*int64 `json:"ManagementScopeUins,omitempty" name:"ManagementScopeUins"`
	// 委派管理员Uin列表。 最大长度20个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 集团服务产品标识。和集团服务ID二选一必填，可以通过ListOrganizationService获取

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集团服务ID。和集团服务产品标识二选一必填，可以通过ListOrganizationService获取

	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *CreateOrgServiceAssignRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrgServiceAssignRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrgServiceAssignResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrgServiceAssignResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 创建人。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 组织管理员Uin。

	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员列表。列表最大长度10

	Members []*AddOrgMember `json:"Members,omitempty" name:"Members"`
}

func (r *CreateOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 添加状态列表

		CreateStatus []*StatusInfo `json:"CreateStatus,omitempty" name:"CreateStatus"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrgServiceAssignRequest struct {
	*tchttp.BaseRequest

	// 委派管理员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 集团服务产品标识。和集团服务ID二选一必填，可以通过ListOrganizationService获取

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集团服务ID。和集团服务产品标识二选一必填，可以通过ListOrganizationService获取

	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DeleteOrgServiceAssignRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrgServiceAssignRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrgServiceAssignResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrgServiceAssignResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeTagsRequest struct {
	*tchttp.BaseRequest

	// 组织节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 标签列表

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *DeleteOrganizationNodeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationNodeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationNodeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationNodeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyMemberChangePermissionRequest struct {
	*tchttp.BaseRequest

	// 变更记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DenyMemberChangePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyMemberChangePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyMemberChangePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DenyMemberChangePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyMemberChangePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationCreateRecordRequest struct {
	*tchttp.BaseRequest

	// 创建记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DenyOrganizationCreateRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyOrganizationCreateRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationCreateRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DenyOrganizationCreateRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyOrganizationCreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationInvitationRequest struct {
	*tchttp.BaseRequest

	// 邀请ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DenyOrganizationInvitationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyOrganizationInvitationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenyOrganizationInvitationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DenyOrganizationInvitationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DenyOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllParentNodesRequest struct {
	*tchttp.BaseRequest

	// 节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DescribeAllParentNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllParentNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllParentNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点列表

		Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllParentNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllParentNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEffectivePolicyRequest struct {
	*tchttp.BaseRequest

	// 账号uin或者节点id。

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
}

func (r *DescribeEffectivePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEffectivePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEffectivePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 有效策略。

		EffectivePolicy *EffectivePolicy `json:"EffectivePolicy,omitempty" name:"EffectivePolicy"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEffectivePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEffectivePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemberChangePermissionRecordsRequest struct {
	*tchttp.BaseRequest

	// 国际站：en，国内站：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索字段。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeMemberChangePermissionRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemberChangePermissionRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemberChangePermissionRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员权限变更记录列表

		Items []*MemberPermissionChangeRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMemberChangePermissionRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemberChangePermissionRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemberDeletionPermissionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMemberDeletionPermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemberDeletionPermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemberDeletionPermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员删除许可状态。Enabled-开启、Disabled-关闭

		Status *string `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMemberDeletionPermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemberDeletionPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationAuthPoliciesRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。默认值：10。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量。取值是limit的整数倍。默认值 : 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 成员访问策略Id。

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 支持策略名称或者Id查询。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationAuthPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationAuthPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationAuthPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权策略列表。

		Items []*OrgAuthPolicy `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationAuthPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationAuthPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationBeInviteRecordRequest struct {
	*tchttp.BaseRequest

	// 创建起始时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 国际站：en，国内站：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationBeInviteRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationBeInviteRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationBeInviteRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 邀请信息列表

		Items []*OrgBeInviteRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationBeInviteRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationBeInviteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationCollPoliciesRequest struct {
	*tchttp.BaseRequest

	// 限制数目。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 偏移量。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationCollPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationCollPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationCollPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Items []*OrgCollPolicy `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationCollPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationCollPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationCreateRecordRequest struct {
	*tchttp.BaseRequest

	// 国际站：en，国内站：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索字段。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationCreateRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationCreateRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationCreateRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Items []*OrgCreateRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationCreateRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationCreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByMemberRequest struct {
	*tchttp.BaseRequest

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份。

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
	// 限制数目。取值范围：1~50，默认值：10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询成员列表。 最大100个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 查询开始月份。格式：yyyy-mm，例如：2021-01。

	Month *string `json:"Month,omitempty" name:"Month"`
	// 偏移量。取值是limit的整数倍，默认值 : 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询产品列表。 最大100个

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员消耗详情。

		Items []*OrgMemberFinancial `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 当月总消耗。

		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByMonthRequest struct {
	*tchttp.BaseRequest

	// 查询结束月份。格式：yyyy-mm，例如：2021-01

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
	// 查询月数。取值范围：1~6，默认值：6

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询成员列表。 最大100个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 查询产品列表。 最大100个

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMonthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByMonthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByMonthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品消耗详情。

		Items []*OrgFinancialByMonth `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMonthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByProductRequest struct {
	*tchttp.BaseRequest

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
	// 限制数目。取值范围：1~50，默认值：10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询成员列表。 最大100个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 查询开始月份。格式：yyyy-mm，例如：2021-01

	Month *string `json:"Month,omitempty" name:"Month"`
	// 偏移量。取值是limit的整数倍，默认值 : 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询产品列表。 最大100个

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialByProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品消耗详情。

		Items []*OrgProductFinancial `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 当月总消耗。

		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationFinancialByProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialMemberNumRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOrganizationFinancialMemberNumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialMemberNumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationFinancialMemberNumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主体成员财务统计数量。

		AuthInfos []*FinancialMemberNum `json:"AuthInfos,omitempty" name:"AuthInfos"`
		// 组织代付费成员数。

		BehalfPay *int64 `json:"BehalfPay,omitempty" name:"BehalfPay"`
		// 组织自付费优惠继承成员数。

		SelfPayInherit *int64 `json:"SelfPayInherit,omitempty" name:"SelfPayInherit"`
		// 组织自付费非优惠继承成员数。

		SelfPayNoInherit *int64 `json:"SelfPayNoInherit,omitempty" name:"SelfPayNoInherit"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationFinancialMemberNumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationFinancialMemberNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationIdentityRequest struct {
	*tchttp.BaseRequest

	// 身份ID

	IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
}

func (r *DescribeOrganizationIdentityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationIdentityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationIdentityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 身份名称

		IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
		// 身份ID

		IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
		// 身份策略

		IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitempty" name:"IdentityPolicy"`
		// 身份类型 1-预设 2-自定义

		IdentityType *uint64 `json:"IdentityType,omitempty" name:"IdentityType"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationIdentityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationInviteRecordRequest struct {
	*tchttp.BaseRequest

	// 国际站：en，国内站：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索字段。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationInviteRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationInviteRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationInviteRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 邀请信息列表

		Items []*OrgInvitation `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationInviteRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationInviteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberNodesRequest struct {
	*tchttp.BaseRequest

	// 成员uin列表。最大50

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
}

func (r *DescribeOrganizationMemberNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员节点列表

		MemberNodes []*MemberNodes `json:"MemberNodes,omitempty" name:"MemberNodes"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberPoliciesRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。默认值：10。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 偏移量。取值是limit的整数倍。默认值 : 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索关键字。可用于策略名或描述搜索

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMemberPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMemberPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表。

		Items []*OrgMemberPolicy `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMemberPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。默认值：10。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 成员uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 偏移量。取值是limit的整数倍。默认值 : 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 集团管理员子账号uin。

	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
	// 支持成员uin或者子账号uin查询。

	SearchKey *int64 `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMembersAuthAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersAuthAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersAuthAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问授权关系列表。

		Items []*OrgMembersAuthAccount `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersAuthAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersAuthPolicyRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。默认值：10。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 成员uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 偏移量。取值是limit的整数倍。默认值 : 0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 集团管理员子账号uin。

	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
	// 成员访问策略Id。

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 支持成员uin或者子账号uin查询。

	SearchKey *int64 `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMembersAuthPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersAuthPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersAuthPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问授权策略列表。

		Items []*OrgMembersAuthPolicy `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersAuthPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersAuthPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersCanAuthIdentitiesRequest struct {
	*tchttp.BaseRequest

	// 成员uin列表。最大10个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
}

func (r *DescribeOrganizationMembersCanAuthIdentitiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersCanAuthIdentitiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersCanAuthIdentitiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可授权身份列表。

		Items []*OrgMemberAuthIdentity `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersCanAuthIdentitiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersCanAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeByNameRequest struct {
	*tchttp.BaseRequest

	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeOrganizationNodeByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点列表

		Level *int64 `json:"Level,omitempty" name:"Level"`
		// 当前节点信息

		Node []*OrgNode `json:"Node,omitempty" name:"Node"`
		// 当前节点的父节点信息列表

		ParentNodes []*OrgNode `json:"ParentNodes,omitempty" name:"ParentNodes"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMemberRecordsRequest struct {
	*tchttp.BaseRequest

	// 限制数目。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索关键字。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationNodeMemberRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMemberRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeMemberRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Items []*OrgNodeMemberRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeMemberRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeMemberRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeRecordsRequest struct {
	*tchttp.BaseRequest

	// 限制数目。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索关键字。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationNodeRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 邀请信息列表

		Items []*OrgNodeRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeTagsRequest struct {
	*tchttp.BaseRequest

	// 组织节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DescribeOrganizationNodeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationNodeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点ID

		NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
		// 节点标签列表

		Tags []*NodeTag `json:"Tags,omitempty" name:"Tags"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationNodeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationNodeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationOverViewRequest struct {
	*tchttp.BaseRequest

	// 不传默认获取前一天数据。格式：yyyy-mm-dd，示例：2021-06-10。 数据按天统计，当天可以获取前一天之前的数据。

	Day *string `json:"Day,omitempty" name:"Day"`
}

func (r *DescribeOrganizationOverViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationOverViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationOverViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 互信主体数量。

		AuthRelationNum *int64 `json:"AuthRelationNum,omitempty" name:"AuthRelationNum"`
		// 集团管理主体名称。

		HostAuthName *string `json:"HostAuthName,omitempty" name:"HostAuthName"`
		// 成员登录权限数。

		IdentityNum *int64 `json:"IdentityNum,omitempty" name:"IdentityNum"`
		// 可登录的成员数。仅子账号查询时有效

		LoginMemberNum *int64 `json:"LoginMemberNum,omitempty" name:"LoginMemberNum"`
		// 可登录成员控制台的子用户数。仅主账号查询时有效

		LoginSubAccountNum *int64 `json:"LoginSubAccountNum,omitempty" name:"LoginSubAccountNum"`
		// 成员总数。

		MemberNum *int64 `json:"MemberNum,omitempty" name:"MemberNum"`
		// 节点总数。

		NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
		// 角色总数。

		RoleNum *int64 `json:"RoleNum,omitempty" name:"RoleNum"`
		// 子账号总数。

		SubAccountNum *int64 `json:"SubAccountNum,omitempty" name:"SubAccountNum"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationOverViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationOverViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationPendingCreateRecordRequest struct {
	*tchttp.BaseRequest

	// 国际站：en，国内站：zh

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationPendingCreateRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationPendingCreateRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationPendingCreateRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Items []*OrgCreateRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationPendingCreateRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationPendingCreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationRecordsRequest struct {
	*tchttp.BaseRequest

	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Items []*OrgRecord `json:"Items,omitempty" name:"Items"`
		// 总数目

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationServiceRoleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOrganizationServiceRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationServiceRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationServiceRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 企业组织角色

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationServiceRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationServiceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationSubAccountByDayRequest struct {
	*tchttp.BaseRequest

	// 最近的天数。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOrganizationSubAccountByDayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationSubAccountByDayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationSubAccountByDayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织子账号详情。

		Items []*OrgSubAccountByDay `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationSubAccountByDayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationSubAccountByDayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationSubAccountByMonthRequest struct {
	*tchttp.BaseRequest

	// 最近的月份。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOrganizationSubAccountByMonthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationSubAccountByMonthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationSubAccountByMonthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织子账号和成员详情。

		Items []*OrgSubAccountByMonth `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationSubAccountByMonthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationSubAccountByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportCreationRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeReportCreationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportCreationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportCreationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cos链接生成报告的COS地址。

		COSLocationList []*string `json:"COSLocationList,omitempty" name:"COSLocationList"`
		// 最后生成时间。

		LastUpdatedTimestamp *uint64 `json:"LastUpdatedTimestamp,omitempty" name:"LastUpdatedTimestamp"`
		// 状态。 RUNNING-生成中、SUCCEEDED-生成成功 、FAILED-生成失败 、NO REPORT-没有生成请求

		Status *string `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportCreationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportCreationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InviteOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 互信主体证明文件列表

	AuthFile []*AuthRelationFile `json:"AuthFile,omitempty" name:"AuthFile"`
	// 是否允许成员退出。允许：Allow，不允许：Denied。

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 被邀请账户UIN

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 需要调节的节点

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 代付者Uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 关系权限

	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`
	// 关系策略

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 互信实名主体名称

	RelationAuthName *string `json:"RelationAuthName,omitempty" name:"RelationAuthName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *InviteOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InviteOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InviteOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InviteOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InviteOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListComplianceSummaryRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。默认值：10。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 成员Uin。

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 偏移量。取值是limit的整数倍。默认值 : 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListComplianceSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListComplianceSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListComplianceSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员标签合规列表。

		Items []*MemberTagCompliance `json:"Items,omitempty" name:"Items"`
		// 总数目。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListComplianceSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListComplianceSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFinancialProductRequest struct {
	*tchttp.BaseRequest
}

func (r *ListFinancialProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFinancialProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFinancialProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		Items []*FinancialProductInfo `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFinancialProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFinancialProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNonCompliantResourceRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50。

	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 成员Uin。

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 从上一页的响应中获取的下一页的Token值。 如果是第一次请求，设置为空。

	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`
	// 标签键。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *ListNonCompliantResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNonCompliantResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNonCompliantResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源及标签合规信息。

		Items []*ResourceTagMapping `json:"Items,omitempty" name:"Items"`
		// 获取的下一页的Token值。

		PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNonCompliantResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNonCompliantResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrgMemberSubAccountRequest struct {
	*tchttp.BaseRequest

	// 查询关键字，支持按账号、部门、角色查询

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 页码，默认值是 1

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页数量，必须大于 0 且小于或等于 50

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListOrgMemberSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrgMemberSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrgMemberSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 平台子账号数量

		AllPlatNum *uint64 `json:"AllPlatNum,omitempty" name:"AllPlatNum"`
		// 个人子账号数量

		AllUserNum *uint64 `json:"AllUserNum,omitempty" name:"AllUserNum"`
		// 条目详情。

		Items []*SubAccountInfo `json:"Items,omitempty" name:"Items"`
		// 总数。

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrgMemberSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrgMemberSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrgServiceAssignMemberRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50，默认值：10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量。取值是limit的整数倍，默认值 : 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 集团服务产品标识。和集团服务ID二选一必填，可以通过ListOrganizationService获取

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集团服务ID。和集团服务产品标识二选一必填，可以通过ListOrganizationService获取

	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *ListOrgServiceAssignMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrgServiceAssignMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrgServiceAssignMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 委派管理员列表。

		Items []*OrganizationServiceAssignMember `json:"Items,omitempty" name:"Items"`
		// 总数。

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrgServiceAssignMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrgServiceAssignMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationServiceRequest struct {
	*tchttp.BaseRequest

	// 限制数目。取值范围：1~50，默认值：10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量。取值是limit的整数倍，默认值 : 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 名称搜索关键字。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *ListOrganizationServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrganizationServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrganizationServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织委派管理员总数。

		AssignManageTotal *uint64 `json:"AssignManageTotal,omitempty" name:"AssignManageTotal"`
		// 集团服务列表。

		Items []*OrganizationServiceAssign `json:"Items,omitempty" name:"Items"`
		// 总数。

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrganizationServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrganizationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesForTargetRequest struct {
	*tchttp.BaseRequest

	// 搜索关键字。按照策略名称搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 页码。默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 账号uin或者节点id。

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
}

func (r *ListPoliciesForTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesForTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesForTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 目标关联的策略列表。

		List []*ListPoliciesForTarget `json:"List,omitempty" name:"List"`
		// 总数。

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesForTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesForTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendOrgMemberAccountBindEmailRequest struct {
	*tchttp.BaseRequest

	// 绑定ID。可以通过 DescribeOrganizationMemberEmailBind 获取

	BindId *int64 `json:"BindId,omitempty" name:"BindId"`
	// 成员Uin。

	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *SendOrgMemberAccountBindEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendOrgMemberAccountBindEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendOrgMemberAccountBindEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendOrgMemberAccountBindEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendOrgMemberAccountBindEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsVerifyCodeForBindPhoneRequest struct {
	*tchttp.BaseRequest

	// 授权码

	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
	// 滑动验证码AppId

	CaptchaAppId *string `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`
	// 用户ip

	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
	// 发送语言，中文：zh；英文： en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 滑动验证码随机串

	Random *string `json:"Random,omitempty" name:"Random"`
	// 滑动验证码参数

	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`
}

func (r *SendSmsVerifyCodeForBindPhoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendSmsVerifyCodeForBindPhoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendSmsVerifyCodeForBindPhoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendSmsVerifyCodeForBindPhoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendSmsVerifyCodeForBindPhoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMemberDeletionPermissionRequest struct {
	*tchttp.BaseRequest

	// 成员删除许可状态。取值：Enabled-开启、Disabled-关闭

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *SetMemberDeletionPermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMemberDeletionPermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMemberDeletionPermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMemberDeletionPermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMemberDeletionPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartReportCreationRequest struct {
	*tchttp.BaseRequest

	// 对象存储桶。

	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`
	// 对象存储地域。

	COSRegion *string `json:"COSRegion,omitempty" name:"COSRegion"`
}

func (r *StartReportCreationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartReportCreationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartReportCreationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartReportCreationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartReportCreationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略描述。最大长度为128个字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 成员访问身份ID。可通过ListOrganizationIdentity获取

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 成员Uin列表。最多10个

	MemberUins []*int64 `json:"MemberUins,omitempty" name:"MemberUins"`
	// 成员访问策略Id。可通过DescribeOrganizationMemberPolicies获取

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *UpdateOrganizationMembersPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationMembersPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeTagRequest struct {
	*tchttp.BaseRequest

	// 组织节点ID

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *UpdateOrganizationNodeTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationNodeTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationNodeTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationNodeTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

package v20210331

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type SAMLIdentityProviderConfiguration struct {
	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 元数据文档（Base64 编码）。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// X509证书ID。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type UserInfo struct {
	// 查询username。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的名。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的姓。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户状态 Enabled：启用， Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户的修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否选中
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`
}

type UserProvisioningsTask struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 授权的集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型。StartProvisioning：用户同步，DeleteProvisioning：删除用户同步
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态：InProgress: 进行中，Failed: 失败 3:Success: 成功
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 用户同步ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	//  CAM 用户同步的身份 ID。取值： 当PrincipalType取值为Group时，该值为CIC 用户组 ID（g-********）。 当PrincipalType取值为User时，该值为CIC 用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值： User：表示该 CAM 用户同步的身份是CIC 用户。 Group：表示该 CAM 用户同步的身份是CIC 用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户或者用户组名称。
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// 冲突策略。KeepBoth:两者都保留;TakeOver:替换
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。Delete:删除;Keep:保留
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`
}

type UserSyncProvisioning struct {
	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值：
	//
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

// SCIMCredential SCIM密钥信息
type SCIMCredential struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥状态，Enabled已开启，Disabled已关闭。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥类型。
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// SCIM 密钥的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SCIM 密钥的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type SendOrgMemberAccountBindEmailRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`
}

type SendOrgMemberAccountBindEmailRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`
}
type SAMLServiceProvider struct {
	// https://tencentcloudsso.com/saml/sp/z-sjw8ensa**
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SP 元数据文档（Base64 编码）。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SP 的 ACS URL。
	AcsUrl *string `json:"AcsUrl,omitnil,omitempty" name:"AcsUrl"`
}

type GetZoneSAMLServiceProviderInfoRequest struct {
	*tchttp.BaseRequest

	// 空间ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetZoneSAMLServiceProviderInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneSAMLServiceProviderInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneSAMLServiceProviderInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// saml服务提供商配置信息
		SAMLServiceProvider *SAMLServiceProvider `json:"SAMLServiceProvider,omitnil,omitempty" name:"SAMLServiceProvider"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	}
}

func (r *GetZoneSAMLServiceProviderInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneSAMLServiceProviderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetExternalSAMLIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// saml 身份提供商配置信息。
		SAMLIdentityProviderConfiguration *SAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitnil,omitempty" name:"SAMLIdentityProviderConfiguration"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExternalSAMLIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// IdentityPolicy 身份策略
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

// GroupInfo 用户组信息
type GroupInfo struct {
	// 用户组的名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户组的类型  Manual：手动创建，Synchronized：外部导入。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 用户组的修改时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组员数量。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// 如果有入参FilterUsers，用户在用户组返回true，否则返回false
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`
}

// Predefined struct for user
type UpdateGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 新的用户组名称。
	NewGroupName *string `json:"NewGroupName,omitnil,omitempty" name:"NewGroupName"`

	// 新的用户组描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 新的用户组名称。
	NewGroupName *string `json:"NewGroupName,omitnil,omitempty" name:"NewGroupName"`

	// 新的用户组描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 用户组信息。
		GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type SetExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IdP 元数据文档（Base64 编码）。  由支持 SAML 2.0 协议的 IdP 提供。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// PEM 格式的 X509 证书。指定该参数会替换所有已经存在的证书。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

type SetExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IdP 元数据文档（Base64 编码）。  由支持 SAML 2.0 协议的 IdP 提供。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// PEM 格式的 X509 证书。指定该参数会替换所有已经存在的证书。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

func (r *SetExternalSAMLIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetExternalSAMLIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type ClearExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *ClearExternalSAMLIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearExternalSAMLIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户名称。空间内必须唯一。不支持修改。  格式：包含数字、英文字母和特殊符号+ = , . @ - _ 。  长度：最大 64 个字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的姓。  长度：最大 64 个字符。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的名。  长度：最大 64 个字符。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。  长度：最大 256 个字符。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。  长度：最大 128 个字符。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户的状态。取值：  Enabled（默认值）：启用。 Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户名称。空间内必须唯一。不支持修改。  格式：包含数字、英文字母和特殊符号+ = , . @ - _ 。  长度：最大 64 个字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的姓。  长度：最大 64 个字符。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的名。  长度：最大 64 个字符。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。  长度：最大 256 个字符。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。  长度：最大 128 个字符。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户的状态。取值：  Enabled（默认值）：启用。 Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

func (r *CreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 用户详情
		UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM用户同步信息。
	UserSyncProvisionings []*UserSyncProvisioning `json:"UserSyncProvisionings,omitnil,omitempty" name:"UserSyncProvisionings"`
}

type CreateUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM用户同步信息。
	UserSyncProvisionings []*UserSyncProvisioning `json:"UserSyncProvisionings,omitnil,omitempty" name:"UserSyncProvisionings"`
}

func (r *CreateUserSyncProvisioningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSyncProvisioningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务详细。
		Tasks []*UserProvisioningsTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserSyncProvisioningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserRequestParams struct {
	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetUserRequest struct {
	*tchttp.BaseRequest

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 用户信息。
		UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的状态。取值：  Enabled：启用。 Disabled：禁用。
	NewUserStatus *string `json:"NewUserStatus,omitnil,omitempty" name:"NewUserStatus"`
}

type UpdateUserStatusRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的状态。取值：  Enabled：启用。 Disabled：禁用。
	NewUserStatus *string `json:"NewUserStatus,omitnil,omitempty" name:"NewUserStatus"`
}

func (r *UpdateUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserStatusResponseParams `json:"Response"`
}

func (r *UpdateUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleAssignmentRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 授权成员账号信息，最多授权50条。
	RoleAssignmentInfo []*RoleAssignmentInfo `json:"RoleAssignmentInfo,omitnil,omitempty" name:"RoleAssignmentInfo"`
}

type CreateRoleAssignmentRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 授权成员账号信息，最多授权50条。
	RoleAssignmentInfo []*RoleAssignmentInfo `json:"RoleAssignmentInfo,omitnil,omitempty" name:"RoleAssignmentInfo"`
}

func (r *CreateRoleAssignmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleAssignmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleAssignmentResponseParams struct {
	// 任务详情。
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleAssignmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleAssignmentResponseParams `json:"Response"`
}

func (r *CreateRoleAssignmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfo struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 授权的目标成员账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 用户授权的身份ID,如果是身份类型是CIC用户,则为用户ID; 如果是用户组，则为用户组ID;
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 用户授权的身份类型, User代表CIC用户, Group代表CIC用户组
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

type TaskStatus struct {
	// 任务状态。取值：  InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型。取值：
	// ProvisionRoleConfiguration：部署权限配置。
	// DeprovisionRoleConfiguration：解除权限配置部署。
	// CreateRoleAssignment：在成员 账号上授权。
	// DeleteRoleAssignment：移除 成员 账号上的授权。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务失败原因。
	// 说明
	// 只有Status为Failed，才会显示该参数。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

type RoleAssignmentInfo struct {
	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值：
	//
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type RoleAssignments struct {
	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CAM 用户同步的身份 ID。取值： 当PrincipalType取值为Group时，该值为CIC 用户组 ID（g-********）。 当PrincipalType取值为User时，该值为CIC 用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值： User：表示该 CAM 用户同步的身份是CIC用户。 Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户名称或者用户组名称
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`
}

type RoleConfiguration struct {
	// 权限配置配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置配名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。CIC 用户使用访问配置访问成员账号时，会话最多保持的时间。
	// 单位：秒。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。CIC 用户使用访问配置访问成员账号时，初始访问的页面地址。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`

	// 权限配置的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 权限配置的更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 如果有入参FilterTargets查询成员账号是否配置过权限，配置了返回true，否则返回false。
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`
}

type RoleConfigurationProvisionings struct {
	// Deployed: 部署成功 DeployedRequired：需要重新部署 DeployFailed：部署失败
	DeploymentStatus *string `json:"DeploymentStatus,omitnil,omitempty" name:"DeploymentStatus"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 创建时间，
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间，
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

type RolePolicie struct {
	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RolePolicyId *int64 `json:"RolePolicyId,omitnil,omitempty" name:"RolePolicyId"`

	// 权限策略名称
	RolePolicyName *string `json:"RolePolicyName,omitnil,omitempty" name:"RolePolicyName"`

	// 权限策略类型
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 自定义策略内容。仅自定义策略返回该参数。
	RolePolicyDocument *string `json:"RolePolicyDocument,omitnil,omitempty" name:"RolePolicyDocument"`

	// 权限策略被添加到权限配置的时间。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`
}

type RoleProvisioningsTask struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 授权的集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态：InProgress: 进行中，Failed: 失败 3:Success: 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

// Predefined struct for user
type ListRoleAssignmentsResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 集团账号目标账号的授权列表。
	RoleAssignments []*RoleAssignments `json:"RoleAssignments,omitnil,omitempty" name:"RoleAssignments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRoleAssignmentsResponse struct {
	*tchttp.BaseResponse
	Response *ListRoleAssignmentsResponseParams `json:"Response"`
}

func (r *ListRoleAssignmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleAssignmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesForTarget struct {
	// 策略Id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关联的账号或节点
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 关联类型 1-节点 2-用户
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略绑定时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachTime *string `json:"AttachTime,omitnil,omitempty" name:"AttachTime"`
}

// Predefined struct for user
type ListPoliciesForTargetRequestParams struct {
	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字。按照策略名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListPoliciesForTargetRequest struct {
	*tchttp.BaseRequest

	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字。按照策略名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListPoliciesForTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesForTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesForTargetResponseParams struct {
	// 总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 目标关联的策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ListPoliciesForTarget `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPoliciesForTargetResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesForTargetResponseParams `json:"Response"`
}

func (r *ListPoliciesForTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesForTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type ListRoleAssignmentsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****)，当PrincipalType取值为User时，该值为用户 ID （u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询条件，目前只支持权限配置名称查询。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListRoleAssignmentsRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****)，当PrincipalType取值为User时，该值为用户 ID （u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询条件，目前只支持权限配置名称查询。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListRoleAssignmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleAssignmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleAssignmentRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-********）， 当PrincipalType取值为User时，该值为用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 当您移除一个集团账号目标账号上使用某权限配置的最后一个授权时，是否同时解除权限配置部署。取值： DeprovisionForLastRoleAssignmentOnAccount：解除权限配置部署。 None（默认值）：不解除权限配置部署。
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitnil,omitempty" name:"DeprovisionStrategy"`
}

type DeleteRoleAssignmentRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-********）， 当PrincipalType取值为User时，该值为用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 当您移除一个集团账号目标账号上使用某权限配置的最后一个授权时，是否同时解除权限配置部署。取值： DeprovisionForLastRoleAssignmentOnAccount：解除权限配置部署。 None（默认值）：不解除权限配置部署。
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitnil,omitempty" name:"DeprovisionStrategy"`
}

func (r *DeleteRoleAssignmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleAssignmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleAssignmentResponseParams struct {
	// 任务详情
	Task *TaskInfo `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleAssignmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleAssignmentResponseParams `json:"Response"`
}

func (r *DeleteRoleAssignmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type DeleteRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

func (r *DeleteRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleConfigurationResponseParams `json:"Response"`
}

func (r *DeleteRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismantleRoleConfigurationRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type DismantleRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *DismantleRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismantleRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismantleRoleConfigurationResponseParams struct {
	// 任务详情。
	Task *RoleProvisioningsTask `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DismantleRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DismantleRoleConfigurationResponseParams `json:"Response"`
}

func (r *DismantleRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismantleRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的名称。  格式：允许英文字母、数字和特殊字符-。 长度：最大 128 个字符。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组类型  Manual：手动创建，Synchronized：外部导入
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// 用户组信息。
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *GetGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupResponseParams struct {
	// 用户组信息
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupResponseParams `json:"Response"`
}

func (r *GetGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusResponseParams struct {
	// 任务状态信息。
	TaskStatus *TaskStatus `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStatusResponseParams `json:"Response"`
}

func (r *GetTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleConfigurationRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置名称。格式：包含英文字母、数字或短划线（-）。 长度：最大 128 个字符。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。 长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。 CIC用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。 单位：秒。 取值范围：900 ~ 43200（15 分钟~12 小时）。 默认值：3600（1 小时）。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。 CIC用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。 该页面必须是腾讯云控制台页面。默认为空，表示跳转到腾讯云控制台首页。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`
}

type CreateRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置名称。格式：包含英文字母、数字或短划线（-）。 长度：最大 128 个字符。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。 长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。 CIC用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。 单位：秒。 取值范围：900 ~ 43200（15 分钟~12 小时）。 默认值：3600（1 小时）。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。 CIC用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。 该页面必须是腾讯云控制台页面。默认为空，表示跳转到腾讯云控制台首页。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`
}

func (r *CreateRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleConfigurationResponseParams struct {
	// 配置访问详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleConfigurationResponseParams `json:"Response"`
}

func (r *CreateRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type GetRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

func (r *GetRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleConfigurationResponseParams struct {
	// 权限配置详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *GetRoleConfigurationResponseParams `json:"Response"`
}

func (r *GetRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 新的权限配置描述。  长度：最大 1024 个字符。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 新的会话持续时间。  CIC 用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。  单位：秒。  取值范围：900-43200（15 分钟-12 小时）。
	NewSessionDuration *int64 `json:"NewSessionDuration,omitnil,omitempty" name:"NewSessionDuration"`

	// 新的初始访问页面。  CIC 用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。  该页面必须是腾讯云控制台页面。
	NewRelayState *string `json:"NewRelayState,omitnil,omitempty" name:"NewRelayState"`
}

type UpdateRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 新的权限配置描述。  长度：最大 1024 个字符。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 新的会话持续时间。  CIC 用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。  单位：秒。  取值范围：900-43200（15 分钟-12 小时）。
	NewSessionDuration *int64 `json:"NewSessionDuration,omitnil,omitempty" name:"NewSessionDuration"`

	// 新的初始访问页面。  CIC 用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。  该页面必须是腾讯云控制台页面。
	NewRelayState *string `json:"NewRelayState,omitnil,omitempty" name:"NewRelayState"`
}

func (r *UpdateRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConfigurationResponseParams struct {
	// 权限配置详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleConfigurationResponseParams `json:"Response"`
}

func (r *UpdateRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyDetail struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

type AddPermissionPolicyToRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 20策略，每个策略长度最大32个字符。如果要添加系统策略，建议使用RolePolicies参数。自定义策略时，数组长度最大为1。
	RolePolicyNames []*string `json:"RolePolicyNames,omitnil,omitempty" name:"RolePolicyNames"`

	// 添加的系统策略详情。
	RolePolicies []*PolicyDetail `json:"RolePolicies,omitnil,omitempty" name:"RolePolicies"`

	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。
	CustomPolicyDocument *string `json:"CustomPolicyDocument,omitnil,omitempty" name:"CustomPolicyDocument"`

	// 自定义策略内容列表（跟RolePolicyNames一一对应）
	CustomPolicyDocuments []*string `json:"CustomPolicyDocuments,omitnil,omitempty" name:"CustomPolicyDocuments"`
}

func (r *AddPermissionPolicyToRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddPermissionPolicyToRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddPermissionPolicyToRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddPermissionPolicyToRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *AddPermissionPolicyToRoleConfigurationResponseParams `json:"Response"`
}

func (r *AddPermissionPolicyToRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddPermissionPolicyToRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPermissionPoliciesInRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 按策略名称搜索
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListPermissionPoliciesInRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 按策略名称搜索
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListPermissionPoliciesInRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPermissionPoliciesInRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPermissionPoliciesInRoleConfigurationResponseParams struct {
	// 权限策略总个数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 权限策略列表。
	RolePolicies []*RolePolicie `json:"RolePolicies,omitnil,omitempty" name:"RolePolicies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPermissionPoliciesInRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ListPermissionPoliciesInRoleConfigurationResponseParams `json:"Response"`
}

func (r *ListPermissionPoliciesInRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPermissionPoliciesInRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemovePermissionPolicyFromRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 32 个字符。
	RolePolicyName *string `json:"RolePolicyName,omitnil,omitempty" name:"RolePolicyName"`

	// 策略ID。
	RolePolicyId *int64 `json:"RolePolicyId,omitnil,omitempty" name:"RolePolicyId"`
}

func (r *RemovePermissionPolicyFromRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePermissionPolicyFromRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePermissionPolicyFromRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemovePermissionPolicyFromRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *RemovePermissionPolicyFromRoleConfigurationResponseParams `json:"Response"`
}

func (r *RemovePermissionPolicyFromRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePermissionPolicyFromRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSCIMCredentialRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type CreateSCIMCredentialRequest struct {
	*tchttp.BaseRequest

	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *CreateSCIMCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSCIMCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSCIMCredentialResponseParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥类型。
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// SCIM 密钥的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SCIM 密钥的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// SCIM密钥状态，Enabled已开启，Disabled已关闭。
	CredentialStatus *string `json:"CredentialStatus,omitnil,omitempty" name:"CredentialStatus"`

	// SCIM密钥。
	CredentialSecret *string `json:"CredentialSecret,omitnil,omitempty" name:"CredentialSecret"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSCIMCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateSCIMCredentialResponseParams `json:"Response"`
}

func (r *CreateSCIMCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSCIMCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSCIMCredentialsRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

type ListSCIMCredentialsRequest struct {
	*tchttp.BaseRequest

	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

func (r *ListSCIMCredentialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSCIMCredentialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSCIMCredentialsResponseParams struct {
	// SCIM密钥数量。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// SCIM 密钥信息。
	SCIMCredentials []*SCIMCredential `json:"SCIMCredentials,omitnil,omitempty" name:"SCIMCredentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSCIMCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *ListSCIMCredentialsResponseParams `json:"Response"`
}

func (r *ListSCIMCredentialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSCIMCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSCIMCredentialRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

type DeleteSCIMCredentialRequest struct {
	*tchttp.BaseRequest

	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

func (r *DeleteSCIMCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSCIMCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSCIMCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSCIMCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSCIMCredentialResponseParams `json:"Response"`
}

func (r *DeleteSCIMCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSCIMCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSCIMSynchronizationStatusRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetSCIMSynchronizationStatusRequest struct {
	*tchttp.BaseRequest

	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetSCIMSynchronizationStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSCIMSynchronizationStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSCIMSynchronizationStatusResponseParams struct {
	// SCIM 同步状态。Enabled：启用。 Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSCIMSynchronizationStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetSCIMSynchronizationStatusResponseParams `json:"Response"`
}

func (r *GetSCIMSynchronizationStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSCIMSynchronizationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMSynchronizationStatusRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM 同步状态。Enabled：启用。Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`
}

type UpdateSCIMSynchronizationStatusRequest struct {
	*tchttp.BaseRequest

	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM 同步状态。Enabled：启用。Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`
}

func (r *UpdateSCIMSynchronizationStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMSynchronizationStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMSynchronizationStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSCIMSynchronizationStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSCIMSynchronizationStatusResponseParams `json:"Response"`
}

func (r *UpdateSCIMSynchronizationStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMSynchronizationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的名。
	NewFirstName *string `json:"NewFirstName,omitnil,omitempty" name:"NewFirstName"`

	// 用户的姓。
	NewLastName *string `json:"NewLastName,omitnil,omitempty" name:"NewLastName"`

	// 用户的显示名称。
	NewDisplayName *string `json:"NewDisplayName,omitnil,omitempty" name:"NewDisplayName"`

	// 用户的描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 用户的电子邮箱。
	NewEmail *string `json:"NewEmail,omitnil,omitempty" name:"NewEmail"`
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的名。
	NewFirstName *string `json:"NewFirstName,omitnil,omitempty" name:"NewFirstName"`

	// 用户的姓。
	NewLastName *string `json:"NewLastName,omitnil,omitempty" name:"NewLastName"`

	// 用户的显示名称。
	NewDisplayName *string `json:"NewDisplayName,omitnil,omitempty" name:"NewDisplayName"`

	// 用户的描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 用户的电子邮箱。
	NewEmail *string `json:"NewEmail,omitnil,omitempty" name:"NewEmail"`
}

func (r *UpdateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
	// 用户信息
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserResponseParams `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUserToGroupResponseParams `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJoinedGroupsForUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

type ListJoinedGroupsForUserRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

func (r *ListJoinedGroupsForUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJoinedGroupsForUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJoinedGroupsForUserResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 用户加入的用户组列表
	JoinedGroups []*JoinedGroups `json:"JoinedGroups,omitnil,omitempty" name:"JoinedGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListJoinedGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *ListJoinedGroupsForUserResponseParams `json:"Response"`
}

func (r *ListJoinedGroupsForUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJoinedGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinedGroups struct {
	// 用户组的名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组的类型。取值：
	//
	// Manual：手动创建。
	// Synchronized：外部同步。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 加入用户组的时间
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`
}

type RemoveUserFromGroupRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserFromGroupResponseParams `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProvisioningTaskStatusRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetProvisioningTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetProvisioningTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvisioningTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProvisioningTaskStatusResponseParams struct {
	// 任务状态信息。
	TaskStatus *TaskStatus `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetProvisioningTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetProvisioningTaskStatusResponseParams `json:"Response"`
}

func (r *GetProvisioningTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvisioningTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM 用户同步的 ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

type GetUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM 用户同步的 ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

func (r *GetUserSyncProvisioningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserSyncProvisioningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserSyncProvisioningResponseParams struct {
	// CAM 用户同步信息。
	UserProvisioning *UserProvisioning `json:"UserProvisioning,omitnil,omitempty" name:"UserProvisioning"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *GetUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *GetUserSyncProvisioningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserProvisioning struct {
	// CAM 用户同步的状态。取值：
	//
	// Enabled：CAM 用户同步已启用。
	// Disabled：CAM 用户同步未启用。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CAM 用户同步的状态。取值：
	// Enabled：CAM 用户同步已启用。
	// Disabled：CAM 用户同步未启用。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份名称。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组名称。
	// 当PrincipalType取值为User时，该值为CIC用户名称。
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// CAM 用户同步的身份类型。取值：
	//
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

// Predefined struct for user
type UpdateUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的iD
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 用户同步描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	NewDuplicationStateful *string `json:"NewDuplicationStateful,omitnil,omitempty" name:"NewDuplicationStateful"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitnil,omitempty" name:"NewDeletionStrategy"`
}

type UpdateUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的iD
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 用户同步描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	NewDuplicationStateful *string `json:"NewDuplicationStateful,omitnil,omitempty" name:"NewDuplicationStateful"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitnil,omitempty" name:"NewDeletionStrategy"`
}

func (r *UpdateUserSyncProvisioningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSyncProvisioningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserSyncProvisioningResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *UpdateUserSyncProvisioningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

type DeleteUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

func (r *DeleteUserSyncProvisioningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserSyncProvisioningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserSyncProvisioningResponseParams struct {
	// 任务详情。
	Tasks *UserProvisioningsTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *DeleteUserSyncProvisioningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ListGroups

type ListGroupsRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用不需要 NextToken。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数，取值范围：1~100，默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤条件。格式：<Attribute> <Operator> <Value>，不区分大小写。
	// 目前 <Attribute> 只支持 GroupName，<Operator> 只支持 eq 和 sw。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 用户组的类型。Manual：手动创建，Synchronized：外部导入。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 筛选的用户，该用户关联的用户组会返回 IsSelected=true。
	FilterUsers []*string `json:"FilterUsers,omitnil,omitempty" name:"FilterUsers"`

	// 排序字段，目前只支持 CreateTime，默认是 CreateTime。
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型：Desc 倒序，Asc 正序。需与 SortField 一起设置。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupsResponseParams struct {
	// 查询返回结果下一页的令牌，仅 IsTruncated 为 true 时显示。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 用户组列表。
	Groups []*GroupInfo `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 返回结果是否被截断。true：已截断；false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 唯一请求 ID。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupsResponseParams `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ListRoleConfigurations

type ListRoleConfigurationsRequest struct {
	*tchttp.BaseRequest

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用不需要 NextToken。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数，取值范围：1~100，默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤文本，不区分大小写。支持 RoleConfigurationName 和 Description。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 检索成员账号是否配置过权限，如果配置过返回 IsSelected: true，否则返回 false。
	FilterTargets []*int64 `json:"FilterTargets,omitnil,omitempty" name:"FilterTargets"`

	// 授权的用户 UserId 或用户组 GroupId，必须和 FilterTargets 一起设置。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`
}

func (r *ListRoleConfigurationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRoleConfigurationsResponseParams struct {
	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。true：已截断；false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 查询返回结果下一页的令牌，仅 IsTruncated 为 true 时显示。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 权限配置列表。
	RoleConfigurations []*RoleConfiguration `json:"RoleConfigurations,omitnil,omitempty" name:"RoleConfigurations"`

	// 唯一请求 ID。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRoleConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *ListRoleConfigurationsResponseParams `json:"Response"`
}

func (r *ListRoleConfigurationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Account struct {

	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}

type AccountUserProvisioning struct {

	// 同步的目标成员账号名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 同步创建的子账号名

	SyncSubName *string `json:"SyncSubName,omitempty" name:"SyncSubName"`
	// 通过创建的子账号uin

	SyncSubUin *string `json:"SyncSubUin,omitempty" name:"SyncSubUin"`
	// 同步的目标成员账号uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type CreateUserInfo struct {

	// 用户的描述。 长度：最大 1024 个字符。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户的显示名称。 长度：最大 256 个字符。

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 用户的电子邮箱。目录内必须唯一。 长度：最大 128 个字符。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户的姓。 长度：最大 64 个字符。

	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`
	// 用户的名。 长度：最大 64 个字符。

	LastName *string `json:"LastName,omitempty" name:"LastName"`
	// 用户名称。空间内必须唯一。不支持修改。 格式：包含数字、英文字母和特殊符号+ = , . @ - _ 。 长度：最大 64 个字符

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户的状态。取值： Enabled（默认值）：启用。 Disabled：禁用。

	UserStatus *string `json:"UserStatus,omitempty" name:"UserStatus"`
}

type GroupIdAndUserId struct {

	// 用户组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type GroupMembers struct {

	// 用户的描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户的显示名称。

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 用户的电子邮箱。目录内必须唯一。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户加入用户组的时间

	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
	// 用户 ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 查询username。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户状态 Enabled：启用， Disabled：禁用。

	UserStatus *string `json:"UserStatus,omitempty" name:"UserStatus"`
	// 用户类型 Manual：手动创建，Synchronized：外部导入。

	UserType *string `json:"UserType,omitempty" name:"UserType"`
}

type SAMLIdPCertificate struct {

	// 证书ID。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 证书颁发者。

	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`
	// 证书的过期时间。

	NotAfter *string `json:"NotAfter,omitempty" name:"NotAfter"`
	// 证书的创建时间。

	NotBefore *string `json:"NotBefore,omitempty" name:"NotBefore"`
	// PEM 格式的公钥证书（Base64 编码）。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// 证书序列号。

	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
	// 证书的签名算法。

	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" name:"SignatureAlgorithm"`
	// 证书的主体。

	Subject *string `json:"Subject,omitempty" name:"Subject"`
	// 证书版本。

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// PEM 格式的 X509 证书。

	X509Certificate *string `json:"X509Certificate,omitempty" name:"X509Certificate"`
}

type UserProvisioningEvent struct {

	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。

	DeletionStrategy *string `json:"DeletionStrategy,omitempty" name:"DeletionStrategy"`
	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。

	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" name:"DuplicationStrategy"`
	// 执行失败次数。

	ErrorCount *int64 `json:"ErrorCount,omitempty" name:"ErrorCount"`
	// CAM 用户同步事件上次执行失败的错误信息。

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 同步事件ID。

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// CAM 用户同步的身份 ID。取值： 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。

	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`
	// CAM 用户同步的身份名称。取值：当PrincipalType取值为Group时，该值为CIC用户组名称。当PrincipalType取值为User时，该值为CIC用户名称。

	PrincipalName *string `json:"PrincipalName,omitempty" name:"PrincipalName"`
	// CAM 用户同步的身份类型。取值：User：表示该 CAM 用户同步的身份是CIC用户。Group：表示该 CAM 用户同步的身份是CIC用户组。

	PrincipalType *string `json:"PrincipalType,omitempty" name:"PrincipalType"`
	// 源操作类型。取值： StartProvisioning：启用 CAM 用户同步。 DeleteProvisioning：删除 CAM 用户同步。 AddUserToGroup：用户加入用户组。 RemoveUserFromGroup：用户从用户组移出。

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 集团账号目标账号名称。

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 同步的集团账号目标账号类型，ManagerUin管理账号;MemberUin成员账号

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 集团账号目标账号的UIN。

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// CAM 用户同步的状态。取值： Enabled：CAM 用户同步已启用。 Disabled：CAM 用户同步未启用。

	UserProvisioningId *string `json:"UserProvisioningId,omitempty" name:"UserProvisioningId"`
}

type ZoneStatistics struct {

	// 用户组数。

	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`
	// 用户组配额。

	GroupQuota *int64 `json:"GroupQuota,omitempty" name:"GroupQuota"`
	// 权限配置数

	RoleConfigurationCount *int64 `json:"RoleConfigurationCount,omitempty" name:"RoleConfigurationCount"`
	// 权限配置配额。

	RoleConfigurationQuota *int64 `json:"RoleConfigurationQuota,omitempty" name:"RoleConfigurationQuota"`
	// 同步角色数。

	RoleConfigurationSyncCount *int64 `json:"RoleConfigurationSyncCount,omitempty" name:"RoleConfigurationSyncCount"`
	// 权限配置绑定的系统策略配额。

	SystemPolicyPerRoleConfigurationQuota *int64 `json:"SystemPolicyPerRoleConfigurationQuota,omitempty" name:"SystemPolicyPerRoleConfigurationQuota"`
	// 用户数。

	UserCount *int64 `json:"UserCount,omitempty" name:"UserCount"`
	// 同步用户数。

	UserProvisioningCount *int64 `json:"UserProvisioningCount,omitempty" name:"UserProvisioningCount"`
	// 用户配额。

	UserQuota *int64 `json:"UserQuota,omitempty" name:"UserQuota"`
}

type AddExternalSAMLIdPCertificateRequest struct {
	*tchttp.BaseRequest

	// PEM 格式的 X509 证书。 由 SAML 身份提供商提供。

	X509Certificate *string `json:"X509Certificate,omitempty" name:"X509Certificate"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *AddExternalSAMLIdPCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddExternalSAMLIdPCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddExternalSAMLIdPCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SAML 签名证书 ID。

		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddExternalSAMLIdPCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddExternalSAMLIdPCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAddUserToGroupRequest struct {
	*tchttp.BaseRequest

	// 添加的用户组ID和用户ID

	GroupInfo []*GroupIdAndUserId `json:"GroupInfo,omitempty" name:"GroupInfo"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *BatchAddUserToGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddUserToGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchAddUserToGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchAddUserToGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchAddUserToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateUserRequest struct {
	*tchttp.BaseRequest

	// 批量创建的用户信息

	Users []*CreateUserInfo `json:"Users,omitempty" name:"Users"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *BatchCreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchCreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchRemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID和用户ID

	GroupInfo []*GroupIdAndUserId `json:"GroupInfo,omitempty" name:"GroupInfo"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *BatchRemoveUserFromGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchRemoveUserFromGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchRemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchRemoveUserFromGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchRemoveUserFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentityCenterRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIdentityCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdentityCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentityCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// SCIM 同步状态。Enabled：启用。 Disabled：禁用。

		ScimSyncStatus *string `json:"ScimSyncStatus,omitempty" name:"ScimSyncStatus"`
		// 服务开启状态，Disabled代表未开通，Enabled代表已开通

		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。

		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdentityCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdentityCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenSamlResponseRequest struct {
	*tchttp.BaseRequest

	// 登陆态

	LoginToken *string `json:"LoginToken,omitempty" name:"LoginToken"`
	// 登陆类型，UserSAML代表用户SSO，RoleSAML代表角色

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 权限配置ID

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// 目标成员账号uin

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *GenSamlResponseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenSamlResponseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenSamlResponseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SAML响应

		SAMLResponse *string `json:"SAMLResponse,omitempty" name:"SAMLResponse"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenSamlResponseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenSamlResponseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneDetailRequest struct {
	*tchttp.BaseRequest

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *GetZoneDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZoneDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// idp entity ID。

		IdpEntityID *string `json:"IdpEntityID,omitempty" name:"IdpEntityID"`
		// SSO登陆开启时的IDP登陆URL

		IdpLoginUrl *string `json:"IdpLoginUrl,omitempty" name:"IdpLoginUrl"`
		// sp向idp发送的认证请求。

		SAMLRequest *string `json:"SAMLRequest,omitempty" name:"SAMLRequest"`
		// 服务开启状态，Disabled代表未开通，Enabled代表已开通

		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
		// SSO登陆开启状态。Disabled代表未开启，Enabled代表已开启

		SSOStatus *string `json:"SSOStatus,omitempty" name:"SSOStatus"`
		// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetZoneDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZoneDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneStatisticsRequest struct {
	*tchttp.BaseRequest

	// 空间ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *GetZoneStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZoneStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZoneStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 空间的统计信息。

		ZoneStatistics *ZoneStatistics `json:"ZoneStatistics,omitempty" name:"ZoneStatistics"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetZoneStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZoneStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForUserProvisioningRequest struct {
	*tchttp.BaseRequest

	// 登陆token

	LoginToken *string `json:"LoginToken,omitempty" name:"LoginToken"`
	// 每页的最大数据条数。 取值范围：1~20。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *ListAccountsForUserProvisioningRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccountsForUserProvisioningRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForUserProvisioningResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步的成员账号和创建的子账号信息列表

		AccountUserProvisionings []*AccountUserProvisioning `json:"AccountUserProvisionings,omitempty" name:"AccountUserProvisionings"`
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAccountsForUserProvisioningResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccountsForUserProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListExternalSAMLIdPCertificatesRequest struct {
	*tchttp.BaseRequest

	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListExternalSAMLIdPCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListExternalSAMLIdPCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListExternalSAMLIdPCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SAML 签名证书列表

		SAMLIdPCertificates []*SAMLIdPCertificate `json:"SAMLIdPCertificates,omitempty" name:"SAMLIdPCertificates"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListExternalSAMLIdPCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListExternalSAMLIdPCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 用户类型 Manual：手动创建，Synchronized：外部导入。

	UserType *string `json:"UserType,omitempty" name:"UserType"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组的用户列表

		GroupMembers []*GroupMembers `json:"GroupMembers,omitempty" name:"GroupMembers"`
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserSyncProvisioningsRequest struct {
	*tchttp.BaseRequest

	// 检测条件。

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-********）。 当PrincipalType取值为User时，该值为用户 ID（u-********）。

	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`
	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。

	PrincipalType *string `json:"PrincipalType,omitempty" name:"PrincipalType"`
	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 集团账号目标账号的UIN。

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListUserSyncProvisioningsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserSyncProvisioningsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserSyncProvisioningsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// CAM同步的用户列表。

		UserProvisionings []*UserProvisioning `json:"UserProvisionings,omitempty" name:"UserProvisionings"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserSyncProvisioningsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserSyncProvisioningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。 目前仅支持用户名，邮箱，用户userId，描述

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 筛选的用户组，该用户组关联的子用户会返回IsSelected=1

	FilterGroups []*string `json:"FilterGroups,omitempty" name:"FilterGroups"`
	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法经过多次查询，直到IsTruncated为false时，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序的字段，目前只支持CreateTime，默认是CreateTime字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 排序类型：Desc 倒序 Asc 正序，需要您和SortField一起设置

	SortType *string `json:"SortType,omitempty" name:"SortType"`
	// 用户状态 Enabled：启用， Disabled：禁用。

	UserStatus *string `json:"UserStatus,omitempty" name:"UserStatus"`
	// 用户类型 Manual：手动创建，Synchronized：外部导入。

	UserType *string `json:"UserType,omitempty" name:"UserType"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// 用户列表。

		Users []*UserInfo `json:"Users,omitempty" name:"Users"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginCICRequest struct {
	*tchttp.BaseRequest

	// 用户名密码。

	Password *string `json:"Password,omitempty" name:"Password"`
	// samlResponse，base64加密

	SAMLResponse *string `json:"SAMLResponse,omitempty" name:"SAMLResponse"`
	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *LoginCICRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginCICRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginCICResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登陆态

		LoginSkey *string `json:"LoginSkey,omitempty" name:"LoginSkey"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LoginCICResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginCICResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogoutCICRequest struct {
	*tchttp.BaseRequest

	// 登陆态

	LoginSkey *string `json:"LoginSkey,omitempty" name:"LoginSkey"`
}

func (r *LogoutCICRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogoutCICRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogoutCICResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LogoutCICResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogoutCICResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenIdentityCenterRequest struct {
	*tchttp.BaseRequest

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *OpenIdentityCenterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenIdentityCenterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenIdentityCenterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenIdentityCenterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenIdentityCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveExternalSAMLIdPCertificateRequest struct {
	*tchttp.BaseRequest

	// PEM 格式的 X509 证书。 由 SAML 身份提供商提供。

	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *RemoveExternalSAMLIdPCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveExternalSAMLIdPCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveExternalSAMLIdPCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveExternalSAMLIdPCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveExternalSAMLIdPCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSCIMCredentialStatusRequest struct {
	*tchttp.BaseRequest

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。

	CredentialId *string `json:"CredentialId,omitempty" name:"CredentialId"`
	// SCIM密钥状态。Enabled：启用。 Disabled：禁用。

	NewStatus *string `json:"NewStatus,omitempty" name:"NewStatus"`
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *UpdateSCIMCredentialStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSCIMCredentialStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSCIMCredentialStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSCIMCredentialStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSCIMCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateZoneRequest struct {
	*tchttp.BaseRequest

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。

	NewZoneName *string `json:"NewZoneName,omitempty" name:"NewZoneName"`
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *UpdateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLoginSkeyRequest struct {
	*tchttp.BaseRequest

	// 登陆态

	LoginSkey *string `json:"LoginSkey,omitempty" name:"LoginSkey"`
}

func (r *VerifyLoginSkeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLoginSkeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLoginSkeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户名称。空间内必须唯一。不支持修改。 格式：包含数字、英文字母和特殊符号@_-.。 长度：最大 64 个字符

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 空间ID。z-前缀开头，后面是12位随机数字/小写字母

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// 空间名，必须全局唯一。

		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyLoginSkeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLoginSkeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// 临时证书密钥ID

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时证书密钥Key

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 访问token

	Token *string `json:"Token,omitempty" name:"Token"`
}

type CheckRelayStateIsLegalRequest struct {
	*tchttp.BaseRequest

	// 登陆方式，目前只支持tccli，代表CLI登陆。

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆方式所对应合法的回跳域名。

	RelayState *string `json:"RelayState,omitempty" name:"RelayState"`
}

func (r *CheckRelayStateIsLegalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRelayStateIsLegalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRelayStateIsLegalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表不合法，1代表合法

		IsLegal *int64 `json:"IsLegal,omitempty" name:"IsLegal"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckRelayStateIsLegalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRelayStateIsLegalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentityCenterResourceByRoleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIdentityCenterResourceByRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdentityCenterResourceByRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdentityCenterResourceByRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户CIC相关资源

		List []*string `json:"List,omitempty" name:"List"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdentityCenterResourceByRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdentityCenterResourceByRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRoleConfigurationCredentialRequest struct {
	*tchttp.BaseRequest

	// 登陆态

	LoginToken *string `json:"LoginToken,omitempty" name:"LoginToken"`
	// 权限配置ID

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// 目标成员账号uin

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *GenRoleConfigurationCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRoleConfigurationCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRoleConfigurationCredentialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时访问凭证

		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`
		// 临时访问凭证的过期时间，返回 Unix 时间戳，精确到秒

		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenRoleConfigurationCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRoleConfigurationCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSyncProvisioningEventRequest struct {
	*tchttp.BaseRequest

	// CAM 用户同步事件的 ID。

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *GetUserSyncProvisioningEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSyncProvisioningEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserSyncProvisioningEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CAM 用户同步事件信息。

		UserProvisioningEvent *UserProvisioningEvent `json:"UserProvisioningEvent,omitempty" name:"UserProvisioningEvent"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserSyncProvisioningEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserSyncProvisioningEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForAccessAssignmentRequest struct {
	*tchttp.BaseRequest

	// 登陆token

	LoginToken *string `json:"LoginToken,omitempty" name:"LoginToken"`
	// 页码数量，从1开始.

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页的最大数据条数。 取值范围：1~20。默认值为10.

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListAccountsForAccessAssignmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccountsForAccessAssignmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForAccessAssignmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步的成员账号列表

		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAccountsForAccessAssignmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAccountsForAccessAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRoleConfigurationProvisioningsRequest struct {
	*tchttp.BaseRequest

	// Deployed: 部署成功 DeployedRequired：需要重新部署 DeployFailed：部署失败

	DeploymentStatus *string `json:"DeploymentStatus,omitempty" name:"DeploymentStatus"`
	// 支持配置名称搜索。

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 权限配置ID。

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 同步的集团账号目标账号的UIN。

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListRoleConfigurationProvisioningsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationProvisioningsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRoleConfigurationProvisioningsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 部成员账号列表。

		RoleConfigurationProvisionings []*RoleConfigurationProvisionings `json:"RoleConfigurationProvisionings,omitempty" name:"RoleConfigurationProvisionings"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRoleConfigurationProvisioningsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationProvisioningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRoleConfigurationsForAccountRequest struct {
	*tchttp.BaseRequest

	// 登陆token

	LoginToken *string `json:"LoginToken,omitempty" name:"LoginToken"`
	// 每页的最大数据条数。 取值范围：1~20。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 目标uin

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *ListRoleConfigurationsForAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationsForAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRoleConfigurationsForAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 同步的权限配置信息列表

		RoleConfigurationsForAccount []*RoleConfiguration `json:"RoleConfigurationsForAccount,omitempty" name:"RoleConfigurationsForAccount"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRoleConfigurationsForAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRoleConfigurationsForAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTasksRequest struct {
	*tchttp.BaseRequest

	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****）， 当PrincipalType取值为User时，该值为用户 ID（u-****）。

	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`
	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。

	PrincipalType *string `json:"PrincipalType,omitempty" name:"PrincipalType"`
	// 权限配置ID。

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 同步的集团账号目标账号的UIN。

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 任务类型。

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 任务详情

		Tasks []*TaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserSyncProvisioningEventsRequest struct {
	*tchttp.BaseRequest

	// 每页的最大数据条数。 取值范围：1~100。 默认值：10。

	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。 当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 用户同步ID。

	UserProvisioningId *string `json:"UserProvisioningId,omitempty" name:"UserProvisioningId"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ListUserSyncProvisioningEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserSyncProvisioningEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserSyncProvisioningEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果是否被截断。取值： true：已截断。 false：未截断。

		IsTruncated *bool `json:"IsTruncated,omitempty" name:"IsTruncated"`
		// 每页的最大数据条数。

		MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
		// 查询返回结果下一页的令牌。 说明 只有IsTruncated为true时，才显示该参数。

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// 符合请求参数条件的数据总条数。

		TotalCounts *int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// CAM同步的用户事件列表

		UserProvisioningEvents []*UserProvisioningEvent `json:"UserProvisioningEvents,omitempty" name:"UserProvisioningEvents"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserSyncProvisioningEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserSyncProvisioningEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProvisionRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 权限配置ID。

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 集团账号目标账号的UIN。

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 空间 ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ProvisionRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProvisionRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProvisionRoleConfigurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情。

		Task *RoleProvisioningsTask `json:"Task,omitempty" name:"Task"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProvisionRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProvisionRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryUserSyncProvisioningEventRequest struct {
	*tchttp.BaseRequest

	// 冲突策略。KeepBoth:两者都保留;TakeOver:替换

	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" name:"DuplicationStrategy"`
	// CAM 用户同步事件 ID。

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *RetryUserSyncProvisioningEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryUserSyncProvisioningEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryUserSyncProvisioningEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryUserSyncProvisioningEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryUserSyncProvisioningEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTemporarySecretShowStatusRequest struct {
	*tchttp.BaseRequest

	// Enabled： 展示， Disabled： 不展示。

	TemporarySecretShowStatus *string `json:"TemporarySecretShowStatus,omitempty" name:"TemporarySecretShowStatus"`
	// 空间ID。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *SetTemporarySecretShowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTemporarySecretShowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTemporarySecretShowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTemporarySecretShowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTemporarySecretShowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomPolicyForRoleConfigurationRequest struct {
	*tchttp.BaseRequest

	// 权限策略名称，长度最大为 32 个字符。

	CustomPolicyName *string `json:"CustomPolicyName,omitempty" name:"CustomPolicyName"`
	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。

	NewCustomPolicyDocument *string `json:"NewCustomPolicyDocument,omitempty" name:"NewCustomPolicyDocument"`
	// 权限配置 ID

	RoleConfigurationId *string `json:"RoleConfigurationId,omitempty" name:"RoleConfigurationId"`
	// 空间 ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *UpdateCustomPolicyForRoleConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomPolicyForRoleConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomPolicyForRoleConfigurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomPolicyForRoleConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomPolicyForRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

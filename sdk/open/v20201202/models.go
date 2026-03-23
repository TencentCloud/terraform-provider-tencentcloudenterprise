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

package v20201202

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type UpdateSamlConfigRequest struct {
	*tchttp.BaseRequest

	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// idp主键id

	IdpId *uint64 `json:"IdpId,omitempty" name:"IdpId"`
	// 是否同步&nbsp;idp&nbsp;用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// 辅助域名

	AssistDomain *string `json:"AssistDomain,omitempty" name:"AssistDomain"`
	// saml配置

	SamlMetaData *string `json:"SamlMetaData,omitempty" name:"SamlMetaData"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateSamlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSamlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOidcConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOidcConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOidcConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
}

func (r *CreateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AgentId

		AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
		// AppName

		AppName *string `json:"AppName,omitempty" name:"AppName"`
		// CorpId

		CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
		// CreatedTime

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// Id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Secret

		Secret *string `json:"Secret,omitempty" name:"Secret"`
		// UpdateTime

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// Item

		Item *WorkWeixinItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepartmentAttr struct {

	// 部门id

	DepartmentId *uint64 `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 部门名称

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// 父部门id

	ParentDepartmentId *uint64 `json:"ParentDepartmentId,omitempty" name:"ParentDepartmentId"`
	// 部门成员列表

	UserList []*WechatUserAttr `json:"UserList,omitempty" name:"UserList"`
}

type DepartmentList struct {

	// DepartmentId

	DepartmentId *int64 `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// DepartmentName

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// ParentDepartmentId

	ParentDepartmentId *int64 `json:"ParentDepartmentId,omitempty" name:"ParentDepartmentId"`
	// UserList

	UserList []*UserList `json:"UserList,omitempty" name:"UserList"`
}

type LdapGetItem struct {

	// LdapIdpList

	LdapIdpList []*LdapIdpConfig `json:"LdapIdpList,omitempty" name:"LdapIdpList"`
}

type LdapIdpConfigOut struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// IdpName

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// PropertiesMap

	PropertiesMap *PropertiesMap `json:"PropertiesMap,omitempty" name:"PropertiesMap"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// IsEnabled

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// CreatedTime

	CreatedTime *uint64 `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// UpdateTime

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// CreatorUin

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
}

type SamlIdpConfig struct {

	// IdpId

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
	// IdpName

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// IsEnabled

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// CreatorUin

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// CreatedTime

	CreatedTime *uint64 `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// UpdateTime

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// SyncFlag

	SyncFlag *int64 `json:"SyncFlag,omitempty" name:"SyncFlag"`
	// SamlMetaData

	SamlMetaData *string `json:"SamlMetaData,omitempty" name:"SamlMetaData"`
	// AssistDomain

	AssistDomain *string `json:"AssistDomain,omitempty" name:"AssistDomain"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

type EnabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
	// 用户登录名

	TuserId *string `json:"TuserId,omitempty" name:"TuserId"`
}

func (r *EnabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantInnerUserRequest struct {
	*tchttp.BaseRequest

	// GrantUin

	GrantUin *string `json:"GrantUin,omitempty" name:"GrantUin"`
	// GrantOwner

	GrantOwner *string `json:"GrantOwner,omitempty" name:"GrantOwner"`
	// GrantType

	GrantType *string `json:"GrantType,omitempty" name:"GrantType"`
}

func (r *GrantInnerUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantInnerUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Item struct {

	// 配置&nbsp;id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// &nbsp;配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主账号&nbsp;uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 身份提供商类型

	ProviderType *int64 `json:"ProviderType,omitempty" name:"ProviderType"`
	// 企业认证状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// &nbsp;oauth协议映射

	Oauth *string `json:"Oauth,omitempty" name:"Oauth"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// IsGlobal

	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// Cas

	Cas *string `json:"Cas,omitempty" name:"Cas"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// SamlMetaData

	SamlMetaData *string `json:"SamlMetaData,omitempty" name:"SamlMetaData"`
	// AssistDomain

	AssistDomain *string `json:"AssistDomain,omitempty" name:"AssistDomain"`
	// Ldap

	Ldap *string `json:"Ldap,omitempty" name:"Ldap"`
	// Oidc

	Oidc *string `json:"Oidc,omitempty" name:"Oidc"`
}

type TestResult struct {

	// StepNo

	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`
	// Code

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// CodeDesc

	CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
	// Message

	Message *string `json:"Message,omitempty" name:"Message"`
}

type WorkWeixinItem struct {

	// AppName

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// AgentId

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// Secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// CorpId

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// IsGlobal

	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// SyncFlag

	SyncFlag *int64 `json:"SyncFlag,omitempty" name:"SyncFlag"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// UpdateTime

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// CreatedTime

	CreatedTime *int64 `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateWorkWeixinConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *WorkWeixinItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkWeixinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否同步&nbsp;idp&nbsp;用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *UpdateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LdapIdpConfig struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// CreatedTime

	CreatedTime *uint64 `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// UpdateTime

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// AuthType

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// IdpName

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// LdapType

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// LdapUrl

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// BaseDn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// Cert

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// ConnectType

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// AdminAccount

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// AdminPassword

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// FilterCondition

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// UserAccountField

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// UserMailField

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// UserNicknameField

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// UserPhoneField

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// IsEnabled

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

type SamlItem struct {

	// IdpConfig

	IdpConfig *SamlIdpConfig `json:"IdpConfig,omitempty" name:"IdpConfig"`
}

type GetWorkWeixinConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *LdapItemOut `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *TestLdapItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestLdapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOidcConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOidcConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOidcConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
}

func (r *DisabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLdapIdpConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLdapIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLdapIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *LdapGetItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLdapIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内部应用id

		AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
		// 内部应用名称

		AppName *string `json:"AppName,omitempty" name:"AppName"`
		// 企业id

		CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
		// 关联时间

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 应用存储id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 内部应用秘钥

		Secret *string `json:"Secret,omitempty" name:"Secret"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// Item

		Item *WorkWeixinItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinConfigRequest struct {
	*tchttp.BaseRequest

	// AppName

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// AgentId

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// CorpId

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// Secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// SyncFlag

	SyncFlag *uint64 `json:"SyncFlag,omitempty" name:"SyncFlag"`
}

func (r *CreateWorkWeixinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchBindWorkWeixinAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchBindWorkWeixinAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchBindWorkWeixinAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCICUserSAMLConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCICUserSAMLConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCICUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserList struct {

	// WechatUserId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
	// WechatUserName

	WechatUserName *string `json:"WechatUserName,omitempty" name:"WechatUserName"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type AddSamlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *SamlItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSamlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSamlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *ProviderData `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSamlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *SamlItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSamlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSamlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinConfigRequest struct {
	*tchttp.BaseRequest

	// AppName

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// AgentId

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// Secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// CorpId

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// SyncFlag

	SyncFlag *uint64 `json:"SyncFlag,omitempty" name:"SyncFlag"`
}

func (r *UpdateWorkWeixinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkWeixinAuthorizationItem struct {

	// DepartmentList

	DepartmentList []*DepartmentList `json:"DepartmentList,omitempty" name:"DepartmentList"`
	// VisibleUserList

	VisibleUserList []*UserList `json:"VisibleUserList,omitempty" name:"VisibleUserList"`
}

type GetWorkWeixinOpenAppMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 企业微信部门成员信息

		DepartmentList []*DepartmentAttr `json:"DepartmentList,omitempty" name:"DepartmentList"`
		// 企业微信可见成员信息

		VisibleUserList []*WechatUserAttr `json:"VisibleUserList,omitempty" name:"VisibleUserList"`
		// Items

		Items *AppMemberItem `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkWeixinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOidcConfigRequest struct {
	*tchttp.BaseRequest

	// IdpName

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// IdentityUrl

	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`
	// IdentityKey

	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`
	// ClientId

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// AuthorizationEndpoint

	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`
	// Scope

	Scope []*string `json:"Scope,omitempty" name:"Scope"`
	// ResponseType

	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`
	// ResponseMode

	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// EmailField

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// NickNameField

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// PhoneNumField

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// LoginAccountField

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// CountryCodeField

	CountryCodeField *string `json:"CountryCodeField,omitempty" name:"CountryCodeField"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// LogoutUrl

	LogoutUrl *string `json:"LogoutUrl,omitempty" name:"LogoutUrl"`
}

func (r *UpdateOidcConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOidcConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LdapItemOut struct {

	// IdpConfig

	IdpConfig *LdapIdpConfigOut `json:"IdpConfig,omitempty" name:"IdpConfig"`
}

type CreateCICUserSAMLConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SAMLProviderArn

		SAMLProviderArn *string `json:"SAMLProviderArn,omitempty" name:"SAMLProviderArn"`
		// Name

		Name *string `json:"Name,omitempty" name:"Name"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCICUserSAMLConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCICUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 是否同步&nbsp;idp&nbsp;用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *CreateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderRequest struct {
	*tchttp.BaseRequest
}

func (r *ListIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderList struct {

	// 身份提供商列表

	List []*Item `json:"List,omitempty" name:"List"`
}

type TestLdapItem struct {

	// TestResult

	TestResult []*TestResult `json:"TestResult,omitempty" name:"TestResult"`
}

type GetWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppMemberRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProviderItem struct {

	// &nbsp;身份提供商配置

	Item *ProviderData `json:"Item,omitempty" name:"Item"`
}

type BindWorkWeixinAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindWorkWeixinAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCICUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCICUserSAMLConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCICUserSAMLConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCICUserSAMLConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SAMLMetadata

		SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`
		// Status

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCICUserSAMLConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCICUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinAuthorizationScopeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Items

		Items *WorkWeixinAuthorizationItem `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinAuthorizationScopeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinAuthorizationScopeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCICUserSAMLConfigRequest struct {
	*tchttp.BaseRequest

	// Uin

	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`
	// Operate

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// SAMLMetadataDocument

	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *UpdateCICUserSAMLConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCICUserSAMLConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LdapItem struct {

	// IdpConf

	IdpConf *LdapIdpConfig `json:"IdpConf,omitempty" name:"IdpConf"`
}

type WechatUserAttr struct {

	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
	// 企业微信用户名

	WechatUserName *string `json:"WechatUserName,omitempty" name:"WechatUserName"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type GetWorkWeixinAuthorizationScopeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinAuthorizationScopeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinAuthorizationScopeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSamlConfigRequest struct {
	*tchttp.BaseRequest

	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议名称

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否同步&nbsp;idp&nbsp;账号

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// Saml配置

	SamlMetaData *string `json:"SamlMetaData,omitempty" name:"SamlMetaData"`
	// 辅助域名

	AssistDomain *string `json:"AssistDomain,omitempty" name:"AssistDomain"`
}

func (r *AddSamlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSamlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchBindWorkWeixinAccountRequest struct {
	*tchttp.BaseRequest

	// Items

	Items []*Items `json:"Items,omitempty" name:"Items"`
}

func (r *BatchBindWorkWeixinAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchBindWorkWeixinAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *LdapItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 配置id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOidcConfigRequest struct {
	*tchttp.BaseRequest

	// IdpName

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// IdentityUrl

	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`
	// IdentityKey

	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`
	// ClientId

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// AuthorizationEndpoint

	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`
	// Scope

	Scope []*string `json:"Scope,omitempty" name:"Scope"`
	// ResponseType

	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`
	// ResponseMode

	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// IsSyncIdpUser

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// EmailField

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// NickNameField

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// PhoneNumField

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// LoginAccountField

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// CountryCodeField

	CountryCodeField *string `json:"CountryCodeField,omitempty" name:"CountryCodeField"`
	// LogoutUrl

	LogoutUrl *string `json:"LogoutUrl,omitempty" name:"LogoutUrl"`
}

func (r *AddOidcConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOidcConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindWorkWeixinAccountRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
}

func (r *BindWorkWeixinAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 提供商名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// 地址

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// basedn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 测试账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 测试账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *TestLdapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppMemberItem struct {

	// DepartmentList

	DepartmentList []*DepartmentAttr `json:"DepartmentList,omitempty" name:"DepartmentList"`
	// VisibleUserList

	VisibleUserList []*WechatUserAttr `json:"VisibleUserList,omitempty" name:"VisibleUserList"`
}

type Items struct {

	// Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// OwnerUin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// WechatUserId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
}

type PropertiesMap struct {

	// Cert

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// BaseDn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// LdapUrl

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// LdapType

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// ConnectType

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// AdminAccount

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// AdminPassword

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// UserMailField

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// UserPhoneField

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// FilterCondition

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// UserAccountField

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// UserNicknameField

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
}

type ProviderData struct {

	// 身份提供商数据

	Data *ProviderList `json:"Data,omitempty" name:"Data"`
}

type CreateCICUserSAMLConfigRequest struct {
	*tchttp.BaseRequest

	// SAMLMetadataDocument

	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *CreateCICUserSAMLConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCICUserSAMLConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *WorkWeixinItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrantInnerUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UserAccessToken

		UserAccessToken *string `json:"UserAccessToken,omitempty" name:"UserAccessToken"`
		// ExpiresAt

		ExpiresAt *uint64 `json:"ExpiresAt,omitempty" name:"ExpiresAt"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantInnerUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GrantInnerUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Item

		Item *LdapItem `json:"Item,omitempty" name:"Item"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// SAMLMetadataDocument

	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSAMLProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSAMLProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetSAMLProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSAMLProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Name

		Name *string `json:"Name,omitempty" name:"Name"`
		// Description

		Description *string `json:"Description,omitempty" name:"Description"`
		// CreateTime

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// ModifyTime

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// SAMLMetadata

		SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSAMLProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSAMLProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteSAMLProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSAMLProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSAMLProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

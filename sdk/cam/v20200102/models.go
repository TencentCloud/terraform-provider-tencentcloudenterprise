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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetUidByUinArrRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *GetUidByUinArrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinListByRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin列表

		UinList []*UinList `json:"UinList,omitempty" name:"UinList"`
		// 组列表

		GroupList []*GroupList `json:"GroupList,omitempty" name:"GroupList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUinListByRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinListByRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssumeRole struct {

	// 主账号UIN

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色ID

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 角色状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type BatchCheckUserHavePoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查成功结果

		Success []*BatchCheckResult `json:"Success,omitempty" name:"Success"`
		// 检查失败结果

		Failure []*BatchCheckResult `json:"Failure,omitempty" name:"Failure"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchCheckUserHavePoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCheckUserHavePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCheckUserHavePoliciesRequest struct {
	*tchttp.BaseRequest

	// 查询策略

	Policies []*PolicyToCheck `json:"Policies,omitempty" name:"Policies"`
	// 查询子账号uin

	UserUins []*int64 `json:"UserUins,omitempty" name:"UserUins"`
}

func (r *BatchCheckUserHavePoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCheckUserHavePoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUidByUinArrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin对应的Uid信息

		UinInfo []*UinInfo `json:"UinInfo,omitempty" name:"UinInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUidByUinArrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupList struct {

	// 组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GetRoleListByOwnerUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRoleListByOwnerUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleListByOwnerUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRelatedRolePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRelatedRolePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRelatedRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRelatedRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 需关联的uin列表

	TargetUin []*uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 需关联的GroupId列表

	TargetGroupId []*uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
}

func (r *UpdateRelatedRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRelatedRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyCheckResult struct {

	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否拥有该策略

	OK *bool `json:"OK,omitempty" name:"OK"`
}

type LoginData struct {

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 登陆密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type GetRoleListByOwnerUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色列表

		RoleList []*AssumeRole `json:"RoleList,omitempty" name:"RoleList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleListByOwnerUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRoleListByOwnerUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyToCheck struct {

	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 策略关联项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type VerifyMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UinInfo struct {

	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户uid

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type GroupUserInfo struct {

	// 接收者id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 手机号标识

	PhoneFlag *string `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱标识

	EmailFlag *string `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否是主账户

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 账户类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
}

type VerifyMenuRequest struct {
	*tchttp.BaseRequest

	// 菜单

	Menu *string `json:"Menu,omitempty" name:"Menu"`
	// 子菜单

	Submenu *string `json:"Submenu,omitempty" name:"Submenu"`
}

func (r *VerifyMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductMenuItem struct {

	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 菜单uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type CheckLoginDetail struct {

	// 登陆类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆到到地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
	// 是否使用多因子认证

	MFAUsed *string `json:"MFAUsed,omitempty" name:"MFAUsed"`
}

type LoginDetail struct {

	// 登陆账户类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆的平台地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
}

type BatchCheckResult struct {

	// 批量查询结果

	PolicyCheckResults []*PolicyCheckResult `json:"PolicyCheckResults,omitempty" name:"PolicyCheckResults"`
	// 查询子账号Uin

	UserUin *int64 `json:"UserUin,omitempty" name:"UserUin"`
	// Error

	Error *string `json:"Error,omitempty" name:"Error"`
}

type UinList struct {

	// 账号UIN

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type GetUinListByRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色所属OwnerUin

	RoleOwnerUin *string `json:"RoleOwnerUin,omitempty" name:"RoleOwnerUin"`
}

func (r *GetUinListByRolePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinListByRolePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

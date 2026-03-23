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

package v20190325

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeRelatedUinSessionKeyRequest struct {
	*tchttp.BaseRequest

	// 当前账号登录态

	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`
}

func (r *DescribeRelatedUinSessionKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelatedUinSessionKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// ownerUin

		Owneruin *uint64 `json:"Owneruin,omitempty" name:"Owneruin"`
		// 账户名

		Account *string `json:"Account,omitempty" name:"Account"`
		// 用户类型

		Usertype *string `json:"Usertype,omitempty" name:"Usertype"`
		// 类型

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// 地区

		Area *string `json:"Area,omitempty" name:"Area"`
		// 地区

		Curarea *int64 `json:"Curarea,omitempty" name:"Curarea"`
		// 邮箱已验证

		Mailverify *int64 `json:"Mailverify,omitempty" name:"Mailverify"`
		// 身份证号

		Idcard *string `json:"Idcard,omitempty" name:"Idcard"`
		// 身份证类型

		Idcardtype *int64 `json:"Idcardtype,omitempty" name:"Idcardtype"`
		// 身份证类型

		Curidcardtype *int64 `json:"Curidcardtype,omitempty" name:"Curidcardtype"`
		// url

		Idcardurl *string `json:"Idcardurl,omitempty" name:"Idcardurl"`
		// 身份证uuid

		Idcarduuid *string `json:"Idcarduuid,omitempty" name:"Idcarduuid"`
		// Idcardcacheurl

		Idcardcacheurl *string `json:"Idcardcacheurl,omitempty" name:"Idcardcacheurl"`
		// Idcardcacheuuid

		Idcardcacheuuid *string `json:"Idcardcacheuuid,omitempty" name:"Idcardcacheuuid"`
		// Idcardurlnew

		Idcardurlnew *string `json:"Idcardurlnew,omitempty" name:"Idcardurlnew"`
		// Idcardcacheuuidnew

		Idcardcacheuuidnew *string `json:"Idcardcacheuuidnew,omitempty" name:"Idcardcacheuuidnew"`
		// Entrycardcacheurlnew

		Entrycardcacheurlnew *string `json:"Entrycardcacheurlnew,omitempty" name:"Entrycardcacheurlnew"`
		// Entrycardurlnew

		Entrycardurlnew *string `json:"Entrycardurlnew,omitempty" name:"Entrycardurlnew"`
		// Contact

		Contact *string `json:"Contact,omitempty" name:"Contact"`
		// 电话

		Tel *string `json:"Tel,omitempty" name:"Tel"`
		// 区号

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 是否需要信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 邮箱

		Curmail *string `json:"Curmail,omitempty" name:"Curmail"`
		// 邮箱已验证

		Mailpass *int64 `json:"Mailpass,omitempty" name:"Mailpass"`
		// Curmailpass

		Curmailpass *int64 `json:"Curmailpass,omitempty" name:"Curmailpass"`
		// Registerstatus

		Registerstatus *string `json:"Registerstatus,omitempty" name:"Registerstatus"`
		// Curdevcheckstate

		Curdevcheckstate *int64 `json:"Curdevcheckstate,omitempty" name:"Curdevcheckstate"`
		// Devcheckpass

		Devcheckpass *int64 `json:"Devcheckpass,omitempty" name:"Devcheckpass"`
		// 银行账户

		Bankaccount *string `json:"Bankaccount,omitempty" name:"Bankaccount"`
		// 银行卡号

		Banknumber *string `json:"Banknumber,omitempty" name:"Banknumber"`
		// Perfection

		Perfection *int64 `json:"Perfection,omitempty" name:"Perfection"`
		// Collname

		Collname *string `json:"Collname,omitempty" name:"Collname"`
		// Colltel

		Colltel *string `json:"Colltel,omitempty" name:"Colltel"`
		// Collmail

		Collmail *string `json:"Collmail,omitempty" name:"Collmail"`
		// WanIpTime

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// WanRestrict

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// Accredit

		Accredit *string `json:"Accredit,omitempty" name:"Accredit"`
		// Curaccredit

		Curaccredit *string `json:"Curaccredit,omitempty" name:"Curaccredit"`
		// 地址

		Addr *string `json:"Addr,omitempty" name:"Addr"`
		// Addrprefix

		Addrprefix *string `json:"Addrprefix,omitempty" name:"Addrprefix"`
		// Devcheckmsg

		Devcheckmsg *string `json:"Devcheckmsg,omitempty" name:"Devcheckmsg"`
		// Curdevcheckmsg

		Curdevcheckmsg *string `json:"Curdevcheckmsg,omitempty" name:"Curdevcheckmsg"`
		// Organizationcode

		Organizationcode *string `json:"Organizationcode,omitempty" name:"Organizationcode"`
		// Curorganizationcode

		Curorganizationcode *string `json:"Curorganizationcode,omitempty" name:"Curorganizationcode"`
		// AuthenticateType

		AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`
		// CurauthenticateType

		CurauthenticateType *string `json:"CurauthenticateType,omitempty" name:"CurauthenticateType"`
		// IsModify

		IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
		// Curname

		Curname *string `json:"Curname,omitempty" name:"Curname"`
		// Curidcard

		Curidcard *string `json:"Curidcard,omitempty" name:"Curidcard"`
		// Curtype

		Curtype *int64 `json:"Curtype,omitempty" name:"Curtype"`
		// 平台

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// Devchecktime

		Devchecktime *string `json:"Devchecktime,omitempty" name:"Devchecktime"`
		// Mailchecktime

		Mailchecktime *string `json:"Mailchecktime,omitempty" name:"Mailchecktime"`
		// Firstdevchecktime

		Firstdevchecktime *string `json:"Firstdevchecktime,omitempty" name:"Firstdevchecktime"`
		// Acceptopenprotocol

		Acceptopenprotocol *int64 `json:"Acceptopenprotocol,omitempty" name:"Acceptopenprotocol"`
		// Acceptyunprotocol

		Acceptyunprotocol *int64 `json:"Acceptyunprotocol,omitempty" name:"Acceptyunprotocol"`
		// Openuser

		Openuser *int64 `json:"Openuser,omitempty" name:"Openuser"`
		// Acceptopen

		Acceptopen *string `json:"Acceptopen,omitempty" name:"Acceptopen"`
		// BankStatus

		BankStatus *string `json:"BankStatus,omitempty" name:"BankStatus"`
		// BankMsg

		BankMsg *string `json:"BankMsg,omitempty" name:"BankMsg"`
		// DeployName

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// SubExist

		SubExist *int64 `json:"SubExist,omitempty" name:"SubExist"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 授权方法

		Authmethod *int64 `json:"Authmethod,omitempty" name:"Authmethod"`
		// Isreceiverowner

		Isreceiverowner *int64 `json:"Isreceiverowner,omitempty" name:"Isreceiverowner"`
		// Contactemail

		Contactemail *string `json:"Contactemail,omitempty" name:"Contactemail"`
		// Addtimestamp

		Addtimestamp *string `json:"Addtimestamp,omitempty" name:"Addtimestamp"`
		// Checkstatus

		Checkstatus *int64 `json:"Checkstatus,omitempty" name:"Checkstatus"`
		// Srcplatform

		Srcplatform *string `json:"Srcplatform,omitempty" name:"Srcplatform"`
		// Updatetimestamp

		Updatetimestamp *string `json:"Updatetimestamp,omitempty" name:"Updatetimestamp"`
		// Bindstatus

		Bindstatus *int64 `json:"Bindstatus,omitempty" name:"Bindstatus"`
		// Needresetpassword

		Needresetpassword *int64 `json:"Needresetpassword,omitempty" name:"Needresetpassword"`
		// Qywxuserid

		Qywxuserid *string `json:"Qywxuserid,omitempty" name:"Qywxuserid"`
		// Typebind

		Typebind *int64 `json:"Typebind,omitempty" name:"Typebind"`
		// Username

		Username *string `json:"Username,omitempty" name:"Username"`
		// Mailstatus

		Mailstatus *int64 `json:"Mailstatus,omitempty" name:"Mailstatus"`
		// Identifytype

		Identifytype *int64 `json:"Identifytype,omitempty" name:"Identifytype"`
		// Updatetime

		Updatetime *string `json:"Updatetime,omitempty" name:"Updatetime"`
		// Contactemailflag

		Contactemailflag *int64 `json:"Contactemailflag,omitempty" name:"Contactemailflag"`
		// Phoneflag

		Phoneflag *int64 `json:"Phoneflag,omitempty" name:"Phoneflag"`
		// Remark

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// Istestuser

		Istestuser *int64 `json:"Istestuser,omitempty" name:"Istestuser"`
		// Isowner

		Isowner *int64 `json:"Isowner,omitempty" name:"Isowner"`
		// Uid

		Uid *int64 `json:"Uid,omitempty" name:"Uid"`
		// Wxzsstatus

		Wxzsstatus *int64 `json:"Wxzsstatus,omitempty" name:"Wxzsstatus"`
		// Createuin

		Createuin *uint64 `json:"Createuin,omitempty" name:"Createuin"`
		// Defaultowner

		Defaultowner *uint64 `json:"Defaultowner,omitempty" name:"Defaultowner"`
		// Isauthenticate

		Isauthenticate *int64 `json:"Isauthenticate,omitempty" name:"Isauthenticate"`
		// Countrycode

		Countrycode *string `json:"Countrycode,omitempty" name:"Countrycode"`
		// Isregaccount

		Isregaccount *int64 `json:"Isregaccount,omitempty" name:"Isregaccount"`
		// Phonenum

		Phonenum *string `json:"Phonenum,omitempty" name:"Phonenum"`
		// Accounttype

		Accounttype *int64 `json:"Accounttype,omitempty" name:"Accounttype"`
		// Consolelogin

		Consolelogin *int64 `json:"Consolelogin,omitempty" name:"Consolelogin"`
		// Createtime

		Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
		// Isregsucc

		Isregsucc *int64 `json:"Isregsucc,omitempty" name:"Isregsucc"`
		// Loginaccount

		Loginaccount *string `json:"Loginaccount,omitempty" name:"Loginaccount"`
		// Loginstatus

		Loginstatus *int64 `json:"Loginstatus,omitempty" name:"Loginstatus"`
		// Wxzsopenid

		Wxzsopenid *string `json:"Wxzsopenid,omitempty" name:"Wxzsopenid"`
		// Curinfo

		Curinfo *AccountCurInfo `json:"Curinfo,omitempty" name:"Curinfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserInfoByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendLoginVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SendLoginVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLoginVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginMfa struct {

	// 是否需要认证

	NeedAuth *int64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
	// 认证类型

	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
	// 认证手机号

	AuthPhone *string `json:"AuthPhone,omitempty" name:"AuthPhone"`
}

type GetUserAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// 国家名称

		CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedLoginTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生成的token字符串

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeedLoginTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *EnableSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAutoLockPeriodRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAutoLockPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAutoLockPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 语言类型，zh或en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *ChangeMailPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeMailPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaxSubAccountAndGroupNumRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMaxSubAccountAndGroupNumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaxSubAccountAndGroupNumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNicknameRequest struct {
	*tchttp.BaseRequest

	// 临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *GetNicknameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNicknameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子账号列表

		SubAccounts []*SubAccountUser `json:"SubAccounts,omitempty" name:"SubAccounts"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserByAttributeValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的用户总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息列表

		UserInfo []*AttributeUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserByAttributeValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByAttributeValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {

	// id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
}

type AccountBankInfo struct {

	// 账号id

	Accountid *string `json:"Accountid,omitempty" name:"Accountid"`
	// 账号名称

	Accountname *string `json:"Accountname,omitempty" name:"Accountname"`
	// 账号所属银行

	Accountbank *string `json:"Accountbank,omitempty" name:"Accountbank"`
	// 省

	Provincename *string `json:"Provincename,omitempty" name:"Provincename"`
	// 省ID

	Provinceid *string `json:"Provinceid,omitempty" name:"Provinceid"`
	// 城市

	Cityname *string `json:"Cityname,omitempty" name:"Cityname"`
	// 城市ID

	Cityid *string `json:"Cityid,omitempty" name:"Cityid"`
	// 银行名称

	Bankname *string `json:"Bankname,omitempty" name:"Bankname"`
	// 银行ID

	Bankid *string `json:"Bankid,omitempty" name:"Bankid"`
}

type SubAccounts struct {

	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 秘钥Id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 秘钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type ChangeSubAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeSubAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeSubAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaskedUserInfoByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 主账号UIN

		Owneruin *int64 `json:"Owneruin,omitempty" name:"Owneruin"`
		// 登陆账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 账号类型名称

		Usertype *string `json:"Usertype,omitempty" name:"Usertype"`
		// 账号类型

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// 地区

		Area *string `json:"Area,omitempty" name:"Area"`
		// 当前区域

		Curarea *int64 `json:"Curarea,omitempty" name:"Curarea"`
		// 邮件验证状态

		Mailverify *int64 `json:"Mailverify,omitempty" name:"Mailverify"`
		// 证件名称

		Idcard *string `json:"Idcard,omitempty" name:"Idcard"`
		// 证件类型

		Idcardtype *int64 `json:"Idcardtype,omitempty" name:"Idcardtype"`
		// 当前证件类型

		Curidcardtype *int64 `json:"Curidcardtype,omitempty" name:"Curidcardtype"`
		// 证件url

		Idcardurl *string `json:"Idcardurl,omitempty" name:"Idcardurl"`
		// 证件uuid

		Idcarduuid *string `json:"Idcarduuid,omitempty" name:"Idcarduuid"`
		// 证件缓存地址

		Idcardcacheurl *string `json:"Idcardcacheurl,omitempty" name:"Idcardcacheurl"`
		// 证件缓存uuid

		Idcardcacheuuid *string `json:"Idcardcacheuuid,omitempty" name:"Idcardcacheuuid"`
		// 证件新url

		Idcardurlnew *string `json:"Idcardurlnew,omitempty" name:"Idcardurlnew"`
		// 证件缓存新uuid

		Idcardcacheuuidnew *string `json:"Idcardcacheuuidnew,omitempty" name:"Idcardcacheuuidnew"`
		// 证件缓存新url

		Entrycardcacheurlnew *string `json:"Entrycardcacheurlnew,omitempty" name:"Entrycardcacheurlnew"`
		// 入境卡新url

		Entrycardurlnew *string `json:"Entrycardurlnew,omitempty" name:"Entrycardurlnew"`
		// 联系信息

		Contact *string `json:"Contact,omitempty" name:"Contact"`
		// 电话

		Tel *string `json:"Tel,omitempty" name:"Tel"`
		// 国家编码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 需求信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 当前邮箱

		Curmail *string `json:"Curmail,omitempty" name:"Curmail"`
		// 邮箱验证状态

		Mailpass *int64 `json:"Mailpass,omitempty" name:"Mailpass"`
		// 当前邮箱验证状态

		Curmailpass *int64 `json:"Curmailpass,omitempty" name:"Curmailpass"`
		// 注册状态

		Registerstatus *string `json:"Registerstatus,omitempty" name:"Registerstatus"`
		// 当前开发检查状态

		Curdevcheckstate *int64 `json:"Curdevcheckstate,omitempty" name:"Curdevcheckstate"`
		// 开发检查状态

		Devcheckpass *int64 `json:"Devcheckpass,omitempty" name:"Devcheckpass"`
		// 银行账号

		Bankaccount *string `json:"Bankaccount,omitempty" name:"Bankaccount"`
		// 银行号码

		Banknumber *string `json:"Banknumber,omitempty" name:"Banknumber"`
		// 完成进度

		Perfection *int64 `json:"Perfection,omitempty" name:"Perfection"`
		// 协作者名称

		Collname *string `json:"Collname,omitempty" name:"Collname"`
		// 协作者手机号码

		Colltel *string `json:"Colltel,omitempty" name:"Colltel"`
		// 协作者邮箱

		Collmail *string `json:"Collmail,omitempty" name:"Collmail"`
		// wan&nbsp;ip时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// wan限制

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 授权

		Accredit *string `json:"Accredit,omitempty" name:"Accredit"`
		// 当前授权

		Curaccredit *string `json:"Curaccredit,omitempty" name:"Curaccredit"`
		// 地址

		Addr *string `json:"Addr,omitempty" name:"Addr"`
		// 地址前缀

		Addrprefix *string `json:"Addrprefix,omitempty" name:"Addrprefix"`
		// 开发检查信息

		Devcheckmsg *string `json:"Devcheckmsg,omitempty" name:"Devcheckmsg"`
		// 当前开发检查信息

		Curdevcheckmsg *string `json:"Curdevcheckmsg,omitempty" name:"Curdevcheckmsg"`
		// 组织代码

		Organizationcode *string `json:"Organizationcode,omitempty" name:"Organizationcode"`
		// 当前组织代码

		Curorganizationcode *string `json:"Curorganizationcode,omitempty" name:"Curorganizationcode"`
		// 认证类型

		AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`
		// 当前认证类型

		CurauthenticateType *string `json:"CurauthenticateType,omitempty" name:"CurauthenticateType"`
		// 是否修改

		IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
		// 当前名称

		Curname *string `json:"Curname,omitempty" name:"Curname"`
		// 当前身份证件

		Curidcard *string `json:"Curidcard,omitempty" name:"Curidcard"`
		// 当前类型

		Curtype *int64 `json:"Curtype,omitempty" name:"Curtype"`
		// 账号来源

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 开发检查时间

		Devchecktime *string `json:"Devchecktime,omitempty" name:"Devchecktime"`
		// 邮件检查时间

		Mailchecktime *string `json:"Mailchecktime,omitempty" name:"Mailchecktime"`
		// 开发首次检查时间

		Firstdevchecktime *string `json:"Firstdevchecktime,omitempty" name:"Firstdevchecktime"`
		// 接受开放协议

		Acceptopenprotocol *int64 `json:"Acceptopenprotocol,omitempty" name:"Acceptopenprotocol"`
		// 接受云协议

		Acceptyunprotocol *int64 `json:"Acceptyunprotocol,omitempty" name:"Acceptyunprotocol"`
		// 开放平台用户

		Openuser *int64 `json:"Openuser,omitempty" name:"Openuser"`
		// 接受开放平台

		Acceptopen *string `json:"Acceptopen,omitempty" name:"Acceptopen"`
		// 通过审核银行账号信息

		PassedBankInfo *AccountBankInfo `json:"PassedBankInfo,omitempty" name:"PassedBankInfo"`
		// 待审核银行账号信息

		NeedCheckBankInfo *AccountBankInfo `json:"NeedCheckBankInfo,omitempty" name:"NeedCheckBankInfo"`
		// 银行账号状态

		BankStatus *string `json:"BankStatus,omitempty" name:"BankStatus"`
		// 银行账号审核信息

		BankMsg *string `json:"BankMsg,omitempty" name:"BankMsg"`
		// 部署名称

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 存在子账号

		SubExist *int64 `json:"SubExist,omitempty" name:"SubExist"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 账号名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 消息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 账号当前身份信息

		Curinfo *AccountCurInfo `json:"Curinfo,omitempty" name:"Curinfo"`
		// 认证方式

		Authmethod *int64 `json:"Authmethod,omitempty" name:"Authmethod"`
		// Isreceiverowner

		Isreceiverowner *int64 `json:"Isreceiverowner,omitempty" name:"Isreceiverowner"`
		// Contactemail

		Contactemail *string `json:"Contactemail,omitempty" name:"Contactemail"`
		// Addtimestamp

		Addtimestamp *string `json:"Addtimestamp,omitempty" name:"Addtimestamp"`
		// Checkstatus

		Checkstatus *int64 `json:"Checkstatus,omitempty" name:"Checkstatus"`
		// Srcplatform

		Srcplatform *string `json:"Srcplatform,omitempty" name:"Srcplatform"`
		// Updatetimestamp

		Updatetimestamp *string `json:"Updatetimestamp,omitempty" name:"Updatetimestamp"`
		// Bindstatus

		Bindstatus *int64 `json:"Bindstatus,omitempty" name:"Bindstatus"`
		// Needresetpassword

		Needresetpassword *int64 `json:"Needresetpassword,omitempty" name:"Needresetpassword"`
		// Qywxuserid

		Qywxuserid *string `json:"Qywxuserid,omitempty" name:"Qywxuserid"`
		// Typebind

		Typebind *int64 `json:"Typebind,omitempty" name:"Typebind"`
		// Username

		Username *string `json:"Username,omitempty" name:"Username"`
		// Mailstatus

		Mailstatus *int64 `json:"Mailstatus,omitempty" name:"Mailstatus"`
		// Identifytype

		Identifytype *int64 `json:"Identifytype,omitempty" name:"Identifytype"`
		// Updatetime

		Updatetime *string `json:"Updatetime,omitempty" name:"Updatetime"`
		// Contactemailflag

		Contactemailflag *int64 `json:"Contactemailflag,omitempty" name:"Contactemailflag"`
		// Phoneflag

		Phoneflag *int64 `json:"Phoneflag,omitempty" name:"Phoneflag"`
		// Remark

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// Istestuser

		Istestuser *int64 `json:"Istestuser,omitempty" name:"Istestuser"`
		// Isowner

		Isowner *int64 `json:"Isowner,omitempty" name:"Isowner"`
		// Uid

		Uid *int64 `json:"Uid,omitempty" name:"Uid"`
		// Wxzsstatus

		Wxzsstatus *int64 `json:"Wxzsstatus,omitempty" name:"Wxzsstatus"`
		// Createuin

		Createuin *uint64 `json:"Createuin,omitempty" name:"Createuin"`
		// Defaultowner

		Defaultowner *uint64 `json:"Defaultowner,omitempty" name:"Defaultowner"`
		// Isauthenticate

		Isauthenticate *int64 `json:"Isauthenticate,omitempty" name:"Isauthenticate"`
		// Countrycode

		Countrycode *string `json:"Countrycode,omitempty" name:"Countrycode"`
		// Isregaccount

		Isregaccount *int64 `json:"Isregaccount,omitempty" name:"Isregaccount"`
		// Phonenum

		Phonenum *string `json:"Phonenum,omitempty" name:"Phonenum"`
		// Accounttype

		Accounttype *int64 `json:"Accounttype,omitempty" name:"Accounttype"`
		// Consolelogin

		Consolelogin *int64 `json:"Consolelogin,omitempty" name:"Consolelogin"`
		// Createtime

		Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
		// Isregsucc

		Isregsucc *int64 `json:"Isregsucc,omitempty" name:"Isregsucc"`
		// Loginaccount

		Loginaccount *string `json:"Loginaccount,omitempty" name:"Loginaccount"`
		// Loginstatus

		Loginstatus *int64 `json:"Loginstatus,omitempty" name:"Loginstatus"`
		// Wxzsopenid

		Wxzsopenid *string `json:"Wxzsopenid,omitempty" name:"Wxzsopenid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaskedUserInfoByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaskedUserInfoByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSafeAuthFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSafeAuthFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSafeAuthFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultIFactor struct {

	// MultiFactorChoices

	MultiFactorChoices *int64 `json:"MultiFactorChoices,omitempty" name:"MultiFactorChoices"`
	// MultiFactorcChangable

	MultiFactorcChangable *int64 `json:"MultiFactorcChangable,omitempty" name:"MultiFactorcChangable"`
	// MultiActionChoices

	MultiActionChoices *int64 `json:"MultiActionChoices,omitempty" name:"MultiActionChoices"`
	// MultiActionChangable

	MultiActionChangable *int64 `json:"MultiActionChangable,omitempty" name:"MultiActionChangable"`
	// MfaHardTokenSource

	MfaHardTokenSource *string `json:"MfaHardTokenSource,omitempty" name:"MfaHardTokenSource"`
}

type GetTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的token

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBindAccountByUinRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBindAccountByUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBindAccountByUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMasterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MfaActions struct {

	// 云api名字

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否激活，0代表未激活，1代表激活

	IsActivate *int64 `json:"IsActivate,omitempty" name:"IsActivate"`
}

type RetentionSubAccountsUserInfo struct {

	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// Uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
}

type GetAccountLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *GetAccountLoginStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountLoginStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginVerifyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// Uin

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// OwnerUin

		OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LoginVerifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountCurInfo struct {

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Type

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Area

	Area *int64 `json:"Area,omitempty" name:"Area"`
	// Idcardtype

	Idcardtype *int64 `json:"Idcardtype,omitempty" name:"Idcardtype"`
	// Idcard

	Idcard *string `json:"Idcard,omitempty" name:"Idcard"`
	// Organizationcode

	Organizationcode *string `json:"Organizationcode,omitempty" name:"Organizationcode"`
	// AuthenticateType

	AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`
}

type GetCustomSessionSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCustomSessionSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomSessionSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInfoByFieldsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInfoByFieldsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInfoByFieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginVerifyRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 域名id

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 应答是否需要昵称

	NeedNick *int64 `json:"NeedNick,omitempty" name:"NeedNick"`
	// 客户端代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
}

func (r *LoginVerifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginVerifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 持久密钥

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOwnerUinByAppidRequest struct {
	*tchttp.BaseRequest

	// appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *GetOwnerUinByAppidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOwnerUinByAppidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCaptchaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Token

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCaptchaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCaptchaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverRetentionSubaccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverRetentionSubaccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverRetentionSubaccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionLoginFlag struct {

	// 电话,&nbsp;传0不开启敏感操作保护,&nbsp;传1开启敏感操作保护。

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 软Token,&nbsp;传0不开启敏感操作保护,&nbsp;传1开启敏感操作保护。

	Stoken *string `json:"Stoken,omitempty" name:"Stoken"`
	// 硬Token,&nbsp;传0不开启敏感操作保护,&nbsp;传1开启敏感操作保护。

	Token *string `json:"Token,omitempty" name:"Token"`
}

type ModifySensitiveActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MfaActions

		MfaActions []*MfaActions `json:"MfaActions,omitempty" name:"MfaActions"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySensitiveActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySensitiveActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBindAccountByUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户唯一id

		Fakeuin *uint64 `json:"Fakeuin,omitempty" name:"Fakeuin"`
		// 用户名

		Account *string `json:"Account,omitempty" name:"Account"`
		// 账户昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 账户绑定类型

		TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
		// 是否为注册账户

		IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
		// 账户状态

		BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`
		// 类型名称

		TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBindAccountByUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBindAccountByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 属性列表

		Attributes []*AccountAttribute `json:"Attributes,omitempty" name:"Attributes"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 是否为主账号

	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 审核状态

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 邮箱是否审核通过

	MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
	// 线下审核状态

	OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
	// 首次购买带外网IP的cvm设备的时间

	WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
	// 外网是否受限

	WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
	// 返回的字段

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IdentifyType

		IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
		// 扩展属性值

		Items []*AccountAttributeAndValue `json:"Items,omitempty" name:"Items"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubAccountInfoRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 主账户uin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *GetSubAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UIN

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 主账号UIN

		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// 是否为主账号

		IsOwner *int64 `json:"IsOwner,omitempty" name:"IsOwner"`
		// 资料是否审核通过

		CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 是否实名认证：0未认证，1已认证

		IsAuthenticate *int64 `json:"IsAuthenticate,omitempty" name:"IsAuthenticate"`
		// 邮箱是否审核通过

		MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 手机号码

		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
		// 用户指引标识位

		GuideBit *int64 `json:"GuideBit,omitempty" name:"GuideBit"`
		// 用户首次购买带外网IP的cvm设备的时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// 标识外网是否受限

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// 修改时间

		ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
		// 线下审核状态

		OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 来源平台

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 是否测试用户

		IsTestUser *int64 `json:"IsTestUser,omitempty" name:"IsTestUser"`
		// 客户来源

		ClientFrom *string `json:"ClientFrom,omitempty" name:"ClientFrom"`
		// register&nbsp;refer

		Referer *string `json:"Referer,omitempty" name:"Referer"`
		// 是否注册成功

		IsRegSucc *bool `json:"IsRegSucc,omitempty" name:"IsRegSucc"`
		// 用户属性集合

		Attributes *int64 `json:"Attributes,omitempty" name:"Attributes"`
		// 否导入了即时通白名单

		Isprotect *int64 `json:"Isprotect,omitempty" name:"Isprotect"`
		// 用户指引

		IsSeeGuidelines *int64 `json:"IsSeeGuidelines,omitempty" name:"IsSeeGuidelines"`
		// 是否接收推广信息

		IsAcceptProMsg *int64 `json:"IsAcceptProMsg,omitempty" name:"IsAcceptProMsg"`
		// 部署模块

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 账号列表

		AccountList *string `json:"AccountList,omitempty" name:"AccountList"`
		// 账号类型

		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
		// 默认开发商

		DefaultOwner *int64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 邮箱验证

		MailVerify *int64 `json:"MailVerify,omitempty" name:"MailVerify"`
		// 接收信息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 地域

		Area *string `json:"Area,omitempty" name:"Area"`
		// 是否需要完善信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 绑定账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 企业微信用户id

		QywxUserId *string `json:"QywxUserId,omitempty" name:"QywxUserId"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否受保护

		IsProtect *int64 `json:"IsProtect,omitempty" name:"IsProtect"`
		// Needresetpassword

		Needresetpassword *int64 `json:"Needresetpassword,omitempty" name:"Needresetpassword"`
		// Updatetime

		Updatetime *string `json:"Updatetime,omitempty" name:"Updatetime"`
		// Username

		Username *string `json:"Username,omitempty" name:"Username"`
		// Istestuser

		Istestuser *int64 `json:"Istestuser,omitempty" name:"Istestuser"`
		// Phonenum

		Phonenum *string `json:"Phonenum,omitempty" name:"Phonenum"`
		// Accounttype

		Accounttype *int64 `json:"Accounttype,omitempty" name:"Accounttype"`
		// Isauthenticate

		Isauthenticate *int64 `json:"Isauthenticate,omitempty" name:"Isauthenticate"`
		// Contactemailflag

		Contactemailflag *int64 `json:"Contactemailflag,omitempty" name:"Contactemailflag"`
		// Consolelogin

		Consolelogin *int64 `json:"Consolelogin,omitempty" name:"Consolelogin"`
		// Mailstatus

		Mailstatus *int64 `json:"Mailstatus,omitempty" name:"Mailstatus"`
		// Countrycode

		Countrycode *string `json:"Countrycode,omitempty" name:"Countrycode"`
		// Checkstatus

		Checkstatus *int64 `json:"Checkstatus,omitempty" name:"Checkstatus"`
		// Wxzsopenid

		Wxzsopenid *string `json:"Wxzsopenid,omitempty" name:"Wxzsopenid"`
		// Srcplatform

		Srcplatform *string `json:"Srcplatform,omitempty" name:"Srcplatform"`
		// Isregsucc

		Isregsucc *int64 `json:"Isregsucc,omitempty" name:"Isregsucc"`
		// Bindstatus

		Bindstatus *int64 `json:"Bindstatus,omitempty" name:"Bindstatus"`
		// Updatetimestamp

		Updatetimestamp *string `json:"Updatetimestamp,omitempty" name:"Updatetimestamp"`
		// Contactemail

		Contactemail *string `json:"Contactemail,omitempty" name:"Contactemail"`
		// Isreceiverowner

		Isreceiverowner *int64 `json:"Isreceiverowner,omitempty" name:"Isreceiverowner"`
		// Defaultowner

		Defaultowner *uint64 `json:"Defaultowner,omitempty" name:"Defaultowner"`
		// Createtime

		Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
		// Qywxuserid

		Qywxuserid *string `json:"Qywxuserid,omitempty" name:"Qywxuserid"`
		// Identifytype

		Identifytype *int64 `json:"Identifytype,omitempty" name:"Identifytype"`
		// Createuin

		Createuin *uint64 `json:"Createuin,omitempty" name:"Createuin"`
		// Wxzsstatus

		Wxzsstatus *int64 `json:"Wxzsstatus,omitempty" name:"Wxzsstatus"`
		// Phoneflag

		Phoneflag *int64 `json:"Phoneflag,omitempty" name:"Phoneflag"`
		// Isowner

		Isowner *int64 `json:"Isowner,omitempty" name:"Isowner"`
		// Isregaccount

		Isregaccount *int64 `json:"Isregaccount,omitempty" name:"Isregaccount"`
		// Addtimestamp

		Addtimestamp *string `json:"Addtimestamp,omitempty" name:"Addtimestamp"`
		// Owneruin

		Owneruin *uint64 `json:"Owneruin,omitempty" name:"Owneruin"`
		// Loginaccount

		Loginaccount *string `json:"Loginaccount,omitempty" name:"Loginaccount"`
		// Typebind

		Typebind *int64 `json:"Typebind,omitempty" name:"Typebind"`
		// Name

		Name *string `json:"Name,omitempty" name:"Name"`
		// Type

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// Uid

		Uid *int64 `json:"Uid,omitempty" name:"Uid"`
		// Loginstatus

		Loginstatus *int64 `json:"Loginstatus,omitempty" name:"Loginstatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMfaDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMfaDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMfaDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdRequest struct {
	*tchttp.BaseRequest

	// 业务传过来的uin

	ParaUin *int64 `json:"ParaUin,omitempty" name:"ParaUin"`
}

func (r *GetAppIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLastLoginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// &nbsp;登录信息

		Data *LoginInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLastLoginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLastLoginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenUnBindResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TokenUnBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenUnBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMailPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 密码，密文

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CheckMailPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMailPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 验证码类型，1-手机验证，2-邮箱验证

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 验证码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 地区码，默认86

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

func (r *CheckVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTokenRequest struct {
	*tchttp.BaseRequest

	// 待删除token

	DelToken *string `json:"DelToken,omitempty" name:"DelToken"`
}

func (r *DeleteTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 多因子配置

		Data *MultIFactor `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMultiFactorParasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMultiFactorParasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoLockPeriodRequest struct {
	*tchttp.BaseRequest

	// IsEnabled

	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// Period

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyAutoLockPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoLockPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OwnerInfo struct {

	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 显示名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 账户状态

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 是否为该子账户的默认主账户

	IsDefaultOwner *bool `json:"IsDefaultOwner,omitempty" name:"IsDefaultOwner"`
	// 手机号，打掩码处理

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱，掩码处理

	Email *string `json:"Email,omitempty" name:"Email"`
	// 扩展属性

	ExtAttr *ExtAttr `json:"ExtAttr,omitempty" name:"ExtAttr"`
	// 多因子认证

	LoginMfa *LoginMfa `json:"LoginMfa,omitempty" name:"LoginMfa"`
}

type UserInfo struct {

	// 子账号类型，传0不可登陆控制台，传1可以登陆控制台。示例值:&nbsp;1

	CanLogin *string `json:"CanLogin,omitempty" name:"CanLogin"`
	// 区号。

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 详情

	Detail *AccountDetail `json:"Detail,omitempty" name:"Detail"`
	// 名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 电话号码。

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 安全邮箱。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 昵称。

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 账号类型。

	SyStemType *string `json:"SyStemType,omitempty" name:"SyStemType"`
}

type GetSafeAuthConfigRequest struct {
	*tchttp.BaseRequest

	// 查询的用户uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *GetSafeAuthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSafeAuthConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSingleLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSingleLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSingleLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRetentionPolicyRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOwnerUinByAppidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// appid

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// 账户创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOwnerUinByAppidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOwnerUinByAppidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtAttr struct {

	// 需要重置mfa的token

	NeedResetToken *int64 `json:"NeedResetToken,omitempty" name:"NeedResetToken"`
	// 需要重置mfa的token

	NeedResetStoken *int64 `json:"NeedResetStoken,omitempty" name:"NeedResetStoken"`
}

type CheckSubAccountUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的子账户uin列表

		Exist []*uint64 `json:"Exist,omitempty" name:"Exist"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSubAccountUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelatedUinSessionKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 关联账号

		RelatedUin *string `json:"RelatedUin,omitempty" name:"RelatedUin"`
		// 关联账号的主账号

		RelatedOwnerUin *string `json:"RelatedOwnerUin,omitempty" name:"RelatedOwnerUin"`
		// 关联登录态

		RelatedSessionKey *string `json:"RelatedSessionKey,omitempty" name:"RelatedSessionKey"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelatedUinSessionKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelatedUinSessionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRetentionSubAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalNum

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// UserInfo

		UserInfo []*RetentionUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRetentionSubAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRetentionSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSafeAuthConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token信息

		TokenInfo *SafeAuthTokenInfo `json:"TokenInfo,omitempty" name:"TokenInfo"`
		// 标识信息

		Flag *SafeAuthFlagInfo `json:"Flag,omitempty" name:"Flag"`
		// ukey信息

		UkeyInfo *SafeAuthTokenInfo `json:"UkeyInfo,omitempty" name:"UkeyInfo"`
		// MfaHardTokenSource

		MfaHardTokenSource *string `json:"MfaHardTokenSource,omitempty" name:"MfaHardTokenSource"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSafeAuthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSafeAuthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OmitOffsiteDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OmitOffsiteDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OmitOffsiteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAccountsRequest struct {
	*tchttp.BaseRequest

	// 子账号列表

	FilterSubAccountUin []*uint64 `json:"FilterSubAccountUin,omitempty" name:"FilterSubAccountUin"`
}

func (r *DescribeSubAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomSessionSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SessionTimeout

		SessionTimeout *int64 `json:"SessionTimeout,omitempty" name:"SessionTimeout"`
		// Scope

		Scope *int64 `json:"Scope,omitempty" name:"Scope"`
		// LastModifyTime

		LastModifyTime *string `json:"LastModifyTime,omitempty" name:"LastModifyTime"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomSessionSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomSessionSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户信息列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// 用户id

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListV2Request struct {
	*tchttp.BaseRequest

	// 是否获取账户属性

	WithAttr *int64 `json:"WithAttr,omitempty" name:"WithAttr"`
}

func (r *GetMasterListV2Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListV2Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKeyDetail `json:"IdKeys,omitempty" name:"IdKeys"`
		// 账号Uin

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// 项目Id

		SecretProjectId *uint64 `json:"SecretProjectId,omitempty" name:"SecretProjectId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountLoginStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可登陆状态，0-正常，1-临时锁定，2-运营端锁定

		LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountLoginStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNicknameRequest struct {
	*tchttp.BaseRequest

	// 用户新昵称

	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}

func (r *ModifyNicknameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNicknameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListV2Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// 账户uid

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListV2Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListWithStatusRequest struct {
	*tchttp.BaseRequest

	// 客户端用户代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *GetMasterListWithStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListWithStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户信息

	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 是否在白名单

	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`
	// 从Api创建

	FromAPI *uint64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *AddSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountExistRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *CheckAccountExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCountryCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCountryCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaxSubAccountAndGroupNumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MaxSubAccountNum

		MaxSubAccountNum *int64 `json:"MaxSubAccountNum,omitempty" name:"MaxSubAccountNum"`
		// MaxGroupNum

		MaxGroupNum *int64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaxSubAccountAndGroupNumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaxSubAccountAndGroupNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretIDLastUsedRow struct {

	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 最后访问日期

	LastUsedDate *string `json:"LastUsedDate,omitempty" name:"LastUsedDate"`
	// 最后密钥访问日期

	LastSecretUsedDate *int64 `json:"LastSecretUsedDate,omitempty" name:"LastSecretUsedDate"`
}

type QuserGetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UIN

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 主账号UIN

		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// 是否为主账号

		IsOwner *int64 `json:"IsOwner,omitempty" name:"IsOwner"`
		// 资料是否审核通过

		CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 是否实名认证：0未认证，1已认证

		IsAuthenticate *int64 `json:"IsAuthenticate,omitempty" name:"IsAuthenticate"`
		// 邮箱是否审核通过

		MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 手机号码

		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
		// 用户指引标识位

		GuideBit *int64 `json:"GuideBit,omitempty" name:"GuideBit"`
		// 用户首次购买带外网IP的cvm设备的时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// 标识外网是否受限

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// 修改时间

		ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
		// 线下审核状态

		OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 来源平台

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 是否测试用户

		IsTestUser *int64 `json:"IsTestUser,omitempty" name:"IsTestUser"`
		// 客户来源

		ClientFrom *string `json:"ClientFrom,omitempty" name:"ClientFrom"`
		// register&nbsp;refer

		Referer *string `json:"Referer,omitempty" name:"Referer"`
		// 是否注册成功

		IsRegSucc *bool `json:"IsRegSucc,omitempty" name:"IsRegSucc"`
		// 用户属性集合

		Attributes *int64 `json:"Attributes,omitempty" name:"Attributes"`
		// 否导入了即时通白名单

		Isprotect *int64 `json:"Isprotect,omitempty" name:"Isprotect"`
		// 用户指引

		IsSeeGuidelines *int64 `json:"IsSeeGuidelines,omitempty" name:"IsSeeGuidelines"`
		// 是否接收推广信息

		IsAcceptProMsg *int64 `json:"IsAcceptProMsg,omitempty" name:"IsAcceptProMsg"`
		// 部署模块

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 账号列表

		AccountList *string `json:"AccountList,omitempty" name:"AccountList"`
		// 账号类型

		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
		// 默认开发商

		DefaultOwner *int64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 邮箱验证

		MailVerify *int64 `json:"MailVerify,omitempty" name:"MailVerify"`
		// 接收信息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 地域

		Area *string `json:"Area,omitempty" name:"Area"`
		// 是否需要完善信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 绑定账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuserGetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuserGetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenBindResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TokenBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAutoLockPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Period

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// IsEnabled

		IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAutoLockPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAutoLockPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CategoryList struct {

	// 类别Id

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 类别键

	CategoryKey *string `json:"CategoryKey,omitempty" name:"CategoryKey"`
	// 类别值

	CategoryValue *string `json:"CategoryValue,omitempty" name:"CategoryValue"`
}

type GetNicknameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 显示的名称

		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
		// 用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 主账号

		DefaultOwner *uint64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNicknameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNicknameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
	// 备注

	ApiSecretIdRemark *string `json:"ApiSecretIdRemark,omitempty" name:"ApiSecretIdRemark"`
}

func (r *ModifySelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCountryCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCountryCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appid

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecurityLastUsedRequest struct {
	*tchttp.BaseRequest

	// 查询密钥ID列表

	SecretIdList []*string `json:"SecretIdList,omitempty" name:"SecretIdList"`
}

func (r *GetSecurityLastUsedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecurityLastUsedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinOwnerInOpenRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUinOwnerInOpenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinOwnerInOpenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNicknameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNicknameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNicknameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserAreaByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserByAttributeValueRequest struct {
	*tchttp.BaseRequest

	// 扩展属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 扩展属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 分页每页大小，默认20

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 分页值，从1开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *GetUserByAttributeValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByAttributeValueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被删除的目标uin，不填则删除登录账户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 属性列表

	Attributes []*string `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *DeleteAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecurityLastUsedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥ID最近访问列表

		SecretIdLastUsedRows []*SecretIDLastUsedRow `json:"SecretIdLastUsedRows,omitempty" name:"SecretIdLastUsedRows"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSecurityLastUsedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecurityLastUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionUserInfo struct {

	// Uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// OwnerUin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Nickname

	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
	// PhoneNum

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// CountryCode

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// PhoneFlag

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// Email

	Email *string `json:"Email,omitempty" name:"Email"`
	// EmailFlag

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// SystemType

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// LoginStatus

	LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
}

type DescribeUserCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 类别

		CategoryList []*CategoryList `json:"CategoryList,omitempty" name:"CategoryList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTokenRequest struct {
	*tchttp.BaseRequest

	// 过期时间，单位s

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *GetTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地区枚举值

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// 国家名称

		CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAreaByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSubAccountUinRequest struct {
	*tchttp.BaseRequest

	// 子账户uin列表

	AccountUin []*uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *CheckSubAccountUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRetentionSubaccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRetentionSubaccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionSubaccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthOffsiteFlag struct {

	// 验证标识

	VerifyFlag *int64 `json:"VerifyFlag,omitempty" name:"VerifyFlag"`
	// 是否进行电话通知

	NotifyPhone *int64 `json:"NotifyPhone,omitempty" name:"NotifyPhone"`
	// 是否进行

	NotifyEmail *int64 `json:"NotifyEmail,omitempty" name:"NotifyEmail"`
}

type SetCaptchaRequest struct {
	*tchttp.BaseRequest

	// 验证码关联token

	CaptchaToken *string `json:"CaptchaToken,omitempty" name:"CaptchaToken"`
	// 新设置的图形验证码

	Captcha *string `json:"Captcha,omitempty" name:"Captcha"`
}

func (r *SetCaptchaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCaptchaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginInfo struct {

	// 客户端ip

	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
	// 客户端代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 设备id

	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 登录地址

	Location *string `json:"Location,omitempty" name:"Location"`
	// 登录标识

	LoginFlag *int64 `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// 登录时间

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 主账户uin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// LoginId

	LoginId *int64 `json:"LoginId,omitempty" name:"LoginId"`
}

type SetMfaDeviceRequest struct {
	*tchttp.BaseRequest

	// 1-三方设备,&nbsp;2-虚拟mfa设备

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// 被分配客户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *SetMfaDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMfaDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRetentionPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RetentionDays

		RetentionDays *int64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
		// RetentionPolicy

		RetentionPolicy *int64 `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMailPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否正确

		IsRight *bool `json:"IsRight,omitempty" name:"IsRight"`
		// 账户唯一id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMailPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMailPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverRetentionSubaccountRequest struct {
	*tchttp.BaseRequest

	// UserInfo

	UserInfo []*RetentionSubAccountsUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
}

func (r *RecoverRetentionSubaccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverRetentionSubaccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthTokenInfo struct {

	// 状态，0-未分配，1-已分配，2-已绑定，3-已解绑

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// token序列号

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
	// token类型，1-hard&nbsp;token,&nbsp;2-soft&nbsp;token，3-ukey

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// &nbsp;设备厂商来源

	TokenSource *string `json:"TokenSource,omitempty" name:"TokenSource"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type SubAccountUser struct {

	// 子用户&nbsp;UID

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 子用户用户&nbsp;ID

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 主账号ID

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 联系手机

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 电话号码是否验证

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 联系邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱是否验证

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 是否允许登录

	CanLogin *int64 `json:"CanLogin,omitempty" name:"CanLogin"`
	// 是否允许控制台登录

	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 微信公众号关注状态

	WxzsStatus *int64 `json:"WxzsStatus,omitempty" name:"WxzsStatus"`
	// 是否需要重置密码

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 额外属性

	ExtAttr *ExtAttr `json:"ExtAttr,omitempty" name:"ExtAttr"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// UserType

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// LastLoginTime

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastLoginIP

	LastLoginIP *string `json:"LastLoginIP,omitempty" name:"LastLoginIP"`
	// LastLoginIp

	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`
}

type CheckSubAccountNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的子账户用户名列表

		Exist []*string `json:"Exist,omitempty" name:"Exist"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSubAccountNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 自定义SecretId

	CustomSecretId *string `json:"CustomSecretId,omitempty" name:"CustomSecretId"`
	// 自定义SecretKey

	CustomSecretKey *string `json:"CustomSecretKey,omitempty" name:"CustomSecretKey"`
}

func (r *CreateSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DisableSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAppIdByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckTokenRequest struct {
	*tchttp.BaseRequest

	// 待校验的token

	CheckToken *string `json:"CheckToken,omitempty" name:"CheckToken"`
}

func (r *CheckTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByOwnerUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdByOwnerUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByOwnerUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被查询的账户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *GetAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFrequentLoginConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 常用登录地计算天数

		FrequentLocationDays *int64 `json:"FrequentLocationDays,omitempty" name:"FrequentLocationDays"`
		// 常用登录地登陆次数

		FrequentLocationTimes *int64 `json:"FrequentLocationTimes,omitempty" name:"FrequentLocationTimes"`
		// 常用登录地数量

		FrequentLocationCount *int64 `json:"FrequentLocationCount,omitempty" name:"FrequentLocationCount"`
		// 常用登录设备计算天数

		FrequentDeviceDays *int64 `json:"FrequentDeviceDays,omitempty" name:"FrequentDeviceDays"`
		// 常用登录设备登陆次数

		FrequentDeviceTimes *int64 `json:"FrequentDeviceTimes,omitempty" name:"FrequentDeviceTimes"`
		// 常用登录设备数量

		FrequentDeviceCount *int64 `json:"FrequentDeviceCount,omitempty" name:"FrequentDeviceCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFrequentLoginConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFrequentLoginConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserIdAttrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserIdAttrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserIdAttrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoLockPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoLockPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoLockPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeMailPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeMailPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttributeValue struct {

	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type AccountExistData struct {

	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 账户状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 账户绑定类型

	TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
	// 是否为注册账户

	IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
	// 最近一次更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ApiKey struct {

	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效,&nbsp;3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 数据源

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 密钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type ClearLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *ClearLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFrequentLoginConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetFrequentLoginConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFrequentLoginConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserInfoByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 登陆期临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 登陆临时id

	Tinyid *string `json:"Tinyid,omitempty" name:"Tinyid"`
}

func (r *SeedLoginTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedLoginTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRetentionSubaccountRequest struct {
	*tchttp.BaseRequest

	// UserInfo

	UserInfo []*RetentionSubAccountsUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
}

func (r *DeleteRetentionSubaccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionSubaccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthFlagInfo struct {

	// 登陆标识

	LoginFlag *SafeAuthFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// action标识

	ActionFlag *SafeAuthFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 外围标识

	OffsiteFlag *SafeAuthOffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`
}

type CreateSoftTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token序列号

		TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
		// 链接数据，用于生成google&nbsp;authentication客户端绑定二维码，base64编码

		QrcodeData *string `json:"QrcodeData,omitempty" name:"QrcodeData"`
		// 过期时间戳，秒

		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// token标识

		Mark *string `json:"Mark,omitempty" name:"Mark"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSoftTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoftTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInfoByFieldsRequest struct {
	*tchttp.BaseRequest

	// 账户信息

	Account *AccountRequest `json:"Account,omitempty" name:"Account"`
	// 查询属性列表

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetInfoByFieldsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInfoByFieldsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMultiFactorParasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMultiFactorParasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserIdAttrRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinArr []*uint64 `json:"UinArr,omitempty" name:"UinArr"`
	// appid列表

	AppIdArr []*uint64 `json:"AppIdArr,omitempty" name:"AppIdArr"`
}

func (r *GetUserIdAttrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserIdAttrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRetentionPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RetentionDays

		RetentionDays *int64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
		// RetentionPolicy

		RetentionPolicy *int64 `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaskedUserInfoByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMaskedUserInfoByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaskedUserInfoByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SetLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSingleLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SetSingleLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSingleLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuserGetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 查询字段，字段名逗号分隔

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *QuserGetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuserGetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCaptchaRequest struct {
	*tchttp.BaseRequest

	// 验证码

	Captcha *string `json:"Captcha,omitempty" name:"Captcha"`
	// 验证生命周期用的token

	CaptchaToken *string `json:"CaptchaToken,omitempty" name:"CaptchaToken"`
}

func (r *CheckCaptchaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCaptchaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCategoryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByOwnerUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAppIdByOwnerUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByOwnerUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttributeAndValue struct {

	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 值id

	ValueId *int64 `json:"ValueId,omitempty" name:"ValueId"`
	// uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckAccountExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户唯一id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		Account *string `json:"Account,omitempty" name:"Account"`
		// 昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 账户状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 账户绑定类型

		TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
		// 是否为注册账户

		IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
		// 最近一次更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAccountExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DeleteSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLastLoginInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLastLoginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLastLoginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthFlag struct {

	// 是否进行手机号认证，0-否，1-是

	Phone *int64 `json:"Phone,omitempty" name:"Phone"`
	// 是否进行token认证，0-否，1-是

	Token *int64 `json:"Token,omitempty" name:"Token"`
	// 是否进行stoken认证，0-否，1-是

	Stoken *int64 `json:"Stoken,omitempty" name:"Stoken"`
	// 是否进行ukey检查

	Ukey *int64 `json:"Ukey,omitempty" name:"Ukey"`
	// Global

	Global *int64 `json:"Global,omitempty" name:"Global"`
}

type ApiKeyDetail struct {

	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 密钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效,&nbsp;3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 数据源

	Source *uint64 `json:"Source,omitempty" name:"Source"`
}

type OmitOffsiteDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备id

	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
	// 登录记录id

	LoginId *int64 `json:"LoginId,omitempty" name:"LoginId"`
}

func (r *OmitOffsiteDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OmitOffsiteDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSubAccountNameRequest struct {
	*tchttp.BaseRequest

	// 用户名列表

	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *CheckSubAccountNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaskedUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UIN

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 主账号UIN

		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// 是否为主账号

		IsOwner *int64 `json:"IsOwner,omitempty" name:"IsOwner"`
		// 资料是否审核通过

		CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 是否实名认证：0未认证，1已认证

		IsAuthenticate *int64 `json:"IsAuthenticate,omitempty" name:"IsAuthenticate"`
		// 邮箱是否审核通过

		MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 手机号码

		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
		// 用户指引标识位

		GuideBit *int64 `json:"GuideBit,omitempty" name:"GuideBit"`
		// 用户首次购买带外网IP的cvm设备的时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// 标识外网是否受限

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// 修改时间

		ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
		// 线下审核状态

		OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 来源平台

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 是否测试用户

		IsTestUser *int64 `json:"IsTestUser,omitempty" name:"IsTestUser"`
		// 客户来源

		ClientFrom *string `json:"ClientFrom,omitempty" name:"ClientFrom"`
		// register&nbsp;refer

		Referer *string `json:"Referer,omitempty" name:"Referer"`
		// 是否注册成功

		IsRegSucc *bool `json:"IsRegSucc,omitempty" name:"IsRegSucc"`
		// 用户属性集合

		Attributes *int64 `json:"Attributes,omitempty" name:"Attributes"`
		// 否导入了即时通白名单

		Isprotect *int64 `json:"Isprotect,omitempty" name:"Isprotect"`
		// 用户指引

		IsSeeGuidelines *int64 `json:"IsSeeGuidelines,omitempty" name:"IsSeeGuidelines"`
		// 是否接收推广信息

		IsAcceptProMsg *int64 `json:"IsAcceptProMsg,omitempty" name:"IsAcceptProMsg"`
		// 部署模块

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 账号列表

		AccountList *string `json:"AccountList,omitempty" name:"AccountList"`
		// 账号类型

		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
		// 默认开发商

		DefaultOwner *int64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 邮箱验证

		MailVerify *int64 `json:"MailVerify,omitempty" name:"MailVerify"`
		// 接收信息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 地域

		Area *string `json:"Area,omitempty" name:"Area"`
		// 是否需要完善信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 绑定账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 企业微信用户id

		QywxUserId *string `json:"QywxUserId,omitempty" name:"QywxUserId"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 是否受保护

		IsProtect *int64 `json:"IsProtect,omitempty" name:"IsProtect"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaskedUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaskedUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListWithStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户信息列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// uid

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListWithStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListWithStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户ip

		ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
		// 客户userAgent

		ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
		// 登录地址

		Location *string `json:"Location,omitempty" name:"Location"`
		// 登录时间

		LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
		// 设备id

		DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
		// 登录记录id

		LoginId *int64 `json:"LoginId,omitempty" name:"LoginId"`
		// 登录标识。1-非常用设备登录

		LoginFlag *int64 `json:"LoginFlag,omitempty" name:"LoginFlag"`
		// &nbsp;登录信息

		Data *LoginInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLoginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetTradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MFAStatusRequest struct {
	*tchttp.BaseRequest

	// 接口名

	Interface *string `json:"Interface,omitempty" name:"Interface"`
	// 客户端代理信息

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *MFAStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MFAStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被设置的目标uin，不填则设置登录用户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 属性值列表

	Attributes []*AccountAttributeValue `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *SetAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MFAStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否需要认证，0-不需要，1-需要

		NeedAuth *int64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
		// 过期时间

		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 认证类型

		AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
		// 认证的手机号

		AuthPhone *string `json:"AuthPhone,omitempty" name:"AuthPhone"`
		// 认证的token

		AuthHardtoken *string `json:"AuthHardtoken,omitempty" name:"AuthHardtoken"`
		// 是否来自小程序

		IsWeapp *int64 `json:"IsWeapp,omitempty" name:"IsWeapp"`
		// 认证的email

		AuthEmail *string `json:"AuthEmail,omitempty" name:"AuthEmail"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MFAStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MFAStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRetentionSubAccountsRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Keyword

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListRetentionSubAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRetentionSubAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenUnBindRequest struct {
	*tchttp.BaseRequest

	// 解绑目标账号

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// token类型，1-hard&nbsp;token&nbsp;2-soft&nbsp;token&nbsp;3-ukey

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// 序列号，ukey类型会进行校验

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
}

func (r *TokenUnBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenUnBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRetentionPolicyRequest struct {
	*tchttp.BaseRequest

	// RetentionDays

	RetentionDays *int64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// RetentionPolicy

	RetentionPolicy *int64 `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *ModifyRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRetentionPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginInfoRequest struct {
	*tchttp.BaseRequest

	// 数据类型，0-上次登录信息，1-上次非常用设备登录信息

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetLoginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinOwnerInOpenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUinOwnerInOpenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinOwnerInOpenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 账号uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 项目id

		SecretProjectId *int64 `json:"SecretProjectId,omitempty" name:"SecretProjectId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountRequest struct {

	// 账户类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 账户id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主账户id

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 子账户id列表

	SubAccountList *string `json:"SubAccountList,omitempty" name:"SubAccountList"`
	// uid列表

	UidList *string `json:"UidList,omitempty" name:"UidList"`
	// 角色列表

	RoleList *string `json:"RoleList,omitempty" name:"RoleList"`
	// 用户组列表

	GroupList *string `json:"GroupList,omitempty" name:"GroupList"`
}

type CreateSoftTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateSoftTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoftTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *GetSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeUserInfo struct {

	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type GetSubAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子账户uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 是否是控制台登陆

		ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
		// 是否需要重置密码。需要-1

		NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubLoginUinListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubLoginUinListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubLoginUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendLoginVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendLoginVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLoginVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySensitiveActionRequest struct {
	*tchttp.BaseRequest

	// Operate

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// ActionName

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySensitiveActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySensitiveActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySelfApiKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *QuerySelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 需要验证的操作

	VerifyAction *string `json:"VerifyAction,omitempty" name:"VerifyAction"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 验证码类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 手机号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 手机区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 邮箱号码

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 修改来源

	From *string `json:"From,omitempty" name:"From"`
	// 目标账号uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *SendVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSafeAuthFlagRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 登陆标识

	LoginFlag *SafeAuthFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// action标识

	ActionFlag *SafeAuthFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 设备序列号

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
}

func (r *SetSafeAuthFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSafeAuthFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenBindRequest struct {
	*tchttp.BaseRequest

	// token类型，1-hard&nbsp;token,&nbsp;2-soft&nbsp;token;3-ukey

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// 被设置的用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// token序列号，tokenType&nbsp;=&nbsp;soft&nbsp;token时，必填

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
	// 验证码，动态口令

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *TokenBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountDetail struct {

	// 敏感操作标识

	ActionFlag *ActionLoginFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 是否允许控制台登录,&nbsp;传0不可登陆控制台,&nbsp;传1可以登陆控制台。

	ConsoleLogin *string `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 登录保护

	LoginFlag *ActionLoginFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// 是否需要重置密码,&nbsp;传0不需要重置密码,&nbsp;传1需要重置密码。

	NeedResetPassword *string `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 用户密码。示例值:&nbsp;password

	Password *string `json:"Password,omitempty" name:"Password"`
	// 使用Api,&nbsp;传0不使用Api,&nbsp;传1使用Api。示例值:&nbsp;1

	UseApi *string `json:"UseApi,omitempty" name:"UseApi"`
}

type UserProductUsableInfo struct {

	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 产品id

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 是否打开

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 总量

	Nums *int64 `json:"Nums,omitempty" name:"Nums"`
	// DeadNums

	DeadNums *int64 `json:"DeadNums,omitempty" name:"DeadNums"`
	// 前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 区域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type ChangeSubAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ChangeSubAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeSubAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaskedUserInfoRequest struct {
	*tchttp.BaseRequest

	// 是否为主账号

	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 审核状态

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 邮箱是否审核通过

	MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
	// 线下审核状态

	OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
	// 首次购买带外网IP的cvm设备的时间

	WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
	// 外网是否受限

	WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
	// 返回的字段

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetMaskedUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaskedUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubLoginUinListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSubLoginUinListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubLoginUinListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCaptchaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCaptchaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCaptchaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号Uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 子用户uin列表

		List []*uint64 `json:"List,omitempty" name:"List"`
		// 子用户Uid

		Uids []*uint64 `json:"Uids,omitempty" name:"Uids"`
		// 子用户详情

		SubAccounts []*SubAccounts `json:"SubAccounts,omitempty" name:"SubAccounts"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeNameRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAttributeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

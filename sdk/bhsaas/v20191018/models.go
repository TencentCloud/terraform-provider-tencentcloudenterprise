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

package v20191018

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CheckLDAPConnectionRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，必须为true

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 网络域id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base&nbsp;DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
}

func (r *CheckLDAPConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLDAPConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {

	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户名,&nbsp;3-20个字符&nbsp;必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户姓名，&nbsp;最大20个字符，不能包含空白字符

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 手机号码，&nbsp;大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如:&nbsp;"+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组列表

	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
	// 认证方式，0&nbsp;-&nbsp;本地，1&nbsp;-&nbsp;LDAP，2&nbsp;-&nbsp;OAuth

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制，&nbsp;由0、1组成的字符串，长度168(7&nbsp;×&nbsp;24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时，&nbsp;0&nbsp;-&nbsp;代表不允许访问，1&nbsp;-&nbsp;代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 用户所属部门（用于出参）

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 用户所属部门（用于入参）

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 激活状态&nbsp;0&nbsp;-&nbsp;未激活&nbsp;1&nbsp;-&nbsp;激活

	ActiveStatus *uint64 `json:"ActiveStatus,omitempty" name:"ActiveStatus"`
	// 锁定状态&nbsp;0&nbsp;-&nbsp;未锁定&nbsp;1&nbsp;-&nbsp;锁定

	LockStatus *uint64 `json:"LockStatus,omitempty" name:"LockStatus"`
	// 状态&nbsp;与Filter中一致

	Status *string `json:"Status,omitempty" name:"Status"`
	// ioa同步过来的用户相关信息

	IOAUserGroup *IOAUserGroup `json:"IOAUserGroup,omitempty" name:"IOAUserGroup"`
	// 权限版本

	AclVersion *uint64 `json:"AclVersion,omitempty" name:"AclVersion"`
	// 用户来源，0-bh,1-ioa

	UserFrom *uint64 `json:"UserFrom,omitempty" name:"UserFrom"`
	// ukey绑定状态&nbsp;0&nbsp;-&nbsp;未绑定&nbsp;1&nbsp;-&nbsp;已绑定

	UKeyStatus *int64 `json:"UKeyStatus,omitempty" name:"UKeyStatus"`
}

type DeleteUKeysRequest struct {
	*tchttp.BaseRequest

	// UKey&nbsp;ID集合

	KeyIdSet []*string `json:"KeyIdSet,omitempty" name:"KeyIdSet"`
}

func (r *DeleteUKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableIntranetAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableIntranetAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableIntranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetDeviceAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令列表

		Commands []*SearchCommandResult `json:"Commands,omitempty" name:"Commands"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclsRequest struct {
	*tchttp.BaseRequest

	// 有访问权限的应用资产ID集合

	AuthorizedAppAssetIdSet []*uint64 `json:"AuthorizedAppAssetIdSet,omitempty" name:"AuthorizedAppAssetIdSet"`
	// 访问权限ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 访问权限名称，模糊查询，最长64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否根据Name进行精确查询，默认值false

	Exact *bool `json:"Exact,omitempty" name:"Exact"`
	// 过滤数组

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 有访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
	// 有访问权限的资产ID集合

	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`
	// 访问权限状态，1&nbsp;-&nbsp;已生效，2&nbsp;-&nbsp;未生效，3&nbsp;-&nbsp;已过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否根据AuthorizedDeviceIdSet,对资产账号进行精确匹配，默认false,&nbsp;设置true时，确保AuthorizedDeviceIdSet只有一个元素

	ExactAccount *bool `json:"ExactAccount,omitempty" name:"ExactAccount"`
	// 部门ID，用于过滤属于某个部门的访问权限

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceAccount struct {

	// 账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机ID

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
	// true-已托管密码，false-未托管密码

	BoundPassword *bool `json:"BoundPassword,omitempty" name:"BoundPassword"`
	// true-已托管私钥，false-未托管私钥

	BoundPrivateKey *bool `json:"BoundPrivateKey,omitempty" name:"BoundPrivateKey"`
	// 是否托管凭证,&nbsp;true-托管，false-未托管

	BoundKubeconfig *bool `json:"BoundKubeconfig,omitempty" name:"BoundKubeconfig"`
	// 是否为k8s资产管理账号

	IsK8SManageAccount *bool `json:"IsK8SManageAccount,omitempty" name:"IsK8SManageAccount"`
}

type ExternalDevice struct {

	// 主机名，可为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统名称，只能是Linux、Windows或MySQL

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产所属的部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 资产多节点：字段ip和端口

	IpPortSet []*string `json:"IpPortSet,omitempty" name:"IpPortSet"`
	// SSL证书，EnableSSL时必填

	SSLCert *string `json:"SSLCert,omitempty" name:"SSLCert"`
	// 是否启用SSL,1:启用&nbsp;0：禁用，仅支持Redis资产

	EnableSSL *int64 `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// SSL证书名称，EnableSSL时必填

	SSLCertName *string `json:"SSLCertName,omitempty" name:"SSLCertName"`
}

type DestroyDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogOutputSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogOutputSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogOutputSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDeviceCountSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangePwdTaskDetail struct {

	// 资产信息

	Device *Device `json:"Device,omitempty" name:"Device"`
	// 资产账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 上次改密结果。0-未改密&nbsp;&nbsp;1-改密成功&nbsp;2-改密失败

	LastChangeStatus *uint64 `json:"LastChangeStatus,omitempty" name:"LastChangeStatus"`
}

type DescribeDomainInstallScriptRequest struct {
	*tchttp.BaseRequest

	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainInstallScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainInstallScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogOutputSettings struct {

	// 是否已开启日志外发

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 是否已开启产品登录日志发送

	ProductLogin *bool `json:"ProductLogin,omitempty" name:"ProductLogin"`
	// 是否已开启资产登录日志发送

	DeviceLogin *bool `json:"DeviceLogin,omitempty" name:"DeviceLogin"`
}

type DeleteReportTaskHistoryRequest struct {
	*tchttp.BaseRequest

	// 报表任务记录&nbsp;IdSet

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteReportTaskHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportTaskHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginEventRequest struct {
	*tchttp.BaseRequest

	// 用户名，如果不包含其他条件时对user_name&nbsp;or&nbsp;real_name两个字段模糊查询

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，模糊查询

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 查询时间范围，起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询时间范围，结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 来源IP，模糊查询

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 登录入口：1-字符界面,2-图形界面，3-web页面,&nbsp;4-API

	Entry *uint64 `json:"Entry,omitempty" name:"Entry"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页记录数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLoginEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesPortRequest struct {
	*tchttp.BaseRequest

	// 主机记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyDevicesPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReportTaskRequest struct {
	*tchttp.BaseRequest

	// 创建用户name

	User *string `json:"User,omitempty" name:"User"`
	// 周期cron&nbsp;表达式

	Period *string `json:"Period,omitempty" name:"Period"`
	// 报表任务名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 报表模板&nbsp;id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 报表模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 报表任务说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 单次报表-数据开始时间

	DataStartTime *string `json:"DataStartTime,omitempty" name:"DataStartTime"`
	// 单次报表-数据结束时间

	DataEndTime *string `json:"DataEndTime,omitempty" name:"DataEndTime"`
	// 数据时间范围，单位：小时

	DataTimeRange *uint64 `json:"DataTimeRange,omitempty" name:"DataTimeRange"`
	// 所属部门id。“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 报表任务类型：1&nbsp;单次报表，2&nbsp;周期报表类型

	ReportType *uint64 `json:"ReportType,omitempty" name:"ReportType"`
}

func (r *CreateReportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 权限控制模版列表

		TemplateSet []*ACTemplateResponse `json:"TemplateSet,omitempty" name:"TemplateSet"`
		// total行数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产组名，最大长度32字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产组所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateDeviceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 模版描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *ModifyAccessControlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteChangePwdTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteChangePwdTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableIntranetAccessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableIntranetAccessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableIntranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Domain struct {

	// 自增id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 网络域id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 堡垒机id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// ip，网段

	WhiteIpSet []*string `json:"WhiteIpSet,omitempty" name:"WhiteIpSet"`
	// 是否启用&nbsp;&nbsp;默认&nbsp;1启用&nbsp;0禁用

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 网络域创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 网络域名称

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// 状态&nbsp;0-已断开&nbsp;&nbsp;1-已连接

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否资源默认网络域&nbsp;1-资源默认网络域&nbsp;0-用户添加网络域

	Default *uint64 `json:"Default,omitempty" name:"Default"`
}

type DescribeLoginEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登录日志列表

		LoginEventSet []*LoginEvent `json:"LoginEventSet,omitempty" name:"LoginEventSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveDevice struct {

	// 主机ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 地域信息

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 活跃主机的活跃次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type AuthModeSetting struct {

	// 双因子认证，0-不开启，1-OTP，2-短信

	AuthMode *uint64 `json:"AuthMode,omitempty" name:"AuthMode"`
}

type DescribeOperationTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作类型名称与值的对应关系

		OperationTypeSet []*Map `json:"OperationTypeSet,omitempty" name:"OperationTypeSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePushAccountTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤数组，支持以下FilterName:
	// InstanceId&nbsp;实例ID
	// DeviceName&nbsp;实例名称
	// Ip&nbsp;公网ip或vpc&nbsp;ip

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移位置，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20,&nbsp;最大500。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePushAccountTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushAccountTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialGuideRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTrialGuideRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialGuideRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDeviceAccountRequest struct {
	*tchttp.BaseRequest

	// 账号列表

	DeviceAccountSet []*ParamDeviceAccount `json:"DeviceAccountSet,omitempty" name:"DeviceAccountSet"`
}

func (r *ImportDeviceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeviceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlTemplatesRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 描述键值对过滤器数据，用于条件过滤查询。filter的name目前支持template_id，template_name

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccessControlTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 要查询的运维任务ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 分页每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 资产名称或资产IP

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产地域过滤

	ApCode []*string `json:"ApCode,omitempty" name:"ApCode"`
}

func (r *DescribeOperationTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecuritySettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecuritySettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecuritySettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOAUserGroup struct {

	// ioa用户组织id

	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
	// ioa用户组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// ioa用户组织id路径

	OrgIdPath *string `json:"OrgIdPath,omitempty" name:"OrgIdPath"`
	// ioa用户组织名称路径

	OrgNamePath *string `json:"OrgNamePath,omitempty" name:"OrgNamePath"`
	// ioa关联用户源类型

	Source *uint64 `json:"Source,omitempty" name:"Source"`
}

type AccessDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 认证信息

		AccessInfo *AccessInfo `json:"AccessInfo,omitempty" name:"AccessInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户名或用户姓名，最长64个字符，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// true&nbsp;-&nbsp;查询已添加到该用户组的用户，false&nbsp;-&nbsp;查询未添加到该用户组的用户

	Bound *bool `json:"Bound,omitempty" name:"Bound"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据库操作列表

		StatementSet []*DatabaseStatement `json:"StatementSet,omitempty" name:"StatementSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStatementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginEvent struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 来源IP

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 登录入口：1-字符界面,2-图形界面，3-web页面,&nbsp;4-API

	Entry *uint64 `json:"Entry,omitempty" name:"Entry"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
}

type KeyCount struct {

	// 用户名或日期

	Key *string `json:"Key,omitempty" name:"Key"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeAccountsWithDeviceCountRequest struct {
	*tchttp.BaseRequest

	// 主机ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 主机组ID集合

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
}

func (r *DescribeAccountsWithDeviceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsWithDeviceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		ReportTasks []*ReportTask `json:"ReportTasks,omitempty" name:"ReportTasks"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchPushAccountTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 搜索区间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchPushAccountTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchPushAccountTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest

	// 待删除的用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLDAPUnitSetRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base&nbsp;DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 用户名映射属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 部门过滤

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
}

func (r *DescribeLDAPUnitSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLDAPUnitSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ACRuleResponse struct {

	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeleteDeviceAccountsRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDeviceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维任务执行结果

		TaskResult []*TaskResult `json:"TaskResult,omitempty" name:"TaskResult"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowGraphResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折线图

		Graph *GraphResult `json:"Graph,omitempty" name:"Graph"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShowGraphResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTaskDetail struct {

	// 运维任务主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务类型，1&nbsp;-&nbsp;手工执行任务，&nbsp;2&nbsp;-&nbsp;周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期性任务执行间隔，单位天

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 首次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 执行账户

	Account *string `json:"Account,omitempty" name:"Account"`
	// 任务超时时间，单位秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 任务命令

	Script *string `json:"Script,omitempty" name:"Script"`
	// 执行任务的资产信息

	Devices []*Device `json:"Devices,omitempty" name:"Devices"`
	// 下次任务执行时间

	NextTime *string `json:"NextTime,omitempty" name:"NextTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
}

type BindDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机账号私钥，最新长度128字节，最大长度8192字节

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 主机账号私钥口令，最大长度256字节

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

func (r *BindDeviceAccountPrivateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlRulesRequest struct {
	*tchttp.BaseRequest

	// 策略ID集合

	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

func (r *DeleteAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicesPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppAssetGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 应用资产id集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *DeleteAppAssetGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppAssetGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportAuditLogTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出任务Id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportAuditLogTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportAuditLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePushAccountTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 资产Id数组

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 资产账号home路径，默认&nbsp;为空不指定

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 资产账号所属组，默认为空不创建。

	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 需要创建的资产账号

	CreateAccount *string `json:"CreateAccount,omitempty" name:"CreateAccount"`
	// 认证方式。
	// 1&nbsp;密码
	// 2&nbsp;私钥

	AuthenticationMethod []*uint64 `json:"AuthenticationMethod,omitempty" name:"AuthenticationMethod"`
	// 需要创建的资产账号密码。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 需要创建的资产账号密钥。

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 需要创建的资产账号密钥密码。

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
	// 认证生成方式。
	// 1:自动生成相同密码/密
	// 2:自动生成不同密码/密
	// 3:手动指定相同密码/密钥

	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreatePushAccountTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePushAccountTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维工单总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维工单列表

		TicketSet []*Ticket `json:"TicketSet,omitempty" name:"TicketSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTicketSubmitFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTicketSubmitFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTicketSubmitFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSessionCommandResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令和所属会话

		CommandSessionSet []*SessionCommand `json:"CommandSessionSet,omitempty" name:"CommandSessionSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSessionCommandResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaInstance struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// ckafka实例Id

	CkafkaInstanceId *string `json:"CkafkaInstanceId,omitempty" name:"CkafkaInstanceId"`
	// ckafka实例名

	CkafkaInstanceName *string `json:"CkafkaInstanceName,omitempty" name:"CkafkaInstanceName"`
	// vpc&nbsp;Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// subnet&nbsp;Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type BindDeviceResourceRequest struct {
	*tchttp.BaseRequest

	// 资产ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 网络域ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *BindDeviceResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmdTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名，最大长度32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，\n分隔，最大长度32768字节

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
	// 标识cmdlist字段前端是否为base64加密传值.
	// 0:表示非base64加密
	// 1:表示是base64加密

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
}

func (r *CreateCmdTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmdTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsWithDeviceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 携带拥有该同名账号主机数量的账号信息列表

		AccountWithDeviceCountSet []*AccountWithDeviceCount `json:"AccountWithDeviceCountSet,omitempty" name:"AccountWithDeviceCountSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsWithDeviceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsWithDeviceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest

	// 每页条目数量，默认20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecuritySettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		SecuritySetting *SecuritySetting `json:"SecuritySetting,omitempty" name:"SecuritySetting"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecuritySettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecuritySettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReportTaskRequest struct {
	*tchttp.BaseRequest

	// 开关

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 报表任务&nbsp;id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 报表任务名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 周期cron&nbsp;表达式

	Period *string `json:"Period,omitempty" name:"Period"`
	// 报表任务说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 报表模板&nbsp;id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 单次报表-数据开始时间

	DataStartTime *string `json:"DataStartTime,omitempty" name:"DataStartTime"`
	// 数据时间范围，单位：小时

	DataTimeRange *uint64 `json:"DataTimeRange,omitempty" name:"DataTimeRange"`
	// 报表任务类型：1&nbsp;单次报表，2&nbsp;周期报表类型

	ReportType *uint64 `json:"ReportType,omitempty" name:"ReportType"`
	// 报表模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 创建用户name

	User *string `json:"User,omitempty" name:"User"`
	// 单次报表-数据结束时间

	DataEndTime *string `json:"DataEndTime,omitempty" name:"DataEndTime"`
	// 所属部门id。“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyReportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTrialGuideStepResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTrialGuideStepResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrialGuideStepResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUKeyBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUKeyBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUKeyBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmdTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmdTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseItemList struct {

	// 计费项id

	RuleKey *string `json:"RuleKey,omitempty" name:"RuleKey"`
	// 用量

	Usage *uint64 `json:"Usage,omitempty" name:"Usage"`
}

type ParamDeviceAccount struct {

	// 资产ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 账号名称，1-64，不能包含空白字符

	Account *string `json:"Account,omitempty" name:"Account"`
	// 账号密码，1-64字符

	Password *string `json:"Password,omitempty" name:"Password"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 账号密钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 账号密钥的密码

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
	// 容器账号凭据

	Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`
}

type ModifyAccessControlTemplateRuleOrderRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 规则id列表，按排序顺序传递。ruleId格式：acrule-xxx

	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *ModifyAccessControlTemplateRuleOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlTemplateRuleOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppAssetGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 应用资产类型集合

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 资产组ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// true-查询已绑定的成员，false-未绑定的成员

	Bound *bool `json:"Bound,omitempty" name:"Bound"`
	// 资产组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAppAssetGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppAssetGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportDeviceTaskRequest struct {
	*tchttp.BaseRequest

	// 部门&nbsp;Id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeExportDeviceTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportDeviceTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginSetting struct {

	// 登录会话超时，10分钟，20分钟，30分钟，默认20分钟

	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 连续密码错误次数，超过锁定账号，3-5

	LockThreshold *uint64 `json:"LockThreshold,omitempty" name:"LockThreshold"`
	// 账号锁定时长，10分钟，20分钟，30分钟

	LockTime *uint64 `json:"LockTime,omitempty" name:"LockTime"`
	// 用户多少天不活跃，账号自动锁定

	InactiveUserLock *uint64 `json:"InactiveUserLock,omitempty" name:"InactiveUserLock"`
}

type CanCreateTrialResourceRequest struct {
	*tchttp.BaseRequest
}

func (r *CanCreateTrialResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CanCreateTrialResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 过滤数组。支持的名称：RuleId&nbsp;规则ID&nbsp;，RuleName&nbsp;规则名称,
	// UnBindTemplateId&nbsp;未绑定模板id的规则

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordSettingRequest struct {
	*tchttp.BaseRequest

	// 密码最小长度，8-20，默认8

	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`
	// 密码复杂度，0不限制，1包含字母和数字，2至少包括大写字母、小写字母、数字、特殊符号，默认2

	Complexity *uint64 `json:"Complexity,omitempty" name:"Complexity"`
	// 密码有效期，0不限制，30天，90天，180天

	ValidTerm *uint64 `json:"ValidTerm,omitempty" name:"ValidTerm"`
	// 检查最近n次密码设置是否存在相同密码，2-10，默认5

	CheckHistory *uint64 `json:"CheckHistory,omitempty" name:"CheckHistory"`
}

func (r *ModifyPasswordSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunChangePwdTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunChangePwdTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Acl struct {

	// 访问权限ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 访问权限名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（目前含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（目前含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启文件传输上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，暂未启用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启文件传输下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，暂未启用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户列表

	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
	// 关联的用户组列表

	UserGroupSet []*Group `json:"UserGroupSet,omitempty" name:"UserGroupSet"`
	// 关联的资产列表

	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// 关联的资产组列表

	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitempty" name:"DeviceGroupSet"`
	// 关联的账号列表

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 关联的高危命令模板列表

	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`
	// 是否开启&nbsp;RDP&nbsp;磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启&nbsp;RDP&nbsp;磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启&nbsp;rz&nbsp;sz&nbsp;文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启&nbsp;rz&nbsp;sz&nbsp;文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启&nbsp;SFTP&nbsp;文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 访问权限状态，1&nbsp;-&nbsp;已生效，2&nbsp;-&nbsp;未生效，3&nbsp;-&nbsp;已过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 所属部门的信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 是否允许使用访问串，默认允许

	AllowAccessCredential *bool `json:"AllowAccessCredential,omitempty" name:"AllowAccessCredential"`
	// 关联的数据库高危命令列表

	ACTemplateSet []*ACTemplate `json:"ACTemplateSet,omitempty" name:"ACTemplateSet"`
	// 关联的白命令命令

	WhiteCmds []*string `json:"WhiteCmds,omitempty" name:"WhiteCmds"`
	// 是否允许记录键盘

	AllowKeyboardLogger *bool `json:"AllowKeyboardLogger,omitempty" name:"AllowKeyboardLogger"`
	// 关联的应用资产列表

	AppAssetSet []*AppAsset `json:"AppAssetSet,omitempty" name:"AppAssetSet"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 堡垒机资源列表

		ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 堡垒机资源数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproveTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApproveTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 策略名称，1~60字符，仅支持中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，RuleName不可重复

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 策略描述，0~100字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 客户端IP，最多10组

	ClientIPList []*string `json:"ClientIPList,omitempty" name:"ClientIPList"`
	// 数据库名称，最多10组

	DBName []*string `json:"DBName,omitempty" name:"DBName"`
	// 表名称，最多10组

	TableName []*string `json:"TableName,omitempty" name:"TableName"`
	// SQL命令，最多10组

	SQLCommand []*string `json:"SQLCommand,omitempty" name:"SQLCommand"`
	// 限制SELECT结果集行数大于等于此数。默认为0，不限制。

	SQLRowCount *uint64 `json:"SQLRowCount,omitempty" name:"SQLRowCount"`
	// 字段名，最多50组

	FieldName []*string `json:"FieldName,omitempty" name:"FieldName"`
	// 0：&nbsp;不限&nbsp;1：每天&nbsp;2：指定时间段

	ExecutePeriod *uint64 `json:"ExecutePeriod,omitempty" name:"ExecutePeriod"`
	// 开始日期，格式:：YYYY-MM-DD

	ExecuteBeginDate *string `json:"ExecuteBeginDate,omitempty" name:"ExecuteBeginDate"`
	// 开始时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteBeginTime *string `json:"ExecuteBeginTime,omitempty" name:"ExecuteBeginTime"`
	// 结束日期，格式:：YYYY-MM-DD

	ExecuteEndDate *string `json:"ExecuteEndDate,omitempty" name:"ExecuteEndDate"`
	// 结束时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteEndTime *string `json:"ExecuteEndTime,omitempty" name:"ExecuteEndTime"`
	// 0：阻断&nbsp;1：放行

	ExecuteMode *uint64 `json:"ExecuteMode,omitempty" name:"ExecuteMode"`
	// 字段名和字段敏感数据分类的逻辑条件。0:&nbsp;or;&nbsp;&nbsp;1:and

	FieldLogicGate *uint64 `json:"FieldLogicGate,omitempty" name:"FieldLogicGate"`
}

func (r *CreateAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncStatusRequest struct {
	*tchttp.BaseRequest

	// 查询的资产同步类型。1&nbsp;-主机资产，&nbsp;2&nbsp;-&nbsp;数据库资产

	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

func (r *DescribeAssetSyncStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAppAssetGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAppAssetGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAppAssetGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePushAccountTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePushAccountTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePushAccountTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportExternalDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产参数列表

	DeviceSet []*ExternalDevice `json:"DeviceSet,omitempty" name:"DeviceSet"`
}

func (r *ImportExternalDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportExternalDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskResultDetail struct {

	// 运维任务结果日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务类型&nbsp;1&nbsp;-&nbsp;手工任务，&nbsp;2&nbsp;-&nbsp;周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 运维任务周期（单位：天）

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务创建人员用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维任务创建人员姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 运维任务执行账户

	Account *string `json:"Account,omitempty" name:"Account"`
	// 运维任务超时时间,&nbsp;单位秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 运维任务命令

	Script *string `json:"Script,omitempty" name:"Script"`
	// 运维任务来源IP

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
}

type ApproveTicketRequest struct {
	*tchttp.BaseRequest

	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否允许使用访问串，默认允许

	AllowAccessCredential *bool `json:"AllowAccessCredential,omitempty" name:"AllowAccessCredential"`
	// 工单唯一id，数据库索引id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 申请时间戳，单位微妙

	ApplyTimeVersion *int64 `json:"ApplyTimeVersion,omitempty" name:"ApplyTimeVersion"`
	// 审批权限工单时，设置的有效期开始时间，默认为当前时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 是否开启rz&nbsp;sz文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 关联高危DB模板ID

	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitempty" name:"ACTemplateIdSet"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启&nbsp;SFTP&nbsp;文件下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 是否开启rdp磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 是否允许键盘记录

	AllowKeyboardLogger *bool `json:"AllowKeyboardLogger,omitempty" name:"AllowKeyboardLogger"`
	// 2-审批通过，4-审批驳回

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批意见

	ApproveDesc *string `json:"ApproveDesc,omitempty" name:"ApproveDesc"`
	// 是否开启rdp磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启&nbsp;SFTP&nbsp;文件上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，目前暂未使用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 文件传输下载大小限制（预留参数，目前暂未使用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否开启rz&nbsp;sz文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启&nbsp;SFTP&nbsp;文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 审批权限工单时，设置的有效期的结束时间，默认有限期为1天

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是否开启剪贴板文本（含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 关联的高危命令模板ID

	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`
}

func (r *ApproveTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApproveTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CanCreateTrialResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true:可以创建试用版，false:不能创建试用版

		Result *bool `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CanCreateTrialResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CanCreateTrialResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板列表

		Templates []*SystemTaskTemplate `json:"Templates,omitempty" name:"Templates"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateChangePwdTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChangePwdTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetLogDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetLogDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CanCreateShareClbResourceRequest struct {
	*tchttp.BaseRequest
}

func (r *CanCreateShareClbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CanCreateShareClbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件告警设置

		EventAlarmSetting *EventAlarmSetting `json:"EventAlarmSetting,omitempty" name:"EventAlarmSetting"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConnectDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConnectDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShareClbIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 共享clb

		ShareClbSet []*ShareClbIp `json:"ShareClbSet,omitempty" name:"ShareClbSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShareClbIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShareClbIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSyncFlagRequest struct {
	*tchttp.BaseRequest

	// 是否开启资产自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
}

func (r *ModifyAssetSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Departments struct {

	// 部门列表

	DepartmentSet []*Department `json:"DepartmentSet,omitempty" name:"DepartmentSet"`
	// 是否开启了部门管理&nbsp;true&nbsp;-&nbsp;已开启,&nbsp;false&nbsp;-&nbsp;未开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 当前操作UIN是否是根部门管理员

	RootManager *bool `json:"RootManager,omitempty" name:"RootManager"`
}

type DescribeTaskTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板类型：1账号推送&nbsp;&nbsp;2&nbsp;改密计划&nbsp;&nbsp;3&nbsp;审计报表

	TemplateType *uint64 `json:"TemplateType,omitempty" name:"TemplateType"`
}

func (r *DescribeTaskTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockUserRequest struct {
	*tchttp.BaseRequest

	// 用户id

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *LockUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportDeviceTaskRequest struct {
	*tchttp.BaseRequest

	// 任务&nbsp;Id列表

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteExportDeviceTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportDeviceTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExternalDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountRequest struct {
	*tchttp.BaseRequest

	// 地域码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 用户VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资产类型,1-Linux,&nbsp;2-Windows,3-MySQL,4-SqlServer&nbsp;不传-全部

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 是否绑定服务,1-已绑定,&nbsp;2-未绑定，&nbsp;不传-全部

	BindResource *uint64 `json:"BindResource,omitempty" name:"BindResource"`
}

func (r *DescribeDeviceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainInstallScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 脚本信息

		Script *string `json:"Script,omitempty" name:"Script"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainInstallScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainInstallScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorSessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *MonitorSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandSessionRequest struct {
	*tchttp.BaseRequest

	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 默认值为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 检索的目标命令，为模糊搜索

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// Cmd字段前端是否做bae64加密
	// 0：否&nbsp;1：是

	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`
	// 开始时间，不得早于当前时间的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SearchCommandSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 终端会话所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessWhiteListRule struct {

	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// IP或者网段

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CreateDepartmentRequest struct {
	*tchttp.BaseRequest

	// 部门名称，1&nbsp;-&nbsp;256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 上级部门的ID

	Parent *string `json:"Parent,omitempty" name:"Parent"`
	// 部门管理员的账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 部门管理员信息

	ManagerUsers []*DepartmentManagerUser `json:"ManagerUsers,omitempty" name:"ManagerUsers"`
}

func (r *CreateDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlTemplateRuleRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteAccessControlTemplateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlTemplateRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportUserTaskRequest struct {
	*tchttp.BaseRequest

	// 部门&nbsp;Id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeExportUserTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportUserTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidRequest struct {
	*tchttp.BaseRequest

	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 可填写路径名或文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 分页用偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否创建审计日志,通过查看按钮调用时为true,其他为false

	AuditLog *bool `json:"AuditLog,omitempty" name:"AuditLog"`
	// 1-已执行，&nbsp;&nbsp;2-被阻断

	AuditAction *int64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 以Protocol和Method为条件查询

	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitempty" name:"TypeFilters"`
}

func (r *SearchFileBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTasksRequest struct {
	*tchttp.BaseRequest

	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运维任务类型，1&nbsp;-&nbsp;手工执行任务，&nbsp;2&nbsp;-&nbsp;周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperationTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessTrackPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问日志id

		VisitId *string `json:"VisitId,omitempty" name:"VisitId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AccessTrackPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessTrackPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableInstanceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例类型列表

		InstanceTypeSet []*InstanceType `json:"InstanceTypeSet,omitempty" name:"InstanceTypeSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableInstanceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePushAccountTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情

		Tasks []*PushAccountTaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 任务总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushAccountTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushAccountTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListAutoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListAutoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppAssetRequest struct {
	*tchttp.BaseRequest

	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 资产组id列表

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 应用服务器账号id

	DeviceAccountId *uint64 `json:"DeviceAccountId,omitempty" name:"DeviceAccountId"`
	// 应用资产类型

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 应用资产url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 客户端工具类型

	ClientAppKind *string `json:"ClientAppKind,omitempty" name:"ClientAppKind"`
	// 应用资产id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 应用资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用服务器id

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 客户端工具路径

	ClientAppPath *string `json:"ClientAppPath,omitempty" name:"ClientAppPath"`
}

func (r *ModifyAppAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// 服务实例ID，如bh-saas-s3ed4r5e

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域编码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 服务实例规格信息

	SvArgs *string `json:"SvArgs,omitempty" name:"SvArgs"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 服务规格对应的资产数

	Nodes *uint64 `json:"Nodes,omitempty" name:"Nodes"`
	// 自动续费标记，0&nbsp;-&nbsp;表示默认状态，1&nbsp;-&nbsp;表示自动续费，2&nbsp;-&nbsp;表示明确不自动续费

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 资源状态，0&nbsp;-&nbsp;未初始化，1&nbsp;-&nbsp;正常，2&nbsp;-&nbsp;隔离，3&nbsp;-&nbsp;销毁，4&nbsp;-&nbsp;初始化失败，5&nbsp;-&nbsp;初始化中

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 服务实例名，如T-Sec-堡垒机（SaaS型）

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 定价模型ID

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 资源创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 商品码,&nbsp;p_cds_dasb

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码,&nbsp;sp_cds_dasb_bh_saas

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 是否过期，true-过期，false-未过期

	Expired *bool `json:"Expired,omitempty" name:"Expired"`
	// 是否开通，true-开通，false-未开通

	Deployed *bool `json:"Deployed,omitempty" name:"Deployed"`
	// 开通服务的&nbsp;VPC&nbsp;名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 开通服务的&nbsp;VPC&nbsp;对应的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 开通服务的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 开通服务的子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 开通服务的子网网段

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 外部IP

	PublicIpSet []*string `json:"PublicIpSet,omitempty" name:"PublicIpSet"`
	// 内部IP

	PrivateIpSet []*string `json:"PrivateIpSet,omitempty" name:"PrivateIpSet"`
	// 服务开通的高级功能列表，如:[DB]

	ModuleSet []*string `json:"ModuleSet,omitempty" name:"ModuleSet"`
	// 已使用的授权点数

	UsedNodes *uint64 `json:"UsedNodes,omitempty" name:"UsedNodes"`
	// 命令行运维端口号

	TUICmdPort *uint64 `json:"TUICmdPort,omitempty" name:"TUICmdPort"`
	// 字符直连端口号

	TUIDirectPort *uint64 `json:"TUIDirectPort,omitempty" name:"TUIDirectPort"`
	// 扩展点数

	ExtendPoints *uint64 `json:"ExtendPoints,omitempty" name:"ExtendPoints"`
	// 带宽扩展包个数(4M)

	PackageBandwidth *uint64 `json:"PackageBandwidth,omitempty" name:"PackageBandwidth"`
	// 授权点数扩展包个数(50点)

	PackageNode *uint64 `json:"PackageNode,omitempty" name:"PackageNode"`
	// 日志投递规格信息

	LogDeliveryArgs *string `json:"LogDeliveryArgs,omitempty" name:"LogDeliveryArgs"`
	// 是否共享clb，true-共享clb，false-独享clb

	ShareClb *bool `json:"ShareClb,omitempty" name:"ShareClb"`
	// 运营商类型，如果为空字符串表示内网类型

	LbVipIsp *string `json:"LbVipIsp,omitempty" name:"LbVipIsp"`
	// 部署模式&nbsp;默认0&nbsp;0-cvm&nbsp;1-tke

	DeployModel *uint64 `json:"DeployModel,omitempty" name:"DeployModel"`
	// 共享clb&nbsp;id

	OpenClbId *string `json:"OpenClbId,omitempty" name:"OpenClbId"`
	// 已经使用的网络域个数

	UsedDomainCount *uint64 `json:"UsedDomainCount,omitempty" name:"UsedDomainCount"`
	// 堡垒机实例对应的零信任实例id

	IOAResourceId *string `json:"IOAResourceId,omitempty" name:"IOAResourceId"`
	// 内网访问的ip

	IntranetPrivateIpSet []*string `json:"IntranetPrivateIpSet,omitempty" name:"IntranetPrivateIpSet"`
	// 堡垒机资源LB

	ClbSet []*Clb `json:"ClbSet,omitempty" name:"ClbSet"`
	// 开通内网访问vpc的网段

	IntranetVpcCidr *string `json:"IntranetVpcCidr,omitempty" name:"IntranetVpcCidr"`
	// 1&nbsp;默认值，客户单访问开启，0&nbsp;客户端访问关闭，2&nbsp;客户端访问开通中，3&nbsp;客户端访问关闭中

	ClientAccess *uint64 `json:"ClientAccess,omitempty" name:"ClientAccess"`
	// 零信任堡垒机用户扩展包个数。1个扩展包对应20个用户数

	PackageIOAUserCount *uint64 `json:"PackageIOAUserCount,omitempty" name:"PackageIOAUserCount"`
	// cdc集群id

	CdcClusterId *string `json:"CdcClusterId,omitempty" name:"CdcClusterId"`
	// 1&nbsp;默认值，外网访问开启，0&nbsp;外网访问关闭，2&nbsp;外网访问开通中，3&nbsp;外网访问关闭中

	ExternalAccess *uint64 `json:"ExternalAccess,omitempty" name:"ExternalAccess"`
	// 网络域个数

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 1&nbsp;默认值，web访问开启，0&nbsp;web访问关闭，2&nbsp;web访问开通中，3&nbsp;web访问关闭中

	WebAccess *uint64 `json:"WebAccess,omitempty" name:"WebAccess"`
	// 0默认值。0-免费版（试用版）ioa，1-付费版ioa

	IOAResource *uint64 `json:"IOAResource,omitempty" name:"IOAResource"`
	// &nbsp;零信任堡垒机带宽扩展包个数。一个扩展包表示4M带宽

	PackageIOABandwidth *uint64 `json:"PackageIOABandwidth,omitempty" name:"PackageIOABandwidth"`
	// 0&nbsp;非试用版，1&nbsp;试用版

	Trial *uint64 `json:"Trial,omitempty" name:"Trial"`
	// 日志投递规格信息

	LogDelivery *string `json:"LogDelivery,omitempty" name:"LogDelivery"`
	// 0&nbsp;默认值，非内网访问，1&nbsp;内网访问，2&nbsp;内网访问开通中，3&nbsp;内网访问关闭中

	IntranetAccess *uint64 `json:"IntranetAccess,omitempty" name:"IntranetAccess"`
	// 开通内网访问的vpc

	IntranetVpcId *string `json:"IntranetVpcId,omitempty" name:"IntranetVpcId"`
	// 开通内网访问的subnet

	IntranetSubnetId *string `json:"IntranetSubnetId,omitempty" name:"IntranetSubnetId"`
}

type UKeyUser struct {

	// UKEY&nbsp;ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// UKEY名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// UKEY容器名

	KeyContainerName *string `json:"KeyContainerName,omitempty" name:"KeyContainerName"`
	// UKEY证书

	KeyCertificate *string `json:"KeyCertificate,omitempty" name:"KeyCertificate"`
	// 待绑定用户的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。
	// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateAccessControlRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// acrule-******

		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的资产组ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListAutoStatusRequest struct {
	*tchttp.BaseRequest

	// true：放开自动添加IP；false：不放开自动添加IP

	AllowAuto *bool `json:"AllowAuto,omitempty" name:"AllowAuto"`
}

func (r *ModifyAccessWhiteListAutoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListAutoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunChangePwdTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 改密任务详情

	Details []*RunChangePwdTaskDetail `json:"Details,omitempty" name:"Details"`
}

func (r *RunChangePwdTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunChangePwdTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessEntryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维人员WEB访问入口

		AccessEntrySet []*string `json:"AccessEntrySet,omitempty" name:"AccessEntrySet"`
		// 运维人员WEB访问零信任堡垒机入口

		IOAAccessEntry *string `json:"IOAAccessEntry,omitempty" name:"IOAAccessEntry"`
		// 是否可以试用IOA,&nbsp;0-否&nbsp;1-是

		CanTrialIOA *uint64 `json:"CanTrialIOA,omitempty" name:"CanTrialIOA"`
		// 0&nbsp;内网访问未开通，1&nbsp;内网访问正在开通中，2&nbsp;已开通

		IntranetAccessStatus *uint64 `json:"IntranetAccessStatus,omitempty" name:"IntranetAccessStatus"`
		// 运维人员WEB访问内网入口

		IntranetAccessEntry *string `json:"IntranetAccessEntry,omitempty" name:"IntranetAccessEntry"`
		// 国密访问状态，0：未开通，1：已开通

		GMAccessStatus *uint64 `json:"GMAccessStatus,omitempty" name:"GMAccessStatus"`
		// 运维人员WEB访问国密入口

		GMAccessEntry *string `json:"GMAccessEntry,omitempty" name:"GMAccessEntry"`
		// 运维人员WEB访问内网入口

		IntranetAccessEntrySet []*IntranetAccessEntry `json:"IntranetAccessEntrySet,omitempty" name:"IntranetAccessEntrySet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmdTemplatesRequest struct {
	*tchttp.BaseRequest

	// 命令模板ID集合，非必需

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 命令模板名，模糊查询，最大长度64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令模板类型&nbsp;1-内置模板&nbsp;2-自定义模板

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCmdTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmdTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作日志列表

		OperationEventSet []*OperationEvent `json:"OperationEventSet,omitempty" name:"OperationEventSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationEventRequest struct {
	*tchttp.BaseRequest

	// 用户名，如果不包含其他条件时对user_name&nbsp;or&nbsp;real_name两个字段模糊查询

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，模糊查询

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 查询时间范围，起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询时间范围，结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 来源IP，模糊查询

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 操作类型，参考DescribeOperationType返回结果

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页记录数，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperationEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableIntranetAccessRequest struct {
	*tchttp.BaseRequest

	// 堡垒机id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DisableIntranetAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableIntranetAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDeviceAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportDeviceTaskRequest struct {
	*tchttp.BaseRequest

	// 资产IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 地域码集合

	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`
	// 操作系统类型,&nbsp;1&nbsp;-&nbsp;Linux,&nbsp;2&nbsp;-&nbsp;Windows,&nbsp;3&nbsp;-&nbsp;MySQL,&nbsp;4&nbsp;-&nbsp;SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 每页条目数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 任务类型：0-导出主机device，2-导出应用资产

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 过滤条件，资产绑定的堡垒机服务ID集合

	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`
	// 可提供按照多种类型过滤,&nbsp;1&nbsp;-&nbsp;Linux,&nbsp;2&nbsp;-&nbsp;Windows,&nbsp;3&nbsp;-&nbsp;MySQL,&nbsp;4&nbsp;-&nbsp;SQLServer

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 过滤条件，可按照部门ID进行过滤

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 过滤数组。支持的Name：
	// BindingStatus&nbsp;绑定状态

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 导出任务的部门&nbsp;id

	TaskDepartmentId *string `json:"TaskDepartmentId,omitempty" name:"TaskDepartmentId"`
	// 资产名或资产IP，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤条件，查询应用资产

	AppAssetParam *ExportAppAssetFilter `json:"AppAssetParam,omitempty" name:"AppAssetParam"`
}

func (r *CreateExportDeviceTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportDeviceTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 应用资产集合

		AppAssetSet []*AppAsset `json:"AppAssetSet,omitempty" name:"AppAssetSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketSubmitFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTicketSubmitFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketSubmitFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountWithDeviceCount struct {

	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 拥有该同名账号的主机数量

	DeviceCount *uint64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
}

type SearchSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 会话信息列表

		SessionSet []*SessionResult `json:"SessionSet,omitempty" name:"SessionSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSessionCommandRequest struct {
	*tchttp.BaseRequest

	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 默认值为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 检索的目标命令，为模糊搜索

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// Cmd字段前端是否做base64加密
	// 0：否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// 开始时间，不得早于当前时间的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SearchSessionCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidResult struct {

	// 文件操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 1-上传文件&nbsp;2-下载文件&nbsp;3-删除文件&nbsp;4-移动文件&nbsp;5-重命名文件&nbsp;6-新建文件夹&nbsp;7-移动文件夹&nbsp;8-重命名文件夹&nbsp;9-删除文件夹

	Method *int64 `json:"Method,omitempty" name:"Method"`
	// 文件传输协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// method为上传、下载、删除时文件在服务器上的位置,&nbsp;或重命名、移动文件前文件的位置

	FileCurr *string `json:"FileCurr,omitempty" name:"FileCurr"`
	// method为重命名、移动文件时代表移动后的新位置.其他情况为null

	FileNew *string `json:"FileNew,omitempty" name:"FileNew"`
	// method为上传文件、下载文件、删除文件时显示文件大小。其他情况为null

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 堡垒机拦截情况,&nbsp;1-已执行，&nbsp;&nbsp;2-被阻断

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 设备部门name

	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitempty" name:"DeviceDepartmentName"`
	// 资产账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 设备部门id

	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitempty" name:"DeviceDepartmentId"`
	// 签名值

	SignValue *string `json:"SignValue,omitempty" name:"SignValue"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 复核时间，当Action是3时，需有复核时间

	ConfirmTime *string `json:"ConfirmTime,omitempty" name:"ConfirmTime"`
	// 用户部门Id

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 用户部门name

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

type CreateUKeyBatchRequest struct {
	*tchttp.BaseRequest

	// UKey和待绑定用户集合

	UKeyUserSet []*UKeyUser `json:"UKeyUserSet,omitempty" name:"UKeyUserSet"`
	// 部门ID，用于判断用户是否有此部门的权限

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUKeyBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUKeyBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppAssetsRequest struct {
	*tchttp.BaseRequest

	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 过滤数组。支持的条件：BindingStatus、InstanceId、Url、DeviceName、DeviceInstanceId、ResourceId、DomainId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 应用资产id

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 应用资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 应用资产类型集合

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 有该资产访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
}

func (r *DescribeAppAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户姓名，最大长度20个字符，不能包含空格

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如:&nbsp;"+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 认证方式，0&nbsp;-&nbsp;本地，1&nbsp;-&nbsp;LDAP，2&nbsp;-&nbsp;OAuth&nbsp;不传则默认为0

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制，&nbsp;由0、1组成的字符串，长度168(7&nbsp;×&nbsp;24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时，&nbsp;0&nbsp;-&nbsp;代表不允许访问，1&nbsp;-&nbsp;代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 用户所属部门的ID，如1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployResourceRequest struct {
	*tchttp.BaseRequest

	// 需要开通服务的资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 需要开通服务的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 子网所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 需要开通服务的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 需要开通服务的子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 需要开通服务的子网网段

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 需要开通服务的VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 需要开通服务的VPC对应的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 需要开通实例所属的CDC集群ID

	CdcClusterId *string `json:"CdcClusterId,omitempty" name:"CdcClusterId"`
	// 需要开通服务的子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 开通堡垒机指定共享的clbId

	ShareClbId *string `json:"ShareClbId,omitempty" name:"ShareClbId"`
	// 需要开通内网访问的VPCId

	InternalVpcId *string `json:"InternalVpcId,omitempty" name:"InternalVpcId"`
	// 需要开通内网访问的VPC对应的网段

	InternalVpcCidrBlock *string `json:"InternalVpcCidrBlock,omitempty" name:"InternalVpcCidrBlock"`
	// 需要开通内网访问的子网ID

	InternalSubnetId *string `json:"InternalSubnetId,omitempty" name:"InternalSubnetId"`
}

func (r *DeployResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LDAPSetting struct {

	// 是否开启LDAP认证，false-不开启，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base&nbsp;DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 用户属性

	AttributeUser *string `json:"AttributeUser,omitempty" name:"AttributeUser"`
	// 用户名属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
	// 覆盖用户信息，false-不开启，true-开启

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
	// 同步周期，30～60000之间的整数

	SyncPeriod *uint64 `json:"SyncPeriod,omitempty" name:"SyncPeriod"`
	// 是否同步全部，false-不开启，true-开启

	SyncAll *bool `json:"SyncAll,omitempty" name:"SyncAll"`
	// 同步OU列表

	SyncUnitSet []*string `json:"SyncUnitSet,omitempty" name:"SyncUnitSet"`
	// 组织单元属性

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
	// 用户姓名属性

	AttributeRealName *string `json:"AttributeRealName,omitempty" name:"AttributeRealName"`
	// 手机号属性

	AttributePhone *string `json:"AttributePhone,omitempty" name:"AttributePhone"`
	// 邮箱属性

	AttributeEmail *string `json:"AttributeEmail,omitempty" name:"AttributeEmail"`
	// 请求LDAP服务的堡垒机实例

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeChangePwdTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情

		Details []*ChangePwdTaskDetail `json:"Details,omitempty" name:"Details"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChangePwdTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChangePwdTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需要添加到资产组的资产ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *AddDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandBySidRequest struct {
	*tchttp.BaseRequest

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 命令，可模糊搜索

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据拦截状态进行过滤

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
}

func (r *SearchCommandBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisitTrackPageRequest struct {
	*tchttp.BaseRequest

	// 访问id

	VisitId *string `json:"VisitId,omitempty" name:"VisitId"`
	// 访问时间

	VisitTime *string `json:"VisitTime,omitempty" name:"VisitTime"`
	// 按钮id

	ButtonId *uint64 `json:"ButtonId,omitempty" name:"ButtonId"`
}

func (r *VisitTrackPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisitTrackPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppAssetGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产组成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产组成员列表

		AppAssetSet []*AppAsset `json:"AppAssetSet,omitempty" name:"AppAssetSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppAssetGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppAssetGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessTimePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessTimePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessTimePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileSessionRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 主机内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 主机外网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 文件名

	File *string `json:"File,omitempty" name:"File"`
	// 对文件的动作，1为上传，2为下载，3为删除，4为重命名

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 只能取&nbsp;"SFTP"、"RDP"、"rz/sz"

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页中一页内的记录数，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchFileSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplayInformation struct {

	// 令牌

	Token *string `json:"Token,omitempty" name:"Token"`
	// 会话开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 回放链接

	Address *string `json:"Address,omitempty" name:"Address"`
	// 回放类型&nbsp;，默认0，&nbsp;1-rfb&nbsp;2-mp4&nbsp;3-ssh

	ReplayType *uint64 `json:"ReplayType,omitempty" name:"ReplayType"`
}

type DeleteAppAssetsRequest struct {
	*tchttp.BaseRequest

	// 删除的资产ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteAppAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessControlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChangePwdTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyChangePwdTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultDetailRequest struct {
	*tchttp.BaseRequest

	// 运维任务日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *SearchTaskResultDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDelivery struct {

	// 接入方式。
	// Domain&nbsp;域名接入
	// Internal&nbsp;支撑环境接入

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// ckafka实例Id。

	CkafkaInstanceId *string `json:"CkafkaInstanceId,omitempty" name:"CkafkaInstanceId"`
	// 接入地址。
	// 域名接入的情况下是域名和端口；支撑环境接入的情况下是虚拟地址和端口。

	Address *string `json:"Address,omitempty" name:"Address"`
	// ckafka实例所属VpcId。

	CkafkaUniqVpcId *string `json:"CkafkaUniqVpcId,omitempty" name:"CkafkaUniqVpcId"`
	// ckafka实例所属SubnetId。

	CkafkaUniqSubnetId *string `json:"CkafkaUniqSubnetId,omitempty" name:"CkafkaUniqSubnetId"`
	// ckafka实例所在地域。

	CkafkaRegion *string `json:"CkafkaRegion,omitempty" name:"CkafkaRegion"`
	// 日志投递详情。

	Details []*LogDeliveryDetail `json:"Details,omitempty" name:"Details"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组名，最大长度32字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户组所属部门的ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReconnectionSetting struct {

	// 重连次数

	ReconnectionMaxCount *uint64 `json:"ReconnectionMaxCount,omitempty" name:"ReconnectionMaxCount"`
	// true：可以重连，false：不可以重连

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type CreateAccessWhiteListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessWhiteListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessWhiteListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAuditLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审计日志

		AuditLogSet []*AuditLogResult `json:"AuditLogSet,omitempty" name:"AuditLogSet"`
		// 日志总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAuditLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetLogDeliveryRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例Id。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ResetLogDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetLogDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportAuditLogTaskRequest struct {
	*tchttp.BaseRequest

	// 账号，长度不超过64

	Account *string `json:"Account,omitempty" name:"Account"`
	// 来源IP，模糊查询

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 执行的命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 资产实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 来源IP，模糊查询,兼容不同接口中使用的名称不一致

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 登录入口：1-字符界面,2-图形界面，3-web页面,&nbsp;4-API

	Entry *uint64 `json:"Entry,omitempty" name:"Entry"`
	// 应用资产url

	AppAssetUrl *string `json:"AppAssetUrl,omitempty" name:"AppAssetUrl"`
	// 操作类型，参考DescribeOperationEvent返回结果

	EventKind *uint64 `json:"EventKind,omitempty" name:"EventKind"`
	// Cmd字段是前端传值是否进行base64.&nbsp;0:否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// 操作/登录结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 查询时间范围，结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 资产的公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 应用资产类型

	AppAssetKindSet []*uint64 `json:"AppAssetKindSet,omitempty" name:"AppAssetKindSet"`
	// 资产的内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作类型：1&nbsp;-&nbsp;文件上传，2&nbsp;-&nbsp;文件下载，3&nbsp;-&nbsp;文件删除，4&nbsp;-&nbsp;文件(夹)移动，5&nbsp;-&nbsp;文件(夹)重命名，6&nbsp;-&nbsp;新建文件夹，9&nbsp;-&nbsp;删除文件夹，10&nbsp;-&nbsp;剪贴板上传文件，11&nbsp;-&nbsp;剪贴板下载文件，12&nbsp;-&nbsp;剪贴板上行文本，13&nbsp;-&nbsp;剪贴板下行文本，14&nbsp;-&nbsp;剪贴板上行位图，15&nbsp;-&nbsp;剪贴板下行位图，16&nbsp;-&nbsp;剪贴板上行其它类型文件，17&nbsp;-&nbsp;剪贴板下行其它类型文件

	Method []*uint64 `json:"Method,omitempty" name:"Method"`
	// 可填写路径名或文件（夹）名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 用户名，如果不包含其他条件时对user_name&nbsp;or&nbsp;real_name两个字段模糊查询

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 导出日志类型&nbsp;1-字符会话&nbsp;2-图形会话&nbsp;3-文件传输会话&nbsp;4-数据库会话&nbsp;5-审计日志&nbsp;6-命令执行&nbsp;7-文件传输&nbsp;8-登录日志&nbsp;9-操作日志

	LogKind *uint64 `json:"LogKind,omitempty" name:"LogKind"`
	// 姓名，模糊查询

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 根据拦截状态进行过滤：1&nbsp;-&nbsp;已执行，2&nbsp;-&nbsp;被阻断,&nbsp;3-复核

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 状态，0为全部，1为活跃，2为结束，3为强制离线，4为其他错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 查询时间范围，起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *CreateExportAuditLogTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportAuditLogTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmdTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmdTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LeaveTrackPageRequest struct {
	*tchttp.BaseRequest

	// 访问id

	VisitId *string `json:"VisitId,omitempty" name:"VisitId"`
	// 离开时间

	LeaveTime *string `json:"LeaveTime,omitempty" name:"LeaveTime"`
}

func (r *LeaveTrackPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LeaveTrackPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaySessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回放所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaySessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件和所属会话

		FileSessionSet []*SessionFile `json:"FileSessionSet,omitempty" name:"FileSessionSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemTaskTemplate struct {

	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板类型：1&nbsp;账号推送&nbsp;&nbsp;2&nbsp;改密&nbsp;&nbsp;3&nbsp;审计报表

	TemplateType *uint64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板内容

	Script *string `json:"Script,omitempty" name:"Script"`
	// 超时时间

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 模板&nbsp;id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type UKeyUserDetail struct {

	// UKEY&nbsp;ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// UKEY名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 已绑定用户的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 已绑定用户的姓名

	UserRealName *string `json:"UserRealName,omitempty" name:"UserRealName"`
	// 已绑定用户的部门信息,仅返回DepartmentName

	Department *Department `json:"Department,omitempty" name:"Department"`
}

type DescribeAclsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问权限总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问权限列表

		AclSet []*Acl `json:"AclSet,omitempty" name:"AclSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowGraphRequest struct {
	*tchttp.BaseRequest

	// 开始时间，不得早于当前的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ShowGraphRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowGraphRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserBatchRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 认证方式，0-本地&nbsp;1-ldap&nbsp;2-oauth,不传则默认为0

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
}

func (r *ModifyUserBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionCommand struct {

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 账号

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 设备名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 命令列表

	Commands []*Command `json:"Commands,omitempty" name:"Commands"`
	// 记录数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 设备id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 设备所属的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
}

type DescribeReportTaskRequest struct {
	*tchttp.BaseRequest

	// 过滤数组。过滤数组。Name支持以下值:&nbsp;ReportName&nbsp;名称Description&nbsp;说明
	// ReportType&nbsp;&nbsp;类型
	// TemplateName&nbsp;模板名称
	// Status&nbsp;&nbsp;&nbsp;任务状态

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 部门&nbsp;id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 个数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportTaskHistory struct {

	// 模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 数据开始时间

	DataStartTime *string `json:"DataStartTime,omitempty" name:"DataStartTime"`
	// 数据结束时间

	DataEndTime *string `json:"DataEndTime,omitempty" name:"DataEndTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 记录&nbsp;id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 关联任务&nbsp;id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 报表类型：1&nbsp;单次&nbsp;2&nbsp;周期

	PeriodType *uint64 `json:"PeriodType,omitempty" name:"PeriodType"`
	// 模板&nbsp;id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 报表名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 报表描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建者&nbsp;name

	User *string `json:"User,omitempty" name:"User"`
}

type CreateChangePwdTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 资产id数组

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 修改的账户数组

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 改密方式。1：使用执行账号修改密码；2：修改自身密码

	ChangeMethod *int64 `json:"ChangeMethod,omitempty" name:"ChangeMethod"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 认证生成方式。&nbsp;1:自动生成相同密码&nbsp;2:自动生成不同密码&nbsp;3:手动指定相同密码

	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 手动指定密码需要传该字段

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密码限制长度

	PasswordLength *int64 `json:"PasswordLength,omitempty" name:"PasswordLength"`
	// 密码包含小写字母。1：包含

	SmallLetter *int64 `json:"SmallLetter,omitempty" name:"SmallLetter"`
	// 密码包含大写字母。1：包含

	BigLetter *int64 `json:"BigLetter,omitempty" name:"BigLetter"`
	// 密码包含数字。1：包含

	Digit *int64 `json:"Digit,omitempty" name:"Digit"`
	// 密码包含的特殊字符：base64编码

	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`
	// 改密完成通知。0：不通知&nbsp;
	// &nbsp;&nbsp;1：通知

	CompleteNotify *int64 `json:"CompleteNotify,omitempty" name:"CompleteNotify"`
	// 通知邮箱

	NotifyEmails []*string `json:"NotifyEmails,omitempty" name:"NotifyEmails"`
	// 加密压缩文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 所属部门id。“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 任务类型&nbsp;&nbsp;4-手工执行&nbsp;&nbsp;5-周期自动执行

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 执行周期，单位天

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 首次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
}

func (r *CreateChangePwdTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateChangePwdTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketsRequest struct {
	*tchttp.BaseRequest

	// 工单类型：0-运维工单，1-权限工单

	TicketType *string `json:"TicketType,omitempty" name:"TicketType"`
	// 过滤数组

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 状态：0待提交(草稿)，1待审批，2批准，3已执行，4驳回

	StatusSet []*uint64 `json:"StatusSet,omitempty" name:"StatusSet"`
}

func (r *DescribeTicketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionResult struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 会话大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 设备ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 设备名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 会话持续时长

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 该会话内命令数量&nbsp;，搜索图形会话时该字段无意义

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 该会话内高危命令数，搜索图形时该字段无意义

	DangerCount *uint64 `json:"DangerCount,omitempty" name:"DangerCount"`
	// 会话状态，如1会话活跃&nbsp;&nbsp;2会话结束&nbsp;&nbsp;3强制离线&nbsp;&nbsp;4其他错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 设备所属的地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 会话协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 应用资产url

	AppAssetUrl *string `json:"AppAssetUrl,omitempty" name:"AppAssetUrl"`
	// 会话资产类型

	DeviceKind *string `json:"DeviceKind,omitempty" name:"DeviceKind"`
	// K8S集群命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// K8S集群工作负载

	Workload *string `json:"Workload,omitempty" name:"Workload"`
	// 应用资产类型：1-web

	AppAssetKind *uint64 `json:"AppAssetKind,omitempty" name:"AppAssetKind"`
	// 回放类型&nbsp;默认0,&nbsp;1-rfb&nbsp;2-mp4&nbsp;3-ssh

	ReplayType *uint64 `json:"ReplayType,omitempty" name:"ReplayType"`
	// K8S集群容器名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 访问方式&nbsp;1-直链&nbsp;2-客户端&nbsp;3-web&nbsp;大部分情况下是2

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
}

type DeleteDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlTemplateRuleOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessControlTemplateRuleOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlTemplateRuleOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesSSLRequest struct {
	*tchttp.BaseRequest

	// 资产ID数组

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 是否启用SSL,&nbsp;仅Redis资产可用。1:启用&nbsp;0:禁用

	EnableSSL *int64 `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// SSL证书，EnableSSL时必填

	SSLCert *string `json:"SSLCert,omitempty" name:"SSLCert"`
	// SSL证书名称，EnableSSL时必填

	SSLCertName *string `json:"SSLCertName,omitempty" name:"SSLCertName"`
}

func (r *ModifyDevicesSSLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesSSLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAuditLogRequest struct {
	*tchttp.BaseRequest

	// 开始时间，不得早于当前时间的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchAuditLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAuditLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCountSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各种类型的资产总数

		DeviceCountSet []*DeviceCount `json:"DeviceCountSet,omitempty" name:"DeviceCountSet"`
		// 各种类型应用资产总数

		AppAssetCountSet []*DeviceCount `json:"AppAssetCountSet,omitempty" name:"AppAssetCountSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceCountSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceCountSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的用户组ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cvm struct {

	// cvm实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 外网IP

	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 内网IP

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// vpc&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type PushAccountTaskDetail struct {

	// 资产信息

	Device *Device `json:"Device,omitempty" name:"Device"`
	// 任务状态
	// 0&nbsp;未执行
	// 1&nbsp;执行中
	// 2&nbsp;成功
	// 3&nbsp;失败
	// 4&nbsp;超时

	Status *string `json:"Status,omitempty" name:"Status"`
	// 任务状态码
	// 0&nbsp;未执行
	// 1&nbsp;执行中
	// 2&nbsp;成功
	// 3&nbsp;失败
	// 4&nbsp;超时

	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// 网络域id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 网络域名称

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// ip&nbsp;，网段

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandRequest struct {
	*tchttp.BaseRequest

	// 搜索区间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间，不填默认为开始时间到现在为止

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产的公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产的内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 执行的命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// 根据拦截状态进行过滤：1&nbsp;-&nbsp;已执行，2&nbsp;-&nbsp;被阻断

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchCommandRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSessionRequest struct {
	*tchttp.BaseRequest

	// 内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 用户名，长度不超过20

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 账号，长度不超过64

	Account *string `json:"Account,omitempty" name:"Account"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 会话协议类型，只能是1、2、3或4&nbsp;对应关系为1-tui&nbsp;2-gui&nbsp;3-file&nbsp;4-数据库。若入参是Id，则非必传，否则为必传。

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 姓名，长度不超过20

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机名，长度不超过64

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 应用资产Url

	AppAssetUrl *string `json:"AppAssetUrl,omitempty" name:"AppAssetUrl"`
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话

	Id *string `json:"Id,omitempty" name:"Id"`
	// 应用资产类型,&nbsp;1-web

	AppAssetKindSet []*uint64 `json:"AppAssetKindSet,omitempty" name:"AppAssetKindSet"`
}

func (r *SearchSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUsersDepartmentRequest struct {
	*tchttp.BaseRequest

	// 用户记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 用户所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUsersDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsersDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务详情

		OperationTaskDetail *OperationTaskDetail `json:"OperationTaskDetail,omitempty" name:"OperationTaskDetail"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id&nbsp;bh-saas-xxxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型。标准版(standard)或者专业版(pro)

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源默认节点数。取值：10/20/50/100/200/500/1000

	ResourceNode *uint64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 带宽扩展包数量

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包数量

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
	// 日志投递

	LogDelivery *uint64 `json:"LogDelivery,omitempty" name:"LogDelivery"`
}

func (r *ModifyDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需删除的成员用户ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *DeleteUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 策略ID，策略ID&nbsp;和策略名称&nbsp;二选一

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略名称，策略ID&nbsp;和策略名称&nbsp;二选一

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 资产ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产名或资产IP，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 暂未使用

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 地域码集合

	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`
	// 操作系统类型,&nbsp;1&nbsp;-&nbsp;Linux,&nbsp;2&nbsp;-&nbsp;Windows,&nbsp;3&nbsp;-&nbsp;MySQL,&nbsp;4&nbsp;-&nbsp;SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 有该资产访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
	// 过滤条件，资产绑定的堡垒机服务ID集合

	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`
	// 可提供按照多种类型过滤,&nbsp;1&nbsp;-&nbsp;Linux,&nbsp;2&nbsp;-&nbsp;Windows,&nbsp;3&nbsp;-&nbsp;MySQL,&nbsp;4&nbsp;-&nbsp;SQLServer

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 资产是否包含托管账号。1，包含；0，不包含

	ManagedAccount *string `json:"ManagedAccount,omitempty" name:"ManagedAccount"`
	// 过滤条件，可按照部门ID进行过滤

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 过滤数组。支持的Name：
	// BindingStatus&nbsp;绑定状态

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDeviceTask struct {

	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务状态：0&nbsp;未开始，1&nbsp;进行中，2&nbsp;已完成，3&nbsp;已失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务剩余时间（秒）

	RemainTime *uint64 `json:"RemainTime,omitempty" name:"RemainTime"`
	// 文件cos地址

	CosAddr *string `json:"CosAddr,omitempty" name:"CosAddr"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 任务&nbsp;Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 任务进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 解压文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
}

type ModifyDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditLogResult struct {

	// 被审计会话的Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 审计者的编号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 审计动作发生的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 审计者的Ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 审计动作类型，1--回放、2--中断、3--监控

	Operation *int64 `json:"Operation,omitempty" name:"Operation"`
	// 被审计主机的Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 被审计主机的主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 被审计会话所属的类型，如字符会话

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 被审计主机的内部Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 被审计主机的外部Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 审计者的子账号

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 被审计资产url

	AppAssetUrl *string `json:"AppAssetUrl,omitempty" name:"AppAssetUrl"`
}

type DeleteDeviceGroupsRequest struct {
	*tchttp.BaseRequest

	// 待删除的资产组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDeviceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceAccountPrivateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUKeyRequest struct {
	*tchttp.BaseRequest

	// UKey&nbsp;ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 待绑定的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// UKey名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 部门ID，用于判断用户是否有此部门的权限

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchChangePwdTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 改密任务详细信息

		ChangePwdTask []*ChangePwdTaskInfoResult `json:"ChangePwdTask,omitempty" name:"ChangePwdTask"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchChangePwdTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchChangePwdTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementBySidRequest struct {
	*tchttp.BaseRequest

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// Statement字段前端是否做base64加密&nbsp;0：否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// SQL语句&nbsp;base64编码

	Statement *string `json:"Statement,omitempty" name:"Statement"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据拦截状态进行过滤

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
}

func (r *SearchStatementBySidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementBySidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepartmentManagerUser struct {

	// 管理员Id

	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
	// 管理员姓名

	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`
}

type DeployResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmdTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名，最长32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，\n分隔，最长32768字节

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
	// 命令模板ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// CmdList字段前端是否base64传值。
	// 0：否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// 命令模板类型&nbsp;1-内置模板&nbsp;2-自定义模板

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyCmdTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmdTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问控制规则总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表

		RuleSet []*AccessControlRuleDetail `json:"RuleSet,omitempty" name:"RuleSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用资产ID列表

		AssetIdSet []*uint64 `json:"AssetIdSet,omitempty" name:"AssetIdSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclsRequest struct {
	*tchttp.BaseRequest

	// 待删除的权限ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaInstanceListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCkafkaInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaInstanceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作类型：1&nbsp;-&nbsp;文件上传，2&nbsp;-&nbsp;文件下载，3&nbsp;-&nbsp;文件删除，4&nbsp;-&nbsp;文件(夹)移动，5&nbsp;-&nbsp;文件(夹)重命名，6&nbsp;-&nbsp;新建文件夹，9&nbsp;-&nbsp;删除文件夹

	Method []*uint64 `json:"Method,omitempty" name:"Method"`
	// 可填写路径名或文件（夹）名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 1-已执行，&nbsp;&nbsp;2-被阻断

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppAssetGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppAssetGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppAssetGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ACTemplate struct {

	// 模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模板名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建部门的ID，如：“1.2.3”

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetUserPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunPushAccountTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 资产Id数组，为空的情况下所有资产都执行任务

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 所属部门ID，如：“1.2.3”。

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *RunPushAccountTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunPushAccountTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDepartmentRequest struct {
	*tchttp.BaseRequest

	// 待删除的部门ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLDAPSyncFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *SetLDAPSyncFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLDAPSyncFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产组成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产组成员列表

		DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产组总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产组列表

		GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLDAPUnitSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ou&nbsp;列表

		UnitSet []*string `json:"UnitSet,omitempty" name:"UnitSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLDAPUnitSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLDAPUnitSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTaskResultDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务执行结果详情

		TaskResultDetail *TaskResultDetail `json:"TaskResultDetail,omitempty" name:"TaskResultDetail"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTaskResultDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// InstanceId集合

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户组所属的部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据库操作列表

		StatementSet []*DatabaseStatement `json:"StatementSet,omitempty" name:"StatementSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStatementBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSubtaskResultByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维子任务执行结果

		SubtaskResult []*SubtaskResult `json:"SubtaskResult,omitempty" name:"SubtaskResult"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSubtaskResultByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSubtaskResultByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Map struct {

	// 对应map的key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应map的value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeSystemTaskStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统任务统计信息

		SystemTaskStatistics *SystemTaskStatistics `json:"SystemTaskStatistics,omitempty" name:"SystemTaskStatistics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemTaskStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportExternalDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产ID列表

		DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportExternalDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportExternalDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmdTemplatesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteCmdTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmdTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LockUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LockUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppAsset struct {

	// 资产所属部门

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 应用服务器id

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 托管状态。0-未托管，1-已托管

	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`
	// 应用服务器实例id

	DeviceInstanceId *string `json:"DeviceInstanceId,omitempty" name:"DeviceInstanceId"`
	// 网络域名称

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// 应用服务器名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 堡垒机实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 堡垒机实例信息

	Resource *Resource `json:"Resource,omitempty" name:"Resource"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 客户端工具类型

	ClientAppKind *string `json:"ClientAppKind,omitempty" name:"ClientAppKind"`
	// 应用资产url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 应用服务器账号名称

	DeviceAccountName *string `json:"DeviceAccountName,omitempty" name:"DeviceAccountName"`
	// 资产组信息

	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
	// 应用资产id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 应用服务器账号id

	DeviceAccountId *uint64 `json:"DeviceAccountId,omitempty" name:"DeviceAccountId"`
	// 应用资产类型。1-web应用

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 客户端工具路径

	ClientAppPath *string `json:"ClientAppPath,omitempty" name:"ClientAppPath"`
	// 网络域id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type ModifyAccessTimePolicyRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 生效时间段,&nbsp;0、1组成的字符串，长度168(7*24),&nbsp;代表该用户的生效时间.&nbsp;0&nbsp;-&nbsp;未生效，1&nbsp;-&nbsp;生效

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
}

func (r *ModifyAccessTimePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessTimePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangePwdTaskInfo struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 任务id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 所属部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 改密方式。1：使用执行账号。2：修改自身密码

	ChangeMethod *uint64 `json:"ChangeMethod,omitempty" name:"ChangeMethod"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 密码生成策略

	AuthGenerationStrategy *uint64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 密码长度

	PasswordLength *uint64 `json:"PasswordLength,omitempty" name:"PasswordLength"`
	// 包含小写字母

	SmallLetter *uint64 `json:"SmallLetter,omitempty" name:"SmallLetter"`
	// 包含大写字母

	BigLetter *uint64 `json:"BigLetter,omitempty" name:"BigLetter"`
	// 包含数字

	Digit *uint64 `json:"Digit,omitempty" name:"Digit"`
	// 包含的特殊字符，base64

	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`
	// 改密完成通知。0-通知，1-不通知

	CompleteNotify *uint64 `json:"CompleteNotify,omitempty" name:"CompleteNotify"`
	// 通知人邮箱

	NotifyEmails []*string `json:"NotifyEmails,omitempty" name:"NotifyEmails"`
	// 加密附件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 需要改密的账户

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 需要改密的主机

	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// 任务类型：4手动，5自动

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 首次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 下次执行时间

	NextTime *string `json:"NextTime,omitempty" name:"NextTime"`
	// 上次执行时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
}

type CreateAssetSyncJobRequest struct {
	*tchttp.BaseRequest

	// 同步资产类别，1&nbsp;-&nbsp;主机资产,&nbsp;2&nbsp;-&nbsp;数据库资产

	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

func (r *CreateAssetSyncJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetSyncJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableInstanceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableInstanceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableInstanceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产组ID集合，传Id，IdSet不生效。

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产名或资产IP，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// true&nbsp;-&nbsp;查询已在该资产组的资产，false&nbsp;-&nbsp;查询未在该资产组的资产

	Bound *bool `json:"Bound,omitempty" name:"Bound"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 资产类型，1&nbsp;-&nbsp;Linux，2&nbsp;-&nbsp;Windows，3&nbsp;-&nbsp;MySQL，4&nbsp;-&nbsp;SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 资产类型集合，1&nbsp;-&nbsp;Linux，2&nbsp;-&nbsp;Windows，3&nbsp;-&nbsp;MySQL，4&nbsp;-&nbsp;SQLServer

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 需要删除的资产ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *DeleteDeviceGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表cos&nbsp;存储链接

		ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ViewReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ViewReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ACTemplateResponse struct {

	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模版名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 模版描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 模版关联的acl数量

	AclCount *uint64 `json:"AclCount,omitempty" name:"AclCount"`
	// 模版关联的rule数量

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
}

type Device struct {

	// 资产ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID，对应CVM、CDB等实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 地域编码

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 资产类型&nbsp;1&nbsp;-&nbsp;Linux,&nbsp;2&nbsp;-&nbsp;Windows,&nbsp;3&nbsp;-&nbsp;MySQL,&nbsp;4&nbsp;-&nbsp;SQLServer

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 所属资产组列表

	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
	// 资产绑定的账号数

	AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 堡垒机服务信息，注意没有绑定服务时为null

	Resource *Resource `json:"Resource,omitempty" name:"Resource"`
	// 资产所属部门

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 数据库资产的多节点

	IpPortSet []*string `json:"IpPortSet,omitempty" name:"IpPortSet"`
	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 已上传的SSL证书名称

	SSLCertName *string `json:"SSLCertName,omitempty" name:"SSLCertName"`
	// IOA侧的资源ID

	IOAId *int64 `json:"IOAId,omitempty" name:"IOAId"`
	// K8S集群工作负载

	Workload *string `json:"Workload,omitempty" name:"Workload"`
	// K8S集群pod已同步数量

	SyncPodCount *uint64 `json:"SyncPodCount,omitempty" name:"SyncPodCount"`
	// K8S集群pod总数量

	TotalPodCount *uint64 `json:"TotalPodCount,omitempty" name:"TotalPodCount"`
	// 网络域名称

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// 是否启用SSL，仅支持Redis资产。0：禁用&nbsp;1：启用

	EnableSSL *int64 `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// K8S集群托管维度。1-集群，2-命名空间，3-工作负载

	ManageDimension *uint64 `json:"ManageDimension,omitempty" name:"ManageDimension"`
	// K8S集群托管账号id

	ManageAccountId *uint64 `json:"ManageAccountId,omitempty" name:"ManageAccountId"`
	// K8S集群命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DeleteUserGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginOpserverResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 登录token

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LoginOpserverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginOpserverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessInfo struct {

	// 地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 账号

	User *string `json:"User,omitempty" name:"User"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 唤起链接｜wss链接

	AccessURL *string `json:"AccessURL,omitempty" name:"AccessURL"`
}

type ModifyUsersDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsersDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUsersDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessDeviceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 私钥密码

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// &nbsp;协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *AccessDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListRuleRequest struct {
	*tchttp.BaseRequest

	// 白名单规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// ip或网段信息，如10.10.10.1或10.10.10.0/24，最大长度40字节

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息，最大长度64字符。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccessWhiteListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchEventRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 高危事件类型

	EventStyle *int64 `json:"EventStyle,omitempty" name:"EventStyle"`
	// 开始时间，若传该参数，则不得早于当前的180天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量，默认为5，最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventResult struct {

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 地域信息

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 主机id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 高危命令或文件传输风险操作的文件名

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 高危命令执行时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 事件类型，如shell命令、文件传输等

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 各Type所属子类型，如:Type为文件传输时，此字段代表文件操作类型

	Method *int64 `json:"Method,omitempty" name:"Method"`
}

type CmdTemplate struct {

	// 高危命令模板ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 高危命令模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命令列表，命令之间用换行符（"\n"）分隔

	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
	// 命令模板类型&nbsp;1-内置&nbsp;2-自定义

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type PasswordSetting struct {

	// 密码最小长度，8-20，默认8。

	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`
	// 密码复杂度，0不限制，1包含字母和数字，2至少包括大写字母、小写字母、数字、特殊符号，默认2。

	Complexity *uint64 `json:"Complexity,omitempty" name:"Complexity"`
	// 密码有效期，0不限制，30天，90天，180天。

	ValidTerm *uint64 `json:"ValidTerm,omitempty" name:"ValidTerm"`
	// 检查最近n次密码设置是否存在相同密码，2-10，默认5。

	CheckHistory *uint64 `json:"CheckHistory,omitempty" name:"CheckHistory"`
}

type CreateExportDeviceTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务&nbsp;ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportDeviceTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportDeviceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账号信息列表

		DeviceAccountSet []*DeviceAccount `json:"DeviceAccountSet,omitempty" name:"DeviceAccountSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 堡垒机资源列表

		Item []*Resource `json:"Item,omitempty" name:"Item"`
		// 堡垒机实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetUserPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetUserPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetUserPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckLDAPConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLDAPConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLDAPConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductList struct {

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// license对象组列表

	LicenseItemList []*LicenseItemList `json:"LicenseItemList,omitempty" name:"LicenseItemList"`
}

type DescribeAccessEntryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccessEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessEntryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchPushAccountTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账号推送任务信息

		PushAccountTask []*PushAccountTaskInfoResult `json:"PushAccountTask,omitempty" name:"PushAccountTask"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchPushAccountTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchPushAccountTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangePwdTaskInfoResult struct {

	// 运维任务结果ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 改密方式

	ChangeMethod *uint64 `json:"ChangeMethod,omitempty" name:"ChangeMethod"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 密码生成策略

	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 密码长度

	PasswordLength *int64 `json:"PasswordLength,omitempty" name:"PasswordLength"`
	// 包含小写字母

	SmallLetter *uint64 `json:"SmallLetter,omitempty" name:"SmallLetter"`
	// 包含大写字母

	BigLetter *uint64 `json:"BigLetter,omitempty" name:"BigLetter"`
	// 包含数字：1包含

	Digit *uint64 `json:"Digit,omitempty" name:"Digit"`
	// 包含字符

	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`
	// 改密完成通知

	CompleteNotify *int64 `json:"CompleteNotify,omitempty" name:"CompleteNotify"`
	// 改密通知邮箱

	NotifyEmails []*string `json:"NotifyEmails,omitempty" name:"NotifyEmails"`
	// 加密文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 部门ID

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 部门名称

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
}

type CanCreateShareClbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true:能创建共享clb的堡垒机，false:不能创建共享clb的堡垒机

		Result *bool `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CanCreateShareClbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CanCreateShareClbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePushAccountTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePushAccountTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePushAccountTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChangePwdTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情

		Tasks []*ChangePwdTaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 任务总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChangePwdTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushAccountTaskInfo struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 需要创建的资产账号

	CreateAccount *string `json:"CreateAccount,omitempty" name:"CreateAccount"`
	// 认证方式
	// 1&nbsp;密码
	// 2&nbsp;私钥

	AuthenticationMethod []*uint64 `json:"AuthenticationMethod,omitempty" name:"AuthenticationMethod"`
	// 认证生成方式

	AuthGenerationStrategy *uint64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 资产账号home路径，默认&nbsp;为空不指定

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 资产账号所属组，默认为空不创建

	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet"`
	// 所属部门的信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 资产数量

	DeviceCount *int64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
	// 资产列表

	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// shell所在路径

	ShellPath *string `json:"ShellPath,omitempty" name:"ShellPath"`
}

type SystemTaskStatistic struct {

	// 任务数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 已使用任务数量

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 运行中的任务数量

	Running *int64 `json:"Running,omitempty" name:"Running"`
	// 最近一周执行任务数量

	LastWeekRunningTask *int64 `json:"LastWeekRunningTask,omitempty" name:"LastWeekRunningTask"`
	// 最近一个月执行任务数量

	LastMonthRunningTask *int64 `json:"LastMonthRunningTask,omitempty" name:"LastMonthRunningTask"`
}

type ConnectDomainRequest struct {
	*tchttp.BaseRequest

	// 网络域id

	DomainIdSet []*string `json:"DomainIdSet,omitempty" name:"DomainIdSet"`
}

func (r *ConnectDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConnectDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportTaskHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 报表记录列表

		ReportHistory []*ReportTaskHistory `json:"ReportHistory,omitempty" name:"ReportHistory"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReportTaskHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTaskHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOAuthSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOAuthSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOAuthSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisconnectDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisconnectDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisconnectDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReportTaskHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReportTaskHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportTaskHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportUserTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		Tasks []*ExportDeviceTask `json:"Tasks,omitempty" name:"Tasks"`
		// 任务数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportUserTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportUserTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlTemplateRulesRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeAccessControlTemplateRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlTemplateRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesSSLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicesSSLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Department struct {

	// 部门ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 部门名称，1&nbsp;-&nbsp;256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部门管理员账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 管理员用户

	ManagerUsers []*DepartmentManagerUser `json:"ManagerUsers,omitempty" name:"ManagerUsers"`
	// ioa分组id

	IOAId *uint64 `json:"IOAId,omitempty" name:"IOAId"`
}

type ModifyDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产所属组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 资产所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPushAccountTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 资产Id数组

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyPushAccountTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPushAccountTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PwdLog struct {

	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 订单关联标识

	IncKey *string `json:"IncKey,omitempty" name:"IncKey"`
}

type ModifyDevicesDepartmentRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyDevicesDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchKeyboardLoggerRequest struct {
	*tchttp.BaseRequest

	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 会话ID

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 分页用偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchKeyboardLoggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchKeyboardLoggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTask struct {

	// 运维任务主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维人员姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 任务类型，1&nbsp;-&nbsp;手工执行任务，&nbsp;2&nbsp;-&nbsp;周期性任务

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期性任务执行间隔，单位天

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 执行账户

	NextTime *string `json:"NextTime,omitempty" name:"NextTime"`
	// 下一次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
}

type BindAppAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAppAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAppAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id&nbsp;bh-saas-xxxx

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DestroyDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginSettingRequest struct {
	*tchttp.BaseRequest

	// 登录会话超时，10分钟，20分钟，30分钟,&nbsp;60分钟，默认30分钟

	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 连续密码错误次数，超过锁定账号，3-5

	LockThreshold *uint64 `json:"LockThreshold,omitempty" name:"LockThreshold"`
	// 账号锁定时长，10分钟，20分钟，30分钟

	LockTime *uint64 `json:"LockTime,omitempty" name:"LockTime"`
}

func (r *ModifyLoginSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessWhiteListRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问白名单规则列表

		AccessWhiteListRuleSet []*AccessWhiteListRule `json:"AccessWhiteListRuleSet,omitempty" name:"AccessWhiteListRuleSet"`
		// 是否放开全部来源IP，如果为true，TotalCount为0，AccessWhiteListRuleSet为空

		AllowAny *bool `json:"AllowAny,omitempty" name:"AllowAny"`
		// 是否开启自动添加来源IP,&nbsp;如果为true,&nbsp;在开启访问白名单的情况下将自动添加来源IP

		AllowAuto *bool `json:"AllowAuto,omitempty" name:"AllowAuto"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessWhiteListRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessWhiteListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产同步结果

		Status *AssetSyncStatus `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChangePwdTaskRequest struct {
	*tchttp.BaseRequest

	// 过滤数组。过滤数组。Name支持以下值:&nbsp;OperationId&nbsp;任务ID&nbsp;TaskName&nbsp;任务名

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移量，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeChangePwdTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChangePwdTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Group struct {

	// 组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 个数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DeleteAccessWhiteListRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessWhiteListRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessWhiteListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppAssetsDepartmentRequest struct {
	*tchttp.BaseRequest

	// 应用资产id集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 所属部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyAppAssetsDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppAssetsDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionFile struct {

	// 会话Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 设备ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 主机内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 主机外网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 来源Ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 属于该会话的文件集合

	Files []*FileInformation `json:"Files,omitempty" name:"Files"`
}

type AccessTrackPageRequest struct {
	*tchttp.BaseRequest

	// 进入埋点页面时间

	AccessTime *string `json:"AccessTime,omitempty" name:"AccessTime"`
	// 页面名称：Buy

	PageName *string `json:"PageName,omitempty" name:"PageName"`
}

func (r *AccessTrackPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AccessTrackPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDeliveryDetail struct {

	// 日志类型。目前取值：
	// CommandLog&nbsp;命令日志
	// FileLog&nbsp;文件日志

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// Ckafka&nbsp;topic。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 是否启用。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ReplaySessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *ReplaySessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaySessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnlockUserRequest struct {
	*tchttp.BaseRequest

	// 用户id

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *UnlockUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaInstanceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// ckafka实例信息

		CkafkaInstances []*CkafkaInstance `json:"CkafkaInstances,omitempty" name:"CkafkaInstances"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户组列表

		GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnlockUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnlockUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnlockUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件操作列表

		Files []*SearchFileResult `json:"Files,omitempty" name:"Files"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportDeviceTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportDeviceTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportDeviceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportUserTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportUserTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportUserTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlTemplateRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则对象列表

		RuleSet []*ACRuleResponse `json:"RuleSet,omitempty" name:"RuleSet"`
		// total行数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlTemplateRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlTemplateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseStatement struct {

	// 数据库操作

	Statement *string `json:"Statement,omitempty" name:"Statement"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 执行情况，1--允许，2--拒绝，3--确认

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 数据库类型&nbsp;1-MySQL&nbsp;2-PostgreSQL&nbsp;&nbsp;3-&nbsp;MongoDB

	DBType *int64 `json:"DBType,omitempty" name:"DBType"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户真实姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type OperationEvent struct {

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 操作时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 来源IP

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 操作类型

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 具体操作内容

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操作结果，1-成功，2-失败

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 签名值

	SignValue *string `json:"SignValue,omitempty" name:"SignValue"`
}

type CreateUKeyRequest struct {
	*tchttp.BaseRequest

	// UKey&nbsp;ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// UKey名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// UKey容器名

	KeyContainerName *string `json:"KeyContainerName,omitempty" name:"KeyContainerName"`
	// UKey证书

	KeyCertificate *string `json:"KeyCertificate,omitempty" name:"KeyCertificate"`
	// 待绑定的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 部门ID，用于判断用户是否有此部门的权限

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 活跃用户、主机、高危用户和在线用户

		Top *TopResult `json:"Top,omitempty" name:"Top"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShowTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令和所属会话

		CommandSessionSet []*SessionCommand `json:"CommandSessionSet,omitempty" name:"CommandSessionSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAuditLogTask struct {

	// 任务导出的日志总记录数

	TotalRecords *uint64 `json:"TotalRecords,omitempty" name:"TotalRecords"`
	// 导出任务类型

	LogKind *uint64 `json:"LogKind,omitempty" name:"LogKind"`
	// 任务状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// cos地址

	CosAddr *string `json:"CosAddr,omitempty" name:"CosAddr"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 导出的日志数据开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 导出的日志数据结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 任务剩余时间（秒）

	RemainTime *uint64 `json:"RemainTime,omitempty" name:"RemainTime"`
}

type InstanceType struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

type DeleteAccessControlTemplateRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessControlTemplateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlTemplateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportTask struct {

	// 报表模板&nbsp;id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 数据时间范围，单位&nbsp;小时（周期报表）

	DataTimeRange *uint64 `json:"DataTimeRange,omitempty" name:"DataTimeRange"`
	// 任务状态:&nbsp;0&nbsp;待生成&nbsp;1&nbsp;已生成&nbsp;&nbsp;2已暂停&nbsp;&nbsp;3已失败&nbsp;4&nbsp;生成中

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 报表任务&nbsp;id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 周期&nbsp;cron&nbsp;表达式

	Period *string `json:"Period,omitempty" name:"Period"`
	// 创建者名称

	User *string `json:"User,omitempty" name:"User"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 上次报告触发时间

	LastReportTime *string `json:"LastReportTime,omitempty" name:"LastReportTime"`
	// 报表任务类型：1&nbsp;单次报表&nbsp;2&nbsp;周期报表

	ReportType *uint64 `json:"ReportType,omitempty" name:"ReportType"`
	// 报表名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 开关：0&nbsp;关闭&nbsp;&nbsp;1&nbsp;打开

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 部门信息

	Department *Department `json:"Department,omitempty" name:"Department"`
	// 数据开始时间（单次报表）

	DataStartTime *string `json:"DataStartTime,omitempty" name:"DataStartTime"`
	// 已生成报表数量

	ReportCount *uint64 `json:"ReportCount,omitempty" name:"ReportCount"`
	// 数据结束时间（单次报表）

	DataEndTime *string `json:"DataEndTime,omitempty" name:"DataEndTime"`
	// 下次触发时间（周期报表）

	NextReportTime *string `json:"NextReportTime,omitempty" name:"NextReportTime"`
	// 报表任务名称

	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
	// 报表说明

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeCmdTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命令模板列表

		CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmdTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessControlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAppAssetFilter struct {

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 每页条目数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 应用资产ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 应用资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 有该主机访问权限的用户ID集合

	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 应用资产类型

	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
	// 所属部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤数组。支持的条件：BindingStatus、InstanceId、Url、DeviceName、DeviceInstanceId、ResourceId、DomainId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type CreateAssetSyncJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetSyncJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReportTaskRequest struct {
	*tchttp.BaseRequest

	// 报表任务&nbsp;idset

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteReportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品license用量列表

		ProductList []*ProductList `json:"ProductList,omitempty" name:"ProductList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReportTaskHistoryRequest struct {
	*tchttp.BaseRequest

	// 过滤数组。过滤数组。Name支持以下值:&nbsp;ReportName&nbsp;名称Description&nbsp;说明
	// ReportType&nbsp;&nbsp;类型
	// TemplateName&nbsp;模板名称

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 部门&nbsp;id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页个数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序：desc、asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeReportTaskHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReportTaskHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOAuthSettingRequest struct {
	*tchttp.BaseRequest

	// OAuth认证方式。

	AuthMethod *string `json:"AuthMethod,omitempty" name:"AuthMethod"`
	// OAuth认证客户端Id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// OAuth认证客户端密钥

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// 获取OAuth认证授权码URL

	CodeUrl *string `json:"CodeUrl,omitempty" name:"CodeUrl"`
	// 获取OAuth令牌URL

	TokenUrl *string `json:"TokenUrl,omitempty" name:"TokenUrl"`
	// 获取OAuth用户信息URL

	UserInfoUrl *string `json:"UserInfoUrl,omitempty" name:"UserInfoUrl"`
	// 使用Okta认证时指定范围。为空时默认使用&nbsp;openid、profile、email。

	Scopes []*string `json:"Scopes,omitempty" name:"Scopes"`
	// 是否开启OAuth认证，false-不开启，true-开启。

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyOAuthSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOAuthSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileResult struct {

	// 文件传输的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 操作结果：1&nbsp;-&nbsp;已执行，2&nbsp;-&nbsp;已阻断

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 操作类型：1&nbsp;-&nbsp;文件上传，2&nbsp;-&nbsp;文件下载，3&nbsp;-&nbsp;文件删除，4&nbsp;-&nbsp;文件(夹)移动，5&nbsp;-&nbsp;文件(夹)重命名，6&nbsp;-&nbsp;新建文件夹，9&nbsp;-&nbsp;删除文件夹

	Method *uint64 `json:"Method,omitempty" name:"Method"`
	// 下载的文件（夹）路径及名称

	FileCurr *string `json:"FileCurr,omitempty" name:"FileCurr"`
	// 上传或新建文件（夹）路径及名称

	FileNew *string `json:"FileNew,omitempty" name:"FileNew"`
	// 会话id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 复核时间

	ConfirmTime *string `json:"ConfirmTime,omitempty" name:"ConfirmTime"`
	// 用户部门id

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 设备部门id

	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitempty" name:"DeviceDepartmentId"`
	// 来源id

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 设备部门name

	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitempty" name:"DeviceDepartmentName"`
	// 签名值

	SignValue *string `json:"SignValue,omitempty" name:"SignValue"`
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 文件大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 用户部门name

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
}

type SearchTaskResultRequest struct {
	*tchttp.BaseRequest

	// 搜索区间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户名，长度不超过20

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名，长度不超过20

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 任务类型
	// 1&nbsp;手工运维任务
	// 2&nbsp;定时任务
	// 3&nbsp;账号推送任务

	TaskType []*uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDeviceAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// 地域码集合

	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`
}

func (r *DescribeInstanceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTicketSubmitFlagRequest struct {
	*tchttp.BaseRequest

	// false:未开启自动提交&nbsp;true:自动提交

	TicketSubmitFlag *string `json:"TicketSubmitFlag,omitempty" name:"TicketSubmitFlag"`
}

func (r *ModifyTicketSubmitFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTicketSubmitFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDepartmentRequest struct {
	*tchttp.BaseRequest

	// 部门ID，如：“1.2.3”

	Id *string `json:"Id,omitempty" name:"Id"`
	// 部门名称，1&nbsp;-&nbsp;256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部门管理员的账号ID

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 部门管理员信息

	ManagerUsers []*DepartmentManagerUser `json:"ManagerUsers,omitempty" name:"ManagerUsers"`
}

func (r *ModifyDepartmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepartmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchChangePwdTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 搜索区间开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchChangePwdTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchChangePwdTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 策略名称，1~60字符，仅支持中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，RuleName不可重复

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 策略描述，0~100字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 客户端IP，最多10组

	ClientIPList []*string `json:"ClientIPList,omitempty" name:"ClientIPList"`
	// 数据库名称，最多10组

	DBName []*string `json:"DBName,omitempty" name:"DBName"`
	// 表名称，最多10组

	TableName []*string `json:"TableName,omitempty" name:"TableName"`
	// SQL命令，最多10组

	SQLCommand []*string `json:"SQLCommand,omitempty" name:"SQLCommand"`
	// 限制SELECT结果集行数大于等于此数。默认为0，不限制。

	SQLRowCount *uint64 `json:"SQLRowCount,omitempty" name:"SQLRowCount"`
	// 字段名，最多50组

	FieldName []*string `json:"FieldName,omitempty" name:"FieldName"`
	// 0：&nbsp;不限&nbsp;1：每天&nbsp;2：指定时间段

	ExecutePeriod *uint64 `json:"ExecutePeriod,omitempty" name:"ExecutePeriod"`
	// 开始日期，格式:：YYYY-MM-DD

	ExecuteBeginDate *string `json:"ExecuteBeginDate,omitempty" name:"ExecuteBeginDate"`
	// 开始时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteBeginTime *string `json:"ExecuteBeginTime,omitempty" name:"ExecuteBeginTime"`
	// 结束日期，格式:：YYYY-MM-DD

	ExecuteEndDate *string `json:"ExecuteEndDate,omitempty" name:"ExecuteEndDate"`
	// 结束时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteEndTime *string `json:"ExecuteEndTime,omitempty" name:"ExecuteEndTime"`
	// 0：阻断&nbsp;1：放行

	ExecuteMode *uint64 `json:"ExecuteMode,omitempty" name:"ExecuteMode"`
	// 字段名和字段敏感数据分类的逻辑条件。0:&nbsp;or;&nbsp;&nbsp;1:and

	FieldLogicGate *uint64 `json:"FieldLogicGate,omitempty" name:"FieldLogicGate"`
	// 策略ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *ModifyAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTrialGuideStepRequest struct {
	*tchttp.BaseRequest

	// 试用引导步骤，取值0-27

	GuideStep *uint64 `json:"GuideStep,omitempty" name:"GuideStep"`
}

func (r *UpdateTrialGuideStepRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrialGuideStepRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportUserTaskRequest struct {
	*tchttp.BaseRequest

	// 文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 导出任务的部门&nbsp;id

	TaskDepartmentId *string `json:"TaskDepartmentId,omitempty" name:"TaskDepartmentId"`
	// 如果IdSet不为空，则忽略其他参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 精确查询，IdSet为空时才生效

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 邮箱，精确查询

	Email *string `json:"Email,omitempty" name:"Email"`
	// 查询具有指定资产ID访问权限的用户

	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`
	// 认证方式，0&nbsp;-&nbsp;本地,&nbsp;1&nbsp;-&nbsp;LDAP,&nbsp;2&nbsp;-&nbsp;OAuth,&nbsp;不传为全部

	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`
	// 部门ID，用于过滤属于某个部门的用户

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 参数过滤数组

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 精确查询，IdSet、UserName为空时才生效。如:&nbsp;"+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 每页条目数量，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *CreateExportUserTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportUserTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportAuditLogTaskRequest struct {
	*tchttp.BaseRequest

	// 导出日志类型&nbsp;1-字符会话&nbsp;2-图形会话&nbsp;3-文件传输会话&nbsp;4-数据库会话&nbsp;5-审计日志&nbsp;6-命令执行&nbsp;7-文件传输&nbsp;8-登录日志&nbsp;9-操作日志

	LogKind *uint64 `json:"LogKind,omitempty" name:"LogKind"`
	// 每页条目数，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeExportAuditLogTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportAuditLogTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInformation struct {

	// 文件名

	File *string `json:"File,omitempty" name:"File"`
	// 对文件的动作，1为上传，2为下载，3为删除，4为重命名

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 操作文件的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 文件的大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 操作文件时所用的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DeleteAccessControlRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChangePwdTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 过滤数组，支持：InstanceId&nbsp;实例ID，DeviceName&nbsp;实例名称，Ip&nbsp;内外IP，Account&nbsp;资产账号，LastChangeStatus&nbsp;上次改密状态。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移位置，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目。默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeChangePwdTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChangePwdTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportUserTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id列表

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteExportUserTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportUserTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 高危事件

		EventSet []*EventResult `json:"EventSet,omitempty" name:"EventSet"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDeliveryDetail struct {

	// 日志类型。目前取值：
	// CommandLog&nbsp;命令日志
	// FileLog&nbsp;文件日志

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// Ckafka&nbsp;name。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 是否启用

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 投递状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type TrialGuide struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 客户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 引导步骤

	GuideStep *uint64 `json:"GuideStep,omitempty" name:"GuideStep"`
	// 引导状态，0-未开始引导，1-引导中，2-引导结束

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type Command struct {

	// 命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 命令输入的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms

	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
	// 命令执行情况，1--允许，2--拒绝，3--确认

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 会话id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 设备account

	Account *string `json:"Account,omitempty" name:"Account"`
	// 设备ip

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// source&nbsp;ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 该命令所属会话的会话开始时间

	SessionTime *string `json:"SessionTime,omitempty" name:"SessionTime"`
	// 该命令所属会话的会话开始时间

	SessTime *string `json:"SessTime,omitempty" name:"SessTime"`
	// 复核时间

	ConfirmTime *string `json:"ConfirmTime,omitempty" name:"ConfirmTime"`
	// 用户部门id

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 用户部门name

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
	// 设备部门id

	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitempty" name:"DeviceDepartmentId"`
	// 设备部门name

	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitempty" name:"DeviceDepartmentName"`
	// 会话大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 签名值

	SignValue *string `json:"SignValue,omitempty" name:"SignValue"`
	// 资产类型

	DeviceKind *string `json:"DeviceKind,omitempty" name:"DeviceKind"`
}

type RunChangePwdTaskDetail struct {

	// 资产id

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 资产账号

	Account *string `json:"Account,omitempty" name:"Account"`
}

type DescribeAccessControlRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略详情

		Rule *AccessControlRuleDetail `json:"Rule,omitempty" name:"Rule"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListStatusRequest struct {
	*tchttp.BaseRequest

	// true：放开全部来源IP；false：不放开全部来源IP

	AllowAny *bool `json:"AllowAny,omitempty" name:"AllowAny"`
}

func (r *ModifyAccessWhiteListStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceRequest struct {
	*tchttp.BaseRequest

	// 需要开通服务的资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 已废弃

	Status *string `json:"Status,omitempty" name:"Status"`
	// 已废弃

	ModuleSet []*string `json:"ModuleSet,omitempty" name:"ModuleSet"`
	// 实例版本

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源节点数

	ResourceNode *int64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 带宽扩展包个数(4M)

	PackageBandwidth *int64 `json:"PackageBandwidth,omitempty" name:"PackageBandwidth"`
	// 授权点数扩展包个数(50点)

	PackageNode *int64 `json:"PackageNode,omitempty" name:"PackageNode"`
	// 日志投递

	LogDelivery *int64 `json:"LogDelivery,omitempty" name:"LogDelivery"`
}

func (r *ModifyResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemTaskStatistics struct {

	// 账号推送任务统计信息

	PushAccountTask *SystemTaskStatistic `json:"PushAccountTask,omitempty" name:"PushAccountTask"`
}

type DescribeLicenseUsageRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户列表

		UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDeviceRequest struct {
	*tchttp.BaseRequest

	// 资产记录ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产名，不能为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 资产IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 管理端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 资产组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 资产所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 资产多节点时填写

	IpPortSet []*string `json:"IpPortSet,omitempty" name:"IpPortSet"`
}

func (r *ModifyExternalDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest

	// 待删除的用户组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteUserGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepartmentsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDepartmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepartmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络域总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络域列表

		DomainSet []*Domain `json:"DomainSet,omitempty" name:"DomainSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableIntranetAccessRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 开通内网访问的vpc&nbsp;id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 开通内网访问的subnet&nbsp;id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *EnableIntranetAccessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableIntranetAccessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GraphResult struct {

	// 日期和当日新开会话数量

	Sessions []*KeyCount `json:"Sessions,omitempty" name:"Sessions"`
	// 日期和当日命令总数

	Commands []*KeyCount `json:"Commands,omitempty" name:"Commands"`
	// 日期和高危命令总数

	Dangers []*KeyCount `json:"Dangers,omitempty" name:"Dangers"`
	// 日期和sql总数

	Statements []*KeyCount `json:"Statements,omitempty" name:"Statements"`
	// 日期和高危sql总数

	DangerStatements []*KeyCount `json:"DangerStatements,omitempty" name:"DangerStatements"`
}

type TaskResult struct {

	// 运维任务结果日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 执行任务来源IP

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 运维任务所属用户

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 运维任务所属用户的姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 运维任务执行状态&nbsp;1&nbsp;-&nbsp;执行中，2&nbsp;-&nbsp;成功，3&nbsp;-&nbsp;失败，4&nbsp;-&nbsp;部分失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateAccessControlTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模版id

		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessControlTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessWhiteListRulesRequest struct {
	*tchttp.BaseRequest

	// 待删除的ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteAccessWhiteListRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessWhiteListRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrialGuideResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 试用引导

		TrialGuide *TrialGuide `json:"TrialGuide,omitempty" name:"TrialGuide"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrialGuideResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrialGuideResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventAlarmSetting struct {

	// 聚合开关。0关闭，1开启

	AggregateSwitch *uint64 `json:"AggregateSwitch,omitempty" name:"AggregateSwitch"`
	// 实时告警次数上限

	RealTimeCountLimit *uint64 `json:"RealTimeCountLimit,omitempty" name:"RealTimeCountLimit"`
	// 事件告警规则

	EventAlarmRule []*EventAlarmRule `json:"EventAlarmRule,omitempty" name:"EventAlarmRule"`
}

type DescribeUserCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisconnectDomainRequest struct {
	*tchttp.BaseRequest

	// 网络域id

	DomainIdSet []*string `json:"DomainIdSet,omitempty" name:"DomainIdSet"`
}

func (r *DisconnectDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisconnectDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产组名，最大长度32字符，不能为空

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资产组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 资产组所属部门ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *ModifyDeviceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchStatementRequest struct {
	*tchttp.BaseRequest

	// 会话Id

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 用户名，长度不超过20

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称，长度不超过64

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// Statement字段前端是否做base64加密&nbsp;0：否，1：是

	Encoding *uint64 `json:"Encoding,omitempty" name:"Encoding"`
	// SQL语句&nbsp;Base64编码

	Statement *string `json:"Statement,omitempty" name:"Statement"`
	// 真实姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 根据拦截状态进行过滤

	AuditAction []*uint64 `json:"AuditAction,omitempty" name:"AuditAction"`
	// 数据库类型&nbsp;1-MySQL&nbsp;2-PostgreSQL&nbsp;&nbsp;3-&nbsp;MongoDB

	DBTypes []*uint64 `json:"DBTypes,omitempty" name:"DBTypes"`
	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。
	// 2020-01-02T01:01:01Z

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 搜索区间的结束时间
	// 2020-01-02T01:01:01Z

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchStatementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchStatementRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LeaveTrackPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LeaveTrackPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LeaveTrackPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCost struct {

	// 总费用

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 优惠后总价

	RealCost *float64 `json:"RealCost,omitempty" name:"RealCost"`
	// 货币

	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

type DescribeShareClbIpsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeShareClbIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShareClbIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchKeyboardLoggerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 键盘记录事件数组

		SearchKeyboardLoggerResult []*KeyboardLogger `json:"SearchKeyboardLoggerResult,omitempty" name:"SearchKeyboardLoggerResult"`
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchKeyboardLoggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchKeyboardLoggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 主机账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *BindDeviceAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDeviceAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChangePwdTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 资产id数组

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 账号数组

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 1-正常修改&nbsp;&nbsp;2-关联资产账号

	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`
	// 所属部门ID，"1,2,3"

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 改密方式。1：使用执行账号修改密码；2：修改自身密码

	ChangeMethod *int64 `json:"ChangeMethod,omitempty" name:"ChangeMethod"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 认证生成方式。&nbsp;1:自动生成相同密码&nbsp;2:自动生成不同密码&nbsp;3:手动指定相同密码

	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// 手动指定密码需要传该字段

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密码限制长度

	PasswordLength *int64 `json:"PasswordLength,omitempty" name:"PasswordLength"`
	// 密码包含小写字母。1：包含

	SmallLetter *int64 `json:"SmallLetter,omitempty" name:"SmallLetter"`
	// 密码包含大写字母。1：包含

	BigLetter *int64 `json:"BigLetter,omitempty" name:"BigLetter"`
	// 密码包含数字。1：包含

	Digit *int64 `json:"Digit,omitempty" name:"Digit"`
	// 密码包含的特殊字符：base64编码

	Symbol *string `json:"Symbol,omitempty" name:"Symbol"`
	// 改密完成通知。0：不通知&nbsp;1：通知

	CompleteNotify *int64 `json:"CompleteNotify,omitempty" name:"CompleteNotify"`
	// 通知邮箱

	NotifyEmails []*string `json:"NotifyEmails,omitempty" name:"NotifyEmails"`
	// 加密压缩文件密码

	FilePassword *string `json:"FilePassword,omitempty" name:"FilePassword"`
	// 任务类型&nbsp;4-手工执行&nbsp;5-周期自动执行

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 周期>=7

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 首次执行时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
}

func (r *ModifyChangePwdTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChangePwdTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShowTopRequest struct {
	*tchttp.BaseRequest
}

func (r *ShowTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShowTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevicesDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevicesDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevicesDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 模版描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccessControlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessWhiteListRulesRequest struct {
	*tchttp.BaseRequest

	// 用户ID集合，非必需，如果使用IdSet参数则忽略Name参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 来源IP或网段，模糊查询，最大长度64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessWhiteListRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessWhiteListRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 用户名,&nbsp;3-20个字符,&nbsp;必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户姓名，最大长度20个字符，不能包含空白字符

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 大陆手机号直接填写，如果是其他国家、地区号码，&nbsp;按照"国家地区代码|手机号"的格式输入。如:&nbsp;"+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 电子邮件

	Email *string `json:"Email,omitempty" name:"Email"`
	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 所属用户组ID集合

	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`
	// 认证方式，0&nbsp;-&nbsp;本地，&nbsp;1&nbsp;-&nbsp;LDAP，&nbsp;2&nbsp;-&nbsp;OAuth&nbsp;不传则默认为0

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 访问时间段限制，&nbsp;由0、1组成的字符串，长度168(7&nbsp;×&nbsp;24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时，&nbsp;0&nbsp;-&nbsp;代表不允许访问，1&nbsp;-&nbsp;代表允许访问

	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
	// 所属部门ID，如：“1.2.3”

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *CreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAppAssetGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 资产组id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 应用资产id集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *AddAppAssetGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAppAssetGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserGroupMembersRequest struct {
	*tchttp.BaseRequest

	// 用户组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 成员用户ID集合

	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

func (r *AddUserGroupMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserGroupMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产同步标志

		AssetSyncFlags *AssetSyncFlags `json:"AssetSyncFlags,omitempty" name:"AssetSyncFlags"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupsRequest struct {
	*tchttp.BaseRequest

	// 资产组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 资产组名，最长64个字符，模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，缺省20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 部门ID，用于过滤属于某个部门的资产组

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeDeviceGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest

	// 如果IdSet不为空，则忽略其他参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20,&nbsp;最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 精确查询，IdSet为空时才生效

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如:&nbsp;"+852|xxxxxxxx"

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱，精确查询

	Email *string `json:"Email,omitempty" name:"Email"`
	// 查询具有指定资产ID访问权限的用户

	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`
	// 查询具有指定应用资产ID访问权限的用户

	AuthorizedAppAssetIdSet []*uint64 `json:"AuthorizedAppAssetIdSet,omitempty" name:"AuthorizedAppAssetIdSet"`
	// 认证方式，0&nbsp;-&nbsp;本地,&nbsp;1&nbsp;-&nbsp;LDAP,&nbsp;2&nbsp;-&nbsp;OAuth,&nbsp;不传为全部

	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`
	// 部门ID，用于过滤属于某个部门的用户

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 参数过滤数组
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMFAPreCheckRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMFAPreCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMFAPreCheckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSyncFlags struct {

	// 是否已完成角色授权

	RoleGranted *bool `json:"RoleGranted,omitempty" name:"RoleGranted"`
	// 是否已开启自动资产同步

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
}

type DescribeDepartmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 部门列表

		Departments *Departments `json:"Departments,omitempty" name:"Departments"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepartmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepartmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAppAssetRequest struct {
	*tchttp.BaseRequest

	// 应用资产id集合

	AppAssetIdSet []*uint64 `json:"AppAssetIdSet,omitempty" name:"AppAssetIdSet"`
	// 托管状态。&nbsp;0-不托管,&nbsp;1-托管

	BindStatus *uint64 `json:"BindStatus,omitempty" name:"BindStatus"`
}

func (r *BindAppAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAppAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest

	// 部署region

	DeployRegion *string `json:"DeployRegion,omitempty" name:"DeployRegion"`
	// 部署zone

	DeployZone *string `json:"DeployZone,omitempty" name:"DeployZone"`
	// 部署堡垒机的VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 部署堡垒机的SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 0非试用版，1试用版

	Trial *uint64 `json:"Trial,omitempty" name:"Trial"`
	// 是否共享clb，0：不共享，1：共享

	ShareClb *uint64 `json:"ShareClb,omitempty" name:"ShareClb"`
	// 资源类型。取值:standard/pro

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源节点数

	ResourceNode *int64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费时长

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 计费模式&nbsp;1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 自动续费

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 带宽扩展包

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
	// lb&nbsp;vip运营商。默认为空，内网访问。&nbsp;CMCC&nbsp;移动&nbsp;CTCC&nbsp;电信&nbsp;CUCC&nbsp;联通&nbsp;BGP&nbsp;外网CAP

	LbVipIsp *string `json:"LbVipIsp,omitempty" name:"LbVipIsp"`
}

func (r *CreateResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEnvSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDeviceAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDeviceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeviceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTaskStatistics struct {

	// 当前允许配置的运维任务最大数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 当前已配置的运维任务数量

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 当前已配置的周期性触发运维任务数量

	Periodic *int64 `json:"Periodic,omitempty" name:"Periodic"`
	// 当前已配置的手工触发运维任务数量

	Manual *int64 `json:"Manual,omitempty" name:"Manual"`
	// 当前正在执行中的运维任务数量

	Running *int64 `json:"Running,omitempty" name:"Running"`
}

type TopResult struct {

	// 在线用户数

	OnlineUser *int64 `json:"OnlineUser,omitempty" name:"OnlineUser"`
	// 高危用户和高危操作次数

	DangerUsers []*KeyCount `json:"DangerUsers,omitempty" name:"DangerUsers"`
	// 活跃用户和活跃次数

	ActiveUsers []*KeyCount `json:"ActiveUsers,omitempty" name:"ActiveUsers"`
	// 活跃主机和活跃次数及主机信息

	ActiveDevices []*ActiveDevice `json:"ActiveDevices,omitempty" name:"ActiveDevices"`
}

type CreateAccessControlTemplateRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessControlTemplateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlTemplateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmSettingRequest struct {
	*tchttp.BaseRequest

	// 聚合开关。0关闭，1开启

	AggregateSwitch *uint64 `json:"AggregateSwitch,omitempty" name:"AggregateSwitch"`
	// 实时告警次数限制

	RealTimeCountLimit *uint64 `json:"RealTimeCountLimit,omitempty" name:"RealTimeCountLimit"`
	// 事件规则信息

	EventAlarmRule *EventAlarmRule `json:"EventAlarmRule,omitempty" name:"EventAlarmRule"`
}

func (r *ModifyAlarmSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceAllRequest struct {
	*tchttp.BaseRequest

	// 每页条目数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeResourceAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalAppAsset struct {

	// 应用资产类型。1-web应用

	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`
	// 客户端工具类型

	ClientAppKind *string `json:"ClientAppKind,omitempty" name:"ClientAppKind"`
	// 应用服务器ID，添加单台使用

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 应用服务器实例ID，批量导入使用

	DeviceInstanceId *string `json:"DeviceInstanceId,omitempty" name:"DeviceInstanceId"`
	// 应用资产名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 客户端工具路径

	ClientAppPath *string `json:"ClientAppPath,omitempty" name:"ClientAppPath"`
	// 应用服务器账号ID，添加单台使用

	DeviceAccountId *uint64 `json:"DeviceAccountId,omitempty" name:"DeviceAccountId"`
	// 应用服务器账号名，批量导入使用

	DeviceAccount *string `json:"DeviceAccount,omitempty" name:"DeviceAccount"`
	// web应用资产url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 部门id

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建用户的ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportDeviceTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		Tasks []*ExportDeviceTask `json:"Tasks,omitempty" name:"Tasks"`
		// 任务数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportDeviceTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportDeviceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDeviceAccountPrivateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境

		Environment *string `json:"Environment,omitempty" name:"Environment"`
		// 站点

		Site *string `json:"Site,omitempty" name:"Site"`
		// 环境

		RunTimeEnv *string `json:"RunTimeEnv,omitempty" name:"RunTimeEnv"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceAccountRequest struct {
	*tchttp.BaseRequest

	// 主机记录ID

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *CreateDeviceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlRuleDetail struct {

	// RuleID&nbsp;&nbsp;策略ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略名称，1~60字符，仅支持中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，RuleName不可重复

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 0:&nbsp;内置&nbsp;1:&nbsp;自定义

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 策略描述，0~100字符

	Description *string `json:"Description,omitempty" name:"Description"`
	// 客户端IP，最多10组

	ClientIPList []*string `json:"ClientIPList,omitempty" name:"ClientIPList"`
	// 数据库名称，最多10组

	DBName []*string `json:"DBName,omitempty" name:"DBName"`
	// 表名称，最多10组

	TableName []*string `json:"TableName,omitempty" name:"TableName"`
	// SQL命令，最多10组

	SQLCommand []*string `json:"SQLCommand,omitempty" name:"SQLCommand"`
	// 限制SELECT结果集行数大于等于此数。默认为0，不限制。

	SelectRowsCount *uint64 `json:"SelectRowsCount,omitempty" name:"SelectRowsCount"`
	// 字段名，最多50组

	FieldName []*string `json:"FieldName,omitempty" name:"FieldName"`
	// 0：&nbsp;不限&nbsp;1：每天&nbsp;2：指定时间段

	ExecutePeriod *uint64 `json:"ExecutePeriod,omitempty" name:"ExecutePeriod"`
	// 开始日期，格式:：YYYY-MM-DD

	ExecuteBeginDate *string `json:"ExecuteBeginDate,omitempty" name:"ExecuteBeginDate"`
	// 开始时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteBeginTime *string `json:"ExecuteBeginTime,omitempty" name:"ExecuteBeginTime"`
	// 结束日期，格式:：YYYY-MM-DD

	ExecuteEndDate *string `json:"ExecuteEndDate,omitempty" name:"ExecuteEndDate"`
	// 结束时间，24小时制，格式:&nbsp;HH:mm:ss

	ExecuteEndTime *string `json:"ExecuteEndTime,omitempty" name:"ExecuteEndTime"`
	// 0：阻断&nbsp;1：放行

	ExecuteMode *uint64 `json:"ExecuteMode,omitempty" name:"ExecuteMode"`
	// 字段名和字段敏感数据分类的逻辑条件。0:&nbsp;or;&nbsp;&nbsp;1:and

	FieldLogicGate *uint64 `json:"FieldLogicGate,omitempty" name:"FieldLogicGate"`
	// 关联模板数量

	TemplateCount *int64 `json:"TemplateCount,omitempty" name:"TemplateCount"`
}

type DescribePushAccountTaskRequest struct {
	*tchttp.BaseRequest

	// 过滤数组。Name支持以下值:
	// OperationId&nbsp;任务ID
	// TaskName&nbsp;任务名
	// RunAccount&nbsp;执行账号

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 所属部门ID，如：“1.2.3”。

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移位置，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20,&nbsp;最大500。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePushAccountTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushAccountTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireCreateDasbResourceRequest struct {
	*tchttp.BaseRequest

	// 私有网络vpcId&nbsp;&nbsp;vpc-xxxx

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网subnetId&nbsp;subnet-xxxx

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// lb&nbsp;vip运营商。默认为空，内网访问。&nbsp;CMCC&nbsp;移动&nbsp;CTCC&nbsp;电信&nbsp;CUCC&nbsp;联通&nbsp;BGP&nbsp;外网CAP

	LbVipIsp *string `json:"LbVipIsp,omitempty" name:"LbVipIsp"`
	// 资源类型。标准版(standard)或者专业版(pro)

	ResourceEdition *string `json:"ResourceEdition,omitempty" name:"ResourceEdition"`
	// 资源默认节点数。取值：10/20/50/100/200/500/1000

	ResourceNode *uint64 `json:"ResourceNode,omitempty" name:"ResourceNode"`
	// 带宽扩展包&nbsp;

	MbpExp *uint64 `json:"MbpExp,omitempty" name:"MbpExp"`
	// 节点扩展包

	NodeExp *uint64 `json:"NodeExp,omitempty" name:"NodeExp"`
	// 计费周期&nbsp;只支持&nbsp;m

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费时长

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 付费方式。0表示按需计费/后付费，1表示预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 部署region

	DeployRegion *string `json:"DeployRegion,omitempty" name:"DeployRegion"`
	// 部署的zone

	DeployZone *string `json:"DeployZone,omitempty" name:"DeployZone"`
	// 日志投递

	LogDelivery *uint64 `json:"LogDelivery,omitempty" name:"LogDelivery"`
}

func (r *InquireCreateDasbResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreateDasbResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewReportRequest struct {
	*tchttp.BaseRequest

	// 报表记录&nbsp;Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ViewReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ViewReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCount struct {

	// 资产数目

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 资产类型

	Kind *int64 `json:"Kind,omitempty" name:"Kind"`
}

type DeleteExportAuditLogTaskRequest struct {
	*tchttp.BaseRequest

	// 审计日志导出任务Id

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteExportAuditLogTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportAuditLogTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志推送配置。

		DeliveryConfig *LogDelivery `json:"DeliveryConfig,omitempty" name:"DeliveryConfig"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDeliveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessControlRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceAccountBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceAccountBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPushAccountTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPushAccountTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPushAccountTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileTypeFilter struct {

	// 需要查询的文件传输类型，如SFTP/CLIP/RDP/RZSZ

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 在当前指定的protocol下进一步过滤具体操作类型,如剪贴板文件上传，剪贴板文件下载等

	Method []*int64 `json:"Method,omitempty" name:"Method"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 地域码,&nbsp;如:&nbsp;ap-city

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 按照堡垒机开通的&nbsp;VPC&nbsp;实例ID查询

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 资源ID集合，当传入ID集合时忽略&nbsp;ApCode&nbsp;和&nbsp;VpcId

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 是否使用公参&nbsp;Region，1&nbsp;是，0&nbsp;否

	UseRegion *string `json:"UseRegion,omitempty" name:"UseRegion"`
	// 每页条目数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorSessionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 会话监控所需信息

		ReplayInfo *ReplayInformation `json:"ReplayInfo,omitempty" name:"ReplayInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MonitorSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MonitorSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmdTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功后返回的记录ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmdTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户组成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户组成员列表

		UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest

	// 访问权限名称，最大32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启文件传输上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，目前暂未使用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启文件传输下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，目前暂未使用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户ID

	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`
	// 关联的用户组ID

	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`
	// 关联的资产ID

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 关联的资产组ID

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
	// 关联的账号

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 是否允许键盘记录

	AllowKeyboardLogger *bool `json:"AllowKeyboardLogger,omitempty" name:"AllowKeyboardLogger"`
	// 关联的高危命令模板ID

	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`
	// 关联高危DB模板ID

	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitempty" name:"ACTemplateIdSet"`
	// 访问权限ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否开启&nbsp;RDP&nbsp;磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启&nbsp;RDP&nbsp;磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启rz&nbsp;sz文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启rz&nbsp;sz文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启&nbsp;SFTP&nbsp;文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 关联的应用资产ID集合

	AppAssetIdSet []*uint64 `json:"AppAssetIdSet,omitempty" name:"AppAssetIdSet"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 权限所属部门的ID，如：1.2.3

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 是否允许使用访问串

	AllowAccessCredential *bool `json:"AllowAccessCredential,omitempty" name:"AllowAccessCredential"`
}

func (r *ModifyAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLDAPSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLDAPSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLDAPSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest

	// ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *ResetDeviceAccountPrivateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunPushAccountTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunPushAccountTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunPushAccountTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppAssetRequest struct {
	*tchttp.BaseRequest

	// 应用资产参数列表

	AppAssetSet []*ExternalAppAsset `json:"AppAssetSet,omitempty" name:"AppAssetSet"`
}

func (r *CreateAppAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOperationTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogOutputSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志外发配置

		LogOutputSettings *LogOutputSettings `json:"LogOutputSettings,omitempty" name:"LogOutputSettings"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogOutputSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogOutputSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessWhiteListRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessWhiteListRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessWhiteListRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命令列表

		CommandSet []*Command `json:"CommandSet,omitempty" name:"CommandSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCommandBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCommandBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecuritySetting struct {

	// 认证方式设置

	AuthMode *AuthModeSetting `json:"AuthMode,omitempty" name:"AuthMode"`
	// 密码安全设置

	Password *PasswordSetting `json:"Password,omitempty" name:"Password"`
	// 登录安全设置

	Login *LoginSetting `json:"Login,omitempty" name:"Login"`
	// LDAP配置信息

	LDAP *LDAPSetting `json:"LDAP,omitempty" name:"LDAP"`
	// OAuth配置信息

	OAuth *OAuthSetting `json:"OAuth,omitempty" name:"OAuth"`
	// 国密认证方式设置

	AuthModeGM *AuthModeSetting `json:"AuthModeGM,omitempty" name:"AuthModeGM"`
	// 资产重连次数

	Reconnection *ReconnectionSetting `json:"Reconnection,omitempty" name:"Reconnection"`
}

type CreateAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建成功的访问权限ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncUserFromCamResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncUserFromCamResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncUserFromCamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdcSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCdcSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogDeliveryRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 接入方式。
	// Domain&nbsp;域名接入
	// Internal&nbsp;支撑环境接入

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// ckafka实例Id。

	CkafkaInstanceId *string `json:"CkafkaInstanceId,omitempty" name:"CkafkaInstanceId"`
	// 接入地址。
	// 域名接入的情况下是域名和端口；支撑环境接入的情况下是虚拟地址和端口。

	Address *string `json:"Address,omitempty" name:"Address"`
	// 账号。域名接入时支持，支撑环境接入时不支持

	User *string `json:"User,omitempty" name:"User"`
	// 密码。域名接入时支持，支撑环境接入时不支持

	Password *string `json:"Password,omitempty" name:"Password"`
	// Ckafka实例所在地域。

	CkafkaRegion *string `json:"CkafkaRegion,omitempty" name:"CkafkaRegion"`
	// 日志投递配置。

	Details []*CreateLogDeliveryDetail `json:"Details,omitempty" name:"Details"`
	// Ckafka实例所在vpc&nbsp;Id

	CkafkaUniqVpcId *string `json:"CkafkaUniqVpcId,omitempty" name:"CkafkaUniqVpcId"`
	// Ckafka实例所在子网Id

	CkafkaUniqSubnetId *string `json:"CkafkaUniqSubnetId,omitempty" name:"CkafkaUniqSubnetId"`
}

func (r *CreateLogDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketSubmitFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// false:未开启自动提交
		// true:自动提交

		TicketSubmitFlag *bool `json:"TicketSubmitFlag,omitempty" name:"TicketSubmitFlag"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketSubmitFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketSubmitFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeLogDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteChangePwdTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id数组

	IdSet []*int64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeleteChangePwdTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteChangePwdTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlTemplateRuleRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 规则id列表，id格式为：acrule-xxxxx

	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`
}

func (r *CreateAccessControlTemplateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlTemplateRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogOutputSettingsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogOutputSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogOutputSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppAssetsDepartmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppAssetsDepartmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppAssetsDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogDeliveryRequest struct {
	*tchttp.BaseRequest

	// 堡垒机实例Id。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 需要修改的配置。

	Details []*ModifyLogDeliveryDetail `json:"Details,omitempty" name:"Details"`
}

func (r *ModifyLogDeliveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogDeliveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogOutputSettingsRequest struct {
	*tchttp.BaseRequest

	// 是否开启日志外发，false-不开启，true-开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 是否开启产品登录日志外发，false-不开启，true-开启

	ProductLogin *bool `json:"ProductLogin,omitempty" name:"ProductLogin"`
	// 是否开启资产登录日志外发，false-不开启，true-开启

	DeviceLogin *bool `json:"DeviceLogin,omitempty" name:"DeviceLogin"`
}

func (r *ModifyLogOutputSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogOutputSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 某会话的文件操作列表

		SearchFileBySidResult []*SearchFileBySidResult `json:"SearchFileBySidResult,omitempty" name:"SearchFileBySidResult"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFileBySidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchFileBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupWithCount struct {

	// 组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 成员数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DescribeOperationTaskStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务统计信息

		OperationTaskStatistics *OperationTaskStatistics `json:"OperationTaskStatistics,omitempty" name:"OperationTaskStatistics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTaskStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePushAccountTasksRequest struct {
	*tchttp.BaseRequest

	// 任务Id数组。

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

func (r *DeletePushAccountTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePushAccountTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePushAccountTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情

		Details []*PushAccountTaskDetail `json:"Details,omitempty" name:"Details"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushAccountTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushAccountTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthModeSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuthModeSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuthModeSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyboardLogger struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户部门名称

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
	// 设备部门名称

	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitempty" name:"DeviceDepartmentName"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 客户端IP

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 写ES时间

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 资产ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 用户部门ID

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 会话ID

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 按键事件时间，按分钟级粒度

	KeyTime *string `json:"KeyTime,omitempty" name:"KeyTime"`
	// 按键信息

	Keys *string `json:"Keys,omitempty" name:"Keys"`
	// 会话时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 设备部门ID

	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitempty" name:"DeviceDepartmentId"`
}

type CreateDeviceAccountBatchRequest struct {
	*tchttp.BaseRequest

	// 主机ID

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 主机账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 账号密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 私钥口令

	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

func (r *CreateDeviceAccountBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceAccountBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLDAPSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启LDAP认证，false-不开启，true-开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 服务器地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用服务器地址

	IpBackup *string `json:"IpBackup,omitempty" name:"IpBackup"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否开启SSL，false-不开启，true-开启

	EnableSSL *bool `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// Base&nbsp;DN

	BaseDN *string `json:"BaseDN,omitempty" name:"BaseDN"`
	// 管理员账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理员密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 用户属性

	AttributeUser *string `json:"AttributeUser,omitempty" name:"AttributeUser"`
	// 用户名属性

	AttributeUserName *string `json:"AttributeUserName,omitempty" name:"AttributeUserName"`
	// 自动同步，false-不开启，true-开启

	AutoSync *bool `json:"AutoSync,omitempty" name:"AutoSync"`
	// 覆盖用户信息，false-不开启，true-开启

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
	// 网络域Id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 同步周期，30～60000之间的整数

	SyncPeriod *uint64 `json:"SyncPeriod,omitempty" name:"SyncPeriod"`
	// 是否同步全部，false-不开启，true-开启

	SyncAll *bool `json:"SyncAll,omitempty" name:"SyncAll"`
	// 同步OU列表，SyncAll为false时必传

	SyncUnitSet []*string `json:"SyncUnitSet,omitempty" name:"SyncUnitSet"`
	// 组织单元属性

	AttributeUnit *string `json:"AttributeUnit,omitempty" name:"AttributeUnit"`
	// 用户姓名属性

	AttributeRealName *string `json:"AttributeRealName,omitempty" name:"AttributeRealName"`
	// 手机号属性

	AttributePhone *string `json:"AttributePhone,omitempty" name:"AttributePhone"`
	// 邮箱属性

	AttributeEmail *string `json:"AttributeEmail,omitempty" name:"AttributeEmail"`
}

func (r *ModifyLDAPSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLDAPSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubtaskResult struct {

	// 执行日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 执行主机实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 执行主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 执行主机地域

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// 执行主机外网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 执行主机内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 运维任务状态&nbsp;1&nbsp;-&nbsp;执行中，2&nbsp;-&nbsp;成功，&nbsp;3&nbsp;-&nbsp;失败，4&nbsp;-&nbsp;超时

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 运维任务失败原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 运维任务命令退出码

	ExitCode *int64 `json:"ExitCode,omitempty" name:"ExitCode"`
	// 运维任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 运维任务执行结果输出

	StdOut *string `json:"StdOut,omitempty" name:"StdOut"`
	// 运维任务执行结果错误

	StdErr *string `json:"StdErr,omitempty" name:"StdErr"`
	// 资产名

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产账号

	Account *string `json:"Account,omitempty" name:"Account"`
}

type DescribeUKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// USB&nbsp;Key总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// USB&nbsp;Key详情列表

		UKeyUserSet []*UKeyUserDetail `json:"UKeyUserSet,omitempty" name:"UKeyUserSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUKeysRequest struct {
	*tchttp.BaseRequest

	// UKey名称，模糊查询，最长64字符

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 已绑定用户的用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 	部门ID，用于过滤属于某个部门的用户

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// UKey&nbsp;ID集合

	KeyIdSet []*string `json:"KeyIdSet,omitempty" name:"KeyIdSet"`
}

func (r *DescribeUKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络域id

		DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemTaskStatisticsRequest struct {
	*tchttp.BaseRequest

	// 0：默认账号推送；3：账号推送；4：手动改密任务；5：周期改密任务。

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSystemTaskStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemTaskStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireCreateDasbResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品价格

		ProductCost *ProductCost `json:"ProductCost,omitempty" name:"ProductCost"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireCreateDasbResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquireCreateDasbResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginOpserverRequest struct {
	*tchttp.BaseRequest
}

func (r *LoginOpserverRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginOpserverRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Clb struct {

	// 负载均衡IP

	ClbIp *string `json:"ClbIp,omitempty" name:"ClbIp"`
}

type CreateExportUserTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务&nbsp;ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportUserTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportUserTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisitTrackPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VisitTrackPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisitTrackPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ticket struct {

	// 运维开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运维结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 工单状态，0待提交(草稿)，1待审批，2批准，3已执行，4驳回,5删除，6过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批人uin

	ApproveUser *string `json:"ApproveUser,omitempty" name:"ApproveUser"`
	// acl权限Id

	AclId *uint64 `json:"AclId,omitempty" name:"AclId"`
	// 申请人用户名

	ApplyUserName *string `json:"ApplyUserName,omitempty" name:"ApplyUserName"`
	// 数据库主键

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 工单类型，0-运维访问工单，1-权限工单,&nbsp;2-数据库变更工单,3-自动执行命令/脚本工单

	TicketType *uint64 `json:"TicketType,omitempty" name:"TicketType"`
	// 申请人真实姓名

	ApplyRealName *string `json:"ApplyRealName,omitempty" name:"ApplyRealName"`
	// 审批时间

	ApproveTime *string `json:"ApproveTime,omitempty" name:"ApproveTime"`
	// 部门名称

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// 申请时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 待审批状态下，可以审批的人

	ApproveUserSet []*string `json:"ApproveUserSet,omitempty" name:"ApproveUserSet"`
	// 工单审批备注

	ApproveDesc *string `json:"ApproveDesc,omitempty" name:"ApproveDesc"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 部门ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 部门Sequence

	DepartmentSequence *uint64 `json:"DepartmentSequence,omitempty" name:"DepartmentSequence"`
	// 工单ID,id-s3ed4r5e

	TicketId *string `json:"TicketId,omitempty" name:"TicketId"`
	// 申请人用户Id

	ApplyUser *uint64 `json:"ApplyUser,omitempty" name:"ApplyUser"`
	// 申请权限

	Name *string `json:"Name,omitempty" name:"Name"`
	// 申请时间戳，微妙

	ApplyTimeVersion *int64 `json:"ApplyTimeVersion,omitempty" name:"ApplyTimeVersion"`
	// acl权限

	Acl *Acl `json:"Acl,omitempty" name:"Acl"`
	// 审批人信息

	ApproveUserInfoSet []*DepartmentManagerUser `json:"ApproveUserInfoSet,omitempty" name:"ApproveUserInfoSet"`
}

type EventAlarmRule struct {

	// 告警开关。0关闭，1开启

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 全天告警。0不允许，1允许

	AllowAllDay *uint64 `json:"AllowAllDay,omitempty" name:"AllowAllDay"`
	// 允许开始告警小时。取值0-23

	AllowStartHour *uint64 `json:"AllowStartHour,omitempty" name:"AllowStartHour"`
	// 允许开始告警分钟。取值0-59

	AllowStartMinute *uint64 `json:"AllowStartMinute,omitempty" name:"AllowStartMinute"`
	// 允许结束告警小时。取值0-23

	AllowEndHour *uint64 `json:"AllowEndHour,omitempty" name:"AllowEndHour"`
	// 允许结束告警分钟。取值0-59

	AllowEndMinute *uint64 `json:"AllowEndMinute,omitempty" name:"AllowEndMinute"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 告警事件类型。1-linux命令拦截，2-sql语句拦截，3-文件操作拦截

	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`
}

type DescribeDeviceAccountsRequest struct {
	*tchttp.BaseRequest

	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交

	Account *string `json:"Account,omitempty" name:"Account"`
	// 主机ID，未使用IdSet时必须携带

	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeviceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportAuditLogTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审计日志导出任务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 审计日志导出任务列表

		Tasks []*ExportAuditLogTask `json:"Tasks,omitempty" name:"Tasks"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportAuditLogTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportAuditLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSubtaskResultByIdRequest struct {
	*tchttp.BaseRequest

	// 运维任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页的页内记录数，默认为20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 运维父任务执行日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维父任务执行状态

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SearchSubtaskResultByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSubtaskResultByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogDeliveryDetail struct {

	// 日志类型。目前取值：
	// CommandLog&nbsp;命令日志
	// FileLog&nbsp;文件日志

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// Ckafka&nbsp;topic。需要为已存在的topic。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 是否启用。默认是。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ImportDevicesRequest struct {
	*tchttp.BaseRequest

	// 地域参数

	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`
	// cvm参数列表

	CvmSet []*Cvm `json:"CvmSet,omitempty" name:"CvmSet"`
	// 用户主机所在的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 堡垒机服务ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *ImportDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// ip&nbsp;，网段

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 网络域名称

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *CreateDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest

	// 用户组ID集合

	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
	// 用户组名，模糊查询,长度：0-64字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页偏移位置，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条目数量，缺省20，最大500

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 部门ID，用于过滤属于某个部门的用户组

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

func (r *DescribeUserGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSyncStatus struct {

	// 上一次同步完成的时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 上一次同步的结果。&nbsp;0&nbsp;-&nbsp;从未进行,&nbsp;1&nbsp;-&nbsp;成功，&nbsp;2&nbsp;-&nbsp;失败

	LastStatus *uint64 `json:"LastStatus,omitempty" name:"LastStatus"`
	// 同步任务是否正在进行中

	InProcess *bool `json:"InProcess,omitempty" name:"InProcess"`
	// 任务错误消息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type CreateReportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表任务&nbsp;Id

		ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessControlTemplateRequest struct {
	*tchttp.BaseRequest

	// 需要删除的模版id列表

	TemplateIds []*string `json:"TemplateIds,omitempty" name:"TemplateIds"`
}

func (r *DeleteAccessControlTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAccessControlTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdcSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true&nbsp;--&nbsp;用户可以在CDC环境下开通堡垒机资源；
		// false&nbsp;--&nbsp;用户不能在CDC环境下开通堡垒机资源；

		CdcEnabled *bool `json:"CdcEnabled,omitempty" name:"CdcEnabled"`
		// 当CdcEnabled为true时，这个参数返回用户可以在CDC环境下开通堡垒机实例的CDC集群ID列表；当CdcEnabled位false时，这个参数返回空列表；

		CdcEnabledClusterIds []*string `json:"CdcEnabledClusterIds,omitempty" name:"CdcEnabledClusterIds"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdcSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// 权限名称，最大32字符，不能包含空白字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启磁盘映射

	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`
	// 是否开启剪贴板文件上行

	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`
	// 是否开启剪贴板文件下行

	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`
	// 是否开启剪贴板文本（含图片）上行

	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`
	// 是否开启剪贴板文本（含图片）下行

	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`
	// 是否开启&nbsp;SFTP&nbsp;文件上传

	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`
	// 文件传输上传大小限制（预留参数，目前暂未使用）

	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`
	// 是否开启&nbsp;SFTP&nbsp;文件下载

	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`
	// 文件传输下载大小限制（预留参数，目前暂未使用）

	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`
	// 是否允许任意账号登录

	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`
	// 关联的用户ID集合

	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`
	// 关联的用户组ID

	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`
	// 关联的资产ID集合

	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`
	// 关联的资产组ID

	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`
	// 关联的账号

	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`
	// 关联的高危命令模板ID

	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`
	// 关联高危DB模板ID

	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitempty" name:"ACTemplateIdSet"`
	// 是否开启rdp磁盘映射文件上传

	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`
	// 是否开启rdp磁盘映射文件下载

	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`
	// 是否开启rz&nbsp;sz文件上传

	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`
	// 是否开启rz&nbsp;sz文件下载

	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`
	// 是否开启&nbsp;SFTP&nbsp;文件删除

	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`
	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`
	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效

	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
	// 访问权限所属部门的ID

	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 是否允许键盘记录

	AllowKeyboardLogger *bool `json:"AllowKeyboardLogger,omitempty" name:"AllowKeyboardLogger"`
	// 关联的应用资产ID集合

	AppAssetIdSet []*uint64 `json:"AppAssetIdSet,omitempty" name:"AppAssetIdSet"`
	// 是否允许使用访问串，默认允许

	AllowAccessCredential *bool `json:"AllowAccessCredential,omitempty" name:"AllowAccessCredential"`
}

func (r *CreateAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserBatchRequest struct {
	*tchttp.BaseRequest

	// 用户信息

	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`
}

func (r *CreateUserBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type ModifyAssetSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShareClbIp struct {

	// clb的实例id

	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`
	// clb的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// clb所属vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// clb的公网ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 监听的端口数

	ListenedPortCount *uint64 `json:"ListenedPortCount,omitempty" name:"ListenedPortCount"`
}

type DescribeOperationTaskStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOperationTaskStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTaskStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntranetAccessEntry struct {

	// 堡垒机实例id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 内网访问url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 内网访问的vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 内网访问的网段

	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`
	// 0&nbsp;内网未开通，1&nbsp;内网已开通，2&nbsp;内网访问开通中

	IntranetAccessStatus *uint64 `json:"IntranetAccessStatus,omitempty" name:"IntranetAccessStatus"`
}

type DeleteDomainsRequest struct {
	*tchttp.BaseRequest

	// 待删除的网络域ID集合

	DomainIdSet []*string `json:"DomainIdSet,omitempty" name:"DomainIdSet"`
}

func (r *DeleteDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDomainsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KillSessionRequest struct {
	*tchttp.BaseRequest

	// 会话Sid

	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

func (r *KillSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *KillSessionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资产信息列表

		DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维任务列表

		OperationTasks []*OperationTask `json:"OperationTasks,omitempty" name:"OperationTasks"`
		// 任务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuthModeSettingRequest struct {
	*tchttp.BaseRequest

	// 双因子认证，0-不开启，1-OTP，2-短信

	AuthMode *uint64 `json:"AuthMode,omitempty" name:"AuthMode"`
	// 资源类型，0：普通&nbsp;1：国密

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *ModifyAuthModeSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuthModeSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLDAPSyncFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLDAPSyncFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLDAPSyncFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncUserFromCamRequest struct {
	*tchttp.BaseRequest
}

func (r *SyncUserFromCamRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncUserFromCamRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMFAPreCheckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMFAPreCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMFAPreCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OAuthSetting struct {

	// 使用Okta认证时指定范围。

	Scopes []*string `json:"Scopes,omitempty" name:"Scopes"`
	// 是否开启OAuth认证

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// OAuth认证方式。

	AuthMethod *string `json:"AuthMethod,omitempty" name:"AuthMethod"`
	// OAuth认证客户端Id。

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 获取OAuth认证授权码URL。

	CodeUrl *string `json:"CodeUrl,omitempty" name:"CodeUrl"`
	// 获取OAuth令牌URL。

	TokenUrl *string `json:"TokenUrl,omitempty" name:"TokenUrl"`
	// 获取OAuth用户信息URL。

	UserInfoUrl *string `json:"UserInfoUrl,omitempty" name:"UserInfoUrl"`
}

type PushAccountTaskInfoResult struct {

	// 运维任务结果日志ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 运维任务ID

	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`
	// 执行账号

	RunAccount *string `json:"RunAccount,omitempty" name:"RunAccount"`
	// 待创建账号

	CreateAccount *string `json:"CreateAccount,omitempty" name:"CreateAccount"`
	// 密码认证

	PasswordAuth *int64 `json:"PasswordAuth,omitempty" name:"PasswordAuth"`
	// 私钥认证

	PrivateKeyAuth *int64 `json:"PrivateKeyAuth,omitempty" name:"PrivateKeyAuth"`
	// 认证生成策略

	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitempty" name:"AuthGenerationStrategy"`
	// home路径

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// shell路径

	ShellPath *string `json:"ShellPath,omitempty" name:"ShellPath"`
	// 用户组

	GroupSet *string `json:"GroupSet,omitempty" name:"GroupSet"`
	// 用户部门名称

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
	// 用户部门id

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
}

type SearchCommandResult struct {

	// 命令输入的时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 姓名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
	// 资产ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 资产公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 资产内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 命令

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 命令执行情况，1--允许，2--拒绝

	Action *uint64 `json:"Action,omitempty" name:"Action"`
	// 命令所属的会话ID

	Sid *string `json:"Sid,omitempty" name:"Sid"`
	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms

	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// source&nbsp;ip

	FromIp *string `json:"FromIp,omitempty" name:"FromIp"`
	// 该命令所属会话的会话开始时间

	SessionTime *string `json:"SessionTime,omitempty" name:"SessionTime"`
	// 该命令所属会话的会话开始时间（废弃，使用SessionTime）

	SessTime *string `json:"SessTime,omitempty" name:"SessTime"`
	// 复核时间

	ConfirmTime *string `json:"ConfirmTime,omitempty" name:"ConfirmTime"`
	// 部门id

	UserDepartmentId *string `json:"UserDepartmentId,omitempty" name:"UserDepartmentId"`
	// 用户部门名称

	UserDepartmentName *string `json:"UserDepartmentName,omitempty" name:"UserDepartmentName"`
	// 设备部门id

	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitempty" name:"DeviceDepartmentId"`
	// 设备部门名称

	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitempty" name:"DeviceDepartmentName"`
	// 会话大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 资产类型

	DeviceKind *string `json:"DeviceKind,omitempty" name:"DeviceKind"`
	// 签名值

	SignValue *string `json:"SignValue,omitempty" name:"SignValue"`
}

type CreateAccessWhiteListRuleRequest struct {
	*tchttp.BaseRequest

	// ip&nbsp;10.10.10.1或者网段10.10.10.0/24，最小长度4字节，最大长度40字节。

	Source *string `json:"Source,omitempty" name:"Source"`
	// 备注信息，最小长度0字符，最大长度40字符。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAccessWhiteListRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessWhiteListRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportAuditLogTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportAuditLogTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportAuditLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

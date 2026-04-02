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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type BindVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// VPC信息

	VpcInfos []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

func (r *BindVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID，以逗号分隔

	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *DeleteVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordDetail struct {

	// 记录ID

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 记录对应的域名ID

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 主机记录

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// TTL

	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`
	// MX优先级

	Mx *int64 `json:"Mx,omitempty" name:"Mx"`
	// 记录是否可用，0不可用，1可用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 记录状态，启用，暂停等

	Status *string `json:"Status,omitempty" name:"Status"`
	// 其他信息

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 记录创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 记录最后修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 记录权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type DescribeVpcDnsDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 域名数量信息

		Info *DomainCountInfo `json:"Info,omitempty" name:"Info"`
		// 域名信息

		Domains []*DomainDetail `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// MX优先级

	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
}

func (r *CreateVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordId struct {

	// 记录id

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

type VpcInfos struct {

	// VpcId

	VpcId *int64 `json:"vpc_id,omitempty" name:"VpcId"`
	// RegionId

	RegionId *int64 `json:"region_id,omitempty" name:"RegionId"`
	// UnVpcId

	UnVpcId *string `json:"un_vpc_id,omitempty" name:"UnVpcId"`
}

type Record struct {

	// 记录id

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

type CreateVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 标签数组

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启子域名递归，只接受"ENABLED”和“DISABLED”

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *CreateVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 记录数量描述

		Info *RecordListCountInfo `json:"Info,omitempty" name:"Info"`
		// 解析记录信息

		Records []*RecordDetail `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 记录ID

	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *string `json:"Value,omitempty" name:"Value"`
	// MX优先级

	Mx *uint64 `json:"Mx,omitempty" name:"Mx"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainDetail struct {

	// 域名id

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 域名所有者uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 最后修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 域名记录数量

	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
	// 域名备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否开启子域名递归（“ENABLED”和“DISABLED”）

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
	// VPC信息

	VpcInfos []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

type RecordListFilters struct {

	// 过滤类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateVpcDnsDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordListCountInfo struct {

	// 该域名的记录总数量

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 本次返回的记录数量

	RecordTotal *int64 `json:"RecordTotal,omitempty" name:"RecordTotal"`
}

type ModifyVpcDnsDomainRequest struct {
	*tchttp.BaseRequest

	// 域名ID，以逗号分割

	DomainIds *string `json:"DomainIds,omitempty" name:"DomainIds"`
	// 是否开启子域名递归，只接受"ENABLED"和"DISABLED"两种

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *ModifyVpcDnsDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsRecordListRequest struct {
	*tchttp.BaseRequest

	// 长度

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*RecordListFilters `json:"Filters,omitempty" name:"Filters"`
	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeVpcDnsRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRemarkRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateVpcDnsDomainRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 记录Id

		Data *RecordId `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsDomainListRequest struct {
	*tchttp.BaseRequest

	// 长度

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*DomainListFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcDnsDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsDomainRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求响应时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsDomainRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsDomainRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsRecordRequest struct {
	*tchttp.BaseRequest

	// 域名ID

	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
	// 记录ID，逗号分隔

	RecordIds *string `json:"RecordIds,omitempty" name:"RecordIds"`
}

func (r *DeleteVpcDnsRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCountInfo struct {

	// 全部域名数量

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 返回的域名数量

	DomainTotal *int64 `json:"DomainTotal,omitempty" name:"DomainTotal"`
}

type DomainListFilters struct {

	// 过滤类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MessageDefault struct {

	// 返回码，正常为0

	Code *string `json:"Code,omitempty" name:"Code"`
	// 操作码说明

	Message *string `json:"Message,omitempty" name:"Message"`
}

type VpcDnsForwardRule struct {
	RuleId         *string   `json:"RuleId,omitempty" name:"RuleId"`
	Remark         *string   `json:"Remark,omitempty" name:"Remark"`
	DomainId       *string   `json:"DomainId,omitempty" name:"DomainId"`
	ForwardAddress []*string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
}

type VpcDnsForwardRuleDetail struct {
	RuleId         *string     `json:"RuleId,omitempty" name:"RuleId"`
	Remark         *string     `json:"Remark,omitempty" name:"Remark"`
	DomainId       *string     `json:"DomainId,omitempty" name:"DomainId"`
	DomainName     *string     `json:"DomainName,omitempty" name:"DomainName"`
	ForwardAddress []*string   `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
	VpcInfos       []*VpcInfos `json:"VpcInfos,omitempty" name:"VpcInfos"`
	CreatedOn      *string     `json:"CreatedOn,omitempty" name:"CreatedOn"`
	UpdatedOn      *string     `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
}

type CreateVpcDnsForwardRuleRequest struct {
	*tchttp.BaseRequest
	Remark         *string   `json:"Remark,omitempty" name:"Remark"`
	DomainIdList   []*string `json:"DomainIdList,omitempty" name:"DomainIdList"`
	ForwardAddress []*string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
}

type CreateVpcDnsForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 请求响应时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId  *string   `json:"RequestId,omitempty" name:"RequestId"`
		RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
	} `json:"Response"`
}

type ForwardRuleFilter struct {
	Name   *string   `json:"Name,omitempty" name:"Name"`
	Values []*string `json:"Values,omitempty" name:"Values"`
}

func (r *CreateVpcDnsForwardRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsForwardRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsForwardRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func (r *CreateVpcDnsForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsForwardRuleRequest struct {
	*tchttp.BaseRequest
	// 长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移
	Offset  *uint64              `json:"Offset,omitempty" name:"Offset"`
	Filters []*ForwardRuleFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeVpcDnsForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		CreatedAt       *string                    `json:"CreatedAt,omitempty" name:"CreatedAt"`
		RequestId       *string                    `json:"RequestId,omitempty" name:"RequestId"`
		ForwardRuleList []*VpcDnsForwardRuleDetail `json:"ForwardRuleList,omitempty" name:"ForwardRuleList"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsForwardRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsForwardRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsForwardRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func (r *DescribeVpcDnsForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsForwardRuleRequest struct {
	*tchttp.BaseRequest
	RuleId         *string   `json:"RuleId,omitempty" name:"RuleId"`
	ForwardAddress []*string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
	Remark         *string   `json:"Remark,omitempty" name:"Remark"`
}

type ModifyVpcDnsForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
		// 请求响应时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 规则ID
		RuleId         *string   `json:"RuleId,omitempty" name:"RuleId"`
		ForwardAddress []*string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
		Remark         *string   `json:"Remark,omitempty" name:"Remark"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsForwardRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsForwardRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsForwardRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func (r *ModifyVpcDnsForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsForwardRuleRequest struct {
	*tchttp.BaseRequest
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DeleteVpcDnsForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
		// 请求响应时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsForwardRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsForwardRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsForwardRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func (r *DeleteVpcDnsForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {
	// VpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Vpc所属地区

	Region *string `json:"Region,omitempty" name:"Region"`
}

type TagInfo struct {
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

type AccountVpcInfo struct {
	// vpc资源名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// VpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Vpc所属的地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// Vpc所属账号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type AccountVpcInfoOutput struct {
	// vpcid

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 关联账户的uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数值数组

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type PrivateZone struct {
	// 私有域id: zone-xxxxxxxx

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 私有域绑定VPC状态，未关联vpc：SUSPEND，已关联VPC：ENABLED，关联VPC失败：FAILED

	Status *string `json:"Status,omitempty" name:"Status"`
	// 转发规则名称

	ForwardRuleName *string `json:"ForwardRuleName,omitempty" name:"ForwardRuleName"`
	// 终端节点名称

	EndPointName *string `json:"EndPointName,omitempty" name:"EndPointName"`
	// 已删除的vpc

	DeletedVpcSet []*VpcInfo `json:"DeletedVpcSet,omitempty" name:"DeletedVpcSet"`
	// 记录数

	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
	// 绑定的Vpc列表

	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`
	// 域名递归解析状态：开通：ENABLED, 关闭，DISABLED

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
	// CNAME加速状态：开通：ENABLED, 关闭，DISABLED

	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
	// 域名所有者uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 转发规则类型：云上到云下，DOWN；云下到云上，UP

	ForwardRuleType *string `json:"ForwardRuleType,omitempty" name:"ForwardRuleType"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 标签键值对集合

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 绑定的关联账号的vpc列表

	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
	// 是否自定义TLD

	IsCustomTld *bool `json:"IsCustomTld,omitempty" name:"IsCustomTld"`
	// 转发的地址

	ForwardAddress *string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`
	// 域名数字ID

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

type PrivateZoneRecord struct {
	// 附加信息

	Extra *string `json:"Extra,omitempty" name:"Extra"`
	// 0暂停，1启用

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 私有域id: zone-xxxxxxxx

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
	// MX优先级

	MX *int64 `json:"MX,omitempty" name:"MX"`
	// 记录状态：ENABLED

	Status *string `json:"Status,omitempty" name:"Status"`
	// 记录id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 记录缓存时间

	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
	// 记录权重，值为1-100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 记录创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 记录更新时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
}

// ---- CreatePrivateZone ----

type CreatePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 创建私有域的同时，为其打上标签

	TagSet []*TagInfo `json:"TagSet,omitempty" name:"TagSet"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否开启子域名递归, ENABLED, DISABLED。默认值为ENABLED

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
	// 创建私有域的同时，将其关联至VPC

	Vpcs []*VpcInfo `json:"Vpcs,omitempty" name:"Vpcs"`
	// 创建私有域同时绑定关联账号的VPC

	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
	// 是否CNAME加速：ENABLED，DISABLED，默认值为ENABLED

	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
	// 是否强制绑定冲突的vpc

	ForceBindVpc *bool `json:"ForceBindVpc,omitempty" name:"ForceBindVpc"`
	// 域名，格式必须是标准的TLD

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 创建私有域的同时，将其关联至VPC

	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`
}

func (r *CreatePrivateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 私有域ID

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// 私有域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- DescribePrivateZone ----

type DescribePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 私有域id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribePrivateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 私有域详情

		PrivateZone *PrivateZone `json:"PrivateZone,omitempty" name:"PrivateZone"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- DescribePrivateZoneList ----

type DescribePrivateZoneListRequest struct {
	*tchttp.BaseRequest

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Filters
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrivateZoneListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		PrivateZoneSet []*PrivateZone `json:"PrivateZoneSet,omitempty" name:"PrivateZoneSet"`

		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

func (r *DescribePrivateZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- ModifyPrivateZone ----

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 是否开启子域名递归, ENABLED, DISABLED

	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
	// 是否开启CNAME加速：ENABLED, DISABLED

	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyPrivateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- DeletePrivateZone ----

type DeletePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有域ID数组，ZoneId 优先

	ZoneIdSet []*string `json:"ZoneIdSet,omitempty" name:"ZoneIdSet"`
}

func (r *DeletePrivateZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- CreatePrivateZoneRecord ----

type CreatePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 子域名，例如 "www", "m", "@"

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录值

	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
	// 记录权重，值为1-100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 记录类型："A", "AAAA", "CNAME", "MX", "TXT", "PTR"

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// MX优先级

	MX *int64 `json:"MX,omitempty" name:"MX"`
	// 记录缓存时间，取值1-86400s, 默认 600

	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *CreatePrivateZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateZoneRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录Id

		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- ModifyPrivateZoneRecord ----

type ModifyPrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 记录ID

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 记录类型："A", "AAAA", "CNAME", "MX", "TXT", "PTR"

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录值

	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
	// 记录权重，值为1-100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// MX优先级

	MX *int64 `json:"MX,omitempty" name:"MX"`
	// 记录缓存时间，取值1-86400s, 默认 600

	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *ModifyPrivateZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- DeletePrivateZoneRecord ----

type DeletePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 记录ID

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 记录ID数组，RecordId 优先

	RecordIdSet []*string `json:"RecordIdSet,omitempty" name:"RecordIdSet"`
}

func (r *DeletePrivateZoneRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateZoneRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- DescribePrivateZoneRecordList ----

type DescribePrivateZoneRecordListRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制数目, 最大200，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 私有域ID: zone-xxxxxxxx

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribePrivateZoneRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解析记录数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 解析记录列表

		RecordSet []*PrivateZoneRecord `json:"RecordSet,omitempty" name:"RecordSet"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivateZoneRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- ModifyPrivateZoneVpc ----

type ModifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest

	// 私有域账号关联的全部VPC列表

	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
	// 是否强制绑定冲突的vpc

	ForceBindVpc *bool `json:"ForceBindVpc,omitempty" name:"ForceBindVpc"`
	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 私有域关联的全部VPC列表

	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`
}

func (r *ModifyPrivateZoneVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 私有域ID

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// 解析域关联的VPC列表

		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`
		// 私有域账号关联的全部VPC列表

		AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// ---- ModifyRecordsStatus ----

type ModifyRecordsStatusRequest struct {
	*tchttp.BaseRequest

	// enabled：生效，disabled：失效

	Status *string `json:"Status,omitempty" name:"Status"`
	// 私有域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 解析记录ID列表

	RecordIds []*int64 `json:"RecordIds,omitempty" name:"RecordIds"`
}

func (r *ModifyRecordsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 私有域ID

		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
		// 解析记录ID列表

		RecordIds []*int64 `json:"RecordIds,omitempty" name:"RecordIds"`
		// enabled：生效，disabled：失效

		Status *string `json:"Status,omitempty" name:"Status"`
		// 请求返回时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

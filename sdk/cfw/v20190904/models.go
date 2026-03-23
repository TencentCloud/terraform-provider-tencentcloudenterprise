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

package v20190904

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeRiskAssetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计数据

		Data []*RiskAsset `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskAssetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwGroupSwitchRequest struct {
	*tchttp.BaseRequest

	// 是否操作全部开关&nbsp;0&nbsp;不操作全部开关，1&nbsp;操作全部开关

	AllSwitch *int64 `json:"AllSwitch,omitempty" name:"AllSwitch"`
	// 打开或关闭开关
	// 0：关闭开关
	// 1：打开开关

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 开关列表

	SwitchList []*FwGroupSwitch `json:"SwitchList,omitempty" name:"SwitchList"`
}

func (r *ModifyFwGroupSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwGroupSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，-1代表失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success代表成功，failed代表失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 删除成功后返回被删除策略的uuid

		RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderList struct {

	// 订单的开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 货币，国内站固定usd，国际站固定

	Currency *string `json:"Currency,omitempty" name:"Currency"`
	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 订单的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 场景Id

	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 计费产品详情

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 购买量，固定值1

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 资源隔离时间

	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`
	// 计费模式&nbsp;0&nbsp;后付费&nbsp;1&nbsp;预付费(包年包月)

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 产品id

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 所属产品&nbsp;p_

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 套餐规格，PayMode=1时有值

	ProductInfo *OrderProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 0表示默认状态(用户未设置，即初始状态)，&nbsp;1表示自动续费，2表示明确不自动续费(用户设置)

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 订单状态&nbsp;1&nbsp;正常&nbsp;2隔离&nbsp;3销毁

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 子产品&nbsp;sp_

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeAccountGroupsData struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 扩展信息

	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// id&nbsp;path

	IdPath *string `json:"IdPath,omitempty" name:"IdPath"`
	// id&nbsp;patch数组(只支持32位)

	IdPathArr []*int64 `json:"IdPathArr,omitempty" name:"IdPathArr"`
	// 导入类型

	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`
	// 是否已接入

	IsAccessed *int64 `json:"IsAccessed,omitempty" name:"IsAccessed"`
	// 是否叶子节点

	IsLeaf *bool `json:"IsLeaf,omitempty" name:"IsLeaf"`
	// 创建时间

	Itime *string `json:"Itime,omitempty" name:"Itime"`
	// miniIAM&nbsp;id

	MiniIamId *string `json:"MiniIamId,omitempty" name:"MiniIamId"`
	// 账户组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 名称path

	NamePath *string `json:"NamePath,omitempty" name:"NamePath"`
	// 组织id

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 父id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 父组织id

	ParentOrgId *string `json:"ParentOrgId,omitempty" name:"ParentOrgId"`
	// 是否该账户的直接权限

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 根节点

	RootGroupId *string `json:"RootGroupId,omitempty" name:"RootGroupId"`
	// 同步数据源

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 该分组下用户总数

	UserTotal *int64 `json:"UserTotal,omitempty" name:"UserTotal"`
	// 最后更新时间

	Utime *string `json:"Utime,omitempty" name:"Utime"`
}

type NatFwEipsInfo struct {

	// 弹性公网ip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// 所属的Nat网关Id

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// Nat网关名称

	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
}

type PresetAddressTemplateInfo struct {

	// 模版描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 模版内容

	IpString *string `json:"IpString,omitempty" name:"IpString"`
	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则数量

	RuleNum *int64 `json:"RuleNum,omitempty" name:"RuleNum"`
	// 模版类型&nbsp;1云产资源&nbsp;2威胁情报

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模版排序

	TmplIndex *int64 `json:"TmplIndex,omitempty" name:"TmplIndex"`
	// 模版类型&nbsp;1ip&nbsp;5域名

	TmplType *int64 `json:"TmplType,omitempty" name:"TmplType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeLogDeliverCongRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogDeliverCongRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliverCongRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProbeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProbeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpTopEvents struct {

	// 告警数

	AlarmCount *int64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
	// 拦截数据

	BlockCount *int64 `json:"BlockCount,omitempty" name:"BlockCount"`
	// 欺骗数据

	DeceptionCount *int64 `json:"DeceptionCount,omitempty" name:"DeceptionCount"`
	// ip数据

	IP *string `json:"IP,omitempty" name:"IP"`
	// 忽略数据

	IgnoreCount *int64 `json:"IgnoreCount,omitempty" name:"IgnoreCount"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DeleteBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlockIgnoreRuleNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageHistogramShow struct {

	// 数据

	Data []*IntArray `json:"Data,omitempty" name:"Data"`
	// 日期

	Dates []*string `json:"Dates,omitempty" name:"Dates"`
	// 存储类型

	StorageType []*string `json:"StorageType,omitempty" name:"StorageType"`
}

type CreateChooseVpcsRequest struct {
	*tchttp.BaseRequest

	// zone列表

	AllZoneList []*VpcZoneData `json:"AllZoneList,omitempty" name:"AllZoneList"`
	// vpc列表

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *CreateChooseVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateChooseVpcsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAcRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 规则Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 限速的类型，可以是具体的IP也可以是CIDR，比如192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 需要添加限速规则的防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
	// 流量控制策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 入向限制速度，单位是kbps

	InLimit *uint64 `json:"InLimit,omitempty" name:"InLimit"`
	// 限速类型
	// IP限速：IP
	// 端口限速：PORT
	// 当前仅支持IP限速

	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`
	// 出向限制速度，单位是kbps

	OutLimit *uint64 `json:"OutLimit,omitempty" name:"OutLimit"`
	// 流量控制策略的端口。取值：
	// -1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyVpcFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionZone struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	ZoneLst []*string `json:"ZoneLst,omitempty" name:"ZoneLst"`
	// 可用区列表中文

	ZoneZhLst []*string `json:"ZoneZhLst,omitempty" name:"ZoneZhLst"`
}

type DescribeCcnWithEdgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开白路由表的云联网CCN列表

		Data []*CcnInstance `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnWithEdgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnWithEdgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwSwitchRequest struct {
	*tchttp.BaseRequest

	// vpc间防火墙实例id列表，其中CfwInsIdList，CfwSideList只能传递一种。

	CfwInsIdList []*string `json:"CfwInsIdList,omitempty" name:"CfwInsIdList"`
	// 本端VPC与对端VPC组成的fwsId，其中CfwInsIdList，CfwSideList只能传递一种。

	CfwSideList []*string `json:"CfwSideList,omitempty" name:"CfwSideList"`
	// 开关，0：关闭，1：开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 是否跨租户

	Ispeer *int64 `json:"Ispeer,omitempty" name:"Ispeer"`
	// 实例的开关模式：1，同开同关；0，单开单关

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

func (r *ModifyVpcFwSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeWarnListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		Data []*IpsLogsData `json:"Data,omitempty" name:"Data"`
		// 判断来源列表

		DiffFromLists []*string `json:"DiffFromLists,omitempty" name:"DiffFromLists"`
		// 威胁类型名称列表

		EventNameLists []*string `json:"EventNameLists,omitempty" name:"EventNameLists"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeWarnListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeWarnListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 编辑成功后返回新策略ID列表

		RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupSimplifyRule struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)

	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`
	// 协议；TCP/UDP/ICMP/ANY

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 规则对应的唯一id

	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
	// 规则序号

	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
}

type DescribeAccessControlOverviewRequest struct {
	*tchttp.BaseRequest

	// 0：按方向，1：按策略

	Dimension *uint64 `json:"Dimension,omitempty" name:"Dimension"`
	// 0:&nbsp;7天，1:24小时

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAccessControlOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchMasterRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
}

func (r *SwitchMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZeroTrustVpc struct {

	// cidr

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 内网CVM数量

	PrivateCvmCount *uint64 `json:"PrivateCvmCount,omitempty" name:"PrivateCvmCount"`
	// 内网DB数量

	PrivateDbCount *uint64 `json:"PrivateDbCount,omitempty" name:"PrivateDbCount"`
	// 内网WEB防护数量

	PrivateWebCount *uint64 `json:"PrivateWebCount,omitempty" name:"PrivateWebCount"`
	// 公网CVM数量

	PublicCvmCount *uint64 `json:"PublicCvmCount,omitempty" name:"PublicCvmCount"`
	// 公网DB数量

	PublicDbCount *uint64 `json:"PublicDbCount,omitempty" name:"PublicDbCount"`
	// 公网WEB防护数量

	PublicWebCount *uint64 `json:"PublicWebCount,omitempty" name:"PublicWebCount"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 0:完成,&nbsp;1:开启中，2：关闭中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 零信任防护开关
	// 1&nbsp;打开
	// 0&nbsp;关闭

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
	// 0&nbsp;正常&nbsp;1&nbsp;隐藏开关或开关置灰

	SwitchHide *int64 `json:"SwitchHide,omitempty" name:"SwitchHide"`
	// vpc唯一id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeBorderACLListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 不算筛选条数的总条数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// 访问控制列表数据

		Data []*BorderACL `json:"Data,omitempty" name:"Data"`
		// 访问控制规则全部启用1/全部停用&nbsp;0

		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBorderACLListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderACLListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDefenseSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcConnectionsTrendsRequest struct {
	*tchttp.BaseRequest

	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Vpc间防火墙组id

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// Vpc间防火墙示例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcConnectionsTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcConnectionsTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建结果,0成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一Id

		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcProtectDetail struct {

	// 目的vpc&nbsp;cidr

	DstCidr *string `json:"DstCidr,omitempty" name:"DstCidr"`
	// 目的vpc&nbsp;ID

	DstId *string `json:"DstId,omitempty" name:"DstId"`
	// 目的vpc名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 目的vpc地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// vpc边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// vpc边名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 0&nbsp;观察&nbsp;1&nbsp;拦截&nbsp;2&nbsp;严格&nbsp;3&nbsp;关闭

	Idpsaction *int64 `json:"Idpsaction,omitempty" name:"Idpsaction"`
	// vpc&nbsp;cidr

	SrcCidr *string `json:"SrcCidr,omitempty" name:"SrcCidr"`
	// 源vpc&nbsp;ID

	SrcId *string `json:"SrcId,omitempty" name:"SrcId"`
	// 源vpc名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 源vpc地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
}

type DescribeCfwVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1:&nbsp;dpdk&nbsp;模式，新引擎&nbsp;2:&nbsp;内核版模式，旧引擎

		EngineMode *int64 `json:"EngineMode,omitempty" name:"EngineMode"`
		// 是否需要升级引擎支持nat拨测探测&nbsp;1需要升级&nbsp;&nbsp;&nbsp;0不需要升级&nbsp;

		NeedProbeEngineUpdate *int64 `json:"NeedProbeEngineUpdate,omitempty" name:"NeedProbeEngineUpdate"`
		// 1需要升级&nbsp;0不需要升级&nbsp;

		NeedTcEngineUpdate *int64 `json:"NeedTcEngineUpdate,omitempty" name:"NeedTcEngineUpdate"`
		// 引擎的具体版本号

		Version *string `json:"Version,omitempty" name:"Version"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipInfo struct {

	// 防火墙出口绑定类型:0,cvm类型;1,子网类型;2,vpc类型

	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
	// 弹性公网ip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// 防火墙出口绑定策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DescribeAclTagRequest struct {
	*tchttp.BaseRequest

	// 访问控制标签类型：
	// border&nbsp;互联网边界
	// nat&nbsp;&nbsp;NAT防火墙
	// vpc&nbsp;VPC防火墙
	// sg&nbsp;企业安全组
	// dns&nbsp;DNS防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

func (r *DescribeAclTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogTopIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据集合

		Data []*IpTopEvents `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTLogTopIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogTopIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSessionConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 会话连接数统计列表

		Data []*SessionConnections `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcSessionConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSessionConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSwitchListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcSwitchListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGuideUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGuideUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGuideUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveAutoBackUpSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveAutoBackUpSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveAutoBackUpSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcFwInstanceName struct {

	// VPC实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// VPC实例名称

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// 0：旧VPC&nbsp;防火墙模式&nbsp;，&nbsp;1:&nbsp;新CCN&nbsp;VPC防火墙模式

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
}

type CreateGuideOneClickScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGuideOneClickScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGuideOneClickScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指定边启用的内网规则数量

		EdgeEnableRuleNum *uint64 `json:"EdgeEnableRuleNum,omitempty" name:"EdgeEnableRuleNum"`
		// 指定边的内网规则数量

		EdgeRuleNum *uint64 `json:"EdgeRuleNum,omitempty" name:"EdgeRuleNum"`
		// 全局范围启用的内网规则数量

		GlobalEnableRuleNum *uint64 `json:"GlobalEnableRuleNum,omitempty" name:"GlobalEnableRuleNum"`
		// 全局范围的内网规则数量

		GlobalRuleNum *uint64 `json:"GlobalRuleNum,omitempty" name:"GlobalRuleNum"`
		// 预操作是否超过配额，1：超过，0：未超过

		IsOverLimit *int64 `json:"IsOverLimit,omitempty" name:"IsOverLimit"`
		// 最近一次规则备份的时间

		LastBackupTime *string `json:"LastBackupTime,omitempty" name:"LastBackupTime"`
		// 预操作需要额外配额数

		NeedQuota *int64 `json:"NeedQuota,omitempty" name:"NeedQuota"`
		// 内网间规则配额

		Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
		// 剩余通用规则扩展数

		RemainExtensionNum *int64 `json:"RemainExtensionNum,omitempty" name:"RemainExtensionNum"`
		// 剩余配额数

		RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 已使用规格数

		UsedNum *uint64 `json:"UsedNum,omitempty" name:"UsedNum"`
		// Ipv4开启规则总数

		V4EnableRuleTotal *int64 `json:"V4EnableRuleTotal,omitempty" name:"V4EnableRuleTotal"`
		// Ipv4规则当前最大序列号

		V4RuleMaxSeq *int64 `json:"V4RuleMaxSeq,omitempty" name:"V4RuleMaxSeq"`
		// Ipv4规则总数

		V4RuleTotal *int64 `json:"V4RuleTotal,omitempty" name:"V4RuleTotal"`
		// Ipv6开启规则总数

		V6EnableRuleTotal *int64 `json:"V6EnableRuleTotal,omitempty" name:"V6EnableRuleTotal"`
		// Ipv6规则当前最大序列号

		V6RuleMaxSeq *int64 `json:"V6RuleMaxSeq,omitempty" name:"V6RuleMaxSeq"`
		// Ipv6规则总数

		V6RuleTotal *int64 `json:"V6RuleTotal,omitempty" name:"V6RuleTotal"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetWebScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态值&nbsp;0：成功，1&nbsp;执行扫描中,其他：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetWebScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetWebScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceAclRulesRequest struct {
	*tchttp.BaseRequest

	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号

	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitempty" name:"RuleChangeItems"`
}

func (r *ModifySequenceAclRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySequenceAclRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonIdName struct {

	// 资源ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资源名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeStrictModeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否自动封禁域名

		AutoBlockDomain *int64 `json:"AutoBlockDomain,omitempty" name:"AutoBlockDomain"`
		// 是否自动封禁IP

		AutoBlockIp *int64 `json:"AutoBlockIp,omitempty" name:"AutoBlockIp"`
		// 封禁时长类型，&nbsp;Fixed固定时长，Step阶梯时长

		BlockDurationType *string `json:"BlockDurationType,omitempty" name:"BlockDurationType"`
		// 第一阶梯时长，单位分钟

		FirstBlockStepDuration *int64 `json:"FirstBlockStepDuration,omitempty" name:"FirstBlockStepDuration"`
		// 固定封禁时长

		FixedBlockDuration *int64 `json:"FixedBlockDuration,omitempty" name:"FixedBlockDuration"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 第二阶梯时长，单位分钟

		SecondBlockStepDuration *int64 `json:"SecondBlockStepDuration,omitempty" name:"SecondBlockStepDuration"`
		// 第三阶梯时长，单位分钟

		ThirdBlockStepDuration *int64 `json:"ThirdBlockStepDuration,omitempty" name:"ThirdBlockStepDuration"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStrictModeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrictModeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRuleLimitsForWidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例是否修改过&nbsp;1是&nbsp;0否&nbsp;默认0

		FwInsIsFix *int64 `json:"FwInsIsFix,omitempty" name:"FwInsIsFix"`
		// 不同带宽范围的规则最大数限制

		RuleLimit []*WidthRuleLimit `json:"RuleLimit,omitempty" name:"RuleLimit"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRuleLimitsForWidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRuleLimitsForWidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewEvent struct {

	// 事件集合

	EventList []*GetOverviewEventStruct `json:"EventList,omitempty" name:"EventList"`
	// 事件类型

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 创建规则数据

	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`
	// 边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 访问控制规则状态

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// portScan:&nbsp;来自于端口扫描,&nbsp;patchImport:&nbsp;来自于批量导入

	From *string `json:"From,omitempty" name:"From"`
	// NAT实例ID,&nbsp;参数Area存在的时候这个必传

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 0：添加，1：覆盖

	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`
	// 0：添加，1：插入

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateAcRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAcRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStatusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// ip状态信息

		StatusList []*IPDefendStatus `json:"StatusList,omitempty" name:"StatusList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStatusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwInsInfo struct {

	// 租户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 机器id

	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`
	// 错误等级

	ErrLevel *int64 `json:"ErrLevel,omitempty" name:"ErrLevel"`
	// 错误类型

	ErrType *string `json:"ErrType,omitempty" name:"ErrType"`
	// 防火墙类型

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 实例id

	InsId *string `json:"InsId,omitempty" name:"InsId"`
	// 租户名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type SwitchError struct {

	// 开关唯一标识

	ErrIns *string `json:"ErrIns,omitempty" name:"ErrIns"`
	// 错误类型区分

	ErrKey *string `json:"ErrKey,omitempty" name:"ErrKey"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 错误时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type ModifyFwGroupSwitchAllRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;关闭;&nbsp;1&nbsp;开启

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ModifyFwGroupSwitchAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwGroupSwitchAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebServiceInfo struct {

	// 接入域名

	AccessDomain *string `json:"AccessDomain,omitempty" name:"AccessDomain"`
	// 接入类型:1自有域名接入,2防火墙域名接入

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 权限模式：1，精细模式，0，极简模式

	AclMode *int64 `json:"AclMode,omitempty" name:"AclMode"`
	// web服务受到的攻击次数

	AttackCount *int64 `json:"AttackCount,omitempty" name:"AttackCount"`
	// 授权的运维用户，多个用户用逗号隔开

	AuthUsers *string `json:"AuthUsers,omitempty" name:"AuthUsers"`
	// 组件信息

	Component *string `json:"Component,omitempty" name:"Component"`
	// 频繁访问的运维用户，多个用户用逗号隔开

	FrequenVisit *string `json:"FrequenVisit,omitempty" name:"FrequenVisit"`
	// 频繁访问的运维用户，多个用户用逗号隔开

	FrequentVisit *string `json:"FrequentVisit,omitempty" name:"FrequentVisit"`
	// 服务来源，有三种途径，分别是：Scan：资产扫描；Manual：手动添加；FlowAware：流量感知

	From *string `json:"From,omitempty" name:"From"`
	// 添加时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 承载这个web服务所用到的资产，可能是cvm、clb等

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Instance字段对应的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// ip状态，true为ip和实例对应的ip一致，&nbsp;false：为ip与实例对应的ip（包括内外网ip）存在不一致

	IpStatus *bool `json:"IpStatus,omitempty" name:"IpStatus"`
	// 1接入成功&nbsp;2未接入&nbsp;3正在接入&nbsp;4接入失败

	IsAccessed *int64 `json:"IsAccessed,omitempty" name:"IsAccessed"`
	// true代表新增实例，false非新增

	IsNew *bool `json:"IsNew,omitempty" name:"IsNew"`
	// true正在做漏洞扫描，false没在做漏洞扫描

	IsScanning *bool `json:"IsScanning,omitempty" name:"IsScanning"`
	// 接入内网域名

	PrivateAccessDomain *string `json:"PrivateAccessDomain,omitempty" name:"PrivateAccessDomain"`
	// 接入防护类型

	ProtectType *uint64 `json:"ProtectType,omitempty" name:"ProtectType"`
	// 最近访问的运维用户，多个用户用逗号隔开

	RecentVisit *string `json:"RecentVisit,omitempty" name:"RecentVisit"`
	// web服务所在的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// web服务备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// web服务对应的IP地址

	ServiceAddress *string `json:"ServiceAddress,omitempty" name:"ServiceAddress"`
	// web服务对应的域名

	ServiceDomain *string `json:"ServiceDomain,omitempty" name:"ServiceDomain"`
	// web服务对应的端口

	ServicePort *int64 `json:"ServicePort,omitempty" name:"ServicePort"`
	// public：公网开放服务；&nbsp;private：内网开放服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 授权用户最后访问web服务的时间

	UserLastVisitTime *string `json:"UserLastVisitTime,omitempty" name:"UserLastVisitTime"`
	// web服务存在的漏洞

	VULNCount *int64 `json:"VULNCount,omitempty" name:"VULNCount"`
	// web服务的访问次数

	VisitCount *int64 `json:"VisitCount,omitempty" name:"VisitCount"`
	// 所属vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 所属vpcid名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// web服务存在的漏洞数量

	VulnCount *int64 `json:"VulnCount,omitempty" name:"VulnCount"`
	// web服务的id

	WebId *string `json:"WebId,omitempty" name:"WebId"`
}

type ModifyDnsResolveIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDnsResolveIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnsResolveIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetOutOverviewRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeInternetOutOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetOutOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSFWAssetStat struct {

	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资产内网ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 峰值请求频率

	RequestQpsMax *uint64 `json:"RequestQpsMax,omitempty" name:"RequestQpsMax"`
	// 开关状态，1开启，0关闭

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 累计请求次数

	TotalRequestQuantity *uint64 `json:"TotalRequestQuantity,omitempty" name:"TotalRequestQuantity"`
	// VPC&nbsp;Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC&nbsp;名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeDefenceApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenceApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作总条数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAcRuleAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportWeChatSpec struct {

	// 超出额度数量

	ExceedSpec *int64 `json:"ExceedSpec,omitempty" name:"ExceedSpec"`
	// 导入手机号数量

	PhoneLst *int64 `json:"PhoneLst,omitempty" name:"PhoneLst"`
	// true:配额超量

	WxSpecStatus *bool `json:"WxSpecStatus,omitempty" name:"WxSpecStatus"`
}

type DeleteBlockIgnoreRuleAllRequest struct {
	*tchttp.BaseRequest

	// 方向：1互联网入站，0互联网出站，3内网，空&nbsp;全部方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 规则类型：1封禁，2放通

	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`
	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// blocklist&nbsp;封禁列表
	// whitelist&nbsp;白名单列表

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DeleteBlockIgnoreRuleAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenceApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeDefenceApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetScanFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetScanFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetScanFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatFwFilter struct {

	// 过滤的内容，以','&nbsp;分隔

	FilterContent *string `json:"FilterContent,omitempty" name:"FilterContent"`
	// 过滤的类型，例如实例id

	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type ProductStatusDataType struct {

	// 产品对应分层&nbsp;1:&nbsp;第一道防线&nbsp;2:&nbsp;第二道防线&nbsp;3:&nbsp;第三道防线&nbsp;4:&nbsp;加固防线

	ProductLayer *int64 `json:"ProductLayer,omitempty" name:"ProductLayer"`
	// 产品名称&nbsp;CFW：云防火墙&nbsp;WAF：Web应用防火墙&nbsp;CWP：主机安全&nbsp;SOC：安全运营中心&nbsp;TCSS：容器安全

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品是否已上线&nbsp;0:&nbsp;未上线&nbsp;1:&nbsp;已上线

	ProductOnline *int64 `json:"ProductOnline,omitempty" name:"ProductOnline"`
	// 产品当前状态&nbsp;0可试用，1未防护，2防护中（已购买），3试用中，4已过期

	ProductStatus *int64 `json:"ProductStatus,omitempty" name:"ProductStatus"`
	// 返回消息，一般为不可试用原因

	TrialMessage *string `json:"TrialMessage,omitempty" name:"TrialMessage"`
	// 是否提供试用功能&nbsp;0不提供&nbsp;1提供

	TrialStatus *int64 `json:"TrialStatus,omitempty" name:"TrialStatus"`
}

type DescribeAssetStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data []*AssetStatic `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1展示迭代发布横幅

		IsPublishVersion *int64 `json:"IsPublishVersion,omitempty" name:"IsPublishVersion"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGlobalSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TiNews struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 详情链接

	DetailUrl *string `json:"DetailUrl,omitempty" name:"DetailUrl"`
	// 共享链接

	ShareUrl *string `json:"ShareUrl,omitempty" name:"ShareUrl"`
	// 0&nbsp;隐藏&nbsp;1&nbsp;展示

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type CreateGuideIdpsInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateGuideIdpsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGuideIdpsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data *AssetOverviewEvent `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeLogApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogsOfflineRequest struct {
	*tchttp.BaseRequest

	// Cos存储桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// Cos存储桶地域

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// 压缩方式，.zip、.tar.gz、.tar.zst、.tar.lz4、传空不压缩

	CompressionFormat *string `json:"CompressionFormat,omitempty" name:"CompressionFormat"`
	// 数据格式

	DataFormat *string `json:"DataFormat,omitempty" name:"DataFormat"`
	// 筛选结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 日志类型标识

	Index *string `json:"Index,omitempty" name:"Index"`
	// 日志数量，传0全部导出

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 日志排序，desc时间降序，asc时间升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 筛选开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 文件过期时长，1一天，7七天，-1永久

	StorageDays *int64 `json:"StorageDays,omitempty" name:"StorageDays"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *ExportLogsOfflineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsOfflineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicySet struct {

	// &nbsp;出站规则。

	Egress []*SecurityGroupPolicy `json:"Egress,omitempty" name:"Egress"`
	// &nbsp;入站规则。

	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty" name:"Ingress"`
}

type UnHandleEventDetail struct {

	// 安全事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 未处置事件数量

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type ZeroTrustAuthorityOverview struct {

	// 访问数量

	AccessCount *int64 `json:"AccessCount,omitempty" name:"AccessCount"`
	// 拦截数量

	DropCount *int64 `json:"DropCount,omitempty" name:"DropCount"`
	// 命中数量

	HitCount *int64 `json:"HitCount,omitempty" name:"HitCount"`
	// 规则数量

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 规则开启数量

	RuleEnableCount *int64 `json:"RuleEnableCount,omitempty" name:"RuleEnableCount"`
}

type CreateVpcFwInstanceRequest struct {
	*tchttp.BaseRequest

	// 云联网模式下CcnId；仅CCN云联网模式使用

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 部署方式&nbsp;1：单地域部署；0：全地域部署。该字段仅CCN云联网模式场景有效，默认全地域部署

	DeployType *int64 `json:"DeployType,omitempty" name:"DeployType"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 部署地域信息

	FwDeploy []*FwDeploy `json:"FwDeploy,omitempty" name:"FwDeploy"`
	// 模式&nbsp;1：CCN云联网模式；0：私有网络模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 防火墙实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 私有网络模式下接入的VpcId列表；仅私有网络模式使用

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
}

func (r *CreateVpcFwInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcFwInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveOfflineExportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveOfflineExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveOfflineExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthBannerRequest struct {
	*tchttp.BaseRequest

	// 串行地域，
	// 传空为旁路

	SerialRegion *string `json:"SerialRegion,omitempty" name:"SerialRegion"`
}

func (r *DescribeBandWidthBannerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandWidthBannerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideNetFlowSwitchAPIRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuideNetFlowSwitchAPIRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideNetFlowSwitchAPIRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEdgeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySysLogAddrConfigRequest struct {
	*tchttp.BaseRequest

	// 地址配置

	AddrConfigList []*SysLogAddrConfig `json:"AddrConfigList,omitempty" name:"AddrConfigList"`
}

func (r *ModifySysLogAddrConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySysLogAddrConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeCommonRequest struct {
	*tchttp.BaseRequest

	// AutoRebuild、DelayRebuild
	//

	NoticeCommon *NoticeCommon `json:"NoticeCommon,omitempty" name:"NoticeCommon"`
}

func (r *ModifyNoticeCommonRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeCommonRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*InstanceInfo `json:"Data,omitempty" name:"Data"`
		// 返回数据总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域集合

		ZoneList []*AssetZone `json:"ZoneList,omitempty" name:"ZoneList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNatAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功后返回新策略ID列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNatAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EniBindInfo struct {

	// 网卡id

	EniId *string `json:"EniId,omitempty" name:"EniId"`
	// 网卡关联的私有ip信息

	IpInfo *string `json:"IpInfo,omitempty" name:"IpInfo"`
	// 网卡类型，1，主网卡；0，辅助网卡

	Role *string `json:"Role,omitempty" name:"Role"`
}

type CfwUpdateStatus struct {

	// 实例升级状态

	InsUpdateStatus []*CfwInsUpdateStatus `json:"InsUpdateStatus,omitempty" name:"InsUpdateStatus"`
	// 当前最新版本

	LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
	// 整体可升级状态，0：无需升级，1：可升级

	TotalEnable *int64 `json:"TotalEnable,omitempty" name:"TotalEnable"`
	// 整体升级状态，0：已完成，1：正在升级

	TotalStatus *int64 `json:"TotalStatus,omitempty" name:"TotalStatus"`
}

type DescribeLogsResultAsyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询状态，0查询已完成，1查询进行中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsResultAsyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsResultAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyAuthorityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyAuthorityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyAuthorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAclRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功后返回被删除策略的uuid列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveAclRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IocListData struct {

	// 只能为0或者1&nbsp;&nbsp;&nbsp;0代表出站&nbsp;1代表入站

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 待处置域名，IP/Domain字段二选一

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 待处置IP地址，IP/Domain字段二选一

	IP *string `json:"IP,omitempty" name:"IP"`
}

type OperatorLogsData struct {

	// 操作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 操作类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 客户端ip

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 更多的信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 威胁等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 操作账号

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 运营商

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 防火墙类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type SecurityGroupOrderIndexData struct {

	// 企业安全组规则更新目标执行顺序

	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
	// 企业安全组规则当前执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
}

type WebAssetsZoneInfo struct {

	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type CreateBlockIgnoreRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回

		List []*IocListData `json:"List,omitempty" name:"List"`
		// 所需配额

		RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBlockIgnoreRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBlockIgnoreRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 描述过滤

	Description *string `json:"Description,omitempty" name:"Description"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 规则Id过滤

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeAclRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertAd struct {

	// 链接信息

	AlertLink []*AlertLink `json:"AlertLink,omitempty" name:"AlertLink"`
	// 序号

	Id *string `json:"Id,omitempty" name:"Id"`
	// tag文案

	TagText *string `json:"TagText,omitempty" name:"TagText"`
	// 类型

	TagType *int64 `json:"TagType,omitempty" name:"TagType"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type DescribeLogsAsyncRequest struct {
	*tchttp.BaseRequest

	// 筛选结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 日志类型标识

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeLogsAsyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsAsyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstancesInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例卡片信息数组

		NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitempty" name:"NatinsLst"`
		// nat&nbsp;防火墙个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstancesInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstancesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotifyTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Info *string `json:"Info,omitempty" name:"Info"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotifyTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotifyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BetaTask struct {

	// 规则类型&nbsp;1border&nbsp;2nat&nbsp;3vpc&nbsp;4sg

	AclType *int64 `json:"AclType,omitempty" name:"AclType"`
	// 预留字段处理特殊异常

	Exception *string `json:"Exception,omitempty" name:"Exception"`
	// 执行时间

	ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
	// 执行次数

	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`
	// 任务异常状态:&nbsp;0无意义&nbsp;1任务异常&nbsp;2无可执行任务对象

	FailedCode *int64 `json:"FailedCode,omitempty" name:"FailedCode"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 任务状态&nbsp;1正常&nbsp;2异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateAlertCenterIsolateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码：
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息：
		// success&nbsp;成功
		// 其他

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 处置状态码：
		// 0&nbsp;&nbsp;处置成功
		// -1&nbsp;通用错误，不用处理
		// -3&nbsp;表示重复，需重新刷新列表
		// 其他

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertCenterIsolateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterIsolateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwEngineZoneListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎区域信息列表

		Data []*UserEngineZoneInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwEngineZoneListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEngineZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplayUrlRequest struct {
	*tchttp.BaseRequest

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 系统类型

	OsType *int64 `json:"OsType,omitempty" name:"OsType"`
}

func (r *DescribeReplayUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplayUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEdgeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vpc防火墙边状态

		Data []*EdgeStatus `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEdgeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEdgeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleItem struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 告警Id

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 规则序号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 协议端口组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 访问控制策略的端口。取值：&nbsp;-1/-1：全部端口&nbsp;80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：&nbsp;TCP&nbsp;UDP&nbsp;ICMP&nbsp;ANY&nbsp;HTTP&nbsp;HTTPS&nbsp;HTTP/HTTPS&nbsp;SMTP&nbsp;SMTPS&nbsp;SMTP/SMTPS&nbsp;FTP&nbsp;DNS

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：&nbsp;accept：放行&nbsp;drop：拒绝&nbsp;log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 0，正常规则添加；1，入侵检测添加

	RuleSource *int64 `json:"RuleSource,omitempty" name:"RuleSource"`
	// all

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 访问源示例：&nbsp;net：IP/CIDR(192.168.0.2)

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源类型：入向规则时类型可以为&nbsp;ip,net,template,location；出向规则时可以为&nbsp;ip,net,template,instance,group,tag

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 访问目的示例：&nbsp;net：IP/CIDR(192.168.0.2)&nbsp;domain：域名规则，例如*.domain.com

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为&nbsp;&nbsp;ip,net,domain,template,location

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则对应的唯一id，创建规则时无需填写

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type StringValueText struct {

	// 描述

	Text *string `json:"Text,omitempty" name:"Text"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAssetSyncRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetSyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosBucketListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCosBucketListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosBucketListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwJoinInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙接入的网络实例列表

		Data []*VpcFwJoinInstance `json:"Data,omitempty" name:"Data"`
		// VPC防火墙接入的网络实例列表

		JoinData *VpcFwJoinInstance `json:"JoinData,omitempty" name:"JoinData"`
		// VPC防火墙个数

		JoinTotal *int64 `json:"JoinTotal,omitempty" name:"JoinTotal"`
		// VPC防火墙个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwJoinInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwJoinInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAclRulesRequest struct {
	*tchttp.BaseRequest

	// Acl类型

	AclType *string `json:"AclType,omitempty" name:"AclType"`
	// 是否全部导出

	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`
	// 自定义导出列

	ExportFields *string `json:"ExportFields,omitempty" name:"ExportFields"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ExportAclRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAclRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwUserStatusData struct {

	// 待处理安全事件数

	AlertCount *uint64 `json:"AlertCount,omitempty" name:"AlertCount"`
	// 防护资产数

	ProtectAssets *uint64 `json:"ProtectAssets,omitempty" name:"ProtectAssets"`
	// 防火墙用户状态：1(未使用cfw)、2(开启cfw)、3(过期用户)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 自动拦截网络攻击数量

	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
}

type CreateIdsWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 目的IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 白名单生效防火墙范围：
	// 1&nbsp;边界防火墙
	// 2&nbsp;nat防火墙
	// 4&nbsp;vpc防火墙
	// 7&nbsp;=&nbsp;1+2+4&nbsp;&nbsp;所有防火墙

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// 入侵防御规则ID

	IdsRuleId *string `json:"IdsRuleId,omitempty" name:"IdsRuleId"`
	// 源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 白名单类型：
	// src&nbsp;针对源放通
	// dst&nbsp;针对目的放通
	// srcdst&nbsp;针对源和目的放通

	WhiteRuleType *string `json:"WhiteRuleType,omitempty" name:"WhiteRuleType"`
}

func (r *CreateIdsWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIdsWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHoneyPotListRequest struct {
	*tchttp.BaseRequest

	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 探针Id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 搜索参数，json格式字符串，空则传"{}"，蜜罐服务：service_type，交互类型：interaction_type，地域：region，蜜罐Id：service_id，蜜罐名称：service_name，诱饵：bait，开关：switch，状态：status，模糊搜索：common

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 排序参数，探针数量从小到大probe_count:asc，探针数量从大到小probe_count:desc，命中次数从小到大hit_times:asc，命中次数从大到小hit_times:desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeHoneyPotListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHoneyPotListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapDownUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownUrl []*string `json:"DownUrl,omitempty" name:"DownUrl"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePcapDownUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePcapDownUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIgnoreRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertLogEventNameSelectListRequest struct {
	*tchttp.BaseRequest

	// 告警来源，Idps规则告警，Honeypot蜜罐告警

	AlertSourceType *string `json:"AlertSourceType,omitempty" name:"AlertSourceType"`
	// 资产实例ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 方向，1入站，0出站，3横向移动，4主机事件仅AlertSourceType=Honeypot时生效

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 检索结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 检索开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeAlertLogEventNameSelectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertLogEventNameSelectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreAsyncTaskErrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态0成功&nbsp;其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIgnoreAsyncTaskErrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreAsyncTaskErrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 开关状态，true代表开启，false代表关闭

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 开关类型，比如dns代表查询nat防火墙的dns开关

	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
}

func (r *ModifySwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdcIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地专用集群列表

		CdcLst []*DescCdc `json:"CdcLst,omitempty" name:"CdcLst"`
		// 列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdcIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClsDeliverTaskRequest struct {
	*tchttp.BaseRequest

	// 新SecurityID

	SecurityID *string `json:"SecurityID,omitempty" name:"SecurityID"`
	// 新SecurityKey

	SecurityKey *string `json:"SecurityKey,omitempty" name:"SecurityKey"`
	// 来源&nbsp;区分cls投递和企业安全组日志

	SourceFrom *string `json:"SourceFrom,omitempty" name:"SourceFrom"`
}

func (r *ModifyClsDeliverTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClsDeliverTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogpushTypeInfo struct {

	// 提示

	Alert *int64 `json:"Alert,omitempty" name:"Alert"`
	// 自定义

	Custom *int64 `json:"Custom,omitempty" name:"Custom"`
	// 自定义选项

	CustomInfo []*string `json:"CustomInfo,omitempty" name:"CustomInfo"`
	// 错误

	Error *int64 `json:"Error,omitempty" name:"Error"`
	// 警告

	Warning *int64 `json:"Warning,omitempty" name:"Warning"`
}

type DeleteIdsWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 入侵防御白名单id
	// 参考DescribeIdsWhiteRule接口返回的Id字段

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteIdsWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIdsWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleLogsData struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 1:出站，0：入站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源端口

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 命中时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 规则uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type NetFlowDomainTopInfo struct {

	// 域名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 流量

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeFwSwitchOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFwSwitchOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwSwitchOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatProtectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT防火墙接入的开关列表

		Data []*NatProtectSwitchDetail `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatProtectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatProtectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdpsInfo struct {

	// 是否赠送过

	IsGive *bool `json:"IsGive,omitempty" name:"IsGive"`
	// 是否开启拦截模式

	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
	// 是否付费用户

	IsPay *bool `json:"IsPay,omitempty" name:"IsPay"`
	// 拦截模式开启率

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 0：高级版&nbsp;1：企业版&nbsp;2：旗舰版&nbsp;3：IPS版

	VersionType *int64 `json:"VersionType,omitempty" name:"VersionType"`
}

type DescribeRiskAssetsRequest struct {
	*tchttp.BaseRequest

	// 统计类型：0:&nbsp;安全事件，1:暴露端口，2:暴露漏洞

	Field *string `json:"Field,omitempty" name:"Field"`
	// 时间类型：0:&nbsp;7天，1:24小时

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeRiskAssetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskAssetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwInsUpdateStatus struct {

	// 是否可回退，0，不可回退；1，可回退

	BackEnable *int64 `json:"BackEnable,omitempty" name:"BackEnable"`
	// 实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 实例名称

	CfwInsName *string `json:"CfwInsName,omitempty" name:"CfwInsName"`
	// 当前版本

	CurVersion *string `json:"CurVersion,omitempty" name:"CurVersion"`
	// 防火墙类型，ew，vpc间防火墙；nat，nat防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 上次更新时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 低版本风险，是否需要提示强制升级，1：需要；0：不需要

	NeedForceUpdate *int64 `json:"NeedForceUpdate,omitempty" name:"NeedForceUpdate"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 预约任务id

	ReserveId *int64 `json:"ReserveId,omitempty" name:"ReserveId"`
	// 预约升级时间

	ReserveTime *string `json:"ReserveTime,omitempty" name:"ReserveTime"`
	// 预约版本

	ReserveVersion *string `json:"ReserveVersion,omitempty" name:"ReserveVersion"`
	// 引擎预约升级版本状态

	ReserveVersionState *string `json:"ReserveVersionState,omitempty" name:"ReserveVersionState"`
	// 可升级状态，1：可升级；0：已是最新版无需升级

	UpdateEnable *int64 `json:"UpdateEnable,omitempty" name:"UpdateEnable"`
	// 升级状态，0：升级完成；1：正在升级

	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`
}

type RuleImpact struct {

	// 互联网

	InternetRule []*RuleImpactInfo `json:"InternetRule,omitempty" name:"InternetRule"`
	// nat

	NatRule []*RuleImpactInfo `json:"NatRule,omitempty" name:"NatRule"`
	// 规则数

	RuleTotalCount *int64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
	// vpc

	VpcRule []*RuleImpactInfo `json:"VpcRule,omitempty" name:"VpcRule"`
	// 安全组

	SecurityGroupRule []*RuleImpactInfo `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
}

type DescribeNtaFileDownUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownUrl []*string `json:"DownUrl,omitempty" name:"DownUrl"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNtaFileDownUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNtaFileDownUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetOverviewRequest struct {
	*tchttp.BaseRequest

	// 结束

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeAssetOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBetaTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务集合

		Data []*BetaTask `json:"Data,omitempty" name:"Data"`
		// 自动化助手配额

		Quota *int64 `json:"Quota,omitempty" name:"Quota"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBetaTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewNatCheckInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关型号大型使用量

		LargeUseNum *uint64 `json:"LargeUseNum,omitempty" name:"LargeUseNum"`
		// 网关型号中型使用量

		MediumUseNum *uint64 `json:"MediumUseNum,omitempty" name:"MediumUseNum"`
		// 可创建实例地域列表&nbsp;根据环境真实地域

		RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
		// 网关型号小型使用量

		SmallUseNum *uint64 `json:"SmallUseNum,omitempty" name:"SmallUseNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewNatCheckInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewNatCheckInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙接入的开关列表

		FwSwitchLst []*VpcFwSwitch `json:"FwSwitchLst,omitempty" name:"FwSwitchLst"`
		// VPC防火墙开关个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySysLogAddrConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;1异常&nbsp;0正常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySysLogAddrConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySysLogAddrConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwGroupSwitchShow struct {

	// 0-非sase实例，忽略，1-未绑定状态，2-已绑定

	AttachWithEdge *int64 `json:"AttachWithEdge,omitempty" name:"AttachWithEdge"`
	// 开关是否处于bypass状态
	// 0：正常状态
	// 1：bypass状态

	BypassStatus *int64 `json:"BypassStatus,omitempty" name:"BypassStatus"`
	// 连接ID

	ConnectId *string `json:"ConnectId,omitempty" name:"ConnectId"`
	// 连接名称

	ConnectName *string `json:"ConnectName,omitempty" name:"ConnectName"`
	// 开关边连接类型&nbsp;0：对等连接，&nbsp;1：云连网

	ConnectType *int64 `json:"ConnectType,omitempty" name:"ConnectType"`
	// 对等防火墙和开关状态&nbsp;0：正常，&nbsp;1：对等未创建防火墙，2：对等已创建防火墙，未打开开关

	CrossEdgeStatus *int64 `json:"CrossEdgeStatus,omitempty" name:"CrossEdgeStatus"`
	// 目的实例信息

	DstInstancesInfo []*NetInstancesInfo `json:"DstInstancesInfo,omitempty" name:"DstInstancesInfo"`
	// 开关状态&nbsp;0：关&nbsp;，&nbsp;1：开

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 防火墙(组)数据

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// 开关关联的防火墙实例列表

	FwInsLst []*VpcFwInstanceShow `json:"FwInsLst,omitempty" name:"FwInsLst"`
	// 网络经过VPC防火墙CVM所在地域

	FwInsRegion []*string `json:"FwInsRegion,omitempty" name:"FwInsRegion"`
	// 0:&nbsp;ipv4&nbsp;,&nbsp;1:ipv6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 0&nbsp;观察&nbsp;1&nbsp;拦截&nbsp;2&nbsp;严格&nbsp;3&nbsp;关闭

	IpsAction *int64 `json:"IpsAction,omitempty" name:"IpsAction"`
	// 源实例信息

	SrcInstancesInfo []*NetInstancesInfo `json:"SrcInstancesInfo,omitempty" name:"SrcInstancesInfo"`
	// 开关的状态&nbsp;0：正常，&nbsp;1：转换中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 防火墙开关ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 互通模式

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
	// 防火墙开关NAME

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
}

type CreatePcapTaskRequest struct {
	*tchttp.BaseRequest

	// 抓包时长，单位&nbsp;s

	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`
	// 实例id，nat则填nat实例id；vpc间防火墙则填防火墙组实例id

	FwInsIdLst []*string `json:"FwInsIdLst,omitempty" name:"FwInsIdLst"`
	// 抓包的防火墙类型

	FwType []*string `json:"FwType,omitempty" name:"FwType"`
	// 本端ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// ip配置类型，1&nbsp;单IP，2&nbsp;IP对

	IpConfig *int64 `json:"IpConfig,omitempty" name:"IpConfig"`
	// 最大字节数

	PkgSize *int64 `json:"PkgSize,omitempty" name:"PkgSize"`
	// 本端端口，-1表示所有

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议，值TCP，UDP，ICMP，ANY

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 对端ip

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 对端端口，-1表示所有

	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *CreatePcapTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePcapTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcConnectionsTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 峰值连接数

		ConnectionsMax *int64 `json:"ConnectionsMax,omitempty" name:"ConnectionsMax"`
		// 连接数配额占比

		ConnectionsPercent *float64 `json:"ConnectionsPercent,omitempty" name:"ConnectionsPercent"`
		// 连接数折线图

		ConnectionsTrendData []*FloatTrendsItem `json:"ConnectionsTrendData,omitempty" name:"ConnectionsTrendData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcConnectionsTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcConnectionsTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveNatAcRuleRequest struct {
	*tchttp.BaseRequest

	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则

	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *RemoveNatAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCfwInsBypassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCfwInsBypassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCfwInsBypassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonIdNameStatus struct {

	// 资源id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资源name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ZeroTrustOverview struct {

	// 已授权资产数量

	AuthCount *uint64 `json:"AuthCount,omitempty" name:"AuthCount"`
	// 数据库个数

	DbCount *uint64 `json:"DbCount,omitempty" name:"DbCount"`
	// 数据库登录次数

	DbLoginCount *uint64 `json:"DbLoginCount,omitempty" name:"DbLoginCount"`
	// 资产实例个数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例登录次数

	InstanceLoginCount *uint64 `json:"InstanceLoginCount,omitempty" name:"InstanceLoginCount"`
	// 剩余规格

	RemainAuthCount *uint64 `json:"RemainAuthCount,omitempty" name:"RemainAuthCount"`
	// web资产个数

	WebCount *uint64 `json:"WebCount,omitempty" name:"WebCount"`
	// web登录次数

	WebLoginCount *uint64 `json:"WebLoginCount,omitempty" name:"WebLoginCount"`
}

type DescribeVpcFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量控制列表数据

		Data []*TcRule `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatProtectSwitchDetail struct {

	// 开关值

	Idpsaction *int64 `json:"Idpsaction,omitempty" name:"Idpsaction"`
	// nat实例id

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// nat实例名称

	NatInsName *string `json:"NatInsName,omitempty" name:"NatInsName"`
	// 子网网段

	SubnetCidr *string `json:"SubnetCidr,omitempty" name:"SubnetCidr"`
	// 开关子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// vpc的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type WarnEventData struct {

	// 地理位置

	Address *string `json:"Address,omitempty" name:"Address"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产地域

	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 安全事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 是否生成规则，0：未生成，1：已生成

	GenRule *uint64 `json:"GenRule,omitempty" name:"GenRule"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// NAT实例id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// NAT实例名称

	NatInstanceName *string `json:"NatInstanceName,omitempty" name:"NatInstanceName"`
	// payload信息

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源/目的IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 访问源/目的端口

	PublicPort *string `json:"PublicPort,omitempty" name:"PublicPort"`
	// 判断来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源端口

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 安全事件/规则描述

	ThreatDesc *string `json:"ThreatDesc,omitempty" name:"ThreatDesc"`
	// 命中规则/威胁类型

	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeCfwInsCvmListRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCfwInsCvmListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsCvmListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Dnat规则列表

		Data []*DescNatDnatRule `json:"Data,omitempty" name:"Data"`
		// 列表总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebServicesRequest struct {
	*tchttp.BaseRequest

	// 排序字段，VulnCount按漏洞数量排序，VisitCountan访问次数排序，AttackCount按攻击次数排序，AuthUser按照授权用户的数量进行排序，LastVisitTime最近访问时间排序

	By *string `json:"By,omitempty" name:"By"`
	// 实例ID，基于实例ID进行检索

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 1表示只看接入的实例，2表示只看未接入实例，0表示查看全部实例

	IsAccessed *int64 `json:"IsAccessed,omitempty" name:"IsAccessed"`
	// 筛选授权资产

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 1表示只看新的实例，0全查看

	IsNew *string `json:"IsNew,omitempty" name:"IsNew"`
	// 长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 升序降序&nbsp;&nbsp;asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 已授权vpc筛选

	RemoteOMFilter *int64 `json:"RemoteOMFilter,omitempty" name:"RemoteOMFilter"`
	// 查询条件
	// "SearchValue":"{\"CloudCode\":\"1,2\",\"common\":\"\"}"
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;查询排序的字段
	// 	"SrcIp":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//源&nbsp;&nbsp;多条件&nbsp;1,2,3&nbsp;格式
	// 	"Uuid":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//规则id
	// 	"DstContent"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//目的
	// 	"CloudCode":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//云厂商
	// 	"UpdateTime":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//时间
	// 	"Disable":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//状态
	// 	"Resource":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//资产组
	// 	"InsId":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//资产
	// 	"DetectedTimes"&nbsp;//次数&nbsp;
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"ProtectType"//&nbsp;防护方式（-1:未防护，0:基础防护，1:高级防护）&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;多条件&nbsp;&nbsp;&nbsp;-1,0,1格式
	//
	// common&nbsp;不指定字段&nbsp;和下面FuzzySearch&nbsp;&nbsp;功能&nbsp;一直&nbsp;&nbsp;前端那个好处理使用那个

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 是否展示运维用户：1:展示，其他不展示

	ShowYwUser *string `json:"ShowYwUser,omitempty" name:"ShowYwUser"`
	// 查询时间段0代表24小时，1代表7天

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
	// 运维id列表，多个id用英文逗号隔开，当ShowYwUser为1时，该参数可过滤出某一个运维用户的服务

	YwId *string `json:"YwId,omitempty" name:"YwId"`
}

func (r *DescribeWebServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SGInstanceInfo struct {

	// 子网cidr

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例对应的安全组

	SgSet *string `json:"SgSet,omitempty" name:"SgSet"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 资产类型
	// &nbsp;3是cvm实例,4是clb实例,5是eni实例,6是mysql,7是redis,8是NAT,9是VPN,10是ES,11是MARIADB,12是KAFKA

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// vpcid信息

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 弹性带宽上限，单位Mbps

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 弹性开关&nbsp;1打开&nbsp;0&nbsp;关闭

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// nat：nat防火墙，ew：东西向防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 按量计费标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 带宽值

	Width *uint64 `json:"Width,omitempty" name:"Width"`
}

func (r *ExpandCfwVerticalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandCfwVerticalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertCenterRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码：
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息：
		// success&nbsp;成功
		// 其他

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 处置状态码：
		// 0&nbsp;&nbsp;处置成功
		// -1&nbsp;通用错误，不用处理
		// -3&nbsp;表示重复，需重新刷新列表
		// 其他

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertCenterRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGuideIdpsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGuideIdpsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGuideIdpsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapDownUrlRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribePcapDownUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePcapDownUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpsRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0&nbsp;表示都操作成功，1表示部分引擎不支持，需要升级

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpsRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpsRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwInsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙实例运行状态

		CfwInsStatus []*CfwInsStatus `json:"CfwInsStatus,omitempty" name:"CfwInsStatus"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwInsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBetaTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBetaTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebAssetFilterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*WebAssetsInfo `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebAssetFilterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetFilterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceGroupChangeImpactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源组url

		ResourceGroupPath *string `json:"ResourceGroupPath,omitempty" name:"ResourceGroupPath"`
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 影响规则数量

		RuleCount *RuleCount `json:"RuleCount,omitempty" name:"RuleCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceGroupChangeImpactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceGroupChangeImpactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMessageCenterSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMessageCenterSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMessageCenterSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZeroTrustAccess struct {

	// 接入类型，微信wechat，IOA&nbsp;ioa

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// 目录数

	DirectoryTotal *string `json:"DirectoryTotal,omitempty" name:"DirectoryTotal"`
	// 公司名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 接入状态，1已接入，0未接入

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 接入总数

	UserTotal *int64 `json:"UserTotal,omitempty" name:"UserTotal"`
	// 7天访问数

	VisitsTimes *int64 `json:"VisitsTimes,omitempty" name:"VisitsTimes"`
}

type DeleteIdsWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码：
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息：
		// success&nbsp;成功
		// 其他

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回状态码：
		// 0&nbsp;&nbsp;处置成功
		// -1&nbsp;通用错误，不用处理

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIdsWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIdsWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递地址配置

		AddrConfigList []*SysLogAddrConfig `json:"AddrConfigList,omitempty" name:"AddrConfigList"`
		// 投递类型配置

		DeliveryConfigList []*SysLogTypeConfig `json:"DeliveryConfigList,omitempty" name:"DeliveryConfigList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDeliveryConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfoResource struct {

	// '策略，1阻断，2放行,ACCEPT&nbsp;或&nbsp;DROP

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// cidr信息

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// RegCode2码

	CityRegCode *string `json:"CityRegCode,omitempty" name:"CityRegCode"`
	// RegName2码

	CityRegName *string `json:"CityRegName,omitempty" name:"CityRegName"`
	// tencentcloudenterprise_code

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// RegCode1码

	CountryRegCode *string `json:"CountryRegCode,omitempty" name:"CountryRegCode"`
	// RegName1码

	CountryRegName *string `json:"CountryRegName,omitempty" name:"CountryRegName"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向&nbsp;0出&nbsp;1入

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 规则是否开启，0未开启，1开启

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 目的ip

	DstContent *string `json:"DstContent,omitempty" name:"DstContent"`
	// 目的路径

	DstPath *string `json:"DstPath,omitempty" name:"DstPath"`
	// 目的类型

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 记录id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 资产名称

	InsName *string `json:"InsName,omitempty" name:"InsName"`
	// 0为正常规则,1为云厂商规则

	IsCloud *int64 `json:"IsCloud,omitempty" name:"IsCloud"`
	// 是否最新一次改动的规则,0否，1是

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 0为正常规则,1为地域规则

	IsReg *int64 `json:"IsReg,omitempty" name:"IsReg"`
	// 是否双向下发，0否，1是

	IsReverse *int64 `json:"IsReverse,omitempty" name:"IsReverse"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 协议端口参数模板

	ProtocolType *int64 `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 规则来源，0为控制台

	RuleSource *int64 `json:"RuleSource,omitempty" name:"RuleSource"`
	// 序号

	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`
	// 协议端口参数模板id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 端口模板名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 访问源Cidr

	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`
	// 访问源资产名

	SourceInsName *string `json:"SourceInsName,omitempty" name:"SourceInsName"`
	// 访问源内网ip

	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" name:"SourcePrivateIp"`
	// 访问源公网ip

	SourcePublicIp *string `json:"SourcePublicIp,omitempty" name:"SourcePublicIp"`
	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 源路径

	SrcPath *string `json:"SrcPath,omitempty" name:"SrcPath"`
	// 源类型

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 所在的子网id'

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 地址模板addr

	TmplAddr *string `json:"TmplAddr,omitempty" name:"TmplAddr"`
	// 地址模板描述

	TmplDetail *string `json:"TmplDetail,omitempty" name:"TmplDetail"`
	// 模板名称&nbsp;地址模板和安全组的参数模板

	TmplName *string `json:"TmplName,omitempty" name:"TmplName"`
	// 规则uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 所在的vpc&nbsp;id'

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeLogDeliverCongResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*StringValueText `json:"Data,omitempty" name:"Data"`
		// 返回状态&nbsp;0成功&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回消息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDeliverCongResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliverCongResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 阻断规则条数

		BlockNum *int64 `json:"BlockNum,omitempty" name:"BlockNum"`
		// 端口转发规则数量

		DnatNum *int64 `json:"DnatNum,omitempty" name:"DnatNum"`
		// 弹性IP列表

		EipList []*string `json:"EipList,omitempty" name:"EipList"`
		// 启用规则条数

		EnableNum *int64 `json:"EnableNum,omitempty" name:"EnableNum"`
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 访问控制规则剩余配额

		RemainNum *int64 `json:"RemainNum,omitempty" name:"RemainNum"`
		// 访问控制规则总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙实例卡片信息数组

		FwInstanceLst []*VpcFwInstanceName `json:"FwInstanceLst,omitempty" name:"FwInstanceLst"`
		// 接入网络节点地域列表

		RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUpdateResourceTaskStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryUpdateResourceTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUpdateResourceTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSysLogTaskConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSysLogTaskConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSysLogTaskConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwExportBindRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DescribeCfwExportBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwExportBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRunSyncAssetRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;互联网防火墙开关，1：vpc&nbsp;防火墙开关

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyRunSyncAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRunSyncAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BanAndAllowRule struct {

	// 规则评论

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 自定义白名单规则

	CustomRule *CustomWhiteRule `json:"CustomRule,omitempty" name:"CustomRule"`
	// 0互联网出站&nbsp;1互联网入站&nbsp;5内网访问源&nbsp;6内网访问目的

	DirectionList *string `json:"DirectionList,omitempty" name:"DirectionList"`
	// 规则截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 放通的引擎:&nbsp;1针对互联网边界&nbsp;2针对nat防火墙&nbsp;4针对vpc防火墙

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// 封禁和放通对象

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
}

type NewsDetail struct {

	// 影响面

	Affect *string `json:"Affect,omitempty" name:"Affect"`
	// 攻击类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 安全事件描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 自增id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 保留字段

	Message *string `json:"Message,omitempty" name:"Message"`
	// 消息类型&nbsp;1(漏洞信息),2(产品功能动态)

	News *int64 `json:"News,omitempty" name:"News"`
	// 发布时间

	ReleaseDate *string `json:"ReleaseDate,omitempty" name:"ReleaseDate"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 文章标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// 版本信息

	VersionEn *string `json:"VersionEn,omitempty" name:"VersionEn"`
	// 影响面

	AffectEn *string `json:"AffectEn,omitempty" name:"AffectEn"`
	// 攻击类型

	CategoryEn *string `json:"CategoryEn,omitempty" name:"CategoryEn"`
	// 文章标题

	TitleEn *string `json:"TitleEn,omitempty" name:"TitleEn"`
	// 安全事件描述

	DescriptEn *string `json:"DescriptEn,omitempty" name:"DescriptEn"`
	// 危险等级

	LevelEn *string `json:"LevelEn,omitempty" name:"LevelEn"`
}

type SwitchFailInfo struct {

	// 自增唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 开关名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 防火墙开关变动状态，小于0

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest

	// NAT地域，&nbsp;如ap-NewYork

	Area *string `json:"Area,omitempty" name:"Area"`
	// 出站还是入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckNatFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// IP地址或者IP掩码，比如192.168.0.1或者192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// NAT防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
}

func (r *CheckNatFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNatFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsResolveIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ip列表

		IpList []*string `json:"IpList,omitempty" name:"IpList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDnsResolveIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDnsResolveIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreRuleRequest struct {
	*tchttp.BaseRequest

	// 规则

	Rule *IntrusionDefenseRule `json:"Rule,omitempty" name:"Rule"`
	// 规则类型，1封禁，2放通

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *ModifyBlockIgnoreRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwUserInfo struct {

	// 资产更新时间

	AssetUpdateTime *string `json:"AssetUpdateTime,omitempty" name:"AssetUpdateTime"`
	// 拦截攻击个数

	AttackNum *int64 `json:"AttackNum,omitempty" name:"AttackNum"`
	// 处理告警

	HandleCount *int64 `json:"HandleCount,omitempty" name:"HandleCount"`
	// 规则库更新时间

	IPSRuleUpdateTime *string `json:"IPSRuleUpdateTime,omitempty" name:"IPSRuleUpdateTime"`
	// 0观察，1拦截，2严格

	IdpsAction *int64 `json:"IdpsAction,omitempty" name:"IdpsAction"`
	// 互联网拦截模式数量

	InternetBlockActionCount *int64 `json:"InternetBlockActionCount,omitempty" name:"InternetBlockActionCount"`
	// 互联网观察模式数量

	InternetMonitorActionCount *int64 `json:"InternetMonitorActionCount,omitempty" name:"InternetMonitorActionCount"`
	// 互联网严格模式数量

	InternetStrictActionCount *int64 `json:"InternetStrictActionCount,omitempty" name:"InternetStrictActionCount"`
	// nat拦截模式数量

	NatBlockActionCount *int64 `json:"NatBlockActionCount,omitempty" name:"NatBlockActionCount"`
	// nat关闭数量

	NatCloseActionCount *int64 `json:"NatCloseActionCount,omitempty" name:"NatCloseActionCount"`
	// nat防火墙数量

	NatFWCount *int64 `json:"NatFWCount,omitempty" name:"NatFWCount"`
	// nat观察模式数量

	NatMonitorActionCount *int64 `json:"NatMonitorActionCount,omitempty" name:"NatMonitorActionCount"`
	// nat严格模式数量

	NatStrictActionCount *int64 `json:"NatStrictActionCount,omitempty" name:"NatStrictActionCount"`
	// 防护ip数量

	ProtectedIP *int64 `json:"ProtectedIP,omitempty" name:"ProtectedIP"`
	// ip总数

	TotalIP *int64 `json:"TotalIP,omitempty" name:"TotalIP"`
	// 虚拟补丁是否开始，1开启&nbsp;0关闭

	VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitempty" name:"VirtualPatchSwitch"`
	// vpc拦截模式数量

	VpcBlockActionCount *int64 `json:"VpcBlockActionCount,omitempty" name:"VpcBlockActionCount"`
	// vpc关闭数量

	VpcCloseActionCount *int64 `json:"VpcCloseActionCount,omitempty" name:"VpcCloseActionCount"`
	// vpc观察模式数量

	VpcMonitorActionCount *int64 `json:"VpcMonitorActionCount,omitempty" name:"VpcMonitorActionCount"`
	// vpc严格模式数量

	VpcStrictActionCount *int64 `json:"VpcStrictActionCount,omitempty" name:"VpcStrictActionCount"`
}

type UserEngineZoneInfo struct {

	// 引擎实例类型

	FwEnginType *string `json:"FwEnginType,omitempty" name:"FwEnginType"`
	// 引擎实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否在单区域

	IsSingleZone *bool `json:"IsSingleZone,omitempty" name:"IsSingleZone"`
	// 地域列表

	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
	// 可用区列表

	ZoneList []*string `json:"ZoneList,omitempty" name:"ZoneList"`
	// 可用区名称列表

	ZoneNameList []*string `json:"ZoneNameList,omitempty" name:"ZoneNameList"`
}

type DescribeWarnAssetEventListRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，desc：降序，asc：升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 判断来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 事件处置状态过滤，0：显示全部，1：只显示未处置

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeWarnAssetEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarnAssetEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveNatFwObjRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 模式&nbsp;1：接入模式；0：新增模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 接入模式移除接入的nat网关列表，其中NatGwList和VpcList只能传递一个。

	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`
	// 新增模式移除接入的vpc列表，其中NatGwList和NatgwList只能传递一个。

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *RemoveNatFwObjRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatFwObjRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBlockIgnoreRulesImportRequest struct {
	*tchttp.BaseRequest
}

func (r *StopBlockIgnoreRulesImportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBlockIgnoreRulesImportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwEdgeIpsRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 筛选公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFwEdgeIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEdgeIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功
		// 其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回状态
		// 0&nbsp;请求成功
		// 2&nbsp;请求失败
		// 3&nbsp;请求失败-频率限制

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetSyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteMtInstance struct {

	// 0&nbsp;手动接入&nbsp;1&nbsp;自动管理接入

	AccessMode *int64 `json:"AccessMode,omitempty" name:"AccessMode"`
	// cvm数量

	CvmCount *uint64 `json:"CvmCount,omitempty" name:"CvmCount"`
	// db数量

	DBCount *uint64 `json:"DBCount,omitempty" name:"DBCount"`
	// EIP数组

	Eips []*string `json:"Eips,omitempty" name:"Eips"`
	// 开关状态0，1

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 是否开启高可用实例&nbsp;0&nbsp;关闭&nbsp;1&nbsp;打开

	HighAvailabilitySwitch *int64 `json:"HighAvailabilitySwitch,omitempty" name:"HighAvailabilitySwitch"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否为低版本免费版
	// 0&nbsp;不可以创建EIP
	// 1&nbsp;可以创建EIP&nbsp;但只可以创建1个
	// 2&nbsp;高级版&nbsp;可以创建两个EIP

	OldUser *int64 `json:"OldUser,omitempty" name:"OldUser"`
	// 地域标识

	RegionArea *string `json:"RegionArea,omitempty" name:"RegionArea"`
	// 0:完成,&nbsp;1:开启中，2：关闭中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// vpc数量

	VpcCount *uint64 `json:"VpcCount,omitempty" name:"VpcCount"`
	// web数量

	WebCount *uint64 `json:"WebCount,omitempty" name:"WebCount"`
}

type DescribeCfwEipsRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id，当前仅支持接入模式的实例

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 1：cfw接入模式，目前仅支持接入模式实例

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// ALL：查询所有弹性公网ip;&nbsp;nat-xxxxx：接入模式场景指定网关的弹性公网ip

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
}

func (r *DescribeCfwEipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwEipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwInstancesInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙实例卡片信息数组

		Data []*VpcFwInstanceInfo `json:"Data,omitempty" name:"Data"`
		// VPC防火墙个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwInstancesInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwInstancesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetScanListRequest struct {
	*tchttp.BaseRequest

	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 资产类型

	InsType *string `json:"InsType,omitempty" name:"InsType"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 资产组id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeAssetScanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociatedInstanceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		Data []*AssociatedInstanceInfo `json:"Data,omitempty" name:"Data"`
		// 实例数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssociatedInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssociatedInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIsolateNetCardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果合计

		Data []*IsolateNetCard `json:"Data,omitempty" name:"Data"`
		// 0成功&nbsp;非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIsolateNetCardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIsolateNetCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewFlowStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽统计数据列表

		Data []*OverViewEdgeFlowStat `json:"Data,omitempty" name:"Data"`
		// 入均值

		InAvg *int64 `json:"InAvg,omitempty" name:"InAvg"`
		// 入峰值

		InMax *int64 `json:"InMax,omitempty" name:"InMax"`
		// 出均值

		OutAvg *int64 `json:"OutAvg,omitempty" name:"OutAvg"`
		// 出峰值

		OutMax *int64 `json:"OutMax,omitempty" name:"OutMax"`
		// 出入向均值

		TotalAvg *int64 `json:"TotalAvg,omitempty" name:"TotalAvg"`
		// 出入向流量值

		TotalFlow *int64 `json:"TotalFlow,omitempty" name:"TotalFlow"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewFlowStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewFlowStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		Data []*BlockIgnoreRule `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 对应规则类型的数量，示例：[0,122,30,55,12,232,0]，封禁0个，IP地址122个，域名30个，威胁情报55个，资产实例12个，自定义策略232个，入侵防御规则0个

		RuleTypeDataList []*int64 `json:"RuleTypeDataList,omitempty" name:"RuleTypeDataList"`
		// 安全事件来源下拉框&nbsp;["基础防御","手动添加","网络蜜罐（严格模式）","虚拟补丁"]

		SourceList []*string `json:"SourceList,omitempty" name:"SourceList"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIgnoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEdgeIpSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwAZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFwAZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwAZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearResourceGroupRequest struct {
	*tchttp.BaseRequest
}

func (r *ClearResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetflowBorderUsedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1已超量，非1未超量

		IsOverUsed *int64 `json:"IsOverUsed,omitempty" name:"IsOverUsed"`
		// 超量策略生效时间

		PolicyChangeTime *string `json:"PolicyChangeTime,omitempty" name:"PolicyChangeTime"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 开始超量时间

		StartOverTime *string `json:"StartOverTime,omitempty" name:"StartOverTime"`
		// 使用配额比例

		UsedPercentage *float64 `json:"UsedPercentage,omitempty" name:"UsedPercentage"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetflowBorderUsedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetflowBorderUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetflowCenterTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*TrendsItem `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetflowCenterTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetflowCenterTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志存储配置详情

		Data *StorageSetting `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口风险列表

		Data []*PortRiskData `json:"Data,omitempty" name:"Data"`
		// 入站规则数

		InNum *uint64 `json:"InNum,omitempty" name:"InNum"`
		// 上次扫描成功时间

		LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
		// 公网ip数组

		PublicIpList []*string `json:"PublicIpList,omitempty" name:"PublicIpList"`
		// 是否扫描成功过，0&nbsp;未扫描，1：已扫描

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreAsyncTaskErrRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// acl_border_import&nbsp;&nbsp;acl_secgroupnew_import&nbsp;acl_vpc_import&nbsp;acl_nat_import&nbsp;acl_dns_import&nbsp;

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *ModifyIgnoreAsyncTaskErrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIgnoreAsyncTaskErrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupTableStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupTableStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcFwJoinInstanceType struct {

	// JoinType

	JoinType *string `json:"JoinType,omitempty" name:"JoinType"`
	// 接入的对应网络实例类型的数量

	Num *int64 `json:"Num,omitempty" name:"Num"`
}

type CreateAcRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回多余的信息

		Info *string `json:"Info,omitempty" name:"Info"`
		// 状态值&nbsp;0正常&nbsp;其他异常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAcRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAcRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainInfoRequest struct {
	*tchttp.BaseRequest

	// 实例Id，不指定则传空

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，风险评估：risk_assess

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 排序参数，访问流量从小到大flowCount:asc，访问流量从大到小flowCount:desc，最近访问时间从小到大end_time:asc，最近访问时间从大到小end_time:desc，访问次数从小到大total_count:asc，访问次数从大到小total_count:desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 时间类型，0：7天，1：24小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNetFlowDomainInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeHoneyPotSelector struct {

	// 1&nbsp;允许修改
	// 0&nbsp;不允许修改

	AllowModify *int64 `json:"AllowModify,omitempty" name:"AllowModify"`
	// web&nbsp;服务默认使用路径

	DefaultPath []*string `json:"DefaultPath,omitempty" name:"DefaultPath"`
	// 蜜罐服务使用的默认端口

	DefaultPort *int64 `json:"DefaultPort,omitempty" name:"DefaultPort"`
	// 蜜罐的交互类型

	InteractionType *string `json:"InteractionType,omitempty" name:"InteractionType"`
	// 蜜罐的服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 蜜罐类型&nbsp;服务端口&nbsp;蜜罐名称

	ServiceNameWithPort *string `json:"ServiceNameWithPort,omitempty" name:"ServiceNameWithPort"`
	// 支持部署模式，1:公网IP探针，2:负载均衡探针，3:内网IP探针

	SupportDeployMode *string `json:"SupportDeployMode,omitempty" name:"SupportDeployMode"`
	// 是否是web类型蜜罐，0：非web类型;&nbsp;1：web类型,&nbsp;支持选择lb探针

	WebType *int64 `json:"WebType,omitempty" name:"WebType"`
}

type CreateBakRuleListRequest struct {
	*tchttp.BaseRequest

	// 备份规则描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 备份规则为NAT边界规则时的地域

	NATRegion *string `json:"NATRegion,omitempty" name:"NATRegion"`
	// 备份列表名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateBakRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBakRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportLogsRequest struct {
	*tchttp.BaseRequest

	// 筛选结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 日志类型标识

	Index *string `json:"Index,omitempty" name:"Index"`
	// 筛选开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeExportLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWarnEventRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;下发成功，-1:&nbsp;通用错误，-3:&nbsp;规则重复，-99:&nbsp;下发失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWarnEventRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWarnEventRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyLogApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfoData struct {

	// 城市Code

	City *uint64 `json:"City,omitempty" name:"City"`
	// 城市

	CityName *string `json:"CityName,omitempty" name:"CityName"`
	// 云厂商地域

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 国家Code

	Country *uint64 `json:"Country,omitempty" name:"Country"`
	// 国家

	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否地域&nbsp;1是0不是

	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// log

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 国家二位iso代码或者省份缩写代码

	RegionIso *string `json:"RegionIso,omitempty" name:"RegionIso"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
}

type TemplateListInfo struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 模板包含地址数量

	IPNum *int64 `json:"IPNum,omitempty" name:"IPNum"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// IP模版

	IpString *string `json:"IpString,omitempty" name:"IpString"`
	// IP版本,0,IPv4;1,IPv6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 协议端口模板，协议类型，4:4层协议，7:7层协议

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 关联规则条数

	RulesNum *int64 `json:"RulesNum,omitempty" name:"RulesNum"`
	// 模板Id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模版类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 模版ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AddAcRuleRequest struct {
	*tchttp.BaseRequest

	// 七层协议，取值：
	// HTTP/HTTPS
	// TLS/SSL

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// 访问控制策略的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 访问控制策略中的目的地址。取值：
	// 当DestType为net时，DestContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	//
	// 当DestType为template时，DestContent为源地址模板id。
	//
	// 当DestType为location时，DestContent为源区域。
	// 例如["BJ11",&nbsp;"ZB"]
	//
	// 当DestType为instance时，DestContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	//
	// 当DestType为domain时，DestContent为该实例id对应的域名规则。
	// 例如*.domain.com
	//
	// 当DestType为vendor时，DestContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tx,aliyun,azure,all代表以上五个

	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`
	// 访问控制策略中的目的地址类型。取值：
	// net：目的IP或者网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商
	// domain:&nbsp;域名或者ip

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 访问控制策略的流量方向。取值：
	// in：外对内流量访问控制
	// out：内对外流量访问控制

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// -1表示优先级最低，1表示优先级最高

	OrderIndex *string `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80,443：80或者443

	Port *string `json:"Port,omitempty" name:"Port"`
	// 访问控制策略中流量访问的协议类型。取值：TCP，目前互联网边界规则只能支持TCP，不传参数默认就是TCP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 访问控制策略中的源地址。取值：
	// 当SourceType为net时，SourceContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	//
	// 当SourceType为template时，SourceContent为源地址模板id。
	//
	// 当SourceType为location时，SourceContent为源区域。
	// 例如["BJ11",&nbsp;"ZB"]
	//
	// 当SourceType为instance时，SourceContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	//
	// 当SourceType为vendor时，SourceContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tx,aliyun,azure,all代表以上五个

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问控制策略中的源地址类型。取值：
	// net：源IP或网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *AddAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGlobalSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopBlockIgnoreRulesImportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopBlockIgnoreRulesImportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopBlockIgnoreRulesImportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的资源分组id

		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
		// 状态值，0：编辑成功，非0：编辑失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeWarnListsRequest struct {
	*tchttp.BaseRequest

	// 方向，0：入站，1：出站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 危险等级,&nbsp;提示->信息，中危->低，高危->高，极高危->极高

	Level *string `json:"Level,omitempty" name:"Level"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 判断来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 0:&nbsp;观察，1：阻断，2：放行

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 类型，0：7天，1：24小时，2：近30天，3：6个月

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSafeWarnListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeWarnListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllProbePort struct {

	// 探针部署模式1、EIP，&nbsp;2、负载均衡

	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`
	// 探针ID

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 探针使用的路径

	ProbePath *string `json:"ProbePath,omitempty" name:"ProbePath"`
	// 探针使用的端口

	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
}

type DescribeAlertLogEventNameSelectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 返回值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回值

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertLogEventNameSelectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertLogEventNameSelectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartUpdateResourceTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartUpdateResourceTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartUpdateResourceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefineBait struct {

	// 蜜罐关联其他蜜罐(天幕蜜罐剧本)

	RelateHoneyPot []*string `json:"RelateHoneyPot,omitempty" name:"RelateHoneyPot"`
	// 蜜罐root密码，支持多用户密码

	SshPasswd *string `json:"SshPasswd,omitempty" name:"SshPasswd"`
	// 蜜罐&nbsp;web页面&nbsp;自定义诱饵数据

	StaticBait *string `json:"StaticBait,omitempty" name:"StaticBait"`
}

type RegionStatusData struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态&nbsp;0正常&nbsp;其他异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeImportCredentialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传桶

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// cos域名

		CosDomain *string `json:"CosDomain,omitempty" name:"CosDomain"`
		// 凭证过期时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 桶所在地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 凭证有效开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 临时SecretId

		TemplateSecretId *string `json:"TemplateSecretId,omitempty" name:"TemplateSecretId"`
		// 临时SecretKey

		TemplateSecretKey *string `json:"TemplateSecretKey,omitempty" name:"TemplateSecretKey"`
		// Token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 上传链接前缀

		UrlPrefix *string `json:"UrlPrefix,omitempty" name:"UrlPrefix"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEdgeIpSwitchAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAclRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功后返回新策略ID列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAclRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCheckCLSStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启&nbsp;1开启0关闭

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCheckCLSStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckCLSStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupChangeImpactNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数据

		RuleImpactList *RuleImpactNew `json:"RuleImpactList,omitempty" name:"RuleImpactList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupChangeImpactNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupChangeImpactNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetOverviewEventNew struct {

	// 待处理告警总数

	AlertCount *int64 `json:"AlertCount,omitempty" name:"AlertCount"`
	// 服务器资产总数

	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`
	// clb总数

	ClbCount *int64 `json:"ClbCount,omitempty" name:"ClbCount"`
	// 暴力破解次数

	CrackCount *int64 `json:"CrackCount,omitempty" name:"CrackCount"`
	// 新增数据库

	DbAddCount *int64 `json:"DbAddCount,omitempty" name:"DbAddCount"`
	// 数据库总数

	DbCount *int64 `json:"DbCount,omitempty" name:"DbCount"`
	// 高危事件总数

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 新增公网ip

	IpAddCount *int64 `json:"IpAddCount,omitempty" name:"IpAddCount"`
	// 公网ip总数

	IpCount *int64 `json:"IpCount,omitempty" name:"IpCount"`
	// 新增漏洞

	LeakAddCount *int64 `json:"LeakAddCount,omitempty" name:"LeakAddCount"`
	// 暴露漏洞总数

	LeakCount *int64 `json:"LeakCount,omitempty" name:"LeakCount"`
	// nat总数

	NatCount *int64 `json:"NatCount,omitempty" name:"NatCount"`
	// 网络资产总数

	NetCount *int64 `json:"NetCount,omitempty" name:"NetCount"`
	// 恶意主动外联总数

	OutRelateCount *int64 `json:"OutRelateCount,omitempty" name:"OutRelateCount"`
	// 已失陷主机总数

	OutTiCount *int64 `json:"OutTiCount,omitempty" name:"OutTiCount"`
	// 新增暴露端口

	PortAddCount *int64 `json:"PortAddCount,omitempty" name:"PortAddCount"`
	// 暴露端口总数

	PortCount *int64 `json:"PortCount,omitempty" name:"PortCount"`
	// 新增ssh服务

	SshAddCount *int64 `json:"SshAddCount,omitempty" name:"SshAddCount"`
	// ssh服务总数

	SshCount *int64 `json:"SshCount,omitempty" name:"SshCount"`
	// 漏洞攻击次数

	VulAttachCount *int64 `json:"VulAttachCount,omitempty" name:"VulAttachCount"`
	// 新增web服务

	WebAddCount *int64 `json:"WebAddCount,omitempty" name:"WebAddCount"`
	// web服务总数

	WebCount *int64 `json:"WebCount,omitempty" name:"WebCount"`
	// 新增rdp服务

	WindowsAddCount *int64 `json:"WindowsAddCount,omitempty" name:"WindowsAddCount"`
	// rdp服务总数

	WindowsCount *int64 `json:"WindowsCount,omitempty" name:"WindowsCount"`
}

type InstanceInfo struct {

	// appid信息

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 1，公网&nbsp;2内网

	InsSource *string `json:"InsSource,omitempty" name:"InsSource"`
	// 资产类型
	// &nbsp;3是cvm实例,4是clb实例,5是eni实例,6是mysql,7是redis,8是NAT,9是VPN,10是ES,11是MARIADB,12是KAFKA&nbsp;13&nbsp;NATFW

	InsType *int64 `json:"InsType,omitempty" name:"InsType"`
	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 漏洞数

	LeakNum *string `json:"LeakNum,omitempty" name:"LeakNum"`
	// 端口数

	PortNum *string `json:"PortNum,omitempty" name:"PortNum"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域

	RegionKey *string `json:"RegionKey,omitempty" name:"RegionKey"`
	// [a,b]

	ResourcePath []*string `json:"ResourcePath,omitempty" name:"ResourcePath"`
	// 扫描结果

	Server []*string `json:"Server,omitempty" name:"Server"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// vpc名称

	VPCName *string `json:"VPCName,omitempty" name:"VPCName"`
	// vpcid信息

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type VpcFwCvmInsInfo struct {

	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 防火墙CVM带宽值

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// VPC防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// CVM所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// CVM所在地域详情

	RegionDetail *string `json:"RegionDetail,omitempty" name:"RegionDetail"`
	// CVM所在地域中文

	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`
	// 实例主机所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例备机所在可用区

	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`
	// 主机所在可用区

	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`
	// 备机所在可用区

	ZoneZhBack *string `json:"ZoneZhBack,omitempty" name:"ZoneZhBack"`
	// 实例主机id

	CvmIdFirst *string `json:"CvmIdFirst,omitempty" name:"CvmIdFirst"`
	// 实例备机id

	CvmIdSecond *string `json:"CvmIdSecond,omitempty" name:"CvmIdSecond"`
	// 实例主机ip

	CvmIpFirst *string `json:"CvmIpFirst,omitempty" name:"CvmIpFirst"`
	// 实例备机ip

	CvmIpSecond *string `json:"CvmIpSecond,omitempty" name:"CvmIpSecond"`
}

type DescribeAssetHubDbAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*DescribeAssetHubDbAssetListInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHubDbAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHubDbAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaradStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*BaradStatus `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaradStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaradStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。

	CfwInsIdList []*string `json:"CfwInsIdList,omitempty" name:"CfwInsIdList"`
	// 开关，0：关闭，1：开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。

	RouteTableIdList []*string `json:"RouteTableIdList,omitempty" name:"RouteTableIdList"`
	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。

	SubnetIdList []*string `json:"SubnetIdList,omitempty" name:"SubnetIdList"`
}

func (r *ModifyNatFwSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeTaskHistory struct {

	// 序号

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 拨测次数

	ProbeCount *uint64 `json:"ProbeCount,omitempty" name:"ProbeCount"`
	// 拨测结果，1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	Result *uint64 `json:"Result,omitempty" name:"Result"`
	// 拨测结果说明

	Status *string `json:"Status,omitempty" name:"Status"`
	// 拨测时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type TagInfo struct {

	// 目标key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 目标值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeBaradStatusRequest struct {
	*tchttp.BaseRequest

	// 页高

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBaradStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaradStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEdgeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 边的目的地域集合

		DstRegion []*string `json:"DstRegion,omitempty" name:"DstRegion"`
		// vpc防火墙边状态

		EdgeInstance []*EdgeStatus `json:"EdgeInstance,omitempty" name:"EdgeInstance"`
		// 失败状态的边

		FailData []*EdgeStatus `json:"FailData,omitempty" name:"FailData"`
		// 关闭状态的边个数

		OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`
		// 开启状态的边个数

		OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`
		// 边的源地域集合

		SrcRegion []*string `json:"SrcRegion,omitempty" name:"SrcRegion"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEdgeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEdgeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRuleLimitsForWidthRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
}

func (r *QueryRuleLimitsForWidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRuleLimitsForWidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWarnEventData struct {

	// 告警次数

	AlertCount *uint64 `json:"AlertCount,omitempty" name:"AlertCount"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产ip

	AssetIp *string `json:"AssetIp,omitempty" name:"AssetIp"`
	// 拦截次数

	BlockCount *uint64 `json:"BlockCount,omitempty" name:"BlockCount"`
	// 待处置数

	LeftNum *uint64 `json:"LeftNum,omitempty" name:"LeftNum"`
	// 安全事件类型list

	List []*AssetDayEventInfo `json:"List,omitempty" name:"List"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 安全事件数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type ListenerInfo struct {

	// 监听器包含的域名列表

	Domain []*string `json:"Domain,omitempty" name:"Domain"`
	// 监听器id

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 路径绑定cvm的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeGuidePortProtectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 禁封端口页查询信息

		Data *PortProtect `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuidePortProtectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuidePortProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwGroupInsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙(组)列表信息

		FwGroupLst []*VpcFwGroupInsShow `json:"FwGroupLst,omitempty" name:"FwGroupLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwGroupInsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupInsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreImportCredentialRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlockIgnoreImportCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreImportCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTreatInfoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;开启或关闭成功，1：失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTreatInfoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTreatInfoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackListQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 需要配额

		NeedQuota *uint64 `json:"NeedQuota,omitempty" name:"NeedQuota"`
		// 剩余配额

		RemainQuota *uint64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 0成功&nbsp;非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总配额

		TotalQuota *uint64 `json:"TotalQuota,omitempty" name:"TotalQuota"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlackListQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlackListQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningCenterOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警中心概览数据

		Data *WarningCenterOverviewData `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarningCenterOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningCenterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRuleItem struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 内部id

	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`
	// 规则序号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口协议组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 访问控制策略的端口。取值：&nbsp;-1/-1：全部端口&nbsp;80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：&nbsp;TCP&nbsp;UDP&nbsp;ICMP&nbsp;ANY&nbsp;HTTP&nbsp;HTTPS&nbsp;HTTP/HTTPS&nbsp;SMTP&nbsp;SMTPS&nbsp;SMTP/SMTPS&nbsp;FTP&nbsp;DNS

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：&nbsp;accept：放行&nbsp;drop：拒绝&nbsp;log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 规则生效的范围：ALL，全局生效；ap-NewYork，生效的地域；cfwnat-xxx，生效基于实例维度

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 访问源示例：&nbsp;net：IP/CIDR(192.168.0.2)

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源类型：入向规则时类型可以为&nbsp;ip,net,template,location；出向规则时可以为&nbsp;ip,net,template,instance,group,tag

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 访问目的示例：&nbsp;net：IP/CIDR(192.168.0.2)&nbsp;domain：域名规则，例如*.domain.com

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为&nbsp;&nbsp;ip,net,domain,template,location

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则对应的唯一id，创建规则时无需填写

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type TiWhiteInfo struct {

	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 白名单对象

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeRuleOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// 剩余配额

		RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`
		// 启用规则数量

		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`
		// 停用规则数量

		StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`
		// 阻断策略规则数量

		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogsOfflineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLogsOfflineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsOfflineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 当前规则ID

	RuleUuid []*uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
	// 是否开关开启，0：未开启，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAclRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIgnoreListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindEdgeCFWResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindEdgeCFWResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindEdgeCFWResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSProtectDetail struct {

	// 入侵防御模式

	IpsAction *int64 `json:"IpsAction,omitempty" name:"IpsAction"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// cidr

	VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type AddVpcFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 限速的类型，可以是具体的IP也可以是CIDR，比如192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 需要添加限速规则的防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
	// 入向限制速度，单位是kbps

	InLimit *uint64 `json:"InLimit,omitempty" name:"InLimit"`
	// 0&nbsp;IPV4&nbsp;1&nbsp;IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 限速类型
	// IP限速：IP
	// 端口限速：PORT
	// 当前仅支持IP限速

	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`
	// 出向限制速度，单位是kbps

	OutLimit *uint64 `json:"OutLimit,omitempty" name:"OutLimit"`
	// 流量控制策略的端口。取值：
	// -1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

func (r *AddVpcFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcRuleOverviewData struct {

	// 启用规则数量

	StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`
	// 阻断策略规则数量

	StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`
}

type DescribeBorderACLListRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 出站还是入站，1：入站，0：出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 规则是否开启，'0':&nbsp;未开启，'1':&nbsp;开启,&nbsp;默认为'0'

	Status *string `json:"Status,omitempty" name:"Status"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeBorderACLListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderACLListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTiWhiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTiWhiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTiWhiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNewBindRoleUserRequest struct {
	*tchttp.BaseRequest

	// 是否发起扫描，0不发起，1发起

	IsScan *int64 `json:"IsScan,omitempty" name:"IsScan"`
}

func (r *AddNewBindRoleUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNewBindRoleUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupApiRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0:添加成功，非0：添加失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupApiRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWInfoRequest struct {
	*tchttp.BaseRequest

	// -

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeCFWInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportBlockIgnoreRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportBlockIgnoreRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportBlockIgnoreRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescDNSFWSwitch struct {

	// 请求频率均值

	AverageFrequency *string `json:"AverageFrequency,omitempty" name:"AverageFrequency"`
	// DNS服务器IP

	DNS *string `json:"DNS,omitempty" name:"DNS"`
	// 0-关闭，1-开关，2-NAT已防护，3-DNS&nbsp;IP非云，4-不支持

	DNSSwitch *uint64 `json:"DNSSwitch,omitempty" name:"DNSSwitch"`
	// 请求频率峰值

	PeakFrequency *string `json:"PeakFrequency,omitempty" name:"PeakFrequency"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 风险请求次数

	RiskRequestCnt *string `json:"RiskRequestCnt,omitempty" name:"RiskRequestCnt"`
	// 累计请求次数

	TotalRequestCnt *string `json:"TotalRequestCnt,omitempty" name:"TotalRequestCnt"`
	// vpc&nbsp;cidr

	VpcCIDR *string `json:"VpcCIDR,omitempty" name:"VpcCIDR"`
	// vpc&nbsp;id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type ServiceTemplateSpecification struct {

	// 协议端口组ID，例如：ppmg-f5n1f8da

	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
	// &nbsp;协议端口ID，例如：ppm-f5n1f8da。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type DescribeNatNewFlowStatsDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 入向峰值带宽，单位bps

		InputMax *int64 `json:"InputMax,omitempty" name:"InputMax"`
		// 入向峰值带宽配额占比

		InputPercent *float64 `json:"InputPercent,omitempty" name:"InputPercent"`
		// 入向带宽峰值折线图

		InputTrendData []*TrendsItem `json:"InputTrendData,omitempty" name:"InputTrendData"`
		// 实例带宽规格，单位Mbps

		InstanceBandWidth *int64 `json:"InstanceBandWidth,omitempty" name:"InstanceBandWidth"`
		// 出向峰值带宽，单位bps

		OutputMax *int64 `json:"OutputMax,omitempty" name:"OutputMax"`
		// 出向峰值带宽配额占比

		OutputPercent *float64 `json:"OutputPercent,omitempty" name:"OutputPercent"`
		// 出向带宽峰值折线图

		OutputTrendData []*TrendsItem `json:"OutputTrendData,omitempty" name:"OutputTrendData"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatNewFlowStatsDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatNewFlowStatsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseWhiteListRuleData struct {

	// 云厂商码

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 规则描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 是否启用，0&nbsp;不启用，1启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 地域码1

	FirstLevelRegionCode *int64 `json:"FirstLevelRegionCode,omitempty" name:"FirstLevelRegionCode"`
	// 地域名称1

	FirstLevelRegionName *string `json:"FirstLevelRegionName,omitempty" name:"FirstLevelRegionName"`
	// 是否云厂商规则，0不是&nbsp;1&nbsp;时

	IsCloudRule *int64 `json:"IsCloudRule,omitempty" name:"IsCloudRule"`
	// 是否地域规则，0不是&nbsp;1是

	IsRegionRule *int64 `json:"IsRegionRule,omitempty" name:"IsRegionRule"`
	// 地域码2

	SecondLevelRegionCode *int64 `json:"SecondLevelRegionCode,omitempty" name:"SecondLevelRegionCode"`
	// 地域名称2

	SecondLevelRegionName *string `json:"SecondLevelRegionName,omitempty" name:"SecondLevelRegionName"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 访问源类型，1&nbsp;ip；6&nbsp;实例；100&nbsp;资源分组

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 访问目的类型，1&nbsp;ip；6&nbsp;实例；100&nbsp;资源分组

	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
}

type ModuleConf struct {

	// 资产中心版本，旧版0，新版1，tce默认0，公有云默认1

	AssetCenterVersion *int64 `json:"AssetCenterVersion,omitempty" name:"AssetCenterVersion"`
	// 是否支持Barad&nbsp;1支持&nbsp;0不支持

	BaradBandWidth *int64 `json:"BaradBandWidth,omitempty" name:"BaradBandWidth"`
	// 是否支持安全基线&nbsp;1支持&nbsp;0不支持

	BaseLine *int64 `json:"BaseLine,omitempty" name:"BaseLine"`
	// tce&nbsp;计费模式&nbsp;1&nbsp;计量&nbsp;0&nbsp;计费&nbsp;&nbsp;2异常

	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`
	// 是否支持批量操作互联网规则&nbsp;1支持&nbsp;0不支持

	BorderAclBulk *int64 `json:"BorderAclBulk,omitempty" name:"BorderAclBulk"`
	// 是否支持互联网边界ACL&nbsp;特性&nbsp;1支持&nbsp;0不支持

	BorderAclSpecial *int64 `json:"BorderAclSpecial,omitempty" name:"BorderAclSpecial"`
	// 边界&nbsp;旁路&nbsp;1支持&nbsp;0不支持

	BorderBypass *int64 `json:"BorderBypass,omitempty" name:"BorderBypass"`
	// 边界&nbsp;山石&nbsp;1支持&nbsp;0不支持

	BorderHillStone *int64 `json:"BorderHillStone,omitempty" name:"BorderHillStone"`
	// 边界&nbsp;串行&nbsp;1支持&nbsp;0不支持

	BorderSerial *int64 `json:"BorderSerial,omitempty" name:"BorderSerial"`
	// 是否支持vpcfw的自定义路由&nbsp;1支持&nbsp;0不支持

	CustomRoute *int64 `json:"CustomRoute,omitempty" name:"CustomRoute"`
	// 是否支持ips白名单自定义策略&nbsp;1支持&nbsp;0不支持

	CustomWhiteRule *int64 `json:"CustomWhiteRule,omitempty" name:"CustomWhiteRule"`
	// 是否支持DNS防火墙&nbsp;1支持&nbsp;0不支持

	DNSFW *int64 `json:"DNSFW,omitempty" name:"DNSFW"`
	// 用户是否为高级版折扣用户时&nbsp;1支持&nbsp;0不支持

	DiscountUser *int64 `json:"DiscountUser,omitempty" name:"DiscountUser"`
	// 是否允许域名解析IP能力&nbsp;1支持&nbsp;0不支持

	DnsAcl *int64 `json:"DnsAcl,omitempty" name:"DnsAcl"`
	// 环境&nbsp;380/3100/3110/3120

	EnvVersion *int64 `json:"EnvVersion,omitempty" name:"EnvVersion"`
	// 是否有互联网边界&nbsp;1支持&nbsp;0不支持

	Flow *int64 `json:"Flow,omitempty" name:"Flow"`
	// FlowCenterVersio&nbsp;1支持&nbsp;0不支持n

	FlowCenterVersion *int64 `json:"FlowCenterVersion,omitempty" name:"FlowCenterVersion"`
	// 是否支持网络蜜罐&nbsp;1支持&nbsp;0不支持，TCE默认关闭

	HoneypotVersion *int64 `json:"HoneypotVersion,omitempty" name:"HoneypotVersion"`
	// 是否支持自动化助手-互联网边界&nbsp;1支持&nbsp;0不支持

	InternetACLBetaTask *int64 `json:"InternetACLBetaTask,omitempty" name:"InternetACLBetaTask"`
	// 1&nbsp;多语言环境&nbsp;0&nbsp;简体中文

	Intl *uint64 `json:"Intl,omitempty" name:"Intl"`
	// 边界墙IDPS规则控制&nbsp;1支持&nbsp;0不支持

	IpsActionBorder *int64 `json:"IpsActionBorder,omitempty" name:"IpsActionBorder"`
	// 是否有零信任防护&nbsp;1支持&nbsp;0不支持

	IsShowIdentityAuth *int64 `json:"IsShowIdentityAuth,omitempty" name:"IsShowIdentityAuth"`
	// 是否有NAT&nbsp;1支持&nbsp;0不支持

	IsShowNat *int64 `json:"IsShowNat,omitempty" name:"IsShowNat"`
	// 是否支持日志投递&nbsp;1支持&nbsp;0不支持，TCE默认关闭

	LogDelivery *int64 `json:"LogDelivery,omitempty" name:"LogDelivery"`
	// 是否支持ulp，编辑&nbsp;&nbsp;1支持&nbsp;0不支持日志存储功能

	LogStorage *int64 `json:"LogStorage,omitempty" name:"LogStorage"`
	// 是否支持安全组&nbsp;1支持&nbsp;0不支持

	SecurityGroup *int64 `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
	// 是否支持natdns开关&nbsp;1支持&nbsp;0不支持

	ShowNatDnsSwitch *int64 `json:"ShowNatDnsSwitch,omitempty" name:"ShowNatDnsSwitch"`
	// 是否支持选择扩展网段优先&nbsp;1支持&nbsp;0不支持

	SupportAssisCidr *int64 `json:"SupportAssisCidr,omitempty" name:"SupportAssisCidr"`
	// TCE&nbsp;是否支持租户&nbsp;1支持&nbsp;0不支持端COS

	SupportCos *int64 `json:"SupportCos,omitempty" name:"SupportCos"`
	// 是否支持NTA&nbsp;1支持&nbsp;0不支持

	SupportNTA *int64 `json:"SupportNTA,omitempty" name:"SupportNTA"`
	// 是否支持syslog&nbsp;日志投递&nbsp;1支持&nbsp;0不支持

	SupportSysLog *int64 `json:"SupportSysLog,omitempty" name:"SupportSysLog"`
	// 是否支持试用&nbsp;1支持&nbsp;0不支持

	SupportTrial *int64 `json:"SupportTrial,omitempty" name:"SupportTrial"`
	// 是否支持IPV6&nbsp;1支持&nbsp;0不支持

	SupportVpcIPV6 *int64 `json:"SupportVpcIPV6,omitempty" name:"SupportVpcIPV6"`
	// 是否支持情报中心&nbsp;1支持&nbsp;0不支持

	TiCenter *int64 `json:"TiCenter,omitempty" name:"TiCenter"`
	// 跨租户开关&nbsp;&nbsp;1支持&nbsp;0不支持

	VpcCrossTenant *int64 `json:"VpcCrossTenant,omitempty" name:"VpcCrossTenant"`
	// VPC是否支持入侵防御&nbsp;1支持&nbsp;0不支持

	VpcMode *int64 `json:"VpcMode,omitempty" name:"VpcMode"`
	// 是否支持告警中心&nbsp;1支持&nbsp;0不支持，TCE默认关闭

	WarnCenterVersion *int64 `json:"WarnCenterVersion,omitempty" name:"WarnCenterVersion"`
	// 是否支持私有云CDCVPC

	SupportCdcVpc *int64 `json:"SupportCdcVpc,omitempty" name:"SupportCdcVpc"`
}

type ZeroTrustDbAssetInfo struct {

	// 接入状态&nbsp;

	AccessStatus *int64 `json:"AccessStatus,omitempty" name:"AccessStatus"`
	// 实例id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 实例名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 实例类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 授权用户/组织

	AuthUsers *string `json:"AuthUsers,omitempty" name:"AuthUsers"`
	// 地址

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 防护状态&nbsp;0&nbsp;未防护&nbsp;1&nbsp;已防护

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 支持域名接入&nbsp;0&nbsp;不支持&nbsp;1支持

	SupportDomainAccess *int64 `json:"SupportDomainAccess,omitempty" name:"SupportDomainAccess"`
	// 资源标签

	Tag []*Tags `json:"Tag,omitempty" name:"Tag"`
	// 最近访问时间

	UserLastVisitTime *string `json:"UserLastVisitTime,omitempty" name:"UserLastVisitTime"`
	// 所属私有网络id

	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`
	// 所属私有网络

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 运维端口

	YWPort *int64 `json:"YWPort,omitempty" name:"YWPort"`
}

type FwGateway struct {

	// 防火墙网关id

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关ip地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 网关所属vpc&nbsp;id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ClearResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0清空成功，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefenseError struct {

	// 错误类型区分

	ErrKey *string `json:"ErrKey,omitempty" name:"ErrKey"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 错误ip列表

	ErrorIpInfo []*string `json:"ErrorIpInfo,omitempty" name:"ErrorIpInfo"`
	// 错误内部id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 错误时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type InstanceKeyInfo struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 托管密码

	PassWordDatas []*PassWordData `json:"PassWordDatas,omitempty" name:"PassWordDatas"`
	// 托管密钥

	SecretKeyDatas []*SecretKeyData `json:"SecretKeyDatas,omitempty" name:"SecretKeyDatas"`
}

type ModifyAssetScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态值&nbsp;0：成功，1&nbsp;执行扫描中,其他：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssetScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalanceIns struct {

	// 负载均衡id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
}

type PortRiskData struct {

	// 风险端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 端口服务

	Server *string `json:"Server,omitempty" name:"Server"`
	// 防火墙防护状态，1：已生成防护规则，2：未处理

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修复建议，1：保持现状，2：限制访问，3：建议关闭

	Suggest *int64 `json:"Suggest,omitempty" name:"Suggest"`
	// 开关状态，0：关闭，1：开启

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 组件

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 规则uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeleteClsDeliverTaskRequest struct {
	*tchttp.BaseRequest

	// 来源

	SourceFrom *string `json:"SourceFrom,omitempty" name:"SourceFrom"`
}

func (r *DeleteClsDeliverTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClsDeliverTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBetaTaskAclRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则类型1border&nbsp;2nat&nbsp;3vpc&nbsp;4sg

		AclType *int64 `json:"AclType,omitempty" name:"AclType"`
		// 互联网规则列表

		BorderData []*BorderACL `json:"BorderData,omitempty" name:"BorderData"`
		// nat规则列表

		NatData []*DescAcItem `json:"NatData,omitempty" name:"NatData"`
		// 安全组规则列表

		SecurityGroupData []*BetaSGAcl `json:"SecurityGroupData,omitempty" name:"SecurityGroupData"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// vpc规则列表

		VpcData []*VpcRuleItem `json:"VpcData,omitempty" name:"VpcData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBetaTaskAclRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskAclRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHitRuleLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*RuleLogsData `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHitRuleLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHitRuleLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupChangeRule struct {

	// 规则

	PolicySet *SecurityGroupPolicySet `json:"PolicySet,omitempty" name:"PolicySet"`
	// 安全组id

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
}

type WarningCenterOverviewData struct {

	// 高危告警

	HighRiskNum *uint64 `json:"HighRiskNum,omitempty" name:"HighRiskNum"`
	// 低危告警

	LowRiskNum *uint64 `json:"LowRiskNum,omitempty" name:"LowRiskNum"`
	// 中危告警

	MidRiskNum *uint64 `json:"MidRiskNum,omitempty" name:"MidRiskNum"`
	// 严重告警

	SeriousNum *uint64 `json:"SeriousNum,omitempty" name:"SeriousNum"`
	// 提示告警

	SuggestRiskNum *uint64 `json:"SuggestRiskNum,omitempty" name:"SuggestRiskNum"`
}

type ZeroTrustSelectAssets struct {

	// 资产实例Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产类型，web&nbsp;WEB服务，ins&nbsp;服务七，db&nbsp;数据库

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 权限细则是否能被删除

	CanDelete *int64 `json:"CanDelete,omitempty" name:"CanDelete"`
	// 资产实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// web服务，ip状态

	IpStatus *bool `json:"IpStatus,omitempty" name:"IpStatus"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域Id

	RegionKey *string `json:"RegionKey,omitempty" name:"RegionKey"`
	// 关联权限Id

	Rules *string `json:"Rules,omitempty" name:"Rules"`
	// WEB服务地址

	ServiceAddress *string `json:"ServiceAddress,omitempty" name:"ServiceAddress"`
	// WEB服务域名

	ServiceDomain *string `json:"ServiceDomain,omitempty" name:"ServiceDomain"`
	// WEB服务端口

	ServicePort *int64 `json:"ServicePort,omitempty" name:"ServicePort"`
	// 服务类型，Public公网服务，Private内网服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// WEB服务类型，1内网，2公网

	WebType *int64 `json:"WebType,omitempty" name:"WebType"`
}

type DescribeAssetStatisticRequest struct {
	*tchttp.BaseRequest

	// 资产id

	InsId *string `json:"InsId,omitempty" name:"InsId"`
	// 时间类型&nbsp;1.24小时&nbsp;2.7天3.30天&nbsp;4.6个月

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeAssetStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZeroTrustBlockStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZeroTrustBlockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZeroTrustBlockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupRule struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)

	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`
	// 访问目的类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 规则对应的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 规则顺序，-1表示最低，1表示最高

	OrderIndex *string `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议；TCP/UDP/ICMP/ANY

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 端口协议类型参数模板id；协议端口模板id；与Protocol,Port互斥

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 规则对应的唯一内部id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

type DescribeAllZoneListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		Data []*ZoneConfig `json:"Data,omitempty" name:"Data"`
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllZoneListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwRuleHitDetailRequest struct {
	*tchttp.BaseRequest

	// 规则对应的唯一id

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeNatFwRuleHitDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwRuleHitDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteVpcInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSwitchListRequest struct {
	*tchttp.BaseRequest

	// 筛选NAT防火墙子网开关所属地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 条数，分页用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选NAT防火墙子网开关所属NAT网关

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 筛选NAT防火墙子网开关所属NAT防火墙实例

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// 偏移量，分页用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开关，1打开，0关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 筛选NAT防火墙子网开关所属VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeNatSwitchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSwitchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeCommonResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关基础信息

		Data *NoticeCommon `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNoticeCommonResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeCommonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeNameRequest struct {
	*tchttp.BaseRequest

	// 改名的边edge&nbsp;ID

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 改名新名称

	EdgeNewName *string `json:"EdgeNewName,omitempty" name:"EdgeNewName"`
}

func (r *ModifyEdgeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleOverviewData struct {

	// 规则总数

	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 剩余配额

	RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`
	// 启用规则数量

	StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`
	// 停用规则数量

	StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`
	// 阻断策略规则数量

	StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`
}

type DescribeAssetScanStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetScanStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertCenterEventNameSelectListRequest struct {
	*tchttp.BaseRequest

	// 资产实例ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 欺骗告警策略类型，Strategy=Deception时生效，端口探测Detect，入侵事件Invade，横向移动Move

	DeceptionEventType *string `json:"DeceptionEventType,omitempty" name:"DeceptionEventType"`
	// 方向，1入站，0出站，3横向移动，为空不筛选

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 检索结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 筛选攻击链

	KillChainList []*string `json:"KillChainList,omitempty" name:"KillChainList"`
	// 检索开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 告警策略：观察Alert,拦截Block,欺骗Deception

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeAlertCenterEventNameSelectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertCenterEventNameSelectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoInternetSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启新增资产自动开启云防火墙

		AutoInternetSwitch *int64 `json:"AutoInternetSwitch,omitempty" name:"AutoInternetSwitch"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoInternetSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoInternetSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetScanInfo struct {

	// 资产ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 扫描深度

	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`
	// 扫描状态

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 最近一次扫描成功时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 端口风险数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetEventListStruct struct {

	// 无

	AffectAssetNum *int64 `json:"AffectAssetNum,omitempty" name:"AffectAssetNum"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 无

	EventTitle *string `json:"EventTitle,omitempty" name:"EventTitle"`
	// 无

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 无

	Level *string `json:"Level,omitempty" name:"Level"`
	// 无

	Related *int64 `json:"Related,omitempty" name:"Related"`
	// 无

	Time *string `json:"Time,omitempty" name:"Time"`
	// 无

	Url *string `json:"Url,omitempty" name:"Url"`
}

type DescribeEventAssetInfoRequest struct {
	*tchttp.BaseRequest

	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeEventAssetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventAssetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleRequest struct {
	*tchttp.BaseRequest

	// 需要编辑的规则数组

	Rules []*VpcRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyVpcAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SysLogTypeConfig struct {

	// 投递错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 投递类型
	// 范围&nbsp;
	// internetFlowLog,natFlowLog,vpcFlowLog,dnsFlowLog,eventLog,internetRuleLog,natRuleLog,vpcRuleLog,dnsRuleLog,operateLog,borderNtaFlowLog
	// 这里的类型和LogTypeName&nbsp;含义对应

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 类型名称
	// 范围
	// 流量日志-互联网边界,流量日志-NAT边界,流量日志-VPC边界,流量日志-DNS,入侵防御日志,访问控制日志-互联网边界,访问控制日志-NAT边界,访问控制日志-VPC边界,访问控制日志-DNS,操作日志,流量分析日志-互联网边界

	LogTypeName *string `json:"LogTypeName,omitempty" name:"LogTypeName"`
	// 投递状态:1:正常&nbsp;2:存在错误&nbsp;3:失败&nbsp;0&nbsp;init

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 投递开关&nbsp;1:开启&nbsp;2:关闭

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// tag

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 预留kafka&nbsp;投递topic名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 预留kafka&nbsp;投递topicid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type CreateWarnEventRulesRequest struct {
	*tchttp.BaseRequest

	// 拦截/忽略天数，可选：1&nbsp;天&nbsp;or&nbsp;7&nbsp;天

	EffectiveTime *uint64 `json:"EffectiveTime,omitempty" name:"EffectiveTime"`
	// id&nbsp;数组

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 忽略原因，1：重复，2：误报

	IgnoreReason *uint64 `json:"IgnoreReason,omitempty" name:"IgnoreReason"`
	// 下发规则类型，1：拦截，2：忽略

	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *CreateWarnEventRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWarnEventRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpsRuleListNewRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填	官网公开展示

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeIpsRuleListNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsRuleListNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupAssociationRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数据

		RuleImpactList *RuleImpact `json:"RuleImpactList,omitempty" name:"RuleImpactList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupAssociationRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupAssociationRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarnEventListRequest struct {
	*tchttp.BaseRequest

	// 资产&nbsp;Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，desc：降序，asc：升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 判断来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 事件处置状态过滤，0：显示全部，1：只显示未处置

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeWarnEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarnEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 忽略条件导出

	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`
	// 导出字段

	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *ExportNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleRegionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		Data []*RuleRegionCount `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleRegionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleRegionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwPricesInfo struct {

	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 价格

	Price *string `json:"Price,omitempty" name:"Price"`
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费周期&nbsp;day天&nbsp;hour&nbsp;小时

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type FwGroupSwitchFlowStat struct {

	// 均值带宽

	FlowAvg *float64 `json:"FlowAvg,omitempty" name:"FlowAvg"`
	// 峰值流量

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 开关ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 时间

	UpTime *string `json:"UpTime,omitempty" name:"UpTime"`
}

type HoneyPotInfo struct {

	// 捕获攻击者

	AttackerCount *int64 `json:"AttackerCount,omitempty" name:"AttackerCount"`
	// 诱饵

	Bait *string `json:"Bait,omitempty" name:"Bait"`
	// 命中次数

	HitTimes *int64 `json:"HitTimes,omitempty" name:"HitTimes"`
	// 交互类型：高、中、低交互

	InteractionType *string `json:"InteractionType,omitempty" name:"InteractionType"`
	// 连接探针数量

	ProbeCount *int64 `json:"ProbeCount,omitempty" name:"ProbeCount"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 蜜罐服务唯一ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务类型：ssh、rdp等

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 服务实例状态&nbsp;1&nbsp;未创建&nbsp;2&nbsp;创建中&nbsp;&nbsp;3&nbsp;正删除&nbsp;4&nbsp;开关转变中&nbsp;0完成&nbsp;-1&nbsp;失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 服务开关&nbsp;0关闭&nbsp;1打开

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

type DescribeIdsWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Data []*IdsWhiteInfo `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdsWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdsWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true是防火墙用户，false不是防火墙用户

		Status *bool `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverViewEdgeFlowStat struct {

	// 时间

	Label *string `json:"Label,omitempty" name:"Label"`
	// 带宽方向

	Name *string `json:"Name,omitempty" name:"Name"`
	// 带宽值&nbsp;单位&nbsp;bps

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeBlockTimesListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折线数据

		Data []*IpStatic `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockTimesListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockTimesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchErrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息列表

		Data []*SwitchError `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchErrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserBandwidthUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽数据列表

		Data []*UserBandwidthUsage `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserBandwidthUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserBandwidthUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetWebScanRequest struct {
	*tchttp.BaseRequest

	// 资产类型，1为ip&nbsp;2为Web服务&nbsp;3为两者都有

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 扫描类型：1立即扫描&nbsp;2&nbsp;周期任务

	RangeType *int64 `json:"RangeType,omitempty" name:"RangeType"`
	// 扫描深度：'heavy',&nbsp;'medium',&nbsp;'light'

	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`
	// 立即扫描这个字段传过滤的IP资产扫描集合

	ScanFilterIp []*string `json:"ScanFilterIp,omitempty" name:"ScanFilterIp"`
	// 立即扫描这个字段传过滤的Web资产ID扫描集合

	ScanFilterWebId []*string `json:"ScanFilterWebId,omitempty" name:"ScanFilterWebId"`
	// RangeType为2&nbsp;是必须添加，定时任务时间

	ScanPeriod *string `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
	// 扫描范围：1端口,&nbsp;2端口+漏扫,&nbsp;3端口+弱口令,&nbsp;4端口+漏扫+弱口令

	ScanRange *int64 `json:"ScanRange,omitempty" name:"ScanRange"`
}

func (r *ModifyAssetWebScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetWebScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnatInfo struct {

	// 远程运维Dnat规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// Dnat协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Dnat外部IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// Dnat外部端口

	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`
}

type DescribeCfwBandWidthRequest struct {
	*tchttp.BaseRequest

	// 最大带宽时间段，默认为1

	DayNum *int64 `json:"DayNum,omitempty" name:"DayNum"`
}

func (r *DescribeCfwBandWidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwBandWidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceWithRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例数组

		NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstanceWithRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstanceWithRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启用规则数量

		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`
		// 阻断策略规则数量

		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`
		// 规则总量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcRuleOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;不自动选择子网
	// 1&nbsp;自动选择子网创建私有连接

	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitempty" name:"AutoChooseSubnet"`
	// 操作开关详情

	EdgeIpSwitchLst []*EdgeIpSwitch `json:"EdgeIpSwitchLst,omitempty" name:"EdgeIpSwitchLst"`
	// 0&nbsp;关闭开关
	// 1&nbsp;打开开关
	// 2&nbsp;不操作开关，此次切换模式

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 0&nbsp;切换为旁路
	// 1&nbsp;切换为串行
	// 2&nbsp;不切换模式，此次操作开关

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

func (r *ModifyEdgeIpSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCfwInsBypassRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 防火墙类型，nat：nat防火墙，ew：东西向防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 操作，open：开启，close：关闭

	Opr *string `json:"Opr,omitempty" name:"Opr"`
	// Normal:&nbsp;正常模式&nbsp;OnlyRoute:&nbsp;透明模式

	TrafficMode *string `json:"TrafficMode,omitempty" name:"TrafficMode"`
}

func (r *SetCfwInsBypassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCfwInsBypassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplyTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请失败的原因

		FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplyTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosBucket struct {

	// 创建时间

	CreationDate *string `json:"CreationDate,omitempty" name:"CreationDate"`
	// 存储桶名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 存储桶地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type FlowCenterOutAddress struct {

	// 地理位置

	Address *string `json:"Address,omitempty" name:"Address"`
	// 请求结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是互联网还是nat

	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`
	// 访问目的

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 目的端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 访问流量请求

	RequestFlow *int64 `json:"RequestFlow,omitempty" name:"RequestFlow"`
	// 访问流量响应

	ResponseFlow *int64 `json:"ResponseFlow,omitempty" name:"ResponseFlow"`
	// 风险评估

	RiskAssess *string `json:"RiskAssess,omitempty" name:"RiskAssess"`
	// 请求开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 会话数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建结果,0成功&nbsp;非0失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一Id

		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteBackupRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 下一跳id，vpc间防火墙的云联网模式下的下一跳id为ccn的路由表id，形如ccnrtb-xxx

	NextHopId *string `json:"NextHopId,omitempty" name:"NextHopId"`
	// 对象id，vpc间防火墙的云联网模式下的对象id为接入ccn的实例id，如vpc-xxx

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
}

func (r *ModifyRouteBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpsLogsData struct {

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 规则描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 判断来源

	DiffFrom *string `json:"DiffFrom,omitempty" name:"DiffFrom"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 告警事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 是否下发过规则

	GenRule *uint64 `json:"GenRule,omitempty" name:"GenRule"`
	// ID值

	Id *string `json:"Id,omitempty" name:"Id"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源端口

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 威胁描述

	ThreatDesc *string `json:"ThreatDesc,omitempty" name:"ThreatDesc"`
	// 家族组织

	ThreatOrganization *string `json:"ThreatOrganization,omitempty" name:"ThreatOrganization"`
	// 处置建议

	ThreatSuggestion *string `json:"ThreatSuggestion,omitempty" name:"ThreatSuggestion"`
	// 战术技术过程

	ThreatTtps *string `json:"ThreatTtps,omitempty" name:"ThreatTtps"`
	// 威胁类型

	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type SecretKeyData struct {

	// 密钥内容

	SecretKeyContent *string `json:"SecretKeyContent,omitempty" name:"SecretKeyContent"`
	// 密钥文件名

	SecretKeyFileName *string `json:"SecretKeyFileName,omitempty" name:"SecretKeyFileName"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeAssetOverviewNewRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeAssetOverviewNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetOverviewNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwInsCvmListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data []*InsInfo `json:"Data,omitempty" name:"Data"`
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwInsCvmListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsCvmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatProbeEipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回集合

		NatProbeEipList []*NatProbeEip `json:"NatProbeEipList,omitempty" name:"NatProbeEipList"`
		// 是否需要升级引擎支持nat拨测&nbsp;1&nbsp;需要&nbsp;0不需要

		NeedProbeEngineUpdate *int64 `json:"NeedProbeEngineUpdate,omitempty" name:"NeedProbeEngineUpdate"`
		// 支持最多绑定的ip

		ProbeEipLimit *int64 `json:"ProbeEipLimit,omitempty" name:"ProbeEipLimit"`
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatProbeEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatProbeEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatInstanceInfo struct {

	// 实例带宽大小&nbsp;Mbps

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 公网ip数组

	EipAddress []*string `json:"EipAddress,omitempty" name:"EipAddress"`
	// 弹性带宽，单位Mbps

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 弹性开关
	// 1&nbsp;打开
	// 0&nbsp;关闭

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// 实例引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 0:&nbsp;新增模式，1:接入模式

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
	// 入向带宽峰值&nbsp;bps

	InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`
	// 是否首次开通按量付费
	// 1&nbsp;是
	// 0&nbsp;不是

	IsFirstAfterPay *int64 `json:"IsFirstAfterPay,omitempty" name:"IsFirstAfterPay"`
	// nat实例id

	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`
	// nat实例名称

	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`
	// 是的需要升级引擎&nbsp;支持&nbsp;nat拨测&nbsp;1需要&nbsp;0不需要

	NeedProbeEngineUpdate *int64 `json:"NeedProbeEngineUpdate,omitempty" name:"NeedProbeEngineUpdate"`
	// 出向带宽峰值&nbsp;bps

	OutFlowMax *uint64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域区域信息

	RegionDetail *string `json:"RegionDetail,omitempty" name:"RegionDetail"`
	// 地域中文信息

	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`
	// 引擎预约升级时间

	ReserveTime *string `json:"ReserveTime,omitempty" name:"ReserveTime"`
	// 引擎预约升级版本

	ReserveVersion *string `json:"ReserveVersion,omitempty" name:"ReserveVersion"`
	// 引擎预约升级版本状态

	ReserveVersionState *string `json:"ReserveVersionState,omitempty" name:"ReserveVersionState"`
	// 实例的规则限制最大规格数

	RuleMax *uint64 `json:"RuleMax,omitempty" name:"RuleMax"`
	// 已使用规则数

	RuleUsed *uint64 `json:"RuleUsed,omitempty" name:"RuleUsed"`
	// 0&nbsp;:正常&nbsp;1：正在初始化

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例关联子网数组

	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`
	// 引擎运行模式，Normal:正常,&nbsp;OnlyRoute:透明模式

	TrafficMode *string `json:"TrafficMode,omitempty" name:"TrafficMode"`
	// 引擎是否可升级：0，不可升级；1，可升级

	UpdateEnable *int64 `json:"UpdateEnable,omitempty" name:"UpdateEnable"`
	// 内外使用ip数组

	VpcIp []*string `json:"VpcIp,omitempty" name:"VpcIp"`
	// 实例主所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例备所在可用区

	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`
	// 实例所在可用区

	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`
	// 实例所在可用区

	ZoneZhBak *string `json:"ZoneZhBak,omitempty" name:"ZoneZhBak"`
	// 实例所在引擎机器信息

	FwCvmList *VpcFwCvmInsInfo `json:"FwCvmList,omitempty" name:"FwCvmList"`
}

type SwitchListsData struct {

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网IP

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
	// 最近扫描时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 风险端口数

	PortTimes *uint64 `json:"PortTimes,omitempty" name:"PortTimes"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 公网&nbsp;IP&nbsp;类型

	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 扫描深度&nbsp;light轻度&nbsp;

	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`
	// 扫描状态

	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 防火墙开关

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

type DescribeCcnSASEEdgeRequest struct {
	*tchttp.BaseRequest

	// CCN&nbsp;ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DescribeCcnSASEEdgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnSASEEdgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP连接数统计列表

		Data []*IpConnections `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcIpConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcIpConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicDetail struct {

	// 影响面

	Affect *string `json:"Affect,omitempty" name:"Affect"`
	// 攻击类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// Cve编号

	Cve *string `json:"Cve,omitempty" name:"Cve"`
	// 安全事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 自增id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 保留字段

	Message *string `json:"Message,omitempty" name:"Message"`
	// 消息类型&nbsp;1(漏洞信息),2(产品功能动态)

	News *int64 `json:"News,omitempty" name:"News"`
	// 发布时间

	ReleaseDate *string `json:"ReleaseDate,omitempty" name:"ReleaseDate"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 文章标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
}

type PieData struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type QueryFilter struct {

	// 过滤key

	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`
	// 操作符(只支持32位)

	FilterOperatorType *int64 `json:"FilterOperatorType,omitempty" name:"FilterOperatorType"`
	// 过滤value

	FilterValue *string `json:"FilterValue,omitempty" name:"FilterValue"`
}

type SecurityGroupApiRuleData struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 执行顺序，中间插入必传，前插、后插非必传

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口,&nbsp;当Protocol为ANY或ICMP时，Port为-1/-1

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，支持ANY/TCP/UDP/ICMP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 规则类型，1：VpcId+Ip/Cidr,&nbsp;2:&nbsp;实例ID，入站时为访问目的类型，出站时为访问源类型

	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`
	// 访问源，入站时为Ip/Cidr，默认为0.0.0.0/0；&nbsp;出站时当RuleType为1时，支持内网Ip/Cidr,&nbsp;当RuleType为2时，填实例ID

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 策略,&nbsp;1：阻断，2：放行

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的，出站时为Ip/Cidr，默认为0.0.0.0/0；入站时当RuleType为1时，支持内网Ip/Cidr,&nbsp;当RuleType为2时，填实例ID

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 私有网络ID，当RuleType为1时必传

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeBillingReleaseResourceCheckRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBillingReleaseResourceCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingReleaseResourceCheckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwLimitsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfwLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetflowRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：修改成功，非0：修改失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetflowRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetflowRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockIgnoreRule struct {

	// 规则类型：1封禁，2放通

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 地理位置

	Address *string `json:"Address,omitempty" name:"Address"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 自定义规则细节

	CustomRule *CustomWhiteRule `json:"CustomRule,omitempty" name:"CustomRule"`
	// 方向：1入站，0出站

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 所有方向聚合成字符串

	DirectionList *string `json:"DirectionList,omitempty" name:"DirectionList"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 来源事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 1&nbsp;border&nbsp;2&nbsp;nat&nbsp;4&nbsp;vpc&nbsp;8&nbsp;border-serial

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 忽略原因

	IgnoreReason *string `json:"IgnoreReason,omitempty" name:"IgnoreReason"`
	// 规则ip

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 白名单信息

	IocInfo *string `json:"IocInfo,omitempty" name:"IocInfo"`
	// 资产实例名称、自定义策略名称等

	IocName *string `json:"IocName,omitempty" name:"IocName"`
	// 上次命中时间

	LastHitTime *string `json:"LastHitTime,omitempty" name:"LastHitTime"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 规则命中次数

	MatchTimes *int64 `json:"MatchTimes,omitempty" name:"MatchTimes"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 1&nbsp;封禁&nbsp;2外部IP&nbsp;3域名&nbsp;4情报&nbsp;5assets&nbsp;6udf&nbsp;&nbsp;7入侵防御规则id&nbsp;（2-7属于白名单类型）

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 安全事件来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 规则生效开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 规则id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
}

type HoneyPotNameInfo struct {

	// 蜜罐服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 蜜罐服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ModifyVpcFwReSelectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwReSelectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwReSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZeroTrustBlockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 按钮状态
		// 0&nbsp;关闭
		// 1&nbsp;打开

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZeroTrustBlockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZeroTrustBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRunSyncAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：同步成功，1：资产更新中，2：后台同步调用失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRunSyncAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogSelectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问控制类型

		ACLTypeLists []*IntValueText `json:"ACLTypeLists,omitempty" name:"ACLTypeLists"`
		// 操作行为

		ActionLists []*string `json:"ActionLists,omitempty" name:"ActionLists"`
		// 操作类型

		ActionTypeLists []*string `json:"ActionTypeLists,omitempty" name:"ActionTypeLists"`
		// 资产分类

		AssetTypeList []*StringValueText `json:"AssetTypeList,omitempty" name:"AssetTypeList"`
		// 安全基线类型

		BaseLineLists []*string `json:"BaseLineLists,omitempty" name:"BaseLineLists"`
		// 防火墙类型

		FwTypeLists []*IntValueText `json:"FwTypeLists,omitempty" name:"FwTypeLists"`
		// 零信任防护类型

		IdentityAuthTypeList []*string `json:"IdentityAuthTypeList,omitempty" name:"IdentityAuthTypeList"`
		// 防火墙实例类型

		InstanceFwTypeLists []*IntValueText `json:"InstanceFwTypeLists,omitempty" name:"InstanceFwTypeLists"`
		// 日志操作类型

		LogStorageLists []*string `json:"LogStorageLists,omitempty" name:"LogStorageLists"`
		// 地域列表

		RegionLists []*string `json:"RegionLists,omitempty" name:"RegionLists"`
		// 已创建老版安全组的的地域

		RegionSafe []*string `json:"RegionSafe,omitempty" name:"RegionSafe"`
		// 功能模块

		SystemLogModuleLists []*StringValueText `json:"SystemLogModuleLists,omitempty" name:"SystemLogModuleLists"`
		// 系统模块

		SystemModuleLists []*string `json:"SystemModuleLists,omitempty" name:"SystemModuleLists"`
		// 防火墙类型

		TypeLists []*string `json:"TypeLists,omitempty" name:"TypeLists"`
		// vpc列表

		VpcLists []*EdgeInfo `json:"VpcLists,omitempty" name:"VpcLists"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateLogSelectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZeroTrustBlockPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP列表

		ZeroTrustCvmDst *string `json:"ZeroTrustCvmDst,omitempty" name:"ZeroTrustCvmDst"`
		// 端口列表

		ZeroTrustCvmPort *string `json:"ZeroTrustCvmPort,omitempty" name:"ZeroTrustCvmPort"`
		// IP列表

		ZeroTrustSrc *string `json:"ZeroTrustSrc,omitempty" name:"ZeroTrustSrc"`
		// IP列表

		ZeroTrustWebDst *string `json:"ZeroTrustWebDst,omitempty" name:"ZeroTrustWebDst"`
		// 端口列表

		ZeroTrustWebPort *string `json:"ZeroTrustWebPort,omitempty" name:"ZeroTrustWebPort"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZeroTrustBlockPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZeroTrustBlockPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态列表

		Data []*NatSwitchStatusInfo `json:"Data,omitempty" name:"Data"`
		// 错误码列表

		FailData []*SwitchFailInfo `json:"FailData,omitempty" name:"FailData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVpcFwSwitchModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例支持的开关模式：1，同开同关；0，单开单关

		SwitchMode *uint64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryVpcFwSwitchModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVpcFwSwitchModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrendsItem struct {

	// 时间坐标

	Label *string `json:"Label,omitempty" name:"Label"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeCfwExportBindResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值信息

		ExportBindLst []*ExportBindLst `json:"ExportBindLst,omitempty" name:"ExportBindLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwExportBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwExportBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEdgeCFWSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0&nbsp;：不允许使用&nbsp;1:&nbsp;可以使用

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEdgeCFWSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeCFWSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 预操作增加规则列表个数

	AddCnt *int64 `json:"AddCnt,omitempty" name:"AddCnt"`
	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 预操作类型，rollback：回滚；cover_import_in：覆盖入向导入；cover_import_out：覆盖出向导入；其他：add

	OprTypeFrom *string `json:"OprTypeFrom,omitempty" name:"OprTypeFrom"`
}

func (r *DescribeNatRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 开关类型，比如dns代表查询nat防火墙的dns开关

	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
}

func (r *DescribeSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociatedInstanceInfo struct {

	// 关联数据库代理Id

	CdbId *string `json:"CdbId,omitempty" name:"CdbId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 关联安全组数量

	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitempty" name:"SecurityGroupCount"`
	// 关联安全组规则数量

	SecurityGroupRuleCount *uint64 `json:"SecurityGroupRuleCount,omitempty" name:"SecurityGroupRuleCount"`
	// 实例类型，3是cvm实例,4是clb实例,5是eni实例,6是云数据库

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeAnalyzeLogRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询字段名称

	Field *string `json:"Field,omitempty" name:"Field"`
	// ES查询索引:&nbsp;&nbsp;&nbsp;netflow_border_1300448058*

	Index *string `json:"Index,omitempty" name:"Index"`
	// 柱状图间隔

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 返回页高

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回页刻度

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 逻辑表达式

	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询分类1&nbsp;-&nbsp;6

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAnalyzeLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAnalyzeLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest

	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// vpc间防火墙开关边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ip查询条件

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 日志来源&nbsp;move：vpc间防火墙

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeBlockByIpTimesListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockByIpTimesListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwCidrRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DescribeCfwCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 状态值，0:&nbsp;关闭&nbsp;,1:开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0:&nbsp;互联网边界边界防火墙开关，1：vpc防火墙开关

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyItemSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyItemSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUpdateResourceTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态：1，正在同步中，0，同步成功,&nbsp;-1，&nbsp;同步失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 任务状态最新时间：格式：2022-01-01&nbsp;01:01:01

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUpdateResourceTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUpdateResourceTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOASort struct {

	// 排序字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeExportOperatorLogsRequest struct {
	*tchttp.BaseRequest

	// 操作行为

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 结束值

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 防火墙类型

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 'ids',&nbsp;'acl',&nbsp;'login',&nbsp;'switch'

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始值

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeExportOperatorLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportOperatorLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDoaminTopRequest struct {
	*tchttp.BaseRequest

	// 实例Id，不指定则传空

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 时间类型，0：7天，1：24小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNetFlowDoaminTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDoaminTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回树形结构

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码；0为请求成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回消息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 未分类实例数量

		UnResourceNum *int64 `json:"UnResourceNum,omitempty" name:"UnResourceNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserAuthCheckStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserAuthCheckStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserAuthCheckStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupListData struct {

	// 资产分组名称

	AssetGroupNameIn *string `json:"AssetGroupNameIn,omitempty" name:"AssetGroupNameIn"`
	// 资产分组名称

	AssetGroupNameOut *string `json:"AssetGroupNameOut,omitempty" name:"AssetGroupNameOut"`
	// 单/双向下发，0:单向下发，1：双向下发

	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`
	// 生成双向下发规则

	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitempty" name:"BothWayInfo"`
	// 掩码地址，多个以英文逗号分隔

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，0：出站，1：入站，默认1

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否是正常规则，0：正常，1：异常

	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 模板名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 目的端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 内网IP，多个以英文逗号分隔

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口协议类型参数模板名称

	ProtocolPortName *string `json:"ProtocolPortName,omitempty" name:"ProtocolPortName"`
	// 是否使用端口协议模板，0：否，1：是

	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
	// 公网IP，多个以英文逗号分隔

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 端口协议类型参数模板id

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 访问源

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 访问源类型，默认为0，1:&nbsp;VPC,&nbsp;2:&nbsp;SUBNET,&nbsp;3:&nbsp;CVM,&nbsp;4:&nbsp;CLB,&nbsp;5:&nbsp;ENI,&nbsp;6:&nbsp;CDB,&nbsp;7:&nbsp;参数模板,&nbsp;100:&nbsp;资源组

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 是否开关开启，0：未开启，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略,&nbsp;1：阻断，2：放行

	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 访问目的

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 访问目的类型，默认为0，1:&nbsp;VPC,&nbsp;2:&nbsp;SUBNET,&nbsp;3:&nbsp;CVM,&nbsp;4:&nbsp;CLB,&nbsp;5:&nbsp;ENI,&nbsp;6:&nbsp;CDB,&nbsp;7:&nbsp;参数模板,&nbsp;100:资源组

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeVpcFwStatusBarRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcFwStatusBarRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwStatusBarRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwGroupSwitchAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作开关个数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFwGroupSwitchAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwGroupSwitchAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 修改成功与否状态值&nbsp;0：修改成功，非&nbsp;0：修改失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyItemSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyItemSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveActivityVoucherRequest struct {
	*tchttp.BaseRequest
}

func (r *ReceiveActivityVoucherRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveActivityVoucherRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogListsRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;入站，1：出站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束日期

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始日期

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFlowLogListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchAllRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;不自动选择子网
	// 1&nbsp;自动选择子网&nbsp;创建私有连接

	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitempty" name:"AutoChooseSubnet"`
	// 0:&nbsp;全部关闭
	// 1:&nbsp;全部打开
	// 2&nbsp;不操作开关，此次切换模式

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 0&nbsp;全部切换为旁路
	// 1&nbsp;全部切换为串行
	// 2&nbsp;不切换模式，此次操作开关

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

func (r *ModifyEdgeIpSwitchAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertSourceConfig struct {

	// 安全基线

	BaseLine *int64 `json:"BaseLine,omitempty" name:"BaseLine"`
	// 基础防御

	BasicRule *int64 `json:"BasicRule,omitempty" name:"BasicRule"`
	// 网络蜜罐

	HoneyPot *int64 `json:"HoneyPot,omitempty" name:"HoneyPot"`
	// 威胁情报

	ThreatenInfo *int64 `json:"ThreatenInfo,omitempty" name:"ThreatenInfo"`
	// 虚拟补丁

	VirtualPatch *int64 `json:"VirtualPatch,omitempty" name:"VirtualPatch"`
}

type FwDeploy struct {

	// 若为cdc防火墙时填充该id

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 异地灾备&nbsp;1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备

	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`
	// 防火墙部署地域

	DeployRegion *string `json:"DeployRegion,omitempty" name:"DeployRegion"`
	// 带宽，单位：Mbps

	Width *int64 `json:"Width,omitempty" name:"Width"`
	// 主可用区，为空则选择默认可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 备可用区，为空则选择默认可用区

	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`
}

type DeleteVpcFwGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcFwGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest

	// 筛选结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 日志类型标识

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeCommonResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNoticeCommonResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeCommonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSessionConnectionsRequest struct {
	*tchttp.BaseRequest

	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙组Id

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// Vpc间防火墙示例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcSessionConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSessionConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCFWFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态0成功&nbsp;非0失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCFWFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCFWFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyAlertApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户配置信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// DEVLOPMENT:&nbsp;测试环境，&nbsp;PRODUCTION:&nbsp;正式环境

		Env *string `json:"Env,omitempty" name:"Env"`
		// 1:有权限访问,&nbsp;0:无权限访问

		HasAuthority *uint64 `json:"HasAuthority,omitempty" name:"HasAuthority"`
		// 是否为白名单用户

		IsWhiteListUser *int64 `json:"IsWhiteListUser,omitempty" name:"IsWhiteListUser"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwSwitchOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DNS防火墙打开开关数

		DnsOpenSwitchNum *int64 `json:"DnsOpenSwitchNum,omitempty" name:"DnsOpenSwitchNum"`
		// 互联网边界防火墙打开开关数

		EdgeOpenSwitchNum *int64 `json:"EdgeOpenSwitchNum,omitempty" name:"EdgeOpenSwitchNum"`
		// NAT防火墙打开开关数

		NatOpenSwitchNum []*StringValueText `json:"NatOpenSwitchNum,omitempty" name:"NatOpenSwitchNum"`
		// VPC间防火墙打开开关数

		VpcOpenSwitchNum *int64 `json:"VpcOpenSwitchNum,omitempty" name:"VpcOpenSwitchNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwSwitchOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwSwitchOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalAccountAccountVirtualGroupsData struct {

	// 虚拟组Id(只支持32位)

	AccountVirtualGroupId *int64 `json:"AccountVirtualGroupId,omitempty" name:"AccountVirtualGroupId"`
	// 虚拟组名称

	AccountVirtualGroupName *string `json:"AccountVirtualGroupName,omitempty" name:"AccountVirtualGroupName"`
}

type CreateBetaTaskRequest struct {
	*tchttp.BaseRequest

	// 选中规则列表

	AclList []*string `json:"AclList,omitempty" name:"AclList"`
	// 规则类型&nbsp;&nbsp;1border&nbsp;2nat&nbsp;3vpc&nbsp;4sg

	AclType *int64 `json:"AclType,omitempty" name:"AclType"`
	// 动作列表

	TaskActionList []*BetaTaskAction `json:"TaskActionList,omitempty" name:"TaskActionList"`
	// 任务对象信息

	TaskInfo *AutoAssistantUserConf `json:"TaskInfo,omitempty" name:"TaskInfo"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *CreateBetaTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBetaTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchErrorRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSwitchErrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchErrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 时间格式：yyyy-MM-dd&nbsp;HH:mm:ss，IocAction&nbsp;为edit或add时必填，必须大于当前时间且大于StartTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// IP、Domain二选一（注：封禁列表，只能填写IP），不能同时为空

	IOC []*IocListData `json:"IOC,omitempty" name:"IOC"`
	// 可选值：delete（删除）、edit（编辑）、add（添加）&nbsp;&nbsp;其他值无效

	IocAction *string `json:"IocAction,omitempty" name:"IocAction"`
	// 1封禁列表&nbsp;2&nbsp;放通列表

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 时间格式：yyyy-MM-dd&nbsp;HH:mm:ss，IocAction&nbsp;为edit或add时必填

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *ModifyBlockIgnoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNatFwObjRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 模式&nbsp;1：接入模式；0：新增模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 接入模式增加接入的nat网关列表，其中NatGwList和VpcList只能传递一个。

	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`
	// 新增模式增加接入的vpc列表，其中NatGwList和NatgwList只能传递一个。

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *AddNatFwObjRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatFwObjRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest

	// 是否删除全部

	DeleteAll *int64 `json:"DeleteAll,omitempty" name:"DeleteAll"`
	// 封禁：1，放通：100，
	// 主要用于全部删除时区分列表类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则列表

	Rules []*BanAndAllowRuleDel `json:"Rules,omitempty" name:"Rules"`
	// blocklist&nbsp;封禁列表&nbsp;whitelist&nbsp;白名单列表

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DeleteBlockIgnoreRuleNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryNotEmptyRuleListInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQueryNotEmptyRuleListInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryNotEmptyRuleListInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcFwGroupRequest struct {
	*tchttp.BaseRequest

	// 云联网id&nbsp;，适用于云联网模式

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 跨租户管理员模式&nbsp;&nbsp;1管理员&nbsp;2多账号

	CrossUserMode *string `json:"CrossUserMode,omitempty" name:"CrossUserMode"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// auto&nbsp;自动选择防火墙网段
	// 10.10.10.0/24&nbsp;用户输入的防火墙网段

	FwVpcCidr *string `json:"FwVpcCidr,omitempty" name:"FwVpcCidr"`
	// 模式&nbsp;1：CCN云联网模式；0：私有网络模式&nbsp;2:&nbsp;sase&nbsp;模式&nbsp;3：ccn&nbsp;高级模式&nbsp;4:&nbsp;私有网络(跨租户单边模式)

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// VPC防火墙(组)名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 防火墙实例的开关模式
	// 1:&nbsp;单点互通
	// 2:&nbsp;多点互通
	// 3:&nbsp;全互通
	// 4:&nbsp;自定义路由

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
	// 防火墙(组)下的防火墙实例列表

	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitempty" name:"VpcFwInstances"`
}

func (r *CreateVpcFwGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcFwGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPBySGSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlockIPBySGSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPBySGSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDeliveryConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogDeliveryConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDeliveryConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStorageSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDnsResolveIpRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// dns解析地址列表，为空时将使用云默认的ip列表

	ResolveIpLst []*string `json:"ResolveIpLst,omitempty" name:"ResolveIpLst"`
}

func (r *ModifyDnsResolveIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnsResolveIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeIpFlowStat struct {

	// 实例类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 入站均值带宽

	InFlowAvg *int64 `json:"InFlowAvg,omitempty" name:"InFlowAvg"`
	// 入站峰值带宽

	InputMax *int64 `json:"InputMax,omitempty" name:"InputMax"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 出站均值带宽

	OutFlowAvg *int64 `json:"OutFlowAvg,omitempty" name:"OutFlowAvg"`
	// 出站峰值带宽

	OutputMax *int64 `json:"OutputMax,omitempty" name:"OutputMax"`
	// ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 开关状态&nbsp;0:关闭&nbsp;1:打开

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest

	// 地域代码（例：ap-NewYork）,支持云全地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 每页记录条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 列表偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式（asc:升序,desc:降序）

	Order *string `json:"Order,omitempty" name:"Order"`
	// 额外检索条件（JSON字符串）

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 安全组ID

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAssociatedInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssociatedInstanceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEdgeIpFlowListsRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 地域key

	FwRegion *string `json:"FwRegion,omitempty" name:"FwRegion"`
	// 开关流量统计页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开关流量统计页分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序&nbsp;升序&nbsp;asc&nbsp;降序&nbsp;desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeEdgeIpFlowListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeIpFlowListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsInfo struct {

	// 租户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 机器id

	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`
	// Device

	Device *string `json:"Device,omitempty" name:"Device"`
	// Fstype

	Fstype *string `json:"Fstype,omitempty" name:"Fstype"`
	// 防火墙类型

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 机器ip

	Host *string `json:"Host,omitempty" name:"Host"`
	// Mode

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 租户名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// Path

	Path *string `json:"Path,omitempty" name:"Path"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// QueueNumber

	QueueNumber *int64 `json:"QueueNumber,omitempty" name:"QueueNumber"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeOfflineExportTaskRequest struct {
	*tchttp.BaseRequest

	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 任务ID/任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *DescribeOfflineExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOfflineExportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWProductPricesRequest struct {
	*tchttp.BaseRequest

	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeCFWProductPricesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWProductPricesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCommonStatusRequest struct {
	*tchttp.BaseRequest

	// 操作对象唯一key

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 操作动作

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
	// 前后端调用约定

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
}

func (r *ModifyCommonStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCommonStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeCommonNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNoticeCommonNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeCommonNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseWhiteListInfo struct {

	// 接入数据库数量

	AccessDatabaseCount *uint64 `json:"AccessDatabaseCount,omitempty" name:"AccessDatabaseCount"`
	// AccessDomainList

	AccessDomainList []*string `json:"AccessDomainList,omitempty" name:"AccessDomainList"`
	// 登录次数

	LoginTimes *uint64 `json:"LoginTimes,omitempty" name:"LoginTimes"`
	// 接入白名单数量

	WhiteListCount *uint64 `json:"WhiteListCount,omitempty" name:"WhiteListCount"`
}

type RuleImpactInfo struct {

	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 边id

	EdgeID *string `json:"EdgeID,omitempty" name:"EdgeID"`
	// 边名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域name

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 规则列表

	RuleList []*RuleInfoResource `json:"RuleList,omitempty" name:"RuleList"`
	// 规则数

	RuleTotalCount *int64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.domain.com,www.domain.com

	IpString *string `json:"IpString,omitempty" name:"IpString"`
	// IP版本,0&nbsp;IPV4;1&nbsp;IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 协议端口模板，协议类型，4:4层协议，7:7层协议，Type=6时必填

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 1&nbsp;ip模板
	// 5&nbsp;域名模板

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBakRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBakRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBakRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOnceClickScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一键体检用户信息

		Data *UserDetectedInfo `json:"Data,omitempty" name:"Data"`
		// 状态：0正常，非0异常

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOnceClickScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOnceClickScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwJoinInstancesRequest struct {
	*tchttp.BaseRequest

	// VPC防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 是否跨租户&nbsp;1是&nbsp;&nbsp;0不是

	IsPeer *int64 `json:"IsPeer,omitempty" name:"IsPeer"`
	// 每页长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 对等实例id，跨租户需要

	PeerFwInsId *string `json:"PeerFwInsId,omitempty" name:"PeerFwInsId"`
}

func (r *DescribeVpcFwJoinInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwJoinInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDnsRuleItem struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 内部id

	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`
	// 关联告警Id

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 规则序号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：&nbsp;accept：放行&nbsp;drop：拒绝&nbsp;log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 访问源示例：ip地址，192.168.0.2；ip段，192.168.0.0/16

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源类型，可以为&nbsp;ip,net,template,instance,group,tag

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 访问目的示例：域名规则，例如*.domain.com

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 访问目的类型，可以为&nbsp;&nbsp;domain,template

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则对应的唯一id，创建规则时无需填写

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeTiWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关基础信息

		Data []*TiWhiteInfo `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTiWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest

	// 需要编辑的规则数组

	Rules []*CreateRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyAclRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryActivityInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActivityInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;不自动创建私有连接
		// 1&nbsp;自动创建私有连接

		AutoCreateEndpoint *int64 `json:"AutoCreateEndpoint,omitempty" name:"AutoCreateEndpoint"`
		// 0&nbsp;关闭
		// 1&nbsp;打开新增资产自动开启防火墙

		AutoDefence *int64 `json:"AutoDefence,omitempty" name:"AutoDefence"`
		// nat防火墙过载持续时间，单位s，持续超过该时间将bypass；0，不触发bypass；-1，不支持触发bypass；大于0时为具体持续时间

		NatFwOverLoadTime *int64 `json:"NatFwOverLoadTime,omitempty" name:"NatFwOverLoadTime"`
		// nat防火墙管理是否保留云服务器的公网IP类型路由启用状态，0:&nbsp;默认关闭&nbsp;&nbsp;1:&nbsp;启用，将保留公网ip类型路由

		NatHoldPublicIPTypeRoute *int64 `json:"NatHoldPublicIPTypeRoute,omitempty" name:"NatHoldPublicIPTypeRoute"`
		// nat防火墙管理网络异常bypass，1，开启；0，关闭

		NatManageNetworkDownBypass *int64 `json:"NatManageNetworkDownBypass,omitempty" name:"NatManageNetworkDownBypass"`
		// 0&nbsp;旁路
		// 1&nbsp;串行

		SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
		// vpc间防火墙过载持续时间，单位s，持续超过该时间将bypass；0，不触发bypass；-1，不支持触发bypass；大于0时为具体持续时间

		VpcFwOverLoadTime *int64 `json:"VpcFwOverLoadTime,omitempty" name:"VpcFwOverLoadTime"`
		// vpc间防火墙管理网络异常bypass，1，开启；0，关闭

		VpcManageNetworkDownBypass *int64 `json:"VpcManageNetworkDownBypass,omitempty" name:"VpcManageNetworkDownBypass"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量控制列表数据

		Data []*TcRule `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetflowRuleStatusAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的开关个数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetflowRuleStatusAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetflowRuleStatusAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderProductInfo struct {

	// 地址模板数量

	AddressTemplateNum *int64 `json:"AddressTemplateNum,omitempty" name:"AddressTemplateNum"`
	// 南北向通用带宽，单位Mbps

	EdgeBandwidth *int64 `json:"EdgeBandwidth,omitempty" name:"EdgeBandwidth"`
	// 扩展通用规则数

	ExtendAclRuleNum *int64 `json:"ExtendAclRuleNum,omitempty" name:"ExtendAclRuleNum"`
	// 通用扩展实例，单位个

	FirewallInsNum *int64 `json:"FirewallInsNum,omitempty" name:"FirewallInsNum"`
	// 蜜罐数

	HoneypotNum *int64 `json:"HoneypotNum,omitempty" name:"HoneypotNum"`
	// 是否开启日志分析

	LogAnalysis *int64 `json:"LogAnalysis,omitempty" name:"LogAnalysis"`
	// 日志存储容量，单位GB

	LogStorage *int64 `json:"LogStorage,omitempty" name:"LogStorage"`
	// 套餐版本

	VersionType *int64 `json:"VersionType,omitempty" name:"VersionType"`
	// 套餐版本名称

	VersionTypeName *string `json:"VersionTypeName,omitempty" name:"VersionTypeName"`
	// 零信任远程运维扩展包数

	ZeroTrustOmPacketNum *int64 `json:"ZeroTrustOmPacketNum,omitempty" name:"ZeroTrustOmPacketNum"`
	// 零信任远程运维版本，1基础版，2高级运维版

	ZeroTrustOmVersion *int64 `json:"ZeroTrustOmVersion,omitempty" name:"ZeroTrustOmVersion"`
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest

	// Nat所在地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 0：&nbsp;出向，1：入向

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 状态值，1：锁表，2：解锁表

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyTableStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTableStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelatedHoneyPot struct {

	// 蜜罐服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 蜜罐服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type DescribeProbeTaskRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 条件过滤

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeProbeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回修改成功的规则ID列表

		RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAcRuleSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveNatFwObjResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveNatFwObjResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatFwObjResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseInstanceInfoDetail struct {

	// 数据库实例详情

	DatabaseInstanceList []*DatabaseInstanceInfo `json:"DatabaseInstanceList,omitempty" name:"DatabaseInstanceList"`
	// 地域总数

	Region *string `json:"Region,omitempty" name:"Region"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CreateApplyTrialRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateApplyTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplyTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsAsyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Data *LogAsyncRespond `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsAsyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwGroupInstanceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙(组)个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 防火墙(组)

		VpcFwGroupLst []*VpcFwGroupInfo `json:"VpcFwGroupLst,omitempty" name:"VpcFwGroupLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwGroupInstanceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwGroupInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作的开关个数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwSwitchAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwSwitchAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwGroupIdName struct {

	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
}

type DescribeOperateLogSelectRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 操作日志类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 筛选值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeOperateLogSelectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogSelectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLogEdgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回可日志查询得VPC防火墙开关列表

		EdgeLst []*EdgeInfo `json:"EdgeLst,omitempty" name:"EdgeLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcLogEdgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLogEdgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcZoneData struct {

	// vpc节点地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeRouteBackupLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份路由列表

		Data []*RouteBackupInfo `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteBackupLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteBackupLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatExistRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatExistRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatExistRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest

	// 查询示例ID

	AssetID *string `json:"AssetID,omitempty" name:"AssetID"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeUnHandleEventTabListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnHandleEventTabListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDynamicCheckRequest struct {
	*tchttp.BaseRequest

	// 取消的动态id集合

	ReadId []*string `json:"ReadId,omitempty" name:"ReadId"`
}

func (r *ModifyDynamicCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDynamicCheckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 模式&nbsp;1：接入模式；0：新增模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。

	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`
	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *ModifyNatFwReSelectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwReSelectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeInfoRsp struct {

	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 部署方式：EIP、CLB

	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`
	// 命中次数

	HitTimes *int64 `json:"HitTimes,omitempty" name:"HitTimes"`
	// 转发到蜜罐数量

	HoneyPotCount *int64 `json:"HoneyPotCount,omitempty" name:"HoneyPotCount"`
	// 部署实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 部署实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 探针Id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 探针名称

	ProbeName *string `json:"ProbeName,omitempty" name:"ProbeName"`
	// 转发器Id

	ProxyId *string `json:"ProxyId,omitempty" name:"ProxyId"`
	// 转发器名称

	ProxyName *string `json:"ProxyName,omitempty" name:"ProxyName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 服务实例状态&nbsp;0&nbsp;完成&nbsp;1未创建&nbsp;2创建中&nbsp;3正删除&nbsp;4&nbsp;开关转变中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 服务开关&nbsp;0关闭&nbsp;1打开

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 所属私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 所属私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeIpsLogsRequest struct {
	*tchttp.BaseRequest

	// 方向，0：入站，1：出站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束日期

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 威胁类型名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 危险等级,&nbsp;提示->信息，中危->低，高危->高，极高危->极高

	Level *string `json:"Level,omitempty" name:"Level"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始日期

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 0:&nbsp;观察，1：阻断，2：放行

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeIpsLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageCenterSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMessageCenterSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMessageCenterSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnHandleEventTabListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户伪攻击链未处置事件

		Data *UnHandleEvent `json:"Data,omitempty" name:"Data"`
		// 错误码，0成功&nbsp;非0错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;success成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnHandleEventTabListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnHandleEventTabListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIPSProtectModeRequest struct {
	*tchttp.BaseRequest

	// 防火墙类型："border",&nbsp;"nat",&nbsp;"vpc"

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 0&nbsp;观察&nbsp;1&nbsp;拦截&nbsp;2&nbsp;严格&nbsp;3&nbsp;关闭

	IdpsAction *int64 `json:"IdpsAction,omitempty" name:"IdpsAction"`
	// 0&nbsp;局部&nbsp;1&nbsp;全局

	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 变更保护详情

	ProtectDetail *ProtectInfo `json:"ProtectDetail,omitempty" name:"ProtectDetail"`
}

func (r *ModifyIPSProtectModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIPSProtectModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSelectAssetGroupRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 组id&nbsp;&nbsp;&nbsp;未分组资产id传&nbsp;no-resource

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 顺序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 类型&nbsp;1资产组2..资产标签3.子网

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 筛选零信任远程运维已接入资产

	RemoteOMFilter *int64 `json:"RemoteOMFilter,omitempty" name:"RemoteOMFilter"`
	// 查询

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 查询类型&nbsp;all&nbsp;查看自己和子组&nbsp;own&nbsp;查看自己

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeSelectAssetGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSelectAssetGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDetailRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否加载其他用户的跨租户vpc

	NeedRemoteVpc *int64 `json:"NeedRemoteVpc,omitempty" name:"NeedRemoteVpc"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 地域列表

	RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// vpc&nbsp;模式

	VpcMode *int64 `json:"VpcMode,omitempty" name:"VpcMode"`
}

func (r *DescribeVpcDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActivityInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 活动信息

		ActivityInfo *ActivityInfo `json:"ActivityInfo,omitempty" name:"ActivityInfo"`
		// 代金券领取，0未领取，1已领取

		AlreadyReceivedVoucher *uint64 `json:"AlreadyReceivedVoucher,omitempty" name:"AlreadyReceivedVoucher"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActivityInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActivityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainData struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type GetEventAssetStruct struct {

	// 无

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 无

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 无

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeNatRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// NAT地域&nbsp;&nbsp;这个是必填项，填入相关的英文，'ap-NY-fsi':&nbsp;'纽约金融'

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向，0：出站，1：入站&nbsp;默认值：0

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeNatRuleOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapTaskRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribePcapTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePcapTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwVpcDnsSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwVpcDnsSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwVpcDnsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleLogDetail struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略

	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
}

type IsolateTable struct {

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 资产ID，唯一key

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资产类型说明

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 1&nbsp;入站&nbsp;2&nbsp;出战&nbsp;4&nbsp;内网

	IsolateType *int64 `json:"IsolateType,omitempty" name:"IsolateType"`
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 资产地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 安全组规则id

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// vpcID，唯一key

	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 1&nbsp;手动白名单&nbsp;2&nbsp;身份认证

	YwMode *int64 `json:"YwMode,omitempty" name:"YwMode"`
}

type DescribeEsLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果或redis缓存key

		Data *string `json:"Data,omitempty" name:"Data"`
		// 值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEsLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertCenterOmitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码：
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息：
		// success&nbsp;成功
		// 其他

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 处置状态码：
		// 0&nbsp;&nbsp;处置成功
		// -1&nbsp;通用错误，不用处理
		// -3&nbsp;表示重复，需重新刷新列表
		// 其他

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertCenterOmitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterOmitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchCommonResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关基础信息

		Data *SwitchCommon `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchCommonResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchCommonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeThreatInfoNewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果合计

		Data *TiNews `json:"Data,omitempty" name:"Data"`
		// 0成功&nbsp;非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeThreatInfoNewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeThreatInfoNewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdsWhiteInfo struct {

	// 目的IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 白名单生效防火墙范围：&nbsp;1&nbsp;边界防火墙&nbsp;2&nbsp;nat防火墙&nbsp;4&nbsp;vpc防火墙&nbsp;7&nbsp;=&nbsp;1+2+4&nbsp;所有防火墙

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// 白名单唯一ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 入侵防御规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 规则类型

	WhiteRuleType *string `json:"WhiteRuleType,omitempty" name:"WhiteRuleType"`
}

type ProbeTaskDetail struct {

	// 防火墙实例

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 防火墙实例名

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// 防火墙类型，ew:vpc防火墙，nat:nat防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 最近拨测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 0:&nbsp;关&nbsp;1：打开

	LeftInnerProbe *uint64 `json:"LeftInnerProbe,omitempty" name:"LeftInnerProbe"`
	// 1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	LeftInnerProbeResult *uint64 `json:"LeftInnerProbeResult,omitempty" name:"LeftInnerProbeResult"`
	// 左侧内网拨测结果说明

	LeftInnerProbeStatus *string `json:"LeftInnerProbeStatus,omitempty" name:"LeftInnerProbeStatus"`
	// 左侧ip

	LeftIp *string `json:"LeftIp,omitempty" name:"LeftIp"`
	// 0:&nbsp;关&nbsp;1：打开

	LeftProbe *uint64 `json:"LeftProbe,omitempty" name:"LeftProbe"`
	// 1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	LeftProbeResult *uint64 `json:"LeftProbeResult,omitempty" name:"LeftProbeResult"`
	// 左侧拨测结果说明

	LeftProbeStatus *string `json:"LeftProbeStatus,omitempty" name:"LeftProbeStatus"`
	// 任务编辑时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 失败重试次数

	RetryTimes *uint64 `json:"RetryTimes,omitempty" name:"RetryTimes"`
	// 0:&nbsp;关&nbsp;1：打开

	RightInternetProbe *uint64 `json:"RightInternetProbe,omitempty" name:"RightInternetProbe"`
	// 1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	RightInternetProbeResult *uint64 `json:"RightInternetProbeResult,omitempty" name:"RightInternetProbeResult"`
	// 右侧互联网拨测结果说明

	RightInternetProbeStatus *string `json:"RightInternetProbeStatus,omitempty" name:"RightInternetProbeStatus"`
	// 右侧ip

	RightIp *string `json:"RightIp,omitempty" name:"RightIp"`
	// 0:&nbsp;关&nbsp;1：打开

	RightProbe *uint64 `json:"RightProbe,omitempty" name:"RightProbe"`
	// 1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	RightProbeResult *uint64 `json:"RightProbeResult,omitempty" name:"RightProbeResult"`
	// 右侧拨测结果说明

	RightProbeStatus *string `json:"RightProbeStatus,omitempty" name:"RightProbeStatus"`
	// 拨测动作：0：无动作，1：拨测中

	Running *uint64 `json:"Running,omitempty" name:"Running"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务拨测结果，1：未配置，2：待执行，3：全部健康，4、部分异常，5：全部异常

	TaskProbeResult *uint64 `json:"TaskProbeResult,omitempty" name:"TaskProbeResult"`
	// 时间格式or周期格式如month|3|00:00:00

	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`
	// 任务类型，1：临时任务，2：周期任务

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type ModifyAclApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyAclApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicy struct {

	// 动作：ACCEPT&nbsp;或&nbsp;DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// IP地址ID或者ID地址组ID。

	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitempty" name:"AddressTemplate"`
	// 网段或IP(互斥)。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 网段或IPv6(互斥)。

	IpCidrBlock *string `json:"IpCidrBlock,omitempty" name:"IpCidrBlock"`
	// 安全组规则描述。

	PolicyDescription *string `json:"PolicyDescription,omitempty" name:"PolicyDescription"`
	// 安全组规则索引号

	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
	// 规则类型&nbsp;CREATE&nbsp;&nbsp;DEL&nbsp;NORMAL

	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 端口(all,&nbsp;离散port,&nbsp;&nbsp;range)。

	Port *string `json:"Port,omitempty" name:"Port"`
	// &nbsp;协议,&nbsp;取值:&nbsp;TCP,UDP,ICMP,ICMPv6,ALL。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 安全组实例ID，例如：sg-ohuuioma。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。

	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`
}

type SetNatProbeEipRequest struct {
	*tchttp.BaseRequest

	// eip集合

	EipList []*string `json:"EipList,omitempty" name:"EipList"`
	// 实例id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 操作类型&nbsp;check&nbsp;验证联通性/modify修改&nbsp;

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SetNatProbeEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatProbeEipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeRange struct {

	// 对端网络实例cidr，多个以逗号分隔

	DstCidr *string `json:"DstCidr,omitempty" name:"DstCidr"`
	// 对端网络实例Id

	DstId *string `json:"DstId,omitempty" name:"DstId"`
	// 对端网络实例名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 对端网络实例所属地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 规则生效的范围id，是在哪对vpc之间还是针对所有vpc间生效

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// EdgeId对应的名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 本端网络实例cidr，多个以逗号分隔

	SrcCidr *string `json:"SrcCidr,omitempty" name:"SrcCidr"`
	// 本端网络实例Id

	SrcId *string `json:"SrcId,omitempty" name:"SrcId"`
	// 本端网络实例名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 本端网络实例所属地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
}

type ZeroTrustRegionItem struct {

	// cvm数量

	CvmCount *uint64 `json:"CvmCount,omitempty" name:"CvmCount"`
	// db数量

	DBCount *uint64 `json:"DBCount,omitempty" name:"DBCount"`
	// EIP数组

	EIP []*string `json:"EIP,omitempty" name:"EIP"`
	// 开关状态0，1

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 是否为低版本免费版
	// 0&nbsp;不可以创建EIP
	// 1&nbsp;可以创建EIP&nbsp;但只可以创建1个
	// 2&nbsp;高级版&nbsp;可以创建两个EIP

	OldUser *int64 `json:"OldUser,omitempty" name:"OldUser"`
	// 地域标识

	Region *string `json:"Region,omitempty" name:"Region"`
	// 0:完成,&nbsp;1:开启中，2：关闭中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// vpc数量

	VpcCount *uint64 `json:"VpcCount,omitempty" name:"VpcCount"`
	// web数量

	WebCount *uint64 `json:"WebCount,omitempty" name:"WebCount"`
}

type AddNatFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 限速的类型，可以是具体的IP也可以是CIDR，比如192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 需要添加限速规则的防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
	// 入向限制速度，单位是kbps

	InLimit *uint64 `json:"InLimit,omitempty" name:"InLimit"`
	// 限速类型
	// IP限速：IP
	// 端口限速：PORT
	// 当前仅支持IP限速

	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`
	// 出向限制速度，单位是kbps

	OutLimit *uint64 `json:"OutLimit,omitempty" name:"OutLimit"`
	// 流量控制策略的端口。取值：
	// -1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

func (r *AddNatFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddrTmplSelectListRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAddrTmplSelectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddrTmplSelectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeIpInfo struct {

	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 域名化CLB的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 0:&nbsp;开启开关时提示要创建私有连接。
	// 1:&nbsp;关闭该开关是提示删除私有连接。
	// 如果大于&nbsp;1:&nbsp;关闭开关&nbsp;、开启开关不需提示创建删除私有连接。

	EndpointBindEipNum *int64 `json:"EndpointBindEipNum,omitempty" name:"EndpointBindEipNum"`
	// 私有连接ID

	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
	// 私有连接IP

	EndpointIp *string `json:"EndpointIp,omitempty" name:"EndpointIp"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网IP

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
	// 0:&nbsp;不是公网CLB&nbsp;可以开启串行开关
	// 1:&nbsp;是公网CLB&nbsp;不可以开启串行开关
	//

	IsPublicClb *int64 `json:"IsPublicClb,omitempty" name:"IsPublicClb"`
	// 是否为region&nbsp;eip
	// 0&nbsp;不为region&nbsp;eip，不能选择串行
	// 1&nbsp;为region&nbsp;eip&nbsp;可以选择串行

	IsRegionEip *int64 `json:"IsRegionEip,omitempty" name:"IsRegionEip"`
	// 0:&nbsp;该地域暂未支持串行
	// 1:&nbsp;该用户未在该地域配置串行带宽
	// 2:&nbsp;该用户已在该地域配置串行带宽，可以开启串行开关

	IsSerialRegion *int64 `json:"IsSerialRegion,omitempty" name:"IsSerialRegion"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// IP超量状态

	OverUsedStatus *int64 `json:"OverUsedStatus,omitempty" name:"OverUsedStatus"`
	// 风险端口数

	PortRiskCount *int64 `json:"PortRiskCount,omitempty" name:"PortRiskCount"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 公网&nbsp;IP&nbsp;类型

	PublicIpType *int64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 扫描深度

	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`
	// 扫描状态

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 开关状态
	// 0&nbsp;:&nbsp;关闭
	// 1&nbsp;:&nbsp;开启
	// 2&nbsp;:&nbsp;开启中
	// 3&nbsp;:&nbsp;关闭中
	// 4&nbsp;:&nbsp;异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0&nbsp;:&nbsp;旁路
	// 1&nbsp;:&nbsp;串行
	// 2&nbsp;:&nbsp;正在模式切换

	SwitchMode *uint64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
	// 开关权重

	SwitchWeight *int64 `json:"SwitchWeight,omitempty" name:"SwitchWeight"`
	// EIP&nbsp;所关联的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type SecretKeyNames struct {

	// 密钥文件名

	SecretKeyFileName *string `json:"SecretKeyFileName,omitempty" name:"SecretKeyFileName"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeExportFlowLogsRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;入站，1：出站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束日期

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始日期

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeExportFlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportFlowLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeThreatInfoNewsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeThreatInfoNewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeThreatInfoNewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAcRuleRequest struct {
	*tchttp.BaseRequest

	// 规则的uuid，可通过查询规则列表获取

	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *RemoveAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCdcIdsRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeCdcIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCdcIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsCountAsyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询状态，0查询已完成，1查询进行中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsCountAsyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsCountAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeWarnTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警趋势数据

		Data []*AccessControlOverviewData `json:"Data,omitempty" name:"Data"`
		// 告警统计数据

		NumData *EventNumData `json:"NumData,omitempty" name:"NumData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeWarnTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeWarnTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatSequenceRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatSequenceRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntrusionDefenseCustomRule struct {

	// 备注信息，长度不能超过200字符

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 目的IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 规则结束时间，格式：2023-12-30&nbsp;15:04:05，必须大于当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 入侵防御规则id

	IdsRuleId *string `json:"IdsRuleId,omitempty" name:"IdsRuleId"`
	// 自定义策略名称

	IdsRuleName *string `json:"IdsRuleName,omitempty" name:"IdsRuleName"`
	// 源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 规则开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type ModifyRuleItem struct {

	// 企业安全组规则ID

	RuleUuid *uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

type StaticInfo struct {

	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 资产id

	InsID *string `json:"InsID,omitempty" name:"InsID"`
	// 资产名称

	InsName *string `json:"InsName,omitempty" name:"InsName"`
	// ip信息

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 数

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

type DescribeAllRegionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置信息

		Data []*RegionConfig `json:"Data,omitempty" name:"Data"`
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllRegionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRegionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportIpsLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出地址

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 状态码&nbsp;0成功&nbsp;非0失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportIpsLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportIpsLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockByIpTimesListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*IpStatic `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockByIpTimesListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockByIpTimesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcRuleData struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
}

type DeleteNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupTableStatusRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifySecurityGroupTableStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupTableStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGuideOneClickScanRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateGuideOneClickScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGuideOneClickScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetOutOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data *NetflowCenterOutOverview `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternetOutOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetOutOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnWithRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开白路由表的云联网CCN列表

		Data []*CcnInstance `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnWithRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnWithRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMessageCenterSwitchRequest struct {
	*tchttp.BaseRequest

	// 开关，0：关闭，1：开启

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyMessageCenterSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMessageCenterSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHoneyPotAttackerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Data []*HoneyPotAttacker `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHoneyPotAttackerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHoneyPotAttackerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteNextHopLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份路由的下一跳列表

		Data []*BackupNextHop `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteNextHopLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteNextHopLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebCosUrlRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWebCosUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebCosUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsList struct {

	// 来源或目标网段

	NetWorkID *string `json:"NetWorkID,omitempty" name:"NetWorkID"`
	// //&nbsp;TypeVPC&nbsp;vpc实例&nbsp;TypeVPC&nbsp;int&nbsp;=&nbsp;1&nbsp;//&nbsp;TypeSubnet&nbsp;subnet实例&nbsp;TypeSubnet&nbsp;int&nbsp;=&nbsp;2&nbsp;//&nbsp;TypeCVM&nbsp;CVM实例&nbsp;TypeCVM&nbsp;int&nbsp;=&nbsp;3&nbsp;//&nbsp;TypeCLB&nbsp;CLB实例&nbsp;TypeCLB&nbsp;int&nbsp;=&nbsp;4&nbsp;//&nbsp;TypeENI&nbsp;eni实例&nbsp;TypeENI&nbsp;int&nbsp;=&nbsp;5&nbsp;//&nbsp;TypeCDB&nbsp;cdb实例&nbsp;TypeCDB&nbsp;int&nbsp;=&nbsp;6&nbsp;//&nbsp;TypeRedis&nbsp;Redis实例&nbsp;TypeRedis&nbsp;int&nbsp;=&nbsp;6&nbsp;//&nbsp;TypeCDBMySQL&nbsp;mysql&nbsp;TypeCDBMySQL&nbsp;int&nbsp;=&nbsp;61&nbsp;//&nbsp;TypeCDBREDIS&nbsp;redis&nbsp;TypeCDBREDIS&nbsp;int&nbsp;=&nbsp;62&nbsp;//&nbsp;TypeALL&nbsp;所有类型&nbsp;TypeALL&nbsp;=&nbsp;-1

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type ProtectInfo struct {

	// 告警IP列表/vpc-edge列表

	Alert []*string `json:"Alert,omitempty" name:"Alert"`
	// 告警IP列表/vpc-edge列表

	Block []*string `json:"Block,omitempty" name:"Block"`
	// 告警IP列表/vpc-edge列表

	Close []*string `json:"Close,omitempty" name:"Close"`
	// 告警IP列表/vpc-edge列表

	Strict []*string `json:"Strict,omitempty" name:"Strict"`
}

type DescribeAlertFeedBackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 反馈列表

		Data []*AlterFeedBackShow `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertFeedBackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertFeedBackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DbInstance struct {

	// 总数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 显示名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RemoteAccessInfo struct {

	// 远程运维Dnat规则

	Dnat []*DnatInfo `json:"Dnat,omitempty" name:"Dnat"`
	// 远程运维域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名上次修改时间戳，如果为0则未修改

	LastModifiedTime *uint64 `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`
	// Region,如ap-NewYork

	Region *string `json:"Region,omitempty" name:"Region"`
	// 远程运维开关状态，1代表开启，0代表关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest

	// ChooseType为1，查询已经分组的资产；ChooseType不为1查询没有分组的资产

	ChooseType *string `json:"ChooseType,omitempty" name:"ChooseType"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 资产类型&nbsp;1公网&nbsp;2内网

	InsType *string `json:"InsType,omitempty" name:"InsType"`
	// 查询单页的最大值；eg：10；则最多返回10条结果

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询结果的偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeSourceAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleScopesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatRuleScopesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleScopesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网间访问控制列表数据

		Data []*VpcRuleItem `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChooseResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyChooseResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChooseResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIsolateTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIsolateTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIsolateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPublicIPSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchErrorIgnoreResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySwitchErrorIgnoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchErrorIgnoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclRuleExportStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件Url地址，导出状态

		Data *AclRulePathStatus `json:"Data,omitempty" name:"Data"`
		// Code&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// Msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclRuleExportStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclRuleExportStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupAddressIpListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		IpList []*string `json:"IpList,omitempty" name:"IpList"`
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupAddressIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupAddressIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwViewStatRequest struct {
	*tchttp.BaseRequest

	// 是否跨租户&nbsp;1是&nbsp;0不是

	IsPeer *int64 `json:"IsPeer,omitempty" name:"IsPeer"`
}

func (r *DescribeVpcFwViewStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwViewStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClsDeliverSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值&nbsp;0：成功，非&nbsp;0：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClsDeliverSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClsDeliverSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProbeTaskRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// ew：vpc防火墙，nat：nat防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 访问内网,1-开，0-关

	LeftInnerProbe *uint64 `json:"LeftInnerProbe,omitempty" name:"LeftInnerProbe"`
	// 左侧IP

	LeftIp *string `json:"LeftIp,omitempty" name:"LeftIp"`
	// 左侧访问防火墙,1-开，0-关

	LeftProbe *uint64 `json:"LeftProbe,omitempty" name:"LeftProbe"`
	// 失败重试次数

	RetryTimes *uint64 `json:"RetryTimes,omitempty" name:"RetryTimes"`
	// 访问互联网,1-开，0-关

	RightInternetProbe *uint64 `json:"RightInternetProbe,omitempty" name:"RightInternetProbe"`
	// 右侧IP

	RightIp *string `json:"RightIp,omitempty" name:"RightIp"`
	// 右侧访问防火墙,1-开，0-关

	RightProbe *uint64 `json:"RightProbe,omitempty" name:"RightProbe"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 时间格式or周期格式如month|3|00:00:00

	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`
	// 1：临时任务，2：周期任务

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *ModifyProbeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProbeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserAuthCheckStatusRequest struct {
	*tchttp.BaseRequest

	// 开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// autoScan自动同步，非autoScan检查授权

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyUserAuthCheckStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserAuthCheckStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainInsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Data []*InstanceDomainInfo `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetFlowDomainInsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainInsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNotifyTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNotifyTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNotifyTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShowBakRuleListRequest struct {
	*tchttp.BaseRequest

	// 实例信息

	InsMsg *string `json:"InsMsg,omitempty" name:"InsMsg"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 备份规则列表名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 模糊搜索

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeShowBakRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShowBakRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwUpdateStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfwUpdateStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwUpdateStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountGroupsPageResp struct {

	// 账户分响应对象集合

	Items []*DescribeAccountGroupsData `json:"Items,omitempty" name:"Items"`
	// 分页公共对象

	Page *Paging `json:"Page,omitempty" name:"Page"`
}

type DescribeProbeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		AreaLists []*SelectInfo `json:"AreaLists,omitempty" name:"AreaLists"`
		// 暴露探针列表数据

		Data []*ProbeInfoRsp `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProbeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAclEdgeRangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网间访问控制规则的生效范围列表

		EdgeRanges []*EdgeRange `json:"EdgeRanges,omitempty" name:"EdgeRanges"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcAclEdgeRangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAclEdgeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSyncAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：同步成功，1：资产更新中，2：后台同步调用失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunSyncAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ip流量控制数据

		Data *TcRule `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFlowCenterLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*VpcFlowCenterData `json:"Data,omitempty" name:"Data"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFlowCenterLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFlowCenterLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthUserInfo struct {

	// 正常日期：时间戳
	// 永久过期：-1

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 备注

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 1

	YwId *int64 `json:"YwId,omitempty" name:"YwId"`
}

type DeleteBlockIgnoreRuleAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作名单个数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 返回状态码&nbsp;0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlockIgnoreRuleAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLogInfo struct {

	// 封禁列表

	BanNum *int64 `json:"BanNum,omitempty" name:"BanNum"`
	// 暴力破解

	BruteForceNum *int64 `json:"BruteForceNum,omitempty" name:"BruteForceNum"`
	// 待处置告警

	HandleNum *int64 `json:"HandleNum,omitempty" name:"HandleNum"`
	// 网络探测

	NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`
	// 失陷主机

	OutNum *int64 `json:"OutNum,omitempty" name:"OutNum"`
	// 漏洞攻击

	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`
}

type CreateOperateLogsRequest struct {
	*tchttp.BaseRequest

	// 地址模板名称

	AddrName *string `json:"AddrName,omitempty" name:"AddrName"`
	// NAT地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 日志投递操作类型

	DeliveryType *string `json:"DeliveryType,omitempty" name:"DeliveryType"`
	// 边Id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 零信任操作类型

	RemoteType *string `json:"RemoteType,omitempty" name:"RemoteType"`
	// 操作行为

	ReportAction *string `json:"ReportAction,omitempty" name:"ReportAction"`
	// 资产id

	ReportAssetId *string `json:"ReportAssetId,omitempty" name:"ReportAssetId"`
	// 资产地域

	ReportAssetRegion *string `json:"ReportAssetRegion,omitempty" name:"ReportAssetRegion"`
	// 安全基线类型

	ReportBaseType *string `json:"ReportBaseType,omitempty" name:"ReportBaseType"`
	// 描述

	ReportDetail *string `json:"ReportDetail,omitempty" name:"ReportDetail"`
	// 0,1&nbsp;防火墙类型

	ReportFwType *string `json:"ReportFwType,omitempty" name:"ReportFwType"`
	// 多余的信息

	ReportInfo *string `json:"ReportInfo,omitempty" name:"ReportInfo"`
	// 严重级别，严重/高危/中危/低危/提示

	ReportLevel *string `json:"ReportLevel,omitempty" name:"ReportLevel"`
	// 操作人

	ReportOperator *string `json:"ReportOperator,omitempty" name:"ReportOperator"`
	// 上报规则名称

	ReportRule *string `json:"ReportRule,omitempty" name:"ReportRule"`
	// ['ids',&nbsp;'acl',&nbsp;'login',&nbsp;'switch']

	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`
}

func (r *CreateOperateLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddrTemplateListRequest struct {
	*tchttp.BaseRequest

	// 排序字段，取值&nbsp;'UpdateTime'&nbsp;|&nbsp;'RulesNum'

	By *string `json:"By,omitempty" name:"By"`
	// 条数，分页用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，分页用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，取值&nbsp;'asc'|'desc'

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeAddrTemplateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddrTemplateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTreatInfoStatusRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;关闭，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyTreatInfoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTreatInfoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAclRuleRequest struct {
	*tchttp.BaseRequest

	// 添加规则的来源，一般不需要使用，值insert_rule&nbsp;表示插入指定位置的规则；值batch_import&nbsp;表示批量导入规则；为空时表示添加规则

	From *string `json:"From,omitempty" name:"From"`
	// 需要添加的访问控制规则列表

	Rules []*CreateRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddAclRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAclRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpsRuleListNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 类别

		Category []*string `json:"Category,omitempty" name:"Category"`
		// 数据

		Data []*IpsRuleDetailNew `json:"Data,omitempty" name:"Data"`
		// code&nbsp;0成功其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// code&nbsp;描述

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpsRuleListNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsRuleListNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回多余的信息

		Info *string `json:"Info,omitempty" name:"Info"`
		// 状态值&nbsp;0成功&nbsp;其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIsolateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果合计

		Data []*IsolateTable `json:"Data,omitempty" name:"Data"`
		// 0&nbsp;返回成功&nbsp;非0返回失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;返回成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIsolateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIsolateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleScopesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生效范围列表

		ScopeItems []*RuleScopeInfo `json:"ScopeItems,omitempty" name:"ScopeItems"`
		// 可选生效范围总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleScopesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleScopesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveNatFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 流量控制规则的Id，可通过查询流量控制规则列表获取

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RemoveNatFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportCredentialRequest struct {
	*tchttp.BaseRequest

	// 导入类型

	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`
}

func (r *DescribeImportCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZeroTrustEffectiveScopeItem struct {

	// 生效范围资产或人员Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型，user人员，group组织，ins服务器资产

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ZoneConfig struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区中文

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// nat&nbsp;&nbsp;地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// vpc&nbsp;开关id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// idlist

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteAllAccessControlRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllAccessControlRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeFlowApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest

	// 获取实例列表过滤字段

	Filter []*NatFwFilter `json:"Filter,omitempty" name:"Filter"`
	// 每页长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNatFwInstancesInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstancesInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BorderRuleDetail struct {

	// 城市地域码

	City *int64 `json:"City,omitempty" name:"City"`
	// 云厂商代码

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 国家地域码

	Country *int64 `json:"Country,omitempty" name:"Country"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，取值：0，出站；1，入站，vpc间规则时无方向默认为0

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 访问目的名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 访问目的类型，1是ip，2是url，3是http域名

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 是否为地域规则

	IsRegion *int64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 规则顺序编号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 协议端口组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 协议端口组名称

	ParamTemplateName *string `json:"ParamTemplateName,omitempty" name:"ParamTemplateName"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：
	// TCP
	// UDP
	// ICMP
	// ANY
	// HTTP
	// HTTPS
	// HTTP/HTTPS
	// SMTP
	// SMTPS
	// SMTP/SMTPS
	// FTP
	// DNS
	// TLS/SSL

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 生效范围:0,全局;1,串行;2,旁路

	Scope *int64 `json:"Scope,omitempty" name:"Scope"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 访问源名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 源类型，1是ip，2是url，3是http域名

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 规则被删除：1，已删除；0，未删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// 0：观察
	// 1：拒绝
	// 2：放行

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.domain.com

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
}

type StringValueTextEx struct {

	// 保留字段

	Reserved *string `json:"Reserved,omitempty" name:"Reserved"`
	// 描述

	Text *string `json:"Text,omitempty" name:"Text"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySerialRegionRequest struct {
	*tchttp.BaseRequest

	// 旁路带宽值修改
	// 可以不填，不填即不修改

	BypassWidth *int64 `json:"BypassWidth,omitempty" name:"BypassWidth"`
	// 旁路防火墙弹性带宽上限

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 旁路防火墙弹性按量计费开关

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// 串行防火墙弹性按量计费开关

	SerialElasticSwitch *int64 `json:"SerialElasticSwitch,omitempty" name:"SerialElasticSwitch"`
	// 串行地域

	SerialRegionLst []*SerialRegionInfo `json:"SerialRegionLst,omitempty" name:"SerialRegionLst"`
	// 按量计费标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifySerialRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySerialRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例数组

		NatinsLst []*NatFwInstance `json:"NatinsLst,omitempty" name:"NatinsLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefenseSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：修改成功，非0：修改成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：修改成功，fail：修改成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDefenseSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDefenseSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwIpStatRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙组ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙的实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeVpcFwIpStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwIpStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackListQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlackListQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlackListQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwViewStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json&nbsp;数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwViewStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwViewStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwFirstConfigRequest struct {
	*tchttp.BaseRequest

	// 页面来源：1、开关页

	FromPage *string `json:"FromPage,omitempty" name:"FromPage"`
}

func (r *DescribeFwFirstConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwFirstConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveVpcFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 流量控制规则的Id，可通过查询流量控制规则列表获取

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RemoveVpcFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVpcFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessDomainInfo struct {

	// 接入功能详情

	AccessServiceInfoList []*AccessServiceInfo `json:"AccessServiceInfoList,omitempty" name:"AccessServiceInfoList"`
	// 创建状态，0创建成功，1正在创建，大于1创建失败

	CreateStatus *int64 `json:"CreateStatus,omitempty" name:"CreateStatus"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 解析地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 是否影响nat防火墙，0不影响，1影响

	IsChangeNat *int64 `json:"IsChangeNat,omitempty" name:"IsChangeNat"`
	// 是否主域名

	IsMaster *uint64 `json:"IsMaster,omitempty" name:"IsMaster"`
	// 域名上次修改时间

	LastModifiedTime *uint64 `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`
	// 接入地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 开关，0关闭，1开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type FloatTrendsItem struct {

	// 时间坐标

	Label *string `json:"Label,omitempty" name:"Label"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type ModifyAllPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllPublicIPSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllRegionListRequest struct {
	*tchttp.BaseRequest

	// 页高

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *DescribeAllRegionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRegionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeCommonRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNoticeCommonRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeCommonRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DbRule struct {

	// 地址描述

	AddDetail *string `json:"AddDetail,omitempty" name:"AddDetail"`
	// 地址ip

	AddIp *string `json:"AddIp,omitempty" name:"AddIp"`
	// 地址名称

	AddName *string `json:"AddName,omitempty" name:"AddName"`
	// 城市code

	CityRegCode *int64 `json:"CityRegCode,omitempty" name:"CityRegCode"`
	// 城市

	CityRegName *string `json:"CityRegName,omitempty" name:"CityRegName"`
	// 云厂商code

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 国家code

	CountryRegCode *int64 `json:"CountryRegCode,omitempty" name:"CountryRegCode"`
	// 国家

	CountryRegName *string `json:"CountryRegName,omitempty" name:"CountryRegName"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 命中次数

	DetectedTimes *int64 `json:"DetectedTimes,omitempty" name:"DetectedTimes"`
	// 开关状态

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 导出用前端忽略

	DisableExport *string `json:"DisableExport,omitempty" name:"DisableExport"`
	// 目的类型

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 资产名称

	InsName *string `json:"InsName,omitempty" name:"InsName"`
	// 是否云厂商

	IsCloud *int64 `json:"IsCloud,omitempty" name:"IsCloud"`
	// 是否是地域

	IsReg *int64 `json:"IsReg,omitempty" name:"IsReg"`
	// 资产路径

	ResourcePath *string `json:"ResourcePath,omitempty" name:"ResourcePath"`
	// 源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源类型

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 规则id

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeFlowCenterLogsV1Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*FlowCenterData `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowCenterLogsV1Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterLogsV1Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest

	// 查询个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤的vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录集合

		Data []*WebServiceInfo `json:"Data,omitempty" name:"Data"`
		// 返回值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回值

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetFlowSwitchInfo struct {

	// 存在关闭开关的公网ip

	HaveClose *bool `json:"HaveClose,omitempty" name:"HaveClose"`
	// 存在开启开关的公网ip

	HaveOpen *bool `json:"HaveOpen,omitempty" name:"HaveOpen"`
	// 0&nbsp;扫描中&nbsp;1没有扫描&nbsp;2&nbsp;扫描完成

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type DescribeAcListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 不算筛选条数的总条数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// 访问控制列表数据

		Data []*AcListsData `json:"Data,omitempty" name:"Data"`
		// 访问控制规则全部启用/全部停用

		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAcListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SolidProtectionDataResp struct {

	// 云防火墙待处理告警数

	CfwAlertCount *int64 `json:"CfwAlertCount,omitempty" name:"CfwAlertCount"`
	// 云防火墙防护公网IP数

	CfwProtectedIpCount *int64 `json:"CfwProtectedIpCount,omitempty" name:"CfwProtectedIpCount"`
	// 主机安全受影响主机数

	CwpAffectedCount *int64 `json:"CwpAffectedCount,omitempty" name:"CwpAffectedCount"`
	// 主机安全待处理风险数

	CwpRiskCount *int64 `json:"CwpRiskCount,omitempty" name:"CwpRiskCount"`
	// 超过多少XX选择三道防线

	ProductPercentage *float64 `json:"ProductPercentage,omitempty" name:"ProductPercentage"`
	// 产品状态数组

	ProductStatata []*ProductStatusDataType `json:"ProductStatata,omitempty" name:"ProductStatata"`
	// 安全运营中心失陷告警数量

	SocAlertCount *int64 `json:"SocAlertCount,omitempty" name:"SocAlertCount"`
	// 安全运营中心云配置风险数

	SocRiskCount *int64 `json:"SocRiskCount,omitempty" name:"SocRiskCount"`
	// 容器安全防护实例数量

	TcssCvmCount *int64 `json:"TcssCvmCount,omitempty" name:"TcssCvmCount"`
	// 容器安全待处理风险数

	TcssRiskCount *int64 `json:"TcssRiskCount,omitempty" name:"TcssRiskCount"`
	// Web应用防火墙攻击次数

	WafAttackCount *int64 `json:"WafAttackCount,omitempty" name:"WafAttackCount"`
	// Web应用防火墙爬虫风险数

	WafCrawlerCount *int64 `json:"WafCrawlerCount,omitempty" name:"WafCrawlerCount"`
}

type ScanResultInfo struct {

	// 是否禁封端口

	BanStatus *bool `json:"BanStatus,omitempty" name:"BanStatus"`
	// 防护ip数量

	IPNum *uint64 `json:"IPNum,omitempty" name:"IPNum"`
	// 是否开启防护

	IPStatus *bool `json:"IPStatus,omitempty" name:"IPStatus"`
	// 是否拦截攻击

	IdpStatus *bool `json:"IdpStatus,omitempty" name:"IdpStatus"`
	// 暴露漏洞数量

	LeakNum *uint64 `json:"LeakNum,omitempty" name:"LeakNum"`
	// 暴露端口数量

	PortNum *uint64 `json:"PortNum,omitempty" name:"PortNum"`
}

type DescribeFlowStatisticsOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量统计

		Data []*AccessControlOverviewData `json:"Data,omitempty" name:"Data"`
		// 入站流量平均

		InAvg *float64 `json:"InAvg,omitempty" name:"InAvg"`
		// 入站流量峰值

		InMax *uint64 `json:"InMax,omitempty" name:"InMax"`
		// 出站流量平均

		OutAvg *float64 `json:"OutAvg,omitempty" name:"OutAvg"`
		// 出站流量峰值

		OutMax *uint64 `json:"OutMax,omitempty" name:"OutMax"`
		// 总流量平均

		TotalAvg *float64 `json:"TotalAvg,omitempty" name:"TotalAvg"`
		// 总高峰值

		TotalMax *uint64 `json:"TotalMax,omitempty" name:"TotalMax"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowStatisticsOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowStatisticsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 添加或删除操作的Dnat规则列表。

	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitempty" name:"AddOrDelDnatRules"`
	// 防火墙实例id，该字段必须传递。

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 0：cfw新增模式，1：cfw接入模式。

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 修改操作的新的Dnat规则

	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
	// 操作类型，可选值：add，del，modify。

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 修改操作的原始Dnat规则

	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`
}

func (r *SetNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id，该字段必须传递。

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 添加或删除操作的Dnat规则列表。

	DnatRules []*CfwNatDnatRule `json:"DnatRules,omitempty" name:"DnatRules"`
	// 0：cfw新增模式，1：cfw接入模式。

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *DeleteNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineUpdateDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎版本信息列表

		EngineVersionLst []*EngineVersionInfo `json:"EngineVersionLst,omitempty" name:"EngineVersionLst"`
		// 实例版本信息列表

		FwInsVersionLst []*CfwInsUpdateStatus `json:"FwInsVersionLst,omitempty" name:"FwInsVersionLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineUpdateDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineUpdateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProbeTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关列表

		Data *ProbeTaskDetail `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProbeTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WidthRuleLimit struct {

	// 带宽在该范围的最大规则数限制

	MaxRuleLimit *uint64 `json:"MaxRuleLimit,omitempty" name:"MaxRuleLimit"`
	// 带宽范围最大值

	MaxWidth *uint64 `json:"MaxWidth,omitempty" name:"MaxWidth"`
	// 带宽范围最小值

	MinWidth *uint64 `json:"MinWidth,omitempty" name:"MinWidth"`
}

type DescribeDefenseErrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息列表

		Data []*DefenseError `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenseErrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefenceApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyDefenceApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDefenceApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateCategory struct {

	// 蜜罐分类名称

	Category *string `json:"Category,omitempty" name:"Category"`
	// 该分类下的蜜罐服务列表

	Templates []*TemplateCategoryInfo `json:"Templates,omitempty" name:"Templates"`
}

type VpcFlowCenterData struct {

	// 会话数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 请求结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// index的名称

	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`
	// 访问流量请求

	RequestCount *int64 `json:"RequestCount,omitempty" name:"RequestCount"`
	// 访问流量响应

	ResponseCount *int64 `json:"ResponseCount,omitempty" name:"ResponseCount"`
	// 源资产id

	SourceAssetId *string `json:"SourceAssetId,omitempty" name:"SourceAssetId"`
	// 源资产名称

	SourceAssetName *string `json:"SourceAssetName,omitempty" name:"SourceAssetName"`
	// 源资产地域

	SourceAssetRegion *string `json:"SourceAssetRegion,omitempty" name:"SourceAssetRegion"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源vpc的id

	SourceVpcId *string `json:"SourceVpcId,omitempty" name:"SourceVpcId"`
	// 目的vpc的名称

	SourceVpcName *string `json:"SourceVpcName,omitempty" name:"SourceVpcName"`
	// 请求开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 目的资产id

	TargetAssetId *string `json:"TargetAssetId,omitempty" name:"TargetAssetId"`
	// 目的资产名称

	TargetAssetName *string `json:"TargetAssetName,omitempty" name:"TargetAssetName"`
	// 目的资产地域

	TargetAssetRegion *string `json:"TargetAssetRegion,omitempty" name:"TargetAssetRegion"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 目的vpc的id

	TargetVpcId *string `json:"TargetVpcId,omitempty" name:"TargetVpcId"`
	// 目的vpc的名称

	TargetVpcName *string `json:"TargetVpcName,omitempty" name:"TargetVpcName"`
}

type DescribeFlowApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserBandwidthUsageRequest struct {
	*tchttp.BaseRequest

	// 带宽统计结束时间&nbsp;默认为当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 带宽统计开始时间&nbsp;默认为三天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeUserBandwidthUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserBandwidthUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwGroupRequest struct {
	*tchttp.BaseRequest

	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 编辑的防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 修改防火墙(组)名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 编辑的防火墙实例列表

	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitempty" name:"VpcFwInstances"`
}

func (r *ModifyVpcFwGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateNetCard struct {

	// 资产实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 资产实例网卡列表

	SgInstList []*string `json:"SgInstList,omitempty" name:"SgInstList"`
}

type RegionConfig struct {

	// 地域

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 是否国外&nbsp;1是&nbsp;0不是

	Foreign *int64 `json:"Foreign,omitempty" name:"Foreign"`
	// 是否自驾云&nbsp;1是&nbsp;0不是

	IsAutoDriveCloud *int64 `json:"IsAutoDriveCloud,omitempty" name:"IsAutoDriveCloud"`
	// 是否支持nat&nbsp;1是&nbsp;0不是

	IsSupportNat *int64 `json:"IsSupportNat,omitempty" name:"IsSupportNat"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区信息

	RegionArea *string `json:"RegionArea,omitempty" name:"RegionArea"`
	// 地域中文

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DescribeTiCenterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集合

		Data []*TiCenterInfo `json:"Data,omitempty" name:"Data"`
		// 状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTiCenterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiCenterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwConfigRequest struct {
	*tchttp.BaseRequest

	// 管理是否保留云服务器的公网IP类型路由启用状态，0:&nbsp;默认关闭&nbsp;&nbsp;1:&nbsp;启用，将保留公网ip类型路由

	HoldPublicIPTypeRoute *int64 `json:"HoldPublicIPTypeRoute,omitempty" name:"HoldPublicIPTypeRoute"`
	// 管理网络异常bypass，1，开启；0，关闭

	ManageNetworkDownBypass *int64 `json:"ManageNetworkDownBypass,omitempty" name:"ManageNetworkDownBypass"`
	// 防火墙满载可持续时间，单位s，超过该时间将bypass防火墙，为0时不做任何操作

	OverLoadLimitTime *int64 `json:"OverLoadLimitTime,omitempty" name:"OverLoadLimitTime"`
}

func (r *ModifyNatFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindWechatUser struct {

	// 绑定时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 是否推送告警&nbsp;1是&nbsp;0否

	IsNotify *int64 `json:"IsNotify,omitempty" name:"IsNotify"`
	// 微信头像

	Photo *string `json:"Photo,omitempty" name:"Photo"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 微信公众平台unionId

	UnionId *string `json:"UnionId,omitempty" name:"UnionId"`
	// 微信昵称

	Username *string `json:"Username,omitempty" name:"Username"`
}

type RouteBackupInfo struct {

	// 目的cidr，vpc间防火墙的云联网模式的防火墙实例，该值为空

	DestCidr *string `json:"DestCidr,omitempty" name:"DestCidr"`
	// 对象名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// vpc间防火墙的云联网模式的下一跳为云联网路由表id

	NextHopId *string `json:"NextHopId,omitempty" name:"NextHopId"`
	// 下一跳的名称

	NextHopName *string `json:"NextHopName,omitempty" name:"NextHopName"`
	// 下一跳类型：CCN表示下一跳为云联网

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// 对象id，实例id或路由表id；云联网模式为接入实例id，vpc-xxx等

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 是否可编辑，0：可编辑；1：不可编辑

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 状态值，ACTIVE：有效的；NEXTHOP_INVALID：下一跳失效

	State *string `json:"State,omitempty" name:"State"`
	// 失效状态时，返回失效原因

	StateDesc *string `json:"StateDesc,omitempty" name:"StateDesc"`
}

type ScanInfo struct {

	// 进度

	ScanPercent *float64 `json:"ScanPercent,omitempty" name:"ScanPercent"`
	// 扫描结果信息

	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitempty" name:"ScanResultInfo"`
	// 扫描状态&nbsp;0扫描中&nbsp;1完成&nbsp;&nbsp;2未勾选自动扫描

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 预计完成时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type SequenceData struct {

	// 规则Id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 修改后执行顺序

	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
	// 修改前执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
}

type DescDnsAcItem struct {

	// 关联任务详情

	BetaList []*BetaInfoByACL `json:"BetaList,omitempty" name:"BetaList"`
	// 命中次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 内部使用的uuid，一般情况下不会使用到该字段

	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`
	// 规则有效性

	Invalid *uint64 `json:"Invalid,omitempty" name:"Invalid"`
	// 上次命中时间

	LastHitTime *string `json:"LastHitTime,omitempty" name:"LastHitTime"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：&nbsp;accept：放行&nbsp;drop：拒绝&nbsp;log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 访问源

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 访问源类型可以为&nbsp;ip,net,template,instance,group,tag

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 规则状态，查询规则命中详情时该字段有效，0：新增，1:&nbsp;已删除,&nbsp;2:&nbsp;编辑删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 访问目的

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 访问目的名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 访问目的类型可以为domain,template

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则对应的唯一id

	Uuid *uint64 `json:"Uuid,omitempty" name:"Uuid"`
}

type SwitchListData struct {

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内网IP

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
	// 最近扫描时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 风险端口数

	PortTimes *int64 `json:"PortTimes,omitempty" name:"PortTimes"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 公网&nbsp;IP&nbsp;类型

	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 扫描深度

	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`
	// 扫描状态

	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 防火墙开关

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

type AddAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，-1代表失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success代表成功，failed代表失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 创建成功后返回新策略的uuid

		RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id&nbsp;&nbsp;不是必填项可以为空，就是所有vpc间的访问控制规则

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *DescribeVpcRuleOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcRuleOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIPSProtectModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIPSProtectModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIPSProtectModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePcapTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 抓包任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePcapTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePcapTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleSwitchRequest struct {
	*tchttp.BaseRequest

	// 需要操作开关的规则ID列表，注意：如果传入的是[-1]将开启或关闭所有规则开关

	RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
	// on开启规则，off关闭规则

	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyVpcAcRuleSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoAssistantConf struct {

	// 规则分类

	Action []*StringValueTextEx `json:"Action,omitempty" name:"Action"`
	// 告警分类

	AlertKillChain []*StringValueTextEx `json:"AlertKillChain,omitempty" name:"AlertKillChain"`
	// 告警来源

	AlertSource []*StringValueTextEx `json:"AlertSource,omitempty" name:"AlertSource"`
	// 事件列表

	EventName []*string `json:"EventName,omitempty" name:"EventName"`
	// 告警等级

	Level []*StringValueTextEx `json:"Level,omitempty" name:"Level"`
	// 对象类型

	TargetType []*StringValueTextEx `json:"TargetType,omitempty" name:"TargetType"`
}

type LogAsyncRespond struct {

	// 查询结果

	Data *string `json:"Data,omitempty" name:"Data"`
	// 查询唯一Id

	QueryId *string `json:"QueryId,omitempty" name:"QueryId"`
	// 查询分片

	RangeField *string `json:"RangeField,omitempty" name:"RangeField"`
	// 查询结果，1可用结果接口轮询，非1查询错误

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 查询结果数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type NdrSwitchError struct {

	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 错误内部id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作类型&nbsp;开启&nbsp;1&nbsp;关闭&nbsp;0

	OperationType *int64 `json:"OperationType,omitempty" name:"OperationType"`
}

type DescribeExportHitLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出地址

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 状态码&nbsp;

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportHitLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportHitLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcTopolListsRequest struct {
	*tchttp.BaseRequest

	// json字符串

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeVpcTopolListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcTopolListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回防火墙部署的可用区列表

		RegionZoneLst []*RegionZone `json:"RegionZoneLst,omitempty" name:"RegionZoneLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleScopeInfo struct {

	// NAT边界防火墙：ALL，全局生效；ap-NewYork，生效的地域；cfwnat-xxx，生效基于实例维度

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 生效范围描述

	ScopeDesc *string `json:"ScopeDesc,omitempty" name:"ScopeDesc"`
}

type DescribeExportHitLogsRequest struct {
	*tchttp.BaseRequest

	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束值

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始值

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeExportHitLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportHitLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebAssetScanListRequest struct {
	*tchttp.BaseRequest

	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeWebAssetScanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetScanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PassWordData struct {

	// 密码

	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type UserEsData struct {

	// es地址

	Es *string `json:"Es,omitempty" name:"Es"`
	// 集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户列表

	Users *string `json:"Users,omitempty" name:"Users"`
}

type StopAsyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码0&nbsp;成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 0成功&nbsp;其他失败&nbsp;状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopAsyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopAsyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomWhiteRule struct {

	// 访问目的

	DstIP *string `json:"DstIP,omitempty" name:"DstIP"`
	// 规则ID

	IdsRuleId *string `json:"IdsRuleId,omitempty" name:"IdsRuleId"`
	// 规则名称

	IdsRuleName *string `json:"IdsRuleName,omitempty" name:"IdsRuleName"`
	// 访问源

	SrcIP *string `json:"SrcIP,omitempty" name:"SrcIP"`
}

type DescribeAssetsOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产图

		Data []*AssetsOverviewData `json:"Data,omitempty" name:"Data"`
		// 未开启防火墙资产个数

		OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`
		// 开启防火墙资产个数

		OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwReSelectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwReSelectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwReSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOACondition struct {

	// FilterGroups&nbsp;条件过滤组

	FilterGroups []*IOAFilterGroup `json:"FilterGroups,omitempty" name:"FilterGroups"`
	// Filters&nbsp;条件过滤

	Filters []*IOAFilter `json:"Filters,omitempty" name:"Filters"`
	// PageNum&nbsp;获取第几页(只支持32位)

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// PageSize&nbsp;每页获取数(只支持32位)

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// Sort&nbsp;排序字段

	Sort *IOASort `json:"Sort,omitempty" name:"Sort"`
}

type AddZBTiNoticeRequest struct {
	*tchttp.BaseRequest

	// 动作：status轮询状态，其他无意义

	Modus *string `json:"Modus,omitempty" name:"Modus"`
}

func (r *AddZBTiNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddZBTiNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyslogConnectionStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;1异常&nbsp;0正常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyslogConnectionStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSyslogConnectionStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllVPCSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除规则详情

		Info *int64 `json:"Info,omitempty" name:"Info"`
		// 状态&nbsp;0成功&nbsp;其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAllAccessControlRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapVisual struct {

	// 名称

	ChineseName *string `json:"ChineseName,omitempty" name:"ChineseName"`
	// 名称标识

	Name *string `json:"Name,omitempty" name:"Name"`
	// 流量大小

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 访问次数

	VisitTimes *int64 `json:"VisitTimes,omitempty" name:"VisitTimes"`
}

type AddNatAcRuleRequest struct {
	*tchttp.BaseRequest

	// 添加规则的来源，一般不需要使用，值insert_rule&nbsp;表示插入指定位置的规则；值batch_import&nbsp;表示批量导入规则；为空时表示添加规则

	From *string `json:"From,omitempty" name:"From"`
	// 需要添加的nat访问控制规则列表

	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddNatAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTiWhiteRequest struct {
	*tchttp.BaseRequest

	// 前后端调用约定&nbsp;TiSwitchIn/TiSwitchOut

	IocList []*string `json:"IocList,omitempty" name:"IocList"`
	// add&nbsp;添加
	// delete&nbsp;删除

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyTiWhiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTiWhiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcListsData struct {

	// 自动化助手信息

	AutoTask *string `json:"AutoTask,omitempty" name:"AutoTask"`
	// 城市代码

	City *int64 `json:"City,omitempty" name:"City"`
	// 云厂商代码

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 命中次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家代码

	Country *int64 `json:"Country,omitempty" name:"Country"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则目的类型

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 规则有效性
	// 1&nbsp;有效
	// 0&nbsp;无效

	Invalid *int64 `json:"Invalid,omitempty" name:"Invalid"`
	// 是否地域规则

	IsRegion *int64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 告警规则id

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 国家名称

	RegName1 *string `json:"RegName1,omitempty" name:"RegName1"`
	// 城市名称

	RegName2 *string `json:"RegName2,omitempty" name:"RegName2"`
	// 地域码信息

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 规则源类型

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 规则开关状态&nbsp;1打开&nbsp;0关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 策略

	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 规则唯一ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAddrTmplSelectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地址模版信息列表

		Data []*RuleSelectAddrTmplInfo `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddrTmplSelectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddrTmplSelectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果，-1:不存在&nbsp;1:创建中&nbsp;0:创建成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsDeliverTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsDeliverTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsDeliverTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDynamicListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 轮播banner信息

		AlertList []*AlertAd `json:"AlertList,omitempty" name:"AlertList"`
		// 最近一次更新时间

		LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
		// 动态信息

		ProductList []*ProductInfo `json:"ProductList,omitempty" name:"ProductList"`
		// 未读动态集合

		ReadList []*string `json:"ReadList,omitempty" name:"ReadList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDynamicListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDynamicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStatusListRequest struct {
	*tchttp.BaseRequest

	// 资产Id

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
}

func (r *DescribeIPStatusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPStatusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySerialRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySerialRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySerialRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessDomainItem struct {

	// 接入数据库数量

	AccessDataBaseCount *int64 `json:"AccessDataBaseCount,omitempty" name:"AccessDataBaseCount"`
	// 接入运维实例数量

	AccessInstanceCount *int64 `json:"AccessInstanceCount,omitempty" name:"AccessInstanceCount"`
	// 访问统计

	AccessStat []*RemoteAccessStat `json:"AccessStat,omitempty" name:"AccessStat"`
	// 接入web服务地址

	AccessWebAddress *string `json:"AccessWebAddress,omitempty" name:"AccessWebAddress"`
	// cname

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名类型；运维域名：1；Web域名：2；

	DomainType *int64 `json:"DomainType,omitempty" name:"DomainType"`
	// 0为新版，1是老版域名

	DomainVersion *int64 `json:"DomainVersion,omitempty" name:"DomainVersion"`
	// 0为未接入，1为已接入

	HasCertificate *int64 `json:"HasCertificate,omitempty" name:"HasCertificate"`
	// 0为未接入，1为已接入

	IsAccess *int64 `json:"IsAccess,omitempty" name:"IsAccess"`
	// 域名最后修改时间

	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`
	// 最近访问时间

	LastVisitTime *string `json:"LastVisitTime,omitempty" name:"LastVisitTime"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type CheckVpcFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态返回，true为冲突，false不冲突

		Status *bool `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckVpcFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVpcFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetflowRuleStatusAllRequest struct {
	*tchttp.BaseRequest

	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// delete&nbsp;全部删除&nbsp;(不删除置空)

	IsDelete *string `json:"IsDelete,omitempty" name:"IsDelete"`
	// 0&nbsp;关闭&nbsp;1&nbsp;打开

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyNetflowRuleStatusAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetflowRuleStatusAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAcRuleAllRequest struct {
	*tchttp.BaseRequest

	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// delete&nbsp;执行删除，其余置空

	IsDelete *string `json:"IsDelete,omitempty" name:"IsDelete"`
	// on&nbsp;开启&nbsp;，off&nbsp;关闭

	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyVpcAcRuleAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAcRuleAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleSelectAddrTmplInfo struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 写入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 模版内容

	IpString *string `json:"IpString,omitempty" name:"IpString"`
	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 协议类型

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 模版来源&nbsp;0是自定义模版&nbsp;1是预设模版

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模版排序

	TmplIndex *int64 `json:"TmplIndex,omitempty" name:"TmplIndex"`
	// 模版类型&nbsp;1ip&nbsp;5&nbsp;域名&nbsp;6&nbsp;端口

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 模版唯一值

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeleteBlockIgnoreRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlockIgnoreRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GuideUserStatus struct {

	// 1完成&nbsp;资产扫描界面&nbsp;0未完成

	AssetStatus *int64 `json:"AssetStatus,omitempty" name:"AssetStatus"`
	// 1完成&nbsp;0为完成

	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`
	// 1完成&nbsp;完成入侵防御&nbsp;0未完成

	IdpsStatus *int64 `json:"IdpsStatus,omitempty" name:"IdpsStatus"`
	// 1完成&nbsp;完成禁封端口&nbsp;0未完成

	PortStatus *int64 `json:"PortStatus,omitempty" name:"PortStatus"`
	// 1完成远程运维&nbsp;0未完成

	RemoteStatus *int64 `json:"RemoteStatus,omitempty" name:"RemoteStatus"`
	// 1完成互联网边界防火墙开关&nbsp;0未完成

	SwitchStatus *int64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
}

type CreateProbeTaskRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// ew：vpc防火墙，nat：nat防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 访问内网,1-开，0-关

	LeftInnerProbe *uint64 `json:"LeftInnerProbe,omitempty" name:"LeftInnerProbe"`
	// 左侧IP

	LeftIp *string `json:"LeftIp,omitempty" name:"LeftIp"`
	// 左侧访问防火墙,1-开，0-关

	LeftProbe *uint64 `json:"LeftProbe,omitempty" name:"LeftProbe"`
	// 失败重试次数

	RetryTimes *uint64 `json:"RetryTimes,omitempty" name:"RetryTimes"`
	// 访问互联网,1-开，0-关

	RightInternetProbe *uint64 `json:"RightInternetProbe,omitempty" name:"RightInternetProbe"`
	// 右侧IP

	RightIp *string `json:"RightIp,omitempty" name:"RightIp"`
	// 右侧访问防火墙,1-开，0-关

	RightProbe *uint64 `json:"RightProbe,omitempty" name:"RightProbe"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 时间格式or周期格式如month|3|00:00:00
	// week|Sunday|00:00:00
	// day|1|20:00:00。临时任务直接填时间

	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`
	// 1：临时任务，2：周期任务

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *CreateProbeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProbeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackWhiteQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlackWhiteQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlackWhiteQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 公网IP

	FireWallPublicIP *string `json:"FireWallPublicIP,omitempty" name:"FireWallPublicIP"`
	// 状态值，0:&nbsp;关闭&nbsp;,1:开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPublicIPSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BetaSGAcl struct {

	// cidr

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例name

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 双向下发

	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`
	// 是否新建

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 顺序

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 参数名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 模板id

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// cidr

	SouCidr *string `json:"SouCidr,omitempty" name:"SouCidr"`
	// 模板名称

	SouInstanceName *string `json:"SouInstanceName,omitempty" name:"SouInstanceName"`
	// 参数名称

	SouParameterName *string `json:"SouParameterName,omitempty" name:"SouParameterName"`
	// 内网ip

	SouPrivateIp *string `json:"SouPrivateIp,omitempty" name:"SouPrivateIp"`
	// 公网ip

	SouPublicIp *string `json:"SouPublicIp,omitempty" name:"SouPublicIp"`
	// 源

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 源类型

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 动作

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 目的

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 目的类型

	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type SecurityGroupInfo struct {

	// 安全组id

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组name

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
}

type SubAccount struct {

	// 子账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// open&nbsp;id

	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeBlockTimesListRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeBlockTimesListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockTimesListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全基线开关

		BaselineAllSwitch *int64 `json:"BaselineAllSwitch,omitempty" name:"BaselineAllSwitch"`
		// 基础防御开关

		BasicRuleSwitch *int64 `json:"BasicRuleSwitch,omitempty" name:"BasicRuleSwitch"`
		// 是否历史开启

		HistoryOpen *int64 `json:"HistoryOpen,omitempty" name:"HistoryOpen"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 威胁情报开关

		TiSwitch *int64 `json:"TiSwitch,omitempty" name:"TiSwitch"`
		// 虚拟补丁开关

		VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitempty" name:"VirtualPatchSwitch"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenseSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowCenterAssetListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 外联资产地域列表

		AreaList []*string `json:"AreaList,omitempty" name:"AreaList"`
		// 数据返回列表

		Data []*FlowCenterOutAsset `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowCenterAssetListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterAssetListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProbeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeProbeTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TiCenterInfo struct {

	// 作者

	Author *string `json:"Author,omitempty" name:"Author"`
	// 记录id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// ioc信息

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 关键字

	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`
	// 详情

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 漏洞id

	PocId *string `json:"PocId,omitempty" name:"PocId"`
	// 危险等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 防护状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// URL地址

	Url *string `json:"Url,omitempty" name:"Url"`
}

type ZeroTrustAccessSpecifications struct {

	// 授权地域数量

	AuthorizedRegionCount *uint64 `json:"AuthorizedRegionCount,omitempty" name:"AuthorizedRegionCount"`
	// 可接入身份数量

	IdentityCount *uint64 `json:"IdentityCount,omitempty" name:"IdentityCount"`
	// 剩余可接入身份数量

	RemainIdentityCount *int64 `json:"RemainIdentityCount,omitempty" name:"RemainIdentityCount"`
	// 剩余地域数量

	RemainRegionCount *uint64 `json:"RemainRegionCount,omitempty" name:"RemainRegionCount"`
	// 可接入VPC数量

	VpcCount *uint64 `json:"VpcCount,omitempty" name:"VpcCount"`
}

type DescribeGroupAddressIpListRequest struct {
	*tchttp.BaseRequest

	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
	// all/own

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeGroupAddressIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupAddressIpListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAclRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回查询MD5

		ReturnMd5 *string `json:"ReturnMd5,omitempty" name:"ReturnMd5"`
		// 返回信息&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAclRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BorderProtectDetail struct {

	// 域名化CLB的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip地址

	IP *string `json:"IP,omitempty" name:"IP"`
	// ip类型

	IPType *string `json:"IPType,omitempty" name:"IPType"`
	// 0观察&nbsp;1阻断&nbsp;2严格

	Idpsaction *int64 `json:"Idpsaction,omitempty" name:"Idpsaction"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// ip类型

	PublicIpType *int64 `json:"PublicIpType,omitempty" name:"PublicIpType"`
	// 地域

	ZoneAz *string `json:"ZoneAz,omitempty" name:"ZoneAz"`
}

type EngineVersionInfo struct {

	// 可回退到该版本的数量

	AbleBackCnt *int64 `json:"AbleBackCnt,omitempty" name:"AbleBackCnt"`
	// 可升级到该版本数量

	AbleUpdateCnt *int64 `json:"AbleUpdateCnt,omitempty" name:"AbleUpdateCnt"`
	// 版本描述详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 版本状态，stable：稳定版；previewed：预览版

	State *string `json:"State,omitempty" name:"State"`
	// 版本发布时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeProtectModeCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *ProtectCount `json:"Data,omitempty" name:"Data"`
		// 返回码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回值

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectModeCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectModeCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NATRegionStatus struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type FwGroupIpFlowStat struct {

	// 均值带宽

	FlowAvg *float64 `json:"FlowAvg,omitempty" name:"FlowAvg"`
	// 峰值流量

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 防火墙ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 0&nbsp;IPV4&nbsp;,&nbsp;1&nbsp;IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 限速状态

	TcRuleStatus *int64 `json:"TcRuleStatus,omitempty" name:"TcRuleStatus"`
	// VPCID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC&nbsp;名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 异地灾备&nbsp;1：使用异地灾备；0：不使用异地灾备

	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 是否建立

	IsCreateDomain *int64 `json:"IsCreateDomain,omitempty" name:"IsCreateDomain"`
	// 模式&nbsp;1：接入模式；0：新增模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 防火墙实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。

	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`
	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。

	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`
	// 带宽

	Width *int64 `json:"Width,omitempty" name:"Width"`
	// 主可用区，为空则选择默认可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 备可用区，为空则选择默认可用区

	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`
}

func (r *CreateNatFwInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwEdgeBarResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙通用实例总数

		AllInsQuota *int64 `json:"AllInsQuota,omitempty" name:"AllInsQuota"`
		// 绑定资产数

		AssetNum *int64 `json:"AssetNum,omitempty" name:"AssetNum"`
		// 0&nbsp;不开启自动防护
		// 1&nbsp;开启自动防护

		AutoDefence *int64 `json:"AutoDefence,omitempty" name:"AutoDefence"`
		// 购买配额

		BuyWidth *int64 `json:"BuyWidth,omitempty" name:"BuyWidth"`
		// 旁路入站峰值带宽，单位bps

		BypassInFlowMax *int64 `json:"BypassInFlowMax,omitempty" name:"BypassInFlowMax"`
		// 旁路入站峰值带宽，规格占比

		BypassInFlowPercent *float64 `json:"BypassInFlowPercent,omitempty" name:"BypassInFlowPercent"`
		// 旁路专用带宽

		BypassOnlyBandwidth *int64 `json:"BypassOnlyBandwidth,omitempty" name:"BypassOnlyBandwidth"`
		// 旁路出站峰值带宽，单位bps

		BypassOutFlowMax *int64 `json:"BypassOutFlowMax,omitempty" name:"BypassOutFlowMax"`
		// 旁路出站峰值带宽，规格占比

		BypassOutFlowPercent *float64 `json:"BypassOutFlowPercent,omitempty" name:"BypassOutFlowPercent"`
		// 可分配到串行防火墙和NAT防火墙的带宽值，将从旁路防火墙等额扣除

		CanUsedBandwidth *int64 `json:"CanUsedBandwidth,omitempty" name:"CanUsedBandwidth"`
		// 关闭开关数

		CloseSwitchNum *int64 `json:"CloseSwitchNum,omitempty" name:"CloseSwitchNum"`
		// 旁路防火墙弹性带宽上限

		ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
		// 弹性付费状态，-1无计费项，1正常，2隔离，3销毁

		ElasticPayStatus *int64 `json:"ElasticPayStatus,omitempty" name:"ElasticPayStatus"`
		// 旁路防火墙弹性按量计费开关

		ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
		// 入站峰值带宽&nbsp;单位&nbsp;bit

		InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`
		// 入站峰值带宽占用规格百分比

		InFlowPercent *float64 `json:"InFlowPercent,omitempty" name:"InFlowPercent"`
		// 剩余通用实例个数

		InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
		// 规格占比最大的串行地域

		MaxSerialRegion *string `json:"MaxSerialRegion,omitempty" name:"MaxSerialRegion"`
		// natfw&nbsp;实例个数

		NatInstanceNum *int64 `json:"NatInstanceNum,omitempty" name:"NatInstanceNum"`
		// NAT&nbsp;分配带宽&nbsp;单位&nbsp;M

		NatWidth *int64 `json:"NatWidth,omitempty" name:"NatWidth"`
		// 开启开关数

		OpenSwitchNum *int64 `json:"OpenSwitchNum,omitempty" name:"OpenSwitchNum"`
		// 出站峰值带宽&nbsp;单位&nbsp;bit

		OutFlowMax *int64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`
		// 出站峰值带宽占用规格百分比

		OutFlowPercent *float64 `json:"OutFlowPercent,omitempty" name:"OutFlowPercent"`
		// 串行防火墙弹性总带宽上限

		SerialElasticBandWidth *int64 `json:"SerialElasticBandWidth,omitempty" name:"SerialElasticBandWidth"`
		// 串行防火墙弹性按量计费开关

		SerialElasticSwitch *int64 `json:"SerialElasticSwitch,omitempty" name:"SerialElasticSwitch"`
		// 规格占比最大的串行地域入站峰值带宽，单位bps

		SerialRegionInFlowMax *int64 `json:"SerialRegionInFlowMax,omitempty" name:"SerialRegionInFlowMax"`
		// 串行地域入站峰值带宽，规格占比

		SerialRegionInFlowPercent *float64 `json:"SerialRegionInFlowPercent,omitempty" name:"SerialRegionInFlowPercent"`
		// 串行地域个数

		SerialRegionNum *int64 `json:"SerialRegionNum,omitempty" name:"SerialRegionNum"`
		// 规格占比最大的串行地域出站峰值带宽，单位bps

		SerialRegionOutFlowMax *int64 `json:"SerialRegionOutFlowMax,omitempty" name:"SerialRegionOutFlowMax"`
		// 串行地域出站峰值带宽，规格占比

		SerialRegionOutFlowPercent *float64 `json:"SerialRegionOutFlowPercent,omitempty" name:"SerialRegionOutFlowPercent"`
		// 串行带宽&nbsp;单位&nbsp;M

		SerialWidth *int64 `json:"SerialWidth,omitempty" name:"SerialWidth"`
		// 是否支持弹性

		SupportElastic *int64 `json:"SupportElastic,omitempty" name:"SupportElastic"`
		// 开关总个数

		SwitchNum *int64 `json:"SwitchNum,omitempty" name:"SwitchNum"`
		// 开关配额数

		SwitchQuota *int64 `json:"SwitchQuota,omitempty" name:"SwitchQuota"`
		// 剩余带宽&nbsp;单位&nbsp;M

		UnUsedWidth *int64 `json:"UnUsedWidth,omitempty" name:"UnUsedWidth"`
		// vpcfw&nbsp;实例个数

		VpcInstanceNum *int64 `json:"VpcInstanceNum,omitempty" name:"VpcInstanceNum"`
		// 购买带宽&nbsp;单位&nbsp;M

		Width *int64 `json:"Width,omitempty" name:"Width"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwEdgeBarResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEdgeBarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpsLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*IpsLogsData `json:"Data,omitempty" name:"Data"`
		// 判断来源列表

		DiffFromLists []*string `json:"DiffFromLists,omitempty" name:"DiffFromLists"`
		// 威胁类型名称列表

		EventNameLists []*string `json:"EventNameLists,omitempty" name:"EventNameLists"`
		// 条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpsLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySysLogTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 类型配置

	DeliveryConfigList []*SysLogTypeConfig `json:"DeliveryConfigList,omitempty" name:"DeliveryConfigList"`
}

func (r *ModifySysLogTypeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySysLogTypeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVpcSwitch struct {

	// 0：设置为关闭&nbsp;1:设置为打开

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// vpc&nbsp;id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeFuncDynamicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击类型

		Category []*string `json:"Category,omitempty" name:"Category"`
		// 规则详情

		Data []*DynamicDetail `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFuncDynamicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFuncDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeQuicklyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProbeQuicklyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProbeQuicklyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProbeSwitch struct {

	// 探针id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
}

type CreateSecurityGroupApiRulesRequest struct {
	*tchttp.BaseRequest

	// 云地域的英文简写

	Area *string `json:"Area,omitempty" name:"Area"`
	// 创建规则数据

	Data []*SecurityGroupApiRuleData `json:"Data,omitempty" name:"Data"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 插入类型，0：后插，1：前插，2：中插

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateSecurityGroupApiRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupApiRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleLogsRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每页个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 游标

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeRuleLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIgnoreRuleNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePcapTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePcapTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePcapTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStaticListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*StaticInfo `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockStaticListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockStaticListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBuyPageInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBuyPageInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBuyPageInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyCvmInfo struct {

	// cvm&nbsp;id

	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`
	// cvm&nbsp;名称

	CvmName *string `json:"CvmName,omitempty" name:"CvmName"`
}

type ModifyAclRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：修改成功，非0：修改失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 修改过后的uuid

		RuleUuid []*uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAclRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySysLogTypeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;1异常&nbsp;0成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySysLogTypeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySysLogTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatProbeTrafficStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatProbeTrafficStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatProbeTrafficStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescPcapTask struct {

	// 下载状态,1:可下载，0：不可下载

	DownEnable *int64 `json:"DownEnable,omitempty" name:"DownEnable"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 本端ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 抓包大小，单位B

	PkgSize *int64 `json:"PkgSize,omitempty" name:"PkgSize"`
	// 本端端口，-1表示所有

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// Protocol&nbsp;TCP,UDP,ICMP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务状态,-1:抓包失败，1：抓包中，0：抓包结束

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 对端ip

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 对端端口，-1表示所有

	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`
	// string

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type DescribeFilter struct {

	// 过滤的内容

	FilterContent *string `json:"FilterContent,omitempty" name:"FilterContent"`
	// 过滤的类型

	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type NatGatewayIpData struct {

	// 网关id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 可选弹性ip数组

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// 网关名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ActivityInfo struct {

	// 活动名称

	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
	// 活动结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 活动开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 活动代金券剩余数量

	VoucherCnt *uint64 `json:"VoucherCnt,omitempty" name:"VoucherCnt"`
}

type DescribeAssetEventTreeRequest struct {
	*tchttp.BaseRequest

	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeAssetEventTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEventTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHubDbAssetListRequest struct {
	*tchttp.BaseRequest

	// 资产大类，根据此字段时返回不同的子结构,AssetBasicType(只支持32位)

	AssetBasicType *uint64 `json:"AssetBasicType,omitempty" name:"AssetBasicType"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤

	Filter []*QueryFilterV3 `json:"Filter,omitempty" name:"Filter"`
	// 聚合字段

	GroupByField *string `json:"GroupByField,omitempty" name:"GroupByField"`
	// limit,最大值200(只支持32位)

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 逻辑字段(只支持32位)

	Logic *uint64 `json:"Logic,omitempty" name:"Logic"`
	// 起始，从0开始(只支持32位)

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// TimeRange&nbsp;（7天）&nbsp;24(小时）(只支持32位)

	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`
}

func (r *DescribeAssetHubDbAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHubDbAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceGroupOrderRequest struct {
	*tchttp.BaseRequest

	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 位置

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 移动的父id

	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
}

func (r *ModifyResourceGroupOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceGroupOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Threshold struct {

	// 第一级阈值

	FirstLevelThreshold *int64 `json:"FirstLevelThreshold,omitempty" name:"FirstLevelThreshold"`
	// 第二级阈值

	SecondLevelThreshold *int64 `json:"SecondLevelThreshold,omitempty" name:"SecondLevelThreshold"`
	// 第三级阈值

	ThirdLevelThreshold *int64 `json:"ThirdLevelThreshold,omitempty" name:"ThirdLevelThreshold"`
}

type CheckIpsRuleSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckIpsRuleSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpsRuleSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdsWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeIdsWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdsWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsCountAsyncRequest struct {
	*tchttp.BaseRequest

	// 查询唯一Id

	QueryId *string `json:"QueryId,omitempty" name:"QueryId"`
}

func (r *DescribeLogsCountAsyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsCountAsyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBetaTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBetaTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBetaTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoAssistantConfRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAutoAssistantConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoAssistantConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeEdgeRequest struct {
	*tchttp.BaseRequest

	// 查询过滤筛选字段内容

	Filter []*DescribeFilter `json:"Filter,omitempty" name:"Filter"`
	// 是否跨租户

	IsPeer *int64 `json:"IsPeer,omitempty" name:"IsPeer"`
}

func (r *DescribeNodeEdgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeEdgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOnceClickScanRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOnceClickScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOnceClickScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NdrSwitchWeight struct {

	// 实例ID开关

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开关权重

	SwitchWeight *int64 `json:"SwitchWeight,omitempty" name:"SwitchWeight"`
}

type ModifyWidthCheckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：不需要调整；1：需要调整

		AdjustFlag *int64 `json:"AdjustFlag,omitempty" name:"AdjustFlag"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWidthCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWidthCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveVpcAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功后返回被删除策略的uuid列表

		RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveVpcAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwSyncStatusRequest struct {
	*tchttp.BaseRequest

	// 查询的同步状态类型：Route,同步路由状态

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
}

func (r *DescribeFwSyncStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwSyncStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnWithRouteRequest struct {
	*tchttp.BaseRequest

	// 筛选过滤ccn

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCcnWithRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnWithRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值信息

		NatFwEipList []*NatFwEipsInfo `json:"NatFwEipList,omitempty" name:"NatFwEipList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwEipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwGroupSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关列表

		SwitchList []*FwGroupSwitchShow `json:"SwitchList,omitempty" name:"SwitchList"`
		// 开关总个数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwGroupSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeSwitchRecoveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEdgeSwitchRecoveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeSwitchRecoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTimesAndFlowAssetMaxTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*FlowCenterAssetTop `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVisitTimesAndFlowAssetMaxTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVisitTimesAndFlowAssetMaxTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarnAssetEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		Data []*WarnEventData `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarnAssetEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarnAssetEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstance struct {

	// ccn&nbsp;云联网实例ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// ccn&nbsp;云联网实例名称

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// ccn&nbsp;接入实例个数

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// ccn&nbsp;接入实例的名称ID信息

	InstanceLst []*CcnInstanceNameId `json:"InstanceLst,omitempty" name:"InstanceLst"`
	// 0&nbsp;当前ccn&nbsp;未被ccn防火墙接入，1：已被其他ccn防火墙接入；

	IsUsed *int64 `json:"IsUsed,omitempty" name:"IsUsed"`
	// ccn&nbsp;是否开白带路由表

	RouteTableFlag *bool `json:"RouteTableFlag,omitempty" name:"RouteTableFlag"`
	// ccn&nbsp;云联网实例状态

	State *string `json:"State,omitempty" name:"State"`
}

type NetflowCenterOutOverview struct {

	// 外联资产数量

	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`
	// 外联域名数量

	DomainCount *int64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 外联目的地址数量

	DstCount *int64 `json:"DstCount,omitempty" name:"DstCount"`
	// 失陷资产数量

	RiskAssetCount *int64 `json:"RiskAssetCount,omitempty" name:"RiskAssetCount"`
	// 外联风险域名数量

	RiskDomainCount *int64 `json:"RiskDomainCount,omitempty" name:"RiskDomainCount"`
	// 外联风险目的地址数量

	RiskDstCount *int64 `json:"RiskDstCount,omitempty" name:"RiskDstCount"`
}

type ZeroTrustRuleItem struct {

	// 用户APPID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 命令阻断acl&nbsp;id

	CommandGroupId *string `json:"CommandGroupId,omitempty" name:"CommandGroupId"`
	// 命中次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则是否开启，0未开启，1开启

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 访问目的示例：

	DstContent *string `json:"DstContent,omitempty" name:"DstContent"`
	// 资产

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 访问目的类型，类型可以为：资产实例ins,WEB服务web,数据库db,资产分组asset_group,资源标签tag

	DstType *string `json:"DstType,omitempty" name:"DstType"`
	// 过期时间

	Expire *string `json:"Expire,omitempty" name:"Expire"`
	// 权限规则ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 规则有效性

	Invalid *int64 `json:"Invalid,omitempty" name:"Invalid"`
	// 访问源示例：

	SrcContent *string `json:"SrcContent,omitempty" name:"SrcContent"`
	// 身份信息

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 访问源类型，类型可以为：用户user,用户组group,ip/cidr&nbsp;net

	SrcType *string `json:"SrcType,omitempty" name:"SrcType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 权限规则UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeSwitchListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域列表

		AreaLists []*string `json:"AreaLists,omitempty" name:"AreaLists"`
		// 列表数据

		Data []*SwitchListsData `json:"Data,omitempty" name:"Data"`
		// 关闭个数

		OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`
		// 打开个数

		OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcAclEdgeRangeRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 规则列表来源：rules，来源于当前已配置的规则；switchs，来源于开关列表

	FromList *string `json:"FromList,omitempty" name:"FromList"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcAclEdgeRangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcAclEdgeRangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 选中的防火墙开关Id

	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitempty" name:"FireWallPublicIPs"`
	// 状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAllPublicIPSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpsRuleListRequest struct {
	*tchttp.BaseRequest

	// 规则开关：Close/Open，规则动作：Alert/Drop，重置：Reset

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 修改的规则ID

	RuleId []*uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 是否选中所有规则。

	SelectAll *bool `json:"SelectAll,omitempty" name:"SelectAll"`
}

func (r *ModifyIpsRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpsRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlOverviewData struct {

	// 折线图x轴值

	Label *string `json:"Label,omitempty" name:"Label"`
	// 方向

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type FwGroupVpcFlowStat struct {

	// vpc&nbsp;的cidr

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 均值带宽

	FlowAvg *float64 `json:"FlowAvg,omitempty" name:"FlowAvg"`
	// 峰值流量

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 防火墙ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// vpc&nbsp;的ipv6&nbsp;cidr

	IPv6Cidr *string `json:"IPv6Cidr,omitempty" name:"IPv6Cidr"`
	// 是否开启限速

	TcRuleStatus *int64 `json:"TcRuleStatus,omitempty" name:"TcRuleStatus"`
	// VPCID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type HoneyPotTemplateItem struct {

	// 诱饵

	Bait *string `json:"Bait,omitempty" name:"Bait"`
	// 蜜罐使用负载均衡时默认路径

	DefaultPath *string `json:"DefaultPath,omitempty" name:"DefaultPath"`
	// 自定义诱饵项

	DefineBait *string `json:"DefineBait,omitempty" name:"DefineBait"`
	// 蜜罐描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 交互类型

	InteractionType *string `json:"InteractionType,omitempty" name:"InteractionType"`
	// 是否支持用户自定义密码

	NeedPasswd *int64 `json:"NeedPasswd,omitempty" name:"NeedPasswd"`
	// 默认暴露的端口

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 蜜罐名称类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 模版ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 是否是web类型蜜罐，0：非web类型;&nbsp;1：web类型,&nbsp;支持选择lb探针

	WebType *int64 `json:"WebType,omitempty" name:"WebType"`
}

type DescribeAclRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未过滤的总条数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// nat访问控制列表数据

		Data []*DescAcItem `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetOverviewNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *AssetOverviewEventNew `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetOverviewNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetOverviewNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHoneyPotListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		AreaLists []*SelectInfo `json:"AreaLists,omitempty" name:"AreaLists"`
		// 蜜罐服务列表数据

		Data []*HoneyPotInfo `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHoneyPotListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHoneyPotListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageLogTypeSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStorageLogTypeSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageLogTypeSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 限速的类型，可以是具体的IP也可以是CIDR，比如192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 需要添加限速规则的防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
	// 流量控制策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 入向限制速度，单位是kbps

	InLimit *uint64 `json:"InLimit,omitempty" name:"InLimit"`
	// 限速类型
	// IP限速：IP
	// 端口限速：PORT
	// 当前仅支持IP限速

	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`
	// 出向限制速度，单位是kbps

	OutLimit *uint64 `json:"OutLimit,omitempty" name:"OutLimit"`
	// 流量控制策略的端口。取值：
	// -1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyNatFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所需配额

		RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBlockIgnoreRuleNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSelectAssetGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*InstanceInfo `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域集合

		ZoneList []*AssetZone `json:"ZoneList,omitempty" name:"ZoneList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSelectAssetGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSelectAssetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回VPC防火墙部署的地域列表

		RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportFlowLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出地址

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 状态码&nbsp;1异常&nbsp;0正常

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportFlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatAcRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 描述过滤

	Description *string `json:"Description,omitempty" name:"Description"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 规则Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEdgeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcEdgeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEdgeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCfwIdpsModeRequest struct {
	*tchttp.BaseRequest

	// 是否支持公网ip&nbsp;1支持&nbsp;0不支持

	IsOpenPublicIp *int64 `json:"IsOpenPublicIp,omitempty" name:"IsOpenPublicIp"`
	// 值：IdpsAction，&nbsp;VirtualPatchSwitch

	SetKey *string `json:"SetKey,omitempty" name:"SetKey"`
	// 设置的值

	SetValue *int64 `json:"SetValue,omitempty" name:"SetValue"`
}

func (r *UpdateCfwIdpsModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfwIdpsModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwInsStatus struct {

	// 防火墙实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 实例名称

	CfwInsName *string `json:"CfwInsName,omitempty" name:"CfwInsName"`
	// 事件时间

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// 防火墙类型，nat：nat防火墙；ew：vpc间防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 恢复时间

	RecoverTime *string `json:"RecoverTime,omitempty" name:"RecoverTime"`
	// 实例所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例运行状态，Running：正常运行；BypassAutoFix：bypass修复；Updating：升级中；Expand：扩容中；BypassManual：手动触发bypass中；BypassAuto：自动触发bypass中

	Status *string `json:"Status,omitempty" name:"Status"`
	// Normal:&nbsp;正常模式
	// OnlyRoute:&nbsp;透明模式

	TrafficMode *string `json:"TrafficMode,omitempty" name:"TrafficMode"`
}

type DescribeFlowCenterAddressListsRequest struct {
	*tchttp.BaseRequest

	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 检索结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 排序条件

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 检索开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFlowCenterAddressListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterAddressListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwGroupSwitchRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcFwGroupSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceList struct {

	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type VpcFwInstanceShow struct {

	// VPC防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// VPC防火墙实例名称

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// 网络经过VPC防火墙CVM所在地域

	FwInsRegion *string `json:"FwInsRegion,omitempty" name:"FwInsRegion"`
}

type DescribeNatNewFlowStatsDataRequest struct {
	*tchttp.BaseRequest

	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// nat实例Id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间类型

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNatNewFlowStatsDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatNewFlowStatsDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本授权规则数

		AllowedNum *uint64 `json:"AllowedNum,omitempty" name:"AllowedNum"`
		// 入站启用规则数

		InEnableNum *uint64 `json:"InEnableNum,omitempty" name:"InEnableNum"`
		// 入站规则数

		InTotal *uint64 `json:"InTotal,omitempty" name:"InTotal"`
		// 预操作是否超过配额，1：超过，0：未超过

		IsOverLimit *int64 `json:"IsOverLimit,omitempty" name:"IsOverLimit"`
		// 预操作需要额外配额数

		NeedQuota *int64 `json:"NeedQuota,omitempty" name:"NeedQuota"`
		// 出站启用规则数

		OutEnableNum *uint64 `json:"OutEnableNum,omitempty" name:"OutEnableNum"`
		// 出站规则数

		OutTotal *uint64 `json:"OutTotal,omitempty" name:"OutTotal"`
		// 剩余通用规则扩展数

		RemainExtensionNum *int64 `json:"RemainExtensionNum,omitempty" name:"RemainExtensionNum"`
		// 剩余配额数

		RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 已使用规格数

		UsedNum *uint64 `json:"UsedNum,omitempty" name:"UsedNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetNatProbeEipTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeSetNatProbeEipTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetNatProbeEipTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 预操作增加规则列表个数

	AddCnt *int64 `json:"AddCnt,omitempty" name:"AddCnt"`
	// 预操作类型，rollback：回滚；cover_import：覆盖导入；其他：add

	OprTypeFrom *string `json:"OprTypeFrom,omitempty" name:"OprTypeFrom"`
}

func (r *DescribeVpcRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpsRuleDetail struct {

	// 置信度

	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`
	// 安全事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 自增id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 规则描述

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type NatRuleDetail struct {

	// 城市地域码

	City *int64 `json:"City,omitempty" name:"City"`
	// 云厂商代码

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 国家地域码

	Country *int64 `json:"Country,omitempty" name:"Country"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，取值：0，出站；1，入站，vpc间规则时无方向默认为0

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 访问目的名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 访问目的类型，1是ip，2是url，3是http域名

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 是否为地域规则

	IsRegion *int64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 规则顺序编号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 协议端口组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 协议端口组名称

	ParamTemplateName *string `json:"ParamTemplateName,omitempty" name:"ParamTemplateName"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：
	// TCP
	// UDP
	// ICMP
	// ANY
	// HTTP
	// HTTPS
	// HTTP/HTTPS
	// SMTP
	// SMTPS
	// SMTP/SMTPS
	// FTP
	// DNS
	// TLS/SSL

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// NAT边界防火墙：ALL，全局生效；ap-NewYork，生效的地域；cfwnat-xxx，生效基于实例维度

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 生效范围描述

	ScopeDesc *string `json:"ScopeDesc,omitempty" name:"ScopeDesc"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 访问源名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 源类型，1是ip，2是url，3是http域名

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 规则被删除：1，已删除；0，未删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// 0：观察
	// 1：拒绝
	// 2：放行

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.domain.com

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 规则对应的唯一id

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeNetFlowDoaminTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Data []*NetFlowDomainTopInfo `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态&nbsp;0：查询成功，非0：查询失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetFlowDoaminTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDoaminTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBorderFwRuleHitDetailRequest struct {
	*tchttp.BaseRequest

	// 规则对应的唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeBorderFwRuleHitDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderFwRuleHitDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwConfigRequest struct {
	*tchttp.BaseRequest

	// 防火墙满载可持续时间，单位s，超过该时间将bypass防火墙，为0时不做任何操作

	OverLoadLimitTime *int64 `json:"OverLoadLimitTime,omitempty" name:"OverLoadLimitTime"`
}

func (r *SetNatFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatProbeEipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 任务id集合

		TaskId []*string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatProbeEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatProbeEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwGroupIdNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFwGroupIdNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwGroupIdNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：正常，其它：不正常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcFwGroupInsShow struct {

	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// 防火墙(组)下防火墙实例

	FwInstanceLst []*VpcFwInstanceShow `json:"FwInstanceLst,omitempty" name:"FwInstanceLst"`
}

type AutoBackUpSetting struct {

	// 自动备份周期

	Period *string `json:"Period,omitempty" name:"Period"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 展示名

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 自动备份设置状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 自动备份子周期

	SubPeriod *string `json:"SubPeriod,omitempty" name:"SubPeriod"`
	// 自动备份执行时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 自动备份设置类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDefenseErrorRequest struct {
	*tchttp.BaseRequest

	// 错误类型，导入错误INTRUSION_DEFENSE_IMPORT

	ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`
	// 规则类型&nbsp;1封禁列表&nbsp;2放通列表

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// blocklist&nbsp;封禁列表
	// whitelist&nbsp;白名单列表

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeDefenseErrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseErrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideIdpsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新手引导拦截模式界面基本信息

		Data *IdpsInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuideIdpsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideIdpsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetInfoRequest struct {
	*tchttp.BaseRequest

	// 结束

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNetInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupChangeImpactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数据

		RuleImpactList *RuleImpact `json:"RuleImpactList,omitempty" name:"RuleImpactList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupChangeImpactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupChangeImpactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchCommonRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSwitchCommonRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchCommonRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回标签列表

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功
		// 其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRequest struct {
	*tchttp.BaseRequest

	// 需查询的地域列表

	RegionLst *string `json:"RegionLst,omitempty" name:"RegionLst"`
}

func (r *DescribeZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwEdgeIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ip&nbsp;开关列表

		Data []*EdgeIpInfo `json:"Data,omitempty" name:"Data"`
		// 实例类型列表

		InstanceTypeLst []*string `json:"InstanceTypeLst,omitempty" name:"InstanceTypeLst"`
		// 地域列表

		RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
		// ip&nbsp;开关列表个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwEdgeIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEdgeIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeScoreOverviewData struct {

	// 字段描述解释

	Info *string `json:"Info,omitempty" name:"Info"`
	// 维度名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评分数值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DescribeEngineUpdateOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEngineUpdateOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineUpdateOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcCfwWidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcCfwWidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcCfwWidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveOfflineExportTaskRequest struct {
	*tchttp.BaseRequest

	// 是否保留文件，1保留，非1删除

	KeepFile *int64 `json:"KeepFile,omitempty" name:"KeepFile"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RemoveOfflineExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveOfflineExportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatSwitchStatusInfo struct {

	// 防火墙开关状态，0：开启，1：关闭

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 自增唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 路由表id

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 路由表名称

	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`
	// 防火墙开关变动状态，0：稳定，1：开启中，2：关闭中，<0：错误码

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeEdgeFwFlowStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 旁路入站峰值流量

		InBypassFlowMax *int64 `json:"InBypassFlowMax,omitempty" name:"InBypassFlowMax"`
		// 旁路入站折线图数据

		InBypassLineCharLst []*FlowLineChar `json:"InBypassLineCharLst,omitempty" name:"InBypassLineCharLst"`
		// 入站峰值流量

		InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`
		// 入站折线图数据

		InLineCharLst []*FlowLineChar `json:"InLineCharLst,omitempty" name:"InLineCharLst"`
		// 串行入站峰值流量

		InSerialFlowMax *int64 `json:"InSerialFlowMax,omitempty" name:"InSerialFlowMax"`
		// 串行入站折线图数据

		InSerialLineCharLst []*FlowLineChar `json:"InSerialLineCharLst,omitempty" name:"InSerialLineCharLst"`
		// 出站旁路峰值流量

		OutBypassFlowMax *int64 `json:"OutBypassFlowMax,omitempty" name:"OutBypassFlowMax"`
		// 旁路出站折线图数据

		OutBypassLineCharLst []*FlowLineChar `json:"OutBypassLineCharLst,omitempty" name:"OutBypassLineCharLst"`
		// 出站峰值流量

		OutFlowMax *int64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`
		// 出站折线图数据

		OutLineCharLst []*FlowLineChar `json:"OutLineCharLst,omitempty" name:"OutLineCharLst"`
		// 出站串行峰值流量

		OutSerialFlowMax *int64 `json:"OutSerialFlowMax,omitempty" name:"OutSerialFlowMax"`
		// 串行出站折线图数据

		OutSerialLineCharLst []*FlowLineChar `json:"OutSerialLineCharLst,omitempty" name:"OutSerialLineCharLst"`
		// 串行带宽规格

		SerialWidth *int64 `json:"SerialWidth,omitempty" name:"SerialWidth"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 带宽规格

		Width *int64 `json:"Width,omitempty" name:"Width"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEdgeFwFlowStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeFwFlowStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessDatabaseDetail struct {

	// 数库库接入详情列表

	DatabaseInstanceInfoDetailList []*DatabaseInstanceInfoDetail `json:"DatabaseInstanceInfoDetailList,omitempty" name:"DatabaseInstanceInfoDetailList"`
	// 接入地域总量

	RegionCount *int64 `json:"RegionCount,omitempty" name:"RegionCount"`
	// 实例总量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type FlowLineChar struct {

	// 均值带宽

	FlowAvg *float64 `json:"FlowAvg,omitempty" name:"FlowAvg"`
	// 峰值流量

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 时间

	UpTime *string `json:"UpTime,omitempty" name:"UpTime"`
}

type CreateClsDeliverTaskRequest struct {
	*tchttp.BaseRequest

	// SecurityID

	SecurityID *string `json:"SecurityID,omitempty" name:"SecurityID"`
	// SecurityKey

	SecurityKey *string `json:"SecurityKey,omitempty" name:"SecurityKey"`
	// 来源

	SourceFrom *string `json:"SourceFrom,omitempty" name:"SourceFrom"`
}

func (r *CreateClsDeliverTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsDeliverTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEdgeIpFlowListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		AreaLists []*string `json:"AreaLists,omitempty" name:"AreaLists"`
		// ip&nbsp;流量列表

		Data []*EdgeIpFlowStat `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEdgeIpFlowListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeIpFlowListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSessionConnectionsRequest struct {
	*tchttp.BaseRequest

	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// nat实例Id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatSessionConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSessionConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest

	// 地域代码（例:&nbsp;ap-NewYork),支持云全部地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 0:&nbsp;出站规则，1：入站规则

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 0:&nbsp;不过滤，1：过滤掉正常规则，保留下发异常规则

	Filter *uint64 `json:"Filter,omitempty" name:"Filter"`
	// 每页条数，默认为10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 状态，'':&nbsp;全部，'0'：筛选停用规则，'1'：筛选启用规则

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeSecurityGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPayReleaseResourceRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyPayReleaseResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPayReleaseResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowTimesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*AssetStatic `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowTimesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowTimesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandCfwVerticalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandCfwVerticalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpandCfwVerticalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BetaInfoByACL struct {

	// 上次执行时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type VpcFwInstanceInfo struct {

	// VPC&nbsp;相关云联网ID列表

	CcnId []*string `json:"CcnId,omitempty" name:"CcnId"`
	// VPC&nbsp;相关云联网名称列表

	CcnName []*string `json:"CcnName,omitempty" name:"CcnName"`
	// 弹性带宽，单位Mbps

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 弹性开关&nbsp;1打开&nbsp;0关闭

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// 实例引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 内网间峰值带宽&nbsp;(单位&nbsp;bps&nbsp;)

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 对端部署信息

	FwCvmListPeerAll []*VpcFwCvmInsInfo `json:"FwCvmListPeerAll,omitempty" name:"FwCvmListPeerAll"`
	// VPC防火墙CVM的列表

	FwCvmLst []*VpcFwCvmInsInfo `json:"FwCvmLst,omitempty" name:"FwCvmLst"`
	// 防火墙网关信息

	FwGateway []*FwGateway `json:"FwGateway,omitempty" name:"FwGateway"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// VPC防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 对端id结合

	FwInsIdListPeerAll []*string `json:"FwInsIdListPeerAll,omitempty" name:"FwInsIdListPeerAll"`
	// 对端实例id

	FwInsIdPeer *string `json:"FwInsIdPeer,omitempty" name:"FwInsIdPeer"`
	// FwInsName

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// 对端name

	FwInsNameListPeerAll []*string `json:"FwInsNameListPeerAll,omitempty" name:"FwInsNameListPeerAll"`
	// VPC防火墙实例名称

	FwInsNamePeer *string `json:"FwInsNamePeer,omitempty" name:"FwInsNamePeer"`
	// VPC防火墙实例模式&nbsp;0:&nbsp;旧VPC模式防火墙&nbsp;1:&nbsp;CCN模式防火墙

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
	// VPC对端防火墙实例模式&nbsp;0:&nbsp;旧VPC模式防火墙&nbsp;1:&nbsp;CCN模式防火墙

	FwModePeer *int64 `json:"FwModePeer,omitempty" name:"FwModePeer"`
	// VPC防火墙开关个数

	FwSwitchNum *int64 `json:"FwSwitchNum,omitempty" name:"FwSwitchNum"`
	// 是否首次开通按量付费
	// 1&nbsp;是
	// 0&nbsp;不是

	IsFirstAfterPay *int64 `json:"IsFirstAfterPay,omitempty" name:"IsFirstAfterPay"`
	// 接入的vpc列表

	JoinInsIdLst []*string `json:"JoinInsIdLst,omitempty" name:"JoinInsIdLst"`
	// VPC防火墙接入网络实例类型列表

	JoinInsLst []*VpcFwJoinInstanceType `json:"JoinInsLst,omitempty" name:"JoinInsLst"`
	// VPC防火墙接入网络实例个数

	JoinInsNum *int64 `json:"JoinInsNum,omitempty" name:"JoinInsNum"`
	// VPC防火墙接入网络实例个数

	JoinPeerInsLst []*VpcFwJoinInstanceType `json:"JoinPeerInsLst,omitempty" name:"JoinPeerInsLst"`
	// VPC防火墙接入网络实例个数

	JoinPeerInsNum *int64 `json:"JoinPeerInsNum,omitempty" name:"JoinPeerInsNum"`
	// VPC&nbsp;相关对等连接ID列表

	PeerConnectionId []*string `json:"PeerConnectionId,omitempty" name:"PeerConnectionId"`
	// VPC&nbsp;相关对等连接名称列表

	PeerConnectionName []*string `json:"PeerConnectionName,omitempty" name:"PeerConnectionName"`
	// 跨租户状态&nbsp;0&nbsp;正常&nbsp;&nbsp;&nbsp;1&nbsp;对等销毁&nbsp;&nbsp;2还未接入对端

	PeerStatus *int64 `json:"PeerStatus,omitempty" name:"PeerStatus"`
	// 引擎预约升级时间

	ReserveTime *string `json:"ReserveTime,omitempty" name:"ReserveTime"`
	// 预约引擎升级版本

	ReserveVersion *string `json:"ReserveVersion,omitempty" name:"ReserveVersion"`
	// 引擎预约升级版本状态

	ReserveVersionState *string `json:"ReserveVersionState,omitempty" name:"ReserveVersionState"`
	// 最大规则数

	RuleMax *int64 `json:"RuleMax,omitempty" name:"RuleMax"`
	// 已使用规则数

	RuleUsed *int64 `json:"RuleUsed,omitempty" name:"RuleUsed"`
	// VPC防火墙状态&nbsp;0:正常&nbsp;，&nbsp;1：创建中&nbsp;2:&nbsp;变更中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// VPC防火墙创建时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 引擎运行模式，Normal:正常,&nbsp;OnlyRoute:透明模式

	TrafficMode *string `json:"TrafficMode,omitempty" name:"TrafficMode"`
	// 引擎是否可升级：0，不可升级；1，可升级

	UpdateEnable *int64 `json:"UpdateEnable,omitempty" name:"UpdateEnable"`
	// 用户VPC墙总带宽

	UserVpcWidth *int64 `json:"UserVpcWidth,omitempty" name:"UserVpcWidth"`
	// 防火墙实例带宽

	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type ZeroTrustCommandRuleItem struct {

	// 用户APPID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 生效范围资产，全部资产，Id填ALL

	Assets []*ZeroTrustEffectiveScopeItem `json:"Assets,omitempty" name:"Assets"`
	// 命令详情

	Commands *string `json:"Commands,omitempty" name:"Commands"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则是否开启，0未开启，1开启

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 命中次数

	HitTimes *int64 `json:"HitTimes,omitempty" name:"HitTimes"`
	// 策略ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 策略类型，whitelist白名单，blacklist黑名单

	Type *string `json:"Type,omitempty" name:"Type"`
	// 生效范围用户，全部用户，Id填ALL

	Users []*ZeroTrustEffectiveScopeItem `json:"Users,omitempty" name:"Users"`
	// 策略UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeShowBakRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份规则列表数据

		Data []*BakRuleListObj `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 备份规则总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShowBakRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShowBakRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogTopIpRequest struct {
	*tchttp.BaseRequest

	// 结束

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// top数

	Top *int64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeTLogTopIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogTopIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNotifyTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改状态&nbsp;0成功&nbsp;其他失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNotifyTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotifyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 编辑成功后返回新策略ID列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAclRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefenceApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDefenceApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDefenceApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChooseAssetRequest struct {
	*tchttp.BaseRequest

	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 资产类型&nbsp;1

	InsType *string `json:"InsType,omitempty" name:"InsType"`
	// 页大小&nbsp;&nbsp;0&nbsp;全部

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeChooseAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChooseAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowCenterLogsV1Request struct {
	*tchttp.BaseRequest

	// 外部地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 方向，1：外部访问&nbsp;0主动外联

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 地域位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 端口号

	Port *string `json:"Port,omitempty" name:"Port"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeFlowCenterLogsV1Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterLogsV1Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HoneyPotPortSelector struct {

	// 蜜罐服务默认ip

	DefaultPort *int64 `json:"DefaultPort,omitempty" name:"DefaultPort"`
	// 蜜罐服务交互类型

	InteractionType *string `json:"InteractionType,omitempty" name:"InteractionType"`
	// 负载均衡路径

	ProbePath *string `json:"ProbePath,omitempty" name:"ProbePath"`
	// 探针与蜜罐绑定的端口

	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
	// 探针与蜜罐绑定的端口范围

	ProbePortRange *string `json:"ProbePortRange,omitempty" name:"ProbePortRange"`
	// 蜜罐服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 蜜罐服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type AddIocFeedbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIocFeedbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIocFeedbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWProductPricesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品价格详情

		Data []*CfwPricesInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFWProductPricesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWProductPricesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFuncDynamicsRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 筛选结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 筛选开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFuncDynamicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFuncDynamicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupNewRequest struct {
	*tchttp.BaseRequest

	// 资产组id&nbsp;&nbsp;全部传0

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 查询类型&nbsp;网络结构-vpc，业务识别-resource&nbsp;，资源标签-tag

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// all&nbsp;&nbsp;包含子组&nbsp;own自己

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeResourceGroupNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest

	// Nat所在地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 状态值，0：检查表的状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeTableStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTableStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.domain.com,www.domain.com

	IpString *string `json:"IpString,omitempty" name:"IpString"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 协议端口模板，协议类型，4:4层协议，7:7层协议。Type=6时必填。

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 1&nbsp;ip模板
	// 5&nbsp;域名模板

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 地址模板唯一Id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ModifyAddressTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainParseIpLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 域名正向查询是否cname

		HasCName *uint64 `json:"HasCName,omitempty" name:"HasCName"`
		// ip列表

		IpList []*string `json:"IpList,omitempty" name:"IpList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainParseIpLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainParseIpLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 类型&nbsp;1&nbsp;告警&nbsp;2阻断

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 查询条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// top数

	Top *int64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeTLogIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogIpListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回多余的信息

		Info *string `json:"Info,omitempty" name:"Info"`
		// 状态值&nbsp;0正常&nbsp;其他异常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0：正常，-1：不正常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BakRuleAddName struct {

	// 地址模板数量

	AddressTemplate *int64 `json:"AddressTemplate,omitempty" name:"AddressTemplate"`
	// DNS规则标识

	DNS *int64 `json:"DNS,omitempty" name:"DNS"`
	// 放通列表数量

	Ignore *int64 `json:"Ignore,omitempty" name:"Ignore"`
	// 拦截列表数量

	Intercept *int64 `json:"Intercept,omitempty" name:"Intercept"`
	// 互联网边界规则列表数量

	Internet *int64 `json:"Internet,omitempty" name:"Internet"`
	// NAT规则地域状态列表

	NAT []*NATRegionStatus `json:"NAT,omitempty" name:"NAT"`
	// 企业安全规则组新数量

	SgNew *int64 `json:"SgNew,omitempty" name:"SgNew"`
	// 内网间规则标识

	VPC *string `json:"VPC,omitempty" name:"VPC"`
}

type AssetDayEventInfo struct {

	// 次数

	Count *string `json:"Count,omitempty" name:"Count"`
	// 事件详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，0：出站，1：入站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 安全事件类型

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AuthInfo struct {

	// 1&nbsp;&nbsp;赠送过&nbsp;0没有&nbsp;&nbsp;开启开关

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 1&nbsp;&nbsp;赠送过&nbsp;0没有&nbsp;扫描

	IsScan *int64 `json:"IsScan,omitempty" name:"IsScan"`
}

type DescribeLocalAccountAccountGroupsData struct {

	// 组Id(只支持32位)

	AccountGroupId *int64 `json:"AccountGroupId,omitempty" name:"AccountGroupId"`
	// 组名称

	AccountGroupName *string `json:"AccountGroupName,omitempty" name:"AccountGroupName"`
	// 组路径

	AccountGroupNamePaths []*string `json:"AccountGroupNamePaths,omitempty" name:"AccountGroupNamePaths"`
	// 组路径Id(只支持32位)

	AccountGroupPathIds []*string `json:"AccountGroupPathIds,omitempty" name:"AccountGroupPathIds"`
}

type DescribeAsyncTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务类型t&nbsp;&nbsp;acl_secgroupnew_import&nbsp;acl_vpc_import&nbsp;acl_dns_import&nbsp;

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeAsyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupChangeImpactRequest struct {
	*tchttp.BaseRequest

	// 资源组变更类型

	ResourceGroupChangeType *int64 `json:"ResourceGroupChangeType,omitempty" name:"ResourceGroupChangeType"`
	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
	// 资源组名称，名称变更场景必填

	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`
	// 资源组新父id，移动场景必填

	ResourceGroupNowParentId *string `json:"ResourceGroupNowParentId,omitempty" name:"ResourceGroupNowParentId"`
}

func (r *DescribeResourceGroupChangeImpactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupChangeImpactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcTopolListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcTopolListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcTopolListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatAcRuleSwitchAllResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作总条数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatAcRuleSwitchAllResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAcRuleSwitchAllResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveAclRuleRequest struct {
	*tchttp.BaseRequest

	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则

	RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *RemoveAclRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveAclRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwNatDnatRule struct {

	// NAT防火墙转发规则描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 网络协议，可选值：TCP、UDP。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 内网地址。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 内网端口。

	PrivatePort *int64 `json:"PrivatePort,omitempty" name:"PrivatePort"`
	// 弹性IP。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 公网端口。

	PublicPort *int64 `json:"PublicPort,omitempty" name:"PublicPort"`
}

type SelectInfo struct {

	// 名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagLog struct {

	// 新的标签列表

	NewTags []*TagInfo `json:"NewTags,omitempty" name:"NewTags"`
	// 旧的标签列表

	OldTags []*TagInfo `json:"OldTags,omitempty" name:"OldTags"`
	// 资源Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type DescribeApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowDistributeListRequest struct {
	*tchttp.BaseRequest

	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序类型，0流量大小倒序，1访问次数倒序

	SortType *int64 `json:"SortType,omitempty" name:"SortType"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
	// 地域类型，0全球，1中国

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeFlowDistributeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowDistributeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户信息概览的JSON字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveVpcFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveVpcFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVpcFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHubDbAssetListInfo struct {

	// 运维端口,0正常,1正在接入,2接入失败

	AccessStatus *int64 `json:"AccessStatus,omitempty" name:"AccessStatus"`
	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 配置风险

	ConfigureRisk *uint64 `json:"ConfigureRisk,omitempty" name:"ConfigureRisk"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 资产发现时间

	DiscoverTime *string `json:"DiscoverTime,omitempty" name:"DiscoverTime"`
	// 地址

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 最近扫描时间

	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
	// 网络访问

	NetworkAccess *uint64 `json:"NetworkAccess,omitempty" name:"NetworkAccess"`
	// 网络攻击

	NetworkAttack *uint64 `json:"NetworkAttack,omitempty" name:"NetworkAttack"`
	// 私有Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 扫描任务

	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`
	// 暴露状态&nbsp;0为未暴露&nbsp;1为暴露

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 标签列表

	Tag []*Tags `json:"Tag,omitempty" name:"Tag"`
	// 私有网络ID

	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`
	// 私有网络

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 运维端口,0为未接入

	YWPort *int64 `json:"YWPort,omitempty" name:"YWPort"`
}

type IpStatStruct struct {

	// 平均入向带宽&nbsp;（单位b)

	AvgInPkt *int64 `json:"AvgInPkt,omitempty" name:"AvgInPkt"`
	// 平均出向带宽&nbsp;（单位b)

	AvgOutPkt *int64 `json:"AvgOutPkt,omitempty" name:"AvgOutPkt"`
	// 平均出/入向带宽的条数

	Cnt *int64 `json:"Cnt,omitempty" name:"Cnt"`
	// 入向峰值带宽&nbsp;（单位b)

	InPktMax *int64 `json:"InPktMax,omitempty" name:"InPktMax"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// cvm或其他网络实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// IP&nbsp;地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 出向峰值带宽&nbsp;（单位b)

	OutPktMax *int64 `json:"OutPktMax,omitempty" name:"OutPktMax"`
	// 是否已开启限速开关&nbsp;1：已开启&nbsp;0&nbsp;:&nbsp;未开启

	TcRuleStatus *int64 `json:"TcRuleStatus,omitempty" name:"TcRuleStatus"`
	// 所属vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 所属vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeOfflineExportTemporaryCredentialsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		Link *string `json:"Link,omitempty" name:"Link"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOfflineExportTemporaryCredentialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOfflineExportTemporaryCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回树形结构

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagIpListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ip数组

		IpList []*string `json:"IpList,omitempty" name:"IpList"`
		// 0表示成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据

	CoverDuplicate *int64 `json:"CoverDuplicate,omitempty" name:"CoverDuplicate"`
	// 导入文件地址

	ImportFileUrl *string `json:"ImportFileUrl,omitempty" name:"ImportFileUrl"`
	// 规则类型，1封禁，2放通，不支持域名封禁

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// blocklist&nbsp;封禁列表
	// whitelist&nbsp;放通列表

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *ImportBlockIgnoreRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportBlockIgnoreRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlterFeedBackShow struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 告警规则误报，告警方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 告警规则误报，告警事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 误报反馈时间

	FeedBackTime *string `json:"FeedBackTime,omitempty" name:"FeedBackTime"`
	// 误报Ioc

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 误报Ioc类型，0&nbsp;误报IP/域名，1告警规则误报

	IocType *int64 `json:"IocType,omitempty" name:"IocType"`
	// 告警规则误报，告警攻击链

	KillChain *string `json:"KillChain,omitempty" name:"KillChain"`
	// 告警规则误报，告警事件来源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 反馈状态

	ReScanReason *int64 `json:"ReScanReason,omitempty" name:"ReScanReason"`
	// 告警规则误报，告警策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

type CreateChooseVpcsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateChooseVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateChooseVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwStatusBarResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙状态栏数据

		FwBarStatus *VpcFwBarStatus `json:"FwBarStatus,omitempty" name:"FwBarStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwStatusBarResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwStatusBarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAllRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;修改成功,&nbsp;其他:&nbsp;修改失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupAllRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventNewsRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeEventNewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventNewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceProbeInfo struct {

	// 容器内端口

	DefaultPort *int64 `json:"DefaultPort,omitempty" name:"DefaultPort"`
	// 探针端口

	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
	// 蜜罐服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type DescribeAlertCenterEventNameSelectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 返回值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回值

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertCenterEventNameSelectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertCenterEventNameSelectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetflowCenterTrendsRequest struct {
	*tchttp.BaseRequest

	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 方向，0主动外联，1外部访问

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 检索结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 检索开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNetflowCenterTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetflowCenterTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorDetail struct {

	// 错误类型
	// &nbsp;NAT_FW_PROBE_FAIL拨测切换失败&nbsp;&nbsp;NAT_FW_PROBE_SUC&nbsp;拨测切换成功

	ErrType *string `json:"ErrType,omitempty" name:"ErrType"`
	// 错误信息

	Errmsg *string `json:"Errmsg,omitempty" name:"Errmsg"`
	// 事件类型：0，操作性事件；1，状态型事件

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// 错误内部id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 通常为请求实例唯一ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 操作类型&nbsp;创建实例;配置;重新选择等

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	// 云API请求ID

	RequestID *string `json:"RequestID,omitempty" name:"RequestID"`
	// 状态码&nbsp;0&nbsp;存在&nbsp;-1&nbsp;忽略

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务类别&nbsp;CFWINS:防火墙实例

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type FlowCenterAssetTop struct {

	// 资产id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 数据值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeNatFwVpcDnsLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回参数

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 开关总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// nat防火墙vpc&nbsp;dns&nbsp;信息数组

		VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitempty" name:"VpcDnsSwitchLst"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwVpcDnsLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwVpcDnsLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetScanFilterRequest struct {
	*tchttp.BaseRequest

	// 过滤的ip集合

	FilterList []*string `json:"FilterList,omitempty" name:"FilterList"`
}

func (r *ModifyAssetScanFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetScanFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcFwGroupRequest struct {
	*tchttp.BaseRequest

	// 是否删除整个防火墙(组)
	// 0：不删除防火墙(组)，只删除单独实例
	// 1：删除整个防火墙(组)

	DeleteFwGroup *int64 `json:"DeleteFwGroup,omitempty" name:"DeleteFwGroup"`
	// 防火墙(组)Id

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 待删除的防火墙实例数组

	VpcFwInsList []*string `json:"VpcFwInsList,omitempty" name:"VpcFwInsList"`
}

func (r *DeleteVpcFwGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcFwGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddrTemplateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板列表数据

		Data []*TemplateListInfo `json:"Data,omitempty" name:"Data"`
		// 模板名称列表

		NameList []*string `json:"NameList,omitempty" name:"NameList"`
		// 模板总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddrTemplateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddrTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEdgeFwFlowStatRequest struct {
	*tchttp.BaseRequest

	// asc&nbsp;升序
	// desc&nbsp;降序

	By *string `json:"By,omitempty" name:"By"`
	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 防火墙实例ID

	FwRegion *string `json:"FwRegion,omitempty" name:"FwRegion"`
	// 开关流量统计页分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开关流量统计页分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeEdgeFwFlowStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeFwFlowStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwVpcStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vpc&nbsp;峰值带宽统计列表

		Data []*FwGroupVpcFlowStat `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总行数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwVpcStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwVpcStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreRulesImportProgressRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBlockIgnoreRulesImportProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreRulesImportProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 返回值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回值

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHitRuleLogsRequest struct {
	*tchttp.BaseRequest

	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeHitRuleLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHitRuleLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNtaFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Data []*NtaFileInfo `json:"Data,omitempty" name:"Data"`
		// 文件类型

		FileTypeList []*string `json:"FileTypeList,omitempty" name:"FileTypeList"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNtaFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNtaFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwVersionRequest struct {
	*tchttp.BaseRequest

	// 查询的防火墙实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
}

func (r *DescribeCfwVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNtaFileDownUrlRequest struct {
	*tchttp.BaseRequest

	// cos文件名

	CosFile []*string `json:"CosFile,omitempty" name:"CosFile"`
}

func (r *DescribeNtaFileDownUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNtaFileDownUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperatorLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作日志对象

		Data []*OperatorLogsData `json:"Data,omitempty" name:"Data"`
		// 是否到底了

		IsEnd *bool `json:"IsEnd,omitempty" name:"IsEnd"`
		// 游标

		Offset *string `json:"Offset,omitempty" name:"Offset"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperatorLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperatorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveVpcAcRuleRequest struct {
	*tchttp.BaseRequest

	// 仅当RuleUuids为-1时有效；0：删除Ipv4规则，1：删除Ipv6规则；默认为Ipv4类型规则

	IpVersion *uint64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则

	RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
}

func (r *RemoveVpcAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveVpcAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVpcFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVpcFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainInsInfoRequest struct {
	*tchttp.BaseRequest

	// 查询指定域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序参数，访问流量从小到大flowCount:asc，访问流量从大到小flowCount:desc，最近访问时间从小到大end_time:asc，最近访问时间从大到小end_time:desc，访问次数从小到大total_count:asc，访问次数从大到小total_count:desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 时间类型，0：7天，1：24小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNetFlowDomainInsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainInsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwRuleHitDetailRequest struct {
	*tchttp.BaseRequest

	// 规则对应的唯一id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DescribeVpcFwRuleHitDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwRuleHitDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVpcFwSwitchModeRequest struct {
	*tchttp.BaseRequest

	// vpc间防火墙实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 当前需要执行的开关操作，1：开；0：关

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *QueryVpcFwSwitchModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVpcFwSwitchModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopAsyncTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务类型&nbsp;acl_border_import&nbsp;_in&nbsp;
	// acl_border_import&nbsp;_out
	// acl_secgroupnew_import&nbsp;acl_vpc_import&nbsp;acl_dns_import&nbsp;
	// acl_nat_import_in
	// acl_nat_import_out

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *StopAsyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopAsyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClsDeliverTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值&nbsp;0：成功，非&nbsp;0：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClsDeliverTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsDeliverTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrictModeConfigRequest struct {
	*tchttp.BaseRequest

	// 是否自动封禁域名

	AutoBlockDomain *int64 `json:"AutoBlockDomain,omitempty" name:"AutoBlockDomain"`
	// 是否自动封禁IP

	AutoBlockIp *int64 `json:"AutoBlockIp,omitempty" name:"AutoBlockIp"`
	// 封禁时长类型，
	// Fixed固定时长，Step阶梯时长

	BlockDurationType *string `json:"BlockDurationType,omitempty" name:"BlockDurationType"`
	// 第一阶梯时长，分钟单位

	FirstBlockStepDuration *int64 `json:"FirstBlockStepDuration,omitempty" name:"FirstBlockStepDuration"`
	// 固定封禁时长，单位分钟

	FixedBlockDuration *int64 `json:"FixedBlockDuration,omitempty" name:"FixedBlockDuration"`
	// 第二阶梯时长，单位分钟

	SecondBlockStepDuration *int64 `json:"SecondBlockStepDuration,omitempty" name:"SecondBlockStepDuration"`
	// 第三阶梯时长，单位分钟

	ThirdBlockStepDuration *int64 `json:"ThirdBlockStepDuration,omitempty" name:"ThirdBlockStepDuration"`
}

func (r *ModifyStrictModeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrictModeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpTcRuleRequest struct {
	*tchttp.BaseRequest

	// 查询的ip或网段

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 防火墙实例

	FwId *string `json:"FwId,omitempty" name:"FwId"`
}

func (r *DescribeIpTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagIpListRequest struct {
	*tchttp.BaseRequest

	// 	nat：nat防火墙，ew：vpc间(东西向)防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// ip类型，公网IP：Public；内网IP：Private

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 	Ip版本，0：IPv4，1：IPv6，默认或空值时为IPv4

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 标签值

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeTagIpListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagIpListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPayReleaseResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 销毁状态&nbsp;0成功&nbsp;1失败

		PayReleaseResourceStatus *uint64 `json:"PayReleaseResourceStatus,omitempty" name:"PayReleaseResourceStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPayReleaseResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPayReleaseResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpcRuleHitTimesRequest struct {
	*tchttp.BaseRequest

	// 规则的uuid，可通过查询规则列表获取

	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *ResetVpcRuleHitTimesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpcRuleHitTimesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwVpcInfo struct {

	// vpc的cidr

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// vpc内cvm数量

	CvmNum *int64 `json:"CvmNum,omitempty" name:"CvmNum"`
	// dns地址

	DnsServer *string `json:"DnsServer,omitempty" name:"DnsServer"`
	// ipv6网段

	Ipv6Cidr *string `json:"Ipv6Cidr,omitempty" name:"Ipv6Cidr"`
	// 是否是新增资产，1：新增资产

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 要求选中对端vpc&nbsp;id&nbsp;列表

	PeerVpcLst []*string `json:"PeerVpcLst,omitempty" name:"PeerVpcLst"`
	// 地域信息

	RegionStr *string `json:"RegionStr,omitempty" name:"RegionStr"`
	// 0&nbsp;：正常vpc&nbsp;1:&nbsp;跨帐号vpc

	RemoteVpc *int64 `json:"RemoteVpc,omitempty" name:"RemoteVpc"`
	// vpc内子网数量

	SubnetNum *int64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// 日志主题id

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 0：未接入防火墙；1：已接入防火墙

	VpcFwJoinUsed *int64 `json:"VpcFwJoinUsed,omitempty" name:"VpcFwJoinUsed"`
	// vpc的Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeProtectModeCountRequest struct {
	*tchttp.BaseRequest

	// border&nbsp;互联网,&nbsp;nat&nbsp;natfw,&nbsp;vpc&nbsp;vpcfw,&nbsp;all&nbsp;&nbsp;四个值取一个

	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

func (r *DescribeProtectModeCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectModeCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPDefendStatus struct {

	// ip地址

	IP *string `json:"IP,omitempty" name:"IP"`
	// 防护状态&nbsp;&nbsp;&nbsp;1:防护打开;&nbsp;-1:地址错误;&nbsp;其他:未防护

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type RuleCount struct {

	// 互联网规则数量

	InternetRuleCount *int64 `json:"InternetRuleCount,omitempty" name:"InternetRuleCount"`
	// nat规则数量

	NatRuleCount []*RuleCountInfo `json:"NatRuleCount,omitempty" name:"NatRuleCount"`
	// 规则总量

	RuleTotalCount *int64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
	// 企业安全组规则数量

	SecurityGroupRuleCount []*RuleCountInfo `json:"SecurityGroupRuleCount,omitempty" name:"SecurityGroupRuleCount"`
	// vpc规则数量

	VpcRuleCount []*RuleCountInfo `json:"VpcRuleCount,omitempty" name:"VpcRuleCount"`
}

type DescribeAsyncTaskErrRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务类型&nbsp;acl_secgroupnew_import&nbsp;acl_vpc_import&nbsp;acl_dns_import&nbsp;

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeAsyncTaskErrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncTaskErrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户使用情况

		Data *CfwUserStatusData `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBorderProtectListRequest struct {
	*tchttp.BaseRequest

	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索参数，json格式字符串，空则传"{}"

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeBorderProtectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderProtectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowCenterLogsRequest struct {
	*tchttp.BaseRequest

	// 外部地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 方向，1：外部访问&nbsp;0主动外联

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 地域位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 端口号

	Port *string `json:"Port,omitempty" name:"Port"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeFlowCenterLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NtaSwitch struct {

	// 创建私有连接支持指定ip创建

	EndpointIp *string `json:"EndpointIp,omitempty" name:"EndpointIp"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 选择引流子网，用于创建私有连接

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeEdgeCFWSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEdgeCFWSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEdgeCFWSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwSwitchRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 条数，分页用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，分页用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNatFwSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewAuthInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNewAuthInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewAuthInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 当OperationType&nbsp;为bind或unbind操作时，使用该字段。

	EipList []*string `json:"EipList,omitempty" name:"EipList"`
	// bind：绑定eip；unbind：解绑eip；newAdd：新增防火墙弹性公网ip

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SetNatFwEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwEipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnWithEdgeRequest struct {
	*tchttp.BaseRequest

	// 筛选过滤ccn

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCcnWithEdgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnWithEdgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryNotEmptyRuleListInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份规则新增列表名称

		Data *BakRuleAddName `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 备份规则查询列表名称

		SearName []*BakRuleSearchName `json:"SearName,omitempty" name:"SearName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueryNotEmptyRuleListInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryNotEmptyRuleListInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwGroupInsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcFwGroupInsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupInsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwSequenceRulesRequest struct {
	*tchttp.BaseRequest

	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号

	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitempty" name:"RuleChangeItems"`
}

func (r *ModifyVpcFwSequenceRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwSequenceRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WeChatWorkUser struct {

	// 企业微信组织Id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 部门Id列表

	DepartmentIds []*string `json:"DepartmentIds,omitempty" name:"DepartmentIds"`
	// 接入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 备注

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 7天访问数

	VisitsTimes *int64 `json:"VisitsTimes,omitempty" name:"VisitsTimes"`
	// 企业微信Id

	WeChatWorkId *string `json:"WeChatWorkId,omitempty" name:"WeChatWorkId"`
	// 运维Id

	YwId *string `json:"YwId,omitempty" name:"YwId"`
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest

	// 每页最多个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// natfw&nbsp;防火墙实例id

	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`
	// natfw&nbsp;过滤，以','分隔

	NatInsIdFilter *string `json:"NatInsIdFilter,omitempty" name:"NatInsIdFilter"`
	// 分页页数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNatFwVpcDnsLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwVpcDnsLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsDeliverTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集ID

		LogID *string `json:"LogID,omitempty" name:"LogID"`
		// 日志集名称

		LogName *string `json:"LogName,omitempty" name:"LogName"`
		// SecurityID

		SecurityID *string `json:"SecurityID,omitempty" name:"SecurityID"`
		// SecurityKey

		SecurityKey *string `json:"SecurityKey,omitempty" name:"SecurityKey"`
		// 投递任务

		Topics []*ClsTopicInfo `json:"Topics,omitempty" name:"Topics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsDeliverTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsDeliverTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子网列表

		Data []*SubnetInfo `json:"Data,omitempty" name:"Data"`
		// 子网总个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;修改成功,&nbsp;!0:&nbsp;修改失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBetaTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBetaTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBetaTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNotifyTypeRequest struct {
	*tchttp.BaseRequest

	// 修改的对象

	Info *string `json:"Info,omitempty" name:"Info"`
}

func (r *ModifyNotifyTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotifyTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAclRuleHitTimesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重置成功后返回被重置的策略uuid

		RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAclRuleHitTimesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAclRuleHitTimesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBorderRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 预操作增加规则列表个数

	AddCnt *int64 `json:"AddCnt,omitempty" name:"AddCnt"`
	// 预操作类型，rollback：回滚；cover_import_in：覆盖入向导入；cover_import_out：覆盖出向导入；其他：add

	OprTypeFrom *string `json:"OprTypeFrom,omitempty" name:"OprTypeFrom"`
}

func (r *DescribeBorderRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFlowCenterLogsV1Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*VpcFlowCenterData `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFlowCenterLogsV1Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFlowCenterLogsV1Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Data []*NetFlowDomainTopInfo `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetFlowDomainTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleChangeItem struct {

	// Ip版本，0：IPv4，1：IPv6，默认为IPv4

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 新的sequence&nbsp;值

	NewOrderIndex *int64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
	// 原始sequence&nbsp;值

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
}

type CreateBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据

	CoverDuplicate *int64 `json:"CoverDuplicate,omitempty" name:"CoverDuplicate"`
	// 规则类型，1封禁，2放通，不支持域名封禁

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则列表

	Rules []*IntrusionDefenseRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateBlockIgnoreRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBlockIgnoreRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstanceNameId struct {

	// ccn接入实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// ccn接入实例Name

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type NormalRouteBackupInfo struct {

	// 关联子网列表信息

	AttachSubnets []*SubnetInfo `json:"AttachSubnets,omitempty" name:"AttachSubnets"`
	// 对象名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 下一跳备份信息，默认为0.0.0.0/0的下一跳

	NextHopInfo []*RouteNextHopInfo `json:"NextHopInfo,omitempty" name:"NextHopInfo"`
	// 对象id，即路由表id

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 是否可编辑，0：可编辑；1：不可编辑

	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 状态值，ACTIVE：有效的；NEXTHOP_INVALID：下一跳失效

	State *string `json:"State,omitempty" name:"State"`
	// 失效状态时，返回失效原因

	StateDesc *string `json:"StateDesc,omitempty" name:"StateDesc"`
}

type DescribeAccessControlOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折线图数据

		Data []*AccessControlOverviewData `json:"Data,omitempty" name:"Data"`
		// 入站数量

		InNum *uint64 `json:"InNum,omitempty" name:"InNum"`
		// 剩余配额

		LeftNum *uint64 `json:"LeftNum,omitempty" name:"LeftNum"`
		// 出站数量

		OutNum *uint64 `json:"OutNum,omitempty" name:"OutNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlackWhiteQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 封禁列表套餐规格

		BlockListVersionSpecNum *int64 `json:"BlockListVersionSpecNum,omitempty" name:"BlockListVersionSpecNum"`
		// 放通列表自定义白名单配额

		CustomRuleQuota *int64 `json:"CustomRuleQuota,omitempty" name:"CustomRuleQuota"`
		// 黑名单扩容配额

		ExpandBlackList *uint64 `json:"ExpandBlackList,omitempty" name:"ExpandBlackList"`
		// 最大导出黑名单

		MaxImportQuotaBlackList *uint64 `json:"MaxImportQuotaBlackList,omitempty" name:"MaxImportQuotaBlackList"`
		// 最大导出白名单

		MaxImportQuotaWhiteList *uint64 `json:"MaxImportQuotaWhiteList,omitempty" name:"MaxImportQuotaWhiteList"`
		// 黑名单需要配额

		NeedQuotaBlackList *uint64 `json:"NeedQuotaBlackList,omitempty" name:"NeedQuotaBlackList"`
		// 剩余通用扩展规则配额

		RemainExtensionNum *int64 `json:"RemainExtensionNum,omitempty" name:"RemainExtensionNum"`
		// 黑名单剩余配额

		RemainQuotaBlackList *uint64 `json:"RemainQuotaBlackList,omitempty" name:"RemainQuotaBlackList"`
		// 白名单剩余配额

		RemainQuotaWhiteList *uint64 `json:"RemainQuotaWhiteList,omitempty" name:"RemainQuotaWhiteList"`
		// 0成功&nbsp;非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功&nbsp;其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 黑名单总配额

		TotalQuotaBlackList *uint64 `json:"TotalQuotaBlackList,omitempty" name:"TotalQuotaBlackList"`
		// 白名单总配额

		TotalQuotaWhiteList *uint64 `json:"TotalQuotaWhiteList,omitempty" name:"TotalQuotaWhiteList"`
		// 黑名单已使用配额

		UsedBlackList *uint64 `json:"UsedBlackList,omitempty" name:"UsedBlackList"`
		// 互联网边界入站封禁列表数量

		UsedInBlackList *int64 `json:"UsedInBlackList,omitempty" name:"UsedInBlackList"`
		// 互联网边界出站封禁列表数量

		UsedOutBlackList *int64 `json:"UsedOutBlackList,omitempty" name:"UsedOutBlackList"`
		// 白名单已使用配额

		UsedWhiteList *uint64 `json:"UsedWhiteList,omitempty" name:"UsedWhiteList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlackWhiteQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlackWhiteQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpConnectionsRequest struct {
	*tchttp.BaseRequest

	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙组id

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// Vpc间防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeVpcIpConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcIpConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockInfo struct {

	// appid信息

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 拦截频率

	AvgCount *float64 `json:"AvgCount,omitempty" name:"AvgCount"`
	// 1是&nbsp;0不是

	Block *int64 `json:"Block,omitempty" name:"Block"`
	// 来源&nbsp;&nbsp;1拦截列表，2虚拟补丁，3威胁情报，4基础防御

	BlockSource *int64 `json:"BlockSource,omitempty" name:"BlockSource"`
	// 1是&nbsp;0不是

	ByPass *int64 `json:"ByPass,omitempty" name:"ByPass"`
	// 次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 国加

	Country *string `json:"Country,omitempty" name:"Country"`
	// 方向&nbsp;&nbsp;1入站&nbsp;0出战

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// vpc间防火墙日志，目的资产id

	DstAssetId *string `json:"DstAssetId,omitempty" name:"DstAssetId"`
	// vpc间防火墙日志，目的资产名称

	DstAssetName *string `json:"DstAssetName,omitempty" name:"DstAssetName"`
	// 目的ip

	DstIP *string `json:"DstIP,omitempty" name:"DstIP"`
	// 端口

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// vpc间防火墙日志，目的VPC&nbsp;id

	DstVpc *string `json:"DstVpc,omitempty" name:"DstVpc"`
	// vpc间防火墙日志，目的VPC名称

	DstVpcName *string `json:"DstVpcName,omitempty" name:"DstVpcName"`
	// vpc间防火墙开关边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// vpc间防火墙开关边名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 是否加入隔离列表，1是&nbsp;0不是

	Hide *int64 `json:"Hide,omitempty" name:"Hide"`
	// 1是&nbsp;0不是

	Ignore *int64 `json:"Ignore,omitempty" name:"Ignore"`
	// 置顶位

	IsTop *int64 `json:"IsTop,omitempty" name:"IsTop"`
	// 是否加入隔离列表，1是&nbsp;0不是

	Isolate *int64 `json:"Isolate,omitempty" name:"Isolate"`
	// 日志来源&nbsp;move：vpc间防火墙

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 最晚时间

	MaxDate *string `json:"MaxDate,omitempty" name:"MaxDate"`
	// 最近时间

	MinDate *string `json:"MinDate,omitempty" name:"MinDate"`
	// 源IP

	SrcIP *string `json:"SrcIP,omitempty" name:"SrcIP"`
	// vpc间防火墙日志，源VPC&nbsp;id

	SrcVpc *string `json:"SrcVpc,omitempty" name:"SrcVpc"`
	// vpc间防火墙日志，源VPC名称

	SrcVpcName *string `json:"SrcVpcName,omitempty" name:"SrcVpcName"`
	// 记录唯一标识

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// WhiteTag信息

	WhiteTag *string `json:"WhiteTag,omitempty" name:"WhiteTag"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type VpcFwJoinInstance struct {

	// 0-非sase实例，忽略，1-未绑定状态，2-已绑定

	AttachWithEdge *int64 `json:"AttachWithEdge,omitempty" name:"AttachWithEdge"`
	// 所属网络ID

	BelongNet *string `json:"BelongNet,omitempty" name:"BelongNet"`
	// 所属网络名称

	BelongNetName *string `json:"BelongNetName,omitempty" name:"BelongNetName"`
	// 网络实例内网网段

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 对端防火墙和开关状态&nbsp;0：正常，&nbsp;1：对等未创建防火墙，2：对等已创建防火墙，未打开开关

	CrossEdgeStatus *int64 `json:"CrossEdgeStatus,omitempty" name:"CrossEdgeStatus"`
	// 网络实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 网络实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 网络实例类型&nbsp;目前就vpc

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 网络实例内网IPV6网段

	Ipv6Cidr *string `json:"Ipv6Cidr,omitempty" name:"Ipv6Cidr"`
	// 网络实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 网络实例所在地域中文

	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`
}

type DescribeBetaTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// acl规则数

		AclNum *int64 `json:"AclNum,omitempty" name:"AclNum"`
		// 任务类型&nbsp;&nbsp;1border&nbsp;2nat&nbsp;3vpc&nbsp;4sg

		AclType *int64 `json:"AclType,omitempty" name:"AclType"`
		// 动作集合

		TaskActionList []*BetaTaskAction `json:"TaskActionList,omitempty" name:"TaskActionList"`
		// 任务详情

		TaskInfo *AutoAssistantUserConf `json:"TaskInfo,omitempty" name:"TaskInfo"`
		// 任务名称

		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBetaTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// nat实例唯一ID

		NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
		// 0&nbsp;正常
		// -1&nbsp;异常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowTimesRequest struct {
	*tchttp.BaseRequest

	// 时间类型&nbsp;1&nbsp;小时/2&nbsp;7天/3&nbsp;30天/4&nbsp;180天

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeFlowTimesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowTimesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetInfo struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产地域

	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 公网ip

	PublicIp []*string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type DescribeOperateLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作行为列表

		ActionLists []*string `json:"ActionLists,omitempty" name:"ActionLists"`
		// 操作日志列表

		Data []*OperatorLogsData `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 防火墙类型

		TypeLists []*string `json:"TypeLists,omitempty" name:"TypeLists"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperateLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyAuthorityRequest struct {
	*tchttp.BaseRequest

	// 策略状态&nbsp;&nbsp;0正常&nbsp;其他异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPolicyAuthorityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyAuthorityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideUserStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuideUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEdgeIpSwitchWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZeroTrustBlockStatusRequest struct {
	*tchttp.BaseRequest

	// 零信任未授权封禁按钮
	// 1&nbsp;打开
	// 0&nbsp;关闭

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 模板ID

	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *ModifyZeroTrustBlockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZeroTrustBlockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HoneyPotMergeTemplate struct {

	// 1&nbsp;允许修改
	// 0&nbsp;禁止修改

	AllowModify *int64 `json:"AllowModify,omitempty" name:"AllowModify"`
	// 蜜罐使用负载均衡时默认路径

	DefaultPath *string `json:"DefaultPath,omitempty" name:"DefaultPath"`
	// 自定义诱饵项

	DefineBait []*string `json:"DefineBait,omitempty" name:"DefineBait"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 默认暴露的端口

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 模版ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type IOAFilter struct {

	// 过滤字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 过滤方式&nbsp;eq:等于,net:不等于,like,nlike,gt:大于,lt:小于,egt:大于等于,elt:小于等于

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤条件

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ModifyWidthCheckRequest struct {
	*tchttp.BaseRequest

	// ew：东西向防火墙；nat：nat防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 带宽对比信息

	WidthCompareInfo []*WidthCompare `json:"WidthCompareInfo,omitempty" name:"WidthCompareInfo"`
}

func (r *ModifyWidthCheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWidthCheckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaitTemplate struct {

	// 诱饵

	Bait *string `json:"Bait,omitempty" name:"Bait"`
	// 蜜罐模版详细信息

	Template *HoneyPotMergeTemplate `json:"Template,omitempty" name:"Template"`
}

type DescribeGuideUserStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新手引导完成标志位

		Data *GuideUserStatus `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuideUserStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTiWhiteListRequest struct {
	*tchttp.BaseRequest

	// 页高

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTiWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefineBaitShow struct {

	// 自定义诱饵

	DefineBait *string `json:"DefineBait,omitempty" name:"DefineBait"`
	// Selector&nbsp;下拉框，&nbsp;OptionButton&nbsp;单选框

	Type *string `json:"Type,omitempty" name:"Type"`
	// 自定义诱饵值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type PortProtect struct {

	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 是否进行了奖励

	IsGive *bool `json:"IsGive,omitempty" name:"IsGive"`
	// 是否付费用户

	IsPay *bool `json:"IsPay,omitempty" name:"IsPay"`
	// 是否正在扫描

	IsScan *bool `json:"IsScan,omitempty" name:"IsScan"`
	// 暴露端口

	PortNum *int64 `json:"PortNum,omitempty" name:"PortNum"`
	// 防护的暴露端口

	ProPortNum *int64 `json:"ProPortNum,omitempty" name:"ProPortNum"`
	// 0&nbsp;&nbsp;扫描中&nbsp;1没有扫描&nbsp;2&nbsp;扫秒完成.

	ScanStatus *int64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type ZeroTrustAccessOverview struct {

	// 其他接入统计

	AccessSourceList []*ZeroTrustAccessSourceStat `json:"AccessSourceList,omitempty" name:"AccessSourceList"`
	// 已接入域名

	DomainCount *uint64 `json:"DomainCount,omitempty" name:"DomainCount"`
	// 显示接入的企业名称，如果未接入显示未接

	IOA *string `json:"IOA,omitempty" name:"IOA"`
	// 开启地域

	Region []*string `json:"Region,omitempty" name:"Region"`
	// 已接入vpc数量

	VpcCount *uint64 `json:"VpcCount,omitempty" name:"VpcCount"`
	// 微信身份

	WeChatCount *uint64 `json:"WeChatCount,omitempty" name:"WeChatCount"`
	// web域名

	WebDomainCount *uint64 `json:"WebDomainCount,omitempty" name:"WebDomainCount"`
}

type ModifyAutoInternetSwitchRequest struct {
	*tchttp.BaseRequest

	// 修改开关状态，1：开启，0：关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoInternetSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoInternetSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwAZoneRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 迁移主或备机的可用区，main：迁移主机可用区；backup：迁移备机可用区

	Role *string `json:"Role,omitempty" name:"Role"`
	// 迁移目标可用区

	SwitchAZone *string `json:"SwitchAZone,omitempty" name:"SwitchAZone"`
}

func (r *ModifyFwAZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwAZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwInsListRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCfwInsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperatorLogsRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 游标

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeOperatorLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperatorLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrictModeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStrictModeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrictModeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 攻击拦截次数

		AttackNum *int64 `json:"AttackNum,omitempty" name:"AttackNum"`
		// 蜜罐欺骗次数

		DeceptionNum *uint64 `json:"DeceptionNum,omitempty" name:"DeceptionNum"`
		// 网络访问次数

		FlowNum *int64 `json:"FlowNum,omitempty" name:"FlowNum"`
		// 攻击告警次数

		GiveNum *int64 `json:"GiveNum,omitempty" name:"GiveNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChooseResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产集合

	FilterList []*string `json:"FilterList,omitempty" name:"FilterList"`
	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 新组id

	NewGroupId *string `json:"NewGroupId,omitempty" name:"NewGroupId"`
	// 操作&nbsp;0移动到其他分组NewGroupId必填，2删除，1资产组资产变更FilterList为最终的资产

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
}

func (r *ModifyChooseResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChooseResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest

	// DNS&nbsp;开关切换列表

	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitempty" name:"DnsVpcSwitchLst"`
	// nat&nbsp;防火墙&nbsp;id

	NatFwInsId *string `json:"NatFwInsId,omitempty" name:"NatFwInsId"`
}

func (r *ModifyNatFwVpcDnsSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwVpcDnsSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEventListResult struct {

	// 无

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// 无

	List []*GetEventListStruct `json:"List,omitempty" name:"List"`
	// 无

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 无

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 无

	SearchLabelLabel []*string `json:"SearchLabelLabel,omitempty" name:"SearchLabelLabel"`
	// 无

	Title *string `json:"Title,omitempty" name:"Title"`
	// 无

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 无

	Url *string `json:"Url,omitempty" name:"Url"`
}

type ModifyAclApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAclApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAclApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSessionConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 会话连接数统计列表

		Data []*SessionConnections `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSessionConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSessionConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStorageSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStorageSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchErrorIgnoreRequest struct {
	*tchttp.BaseRequest

	// 错误实例

	ErrIns *string `json:"ErrIns,omitempty" name:"ErrIns"`
	// 错误类型&nbsp;&nbsp;ERR_ENDPOINT_SUBNET&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	ERR_ENDPOINT_ERR_CLOSE&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	ERR_ENDPOINT_ERR_TO_BYPASS&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	ERR_SERIAL_OVER_FLOW_BYPASS&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	ERR_SERIAL_OVER_FLOW_BYPASS_RULE&nbsp;	ERR_SERIAL_OVER_FLOW_CLOSE&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	ERR_SERIAL_OVER_FLOW_WARN&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;

	ErrKey *string `json:"ErrKey,omitempty" name:"ErrKey"`
}

func (r *ModifySwitchErrorIgnoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchErrorIgnoreRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatGwInfo struct {

	// nat网关Eip列表

	Eips []*string `json:"Eips,omitempty" name:"Eips"`
	// nat网关id

	NatGwId *string `json:"NatGwId,omitempty" name:"NatGwId"`
	// nat网关名称

	NatGwName *string `json:"NatGwName,omitempty" name:"NatGwName"`
	// region对应的字符串

	Region *string `json:"Region,omitempty" name:"Region"`
	// vpc的CIDR

	VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`
	// vpc的Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type RiskAsset struct {

	// 标签名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 标签值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeAutoBackUpSettingListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*AutoBackUpSetting `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoBackUpSettingListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackUpSettingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreRulesImportProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入进度百分比

		Progress *int64 `json:"Progress,omitempty" name:"Progress"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态，0导入中，1导入已完成，2导入失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIgnoreRulesImportProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreRulesImportProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleLogDetailRequest struct {
	*tchttp.BaseRequest

	// 类型，0：互联网边界防火墙，1：VPC边界防火墙

	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
	// uuid值

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeRuleLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLogDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChooseAssetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*InstanceInfo `json:"Data,omitempty" name:"Data"`
		// 数据总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChooseAssetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChooseAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorDetailRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeErrorDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SolidProtectionSingleApplyTrial struct {

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 试用失败原因，成功为空

	TrialMessage *string `json:"TrialMessage,omitempty" name:"TrialMessage"`
	// 试用状态&nbsp;0失败&nbsp;1成功

	TrialStatus *int64 `json:"TrialStatus,omitempty" name:"TrialStatus"`
}

type ModifyNatFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcCfwWidthRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 弹性带宽

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 弹性开关&nbsp;1打开&nbsp;0关闭

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// 防火墙地域调整带宽信息

	FwDeploy []*FwDeploy `json:"FwDeploy,omitempty" name:"FwDeploy"`
	// ew：东西向防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 按量计费标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyVpcCfwWidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcCfwWidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteAccessStat struct {

	// 标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 数量

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type Tags struct {

	// 标签名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type VpcFwGroupInfo struct {

	// 云联网模式下，实例关联的云联网id

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// cdc专用集群场景时表示部署所属的cdc

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// cdc专用集群场景时表示cdc名称

	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`
	// 跨租户模式&nbsp;1管理员&nbsp;2单边&nbsp;0&nbsp;非跨租户

	CrossUserMode *string `json:"CrossUserMode,omitempty" name:"CrossUserMode"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// VPC防火墙实例卡片信息数组

	FwInstanceLst []*VpcFwInstanceInfo `json:"FwInstanceLst,omitempty" name:"FwInstanceLst"`
	// 防火墙组涉及到的开关个数

	FwSwitchNum *int64 `json:"FwSwitchNum,omitempty" name:"FwSwitchNum"`
	// auto&nbsp;:自动选择
	// 如果为网段，则为用户自定义&nbsp;192.168.0.0/20&nbsp;

	FwVpcCidr *string `json:"FwVpcCidr,omitempty" name:"FwVpcCidr"`
	// 模式&nbsp;1：CCN云联网模式；0：私有网络模式&nbsp;2:&nbsp;sase&nbsp;模式&nbsp;3：ccn&nbsp;高级模式&nbsp;4:&nbsp;私有网络(跨租户单边模式)

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 云联网模式下，当前实例是否需要开启重叠路由开关，1：需要开启，0：不需要开启

	NeedSwitchCcnOverlap *int64 `json:"NeedSwitchCcnOverlap,omitempty" name:"NeedSwitchCcnOverlap"`
	// 防火墙(组)部署的地域

	RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
	// 防火墙(状态)&nbsp;0：正常&nbsp;1:&nbsp;初始化或操作中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 防火墙实例的开关模式&nbsp;1:&nbsp;单点互通&nbsp;2:&nbsp;多点互通&nbsp;3:&nbsp;全互通&nbsp;4:&nbsp;自定义路由

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

type DescribeIpsRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Data []*IpsRuleDetail `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpsRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcFwInstance struct {

	// 部署地域信息

	FwDeploy *FwDeploy `json:"FwDeploy,omitempty" name:"FwDeploy"`
	// 防火墙实例ID&nbsp;(编辑场景传)

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 防火墙实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 私有网络模式下接入的VpcId列表；仅私有网络模式使用

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
}

type DescribeLogsResultAsyncRequest struct {
	*tchttp.BaseRequest

	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询唯一Id

	QueryId *string `json:"QueryId,omitempty" name:"QueryId"`
	// 查询分片

	RangeField *string `json:"RangeField,omitempty" name:"RangeField"`
}

func (r *DescribeLogsResultAsyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsResultAsyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTimesAndFlowAssetMaxTopRequest struct {
	*tchttp.BaseRequest

	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 排序类型，0访问次数倒序，1流量大小倒序

	DataType *int64 `json:"DataType,omitempty" name:"DataType"`
	// 方向，0主动外联，1外部访问

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
	// 地域类型，0全球，1中国

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeVisitTimesAndFlowAssetMaxTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVisitTimesAndFlowAssetMaxTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFlowCenterLogsRequest struct {
	*tchttp.BaseRequest

	// 关系&nbsp;id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 端口号

	Port *string `json:"Port,omitempty" name:"Port"`
	// 源地址

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 目标地址

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeVpcFlowCenterLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFlowCenterLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceGroupOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceGroupOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceGroupOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckNatFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态返回，true为冲突，false不冲突

		Status *bool `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckNatFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNatFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackUpSettingListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAutoBackUpSettingListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoBackUpSettingListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZeroTrustBlockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZeroTrustBlockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZeroTrustBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleDetail struct {

	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，取值：0，出站；1，入站，vpc间规则时无方向默认为0

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 访问目的名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 访问目的类型，1是ip，2是url，3是http域名

	DstType *int64 `json:"DstType,omitempty" name:"DstType"`
	// 规则生效的范围，仅vpc间规则使用，是在哪对vpc之间还是针对所有vpc间生效

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// EdgeId对应的这对VPC间防火墙的描述，仅vpc规则使用

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// 规则顺序编号

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 协议端口组唯一ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 协议端口组名称

	ParamTemplateName *string `json:"ParamTemplateName,omitempty" name:"ParamTemplateName"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：
	// TCP
	// UDP
	// ICMP
	// ANY
	// HTTP
	// HTTPS
	// HTTP/HTTPS
	// SMTP
	// SMTPS
	// SMTP/SMTPS
	// FTP
	// DNS
	// TLS/SSL

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 规则对应的唯一id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 访问源名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 源类型，1是ip，2是url，3是http域名

	SrcType *int64 `json:"SrcType,omitempty" name:"SrcType"`
	// 规则被删除：1，已删除；0，未删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// 0：观察
	// 1：拒绝
	// 2：放行

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.domain.com

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
}

type DescribeOfflineExportTemporaryCredentialsRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeOfflineExportTemporaryCredentialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOfflineExportTemporaryCredentialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*RuleLogsData `json:"Data,omitempty" name:"Data"`
		// 是否到底了

		IsEnd *bool `json:"IsEnd,omitempty" name:"IsEnd"`
		// 游标

		Offset *string `json:"Offset,omitempty" name:"Offset"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCommonStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCommonStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCommonStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeCommonNewRequest struct {
	*tchttp.BaseRequest

	// 告警配置,用于外部入侵、主机失陷、网络蜜罐

	AlertConfig *AlertConfig `json:"AlertConfig,omitempty" name:"AlertConfig"`
	// 告警来源配置

	AlertSourceConfig *AlertSourceConfig `json:"AlertSourceConfig,omitempty" name:"AlertSourceConfig"`
	// 告警时间范围列表

	AlertTimeRangeList []*AlertTimeRange `json:"AlertTimeRangeList,omitempty" name:"AlertTimeRangeList"`
	// 勾选的暴露面类型列表

	ExposedSideList []*string `json:"ExposedSideList,omitempty" name:"ExposedSideList"`
	// 日志推送接收类型

	LogpushTypeInfo *LogpushTypeInfo `json:"LogpushTypeInfo,omitempty" name:"LogpushTypeInfo"`
	// 告警对象是否勾选主账号

	MasterAccount *int64 `json:"MasterAccount,omitempty" name:"MasterAccount"`
	// 请求操作的目标任务

	OperateTarget *string `json:"OperateTarget,omitempty" name:"OperateTarget"`
	// 是否刷新子账号

	Refresh *int64 `json:"Refresh,omitempty" name:"Refresh"`
	// 通知设置页面部分开关的状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 子账号数据列表

	SubAccountList []*SubAccount `json:"SubAccountList,omitempty" name:"SubAccountList"`
	// 标签设置的修改内容

	TagLog *TagLog `json:"TagLog,omitempty" name:"TagLog"`
	// 阈值

	Threshold *Threshold `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *ModifyNoticeCommonNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeCommonNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BanAndAllowRuleDel struct {

	// 0互联网出站&nbsp;1互联网入站&nbsp;5内网访问源&nbsp;6内网访问目的

	DirectionList *string `json:"DirectionList,omitempty" name:"DirectionList"`
	// 封禁和放通对象

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

type IntValueText struct {

	// 描述

	Text *string `json:"Text,omitempty" name:"Text"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeExportOperatorLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出地址

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 状态码

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportOperatorLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportOperatorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatConnectionsTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 峰值连接数

		ConnectionsMax *int64 `json:"ConnectionsMax,omitempty" name:"ConnectionsMax"`
		// 连接数配额占比

		ConnectionsPercent *float64 `json:"ConnectionsPercent,omitempty" name:"ConnectionsPercent"`
		// 连接数折线图

		ConnectionsTrendData []*FloatTrendsItem `json:"ConnectionsTrendData,omitempty" name:"ConnectionsTrendData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatConnectionsTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatConnectionsTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetflowBorderUsedRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNetflowBorderUsedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetflowBorderUsedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcProtectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC防火墙防护列表数据

		Data []*VpcProtectDetail `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcProtectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcProtectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescCdc struct {

	// 本地专用集群CDC&nbsp;Id

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 所属站点Id

	SiteId *string `json:"SiteId,omitempty" name:"SiteId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeAnalyzeLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果或redis缓存key

		Data *string `json:"Data,omitempty" name:"Data"`
		// 值&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 结果信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAnalyzeLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAnalyzeLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEsLogRequest struct {
	*tchttp.BaseRequest

	// es结构体

	Body *string `json:"Body,omitempty" name:"Body"`
	// 查询字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// ES查询索引:&nbsp;&nbsp;&nbsp;netflow_border_1300448058*

	Index *string `json:"Index,omitempty" name:"Index"`
	// 查询分类

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeEsLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEsLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// NAT地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// Edge&nbsp;ID值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 批量操作时的ID

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 状态，0：全部停用，1：全部启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAllRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwGroupAndInstance struct {

	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 关联的防火墙实例ID列表

	FwInsId []*string `json:"FwInsId,omitempty" name:"FwInsId"`
}

type SaveAutoBackUpSetting struct {

	// 周期：
	// day：每天
	// week：每周
	// month：每月

	Period *string `json:"Period,omitempty" name:"Period"`
	// 地域，ap-NewYork

	Region *string `json:"Region,omitempty" name:"Region"`
	// 1:&nbsp;开启，&nbsp;0:&nbsp;失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 子周期：
	// Period为day时不用填
	// Period为week时填1～7，对应周一到周日
	// Period为month时填1-31，对应每月的日期，31表示每月最后一天

	SubPeriod *string `json:"SubPeriod,omitempty" name:"SubPeriod"`
	// 时间：12点0分0秒执行

	Time *string `json:"Time,omitempty" name:"Time"`
	// 备份类型：
	// 1:&nbsp;互联网规则
	// 2:&nbsp;nat规则
	// 3:&nbsp;vpc规则

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteBetaTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态&nbsp;0成功&nbsp;1失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBetaTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBetaTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupAssociationRulesRequest struct {
	*tchttp.BaseRequest

	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
}

func (r *DescribeResourceGroupAssociationRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupAssociationRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户vpc信息列表

		Data []*FwVpcInfo `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClsDeliverSwitchRequest struct {
	*tchttp.BaseRequest

	// 修改模式；0修改指定分区，1全开启，2全关闭

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// Mode为0时，指定topic

	Topics []*ClsTopicInfo `json:"Topics,omitempty" name:"Topics"`
}

func (r *ModifyClsDeliverSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClsDeliverSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIocFeedbackRequest struct {
	*tchttp.BaseRequest

	// Ioc误报IP&nbsp;Description&nbsp;描述

	IocInfo []*IOCFeedBack `json:"IocInfo,omitempty" name:"IocInfo"`
}

func (r *AddIocFeedbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIocFeedbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeElasticBandWidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 边界弹性开关

		BorderElasticSwitch *int64 `json:"BorderElasticSwitch,omitempty" name:"BorderElasticSwitch"`
		// 弹性带宽区间

		Data []*ElasticBandWidth `json:"Data,omitempty" name:"Data"`
		// nat弹性开关

		NatElasticSwitch *int64 `json:"NatElasticSwitch,omitempty" name:"NatElasticSwitch"`
		// 是否为nat弹性白名单&nbsp;1是&nbsp;0不是

		NatWhiteUser *int64 `json:"NatWhiteUser,omitempty" name:"NatWhiteUser"`
		// vpc弹性开关

		VpcElasticSwitch *int64 `json:"VpcElasticSwitch,omitempty" name:"VpcElasticSwitch"`
		// 是否为vpc弹性白名单&nbsp;1是&nbsp;0不是

		VpcWhiteUser *int64 `json:"VpcWhiteUser,omitempty" name:"VpcWhiteUser"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeElasticBandWidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeElasticBandWidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventAssetInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资产信息

		AssetInfo *AssetInfo `json:"AssetInfo,omitempty" name:"AssetInfo"`
		// 高危端口Top5

		BarData []*BarData `json:"BarData,omitempty" name:"BarData"`
		// 安全事件统计

		PieData []*PieData `json:"PieData,omitempty" name:"PieData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventAssetInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessDomainCreateInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 0未创建&nbsp;1已创建&nbsp;2创建中

	IsCreateDomain *int64 `json:"IsCreateDomain,omitempty" name:"IsCreateDomain"`
	// 0未创建&nbsp;1已创建&nbsp;2创建中

	IsCreateNat *int64 `json:"IsCreateNat,omitempty" name:"IsCreateNat"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type BetaTaskAction struct {

	// 动作&nbsp;&nbsp;0关闭,&nbsp;1开启

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 预留字段处理特殊异常

	Exception *string `json:"Exception,omitempty" name:"Exception"`
	// 执行次数

	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`
	// 任务异常状态:&nbsp;0无意义&nbsp;1任务异常&nbsp;2无可执行任务对象

	FailedCode *int64 `json:"FailedCode,omitempty" name:"FailedCode"`
	// 周期&nbsp;&nbsp;day/week/month

	Period *string `json:"Period,omitempty" name:"Period"`
	// 顺序

	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`
	// 状态&nbsp;&nbsp;0待执行&nbsp;1正常&nbsp;2异常&nbsp;3已执行

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 子周期&nbsp;&nbsp;1-31/1-7&nbsp;

	SubPeriod *string `json:"SubPeriod,omitempty" name:"SubPeriod"`
	// 1临时任务&nbsp;2周期任务&nbsp;0未定义

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 执行日期&nbsp;&nbsp;临时任务yyyy-mm-dd&nbsp;&nbsp;hh:mm:ss&nbsp;格式&nbsp;&nbsp;周期任务&nbsp;hh:mm:ss&nbsp;格式

	Time *string `json:"Time,omitempty" name:"Time"`
}

type RuleRegionCount struct {

	// 规则数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeAssetScanStatusRequest struct {
	*tchttp.BaseRequest

	// json字符串

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeAssetScanStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCidrRelatedInstancesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的网段

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 开关边id：all，所有edgeId；cfws-xxx，指定边id关联的实例中查找

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 防火墙id：all，所有；cfwg-xxx，指定防火墙组关联的实例中查找

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
}

func (r *DescribeCidrRelatedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCidrRelatedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainParseIpLstRequest struct {
	*tchttp.BaseRequest

	// 需要解析的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防火墙类型：nat，nat防火墙；ew，vpc间防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// Ip版本，0：IPv4，1：IPv6，默认为IPv4

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *DescribeDomainParseIpLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDomainParseIpLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplayUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取回放链接

		Url *string `json:"Url,omitempty" name:"Url"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplayUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplayUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOCFeedBack struct {

	// 误报描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 误报IP

	Ioc *string `json:"Ioc,omitempty" name:"Ioc"`
	// 误报类型，0&nbsp;误报IP/域名，1告警规则误报

	IocType *int64 `json:"IocType,omitempty" name:"IocType"`
}

type DescribeErrorDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Data []*ErrorDetail `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwCrossStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;正常&nbsp;，&nbsp;1：对端账号VPC防火墙实例暂未创建，2：对端账号暂未开启当前防火墙开关，流量不会经过防火墙

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwCrossStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwCrossStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveAutoBackUpSettingRequest struct {
	*tchttp.BaseRequest

	// 参数列表

	Settings []*SaveAutoBackUpSetting `json:"Settings,omitempty" name:"Settings"`
}

func (r *SaveAutoBackUpSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveAutoBackUpSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSysLogTaskConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteSysLogTaskConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSysLogTaskConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwUserStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfwUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwEngineZoneListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFwEngineZoneListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEngineZoneListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInfoCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInfoCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInfoCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserStorageTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserStorageTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStorageTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperateLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值&nbsp;0正常其他异常

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperateLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatProtectListRequest struct {
	*tchttp.BaseRequest

	// 每页长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 获取NAT防护模式过滤参数

	SearchValue []*Filter `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeNatProtectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatProtectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 不算筛选条数的总条数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// 安全组规则列表数据

		Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data"`
		// 访问控制规则全部启用/全部停用

		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
		// 列表当前规则总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest

	// 规则

	Rule *BanAndAllowRule `json:"Rule,omitempty" name:"Rule"`
	// RuleType:&nbsp;1放通列表&nbsp;2外部IP&nbsp;3域名&nbsp;4情报&nbsp;5资产&nbsp;6自定义规则&nbsp;&nbsp;7入侵防御规则

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *ModifyBlockIgnoreRuleNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BorderACL struct {

	// 关联资产分组名称

	AssetGroupName *string `json:"AssetGroupName,omitempty" name:"AssetGroupName"`
	// 关联beta任务详情

	BetaList []*BetaInfoByACL `json:"BetaList,omitempty" name:"BetaList"`
	// 城市id

	CityCode *uint64 `json:"CityCode,omitempty" name:"CityCode"`
	// 城市名称

	CityName *string `json:"CityName,omitempty" name:"CityName"`
	// 云厂商

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 命中次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家id

	CountryCode *uint64 `json:"CountryCode,omitempty" name:"CountryCode"`
	// 国家名称

	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 规则有效性

	Invalid *uint64 `json:"Invalid,omitempty" name:"Invalid"`
	// 0为正常规则,1为地域规则

	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 告警规则id

	LogId *string `json:"LogId,omitempty" name:"LogId"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源类型,1是ip,2是url,3是http域名

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 1为启用状态,0为禁用状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略

	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的类型,1是ip,2是url,3是http域名,4.地域

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
	// 规则组uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type NDRAssetPublicIpInfo struct {

	// ip类型&nbsp;
	// 0&nbsp;未知,1&nbsp;公网,2&nbsp;弹性,3&nbsp;弹性ipv6,4&nbsp;anycastIP

	IpType *int64 `json:"IpType,omitempty" name:"IpType"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPBySGSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// true代表开启，false代表关闭

		Switch *bool `json:"Switch,omitempty" name:"Switch"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIPBySGSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPBySGSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageSetting struct {

	// 用户剩余清空日志次数

	CleanUpCount *int64 `json:"CleanUpCount,omitempty" name:"CleanUpCount"`
	// 上次清空时间

	LastCleanUpDate *string `json:"LastCleanUpDate,omitempty" name:"LastCleanUpDate"`
	// 上次设置存储类型时间

	LastLogTypeDate *string `json:"LastLogTypeDate,omitempty" name:"LastLogTypeDate"`
	// 上次设置时间

	LastStorageDate *string `json:"LastStorageDate,omitempty" name:"LastStorageDate"`
	// 用户当前存储时间

	StorageDay *int64 `json:"StorageDay,omitempty" name:"StorageDay"`
	// 用户时间参数可选列表&nbsp;[30,60,90]

	StorageDayList []*int64 `json:"StorageDayList,omitempty" name:"StorageDayList"`
}

type DescribeAssetSyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success&nbsp;成功
		// 其他失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回状态
		// 1&nbsp;更新中
		// 2&nbsp;更新完成
		// 3&nbsp;更新失败
		// 4&nbsp;更新失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBorderRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本授权规则数

		AllowedNum *uint64 `json:"AllowedNum,omitempty" name:"AllowedNum"`
		// 入站启用规则数

		InEnableNum *uint64 `json:"InEnableNum,omitempty" name:"InEnableNum"`
		// 入站规则数

		InTotal *uint64 `json:"InTotal,omitempty" name:"InTotal"`
		// 预操作是否超过配额，1：超过，0：未超过

		IsOverLimit *int64 `json:"IsOverLimit,omitempty" name:"IsOverLimit"`
		// 预操作需要额外配额数

		NeedQuota *int64 `json:"NeedQuota,omitempty" name:"NeedQuota"`
		// 出站启用规则数

		OutEnableNum *uint64 `json:"OutEnableNum,omitempty" name:"OutEnableNum"`
		// 出站规则数

		OutTotal *uint64 `json:"OutTotal,omitempty" name:"OutTotal"`
		// 剩余通用规则扩展数

		RemainExtensionNum *int64 `json:"RemainExtensionNum,omitempty" name:"RemainExtensionNum"`
		// 剩余配额数

		RemainQuota *int64 `json:"RemainQuota,omitempty" name:"RemainQuota"`
		// 已使用规格数

		UsedNum *uint64 `json:"UsedNum,omitempty" name:"UsedNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBorderRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBuyPageInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟补丁漏洞情况统计

		Leaks []*LeakTrendsHead `json:"Leaks,omitempty" name:"Leaks"`
		// 基础防御漏洞情况统计

		OtherAttacks []*LeakTrendsHead `json:"OtherAttacks,omitempty" name:"OtherAttacks"`
		// 防火墙防护情况统计

		ProtectInfo *CfwProtectInfo `json:"ProtectInfo,omitempty" name:"ProtectInfo"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBuyPageInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBuyPageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockListRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 是否处置&nbsp;1处置&nbsp;0不是

	ByPass *string `json:"ByPass,omitempty" name:"ByPass"`
	// 国家&nbsp;1国内

	Country *string `json:"Country,omitempty" name:"Country"`
	// 方向&nbsp;1&nbsp;入&nbsp;0出

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 日志来源&nbsp;move：vpc间防火墙

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 频率类型

	StatTimeSpan *int64 `json:"StatTimeSpan,omitempty" name:"StatTimeSpan"`
}

func (r *DescribeBlockListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 类型&nbsp;1&nbsp;告警&nbsp;2阻断

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 查询条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeTLogInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest

	// NAT地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 修改数据

	Data []*SequenceData `json:"Data,omitempty" name:"Data"`
	// 方向，0：出向，1：入向

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 边Id值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *ModifySequenceRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySequenceRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainModifyResult struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 修改成功与否状态值&nbsp;0：修改成功，非&nbsp;0：修改失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SysLogAddrConfig struct {

	// &nbsp;接入地址&nbsp;ip:port&nbsp;格式

	AccessAddr *string `json:"AccessAddr,omitempty" name:"AccessAddr"`
	// 接入方式&nbsp;[EnableSSL=0]11:udp&nbsp;12:tcp&nbsp;[EnableSSL=1]12:tls

	AccessMethod *int64 `json:"AccessMethod,omitempty" name:"AccessMethod"`
	// 是否使用tls加密&nbsp;1启用&nbsp;0不启用

	EnableSSL *int64 `json:"EnableSSL,omitempty" name:"EnableSSL"`
	// 预留后续投递&nbsp;kafka实例id

	KafkaId *string `json:"KafkaId,omitempty" name:"KafkaId"`
	// 预留后续投递&nbsp;kafka实例名称

	KafkaName *string `json:"KafkaName,omitempty" name:"KafkaName"`
	// 预留后续投递kafka地域

	KafkaRegion *string `json:"KafkaRegion,omitempty" name:"KafkaRegion"`
	// 预留后续投递&nbsp;kafkauin

	KafkaUin *string `json:"KafkaUin,omitempty" name:"KafkaUin"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 消息协议&nbsp;0:兼容协议&nbsp;1:RFC3164&nbsp;2:RFC5424

	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`
	// 连通性状态&nbsp;0:健康&nbsp;1:异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ModifySecurityGroupAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// NAT地域,&nbsp;云地域的英文简写

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// Edge&nbsp;ID值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 列表规则状态，0：全部停用，1：全部启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifySecurityGroupAllRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupAllRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessMasterDomainInfo struct {

	// 接入数据库

	AccessDataBaseCount *int64 `json:"AccessDataBaseCount,omitempty" name:"AccessDataBaseCount"`
	// 接入运维实例

	AccessInstanceCount *int64 `json:"AccessInstanceCount,omitempty" name:"AccessInstanceCount"`
	// 访问统计

	AccessStat []*RemoteAccessStat `json:"AccessStat,omitempty" name:"AccessStat"`
	// 接入的子域名详情

	AccessSubsetDomainInfo []*AccessSubsetDomainInfo `json:"AccessSubsetDomainInfo,omitempty" name:"AccessSubsetDomainInfo"`
	// 接入web服务

	AccessWebCount *int64 `json:"AccessWebCount,omitempty" name:"AccessWebCount"`
	// 创建状态

	CreateStatus *int64 `json:"CreateStatus,omitempty" name:"CreateStatus"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 失败原因

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 防火墙实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 防火墙实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 解析地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 域名最后修改时间：时间戳：秒

	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeCidrRelatedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详情信息

		RelateInstances []*FwRelateInstances `json:"RelateInstances,omitempty" name:"RelateInstances"`
		// 实例总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCidrRelatedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCidrRelatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProbeHistoryRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 条件过滤

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeProbeHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleLogDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命中日志详情

		Data []*RuleLogDetail `json:"Data,omitempty" name:"Data"`
		// VPC相应信息,保存为json字符串

		Info *string `json:"Info,omitempty" name:"Info"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatProbeEip struct {

	// 租户id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// 写入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 实例id

	NatInsID *string `json:"NatInsID,omitempty" name:"NatInsID"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态&nbsp;0待拨测&nbsp;&nbsp;1&nbsp;链路异常&nbsp;2健康&nbsp;&nbsp;3&nbsp;部分故障&nbsp;其他后续补充

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyGuideUserStatusRequest struct {
	*tchttp.BaseRequest

	// banport&nbsp;&nbsp;禁封端口&nbsp;scan资产扫描&nbsp;&nbsp;idps&nbsp;入侵防御

	StatusAction *string `json:"StatusAction,omitempty" name:"StatusAction"`
}

func (r *ModifyGuideUserStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGuideUserStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 编辑成功后返回新策略ID列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfwProtectInfo struct {

	// 资产防护次数

	AssetProtectNum *int64 `json:"AssetProtectNum,omitempty" name:"AssetProtectNum"`
	// 主机失陷次数

	CompromisedNum *int64 `json:"CompromisedNum,omitempty" name:"CompromisedNum"`
	// 系数

	EnLarge *float64 `json:"EnLarge,omitempty" name:"EnLarge"`
	// 拦截攻击次数

	InterceptAttackNum *int64 `json:"InterceptAttackNum,omitempty" name:"InterceptAttackNum"`
	// 主动外联次数

	OutBindVulNum *int64 `json:"OutBindVulNum,omitempty" name:"OutBindVulNum"`
	// 漏洞防护次数

	VulProtectNum *int64 `json:"VulProtectNum,omitempty" name:"VulProtectNum"`
}

type FlowCenterOutAsset struct {

	// 资产地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 会话数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 防火墙开关id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 请求结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是互联网还是nat

	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 访问流量请求

	RequestFlow *int64 `json:"RequestFlow,omitempty" name:"RequestFlow"`
	// 访问流量响应

	ResponseFlow *int64 `json:"ResponseFlow,omitempty" name:"ResponseFlow"`
	// 危险等级

	SenseLoss *string `json:"SenseLoss,omitempty" name:"SenseLoss"`
	// 请求开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type CancelUpdateEngineReserveTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelUpdateEngineReserveTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelUpdateEngineReserveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcFwInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcFwInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStorageLogTypeSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStorageLogTypeSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageLogTypeSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatProbeTrafficStatusRequest struct {
	*tchttp.BaseRequest

	// 事件id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 操作类型
	// ChangeIp&nbsp;切换主ip
	// 其他后续补充
	//

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SetNatProbeTrafficStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatProbeTrafficStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVpcAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功后返回新策略ID列表

		RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVpcAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeEdgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeEdgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeEdgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntrusionDefenseRule struct {

	// 备注信息，长度不能超过200

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则方向，0出站，1入站，3内网间

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 规则域名，IP与Domain必填其中之一

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则结束时间，格式：2006-01-02&nbsp;15:04:05，必须大于当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 规则IP地址，IP与Domain必填其中之一

	IP *string `json:"IP,omitempty" name:"IP"`
	// 规则开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type DescribeCfwCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vpc中防火墙使用的网段信息

		FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitempty" name:"FwCidrLst"`
		// 防火墙使用的网段类型，值VpcSelf/Assis/Custom分别代表自有网段优先/扩展网段优先/自定义

		FwCidrType *string `json:"FwCidrType,omitempty" name:"FwCidrType"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewFlowStatRequest struct {
	*tchttp.BaseRequest

	// 资产ID，VPC&nbsp;ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月&nbsp;5：6个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeOverviewFlowStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewFlowStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// NAT实例所在地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// NAT开关切换类型，1,单个子网，2，同开同关，3，全部

	ChangeType *int64 `json:"ChangeType,omitempty" name:"ChangeType"`
	// 选中的防火墙开关Id

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// type为NAT防火墙时需要。0：新增模式，1：接入模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0:&nbsp;互联网边界防火墙开关，1：vpc防火墙开关

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyAllSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVpcFwSupportSwitchModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例支持的开关模式

		SupportSwitchMode *SwitchModeSupport `json:"SupportSwitchMode,omitempty" name:"SupportSwitchMode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryVpcFwSupportSwitchModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVpcFwSupportSwitchModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleRequest struct {
	*tchttp.BaseRequest

	// 地域的英文简写

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DeleteSecurityGroupAllRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupAllRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据的json字符串

		Info *int64 `json:"Info,omitempty" name:"Info"`
		// 0:&nbsp;操作成功，非0：操作失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupAllRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupAllRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwSequenceRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwSequenceRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteRuleItem struct {

	// 用户APPID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 命令阻断acl&nbsp;id

	CommandGroupId *string `json:"CommandGroupId,omitempty" name:"CommandGroupId"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则是否开启，0未开启，1开启

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 访问目的示例：

	DstContent *string `json:"DstContent,omitempty" name:"DstContent"`
	// 访问目的类型，类型可以为：资产实例ins,WEB服务web,数据库db,资产分组asset_group,资源标签tag

	DstType *string `json:"DstType,omitempty" name:"DstType"`
	// 过期时间

	Expire *string `json:"Expire,omitempty" name:"Expire"`
	// 权限规则ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 规则有效性

	Invalid *int64 `json:"Invalid,omitempty" name:"Invalid"`
	// 访问源示例：

	SrcContent *string `json:"SrcContent,omitempty" name:"SrcContent"`
	// 访问源类型，类型可以为：用户user,用户组group,ip/cidr&nbsp;net

	SrcType *string `json:"SrcType,omitempty" name:"SrcType"`
	// 权限规则UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type SetNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEdgeStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcEdgeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEdgeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClsTopicInfo struct {

	// 投递状态&nbsp;1正常其他失败

	DeliverStatus *int64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 投递开关&nbsp;1开启&nbsp;0关闭

	DeliverSwitch *int64 `json:"DeliverSwitch,omitempty" name:"DeliverSwitch"`
	// 日志名称

	LogName *string `json:"LogName,omitempty" name:"LogName"`
	// 日志集ID

	LogSetID *string `json:"LogSetID,omitempty" name:"LogSetID"`
	// 日志集名称

	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ProductInfo struct {

	// 附注

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 记录序号

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否展示&nbsp;0不展示&nbsp;1展示

	LabelShow *int64 `json:"LabelShow,omitempty" name:"LabelShow"`
	// 链接信息

	Links []*Link `json:"Links,omitempty" name:"Links"`
	// 正文

	Notes *string `json:"Notes,omitempty" name:"Notes"`
	// 图片路径

	PicUrl *string `json:"PicUrl,omitempty" name:"PicUrl"`
	// tagText文案

	TagText *string `json:"TagText,omitempty" name:"TagText"`
	// tag类别

	TagType *int64 `json:"TagType,omitempty" name:"TagType"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type CreateNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id，该字段必须传递。

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 添加或删除操作的Dnat规则列表。

	DnatRules []*CfwNatDnatRule `json:"DnatRules,omitempty" name:"DnatRules"`
	// 0：cfw新增模式，1：cfw接入模式。

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *CreateNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProbeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关列表

		Data []*ProbeTaskDetail `json:"Data,omitempty" name:"Data"`
		// 拨测任务配额数

		Quota *int64 `json:"Quota,omitempty" name:"Quota"`
		// 剩余次数

		RemainCount *uint64 `json:"RemainCount,omitempty" name:"RemainCount"`
		// 列表总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProbeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwIpStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 单台cvm峰值带宽统计列表

		Data []*FwGroupIpFlowStat `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总行数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwIpStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwIpStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDefenseSwitchRequest struct {
	*tchttp.BaseRequest

	// 开关状态，1：打开，0：关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 按钮别名，基础防御basicruleswitch，安全基线baselineallswitch，虚拟补丁virtualpatchswitch，威胁情报tiswitch

	SwitchKey *string `json:"SwitchKey,omitempty" name:"SwitchKey"`
}

func (r *ModifyDefenseSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDefenseSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindEdgeCFWResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEdgeCFWResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindEdgeCFWResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertFeedBackListRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAlertFeedBackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertFeedBackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBetaTaskRequest struct {
	*tchttp.BaseRequest

	// 动作集合&nbsp;这里传递修改后的

	TaskActionList []*BetaTaskAction `json:"TaskActionList,omitempty" name:"TaskActionList"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 规则对象修改

	TaskInfo *AutoAssistantUserConf `json:"TaskInfo,omitempty" name:"TaskInfo"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *ModifyBetaTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBetaTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventNameListRequest struct {
	*tchttp.BaseRequest

	// 方向，1入站，0出站，为空不筛选

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 防火墙类型，0互联网，1vpc，2nat

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// 筛选攻击链

	KillChainList []*string `json:"KillChainList,omitempty" name:"KillChainList"`
}

func (r *DescribeEventNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BakRuleListObj struct {

	// 规则列表数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 实例信息,NAT时为region，VPC时为edge_id

	InsMsg *string `json:"InsMsg,omitempty" name:"InsMsg"`
	// 备份时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 规则列表名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型&nbsp;1互联网&nbsp;2nat&nbsp;3vpc&nbsp;4封禁名单&nbsp;5放通名单

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 规则列表uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AddZBTiNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 任务完成状态：0已完成&nbsp;1进行中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddZBTiNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddZBTiNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeAclApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnSASEEdgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SASE云防火墙实例数组

		SASEInstanceList []*SASEInstanceInfo `json:"SASEInstanceList,omitempty" name:"SASEInstanceList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnSASEEdgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnSASEEdgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙配额扩展数，可用于NAT防火墙或VPC间防火墙配额的扩展

		FwQuotaExtendCnt *uint64 `json:"FwQuotaExtendCnt,omitempty" name:"FwQuotaExtendCnt"`
		// 接入模式下，单实例最大接入nat网关数

		NatAccessNum *uint64 `json:"NatAccessNum,omitempty" name:"NatAccessNum"`
		// nat防火墙实例最大规格数

		NatFwInsLimit *uint64 `json:"NatFwInsLimit,omitempty" name:"NatFwInsLimit"`
		// nat新增模式支持最大关联弹性IP数

		NatHavipLimit *uint64 `json:"NatHavipLimit,omitempty" name:"NatHavipLimit"`
		// nat实例数剩余配额

		NatRemainNum *uint64 `json:"NatRemainNum,omitempty" name:"NatRemainNum"`
		// vpc防火墙组最大规格数

		VpcFwGroupLimit *uint64 `json:"VpcFwGroupLimit,omitempty" name:"VpcFwGroupLimit"`
		// vpc防火墙组剩余配额

		VpcFwGroupRemainNum *uint64 `json:"VpcFwGroupRemainNum,omitempty" name:"VpcFwGroupRemainNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePresetAddrTmplListRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePresetAddrTmplListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePresetAddrTmplListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// json字符串

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeVpcSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessProxyLog struct {

	// 访问类型

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 访问的服务所属appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 防火墙实例id

	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`
	// 访问详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 访问的web服务的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 访问的web实例id

	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`
	// 目标ip

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 结束时间（秒时间戳）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 源地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 访问源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 访问地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
	// 开始时间（秒时间戳）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 访问的用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 微信名称

	WxName *string `json:"WxName,omitempty" name:"WxName"`
}

type RegionInfo struct {

	// 地域key

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域详细信息

	RegionDetail *string `json:"RegionDetail,omitempty" name:"RegionDetail"`
	// 地域名称

	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`
}

type DescribeCosBucketListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储桶列表

		Data []*CosBucket `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCosBucketListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosBucketListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInstanceWithRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNatFwInstanceWithRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInstanceWithRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwConfigRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;不自动创建私有连接&nbsp;1&nbsp;自动创建私有连接

	AutoCreateEndpoint *int64 `json:"AutoCreateEndpoint,omitempty" name:"AutoCreateEndpoint"`
	// 0&nbsp;关闭&nbsp;1&nbsp;打开新增资产自动开启防火墙

	AutoDefence *int64 `json:"AutoDefence,omitempty" name:"AutoDefence"`
	// 0&nbsp;旁路&nbsp;1&nbsp;串行

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

func (r *ModifyFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 出站还是入站，0：入站，1：出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 规则是否开启，'0':&nbsp;未开启，'1':&nbsp;开启,&nbsp;默认为'0'

	Status *string `json:"Status,omitempty" name:"Status"`
	// 策略&nbsp;deny阻断&nbsp;

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeAcListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAcListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 自增id数组

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DescribeNatSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteNextHopLstRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeRouteNextHopLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteNextHopLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDnsResolveIpRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id或组id：nat防火墙时为实例id，形如cfwnat-xxx；vpc间防火墙时为防火墙组id，形如cfwg-xxx

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
}

func (r *DescribeDnsResolveIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDnsResolveIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTiContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回详情数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// 状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTiContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineCvmInfo struct {

	// 租户id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// cvmid

	CvmID *string `json:"CvmID,omitempty" name:"CvmID"`
	// 网卡名称

	Device *string `json:"Device,omitempty" name:"Device"`
	// 防火墙类型&nbsp;nat/vpc

	FsType *string `json:"FsType,omitempty" name:"FsType"`
	// 防火墙类型&nbsp;nat/vpc

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 机器ip

	Host *string `json:"Host,omitempty" name:"Host"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 用户名称

	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 队列数

	QueueNumber *string `json:"QueueNumber,omitempty" name:"QueueNumber"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type SecurityGroupVersionInfo struct {

	// 升级规则后将会触发重新下发安全组，已经配置的企业安全组规则将会重新关联新增类型资产；建议关闭自动下发后更新规则，并预览变动内容确认后下发

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 发布时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 支持规则变动类型&nbsp;前端忽略

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 安全组下发资产对象新增支持&nbsp;MariaDB

	Title *string `json:"Title,omitempty" name:"Title"`
	// 版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeNatFwRuleHitDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则命中详情

		Data []*NatRuleDetail `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwRuleHitDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwRuleHitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineUpdateOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎更新信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// nat引擎更新状态

		NatStatus *int64 `json:"NatStatus,omitempty" name:"NatStatus"`
		// 版本更新时间

		Time *string `json:"Time,omitempty" name:"Time"`
		// 版本号

		Version *string `json:"Version,omitempty" name:"Version"`
		// vpc引擎更新状态

		VpcStatus *int64 `json:"VpcStatus,omitempty" name:"VpcStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEngineUpdateOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineUpdateOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartFwInsRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
}

func (r *RestartFwInsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartFwInsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrAclInfo struct {

	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 顺序

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 策略阻断

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// vpc和互联网需要&nbsp;

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 源内容

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 源类型

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 目的内容

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 目的类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type NatSubnetStatStruct struct {

	// 开关是否打开&nbsp;1开&nbsp;0关闭

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 入向峰值带宽

	InputMax *int64 `json:"InputMax,omitempty" name:"InputMax"`
	// 防火墙id

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// 出向峰值带宽

	OutputMax *int64 `json:"OutputMax,omitempty" name:"OutputMax"`
	// 子网地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 子网CIDR

	SubnetCidr *string `json:"SubnetCidr,omitempty" name:"SubnetCidr"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 访问控制日志存储量，单位B

		AclSize *int64 `json:"AclSize,omitempty" name:"AclSize"`
		// 后付费模式存储状态，0正常，1欠费停止写入

		ArrearsStopWriting *int64 `json:"ArrearsStopWriting,omitempty" name:"ArrearsStopWriting"`
		// 入侵防御日志存储量，单位B

		IdsSize *int64 `json:"IdsSize,omitempty" name:"IdsSize"`
		// 剩余存储量，单位B

		LeftSize *int64 `json:"LeftSize,omitempty" name:"LeftSize"`
		// 流量日志存储量，单位B

		NetFlowSize *int64 `json:"NetFlowSize,omitempty" name:"NetFlowSize"`
		// 操作日志存储量，单位B

		OperateSize *int64 `json:"OperateSize,omitempty" name:"OperateSize"`
		// 计费模式，0后付费，1预付费

		PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 存储天数

		StorageDay *int64 `json:"StorageDay,omitempty" name:"StorageDay"`
		// 每日增加日志存储量柱状图

		TimeHistogram []*StorageHistogram `json:"TimeHistogram,omitempty" name:"TimeHistogram"`
		// 柱形图格式数据

		TimeHistogramShow *StorageHistogramShow `json:"TimeHistogramShow,omitempty" name:"TimeHistogramShow"`
		// 配额存储总量，单位B

		TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 已使用存储量，单位B

		UsedSize *int64 `json:"UsedSize,omitempty" name:"UsedSize"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatExistRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
		// 地域状态列表

		RegionStatusList []*RegionStatusData `json:"RegionStatusList,omitempty" name:"RegionStatusList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatExistRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatExistRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwInfoCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前租户的nat实例个数

		NatFwInsCount *int64 `json:"NatFwInsCount,omitempty" name:"NatFwInsCount"`
		// 打开开关个数

		OpenSwitchCount *int64 `json:"OpenSwitchCount,omitempty" name:"OpenSwitchCount"`
		// 返回参数

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 当前租户接入子网个数

		SubnetCount *int64 `json:"SubnetCount,omitempty" name:"SubnetCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwInfoCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwInfoCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAssetInfoStruct struct {

	// 无

	Time *string `json:"Time,omitempty" name:"Time"`
	// 无

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type Paging struct {

	// 总页数(只支持32位)

	PageCount *int64 `json:"PageCount,omitempty" name:"PageCount"`
	// 页码(只支持32位)

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页条数(只支持32位)

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 记录总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type DescribeCfwBandWidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// nat流量详情

		NatDetail []*NatInstanceDetail `json:"NatDetail,omitempty" name:"NatDetail"`
		// 互联网边界带宽Mbps

		NetBandWidth *float64 `json:"NetBandWidth,omitempty" name:"NetBandWidth"`
		// 流量峰值Mbps

		NetMaxFlow *float64 `json:"NetMaxFlow,omitempty" name:"NetMaxFlow"`
		// 状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwBandWidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwBandWidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpcRuleHitTimesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重置成功后返回被重置的策略uuid

		RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpcRuleHitTimesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpcRuleHitTimesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchCommon struct {

	// 入向情报开关

	TiSwitchIn *string `json:"TiSwitchIn,omitempty" name:"TiSwitchIn"`
	// 出向情报开关

	TiSwitchOut *string `json:"TiSwitchOut,omitempty" name:"TiSwitchOut"`
}

type QueryFilterV3 struct {

	// 过滤条件

	Filter *QueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 有无子条件

	HasSub *bool `json:"HasSub,omitempty" name:"HasSub"`
	// 逻辑操作(只支持32位)

	Logic *uint64 `json:"Logic,omitempty" name:"Logic"`
	// 查询条件

	SubFilters []*QueryFilter `json:"SubFilters,omitempty" name:"SubFilters"`
}

type DescribeNatGwJoinFwStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙实例id

		FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
		// 0:未接入，1，已接入

		JoinFwStatus *int64 `json:"JoinFwStatus,omitempty" name:"JoinFwStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGwJoinFwStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGwJoinFwStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceGroupChangeImpactRequest struct {
	*tchttp.BaseRequest

	// 资源组变更类型

	ResourceGroupChangeType *int64 `json:"ResourceGroupChangeType,omitempty" name:"ResourceGroupChangeType"`
	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
	// 资源组名称，名称变更场景必填

	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`
	// 资源组新父id，移动场景必填

	ResourceGroupNowParentId *string `json:"ResourceGroupNowParentId,omitempty" name:"ResourceGroupNowParentId"`
}

func (r *GetResourceGroupChangeImpactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceGroupChangeImpactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetSyncRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyAssetSyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetSyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeStatus struct {

	// 边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 边名字

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 开关操作状态，0:未操作,1:开启中,2:关闭中

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SerialRegionInfo struct {

	// 弹性带宽上限，单位Mbps

	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
	// 弹性开关

	ElasticSwitch *int64 `json:"ElasticSwitch,omitempty" name:"ElasticSwitch"`
	// 七天入向峰值带宽，单位bps

	InFlowMax *int64 `json:"InFlowMax,omitempty" name:"InFlowMax"`
	// 七天出向峰值带宽，单位bps

	OutFlowMax *int64 `json:"OutFlowMax,omitempty" name:"OutFlowMax"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 分配带宽值&nbsp;单位M

	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type AddVpcAcRuleRequest struct {
	*tchttp.BaseRequest

	// 添加规则的来源，一般不需要使用，值insert_rule&nbsp;表示插入指定位置的规则；值batch_import&nbsp;表示批量导入规则；为空时表示添加规则

	From *string `json:"From,omitempty" name:"From"`
	// 需要添加的vpc内网间规则列表

	Rules []*VpcRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddVpcAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 类型&nbsp;add&nbsp;手动添加&nbsp;&nbsp;import&nbsp;&nbsp;模板导入

	AddType *string `json:"AddType,omitempty" name:"AddType"`
	// 名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 父资产组id

	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
	// 模板id

	TempId *string `json:"TempId,omitempty" name:"TempId"`
}

func (r *CreateResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProbeHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关列表

		Data []*ProbeTaskHistory `json:"Data,omitempty" name:"Data"`
		// 列表总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProbeHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchModeSupport struct {

	// 全互通是否支持,1:支持,0:不支持

	AllMode *int64 `json:"AllMode,omitempty" name:"AllMode"`
	// 自定义路由是否支持,1:支持,0:不支持

	DefineMode *int64 `json:"DefineMode,omitempty" name:"DefineMode"`
	// 单点互通是否支持,1:支持,0:不支持

	EdgeMode *int64 `json:"EdgeMode,omitempty" name:"EdgeMode"`
	// 多点互通是否支持,1:支持,0:不支持

	VpcMode *int64 `json:"VpcMode,omitempty" name:"VpcMode"`
}

type DescribeAreaStatusRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// nat防火墙实例id

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
}

func (r *DescribeAreaStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAreaStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwInsStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCfwInsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatFwSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT边界防火墙开关列表数据

		Data []*NatSwitchListData `json:"Data,omitempty" name:"Data"`
		// 失败开关状态

		FailData []*CommonIdNameStatus `json:"FailData,omitempty" name:"FailData"`
		// 开关相关NAT列表

		NatList []*CommonIdName `json:"NatList,omitempty" name:"NatList"`
		// 关闭开关个数

		OffNum *int64 `json:"OffNum,omitempty" name:"OffNum"`
		// 开启开关个数

		OnNum *int64 `json:"OnNum,omitempty" name:"OnNum"`
		// 开关相关ROUTE列表

		RouteList []*CommonIdName `json:"RouteList,omitempty" name:"RouteList"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 开关相关VPC列表

		VpcList []*CommonIdName `json:"VpcList,omitempty" name:"VpcList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeSwitchRecoveryRequest struct {
	*tchttp.BaseRequest

	// 忽略的ip列表

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *ModifyEdgeSwitchRecoveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeSwitchRecoveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeWarnTrendsRequest struct {
	*tchttp.BaseRequest

	// 类型，0：7天，1：24小时，2：近30天，3：6个月

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSafeWarnTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeWarnTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcSwitchListsRequest struct {
	*tchttp.BaseRequest

	// json字符串

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeVpcSwitchListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcSwitchListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelEipProbe struct {

	// EIP&nbsp;类型的探针的公网ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 探针ID

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 探针名称

	ProbeName *string `json:"ProbeName,omitempty" name:"ProbeName"`
}

type SASEInstanceInfo struct {

	// 带宽规格（Mbps）

	Bandwidth *string `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// Edge&nbsp;id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 七层协议，NTA日志有效

		AppProtocolList []*string `json:"AppProtocolList,omitempty" name:"AppProtocolList"`
		// 日志列表

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id，该字段必须传递。

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 0：cfw新增模式，1：cfw接入模式。

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 修改操作的新的Dnat规则

	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
	// 修改操作的原始Dnat规则

	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`
}

func (r *ModifyNatFwDnatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwDnatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCfwIdpsModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCfwIdpsModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCfwIdpsModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProbeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProbeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeElasticBandWidthRequest struct {
	*tchttp.BaseRequest

	// 防火墙的实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 防火墙类型：nat、vpc，或者空

	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

func (r *DescribeElasticBandWidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeElasticBandWidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePresetAddrTmplListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地址模版信息列表

		Data []*PresetAddressTemplateInfo `json:"Data,omitempty" name:"Data"`
		// 域名模版数

		DomainTemplateCount *int64 `json:"DomainTemplateCount,omitempty" name:"DomainTemplateCount"`
		// ip模版数

		IpTemplateCount *int64 `json:"IpTemplateCount,omitempty" name:"IpTemplateCount"`
		// 端口模版数

		PortTemplateCount *int64 `json:"PortTemplateCount,omitempty" name:"PortTemplateCount"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePresetAddrTmplListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePresetAddrTmplListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageLogTypeSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志存储类型配置详情

		Data *LogStorageTypeSetting `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageLogTypeSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageLogTypeSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetNatRuleHitTimesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重置成功后返回被重置的策略uuid

		RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetNatRuleHitTimesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetNatRuleHitTimesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcFwSwitch struct {

	// 0-非sase实例，忽略，1-未绑定状态，2-已绑定

	AttachWithEdge *int64 `json:"AttachWithEdge,omitempty" name:"AttachWithEdge"`
	// 连接ID

	ConnectId *string `json:"ConnectId,omitempty" name:"ConnectId"`
	// 云联网ccn防火墙开发--勿删

	ConnectName *string `json:"ConnectName,omitempty" name:"ConnectName"`
	// 开关边连接类型&nbsp;0：对等连接，&nbsp;1：云连网

	ConnectType *int64 `json:"ConnectType,omitempty" name:"ConnectType"`
	// 对等防火墙和开关状态&nbsp;0：正常，&nbsp;1：对等未创建防火墙，2：对等已创建防火墙，未打开开关

	CrossEdgeStatus *int64 `json:"CrossEdgeStatus,omitempty" name:"CrossEdgeStatus"`
	// 开关边的ID

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 开关边的名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 开关状态&nbsp;0：关&nbsp;，&nbsp;1：开

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// VPC防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// VPC防火墙实例名称

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// 网络经过VPC防火墙CVM所在地域

	FwInsRegion *string `json:"FwInsRegion,omitempty" name:"FwInsRegion"`
	// 本端网络实例Cidr

	SourceInstanceCidr *string `json:"SourceInstanceCidr,omitempty" name:"SourceInstanceCidr"`
	// 本端网络实例ID

	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	// 本端网络实例名称

	SourceInstanceName *string `json:"SourceInstanceName,omitempty" name:"SourceInstanceName"`
	// 本端网络实例地域

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 开关的状态&nbsp;0：正常，&nbsp;1：转换中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 对端网络实例Cidr

	TargetInstanceCidr *string `json:"TargetInstanceCidr,omitempty" name:"TargetInstanceCidr"`
	// 对端网络实例ID

	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	// 对端网络实例名称

	TargetInstanceName *string `json:"TargetInstanceName,omitempty" name:"TargetInstanceName"`
	// 对端网络实例地域

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
}

type DescribeNatFwRouteBackupLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份路由列表

		Data []*NormalRouteBackupInfo `json:"Data,omitempty" name:"Data"`
		// 下一跳类型列表

		NextHopTypeLst []*StringValueText `json:"NextHopTypeLst,omitempty" name:"NextHopTypeLst"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatFwRouteBackupLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwRouteBackupLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyslogConnectionStateRequest struct {
	*tchttp.BaseRequest

	// 投递地址配置

	AddrConfig *SysLogAddrConfig `json:"AddrConfig,omitempty" name:"AddrConfig"`
}

func (r *DescribeSyslogConnectionStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSyslogConnectionStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyLoginTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVpcFwConfigRequest struct {
	*tchttp.BaseRequest

	// 防火墙满载可持续时间，单位s，超过该时间将bypass防火墙，为0时不做任何操作

	OverLoadLimitTime *int64 `json:"OverLoadLimitTime,omitempty" name:"OverLoadLimitTime"`
}

func (r *SetVpcFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetVpcFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcRule struct {

	// 限速的类型，可以是具体的IP也可以是CIDR，比如192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 规则编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 入向限制速度，单位是kbps

	InLimit *uint64 `json:"InLimit,omitempty" name:"InLimit"`
	// 0&nbsp;IPV4,&nbsp;1&nbsp;IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// IP：IP限速
	// PORT：端口限速

	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`
	// 出向限制速度，单位是kbps

	OutLimit *uint64 `json:"OutLimit,omitempty" name:"OutLimit"`
	// 流量控制策略的端口。取值：
	// -1：全部端口
	// 80：80端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeAddressTemplateListRequest struct {
	*tchttp.BaseRequest

	// 排序字段，取值&nbsp;'UpdateTime'&nbsp;|&nbsp;'RulesNum'

	By *string `json:"By,omitempty" name:"By"`
	// 条数，分页用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，分页用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，取值&nbsp;'asc'|'desc'

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 是否展示已删除数据

	ShowDelete *bool `json:"ShowDelete,omitempty" name:"ShowDelete"`
	// 模板Id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 1：ip模板，5：域名模板，6：协议端口模板

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 检索地址模板唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAddressTemplateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGwListRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 地域列表；为空时查询所有nat网关

	RegionLst []*string `json:"RegionLst,omitempty" name:"RegionLst"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatGwListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGwListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpConnections struct {

	// 地理位置,IPType=2时有效

	Address *string `json:"Address,omitempty" name:"Address"`
	// 连接数

	Connections *int64 `json:"Connections,omitempty" name:"Connections"`
	// IP,IPType=1时为资产id，IPType=2时为外部IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 资产ID,IPType=1时有效

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产名称,IPType=1时有效

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Ip类型

	IpType *int64 `json:"IpType,omitempty" name:"IpType"`
	// ip版本,0,IPV4;1,IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 封禁状态,IPType=2时有效

	IsBlock *int64 `json:"IsBlock,omitempty" name:"IsBlock"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 限速状态,IPType=1时有效

	TcRuleStatus *int64 `json:"TcRuleStatus,omitempty" name:"TcRuleStatus"`
	// VPC&nbsp;id,IPType=1时有效

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC名称,IPType=1时有效

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeAssetFilterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*AssetFilterInfo `json:"Data,omitempty" name:"Data"`
		// 0

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetFilterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetFilterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwInstancesInfoRequest struct {
	*tchttp.BaseRequest

	// 获取实例列表过滤字段

	Filter []*DescribeFilter `json:"Filter,omitempty" name:"Filter"`
	// 模式&nbsp;1跨租户&nbsp;0正常

	IsPeer *int64 `json:"IsPeer,omitempty" name:"IsPeer"`
	// 每页长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcFwInstancesInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwInstancesInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertCenterIsolateRequest struct {
	*tchttp.BaseRequest

	// 当前日志方向：&nbsp;0&nbsp;出向&nbsp;1&nbsp;入向

	AlertDirection *int64 `json:"AlertDirection,omitempty" name:"AlertDirection"`
	// 处置对象,资产列表

	HandleAssetList []*string `json:"HandleAssetList,omitempty" name:"HandleAssetList"`
	// 处置时间
	// 1&nbsp;&nbsp;1天
	// 7&nbsp;&nbsp;&nbsp;7天
	// -2&nbsp;永久

	HandleTime *int64 `json:"HandleTime,omitempty" name:"HandleTime"`
	// 隔离类型&nbsp;
	// 1&nbsp;互联网入站
	// 2&nbsp;互联网出站
	// 4&nbsp;内网访问

	IsolateType []*int64 `json:"IsolateType,omitempty" name:"IsolateType"`
	// 运维模式&nbsp;1&nbsp;IP白名单&nbsp;2&nbsp;身份认证

	OmMode *int64 `json:"OmMode,omitempty" name:"OmMode"`
}

func (r *CreateAlertCenterIsolateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterIsolateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthBannerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关基础信息

		Data *BandWidthBanner `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandWidthBannerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandWidthBannerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideNetFlowSwitchAPIResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 互联网边界防火墙开关查询信息

		Data *NetFlowSwitchInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuideNetFlowSwitchAPIResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideNetFlowSwitchAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模块基础配置

		Data *ModuleConf `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewTopic struct {

	// 数据集合

	List []*GetOverviewTopicStruct `json:"List,omitempty" name:"List"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// topcid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
}

type DescribeAutoAssistantConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 初始配置

		Data *AutoAssistantConf `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoAssistantConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoAssistantConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTiContentRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	PocID *int64 `json:"PocID,omitempty" name:"PocID"`
}

func (r *DescribeTiContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatAcRuleSwitchAllRequest struct {
	*tchttp.BaseRequest

	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// delete&nbsp;：执行删除；不删除则置空

	IsDelete *string `json:"IsDelete,omitempty" name:"IsDelete"`
	// 是否开关开启，0：未开启，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyNatAcRuleSwitchAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAcRuleSwitchAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStorageLogTypeSettingRequest struct {
	*tchttp.BaseRequest

	// 是否清理历史流量日志，1否，2是

	CleanHistoryLog *int64 `json:"CleanHistoryLog,omitempty" name:"CleanHistoryLog"`
	// 租户日志存储类型配置

	LogStorageTypeSetting *LogStorageTypeSetting `json:"LogStorageTypeSetting,omitempty" name:"LogStorageTypeSetting"`
}

func (r *ModifyStorageLogTypeSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageLogTypeSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetListsData struct {

	// 资产实例ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产实例名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
}

type DescribeNatIpConnectionsRequest struct {
	*tchttp.BaseRequest

	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// IP类型，1外部IP，2资产IP

	IpType *int64 `json:"IpType,omitempty" name:"IpType"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// nat实例Id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatIpConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatIpConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetflowRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则方向，0：出站，1：入站，默认1

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 更改的规则当前执行顺序

	RuleSequence *uint64 `json:"RuleSequence,omitempty" name:"RuleSequence"`
	// 是否开关开启，0：未开启，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyNetflowRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetflowRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回多余的信息

		Info *string `json:"Info,omitempty" name:"Info"`
		// 状态值，0：成功，非0：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoInternetSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoInternetSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoInternetSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonFilter struct {

	// 查询key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 枚举类型，代表name与values之间的匹配关系&nbsp;enum&nbsp;FilterOperatorType&nbsp;{&nbsp;//INVALID&nbsp;FILTER_OPERATOR_TYPE_INVALID&nbsp;=&nbsp;0;&nbsp;//等于&nbsp;FILTER_OPERATOR_TYPE_EQUAL&nbsp;=&nbsp;1;&nbsp;//大于&nbsp;FILTER_OPERATOR_TYPE_GREATER&nbsp;=&nbsp;2;&nbsp;//小于&nbsp;FILTER_OPERATOR_TYPE_LESS&nbsp;=&nbsp;3;&nbsp;//大于等于&nbsp;FILTER_OPERATOR_TYPE_GREATER_EQ&nbsp;=&nbsp;4;&nbsp;//小于等于&nbsp;FILTER_OPERATOR_TYPE_LESS_EQ&nbsp;=&nbsp;5;&nbsp;//不等于&nbsp;FILTER_OPERATOR_TYPE_NO_EQ&nbsp;=&nbsp;6;&nbsp;//in，数组中包含&nbsp;FILTER_OPERATOR_TYPE_IN&nbsp;=&nbsp;7;&nbsp;//not&nbsp;in&nbsp;FILTER_OPERATOR_TYPE_NOT_IN&nbsp;=&nbsp;8;&nbsp;//模糊匹配&nbsp;FILTER_OPERATOR_TYPE_FUZZINESS&nbsp;=&nbsp;9;&nbsp;//存在&nbsp;FILTER_OPERATOR_TYPE_EXIST&nbsp;=&nbsp;10;&nbsp;//不存在&nbsp;FILTER_OPERATOR_TYPE_NOT_EXIST&nbsp;=&nbsp;11;&nbsp;//正则&nbsp;FILTER_OPERATOR_TYPE_REGULAR&nbsp;=&nbsp;12;&nbsp;}

	OperatorType *int64 `json:"OperatorType,omitempty" name:"OperatorType"`
	// 查询值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteBetaTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteBetaTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBetaTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarnEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表数据

		Data []*WarnEventData `json:"Data,omitempty" name:"Data"`
		// 安全事件类型名称列表

		EventNameList []*string `json:"EventNameList,omitempty" name:"EventNameList"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWarnEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarnEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeIpSwitchWeight struct {

	// ip&nbsp;开关

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 开关权重

	SwitchWeight *int64 `json:"SwitchWeight,omitempty" name:"SwitchWeight"`
}

type DescribeAlertApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeAlertApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDetectedInfo struct {

	// 资产个数

	AssetsCount *int64 `json:"AssetsCount,omitempty" name:"AssetsCount"`
	// 外部入侵检测次数

	InboundEventCount *int64 `json:"InboundEventCount,omitempty" name:"InboundEventCount"`
	// 主动外联管控次数

	OutboundEventCount *int64 `json:"OutboundEventCount,omitempty" name:"OutboundEventCount"`
	// 暴露端口个数

	PortCount *int64 `json:"PortCount,omitempty" name:"PortCount"`
	// 检测进度

	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
	// 检测时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// vpc个数

	VpcCount *int64 `json:"VpcCount,omitempty" name:"VpcCount"`
	// 漏洞个数

	VulCount *int64 `json:"VulCount,omitempty" name:"VulCount"`
}

type DescribeFlowCenterAssetListsRequest struct {
	*tchttp.BaseRequest

	// 资产实例&nbsp;id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 检索结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 检索条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 排序条件

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 检索开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFlowCenterAssetListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterAssetListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowCenterLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*FlowCenterData `json:"Data,omitempty" name:"Data"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowCenterLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWarningCenterOverviewRequest struct {
	*tchttp.BaseRequest

	// 类型，0：7天，1：24小时

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeWarningCenterOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWarningCenterOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateCategoryInfo struct {

	// 蜜罐诱饵

	Bait *string `json:"Bait,omitempty" name:"Bait"`
	// 1:允许快速创建，0:需要填写诱饵，不允许快速创建

	CanFastCreate *int64 `json:"CanFastCreate,omitempty" name:"CanFastCreate"`
	// 蜜罐所属分类列表

	Category *string `json:"Category,omitempty" name:"Category"`
	// 是否允许创建灰度

	CreateGrayscale *int64 `json:"CreateGrayscale,omitempty" name:"CreateGrayscale"`
	// 蜜罐描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 蜜罐展示排名先后

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 蜜罐服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 支持探针类型

	SupportDeployMode *string `json:"SupportDeployMode,omitempty" name:"SupportDeployMode"`
	// 蜜罐模版ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeSerialRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 旁路带宽数据

		BypassWidth *int64 `json:"BypassWidth,omitempty" name:"BypassWidth"`
		// 赠送的旁路带宽数据

		SendBypassWidth *int64 `json:"SendBypassWidth,omitempty" name:"SendBypassWidth"`
		// 串行地域带宽分配

		SerialRegionLst []*SerialRegionInfo `json:"SerialRegionLst,omitempty" name:"SerialRegionLst"`
		// 可配置实例个数

		UnUsedQuota *int64 `json:"UnUsedQuota,omitempty" name:"UnUsedQuota"`
		// 剩余可分配通用带宽&nbsp;单位M

		UnUsedWidth *int64 `json:"UnUsedWidth,omitempty" name:"UnUsedWidth"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSerialRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSerialRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalanceInsInfo struct {

	// 当前负载均衡能绑定的proxy机器信息

	CvmList []*ProxyCvmInfo `json:"CvmList,omitempty" name:"CvmList"`
	// 负载均衡实例&nbsp;的&nbsp;监听器列表

	ListenerList []*ListenerInfo `json:"ListenerList,omitempty" name:"ListenerList"`
	// 负载均衡实例id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡所属vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type NoticeCommon struct {

	// 0&nbsp;关闭
	// 1&nbsp;打开

	AutoRebuild *string `json:"AutoRebuild,omitempty" name:"AutoRebuild"`
	// auto&nbsp;&nbsp;&nbsp;0&nbsp;&nbsp;&nbsp;30&nbsp;&nbsp;&nbsp;60&nbsp;180&nbsp;720&nbsp;1440

	DelayRebuild *string `json:"DelayRebuild,omitempty" name:"DelayRebuild"`
}

type DescribeAutoInternetSwitchRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAutoInternetSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoInternetSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeScoreOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSafeScoreOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeScoreOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlowApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatAclRuleSwitchRequest struct {
	*tchttp.BaseRequest

	// 规则方向：1，入站；0，出站

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 需要操作开关的规则ID列表，注意：如果传入的是[-1]将开启或关闭所有规则开关

	RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
	// on开启规则，off关闭规则

	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyNatAclRuleSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAclRuleSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveNatAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功后返回被删除策略的uuid列表

		RuleUuid []*int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveNatAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIpsRuleSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0-不支持编辑、1-支持编辑

		EnableModify *uint64 `json:"EnableModify,omitempty" name:"EnableModify"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckIpsRuleSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIpsRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClsDeliverTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值&nbsp;0：成功，非&nbsp;0：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClsDeliverTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClsDeliverTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncTask struct {

	// 是否覆盖导入&nbsp;1是

	CoverDuplicate *int64 `json:"CoverDuplicate,omitempty" name:"CoverDuplicate"`
	// 自定义的参数

	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
	// 任务进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务状态
	// 0进行中，1已完成，2异常失败

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type CreateNatFwInstanceWithDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// nat实例信息

		CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatFwInstanceWithDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwInstanceWithDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwSyncStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步状态：1，同步中；0，同步完成

		SyncStatus *int64 `json:"SyncStatus,omitempty" name:"SyncStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwSyncStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchCommonRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;关闭&nbsp;
	// 1&nbsp;打开

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
	// 前后端调用约定&nbsp;TiSwitchIn/TiSwitchOut

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
}

func (r *ModifySwitchCommonRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchCommonRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwNameRequest struct {
	*tchttp.BaseRequest

	// 改名的实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 改名新名称

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
}

func (r *ModifyVpcFwNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LeakTrendsHead struct {

	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 事件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 事件次数

	Times *int64 `json:"Times,omitempty" name:"Times"`
}

type DescribeImportWeChatSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 零信任微信批量导入返回结构体

		Data *ImportWeChatSpec `json:"Data,omitempty" name:"Data"`
		// 状态&nbsp;0成功&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误码

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportWeChatSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportWeChatSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionConnections struct {

	// 连接数

	Connections *int64 `json:"Connections,omitempty" name:"Connections"`
	// 访问目的IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 访问目的端口

	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
	// ip版本,0,IPV4;1,IPV6

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 访问源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 限速状态

	TcRuleStatus *int64 `json:"TcRuleStatus,omitempty" name:"TcRuleStatus"`
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除结果,0成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatFwInstanceRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DeleteNatFwInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatFwInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebCosUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cos域名

		Data *string `json:"Data,omitempty" name:"Data"`
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebCosUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebCosUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwEipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetNatFwEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetNatFwEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreListRequest struct {
	*tchttp.BaseRequest

	// 排序类型：desc降序，asc正序

	By *string `json:"By,omitempty" name:"By"`
	// 方向：1互联网入站，0互联网出站，3内网，空&nbsp;全部方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序列：EndTime结束时间，StartTime开始时间，MatchTimes命中次数

	Order *string `json:"Order,omitempty" name:"Order"`
	// 规则类型：1封禁，2放通

	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`
	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// blocklist&nbsp;封禁列表
	// whitelist&nbsp;白名单列表

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeBlockIgnoreListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatConnectionsTrendsRequest struct {
	*tchttp.BaseRequest

	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// nat实例Id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatConnectionsTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatConnectionsTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOAFilterGroup struct {

	// Filters&nbsp;条件过滤

	Filters []*IOAFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeSerialRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSerialRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSerialRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogStatus struct {

	// 可用网卡数

	AvailableEniNum *int64 `json:"AvailableEniNum,omitempty" name:"AvailableEniNum"`
	// 1表示开启，0表示关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// vpcid对应的vpc的日志开关状态，vpcid为空则代表所有vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type FwCidrInfo struct {

	// 其他防火墙占用网段，一般是防火墙需要独占vpc时指定的网段

	ComFwCidr *string `json:"ComFwCidr,omitempty" name:"ComFwCidr"`
	// 为每个vpc指定防火墙的网段

	FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitempty" name:"FwCidrLst"`
	// 防火墙使用的网段类型，值VpcSelf/Assis/Custom分别代表自有网段优先/扩展网段优先/自定义

	FwCidrType *string `json:"FwCidrType,omitempty" name:"FwCidrType"`
}

type ResourceInfo struct {

	// 资源ID，全局唯一

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源ID名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type BakRuleSearchName struct {

	// 实例信息

	InsMsg *string `json:"InsMsg,omitempty" name:"InsMsg"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type CancelUpdateEngineReserveTaskRequest struct {
	*tchttp.BaseRequest

	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
}

func (r *CancelUpdateEngineReserveTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelUpdateEngineReserveTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncTaskRequest struct {
	*tchttp.BaseRequest

	// 任务类型&nbsp;sg_rule_compare&nbsp;安全组预览变动任务，其他后续补充

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *CreateAsyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDynamicListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDynamicListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDynamicListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnHandleEvent struct {

	// 1&nbsp;打开&nbsp;0&nbsp;关闭

	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitempty" name:"BaseLineInSwitch"`
	// 1&nbsp;打开&nbsp;0&nbsp;关闭

	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitempty" name:"BaseLineOutSwitch"`
	// 1&nbsp;是&nbsp;&nbsp;0否

	BaseLineUser *uint64 `json:"BaseLineUser,omitempty" name:"BaseLineUser"`
	// 伪攻击链类型

	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitempty" name:"EventTableListStruct"`
	// vpc间防火墙实例数量

	VpcFwCount *uint64 `json:"VpcFwCount,omitempty" name:"VpcFwCount"`
}

type ModifyClsDeliverTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值&nbsp;0：成功，非&nbsp;0：失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClsDeliverTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClsDeliverTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEngineRequest struct {
	*tchttp.BaseRequest

	// nat防火墙实例列表

	NatFwInsLst []*string `json:"NatFwInsLst,omitempty" name:"NatFwInsLst"`
	// 预约引擎升级时间，为空时表示直接升级

	ReserveTime *string `json:"ReserveTime,omitempty" name:"ReserveTime"`
	// sdwan防火墙实例列表

	SdWanFwInsLst []*string `json:"SdWanFwInsLst,omitempty" name:"SdWanFwInsLst"`
	// 是否升级NAT防火墙，可不需要该字段

	UpdateNatFw *bool `json:"UpdateNatFw,omitempty" name:"UpdateNatFw"`
	// 升级版本号

	UpdateVersion *string `json:"UpdateVersion,omitempty" name:"UpdateVersion"`
	// 是否升级vpc间防火墙，可不需要该字段

	UpdateVpcFw *bool `json:"UpdateVpcFw,omitempty" name:"UpdateVpcFw"`
	// vpc防火墙实例列表

	VpcFwInsLst []*string `json:"VpcFwInsLst,omitempty" name:"VpcFwInsLst"`
}

func (r *UpdateEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertCenterOmitRequest struct {
	*tchttp.BaseRequest

	// 处置对象,ID列表，&nbsp;&nbsp;IdLists和IpList二选一

	HandleIdList []*string `json:"HandleIdList,omitempty" name:"HandleIdList"`
	// 忽略数据来源：
	// AlertTable&nbsp;告警中心&nbsp;&nbsp;InterceptionTable拦截列表

	TableType *string `json:"TableType,omitempty" name:"TableType"`
}

func (r *CreateAlertCenterOmitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterOmitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcFwGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙组ID

		FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcFwGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbOverviewRequest struct {
	*tchttp.BaseRequest

	// 时间类型：1:&nbsp;24小时数据，2:&nbsp;7天数据

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeDbOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperateLogsRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 防火墙类型

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 危险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 日志类型，ids:入侵防御操作，acl:访问控制，login:登录控制,&nbsp;switch:&nbsp;防火墙开关

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 规则名称

	ReportRule *string `json:"ReportRule,omitempty" name:"ReportRule"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 不再使用

	SwitchType *uint64 `json:"SwitchType,omitempty" name:"SwitchType"`
}

func (r *DescribeOperateLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperateLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewModeItems struct {

	// 新增模式下新增绑定的出口弹性公网ip个数，其中Eips和AddCount至少传递一个。

	AddCount *int64 `json:"AddCount,omitempty" name:"AddCount"`
	// 新增模式下绑定的出口弹性公网ip列表，其中Eips和AddCount至少传递一个。

	Eips []*string `json:"Eips,omitempty" name:"Eips"`
	// 新增模式下接入的vpc列表

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

type DescribeGuideScanInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuideScanInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideScanInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCFWFileRequest struct {
	*tchttp.BaseRequest

	// 是否异步导入&nbsp;1异步&nbsp;2同步

	AsyncType *int64 `json:"AsyncType,omitempty" name:"AsyncType"`
	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据

	CoverDuplicate *int64 `json:"CoverDuplicate,omitempty" name:"CoverDuplicate"`
	// 自定义参数&nbsp;扩展预留

	CustomParameter *string `json:"CustomParameter,omitempty" name:"CustomParameter"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 文件cos地址

	ImportFileUrl *string `json:"ImportFileUrl,omitempty" name:"ImportFileUrl"`
	// 导入类型acl_secgroupnew_import&nbsp;
	// acl_vpc_import&nbsp;acl_dns_import&nbsp;
	// acl_border_import_in
	// acl_border_import_out
	// acl_nat_import_in
	// acl_nat_import_out
	// 其他后续补充&nbsp;

	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`
	// Ip版本，0：IPv4，1：IPv6，默认为IPv4

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *ImportCFWFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCFWFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSyncAssetRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;互联网防火墙开关，1：vpc&nbsp;防火墙开关

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *RunSyncAssetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunSyncAssetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleRegionListRequest struct {
	*tchttp.BaseRequest

	// 页高

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNatRuleRegionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRuleRegionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartUpdateResourceTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *StartUpdateResourceTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartUpdateResourceTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatFwInstance struct {

	// 0:新增模式，1:接入模式

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
	// ip

	NatIp *string `json:"NatIp,omitempty" name:"NatIp"`
	// nat实例id

	NatinsId *string `json:"NatinsId,omitempty" name:"NatinsId"`
	// nat实例名称

	NatinsName *string `json:"NatinsName,omitempty" name:"NatinsName"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 0:正常状态，&nbsp;1:&nbsp;正在创建

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeIpStatLstRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 防火墙的实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeIpStatLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpStatLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewAuthInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权页复选框&nbsp;返回信息

		Data *AuthInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewAuthInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncTaskErrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常信息

		Data *AsyncErrInfo `json:"Data,omitempty" name:"Data"`
		// 0&nbsp;正常&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncTaskErrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncTaskErrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertConfig struct {

	// 高危告警

	HighWarn *int64 `json:"HighWarn,omitempty" name:"HighWarn"`
	// 低危告警

	LowWarn *int64 `json:"LowWarn,omitempty" name:"LowWarn"`
	// 中危告警

	MiddleWarn *int64 `json:"MiddleWarn,omitempty" name:"MiddleWarn"`
	// 严重告警

	SeriousWarn *int64 `json:"SeriousWarn,omitempty" name:"SeriousWarn"`
	// 提示告警

	SuggestWarn *int64 `json:"SuggestWarn,omitempty" name:"SuggestWarn"`
}

type BackupNextHop struct {

	// vpc间防火墙的云联网模式的下一跳为云联网路由表id

	NextHopId *string `json:"NextHopId,omitempty" name:"NextHopId"`
	// 下一跳的名称

	NextHopName *string `json:"NextHopName,omitempty" name:"NextHopName"`
}

type DescribeOfflineExportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 离线导出任务列表

		Data []*OfflineExportTask `json:"Data,omitempty" name:"Data"`
		// 导出数据限制

		ExportLimit *int64 `json:"ExportLimit,omitempty" name:"ExportLimit"`
		// 导出文件配额，单位B

		ExportQuota *int64 `json:"ExportQuota,omitempty" name:"ExportQuota"`
		// 剩余导出文件配额，单位B

		ExportRemainQuota *int64 `json:"ExportRemainQuota,omitempty" name:"ExportRemainQuota"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 任务数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOfflineExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOfflineExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true代表开启，false代表关闭

		Status *bool `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwCrossStatusRequest struct {
	*tchttp.BaseRequest

	// 边id&nbsp;或者VPC&nbsp;id

	CrossInsId *string `json:"CrossInsId,omitempty" name:"CrossInsId"`
}

func (r *DescribeVpcFwCrossStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwCrossStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwGroupIdNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙(组列表)

		Data []*FwGroupIdName `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwGroupIdNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwGroupIdNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSubnetStatLstRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 防火墙的实例id

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNatSubnetStatLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSubnetStatLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePcapTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 抓包任务列表

		Data []*DescPcapTask `json:"Data,omitempty" name:"Data"`
		// 当天剩余可创建任务数

		RemainCnt *uint64 `json:"RemainCnt,omitempty" name:"RemainCnt"`
		// 列表总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePcapTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePcapTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeScoreOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全评分维度数据展示

		Data []*SafeScoreOverviewData `json:"Data,omitempty" name:"Data"`
		// 评分

		Score *uint64 `json:"Score,omitempty" name:"Score"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeScoreOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSafeScoreOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclRulePathStatus struct {

	// 文件Url路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 状态&nbsp;0成功&nbsp;1失败

	Status *string `json:"Status,omitempty" name:"Status"`
}

type FwVpcCidr struct {

	// 防火墙网段，最少/24的网段

	FwCidr *string `json:"FwCidr,omitempty" name:"FwCidr"`
	// vpc的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeIpStatLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 单台cvm峰值带宽统计列表

		Data []*IpStatStruct `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总行数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpStatLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpStatLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIsolateListRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 升序降序&nbsp;&nbsp;asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 查询条件
	// "SearchValue":{"InstanceID":"1",&nbsp;&nbsp;"InstanceName":&nbsp;"v_tpq"},
	// &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;查询排序的字段
	// 	"SrcIp":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//源&nbsp;&nbsp;多条件&nbsp;1,2,3&nbsp;格式
	// 	"Uuid":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//规则id
	// 	"DstContent"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//目的
	// 	"CloudCode":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//云厂商
	// 	"UpdateTime":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//时间
	// 	"Disable":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//状态
	// 	"Resource":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//资产组
	// 	"InsId":&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//资产
	// 	"DetectedTimes"&nbsp;//次数&nbsp;
	//
	// common&nbsp;不指定字段&nbsp;和下面FuzzySearch&nbsp;&nbsp;功能&nbsp;一直&nbsp;&nbsp;前端那个好处理使用那个

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeIsolateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIsolateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcFwRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwGroupSwitch struct {

	// 防火墙开关ID,支持三种类型,1.&nbsp;边开关(单点互通),2.&nbsp;点开关(多点互通),3.&nbsp;全开关(全互通)

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 防火墙实例的开关模式&nbsp;1:&nbsp;单点互通&nbsp;2:&nbsp;多点互通&nbsp;3:&nbsp;全互通&nbsp;4:&nbsp;自定义路由

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

type DeleteNatFwInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatFwInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*FlowLogData `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetsOverviewData struct {

	// 资产名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 未打开防火墙个数

	Off *uint64 `json:"Off,omitempty" name:"Off"`
	// 已打开防火墙个数

	On *uint64 `json:"On,omitempty" name:"On"`
}

type ProbeInfo struct {

	// 容器内端口

	DefaultPort *int64 `json:"DefaultPort,omitempty" name:"DefaultPort"`
	// 部署模式：1&nbsp;公网IP&nbsp;2&nbsp;负载均衡

	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`
	// 探针为eip时的外网ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 探针id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 探针名称

	ProbeName *string `json:"ProbeName,omitempty" name:"ProbeName"`
	// 探针负载均衡模式路径

	ProbePath *string `json:"ProbePath,omitempty" name:"ProbePath"`
	// 探针要使用的端口

	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
	// 探针要使用的端口范围

	ProbePortRange *string `json:"ProbePortRange,omitempty" name:"ProbePortRange"`
}

type ZeroTrustError struct {

	// 并发账号account，带宽bandwidth

	AlertType *string `json:"AlertType,omitempty" name:"AlertType"`
	// 告警值

	AlertValue *float64 `json:"AlertValue,omitempty" name:"AlertValue"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 告警当时规格

	Spec *int64 `json:"Spec,omitempty" name:"Spec"`
	// 告警时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeCheckCLSStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCheckCLSStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckCLSStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IOAAccessDirectory struct {

	// 接入目录Id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 已接入IOA子用户组列表

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 接入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 目录名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// IOA用户组信息

	RootAccountGroup *DescribeAccountGroupsData `json:"RootAccountGroup,omitempty" name:"RootAccountGroup"`
	// 0:完成,&nbsp;1:接入中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 同步数据源,0表示本地账号,LDAP&nbsp;10001,&nbsp;企业微信&nbsp;10002,&nbsp;政务微信/私有化企业微信&nbsp;10003,&nbsp;飞书&nbsp;10004,&nbsp;钉钉&nbsp;10005,&nbsp;WindowsAD&nbsp;10006,&nbsp;SCIM&nbsp;10007

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 接入数量

	UserTotal *string `json:"UserTotal,omitempty" name:"UserTotal"`
	// 7天访问次数

	VisitsTimes *int64 `json:"VisitsTimes,omitempty" name:"VisitsTimes"`
}

type DescribeNatFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrictModeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrictModeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrictModeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertTimeRange struct {

	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
}

type NetInstancesInfo struct {

	// 网络cidr&nbsp;(多段以逗号分隔)

	InstanceCidr *string `json:"InstanceCidr,omitempty" name:"InstanceCidr"`
	// 网络实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 网络实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 网络实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeNatGwListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户vpc信息列表

		Data []*NatGwInfo `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGwListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGwListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNtaFileRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNtaFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNtaFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest

	// NAT地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 规则数组

	Data []*RuleInfoData `json:"Data,omitempty" name:"Data"`
	// EdgeId值

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 访问规则状态

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterDataObject struct {

	// 显示名称

	Text *string `json:"Text,omitempty" name:"Text"`
	// 实际值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeNatFwRouteBackupLstRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeNatFwRouteBackupLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatFwRouteBackupLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwGroupSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFwGroupSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwGroupSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProposeListStruct struct {

	// 影响资产数

	AffectAssetNum *int64 `json:"AffectAssetNum,omitempty" name:"AffectAssetNum"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 标签列表

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 0不限关&nbsp;1与我相关

	Related *int64 `json:"Related,omitempty" name:"Related"`
	// insert_time

	Time *string `json:"Time,omitempty" name:"Time"`
	// 事件名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 图片地址

	Url *string `json:"Url,omitempty" name:"Url"`
}

type RulePair struct {

	// 规则源

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 规则目的

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
}

type ModifyVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 公网IP

	FirewallVpcId *string `json:"FirewallVpcId,omitempty" name:"FirewallVpcId"`
	// 状态值，0:&nbsp;关闭&nbsp;,1:开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVPCSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVPCSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetZone struct {

	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 地域英文

	ZoneEng *string `json:"ZoneEng,omitempty" name:"ZoneEng"`
}

type ModifyNatAcRuleRequest struct {
	*tchttp.BaseRequest

	// 需要编辑的规则数组

	Rules []*CreateNatRuleItem `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyNatAcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowCenterData struct {

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 资产名称

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产地域

	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`
	// 会话数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 方向：1：外部访问&nbsp;0主动外联

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
	// 防火墙开关id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 请求结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是互联网还是nat

	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`
	// 地理位置

	Region *string `json:"Region,omitempty" name:"Region"`
	// 访问流量请求

	RequestCount *int64 `json:"RequestCount,omitempty" name:"RequestCount"`
	// 访问流量响应

	ResponseCount *int64 `json:"ResponseCount,omitempty" name:"ResponseCount"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 请求开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
}

type DescribeZeroTrustBlockPreviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZeroTrustBlockPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZeroTrustBlockPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindEdgeCFWRequest struct {
	*tchttp.BaseRequest

	// 云防火墙（东西向防火墙）实例ID

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// Edge设备ID

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *UnbindEdgeCFWRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindEdgeCFWRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcProtectListRequest struct {
	*tchttp.BaseRequest

	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索参数，json格式字符串，空则传"{}"

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeVpcProtectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcProtectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 修改成功与否的状态值&nbsp;0：修改成功，非&nbsp;0：修改失败

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwRelateInstances struct {

	// 防火墙实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 防火墙实例名称

	CfwInsName *string `json:"CfwInsName,omitempty" name:"CfwInsName"`
	// 网络实例网段列表

	Cidrs *string `json:"Cidrs,omitempty" name:"Cidrs"`
	// 网络实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例类型

	InstanceIdType *string `json:"InstanceIdType,omitempty" name:"InstanceIdType"`
	// 网络实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type VpcDnsInfo struct {

	// 外网弹性ip，防火墙&nbsp;dns解析地址

	DNSEip *string `json:"DNSEip,omitempty" name:"DNSEip"`
	// nat&nbsp;防火墙模式&nbsp;0：新增模式，&nbsp;1:&nbsp;接入模式

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
	// nat网关id

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// nat网关名称

	NatInsName *string `json:"NatInsName,omitempty" name:"NatInsName"`
	// 0：未防护，&nbsp;1:&nbsp;已防护，2：忽略此字段

	ProtectedStatus *uint64 `json:"ProtectedStatus,omitempty" name:"ProtectedStatus"`
	// 是否支持DNS&nbsp;FW，0-不支持、1-支持

	SupportDNSFW *uint64 `json:"SupportDNSFW,omitempty" name:"SupportDNSFW"`
	// 0：开关关闭&nbsp;，&nbsp;1:&nbsp;开关打开

	SwitchStatus *int64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
	// vpc&nbsp;id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc&nbsp;ipv4网段范围&nbsp;CIDR（Classless&nbsp;Inter-Domain&nbsp;Routing，无类域间路由选择）

	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitempty" name:"VpcIpv4Cidr"`
	// vpc&nbsp;名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeUserStorageTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;es&nbsp;1&nbsp;filtration&nbsp;2&nbsp;ck

		Data *string `json:"Data,omitempty" name:"Data"`
		// 值

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 结果信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserStorageTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserStorageTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProbeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProbeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideScanInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 扫描信息

		Data *ScanInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGuideScanInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIgnoreImportCredentialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 凭证过期时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 凭证有效开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 临时SecretId

		TemplateSecretId *string `json:"TemplateSecretId,omitempty" name:"TemplateSecretId"`
		// 临时SecretKey

		TemplateSecretKey *string `json:"TemplateSecretKey,omitempty" name:"TemplateSecretKey"`
		// Token

		Token *string `json:"Token,omitempty" name:"Token"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIgnoreImportCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIgnoreImportCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventNewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Data []*NewsDetail `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventNewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventNewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessServiceInfo struct {

	// 实例数量

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 接入功能

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 开关状态，0关闭，1开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ResetNatRuleHitTimesRequest struct {
	*tchttp.BaseRequest

	// 规则的uuid，可通过查询规则列表获取

	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *ResetNatRuleHitTimesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetNatRuleHitTimesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpStatic struct {

	// 值

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 折线图横坐标时间

	StatTime *string `json:"StatTime,omitempty" name:"StatTime"`
}

type NDRAssetInfo struct {

	// agent状态&nbsp;0&nbsp;离线,1&nbsp;防护中&nbsp;2&nbsp;未防护&nbsp;

	AgentStatus *int64 `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 采集类型&nbsp;&nbsp;&nbsp;0&nbsp;agent&nbsp;&nbsp;1&nbsp;流量镜像

	CollectionType *int64 `json:"CollectionType,omitempty" name:"CollectionType"`
	// NDR开关&nbsp;&nbsp;0&nbsp;关闭&nbsp;1&nbsp;开启

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 开关状态&nbsp;&nbsp;0&nbsp;为操作完成&nbsp;&nbsp;1&nbsp;为开启中&nbsp;2关闭中

	EnabledStatus *int64 `json:"EnabledStatus,omitempty" name:"EnabledStatus"`
	// 私有连接ID

	EndpointId *string `json:"EndpointId,omitempty" name:"EndpointId"`
	// 私有连接使用的CVM个数

	EndpointInsNum *int64 `json:"EndpointInsNum,omitempty" name:"EndpointInsNum"`
	// 私有连接IP

	EndpointIp *string `json:"EndpointIp,omitempty" name:"EndpointIp"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资产类型&nbsp;CVM/LH

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 内网ip列表

	PrivateIpList []*string `json:"PrivateIpList,omitempty" name:"PrivateIpList"`
	// 公网信息(将废弃)

	PublicIp []*NDRAssetPublicIpInfo `json:"PublicIp,omitempty" name:"PublicIp"`
	// 公网ip列表

	PublicIpList []*string `json:"PublicIpList,omitempty" name:"PublicIpList"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 开关权重&nbsp;0&nbsp;-&nbsp;100

	SwitchWeight *int64 `json:"SwitchWeight,omitempty" name:"SwitchWeight"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type ModifyNatSequenceRulesRequest struct {
	*tchttp.BaseRequest

	// 规则方向：1，入站；0，出站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号

	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitempty" name:"RuleChangeItems"`
}

func (r *ModifyNatSequenceRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatSequenceRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowStatisticsOverviewRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFlowStatisticsOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowStatisticsOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest

	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 排序字段&nbsp;&nbsp;LastTime最近扫描时间，PortTimes风险端口数

	By *string `json:"By,omitempty" name:"By"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，desc：降序，asc：升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 防火墙状态&nbsp;0正常其他异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 资产类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSwitchListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTLogInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "NetworkNum":网络扫描探测
		// &nbsp;"HandleNum":&nbsp;待处理事件
		// "BanNum":&nbsp;
		// &nbsp;&nbsp;"VulNum":&nbsp;漏洞利用
		// &nbsp;&nbsp;"OutNum":&nbsp;失陷主机
		// "BruteForceNum":&nbsp;0

		Data *TLogInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTLogInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBorderFwRuleHitDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则命中详情

		Data []*BorderRuleDetail `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBorderFwRuleHitDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderFwRuleHitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuidePortProtectRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuidePortProtectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuidePortProtectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 资产组id&nbsp;&nbsp;全部传0

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 查询类型&nbsp;网络结构&nbsp;vpc，业务识别-&nbsp;resource&nbsp;，资源标签-tag

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// all&nbsp;&nbsp;包含子组&nbsp;own自己

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
}

func (r *DescribeResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescAcItem struct {

	// 关联任务详情

	BetaList []*BetaInfoByACL `json:"BetaList,omitempty" name:"BetaList"`
	// 城市id

	CityCode *uint64 `json:"CityCode,omitempty" name:"CityCode"`
	// 省名称

	CityName *string `json:"CityName,omitempty" name:"CityName"`
	// 云厂商code

	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`
	// 命中次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 国家id

	CountryCode *uint64 `json:"CountryCode,omitempty" name:"CountryCode"`
	// 国家名称

	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 规则方向：1，入向；0，出向

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 若访问目的类型是dnsparse，即协议解析类型时，该值有效

	DnsParseCnt *int64 `json:"DnsParseCnt,omitempty" name:"DnsParseCnt"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 内部使用的uuid，一般情况下不会使用到该字段

	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`
	// 互联网边界防火墙使用的内部规则id

	InternetBorderUuid *string `json:"InternetBorderUuid,omitempty" name:"InternetBorderUuid"`
	// 规则有效性

	Invalid *uint64 `json:"Invalid,omitempty" name:"Invalid"`
	// 0为正常规则,1为云厂商规则

	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`
	// 0为正常规则,1为地域规则

	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`
	// 规则最近命中时间

	LastHitTime *string `json:"LastHitTime,omitempty" name:"LastHitTime"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口协议组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 端口协议组名称

	ParamTemplateName *string `json:"ParamTemplateName,omitempty" name:"ParamTemplateName"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：&nbsp;accept：放行&nbsp;drop：拒绝&nbsp;log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 生效范围：serial，串行；side，旁路；all，全局

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 生效范围描述

	ScopeDesc *string `json:"ScopeDesc,omitempty" name:"ScopeDesc"`
	// 访问源

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 访问源类型：入向规则时类型可以为&nbsp;ip,net,template,location；出向规则时可以为&nbsp;ip,net,template,instance,group,tag

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 规则状态，查询规则命中详情时该字段有效，0：新增，1:&nbsp;已删除,&nbsp;2:&nbsp;编辑删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 访问目的

	TargetContent *string `json:"TargetContent,omitempty" name:"TargetContent"`
	// 访问目的名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为&nbsp;ip,net,domain,template,location

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 规则对应的唯一id

	Uuid *uint64 `json:"Uuid,omitempty" name:"Uuid"`
}

type NDRConfig struct {

	// NDR带宽

	BuyNdrBandwidth *int64 `json:"BuyNdrBandwidth,omitempty" name:"BuyNdrBandwidth"`
	// 是否试用用户

	IsWhiteUser *int64 `json:"IsWhiteUser,omitempty" name:"IsWhiteUser"`
	// 全部新增资产自动开启；0关闭，1开启

	NewAssetsAutoStart *int64 `json:"NewAssetsAutoStart,omitempty" name:"NewAssetsAutoStart"`
	// 仅公网新增资产自动开启；0关闭，1开启

	NewPublicAssetsAutoStart *int64 `json:"NewPublicAssetsAutoStart,omitempty" name:"NewPublicAssetsAutoStart"`
	// NDR功能状态，

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 试用到期时间

	WhiteExpireTime *string `json:"WhiteExpireTime,omitempty" name:"WhiteExpireTime"`
}

type DescribeAssetsOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetsOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetsOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEngineUpdateDetailRequest struct {
	*tchttp.BaseRequest

	// 选定的版本，基于选定的版本判断实例是否可升级和回退到该版本

	SelectedVersion *string `json:"SelectedVersion,omitempty" name:"SelectedVersion"`
}

func (r *DescribeEngineUpdateDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEngineUpdateDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowDistributeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*MapVisual `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowDistributeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowDistributeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserObject struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 1新用户&nbsp;0付费

	Isnew *uint64 `json:"Isnew,omitempty" name:"Isnew"`
	// 通知等级

	NotifyLevel []*string `json:"NotifyLevel,omitempty" name:"NotifyLevel"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateAsyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口状态0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态码&nbsp;success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态&nbsp;1失败&nbsp;0成功

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAsyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAclRuleHitTimesRequest struct {
	*tchttp.BaseRequest

	// 规则的uuid，可通过查询规则列表获取

	RuleUuid *int64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
}

func (r *ResetAclRuleHitTimesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAclRuleHitTimesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpsRuleDetailNew struct {

	// 0观察,&nbsp;1阻断

	Action *int64 `json:"Action,omitempty" name:"Action"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 置信度

	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`
	// cve

	Cve *string `json:"Cve,omitempty" name:"Cve"`
	// 默认动作

	DefaultAction *int64 `json:"DefaultAction,omitempty" name:"DefaultAction"`
	// 规则名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 描述

	EventNameDesc *string `json:"EventNameDesc,omitempty" name:"EventNameDesc"`
	// FwType字段&nbsp;1&nbsp;border&nbsp;2&nbsp;nat&nbsp;4&nbsp;vpc

	FwType *int64 `json:"FwType,omitempty" name:"FwType"`
	// iD

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否重保规则
	// 1是
	// 0不是

	ImportantRule *int64 `json:"ImportantRule,omitempty" name:"ImportantRule"`
	// 等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 1111

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 基础防御/虚拟补丁

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 状态&nbsp;0&nbsp;关闭&nbsp;1打开

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 	漏洞对象

	VulTarget *string `json:"VulTarget,omitempty" name:"VulTarget"`
}

type BaradStatus struct {

	// 云监控默认字段默认值

	Status *string `json:"Status,omitempty" name:"Status"`
}

type CreateIdsWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码：
		// 0&nbsp;成功
		// 非0&nbsp;失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息：
		// success&nbsp;成功
		// 其他

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 返回状态码：
		// 0&nbsp;&nbsp;处置成功
		// -1&nbsp;通用错误，不用处理

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIdsWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIdsWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBetaTaskAclRuleListRequest struct {
	*tchttp.BaseRequest

	// 规则集合&nbsp;

	AclList []*string `json:"AclList,omitempty" name:"AclList"`
	// &nbsp;1border&nbsp;2nat&nbsp;3vpc&nbsp;4sg&nbsp;规则类型

	AclType *int64 `json:"AclType,omitempty" name:"AclType"`
	// limit&nbsp;默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// offset&nbsp;默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBetaTaskAclRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskAclRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest

	// 组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 上级组id

	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
}

func (r *ModifyResourceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 阻断记录集合

		Data []*BlockInfo `json:"Data,omitempty" name:"Data"`
		// 置顶数量

		TopCount *int64 `json:"TopCount,omitempty" name:"TopCount"`
		// 置顶阻断记录集合

		TopData []*BlockInfo `json:"TopData,omitempty" name:"TopData"`
		// 数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEWRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：修改成功，非0：修改失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEWRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEWRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatFwSwitchAllRequest struct {
	*tchttp.BaseRequest

	// 0&nbsp;关闭开关；&nbsp;1&nbsp;打开开关

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ModifyNatFwSwitchAllRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatFwSwitchAllRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwFirstConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;第一次进入该页面&nbsp;1&nbsp;非第一次进入该页面

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFwFirstConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwFirstConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoAssistantUserConf struct {

	// 规则ID列表-新

	AclList *string `json:"AclList,omitempty" name:"AclList"`
	// 规则分类

	Action []*string `json:"Action,omitempty" name:"Action"`
	// 告警分类

	AlertKillChain []*string `json:"AlertKillChain,omitempty" name:"AlertKillChain"`
	// 告警来源

	AlertSource []*string `json:"AlertSource,omitempty" name:"AlertSource"`
	// 关联资产

	AssetsList *string `json:"AssetsList,omitempty" name:"AssetsList"`
	// 分类方式&nbsp;1预设分类&nbsp;2详细分类

	Classification *int64 `json:"Classification,omitempty" name:"Classification"`
	// 事件类型

	EventNameList *string `json:"EventNameList,omitempty" name:"EventNameList"`
	// 告警等级high,low,medium

	Level []*string `json:"Level,omitempty" name:"Level"`
	// 选择范围&nbsp;1全部对象&nbsp;2分类选择&nbsp;3自定义选择&nbsp;0未选择

	SelectLimits *int64 `json:"SelectLimits,omitempty" name:"SelectLimits"`
	// 对象类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
}

type GetOverviewTopicStruct struct {

	// 影响资产数量

	AffectAssetNum *int64 `json:"AffectAssetNum,omitempty" name:"AffectAssetNum"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 专题标题

	EventTitle *string `json:"EventTitle,omitempty" name:"EventTitle"`
	// 标签列表

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 安全等级，safe-安全，low-低危，mid-中危，high-高危

	Level *string `json:"Level,omitempty" name:"Level"`
	// 与我相关

	Related *int64 `json:"Related,omitempty" name:"Related"`
	// 发布时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 专题topic

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 专题URL

	Url *string `json:"Url,omitempty" name:"Url"`
}

type VpcFwBarStatus struct {

	// 单位为M

	BandWidthUsed *int64 `json:"BandWidthUsed,omitempty" name:"BandWidthUsed"`
	// 跨地域峰值带宽对应实例ID

	DiffRegionFwInsId *string `json:"DiffRegionFwInsId,omitempty" name:"DiffRegionFwInsId"`
	// 跨地域峰值带宽对应实例名称

	DiffRegionFwInsName *string `json:"DiffRegionFwInsName,omitempty" name:"DiffRegionFwInsName"`
	// 跨地域峰值带宽数值

	DiffRegionMaxFlow *int64 `json:"DiffRegionMaxFlow,omitempty" name:"DiffRegionMaxFlow"`
	// 防火墙(组)个数

	FwGroupCount *int64 `json:"FwGroupCount,omitempty" name:"FwGroupCount"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙(组)最大峰值带宽

	FwGroupMaxFlow *int64 `json:"FwGroupMaxFlow,omitempty" name:"FwGroupMaxFlow"`
	// 单个组最多建多少防火墙实例

	FwGroupMaxInsCount *int64 `json:"FwGroupMaxInsCount,omitempty" name:"FwGroupMaxInsCount"`
	// 防火墙(组)名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// VPC防火墙实例数

	FwInsCount *int64 `json:"FwInsCount,omitempty" name:"FwInsCount"`
	// 防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 防火墙实例最大峰值带宽

	FwInsMaxFlow *int64 `json:"FwInsMaxFlow,omitempty" name:"FwInsMaxFlow"`
	// 防火墙实例名称

	FwInsName *string `json:"FwInsName,omitempty" name:"FwInsName"`
	// VPC防火墙总开关数量

	FwSwitchCount *int64 `json:"FwSwitchCount,omitempty" name:"FwSwitchCount"`
	// 剩余通用实例个数

	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
	// VPC防火墙总接入网络实例数量

	JoinInstanceCount *int64 `json:"JoinInstanceCount,omitempty" name:"JoinInstanceCount"`
	// VPC防火墙总接入网络实例数量

	JoinPeerInstanceCount *int64 `json:"JoinPeerInstanceCount,omitempty" name:"JoinPeerInstanceCount"`
	// 私有模式下，最大支持接入的Vpc实例个数

	JoinVpcCntLimit *int64 `json:"JoinVpcCntLimit,omitempty" name:"JoinVpcCntLimit"`
	// 同地域峰值带宽对应实例ID

	SameRegionFwInsId *string `json:"SameRegionFwInsId,omitempty" name:"SameRegionFwInsId"`
	// 同地域峰值带宽对应实例名称

	SameRegionFwInsName *string `json:"SameRegionFwInsName,omitempty" name:"SameRegionFwInsName"`
	// 同地域峰值带宽数值

	SameRegionMaxFlow *int64 `json:"SameRegionMaxFlow,omitempty" name:"SameRegionMaxFlow"`
	// 租户最大支持实例个数

	VpcFwInsCntLimit *int64 `json:"VpcFwInsCntLimit,omitempty" name:"VpcFwInsCntLimit"`
	// 单实例支持最大地域个数

	VpcRegionCntLimit *int64 `json:"VpcRegionCntLimit,omitempty" name:"VpcRegionCntLimit"`
	// VPC防火墙的购买带宽&nbsp;2GB&nbsp;单位为G

	VpcThroughPut *int64 `json:"VpcThroughPut,omitempty" name:"VpcThroughPut"`
}

type DescribeBetaTaskRequest struct {
	*tchttp.BaseRequest

	// 查询条件&nbsp;

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 默认不填为20&nbsp;Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 默认不填为0&nbsp;Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBetaTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBetaTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HoneyPotProbeSelector struct {

	// 部署模式：1&nbsp;公网IP&nbsp;2&nbsp;负载均衡

	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`
	// 探针id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 探针名称

	ProbeName *string `json:"ProbeName,omitempty" name:"ProbeName"`
	// 探针负载均衡模式可使用的路径

	ProxyPath []*string `json:"ProxyPath,omitempty" name:"ProxyPath"`
}

type ZeroTrustUserInfo struct {

	// 头像链接

	AvatarUrl *string `json:"AvatarUrl,omitempty" name:"AvatarUrl"`
	// 组织Id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 组织路径

	Depts []*string `json:"Depts,omitempty" name:"Depts"`
	// 过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 是否有子节点

	HasChildren *bool `json:"HasChildren,omitempty" name:"HasChildren"`
	// 用户Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 接入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 最后一次登录资产

	LastLoginCvm *string `json:"LastLoginCvm,omitempty" name:"LastLoginCvm"`
	// 最后一次登录Ip

	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`
	// 最后一次登录地址

	LastLoginPlace *string `json:"LastLoginPlace,omitempty" name:"LastLoginPlace"`
	// 最后一次登录时间

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 最后一次登录的用户Id

	LastLoginUser *string `json:"LastLoginUser,omitempty" name:"LastLoginUser"`
	// 最后一次登录WEB服务

	LastVisitWebDomain *string `json:"LastVisitWebDomain,omitempty" name:"LastVisitWebDomain"`
	// 人员数量

	MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
	// 名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 服务号Id

	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
	// 顺序

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 权限细则是否支持删除

	PermissionCanDelete *int64 `json:"PermissionCanDelete,omitempty" name:"PermissionCanDelete"`
	// 手机号

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 描述

	Remarks []*string `json:"Remarks,omitempty" name:"Remarks"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用户Id

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 访问次数

	VisitsTimes *int64 `json:"VisitsTimes,omitempty" name:"VisitsTimes"`
}

type CreateAlertCenterRuleRequest struct {
	*tchttp.BaseRequest

	// 当前日志方向：&nbsp;0&nbsp;出向&nbsp;1&nbsp;入向

	AlertDirection *int64 `json:"AlertDirection,omitempty" name:"AlertDirection"`
	// 封禁域名-保留字段

	BlockDomain *string `json:"BlockDomain,omitempty" name:"BlockDomain"`
	// 处置描述

	HandleComment *string `json:"HandleComment,omitempty" name:"HandleComment"`
	// 处置方向：&nbsp;0出向&nbsp;1入向&nbsp;0,1出入向&nbsp;3内网

	HandleDirection *string `json:"HandleDirection,omitempty" name:"HandleDirection"`
	// 处置对象,ID列表，&nbsp;&nbsp;IdLists和IpList二选一

	HandleIdList []*string `json:"HandleIdList,omitempty" name:"HandleIdList"`
	// 处置对象,IP列表，&nbsp;&nbsp;IdLists和IpList二选一

	HandleIpList []*string `json:"HandleIpList,omitempty" name:"HandleIpList"`
	// 处置时间
	// 1&nbsp;&nbsp;1天
	// 7&nbsp;&nbsp;&nbsp;7天
	// -2&nbsp;永久

	HandleTime *int64 `json:"HandleTime,omitempty" name:"HandleTime"`
	// 处置类型
	// 当HandleIdList&nbsp;不为空时：1封禁&nbsp;2放通&nbsp;&nbsp;
	// 当HandleIpList&nbsp;不为空时：3放通&nbsp;4封禁

	HandleType *int64 `json:"HandleType,omitempty" name:"HandleType"`
	// 放通原因:
	// 0默认&nbsp;1重复&nbsp;2误报&nbsp;3紧急放通

	IgnoreReason *int64 `json:"IgnoreReason,omitempty" name:"IgnoreReason"`
}

func (r *CreateAlertCenterRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertCenterRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortRiskListRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 检索条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

func (r *DescribeAssetPortRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCFWInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙用户信息

		Data *CfwUserInfo `json:"Data,omitempty" name:"Data"`
		// 返回状态码

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCFWInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCFWInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupBothWayInfo struct {

	// 单/双向下发，0:单向下发，1：双向下发

	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`
	// 掩码地址，多个以英文逗号分隔

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 方向，0：出站，1：入站，默认1

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否是正常规则，0：正常，1：异常

	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`
	// 执行顺序

	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 目的端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 内网IP，多个以英文逗号分隔

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否使用端口协议模板，0：否，1：是

	ProtocolPortType *uint64 `json:"ProtocolPortType,omitempty" name:"ProtocolPortType"`
	// 公网IP，多个以英文逗号分隔

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 端口协议类型参数模板id

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 访问源

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 访问源类型，默认为0，0:&nbsp;IP,&nbsp;1:&nbsp;VPC,&nbsp;2:&nbsp;SUBNET,&nbsp;3:&nbsp;CVM,&nbsp;4:&nbsp;CLB,&nbsp;5:&nbsp;ENI,&nbsp;6:&nbsp;CDB,&nbsp;7:&nbsp;参数模板,&nbsp;100:&nbsp;资产分组

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 是否开关开启，0：未开启，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 策略,&nbsp;1：阻断，2：放行

	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 访问目的

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 访问目的类型，默认为0，0:&nbsp;IP,&nbsp;1:&nbsp;VPC,&nbsp;2:&nbsp;SUBNET,&nbsp;3:&nbsp;CVM,&nbsp;4:&nbsp;CLB,&nbsp;5:&nbsp;ENI,&nbsp;6:&nbsp;CDB,&nbsp;7:&nbsp;参数模板,&nbsp;100:&nbsp;资产分组

	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DeleteBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest

	// 规则类型，1封禁，2放通，不支持域名封禁

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 规则列表

	Rules []*IocListData `json:"Rules,omitempty" name:"Rules"`
}

func (r *DeleteBlockIgnoreRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBlockIgnoreRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 列表类型，只能是下面三种之一：port、address、ip

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
	// 查询条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// top数

	Top *int64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeBlockStaticListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockStaticListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteNextHopInfo struct {

	// 目的cidr，nat防火墙下一跳默认值为0.0.0.0/0

	DestCidr *string `json:"DestCidr,omitempty" name:"DestCidr"`
	// 下一跳id

	NextHopId *string `json:"NextHopId,omitempty" name:"NextHopId"`
	// 下一跳的名称

	NextHopName *string `json:"NextHopName,omitempty" name:"NextHopName"`
	// 下一跳类型，NAT表示NAT网关；EIP：云服务器公网IP

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// NAT网关

	NextHopTypeDesc *string `json:"NextHopTypeDesc,omitempty" name:"NextHopTypeDesc"`
}

type QueryVpcFwSupportSwitchModeRequest struct {
	*tchttp.BaseRequest

	// 云联网模式下该云联网可支持的开关模式

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 部署地域列表

	DeployRegions []*string `json:"DeployRegions,omitempty" name:"DeployRegions"`
	// 防火墙模式

	FwMode *int64 `json:"FwMode,omitempty" name:"FwMode"`
}

func (r *QueryVpcFwSupportSwitchModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVpcFwSupportSwitchModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatFwInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙实例id

		CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatFwInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTreatInfoStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 威胁情报开关

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTreatInfoStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTreatInfoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 选中的防火墙开关Id

	FireWallVpcIds []*string `json:"FireWallVpcIds,omitempty" name:"FireWallVpcIds"`
	// 状态，0：关闭，1：开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAllVPCSwitchStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAllVPCSwitchStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNewBindRoleUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0成功通知，其他失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNewBindRoleUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNewBindRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupAssociationRulesNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// code

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// msg

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数据

		RuleImpactList *RuleImpactNew `json:"RuleImpactList,omitempty" name:"RuleImpactList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceGroupAssociationRulesNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupAssociationRulesNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTiCenterListRequest struct {
	*tchttp.BaseRequest

	// 查询条件&nbsp;"Filter":&nbsp;"{}"

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 页大小

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 查询条件&nbsp;"SearchValue":&nbsp;"{}"

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 页数

	Start *int64 `json:"Start,omitempty" name:"Start"`
}

func (r *DescribeTiCenterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTiCenterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDynamicCheckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDynamicCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDynamicCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStorageSettingRequest struct {
	*tchttp.BaseRequest

	// 执行动作：storagetime&nbsp;修改时间&nbsp;storagecount&nbsp;清空日志

	StorageAction *string `json:"StorageAction,omitempty" name:"StorageAction"`
	// 存储天数

	StorageDay *int64 `json:"StorageDay,omitempty" name:"StorageDay"`
}

func (r *ModifyStorageSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStorageSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclRuleExportStatusRequest struct {
	*tchttp.BaseRequest

	// Md5

	RequestMd5 *string `json:"RequestMd5,omitempty" name:"RequestMd5"`
}

func (r *DescribeAclRuleExportStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclRuleExportStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 0:&nbsp;互联网防火墙开关，1：vpc&nbsp;防火墙开关

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSyncAssetStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSyncAssetStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceSavedKeyInfo struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 已保存密码用户

	PassWordNames []*string `json:"PassWordNames,omitempty" name:"PassWordNames"`
	// 已保存密钥用户

	SecretKeyNames []*SecretKeyNames `json:"SecretKeyNames,omitempty" name:"SecretKeyNames"`
}

type DescribeImportWeChatSpecRequest struct {
	*tchttp.BaseRequest

	// 手机号集合

	PhoneLst []*string `json:"PhoneLst,omitempty" name:"PhoneLst"`
}

func (r *DescribeImportWeChatSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImportWeChatSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Data []*InstanceDomainInfo `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetFlowDomainInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandWidthBanner struct {

	// 服务降级次数

	DowngradeCount *int64 `json:"DowngradeCount,omitempty" name:"DowngradeCount"`
	// &nbsp;1显示&nbsp;0不显示

	IsOverUsed *int64 `json:"IsOverUsed,omitempty" name:"IsOverUsed"`
	// 0无意义&nbsp;1:超额100%&nbsp;2:&nbsp;超额120%，3:超额150%

	OverUsedLevel *int64 `json:"OverUsedLevel,omitempty" name:"OverUsedLevel"`
}

type ReceiveActivityVoucherResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态值，0：领取成功，非0：领取失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReceiveActivityVoucherResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveActivityVoucherResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventNumData struct {

	// 高危事件日/周平均个数

	HighAvgNum *float64 `json:"HighAvgNum,omitempty" name:"HighAvgNum"`
	// 高危事件个数

	HighNum *uint64 `json:"HighNum,omitempty" name:"HighNum"`
	// 待处置入侵事件个数

	InNum *uint64 `json:"InNum,omitempty" name:"InNum"`
	// 待处置外联事件个数

	OutNum *uint64 `json:"OutNum,omitempty" name:"OutNum"`
	// 严重事件日/周平均个数

	SeriousAvgNum *float64 `json:"SeriousAvgNum,omitempty" name:"SeriousAvgNum"`
	// 严重事件个数

	SeriousNum *uint64 `json:"SeriousNum,omitempty" name:"SeriousNum"`
}

type AddNatFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNatFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HoneyPotSwitch struct {

	// 蜜罐服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type DescribeNatProbeEipRequest struct {
	*tchttp.BaseRequest

	// 页数&nbsp;&nbsp;不需要&nbsp;&nbsp;可以不传递

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 资产实例id

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
	// 页位移&nbsp;&nbsp;&nbsp;&nbsp;不需要可以不传递

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNatProbeEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatProbeEipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest

	// 扫描类型：1立即扫描&nbsp;2&nbsp;周期任务

	RangeType *int64 `json:"RangeType,omitempty" name:"RangeType"`
	// 扫描深度：'heavy',&nbsp;'medium',&nbsp;'light'

	ScanDeep *string `json:"ScanDeep,omitempty" name:"ScanDeep"`
	// 立即扫描这个字段传过滤的扫描集合

	ScanFilterIp []*string `json:"ScanFilterIp,omitempty" name:"ScanFilterIp"`
	// RangeType为2&nbsp;是必须添加，定时任务时间

	ScanPeriod *string `json:"ScanPeriod,omitempty" name:"ScanPeriod"`
	// 扫描范围：1端口,&nbsp;2端口+漏扫

	ScanRange *int64 `json:"ScanRange,omitempty" name:"ScanRange"`
	// 1全量2单个

	ScanType *int64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyAssetScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssetScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1-更新中&nbsp;2-更新完成&nbsp;3、4-更新失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncAssetStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSyncAssetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIPBySGSwitchRequest struct {
	*tchttp.BaseRequest

	// 开关，false：关闭，true：开启

	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyBlockIPBySGSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIPBySGSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewEventStruct struct {

	// 资产数

	AffectAssetNum *int64 `json:"AffectAssetNum,omitempty" name:"AffectAssetNum"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 标题

	EventTitle *string `json:"EventTitle,omitempty" name:"EventTitle"`
	// 标题

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// Label

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// Related

	Related *int64 `json:"Related,omitempty" name:"Related"`
	// 发布时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
}

type CheckVpcFwTcRuleRequest struct {
	*tchttp.BaseRequest

	// IP地址或者IP掩码，比如192.168.0.1或者192.168.0.0/24

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// NAT防火墙实例ID

	FwId *string `json:"FwId,omitempty" name:"FwId"`
}

func (r *CheckVpcFwTcRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVpcFwTcRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeQuicklyRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ProbeQuicklyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProbeQuicklyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFlowCenterLogsV1Request struct {
	*tchttp.BaseRequest

	// 关系&nbsp;id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 端口号

	Port *string `json:"Port,omitempty" name:"Port"`
	// 源地址

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 目标地址

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 时间格式Type&nbsp;0七天&nbsp;1&nbsp;今天&nbsp;5&nbsp;小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeVpcFlowCenterLogsV1Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFlowCenterLogsV1Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnterpriseSecurityGroupRuleRuleInfo struct {

	// 标签

	AclTags *string `json:"AclTags,omitempty" name:"AclTags"`
	// 自动化任务信息

	BetaList []*EnterpriseSecurityGroupRuleBetaInfo `json:"BetaList,omitempty" name:"BetaList"`
	// 目的资产网段信息

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则id&nbsp;&nbsp;等同RuleUuid

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 目的资产名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否延迟下发规则&nbsp;
	// 取值范围：0/1
	// 0:立即下发&nbsp;1:延迟下发

	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`
	// 规则最新一次是否有改动
	// 取值范围：0/1
	// 0:否
	// 1:是

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 排序

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 目的模板名称

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 目的资产内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 协议名称
	// 取值范围:TCP/ANY/ICMP/UDP
	// ANY:表示所有
	//

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口模板名称

	ProtocolPortName *string `json:"ProtocolPortName,omitempty" name:"ProtocolPortName"`
	// 目的资产公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 主键id

	RuleUuid *uint64 `json:"RuleUuid,omitempty" name:"RuleUuid"`
	// 服务模板id

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 源资产网段信息

	SouCidr *string `json:"SouCidr,omitempty" name:"SouCidr"`
	// 源资产名称

	SouInstanceName *string `json:"SouInstanceName,omitempty" name:"SouInstanceName"`
	// 源模板名称

	SouParameterName *string `json:"SouParameterName,omitempty" name:"SouParameterName"`
	// 源资产内网ip

	SouPrivateIp *string `json:"SouPrivateIp,omitempty" name:"SouPrivateIp"`
	// 源资产公网ip

	SouPublicIp *string `json:"SouPublicIp,omitempty" name:"SouPublicIp"`
	// 源规则内容

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 源规则类型&nbsp;
	// 取值范围&nbsp;0/1/2/3/4/5/6/7/8/9/100
	// 0表示ip(net),
	// 1表示VPC实例(intance)
	// 2表示子网实例(intance)
	// 3表示CVM实例(intance)
	// 4表示CLB实例(intance)
	// 5表示ENI实例(intance)
	// 6表示数据库实例(intance)
	// 7表示模版(template)
	// 8表示标签(tag)
	// 9表示地域(region)
	// 100表示资产分组(resourcegroup)

	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
	// 规则启用状态&nbsp;
	// 取值范围：&nbsp;0/1
	// 0:未开启
	// 1:开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则策略
	// 取值范围:1/2
	// 1:阻断
	// 2:放行

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 目的规则内容

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 目的规则类型&nbsp;
	// 取值范围&nbsp;0/1/2/3/4/5/6/7/8/9/100
	// 0表示ip(net),
	// 1表示VPC实例(intance)
	// 2表示子网实例(intance)
	// 3表示CVM实例(intance)
	// 4表示CLB实例(intance)
	// 5表示ENI实例(intance)
	// 6表示数据库实例(intance)
	// 7表示模版(template)
	// 8表示标签(tag)
	// 9表示地域(region)
	// 100表示资产分组(resourcegroup)

	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
	// 规则uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type InstanceDomainInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目的地址

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 最近访问时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 流量日志索引

	IndexType *string `json:"IndexType,omitempty" name:"IndexType"`
	// 资产实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资产实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 私有IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 7天内请求流量统计

	ResCnt *int64 `json:"ResCnt,omitempty" name:"ResCnt"`
	// 威胁情报风险评估

	Risk *int64 `json:"Risk,omitempty" name:"Risk"`
	// 7天内响应流量统计

	RspCnt *int64 `json:"RspCnt,omitempty" name:"RspCnt"`
	// 7天内会话数

	SessionCnt *int64 `json:"SessionCnt,omitempty" name:"SessionCnt"`
	// 7天内最小访问时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 威胁情报标签

	Tags *string `json:"Tags,omitempty" name:"Tags"`
	// 7天内总流量统计

	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeFlowCenterAddressListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据返回列表

		Data []*FlowCenterOutAddress `json:"Data,omitempty" name:"Data"`
		// 返回状态码&nbsp;0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowCenterAddressListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowCenterAddressListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSwitchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT边界防火墙开关列表数据

		Data []*NatSwitchListData `json:"Data,omitempty" name:"Data"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSwitchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnterpriseSecurityGroupRuleBetaInfo struct {

	// 时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type HoneyPotAttacker struct {

	// 攻击ip

	AttackerIp *string `json:"AttackerIp,omitempty" name:"AttackerIp"`
	// 是否封禁

	Block *int64 `json:"Block,omitempty" name:"Block"`
	// 告警次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 最近时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 攻击事件类型

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 是否放通

	Ignore *int64 `json:"Ignore,omitempty" name:"Ignore"`
	// 危险等级

	MaxLevel *string `json:"MaxLevel,omitempty" name:"MaxLevel"`
	// 蜜罐id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 蜜罐名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 威胁情报标签

	ThreatType *string `json:"ThreatType,omitempty" name:"ThreatType"`
	// 反制信息

	TraceCount *int64 `json:"TraceCount,omitempty" name:"TraceCount"`
}

type DeletePcapTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeletePcapTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePcapTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest

	// 云地域的英文简写

	Area *string `json:"Area,omitempty" name:"Area"`
	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// 所需要删除规则的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否删除反向规则，0：否，1：是

	IsDelReverse *uint64 `json:"IsDelReverse,omitempty" name:"IsDelReverse"`
}

func (r *DeleteSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpsRuleListRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeIpsRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpsRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIsolateTableRequest struct {
	*tchttp.BaseRequest

	// 操作动作：编辑、删除

	ButtonAction *string `json:"ButtonAction,omitempty" name:"ButtonAction"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 资产唯一id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *ModifyIsolateTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIsolateTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipBindInfo struct {

	// 绑定方式，1，独占出口ip；0，不独占ip

	BindType *string `json:"BindType,omitempty" name:"BindType"`
	// 出口Eip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// havip，上述eip绑定的havip

	HaVip *string `json:"HaVip,omitempty" name:"HaVip"`
}

type ZeroTrustRuleHitDetailItem struct {

	// 租户id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 命令阻断acl&nbsp;id

	CommandGroupId *string `json:"CommandGroupId,omitempty" name:"CommandGroupId"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则是否开启，0未开启，1开启

	Disable *uint64 `json:"Disable,omitempty" name:"Disable"`
	// 类型,1是ip,2是url,3是http域名

	DstContent *string `json:"DstContent,omitempty" name:"DstContent"`
	// 目的资产名称

	DstName *string `json:"DstName,omitempty" name:"DstName"`
	// 资源类型,资产实例ins,WEB服务web,数据库db,资产分组asset_group,资源标签tag

	DstType *string `json:"DstType,omitempty" name:"DstType"`
	// 过期时间

	Expire *string `json:"Expire,omitempty" name:"Expire"`
	// 权限id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 规则有效性

	Invalid *uint64 `json:"Invalid,omitempty" name:"Invalid"`
	// 访问源，逗号分隔，用户类型为前缀，例如：&nbsp;wx_user:1111,ios_user:2222,ioa_group:33333

	SrcContent *string `json:"SrcContent,omitempty" name:"SrcContent"`
	// 来源身份名称

	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`
	// 用户类型,微信用户wx_user,ioa用户ioa_user,ioa用户组ioa_group,ip/cidr&nbsp;net

	SrcType *string `json:"SrcType,omitempty" name:"SrcType"`
	// 状态，0:未删除,1:删除,2:编辑删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 规则组id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeWebAssetScanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*WebAssetsInfo `json:"Data,omitempty" name:"Data"`
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 数据集合总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域集合

		ZoneList []*WebAssetsZoneInfo `json:"ZoneList,omitempty" name:"ZoneList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebAssetScanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetScanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeIpSwitch struct {

	// 创建私有连接指定IP

	EndpointIp *string `json:"EndpointIp,omitempty" name:"EndpointIp"`
	// 公网IP

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// vpc&nbsp;中第一个EIP开关打开，需要指定子网创建私有连接

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 0&nbsp;:&nbsp;旁路&nbsp;1&nbsp;:&nbsp;串行

	SwitchMode *int64 `json:"SwitchMode,omitempty" name:"SwitchMode"`
}

type ExportBindLst struct {

	// 防火墙实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 出口绑定公网ip返回值信息

	EipInfo []*EipInfo `json:"EipInfo,omitempty" name:"EipInfo"`
	// 绑定实例id，可以是cvmId,子网Id,vpcId等

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 绑定实例的名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 绑定实例类型：0,cvm类型；1,子网类型；2，vpc类型

	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// 所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type NatInstanceDetail struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 流量峰值

	MaxFlow *float64 `json:"MaxFlow,omitempty" name:"MaxFlow"`
	// 带宽

	NatWidth *float64 `json:"NatWidth,omitempty" name:"NatWidth"`
}

type ProtectCount struct {

	// 观察数量

	Alert *int64 `json:"Alert,omitempty" name:"Alert"`
	// 拦截数量

	Block *int64 `json:"Block,omitempty" name:"Block"`
	// 关闭数量

	Close *int64 `json:"Close,omitempty" name:"Close"`
	// 严格数量

	Strict *int64 `json:"Strict,omitempty" name:"Strict"`
}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageCenterSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 1代表开启，0代表关闭

		Switch *int64 `json:"Switch,omitempty" name:"Switch"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMessageCenterSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMessageCenterSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetStatic struct {

	// 拦截数据

	BlockNum *int64 `json:"BlockNum,omitempty" name:"BlockNum"`
	// 蜜罐欺骗次数

	DeceptionNum *int64 `json:"DeceptionNum,omitempty" name:"DeceptionNum"`
	// 告警数据

	IgnoreNum *int64 `json:"IgnoreNum,omitempty" name:"IgnoreNum"`
	// 入站流量峰值

	InFlow *float64 `json:"InFlow,omitempty" name:"InFlow"`
	// 网络访问次数

	NetNum *int64 `json:"NetNum,omitempty" name:"NetNum"`
	// 出站流量峰值

	OutFlow *float64 `json:"OutFlow,omitempty" name:"OutFlow"`
	// 时间

	StatTime *string `json:"StatTime,omitempty" name:"StatTime"`
}

type DescribeBorderProtectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 互联网边界防护列表数据

		Data []*BorderProtectDetail `json:"Data,omitempty" name:"Data"`
		// 状态值，0：查询成功，非0：查询失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态信息，success：查询成功，fail：查询失败

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 查询结果总数，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBorderProtectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBorderProtectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatAcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未过滤的总条数

		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`
		// nat访问控制列表数据

		Data []*DescAcItem `json:"Data,omitempty" name:"Data"`
		// 总条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatAcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserInfoOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseInstanceInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// MYSQL:6;REDIS:7;ES:10;MARIADB:11;KAFKA:12

	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 1是内网实例，0是公网实例

	IsInternalInstance *int64 `json:"IsInternalInstance,omitempty" name:"IsInternalInstance"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type Filters struct {

	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values *string `json:"Values,omitempty" name:"Values"`
}

type DescribeVpcFwGroupFlowStatRequest struct {
	*tchttp.BaseRequest

	// 选择结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 防火墙组ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 开关流量统计页分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开关流量统计页分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 选择开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeVpcFwGroupFlowStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupFlowStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEdgeIpSwitchWeightRequest struct {
	*tchttp.BaseRequest

	// ip&nbsp;列表权重修改

	IpLst []*EdgeIpSwitchWeight `json:"IpLst,omitempty" name:"IpLst"`
}

func (r *ModifyEdgeIpSwitchWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEdgeIpSwitchWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *ModifyFlowApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwConfigRequest struct {
	*tchttp.BaseRequest

	// 管理网络异常bypass，1，开启；0，关闭

	ManageNetworkDownBypass *int64 `json:"ManageNetworkDownBypass,omitempty" name:"ManageNetworkDownBypass"`
	// 防火墙满载可持续时间，单位s，超过该时间将bypass防火墙，为0时不做任何操作

	OverLoadLimitTime *int64 `json:"OverLoadLimitTime,omitempty" name:"OverLoadLimitTime"`
}

func (r *ModifyVpcFwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetOverviewEvent struct {

	// 公网资产总数

	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`
	// 待处理事件

	Event *int64 `json:"Event,omitempty" name:"Event"`
	// 暴露端口总数

	ExposedPortCount *int64 `json:"ExposedPortCount,omitempty" name:"ExposedPortCount"`
	// 入侵防御模式

	IdpStatus *int64 `json:"IdpStatus,omitempty" name:"IdpStatus"`
	// 新增资产数

	NewAsset *int64 `json:"NewAsset,omitempty" name:"NewAsset"`
	// 失陷主机

	OutNum *int64 `json:"OutNum,omitempty" name:"OutNum"`
	// 端口扫描次数

	PortScanCount *int64 `json:"PortScanCount,omitempty" name:"PortScanCount"`
	// 暴露公网ip总数

	PublicIpCount *int64 `json:"PublicIpCount,omitempty" name:"PublicIpCount"`
	// 暴露公网ip已防护总数

	PublicIpProtectCount *int64 `json:"PublicIpProtectCount,omitempty" name:"PublicIpProtectCount"`
	// 远程运维登陆

	RemoteLogin *int64 `json:"RemoteLogin,omitempty" name:"RemoteLogin"`
	// 待封禁端口总数

	UnBannedPortCount *int64 `json:"UnBannedPortCount,omitempty" name:"UnBannedPortCount"`
	// 漏洞攻击次数

	VulAttachNum *int64 `json:"VulAttachNum,omitempty" name:"VulAttachNum"`
	// 暴露漏洞总数

	VulCount *int64 `json:"VulCount,omitempty" name:"VulCount"`
}

type RuleCountInfo struct {

	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域或实例名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DescribeNatIpConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP连接数统计列表

		Data []*IpConnections `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatIpConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatIpConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowLogData struct {

	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 终止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 流字节数

	FlowBytesNum *uint64 `json:"FlowBytesNum,omitempty" name:"FlowBytesNum"`
	// 流报文数

	FlowPackageNum *uint64 `json:"FlowPackageNum,omitempty" name:"FlowPackageNum"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 接收字节数

	ReceiveBytesNum *uint64 `json:"ReceiveBytesNum,omitempty" name:"ReceiveBytesNum"`
	// 接收报文数

	ReceivePackageNum *uint64 `json:"ReceivePackageNum,omitempty" name:"ReceivePackageNum"`
	// 发送字节数

	SendBytesNum *uint64 `json:"SendBytesNum,omitempty" name:"SendBytesNum"`
	// 发送报文数

	SendPackageNum *uint64 `json:"SendPackageNum,omitempty" name:"SendPackageNum"`
	// 访问源

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 源端口

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运营商

	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`
	// 访问目的

	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`
	// 目的端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
}

type ZeroTrustCommandRuleHitDetail struct {

	// 用户APPID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 生效范围资产，全部资产，Id填ALL

	Assets []*ZeroTrustEffectiveScopeItem `json:"Assets,omitempty" name:"Assets"`
	// 命令详情

	Commands *string `json:"Commands,omitempty" name:"Commands"`
	// 描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则是否开启，0未开启，1开启

	Disable *int64 `json:"Disable,omitempty" name:"Disable"`
	// 命中次数

	HitTimes *int64 `json:"HitTimes,omitempty" name:"HitTimes"`
	// 策略ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
	// 状态，0:未删除,1:删除,2:编辑删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 策略类型，whitelist白名单，blacklist黑名单

	Type *string `json:"Type,omitempty" name:"Type"`
	// 生效范围用户，全部用户，Id填ALL

	Users []*ZeroTrustEffectiveScopeItem `json:"Users,omitempty" name:"Users"`
	// 策略UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeHoneyPotAttackerRequest struct {
	*tchttp.BaseRequest

	// 结束事件

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 探针id

	ProbeId *string `json:"ProbeId,omitempty" name:"ProbeId"`
	// 搜索条件

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 蜜罐id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 排序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 告警状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeHoneyPotAttackerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHoneyPotAttackerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceGroupChangeImpactNewRequest struct {
	*tchttp.BaseRequest

	// 资源组变更类型

	ResourceGroupChangeType *int64 `json:"ResourceGroupChangeType,omitempty" name:"ResourceGroupChangeType"`
	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
	// 资源组名称，名称变更场景必填

	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`
	// 资源组新父id，移动场景必填

	ResourceGroupNowParentId *string `json:"ResourceGroupNowParentId,omitempty" name:"ResourceGroupNowParentId"`
}

func (r *DescribeResourceGroupChangeImpactNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupChangeImpactNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressTemplateSpecification struct {

	// &nbsp;IP地址组ID，例如：ipmg-2uw6ujo6。

	AddressGroupId *string `json:"AddressGroupId,omitempty" name:"AddressGroupId"`
	// IP地址ID，例如：ipm-2uw6ujo6。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
}

type GetEventResult struct {

	// 无

	AffectAssetList []*GetEventAssetStruct `json:"AffectAssetList,omitempty" name:"AffectAssetList"`
	// 无

	AffectAssetNum *int64 `json:"AffectAssetNum,omitempty" name:"AffectAssetNum"`
	// 无

	AffectDescription *string `json:"AffectDescription,omitempty" name:"AffectDescription"`
	// 无

	AffectTrendList []*GetEventTrendStruct `json:"AffectTrendList,omitempty" name:"AffectTrendList"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 无

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 无

	EventTitle *string `json:"EventTitle,omitempty" name:"EventTitle"`
	// 无

	ExtraAffectTrend []*GetAssetInfoStruct `json:"ExtraAffectTrend,omitempty" name:"ExtraAffectTrend"`
	// 无

	IsDownload *bool `json:"IsDownload,omitempty" name:"IsDownload"`
	// 无

	Label []*string `json:"Label,omitempty" name:"Label"`
	// 无

	Level *string `json:"Level,omitempty" name:"Level"`
	// 无

	QRCode *string `json:"QRCode,omitempty" name:"QRCode"`
	// 无

	Related *int64 `json:"Related,omitempty" name:"Related"`
	// 无

	Suggest *string `json:"Suggest,omitempty" name:"Suggest"`
	// 无

	Time *string `json:"Time,omitempty" name:"Time"`
	// 无

	Url *string `json:"Url,omitempty" name:"Url"`
}

type WidthCompare struct {

	// 所属地域

	BelongRegion *string `json:"BelongRegion,omitempty" name:"BelongRegion"`
	// 实例id

	CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
	// 新带宽值

	NewWidth *int64 `json:"NewWidth,omitempty" name:"NewWidth"`
	// 原有带宽值

	OriginWidth *int64 `json:"OriginWidth,omitempty" name:"OriginWidth"`
}

type DeleteVpcFwInstanceRequest struct {
	*tchttp.BaseRequest

	// Vpc间防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *DeleteVpcFwInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcFwInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageHistogram struct {

	// 访问控制日志存储量，单位B

	AclSize *int64 `json:"AclSize,omitempty" name:"AclSize"`
	// 入侵防御日志存储量，单位B

	IdsSize *int64 `json:"IdsSize,omitempty" name:"IdsSize"`
	// 流量日志存储量，单位B

	NetFlowSize *int64 `json:"NetFlowSize,omitempty" name:"NetFlowSize"`
	// 操作日志存储量，单位B

	OperateSize *int64 `json:"OperateSize,omitempty" name:"OperateSize"`
	// 统计时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeAssetScanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*AssetFilterInfo `json:"Data,omitempty" name:"Data"`
		// 0

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 地域集合

		ZoneList []*AssetZone `json:"ZoneList,omitempty" name:"ZoneList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetScanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetScanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIsolateNetCardRequest struct {
	*tchttp.BaseRequest

	// 查询资产实例列表

	InstanceIDList []*string `json:"InstanceIDList,omitempty" name:"InstanceIDList"`
}

func (r *DescribeIsolateNetCardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIsolateNetCardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchCommonResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySwitchCommonResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchCommonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertLink struct {

	// 链接地址

	Link *string `json:"Link,omitempty" name:"Link"`
	// 图标类别

	LinkIcon *string `json:"LinkIcon,omitempty" name:"LinkIcon"`
	// 文案

	LinkTitle *string `json:"LinkTitle,omitempty" name:"LinkTitle"`
}

type DescribeAllZoneListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAllZoneListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllZoneListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGuideIdpsInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGuideIdpsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGuideIdpsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InteractionTypeTemplate struct {

	// 当前交互类型下的诱饵选择项

	BaitTemplateLst []*BaitTemplate `json:"BaitTemplateLst,omitempty" name:"BaitTemplateLst"`
	// 蜜罐交互类型选择

	InteractionType *string `json:"InteractionType,omitempty" name:"InteractionType"`
	// 支持部署模式，1:公网IP探针，2:负载均衡探针，3:内网IP探针

	SupportDeployMode *string `json:"SupportDeployMode,omitempty" name:"SupportDeployMode"`
}

type DescribeAsyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务信息

		Data *AsyncTask `json:"Data,omitempty" name:"Data"`
		// 0

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAsyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportIpsLogsRequest struct {
	*tchttp.BaseRequest

	// 资产id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 方向，0：入站，1：出站

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 结束日期

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 威胁类型名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 危险等级,&nbsp;提示->信息，中危->低，高危->高，极高危->极高

	Level *string `json:"Level,omitempty" name:"Level"`
	// 排序方向

	Order *string `json:"Order,omitempty" name:"Order"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 判断来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 开始日期

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 0:&nbsp;观察，1：阻断，2：放行

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *DescribeExportIpsLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportIpsLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出文件获取链接

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 返回状态码&nbsp;0&nbsp;成功&nbsp;非0不成功

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 返回信息&nbsp;&nbsp;success&nbsp;成功&nbsp;其他&nbsp;不成功

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 导出状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetNatProbeEipTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务完成情况
		// 0&nbsp;成功&nbsp;1&nbsp;拨测中

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 任务结果
		// 0待拨测&nbsp;1&nbsp;链路异常&nbsp;2健康&nbsp;3&nbsp;部分故障&nbsp;&nbsp;4检测超时

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetNatProbeEipTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetNatProbeEipTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BarData struct {

	// 名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type Filter struct {

	// 过滤类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeTLogIpListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据集合

		Data []*StaticInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTLogIpListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTLogIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeInfo struct {

	// 边ID

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 边名称

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 是否连接互联网&nbsp;1是&nbsp;0不是

	ToInternet *int64 `json:"ToInternet,omitempty" name:"ToInternet"`
}

type DeleteProbeTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteProbeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProbeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatAclRuleSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回修改成功的规则ID列表

		RuleUuids []*int64 `json:"RuleUuids,omitempty" name:"RuleUuids"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatAclRuleSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatAclRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceAclRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySequenceAclRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySequenceAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZeroTrustAccessSourceStat struct {

	// 接入类型

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 接入类型名称

	AccessTypeName *string `json:"AccessTypeName,omitempty" name:"AccessTypeName"`
	// 接入组织数量

	GroupNum *int64 `json:"GroupNum,omitempty" name:"GroupNum"`
	// 接入用户数量

	UserNum *int64 `json:"UserNum,omitempty" name:"UserNum"`
}

type ExportNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码&nbsp;0为成功&nbsp;其他值失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 任务id

		ReturnMd5 *string `json:"ReturnMd5,omitempty" name:"ReturnMd5"`
		// 消息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportNatFwDnatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据

	CoverDuplicate *int64 `json:"CoverDuplicate,omitempty" name:"CoverDuplicate"`
	// RuleType:&nbsp;1黑名单&nbsp;2外部IP&nbsp;3域名&nbsp;4情报&nbsp;5资产&nbsp;6自定义规则&nbsp;&nbsp;7入侵防御规则

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 非自定义类型规则列表

	Rules []*BanAndAllowRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateBlockIgnoreRuleNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwReSelectRequest struct {
	*tchttp.BaseRequest

	// vpc间防火墙实例id

	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 新增部署地域信息

	FwDeploy []*FwDeploy `json:"FwDeploy,omitempty" name:"FwDeploy"`
	// 模式&nbsp;2跨租户&nbsp;

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 私有网络模式重新接入的vpc列表

	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
}

func (r *ModifyVpcFwReSelectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwReSelectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNatFwObjResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNatFwObjResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatFwObjResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwVpcStatRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙组ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙的实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 选择的时间区间&nbsp;1:1小时内，2:&nbsp;24&nbsp;小时&nbsp;，3：7天，&nbsp;4：1个月

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeVpcFwVpcStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwVpcStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回错误码，0请求成功&nbsp;&nbsp;非0失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 接口返回信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVPCSwitchStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartFwInsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartFwInsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartFwInsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwGroupLineChar struct {

	// 均值带宽

	FlowAvg *float64 `json:"FlowAvg,omitempty" name:"FlowAvg"`
	// 峰值流量

	FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
	// 防火墙(组)ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙实例ID

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 时间

	UpTime *string `json:"UpTime,omitempty" name:"UpTime"`
}

type AccessSubsetDomainInfo struct {

	// 访问统计

	AccessStat []*RemoteAccessStat `json:"AccessStat,omitempty" name:"AccessStat"`
	// 访问次数

	AccessTimes *int64 `json:"AccessTimes,omitempty" name:"AccessTimes"`
	// 解析地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 创建状态

	CreateStatus *int64 `json:"CreateStatus,omitempty" name:"CreateStatus"`
	// 服务域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 修改时间

	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
	// 状态，1已接入，0未接入

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAssetEventTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Data []*AssetWarnEventData `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetEventTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEventTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLogEdgeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcLogEdgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLogEdgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetFilterInfo struct {

	// 无

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 无

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 无

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 无

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 无

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 无

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeProbeListRequest struct {
	*tchttp.BaseRequest

	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索参数，json格式字符串，空则传"{}"，部署模式：deploy_mode(1:EIP，2:CLB)，地域：region，探针Id：probe_id，探针名称：probe_name，地址：address，开关：switch，状态：status，所属私有网络Id：vpc_id，所属私有网络名称：vpc_name，部署实例Id：instance_id，部署实例名称：instance_name，转发器Id：proxy_id，转发器名称：proxy_name，模糊搜索：common

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 蜜罐服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 排序参数，转发到蜜罐数量从小到大honeypot_count:asc，转发到蜜罐数量从大到小honeypot_count:desc，命中次数从小到大hit_times:asc，命中次数从大到小hit_times:desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeProbeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProbeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTreatInfoStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTreatInfoStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTreatInfoStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetFilterListRequest struct {
	*tchttp.BaseRequest

	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 资产类型

	InsType *string `json:"InsType,omitempty" name:"InsType"`
	// 页大小&nbsp;0查询全部

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 资产组id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeAssetFilterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetFilterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSubnetStatLstResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 单台cvm峰值带宽统计列表

		Data []*NatSubnetStatStruct `json:"Data,omitempty" name:"Data"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 总行数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSubnetStatLstResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSubnetStatLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEWRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 规则方向，0：出站，1：入站，默认1

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
	// vpc规则必填，边id

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// 更改的规则当前执行顺序

	RuleSequence *uint64 `json:"RuleSequence,omitempty" name:"RuleSequence"`
	// 规则类型，vpc：VPC间规则、nat：Nat边界规则

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 是否开关开启，0：未开启，1：开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEWRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEWRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalAccountsPage struct {

	// 获取账号列表响应的单个对象

	Items []*DescribeLocalAccountsData `json:"Items,omitempty" name:"Items"`
	// 公共分页对象

	Page *Paging `json:"Page,omitempty" name:"Page"`
}

type Link struct {

	// 图标类型

	LinkIcon *string `json:"LinkIcon,omitempty" name:"LinkIcon"`
	// 外链接文案

	LinkOutChain *string `json:"LinkOutChain,omitempty" name:"LinkOutChain"`
	// 链接文案

	LinkText *string `json:"LinkText,omitempty" name:"LinkText"`
}

type DescribeCfwUpdateStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// nat防火墙升级状态

		NatFwUpdateStatus *CfwUpdateStatus `json:"NatFwUpdateStatus,omitempty" name:"NatFwUpdateStatus"`
		// vpc间防火墙升级状态

		VpcFwUpdateStatus *CfwUpdateStatus `json:"VpcFwUpdateStatus,omitempty" name:"VpcFwUpdateStatus"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwUpdateStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwUpdateStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGwJoinFwStatusRequest struct {
	*tchttp.BaseRequest

	// nat网关id

	NatGwId *string `json:"NatGwId,omitempty" name:"NatGwId"`
}

func (r *DescribeNatGwJoinFwStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGwJoinFwStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebAssetsInfo struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip:port

	Host *string `json:"Host,omitempty" name:"Host"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// webid

	WebId *string `json:"WebId,omitempty" name:"WebId"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ZeroTrustConfig struct {

	// 并发账号数规格

	AccountSpec *int64 `json:"AccountSpec,omitempty" name:"AccountSpec"`
	// 运维资产数

	AssetSpec *int64 `json:"AssetSpec,omitempty" name:"AssetSpec"`
	// 带宽规格

	BandwidthSpec *int64 `json:"BandwidthSpec,omitempty" name:"BandwidthSpec"`
	// 是否支持试用

	CanTrial *int64 `json:"CanTrial,omitempty" name:"CanTrial"`
	// 域名后缀

	DomainSuffix *string `json:"DomainSuffix,omitempty" name:"DomainSuffix"`
	// 购买扩展包数量

	ExtendPacket *int64 `json:"ExtendPacket,omitempty" name:"ExtendPacket"`
	// iOA映射状态,0无IOA权限，1已映射IOA，2已销毁映射

	IoaStatus *int64 `json:"IoaStatus,omitempty" name:"IoaStatus"`
	// 是否开白用户

	IsWhiteUser *int64 `json:"IsWhiteUser,omitempty" name:"IsWhiteUser"`
	// 地域接入数

	RegionSpec *int64 `json:"RegionSpec,omitempty" name:"RegionSpec"`
	// 零信任远程运维状态，0不可用，1可用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 是否支持数据库域名接入

	SupportAccessDbDomain *int64 `json:"SupportAccessDbDomain,omitempty" name:"SupportAccessDbDomain"`
	// 是否支持运维实例

	SupportCVM *int64 `json:"SupportCVM,omitempty" name:"SupportCVM"`
	// 是否支持数据库

	SupportDB *int64 `json:"SupportDB,omitempty" name:"SupportDB"`
	// 是否支持多运维ip

	SupportHA *int64 `json:"SupportHA,omitempty" name:"SupportHA"`
	// 是否支持Web服务

	SupportWeb *int64 `json:"SupportWeb,omitempty" name:"SupportWeb"`
	// 0存量免费版，1基础版，2高级运维版

	VersionType *int64 `json:"VersionType,omitempty" name:"VersionType"`
	// vpc接入数

	VpcSpec *int64 `json:"VpcSpec,omitempty" name:"VpcSpec"`
	// 微信账号数

	WechatSpec *int64 `json:"WechatSpec,omitempty" name:"WechatSpec"`
	// 试用过期时间

	WhiteExpireTime *string `json:"WhiteExpireTime,omitempty" name:"WhiteExpireTime"`
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest

	// 操作类型&nbsp;1&nbsp;置顶&nbsp;0取消

	OpeType *string `json:"OpeType,omitempty" name:"OpeType"`
	// 记录id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
}

func (r *ModifyBlockTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:&nbsp;修改成功,&nbsp;!0:&nbsp;修改失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySequenceRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLocalAccountsData struct {

	// 所属组

	AccountGroups []*DescribeLocalAccountAccountGroupsData `json:"AccountGroups,omitempty" name:"AccountGroups"`
	// 账号id，同Id字段

	AccountId *int64 `json:"AccountId,omitempty" name:"AccountId"`
	// 虚拟组

	AccountVirtualGroups []*DescribeLocalAccountAccountVirtualGroupsData `json:"AccountVirtualGroups,omitempty" name:"AccountVirtualGroups"`
	// 账号活跃状态&nbsp;1：活跃&nbsp;2：非活跃

	ActiveStatus *int64 `json:"ActiveStatus,omitempty" name:"ActiveStatus"`
	// 域id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 账号的扩展信息，包含邮箱、手机号、身份证、职位等信息

	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	// 账号所在的分组id

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 账号所在的分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// uid，数据库中唯一

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 节点路径

	IdPath *string `json:"IdPath,omitempty" name:"IdPath"`
	// 账号的创建时间

	Itime *string `json:"Itime,omitempty" name:"Itime"`
	// 账号登录时间

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 账号登出时间

	LogoutTime *string `json:"LogoutTime,omitempty" name:"LogoutTime"`
	// 绑定手机端设备数

	MobileBindNum *int64 `json:"MobileBindNum,omitempty" name:"MobileBindNum"`
	// 账号所在的分组名称路径，用英文.分割

	NamePath *string `json:"NamePath,omitempty" name:"NamePath"`
	// 账号在线状态&nbsp;1：在线&nbsp;2：离线

	OnlineStatus *int64 `json:"OnlineStatus,omitempty" name:"OnlineStatus"`
	// 绑定Pc端设备数

	PcBindNum *int64 `json:"PcBindNum,omitempty" name:"PcBindNum"`
	// 用户风险等级，枚举：none,&nbsp;low,&nbsp;middle,&nbsp;high

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 根节点

	RootGroupId *string `json:"RootGroupId,omitempty" name:"RootGroupId"`
	// 账号来源,0表示本地账号(只支持32位)

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 账号状态,0禁用，1启用(只支持32位)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 账号，登录账号

	UserID *string `json:"UserID,omitempty" name:"UserID"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 账号的最后更新时间

	Utime *string `json:"Utime,omitempty" name:"Utime"`
}

type ElasticBandWidth struct {

	// 最大带宽，单位Mbps

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 最小带宽，单位Mbps

	MinBandwidth *int64 `json:"MinBandwidth,omitempty" name:"MinBandwidth"`
}

type BindEdgeCFWRequest struct {
	*tchttp.BaseRequest

	// Edge设备ID

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *BindEdgeCFWRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindEdgeCFWRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeModuleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceTypeTemplate struct {

	// 蜜罐描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 服务类型的交互类型选项

	InteractionTypeLst []*InteractionTypeTemplate `json:"InteractionTypeLst,omitempty" name:"InteractionTypeLst"`
	// 蜜罐服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 是否是web类型蜜罐，0：非web类型;&nbsp;1：web类型,&nbsp;支持选择lb探针

	WebType *int64 `json:"WebType,omitempty" name:"WebType"`
}

type CreateNatFwInstanceWithDomainRequest struct {
	*tchttp.BaseRequest

	// 异地灾备&nbsp;1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备

	CrossAZone *int64 `json:"CrossAZone,omitempty" name:"CrossAZone"`
	// 如果要创建域名则必填

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 指定防火墙使用网段信息

	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitempty" name:"FwCidrInfo"`
	// 0不创建域名,1创建域名

	IsCreateDomain *int64 `json:"IsCreateDomain,omitempty" name:"IsCreateDomain"`
	// 模式&nbsp;1：接入模式；0：新增模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 防火墙实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。

	NatGwList []*string `json:"NatGwList,omitempty" name:"NatGwList"`
	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。

	NewModeItems *NewModeItems `json:"NewModeItems,omitempty" name:"NewModeItems"`
	// 带宽

	Width *int64 `json:"Width,omitempty" name:"Width"`
	// 主可用区，为空则选择默认可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 备可用区，为空则选择默认可用区

	ZoneBak *string `json:"ZoneBak,omitempty" name:"ZoneBak"`
}

func (r *CreateNatFwInstanceWithDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatFwInstanceWithDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetFlowDomainTopRequest struct {
	*tchttp.BaseRequest

	// 实例Id，不指定则传空

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 时间类型，0：7天，1：24小时

	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`
}

func (r *DescribeNetFlowDomainTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetFlowDomainTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebAssetFilterListRequest struct {
	*tchttp.BaseRequest

	// 模糊查询

	FuzzySearch *string `json:"FuzzySearch,omitempty" name:"FuzzySearch"`
	// 页大小&nbsp;0查询全部

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 地域

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeWebAssetFilterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebAssetFilterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescNatDnatRule struct {

	// NAT防火墙转发规则描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 所属防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 网络协议，可选值：TCP、UDP。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 是否被关联引用，如被远程运维使用

	IsReferenced *int64 `json:"IsReferenced,omitempty" name:"IsReferenced"`
	// 关联的nat网关Id

	NatGwId *string `json:"NatGwId,omitempty" name:"NatGwId"`
	// 内网地址。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 内网端口。

	PrivatePort *int64 `json:"PrivatePort,omitempty" name:"PrivatePort"`
	// 弹性IP。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 公网端口。

	PublicPort *int64 `json:"PublicPort,omitempty" name:"PublicPort"`
}

type DescribeResourceGroupAssociationRulesNewRequest struct {
	*tchttp.BaseRequest

	// 资源组id

	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
}

func (r *DescribeResourceGroupAssociationRulesNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceGroupAssociationRulesNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlockIPBySGSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表成功，其他代表错误

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 错误时返回错误原因

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlockIPBySGSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBlockIPBySGSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncFwOperateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncFwOperateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncFwOperateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatSwitchListData struct {

	// 开关是否异常,0:正常,1:异常

	Abnormal *int64 `json:"Abnormal,omitempty" name:"Abnormal"`
	// 云服务器个数

	CvmNum *uint64 `json:"CvmNum,omitempty" name:"CvmNum"`
	// 是否生效

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 列表ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// NAT网关ID

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// NAT防火墙实例ID

	NatInsId *string `json:"NatInsId,omitempty" name:"NatInsId"`
	// NAT防火墙实例名称

	NatInsName *string `json:"NatInsName,omitempty" name:"NatInsName"`
	// NAT网关名称

	NatName *string `json:"NatName,omitempty" name:"NatName"`
	// nat防火墙出口路由表id

	ORTableId *string `json:"ORTableId,omitempty" name:"ORTableId"`
	// nat防火墙出口路由表名称

	ORTableName *string `json:"ORTableName,omitempty" name:"ORTableName"`
	// 出口Snat&nbsp;Ip列表

	Ohavips []*string `json:"Ohavips,omitempty" name:"Ohavips"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 关联路由ID

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 关联路由名称

	RouteName *string `json:"RouteName,omitempty" name:"RouteName"`
	// 开关状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// IPv4&nbsp;CIDR

	SubnetCidr *string `json:"SubnetCidr,omitempty" name:"SubnetCidr"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 所属VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 所属VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type DescribeVpcFwGroupFlowStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 峰值流量(防火墙组&nbsp;或&nbsp;防火墙实例)

		FlowMax *int64 `json:"FlowMax,omitempty" name:"FlowMax"`
		// 折线图数据

		LineCharLst []*FwGroupLineChar `json:"LineCharLst,omitempty" name:"LineCharLst"`
		// 异步查询状态，1查询执行中，0查询已结束

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 开关流量统计数

		SwitchFlowLst []*FwGroupSwitchFlowStat `json:"SwitchFlowLst,omitempty" name:"SwitchFlowLst"`
		// 开关流量总数

		SwitchFlowTotal *int64 `json:"SwitchFlowTotal,omitempty" name:"SwitchFlowTotal"`
		// 防火墙(组)或防火墙实例的规格

		Width *int64 `json:"Width,omitempty" name:"Width"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwGroupFlowStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwGroupFlowStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsyncErrInfo struct {

	// 异常的错误acl集合

	ErrAcl []*ErrAclInfo `json:"ErrAcl,omitempty" name:"ErrAcl"`
	// 异常信息&nbsp;比如文件下载失败&nbsp;等&nbsp;

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type LogStorageTypeSetting struct {

	// 互联网边界防火墙流量日志，1写入，2停写

	BorderTraffic *int64 `json:"BorderTraffic,omitempty" name:"BorderTraffic"`
	// DNS流量日志，1写入，2停写

	DnsTraffic *int64 `json:"DnsTraffic,omitempty" name:"DnsTraffic"`
	// NAT边界防火墙流量日志，1写入，2停写

	NatTraffic *int64 `json:"NatTraffic,omitempty" name:"NatTraffic"`
	// 流量分析入向日志，1写入，2停写

	NtaInTraffic *int64 `json:"NtaInTraffic,omitempty" name:"NtaInTraffic"`
	// 流量分析出向日志，1写入，2停写

	NtaOutTraffic *int64 `json:"NtaOutTraffic,omitempty" name:"NtaOutTraffic"`
	// 流量分析内网日志，1写入，2停写

	NtaPrivateTraffic *int64 `json:"NtaPrivateTraffic,omitempty" name:"NtaPrivateTraffic"`
	// 流量分析日志，1写入，2停写

	NtaTraffic *int64 `json:"NtaTraffic,omitempty" name:"NtaTraffic"`
	// 远程运维行为监控日志，1写入，2停写

	OperateRemoteCommand *int64 `json:"OperateRemoteCommand,omitempty" name:"OperateRemoteCommand"`
	// 远程运维登录日志，1写入，2停写

	OperateRemoteOM *int64 `json:"OperateRemoteOM,omitempty" name:"OperateRemoteOM"`
	// Web服务访问日志，1写入，2停写

	OperateWebAccess *int64 `json:"OperateWebAccess,omitempty" name:"OperateWebAccess"`
	// 内网流量日志，1写入，2停写

	PrivateTraffic *int64 `json:"PrivateTraffic,omitempty" name:"PrivateTraffic"`
	// 互联网边界规则访问控制日志，1写入，2停写

	RuleACLBorder *int64 `json:"RuleACLBorder,omitempty" name:"RuleACLBorder"`
	// 企业安全组日志，1写入，2停写

	RuleACLSecurityGroup *int64 `json:"RuleACLSecurityGroup,omitempty" name:"RuleACLSecurityGroup"`
	// DNS规则访问控制日志，1写入，2停写

	RuleDnsACL *int64 `json:"RuleDnsACL,omitempty" name:"RuleDnsACL"`
	// 入侵防御日志-安全基线，1写入，2停写

	RuleThreatInfoBaseline *int64 `json:"RuleThreatInfoBaseline,omitempty" name:"RuleThreatInfoBaseline"`
	// 入侵防御日志-封禁列表，1写入，2停写

	RuleThreatInfoBlockList *int64 `json:"RuleThreatInfoBlockList,omitempty" name:"RuleThreatInfoBlockList"`
	// 入侵防御日志-网络蜜罐，1写入，2停写

	RuleThreatInfoHoneypot *int64 `json:"RuleThreatInfoHoneypot,omitempty" name:"RuleThreatInfoHoneypot"`
	// 入侵防御日志-基础防御，1写入，2停写

	RuleThreatInfoIDS *int64 `json:"RuleThreatInfoIDS,omitempty" name:"RuleThreatInfoIDS"`
	// 入侵防御日志-虚拟补丁，1写入，2停写

	RuleThreatInfoIPS *int64 `json:"RuleThreatInfoIPS,omitempty" name:"RuleThreatInfoIPS"`
	// 入侵防御日志-威胁情报，1写入，2停写

	RuleThreatInfoTI *int64 `json:"RuleThreatInfoTI,omitempty" name:"RuleThreatInfoTI"`
	// 数据库访问日志，1写入，2停写

	RuleVpcACLDb *int64 `json:"RuleVpcACLDb,omitempty" name:"RuleVpcACLDb"`
	// nat规则访问控制日志，1写入，2停写

	RuleVpcACLNat *int64 `json:"RuleVpcACLNat,omitempty" name:"RuleVpcACLNat"`
	// 内网间规则访问控制日志，1写入，2停写

	RuleVpcACLVpc *int64 `json:"RuleVpcACLVpc,omitempty" name:"RuleVpcACLVpc"`
	// VPC间防火墙流量日志，1写入，2停写

	VpcTraffic *int64 `json:"VpcTraffic,omitempty" name:"VpcTraffic"`
}

type NtaFileInfo struct {

	// 文件连接

	CosFile *string `json:"CosFile,omitempty" name:"CosFile"`
	// 文件发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 文件MD5

	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件大小

	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 流量唯一ID

	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
	// 防火墙类型

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 文件属性，如下值：
	// 1:&nbsp;安全,2:&nbsp;恶意,3:&nbsp;未知,4:&nbsp;感染,6:&nbsp;风险,7:&nbsp;可信

	SafeLevel *int64 `json:"SafeLevel,omitempty" name:"SafeLevel"`
	// 文件来源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
}

type DescribeFwEdgeBarRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFwEdgeBarRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwEdgeBarRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewNatCheckInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNewNatCheckInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewNatCheckInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站

	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeRuleOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetInfo struct {

	// 租户子CIDR

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// 租户子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 租户子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 租户VPCID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeDbOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据库总数

		DbNum *int64 `json:"DbNum,omitempty" name:"DbNum"`
		// 数据库类型实例

		DbTypeInstances []*DbInstance `json:"DbTypeInstances,omitempty" name:"DbTypeInstances"`
		// 数据库七日曲线图形

		DbVisitChart []*DbInstance `json:"DbVisitChart,omitempty" name:"DbVisitChart"`
		// 数据库访问top5

		DbVisitItems []*DbInstance `json:"DbVisitItems,omitempty" name:"DbVisitItems"`
		// 数据库未防护数

		NoDefendDbNum *int64 `json:"NoDefendDbNum,omitempty" name:"NoDefendDbNum"`
		// 风险数据库数

		RiskDbNum *int64 `json:"RiskDbNum,omitempty" name:"RiskDbNum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDbOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFwGroupInstanceInfoRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeFwGroupInstanceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFwGroupInstanceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteBackupLstRequest struct {
	*tchttp.BaseRequest

	// 排序所用到的字段

	By *string `json:"By,omitempty" name:"By"`
	// 检索的截止时间，可不传

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件组合

	Filters []*CommonFilter `json:"Filters,omitempty" name:"Filters"`
	// 防火墙实例id

	FwInsId *string `json:"FwInsId,omitempty" name:"FwInsId"`
	// 需要查询的索引，特定场景使用，可不填

	Index *string `json:"Index,omitempty" name:"Index"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值

	Order *string `json:"Order,omitempty" name:"Order"`
	// 检索的起始时间，可不传

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeRouteBackupLstRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteBackupLstRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcRuleItem struct {

	// beta任务详情

	BetaList []*BetaInfoByACL `json:"BetaList,omitempty" name:"BetaList"`
	// 规则被删除：1，已删除；0，未删除

	Deleted *int64 `json:"Deleted,omitempty" name:"Deleted"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.domain.com

	DestContent *string `json:"DestContent,omitempty" name:"DestContent"`
	// 访问目的类型，类型可以为：net，domain

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 规则的命中次数，增删改查规则时无需传入此参数，主要用于返回查询结果数据

	DetectedTimes *int64 `json:"DetectedTimes,omitempty" name:"DetectedTimes"`
	// 若访问目的类型是dnsparse，即协议解析类型时，该值有效

	DnsParseCnt *int64 `json:"DnsParseCnt,omitempty" name:"DnsParseCnt"`
	// 规则生效的范围，是在哪对vpc之间还是针对所有vpc间生效

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
	// EdgeId对应的这对VPC间防火墙的描述

	EdgeName *string `json:"EdgeName,omitempty" name:"EdgeName"`
	// 规则状态，true表示启用，false表示禁用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 规则生效的防火墙实例ID

	FwGroupId *string `json:"FwGroupId,omitempty" name:"FwGroupId"`
	// 防火墙名称

	FwGroupName *string `json:"FwGroupName,omitempty" name:"FwGroupName"`
	// 内部使用的uuid，一般情况下不会使用到该字段

	InternalUuid *int64 `json:"InternalUuid,omitempty" name:"InternalUuid"`
	// 是否是无效规则，0&nbsp;表示有效规则，1&nbsp;表示无效规则，出参场景返回使用

	Invalid *int64 `json:"Invalid,omitempty" name:"Invalid"`
	// Ip版本，0：IPv4，1：IPv6，默认为IPv4

	IpVersion *int64 `json:"IpVersion,omitempty" name:"IpVersion"`
	// 规则上次命中时间

	LastHitTime *string `json:"LastHitTime,omitempty" name:"LastHitTime"`
	// 规则顺序，-1表示最低，1表示最高

	OrderIndex *int64 `json:"OrderIndex,omitempty" name:"OrderIndex"`
	// 端口协议组ID

	ParamTemplateId *string `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
	// 端口协议组名称

	ParamTemplateName *string `json:"ParamTemplateName,omitempty" name:"ParamTemplateName"`
	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议，可选的值：
	// TCP
	// UDP
	// ICMP
	// ANY
	// HTTP
	// HTTPS
	// HTTP/HTTPS
	// SMTP
	// SMTPS
	// SMTP/SMTPS
	// FTP
	// DNS
	// TLS/SSL

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察

	RuleAction *string `json:"RuleAction,omitempty" name:"RuleAction"`
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)

	SourceContent *string `json:"SourceContent,omitempty" name:"SourceContent"`
	// 访问源名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 访问源类型，类型可以为：net

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 访问目的名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 规则对应的唯一id

	Uuid *int64 `json:"Uuid,omitempty" name:"Uuid"`
}

type RemoveNatFwTcRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveNatFwTcRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveNatFwTcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntArray struct {

	// 数组

	List []*int64 `json:"List,omitempty" name:"List"`
}

type RuleImpactNew struct {

	// 互联网

	InternetRule []*RuleImpactInfo `json:"InternetRule,omitempty" name:"InternetRule"`
	// nat

	NatRule []*RuleImpactInfo `json:"NatRule,omitempty" name:"NatRule"`
	// 新安全组

	NewSecurityGroupRule []*RuleImpactInfo `json:"NewSecurityGroupRule,omitempty" name:"NewSecurityGroupRule"`
	// 规则数

	RuleTotalCount *int64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
	// 安全组

	SecurityGroupRule []*RuleImpactInfo `json:"SecurityGroupRule,omitempty" name:"SecurityGroupRule"`
	// vpc

	VpcRule []*RuleImpactInfo `json:"VpcRule,omitempty" name:"VpcRule"`
	// 数据库

	WhiteDbRule []*RuleImpactInfo `json:"WhiteDbRule,omitempty" name:"WhiteDbRule"`
}

type DescribeBillingReleaseResourceCheckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0:未释放，1:释放

		FwSwitchStatus *uint64 `json:"FwSwitchStatus,omitempty" name:"FwSwitchStatus"`
		// 状态码0成功&nbsp;其他失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 状态值&nbsp;success

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingReleaseResourceCheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingReleaseResourceCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwSwitchRequest struct {
	*tchttp.BaseRequest

	// 获取实例开关列表过滤筛选字段

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 是否跨租户&nbsp;1跨租户&nbsp;0同租户

	IsPeer *int64 `json:"IsPeer,omitempty" name:"IsPeer"`
	// 每页长度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcFwSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineExportTask struct {

	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 日志总数

	DataLength *uint64 `json:"DataLength,omitempty" name:"DataLength"`
	// 导出失败信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 文件过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 下载进度

	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
	// 任务状态，0等待下载，1下载中，2下载完成，3下载失败，4文件过期

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 是否使用Cos

	UseUserCos *int64 `json:"UseUserCos,omitempty" name:"UseUserCos"`
}

type UserBandwidthUsage struct {

	// 引擎实例类型

	FwEnginType *string `json:"FwEnginType,omitempty" name:"FwEnginType"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 使用百分比

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 带宽规格&nbsp;单位是Mbps

	Quota *float64 `json:"Quota,omitempty" name:"Quota"`
	// 实例地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 使用带宽情况&nbsp;单位是Mbps

	UsedWidth *float64 `json:"UsedWidth,omitempty" name:"UsedWidth"`
}

type DescribeCfwInsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data []*CfwInsInfo `json:"Data,omitempty" name:"Data"`
		// Total

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfwInsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwInsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcFwRuleHitDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则命中详情

		VpcRuleDetail *RuleDetail `json:"VpcRuleDetail,omitempty" name:"VpcRuleDetail"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcFwRuleHitDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcFwRuleHitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatInstanceRequest struct {
	*tchttp.BaseRequest

	// NAT防火墙实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// NAT防火墙实例ID

	NatInstanceId *string `json:"NatInstanceId,omitempty" name:"NatInstanceId"`
}

func (r *ModifyNatInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVpcFwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetVpcFwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetVpcFwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncFwOperateRequest struct {
	*tchttp.BaseRequest

	// 防火墙类型；nat,nat防火墙;ew,vpc间防火墙

	FwType *string `json:"FwType,omitempty" name:"FwType"`
	// 同步操作类型：Route，同步防火墙路由

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
}

func (r *SyncFwOperateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncFwOperateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEventTrendStruct struct {

	// 无

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 无

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type CreateVpcFwInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 防火墙实例id

		CfwInsId *string `json:"CfwInsId,omitempty" name:"CfwInsId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcFwInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板列表数据

		Data []*TemplateListInfo `json:"Data,omitempty" name:"Data"`
		// 域名地址模板数量

		DomainTemplateCount *int64 `json:"DomainTemplateCount,omitempty" name:"DomainTemplateCount"`
		// Ip地址模板数量

		IpTemplateCount *int64 `json:"IpTemplateCount,omitempty" name:"IpTemplateCount"`
		// 模板名称列表

		NameList []*string `json:"NameList,omitempty" name:"NameList"`
		// 协议端口模板数量

		PortTemplateCount *int64 `json:"PortTemplateCount,omitempty" name:"PortTemplateCount"`
		// 地址模板配额数量

		TemplateQuotaCount *int64 `json:"TemplateQuotaCount,omitempty" name:"TemplateQuotaCount"`
		// 模板总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 已使用的地址模板数

		UsedTemplateCount *int64 `json:"UsedTemplateCount,omitempty" name:"UsedTemplateCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

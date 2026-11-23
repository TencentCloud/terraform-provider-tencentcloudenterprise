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

package v20180317

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeCertsRequest struct {
	*tchttp.BaseRequest

	// 数据偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回证书的数量，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 证书类型（目前支持:CA=客户端证书,SVR=服务器证书）

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 证书id 列表

	CertIdList []*string `json:"CertIdList,omitempty" name:"CertIdList"`
	// 是否同时获取证书内容

	WithCert *bool `json:"WithCert,omitempty" name:"WithCert"`
}

func (r *DescribeCertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的ID，当绑定后端服务到七层转发规则时，必须提供此参数或Domain+Url两者之一

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标转发规则的域名，提供LocationId参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标转发规则的URL，提供LocationId参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 待绑定的后端服务列表，数组长度最大支持20

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
}

func (r *RegisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionTarget struct {

	// 云函数相关信息

	Function *FunctionInfo `json:"Function,omitempty" name:"Function"`
	// 权重

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type HealthCheck struct {

	// 是否开启健康检查：1（开启）、0（关闭）。

	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`
	// 健康检查的响应超时时间（仅适用于四层监听器），可选值：2~60，默认值：2，单位：秒。响应超时时间要小于检查间隔时间。

	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 健康检查探测间隔时间，默认值：5，可选值：5~300，单位：秒。

	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`
	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2~10，单位：次。

	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`
	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发异常，可选值：2~10，单位：次。

	UnHealthNum *int64 `json:"UnHealthNum,omitempty" name:"UnHealthNum"`
	// 健康检查状态码（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。可选值：1~31，默认 31。
	// 1 表示探测后返回值 1xx 代表健康，2 表示返回 2xx 代表健康，4 表示返回 3xx 代表健康，8 表示返回 4xx 代表健康，16 表示返回 5xx 代表健康。若希望多种返回码都可代表健康，则将相应的值相加。注意：TCP监听器的HTTP健康检查方式，只支持指定一种健康检查状态码。

	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`
	// 健康检查路径（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。

	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`
	// 健康检查域名（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。

	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`
	// 健康检查方法（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式），默认值：HEAD，可选值HEAD或GET。

	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" name:"HttpCheckMethod"`
	// 自定义探测相关参数。健康检查端口，默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。（仅适用于TCP/UDP监听器）。

	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`
	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查的输入格式，可取值：HEX或TEXT；取值为HEX时，SendContext和RecvContext的字符只能在0123456789ABCDEF中选取且长度必须是偶数位。（仅适用于TCP/UDP监听器）

	ContextType *string `json:"ContextType,omitempty" name:"ContextType"`
	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查发送的请求内容，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。

	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`
	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查返回的结果，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。

	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`
	// 自定义探测相关参数。健康检查使用的协议：TCP | HTTP | CUSTOM（仅适用于TCP/UDP监听器，其中UDP监听器只支持CUSTOM；如果使用自定义健康检查功能，则必传）。

	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
	// 自定义探测相关参数。健康检查协议CheckType的值取HTTP时，必传此字段，代表后端服务的HTTP版本：HTTP/1.0、HTTP/1.1；（仅适用于TCP监听器）

	HttpVersion *string `json:"HttpVersion,omitempty" name:"HttpVersion"`
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest

	// 服务端证书的ID，或客户端证书的ID

	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

func (r *DescribeLoadBalancerListByCertIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerListByCertIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetNameVIP struct {

	// 集群名

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口号

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例的网络类型：
	// OPEN：公网属性，&nbsp;INTERNAL：内网属性。

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡实例的类型。1：通用的负载均衡实例，目前只支持传入1

	Forward *int64 `json:"Forward,omitempty" name:"Forward"`
	// 负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-50&nbsp;个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡后端目标设备所属的网络&nbsp;ID，如vpc-j9x73ko1，可以通过&nbsp;DescribeVpcEx&nbsp;接口获取。&nbsp;不传此参数则默认为基础网络（"0"）。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 在私有网络内购买内网负载均衡实例的情况下，必须指定子网&nbsp;ID，内网负载均衡实例的&nbsp;VIP&nbsp;将从这个子网中产生。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 负载均衡实例所属的项目&nbsp;ID，可以通过&nbsp;DescribeProject&nbsp;接口获取。不传此参数则视为默认项目。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 仅适用于公网负载均衡。IP版本，可取值：IPV4、IPV6、IPv6FullChain，默认值&nbsp;IPV4。

	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`
	// 创建负载均衡的个数，默认值&nbsp;1。

	Number *uint64 `json:"Number,omitempty" name:"Number"`
	// 设置跨可用区容灾时的主可用区ID，例如&nbsp;100001&nbsp;或&nbsp;ap-region1-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区。可通过&nbsp;DescribeMasterZones&nbsp;接口查询一个地域的主可用区的列表。

	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`
	// 仅适用于公网负载均衡。可用区ID，指定可用区以创建负载均衡实例。如：ap-region1-1。私有云暂不使用

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 仅适用于公网负载均衡。Anycast的发布域，可取&nbsp;ZONE_A&nbsp;或&nbsp;ZONE_B。仅带宽非上移用户支持此参数。（已下线）

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
	// 仅适用于公网负载均衡。负载均衡的网络计费模式。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 仅适用于公网负载均衡。CMCC&nbsp;|&nbsp;CTCC&nbsp;|&nbsp;CUCC&nbsp;|&nbsp;BGP，如果不指定本参数，则默认使用BGP。可通过&nbsp;DescribeIspInfo&nbsp;接口查询一个地域所支持的Isp。

	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`
	// 购买负载均衡同时，给负载均衡打上标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 是否支持直通。私有云不支持

	ZhiTong *bool `json:"ZhiTong,omitempty" name:"ZhiTong"`
	// 指定Vip申请负载均衡，必须同时指定&nbsp;TgwGroupName&nbsp;参数

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Tgw独占集群的名称

	TgwGroupName *string `json:"TgwGroupName,omitempty" name:"TgwGroupName"`
	// 是否可绑定高防包。私有云不支持

	IsDDos *bool `json:"IsDDos,omitempty" name:"IsDDos"`
	// 带宽包ID，网络计费方式选择带宽包时必须指定带宽包ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 独占集群信息

	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
	// 设置跨可用区容灾时的备可用区ID，例如&nbsp;100001&nbsp;或&nbsp;ap-region1-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，自动切换至备可用区。可通过&nbsp;DescribeMasterZones&nbsp;接口查询一个地域的主备可用区的列表。

	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`
	// 私有云专用参数，申请内网clb时，可传入四层集群标签，指定四层独占集群

	TgwSetLabels []*string `json:"TgwSetLabels,omitempty" name:"TgwSetLabels"`
	// 私有云专用参数，申请内网clb时，可传入七层集群标签，指定七层独占集群

	StgwSetLabels []*string `json:"StgwSetLabels,omitempty" name:"StgwSetLabels"`
	// EIP&nbsp;的唯一&nbsp;ID，形如：eip-11112222，仅适用于内网负载均衡绑定EIP。

	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`
	// 本地专用集群ID

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 是否从vpc侧申请v6地址，1表示跟vpc解耦，不从vpc侧申请地址，0表示不跟vpc解耦，从vpc侧申请地址

	IPv6DecoupleVpc *int64 `json:"IPv6DecoupleVpc,omitempty" name:"IPv6DecoupleVpc"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 新建转发规则的信息

	Rules []*RuleInput `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 running，2 fail，6 succ

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIPTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManualRewriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManualRewriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManualRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpLoadCertRequest struct {
	*tchttp.BaseRequest

	// 证书内容

	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`
	// 证书类型（目前支持:CA=客户端证书,SVR=服务器证书）

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 证书私钥

	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`
	// 证书备注

	CertAlias *string `json:"CertAlias,omitempty" name:"CertAlias"`
	// 国密签名证书内容

	SignCert *string `json:"SignCert,omitempty" name:"SignCert"`
	// 国密签名证书私钥

	SignCertPrivateKey *string `json:"SignCertPrivateKey,omitempty" name:"SignCertPrivateKey"`
}

func (r *UpLoadCertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpLoadCertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlockedIP struct {

	// 黑名单IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 加入黑名单的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type CertificateOutput struct {

	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证

	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`
	// 服务端证书的 ID。(兼容早期版本，默认RSA证书)

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 客户端证书的 ID。

	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`
	// 服务端SM证书的 ID。

	SmCertId *string `json:"SmCertId,omitempty" name:"SmCertId"`
	// 服务端ECC证书的 ID。

	EccCertId *string `json:"EccCertId,omitempty" name:"EccCertId"`
}

type ModifyDnatRequest struct {
	*tchttp.BaseRequest

	// clb的字符串id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 开关参数。1 开启直通，0 关闭直通。

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyDnatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 切换负载均衡计费方式时，可用此参数查询切换任务是否成功。

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancerAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest

	// 安全组ID，如 sg-12345678

	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
	// ADD 绑定安全组；
	// DEL 解绑安全组

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 负载均衡实例ID数组

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalanceListenerIds struct {

	// 负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器ID数组

	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type ExtraInfo struct {

	// 是否开通VIP直通

	ZhiTong *bool `json:"ZhiTong,omitempty" name:"ZhiTong"`
	// TgwGroup名称

	TgwGroupName *string `json:"TgwGroupName,omitempty" name:"TgwGroupName"`
}

type ProjectInfo struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目所有人

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建者

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 项目信息

	Info *string `json:"Info,omitempty" name:"Info"`
}

type RuleInput struct {

	// 转发规则的域名。长度限制为：1~80。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发规则的路径。长度限制为：1~200。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 会话保持时间。设置为0表示关闭会话保持，开启会话保持可取值30~3600，单位：秒。

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 健康检查信息

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 证书信息

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 负载均衡与后端服务之间的转发协议，目前支持 HTTP/HTTPS/TRPC

	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`
	// 是否将该域名设为默认域名，注意，一个监听器下只能设置一个默认域名。

	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。

	Http2 *bool `json:"Http2,omitempty" name:"Http2"`
	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// TRPC被调服务器路由，ForwardType为TRPC时必填

	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`
	// TRPC调用服务接口，ForwardType为TRPC时必填

	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

type DeleteCertRequest struct {
	*tchttp.BaseRequest

	// 证书ID 列表

	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

func (r *DeleteCertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlockIPTaskRequest struct {
	*tchttp.BaseRequest

	// ModifyBlockIPList 接口返回的异步任务的ID。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBlockIPTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监听器后端绑定的机器信息

		Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 询价的负载均衡类型，OPEN为公网类型，INTERNAL为内网类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 询价的收费类型，POSTPAID为按量计费，"PREPAID"为预付费包年包月

	LoadBalancerChargeType *string `json:"LoadBalancerChargeType,omitempty" name:"LoadBalancerChargeType"`
	// 询价的收费周期

	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitempty" name:"LoadBalancerChargePrepaid"`
	// 询价的网络计费方式

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 询价的负载均衡实例个数，默认为1

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
}

func (r *InquiryPriceCreateLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigListItem struct {

	// 配置ID

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
	// 配置类型

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置名字

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置内容

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
	// 增加配置时间

	CreateTimestamp *string `json:"CreateTimestamp,omitempty" name:"CreateTimestamp"`
	// 修改配置时间

	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" name:"UpdateTimestamp"`
}

type TargetGroupInstance struct {

	// 目标组实例的内网IP

	BindIP *string `json:"BindIP,omitempty" name:"BindIP"`
	// 目标组实例的端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 目标组实例的权重

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 目标组实例的新端口

	NewPort *uint64 `json:"NewPort,omitempty" name:"NewPort"`
}

type DescribeMasterZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区信息

		MasterZoneSet []*MasterZoneInfo `json:"MasterZoneSet,omitempty" name:"MasterZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMasterZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMasterZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetCountForLoadBalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询负载均衡实例，及其绑定的后端服务总数

		LoadBalancerSet []*TargetCountForLoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetCountForLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetCountForLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionInfo struct {

	// 函数命名空间

	FunctionNamespace *string `json:"FunctionNamespace,omitempty" name:"FunctionNamespace"`
	// 函数名称

	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
	// 函数的版本名称或别名

	FunctionQualifier *string `json:"FunctionQualifier,omitempty" name:"FunctionQualifier"`
	// 标识 FunctionQualifier 参数的类型，可取值： VERSION（版本）、ALIAS（别名）

	FunctionQualifierType *string `json:"FunctionQualifierType,omitempty" name:"FunctionQualifierType"`
}

type SetSubUinQuotasRequest struct {
	*tchttp.BaseRequest

	// 子账号配额列表

	SubQuota []*SubQuota `json:"SubQuota,omitempty" name:"SubQuota"`
}

func (r *SetSubUinQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSubUinQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LBChargePrepaid struct {

	// 续费类型：AUTO_RENEW 自动续费，  MANUAL_RENEW 手动续费

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 购买时长，单位：月

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type DescribeClusterResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群中资源列表

		ClusterResourceSet []*ClusterResource `json:"ClusterResourceSet,omitempty" name:"ClusterResourceSet"`
		// 集群中资源总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigListRequest struct {
	*tchttp.BaseRequest

	// 配置类型:CLB 负载均衡维度。 SERVER 域名维度。 LOCATION 规则维度。

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 拉取页偏移，默认值0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 拉取数目，默认值20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 拉取指定配置名字，模糊匹配。

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置ID

	UconfigIds []*string `json:"UconfigIds,omitempty" name:"UconfigIds"`
}

func (r *DescribeCustomizedConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomizedConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述价格信息

		Price []*Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlaItem struct {

	// 性能保障规格类型

	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
	// 性能保障规格名字

	SlaName *string `json:"SlaName,omitempty" name:"SlaName"`
	// 最大连接数

	MaxConn *uint64 `json:"MaxConn,omitempty" name:"MaxConn"`
	// 最大新增连接数

	MaxCps *uint64 `json:"MaxCps,omitempty" name:"MaxCps"`
	// 最大出带宽。单位：bps

	MaxOutBits *uint64 `json:"MaxOutBits,omitempty" name:"MaxOutBits"`
	// 最大入带宽。单位：bps

	MaxInBits *uint64 `json:"MaxInBits,omitempty" name:"MaxInBits"`
	// 最大QPS

	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

type TargetHealth struct {

	// Target的内网IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// Target绑定的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 当前健康状态，true：健康，false：不健康（包括尚未开始探测、探测中、状态异常等几种状态）。只有处于健康状态（且权重大于0），负载均衡才会向其转发流量。

	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// Target的实例ID，如 ins-12345678

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 当前健康状态的详细信息。如：Alive、Dead、Unknown。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知。

	HealthStatusDetial *string `json:"HealthStatusDetial,omitempty" name:"HealthStatusDetial"`
}

type InquiryPriceModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 修改后的网络带宽信息

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 性能独享型规格型号

	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
}

func (r *InquiryPriceModifyLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器

	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`
	// 监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL（TCP_SSL 正在内测中，如需使用请通过工单申请）

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数

	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`
	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 证书相关信息，此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL监听器。

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。

	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 监听器是否支持设置默认域名，默认为支持

	DefaultSerSwitch *bool `json:"DefaultSerSwitch,omitempty" name:"DefaultSerSwitch"`
}

func (r *CreateListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest

	// 目标组ID，与Filters互斥

	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`
	// 显示条数限制，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 显示的偏移起始量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件数组，与TargetGroupIds互斥，支持TargetGroupVpcId和TargetGroupName

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTargetGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该参数表示对应的价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIspInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运营商信息列表

		IspSet []*IspSet `json:"IspSet,omitempty" name:"IspSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIspInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIspInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancer struct {

	// 负载均衡实例&nbsp;ID。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例的名称。

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡实例的网络类型：
	// OPEN：公网属性，&nbsp;INTERNAL：内网属性。

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡类型标识，1：负载均衡，0：传统型负载均衡。

	Forward *uint64 `json:"Forward,omitempty" name:"Forward"`
	// 负载均衡实例的域名，仅公网传统型负载均衡实例才提供该字段

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 负载均衡实例的&nbsp;VIP&nbsp;列表。

	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`
	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 负载均衡实例的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 负载均衡实例的上次状态转换时间。

	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`
	// 负载均衡实例所属的项目&nbsp;ID，&nbsp;0&nbsp;表示默认项目。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 私有网络的&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 高防&nbsp;LB&nbsp;的标识，1：高防负载均衡&nbsp;0：非高防负载均衡。

	OpenBgp *uint64 `json:"OpenBgp,omitempty" name:"OpenBgp"`
	// 在&nbsp;2016&nbsp;年&nbsp;12&nbsp;月份之前的传统型内网负载均衡都是开启了&nbsp;snat&nbsp;的。

	Snat *bool `json:"Snat,omitempty" name:"Snat"`
	// 0：表示未被隔离，1：表示被隔离。

	Isolation *uint64 `json:"Isolation,omitempty" name:"Isolation"`
	// 用户开启日志的信息，日志只有公网属性创建了&nbsp;HTTP&nbsp;、HTTPS&nbsp;监听器的负载均衡才会有日志。

	Log *string `json:"Log,omitempty" name:"Log"`
	// 负载均衡实例所在的子网（仅对内网VPC型LB有意义）

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 负载均衡实例的标签信息

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 负载均衡实例的安全组

	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups"`
	// 负载均衡实例绑定的后端设备的基本信息

	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`
	// anycast负载均衡的发布域，对于非anycast的负载均衡，此字段返回为空字符串

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
	// IP版本，ipv4&nbsp;|&nbsp;ipv6

	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`
	// 数值形式的私有网络&nbsp;ID

	NumericalVpcId *uint64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`
	// 负载均衡IP地址所属的ISP

	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`
	// 主可用区

	MasterZone *ZoneInfo `json:"MasterZone,omitempty" name:"MasterZone"`
	// 备可用区

	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet"`
	// 负载均衡实例被隔离的时间

	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`
	// 负载均衡实例的过期时间，仅对预付费负载均衡生效

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 负载均衡实例的计费类型

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 负载均衡实例的网络属性

	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`
	// 负载均衡实例的预付费相关属性

	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitempty" name:"PrepaidAttributes"`
	// 负载均衡日志服务(CLS)的日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 负载均衡日志服务(CLS)的日志主题ID

	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`
	// 负载均衡实例的IPv6地址

	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`
	// 暂做保留，一般用户无需关注。

	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	// 是否可绑定高防包

	IsDDos *bool `json:"IsDDos,omitempty" name:"IsDDos"`
	// 负载均衡维度的个性化配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 后端服务是否放通来自LB的流量

	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
	// 内网独占集群

	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
	// IP地址版本为ipv6时此字段有意义，&nbsp;IPv6Nat64&nbsp;|&nbsp;IPv6FullChain

	IPv6Mode *string `json:"IPv6Mode,omitempty" name:"IPv6Mode"`
	// 是否开启SnatPro

	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`
	// 开启SnatPro负载均衡后，SnatIp列表

	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`
	// 性能保障规格

	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
	// vip是否被封堵

	IsBlock *bool `json:"IsBlock,omitempty" name:"IsBlock"`
	// 负载均衡IP地址所属的ISP的ID

	VipIspId *int64 `json:"VipIspId,omitempty" name:"VipIspId"`
	// 负载均衡IP地址所属的ISP

	VipIspName *string `json:"VipIspName,omitempty" name:"VipIspName"`
	// 属性标签列表

	AttributeFlags []*string `json:"AttributeFlags,omitempty" name:"AttributeFlags"`
	// 私有云专用，clb实例绑定的四层集群标签

	TgwSetLabels []*string `json:"TgwSetLabels,omitempty" name:"TgwSetLabels"`
	// 私有云专用，clb实例绑定的七层集群标签

	StgwSetLabels []*string `json:"StgwSetLabels,omitempty" name:"StgwSetLabels"`
	// 私有网络内网负载均衡，就近接入模式下规则所落在的可用区

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// CLB是否为NFV，空：不是，l7nfv：七层是NFV。

	NfvInfo *string `json:"NfvInfo,omitempty" name:"NfvInfo"`
	// 封堵或解封时间

	IsBlockTime *string `json:"IsBlockTime,omitempty" name:"IsBlockTime"`
	// 负载均衡日志服务(CLS)的健康检查日志集ID

	HealthLogSetId *string `json:"HealthLogSetId,omitempty" name:"HealthLogSetId"`
	// 负载均衡日志服务(CLS)的健康检查日志主题ID

	HealthLogTopicId *string `json:"HealthLogTopicId,omitempty" name:"HealthLogTopicId"`
	// IP类型是否是本地BGP

	LocalBgp *bool `json:"LocalBgp,omitempty" name:"LocalBgp"`
	// 开启IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标功能。

	MixIpTarget *bool `json:"MixIpTarget,omitempty" name:"MixIpTarget"`
	// 7层独占标签。

	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`
	// 7层独占集群ID。

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 负载均衡IP地址所属的ISP别名

	VipIspAlias *string `json:"VipIspAlias,omitempty" name:"VipIspAlias"`
}

type ModifyDomainAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoadBalancerLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersForVpcRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx。兼容整型的VpcId。

	VpcId []*string `json:"VpcId,omitempty" name:"VpcId"`
	// 负载均衡实例所属子网唯一ID，暂时只支持整型SubnetId。
	// 后续会兼容字符串的子网ID（subnet-1122ccbb）。

	SubnetId []*string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 数据偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回负载均衡实例的数量，默认为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLoadBalancersForVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersForVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重定向转发规则构成的数组，若无重定向规则，则返回空数组

		RewriteSet []*RuleOutput `json:"RewriteSet,omitempty" name:"RewriteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRewriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainAttributesRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 域名（必须是已经创建的转发规则下的域名）

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 要修改的新域名

	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`
	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用。

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。

	Http2 *bool `json:"Http2,omitempty" name:"Http2"`
	// 是否设为默认域名，注意，一个监听器下只能设置一个默认域名。

	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// 监听器下必须配置一个默认域名，若要关闭原默认域名，必须同时指定另一个域名作为新的默认域名。

	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

func (r *ModifyDomainAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAttributes struct {

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 是否设为默认域名，默认值

	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// 是否开启Http2，默认值

	Http2 *bool `json:"Http2,omitempty" name:"Http2"`
	// 监听器开启SNI的情况下，此域名所关联的证书。

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
}

type InternetAccessible struct {

	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;
	// BANDWIDTH_PACKAGE 按带宽包计费（当前，只有指定运营商时才支持此种计费模式）

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 最大出带宽，单位Mbps，范围支持0到65535，仅对公网属性的LB生效，默认值 10

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 带宽包的类型，如SINGLEISP

	BandwidthpkgSubType *string `json:"BandwidthpkgSubType,omitempty" name:"BandwidthpkgSubType"`
}

type LoadBalancerIdPair struct {

	// 源负载均衡ID

	SrcId *string `json:"SrcId,omitempty" name:"SrcId"`
	// 目的负载均衡ID

	DstId *string `json:"DstId,omitempty" name:"DstId"`
}

type MasterZoneInfo struct {

	// 可用区数值形式的唯一ID，如：100001

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区字符串形式的唯一ID，如：ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称，如：广州一区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 备可用区列表

	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet"`
}

type ItemPrice struct {

	// 后付费单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后续计价单元，可取值范围：
	// HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）；
	// GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预支费用的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 后付费的折扣单价，单位:元

	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`
	// 折扣，如20.0代表2折。

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
}

type LbRsItem struct {

	// vpc的字符串id，只支持字符串id。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 需要查询后端的内网ip，可以是cvm和弹性网卡。

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID，格式如 lb-12345678

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器 ID，格式如 lbl-12345678

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的ID，格式如 loc-12345678，当从七层转发规则解绑机器时，必须提供此参数或Domain+Url两者之一

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标规则的域名，提供LocationId参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标规则的URL，提供LocationId参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 要解绑的后端服务列表，数组长度最大支持20

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
}

func (r *DeregisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLBListenersRequest struct {
	*tchttp.BaseRequest

	// 需要查询的内网ip列表

	Backends []*LbRsItem `json:"Backends,omitempty" name:"Backends"`
}

func (r *DescribeLBListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLBListenersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID。实例ID数量上限为20个。

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
	// 负载均衡实例的网络类型：
	// OPEN：公网属性，&nbsp;INTERNAL：内网属性。

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 负载均衡实例的类型。1：通用的负载均衡实例。如果不传此参数，则查询所有类型的负载均衡实例。

	Forward *int64 `json:"Forward,omitempty" name:"Forward"`
	// 负载均衡实例的名称。

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 云平台为负载均衡实例分配的域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 负载均衡实例的&nbsp;VIP&nbsp;地址，支持多个。

	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`
	// 负载均衡绑定的后端服务的外网&nbsp;IP。

	BackendPublicIps []*string `json:"BackendPublicIps,omitempty" name:"BackendPublicIps"`
	// 负载均衡绑定的后端服务的内网&nbsp;IP。

	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`
	// 数据偏移量，默认为&nbsp;0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回负载均衡实例的数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数，支持以下字段：LoadBalancerName，CreateTime，Domain，LoadBalancerType。LoadBalancerName表示lb实例名称，CreateTime表示创建时间，Domain表示域名，LoadBalancerType表示实例类型。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 1：倒序，0：顺序，默认按照创建时间倒序。

	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`
	// 搜索字段，模糊匹配名称、域名、VIP。

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
	// 负载均衡实例所属的项目&nbsp;ID，可以通过&nbsp;DescribeProject&nbsp;接口获取。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。

	WithRs *int64 `json:"WithRs,omitempty" name:"WithRs"`
	// 负载均衡实例所属私有网络唯一ID，如&nbsp;vpc-bhqkbhdx。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 安全组ID，如&nbsp;sg-m1cc9123

	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
	// 主可用区ID，如&nbsp;："100001"&nbsp;

	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。详细的过滤条件如下：
	// <li>&nbsp;internet-charge-type&nbsp;-&nbsp;String&nbsp;-&nbsp;是否必填：否&nbsp;-&nbsp;（过滤条件）按照&nbsp;CLB&nbsp;的网络计费模式过滤，包括"BANDWIDTH_PREPAID","TRAFFIC_POSTPAID_BY_HOUR","BANDWIDTH_POSTPAID_BY_HOUR","BANDWIDTH_PACKAGE"。</li><li>&nbsp;master-zone-id&nbsp;-&nbsp;String&nbsp;-&nbsp;是否必填：否&nbsp;-&nbsp;（过滤条件）按照&nbsp;CLB&nbsp;的主可用区ID过滤，如&nbsp;："100001"&nbsp;。</li><li>&nbsp;tag-key&nbsp;-&nbsp;String&nbsp;-&nbsp;是否必填：否&nbsp;-&nbsp;（过滤条件）按照&nbsp;CLB&nbsp;标签的键过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest

	// 要查询的负载均衡实例 ID列表

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancerPriceInfo struct {

	// 后付费价格，单位为元，

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后付费计费周期

	ChargePeriod *string `json:"ChargePeriod,omitempty" name:"ChargePeriod"`
	// 预付费折前价格

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预付费折后价格

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type DescribeSingleIspResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMCC：移动、CTCC：电信、CUCC：联通

		IspSet []*string `json:"IspSet,omitempty" name:"IspSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSingleIspResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSingleIspResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteListSupportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取到的白名单信息

		WhiteLists []*CLBWhiteList `json:"WhiteLists,omitempty" name:"WhiteLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteListSupportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteListSupportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizedConfig struct {

	// 个性化配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 个性化配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 个性化配置内容

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
	// 个性化配置绑定的负载均衡

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
	// 配置的添加时间

	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
	// 配置的最后更新时间

	ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
}

type AssociateCustomizedConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateCustomizedConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateCustomizedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的负载均衡实例数组。

		LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubUinQuotasRequest struct {
	*tchttp.BaseRequest

	// 子账号UIN列表

	SubUin []*string `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *DescribeSubUinQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubUinQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindConfigItem struct {

	// 配置ID

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
	// 配置绑定的CLB ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 配置绑定的监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 配置绑定的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 配置绑定的规则

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type Target struct {

	// 后端服务的类型，可取：CVM（云服务器）、ENI（弹性网卡）、BMS（裸金属服务器）；作为入参时，目前本参数暂不生效。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 绑定CVM时需要传入此参数，代表CVM的唯一 ID，可通过 DescribeInstances 接口返回字段中的 InstanceId 字段获取。
	// 注意：参数 InstanceId 、 EniIp 、InstanceBmsId 只能传入一个且必须传入一个。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 后端服务的监听端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 绑定弹性网卡时需要传入此参数，代表弹性网卡的IP，弹性网卡必须先绑定至CVM，然后才能绑定到负载均衡实例。注意：参数 InstanceId 、EniIp、InstanceBmsId只能传入一个且必须传入一个。注意：绑定弹性网卡需要先提交工单开白名单使用。

	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
	// 绑定BMS时需要传入此参数，代表BMS的唯一 ID，可通过BMS产品的DescribeInstances 接口返回字段中的 bmsId 字段获取。
	// 注意：参数 InstanceId 、 EniIp 、InstanceBmsId 只能传入一个且必须传入一个，不支持同时绑定bms和cvm或eni。

	InstanceBmsId *string `json:"InstanceBmsId,omitempty" name:"InstanceBmsId"`
}

type DescribeSetInnerNameRequest struct {
	*tchttp.BaseRequest

	// 集群类型。四层："L4"，七层：“L7”

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群名字

	SetNames []*string `json:"SetNames,omitempty" name:"SetNames"`
	// 集群ID

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeSetInnerNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetInnerNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancerLocationRule struct {

	// 负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 规则信息

	LocationRules []*LocationRule `json:"LocationRules,omitempty" name:"LocationRules"`
}

type DeregisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监听器列表

		Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners"`
		// 总的监听器个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterResource struct {

	// 集群唯一ID，如tgw-12345678。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// ip地址。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡唯一ID，如lb-12345678。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 资源是否闲置。

	Idle *string `json:"Idle,omitempty" name:"Idle"`
	// 集群名称。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type RsWeightRule struct {

	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的ID，七层规则时需要此参数，4层规则不需要

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标规则的域名，提供LocationId参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标规则的URL，提供LocationId参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 要修改权重的后端机器列表

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
	// 后端服务新的转发权重，取值范围：0~100。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type SnatIp struct {

	// 私有网络子网的唯一性id，如subnet-12345678

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// IP地址，如192.168.0.1

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribeAppIdLabelRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAppIdLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 请求ID，即接口返回的 RequestId 参数

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 订单id

	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LbRsTargets struct {

	// 内网ip类型。“cvm”或“eni”

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端实例的内网ip。

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 绑定后端实例的端口。

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// rs的vpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// rs的权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type Resource struct {

	// 运营商信息，如"CMCC", "CUCC", "CTCC", "BGP", "INTERNAL"。

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 运营商内具体资源信息，如"CMCC", "CUCC", "CTCC", "BGP", "INTERNAL"。

	Type []*string `json:"Type,omitempty" name:"Type"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要删除的监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVip struct {

	// 集群名字

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 已使用的vip

	UsedVip []*string `json:"UsedVip,omitempty" name:"UsedVip"`
	// 未使用vip

	UnuseVip []*string `json:"UnuseVip,omitempty" name:"UnuseVip"`
}

type TagInfo struct {

	// 标签的键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签的值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DeleteRewriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRewriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertList struct {

	// 所有者

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 证书来源

	From *string `json:"From,omitempty" name:"From"`
	// 产品中文类型

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 中文名称

	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`
	// 项目域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 证书备注或名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 证书状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 漏洞状态

	VulnerAbilityStatus *string `json:"VulnerAbilityStatus,omitempty" name:"VulnerAbilityStatus"`
	// 开始时间

	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`
	// 过期时间

	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`
	// 有效期

	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`
	// 入库时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 证书ID

	CertID *string `json:"CertID,omitempty" name:"CertID"`
	// 状态名称

	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`
	// 项目信息

	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitempty" name:"ProjectInfo"`
	// 证书内容

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// 国密证书签名内容

	SignCert *string `json:"SignCert,omitempty" name:"SignCert"`
	// 备用域名列表

	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`
	// 取值类型：rsa： ecc： sm2:  默认为 rsa

	CertAlgorithm *string `json:"CertAlgorithm,omitempty" name:"CertAlgorithm"`
}

type AutoRewriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AutoRewriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AutoRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的ID，当后端服务绑定到七层转发规则时，必须提供此参数或Domain+Url两者之一

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标规则的域名，提供LocationId参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标规则的URL，提供LocationId参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 要修改端口的后端服务列表

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
	// 后端服务绑定到监听器或转发规则的新端口

	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestEzioResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// aa

		Hello *string `json:"Hello,omitempty" name:"Hello"`
		// aa

		World *string `json:"World,omitempty" name:"World"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestEzioResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestEzioResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceAmount struct {

	// asd

	FunctionNum *int64 `json:"FunctionNum,omitempty" name:"FunctionNum"`
}

type TargetGroupAssociation struct {

	// 负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
}

type RewriteLocationMap struct {

	// 源转发规则ID

	SourceLocationId *string `json:"SourceLocationId,omitempty" name:"SourceLocationId"`
	// 重定向至的目标转发规则ID

	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`
	// 重定向状态码，可取值301,302,307

	RewriteCode *uint64 `json:"RewriteCode,omitempty" name:"RewriteCode"`
	// 重定向是否携带匹配的url，配置RewriteCode时必填

	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`
	// 源转发的域名，必须是SourceLocationId对应的域名，配置RewriteCode时必填

	SourceDomain *string `json:"SourceDomain,omitempty" name:"SourceDomain"`
}

type DescribeLBActionLimitRequest struct {
	*tchttp.BaseRequest

	// 负载均衡唯一ID列表

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
	// 负载均衡相关操作，当前仅支持"ModifyLBNetwork"

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeLBActionLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLBActionLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubUinQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账号下各类型及其子账号配额列表

		QuotaData []*QuotaData `json:"QuotaData,omitempty" name:"QuotaData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubUinQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubUinQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器下的某个旧域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 新域名，	长度限制为：1-120。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。

	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicalTargetInfo struct {

	// 后端实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 权重，取值范围 [0, 100]

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ListenerRule struct {

	// 健康健康配置

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 会话保持时间。设置为0表示关闭会话保持，开启会话保持可取值30~3600，单位：秒。

	SessionExpireTime *uint64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 待绑定的后端服务列表，数组长度最大支持20

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
}

type DescribeClassicalLBHealthStatusRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DescribeClassicalLBHealthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClassicalLBHealthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBlockIPTaskRequest struct {
	*tchttp.BaseRequest

	// 异步任务的ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryBlockIPTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBlockIPTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetItem struct {

	// 集群Id

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 集群名字

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 集群类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群标签

	SetTag *string `json:"SetTag,omitempty" name:"SetTag"`
	// 最大连接数

	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`
	// 最大入带宽

	MaxInFlow *int64 `json:"MaxInFlow,omitempty" name:"MaxInFlow"`
	// 最大入包量

	MaxInPkg *int64 `json:"MaxInPkg,omitempty" name:"MaxInPkg"`
	// 最大出带宽

	MaxOutFlow *int64 `json:"MaxOutFlow,omitempty" name:"MaxOutFlow"`
	// 最大出包量

	MaxOutPkg *int64 `json:"MaxOutPkg,omitempty" name:"MaxOutPkg"`
	// 最大新建连接数

	MaxNewConn *int64 `json:"MaxNewConn,omitempty" name:"MaxNewConn"`
	// http最大新建连接数

	HTTPMaxNewConn *int64 `json:"HTTPMaxNewConn,omitempty" name:"HTTPMaxNewConn"`
	// https最大新建连接数

	HTTPSMaxNewConn *int64 `json:"HTTPSMaxNewConn,omitempty" name:"HTTPSMaxNewConn"`
	// http QPS

	HTTPQps *int64 `json:"HTTPQps,omitempty" name:"HTTPQps"`
	// https QPS

	HTTPSQps *int64 `json:"HTTPSQps,omitempty" name:"HTTPSQps"`
	// 集群已使用vip

	Vips []*VipType `json:"Vips,omitempty" name:"Vips"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type ZoneInfo struct {

	// 可用区数值形式的唯一ID，如：100001

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区字符串形式的唯一ID，如：ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称，如：广州一区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区所在地域

	ZoneRegion *string `json:"ZoneRegion,omitempty" name:"ZoneRegion"`
	// 可用区是否是LocalZone可用区

	LocalZone *bool `json:"LocalZone,omitempty" name:"LocalZone"`
}

type ListenerHealth struct {

	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 监听器的转发规则列表

	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules"`
}

type RuleHealth struct {

	// 转发规则ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 转发规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发规则的Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 本规则上绑定的后端的健康检查状态

	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets"`
	// 后端云函数的信息（TCE暂不支持）

	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`
}

type RulesItems struct {

	// 规则id

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// uri

	Url *string `json:"Url,omitempty" name:"Url"`
	// 绑定的后端对象

	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`
}

type ZoneSlaItem struct {

	// 可用区id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可以购买的性能保障规格

	SlaSet []*string `json:"SlaSet,omitempty" name:"SlaSet"`
}

type DescribeLoadBalancerCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一个账户在一个地域的负载均衡实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertRequest struct {
	*tchttp.BaseRequest

	// 原证书ID

	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`
	// 新证书ID；注：新证书公约、备注、私钥和该参数不能同时传入

	NewCertId *string `json:"NewCertId,omitempty" name:"NewCertId"`
	// 新证书公钥；注：新证书ID和该参数不能同时传入

	NewCertContent *string `json:"NewCertContent,omitempty" name:"NewCertContent"`
	// 新证书备注；注：新证书ID和该参数不能同时传入

	NewCertAlias *string `json:"NewCertAlias,omitempty" name:"NewCertAlias"`
	// 新证书私钥；注：新证书ID和该参数不能同时传入

	NewCertKey *string `json:"NewCertKey,omitempty" name:"NewCertKey"`
	// 新国密证书签名公钥；注：新证书ID和该参数不能同时传入

	NewSignCert *string `json:"NewSignCert,omitempty" name:"NewSignCert"`
	// 新国密证书签名私钥；注：新证书ID和该参数不能同时传入

	NewSignCertPrivateKey *string `json:"NewSignCertPrivateKey,omitempty" name:"NewSignCertPrivateKey"`
}

func (r *ReplaceCertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubQuota struct {

	// 子账号Uin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 配额类型（4:公网实例配额；7:内网实例配额）

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 配额值

	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要批量修改权重的列表，例如"ModifyList": [{"ListenerId": "lbl-xxxxxxxx", "LocationId":"", "Domain": "", "Url": "", "Targets": [{"Type":"CVM", "InstanceId": "ins-xxxxxxx", "Port": "443", "Weight": 10, "EniIp":""}], "Weight": 10}]

	ModifyList []*RsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 要修改的转发规则的 ID。

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数

	Url *string `json:"Url,omitempty" name:"Url"`
	// 健康检查信息

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 会话保持时间

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 负载均衡实例与后端服务之间的转发协议，默认HTTP，可取值：HTTP、HTTPS

	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`
	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Test2 struct {

	// xxx

	X *int64 `json:"X,omitempty" name:"X"`
}

type TestBAC struct {

	// 阿打发斯蒂芬

	Abc *int64 `json:"Abc,omitempty" name:"Abc"`
}

type ZoneResource struct {

	// 主可用区，如"ap-guangzhou-1"。

	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
	// 资源列表。

	ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
	// 备可用区，如"ap-guangzhou-2"，单可用区时，备可用区为null。

	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
	// IP版本，如IPv4，IPv6，IPv6_Nat。

	IPVersion *string `json:"IPVersion,omitempty" name:"IPVersion"`
}

type ListenerBackend struct {

	// 监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 监听器下的规则信息（仅适用于HTTP/HTTPS监听器）

	Rules []*RuleTargets `json:"Rules,omitempty" name:"Rules"`
	// 监听器上绑定的后端服务列表（仅适用于TCP/UDP/TCP_SSL监听器）

	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`
	// 若支持端口段，则为端口段结束端口；若不支持端口段，则为0

	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type DeleteRewriteRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 源监听器ID

	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`
	// 目标监听器ID

	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`
	// 转发规则之间的重定向关系

	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

func (r *DeleteRewriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRewriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigAssociateListRequest struct {
	*tchttp.BaseRequest

	// 配置ID

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
	// 拉取绑定关系列表开始位置，默认值 0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 拉取绑定关系列表数目，默认值 20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// CLB维度的配置ID列表，不支持根据 Domain 过滤

	UconfigIds []*string `json:"UconfigIds,omitempty" name:"UconfigIds"`
}

func (r *DescribeCustomizedConfigAssociateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomizedConfigAssociateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetCountForLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// 负载均衡ID

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeTargetCountForLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetCountForLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务的当前状态。 0：成功，1：失败，2：进行中。

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicalLoadBalancerInfo struct {

	// 后端实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 负载均衡实例ID列表

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type ExclusiveCluster struct {

	// 4层独占集群列表

	L4Clusters []*ClusterItem `json:"L4Clusters,omitempty" name:"L4Clusters"`
	// 7层独占集群列表

	L7Clusters []*ClusterItem `json:"L7Clusters,omitempty" name:"L7Clusters"`
	// vpcgw集群

	ClassicalCluster *ClusterItem `json:"ClassicalCluster,omitempty" name:"ClassicalCluster"`
}

type SetNameItem struct {

	// 集群对外名字

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 集群对内名字

	SetChName *string `json:"SetChName,omitempty" name:"SetChName"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 由负载均衡实例唯一 ID 组成的数组。

		LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLoadBalancerCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMasterZonesRequest struct {
	*tchttp.BaseRequest

	// 园区标识字符串

	Groups []*string `json:"Groups,omitempty" name:"Groups"`
	// 查询类型，只支持TCE 侧查询，取值REAL

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeMasterZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMasterZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 显示的结果数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 显示的目标组信息集合

		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSecurityGroupForLoadbalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestEzioRequest struct {
	*tchttp.BaseRequest

	// name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *TestEzioRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestEzioRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleOutput struct {

	// 转发规则的 ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 转发规则的域名。

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发规则的路径。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 会话保持时间

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 健康检查信息

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 证书信息

	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`
	// 规则的请求转发方式

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 转发规则所属的监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的重定向目标信息

	RewriteTarget *RewriteTarget `json:"RewriteTarget,omitempty" name:"RewriteTarget"`
	// 是否开启gzip

	HttpGzip *bool `json:"HttpGzip,omitempty" name:"HttpGzip"`
	// 转发规则是否为自动创建

	BeAutoCreated *bool `json:"BeAutoCreated,omitempty" name:"BeAutoCreated"`
	// 是否作为默认域名

	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// 是否开启Http2

	Http2 *bool `json:"Http2,omitempty" name:"Http2"`
	// 负载均衡与后端服务之间的转发协议

	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`
	// 转发规则的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 后端服务器类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 绑定的目标组基本信息；当规则绑定目标组时，会返回该字段

	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`
	// WAF实例ID

	WafDomainId *string `json:"WafDomainId,omitempty" name:"WafDomainId"`
	// QUIC状态

	QuicStatus *string `json:"QuicStatus,omitempty" name:"QuicStatus"`
	// TRPC被调服务器路由，ForwardType为TRPC时有效

	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`
	// TRPC调用服务接口，ForwardType为TRPC时有效

	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

type VipType struct {

	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 负载均衡类型。1 应用型，0 传统型。

	Forward *int64 `json:"Forward,omitempty" name:"Forward"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡ID。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 绑定目标。

	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleIspRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSingleIspRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSingleIspRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubQuotaRsp struct {

	// 子账号Uin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 子账号已使用配额

	UsedQuota *int64 `json:"UsedQuota,omitempty" name:"UsedQuota"`
	// 子账号当前配额

	CurQuota *int64 `json:"CurQuota,omitempty" name:"CurQuota"`
	// 子账号修改时间

	LastSetTime *string `json:"LastSetTime,omitempty" name:"LastSetTime"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的转发规则的唯一标识数组

		LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersForVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的负载均衡实例数组。

		LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersForVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersForVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpLoadCertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书ID

		CertID *string `json:"CertID,omitempty" name:"CertID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpLoadCertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpLoadCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicalListener struct {

	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 负载均衡监听器端口

	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`
	// 监听器后端转发端口

	InstancePort *int64 `json:"InstancePort,omitempty" name:"InstancePort"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 会话保持时间

	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`
	// 是否开启了健康检查：1（开启）、0（关闭）

	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`
	// 响应超时时间

	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 检查间隔

	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`
	// 健康阈值

	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`
	// 不健康阈值

	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`
	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的请求均衡方法。wrr 表示按权重轮询，ip_hash 表示根据访问的源 IP 进行一致性哈希方式来分发

	HttpHash *string `json:"HttpHash,omitempty" name:"HttpHash"`
	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查返回码。具体可参考创建监听器中对该字段的解释

	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`
	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查路径

	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`
	// 传统型公网负载均衡的 HTTPS 监听器的认证方式

	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`
	// 传统型公网负载均衡的 HTTPS 监听器的服务端证书 ID

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 传统型公网负载均衡的 HTTPS 监听器的客户端证书 ID

	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`
	// 监听器的状态，0 表示创建中，1 表示运行中

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AssociationItem struct {

	// 关联到的负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 关联到的监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 关联到的转发规则ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 关联到的监听器协议类型，如HTTP,TCP,

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 关联到的监听器端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 关联到的转发规则域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 关联到的转发规则URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 负载均衡名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchModifyTargetWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的监听器的唯一标识数组

		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID数组

	SourceListenerIds []*string `json:"SourceListenerIds,omitempty" name:"SourceListenerIds"`
	// 负载均衡转发规则的ID数组

	SourceLocationIds []*string `json:"SourceLocationIds,omitempty" name:"SourceLocationIds"`
}

func (r *DescribeRewriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRewriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerClsLogRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 日志服务(CLS)的日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志服务(CLS)的日志主题ID

	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`
	// 日志类型，ACCESS：访问日志

	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *SetLoadBalancerClsLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerClsLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicalTarget struct {

	// 后端服务的类型，可取值：CVM、ENI（即将支持）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务的唯一 ID，可通过 DescribeInstances 接口返回字段中的 unInstanceId 字段获取

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 后端服务的外网 IP

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 后端服务的内网 IP

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 后端服务的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 后端服务的状态
	// 1：故障，2：运行中，3：创建中，4：已关机，5：已退还，6：退还中， 7：重启中，8：开机中，9：关机中，10：密码重置中，11：格式化中，12：镜像制作中，13：带宽设置中，14：重装系统中，19：升级中，21：热迁移中

	RunFlag *int64 `json:"RunFlag,omitempty" name:"RunFlag"`
}

type DescribeLBActionLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 负载均衡的操作配额详细信息

		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLBActionLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLBActionLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManualRewriteRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 源监听器ID

	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`
	// 目标监听器ID

	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`
	// 转发规则之间的重定向关系

	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

func (r *ManualRewriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManualRewriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCustomizedConfigForLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 操作类型：'ADD', 'DELETE', 'UPDATE', 'BIND', 'UNBIND'

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 除了创建个性化配置外，必传此字段，如：pz-1234abcd

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
	// 创建个性化配置或修改个性化配置的内容时，必传此字段

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
	// 创建个性化配置或修改个性化配置的名字时，必传此字段

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 绑定解绑时，必传此字段

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *SetCustomizedConfigForLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCustomizedConfigForLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLBWhiteList struct {

	// 白名单关键字

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
	// 白名单值

	WhiteListValue []*string `json:"WhiteListValue,omitempty" name:"WhiteListValue"`
}

type TargetGroupBackend struct {

	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 后端服务的类型，可取：CVM、ENI（即将支持）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务的唯一 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 后端服务的监听端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 后端服务的外网 IP

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 后端服务的内网 IP

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 后端服务的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 后端服务被绑定的时间

	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`
	// 弹性网卡唯一ID

	EniId *string `json:"EniId,omitempty" name:"EniId"`
}

type DescribeCustomizedConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置列表

		ConfigList []*ConfigListItem `json:"ConfigList,omitempty" name:"ConfigList"`
		// 配置数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomizedConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomizedConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要查询的负载均衡监听器 ID数组

	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 要查询的监听器的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoRewriteRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// HTTPS:443监听器的ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// HTTPS:443监听器下需要重定向的域名

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *AutoRewriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AutoRewriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定失败的监听器ID，如为空表示全部绑定成功。

		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicalLBHealthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 后端健康状态列表

		HealthList []*ClassicalHealth `json:"HealthList,omitempty" name:"HealthList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicalLBHealthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClassicalLBHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 新的监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 证书相关信息，此参数仅适用于HTTPS/TCP_SSL监听器

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。注意：未开启SNI的监听器可以开启SNI；已开启SNI的监听器不能关闭SNI

	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 监听器是否支持设置默认域名，默认为支持

	DefaultSerSwitch *bool `json:"DefaultSerSwitch,omitempty" name:"DefaultSerSwitch"`
	// 开启监听器默认域名开关时，指定该监听器下任意域名为默认域名，当DefaultSerSwitch传值为ture时，且该监听器有存量域名时，该参数为必填，并同时生效。

	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyListenerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfgProductCode struct {

	// 购买的产品的标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 购买的子产品的标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
}

type ZoneSet struct {

	// 是否为主可用区

	Master *bool `json:"Master,omitempty" name:"Master"`
	// 可用区数值形式的唯一ID，如：100001

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区字符串形式的唯一ID，如：ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称，如：广州一区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 安全组ID构成的数组，一个负载均衡实例最多可绑定50个安全组，如果要解绑所有安全组，可不传此参数，或传入空数组。

	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListenerItem struct {

	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 绑定规则

	Rules []*RulesItems `json:"Rules,omitempty" name:"Rules"`
	// 四层绑定对象

	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`
	// 端口段监听器的结束端口

	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type VipAndVport struct {

	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口号

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡ID。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 解绑目标。

	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书列表

		CertList []*CertList `json:"CertList,omitempty" name:"CertList"`
		// 证书总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancerHealth struct {

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 监听器列表

	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners"`
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解绑失败的监听器ID。

		FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backend struct {

	// 后端服务的类型，可取：CVM、BMS、ENI

	Type *string `json:"Type,omitempty" name:"Type"`
	// 后端服务的唯一 ID，如 ins-abcd1234

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 后端服务的监听端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 后端服务的外网 IP

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 后端服务的内网 IP

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 后端服务的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 后端服务被绑定的时间

	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`
	// 弹性网卡唯一ID，如 eni-1234abcd

	EniId *string `json:"EniId,omitempty" name:"EniId"`
}

type BatchTarget struct {

	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 子机ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 弹性网卡ip

	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
	// 绑定端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 子机权重，范围[0, 100]。绑定时如果不存在，则默认为10。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 七层规则ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 目标规则的域名，提供LocationId参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 目标规则的URL，提供LocationId参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 要修改权重的后端服务列表

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 要删除的转发规则的ID组成的数组

	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`
	// 要删除的转发规则的域名，已提供LocationIds参数时本参数不生效

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 要删除的转发规则的转发路径，已提供LocationIds参数时本参数不生效

	Url *string `json:"Url,omitempty" name:"Url"`
	// 监听器下必须配置一个默认域名，当需要删除默认域名时，可以指定另一个域名作为新的默认域名。

	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancerListByCertIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书ID，以及与该证书ID关联的负载均衡实例列表

		CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitempty" name:"CertSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancerListByCertIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancerListByCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest

	// 负载均衡的唯一ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡绑定的后端服务的地域信息

	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`
	// 网络计费相关参数

	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`
	// Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。

	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
	// 负载均衡实例的计费类型，后付费：POSTPAID_BY_HOUR，预付费：PREPAID。

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 不同计费模式之间的切换：0表示不切换，1表示预付费和后付费切换，2表示后付费之间切换。默认值：0

	SwitchFlag *uint64 `json:"SwitchFlag,omitempty" name:"SwitchFlag"`
	// 负载均衡实例的预付费相关属性

	PrepaidInfo *LBChargePrepaid `json:"PrepaidInfo,omitempty" name:"PrepaidInfo"`
	// 7层集群列表

	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
	// 是否开启跨地域绑定2.0功能（私有云不支持）

	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OneCapacity struct {

	// 集群名字

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 最大连接数

	MaxConn *uint64 `json:"MaxConn,omitempty" name:"MaxConn"`
	// 最大Qps

	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`
	// 最大入包量，单位pps

	MaxInPkg *uint64 `json:"MaxInPkg,omitempty" name:"MaxInPkg"`
	// 最大出包量，单位pps

	MaxOutPkg *uint64 `json:"MaxOutPkg,omitempty" name:"MaxOutPkg"`
	// 最大入带宽 ，单位Mbps

	MaxInFlow *uint64 `json:"MaxInFlow,omitempty" name:"MaxInFlow"`
	// 最大出带宽 ，单位Mbps

	MaxOutFlow *uint64 `json:"MaxOutFlow,omitempty" name:"MaxOutFlow"`
	// 最大新建连接速率，单位：cps

	MaxNewConn *uint64 `json:"MaxNewConn,omitempty" name:"MaxNewConn"`
	// 最大新建连接速率，单位：cps

	HttpMaxNewConn *uint64 `json:"HttpMaxNewConn,omitempty" name:"HttpMaxNewConn"`
	// 最大新建连接速率，单位：cps

	HttpsMaxNewConn *uint64 `json:"HttpsMaxNewConn,omitempty" name:"HttpsMaxNewConn"`
	// 最大新建连接速率，单位：cps

	HttpsQps *uint64 `json:"HttpsQps,omitempty" name:"HttpsQps"`
}

type RewriteTarget struct {

	// 重定向目标的监听器ID
	// 注意：此字段可能返回 null，表示无重定向。

	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`
	// 重定向目标的转发规则ID
	// 注意：此字段可能返回 null，表示无重定向。

	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`
	// 重定向状态码

	RewriteCode *int64 `json:"RewriteCode,omitempty" name:"RewriteCode"`
	// 重定向是否携带匹配的url

	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`
	// 重定向类型，Manual: 手动重定向，Auto:  自动重定向

	RewriteType *string `json:"RewriteType,omitempty" name:"RewriteType"`
}

type SetSubUinQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSubUinQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSubUinQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤器的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤器的值数组

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type LoadBalancerListenerRule struct {

	// 负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 四层七层监听器信息

	ListenerAttributes []*ListenerAttribute `json:"ListenerAttributes,omitempty" name:"ListenerAttributes"`
}

type SetLoadBalancerClsLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoadBalancerClsLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertificateInput struct {

	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证

	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`
	// 服务端RSA证书的 ID，如果不填写此项则必须上传证书，包括 CertContent，CertKey，CertName。

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 客户端CA证书的 ID，当监听器采用双向认证，即 SSLMode=MUTUAL 时，如果不填写此项则必须上传客户端证书，包括 CertCaContent，CertCaName。

	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`
	// 上传服务端RSA证书的名称，如果没有 CertId，则此项必传。

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 上传服务端RSA证书的 key，如果没有 CertId，则此项必传。

	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`
	// 上传服务端RSA证书的内容，如果没有 CertId，则此项必传。

	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`
	// 上传客户端 CA 证书的名称，如果 SSLMode=mutual，如果没有 CertCaId，则此项必传。

	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`
	// 上传客户端证书的内容，如果 SSLMode=mutual，如果没有 CertCaId，则此项必传。

	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`
	// 服务端ECC证书的 ID，如果不填写此项则必须上传证书，包括 CertContent，CertKey，CertName。

	EccCertId *string `json:"EccCertId,omitempty" name:"EccCertId"`
	// 上传服务端ECC证书的名称，如果没有 CertEccId，则此项必传。

	EccCertName *string `json:"EccCertName,omitempty" name:"EccCertName"`
	// 上传服务端ECC证书的内容，如果没有 CertEccId，则此项必传。

	EccCertContent *string `json:"EccCertContent,omitempty" name:"EccCertContent"`
	// 上传服务端ECC证书的key，如果没有 CertEccId，则此项必传。

	EccCertKey *string `json:"EccCertKey,omitempty" name:"EccCertKey"`
	// 服务端国密SM证书的 ID，如果不填写此项则必须上传证书，包括 CertSMContent，CertSmKey，CertSMName，CertSMSignContent，CertSMSignKey。

	SmCertId *string `json:"SmCertId,omitempty" name:"SmCertId"`
	// 上传服务端国密SM证书的名称，如果没有 CertSMId，则此项必传。

	SmCertName *string `json:"SmCertName,omitempty" name:"SmCertName"`
	// 上传服务端国密SM证书的签名公约，如果没有 CertSMId，则此项必传。

	SmCertSignContent *string `json:"SmCertSignContent,omitempty" name:"SmCertSignContent"`
	// 上传服务端国密SM证书的签名密钥，如果没有 CertSMId，则此项必传。

	SmCertSignKey *string `json:"SmCertSignKey,omitempty" name:"SmCertSignKey"`
	// 上传服务端国密SM证书的证书公约，如果没有 CertSMId，则此项必传。

	SmCertContent *string `json:"SmCertContent,omitempty" name:"SmCertContent"`
	// 上传服务端国密SM证书的证书密钥，如果没有 CertSMId，则此项必传。

	SmCertKey *string `json:"SmCertKey,omitempty" name:"SmCertKey"`
}

type RuleTargets struct {

	// 转发规则的 ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 转发规则的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发规则的路径。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 后端服务的信息

	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`
	// 后端云函数的信息（TCE暂不支持）

	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`
}

type ModifyCertAliasRequest struct {
	*tchttp.BaseRequest

	// 证书ID

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 证书新备注

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyCertAliasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertAliasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCustomizedConfigForLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个性化配置ID，如：pz-1234abcd

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCustomizedConfigForLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCustomizedConfigForLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerLogRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 操作类型，ADD | DEL | UPDATE

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 如果用户开启日志访问，则此字段指定日志存放位置（COS的Bucket），OperationType 为ADD时，必须穿入此参数。

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 日志导出到bucket的频率，单位：秒，有效取值范围为10分钟到12小时，建议取3600秒

	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *SetLoadBalancerLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDetailItem struct {

	// 配置绑定的CLB ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 配置绑定的监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 配置绑定的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 配置绑定的规则

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 监听器名字

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// location的url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 配置ID

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
}

type ClassicalHealth struct {

	// 后端服务的内网 IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 后端服务的端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 负载均衡的监听端口

	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`
	// 转发协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 健康检查结果，1 表示健康，0 表示不健康

	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type Cluster struct {

	// 集群唯一ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型，如TGW，STGW，VPCGW

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群标签，只有STGW集群有标签

	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`
	// 集群所在可用区，如ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 集群网络类型，如Public，Private

	Network *string `json:"Network,omitempty" name:"Network"`
	// 最大连接数

	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`
	// 最大入带宽

	MaxInFlow *int64 `json:"MaxInFlow,omitempty" name:"MaxInFlow"`
	// 最大入包量

	MaxInPkg *int64 `json:"MaxInPkg,omitempty" name:"MaxInPkg"`
	// 最大出带宽

	MaxOutFlow *int64 `json:"MaxOutFlow,omitempty" name:"MaxOutFlow"`
	// 最大出包量

	MaxOutPkg *int64 `json:"MaxOutPkg,omitempty" name:"MaxOutPkg"`
	// 最大新建连接数

	MaxNewConn *int64 `json:"MaxNewConn,omitempty" name:"MaxNewConn"`
	// http最大新建连接数

	HTTPMaxNewConn *int64 `json:"HTTPMaxNewConn,omitempty" name:"HTTPMaxNewConn"`
	// https最大新建连接数

	HTTPSMaxNewConn *int64 `json:"HTTPSMaxNewConn,omitempty" name:"HTTPSMaxNewConn"`
	// http QPS

	HTTPQps *int64 `json:"HTTPQps,omitempty" name:"HTTPQps"`
	// https QPS

	HTTPSQps *int64 `json:"HTTPSQps,omitempty" name:"HTTPSQps"`
	// 集群内资源总数目

	ResourceCount *int64 `json:"ResourceCount,omitempty" name:"ResourceCount"`
	// 集群内空闲资源数目

	IdleResourceCount *int64 `json:"IdleResourceCount,omitempty" name:"IdleResourceCount"`
	// 集群内转发机的数目

	LoadBalanceDirectorCount *int64 `json:"LoadBalanceDirectorCount,omitempty" name:"LoadBalanceDirectorCount"`
	// 集群的Isp属性，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。

	Isp *string `json:"Isp,omitempty" name:"Isp"`
}

type IspSet struct {

	// 运营商ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 运营商类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运营商名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运营商IPv4状态

	IspStatus *uint64 `json:"IspStatus,omitempty" name:"IspStatus"`
	// 运营商IPv6状态

	IspStatusIPv6 *uint64 `json:"IspStatusIPv6,omitempty" name:"IspStatusIPv6"`
	// 运营商v4别名

	V4alias *string `json:"V4alias,omitempty" name:"V4alias"`
	// 运营商v6别名

	V6alias *string `json:"V6alias,omitempty" name:"V6alias"`
}

type AssociateCustomizedConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID

	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
	// 关联的server或location

	BindList []*BindItem `json:"BindList,omitempty" name:"BindList"`
}

func (r *AssociateCustomizedConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateCustomizedConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetInnerNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群实际名字

		SetInnerNames []*SetNameItem `json:"SetInnerNames,omitempty" name:"SetInnerNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetInnerNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetInnerNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {

	// 配额名称

	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`
	// 使用数量

	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`
	// 配额数量

	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type DeleteCertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLBListenersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定的后端规则

		LoadBalancers []*LBItem `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLBListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLBListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteListSupportRequest struct {
	*tchttp.BaseRequest

	// 查询的白名单关键字

	WhiteListKeys []*string `json:"WhiteListKeys,omitempty" name:"WhiteListKeys"`
}

func (r *DescribeWhiteListSupportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteListSupportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 要删除的负载均衡实例 ID数组，数组大小最大支持20

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoadBalancerSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoadBalancerSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Listener struct {

	// 负载均衡监听器 ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 监听器绑定的证书信息

	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`
	// 监听器的健康检查信息

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 请求的调度方式

	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
	// 会话保持时间

	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`
	// 是否开启SNI特性（本参数仅对于HTTPS监听器有意义）

	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`
	// 监听器下的全部转发规则（本参数仅对于HTTP/HTTPS监听器有意义）

	Rules []*RuleOutput `json:"Rules,omitempty" name:"Rules"`
	// 监听器的名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 端口段结束端口

	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
	// 后端服务器类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 绑定的目标组基本信息；当监听器绑定目标组时，会返回该字段

	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`
	// 监听器是否开启设置默认域名标识，0为未开启；1为开启

	DefaultSerSwitch *uint64 `json:"DefaultSerSwitch,omitempty" name:"DefaultSerSwitch"`
	// 解绑后端目标时，是否发RST给客户端，（此参数仅对于TCP监听器有意义）。

	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`
	// 仅支持Nat64 CLB TCP监听器

	Toa *bool `json:"Toa,omitempty" name:"Toa"`
	// 会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。

	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`
	// 是否开启长连接，1开启，0关闭，（本参数仅对于HTTP/HTTPS监听器有意义）

	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`
}

type DescribeBlockIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的IP的数量

		BlockedIPCount *uint64 `json:"BlockedIPCount,omitempty" name:"BlockedIPCount"`
		// 获取用户真实IP的字段

		ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`
		// 加入了12360黑名单的IP列表

		BlockedIPList []*BlockedIP `json:"BlockedIPList,omitempty" name:"BlockedIPList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlockIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindItem struct {

	// 配置绑定的CLB ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 配置绑定的监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 配置绑定的域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 配置绑定的规则

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type RegionItem struct {

	// 地域，如"ap-guangzhou"。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域ID，如1。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域缩写，如"gz"。

	RegionPrefix *string `json:"RegionPrefix,omitempty" name:"RegionPrefix"`
}

type Price struct {

	// 描述了实例价格。

	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
	// 描述了网络价格。

	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type DescribeAppIdLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户和标签对应关系

		OwnerLabelSet []*OwnerLabel `json:"OwnerLabelSet,omitempty" name:"OwnerLabelSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppIdLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest

	// 目标组ID数组

	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`
	// 过滤条件数组，支持TargetGroupVpcId和TargetGroupName。与TargetGroupIds互斥，优先使用目标组ID，

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 显示的偏移起始量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 显示条数限制，默认为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTargetGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResourcesRequest struct {
	*tchttp.BaseRequest

	// 返回集群中资源列表数目，默认20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回集群中资源列表起始偏移量，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询集群中资源列表条件，详细的过滤条件如下：
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照vip过滤。</li>
	// <li> loadblancer-id - String - 是否必填：否 - （过滤条件）按照负载均衡唯一ID过滤。</li>
	// <li> idle - String 是否必填：否 - （过滤条件）按照是否闲置过滤，如"True","False"。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器 ID列表

	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
	// 监听器协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type NetworkPriceInfo struct {

	// 网络收费单位价格，

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 网络收费单位，GB按流量计费，HOUR按小时计算

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

type QuotaData struct {

	// 配额类型（4:公网实例配额；7:内网实例配额）

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 已使用配额

	UsedQuota *int64 `json:"UsedQuota,omitempty" name:"UsedQuota"`
	// 主账号配额

	CurQuota *int64 `json:"CurQuota,omitempty" name:"CurQuota"`
	// 该类型的子账号配额列表

	SubQuota []*SubQuotaRsp `json:"SubQuota,omitempty" name:"SubQuota"`
	// 子账户使用配额总和

	SubUsedTotalQuota *int64 `json:"SubUsedTotalQuota,omitempty" name:"SubUsedTotalQuota"`
}

type QueryBlockIPTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 running，2 fail，6 succ

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBlockIPTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBlockIPTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LBItem struct {

	// lb的字符串id

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// lb的vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 监听器规则

	Listeners []*ListenerItem `json:"Listeners,omitempty" name:"Listeners"`
	// LB所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ListenerAttribute struct {

	// 监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL（TCP_SSL 正在内测中，如需使用请通过工单申请）

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。取值"ON"或者"OFF"

	SniSwitch *string `json:"SniSwitch,omitempty" name:"SniSwitch"`
	// 证书相关信息，此参数仅适用于TCP_SSL监听器或者不开启SNI特性的HTTPS监听器

	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
	// 七层监听器的规则信息

	LocationRules []*LocationRule `json:"LocationRules,omitempty" name:"LocationRules"`
	// 四层规则的信息

	ListenerRule *ListenerRule `json:"ListenerRule,omitempty" name:"ListenerRule"`
}

type TargetCountForLoadBalancer struct {

	// 负载均衡ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例绑定的后端服务总数

	TargetCount *uint64 `json:"TargetCount,omitempty" name:"TargetCount"`
}

type TargetGroupInfo struct {

	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 目标组的vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 目标组的名字

	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`
	// 目标组的默认端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 目标组的创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 目标组的修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 关联到的规则数组

	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitempty" name:"AssociatedRule"`
}

type Test struct {

	// A

	A *string `json:"A,omitempty" name:"A"`
}

type ModifyDnatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDnatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDnatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomizedConfigAssociateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定关系列表

		BindList []*BindDetailItem `json:"BindList,omitempty" name:"BindList"`
		// 绑定关系总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomizedConfigAssociateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomizedConfigAssociateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 显示的结果数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 显示的目标组信息集合

		TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 负载均衡实例列表

		LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocationRule struct {

	// 新建转发规则的信息

	Rule *RuleInput `json:"Rule,omitempty" name:"Rule"`
	// 待绑定的后端服务列表，数组长度最大支持20

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
}

type DescribeBlockIPListRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID。

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 数据偏移量，默认为 0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回IP的最大个数，默认为 100000。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBlockIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBlockIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIspInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIspInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIspInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertAliasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertAliasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCertAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertIdRelatedWithLoadBalancers struct {

	// 证书ID

	CertId *string `json:"CertId,omitempty" name:"CertId"`
	// 与证书关联的负载均衡实例列表

	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
}

type ClusterItem struct {

	// 集群唯一ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群所在可用区，如ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type OwnerLabel struct {

	// 用户AppId

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// 集群标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 集群标签类型，取值为L4_LAN_CLB，L4_WAN_CLB，L4_CLB，L7_LAN_CLB，L7_WAN_CLB，L7_CLB

	SetType *string `json:"SetType,omitempty" name:"SetType"`
}

type TargetRegionInfo struct {

	// Target所属地域，如 ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// Target所属网络，私有网络格式如 vpc-abcd1234，如果是基础网络，则为"0"

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicTargetGroupInfo struct {

	// 目标组ID

	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 目标组名称

	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`
}

type DescribeClsLogSetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsLogSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
		// 健康检查日志集ID
		HealthLogsetId *string `json:"HealthLogsetId,omitempty" name:"HealthLogsetId"`
		// 唯一请求 ID，每次请求都会返回。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClsLogSetRequest struct {
	*tchttp.BaseRequest

	// 日志集名称
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志集保存时间（天）
	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 日志集类型，ACCESS：访问日志，HEALTH：健康检查日志
	LogsetType *string `json:"LogsetType,omitempty" name:"LogsetType"`
}

func (r *CreateClsLogSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsLogSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClsLogSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
		// 唯一请求 ID，每次请求都会返回。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClsLogSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题的名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志主题的分区个数
	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
}

func (r *CreateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 唯一请求 ID，每次请求都会返回。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

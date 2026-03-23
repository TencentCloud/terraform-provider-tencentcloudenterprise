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

package v20190924

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeSourceCodeAuthUserInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSourceCodeAuthUserInfoPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeAuthUserInfoPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardRequestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代理请求响应body

		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`
		// 代理请求响应的header

		ResponseHeaders []*ResponseHeader `json:"ResponseHeaders,omitempty" name:"ResponseHeaders"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForwardRequestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ForwardRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageImageLifecycleGlobalPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 用户在registry中的用户名

	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`
	// 镜像名称，不包含tag

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag的格式

	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`
	// 源码所在的git服务

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 源码在git服务器上的仓库名

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 用户在git服务器上的用户名

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// Trigger

	Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`
	// dockerfile在仓库中的路径

	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`
	// 工作目录

	WorkDir *string `json:"WorkDir,omitempty" name:"WorkDir"`
	// 构建出的镜像覆盖该Tag

	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`
	// Args

	Args []*string `json:"Args,omitempty" name:"Args"`
}

func (r *ModifyImageBuildPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type TcrRepositoryInfo struct {
	// 仓库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 创建时间，格式"2006-01-02 15:04:05.999999999 -0700 MST"
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 是否公开
	Public *bool `json:"Public,omitnil,omitempty" name:"Public"`

	// 仓库详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 简单描述
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// 更新时间，格式"2006-01-02 15:04:05.999999999 -0700 MST"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type BatchDeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DuplicateImagePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 复制镜像返回值

		Data *DupImageTagResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DuplicateImagePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DuplicateImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerInvokeCondition struct {

	// 触发方式

	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`
	// 触发表达式

	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

type CreateApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationTriggerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageLifecyclePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryFilterPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库信息

		Data *SearchUserRepositoryResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryFilterPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryFilterPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageInternalEndpointRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 需要接入的用户vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 需要接入的用户子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 请求的地域名称，用于实例复制地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type ManageInternalEndpointRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 需要接入的用户vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 需要接入的用户子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 请求的地域名称，用于实例复制地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}
func (r *ManageInternalEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageInternalEndpointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest

	// 用户凭证

	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`
}

func (r *ValidateApplicationTokenPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateApplicationTokenPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageLifecyclePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名字

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerLogResp struct {

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Tag名称

	TagName *string `json:"TagName,omitempty" name:"TagName"`
	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
	// 触发方式

	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`
	// 触发动作

	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`
	// 触发时间

	InvokeTime *string `json:"InvokeTime,omitempty" name:"InvokeTime"`
	// 触发条件

	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`
	// 触发参数

	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`
	// 触发结果

	InvokeResult *TriggerInvokeResult `json:"InvokeResult,omitempty" name:"InvokeResult"`
}

type DeleteFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavorRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个人收藏仓库列表返回信息

		Data *FavorResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFavorRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFavorRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryAccessPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRepositoryAccessPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepositoryAccessPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTokenPersonalResp struct {

	// 镜像仓库凭证

	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`
}

type DesGitAuthRsp struct {

	// 用户GitAuth信息

	AuthMap *string `json:"AuthMap,omitempty" name:"AuthMap"`
}

type RegistryStatus struct {

	// 实例的Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 实例的状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 附加状态

	Conditions []*RegistryCondition `json:"Conditions,omitempty" name:"Conditions"`
}

type DockerHubRepoinfo struct {

	// 仓库名称

	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`
	// 仓库类型

	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`
	// 仓库Logo

	Logo *string `json:"Logo,omitempty" name:"Logo"`
	// 简述

	SimpleDesc *string `json:"SimpleDesc,omitempty" name:"SimpleDesc"`
	// 详述

	DetailDesc *string `json:"DetailDesc,omitempty" name:"DetailDesc"`
	// 收藏次数

	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
	// 是否用户的收藏

	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`
}

type DescribeApplicationTriggerLogPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 触发日志返回值

		Data *DescribeApplicationTriggerLogPersonalResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTriggerLogPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTriggerLogPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeerReplicationOption struct {
	// 待同步实例的uin
	PeerRegistryUin *string `json:"PeerRegistryUin,omitnil,omitempty" name:"PeerRegistryUin"`

	// 待同步实例的访问永久Token
	PeerRegistryToken *string `json:"PeerRegistryToken,omitnil,omitempty" name:"PeerRegistryToken"`

	// 是否开启跨主账号实例同步
	EnablePeerReplication *bool `json:"EnablePeerReplication,omitnil,omitempty" name:"EnablePeerReplication"`
}

type ManageReplicationRequest struct {
	*tchttp.BaseRequest

	// 复制源实例ID
	SourceRegistryId *string `json:"SourceRegistryId,omitnil,omitempty" name:"SourceRegistryId"`

	// 复制目标实例ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitnil,omitempty" name:"DestinationRegistryId"`

	// 同步规则
	Rule *ReplicationRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 目标实例的地域ID，如广州是1
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitnil,omitempty" name:"DestinationRegionId"`

	// 开启跨主账号实例同步配置项
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitnil,omitempty" name:"PeerReplicationOption"`
}
func (r *ManageReplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageReplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageBuildPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthUserInfoResp struct {

	// 构建用户信息

	User *AuthUser `json:"User,omitempty" name:"User"`
}

type ResponseHeader struct {

	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UserIsExistsResp struct {

	// 用户是否存在

	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
	// 主账号是否存在

	MainIsExist *bool `json:"MainIsExist,omitempty" name:"MainIsExist"`
}

type SecurityPolicy struct {

	// 策略索引

	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 192.168.1.0/24

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 安全策略的版本

	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 实例的规格

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RespLimit struct {

	// 配额信息

	LimitInfo []*Limit `json:"LimitInfo,omitempty" name:"LimitInfo"`
}

type CreateImageBuildTaskDockerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 构建Id

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildTaskDockerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildTaskDockerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Limit用于分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量用于分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDockerHubRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BuildRule信息返回值

		Data *BuildRuleResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像tag信息

		Data *TagInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalEndpointsRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeInternalEndpointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalEndpointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 用户在registry中的用户名

	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`
	// 镜像名称，不包含tag

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag的格式

	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`
	// 源码所在的git服务

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 源码在git服务器上的仓库名

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 用户在git服务器上的用户名

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// Trigger信息

	Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`
	// dockerfile在仓库中的路径

	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`
	// 工作目录

	WorkDir *string `json:"WorkDir,omitempty" name:"WorkDir"`
	// 构建出的镜像覆盖该Tag

	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`
	// Args

	Args []*string `json:"Args,omitempty" name:"Args"`
}

func (r *CreateImageBuildPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildBranchResp struct {

	// 构建的分支信息

	Branches []*string `json:"Branches,omitempty" name:"Branches"`
}

type BuildInfoResp struct {

	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 构建信息列表

	BuildList []*BuildInfo `json:"BuildList,omitempty" name:"BuildList"`
}

type BatchDeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称数组

	RepoNames []*string `json:"RepoNames,omitempty" name:"RepoNames"`
}

func (r *BatchDeleteRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryOwnerPersonalRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20, 最大值 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryOwnerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryOwnerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageExternalEndpointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageExternalEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageExternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserPasswordPersonalRequest struct {
	*tchttp.BaseRequest

	// 更新后的密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyUserPasswordPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserPasswordPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildRule struct {

	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 用户在registry中的用户名

	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`
	// 镜像名称，不包含tag

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag的格式

	ImageTagFormat *string `json:"ImageTagFormat,omitempty" name:"ImageTagFormat"`
	// 源码所在的git服务

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 源码在git服务器上的仓库名

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 用户在git服务器上的用户名

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 分支

	Branches []*string `json:"Branches,omitempty" name:"Branches"`
	// Tag

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
	// dockerfile在仓库中的路径

	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`
	// 工作目录

	BuildWorkDir *string `json:"BuildWorkDir,omitempty" name:"BuildWorkDir"`
	// 构建出的镜像覆盖该Tag

	ForceTag *string `json:"ForceTag,omitempty" name:"ForceTag"`
	// Args

	BuildArgs *string `json:"BuildArgs,omitempty" name:"BuildArgs"`
}

type DescribeApplicationTriggerLogPersonalResp struct {

	// 返回总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 触发日志列表

	LogInfo []*TriggerLogResp `json:"LogInfo,omitempty" name:"LogInfo"`
}

type RepoInfo struct {

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
	// Tag数量

	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`
	// 是否为公开

	Public *int64 `json:"Public,omitempty" name:"Public"`
	// 是否为用户收藏

	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`
	// 是否为腾讯云官方仓库

	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`
	// 被收藏的个数

	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
	// 拉取的数量

	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 仓库创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 仓库更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteImagePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例的状态列表

		RegistryStatusSet []*RegistryStatus `json:"RegistryStatusSet,omitempty" name:"RegistryStatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryAccessPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 默认值为0

	Public *int64 `json:"Public,omitempty" name:"Public"`
}

func (r *ModifyRepositoryAccessPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepositoryAccessPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildRuleResp struct {

	// BuildRule信息

	BuildRule *BuildRule `json:"BuildRule,omitempty" name:"BuildRule"`
}

type RespDockerHubRepoList struct {

	// 仓库总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 仓库信息列表

	RepoInfo []*HubSimpleInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`
}

type DeleteImageBuildTaskPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildTaskPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildTaskPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryBranchPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 构建仓库分支信息

		Data *BuildBranchResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeRepositoryBranchPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeRepositoryBranchPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FavorResp struct {

	// 收藏仓库的总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 仓库信息数组

	RepoInfo []*Favors `json:"RepoInfo,omitempty" name:"RepoInfo"`
}

type ManageReplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageReplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryFilterPersonalRequest struct {
	*tchttp.BaseRequest

	// 搜索镜像名

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20，最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件：1表示public，0表示private

	Public *int64 `json:"Public,omitempty" name:"Public"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeRepositoryFilterPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryFilterPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryOwnerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库信息

		Data *RepoInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryOwnerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryOwnerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateApplicationTokenPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationTokenPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryAllPersonalRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20, 最大值 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 筛选条件，支持pullCount和official两个值

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序还是降序，默认是desc，还可以选择asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRepositoryAllPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryAllPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcrImageInfo struct {
	// 哈希值
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// 镜像体积（单位：字节）
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 制品类型
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// KMS 签名信息
	KmsSignature *string `json:"KmsSignature,omitnil,omitempty" name:"KmsSignature"`
}

type BuildInfo struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 构建类型

	BuildType *string `json:"BuildType,omitempty" name:"BuildType"`
	// 是否手动构建

	BuildManually *int64 `json:"BuildManually,omitempty" name:"BuildManually"`
	// 构建工作目录

	BuildWorkDir *string `json:"BuildWorkDir,omitempty" name:"BuildWorkDir"`
	// 构建参数

	Args *string `json:"Args,omitempty" name:"Args"`
	// 构建状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 构建开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 构建结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 构建仓库地址

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 构建所在的repo

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// Repo url地址

	RepoUrl *string `json:"RepoUrl,omitempty" name:"RepoUrl"`
	// 用户在git服务器上的用户名

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 构建的分支

	Branch *string `json:"Branch,omitempty" name:"Branch"`
	// dockerfile在仓库中的路径

	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`
	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 用户在registry中的用户名

	RegistryUsername *string `json:"RegistryUsername,omitempty" name:"RegistryUsername"`
	// 镜像名称，不包含tag

	Image *string `json:"Image,omitempty" name:"Image"`
	// 镜像名

	ForceImage *string `json:"ForceImage,omitempty" name:"ForceImage"`
	// 提交的SHA

	CommitSHA *string `json:"CommitSHA,omitempty" name:"CommitSHA"`
	// 提交的作者

	CommitAuthor *string `json:"CommitAuthor,omitempty" name:"CommitAuthor"`
	// 提交的信息

	CommitMessage *string `json:"CommitMessage,omitempty" name:"CommitMessage"`
	// 提交的时间

	CommitTime *string `json:"CommitTime,omitempty" name:"CommitTime"`
	// 构建日志

	BuildLog *string `json:"BuildLog,omitempty" name:"BuildLog"`
}

type DeleteImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Tag名

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DeleteImagePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRepositoryInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyRepositoryInfoPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepositoryInfoPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateUserPersonalRequest struct {
	*tchttp.BaseRequest

	// 自定义用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ValidateUserPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateUserPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoDelStrategyInfoResp struct {

	// 总数目

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 自动删除策略列表

	StrategyInfo []*AutoDelStrategyInfo `json:"StrategyInfo,omitempty" name:"StrategyInfo"`
}

type BuildTaskInfoResponse struct {

	// 构建ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 构建状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 构建开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 构建结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 构建日志

	BuildLog *string `json:"BuildLog,omitempty" name:"BuildLog"`
}

type DescribeApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 触发器列表返回值

		Data *DescribeApplicationTriggerPersonalResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTriggerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest

	// global_keep_last_days:全局保留最近几天的数据;global_keep_last_nums:全局保留最近多少个

	Type *string `json:"Type,omitempty" name:"Type"`
	// 策略值

	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *ManageImageLifecycleGlobalPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 白名单Id

	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
	// 白名单版本

	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageFilterPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// payload

		Data *SameImagesResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageFilterPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageFilterPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthUser struct {

	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
}

type CreateRepoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRepoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceRequestParams struct {
	// 实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 是否删除存储桶，默认为false
	DeleteBucket *bool `json:"DeleteBucket,omitnil,omitempty" name:"DeleteBucket"`

	// 是否dryRun模式，缺省值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 是否删除存储桶，默认为false
	DeleteBucket *bool `json:"DeleteBucket,omitnil,omitempty" name:"DeleteBucket"`

	// 是否dryRun模式，缺省值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogPersonalRequest struct {
	*tchttp.BaseRequest

	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// registry在仓库中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// Offset用于分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit用于分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageBuildTaskLogPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildTaskLogPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageLifecycleGlobalPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageInternalEndpointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageInternalEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Limit struct {

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 配额的类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置的值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type TcrRepositoryInfoResp struct {

	// 仓库信息列表

	RepositoryList []*TcrRepositoryInfo `json:"RepositoryList,omitempty" name:"RepositoryList"`
}

type RepositoryInfoResp struct {

	// 镜像仓库名字

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 镜像仓库类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
	// 镜像仓库服务地址

	Server *string `json:"Server,omitempty" name:"Server"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 镜像仓库描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否为公有镜像

	Public *int64 `json:"Public,omitempty" name:"Public"`
	// 下载次数

	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`
	// 收藏次数

	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
	// 是否为用户收藏

	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`
	// 是否为腾讯云官方镜像

	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`
}

type BatchDeleteFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 收藏仓库的列表

	Favors []*RequestFavor `json:"Favors,omitempty" name:"Favors"`
}

func (r *BatchDeleteFavorRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteFavorRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
}

func (r *DeleteApplicationTriggerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageConfigPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Tag名称

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImageConfigPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageConfigPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例安全策略组

		SecurityPolicySet []*SecurityPolicy `json:"SecurityPolicySet,omitempty" name:"SecurityPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器关联的镜像仓库，library/test格式

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发

	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`
	// 触发方式对应的表达式

	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
	// 应用所在TKE集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 应用所在TKE集群命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 应用所在TKE集群工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 应用所在TKE集群工作负载下容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 应用所在TKE集群地域数字ID，如1（广州）、16（成都）

	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
	// 新触发器名称

	NewTriggerName *string `json:"NewTriggerName,omitempty" name:"NewTriggerName"`
}

func (r *ModifyApplicationTriggerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 云标签的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 云标签的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagSpecification struct {
	// 默认值为instance
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 云标签数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RegistryChargePrepaid struct {
	// 购买实例的时长，单位：月
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识，0：手动续费，1：自动续费，2：不续费并且不通知
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 企业版实例名称
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// 企业版实例类型（basic 基础版；standard 标准版；premium 高级版）
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例计费类型，0表示按量计费，1表示预付费，默认为按量计费
	RegistryChargeType *int64 `json:"RegistryChargeType,omitnil,omitempty" name:"RegistryChargeType"`

	// 预付费自动续费标识和购买时长
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`

	// 是否开启Cos桶多AZ特性
	EnableCosMAZ *bool `json:"EnableCosMAZ,omitnil,omitempty" name:"EnableCosMAZ"`

	// 是否开启实例删除保护
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}
type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 企业版实例名称
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// 企业版实例类型（basic 基础版；standard 标准版；premium 高级版）
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例计费类型，0表示按量计费，1表示预付费，默认为按量计费
	RegistryChargeType *int64 `json:"RegistryChargeType,omitnil,omitempty" name:"RegistryChargeType"`

	// 预付费自动续费标识和购买时长
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`

	// 是否开启Cos桶多AZ特性
	EnableCosMAZ *bool `json:"EnableCosMAZ,omitnil,omitempty" name:"EnableCosMAZ"`

	// 是否开启实例删除保护
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20, 最大值 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// tag名称，可根据输入搜索

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImagePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间，支持模糊查询

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 单页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNamespacePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageExternalEndpointRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 操作（Create/Delete）

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageExternalEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageExternalEndpointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespaceInfo struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 命名空间下仓库数量

	RepoCount *int64 `json:"RepoCount,omitempty" name:"RepoCount"`
}

type BatchDeleteImagePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteImagePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeImageLifecyclePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageLifecyclePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFavorRepositoryRegionPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavorRepositoryRegionPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorRepositoryRegionPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额返回信息

		Data *RespLimit `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserQuotaPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserQuotaPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateGitHubAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// Code

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *ValidateGitHubAuthPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateGitHubAuthPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DupImageTagResp struct {

	// 镜像Digest值

	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

type RepoInfoResp struct {

	// 仓库总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 仓库信息列表

	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`
	// Server信息

	Server *string `json:"Server,omitempty" name:"Server"`
}

type DescribeImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动删除策略信息

		Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageLifecyclePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNamespacePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespacePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcrInstanceToken struct {

	// 令牌ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 令牌描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 令牌所属实例ID

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 令牌启用状态

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 令牌创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 令牌过期时间戳

	ExpiredAt *int64 `json:"ExpiredAt,omitempty" name:"ExpiredAt"`
}

type BuildReposList struct {

	// 仓库信息列表

	Repos []*BuildRepo `json:"Repos,omitempty" name:"Repos"`
}

type DeleteFavorRepositoryRegionPersonalRequest struct {
	*tchttp.BaseRequest

	// 被收藏的仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteFavorRepositoryRegionPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorRepositoryRegionPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表(为空时，
	// 表示获取账号下所有实例)

	Registryids []*string `json:"Registryids,omitempty" name:"Registryids"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 获取所有地域的实例，默认为False

	AllRegion *bool `json:"AllRegion,omitempty" name:"AllRegion"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域信息列表

		Regions []*Region `json:"Regions,omitempty" name:"Regions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateNamespaceExistPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间是否存在

		Data *NamespaceIsExistsResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateNamespaceExistPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateNamespaceExistPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建第三方应用访问凭证执行结果

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationTokenPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationTokenPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteImageLifecyclePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageLifecyclePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcrNamespaceInfo struct {
	// 命名空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 访问级别
	Public *bool `json:"Public,omitnil,omitempty" name:"Public"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 实例云标签
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 命名空间元数据
	Metadata []*KeyValueString `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitnil,omitempty" name:"CVEWhitelistItems"`

	// 扫描级别，true为自动，false为手动
	AutoScan *bool `json:"AutoScan,omitnil,omitempty" name:"AutoScan"`

	// 安全阻断级别，true为开启，false为关闭
	PreventVUL *bool `json:"PreventVUL,omitnil,omitempty" name:"PreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high, 为""时表示没有设置
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`
}


type NamespaceIsExistsResp struct {

	// 命名空间是否存在

	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
	// 是否为保留命名空间

	IsPreserved *bool `json:"IsPreserved,omitempty" name:"IsPreserved"`
}

type DeleteImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteImageLifecycleGlobalPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取镜像仓库个人版凭证

		Data *DescribeApplicationTokenPersonalResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationTokenPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTokenPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HubSimpleInfo struct {

	// 仓库名称

	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`
	// 仓库类型

	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`
	// 仓库Logo

	Logo *string `json:"Logo,omitempty" name:"Logo"`
	// 仓库简述

	SimpleDesc *string `json:"SimpleDesc,omitempty" name:"SimpleDesc"`
	// 是否为收藏

	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`
	// 收藏数量

	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
}

type CreateImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// keep_last_days:保留最近几天的数据;keep_last_nums:保留最近多少个

	Type *string `json:"Type,omitempty" name:"Type"`
	// 策略值

	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *CreateImageLifecyclePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageLifecyclePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 是否公共,1:公共,0:私有

	Public *uint64 `json:"Public,omitempty" name:"Public"`
	// 仓库描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeDockerHubImagePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubImagePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityPoliciesRequest struct {
	*tchttp.BaseRequest

	// 实例的Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeSecurityPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerInvokeResult struct {

	// 请求TKE返回值

	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
	// 请求TKE返回信息

	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
}

type CreateFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库的名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库的类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *CreateFavorRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFavorRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskDockerPersonalRequest struct {
	*tchttp.BaseRequest

	// dockerfile在仓库中的路径

	Dockerfile *string `json:"Dockerfile,omitempty" name:"Dockerfile"`
	// 用户在registry中的用户名

	RegistryUserName *string `json:"RegistryUserName,omitempty" name:"RegistryUserName"`
	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 镜像名称

	Image *string `json:"Image,omitempty" name:"Image"`
}

func (r *CreateImageBuildTaskDockerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildTaskDockerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchUserRepositoryResp struct {

	// 总个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 仓库列表

	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`
	// Server

	Server *string `json:"Server,omitempty" name:"Server"`
	// PrivilegeFiltered

	PrivilegeFiltered *bool `json:"PrivilegeFiltered,omitempty" name:"PrivilegeFiltered"`
}

type Favors struct {

	// 仓库名字

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
	// Pull总共的次数

	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`
	// 仓库收藏次数

	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`
	// 仓库是否公开

	Public *int64 `json:"Public,omitempty" name:"Public"`
	// 是否为官方所有

	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`
	// 仓库Tag的数量

	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`
	// Logo

	Logo *string `json:"Logo,omitempty" name:"Logo"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域的Id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type NamespaceInfoResp struct {

	// 命名空间数量

	NamespaceCount *int64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`
	// 命名空间信息

	NamespaceInfo []*NamespaceInfo `json:"NamespaceInfo,omitempty" name:"NamespaceInfo"`
}

type TriggerResp struct {

	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
	// 触发来源

	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`
	// 触发动作

	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 触发条件

	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`
	// 触发器参数

	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`
}

type DescribeImageConfigPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// payload

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageConfigPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageConfigPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInfoListResp struct {

	// 镜像信息列表

	ImageInfoList []*TcrImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`
}

type RegistryCondition struct {

	// 实例创建过程类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 实例创建过程状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 转换到该过程的简明原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type CreateSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSourceCodeAuthPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalEndpointStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeExternalEndpointStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalEndpointStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// registry中的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// Image名称

	Image *string `json:"Image,omitempty" name:"Image"`
}

func (r *DescribeImageBuildPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationFilter struct {

	// 类型（name、tag和resource）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 默认为空

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeImageBuildTaskLogInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 构建日志ID

	BuildId *int64 `json:"BuildId,omitempty" name:"BuildId"`
}

func (r *DescribeImageBuildTaskLogInfoPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildTaskLogInfoPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的用户信息

		Data *UserInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalEndpointStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开启公网访问状态，包括开启中，开启成功以及关闭和更新失败等

		Status *string `json:"Status,omitempty" name:"Status"`
		// 原因

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExternalEndpointStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserPasswordPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserPasswordPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserPasswordPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateRepositoryExistPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库是否存在

		Data *RepoIsExistResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateRepositoryExistPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateRepositoryExistPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateApplicationTokenPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 验证结果

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateApplicationTokenPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateApplicationTokenPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateNamespaceExistPersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ValidateNamespaceExistPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateNamespaceExistPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Region struct {

	// gz

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 1

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// ap-guangzhou

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// alluser

	Status *string `json:"Status,omitempty" name:"Status"`
	// remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DeleteImageBuildPersonalRequest struct {
	*tchttp.BaseRequest

	// 镜像列表

	Images []*string `json:"Images,omitempty" name:"Images"`
}

func (r *DeleteImageBuildPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 构建日志

		Data *BuildInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildTaskLogPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildTaskLogPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询信息

		Data *DesGitAuthRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeAuthPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerPersonalResp struct {

	// 返回条目总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 触发器列表

	TriggerInfo []*TriggerResp `json:"TriggerInfo,omitempty" name:"TriggerInfo"`
}

type SameImagesResp struct {

	// tag列表

	SameImages []*string `json:"SameImages,omitempty" name:"SameImages"`
}

type DeleteImageBuildPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全局自动删除策略信息

		Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageLifecycleGlobalPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalEndpointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网接入信息的列表

		AccessVpcSet []*AccessVpc `json:"AccessVpcSet,omitempty" name:"AccessVpcSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalEndpointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 构建仓库列表

		Data *BuildReposList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepoIsExistResp struct {

	// 仓库是否存在

	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
}

type CreateSecurityPoliciesRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 192.168.0.0/24

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildTaskPersonalRequest struct {
	*tchttp.BaseRequest

	// 构建任务ID

	BuildId *int64 `json:"BuildId,omitempty" name:"BuildId"`
}

func (r *DeleteImageBuildTaskPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildTaskPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTokenPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApplicationTokenPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTokenPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例的Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// PolicyId

	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
	// 192.168.0.0/24 白名单Ip

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateGitHubAuthPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateGitHubAuthPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateGitHubAuthPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20, 最大值 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationTriggerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// Tag名称

	TagName *string `json:"TagName,omitempty" name:"TagName"`
	// 镜像Id

	TagId *string `json:"TagId,omitempty" name:"TagId"`
	// docker image 可以看到的id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 大小

	Size *string `json:"Size,omitempty" name:"Size"`
	// 镜像的创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 镜像创建至今时间长度

	DurationDays *string `json:"DurationDays,omitempty" name:"DurationDays"`
	// 镜像的作者

	Author *string `json:"Author,omitempty" name:"Author"`
	// 次镜像建议运行的系统架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 创建此镜像的docker版本

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// 此镜像建议运行系统

	OS *string `json:"OS,omitempty" name:"OS"`
	// SizeByte

	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 数据更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 镜像更新时间

	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`
}

type DescribeApplicationTriggerLogPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数量，默认 20, 最大值 100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序或降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 按某列排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeApplicationTriggerLogPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationTriggerLogPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DuplicateImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 源镜像名称，不包含domain。例如： tencentyun/foo:v1

	SrcImage *string `json:"SrcImage,omitempty" name:"SrcImage"`
	// 目的镜像名称，不包含domain。例如： tencentyun/foo:latest

	DestImage *string `json:"DestImage,omitempty" name:"DestImage"`
}

func (r *DuplicateImagePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DuplicateImagePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 密码

	ApplicationToken *string `json:"ApplicationToken,omitempty" name:"ApplicationToken"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type BuildRepo struct {

	// 源码所在的git服务

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 源码在git服务器上的仓库名

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库Id

	RepoId *int64 `json:"RepoId,omitempty" name:"RepoId"`
	// 仓库描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 是否为私有

	Private *bool `json:"Private,omitempty" name:"Private"`
	// WebUrl

	WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`
}

type BatchDeleteFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDeleteFavorRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteFavorRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest

	// 触发器关联的镜像仓库，library/test格式

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 触发器名称

	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发

	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`
	// 触发方式对应的表达式

	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
	// 应用所在TKE集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 应用所在TKE集群命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 应用所在TKE集群工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 应用所在TKE集群工作负载下容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 应用所在TKE集群地域

	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *CreateApplicationTriggerPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRepoRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`
	// 是否公共,1:公共,0:私有

	Public *uint64 `json:"Public,omitempty" name:"Public"`
	// 仓库描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRepoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildTaskLogInfoPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 构建历史信息

		Data *BuildHistoryResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildTaskLogInfoPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildTaskLogInfoPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateUserPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 判断用户存在的返回值

		Data *UserIsExistsResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateUserPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateUserPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoDelStrategyInfo struct {

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 仓库名

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 策略值

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// Valid

	Valid *int64 `json:"Valid,omitempty" name:"Valid"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type TcrNamespaceInfoResp struct {

	// 命名空间信息列表

	NamespaceList []*TcrNamespaceInfo `json:"NamespaceList,omitempty" name:"NamespaceList"`
}

type ModifyRepositoryInfoPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRepositoryInfoPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepositoryInfoPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteImagePersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Tag列表

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *BatchDeleteImagePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteImagePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
	// 192.168.0.0/24

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteNamespacePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNamespacePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 代码仓库类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Page

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// PerPage

	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSourceCodeRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// GitServer地址

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// 用户信息

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// git 秘钥

	GitToken *string `json:"GitToken,omitempty" name:"GitToken"`
}

func (r *CreateSourceCodeAuthPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthUserInfoPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 源代码授权用户信息

		Data *AuthUserInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSourceCodeAuthUserInfoPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeAuthUserInfoPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationRule struct {

	// 同步规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 目标命名空间

	DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`
	// 是否覆盖

	Override *bool `json:"Override,omitempty" name:"Override"`
	// 同步过滤条件

	Filters []*ReplicationFilter `json:"Filters,omitempty" name:"Filters"`
}

type CreateInstanceTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户名
		Username *string `json:"Username,omitnil,omitempty" name:"Username"`

		// 访问凭证
		Token *string `json:"Token,omitnil,omitempty" name:"Token"`

		// 访问凭证过期时间戳，是一个时间戳数字，无单位
		ExpTime *int64 `json:"ExpTime,omitnil,omitempty" name:"ExpTime"`

		// 长期凭证的TokenId，短期凭证没有TokenId
		TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

		// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserPersonalRequest struct {
	*tchttp.BaseRequest

	// 用户密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest

	// Server 地址（GitHub或者GitLab）

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
}

func (r *DeleteSourceCodeAuthPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildHistoryResp struct {

	// 构建信息

	BuildHistory *BuildInfo `json:"BuildHistory,omitempty" name:"BuildHistory"`
}

type DescribeRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仓库信息

		Data *RepositoryInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskManuallyPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildTaskManuallyPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildTaskManuallyPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dockerhub仓库列表

		Data *RespDockerHubRepoList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceStatusRequest struct {
	*tchttp.BaseRequest

	// 实例ID的数组

	RegistryIds []*string `json:"RegistryIds,omitempty" name:"RegistryIds"`
}

func (r *DescribeInstanceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateRepositoryExistPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *ValidateRepositoryExistPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateRepositoryExistPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GitAuth struct {

	// appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// GitServer地址

	GitServer *string `json:"GitServer,omitempty" name:"GitServer"`
	// Scope

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 类型

	GitType *string `json:"GitType,omitempty" name:"GitType"`
}

type Trigger struct {

	// 分支

	Branches []*string `json:"Branches,omitempty" name:"Branches"`
	// Tag

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

type CreateImageBuildPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubImagePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询dockerhub仓库镜像列表的返回值

		Data *DockerHubTagList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubImagePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryInfoPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeDockerHubRepositoryInfoPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubRepositoryInfoPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总实例个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例信息列表

		Registries []*Registry `json:"Registries,omitempty" name:"Registries"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户命名空间返回信息

		Data *NamespaceInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeAuthPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSourceCodeAuthPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeAuthPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessVpc struct {

	// Vpc的Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 内网接入状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 内网接入Ip

	AccessIp *string `json:"AccessIp,omitempty" name:"AccessIp"`
}

type DeleteFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 仓库类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *DeleteFavorRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavorRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfoResp struct {

	// Tag的总数

	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`
	// TagInfo列表

	TagInfo []*TagInfo `json:"TagInfo,omitempty" name:"TagInfo"`
	// Server

	Server *string `json:"Server,omitempty" name:"Server"`
	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DeleteImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageLifecycleGlobalPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNamespacePersonalRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateNamespacePersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNamespacePersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespacePersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespacePersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset用于分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFavorRepositoryPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFavorRepositoryPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSourceCodeRepositoryBranchPersonalRequest struct {
	*tchttp.BaseRequest

	// 代码仓库类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// repo所在的group

	Group *string `json:"Group,omitempty" name:"Group"`
	// 源码在git服务器上的仓库名

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// Page

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// PerPage

	PerPage *int64 `json:"PerPage,omitempty" name:"PerPage"`
}

func (r *DescribeSourceCodeRepositoryBranchPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSourceCodeRepositoryBranchPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardRequestRequest struct {
	*tchttp.BaseRequest

	// 请求tcr对应的方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 请求tcr对应的路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 请求tcr中的HTTP头中Accept参数

	Accept *string `json:"Accept,omitempty" name:"Accept"`
	// 请求tcr http头中ContentType参数

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// 请求tcr中的body信息

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 请求的实例ID

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *ForwardRequestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ForwardRequestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DockerHubTagList struct {

	// DockerHub的仓库名称

	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`
	// Tag的列表

	TagList []*string `json:"TagList,omitempty" name:"TagList"`
}
type Registry struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 实例名称
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// 实例规格
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// 实例状态。有以下状态：
	// Pending, 初始化中
	// Deploying, 创建中
	// Running, 运行中
	// Unhealthy, 状态异常
	// FailedCreated, 创建失败
	// FailedUpdated, 更新失败
	// Bucket-Error, 存储桶异常
	// Isolate, 待回收
	// Deleting, 删除中
	// DeleteBucketFailed, 实例删除存储桶失败
	// DeleteFailed, 实例删除失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例的公共访问地址
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 实例创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域Id
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 是否支持匿名
	EnableAnonymous *bool `json:"EnableAnonymous,omitnil,omitempty" name:"EnableAnonymous"`

	// Token有效时间
	TokenValidTime *uint64 `json:"TokenValidTime,omitnil,omitempty" name:"TokenValidTime"`

	// 实例内部访问地址
	InternalEndpoint *string `json:"InternalEndpoint,omitnil,omitempty" name:"InternalEndpoint"`

	// 实例云标签
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例过期时间（预付费）
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// 实例付费类型，0表示后付费，1表示预付费
	PayMod *int64 `json:"PayMod,omitnil,omitempty" name:"PayMod"`

	// 预付费续费标识，0表示手动续费，1表示自动续费，2不续费并且不通知
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 是否开启实例删除保护，false表示不开启
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

// Predefined struct for user
type CreateInstanceTokenRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 访问凭证类型，longterm 为长期访问凭证，temp 为临时访问凭证，默认是临时访问凭证，有效期1小时
	TokenType *string `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// 长期访问凭证描述信息
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 访问凭证类型，longterm 为长期访问凭证，temp 为临时访问凭证，默认是临时访问凭证，有效期1小时
	TokenType *string `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// 长期访问凭证描述信息
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}


func (r *CreateInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRepositoryAllPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像仓库列表

		Data *RepoInfoResp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRepositoryAllPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryAllPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaPersonalRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserQuotaPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserQuotaPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RequestFavor struct {

	// 收藏的镜像仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// 收藏的镜像仓库类型

	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
	// 地域Id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type TriggerInvokePara struct {

	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// TKE集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// TKE集群命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// TKE集群工作负载名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// TKE集群工作负载中容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// TKE集群地域数字ID

	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type DescribeImageFilterPersonalRequest struct {
	*tchttp.BaseRequest

	// 仓库名称

	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
	// Tag名

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImageFilterPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageFilterPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFavorRepositoryPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFavorRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildTaskManuallyPersonalRequest struct {
	*tchttp.BaseRequest

	// registry所在的namespace

	RegistryNamespace *string `json:"RegistryNamespace,omitempty" name:"RegistryNamespace"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// Tag的格式

	TagFormat *string `json:"TagFormat,omitempty" name:"TagFormat"`
	// 仓库类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 仓库分支或者是commit

	BranchOrCommit *string `json:"BranchOrCommit,omitempty" name:"BranchOrCommit"`
}

func (r *CreateImageBuildTaskManuallyPersonalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildTaskManuallyPersonalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 企业版实例Id

		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationTriggerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSourceCodeAuthPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSourceCodeAuthPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSourceCodeAuthPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDockerHubRepositoryInfoPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DockerHub仓库信息

		Data *DockerHubRepoinfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDockerHubRepositoryInfoPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDockerHubRepositoryInfoPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationTriggerPersonalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Permission struct {
	// 资源路径，目前仅支持Namespace
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 动作，目前仅支持：tcr:PushRepository、tcr:PullRepository、tcr:CreateRepository、tcr:CreateHelmChart、tcr:DescribeHelmCharts
	Actions []*string `json:"Actions,omitnil,omitempty" name:"Actions"`
}

// Predefined struct for user
type CreateServiceAccountRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略列表
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 服务级账号描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 有效期(单位：天)，从当前时间开始计算，优先级高于ExpiresAt
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 过期时间（时间戳，单位:毫秒）
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 是否禁用服务级账号
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`
}

type CreateServiceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略列表
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 服务级账号描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 有效期(单位：天)，从当前时间开始计算，优先级高于ExpiresAt
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 过期时间（时间戳，单位:毫秒）
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 是否禁用服务级账号
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`
}

func (r *CreateServiceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceAccountResponseParams struct {
	// 服务级账号名（会自动加上前缀tcr$）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务级账号密码，仅展示一次，请注意留存
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 服务级账号失效时间（时间戳）
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 服务级账号创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceAccountResponseParams `json:"Response"`
}

func (r *CreateServiceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CVEWhitelistItem struct {
	// 漏洞白名单 ID
	CVEID *string `json:"CVEID,omitnil,omitempty" name:"CVEID"`
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称（长度2-30个字符，只能包含小写字母、数字及分隔符("."、"_"、"-")，且不能以分隔符开头、结尾或连续）
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 是否公开，true为公开，fale为私有
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 自动扫描级别，true为自动，false为手动
	IsAutoScan *bool `json:"IsAutoScan,omitnil,omitempty" name:"IsAutoScan"`

	// 安全阻断级别，true为自动，false为手动
	IsPreventVUL *bool `json:"IsPreventVUL,omitnil,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitnil,omitempty" name:"CVEWhitelistItems"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称（长度2-30个字符，只能包含小写字母、数字及分隔符("."、"_"、"-")，且不能以分隔符开头、结尾或连续）
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 是否公开，true为公开，fale为私有
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 自动扫描级别，true为自动，false为手动
	IsAutoScan *bool `json:"IsAutoScan,omitnil,omitempty" name:"IsAutoScan"`

	// 安全阻断级别，true为自动，false为手动
	IsPreventVUL *bool `json:"IsPreventVUL,omitnil,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitnil,omitempty" name:"CVEWhitelistItems"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 访问级别，True为公开，False为私有
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 扫描级别，True为自动，False为手动
	IsAutoScan *bool `json:"IsAutoScan,omitnil,omitempty" name:"IsAutoScan"`

	// 阻断开关，True为开放，False为关闭
	IsPreventVUL *bool `json:"IsPreventVUL,omitnil,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持 low、medium、high
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitnil,omitempty" name:"CVEWhitelistItems"`
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 访问级别，True为公开，False为私有
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 扫描级别，True为自动，False为手动
	IsAutoScan *bool `json:"IsAutoScan,omitnil,omitempty" name:"IsAutoScan"`

	// 阻断开关，True为开放，False为关闭
	IsPreventVUL *bool `json:"IsPreventVUL,omitnil,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持 low、medium、high
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitnil,omitempty" name:"CVEWhitelistItems"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNamespaceResponseParams `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespaceResponseParams `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认查询所有命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 每页个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页面偏移（第几页）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列出所有命名空间
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// 过滤条件
	// - 按照【标签】过滤
	//    Name: Tags
	//    Value:   tagKey:tagVal
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 仅查询启用了 KMS 镜像签名的空间
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitnil,omitempty" name:"KmsSignPolicy"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认查询所有命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 每页个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页面偏移（第几页）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列出所有命名空间
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// 过滤条件
	// - 按照【标签】过滤
	//    Name: Tags
	//    Value:   tagKey:tagVal
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 仅查询启用了 KMS 镜像签名的空间
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitnil,omitempty" name:"KmsSignPolicy"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// 命名空间列表信息
	NamespaceList []*TcrNamespaceInfo `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 仓库名称，需满足以下规则：
	// 1. 长度需 ​大于2且小于245个字符
	// 2. 仅允许小写字母、数字及符号 . _ -
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateRepositoryRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 仓库名称，需满足以下规则：
	// 1. 长度需 ​大于2且小于245个字符
	// 2. 仅允许小写字母、数字及符号 . _ -
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateRepositoryResponseParams `json:"Response"`
}

func (r *CreateRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyRepositoryRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRepositoryResponseParams `json:"Response"`
}

func (r *ModifyRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库的名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`
}

type DeleteRepositoryRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库的名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryResponseParams `json:"Response"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceAccount struct {
	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否禁用
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// 过期时间
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

// Predefined struct for user
type DescribeServiceAccountsRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 列出所有服务级账号
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// 是否填充权限信息
	EmbedPermission *bool `json:"EmbedPermission,omitnil,omitempty" name:"EmbedPermission"`

	// 过滤条件，key 目前只支持ServiceAccountName
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100（超出最大值，调整到最大值）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeServiceAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 列出所有服务级账号
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// 是否填充权限信息
	EmbedPermission *bool `json:"EmbedPermission,omitnil,omitempty" name:"EmbedPermission"`

	// 过滤条件，key 目前只支持ServiceAccountName
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100（超出最大值，调整到最大值）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeServiceAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceAccountsResponseParams struct {
	// 服务级账号列表
	ServiceAccounts []*ServiceAccount `json:"ServiceAccounts,omitnil,omitempty" name:"ServiceAccounts"`

	// 服务级账户数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceAccountsResponseParams `json:"Response"`
}

func (r *DescribeServiceAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务级账号描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 有效期(单位：天)，从当前时间开始计算，优先级高于ExpiresAt
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 过期时间（时间戳，单位:毫秒）
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 是否禁用服务级账号
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// 策略列表
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type ModifyServiceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务级账号描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 有效期(单位：天)，从当前时间开始计算，优先级高于ExpiresAt
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 过期时间（时间戳，单位:毫秒）
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 是否禁用服务级账号
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// 策略列表
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

func (r *ModifyServiceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceAccountResponseParams `json:"Response"`
}

func (r *ModifyServiceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceAccountRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteServiceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteServiceAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceAccountResponseParams `json:"Response"`
}

func (r *DeleteServiceAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionExecution struct {
	// 执行Id
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// 所属规则id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 执行的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 执行的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行的状态，Failed, Succeed, Stopped, InProgress
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeTagRetentionExecutionRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionExecutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionResponseParams struct {
	// 版本保留执行记录列表
	RetentionExecutionList []*RetentionExecution `json:"RetentionExecutionList,omitnil,omitempty" name:"RetentionExecutionList"`

	// 版本保留执行记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 版本保留规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 是否模拟执行，默认值为false，即非模拟执行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type CreateTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 版本保留规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 是否模拟执行，默认值为false，即非模拟执行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *CreateTagRetentionExecutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *CreateTagRetentionExecutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认为查询所有命名空间下镜像仓库
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 指定镜像仓库，不填写默认查询指定命名空间下所有镜像仓库
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 页数，第几页，用于分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页个数，用于分页，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 基于字段排序，支持的值有-creation_time,-name, -update_time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeRepositoriesRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认为查询所有命名空间下镜像仓库
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 指定镜像仓库，不填写默认查询指定命名空间下所有镜像仓库
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 页数，第几页，用于分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页个数，用于分页，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 基于字段排序，支持的值有-creation_time,-name, -update_time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

func (r *DescribeRepositoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesResponseParams struct {
	// 仓库信息列表
	RepositoryList []*TcrRepositoryInfo `json:"RepositoryList,omitnil,omitempty" name:"RepositoryList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoriesResponseParams `json:"Response"`
}

func (r *DescribeRepositoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTokenRequestParams struct {
	// 实例长期访问凭证 ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 启用或禁用实例长期访问凭证
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 访问凭证描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 1为修改描述 2为操作启动禁用，默认值为2
	ModifyFlag *int64 `json:"ModifyFlag,omitnil,omitempty" name:"ModifyFlag"`
}

type ModifyInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// 实例长期访问凭证 ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 启用或禁用实例长期访问凭证
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 访问凭证描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 1为修改描述 2为操作启动禁用，默认值为2
	ModifyFlag *int64 `json:"ModifyFlag,omitnil,omitempty" name:"ModifyFlag"`
}

func (r *ModifyInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceTokenResponseParams `json:"Response"`
}

func (r *ModifyInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenRequestParams struct {
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 访问凭证 ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`
}

type DeleteInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 访问凭证 ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`
}

func (r *DeleteInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceTokenResponseParams `json:"Response"`
}

func (r *DeleteInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenRequestParams struct {
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenResponseParams struct {
	// 长期访问凭证总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 长期访问凭证列表
	Tokens []*TcrInstanceToken `json:"Tokens,omitnil,omitempty" name:"Tokens"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTokenResponseParams `json:"Response"`
}

func (r *DescribeInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInternalEndpointDnsRequestParams struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitnil,omitempty" name:"EniLBIp"`

	// true：为默认域名，公网域名一致
	// false: 使用vpc域名
	// 默认为vpc域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitnil,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type CreateInternalEndpointDnsRequest struct {
	*tchttp.BaseRequest

	// tcr实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitnil,omitempty" name:"EniLBIp"`

	// true：为默认域名，公网域名一致
	// false: 使用vpc域名
	// 默认为vpc域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitnil,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

func (r *CreateInternalEndpointDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalEndpointDnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalEndpointDnsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInternalEndpointDnsResponse struct {
	*tchttp.BaseResponse
	Response *CreateInternalEndpointDnsResponseParams `json:"Response"`
}

func (r *CreateInternalEndpointDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInternalEndpointDnsRequestParams struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitnil,omitempty" name:"EniLBIp"`

	// true：使用默认域名
	// false:  使用带有vpc的域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitnil,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type DeleteInternalEndpointDnsRequest struct {
	*tchttp.BaseRequest

	// tcr实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitnil,omitempty" name:"EniLBIp"`

	// true：使用默认域名
	// false:  使用带有vpc的域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitnil,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

func (r *DeleteInternalEndpointDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInternalEndpointDnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInternalEndpointDnsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInternalEndpointDnsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInternalEndpointDnsResponseParams `json:"Response"`
}

func (r *DeleteInternalEndpointDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointDnsStatusRequestParams struct {
	// vpc列表
	VpcSet []*VpcAndDomainInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`
}

type DescribeInternalEndpointDnsStatusRequest struct {
	*tchttp.BaseRequest

	// vpc列表
	VpcSet []*VpcAndDomainInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`
}

func (r *DescribeInternalEndpointDnsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointDnsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcAndDomainInfo struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitnil,omitempty" name:"EniLBIp"`

	// true：use instance name as subdomain
	// false: use instancename+"-vpc" as subdomain
	UsePublicDomain *bool `json:"UsePublicDomain,omitnil,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type VpcPrivateDomainStatus struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// unique vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ENABLE代表已经开启，DISABLE代表未开启，ERROR代表查询出错
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeInternalEndpointDnsStatusResponseParams struct {
	// vpc私有域名解析状态列表
	VpcSet []*VpcPrivateDomainStatus `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInternalEndpointDnsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternalEndpointDnsStatusResponseParams `json:"Response"`
}

func (r *DescribeInternalEndpointDnsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointDnsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 复制实例地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// 复制实例地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`
}

type CreateReplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 复制实例地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// 复制实例地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`
}

func (r *CreateReplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceResponseParams struct {
	// 企业版复制实例Id
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplicationInstanceResponseParams `json:"Response"`
}

func (r *CreateReplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationRegistry struct {
	// 主实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// 复制实例的地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// 复制实例的地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// 复制实例的状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}


// Predefined struct for user
type DescribeReplicationInstancesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeReplicationInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstancesResponseParams struct {
	// 总实例个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 同步实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicationRegistries []*ReplicationRegistry `json:"ReplicationRegistries,omitnil,omitempty" name:"ReplicationRegistries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstancesResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceRequestParams struct {
	// 实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// 复制实例地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

type DeleteReplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// 复制实例地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

func (r *DeleteReplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReplicationInstanceResponseParams `json:"Response"`
}

func (r *DeleteReplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionRulesRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTagRetentionRulesRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionPolicy struct {
	// 版本保留策略Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则列表
	RetentionRuleList []*RetentionRule `json:"RetentionRuleList,omitnil,omitempty" name:"RetentionRuleList"`

	// 定期执行方式
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// 是否启用规则
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 基于当前时间根据cronSetting后下一次任务要执行的时间，仅做参考使用
	NextExecutionTime *string `json:"NextExecutionTime,omitnil,omitempty" name:"NextExecutionTime"`
}

type RetentionRule struct {
	// 支持的策略，可选值为latestPushedK（保留最新推送多少个版本）nDaysSinceLastPush（保留近天内推送）
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 规则设置下的对应值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type RetentionTask struct {
	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 所属的规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// 任务开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务的执行状态，Failed, Succeed, Stopped, InProgress
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 总tag数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 保留tag数
	Retained *int64 `json:"Retained,omitnil,omitempty" name:"Retained"`

	// 应用的仓库
	Repository *string `json:"Repository,omitnil,omitempty" name:"Repository"`
}

// Predefined struct for user
type DescribeTagRetentionRulesResponseParams struct {
	// 版本保留策略列表
	RetentionPolicyList []*RetentionPolicy `json:"RetentionPolicyList,omitnil,omitempty" name:"RetentionPolicyList"`

	// 版本保留策略总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionRulesResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 版本保留规则的Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`
}

type DeleteTagRetentionRuleRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 版本保留规则的Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`
}

func (r *DeleteTagRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagRetentionRuleResponseParams `json:"Response"`
}

func (r *DeleteTagRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeWebhookTriggerRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Header struct {
	// Header Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Header Values
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type WebhookTarget struct {
	// 目标地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 自定义 Headers
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type WebhookTrigger struct {
	// 触发器名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 触发器目标
	Targets []*WebhookTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 触发动作
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// 触发规则
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 启用触发器
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 触发器Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 触发器描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 触发器所属命名空间 Id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 触发器所属命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

type WebhookTriggerLog struct {
	// 日志 Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 触发器 Id
	TriggerId *int64 `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// 事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 通知类型
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 详情
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeWebhookTriggerResponseParams struct {
	// 触发器总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 触发器列表
	Triggers []*WebhookTrigger `json:"Triggers,omitnil,omitempty" name:"Triggers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteWebhookTriggerRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebhookTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebhookTriggerResponseParams `json:"Response"`
}

func (r *DeleteWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeWebhookTriggerLogRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeWebhookTriggerLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志列表
	Logs []*WebhookTriggerLog `json:"Logs,omitnil,omitempty" name:"Logs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerLogResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


// Predefined struct for user
type DescribeInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceCustomizedDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCustomizedDomainResponseParams struct {
	// 域名信息列表
	DomainInfoList []*CustomizedDomainInfo `json:"DomainInfoList,omitnil,omitempty" name:"DomainInfoList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DescribeInstanceCustomizedDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizedDomainInfo struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 域名名称
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 域名创建状态（SUCCESS, FAILURE, CREATING, DELETING）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteApplicationTriggerPersonalRequestParams struct {
	// 触发器名称
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`
}

// Predefined struct for user
type DeleteApplicationTriggerPersonalResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type DeleteImageAccelerateServiceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DeleteImageAccelerateServiceRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DeleteImageAccelerateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageAccelerateServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageAccelerateServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageAccelerateServiceResponseParams `json:"Response"`
}

func (r *DeleteImageAccelerateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageLifecycleGlobalPersonalRequestParams struct {

}

// Predefined struct for user
type DeleteImageLifecycleGlobalPersonalResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type DeleteImagePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type DeleteImagePersonalResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

// Predefined struct for user
type DeleteImageRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageResponseParams `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteImmutableTagRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DeleteImmutableTagRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DeleteInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DeleteInstanceCustomizedDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DeleteInstanceCustomizedDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 页数，默认为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页展示个数，最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 页数，默认为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页展示个数，最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeImmutableTagRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImmutableTagRule struct {
	// 仓库匹配规则
	RepositoryPattern *string `json:"RepositoryPattern,omitnil,omitempty" name:"RepositoryPattern"`

	// Tag 匹配规则
	TagPattern *string `json:"TagPattern,omitnil,omitempty" name:"TagPattern"`

	// repoMatches或repoExcludes
	RepositoryDecoration *string `json:"RepositoryDecoration,omitnil,omitempty" name:"RepositoryDecoration"`

	// matches或excludes
	TagDecoration *string `json:"TagDecoration,omitnil,omitempty" name:"TagDecoration"`

	// 禁用规则
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 命名空间
	NsName *string `json:"NsName,omitnil,omitempty" name:"NsName"`
}

type KeyValueString struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}


// Predefined struct for user
type ManageExternalEndpointRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 操作（Create/Delete）
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}


// Predefined struct for user
type DescribeImmutableTagRulesResponseParams struct {
	// 规则列表
	Rules []*ImmutableTagRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 未创建规则的命名空间
	EmptyNs []*string `json:"EmptyNs,omitnil,omitempty" name:"EmptyNs"`

	// 规则总量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DescribeImmutableTagRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 指定镜像版本进行查找，当前为模糊搜索
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 每页个数，用于分页，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，默认值为1
	// 补充说明：limit指的是每页的大小，offset指的是具体第几页。
	// 举例：limit 20 offset 1指的是1-20；limit 20 offset 2 指的是21-40；limit 30 offset 4 是指 90-120。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定镜像 Digest 进行查找
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// 指定是否为精准匹配，true为精准匹配，不填为模糊匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 指定镜像版本进行查找，当前为模糊搜索
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 每页个数，用于分页，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，默认值为1
	// 补充说明：limit指的是每页的大小，offset指的是具体第几页。
	// 举例：limit 20 offset 1指的是1-20；limit 20 offset 2 指的是21-40；limit 30 offset 4 是指 90-120。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定镜像 Digest 进行查找
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// 指定是否为精准匹配，true为精准匹配，不填为模糊匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 容器镜像信息列表
	ImageInfoList []*TcrImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// 容器镜像总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

type DescribeImageManifestsRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

func (r *DescribeImageManifestsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsResponseParams struct {
	// 镜像的Manifest信息
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 镜像的配置信息
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 镜像的Labels信息
	Labels []*KeyValueString `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 镜像大小，单位：byte
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageManifestsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageManifestsResponseParams `json:"Response"`
}

func (r *DescribeImageManifestsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// 页数，第几页，用于分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页个数，用于分页，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTagRetentionExecutionTaskRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// 页数，第几页，用于分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页个数，用于分页，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTagRetentionExecutionTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskResponseParams struct {
	// 版本保留执行任务列表
	RetentionTaskList []*RetentionTask `json:"RetentionTaskList,omitnil,omitempty" name:"RetentionTaskList"`

	// 版本保留执行任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionTaskResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


// Predefined struct for user
type ModifyServiceAccountPasswordRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否随机生成密码
	Random *bool `json:"Random,omitnil,omitempty" name:"Random"`

	// 服务级账号密码，长度在8到20之间且需包含至少一个大写字符，一个小写字符和一个数字
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModifyServiceAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 服务级账号名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否随机生成密码
	Random *bool `json:"Random,omitnil,omitempty" name:"Random"`

	// 服务级账号密码，长度在8到20之间且需包含至少一个大写字符，一个小写字符和一个数字
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModifyServiceAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountPasswordResponseParams struct {
	// 自定义用户密码，仅展示一次，请注意留存
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceAccountPasswordResponseParams `json:"Response"`
}

func (r *ModifyServiceAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateMultipleSecurityPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateMultipleSecurityPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteMultipleSecurityPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteMultipleSecurityPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// 执行周期，当前只能选择： manual;daily;weekly;monthly
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// 是否禁用规则，默认值为false
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

type CreateTagRetentionRuleRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// 执行周期，当前只能选择： manual;daily;weekly;monthly
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// 是否禁用规则，默认值为false
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

func (r *CreateTagRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionRuleResponseParams `json:"Response"`
}

func (r *CreateTagRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


// Predefined struct for user
type ModifyTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的Id，必须填写原有的命名空间id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// 执行周期，必须填写为原来的设置
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 是否禁用规则
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

type ModifyTagRetentionRuleRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间的Id，必须填写原有的命名空间id
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// 执行周期，必须填写为原来的设置
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// 是否禁用规则
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

func (r *ModifyTagRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTagRetentionRuleResponseParams `json:"Response"`
}

func (r *ModifyTagRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type CreateImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *CreateImmutableTagRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateImmutableTagRulesResponseParams `json:"Response"`
}

func (r *CreateImmutableTagRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyImmutableTagRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImmutableTagRulesResponseParams `json:"Response"`
}

func (r *ModifyImmutableTagRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CreateInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest

	// 主实例iD
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CreateInstanceCustomizedDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *CreateInstanceCustomizedDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

type CreateSignatureRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

func (r *CreateSignatureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSignatureResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignatureResponseParams `json:"Response"`
}

func (r *CreateSignatureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type CreateWebhookTriggerRequest struct {
	*tchttp.BaseRequest

	// 实例 Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *CreateWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerResponseParams struct {
	// 新建的触发器
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebhookTriggerResponseParams `json:"Response"`
}

func (r *CreateWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type ModifyWebhookTriggerRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *ModifyWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebhookTriggerResponseParams `json:"Response"`
}

func (r *ModifyWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

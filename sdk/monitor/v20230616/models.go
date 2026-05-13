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

package v20230616

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type NoticeContentTmplItem struct {

	// 自定义通知渠道配置

	CustomCallback []*CustomCallbackTmplMatcher `json:"CustomCallback,omitempty" name:"CustomCallback"`
}

type DeleteNoticeContentTmplsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNoticeContentTmplsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNoticeContentTmplsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNoticeContentTmplRequest struct {
	*tchttp.BaseRequest

	// 模版名称

	TmplName *string `json:"TmplName,omitempty" name:"TmplName"`
	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 模板内容

	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitempty" name:"TmplContents"`
	// 模板语言&nbsp;en/zh

	TmplLanguage *string `json:"TmplLanguage,omitempty" name:"TmplLanguage"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *CreateNoticeContentTmplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNoticeContentTmplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeContentTmplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNoticeContentTmplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeContentTmpl struct {

	// 模版id

	TmplID *string `json:"TmplID,omitempty" name:"TmplID"`
	// 模版名称

	TmplName *string `json:"TmplName,omitempty" name:"TmplName"`
	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 内容模版详情

	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitempty" name:"TmplContents"`
	// 模板语言&nbsp;en/zh

	TmplLanguage *string `json:"TmplLanguage,omitempty" name:"TmplLanguage"`
	// 创建时间，Unix时间戳，秒

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间，Unix时间戳，秒

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最后修改人

	LastModifier *string `json:"LastModifier,omitempty" name:"LastModifier"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 是否是预设&nbsp;0&nbsp;非预设&nbsp;1预设

	IsPreset *int64 `json:"IsPreset,omitempty" name:"IsPreset"`
}

type DescribeNoticeContentTmplRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TmplIDs []*string `json:"TmplIDs,omitempty" name:"TmplIDs"`
	// 模版名称

	TmplName *string `json:"TmplName,omitempty" name:"TmplName"`
	// 模板语言&nbsp;en/zh&nbsp;缺省不过滤

	TmplLanguage *string `json:"TmplLanguage,omitempty" name:"TmplLanguage"`
	// 分页数

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeNoticeContentTmplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeContentTmplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomCallbackTmplMatcher struct {

	// 匹配状态&nbsp;Invalid;&nbsp;Trigger&nbsp;;Recovery

	MatchingStatus []*string `json:"MatchingStatus,omitempty" name:"MatchingStatus"`
	// 模板配置

	Template *CustomCallbackNoticeTmpl `json:"Template,omitempty" name:"Template"`
}

type ModifyNoticeContentTmplRequest struct {
	*tchttp.BaseRequest

	// 模版名称

	TmplName *string `json:"TmplName,omitempty" name:"TmplName"`
	// 模版id

	TmplID *string `json:"TmplID,omitempty" name:"TmplID"`
	// 模板内容

	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitempty" name:"TmplContents"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *ModifyNoticeContentTmplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeContentTmplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultNoticeTmplRequest struct {
	*tchttp.BaseRequest

	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
}

func (r *DescribeDefaultNoticeTmplRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultNoticeTmplRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNoticeContentTmplsRequest struct {
	*tchttp.BaseRequest

	// 模版id

	TmplIDs []*string `json:"TmplIDs,omitempty" name:"TmplIDs"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DeleteNoticeContentTmplsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNoticeContentTmplsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeContentTmplBindPolicyCount struct {

	// 通知内容模板ID

	NoticeContentTmplID *string `json:"NoticeContentTmplID,omitempty" name:"NoticeContentTmplID"`
	// 绑定告警策略数量

	BindCount *uint64 `json:"BindCount,omitempty" name:"BindCount"`
}

type CustomCallbackNoticeTmpl struct {

	// 告警通知内容模板

	ContentTmpl *string `json:"ContentTmpl,omitempty" name:"ContentTmpl"`
	// 告警通知主题模板

	TitleTmpl *string `json:"TitleTmpl,omitempty" name:"TitleTmpl"`
}

type CreateNoticeContentTmplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义内容模板ID

		TmplID *string `json:"TmplID,omitempty" name:"TmplID"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNoticeContentTmplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultNoticeTmplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义通知内容模板

		NoticeContentTmpls []*NoticeContentTmpl `json:"NoticeContentTmpls,omitempty" name:"NoticeContentTmpls"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultNoticeTmplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultNoticeTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeContentTmplResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义通知内容模板

		NoticeContentTmpls []*NoticeContentTmpl `json:"NoticeContentTmpls,omitempty" name:"NoticeContentTmpls"`
		// 通知内容模板绑定的告警策略数量

		NoticeContentTmplBindPolicyCounts []*NoticeContentTmplBindPolicyCount `json:"NoticeContentTmplBindPolicyCounts,omitempty" name:"NoticeContentTmplBindPolicyCounts"`
		// 分页数

		PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
		// 分页大小

		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNoticeContentTmplResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

package v20190304

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CreatExporJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatExporJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatExporJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExporLoggingListRequest struct {
	*tchttp.BaseRequest

	// 页数

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页显示数量

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *GetExporLoggingListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExporLoggingListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatExporJobRequest struct {
	*tchttp.BaseRequest

	// 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 文件类型。取值：json、csv

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 存储类型。取值：cos、csp

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *CreatExporJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatExporJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookupEventsRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 每次最大查询数量

	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询分页

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 查询用户uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 查询属性细节

	LookupAttributes []*Attr `json:"LookupAttributes,omitempty" name:"LookupAttributes"`
	// 搜索类型

	LookupType *string `json:"LookupType,omitempty" name:"LookupType"`
	// 全局内容模糊查询

	ContentValue *string `json:"ContentValue,omitempty" name:"ContentValue"`
}

func (r *LookupEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LookupEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {

	// 事件

	Events []*Events `json:"Events,omitempty" name:"Events"`
	// 分页

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 分页是否结束

	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
}

type CreateBucketRequest struct {
	*tchttp.BaseRequest

	// 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 存储类型。取值：cos、csp

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *CreateBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resources struct {

	// 资源名字

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源中文名

	ResourceCnName *string `json:"ResourceCnName,omitempty" name:"ResourceCnName"`
	// 服务英文名

	ResourceEnName *string `json:"ResourceEnName,omitempty" name:"ResourceEnName"`
	// 资源地域

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 服务类型

	ResourceTypeName *string `json:"ResourceTypeName,omitempty" name:"ResourceTypeName"`
}

type GetExporLoggingListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExporLoggingListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExporLoggingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Events struct {

	// 事件内容

	CloudAuditEvent *string `json:"CloudAuditEvent,omitempty" name:"CloudAuditEvent"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件时间

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// SecretId

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 错误码

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 请求ID

	RequestID *uint64 `json:"RequestID,omitempty" name:"RequestID"`
	// 账户ID

	AccountID *string `json:"AccountID,omitempty" name:"AccountID"`
	// 源ip地址

	SourceIPAddress *string `json:"SourceIPAddress,omitempty" name:"SourceIPAddress"`
	// 事件源

	EventSource *string `json:"EventSource,omitempty" name:"EventSource"`
	// 事件域

	EventRegion *string `json:"EventRegion,omitempty" name:"EventRegion"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 资源内容

	Resources *Resources `json:"Resources,omitempty" name:"Resources"`
	// yunapi错误码

	ApiErrorCode *string `json:"ApiErrorCode,omitempty" name:"ApiErrorCode"`
	// yunapi错误信息

	ApiErrorMessage *string `json:"ApiErrorMessage,omitempty" name:"ApiErrorMessage"`
	// EventNameCn

	EventNameCn *string `json:"EventNameCn,omitempty" name:"EventNameCn"`
	// EventZhName

	EventZhName *string `json:"EventZhName,omitempty" name:"EventZhName"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 接口类型

	ReadWriteDetail *string `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`
	// ResourceRegion

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// ResourceTypeCn

	ResourceTypeCn *string `json:"ResourceTypeCn,omitempty" name:"ResourceTypeCn"`
	// 数字签名

	UkeySign *string `json:"UkeySign,omitempty" name:"UkeySign"`
}

type CreateBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExporJobRequest struct {
	*tchttp.BaseRequest

	// IDs

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteExporJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExporJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExporJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExporJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExporJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookupEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LookupEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LookupEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Attr struct {

	// 查询属性

	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`
	// 查询内容

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type Attribute struct {

	// 查询属性

	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`
	// 查询内容

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type ExportList struct {

	// 存储桶

	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`
	// 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 文件大小

	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 用户ID

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务状态 1:待处理，2:正在处理 3:已处理 4:错误

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 存储类型。取值：cos、csp

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 子账号ID

	SubAccountUin *uint64 `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Track struct {

	// 事件类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 接口名称

	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
	// 跟踪集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 同步开关：0=未开启，1=开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 存储位置

	Storage *TrackStorage `json:"Storage,omitempty" name:"Storage"`
	// 跟踪集id

	TrackId *int64 `json:"TrackId,omitempty" name:"TrackId"`
}

type TrackStorage struct {

	// 被指定存储用户ID

	StorageAccountId *int64 `json:"StorageAccountId,omitempty" name:"StorageAccountId"`
	// 被指定存储用户appid

	StorageAppId *int64 `json:"StorageAppId,omitempty" name:"StorageAppId"`
	// 加密方式

	StorageEncryption *int64 `json:"StorageEncryption,omitempty" name:"StorageEncryption"`
	// bucket名称

	StorageName *string `json:"StorageName,omitempty" name:"StorageName"`
	// 日志文件前缀

	StoragePrefix *string `json:"StoragePrefix,omitempty" name:"StoragePrefix"`
	// 所属地域

	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`
	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

type CreateAuditTrackRequest struct {
	*tchttp.BaseRequest

	// 事件类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 接口名称

	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
	// 跟踪集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 同步开关：0=未开启；1=开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 存储位置

	Storage *TrackStorage `json:"Storage,omitempty" name:"Storage"`
	// 是否投递到集团账号

	TrackForAllMembers *int64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
}

func (r *CreateAuditTrackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAuditTrackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuditTrackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跟踪集id

		TrackId *int64 `json:"TrackId,omitempty" name:"TrackId"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditTrackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditTrackRequest struct {
	*tchttp.BaseRequest

	// 跟踪集id

	TrackId *int64 `json:"TrackId,omitempty" name:"TrackId"`
}

func (r *DeleteAuditTrackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAuditTrackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditTrackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditTrackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditTrackRequest struct {
	*tchttp.BaseRequest

	// 跟踪集id

	TrackId *int64 `json:"TrackId,omitempty" name:"TrackId"`
}

func (r *DescribeAuditTrackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditTrackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditTrackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件类型

		ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 接口名称

		EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
		// 补齐id

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// 跟踪集名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 资源类型

		ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
		// 同步开关：0=未开启，1=开启

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 存储位置

		Storage *TrackStorage `json:"Storage,omitempty" name:"Storage"`
		// 是否投递集团账号

		TrackForAllMembers *int64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditTrackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditTracksRequest struct {
	*tchttp.BaseRequest

	// 当前页码

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 页面大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeAuditTracksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditTracksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditTracksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 跟踪集列表

		Tracks []*Track `json:"Tracks,omitempty" name:"Tracks"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditTracksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAuditTracksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest

	// 全局内容模糊查询

	ContentValue *string `json:"ContentValue,omitempty" name:"ContentValue"`
	// 查询结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 查询属性细节

	LookupAttributes []*Attr `json:"LookupAttributes,omitempty" name:"LookupAttributes"`
	// 搜索类型

	LookupType *string `json:"LookupType,omitempty" name:"LookupType"`
	// 每次最大查询数量

	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 查询分页

	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
	// 查询用户uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 查询开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件

		Events []*Events `json:"Events,omitempty" name:"Events"`
		// 是否最后一页

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 下一页token

		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
		// ReturnMessage

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExportTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetExportTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExportTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExportTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的message

		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
		// 返回的value

		ReturnValue *string `json:"ReturnValue,omitempty" name:"ReturnValue"`
		// 导出类型 1:cos 2:csp

		Type *uint64 `json:"Type,omitempty" name:"Type"`
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExportTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExportTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditTrackRequest struct {
	*tchttp.BaseRequest

	// 事件类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 接口名称

	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
	// 跟踪集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 同步开关：0=未开启；1=可开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 存储位置

	Storage *TrackStorage `json:"Storage,omitempty" name:"Storage"`
	// 是否投递集团账号

	TrackForAllMembers *int64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
	// 跟踪集id

	TrackId *int64 `json:"TrackId,omitempty" name:"TrackId"`
}

func (r *ModifyAuditTrackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuditTrackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAuditTrackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAuditTrackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

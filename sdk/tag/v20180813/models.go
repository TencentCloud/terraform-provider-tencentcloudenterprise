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

package v20180813

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribeResourceMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据位移偏量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 每页大小

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 资源标签

		Rows []*ResourceIdTag `json:"Rows,omitempty" name:"Rows"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesByTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagsRequest struct {
	*tchttp.BaseRequest

	// 创建标签者uin

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 数据偏移量，默认为&nbsp;0,&nbsp;必须为Limit参数的整数倍

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小，默认为&nbsp;15

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 资源唯一标记

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源所在地域

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 业务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 标签过滤数组

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 是否忽略标签过滤数组（0：否；1：是），忽略后传入TagFilters参数无效，默认为0

	IgnoreFilter *uint64 `json:"IgnoreFilter,omitempty" name:"IgnoreFilter"`
	// CapiAction

	CapiAction *string `json:"CapiAction,omitempty" name:"CapiAction"`
	// CapiServiceType

	CapiServiceType *string `json:"CapiServiceType,omitempty" name:"CapiServiceType"`
	// CapiTransformInFunc

	CapiTransformInFunc *string `json:"CapiTransformInFunc,omitempty" name:"CapiTransformInFunc"`
	// CapiTransformOutFunc

	CapiTransformOutFunc *string `json:"CapiTransformOutFunc,omitempty" name:"CapiTransformOutFunc"`
	// ServiceName

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// ResourceType

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourcesByTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesByTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByTagKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据位移偏量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 每页大小

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 资源标签

		Rows []*ResourceIdTag `json:"Rows,omitempty" name:"Rows"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsByTagKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsByTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTagMapping struct {

	// 资源六段式。云使用资源六段式描述一个资源。
	// 例如：ResourceList.1&nbsp;=&nbsp;qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。

	Resource *string `json:"Resource,omitempty" name:"Resource"`
	// 资源关联的标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type DescribeResourceMenuRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest

	// 键

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceTagRequest struct {
	*tchttp.BaseRequest

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 资源六段式描述

	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *DeleteResourceTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTagRequest struct {
	*tchttp.BaseRequest

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *CreateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据位移偏量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 每页大小

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 资源标签

		Rows []*TagResource `json:"Rows,omitempty" name:"Rows"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateTagRequest struct {
	*tchttp.BaseRequest

	// 标签列表

	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`
}

func (r *BatchCreateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// 标签键,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值,与标签键同时存在或同时不存在，不存在时表示查询该用户所有标签

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 数据偏移量，默认为&nbsp;0,&nbsp;必须为Limit参数的整数倍

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小，默认为&nbsp;15

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 创建者用户&nbsp;Uin，不传或为空只将&nbsp;Uin&nbsp;作为条件查询

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 标签键数组,与标签值同时存在或同时不存在，不存在时表示查询该用户所有标签,当与TagKey同时传递时只会本值

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败资源信息。
		// 创建并绑定标签成功时，返回的FailedResources为空。
		// 创建并绑定标签失败或部分失败时，返回的FailedResources会显示失败资源的详细信息。

		FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TagResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagResourcesRequest struct {
	*tchttp.BaseRequest

	// 待绑定的云资源，用标准的资源六段式表示。正确的资源六段式请参考：
	// N取值范围：0~9

	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`
	// 标签键和标签值。
	// 如果指定多个标签，则会为指定资源同时创建并绑定该多个标签。
	// 同一个资源上的同一个标签键只能对应一个标签值。如果您尝试添加已有标签键，则对应的标签值会更新为新值。
	// 如果标签不存在会为您自动创建标签。
	// N取值范围：0~9

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否通过TagResourcesAllocateQuotas预分配配额。
	// -&nbsp;1：否（缺省值）
	// -&nbsp;2：是

	IsAllocatedQuotas *uint64 `json:"IsAllocatedQuotas,omitempty" name:"IsAllocatedQuotas"`
}

func (r *TagResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TagResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceTypeFilter struct {

	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix []*string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type TagEntry struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type AddResourceTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddResourceTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据位移偏量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 每页大小

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 标签列表

		Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceIdTag struct {

	// 资源唯一标识

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 标签键值对

	TagKeyValues *string `json:"TagKeyValues,omitempty" name:"TagKeyValues"`
}

type DeleteResourceTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByResourceIdsRequest struct {
	*tchttp.BaseRequest

	// 业务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 资源唯一标记

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 资源所在地域

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 数据偏移量，默认为&nbsp;0,&nbsp;必须为Limit参数的整数倍

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小，默认为&nbsp;15

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceTagsByResourceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsByResourceIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourcesRequest struct {
	*tchttp.BaseRequest

	// 资源六段式列表。云使用资源六段式描述一个资源。
	// 例如：ResourceList.1&nbsp;=&nbsp;qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。
	// 如果传入了此参数会返回所有匹配的资源列表，指定的MaxResults会失效。
	// N取值范围：0~9

	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`
	// 标签键和标签值。
	// 指定多个标签，会查询同时绑定了该多个标签的资源。
	// N取值范围：0~5。
	// 每个TagFilters中的TagValue最多支持10个

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。

	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`
	// 每一页返回的数据最大条数，最大200。
	// 缺省值：50。

	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
	// 业务类型，资源六段式的第3段。如果传入ResourceList参数，此参数会被忽略。

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源所在地域，资源六段式的第4段。如果传入ResourceList参数，此参数会被忽略。

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 资源前缀，资源六段式的第6段"/"前的内容。如果传入ResourceList参数，此参数会被忽略。

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *GetResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnTagResourcesRequest struct {
	*tchttp.BaseRequest

	// 标签键。
	// 取值范围：0~9

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
	// 资源六段式列表。云使用资源六段式描述一个资源。
	// 例如：ResourceList.1&nbsp;=&nbsp;qcs::${ServiceType}:${Region}:uin/${Account}:${ResourcePrefix}/${ResourceId}。
	// N取值范围：0~9

	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`
}

func (r *UnTagResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnTagResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedResource struct {

	// 失败的资源六段式

	Resource *string `json:"Resource,omitempty" name:"Resource"`
	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type TagWithDelete struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 是否可以删除

	CanDelete *uint64 `json:"CanDelete,omitempty" name:"CanDelete"`
}

type DescribeResourcesBindTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Limit

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// Offset

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 标签列表

		Rows *TagResource `json:"Rows,omitempty" name:"Rows"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesBindTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesBindTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagKeyObject struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type CreateTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 每页数量

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 标签

		Tags []*string `json:"Tags,omitempty" name:"Tags"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 每页数量

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 标签条目

		Tags []*TagEntry `json:"Tags,omitempty" name:"Tags"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest

	// 资源的六段式描述

	Resource *string `json:"Resource,omitempty" name:"Resource"`
	// 需要增加或修改的标签集合。如果Resource描述的资源未关联输入的标签键，则增加关联；若已关联，则将该资源关联的键对应的标签值修改为输入值。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键

	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags"`
	// 需要解关联的标签集合。本接口中ReplaceTags和DeleteTags二者必须存在其一，且二者不能包含相同的标签键

	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsRequest struct {
	*tchttp.BaseRequest

	// 创建者uin

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 资源所在地域

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 业务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 资源唯一标识

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 数据偏移量，默认为&nbsp;0,&nbsp;必须为Limit参数的整数倍

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小，默认为&nbsp;15

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否是Cos的资源id

	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 标签类型。取值：&nbsp;Custom：自定义标签。&nbsp;System：系统标签。&nbsp;All：全部标签。&nbsp;默认值：All。

	Category *string `json:"Category,omitempty" name:"Category"`
}

type DescribeResourceTagsByResourceIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据位移偏量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 每页大小

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 标签列表

		Tags []*TagResource `json:"Tags,omitempty" name:"Tags"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsByResourceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagKeysRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTagKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取的下一页的Token值

		PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`
		// 资源及关联的标签(键和值)列表

		ResourceTagMappingList []*ResourceTagMapping `json:"ResourceTagMappingList,omitempty" name:"ResourceTagMappingList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchCreateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesBindTagRequest struct {
	*tchttp.BaseRequest

	// 地域列表

	ResourceRegions []*string `json:"ResourceRegions,omitempty" name:"ResourceRegions"`
	// 服务列表

	ServiceTypeFilters []*ServiceTypeFilter `json:"ServiceTypeFilters,omitempty" name:"ServiceTypeFilters"`
	// 开始

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 标签过滤

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeResourcesBindTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesBindTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByTagKeysRequest struct {
	*tchttp.BaseRequest

	// 业务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 资源地域

	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
	// 资源唯一标识

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 资源标签键

	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
	// 每页大小，默认为&nbsp;400

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移量，默认为&nbsp;0,&nbsp;必须为Limit参数的整数倍

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 云api名

	CapiAction *string `json:"CapiAction,omitempty" name:"CapiAction"`
	// 云api服务类型

	CapiServiceType *string `json:"CapiServiceType,omitempty" name:"CapiServiceType"`
	// 云api输入

	CapiTransformInFunc *string `json:"CapiTransformInFunc,omitempty" name:"CapiTransformInFunc"`
	// 云api输出

	CapiTransformOutFunc *string `json:"CapiTransformOutFunc,omitempty" name:"CapiTransformOutFunc"`
	// 资源名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeResourceTagsByTagKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceTagsByTagKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnTagResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败资源信息。
		// 解绑标签成功时，返回的FailedResources为空。
		// 解绑标签失败或部分失败时，返回的FailedResources会显示失败资源的详细信息。

		FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnTagResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnTagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值数组&nbsp;多个值的话是或的关系

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagResource struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 标签键MD5值

	TagKeyMd5 *string `json:"TagKeyMd5,omitempty" name:"TagKeyMd5"`
	// 标签值MD5值

	TagValueMd5 *string `json:"TagValueMd5,omitempty" name:"TagValueMd5"`
	// 资源类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type AddResourceTagRequest struct {
	*tchttp.BaseRequest

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// 资源六段式描述

	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *AddResourceTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagRequest struct {
	*tchttp.BaseRequest

	// 需要删除的标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 需要删除的标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *DeleteTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

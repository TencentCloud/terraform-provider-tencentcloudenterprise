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

package v20200107

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetStatDayRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 统计信息日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 存储类型 1:标准存储 2:低频存储

	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *GetStatDayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatDayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBucketsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认返回全量。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶字段搜索条件，模糊匹配

	BucketFilter []*string `json:"BucketFilter,omitempty" name:"BucketFilter"`
}

func (r *ListBucketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBucketsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBucketsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的存储桶数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的存储桶列表

		Buckets []*Bucket `json:"Buckets,omitempty" name:"Buckets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBucketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutBucketRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 定义存储桶的访问控制列表（ACL）属性。枚举值请参见 ACL 概述 文档中存储桶的预设 ACL 部分，例如 private，public-read 等，默认为 private

	ACL *string `json:"ACL,omitempty" name:"ACL"`
	// 赋予被授权者读取存储桶的权限，格式：id="[OwnerUin]"，可使用半角逗号（,）分隔多组被授权者： 当需要给子账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>" 当需要给主账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>" 例如'id="qcs::cam::uin/100000000001:uin/100000000001", id="qcs::cam::uin/100000000001:uin/100000000011"'

	GrantRead *string `json:"GrantRead,omitempty" name:"GrantRead"`
	// 赋予被授权者写入存储桶的权限，格式：id="[OwnerUin]"，可使用半角逗号（,）分隔多组被授权者： 当需要给子账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>" 当需要给主账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>" 例如'id="qcs::cam::uin/100000000001:uin/100000000001", id="qcs::cam::uin/100000000001:uin/100000000011"'

	GrantWrite *string `json:"GrantWrite,omitempty" name:"GrantWrite"`
	// 赋予被授权者读取存储桶的访问控制列表（ACL）和存储桶策略（Policy）的权限。可使用半角逗号（,）分隔多组被授权者： 当需要给子账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>" 当需要给主账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>" 例如'id="qcs::cam::uin/100000000001:uin/100000000001", id="qcs::cam::uin/100000000001:uin/100000000011"'

	GrantReadAcp *string `json:"GrantReadAcp,omitempty" name:"GrantReadAcp"`
	// 赋予被授权者写入存储桶的访问控制列表（ACL）和存储桶策略（Policy）的权限，可使用半角逗号（,）分隔多组被授权者： 当需要给子账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>" 当需要给主账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>" 例如'id="qcs::cam::uin/100000000001:uin/100000000001",id="qcs::cam::uin/100000000001:uin/100000000011"'

	GrantWriteAcp *string `json:"GrantWriteAcp,omitempty" name:"GrantWriteAcp"`
	// 赋予被授权者操作存储桶的所有权限，格式：id="[OwnerUin]"，可使用半角逗号（,）分隔多组被授权者： 当需要给子账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<SubUin>" 当需要给主账号授权时，id="qcs::cam::uin/<OwnerUin>:uin/<OwnerUin>" 例如'id="qcs::cam::uin/100000000001:uin/100000000001", id="qcs::cam::uin/100000000001:uin/100000000011"'

	GrantFullControl *string `json:"GrantFullControl,omitempty" name:"GrantFullControl"`
}

func (r *PutBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPriceRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HeadObjectRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 对象键（Object 的名称），对象在存储桶中的唯一标识

	Key *string `json:"Key,omitempty" name:"Key"`
	// 对象的VersionId

	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *HeadObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeadObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Bucket struct {

	// 存储桶名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 存储桶所在 COS 地域

	Location *string `json:"Location,omitempty" name:"Location"`
	// 存储桶创建时间

	CreationDate *string `json:"CreationDate,omitempty" name:"CreationDate"`
}

type DeleteObjectObjects struct {

	// 对象键（Object 的名称），对象在存储桶中的唯一标识

	Key *string `json:"Key,omitempty" name:"Key"`
	// 要删除的对象版本 ID 或 DeleteMarker 版本 ID

	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type HeadObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HeadObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeadObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatHttpDayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatHttpDayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatHttpDayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatHttpDaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatHttpDaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatHttpDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultipleObjectRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 要删除的对象列表

	Objects []*DeleteObjectObjects `json:"Objects,omitempty" name:"Objects"`
}

func (r *DeleteMultipleObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMultipleObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatHttpDaysRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 统计信息开始日期

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 统计信息结束日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 存储类型，1: 标准存储; 2: 低频存储

	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *GetStatHttpDaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatHttpDaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosProxyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CosProxyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CosProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
}

func (r *GetServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatHttpDayRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 统计信息日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 存储类型，1: 标准存储; 2: 低频存储

	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *GetStatHttpDayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatHttpDayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteObjectRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 对象键（Object 的名称），对象在存储桶中的唯一标识

	ObjectKey *string `json:"ObjectKey,omitempty" name:"ObjectKey"`
	// 要删除的对象版本 ID 或 DeleteMarker 版本 ID

	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeleteObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
}

func (r *GetOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenCosBillingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenCosBillingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenCosBillingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatDaysRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 统计信息开始日期

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 统计信息结束日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 存储类型 1: 标准存储 2: 低频存储

	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *GetStatDaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatDaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutBucketRefererRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 防盗链配置信息

	RefererConfiguration *RefererConfiguration `json:"RefererConfiguration,omitempty" name:"RefererConfiguration"`
}

func (r *PutBucketRefererRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutBucketRefererRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosProxyRequest struct {
	*tchttp.BaseRequest

	// 操作链接

	Url *string `json:"Url,omitempty" name:"Url"`
	// 操作方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 操作 Header 参数，编码为 JSON 字符串

	Headers *string `json:"Headers,omitempty" name:"Headers"`
	// 操作 Url 参数，编码为 JSON 字符串

	Query *string `json:"Query,omitempty" name:"Query"`
	// 操作内容

	Body *string `json:"Body,omitempty" name:"Body"`
	// Body格式类型，可以传 base64、utf8，默认 utf8 字符串

	BodyType *string `json:"BodyType,omitempty" name:"BodyType"`
}

func (r *CosProxyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CosProxyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBillingTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBillingTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatDayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatDayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatDayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserStatRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenCosBillingRequest struct {
	*tchttp.BaseRequest
}

func (r *OpenCosBillingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenCosBillingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefererConfiguration struct {

	// 是否开启防盗链，枚举值：Enabled、Disabled

	Status *string `json:"Status,omitempty" name:"Status"`
	// 防盗链类型，枚举值：Black-List、White-List

	RefererType *string `json:"RefererType,omitempty" name:"RefererType"`
	// 是否允许空 Referer 访问，枚举值：Allow、Deny，默认值为 Deny

	EmptyReferConfiguration *string `json:"EmptyReferConfiguration,omitempty" name:"EmptyReferConfiguration"`
	// 生效域名列表， 支持多个域名且为前缀匹配， 支持带端口的域名和 IP， 支持通配符*，做二级域名或多级域名的通配

	DomainList *RefererConfigurationDomainList `json:"DomainList,omitempty" name:"DomainList"`
}

type GetBucketRefererResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBucketRefererResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketRefererResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCosStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutBucketRefererResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutBucketRefererResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutBucketRefererResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultipleObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMultipleObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMultipleObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBucketRefererRequest struct {
	*tchttp.BaseRequest

	// 临时密钥的 tmpSecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时密钥的 tmpSecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
	// 临时密钥的 sessionToken

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 存储桶所在的 COS 地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
	// 存储桶的名称，命名格式为 BucketName-APPID

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

func (r *GetBucketRefererRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketRefererRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefererConfigurationDomainList struct {

	// 生效域名列表， 支持多个域名且为前缀匹配， 支持带端口的域名和 IP， 支持通配符*，做二级域名或多级域名的通配

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type GetBillingTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前cos部署是否为计费

		IsBilling *bool `json:"IsBilling,omitempty" name:"IsBilling"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBillingTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatDaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatDaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCosStatRequest struct {
	*tchttp.BaseRequest

	// 存储桶所在地域

	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
}

func (r *GetCosStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

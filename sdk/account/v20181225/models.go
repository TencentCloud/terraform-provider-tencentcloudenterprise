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

package v20181225

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type TestDescRequest struct {
	*tchttp.BaseRequest

	// 数据

	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *TestDescRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestDescRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Description

		Description *string `json:"Description,omitempty" name:"Description"`
		// DisplayName

		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
		// UserGroupId

		UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MatchedWhitelist struct {

	// WhitelistUinList

	WhitelistUinList []*string `json:"WhitelistUinList,omitempty" name:"WhitelistUinList"`
	// WhitelistKey

	WhitelistKey *string `json:"WhitelistKey,omitempty" name:"WhitelistKey"`
}

type DescribeUserGroupRequest struct {
	*tchttp.BaseRequest

	// UserGroupId

	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *DescribeUserGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNicknameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DisplayName

		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
		// Nickname

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNicknameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNicknameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCheckWhitelistRequest struct {
	*tchttp.BaseRequest

	// WhitelistKeyList

	WhitelistKeyList []*string `json:"WhitelistKeyList,omitempty" name:"WhitelistKeyList"`
	// WhitelistUinList

	WhitelistUinList []*string `json:"WhitelistUinList,omitempty" name:"WhitelistUinList"`
	// Platform

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// DefaultWhitelistUinFlag

	DefaultWhitelistUinFlag *int64 `json:"DefaultWhitelistUinFlag,omitempty" name:"DefaultWhitelistUinFlag"`
}

func (r *BatchCheckWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCheckWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByUinRequest struct {
	*tchttp.BaseRequest

	// OpUin

	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`
}

func (r *GetAppIdByUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCheckWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MatchedWhitelist

		MatchedWhitelist []*MatchedWhitelist `json:"MatchedWhitelist,omitempty" name:"MatchedWhitelist"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchCheckWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCheckWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNicknameRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNicknameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNicknameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAreaByUinRequest struct {
	*tchttp.BaseRequest

	// OpUin

	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`
}

func (r *GetAreaByUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAreaByUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AppId

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdByUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestDescResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestDescResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAreaByUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Area

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// CountryName

		CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
		// CountryCode

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAreaByUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAreaByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

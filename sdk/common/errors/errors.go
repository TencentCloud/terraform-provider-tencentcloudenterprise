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
package errors

import (
	"fmt"
)

type CloudSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *CloudSDKError) Error() string {
	return fmt.Sprintf("[TCECloudSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewCloudSDKError(code, message, requestId string) error {
	return &CloudSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *CloudSDKError) GetCode() string {
	return e.Code
}

func (e *CloudSDKError) GetMessage() string {
	return e.Message
}

func (e *CloudSDKError) GetRequestId() string {
	return e.RequestId
}

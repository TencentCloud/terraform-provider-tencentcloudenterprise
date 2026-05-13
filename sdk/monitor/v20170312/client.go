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

package v20170312

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2017-03-12"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewCancelStarUnifyDashboardRequest() (request *CancelStarUnifyDashboardRequest) {
	request = &CancelStarUnifyDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "CancelStarUnifyDashboard")
	return
}

func NewCancelStarUnifyDashboardResponse() (response *CancelStarUnifyDashboardResponse) {
	response = &CancelStarUnifyDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消收藏 dashboard
func (c *Client) CancelStarUnifyDashboard(request *CancelStarUnifyDashboardRequest) (response *CancelStarUnifyDashboardResponse, err error) {
	if request == nil {
		request = NewCancelStarUnifyDashboardRequest()
	}
	response = NewCancelStarUnifyDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewSetDashboardStatusRequest() (request *SetDashboardStatusRequest) {
	request = &SetDashboardStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "SetDashboardStatus")
	return
}

func NewSetDashboardStatusResponse() (response *SetDashboardStatusResponse) {
	response = &SetDashboardStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置dashboard状态
func (c *Client) SetDashboardStatus(request *SetDashboardStatusRequest) (response *SetDashboardStatusResponse, err error) {
	if request == nil {
		request = NewSetDashboardStatusRequest()
	}
	response = NewSetDashboardStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
	request = &DescribeBaseMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
	return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
	response = &DescribeBaseMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基础指标详情
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeBaseMetricsRequest()
	}
	response = NewDescribeBaseMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
	request = &GetMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
	return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
	response = &GetMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云产品的监控数据。传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
// 接口调用频率限制为：50次/秒，500次/分钟。
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
	if request == nil {
		request = NewGetMonitorDataRequest()
	}
	response = NewGetMonitorDataResponse()
	err = c.Send(request, response)
	return
}

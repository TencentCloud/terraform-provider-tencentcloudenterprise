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

package v20201214

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-12-14"

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

func NewDescribeUserDefinedOrderDefaultValuesRequest() (request *DescribeUserDefinedOrderDefaultValuesRequest) {
	request = &DescribeUserDefinedOrderDefaultValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeUserDefinedOrderDefaultValues")
	return
}

func NewDescribeUserDefinedOrderDefaultValuesResponse() (response *DescribeUserDefinedOrderDefaultValuesResponse) {
	response = &DescribeUserDefinedOrderDefaultValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户自定义订单各项配置属性的默认值、初始值、范围等
func (c *Client) DescribeUserDefinedOrderDefaultValues(request *DescribeUserDefinedOrderDefaultValuesRequest) (response *DescribeUserDefinedOrderDefaultValuesResponse, err error) {
	if request == nil {
		request = NewDescribeUserDefinedOrderDefaultValuesRequest()
	}
	response = NewDescribeUserDefinedOrderDefaultValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSitesRequest() (request *DescribeSitesRequest) {
	request = &DescribeSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeSites")
	return
}

func NewDescribeSitesResponse() (response *DescribeSitesResponse) {
	response = &DescribeSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点列表
func (c *Client) DescribeSites(request *DescribeSitesRequest) (response *DescribeSitesResponse, err error) {
	if request == nil {
		request = NewDescribeSitesRequest()
	}
	response = NewDescribeSitesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSitesRequest() (request *DeleteSitesRequest) {
	request = &DeleteSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DeleteSites")
	return
}

func NewDeleteSitesResponse() (response *DeleteSitesResponse) {
	response = &DeleteSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除站点
func (c *Client) DeleteSites(request *DeleteSitesRequest) (response *DeleteSitesResponse, err error) {
	if request == nil {
		request = NewDeleteSitesRequest()
	}
	response = NewDeleteSitesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterInstanceTypesRequest() (request *DescribeDedicatedClusterInstanceTypesRequest) {
	request = &DescribeDedicatedClusterInstanceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterInstanceTypes")
	return
}

func NewDescribeDedicatedClusterInstanceTypesResponse() (response *DescribeDedicatedClusterInstanceTypesResponse) {
	response = &DescribeDedicatedClusterInstanceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群支持的实例规格列表
func (c *Client) DescribeDedicatedClusterInstanceTypes(request *DescribeDedicatedClusterInstanceTypesRequest) (response *DescribeDedicatedClusterInstanceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterInstanceTypesRequest()
	}
	response = NewDescribeDedicatedClusterInstanceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterCosCapacityRequest() (request *DescribeDedicatedClusterCosCapacityRequest) {
	request = &DescribeDedicatedClusterCosCapacityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterCosCapacity")
	return
}

func NewDescribeDedicatedClusterCosCapacityResponse() (response *DescribeDedicatedClusterCosCapacityResponse) {
	response = &DescribeDedicatedClusterCosCapacityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群内cos的容量信息
func (c *Client) DescribeDedicatedClusterCosCapacity(request *DescribeDedicatedClusterCosCapacityRequest) (response *DescribeDedicatedClusterCosCapacityResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterCosCapacityRequest()
	}
	response = NewDescribeDedicatedClusterCosCapacityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterHostsRequest() (request *DescribeDedicatedClusterHostsRequest) {
	request = &DescribeDedicatedClusterHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterHosts")
	return
}

func NewDescribeDedicatedClusterHostsResponse() (response *DescribeDedicatedClusterHostsResponse) {
	response = &DescribeDedicatedClusterHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群宿主机信息
func (c *Client) DescribeDedicatedClusterHosts(request *DescribeDedicatedClusterHostsRequest) (response *DescribeDedicatedClusterHostsResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterHostsRequest()
	}
	response = NewDescribeDedicatedClusterHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterHostsInfoRequest() (request *DescribeDedicatedClusterHostsInfoRequest) {
	request = &DescribeDedicatedClusterHostsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterHostsInfo")
	return
}

func NewDescribeDedicatedClusterHostsInfoResponse() (response *DescribeDedicatedClusterHostsInfoResponse) {
	response = &DescribeDedicatedClusterHostsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群宿主机信息
func (c *Client) DescribeDedicatedClusterHostsInfo(request *DescribeDedicatedClusterHostsInfoRequest) (response *DescribeDedicatedClusterHostsInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterHostsInfoRequest()
	}
	response = NewDescribeDedicatedClusterHostsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterTypesRequest() (request *DescribeDedicatedClusterTypesRequest) {
	request = &DescribeDedicatedClusterTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterTypes")
	return
}

func NewDescribeDedicatedClusterTypesResponse() (response *DescribeDedicatedClusterTypesResponse) {
	response = &DescribeDedicatedClusterTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专有集群配置列表
func (c *Client) DescribeDedicatedClusterTypes(request *DescribeDedicatedClusterTypesRequest) (response *DescribeDedicatedClusterTypesResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterTypesRequest()
	}
	response = NewDescribeDedicatedClusterTypesResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceCreateDedicatedClusterOrderRequest() (request *InquirePriceCreateDedicatedClusterOrderRequest) {
	request = &InquirePriceCreateDedicatedClusterOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "InquirePriceCreateDedicatedClusterOrder")
	return
}

func NewInquirePriceCreateDedicatedClusterOrderResponse() (response *InquirePriceCreateDedicatedClusterOrderResponse) {
	response = &InquirePriceCreateDedicatedClusterOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群订单询价
func (c *Client) InquirePriceCreateDedicatedClusterOrder(request *InquirePriceCreateDedicatedClusterOrderRequest) (response *InquirePriceCreateDedicatedClusterOrderResponse, err error) {
	if request == nil {
		request = NewInquirePriceCreateDedicatedClusterOrderRequest()
	}
	response = NewInquirePriceCreateDedicatedClusterOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedSupportedZonesRequest() (request *DescribeDedicatedSupportedZonesRequest) {
	request = &DescribeDedicatedSupportedZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedSupportedZones")
	return
}

func NewDescribeDedicatedSupportedZonesResponse() (response *DescribeDedicatedSupportedZonesResponse) {
	response = &DescribeDedicatedSupportedZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群支持的可用区列表
func (c *Client) DescribeDedicatedSupportedZones(request *DescribeDedicatedSupportedZonesRequest) (response *DescribeDedicatedSupportedZonesResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedSupportedZonesRequest()
	}
	response = NewDescribeDedicatedSupportedZonesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDedicatedClusterOrderRequest() (request *CreateDedicatedClusterOrderRequest) {
	request = &CreateDedicatedClusterOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "CreateDedicatedClusterOrder")
	return
}

func NewCreateDedicatedClusterOrderResponse() (response *CreateDedicatedClusterOrderResponse) {
	response = &CreateDedicatedClusterOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专用集群订单
func (c *Client) CreateDedicatedClusterOrder(request *CreateDedicatedClusterOrderRequest) (response *CreateDedicatedClusterOrderResponse, err error) {
	if request == nil {
		request = NewCreateDedicatedClusterOrderRequest()
	}
	response = NewCreateDedicatedClusterOrderResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDedicatedClusterRequest() (request *CreateDedicatedClusterRequest) {
	request = &CreateDedicatedClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "CreateDedicatedCluster")
	return
}

func NewCreateDedicatedClusterResponse() (response *CreateDedicatedClusterResponse) {
	response = &CreateDedicatedClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专用集群
func (c *Client) CreateDedicatedCluster(request *CreateDedicatedClusterRequest) (response *CreateDedicatedClusterResponse, err error) {
	if request == nil {
		request = NewCreateDedicatedClusterRequest()
	}
	response = NewCreateDedicatedClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSitesDetailRequest() (request *DescribeSitesDetailRequest) {
	request = &DescribeSitesDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeSitesDetail")
	return
}

func NewDescribeSitesDetailResponse() (response *DescribeSitesDetailResponse) {
	response = &DescribeSitesDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询站点详情
func (c *Client) DescribeSitesDetail(request *DescribeSitesDetailRequest) (response *DescribeSitesDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSitesDetailRequest()
	}
	response = NewDescribeSitesDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifySiteDeviceInfoRequest() (request *ModifySiteDeviceInfoRequest) {
	request = &ModifySiteDeviceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "ModifySiteDeviceInfo")
	return
}

func NewModifySiteDeviceInfoResponse() (response *ModifySiteDeviceInfoResponse) {
	response = &ModifySiteDeviceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机房设备信息
func (c *Client) ModifySiteDeviceInfo(request *ModifySiteDeviceInfoRequest) (response *ModifySiteDeviceInfoResponse, err error) {
	if request == nil {
		request = NewModifySiteDeviceInfoRequest()
	}
	response = NewModifySiteDeviceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterAlarmsRequest() (request *DescribeDedicatedClusterAlarmsRequest) {
	request = &DescribeDedicatedClusterAlarmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterAlarms")
	return
}

func NewDescribeDedicatedClusterAlarmsResponse() (response *DescribeDedicatedClusterAlarmsResponse) {
	response = &DescribeDedicatedClusterAlarmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群告警信息列表
func (c *Client) DescribeDedicatedClusterAlarms(request *DescribeDedicatedClusterAlarmsRequest) (response *DescribeDedicatedClusterAlarmsResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterAlarmsRequest()
	}
	response = NewDescribeDedicatedClusterAlarmsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySiteInfoRequest() (request *ModifySiteInfoRequest) {
	request = &ModifySiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "ModifySiteInfo")
	return
}

func NewModifySiteInfoResponse() (response *ModifySiteInfoResponse) {
	response = &ModifySiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机房信息
func (c *Client) ModifySiteInfo(request *ModifySiteInfoRequest) (response *ModifySiteInfoResponse, err error) {
	if request == nil {
		request = NewModifySiteInfoRequest()
	}
	response = NewModifySiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterCreateDedicatedClusterOrderRequest() (request *SwitchParameterCreateDedicatedClusterOrderRequest) {
	request = &SwitchParameterCreateDedicatedClusterOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "SwitchParameterCreateDedicatedClusterOrder")
	return
}

func NewSwitchParameterCreateDedicatedClusterOrderResponse() (response *SwitchParameterCreateDedicatedClusterOrderResponse) {
	response = &SwitchParameterCreateDedicatedClusterOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群订单参数转换
func (c *Client) SwitchParameterCreateDedicatedClusterOrder(request *SwitchParameterCreateDedicatedClusterOrderRequest) (response *SwitchParameterCreateDedicatedClusterOrderResponse, err error) {
	if request == nil {
		request = NewSwitchParameterCreateDedicatedClusterOrderRequest()
	}
	response = NewSwitchParameterCreateDedicatedClusterOrderResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteRequest() (request *CreateSiteRequest) {
	request = &CreateSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "CreateSite")
	return
}

func NewCreateSiteResponse() (response *CreateSiteResponse) {
	response = &CreateSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建站点
func (c *Client) CreateSite(request *CreateSiteRequest) (response *CreateSiteResponse, err error) {
	if request == nil {
		request = NewCreateSiteRequest()
	}
	response = NewCreateSiteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClustersRequest() (request *DescribeDedicatedClustersRequest) {
	request = &DescribeDedicatedClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusters")
	return
}

func NewDescribeDedicatedClustersResponse() (response *DescribeDedicatedClustersResponse) {
	response = &DescribeDedicatedClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群列表
func (c *Client) DescribeDedicatedClusters(request *DescribeDedicatedClustersRequest) (response *DescribeDedicatedClustersResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClustersRequest()
	}
	response = NewDescribeDedicatedClustersResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceCreateDedicatedClusterUserDefinedOrderRequest() (request *InquirePriceCreateDedicatedClusterUserDefinedOrderRequest) {
	request = &InquirePriceCreateDedicatedClusterUserDefinedOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "InquirePriceCreateDedicatedClusterUserDefinedOrder")
	return
}

func NewInquirePriceCreateDedicatedClusterUserDefinedOrderResponse() (response *InquirePriceCreateDedicatedClusterUserDefinedOrderResponse) {
	response = &InquirePriceCreateDedicatedClusterUserDefinedOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群自定义订单询价
func (c *Client) InquirePriceCreateDedicatedClusterUserDefinedOrder(request *InquirePriceCreateDedicatedClusterUserDefinedOrderRequest) (response *InquirePriceCreateDedicatedClusterUserDefinedOrderResponse, err error) {
	if request == nil {
		request = NewInquirePriceCreateDedicatedClusterUserDefinedOrderRequest()
	}
	response = NewInquirePriceCreateDedicatedClusterUserDefinedOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterHostStatisticsRequest() (request *DescribeDedicatedClusterHostStatisticsRequest) {
	request = &DescribeDedicatedClusterHostStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterHostStatistics")
	return
}

func NewDescribeDedicatedClusterHostStatisticsResponse() (response *DescribeDedicatedClusterHostStatisticsResponse) {
	response = &DescribeDedicatedClusterHostStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群内宿主机的统计信息
func (c *Client) DescribeDedicatedClusterHostStatistics(request *DescribeDedicatedClusterHostStatisticsRequest) (response *DescribeDedicatedClusterHostStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterHostStatisticsRequest()
	}
	response = NewDescribeDedicatedClusterHostStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDedicatedClusterUserDefinedOrderRequest() (request *CreateDedicatedClusterUserDefinedOrderRequest) {
	request = &CreateDedicatedClusterUserDefinedOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "CreateDedicatedClusterUserDefinedOrder")
	return
}

func NewCreateDedicatedClusterUserDefinedOrderResponse() (response *CreateDedicatedClusterUserDefinedOrderResponse) {
	response = &CreateDedicatedClusterUserDefinedOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专用集群自定义订单
func (c *Client) CreateDedicatedClusterUserDefinedOrder(request *CreateDedicatedClusterUserDefinedOrderRequest) (response *CreateDedicatedClusterUserDefinedOrderResponse, err error) {
	if request == nil {
		request = NewCreateDedicatedClusterUserDefinedOrderRequest()
	}
	response = NewCreateDedicatedClusterUserDefinedOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterOrdersRequest() (request *DescribeDedicatedClusterOrdersRequest) {
	request = &DescribeDedicatedClusterOrdersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterOrders")
	return
}

func NewDescribeDedicatedClusterOrdersResponse() (response *DescribeDedicatedClusterOrdersResponse) {
	response = &DescribeDedicatedClusterOrdersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专用集群订单列表
func (c *Client) DescribeDedicatedClusterOrders(request *DescribeDedicatedClusterOrdersRequest) (response *DescribeDedicatedClusterOrdersResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterOrdersRequest()
	}
	response = NewDescribeDedicatedClusterOrdersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDedicatedClusterInfoRequest() (request *ModifyDedicatedClusterInfoRequest) {
	request = &ModifyDedicatedClusterInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "ModifyDedicatedClusterInfo")
	return
}

func NewModifyDedicatedClusterInfoResponse() (response *ModifyDedicatedClusterInfoResponse) {
	response = &ModifyDedicatedClusterInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改本地专用集群信息
func (c *Client) ModifyDedicatedClusterInfo(request *ModifyDedicatedClusterInfoRequest) (response *ModifyDedicatedClusterInfoResponse, err error) {
	if request == nil {
		request = NewModifyDedicatedClusterInfoRequest()
	}
	response = NewModifyDedicatedClusterInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDedicatedClustersRequest() (request *DeleteDedicatedClustersRequest) {
	request = &DeleteDedicatedClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DeleteDedicatedClusters")
	return
}

func NewDeleteDedicatedClustersResponse() (response *DeleteDedicatedClustersResponse) {
	response = &DeleteDedicatedClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专用集群
func (c *Client) DeleteDedicatedClusters(request *DeleteDedicatedClustersRequest) (response *DeleteDedicatedClustersResponse, err error) {
	if request == nil {
		request = NewDeleteDedicatedClustersRequest()
	}
	response = NewDeleteDedicatedClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedClusterOverviewRequest() (request *DescribeDedicatedClusterOverviewRequest) {
	request = &DescribeDedicatedClusterOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterOverview")
	return
}

func NewDescribeDedicatedClusterOverviewResponse() (response *DescribeDedicatedClusterOverviewResponse) {
	response = &DescribeDedicatedClusterOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专用集群概览信息
func (c *Client) DescribeDedicatedClusterOverview(request *DescribeDedicatedClusterOverviewRequest) (response *DescribeDedicatedClusterOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedClusterOverviewRequest()
	}
	response = NewDescribeDedicatedClusterOverviewResponse()
	err = c.Send(request, response)
	return
}

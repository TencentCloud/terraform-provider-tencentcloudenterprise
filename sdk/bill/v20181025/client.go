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

package v20181025

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-10-25"

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

func NewDeletePartaMailingAddressGatewayRequest() (request *DeletePartaMailingAddressGatewayRequest) {
	request = &DeletePartaMailingAddressGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeletePartaMailingAddressGateway")
	return
}

func NewDeletePartaMailingAddressGatewayResponse() (response *DeletePartaMailingAddressGatewayResponse) {
	response = &DeletePartaMailingAddressGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除邮寄地址
func (c *Client) DeletePartaMailingAddressGateway(request *DeletePartaMailingAddressGatewayRequest) (response *DeletePartaMailingAddressGatewayResponse, err error) {
	if request == nil {
		request = NewDeletePartaMailingAddressGatewayRequest()
	}
	response = NewDeletePartaMailingAddressGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceSummaryGatewayRequest() (request *ListResourceSummaryGatewayRequest) {
	request = &ListResourceSummaryGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceSummaryGateway")
	return
}

func NewListResourceSummaryGatewayResponse() (response *ListResourceSummaryGatewayResponse) {
	response = &ListResourceSummaryGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResourceSummaryGateway(request *ListResourceSummaryGatewayRequest) (response *ListResourceSummaryGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceSummaryGatewayRequest()
	}
	response = NewListResourceSummaryGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagsByMonthGatewayRequest() (request *GetTagsByMonthGatewayRequest) {
	request = &GetTagsByMonthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetTagsByMonthGateway")
	return
}

func NewGetTagsByMonthGatewayResponse() (response *GetTagsByMonthGatewayResponse) {
	response = &GetTagsByMonthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某月设置的标签键列表
func (c *Client) GetTagsByMonthGateway(request *GetTagsByMonthGatewayRequest) (response *GetTagsByMonthGatewayResponse, err error) {
	if request == nil {
		request = NewGetTagsByMonthGatewayRequest()
	}
	response = NewGetTagsByMonthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeListGatewayRequest() (request *DescribeContractTypeListGatewayRequest) {
	request = &DescribeContractTypeListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractTypeListGateway")
	return
}

func NewDescribeContractTypeListGatewayResponse() (response *DescribeContractTypeListGatewayResponse) {
	response = &DescribeContractTypeListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取合同类型接口
func (c *Client) DescribeContractTypeListGateway(request *DescribeContractTypeListGatewayRequest) (response *DescribeContractTypeListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeListGatewayRequest()
	}
	response = NewDescribeContractTypeListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewPrepayDealRequest() (request *PrepayDealRequest) {
	request = &PrepayDealRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "PrepayDeal")
	return
}

func NewPrepayDealResponse() (response *PrepayDealResponse) {
	response = &PrepayDealResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源支付
func (c *Client) PrepayDeal(request *PrepayDealRequest) (response *PrepayDealResponse, err error) {
	if request == nil {
		request = NewPrepayDealRequest()
	}
	response = NewPrepayDealResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradePriceRequest() (request *UpgradePriceRequest) {
	request = &UpgradePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "UpgradePrice")
	return
}

func NewUpgradePriceResponse() (response *UpgradePriceResponse) {
	response = &UpgradePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口无效，请勿调用，管理员请及时下线
func (c *Client) UpgradePrice(request *UpgradePriceRequest) (response *UpgradePriceResponse, err error) {
	if request == nil {
		request = NewUpgradePriceRequest()
	}
	response = NewUpgradePriceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefaultPostInfoGatewayRequest() (request *ModifyDefaultPostInfoGatewayRequest) {
	request = &ModifyDefaultPostInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultPostInfoGateway")
	return
}

func NewModifyDefaultPostInfoGatewayResponse() (response *ModifyDefaultPostInfoGatewayResponse) {
	response = &ModifyDefaultPostInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置默认用户的邮寄信息
func (c *Client) ModifyDefaultPostInfoGateway(request *ModifyDefaultPostInfoGatewayRequest) (response *ModifyDefaultPostInfoGatewayResponse, err error) {
	if request == nil {
		request = NewModifyDefaultPostInfoGatewayRequest()
	}
	response = NewModifyDefaultPostInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaLeftGatewayRequest() (request *GetQuotaLeftGatewayRequest) {
	request = &GetQuotaLeftGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetQuotaLeftGateway")
	return
}

func NewGetQuotaLeftGatewayResponse() (response *GetQuotaLeftGatewayResponse) {
	response = &GetQuotaLeftGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取我的剩余配额
func (c *Client) GetQuotaLeftGateway(request *GetQuotaLeftGatewayRequest) (response *GetQuotaLeftGatewayResponse, err error) {
	if request == nil {
		request = NewGetQuotaLeftGatewayRequest()
	}
	response = NewGetQuotaLeftGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryQuotaRequest() (request *QueryQuotaRequest) {
	request = &QueryQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "QueryQuota")
	return
}

func NewQueryQuotaResponse() (response *QueryQuotaResponse) {
	response = &QueryQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户当前使用配额
func (c *Client) QueryQuota(request *QueryQuotaRequest) (response *QueryQuotaResponse, err error) {
	if request == nil {
		request = NewQueryQuotaRequest()
	}
	response = NewQueryQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetBillingProductCodeGatewayRequest() (request *GetBillingProductCodeGatewayRequest) {
	request = &GetBillingProductCodeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetBillingProductCodeGateway")
	return
}

func NewGetBillingProductCodeGatewayResponse() (response *GetBillingProductCodeGatewayResponse) {
	response = &GetBillingProductCodeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询由计量迁移至计费的产品码
func (c *Client) GetBillingProductCodeGateway(request *GetBillingProductCodeGatewayRequest) (response *GetBillingProductCodeGatewayResponse, err error) {
	if request == nil {
		request = NewGetBillingProductCodeGatewayRequest()
	}
	response = NewGetBillingProductCodeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContractGatewayRequest() (request *CreateContractGatewayRequest) {
	request = &CreateContractGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateContractGateway")
	return
}

func NewCreateContractGatewayResponse() (response *CreateContractGatewayResponse) {
	response = &CreateContractGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建合同
func (c *Client) CreateContractGateway(request *CreateContractGatewayRequest) (response *CreateContractGatewayResponse, err error) {
	if request == nil {
		request = NewCreateContractGatewayRequest()
	}
	response = NewCreateContractGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePartaMailingAddressRequest() (request *DeletePartaMailingAddressRequest) {
	request = &DeletePartaMailingAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeletePartaMailingAddress")
	return
}

func NewDeletePartaMailingAddressResponse() (response *DeletePartaMailingAddressResponse) {
	response = &DeletePartaMailingAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除邮寄地址
func (c *Client) DeletePartaMailingAddress(request *DeletePartaMailingAddressRequest) (response *DeletePartaMailingAddressResponse, err error) {
	if request == nil {
		request = NewDeletePartaMailingAddressRequest()
	}
	response = NewDeletePartaMailingAddressResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceDetailGatewayRequest() (request *ListResourceDetailGatewayRequest) {
	request = &ListResourceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceDetailGateway")
	return
}

func NewListResourceDetailGatewayResponse() (response *ListResourceDetailGatewayResponse) {
	response = &ListResourceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Uin，ResourceId, RegionId, ProductCode查询资源详细信息
func (c *Client) ListResourceDetailGateway(request *ListResourceDetailGatewayRequest) (response *ListResourceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceDetailGatewayRequest()
	}
	response = NewListResourceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewHourUnblockGatewayRequest() (request *HourUnblockGatewayRequest) {
	request = &HourUnblockGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourUnblockGateway")
	return
}

func NewHourUnblockGatewayResponse() (response *HourUnblockGatewayResponse) {
	response = &HourUnblockGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源销毁
func (c *Client) HourUnblockGateway(request *HourUnblockGatewayRequest) (response *HourUnblockGatewayResponse, err error) {
	if request == nil {
		request = NewHourUnblockGatewayRequest()
	}
	response = NewHourUnblockGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGoodsPriceRequest() (request *DescribeGoodsPriceRequest) {
	request = &DescribeGoodsPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeGoodsPrice")
	return
}

func NewDescribeGoodsPriceResponse() (response *DescribeGoodsPriceResponse) {
	response = &DescribeGoodsPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询商品价格
func (c *Client) DescribeGoodsPrice(request *DescribeGoodsPriceRequest) (response *DescribeGoodsPriceResponse, err error) {
	if request == nil {
		request = NewDescribeGoodsPriceRequest()
	}
	response = NewDescribeGoodsPriceResponse()
	err = c.Send(request, response)
	return
}

func NewAddProductRequest() (request *AddProductRequest) {
	request = &AddProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddProduct")
	return
}

func NewAddProductResponse() (response *AddProductResponse) {
	response = &AddProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加产品四层定义
func (c *Client) AddProduct(request *AddProductRequest) (response *AddProductResponse, err error) {
	if request == nil {
		request = NewAddProductRequest()
	}
	response = NewAddProductResponse()
	err = c.Send(request, response)
	return
}

func NewCouponsWaterRequest() (request *CouponsWaterRequest) {
	request = &CouponsWaterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CouponsWater")
	return
}

func NewCouponsWaterResponse() (response *CouponsWaterResponse) {
	response = &CouponsWaterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 代金券流水
func (c *Client) CouponsWater(request *CouponsWaterRequest) (response *CouponsWaterResponse, err error) {
	if request == nil {
		request = NewCouponsWaterRequest()
	}
	response = NewCouponsWaterResponse()
	err = c.Send(request, response)
	return
}

func NewGetMineQuotaGatewayRequest() (request *GetMineQuotaGatewayRequest) {
	request = &GetMineQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMineQuotaGateway")
	return
}

func NewGetMineQuotaGatewayResponse() (response *GetMineQuotaGatewayResponse) {
	response = &GetMineQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询用户自身的配额
func (c *Client) GetMineQuotaGateway(request *GetMineQuotaGatewayRequest) (response *GetMineQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewGetMineQuotaGatewayRequest()
	}
	response = NewGetMineQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewHourCreateRequest() (request *HourCreateRequest) {
	request = &HourCreateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourCreate")
	return
}

func NewHourCreateResponse() (response *HourCreateResponse) {
	response = &HourCreateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源下单开通
func (c *Client) HourCreate(request *HourCreateRequest) (response *HourCreateResponse, err error) {
	if request == nil {
		request = NewHourCreateRequest()
	}
	response = NewHourCreateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPostInfoRequest() (request *ModifyPostInfoRequest) {
	request = &ModifyPostInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyPostInfo")
	return
}

func NewModifyPostInfoResponse() (response *ModifyPostInfoResponse) {
	response = &ModifyPostInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改新增用户的邮寄信息
func (c *Client) ModifyPostInfo(request *ModifyPostInfoRequest) (response *ModifyPostInfoResponse, err error) {
	if request == nil {
		request = NewModifyPostInfoRequest()
	}
	response = NewModifyPostInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceSummaryDownloadGatewayRequest() (request *ListResourceSummaryDownloadGatewayRequest) {
	request = &ListResourceSummaryDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceSummaryDownloadGateway")
	return
}

func NewListResourceSummaryDownloadGatewayResponse() (response *ListResourceSummaryDownloadGatewayResponse) {
	response = &ListResourceSummaryDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResourceSummaryDownloadGateway(request *ListResourceSummaryDownloadGatewayRequest) (response *ListResourceSummaryDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceSummaryDownloadGatewayRequest()
	}
	response = NewListResourceSummaryDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadCommonSummaryRequest() (request *DescribeBillDownloadCommonSummaryRequest) {
	request = &DescribeBillDownloadCommonSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadCommonSummary")
	return
}

func NewDescribeBillDownloadCommonSummaryResponse() (response *DescribeBillDownloadCommonSummaryResponse) {
	response = &DescribeBillDownloadCommonSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载账单通用汇总（产品+项目+地域）
func (c *Client) DescribeBillDownloadCommonSummary(request *DescribeBillDownloadCommonSummaryRequest) (response *DescribeBillDownloadCommonSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadCommonSummaryRequest()
	}
	response = NewDescribeBillDownloadCommonSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRenewInfoRequest() (request *DescribeRenewInfoRequest) {
	request = &DescribeRenewInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeRenewInfo")
	return
}

func NewDescribeRenewInfoResponse() (response *DescribeRenewInfoResponse) {
	response = &DescribeRenewInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询续费信息。用户id通过云api公共字段转入
func (c *Client) DescribeRenewInfo(request *DescribeRenewInfoRequest) (response *DescribeRenewInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRenewInfoRequest()
	}
	response = NewDescribeRenewInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByTagRequest() (request *DescribeBillFeeByTagRequest) {
	request = &DescribeBillFeeByTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByTag")
	return
}

func NewDescribeBillFeeByTagResponse() (response *DescribeBillFeeByTagResponse) {
	response = &DescribeBillFeeByTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询按标签汇总明细数据
func (c *Client) DescribeBillFeeByTag(request *DescribeBillFeeByTagRequest) (response *DescribeBillFeeByTagResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByTagRequest()
	}
	response = NewDescribeBillFeeByTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealListRequest() (request *DescribeDealListRequest) {
	request = &DescribeDealListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDealList")
	return
}

func NewDescribeDealListResponse() (response *DescribeDealListResponse) {
	response = &DescribeDealListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单列表
func (c *Client) DescribeDealList(request *DescribeDealListRequest) (response *DescribeDealListResponse, err error) {
	if request == nil {
		request = NewDescribeDealListRequest()
	}
	response = NewDescribeDealListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeTrendRequest() (request *DescribeBillFeeTrendRequest) {
	request = &DescribeBillFeeTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeTrend")
	return
}

func NewDescribeBillFeeTrendResponse() (response *DescribeBillFeeTrendResponse) {
	response = &DescribeBillFeeTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类汇总明细费用趋势
func (c *Client) DescribeBillFeeTrend(request *DescribeBillFeeTrendRequest) (response *DescribeBillFeeTrendResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeTrendRequest()
	}
	response = NewDescribeBillFeeTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantQuotasGatewayRequest() (request *DescribeTenantQuotasGatewayRequest) {
	request = &DescribeTenantQuotasGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeTenantQuotasGateway")
	return
}

func NewDescribeTenantQuotasGatewayResponse() (response *DescribeTenantQuotasGatewayResponse) {
	response = &DescribeTenantQuotasGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端产品配额列表查询
func (c *Client) DescribeTenantQuotasGateway(request *DescribeTenantQuotasGatewayRequest) (response *DescribeTenantQuotasGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeTenantQuotasGatewayRequest()
	}
	response = NewDescribeTenantQuotasGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceRequest() (request *ListResourceRequest) {
	request = &ListResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResource")
	return
}

func NewListResourceResponse() (response *ListResourceResponse) {
	response = &ListResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResource(request *ListResourceRequest) (response *ListResourceResponse, err error) {
	if request == nil {
		request = NewListResourceRequest()
	}
	response = NewListResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGoodsPriceGatewayRequest() (request *DescribeGoodsPriceGatewayRequest) {
	request = &DescribeGoodsPriceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeGoodsPriceGateway")
	return
}

func NewDescribeGoodsPriceGatewayResponse() (response *DescribeGoodsPriceGatewayResponse) {
	response = &DescribeGoodsPriceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询商品价格
func (c *Client) DescribeGoodsPriceGateway(request *DescribeGoodsPriceGatewayRequest) (response *DescribeGoodsPriceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeGoodsPriceGatewayRequest()
	}
	response = NewDescribeGoodsPriceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaRequest() (request *GetQuotaRequest) {
	request = &GetQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetQuota")
	return
}

func NewGetQuotaResponse() (response *GetQuotaResponse) {
	response = &GetQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询配额
func (c *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
	if request == nil {
		request = NewGetQuotaRequest()
	}
	response = NewGetQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescriPartaMailingAddressListGatewayRequest() (request *DescriPartaMailingAddressListGatewayRequest) {
	request = &DescriPartaMailingAddressListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescriPartaMailingAddressListGateway")
	return
}

func NewDescriPartaMailingAddressListGatewayResponse() (response *DescriPartaMailingAddressListGatewayResponse) {
	response = &DescriPartaMailingAddressListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询邮寄地址列表
func (c *Client) DescriPartaMailingAddressListGateway(request *DescriPartaMailingAddressListGatewayRequest) (response *DescriPartaMailingAddressListGatewayResponse, err error) {
	if request == nil {
		request = NewDescriPartaMailingAddressListGatewayRequest()
	}
	response = NewDescriPartaMailingAddressListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSubscribeManualRenewRequest() (request *SubscribeManualRenewRequest) {
	request = &SubscribeManualRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SubscribeManualRenew")
	return
}

func NewSubscribeManualRenewResponse() (response *SubscribeManualRenewResponse) {
	response = &SubscribeManualRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费手工续费下单
func (c *Client) SubscribeManualRenew(request *SubscribeManualRenewRequest) (response *SubscribeManualRenewResponse, err error) {
	if request == nil {
		request = NewSubscribeManualRenewRequest()
	}
	response = NewSubscribeManualRenewResponse()
	err = c.Send(request, response)
	return
}

func NewHourCreateGatewayRequest() (request *HourCreateGatewayRequest) {
	request = &HourCreateGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourCreateGateway")
	return
}

func NewHourCreateGatewayResponse() (response *HourCreateGatewayResponse) {
	response = &HourCreateGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源下单开通
func (c *Client) HourCreateGateway(request *HourCreateGatewayRequest) (response *HourCreateGatewayResponse, err error) {
	if request == nil {
		request = NewHourCreateGatewayRequest()
	}
	response = NewHourCreateGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountPeriodGatewayRequest() (request *DescribeAccountPeriodGatewayRequest) {
	request = &DescribeAccountPeriodGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeAccountPeriodGateway")
	return
}

func NewDescribeAccountPeriodGatewayResponse() (response *DescribeAccountPeriodGatewayResponse) {
	response = &DescribeAccountPeriodGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账期信息
func (c *Client) DescribeAccountPeriodGateway(request *DescribeAccountPeriodGatewayRequest) (response *DescribeAccountPeriodGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeAccountPeriodGatewayRequest()
	}
	response = NewDescribeAccountPeriodGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListRequest() (request *ExportDealListRequest) {
	request = &ExportDealListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ExportDealList")
	return
}

func NewExportDealListResponse() (response *ExportDealListResponse) {
	response = &ExportDealListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入订单列表信息
func (c *Client) ExportDealList(request *ExportDealListRequest) (response *ExportDealListResponse, err error) {
	if request == nil {
		request = NewExportDealListRequest()
	}
	response = NewExportDealListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsolidatedBillDownloadUrlGatewayRequest() (request *DescribeConsolidatedBillDownloadUrlGatewayRequest) {
	request = &DescribeConsolidatedBillDownloadUrlGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeConsolidatedBillDownloadUrlGateway")
	return
}

func NewDescribeConsolidatedBillDownloadUrlGatewayResponse() (response *DescribeConsolidatedBillDownloadUrlGatewayResponse) {
	response = &DescribeConsolidatedBillDownloadUrlGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrlGateway(request *DescribeConsolidatedBillDownloadUrlGatewayRequest) (response *DescribeConsolidatedBillDownloadUrlGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeConsolidatedBillDownloadUrlGatewayRequest()
	}
	response = NewDescribeConsolidatedBillDownloadUrlGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetBillingProductCodeRequest() (request *GetBillingProductCodeRequest) {
	request = &GetBillingProductCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetBillingProductCode")
	return
}

func NewGetBillingProductCodeResponse() (response *GetBillingProductCodeResponse) {
	response = &GetBillingProductCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询由计量迁移至计费的产品码
func (c *Client) GetBillingProductCode(request *GetBillingProductCodeRequest) (response *GetBillingProductCodeResponse, err error) {
	if request == nil {
		request = NewGetBillingProductCodeRequest()
	}
	response = NewGetBillingProductCodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubUinQuotaGatewayRequest() (request *DeleteSubUinQuotaGatewayRequest) {
	request = &DeleteSubUinQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteSubUinQuotaGateway")
	return
}

func NewDeleteSubUinQuotaGatewayResponse() (response *DeleteSubUinQuotaGatewayResponse) {
	response = &DeleteSubUinQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除父账户下子账户的配额数据
func (c *Client) DeleteSubUinQuotaGateway(request *DeleteSubUinQuotaGatewayRequest) (response *DeleteSubUinQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteSubUinQuotaGatewayRequest()
	}
	response = NewDeleteSubUinQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePropertiesRequest() (request *DescribePropertiesRequest) {
	request = &DescribePropertiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProperties")
	return
}

func NewDescribePropertiesResponse() (response *DescribePropertiesResponse) {
	response = &DescribePropertiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询售卖属性列表
func (c *Client) DescribeProperties(request *DescribePropertiesRequest) (response *DescribePropertiesResponse, err error) {
	if request == nil {
		request = NewDescribePropertiesRequest()
	}
	response = NewDescribePropertiesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceListRequest() (request *DescribeInvoiceListRequest) {
	request = &DescribeInvoiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceList")
	return
}

func NewDescribeInvoiceListResponse() (response *DescribeInvoiceListResponse) {
	response = &DescribeInvoiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取用户开票记录
func (c *Client) DescribeInvoiceList(request *DescribeInvoiceListRequest) (response *DescribeInvoiceListResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceListRequest()
	}
	response = NewDescribeInvoiceListResponse()
	err = c.Send(request, response)
	return
}

func NewGetMineQuotaRequest() (request *GetMineQuotaRequest) {
	request = &GetMineQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMineQuota")
	return
}

func NewGetMineQuotaResponse() (response *GetMineQuotaResponse) {
	response = &GetMineQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询用户自身的配额
func (c *Client) GetMineQuota(request *GetMineQuotaRequest) (response *GetMineQuotaResponse, err error) {
	if request == nil {
		request = NewGetMineQuotaRequest()
	}
	response = NewGetMineQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadByResourceGatewayRequest() (request *DescribeBillDownloadByResourceGatewayRequest) {
	request = &DescribeBillDownloadByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadByResourceGateway")
	return
}

func NewDescribeBillDownloadByResourceGatewayResponse() (response *DescribeBillDownloadByResourceGatewayResponse) {
	response = &DescribeBillDownloadByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单按资源汇总
func (c *Client) DescribeBillDownloadByResourceGateway(request *DescribeBillDownloadByResourceGatewayRequest) (response *DescribeBillDownloadByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadByResourceGatewayRequest()
	}
	response = NewDescribeBillDownloadByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByRegionGatewayRequest() (request *DescribeBillSummaryByRegionGatewayRequest) {
	request = &DescribeBillSummaryByRegionGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByRegionGateway")
	return
}

func NewDescribeBillSummaryByRegionGatewayResponse() (response *DescribeBillSummaryByRegionGatewayResponse) {
	response = &DescribeBillSummaryByRegionGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegionGateway(request *DescribeBillSummaryByRegionGatewayRequest) (response *DescribeBillSummaryByRegionGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByRegionGatewayRequest()
	}
	response = NewDescribeBillSummaryByRegionGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadRecordListRequest() (request *DescribeBillDownloadRecordListRequest) {
	request = &DescribeBillDownloadRecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadRecordList")
	return
}

func NewDescribeBillDownloadRecordListResponse() (response *DescribeBillDownloadRecordListResponse) {
	response = &DescribeBillDownloadRecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBillDownloadRecordList(request *DescribeBillDownloadRecordListRequest) (response *DescribeBillDownloadRecordListResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadRecordListRequest()
	}
	response = NewDescribeBillDownloadRecordListResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePostInfoRequest() (request *DeletePostInfoRequest) {
	request = &DeletePostInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeletePostInfo")
	return
}

func NewDeletePostInfoResponse() (response *DeletePostInfoResponse) {
	response = &DeletePostInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户的邮寄信息
func (c *Client) DeletePostInfo(request *DeletePostInfoRequest) (response *DeletePostInfoResponse, err error) {
	if request == nil {
		request = NewDeletePostInfoRequest()
	}
	response = NewDeletePostInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractListGatewayRequest() (request *DescribeContractListGatewayRequest) {
	request = &DescribeContractListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractListGateway")
	return
}

func NewDescribeContractListGatewayResponse() (response *DescribeContractListGatewayResponse) {
	response = &DescribeContractListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同列表接口
func (c *Client) DescribeContractListGateway(request *DescribeContractListGatewayRequest) (response *DescribeContractListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractListGatewayRequest()
	}
	response = NewDescribeContractListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDownloadRecordRequest() (request *CreateDownloadRecordRequest) {
	request = &CreateDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateDownloadRecord")
	return
}

func NewCreateDownloadRecordResponse() (response *CreateDownloadRecordResponse) {
	response = &CreateDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户点击下载时通知后台
func (c *Client) CreateDownloadRecord(request *CreateDownloadRecordRequest) (response *CreateDownloadRecordResponse, err error) {
	if request == nil {
		request = NewCreateDownloadRecordRequest()
	}
	response = NewCreateDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadListRequest() (request *DescribeBillDownloadListRequest) {
	request = &DescribeBillDownloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadList")
	return
}

func NewDescribeBillDownloadListResponse() (response *DescribeBillDownloadListResponse) {
	response = &DescribeBillDownloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单下载地址列表
func (c *Client) DescribeBillDownloadList(request *DescribeBillDownloadListRequest) (response *DescribeBillDownloadListResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadListRequest()
	}
	response = NewDescribeBillDownloadListResponse()
	err = c.Send(request, response)
	return
}

func NewCancelTagGatewayRequest() (request *CancelTagGatewayRequest) {
	request = &CancelTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CancelTagGateway")
	return
}

func NewCancelTagGatewayResponse() (response *CancelTagGatewayResponse) {
	response = &CancelTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消标签
func (c *Client) CancelTagGateway(request *CancelTagGatewayRequest) (response *CancelTagGatewayResponse, err error) {
	if request == nil {
		request = NewCancelTagGatewayRequest()
	}
	response = NewCancelTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadSubQuotasListRequest() (request *DescribeDownloadSubQuotasListRequest) {
	request = &DescribeDownloadSubQuotasListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadSubQuotasList")
	return
}

func NewDescribeDownloadSubQuotasListResponse() (response *DescribeDownloadSubQuotasListResponse) {
	response = &DescribeDownloadSubQuotasListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端子账号配额导出接口
func (c *Client) DescribeDownloadSubQuotasList(request *DescribeDownloadSubQuotasListRequest) (response *DescribeDownloadSubQuotasListResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadSubQuotasListRequest()
	}
	response = NewDescribeDownloadSubQuotasListResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceGatewayRequest() (request *ListResourceGatewayRequest) {
	request = &ListResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceGateway")
	return
}

func NewListResourceGatewayResponse() (response *ListResourceGatewayResponse) {
	response = &ListResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResourceGateway(request *ListResourceGatewayRequest) (response *ListResourceGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceGatewayRequest()
	}
	response = NewListResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePayOrderRequest() (request *CreatePayOrderRequest) {
	request = &CreatePayOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreatePayOrder")
	return
}

func NewCreatePayOrderResponse() (response *CreatePayOrderResponse) {
	response = &CreatePayOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建现金支付订单
func (c *Client) CreatePayOrder(request *CreatePayOrderRequest) (response *CreatePayOrderResponse, err error) {
	if request == nil {
		request = NewCreatePayOrderRequest()
	}
	response = NewCreatePayOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDebtBillRequest() (request *DescribeUserDebtBillRequest) {
	request = &DescribeUserDebtBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeUserDebtBill")
	return
}

func NewDescribeUserDebtBillResponse() (response *DescribeUserDebtBillResponse) {
	response = &DescribeUserDebtBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询后付费账单
func (c *Client) DescribeUserDebtBill(request *DescribeUserDebtBillRequest) (response *DescribeUserDebtBillResponse, err error) {
	if request == nil {
		request = NewDescribeUserDebtBillRequest()
	}
	response = NewDescribeUserDebtBillResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDownloadRequest() (request *DescribeResourceDownloadRequest) {
	request = &DescribeResourceDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceDownload")
	return
}

func NewDescribeResourceDownloadResponse() (response *DescribeResourceDownloadResponse) {
	response = &DescribeResourceDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订购实例导出
func (c *Client) DescribeResourceDownload(request *DescribeResourceDownloadRequest) (response *DescribeResourceDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDownloadRequest()
	}
	response = NewDescribeResourceDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDetailByResourceGatewayRequest() (request *DescribeBillDetailByResourceGatewayRequest) {
	request = &DescribeBillDetailByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDetailByResourceGateway")
	return
}

func NewDescribeBillDetailByResourceGatewayResponse() (response *DescribeBillDetailByResourceGatewayResponse) {
	response = &DescribeBillDetailByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源花费详情
func (c *Client) DescribeBillDetailByResourceGateway(request *DescribeBillDetailByResourceGatewayRequest) (response *DescribeBillDetailByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDetailByResourceGatewayRequest()
	}
	response = NewDescribeBillDetailByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractPartaInfoRequest() (request *DescribeContractPartaInfoRequest) {
	request = &DescribeContractPartaInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractPartaInfo")
	return
}

func NewDescribeContractPartaInfoResponse() (response *DescribeContractPartaInfoResponse) {
	response = &DescribeContractPartaInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同甲方信息
func (c *Client) DescribeContractPartaInfo(request *DescribeContractPartaInfoRequest) (response *DescribeContractPartaInfoResponse, err error) {
	if request == nil {
		request = NewDescribeContractPartaInfoRequest()
	}
	response = NewDescribeContractPartaInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSetSubUinQuotaGatewayRequest() (request *SetSubUinQuotaGatewayRequest) {
	request = &SetSubUinQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SetSubUinQuotaGateway")
	return
}

func NewSetSubUinQuotaGatewayResponse() (response *SetSubUinQuotaGatewayResponse) {
	response = &SetSubUinQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置子账号配额
func (c *Client) SetSubUinQuotaGateway(request *SetSubUinQuotaGatewayRequest) (response *SetSubUinQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewSetSubUinQuotaGatewayRequest()
	}
	response = NewSetSubUinQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningGatewayRequest() (request *DescribeWarningGatewayRequest) {
	request = &DescribeWarningGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeWarningGateway")
	return
}

func NewDescribeWarningGatewayResponse() (response *DescribeWarningGatewayResponse) {
	response = &DescribeWarningGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户查询预警状态和阈值
func (c *Client) DescribeWarningGateway(request *DescribeWarningGatewayRequest) (response *DescribeWarningGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeWarningGatewayRequest()
	}
	response = NewDescribeWarningGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractPartaInfoGatewayRequest() (request *DescribeContractPartaInfoGatewayRequest) {
	request = &DescribeContractPartaInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractPartaInfoGateway")
	return
}

func NewDescribeContractPartaInfoGatewayResponse() (response *DescribeContractPartaInfoGatewayResponse) {
	response = &DescribeContractPartaInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同甲方信息
func (c *Client) DescribeContractPartaInfoGateway(request *DescribeContractPartaInfoGatewayRequest) (response *DescribeContractPartaInfoGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractPartaInfoGatewayRequest()
	}
	response = NewDescribeContractPartaInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaLeftRequest() (request *GetQuotaLeftRequest) {
	request = &GetQuotaLeftRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetQuotaLeft")
	return
}

func NewGetQuotaLeftResponse() (response *GetQuotaLeftResponse) {
	response = &GetQuotaLeftResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取我的剩余配额
func (c *Client) GetQuotaLeft(request *GetQuotaLeftRequest) (response *GetQuotaLeftResponse, err error) {
	if request == nil {
		request = NewGetQuotaLeftRequest()
	}
	response = NewGetQuotaLeftResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByRegionRequest() (request *DescribeBillFeeByRegionRequest) {
	request = &DescribeBillFeeByRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByRegion")
	return
}

func NewDescribeBillFeeByRegionResponse() (response *DescribeBillFeeByRegionResponse) {
	response = &DescribeBillFeeByRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总明细费用
func (c *Client) DescribeBillFeeByRegion(request *DescribeBillFeeByRegionRequest) (response *DescribeBillFeeByRegionResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByRegionRequest()
	}
	response = NewDescribeBillFeeByRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoFlagGatewayRequest() (request *ModifyAutoFlagGatewayRequest) {
	request = &ModifyAutoFlagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyAutoFlagGateway")
	return
}

func NewModifyAutoFlagGatewayResponse() (response *ModifyAutoFlagGatewayResponse) {
	response = &ModifyAutoFlagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置续费类型。用户id通过云api公共参数传入
func (c *Client) ModifyAutoFlagGateway(request *ModifyAutoFlagGatewayRequest) (response *ModifyAutoFlagGatewayResponse, err error) {
	if request == nil {
		request = NewModifyAutoFlagGatewayRequest()
	}
	response = NewModifyAutoFlagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQuotaGatewayRequest() (request *DeleteQuotaGatewayRequest) {
	request = &DeleteQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteQuotaGateway")
	return
}

func NewDeleteQuotaGatewayResponse() (response *DeleteQuotaGatewayResponse) {
	response = &DeleteQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口描述
func (c *Client) DeleteQuotaGateway(request *DeleteQuotaGatewayRequest) (response *DeleteQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteQuotaGatewayRequest()
	}
	response = NewDeleteQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaDownloadRequest() (request *DescribeQuotaDownloadRequest) {
	request = &DescribeQuotaDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeQuotaDownload")
	return
}

func NewDescribeQuotaDownloadResponse() (response *DescribeQuotaDownloadResponse) {
	response = &DescribeQuotaDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [计量]配额导出
func (c *Client) DescribeQuotaDownload(request *DescribeQuotaDownloadRequest) (response *DescribeQuotaDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaDownloadRequest()
	}
	response = NewDescribeQuotaDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
	request = &DescribeResourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceList")
	return
}

func NewDescribeResourceListResponse() (response *DescribeResourceListResponse) {
	response = &DescribeResourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户产品列表，只支持到一层产品
func (c *Client) DescribeResourceList(request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
	if request == nil {
		request = NewDescribeResourceListRequest()
	}
	response = NewDescribeResourceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadRecordListGatewayRequest() (request *DescribeBillDownloadRecordListGatewayRequest) {
	request = &DescribeBillDownloadRecordListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadRecordListGateway")
	return
}

func NewDescribeBillDownloadRecordListGatewayResponse() (response *DescribeBillDownloadRecordListGatewayResponse) {
	response = &DescribeBillDownloadRecordListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBillDownloadRecordListGateway(request *DescribeBillDownloadRecordListGatewayRequest) (response *DescribeBillDownloadRecordListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadRecordListGatewayRequest()
	}
	response = NewDescribeBillDownloadRecordListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostInfoListRequest() (request *DescribePostInfoListRequest) {
	request = &DescribePostInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribePostInfoList")
	return
}

func NewDescribePostInfoListResponse() (response *DescribePostInfoListResponse) {
	response = &DescribePostInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的邮寄信息
func (c *Client) DescribePostInfoList(request *DescribePostInfoListRequest) (response *DescribePostInfoListResponse, err error) {
	if request == nil {
		request = NewDescribePostInfoListRequest()
	}
	response = NewDescribePostInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaDownloadGatewayRequest() (request *DescribeQuotaDownloadGatewayRequest) {
	request = &DescribeQuotaDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeQuotaDownloadGateway")
	return
}

func NewDescribeQuotaDownloadGatewayResponse() (response *DescribeQuotaDownloadGatewayResponse) {
	response = &DescribeQuotaDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [计量]配额导出
func (c *Client) DescribeQuotaDownloadGateway(request *DescribeQuotaDownloadGatewayRequest) (response *DescribeQuotaDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaDownloadGatewayRequest()
	}
	response = NewDescribeQuotaDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInvoiceGatewayRequest() (request *CreateInvoiceGatewayRequest) {
	request = &CreateInvoiceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateInvoiceGateway")
	return
}

func NewCreateInvoiceGatewayResponse() (response *CreateInvoiceGatewayResponse) {
	response = &CreateInvoiceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户申请开票
func (c *Client) CreateInvoiceGateway(request *CreateInvoiceGatewayRequest) (response *CreateInvoiceGatewayResponse, err error) {
	if request == nil {
		request = NewCreateInvoiceGatewayRequest()
	}
	response = NewCreateInvoiceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
	request = &DescribeProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProducts")
	return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
	response = &DescribeProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取商品信息(第一层)
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
	if request == nil {
		request = NewDescribeProductsRequest()
	}
	response = NewDescribeProductsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadResourceDetailGatewayRequest() (request *DescribeBillDownloadResourceDetailGatewayRequest) {
	request = &DescribeBillDownloadResourceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadResourceDetailGateway")
	return
}

func NewDescribeBillDownloadResourceDetailGatewayResponse() (response *DescribeBillDownloadResourceDetailGatewayResponse) {
	response = &DescribeBillDownloadResourceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单明细
func (c *Client) DescribeBillDownloadResourceDetailGateway(request *DescribeBillDownloadResourceDetailGatewayRequest) (response *DescribeBillDownloadResourceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadResourceDetailGatewayRequest()
	}
	response = NewDescribeBillDownloadResourceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpenInvoiceInfoRequest() (request *ModifyOpenInvoiceInfoRequest) {
	request = &ModifyOpenInvoiceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyOpenInvoiceInfo")
	return
}

func NewModifyOpenInvoiceInfoResponse() (response *ModifyOpenInvoiceInfoResponse) {
	response = &ModifyOpenInvoiceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置开票信息
func (c *Client) ModifyOpenInvoiceInfo(request *ModifyOpenInvoiceInfoRequest) (response *ModifyOpenInvoiceInfoResponse, err error) {
	if request == nil {
		request = NewModifyOpenInvoiceInfoRequest()
	}
	response = NewModifyOpenInvoiceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceDetailRequest() (request *DescribeInvoiceDetailRequest) {
	request = &DescribeInvoiceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceDetail")
	return
}

func NewDescribeInvoiceDetailResponse() (response *DescribeInvoiceDetailResponse) {
	response = &DescribeInvoiceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取发票的开票详情
func (c *Client) DescribeInvoiceDetail(request *DescribeInvoiceDetailRequest) (response *DescribeInvoiceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceDetailRequest()
	}
	response = NewDescribeInvoiceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewSyncContractPartaInfoGatewayRequest() (request *SyncContractPartaInfoGatewayRequest) {
	request = &SyncContractPartaInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SyncContractPartaInfoGateway")
	return
}

func NewSyncContractPartaInfoGatewayResponse() (response *SyncContractPartaInfoGatewayResponse) {
	response = &SyncContractPartaInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步合同甲方信息
func (c *Client) SyncContractPartaInfoGateway(request *SyncContractPartaInfoGatewayRequest) (response *SyncContractPartaInfoGatewayResponse, err error) {
	if request == nil {
		request = NewSyncContractPartaInfoGatewayRequest()
	}
	response = NewSyncContractPartaInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductDeployStatusGatewayRequest() (request *GetProductDeployStatusGatewayRequest) {
	request = &GetProductDeployStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetProductDeployStatusGateway")
	return
}

func NewGetProductDeployStatusGatewayResponse() (response *GetProductDeployStatusGatewayResponse) {
	response = &GetProductDeployStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品部署计费状态
func (c *Client) GetProductDeployStatusGateway(request *GetProductDeployStatusGatewayRequest) (response *GetProductDeployStatusGatewayResponse, err error) {
	if request == nil {
		request = NewGetProductDeployStatusGatewayRequest()
	}
	response = NewGetProductDeployStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByTagGatewayRequest() (request *DescribeBillSummaryByTagGatewayRequest) {
	request = &DescribeBillSummaryByTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByTagGateway")
	return
}

func NewDescribeBillSummaryByTagGatewayResponse() (response *DescribeBillSummaryByTagGatewayResponse) {
	response = &DescribeBillSummaryByTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签汇总费用分布
func (c *Client) DescribeBillSummaryByTagGateway(request *DescribeBillSummaryByTagGatewayRequest) (response *DescribeBillSummaryByTagGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByTagGatewayRequest()
	}
	response = NewDescribeBillSummaryByTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReconciliationListRequest() (request *DescribeReconciliationListRequest) {
	request = &DescribeReconciliationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeReconciliationList")
	return
}

func NewDescribeReconciliationListResponse() (response *DescribeReconciliationListResponse) {
	response = &DescribeReconciliationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调账记录查询调账列表
func (c *Client) DescribeReconciliationList(request *DescribeReconciliationListRequest) (response *DescribeReconciliationListResponse, err error) {
	if request == nil {
		request = NewDescribeReconciliationListRequest()
	}
	response = NewDescribeReconciliationListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantSubUinQuotasRequest() (request *DescribeTenantSubUinQuotasRequest) {
	request = &DescribeTenantSubUinQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeTenantSubUinQuotas")
	return
}

func NewDescribeTenantSubUinQuotasResponse() (response *DescribeTenantSubUinQuotasResponse) {
	response = &DescribeTenantSubUinQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端-子账号登录后，查询当前账户在主账户下的配额信息
func (c *Client) DescribeTenantSubUinQuotas(request *DescribeTenantSubUinQuotasRequest) (response *DescribeTenantSubUinQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeTenantSubUinQuotasRequest()
	}
	response = NewDescribeTenantSubUinQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithAuthRequest() (request *DescribeDownloadWithAuthRequest) {
	request = &DescribeDownloadWithAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadWithAuth")
	return
}

func NewDescribeDownloadWithAuthResponse() (response *DescribeDownloadWithAuthResponse) {
	response = &DescribeDownloadWithAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithAuth(request *DescribeDownloadWithAuthRequest) (response *DescribeDownloadWithAuthResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithAuthRequest()
	}
	response = NewDescribeDownloadWithAuthResponse()
	err = c.Send(request, response)
	return
}

func NewCancelInvoiceRequest() (request *CancelInvoiceRequest) {
	request = &CancelInvoiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CancelInvoice")
	return
}

func NewCancelInvoiceResponse() (response *CancelInvoiceResponse) {
	response = &CancelInvoiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消开票申请
func (c *Client) CancelInvoice(request *CancelInvoiceRequest) (response *CancelInvoiceResponse, err error) {
	if request == nil {
		request = NewCancelInvoiceRequest()
	}
	response = NewCancelInvoiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByRegionGatewayRequest() (request *DescribeBillFeeByRegionGatewayRequest) {
	request = &DescribeBillFeeByRegionGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByRegionGateway")
	return
}

func NewDescribeBillFeeByRegionGatewayResponse() (response *DescribeBillFeeByRegionGatewayResponse) {
	response = &DescribeBillFeeByRegionGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总明细费用
func (c *Client) DescribeBillFeeByRegionGateway(request *DescribeBillFeeByRegionGatewayRequest) (response *DescribeBillFeeByRegionGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByRegionGatewayRequest()
	}
	response = NewDescribeBillFeeByRegionGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetMineResourceBillGatewayRequest() (request *GetMineResourceBillGatewayRequest) {
	request = &GetMineResourceBillGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMineResourceBillGateway")
	return
}

func NewGetMineResourceBillGatewayResponse() (response *GetMineResourceBillGatewayResponse) {
	response = &GetMineResourceBillGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户自身的账单信息
func (c *Client) GetMineResourceBillGateway(request *GetMineResourceBillGatewayRequest) (response *GetMineResourceBillGatewayResponse, err error) {
	if request == nil {
		request = NewGetMineResourceBillGatewayRequest()
	}
	response = NewGetMineResourceBillGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagsByMonthRequest() (request *GetTagsByMonthRequest) {
	request = &GetTagsByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetTagsByMonth")
	return
}

func NewGetTagsByMonthResponse() (response *GetTagsByMonthResponse) {
	response = &GetTagsByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某月设置的标签键列表
func (c *Client) GetTagsByMonth(request *GetTagsByMonthRequest) (response *GetTagsByMonthResponse, err error) {
	if request == nil {
		request = NewGetTagsByMonthRequest()
	}
	response = NewGetTagsByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadCommonSummaryGatewayRequest() (request *DescribeBillDownloadCommonSummaryGatewayRequest) {
	request = &DescribeBillDownloadCommonSummaryGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadCommonSummaryGateway")
	return
}

func NewDescribeBillDownloadCommonSummaryGatewayResponse() (response *DescribeBillDownloadCommonSummaryGatewayResponse) {
	response = &DescribeBillDownloadCommonSummaryGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载账单通用汇总（产品+项目+地域）
func (c *Client) DescribeBillDownloadCommonSummaryGateway(request *DescribeBillDownloadCommonSummaryGatewayRequest) (response *DescribeBillDownloadCommonSummaryGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadCommonSummaryGatewayRequest()
	}
	response = NewDescribeBillDownloadCommonSummaryGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceAmountRequest() (request *DescribeInvoiceAmountRequest) {
	request = &DescribeInvoiceAmountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceAmount")
	return
}

func NewDescribeInvoiceAmountResponse() (response *DescribeInvoiceAmountResponse) {
	response = &DescribeInvoiceAmountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可开票金额
func (c *Client) DescribeInvoiceAmount(request *DescribeInvoiceAmountRequest) (response *DescribeInvoiceAmountResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceAmountRequest()
	}
	response = NewDescribeInvoiceAmountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractStatusGatewayRequest() (request *ModifyContractStatusGatewayRequest) {
	request = &ModifyContractStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyContractStatusGateway")
	return
}

func NewModifyContractStatusGatewayResponse() (response *ModifyContractStatusGatewayResponse) {
	response = &ModifyContractStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同状态
func (c *Client) ModifyContractStatusGateway(request *ModifyContractStatusGatewayRequest) (response *ModifyContractStatusGatewayResponse, err error) {
	if request == nil {
		request = NewModifyContractStatusGatewayRequest()
	}
	response = NewModifyContractStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWarningGatewayRequest() (request *CreateWarningGatewayRequest) {
	request = &CreateWarningGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateWarningGateway")
	return
}

func NewCreateWarningGatewayResponse() (response *CreateWarningGatewayResponse) {
	response = &CreateWarningGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加或更新预警信息
func (c *Client) CreateWarningGateway(request *CreateWarningGatewayRequest) (response *CreateWarningGatewayResponse, err error) {
	if request == nil {
		request = NewCreateWarningGatewayRequest()
	}
	response = NewCreateWarningGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpenInvoiceInfoRequest() (request *DescribeOpenInvoiceInfoRequest) {
	request = &DescribeOpenInvoiceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeOpenInvoiceInfo")
	return
}

func NewDescribeOpenInvoiceInfoResponse() (response *DescribeOpenInvoiceInfoResponse) {
	response = &DescribeOpenInvoiceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户开票信息
func (c *Client) DescribeOpenInvoiceInfo(request *DescribeOpenInvoiceInfoRequest) (response *DescribeOpenInvoiceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeOpenInvoiceInfoRequest()
	}
	response = NewDescribeOpenInvoiceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescriPartaMailingAddressListRequest() (request *DescriPartaMailingAddressListRequest) {
	request = &DescriPartaMailingAddressListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescriPartaMailingAddressList")
	return
}

func NewDescriPartaMailingAddressListResponse() (response *DescriPartaMailingAddressListResponse) {
	response = &DescriPartaMailingAddressListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询邮寄地址列表
func (c *Client) DescriPartaMailingAddressList(request *DescriPartaMailingAddressListRequest) (response *DescriPartaMailingAddressListResponse, err error) {
	if request == nil {
		request = NewDescriPartaMailingAddressListRequest()
	}
	response = NewDescriPartaMailingAddressListResponse()
	err = c.Send(request, response)
	return
}

func NewSyncContractPartaInfoRequest() (request *SyncContractPartaInfoRequest) {
	request = &SyncContractPartaInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SyncContractPartaInfo")
	return
}

func NewSyncContractPartaInfoResponse() (response *SyncContractPartaInfoResponse) {
	response = &SyncContractPartaInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步合同甲方信息
func (c *Client) SyncContractPartaInfo(request *SyncContractPartaInfoRequest) (response *SyncContractPartaInfoResponse, err error) {
	if request == nil {
		request = NewSyncContractPartaInfoRequest()
	}
	response = NewSyncContractPartaInfoResponse()
	err = c.Send(request, response)
	return
}

func NewAddPartaMailingAddressRequest() (request *AddPartaMailingAddressRequest) {
	request = &AddPartaMailingAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddPartaMailingAddress")
	return
}

func NewAddPartaMailingAddressResponse() (response *AddPartaMailingAddressResponse) {
	response = &AddPartaMailingAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加邮寄地址
func (c *Client) AddPartaMailingAddress(request *AddPartaMailingAddressRequest) (response *AddPartaMailingAddressResponse, err error) {
	if request == nil {
		request = NewAddPartaMailingAddressRequest()
	}
	response = NewAddPartaMailingAddressResponse()
	err = c.Send(request, response)
	return
}

func NewSubscribePayRequest() (request *SubscribePayRequest) {
	request = &SubscribePayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SubscribePay")
	return
}

func NewSubscribePayResponse() (response *SubscribePayResponse) {
	response = &SubscribePayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 普通预付费支付
func (c *Client) SubscribePay(request *SubscribePayRequest) (response *SubscribePayResponse, err error) {
	if request == nil {
		request = NewSubscribePayRequest()
	}
	response = NewSubscribePayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealListGatewayRequest() (request *DescribeDealListGatewayRequest) {
	request = &DescribeDealListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDealListGateway")
	return
}

func NewDescribeDealListGatewayResponse() (response *DescribeDealListGatewayResponse) {
	response = &DescribeDealListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单列表
func (c *Client) DescribeDealListGateway(request *DescribeDealListGatewayRequest) (response *DescribeDealListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDealListGatewayRequest()
	}
	response = NewDescribeDealListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadQuotasListRequest() (request *DescribeDownloadQuotasListRequest) {
	request = &DescribeDownloadQuotasListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadQuotasList")
	return
}

func NewDescribeDownloadQuotasListResponse() (response *DescribeDownloadQuotasListResponse) {
	response = &DescribeDownloadQuotasListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端配额导出接口
func (c *Client) DescribeDownloadQuotasList(request *DescribeDownloadQuotasListRequest) (response *DescribeDownloadQuotasListResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadQuotasListRequest()
	}
	response = NewDescribeDownloadQuotasListResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyRequest() (request *DestroyRequest) {
	request = &DestroyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "Destroy")
	return
}

func NewDestroyResponse() (response *DestroyResponse) {
	response = &DestroyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源销毁
func (c *Client) Destroy(request *DestroyRequest) (response *DestroyResponse, err error) {
	if request == nil {
		request = NewDestroyRequest()
	}
	response = NewDestroyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpenInvoiceInfoGatewayRequest() (request *ModifyOpenInvoiceInfoGatewayRequest) {
	request = &ModifyOpenInvoiceInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyOpenInvoiceInfoGateway")
	return
}

func NewModifyOpenInvoiceInfoGatewayResponse() (response *ModifyOpenInvoiceInfoGatewayResponse) {
	response = &ModifyOpenInvoiceInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置开票信息
func (c *Client) ModifyOpenInvoiceInfoGateway(request *ModifyOpenInvoiceInfoGatewayRequest) (response *ModifyOpenInvoiceInfoGatewayResponse, err error) {
	if request == nil {
		request = NewModifyOpenInvoiceInfoGatewayRequest()
	}
	response = NewModifyOpenInvoiceInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubUinQuotaRequest() (request *DeleteSubUinQuotaRequest) {
	request = &DeleteSubUinQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteSubUinQuota")
	return
}

func NewDeleteSubUinQuotaResponse() (response *DeleteSubUinQuotaResponse) {
	response = &DeleteSubUinQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除父账户下子账户的配额数据
func (c *Client) DeleteSubUinQuota(request *DeleteSubUinQuotaRequest) (response *DeleteSubUinQuotaResponse, err error) {
	if request == nil {
		request = NewDeleteSubUinQuotaRequest()
	}
	response = NewDeleteSubUinQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByProductGatewayRequest() (request *DescribeBillFeeByProductGatewayRequest) {
	request = &DescribeBillFeeByProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByProductGateway")
	return
}

func NewDescribeBillFeeByProductGatewayResponse() (response *DescribeBillFeeByProductGatewayResponse) {
	response = &DescribeBillFeeByProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按产品汇总明细费用
func (c *Client) DescribeBillFeeByProductGateway(request *DescribeBillFeeByProductGatewayRequest) (response *DescribeBillFeeByProductGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByProductGatewayRequest()
	}
	response = NewDescribeBillFeeByProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductGatewayRequest() (request *DeleteProductGatewayRequest) {
	request = &DeleteProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteProductGateway")
	return
}

func NewDeleteProductGatewayResponse() (response *DeleteProductGatewayResponse) {
	response = &DeleteProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】删除产品四层定义
func (c *Client) DeleteProductGateway(request *DeleteProductGatewayRequest) (response *DeleteProductGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteProductGatewayRequest()
	}
	response = NewDeleteProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListMineResourceRequest() (request *ListMineResourceRequest) {
	request = &ListMineResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListMineResource")
	return
}

func NewListMineResourceResponse() (response *ListMineResourceResponse) {
	response = &ListMineResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户段用户查询自己拥有的资源
func (c *Client) ListMineResource(request *ListMineResourceRequest) (response *ListMineResourceResponse, err error) {
	if request == nil {
		request = NewListMineResourceRequest()
	}
	response = NewListMineResourceResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateDealsGatewayRequest() (request *GenerateDealsGatewayRequest) {
	request = &GenerateDealsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GenerateDealsGateway")
	return
}

func NewGenerateDealsGatewayResponse() (response *GenerateDealsGatewayResponse) {
	response = &GenerateDealsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建订单
func (c *Client) GenerateDealsGateway(request *GenerateDealsGatewayRequest) (response *GenerateDealsGatewayResponse, err error) {
	if request == nil {
		request = NewGenerateDealsGatewayRequest()
	}
	response = NewGenerateDealsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByPayModeGatewayRequest() (request *DescribeBillSummaryByPayModeGatewayRequest) {
	request = &DescribeBillSummaryByPayModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByPayModeGateway")
	return
}

func NewDescribeBillSummaryByPayModeGatewayResponse() (response *DescribeBillSummaryByPayModeGatewayResponse) {
	response = &DescribeBillSummaryByPayModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayModeGateway(request *DescribeBillSummaryByPayModeGatewayRequest) (response *DescribeBillSummaryByPayModeGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeGatewayRequest()
	}
	response = NewDescribeBillSummaryByPayModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillTrendByMonthRequest() (request *DescribeBillTrendByMonthRequest) {
	request = &DescribeBillTrendByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillTrendByMonth")
	return
}

func NewDescribeBillTrendByMonthResponse() (response *DescribeBillTrendByMonthResponse) {
	response = &DescribeBillTrendByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账单按月消费趋势
func (c *Client) DescribeBillTrendByMonth(request *DescribeBillTrendByMonthRequest) (response *DescribeBillTrendByMonthResponse, err error) {
	if request == nil {
		request = NewDescribeBillTrendByMonthRequest()
	}
	response = NewDescribeBillTrendByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTreeRequest() (request *GetProductTreeRequest) {
	request = &GetProductTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetProductTree")
	return
}

func NewGetProductTreeResponse() (response *GetProductTreeResponse) {
	response = &GetProductTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductTree(request *GetProductTreeRequest) (response *GetProductTreeResponse, err error) {
	if request == nil {
		request = NewGetProductTreeRequest()
	}
	response = NewGetProductTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeListRequest() (request *DescribeContractTypeListRequest) {
	request = &DescribeContractTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractTypeList")
	return
}

func NewDescribeContractTypeListResponse() (response *DescribeContractTypeListResponse) {
	response = &DescribeContractTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取合同类型接口
func (c *Client) DescribeContractTypeList(request *DescribeContractTypeListRequest) (response *DescribeContractTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeListRequest()
	}
	response = NewDescribeContractTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordGatewayRequest() (request *DescribeRemitRecordGatewayRequest) {
	request = &DescribeRemitRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeRemitRecordGateway")
	return
}

func NewDescribeRemitRecordGatewayResponse() (response *DescribeRemitRecordGatewayResponse) {
	response = &DescribeRemitRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户查询汇款记录
func (c *Client) DescribeRemitRecordGateway(request *DescribeRemitRecordGatewayRequest) (response *DescribeRemitRecordGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordGatewayRequest()
	}
	response = NewDescribeRemitRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPartaMailingAddressRequest() (request *ModifyPartaMailingAddressRequest) {
	request = &ModifyPartaMailingAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyPartaMailingAddress")
	return
}

func NewModifyPartaMailingAddressResponse() (response *ModifyPartaMailingAddressResponse) {
	response = &ModifyPartaMailingAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改邮寄地址
func (c *Client) ModifyPartaMailingAddress(request *ModifyPartaMailingAddressRequest) (response *ModifyPartaMailingAddressResponse, err error) {
	if request == nil {
		request = NewModifyPartaMailingAddressRequest()
	}
	response = NewModifyPartaMailingAddressResponse()
	err = c.Send(request, response)
	return
}

func NewSubscribeUpgradeRequest() (request *SubscribeUpgradeRequest) {
	request = &SubscribeUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SubscribeUpgrade")
	return
}

func NewSubscribeUpgradeResponse() (response *SubscribeUpgradeResponse) {
	response = &SubscribeUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费升配下单。
func (c *Client) SubscribeUpgrade(request *SubscribeUpgradeRequest) (response *SubscribeUpgradeResponse, err error) {
	if request == nil {
		request = NewSubscribeUpgradeRequest()
	}
	response = NewSubscribeUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealPriceGatewayRequest() (request *DescribeDealPriceGatewayRequest) {
	request = &DescribeDealPriceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDealPriceGateway")
	return
}

func NewDescribeDealPriceGatewayResponse() (response *DescribeDealPriceGatewayResponse) {
	response = &DescribeDealPriceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单价格
func (c *Client) DescribeDealPriceGateway(request *DescribeDealPriceGatewayRequest) (response *DescribeDealPriceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDealPriceGatewayRequest()
	}
	response = NewDescribeDealPriceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefaultPostInfoRequest() (request *ModifyDefaultPostInfoRequest) {
	request = &ModifyDefaultPostInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultPostInfo")
	return
}

func NewModifyDefaultPostInfoResponse() (response *ModifyDefaultPostInfoResponse) {
	response = &ModifyDefaultPostInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置默认用户的邮寄信息
func (c *Client) ModifyDefaultPostInfo(request *ModifyDefaultPostInfoRequest) (response *ModifyDefaultPostInfoResponse, err error) {
	if request == nil {
		request = NewModifyDefaultPostInfoRequest()
	}
	response = NewModifyDefaultPostInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRelationsRequest() (request *DescribeProductRelationsRequest) {
	request = &DescribeProductRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProductRelations")
	return
}

func NewDescribeProductRelationsResponse() (response *DescribeProductRelationsResponse) {
	response = &DescribeProductRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义列表
func (c *Client) DescribeProductRelations(request *DescribeProductRelationsRequest) (response *DescribeProductRelationsResponse, err error) {
	if request == nil {
		request = NewDescribeProductRelationsRequest()
	}
	response = NewDescribeProductRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemitRecordRequest() (request *CreateRemitRecordRequest) {
	request = &CreateRemitRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateRemitRecord")
	return
}

func NewCreateRemitRecordResponse() (response *CreateRemitRecordResponse) {
	response = &CreateRemitRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户新增汇款记录接口
func (c *Client) CreateRemitRecord(request *CreateRemitRecordRequest) (response *CreateRemitRecordResponse, err error) {
	if request == nil {
		request = NewCreateRemitRecordRequest()
	}
	response = NewCreateRemitRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractListRequest() (request *DescribeContractListRequest) {
	request = &DescribeContractListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeContractList")
	return
}

func NewDescribeContractListResponse() (response *DescribeContractListResponse) {
	response = &DescribeContractListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同列表接口
func (c *Client) DescribeContractList(request *DescribeContractListRequest) (response *DescribeContractListResponse, err error) {
	if request == nil {
		request = NewDescribeContractListRequest()
	}
	response = NewDescribeContractListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefaultMailingAddressGatewayRequest() (request *ModifyDefaultMailingAddressGatewayRequest) {
	request = &ModifyDefaultMailingAddressGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultMailingAddressGateway")
	return
}

func NewModifyDefaultMailingAddressGatewayResponse() (response *ModifyDefaultMailingAddressGatewayResponse) {
	response = &ModifyDefaultMailingAddressGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改默认邮寄地址
func (c *Client) ModifyDefaultMailingAddressGateway(request *ModifyDefaultMailingAddressGatewayRequest) (response *ModifyDefaultMailingAddressGatewayResponse, err error) {
	if request == nil {
		request = NewModifyDefaultMailingAddressGatewayRequest()
	}
	response = NewModifyDefaultMailingAddressGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewPrepayDealGatewayRequest() (request *PrepayDealGatewayRequest) {
	request = &PrepayDealGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "PrepayDealGateway")
	return
}

func NewPrepayDealGatewayResponse() (response *PrepayDealGatewayResponse) {
	response = &PrepayDealGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源支付
func (c *Client) PrepayDealGateway(request *PrepayDealGatewayRequest) (response *PrepayDealGatewayResponse, err error) {
	if request == nil {
		request = NewPrepayDealGatewayRequest()
	}
	response = NewPrepayDealGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductGatewayRequest() (request *UpdateProductGatewayRequest) {
	request = &UpdateProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "UpdateProductGateway")
	return
}

func NewUpdateProductGatewayResponse() (response *UpdateProductGatewayResponse) {
	response = &UpdateProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】修改产品四层定义
func (c *Client) UpdateProductGateway(request *UpdateProductGatewayRequest) (response *UpdateProductGatewayResponse, err error) {
	if request == nil {
		request = NewUpdateProductGatewayRequest()
	}
	response = NewUpdateProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByProductRequest() (request *DescribeBillFeeByProductRequest) {
	request = &DescribeBillFeeByProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByProduct")
	return
}

func NewDescribeBillFeeByProductResponse() (response *DescribeBillFeeByProductResponse) {
	response = &DescribeBillFeeByProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按产品汇总明细费用
func (c *Client) DescribeBillFeeByProduct(request *DescribeBillFeeByProductRequest) (response *DescribeBillFeeByProductResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByProductRequest()
	}
	response = NewDescribeBillFeeByProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantSubUinQuotasGatewayRequest() (request *DescribeTenantSubUinQuotasGatewayRequest) {
	request = &DescribeTenantSubUinQuotasGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeTenantSubUinQuotasGateway")
	return
}

func NewDescribeTenantSubUinQuotasGatewayResponse() (response *DescribeTenantSubUinQuotasGatewayResponse) {
	response = &DescribeTenantSubUinQuotasGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端-子账号登录后，查询当前账户在主账户下的配额信息
func (c *Client) DescribeTenantSubUinQuotasGateway(request *DescribeTenantSubUinQuotasGatewayRequest) (response *DescribeTenantSubUinQuotasGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeTenantSubUinQuotasGatewayRequest()
	}
	response = NewDescribeTenantSubUinQuotasGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadQuotasListGatewayRequest() (request *DescribeDownloadQuotasListGatewayRequest) {
	request = &DescribeDownloadQuotasListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadQuotasListGateway")
	return
}

func NewDescribeDownloadQuotasListGatewayResponse() (response *DescribeDownloadQuotasListGatewayResponse) {
	response = &DescribeDownloadQuotasListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端配额导出接口
func (c *Client) DescribeDownloadQuotasListGateway(request *DescribeDownloadQuotasListGatewayRequest) (response *DescribeDownloadQuotasListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadQuotasListGatewayRequest()
	}
	response = NewDescribeDownloadQuotasListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCouponsDeductInfoRequest() (request *DescribeCouponsDeductInfoRequest) {
	request = &DescribeCouponsDeductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeCouponsDeductInfo")
	return
}

func NewDescribeCouponsDeductInfoResponse() (response *DescribeCouponsDeductInfoResponse) {
	response = &DescribeCouponsDeductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询抵扣券作用在物品时，抵扣金额信息
func (c *Client) DescribeCouponsDeductInfo(request *DescribeCouponsDeductInfoRequest) (response *DescribeCouponsDeductInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCouponsDeductInfoRequest()
	}
	response = NewDescribeCouponsDeductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubUinQuotasRequest() (request *DescribeSubUinQuotasRequest) {
	request = &DescribeSubUinQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeSubUinQuotas")
	return
}

func NewDescribeSubUinQuotasResponse() (response *DescribeSubUinQuotasResponse) {
	response = &DescribeSubUinQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子账号配额
func (c *Client) DescribeSubUinQuotas(request *DescribeSubUinQuotasRequest) (response *DescribeSubUinQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeSubUinQuotasRequest()
	}
	response = NewDescribeSubUinQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoFlagRequest() (request *ModifyAutoFlagRequest) {
	request = &ModifyAutoFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyAutoFlag")
	return
}

func NewModifyAutoFlagResponse() (response *ModifyAutoFlagResponse) {
	response = &ModifyAutoFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置续费类型。用户id通过云api公共参数传入
func (c *Client) ModifyAutoFlag(request *ModifyAutoFlagRequest) (response *ModifyAutoFlagResponse, err error) {
	if request == nil {
		request = NewModifyAutoFlagRequest()
	}
	response = NewModifyAutoFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInfoGatewayRequest() (request *DescribeProductInfoGatewayRequest) {
	request = &DescribeProductInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProductInfoGateway")
	return
}

func NewDescribeProductInfoGatewayResponse() (response *DescribeProductInfoGatewayResponse) {
	response = &DescribeProductInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据场景信息以及额外参数显示用户已经拥有的产品信息
func (c *Client) DescribeProductInfoGateway(request *DescribeProductInfoGatewayRequest) (response *DescribeProductInfoGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductInfoGatewayRequest()
	}
	response = NewDescribeProductInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByTagGatewayRequest() (request *DescribeBillFeeByTagGatewayRequest) {
	request = &DescribeBillFeeByTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByTagGateway")
	return
}

func NewDescribeBillFeeByTagGatewayResponse() (response *DescribeBillFeeByTagGatewayResponse) {
	response = &DescribeBillFeeByTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询按标签汇总明细数据
func (c *Client) DescribeBillFeeByTagGateway(request *DescribeBillFeeByTagGatewayRequest) (response *DescribeBillFeeByTagGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByTagGatewayRequest()
	}
	response = NewDescribeBillFeeByTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQuotaRequest() (request *DeleteQuotaRequest) {
	request = &DeleteQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteQuota")
	return
}

func NewDeleteQuotaResponse() (response *DeleteQuotaResponse) {
	response = &DeleteQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口描述
func (c *Client) DeleteQuota(request *DeleteQuotaRequest) (response *DeleteQuotaResponse, err error) {
	if request == nil {
		request = NewDeleteQuotaRequest()
	}
	response = NewDeleteQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillTrendByMonthGatewayRequest() (request *DescribeBillTrendByMonthGatewayRequest) {
	request = &DescribeBillTrendByMonthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillTrendByMonthGateway")
	return
}

func NewDescribeBillTrendByMonthGatewayResponse() (response *DescribeBillTrendByMonthGatewayResponse) {
	response = &DescribeBillTrendByMonthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账单按月消费趋势
func (c *Client) DescribeBillTrendByMonthGateway(request *DescribeBillTrendByMonthGatewayRequest) (response *DescribeBillTrendByMonthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillTrendByMonthGatewayRequest()
	}
	response = NewDescribeBillTrendByMonthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsolidatedBillDownloadUrlRequest() (request *DescribeConsolidatedBillDownloadUrlRequest) {
	request = &DescribeConsolidatedBillDownloadUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeConsolidatedBillDownloadUrl")
	return
}

func NewDescribeConsolidatedBillDownloadUrlResponse() (response *DescribeConsolidatedBillDownloadUrlResponse) {
	response = &DescribeConsolidatedBillDownloadUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrl(request *DescribeConsolidatedBillDownloadUrlRequest) (response *DescribeConsolidatedBillDownloadUrlResponse, err error) {
	if request == nil {
		request = NewDescribeConsolidatedBillDownloadUrlRequest()
	}
	response = NewDescribeConsolidatedBillDownloadUrlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefaultMailingAddressRequest() (request *ModifyDefaultMailingAddressRequest) {
	request = &ModifyDefaultMailingAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDefaultMailingAddress")
	return
}

func NewModifyDefaultMailingAddressResponse() (response *ModifyDefaultMailingAddressResponse) {
	response = &ModifyDefaultMailingAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改默认邮寄地址
func (c *Client) ModifyDefaultMailingAddress(request *ModifyDefaultMailingAddressRequest) (response *ModifyDefaultMailingAddressResponse, err error) {
	if request == nil {
		request = NewModifyDefaultMailingAddressRequest()
	}
	response = NewModifyDefaultMailingAddressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadSubQuotasListGatewayRequest() (request *DescribeDownloadSubQuotasListGatewayRequest) {
	request = &DescribeDownloadSubQuotasListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadSubQuotasListGateway")
	return
}

func NewDescribeDownloadSubQuotasListGatewayResponse() (response *DescribeDownloadSubQuotasListGatewayResponse) {
	response = &DescribeDownloadSubQuotasListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端子账号配额导出接口
func (c *Client) DescribeDownloadSubQuotasListGateway(request *DescribeDownloadSubQuotasListGatewayRequest) (response *DescribeDownloadSubQuotasListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadSubQuotasListGatewayRequest()
	}
	response = NewDescribeDownloadSubQuotasListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
	request = &DescribeQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeQuota")
	return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
	response = &DescribeQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户当前使用配额
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaRequest()
	}
	response = NewDescribeQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceBillGatewayRequest() (request *GetResourceBillGatewayRequest) {
	request = &GetResourceBillGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetResourceBillGateway")
	return
}

func NewGetResourceBillGatewayResponse() (response *GetResourceBillGatewayResponse) {
	response = &GetResourceBillGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的账单信息
func (c *Client) GetResourceBillGateway(request *GetResourceBillGatewayRequest) (response *GetResourceBillGatewayResponse, err error) {
	if request == nil {
		request = NewGetResourceBillGatewayRequest()
	}
	response = NewGetResourceBillGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDetailRequest() (request *DescribeResourceBillDetailRequest) {
	request = &DescribeResourceBillDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceBillDetail")
	return
}

func NewDescribeResourceBillDetailResponse() (response *DescribeResourceBillDetailResponse) {
	response = &DescribeResourceBillDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件级别明细账单
func (c *Client) DescribeResourceBillDetail(request *DescribeResourceBillDetailRequest) (response *DescribeResourceBillDetailResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDetailRequest()
	}
	response = NewDescribeResourceBillDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithAuthGatewayRequest() (request *DescribeDownloadWithAuthGatewayRequest) {
	request = &DescribeDownloadWithAuthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadWithAuthGateway")
	return
}

func NewDescribeDownloadWithAuthGatewayResponse() (response *DescribeDownloadWithAuthGatewayResponse) {
	response = &DescribeDownloadWithAuthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithAuthGateway(request *DescribeDownloadWithAuthGatewayRequest) (response *DescribeDownloadWithAuthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithAuthGatewayRequest()
	}
	response = NewDescribeDownloadWithAuthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPartaMailingAddressGatewayRequest() (request *ModifyPartaMailingAddressGatewayRequest) {
	request = &ModifyPartaMailingAddressGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyPartaMailingAddressGateway")
	return
}

func NewModifyPartaMailingAddressGatewayResponse() (response *ModifyPartaMailingAddressGatewayResponse) {
	response = &ModifyPartaMailingAddressGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改邮寄地址
func (c *Client) ModifyPartaMailingAddressGateway(request *ModifyPartaMailingAddressGatewayRequest) (response *ModifyPartaMailingAddressGatewayResponse, err error) {
	if request == nil {
		request = NewModifyPartaMailingAddressGatewayRequest()
	}
	response = NewModifyPartaMailingAddressGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpenInvoiceInfoGatewayRequest() (request *DescribeOpenInvoiceInfoGatewayRequest) {
	request = &DescribeOpenInvoiceInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeOpenInvoiceInfoGateway")
	return
}

func NewDescribeOpenInvoiceInfoGatewayResponse() (response *DescribeOpenInvoiceInfoGatewayResponse) {
	response = &DescribeOpenInvoiceInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户开票信息
func (c *Client) DescribeOpenInvoiceInfoGateway(request *DescribeOpenInvoiceInfoGatewayRequest) (response *DescribeOpenInvoiceInfoGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeOpenInvoiceInfoGatewayRequest()
	}
	response = NewDescribeOpenInvoiceInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddPartaMailingAddressGatewayRequest() (request *AddPartaMailingAddressGatewayRequest) {
	request = &AddPartaMailingAddressGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddPartaMailingAddressGateway")
	return
}

func NewAddPartaMailingAddressGatewayResponse() (response *AddPartaMailingAddressGatewayResponse) {
	response = &AddPartaMailingAddressGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加邮寄地址
func (c *Client) AddPartaMailingAddressGateway(request *AddPartaMailingAddressGatewayRequest) (response *AddPartaMailingAddressGatewayResponse, err error) {
	if request == nil {
		request = NewAddPartaMailingAddressGatewayRequest()
	}
	response = NewAddPartaMailingAddressGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByResourceGatewayRequest() (request *DescribeBillSummaryByResourceGatewayRequest) {
	request = &DescribeBillSummaryByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByResourceGateway")
	return
}

func NewDescribeBillSummaryByResourceGatewayResponse() (response *DescribeBillSummaryByResourceGatewayResponse) {
	response = &DescribeBillSummaryByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按资源汇总数据
func (c *Client) DescribeBillSummaryByResourceGateway(request *DescribeBillSummaryByResourceGatewayRequest) (response *DescribeBillSummaryByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByResourceGatewayRequest()
	}
	response = NewDescribeBillSummaryByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRefundRequest() (request *DescribeRefundRequest) {
	request = &DescribeRefundRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeRefund")
	return
}

func NewDescribeRefundResponse() (response *DescribeRefundResponse) {
	response = &DescribeRefundResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询提现(退款)信息
func (c *Client) DescribeRefund(request *DescribeRefundRequest) (response *DescribeRefundResponse, err error) {
	if request == nil {
		request = NewDescribeRefundRequest()
	}
	response = NewDescribeRefundResponse()
	err = c.Send(request, response)
	return
}

func NewSetTagGatewayRequest() (request *SetTagGatewayRequest) {
	request = &SetTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SetTagGateway")
	return
}

func NewSetTagGatewayResponse() (response *SetTagGatewayResponse) {
	response = &SetTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置标签
func (c *Client) SetTagGateway(request *SetTagGatewayRequest) (response *SetTagGatewayResponse, err error) {
	if request == nil {
		request = NewSetTagGatewayRequest()
	}
	response = NewSetTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDetailByResourceRequest() (request *DescribeBillDetailByResourceRequest) {
	request = &DescribeBillDetailByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDetailByResource")
	return
}

func NewDescribeBillDetailByResourceResponse() (response *DescribeBillDetailByResourceResponse) {
	response = &DescribeBillDetailByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源花费详情
func (c *Client) DescribeBillDetailByResource(request *DescribeBillDetailByResourceRequest) (response *DescribeBillDetailByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillDetailByResourceRequest()
	}
	response = NewDescribeBillDetailByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByPayModeGatewayRequest() (request *DescribeBillFeeByPayModeGatewayRequest) {
	request = &DescribeBillFeeByPayModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByPayModeGateway")
	return
}

func NewDescribeBillFeeByPayModeGatewayResponse() (response *DescribeBillFeeByPayModeGatewayResponse) {
	response = &DescribeBillFeeByPayModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按计费模式汇总明细费用
func (c *Client) DescribeBillFeeByPayModeGateway(request *DescribeBillFeeByPayModeGatewayRequest) (response *DescribeBillFeeByPayModeGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByPayModeGatewayRequest()
	}
	response = NewDescribeBillFeeByPayModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetTagRequest() (request *SetTagRequest) {
	request = &SetTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SetTag")
	return
}

func NewSetTagResponse() (response *SetTagResponse) {
	response = &SetTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置标签
func (c *Client) SetTag(request *SetTagRequest) (response *SetTagResponse, err error) {
	if request == nil {
		request = NewSetTagRequest()
	}
	response = NewSetTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByResourceRequest() (request *DescribeBillSummaryByResourceRequest) {
	request = &DescribeBillSummaryByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByResource")
	return
}

func NewDescribeBillSummaryByResourceResponse() (response *DescribeBillSummaryByResourceResponse) {
	response = &DescribeBillSummaryByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按资源汇总数据
func (c *Client) DescribeBillSummaryByResource(request *DescribeBillSummaryByResourceRequest) (response *DescribeBillSummaryByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByResourceRequest()
	}
	response = NewDescribeBillSummaryByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDealStatusRequest() (request *ModifyDealStatusRequest) {
	request = &ModifyDealStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDealStatus")
	return
}

func NewModifyDealStatusResponse() (response *ModifyDealStatusResponse) {
	response = &ModifyDealStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据大订单号修改大订单状态
//
// 当订单状态为待支付时调用此接口可将订单状态改为取消
//
// 当订单状态为取消时调用此接口可将订单状态改为删除
func (c *Client) ModifyDealStatus(request *ModifyDealStatusRequest) (response *ModifyDealStatusResponse, err error) {
	if request == nil {
		request = NewModifyDealStatusRequest()
	}
	response = NewModifyDealStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountPeriodRequest() (request *DescribeAccountPeriodRequest) {
	request = &DescribeAccountPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeAccountPeriod")
	return
}

func NewDescribeAccountPeriodResponse() (response *DescribeAccountPeriodResponse) {
	response = &DescribeAccountPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账期信息
func (c *Client) DescribeAccountPeriod(request *DescribeAccountPeriodRequest) (response *DescribeAccountPeriodResponse, err error) {
	if request == nil {
		request = NewDescribeAccountPeriodRequest()
	}
	response = NewDescribeAccountPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
	request = &DescribeBillSummaryByTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByTag")
	return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
	response = &DescribeBillSummaryByTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签汇总费用分布
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByTagRequest()
	}
	response = NewDescribeBillSummaryByTagResponse()
	err = c.Send(request, response)
	return
}

func NewCancelTagRequest() (request *CancelTagRequest) {
	request = &CancelTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CancelTag")
	return
}

func NewCancelTagResponse() (response *CancelTagResponse) {
	response = &CancelTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消标签
func (c *Client) CancelTag(request *CancelTagRequest) (response *CancelTagResponse, err error) {
	if request == nil {
		request = NewCancelTagRequest()
	}
	response = NewCancelTagResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaGatewayRequest() (request *GetQuotaGatewayRequest) {
	request = &GetQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetQuotaGateway")
	return
}

func NewGetQuotaGatewayResponse() (response *GetQuotaGatewayResponse) {
	response = &GetQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询配额
func (c *Client) GetQuotaGateway(request *GetQuotaGatewayRequest) (response *GetQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewGetQuotaGatewayRequest()
	}
	response = NewGetQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadByResourceRequest() (request *DescribeBillDownloadByResourceRequest) {
	request = &DescribeBillDownloadByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadByResource")
	return
}

func NewDescribeBillDownloadByResourceResponse() (response *DescribeBillDownloadByResourceResponse) {
	response = &DescribeBillDownloadByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单按资源汇总
func (c *Client) DescribeBillDownloadByResource(request *DescribeBillDownloadByResourceRequest) (response *DescribeBillDownloadByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadByResourceRequest()
	}
	response = NewDescribeBillDownloadByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewHourModifyGatewayRequest() (request *HourModifyGatewayRequest) {
	request = &HourModifyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourModifyGateway")
	return
}

func NewHourModifyGatewayResponse() (response *HourModifyGatewayResponse) {
	response = &HourModifyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源变配
func (c *Client) HourModifyGateway(request *HourModifyGatewayRequest) (response *HourModifyGatewayResponse, err error) {
	if request == nil {
		request = NewHourModifyGatewayRequest()
	}
	response = NewHourModifyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostInfoListGatewayRequest() (request *DescribePostInfoListGatewayRequest) {
	request = &DescribePostInfoListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribePostInfoListGateway")
	return
}

func NewDescribePostInfoListGatewayResponse() (response *DescribePostInfoListGatewayResponse) {
	response = &DescribePostInfoListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的邮寄信息
func (c *Client) DescribePostInfoListGateway(request *DescribePostInfoListGatewayRequest) (response *DescribePostInfoListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribePostInfoListGatewayRequest()
	}
	response = NewDescribePostInfoListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddQuotaRequest() (request *AddQuotaRequest) {
	request = &AddQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddQuota")
	return
}

func NewAddQuotaResponse() (response *AddQuotaResponse) {
	response = &AddQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增,修改quota配置  支持一次传入多个配额配置。
//
// quotaKey需符合四层定义
//
func (c *Client) AddQuota(request *AddQuotaRequest) (response *AddQuotaResponse, err error) {
	if request == nil {
		request = NewAddQuotaRequest()
	}
	response = NewAddQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccountGatewayRequest() (request *DescribeUserAccountGatewayRequest) {
	request = &DescribeUserAccountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeUserAccountGateway")
	return
}

func NewDescribeUserAccountGatewayResponse() (response *DescribeUserAccountGatewayResponse) {
	response = &DescribeUserAccountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay
func (c *Client) DescribeUserAccountGateway(request *DescribeUserAccountGatewayRequest) (response *DescribeUserAccountGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccountGatewayRequest()
	}
	response = NewDescribeUserAccountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBspDownloadRecordListGatewayRequest() (request *DescribeBspDownloadRecordListGatewayRequest) {
	request = &DescribeBspDownloadRecordListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBspDownloadRecordListGateway")
	return
}

func NewDescribeBspDownloadRecordListGatewayResponse() (response *DescribeBspDownloadRecordListGatewayResponse) {
	response = &DescribeBspDownloadRecordListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBspDownloadRecordListGateway(request *DescribeBspDownloadRecordListGatewayRequest) (response *DescribeBspDownloadRecordListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBspDownloadRecordListGatewayRequest()
	}
	response = NewDescribeBspDownloadRecordListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantQuotasRequest() (request *DescribeTenantQuotasRequest) {
	request = &DescribeTenantQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeTenantQuotas")
	return
}

func NewDescribeTenantQuotasResponse() (response *DescribeTenantQuotasResponse) {
	response = &DescribeTenantQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端产品配额列表查询
func (c *Client) DescribeTenantQuotas(request *DescribeTenantQuotasRequest) (response *DescribeTenantQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeTenantQuotasRequest()
	}
	response = NewDescribeTenantQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewSubscribeOrderRequest() (request *SubscribeOrderRequest) {
	request = &SubscribeOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SubscribeOrder")
	return
}

func NewSubscribeOrderResponse() (response *SubscribeOrderResponse) {
	response = &SubscribeOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 普通预付费下单
func (c *Client) SubscribeOrder(request *SubscribeOrderRequest) (response *SubscribeOrderResponse, err error) {
	if request == nil {
		request = NewSubscribeOrderRequest()
	}
	response = NewSubscribeOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceDetailGatewayRequest() (request *DescribeInvoiceDetailGatewayRequest) {
	request = &DescribeInvoiceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceDetailGateway")
	return
}

func NewDescribeInvoiceDetailGatewayResponse() (response *DescribeInvoiceDetailGatewayResponse) {
	response = &DescribeInvoiceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取发票的开票详情
func (c *Client) DescribeInvoiceDetailGateway(request *DescribeInvoiceDetailGatewayRequest) (response *DescribeInvoiceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceDetailGatewayRequest()
	}
	response = NewDescribeInvoiceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsGatewayRequest() (request *DescribeProductsGatewayRequest) {
	request = &DescribeProductsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProductsGateway")
	return
}

func NewDescribeProductsGatewayResponse() (response *DescribeProductsGatewayResponse) {
	response = &DescribeProductsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取商品信息(第一层)
func (c *Client) DescribeProductsGateway(request *DescribeProductsGatewayRequest) (response *DescribeProductsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductsGatewayRequest()
	}
	response = NewDescribeProductsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewPayDealsGatewayRequest() (request *PayDealsGatewayRequest) {
	request = &PayDealsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "PayDealsGateway")
	return
}

func NewPayDealsGatewayResponse() (response *PayDealsGatewayResponse) {
	response = &PayDealsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源支付
func (c *Client) PayDealsGateway(request *PayDealsGatewayRequest) (response *PayDealsGatewayResponse, err error) {
	if request == nil {
		request = NewPayDealsGatewayRequest()
	}
	response = NewPayDealsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCouponInfoRequest() (request *CouponInfoRequest) {
	request = &CouponInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CouponInfo")
	return
}

func NewCouponInfoResponse() (response *CouponInfoResponse) {
	response = &CouponInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询代金券信息
func (c *Client) CouponInfo(request *CouponInfoRequest) (response *CouponInfoResponse, err error) {
	if request == nil {
		request = NewCouponInfoRequest()
	}
	response = NewCouponInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByPayModeRequest() (request *DescribeBillFeeByPayModeRequest) {
	request = &DescribeBillFeeByPayModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeByPayMode")
	return
}

func NewDescribeBillFeeByPayModeResponse() (response *DescribeBillFeeByPayModeResponse) {
	response = &DescribeBillFeeByPayModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按计费模式汇总明细费用
func (c *Client) DescribeBillFeeByPayMode(request *DescribeBillFeeByPayModeRequest) (response *DescribeBillFeeByPayModeResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByPayModeRequest()
	}
	response = NewDescribeBillFeeByPayModeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDealStatusGatewayRequest() (request *ModifyDealStatusGatewayRequest) {
	request = &ModifyDealStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyDealStatusGateway")
	return
}

func NewModifyDealStatusGatewayResponse() (response *ModifyDealStatusGatewayResponse) {
	response = &ModifyDealStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据大订单号修改大订单状态
//
// 当订单状态为待支付时调用此接口可将订单状态改为取消
//
// 当订单状态为取消时调用此接口可将订单状态改为删除
func (c *Client) ModifyDealStatusGateway(request *ModifyDealStatusGatewayRequest) (response *ModifyDealStatusGatewayResponse, err error) {
	if request == nil {
		request = NewModifyDealStatusGatewayRequest()
	}
	response = NewModifyDealStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddQuotaGatewayRequest() (request *AddQuotaGatewayRequest) {
	request = &AddQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddQuotaGateway")
	return
}

func NewAddQuotaGatewayResponse() (response *AddQuotaGatewayResponse) {
	response = &AddQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增,修改quota配置  支持一次传入多个配额配置。
//
// quotaKey需符合四层定义
//
func (c *Client) AddQuotaGateway(request *AddQuotaGatewayRequest) (response *AddQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewAddQuotaGatewayRequest()
	}
	response = NewAddQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDetailListGatewayRequest() (request *DescribeResourceDetailListGatewayRequest) {
	request = &DescribeResourceDetailListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceDetailListGateway")
	return
}

func NewDescribeResourceDetailListGatewayResponse() (response *DescribeResourceDetailListGatewayResponse) {
	response = &DescribeResourceDetailListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户续费参数
func (c *Client) DescribeResourceDetailListGateway(request *DescribeResourceDetailListGatewayRequest) (response *DescribeResourceDetailListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDetailListGatewayRequest()
	}
	response = NewDescribeResourceDetailListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetMineResourceBillRequest() (request *GetMineResourceBillRequest) {
	request = &GetMineResourceBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetMineResourceBill")
	return
}

func NewGetMineResourceBillResponse() (response *GetMineResourceBillResponse) {
	response = &GetMineResourceBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户自身的账单信息
func (c *Client) GetMineResourceBill(request *GetMineResourceBillRequest) (response *GetMineResourceBillResponse, err error) {
	if request == nil {
		request = NewGetMineResourceBillRequest()
	}
	response = NewGetMineResourceBillResponse()
	err = c.Send(request, response)
	return
}

func NewAddProductGatewayRequest() (request *AddProductGatewayRequest) {
	request = &AddProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "AddProductGateway")
	return
}

func NewAddProductGatewayResponse() (response *AddProductGatewayResponse) {
	response = &AddProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加产品四层定义
func (c *Client) AddProductGateway(request *AddProductGatewayRequest) (response *AddProductGatewayResponse, err error) {
	if request == nil {
		request = NewAddProductGatewayRequest()
	}
	response = NewAddProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeTrendGatewayRequest() (request *DescribeBillFeeTrendGatewayRequest) {
	request = &DescribeBillFeeTrendGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillFeeTrendGateway")
	return
}

func NewDescribeBillFeeTrendGatewayResponse() (response *DescribeBillFeeTrendGatewayResponse) {
	response = &DescribeBillFeeTrendGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类汇总明细费用趋势
func (c *Client) DescribeBillFeeTrendGateway(request *DescribeBillFeeTrendGatewayRequest) (response *DescribeBillFeeTrendGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeTrendGatewayRequest()
	}
	response = NewDescribeBillFeeTrendGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceModifyLogGatewayRequest() (request *ListResourceModifyLogGatewayRequest) {
	request = &ListResourceModifyLogGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceModifyLogGateway")
	return
}

func NewListResourceModifyLogGatewayResponse() (response *ListResourceModifyLogGatewayResponse) {
	response = &ListResourceModifyLogGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源变配记录
func (c *Client) ListResourceModifyLogGateway(request *ListResourceModifyLogGatewayRequest) (response *ListResourceModifyLogGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceModifyLogGatewayRequest()
	}
	response = NewListResourceModifyLogGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDownloadGatewayRequest() (request *DescribeResourceBillDownloadGatewayRequest) {
	request = &DescribeResourceBillDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceBillDownloadGateway")
	return
}

func NewDescribeResourceBillDownloadGatewayResponse() (response *DescribeResourceBillDownloadGatewayResponse) {
	response = &DescribeResourceBillDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量明细导出
func (c *Client) DescribeResourceBillDownloadGateway(request *DescribeResourceBillDownloadGatewayRequest) (response *DescribeResourceBillDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDownloadGatewayRequest()
	}
	response = NewDescribeResourceBillDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceListGatewayRequest() (request *DescribeInvoiceListGatewayRequest) {
	request = &DescribeInvoiceListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceListGateway")
	return
}

func NewDescribeInvoiceListGatewayResponse() (response *DescribeInvoiceListGatewayResponse) {
	response = &DescribeInvoiceListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取用户开票记录
func (c *Client) DescribeInvoiceListGateway(request *DescribeInvoiceListGatewayRequest) (response *DescribeInvoiceListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceListGatewayRequest()
	}
	response = NewDescribeInvoiceListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
	request = &DescribeRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeRegion")
	return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
	response = &DescribeRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域.  用户id通过公共参数UIN传入
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
	if request == nil {
		request = NewDescribeRegionRequest()
	}
	response = NewDescribeRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccountRequest() (request *DescribeUserAccountRequest) {
	request = &DescribeUserAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeUserAccount")
	return
}

func NewDescribeUserAccountResponse() (response *DescribeUserAccountResponse) {
	response = &DescribeUserAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay
func (c *Client) DescribeUserAccount(request *DescribeUserAccountRequest) (response *DescribeUserAccountResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccountRequest()
	}
	response = NewDescribeUserAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRelationsGatewayRequest() (request *DescribeProductRelationsGatewayRequest) {
	request = &DescribeProductRelationsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeProductRelationsGateway")
	return
}

func NewDescribeProductRelationsGatewayResponse() (response *DescribeProductRelationsGatewayResponse) {
	response = &DescribeProductRelationsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义列表
func (c *Client) DescribeProductRelationsGateway(request *DescribeProductRelationsGatewayRequest) (response *DescribeProductRelationsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductRelationsGatewayRequest()
	}
	response = NewDescribeProductRelationsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewPayDealsRequest() (request *PayDealsRequest) {
	request = &PayDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "PayDeals")
	return
}

func NewPayDealsResponse() (response *PayDealsResponse) {
	response = &PayDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源支付
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
	if request == nil {
		request = NewPayDealsRequest()
	}
	response = NewPayDealsResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagListRequest() (request *GetTagListRequest) {
	request = &GetTagListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetTagList")
	return
}

func NewGetTagListResponse() (response *GetTagListResponse) {
	response = &GetTagListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当月标签列表（分页）
func (c *Client) GetTagList(request *GetTagListRequest) (response *GetTagListResponse, err error) {
	if request == nil {
		request = NewGetTagListRequest()
	}
	response = NewGetTagListResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListGatewayRequest() (request *ExportDealListGatewayRequest) {
	request = &ExportDealListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ExportDealListGateway")
	return
}

func NewExportDealListGatewayResponse() (response *ExportDealListGatewayResponse) {
	response = &ExportDealListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入订单列表信息
func (c *Client) ExportDealListGateway(request *ExportDealListGatewayRequest) (response *ExportDealListGatewayResponse, err error) {
	if request == nil {
		request = NewExportDealListGatewayRequest()
	}
	response = NewExportDealListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetSubUinQuotaRequest() (request *SetSubUinQuotaRequest) {
	request = &SetSubUinQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SetSubUinQuota")
	return
}

func NewSetSubUinQuotaResponse() (response *SetSubUinQuotaResponse) {
	response = &SetSubUinQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置子账号配额
func (c *Client) SetSubUinQuota(request *SetSubUinQuotaRequest) (response *SetSubUinQuotaResponse, err error) {
	if request == nil {
		request = NewSetSubUinQuotaRequest()
	}
	response = NewSetSubUinQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
	request = &DescribeBillSummaryByProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByProduct")
	return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
	response = &DescribeBillSummaryByProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByProductRequest()
	}
	response = NewDescribeBillSummaryByProductResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContractRequest() (request *CreateContractRequest) {
	request = &CreateContractRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateContract")
	return
}

func NewCreateContractResponse() (response *CreateContractResponse) {
	response = &CreateContractResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建合同
func (c *Client) CreateContract(request *CreateContractRequest) (response *CreateContractResponse, err error) {
	if request == nil {
		request = NewCreateContractRequest()
	}
	response = NewCreateContractResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDownloadRequest() (request *DescribeResourceBillDownloadRequest) {
	request = &DescribeResourceBillDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceBillDownload")
	return
}

func NewDescribeResourceBillDownloadResponse() (response *DescribeResourceBillDownloadResponse) {
	response = &DescribeResourceBillDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量明细导出
func (c *Client) DescribeResourceBillDownload(request *DescribeResourceBillDownloadRequest) (response *DescribeResourceBillDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDownloadRequest()
	}
	response = NewDescribeResourceBillDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReconciliationGatewayRequest() (request *ModifyReconciliationGatewayRequest) {
	request = &ModifyReconciliationGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyReconciliationGateway")
	return
}

func NewModifyReconciliationGatewayResponse() (response *ModifyReconciliationGatewayResponse) {
	response = &ModifyReconciliationGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交调账
//
func (c *Client) ModifyReconciliationGateway(request *ModifyReconciliationGatewayRequest) (response *ModifyReconciliationGatewayResponse, err error) {
	if request == nil {
		request = NewModifyReconciliationGatewayRequest()
	}
	response = NewModifyReconciliationGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePayOrderRequest() (request *DescribePayOrderRequest) {
	request = &DescribePayOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribePayOrder")
	return
}

func NewDescribePayOrderResponse() (response *DescribePayOrderResponse) {
	response = &DescribePayOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询支付订单。
func (c *Client) DescribePayOrder(request *DescribePayOrderRequest) (response *DescribePayOrderResponse, err error) {
	if request == nil {
		request = NewDescribePayOrderRequest()
	}
	response = NewDescribePayOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceAmountGatewayRequest() (request *DescribeInvoiceAmountGatewayRequest) {
	request = &DescribeInvoiceAmountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeInvoiceAmountGateway")
	return
}

func NewDescribeInvoiceAmountGatewayResponse() (response *DescribeInvoiceAmountGatewayResponse) {
	response = &DescribeInvoiceAmountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可开票金额
func (c *Client) DescribeInvoiceAmountGateway(request *DescribeInvoiceAmountGatewayRequest) (response *DescribeInvoiceAmountGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceAmountGatewayRequest()
	}
	response = NewDescribeInvoiceAmountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReconciliationListGatewayRequest() (request *DescribeReconciliationListGatewayRequest) {
	request = &DescribeReconciliationListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeReconciliationListGateway")
	return
}

func NewDescribeReconciliationListGatewayResponse() (response *DescribeReconciliationListGatewayResponse) {
	response = &DescribeReconciliationListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调账记录查询调账列表
func (c *Client) DescribeReconciliationListGateway(request *DescribeReconciliationListGatewayRequest) (response *DescribeReconciliationListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeReconciliationListGatewayRequest()
	}
	response = NewDescribeReconciliationListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordRequest() (request *DescribeRemitRecordRequest) {
	request = &DescribeRemitRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeRemitRecord")
	return
}

func NewDescribeRemitRecordResponse() (response *DescribeRemitRecordResponse) {
	response = &DescribeRemitRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户查询汇款记录
func (c *Client) DescribeRemitRecord(request *DescribeRemitRecordRequest) (response *DescribeRemitRecordResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordRequest()
	}
	response = NewDescribeRemitRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGetDownloadUrlWithAuthGatewayRequest() (request *DescribeGetDownloadUrlWithAuthGatewayRequest) {
	request = &DescribeGetDownloadUrlWithAuthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeGetDownloadUrlWithAuthGateway")
	return
}

func NewDescribeGetDownloadUrlWithAuthGatewayResponse() (response *DescribeGetDownloadUrlWithAuthGatewayResponse) {
	response = &DescribeGetDownloadUrlWithAuthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeGetDownloadUrlWithAuthGateway(request *DescribeGetDownloadUrlWithAuthGatewayRequest) (response *DescribeGetDownloadUrlWithAuthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeGetDownloadUrlWithAuthGatewayRequest()
	}
	response = NewDescribeGetDownloadUrlWithAuthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
	request = &DescribeBillSummaryByPayModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByPayMode")
	return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
	response = &DescribeBillSummaryByPayModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeRequest()
	}
	response = NewDescribeBillSummaryByPayModeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInvoiceRequest() (request *CreateInvoiceRequest) {
	request = &CreateInvoiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateInvoice")
	return
}

func NewCreateInvoiceResponse() (response *CreateInvoiceResponse) {
	response = &CreateInvoiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户申请开票
func (c *Client) CreateInvoice(request *CreateInvoiceRequest) (response *CreateInvoiceResponse, err error) {
	if request == nil {
		request = NewCreateInvoiceRequest()
	}
	response = NewCreateInvoiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDetailGatewayRequest() (request *DescribeResourceBillDetailGatewayRequest) {
	request = &DescribeResourceBillDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceBillDetailGateway")
	return
}

func NewDescribeResourceBillDetailGatewayResponse() (response *DescribeResourceBillDetailGatewayResponse) {
	response = &DescribeResourceBillDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件级别明细账单
func (c *Client) DescribeResourceBillDetailGateway(request *DescribeResourceBillDetailGatewayRequest) (response *DescribeResourceBillDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDetailGatewayRequest()
	}
	response = NewDescribeResourceBillDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemitRecordGatewayRequest() (request *CreateRemitRecordGatewayRequest) {
	request = &CreateRemitRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateRemitRecordGateway")
	return
}

func NewCreateRemitRecordGatewayResponse() (response *CreateRemitRecordGatewayResponse) {
	response = &CreateRemitRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户新增汇款记录接口
func (c *Client) CreateRemitRecordGateway(request *CreateRemitRecordGatewayRequest) (response *CreateRemitRecordGatewayResponse, err error) {
	if request == nil {
		request = NewCreateRemitRecordGatewayRequest()
	}
	response = NewCreateRemitRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
	request = &DeleteProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeleteProduct")
	return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
	response = &DeleteProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】删除产品四层定义
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
	if request == nil {
		request = NewDeleteProductRequest()
	}
	response = NewDeleteProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealPriceRequest() (request *DescribeDealPriceRequest) {
	request = &DescribeDealPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDealPrice")
	return
}

func NewDescribeDealPriceResponse() (response *DescribeDealPriceResponse) {
	response = &DescribeDealPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单价格
func (c *Client) DescribeDealPrice(request *DescribeDealPriceRequest) (response *DescribeDealPriceResponse, err error) {
	if request == nil {
		request = NewDescribeDealPriceRequest()
	}
	response = NewDescribeDealPriceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPublicAccountRequest() (request *QueryPublicAccountRequest) {
	request = &QueryPublicAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "QueryPublicAccount")
	return
}

func NewQueryPublicAccountResponse() (response *QueryPublicAccountResponse) {
	response = &QueryPublicAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共帐户信息
func (c *Client) QueryPublicAccount(request *QueryPublicAccountRequest) (response *QueryPublicAccountResponse, err error) {
	if request == nil {
		request = NewQueryPublicAccountRequest()
	}
	response = NewQueryPublicAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceListGatewayRequest() (request *DescribeResourceListGatewayRequest) {
	request = &DescribeResourceListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceListGateway")
	return
}

func NewDescribeResourceListGatewayResponse() (response *DescribeResourceListGatewayResponse) {
	response = &DescribeResourceListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户产品列表，只支持到一层产品
func (c *Client) DescribeResourceListGateway(request *DescribeResourceListGatewayRequest) (response *DescribeResourceListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceListGatewayRequest()
	}
	response = NewDescribeResourceListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadListGatewayRequest() (request *DescribeBillDownloadListGatewayRequest) {
	request = &DescribeBillDownloadListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadListGateway")
	return
}

func NewDescribeBillDownloadListGatewayResponse() (response *DescribeBillDownloadListGatewayResponse) {
	response = &DescribeBillDownloadListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单下载地址列表
func (c *Client) DescribeBillDownloadListGateway(request *DescribeBillDownloadListGatewayRequest) (response *DescribeBillDownloadListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadListGatewayRequest()
	}
	response = NewDescribeBillDownloadListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagListGatewayRequest() (request *GetTagListGatewayRequest) {
	request = &GetTagListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetTagListGateway")
	return
}

func NewGetTagListGatewayResponse() (response *GetTagListGatewayResponse) {
	response = &GetTagListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当月标签列表（分页）
func (c *Client) GetTagListGateway(request *GetTagListGatewayRequest) (response *GetTagListGatewayResponse, err error) {
	if request == nil {
		request = NewGetTagListGatewayRequest()
	}
	response = NewGetTagListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceDetailDownloadGatewayRequest() (request *ListResourceDetailDownloadGatewayRequest) {
	request = &ListResourceDetailDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceDetailDownloadGateway")
	return
}

func NewListResourceDetailDownloadGatewayResponse() (response *ListResourceDetailDownloadGatewayResponse) {
	response = &ListResourceDetailDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Uin，ResourceId, RegionId, ProductCode查询资源详细信息
func (c *Client) ListResourceDetailDownloadGateway(request *ListResourceDetailDownloadGatewayRequest) (response *ListResourceDetailDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceDetailDownloadGatewayRequest()
	}
	response = NewListResourceDetailDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReconciliationRequest() (request *ModifyReconciliationRequest) {
	request = &ModifyReconciliationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyReconciliation")
	return
}

func NewModifyReconciliationResponse() (response *ModifyReconciliationResponse) {
	response = &ModifyReconciliationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交调账
//
func (c *Client) ModifyReconciliation(request *ModifyReconciliationRequest) (response *ModifyReconciliationResponse, err error) {
	if request == nil {
		request = NewModifyReconciliationRequest()
	}
	response = NewModifyReconciliationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubUinQuotasGatewayRequest() (request *DescribeSubUinQuotasGatewayRequest) {
	request = &DescribeSubUinQuotasGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeSubUinQuotasGateway")
	return
}

func NewDescribeSubUinQuotasGatewayResponse() (response *DescribeSubUinQuotasGatewayResponse) {
	response = &DescribeSubUinQuotasGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子账号配额
func (c *Client) DescribeSubUinQuotasGateway(request *DescribeSubUinQuotasGatewayRequest) (response *DescribeSubUinQuotasGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeSubUinQuotasGatewayRequest()
	}
	response = NewDescribeSubUinQuotasGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningRequest() (request *DescribeWarningRequest) {
	request = &DescribeWarningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeWarning")
	return
}

func NewDescribeWarningResponse() (response *DescribeWarningResponse) {
	response = &DescribeWarningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户查询预警状态和阈值
func (c *Client) DescribeWarning(request *DescribeWarningRequest) (response *DescribeWarningResponse, err error) {
	if request == nil {
		request = NewDescribeWarningRequest()
	}
	response = NewDescribeWarningResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaGatewayRequest() (request *DescribeQuotaGatewayRequest) {
	request = &DescribeQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeQuotaGateway")
	return
}

func NewDescribeQuotaGatewayResponse() (response *DescribeQuotaGatewayResponse) {
	response = &DescribeQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户当前使用配额
func (c *Client) DescribeQuotaGateway(request *DescribeQuotaGatewayRequest) (response *DescribeQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaGatewayRequest()
	}
	response = NewDescribeQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWarningRequest() (request *CreateWarningRequest) {
	request = &CreateWarningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateWarning")
	return
}

func NewCreateWarningResponse() (response *CreateWarningResponse) {
	response = &CreateWarningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加或更新预警信息
func (c *Client) CreateWarning(request *CreateWarningRequest) (response *CreateWarningResponse, err error) {
	if request == nil {
		request = NewCreateWarningRequest()
	}
	response = NewCreateWarningResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadResourceDetailRequest() (request *DescribeBillDownloadResourceDetailRequest) {
	request = &DescribeBillDownloadResourceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillDownloadResourceDetail")
	return
}

func NewDescribeBillDownloadResourceDetailResponse() (response *DescribeBillDownloadResourceDetailResponse) {
	response = &DescribeBillDownloadResourceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单明细
func (c *Client) DescribeBillDownloadResourceDetail(request *DescribeBillDownloadResourceDetailRequest) (response *DescribeBillDownloadResourceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadResourceDetailRequest()
	}
	response = NewDescribeBillDownloadResourceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCancelInvoiceGatewayRequest() (request *CancelInvoiceGatewayRequest) {
	request = &CancelInvoiceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CancelInvoiceGateway")
	return
}

func NewCancelInvoiceGatewayResponse() (response *CancelInvoiceGatewayResponse) {
	response = &CancelInvoiceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消开票申请
func (c *Client) CancelInvoiceGateway(request *CancelInvoiceGatewayRequest) (response *CancelInvoiceGatewayResponse, err error) {
	if request == nil {
		request = NewCancelInvoiceGatewayRequest()
	}
	response = NewCancelInvoiceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateDealsRequest() (request *GenerateDealsRequest) {
	request = &GenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GenerateDeals")
	return
}

func NewGenerateDealsResponse() (response *GenerateDealsResponse) {
	response = &GenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建订单
func (c *Client) GenerateDeals(request *GenerateDealsRequest) (response *GenerateDealsResponse, err error) {
	if request == nil {
		request = NewGenerateDealsRequest()
	}
	response = NewGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductRequest() (request *UpdateProductRequest) {
	request = &UpdateProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "UpdateProduct")
	return
}

func NewUpdateProductResponse() (response *UpdateProductResponse) {
	response = &UpdateProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】修改产品四层定义
func (c *Client) UpdateProduct(request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
	if request == nil {
		request = NewUpdateProductRequest()
	}
	response = NewUpdateProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByProductGatewayRequest() (request *DescribeBillSummaryByProductGatewayRequest) {
	request = &DescribeBillSummaryByProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByProductGateway")
	return
}

func NewDescribeBillSummaryByProductGatewayResponse() (response *DescribeBillSummaryByProductGatewayResponse) {
	response = &DescribeBillSummaryByProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProductGateway(request *DescribeBillSummaryByProductGatewayRequest) (response *DescribeBillSummaryByProductGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByProductGatewayRequest()
	}
	response = NewDescribeBillSummaryByProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadVoucherListFileRequest() (request *DescribeDownloadVoucherListFileRequest) {
	request = &DescribeDownloadVoucherListFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadVoucherListFile")
	return
}

func NewDescribeDownloadVoucherListFileResponse() (response *DescribeDownloadVoucherListFileResponse) {
	response = &DescribeDownloadVoucherListFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载代金券信息excel
func (c *Client) DescribeDownloadVoucherListFile(request *DescribeDownloadVoucherListFileRequest) (response *DescribeDownloadVoucherListFileResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadVoucherListFileRequest()
	}
	response = NewDescribeDownloadVoucherListFileResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPublicAccountGatewayRequest() (request *QueryPublicAccountGatewayRequest) {
	request = &QueryPublicAccountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "QueryPublicAccountGateway")
	return
}

func NewQueryPublicAccountGatewayResponse() (response *QueryPublicAccountGatewayResponse) {
	response = &QueryPublicAccountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共帐户信息
func (c *Client) QueryPublicAccountGateway(request *QueryPublicAccountGatewayRequest) (response *QueryPublicAccountGatewayResponse, err error) {
	if request == nil {
		request = NewQueryPublicAccountGatewayRequest()
	}
	response = NewQueryPublicAccountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewHourUnblockRequest() (request *HourUnblockRequest) {
	request = &HourUnblockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourUnblock")
	return
}

func NewHourUnblockResponse() (response *HourUnblockResponse) {
	response = &HourUnblockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源销毁
func (c *Client) HourUnblock(request *HourUnblockRequest) (response *HourUnblockResponse, err error) {
	if request == nil {
		request = NewHourUnblockRequest()
	}
	response = NewHourUnblockResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadVoucherListFileGatewayRequest() (request *DescribeDownloadVoucherListFileGatewayRequest) {
	request = &DescribeDownloadVoucherListFileGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeDownloadVoucherListFileGateway")
	return
}

func NewDescribeDownloadVoucherListFileGatewayResponse() (response *DescribeDownloadVoucherListFileGatewayResponse) {
	response = &DescribeDownloadVoucherListFileGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载代金券信息excel
func (c *Client) DescribeDownloadVoucherListFileGateway(request *DescribeDownloadVoucherListFileGatewayRequest) (response *DescribeDownloadVoucherListFileGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadVoucherListFileGatewayRequest()
	}
	response = NewDescribeDownloadVoucherListFileGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewRefundRequest() (request *RefundRequest) {
	request = &RefundRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "Refund")
	return
}

func NewRefundResponse() (response *RefundResponse) {
	response = &RefundResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提现（退款）
func (c *Client) Refund(request *RefundRequest) (response *RefundResponse, err error) {
	if request == nil {
		request = NewRefundRequest()
	}
	response = NewRefundResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPostInfoGatewayRequest() (request *ModifyPostInfoGatewayRequest) {
	request = &ModifyPostInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyPostInfoGateway")
	return
}

func NewModifyPostInfoGatewayResponse() (response *ModifyPostInfoGatewayResponse) {
	response = &ModifyPostInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改新增用户的邮寄信息
func (c *Client) ModifyPostInfoGateway(request *ModifyPostInfoGatewayRequest) (response *ModifyPostInfoGatewayResponse, err error) {
	if request == nil {
		request = NewModifyPostInfoGatewayRequest()
	}
	response = NewModifyPostInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTreeGatewayRequest() (request *GetProductTreeGatewayRequest) {
	request = &GetProductTreeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetProductTreeGateway")
	return
}

func NewGetProductTreeGatewayResponse() (response *GetProductTreeGatewayResponse) {
	response = &GetProductTreeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductTreeGateway(request *GetProductTreeGatewayRequest) (response *GetProductTreeGatewayResponse, err error) {
	if request == nil {
		request = NewGetProductTreeGatewayRequest()
	}
	response = NewGetProductTreeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceModifyLogRequest() (request *ListResourceModifyLogRequest) {
	request = &ListResourceModifyLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListResourceModifyLog")
	return
}

func NewListResourceModifyLogResponse() (response *ListResourceModifyLogResponse) {
	response = &ListResourceModifyLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源变配记录
func (c *Client) ListResourceModifyLog(request *ListResourceModifyLogRequest) (response *ListResourceModifyLogResponse, err error) {
	if request == nil {
		request = NewListResourceModifyLogRequest()
	}
	response = NewListResourceModifyLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDownloadRecordGatewayRequest() (request *CreateDownloadRecordGatewayRequest) {
	request = &CreateDownloadRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "CreateDownloadRecordGateway")
	return
}

func NewCreateDownloadRecordGatewayResponse() (response *CreateDownloadRecordGatewayResponse) {
	response = &CreateDownloadRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户点击下载时通知后台
func (c *Client) CreateDownloadRecordGateway(request *CreateDownloadRecordGatewayRequest) (response *CreateDownloadRecordGatewayResponse, err error) {
	if request == nil {
		request = NewCreateDownloadRecordGatewayRequest()
	}
	response = NewCreateDownloadRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewHourModifyRequest() (request *HourModifyRequest) {
	request = &HourModifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "HourModify")
	return
}

func NewHourModifyResponse() (response *HourModifyResponse) {
	response = &HourModifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费资源变配
func (c *Client) HourModify(request *HourModifyRequest) (response *HourModifyResponse, err error) {
	if request == nil {
		request = NewHourModifyRequest()
	}
	response = NewHourModifyResponse()
	err = c.Send(request, response)
	return
}

func NewSetAutoFlagRequest() (request *SetAutoFlagRequest) {
	request = &SetAutoFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "SetAutoFlag")
	return
}

func NewSetAutoFlagResponse() (response *SetAutoFlagResponse) {
	response = &SetAutoFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置续费类型。用户id通过云api公共参数传入
func (c *Client) SetAutoFlag(request *SetAutoFlagRequest) (response *SetAutoFlagResponse, err error) {
	if request == nil {
		request = NewSetAutoFlagRequest()
	}
	response = NewSetAutoFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePostInfoGatewayRequest() (request *DeletePostInfoGatewayRequest) {
	request = &DeletePostInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DeletePostInfoGateway")
	return
}

func NewDeletePostInfoGatewayResponse() (response *DeletePostInfoGatewayResponse) {
	response = &DeletePostInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户的邮寄信息
func (c *Client) DeletePostInfoGateway(request *DeletePostInfoGatewayRequest) (response *DeletePostInfoGatewayResponse, err error) {
	if request == nil {
		request = NewDeletePostInfoGatewayRequest()
	}
	response = NewDeletePostInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
	request = &DescribeBillSummaryByRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeBillSummaryByRegion")
	return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
	response = &DescribeBillSummaryByRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByRegionRequest()
	}
	response = NewDescribeBillSummaryByRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDetailListRequest() (request *DescribeResourceDetailListRequest) {
	request = &DescribeResourceDetailListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceDetailList")
	return
}

func NewDescribeResourceDetailListResponse() (response *DescribeResourceDetailListResponse) {
	response = &DescribeResourceDetailListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户续费参数
func (c *Client) DescribeResourceDetailList(request *DescribeResourceDetailListRequest) (response *DescribeResourceDetailListResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDetailListRequest()
	}
	response = NewDescribeResourceDetailListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDownloadGatewayRequest() (request *DescribeResourceDownloadGatewayRequest) {
	request = &DescribeResourceDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeResourceDownloadGateway")
	return
}

func NewDescribeResourceDownloadGatewayResponse() (response *DescribeResourceDownloadGatewayResponse) {
	response = &DescribeResourceDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订购实例导出
func (c *Client) DescribeResourceDownloadGateway(request *DescribeResourceDownloadGatewayRequest) (response *DescribeResourceDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDownloadGatewayRequest()
	}
	response = NewDescribeResourceDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractStatusRequest() (request *ModifyContractStatusRequest) {
	request = &ModifyContractStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ModifyContractStatus")
	return
}

func NewModifyContractStatusResponse() (response *ModifyContractStatusResponse) {
	response = &ModifyContractStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同状态
func (c *Client) ModifyContractStatus(request *ModifyContractStatusRequest) (response *ModifyContractStatusResponse, err error) {
	if request == nil {
		request = NewModifyContractStatusRequest()
	}
	response = NewModifyContractStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountWaterRequest() (request *DescribeAccountWaterRequest) {
	request = &DescribeAccountWaterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "DescribeAccountWater")
	return
}

func NewDescribeAccountWaterResponse() (response *DescribeAccountWaterResponse) {
	response = &DescribeAccountWaterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户的还款或销账流水
func (c *Client) DescribeAccountWater(request *DescribeAccountWaterRequest) (response *DescribeAccountWaterResponse, err error) {
	if request == nil {
		request = NewDescribeAccountWaterRequest()
	}
	response = NewDescribeAccountWaterResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceBillRequest() (request *GetResourceBillRequest) {
	request = &GetResourceBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "GetResourceBill")
	return
}

func NewGetResourceBillResponse() (response *GetResourceBillResponse) {
	response = &GetResourceBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的账单信息
func (c *Client) GetResourceBill(request *GetResourceBillRequest) (response *GetResourceBillResponse, err error) {
	if request == nil {
		request = NewGetResourceBillRequest()
	}
	response = NewGetResourceBillResponse()
	err = c.Send(request, response)
	return
}

func NewListMineResourceGatewayRequest() (request *ListMineResourceGatewayRequest) {
	request = &ListMineResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bill", APIVersion, "ListMineResourceGateway")
	return
}

func NewListMineResourceGatewayResponse() (response *ListMineResourceGatewayResponse) {
	response = &ListMineResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户段用户查询自己拥有的资源
func (c *Client) ListMineResourceGateway(request *ListMineResourceGatewayRequest) (response *ListMineResourceGatewayResponse, err error) {
	if request == nil {
		request = NewListMineResourceGatewayRequest()
	}
	response = NewListMineResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

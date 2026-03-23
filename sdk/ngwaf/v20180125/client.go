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

package v20180125

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-01-25"

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

func NewDescribeAccessExportsRequest() (request *DescribeAccessExportsRequest) {
	request = &DescribeAccessExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessExports")
	return
}

func NewDescribeAccessExportsResponse() (response *DescribeAccessExportsResponse) {
	response = &DescribeAccessExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志导出列表
func (c *Client) DescribeAccessExports(request *DescribeAccessExportsRequest) (response *DescribeAccessExportsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessExportsRequest()
	}
	response = NewDescribeAccessExportsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainEngineRequest() (request *ModifyDomainEngineRequest) {
	request = &ModifyDomainEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainEngine")
	return
}

func NewModifyDomainEngineResponse() (response *ModifyDomainEngineResponse) {
	response = &ModifyDomainEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名引擎类型
func (c *Client) ModifyDomainEngine(request *ModifyDomainEngineRequest) (response *ModifyDomainEngineResponse, err error) {
	if request == nil {
		request = NewModifyDomainEngineRequest()
	}
	response = NewModifyDomainEngineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailVersionTwoRequest() (request *DescribeBotDetailVersionTwoRequest) {
	request = &DescribeBotDetailVersionTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotDetailVersionTwo")
	return
}

func NewDescribeBotDetailVersionTwoResponse() (response *DescribeBotDetailVersionTwoResponse) {
	response = &DescribeBotDetailVersionTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot详情显示v2版本
func (c *Client) DescribeBotDetailVersionTwo(request *DescribeBotDetailVersionTwoRequest) (response *DescribeBotDetailVersionTwoResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailVersionTwoRequest()
	}
	response = NewDescribeBotDetailVersionTwoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotApiEventAlarmRequest() (request *DescribeBotApiEventAlarmRequest) {
	request = &DescribeBotApiEventAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotApiEventAlarm")
	return
}

func NewDescribeBotApiEventAlarmResponse() (response *DescribeBotApiEventAlarmResponse) {
	response = &DescribeBotApiEventAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot/api事件告警配置
func (c *Client) DescribeBotApiEventAlarm(request *DescribeBotApiEventAlarmRequest) (response *DescribeBotApiEventAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeBotApiEventAlarmRequest()
	}
	response = NewDescribeBotApiEventAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBPreinstallRuleRequest() (request *DescribeBotUCBPreinstallRuleRequest) {
	request = &DescribeBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotUCBPreinstallRule")
	return
}

func NewDescribeBotUCBPreinstallRuleResponse() (response *DescribeBotUCBPreinstallRuleResponse) {
	response = &DescribeBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB预设策略
func (c *Client) DescribeBotUCBPreinstallRule(request *DescribeBotUCBPreinstallRuleRequest) (response *DescribeBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBPreinstallRuleRequest()
	}
	response = NewDescribeBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSignatureRuleV2Request() (request *DescribeUserSignatureRuleV2Request) {
	request = &DescribeUserSignatureRuleV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserSignatureRuleV2")
	return
}

func NewDescribeUserSignatureRuleV2Response() (response *DescribeUserSignatureRuleV2Response) {
	response = &DescribeUserSignatureRuleV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户特征规则列表
func (c *Client) DescribeUserSignatureRuleV2(request *DescribeUserSignatureRuleV2Request) (response *DescribeUserSignatureRuleV2Response, err error) {
	if request == nil {
		request = NewDescribeUserSignatureRuleV2Request()
	}
	response = NewDescribeUserSignatureRuleV2Response()
	err = c.Send(request, response)
	return
}

func NewAddCustomWhiteRuleRequest() (request *AddCustomWhiteRuleRequest) {
	request = &AddCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddCustomWhiteRule")
	return
}

func NewAddCustomWhiteRuleResponse() (response *AddCustomWhiteRuleResponse) {
	response = &AddCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加精准白名单规则
func (c *Client) AddCustomWhiteRule(request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddCustomWhiteRuleRequest()
	}
	response = NewAddCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiFakeRulesRequest() (request *DescribeAntiFakeRulesRequest) {
	request = &DescribeAntiFakeRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAntiFakeRules")
	return
}

func NewDescribeAntiFakeRulesResponse() (response *DescribeAntiFakeRulesResponse) {
	response = &DescribeAntiFakeRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防篡改url
func (c *Client) DescribeAntiFakeRules(request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiFakeRulesRequest()
	}
	response = NewDescribeAntiFakeRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartUserRequest() (request *DescribeSpartUserRequest) {
	request = &DescribeSpartUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSpartUser")
	return
}

func NewDescribeSpartUserResponse() (response *DescribeSpartUserResponse) {
	response = &DescribeSpartUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息，判断用户是否开通waf，套餐购买情况等信息
func (c *Client) DescribeSpartUser(request *DescribeSpartUserRequest) (response *DescribeSpartUserResponse, err error) {
	if request == nil {
		request = NewDescribeSpartUserRequest()
	}
	response = NewDescribeSpartUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHybridClusterRequest() (request *ModifyHybridClusterRequest) {
	request = &ModifyHybridClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHybridCluster")
	return
}

func NewModifyHybridClusterResponse() (response *ModifyHybridClusterResponse) {
	response = &ModifyHybridClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改混合云集群
func (c *Client) ModifyHybridCluster(request *ModifyHybridClusterRequest) (response *ModifyHybridClusterResponse, err error) {
	if request == nil {
		request = NewModifyHybridClusterRequest()
	}
	response = NewModifyHybridClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiFakeUrlRequest() (request *DescribeAntiFakeUrlRequest) {
	request = &DescribeAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAntiFakeUrl")
	return
}

func NewDescribeAntiFakeUrlResponse() (response *DescribeAntiFakeUrlResponse) {
	response = &DescribeAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防篡改url
func (c *Client) DescribeAntiFakeUrl(request *DescribeAntiFakeUrlRequest) (response *DescribeAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewDescribeAntiFakeUrlRequest()
	}
	response = NewDescribeAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewResourcesOnlineRefundRequest() (request *ResourcesOnlineRefundRequest) {
	request = &ResourcesOnlineRefundRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ResourcesOnlineRefund")
	return
}

func NewResourcesOnlineRefundResponse() (response *ResourcesOnlineRefundResponse) {
	response = &ResourcesOnlineRefundResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot qps白名单
func (c *Client) ResourcesOnlineRefund(request *ResourcesOnlineRefundRequest) (response *ResourcesOnlineRefundResponse, err error) {
	if request == nil {
		request = NewResourcesOnlineRefundRequest()
	}
	response = NewResourcesOnlineRefundResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackIndexRequest() (request *DescribeAttackIndexRequest) {
	request = &DescribeAttackIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackIndex")
	return
}

func NewDescribeAttackIndexResponse() (response *DescribeAttackIndexResponse) {
	response = &DescribeAttackIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 攻击日志索引描述详情
func (c *Client) DescribeAttackIndex(request *DescribeAttackIndexRequest) (response *DescribeAttackIndexResponse, err error) {
	if request == nil {
		request = NewDescribeAttackIndexRequest()
	}
	response = NewDescribeAttackIndexResponse()
	err = c.Send(request, response)
	return
}

func NewKeysLocalStorageRequest() (request *KeysLocalStorageRequest) {
	request = &KeysLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "KeysLocalStorage")
	return
}

func NewKeysLocalStorageResponse() (response *KeysLocalStorageResponse) {
	response = &KeysLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程缓存的所有键
func (c *Client) KeysLocalStorage(request *KeysLocalStorageRequest) (response *KeysLocalStorageResponse, err error) {
	if request == nil {
		request = NewKeysLocalStorageRequest()
	}
	response = NewKeysLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExpiredInstancesRequest() (request *DescribeExpiredInstancesRequest) {
	request = &DescribeExpiredInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeExpiredInstances")
	return
}

func NewDescribeExpiredInstancesResponse() (response *DescribeExpiredInstancesResponse) {
	response = &DescribeExpiredInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有销毁实例的详细信息
func (c *Client) DescribeExpiredInstances(request *DescribeExpiredInstancesRequest) (response *DescribeExpiredInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeExpiredInstancesRequest()
	}
	response = NewDescribeExpiredInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUseCdnWithoutOpenProxyDomainsRequest() (request *DescribeUseCdnWithoutOpenProxyDomainsRequest) {
	request = &DescribeUseCdnWithoutOpenProxyDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUseCdnWithoutOpenProxyDomains")
	return
}

func NewDescribeUseCdnWithoutOpenProxyDomainsResponse() (response *DescribeUseCdnWithoutOpenProxyDomainsResponse) {
	response = &DescribeUseCdnWithoutOpenProxyDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否存在未开启waf代理设置但实际存在七层代理的域名
func (c *Client) DescribeUseCdnWithoutOpenProxyDomains(request *DescribeUseCdnWithoutOpenProxyDomainsRequest) (response *DescribeUseCdnWithoutOpenProxyDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeUseCdnWithoutOpenProxyDomainsRequest()
	}
	response = NewDescribeUseCdnWithoutOpenProxyDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpAccessControlRequest() (request *ModifyIpAccessControlRequest) {
	request = &ModifyIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyIpAccessControl")
	return
}

func NewModifyIpAccessControlResponse() (response *ModifyIpAccessControlResponse) {
	response = &ModifyIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单编辑接口
func (c *Client) ModifyIpAccessControl(request *ModifyIpAccessControlRequest) (response *ModifyIpAccessControlResponse, err error) {
	if request == nil {
		request = NewModifyIpAccessControlRequest()
	}
	response = NewModifyIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAreaBanRuleRequest() (request *CreateAreaBanRuleRequest) {
	request = &CreateAreaBanRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAreaBanRule")
	return
}

func NewCreateAreaBanRuleResponse() (response *CreateAreaBanRuleResponse) {
	response = &CreateAreaBanRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加（编辑）地域封禁中的地域信息
func (c *Client) CreateAreaBanRule(request *CreateAreaBanRuleRequest) (response *CreateAreaBanRuleResponse, err error) {
	if request == nil {
		request = NewCreateAreaBanRuleRequest()
	}
	response = NewCreateAreaBanRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBFeatureRuleRequest() (request *DescribeBotUCBFeatureRuleRequest) {
	request = &DescribeBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotUCBFeatureRule")
	return
}

func NewDescribeBotUCBFeatureRuleResponse() (response *DescribeBotUCBFeatureRuleResponse) {
	response = &DescribeBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB自定义策略
func (c *Client) DescribeBotUCBFeatureRule(request *DescribeBotUCBFeatureRuleRequest) (response *DescribeBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBFeatureRuleRequest()
	}
	response = NewDescribeBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePostCKafkaFlowRequest() (request *CreatePostCKafkaFlowRequest) {
	request = &CreatePostCKafkaFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreatePostCKafkaFlow")
	return
}

func NewCreatePostCKafkaFlowResponse() (response *CreatePostCKafkaFlowResponse) {
	response = &CreatePostCKafkaFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CKafka投递流任务
func (c *Client) CreatePostCKafkaFlow(request *CreatePostCKafkaFlowRequest) (response *CreatePostCKafkaFlowResponse, err error) {
	if request == nil {
		request = NewCreatePostCKafkaFlowRequest()
	}
	response = NewCreatePostCKafkaFlowResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProtectionModesRequest() (request *UpdateProtectionModesRequest) {
	request = &UpdateProtectionModesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpdateProtectionModes")
	return
}

func NewUpdateProtectionModesResponse() (response *UpdateProtectionModesResponse) {
	response = &UpdateProtectionModesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Tiga引擎下大类规则的防护模式
func (c *Client) UpdateProtectionModes(request *UpdateProtectionModesRequest) (response *UpdateProtectionModesResponse, err error) {
	if request == nil {
		request = NewUpdateProtectionModesRequest()
	}
	response = NewUpdateProtectionModesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartaProtectionInfoRequest() (request *DescribeSpartaProtectionInfoRequest) {
	request = &DescribeSpartaProtectionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSpartaProtectionInfo")
	return
}

func NewDescribeSpartaProtectionInfoResponse() (response *DescribeSpartaProtectionInfoResponse) {
	response = &DescribeSpartaProtectionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-获取防护域名信息
func (c *Client) DescribeSpartaProtectionInfo(request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSpartaProtectionInfoRequest()
	}
	response = NewDescribeSpartaProtectionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewRuleUploadRequest() (request *RuleUploadRequest) {
	request = &RuleUploadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RuleUpload")
	return
}

func NewRuleUploadResponse() (response *RuleUploadResponse) {
	response = &RuleUploadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传tiga规则库信息到mysql
func (c *Client) RuleUpload(request *RuleUploadRequest) (response *RuleUploadResponse, err error) {
	if request == nil {
		request = NewRuleUploadRequest()
	}
	response = NewRuleUploadResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveLocalStorageItemRequest() (request *RemoveLocalStorageItemRequest) {
	request = &RemoveLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RemoveLocalStorageItem")
	return
}

func NewRemoveLocalStorageItemResponse() (response *RemoveLocalStorageItemResponse) {
	response = &RemoveLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除远程缓存
func (c *Client) RemoveLocalStorageItem(request *RemoveLocalStorageItemRequest) (response *RemoveLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewRemoveLocalStorageItemRequest()
	}
	response = NewRemoveLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchCustomRuleRequest() (request *DeleteBatchCustomRuleRequest) {
	request = &DeleteBatchCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBatchCustomRule")
	return
}

func NewDeleteBatchCustomRuleResponse() (response *DeleteBatchCustomRuleResponse) {
	response = &DeleteBatchCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-批量删除接口
func (c *Client) DeleteBatchCustomRule(request *DeleteBatchCustomRuleRequest) (response *DeleteBatchCustomRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBatchCustomRuleRequest()
	}
	response = NewDeleteBatchCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostCKafkaFlowsRequest() (request *DescribePostCKafkaFlowsRequest) {
	request = &DescribePostCKafkaFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePostCKafkaFlows")
	return
}

func NewDescribePostCKafkaFlowsResponse() (response *DescribePostCKafkaFlowsResponse) {
	response = &DescribePostCKafkaFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CKafka投递流任务列表
func (c *Client) DescribePostCKafkaFlows(request *DescribePostCKafkaFlowsRequest) (response *DescribePostCKafkaFlowsResponse, err error) {
	if request == nil {
		request = NewDescribePostCKafkaFlowsRequest()
	}
	response = NewDescribePostCKafkaFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLLMLimitRuleRequest() (request *DeleteLLMLimitRuleRequest) {
	request = &DeleteLLMLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteLLMLimitRule")
	return
}

func NewDeleteLLMLimitRuleResponse() (response *DeleteLLMLimitRuleResponse) {
	response = &DeleteLLMLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除大模型算力消耗防护规则
func (c *Client) DeleteLLMLimitRule(request *DeleteLLMLimitRuleRequest) (response *DeleteLLMLimitRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLLMLimitRuleRequest()
	}
	response = NewDeleteLLMLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportsRequest() (request *DescribeExportsRequest) {
	request = &DescribeExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeExports")
	return
}

func NewDescribeExportsResponse() (response *DescribeExportsResponse) {
	response = &DescribeExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志下载任务列表
func (c *Client) DescribeExports(request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
	if request == nil {
		request = NewDescribeExportsRequest()
	}
	response = NewDescribeExportsResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyPostCLSFlowRequest() (request *DestroyPostCLSFlowRequest) {
	request = &DestroyPostCLSFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DestroyPostCLSFlow")
	return
}

func NewDestroyPostCLSFlowResponse() (response *DestroyPostCLSFlowResponse) {
	response = &DestroyPostCLSFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁CLS投递流任务
func (c *Client) DestroyPostCLSFlow(request *DestroyPostCLSFlowRequest) (response *DestroyPostCLSFlowResponse, err error) {
	if request == nil {
		request = NewDestroyPostCLSFlowRequest()
	}
	response = NewDestroyPostCLSFlowResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchAreaBanRuleRequest() (request *ModifyBatchAreaBanRuleRequest) {
	request = &ModifyBatchAreaBanRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchAreaBanRule")
	return
}

func NewModifyBatchAreaBanRuleResponse() (response *ModifyBatchAreaBanRuleResponse) {
	response = &ModifyBatchAreaBanRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改批量地域封禁规则
func (c *Client) ModifyBatchAreaBanRule(request *ModifyBatchAreaBanRuleRequest) (response *ModifyBatchAreaBanRuleResponse, err error) {
	if request == nil {
		request = NewModifyBatchAreaBanRuleRequest()
	}
	response = NewModifyBatchAreaBanRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceLimitRulesRequest() (request *CreateInstanceLimitRulesRequest) {
	request = &CreateInstanceLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateInstanceLimitRules")
	return
}

func NewCreateInstanceLimitRulesResponse() (response *CreateInstanceLimitRulesResponse) {
	response = &CreateInstanceLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例限流规则
func (c *Client) CreateInstanceLimitRules(request *CreateInstanceLimitRulesRequest) (response *CreateInstanceLimitRulesResponse, err error) {
	if request == nil {
		request = NewCreateInstanceLimitRulesRequest()
	}
	response = NewCreateInstanceLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTldTokenRequest() (request *DescribeBotTldTokenRequest) {
	request = &DescribeBotTldTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTldToken")
	return
}

func NewDescribeBotTldTokenResponse() (response *DescribeBotTldTokenResponse) {
	response = &DescribeBotTldTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示图灵盾token的匹配规则
func (c *Client) DescribeBotTldToken(request *DescribeBotTldTokenRequest) (response *DescribeBotTldTokenResponse, err error) {
	if request == nil {
		request = NewDescribeBotTldTokenRequest()
	}
	response = NewDescribeBotTldTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSignatureClassRequest() (request *DescribeUserSignatureClassRequest) {
	request = &DescribeUserSignatureClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserSignatureClass")
	return
}

func NewDescribeUserSignatureClassResponse() (response *DescribeUserSignatureClassResponse) {
	response = &DescribeUserSignatureClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Tiga引擎规则类型及状态
func (c *Client) DescribeUserSignatureClass(request *DescribeUserSignatureClassRequest) (response *DescribeUserSignatureClassResponse, err error) {
	if request == nil {
		request = NewDescribeUserSignatureClassRequest()
	}
	response = NewDescribeUserSignatureClassResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaBanRuleRequest() (request *DescribeAreaBanRuleRequest) {
	request = &DescribeAreaBanRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAreaBanRule")
	return
}

func NewDescribeAreaBanRuleResponse() (response *DescribeAreaBanRuleResponse) {
	response = &DescribeAreaBanRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地域封禁规则配置
func (c *Client) DescribeAreaBanRule(request *DescribeAreaBanRuleRequest) (response *DescribeAreaBanRuleResponse, err error) {
	if request == nil {
		request = NewDescribeAreaBanRuleRequest()
	}
	response = NewDescribeAreaBanRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateSignRequest() (request *GenerateSignRequest) {
	request = &GenerateSignRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GenerateSign")
	return
}

func NewGenerateSignResponse() (response *GenerateSignResponse) {
	response = &GenerateSignResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费产生签名
func (c *Client) GenerateSign(request *GenerateSignRequest) (response *GenerateSignResponse, err error) {
	if request == nil {
		request = NewGenerateSignRequest()
	}
	response = NewGenerateSignResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostFlowModeInnerRequest() (request *ModifyHostFlowModeInnerRequest) {
	request = &ModifyHostFlowModeInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostFlowModeInner")
	return
}

func NewModifyHostFlowModeInnerResponse() (response *ModifyHostFlowModeInnerResponse) {
	response = &ModifyHostFlowModeInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名的流量模式
func (c *Client) ModifyHostFlowModeInner(request *ModifyHostFlowModeInnerRequest) (response *ModifyHostFlowModeInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostFlowModeInnerRequest()
	}
	response = NewModifyHostFlowModeInnerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 释放已经过期的实例
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewEnableClientMsgRequest() (request *EnableClientMsgRequest) {
	request = &EnableClientMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EnableClientMsg")
	return
}

func NewEnableClientMsgResponse() (response *EnableClientMsgResponse) {
	response = &EnableClientMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关开启后，会将客户端的ip和port透传到后端
func (c *Client) EnableClientMsg(request *EnableClientMsgRequest) (response *EnableClientMsgResponse, err error) {
	if request == nil {
		request = NewEnableClientMsgRequest()
	}
	response = NewEnableClientMsgResponse()
	err = c.Send(request, response)
	return
}

func NewAddDomainWhiteRuleRequest() (request *AddDomainWhiteRuleRequest) {
	request = &AddDomainWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddDomainWhiteRule")
	return
}

func NewAddDomainWhiteRuleResponse() (response *AddDomainWhiteRuleResponse) {
	response = &AddDomainWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加域名规则白名单
func (c *Client) AddDomainWhiteRule(request *AddDomainWhiteRuleRequest) (response *AddDomainWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddDomainWhiteRuleRequest()
	}
	response = NewAddDomainWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSessionKeyRequest() (request *DeleteBotSessionKeyRequest) {
	request = &DeleteBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBotSessionKey")
	return
}

func NewDeleteBotSessionKeyResponse() (response *DeleteBotSessionKeyResponse) {
	response = &DeleteBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除bot session配置
func (c *Client) DeleteBotSessionKey(request *DeleteBotSessionKeyRequest) (response *DeleteBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewDeleteBotSessionKeyRequest()
	}
	response = NewDeleteBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneOverviewRequest() (request *DescribeBotSceneOverviewRequest) {
	request = &DescribeBotSceneOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSceneOverview")
	return
}

func NewDescribeBotSceneOverviewResponse() (response *DescribeBotSceneOverviewResponse) {
	response = &DescribeBotSceneOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Bot场景全局概览
func (c *Client) DescribeBotSceneOverview(request *DescribeBotSceneOverviewRequest) (response *DescribeBotSceneOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneOverviewRequest()
	}
	response = NewDescribeBotSceneOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUCBRecordsRequest() (request *DescribeBotUCBRecordsRequest) {
	request = &DescribeBotUCBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotUCBRecords")
	return
}

func NewDescribeBotUCBRecordsResponse() (response *DescribeBotUCBRecordsResponse) {
	response = &DescribeBotUCBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UCB类型Bots
func (c *Client) DescribeBotUCBRecords(request *DescribeBotUCBRecordsRequest) (response *DescribeBotUCBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotUCBRecordsRequest()
	}
	response = NewDescribeBotUCBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCanaryClsUserRequest() (request *GetCanaryClsUserRequest) {
	request = &GetCanaryClsUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetCanaryClsUser")
	return
}

func NewGetCanaryClsUserResponse() (response *GetCanaryClsUserResponse) {
	response = &GetCanaryClsUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cls灰度用户接口
func (c *Client) GetCanaryClsUser(request *GetCanaryClsUserRequest) (response *GetCanaryClsUserResponse, err error) {
	if request == nil {
		request = NewGetCanaryClsUserRequest()
	}
	response = NewGetCanaryClsUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppletAccessRequest() (request *CreateAppletAccessRequest) {
	request = &CreateAppletAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAppletAccess")
	return
}

func NewCreateAppletAccessResponse() (response *CreateAppletAccessResponse) {
	response = &CreateAppletAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加小程序接入
func (c *Client) CreateAppletAccess(request *CreateAppletAccessRequest) (response *CreateAppletAccessResponse, err error) {
	if request == nil {
		request = NewCreateAppletAccessRequest()
	}
	response = NewCreateAppletAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRequestCountRequest() (request *DescribeBotRequestCountRequest) {
	request = &DescribeBotRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRequestCount")
	return
}

func NewDescribeBotRequestCountResponse() (response *DescribeBotRequestCountResponse) {
	response = &DescribeBotRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Bot访问日志的总量
func (c *Client) DescribeBotRequestCount(request *DescribeBotRequestCountRequest) (response *DescribeBotRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeBotRequestCountRequest()
	}
	response = NewDescribeBotRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHybridClusterNodesRequest() (request *DescribeHybridClusterNodesRequest) {
	request = &DescribeHybridClusterNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHybridClusterNodes")
	return
}

func NewDescribeHybridClusterNodesResponse() (response *DescribeHybridClusterNodesResponse) {
	response = &DescribeHybridClusterNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看混合云集群的节点信息
func (c *Client) DescribeHybridClusterNodes(request *DescribeHybridClusterNodesRequest) (response *DescribeHybridClusterNodesResponse, err error) {
	if request == nil {
		request = NewDescribeHybridClusterNodesRequest()
	}
	response = NewDescribeHybridClusterNodesResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateDealsAndPayRequest() (request *GenerateDealsAndPayRequest) {
	request = &GenerateDealsAndPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GenerateDealsAndPay")
	return
}

func NewGenerateDealsAndPayResponse() (response *GenerateDealsAndPayResponse) {
	response = &GenerateDealsAndPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例下单接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费接口下单。目前只支持预付费下单，计费侧接口：https://tcb.xxx.com/magical-brush/docs/754661947
func (c *Client) GenerateDealsAndPay(request *GenerateDealsAndPayRequest) (response *GenerateDealsAndPayResponse, err error) {
	if request == nil {
		request = NewGenerateDealsAndPayRequest()
	}
	response = NewGenerateDealsAndPayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostCLSFlowsRequest() (request *DescribePostCLSFlowsRequest) {
	request = &DescribePostCLSFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePostCLSFlows")
	return
}

func NewDescribePostCLSFlowsResponse() (response *DescribePostCLSFlowsResponse) {
	response = &DescribePostCLSFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CLS投递流任务列表
func (c *Client) DescribePostCLSFlows(request *DescribePostCLSFlowsRequest) (response *DescribePostCLSFlowsResponse, err error) {
	if request == nil {
		request = NewDescribePostCLSFlowsRequest()
	}
	response = NewDescribePostCLSFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportRequest() (request *DeleteExportRequest) {
	request = &DeleteExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteExport")
	return
}

func NewDeleteExportResponse() (response *DeleteExportResponse) {
	response = &DeleteExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志下载任务
func (c *Client) DeleteExport(request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
	if request == nil {
		request = NewDeleteExportRequest()
	}
	response = NewDeleteExportResponse()
	err = c.Send(request, response)
	return
}

func NewTestWafMultiRegionRequest() (request *TestWafMultiRegionRequest) {
	request = &TestWafMultiRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "TestWafMultiRegion")
	return
}

func NewTestWafMultiRegionResponse() (response *TestWafMultiRegionResponse) {
	response = &TestWafMultiRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// WAF测试多地API服务接口
func (c *Client) TestWafMultiRegion(request *TestWafMultiRegionRequest) (response *TestWafMultiRegionResponse, err error) {
	if request == nil {
		request = NewTestWafMultiRegionRequest()
	}
	response = NewTestWafMultiRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafInfoRequest() (request *DescribeWafInfoRequest) {
	request = &DescribeWafInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafInfo")
	return
}

func NewDescribeWafInfoResponse() (response *DescribeWafInfoResponse) {
	response = &DescribeWafInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取负载均衡绑定的WAF信息，可以根据租户负载均衡实例ID、负载均衡监听器ID、负载均衡的域名信息来查询对应绑定的 Waf的状态信息。
// 查询的范围：负载均衡实例ID、负载均衡实例ID+监听器ID、负载均衡实例ID+监听器ID+域名。
// 可能的错误码：ResourceNotFound（没有找到对应的资源）、UnsupportedRegion（目前clb-waf只支持纽约、洛杉矶、莫斯科、巴黎、伦敦、东京地域）。
func (c *Client) DescribeWafInfo(request *DescribeWafInfoRequest) (response *DescribeWafInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWafInfoRequest()
	}
	response = NewDescribeWafInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotDetailDownloadTaskRequest() (request *ModifyBotDetailDownloadTaskRequest) {
	request = &ModifyBotDetailDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotDetailDownloadTask")
	return
}

func NewModifyBotDetailDownloadTaskResponse() (response *ModifyBotDetailDownloadTaskResponse) {
	response = &ModifyBotDetailDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建bot日志详情下载任务
func (c *Client) ModifyBotDetailDownloadTask(request *ModifyBotDetailDownloadTaskRequest) (response *ModifyBotDetailDownloadTaskResponse, err error) {
	if request == nil {
		request = NewModifyBotDetailDownloadTaskRequest()
	}
	response = NewModifyBotDetailDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModuleStatusRequest() (request *DescribeModuleStatusRequest) {
	request = &DescribeModuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeModuleStatus")
	return
}

func NewDescribeModuleStatusResponse() (response *DescribeModuleStatusResponse) {
	response = &DescribeModuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各个waf基础安全模块的开关状态，看每个模块是否开启
func (c *Client) DescribeModuleStatus(request *DescribeModuleStatusRequest) (response *DescribeModuleStatusResponse, err error) {
	if request == nil {
		request = NewDescribeModuleStatusRequest()
	}
	response = NewDescribeModuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneActionRuleRequest() (request *ModifyBotSceneActionRuleRequest) {
	request = &ModifyBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotSceneActionRule")
	return
}

func NewModifyBotSceneActionRuleResponse() (response *ModifyBotSceneActionRuleResponse) {
	response = &ModifyBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增or修改BOT场景的动作策略配置
func (c *Client) ModifyBotSceneActionRule(request *ModifyBotSceneActionRuleRequest) (response *ModifyBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneActionRuleRequest()
	}
	response = NewModifyBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBusinessRiskRuleRequest() (request *ModifyBusinessRiskRuleRequest) {
	request = &ModifyBusinessRiskRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBusinessRiskRule")
	return
}

func NewModifyBusinessRiskRuleResponse() (response *ModifyBusinessRiskRuleResponse) {
	response = &ModifyBusinessRiskRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务安全规则
func (c *Client) ModifyBusinessRiskRule(request *ModifyBusinessRiskRuleRequest) (response *ModifyBusinessRiskRuleResponse, err error) {
	if request == nil {
		request = NewModifyBusinessRiskRuleRequest()
	}
	response = NewModifyBusinessRiskRuleResponse()
	err = c.Send(request, response)
	return
}

func NewClearLocalStorageRequest() (request *ClearLocalStorageRequest) {
	request = &ClearLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ClearLocalStorage")
	return
}

func NewClearLocalStorageResponse() (response *ClearLocalStorageResponse) {
	response = &ClearLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清空远程存储的缓存
func (c *Client) ClearLocalStorage(request *ClearLocalStorageRequest) (response *ClearLocalStorageResponse, err error) {
	if request == nil {
		request = NewClearLocalStorageRequest()
	}
	response = NewClearLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostPointsRequest() (request *DescribePostPointsRequest) {
	request = &DescribePostPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePostPoints")
	return
}

func NewDescribePostPointsResponse() (response *DescribePostPointsResponse) {
	response = &DescribePostPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递流量趋势图
func (c *Client) DescribePostPoints(request *DescribePostPointsRequest) (response *DescribePostPointsResponse, err error) {
	if request == nil {
		request = NewDescribePostPointsRequest()
	}
	response = NewDescribePostPointsResponse()
	err = c.Send(request, response)
	return
}

func NewBypassSpartaProtectionRequest() (request *BypassSpartaProtectionRequest) {
	request = &BypassSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "BypassSpartaProtection")
	return
}

func NewBypassSpartaProtectionResponse() (response *BypassSpartaProtectionResponse) {
	response = &BypassSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键绕过waf
func (c *Client) BypassSpartaProtection(request *BypassSpartaProtectionRequest) (response *BypassSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewBypassSpartaProtectionRequest()
	}
	response = NewBypassSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppletAccessV3Request() (request *CreateAppletAccessV3Request) {
	request = &CreateAppletAccessV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAppletAccessV3")
	return
}

func NewCreateAppletAccessV3Response() (response *CreateAppletAccessV3Response) {
	response = &CreateAppletAccessV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加小程序接入3.0
func (c *Client) CreateAppletAccessV3(request *CreateAppletAccessV3Request) (response *CreateAppletAccessV3Response, err error) {
	if request == nil {
		request = NewCreateAppletAccessV3Request()
	}
	response = NewCreateAppletAccessV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventListRequest() (request *DescribeApiSecEventListRequest) {
	request = &DescribeApiSecEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecEventList")
	return
}

func NewDescribeApiSecEventListResponse() (response *DescribeApiSecEventListResponse) {
	response = &DescribeApiSecEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全事件列表
func (c *Client) DescribeApiSecEventList(request *DescribeApiSecEventListRequest) (response *DescribeApiSecEventListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventListRequest()
	}
	response = NewDescribeApiSecEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotUBRecordsRequest() (request *DescribeBotUBRecordsRequest) {
	request = &DescribeBotUBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotUBRecords")
	return
}

func NewDescribeBotUBRecordsResponse() (response *DescribeBotUBRecordsResponse) {
	response = &DescribeBotUBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取UB类型Bots
func (c *Client) DescribeBotUBRecords(request *DescribeBotUBRecordsRequest) (response *DescribeBotUBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotUBRecordsRequest()
	}
	response = NewDescribeBotUBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHybridClusterRequest() (request *CreateHybridClusterRequest) {
	request = &CreateHybridClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateHybridCluster")
	return
}

func NewCreateHybridClusterResponse() (response *CreateHybridClusterResponse) {
	response = &CreateHybridClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建混合云集群
func (c *Client) CreateHybridCluster(request *CreateHybridClusterRequest) (response *CreateHybridClusterResponse, err error) {
	if request == nil {
		request = NewCreateHybridClusterRequest()
	}
	response = NewCreateHybridClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorDomainsRequest() (request *DescribeMonitorDomainsRequest) {
	request = &DescribeMonitorDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeMonitorDomains")
	return
}

func NewDescribeMonitorDomainsResponse() (response *DescribeMonitorDomainsResponse) {
	response = &DescribeMonitorDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Barad云监控的域名列表
func (c *Client) DescribeMonitorDomains(request *DescribeMonitorDomainsRequest) (response *DescribeMonitorDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorDomainsRequest()
	}
	response = NewDescribeMonitorDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafAutoDenyRulesRequest() (request *DescribeWafAutoDenyRulesRequest) {
	request = &DescribeWafAutoDenyRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafAutoDenyRules")
	return
}

func NewDescribeWafAutoDenyRulesResponse() (response *DescribeWafAutoDenyRulesResponse) {
	response = &DescribeWafAutoDenyRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回ip惩罚规则详细信息
func (c *Client) DescribeWafAutoDenyRules(request *DescribeWafAutoDenyRulesRequest) (response *DescribeWafAutoDenyRulesResponse, err error) {
	if request == nil {
		request = NewDescribeWafAutoDenyRulesRequest()
	}
	response = NewDescribeWafAutoDenyRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotAiRuleRequest() (request *ModifyBotAiRuleRequest) {
	request = &ModifyBotAiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotAiRule")
	return
}

func NewModifyBotAiRuleResponse() (response *ModifyBotAiRuleResponse) {
	response = &ModifyBotAiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-ai模块配置
func (c *Client) ModifyBotAiRule(request *ModifyBotAiRuleRequest) (response *ModifyBotAiRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotAiRuleRequest()
	}
	response = NewModifyBotAiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTldTokenRequest() (request *ModifyBotTldTokenRequest) {
	request = &ModifyBotTldTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotTldToken")
	return
}

func NewModifyBotTldTokenResponse() (response *ModifyBotTldTokenResponse) {
	response = &ModifyBotTldTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot下图灵盾device_token的匹配规则
func (c *Client) ModifyBotTldToken(request *ModifyBotTldTokenRequest) (response *ModifyBotTldTokenResponse, err error) {
	if request == nil {
		request = NewModifyBotTldTokenRequest()
	}
	response = NewModifyBotTldTokenResponse()
	err = c.Send(request, response)
	return
}

func NewModifyObjectRequest() (request *ModifyObjectRequest) {
	request = &ModifyObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyObject")
	return
}

func NewModifyObjectResponse() (response *ModifyObjectResponse) {
	response = &ModifyObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护对象
func (c *Client) ModifyObject(request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
	if request == nil {
		request = NewModifyObjectRequest()
	}
	response = NewModifyObjectResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostModeRequest() (request *ModifyHostModeRequest) {
	request = &ModifyHostModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostMode")
	return
}

func NewModifyHostModeResponse() (response *ModifyHostModeResponse) {
	response = &ModifyHostModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf设置防护域名防护状态
func (c *Client) ModifyHostMode(request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
	if request == nil {
		request = NewModifyHostModeRequest()
	}
	response = NewModifyHostModeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProtectGroupRequest() (request *CreateProtectGroupRequest) {
	request = &CreateProtectGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateProtectGroup")
	return
}

func NewCreateProtectGroupResponse() (response *CreateProtectGroupResponse) {
	response = &CreateProtectGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建防护对象组
func (c *Client) CreateProtectGroup(request *CreateProtectGroupRequest) (response *CreateProtectGroupResponse, err error) {
	if request == nil {
		request = NewCreateProtectGroupRequest()
	}
	response = NewCreateProtectGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotLabelRequest() (request *DescribeBotLabelRequest) {
	request = &DescribeBotLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotLabel")
	return
}

func NewDescribeBotLabelResponse() (response *DescribeBotLabelResponse) {
	response = &DescribeBotLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot标签
func (c *Client) DescribeBotLabel(request *DescribeBotLabelRequest) (response *DescribeBotLabelResponse, err error) {
	if request == nil {
		request = NewDescribeBotLabelRequest()
	}
	response = NewDescribeBotLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
	request = &DescribeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeConfigs")
	return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
	response = &DescribeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集规则配置
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigsRequest()
	}
	response = NewDescribeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceLimitRulesRequest() (request *DescribeInstanceLimitRulesRequest) {
	request = &DescribeInstanceLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInstanceLimitRules")
	return
}

func NewDescribeInstanceLimitRulesResponse() (response *DescribeInstanceLimitRulesResponse) {
	response = &DescribeInstanceLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例限流规则
func (c *Client) DescribeInstanceLimitRules(request *DescribeInstanceLimitRulesRequest) (response *DescribeInstanceLimitRulesResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceLimitRulesRequest()
	}
	response = NewDescribeInstanceLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessLogTwoRequest() (request *DescribeBotAccessLogTwoRequest) {
	request = &DescribeBotAccessLogTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAccessLogTwo")
	return
}

func NewDescribeBotAccessLogTwoResponse() (response *DescribeBotAccessLogTwoResponse) {
	response = &DescribeBotAccessLogTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot详情页日志列表显示v2版本
func (c *Client) DescribeBotAccessLogTwo(request *DescribeBotAccessLogTwoRequest) (response *DescribeBotAccessLogTwoResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessLogTwoRequest()
	}
	response = NewDescribeBotAccessLogTwoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCiphersDetailRequest() (request *DescribeCiphersDetailRequest) {
	request = &DescribeCiphersDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCiphersDetail")
	return
}

func NewDescribeCiphersDetailResponse() (response *DescribeCiphersDetailResponse) {
	response = &DescribeCiphersDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Saas型WAF接入查询加密套件信息
func (c *Client) DescribeCiphersDetail(request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCiphersDetailRequest()
	}
	response = NewDescribeCiphersDetailResponse()
	err = c.Send(request, response)
	return
}

func NewAddAttackWhiteRuleRequest() (request *AddAttackWhiteRuleRequest) {
	request = &AddAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddAttackWhiteRule")
	return
}

func NewAddAttackWhiteRuleResponse() (response *AddAttackWhiteRuleResponse) {
	response = &AddAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，增加Tiga规则引擎白名单。
func (c *Client) AddAttackWhiteRule(request *AddAttackWhiteRuleRequest) (response *AddAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddAttackWhiteRuleRequest()
	}
	response = NewAddAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletAccessV3Request() (request *DescribeAppletAccessV3Request) {
	request = &DescribeAppletAccessV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletAccessV3")
	return
}

func NewDescribeAppletAccessV3Response() (response *DescribeAppletAccessV3Response) {
	response = &DescribeAppletAccessV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序接入列表3.0
func (c *Client) DescribeAppletAccessV3(request *DescribeAppletAccessV3Request) (response *DescribeAppletAccessV3Response, err error) {
	if request == nil {
		request = NewDescribeAppletAccessV3Request()
	}
	response = NewDescribeAppletAccessV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeCaptchaCfgRequest() (request *DescribeCaptchaCfgRequest) {
	request = &DescribeCaptchaCfgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCaptchaCfg")
	return
}

func NewDescribeCaptchaCfgResponse() (response *DescribeCaptchaCfgResponse) {
	response = &DescribeCaptchaCfgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名的验证码配置
func (c *Client) DescribeCaptchaCfg(request *DescribeCaptchaCfgRequest) (response *DescribeCaptchaCfgResponse, err error) {
	if request == nil {
		request = NewDescribeCaptchaCfgRequest()
	}
	response = NewDescribeCaptchaCfgResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotIdRuleRequest() (request *ModifyBotIdRuleRequest) {
	request = &ModifyBotIdRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotIdRule")
	return
}

func NewModifyBotIdRuleResponse() (response *ModifyBotIdRuleResponse) {
	response = &ModifyBotIdRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Bot-ID规则配置1
func (c *Client) ModifyBotIdRule(request *ModifyBotIdRuleRequest) (response *ModifyBotIdRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotIdRuleRequest()
	}
	response = NewModifyBotIdRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletTesterTaskRequest() (request *ModifyAppletTesterTaskRequest) {
	request = &ModifyAppletTesterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletTesterTask")
	return
}

func NewModifyAppletTesterTaskResponse() (response *ModifyAppletTesterTaskResponse) {
	response = &ModifyAppletTesterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改小程序压测任务
func (c *Client) ModifyAppletTesterTask(request *ModifyAppletTesterTaskRequest) (response *ModifyAppletTesterTaskResponse, err error) {
	if request == nil {
		request = NewModifyAppletTesterTaskRequest()
	}
	response = NewModifyAppletTesterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddResponsePageRequest() (request *AddResponsePageRequest) {
	request = &AddResponsePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddResponsePage")
	return
}

func NewAddResponsePageResponse() (response *AddResponsePageResponse) {
	response = &AddResponsePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义响应页面，每个用户最多允许添加5个自定义响应页面，添加完后需要审核通过才可以应用到域名。
func (c *Client) AddResponsePage(request *AddResponsePageRequest) (response *AddResponsePageResponse, err error) {
	if request == nil {
		request = NewAddResponsePageRequest()
	}
	response = NewAddResponsePageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDealsRequest() (request *CreateDealsRequest) {
	request = &CreateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateDeals")
	return
}

func NewCreateDealsResponse() (response *CreateDealsResponse) {
	response = &CreateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费资源购买、续费下单接口
func (c *Client) CreateDeals(request *CreateDealsRequest) (response *CreateDealsResponse, err error) {
	if request == nil {
		request = NewCreateDealsRequest()
	}
	response = NewCreateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTCBRecordsRequest() (request *DescribeBotTCBRecordsRequest) {
	request = &DescribeBotTCBRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTCBRecords")
	return
}

func NewDescribeBotTCBRecordsResponse() (response *DescribeBotTCBRecordsResponse) {
	response = &DescribeBotTCBRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取TCB类型Bots
func (c *Client) DescribeBotTCBRecords(request *DescribeBotTCBRecordsRequest) (response *DescribeBotTCBRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBotTCBRecordsRequest()
	}
	response = NewDescribeBotTCBRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewApplyBlockPageRequest() (request *ApplyBlockPageRequest) {
	request = &ApplyBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ApplyBlockPage")
	return
}

func NewApplyBlockPageResponse() (response *ApplyBlockPageResponse) {
	response = &ApplyBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给域名应用自定义拦截页面，对于CLB-WAF类型的域名只能应用HTTP状态码为624的自定义拦截页面。
func (c *Client) ApplyBlockPage(request *ApplyBlockPageRequest) (response *ApplyBlockPageResponse, err error) {
	if request == nil {
		request = NewApplyBlockPageRequest()
	}
	response = NewApplyBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotWhiteStatusRequest() (request *ModifyBotWhiteStatusRequest) {
	request = &ModifyBotWhiteStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotWhiteStatus")
	return
}

func NewModifyBotWhiteStatusResponse() (response *ModifyBotWhiteStatusResponse) {
	response = &ModifyBotWhiteStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改goodbot的开关
func (c *Client) ModifyBotWhiteStatus(request *ModifyBotWhiteStatusRequest) (response *ModifyBotWhiteStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotWhiteStatusRequest()
	}
	response = NewModifyBotWhiteStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotEventTypeStatRequest() (request *DescribeBotEventTypeStatRequest) {
	request = &DescribeBotEventTypeStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotEventTypeStat")
	return
}

func NewDescribeBotEventTypeStatResponse() (response *DescribeBotEventTypeStatResponse) {
	response = &DescribeBotEventTypeStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-拉取BOT事件分类统计
func (c *Client) DescribeBotEventTypeStat(request *DescribeBotEventTypeStatRequest) (response *DescribeBotEventTypeStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotEventTypeStatRequest()
	}
	response = NewDescribeBotEventTypeStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotFirstPurchaseRequest() (request *DescribeBotFirstPurchaseRequest) {
	request = &DescribeBotFirstPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotFirstPurchase")
	return
}

func NewDescribeBotFirstPurchaseResponse() (response *DescribeBotFirstPurchaseResponse) {
	response = &DescribeBotFirstPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否为第一次购买waf实例，api接口会调用计费侧查询订单接口
func (c *Client) DescribeBotFirstPurchase(request *DescribeBotFirstPurchaseRequest) (response *DescribeBotFirstPurchaseResponse, err error) {
	if request == nil {
		request = NewDescribeBotFirstPurchaseRequest()
	}
	response = NewDescribeBotFirstPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserLevelRequest() (request *ModifyUserLevelRequest) {
	request = &ModifyUserLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserLevel")
	return
}

func NewModifyUserLevelResponse() (response *ModifyUserLevelResponse) {
	response = &ModifyUserLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户防护规则等级
func (c *Client) ModifyUserLevel(request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
	if request == nil {
		request = NewModifyUserLevelRequest()
	}
	response = NewModifyUserLevelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSpartaProtectionRequest() (request *DeleteSpartaProtectionRequest) {
	request = &DeleteSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteSpartaProtection")
	return
}

func NewDeleteSpartaProtectionResponse() (response *DeleteSpartaProtectionResponse) {
	response = &DeleteSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Saas型WAF删除防护域名
func (c *Client) DeleteSpartaProtection(request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewDeleteSpartaProtectionRequest()
	}
	response = NewDeleteSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchCustomWhiteRuleRequest() (request *DeleteBatchCustomWhiteRuleRequest) {
	request = &DeleteBatchCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBatchCustomWhiteRule")
	return
}

func NewDeleteBatchCustomWhiteRuleResponse() (response *DeleteBatchCustomWhiteRuleResponse) {
	response = &DeleteBatchCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除批量精准白名单规则
func (c *Client) DeleteBatchCustomWhiteRule(request *DeleteBatchCustomWhiteRuleRequest) (response *DeleteBatchCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBatchCustomWhiteRuleRequest()
	}
	response = NewDeleteBatchCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTCBRuleRequest() (request *DescribeBotTCBRuleRequest) {
	request = &DescribeBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTCBRule")
	return
}

func NewDescribeBotTCBRuleResponse() (response *DescribeBotTCBRuleResponse) {
	response = &DescribeBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取TCB规则
func (c *Client) DescribeBotTCBRule(request *DescribeBotTCBRuleRequest) (response *DescribeBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotTCBRuleRequest()
	}
	response = NewDescribeBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
	request = &DescribeApiDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiDetail")
	return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
	response = &DescribeApiDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Api请求详情信息
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
	if request == nil {
		request = NewDescribeApiDetailRequest()
	}
	response = NewDescribeApiDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneRequest() (request *ModifyBotSceneRequest) {
	request = &ModifyBotSceneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotScene")
	return
}

func NewModifyBotSceneResponse() (response *ModifyBotSceneResponse) {
	response = &ModifyBotSceneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增or修改BOT场景配置
func (c *Client) ModifyBotScene(request *ModifyBotSceneRequest) (response *ModifyBotSceneResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneRequest()
	}
	response = NewModifyBotSceneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletBackOriginAddressRequest() (request *DescribeAppletBackOriginAddressRequest) {
	request = &DescribeAppletBackOriginAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletBackOriginAddress")
	return
}

func NewDescribeAppletBackOriginAddressResponse() (response *DescribeAppletBackOriginAddressResponse) {
	response = &DescribeAppletBackOriginAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询小程序回源地址
func (c *Client) DescribeAppletBackOriginAddress(request *DescribeAppletBackOriginAddressRequest) (response *DescribeAppletBackOriginAddressResponse, err error) {
	if request == nil {
		request = NewDescribeAppletBackOriginAddressRequest()
	}
	response = NewDescribeAppletBackOriginAddressResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWxUserGroupRequest() (request *ModifyWxUserGroupRequest) {
	request = &ModifyWxUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWxUserGroup")
	return
}

func NewModifyWxUserGroupResponse() (response *ModifyWxUserGroupResponse) {
	response = &ModifyWxUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户分组
func (c *Client) ModifyWxUserGroup(request *ModifyWxUserGroupRequest) (response *ModifyWxUserGroupResponse, err error) {
	if request == nil {
		request = NewModifyWxUserGroupRequest()
	}
	response = NewModifyWxUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotUCBPreinstallRuleRequest() (request *ModifyBotUCBPreinstallRuleRequest) {
	request = &ModifyBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotUCBPreinstallRule")
	return
}

func NewModifyBotUCBPreinstallRuleResponse() (response *ModifyBotUCBPreinstallRuleResponse) {
	response = &ModifyBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB预设规则更新
func (c *Client) ModifyBotUCBPreinstallRule(request *ModifyBotUCBPreinstallRuleRequest) (response *ModifyBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotUCBPreinstallRuleRequest()
	}
	response = NewModifyBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotUCBPreinstallRuleRequest() (request *CopyBotUCBPreinstallRuleRequest) {
	request = &CopyBotUCBPreinstallRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CopyBotUCBPreinstallRule")
	return
}

func NewCopyBotUCBPreinstallRuleResponse() (response *CopyBotUCBPreinstallRuleResponse) {
	response = &CopyBotUCBPreinstallRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB预设策略拷贝
func (c *Client) CopyBotUCBPreinstallRule(request *CopyBotUCBPreinstallRuleRequest) (response *CopyBotUCBPreinstallRuleResponse, err error) {
	if request == nil {
		request = NewCopyBotUCBPreinstallRuleRequest()
	}
	response = NewCopyBotUCBPreinstallRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneListRequest() (request *DescribeBotSceneListRequest) {
	request = &DescribeBotSceneListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSceneList")
	return
}

func NewDescribeBotSceneListResponse() (response *DescribeBotSceneListResponse) {
	response = &DescribeBotSceneListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT场景列表与概览
func (c *Client) DescribeBotSceneList(request *DescribeBotSceneListRequest) (response *DescribeBotSceneListResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneListRequest()
	}
	response = NewDescribeBotSceneListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainRulesEnRequest() (request *DescribeDomainRulesEnRequest) {
	request = &DescribeDomainRulesEnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainRulesEn")
	return
}

func NewDescribeDomainRulesEnResponse() (response *DescribeDomainRulesEnResponse) {
	response = &DescribeDomainRulesEnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取域名的防护规则列表支持英文
func (c *Client) DescribeDomainRulesEn(request *DescribeDomainRulesEnRequest) (response *DescribeDomainRulesEnResponse, err error) {
	if request == nil {
		request = NewDescribeDomainRulesEnRequest()
	}
	response = NewDescribeDomainRulesEnResponse()
	err = c.Send(request, response)
	return
}

func NewImportIpAccessControlRequest() (request *ImportIpAccessControlRequest) {
	request = &ImportIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ImportIpAccessControl")
	return
}

func NewImportIpAccessControlResponse() (response *ImportIpAccessControlResponse) {
	response = &ImportIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入IP黑白名单
func (c *Client) ImportIpAccessControl(request *ImportIpAccessControlRequest) (response *ImportIpAccessControlResponse, err error) {
	if request == nil {
		request = NewImportIpAccessControlRequest()
	}
	response = NewImportIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpRelatedEventListRequest() (request *DescribeIpRelatedEventListRequest) {
	request = &DescribeIpRelatedEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeIpRelatedEventList")
	return
}

func NewDescribeIpRelatedEventListResponse() (response *DescribeIpRelatedEventListResponse) {
	response = &DescribeIpRelatedEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取攻击源IP的关联事件信息
func (c *Client) DescribeIpRelatedEventList(request *DescribeIpRelatedEventListRequest) (response *DescribeIpRelatedEventListResponse, err error) {
	if request == nil {
		request = NewDescribeIpRelatedEventListRequest()
	}
	response = NewDescribeIpRelatedEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrafficMarkingRequest() (request *DescribeTrafficMarkingRequest) {
	request = &DescribeTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTrafficMarking")
	return
}

func NewDescribeTrafficMarkingResponse() (response *DescribeTrafficMarkingResponse) {
	response = &DescribeTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询流量标记配置V2
func (c *Client) DescribeTrafficMarking(request *DescribeTrafficMarkingRequest) (response *DescribeTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewDescribeTrafficMarkingRequest()
	}
	response = NewDescribeTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRegionsStatRequest() (request *DescribeBotRegionsStatRequest) {
	request = &DescribeBotRegionsStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRegionsStat")
	return
}

func NewDescribeBotRegionsStatResponse() (response *DescribeBotRegionsStatResponse) {
	response = &DescribeBotRegionsStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot地理纬度统计
func (c *Client) DescribeBotRegionsStat(request *DescribeBotRegionsStatRequest) (response *DescribeBotRegionsStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotRegionsStatRequest()
	}
	response = NewDescribeBotRegionsStatResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOwaspWhiteRuleRequest() (request *DeleteOwaspWhiteRuleRequest) {
	request = &DeleteOwaspWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteOwaspWhiteRule")
	return
}

func NewDeleteOwaspWhiteRuleResponse() (response *DeleteOwaspWhiteRuleResponse) {
	response = &DeleteOwaspWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户规则引擎白名单
func (c *Client) DeleteOwaspWhiteRule(request *DeleteOwaspWhiteRuleRequest) (response *DeleteOwaspWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteOwaspWhiteRuleRequest()
	}
	response = NewDeleteOwaspWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMStatusRequest() (request *DescribeLLMStatusRequest) {
	request = &DescribeLLMStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMStatus")
	return
}

func NewDescribeLLMStatusResponse() (response *DescribeLLMStatusResponse) {
	response = &DescribeLLMStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取大模型开关状态
func (c *Client) DescribeLLMStatus(request *DescribeLLMStatusRequest) (response *DescribeLLMStatusResponse, err error) {
	if request == nil {
		request = NewDescribeLLMStatusRequest()
	}
	response = NewDescribeLLMStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJsInjectRuleStatusRequest() (request *ModifyJsInjectRuleStatusRequest) {
	request = &ModifyJsInjectRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyJsInjectRuleStatus")
	return
}

func NewModifyJsInjectRuleStatusResponse() (response *ModifyJsInjectRuleStatusResponse) {
	response = &ModifyJsInjectRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新前端对抗规则的状态，可以开启和关闭动态前端对抗规则。
func (c *Client) ModifyJsInjectRuleStatus(request *ModifyJsInjectRuleStatusRequest) (response *ModifyJsInjectRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyJsInjectRuleStatusRequest()
	}
	response = NewModifyJsInjectRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateInstanceLimitRulesRequest() (request *UpdateInstanceLimitRulesRequest) {
	request = &UpdateInstanceLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpdateInstanceLimitRules")
	return
}

func NewUpdateInstanceLimitRulesResponse() (response *UpdateInstanceLimitRulesResponse) {
	response = &UpdateInstanceLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实例限流规则
func (c *Client) UpdateInstanceLimitRules(request *UpdateInstanceLimitRulesRequest) (response *UpdateInstanceLimitRulesResponse, err error) {
	if request == nil {
		request = NewUpdateInstanceLimitRulesRequest()
	}
	response = NewUpdateInstanceLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotGlobalUAStrategyRequest() (request *DescribeBotGlobalUAStrategyRequest) {
	request = &DescribeBotGlobalUAStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotGlobalUAStrategy")
	return
}

func NewDescribeBotGlobalUAStrategyResponse() (response *DescribeBotGlobalUAStrategyResponse) {
	response = &DescribeBotGlobalUAStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Bot UA策略全局配置
func (c *Client) DescribeBotGlobalUAStrategy(request *DescribeBotGlobalUAStrategyRequest) (response *DescribeBotGlobalUAStrategyResponse, err error) {
	if request == nil {
		request = NewDescribeBotGlobalUAStrategyRequest()
	}
	response = NewDescribeBotGlobalUAStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStatusInnerRequest() (request *ModifyHostStatusInnerRequest) {
	request = &ModifyHostStatusInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostStatusInner")
	return
}

func NewModifyHostStatusInnerResponse() (response *ModifyHostStatusInnerResponse) {
	response = &ModifyHostStatusInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名WAF开关
// 支持批量操作。
func (c *Client) ModifyHostStatusInner(request *ModifyHostStatusInnerRequest) (response *ModifyHostStatusInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostStatusInnerRequest()
	}
	response = NewModifyHostStatusInnerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOwaspRuleStatusRequest() (request *ModifyOwaspRuleStatusRequest) {
	request = &ModifyOwaspRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOwaspRuleStatus")
	return
}

func NewModifyOwaspRuleStatusResponse() (response *ModifyOwaspRuleStatusResponse) {
	response = &ModifyOwaspRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新规则的开关
func (c *Client) ModifyOwaspRuleStatus(request *ModifyOwaspRuleStatusRequest) (response *ModifyOwaspRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyOwaspRuleStatusRequest()
	}
	response = NewModifyOwaspRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessLogsRequest() (request *DescribeAccessLogsRequest) {
	request = &DescribeAccessLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessLogs")
	return
}

func NewDescribeAccessLogsResponse() (response *DescribeAccessLogsResponse) {
	response = &DescribeAccessLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问日志
func (c *Client) DescribeAccessLogs(request *DescribeAccessLogsRequest) (response *DescribeAccessLogsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessLogsRequest()
	}
	response = NewDescribeAccessLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecurityFirstPurchaseRequest() (request *DescribeApiSecurityFirstPurchaseRequest) {
	request = &DescribeApiSecurityFirstPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecurityFirstPurchase")
	return
}

func NewDescribeApiSecurityFirstPurchaseResponse() (response *DescribeApiSecurityFirstPurchaseResponse) {
	response = &DescribeApiSecurityFirstPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否为第一次购买api安全资源，api接口会调用计费侧查询订单接口
func (c *Client) DescribeApiSecurityFirstPurchase(request *DescribeApiSecurityFirstPurchaseRequest) (response *DescribeApiSecurityFirstPurchaseResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecurityFirstPurchaseRequest()
	}
	response = NewDescribeApiSecurityFirstPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomRuleRequest() (request *AddCustomRuleRequest) {
	request = &AddCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddCustomRule")
	return
}

func NewAddCustomRuleResponse() (response *AddCustomRuleResponse) {
	response = &AddCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加访问控制（自定义策略）
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
	if request == nil {
		request = NewAddCustomRuleRequest()
	}
	response = NewAddCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddBatchAreaBanRequest() (request *AddBatchAreaBanRequest) {
	request = &AddBatchAreaBanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBatchAreaBan")
	return
}

func NewAddBatchAreaBanResponse() (response *AddBatchAreaBanResponse) {
	response = &AddBatchAreaBanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-批量新增接口
func (c *Client) AddBatchAreaBan(request *AddBatchAreaBanRequest) (response *AddBatchAreaBanResponse, err error) {
	if request == nil {
		request = NewAddBatchAreaBanRequest()
	}
	response = NewAddBatchAreaBanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletTesterWXTemplateRequest() (request *DescribeAppletTesterWXTemplateRequest) {
	request = &DescribeAppletTesterWXTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletTesterWXTemplate")
	return
}

func NewDescribeAppletTesterWXTemplateResponse() (response *DescribeAppletTesterWXTemplateResponse) {
	response = &DescribeAppletTesterWXTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序压测微信请求模版
func (c *Client) DescribeAppletTesterWXTemplate(request *DescribeAppletTesterWXTemplateRequest) (response *DescribeAppletTesterWXTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeAppletTesterWXTemplateRequest()
	}
	response = NewDescribeAppletTesterWXTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpAttackWhiteRuleNewRequest() (request *DescribeOpAttackWhiteRuleNewRequest) {
	request = &DescribeOpAttackWhiteRuleNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpAttackWhiteRuleNew")
	return
}

func NewDescribeOpAttackWhiteRuleNewResponse() (response *DescribeOpAttackWhiteRuleNewResponse) {
	response = &DescribeOpAttackWhiteRuleNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpAttackWhiteRuleNew(request *DescribeOpAttackWhiteRuleNewRequest) (response *DescribeOpAttackWhiteRuleNewResponse, err error) {
	if request == nil {
		request = NewDescribeOpAttackWhiteRuleNewRequest()
	}
	response = NewDescribeOpAttackWhiteRuleNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWxUserInfoRequest() (request *DescribeWxUserInfoRequest) {
	request = &DescribeWxUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWxUserInfo")
	return
}

func NewDescribeWxUserInfoResponse() (response *DescribeWxUserInfoResponse) {
	response = &DescribeWxUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微信用户信息
func (c *Client) DescribeWxUserInfo(request *DescribeWxUserInfoRequest) (response *DescribeWxUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWxUserInfoRequest()
	}
	response = NewDescribeWxUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJsInjectRuleRequest() (request *ModifyJsInjectRuleRequest) {
	request = &ModifyJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyJsInjectRule")
	return
}

func NewModifyJsInjectRuleResponse() (response *ModifyJsInjectRuleResponse) {
	response = &ModifyJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新域名的前端对抗规则
func (c *Client) ModifyJsInjectRule(request *ModifyJsInjectRuleRequest) (response *ModifyJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewModifyJsInjectRuleRequest()
	}
	response = NewModifyJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserSignaturePolicyStatusRequest() (request *GetUserSignaturePolicyStatusRequest) {
	request = &GetUserSignaturePolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetUserSignaturePolicyStatus")
	return
}

func NewGetUserSignaturePolicyStatusResponse() (response *GetUserSignaturePolicyStatusResponse) {
	response = &GetUserSignaturePolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防护规则开关
func (c *Client) GetUserSignaturePolicyStatus(request *GetUserSignaturePolicyStatusRequest) (response *GetUserSignaturePolicyStatusResponse, err error) {
	if request == nil {
		request = NewGetUserSignaturePolicyStatusRequest()
	}
	response = NewGetUserSignaturePolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostAccessLogStatusRequest() (request *ModifyHostAccessLogStatusRequest) {
	request = &ModifyHostAccessLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostAccessLogStatus")
	return
}

func NewModifyHostAccessLogStatusResponse() (response *ModifyHostAccessLogStatusResponse) {
	response = &ModifyHostAccessLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置单个域名是否开启记录访问日志功能的开关状态
func (c *Client) ModifyHostAccessLogStatus(request *ModifyHostAccessLogStatusRequest) (response *ModifyHostAccessLogStatusResponse, err error) {
	if request == nil {
		request = NewModifyHostAccessLogStatusRequest()
	}
	response = NewModifyHostAccessLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageExpireRequest() (request *SetLocalStorageExpireRequest) {
	request = &SetLocalStorageExpireRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SetLocalStorageExpire")
	return
}

func NewSetLocalStorageExpireResponse() (response *SetLocalStorageExpireResponse) {
	response = &SetLocalStorageExpireResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置远程存储缓存失效时间
func (c *Client) SetLocalStorageExpire(request *SetLocalStorageExpireRequest) (response *SetLocalStorageExpireResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageExpireRequest()
	}
	response = NewSetLocalStorageExpireResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchTrafficMarkingRequest() (request *SwitchTrafficMarkingRequest) {
	request = &SwitchTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchTrafficMarking")
	return
}

func NewSwitchTrafficMarkingResponse() (response *SwitchTrafficMarkingResponse) {
	response = &SwitchTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换流量标记开关
func (c *Client) SwitchTrafficMarking(request *SwitchTrafficMarkingRequest) (response *SwitchTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewSwitchTrafficMarkingRequest()
	}
	response = NewSwitchTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticTypesRequest() (request *DescribeStatisticTypesRequest) {
	request = &DescribeStatisticTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeStatisticTypes")
	return
}

func NewDescribeStatisticTypesResponse() (response *DescribeStatisticTypesResponse) {
	response = &DescribeStatisticTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取攻击类型统计数据
func (c *Client) DescribeStatisticTypes(request *DescribeStatisticTypesRequest) (response *DescribeStatisticTypesResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticTypesRequest()
	}
	response = NewDescribeStatisticTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllDomainRequest() (request *DescribeAllDomainRequest) {
	request = &DescribeAllDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAllDomain")
	return
}

func NewDescribeAllDomainResponse() (response *DescribeAllDomainResponse) {
	response = &DescribeAllDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clbwaf和saaswaf所有可用域名列表及域名cls开启状态
func (c *Client) DescribeAllDomain(request *DescribeAllDomainRequest) (response *DescribeAllDomainResponse, err error) {
	if request == nil {
		request = NewDescribeAllDomainRequest()
	}
	response = NewDescribeAllDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletDomainAccessAddressRequest() (request *DescribeAppletDomainAccessAddressRequest) {
	request = &DescribeAppletDomainAccessAddressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletDomainAccessAddress")
	return
}

func NewDescribeAppletDomainAccessAddressResponse() (response *DescribeAppletDomainAccessAddressResponse) {
	response = &DescribeAppletDomainAccessAddressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询小程序回源信息
func (c *Client) DescribeAppletDomainAccessAddress(request *DescribeAppletDomainAccessAddressRequest) (response *DescribeAppletDomainAccessAddressResponse, err error) {
	if request == nil {
		request = NewDescribeAppletDomainAccessAddressRequest()
	}
	response = NewDescribeAppletDomainAccessAddressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotActionRuleRequest() (request *DescribeBotActionRuleRequest) {
	request = &DescribeBotActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotActionRule")
	return
}

func NewDescribeBotActionRuleResponse() (response *DescribeBotActionRuleResponse) {
	response = &DescribeBotActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot用户配置的动作规则分布
func (c *Client) DescribeBotActionRule(request *DescribeBotActionRuleRequest) (response *DescribeBotActionRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotActionRuleRequest()
	}
	response = NewDescribeBotActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostFlowModeRequest() (request *ModifyHostFlowModeRequest) {
	request = &ModifyHostFlowModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostFlowMode")
	return
}

func NewModifyHostFlowModeResponse() (response *ModifyHostFlowModeResponse) {
	response = &ModifyHostFlowModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名的流量模式
func (c *Client) ModifyHostFlowMode(request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
	if request == nil {
		request = NewModifyHostFlowModeRequest()
	}
	response = NewModifyHostFlowModeResponse()
	err = c.Send(request, response)
	return
}

func NewAddPocTestTaskRequest() (request *AddPocTestTaskRequest) {
	request = &AddPocTestTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddPocTestTask")
	return
}

func NewAddPocTestTaskResponse() (response *AddPocTestTaskResponse) {
	response = &AddPocTestTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建POC测试任务
func (c *Client) AddPocTestTask(request *AddPocTestTaskRequest) (response *AddPocTestTaskResponse, err error) {
	if request == nil {
		request = NewAddPocTestTaskRequest()
	}
	response = NewAddPocTestTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessCursorRequest() (request *DescribeAccessCursorRequest) {
	request = &DescribeAccessCursorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessCursor")
	return
}

func NewDescribeAccessCursorResponse() (response *DescribeAccessCursorResponse) {
	response = &DescribeAccessCursorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据指定的时间来获取访问日志的游标
func (c *Client) DescribeAccessCursor(request *DescribeAccessCursorRequest) (response *DescribeAccessCursorResponse, err error) {
	if request == nil {
		request = NewDescribeAccessCursorRequest()
	}
	response = NewDescribeAccessCursorResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotGlobalUAStrategyRequest() (request *ModifyBotGlobalUAStrategyRequest) {
	request = &ModifyBotGlobalUAStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotGlobalUAStrategy")
	return
}

func NewModifyBotGlobalUAStrategyResponse() (response *ModifyBotGlobalUAStrategyResponse) {
	response = &ModifyBotGlobalUAStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Bot UA策略全局配置
func (c *Client) ModifyBotGlobalUAStrategy(request *ModifyBotGlobalUAStrategyRequest) (response *ModifyBotGlobalUAStrategyResponse, err error) {
	if request == nil {
		request = NewModifyBotGlobalUAStrategyRequest()
	}
	response = NewModifyBotGlobalUAStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafThreatenIntelligenceRequest() (request *ModifyWafThreatenIntelligenceRequest) {
	request = &ModifyWafThreatenIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWafThreatenIntelligence")
	return
}

func NewModifyWafThreatenIntelligenceResponse() (response *ModifyWafThreatenIntelligenceResponse) {
	response = &ModifyWafThreatenIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置WAF威胁情报封禁模块详情
func (c *Client) ModifyWafThreatenIntelligence(request *ModifyWafThreatenIntelligenceRequest) (response *ModifyWafThreatenIntelligenceResponse, err error) {
	if request == nil {
		request = NewModifyWafThreatenIntelligenceRequest()
	}
	response = NewModifyWafThreatenIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAIModelStatusRequest() (request *DescribeAIModelStatusRequest) {
	request = &DescribeAIModelStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAIModelStatus")
	return
}

func NewDescribeAIModelStatusResponse() (response *DescribeAIModelStatusResponse) {
	response = &DescribeAIModelStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ai模型状态
func (c *Client) DescribeAIModelStatus(request *DescribeAIModelStatusRequest) (response *DescribeAIModelStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAIModelStatusRequest()
	}
	response = NewDescribeAIModelStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDNSDetectDomainRequest() (request *CreateDNSDetectDomainRequest) {
	request = &CreateDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateDNSDetectDomain")
	return
}

func NewCreateDNSDetectDomainResponse() (response *CreateDNSDetectDomainResponse) {
	response = &CreateDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【添加域名】这个功能。
func (c *Client) CreateDNSDetectDomain(request *CreateDNSDetectDomainRequest) (response *CreateDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewCreateDNSDetectDomainRequest()
	}
	response = NewCreateDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCLSPackageRequest() (request *DescribeCLSPackageRequest) {
	request = &DescribeCLSPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCLSPackage")
	return
}

func NewDescribeCLSPackageResponse() (response *DescribeCLSPackageResponse) {
	response = &DescribeCLSPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的CLS信息
func (c *Client) DescribeCLSPackage(request *DescribeCLSPackageRequest) (response *DescribeCLSPackageResponse, err error) {
	if request == nil {
		request = NewDescribeCLSPackageRequest()
	}
	response = NewDescribeCLSPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainDetailsClbRequest() (request *DescribeDomainDetailsClbRequest) {
	request = &DescribeDomainDetailsClbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainDetailsClb")
	return
}

func NewDescribeDomainDetailsClbResponse() (response *DescribeDomainDetailsClbResponse) {
	response = &DescribeDomainDetailsClbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一个clbwaf域名详情
func (c *Client) DescribeDomainDetailsClb(request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
	if request == nil {
		request = NewDescribeDomainDetailsClbRequest()
	}
	response = NewDescribeDomainDetailsClbResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSessionKeyRequest() (request *ModifyBotSessionKeyRequest) {
	request = &ModifyBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotSessionKey")
	return
}

func NewModifyBotSessionKeyResponse() (response *ModifyBotSessionKeyResponse) {
	response = &ModifyBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot session 配置
func (c *Client) ModifyBotSessionKey(request *ModifyBotSessionKeyRequest) (response *ModifyBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewModifyBotSessionKeyRequest()
	}
	response = NewModifyBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeElasticFlowRequest() (request *DescribeElasticFlowRequest) {
	request = &DescribeElasticFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeElasticFlow")
	return
}

func NewDescribeElasticFlowResponse() (response *DescribeElasticFlowResponse) {
	response = &DescribeElasticFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务扩展包黑白名单
func (c *Client) DescribeElasticFlow(request *DescribeElasticFlowRequest) (response *DescribeElasticFlowResponse, err error) {
	if request == nil {
		request = NewDescribeElasticFlowRequest()
	}
	response = NewDescribeElasticFlowResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStatusRequest() (request *ModifyHostStatusRequest) {
	request = &ModifyHostStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostStatus")
	return
}

func NewModifyHostStatusResponse() (response *ModifyHostStatusResponse) {
	response = &ModifyHostStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf 设置防护域名WAF开关
// 支持批量操作。
func (c *Client) ModifyHostStatus(request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
	if request == nil {
		request = NewModifyHostStatusRequest()
	}
	response = NewModifyHostStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiStatPointsRequest() (request *DescribeApiStatPointsRequest) {
	request = &DescribeApiStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiStatPoints")
	return
}

func NewDescribeApiStatPointsResponse() (response *DescribeApiStatPointsResponse) {
	response = &DescribeApiStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api的访问趋势图
func (c *Client) DescribeApiStatPoints(request *DescribeApiStatPointsRequest) (response *DescribeApiStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeApiStatPointsRequest()
	}
	response = NewDescribeApiStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpartaProtectionsAutoRequest() (request *AddSpartaProtectionsAutoRequest) {
	request = &AddSpartaProtectionsAutoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddSpartaProtectionsAuto")
	return
}

func NewAddSpartaProtectionsAutoResponse() (response *AddSpartaProtectionsAutoResponse) {
	response = &AddSpartaProtectionsAutoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加域名
func (c *Client) AddSpartaProtectionsAuto(request *AddSpartaProtectionsAutoRequest) (response *AddSpartaProtectionsAutoResponse, err error) {
	if request == nil {
		request = NewAddSpartaProtectionsAutoRequest()
	}
	response = NewAddSpartaProtectionsAutoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletTesterReserveInfoRequest() (request *DescribeAppletTesterReserveInfoRequest) {
	request = &DescribeAppletTesterReserveInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletTesterReserveInfo")
	return
}

func NewDescribeAppletTesterReserveInfoResponse() (response *DescribeAppletTesterReserveInfoResponse) {
	response = &DescribeAppletTesterReserveInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序预约任务详情
func (c *Client) DescribeAppletTesterReserveInfo(request *DescribeAppletTesterReserveInfoRequest) (response *DescribeAppletTesterReserveInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAppletTesterReserveInfoRequest()
	}
	response = NewDescribeAppletTesterReserveInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHistogramRequest() (request *DescribeHistogramRequest) {
	request = &DescribeHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHistogram")
	return
}

func NewDescribeHistogramResponse() (response *DescribeHistogramResponse) {
	response = &DescribeHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询多种条件的聚类分析
func (c *Client) DescribeHistogram(request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeHistogramRequest()
	}
	response = NewDescribeHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpAttackWhiteRuleRequest() (request *DescribeOpAttackWhiteRuleRequest) {
	request = &DescribeOpAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpAttackWhiteRule")
	return
}

func NewDescribeOpAttackWhiteRuleResponse() (response *DescribeOpAttackWhiteRuleResponse) {
	response = &DescribeOpAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpAttackWhiteRule(request *DescribeOpAttackWhiteRuleRequest) (response *DescribeOpAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpAttackWhiteRuleRequest()
	}
	response = NewDescribeOpAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotThreatIntelligenceRequest() (request *DescribeBotThreatIntelligenceRequest) {
	request = &DescribeBotThreatIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotThreatIntelligence")
	return
}

func NewDescribeBotThreatIntelligenceResponse() (response *DescribeBotThreatIntelligenceResponse) {
	response = &DescribeBotThreatIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的威胁情报配置列表
func (c *Client) DescribeBotThreatIntelligence(request *DescribeBotThreatIntelligenceRequest) (response *DescribeBotThreatIntelligenceResponse, err error) {
	if request == nil {
		request = NewDescribeBotThreatIntelligenceRequest()
	}
	response = NewDescribeBotThreatIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewAddAntiFakeUrlRequest() (request *AddAntiFakeUrlRequest) {
	request = &AddAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddAntiFakeUrl")
	return
}

func NewAddAntiFakeUrlResponse() (response *AddAntiFakeUrlResponse) {
	response = &AddAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加防篡改url
func (c *Client) AddAntiFakeUrl(request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewAddAntiFakeUrlRequest()
	}
	response = NewAddAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotPointsRequest() (request *DescribeBotPointsRequest) {
	request = &DescribeBotPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotPoints")
	return
}

func NewDescribeBotPointsResponse() (response *DescribeBotPointsResponse) {
	response = &DescribeBotPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Bot趋势数据
func (c *Client) DescribeBotPoints(request *DescribeBotPointsRequest) (response *DescribeBotPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotPointsRequest()
	}
	response = NewDescribeBotPointsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackDownloadRecordsRequest() (request *GetAttackDownloadRecordsRequest) {
	request = &GetAttackDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackDownloadRecords")
	return
}

func NewGetAttackDownloadRecordsResponse() (response *GetAttackDownloadRecordsResponse) {
	response = &GetAttackDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载攻击日志任务记录列表
func (c *Client) GetAttackDownloadRecords(request *GetAttackDownloadRecordsRequest) (response *GetAttackDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewGetAttackDownloadRecordsRequest()
	}
	response = NewGetAttackDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebshellStatusRequest() (request *ModifyWebshellStatusRequest) {
	request = &ModifyWebshellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWebshellStatus")
	return
}

func NewModifyWebshellStatusResponse() (response *ModifyWebshellStatusResponse) {
	response = &ModifyWebshellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置域名的webshell状态。
func (c *Client) ModifyWebshellStatus(request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
	if request == nil {
		request = NewModifyWebshellStatusRequest()
	}
	response = NewModifyWebshellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBypassAllStatusRequest() (request *QueryBypassAllStatusRequest) {
	request = &QueryBypassAllStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryBypassAllStatus")
	return
}

func NewQueryBypassAllStatusResponse() (response *QueryBypassAllStatusResponse) {
	response = &QueryBypassAllStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询该用户是否被加入了全局的bypass列表
func (c *Client) QueryBypassAllStatus(request *QueryBypassAllStatusRequest) (response *QueryBypassAllStatusResponse, err error) {
	if request == nil {
		request = NewQueryBypassAllStatusRequest()
	}
	response = NewQueryBypassAllStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDomainPostArgsListRequest() (request *QueryDomainPostArgsListRequest) {
	request = &QueryDomainPostArgsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryDomainPostArgsList")
	return
}

func NewQueryDomainPostArgsListResponse() (response *QueryDomainPostArgsListResponse) {
	response = &QueryDomainPostArgsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [WAF请求体参数长度下发]-编辑接口
func (c *Client) QueryDomainPostArgsList(request *QueryDomainPostArgsListRequest) (response *QueryDomainPostArgsListResponse, err error) {
	if request == nil {
		request = NewQueryDomainPostArgsListRequest()
	}
	response = NewQueryDomainPostArgsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRequestLogRequest() (request *DescribeBotRequestLogRequest) {
	request = &DescribeBotRequestLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRequestLog")
	return
}

func NewDescribeBotRequestLogResponse() (response *DescribeBotRequestLogResponse) {
	response = &DescribeBotRequestLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot攻击日志
func (c *Client) DescribeBotRequestLog(request *DescribeBotRequestLogRequest) (response *DescribeBotRequestLogResponse, err error) {
	if request == nil {
		request = NewDescribeBotRequestLogRequest()
	}
	response = NewDescribeBotRequestLogResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogRequest() (request *EditOpRuleUpdateLogRequest) {
	request = &EditOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EditOpRuleUpdateLog")
	return
}

func NewEditOpRuleUpdateLogResponse() (response *EditOpRuleUpdateLogResponse) {
	response = &EditOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志
func (c *Client) EditOpRuleUpdateLog(request *EditOpRuleUpdateLogRequest) (response *EditOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogRequest()
	}
	response = NewEditOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSceneJsInjectRuleRequest() (request *DescribeSceneJsInjectRuleRequest) {
	request = &DescribeSceneJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSceneJsInjectRule")
	return
}

func NewDescribeSceneJsInjectRuleResponse() (response *DescribeSceneJsInjectRuleResponse) {
	response = &DescribeSceneJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】获取对应域名的前端对抗规则，查询全局的前端对抗配置时IsGlobal==1，返回的JsInjectRule中只含有全局设置对应的字段；查询场景的前端对抗配置时IsGlobal==0，返回的JsInjectRule中只含有场景设置对应的字段。
func (c *Client) DescribeSceneJsInjectRule(request *DescribeSceneJsInjectRuleRequest) (response *DescribeSceneJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewDescribeSceneJsInjectRuleRequest()
	}
	response = NewDescribeSceneJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopAttackIPRequest() (request *GetTopAttackIPRequest) {
	request = &GetTopAttackIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetTopAttackIP")
	return
}

func NewGetTopAttackIPResponse() (response *GetTopAttackIPResponse) {
	response = &GetTopAttackIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意IP的top10
func (c *Client) GetTopAttackIP(request *GetTopAttackIPRequest) (response *GetTopAttackIPResponse, err error) {
	if request == nil {
		request = NewGetTopAttackIPRequest()
	}
	response = NewGetTopAttackIPResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaProtectionRequest() (request *ModifySpartaProtectionRequest) {
	request = &ModifySpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifySpartaProtection")
	return
}

func NewModifySpartaProtectionResponse() (response *ModifySpartaProtectionResponse) {
	response = &ModifySpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名配置
func (c *Client) ModifySpartaProtection(request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
	if request == nil {
		request = NewModifySpartaProtectionRequest()
	}
	response = NewModifySpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalWhiteRuleRequest() (request *DescribeGlobalWhiteRuleRequest) {
	request = &DescribeGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeGlobalWhiteRule")
	return
}

func NewDescribeGlobalWhiteRuleResponse() (response *DescribeGlobalWhiteRuleResponse) {
	response = &DescribeGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，获取全局白名单列表。
func (c *Client) DescribeGlobalWhiteRule(request *DescribeGlobalWhiteRuleRequest) (response *DescribeGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalWhiteRuleRequest()
	}
	response = NewDescribeGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainPKIValidationRequest() (request *ModifyDomainPKIValidationRequest) {
	request = &ModifyDomainPKIValidationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainPKIValidation")
	return
}

func NewModifyDomainPKIValidationResponse() (response *ModifyDomainPKIValidationResponse) {
	response = &ModifyDomainPKIValidationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云SSL调用，用于修改域名的PKI验证路径
func (c *Client) ModifyDomainPKIValidation(request *ModifyDomainPKIValidationRequest) (response *ModifyDomainPKIValidationResponse, err error) {
	if request == nil {
		request = NewModifyDomainPKIValidationRequest()
	}
	response = NewModifyDomainPKIValidationResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRateLimitsV2Request() (request *EnableRateLimitsV2Request) {
	request = &EnableRateLimitsV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EnableRateLimitsV2")
	return
}

func NewEnableRateLimitsV2Response() (response *EnableRateLimitsV2Response) {
	response = &EnableRateLimitsV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量更改自研版限流规则开关
func (c *Client) EnableRateLimitsV2(request *EnableRateLimitsV2Request) (response *EnableRateLimitsV2Response, err error) {
	if request == nil {
		request = NewEnableRateLimitsV2Request()
	}
	response = NewEnableRateLimitsV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBotIdTypeRequest() (request *DescribeBotIdTypeRequest) {
	request = &DescribeBotIdTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotIdType")
	return
}

func NewDescribeBotIdTypeResponse() (response *DescribeBotIdTypeResponse) {
	response = &DescribeBotIdTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot-id规则类型
func (c *Client) DescribeBotIdType(request *DescribeBotIdTypeRequest) (response *DescribeBotIdTypeResponse, err error) {
	if request == nil {
		request = NewDescribeBotIdTypeRequest()
	}
	response = NewDescribeBotIdTypeResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchApiRulesRequest() (request *SwitchApiRulesRequest) {
	request = &SwitchApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchApiRules")
	return
}

func NewSwitchApiRulesResponse() (response *SwitchApiRulesResponse) {
	response = &SwitchApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用或者停用Api规则
func (c *Client) SwitchApiRules(request *SwitchApiRulesRequest) (response *SwitchApiRulesResponse, err error) {
	if request == nil {
		request = NewSwitchApiRulesRequest()
	}
	response = NewSwitchApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectHijackDataRequest() (request *DescribeDNSDetectHijackDataRequest) {
	request = &DescribeDNSDetectHijackDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDNSDetectHijackData")
	return
}

func NewDescribeDNSDetectHijackDataResponse() (response *DescribeDNSDetectHijackDataResponse) {
	response = &DescribeDNSDetectHijackDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的劫持记录，对应【Web安全检测】【DNS劫持检测】【劫持记录】这个功能。
func (c *Client) DescribeDNSDetectHijackData(request *DescribeDNSDetectHijackDataRequest) (response *DescribeDNSDetectHijackDataResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectHijackDataRequest()
	}
	response = NewDescribeDNSDetectHijackDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRulesInfoRequest() (request *DescribeRulesInfoRequest) {
	request = &DescribeRulesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRulesInfo")
	return
}

func NewDescribeRulesInfoResponse() (response *DescribeRulesInfoResponse) {
	response = &DescribeRulesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则库的描述信息
func (c *Client) DescribeRulesInfo(request *DescribeRulesInfoRequest) (response *DescribeRulesInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRulesInfoRequest()
	}
	response = NewDescribeRulesInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRateLimitsRequest() (request *CreateRateLimitsRequest) {
	request = &CreateRateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateRateLimits")
	return
}

func NewCreateRateLimitsResponse() (response *CreateRateLimitsResponse) {
	response = &CreateRateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建限流规则
func (c *Client) CreateRateLimits(request *CreateRateLimitsRequest) (response *CreateRateLimitsResponse, err error) {
	if request == nil {
		request = NewCreateRateLimitsRequest()
	}
	response = NewCreateRateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomPayloadRequest() (request *AddCustomPayloadRequest) {
	request = &AddCustomPayloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddCustomPayload")
	return
}

func NewAddCustomPayloadResponse() (response *AddCustomPayloadResponse) {
	response = &AddCustomPayloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义载荷。
func (c *Client) AddCustomPayload(request *AddCustomPayloadRequest) (response *AddCustomPayloadResponse, err error) {
	if request == nil {
		request = NewAddCustomPayloadRequest()
	}
	response = NewAddCustomPayloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDomainInfoRequest() (request *DescribeUserDomainInfoRequest) {
	request = &DescribeUserDomainInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserDomainInfo")
	return
}

func NewDescribeUserDomainInfoResponse() (response *DescribeUserDomainInfoResponse) {
	response = &DescribeUserDomainInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询saas和clb的域名信息
func (c *Client) DescribeUserDomainInfo(request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserDomainInfoRequest()
	}
	response = NewDescribeUserDomainInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyCustomWhiteRuleRequest() (request *CreateCopyCustomWhiteRuleRequest) {
	request = &CreateCopyCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateCopyCustomWhiteRule")
	return
}

func NewCreateCopyCustomWhiteRuleResponse() (response *CreateCopyCustomWhiteRuleResponse) {
	response = &CreateCopyCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制精准白名单到其他域名
func (c *Client) CreateCopyCustomWhiteRule(request *CreateCopyCustomWhiteRuleRequest) (response *CreateCopyCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewCreateCopyCustomWhiteRuleRequest()
	}
	response = NewCreateCopyCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessStatPointsRequest() (request *DescribeBotAccessStatPointsRequest) {
	request = &DescribeBotAccessStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAccessStatPoints")
	return
}

func NewDescribeBotAccessStatPointsResponse() (response *DescribeBotAccessStatPointsResponse) {
	response = &DescribeBotAccessStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot4.0的趋势图接口
func (c *Client) DescribeBotAccessStatPoints(request *DescribeBotAccessStatPointsRequest) (response *DescribeBotAccessStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessStatPointsRequest()
	}
	response = NewDescribeBotAccessStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotItemsRequest() (request *DescribeBotItemsRequest) {
	request = &DescribeBotItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotItems")
	return
}

func NewDescribeBotItemsResponse() (response *DescribeBotItemsResponse) {
	response = &DescribeBotItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志详情对应的字段
func (c *Client) DescribeBotItems(request *DescribeBotItemsRequest) (response *DescribeBotItemsResponse, err error) {
	if request == nil {
		request = NewDescribeBotItemsRequest()
	}
	response = NewDescribeBotItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSessionKeyRequest() (request *DescribeBotSessionKeyRequest) {
	request = &DescribeBotSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSessionKey")
	return
}

func NewDescribeBotSessionKeyResponse() (response *DescribeBotSessionKeyResponse) {
	response = &DescribeBotSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的session配置
func (c *Client) DescribeBotSessionKey(request *DescribeBotSessionKeyRequest) (response *DescribeBotSessionKeyResponse, err error) {
	if request == nil {
		request = NewDescribeBotSessionKeyRequest()
	}
	response = NewDescribeBotSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCachePathRequest() (request *DescribeCachePathRequest) {
	request = &DescribeCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCachePath")
	return
}

func NewDescribeCachePathResponse() (response *DescribeCachePathResponse) {
	response = &DescribeCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-查找缓存路径
func (c *Client) DescribeCachePath(request *DescribeCachePathRequest) (response *DescribeCachePathResponse, err error) {
	if request == nil {
		request = NewDescribeCachePathRequest()
	}
	response = NewDescribeCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomRuleListRequest() (request *DescribeCustomRuleListRequest) {
	request = &DescribeCustomRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomRuleList")
	return
}

func NewDescribeCustomRuleListResponse() (response *DescribeCustomRuleListResponse) {
	response = &DescribeCustomRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的访问控制策略列表
func (c *Client) DescribeCustomRuleList(request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeCustomRuleListRequest()
	}
	response = NewDescribeCustomRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolePlayStateRequest() (request *DescribeRolePlayStateRequest) {
	request = &DescribeRolePlayStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRolePlayState")
	return
}

func NewDescribeRolePlayStateResponse() (response *DescribeRolePlayStateResponse) {
	response = &DescribeRolePlayStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户授权的信息
func (c *Client) DescribeRolePlayState(request *DescribeRolePlayStateRequest) (response *DescribeRolePlayStateResponse, err error) {
	if request == nil {
		request = NewDescribeRolePlayStateRequest()
	}
	response = NewDescribeRolePlayStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRateLimitsV2Request() (request *DescribeRateLimitsV2Request) {
	request = &DescribeRateLimitsV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRateLimitsV2")
	return
}

func NewDescribeRateLimitsV2Response() (response *DescribeRateLimitsV2Response) {
	response = &DescribeRateLimitsV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询限流规则列表接口
func (c *Client) DescribeRateLimitsV2(request *DescribeRateLimitsV2Request) (response *DescribeRateLimitsV2Response, err error) {
	if request == nil {
		request = NewDescribeRateLimitsV2Request()
	}
	response = NewDescribeRateLimitsV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
	request = &DescribeUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserInfo")
	return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
	response = &DescribeUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clb-waf的套餐信息
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserInfoRequest()
	}
	response = NewDescribeUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiInfoLeakageRulesRequest() (request *DescribeAntiInfoLeakageRulesRequest) {
	request = &DescribeAntiInfoLeakageRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAntiInfoLeakageRules")
	return
}

func NewDescribeAntiInfoLeakageRulesResponse() (response *DescribeAntiInfoLeakageRulesResponse) {
	response = &DescribeAntiInfoLeakageRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取得信息防泄漏规则列表
func (c *Client) DescribeAntiInfoLeakageRules(request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiInfoLeakageRulesRequest()
	}
	response = NewDescribeAntiInfoLeakageRulesResponse()
	err = c.Send(request, response)
	return
}

func NewWafCreateResourceAfterPayRequest() (request *WafCreateResourceAfterPayRequest) {
	request = &WafCreateResourceAfterPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafCreateResourceAfterPay")
	return
}

func NewWafCreateResourceAfterPayResponse() (response *WafCreateResourceAfterPayResponse) {
	response = &WafCreateResourceAfterPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 购买WAF
func (c *Client) WafCreateResourceAfterPay(request *WafCreateResourceAfterPayRequest) (response *WafCreateResourceAfterPayResponse, err error) {
	if request == nil {
		request = NewWafCreateResourceAfterPayRequest()
	}
	response = NewWafCreateResourceAfterPayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostRequest() (request *DeleteHostRequest) {
	request = &DeleteHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteHost")
	return
}

func NewDeleteHostResponse() (response *DeleteHostResponse) {
	response = &DeleteHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CLB-WAF防护域名
// 支持批量操作
func (c *Client) DeleteHost(request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
	if request == nil {
		request = NewDeleteHostRequest()
	}
	response = NewDeleteHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotIdRuleRequest() (request *DescribeBotIdRuleRequest) {
	request = &DescribeBotIdRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotIdRule")
	return
}

func NewDescribeBotIdRuleResponse() (response *DescribeBotIdRuleResponse) {
	response = &DescribeBotIdRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BotId规则列表1
func (c *Client) DescribeBotIdRule(request *DescribeBotIdRuleRequest) (response *DescribeBotIdRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotIdRuleRequest()
	}
	response = NewDescribeBotIdRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotOverviewRequest() (request *DescribeBotOverviewRequest) {
	request = &DescribeBotOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotOverview")
	return
}

func NewDescribeBotOverviewResponse() (response *DescribeBotOverviewResponse) {
	response = &DescribeBotOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的概览数据
func (c *Client) DescribeBotOverview(request *DescribeBotOverviewRequest) (response *DescribeBotOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotOverviewRequest()
	}
	response = NewDescribeBotOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKeepProtectionConfRequest() (request *DescribeKeepProtectionConfRequest) {
	request = &DescribeKeepProtectionConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeKeepProtectionConf")
	return
}

func NewDescribeKeepProtectionConfResponse() (response *DescribeKeepProtectionConfResponse) {
	response = &DescribeKeepProtectionConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看用户是否开启防护配置自动删除开关，默认不开启
func (c *Client) DescribeKeepProtectionConf(request *DescribeKeepProtectionConfRequest) (response *DescribeKeepProtectionConfResponse, err error) {
	if request == nil {
		request = NewDescribeKeepProtectionConfRequest()
	}
	response = NewDescribeKeepProtectionConfResponse()
	err = c.Send(request, response)
	return
}

func NewPostAttackDownloadTaskRequest() (request *PostAttackDownloadTaskRequest) {
	request = &PostAttackDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "PostAttackDownloadTask")
	return
}

func NewPostAttackDownloadTaskResponse() (response *PostAttackDownloadTaskResponse) {
	response = &PostAttackDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建搜索下载攻击日志任务，使用CLS新版本的搜索下载getlog接口
func (c *Client) PostAttackDownloadTask(request *PostAttackDownloadTaskRequest) (response *PostAttackDownloadTaskResponse, err error) {
	if request == nil {
		request = NewPostAttackDownloadTaskRequest()
	}
	response = NewPostAttackDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHybridClusterRequest() (request *DescribeHybridClusterRequest) {
	request = &DescribeHybridClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHybridCluster")
	return
}

func NewDescribeHybridClusterResponse() (response *DescribeHybridClusterResponse) {
	response = &DescribeHybridClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看单个混合云集群信息
func (c *Client) DescribeHybridCluster(request *DescribeHybridClusterRequest) (response *DescribeHybridClusterResponse, err error) {
	if request == nil {
		request = NewDescribeHybridClusterRequest()
	}
	response = NewDescribeHybridClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCertificatedDomainRequest() (request *ModifyCertificatedDomainRequest) {
	request = &ModifyCertificatedDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCertificatedDomain")
	return
}

func NewModifyCertificatedDomainResponse() (response *ModifyCertificatedDomainResponse) {
	response = &ModifyCertificatedDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名的证书
func (c *Client) ModifyCertificatedDomain(request *ModifyCertificatedDomainRequest) (response *ModifyCertificatedDomainResponse, err error) {
	if request == nil {
		request = NewModifyCertificatedDomainRequest()
	}
	response = NewModifyCertificatedDomainResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshAccessCheckResultRequest() (request *RefreshAccessCheckResultRequest) {
	request = &RefreshAccessCheckResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RefreshAccessCheckResult")
	return
}

func NewRefreshAccessCheckResultResponse() (response *RefreshAccessCheckResultResponse) {
	response = &RefreshAccessCheckResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新接入检查的结果，后台会生成接入检查任务
func (c *Client) RefreshAccessCheckResult(request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
	if request == nil {
		request = NewRefreshAccessCheckResultRequest()
	}
	response = NewRefreshAccessCheckResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessDownloadRecordsRequest() (request *DescribeAccessDownloadRecordsRequest) {
	request = &DescribeAccessDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessDownloadRecords")
	return
}

func NewDescribeAccessDownloadRecordsResponse() (response *DescribeAccessDownloadRecordsResponse) {
	response = &DescribeAccessDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询以创建的访问日志下载任务
func (c *Client) DescribeAccessDownloadRecords(request *DescribeAccessDownloadRecordsRequest) (response *DescribeAccessDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessDownloadRecordsRequest()
	}
	response = NewDescribeAccessDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrdersAndPayRequest() (request *CreateOrdersAndPayRequest) {
	request = &CreateOrdersAndPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateOrdersAndPay")
	return
}

func NewCreateOrdersAndPayResponse() (response *CreateOrdersAndPayResponse) {
	response = &CreateOrdersAndPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费资源购买、续费下单接口
func (c *Client) CreateOrdersAndPay(request *CreateOrdersAndPayRequest) (response *CreateOrdersAndPayResponse, err error) {
	if request == nil {
		request = NewCreateOrdersAndPayRequest()
	}
	response = NewCreateOrdersAndPayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessDownloadRecordRequest() (request *DeleteAccessDownloadRecordRequest) {
	request = &DeleteAccessDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAccessDownloadRecord")
	return
}

func NewDeleteAccessDownloadRecordResponse() (response *DeleteAccessDownloadRecordResponse) {
	response = &DeleteAccessDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问日志下载任务记录
func (c *Client) DeleteAccessDownloadRecord(request *DeleteAccessDownloadRecordRequest) (response *DeleteAccessDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteAccessDownloadRecordRequest()
	}
	response = NewDeleteAccessDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebshellStatusRequest() (request *DescribeWebshellStatusRequest) {
	request = &DescribeWebshellStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWebshellStatus")
	return
}

func NewDescribeWebshellStatusResponse() (response *DescribeWebshellStatusResponse) {
	response = &DescribeWebshellStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名的webshell状态
func (c *Client) DescribeWebshellStatus(request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
	if request == nil {
		request = NewDescribeWebshellStatusRequest()
	}
	response = NewDescribeWebshellStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTiRuleRequest() (request *DescribeBotTiRuleRequest) {
	request = &DescribeBotTiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTiRule")
	return
}

func NewDescribeBotTiRuleResponse() (response *DescribeBotTiRuleResponse) {
	response = &DescribeBotTiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot模块的威胁情报配置
func (c *Client) DescribeBotTiRule(request *DescribeBotTiRuleRequest) (response *DescribeBotTiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotTiRuleRequest()
	}
	response = NewDescribeBotTiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSceneInfoRequest() (request *DescribeUserSceneInfoRequest) {
	request = &DescribeUserSceneInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserSceneInfo")
	return
}

func NewDescribeUserSceneInfoResponse() (response *DescribeUserSceneInfoResponse) {
	response = &DescribeUserSceneInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名下的场景and动作策略列表
func (c *Client) DescribeUserSceneInfo(request *DescribeUserSceneInfoRequest) (response *DescribeUserSceneInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserSceneInfoRequest()
	}
	response = NewDescribeUserSceneInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAttackWhiteRuleRequest() (request *ModifyAttackWhiteRuleRequest) {
	request = &ModifyAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAttackWhiteRule")
	return
}

func NewModifyAttackWhiteRuleResponse() (response *ModifyAttackWhiteRuleResponse) {
	response = &ModifyAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，修改Tiga规则引擎白名单。
func (c *Client) ModifyAttackWhiteRule(request *ModifyAttackWhiteRuleRequest) (response *ModifyAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyAttackWhiteRuleRequest()
	}
	response = NewModifyAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAccessLogRequest() (request *SearchAccessLogRequest) {
	request = &SearchAccessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SearchAccessLog")
	return
}

func NewSearchAccessLogResponse() (response *SearchAccessLogResponse) {
	response = &SearchAccessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于搜索WAF访问日志
func (c *Client) SearchAccessLog(request *SearchAccessLogRequest) (response *SearchAccessLogResponse, err error) {
	if request == nil {
		request = NewSearchAccessLogRequest()
	}
	response = NewSearchAccessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBypassSpartaProtectionRequest() (request *DescribeBypassSpartaProtectionRequest) {
	request = &DescribeBypassSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBypassSpartaProtection")
	return
}

func NewDescribeBypassSpartaProtectionResponse() (response *DescribeBypassSpartaProtectionResponse) {
	response = &DescribeBypassSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询bypass信息
func (c *Client) DescribeBypassSpartaProtection(request *DescribeBypassSpartaProtectionRequest) (response *DescribeBypassSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewDescribeBypassSpartaProtectionRequest()
	}
	response = NewDescribeBypassSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGlobalWhiteRuleRequest() (request *ModifyGlobalWhiteRuleRequest) {
	request = &ModifyGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyGlobalWhiteRule")
	return
}

func NewModifyGlobalWhiteRuleResponse() (response *ModifyGlobalWhiteRuleResponse) {
	response = &ModifyGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，增加规则引擎全局白名单列表。
func (c *Client) ModifyGlobalWhiteRule(request *ModifyGlobalWhiteRuleRequest) (response *ModifyGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyGlobalWhiteRuleRequest()
	}
	response = NewModifyGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceName")
	return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的名称
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBusinessRiskRequest() (request *DeleteBusinessRiskRequest) {
	request = &DeleteBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBusinessRisk")
	return
}

func NewDeleteBusinessRiskResponse() (response *DeleteBusinessRiskResponse) {
	response = &DeleteBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务安全策略
func (c *Client) DeleteBusinessRisk(request *DeleteBusinessRiskRequest) (response *DeleteBusinessRiskResponse, err error) {
	if request == nil {
		request = NewDeleteBusinessRiskRequest()
	}
	response = NewDeleteBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotEventHistoryRequest() (request *DescribeBotEventHistoryRequest) {
	request = &DescribeBotEventHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotEventHistory")
	return
}

func NewDescribeBotEventHistoryResponse() (response *DescribeBotEventHistoryResponse) {
	response = &DescribeBotEventHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-拉取BOT事件的变更历史
func (c *Client) DescribeBotEventHistory(request *DescribeBotEventHistoryRequest) (response *DescribeBotEventHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeBotEventHistoryRequest()
	}
	response = NewDescribeBotEventHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewQueryConfigDeliveryRequest() (request *QueryConfigDeliveryRequest) {
	request = &QueryConfigDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryConfigDelivery")
	return
}

func NewQueryConfigDeliveryResponse() (response *QueryConfigDeliveryResponse) {
	response = &QueryConfigDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [WAF运维管理系统]-查询配置下发接口
func (c *Client) QueryConfigDelivery(request *QueryConfigDeliveryRequest) (response *QueryConfigDeliveryResponse, err error) {
	if request == nil {
		request = NewQueryConfigDeliveryRequest()
	}
	response = NewQueryConfigDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWhiteAllDomainRequest() (request *DescribeWhiteAllDomainRequest) {
	request = &DescribeWhiteAllDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWhiteAllDomain")
	return
}

func NewDescribeWhiteAllDomainResponse() (response *DescribeWhiteAllDomainResponse) {
	response = &DescribeWhiteAllDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询saas和clb的域名信息v2
func (c *Client) DescribeWhiteAllDomain(request *DescribeWhiteAllDomainRequest) (response *DescribeWhiteAllDomainResponse, err error) {
	if request == nil {
		request = NewDescribeWhiteAllDomainRequest()
	}
	response = NewDescribeWhiteAllDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneRequest() (request *DeleteBotSceneRequest) {
	request = &DeleteBotSceneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBotScene")
	return
}

func NewDeleteBotSceneResponse() (response *DeleteBotSceneResponse) {
	response = &DeleteBotSceneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除BOT场景配置
func (c *Client) DeleteBotScene(request *DeleteBotSceneRequest) (response *DeleteBotSceneResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneRequest()
	}
	response = NewDeleteBotSceneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchIpPunishmentListRequest() (request *DescribeBatchIpPunishmentListRequest) {
	request = &DescribeBatchIpPunishmentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchIpPunishmentList")
	return
}

func NewDescribeBatchIpPunishmentListResponse() (response *DescribeBatchIpPunishmentListResponse) {
	response = &DescribeBatchIpPunishmentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]批量列表接口
func (c *Client) DescribeBatchIpPunishmentList(request *DescribeBatchIpPunishmentListRequest) (response *DescribeBatchIpPunishmentListResponse, err error) {
	if request == nil {
		request = NewDescribeBatchIpPunishmentListRequest()
	}
	response = NewDescribeBatchIpPunishmentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatPointsRequest() (request *DescribeBotStatPointsRequest) {
	request = &DescribeBotStatPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotStatPoints")
	return
}

func NewDescribeBotStatPointsResponse() (response *DescribeBotStatPointsResponse) {
	response = &DescribeBotStatPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的趋势图接口
func (c *Client) DescribeBotStatPoints(request *DescribeBotStatPointsRequest) (response *DescribeBotStatPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatPointsRequest()
	}
	response = NewDescribeBotStatPointsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMStatusRequest() (request *ModifyLLMStatusRequest) {
	request = &ModifyLLMStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMStatus")
	return
}

func NewModifyLLMStatusResponse() (response *ModifyLLMStatusResponse) {
	response = &ModifyLLMStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大模型开关更新
func (c *Client) ModifyLLMStatus(request *ModifyLLMStatusRequest) (response *ModifyLLMStatusResponse, err error) {
	if request == nil {
		request = NewModifyLLMStatusRequest()
	}
	response = NewModifyLLMStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePeakValueRequest() (request *DescribePeakValueRequest) {
	request = &DescribePeakValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePeakValue")
	return
}

func NewDescribePeakValueResponse() (response *DescribePeakValueResponse) {
	response = &DescribePeakValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取业务和攻击概览峰值
func (c *Client) DescribePeakValue(request *DescribePeakValueRequest) (response *DescribePeakValueResponse, err error) {
	if request == nil {
		request = NewDescribePeakValueRequest()
	}
	response = NewDescribePeakValueResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomRuleStatusRequest() (request *ModifyCustomRuleStatusRequest) {
	request = &ModifyCustomRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCustomRuleStatus")
	return
}

func NewModifyCustomRuleStatusResponse() (response *ModifyCustomRuleStatusResponse) {
	response = &ModifyCustomRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或禁用访问控制（自定义策略）
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyCustomRuleStatusRequest()
	}
	response = NewModifyCustomRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRuleRollbackRequest() (request *RuleRollbackRequest) {
	request = &RuleRollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RuleRollback")
	return
}

func NewRuleRollbackResponse() (response *RuleRollbackResponse) {
	response = &RuleRollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚tiga规则库
func (c *Client) RuleRollback(request *RuleRollbackRequest) (response *RuleRollbackResponse, err error) {
	if request == nil {
		request = NewRuleRollbackRequest()
	}
	response = NewRuleRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewCheckGlobalWhiteRuleNameRequest() (request *CheckGlobalWhiteRuleNameRequest) {
	request = &CheckGlobalWhiteRuleNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CheckGlobalWhiteRuleName")
	return
}

func NewCheckGlobalWhiteRuleNameResponse() (response *CheckGlobalWhiteRuleNameResponse) {
	response = &CheckGlobalWhiteRuleNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：检测全局白名单名字是否重复
func (c *Client) CheckGlobalWhiteRuleName(request *CheckGlobalWhiteRuleNameRequest) (response *CheckGlobalWhiteRuleNameResponse, err error) {
	if request == nil {
		request = NewCheckGlobalWhiteRuleNameRequest()
	}
	response = NewCheckGlobalWhiteRuleNameResponse()
	err = c.Send(request, response)
	return
}

func NewConvertStrategyToCELRequest() (request *ConvertStrategyToCELRequest) {
	request = &ConvertStrategyToCELRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ConvertStrategyToCEL")
	return
}

func NewConvertStrategyToCELResponse() (response *ConvertStrategyToCELResponse) {
	response = &ConvertStrategyToCELResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将自定义规则的匹配条件列表转换为Google的CEL表达式
func (c *Client) ConvertStrategyToCEL(request *ConvertStrategyToCELRequest) (response *ConvertStrategyToCELResponse, err error) {
	if request == nil {
		request = NewConvertStrategyToCELRequest()
	}
	response = NewConvertStrategyToCELResponse()
	err = c.Send(request, response)
	return
}

func NewAddBlockPageRequest() (request *AddBlockPageRequest) {
	request = &AddBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBlockPage")
	return
}

func NewAddBlockPageResponse() (response *AddBlockPageResponse) {
	response = &AddBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义拦截页面，每个用户最多允许添加5个自定义拦截页面，添加完后需要审核通过才可以应用到域名。
func (c *Client) AddBlockPage(request *AddBlockPageRequest) (response *AddBlockPageResponse, err error) {
	if request == nil {
		request = NewAddBlockPageRequest()
	}
	response = NewAddBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVIPConfigsRequest() (request *ModifyVIPConfigsRequest) {
	request = &ModifyVIPConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyVIPConfigs")
	return
}

func NewModifyVIPConfigsResponse() (response *ModifyVIPConfigsResponse) {
	response = &ModifyVIPConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改配置
func (c *Client) ModifyVIPConfigs(request *ModifyVIPConfigsRequest) (response *ModifyVIPConfigsResponse, err error) {
	if request == nil {
		request = NewModifyVIPConfigsRequest()
	}
	response = NewModifyVIPConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryClientMsgRequest() (request *QueryClientMsgRequest) {
	request = &QueryClientMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryClientMsg")
	return
}

func NewQueryClientMsgResponse() (response *QueryClientMsgResponse) {
	response = &QueryClientMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否开启透传客户端信息
func (c *Client) QueryClientMsg(request *QueryClientMsgRequest) (response *QueryClientMsgResponse, err error) {
	if request == nil {
		request = NewQueryClientMsgRequest()
	}
	response = NewQueryClientMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRuleOverviewRequest() (request *DescribeBotRuleOverviewRequest) {
	request = &DescribeBotRuleOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRuleOverview")
	return
}

func NewDescribeBotRuleOverviewResponse() (response *DescribeBotRuleOverviewResponse) {
	response = &DescribeBotRuleOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的规则概览
func (c *Client) DescribeBotRuleOverview(request *DescribeBotRuleOverviewRequest) (response *DescribeBotRuleOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotRuleOverviewRequest()
	}
	response = NewDescribeBotRuleOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHybridClustersRequest() (request *DescribeHybridClustersRequest) {
	request = &DescribeHybridClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHybridClusters")
	return
}

func NewDescribeHybridClustersResponse() (response *DescribeHybridClustersResponse) {
	response = &DescribeHybridClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看混合云集群信息
func (c *Client) DescribeHybridClusters(request *DescribeHybridClustersRequest) (response *DescribeHybridClustersResponse, err error) {
	if request == nil {
		request = NewDescribeHybridClustersRequest()
	}
	response = NewDescribeHybridClustersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpSignatureRuleRequest() (request *ModifyOpSignatureRuleRequest) {
	request = &ModifyOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOpSignatureRule")
	return
}

func NewModifyOpSignatureRuleResponse() (response *ModifyOpSignatureRuleResponse) {
	response = &ModifyOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：修改规则
func (c *Client) ModifyOpSignatureRule(request *ModifyOpSignatureRuleRequest) (response *ModifyOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewModifyOpSignatureRuleRequest()
	}
	response = NewModifyOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackTotalCountRequest() (request *GetAttackTotalCountRequest) {
	request = &GetAttackTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackTotalCount")
	return
}

func NewGetAttackTotalCountResponse() (response *GetAttackTotalCountResponse) {
	response = &GetAttackTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照条件查询展示攻击总次数
func (c *Client) GetAttackTotalCount(request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
	if request == nil {
		request = NewGetAttackTotalCountRequest()
	}
	response = NewGetAttackTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDataChartRequest() (request *DescribeDNSDetectDataChartRequest) {
	request = &DescribeDNSDetectDataChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDNSDetectDataChart")
	return
}

func NewDescribeDNSDetectDataChartResponse() (response *DescribeDNSDetectDataChartResponse) {
	response = &DescribeDNSDetectDataChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的被劫持用户数的数据，对应【Web安全检测】【DNS劫持检测】【数据统计】【被劫持用户数】这个功能。
func (c *Client) DescribeDNSDetectDataChart(request *DescribeDNSDetectDataChartRequest) (response *DescribeDNSDetectDataChartResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDataChartRequest()
	}
	response = NewDescribeDNSDetectDataChartResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateInstanceInnerVipRequest() (request *AllocateInstanceInnerVipRequest) {
	request = &AllocateInstanceInnerVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AllocateInstanceInnerVip")
	return
}

func NewAllocateInstanceInnerVipResponse() (response *AllocateInstanceInnerVipResponse) {
	response = &AllocateInstanceInnerVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为SAASWAF实例分配内网VIP，提供给高防EIP调用。
func (c *Client) AllocateInstanceInnerVip(request *AllocateInstanceInnerVipRequest) (response *AllocateInstanceInnerVipResponse, err error) {
	if request == nil {
		request = NewAllocateInstanceInnerVipRequest()
	}
	response = NewAllocateInstanceInnerVipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventOverviewRequest() (request *DescribeApiSecEventOverviewRequest) {
	request = &DescribeApiSecEventOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecEventOverview")
	return
}

func NewDescribeApiSecEventOverviewResponse() (response *DescribeApiSecEventOverviewResponse) {
	response = &DescribeApiSecEventOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// API安全事件概览
func (c *Client) DescribeApiSecEventOverview(request *DescribeApiSecEventOverviewRequest) (response *DescribeApiSecEventOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventOverviewRequest()
	}
	response = NewDescribeApiSecEventOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchCustomRuleStatusRequest() (request *ModifyBatchCustomRuleStatusRequest) {
	request = &ModifyBatchCustomRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchCustomRuleStatus")
	return
}

func NewModifyBatchCustomRuleStatusResponse() (response *ModifyBatchCustomRuleStatusResponse) {
	response = &ModifyBatchCustomRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-批量规则开关接口
func (c *Client) ModifyBatchCustomRuleStatus(request *ModifyBatchCustomRuleStatusRequest) (response *ModifyBatchCustomRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyBatchCustomRuleStatusRequest()
	}
	response = NewModifyBatchCustomRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObjectsRequest() (request *DescribeObjectsRequest) {
	request = &DescribeObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeObjects")
	return
}

func NewDescribeObjectsResponse() (response *DescribeObjectsResponse) {
	response = &DescribeObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看防护对象列表
func (c *Client) DescribeObjects(request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
	if request == nil {
		request = NewDescribeObjectsRequest()
	}
	response = NewDescribeObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewAddAntiInfoLeakRulesRequest() (request *AddAntiInfoLeakRulesRequest) {
	request = &AddAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddAntiInfoLeakRules")
	return
}

func NewAddAntiInfoLeakRulesResponse() (response *AddAntiInfoLeakRulesResponse) {
	response = &AddAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加信息防泄漏规则
func (c *Client) AddAntiInfoLeakRules(request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewAddAntiInfoLeakRulesRequest()
	}
	response = NewAddAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafAutoDenyStatusRequest() (request *DescribeWafAutoDenyStatusRequest) {
	request = &DescribeWafAutoDenyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafAutoDenyStatus")
	return
}

func NewDescribeWafAutoDenyStatusResponse() (response *DescribeWafAutoDenyStatusResponse) {
	response = &DescribeWafAutoDenyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF自动封禁模块详情
func (c *Client) DescribeWafAutoDenyStatus(request *DescribeWafAutoDenyStatusRequest) (response *DescribeWafAutoDenyStatusResponse, err error) {
	if request == nil {
		request = NewDescribeWafAutoDenyStatusRequest()
	}
	response = NewDescribeWafAutoDenyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyModuleStatusRequest() (request *ModifyModuleStatusRequest) {
	request = &ModifyModuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyModuleStatus")
	return
}

func NewModifyModuleStatusResponse() (response *ModifyModuleStatusResponse) {
	response = &ModifyModuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置某个domain下基础安全模块的开关
func (c *Client) ModifyModuleStatus(request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyModuleStatusRequest()
	}
	response = NewModifyModuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackOverviewRequest() (request *DescribeAttackOverviewRequest) {
	request = &DescribeAttackOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackOverview")
	return
}

func NewDescribeAttackOverviewResponse() (response *DescribeAttackOverviewResponse) {
	response = &DescribeAttackOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 攻击总览
func (c *Client) DescribeAttackOverview(request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeAttackOverviewRequest()
	}
	response = NewDescribeAttackOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertCCRuleRequest() (request *UpsertCCRuleRequest) {
	request = &UpsertCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpsertCCRule")
	return
}

func NewUpsertCCRuleResponse() (response *UpsertCCRuleResponse) {
	response = &UpsertCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Upsert接口
func (c *Client) UpsertCCRule(request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
	if request == nil {
		request = NewUpsertCCRuleRequest()
	}
	response = NewUpsertCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddBatchIpPunishmentRequest() (request *AddBatchIpPunishmentRequest) {
	request = &AddBatchIpPunishmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBatchIpPunishment")
	return
}

func NewAddBatchIpPunishmentResponse() (response *AddBatchIpPunishmentResponse) {
	response = &AddBatchIpPunishmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]批量新增接口
func (c *Client) AddBatchIpPunishment(request *AddBatchIpPunishmentRequest) (response *AddBatchIpPunishmentResponse, err error) {
	if request == nil {
		request = NewAddBatchIpPunishmentRequest()
	}
	response = NewAddBatchIpPunishmentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTsResourceRequest() (request *DeleteTsResourceRequest) {
	request = &DeleteTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteTsResource")
	return
}

func NewDeleteTsResourceResponse() (response *DeleteTsResourceResponse) {
	response = &DeleteTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自研引擎服务资源
func (c *Client) DeleteTsResource(request *DeleteTsResourceRequest) (response *DeleteTsResourceResponse, err error) {
	if request == nil {
		request = NewDeleteTsResourceRequest()
	}
	response = NewDeleteTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCheckOverviewRequest() (request *DescribeFlowCheckOverviewRequest) {
	request = &DescribeFlowCheckOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFlowCheckOverview")
	return
}

func NewDescribeFlowCheckOverviewResponse() (response *DescribeFlowCheckOverviewResponse) {
	response = &DescribeFlowCheckOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取流量体检概览数据
func (c *Client) DescribeFlowCheckOverview(request *DescribeFlowCheckOverviewRequest) (response *DescribeFlowCheckOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeFlowCheckOverviewRequest()
	}
	response = NewDescribeFlowCheckOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstanceCountRequest() (request *DescribeUserInstanceCountRequest) {
	request = &DescribeUserInstanceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserInstanceCount")
	return
}

func NewDescribeUserInstanceCountResponse() (response *DescribeUserInstanceCountResponse) {
	response = &DescribeUserInstanceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过客户appid获取客户在所有支持地域的实例数量，返回一个list，key为地域，value为数量
func (c *Client) DescribeUserInstanceCount(request *DescribeUserInstanceCountRequest) (response *DescribeUserInstanceCountResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstanceCountRequest()
	}
	response = NewDescribeUserInstanceCountResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpartaProtectionRequest() (request *AddSpartaProtectionRequest) {
	request = &AddSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddSpartaProtection")
	return
}

func NewAddSpartaProtectionResponse() (response *AddSpartaProtectionResponse) {
	response = &AddSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加SaaS型WAF防护域名
func (c *Client) AddSpartaProtection(request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewAddSpartaProtectionRequest()
	}
	response = NewAddSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateDealsAndPayNewRequest() (request *GenerateDealsAndPayNewRequest) {
	request = &GenerateDealsAndPayNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GenerateDealsAndPayNew")
	return
}

func NewGenerateDealsAndPayNewResponse() (response *GenerateDealsAndPayNewResponse) {
	response = &GenerateDealsAndPayNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费资源购买、续费下单接口
func (c *Client) GenerateDealsAndPayNew(request *GenerateDealsAndPayNewRequest) (response *GenerateDealsAndPayNewResponse, err error) {
	if request == nil {
		request = NewGenerateDealsAndPayNewRequest()
	}
	response = NewGenerateDealsAndPayNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackCountRequest() (request *DescribeAttackCountRequest) {
	request = &DescribeAttackCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackCount")
	return
}

func NewDescribeAttackCountResponse() (response *DescribeAttackCountResponse) {
	response = &DescribeAttackCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定域名，指定时间段的攻击数量
func (c *Client) DescribeAttackCount(request *DescribeAttackCountRequest) (response *DescribeAttackCountResponse, err error) {
	if request == nil {
		request = NewDescribeAttackCountRequest()
	}
	response = NewDescribeAttackCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCRuleListRequest() (request *DescribeCCRuleListRequest) {
	request = &DescribeCCRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCCRuleList")
	return
}

func NewDescribeCCRuleListResponse() (response *DescribeCCRuleListResponse) {
	response = &DescribeCCRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据多条件查询CC规则
func (c *Client) DescribeCCRuleList(request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeCCRuleListRequest()
	}
	response = NewDescribeCCRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTokenRequest() (request *ModifyBotTokenRequest) {
	request = &ModifyBotTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotToken")
	return
}

func NewModifyBotTokenResponse() (response *ModifyBotTokenResponse) {
	response = &ModifyBotTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot会话token配置v2
func (c *Client) ModifyBotToken(request *ModifyBotTokenRequest) (response *ModifyBotTokenResponse, err error) {
	if request == nil {
		request = NewModifyBotTokenRequest()
	}
	response = NewModifyBotTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletTesterTaskRequest() (request *DescribeAppletTesterTaskRequest) {
	request = &DescribeAppletTesterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletTesterTask")
	return
}

func NewDescribeAppletTesterTaskResponse() (response *DescribeAppletTesterTaskResponse) {
	response = &DescribeAppletTesterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序压测任务列表
func (c *Client) DescribeAppletTesterTask(request *DescribeAppletTesterTaskRequest) (response *DescribeAppletTesterTaskResponse, err error) {
	if request == nil {
		request = NewDescribeAppletTesterTaskRequest()
	}
	response = NewDescribeAppletTesterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserEditionRequest() (request *DescribeUserEditionRequest) {
	request = &DescribeUserEditionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserEdition")
	return
}

func NewDescribeUserEditionResponse() (response *DescribeUserEditionResponse) {
	response = &DescribeUserEditionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户购买的WAF的类型，客户是购买的SaaS WAF还是CLB WAF。
func (c *Client) DescribeUserEdition(request *DescribeUserEditionRequest) (response *DescribeUserEditionResponse, err error) {
	if request == nil {
		request = NewDescribeUserEditionRequest()
	}
	response = NewDescribeUserEditionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiAnalyzeStatusRequest() (request *ModifyApiAnalyzeStatusRequest) {
	request = &ModifyApiAnalyzeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiAnalyzeStatus")
	return
}

func NewModifyApiAnalyzeStatusResponse() (response *ModifyApiAnalyzeStatusResponse) {
	response = &ModifyApiAnalyzeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api分析页面开关
func (c *Client) ModifyApiAnalyzeStatus(request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
	if request == nil {
		request = NewModifyApiAnalyzeStatusRequest()
	}
	response = NewModifyApiAnalyzeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiListRequest() (request *DescribeApiListRequest) {
	request = &DescribeApiListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiList")
	return
}

func NewDescribeApiListResponse() (response *DescribeApiListResponse) {
	response = &DescribeApiListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示api规则列表信息
func (c *Client) DescribeApiList(request *DescribeApiListRequest) (response *DescribeApiListResponse, err error) {
	if request == nil {
		request = NewDescribeApiListRequest()
	}
	response = NewDescribeApiListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRenewWhiteListRequest() (request *DescribeRenewWhiteListRequest) {
	request = &DescribeRenewWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRenewWhiteList")
	return
}

func NewDescribeRenewWhiteListResponse() (response *DescribeRenewWhiteListResponse) {
	response = &DescribeRenewWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的CLS信息
func (c *Client) DescribeRenewWhiteList(request *DescribeRenewWhiteListRequest) (response *DescribeRenewWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeRenewWhiteListRequest()
	}
	response = NewDescribeRenewWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMLimitRuleRequest() (request *ModifyLLMLimitRuleRequest) {
	request = &ModifyLLMLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMLimitRule")
	return
}

func NewModifyLLMLimitRuleResponse() (response *ModifyLLMLimitRuleResponse) {
	response = &ModifyLLMLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】算力消耗防护规则，两个调用位置：1.添加算力消耗防护规则 2.编辑算力消耗防护规则
func (c *Client) ModifyLLMLimitRule(request *ModifyLLMLimitRuleRequest) (response *ModifyLLMLimitRuleResponse, err error) {
	if request == nil {
		request = NewModifyLLMLimitRuleRequest()
	}
	response = NewModifyLLMLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAttackDownloadTaskRequest() (request *CreateAttackDownloadTaskRequest) {
	request = &CreateAttackDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAttackDownloadTask")
	return
}

func NewCreateAttackDownloadTaskResponse() (response *CreateAttackDownloadTaskResponse) {
	response = &CreateAttackDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建攻击日志下载任务
func (c *Client) CreateAttackDownloadTask(request *CreateAttackDownloadTaskRequest) (response *CreateAttackDownloadTaskResponse, err error) {
	if request == nil {
		request = NewCreateAttackDownloadTaskRequest()
	}
	response = NewCreateAttackDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDetailRequest() (request *DescribeAttackDetailRequest) {
	request = &DescribeAttackDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackDetail")
	return
}

func NewDescribeAttackDetailResponse() (response *DescribeAttackDetailResponse) {
	response = &DescribeAttackDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取攻击日志详情
func (c *Client) DescribeAttackDetail(request *DescribeAttackDetailRequest) (response *DescribeAttackDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDetailRequest()
	}
	response = NewDescribeAttackDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailRequest() (request *DescribeBotDetailRequest) {
	request = &DescribeBotDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotDetail")
	return
}

func NewDescribeBotDetailResponse() (response *DescribeBotDetailResponse) {
	response = &DescribeBotDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot的具体详情
func (c *Client) DescribeBotDetail(request *DescribeBotDetailRequest) (response *DescribeBotDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailRequest()
	}
	response = NewDescribeBotDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanIpRequest() (request *DescribeScanIpRequest) {
	request = &DescribeScanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeScanIp")
	return
}

func NewDescribeScanIpResponse() (response *DescribeScanIpResponse) {
	response = &DescribeScanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询扫描ip
func (c *Client) DescribeScanIp(request *DescribeScanIpRequest) (response *DescribeScanIpResponse, err error) {
	if request == nil {
		request = NewDescribeScanIpRequest()
	}
	response = NewDescribeScanIpResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostInnerRequest() (request *DeleteHostInnerRequest) {
	request = &DeleteHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteHostInner")
	return
}

func NewDeleteHostInnerResponse() (response *DeleteHostInnerResponse) {
	response = &DeleteHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CLB-WAF防护域名
// 支持批量操作
func (c *Client) DeleteHostInner(request *DeleteHostInnerRequest) (response *DeleteHostInnerResponse, err error) {
	if request == nil {
		request = NewDeleteHostInnerRequest()
	}
	response = NewDeleteHostInnerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTsResourceRequest() (request *ModifyTsResourceRequest) {
	request = &ModifyTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyTsResource")
	return
}

func NewModifyTsResourceResponse() (response *ModifyTsResourceResponse) {
	response = &ModifyTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Tiga资源列表，如果id=0则添加，>0则修改
func (c *Client) ModifyTsResource(request *ModifyTsResourceRequest) (response *ModifyTsResourceResponse, err error) {
	if request == nil {
		request = NewModifyTsResourceRequest()
	}
	response = NewModifyTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletAuthorizeV3Request() (request *DescribeAppletAuthorizeV3Request) {
	request = &DescribeAppletAuthorizeV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletAuthorizeV3")
	return
}

func NewDescribeAppletAuthorizeV3Response() (response *DescribeAppletAuthorizeV3Response) {
	response = &DescribeAppletAuthorizeV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序授权列表3.0
func (c *Client) DescribeAppletAuthorizeV3(request *DescribeAppletAuthorizeV3Request) (response *DescribeAppletAuthorizeV3Response, err error) {
	if request == nil {
		request = NewDescribeAppletAuthorizeV3Request()
	}
	response = NewDescribeAppletAuthorizeV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTokenRequest() (request *DescribeBotTokenRequest) {
	request = &DescribeBotTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotToken")
	return
}

func NewDescribeBotTokenResponse() (response *DescribeBotTokenResponse) {
	response = &DescribeBotTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot用户自定义会话配置列表
func (c *Client) DescribeBotToken(request *DescribeBotTokenRequest) (response *DescribeBotTokenResponse, err error) {
	if request == nil {
		request = NewDescribeBotTokenRequest()
	}
	response = NewDescribeBotTokenResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFlowCheckStatusRequest() (request *ModifyFlowCheckStatusRequest) {
	request = &ModifyFlowCheckStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyFlowCheckStatus")
	return
}

func NewModifyFlowCheckStatusResponse() (response *ModifyFlowCheckStatusResponse) {
	response = &ModifyFlowCheckStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改体检状态接口
func (c *Client) ModifyFlowCheckStatus(request *ModifyFlowCheckStatusRequest) (response *ModifyFlowCheckStatusResponse, err error) {
	if request == nil {
		request = NewModifyFlowCheckStatusRequest()
	}
	response = NewModifyFlowCheckStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageItemRequest() (request *SetLocalStorageItemRequest) {
	request = &SetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SetLocalStorageItem")
	return
}

func NewSetLocalStorageItemResponse() (response *SetLocalStorageItemResponse) {
	response = &SetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加远程缓存
func (c *Client) SetLocalStorageItem(request *SetLocalStorageItemRequest) (response *SetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageItemRequest()
	}
	response = NewSetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserQPSRequest() (request *DescribeUserQPSRequest) {
	request = &DescribeUserQPSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserQPS")
	return
}

func NewDescribeUserQPSResponse() (response *DescribeUserQPSResponse) {
	response = &DescribeUserQPSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取客户的QPS峰值
func (c *Client) DescribeUserQPS(request *DescribeUserQPSRequest) (response *DescribeUserQPSResponse, err error) {
	if request == nil {
		request = NewDescribeUserQPSRequest()
	}
	response = NewDescribeUserQPSResponse()
	err = c.Send(request, response)
	return
}

func NewAddAreaBanAreasRequest() (request *AddAreaBanAreasRequest) {
	request = &AddAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddAreaBanAreas")
	return
}

func NewAddAreaBanAreasResponse() (response *AddAreaBanAreasResponse) {
	response = &AddAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加地域封禁中的地域信息
func (c *Client) AddAreaBanAreas(request *AddAreaBanAreasRequest) (response *AddAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewAddAreaBanAreasRequest()
	}
	response = NewAddAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainDetailsSaasRequest() (request *DescribeDomainDetailsSaasRequest) {
	request = &DescribeDomainDetailsSaasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainDetailsSaas")
	return
}

func NewDescribeDomainDetailsSaasResponse() (response *DescribeDomainDetailsSaasResponse) {
	response = &DescribeDomainDetailsSaasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个saaswaf域名详情
func (c *Client) DescribeDomainDetailsSaas(request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
	if request == nil {
		request = NewDescribeDomainDetailsSaasRequest()
	}
	response = NewDescribeDomainDetailsSaasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchAreaBanRulesRequest() (request *DescribeBatchAreaBanRulesRequest) {
	request = &DescribeBatchAreaBanRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchAreaBanRules")
	return
}

func NewDescribeBatchAreaBanRulesResponse() (response *DescribeBatchAreaBanRulesResponse) {
	response = &DescribeBatchAreaBanRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取批量地域封禁规则
func (c *Client) DescribeBatchAreaBanRules(request *DescribeBatchAreaBanRulesRequest) (response *DescribeBatchAreaBanRulesResponse, err error) {
	if request == nil {
		request = NewDescribeBatchAreaBanRulesRequest()
	}
	response = NewDescribeBatchAreaBanRulesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackHistogramRequest() (request *GetAttackHistogramRequest) {
	request = &GetAttackHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackHistogram")
	return
}

func NewGetAttackHistogramResponse() (response *GetAttackHistogramResponse) {
	response = &GetAttackHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成攻击日志的产生时间柱状图
func (c *Client) GetAttackHistogram(request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
	if request == nil {
		request = NewGetAttackHistogramRequest()
	}
	response = NewGetAttackHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBatchIpPunishmentAppIdDomainRequest() (request *QueryBatchIpPunishmentAppIdDomainRequest) {
	request = &QueryBatchIpPunishmentAppIdDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryBatchIpPunishmentAppIdDomain")
	return
}

func NewQueryBatchIpPunishmentAppIdDomainResponse() (response *QueryBatchIpPunishmentAppIdDomainResponse) {
	response = &QueryBatchIpPunishmentAppIdDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]查询当前AppId下域名接口
func (c *Client) QueryBatchIpPunishmentAppIdDomain(request *QueryBatchIpPunishmentAppIdDomainRequest) (response *QueryBatchIpPunishmentAppIdDomainResponse, err error) {
	if request == nil {
		request = NewQueryBatchIpPunishmentAppIdDomainRequest()
	}
	response = NewQueryBatchIpPunishmentAppIdDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolicyStatusRequest() (request *DescribePolicyStatusRequest) {
	request = &DescribePolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePolicyStatus")
	return
}

func NewDescribePolicyStatusResponse() (response *DescribePolicyStatusResponse) {
	response = &DescribePolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护状态以及生效的实例id
func (c *Client) DescribePolicyStatus(request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
	if request == nil {
		request = NewDescribePolicyStatusRequest()
	}
	response = NewDescribePolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResponsePagesRequest() (request *DescribeResponsePagesRequest) {
	request = &DescribeResponsePagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeResponsePages")
	return
}

func NewDescribeResponsePagesResponse() (response *DescribeResponsePagesResponse) {
	response = &DescribeResponsePagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示自定义响应页面列表信息。
func (c *Client) DescribeResponsePages(request *DescribeResponsePagesRequest) (response *DescribeResponsePagesResponse, err error) {
	if request == nil {
		request = NewDescribeResponsePagesRequest()
	}
	response = NewDescribeResponsePagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLimitInstanceQpsConfRequest() (request *DescribeLimitInstanceQpsConfRequest) {
	request = &DescribeLimitInstanceQpsConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLimitInstanceQpsConf")
	return
}

func NewDescribeLimitInstanceQpsConfResponse() (response *DescribeLimitInstanceQpsConfResponse) {
	response = &DescribeLimitInstanceQpsConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询需要限流的实例配置信息
func (c *Client) DescribeLimitInstanceQpsConf(request *DescribeLimitInstanceQpsConfRequest) (response *DescribeLimitInstanceQpsConfResponse, err error) {
	if request == nil {
		request = NewDescribeLimitInstanceQpsConfRequest()
	}
	response = NewDescribeLimitInstanceQpsConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpRuleUpdateLogRequest() (request *DescribeOpRuleUpdateLogRequest) {
	request = &DescribeOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpRuleUpdateLog")
	return
}

func NewDescribeOpRuleUpdateLogResponse() (response *DescribeOpRuleUpdateLogResponse) {
	response = &DescribeOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：获取特征规则更新动态
func (c *Client) DescribeOpRuleUpdateLog(request *DescribeOpRuleUpdateLogRequest) (response *DescribeOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDescribeOpRuleUpdateLogRequest()
	}
	response = NewDescribeOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchIpPunishmentRequest() (request *ModifyBatchIpPunishmentRequest) {
	request = &ModifyBatchIpPunishmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchIpPunishment")
	return
}

func NewModifyBatchIpPunishmentResponse() (response *ModifyBatchIpPunishmentResponse) {
	response = &ModifyBatchIpPunishmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]批量编辑接口
func (c *Client) ModifyBatchIpPunishment(request *ModifyBatchIpPunishmentRequest) (response *ModifyBatchIpPunishmentResponse, err error) {
	if request == nil {
		request = NewModifyBatchIpPunishmentRequest()
	}
	response = NewModifyBatchIpPunishmentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotAutoStatisticRuleRequest() (request *ModifyBotAutoStatisticRuleRequest) {
	request = &ModifyBotAutoStatisticRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotAutoStatisticRule")
	return
}

func NewModifyBotAutoStatisticRuleResponse() (response *ModifyBotAutoStatisticRuleResponse) {
	response = &ModifyBotAutoStatisticRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-智能统计分析模块配置
func (c *Client) ModifyBotAutoStatisticRule(request *ModifyBotAutoStatisticRuleRequest) (response *ModifyBotAutoStatisticRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotAutoStatisticRuleRequest()
	}
	response = NewModifyBotAutoStatisticRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceVipStatusRequest() (request *ModifyInstanceVipStatusRequest) {
	request = &ModifyInstanceVipStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceVipStatus")
	return
}

func NewModifyInstanceVipStatusResponse() (response *ModifyInstanceVipStatusResponse) {
	response = &ModifyInstanceVipStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作WAF实例的原生VIP状态
func (c *Client) ModifyInstanceVipStatus(request *ModifyInstanceVipStatusRequest) (response *ModifyInstanceVipStatusResponse, err error) {
	if request == nil {
		request = NewModifyInstanceVipStatusRequest()
	}
	response = NewModifyInstanceVipStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotCountRequest() (request *DescribeBotCountRequest) {
	request = &DescribeBotCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotCount")
	return
}

func NewDescribeBotCountResponse() (response *DescribeBotCountResponse) {
	response = &DescribeBotCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定域名
func (c *Client) DescribeBotCount(request *DescribeBotCountRequest) (response *DescribeBotCountResponse, err error) {
	if request == nil {
		request = NewDescribeBotCountRequest()
	}
	response = NewDescribeBotCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTCBRuleRequest() (request *ModifyBotTCBRuleRequest) {
	request = &ModifyBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotTCBRule")
	return
}

func NewModifyBotTCBRuleResponse() (response *ModifyBotTCBRuleResponse) {
	response = &ModifyBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 TCB策略更新
func (c *Client) ModifyBotTCBRule(request *ModifyBotTCBRuleRequest) (response *ModifyBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotTCBRuleRequest()
	}
	response = NewModifyBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpHitItemsRequest() (request *DescribeIpHitItemsRequest) {
	request = &DescribeIpHitItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeIpHitItems")
	return
}

func NewDescribeIpHitItemsResponse() (response *DescribeIpHitItemsResponse) {
	response = &DescribeIpHitItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  IP封堵状态查询
func (c *Client) DescribeIpHitItems(request *DescribeIpHitItemsRequest) (response *DescribeIpHitItemsResponse, err error) {
	if request == nil {
		request = NewDescribeIpHitItemsRequest()
	}
	response = NewDescribeIpHitItemsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchElasticModeRequest() (request *SwitchElasticModeRequest) {
	request = &SwitchElasticModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchElasticMode")
	return
}

func NewSwitchElasticModeResponse() (response *SwitchElasticModeResponse) {
	response = &SwitchElasticModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换弹性的开关
func (c *Client) SwitchElasticMode(request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
	if request == nil {
		request = NewSwitchElasticModeRequest()
	}
	response = NewSwitchElasticModeResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePostCLSFlowRequest() (request *CreatePostCLSFlowRequest) {
	request = &CreatePostCLSFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreatePostCLSFlow")
	return
}

func NewCreatePostCLSFlowResponse() (response *CreatePostCLSFlowResponse) {
	response = &CreatePostCLSFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CLS投递流任务
func (c *Client) CreatePostCLSFlow(request *CreatePostCLSFlowRequest) (response *CreatePostCLSFlowResponse, err error) {
	if request == nil {
		request = NewCreatePostCLSFlowRequest()
	}
	response = NewCreatePostCLSFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowTrendRequest() (request *DescribeFlowTrendRequest) {
	request = &DescribeFlowTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFlowTrend")
	return
}

func NewDescribeFlowTrendResponse() (response *DescribeFlowTrendResponse) {
	response = &DescribeFlowTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取waf流量访问趋势
func (c *Client) DescribeFlowTrend(request *DescribeFlowTrendRequest) (response *DescribeFlowTrendResponse, err error) {
	if request == nil {
		request = NewDescribeFlowTrendRequest()
	}
	response = NewDescribeFlowTrendResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectGroupRequest() (request *ModifyProtectGroupRequest) {
	request = &ModifyProtectGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyProtectGroup")
	return
}

func NewModifyProtectGroupResponse() (response *ModifyProtectGroupResponse) {
	response = &ModifyProtectGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑防护对象组
func (c *Client) ModifyProtectGroup(request *ModifyProtectGroupRequest) (response *ModifyProtectGroupResponse, err error) {
	if request == nil {
		request = NewModifyProtectGroupRequest()
	}
	response = NewModifyProtectGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCLSRequest() (request *DescribeCLSRequest) {
	request = &DescribeCLSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCLS")
	return
}

func NewDescribeCLSResponse() (response *DescribeCLSResponse) {
	response = &DescribeCLSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前购买的日志服务的使用状况，用户必须已经购买了日志服务才能调用此接口
func (c *Client) DescribeCLS(request *DescribeCLSRequest) (response *DescribeCLSResponse, err error) {
	if request == nil {
		request = NewDescribeCLSRequest()
	}
	response = NewDescribeCLSResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePocTestAuthorizationStatusRequest() (request *DescribePocTestAuthorizationStatusRequest) {
	request = &DescribePocTestAuthorizationStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePocTestAuthorizationStatus")
	return
}

func NewDescribePocTestAuthorizationStatusResponse() (response *DescribePocTestAuthorizationStatusResponse) {
	response = &DescribePocTestAuthorizationStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询APPID授权状态
func (c *Client) DescribePocTestAuthorizationStatus(request *DescribePocTestAuthorizationStatusRequest) (response *DescribePocTestAuthorizationStatusResponse, err error) {
	if request == nil {
		request = NewDescribePocTestAuthorizationStatusRequest()
	}
	response = NewDescribePocTestAuthorizationStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppletAccessRequest() (request *DeleteAppletAccessRequest) {
	request = &DeleteAppletAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAppletAccess")
	return
}

func NewDeleteAppletAccessResponse() (response *DeleteAppletAccessResponse) {
	response = &DeleteAppletAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除小程序接入
func (c *Client) DeleteAppletAccess(request *DeleteAppletAccessRequest) (response *DeleteAppletAccessResponse, err error) {
	if request == nil {
		request = NewDeleteAppletAccessRequest()
	}
	response = NewDeleteAppletAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotEventListRequest() (request *DescribeBotEventListRequest) {
	request = &DescribeBotEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotEventList")
	return
}

func NewDescribeBotEventListResponse() (response *DescribeBotEventListResponse) {
	response = &DescribeBotEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-拉取BOT事件列表
func (c *Client) DescribeBotEventList(request *DescribeBotEventListRequest) (response *DescribeBotEventListResponse, err error) {
	if request == nil {
		request = NewDescribeBotEventListRequest()
	}
	response = NewDescribeBotEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExperienceVersionInfoRequest() (request *DescribeExperienceVersionInfoRequest) {
	request = &DescribeExperienceVersionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeExperienceVersionInfo")
	return
}

func NewDescribeExperienceVersionInfoResponse() (response *DescribeExperienceVersionInfoResponse) {
	response = &DescribeExperienceVersionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取体验版信息
func (c *Client) DescribeExperienceVersionInfo(request *DescribeExperienceVersionInfoRequest) (response *DescribeExperienceVersionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeExperienceVersionInfoRequest()
	}
	response = NewDescribeExperienceVersionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserDefaultEngineRequest() (request *ModifyUserDefaultEngineRequest) {
	request = &ModifyUserDefaultEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserDefaultEngine")
	return
}

func NewModifyUserDefaultEngineResponse() (response *ModifyUserDefaultEngineResponse) {
	response = &ModifyUserDefaultEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户默认规则引擎类型
func (c *Client) ModifyUserDefaultEngine(request *ModifyUserDefaultEngineRequest) (response *ModifyUserDefaultEngineResponse, err error) {
	if request == nil {
		request = NewModifyUserDefaultEngineRequest()
	}
	response = NewModifyUserDefaultEngineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBatchIpAccessControlRequest() (request *CreateBatchIpAccessControlRequest) {
	request = &CreateBatchIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateBatchIpAccessControl")
	return
}

func NewCreateBatchIpAccessControlResponse() (response *CreateBatchIpAccessControlResponse) {
	response = &CreateBatchIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量IP黑白名单新增接口
func (c *Client) CreateBatchIpAccessControl(request *CreateBatchIpAccessControlRequest) (response *CreateBatchIpAccessControlResponse, err error) {
	if request == nil {
		request = NewCreateBatchIpAccessControlRequest()
	}
	response = NewCreateBatchIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWxUserGroupRequest() (request *DescribeWxUserGroupRequest) {
	request = &DescribeWxUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWxUserGroup")
	return
}

func NewDescribeWxUserGroupResponse() (response *DescribeWxUserGroupResponse) {
	response = &DescribeWxUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户分组
func (c *Client) DescribeWxUserGroup(request *DescribeWxUserGroupRequest) (response *DescribeWxUserGroupResponse, err error) {
	if request == nil {
		request = NewDescribeWxUserGroupRequest()
	}
	response = NewDescribeWxUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAutoStatisticRuleRequest() (request *DescribeBotAutoStatisticRuleRequest) {
	request = &DescribeBotAutoStatisticRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAutoStatisticRule")
	return
}

func NewDescribeBotAutoStatisticRuleResponse() (response *DescribeBotAutoStatisticRuleResponse) {
	response = &DescribeBotAutoStatisticRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot-智能统计模块配置
func (c *Client) DescribeBotAutoStatisticRule(request *DescribeBotAutoStatisticRuleRequest) (response *DescribeBotAutoStatisticRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotAutoStatisticRuleRequest()
	}
	response = NewDescribeBotAutoStatisticRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRateLimitV2Request() (request *CreateRateLimitV2Request) {
	request = &CreateRateLimitV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateRateLimitV2")
	return
}

func NewCreateRateLimitV2Response() (response *CreateRateLimitV2Response) {
	response = &CreateRateLimitV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建限流规则
func (c *Client) CreateRateLimitV2(request *CreateRateLimitV2Request) (response *CreateRateLimitV2Response, err error) {
	if request == nil {
		request = NewCreateRateLimitV2Request()
	}
	response = NewCreateRateLimitV2Response()
	err = c.Send(request, response)
	return
}

func NewModifyWhiteListRequest() (request *ModifyWhiteListRequest) {
	request = &ModifyWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWhiteList")
	return
}

func NewModifyWhiteListResponse() (response *ModifyWhiteListResponse) {
	response = &ModifyWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新业务白名单，目前是前端对抗中的白名单
func (c *Client) ModifyWhiteList(request *ModifyWhiteListRequest) (response *ModifyWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyWhiteListRequest()
	}
	response = NewModifyWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotWhiteStatusRequest() (request *DescribeBotWhiteStatusRequest) {
	request = &DescribeBotWhiteStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotWhiteStatus")
	return
}

func NewDescribeBotWhiteStatusResponse() (response *DescribeBotWhiteStatusResponse) {
	response = &DescribeBotWhiteStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取goodbot的开关列表
func (c *Client) DescribeBotWhiteStatus(request *DescribeBotWhiteStatusRequest) (response *DescribeBotWhiteStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBotWhiteStatusRequest()
	}
	response = NewDescribeBotWhiteStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOwaspRuleTypeLevelRequest() (request *ModifyOwaspRuleTypeLevelRequest) {
	request = &ModifyOwaspRuleTypeLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOwaspRuleTypeLevel")
	return
}

func NewModifyOwaspRuleTypeLevelResponse() (response *ModifyOwaspRuleTypeLevelResponse) {
	response = &ModifyOwaspRuleTypeLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新规则类型的防护等级
func (c *Client) ModifyOwaspRuleTypeLevel(request *ModifyOwaspRuleTypeLevelRequest) (response *ModifyOwaspRuleTypeLevelResponse, err error) {
	if request == nil {
		request = NewModifyOwaspRuleTypeLevelRequest()
	}
	response = NewModifyOwaspRuleTypeLevelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaBanRuleRequest() (request *ModifyAreaBanRuleRequest) {
	request = &ModifyAreaBanRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAreaBanRule")
	return
}

func NewModifyAreaBanRuleResponse() (response *ModifyAreaBanRuleResponse) {
	response = &ModifyAreaBanRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加（编辑）地域封禁中的地域信息
func (c *Client) ModifyAreaBanRule(request *ModifyAreaBanRuleRequest) (response *ModifyAreaBanRuleResponse, err error) {
	if request == nil {
		request = NewModifyAreaBanRuleRequest()
	}
	response = NewModifyAreaBanRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTceAlarmRequest() (request *DescribeTceAlarmRequest) {
	request = &DescribeTceAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTceAlarm")
	return
}

func NewDescribeTceAlarmResponse() (response *DescribeTceAlarmResponse) {
	response = &DescribeTceAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警配置
func (c *Client) DescribeTceAlarm(request *DescribeTceAlarmRequest) (response *DescribeTceAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeTceAlarmRequest()
	}
	response = NewDescribeTceAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchAreaBanListRequest() (request *DescribeBatchAreaBanListRequest) {
	request = &DescribeBatchAreaBanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchAreaBanList")
	return
}

func NewDescribeBatchAreaBanListResponse() (response *DescribeBatchAreaBanListResponse) {
	response = &DescribeBatchAreaBanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-批量列表接口
func (c *Client) DescribeBatchAreaBanList(request *DescribeBatchAreaBanListRequest) (response *DescribeBatchAreaBanListResponse, err error) {
	if request == nil {
		request = NewDescribeBatchAreaBanListRequest()
	}
	response = NewDescribeBatchAreaBanListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchIpPunishmentRequest() (request *DeleteBatchIpPunishmentRequest) {
	request = &DeleteBatchIpPunishmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBatchIpPunishment")
	return
}

func NewDeleteBatchIpPunishmentResponse() (response *DeleteBatchIpPunishmentResponse) {
	response = &DeleteBatchIpPunishmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]批量删除接口
func (c *Client) DeleteBatchIpPunishment(request *DeleteBatchIpPunishmentRequest) (response *DeleteBatchIpPunishmentResponse, err error) {
	if request == nil {
		request = NewDeleteBatchIpPunishmentRequest()
	}
	response = NewDeleteBatchIpPunishmentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCertificateVerifyResultRequest() (request *DescribeCertificateVerifyResultRequest) {
	request = &DescribeCertificateVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCertificateVerifyResult")
	return
}

func NewDescribeCertificateVerifyResultResponse() (response *DescribeCertificateVerifyResultResponse) {
	response = &DescribeCertificateVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取证书的检查结果
func (c *Client) DescribeCertificateVerifyResult(request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeCertificateVerifyResultRequest()
	}
	response = NewDescribeCertificateVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessDownloadRecordRequest() (request *CreateAccessDownloadRecordRequest) {
	request = &CreateAccessDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAccessDownloadRecord")
	return
}

func NewCreateAccessDownloadRecordResponse() (response *CreateAccessDownloadRecordResponse) {
	response = &CreateAccessDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载访问日志，必须购买日志服务
func (c *Client) CreateAccessDownloadRecord(request *CreateAccessDownloadRecordRequest) (response *CreateAccessDownloadRecordResponse, err error) {
	if request == nil {
		request = NewCreateAccessDownloadRecordRequest()
	}
	response = NewCreateAccessDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBlockPageRequest() (request *DeleteBlockPageRequest) {
	request = &DeleteBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBlockPage")
	return
}

func NewDeleteBlockPageResponse() (response *DeleteBlockPageResponse) {
	response = &DeleteBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义拦截页面，所有应用了该拦截页面的域名都会自动删除该拦截页面。
func (c *Client) DeleteBlockPage(request *DeleteBlockPageRequest) (response *DeleteBlockPageResponse, err error) {
	if request == nil {
		request = NewDeleteBlockPageRequest()
	}
	response = NewDeleteBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletAccessRequest() (request *DescribeAppletAccessRequest) {
	request = &DescribeAppletAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletAccess")
	return
}

func NewDescribeAppletAccessResponse() (response *DescribeAppletAccessResponse) {
	response = &DescribeAppletAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序接入列表
func (c *Client) DescribeAppletAccess(request *DescribeAppletAccessRequest) (response *DescribeAppletAccessResponse, err error) {
	if request == nil {
		request = NewDescribeAppletAccessRequest()
	}
	response = NewDescribeAppletAccessResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMProtectPathRequest() (request *ModifyLLMProtectPathRequest) {
	request = &ModifyLLMProtectPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMProtectPath")
	return
}

func NewModifyLLMProtectPathResponse() (response *ModifyLLMProtectPathResponse) {
	response = &ModifyLLMProtectPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护路径配置
func (c *Client) ModifyLLMProtectPath(request *ModifyLLMProtectPathRequest) (response *ModifyLLMProtectPathResponse, err error) {
	if request == nil {
		request = NewModifyLLMProtectPathRequest()
	}
	response = NewModifyLLMProtectPathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTypeStatRequest() (request *DescribeBotTypeStatRequest) {
	request = &DescribeBotTypeStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTypeStat")
	return
}

func NewDescribeBotTypeStatResponse() (response *DescribeBotTypeStatResponse) {
	response = &DescribeBotTypeStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 Bot类别统计
func (c *Client) DescribeBotTypeStat(request *DescribeBotTypeStatRequest) (response *DescribeBotTypeStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotTypeStatRequest()
	}
	response = NewDescribeBotTypeStatResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotUCBFeatureRuleRequest() (request *DeleteBotUCBFeatureRuleRequest) {
	request = &DeleteBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBotUCBFeatureRule")
	return
}

func NewDeleteBotUCBFeatureRuleResponse() (response *DeleteBotUCBFeatureRuleResponse) {
	response = &DeleteBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB自定义策略删除
func (c *Client) DeleteBotUCBFeatureRule(request *DeleteBotUCBFeatureRuleRequest) (response *DeleteBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotUCBFeatureRuleRequest()
	}
	response = NewDeleteBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTsRegionRequest() (request *DeleteTsRegionRequest) {
	request = &DeleteTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteTsRegion")
	return
}

func NewDeleteTsRegionResponse() (response *DeleteTsRegionResponse) {
	response = &DeleteTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ts接入区域
func (c *Client) DeleteTsRegion(request *DeleteTsRegionRequest) (response *DeleteTsRegionResponse, err error) {
	if request == nil {
		request = NewDeleteTsRegionRequest()
	}
	response = NewDeleteTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainEngineRequest() (request *DescribeDomainEngineRequest) {
	request = &DescribeDomainEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainEngine")
	return
}

func NewDescribeDomainEngineResponse() (response *DescribeDomainEngineResponse) {
	response = &DescribeDomainEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询域名引擎类型
func (c *Client) DescribeDomainEngine(request *DescribeDomainEngineRequest) (response *DescribeDomainEngineResponse, err error) {
	if request == nil {
		request = NewDescribeDomainEngineRequest()
	}
	response = NewDescribeDomainEngineResponse()
	err = c.Send(request, response)
	return
}

func NewDisableUserSignaturePolicyRequest() (request *DisableUserSignaturePolicyRequest) {
	request = &DisableUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DisableUserSignaturePolicy")
	return
}

func NewDisableUserSignaturePolicyResponse() (response *DisableUserSignaturePolicyResponse) {
	response = &DisableUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭用户防护规则开关
func (c *Client) DisableUserSignaturePolicy(request *DisableUserSignaturePolicyRequest) (response *DisableUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewDisableUserSignaturePolicyRequest()
	}
	response = NewDisableUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiInfoLeakRulesRequest() (request *ModifyAntiInfoLeakRulesRequest) {
	request = &ModifyAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAntiInfoLeakRules")
	return
}

func NewModifyAntiInfoLeakRulesResponse() (response *ModifyAntiInfoLeakRulesResponse) {
	response = &ModifyAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑信息防泄漏规则
func (c *Client) ModifyAntiInfoLeakRules(request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewModifyAntiInfoLeakRulesRequest()
	}
	response = NewModifyAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiRequestParameterRequest() (request *ModifyApiRequestParameterRequest) {
	request = &ModifyApiRequestParameterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiRequestParameter")
	return
}

func NewModifyApiRequestParameterResponse() (response *ModifyApiRequestParameterResponse) {
	response = &ModifyApiRequestParameterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改api自动发现的参数
func (c *Client) ModifyApiRequestParameter(request *ModifyApiRequestParameterRequest) (response *ModifyApiRequestParameterResponse, err error) {
	if request == nil {
		request = NewModifyApiRequestParameterRequest()
	}
	response = NewModifyApiRequestParameterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneStatusRequest() (request *ModifyBotSceneStatusRequest) {
	request = &ModifyBotSceneStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotSceneStatus")
	return
}

func NewModifyBotSceneStatusResponse() (response *ModifyBotSceneStatusResponse) {
	response = &ModifyBotSceneStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot子场景开关
func (c *Client) ModifyBotSceneStatus(request *ModifyBotSceneStatusRequest) (response *ModifyBotSceneStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneStatusRequest()
	}
	response = NewModifyBotSceneStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBatchAreaBanAppIdDomainRequest() (request *QueryBatchAreaBanAppIdDomainRequest) {
	request = &QueryBatchAreaBanAppIdDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryBatchAreaBanAppIdDomain")
	return
}

func NewQueryBatchAreaBanAppIdDomainResponse() (response *QueryBatchAreaBanAppIdDomainResponse) {
	response = &QueryBatchAreaBanAppIdDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-查询当前AppId下批量域名接口
func (c *Client) QueryBatchAreaBanAppIdDomain(request *QueryBatchAreaBanAppIdDomainRequest) (response *QueryBatchAreaBanAppIdDomainResponse, err error) {
	if request == nil {
		request = NewQueryBatchAreaBanAppIdDomainRequest()
	}
	response = NewQueryBatchAreaBanAppIdDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRequestCountRequest() (request *DescribeRequestCountRequest) {
	request = &DescribeRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRequestCount")
	return
}

func NewDescribeRequestCountResponse() (response *DescribeRequestCountResponse) {
	response = &DescribeRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定时间段内请求总数，这个接口不需要客户购买日志服务
func (c *Client) DescribeRequestCount(request *DescribeRequestCountRequest) (response *DescribeRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeRequestCountRequest()
	}
	response = NewDescribeRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletQpsLimitRequest() (request *ModifyAppletQpsLimitRequest) {
	request = &ModifyAppletQpsLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletQpsLimit")
	return
}

func NewModifyAppletQpsLimitResponse() (response *ModifyAppletQpsLimitResponse) {
	response = &ModifyAppletQpsLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置小程序的弹性qps上限
func (c *Client) ModifyAppletQpsLimit(request *ModifyAppletQpsLimitRequest) (response *ModifyAppletQpsLimitResponse, err error) {
	if request == nil {
		request = NewModifyAppletQpsLimitRequest()
	}
	response = NewModifyAppletQpsLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomPayloadsRequest() (request *DescribeCustomPayloadsRequest) {
	request = &DescribeCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomPayloads")
	return
}

func NewDescribeCustomPayloadsResponse() (response *DescribeCustomPayloadsResponse) {
	response = &DescribeCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义payload列表
func (c *Client) DescribeCustomPayloads(request *DescribeCustomPayloadsRequest) (response *DescribeCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomPayloadsRequest()
	}
	response = NewDescribeCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchAreaBanStatusRequest() (request *ModifyBatchAreaBanStatusRequest) {
	request = &ModifyBatchAreaBanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchAreaBanStatus")
	return
}

func NewModifyBatchAreaBanStatusResponse() (response *ModifyBatchAreaBanStatusResponse) {
	response = &ModifyBatchAreaBanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-批量规则开关接口
func (c *Client) ModifyBatchAreaBanStatus(request *ModifyBatchAreaBanStatusRequest) (response *ModifyBatchAreaBanStatusResponse, err error) {
	if request == nil {
		request = NewModifyBatchAreaBanStatusRequest()
	}
	response = NewModifyBatchAreaBanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAggregateTopNRequest() (request *DescribeBotAggregateTopNRequest) {
	request = &DescribeBotAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAggregateTopN")
	return
}

func NewDescribeBotAggregateTopNResponse() (response *DescribeBotAggregateTopNResponse) {
	response = &DescribeBotAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT的日志聚合topN
func (c *Client) DescribeBotAggregateTopN(request *DescribeBotAggregateTopNRequest) (response *DescribeBotAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotAggregateTopNRequest()
	}
	response = NewDescribeBotAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceIdListRequest() (request *DescribeInstanceIdListRequest) {
	request = &DescribeInstanceIdListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInstanceIdList")
	return
}

func NewDescribeInstanceIdListResponse() (response *DescribeInstanceIdListResponse) {
	response = &DescribeInstanceIdListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按地域，返回实例id的列表；key为地域，值为实例列表
func (c *Client) DescribeInstanceIdList(request *DescribeInstanceIdListRequest) (response *DescribeInstanceIdListResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceIdListRequest()
	}
	response = NewDescribeInstanceIdListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMiniProgramPeakPointRequest() (request *DescribeMiniProgramPeakPointRequest) {
	request = &DescribeMiniProgramPeakPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeMiniProgramPeakPoint")
	return
}

func NewDescribeMiniProgramPeakPointResponse() (response *DescribeMiniProgramPeakPointResponse) {
	response = &DescribeMiniProgramPeakPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序各种趋势接口
func (c *Client) DescribeMiniProgramPeakPoint(request *DescribeMiniProgramPeakPointRequest) (response *DescribeMiniProgramPeakPointResponse, err error) {
	if request == nil {
		request = NewDescribeMiniProgramPeakPointRequest()
	}
	response = NewDescribeMiniProgramPeakPointResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOwaspWhiteRulesRequest() (request *DescribeOwaspWhiteRulesRequest) {
	request = &DescribeOwaspWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOwaspWhiteRules")
	return
}

func NewDescribeOwaspWhiteRulesResponse() (response *DescribeOwaspWhiteRulesResponse) {
	response = &DescribeOwaspWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则引擎白名单列表
func (c *Client) DescribeOwaspWhiteRules(request *DescribeOwaspWhiteRulesRequest) (response *DescribeOwaspWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDescribeOwaspWhiteRulesRequest()
	}
	response = NewDescribeOwaspWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleTypesEnRequest() (request *DescribeRuleTypesEnRequest) {
	request = &DescribeRuleTypesEnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRuleTypesEn")
	return
}

func NewDescribeRuleTypesEnResponse() (response *DescribeRuleTypesEnResponse) {
	response = &DescribeRuleTypesEnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则类型描述支持英文
func (c *Client) DescribeRuleTypesEn(request *DescribeRuleTypesEnRequest) (response *DescribeRuleTypesEnResponse, err error) {
	if request == nil {
		request = NewDescribeRuleTypesEnRequest()
	}
	response = NewDescribeRuleTypesEnResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRenewFlagRequest() (request *ModifyInstanceRenewFlagRequest) {
	request = &ModifyInstanceRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceRenewFlag")
	return
}

func NewModifyInstanceRenewFlagResponse() (response *ModifyInstanceRenewFlagResponse) {
	response = &ModifyInstanceRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的自动续费开关
func (c *Client) ModifyInstanceRenewFlag(request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRenewFlagRequest()
	}
	response = NewModifyInstanceRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMainClassRequest() (request *DescribeMainClassRequest) {
	request = &DescribeMainClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeMainClass")
	return
}

func NewDescribeMainClassResponse() (response *DescribeMainClassResponse) {
	response = &DescribeMainClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主类及子类信息
func (c *Client) DescribeMainClass(request *DescribeMainClassRequest) (response *DescribeMainClassResponse, err error) {
	if request == nil {
		request = NewDescribeMainClassRequest()
	}
	response = NewDescribeMainClassResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveBypassAllRuleRequest() (request *RemoveBypassAllRuleRequest) {
	request = &RemoveBypassAllRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RemoveBypassAllRule")
	return
}

func NewRemoveBypassAllRuleResponse() (response *RemoveBypassAllRuleResponse) {
	response = &RemoveBypassAllRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除一键bypass规则
func (c *Client) RemoveBypassAllRule(request *RemoveBypassAllRuleRequest) (response *RemoveBypassAllRuleResponse, err error) {
	if request == nil {
		request = NewRemoveBypassAllRuleRequest()
	}
	response = NewRemoveBypassAllRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotUCBFeatureRulesRequest() (request *CopyBotUCBFeatureRulesRequest) {
	request = &CopyBotUCBFeatureRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CopyBotUCBFeatureRules")
	return
}

func NewCopyBotUCBFeatureRulesResponse() (response *CopyBotUCBFeatureRulesResponse) {
	response = &CopyBotUCBFeatureRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB自定义策略拷贝
func (c *Client) CopyBotUCBFeatureRules(request *CopyBotUCBFeatureRulesRequest) (response *CopyBotUCBFeatureRulesResponse, err error) {
	if request == nil {
		request = NewCopyBotUCBFeatureRulesRequest()
	}
	response = NewCopyBotUCBFeatureRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSearchLogTemplateRequest() (request *DeleteSearchLogTemplateRequest) {
	request = &DeleteSearchLogTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteSearchLogTemplate")
	return
}

func NewDeleteSearchLogTemplateResponse() (response *DeleteSearchLogTemplateResponse) {
	response = &DeleteSearchLogTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志搜索模板
func (c *Client) DeleteSearchLogTemplate(request *DeleteSearchLogTemplateRequest) (response *DeleteSearchLogTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSearchLogTemplateRequest()
	}
	response = NewDeleteSearchLogTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBusinessRiskRequest() (request *CreateBusinessRiskRequest) {
	request = &CreateBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateBusinessRisk")
	return
}

func NewCreateBusinessRiskResponse() (response *CreateBusinessRiskResponse) {
	response = &CreateBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建业务安全规则
func (c *Client) CreateBusinessRisk(request *CreateBusinessRiskRequest) (response *CreateBusinessRiskResponse, err error) {
	if request == nil {
		request = NewCreateBusinessRiskRequest()
	}
	response = NewCreateBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLabelsRequest() (request *DescribeUserLabelsRequest) {
	request = &DescribeUserLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserLabels")
	return
}

func NewDescribeUserLabelsResponse() (response *DescribeUserLabelsResponse) {
	response = &DescribeUserLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的标签列表
func (c *Client) DescribeUserLabels(request *DescribeUserLabelsRequest) (response *DescribeUserLabelsResponse, err error) {
	if request == nil {
		request = NewDescribeUserLabelsRequest()
	}
	response = NewDescribeUserLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchIpAccessControlRequest() (request *ModifyBatchIpAccessControlRequest) {
	request = &ModifyBatchIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchIpAccessControl")
	return
}

func NewModifyBatchIpAccessControlResponse() (response *ModifyBatchIpAccessControlResponse) {
	response = &ModifyBatchIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量IP黑白名单新增接口
func (c *Client) ModifyBatchIpAccessControl(request *ModifyBatchIpAccessControlRequest) (response *ModifyBatchIpAccessControlResponse, err error) {
	if request == nil {
		request = NewModifyBatchIpAccessControlRequest()
	}
	response = NewModifyBatchIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccessConfigRequest() (request *DescribeUserAccessConfigRequest) {
	request = &DescribeUserAccessConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserAccessConfig")
	return
}

func NewDescribeUserAccessConfigResponse() (response *DescribeUserAccessConfigResponse) {
	response = &DescribeUserAccessConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户网关信息
func (c *Client) DescribeUserAccessConfig(request *DescribeUserAccessConfigRequest) (response *DescribeUserAccessConfigResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccessConfigRequest()
	}
	response = NewDescribeUserAccessConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMContentRuleStatusRequest() (request *ModifyLLMContentRuleStatusRequest) {
	request = &ModifyLLMContentRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMContentRuleStatus")
	return
}

func NewModifyLLMContentRuleStatusResponse() (response *ModifyLLMContentRuleStatusResponse) {
	response = &ModifyLLMContentRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换大模型内容防护规则开关
func (c *Client) ModifyLLMContentRuleStatus(request *ModifyLLMContentRuleStatusRequest) (response *ModifyLLMContentRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyLLMContentRuleStatusRequest()
	}
	response = NewModifyLLMContentRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRuleListRequest() (request *DescribeApiRuleListRequest) {
	request = &DescribeApiRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiRuleList")
	return
}

func NewDescribeApiRuleListResponse() (response *DescribeApiRuleListResponse) {
	response = &DescribeApiRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据多条件查询API规则，包括规则ID（精准搜索）/API参数（模糊搜索）/接口名称（模糊搜索）/执行动作（精准搜索）/请求方法（精准搜索）/来源（精准搜索）/开关状态（精准搜索）
func (c *Client) DescribeApiRuleList(request *DescribeApiRuleListRequest) (response *DescribeApiRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeApiRuleListRequest()
	}
	response = NewDescribeApiRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGradeRequest() (request *DescribeUserGradeRequest) {
	request = &DescribeUserGradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserGrade")
	return
}

func NewDescribeUserGradeResponse() (response *DescribeUserGradeResponse) {
	response = &DescribeUserGradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户Uin查询用户的等级和类型
func (c *Client) DescribeUserGrade(request *DescribeUserGradeRequest) (response *DescribeUserGradeResponse, err error) {
	if request == nil {
		request = NewDescribeUserGradeRequest()
	}
	response = NewDescribeUserGradeResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaPackageRenewRequest() (request *ModifySpartaPackageRenewRequest) {
	request = &ModifySpartaPackageRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifySpartaPackageRenew")
	return
}

func NewModifySpartaPackageRenewResponse() (response *ModifySpartaPackageRenewResponse) {
	response = &ModifySpartaPackageRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Sparta版本WAF自动续费开关设置
func (c *Client) ModifySpartaPackageRenew(request *ModifySpartaPackageRenewRequest) (response *ModifySpartaPackageRenewResponse, err error) {
	if request == nil {
		request = NewModifySpartaPackageRenewRequest()
	}
	response = NewModifySpartaPackageRenewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCRuleRequest() (request *DescribeCCRuleRequest) {
	request = &DescribeCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCCRule")
	return
}

func NewDescribeCCRuleResponse() (response *DescribeCCRuleResponse) {
	response = &DescribeCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Query接口
func (c *Client) DescribeCCRule(request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCCRuleRequest()
	}
	response = NewDescribeCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVIPRSRequest() (request *ModifyVIPRSRequest) {
	request = &ModifyVIPRSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyVIPRS")
	return
}

func NewModifyVIPRSResponse() (response *ModifyVIPRSResponse) {
	response = &ModifyVIPRSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移流量
func (c *Client) ModifyVIPRS(request *ModifyVIPRSRequest) (response *ModifyVIPRSResponse, err error) {
	if request == nil {
		request = NewModifyVIPRSRequest()
	}
	response = NewModifyVIPRSResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchIpAccessControlRequest() (request *DescribeBatchIpAccessControlRequest) {
	request = &DescribeBatchIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchIpAccessControl")
	return
}

func NewDescribeBatchIpAccessControlResponse() (response *DescribeBatchIpAccessControlResponse) {
	response = &DescribeBatchIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 多域名ip黑白名单查询
func (c *Client) DescribeBatchIpAccessControl(request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDescribeBatchIpAccessControlRequest()
	}
	response = NewDescribeBatchIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostInnerRequest() (request *CreateHostInnerRequest) {
	request = &CreateHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateHostInner")
	return
}

func NewCreateHostInnerResponse() (response *CreateHostInnerResponse) {
	response = &CreateHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中添加防护的域名
func (c *Client) CreateHostInner(request *CreateHostInnerRequest) (response *CreateHostInnerResponse, err error) {
	if request == nil {
		request = NewCreateHostInnerRequest()
	}
	response = NewCreateHostInnerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoDenyIPRequest() (request *DescribeAutoDenyIPRequest) {
	request = &DescribeAutoDenyIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAutoDenyIP")
	return
}

func NewDescribeAutoDenyIPResponse() (response *DescribeAutoDenyIPResponse) {
	response = &DescribeAutoDenyIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF自动封禁IP详情,对齐自动封堵状态
func (c *Client) DescribeAutoDenyIP(request *DescribeAutoDenyIPRequest) (response *DescribeAutoDenyIPResponse, err error) {
	if request == nil {
		request = NewDescribeAutoDenyIPRequest()
	}
	response = NewDescribeAutoDenyIPResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserAccessConfigRequest() (request *ModifyUserAccessConfigRequest) {
	request = &ModifyUserAccessConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserAccessConfig")
	return
}

func NewModifyUserAccessConfigResponse() (response *ModifyUserAccessConfigResponse) {
	response = &ModifyUserAccessConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户网关配置
func (c *Client) ModifyUserAccessConfig(request *ModifyUserAccessConfigRequest) (response *ModifyUserAccessConfigResponse, err error) {
	if request == nil {
		request = NewModifyUserAccessConfigRequest()
	}
	response = NewModifyUserAccessConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGenerateDealsRequest() (request *ModifyGenerateDealsRequest) {
	request = &ModifyGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyGenerateDeals")
	return
}

func NewModifyGenerateDealsResponse() (response *ModifyGenerateDealsResponse) {
	response = &ModifyGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例下单接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费接口下单。目前只支持预付费下单，计费侧接口：https://tcb.xxx.com/magical-brush/docs/754661947
func (c *Client) ModifyGenerateDeals(request *ModifyGenerateDealsRequest) (response *ModifyGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyGenerateDealsRequest()
	}
	response = NewModifyGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAIModelRequest() (request *ModifyAIModelRequest) {
	request = &ModifyAIModelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAIModel")
	return
}

func NewModifyAIModelResponse() (response *ModifyAIModelResponse) {
	response = &ModifyAIModelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启ai模型训练
func (c *Client) ModifyAIModel(request *ModifyAIModelRequest) (response *ModifyAIModelResponse, err error) {
	if request == nil {
		request = NewModifyAIModelRequest()
	}
	response = NewModifyAIModelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpSignatureRuleRequest() (request *DescribeOpSignatureRuleRequest) {
	request = &DescribeOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpSignatureRule")
	return
}

func NewDescribeOpSignatureRuleResponse() (response *DescribeOpSignatureRuleResponse) {
	response = &DescribeOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口，获取全量特征规则列表
func (c *Client) DescribeOpSignatureRule(request *DescribeOpSignatureRuleRequest) (response *DescribeOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpSignatureRuleRequest()
	}
	response = NewDescribeOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClsForAfterpayRequest() (request *CreateClsForAfterpayRequest) {
	request = &CreateClsForAfterpayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateClsForAfterpay")
	return
}

func NewCreateClsForAfterpayResponse() (response *CreateClsForAfterpayResponse) {
	response = &CreateClsForAfterpayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费创建日志服务
func (c *Client) CreateClsForAfterpay(request *CreateClsForAfterpayRequest) (response *CreateClsForAfterpayResponse, err error) {
	if request == nil {
		request = NewCreateClsForAfterpayRequest()
	}
	response = NewCreateClsForAfterpayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInactivationApiRequest() (request *DescribeInactivationApiRequest) {
	request = &DescribeInactivationApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInactivationApi")
	return
}

func NewDescribeInactivationApiResponse() (response *DescribeInactivationApiResponse) {
	response = &DescribeInactivationApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全失活api查询接口
func (c *Client) DescribeInactivationApi(request *DescribeInactivationApiRequest) (response *DescribeInactivationApiResponse, err error) {
	if request == nil {
		request = NewDescribeInactivationApiRequest()
	}
	response = NewDescribeInactivationApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQpsWhiteListRequest() (request *DescribeQpsWhiteListRequest) {
	request = &DescribeQpsWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeQpsWhiteList")
	return
}

func NewDescribeQpsWhiteListResponse() (response *DescribeQpsWhiteListResponse) {
	response = &DescribeQpsWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务扩展包黑白名单
func (c *Client) DescribeQpsWhiteList(request *DescribeQpsWhiteListRequest) (response *DescribeQpsWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeQpsWhiteListRequest()
	}
	response = NewDescribeQpsWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
	request = &DescribeRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRoles")
	return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
	response = &DescribeRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查CAM角色服务是否开通
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRolesRequest()
	}
	response = NewDescribeRolesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncClbAssetsProcessRequest() (request *SyncClbAssetsProcessRequest) {
	request = &SyncClbAssetsProcessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SyncClbAssetsProcess")
	return
}

func NewSyncClbAssetsProcessResponse() (response *SyncClbAssetsProcessResponse) {
	response = &SyncClbAssetsProcessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产同步任务状态
func (c *Client) SyncClbAssetsProcess(request *SyncClbAssetsProcessRequest) (response *SyncClbAssetsProcessResponse, err error) {
	if request == nil {
		request = NewSyncClbAssetsProcessRequest()
	}
	response = NewSyncClbAssetsProcessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackWhiteRuleRequest() (request *DescribeAttackWhiteRuleRequest) {
	request = &DescribeAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackWhiteRule")
	return
}

func NewDescribeAttackWhiteRuleResponse() (response *DescribeAttackWhiteRuleResponse) {
	response = &DescribeAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户规则白名单列表
func (c *Client) DescribeAttackWhiteRule(request *DescribeAttackWhiteRuleRequest) (response *DescribeAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeAttackWhiteRuleRequest()
	}
	response = NewDescribeAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectionLevelRequest() (request *ModifyProtectionLevelRequest) {
	request = &ModifyProtectionLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyProtectionLevel")
	return
}

func NewModifyProtectionLevelResponse() (response *ModifyProtectionLevelResponse) {
	response = &ModifyProtectionLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改防护等级
func (c *Client) ModifyProtectionLevel(request *ModifyProtectionLevelRequest) (response *ModifyProtectionLevelResponse, err error) {
	if request == nil {
		request = NewModifyProtectionLevelRequest()
	}
	response = NewModifyProtectionLevelResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyAreaBanRequest() (request *CreateCopyAreaBanRequest) {
	request = &CreateCopyAreaBanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateCopyAreaBan")
	return
}

func NewCreateCopyAreaBanResponse() (response *CreateCopyAreaBanResponse) {
	response = &CreateCopyAreaBanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制域名的地域封禁到其他域名，其他域名的地域封禁规则会被清空
func (c *Client) CreateCopyAreaBan(request *CreateCopyAreaBanRequest) (response *CreateCopyAreaBanResponse, err error) {
	if request == nil {
		request = NewCreateCopyAreaBanRequest()
	}
	response = NewCreateCopyAreaBanResponse()
	err = c.Send(request, response)
	return
}

func NewCopyCustomRuleRequest() (request *CopyCustomRuleRequest) {
	request = &CopyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CopyCustomRule")
	return
}

func NewCopyCustomRuleResponse() (response *CopyCustomRuleResponse) {
	response = &CopyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从一个域名复制一条自定义策略到其他域名
func (c *Client) CopyCustomRule(request *CopyCustomRuleRequest) (response *CopyCustomRuleResponse, err error) {
	if request == nil {
		request = NewCopyCustomRuleRequest()
	}
	response = NewCopyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMRuleInfoRequest() (request *DescribeLLMRuleInfoRequest) {
	request = &DescribeLLMRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMRuleInfo")
	return
}

func NewDescribeLLMRuleInfoResponse() (response *DescribeLLMRuleInfoResponse) {
	response = &DescribeLLMRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取大模型内容检测匹配内容信息
func (c *Client) DescribeLLMRuleInfo(request *DescribeLLMRuleInfoRequest) (response *DescribeLLMRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLLMRuleInfoRequest()
	}
	response = NewDescribeLLMRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiFakeUrlRequest() (request *ModifyAntiFakeUrlRequest) {
	request = &ModifyAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAntiFakeUrl")
	return
}

func NewModifyAntiFakeUrlResponse() (response *ModifyAntiFakeUrlResponse) {
	response = &ModifyAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑防篡改url
func (c *Client) ModifyAntiFakeUrl(request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewModifyAntiFakeUrlRequest()
	}
	response = NewModifyAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewAddCaptchaCfgRequest() (request *AddCaptchaCfgRequest) {
	request = &AddCaptchaCfgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddCaptchaCfg")
	return
}

func NewAddCaptchaCfgResponse() (response *AddCaptchaCfgResponse) {
	response = &AddCaptchaCfgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置域名的验证码配置
func (c *Client) AddCaptchaCfg(request *AddCaptchaCfgRequest) (response *AddCaptchaCfgResponse, err error) {
	if request == nil {
		request = NewAddCaptchaCfgRequest()
	}
	response = NewAddCaptchaCfgResponse()
	err = c.Send(request, response)
	return
}

func NewFreshAntiFakeUrlRequest() (request *FreshAntiFakeUrlRequest) {
	request = &FreshAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "FreshAntiFakeUrl")
	return
}

func NewFreshAntiFakeUrlResponse() (response *FreshAntiFakeUrlResponse) {
	response = &FreshAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新防篡改url
func (c *Client) FreshAntiFakeUrl(request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewFreshAntiFakeUrlRequest()
	}
	response = NewFreshAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomWhiteRuleRequest() (request *ModifyCustomWhiteRuleRequest) {
	request = &ModifyCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCustomWhiteRule")
	return
}

func NewModifyCustomWhiteRuleResponse() (response *ModifyCustomWhiteRuleResponse) {
	response = &ModifyCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑精准白名单
func (c *Client) ModifyCustomWhiteRule(request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyCustomWhiteRuleRequest()
	}
	response = NewModifyCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddApiRuleRequest() (request *AddApiRuleRequest) {
	request = &AddApiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddApiRule")
	return
}

func NewAddApiRuleResponse() (response *AddApiRuleResponse) {
	response = &AddApiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户手动录入api规则
func (c *Client) AddApiRule(request *AddApiRuleRequest) (response *AddApiRuleResponse, err error) {
	if request == nil {
		request = NewAddApiRuleRequest()
	}
	response = NewAddApiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessPeriodRequest() (request *ModifyAccessPeriodRequest) {
	request = &ModifyAccessPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAccessPeriod")
	return
}

func NewModifyAccessPeriodResponse() (response *ModifyAccessPeriodResponse) {
	response = &ModifyAccessPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改访问日志保存期限等字段信息
func (c *Client) ModifyAccessPeriod(request *ModifyAccessPeriodRequest) (response *ModifyAccessPeriodResponse, err error) {
	if request == nil {
		request = NewModifyAccessPeriodRequest()
	}
	response = NewModifyAccessPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWhiteListRequest() (request *CreateWhiteListRequest) {
	request = &CreateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateWhiteList")
	return
}

func NewCreateWhiteListResponse() (response *CreateWhiteListResponse) {
	response = &CreateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增业务白名单，目前是前端对抗规则中的白名单
func (c *Client) CreateWhiteList(request *CreateWhiteListRequest) (response *CreateWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateWhiteListRequest()
	}
	response = NewCreateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainPostArgsRequest() (request *ModifyDomainPostArgsRequest) {
	request = &ModifyDomainPostArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainPostArgs")
	return
}

func NewModifyDomainPostArgsResponse() (response *ModifyDomainPostArgsResponse) {
	response = &ModifyDomainPostArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [WAF请求体参数长度下发]-编辑接口
func (c *Client) ModifyDomainPostArgs(request *ModifyDomainPostArgsRequest) (response *ModifyDomainPostArgsResponse, err error) {
	if request == nil {
		request = NewModifyDomainPostArgsRequest()
	}
	response = NewModifyDomainPostArgsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletTaskInfoRequest() (request *ModifyAppletTaskInfoRequest) {
	request = &ModifyAppletTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletTaskInfo")
	return
}

func NewModifyAppletTaskInfoResponse() (response *ModifyAppletTaskInfoResponse) {
	response = &ModifyAppletTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改小程序压测任务预约信息(运营平台)
func (c *Client) ModifyAppletTaskInfo(request *ModifyAppletTaskInfoRequest) (response *ModifyAppletTaskInfoResponse, err error) {
	if request == nil {
		request = NewModifyAppletTaskInfoRequest()
	}
	response = NewModifyAppletTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyKeepProtectionConfRequest() (request *ModifyKeepProtectionConfRequest) {
	request = &ModifyKeepProtectionConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyKeepProtectionConf")
	return
}

func NewModifyKeepProtectionConfResponse() (response *ModifyKeepProtectionConfResponse) {
	response = &ModifyKeepProtectionConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置域名删除时是否保留防护配置
func (c *Client) ModifyKeepProtectionConf(request *ModifyKeepProtectionConfRequest) (response *ModifyKeepProtectionConfResponse, err error) {
	if request == nil {
		request = NewModifyKeepProtectionConfRequest()
	}
	response = NewModifyKeepProtectionConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostLimitRequest() (request *DescribeHostLimitRequest) {
	request = &DescribeHostLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHostLimit")
	return
}

func NewDescribeHostLimitResponse() (response *DescribeHostLimitResponse) {
	response = &DescribeHostLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加域名的首先验证是否购买了套餐，是否没有达到购买套餐的限制，域名是否已经添加
func (c *Client) DescribeHostLimit(request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
	if request == nil {
		request = NewDescribeHostLimitRequest()
	}
	response = NewDescribeHostLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomRulesRequest() (request *DescribeCustomRulesRequest) {
	request = &DescribeCustomRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomRules")
	return
}

func NewDescribeCustomRulesResponse() (response *DescribeCustomRulesResponse) {
	response = &DescribeCustomRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的自定义策略列表
func (c *Client) DescribeCustomRules(request *DescribeCustomRulesRequest) (response *DescribeCustomRulesResponse, err error) {
	if request == nil {
		request = NewDescribeCustomRulesRequest()
	}
	response = NewDescribeCustomRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
	request = &DescribeLogHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLogHistogram")
	return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
	response = &DescribeLogHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于构建日志数量直方图
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeLogHistogramRequest()
	}
	response = NewDescribeLogHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomWhiteRuleRequest() (request *DeleteCustomWhiteRuleRequest) {
	request = &DeleteCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCustomWhiteRule")
	return
}

func NewDeleteCustomWhiteRuleResponse() (response *DeleteCustomWhiteRuleResponse) {
	response = &DeleteCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除精准白名单规则
func (c *Client) DeleteCustomWhiteRule(request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCustomWhiteRuleRequest()
	}
	response = NewDeleteCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAiRuleRequest() (request *DescribeBotAiRuleRequest) {
	request = &DescribeBotAiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAiRule")
	return
}

func NewDescribeBotAiRuleResponse() (response *DescribeBotAiRuleResponse) {
	response = &DescribeBotAiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot-ai的配置信息
func (c *Client) DescribeBotAiRule(request *DescribeBotAiRuleRequest) (response *DescribeBotAiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotAiRuleRequest()
	}
	response = NewDescribeBotAiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessMetaRequest() (request *DescribeAccessMetaRequest) {
	request = &DescribeAccessMetaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessMeta")
	return
}

func NewDescribeAccessMetaResponse() (response *DescribeAccessMetaResponse) {
	response = &DescribeAccessMetaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志元数据信息
func (c *Client) DescribeAccessMeta(request *DescribeAccessMetaRequest) (response *DescribeAccessMetaResponse, err error) {
	if request == nil {
		request = NewDescribeAccessMetaRequest()
	}
	response = NewDescribeAccessMetaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotActionStatRequest() (request *DescribeBotActionStatRequest) {
	request = &DescribeBotActionStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotActionStat")
	return
}

func NewDescribeBotActionStatResponse() (response *DescribeBotActionStatResponse) {
	response = &DescribeBotActionStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取bot动作统计
func (c *Client) DescribeBotActionStat(request *DescribeBotActionStatRequest) (response *DescribeBotActionStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotActionStatRequest()
	}
	response = NewDescribeBotActionStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexRequest() (request *DescribeIndexRequest) {
	request = &DescribeIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeIndex")
	return
}

func NewDescribeIndexResponse() (response *DescribeIndexResponse) {
	response = &DescribeIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取索引配置信息
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
	if request == nil {
		request = NewDescribeIndexRequest()
	}
	response = NewDescribeIndexResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainsCLSStatusRequest() (request *ModifyDomainsCLSStatusRequest) {
	request = &ModifyDomainsCLSStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainsCLSStatus")
	return
}

func NewModifyDomainsCLSStatusResponse() (response *ModifyDomainsCLSStatusResponse) {
	response = &ModifyDomainsCLSStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名列表的访问日志开关
func (c *Client) ModifyDomainsCLSStatus(request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
	if request == nil {
		request = NewModifyDomainsCLSStatusRequest()
	}
	response = NewModifyDomainsCLSStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCOpenCustomCountRequest() (request *DescribeCCOpenCustomCountRequest) {
	request = &DescribeCCOpenCustomCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCCOpenCustomCount")
	return
}

func NewDescribeCCOpenCustomCountResponse() (response *DescribeCCOpenCustomCountResponse) {
	response = &DescribeCCOpenCustomCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本查询自定义cc开启状态数目
func (c *Client) DescribeCCOpenCustomCount(request *DescribeCCOpenCustomCountRequest) (response *DescribeCCOpenCustomCountResponse, err error) {
	if request == nil {
		request = NewDescribeCCOpenCustomCountRequest()
	}
	response = NewDescribeCCOpenCustomCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsUsedRequest() (request *DescribeClsUsedRequest) {
	request = &DescribeClsUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeClsUsed")
	return
}

func NewDescribeClsUsedResponse() (response *DescribeClsUsedResponse) {
	response = &DescribeClsUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回日志用量
func (c *Client) DescribeClsUsed(request *DescribeClsUsedRequest) (response *DescribeClsUsedResponse, err error) {
	if request == nil {
		request = NewDescribeClsUsedRequest()
	}
	response = NewDescribeClsUsedResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAttackLogRequest() (request *SearchAttackLogRequest) {
	request = &SearchAttackLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SearchAttackLog")
	return
}

func NewSearchAttackLogResponse() (response *SearchAttackLogResponse) {
	response = &SearchAttackLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新版本CLS接口存在参数变化，query改成了query_string支持lucence语法接口搜索查询。
func (c *Client) SearchAttackLog(request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
	if request == nil {
		request = NewSearchAttackLogRequest()
	}
	response = NewSearchAttackLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatisticPointsRequest() (request *DescribeBotStatisticPointsRequest) {
	request = &DescribeBotStatisticPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotStatisticPoints")
	return
}

func NewDescribeBotStatisticPointsResponse() (response *DescribeBotStatisticPointsResponse) {
	response = &DescribeBotStatisticPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 Bot流量统计
func (c *Client) DescribeBotStatisticPoints(request *DescribeBotStatisticPointsRequest) (response *DescribeBotStatisticPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatisticPointsRequest()
	}
	response = NewDescribeBotStatisticPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiStatCountRequest() (request *DescribeApiStatCountRequest) {
	request = &DescribeApiStatCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiStatCount")
	return
}

func NewDescribeApiStatCountResponse() (response *DescribeApiStatCountResponse) {
	response = &DescribeApiStatCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量体检-控制台-查询api资产各类统计
func (c *Client) DescribeApiStatCount(request *DescribeApiStatCountRequest) (response *DescribeApiStatCountResponse, err error) {
	if request == nil {
		request = NewDescribeApiStatCountRequest()
	}
	response = NewDescribeApiStatCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例询价接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费询价接口。计费侧询价接口：https://tcb.xxx.com/magical-brush/docs/746616075
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	response = NewDescribePriceResponse()
	err = c.Send(request, response)
	return
}

func NewAddBypassAllRuleRequest() (request *AddBypassAllRuleRequest) {
	request = &AddBypassAllRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBypassAllRule")
	return
}

func NewAddBypassAllRuleResponse() (response *AddBypassAllRuleResponse) {
	response = &AddBypassAllRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加一键bypass能力支持,直接添加APPID
func (c *Client) AddBypassAllRule(request *AddBypassAllRuleRequest) (response *AddBypassAllRuleResponse, err error) {
	if request == nil {
		request = NewAddBypassAllRuleRequest()
	}
	response = NewAddBypassAllRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainRulesRequest() (request *DescribeDomainRulesRequest) {
	request = &DescribeDomainRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainRules")
	return
}

func NewDescribeDomainRulesResponse() (response *DescribeDomainRulesResponse) {
	response = &DescribeDomainRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取域名的防护规则列表
func (c *Client) DescribeDomainRules(request *DescribeDomainRulesRequest) (response *DescribeDomainRulesResponse, err error) {
	if request == nil {
		request = NewDescribeDomainRulesRequest()
	}
	response = NewDescribeDomainRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainSecurityGroupTipsCountRequest() (request *DescribeDomainSecurityGroupTipsCountRequest) {
	request = &DescribeDomainSecurityGroupTipsCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainSecurityGroupTipsCount")
	return
}

func NewDescribeDomainSecurityGroupTipsCountResponse() (response *DescribeDomainSecurityGroupTipsCountResponse) {
	response = &DescribeDomainSecurityGroupTipsCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全量域名安全组绑定失败或安全组变更的数量
func (c *Client) DescribeDomainSecurityGroupTipsCount(request *DescribeDomainSecurityGroupTipsCountRequest) (response *DescribeDomainSecurityGroupTipsCountResponse, err error) {
	if request == nil {
		request = NewDescribeDomainSecurityGroupTipsCountRequest()
	}
	response = NewDescribeDomainSecurityGroupTipsCountResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllQpsConfigRequest() (request *GetAllQpsConfigRequest) {
	request = &GetAllQpsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAllQpsConfig")
	return
}

func NewGetAllQpsConfigResponse() (response *GetAllQpsConfigResponse) {
	response = &GetAllQpsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有套餐类型对应的弹性qps配置
func (c *Client) GetAllQpsConfig(request *GetAllQpsConfigRequest) (response *GetAllQpsConfigResponse, err error) {
	if request == nil {
		request = NewGetAllQpsConfigRequest()
	}
	response = NewGetAllQpsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiSecurityGenerateDealsRequest() (request *ModifyApiSecurityGenerateDealsRequest) {
	request = &ModifyApiSecurityGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiSecurityGenerateDeals")
	return
}

func NewModifyApiSecurityGenerateDealsResponse() (response *ModifyApiSecurityGenerateDealsResponse) {
	response = &ModifyApiSecurityGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给使用的waf实例 api安全下单接口
func (c *Client) ModifyApiSecurityGenerateDeals(request *ModifyApiSecurityGenerateDealsRequest) (response *ModifyApiSecurityGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyApiSecurityGenerateDealsRequest()
	}
	response = NewModifyApiSecurityGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseCanBuyTceRequest() (request *DescribeLicenseCanBuyTceRequest) {
	request = &DescribeLicenseCanBuyTceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLicenseCanBuyTce")
	return
}

func NewDescribeLicenseCanBuyTceResponse() (response *DescribeLicenseCanBuyTceResponse) {
	response = &DescribeLicenseCanBuyTceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf是否可购买
func (c *Client) DescribeLicenseCanBuyTce(request *DescribeLicenseCanBuyTceRequest) (response *DescribeLicenseCanBuyTceResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseCanBuyTceRequest()
	}
	response = NewDescribeLicenseCanBuyTceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainResponsePageRequest() (request *DescribeDomainResponsePageRequest) {
	request = &DescribeDomainResponsePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainResponsePage")
	return
}

func NewDescribeDomainResponsePageResponse() (response *DescribeDomainResponsePageResponse) {
	response = &DescribeDomainResponsePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示指定域名的拦截页面信息
func (c *Client) DescribeDomainResponsePage(request *DescribeDomainResponsePageRequest) (response *DescribeDomainResponsePageResponse, err error) {
	if request == nil {
		request = NewDescribeDomainResponsePageRequest()
	}
	response = NewDescribeDomainResponsePageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTceWafBotResourceRequest() (request *DeleteTceWafBotResourceRequest) {
	request = &DeleteTceWafBotResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteTceWafBotResource")
	return
}

func NewDeleteTceWafBotResourceResponse() (response *DeleteTceWafBotResourceResponse) {
	response = &DeleteTceWafBotResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退订WAF-BOT
func (c *Client) DeleteTceWafBotResource(request *DeleteTceWafBotResourceRequest) (response *DeleteTceWafBotResourceResponse, err error) {
	if request == nil {
		request = NewDeleteTceWafBotResourceRequest()
	}
	response = NewDeleteTceWafBotResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventTypeListRequest() (request *DescribeApiSecEventTypeListRequest) {
	request = &DescribeApiSecEventTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecEventTypeList")
	return
}

func NewDescribeApiSecEventTypeListResponse() (response *DescribeApiSecEventTypeListResponse) {
	response = &DescribeApiSecEventTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全事件分类列表
func (c *Client) DescribeApiSecEventTypeList(request *DescribeApiSecEventTypeListRequest) (response *DescribeApiSecEventTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventTypeListRequest()
	}
	response = NewDescribeApiSecEventTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostFlowServiceRequest() (request *DescribePostFlowServiceRequest) {
	request = &DescribePostFlowServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePostFlowService")
	return
}

func NewDescribePostFlowServiceResponse() (response *DescribePostFlowServiceResponse) {
	response = &DescribePostFlowServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查开通投递服务接口
func (c *Client) DescribePostFlowService(request *DescribePostFlowServiceRequest) (response *DescribePostFlowServiceResponse, err error) {
	if request == nil {
		request = NewDescribePostFlowServiceRequest()
	}
	response = NewDescribePostFlowServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletRiskStatusRequest() (request *DescribeAppletRiskStatusRequest) {
	request = &DescribeAppletRiskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletRiskStatus")
	return
}

func NewDescribeAppletRiskStatusResponse() (response *DescribeAppletRiskStatusResponse) {
	response = &DescribeAppletRiskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询小程序风控状态接口
func (c *Client) DescribeAppletRiskStatus(request *DescribeAppletRiskStatusRequest) (response *DescribeAppletRiskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAppletRiskStatusRequest()
	}
	response = NewDescribeAppletRiskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafAntiDDoSEipRequest() (request *ModifyWafAntiDDoSEipRequest) {
	request = &ModifyWafAntiDDoSEipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWafAntiDDoSEip")
	return
}

func NewModifyWafAntiDDoSEipResponse() (response *ModifyWafAntiDDoSEipResponse) {
	response = &ModifyWafAntiDDoSEipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给高防EIP调用，绑定解绑WAF实例
func (c *Client) ModifyWafAntiDDoSEip(request *ModifyWafAntiDDoSEipRequest) (response *ModifyWafAntiDDoSEipResponse, err error) {
	if request == nil {
		request = NewModifyWafAntiDDoSEipRequest()
	}
	response = NewModifyWafAntiDDoSEipResponse()
	err = c.Send(request, response)
	return
}

func NewGetStracePointsRequest() (request *GetStracePointsRequest) {
	request = &GetStracePointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetStracePoints")
	return
}

func NewGetStracePointsResponse() (response *GetStracePointsResponse) {
	response = &GetStracePointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护量趋势图
func (c *Client) GetStracePoints(request *GetStracePointsRequest) (response *GetStracePointsResponse, err error) {
	if request == nil {
		request = NewGetStracePointsRequest()
	}
	response = NewGetStracePointsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotEventStateRequest() (request *ModifyBotEventStateRequest) {
	request = &ModifyBotEventStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotEventState")
	return
}

func NewModifyBotEventStateResponse() (response *ModifyBotEventStateResponse) {
	response = &ModifyBotEventStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-修改事件状态
func (c *Client) ModifyBotEventState(request *ModifyBotEventStateRequest) (response *ModifyBotEventStateResponse, err error) {
	if request == nil {
		request = NewModifyBotEventStateRequest()
	}
	response = NewModifyBotEventStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessExportRequest() (request *DeleteAccessExportRequest) {
	request = &DeleteAccessExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAccessExport")
	return
}

func NewDeleteAccessExportResponse() (response *DeleteAccessExportResponse) {
	response = &DeleteAccessExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除访问日志导出
func (c *Client) DeleteAccessExport(request *DeleteAccessExportRequest) (response *DeleteAccessExportResponse, err error) {
	if request == nil {
		request = NewDeleteAccessExportRequest()
	}
	response = NewDeleteAccessExportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAllGenerateDealsRequest() (request *ModifyAllGenerateDealsRequest) {
	request = &ModifyAllGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAllGenerateDeals")
	return
}

func NewModifyAllGenerateDealsResponse() (response *ModifyAllGenerateDealsResponse) {
	response = &ModifyAllGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给clb等使用的waf实例下单接口，目前只支持clb旗舰版实例的下单，该接口会进行入参校验，然后调用是否为收购用户，然后调用计费接口下单。目前只支持预付费下单，计费侧接口：https://tcb.xxx.com/magical-brush/docs/754661947
func (c *Client) ModifyAllGenerateDeals(request *ModifyAllGenerateDealsRequest) (response *ModifyAllGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyAllGenerateDealsRequest()
	}
	response = NewModifyAllGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchAreaBanRequest() (request *DeleteBatchAreaBanRequest) {
	request = &DeleteBatchAreaBanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBatchAreaBan")
	return
}

func NewDeleteBatchAreaBanResponse() (response *DeleteBatchAreaBanResponse) {
	response = &DeleteBatchAreaBanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-批量删除接口
func (c *Client) DeleteBatchAreaBan(request *DeleteBatchAreaBanRequest) (response *DeleteBatchAreaBanResponse, err error) {
	if request == nil {
		request = NewDeleteBatchAreaBanRequest()
	}
	response = NewDeleteBatchAreaBanResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafAutoDenyRulesRequest() (request *ModifyWafAutoDenyRulesRequest) {
	request = &ModifyWafAutoDenyRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWafAutoDenyRules")
	return
}

func NewModifyWafAutoDenyRulesResponse() (response *ModifyWafAutoDenyRulesResponse) {
	response = &ModifyWafAutoDenyRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ip惩罚规则
func (c *Client) ModifyWafAutoDenyRules(request *ModifyWafAutoDenyRulesRequest) (response *ModifyWafAutoDenyRulesResponse, err error) {
	if request == nil {
		request = NewModifyWafAutoDenyRulesRequest()
	}
	response = NewModifyWafAutoDenyRulesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRateLimitV2Request() (request *UpdateRateLimitV2Request) {
	request = &UpdateRateLimitV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpdateRateLimitV2")
	return
}

func NewUpdateRateLimitV2Response() (response *UpdateRateLimitV2Response) {
	response = &UpdateRateLimitV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新自研版限流规则
func (c *Client) UpdateRateLimitV2(request *UpdateRateLimitV2Request) (response *UpdateRateLimitV2Response, err error) {
	if request == nil {
		request = NewUpdateRateLimitV2Request()
	}
	response = NewUpdateRateLimitV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBotGlobalConfigRequest() (request *DescribeBotGlobalConfigRequest) {
	request = &DescribeBotGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotGlobalConfig")
	return
}

func NewDescribeBotGlobalConfigResponse() (response *DescribeBotGlobalConfigResponse) {
	response = &DescribeBotGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT全局设置信息
func (c *Client) DescribeBotGlobalConfig(request *DescribeBotGlobalConfigRequest) (response *DescribeBotGlobalConfigResponse, err error) {
	if request == nil {
		request = NewDescribeBotGlobalConfigRequest()
	}
	response = NewDescribeBotGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleUpdateLogRequest() (request *DescribeRuleUpdateLogRequest) {
	request = &DescribeRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRuleUpdateLog")
	return
}

func NewDescribeRuleUpdateLogResponse() (response *DescribeRuleUpdateLogResponse) {
	response = &DescribeRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取特征规则更新动态
func (c *Client) DescribeRuleUpdateLog(request *DescribeRuleUpdateLogRequest) (response *DescribeRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDescribeRuleUpdateLogRequest()
	}
	response = NewDescribeRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchCustomRuleRequest() (request *ModifyBatchCustomRuleRequest) {
	request = &ModifyBatchCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchCustomRule")
	return
}

func NewModifyBatchCustomRuleResponse() (response *ModifyBatchCustomRuleResponse) {
	response = &ModifyBatchCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-批量编辑接口
func (c *Client) ModifyBatchCustomRule(request *ModifyBatchCustomRuleRequest) (response *ModifyBatchCustomRuleResponse, err error) {
	if request == nil {
		request = NewModifyBatchCustomRuleRequest()
	}
	response = NewModifyBatchCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiInfoLeakRuleStatusRequest() (request *ModifyAntiInfoLeakRuleStatusRequest) {
	request = &ModifyAntiInfoLeakRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAntiInfoLeakRuleStatus")
	return
}

func NewModifyAntiInfoLeakRuleStatusResponse() (response *ModifyAntiInfoLeakRuleStatusResponse) {
	response = &ModifyAntiInfoLeakRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 信息防泄漏切换规则开关
func (c *Client) ModifyAntiInfoLeakRuleStatus(request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyAntiInfoLeakRuleStatusRequest()
	}
	response = NewModifyAntiInfoLeakRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVIPRSRequest() (request *DescribeVIPRSRequest) {
	request = &DescribeVIPRSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeVIPRS")
	return
}

func NewDescribeVIPRSResponse() (response *DescribeVIPRSResponse) {
	response = &DescribeVIPRSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询详细的vip信息
func (c *Client) DescribeVIPRS(request *DescribeVIPRSRequest) (response *DescribeVIPRSResponse, err error) {
	if request == nil {
		request = NewDescribeVIPRSRequest()
	}
	response = NewDescribeVIPRSResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadBotTldSdkRequest() (request *DownloadBotTldSdkRequest) {
	request = &DownloadBotTldSdkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DownloadBotTldSdk")
	return
}

func NewDownloadBotTldSdkResponse() (response *DownloadBotTldSdkResponse) {
	response = &DownloadBotTldSdkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载图灵的sdk
func (c *Client) DownloadBotTldSdk(request *DownloadBotTldSdkRequest) (response *DownloadBotTldSdkResponse, err error) {
	if request == nil {
		request = NewDownloadBotTldSdkRequest()
	}
	response = NewDownloadBotTldSdkResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainPostActionRequest() (request *ModifyDomainPostActionRequest) {
	request = &ModifyDomainPostActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainPostAction")
	return
}

func NewModifyDomainPostActionResponse() (response *ModifyDomainPostActionResponse) {
	response = &ModifyDomainPostActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改域名投递状态
func (c *Client) ModifyDomainPostAction(request *ModifyDomainPostActionRequest) (response *ModifyDomainPostActionResponse, err error) {
	if request == nil {
		request = NewModifyDomainPostActionRequest()
	}
	response = NewModifyDomainPostActionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyThreatIntelligenceStatusRequest() (request *ModifyThreatIntelligenceStatusRequest) {
	request = &ModifyThreatIntelligenceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyThreatIntelligenceStatus")
	return
}

func NewModifyThreatIntelligenceStatusResponse() (response *ModifyThreatIntelligenceStatusResponse) {
	response = &ModifyThreatIntelligenceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置威胁情报的开关配置
func (c *Client) ModifyThreatIntelligenceStatus(request *ModifyThreatIntelligenceStatusRequest) (response *ModifyThreatIntelligenceStatusResponse, err error) {
	if request == nil {
		request = NewModifyThreatIntelligenceStatusRequest()
	}
	response = NewModifyThreatIntelligenceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHWDomainStateRequest() (request *QueryHWDomainStateRequest) {
	request = &QueryHWDomainStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryHWDomainState")
	return
}

func NewQueryHWDomainStateResponse() (response *QueryHWDomainStateResponse) {
	response = &QueryHWDomainStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [重保护网]-查询重保护网域名状态接口
func (c *Client) QueryHWDomainState(request *QueryHWDomainStateRequest) (response *QueryHWDomainStateResponse, err error) {
	if request == nil {
		request = NewQueryHWDomainStateRequest()
	}
	response = NewQueryHWDomainStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFirstQualificationsRequest() (request *DescribeFirstQualificationsRequest) {
	request = &DescribeFirstQualificationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFirstQualifications")
	return
}

func NewDescribeFirstQualificationsResponse() (response *DescribeFirstQualificationsResponse) {
	response = &DescribeFirstQualificationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// s d
func (c *Client) DescribeFirstQualifications(request *DescribeFirstQualificationsRequest) (response *DescribeFirstQualificationsResponse, err error) {
	if request == nil {
		request = NewDescribeFirstQualificationsRequest()
	}
	response = NewDescribeFirstQualificationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomRuleRequest() (request *ModifyCustomRuleRequest) {
	request = &ModifyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCustomRule")
	return
}

func NewModifyCustomRuleResponse() (response *ModifyCustomRuleResponse) {
	response = &ModifyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑自定义规则
func (c *Client) ModifyCustomRule(request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
	if request == nil {
		request = NewModifyCustomRuleRequest()
	}
	response = NewModifyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
	request = &DeleteSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteSession")
	return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
	response = &DeleteSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CC攻击的session设置
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
	if request == nil {
		request = NewDeleteSessionRequest()
	}
	response = NewDeleteSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessIndexRequest() (request *DescribeAccessIndexRequest) {
	request = &DescribeAccessIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessIndex")
	return
}

func NewDescribeAccessIndexResponse() (response *DescribeAccessIndexResponse) {
	response = &DescribeAccessIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取访问日志索引配置信息
func (c *Client) DescribeAccessIndex(request *DescribeAccessIndexRequest) (response *DescribeAccessIndexResponse, err error) {
	if request == nil {
		request = NewDescribeAccessIndexRequest()
	}
	response = NewDescribeAccessIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserWafAccessInfoRequest() (request *DescribeUserWafAccessInfoRequest) {
	request = &DescribeUserWafAccessInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserWafAccessInfo")
	return
}

func NewDescribeUserWafAccessInfoResponse() (response *DescribeUserWafAccessInfoResponse) {
	response = &DescribeUserWafAccessInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户waf接入信息，给高防EIP调用
func (c *Client) DescribeUserWafAccessInfo(request *DescribeUserWafAccessInfoRequest) (response *DescribeUserWafAccessInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserWafAccessInfoRequest()
	}
	response = NewDescribeUserWafAccessInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeParametersListRequest() (request *DescribeParametersListRequest) {
	request = &DescribeParametersListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeParametersList")
	return
}

func NewDescribeParametersListResponse() (response *DescribeParametersListResponse) {
	response = &DescribeParametersListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的参数类型列表
func (c *Client) DescribeParametersList(request *DescribeParametersListRequest) (response *DescribeParametersListResponse, err error) {
	if request == nil {
		request = NewDescribeParametersListRequest()
	}
	response = NewDescribeParametersListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIpAccessControlRequest() (request *CreateIpAccessControlRequest) {
	request = &CreateIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateIpAccessControl")
	return
}

func NewCreateIpAccessControlResponse() (response *CreateIpAccessControlResponse) {
	response = &CreateIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单新增接口
func (c *Client) CreateIpAccessControl(request *CreateIpAccessControlRequest) (response *CreateIpAccessControlResponse, err error) {
	if request == nil {
		request = NewCreateIpAccessControlRequest()
	}
	response = NewCreateIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordDetailRequest() (request *DescribeBotRecordDetailRequest) {
	request = &DescribeBotRecordDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRecordDetail")
	return
}

func NewDescribeBotRecordDetailResponse() (response *DescribeBotRecordDetailResponse) {
	response = &DescribeBotRecordDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot详情
func (c *Client) DescribeBotRecordDetail(request *DescribeBotRecordDetailRequest) (response *DescribeBotRecordDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordDetailRequest()
	}
	response = NewDescribeBotRecordDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAggregateTopNRequest() (request *DescribeApiAggregateTopNRequest) {
	request = &DescribeApiAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiAggregateTopN")
	return
}

func NewDescribeApiAggregateTopNResponse() (response *DescribeApiAggregateTopNResponse) {
	response = &DescribeApiAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Api安全模块的访问日志聚合topN
func (c *Client) DescribeApiAggregateTopN(request *DescribeApiAggregateTopNRequest) (response *DescribeApiAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeApiAggregateTopNRequest()
	}
	response = NewDescribeApiAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCheckStatusRequest() (request *DescribeFlowCheckStatusRequest) {
	request = &DescribeFlowCheckStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFlowCheckStatus")
	return
}

func NewDescribeFlowCheckStatusResponse() (response *DescribeFlowCheckStatusResponse) {
	response = &DescribeFlowCheckStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取体检状态接口
func (c *Client) DescribeFlowCheckStatus(request *DescribeFlowCheckStatusRequest) (response *DescribeFlowCheckStatusResponse, err error) {
	if request == nil {
		request = NewDescribeFlowCheckStatusRequest()
	}
	response = NewDescribeFlowCheckStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePanDomainDetailRequest() (request *DescribePanDomainDetailRequest) {
	request = &DescribePanDomainDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePanDomainDetail")
	return
}

func NewDescribePanDomainDetailResponse() (response *DescribePanDomainDetailResponse) {
	response = &DescribePanDomainDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泛域名非标支持详情
func (c *Client) DescribePanDomainDetail(request *DescribePanDomainDetailRequest) (response *DescribePanDomainDetailResponse, err error) {
	if request == nil {
		request = NewDescribePanDomainDetailRequest()
	}
	response = NewDescribePanDomainDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchIpPunishmentStatusRequest() (request *ModifyBatchIpPunishmentStatusRequest) {
	request = &ModifyBatchIpPunishmentStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchIpPunishmentStatus")
	return
}

func NewModifyBatchIpPunishmentStatusResponse() (response *ModifyBatchIpPunishmentStatusResponse) {
	response = &ModifyBatchIpPunishmentStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]批量规则开关接口
func (c *Client) ModifyBatchIpPunishmentStatus(request *ModifyBatchIpPunishmentStatusRequest) (response *ModifyBatchIpPunishmentStatusResponse, err error) {
	if request == nil {
		request = NewModifyBatchIpPunishmentStatusRequest()
	}
	response = NewModifyBatchIpPunishmentStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessFastAnalysisRequest() (request *DescribeAccessFastAnalysisRequest) {
	request = &DescribeAccessFastAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessFastAnalysis")
	return
}

func NewDescribeAccessFastAnalysisResponse() (response *DescribeAccessFastAnalysisResponse) {
	response = &DescribeAccessFastAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于访问日志的快速分析
func (c *Client) DescribeAccessFastAnalysis(request *DescribeAccessFastAnalysisRequest) (response *DescribeAccessFastAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeAccessFastAnalysisRequest()
	}
	response = NewDescribeAccessFastAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorBotStatRequest() (request *DescribeMonitorBotStatRequest) {
	request = &DescribeMonitorBotStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeMonitorBotStat")
	return
}

func NewDescribeMonitorBotStatResponse() (response *DescribeMonitorBotStatResponse) {
	response = &DescribeMonitorBotStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 流量体检-控制台-BOT仅启用监控模式
func (c *Client) DescribeMonitorBotStat(request *DescribeMonitorBotStatRequest) (response *DescribeMonitorBotStatResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorBotStatRequest()
	}
	response = NewDescribeMonitorBotStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneActionRuleRequest() (request *DescribeBotSceneActionRuleRequest) {
	request = &DescribeBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSceneActionRule")
	return
}

func NewDescribeBotSceneActionRuleResponse() (response *DescribeBotSceneActionRuleResponse) {
	response = &DescribeBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT场景的动作策略列表
func (c *Client) DescribeBotSceneActionRule(request *DescribeBotSceneActionRuleRequest) (response *DescribeBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneActionRuleRequest()
	}
	response = NewDescribeBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafRequestCountRequest() (request *DescribeWafRequestCountRequest) {
	request = &DescribeWafRequestCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafRequestCount")
	return
}

func NewDescribeWafRequestCountResponse() (response *DescribeWafRequestCountResponse) {
	response = &DescribeWafRequestCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取waf的访问日志总条数
func (c *Client) DescribeWafRequestCount(request *DescribeWafRequestCountRequest) (response *DescribeWafRequestCountResponse, err error) {
	if request == nil {
		request = NewDescribeWafRequestCountRequest()
	}
	response = NewDescribeWafRequestCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOwaspRuleTypeActionRequest() (request *ModifyOwaspRuleTypeActionRequest) {
	request = &ModifyOwaspRuleTypeActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOwaspRuleTypeAction")
	return
}

func NewModifyOwaspRuleTypeActionResponse() (response *ModifyOwaspRuleTypeActionResponse) {
	response = &ModifyOwaspRuleTypeActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新规则类型的防护模式
func (c *Client) ModifyOwaspRuleTypeAction(request *ModifyOwaspRuleTypeActionRequest) (response *ModifyOwaspRuleTypeActionResponse, err error) {
	if request == nil {
		request = NewModifyOwaspRuleTypeActionRequest()
	}
	response = NewModifyOwaspRuleTypeActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPTypeRequest() (request *DescribeIPTypeRequest) {
	request = &DescribeIPTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeIPType")
	return
}

func NewDescribeIPTypeResponse() (response *DescribeIPTypeResponse) {
	response = &DescribeIPTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加IP黑名单的时候查询IP是否为情报白IP
func (c *Client) DescribeIPType(request *DescribeIPTypeRequest) (response *DescribeIPTypeResponse, err error) {
	if request == nil {
		request = NewDescribeIPTypeRequest()
	}
	response = NewDescribeIPTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
	request = &DescribeTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTopics")
	return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
	response = &DescribeTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志主题列表，支持分页
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicsRequest()
	}
	response = NewDescribeTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackLogCountRequest() (request *DescribeAttackLogCountRequest) {
	request = &DescribeAttackLogCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackLogCount")
	return
}

func NewDescribeAttackLogCountResponse() (response *DescribeAttackLogCountResponse) {
	response = &DescribeAttackLogCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询符合搜索条件攻击日志数量
func (c *Client) DescribeAttackLogCount(request *DescribeAttackLogCountRequest) (response *DescribeAttackLogCountResponse, err error) {
	if request == nil {
		request = NewDescribeAttackLogCountRequest()
	}
	response = NewDescribeAttackLogCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectGroupRequest() (request *DescribeProtectGroupRequest) {
	request = &DescribeProtectGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeProtectGroup")
	return
}

func NewDescribeProtectGroupResponse() (response *DescribeProtectGroupResponse) {
	response = &DescribeProtectGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护对象组详情
func (c *Client) DescribeProtectGroup(request *DescribeProtectGroupRequest) (response *DescribeProtectGroupResponse, err error) {
	if request == nil {
		request = NewDescribeProtectGroupRequest()
	}
	response = NewDescribeProtectGroupResponse()
	err = c.Send(request, response)
	return
}

func NewApplyResponsePageRequest() (request *ApplyResponsePageRequest) {
	request = &ApplyResponsePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ApplyResponsePage")
	return
}

func NewApplyResponsePageResponse() (response *ApplyResponsePageResponse) {
	response = &ApplyResponsePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给域名应用自定义响应页面
func (c *Client) ApplyResponsePage(request *ApplyResponsePageRequest) (response *ApplyResponsePageResponse, err error) {
	if request == nil {
		request = NewApplyResponsePageRequest()
	}
	response = NewApplyResponsePageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotTiRuleRequest() (request *ModifyBotTiRuleRequest) {
	request = &ModifyBotTiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotTiRule")
	return
}

func NewModifyBotTiRuleResponse() (response *ModifyBotTiRuleResponse) {
	response = &ModifyBotTiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-威胁情报配置
func (c *Client) ModifyBotTiRule(request *ModifyBotTiRuleRequest) (response *ModifyBotTiRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotTiRuleRequest()
	}
	response = NewModifyBotTiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDetailRequest() (request *DescribeAssetDetailRequest) {
	request = &DescribeAssetDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAssetDetail")
	return
}

func NewDescribeAssetDetailResponse() (response *DescribeAssetDetailResponse) {
	response = &DescribeAssetDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的资产信息
func (c *Client) DescribeAssetDetail(request *DescribeAssetDetailRequest) (response *DescribeAssetDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDetailRequest()
	}
	response = NewDescribeAssetDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFindDomainListRequest() (request *DescribeFindDomainListRequest) {
	request = &DescribeFindDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFindDomainList")
	return
}

func NewDescribeFindDomainListResponse() (response *DescribeFindDomainListResponse) {
	response = &DescribeFindDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取发现域名列表接口
func (c *Client) DescribeFindDomainList(request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeFindDomainListRequest()
	}
	response = NewDescribeFindDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewGetLocalStorageItemRequest() (request *GetLocalStorageItemRequest) {
	request = &GetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetLocalStorageItem")
	return
}

func NewGetLocalStorageItemResponse() (response *GetLocalStorageItemResponse) {
	response = &GetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程存储缓存
func (c *Client) GetLocalStorageItem(request *GetLocalStorageItemRequest) (response *GetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewGetLocalStorageItemRequest()
	}
	response = NewGetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserClbWafRegionsRequest() (request *DescribeUserClbWafRegionsRequest) {
	request = &DescribeUserClbWafRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserClbWafRegions")
	return
}

func NewDescribeUserClbWafRegionsResponse() (response *DescribeUserClbWafRegionsResponse) {
	response = &DescribeUserClbWafRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在负载均衡型WAF的添加、编辑域名配置的时候，需要展示负载均衡型WAF（clb-waf)支持的地域列表，通过DescribeUserClbWafRegions既可以获得当前对客户已经开放的地域列表
func (c *Client) DescribeUserClbWafRegions(request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeUserClbWafRegionsRequest()
	}
	response = NewDescribeUserClbWafRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewAddBatchCustomWhiteRuleRequest() (request *AddBatchCustomWhiteRuleRequest) {
	request = &AddBatchCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBatchCustomWhiteRule")
	return
}

func NewAddBatchCustomWhiteRuleResponse() (response *AddBatchCustomWhiteRuleResponse) {
	response = &AddBatchCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加批量精准白名单规则
func (c *Client) AddBatchCustomWhiteRule(request *AddBatchCustomWhiteRuleRequest) (response *AddBatchCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddBatchCustomWhiteRuleRequest()
	}
	response = NewAddBatchCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstanceQpsLimitRequest() (request *GetInstanceQpsLimitRequest) {
	request = &GetInstanceQpsLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetInstanceQpsLimit")
	return
}

func NewGetInstanceQpsLimitResponse() (response *GetInstanceQpsLimitResponse) {
	response = &GetInstanceQpsLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取套餐实例的弹性qps上限
func (c *Client) GetInstanceQpsLimit(request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
	if request == nil {
		request = NewGetInstanceQpsLimitRequest()
	}
	response = NewGetInstanceQpsLimitResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchDomainRulesRequest() (request *SwitchDomainRulesRequest) {
	request = &SwitchDomainRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchDomainRules")
	return
}

func NewSwitchDomainRulesResponse() (response *SwitchDomainRulesResponse) {
	response = &SwitchDomainRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换域名的规则开关
func (c *Client) SwitchDomainRules(request *SwitchDomainRulesRequest) (response *SwitchDomainRulesResponse, err error) {
	if request == nil {
		request = NewSwitchDomainRulesRequest()
	}
	response = NewSwitchDomainRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAntiFakeUrlRequest() (request *DeleteAntiFakeUrlRequest) {
	request = &DeleteAntiFakeUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAntiFakeUrl")
	return
}

func NewDeleteAntiFakeUrlResponse() (response *DeleteAntiFakeUrlResponse) {
	response = &DeleteAntiFakeUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防篡改url
func (c *Client) DeleteAntiFakeUrl(request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
	if request == nil {
		request = NewDeleteAntiFakeUrlRequest()
	}
	response = NewDeleteAntiFakeUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityRuleCountRequest() (request *DescribeSecurityRuleCountRequest) {
	request = &DescribeSecurityRuleCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSecurityRuleCount")
	return
}

func NewDescribeSecurityRuleCountResponse() (response *DescribeSecurityRuleCountResponse) {
	response = &DescribeSecurityRuleCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 得到基础安全模块中，每个模块的规则数目
func (c *Client) DescribeSecurityRuleCount(request *DescribeSecurityRuleCountRequest) (response *DescribeSecurityRuleCountResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityRuleCountRequest()
	}
	response = NewDescribeSecurityRuleCountResponse()
	err = c.Send(request, response)
	return
}

func NewEnableUserSignaturePolicyRequest() (request *EnableUserSignaturePolicyRequest) {
	request = &EnableUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EnableUserSignaturePolicy")
	return
}

func NewEnableUserSignaturePolicyResponse() (response *EnableUserSignaturePolicyResponse) {
	response = &EnableUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启用户防护规则开关
func (c *Client) EnableUserSignaturePolicy(request *EnableUserSignaturePolicyRequest) (response *EnableUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewEnableUserSignaturePolicyRequest()
	}
	response = NewEnableUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPackageRenewRequest() (request *ModifyPackageRenewRequest) {
	request = &ModifyPackageRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyPackageRenew")
	return
}

func NewModifyPackageRenewResponse() (response *ModifyPackageRenewResponse) {
	response = &ModifyPackageRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置套餐自动续费
func (c *Client) ModifyPackageRenew(request *ModifyPackageRenewRequest) (response *ModifyPackageRenewResponse, err error) {
	if request == nil {
		request = NewModifyPackageRenewRequest()
	}
	response = NewModifyPackageRenewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCdcResourceRequest() (request *DeleteCdcResourceRequest) {
	request = &DeleteCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCdcResource")
	return
}

func NewDeleteCdcResourceResponse() (response *DeleteCdcResourceResponse) {
	response = &DeleteCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cdc资源数据
func (c *Client) DeleteCdcResource(request *DeleteCdcResourceRequest) (response *DeleteCdcResourceResponse, err error) {
	if request == nil {
		request = NewDeleteCdcResourceRequest()
	}
	response = NewDeleteCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserInstancesRequest() (request *DescribeUserInstancesRequest) {
	request = &DescribeUserInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserInstances")
	return
}

func NewDescribeUserInstancesResponse() (response *DescribeUserInstancesResponse) {
	response = &DescribeUserInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户实例展示，展示通过地域参数展示特定地域用户购买的所有实例
func (c *Client) DescribeUserInstances(request *DescribeUserInstancesRequest) (response *DescribeUserInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeUserInstancesRequest()
	}
	response = NewDescribeUserInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewExportAccessLogsRequest() (request *ExportAccessLogsRequest) {
	request = &ExportAccessLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ExportAccessLogs")
	return
}

func NewExportAccessLogsResponse() (response *ExportAccessLogsResponse) {
	response = &ExportAccessLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问日志批量下载接口，不能带搜索条件
func (c *Client) ExportAccessLogs(request *ExportAccessLogsRequest) (response *ExportAccessLogsResponse, err error) {
	if request == nil {
		request = NewExportAccessLogsRequest()
	}
	response = NewExportAccessLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefaultProtectConfigRequest() (request *DescribeDefaultProtectConfigRequest) {
	request = &DescribeDefaultProtectConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDefaultProtectConfig")
	return
}

func NewDescribeDefaultProtectConfigResponse() (response *DescribeDefaultProtectConfigResponse) {
	response = &DescribeDefaultProtectConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户缺省规则引擎类型
func (c *Client) DescribeDefaultProtectConfig(request *DescribeDefaultProtectConfigRequest) (response *DescribeDefaultProtectConfigResponse, err error) {
	if request == nil {
		request = NewDescribeDefaultProtectConfigRequest()
	}
	response = NewDescribeDefaultProtectConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchAreaBanRequest() (request *ModifyBatchAreaBanRequest) {
	request = &ModifyBatchAreaBanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchAreaBan")
	return
}

func NewModifyBatchAreaBanResponse() (response *ModifyBatchAreaBanResponse) {
	response = &ModifyBatchAreaBanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [地域封禁]-批量编辑接口
func (c *Client) ModifyBatchAreaBan(request *ModifyBatchAreaBanRequest) (response *ModifyBatchAreaBanResponse, err error) {
	if request == nil {
		request = NewModifyBatchAreaBanRequest()
	}
	response = NewModifyBatchAreaBanResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDNSDetectDomainRequest() (request *DeleteDNSDetectDomainRequest) {
	request = &DeleteDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteDNSDetectDomain")
	return
}

func NewDeleteDNSDetectDomainResponse() (response *DeleteDNSDetectDomainResponse) {
	response = &DeleteDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【删除】这个功能。
func (c *Client) DeleteDNSDetectDomain(request *DeleteDNSDetectDomainRequest) (response *DeleteDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewDeleteDNSDetectDomainRequest()
	}
	response = NewDeleteDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAntiInfoLeakRulesRequest() (request *DescribeAntiInfoLeakRulesRequest) {
	request = &DescribeAntiInfoLeakRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAntiInfoLeakRules")
	return
}

func NewDescribeAntiInfoLeakRulesResponse() (response *DescribeAntiInfoLeakRulesResponse) {
	response = &DescribeAntiInfoLeakRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取信息防泄漏规则列表
func (c *Client) DescribeAntiInfoLeakRules(request *DescribeAntiInfoLeakRulesRequest) (response *DescribeAntiInfoLeakRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAntiInfoLeakRulesRequest()
	}
	response = NewDescribeAntiInfoLeakRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMLimitRuleStatusRequest() (request *ModifyLLMLimitRuleStatusRequest) {
	request = &ModifyLLMLimitRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMLimitRuleStatus")
	return
}

func NewModifyLLMLimitRuleStatusResponse() (response *ModifyLLMLimitRuleStatusResponse) {
	response = &ModifyLLMLimitRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换大模型算力消耗防护规则开关
func (c *Client) ModifyLLMLimitRuleStatus(request *ModifyLLMLimitRuleStatusRequest) (response *ModifyLLMLimitRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyLLMLimitRuleStatusRequest()
	}
	response = NewModifyLLMLimitRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSessionRequest() (request *DescribeSessionRequest) {
	request = &DescribeSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSession")
	return
}

func NewDescribeSessionResponse() (response *DescribeSessionResponse) {
	response = &DescribeSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 会话定义查询接口
func (c *Client) DescribeSession(request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
	if request == nil {
		request = NewDescribeSessionRequest()
	}
	response = NewDescribeSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCertificatedDomainRequest() (request *DescribeCertificatedDomainRequest) {
	request = &DescribeCertificatedDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCertificatedDomain")
	return
}

func NewDescribeCertificatedDomainResponse() (response *DescribeCertificatedDomainResponse) {
	response = &DescribeCertificatedDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过ssl证书id获取可以绑定的waf中的域名
func (c *Client) DescribeCertificatedDomain(request *DescribeCertificatedDomainRequest) (response *DescribeCertificatedDomainResponse, err error) {
	if request == nil {
		request = NewDescribeCertificatedDomainRequest()
	}
	response = NewDescribeCertificatedDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAnalyzeStatusRequest() (request *DescribeApiAnalyzeStatusRequest) {
	request = &DescribeApiAnalyzeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiAnalyzeStatus")
	return
}

func NewDescribeApiAnalyzeStatusResponse() (response *DescribeApiAnalyzeStatusResponse) {
	response = &DescribeApiAnalyzeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看api资产分析开关
func (c *Client) DescribeApiAnalyzeStatus(request *DescribeApiAnalyzeStatusRequest) (response *DescribeApiAnalyzeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiAnalyzeStatusRequest()
	}
	response = NewDescribeApiAnalyzeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorInstancesRequest() (request *DescribeMonitorInstancesRequest) {
	request = &DescribeMonitorInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeMonitorInstances")
	return
}

func NewDescribeMonitorInstancesResponse() (response *DescribeMonitorInstancesResponse) {
	response = &DescribeMonitorInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云监控的实例
func (c *Client) DescribeMonitorInstances(request *DescribeMonitorInstancesRequest) (response *DescribeMonitorInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorInstancesRequest()
	}
	response = NewDescribeMonitorInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceAttackLogPostRequest() (request *ModifyInstanceAttackLogPostRequest) {
	request = &ModifyInstanceAttackLogPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceAttackLogPost")
	return
}

func NewModifyInstanceAttackLogPostResponse() (response *ModifyInstanceAttackLogPostResponse) {
	response = &ModifyInstanceAttackLogPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例攻击日志投递开关，企业版及以上版本可以开通，否则返回错误
func (c *Client) ModifyInstanceAttackLogPost(request *ModifyInstanceAttackLogPostRequest) (response *ModifyInstanceAttackLogPostResponse, err error) {
	if request == nil {
		request = NewModifyInstanceAttackLogPostRequest()
	}
	response = NewModifyInstanceAttackLogPostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecSensitiveRuleListRequest() (request *DescribeApiSecSensitiveRuleListRequest) {
	request = &DescribeApiSecSensitiveRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecSensitiveRuleList")
	return
}

func NewDescribeApiSecSensitiveRuleListResponse() (response *DescribeApiSecSensitiveRuleListResponse) {
	response = &DescribeApiSecSensitiveRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全敏感规则列表
func (c *Client) DescribeApiSecSensitiveRuleList(request *DescribeApiSecSensitiveRuleListRequest) (response *DescribeApiSecSensitiveRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecSensitiveRuleListRequest()
	}
	response = NewDescribeApiSecSensitiveRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOpRuleUpdateLogRequest() (request *DeleteOpRuleUpdateLogRequest) {
	request = &DeleteOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteOpRuleUpdateLog")
	return
}

func NewDeleteOpRuleUpdateLogResponse() (response *DeleteOpRuleUpdateLogResponse) {
	response = &DeleteOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：删除规则更新日志
func (c *Client) DeleteOpRuleUpdateLog(request *DeleteOpRuleUpdateLogRequest) (response *DeleteOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDeleteOpRuleUpdateLogRequest()
	}
	response = NewDeleteOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainIpv6StatusRequest() (request *ModifyDomainIpv6StatusRequest) {
	request = &ModifyDomainIpv6StatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainIpv6Status")
	return
}

func NewModifyDomainIpv6StatusResponse() (response *ModifyDomainIpv6StatusResponse) {
	response = &ModifyDomainIpv6StatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换ipv6开关
func (c *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
	if request == nil {
		request = NewModifyDomainIpv6StatusRequest()
	}
	response = NewModifyDomainIpv6StatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostInnerRequest() (request *ModifyHostInnerRequest) {
	request = &ModifyHostInnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostInner")
	return
}

func NewModifyHostInnerResponse() (response *ModifyHostInnerResponse) {
	response = &ModifyHostInnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf编辑防护域名配置内网
func (c *Client) ModifyHostInner(request *ModifyHostInnerRequest) (response *ModifyHostInnerResponse, err error) {
	if request == nil {
		request = NewModifyHostInnerRequest()
	}
	response = NewModifyHostInnerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTceAlarmRequest() (request *ModifyTceAlarmRequest) {
	request = &ModifyTceAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyTceAlarm")
	return
}

func NewModifyTceAlarmResponse() (response *ModifyTceAlarmResponse) {
	response = &ModifyTceAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警配置
func (c *Client) ModifyTceAlarm(request *ModifyTceAlarmRequest) (response *ModifyTceAlarmResponse, err error) {
	if request == nil {
		request = NewModifyTceAlarmRequest()
	}
	response = NewModifyTceAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrialQRCodeV3Request() (request *DescribeTrialQRCodeV3Request) {
	request = &DescribeTrialQRCodeV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTrialQRCodeV3")
	return
}

func NewDescribeTrialQRCodeV3Response() (response *DescribeTrialQRCodeV3Response) {
	response = &DescribeTrialQRCodeV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用本接口可以获取小程序的体验版二维码
func (c *Client) DescribeTrialQRCodeV3(request *DescribeTrialQRCodeV3Request) (response *DescribeTrialQRCodeV3Response, err error) {
	if request == nil {
		request = NewDescribeTrialQRCodeV3Request()
	}
	response = NewDescribeTrialQRCodeV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBusinessRiskRequest() (request *DescribeBusinessRiskRequest) {
	request = &DescribeBusinessRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBusinessRisk")
	return
}

func NewDescribeBusinessRiskResponse() (response *DescribeBusinessRiskResponse) {
	response = &DescribeBusinessRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务安全规则查询
func (c *Client) DescribeBusinessRisk(request *DescribeBusinessRiskRequest) (response *DescribeBusinessRiskResponse, err error) {
	if request == nil {
		request = NewDescribeBusinessRiskRequest()
	}
	response = NewDescribeBusinessRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartaProtectionListRequest() (request *DescribeSpartaProtectionListRequest) {
	request = &DescribeSpartaProtectionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSpartaProtectionList")
	return
}

func NewDescribeSpartaProtectionListResponse() (response *DescribeSpartaProtectionListResponse) {
	response = &DescribeSpartaProtectionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询域名列表
func (c *Client) DescribeSpartaProtectionList(request *DescribeSpartaProtectionListRequest) (response *DescribeSpartaProtectionListResponse, err error) {
	if request == nil {
		request = NewDescribeSpartaProtectionListRequest()
	}
	response = NewDescribeSpartaProtectionListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDNSDetectDomainRequest() (request *ModifyDNSDetectDomainRequest) {
	request = &ModifyDNSDetectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDNSDetectDomain")
	return
}

func NewModifyDNSDetectDomainResponse() (response *ModifyDNSDetectDomainResponse) {
	response = &ModifyDNSDetectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑DNS劫持检测的域名，对应【Web安全检测】【DNS劫持检测】【基础设置】【编辑】这个功能。
func (c *Client) ModifyDNSDetectDomain(request *ModifyDNSDetectDomainRequest) (response *ModifyDNSDetectDomainResponse, err error) {
	if request == nil {
		request = NewModifyDNSDetectDomainRequest()
	}
	response = NewModifyDNSDetectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogStatusRequest() (request *EditOpRuleUpdateLogStatusRequest) {
	request = &EditOpRuleUpdateLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EditOpRuleUpdateLogStatus")
	return
}

func NewEditOpRuleUpdateLogStatusResponse() (response *EditOpRuleUpdateLogStatusResponse) {
	response = &EditOpRuleUpdateLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志的发布状态
func (c *Client) EditOpRuleUpdateLogStatus(request *EditOpRuleUpdateLogStatusRequest) (response *EditOpRuleUpdateLogStatusResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogStatusRequest()
	}
	response = NewEditOpRuleUpdateLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceWafInstanceRequest() (request *InquiryPriceWafInstanceRequest) {
	request = &InquiryPriceWafInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "InquiryPriceWafInstance")
	return
}

func NewInquiryPriceWafInstanceResponse() (response *InquiryPriceWafInstanceResponse) {
	response = &InquiryPriceWafInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询价格
func (c *Client) InquiryPriceWafInstance(request *InquiryPriceWafInstanceRequest) (response *InquiryPriceWafInstanceResponse, err error) {
	if request == nil {
		request = NewInquiryPriceWafInstanceRequest()
	}
	response = NewInquiryPriceWafInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMContentRuleRequest() (request *ModifyLLMContentRuleRequest) {
	request = &ModifyLLMContentRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMContentRule")
	return
}

func NewModifyLLMContentRuleResponse() (response *ModifyLLMContentRuleResponse) {
	response = &ModifyLLMContentRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新大模型内容安全规则
func (c *Client) ModifyLLMContentRule(request *ModifyLLMContentRuleRequest) (response *ModifyLLMContentRuleResponse, err error) {
	if request == nil {
		request = NewModifyLLMContentRuleRequest()
	}
	response = NewModifyLLMContentRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppletTaskInfoRequest() (request *DescribeAppletTaskInfoRequest) {
	request = &DescribeAppletTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAppletTaskInfo")
	return
}

func NewDescribeAppletTaskInfoResponse() (response *DescribeAppletTaskInfoResponse) {
	response = &DescribeAppletTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取小程序压测任务预约信息(运营平台)
func (c *Client) DescribeAppletTaskInfo(request *DescribeAppletTaskInfoRequest) (response *DescribeAppletTaskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAppletTaskInfoRequest()
	}
	response = NewDescribeAppletTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
	request = &DescribeProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeProductInfo")
	return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
	response = &DescribeProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户产品信息
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProductInfoRequest()
	}
	response = NewDescribeProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOwaspWhiteRuleRequest() (request *CreateOwaspWhiteRuleRequest) {
	request = &CreateOwaspWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateOwaspWhiteRule")
	return
}

func NewCreateOwaspWhiteRuleResponse() (response *CreateOwaspWhiteRuleResponse) {
	response = &CreateOwaspWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加规则引擎白名单
func (c *Client) CreateOwaspWhiteRule(request *CreateOwaspWhiteRuleRequest) (response *CreateOwaspWhiteRuleResponse, err error) {
	if request == nil {
		request = NewCreateOwaspWhiteRuleRequest()
	}
	response = NewCreateOwaspWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAttackOverviewRequest() (request *DescribeBotAttackOverviewRequest) {
	request = &DescribeBotAttackOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAttackOverview")
	return
}

func NewDescribeBotAttackOverviewResponse() (response *DescribeBotAttackOverviewResponse) {
	response = &DescribeBotAttackOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot攻击报表概览
func (c *Client) DescribeBotAttackOverview(request *DescribeBotAttackOverviewRequest) (response *DescribeBotAttackOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotAttackOverviewRequest()
	}
	response = NewDescribeBotAttackOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewAddBatchCustomRuleRequest() (request *AddBatchCustomRuleRequest) {
	request = &AddBatchCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddBatchCustomRule")
	return
}

func NewAddBatchCustomRuleResponse() (response *AddBatchCustomRuleResponse) {
	response = &AddBatchCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-批量新增接口
func (c *Client) AddBatchCustomRule(request *AddBatchCustomRuleRequest) (response *AddBatchCustomRuleResponse, err error) {
	if request == nil {
		request = NewAddBatchCustomRuleRequest()
	}
	response = NewAddBatchCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLLMSessionRequest() (request *ModifyLLMSessionRequest) {
	request = &ModifyLLMSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyLLMSession")
	return
}

func NewModifyLLMSessionResponse() (response *ModifyLLMSessionResponse) {
	response = &ModifyLLMSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改session会话配置
func (c *Client) ModifyLLMSession(request *ModifyLLMSessionRequest) (response *ModifyLLMSessionResponse, err error) {
	if request == nil {
		request = NewModifyLLMSessionRequest()
	}
	response = NewModifyLLMSessionResponse()
	err = c.Send(request, response)
	return
}

func NewNoticeInstanceChangeRequest() (request *NoticeInstanceChangeRequest) {
	request = &NoticeInstanceChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "NoticeInstanceChange")
	return
}

func NewNoticeInstanceChangeResponse() (response *NoticeInstanceChangeResponse) {
	response = &NoticeInstanceChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对新购实例进行带宽限速，客户在购买、降配业务扩展包/开启关闭弹性计费/弹性计费上限调整等操作后通知接入
func (c *Client) NoticeInstanceChange(request *NoticeInstanceChangeRequest) (response *NoticeInstanceChangeResponse, err error) {
	if request == nil {
		request = NewNoticeInstanceChangeRequest()
	}
	response = NewNoticeInstanceChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountAttributesRequest() (request *DescribeAccountAttributesRequest) {
	request = &DescribeAccountAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccountAttributes")
	return
}

func NewDescribeAccountAttributesResponse() (response *DescribeAccountAttributesResponse) {
	response = &DescribeAccountAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询登陆账号的属性
func (c *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeAccountAttributesRequest()
	}
	response = NewDescribeAccountAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
	request = &DescribeDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomains")
	return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
	response = &DescribeDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有域名的详细信息
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainsRequest()
	}
	response = NewDescribeDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpartUserInfoRequest() (request *DescribeSpartUserInfoRequest) {
	request = &DescribeSpartUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSpartUserInfo")
	return
}

func NewDescribeSpartUserInfoResponse() (response *DescribeSpartUserInfoResponse) {
	response = &DescribeSpartUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息，判断用户是否开通waf，套餐购买情况等信息
func (c *Client) DescribeSpartUserInfo(request *DescribeSpartUserInfoRequest) (response *DescribeSpartUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSpartUserInfoRequest()
	}
	response = NewDescribeSpartUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPublishOrCancelRuleRequest() (request *ModifyPublishOrCancelRuleRequest) {
	request = &ModifyPublishOrCancelRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyPublishOrCancelRule")
	return
}

func NewModifyPublishOrCancelRuleResponse() (response *ModifyPublishOrCancelRuleResponse) {
	response = &ModifyPublishOrCancelRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务安全规则上线or下线
func (c *Client) ModifyPublishOrCancelRule(request *ModifyPublishOrCancelRuleRequest) (response *ModifyPublishOrCancelRuleResponse, err error) {
	if request == nil {
		request = NewModifyPublishOrCancelRuleRequest()
	}
	response = NewModifyPublishOrCancelRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainLabelsRequest() (request *ModifyDomainLabelsRequest) {
	request = &ModifyDomainLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainLabels")
	return
}

func NewModifyDomainLabelsResponse() (response *ModifyDomainLabelsResponse) {
	response = &ModifyDomainLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑域名标签
func (c *Client) ModifyDomainLabels(request *ModifyDomainLabelsRequest) (response *ModifyDomainLabelsResponse, err error) {
	if request == nil {
		request = NewModifyDomainLabelsRequest()
	}
	response = NewModifyDomainLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCdcResourceRequest() (request *DescribeCdcResourceRequest) {
	request = &DescribeCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCdcResource")
	return
}

func NewDescribeCdcResourceResponse() (response *DescribeCdcResourceResponse) {
	response = &DescribeCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供运营后台获取Cdc资源列表
func (c *Client) DescribeCdcResource(request *DescribeCdcResourceRequest) (response *DescribeCdcResourceResponse, err error) {
	if request == nil {
		request = NewDescribeCdcResourceRequest()
	}
	response = NewDescribeCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotGenerateDealsRequest() (request *ModifyBotGenerateDealsRequest) {
	request = &ModifyBotGenerateDealsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotGenerateDeals")
	return
}

func NewModifyBotGenerateDealsResponse() (response *ModifyBotGenerateDealsResponse) {
	response = &ModifyBotGenerateDealsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给使用的waf实例 bot下单接口
func (c *Client) ModifyBotGenerateDeals(request *ModifyBotGenerateDealsRequest) (response *ModifyBotGenerateDealsResponse, err error) {
	if request == nil {
		request = NewModifyBotGenerateDealsRequest()
	}
	response = NewModifyBotGenerateDealsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiSecSensitiveRuleRequest() (request *ModifyApiSecSensitiveRuleRequest) {
	request = &ModifyApiSecSensitiveRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiSecSensitiveRule")
	return
}

func NewModifyApiSecSensitiveRuleResponse() (response *ModifyApiSecSensitiveRuleResponse) {
	response = &ModifyApiSecSensitiveRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改api安全敏感检测规则
func (c *Client) ModifyApiSecSensitiveRule(request *ModifyApiSecSensitiveRuleRequest) (response *ModifyApiSecSensitiveRuleResponse, err error) {
	if request == nil {
		request = NewModifyApiSecSensitiveRuleRequest()
	}
	response = NewModifyApiSecSensitiveRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsResourceRequest() (request *DescribeTsResourceRequest) {
	request = &DescribeTsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTsResource")
	return
}

func NewDescribeTsResourceResponse() (response *DescribeTsResourceResponse) {
	response = &DescribeTsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供运营后台获取自研引擎服务资源列表
func (c *Client) DescribeTsResource(request *DescribeTsResourceRequest) (response *DescribeTsResourceResponse, err error) {
	if request == nil {
		request = NewDescribeTsResourceRequest()
	}
	response = NewDescribeTsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyPostCKafkaFlowRequest() (request *DestroyPostCKafkaFlowRequest) {
	request = &DestroyPostCKafkaFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DestroyPostCKafkaFlow")
	return
}

func NewDestroyPostCKafkaFlowResponse() (response *DestroyPostCKafkaFlowResponse) {
	response = &DestroyPostCKafkaFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁CKafka投递流任务
func (c *Client) DestroyPostCKafkaFlow(request *DestroyPostCKafkaFlowRequest) (response *DestroyPostCKafkaFlowResponse, err error) {
	if request == nil {
		request = NewDestroyPostCKafkaFlowRequest()
	}
	response = NewDestroyPostCKafkaFlowResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportRequest() (request *CreateExportRequest) {
	request = &CreateExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateExport")
	return
}

func NewCreateExportResponse() (response *CreateExportResponse) {
	response = &CreateExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口仅创建下载任务，任务返回的下载地址，请用户调用DescribeExports查看任务列表。其中有下载地址CosPath参数。参考文档https://cloud.tencent.com/document/product/614/56449
func (c *Client) CreateExport(request *CreateExportRequest) (response *CreateExportResponse, err error) {
	if request == nil {
		request = NewCreateExportRequest()
	}
	response = NewCreateExportResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttackWhiteRuleRequest() (request *DeleteAttackWhiteRuleRequest) {
	request = &DeleteAttackWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAttackWhiteRule")
	return
}

func NewDeleteAttackWhiteRuleResponse() (response *DeleteAttackWhiteRuleResponse) {
	response = &DeleteAttackWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供用户控制台调用，删除Tiga规则引擎白名单。
func (c *Client) DeleteAttackWhiteRule(request *DeleteAttackWhiteRuleRequest) (response *DeleteAttackWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAttackWhiteRuleRequest()
	}
	response = NewDeleteAttackWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDataMapRequest() (request *DescribeDNSDetectDataMapRequest) {
	request = &DescribeDNSDetectDataMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDNSDetectDataMap")
	return
}

func NewDescribeDNSDetectDataMapResponse() (response *DescribeDNSDetectDataMapResponse) {
	response = &DescribeDNSDetectDataMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DNS劫持-获取地图数据
func (c *Client) DescribeDNSDetectDataMap(request *DescribeDNSDetectDataMapRequest) (response *DescribeDNSDetectDataMapResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDataMapRequest()
	}
	response = NewDescribeDNSDetectDataMapResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAntiFakeUrlStatusRequest() (request *ModifyAntiFakeUrlStatusRequest) {
	request = &ModifyAntiFakeUrlStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAntiFakeUrlStatus")
	return
}

func NewModifyAntiFakeUrlStatusResponse() (response *ModifyAntiFakeUrlStatusResponse) {
	response = &ModifyAntiFakeUrlStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换防篡改开关
func (c *Client) ModifyAntiFakeUrlStatus(request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
	if request == nil {
		request = NewModifyAntiFakeUrlStatusRequest()
	}
	response = NewModifyAntiFakeUrlStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneUCBRuleRequest() (request *DescribeBotSceneUCBRuleRequest) {
	request = &DescribeBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSceneUCBRule")
	return
}

func NewDescribeBotSceneUCBRuleResponse() (response *DescribeBotSceneUCBRuleResponse) {
	response = &DescribeBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 场景化后Bot获取UCB自定义规则策略
func (c *Client) DescribeBotSceneUCBRule(request *DescribeBotSceneUCBRuleRequest) (response *DescribeBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneUCBRuleRequest()
	}
	response = NewDescribeBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccessDownloadRecordsRequest() (request *GetAccessDownloadRecordsRequest) {
	request = &GetAccessDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAccessDownloadRecords")
	return
}

func NewGetAccessDownloadRecordsResponse() (response *GetAccessDownloadRecordsResponse) {
	response = &GetAccessDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载访问日志任务记录列表
func (c *Client) GetAccessDownloadRecords(request *GetAccessDownloadRecordsRequest) (response *GetAccessDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewGetAccessDownloadRecordsRequest()
	}
	response = NewGetAccessDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescirbeCCRuleRequest() (request *DescirbeCCRuleRequest) {
	request = &DescirbeCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescirbeCCRule")
	return
}

func NewDescirbeCCRuleResponse() (response *DescirbeCCRuleResponse) {
	response = &DescirbeCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Query接口
func (c *Client) DescirbeCCRule(request *DescirbeCCRuleRequest) (response *DescirbeCCRuleResponse, err error) {
	if request == nil {
		request = NewDescirbeCCRuleRequest()
	}
	response = NewDescirbeCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePieChartRequest() (request *DescribePieChartRequest) {
	request = &DescribePieChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePieChart")
	return
}

func NewDescribePieChartResponse() (response *DescribePieChartResponse) {
	response = &DescribePieChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定聚类条件的饼图
func (c *Client) DescribePieChart(request *DescribePieChartRequest) (response *DescribePieChartResponse, err error) {
	if request == nil {
		request = NewDescribePieChartRequest()
	}
	response = NewDescribePieChartResponse()
	err = c.Send(request, response)
	return
}

func NewAccessSpartaProtectionRequest() (request *AccessSpartaProtectionRequest) {
	request = &AccessSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AccessSpartaProtection")
	return
}

func NewAccessSpartaProtectionResponse() (response *AccessSpartaProtectionResponse) {
	response = &AccessSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键关闭bypass
func (c *Client) AccessSpartaProtection(request *AccessSpartaProtectionRequest) (response *AccessSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewAccessSpartaProtectionRequest()
	}
	response = NewAccessSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiOverviewRequest() (request *DescribeApiOverviewRequest) {
	request = &DescribeApiOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiOverview")
	return
}

func NewDescribeApiOverviewResponse() (response *DescribeApiOverviewResponse) {
	response = &DescribeApiOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api资产概览数据
func (c *Client) DescribeApiOverview(request *DescribeApiOverviewRequest) (response *DescribeApiOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeApiOverviewRequest()
	}
	response = NewDescribeApiOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomWhiteRuleRequest() (request *DescribeCustomWhiteRuleRequest) {
	request = &DescribeCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomWhiteRule")
	return
}

func NewDescribeCustomWhiteRuleResponse() (response *DescribeCustomWhiteRuleResponse) {
	response = &DescribeCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防护配置中的精准白名单策略列表
func (c *Client) DescribeCustomWhiteRule(request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCustomWhiteRuleRequest()
	}
	response = NewDescribeCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostRequest() (request *DescribeHostRequest) {
	request = &DescribeHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHost")
	return
}

func NewDescribeHostResponse() (response *DescribeHostResponse) {
	response = &DescribeHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf获取防护域名详情
func (c *Client) DescribeHost(request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
	if request == nil {
		request = NewDescribeHostRequest()
	}
	response = NewDescribeHostResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDnsRuleRequest() (request *ModifyDnsRuleRequest) {
	request = &ModifyDnsRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDnsRule")
	return
}

func NewModifyDnsRuleResponse() (response *ModifyDnsRuleResponse) {
	response = &ModifyDnsRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将接入WAF的域名切回源站/或者切回WAF
func (c *Client) ModifyDnsRule(request *ModifyDnsRuleRequest) (response *ModifyDnsRuleResponse, err error) {
	if request == nil {
		request = NewModifyDnsRuleRequest()
	}
	response = NewModifyDnsRuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackPointsRequest() (request *GetAttackPointsRequest) {
	request = &GetAttackPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackPoints")
	return
}

func NewGetAttackPointsResponse() (response *GetAttackPointsResponse) {
	response = &GetAttackPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意量趋势图
func (c *Client) GetAttackPoints(request *GetAttackPointsRequest) (response *GetAttackPointsResponse, err error) {
	if request == nil {
		request = NewGetAttackPointsRequest()
	}
	response = NewGetAttackPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchLogTemplateRequest() (request *DescribeSearchLogTemplateRequest) {
	request = &DescribeSearchLogTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSearchLogTemplate")
	return
}

func NewDescribeSearchLogTemplateResponse() (response *DescribeSearchLogTemplateResponse) {
	response = &DescribeSearchLogTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志搜索模板
func (c *Client) DescribeSearchLogTemplate(request *DescribeSearchLogTemplateRequest) (response *DescribeSearchLogTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeSearchLogTemplateRequest()
	}
	response = NewDescribeSearchLogTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCCRuleRequest() (request *DeleteCCRuleRequest) {
	request = &DeleteCCRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCCRule")
	return
}

func NewDeleteCCRuleResponse() (response *DeleteCCRuleResponse) {
	response = &DeleteCCRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  CC V2 Delete接口
func (c *Client) DeleteCCRule(request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCCRuleRequest()
	}
	response = NewDeleteCCRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserSignatureRuleRequest() (request *DescribeUserSignatureRuleRequest) {
	request = &DescribeUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserSignatureRule")
	return
}

func NewDescribeUserSignatureRuleResponse() (response *DescribeUserSignatureRuleResponse) {
	response = &DescribeUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户特征规则列表
func (c *Client) DescribeUserSignatureRule(request *DescribeUserSignatureRuleRequest) (response *DescribeUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeUserSignatureRuleRequest()
	}
	response = NewDescribeUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpartaProtectionAutoRequest() (request *AddSpartaProtectionAutoRequest) {
	request = &AddSpartaProtectionAutoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddSpartaProtectionAuto")
	return
}

func NewAddSpartaProtectionAutoResponse() (response *AddSpartaProtectionAutoResponse) {
	response = &AddSpartaProtectionAutoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键接入
func (c *Client) AddSpartaProtectionAuto(request *AddSpartaProtectionAutoRequest) (response *AddSpartaProtectionAutoResponse, err error) {
	if request == nil {
		request = NewAddSpartaProtectionAutoRequest()
	}
	response = NewAddSpartaProtectionAutoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceDomainListRequest() (request *DescribeInstanceDomainListRequest) {
	request = &DescribeInstanceDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInstanceDomainList")
	return
}

func NewDescribeInstanceDomainListResponse() (response *DescribeInstanceDomainListResponse) {
	response = &DescribeInstanceDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [重保护网]-获取实例Id域名列表接口
func (c *Client) DescribeInstanceDomainList(request *DescribeInstanceDomainListRequest) (response *DescribeInstanceDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceDomainListRequest()
	}
	response = NewDescribeInstanceDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceElasticModeRequest() (request *ModifyInstanceElasticModeRequest) {
	request = &ModifyInstanceElasticModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceElasticMode")
	return
}

func NewModifyInstanceElasticModeResponse() (response *ModifyInstanceElasticModeResponse) {
	response = &ModifyInstanceElasticModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的QPS弹性计费开关
func (c *Client) ModifyInstanceElasticMode(request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceElasticModeRequest()
	}
	response = NewModifyInstanceElasticModeResponse()
	err = c.Send(request, response)
	return
}

func NewRuleDeleteRequest() (request *RuleDeleteRequest) {
	request = &RuleDeleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RuleDelete")
	return
}

func NewRuleDeleteResponse() (response *RuleDeleteResponse) {
	response = &RuleDeleteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除tiga规则库
func (c *Client) RuleDelete(request *RuleDeleteRequest) (response *RuleDeleteResponse, err error) {
	if request == nil {
		request = NewRuleDeleteRequest()
	}
	response = NewRuleDeleteResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResponsePageRequest() (request *DeleteResponsePageRequest) {
	request = &DeleteResponsePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteResponsePage")
	return
}

func NewDeleteResponsePageResponse() (response *DeleteResponsePageResponse) {
	response = &DeleteResponsePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义响应页面，所有应用了该响应页面的域名都会自动删除该响应页面。
func (c *Client) DeleteResponsePage(request *DeleteResponsePageRequest) (response *DeleteResponsePageResponse, err error) {
	if request == nil {
		request = NewDeleteResponsePageRequest()
	}
	response = NewDeleteResponsePageResponse()
	err = c.Send(request, response)
	return
}

func NewHavePostFlowServiceRequest() (request *HavePostFlowServiceRequest) {
	request = &HavePostFlowServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "HavePostFlowService")
	return
}

func NewHavePostFlowServiceResponse() (response *HavePostFlowServiceResponse) {
	response = &HavePostFlowServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 是否开通投递服务接口
func (c *Client) HavePostFlowService(request *HavePostFlowServiceRequest) (response *HavePostFlowServiceResponse, err error) {
	if request == nil {
		request = NewHavePostFlowServiceRequest()
	}
	response = NewHavePostFlowServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceQpsLimitRequest() (request *ModifyInstanceQpsLimitRequest) {
	request = &ModifyInstanceQpsLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceQpsLimit")
	return
}

func NewModifyInstanceQpsLimitResponse() (response *ModifyInstanceQpsLimitResponse) {
	response = &ModifyInstanceQpsLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置套餐实例的弹性qps上限
func (c *Client) ModifyInstanceQpsLimit(request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
	if request == nil {
		request = NewModifyInstanceQpsLimitRequest()
	}
	response = NewModifyInstanceQpsLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessHistogramRequest() (request *DescribeAccessHistogramRequest) {
	request = &DescribeAccessHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessHistogram")
	return
}

func NewDescribeAccessHistogramResponse() (response *DescribeAccessHistogramResponse) {
	response = &DescribeAccessHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于访问日志柱状趋势图
func (c *Client) DescribeAccessHistogram(request *DescribeAccessHistogramRequest) (response *DescribeAccessHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeAccessHistogramRequest()
	}
	response = NewDescribeAccessHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecAttackSourceListRequest() (request *DescribeApiSecAttackSourceListRequest) {
	request = &DescribeApiSecAttackSourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecAttackSourceList")
	return
}

func NewDescribeApiSecAttackSourceListResponse() (response *DescribeApiSecAttackSourceListResponse) {
	response = &DescribeApiSecAttackSourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全事件攻击源详情
func (c *Client) DescribeApiSecAttackSourceList(request *DescribeApiSecAttackSourceListRequest) (response *DescribeApiSecAttackSourceListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecAttackSourceListRequest()
	}
	response = NewDescribeApiSecAttackSourceListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyObjectsRequest() (request *ModifyObjectsRequest) {
	request = &ModifyObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyObjects")
	return
}

func NewModifyObjectsResponse() (response *ModifyObjectsResponse) {
	response = &ModifyObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改防护对象
func (c *Client) ModifyObjects(request *ModifyObjectsRequest) (response *ModifyObjectsResponse, err error) {
	if request == nil {
		request = NewModifyObjectsRequest()
	}
	response = NewModifyObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiSecEventChangeRequest() (request *ModifyApiSecEventChangeRequest) {
	request = &ModifyApiSecEventChangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiSecEventChange")
	return
}

func NewModifyApiSecEventChangeResponse() (response *ModifyApiSecEventChangeResponse) {
	response = &ModifyApiSecEventChangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全状态变更接口
func (c *Client) ModifyApiSecEventChange(request *ModifyApiSecEventChangeRequest) (response *ModifyApiSecEventChangeResponse, err error) {
	if request == nil {
		request = NewModifyApiSecEventChangeRequest()
	}
	response = NewModifyApiSecEventChangeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppletAccessV3Request() (request *DeleteAppletAccessV3Request) {
	request = &DeleteAppletAccessV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAppletAccessV3")
	return
}

func NewDeleteAppletAccessV3Response() (response *DeleteAppletAccessV3Response) {
	response = &DeleteAppletAccessV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除小程序接入3.0
func (c *Client) DeleteAppletAccessV3(request *DeleteAppletAccessV3Request) (response *DeleteAppletAccessV3Response, err error) {
	if request == nil {
		request = NewDeleteAppletAccessV3Request()
	}
	response = NewDeleteAppletAccessV3Response()
	err = c.Send(request, response)
	return
}

func NewDetectAppletTesterTaskRequest() (request *DetectAppletTesterTaskRequest) {
	request = &DetectAppletTesterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DetectAppletTesterTask")
	return
}

func NewDetectAppletTesterTaskResponse() (response *DetectAppletTesterTaskResponse) {
	response = &DetectAppletTesterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调试小程序压测任务
func (c *Client) DetectAppletTesterTask(request *DetectAppletTesterTaskRequest) (response *DetectAppletTesterTaskResponse, err error) {
	if request == nil {
		request = NewDetectAppletTesterTaskRequest()
	}
	response = NewDetectAppletTesterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOwaspRuleTypeStatusRequest() (request *ModifyOwaspRuleTypeStatusRequest) {
	request = &ModifyOwaspRuleTypeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOwaspRuleTypeStatus")
	return
}

func NewModifyOwaspRuleTypeStatusResponse() (response *ModifyOwaspRuleTypeStatusResponse) {
	response = &ModifyOwaspRuleTypeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新规则类型的开关
func (c *Client) ModifyOwaspRuleTypeStatus(request *ModifyOwaspRuleTypeStatusRequest) (response *ModifyOwaspRuleTypeStatusResponse, err error) {
	if request == nil {
		request = NewModifyOwaspRuleTypeStatusRequest()
	}
	response = NewModifyOwaspRuleTypeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMLimitRuleListRequest() (request *DescribeLLMLimitRuleListRequest) {
	request = &DescribeLLMLimitRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMLimitRuleList")
	return
}

func NewDescribeLLMLimitRuleListResponse() (response *DescribeLLMLimitRuleListResponse) {
	response = &DescribeLLMLimitRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取LLM算力消耗规则
func (c *Client) DescribeLLMLimitRuleList(request *DescribeLLMLimitRuleListRequest) (response *DescribeLLMLimitRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeLLMLimitRuleListRequest()
	}
	response = NewDescribeLLMLimitRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRateLimitsV2Request() (request *DeleteRateLimitsV2Request) {
	request = &DeleteRateLimitsV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteRateLimitsV2")
	return
}

func NewDeleteRateLimitsV2Response() (response *DeleteRateLimitsV2Response) {
	response = &DeleteRateLimitsV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自研版限流规则
func (c *Client) DeleteRateLimitsV2(request *DeleteRateLimitsV2Request) (response *DeleteRateLimitsV2Response, err error) {
	if request == nil {
		request = NewDeleteRateLimitsV2Request()
	}
	response = NewDeleteRateLimitsV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailDownloadTaskRequest() (request *DescribeBotDetailDownloadTaskRequest) {
	request = &DescribeBotDetailDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotDetailDownloadTask")
	return
}

func NewDescribeBotDetailDownloadTaskResponse() (response *DescribeBotDetailDownloadTaskResponse) {
	response = &DescribeBotDetailDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询bot日志详情下载列表
func (c *Client) DescribeBotDetailDownloadTask(request *DescribeBotDetailDownloadTaskRequest) (response *DescribeBotDetailDownloadTaskResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailDownloadTaskRequest()
	}
	response = NewDescribeBotDetailDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackWorldMapRequest() (request *DescribeAttackWorldMapRequest) {
	request = &DescribeAttackWorldMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackWorldMap")
	return
}

func NewDescribeAttackWorldMapResponse() (response *DescribeAttackWorldMapResponse) {
	response = &DescribeAttackWorldMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定时间段对域名攻击
func (c *Client) DescribeAttackWorldMap(request *DescribeAttackWorldMapRequest) (response *DescribeAttackWorldMapResponse, err error) {
	if request == nil {
		request = NewDescribeAttackWorldMapRequest()
	}
	response = NewDescribeAttackWorldMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportAreasRequest() (request *DescribeSupportAreasRequest) {
	request = &DescribeSupportAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeSupportAreas")
	return
}

func NewDescribeSupportAreasResponse() (response *DescribeSupportAreasResponse) {
	response = &DescribeSupportAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 支持的地域列表
func (c *Client) DescribeSupportAreas(request *DescribeSupportAreasRequest) (response *DescribeSupportAreasResponse, err error) {
	if request == nil {
		request = NewDescribeSupportAreasRequest()
	}
	response = NewDescribeSupportAreasResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOwaspWhiteRuleRequest() (request *ModifyOwaspWhiteRuleRequest) {
	request = &ModifyOwaspWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOwaspWhiteRule")
	return
}

func NewModifyOwaspWhiteRuleResponse() (response *ModifyOwaspWhiteRuleResponse) {
	response = &ModifyOwaspWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑规则引擎白名单
func (c *Client) ModifyOwaspWhiteRule(request *ModifyOwaspWhiteRuleRequest) (response *ModifyOwaspWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyOwaspWhiteRuleRequest()
	}
	response = NewModifyOwaspWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLevelRequest() (request *DescribeUserLevelRequest) {
	request = &DescribeUserLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserLevel")
	return
}

func NewDescribeUserLevelResponse() (response *DescribeUserLevelResponse) {
	response = &DescribeUserLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防护规则等级
func (c *Client) DescribeUserLevel(request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
	if request == nil {
		request = NewDescribeUserLevelRequest()
	}
	response = NewDescribeUserLevelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInactivationApiRequest() (request *ModifyInactivationApiRequest) {
	request = &ModifyInactivationApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInactivationApi")
	return
}

func NewModifyInactivationApiResponse() (response *ModifyInactivationApiResponse) {
	response = &ModifyInactivationApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全失活api定义接口
func (c *Client) ModifyInactivationApi(request *ModifyInactivationApiRequest) (response *ModifyInactivationApiResponse, err error) {
	if request == nil {
		request = NewModifyInactivationApiRequest()
	}
	response = NewModifyInactivationApiResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCachePathRequest() (request *CreateCachePathRequest) {
	request = &CreateCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateCachePath")
	return
}

func NewCreateCachePathResponse() (response *CreateCachePathResponse) {
	response = &CreateCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-添加缓存路径
func (c *Client) CreateCachePath(request *CreateCachePathRequest) (response *CreateCachePathResponse, err error) {
	if request == nil {
		request = NewCreateCachePathRequest()
	}
	response = NewCreateCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaProtectionModeRequest() (request *ModifySpartaProtectionModeRequest) {
	request = &ModifySpartaProtectionModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifySpartaProtectionMode")
	return
}

func NewModifySpartaProtectionModeResponse() (response *ModifySpartaProtectionModeResponse) {
	response = &ModifySpartaProtectionModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置waf防护状态
func (c *Client) ModifySpartaProtectionMode(request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
	if request == nil {
		request = NewModifySpartaProtectionModeRequest()
	}
	response = NewModifySpartaProtectionModeResponse()
	err = c.Send(request, response)
	return
}

func NewAddGroupProtectionRequest() (request *AddGroupProtectionRequest) {
	request = &AddGroupProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddGroupProtection")
	return
}

func NewAddGroupProtectionResponse() (response *AddGroupProtectionResponse) {
	response = &AddGroupProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 收到qps自动告警，完成vip自动限速、迁移、回滚等操作
func (c *Client) AddGroupProtection(request *AddGroupProtectionRequest) (response *AddGroupProtectionResponse, err error) {
	if request == nil {
		request = NewAddGroupProtectionRequest()
	}
	response = NewAddGroupProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleRequest() (request *DescribeRuleRequest) {
	request = &DescribeRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRule")
	return
}

func NewDescribeRuleResponse() (response *DescribeRuleResponse) {
	response = &DescribeRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示tiga规则库详细列表
func (c *Client) DescribeRule(request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
	if request == nil {
		request = NewDescribeRuleRequest()
	}
	response = NewDescribeRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppletTesterTaskRequest() (request *CreateAppletTesterTaskRequest) {
	request = &CreateAppletTesterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAppletTesterTask")
	return
}

func NewCreateAppletTesterTaskResponse() (response *CreateAppletTesterTaskResponse) {
	response = &CreateAppletTesterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加小程序压测任务
func (c *Client) CreateAppletTesterTask(request *CreateAppletTesterTaskRequest) (response *CreateAppletTesterTaskResponse, err error) {
	if request == nil {
		request = NewCreateAppletTesterTaskRequest()
	}
	response = NewCreateAppletTesterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneActionRuleRequest() (request *DeleteBotSceneActionRuleRequest) {
	request = &DeleteBotSceneActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBotSceneActionRule")
	return
}

func NewDeleteBotSceneActionRuleResponse() (response *DeleteBotSceneActionRuleResponse) {
	response = &DeleteBotSceneActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除BOT场景的动作策略
func (c *Client) DeleteBotSceneActionRule(request *DeleteBotSceneActionRuleRequest) (response *DeleteBotSceneActionRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneActionRuleRequest()
	}
	response = NewDeleteBotSceneActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWafAutoDenyStatusRequest() (request *ModifyWafAutoDenyStatusRequest) {
	request = &ModifyWafAutoDenyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyWafAutoDenyStatus")
	return
}

func NewModifyWafAutoDenyStatusResponse() (response *ModifyWafAutoDenyStatusResponse) {
	response = &ModifyWafAutoDenyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置WAF自动封禁模块状态
func (c *Client) ModifyWafAutoDenyStatus(request *ModifyWafAutoDenyStatusRequest) (response *ModifyWafAutoDenyStatusResponse, err error) {
	if request == nil {
		request = NewModifyWafAutoDenyStatusRequest()
	}
	response = NewModifyWafAutoDenyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleTypesRequest() (request *DescribeRuleTypesRequest) {
	request = &DescribeRuleTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRuleTypes")
	return
}

func NewDescribeRuleTypesResponse() (response *DescribeRuleTypesResponse) {
	response = &DescribeRuleTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取规则类型描述
func (c *Client) DescribeRuleTypes(request *DescribeRuleTypesRequest) (response *DescribeRuleTypesResponse, err error) {
	if request == nil {
		request = NewDescribeRuleTypesRequest()
	}
	response = NewDescribeRuleTypesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDomainCustomRuleInfoRequest() (request *QueryDomainCustomRuleInfoRequest) {
	request = &QueryDomainCustomRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryDomainCustomRuleInfo")
	return
}

func NewQueryDomainCustomRuleInfoResponse() (response *QueryDomainCustomRuleInfoResponse) {
	response = &QueryDomainCustomRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-查询域名自定义规则信息接口
func (c *Client) QueryDomainCustomRuleInfo(request *QueryDomainCustomRuleInfoRequest) (response *QueryDomainCustomRuleInfoResponse, err error) {
	if request == nil {
		request = NewQueryDomainCustomRuleInfoRequest()
	}
	response = NewQueryDomainCustomRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppletTesterTaskRequest() (request *DeleteAppletTesterTaskRequest) {
	request = &DeleteAppletTesterTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAppletTesterTask")
	return
}

func NewDeleteAppletTesterTaskResponse() (response *DeleteAppletTesterTaskResponse) {
	response = &DeleteAppletTesterTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除小程序压测任务
func (c *Client) DeleteAppletTesterTask(request *DeleteAppletTesterTaskRequest) (response *DeleteAppletTesterTaskResponse, err error) {
	if request == nil {
		request = NewDeleteAppletTesterTaskRequest()
	}
	response = NewDeleteAppletTesterTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotLevelRequest() (request *DescribeBotLevelRequest) {
	request = &DescribeBotLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotLevel")
	return
}

func NewDescribeBotLevelResponse() (response *DescribeBotLevelResponse) {
	response = &DescribeBotLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot等级配置
func (c *Client) DescribeBotLevel(request *DescribeBotLevelRequest) (response *DescribeBotLevelResponse, err error) {
	if request == nil {
		request = NewDescribeBotLevelRequest()
	}
	response = NewDescribeBotLevelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpAccessControlRequest() (request *DescribeIpAccessControlRequest) {
	request = &DescribeIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeIpAccessControl")
	return
}

func NewDescribeIpAccessControlResponse() (response *DescribeIpAccessControlResponse) {
	response = &DescribeIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf ip黑白名单查询
func (c *Client) DescribeIpAccessControl(request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDescribeIpAccessControlRequest()
	}
	response = NewDescribeIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafInfoByVipRequest() (request *DescribeWafInfoByVipRequest) {
	request = &DescribeWafInfoByVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafInfoByVip")
	return
}

func NewDescribeWafInfoByVipResponse() (response *DescribeWafInfoByVipResponse) {
	response = &DescribeWafInfoByVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户的VIP查询用户的相关信息
func (c *Client) DescribeWafInfoByVip(request *DescribeWafInfoByVipRequest) (response *DescribeWafInfoByVipResponse, err error) {
	if request == nil {
		request = NewDescribeWafInfoByVipRequest()
	}
	response = NewDescribeWafInfoByVipResponse()
	err = c.Send(request, response)
	return
}

func NewSyncClbAssetsRequest() (request *SyncClbAssetsRequest) {
	request = &SyncClbAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SyncClbAssets")
	return
}

func NewSyncClbAssetsResponse() (response *SyncClbAssetsResponse) {
	response = &SyncClbAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发起同步资产请求
func (c *Client) SyncClbAssets(request *SyncClbAssetsRequest) (response *SyncClbAssetsResponse, err error) {
	if request == nil {
		request = NewSyncClbAssetsRequest()
	}
	response = NewSyncClbAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGlobalWhiteRuleRequest() (request *DeleteGlobalWhiteRuleRequest) {
	request = &DeleteGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteGlobalWhiteRule")
	return
}

func NewDeleteGlobalWhiteRuleResponse() (response *DeleteGlobalWhiteRuleResponse) {
	response = &DeleteGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，删除Tiga规则引擎全局白名单。
func (c *Client) DeleteGlobalWhiteRule(request *DeleteGlobalWhiteRuleRequest) (response *DeleteGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteGlobalWhiteRuleRequest()
	}
	response = NewDeleteGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewNotifyOwaspRuleUpdateRequest() (request *NotifyOwaspRuleUpdateRequest) {
	request = &NotifyOwaspRuleUpdateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "NotifyOwaspRuleUpdate")
	return
}

func NewNotifyOwaspRuleUpdateResponse() (response *NotifyOwaspRuleUpdateResponse) {
	response = &NotifyOwaspRuleUpdateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 门神引擎在规则更新的时候通知云WAF的接口（内部使用）
func (c *Client) NotifyOwaspRuleUpdate(request *NotifyOwaspRuleUpdateRequest) (response *NotifyOwaspRuleUpdateResponse, err error) {
	if request == nil {
		request = NewNotifyOwaspRuleUpdateRequest()
	}
	response = NewNotifyOwaspRuleUpdateResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchDomainWhiteRulesRequest() (request *SwitchDomainWhiteRulesRequest) {
	request = &SwitchDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchDomainWhiteRules")
	return
}

func NewSwitchDomainWhiteRulesResponse() (response *SwitchDomainWhiteRulesResponse) {
	response = &SwitchDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换域名白名单的开关
func (c *Client) SwitchDomainWhiteRules(request *SwitchDomainWhiteRulesRequest) (response *SwitchDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewSwitchDomainWhiteRulesRequest()
	}
	response = NewSwitchDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDownloadRecordsRequest() (request *DescribeAttackDownloadRecordsRequest) {
	request = &DescribeAttackDownloadRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackDownloadRecords")
	return
}

func NewDescribeAttackDownloadRecordsResponse() (response *DescribeAttackDownloadRecordsResponse) {
	response = &DescribeAttackDownloadRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询攻击日志的下载记录
func (c *Client) DescribeAttackDownloadRecords(request *DescribeAttackDownloadRecordsRequest) (response *DescribeAttackDownloadRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDownloadRecordsRequest()
	}
	response = NewDescribeAttackDownloadRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletAccessV3Request() (request *ModifyAppletAccessV3Request) {
	request = &ModifyAppletAccessV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletAccessV3")
	return
}

func NewModifyAppletAccessV3Response() (response *ModifyAppletAccessV3Response) {
	response = &ModifyAppletAccessV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改小程序接入3.0
func (c *Client) ModifyAppletAccessV3(request *ModifyAppletAccessV3Request) (response *ModifyAppletAccessV3Response, err error) {
	if request == nil {
		request = NewModifyAppletAccessV3Request()
	}
	response = NewModifyAppletAccessV3Response()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneUAStatusRequest() (request *ModifyBotSceneUAStatusRequest) {
	request = &ModifyBotSceneUAStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotSceneUAStatus")
	return
}

func NewModifyBotSceneUAStatusResponse() (response *ModifyBotSceneUAStatusResponse) {
	response = &ModifyBotSceneUAStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关Bot场景的UA策略开关
func (c *Client) ModifyBotSceneUAStatus(request *ModifyBotSceneUAStatusRequest) (response *ModifyBotSceneUAStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneUAStatusRequest()
	}
	response = NewModifyBotSceneUAStatusResponse()
	err = c.Send(request, response)
	return
}

func NewBatchOperateUserSignatureRulesRequest() (request *BatchOperateUserSignatureRulesRequest) {
	request = &BatchOperateUserSignatureRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "BatchOperateUserSignatureRules")
	return
}

func NewBatchOperateUserSignatureRulesResponse() (response *BatchOperateUserSignatureRulesResponse) {
	response = &BatchOperateUserSignatureRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量操作tiga子规则
func (c *Client) BatchOperateUserSignatureRules(request *BatchOperateUserSignatureRulesRequest) (response *BatchOperateUserSignatureRulesResponse, err error) {
	if request == nil {
		request = NewBatchOperateUserSignatureRulesRequest()
	}
	response = NewBatchOperateUserSignatureRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionedIpRequest() (request *DescribeActionedIpRequest) {
	request = &DescribeActionedIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeActionedIp")
	return
}

func NewDescribeActionedIpResponse() (response *DescribeActionedIpResponse) {
	response = &DescribeActionedIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf Ip状态统一查询接口
func (c *Client) DescribeActionedIp(request *DescribeActionedIpRequest) (response *DescribeActionedIpResponse, err error) {
	if request == nil {
		request = NewDescribeActionedIpRequest()
	}
	response = NewDescribeActionedIpResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBotSceneUCBRuleRequest() (request *DeleteBotSceneUCBRuleRequest) {
	request = &DeleteBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBotSceneUCBRule")
	return
}

func NewDeleteBotSceneUCBRuleResponse() (response *DeleteBotSceneUCBRuleResponse) {
	response = &DeleteBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 场景化后删除Bot的UCB自定义规则
func (c *Client) DeleteBotSceneUCBRule(request *DeleteBotSceneUCBRuleRequest) (response *DeleteBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBotSceneUCBRuleRequest()
	}
	response = NewDeleteBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDownloadRecordRequest() (request *DeleteDownloadRecordRequest) {
	request = &DeleteDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteDownloadRecord")
	return
}

func NewDeleteDownloadRecordResponse() (response *DeleteDownloadRecordResponse) {
	response = &DeleteDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问日志下载记录
func (c *Client) DeleteDownloadRecord(request *DeleteDownloadRecordRequest) (response *DeleteDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteDownloadRecordRequest()
	}
	response = NewDeleteDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsRegionRequest() (request *DescribeTsRegionRequest) {
	request = &DescribeTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTsRegion")
	return
}

func NewDescribeTsRegionResponse() (response *DescribeTsRegionResponse) {
	response = &DescribeTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ts接入区域
func (c *Client) DescribeTsRegion(request *DescribeTsRegionRequest) (response *DescribeTsRegionResponse, err error) {
	if request == nil {
		request = NewDescribeTsRegionRequest()
	}
	response = NewDescribeTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVipInfoRequest() (request *DescribeVipInfoRequest) {
	request = &DescribeVipInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeVipInfo")
	return
}

func NewDescribeVipInfoResponse() (response *DescribeVipInfoResponse) {
	response = &DescribeVipInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据过滤条件查询VIP信息
func (c *Client) DescribeVipInfo(request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVipInfoRequest()
	}
	response = NewDescribeVipInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletAccessRequest() (request *ModifyAppletAccessRequest) {
	request = &ModifyAppletAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletAccess")
	return
}

func NewModifyAppletAccessResponse() (response *ModifyAppletAccessResponse) {
	response = &ModifyAppletAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改小程序接入
func (c *Client) ModifyAppletAccess(request *ModifyAppletAccessRequest) (response *ModifyAppletAccessResponse, err error) {
	if request == nil {
		request = NewModifyAppletAccessRequest()
	}
	response = NewModifyAppletAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaBanAreasRequest() (request *DescribeAreaBanAreasRequest) {
	request = &DescribeAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAreaBanAreas")
	return
}

func NewDescribeAreaBanAreasResponse() (response *DescribeAreaBanAreasResponse) {
	response = &DescribeAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取地域封禁配置包括地域封禁开关，设置封禁的地区信息
func (c *Client) DescribeAreaBanAreas(request *DescribeAreaBanAreasRequest) (response *DescribeAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewDescribeAreaBanAreasRequest()
	}
	response = NewDescribeAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainVerifyResultRequest() (request *DescribeDomainVerifyResultRequest) {
	request = &DescribeDomainVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainVerifyResult")
	return
}

func NewDescribeDomainVerifyResultResponse() (response *DescribeDomainVerifyResultResponse) {
	response = &DescribeDomainVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取添加域名操作的结果
func (c *Client) DescribeDomainVerifyResult(request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeDomainVerifyResultRequest()
	}
	response = NewDescribeDomainVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePocTestTasksRequest() (request *DescribePocTestTasksRequest) {
	request = &DescribePocTestTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePocTestTasks")
	return
}

func NewDescribePocTestTasksResponse() (response *DescribePocTestTasksResponse) {
	response = &DescribePocTestTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询POC测试任务
func (c *Client) DescribePocTestTasks(request *DescribePocTestTasksRequest) (response *DescribePocTestTasksResponse, err error) {
	if request == nil {
		request = NewDescribePocTestTasksRequest()
	}
	response = NewDescribePocTestTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserSignatureRuleRequest() (request *ModifyUserSignatureRuleRequest) {
	request = &ModifyUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserSignatureRule")
	return
}

func NewModifyUserSignatureRuleResponse() (response *ModifyUserSignatureRuleResponse) {
	response = &ModifyUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户防护规则，开启关闭具体的某条规则
func (c *Client) ModifyUserSignatureRule(request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewModifyUserSignatureRuleRequest()
	}
	response = NewModifyUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSDetectDomainListRequest() (request *DescribeDNSDetectDomainListRequest) {
	request = &DescribeDNSDetectDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDNSDetectDomainList")
	return
}

func NewDescribeDNSDetectDomainListResponse() (response *DescribeDNSDetectDomainListResponse) {
	response = &DescribeDNSDetectDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DNS劫持检测的域名列表，对应【Web安全检测】【DNS劫持检测】【基础设置】这个功能。
func (c *Client) DescribeDNSDetectDomainList(request *DescribeDNSDetectDomainListRequest) (response *DescribeDNSDetectDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeDNSDetectDomainListRequest()
	}
	response = NewDescribeDNSDetectDomainListResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackAnalysisRequest() (request *GetAttackAnalysisRequest) {
	request = &GetAttackAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackAnalysis")
	return
}

func NewGetAttackAnalysisResponse() (response *GetAttackAnalysisResponse) {
	response = &GetAttackAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLS新版本攻击日志聚合分析排序
func (c *Client) GetAttackAnalysis(request *GetAttackAnalysisRequest) (response *GetAttackAnalysisResponse, err error) {
	if request == nil {
		request = NewGetAttackAnalysisRequest()
	}
	response = NewGetAttackAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainCountInfoRequest() (request *DescribeDomainCountInfoRequest) {
	request = &DescribeDomainCountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainCountInfo")
	return
}

func NewDescribeDomainCountInfoResponse() (response *DescribeDomainCountInfoResponse) {
	response = &DescribeDomainCountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名概况
func (c *Client) DescribeDomainCountInfo(request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDomainCountInfoRequest()
	}
	response = NewDescribeDomainCountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAttackIpListRequest() (request *DescribeBotAttackIpListRequest) {
	request = &DescribeBotAttackIpListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAttackIpList")
	return
}

func NewDescribeBotAttackIpListResponse() (response *DescribeBotAttackIpListResponse) {
	response = &DescribeBotAttackIpListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-获取BOT事件攻击源IP详情
func (c *Client) DescribeBotAttackIpList(request *DescribeBotAttackIpListRequest) (response *DescribeBotAttackIpListResponse, err error) {
	if request == nil {
		request = NewDescribeBotAttackIpListRequest()
	}
	response = NewDescribeBotAttackIpListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCachePathRequest() (request *DeleteCachePathRequest) {
	request = &DeleteCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCachePath")
	return
}

func NewDeleteCachePathResponse() (response *DeleteCachePathResponse) {
	response = &DeleteCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-删除缓存路径
func (c *Client) DeleteCachePath(request *DeleteCachePathRequest) (response *DeleteCachePathResponse, err error) {
	if request == nil {
		request = NewDeleteCachePathRequest()
	}
	response = NewDeleteCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainInstanceIDsRequest() (request *DescribeDomainInstanceIDsRequest) {
	request = &DescribeDomainInstanceIDsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainInstanceIDs")
	return
}

func NewDescribeDomainInstanceIDsResponse() (response *DescribeDomainInstanceIDsResponse) {
	response = &DescribeDomainInstanceIDsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取添加某个域名的所有实例id
func (c *Client) DescribeDomainInstanceIDs(request *DescribeDomainInstanceIDsRequest) (response *DescribeDomainInstanceIDsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainInstanceIDsRequest()
	}
	response = NewDescribeDomainInstanceIDsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPocTestAuthorizationRequest() (request *ModifyPocTestAuthorizationRequest) {
	request = &ModifyPocTestAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyPocTestAuthorization")
	return
}

func NewModifyPocTestAuthorizationResponse() (response *ModifyPocTestAuthorizationResponse) {
	response = &ModifyPocTestAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改APPID授权状态
func (c *Client) ModifyPocTestAuthorization(request *ModifyPocTestAuthorizationRequest) (response *ModifyPocTestAuthorizationResponse, err error) {
	if request == nil {
		request = NewModifyPocTestAuthorizationRequest()
	}
	response = NewModifyPocTestAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertBotUCBFeatureRuleRequest() (request *UpsertBotUCBFeatureRuleRequest) {
	request = &UpsertBotUCBFeatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpsertBotUCBFeatureRule")
	return
}

func NewUpsertBotUCBFeatureRuleResponse() (response *UpsertBotUCBFeatureRuleResponse) {
	response = &UpsertBotUCBFeatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 UCB策略更新
func (c *Client) UpsertBotUCBFeatureRule(request *UpsertBotUCBFeatureRuleRequest) (response *UpsertBotUCBFeatureRuleResponse, err error) {
	if request == nil {
		request = NewUpsertBotUCBFeatureRuleRequest()
	}
	response = NewUpsertBotUCBFeatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaBanAreasRequest() (request *ModifyAreaBanAreasRequest) {
	request = &ModifyAreaBanAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAreaBanAreas")
	return
}

func NewModifyAreaBanAreasResponse() (response *ModifyAreaBanAreasResponse) {
	response = &ModifyAreaBanAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改地域封禁中的地域信息
func (c *Client) ModifyAreaBanAreas(request *ModifyAreaBanAreasRequest) (response *ModifyAreaBanAreasResponse, err error) {
	if request == nil {
		request = NewModifyAreaBanAreasRequest()
	}
	response = NewModifyAreaBanAreasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppletAuthorizeV3Request() (request *CreateAppletAuthorizeV3Request) {
	request = &CreateAppletAuthorizeV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAppletAuthorizeV3")
	return
}

func NewCreateAppletAuthorizeV3Response() (response *CreateAppletAuthorizeV3Response) {
	response = &CreateAppletAuthorizeV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加小程序授权3.0
func (c *Client) CreateAppletAuthorizeV3(request *CreateAppletAuthorizeV3Request) (response *CreateAppletAuthorizeV3Response, err error) {
	if request == nil {
		request = NewCreateAppletAuthorizeV3Request()
	}
	response = NewCreateAppletAuthorizeV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockPagesRequest() (request *DescribeBlockPagesRequest) {
	request = &DescribeBlockPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBlockPages")
	return
}

func NewDescribeBlockPagesResponse() (response *DescribeBlockPagesResponse) {
	response = &DescribeBlockPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示自定义拦截页面列表信息。
func (c *Client) DescribeBlockPages(request *DescribeBlockPagesRequest) (response *DescribeBlockPagesResponse, err error) {
	if request == nil {
		request = NewDescribeBlockPagesRequest()
	}
	response = NewDescribeBlockPagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiListVersionTwoRequest() (request *DescribeApiListVersionTwoRequest) {
	request = &DescribeApiListVersionTwoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiListVersionTwo")
	return
}

func NewDescribeApiListVersionTwoResponse() (response *DescribeApiListVersionTwoResponse) {
	response = &DescribeApiListVersionTwoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api资产列表
func (c *Client) DescribeApiListVersionTwo(request *DescribeApiListVersionTwoRequest) (response *DescribeApiListVersionTwoResponse, err error) {
	if request == nil {
		request = NewDescribeApiListVersionTwoRequest()
	}
	response = NewDescribeApiListVersionTwoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePeakPointsRequest() (request *DescribePeakPointsRequest) {
	request = &DescribePeakPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePeakPoints")
	return
}

func NewDescribePeakPointsResponse() (response *DescribePeakPointsResponse) {
	response = &DescribePeakPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务和攻击概要趋势
func (c *Client) DescribePeakPoints(request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
	if request == nil {
		request = NewDescribePeakPointsRequest()
	}
	response = NewDescribePeakPointsResponse()
	err = c.Send(request, response)
	return
}

func NewAddGlobalWhiteRuleRequest() (request *AddGlobalWhiteRuleRequest) {
	request = &AddGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddGlobalWhiteRule")
	return
}

func NewAddGlobalWhiteRuleResponse() (response *AddGlobalWhiteRuleResponse) {
	response = &AddGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供内部运营系统调用，增加Tiga规则引擎全局白名单列表。
func (c *Client) AddGlobalWhiteRule(request *AddGlobalWhiteRuleRequest) (response *AddGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddGlobalWhiteRuleRequest()
	}
	response = NewAddGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertIpAccessControlRequest() (request *UpsertIpAccessControlRequest) {
	request = &UpsertIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpsertIpAccessControl")
	return
}

func NewUpsertIpAccessControlResponse() (response *UpsertIpAccessControlResponse) {
	response = &UpsertIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单Upsert接口
func (c *Client) UpsertIpAccessControl(request *UpsertIpAccessControlRequest) (response *UpsertIpAccessControlResponse, err error) {
	if request == nil {
		request = NewUpsertIpAccessControlRequest()
	}
	response = NewUpsertIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFirstPurchaseRequest() (request *DescribeFirstPurchaseRequest) {
	request = &DescribeFirstPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFirstPurchase")
	return
}

func NewDescribeFirstPurchaseResponse() (response *DescribeFirstPurchaseResponse) {
	response = &DescribeFirstPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否为第一次购买waf实例，api接口会调用计费侧查询订单接口，如果订单数量totalNum大于0，则认为不是首购用户，计费侧根据ownerUin查询订单
func (c *Client) DescribeFirstPurchase(request *DescribeFirstPurchaseRequest) (response *DescribeFirstPurchaseResponse, err error) {
	if request == nil {
		request = NewDescribeFirstPurchaseRequest()
	}
	response = NewDescribeFirstPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotApiEventAlarmRequest() (request *ModifyBotApiEventAlarmRequest) {
	request = &ModifyBotApiEventAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotApiEventAlarm")
	return
}

func NewModifyBotApiEventAlarmResponse() (response *ModifyBotApiEventAlarmResponse) {
	response = &ModifyBotApiEventAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot/api事件告警配置
func (c *Client) ModifyBotApiEventAlarm(request *ModifyBotApiEventAlarmRequest) (response *ModifyBotApiEventAlarmResponse, err error) {
	if request == nil {
		request = NewModifyBotApiEventAlarmRequest()
	}
	response = NewModifyBotApiEventAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomRuleActionRequest() (request *ModifyCustomRuleActionRequest) {
	request = &ModifyCustomRuleActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCustomRuleAction")
	return
}

func NewModifyCustomRuleActionResponse() (response *ModifyCustomRuleActionResponse) {
	response = &ModifyCustomRuleActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 匹配修改访问控制动作（自定义策略）
func (c *Client) ModifyCustomRuleAction(request *ModifyCustomRuleActionRequest) (response *ModifyCustomRuleActionResponse, err error) {
	if request == nil {
		request = NewModifyCustomRuleActionRequest()
	}
	response = NewModifyCustomRuleActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCCAutoStatusRequest() (request *DescribeCCAutoStatusRequest) {
	request = &DescribeCCAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCCAutoStatus")
	return
}

func NewDescribeCCAutoStatusResponse() (response *DescribeCCAutoStatusResponse) {
	response = &DescribeCCAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本查询cc自动封堵状态
func (c *Client) DescribeCCAutoStatus(request *DescribeCCAutoStatusRequest) (response *DescribeCCAutoStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCCAutoStatusRequest()
	}
	response = NewDescribeCCAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowCheckUrlRequest() (request *DescribeFlowCheckUrlRequest) {
	request = &DescribeFlowCheckUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeFlowCheckUrl")
	return
}

func NewDescribeFlowCheckUrlResponse() (response *DescribeFlowCheckUrlResponse) {
	response = &DescribeFlowCheckUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [流量体检]-获取流量体检Url接口
func (c *Client) DescribeFlowCheckUrl(request *DescribeFlowCheckUrlRequest) (response *DescribeFlowCheckUrlResponse, err error) {
	if request == nil {
		request = NewDescribeFlowCheckUrlRequest()
	}
	response = NewDescribeFlowCheckUrlResponse()
	err = c.Send(request, response)
	return
}

func NewAddDomainSgTaskRequest() (request *AddDomainSgTaskRequest) {
	request = &AddDomainSgTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddDomainSgTask")
	return
}

func NewAddDomainSgTaskResponse() (response *AddDomainSgTaskResponse) {
	response = &AddDomainSgTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加安全组任务 为一个域名
func (c *Client) AddDomainSgTask(request *AddDomainSgTaskRequest) (response *AddDomainSgTaskResponse, err error) {
	if request == nil {
		request = NewAddDomainSgTaskRequest()
	}
	response = NewAddDomainSgTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePocTestTaskRequest() (request *DeletePocTestTaskRequest) {
	request = &DeletePocTestTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeletePocTestTask")
	return
}

func NewDeletePocTestTaskResponse() (response *DeletePocTestTaskResponse) {
	response = &DeletePocTestTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除POC测试任务
func (c *Client) DeletePocTestTask(request *DeletePocTestTaskRequest) (response *DeletePocTestTaskResponse, err error) {
	if request == nil {
		request = NewDeletePocTestTaskRequest()
	}
	response = NewDeletePocTestTaskResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRateLimitsRequest() (request *EnableRateLimitsRequest) {
	request = &EnableRateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EnableRateLimits")
	return
}

func NewEnableRateLimitsResponse() (response *EnableRateLimitsResponse) {
	response = &EnableRateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量启用/禁用限流规则
func (c *Client) EnableRateLimits(request *EnableRateLimitsRequest) (response *EnableRateLimitsResponse, err error) {
	if request == nil {
		request = NewEnableRateLimitsRequest()
	}
	response = NewEnableRateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotTldSdkChannelRequest() (request *DescribeBotTldSdkChannelRequest) {
	request = &DescribeBotTldSdkChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotTldSdkChannel")
	return
}

func NewDescribeBotTldSdkChannelResponse() (response *DescribeBotTldSdkChannelResponse) {
	response = &DescribeBotTldSdkChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示图灵盾所使用的sdk_channel
func (c *Client) DescribeBotTldSdkChannel(request *DescribeBotTldSdkChannelRequest) (response *DescribeBotTldSdkChannelResponse, err error) {
	if request == nil {
		request = NewDescribeBotTldSdkChannelRequest()
	}
	response = NewDescribeBotTldSdkChannelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMSessionRequest() (request *DescribeLLMSessionRequest) {
	request = &DescribeLLMSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMSession")
	return
}

func NewDescribeLLMSessionResponse() (response *DescribeLLMSessionResponse) {
	response = &DescribeLLMSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取LLMsession配置列表
func (c *Client) DescribeLLMSession(request *DescribeLLMSessionRequest) (response *DescribeLLMSessionResponse, err error) {
	if request == nil {
		request = NewDescribeLLMSessionRequest()
	}
	response = NewDescribeLLMSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMProtectPathRequest() (request *DescribeLLMProtectPathRequest) {
	request = &DescribeLLMProtectPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMProtectPath")
	return
}

func NewDescribeLLMProtectPathResponse() (response *DescribeLLMProtectPathResponse) {
	response = &DescribeLLMProtectPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取大模型防护路径设置 列表
func (c *Client) DescribeLLMProtectPath(request *DescribeLLMProtectPathRequest) (response *DescribeLLMProtectPathResponse, err error) {
	if request == nil {
		request = NewDescribeLLMProtectPathRequest()
	}
	response = NewDescribeLLMProtectPathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordItemsRequest() (request *DescribeBotRecordItemsRequest) {
	request = &DescribeBotRecordItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRecordItems")
	return
}

func NewDescribeBotRecordItemsResponse() (response *DescribeBotRecordItemsResponse) {
	response = &DescribeBotRecordItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot记录访问详情
func (c *Client) DescribeBotRecordItems(request *DescribeBotRecordItemsRequest) (response *DescribeBotRecordItemsResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordItemsRequest()
	}
	response = NewDescribeBotRecordItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVIPStateRequest() (request *DescribeVIPStateRequest) {
	request = &DescribeVIPStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeVIPState")
	return
}

func NewDescribeVIPStateResponse() (response *DescribeVIPStateResponse) {
	response = &DescribeVIPStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询vip表
func (c *Client) DescribeVIPState(request *DescribeVIPStateRequest) (response *DescribeVIPStateResponse, err error) {
	if request == nil {
		request = NewDescribeVIPStateRequest()
	}
	response = NewDescribeVIPStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJsInjectRuleRequest() (request *DescribeJsInjectRuleRequest) {
	request = &DescribeJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeJsInjectRule")
	return
}

func NewDescribeJsInjectRuleResponse() (response *DescribeJsInjectRuleResponse) {
	response = &DescribeJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对应域名的前端对抗规则。
func (c *Client) DescribeJsInjectRule(request *DescribeJsInjectRuleRequest) (response *DescribeJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewDescribeJsInjectRuleRequest()
	}
	response = NewDescribeJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAntiInfoLeakRuleRequest() (request *DeleteAntiInfoLeakRuleRequest) {
	request = &DeleteAntiInfoLeakRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAntiInfoLeakRule")
	return
}

func NewDeleteAntiInfoLeakRuleResponse() (response *DeleteAntiInfoLeakRuleResponse) {
	response = &DeleteAntiInfoLeakRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 信息防泄漏删除规则
func (c *Client) DeleteAntiInfoLeakRule(request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAntiInfoLeakRuleRequest()
	}
	response = NewDeleteAntiInfoLeakRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHybridClusterLogRequest() (request *ModifyHybridClusterLogRequest) {
	request = &ModifyHybridClusterLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHybridClusterLog")
	return
}

func NewModifyHybridClusterLogResponse() (response *ModifyHybridClusterLogResponse) {
	response = &ModifyHybridClusterLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改混合云日志投递设置
func (c *Client) ModifyHybridClusterLog(request *ModifyHybridClusterLogRequest) (response *ModifyHybridClusterLogResponse, err error) {
	if request == nil {
		request = NewModifyHybridClusterLogRequest()
	}
	response = NewModifyHybridClusterLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotRecordPointsRequest() (request *DescribeBotRecordPointsRequest) {
	request = &DescribeBotRecordPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotRecordPoints")
	return
}

func NewDescribeBotRecordPointsResponse() (response *DescribeBotRecordPointsResponse) {
	response = &DescribeBotRecordPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot记录访问趋势图
func (c *Client) DescribeBotRecordPoints(request *DescribeBotRecordPointsRequest) (response *DescribeBotRecordPointsResponse, err error) {
	if request == nil {
		request = NewDescribeBotRecordPointsRequest()
	}
	response = NewDescribeBotRecordPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserTypeRequest() (request *DescribeUserTypeRequest) {
	request = &DescribeUserTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserType")
	return
}

func NewDescribeUserTypeResponse() (response *DescribeUserTypeResponse) {
	response = &DescribeUserTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户类型
func (c *Client) DescribeUserType(request *DescribeUserTypeRequest) (response *DescribeUserTypeResponse, err error) {
	if request == nil {
		request = NewDescribeUserTypeRequest()
	}
	response = NewDescribeUserTypeResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleCosKeyRequest() (request *GetRuleCosKeyRequest) {
	request = &GetRuleCosKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetRuleCosKey")
	return
}

func NewGetRuleCosKeyResponse() (response *GetRuleCosKeyResponse) {
	response = &GetRuleCosKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取cos临时密钥，用于上传tiga规则库
func (c *Client) GetRuleCosKey(request *GetRuleCosKeyRequest) (response *GetRuleCosKeyResponse, err error) {
	if request == nil {
		request = NewGetRuleCosKeyRequest()
	}
	response = NewGetRuleCosKeyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRateLimitsRequest() (request *UpdateRateLimitsRequest) {
	request = &UpdateRateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpdateRateLimits")
	return
}

func NewUpdateRateLimitsResponse() (response *UpdateRateLimitsResponse) {
	response = &UpdateRateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新限流规则
func (c *Client) UpdateRateLimits(request *UpdateRateLimitsRequest) (response *UpdateRateLimitsResponse, err error) {
	if request == nil {
		request = NewUpdateRateLimitsRequest()
	}
	response = NewUpdateRateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttackCountRequest() (request *GetAttackCountRequest) {
	request = &GetAttackCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetAttackCount")
	return
}

func NewGetAttackCountResponse() (response *GetAttackCountResponse) {
	response = &GetAttackCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意攻击总量
func (c *Client) GetAttackCount(request *GetAttackCountRequest) (response *GetAttackCountResponse, err error) {
	if request == nil {
		request = NewGetAttackCountRequest()
	}
	response = NewGetAttackCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifySceneJsInjectRuleRequest() (request *ModifySceneJsInjectRuleRequest) {
	request = &ModifySceneJsInjectRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifySceneJsInjectRule")
	return
}

func NewModifySceneJsInjectRuleResponse() (response *ModifySceneJsInjectRuleResponse) {
	response = &ModifySceneJsInjectRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】前端对抗在场景化版本中，将配置拆分为全局和场景两部分，[SceneId,Status,RuleAction,Redirect]属于场景配置，[AutomatedTool,CrackBehavior,GoodBot,RasSiteID]属于全局配置，字段RuleID变成复用字段，全局时表示tb_waf_jsinject_rule的id，场景时表示tb_waf_jsinject_scene_rule的id
func (c *Client) ModifySceneJsInjectRule(request *ModifySceneJsInjectRuleRequest) (response *ModifySceneJsInjectRuleResponse, err error) {
	if request == nil {
		request = NewModifySceneJsInjectRuleRequest()
	}
	response = NewModifySceneJsInjectRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchCustomRuleListRequest() (request *DescribeBatchCustomRuleListRequest) {
	request = &DescribeBatchCustomRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchCustomRuleList")
	return
}

func NewDescribeBatchCustomRuleListResponse() (response *DescribeBatchCustomRuleListResponse) {
	response = &DescribeBatchCustomRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [自定义规则]-批量列表接口
func (c *Client) DescribeBatchCustomRuleList(request *DescribeBatchCustomRuleListRequest) (response *DescribeBatchCustomRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeBatchCustomRuleListRequest()
	}
	response = NewDescribeBatchCustomRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotDetailCountTopNRequest() (request *DescribeBotDetailCountTopNRequest) {
	request = &DescribeBotDetailCountTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotDetailCountTopN")
	return
}

func NewDescribeBotDetailCountTopNResponse() (response *DescribeBotDetailCountTopNResponse) {
	response = &DescribeBotDetailCountTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// bot详情中请求数量topN
func (c *Client) DescribeBotDetailCountTopN(request *DescribeBotDetailCountTopNRequest) (response *DescribeBotDetailCountTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotDetailCountTopNRequest()
	}
	response = NewDescribeBotDetailCountTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskInfoRequest() (request *DescribeRiskInfoRequest) {
	request = &DescribeRiskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRiskInfo")
	return
}

func NewDescribeRiskInfoResponse() (response *DescribeRiskInfoResponse) {
	response = &DescribeRiskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务安全风险信息查询
func (c *Client) DescribeRiskInfo(request *DescribeRiskInfoRequest) (response *DescribeRiskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRiskInfoRequest()
	}
	response = NewDescribeRiskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOwaspRuleStatusRequest() (request *DeleteOwaspRuleStatusRequest) {
	request = &DeleteOwaspRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteOwaspRuleStatus")
	return
}

func NewDeleteOwaspRuleStatusResponse() (response *DeleteOwaspRuleStatusResponse) {
	response = &DeleteOwaspRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除门神规则的状态锁
func (c *Client) DeleteOwaspRuleStatus(request *DeleteOwaspRuleStatusRequest) (response *DeleteOwaspRuleStatusResponse, err error) {
	if request == nil {
		request = NewDeleteOwaspRuleStatusRequest()
	}
	response = NewDeleteOwaspRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpMainClassRequest() (request *DescribeOpMainClassRequest) {
	request = &DescribeOpMainClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpMainClass")
	return
}

func NewDescribeOpMainClassResponse() (response *DescribeOpMainClassResponse) {
	response = &DescribeOpMainClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取主类及子类信息
func (c *Client) DescribeOpMainClass(request *DescribeOpMainClassRequest) (response *DescribeOpMainClassResponse, err error) {
	if request == nil {
		request = NewDescribeOpMainClassRequest()
	}
	response = NewDescribeOpMainClassResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
	request = &DescribePortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePorts")
	return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
	response = &DescribePortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Saas型WAF防护端口列表
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
	if request == nil {
		request = NewDescribePortsRequest()
	}
	response = NewDescribePortsResponse()
	err = c.Send(request, response)
	return
}

func NewAddHWRecordRequest() (request *AddHWRecordRequest) {
	request = &AddHWRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddHWRecord")
	return
}

func NewAddHWRecordResponse() (response *AddHWRecordResponse) {
	response = &AddHWRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [重保护网]-新增重保护网记录接口
func (c *Client) AddHWRecord(request *AddHWRecordRequest) (response *AddHWRecordResponse, err error) {
	if request == nil {
		request = NewAddHWRecordRequest()
	}
	response = NewAddHWRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignaturePolicyRequest() (request *DescribeOpUserSignaturePolicyRequest) {
	request = &DescribeOpUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpUserSignaturePolicy")
	return
}

func NewDescribeOpUserSignaturePolicyResponse() (response *DescribeOpUserSignaturePolicyResponse) {
	response = &DescribeOpUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取规则状态
func (c *Client) DescribeOpUserSignaturePolicy(request *DescribeOpUserSignaturePolicyRequest) (response *DescribeOpUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignaturePolicyRequest()
	}
	response = NewDescribeOpUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewPostDomainsConfigTaskRequest() (request *PostDomainsConfigTaskRequest) {
	request = &PostDomainsConfigTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "PostDomainsConfigTask")
	return
}

func NewPostDomainsConfigTaskResponse() (response *PostDomainsConfigTaskResponse) {
	response = &PostDomainsConfigTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载域名配置
func (c *Client) PostDomainsConfigTask(request *PostDomainsConfigTaskRequest) (response *PostDomainsConfigTaskResponse, err error) {
	if request == nil {
		request = NewPostDomainsConfigTaskRequest()
	}
	response = NewPostDomainsConfigTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotEventCauseLossRequest() (request *DescribeBotEventCauseLossRequest) {
	request = &DescribeBotEventCauseLossRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotEventCauseLoss")
	return
}

func NewDescribeBotEventCauseLossResponse() (response *DescribeBotEventCauseLossResponse) {
	response = &DescribeBotEventCauseLossResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-计算BOT事件带来的成本损失
func (c *Client) DescribeBotEventCauseLoss(request *DescribeBotEventCauseLossRequest) (response *DescribeBotEventCauseLossResponse, err error) {
	if request == nil {
		request = NewDescribeBotEventCauseLossRequest()
	}
	response = NewDescribeBotEventCauseLossResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCdcClbWafRegionsRequest() (request *DescribeUserCdcClbWafRegionsRequest) {
	request = &DescribeUserCdcClbWafRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserCdcClbWafRegions")
	return
}

func NewDescribeUserCdcClbWafRegionsResponse() (response *DescribeUserCdcClbWafRegionsResponse) {
	response = &DescribeUserCdcClbWafRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在CDC场景下，负载均衡型WAF的添加、编辑域名配置的时候，需要展示CDC负载均衡型WAF（cdc-clb-waf)支持的地域列表，通过DescribeUserCdcClbWafRegions既可以获得当前对客户已经开放的地域列表
func (c *Client) DescribeUserCdcClbWafRegions(request *DescribeUserCdcClbWafRegionsRequest) (response *DescribeUserCdcClbWafRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeUserCdcClbWafRegionsRequest()
	}
	response = NewDescribeUserCdcClbWafRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDefaultEngineRequest() (request *DescribeUserDefaultEngineRequest) {
	request = &DescribeUserDefaultEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserDefaultEngine")
	return
}

func NewDescribeUserDefaultEngineResponse() (response *DescribeUserDefaultEngineResponse) {
	response = &DescribeUserDefaultEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户缺省规则引擎类型
func (c *Client) DescribeUserDefaultEngine(request *DescribeUserDefaultEngineRequest) (response *DescribeUserDefaultEngineResponse, err error) {
	if request == nil {
		request = NewDescribeUserDefaultEngineRequest()
	}
	response = NewDescribeUserDefaultEngineResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCdcResourceRequest() (request *ModifyCdcResourceRequest) {
	request = &ModifyCdcResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCdcResource")
	return
}

func NewModifyCdcResourceResponse() (response *ModifyCdcResourceResponse) {
	response = &ModifyCdcResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改CDC资源列表，如果id=0则添加，>0则修改
func (c *Client) ModifyCdcResource(request *ModifyCdcResourceRequest) (response *ModifyCdcResourceResponse, err error) {
	if request == nil {
		request = NewModifyCdcResourceRequest()
	}
	response = NewModifyCdcResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
	request = &SearchLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SearchLog")
	return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
	response = &SearchLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于检索分析日志，使用该接口时请注意如下事项：
// 1. 该接口除受默认接口请求频率限制外，针对单个日志主题，查询并发数不能超过15。
// 2. 检索语法建议使用CQL语法规则，请使用SyntaxRule参数，将值设置为1。
// 3. API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
	if request == nil {
		request = NewSearchLogRequest()
	}
	response = NewSearchLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomPayloadsRequest() (request *DeleteCustomPayloadsRequest) {
	request = &DeleteCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCustomPayloads")
	return
}

func NewDeleteCustomPayloadsResponse() (response *DeleteCustomPayloadsResponse) {
	response = &DeleteCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义payload
func (c *Client) DeleteCustomPayloads(request *DeleteCustomPayloadsRequest) (response *DeleteCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomPayloadsRequest()
	}
	response = NewDeleteCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecEventChangeListRequest() (request *DescribeApiSecEventChangeListRequest) {
	request = &DescribeApiSecEventChangeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecEventChangeList")
	return
}

func NewDescribeApiSecEventChangeListResponse() (response *DescribeApiSecEventChangeListResponse) {
	response = &DescribeApiSecEventChangeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取api安全事件变更历史
func (c *Client) DescribeApiSecEventChangeList(request *DescribeApiSecEventChangeListRequest) (response *DescribeApiSecEventChangeListResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecEventChangeListRequest()
	}
	response = NewDescribeApiSecEventChangeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRateLimitsRequest() (request *DescribeRateLimitsRequest) {
	request = &DescribeRateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRateLimits")
	return
}

func NewDescribeRateLimitsResponse() (response *DescribeRateLimitsResponse) {
	response = &DescribeRateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询限流规则列表
func (c *Client) DescribeRateLimits(request *DescribeRateLimitsRequest) (response *DescribeRateLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeRateLimitsRequest()
	}
	response = NewDescribeRateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainWhiteRulesRequest() (request *DescribeDomainWhiteRulesRequest) {
	request = &DescribeDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainWhiteRules")
	return
}

func NewDescribeDomainWhiteRulesResponse() (response *DescribeDomainWhiteRulesResponse) {
	response = &DescribeDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名的规则白名单
func (c *Client) DescribeDomainWhiteRules(request *DescribeDomainWhiteRulesRequest) (response *DescribeDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDescribeDomainWhiteRulesRequest()
	}
	response = NewDescribeDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiDetailSampleRequest() (request *ModifyApiDetailSampleRequest) {
	request = &ModifyApiDetailSampleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyApiDetailSample")
	return
}

func NewModifyApiDetailSampleResponse() (response *ModifyApiDetailSampleResponse) {
	response = &ModifyApiDetailSampleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全api详情样例保存接口
func (c *Client) ModifyApiDetailSample(request *ModifyApiDetailSampleRequest) (response *ModifyApiDetailSampleResponse, err error) {
	if request == nil {
		request = NewModifyApiDetailSampleRequest()
	}
	response = NewModifyApiDetailSampleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotStatusRequest() (request *DescribeBotStatusRequest) {
	request = &DescribeBotStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotStatus")
	return
}

func NewDescribeBotStatusResponse() (response *DescribeBotStatusResponse) {
	response = &DescribeBotStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 获取bot开关
func (c *Client) DescribeBotStatus(request *DescribeBotStatusRequest) (response *DescribeBotStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBotStatusRequest()
	}
	response = NewDescribeBotStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletTesterReserveInfoRequest() (request *ModifyAppletTesterReserveInfoRequest) {
	request = &ModifyAppletTesterReserveInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletTesterReserveInfo")
	return
}

func NewModifyAppletTesterReserveInfoResponse() (response *ModifyAppletTesterReserveInfoResponse) {
	response = &ModifyAppletTesterReserveInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改小程序预约任务详情
func (c *Client) ModifyAppletTesterReserveInfo(request *ModifyAppletTesterReserveInfoRequest) (response *ModifyAppletTesterReserveInfoResponse, err error) {
	if request == nil {
		request = NewModifyAppletTesterReserveInfoRequest()
	}
	response = NewModifyAppletTesterReserveInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotActionRuleRequest() (request *ModifyBotActionRuleRequest) {
	request = &ModifyBotActionRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotActionRule")
	return
}

func NewModifyBotActionRuleResponse() (response *ModifyBotActionRuleResponse) {
	response = &ModifyBotActionRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改bot-动作分布配置
func (c *Client) ModifyBotActionRule(request *ModifyBotActionRuleRequest) (response *ModifyBotActionRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotActionRuleRequest()
	}
	response = NewModifyBotActionRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddNewBindRoleUserRequest() (request *AddNewBindRoleUserRequest) {
	request = &AddNewBindRoleUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddNewBindRoleUser")
	return
}

func NewAddNewBindRoleUserResponse() (response *AddNewBindRoleUserResponse) {
	response = &AddNewBindRoleUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf角色授权绑定接口
func (c *Client) AddNewBindRoleUser(request *AddNewBindRoleUserRequest) (response *AddNewBindRoleUserResponse, err error) {
	if request == nil {
		request = NewAddNewBindRoleUserRequest()
	}
	response = NewAddNewBindRoleUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAreaBanSupportAreasRequest() (request *DescribeAreaBanSupportAreasRequest) {
	request = &DescribeAreaBanSupportAreasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAreaBanSupportAreas")
	return
}

func NewDescribeAreaBanSupportAreasResponse() (response *DescribeAreaBanSupportAreasResponse) {
	response = &DescribeAreaBanSupportAreasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取WAF地域封禁支持的地域列表
func (c *Client) DescribeAreaBanSupportAreas(request *DescribeAreaBanSupportAreasRequest) (response *DescribeAreaBanSupportAreasResponse, err error) {
	if request == nil {
		request = NewDescribeAreaBanSupportAreasRequest()
	}
	response = NewDescribeAreaBanSupportAreasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserEngineTypeRequest() (request *DescribeUserEngineTypeRequest) {
	request = &DescribeUserEngineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserEngineType")
	return
}

func NewDescribeUserEngineTypeResponse() (response *DescribeUserEngineTypeResponse) {
	response = &DescribeUserEngineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户规则引擎类型
func (c *Client) DescribeUserEngineType(request *DescribeUserEngineTypeRequest) (response *DescribeUserEngineTypeResponse, err error) {
	if request == nil {
		request = NewDescribeUserEngineTypeRequest()
	}
	response = NewDescribeUserEngineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostRequest() (request *ModifyHostRequest) {
	request = &ModifyHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHost")
	return
}

func NewModifyHostResponse() (response *ModifyHostResponse) {
	response = &ModifyHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf编辑防护域名配置
func (c *Client) ModifyHost(request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
	if request == nil {
		request = NewModifyHostRequest()
	}
	response = NewModifyHostResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTsRegionRequest() (request *ModifyTsRegionRequest) {
	request = &ModifyTsRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyTsRegion")
	return
}

func NewModifyTsRegionResponse() (response *ModifyTsRegionResponse) {
	response = &ModifyTsRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ts接入区域
func (c *Client) ModifyTsRegion(request *ModifyTsRegionRequest) (response *ModifyTsRegionResponse, err error) {
	if request == nil {
		request = NewModifyTsRegionRequest()
	}
	response = NewModifyTsRegionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostRequest() (request *CreateHostRequest) {
	request = &CreateHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateHost")
	return
}

func NewCreateHostResponse() (response *CreateHostResponse) {
	response = &CreateHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中添加防护域名
func (c *Client) CreateHost(request *CreateHostRequest) (response *CreateHostResponse, err error) {
	if request == nil {
		request = NewCreateHostRequest()
	}
	response = NewCreateHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopAttackDomainRequest() (request *DescribeTopAttackDomainRequest) {
	request = &DescribeTopAttackDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTopAttackDomain")
	return
}

func NewDescribeTopAttackDomainResponse() (response *DescribeTopAttackDomainResponse) {
	response = &DescribeTopAttackDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Top5的攻击域名
func (c *Client) DescribeTopAttackDomain(request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
	if request == nil {
		request = NewDescribeTopAttackDomainRequest()
	}
	response = NewDescribeTopAttackDomainResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchCustomWhiteRuleStatusRequest() (request *ModifyBatchCustomWhiteRuleStatusRequest) {
	request = &ModifyBatchCustomWhiteRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchCustomWhiteRuleStatus")
	return
}

func NewModifyBatchCustomWhiteRuleStatusResponse() (response *ModifyBatchCustomWhiteRuleStatusResponse) {
	response = &ModifyBatchCustomWhiteRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新批量精准白名单规则
func (c *Client) ModifyBatchCustomWhiteRuleStatus(request *ModifyBatchCustomWhiteRuleStatusRequest) (response *ModifyBatchCustomWhiteRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyBatchCustomWhiteRuleStatusRequest()
	}
	response = NewModifyBatchCustomWhiteRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserSignatureClassRequest() (request *ModifyUserSignatureClassRequest) {
	request = &ModifyUserSignatureClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserSignatureClass")
	return
}

func NewModifyUserSignatureClassResponse() (response *ModifyUserSignatureClassResponse) {
	response = &ModifyUserSignatureClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换Tiga引擎规则类型的生效开关
func (c *Client) ModifyUserSignatureClass(request *ModifyUserSignatureClassRequest) (response *ModifyUserSignatureClassResponse, err error) {
	if request == nil {
		request = NewModifyUserSignatureClassRequest()
	}
	response = NewModifyUserSignatureClassResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiSecAssetTreeRequest() (request *DescribeApiSecAssetTreeRequest) {
	request = &DescribeApiSecAssetTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiSecAssetTree")
	return
}

func NewDescribeApiSecAssetTreeResponse() (response *DescribeApiSecAssetTreeResponse) {
	response = &DescribeApiSecAssetTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api安全资产树查询接口
func (c *Client) DescribeApiSecAssetTree(request *DescribeApiSecAssetTreeRequest) (response *DescribeApiSecAssetTreeResponse, err error) {
	if request == nil {
		request = NewDescribeApiSecAssetTreeRequest()
	}
	response = NewDescribeApiSecAssetTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOwaspRulesRequest() (request *DescribeOwaspRulesRequest) {
	request = &DescribeOwaspRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOwaspRules")
	return
}

func NewDescribeOwaspRulesResponse() (response *DescribeOwaspRulesResponse) {
	response = &DescribeOwaspRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则引擎的规则列表
func (c *Client) DescribeOwaspRules(request *DescribeOwaspRulesRequest) (response *DescribeOwaspRulesResponse, err error) {
	if request == nil {
		request = NewDescribeOwaspRulesRequest()
	}
	response = NewDescribeOwaspRulesResponse()
	err = c.Send(request, response)
	return
}

func NewValidateCELRequest() (request *ValidateCELRequest) {
	request = &ValidateCELRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ValidateCEL")
	return
}

func NewValidateCELResponse() (response *ValidateCELResponse) {
	response = &ValidateCELResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证CEL表达式是否符合语法
func (c *Client) ValidateCEL(request *ValidateCELRequest) (response *ValidateCELResponse, err error) {
	if request == nil {
		request = NewValidateCELRequest()
	}
	response = NewValidateCELResponse()
	err = c.Send(request, response)
	return
}

func NewQueryConstantEnumRequest() (request *QueryConstantEnumRequest) {
	request = &QueryConstantEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryConstantEnum")
	return
}

func NewQueryConstantEnumResponse() (response *QueryConstantEnumResponse) {
	response = &QueryConstantEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [WAF运维管理系统]-查询常量枚举值
func (c *Client) QueryConstantEnum(request *QueryConstantEnumRequest) (response *QueryConstantEnumResponse, err error) {
	if request == nil {
		request = NewQueryConstantEnumRequest()
	}
	response = NewQueryConstantEnumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessLogCountRequest() (request *DescribeAccessLogCountRequest) {
	request = &DescribeAccessLogCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAccessLogCount")
	return
}

func NewDescribeAccessLogCountResponse() (response *DescribeAccessLogCountResponse) {
	response = &DescribeAccessLogCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询满足搜索条件的访问日志数量
func (c *Client) DescribeAccessLogCount(request *DescribeAccessLogCountRequest) (response *DescribeAccessLogCountResponse, err error) {
	if request == nil {
		request = NewDescribeAccessLogCountRequest()
	}
	response = NewDescribeAccessLogCountResponse()
	err = c.Send(request, response)
	return
}

func NewUpsertCCAutoStatusRequest() (request *UpsertCCAutoStatusRequest) {
	request = &UpsertCCAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpsertCCAutoStatus")
	return
}

func NewUpsertCCAutoStatusResponse() (response *UpsertCCAutoStatusResponse) {
	response = &UpsertCCAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf 斯巴达版本更新cc自动封堵状态
func (c *Client) UpsertCCAutoStatus(request *UpsertCCAutoStatusRequest) (response *UpsertCCAutoStatusResponse, err error) {
	if request == nil {
		request = NewUpsertCCAutoStatusRequest()
	}
	response = NewUpsertCCAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpAccessControlV2Request() (request *DeleteIpAccessControlV2Request) {
	request = &DeleteIpAccessControlV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteIpAccessControlV2")
	return
}

func NewDeleteIpAccessControlV2Response() (response *DeleteIpAccessControlV2Response) {
	response = &DeleteIpAccessControlV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单最新版本删除接口
func (c *Client) DeleteIpAccessControlV2(request *DeleteIpAccessControlV2Request) (response *DeleteIpAccessControlV2Response, err error) {
	if request == nil {
		request = NewDeleteIpAccessControlV2Request()
	}
	response = NewDeleteIpAccessControlV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessLogRequest() (request *DescribeBotAccessLogRequest) {
	request = &DescribeBotAccessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAccessLog")
	return
}

func NewDescribeBotAccessLogResponse() (response *DescribeBotAccessLogResponse) {
	response = &DescribeBotAccessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bot访问日志列表接口
func (c *Client) DescribeBotAccessLog(request *DescribeBotAccessLogRequest) (response *DescribeBotAccessLogResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessLogRequest()
	}
	response = NewDescribeBotAccessLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBatchCustomWhiteRuleRequest() (request *ModifyBatchCustomWhiteRuleRequest) {
	request = &ModifyBatchCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBatchCustomWhiteRule")
	return
}

func NewModifyBatchCustomWhiteRuleResponse() (response *ModifyBatchCustomWhiteRuleResponse) {
	response = &ModifyBatchCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改批量精准白名单规则
func (c *Client) ModifyBatchCustomWhiteRule(request *ModifyBatchCustomWhiteRuleRequest) (response *ModifyBatchCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyBatchCustomWhiteRuleRequest()
	}
	response = NewModifyBatchCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrafficMarkingRequest() (request *ModifyTrafficMarkingRequest) {
	request = &ModifyTrafficMarkingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyTrafficMarking")
	return
}

func NewModifyTrafficMarkingResponse() (response *ModifyTrafficMarkingResponse) {
	response = &ModifyTrafficMarkingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新流量标记配置
func (c *Client) ModifyTrafficMarking(request *ModifyTrafficMarkingRequest) (response *ModifyTrafficMarkingResponse, err error) {
	if request == nil {
		request = NewModifyTrafficMarkingRequest()
	}
	response = NewModifyTrafficMarkingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomCaptchaStatusRequest() (request *DescribeCustomCaptchaStatusRequest) {
	request = &DescribeCustomCaptchaStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomCaptchaStatus")
	return
}

func NewDescribeCustomCaptchaStatusResponse() (response *DescribeCustomCaptchaStatusResponse) {
	response = &DescribeCustomCaptchaStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示用户是否接入了验证码服务
func (c *Client) DescribeCustomCaptchaStatus(request *DescribeCustomCaptchaStatusRequest) (response *DescribeCustomCaptchaStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCustomCaptchaStatusRequest()
	}
	response = NewDescribeCustomCaptchaStatusResponse()
	err = c.Send(request, response)
	return
}

func NewPostAccessDownloadTaskRequest() (request *PostAccessDownloadTaskRequest) {
	request = &PostAccessDownloadTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "PostAccessDownloadTask")
	return
}

func NewPostAccessDownloadTaskResponse() (response *PostAccessDownloadTaskResponse) {
	response = &PostAccessDownloadTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建搜索下载访问日志任务
func (c *Client) PostAccessDownloadTask(request *PostAccessDownloadTaskRequest) (response *PostAccessDownloadTaskResponse, err error) {
	if request == nil {
		request = NewPostAccessDownloadTaskRequest()
	}
	response = NewPostAccessDownloadTaskResponse()
	err = c.Send(request, response)
	return
}

func NewReloadConfigDeliveryRequest() (request *ReloadConfigDeliveryRequest) {
	request = &ReloadConfigDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ReloadConfigDelivery")
	return
}

func NewReloadConfigDeliveryResponse() (response *ReloadConfigDeliveryResponse) {
	response = &ReloadConfigDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [WAF运维管理系统]-防护配置重新下发接口
func (c *Client) ReloadConfigDelivery(request *ReloadConfigDeliveryRequest) (response *ReloadConfigDeliveryResponse, err error) {
	if request == nil {
		request = NewReloadConfigDeliveryRequest()
	}
	response = NewReloadConfigDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchAppletElasticModeRequest() (request *SwitchAppletElasticModeRequest) {
	request = &SwitchAppletElasticModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "SwitchAppletElasticMode")
	return
}

func NewSwitchAppletElasticModeResponse() (response *SwitchAppletElasticModeResponse) {
	response = &SwitchAppletElasticModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换小程序弹性QPS的开关
func (c *Client) SwitchAppletElasticMode(request *SwitchAppletElasticModeRequest) (response *SwitchAppletElasticModeResponse, err error) {
	if request == nil {
		request = NewSwitchAppletElasticModeRequest()
	}
	response = NewSwitchAppletElasticModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackDownloadRecordRequest() (request *DescribeAttackDownloadRecordRequest) {
	request = &DescribeAttackDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackDownloadRecord")
	return
}

func NewDescribeAttackDownloadRecordResponse() (response *DescribeAttackDownloadRecordResponse) {
	response = &DescribeAttackDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下载记录
func (c *Client) DescribeAttackDownloadRecord(request *DescribeAttackDownloadRecordRequest) (response *DescribeAttackDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDescribeAttackDownloadRecordRequest()
	}
	response = NewDescribeAttackDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogContextRequest() (request *DescribeLogContextRequest) {
	request = &DescribeLogContextRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLogContext")
	return
}

func NewDescribeLogContextResponse() (response *DescribeLogContextResponse) {
	response = &DescribeLogContextResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于搜索日志上下文附近的内容，详情参考[上下文检索](https://cloud.tencent.com/document/product/614/53248)。
// API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
func (c *Client) DescribeLogContext(request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
	if request == nil {
		request = NewDescribeLogContextRequest()
	}
	response = NewDescribeLogContextResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAreaBanStatusRequest() (request *ModifyAreaBanStatusRequest) {
	request = &ModifyAreaBanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAreaBanStatus")
	return
}

func NewModifyAreaBanStatusResponse() (response *ModifyAreaBanStatusResponse) {
	response = &ModifyAreaBanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护域名的地域封禁状态
func (c *Client) ModifyAreaBanStatus(request *ModifyAreaBanStatusRequest) (response *ModifyAreaBanStatusResponse, err error) {
	if request == nil {
		request = NewModifyAreaBanStatusRequest()
	}
	response = NewModifyAreaBanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDomainParamLimitRequest() (request *QueryDomainParamLimitRequest) {
	request = &QueryDomainParamLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "QueryDomainParamLimit")
	return
}

func NewQueryDomainParamLimitResponse() (response *QueryDomainParamLimitResponse) {
	response = &QueryDomainParamLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [IP封禁]查询域名参数限制值
func (c *Client) QueryDomainParamLimit(request *QueryDomainParamLimitRequest) (response *QueryDomainParamLimitResponse, err error) {
	if request == nil {
		request = NewQueryDomainParamLimitRequest()
	}
	response = NewQueryDomainParamLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWhiteListRequest() (request *DeleteWhiteListRequest) {
	request = &DeleteWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteWhiteList")
	return
}

func NewDeleteWhiteListResponse() (response *DeleteWhiteListResponse) {
	response = &DeleteWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务白名单，目前是前端对抗规则中的白名单
func (c *Client) DeleteWhiteList(request *DeleteWhiteListRequest) (response *DeleteWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteWhiteListRequest()
	}
	response = NewDeleteWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewCopyBotTCBRuleRequest() (request *CopyBotTCBRuleRequest) {
	request = &CopyBotTCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CopyBotTCBRule")
	return
}

func NewCopyBotTCBRuleResponse() (response *CopyBotTCBRuleResponse) {
	response = &CopyBotTCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 TCB策略域名拷贝
func (c *Client) CopyBotTCBRule(request *CopyBotTCBRuleRequest) (response *CopyBotTCBRuleResponse, err error) {
	if request == nil {
		request = NewCopyBotTCBRuleRequest()
	}
	response = NewCopyBotTCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpAccessControlRequest() (request *DeleteIpAccessControlRequest) {
	request = &DeleteIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteIpAccessControl")
	return
}

func NewDeleteIpAccessControlResponse() (response *DeleteIpAccessControlResponse) {
	response = &DeleteIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf IP黑白名单Delete接口
func (c *Client) DeleteIpAccessControl(request *DeleteIpAccessControlRequest) (response *DeleteIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDeleteIpAccessControlRequest()
	}
	response = NewDeleteIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectionModesRequest() (request *DescribeProtectionModesRequest) {
	request = &DescribeProtectionModesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeProtectionModes")
	return
}

func NewDescribeProtectionModesResponse() (response *DescribeProtectionModesResponse) {
	response = &DescribeProtectionModesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Tiga引擎大类规则及其防护模式
func (c *Client) DescribeProtectionModes(request *DescribeProtectionModesRequest) (response *DescribeProtectionModesResponse, err error) {
	if request == nil {
		request = NewDescribeProtectionModesRequest()
	}
	response = NewDescribeProtectionModesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainWhiteRuleRequest() (request *ModifyDomainWhiteRuleRequest) {
	request = &ModifyDomainWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDomainWhiteRule")
	return
}

func NewModifyDomainWhiteRuleResponse() (response *ModifyDomainWhiteRuleResponse) {
	response = &ModifyDomainWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改某一条规则
func (c *Client) ModifyDomainWhiteRule(request *ModifyDomainWhiteRuleRequest) (response *ModifyDomainWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyDomainWhiteRuleRequest()
	}
	response = NewModifyDomainWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLLMContentRuleRequest() (request *DeleteLLMContentRuleRequest) {
	request = &DeleteLLMContentRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteLLMContentRule")
	return
}

func NewDeleteLLMContentRuleResponse() (response *DeleteLLMContentRuleResponse) {
	response = &DeleteLLMContentRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除大模型内容防护规则
func (c *Client) DeleteLLMContentRule(request *DeleteLLMContentRuleRequest) (response *DeleteLLMContentRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLLMContentRuleRequest()
	}
	response = NewDeleteLLMContentRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBatchAreaBanRuleRequest() (request *CreateBatchAreaBanRuleRequest) {
	request = &CreateBatchAreaBanRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateBatchAreaBanRule")
	return
}

func NewCreateBatchAreaBanRuleResponse() (response *CreateBatchAreaBanRuleResponse) {
	response = &CreateBatchAreaBanRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加批量地域封禁规则
func (c *Client) CreateBatchAreaBanRule(request *CreateBatchAreaBanRuleRequest) (response *CreateBatchAreaBanRuleResponse, err error) {
	if request == nil {
		request = NewCreateBatchAreaBanRuleRequest()
	}
	response = NewCreateBatchAreaBanRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有实例的详细信息
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewEditHWRecordRequest() (request *EditHWRecordRequest) {
	request = &EditHWRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EditHWRecord")
	return
}

func NewEditHWRecordResponse() (response *EditHWRecordResponse) {
	response = &EditHWRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [重保护网]-编辑重保护网记录接口
func (c *Client) EditHWRecord(request *EditHWRecordRequest) (response *EditHWRecordResponse, err error) {
	if request == nil {
		request = NewEditHWRecordRequest()
	}
	response = NewEditHWRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchIpAccessControlRequest() (request *DeleteBatchIpAccessControlRequest) {
	request = &DeleteBatchIpAccessControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteBatchIpAccessControl")
	return
}

func NewDeleteBatchIpAccessControlResponse() (response *DeleteBatchIpAccessControlResponse) {
	response = &DeleteBatchIpAccessControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量黑白名单删除接口
func (c *Client) DeleteBatchIpAccessControl(request *DeleteBatchIpAccessControlRequest) (response *DeleteBatchIpAccessControlResponse, err error) {
	if request == nil {
		request = NewDeleteBatchIpAccessControlRequest()
	}
	response = NewDeleteBatchIpAccessControlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAggregateDomainStatRequest() (request *DescribeBotAggregateDomainStatRequest) {
	request = &DescribeBotAggregateDomainStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAggregateDomainStat")
	return
}

func NewDescribeBotAggregateDomainStatResponse() (response *DescribeBotAggregateDomainStatResponse) {
	response = &DescribeBotAggregateDomainStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 域名bot统计
func (c *Client) DescribeBotAggregateDomainStat(request *DescribeBotAggregateDomainStatRequest) (response *DescribeBotAggregateDomainStatResponse, err error) {
	if request == nil {
		request = NewDescribeBotAggregateDomainStatRequest()
	}
	response = NewDescribeBotAggregateDomainStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAttackDetailRequest() (request *DescribeBotAttackDetailRequest) {
	request = &DescribeBotAttackDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAttackDetail")
	return
}

func NewDescribeBotAttackDetailResponse() (response *DescribeBotAttackDetailResponse) {
	response = &DescribeBotAttackDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 540新增查询bot攻击详情列表
func (c *Client) DescribeBotAttackDetail(request *DescribeBotAttackDetailRequest) (response *DescribeBotAttackDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBotAttackDetailRequest()
	}
	response = NewDescribeBotAttackDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogConfigRequest() (request *DescribeLogConfigRequest) {
	request = &DescribeLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLogConfig")
	return
}

func NewDescribeLogConfigResponse() (response *DescribeLogConfigResponse) {
	response = &DescribeLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志配置接口
func (c *Client) DescribeLogConfig(request *DescribeLogConfigRequest) (response *DescribeLogConfigResponse, err error) {
	if request == nil {
		request = NewDescribeLogConfigRequest()
	}
	response = NewDescribeLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAppletInfoRequest() (request *DescribeUserAppletInfoRequest) {
	request = &DescribeUserAppletInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeUserAppletInfo")
	return
}

func NewDescribeUserAppletInfoResponse() (response *DescribeUserAppletInfoResponse) {
	response = &DescribeUserAppletInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看用户级别小程序配置
func (c *Client) DescribeUserAppletInfo(request *DescribeUserAppletInfoRequest) (response *DescribeUserAppletInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserAppletInfoRequest()
	}
	response = NewDescribeUserAppletInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHWInfoRequest() (request *DescribeHWInfoRequest) {
	request = &DescribeHWInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHWInfo")
	return
}

func NewDescribeHWInfoResponse() (response *DescribeHWInfoResponse) {
	response = &DescribeHWInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [重保护网]-获取重保护网配置信息
func (c *Client) DescribeHWInfo(request *DescribeHWInfoRequest) (response *DescribeHWInfoResponse, err error) {
	if request == nil {
		request = NewDescribeHWInfoRequest()
	}
	response = NewDescribeHWInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceEventsRequest() (request *DescribeInstanceEventsRequest) {
	request = &DescribeInstanceEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeInstanceEvents")
	return
}

func NewDescribeInstanceEventsResponse() (response *DescribeInstanceEventsResponse) {
	response = &DescribeInstanceEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户所有实例事件的详细信息
func (c *Client) DescribeInstanceEvents(request *DescribeInstanceEventsRequest) (response *DescribeInstanceEventsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceEventsRequest()
	}
	response = NewDescribeInstanceEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDefaultProtectConfigRequest() (request *ModifyDefaultProtectConfigRequest) {
	request = &ModifyDefaultProtectConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyDefaultProtectConfig")
	return
}

func NewModifyDefaultProtectConfigResponse() (response *ModifyDefaultProtectConfigResponse) {
	response = &ModifyDefaultProtectConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户默认规则引擎类型
func (c *Client) ModifyDefaultProtectConfig(request *ModifyDefaultProtectConfigRequest) (response *ModifyDefaultProtectConfigResponse, err error) {
	if request == nil {
		request = NewModifyDefaultProtectConfigRequest()
	}
	response = NewModifyDefaultProtectConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotSceneUAStatusRequest() (request *DescribeBotSceneUAStatusRequest) {
	request = &DescribeBotSceneUAStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotSceneUAStatus")
	return
}

func NewDescribeBotSceneUAStatusResponse() (response *DescribeBotSceneUAStatusResponse) {
	response = &DescribeBotSceneUAStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Bot场景的UA策略开关
func (c *Client) DescribeBotSceneUAStatus(request *DescribeBotSceneUAStatusRequest) (response *DescribeBotSceneUAStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBotSceneUAStatusRequest()
	}
	response = NewDescribeBotSceneUAStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProtectGroupRequest() (request *DeleteProtectGroupRequest) {
	request = &DeleteProtectGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteProtectGroup")
	return
}

func NewDeleteProtectGroupResponse() (response *DeleteProtectGroupResponse) {
	response = &DeleteProtectGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防护对象组
func (c *Client) DeleteProtectGroup(request *DeleteProtectGroupRequest) (response *DeleteProtectGroupResponse, err error) {
	if request == nil {
		request = NewDeleteProtectGroupRequest()
	}
	response = NewDeleteProtectGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWafThreatenIntelligenceRequest() (request *DescribeWafThreatenIntelligenceRequest) {
	request = &DescribeWafThreatenIntelligenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeWafThreatenIntelligence")
	return
}

func NewDescribeWafThreatenIntelligenceResponse() (response *DescribeWafThreatenIntelligenceResponse) {
	response = &DescribeWafThreatenIntelligenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述WAF威胁情报封禁模块配置详情
func (c *Client) DescribeWafThreatenIntelligence(request *DescribeWafThreatenIntelligenceRequest) (response *DescribeWafThreatenIntelligenceResponse, err error) {
	if request == nil {
		request = NewDescribeWafThreatenIntelligenceRequest()
	}
	response = NewDescribeWafThreatenIntelligenceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCopyCustomRuleRequest() (request *CreateCopyCustomRuleRequest) {
	request = &CreateCopyCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateCopyCustomRule")
	return
}

func NewCreateCopyCustomRuleResponse() (response *CreateCopyCustomRuleResponse) {
	response = &CreateCopyCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拷贝自定义规则到其他域名
func (c *Client) CreateCopyCustomRule(request *CreateCopyCustomRuleRequest) (response *CreateCopyCustomRuleResponse, err error) {
	if request == nil {
		request = NewCreateCopyCustomRuleRequest()
	}
	response = NewCreateCopyCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceLimitRulesRequest() (request *DeleteInstanceLimitRulesRequest) {
	request = &DeleteInstanceLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteInstanceLimitRules")
	return
}

func NewDeleteInstanceLimitRulesResponse() (response *DeleteInstanceLimitRulesResponse) {
	response = &DeleteInstanceLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除实例限流规则信息
func (c *Client) DeleteInstanceLimitRules(request *DeleteInstanceLimitRulesRequest) (response *DeleteInstanceLimitRulesResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceLimitRulesRequest()
	}
	response = NewDeleteInstanceLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserSignatureRuleV2Request() (request *ModifyUserSignatureRuleV2Request) {
	request = &ModifyUserSignatureRuleV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyUserSignatureRuleV2")
	return
}

func NewModifyUserSignatureRuleV2Response() (response *ModifyUserSignatureRuleV2Response) {
	response = &ModifyUserSignatureRuleV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户防护规则，开启关闭具体的某条规则
func (c *Client) ModifyUserSignatureRuleV2(request *ModifyUserSignatureRuleV2Request) (response *ModifyUserSignatureRuleV2Response, err error) {
	if request == nil {
		request = NewModifyUserSignatureRuleV2Request()
	}
	response = NewModifyUserSignatureRuleV2Response()
	err = c.Send(request, response)
	return
}

func NewUpsertSessionRequest() (request *UpsertSessionRequest) {
	request = &UpsertSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "UpsertSession")
	return
}

func NewUpsertSessionResponse() (response *UpsertSessionResponse) {
	response = &UpsertSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Waf  会话定义 Upsert接口
func (c *Client) UpsertSession(request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
	if request == nil {
		request = NewUpsertSessionRequest()
	}
	response = NewUpsertSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDowngradeQualificationRequest() (request *DescribeDowngradeQualificationRequest) {
	request = &DescribeDowngradeQualificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDowngradeQualification")
	return
}

func NewDescribeDowngradeQualificationResponse() (response *DescribeDowngradeQualificationResponse) {
	response = &DescribeDowngradeQualificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取降配资格
func (c *Client) DescribeDowngradeQualification(request *DescribeDowngradeQualificationRequest) (response *DescribeDowngradeQualificationResponse, err error) {
	if request == nil {
		request = NewDescribeDowngradeQualificationRequest()
	}
	response = NewDescribeDowngradeQualificationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRuleLimitRequest() (request *DescribeRuleLimitRequest) {
	request = &DescribeRuleLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeRuleLimit")
	return
}

func NewDescribeRuleLimitResponse() (response *DescribeRuleLimitResponse) {
	response = &DescribeRuleLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各个模块具体的规格限制
func (c *Client) DescribeRuleLimit(request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
	if request == nil {
		request = NewDescribeRuleLimitRequest()
	}
	response = NewDescribeRuleLimitResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessExportRequest() (request *CreateAccessExportRequest) {
	request = &CreateAccessExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateAccessExport")
	return
}

func NewCreateAccessExportResponse() (response *CreateAccessExportResponse) {
	response = &CreateAccessExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建访问日志导出
func (c *Client) CreateAccessExport(request *CreateAccessExportRequest) (response *CreateAccessExportResponse, err error) {
	if request == nil {
		request = NewCreateAccessExportRequest()
	}
	response = NewCreateAccessExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
	request = &DescribeHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeHosts")
	return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
	response = &DescribeHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf中获取防护域名列表
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
	if request == nil {
		request = NewDescribeHostsRequest()
	}
	response = NewDescribeHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeListenerListRequest() (request *DescribeListenerListRequest) {
	request = &DescribeListenerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeListenerList")
	return
}

func NewDescribeListenerListResponse() (response *DescribeListenerListResponse) {
	response = &DescribeListenerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监听器列表
func (c *Client) DescribeListenerList(request *DescribeListenerListRequest) (response *DescribeListenerListResponse, err error) {
	if request == nil {
		request = NewDescribeListenerListRequest()
	}
	response = NewDescribeListenerListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRewriteCountRequest() (request *ModifyRewriteCountRequest) {
	request = &ModifyRewriteCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyRewriteCount")
	return
}

func NewModifyRewriteCountResponse() (response *ModifyRewriteCountResponse) {
	response = &ModifyRewriteCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志清空次数
func (c *Client) ModifyRewriteCount(request *ModifyRewriteCountRequest) (response *ModifyRewriteCountResponse, err error) {
	if request == nil {
		request = NewModifyRewriteCountRequest()
	}
	response = NewModifyRewriteCountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDomainWhiteRulesRequest() (request *DeleteDomainWhiteRulesRequest) {
	request = &DeleteDomainWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteDomainWhiteRules")
	return
}

func NewDeleteDomainWhiteRulesResponse() (response *DeleteDomainWhiteRulesResponse) {
	response = &DeleteDomainWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除域名规则白名单
func (c *Client) DeleteDomainWhiteRules(request *DeleteDomainWhiteRulesRequest) (response *DeleteDomainWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDeleteDomainWhiteRulesRequest()
	}
	response = NewDeleteDomainWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchCustomWhiteRulesRequest() (request *DescribeBatchCustomWhiteRulesRequest) {
	request = &DescribeBatchCustomWhiteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBatchCustomWhiteRules")
	return
}

func NewDescribeBatchCustomWhiteRulesResponse() (response *DescribeBatchCustomWhiteRulesResponse) {
	response = &DescribeBatchCustomWhiteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取批量精准白名单的规则列表
func (c *Client) DescribeBatchCustomWhiteRules(request *DescribeBatchCustomWhiteRulesRequest) (response *DescribeBatchCustomWhiteRulesResponse, err error) {
	if request == nil {
		request = NewDescribeBatchCustomWhiteRulesRequest()
	}
	response = NewDescribeBatchCustomWhiteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTlsVersionRequest() (request *DescribeTlsVersionRequest) {
	request = &DescribeTlsVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTlsVersion")
	return
}

func NewDescribeTlsVersionResponse() (response *DescribeTlsVersionResponse) {
	response = &DescribeTlsVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户TLS版本
func (c *Client) DescribeTlsVersion(request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
	if request == nil {
		request = NewDescribeTlsVersionRequest()
	}
	response = NewDescribeTlsVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOwaspRuleTypesRequest() (request *DescribeOwaspRuleTypesRequest) {
	request = &DescribeOwaspRuleTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOwaspRuleTypes")
	return
}

func NewDescribeOwaspRuleTypesResponse() (response *DescribeOwaspRuleTypesResponse) {
	response = &DescribeOwaspRuleTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则引擎的规则类型列表
func (c *Client) DescribeOwaspRuleTypes(request *DescribeOwaspRuleTypesRequest) (response *DescribeOwaspRuleTypesResponse, err error) {
	if request == nil {
		request = NewDescribeOwaspRuleTypesRequest()
	}
	response = NewDescribeOwaspRuleTypesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppletRiskStatusRequest() (request *ModifyAppletRiskStatusRequest) {
	request = &ModifyAppletRiskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyAppletRiskStatus")
	return
}

func NewModifyAppletRiskStatusResponse() (response *ModifyAppletRiskStatusResponse) {
	response = &ModifyAppletRiskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 小程序风控开关接口
func (c *Client) ModifyAppletRiskStatus(request *ModifyAppletRiskStatusRequest) (response *ModifyAppletRiskStatusResponse, err error) {
	if request == nil {
		request = NewModifyAppletRiskStatusRequest()
	}
	response = NewModifyAppletRiskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotStatusRequest() (request *ModifyBotStatusRequest) {
	request = &ModifyBotStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotStatus")
	return
}

func NewModifyBotStatusResponse() (response *ModifyBotStatusResponse) {
	response = &ModifyBotStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Bot_V2 bot总开关更新
func (c *Client) ModifyBotStatus(request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
	if request == nil {
		request = NewModifyBotStatusRequest()
	}
	response = NewModifyBotStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRuleUpgradeRequest() (request *RuleUpgradeRequest) {
	request = &RuleUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "RuleUpgrade")
	return
}

func NewRuleUpgradeResponse() (response *RuleUpgradeResponse) {
	response = &RuleUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动tiga规则库升级任务
func (c *Client) RuleUpgrade(request *RuleUpgradeRequest) (response *RuleUpgradeResponse, err error) {
	if request == nil {
		request = NewRuleUpgradeRequest()
	}
	response = NewRuleUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomRuleRequest() (request *DeleteCustomRuleRequest) {
	request = &DeleteCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteCustomRule")
	return
}

func NewDeleteCustomRuleResponse() (response *DeleteCustomRuleResponse) {
	response = &DeleteCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义规则
func (c *Client) DeleteCustomRule(request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCustomRuleRequest()
	}
	response = NewDeleteCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddSearchLogTemplateRequest() (request *AddSearchLogTemplateRequest) {
	request = &AddSearchLogTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddSearchLogTemplate")
	return
}

func NewAddSearchLogTemplateResponse() (response *AddSearchLogTemplateResponse) {
	response = &AddSearchLogTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加日志搜索模板
func (c *Client) AddSearchLogTemplate(request *AddSearchLogTemplateRequest) (response *AddSearchLogTemplateResponse, err error) {
	if request == nil {
		request = NewAddSearchLogTemplateRequest()
	}
	response = NewAddSearchLogTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterInstallPackageRequest() (request *CreateClusterInstallPackageRequest) {
	request = &CreateClusterInstallPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CreateClusterInstallPackage")
	return
}

func NewCreateClusterInstallPackageResponse() (response *CreateClusterInstallPackageResponse) {
	response = &CreateClusterInstallPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建混合云集群安装包
func (c *Client) CreateClusterInstallPackage(request *CreateClusterInstallPackageRequest) (response *CreateClusterInstallPackageResponse, err error) {
	if request == nil {
		request = NewCreateClusterInstallPackageRequest()
	}
	response = NewCreateClusterInstallPackageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomWhiteRuleStatusRequest() (request *ModifyCustomWhiteRuleStatusRequest) {
	request = &ModifyCustomWhiteRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCustomWhiteRuleStatus")
	return
}

func NewModifyCustomWhiteRuleStatusResponse() (response *ModifyCustomWhiteRuleStatusResponse) {
	response = &ModifyCustomWhiteRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或禁用精准白名单
func (c *Client) ModifyCustomWhiteRuleStatus(request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyCustomWhiteRuleStatusRequest()
	}
	response = NewModifyCustomWhiteRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRuleRequest() (request *DescribeApiRuleRequest) {
	request = &DescribeApiRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeApiRule")
	return
}

func NewDescribeApiRuleResponse() (response *DescribeApiRuleResponse) {
	response = &DescribeApiRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户查询一条api规则
func (c *Client) DescribeApiRule(request *DescribeApiRuleRequest) (response *DescribeApiRuleResponse, err error) {
	if request == nil {
		request = NewDescribeApiRuleRequest()
	}
	response = NewDescribeApiRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafCreateBotResourceAfterPayRequest() (request *WafCreateBotResourceAfterPayRequest) {
	request = &WafCreateBotResourceAfterPayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafCreateBotResourceAfterPay")
	return
}

func NewWafCreateBotResourceAfterPayResponse() (response *WafCreateBotResourceAfterPayResponse) {
	response = &WafCreateBotResourceAfterPayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 购买WAF-BOT
func (c *Client) WafCreateBotResourceAfterPay(request *WafCreateBotResourceAfterPayRequest) (response *WafCreateBotResourceAfterPayResponse, err error) {
	if request == nil {
		request = NewWafCreateBotResourceAfterPayRequest()
	}
	response = NewWafCreateBotResourceAfterPayResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopStraceIPRequest() (request *GetTopStraceIPRequest) {
	request = &GetTopStraceIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetTopStraceIP")
	return
}

func NewGetTopStraceIPResponse() (response *GetTopStraceIPResponse) {
	response = &GetTopStraceIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护IP的TOP10
func (c *Client) GetTopStraceIP(request *GetTopStraceIPRequest) (response *GetTopStraceIPResponse, err error) {
	if request == nil {
		request = NewGetTopStraceIPRequest()
	}
	response = NewGetTopStraceIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePayloadAnalysisRequest() (request *DescribePayloadAnalysisRequest) {
	request = &DescribePayloadAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribePayloadAnalysis")
	return
}

func NewDescribePayloadAnalysisResponse() (response *DescribePayloadAnalysisResponse) {
	response = &DescribePayloadAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 攻击日志大模型分析
func (c *Client) DescribePayloadAnalysis(request *DescribePayloadAnalysisRequest) (response *DescribePayloadAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribePayloadAnalysisRequest()
	}
	response = NewDescribePayloadAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewImportApiRulesRequest() (request *ImportApiRulesRequest) {
	request = &ImportApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ImportApiRules")
	return
}

func NewImportApiRulesResponse() (response *ImportApiRulesResponse) {
	response = &ImportApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户手动导入api规则
func (c *Client) ImportApiRules(request *ImportApiRulesRequest) (response *ImportApiRulesResponse, err error) {
	if request == nil {
		request = NewImportApiRulesRequest()
	}
	response = NewImportApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewGetStraceCountRequest() (request *GetStraceCountRequest) {
	request = &GetStraceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetStraceCount")
	return
}

func NewGetStraceCountResponse() (response *GetStraceCountResponse) {
	response = &GetStraceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护总量
func (c *Client) GetStraceCount(request *GetStraceCountRequest) (response *GetStraceCountResponse, err error) {
	if request == nil {
		request = NewGetStraceCountRequest()
	}
	response = NewGetStraceCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBotSceneUCBRuleRequest() (request *ModifyBotSceneUCBRuleRequest) {
	request = &ModifyBotSceneUCBRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyBotSceneUCBRule")
	return
}

func NewModifyBotSceneUCBRuleResponse() (response *ModifyBotSceneUCBRuleResponse) {
	response = &ModifyBotSceneUCBRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【接口复用】场景化后更新Bot的UCB自定义规则，两个调用位置：1.BOT全局白名单 2.BOT场景配置
func (c *Client) ModifyBotSceneUCBRule(request *ModifyBotSceneUCBRuleRequest) (response *ModifyBotSceneUCBRuleResponse, err error) {
	if request == nil {
		request = NewModifyBotSceneUCBRuleRequest()
	}
	response = NewModifyBotSceneUCBRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectionStatusRequest() (request *ModifyProtectionStatusRequest) {
	request = &ModifyProtectionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyProtectionStatus")
	return
}

func NewModifyProtectionStatusResponse() (response *ModifyProtectionStatusResponse) {
	response = &ModifyProtectionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-waf开关
func (c *Client) ModifyProtectionStatus(request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
	if request == nil {
		request = NewModifyProtectionStatusRequest()
	}
	response = NewModifyProtectionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttackDownloadRecordRequest() (request *DeleteAttackDownloadRecordRequest) {
	request = &DeleteAttackDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteAttackDownloadRecord")
	return
}

func NewDeleteAttackDownloadRecordResponse() (response *DeleteAttackDownloadRecordResponse) {
	response = &DeleteAttackDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除攻击日志下载任务记录
func (c *Client) DeleteAttackDownloadRecord(request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
	if request == nil {
		request = NewDeleteAttackDownloadRecordRequest()
	}
	response = NewDeleteAttackDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteObjectsRequest() (request *DeleteObjectsRequest) {
	request = &DeleteObjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteObjects")
	return
}

func NewDeleteObjectsResponse() (response *DeleteObjectsResponse) {
	response = &DeleteObjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除对象接入的对象
func (c *Client) DeleteObjects(request *DeleteObjectsRequest) (response *DeleteObjectsResponse, err error) {
	if request == nil {
		request = NewDeleteObjectsRequest()
	}
	response = NewDeleteObjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLLMContentRuleListRequest() (request *DescribeLLMContentRuleListRequest) {
	request = &DescribeLLMContentRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeLLMContentRuleList")
	return
}

func NewDescribeLLMContentRuleListResponse() (response *DescribeLLMContentRuleListResponse) {
	response = &DescribeLLMContentRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取LLM大模型内容防护规则
func (c *Client) DescribeLLMContentRuleList(request *DescribeLLMContentRuleListRequest) (response *DescribeLLMContentRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeLLMContentRuleListRequest()
	}
	response = NewDescribeLLMContentRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignatureRuleRequest() (request *DescribeOpUserSignatureRuleRequest) {
	request = &DescribeOpUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpUserSignatureRule")
	return
}

func NewDescribeOpUserSignatureRuleResponse() (response *DescribeOpUserSignatureRuleResponse) {
	response = &DescribeOpUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为运营后台获取用户规则
func (c *Client) DescribeOpUserSignatureRule(request *DescribeOpUserSignatureRuleRequest) (response *DescribeOpUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignatureRuleRequest()
	}
	response = NewDescribeOpUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainBlockPageRequest() (request *DescribeDomainBlockPageRequest) {
	request = &DescribeDomainBlockPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeDomainBlockPage")
	return
}

func NewDescribeDomainBlockPageResponse() (response *DescribeDomainBlockPageResponse) {
	response = &DescribeDomainBlockPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示指定域名的拦截页面信息
func (c *Client) DescribeDomainBlockPage(request *DescribeDomainBlockPageRequest) (response *DescribeDomainBlockPageResponse, err error) {
	if request == nil {
		request = NewDescribeDomainBlockPageRequest()
	}
	response = NewDescribeDomainBlockPageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiRulesRequest() (request *DeleteApiRulesRequest) {
	request = &DeleteApiRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteApiRules")
	return
}

func NewDeleteApiRulesResponse() (response *DeleteApiRulesResponse) {
	response = &DeleteApiRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户删除api规则,支持多条删除
func (c *Client) DeleteApiRules(request *DeleteApiRulesRequest) (response *DeleteApiRulesResponse, err error) {
	if request == nil {
		request = NewDeleteApiRulesRequest()
	}
	response = NewDeleteApiRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProtectGroupDomainRequest() (request *DeleteProtectGroupDomainRequest) {
	request = &DeleteProtectGroupDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteProtectGroupDomain")
	return
}

func NewDeleteProtectGroupDomainResponse() (response *DeleteProtectGroupDomainResponse) {
	response = &DeleteProtectGroupDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除防护对象组中的域名绑定
func (c *Client) DeleteProtectGroupDomain(request *DeleteProtectGroupDomainRequest) (response *DeleteProtectGroupDomainResponse, err error) {
	if request == nil {
		request = NewDeleteProtectGroupDomainRequest()
	}
	response = NewDeleteProtectGroupDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAIVerifyResultRequest() (request *DescribeAIVerifyResultRequest) {
	request = &DescribeAIVerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAIVerifyResult")
	return
}

func NewDescribeAIVerifyResultResponse() (response *DescribeAIVerifyResultResponse) {
	response = &DescribeAIVerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf ai在线验证
func (c *Client) DescribeAIVerifyResult(request *DescribeAIVerifyResultRequest) (response *DescribeAIVerifyResultResponse, err error) {
	if request == nil {
		request = NewDescribeAIVerifyResultRequest()
	}
	response = NewDescribeAIVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackTypeRequest() (request *DescribeAttackTypeRequest) {
	request = &DescribeAttackTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeAttackType")
	return
}

func NewDescribeAttackTypeResponse() (response *DescribeAttackTypeResponse) {
	response = &DescribeAttackTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定域名TOP N攻击类型
func (c *Client) DescribeAttackType(request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
	if request == nil {
		request = NewDescribeAttackTypeRequest()
	}
	response = NewDescribeAttackTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCachePathRequest() (request *ModifyCachePathRequest) {
	request = &ModifyCachePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyCachePath")
	return
}

func NewModifyCachePathResponse() (response *ModifyCachePathResponse) {
	response = &ModifyCachePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf斯巴达-编辑缓存路径
func (c *Client) ModifyCachePath(request *ModifyCachePathRequest) (response *ModifyCachePathResponse, err error) {
	if request == nil {
		request = NewModifyCachePathRequest()
	}
	response = NewModifyCachePathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotEventOverviewRequest() (request *DescribeBotEventOverviewRequest) {
	request = &DescribeBotEventOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotEventOverview")
	return
}

func NewDescribeBotEventOverviewResponse() (response *DescribeBotEventOverviewResponse) {
	response = &DescribeBotEventOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BOT事件管理-拉取BOT事件概览
func (c *Client) DescribeBotEventOverview(request *DescribeBotEventOverviewRequest) (response *DescribeBotEventOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBotEventOverviewRequest()
	}
	response = NewDescribeBotEventOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBotAccessAggregateTopNRequest() (request *DescribeBotAccessAggregateTopNRequest) {
	request = &DescribeBotAccessAggregateTopNRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeBotAccessAggregateTopN")
	return
}

func NewDescribeBotAccessAggregateTopNResponse() (response *DescribeBotAccessAggregateTopNResponse) {
	response = &DescribeBotAccessAggregateTopNResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BOT模块的访问日志聚合topN
func (c *Client) DescribeBotAccessAggregateTopN(request *DescribeBotAccessAggregateTopNRequest) (response *DescribeBotAccessAggregateTopNResponse, err error) {
	if request == nil {
		request = NewDescribeBotAccessAggregateTopNRequest()
	}
	response = NewDescribeBotAccessAggregateTopNResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRateLimitsRequest() (request *DeleteRateLimitsRequest) {
	request = &DeleteRateLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteRateLimits")
	return
}

func NewDeleteRateLimitsResponse() (response *DeleteRateLimitsResponse) {
	response = &DeleteRateLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除限流规则
func (c *Client) DeleteRateLimits(request *DeleteRateLimitsRequest) (response *DeleteRateLimitsResponse, err error) {
	if request == nil {
		request = NewDeleteRateLimitsRequest()
	}
	response = NewDeleteRateLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRegionRequest() (request *ModifyInstanceRegionRequest) {
	request = &ModifyInstanceRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyInstanceRegion")
	return
}

func NewModifyInstanceRegionResponse() (response *ModifyInstanceRegionResponse) {
	response = &ModifyInstanceRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例的地域属性
func (c *Client) ModifyInstanceRegion(request *ModifyInstanceRegionRequest) (response *ModifyInstanceRegionResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRegionRequest()
	}
	response = NewModifyInstanceRegionResponse()
	err = c.Send(request, response)
	return
}

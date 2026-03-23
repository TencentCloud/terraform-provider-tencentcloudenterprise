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

package v20220501

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-05-01"

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

func NewDescribeHealthCheckPolicyBindingsRequest() (request *DescribeHealthCheckPolicyBindingsRequest) {
	request = &DescribeHealthCheckPolicyBindingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckPolicyBindings")
	return
}

func NewDescribeHealthCheckPolicyBindingsResponse() (response *DescribeHealthCheckPolicyBindingsResponse) {
	response = &DescribeHealthCheckPolicyBindingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询健康检测策略绑定关系
func (c *Client) DescribeHealthCheckPolicyBindings(request *DescribeHealthCheckPolicyBindingsRequest) (response *DescribeHealthCheckPolicyBindingsResponse, err error) {
	if request == nil {
		request = NewDescribeHealthCheckPolicyBindingsRequest()
	}
	response = NewDescribeHealthCheckPolicyBindingsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstantInspectJobRequest() (request *CreateInstantInspectJobRequest) {
	request = &CreateInstantInspectJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateInstantInspectJob")
	return
}

func NewCreateInstantInspectJobResponse() (response *CreateInstantInspectJobResponse) {
	response = &CreateInstantInspectJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TKE集群巡检任务
func (c *Client) CreateInstantInspectJob(request *CreateInstantInspectJobRequest) (response *CreateInstantInspectJobResponse, err error) {
	if request == nil {
		request = NewCreateInstantInspectJobRequest()
	}
	response = NewCreateInstantInspectJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMachinesRequest() (request *DescribeClusterMachinesRequest) {
	request = &DescribeClusterMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterMachines")
	return
}

func NewDescribeClusterMachinesResponse() (response *DescribeClusterMachinesResponse) {
	response = &DescribeClusterMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询托原生点列表
func (c *Client) DescribeClusterMachines(request *DescribeClusterMachinesRequest) (response *DescribeClusterMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMachinesRequest()
	}
	response = NewDescribeClusterMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterResourceLabelsRequest() (request *DescribeClusterResourceLabelsRequest) {
	request = &DescribeClusterResourceLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterResourceLabels")
	return
}

func NewDescribeClusterResourceLabelsResponse() (response *DescribeClusterResourceLabelsResponse) {
	response = &DescribeClusterResourceLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群资源标签列表（仅内部可见）
func (c *Client) DescribeClusterResourceLabels(request *DescribeClusterResourceLabelsRequest) (response *DescribeClusterResourceLabelsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterResourceLabelsRequest()
	}
	response = NewDescribeClusterResourceLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewRebootMachinesRequest() (request *RebootMachinesRequest) {
	request = &RebootMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RebootMachines")
	return
}

func NewRebootMachinesResponse() (response *RebootMachinesResponse) {
	response = &RebootMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启原生节点实例
func (c *Client) RebootMachines(request *RebootMachinesRequest) (response *RebootMachinesResponse, err error) {
	if request == nil {
		request = NewRebootMachinesRequest()
	}
	response = NewRebootMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceConvertNativeNodeRequest() (request *InquirePriceConvertNativeNodeRequest) {
	request = &InquirePriceConvertNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InquirePriceConvertNativeNode")
	return
}

func NewInquirePriceConvertNativeNodeResponse() (response *InquirePriceConvertNativeNodeResponse) {
	response = &InquirePriceConvertNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 存量转换原生节点询价
func (c *Client) InquirePriceConvertNativeNode(request *InquirePriceConvertNativeNodeRequest) (response *InquirePriceConvertNativeNodeResponse, err error) {
	if request == nil {
		request = NewInquirePriceConvertNativeNodeRequest()
	}
	response = NewInquirePriceConvertNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterRenewNativeNodeRequest() (request *SwitchParameterRenewNativeNodeRequest) {
	request = &SwitchParameterRenewNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SwitchParameterRenewNativeNode")
	return
}

func NewSwitchParameterRenewNativeNodeResponse() (response *SwitchParameterRenewNativeNodeResponse) {
	response = &SwitchParameterRenewNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费原生节点池参数转换
func (c *Client) SwitchParameterRenewNativeNode(request *SwitchParameterRenewNativeNodeRequest) (response *SwitchParameterRenewNativeNodeResponse, err error) {
	if request == nil {
		request = NewSwitchParameterRenewNativeNodeRequest()
	}
	response = NewSwitchParameterRenewNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceRenewNativeNodeRequest() (request *InquirePriceRenewNativeNodeRequest) {
	request = &InquirePriceRenewNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InquirePriceRenewNativeNode")
	return
}

func NewInquirePriceRenewNativeNodeResponse() (response *InquirePriceRenewNativeNodeResponse) {
	response = &InquirePriceRenewNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费原生节点询价
func (c *Client) InquirePriceRenewNativeNode(request *InquirePriceRenewNativeNodeRequest) (response *InquirePriceRenewNativeNodeResponse, err error) {
	if request == nil {
		request = NewInquirePriceRenewNativeNodeRequest()
	}
	response = NewInquirePriceRenewNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNodePoolRequest() (request *ModifyNodePoolRequest) {
	request = &ModifyNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePool")
	return
}

func NewModifyNodePoolResponse() (response *ModifyNodePoolResponse) {
	response = &ModifyNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新 TKE 节点池
func (c *Client) ModifyNodePool(request *ModifyNodePoolRequest) (response *ModifyNodePoolResponse, err error) {
	if request == nil {
		request = NewModifyNodePoolRequest()
	}
	response = NewModifyNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHealthCheckPoliciesRequest() (request *DescribeHealthCheckPoliciesRequest) {
	request = &DescribeHealthCheckPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckPolicies")
	return
}

func NewDescribeHealthCheckPoliciesResponse() (response *DescribeHealthCheckPoliciesResponse) {
	response = &DescribeHealthCheckPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询健康检测策略
func (c *Client) DescribeHealthCheckPolicies(request *DescribeHealthCheckPoliciesRequest) (response *DescribeHealthCheckPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeHealthCheckPoliciesRequest()
	}
	response = NewDescribeHealthCheckPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewConvertRegularToNativeNodeRequest() (request *ConvertRegularToNativeNodeRequest) {
	request = &ConvertRegularToNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ConvertRegularToNativeNode")
	return
}

func NewConvertRegularToNativeNodeResponse() (response *ConvertRegularToNativeNodeResponse) {
	response = &ConvertRegularToNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// **普通节点升级为原生节点**
// 将普通节点升级为原生节点，支持两种升级模式，原地模式和重装模式。
//
//   - 原地模式
//     `Reinstall`: 原地模式，这个参数值为 `false`，表示升级的过程中，CVM 的 OS 不会被重装 。
//     `ClusterID`: 集群的 ID。
//     `TargetNodePoolID`: 目标原生节点池 ID。如果为空，表示升级后的节点不会加入到一个原生节点池，如果不为空，则表示升级之后的节点会加入到这个原生节点池。
//     `InstanceIDs`: 需要升级的普通节点列表。
//     >  **重要声明**
//     升级过程中您的业务将保持正常运行（即 Pod 不会删除重建），平台不会对节点使用的操作系统进行重装;
//     该模式下的服务升级仅能体验到部分原生节点增值特性，其中Qos隔离、内存压缩、运维参数维护等功能将使用受限;
//     完成升级后您仍可在 CVM 控制台查看机器详情，为保障信息一致性，建议在TKE平台操作节点;
//     服务升级流程预估5分钟左右，点击确认升级后可刷新页面在状态栏获取进度;
//     升级属于不可逆流程，请谨慎操作;
//
//   - 重装模式
//     `Reinstall`: 重装模式，这个参数值为`true`，表示升级的过程中，CVM 的 OS 会被重装。
//     `ClusterID`: 集群的 ID。
//     `TargetNodePoolID`: 目标原生节点池 ID。如果为空，表示升级后的节点不会加入到一个原生节点池，如果不为空，则表示升级之后的节点会加入到这个原生节点池。
//     `InstanceIDs`: 需要升级为原生节点的普通节点列表。
//
// > **重要声明**
//
//	服务升级需重装节点镜像为原生节点专有镜像，建议您提前封锁并排空节点;
//	升级过程中将按照所选节点池的参数配置对节点初始化;
//	完成升级后您仍可在 CVM 控制台查看机器详情，为保障信息一致性，建议在TKE平台操作节点;
//	服务升级流程预估5分钟左右，点击确认升级后可刷新页面在状态栏获取进度;
func (c *Client) ConvertRegularToNativeNode(request *ConvertRegularToNativeNodeRequest) (response *ConvertRegularToNativeNodeResponse, err error) {
	if request == nil {
		request = NewConvertRegularToNativeNodeRequest()
	}
	response = NewConvertRegularToNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTradeTypesRequest() (request *DescribeTradeTypesRequest) {
	request = &DescribeTradeTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTradeTypes")
	return
}

func NewDescribeTradeTypesResponse() (response *DescribeTradeTypesResponse) {
	response = &DescribeTradeTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Housekeeper对应的计费相关参数
func (c *Client) DescribeTradeTypes(request *DescribeTradeTypesRequest) (response *DescribeTradeTypesResponse, err error) {
	if request == nil {
		request = NewDescribeTradeTypesRequest()
	}
	response = NewDescribeTradeTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodePoolsRequest() (request *DescribeNodePoolsRequest) {
	request = &DescribeNodePoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeNodePools")
	return
}

func NewDescribeNodePoolsResponse() (response *DescribeNodePoolsResponse) {
	response = &DescribeNodePoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 TKE 节点池列表
func (c *Client) DescribeNodePools(request *DescribeNodePoolsRequest) (response *DescribeNodePoolsResponse, err error) {
	if request == nil {
		request = NewDescribeNodePoolsRequest()
	}
	response = NewDescribeNodePoolsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeParamUpdateProcessRequest() (request *DescribeNodeParamUpdateProcessRequest) {
	request = &DescribeNodeParamUpdateProcessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeNodeParamUpdateProcess")
	return
}

func NewDescribeNodeParamUpdateProcessResponse() (response *DescribeNodeParamUpdateProcessResponse) {
	response = &DescribeNodeParamUpdateProcessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 原生节点支持对存量节点的指定参数进行更新，该接口用于查询节点池下存量节点更新某一指定参数的进度
func (c *Client) DescribeNodeParamUpdateProcess(request *DescribeNodeParamUpdateProcessRequest) (response *DescribeNodeParamUpdateProcessResponse, err error) {
	if request == nil {
		request = NewDescribeNodeParamUpdateProcessRequest()
	}
	response = NewDescribeNodeParamUpdateProcessResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusters")
	return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
	if request == nil {
		request = NewDescribeClustersRequest()
	}
	response = NewDescribeClustersResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyQGPURequest() (request *VerifyQGPURequest) {
	request = &VerifyQGPURequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "VerifyQGPU")
	return
}

func NewVerifyQGPUResponse() (response *VerifyQGPUResponse) {
	response = &VerifyQGPUResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于校验入参中的机型和GPU驱动版本，能否正常使用QGPU能力
func (c *Client) VerifyQGPU(request *VerifyQGPURequest) (response *VerifyQGPUResponse, err error) {
	if request == nil {
		request = NewVerifyQGPURequest()
	}
	response = NewVerifyQGPUResponse()
	err = c.Send(request, response)
	return
}

func NewStartMachinesRequest() (request *StartMachinesRequest) {
	request = &StartMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "StartMachines")
	return
}

func NewStartMachinesResponse() (response *StartMachinesResponse) {
	response = &StartMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StartMachines) 用于启动一个或多个原生节点实例。
//
// 只有状态为 Stopped 的实例才可以进行此操作。
// 接口调用成功后，等待一分钟左右，实例会进入 Running 状态。
// 支持批量操作。每次请求批量实例的上限为100。
// 本接口为同步接口，启动实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为 Running，则代表启动实例操作成功。
func (c *Client) StartMachines(request *StartMachinesRequest) (response *StartMachinesResponse, err error) {
	if request == nil {
		request = NewStartMachinesRequest()
	}
	response = NewStartMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewSetMachineLoginRequest() (request *SetMachineLoginRequest) {
	request = &SetMachineLoginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SetMachineLogin")
	return
}

func NewSetMachineLoginResponse() (response *SetMachineLoginResponse) {
	response = &SetMachineLoginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置是否开启节点登录
func (c *Client) SetMachineLogin(request *SetMachineLoginRequest) (response *SetMachineLoginResponse, err error) {
	if request == nil {
		request = NewSetMachineLoginRequest()
	}
	response = NewSetMachineLoginResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachinesKeyIdsRequest() (request *ModifyMachinesKeyIdsRequest) {
	request = &ModifyMachinesKeyIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyMachinesKeyIds")
	return
}

func NewModifyMachinesKeyIdsResponse() (response *ModifyMachinesKeyIdsResponse) {
	response = &ModifyMachinesKeyIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于批量的下发节点ssh公钥
func (c *Client) ModifyMachinesKeyIds(request *ModifyMachinesKeyIdsRequest) (response *ModifyMachinesKeyIdsResponse, err error) {
	if request == nil {
		request = NewModifyMachinesKeyIdsRequest()
	}
	response = NewModifyMachinesKeyIdsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHealthCheckPolicyRequest() (request *ModifyHealthCheckPolicyRequest) {
	request = &ModifyHealthCheckPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyHealthCheckPolicy")
	return
}

func NewModifyHealthCheckPolicyResponse() (response *ModifyHealthCheckPolicyResponse) {
	response = &ModifyHealthCheckPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改健康检测策略
func (c *Client) ModifyHealthCheckPolicy(request *ModifyHealthCheckPolicyRequest) (response *ModifyHealthCheckPolicyResponse, err error) {
	if request == nil {
		request = NewModifyHealthCheckPolicyRequest()
	}
	response = NewModifyHealthCheckPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportedInstanceFamiliesRequest() (request *DescribeSupportedInstanceFamiliesRequest) {
	request = &DescribeSupportedInstanceFamiliesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeSupportedInstanceFamilies")
	return
}

func NewDescribeSupportedInstanceFamiliesResponse() (response *DescribeSupportedInstanceFamiliesResponse) {
	response = &DescribeSupportedInstanceFamiliesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ### 查询原生节点支持的实例族列表
// 目前原生节点支持两种实例类型：Navite 和 NativeCVM
// **Native**: CXM 模式，底层的实例是 CXM （默认值）
// **NativeCVM**: CVM 模式，底层的实例是 CVM
//
// #### CVM 模式
// 对于 CVM 模式，这里会区分内部账号和外部账号，账号是否在白名单内。
//
// ##### 内部账号
// 返回 `{"all"}`, 表示支持的实例族和 CVM 完全一致。
//
// ##### 外部账号
// ###### 在白名单内
// 返回当前原生节点支持的所有 CVM 实例族列表。
//
// ###### 没在白名单
// 返回当前原生节点支持的 HCC 和裸金属实例族列表。
//
// #### CXM 模式
// 返回当前原生节点已经支持的所有 CXM 实例族列表。
func (c *Client) DescribeSupportedInstanceFamilies(request *DescribeSupportedInstanceFamiliesRequest) (response *DescribeSupportedInstanceFamiliesResponse, err error) {
	if request == nil {
		request = NewDescribeSupportedInstanceFamiliesRequest()
	}
	response = NewDescribeSupportedInstanceFamiliesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHealthCheckTemplateRequest() (request *DescribeHealthCheckTemplateRequest) {
	request = &DescribeHealthCheckTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckTemplate")
	return
}

func NewDescribeHealthCheckTemplateResponse() (response *DescribeHealthCheckTemplateResponse) {
	response = &DescribeHealthCheckTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询健康检测策略模板
func (c *Client) DescribeHealthCheckTemplate(request *DescribeHealthCheckTemplateRequest) (response *DescribeHealthCheckTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeHealthCheckTemplateRequest()
	}
	response = NewDescribeHealthCheckTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineConfigurationRequest() (request *DescribeMachineConfigurationRequest) {
	request = &DescribeMachineConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeMachineConfiguration")
	return
}

func NewDescribeMachineConfigurationResponse() (response *DescribeMachineConfigurationResponse) {
	response = &DescribeMachineConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机器配置信息（计费回调，仅供控制台调用）
func (c *Client) DescribeMachineConfiguration(request *DescribeMachineConfigurationRequest) (response *DescribeMachineConfigurationResponse, err error) {
	if request == nil {
		request = NewDescribeMachineConfigurationRequest()
	}
	response = NewDescribeMachineConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterMachinesRequest() (request *DeleteClusterMachinesRequest) {
	request = &DeleteClusterMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterMachines")
	return
}

func NewDeleteClusterMachinesResponse() (response *DeleteClusterMachinesResponse) {
	response = &DeleteClusterMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除原生节点池节点
func (c *Client) DeleteClusterMachines(request *DeleteClusterMachinesRequest) (response *DeleteClusterMachinesResponse, err error) {
	if request == nil {
		request = NewDeleteClusterMachinesRequest()
	}
	response = NewDeleteClusterMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceRefundNativeNodeRequest() (request *InquirePriceRefundNativeNodeRequest) {
	request = &InquirePriceRefundNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InquirePriceRefundNativeNode")
	return
}

func NewInquirePriceRefundNativeNodeResponse() (response *InquirePriceRefundNativeNodeResponse) {
	response = &InquirePriceRefundNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退费原生节点询价
func (c *Client) InquirePriceRefundNativeNode(request *InquirePriceRefundNativeNodeRequest) (response *InquirePriceRefundNativeNodeResponse, err error) {
	if request == nil {
		request = NewInquirePriceRefundNativeNodeRequest()
	}
	response = NewInquirePriceRefundNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHealthCheckPolicyRequest() (request *DeleteHealthCheckPolicyRequest) {
	request = &DeleteHealthCheckPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteHealthCheckPolicy")
	return
}

func NewDeleteHealthCheckPolicyResponse() (response *DeleteHealthCheckPolicyResponse) {
	response = &DeleteHealthCheckPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除健康检测策略
func (c *Client) DeleteHealthCheckPolicy(request *DeleteHealthCheckPolicyRequest) (response *DeleteHealthCheckPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteHealthCheckPolicyRequest()
	}
	response = NewDeleteHealthCheckPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGPUInfoRequest() (request *DescribeGPUInfoRequest) {
	request = &DescribeGPUInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeGPUInfo")
	return
}

func NewDescribeGPUInfoResponse() (response *DescribeGPUInfoResponse) {
	response = &DescribeGPUInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求该接口，会返回所有适配该机型和操作系统组合的gpu驱动版本
func (c *Client) DescribeGPUInfo(request *DescribeGPUInfoRequest) (response *DescribeGPUInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGPUInfoRequest()
	}
	response = NewDescribeGPUInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNodePoolRequest() (request *DeleteNodePoolRequest) {
	request = &DeleteNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteNodePool")
	return
}

func NewDeleteNodePoolResponse() (response *DeleteNodePoolResponse) {
	response = &DeleteNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除 TKE 节点池
func (c *Client) DeleteNodePool(request *DeleteNodePoolRequest) (response *DeleteNodePoolResponse, err error) {
	if request == nil {
		request = NewDeleteNodePoolRequest()
	}
	response = NewDeleteNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewStopMachinesRequest() (request *StopMachinesRequest) {
	request = &StopMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "StopMachines")
	return
}

func NewStopMachinesResponse() (response *StopMachinesResponse) {
	response = &StopMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StopMachines) 用于关闭一个或多个原生节点实例。
//
// 只有状态为 Running 的实例才可以进行此操作。
// 接口调用成功时，实例会进入 Stopping 状态；关闭实例成功时，实例会进入 Stopped 状态。
// 支持强制关闭。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
// 支持批量操作。每次请求批量实例的上限为 100。
// 本接口为同步接口，关闭实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为stopped_with_charging，则代表关闭实例操作成功。
func (c *Client) StopMachines(request *StopMachinesRequest) (response *StopMachinesResponse, err error) {
	if request == nil {
		request = NewStopMachinesRequest()
	}
	response = NewStopMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
	request = &DescribeZoneInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeZoneInstanceConfigInfos")
	return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
	response = &DescribeZoneInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询原生节点机型配置
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneInstanceConfigInfosRequest()
	}
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewListHousekeeperRegionRequest() (request *ListHousekeeperRegionRequest) {
	request = &ListHousekeeperRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListHousekeeperRegion")
	return
}

func NewListHousekeeperRegionResponse() (response *ListHousekeeperRegionResponse) {
	response = &ListHousekeeperRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户可见的地域列表
func (c *Client) ListHousekeeperRegion(request *ListHousekeeperRegionRequest) (response *ListHousekeeperRegionResponse, err error) {
	if request == nil {
		request = NewListHousekeeperRegionRequest()
	}
	response = NewListHousekeeperRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterMachineRequest() (request *ModifyClusterMachineRequest) {
	request = &ModifyClusterMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterMachine")
	return
}

func NewModifyClusterMachineResponse() (response *ModifyClusterMachineResponse) {
	response = &ModifyClusterMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改原生节点
func (c *Client) ModifyClusterMachine(request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
	if request == nil {
		request = NewModifyClusterMachineRequest()
	}
	response = NewModifyClusterMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHealthCheckPolicyRequest() (request *CreateHealthCheckPolicyRequest) {
	request = &CreateHealthCheckPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateHealthCheckPolicy")
	return
}

func NewCreateHealthCheckPolicyResponse() (response *CreateHealthCheckPolicyResponse) {
	response = &CreateHealthCheckPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建健康检测策略
func (c *Client) CreateHealthCheckPolicy(request *CreateHealthCheckPolicyRequest) (response *CreateHealthCheckPolicyResponse, err error) {
	if request == nil {
		request = NewCreateHealthCheckPolicyRequest()
	}
	response = NewCreateHealthCheckPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationEventsRequest() (request *DescribeOperationEventsRequest) {
	request = &DescribeOperationEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeOperationEvents")
	return
}

func NewDescribeOperationEventsResponse() (response *DescribeOperationEventsResponse) {
	response = &DescribeOperationEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询原生节点池运维事件
func (c *Client) DescribeOperationEvents(request *DescribeOperationEventsRequest) (response *DescribeOperationEventsResponse, err error) {
	if request == nil {
		request = NewDescribeOperationEventsRequest()
	}
	response = NewDescribeOperationEventsResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePricePurchaseNativeNodeRequest() (request *InquirePricePurchaseNativeNodeRequest) {
	request = &InquirePricePurchaseNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InquirePricePurchaseNativeNode")
	return
}

func NewInquirePricePurchaseNativeNodeResponse() (response *InquirePricePurchaseNativeNodeResponse) {
	response = &InquirePricePurchaseNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建原生节点询价
func (c *Client) InquirePricePurchaseNativeNode(request *InquirePricePurchaseNativeNodeRequest) (response *InquirePricePurchaseNativeNodeResponse, err error) {
	if request == nil {
		request = NewInquirePricePurchaseNativeNodeRequest()
	}
	response = NewInquirePricePurchaseNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchParameterCreateNativeNodeRequest() (request *SwitchParameterCreateNativeNodeRequest) {
	request = &SwitchParameterCreateNativeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SwitchParameterCreateNativeNode")
	return
}

func NewSwitchParameterCreateNativeNodeResponse() (response *SwitchParameterCreateNativeNodeResponse) {
	response = &SwitchParameterCreateNativeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建原生节点池参数转换
func (c *Client) SwitchParameterCreateNativeNode(request *SwitchParameterCreateNativeNodeRequest) (response *SwitchParameterCreateNativeNodeResponse, err error) {
	if request == nil {
		request = NewSwitchParameterCreateNativeNodeRequest()
	}
	response = NewSwitchParameterCreateNativeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewRestartMachineRequest() (request *RestartMachineRequest) {
	request = &RestartMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RestartMachine")
	return
}

func NewRestartMachineResponse() (response *RestartMachineResponse) {
	response = &RestartMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启 Machine
func (c *Client) RestartMachine(request *RestartMachineRequest) (response *RestartMachineResponse, err error) {
	if request == nil {
		request = NewRestartMachineRequest()
	}
	response = NewRestartMachineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstantInspectTaskRequest() (request *DescribeInstantInspectTaskRequest) {
	request = &DescribeInstantInspectTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeInstantInspectTask")
	return
}

func NewDescribeInstantInspectTaskResponse() (response *DescribeInstantInspectTaskResponse) {
	response = &DescribeInstantInspectTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TKE集群巡检任务结果
func (c *Client) DescribeInstantInspectTask(request *DescribeInstantInspectTaskRequest) (response *DescribeInstantInspectTaskResponse, err error) {
	if request == nil {
		request = NewDescribeInstantInspectTaskRequest()
	}
	response = NewDescribeInstantInspectTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
	request = &DescribeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstances")
	return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
	response = &DescribeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群下节点实例信息
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInstancesRequest()
	}
	response = NewDescribeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewScaleNodePoolRequest() (request *ScaleNodePoolRequest) {
	request = &ScaleNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ScaleNodePool")
	return
}

func NewScaleNodePoolResponse() (response *ScaleNodePoolResponse) {
	response = &ScaleNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置 TKE 节点池期望节点数
func (c *Client) ScaleNodePool(request *ScaleNodePoolRequest) (response *ScaleNodePoolResponse, err error) {
	if request == nil {
		request = NewScaleNodePoolRequest()
	}
	response = NewScaleNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNodePoolRequest() (request *CreateNodePoolRequest) {
	request = &CreateNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateNodePool")
	return
}

func NewCreateNodePoolResponse() (response *CreateNodePoolResponse) {
	response = &CreateNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建 TKE 节点池
func (c *Client) CreateNodePool(request *CreateNodePoolRequest) (response *CreateNodePoolResponse, err error) {
	if request == nil {
		request = NewCreateNodePoolRequest()
	}
	response = NewCreateNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewSetMachineScaleDownProtectionRequest() (request *SetMachineScaleDownProtectionRequest) {
	request = &SetMachineScaleDownProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SetMachineScaleDownProtection")
	return
}

func NewSetMachineScaleDownProtectionResponse() (response *SetMachineScaleDownProtectionResponse) {
	response = &SetMachineScaleDownProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点是否开启缩容保护设置
func (c *Client) SetMachineScaleDownProtection(request *SetMachineScaleDownProtectionRequest) (response *SetMachineScaleDownProtectionResponse, err error) {
	if request == nil {
		request = NewSetMachineScaleDownProtectionRequest()
	}
	response = NewSetMachineScaleDownProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewInquirePriceHousekeeperRequest() (request *InquirePriceHousekeeperRequest) {
	request = &InquirePriceHousekeeperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InquirePriceHousekeeper")
	return
}

func NewInquirePriceHousekeeperResponse() (response *InquirePriceHousekeeperResponse) {
	response = &InquirePriceHousekeeperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例询价
func (c *Client) InquirePriceHousekeeper(request *InquirePriceHousekeeperRequest) (response *InquirePriceHousekeeperResponse, err error) {
	if request == nil {
		request = NewInquirePriceHousekeeperRequest()
	}
	response = NewInquirePriceHousekeeperResponse()
	err = c.Send(request, response)
	return
}

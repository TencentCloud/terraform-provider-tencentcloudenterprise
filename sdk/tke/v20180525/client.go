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

package v20180525

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-05-25"

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

func NewDeleteExternalNodePoolRequest() (request *DeleteExternalNodePoolRequest) {
	request = &DeleteExternalNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteExternalNodePool")
	return
}

func NewDeleteExternalNodePoolResponse() (response *DeleteExternalNodePoolResponse) {
	response = &DeleteExternalNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除第三方节点池
func (c *Client) DeleteExternalNodePool(request *DeleteExternalNodePoolRequest) (response *DeleteExternalNodePoolResponse, err error) {
	if request == nil {
		request = NewDeleteExternalNodePoolRequest()
	}
	response = NewDeleteExternalNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirtualServiceRequest() (request *ModifyVirtualServiceRequest) {
	request = &ModifyVirtualServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyVirtualService")
	return
}

func NewModifyVirtualServiceResponse() (response *ModifyVirtualServiceResponse) {
	response = &ModifyVirtualServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改虚拟服务.
func (c *Client) ModifyVirtualService(request *ModifyVirtualServiceRequest) (response *ModifyVirtualServiceResponse, err error) {
	if request == nil {
		request = NewModifyVirtualServiceRequest()
	}
	response = NewModifyVirtualServiceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterNodePoolFromExistingAsgRequest() (request *CreateClusterNodePoolFromExistingAsgRequest) {
	request = &CreateClusterNodePoolFromExistingAsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterNodePoolFromExistingAsg")
	return
}

func NewCreateClusterNodePoolFromExistingAsgResponse() (response *CreateClusterNodePoolFromExistingAsgResponse) {
	response = &CreateClusterNodePoolFromExistingAsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从伸缩组创建节点池
func (c *Client) CreateClusterNodePoolFromExistingAsg(request *CreateClusterNodePoolFromExistingAsgRequest) (response *CreateClusterNodePoolFromExistingAsgResponse, err error) {
	if request == nil {
		request = NewCreateClusterNodePoolFromExistingAsgRequest()
	}
	response = NewCreateClusterNodePoolFromExistingAsgResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHubClusterRequest() (request *CreateHubClusterRequest) {
	request = &CreateHubClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateHubCluster")
	return
}

func NewCreateHubClusterResponse() (response *CreateHubClusterResponse) {
	response = &CreateHubClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Hub集群
func (c *Client) CreateHubCluster(request *CreateHubClusterRequest) (response *CreateHubClusterResponse, err error) {
	if request == nil {
		request = NewCreateHubClusterRequest()
	}
	response = NewCreateHubClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddonValuesRequest() (request *DescribeAddonValuesRequest) {
	request = &DescribeAddonValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAddonValues")
	return
}

func NewDescribeAddonValuesResponse() (response *DescribeAddonValuesResponse) {
	response = &DescribeAddonValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一个addon的参数
func (c *Client) DescribeAddonValues(request *DescribeAddonValuesRequest) (response *DescribeAddonValuesResponse, err error) {
	if request == nil {
		request = NewDescribeAddonValuesRequest()
	}
	response = NewDescribeAddonValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdSnapshotPreUploadURLRequest() (request *DescribeEtcdSnapshotPreUploadURLRequest) {
	request = &DescribeEtcdSnapshotPreUploadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdSnapshotPreUploadURL")
	return
}

func NewDescribeEtcdSnapshotPreUploadURLResponse() (response *DescribeEtcdSnapshotPreUploadURLResponse) {
	response = &DescribeEtcdSnapshotPreUploadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取etcd快照预上传url
func (c *Client) DescribeEtcdSnapshotPreUploadURL(request *DescribeEtcdSnapshotPreUploadURLRequest) (response *DescribeEtcdSnapshotPreUploadURLResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdSnapshotPreUploadURLRequest()
	}
	response = NewDescribeEtcdSnapshotPreUploadURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAvailableExtraArgsRequest() (request *DescribeClusterAvailableExtraArgsRequest) {
	request = &DescribeClusterAvailableExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAvailableExtraArgs")
	return
}

func NewDescribeClusterAvailableExtraArgsResponse() (response *DescribeClusterAvailableExtraArgsResponse) {
	response = &DescribeClusterAvailableExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群可用的自定义参数
func (c *Client) DescribeClusterAvailableExtraArgs(request *DescribeClusterAvailableExtraArgsRequest) (response *DescribeClusterAvailableExtraArgsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAvailableExtraArgsRequest()
	}
	response = NewDescribeClusterAvailableExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHelmChartDetailRequest() (request *DescribeHelmChartDetailRequest) {
	request = &DescribeHelmChartDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHelmChartDetail")
	return
}

func NewDescribeHelmChartDetailResponse() (response *DescribeHelmChartDetailResponse) {
	response = &DescribeHelmChartDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取chart详情
func (c *Client) DescribeHelmChartDetail(request *DescribeHelmChartDetailRequest) (response *DescribeHelmChartDetailResponse, err error) {
	if request == nil {
		request = NewDescribeHelmChartDetailRequest()
	}
	response = NewDescribeHelmChartDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmPolicyRequest() (request *ModifyAlarmPolicyRequest) {
	request = &ModifyAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyAlarmPolicy")
	return
}

func NewModifyAlarmPolicyResponse() (response *ModifyAlarmPolicyResponse) {
	response = &ModifyAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警策略
func (c *Client) ModifyAlarmPolicy(request *ModifyAlarmPolicyRequest) (response *ModifyAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewModifyAlarmPolicyRequest()
	}
	response = NewModifyAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterExtraArgsRequest() (request *ModifyClusterExtraArgsRequest) {
	request = &ModifyClusterExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterExtraArgs")
	return
}

func NewModifyClusterExtraArgsResponse() (response *ModifyClusterExtraArgsResponse) {
	response = &ModifyClusterExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群自定义参数，只支持托管集群
func (c *Client) ModifyClusterExtraArgs(request *ModifyClusterExtraArgsRequest) (response *ModifyClusterExtraArgsResponse, err error) {
	if request == nil {
		request = NewModifyClusterExtraArgsRequest()
	}
	response = NewModifyClusterExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogConfigsRequest() (request *DeleteLogConfigsRequest) {
	request = &DeleteLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteLogConfigs")
	return
}

func NewDeleteLogConfigsResponse() (response *DeleteLogConfigsResponse) {
	response = &DeleteLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群内采集规则
func (c *Client) DeleteLogConfigs(request *DeleteLogConfigsRequest) (response *DeleteLogConfigsResponse, err error) {
	if request == nil {
		request = NewDeleteLogConfigsRequest()
	}
	response = NewDeleteLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterStatusRequest() (request *DescribeClusterStatusRequest) {
	request = &DescribeClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterStatus")
	return
}

func NewDescribeClusterStatusResponse() (response *DescribeClusterStatusResponse) {
	response = &DescribeClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看集群状态列表
func (c *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterStatusRequest()
	}
	response = NewDescribeClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIndependentClusterRequest() (request *CreateIndependentClusterRequest) {
	request = &CreateIndependentClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateIndependentCluster")
	return
}

func NewCreateIndependentClusterResponse() (response *CreateIndependentClusterResponse) {
	response = &CreateIndependentClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建独立集群
func (c *Client) CreateIndependentCluster(request *CreateIndependentClusterRequest) (response *CreateIndependentClusterResponse, err error) {
	if request == nil {
		request = NewCreateIndependentClusterRequest()
	}
	response = NewCreateIndependentClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRuntimeVersionRequest() (request *DescribeClusterRuntimeVersionRequest) {
	request = &DescribeClusterRuntimeVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRuntimeVersion")
	return
}

func NewDescribeClusterRuntimeVersionResponse() (response *DescribeClusterRuntimeVersionResponse) {
	response = &DescribeClusterRuntimeVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前集群运行时版本，以及升级可选版本信息
func (c *Client) DescribeClusterRuntimeVersion(request *DescribeClusterRuntimeVersionRequest) (response *DescribeClusterRuntimeVersionResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRuntimeVersionRequest()
	}
	response = NewDescribeClusterRuntimeVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
	request = &DescribeSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeSecret")
	return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
	response = &DescribeSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询秘钥
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
	if request == nil {
		request = NewDescribeSecretRequest()
	}
	response = NewDescribeSecretResponse()
	err = c.Send(request, response)
	return
}

func NewGetTkeAppDiffRequest() (request *GetTkeAppDiffRequest) {
	request = &GetTkeAppDiffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetTkeAppDiff")
	return
}

func NewGetTkeAppDiffResponse() (response *GetTkeAppDiffResponse) {
	response = &GetTkeAppDiffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TKEApp的Diff信息
func (c *Client) GetTkeAppDiff(request *GetTkeAppDiffRequest) (response *GetTkeAppDiffResponse, err error) {
	if request == nil {
		request = NewGetTkeAppDiffRequest()
	}
	response = NewGetTkeAppDiffResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterExtraArgsTaskStateRequest() (request *ModifyClusterExtraArgsTaskStateRequest) {
	request = &ModifyClusterExtraArgsTaskStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterExtraArgsTaskState")
	return
}

func NewModifyClusterExtraArgsTaskStateResponse() (response *ModifyClusterExtraArgsTaskStateResponse) {
	response = &ModifyClusterExtraArgsTaskStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 暂停或者取消集群更新参数任务
func (c *Client) ModifyClusterExtraArgsTaskState(request *ModifyClusterExtraArgsTaskStateRequest) (response *ModifyClusterExtraArgsTaskStateResponse, err error) {
	if request == nil {
		request = NewModifyClusterExtraArgsTaskStateRequest()
	}
	response = NewModifyClusterExtraArgsTaskStateResponse()
	err = c.Send(request, response)
	return
}

func NewInstallClustersCostsRequest() (request *InstallClustersCostsRequest) {
	request = &InstallClustersCostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallClustersCosts")
	return
}

func NewInstallClustersCostsResponse() (response *InstallClustersCostsResponse) {
	response = &InstallClustersCostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装集群成本分析服务到用户集群
func (c *Client) InstallClustersCosts(request *InstallClustersCostsRequest) (response *InstallClustersCostsResponse, err error) {
	if request == nil {
		request = NewInstallClustersCostsRequest()
	}
	response = NewInstallClustersCostsResponse()
	err = c.Send(request, response)
	return
}

func NewStopContainerRequest() (request *StopContainerRequest) {
	request = &StopContainerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "StopContainer")
	return
}

func NewStopContainerResponse() (response *StopContainerResponse) {
	response = &StopContainerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止pod内的容器
func (c *Client) StopContainer(request *StopContainerRequest) (response *StopContainerResponse, err error) {
	if request == nil {
		request = NewStopContainerRequest()
	}
	response = NewStopContainerResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGrayScaleStatefulSetGridRequest() (request *UpdateGrayScaleStatefulSetGridRequest) {
	request = &UpdateGrayScaleStatefulSetGridRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateGrayScaleStatefulSetGrid")
	return
}

func NewUpdateGrayScaleStatefulSetGridResponse() (response *UpdateGrayScaleStatefulSetGridResponse) {
	response = &UpdateGrayScaleStatefulSetGridResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 灰度StatefulSetGrid
func (c *Client) UpdateGrayScaleStatefulSetGrid(request *UpdateGrayScaleStatefulSetGridRequest) (response *UpdateGrayScaleStatefulSetGridResponse, err error) {
	if request == nil {
		request = NewUpdateGrayScaleStatefulSetGridRequest()
	}
	response = NewUpdateGrayScaleStatefulSetGridResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpgradeClusterProgressRequest() (request *DescribeUpgradeClusterProgressRequest) {
	request = &DescribeUpgradeClusterProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeUpgradeClusterProgress")
	return
}

func NewDescribeUpgradeClusterProgressResponse() (response *DescribeUpgradeClusterProgressResponse) {
	response = &DescribeUpgradeClusterProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群升级进度
func (c *Client) DescribeUpgradeClusterProgress(request *DescribeUpgradeClusterProgressRequest) (response *DescribeUpgradeClusterProgressResponse, err error) {
	if request == nil {
		request = NewDescribeUpgradeClusterProgressRequest()
	}
	response = NewDescribeUpgradeClusterProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSClusterCommonNamesRequest() (request *DescribeEKSClusterCommonNamesRequest) {
	request = &DescribeEKSClusterCommonNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterCommonNames")
	return
}

func NewDescribeEKSClusterCommonNamesResponse() (response *DescribeEKSClusterCommonNamesResponse) {
	response = &DescribeEKSClusterCommonNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定子账户的对应kube-apiserver的客户端证书中CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为50
func (c *Client) DescribeEKSClusterCommonNames(request *DescribeEKSClusterCommonNamesRequest) (response *DescribeEKSClusterCommonNamesResponse, err error) {
	if request == nil {
		request = NewDescribeEKSClusterCommonNamesRequest()
	}
	response = NewDescribeEKSClusterCommonNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusRecordRuleYamlRequest() (request *ModifyPrometheusRecordRuleYamlRequest) {
	request = &ModifyPrometheusRecordRuleYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusRecordRuleYaml")
	return
}

func NewModifyPrometheusRecordRuleYamlResponse() (response *ModifyPrometheusRecordRuleYamlResponse) {
	response = &ModifyPrometheusRecordRuleYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改聚合规则yaml方式
func (c *Client) ModifyPrometheusRecordRuleYaml(request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusRecordRuleYamlRequest()
	}
	response = NewModifyPrometheusRecordRuleYamlResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUseEksRequest() (request *CheckUseEksRequest) {
	request = &CheckUseEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckUseEks")
	return
}

func NewCheckUseEksResponse() (response *CheckUseEksResponse) {
	response = &CheckUseEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查用户是否使用弹性容器服务
func (c *Client) CheckUseEks(request *CheckUseEksRequest) (response *CheckUseEksResponse, err error) {
	if request == nil {
		request = NewCheckUseEksRequest()
	}
	response = NewCheckUseEksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteECMPodsByIdRequest() (request *DeleteECMPodsByIdRequest) {
	request = &DeleteECMPodsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteECMPodsById")
	return
}

func NewDeleteECMPodsByIdResponse() (response *DeleteECMPodsByIdResponse) {
	response = &DeleteECMPodsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定eksId删除ecm pod
func (c *Client) DeleteECMPodsById(request *DeleteECMPodsByIdRequest) (response *DeleteECMPodsByIdResponse, err error) {
	if request == nil {
		request = NewDeleteECMPodsByIdRequest()
	}
	response = NewDeleteECMPodsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewListEKSK8SVersionRequest() (request *ListEKSK8SVersionRequest) {
	request = &ListEKSK8SVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListEKSK8SVersion")
	return
}

func NewListEKSK8SVersionResponse() (response *ListEKSK8SVersionResponse) {
	response = &ListEKSK8SVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取EKS支持的k8s版本
func (c *Client) ListEKSK8SVersion(request *ListEKSK8SVersionRequest) (response *ListEKSK8SVersionResponse, err error) {
	if request == nil {
		request = NewListEKSK8SVersionRequest()
	}
	response = NewListEKSK8SVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusClusterAgentRequest() (request *DeletePrometheusClusterAgentRequest) {
	request = &DeletePrometheusClusterAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusClusterAgent")
	return
}

func NewDeletePrometheusClusterAgentResponse() (response *DeletePrometheusClusterAgentResponse) {
	response = &DeletePrometheusClusterAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除TMP实例的集群关联
func (c *Client) DeletePrometheusClusterAgent(request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusClusterAgentRequest()
	}
	response = NewDeletePrometheusClusterAgentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusAlertPolicyRequest() (request *ModifyPrometheusAlertPolicyRequest) {
	request = &ModifyPrometheusAlertPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAlertPolicy")
	return
}

func NewModifyPrometheusAlertPolicyResponse() (response *ModifyPrometheusAlertPolicyResponse) {
	response = &ModifyPrometheusAlertPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改2.0实例告警策略
func (c *Client) ModifyPrometheusAlertPolicy(request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusAlertPolicyRequest()
	}
	response = NewModifyPrometheusAlertPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewResizeEtcdInstanceRequest() (request *ResizeEtcdInstanceRequest) {
	request = &ResizeEtcdInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ResizeEtcdInstance")
	return
}

func NewResizeEtcdInstanceResponse() (response *ResizeEtcdInstanceResponse) {
	response = &ResizeEtcdInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容etcd实例
func (c *Client) ResizeEtcdInstance(request *ResizeEtcdInstanceRequest) (response *ResizeEtcdInstanceResponse, err error) {
	if request == nil {
		request = NewResizeEtcdInstanceRequest()
	}
	response = NewResizeEtcdInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterEndpointRequest() (request *CreateClusterEndpointRequest) {
	request = &CreateClusterEndpointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpoint")
	return
}

func NewCreateClusterEndpointResponse() (response *CreateClusterEndpointResponse) {
	response = &CreateClusterEndpointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群访问端口
func (c *Client) CreateClusterEndpoint(request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
	if request == nil {
		request = NewCreateClusterEndpointRequest()
	}
	response = NewCreateClusterEndpointResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusAgentRequest() (request *CreatePrometheusAgentRequest) {
	request = &CreatePrometheusAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAgent")
	return
}

func NewCreatePrometheusAgentResponse() (response *CreatePrometheusAgentResponse) {
	response = &CreatePrometheusAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联集群
func (c *Client) CreatePrometheusAgent(request *CreatePrometheusAgentRequest) (response *CreatePrometheusAgentResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusAgentRequest()
	}
	response = NewCreatePrometheusAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeletionEventsRequest() (request *DescribeDeletionEventsRequest) {
	request = &DescribeDeletionEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeDeletionEvents")
	return
}

func NewDescribeDeletionEventsResponse() (response *DescribeDeletionEventsResponse) {
	response = &DescribeDeletionEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群内删除被拦截事件列表
func (c *Client) DescribeDeletionEvents(request *DescribeDeletionEventsRequest) (response *DescribeDeletionEventsResponse, err error) {
	if request == nil {
		request = NewDescribeDeletionEventsRequest()
	}
	response = NewDescribeDeletionEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDedicatedAPIServersRequest() (request *DescribeDedicatedAPIServersRequest) {
	request = &DescribeDedicatedAPIServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeDedicatedAPIServers")
	return
}

func NewDescribeDedicatedAPIServersResponse() (response *DescribeDedicatedAPIServersResponse) {
	response = &DescribeDedicatedAPIServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看集群专属APIServer列表
func (c *Client) DescribeDedicatedAPIServers(request *DescribeDedicatedAPIServersRequest) (response *DescribeDedicatedAPIServersResponse, err error) {
	if request == nil {
		request = NewDescribeDedicatedAPIServersRequest()
	}
	response = NewDescribeDedicatedAPIServersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteECMPodRequest() (request *DeleteECMPodRequest) {
	request = &DeleteECMPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteECMPod")
	return
}

func NewDeleteECMPodResponse() (response *DeleteECMPodResponse) {
	response = &DeleteECMPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定ECM pod
func (c *Client) DeleteECMPod(request *DeleteECMPodRequest) (response *DeleteECMPodResponse, err error) {
	if request == nil {
		request = NewDeleteECMPodRequest()
	}
	response = NewDeleteECMPodResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEtcdSnapshotPolicyRequest() (request *CreateEtcdSnapshotPolicyRequest) {
	request = &CreateEtcdSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEtcdSnapshotPolicy")
	return
}

func NewCreateEtcdSnapshotPolicyResponse() (response *CreateEtcdSnapshotPolicyResponse) {
	response = &CreateEtcdSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建etcd快照策略
func (c *Client) CreateEtcdSnapshotPolicy(request *CreateEtcdSnapshotPolicyRequest) (response *CreateEtcdSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewCreateEtcdSnapshotPolicyRequest()
	}
	response = NewCreateEtcdSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAddonRequest() (request *UpdateAddonRequest) {
	request = &UpdateAddonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateAddon")
	return
}

func NewUpdateAddonResponse() (response *UpdateAddonResponse) {
	response = &UpdateAddonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新一个addon的参数和版本
func (c *Client) UpdateAddon(request *UpdateAddonRequest) (response *UpdateAddonResponse, err error) {
	if request == nil {
		request = NewUpdateAddonRequest()
	}
	response = NewUpdateAddonResponse()
	err = c.Send(request, response)
	return
}

func NewInstallAddonRequest() (request *InstallAddonRequest) {
	request = &InstallAddonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallAddon")
	return
}

func NewInstallAddonResponse() (response *InstallAddonResponse) {
	response = &InstallAddonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为目标集群安装一个addon
func (c *Client) InstallAddon(request *InstallAddonRequest) (response *InstallAddonResponse, err error) {
	if request == nil {
		request = NewInstallAddonRequest()
	}
	response = NewInstallAddonResponse()
	err = c.Send(request, response)
	return
}

func NewListExpiredClusterAuthRequest() (request *ListExpiredClusterAuthRequest) {
	request = &ListExpiredClusterAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListExpiredClusterAuth")
	return
}

func NewListExpiredClusterAuthResponse() (response *ListExpiredClusterAuthResponse) {
	response = &ListExpiredClusterAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群内过期的权限信息
func (c *Client) ListExpiredClusterAuth(request *ListExpiredClusterAuthRequest) (response *ListExpiredClusterAuthResponse, err error) {
	if request == nil {
		request = NewListExpiredClusterAuthRequest()
	}
	response = NewListExpiredClusterAuthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterNodePoolRequest() (request *ModifyClusterNodePoolRequest) {
	request = &ModifyClusterNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterNodePool")
	return
}

func NewModifyClusterNodePoolResponse() (response *ModifyClusterNodePoolResponse) {
	response = &ModifyClusterNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑节点池
func (c *Client) ModifyClusterNodePool(request *ModifyClusterNodePoolRequest) (response *ModifyClusterNodePoolResponse, err error) {
	if request == nil {
		request = NewModifyClusterNodePoolRequest()
	}
	response = NewModifyClusterNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusInstanceRequest() (request *CreatePrometheusInstanceRequest) {
	request = &CreatePrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusInstance")
	return
}

func NewCreatePrometheusInstanceResponse() (response *CreatePrometheusInstanceResponse) {
	response = &CreatePrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例
func (c *Client) CreatePrometheusInstance(request *CreatePrometheusInstanceRequest) (response *CreatePrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusInstanceRequest()
	}
	response = NewCreatePrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterLevelChangeRecordsRequest() (request *DescribeClusterLevelChangeRecordsRequest) {
	request = &DescribeClusterLevelChangeRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelChangeRecords")
	return
}

func NewDescribeClusterLevelChangeRecordsResponse() (response *DescribeClusterLevelChangeRecordsResponse) {
	response = &DescribeClusterLevelChangeRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群变配记录
func (c *Client) DescribeClusterLevelChangeRecords(request *DescribeClusterLevelChangeRecordsRequest) (response *DescribeClusterLevelChangeRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterLevelChangeRecordsRequest()
	}
	response = NewDescribeClusterLevelChangeRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeECMEKSClustersRequest() (request *DescribeECMEKSClustersRequest) {
	request = &DescribeECMEKSClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeECMEKSClusters")
	return
}

func NewDescribeECMEKSClustersResponse() (response *DescribeECMEKSClustersResponse) {
	response = &DescribeECMEKSClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ECM弹性集群列表
func (c *Client) DescribeECMEKSClusters(request *DescribeECMEKSClustersRequest) (response *DescribeECMEKSClustersResponse, err error) {
	if request == nil {
		request = NewDescribeECMEKSClustersRequest()
	}
	response = NewDescribeECMEKSClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudRunClusterQuotaRequest() (request *DescribeCloudRunClusterQuotaRequest) {
	request = &DescribeCloudRunClusterQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCloudRunClusterQuota")
	return
}

func NewDescribeCloudRunClusterQuotaResponse() (response *DescribeCloudRunClusterQuotaResponse) {
	response = &DescribeCloudRunClusterQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CloudRun集群配额
func (c *Client) DescribeCloudRunClusterQuota(request *DescribeCloudRunClusterQuotaRequest) (response *DescribeCloudRunClusterQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeCloudRunClusterQuotaRequest()
	}
	response = NewDescribeCloudRunClusterQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewRunPrometheusInstanceRequest() (request *RunPrometheusInstanceRequest) {
	request = &RunPrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RunPrometheusInstance")
	return
}

func NewRunPrometheusInstanceResponse() (response *RunPrometheusInstanceResponse) {
	response = &RunPrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化TMP实例，开启集成中心时调用
func (c *Client) RunPrometheusInstance(request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewRunPrometheusInstanceRequest()
	}
	response = NewRunPrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCheckMigrateVmRequest() (request *CheckMigrateVmRequest) {
	request = &CheckMigrateVmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckMigrateVm")
	return
}

func NewCheckMigrateVmResponse() (response *CheckMigrateVmResponse) {
	response = &CheckMigrateVmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查子机路由是否正常
func (c *Client) CheckMigrateVm(request *CheckMigrateVmRequest) (response *CheckMigrateVmResponse, err error) {
	if request == nil {
		request = NewCheckMigrateVmRequest()
	}
	response = NewCheckMigrateVmResponse()
	err = c.Send(request, response)
	return
}

func NewDrainClusterNodeRequest() (request *DrainClusterNodeRequest) {
	request = &DrainClusterNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DrainClusterNode")
	return
}

func NewDrainClusterNodeResponse() (response *DrainClusterNodeResponse) {
	response = &DrainClusterNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐集群中的节点
func (c *Client) DrainClusterNode(request *DrainClusterNodeRequest) (response *DrainClusterNodeResponse, err error) {
	if request == nil {
		request = NewDrainClusterNodeRequest()
	}
	response = NewDrainClusterNodeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterVirtualNodeRequest() (request *CreateClusterVirtualNodeRequest) {
	request = &CreateClusterVirtualNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNode")
	return
}

func NewCreateClusterVirtualNodeResponse() (response *CreateClusterVirtualNodeResponse) {
	response = &CreateClusterVirtualNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建按量计费超级节点
func (c *Client) CreateClusterVirtualNode(request *CreateClusterVirtualNodeRequest) (response *CreateClusterVirtualNodeResponse, err error) {
	if request == nil {
		request = NewCreateClusterVirtualNodeRequest()
	}
	response = NewCreateClusterVirtualNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusRegionsRequest() (request *DescribePrometheusRegionsRequest) {
	request = &DescribePrometheusRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusRegions")
	return
}

func NewDescribePrometheusRegionsResponse() (response *DescribePrometheusRegionsResponse) {
	response = &DescribePrometheusRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TPS当前支持的地域
func (c *Client) DescribePrometheusRegions(request *DescribePrometheusRegionsRequest) (response *DescribePrometheusRegionsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusRegionsRequest()
	}
	response = NewDescribePrometheusRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCLogSwitchesRequest() (request *DescribeTDCCLogSwitchesRequest) {
	request = &DescribeTDCCLogSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCLogSwitches")
	return
}

func NewDescribeTDCCLogSwitchesResponse() (response *DescribeTDCCLogSwitchesResponse) {
	response = &DescribeTDCCLogSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询注册集群是否安装日志组件
func (c *Client) DescribeTDCCLogSwitches(request *DescribeTDCCLogSwitchesRequest) (response *DescribeTDCCLogSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCLogSwitchesRequest()
	}
	response = NewDescribeTDCCLogSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateImageRegistryCredentialRequest() (request *UpdateImageRegistryCredentialRequest) {
	request = &UpdateImageRegistryCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateImageRegistryCredential")
	return
}

func NewUpdateImageRegistryCredentialResponse() (response *UpdateImageRegistryCredentialResponse) {
	response = &UpdateImageRegistryCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像仓库凭证接口
func (c *Client) UpdateImageRegistryCredential(request *UpdateImageRegistryCredentialRequest) (response *UpdateImageRegistryCredentialResponse, err error) {
	if request == nil {
		request = NewUpdateImageRegistryCredentialRequest()
	}
	response = NewUpdateImageRegistryCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewAcquireECMEKSClusterAdminRoleRequest() (request *AcquireECMEKSClusterAdminRoleRequest) {
	request = &AcquireECMEKSClusterAdminRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AcquireECMEKSClusterAdminRole")
	return
}

func NewAcquireECMEKSClusterAdminRoleResponse() (response *AcquireECMEKSClusterAdminRoleResponse) {
	response = &AcquireECMEKSClusterAdminRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过此接口，可以获取EKS集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
func (c *Client) AcquireECMEKSClusterAdminRole(request *AcquireECMEKSClusterAdminRoleRequest) (response *AcquireECMEKSClusterAdminRoleResponse, err error) {
	if request == nil {
		request = NewAcquireECMEKSClusterAdminRoleRequest()
	}
	response = NewAcquireECMEKSClusterAdminRoleResponse()
	err = c.Send(request, response)
	return
}

func NewAdmitResourceOperationRequest() (request *AdmitResourceOperationRequest) {
	request = &AdmitResourceOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AdmitResourceOperation")
	return
}

func NewAdmitResourceOperationResponse() (response *AdmitResourceOperationResponse) {
	response = &AdmitResourceOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源操作许可
func (c *Client) AdmitResourceOperation(request *AdmitResourceOperationRequest) (response *AdmitResourceOperationResponse, err error) {
	if request == nil {
		request = NewAdmitResourceOperationRequest()
	}
	response = NewAdmitResourceOperationResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReservedInstanceNameRequest() (request *ModifyReservedInstanceNameRequest) {
	request = &ModifyReservedInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyReservedInstanceName")
	return
}

func NewModifyReservedInstanceNameResponse() (response *ModifyReservedInstanceNameResponse) {
	response = &ModifyReservedInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改预留券名称
func (c *Client) ModifyReservedInstanceName(request *ModifyReservedInstanceNameRequest) (response *ModifyReservedInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyReservedInstanceNameRequest()
	}
	response = NewModifyReservedInstanceNameResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeEKSClusterAuthorizationModeRequest() (request *UpgradeEKSClusterAuthorizationModeRequest) {
	request = &UpgradeEKSClusterAuthorizationModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeEKSClusterAuthorizationMode")
	return
}

func NewUpgradeEKSClusterAuthorizationModeResponse() (response *UpgradeEKSClusterAuthorizationModeResponse) {
	response = &UpgradeEKSClusterAuthorizationModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级EKS集群授权模式为RBAC，默认的授权模式是子账户拉取kubeconfig都为admin token，且访问集群内Kubernetes资源时，不做鉴权，升级为RBAC之后，可以提升集群安全性及权限的细粒度控制
func (c *Client) UpgradeEKSClusterAuthorizationMode(request *UpgradeEKSClusterAuthorizationModeRequest) (response *UpgradeEKSClusterAuthorizationModeResponse, err error) {
	if request == nil {
		request = NewUpgradeEKSClusterAuthorizationModeRequest()
	}
	response = NewUpgradeEKSClusterAuthorizationModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksLogSwitchesRequest() (request *DescribeEksLogSwitchesRequest) {
	request = &DescribeEksLogSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksLogSwitches")
	return
}

func NewDescribeEksLogSwitchesResponse() (response *DescribeEksLogSwitchesResponse) {
	response = &DescribeEksLogSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Eks集群日志（审计、事件、普通日志）开关列表
func (c *Client) DescribeEksLogSwitches(request *DescribeEksLogSwitchesRequest) (response *DescribeEksLogSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeEksLogSwitchesRequest()
	}
	response = NewDescribeEksLogSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRouteTableRequest() (request *CreateClusterRouteTableRequest) {
	request = &CreateClusterRouteTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRouteTable")
	return
}

func NewCreateClusterRouteTableResponse() (response *CreateClusterRouteTableResponse) {
	response = &CreateClusterRouteTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群路由表
func (c *Client) CreateClusterRouteTable(request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
	if request == nil {
		request = NewCreateClusterRouteTableRequest()
	}
	response = NewCreateClusterRouteTableResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirtualClusterAddonRequest() (request *CreateVirtualClusterAddonRequest) {
	request = &CreateVirtualClusterAddonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateVirtualClusterAddon")
	return
}

func NewCreateVirtualClusterAddonResponse() (response *CreateVirtualClusterAddonResponse) {
	response = &CreateVirtualClusterAddonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建虚拟集群插件
func (c *Client) CreateVirtualClusterAddon(request *CreateVirtualClusterAddonRequest) (response *CreateVirtualClusterAddonResponse, err error) {
	if request == nil {
		request = NewCreateVirtualClusterAddonRequest()
	}
	response = NewCreateVirtualClusterAddonResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTkeEdgeAlarmPoliciesRequest() (request *DeleteTkeEdgeAlarmPoliciesRequest) {
	request = &DeleteTkeEdgeAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteTkeEdgeAlarmPolicies")
	return
}

func NewDeleteTkeEdgeAlarmPoliciesResponse() (response *DeleteTkeEdgeAlarmPoliciesResponse) {
	response = &DeleteTkeEdgeAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除边缘集群告警策略，支持批量删除
func (c *Client) DeleteTkeEdgeAlarmPolicies(request *DeleteTkeEdgeAlarmPoliciesRequest) (response *DeleteTkeEdgeAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteTkeEdgeAlarmPoliciesRequest()
	}
	response = NewDeleteTkeEdgeAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeDashboardIDRequest() (request *DescribeEdgeDashboardIDRequest) {
	request = &DescribeEdgeDashboardIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeDashboardID")
	return
}

func NewDescribeEdgeDashboardIDResponse() (response *DescribeEdgeDashboardIDResponse) {
	response = &DescribeEdgeDashboardIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据集群、仪表盘类型查询dashboardID信息
func (c *Client) DescribeEdgeDashboardID(request *DescribeEdgeDashboardIDRequest) (response *DescribeEdgeDashboardIDResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeDashboardIDRequest()
	}
	response = NewDescribeEdgeDashboardIDResponse()
	err = c.Send(request, response)
	return
}

func NewForwardRequestTDCCRequest() (request *ForwardRequestTDCCRequest) {
	request = &ForwardRequestTDCCRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardRequestTDCC")
	return
}

func NewForwardRequestTDCCResponse() (response *ForwardRequestTDCCResponse) {
	response = &ForwardRequestTDCCResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 转发注册集群 workload 查询
func (c *Client) ForwardRequestTDCC(request *ForwardRequestTDCCRequest) (response *ForwardRequestTDCCResponse, err error) {
	if request == nil {
		request = NewForwardRequestTDCCRequest()
	}
	response = NewForwardRequestTDCCResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterVersionRequest() (request *UpdateClusterVersionRequest) {
	request = &UpdateClusterVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterVersion")
	return
}

func NewUpdateClusterVersionResponse() (response *UpdateClusterVersionResponse) {
	response = &UpdateClusterVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级集群 Master 组件到指定版本
func (c *Client) UpdateClusterVersion(request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
	if request == nil {
		request = NewUpdateClusterVersionRequest()
	}
	response = NewUpdateClusterVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInspectionResultsOverviewRequest() (request *DescribeClusterInspectionResultsOverviewRequest) {
	request = &DescribeClusterInspectionResultsOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspectionResultsOverview")
	return
}

func NewDescribeClusterInspectionResultsOverviewResponse() (response *DescribeClusterInspectionResultsOverviewResponse) {
	response = &DescribeClusterInspectionResultsOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户单个Region下的所有集群巡检结果概览信息
func (c *Client) DescribeClusterInspectionResultsOverview(request *DescribeClusterInspectionResultsOverviewRequest) (response *DescribeClusterInspectionResultsOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInspectionResultsOverviewRequest()
	}
	response = NewDescribeClusterInspectionResultsOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterLevelAttributeRequest() (request *DescribeClusterLevelAttributeRequest) {
	request = &DescribeClusterLevelAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelAttribute")
	return
}

func NewDescribeClusterLevelAttributeResponse() (response *DescribeClusterLevelAttributeResponse) {
	response = &DescribeClusterLevelAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群规模
func (c *Client) DescribeClusterLevelAttribute(request *DescribeClusterLevelAttributeRequest) (response *DescribeClusterLevelAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeClusterLevelAttributeRequest()
	}
	response = NewDescribeClusterLevelAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdInstancesRequest() (request *DescribeEtcdInstancesRequest) {
	request = &DescribeEtcdInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdInstances")
	return
}

func NewDescribeEtcdInstancesResponse() (response *DescribeEtcdInstancesResponse) {
	response = &DescribeEtcdInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询etcd实例列表
func (c *Client) DescribeEtcdInstances(request *DescribeEtcdInstancesRequest) (response *DescribeEtcdInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdInstancesRequest()
	}
	response = NewDescribeEtcdInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodChargeInfoRequest() (request *GetPodChargeInfoRequest) {
	request = &GetPodChargeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodChargeInfo")
	return
}

func NewGetPodChargeInfoResponse() (response *GetPodChargeInfoResponse) {
	response = &GetPodChargeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Pod的计费信息查询
func (c *Client) GetPodChargeInfo(request *GetPodChargeInfoRequest) (response *GetPodChargeInfoResponse, err error) {
	if request == nil {
		request = NewGetPodChargeInfoRequest()
	}
	response = NewGetPodChargeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableInstanceConfigInfoRequest() (request *DescribeAvailableInstanceConfigInfoRequest) {
	request = &DescribeAvailableInstanceConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableInstanceConfigInfo")
	return
}

func NewDescribeAvailableInstanceConfigInfoResponse() (response *DescribeAvailableInstanceConfigInfoResponse) {
	response = &DescribeAvailableInstanceConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用的机型信息
func (c *Client) DescribeAvailableInstanceConfigInfo(request *DescribeAvailableInstanceConfigInfoRequest) (response *DescribeAvailableInstanceConfigInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableInstanceConfigInfoRequest()
	}
	response = NewDescribeAvailableInstanceConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersCostSumAggregatedRequest() (request *DescribeClustersCostSumAggregatedRequest) {
	request = &DescribeClustersCostSumAggregatedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersCostSumAggregated")
	return
}

func NewDescribeClustersCostSumAggregatedResponse() (response *DescribeClustersCostSumAggregatedResponse) {
	response = &DescribeClustersCostSumAggregatedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照维度聚合后的该维度对应的总成本数据，比如按照 聚合维度，cluster、namespace、node、deployment、daemonset、statefulset、job、controller、pod、container 聚合后的数据
func (c *Client) DescribeClustersCostSumAggregated(request *DescribeClustersCostSumAggregatedRequest) (response *DescribeClustersCostSumAggregatedResponse, err error) {
	if request == nil {
		request = NewDescribeClustersCostSumAggregatedRequest()
	}
	response = NewDescribeClustersCostSumAggregatedResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRouteTableRequest() (request *DeleteClusterRouteTableRequest) {
	request = &DeleteClusterRouteTableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRouteTable")
	return
}

func NewDeleteClusterRouteTableResponse() (response *DeleteClusterRouteTableResponse) {
	response = &DeleteClusterRouteTableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群路由表
func (c *Client) DeleteClusterRouteTable(request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRouteTableRequest()
	}
	response = NewDeleteClusterRouteTableResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMasterLogRequest() (request *DescribeMasterLogRequest) {
	request = &DescribeMasterLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeMasterLog")
	return
}

func NewDescribeMasterLogResponse() (response *DescribeMasterLogResponse) {
	response = &DescribeMasterLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询master组件日志开启信息
func (c *Client) DescribeMasterLog(request *DescribeMasterLogRequest) (response *DescribeMasterLogResponse, err error) {
	if request == nil {
		request = NewDescribeMasterLogRequest()
	}
	response = NewDescribeMasterLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTasks")
	return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务相关信息，只会查询对应任务类型的最新的一条任务状态
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	if request == nil {
		request = NewDescribeTasksRequest()
	}
	response = NewDescribeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeAvailableExtraArgsRequest() (request *DescribeEdgeAvailableExtraArgsRequest) {
	request = &DescribeEdgeAvailableExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeAvailableExtraArgs")
	return
}

func NewDescribeEdgeAvailableExtraArgsResponse() (response *DescribeEdgeAvailableExtraArgsResponse) {
	response = &DescribeEdgeAvailableExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘容器集群可用的自定义参数
func (c *Client) DescribeEdgeAvailableExtraArgs(request *DescribeEdgeAvailableExtraArgsRequest) (response *DescribeEdgeAvailableExtraArgsResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeAvailableExtraArgsRequest()
	}
	response = NewDescribeEdgeAvailableExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusInstancesOverviewRequest() (request *DescribePrometheusInstancesOverviewRequest) {
	request = &DescribePrometheusInstancesOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstancesOverview")
	return
}

func NewDescribePrometheusInstancesOverviewResponse() (response *DescribePrometheusInstancesOverviewResponse) {
	response = &DescribePrometheusInstancesOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取与云监控融合实例列表
func (c *Client) DescribePrometheusInstancesOverview(request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusInstancesOverviewRequest()
	}
	response = NewDescribePrometheusInstancesOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取容器服务支持的所有地域
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response = NewDescribeRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEKSInstancesRequest() (request *CreateEKSInstancesRequest) {
	request = &CreateEKSInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEKSInstances")
	return
}

func NewCreateEKSInstancesResponse() (response *CreateEKSInstancesResponse) {
	response = &CreateEKSInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性容器实例
func (c *Client) CreateEKSInstances(request *CreateEKSInstancesRequest) (response *CreateEKSInstancesResponse, err error) {
	if request == nil {
		request = NewCreateEKSInstancesRequest()
	}
	response = NewCreateEKSInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupStorageLocationRequest() (request *CreateBackupStorageLocationRequest) {
	request = &CreateBackupStorageLocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateBackupStorageLocation")
	return
}

func NewCreateBackupStorageLocationResponse() (response *CreateBackupStorageLocationResponse) {
	response = &CreateBackupStorageLocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建备份仓库，指定了存储仓库类型（如COS）、COS桶地区、名称等信息，当前最多允许创建100个仓库， 注意此接口当前是全局接口，多个地域的TKE集群如果要备份到相同的备份仓库中，不需要重复创建备份仓库
func (c *Client) CreateBackupStorageLocation(request *CreateBackupStorageLocationRequest) (response *CreateBackupStorageLocationResponse, err error) {
	if request == nil {
		request = NewCreateBackupStorageLocationRequest()
	}
	response = NewCreateBackupStorageLocationResponse()
	err = c.Send(request, response)
	return
}

func NewEnableMasterLogRequest() (request *EnableMasterLogRequest) {
	request = &EnableMasterLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableMasterLog")
	return
}

func NewEnableMasterLogResponse() (response *EnableMasterLogResponse) {
	response = &EnableMasterLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启master日志组件收集
func (c *Client) EnableMasterLog(request *EnableMasterLogRequest) (response *EnableMasterLogResponse, err error) {
	if request == nil {
		request = NewEnableMasterLogRequest()
	}
	response = NewEnableMasterLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetEksSpecsRequest() (request *GetEksSpecsRequest) {
	request = &GetEksSpecsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEksSpecs")
	return
}

func NewGetEksSpecsResponse() (response *GetEksSpecsResponse) {
	response = &GetEksSpecsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EKS规格
func (c *Client) GetEksSpecs(request *GetEksSpecsRequest) (response *GetEksSpecsResponse, err error) {
	if request == nil {
		request = NewGetEksSpecsRequest()
	}
	response = NewGetEksSpecsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSubaccountAuthorityRequest() (request *CheckSubaccountAuthorityRequest) {
	request = &CheckSubaccountAuthorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckSubaccountAuthority")
	return
}

func NewCheckSubaccountAuthorityResponse() (response *CheckSubaccountAuthorityResponse) {
	response = &CheckSubaccountAuthorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验子账户是否拥有操作某资源的权限，该接口类似Kubernetes APIServer的SelfSubjectAccessReview
func (c *Client) CheckSubaccountAuthority(request *CheckSubaccountAuthorityRequest) (response *CheckSubaccountAuthorityResponse, err error) {
	if request == nil {
		request = NewCheckSubaccountAuthorityRequest()
	}
	response = NewCheckSubaccountAuthorityResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusRecordRuleYamlRequest() (request *CreatePrometheusRecordRuleYamlRequest) {
	request = &CreatePrometheusRecordRuleYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusRecordRuleYaml")
	return
}

func NewCreatePrometheusRecordRuleYamlResponse() (response *CreatePrometheusRecordRuleYamlResponse) {
	response = &CreatePrometheusRecordRuleYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建聚合规则yaml方式
func (c *Client) CreatePrometheusRecordRuleYaml(request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusRecordRuleYamlRequest()
	}
	response = NewCreatePrometheusRecordRuleYamlResponse()
	err = c.Send(request, response)
	return
}

func NewNotifyResultRequest() (request *NotifyResultRequest) {
	request = &NotifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "NotifyResult")
	return
}

func NewNotifyResultResponse() (response *NotifyResultResponse) {
	response = &NotifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用作cube场景下通知操作的执行结果。
func (c *Client) NotifyResult(request *NotifyResultRequest) (response *NotifyResultResponse, err error) {
	if request == nil {
		request = NewNotifyResultRequest()
	}
	response = NewNotifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterCIDRFromCcnRequest() (request *DeleteClusterCIDRFromCcnRequest) {
	request = &DeleteClusterCIDRFromCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterCIDRFromCcn")
	return
}

func NewDeleteClusterCIDRFromCcnResponse() (response *DeleteClusterCIDRFromCcnResponse) {
	response = &DeleteClusterCIDRFromCcnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从云联网删除集群CIDR的路由
func (c *Client) DeleteClusterCIDRFromCcn(request *DeleteClusterCIDRFromCcnRequest) (response *DeleteClusterCIDRFromCcnResponse, err error) {
	if request == nil {
		request = NewDeleteClusterCIDRFromCcnRequest()
	}
	response = NewDeleteClusterCIDRFromCcnResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksAlarmSettingRequest() (request *DescribeEksAlarmSettingRequest) {
	request = &DescribeEksAlarmSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksAlarmSetting")
	return
}

func NewDescribeEksAlarmSettingResponse() (response *DescribeEksAlarmSettingResponse) {
	response = &DescribeEksAlarmSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表的监控告警是否设置
func (c *Client) DescribeEksAlarmSetting(request *DescribeEksAlarmSettingRequest) (response *DescribeEksAlarmSettingResponse, err error) {
	if request == nil {
		request = NewDescribeEksAlarmSettingRequest()
	}
	response = NewDescribeEksAlarmSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdCredentialsRequest() (request *DescribeEtcdCredentialsRequest) {
	request = &DescribeEtcdCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdCredentials")
	return
}

func NewDescribeEtcdCredentialsResponse() (response *DescribeEtcdCredentialsResponse) {
	response = &DescribeEtcdCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询etcd访问凭证
func (c *Client) DescribeEtcdCredentials(request *DescribeEtcdCredentialsRequest) (response *DescribeEtcdCredentialsResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdCredentialsRequest()
	}
	response = NewDescribeEtcdCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusAlertHistoryRequest() (request *DescribePrometheusAlertHistoryRequest) {
	request = &DescribePrometheusAlertHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertHistory")
	return
}

func NewDescribePrometheusAlertHistoryResponse() (response *DescribePrometheusAlertHistoryResponse) {
	response = &DescribePrometheusAlertHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警历史
func (c *Client) DescribePrometheusAlertHistory(request *DescribePrometheusAlertHistoryRequest) (response *DescribePrometheusAlertHistoryResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusAlertHistoryRequest()
	}
	response = NewDescribePrometheusAlertHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewRunInContainerRequest() (request *RunInContainerRequest) {
	request = &RunInContainerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RunInContainer")
	return
}

func NewRunInContainerResponse() (response *RunInContainerResponse) {
	response = &RunInContainerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RunInContainer
func (c *Client) RunInContainer(request *RunInContainerRequest) (response *RunInContainerResponse, err error) {
	if request == nil {
		request = NewRunInContainerRequest()
	}
	response = NewRunInContainerResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterKubeconfigRequest() (request *UpdateClusterKubeconfigRequest) {
	request = &UpdateClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterKubeconfig")
	return
}

func NewUpdateClusterKubeconfigResponse() (response *UpdateClusterKubeconfigResponse) {
	response = &UpdateClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对集群的Kubeconfig信息进行更新
func (c *Client) UpdateClusterKubeconfig(request *UpdateClusterKubeconfigRequest) (response *UpdateClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewUpdateClusterKubeconfigRequest()
	}
	response = NewUpdateClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifySecretRequest() (request *ModifySecretRequest) {
	request = &ModifySecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifySecret")
	return
}

func NewModifySecretResponse() (response *ModifySecretResponse) {
	response = &ModifySecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改秘钥
func (c *Client) ModifySecret(request *ModifySecretRequest) (response *ModifySecretResponse, err error) {
	if request == nil {
		request = NewModifySecretRequest()
	}
	response = NewModifySecretResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTkeMigrateRequest() (request *DescribeTkeMigrateRequest) {
	request = &DescribeTkeMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTkeMigrate")
	return
}

func NewDescribeTkeMigrateResponse() (response *DescribeTkeMigrateResponse) {
	response = &DescribeTkeMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TKE集群迁移状态，用于控制台流量切换
func (c *Client) DescribeTkeMigrate(request *DescribeTkeMigrateRequest) (response *DescribeTkeMigrateResponse, err error) {
	if request == nil {
		request = NewDescribeTkeMigrateRequest()
	}
	response = NewDescribeTkeMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableTKEEdgeVersionRequest() (request *DescribeAvailableTKEEdgeVersionRequest) {
	request = &DescribeAvailableTKEEdgeVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableTKEEdgeVersion")
	return
}

func NewDescribeAvailableTKEEdgeVersionResponse() (response *DescribeAvailableTKEEdgeVersionResponse) {
	response = &DescribeAvailableTKEEdgeVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 边缘计算支持版本和k8s版本
func (c *Client) DescribeAvailableTKEEdgeVersion(request *DescribeAvailableTKEEdgeVersionRequest) (response *DescribeAvailableTKEEdgeVersionResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableTKEEdgeVersionRequest()
	}
	response = NewDescribeAvailableTKEEdgeVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersResourceStatusRequest() (request *DescribeClustersResourceStatusRequest) {
	request = &DescribeClustersResourceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersResourceStatus")
	return
}

func NewDescribeClustersResourceStatusResponse() (response *DescribeClustersResourceStatusResponse) {
	response = &DescribeClustersResourceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群资源状态
func (c *Client) DescribeClustersResourceStatus(request *DescribeClustersResourceStatusRequest) (response *DescribeClustersResourceStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClustersResourceStatusRequest()
	}
	response = NewDescribeClustersResourceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpcCniNetworkTypeRequest() (request *EnableVpcCniNetworkTypeRequest) {
	request = &EnableVpcCniNetworkTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableVpcCniNetworkType")
	return
}

func NewEnableVpcCniNetworkTypeResponse() (response *EnableVpcCniNetworkTypeResponse) {
	response = &EnableVpcCniNetworkTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GR集群可以通过本接口附加vpc-cni容器网络插件，开启vpc-cni容器网络能力
func (c *Client) EnableVpcCniNetworkType(request *EnableVpcCniNetworkTypeRequest) (response *EnableVpcCniNetworkTypeResponse, err error) {
	if request == nil {
		request = NewEnableVpcCniNetworkTypeRequest()
	}
	response = NewEnableVpcCniNetworkTypeResponse()
	err = c.Send(request, response)
	return
}

func NewListClustersRequest() (request *ListClustersRequest) {
	request = &ListClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListClusters")
	return
}

func NewListClustersResponse() (response *ListClustersResponse) {
	response = &ListClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表，返回一个用户下某个地域的所有集群(YUNAPI V3版本)
func (c *Client) ListClusters(request *ListClustersRequest) (response *ListClustersResponse, err error) {
	if request == nil {
		request = NewListClustersRequest()
	}
	response = NewListClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersCostTopAggregatedRequest() (request *DescribeClustersCostTopAggregatedRequest) {
	request = &DescribeClustersCostTopAggregatedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersCostTopAggregated")
	return
}

func NewDescribeClustersCostTopAggregatedResponse() (response *DescribeClustersCostTopAggregatedResponse) {
	response = &DescribeClustersCostTopAggregatedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群成本相关聚合数据的TOP N排序列表，比如按照 聚合维度，cluster、namespace、node、deployment、daemonset、statefulset、job、controller、pod、container 聚合后的数据，包括 cpu、ram使用效率等
func (c *Client) DescribeClustersCostTopAggregated(request *DescribeClustersCostTopAggregatedRequest) (response *DescribeClustersCostTopAggregatedResponse, err error) {
	if request == nil {
		request = NewDescribeClustersCostTopAggregatedRequest()
	}
	response = NewDescribeClustersCostTopAggregatedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSInstanceRegionsRequest() (request *DescribeEKSInstanceRegionsRequest) {
	request = &DescribeEKSInstanceRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSInstanceRegions")
	return
}

func NewDescribeEKSInstanceRegionsResponse() (response *DescribeEKSInstanceRegionsResponse) {
	response = &DescribeEKSInstanceRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器实例支持的地域
func (c *Client) DescribeEKSInstanceRegions(request *DescribeEKSInstanceRegionsRequest) (response *DescribeEKSInstanceRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeEKSInstanceRegionsRequest()
	}
	response = NewDescribeEKSInstanceRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSecurityRequest() (request *DescribeClusterSecurityRequest) {
	request = &DescribeClusterSecurityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSecurity")
	return
}

func NewDescribeClusterSecurityResponse() (response *DescribeClusterSecurityResponse) {
	response = &DescribeClusterSecurityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群的密钥信息
func (c *Client) DescribeClusterSecurity(request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSecurityRequest()
	}
	response = NewDescribeClusterSecurityResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodStatusRequest() (request *GetPodStatusRequest) {
	request = &GetPodStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodStatus")
	return
}

func NewGetPodStatusResponse() (response *GetPodStatusResponse) {
	response = &GetPodStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetPodStatus
func (c *Client) GetPodStatus(request *GetPodStatusRequest) (response *GetPodStatusResponse, err error) {
	if request == nil {
		request = NewGetPodStatusRequest()
	}
	response = NewGetPodStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEKSDiskRequest() (request *DeleteEKSDiskRequest) {
	request = &DeleteEKSDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSDisk")
	return
}

func NewDeleteEKSDiskResponse() (response *DeleteEKSDiskResponse) {
	response = &DeleteEKSDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用 CBS. 接口，删除 EKS 创建的隐藏云硬盘
func (c *Client) DeleteEKSDisk(request *DeleteEKSDiskRequest) (response *DeleteEKSDiskResponse, err error) {
	if request == nil {
		request = NewDeleteEKSDiskRequest()
	}
	response = NewDeleteEKSDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusScrapeStatisticsRequest() (request *DescribePrometheusScrapeStatisticsRequest) {
	request = &DescribePrometheusScrapeStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusScrapeStatistics")
	return
}

func NewDescribePrometheusScrapeStatisticsResponse() (response *DescribePrometheusScrapeStatisticsResponse) {
	response = &DescribePrometheusScrapeStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例采集速率信息
func (c *Client) DescribePrometheusScrapeStatistics(request *DescribePrometheusScrapeStatisticsRequest) (response *DescribePrometheusScrapeStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusScrapeStatisticsRequest()
	}
	response = NewDescribePrometheusScrapeStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCreateProgressRequest() (request *DescribeClusterCreateProgressRequest) {
	request = &DescribeClusterCreateProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCreateProgress")
	return
}

func NewDescribeClusterCreateProgressResponse() (response *DescribeClusterCreateProgressResponse) {
	response = &DescribeClusterCreateProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群创建进度
func (c *Client) DescribeClusterCreateProgress(request *DescribeClusterCreateProgressRequest) (response *DescribeClusterCreateProgressResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCreateProgressRequest()
	}
	response = NewDescribeClusterCreateProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sWorkloadPodsRequest() (request *DescribeK8sWorkloadPodsRequest) {
	request = &DescribeK8sWorkloadPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeK8sWorkloadPods")
	return
}

func NewDescribeK8sWorkloadPodsResponse() (response *DescribeK8sWorkloadPodsResponse) {
	response = &DescribeK8sWorkloadPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作负载的配置
func (c *Client) DescribeK8sWorkloadPods(request *DescribeK8sWorkloadPodsRequest) (response *DescribeK8sWorkloadPodsResponse, err error) {
	if request == nil {
		request = NewDescribeK8sWorkloadPodsRequest()
	}
	response = NewDescribeK8sWorkloadPodsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalNodeScriptRequest() (request *DescribeExternalNodeScriptRequest) {
	request = &DescribeExternalNodeScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalNodeScript")
	return
}

func NewDescribeExternalNodeScriptResponse() (response *DescribeExternalNodeScriptResponse) {
	response = &DescribeExternalNodeScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取第三方节点添加脚本
func (c *Client) DescribeExternalNodeScript(request *DescribeExternalNodeScriptRequest) (response *DescribeExternalNodeScriptResponse, err error) {
	if request == nil {
		request = NewDescribeExternalNodeScriptRequest()
	}
	response = NewDescribeExternalNodeScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCOpenStatusRequest() (request *DescribeTDCCOpenStatusRequest) {
	request = &DescribeTDCCOpenStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCOpenStatus")
	return
}

func NewDescribeTDCCOpenStatusResponse() (response *DescribeTDCCOpenStatusResponse) {
	response = &DescribeTDCCOpenStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查TDCC产品开通状态
func (c *Client) DescribeTDCCOpenStatus(request *DescribeTDCCOpenStatusRequest) (response *DescribeTDCCOpenStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCOpenStatusRequest()
	}
	response = NewDescribeTDCCOpenStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEksAuditRequest() (request *DisableEksAuditRequest) {
	request = &DisableEksAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEksAudit")
	return
}

func NewDisableEksAuditResponse() (response *DisableEksAuditResponse) {
	response = &DisableEksAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭弹性集群审计
func (c *Client) DisableEksAudit(request *DisableEksAuditRequest) (response *DisableEksAuditResponse, err error) {
	if request == nil {
		request = NewDisableEksAuditRequest()
	}
	response = NewDisableEksAuditResponse()
	err = c.Send(request, response)
	return
}

func NewCreateKnativeServiceRequest() (request *CreateKnativeServiceRequest) {
	request = &CreateKnativeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateKnativeService")
	return
}

func NewCreateKnativeServiceResponse() (response *CreateKnativeServiceResponse) {
	response = &CreateKnativeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建knative服务
func (c *Client) CreateKnativeService(request *CreateKnativeServiceRequest) (response *CreateKnativeServiceResponse, err error) {
	if request == nil {
		request = NewCreateKnativeServiceRequest()
	}
	response = NewCreateKnativeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodDeductionRateRequest() (request *DescribePodDeductionRateRequest) {
	request = &DescribePodDeductionRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePodDeductionRate")
	return
}

func NewDescribePodDeductionRateResponse() (response *DescribePodDeductionRateResponse) {
	response = &DescribePodDeductionRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各个规格的 Pod 的抵扣率
func (c *Client) DescribePodDeductionRate(request *DescribePodDeductionRateRequest) (response *DescribePodDeductionRateResponse, err error) {
	if request == nil {
		request = NewDescribePodDeductionRateRequest()
	}
	response = NewDescribePodDeductionRateResponse()
	err = c.Send(request, response)
	return
}

func NewTriggerAlarmBreakerRequest() (request *TriggerAlarmBreakerRequest) {
	request = &TriggerAlarmBreakerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "TriggerAlarmBreaker")
	return
}

func NewTriggerAlarmBreakerResponse() (response *TriggerAlarmBreakerResponse) {
	response = &TriggerAlarmBreakerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警触发熔断的回调接口
func (c *Client) TriggerAlarmBreaker(request *TriggerAlarmBreakerRequest) (response *TriggerAlarmBreakerResponse, err error) {
	if request == nil {
		request = NewTriggerAlarmBreakerRequest()
	}
	response = NewTriggerAlarmBreakerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNodeUnitRequest() (request *DeleteNodeUnitRequest) {
	request = &DeleteNodeUnitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteNodeUnit")
	return
}

func NewDeleteNodeUnitResponse() (response *DeleteNodeUnitResponse) {
	response = &DeleteNodeUnitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除NodeUnit
func (c *Client) DeleteNodeUnit(request *DeleteNodeUnitRequest) (response *DeleteNodeUnitResponse, err error) {
	if request == nil {
		request = NewDeleteNodeUnitRequest()
	}
	response = NewDeleteNodeUnitResponse()
	err = c.Send(request, response)
	return
}

func NewForwardApplicationRequestV3Request() (request *ForwardApplicationRequestV3Request) {
	request = &ForwardApplicationRequestV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardApplicationRequestV3")
	return
}

func NewForwardApplicationRequestV3Response() (response *ForwardApplicationRequestV3Response) {
	response = &ForwardApplicationRequestV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作TKE集群的addon
func (c *Client) ForwardApplicationRequestV3(request *ForwardApplicationRequestV3Request) (response *ForwardApplicationRequestV3Response, err error) {
	if request == nil {
		request = NewForwardApplicationRequestV3Request()
	}
	response = NewForwardApplicationRequestV3Response()
	err = c.Send(request, response)
	return
}

func NewDescribeHelmChartRequest() (request *DescribeHelmChartRequest) {
	request = &DescribeHelmChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHelmChart")
	return
}

func NewDescribeHelmChartResponse() (response *DescribeHelmChartResponse) {
	response = &DescribeHelmChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取chart列表
func (c *Client) DescribeHelmChart(request *DescribeHelmChartRequest) (response *DescribeHelmChartResponse, err error) {
	if request == nil {
		request = NewDescribeHelmChartRequest()
	}
	response = NewDescribeHelmChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusAgentInstancesRequest() (request *DescribePrometheusAgentInstancesRequest) {
	request = &DescribePrometheusAgentInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAgentInstances")
	return
}

func NewDescribePrometheusAgentInstancesResponse() (response *DescribePrometheusAgentInstancesResponse) {
	response = &DescribePrometheusAgentInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取关联目标集群的实例列表
func (c *Client) DescribePrometheusAgentInstances(request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusAgentInstancesRequest()
	}
	response = NewDescribePrometheusAgentInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeClusterStatusRequest() (request *DescribeTKEEdgeClusterStatusRequest) {
	request = &DescribeTKEEdgeClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterStatus")
	return
}

func NewDescribeTKEEdgeClusterStatusResponse() (response *DescribeTKEEdgeClusterStatusResponse) {
	response = &DescribeTKEEdgeClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘计算集群的当前状态以及过程信息
func (c *Client) DescribeTKEEdgeClusterStatus(request *DescribeTKEEdgeClusterStatusRequest) (response *DescribeTKEEdgeClusterStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeClusterStatusRequest()
	}
	response = NewDescribeTKEEdgeClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEksAlarmPolicyRequest() (request *ModifyEksAlarmPolicyRequest) {
	request = &ModifyEksAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyEksAlarmPolicy")
	return
}

func NewModifyEksAlarmPolicyResponse() (response *ModifyEksAlarmPolicyResponse) {
	response = &ModifyEksAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改弹性集群告警策略
func (c *Client) ModifyEksAlarmPolicy(request *ModifyEksAlarmPolicyRequest) (response *ModifyEksAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewModifyEksAlarmPolicyRequest()
	}
	response = NewModifyEksAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewForwardRequestForDev1Request() (request *ForwardRequestForDev1Request) {
	request = &ForwardRequestForDev1Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardRequestForDev1")
	return
}

func NewForwardRequestForDev1Response() (response *ForwardRequestForDev1Response) {
	response = &ForwardRequestForDev1Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// YUNAPI 转发请求给TKE APIServer接口
// 该接口用于开发联调或测试
func (c *Client) ForwardRequestForDev1(request *ForwardRequestForDev1Request) (response *ForwardRequestForDev1Response, err error) {
	if request == nil {
		request = NewForwardRequestForDev1Request()
	}
	response = NewForwardRequestForDev1Response()
	err = c.Send(request, response)
	return
}

func NewScaleInClusterMasterRequest() (request *ScaleInClusterMasterRequest) {
	request = &ScaleInClusterMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ScaleInClusterMaster")
	return
}

func NewScaleInClusterMasterResponse() (response *ScaleInClusterMasterResponse) {
	response = &ScaleInClusterMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 缩容独立集群master节点，本功能为内测能力，使用之前请先提单联系我们。
func (c *Client) ScaleInClusterMaster(request *ScaleInClusterMasterRequest) (response *ScaleInClusterMasterResponse, err error) {
	if request == nil {
		request = NewScaleInClusterMasterRequest()
	}
	response = NewScaleInClusterMasterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCommonNamesRequest() (request *DescribeClusterCommonNamesRequest) {
	request = &DescribeClusterCommonNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCommonNames")
	return
}

func NewDescribeClusterCommonNamesResponse() (response *DescribeClusterCommonNamesResponse) {
	response = &DescribeClusterCommonNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定子账户在RBAC授权模式中对应kube-apiserver客户端证书的CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为50
func (c *Client) DescribeClusterCommonNames(request *DescribeClusterCommonNamesRequest) (response *DescribeClusterCommonNamesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCommonNamesRequest()
	}
	response = NewDescribeClusterCommonNamesResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodSecurityGroupsRequest() (request *GetPodSecurityGroupsRequest) {
	request = &GetPodSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodSecurityGroups")
	return
}

func NewGetPodSecurityGroupsResponse() (response *GetPodSecurityGroupsResponse) {
	response = &GetPodSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全组
func (c *Client) GetPodSecurityGroups(request *GetPodSecurityGroupsRequest) (response *GetPodSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewGetPodSecurityGroupsRequest()
	}
	response = NewGetPodSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCloudResourcesRequest() (request *DescribeClusterCloudResourcesRequest) {
	request = &DescribeClusterCloudResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCloudResources")
	return
}

func NewDescribeClusterCloudResourcesResponse() (response *DescribeClusterCloudResourcesResponse) {
	response = &DescribeClusterCloudResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群内创建的其他云产品资源列表，目前支持clb、cvm
func (c *Client) DescribeClusterCloudResources(request *DescribeClusterCloudResourcesRequest) (response *DescribeClusterCloudResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCloudResourcesRequest()
	}
	response = NewDescribeClusterCloudResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterHealthyStatusRequest() (request *DescribeClusterHealthyStatusRequest) {
	request = &DescribeClusterHealthyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterHealthyStatus")
	return
}

func NewDescribeClusterHealthyStatusResponse() (response *DescribeClusterHealthyStatusResponse) {
	response = &DescribeClusterHealthyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述集群目前的健康状态
func (c *Client) DescribeClusterHealthyStatus(request *DescribeClusterHealthyStatusRequest) (response *DescribeClusterHealthyStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterHealthyStatusRequest()
	}
	response = NewDescribeClusterHealthyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCheckPrometheusAlertWebhookRequest() (request *CheckPrometheusAlertWebhookRequest) {
	request = &CheckPrometheusAlertWebhookRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckPrometheusAlertWebhook")
	return
}

func NewCheckPrometheusAlertWebhookResponse() (response *CheckPrometheusAlertWebhookResponse) {
	response = &CheckPrometheusAlertWebhookResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试告警渠道的Webhook
func (c *Client) CheckPrometheusAlertWebhook(request *CheckPrometheusAlertWebhookRequest) (response *CheckPrometheusAlertWebhookResponse, err error) {
	if request == nil {
		request = NewCheckPrometheusAlertWebhookRequest()
	}
	response = NewCheckPrometheusAlertWebhookResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusRecordRuleRequest() (request *DeletePrometheusRecordRuleRequest) {
	request = &DeletePrometheusRecordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusRecordRule")
	return
}

func NewDeletePrometheusRecordRuleResponse() (response *DeletePrometheusRecordRuleResponse) {
	response = &DeletePrometheusRecordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除聚合规则
func (c *Client) DeletePrometheusRecordRule(request *DeletePrometheusRecordRuleRequest) (response *DeletePrometheusRecordRuleResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusRecordRuleRequest()
	}
	response = NewDeletePrometheusRecordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEdgeClusterVersionRequest() (request *UpdateEdgeClusterVersionRequest) {
	request = &UpdateEdgeClusterVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEdgeClusterVersion")
	return
}

func NewUpdateEdgeClusterVersionResponse() (response *UpdateEdgeClusterVersionResponse) {
	response = &UpdateEdgeClusterVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级边缘集群组件到指定版本，此版本为TKEEdge专用版本。
func (c *Client) UpdateEdgeClusterVersion(request *UpdateEdgeClusterVersionRequest) (response *UpdateEdgeClusterVersionResponse, err error) {
	if request == nil {
		request = NewUpdateEdgeClusterVersionRequest()
	}
	response = NewUpdateEdgeClusterVersionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaFeatureRequest() (request *UpdateMetaFeatureRequest) {
	request = &UpdateMetaFeatureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateMetaFeature")
	return
}

func NewUpdateMetaFeatureResponse() (response *UpdateMetaFeatureResponse) {
	response = &UpdateMetaFeatureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TKE集群更新跨租户弹性网卡全局配置接口
func (c *Client) UpdateMetaFeature(request *UpdateMetaFeatureRequest) (response *UpdateMetaFeatureResponse, err error) {
	if request == nil {
		request = NewUpdateMetaFeatureRequest()
	}
	response = NewUpdateMetaFeatureResponse()
	err = c.Send(request, response)
	return
}

func NewCheckEksClusterCIDRRequest() (request *CheckEksClusterCIDRRequest) {
	request = &CheckEksClusterCIDRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckEksClusterCIDR")
	return
}

func NewCheckEksClusterCIDRResponse() (response *CheckEksClusterCIDRResponse) {
	response = &CheckEksClusterCIDRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查弹性集群的CIDR是否冲突
func (c *Client) CheckEksClusterCIDR(request *CheckEksClusterCIDRRequest) (response *CheckEksClusterCIDRResponse, err error) {
	if request == nil {
		request = NewCheckEksClusterCIDRRequest()
	}
	response = NewCheckEksClusterCIDRResponse()
	err = c.Send(request, response)
	return
}

func NewAddEksAlarmPolicyRequest() (request *AddEksAlarmPolicyRequest) {
	request = &AddEksAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddEksAlarmPolicy")
	return
}

func NewAddEksAlarmPolicyResponse() (response *AddEksAlarmPolicyResponse) {
	response = &AddEksAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加弹性集群告警策略
func (c *Client) AddEksAlarmPolicy(request *AddEksAlarmPolicyRequest) (response *AddEksAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewAddEksAlarmPolicyRequest()
	}
	response = NewAddEksAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageCacheRequest() (request *CreateImageCacheRequest) {
	request = &CreateImageCacheRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateImageCache")
	return
}

func NewCreateImageCacheResponse() (response *CreateImageCacheResponse) {
	response = &CreateImageCacheResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像缓存的接口。创建过程中，请勿删除EKSCI实例和云盘，否则镜像缓存将创建失败。
func (c *Client) CreateImageCache(request *CreateImageCacheRequest) (response *CreateImageCacheResponse, err error) {
	if request == nil {
		request = NewCreateImageCacheRequest()
	}
	response = NewCreateImageCacheResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTKEEdgeClusterRequest() (request *DeleteTKEEdgeClusterRequest) {
	request = &DeleteTKEEdgeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteTKEEdgeCluster")
	return
}

func NewDeleteTKEEdgeClusterResponse() (response *DeleteTKEEdgeClusterResponse) {
	response = &DeleteTKEEdgeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除边缘计算集群
func (c *Client) DeleteTKEEdgeCluster(request *DeleteTKEEdgeClusterRequest) (response *DeleteTKEEdgeClusterResponse, err error) {
	if request == nil {
		request = NewDeleteTKEEdgeClusterRequest()
	}
	response = NewDeleteTKEEdgeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnRoutesRequest() (request *DescribeCcnRoutesRequest) {
	request = &DescribeCcnRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCcnRoutes")
	return
}

func NewDescribeCcnRoutesResponse() (response *DescribeCcnRoutesResponse) {
	response = &DescribeCcnRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询tke集群CIDR是否加入云联网
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnRoutesRequest()
	}
	response = NewDescribeCcnRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExternalNodeRequest() (request *DeleteExternalNodeRequest) {
	request = &DeleteExternalNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteExternalNode")
	return
}

func NewDeleteExternalNodeResponse() (response *DeleteExternalNodeResponse) {
	response = &DeleteExternalNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除第三方节点
func (c *Client) DeleteExternalNode(request *DeleteExternalNodeRequest) (response *DeleteExternalNodeResponse, err error) {
	if request == nil {
		request = NewDeleteExternalNodeRequest()
	}
	response = NewDeleteExternalNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDisableResourceRecommendationRequest() (request *DisableResourceRecommendationRequest) {
	request = &DisableResourceRecommendationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableResourceRecommendation")
	return
}

func NewDisableResourceRecommendationResponse() (response *DisableResourceRecommendationResponse) {
	response = &DisableResourceRecommendationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭集群中的资源推荐服务
func (c *Client) DisableResourceRecommendation(request *DisableResourceRecommendationRequest) (response *DisableResourceRecommendationResponse, err error) {
	if request == nil {
		request = NewDisableResourceRecommendationRequest()
	}
	response = NewDisableResourceRecommendationResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusGlobalNotificationRequest() (request *ModifyPrometheusGlobalNotificationRequest) {
	request = &ModifyPrometheusGlobalNotificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusGlobalNotification")
	return
}

func NewModifyPrometheusGlobalNotificationResponse() (response *ModifyPrometheusGlobalNotificationResponse) {
	response = &ModifyPrometheusGlobalNotificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改全局告警通知渠道
func (c *Client) ModifyPrometheusGlobalNotification(request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusGlobalNotificationRequest()
	}
	response = NewModifyPrometheusGlobalNotificationResponse()
	err = c.Send(request, response)
	return
}

func NewActivateEKSClusterRequest() (request *ActivateEKSClusterRequest) {
	request = &ActivateEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ActivateEKSCluster")
	return
}

func NewActivateEKSClusterResponse() (response *ActivateEKSClusterResponse) {
	response = &ActivateEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 激活弹性集群
func (c *Client) ActivateEKSCluster(request *ActivateEKSClusterRequest) (response *ActivateEKSClusterResponse, err error) {
	if request == nil {
		request = NewActivateEKSClusterRequest()
	}
	response = NewActivateEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEncryptionProtectionRequest() (request *EnableEncryptionProtectionRequest) {
	request = &EnableEncryptionProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEncryptionProtection")
	return
}

func NewEnableEncryptionProtectionResponse() (response *EnableEncryptionProtectionResponse) {
	response = &EnableEncryptionProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启加密数据保护，需要先开启KMS能力，完成KMS授权
func (c *Client) EnableEncryptionProtection(request *EnableEncryptionProtectionRequest) (response *EnableEncryptionProtectionResponse, err error) {
	if request == nil {
		request = NewEnableEncryptionProtectionRequest()
	}
	response = NewEnableEncryptionProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDedicatedAPIServerRequest() (request *CreateDedicatedAPIServerRequest) {
	request = &CreateDedicatedAPIServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateDedicatedAPIServer")
	return
}

func NewCreateDedicatedAPIServerResponse() (response *CreateDedicatedAPIServerResponse) {
	response = &CreateDedicatedAPIServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专属API服务器
func (c *Client) CreateDedicatedAPIServer(request *CreateDedicatedAPIServerRequest) (response *CreateDedicatedAPIServerResponse, err error) {
	if request == nil {
		request = NewCreateDedicatedAPIServerRequest()
	}
	response = NewCreateDedicatedAPIServerResponse()
	err = c.Send(request, response)
	return
}

func NewCheckEdgeClusterCIDRRequest() (request *CheckEdgeClusterCIDRRequest) {
	request = &CheckEdgeClusterCIDRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckEdgeClusterCIDR")
	return
}

func NewCheckEdgeClusterCIDRResponse() (response *CheckEdgeClusterCIDRResponse) {
	response = &CheckEdgeClusterCIDRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查边缘计算集群的CIDR是否冲突
func (c *Client) CheckEdgeClusterCIDR(request *CheckEdgeClusterCIDRRequest) (response *CheckEdgeClusterCIDRResponse, err error) {
	if request == nil {
		request = NewCheckEdgeClusterCIDRRequest()
	}
	response = NewCheckEdgeClusterCIDRResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyPrometheusInstanceRequest() (request *DestroyPrometheusInstanceRequest) {
	request = &DestroyPrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DestroyPrometheusInstance")
	return
}

func NewDestroyPrometheusInstanceResponse() (response *DestroyPrometheusInstanceResponse) {
	response = &DestroyPrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁与云监控融合的V2版本实例
func (c *Client) DestroyPrometheusInstance(request *DestroyPrometheusInstanceRequest) (response *DestroyPrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewDestroyPrometheusInstanceRequest()
	}
	response = NewDestroyPrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusTemplateRequest() (request *ModifyPrometheusTemplateRequest) {
	request = &ModifyPrometheusTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusTemplate")
	return
}

func NewModifyPrometheusTemplateResponse() (response *ModifyPrometheusTemplateResponse) {
	response = &ModifyPrometheusTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改模板内容
func (c *Client) ModifyPrometheusTemplate(request *ModifyPrometheusTemplateRequest) (response *ModifyPrometheusTemplateResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusTemplateRequest()
	}
	response = NewModifyPrometheusTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEtcdSnapshotPolicyRequest() (request *DeleteEtcdSnapshotPolicyRequest) {
	request = &DeleteEtcdSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEtcdSnapshotPolicy")
	return
}

func NewDeleteEtcdSnapshotPolicyResponse() (response *DeleteEtcdSnapshotPolicyResponse) {
	response = &DeleteEtcdSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除etcd快照策略
func (c *Client) DeleteEtcdSnapshotPolicy(request *DeleteEtcdSnapshotPolicyRequest) (response *DeleteEtcdSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteEtcdSnapshotPolicyRequest()
	}
	response = NewDeleteEtcdSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodePoolDetailRequest() (request *DescribeClusterNodePoolDetailRequest) {
	request = &DescribeClusterNodePoolDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePoolDetail")
	return
}

func NewDescribeClusterNodePoolDetailResponse() (response *DescribeClusterNodePoolDetailResponse) {
	response = &DescribeClusterNodePoolDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点池详情
func (c *Client) DescribeClusterNodePoolDetail(request *DescribeClusterNodePoolDetailRequest) (response *DescribeClusterNodePoolDetailResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodePoolDetailRequest()
	}
	response = NewDescribeClusterNodePoolDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUseTKERequest() (request *CheckUseTKERequest) {
	request = &CheckUseTKERequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckUseTKE")
	return
}

func NewCheckUseTKEResponse() (response *CheckUseTKEResponse) {
	response = &CheckUseTKEResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否是TKE新用户，包括回归用户
func (c *Client) CheckUseTKE(request *CheckUseTKERequest) (response *CheckUseTKEResponse, err error) {
	if request == nil {
		request = NewCheckUseTKERequest()
	}
	response = NewCheckUseTKEResponse()
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

func NewModifyReservedInstanceScopeRequest() (request *ModifyReservedInstanceScopeRequest) {
	request = &ModifyReservedInstanceScopeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyReservedInstanceScope")
	return
}

func NewModifyReservedInstanceScopeResponse() (response *ModifyReservedInstanceScopeResponse) {
	response = &ModifyReservedInstanceScopeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改预留券的抵扣范围，抵扣范围取值：Region、Zone 和 Node。
func (c *Client) ModifyReservedInstanceScope(request *ModifyReservedInstanceScopeRequest) (response *ModifyReservedInstanceScopeResponse, err error) {
	if request == nil {
		request = NewModifyReservedInstanceScopeRequest()
	}
	response = NewModifyReservedInstanceScopeResponse()
	err = c.Send(request, response)
	return
}

func NewResumeClusterInstancesRequest() (request *ResumeClusterInstancesRequest) {
	request = &ResumeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ResumeClusterInstances")
	return
}

func NewResumeClusterInstancesResponse() (response *ResumeClusterInstancesResponse) {
	response = &ResumeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复集群节点升级，API 3.0
func (c *Client) ResumeClusterInstances(request *ResumeClusterInstancesRequest) (response *ResumeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewResumeClusterInstancesRequest()
	}
	response = NewResumeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewRotateClusterTokenRequest() (request *RotateClusterTokenRequest) {
	request = &RotateClusterTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RotateClusterToken")
	return
}

func NewRotateClusterTokenResponse() (response *RotateClusterTokenResponse) {
	response = &RotateClusterTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 轮转集群的token
func (c *Client) RotateClusterToken(request *RotateClusterTokenRequest) (response *RotateClusterTokenResponse, err error) {
	if request == nil {
		request = NewRotateClusterTokenRequest()
	}
	response = NewRotateClusterTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSClusterAuthorizationModeRequest() (request *DescribeEKSClusterAuthorizationModeRequest) {
	request = &DescribeEKSClusterAuthorizationModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterAuthorizationMode")
	return
}

func NewDescribeEKSClusterAuthorizationModeResponse() (response *DescribeEKSClusterAuthorizationModeResponse) {
	response = &DescribeEKSClusterAuthorizationModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取EKS集群授权模式
func (c *Client) DescribeEKSClusterAuthorizationMode(request *DescribeEKSClusterAuthorizationModeRequest) (response *DescribeEKSClusterAuthorizationModeResponse, err error) {
	if request == nil {
		request = NewDescribeEKSClusterAuthorizationModeRequest()
	}
	response = NewDescribeEKSClusterAuthorizationModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAsGroupOptionRequest() (request *DescribeClusterAsGroupOptionRequest) {
	request = &DescribeClusterAsGroupOptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroupOption")
	return
}

func NewDescribeClusterAsGroupOptionResponse() (response *DescribeClusterAsGroupOptionResponse) {
	response = &DescribeClusterAsGroupOptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群弹性伸缩配置
func (c *Client) DescribeClusterAsGroupOption(request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAsGroupOptionRequest()
	}
	response = NewDescribeClusterAsGroupOptionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePodRequest() (request *UpdatePodRequest) {
	request = &UpdatePodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdatePod")
	return
}

func NewUpdatePodResponse() (response *UpdatePodResponse) {
	response = &UpdatePodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// UpdatePod
func (c *Client) UpdatePod(request *UpdatePodRequest) (response *UpdatePodResponse, err error) {
	if request == nil {
		request = NewUpdatePodRequest()
	}
	response = NewUpdatePodResponse()
	err = c.Send(request, response)
	return
}

func NewGetEkletConfigRequest() (request *GetEkletConfigRequest) {
	request = &GetEkletConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEkletConfig")
	return
}

func NewGetEkletConfigResponse() (response *GetEkletConfigResponse) {
	response = &GetEkletConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过集群ID和配置项名称查询配置内容。
func (c *Client) GetEkletConfig(request *GetEkletConfigRequest) (response *GetEkletConfigResponse, err error) {
	if request == nil {
		request = NewGetEkletConfigRequest()
	}
	response = NewGetEkletConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListEKSZoneRequest() (request *ListEKSZoneRequest) {
	request = &ListEKSZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListEKSZone")
	return
}

func NewListEKSZoneResponse() (response *ListEKSZoneResponse) {
	response = &ListEKSZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对应地域eks可用的可用区
func (c *Client) ListEKSZone(request *ListEKSZoneRequest) (response *ListEKSZoneResponse, err error) {
	if request == nil {
		request = NewListEKSZoneRequest()
	}
	response = NewListEKSZoneResponse()
	err = c.Send(request, response)
	return
}

func NewScaleOutClusterMasterRequest() (request *ScaleOutClusterMasterRequest) {
	request = &ScaleOutClusterMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ScaleOutClusterMaster")
	return
}

func NewScaleOutClusterMasterResponse() (response *ScaleOutClusterMasterResponse) {
	response = &ScaleOutClusterMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容独立集群master节点
func (c *Client) ScaleOutClusterMaster(request *ScaleOutClusterMasterRequest) (response *ScaleOutClusterMasterResponse, err error) {
	if request == nil {
		request = NewScaleOutClusterMasterRequest()
	}
	response = NewScaleOutClusterMasterResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeVipVirtualClusterRequest() (request *UpgradeVipVirtualClusterRequest) {
	request = &UpgradeVipVirtualClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeVipVirtualCluster")
	return
}

func NewUpgradeVipVirtualClusterResponse() (response *UpgradeVipVirtualClusterResponse) {
	response = &UpgradeVipVirtualClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群升级Vip接口
func (c *Client) UpgradeVipVirtualCluster(request *UpgradeVipVirtualClusterRequest) (response *UpgradeVipVirtualClusterResponse, err error) {
	if request == nil {
		request = NewUpgradeVipVirtualClusterRequest()
	}
	response = NewUpgradeVipVirtualClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCloudRunHPARequest() (request *CreateCloudRunHPARequest) {
	request = &CreateCloudRunHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateCloudRunHPA")
	return
}

func NewCreateCloudRunHPAResponse() (response *CreateCloudRunHPAResponse) {
	response = &CreateCloudRunHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建cloudrun伸缩组。
func (c *Client) CreateCloudRunHPA(request *CreateCloudRunHPARequest) (response *CreateCloudRunHPAResponse, err error) {
	if request == nil {
		request = NewCreateCloudRunHPARequest()
	}
	response = NewCreateCloudRunHPAResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBootstrapTokensRequest() (request *DeleteBootstrapTokensRequest) {
	request = &DeleteBootstrapTokensRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteBootstrapTokens")
	return
}

func NewDeleteBootstrapTokensResponse() (response *DeleteBootstrapTokensResponse) {
	response = &DeleteBootstrapTokensResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除BootstrapToken
func (c *Client) DeleteBootstrapTokens(request *DeleteBootstrapTokensRequest) (response *DeleteBootstrapTokensResponse, err error) {
	if request == nil {
		request = NewDeleteBootstrapTokensRequest()
	}
	response = NewDeleteBootstrapTokensResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterEndpointRequest() (request *DeleteClusterEndpointRequest) {
	request = &DeleteClusterEndpointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpoint")
	return
}

func NewDeleteClusterEndpointResponse() (response *DeleteClusterEndpointResponse) {
	response = &DeleteClusterEndpointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群访问端口
func (c *Client) DeleteClusterEndpoint(request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
	if request == nil {
		request = NewDeleteClusterEndpointRequest()
	}
	response = NewDeleteClusterEndpointResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRouteRequest() (request *DeleteClusterRouteRequest) {
	request = &DeleteClusterRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRoute")
	return
}

func NewDeleteClusterRouteResponse() (response *DeleteClusterRouteResponse) {
	response = &DeleteClusterRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群路由
func (c *Client) DeleteClusterRoute(request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRouteRequest()
	}
	response = NewDeleteClusterRouteResponse()
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

func NewDescribeTKEEdgeScriptRequest() (request *DescribeTKEEdgeScriptRequest) {
	request = &DescribeTKEEdgeScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeScript")
	return
}

func NewDescribeTKEEdgeScriptResponse() (response *DescribeTKEEdgeScriptResponse) {
	response = &DescribeTKEEdgeScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘脚本链接，此接口用于添加第三方节点，通过下载脚本从而将节点添加到边缘集群。
func (c *Client) DescribeTKEEdgeScript(request *DescribeTKEEdgeScriptRequest) (response *DescribeTKEEdgeScriptResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeScriptRequest()
	}
	response = NewDescribeTKEEdgeScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterReleasesRequest() (request *DescribeClusterReleasesRequest) {
	request = &DescribeClusterReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleases")
	return
}

func NewDescribeClusterReleasesResponse() (response *DescribeClusterReleasesResponse) {
	response = &DescribeClusterReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群在应用市场中已安装应用列表
func (c *Client) DescribeClusterReleases(request *DescribeClusterReleasesRequest) (response *DescribeClusterReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterReleasesRequest()
	}
	response = NewDescribeClusterReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewListEKSClusterCertificatesRequest() (request *ListEKSClusterCertificatesRequest) {
	request = &ListEKSClusterCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListEKSClusterCertificates")
	return
}

func NewListEKSClusterCertificatesResponse() (response *ListEKSClusterCertificatesResponse) {
	response = &ListEKSClusterCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的证书列表，此接口接入CAM鉴权，只有管理权限的用户才可以正常调用，返回集群所有账户的证书列表
func (c *Client) ListEKSClusterCertificates(request *ListEKSClusterCertificatesRequest) (response *ListEKSClusterCertificatesResponse, err error) {
	if request == nil {
		request = NewListEKSClusterCertificatesRequest()
	}
	response = NewListEKSClusterCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterNodePoolRequest() (request *DeleteClusterNodePoolRequest) {
	request = &DeleteClusterNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterNodePool")
	return
}

func NewDeleteClusterNodePoolResponse() (response *DeleteClusterNodePoolResponse) {
	response = &DeleteClusterNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点池
func (c *Client) DeleteClusterNodePool(request *DeleteClusterNodePoolRequest) (response *DeleteClusterNodePoolResponse, err error) {
	if request == nil {
		request = NewDeleteClusterNodePoolRequest()
	}
	response = NewDeleteClusterNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像信息
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	response = NewDescribeImagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEtcdAttributeRequest() (request *ModifyEtcdAttributeRequest) {
	request = &ModifyEtcdAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyEtcdAttribute")
	return
}

func NewModifyEtcdAttributeResponse() (response *ModifyEtcdAttributeResponse) {
	response = &ModifyEtcdAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改etcd实例属性
func (c *Client) ModifyEtcdAttribute(request *ModifyEtcdAttributeRequest) (response *ModifyEtcdAttributeResponse, err error) {
	if request == nil {
		request = NewModifyEtcdAttributeRequest()
	}
	response = NewModifyEtcdAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeChartDownloadInfoRequest() (request *DescribeChartDownloadInfoRequest) {
	request = &DescribeChartDownloadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeChartDownloadInfo")
	return
}

func NewDescribeChartDownloadInfoResponse() (response *DescribeChartDownloadInfoResponse) {
	response = &DescribeChartDownloadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Chart的下载信息
func (c *Client) DescribeChartDownloadInfo(request *DescribeChartDownloadInfoRequest) (response *DescribeChartDownloadInfoResponse, err error) {
	if request == nil {
		request = NewDescribeChartDownloadInfoRequest()
	}
	response = NewDescribeChartDownloadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRoutesRequest() (request *DescribeClusterRoutesRequest) {
	request = &DescribeClusterRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRoutes")
	return
}

func NewDescribeClusterRoutesResponse() (response *DescribeClusterRoutesResponse) {
	response = &DescribeClusterRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群路由
func (c *Client) DescribeClusterRoutes(request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRoutesRequest()
	}
	response = NewDescribeClusterRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypesForDirectENIRequest() (request *DescribeInstanceTypesForDirectENIRequest) {
	request = &DescribeInstanceTypesForDirectENIRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeInstanceTypesForDirectENI")
	return
}

func NewDescribeInstanceTypesForDirectENIResponse() (response *DescribeInstanceTypesForDirectENIResponse) {
	response = &DescribeInstanceTypesForDirectENIResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看独立网卡模式可用的机型
func (c *Client) DescribeInstanceTypesForDirectENI(request *DescribeInstanceTypesForDirectENIRequest) (response *DescribeInstanceTypesForDirectENIResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypesForDirectENIRequest()
	}
	response = NewDescribeInstanceTypesForDirectENIResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterEndpointsRequest() (request *DescribeClusterEndpointsRequest) {
	request = &DescribeClusterEndpointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpoints")
	return
}

func NewDescribeClusterEndpointsResponse() (response *DescribeClusterEndpointsResponse) {
	response = &DescribeClusterEndpointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的访问地址，包括内网地址，外网地址，外网域名，外网访问安全策略
func (c *Client) DescribeClusterEndpoints(request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterEndpointsRequest()
	}
	response = NewDescribeClusterEndpointsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTDCCClusterKubeconfigRequest() (request *UpdateTDCCClusterKubeconfigRequest) {
	request = &UpdateTDCCClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateTDCCClusterKubeconfig")
	return
}

func NewUpdateTDCCClusterKubeconfigResponse() (response *UpdateTDCCClusterKubeconfigResponse) {
	response = &UpdateTDCCClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新注册集群的kubeconfig信息
func (c *Client) UpdateTDCCClusterKubeconfig(request *UpdateTDCCClusterKubeconfigRequest) (response *UpdateTDCCClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewUpdateTDCCClusterKubeconfigRequest()
	}
	response = NewUpdateTDCCClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetECMZoneResourceRequest() (request *GetECMZoneResourceRequest) {
	request = &GetECMZoneResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetECMZoneResource")
	return
}

func NewGetECMZoneResourceResponse() (response *GetECMZoneResourceResponse) {
	response = &GetECMZoneResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用区是否有可用资源
func (c *Client) GetECMZoneResource(request *GetECMZoneResourceRequest) (response *GetECMZoneResourceResponse, err error) {
	if request == nil {
		request = NewGetECMZoneResourceRequest()
	}
	response = NewGetECMZoneResourceResponse()
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

func NewDeleteClusterVirtualNodeRequest() (request *DeleteClusterVirtualNodeRequest) {
	request = &DeleteClusterVirtualNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNode")
	return
}

func NewDeleteClusterVirtualNodeResponse() (response *DeleteClusterVirtualNodeResponse) {
	response = &DeleteClusterVirtualNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除超级节点
func (c *Client) DeleteClusterVirtualNode(request *DeleteClusterVirtualNodeRequest) (response *DeleteClusterVirtualNodeResponse, err error) {
	if request == nil {
		request = NewDeleteClusterVirtualNodeRequest()
	}
	response = NewDeleteClusterVirtualNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeECMClustersRequest() (request *DescribeECMClustersRequest) {
	request = &DescribeECMClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeECMClusters")
	return
}

func NewDescribeECMClustersResponse() (response *DescribeECMClustersResponse) {
	response = &DescribeECMClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ECM EKS弹性集群列表
func (c *Client) DescribeECMClusters(request *DescribeECMClustersRequest) (response *DescribeECMClustersResponse, err error) {
	if request == nil {
		request = NewDescribeECMClustersRequest()
	}
	response = NewDescribeECMClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogConfigsRequest() (request *DescribeLogConfigsRequest) {
	request = &DescribeLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeLogConfigs")
	return
}

func NewDescribeLogConfigsResponse() (response *DescribeLogConfigsResponse) {
	response = &DescribeLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志采集规则
func (c *Client) DescribeLogConfigs(request *DescribeLogConfigsRequest) (response *DescribeLogConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeLogConfigsRequest()
	}
	response = NewDescribeLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterCIDRToVbcRequest() (request *AddClusterCIDRToVbcRequest) {
	request = &AddClusterCIDRToVbcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDRToVbc")
	return
}

func NewAddClusterCIDRToVbcResponse() (response *AddClusterCIDRToVbcResponse) {
	response = &AddClusterCIDRToVbcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布tke集群cidr到云联网
func (c *Client) AddClusterCIDRToVbc(request *AddClusterCIDRToVbcRequest) (response *AddClusterCIDRToVbcResponse, err error) {
	if request == nil {
		request = NewAddClusterCIDRToVbcRequest()
	}
	response = NewAddClusterCIDRToVbcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterReleaseDetailsRequest() (request *DescribeClusterReleaseDetailsRequest) {
	request = &DescribeClusterReleaseDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleaseDetails")
	return
}

func NewDescribeClusterReleaseDetailsResponse() (response *DescribeClusterReleaseDetailsResponse) {
	response = &DescribeClusterReleaseDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询通过应用市场安装的某个应用详情
func (c *Client) DescribeClusterReleaseDetails(request *DescribeClusterReleaseDetailsRequest) (response *DescribeClusterReleaseDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterReleaseDetailsRequest()
	}
	response = NewDescribeClusterReleaseDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpcPeerClusterRoutesRequest() (request *EnableVpcPeerClusterRoutesRequest) {
	request = &EnableVpcPeerClusterRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableVpcPeerClusterRoutes")
	return
}

func NewEnableVpcPeerClusterRoutesResponse() (response *EnableVpcPeerClusterRoutesResponse) {
	response = &EnableVpcPeerClusterRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动对等连接容器路由
func (c *Client) EnableVpcPeerClusterRoutes(request *EnableVpcPeerClusterRoutesRequest) (response *EnableVpcPeerClusterRoutesResponse, err error) {
	if request == nil {
		request = NewEnableVpcPeerClusterRoutesRequest()
	}
	response = NewEnableVpcPeerClusterRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewOpUpgradeClusterInstancesRequest() (request *OpUpgradeClusterInstancesRequest) {
	request = &OpUpgradeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "OpUpgradeClusterInstances")
	return
}

func NewOpUpgradeClusterInstancesResponse() (response *OpUpgradeClusterInstancesResponse) {
	response = &OpUpgradeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于控制节点升级任务
func (c *Client) OpUpgradeClusterInstances(request *OpUpgradeClusterInstancesRequest) (response *OpUpgradeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewOpUpgradeClusterInstancesRequest()
	}
	response = NewOpUpgradeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesVersionRequest() (request *DescribeInstancesVersionRequest) {
	request = &DescribeInstancesVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeInstancesVersion")
	return
}

func NewDescribeInstancesVersionResponse() (response *DescribeInstancesVersionResponse) {
	response = &DescribeInstancesVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// worker节点的版本统计
func (c *Client) DescribeInstancesVersion(request *DescribeInstancesVersionRequest) (response *DescribeInstancesVersionResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesVersionRequest()
	}
	response = NewDescribeInstancesVersionResponse()
	err = c.Send(request, response)
	return
}

func NewGetAvailableResourcesRequest() (request *GetAvailableResourcesRequest) {
	request = &GetAvailableResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetAvailableResources")
	return
}

func NewGetAvailableResourcesResponse() (response *GetAvailableResourcesResponse) {
	response = &GetAvailableResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前地域每个可用区的资源信息
func (c *Client) GetAvailableResources(request *GetAvailableResourcesRequest) (response *GetAvailableResourcesResponse, err error) {
	if request == nil {
		request = NewGetAvailableResourcesRequest()
	}
	response = NewGetAvailableResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePodRequest() (request *CreatePodRequest) {
	request = &CreatePodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePod")
	return
}

func NewCreatePodResponse() (response *CreatePodResponse) {
	response = &CreatePodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreatePod 创建pod
func (c *Client) CreatePod(request *CreatePodRequest) (response *CreatePodResponse, err error) {
	if request == nil {
		request = NewCreatePodRequest()
	}
	response = NewCreatePodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeLogSwitchesRequest() (request *DescribeEdgeLogSwitchesRequest) {
	request = &DescribeEdgeLogSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeLogSwitches")
	return
}

func NewDescribeEdgeLogSwitchesResponse() (response *DescribeEdgeLogSwitchesResponse) {
	response = &DescribeEdgeLogSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取事件、审计和日志的状态
func (c *Client) DescribeEdgeLogSwitches(request *DescribeEdgeLogSwitchesRequest) (response *DescribeEdgeLogSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeLogSwitchesRequest()
	}
	response = NewDescribeEdgeLogSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEKSClusterRequest() (request *DeleteEKSClusterRequest) {
	request = &DeleteEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSCluster")
	return
}

func NewDeleteEKSClusterResponse() (response *DeleteEKSClusterResponse) {
	response = &DeleteEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性集群(yunapiv3)
func (c *Client) DeleteEKSCluster(request *DeleteEKSClusterRequest) (response *DeleteEKSClusterResponse, err error) {
	if request == nil {
		request = NewDeleteEKSClusterRequest()
	}
	response = NewDeleteEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusDashboardRequest() (request *DeletePrometheusDashboardRequest) {
	request = &DeletePrometheusDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusDashboard")
	return
}

func NewDeletePrometheusDashboardResponse() (response *DeletePrometheusDashboardResponse) {
	response = &DeletePrometheusDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除grafana监控面板
func (c *Client) DeletePrometheusDashboard(request *DeletePrometheusDashboardRequest) (response *DeletePrometheusDashboardResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusDashboardRequest()
	}
	response = NewDeletePrometheusDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewGetContainerProbeResultRequest() (request *GetContainerProbeResultRequest) {
	request = &GetContainerProbeResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetContainerProbeResult")
	return
}

func NewGetContainerProbeResultResponse() (response *GetContainerProbeResultResponse) {
	response = &GetContainerProbeResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetContainerProbeResult
func (c *Client) GetContainerProbeResult(request *GetContainerProbeResultRequest) (response *GetContainerProbeResultResponse, err error) {
	if request == nil {
		request = NewGetContainerProbeResultRequest()
	}
	response = NewGetContainerProbeResultResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterInstancesRequest() (request *ListClusterInstancesRequest) {
	request = &ListClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListClusterInstances")
	return
}

func NewListClusterInstancesResponse() (response *ListClusterInstancesResponse) {
	response = &ListClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群下节点信息
func (c *Client) ListClusterInstances(request *ListClusterInstancesRequest) (response *ListClusterInstancesResponse, err error) {
	if request == nil {
		request = NewListClusterInstancesRequest()
	}
	response = NewListClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCostRequest() (request *DescribeClusterCostRequest) {
	request = &DescribeClusterCostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCost")
	return
}

func NewDescribeClusterCostResponse() (response *DescribeClusterCostResponse) {
	response = &DescribeClusterCostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看集群费用优化清单，列举集群整体的节点的费用，以及通过部分、全部使用超级节点来优化成本后的总费用。
func (c *Client) DescribeClusterCost(request *DescribeClusterCostRequest) (response *DescribeClusterCostResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCostRequest()
	}
	response = NewDescribeClusterCostResponse()
	err = c.Send(request, response)
	return
}

func NewEnableResourceRecommendationRequest() (request *EnableResourceRecommendationRequest) {
	request = &EnableResourceRecommendationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableResourceRecommendation")
	return
}

func NewEnableResourceRecommendationResponse() (response *EnableResourceRecommendationResponse) {
	response = &EnableResourceRecommendationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启资源推荐服务
func (c *Client) EnableResourceRecommendation(request *EnableResourceRecommendationRequest) (response *EnableResourceRecommendationResponse, err error) {
	if request == nil {
		request = NewEnableResourceRecommendationRequest()
	}
	response = NewEnableResourceRecommendationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCostTaskProgressRequest() (request *DescribeCostTaskProgressRequest) {
	request = &DescribeCostTaskProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCostTaskProgress")
	return
}

func NewDescribeCostTaskProgressResponse() (response *DescribeCostTaskProgressResponse) {
	response = &DescribeCostTaskProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取成本分析运维任务进度
func (c *Client) DescribeCostTaskProgress(request *DescribeCostTaskProgressRequest) (response *DescribeCostTaskProgressResponse, err error) {
	if request == nil {
		request = NewDescribeCostTaskProgressRequest()
	}
	response = NewDescribeCostTaskProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTDCCExternalClusterRequest() (request *DeleteTDCCExternalClusterRequest) {
	request = &DeleteTDCCExternalClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteTDCCExternalCluster")
	return
}

func NewDeleteTDCCExternalClusterResponse() (response *DeleteTDCCExternalClusterResponse) {
	response = &DeleteTDCCExternalClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除注册集群
func (c *Client) DeleteTDCCExternalCluster(request *DeleteTDCCExternalClusterRequest) (response *DeleteTDCCExternalClusterResponse, err error) {
	if request == nil {
		request = NewDeleteTDCCExternalClusterRequest()
	}
	response = NewDeleteTDCCExternalClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalClusterSpecRequest() (request *DescribeExternalClusterSpecRequest) {
	request = &DescribeExternalClusterSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalClusterSpec")
	return
}

func NewDescribeExternalClusterSpecResponse() (response *DescribeExternalClusterSpecResponse) {
	response = &DescribeExternalClusterSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导入第三方集群YAML定义
func (c *Client) DescribeExternalClusterSpec(request *DescribeExternalClusterSpecRequest) (response *DescribeExternalClusterSpecResponse, err error) {
	if request == nil {
		request = NewDescribeExternalClusterSpecRequest()
	}
	response = NewDescribeExternalClusterSpecResponse()
	err = c.Send(request, response)
	return
}

func NewEnableMetaFeatureForEksRequest() (request *EnableMetaFeatureForEksRequest) {
	request = &EnableMetaFeatureForEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableMetaFeatureForEks")
	return
}

func NewEnableMetaFeatureForEksResponse() (response *EnableMetaFeatureForEksResponse) {
	response = &EnableMetaFeatureForEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EKS集群开通跨租户弹性网卡接口
func (c *Client) EnableMetaFeatureForEks(request *EnableMetaFeatureForEksRequest) (response *EnableMetaFeatureForEksResponse, err error) {
	if request == nil {
		request = NewEnableMetaFeatureForEksRequest()
	}
	response = NewEnableMetaFeatureForEksResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEventPersistenceRequest() (request *DisableEventPersistenceRequest) {
	request = &DisableEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEventPersistence")
	return
}

func NewDisableEventPersistenceResponse() (response *DisableEventPersistenceResponse) {
	response = &DisableEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭事件持久化功能
func (c *Client) DisableEventPersistence(request *DisableEventPersistenceRequest) (response *DisableEventPersistenceResponse, err error) {
	if request == nil {
		request = NewDisableEventPersistenceRequest()
	}
	response = NewDisableEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeECMInstancesRequest() (request *DescribeECMInstancesRequest) {
	request = &DescribeECMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeECMInstances")
	return
}

func NewDescribeECMInstancesResponse() (response *DescribeECMInstancesResponse) {
	response = &DescribeECMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ECM实例相关信息
func (c *Client) DescribeECMInstances(request *DescribeECMInstancesRequest) (response *DescribeECMInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeECMInstancesRequest()
	}
	response = NewDescribeECMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEksAuditRequest() (request *EnableEksAuditRequest) {
	request = &EnableEksAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEksAudit")
	return
}

func NewEnableEksAuditResponse() (response *EnableEksAuditResponse) {
	response = &EnableEksAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启弹性集群审计
func (c *Client) EnableEksAudit(request *EnableEksAuditRequest) (response *EnableEksAuditResponse, err error) {
	if request == nil {
		request = NewEnableEksAuditRequest()
	}
	response = NewEnableEksAuditResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusAlertRuleRequest() (request *DeletePrometheusAlertRuleRequest) {
	request = &DeletePrometheusAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAlertRule")
	return
}

func NewDeletePrometheusAlertRuleResponse() (response *DeletePrometheusAlertRuleResponse) {
	response = &DeletePrometheusAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警规则
func (c *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusAlertRuleRequest()
	}
	response = NewDeletePrometheusAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeClusterExtraArgsRequest() (request *DescribeEdgeClusterExtraArgsRequest) {
	request = &DescribeEdgeClusterExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterExtraArgs")
	return
}

func NewDescribeEdgeClusterExtraArgsResponse() (response *DescribeEdgeClusterExtraArgsResponse) {
	response = &DescribeEdgeClusterExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘集群自定义参数
func (c *Client) DescribeEdgeClusterExtraArgs(request *DescribeEdgeClusterExtraArgsRequest) (response *DescribeEdgeClusterExtraArgsResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeClusterExtraArgsRequest()
	}
	response = NewDescribeEdgeClusterExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReservedInstanceUtilizationRateRequest() (request *DescribeReservedInstanceUtilizationRateRequest) {
	request = &DescribeReservedInstanceUtilizationRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeReservedInstanceUtilizationRate")
	return
}

func NewDescribeReservedInstanceUtilizationRateResponse() (response *DescribeReservedInstanceUtilizationRateResponse) {
	response = &DescribeReservedInstanceUtilizationRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各种规格类型的预留券使用率
func (c *Client) DescribeReservedInstanceUtilizationRate(request *DescribeReservedInstanceUtilizationRateRequest) (response *DescribeReservedInstanceUtilizationRateResponse, err error) {
	if request == nil {
		request = NewDescribeReservedInstanceUtilizationRateRequest()
	}
	response = NewDescribeReservedInstanceUtilizationRateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEncryptionStatusRequest() (request *DescribeEncryptionStatusRequest) {
	request = &DescribeEncryptionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEncryptionStatus")
	return
}

func NewDescribeEncryptionStatusResponse() (response *DescribeEncryptionStatusResponse) {
	response = &DescribeEncryptionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询etcd数据是否进行加密
func (c *Client) DescribeEncryptionStatus(request *DescribeEncryptionStatusRequest) (response *DescribeEncryptionStatusResponse, err error) {
	if request == nil {
		request = NewDescribeEncryptionStatusRequest()
	}
	response = NewDescribeEncryptionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDisableClusterAuditRequest() (request *DisableClusterAuditRequest) {
	request = &DisableClusterAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableClusterAudit")
	return
}

func NewDisableClusterAuditResponse() (response *DisableClusterAuditResponse) {
	response = &DisableClusterAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭集群审计
func (c *Client) DisableClusterAudit(request *DisableClusterAuditRequest) (response *DisableClusterAuditResponse, err error) {
	if request == nil {
		request = NewDisableClusterAuditRequest()
	}
	response = NewDisableClusterAuditResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePostNodeResourcesRequest() (request *DescribePostNodeResourcesRequest) {
	request = &DescribePostNodeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePostNodeResources")
	return
}

func NewDescribePostNodeResourcesResponse() (response *DescribePostNodeResourcesResponse) {
	response = &DescribePostNodeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 包括 Pod 资源统计和绑定的预留券资源统计。
func (c *Client) DescribePostNodeResources(request *DescribePostNodeResourcesRequest) (response *DescribePostNodeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribePostNodeResourcesRequest()
	}
	response = NewDescribePostNodeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeDefaultsettingsRequest() (request *DescribeTKEEdgeDefaultsettingsRequest) {
	request = &DescribeTKEEdgeDefaultsettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeDefaultsettings")
	return
}

func NewDescribeTKEEdgeDefaultsettingsResponse() (response *DescribeTKEEdgeDefaultsettingsResponse) {
	response = &DescribeTKEEdgeDefaultsettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 边缘计算默认节点配置
func (c *Client) DescribeTKEEdgeDefaultsettings(request *DescribeTKEEdgeDefaultsettingsRequest) (response *DescribeTKEEdgeDefaultsettingsResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeDefaultsettingsRequest()
	}
	response = NewDescribeTKEEdgeDefaultsettingsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterInspectionRequest() (request *ModifyClusterInspectionRequest) {
	request = &ModifyClusterInspectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterInspection")
	return
}

func NewModifyClusterInspectionResponse() (response *ModifyClusterInspectionResponse) {
	response = &ModifyClusterInspectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群巡检配置
func (c *Client) ModifyClusterInspection(request *ModifyClusterInspectionRequest) (response *ModifyClusterInspectionResponse, err error) {
	if request == nil {
		request = NewModifyClusterInspectionRequest()
	}
	response = NewModifyClusterInspectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBootstrapTokenRequest() (request *CreateBootstrapTokenRequest) {
	request = &CreateBootstrapTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateBootstrapToken")
	return
}

func NewCreateBootstrapTokenResponse() (response *CreateBootstrapTokenResponse) {
	response = &CreateBootstrapTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建BootstrapToken
func (c *Client) CreateBootstrapToken(request *CreateBootstrapTokenRequest) (response *CreateBootstrapTokenResponse, err error) {
	if request == nil {
		request = NewCreateBootstrapTokenRequest()
	}
	response = NewCreateBootstrapTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEKSClusterRequest() (request *CreateEKSClusterRequest) {
	request = &CreateEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEKSCluster")
	return
}

func NewCreateEKSClusterResponse() (response *CreateEKSClusterResponse) {
	response = &CreateEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性集群
func (c *Client) CreateEKSCluster(request *CreateEKSClusterRequest) (response *CreateEKSClusterResponse, err error) {
	if request == nil {
		request = NewCreateEKSClusterRequest()
	}
	response = NewCreateEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewOperateCostTaskRequest() (request *OperateCostTaskRequest) {
	request = &OperateCostTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "OperateCostTask")
	return
}

func NewOperateCostTaskResponse() (response *OperateCostTaskResponse) {
	response = &OperateCostTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作某个成本分析任务
func (c *Client) OperateCostTask(request *OperateCostTaskRequest) (response *OperateCostTaskResponse, err error) {
	if request == nil {
		request = NewOperateCostTaskRequest()
	}
	response = NewOperateCostTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterCIDRToCcnRequest() (request *AddClusterCIDRToCcnRequest) {
	request = &AddClusterCIDRToCcnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDRToCcn")
	return
}

func NewAddClusterCIDRToCcnResponse() (response *AddClusterCIDRToCcnResponse) {
	response = &AddClusterCIDRToCcnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加TKE集群CIDR到云联网
func (c *Client) AddClusterCIDRToCcn(request *AddClusterCIDRToCcnRequest) (response *AddClusterCIDRToCcnResponse, err error) {
	if request == nil {
		request = NewAddClusterCIDRToCcnRequest()
	}
	response = NewAddClusterCIDRToCcnResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserClustersRequest() (request *DescribeUserClustersRequest) {
	request = &DescribeUserClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeUserClusters")
	return
}

func NewDescribeUserClustersResponse() (response *DescribeUserClustersResponse) {
	response = &DescribeUserClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户tke、eks、tkeedge集群列表。
func (c *Client) DescribeUserClusters(request *DescribeUserClustersRequest) (response *DescribeUserClustersResponse, err error) {
	if request == nil {
		request = NewDescribeUserClustersRequest()
	}
	response = NewDescribeUserClustersResponse()
	err = c.Send(request, response)
	return
}

func NewGetPriceRequest() (request *GetPriceRequest) {
	request = &GetPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPrice")
	return
}

func NewGetPriceResponse() (response *GetPriceResponse) {
	response = &GetPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EKS询价接口
func (c *Client) GetPrice(request *GetPriceRequest) (response *GetPriceResponse, err error) {
	if request == nil {
		request = NewGetPriceRequest()
	}
	response = NewGetPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTemplatesRequest() (request *DescribePrometheusTemplatesRequest) {
	request = &DescribePrometheusTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemplates")
	return
}

func NewDescribePrometheusTemplatesResponse() (response *DescribePrometheusTemplatesResponse) {
	response = &DescribePrometheusTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取模板列表，默认模板将总是在最前面
func (c *Client) DescribePrometheusTemplates(request *DescribePrometheusTemplatesRequest) (response *DescribePrometheusTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTemplatesRequest()
	}
	response = NewDescribePrometheusTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEksEventPersistenceRequest() (request *DisableEksEventPersistenceRequest) {
	request = &DisableEksEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEksEventPersistence")
	return
}

func NewDisableEksEventPersistenceResponse() (response *DisableEksEventPersistenceResponse) {
	response = &DisableEksEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Eks集群关闭事件持久化功能
func (c *Client) DisableEksEventPersistence(request *DisableEksEventPersistenceRequest) (response *DisableEksEventPersistenceResponse, err error) {
	if request == nil {
		request = NewDisableEksEventPersistenceRequest()
	}
	response = NewDisableEksEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsLogTopicsRequest() (request *DescribeClsLogTopicsRequest) {
	request = &DescribeClsLogTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClsLogTopics")
	return
}

func NewDescribeClsLogTopicsResponse() (response *DescribeClsLogTopicsResponse) {
	response = &DescribeClsLogTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出CLS日志主题
func (c *Client) DescribeClsLogTopics(request *DescribeClsLogTopicsRequest) (response *DescribeClsLogTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeClsLogTopicsRequest()
	}
	response = NewDescribeClsLogTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterPendingReleasesRequest() (request *DescribeClusterPendingReleasesRequest) {
	request = &DescribeClusterPendingReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterPendingReleases")
	return
}

func NewDescribeClusterPendingReleasesResponse() (response *DescribeClusterPendingReleasesResponse) {
	response = &DescribeClusterPendingReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在应用市场中查询正在安装中的应用列表
func (c *Client) DescribeClusterPendingReleases(request *DescribeClusterPendingReleasesRequest) (response *DescribeClusterPendingReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterPendingReleasesRequest()
	}
	response = NewDescribeClusterPendingReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTargetsRequest() (request *DescribePrometheusTargetsRequest) {
	request = &DescribePrometheusTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTargets")
	return
}

func NewDescribePrometheusTargetsResponse() (response *DescribePrometheusTargetsResponse) {
	response = &DescribePrometheusTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取targets信息
func (c *Client) DescribePrometheusTargets(request *DescribePrometheusTargetsRequest) (response *DescribePrometheusTargetsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTargetsRequest()
	}
	response = NewDescribePrometheusTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewDisableClusterDeletionProtectionRequest() (request *DisableClusterDeletionProtectionRequest) {
	request = &DisableClusterDeletionProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableClusterDeletionProtection")
	return
}

func NewDisableClusterDeletionProtectionResponse() (response *DisableClusterDeletionProtectionResponse) {
	response = &DisableClusterDeletionProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭集群删除保护
func (c *Client) DisableClusterDeletionProtection(request *DisableClusterDeletionProtectionRequest) (response *DisableClusterDeletionProtectionResponse, err error) {
	if request == nil {
		request = NewDisableClusterDeletionProtectionRequest()
	}
	response = NewDisableClusterDeletionProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterSchedulerPolicyRequest() (request *ModifyClusterSchedulerPolicyRequest) {
	request = &ModifyClusterSchedulerPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterSchedulerPolicy")
	return
}

func NewModifyClusterSchedulerPolicyResponse() (response *ModifyClusterSchedulerPolicyResponse) {
	response = &ModifyClusterSchedulerPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群调度策略
func (c *Client) ModifyClusterSchedulerPolicy(request *ModifyClusterSchedulerPolicyRequest) (response *ModifyClusterSchedulerPolicyResponse, err error) {
	if request == nil {
		request = NewModifyClusterSchedulerPolicyRequest()
	}
	response = NewModifyClusterSchedulerPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryEKSDiskTaskRequest() (request *QueryEKSDiskTaskRequest) {
	request = &QueryEKSDiskTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "QueryEKSDiskTask")
	return
}

func NewQueryEKSDiskTaskResponse() (response *QueryEKSDiskTaskResponse) {
	response = &QueryEKSDiskTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过EKS 创盘返回的 TaskId，查询云硬盘状态及盘ID
func (c *Client) QueryEKSDiskTask(request *QueryEKSDiskTaskRequest) (response *QueryEKSDiskTaskResponse, err error) {
	if request == nil {
		request = NewQueryEKSDiskTaskRequest()
	}
	response = NewQueryEKSDiskTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListEKSPodsRequest() (request *ListEKSPodsRequest) {
	request = &ListEKSPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListEKSPods")
	return
}

func NewListEKSPodsResponse() (response *ListEKSPodsResponse) {
	response = &ListEKSPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取pod详细信息
func (c *Client) ListEKSPods(request *ListEKSPodsRequest) (response *ListEKSPodsResponse, err error) {
	if request == nil {
		request = NewListEKSPodsRequest()
	}
	response = NewListEKSPodsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterReleaseRequest() (request *CreateClusterReleaseRequest) {
	request = &CreateClusterReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRelease")
	return
}

func NewCreateClusterReleaseResponse() (response *CreateClusterReleaseResponse) {
	response = &CreateClusterReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群创建应用
func (c *Client) CreateClusterRelease(request *CreateClusterReleaseRequest) (response *CreateClusterReleaseResponse, err error) {
	if request == nil {
		request = NewCreateClusterReleaseRequest()
	}
	response = NewCreateClusterReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductVersionDetailsRequest() (request *DescribeProductVersionDetailsRequest) {
	request = &DescribeProductVersionDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeProductVersionDetails")
	return
}

func NewDescribeProductVersionDetailsResponse() (response *DescribeProductVersionDetailsResponse) {
	response = &DescribeProductVersionDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看应用市场中某个应用版本详情
func (c *Client) DescribeProductVersionDetails(request *DescribeProductVersionDetailsRequest) (response *DescribeProductVersionDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeProductVersionDetailsRequest()
	}
	response = NewDescribeProductVersionDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeEtcdInstanceRequest() (request *UpgradeEtcdInstanceRequest) {
	request = &UpgradeEtcdInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeEtcdInstance")
	return
}

func NewUpgradeEtcdInstanceResponse() (response *UpgradeEtcdInstanceResponse) {
	response = &UpgradeEtcdInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级etcd实例
func (c *Client) UpgradeEtcdInstance(request *UpgradeEtcdInstanceRequest) (response *UpgradeEtcdInstanceResponse, err error) {
	if request == nil {
		request = NewUpgradeEtcdInstanceRequest()
	}
	response = NewUpgradeEtcdInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallLogAgentRequest() (request *UninstallLogAgentRequest) {
	request = &UninstallLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallLogAgent")
	return
}

func NewUninstallLogAgentResponse() (response *UninstallLogAgentResponse) {
	response = &UninstallLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从TKE集群中卸载CLS日志采集组件
func (c *Client) UninstallLogAgent(request *UninstallLogAgentRequest) (response *UninstallLogAgentResponse, err error) {
	if request == nil {
		request = NewUninstallLogAgentRequest()
	}
	response = NewUninstallLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEKSDiskRequest() (request *CreateEKSDiskRequest) {
	request = &CreateEKSDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEKSDisk")
	return
}

func NewCreateEKSDiskResponse() (response *CreateEKSDiskResponse) {
	response = &CreateEKSDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用eks-server创建CBS盘
func (c *Client) CreateEKSDisk(request *CreateEKSDiskRequest) (response *CreateEKSDiskResponse, err error) {
	if request == nil {
		request = NewCreateEKSDiskRequest()
	}
	response = NewCreateEKSDiskResponse()
	err = c.Send(request, response)
	return
}

func NewCollectAllCoreRequest() (request *CollectAllCoreRequest) {
	request = &CollectAllCoreRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CollectAllCore")
	return
}

func NewCollectAllCoreResponse() (response *CollectAllCoreResponse) {
	response = &CollectAllCoreResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// license平台定时采集当前TKE管控集群所使用的cpu总核数
func (c *Client) CollectAllCore(request *CollectAllCoreRequest) (response *CollectAllCoreResponse, err error) {
	if request == nil {
		request = NewCollectAllCoreRequest()
	}
	response = NewCollectAllCoreResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAggregationDataRequest() (request *DescribeAggregationDataRequest) {
	request = &DescribeAggregationDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAggregationData")
	return
}

func NewDescribeAggregationDataResponse() (response *DescribeAggregationDataResponse) {
	response = &DescribeAggregationDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照指定的维度来聚合某类数据，比如：根据云产品实例来聚合成本
func (c *Client) DescribeAggregationData(request *DescribeAggregationDataRequest) (response *DescribeAggregationDataResponse, err error) {
	if request == nil {
		request = NewDescribeAggregationDataRequest()
	}
	response = NewDescribeAggregationDataResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTKEEdgeClusterRequest() (request *UpdateTKEEdgeClusterRequest) {
	request = &UpdateTKEEdgeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateTKEEdgeCluster")
	return
}

func NewUpdateTKEEdgeClusterResponse() (response *UpdateTKEEdgeClusterResponse) {
	response = &UpdateTKEEdgeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改边缘计算集群名称等属性
func (c *Client) UpdateTKEEdgeCluster(request *UpdateTKEEdgeClusterRequest) (response *UpdateTKEEdgeClusterResponse, err error) {
	if request == nil {
		request = NewUpdateTKEEdgeClusterRequest()
	}
	response = NewUpdateTKEEdgeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReservedInstancesRequest() (request *DeleteReservedInstancesRequest) {
	request = &DeleteReservedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteReservedInstances")
	return
}

func NewDeleteReservedInstancesResponse() (response *DeleteReservedInstancesResponse) {
	response = &DeleteReservedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预留券实例如符合退还规则，可通过本接口主动退还。
func (c *Client) DeleteReservedInstances(request *DeleteReservedInstancesRequest) (response *DeleteReservedInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteReservedInstancesRequest()
	}
	response = NewDeleteReservedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
	request = &DescribeResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeResourceUsage")
	return
}

func NewDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
	response = &DescribeResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群资源使用量
func (c *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeResourceUsageRequest()
	}
	response = NewDescribeResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateECMInstancesRequest() (request *CreateECMInstancesRequest) {
	request = &CreateECMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateECMInstances")
	return
}

func NewCreateECMInstancesResponse() (response *CreateECMInstancesResponse) {
	response = &CreateECMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建边缘计算ECM机器
func (c *Client) CreateECMInstances(request *CreateECMInstancesRequest) (response *CreateECMInstancesResponse, err error) {
	if request == nil {
		request = NewCreateECMInstancesRequest()
	}
	response = NewCreateECMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewListEKSRegionRequest() (request *ListEKSRegionRequest) {
	request = &ListEKSRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListEKSRegion")
	return
}

func NewListEKSRegionResponse() (response *ListEKSRegionResponse) {
	response = &ListEKSRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询eks可用地域
func (c *Client) ListEKSRegion(request *ListEKSRegionRequest) (response *ListEKSRegionResponse, err error) {
	if request == nil {
		request = NewListEKSRegionRequest()
	}
	response = NewListEKSRegionResponse()
	err = c.Send(request, response)
	return
}

func NewGetTkeAppChartListRequest() (request *GetTkeAppChartListRequest) {
	request = &GetTkeAppChartListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetTkeAppChartList")
	return
}

func NewGetTkeAppChartListResponse() (response *GetTkeAppChartListResponse) {
	response = &GetTkeAppChartListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TKE支持的App列表
func (c *Client) GetTkeAppChartList(request *GetTkeAppChartListRequest) (response *GetTkeAppChartListResponse, err error) {
	if request == nil {
		request = NewGetTkeAppChartListRequest()
	}
	response = NewGetTkeAppChartListResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallTDCCLogAgentRequest() (request *UninstallTDCCLogAgentRequest) {
	request = &UninstallTDCCLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallTDCCLogAgent")
	return
}

func NewUninstallTDCCLogAgentResponse() (response *UninstallTDCCLogAgentResponse) {
	response = &UninstallTDCCLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载注册集群的日志采集组件
func (c *Client) UninstallTDCCLogAgent(request *UninstallTDCCLogAgentRequest) (response *UninstallTDCCLogAgentResponse, err error) {
	if request == nil {
		request = NewUninstallTDCCLogAgentRequest()
	}
	response = NewUninstallTDCCLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDrainEksNodeRequest() (request *DrainEksNodeRequest) {
	request = &DrainEksNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DrainEksNode")
	return
}

func NewDrainEksNodeResponse() (response *DrainEksNodeResponse) {
	response = &DrainEksNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐弹性集群的节点
func (c *Client) DrainEksNode(request *DrainEksNodeRequest) (response *DrainEksNodeResponse, err error) {
	if request == nil {
		request = NewDrainEksNodeRequest()
	}
	response = NewDrainEksNodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetZoneResourceRequest() (request *GetZoneResourceRequest) {
	request = &GetZoneResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetZoneResource")
	return
}

func NewGetZoneResourceResponse() (response *GetZoneResourceResponse) {
	response = &GetZoneResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用区是否有可用资源
func (c *Client) GetZoneResource(request *GetZoneResourceRequest) (response *GetZoneResourceResponse, err error) {
	if request == nil {
		request = NewGetZoneResourceRequest()
	}
	response = NewGetZoneResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEdgeEventPersistenceRequest() (request *DisableEdgeEventPersistenceRequest) {
	request = &DisableEdgeEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEdgeEventPersistence")
	return
}

func NewDisableEdgeEventPersistenceResponse() (response *DisableEdgeEventPersistenceResponse) {
	response = &DisableEdgeEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭事件持久化功能
func (c *Client) DisableEdgeEventPersistence(request *DisableEdgeEventPersistenceRequest) (response *DisableEdgeEventPersistenceResponse, err error) {
	if request == nil {
		request = NewDisableEdgeEventPersistenceRequest()
	}
	response = NewDisableEdgeEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveNodeFromNodePoolRequest() (request *RemoveNodeFromNodePoolRequest) {
	request = &RemoveNodeFromNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RemoveNodeFromNodePool")
	return
}

func NewRemoveNodeFromNodePoolResponse() (response *RemoveNodeFromNodePoolResponse) {
	response = &RemoveNodeFromNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移出节点池节点，但保留在集群内
func (c *Client) RemoveNodeFromNodePool(request *RemoveNodeFromNodePoolRequest) (response *RemoveNodeFromNodePoolResponse, err error) {
	if request == nil {
		request = NewRemoveNodeFromNodePoolRequest()
	}
	response = NewRemoveNodeFromNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEKSInstancesRequest() (request *DeleteEKSInstancesRequest) {
	request = &DeleteEKSInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSInstances")
	return
}

func NewDeleteEKSInstancesResponse() (response *DeleteEKSInstancesResponse) {
	response = &DeleteEKSInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性容器实例，可批量删除
func (c *Client) DeleteEKSInstances(request *DeleteEKSInstancesRequest) (response *DeleteEKSInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteEKSInstancesRequest()
	}
	response = NewDeleteEKSInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusRecordRuleYamlRequest() (request *DeletePrometheusRecordRuleYamlRequest) {
	request = &DeletePrometheusRecordRuleYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusRecordRuleYaml")
	return
}

func NewDeletePrometheusRecordRuleYamlResponse() (response *DeletePrometheusRecordRuleYamlResponse) {
	response = &DeletePrometheusRecordRuleYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除聚合规则
func (c *Client) DeletePrometheusRecordRuleYaml(request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusRecordRuleYamlRequest()
	}
	response = NewDeletePrometheusRecordRuleYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableMajorVersionRequest() (request *DescribeAvailableMajorVersionRequest) {
	request = &DescribeAvailableMajorVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableMajorVersion")
	return
}

func NewDescribeAvailableMajorVersionResponse() (response *DescribeAvailableMajorVersionResponse) {
	response = &DescribeAvailableMajorVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量获取集群是否可以大版本升级接口
func (c *Client) DescribeAvailableMajorVersion(request *DescribeAvailableMajorVersionRequest) (response *DescribeAvailableMajorVersionResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableMajorVersionRequest()
	}
	response = NewDescribeAvailableMajorVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeEniRequest() (request *DescribeEdgeEniRequest) {
	request = &DescribeEdgeEniRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeEni")
	return
}

func NewDescribeEdgeEniResponse() (response *DescribeEdgeEniResponse) {
	response = &DescribeEdgeEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘集群是否开启支持独立网卡
func (c *Client) DescribeEdgeEni(request *DescribeEdgeEniRequest) (response *DescribeEdgeEniResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeEniRequest()
	}
	response = NewDescribeEdgeEniResponse()
	err = c.Send(request, response)
	return
}

func NewRunClusterInspectionsRequest() (request *RunClusterInspectionsRequest) {
	request = &RunClusterInspectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RunClusterInspections")
	return
}

func NewRunClusterInspectionsResponse() (response *RunClusterInspectionsResponse) {
	response = &RunClusterInspectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发一次集群巡检。
func (c *Client) RunClusterInspections(request *RunClusterInspectionsRequest) (response *RunClusterInspectionsResponse, err error) {
	if request == nil {
		request = NewRunClusterInspectionsRequest()
	}
	response = NewRunClusterInspectionsResponse()
	err = c.Send(request, response)
	return
}

func NewAddTkeEdgeAlarmPolicyRequest() (request *AddTkeEdgeAlarmPolicyRequest) {
	request = &AddTkeEdgeAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddTkeEdgeAlarmPolicy")
	return
}

func NewAddTkeEdgeAlarmPolicyResponse() (response *AddTkeEdgeAlarmPolicyResponse) {
	response = &AddTkeEdgeAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加边缘集群告警策略
func (c *Client) AddTkeEdgeAlarmPolicy(request *AddTkeEdgeAlarmPolicyRequest) (response *AddTkeEdgeAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewAddTkeEdgeAlarmPolicyRequest()
	}
	response = NewAddTkeEdgeAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEncryptionProtectionRequest() (request *DisableEncryptionProtectionRequest) {
	request = &DisableEncryptionProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEncryptionProtection")
	return
}

func NewDisableEncryptionProtectionResponse() (response *DisableEncryptionProtectionResponse) {
	response = &DisableEncryptionProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭加密信息保护
func (c *Client) DisableEncryptionProtection(request *DisableEncryptionProtectionRequest) (response *DisableEncryptionProtectionResponse, err error) {
	if request == nil {
		request = NewDisableEncryptionProtectionRequest()
	}
	response = NewDisableEncryptionProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusTempSyncRequest() (request *DeletePrometheusTempSyncRequest) {
	request = &DeletePrometheusTempSyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTempSync")
	return
}

func NewDeletePrometheusTempSyncResponse() (response *DeletePrometheusTempSyncResponse) {
	response = &DeletePrometheusTempSyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除模板同步，这将会删除目标中该模板所生产的配置，针对V2版本实例
func (c *Client) DeletePrometheusTempSync(request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusTempSyncRequest()
	}
	response = NewDeletePrometheusTempSyncResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEKSNodeRequest() (request *DeleteEKSNodeRequest) {
	request = &DeleteEKSNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSNode")
	return
}

func NewDeleteEKSNodeResponse() (response *DeleteEKSNodeResponse) {
	response = &DeleteEKSNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除EKS节点
func (c *Client) DeleteEKSNode(request *DeleteEKSNodeRequest) (response *DeleteEKSNodeResponse, err error) {
	if request == nil {
		request = NewDeleteEKSNodeRequest()
	}
	response = NewDeleteEKSNodeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterAuthenticationOptionsRequest() (request *ModifyClusterAuthenticationOptionsRequest) {
	request = &ModifyClusterAuthenticationOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAuthenticationOptions")
	return
}

func NewModifyClusterAuthenticationOptionsResponse() (response *ModifyClusterAuthenticationOptionsResponse) {
	response = &ModifyClusterAuthenticationOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群认证配置
func (c *Client) ModifyClusterAuthenticationOptions(request *ModifyClusterAuthenticationOptionsRequest) (response *ModifyClusterAuthenticationOptionsResponse, err error) {
	if request == nil {
		request = NewModifyClusterAuthenticationOptionsRequest()
	}
	response = NewModifyClusterAuthenticationOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSClustersRequest() (request *DescribeEKSClustersRequest) {
	request = &DescribeEKSClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusters")
	return
}

func NewDescribeEKSClustersResponse() (response *DescribeEKSClustersResponse) {
	response = &DescribeEKSClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性集群列表
func (c *Client) DescribeEKSClusters(request *DescribeEKSClustersRequest) (response *DescribeEKSClustersResponse, err error) {
	if request == nil {
		request = NewDescribeEKSClustersRequest()
	}
	response = NewDescribeEKSClustersResponse()
	err = c.Send(request, response)
	return
}

func NewAttachEKSDisksRequest() (request *AttachEKSDisksRequest) {
	request = &AttachEKSDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AttachEKSDisks")
	return
}

func NewAttachEKSDisksResponse() (response *AttachEKSDisksResponse) {
	response = &AttachEKSDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过内部 CBS 接口绑定 CBS 隐藏盘到CXM上
func (c *Client) AttachEKSDisks(request *AttachEKSDisksRequest) (response *AttachEKSDisksResponse, err error) {
	if request == nil {
		request = NewAttachEKSDisksRequest()
	}
	response = NewAttachEKSDisksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirtualServiceRequest() (request *CreateVirtualServiceRequest) {
	request = &CreateVirtualServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateVirtualService")
	return
}

func NewCreateVirtualServiceResponse() (response *CreateVirtualServiceResponse) {
	response = &CreateVirtualServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建虚拟服务.
func (c *Client) CreateVirtualService(request *CreateVirtualServiceRequest) (response *CreateVirtualServiceResponse, err error) {
	if request == nil {
		request = NewCreateVirtualServiceRequest()
	}
	response = NewCreateVirtualServiceResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodByIdRequest() (request *GetPodByIdRequest) {
	request = &GetPodByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodById")
	return
}

func NewGetPodByIdResponse() (response *GetPodByIdResponse) {
	response = &GetPodByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过唯一 ID 查询Pod信息。
func (c *Client) GetPodById(request *GetPodByIdRequest) (response *GetPodByIdResponse, err error) {
	if request == nil {
		request = NewGetPodByIdRequest()
	}
	response = NewGetPodByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodesPriceRequest() (request *DescribeNodesPriceRequest) {
	request = &DescribeNodesPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeNodesPrice")
	return
}

func NewDescribeNodesPriceResponse() (response *DescribeNodesPriceResponse) {
	response = &DescribeNodesPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 tke 集群 node费用信息
func (c *Client) DescribeNodesPrice(request *DescribeNodesPriceRequest) (response *DescribeNodesPriceResponse, err error) {
	if request == nil {
		request = NewDescribeNodesPriceRequest()
	}
	response = NewDescribeNodesPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpotPodDetailRequest() (request *DescribeSpotPodDetailRequest) {
	request = &DescribeSpotPodDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeSpotPodDetail")
	return
}

func NewDescribeSpotPodDetailResponse() (response *DescribeSpotPodDetailResponse) {
	response = &DescribeSpotPodDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取删除的竞价pod详情
func (c *Client) DescribeSpotPodDetail(request *DescribeSpotPodDetailRequest) (response *DescribeSpotPodDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSpotPodDetailRequest()
	}
	response = NewDescribeSpotPodDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTDCCExternalClusterRequest() (request *UpdateTDCCExternalClusterRequest) {
	request = &UpdateTDCCExternalClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateTDCCExternalCluster")
	return
}

func NewUpdateTDCCExternalClusterResponse() (response *UpdateTDCCExternalClusterResponse) {
	response = &UpdateTDCCExternalClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改注册集群信息
func (c *Client) UpdateTDCCExternalCluster(request *UpdateTDCCExternalClusterRequest) (response *UpdateTDCCExternalClusterResponse, err error) {
	if request == nil {
		request = NewUpdateTDCCExternalClusterRequest()
	}
	response = NewUpdateTDCCExternalClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteService")
	return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Service
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRequest()
	}
	response = NewDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetaFeatureProgressRequest() (request *DescribeMetaFeatureProgressRequest) {
	request = &DescribeMetaFeatureProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeMetaFeatureProgress")
	return
}

func NewDescribeMetaFeatureProgressResponse() (response *DescribeMetaFeatureProgressResponse) {
	response = &DescribeMetaFeatureProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 跨租户特性相关接口，不能对公开发
func (c *Client) DescribeMetaFeatureProgress(request *DescribeMetaFeatureProgressRequest) (response *DescribeMetaFeatureProgressResponse, err error) {
	if request == nil {
		request = NewDescribeMetaFeatureProgressRequest()
	}
	response = NewDescribeMetaFeatureProgressResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterCertificatesRequest() (request *ListClusterCertificatesRequest) {
	request = &ListClusterCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListClusterCertificates")
	return
}

func NewListClusterCertificatesResponse() (response *ListClusterCertificatesResponse) {
	response = &ListClusterCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的证书列表，此接口接入CAM鉴权，只有管理权限的用户才可以正常调用，返回集群所有账户的证书列表
func (c *Client) ListClusterCertificates(request *ListClusterCertificatesRequest) (response *ListClusterCertificatesResponse, err error) {
	if request == nil {
		request = NewListClusterCertificatesRequest()
	}
	response = NewListClusterCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewSetNodePoolNodeProtectionRequest() (request *SetNodePoolNodeProtectionRequest) {
	request = &SetNodePoolNodeProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SetNodePoolNodeProtection")
	return
}

func NewSetNodePoolNodeProtectionResponse() (response *SetNodePoolNodeProtectionResponse) {
	response = &SetNodePoolNodeProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 仅能设置节点池中处于伸缩组的节点
func (c *Client) SetNodePoolNodeProtection(request *SetNodePoolNodeProtectionRequest) (response *SetNodePoolNodeProtectionResponse, err error) {
	if request == nil {
		request = NewSetNodePoolNodeProtectionRequest()
	}
	response = NewSetNodePoolNodeProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudRunLogSwitchesRequest() (request *DescribeCloudRunLogSwitchesRequest) {
	request = &DescribeCloudRunLogSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCloudRunLogSwitches")
	return
}

func NewDescribeCloudRunLogSwitchesResponse() (response *DescribeCloudRunLogSwitchesResponse) {
	response = &DescribeCloudRunLogSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cloudrun日志（审计、事件、普通日志）开关列表
func (c *Client) DescribeCloudRunLogSwitches(request *DescribeCloudRunLogSwitchesRequest) (response *DescribeCloudRunLogSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeCloudRunLogSwitchesRequest()
	}
	response = NewDescribeCloudRunLogSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpenPolicyListRequest() (request *ModifyOpenPolicyListRequest) {
	request = &ModifyOpenPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyOpenPolicyList")
	return
}

func NewModifyOpenPolicyListResponse() (response *ModifyOpenPolicyListResponse) {
	response = &ModifyOpenPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改opa策略
func (c *Client) ModifyOpenPolicyList(request *ModifyOpenPolicyListRequest) (response *ModifyOpenPolicyListResponse, err error) {
	if request == nil {
		request = NewModifyOpenPolicyListRequest()
	}
	response = NewModifyOpenPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDisableEdgeClusterAuditRequest() (request *DisableEdgeClusterAuditRequest) {
	request = &DisableEdgeClusterAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableEdgeClusterAudit")
	return
}

func NewDisableEdgeClusterAuditResponse() (response *DisableEdgeClusterAuditResponse) {
	response = &DisableEdgeClusterAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭边缘集群审计
func (c *Client) DisableEdgeClusterAudit(request *DisableEdgeClusterAuditRequest) (response *DisableEdgeClusterAuditResponse, err error) {
	if request == nil {
		request = NewDisableEdgeClusterAuditRequest()
	}
	response = NewDisableEdgeClusterAuditResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterInspectionResultsRequest() (request *ListClusterInspectionResultsRequest) {
	request = &ListClusterInspectionResultsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListClusterInspectionResults")
	return
}

func NewListClusterInspectionResultsResponse() (response *ListClusterInspectionResultsResponse) {
	response = &ListClusterInspectionResultsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定集群的巡检结果信息
func (c *Client) ListClusterInspectionResults(request *ListClusterInspectionResultsRequest) (response *ListClusterInspectionResultsResponse, err error) {
	if request == nil {
		request = NewListClusterInspectionResultsRequest()
	}
	response = NewListClusterInspectionResultsResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterInspectionResultsItemsRequest() (request *ListClusterInspectionResultsItemsRequest) {
	request = &ListClusterInspectionResultsItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListClusterInspectionResultsItems")
	return
}

func NewListClusterInspectionResultsItemsResponse() (response *ListClusterInspectionResultsItemsResponse) {
	response = &ListClusterInspectionResultsItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群巡检结果历史列表
func (c *Client) ListClusterInspectionResultsItems(request *ListClusterInspectionResultsItemsRequest) (response *ListClusterInspectionResultsItemsResponse, err error) {
	if request == nil {
		request = NewListClusterInspectionResultsItemsRequest()
	}
	response = NewListClusterInspectionResultsItemsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateECMPodRequest() (request *CreateECMPodRequest) {
	request = &CreateECMPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateECMPod")
	return
}

func NewCreateECMPodResponse() (response *CreateECMPodResponse) {
	response = &CreateECMPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ECM Pod
func (c *Client) CreateECMPod(request *CreateECMPodRequest) (response *CreateECMPodResponse, err error) {
	if request == nil {
		request = NewCreateECMPodRequest()
	}
	response = NewCreateECMPodResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeClusterProgressRequest() (request *GetUpgradeClusterProgressRequest) {
	request = &GetUpgradeClusterProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeClusterProgress")
	return
}

func NewGetUpgradeClusterProgressResponse() (response *GetUpgradeClusterProgressResponse) {
	response = &GetUpgradeClusterProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回当前集群升级进度
func (c *Client) GetUpgradeClusterProgress(request *GetUpgradeClusterProgressRequest) (response *GetUpgradeClusterProgressResponse, err error) {
	if request == nil {
		request = NewGetUpgradeClusterProgressRequest()
	}
	response = NewGetUpgradeClusterProgressResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusTempRequest() (request *ModifyPrometheusTempRequest) {
	request = &ModifyPrometheusTempRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusTemp")
	return
}

func NewModifyPrometheusTempResponse() (response *ModifyPrometheusTempResponse) {
	response = &ModifyPrometheusTempResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改模板内容
func (c *Client) ModifyPrometheusTemp(request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusTempRequest()
	}
	response = NewModifyPrometheusTempResponse()
	err = c.Send(request, response)
	return
}

func NewEnableExternalNodeSupportRequest() (request *EnableExternalNodeSupportRequest) {
	request = &EnableExternalNodeSupportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableExternalNodeSupport")
	return
}

func NewEnableExternalNodeSupportResponse() (response *EnableExternalNodeSupportResponse) {
	response = &EnableExternalNodeSupportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启第三方节点池支持
func (c *Client) EnableExternalNodeSupport(request *EnableExternalNodeSupportRequest) (response *EnableExternalNodeSupportResponse, err error) {
	if request == nil {
		request = NewEnableExternalNodeSupportRequest()
	}
	response = NewEnableExternalNodeSupportResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallEdgeLogAgentRequest() (request *UninstallEdgeLogAgentRequest) {
	request = &UninstallEdgeLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallEdgeLogAgent")
	return
}

func NewUninstallEdgeLogAgentResponse() (response *UninstallEdgeLogAgentResponse) {
	response = &UninstallEdgeLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从tke@edge集群边缘节点上卸载日志采集组件
func (c *Client) UninstallEdgeLogAgent(request *UninstallEdgeLogAgentRequest) (response *UninstallEdgeLogAgentResponse, err error) {
	if request == nil {
		request = NewUninstallEdgeLogAgentRequest()
	}
	response = NewUninstallEdgeLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageCachesRequest() (request *DescribeImageCachesRequest) {
	request = &DescribeImageCachesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeImageCaches")
	return
}

func NewDescribeImageCachesResponse() (response *DescribeImageCachesResponse) {
	response = &DescribeImageCachesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像缓存信息接口
func (c *Client) DescribeImageCaches(request *DescribeImageCachesRequest) (response *DescribeImageCachesResponse, err error) {
	if request == nil {
		request = NewDescribeImageCachesRequest()
	}
	response = NewDescribeImageCachesResponse()
	err = c.Send(request, response)
	return
}

func NewGetPrometheusTransMappingRequest() (request *GetPrometheusTransMappingRequest) {
	request = &GetPrometheusTransMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPrometheusTransMapping")
	return
}

func NewGetPrometheusTransMappingResponse() (response *GetPrometheusTransMappingResponse) {
	response = &GetPrometheusTransMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TPS迁移至TMP映射关系
func (c *Client) GetPrometheusTransMapping(request *GetPrometheusTransMappingRequest) (response *GetPrometheusTransMappingResponse, err error) {
	if request == nil {
		request = NewGetPrometheusTransMappingRequest()
	}
	response = NewGetPrometheusTransMappingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterGlobalStatisticsRequest() (request *DescribeClusterGlobalStatisticsRequest) {
	request = &DescribeClusterGlobalStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterGlobalStatistics")
	return
}

func NewDescribeClusterGlobalStatisticsResponse() (response *DescribeClusterGlobalStatisticsResponse) {
	response = &DescribeClusterGlobalStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有地域集群数量统计
func (c *Client) DescribeClusterGlobalStatistics(request *DescribeClusterGlobalStatisticsRequest) (response *DescribeClusterGlobalStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterGlobalStatisticsRequest()
	}
	response = NewDescribeClusterGlobalStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeUnitRequest() (request *DescribeNodeUnitRequest) {
	request = &DescribeNodeUnitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeNodeUnit")
	return
}

func NewDescribeNodeUnitResponse() (response *DescribeNodeUnitResponse) {
	response = &DescribeNodeUnitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取NodeUnit信息
func (c *Client) DescribeNodeUnit(request *DescribeNodeUnitRequest) (response *DescribeNodeUnitResponse, err error) {
	if request == nil {
		request = NewDescribeNodeUnitRequest()
	}
	response = NewDescribeNodeUnitResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEKSClusterRequest() (request *UpdateEKSClusterRequest) {
	request = &UpdateEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSCluster")
	return
}

func NewUpdateEKSClusterResponse() (response *UpdateEKSClusterResponse) {
	response = &UpdateEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改弹性集群名称等属性
func (c *Client) UpdateEKSCluster(request *UpdateEKSClusterRequest) (response *UpdateEKSClusterResponse, err error) {
	if request == nil {
		request = NewUpdateEKSClusterRequest()
	}
	response = NewUpdateEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteContainerProbeResultRequest() (request *DeleteContainerProbeResultRequest) {
	request = &DeleteContainerProbeResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteContainerProbeResult")
	return
}

func NewDeleteContainerProbeResultResponse() (response *DeleteContainerProbeResultResponse) {
	response = &DeleteContainerProbeResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteContainerProbeResult
func (c *Client) DeleteContainerProbeResult(request *DeleteContainerProbeResultRequest) (response *DeleteContainerProbeResultResponse, err error) {
	if request == nil {
		request = NewDeleteContainerProbeResultRequest()
	}
	response = NewDeleteContainerProbeResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterResourceDetailsRequest() (request *DescribeClusterResourceDetailsRequest) {
	request = &DescribeClusterResourceDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterResourceDetails")
	return
}

func NewDescribeClusterResourceDetailsResponse() (response *DescribeClusterResourceDetailsResponse) {
	response = &DescribeClusterResourceDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群对应资源的详细信息
func (c *Client) DescribeClusterResourceDetails(request *DescribeClusterResourceDetailsRequest) (response *DescribeClusterResourceDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterResourceDetailsRequest()
	}
	response = NewDescribeClusterResourceDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterAttributeRequest() (request *ModifyClusterAttributeRequest) {
	request = &ModifyClusterAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAttribute")
	return
}

func NewModifyClusterAttributeResponse() (response *ModifyClusterAttributeResponse) {
	response = &ModifyClusterAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群属性
func (c *Client) ModifyClusterAttribute(request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
	if request == nil {
		request = NewModifyClusterAttributeRequest()
	}
	response = NewModifyClusterAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewAddFromExistedCvmRequest() (request *AddFromExistedCvmRequest) {
	request = &AddFromExistedCvmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddFromExistedCvm")
	return
}

func NewAddFromExistedCvmResponse() (response *AddFromExistedCvmResponse) {
	response = &AddFromExistedCvmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将已经存在云服务器加入集群
func (c *Client) AddFromExistedCvm(request *AddFromExistedCvmRequest) (response *AddFromExistedCvmResponse, err error) {
	if request == nil {
		request = NewAddFromExistedCvmRequest()
	}
	response = NewAddFromExistedCvmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudRunHPARequest() (request *DescribeCloudRunHPARequest) {
	request = &DescribeCloudRunHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCloudRunHPA")
	return
}

func NewDescribeCloudRunHPAResponse() (response *DescribeCloudRunHPAResponse) {
	response = &DescribeCloudRunHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cloudrun hpa
func (c *Client) DescribeCloudRunHPA(request *DescribeCloudRunHPARequest) (response *DescribeCloudRunHPAResponse, err error) {
	if request == nil {
		request = NewDescribeCloudRunHPARequest()
	}
	response = NewDescribeCloudRunHPAResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusCreatingProgressRequest() (request *DescribePrometheusCreatingProgressRequest) {
	request = &DescribePrometheusCreatingProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusCreatingProgress")
	return
}

func NewDescribePrometheusCreatingProgressResponse() (response *DescribePrometheusCreatingProgressResponse) {
	response = &DescribePrometheusCreatingProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例创建进度
func (c *Client) DescribePrometheusCreatingProgress(request *DescribePrometheusCreatingProgressRequest) (response *DescribePrometheusCreatingProgressResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusCreatingProgressRequest()
	}
	response = NewDescribePrometheusCreatingProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHubClustersRequest() (request *DescribeHubClustersRequest) {
	request = &DescribeHubClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHubClusters")
	return
}

func NewDescribeHubClustersResponse() (response *DescribeHubClustersResponse) {
	response = &DescribeHubClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Hub集群列表
func (c *Client) DescribeHubClusters(request *DescribeHubClustersRequest) (response *DescribeHubClustersResponse, err error) {
	if request == nil {
		request = NewDescribeHubClustersRequest()
	}
	response = NewDescribeHubClustersResponse()
	err = c.Send(request, response)
	return
}

func NewForwardEKSApplicationRequestV3Request() (request *ForwardEKSApplicationRequestV3Request) {
	request = &ForwardEKSApplicationRequestV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardEKSApplicationRequestV3")
	return
}

func NewForwardEKSApplicationRequestV3Response() (response *ForwardEKSApplicationRequestV3Response) {
	response = &ForwardEKSApplicationRequestV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作EKS集群的组件
func (c *Client) ForwardEKSApplicationRequestV3(request *ForwardEKSApplicationRequestV3Request) (response *ForwardEKSApplicationRequestV3Response, err error) {
	if request == nil {
		request = NewForwardEKSApplicationRequestV3Request()
	}
	response = NewForwardEKSApplicationRequestV3Response()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteCluster")
	return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群(YUNAPI V3版本)
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	response = NewDeleteClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusTemplateRequest() (request *DeletePrometheusTemplateRequest) {
	request = &DeletePrometheusTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemplate")
	return
}

func NewDeletePrometheusTemplateResponse() (response *DeletePrometheusTemplateResponse) {
	response = &DeletePrometheusTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除一个云原生Prometheus配置模板
func (c *Client) DeletePrometheusTemplate(request *DeletePrometheusTemplateRequest) (response *DeletePrometheusTemplateResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusTemplateRequest()
	}
	response = NewDeletePrometheusTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetEksClusterUsedRequest() (request *GetEksClusterUsedRequest) {
	request = &GetEksClusterUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEksClusterUsed")
	return
}

func NewGetEksClusterUsedResponse() (response *GetEksClusterUsedResponse) {
	response = &GetEksClusterUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性集群配额和已创建集群数
func (c *Client) GetEksClusterUsed(request *GetEksClusterUsedRequest) (response *GetEksClusterUsedResponse, err error) {
	if request == nil {
		request = NewGetEksClusterUsedRequest()
	}
	response = NewGetEksClusterUsedResponse()
	err = c.Send(request, response)
	return
}

func NewInstallTDCCLogAgentRequest() (request *InstallTDCCLogAgentRequest) {
	request = &InstallTDCCLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallTDCCLogAgent")
	return
}

func NewInstallTDCCLogAgentResponse() (response *InstallTDCCLogAgentResponse) {
	response = &InstallTDCCLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在注册集群中安装CLS日志采集组件
func (c *Client) InstallTDCCLogAgent(request *InstallTDCCLogAgentRequest) (response *InstallTDCCLogAgentResponse, err error) {
	if request == nil {
		request = NewInstallTDCCLogAgentRequest()
	}
	response = NewInstallTDCCLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateECMEKSClusterRequest() (request *UpdateECMEKSClusterRequest) {
	request = &UpdateECMEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateECMEKSCluster")
	return
}

func NewUpdateECMEKSClusterResponse() (response *UpdateECMEKSClusterResponse) {
	response = &UpdateECMEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ECM弹性集群名称等属性
func (c *Client) UpdateECMEKSCluster(request *UpdateECMEKSClusterRequest) (response *UpdateECMEKSClusterResponse, err error) {
	if request == nil {
		request = NewUpdateECMEKSClusterRequest()
	}
	response = NewUpdateECMEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewInstallEksLogAgentRequest() (request *InstallEksLogAgentRequest) {
	request = &InstallEksLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallEksLogAgent")
	return
}

func NewInstallEksLogAgentResponse() (response *InstallEksLogAgentResponse) {
	response = &InstallEksLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在弹性集群中安装CLS日志采集组件
func (c *Client) InstallEksLogAgent(request *InstallEksLogAgentRequest) (response *InstallEksLogAgentResponse, err error) {
	if request == nil {
		request = NewInstallEksLogAgentRequest()
	}
	response = NewInstallEksLogAgentResponse()
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

func NewUpdateECMEKSClusterKubeconfigRequest() (request *UpdateECMEKSClusterKubeconfigRequest) {
	request = &UpdateECMEKSClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateECMEKSClusterKubeconfig")
	return
}

func NewUpdateECMEKSClusterKubeconfigResponse() (response *UpdateECMEKSClusterKubeconfigResponse) {
	response = &UpdateECMEKSClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新ECM EKS集群的kubeconfig信息
func (c *Client) UpdateECMEKSClusterKubeconfig(request *UpdateECMEKSClusterKubeconfigRequest) (response *UpdateECMEKSClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewUpdateECMEKSClusterKubeconfigRequest()
	}
	response = NewUpdateECMEKSClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEKSClusterKubeconfigRequest() (request *UpdateEKSClusterKubeconfigRequest) {
	request = &UpdateEKSClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSClusterKubeconfig")
	return
}

func NewUpdateEKSClusterKubeconfigResponse() (response *UpdateEKSClusterKubeconfigResponse) {
	response = &UpdateEKSClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新EKS集群的kubeconfig信息
func (c *Client) UpdateEKSClusterKubeconfig(request *UpdateEKSClusterKubeconfigRequest) (response *UpdateEKSClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewUpdateEKSClusterKubeconfigRequest()
	}
	response = NewUpdateEKSClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalNodeRequest() (request *DescribeExternalNodeRequest) {
	request = &DescribeExternalNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalNode")
	return
}

func NewDescribeExternalNodeResponse() (response *DescribeExternalNodeResponse) {
	response = &DescribeExternalNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看第三方节点列表
func (c *Client) DescribeExternalNode(request *DescribeExternalNodeRequest) (response *DescribeExternalNodeResponse, err error) {
	if request == nil {
		request = NewDescribeExternalNodeRequest()
	}
	response = NewDescribeExternalNodeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusRecordRuleYamlRequest() (request *DescribePrometheusRecordRuleYamlRequest) {
	request = &DescribePrometheusRecordRuleYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusRecordRuleYaml")
	return
}

func NewDescribePrometheusRecordRuleYamlResponse() (response *DescribePrometheusRecordRuleYamlResponse) {
	response = &DescribePrometheusRecordRuleYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合规则yaml列表
func (c *Client) DescribePrometheusRecordRuleYaml(request *DescribePrometheusRecordRuleYamlRequest) (response *DescribePrometheusRecordRuleYamlResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusRecordRuleYamlRequest()
	}
	response = NewDescribePrometheusRecordRuleYamlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEtcdSnapshotRequest() (request *CreateEtcdSnapshotRequest) {
	request = &CreateEtcdSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEtcdSnapshot")
	return
}

func NewCreateEtcdSnapshotResponse() (response *CreateEtcdSnapshotResponse) {
	response = &CreateEtcdSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建etcd快照
func (c *Client) CreateEtcdSnapshot(request *CreateEtcdSnapshotRequest) (response *CreateEtcdSnapshotResponse, err error) {
	if request == nil {
		request = NewCreateEtcdSnapshotRequest()
	}
	response = NewCreateEtcdSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusOverviewsRequest() (request *DescribePrometheusOverviewsRequest) {
	request = &DescribePrometheusOverviewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusOverviews")
	return
}

func NewDescribePrometheusOverviewsResponse() (response *DescribePrometheusOverviewsResponse) {
	response = &DescribePrometheusOverviewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例列表
func (c *Client) DescribePrometheusOverviews(request *DescribePrometheusOverviewsRequest) (response *DescribePrometheusOverviewsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusOverviewsRequest()
	}
	response = NewDescribePrometheusOverviewsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchModifyTagsStatusRequest() (request *DescribeBatchModifyTagsStatusRequest) {
	request = &DescribeBatchModifyTagsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeBatchModifyTagsStatus")
	return
}

func NewDescribeBatchModifyTagsStatusResponse() (response *DescribeBatchModifyTagsStatusResponse) {
	response = &DescribeBatchModifyTagsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询批量修改标签状态
func (c *Client) DescribeBatchModifyTagsStatus(request *DescribeBatchModifyTagsStatusRequest) (response *DescribeBatchModifyTagsStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBatchModifyTagsStatusRequest()
	}
	response = NewDescribeBatchModifyTagsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersCostQueryAggregatedRequest() (request *DescribeClustersCostQueryAggregatedRequest) {
	request = &DescribeClustersCostQueryAggregatedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersCostQueryAggregated")
	return
}

func NewDescribeClustersCostQueryAggregatedResponse() (response *DescribeClustersCostQueryAggregatedResponse) {
	response = &DescribeClustersCostQueryAggregatedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索集群成本相关聚合数据，比如按照 聚合维度，cluster、namespace、node、deployment、daemonset、statefulset、job、controller、pod、container 聚合后的数据，包括 cpu、ram使用效率等
func (c *Client) DescribeClustersCostQueryAggregated(request *DescribeClustersCostQueryAggregatedRequest) (response *DescribeClustersCostQueryAggregatedResponse, err error) {
	if request == nil {
		request = NewDescribeClustersCostQueryAggregatedRequest()
	}
	response = NewDescribeClustersCostQueryAggregatedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeInternalKubeconfigRequest() (request *DescribeTKEEdgeInternalKubeconfigRequest) {
	request = &DescribeTKEEdgeInternalKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeInternalKubeconfig")
	return
}

func NewDescribeTKEEdgeInternalKubeconfigResponse() (response *DescribeTKEEdgeInternalKubeconfigResponse) {
	response = &DescribeTKEEdgeInternalKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘计算内部访问kubeconfig
func (c *Client) DescribeTKEEdgeInternalKubeconfig(request *DescribeTKEEdgeInternalKubeconfigRequest) (response *DescribeTKEEdgeInternalKubeconfigResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeInternalKubeconfigRequest()
	}
	response = NewDescribeTKEEdgeInternalKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceQuotaRequest() (request *ModifyResourceQuotaRequest) {
	request = &ModifyResourceQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyResourceQuota")
	return
}

func NewModifyResourceQuotaResponse() (response *ModifyResourceQuotaResponse) {
	response = &ModifyResourceQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群资源配额
func (c *Client) ModifyResourceQuota(request *ModifyResourceQuotaRequest) (response *ModifyResourceQuotaResponse, err error) {
	if request == nil {
		request = NewModifyResourceQuotaRequest()
	}
	response = NewModifyResourceQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBackupStorageLocationRequest() (request *DeleteBackupStorageLocationRequest) {
	request = &DeleteBackupStorageLocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteBackupStorageLocation")
	return
}

func NewDeleteBackupStorageLocationResponse() (response *DeleteBackupStorageLocationResponse) {
	response = &DeleteBackupStorageLocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除备份仓库
func (c *Client) DeleteBackupStorageLocation(request *DeleteBackupStorageLocationRequest) (response *DeleteBackupStorageLocationResponse, err error) {
	if request == nil {
		request = NewDeleteBackupStorageLocationRequest()
	}
	response = NewDeleteBackupStorageLocationResponse()
	err = c.Send(request, response)
	return
}

func NewAddVpcCniSubnetsRequest() (request *AddVpcCniSubnetsRequest) {
	request = &AddVpcCniSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddVpcCniSubnets")
	return
}

func NewAddVpcCniSubnetsResponse() (response *AddVpcCniSubnetsResponse) {
	response = &AddVpcCniSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 针对VPC-CNI模式的集群，增加集群容器网络可使用的子网
func (c *Client) AddVpcCniSubnets(request *AddVpcCniSubnetsRequest) (response *AddVpcCniSubnetsResponse, err error) {
	if request == nil {
		request = NewAddVpcCniSubnetsRequest()
	}
	response = NewAddVpcCniSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateECMEKSClusterRequest() (request *CreateECMEKSClusterRequest) {
	request = &CreateECMEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateECMEKSCluster")
	return
}

func NewCreateECMEKSClusterResponse() (response *CreateECMEKSClusterResponse) {
	response = &CreateECMEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ECM弹性集群
func (c *Client) CreateECMEKSCluster(request *CreateECMEKSClusterRequest) (response *CreateECMEKSClusterResponse, err error) {
	if request == nil {
		request = NewCreateECMEKSClusterRequest()
	}
	response = NewCreateECMEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallEksLogAgentRequest() (request *UninstallEksLogAgentRequest) {
	request = &UninstallEksLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallEksLogAgent")
	return
}

func NewUninstallEksLogAgentResponse() (response *UninstallEksLogAgentResponse) {
	response = &UninstallEksLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载弹性集群的日志采集组件
func (c *Client) UninstallEksLogAgent(request *UninstallEksLogAgentRequest) (response *UninstallEksLogAgentResponse, err error) {
	if request == nil {
		request = NewUninstallEksLogAgentRequest()
	}
	response = NewUninstallEksLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusAgentExternalLabelsRequest() (request *ModifyPrometheusAgentExternalLabelsRequest) {
	request = &ModifyPrometheusAgentExternalLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAgentExternalLabels")
	return
}

func NewModifyPrometheusAgentExternalLabelsResponse() (response *ModifyPrometheusAgentExternalLabelsResponse) {
	response = &ModifyPrometheusAgentExternalLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改被关联集群的external labels
func (c *Client) ModifyPrometheusAgentExternalLabels(request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusAgentExternalLabelsRequest()
	}
	response = NewModifyPrometheusAgentExternalLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewAcquireClusterAdminRoleRequest() (request *AcquireClusterAdminRoleRequest) {
	request = &AcquireClusterAdminRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AcquireClusterAdminRole")
	return
}

func NewAcquireClusterAdminRoleResponse() (response *AcquireClusterAdminRoleResponse) {
	response = &AcquireClusterAdminRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过此接口，可以获取集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
func (c *Client) AcquireClusterAdminRole(request *AcquireClusterAdminRoleRequest) (response *AcquireClusterAdminRoleResponse, err error) {
	if request == nil {
		request = NewAcquireClusterAdminRoleRequest()
	}
	response = NewAcquireClusterAdminRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRIUtilizationDetailRequest() (request *DescribeRIUtilizationDetailRequest) {
	request = &DescribeRIUtilizationDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRIUtilizationDetail")
	return
}

func NewDescribeRIUtilizationDetailResponse() (response *DescribeRIUtilizationDetailResponse) {
	response = &DescribeRIUtilizationDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预留实例用量查询
func (c *Client) DescribeRIUtilizationDetail(request *DescribeRIUtilizationDetailRequest) (response *DescribeRIUtilizationDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRIUtilizationDetailRequest()
	}
	response = NewDescribeRIUtilizationDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallClustersCostsRequest() (request *UninstallClustersCostsRequest) {
	request = &UninstallClustersCostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallClustersCosts")
	return
}

func NewUninstallClustersCostsResponse() (response *UninstallClustersCostsResponse) {
	response = &UninstallClustersCostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载集群成本分析服务
func (c *Client) UninstallClustersCosts(request *UninstallClustersCostsRequest) (response *UninstallClustersCostsResponse, err error) {
	if request == nil {
		request = NewUninstallClustersCostsRequest()
	}
	response = NewUninstallClustersCostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusClusterAgentsRequest() (request *DescribePrometheusClusterAgentsRequest) {
	request = &DescribePrometheusClusterAgentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusClusterAgents")
	return
}

func NewDescribePrometheusClusterAgentsResponse() (response *DescribePrometheusClusterAgentsResponse) {
	response = &DescribePrometheusClusterAgentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TMP实例关联集群列表
func (c *Client) DescribePrometheusClusterAgents(request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusClusterAgentsRequest()
	}
	response = NewDescribePrometheusClusterAgentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOSImageIdRequest() (request *DescribeOSImageIdRequest) {
	request = &DescribeOSImageIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeOSImageId")
	return
}

func NewDescribeOSImageIdResponse() (response *DescribeOSImageIdResponse) {
	response = &DescribeOSImageIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像id
func (c *Client) DescribeOSImageId(request *DescribeOSImageIdRequest) (response *DescribeOSImageIdResponse, err error) {
	if request == nil {
		request = NewDescribeOSImageIdRequest()
	}
	response = NewDescribeOSImageIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirtualServiceReplicasRequest() (request *ModifyVirtualServiceReplicasRequest) {
	request = &ModifyVirtualServiceReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyVirtualServiceReplicas")
	return
}

func NewModifyVirtualServiceReplicasResponse() (response *ModifyVirtualServiceReplicasResponse) {
	response = &ModifyVirtualServiceReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改虚拟服务副本数
func (c *Client) ModifyVirtualServiceReplicas(request *ModifyVirtualServiceReplicasRequest) (response *ModifyVirtualServiceReplicasResponse, err error) {
	if request == nil {
		request = NewModifyVirtualServiceReplicasRequest()
	}
	response = NewModifyVirtualServiceReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExternalNodePoolRequest() (request *CreateExternalNodePoolRequest) {
	request = &CreateExternalNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateExternalNodePool")
	return
}

func NewCreateExternalNodePoolResponse() (response *CreateExternalNodePoolResponse) {
	response = &CreateExternalNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建第三方节点池
func (c *Client) CreateExternalNodePool(request *CreateExternalNodePoolRequest) (response *CreateExternalNodePoolResponse, err error) {
	if request == nil {
		request = NewCreateExternalNodePoolRequest()
	}
	response = NewCreateExternalNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewInstallLogAgentRequest() (request *InstallLogAgentRequest) {
	request = &InstallLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallLogAgent")
	return
}

func NewInstallLogAgentResponse() (response *InstallLogAgentResponse) {
	response = &InstallLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在TKE集群中安装CLS日志采集组件
func (c *Client) InstallLogAgent(request *InstallLogAgentRequest) (response *InstallLogAgentResponse, err error) {
	if request == nil {
		request = NewInstallLogAgentRequest()
	}
	response = NewInstallLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterExtraArgsRequest() (request *CheckClusterExtraArgsRequest) {
	request = &CheckClusterExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckClusterExtraArgs")
	return
}

func NewCheckClusterExtraArgsResponse() (response *CheckClusterExtraArgsResponse) {
	response = &CheckClusterExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查集群自定义参数是否合法
func (c *Client) CheckClusterExtraArgs(request *CheckClusterExtraArgsRequest) (response *CheckClusterExtraArgsResponse, err error) {
	if request == nil {
		request = NewCheckClusterExtraArgsRequest()
	}
	response = NewCheckClusterExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCosInfoRequest() (request *DescribeCosInfoRequest) {
	request = &DescribeCosInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCosInfo")
	return
}

func NewDescribeCosInfoResponse() (response *DescribeCosInfoResponse) {
	response = &DescribeCosInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取chart上传cos信息
func (c *Client) DescribeCosInfo(request *DescribeCosInfoRequest) (response *DescribeCosInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCosInfoRequest()
	}
	response = NewDescribeCosInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHelmChartVersionRequest() (request *DescribeHelmChartVersionRequest) {
	request = &DescribeHelmChartVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeHelmChartVersion")
	return
}

func NewDescribeHelmChartVersionResponse() (response *DescribeHelmChartVersionResponse) {
	response = &DescribeHelmChartVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取chart版本列表
func (c *Client) DescribeHelmChartVersion(request *DescribeHelmChartVersionRequest) (response *DescribeHelmChartVersionResponse, err error) {
	if request == nil {
		request = NewDescribeHelmChartVersionRequest()
	}
	response = NewDescribeHelmChartVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusGlobalConfigRequest() (request *DescribePrometheusGlobalConfigRequest) {
	request = &DescribePrometheusGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusGlobalConfig")
	return
}

func NewDescribePrometheusGlobalConfigResponse() (response *DescribePrometheusGlobalConfigResponse) {
	response = &DescribePrometheusGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得实例级别抓取配置
func (c *Client) DescribePrometheusGlobalConfig(request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusGlobalConfigRequest()
	}
	response = NewDescribePrometheusGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExternalNodePoolRequest() (request *ModifyExternalNodePoolRequest) {
	request = &ModifyExternalNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyExternalNodePool")
	return
}

func NewModifyExternalNodePoolResponse() (response *ModifyExternalNodePoolResponse) {
	response = &ModifyExternalNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改第三方节点池
func (c *Client) ModifyExternalNodePool(request *ModifyExternalNodePoolRequest) (response *ModifyExternalNodePoolResponse, err error) {
	if request == nil {
		request = NewModifyExternalNodePoolRequest()
	}
	response = NewModifyExternalNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusGlobalNotificationRequest() (request *CreatePrometheusGlobalNotificationRequest) {
	request = &CreatePrometheusGlobalNotificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusGlobalNotification")
	return
}

func NewCreatePrometheusGlobalNotificationResponse() (response *CreatePrometheusGlobalNotificationResponse) {
	response = &CreatePrometheusGlobalNotificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建全局告警通知渠道
func (c *Client) CreatePrometheusGlobalNotification(request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusGlobalNotificationRequest()
	}
	response = NewCreatePrometheusGlobalNotificationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmPoliciesRequest() (request *DeleteAlarmPoliciesRequest) {
	request = &DeleteAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteAlarmPolicies")
	return
}

func NewDeleteAlarmPoliciesResponse() (response *DeleteAlarmPoliciesResponse) {
	response = &DeleteAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警策略，支持批量删除
func (c *Client) DeleteAlarmPolicies(request *DeleteAlarmPoliciesRequest) (response *DeleteAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmPoliciesRequest()
	}
	response = NewDeleteAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncPrometheusTempRequest() (request *SyncPrometheusTempRequest) {
	request = &SyncPrometheusTempRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SyncPrometheusTemp")
	return
}

func NewSyncPrometheusTempResponse() (response *SyncPrometheusTempResponse) {
	response = &SyncPrometheusTempResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步模板到实例或者集群，针对V2版本实例
func (c *Client) SyncPrometheusTemp(request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
	if request == nil {
		request = NewSyncPrometheusTempRequest()
	}
	response = NewSyncPrometheusTempResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductVersionsRequest() (request *DescribeProductVersionsRequest) {
	request = &DescribeProductVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeProductVersions")
	return
}

func NewDescribeProductVersionsResponse() (response *DescribeProductVersionsResponse) {
	response = &DescribeProductVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个应用制品版本列表
func (c *Client) DescribeProductVersions(request *DescribeProductVersionsRequest) (response *DescribeProductVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeProductVersionsRequest()
	}
	response = NewDescribeProductVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewStopEksRequest() (request *StopEksRequest) {
	request = &StopEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "StopEks")
	return
}

func NewStopEksResponse() (response *StopEksResponse) {
	response = &StopEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对Pod关机操作
func (c *Client) StopEks(request *StopEksRequest) (response *StopEksResponse, err error) {
	if request == nil {
		request = NewStopEksRequest()
	}
	response = NewStopEksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTDCCExternalClusterRequest() (request *CreateTDCCExternalClusterRequest) {
	request = &CreateTDCCExternalClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateTDCCExternalCluster")
	return
}

func NewCreateTDCCExternalClusterResponse() (response *CreateTDCCExternalClusterResponse) {
	response = &CreateTDCCExternalClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TDCC注册集群
func (c *Client) CreateTDCCExternalCluster(request *CreateTDCCExternalClusterRequest) (response *CreateTDCCExternalClusterResponse, err error) {
	if request == nil {
		request = NewCreateTDCCExternalClusterRequest()
	}
	response = NewCreateTDCCExternalClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterRuntimeConfigRequest() (request *ModifyClusterRuntimeConfigRequest) {
	request = &ModifyClusterRuntimeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterRuntimeConfig")
	return
}

func NewModifyClusterRuntimeConfigResponse() (response *ModifyClusterRuntimeConfigResponse) {
	response = &ModifyClusterRuntimeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群及节点池纬度运行时配置
func (c *Client) ModifyClusterRuntimeConfig(request *ModifyClusterRuntimeConfigRequest) (response *ModifyClusterRuntimeConfigResponse, err error) {
	if request == nil {
		request = NewModifyClusterRuntimeConfigRequest()
	}
	response = NewModifyClusterRuntimeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUpdateNodeUnitRequest() (request *CreateUpdateNodeUnitRequest) {
	request = &CreateUpdateNodeUnitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateUpdateNodeUnit")
	return
}

func NewCreateUpdateNodeUnitResponse() (response *CreateUpdateNodeUnitResponse) {
	response = &CreateUpdateNodeUnitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建更新NodeUnit
func (c *Client) CreateUpdateNodeUnit(request *CreateUpdateNodeUnitRequest) (response *CreateUpdateNodeUnitResponse, err error) {
	if request == nil {
		request = NewCreateUpdateNodeUnitRequest()
	}
	response = NewCreateUpdateNodeUnitResponse()
	err = c.Send(request, response)
	return
}

func NewListECMEKSClusterCertificatesRequest() (request *ListECMEKSClusterCertificatesRequest) {
	request = &ListECMEKSClusterCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListECMEKSClusterCertificates")
	return
}

func NewListECMEKSClusterCertificatesResponse() (response *ListECMEKSClusterCertificatesResponse) {
	response = &ListECMEKSClusterCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的证书列表，此接口接入CAM鉴权，只有管理权限的用户才可以正常调用，返回集群所有账户的证书列表
func (c *Client) ListECMEKSClusterCertificates(request *ListECMEKSClusterCertificatesRequest) (response *ListECMEKSClusterCertificatesResponse, err error) {
	if request == nil {
		request = NewListECMEKSClusterCertificatesRequest()
	}
	response = NewListECMEKSClusterCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateImageCacheRequest() (request *UpdateImageCacheRequest) {
	request = &UpdateImageCacheRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateImageCache")
	return
}

func NewUpdateImageCacheResponse() (response *UpdateImageCacheResponse) {
	response = &UpdateImageCacheResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像缓存接口
func (c *Client) UpdateImageCache(request *UpdateImageCacheRequest) (response *UpdateImageCacheResponse, err error) {
	if request == nil {
		request = NewUpdateImageCacheRequest()
	}
	response = NewUpdateImageCacheResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyEtcdBackupInfoRequest() (request *VerifyEtcdBackupInfoRequest) {
	request = &VerifyEtcdBackupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "VerifyEtcdBackupInfo")
	return
}

func NewVerifyEtcdBackupInfoResponse() (response *VerifyEtcdBackupInfoResponse) {
	response = &VerifyEtcdBackupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验备份信息
func (c *Client) VerifyEtcdBackupInfo(request *VerifyEtcdBackupInfoRequest) (response *VerifyEtcdBackupInfoResponse, err error) {
	if request == nil {
		request = NewVerifyEtcdBackupInfoRequest()
	}
	response = NewVerifyEtcdBackupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeClusterReleaseRequest() (request *UpgradeClusterReleaseRequest) {
	request = &UpgradeClusterReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterRelease")
	return
}

func NewUpgradeClusterReleaseResponse() (response *UpgradeClusterReleaseResponse) {
	response = &UpgradeClusterReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级集群中已安装的应用
func (c *Client) UpgradeClusterRelease(request *UpgradeClusterReleaseRequest) (response *UpgradeClusterReleaseResponse, err error) {
	if request == nil {
		request = NewUpgradeClusterReleaseRequest()
	}
	response = NewUpgradeClusterReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterPodResourcesRequest() (request *DescribeClusterPodResourcesRequest) {
	request = &DescribeClusterPodResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterPodResources")
	return
}

func NewDescribeClusterPodResourcesResponse() (response *DescribeClusterPodResourcesResponse) {
	response = &DescribeClusterPodResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Pod资源配置信息，如：cpu-limit,cpu-request,mem-request,mem-limit
func (c *Client) DescribeClusterPodResources(request *DescribeClusterPodResourcesRequest) (response *DescribeClusterPodResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterPodResourcesRequest()
	}
	response = NewDescribeClusterPodResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewGetECMPodsRequest() (request *GetECMPodsRequest) {
	request = &GetECMPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetECMPods")
	return
}

func NewGetECMPodsResponse() (response *GetECMPodsResponse) {
	response = &GetECMPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群内ecm Pod
func (c *Client) GetECMPods(request *GetECMPodsRequest) (response *GetECMPodsResponse, err error) {
	if request == nil {
		request = NewGetECMPodsRequest()
	}
	response = NewGetECMPodsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterUpgradingStateRequest() (request *ModifyClusterUpgradingStateRequest) {
	request = &ModifyClusterUpgradingStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterUpgradingState")
	return
}

func NewModifyClusterUpgradingStateResponse() (response *ModifyClusterUpgradingStateResponse) {
	response = &ModifyClusterUpgradingStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 暂停或者取消集群升级
func (c *Client) ModifyClusterUpgradingState(request *ModifyClusterUpgradingStateRequest) (response *ModifyClusterUpgradingStateResponse, err error) {
	if request == nil {
		request = NewModifyClusterUpgradingStateRequest()
	}
	response = NewModifyClusterUpgradingStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusAlertPolicyRequest() (request *DeletePrometheusAlertPolicyRequest) {
	request = &DeletePrometheusAlertPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAlertPolicy")
	return
}

func NewDeletePrometheusAlertPolicyResponse() (response *DeletePrometheusAlertPolicyResponse) {
	response = &DeletePrometheusAlertPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除2.0实例告警策略
func (c *Client) DeletePrometheusAlertPolicy(request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusAlertPolicyRequest()
	}
	response = NewDeletePrometheusAlertPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEtcdConfigurationRequest() (request *ModifyEtcdConfigurationRequest) {
	request = &ModifyEtcdConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyEtcdConfiguration")
	return
}

func NewModifyEtcdConfigurationResponse() (response *ModifyEtcdConfigurationResponse) {
	response = &ModifyEtcdConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改etcd实例配置
func (c *Client) ModifyEtcdConfiguration(request *ModifyEtcdConfigurationRequest) (response *ModifyEtcdConfigurationResponse, err error) {
	if request == nil {
		request = NewModifyEtcdConfigurationRequest()
	}
	response = NewModifyEtcdConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneDiskQuotaRequest() (request *DescribeZoneDiskQuotaRequest) {
	request = &DescribeZoneDiskQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeZoneDiskQuota")
	return
}

func NewDescribeZoneDiskQuotaResponse() (response *DescribeZoneDiskQuotaResponse) {
	response = &DescribeZoneDiskQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群使用的系统盘类型及云硬盘在各个可用区的配额
func (c *Client) DescribeZoneDiskQuota(request *DescribeZoneDiskQuotaRequest) (response *DescribeZoneDiskQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeZoneDiskQuotaRequest()
	}
	response = NewDescribeZoneDiskQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterEndpointStatusRequest() (request *DescribeClusterEndpointStatusRequest) {
	request = &DescribeClusterEndpointStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointStatus")
	return
}

func NewDescribeClusterEndpointStatusResponse() (response *DescribeClusterEndpointStatusResponse) {
	response = &DescribeClusterEndpointStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群访问端口状态(独立集群开启内网/外网访问，托管集群支持开启内网访问)
func (c *Client) DescribeClusterEndpointStatus(request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterEndpointStatusRequest()
	}
	response = NewDescribeClusterEndpointStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeClusterResourcesRequest() (request *DescribeTKEEdgeClusterResourcesRequest) {
	request = &DescribeTKEEdgeClusterResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterResources")
	return
}

func NewDescribeTKEEdgeClusterResourcesResponse() (response *DescribeTKEEdgeClusterResourcesResponse) {
	response = &DescribeTKEEdgeClusterResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘计算节点资源详情
func (c *Client) DescribeTKEEdgeClusterResources(request *DescribeTKEEdgeClusterResourcesRequest) (response *DescribeTKEEdgeClusterResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeClusterResourcesRequest()
	}
	response = NewDescribeTKEEdgeClusterResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeLogAgentRequest() (request *UpgradeLogAgentRequest) {
	request = &UpgradeLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeLogAgent")
	return
}

func NewUpgradeLogAgentResponse() (response *UpgradeLogAgentResponse) {
	response = &UpgradeLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在TKE集群中升级CLS日志采集组件
func (c *Client) UpgradeLogAgent(request *UpgradeLogAgentRequest) (response *UpgradeLogAgentResponse, err error) {
	if request == nil {
		request = NewUpgradeLogAgentRequest()
	}
	response = NewUpgradeLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusRecordRuleRequest() (request *ModifyPrometheusRecordRuleRequest) {
	request = &ModifyPrometheusRecordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusRecordRule")
	return
}

func NewModifyPrometheusRecordRuleResponse() (response *ModifyPrometheusRecordRuleResponse) {
	response = &ModifyPrometheusRecordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改聚合规则
func (c *Client) ModifyPrometheusRecordRule(request *ModifyPrometheusRecordRuleRequest) (response *ModifyPrometheusRecordRuleResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusRecordRuleRequest()
	}
	response = NewModifyPrometheusRecordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewReservedResourceRequest() (request *ReservedResourceRequest) {
	request = &ReservedResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ReservedResource")
	return
}

func NewReservedResourceResponse() (response *ReservedResourceResponse) {
	response = &ReservedResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预留券保留资源或者删除保留的资源
func (c *Client) ReservedResource(request *ReservedResourceRequest) (response *ReservedResourceResponse, err error) {
	if request == nil {
		request = NewReservedResourceRequest()
	}
	response = NewReservedResourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExternalClusterInternalRequest() (request *ModifyExternalClusterInternalRequest) {
	request = &ModifyExternalClusterInternalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyExternalClusterInternal")
	return
}

func NewModifyExternalClusterInternalResponse() (response *ModifyExternalClusterInternalResponse) {
	response = &ModifyExternalClusterInternalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新第三方集群信息（内部调用）
func (c *Client) ModifyExternalClusterInternal(request *ModifyExternalClusterInternalRequest) (response *ModifyExternalClusterInternalResponse, err error) {
	if request == nil {
		request = NewModifyExternalClusterInternalRequest()
	}
	response = NewModifyExternalClusterInternalResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardIDRequest() (request *GetDashboardIDRequest) {
	request = &GetDashboardIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetDashboardID")
	return
}

func NewGetDashboardIDResponse() (response *GetDashboardIDResponse) {
	response = &GetDashboardIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据集群、仪表盘类型查询dashboardID信息
func (c *Client) GetDashboardID(request *GetDashboardIDRequest) (response *GetDashboardIDResponse, err error) {
	if request == nil {
		request = NewGetDashboardIDRequest()
	}
	response = NewGetDashboardIDResponse()
	err = c.Send(request, response)
	return
}

func NewAcquireClusterKubeConfigForProductRequest() (request *AcquireClusterKubeConfigForProductRequest) {
	request = &AcquireClusterKubeConfigForProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AcquireClusterKubeConfigForProduct")
	return
}

func NewAcquireClusterKubeConfigForProductResponse() (response *AcquireClusterKubeConfigForProductResponse) {
	response = &AcquireClusterKubeConfigForProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云产品获取集群权限，仅支持服务角色调用，必须在TKE完成注册登记。普通用户无法调用
func (c *Client) AcquireClusterKubeConfigForProduct(request *AcquireClusterKubeConfigForProductRequest) (response *AcquireClusterKubeConfigForProductResponse, err error) {
	if request == nil {
		request = NewAcquireClusterKubeConfigForProductRequest()
	}
	response = NewAcquireClusterKubeConfigForProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcCniPodLimitsRequest() (request *DescribeVpcCniPodLimitsRequest) {
	request = &DescribeVpcCniPodLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeVpcCniPodLimits")
	return
}

func NewDescribeVpcCniPodLimitsResponse() (response *DescribeVpcCniPodLimitsResponse) {
	response = &DescribeVpcCniPodLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口查询当前用户和地域在指定可用区下的机型可支持的最大 TKE VPC-CNI 网络模式的 Pod 数量
func (c *Client) DescribeVpcCniPodLimits(request *DescribeVpcCniPodLimitsRequest) (response *DescribeVpcCniPodLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcCniPodLimitsRequest()
	}
	response = NewDescribeVpcCniPodLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEtcdSnapshotPolicyRequest() (request *ModifyEtcdSnapshotPolicyRequest) {
	request = &ModifyEtcdSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyEtcdSnapshotPolicy")
	return
}

func NewModifyEtcdSnapshotPolicyResponse() (response *ModifyEtcdSnapshotPolicyResponse) {
	response = &ModifyEtcdSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改etcd快照策略
func (c *Client) ModifyEtcdSnapshotPolicy(request *ModifyEtcdSnapshotPolicyRequest) (response *ModifyEtcdSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewModifyEtcdSnapshotPolicyRequest()
	}
	response = NewModifyEtcdSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusConfigRequest() (request *DescribePrometheusConfigRequest) {
	request = &DescribePrometheusConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusConfig")
	return
}

func NewDescribePrometheusConfigResponse() (response *DescribePrometheusConfigResponse) {
	response = &DescribePrometheusConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群采集配置
func (c *Client) DescribePrometheusConfig(request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusConfigRequest()
	}
	response = NewDescribePrometheusConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInspectionOverviewsRequest() (request *DescribeClusterInspectionOverviewsRequest) {
	request = &DescribeClusterInspectionOverviewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspectionOverviews")
	return
}

func NewDescribeClusterInspectionOverviewsResponse() (response *DescribeClusterInspectionOverviewsResponse) {
	response = &DescribeClusterInspectionOverviewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得集群健康检查列表
func (c *Client) DescribeClusterInspectionOverviews(request *DescribeClusterInspectionOverviewsRequest) (response *DescribeClusterInspectionOverviewsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInspectionOverviewsRequest()
	}
	response = NewDescribeClusterInspectionOverviewsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSContainerInstanceEventRequest() (request *DescribeEKSContainerInstanceEventRequest) {
	request = &DescribeEKSContainerInstanceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstanceEvent")
	return
}

func NewDescribeEKSContainerInstanceEventResponse() (response *DescribeEKSContainerInstanceEventResponse) {
	response = &DescribeEKSContainerInstanceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器实例的事件
func (c *Client) DescribeEKSContainerInstanceEvent(request *DescribeEKSContainerInstanceEventRequest) (response *DescribeEKSContainerInstanceEventResponse, err error) {
	if request == nil {
		request = NewDescribeEKSContainerInstanceEventRequest()
	}
	response = NewDescribeEKSContainerInstanceEventResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNodePoolDesiredCapacityAboutAsgRequest() (request *ModifyNodePoolDesiredCapacityAboutAsgRequest) {
	request = &ModifyNodePoolDesiredCapacityAboutAsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePoolDesiredCapacityAboutAsg")
	return
}

func NewModifyNodePoolDesiredCapacityAboutAsgResponse() (response *ModifyNodePoolDesiredCapacityAboutAsgResponse) {
	response = &ModifyNodePoolDesiredCapacityAboutAsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改节点池关联伸缩组的期望实例数
func (c *Client) ModifyNodePoolDesiredCapacityAboutAsg(request *ModifyNodePoolDesiredCapacityAboutAsgRequest) (response *ModifyNodePoolDesiredCapacityAboutAsgResponse, err error) {
	if request == nil {
		request = NewModifyNodePoolDesiredCapacityAboutAsgRequest()
	}
	response = NewModifyNodePoolDesiredCapacityAboutAsgResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterVirtualNodePoolRequest() (request *CreateClusterVirtualNodePoolRequest) {
	request = &CreateClusterVirtualNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNodePool")
	return
}

func NewCreateClusterVirtualNodePoolResponse() (response *CreateClusterVirtualNodePoolResponse) {
	response = &CreateClusterVirtualNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建超级节点池
func (c *Client) CreateClusterVirtualNodePool(request *CreateClusterVirtualNodePoolRequest) (response *CreateClusterVirtualNodePoolResponse, err error) {
	if request == nil {
		request = NewCreateClusterVirtualNodePoolRequest()
	}
	response = NewCreateClusterVirtualNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDrainClusterVirtualNodeRequest() (request *DrainClusterVirtualNodeRequest) {
	request = &DrainClusterVirtualNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DrainClusterVirtualNode")
	return
}

func NewDrainClusterVirtualNodeResponse() (response *DrainClusterVirtualNodeResponse) {
	response = &DrainClusterVirtualNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐超级节点
func (c *Client) DrainClusterVirtualNode(request *DrainClusterVirtualNodeRequest) (response *DrainClusterVirtualNodeResponse, err error) {
	if request == nil {
		request = NewDrainClusterVirtualNodeRequest()
	}
	response = NewDrainClusterVirtualNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusRecordRuleRequest() (request *DescribePrometheusRecordRuleRequest) {
	request = &DescribePrometheusRecordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusRecordRule")
	return
}

func NewDescribePrometheusRecordRuleResponse() (response *DescribePrometheusRecordRuleResponse) {
	response = &DescribePrometheusRecordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合规则列表
func (c *Client) DescribePrometheusRecordRule(request *DescribePrometheusRecordRuleRequest) (response *DescribePrometheusRecordRuleResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusRecordRuleRequest()
	}
	response = NewDescribePrometheusRecordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewForwardECMClusterRequestApiRequest() (request *ForwardECMClusterRequestApiRequest) {
	request = &ForwardECMClusterRequestApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardECMClusterRequestApi")
	return
}

func NewForwardECMClusterRequestApiResponse() (response *ForwardECMClusterRequestApiResponse) {
	response = &ForwardECMClusterRequestApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询、新增、删除、编辑ECM EKS集群内资源
func (c *Client) ForwardECMClusterRequestApi(request *ForwardECMClusterRequestApiRequest) (response *ForwardECMClusterRequestApiResponse, err error) {
	if request == nil {
		request = NewForwardECMClusterRequestApiRequest()
	}
	response = NewForwardECMClusterRequestApiResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterRuntimeConfigRequest() (request *CheckClusterRuntimeConfigRequest) {
	request = &CheckClusterRuntimeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckClusterRuntimeConfig")
	return
}

func NewCheckClusterRuntimeConfigResponse() (response *CheckClusterRuntimeConfigResponse) {
	response = &CheckClusterRuntimeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查集群运行时配置是否兼容目标K8S版本
func (c *Client) CheckClusterRuntimeConfig(request *CheckClusterRuntimeConfigRequest) (response *CheckClusterRuntimeConfigResponse, err error) {
	if request == nil {
		request = NewCheckClusterRuntimeConfigRequest()
	}
	response = NewCheckClusterRuntimeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusTempRequest() (request *CreatePrometheusTempRequest) {
	request = &CreatePrometheusTempRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusTemp")
	return
}

func NewCreatePrometheusTempResponse() (response *CreatePrometheusTempResponse) {
	response = &CreatePrometheusTempResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个云原生Prometheus模板
func (c *Client) CreatePrometheusTemp(request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusTempRequest()
	}
	response = NewCreatePrometheusTempResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterCIDRFromVbcRequest() (request *DeleteClusterCIDRFromVbcRequest) {
	request = &DeleteClusterCIDRFromVbcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterCIDRFromVbc")
	return
}

func NewDeleteClusterCIDRFromVbcResponse() (response *DeleteClusterCIDRFromVbcResponse) {
	response = &DeleteClusterCIDRFromVbcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从云联网删除集群CIDR
func (c *Client) DeleteClusterCIDRFromVbc(request *DeleteClusterCIDRFromVbcRequest) (response *DeleteClusterCIDRFromVbcResponse, err error) {
	if request == nil {
		request = NewDeleteClusterCIDRFromVbcRequest()
	}
	response = NewDeleteClusterCIDRFromVbcResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageRegistryCredentialsRequest() (request *DeleteImageRegistryCredentialsRequest) {
	request = &DeleteImageRegistryCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteImageRegistryCredentials")
	return
}

func NewDeleteImageRegistryCredentialsResponse() (response *DeleteImageRegistryCredentialsResponse) {
	response = &DeleteImageRegistryCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像仓库凭证接口
func (c *Client) DeleteImageRegistryCredentials(request *DeleteImageRegistryCredentialsRequest) (response *DeleteImageRegistryCredentialsResponse, err error) {
	if request == nil {
		request = NewDeleteImageRegistryCredentialsRequest()
	}
	response = NewDeleteImageRegistryCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeSSHRequest() (request *DescribeEdgeSSHRequest) {
	request = &DescribeEdgeSSHRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeSSH")
	return
}

func NewDescribeEdgeSSHResponse() (response *DescribeEdgeSSHResponse) {
	response = &DescribeEdgeSSHResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘集群SSH登录内网节点状态
func (c *Client) DescribeEdgeSSH(request *DescribeEdgeSSHRequest) (response *DescribeEdgeSSHResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeSSHRequest()
	}
	response = NewDescribeEdgeSSHResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterHostNameRequest() (request *CheckClusterHostNameRequest) {
	request = &CheckClusterHostNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckClusterHostName")
	return
}

func NewCheckClusterHostNameResponse() (response *CheckClusterHostNameResponse) {
	response = &CheckClusterHostNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查集群节点主机名称，判断节点主机名称是否符合规则，是否可以加入集群。
func (c *Client) CheckClusterHostName(request *CheckClusterHostNameRequest) (response *CheckClusterHostNameResponse, err error) {
	if request == nil {
		request = NewCheckClusterHostNameRequest()
	}
	response = NewCheckClusterHostNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeNodeScriptRequest() (request *DescribeEdgeNodeScriptRequest) {
	request = &DescribeEdgeNodeScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeNodeScript")
	return
}

func NewDescribeEdgeNodeScriptResponse() (response *DescribeEdgeNodeScriptResponse) {
	response = &DescribeEdgeNodeScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘脚本链接，此接口用于添加第三方节点，通过下载脚本从而将节点添加到边缘集群
func (c *Client) DescribeEdgeNodeScript(request *DescribeEdgeNodeScriptRequest) (response *DescribeEdgeNodeScriptResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeNodeScriptRequest()
	}
	response = NewDescribeEdgeNodeScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeClusterUpgradeInfoRequest() (request *DescribeEdgeClusterUpgradeInfoRequest) {
	request = &DescribeEdgeClusterUpgradeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterUpgradeInfo")
	return
}

func NewDescribeEdgeClusterUpgradeInfoResponse() (response *DescribeEdgeClusterUpgradeInfoResponse) {
	response = &DescribeEdgeClusterUpgradeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 可以查询边缘集群升级信息，包含可以升级的组件，当前升级状态和升级错误信息
func (c *Client) DescribeEdgeClusterUpgradeInfo(request *DescribeEdgeClusterUpgradeInfoRequest) (response *DescribeEdgeClusterUpgradeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeClusterUpgradeInfoRequest()
	}
	response = NewDescribeEdgeClusterUpgradeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSContainerInstanceRegionsRequest() (request *DescribeEKSContainerInstanceRegionsRequest) {
	request = &DescribeEKSContainerInstanceRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstanceRegions")
	return
}

func NewDescribeEKSContainerInstanceRegionsResponse() (response *DescribeEKSContainerInstanceRegionsResponse) {
	response = &DescribeEKSContainerInstanceRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器实例支持的地域
func (c *Client) DescribeEKSContainerInstanceRegions(request *DescribeEKSContainerInstanceRegionsRequest) (response *DescribeEKSContainerInstanceRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeEKSContainerInstanceRegionsRequest()
	}
	response = NewDescribeEKSContainerInstanceRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDrainEksClusterNodeRequest() (request *DrainEksClusterNodeRequest) {
	request = &DrainEksClusterNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DrainEksClusterNode")
	return
}

func NewDrainEksClusterNodeResponse() (response *DrainEksClusterNodeResponse) {
	response = &DrainEksClusterNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐弹性集群节点
func (c *Client) DrainEksClusterNode(request *DrainEksClusterNodeRequest) (response *DrainEksClusterNodeResponse, err error) {
	if request == nil {
		request = NewDrainEksClusterNodeRequest()
	}
	response = NewDrainEksClusterNodeResponse()
	err = c.Send(request, response)
	return
}

func NewPauseClusterInstancesRequest() (request *PauseClusterInstancesRequest) {
	request = &PauseClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "PauseClusterInstances")
	return
}

func NewPauseClusterInstancesResponse() (response *PauseClusterInstancesResponse) {
	response = &PauseClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 暂停集群节点升级, API 3.0
func (c *Client) PauseClusterInstances(request *PauseClusterInstancesRequest) (response *PauseClusterInstancesResponse, err error) {
	if request == nil {
		request = NewPauseClusterInstancesRequest()
	}
	response = NewPauseClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTempRequest() (request *DescribePrometheusTempRequest) {
	request = &DescribePrometheusTempRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemp")
	return
}

func NewDescribePrometheusTempResponse() (response *DescribePrometheusTempResponse) {
	response = &DescribePrometheusTempResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取模板列表，默认模板将总是在最前面
func (c *Client) DescribePrometheusTemp(request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTempRequest()
	}
	response = NewDescribePrometheusTempResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddonRequest() (request *DescribeAddonRequest) {
	request = &DescribeAddonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAddon")
	return
}

func NewDescribeAddonResponse() (response *DescribeAddonResponse) {
	response = &DescribeAddonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取addon列表
func (c *Client) DescribeAddon(request *DescribeAddonRequest) (response *DescribeAddonResponse, err error) {
	if request == nil {
		request = NewDescribeAddonRequest()
	}
	response = NewDescribeAddonResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubnetResourceRequest() (request *GetSubnetResourceRequest) {
	request = &GetSubnetResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetSubnetResource")
	return
}

func NewGetSubnetResourceResponse() (response *GetSubnetResourceResponse) {
	response = &GetSubnetResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子网资源
func (c *Client) GetSubnetResource(request *GetSubnetResourceRequest) (response *GetSubnetResourceResponse, err error) {
	if request == nil {
		request = NewGetSubnetResourceRequest()
	}
	response = NewGetSubnetResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEdgeCVMInstancesRequest() (request *DeleteEdgeCVMInstancesRequest) {
	request = &DeleteEdgeCVMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeCVMInstances")
	return
}

func NewDeleteEdgeCVMInstancesResponse() (response *DeleteEdgeCVMInstancesResponse) {
	response = &DeleteEdgeCVMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除边缘容器CVM实例
func (c *Client) DeleteEdgeCVMInstances(request *DeleteEdgeCVMInstancesRequest) (response *DeleteEdgeCVMInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteEdgeCVMInstancesRequest()
	}
	response = NewDeleteEdgeCVMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyK8sWorkloadRequest() (request *ModifyK8sWorkloadRequest) {
	request = &ModifyK8sWorkloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyK8sWorkload")
	return
}

func NewModifyK8sWorkloadResponse() (response *ModifyK8sWorkloadResponse) {
	response = &ModifyK8sWorkloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改工作负载的配置
func (c *Client) ModifyK8sWorkload(request *ModifyK8sWorkloadRequest) (response *ModifyK8sWorkloadResponse, err error) {
	if request == nil {
		request = NewModifyK8sWorkloadRequest()
	}
	response = NewModifyK8sWorkloadResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
	request = &DeleteClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterInstances")
	return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
	response = &DeleteClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群中的实例
func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteClusterInstancesRequest()
	}
	response = NewDeleteClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksAlarmPoliciesRequest() (request *DescribeEksAlarmPoliciesRequest) {
	request = &DescribeEksAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksAlarmPolicies")
	return
}

func NewDescribeEksAlarmPoliciesResponse() (response *DescribeEksAlarmPoliciesResponse) {
	response = &DescribeEksAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性集群的告警策略列表
func (c *Client) DescribeEksAlarmPolicies(request *DescribeEksAlarmPoliciesRequest) (response *DescribeEksAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeEksAlarmPoliciesRequest()
	}
	response = NewDescribeEksAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallClusterReleaseRequest() (request *UninstallClusterReleaseRequest) {
	request = &UninstallClusterReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UninstallClusterRelease")
	return
}

func NewUninstallClusterReleaseResponse() (response *UninstallClusterReleaseResponse) {
	response = &UninstallClusterReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在应用市场中集群删除某个应用
func (c *Client) UninstallClusterRelease(request *UninstallClusterReleaseRequest) (response *UninstallClusterReleaseResponse, err error) {
	if request == nil {
		request = NewUninstallClusterReleaseRequest()
	}
	response = NewUninstallClusterReleaseResponse()
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

func NewForwardClusterRequestRequest() (request *ForwardClusterRequestRequest) {
	request = &ForwardClusterRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardClusterRequest")
	return
}

func NewForwardClusterRequestResponse() (response *ForwardClusterRequestResponse) {
	response = &ForwardClusterRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询、新增、删除、编辑TKE集群内资源
func (c *Client) ForwardClusterRequest(request *ForwardClusterRequestRequest) (response *ForwardClusterRequestResponse, err error) {
	if request == nil {
		request = NewForwardClusterRequestRequest()
	}
	response = NewForwardClusterRequestResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusTemplateRequest() (request *CreatePrometheusTemplateRequest) {
	request = &CreatePrometheusTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusTemplate")
	return
}

func NewCreatePrometheusTemplateResponse() (response *CreatePrometheusTemplateResponse) {
	response = &CreatePrometheusTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个云原生Prometheus模板实例
func (c *Client) CreatePrometheusTemplate(request *CreatePrometheusTemplateRequest) (response *CreatePrometheusTemplateResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusTemplateRequest()
	}
	response = NewCreatePrometheusTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAsGroupsRequest() (request *DescribeClusterAsGroupsRequest) {
	request = &DescribeClusterAsGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroups")
	return
}

func NewDescribeClusterAsGroupsResponse() (response *DescribeClusterAsGroupsResponse) {
	response = &DescribeClusterAsGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群关联的伸缩组列表
func (c *Client) DescribeClusterAsGroups(request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAsGroupsRequest()
	}
	response = NewDescribeClusterAsGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusClusterAgentRequest() (request *CreatePrometheusClusterAgentRequest) {
	request = &CreatePrometheusClusterAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusClusterAgent")
	return
}

func NewCreatePrometheusClusterAgentResponse() (response *CreatePrometheusClusterAgentResponse) {
	response = &CreatePrometheusClusterAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 与云监控融合的2.0实例关联集群
func (c *Client) CreatePrometheusClusterAgent(request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusClusterAgentRequest()
	}
	response = NewCreatePrometheusClusterAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceCreateProgressRequest() (request *DescribeInstanceCreateProgressRequest) {
	request = &DescribeInstanceCreateProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeInstanceCreateProgress")
	return
}

func NewDescribeInstanceCreateProgressResponse() (response *DescribeInstanceCreateProgressResponse) {
	response = &DescribeInstanceCreateProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点创建进度
func (c *Client) DescribeInstanceCreateProgress(request *DescribeInstanceCreateProgressRequest) (response *DescribeInstanceCreateProgressResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceCreateProgressRequest()
	}
	response = NewDescribeInstanceCreateProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogSwitchesRequest() (request *DescribeLogSwitchesRequest) {
	request = &DescribeLogSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeLogSwitches")
	return
}

func NewDescribeLogSwitchesResponse() (response *DescribeLogSwitchesResponse) {
	response = &DescribeLogSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群日志（审计、事件、普通日志）开关列表
func (c *Client) DescribeLogSwitches(request *DescribeLogSwitchesRequest) (response *DescribeLogSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeLogSwitchesRequest()
	}
	response = NewDescribeLogSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReservedInstancesRequest() (request *DescribeReservedInstancesRequest) {
	request = &DescribeReservedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeReservedInstances")
	return
}

func NewDescribeReservedInstancesResponse() (response *DescribeReservedInstancesResponse) {
	response = &DescribeReservedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询预留实例列表
func (c *Client) DescribeReservedInstances(request *DescribeReservedInstancesRequest) (response *DescribeReservedInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeReservedInstancesRequest()
	}
	response = NewDescribeReservedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVersionsRequest() (request *DescribeVersionsRequest) {
	request = &DescribeVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeVersions")
	return
}

func NewDescribeVersionsResponse() (response *DescribeVersionsResponse) {
	response = &DescribeVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群版本信息
func (c *Client) DescribeVersions(request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeVersionsRequest()
	}
	response = NewDescribeVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusConfigRequest() (request *DeletePrometheusConfigRequest) {
	request = &DeletePrometheusConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusConfig")
	return
}

func NewDeletePrometheusConfigResponse() (response *DeletePrometheusConfigResponse) {
	response = &DeletePrometheusConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群采集配置
func (c *Client) DeletePrometheusConfig(request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusConfigRequest()
	}
	response = NewDeletePrometheusConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersCostCacheWarmedRequest() (request *DescribeClustersCostCacheWarmedRequest) {
	request = &DescribeClustersCostCacheWarmedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersCostCacheWarmed")
	return
}

func NewDescribeClustersCostCacheWarmedResponse() (response *DescribeClustersCostCacheWarmedResponse) {
	response = &DescribeClustersCostCacheWarmedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群的成本数据是否已经缓存
func (c *Client) DescribeClustersCostCacheWarmed(request *DescribeClustersCostCacheWarmedRequest) (response *DescribeClustersCostCacheWarmedResponse, err error) {
	if request == nil {
		request = NewDescribeClustersCostCacheWarmedRequest()
	}
	response = NewDescribeClustersCostCacheWarmedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPAMDRequest() (request *DescribeIPAMDRequest) {
	request = &DescribeIPAMDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeIPAMD")
	return
}

func NewDescribeIPAMDResponse() (response *DescribeIPAMDResponse) {
	response = &DescribeIPAMDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取eniipamd组件信息
func (c *Client) DescribeIPAMD(request *DescribeIPAMDRequest) (response *DescribeIPAMDResponse, err error) {
	if request == nil {
		request = NewDescribeIPAMDRequest()
	}
	response = NewDescribeIPAMDResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterNodePoolRequest() (request *CreateClusterNodePoolRequest) {
	request = &CreateClusterNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterNodePool")
	return
}

func NewCreateClusterNodePoolResponse() (response *CreateClusterNodePoolResponse) {
	response = &CreateClusterNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建节点池
func (c *Client) CreateClusterNodePool(request *CreateClusterNodePoolRequest) (response *CreateClusterNodePoolResponse, err error) {
	if request == nil {
		request = NewCreateClusterNodePoolRequest()
	}
	response = NewCreateClusterNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInspectionsRequest() (request *DescribeClusterInspectionsRequest) {
	request = &DescribeClusterInspectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspections")
	return
}

func NewDescribeClusterInspectionsResponse() (response *DescribeClusterInspectionsResponse) {
	response = &DescribeClusterInspectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群巡检概览
func (c *Client) DescribeClusterInspections(request *DescribeClusterInspectionsRequest) (response *DescribeClusterInspectionsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInspectionsRequest()
	}
	response = NewDescribeClusterInspectionsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEdgeClusterAuditRequest() (request *EnableEdgeClusterAuditRequest) {
	request = &EnableEdgeClusterAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEdgeClusterAudit")
	return
}

func NewEnableEdgeClusterAuditResponse() (response *EnableEdgeClusterAuditResponse) {
	response = &EnableEdgeClusterAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启边缘集群审计
func (c *Client) EnableEdgeClusterAudit(request *EnableEdgeClusterAuditRequest) (response *EnableEdgeClusterAuditResponse, err error) {
	if request == nil {
		request = NewEnableEdgeClusterAuditRequest()
	}
	response = NewEnableEdgeClusterAuditResponse()
	err = c.Send(request, response)
	return
}

func NewUploadHelmChartRequest() (request *UploadHelmChartRequest) {
	request = &UploadHelmChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UploadHelmChart")
	return
}

func NewUploadHelmChartResponse() (response *UploadHelmChartResponse) {
	response = &UploadHelmChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传Chart
func (c *Client) UploadHelmChart(request *UploadHelmChartRequest) (response *UploadHelmChartResponse, err error) {
	if request == nil {
		request = NewUploadHelmChartRequest()
	}
	response = NewUploadHelmChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterVirtualNodePoolsRequest() (request *DescribeClusterVirtualNodePoolsRequest) {
	request = &DescribeClusterVirtualNodePoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNodePools")
	return
}

func NewDescribeClusterVirtualNodePoolsResponse() (response *DescribeClusterVirtualNodePoolsResponse) {
	response = &DescribeClusterVirtualNodePoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看超级节点池列表
func (c *Client) DescribeClusterVirtualNodePools(request *DescribeClusterVirtualNodePoolsRequest) (response *DescribeClusterVirtualNodePoolsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterVirtualNodePoolsRequest()
	}
	response = NewDescribeClusterVirtualNodePoolsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEdgeEniRequest() (request *UpdateEdgeEniRequest) {
	request = &UpdateEdgeEniRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEdgeEni")
	return
}

func NewUpdateEdgeEniResponse() (response *UpdateEdgeEniResponse) {
	response = &UpdateEdgeEniResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新边缘集群支持独立网卡的状态
func (c *Client) UpdateEdgeEni(request *UpdateEdgeEniRequest) (response *UpdateEdgeEniResponse, err error) {
	if request == nil {
		request = NewUpdateEdgeEniRequest()
	}
	response = NewUpdateEdgeEniResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateExternalClusterInternalRequest() (request *UpdateExternalClusterInternalRequest) {
	request = &UpdateExternalClusterInternalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateExternalClusterInternal")
	return
}

func NewUpdateExternalClusterInternalResponse() (response *UpdateExternalClusterInternalResponse) {
	response = &UpdateExternalClusterInternalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改第三方集群信息（内部调用）
func (c *Client) UpdateExternalClusterInternal(request *UpdateExternalClusterInternalRequest) (response *UpdateExternalClusterInternalResponse, err error) {
	if request == nil {
		request = NewUpdateExternalClusterInternalRequest()
	}
	response = NewUpdateExternalClusterInternalResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEtcdInstanceRequest() (request *DeleteEtcdInstanceRequest) {
	request = &DeleteEtcdInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEtcdInstance")
	return
}

func NewDeleteEtcdInstanceResponse() (response *DeleteEtcdInstanceResponse) {
	response = &DeleteEtcdInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除etcd实例
func (c *Client) DeleteEtcdInstance(request *DeleteEtcdInstanceRequest) (response *DeleteEtcdInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteEtcdInstanceRequest()
	}
	response = NewDeleteEtcdInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHelmChartVersionRequest() (request *DeleteHelmChartVersionRequest) {
	request = &DeleteHelmChartVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteHelmChartVersion")
	return
}

func NewDeleteHelmChartVersionResponse() (response *DeleteHelmChartVersionResponse) {
	response = &DeleteHelmChartVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Chart指定版本
func (c *Client) DeleteHelmChartVersion(request *DeleteHelmChartVersionRequest) (response *DeleteHelmChartVersionResponse, err error) {
	if request == nil {
		request = NewDeleteHelmChartVersionRequest()
	}
	response = NewDeleteHelmChartVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTkeEdgeAlarmPoliciesRequest() (request *DescribeTkeEdgeAlarmPoliciesRequest) {
	request = &DescribeTkeEdgeAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTkeEdgeAlarmPolicies")
	return
}

func NewDescribeTkeEdgeAlarmPoliciesResponse() (response *DescribeTkeEdgeAlarmPoliciesResponse) {
	response = &DescribeTkeEdgeAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘集群的告警策略列表
func (c *Client) DescribeTkeEdgeAlarmPolicies(request *DescribeTkeEdgeAlarmPoliciesRequest) (response *DescribeTkeEdgeAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeTkeEdgeAlarmPoliciesRequest()
	}
	response = NewDescribeTkeEdgeAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSuperNodesRequest() (request *DescribeSuperNodesRequest) {
	request = &DescribeSuperNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeSuperNodes")
	return
}

func NewDescribeSuperNodesResponse() (response *DescribeSuperNodesResponse) {
	response = &DescribeSuperNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询超级节点的详细信息
func (c *Client) DescribeSuperNodes(request *DescribeSuperNodesRequest) (response *DescribeSuperNodesResponse, err error) {
	if request == nil {
		request = NewDescribeSuperNodesRequest()
	}
	response = NewDescribeSuperNodesResponse()
	err = c.Send(request, response)
	return
}

func NewForwardRequestRequest() (request *ForwardRequestRequest) {
	request = &ForwardRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardRequest")
	return
}

func NewForwardRequestResponse() (response *ForwardRequestResponse) {
	response = &ForwardRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// YUNAPI 转发请求给TKE APIServer接口
func (c *Client) ForwardRequest(request *ForwardRequestRequest) (response *ForwardRequestResponse, err error) {
	if request == nil {
		request = NewForwardRequestRequest()
	}
	response = NewForwardRequestResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterEndpointVipStatusRequest() (request *DescribeClusterEndpointVipStatusRequest) {
	request = &DescribeClusterEndpointVipStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointVipStatus")
	return
}

func NewDescribeClusterEndpointVipStatusResponse() (response *DescribeClusterEndpointVipStatusResponse) {
	response = &DescribeClusterEndpointVipStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群开启端口流程状态(仅支持托管集群外网端口)
func (c *Client) DescribeClusterEndpointVipStatus(request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterEndpointVipStatusRequest()
	}
	response = NewDescribeClusterEndpointVipStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowIdStatusRequest() (request *DescribeFlowIdStatusRequest) {
	request = &DescribeFlowIdStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeFlowIdStatus")
	return
}

func NewDescribeFlowIdStatusResponse() (response *DescribeFlowIdStatusResponse) {
	response = &DescribeFlowIdStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群开启端口流程状态(仅支持托管集群外网端口)
func (c *Client) DescribeFlowIdStatus(request *DescribeFlowIdStatusRequest) (response *DescribeFlowIdStatusResponse, err error) {
	if request == nil {
		request = NewDescribeFlowIdStatusRequest()
	}
	response = NewDescribeFlowIdStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCHubClustersRequest() (request *DescribeTDCCHubClustersRequest) {
	request = &DescribeTDCCHubClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCHubClusters")
	return
}

func NewDescribeTDCCHubClustersResponse() (response *DescribeTDCCHubClustersResponse) {
	response = &DescribeTDCCHubClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Hub集群列表
func (c *Client) DescribeTDCCHubClusters(request *DescribeTDCCHubClustersRequest) (response *DescribeTDCCHubClustersResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCHubClustersRequest()
	}
	response = NewDescribeTDCCHubClustersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirtualClusterRequest() (request *CreateVirtualClusterRequest) {
	request = &CreateVirtualClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateVirtualCluster")
	return
}

func NewCreateVirtualClusterResponse() (response *CreateVirtualClusterResponse) {
	response = &CreateVirtualClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建虚拟容器集群
func (c *Client) CreateVirtualCluster(request *CreateVirtualClusterRequest) (response *CreateVirtualClusterResponse, err error) {
	if request == nil {
		request = NewCreateVirtualClusterRequest()
	}
	response = NewCreateVirtualClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterCIDRRequest() (request *CheckClusterCIDRRequest) {
	request = &CheckClusterCIDRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckClusterCIDR")
	return
}

func NewCheckClusterCIDRResponse() (response *CheckClusterCIDRResponse) {
	response = &CheckClusterCIDRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查集群的CIDR是否冲突
func (c *Client) CheckClusterCIDR(request *CheckClusterCIDRRequest) (response *CheckClusterCIDRResponse, err error) {
	if request == nil {
		request = NewCheckClusterCIDRRequest()
	}
	response = NewCheckClusterCIDRResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSContainerInstancesRequest() (request *DescribeEKSContainerInstancesRequest) {
	request = &DescribeEKSContainerInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSContainerInstances")
	return
}

func NewDescribeEKSContainerInstancesResponse() (response *DescribeEKSContainerInstancesResponse) {
	response = &DescribeEKSContainerInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器实例
func (c *Client) DescribeEKSContainerInstances(request *DescribeEKSContainerInstancesRequest) (response *DescribeEKSContainerInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEKSContainerInstancesRequest()
	}
	response = NewDescribeEKSContainerInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudRunHPARequest() (request *DeleteCloudRunHPARequest) {
	request = &DeleteCloudRunHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteCloudRunHPA")
	return
}

func NewDeleteCloudRunHPAResponse() (response *DeleteCloudRunHPAResponse) {
	response = &DeleteCloudRunHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cloudrun hpa
func (c *Client) DeleteCloudRunHPA(request *DeleteCloudRunHPARequest) (response *DeleteCloudRunHPAResponse, err error) {
	if request == nil {
		request = NewDeleteCloudRunHPARequest()
	}
	response = NewDeleteCloudRunHPAResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeUserBillingWhiteListRequest() (request *DescribeEdgeUserBillingWhiteListRequest) {
	request = &DescribeEdgeUserBillingWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeUserBillingWhiteList")
	return
}

func NewDescribeEdgeUserBillingWhiteListResponse() (response *DescribeEdgeUserBillingWhiteListResponse) {
	response = &DescribeEdgeUserBillingWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户是否在边缘容器计费白名单中
func (c *Client) DescribeEdgeUserBillingWhiteList(request *DescribeEdgeUserBillingWhiteListRequest) (response *DescribeEdgeUserBillingWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeUserBillingWhiteListRequest()
	}
	response = NewDescribeEdgeUserBillingWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusInstanceRequest() (request *DescribePrometheusInstanceRequest) {
	request = &DescribePrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstance")
	return
}

func NewDescribePrometheusInstanceResponse() (response *DescribePrometheusInstanceResponse) {
	response = &DescribePrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例详细信息
func (c *Client) DescribePrometheusInstance(request *DescribePrometheusInstanceRequest) (response *DescribePrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusInstanceRequest()
	}
	response = NewDescribePrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewAddNodeToNodePoolRequest() (request *AddNodeToNodePoolRequest) {
	request = &AddNodeToNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddNodeToNodePool")
	return
}

func NewAddNodeToNodePoolResponse() (response *AddNodeToNodePoolResponse) {
	response = &AddNodeToNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将集群内节点移入节点池
func (c *Client) AddNodeToNodePool(request *AddNodeToNodePoolRequest) (response *AddNodeToNodePoolResponse, err error) {
	if request == nil {
		request = NewAddNodeToNodePoolRequest()
	}
	response = NewAddNodeToNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewSetContainerProbeResultRequest() (request *SetContainerProbeResultRequest) {
	request = &SetContainerProbeResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SetContainerProbeResult")
	return
}

func NewSetContainerProbeResultResponse() (response *SetContainerProbeResultResponse) {
	response = &SetContainerProbeResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SetContainerProbeResult
func (c *Client) SetContainerProbeResult(request *SetContainerProbeResultRequest) (response *SetContainerProbeResultResponse, err error) {
	if request == nil {
		request = NewSetContainerProbeResultRequest()
	}
	response = NewSetContainerProbeResultResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEdgeClusterInstancesRequest() (request *DeleteEdgeClusterInstancesRequest) {
	request = &DeleteEdgeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeClusterInstances")
	return
}

func NewDeleteEdgeClusterInstancesResponse() (response *DeleteEdgeClusterInstancesResponse) {
	response = &DeleteEdgeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除边缘计算实例
func (c *Client) DeleteEdgeClusterInstances(request *DeleteEdgeClusterInstancesRequest) (response *DeleteEdgeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteEdgeClusterInstancesRequest()
	}
	response = NewDeleteEdgeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterDeletionProtectionStateRequest() (request *ModifyClusterDeletionProtectionStateRequest) {
	request = &ModifyClusterDeletionProtectionStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterDeletionProtectionState")
	return
}

func NewModifyClusterDeletionProtectionStateResponse() (response *ModifyClusterDeletionProtectionStateResponse) {
	response = &ModifyClusterDeletionProtectionStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群删除保护状态
func (c *Client) ModifyClusterDeletionProtectionState(request *ModifyClusterDeletionProtectionStateRequest) (response *ModifyClusterDeletionProtectionStateResponse, err error) {
	if request == nil {
		request = NewModifyClusterDeletionProtectionStateRequest()
	}
	response = NewModifyClusterDeletionProtectionStateResponse()
	err = c.Send(request, response)
	return
}

func NewForwardTKEEdgeApplicationRequestV3Request() (request *ForwardTKEEdgeApplicationRequestV3Request) {
	request = &ForwardTKEEdgeApplicationRequestV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardTKEEdgeApplicationRequestV3")
	return
}

func NewForwardTKEEdgeApplicationRequestV3Response() (response *ForwardTKEEdgeApplicationRequestV3Response) {
	response = &ForwardTKEEdgeApplicationRequestV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作TKEEdge集群的addon
func (c *Client) ForwardTKEEdgeApplicationRequestV3(request *ForwardTKEEdgeApplicationRequestV3Request) (response *ForwardTKEEdgeApplicationRequestV3Response, err error) {
	if request == nil {
		request = NewForwardTKEEdgeApplicationRequestV3Request()
	}
	response = NewForwardTKEEdgeApplicationRequestV3Response()
	err = c.Send(request, response)
	return
}

func NewSyncPrometheusTemplateRequest() (request *SyncPrometheusTemplateRequest) {
	request = &SyncPrometheusTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "SyncPrometheusTemplate")
	return
}

func NewSyncPrometheusTemplateResponse() (response *SyncPrometheusTemplateResponse) {
	response = &SyncPrometheusTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步模板到实例或者集群
func (c *Client) SyncPrometheusTemplate(request *SyncPrometheusTemplateRequest) (response *SyncPrometheusTemplateResponse, err error) {
	if request == nil {
		request = NewSyncPrometheusTemplateRequest()
	}
	response = NewSyncPrometheusTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
	request = &AddExistedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddExistedInstances")
	return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
	response = &AddExistedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加已经存在的实例到集群
func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
	if request == nil {
		request = NewAddExistedInstancesRequest()
	}
	response = NewAddExistedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRIPodDetailRequest() (request *DescribeRIPodDetailRequest) {
	request = &DescribeRIPodDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRIPodDetail")
	return
}

func NewDescribeRIPodDetailResponse() (response *DescribeRIPodDetailResponse) {
	response = &DescribeRIPodDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按规格查询预留券和Pod数量
func (c *Client) DescribeRIPodDetail(request *DescribeRIPodDetailRequest) (response *DescribeRIPodDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRIPodDetailRequest()
	}
	response = NewDescribeRIPodDetailResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCloudRunEventPersistenceRequest() (request *EnableCloudRunEventPersistenceRequest) {
	request = &EnableCloudRunEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableCloudRunEventPersistence")
	return
}

func NewEnableCloudRunEventPersistenceResponse() (response *EnableCloudRunEventPersistenceResponse) {
	response = &EnableCloudRunEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启cloudrun事件持久化功能
func (c *Client) EnableCloudRunEventPersistence(request *EnableCloudRunEventPersistenceRequest) (response *EnableCloudRunEventPersistenceResponse, err error) {
	if request == nil {
		request = NewEnableCloudRunEventPersistenceRequest()
	}
	response = NewEnableCloudRunEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEtcdInstanceRequest() (request *CreateEtcdInstanceRequest) {
	request = &CreateEtcdInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEtcdInstance")
	return
}

func NewCreateEtcdInstanceResponse() (response *CreateEtcdInstanceResponse) {
	response = &CreateEtcdInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建etcd实例
func (c *Client) CreateEtcdInstance(request *CreateEtcdInstanceRequest) (response *CreateEtcdInstanceResponse, err error) {
	if request == nil {
		request = NewCreateEtcdInstanceRequest()
	}
	response = NewCreateEtcdInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReservedInstancesRequest() (request *CreateReservedInstancesRequest) {
	request = &CreateReservedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateReservedInstances")
	return
}

func NewCreateReservedInstancesResponse() (response *CreateReservedInstancesResponse) {
	response = &CreateReservedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预留券实例的购买会预先扣除本次实例购买所需金额，在调用本接口前请确保账户余额充足。
func (c *Client) CreateReservedInstances(request *CreateReservedInstancesRequest) (response *CreateReservedInstancesResponse, err error) {
	if request == nil {
		request = NewCreateReservedInstancesRequest()
	}
	response = NewCreateReservedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSDisksRequest() (request *DescribeEKSDisksRequest) {
	request = &DescribeEKSDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSDisks")
	return
}

func NewDescribeEKSDisksResponse() (response *DescribeEKSDisksResponse) {
	response = &DescribeEKSDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用CBS接口查询 EKS 创建的隐藏云硬盘
func (c *Client) DescribeEKSDisks(request *DescribeEKSDisksRequest) (response *DescribeEKSDisksResponse, err error) {
	if request == nil {
		request = NewDescribeEKSDisksRequest()
	}
	response = NewDescribeEKSDisksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEdgeLogConfigRequest() (request *CreateEdgeLogConfigRequest) {
	request = &CreateEdgeLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeLogConfig")
	return
}

func NewCreateEdgeLogConfigResponse() (response *CreateEdgeLogConfigResponse) {
	response = &CreateEdgeLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建边缘集群日志采集配置
func (c *Client) CreateEdgeLogConfig(request *CreateEdgeLogConfigRequest) (response *CreateEdgeLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateEdgeLogConfigRequest()
	}
	response = NewCreateEdgeLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHelmChartRequest() (request *DeleteHelmChartRequest) {
	request = &DeleteHelmChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteHelmChart")
	return
}

func NewDeleteHelmChartResponse() (response *DeleteHelmChartResponse) {
	response = &DeleteHelmChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Chart
func (c *Client) DeleteHelmChart(request *DeleteHelmChartRequest) (response *DeleteHelmChartResponse, err error) {
	if request == nil {
		request = NewDeleteHelmChartRequest()
	}
	response = NewDeleteHelmChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCcnInstancesRequest() (request *DescribeCcnInstancesRequest) {
	request = &DescribeCcnInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCcnInstances")
	return
}

func NewDescribeCcnInstancesResponse() (response *DescribeCcnInstancesResponse) {
	response = &DescribeCcnInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询vpc是否加入云联网
func (c *Client) DescribeCcnInstances(request *DescribeCcnInstancesRequest) (response *DescribeCcnInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeCcnInstancesRequest()
	}
	response = NewDescribeCcnInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaFeatureForEksRequest() (request *UpdateMetaFeatureForEksRequest) {
	request = &UpdateMetaFeatureForEksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateMetaFeatureForEks")
	return
}

func NewUpdateMetaFeatureForEksResponse() (response *UpdateMetaFeatureForEksResponse) {
	response = &UpdateMetaFeatureForEksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EKS集群更新跨租户弹性网卡配置接口
func (c *Client) UpdateMetaFeatureForEks(request *UpdateMetaFeatureForEksRequest) (response *UpdateMetaFeatureForEksResponse, err error) {
	if request == nil {
		request = NewUpdateMetaFeatureForEksRequest()
	}
	response = NewUpdateMetaFeatureForEksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBootstrapTokensRequest() (request *DescribeBootstrapTokensRequest) {
	request = &DescribeBootstrapTokensRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeBootstrapTokens")
	return
}

func NewDescribeBootstrapTokensResponse() (response *DescribeBootstrapTokensResponse) {
	response = &DescribeBootstrapTokensResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询BootstrapTokens列表
func (c *Client) DescribeBootstrapTokens(request *DescribeBootstrapTokensRequest) (response *DescribeBootstrapTokensResponse, err error) {
	if request == nil {
		request = NewDescribeBootstrapTokensRequest()
	}
	response = NewDescribeBootstrapTokensResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterKubeconfigRequest() (request *DescribeClusterKubeconfigRequest) {
	request = &DescribeClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterKubeconfig")
	return
}

func NewDescribeClusterKubeconfigResponse() (response *DescribeClusterKubeconfigResponse) {
	response = &DescribeClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的kubeconfig文件，不同子账户获取自己的kubeconfig文件，该文件中有每个子账户自己的kube-apiserver的客户端证书，默认首次调此接口时候创建客户端证书，时效20年，未授予任何权限，如果是集群所有者或者主账户，则默认是cluster-admin权限。
func (c *Client) DescribeClusterKubeconfig(request *DescribeClusterKubeconfigRequest) (response *DescribeClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewDescribeClusterKubeconfigRequest()
	}
	response = NewDescribeClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalNodePoolsRequest() (request *DescribeExternalNodePoolsRequest) {
	request = &DescribeExternalNodePoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalNodePools")
	return
}

func NewDescribeExternalNodePoolsResponse() (response *DescribeExternalNodePoolsResponse) {
	response = &DescribeExternalNodePoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看第三方节点池列表
func (c *Client) DescribeExternalNodePools(request *DescribeExternalNodePoolsRequest) (response *DescribeExternalNodePoolsResponse, err error) {
	if request == nil {
		request = NewDescribeExternalNodePoolsRequest()
	}
	response = NewDescribeExternalNodePoolsResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterQuotaRequest() (request *GetClusterQuotaRequest) {
	request = &GetClusterQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetClusterQuota")
	return
}

func NewGetClusterQuotaResponse() (response *GetClusterQuotaResponse) {
	response = &GetClusterQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetClusterQuota
func (c *Client) GetClusterQuota(request *GetClusterQuotaRequest) (response *GetClusterQuotaResponse, err error) {
	if request == nil {
		request = NewGetClusterQuotaRequest()
	}
	response = NewGetClusterQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodSpecificationRequest() (request *GetPodSpecificationRequest) {
	request = &GetPodSpecificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodSpecification")
	return
}

func NewGetPodSpecificationResponse() (response *GetPodSpecificationResponse) {
	response = &GetPodSpecificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 Pod 规格
func (c *Client) GetPodSpecification(request *GetPodSpecificationRequest) (response *GetPodSpecificationResponse, err error) {
	if request == nil {
		request = NewGetPodSpecificationRequest()
	}
	response = NewGetPodSpecificationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSInstancesRequest() (request *DescribeEKSInstancesRequest) {
	request = &DescribeEKSInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSInstances")
	return
}

func NewDescribeEKSInstancesResponse() (response *DescribeEKSInstancesResponse) {
	response = &DescribeEKSInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性容器实例
func (c *Client) DescribeEKSInstances(request *DescribeEKSInstancesRequest) (response *DescribeEKSInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEKSInstancesRequest()
	}
	response = NewDescribeEKSInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterPodsRequest() (request *DescribeClusterPodsRequest) {
	request = &DescribeClusterPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterPods")
	return
}

func NewDescribeClusterPodsResponse() (response *DescribeClusterPodsResponse) {
	response = &DescribeClusterPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口获取集群内Pod相关的详细描述信息，参考kubernetes API获取pod，只对内部短期使用
func (c *Client) DescribeClusterPods(request *DescribeClusterPodsRequest) (response *DescribeClusterPodsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterPodsRequest()
	}
	response = NewDescribeClusterPodsResponse()
	err = c.Send(request, response)
	return
}

func NewForwardPlatformRequestV3Request() (request *ForwardPlatformRequestV3Request) {
	request = &ForwardPlatformRequestV3Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardPlatformRequestV3")
	return
}

func NewForwardPlatformRequestV3Response() (response *ForwardPlatformRequestV3Response) {
	response = &ForwardPlatformRequestV3Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询、新增、删除、编辑TKE集群内的资源
func (c *Client) ForwardPlatformRequestV3(request *ForwardPlatformRequestV3Request) (response *ForwardPlatformRequestV3Response, err error) {
	if request == nil {
		request = NewForwardPlatformRequestV3Request()
	}
	response = NewForwardPlatformRequestV3Response()
	err = c.Send(request, response)
	return
}

func NewGetContainerLogsRequest() (request *GetContainerLogsRequest) {
	request = &GetContainerLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetContainerLogs")
	return
}

func NewGetContainerLogsResponse() (response *GetContainerLogsResponse) {
	response = &GetContainerLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取pod内容器日志
func (c *Client) GetContainerLogs(request *GetContainerLogsRequest) (response *GetContainerLogsResponse, err error) {
	if request == nil {
		request = NewGetContainerLogsRequest()
	}
	response = NewGetContainerLogsResponse()
	err = c.Send(request, response)
	return
}

func NewAddAlarmPolicyRequest() (request *AddAlarmPolicyRequest) {
	request = &AddAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddAlarmPolicy")
	return
}

func NewAddAlarmPolicyResponse() (response *AddAlarmPolicyResponse) {
	response = &AddAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加告警策略
func (c *Client) AddAlarmPolicy(request *AddAlarmPolicyRequest) (response *AddAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewAddAlarmPolicyRequest()
	}
	response = NewAddAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerLogRequest() (request *DescribeContainerLogRequest) {
	request = &DescribeContainerLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeContainerLog")
	return
}

func NewDescribeContainerLogResponse() (response *DescribeContainerLogResponse) {
	response = &DescribeContainerLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器组中容器日志
func (c *Client) DescribeContainerLog(request *DescribeContainerLogRequest) (response *DescribeContainerLogResponse, err error) {
	if request == nil {
		request = NewDescribeContainerLogRequest()
	}
	response = NewDescribeContainerLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusAlertRuleRequest() (request *ModifyPrometheusAlertRuleRequest) {
	request = &ModifyPrometheusAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAlertRule")
	return
}

func NewModifyPrometheusAlertRuleResponse() (response *ModifyPrometheusAlertRuleResponse) {
	response = &ModifyPrometheusAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警规则
func (c *Client) ModifyPrometheusAlertRule(request *ModifyPrometheusAlertRuleRequest) (response *ModifyPrometheusAlertRuleResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusAlertRuleRequest()
	}
	response = NewModifyPrometheusAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
	request = &DescribeExistedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExistedInstances")
	return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
	response = &DescribeExistedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询已经存在的节点，判断是否可以加入集群
func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeExistedInstancesRequest()
	}
	response = NewDescribeExistedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusAgentRequest() (request *DeletePrometheusAgentRequest) {
	request = &DeletePrometheusAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAgent")
	return
}

func NewDeletePrometheusAgentResponse() (response *DeletePrometheusAgentResponse) {
	response = &DeletePrometheusAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除集群关联
func (c *Client) DeletePrometheusAgent(request *DeletePrometheusAgentRequest) (response *DeletePrometheusAgentResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusAgentRequest()
	}
	response = NewDeletePrometheusAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdCreatingProgressRequest() (request *DescribeEtcdCreatingProgressRequest) {
	request = &DescribeEtcdCreatingProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdCreatingProgress")
	return
}

func NewDescribeEtcdCreatingProgressResponse() (response *DescribeEtcdCreatingProgressResponse) {
	response = &DescribeEtcdCreatingProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd集群创建进度
func (c *Client) DescribeEtcdCreatingProgress(request *DescribeEtcdCreatingProgressRequest) (response *DescribeEtcdCreatingProgressResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdCreatingProgressRequest()
	}
	response = NewDescribeEtcdCreatingProgressResponse()
	err = c.Send(request, response)
	return
}

func NewGetMostSuitableImageCacheRequest() (request *GetMostSuitableImageCacheRequest) {
	request = &GetMostSuitableImageCacheRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetMostSuitableImageCache")
	return
}

func NewGetMostSuitableImageCacheResponse() (response *GetMostSuitableImageCacheResponse) {
	response = &GetMostSuitableImageCacheResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据镜像列表，查询匹配的镜像缓存
func (c *Client) GetMostSuitableImageCache(request *GetMostSuitableImageCacheRequest) (response *GetMostSuitableImageCacheResponse, err error) {
	if request == nil {
		request = NewGetMostSuitableImageCacheRequest()
	}
	response = NewGetMostSuitableImageCacheResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLogCollectorHostPathRequest() (request *CheckLogCollectorHostPathRequest) {
	request = &CheckLogCollectorHostPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckLogCollectorHostPath")
	return
}

func NewCheckLogCollectorHostPathResponse() (response *CheckLogCollectorHostPathResponse) {
	response = &CheckLogCollectorHostPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查主机路径
func (c *Client) CheckLogCollectorHostPath(request *CheckLogCollectorHostPathRequest) (response *CheckLogCollectorHostPathResponse, err error) {
	if request == nil {
		request = NewCheckLogCollectorHostPathRequest()
	}
	response = NewCheckLogCollectorHostPathResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportedRuntimeRequest() (request *DescribeSupportedRuntimeRequest) {
	request = &DescribeSupportedRuntimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeSupportedRuntime")
	return
}

func NewDescribeSupportedRuntimeResponse() (response *DescribeSupportedRuntimeResponse) {
	response = &DescribeSupportedRuntimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据K8S版本获取可选运行时版本
func (c *Client) DescribeSupportedRuntime(request *DescribeSupportedRuntimeRequest) (response *DescribeSupportedRuntimeResponse, err error) {
	if request == nil {
		request = NewDescribeSupportedRuntimeRequest()
	}
	response = NewDescribeSupportedRuntimeResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatsSummaryRequest() (request *GetStatsSummaryRequest) {
	request = &GetStatsSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetStatsSummary")
	return
}

func NewGetStatsSummaryResponse() (response *GetStatsSummaryResponse) {
	response = &GetStatsSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询pod监控信息
func (c *Client) GetStatsSummary(request *GetStatsSummaryRequest) (response *GetStatsSummaryResponse, err error) {
	if request == nil {
		request = NewGetStatsSummaryRequest()
	}
	response = NewGetStatsSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCLSLogConfigRequest() (request *CreateCLSLogConfigRequest) {
	request = &CreateCLSLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateCLSLogConfig")
	return
}

func NewCreateCLSLogConfigResponse() (response *CreateCLSLogConfigResponse) {
	response = &CreateCLSLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志采集配置
func (c *Client) CreateCLSLogConfig(request *CreateCLSLogConfigRequest) (response *CreateCLSLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateCLSLogConfigRequest()
	}
	response = NewCreateCLSLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCheckInstancesUpgradeAbleRequest() (request *CheckInstancesUpgradeAbleRequest) {
	request = &CheckInstancesUpgradeAbleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckInstancesUpgradeAble")
	return
}

func NewCheckInstancesUpgradeAbleResponse() (response *CheckInstancesUpgradeAbleResponse) {
	response = &CheckInstancesUpgradeAbleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查给定节点列表中哪些是可升级的
func (c *Client) CheckInstancesUpgradeAble(request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
	if request == nil {
		request = NewCheckInstancesUpgradeAbleRequest()
	}
	response = NewCheckInstancesUpgradeAbleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableClusterVersionRequest() (request *DescribeAvailableClusterVersionRequest) {
	request = &DescribeAvailableClusterVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableClusterVersion")
	return
}

func NewDescribeAvailableClusterVersionResponse() (response *DescribeAvailableClusterVersionResponse) {
	response = &DescribeAvailableClusterVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群可以升级的所有版本
func (c *Client) DescribeAvailableClusterVersion(request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableClusterVersionRequest()
	}
	response = NewDescribeAvailableClusterVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemainingIpForGlobalRouteRequest() (request *DescribeRemainingIpForGlobalRouteRequest) {
	request = &DescribeRemainingIpForGlobalRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRemainingIpForGlobalRoute")
	return
}

func NewDescribeRemainingIpForGlobalRouteResponse() (response *DescribeRemainingIpForGlobalRouteResponse) {
	response = &DescribeRemainingIpForGlobalRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看全局路由模式的集群中，容器网络中剩余IP的数量
func (c *Client) DescribeRemainingIpForGlobalRoute(request *DescribeRemainingIpForGlobalRouteRequest) (response *DescribeRemainingIpForGlobalRouteResponse, err error) {
	if request == nil {
		request = NewDescribeRemainingIpForGlobalRouteRequest()
	}
	response = NewDescribeRemainingIpForGlobalRouteResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterEndpointSPRequest() (request *ModifyClusterEndpointSPRequest) {
	request = &ModifyClusterEndpointSPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterEndpointSP")
	return
}

func NewModifyClusterEndpointSPResponse() (response *ModifyClusterEndpointSPResponse) {
	response = &ModifyClusterEndpointSPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改托管集群外网端口的安全策略（老的方式，仅支持托管集群外网端口）
func (c *Client) ModifyClusterEndpointSP(request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
	if request == nil {
		request = NewModifyClusterEndpointSPRequest()
	}
	response = NewModifyClusterEndpointSPResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusRecordRuleRequest() (request *CreatePrometheusRecordRuleRequest) {
	request = &CreatePrometheusRecordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusRecordRule")
	return
}

func NewCreatePrometheusRecordRuleResponse() (response *CreatePrometheusRecordRuleResponse) {
	response = &CreatePrometheusRecordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增聚合规则
func (c *Client) CreatePrometheusRecordRule(request *CreatePrometheusRecordRuleRequest) (response *CreatePrometheusRecordRuleResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusRecordRuleRequest()
	}
	response = NewCreatePrometheusRecordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterImageRequest() (request *ModifyClusterImageRequest) {
	request = &ModifyClusterImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterImage")
	return
}

func NewModifyClusterImageResponse() (response *ModifyClusterImageResponse) {
	response = &ModifyClusterImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群镜像
func (c *Client) ModifyClusterImage(request *ModifyClusterImageRequest) (response *ModifyClusterImageResponse, err error) {
	if request == nil {
		request = NewModifyClusterImageRequest()
	}
	response = NewModifyClusterImageResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusConfigRequest() (request *CreatePrometheusConfigRequest) {
	request = &CreatePrometheusConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusConfig")
	return
}

func NewCreatePrometheusConfigResponse() (response *CreatePrometheusConfigResponse) {
	response = &CreatePrometheusConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群采集配置
func (c *Client) CreatePrometheusConfig(request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusConfigRequest()
	}
	response = NewCreatePrometheusConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsLogSetsRequest() (request *DescribeClsLogSetsRequest) {
	request = &DescribeClsLogSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClsLogSets")
	return
}

func NewDescribeClsLogSetsResponse() (response *DescribeClsLogSetsResponse) {
	response = &DescribeClsLogSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出CLS日志集
func (c *Client) DescribeClsLogSets(request *DescribeClsLogSetsRequest) (response *DescribeClsLogSetsResponse, err error) {
	if request == nil {
		request = NewDescribeClsLogSetsRequest()
	}
	response = NewDescribeClsLogSetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusDashboardRequest() (request *CreatePrometheusDashboardRequest) {
	request = &CreatePrometheusDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusDashboard")
	return
}

func NewCreatePrometheusDashboardResponse() (response *CreatePrometheusDashboardResponse) {
	response = &CreatePrometheusDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建grafana监控面板
func (c *Client) CreatePrometheusDashboard(request *CreatePrometheusDashboardRequest) (response *CreatePrometheusDashboardResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusDashboardRequest()
	}
	response = NewCreatePrometheusDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdOverrideConfigsRequest() (request *DescribeEtcdOverrideConfigsRequest) {
	request = &DescribeEtcdOverrideConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdOverrideConfigs")
	return
}

func NewDescribeEtcdOverrideConfigsResponse() (response *DescribeEtcdOverrideConfigsResponse) {
	response = &DescribeEtcdOverrideConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看集群元数据拆分存储的Etcd配置列表
func (c *Client) DescribeEtcdOverrideConfigs(request *DescribeEtcdOverrideConfigsRequest) (response *DescribeEtcdOverrideConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdOverrideConfigsRequest()
	}
	response = NewDescribeEtcdOverrideConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewGetECMSubnetResourceRequest() (request *GetECMSubnetResourceRequest) {
	request = &GetECMSubnetResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetECMSubnetResource")
	return
}

func NewGetECMSubnetResourceResponse() (response *GetECMSubnetResourceResponse) {
	response = &GetECMSubnetResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ECM子网资源
func (c *Client) GetECMSubnetResource(request *GetECMSubnetResourceRequest) (response *GetECMSubnetResourceResponse, err error) {
	if request == nil {
		request = NewGetECMSubnetResourceRequest()
	}
	response = NewGetECMSubnetResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCClusterCommonNamesRequest() (request *DescribeTDCCClusterCommonNamesRequest) {
	request = &DescribeTDCCClusterCommonNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCClusterCommonNames")
	return
}

func NewDescribeTDCCClusterCommonNamesResponse() (response *DescribeTDCCClusterCommonNamesResponse) {
	response = &DescribeTDCCClusterCommonNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定子账户的对应kube-apiserver的客户端证书中CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为20
func (c *Client) DescribeTDCCClusterCommonNames(request *DescribeTDCCClusterCommonNamesRequest) (response *DescribeTDCCClusterCommonNamesResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCClusterCommonNamesRequest()
	}
	response = NewDescribeTDCCClusterCommonNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNodePoolInstanceTypesRequest() (request *ModifyNodePoolInstanceTypesRequest) {
	request = &ModifyNodePoolInstanceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePoolInstanceTypes")
	return
}

func NewModifyNodePoolInstanceTypesResponse() (response *ModifyNodePoolInstanceTypesResponse) {
	response = &ModifyNodePoolInstanceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改节点池的机型配置
func (c *Client) ModifyNodePoolInstanceTypes(request *ModifyNodePoolInstanceTypesRequest) (response *ModifyNodePoolInstanceTypesResponse, err error) {
	if request == nil {
		request = NewModifyNodePoolInstanceTypesRequest()
	}
	response = NewModifyNodePoolInstanceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOldStaticIPSubnetsRequest() (request *DescribeOldStaticIPSubnetsRequest) {
	request = &DescribeOldStaticIPSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeOldStaticIPSubnets")
	return
}

func NewDescribeOldStaticIPSubnetsResponse() (response *DescribeOldStaticIPSubnetsResponse) {
	response = &DescribeOldStaticIPSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询旧版固定 IP 集群独占的子网列表，可查询独占的子网和对应的集群
func (c *Client) DescribeOldStaticIPSubnets(request *DescribeOldStaticIPSubnetsRequest) (response *DescribeOldStaticIPSubnetsResponse, err error) {
	if request == nil {
		request = NewDescribeOldStaticIPSubnetsRequest()
	}
	response = NewDescribeOldStaticIPSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeECMEKSClusterCommonNamesRequest() (request *DescribeECMEKSClusterCommonNamesRequest) {
	request = &DescribeECMEKSClusterCommonNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeECMEKSClusterCommonNames")
	return
}

func NewDescribeECMEKSClusterCommonNamesResponse() (response *DescribeECMEKSClusterCommonNamesResponse) {
	response = &DescribeECMEKSClusterCommonNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定子账户的对应kube-apiserver的客户端证书中CommonName字段，如果没有客户端证书，将会签发一个，此接口有最大传入子账户数量上限，当前为50
func (c *Client) DescribeECMEKSClusterCommonNames(request *DescribeECMEKSClusterCommonNamesRequest) (response *DescribeECMEKSClusterCommonNamesResponse, err error) {
	if request == nil {
		request = NewDescribeECMEKSClusterCommonNamesRequest()
	}
	response = NewDescribeECMEKSClusterCommonNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdQuotaRequest() (request *DescribeEtcdQuotaRequest) {
	request = &DescribeEtcdQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdQuota")
	return
}

func NewDescribeEtcdQuotaResponse() (response *DescribeEtcdQuotaResponse) {
	response = &DescribeEtcdQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd集群配额
func (c *Client) DescribeEtcdQuota(request *DescribeEtcdQuotaRequest) (response *DescribeEtcdQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdQuotaRequest()
	}
	response = NewDescribeEtcdQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Service
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalNodeSupportConfigRequest() (request *DescribeExternalNodeSupportConfigRequest) {
	request = &DescribeExternalNodeSupportConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalNodeSupportConfig")
	return
}

func NewDescribeExternalNodeSupportConfigResponse() (response *DescribeExternalNodeSupportConfigResponse) {
	response = &DescribeExternalNodeSupportConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看开启第三方节点池配置信息
func (c *Client) DescribeExternalNodeSupportConfig(request *DescribeExternalNodeSupportConfigRequest) (response *DescribeExternalNodeSupportConfigResponse, err error) {
	if request == nil {
		request = NewDescribeExternalNodeSupportConfigRequest()
	}
	response = NewDescribeExternalNodeSupportConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusBasicMetricsRequest() (request *DescribePrometheusBasicMetricsRequest) {
	request = &DescribePrometheusBasicMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusBasicMetrics")
	return
}

func NewDescribePrometheusBasicMetricsResponse() (response *DescribePrometheusBasicMetricsResponse) {
	response = &DescribePrometheusBasicMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取默认采集指标详情
func (c *Client) DescribePrometheusBasicMetrics(request *DescribePrometheusBasicMetricsRequest) (response *DescribePrometheusBasicMetricsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusBasicMetricsRequest()
	}
	response = NewDescribePrometheusBasicMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
	request = &DescribeQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeQuota")
	return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
	response = &DescribeQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群配额
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaRequest()
	}
	response = NewDescribeQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDetachEKSDisksRequest() (request *DetachEKSDisksRequest) {
	request = &DetachEKSDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DetachEKSDisks")
	return
}

func NewDetachEKSDisksResponse() (response *DetachEKSDisksResponse) {
	response = &DetachEKSDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过内部 CBS 接口解绑 CXM上的 CBS 隐藏盘
func (c *Client) DetachEKSDisks(request *DetachEKSDisksRequest) (response *DetachEKSDisksResponse, err error) {
	if request == nil {
		request = NewDetachEKSDisksRequest()
	}
	response = NewDetachEKSDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdSnapshotsRequest() (request *DescribeEtcdSnapshotsRequest) {
	request = &DescribeEtcdSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdSnapshots")
	return
}

func NewDescribeEtcdSnapshotsResponse() (response *DescribeEtcdSnapshotsResponse) {
	response = &DescribeEtcdSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd快照列表
func (c *Client) DescribeEtcdSnapshots(request *DescribeEtcdSnapshotsRequest) (response *DescribeEtcdSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdSnapshotsRequest()
	}
	response = NewDescribeEtcdSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewGetEKSClusterResourcesRequest() (request *GetEKSClusterResourcesRequest) {
	request = &GetEKSClusterResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEKSClusterResources")
	return
}

func NewGetEKSClusterResourcesResponse() (response *GetEKSClusterResourcesResponse) {
	response = &GetEKSClusterResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询EKS集群资源资源信息
func (c *Client) GetEKSClusterResources(request *GetEKSClusterResourcesRequest) (response *GetEKSClusterResourcesResponse, err error) {
	if request == nil {
		request = NewGetEKSClusterResourcesRequest()
	}
	response = NewGetEKSClusterResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterControllersRequest() (request *DescribeClusterControllersRequest) {
	request = &DescribeClusterControllersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterControllers")
	return
}

func NewDescribeClusterControllersResponse() (response *DescribeClusterControllersResponse) {
	response = &DescribeClusterControllersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询Kubernetes的各个原生控制器是否开启
func (c *Client) DescribeClusterControllers(request *DescribeClusterControllersRequest) (response *DescribeClusterControllersResponse, err error) {
	if request == nil {
		request = NewDescribeClusterControllersRequest()
	}
	response = NewDescribeClusterControllersResponse()
	err = c.Send(request, response)
	return
}

func NewTransferPrometheusInstanceRequest() (request *TransferPrometheusInstanceRequest) {
	request = &TransferPrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "TransferPrometheusInstance")
	return
}

func NewTransferPrometheusInstanceResponse() (response *TransferPrometheusInstanceResponse) {
	response = &TransferPrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移TPS实例到TMP
func (c *Client) TransferPrometheusInstance(request *TransferPrometheusInstanceRequest) (response *TransferPrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewTransferPrometheusInstanceRequest()
	}
	response = NewTransferPrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusInstanceRequest() (request *DeletePrometheusInstanceRequest) {
	request = &DeletePrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusInstance")
	return
}

func NewDeletePrometheusInstanceResponse() (response *DeletePrometheusInstanceResponse) {
	response = &DeletePrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除实例
func (c *Client) DeletePrometheusInstance(request *DeletePrometheusInstanceRequest) (response *DeletePrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusInstanceRequest()
	}
	response = NewDeletePrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVirtualClusterRequest() (request *DeleteVirtualClusterRequest) {
	request = &DeleteVirtualClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteVirtualCluster")
	return
}

func NewDeleteVirtualClusterResponse() (response *DeleteVirtualClusterResponse) {
	response = &DeleteVirtualClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除虚拟集群，只有集群内服务为空时才能删除.
func (c *Client) DeleteVirtualCluster(request *DeleteVirtualClusterRequest) (response *DeleteVirtualClusterResponse, err error) {
	if request == nil {
		request = NewDeleteVirtualClusterRequest()
	}
	response = NewDeleteVirtualClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAuthStatusRequest() (request *DescribeClusterAuthStatusRequest) {
	request = &DescribeClusterAuthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAuthStatus")
	return
}

func NewDescribeClusterAuthStatusResponse() (response *DescribeClusterAuthStatusResponse) {
	response = &DescribeClusterAuthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群是否自动清理过期子账户状态，开启之后在子账户被删除7天之后会自动清理改子账户在集群内的权限和证书信息
func (c *Client) DescribeClusterAuthStatus(request *DescribeClusterAuthStatusRequest) (response *DescribeClusterAuthStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAuthStatusRequest()
	}
	response = NewDescribeClusterAuthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusDashboardRequest() (request *ModifyPrometheusDashboardRequest) {
	request = &ModifyPrometheusDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusDashboard")
	return
}

func NewModifyPrometheusDashboardResponse() (response *ModifyPrometheusDashboardResponse) {
	response = &ModifyPrometheusDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改grafana监控面板
func (c *Client) ModifyPrometheusDashboard(request *ModifyPrometheusDashboardRequest) (response *ModifyPrometheusDashboardResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusDashboardRequest()
	}
	response = NewModifyPrometheusDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterInstancesRequest() (request *UpdateClusterInstancesRequest) {
	request = &UpdateClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterInstances")
	return
}

func NewUpdateClusterInstancesResponse() (response *UpdateClusterInstancesResponse) {
	response = &UpdateClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群节点k8s版本升级，API 3.0
func (c *Client) UpdateClusterInstances(request *UpdateClusterInstancesRequest) (response *UpdateClusterInstancesResponse, err error) {
	if request == nil {
		request = NewUpdateClusterInstancesRequest()
	}
	response = NewUpdateClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAuthorizationModeRequest() (request *DescribeClusterAuthorizationModeRequest) {
	request = &DescribeClusterAuthorizationModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAuthorizationMode")
	return
}

func NewDescribeClusterAuthorizationModeResponse() (response *DescribeClusterAuthorizationModeResponse) {
	response = &DescribeClusterAuthorizationModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的授权模式
func (c *Client) DescribeClusterAuthorizationMode(request *DescribeClusterAuthorizationModeRequest) (response *DescribeClusterAuthorizationModeResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAuthorizationModeRequest()
	}
	response = NewDescribeClusterAuthorizationModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnableVpcCniProgressRequest() (request *DescribeEnableVpcCniProgressRequest) {
	request = &DescribeEnableVpcCniProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEnableVpcCniProgress")
	return
}

func NewDescribeEnableVpcCniProgressResponse() (response *DescribeEnableVpcCniProgressResponse) {
	response = &DescribeEnableVpcCniProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询开启vpc-cni模式的任务进度
func (c *Client) DescribeEnableVpcCniProgress(request *DescribeEnableVpcCniProgressRequest) (response *DescribeEnableVpcCniProgressResponse, err error) {
	if request == nil {
		request = NewDescribeEnableVpcCniProgressRequest()
	}
	response = NewDescribeEnableVpcCniProgressResponse()
	err = c.Send(request, response)
	return
}

func NewInstallEdgeLogAgentRequest() (request *InstallEdgeLogAgentRequest) {
	request = &InstallEdgeLogAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "InstallEdgeLogAgent")
	return
}

func NewInstallEdgeLogAgentResponse() (response *InstallEdgeLogAgentResponse) {
	response = &InstallEdgeLogAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在tke@edge集群的边缘节点上安装日志采集组件
func (c *Client) InstallEdgeLogAgent(request *InstallEdgeLogAgentRequest) (response *InstallEdgeLogAgentResponse, err error) {
	if request == nil {
		request = NewInstallEdgeLogAgentRequest()
	}
	response = NewInstallEdgeLogAgentResponse()
	err = c.Send(request, response)
	return
}

func NewCheckPodRetainRequest() (request *CheckPodRetainRequest) {
	request = &CheckPodRetainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckPodRetain")
	return
}

func NewCheckPodRetainResponse() (response *CheckPodRetainResponse) {
	response = &CheckPodRetainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询删除后的pod是否保留
func (c *Client) CheckPodRetain(request *CheckPodRetainRequest) (response *CheckPodRetainResponse, err error) {
	if request == nil {
		request = NewCheckPodRetainRequest()
	}
	response = NewCheckPodRetainResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEKSContainerInstancesRequest() (request *CreateEKSContainerInstancesRequest) {
	request = &CreateEKSContainerInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEKSContainerInstances")
	return
}

func NewCreateEKSContainerInstancesResponse() (response *CreateEKSContainerInstancesResponse) {
	response = &CreateEKSContainerInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建容器实例
func (c *Client) CreateEKSContainerInstances(request *CreateEKSContainerInstancesRequest) (response *CreateEKSContainerInstancesResponse, err error) {
	if request == nil {
		request = NewCreateEKSContainerInstancesRequest()
	}
	response = NewCreateEKSContainerInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalClusterMetricDataRequest() (request *DescribeExternalClusterMetricDataRequest) {
	request = &DescribeExternalClusterMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalClusterMetricData")
	return
}

func NewDescribeExternalClusterMetricDataResponse() (response *DescribeExternalClusterMetricDataResponse) {
	response = &DescribeExternalClusterMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出第三方集群Metric信息
func (c *Client) DescribeExternalClusterMetricData(request *DescribeExternalClusterMetricDataRequest) (response *DescribeExternalClusterMetricDataResponse, err error) {
	if request == nil {
		request = NewDescribeExternalClusterMetricDataRequest()
	}
	response = NewDescribeExternalClusterMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodChargeInfoRequest() (request *DescribePodChargeInfoRequest) {
	request = &DescribePodChargeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePodChargeInfo")
	return
}

func NewDescribePodChargeInfoResponse() (response *DescribePodChargeInfoResponse) {
	response = &DescribePodChargeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询正在运行中Pod的计费信息。可以通过 Namespace 和 Name 来查询某个 Pod 的信息，也可以通过 Pod 的 Uid 批量查询。
func (c *Client) DescribePodChargeInfo(request *DescribePodChargeInfoRequest) (response *DescribePodChargeInfoResponse, err error) {
	if request == nil {
		request = NewDescribePodChargeInfoRequest()
	}
	response = NewDescribePodChargeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeClusterCredentialRequest() (request *DescribeTKEEdgeClusterCredentialRequest) {
	request = &DescribeTKEEdgeClusterCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterCredential")
	return
}

func NewDescribeTKEEdgeClusterCredentialResponse() (response *DescribeTKEEdgeClusterCredentialResponse) {
	response = &DescribeTKEEdgeClusterCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘计算集群的认证信息
func (c *Client) DescribeTKEEdgeClusterCredential(request *DescribeTKEEdgeClusterCredentialRequest) (response *DescribeTKEEdgeClusterCredentialResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeClusterCredentialRequest()
	}
	response = NewDescribeTKEEdgeClusterCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTKEClusterRequest() (request *CreateTKEClusterRequest) {
	request = &CreateTKEClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateTKECluster")
	return
}

func NewCreateTKEClusterResponse() (response *CreateTKEClusterResponse) {
	response = &CreateTKEClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TKE集群
func (c *Client) CreateTKECluster(request *CreateTKEClusterRequest) (response *CreateTKEClusterResponse, err error) {
	if request == nil {
		request = NewCreateTKEClusterRequest()
	}
	response = NewCreateTKEClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodePoolsRequest() (request *DescribeClusterNodePoolsRequest) {
	request = &DescribeClusterNodePoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePools")
	return
}

func NewDescribeClusterNodePoolsResponse() (response *DescribeClusterNodePoolsResponse) {
	response = &DescribeClusterNodePoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点池列表
func (c *Client) DescribeClusterNodePools(request *DescribeClusterNodePoolsRequest) (response *DescribeClusterNodePoolsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodePoolsRequest()
	}
	response = NewDescribeClusterNodePoolsResponse()
	err = c.Send(request, response)
	return
}

func NewDisableMasterLogRequest() (request *DisableMasterLogRequest) {
	request = &DisableMasterLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableMasterLog")
	return
}

func NewDisableMasterLogResponse() (response *DisableMasterLogResponse) {
	response = &DisableMasterLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭指定master组件日志收集
func (c *Client) DisableMasterLog(request *DisableMasterLogRequest) (response *DisableMasterLogResponse, err error) {
	if request == nil {
		request = NewDisableMasterLogRequest()
	}
	response = NewDisableMasterLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusAlertRuleRequest() (request *DescribePrometheusAlertRuleRequest) {
	request = &DescribePrometheusAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertRule")
	return
}

func NewDescribePrometheusAlertRuleResponse() (response *DescribePrometheusAlertRuleResponse) {
	response = &DescribePrometheusAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警规则列表
func (c *Client) DescribePrometheusAlertRule(request *DescribePrometheusAlertRuleRequest) (response *DescribePrometheusAlertRuleResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusAlertRuleRequest()
	}
	response = NewDescribePrometheusAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterImageRequest() (request *CheckClusterImageRequest) {
	request = &CheckClusterImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckClusterImage")
	return
}

func NewCheckClusterImageResponse() (response *CheckClusterImageResponse) {
	response = &CheckClusterImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查镜像是否支持设置为集群镜像
func (c *Client) CheckClusterImage(request *CheckClusterImageRequest) (response *CheckClusterImageResponse, err error) {
	if request == nil {
		request = NewCheckClusterImageRequest()
	}
	response = NewCheckClusterImageResponse()
	err = c.Send(request, response)
	return
}

func NewAcquireTDCCClusterAdminRoleRequest() (request *AcquireTDCCClusterAdminRoleRequest) {
	request = &AcquireTDCCClusterAdminRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AcquireTDCCClusterAdminRole")
	return
}

func NewAcquireTDCCClusterAdminRoleResponse() (response *AcquireTDCCClusterAdminRoleResponse) {
	response = &AcquireTDCCClusterAdminRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过此接口，可以获取EKS集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
func (c *Client) AcquireTDCCClusterAdminRole(request *AcquireTDCCClusterAdminRoleRequest) (response *AcquireTDCCClusterAdminRoleResponse, err error) {
	if request == nil {
		request = NewAcquireTDCCClusterAdminRoleRequest()
	}
	response = NewAcquireTDCCClusterAdminRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageRegistryCredentialRequest() (request *CreateImageRegistryCredentialRequest) {
	request = &CreateImageRegistryCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateImageRegistryCredential")
	return
}

func NewCreateImageRegistryCredentialResponse() (response *CreateImageRegistryCredentialResponse) {
	response = &CreateImageRegistryCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像仓库凭证接口
func (c *Client) CreateImageRegistryCredential(request *CreateImageRegistryCredentialRequest) (response *CreateImageRegistryCredentialResponse, err error) {
	if request == nil {
		request = NewCreateImageRegistryCredentialRequest()
	}
	response = NewCreateImageRegistryCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRetainPodRequest() (request *DeleteRetainPodRequest) {
	request = &DeleteRetainPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteRetainPod")
	return
}

func NewDeleteRetainPodResponse() (response *DeleteRetainPodResponse) {
	response = &DeleteRetainPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定 Pod 唯一 ID 删除预留资源
func (c *Client) DeleteRetainPod(request *DeleteRetainPodRequest) (response *DeleteRetainPodResponse, err error) {
	if request == nil {
		request = NewDeleteRetainPodRequest()
	}
	response = NewDeleteRetainPodResponse()
	err = c.Send(request, response)
	return
}

func NewGetEksDashboardIDRequest() (request *GetEksDashboardIDRequest) {
	request = &GetEksDashboardIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEksDashboardID")
	return
}

func NewGetEksDashboardIDResponse() (response *GetEksDashboardIDResponse) {
	response = &GetEksDashboardIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据集群、仪表盘类型查询dashboardID信息
func (c *Client) GetEksDashboardID(request *GetEksDashboardIDRequest) (response *GetEksDashboardIDResponse, err error) {
	if request == nil {
		request = NewGetEksDashboardIDRequest()
	}
	response = NewGetEksDashboardIDResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterEndpointVipRequest() (request *CreateClusterEndpointVipRequest) {
	request = &CreateClusterEndpointVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpointVip")
	return
}

func NewCreateClusterEndpointVipResponse() (response *CreateClusterEndpointVipResponse) {
	response = &CreateClusterEndpointVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建托管集群外网访问端口（不再维护，准备下线）请使用新接口：CreateClusterEndpoint
func (c *Client) CreateClusterEndpointVip(request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
	if request == nil {
		request = NewCreateClusterEndpointVipRequest()
	}
	response = NewCreateClusterEndpointVipResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeClusterAuthorizationModeRequest() (request *UpgradeClusterAuthorizationModeRequest) {
	request = &UpgradeClusterAuthorizationModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterAuthorizationMode")
	return
}

func NewUpgradeClusterAuthorizationModeResponse() (response *UpgradeClusterAuthorizationModeResponse) {
	response = &UpgradeClusterAuthorizationModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级授权模式为RBAC，默认的授权模式是子账户拉取kubeconfig都为admin token，且访问集群内Kubernetes资源时，不做鉴权，升级为RBAC之后，可以提升集群安全性及权限的细粒度控制
func (c *Client) UpgradeClusterAuthorizationMode(request *UpgradeClusterAuthorizationModeRequest) (response *UpgradeClusterAuthorizationModeResponse, err error) {
	if request == nil {
		request = NewUpgradeClusterAuthorizationModeRequest()
	}
	response = NewUpgradeClusterAuthorizationModeResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusAlertPolicyRequest() (request *CreatePrometheusAlertPolicyRequest) {
	request = &CreatePrometheusAlertPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAlertPolicy")
	return
}

func NewCreatePrometheusAlertPolicyResponse() (response *CreatePrometheusAlertPolicyResponse) {
	response = &CreatePrometheusAlertPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警策略
func (c *Client) CreatePrometheusAlertPolicy(request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusAlertPolicyRequest()
	}
	response = NewCreatePrometheusAlertPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewEnableClusterDeletionProtectionRequest() (request *EnableClusterDeletionProtectionRequest) {
	request = &EnableClusterDeletionProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableClusterDeletionProtection")
	return
}

func NewEnableClusterDeletionProtectionResponse() (response *EnableClusterDeletionProtectionResponse) {
	response = &EnableClusterDeletionProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用集群删除保护
func (c *Client) EnableClusterDeletionProtection(request *EnableClusterDeletionProtectionRequest) (response *EnableClusterDeletionProtectionResponse, err error) {
	if request == nil {
		request = NewEnableClusterDeletionProtectionRequest()
	}
	response = NewEnableClusterDeletionProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCostAggregatedMetaInfoRequest() (request *DescribeCostAggregatedMetaInfoRequest) {
	request = &DescribeCostAggregatedMetaInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCostAggregatedMetaInfo")
	return
}

func NewDescribeCostAggregatedMetaInfoResponse() (response *DescribeCostAggregatedMetaInfoResponse) {
	response = &DescribeCostAggregatedMetaInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群成本聚合记录的元信息数据，包括某个聚合记录包含的容器的类型（pause、normal）以及容器是否设置了requestlimit等(NoRequestsNoLimits、NoRequestsHasLimits、HasRequestsNoLimits、HasRequestsHasLimits)
func (c *Client) DescribeCostAggregatedMetaInfo(request *DescribeCostAggregatedMetaInfoRequest) (response *DescribeCostAggregatedMetaInfoResponse, err error) {
	if request == nil {
		request = NewDescribeCostAggregatedMetaInfoRequest()
	}
	response = NewDescribeCostAggregatedMetaInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNamespacesRequest() (request *DescribeClusterNamespacesRequest) {
	request = &DescribeClusterNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNamespaces")
	return
}

func NewDescribeClusterNamespacesResponse() (response *DescribeClusterNamespacesResponse) {
	response = &DescribeClusterNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群命名空间列表信息，包括名称、状态、创建时间
func (c *Client) DescribeClusterNamespaces(request *DescribeClusterNamespacesRequest) (response *DescribeClusterNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNamespacesRequest()
	}
	response = NewDescribeClusterNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryCredentialsRequest() (request *DescribeImageRegistryCredentialsRequest) {
	request = &DescribeImageRegistryCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeImageRegistryCredentials")
	return
}

func NewDescribeImageRegistryCredentialsResponse() (response *DescribeImageRegistryCredentialsResponse) {
	response = &DescribeImageRegistryCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像仓库凭证接口
func (c *Client) DescribeImageRegistryCredentials(request *DescribeImageRegistryCredentialsRequest) (response *DescribeImageRegistryCredentialsResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryCredentialsRequest()
	}
	response = NewDescribeImageRegistryCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupStorageLocationsRequest() (request *DescribeBackupStorageLocationsRequest) {
	request = &DescribeBackupStorageLocationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeBackupStorageLocations")
	return
}

func NewDescribeBackupStorageLocationsResponse() (response *DescribeBackupStorageLocationsResponse) {
	response = &DescribeBackupStorageLocationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询备份仓库信息
func (c *Client) DescribeBackupStorageLocations(request *DescribeBackupStorageLocationsRequest) (response *DescribeBackupStorageLocationsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupStorageLocationsRequest()
	}
	response = NewDescribeBackupStorageLocationsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCIDRsRequest() (request *DescribeClusterCIDRsRequest) {
	request = &DescribeClusterCIDRsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCIDRs")
	return
}

func NewDescribeClusterCIDRsResponse() (response *DescribeClusterCIDRsResponse) {
	response = &DescribeClusterCIDRsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群VPC下使用的CIDR
func (c *Client) DescribeClusterCIDRs(request *DescribeClusterCIDRsRequest) (response *DescribeClusterCIDRsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCIDRsRequest()
	}
	response = NewDescribeClusterCIDRsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTkeEdgeAlarmSettingRequest() (request *DescribeTkeEdgeAlarmSettingRequest) {
	request = &DescribeTkeEdgeAlarmSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTkeEdgeAlarmSetting")
	return
}

func NewDescribeTkeEdgeAlarmSettingResponse() (response *DescribeTkeEdgeAlarmSettingResponse) {
	response = &DescribeTkeEdgeAlarmSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘集群列表的监控告警是否设置
func (c *Client) DescribeTkeEdgeAlarmSetting(request *DescribeTkeEdgeAlarmSettingRequest) (response *DescribeTkeEdgeAlarmSettingResponse, err error) {
	if request == nil {
		request = NewDescribeTkeEdgeAlarmSettingRequest()
	}
	response = NewDescribeTkeEdgeAlarmSettingResponse()
	err = c.Send(request, response)
	return
}

func NewEnableClusterAuditRequest() (request *EnableClusterAuditRequest) {
	request = &EnableClusterAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableClusterAudit")
	return
}

func NewEnableClusterAuditResponse() (response *EnableClusterAuditResponse) {
	response = &EnableClusterAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启集群审计
func (c *Client) EnableClusterAudit(request *EnableClusterAuditRequest) (response *EnableClusterAuditResponse, err error) {
	if request == nil {
		request = NewEnableClusterAuditRequest()
	}
	response = NewEnableClusterAuditResponse()
	err = c.Send(request, response)
	return
}

func NewRestartContainerRequest() (request *RestartContainerRequest) {
	request = &RestartContainerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RestartContainer")
	return
}

func NewRestartContainerResponse() (response *RestartContainerResponse) {
	response = &RestartContainerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启pod内的某个容器
func (c *Client) RestartContainer(request *RestartContainerRequest) (response *RestartContainerResponse, err error) {
	if request == nil {
		request = NewRestartContainerRequest()
	}
	response = NewRestartContainerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusRecordRulesRequest() (request *DescribePrometheusRecordRulesRequest) {
	request = &DescribePrometheusRecordRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusRecordRules")
	return
}

func NewDescribePrometheusRecordRulesResponse() (response *DescribePrometheusRecordRulesResponse) {
	response = &DescribePrometheusRecordRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合规则列表，包含关联集群内crd资源创建的record rule
func (c *Client) DescribePrometheusRecordRules(request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusRecordRulesRequest()
	}
	response = NewDescribePrometheusRecordRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTransferProgressRequest() (request *DescribePrometheusTransferProgressRequest) {
	request = &DescribePrometheusTransferProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTransferProgress")
	return
}

func NewDescribePrometheusTransferProgressResponse() (response *DescribePrometheusTransferProgressResponse) {
	response = &DescribePrometheusTransferProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取迁移任务执行状态
func (c *Client) DescribePrometheusTransferProgress(request *DescribePrometheusTransferProgressRequest) (response *DescribePrometheusTransferProgressResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTransferProgressRequest()
	}
	response = NewDescribePrometheusTransferProgressResponse()
	err = c.Send(request, response)
	return
}

func NewGetEksAppDiffRequest() (request *GetEksAppDiffRequest) {
	request = &GetEksAppDiffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEksAppDiff")
	return
}

func NewGetEksAppDiffResponse() (response *GetEksAppDiffResponse) {
	response = &GetEksAppDiffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取EKSApp的Diff信息
func (c *Client) GetEksAppDiff(request *GetEksAppDiffRequest) (response *GetEksAppDiffResponse, err error) {
	if request == nil {
		request = NewGetEksAppDiffRequest()
	}
	response = NewGetEksAppDiffResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteTableConflictsRequest() (request *DescribeRouteTableConflictsRequest) {
	request = &DescribeRouteTableConflictsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRouteTableConflicts")
	return
}

func NewDescribeRouteTableConflictsResponse() (response *DescribeRouteTableConflictsResponse) {
	response = &DescribeRouteTableConflictsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由表冲突列表
func (c *Client) DescribeRouteTableConflicts(request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
	if request == nil {
		request = NewDescribeRouteTableConflictsRequest()
	}
	response = NewDescribeRouteTableConflictsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodSpecQuotaRequest() (request *GetPodSpecQuotaRequest) {
	request = &GetPodSpecQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPodSpecQuota")
	return
}

func NewGetPodSpecQuotaResponse() (response *GetPodSpecQuotaResponse) {
	response = &GetPodSpecQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定规格的 Pod 配额
func (c *Client) GetPodSpecQuota(request *GetPodSpecQuotaRequest) (response *GetPodSpecQuotaResponse, err error) {
	if request == nil {
		request = NewGetPodSpecQuotaRequest()
	}
	response = NewGetPodSpecQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeClusterInstancesRequest() (request *DescribeEdgeClusterInstancesRequest) {
	request = &DescribeEdgeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterInstances")
	return
}

func NewDescribeEdgeClusterInstancesResponse() (response *DescribeEdgeClusterInstancesResponse) {
	response = &DescribeEdgeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘计算集群的节点信息
func (c *Client) DescribeEdgeClusterInstances(request *DescribeEdgeClusterInstancesRequest) (response *DescribeEdgeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeClusterInstancesRequest()
	}
	response = NewDescribeEdgeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSClusterStatusRequest() (request *DescribeEKSClusterStatusRequest) {
	request = &DescribeEKSClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterStatus")
	return
}

func NewDescribeEKSClusterStatusResponse() (response *DescribeEKSClusterStatusResponse) {
	response = &DescribeEKSClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性容器集群的当前状态以及过程信息
func (c *Client) DescribeEKSClusterStatus(request *DescribeEKSClusterStatusRequest) (response *DescribeEKSClusterStatusResponse, err error) {
	if request == nil {
		request = NewDescribeEKSClusterStatusRequest()
	}
	response = NewDescribeEKSClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetEKSAppUpgradeInfoRequest() (request *GetEKSAppUpgradeInfoRequest) {
	request = &GetEKSAppUpgradeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetEKSAppUpgradeInfo")
	return
}

func NewGetEKSAppUpgradeInfoResponse() (response *GetEKSAppUpgradeInfoResponse) {
	response = &GetEKSAppUpgradeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取EKS集群下组件升级信息
func (c *Client) GetEKSAppUpgradeInfo(request *GetEKSAppUpgradeInfoRequest) (response *GetEKSAppUpgradeInfoResponse, err error) {
	if request == nil {
		request = NewGetEKSAppUpgradeInfoRequest()
	}
	response = NewGetEKSAppUpgradeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewRestoreRemoteEtcdSnapshotRequest() (request *RestoreRemoteEtcdSnapshotRequest) {
	request = &RestoreRemoteEtcdSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RestoreRemoteEtcdSnapshot")
	return
}

func NewRestoreRemoteEtcdSnapshotResponse() (response *RestoreRemoteEtcdSnapshotResponse) {
	response = &RestoreRemoteEtcdSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复远程etcd快照
func (c *Client) RestoreRemoteEtcdSnapshot(request *RestoreRemoteEtcdSnapshotRequest) (response *RestoreRemoteEtcdSnapshotResponse, err error) {
	if request == nil {
		request = NewRestoreRemoteEtcdSnapshotRequest()
	}
	response = NewRestoreRemoteEtcdSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmPoliciesRequest() (request *DescribeAlarmPoliciesRequest) {
	request = &DescribeAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeAlarmPolicies")
	return
}

func NewDescribeAlarmPoliciesResponse() (response *DescribeAlarmPoliciesResponse) {
	response = &DescribeAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警策略列表
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmPoliciesRequest()
	}
	response = NewDescribeAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewListTDCCClusterCertificatesRequest() (request *ListTDCCClusterCertificatesRequest) {
	request = &ListTDCCClusterCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListTDCCClusterCertificates")
	return
}

func NewListTDCCClusterCertificatesResponse() (response *ListTDCCClusterCertificatesResponse) {
	response = &ListTDCCClusterCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的证书列表， 返回集群所有账户的证书列表
func (c *Client) ListTDCCClusterCertificates(request *ListTDCCClusterCertificatesRequest) (response *ListTDCCClusterCertificatesResponse, err error) {
	if request == nil {
		request = NewListTDCCClusterCertificatesRequest()
	}
	response = NewListTDCCClusterCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewRestartEKSContainerInstancesRequest() (request *RestartEKSContainerInstancesRequest) {
	request = &RestartEKSContainerInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RestartEKSContainerInstances")
	return
}

func NewRestartEKSContainerInstancesResponse() (response *RestartEKSContainerInstancesResponse) {
	response = &RestartEKSContainerInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启弹性容器实例，支持批量操作
func (c *Client) RestartEKSContainerInstances(request *RestartEKSContainerInstancesRequest) (response *RestartEKSContainerInstancesResponse, err error) {
	if request == nil {
		request = NewRestartEKSContainerInstancesRequest()
	}
	response = NewRestartEKSContainerInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteECMEKSClusterRequest() (request *DeleteECMEKSClusterRequest) {
	request = &DeleteECMEKSClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteECMEKSCluster")
	return
}

func NewDeleteECMEKSClusterResponse() (response *DeleteECMEKSClusterResponse) {
	response = &DeleteECMEKSClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ECM弹性集群
func (c *Client) DeleteECMEKSCluster(request *DeleteECMEKSClusterRequest) (response *DeleteECMEKSClusterResponse, err error) {
	if request == nil {
		request = NewDeleteECMEKSClusterRequest()
	}
	response = NewDeleteECMEKSClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodsRequest() (request *GetPodsRequest) {
	request = &GetPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPods")
	return
}

func NewGetPodsResponse() (response *GetPodsResponse) {
	response = &GetPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群内Pod
func (c *Client) GetPods(request *GetPodsRequest) (response *GetPodsResponse, err error) {
	if request == nil {
		request = NewGetPodsRequest()
	}
	response = NewGetPodsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterAsGroupAttributeRequest() (request *ModifyClusterAsGroupAttributeRequest) {
	request = &ModifyClusterAsGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupAttribute")
	return
}

func NewModifyClusterAsGroupAttributeResponse() (response *ModifyClusterAsGroupAttributeResponse) {
	response = &ModifyClusterAsGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群伸缩组属性
func (c *Client) ModifyClusterAsGroupAttribute(request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyClusterAsGroupAttributeRequest()
	}
	response = NewModifyClusterAsGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTkeEdgeAlarmPolicyRequest() (request *ModifyTkeEdgeAlarmPolicyRequest) {
	request = &ModifyTkeEdgeAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyTkeEdgeAlarmPolicy")
	return
}

func NewModifyTkeEdgeAlarmPolicyResponse() (response *ModifyTkeEdgeAlarmPolicyResponse) {
	response = &ModifyTkeEdgeAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改边缘集群告警策略
func (c *Client) ModifyTkeEdgeAlarmPolicy(request *ModifyTkeEdgeAlarmPolicyRequest) (response *ModifyTkeEdgeAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewModifyTkeEdgeAlarmPolicyRequest()
	}
	response = NewModifyTkeEdgeAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEksLogConfigRequest() (request *CreateEksLogConfigRequest) {
	request = &CreateEksLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEksLogConfig")
	return
}

func NewCreateEksLogConfigResponse() (response *CreateEksLogConfigResponse) {
	response = &CreateEksLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为弹性集群创建日志采集配置
func (c *Client) CreateEksLogConfig(request *CreateEksLogConfigRequest) (response *CreateEksLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateEksLogConfigRequest()
	}
	response = NewCreateEksLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusTempRequest() (request *DeletePrometheusTempRequest) {
	request = &DeletePrometheusTempRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemp")
	return
}

func NewDeletePrometheusTempResponse() (response *DeletePrometheusTempResponse) {
	response = &DeletePrometheusTempResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除一个云原生Prometheus配置模板
func (c *Client) DeletePrometheusTemp(request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusTempRequest()
	}
	response = NewDeletePrometheusTempResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServicesRequest() (request *DescribeClusterServicesRequest) {
	request = &DescribeClusterServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterServices")
	return
}

func NewDescribeClusterServicesResponse() (response *DescribeClusterServicesResponse) {
	response = &DescribeClusterServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口获取集群内Service相关的详细描述信息，参考kubernetes API获取Service，只对内部短期使用
func (c *Client) DescribeClusterServices(request *DescribeClusterServicesRequest) (response *DescribeClusterServicesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServicesRequest()
	}
	response = NewDescribeClusterServicesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccountTypeRequest() (request *GetAccountTypeRequest) {
	request = &GetAccountTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetAccountType")
	return
}

func NewGetAccountTypeResponse() (response *GetAccountTypeResponse) {
	response = &GetAccountTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主账号类型
func (c *Client) GetAccountType(request *GetAccountTypeRequest) (response *GetAccountTypeResponse, err error) {
	if request == nil {
		request = NewGetAccountTypeRequest()
	}
	response = NewGetAccountTypeResponse()
	err = c.Send(request, response)
	return
}

func NewProbePrometheusTargetsRequest() (request *ProbePrometheusTargetsRequest) {
	request = &ProbePrometheusTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ProbePrometheusTargets")
	return
}

func NewProbePrometheusTargetsResponse() (response *ProbePrometheusTargetsResponse) {
	response = &ProbePrometheusTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 探测采集对象
func (c *Client) ProbePrometheusTargets(request *ProbePrometheusTargetsRequest) (response *ProbePrometheusTargetsResponse, err error) {
	if request == nil {
		request = NewProbePrometheusTargetsRequest()
	}
	response = NewProbePrometheusTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLogCollectorNameRequest() (request *CheckLogCollectorNameRequest) {
	request = &CheckLogCollectorNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckLogCollectorName")
	return
}

func NewCheckLogCollectorNameResponse() (response *CheckLogCollectorNameResponse) {
	response = &CheckLogCollectorNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过名称检查日志采集规则是否存在
func (c *Client) CheckLogCollectorName(request *CheckLogCollectorNameRequest) (response *CheckLogCollectorNameResponse, err error) {
	if request == nil {
		request = NewCheckLogCollectorNameRequest()
	}
	response = NewCheckLogCollectorNameResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEventPersistenceRequest() (request *EnableEventPersistenceRequest) {
	request = &EnableEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEventPersistence")
	return
}

func NewEnableEventPersistenceResponse() (response *EnableEventPersistenceResponse) {
	response = &EnableEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启事件持久化功能
func (c *Client) EnableEventPersistence(request *EnableEventPersistenceRequest) (response *EnableEventPersistenceResponse, err error) {
	if request == nil {
		request = NewEnableEventPersistenceRequest()
	}
	response = NewEnableEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTDCCHubClusterRequest() (request *CreateTDCCHubClusterRequest) {
	request = &CreateTDCCHubClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateTDCCHubCluster")
	return
}

func NewCreateTDCCHubClusterResponse() (response *CreateTDCCHubClusterResponse) {
	response = &CreateTDCCHubClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TDCC Hub集群
func (c *Client) CreateTDCCHubCluster(request *CreateTDCCHubClusterRequest) (response *CreateTDCCHubClusterResponse, err error) {
	if request == nil {
		request = NewCreateTDCCHubClusterRequest()
	}
	response = NewCreateTDCCHubClusterResponse()
	err = c.Send(request, response)
	return
}

func NewQueryOverageRequest() (request *QueryOverageRequest) {
	request = &QueryOverageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "QueryOverage")
	return
}

func NewQueryOverageResponse() (response *QueryOverageResponse) {
	response = &QueryOverageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前TKE管控集群的cpu总核数是否超过限制
func (c *Client) QueryOverage(request *QueryOverageRequest) (response *QueryOverageResponse, err error) {
	if request == nil {
		request = NewQueryOverageRequest()
	}
	response = NewQueryOverageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEKSClusterCredentialRequest() (request *DescribeEKSClusterCredentialRequest) {
	request = &DescribeEKSClusterCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEKSClusterCredential")
	return
}

func NewDescribeEKSClusterCredentialResponse() (response *DescribeEKSClusterCredentialResponse) {
	response = &DescribeEKSClusterCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性容器集群的接入认证信息
func (c *Client) DescribeEKSClusterCredential(request *DescribeEKSClusterCredentialRequest) (response *DescribeEKSClusterCredentialResponse, err error) {
	if request == nil {
		request = NewDescribeEKSClusterCredentialRequest()
	}
	response = NewDescribeEKSClusterCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubnetVipRequest() (request *GetSubnetVipRequest) {
	request = &GetSubnetVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetSubnetVip")
	return
}

func NewGetSubnetVipResponse() (response *GetSubnetVipResponse) {
	response = &GetSubnetVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子网内预留VIP
func (c *Client) GetSubnetVip(request *GetSubnetVipRequest) (response *GetSubnetVipResponse, err error) {
	if request == nil {
		request = NewGetSubnetVipRequest()
	}
	response = NewGetSubnetVipResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterVirtualNodePoolRequest() (request *DeleteClusterVirtualNodePoolRequest) {
	request = &DeleteClusterVirtualNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNodePool")
	return
}

func NewDeleteClusterVirtualNodePoolResponse() (response *DeleteClusterVirtualNodePoolResponse) {
	response = &DeleteClusterVirtualNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除超级节点池
func (c *Client) DeleteClusterVirtualNodePool(request *DeleteClusterVirtualNodePoolRequest) (response *DeleteClusterVirtualNodePoolResponse, err error) {
	if request == nil {
		request = NewDeleteClusterVirtualNodePoolRequest()
	}
	response = NewDeleteClusterVirtualNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGrafanaAdminPasswordRequest() (request *ModifyGrafanaAdminPasswordRequest) {
	request = &ModifyGrafanaAdminPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyGrafanaAdminPassword")
	return
}

func NewModifyGrafanaAdminPasswordResponse() (response *ModifyGrafanaAdminPasswordResponse) {
	response = &ModifyGrafanaAdminPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Grafana管理员账号和密码
func (c *Client) ModifyGrafanaAdminPassword(request *ModifyGrafanaAdminPasswordRequest) (response *ModifyGrafanaAdminPasswordResponse, err error) {
	if request == nil {
		request = NewModifyGrafanaAdminPasswordRequest()
	}
	response = NewModifyGrafanaAdminPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterAsGroupOptionAttributeRequest() (request *ModifyClusterAsGroupOptionAttributeRequest) {
	request = &ModifyClusterAsGroupOptionAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupOptionAttribute")
	return
}

func NewModifyClusterAsGroupOptionAttributeResponse() (response *ModifyClusterAsGroupOptionAttributeResponse) {
	response = &ModifyClusterAsGroupOptionAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群弹性伸缩属性
func (c *Client) ModifyClusterAsGroupOptionAttribute(request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
	if request == nil {
		request = NewModifyClusterAsGroupOptionAttributeRequest()
	}
	response = NewModifyClusterAsGroupOptionAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeInstanceProgressRequest() (request *GetUpgradeInstanceProgressRequest) {
	request = &GetUpgradeInstanceProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeInstanceProgress")
	return
}

func NewGetUpgradeInstanceProgressResponse() (response *GetUpgradeInstanceProgressResponse) {
	response = &GetUpgradeInstanceProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得节点升级当前的进度，若集群未处于节点升级状态，则接口会报错：任务未找到。
func (c *Client) GetUpgradeInstanceProgress(request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
	if request == nil {
		request = NewGetUpgradeInstanceProgressRequest()
	}
	response = NewGetUpgradeInstanceProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSecurityGroupRequest() (request *DescribeClusterSecurityGroupRequest) {
	request = &DescribeClusterSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSecurityGroup")
	return
}

func NewDescribeClusterSecurityGroupResponse() (response *DescribeClusterSecurityGroupResponse) {
	response = &DescribeClusterSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群安全组信息
func (c *Client) DescribeClusterSecurityGroup(request *DescribeClusterSecurityGroupRequest) (response *DescribeClusterSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSecurityGroupRequest()
	}
	response = NewDescribeClusterSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusAgentsRequest() (request *DescribePrometheusAgentsRequest) {
	request = &DescribePrometheusAgentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAgents")
	return
}

func NewDescribePrometheusAgentsResponse() (response *DescribePrometheusAgentsResponse) {
	response = &DescribePrometheusAgentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取被关联集群列表
func (c *Client) DescribePrometheusAgents(request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusAgentsRequest()
	}
	response = NewDescribePrometheusAgentsResponse()
	err = c.Send(request, response)
	return
}

func NewDrainExternalNodeRequest() (request *DrainExternalNodeRequest) {
	request = &DrainExternalNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DrainExternalNode")
	return
}

func NewDrainExternalNodeResponse() (response *DrainExternalNodeResponse) {
	response = &DrainExternalNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐第三方节点
func (c *Client) DrainExternalNode(request *DrainExternalNodeRequest) (response *DrainExternalNodeResponse, err error) {
	if request == nil {
		request = NewDrainExternalNodeRequest()
	}
	response = NewDrainExternalNodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetECMPodRequest() (request *GetECMPodRequest) {
	request = &GetECMPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetECMPod")
	return
}

func NewGetECMPodResponse() (response *GetECMPodResponse) {
	response = &GetECMPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ECMpod信息
func (c *Client) GetECMPod(request *GetECMPodRequest) (response *GetECMPodResponse, err error) {
	if request == nil {
		request = NewGetECMPodRequest()
	}
	response = NewGetECMPodResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEdgeSSHRequest() (request *UpdateEdgeSSHRequest) {
	request = &UpdateEdgeSSHRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEdgeSSH")
	return
}

func NewUpdateEdgeSSHResponse() (response *UpdateEdgeSSHResponse) {
	response = &UpdateEdgeSSHResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 控制边缘集群是否支持SSH登录内网节点
func (c *Client) UpdateEdgeSSH(request *UpdateEdgeSSHRequest) (response *UpdateEdgeSSHResponse, err error) {
	if request == nil {
		request = NewUpdateEdgeSSHRequest()
	}
	response = NewUpdateEdgeSSHResponse()
	err = c.Send(request, response)
	return
}

func NewValidateClusterAddVirtualNodeConditionsRequest() (request *ValidateClusterAddVirtualNodeConditionsRequest) {
	request = &ValidateClusterAddVirtualNodeConditionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ValidateClusterAddVirtualNodeConditions")
	return
}

func NewValidateClusterAddVirtualNodeConditionsResponse() (response *ValidateClusterAddVirtualNodeConditionsResponse) {
	response = &ValidateClusterAddVirtualNodeConditionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验集群添加超级节点条件
func (c *Client) ValidateClusterAddVirtualNodeConditions(request *ValidateClusterAddVirtualNodeConditionsRequest) (response *ValidateClusterAddVirtualNodeConditionsResponse, err error) {
	if request == nil {
		request = NewValidateClusterAddVirtualNodeConditionsRequest()
	}
	response = NewValidateClusterAddVirtualNodeConditionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSchedulerPolicyRequest() (request *DescribeClusterSchedulerPolicyRequest) {
	request = &DescribeClusterSchedulerPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSchedulerPolicy")
	return
}

func NewDescribeClusterSchedulerPolicyResponse() (response *DescribeClusterSchedulerPolicyResponse) {
	response = &DescribeClusterSchedulerPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群调度策略
func (c *Client) DescribeClusterSchedulerPolicy(request *DescribeClusterSchedulerPolicyRequest) (response *DescribeClusterSchedulerPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSchedulerPolicyRequest()
	}
	response = NewDescribeClusterSchedulerPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaRequest() (request *GetQuotaRequest) {
	request = &GetQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetQuota")
	return
}

func NewGetQuotaResponse() (response *GetQuotaResponse) {
	response = &GetQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetQuota
func (c *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
	if request == nil {
		request = NewGetQuotaRequest()
	}
	response = NewGetQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewCheckComponentVersionRequest() (request *CheckComponentVersionRequest) {
	request = &CheckComponentVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckComponentVersion")
	return
}

func NewCheckComponentVersionResponse() (response *CheckComponentVersionResponse) {
	response = &CheckComponentVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验当前集群的各组件是否满足指定产品的最低版本要求
func (c *Client) CheckComponentVersion(request *CheckComponentVersionRequest) (response *CheckComponentVersionResponse, err error) {
	if request == nil {
		request = NewCheckComponentVersionRequest()
	}
	response = NewCheckComponentVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateService")
	return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Service
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	if request == nil {
		request = NewCreateServiceRequest()
	}
	response = NewCreateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageCachesRequest() (request *DeleteImageCachesRequest) {
	request = &DeleteImageCachesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteImageCaches")
	return
}

func NewDeleteImageCachesResponse() (response *DeleteImageCachesResponse) {
	response = &DeleteImageCachesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除镜像缓存
func (c *Client) DeleteImageCaches(request *DeleteImageCachesRequest) (response *DeleteImageCachesResponse, err error) {
	if request == nil {
		request = NewDeleteImageCachesRequest()
	}
	response = NewDeleteImageCachesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusAlertPolicyRequest() (request *DescribePrometheusAlertPolicyRequest) {
	request = &DescribePrometheusAlertPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusAlertPolicy")
	return
}

func NewDescribePrometheusAlertPolicyResponse() (response *DescribePrometheusAlertPolicyResponse) {
	response = &DescribePrometheusAlertPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取2.0实例告警策略列表
func (c *Client) DescribePrometheusAlertPolicy(request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusAlertPolicyRequest()
	}
	response = NewDescribePrometheusAlertPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPodSecurityGroupsRequest() (request *ModifyPodSecurityGroupsRequest) {
	request = &ModifyPodSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPodSecurityGroups")
	return
}

func NewModifyPodSecurityGroupsResponse() (response *ModifyPodSecurityGroupsResponse) {
	response = &ModifyPodSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安全组
func (c *Client) ModifyPodSecurityGroups(request *ModifyPodSecurityGroupsRequest) (response *ModifyPodSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewModifyPodSecurityGroupsRequest()
	}
	response = NewModifyPodSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
	request = &DeleteSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteSecret")
	return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
	response = &DeleteSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除秘钥
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
	if request == nil {
		request = NewDeleteSecretRequest()
	}
	response = NewDeleteSecretResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMetricsDataRequest() (request *DescribeClusterMetricsDataRequest) {
	request = &DescribeClusterMetricsDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterMetricsData")
	return
}

func NewDescribeClusterMetricsDataResponse() (response *DescribeClusterMetricsDataResponse) {
	response = &DescribeClusterMetricsDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的监控数据，通过指定的指标名称。
func (c *Client) DescribeClusterMetricsData(request *DescribeClusterMetricsDataRequest) (response *DescribeClusterMetricsDataResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMetricsDataRequest()
	}
	response = NewDescribeClusterMetricsDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterAsGroupsRequest() (request *DeleteClusterAsGroupsRequest) {
	request = &DeleteClusterAsGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterAsGroups")
	return
}

func NewDeleteClusterAsGroupsResponse() (response *DeleteClusterAsGroupsResponse) {
	response = &DeleteClusterAsGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群伸缩组
func (c *Client) DeleteClusterAsGroups(request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteClusterAsGroupsRequest()
	}
	response = NewDeleteClusterAsGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackClusterReleaseRequest() (request *RollbackClusterReleaseRequest) {
	request = &RollbackClusterReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RollbackClusterRelease")
	return
}

func NewRollbackClusterReleaseResponse() (response *RollbackClusterReleaseResponse) {
	response = &RollbackClusterReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在应用市场中集群回滚应用至某个历史版本
func (c *Client) RollbackClusterRelease(request *RollbackClusterReleaseRequest) (response *RollbackClusterReleaseResponse, err error) {
	if request == nil {
		request = NewRollbackClusterReleaseRequest()
	}
	response = NewRollbackClusterReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdAvailableVersionsRequest() (request *DescribeEtcdAvailableVersionsRequest) {
	request = &DescribeEtcdAvailableVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdAvailableVersions")
	return
}

func NewDescribeEtcdAvailableVersionsResponse() (response *DescribeEtcdAvailableVersionsResponse) {
	response = &DescribeEtcdAvailableVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd可用版本
func (c *Client) DescribeEtcdAvailableVersions(request *DescribeEtcdAvailableVersionsRequest) (response *DescribeEtcdAvailableVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdAvailableVersionsRequest()
	}
	response = NewDescribeEtcdAvailableVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdSnapshotPoliciesRequest() (request *DescribeEtcdSnapshotPoliciesRequest) {
	request = &DescribeEtcdSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdSnapshotPolicies")
	return
}

func NewDescribeEtcdSnapshotPoliciesResponse() (response *DescribeEtcdSnapshotPoliciesResponse) {
	response = &DescribeEtcdSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd快照策略
func (c *Client) DescribeEtcdSnapshotPolicies(request *DescribeEtcdSnapshotPoliciesRequest) (response *DescribeEtcdSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdSnapshotPoliciesRequest()
	}
	response = NewDescribeEtcdSnapshotPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodsBySpecRequest() (request *DescribePodsBySpecRequest) {
	request = &DescribePodsBySpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePodsBySpec")
	return
}

func NewDescribePodsBySpecResponse() (response *DescribePodsBySpecResponse) {
	response = &DescribePodsBySpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可以用预留券抵扣的 Pod 信息。
func (c *Client) DescribePodsBySpec(request *DescribePodsBySpecRequest) (response *DescribePodsBySpecResponse, err error) {
	if request == nil {
		request = NewDescribePodsBySpecRequest()
	}
	response = NewDescribePodsBySpecResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOSImagesRequest() (request *DescribeOSImagesRequest) {
	request = &DescribeOSImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeOSImages")
	return
}

func NewDescribeOSImagesResponse() (response *DescribeOSImagesResponse) {
	response = &DescribeOSImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取OS聚合信息
func (c *Client) DescribeOSImages(request *DescribeOSImagesRequest) (response *DescribeOSImagesResponse, err error) {
	if request == nil {
		request = NewDescribeOSImagesRequest()
	}
	response = NewDescribeOSImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCloudRunEventPersistenceRequest() (request *DisableCloudRunEventPersistenceRequest) {
	request = &DisableCloudRunEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableCloudRunEventPersistence")
	return
}

func NewDisableCloudRunEventPersistenceResponse() (response *DisableCloudRunEventPersistenceResponse) {
	response = &DisableCloudRunEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭cloudrun事件持久化功能
func (c *Client) DisableCloudRunEventPersistence(request *DisableCloudRunEventPersistenceRequest) (response *DisableCloudRunEventPersistenceResponse, err error) {
	if request == nil {
		request = NewDisableCloudRunEventPersistenceRequest()
	}
	response = NewDisableCloudRunEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCExternalClusterSpecRequest() (request *DescribeTDCCExternalClusterSpecRequest) {
	request = &DescribeTDCCExternalClusterSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCExternalClusterSpec")
	return
}

func NewDescribeTDCCExternalClusterSpecResponse() (response *DescribeTDCCExternalClusterSpecResponse) {
	response = &DescribeTDCCExternalClusterSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导入第三方集群YAML定义
func (c *Client) DescribeTDCCExternalClusterSpec(request *DescribeTDCCExternalClusterSpecRequest) (response *DescribeTDCCExternalClusterSpecResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCExternalClusterSpecRequest()
	}
	response = NewDescribeTDCCExternalClusterSpecResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusInstanceInitStatusRequest() (request *DescribePrometheusInstanceInitStatusRequest) {
	request = &DescribePrometheusInstanceInitStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstanceInitStatus")
	return
}

func NewDescribePrometheusInstanceInitStatusResponse() (response *DescribePrometheusInstanceInitStatusResponse) {
	response = &DescribePrometheusInstanceInitStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取2.0实例初始化任务状态
func (c *Client) DescribePrometheusInstanceInitStatus(request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusInstanceInitStatusRequest()
	}
	response = NewDescribePrometheusInstanceInitStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterAuthStatusRequest() (request *UpdateClusterAuthStatusRequest) {
	request = &UpdateClusterAuthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterAuthStatus")
	return
}

func NewUpdateClusterAuthStatusResponse() (response *UpdateClusterAuthStatusResponse) {
	response = &UpdateClusterAuthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启之后在子账户被删除7天之后会自动清理改子账户在集群内的权限和证书信息，关闭之后则不会自动清理，需要用户手动在控制台一键清理
func (c *Client) UpdateClusterAuthStatus(request *UpdateClusterAuthStatusRequest) (response *UpdateClusterAuthStatusResponse, err error) {
	if request == nil {
		request = NewUpdateClusterAuthStatusRequest()
	}
	response = NewUpdateClusterAuthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTKEEdgeExternalKubeconfigRequest() (request *UpdateTKEEdgeExternalKubeconfigRequest) {
	request = &UpdateTKEEdgeExternalKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateTKEEdgeExternalKubeconfig")
	return
}

func NewUpdateTKEEdgeExternalKubeconfigResponse() (response *UpdateTKEEdgeExternalKubeconfigResponse) {
	response = &UpdateTKEEdgeExternalKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新边缘计算外部访问的kubeconfig
func (c *Client) UpdateTKEEdgeExternalKubeconfig(request *UpdateTKEEdgeExternalKubeconfigRequest) (response *UpdateTKEEdgeExternalKubeconfigResponse, err error) {
	if request == nil {
		request = NewUpdateTKEEdgeExternalKubeconfigRequest()
	}
	response = NewUpdateTKEEdgeExternalKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifySuperNodeTagRequest() (request *ModifySuperNodeTagRequest) {
	request = &ModifySuperNodeTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifySuperNodeTag")
	return
}

func NewModifySuperNodeTagResponse() (response *ModifySuperNodeTagResponse) {
	response = &ModifySuperNodeTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改包年包月超级节点的标签
func (c *Client) ModifySuperNodeTag(request *ModifySuperNodeTagRequest) (response *ModifySuperNodeTagResponse, err error) {
	if request == nil {
		request = NewModifySuperNodeTagRequest()
	}
	response = NewModifySuperNodeTagResponse()
	err = c.Send(request, response)
	return
}

func NewRenewReservedInstancesRequest() (request *RenewReservedInstancesRequest) {
	request = &RenewReservedInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RenewReservedInstances")
	return
}

func NewRenewReservedInstancesResponse() (response *RenewReservedInstancesResponse) {
	response = &RenewReservedInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 续费时请确保账户余额充足。
func (c *Client) RenewReservedInstances(request *RenewReservedInstancesRequest) (response *RenewReservedInstancesResponse, err error) {
	if request == nil {
		request = NewRenewReservedInstancesRequest()
	}
	response = NewRenewReservedInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEKSContainerInstancesRequest() (request *DeleteEKSContainerInstancesRequest) {
	request = &DeleteEKSContainerInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEKSContainerInstances")
	return
}

func NewDeleteEKSContainerInstancesResponse() (response *DeleteEKSContainerInstancesResponse) {
	response = &DeleteEKSContainerInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容器实例，可批量删除
func (c *Client) DeleteEKSContainerInstances(request *DeleteEKSContainerInstancesRequest) (response *DeleteEKSContainerInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteEKSContainerInstancesRequest()
	}
	response = NewDeleteEKSContainerInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterGlobalStatRequest() (request *DescribeClusterGlobalStatRequest) {
	request = &DescribeClusterGlobalStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterGlobalStat")
	return
}

func NewDescribeClusterGlobalStatResponse() (response *DescribeClusterGlobalStatResponse) {
	response = &DescribeClusterGlobalStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局集群统计数据，包括集群数、节点数、工作负载数
func (c *Client) DescribeClusterGlobalStat(request *DescribeClusterGlobalStatRequest) (response *DescribeClusterGlobalStatResponse, err error) {
	if request == nil {
		request = NewDescribeClusterGlobalStatRequest()
	}
	response = NewDescribeClusterGlobalStatResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePodRequest() (request *DeletePodRequest) {
	request = &DeletePodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePod")
	return
}

func NewDeletePodResponse() (response *DeletePodResponse) {
	response = &DeletePodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定pod
func (c *Client) DeletePod(request *DeletePodRequest) (response *DeletePodResponse, err error) {
	if request == nil {
		request = NewDeletePodRequest()
	}
	response = NewDeletePodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdTasksRequest() (request *DescribeEtcdTasksRequest) {
	request = &DescribeEtcdTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdTasks")
	return
}

func NewDescribeEtcdTasksResponse() (response *DescribeEtcdTasksResponse) {
	response = &DescribeEtcdTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd相关tasks
func (c *Client) DescribeEtcdTasks(request *DescribeEtcdTasksRequest) (response *DescribeEtcdTasksResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdTasksRequest()
	}
	response = NewDescribeEtcdTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAddonRequest() (request *DeleteAddonRequest) {
	request = &DeleteAddonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteAddon")
	return
}

func NewDeleteAddonResponse() (response *DeleteAddonResponse) {
	response = &DeleteAddonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除一个addon
func (c *Client) DeleteAddon(request *DeleteAddonRequest) (response *DeleteAddonResponse, err error) {
	if request == nil {
		request = NewDeleteAddonRequest()
	}
	response = NewDeleteAddonResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInspectionReportRequest() (request *DescribeClusterInspectionReportRequest) {
	request = &DescribeClusterInspectionReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInspectionReport")
	return
}

func NewDescribeClusterInspectionReportResponse() (response *DescribeClusterInspectionReportResponse) {
	response = &DescribeClusterInspectionReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得集群巡检报告页详情
func (c *Client) DescribeClusterInspectionReport(request *DescribeClusterInspectionReportRequest) (response *DescribeClusterInspectionReportResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInspectionReportRequest()
	}
	response = NewDescribeClusterInspectionReportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksTaskResultRequest() (request *DescribeEksTaskResultRequest) {
	request = &DescribeEksTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksTaskResult")
	return
}

func NewDescribeEksTaskResultResponse() (response *DescribeEksTaskResultResponse) {
	response = &DescribeEksTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询异步任务的执行结果。
func (c *Client) DescribeEksTaskResult(request *DescribeEksTaskResultRequest) (response *DescribeEksTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeEksTaskResultRequest()
	}
	response = NewDescribeEksTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEksEventPersistenceRequest() (request *EnableEksEventPersistenceRequest) {
	request = &EnableEksEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEksEventPersistence")
	return
}

func NewEnableEksEventPersistenceResponse() (response *EnableEksEventPersistenceResponse) {
	response = &EnableEksEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Eks集群开启事件持久化功能
func (c *Client) EnableEksEventPersistence(request *EnableEksEventPersistenceRequest) (response *EnableEksEventPersistenceResponse, err error) {
	if request == nil {
		request = NewEnableEksEventPersistenceRequest()
	}
	response = NewEnableEksEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudRunEventRequest() (request *DescribeCloudRunEventRequest) {
	request = &DescribeCloudRunEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCloudRunEvent")
	return
}

func NewDescribeCloudRunEventResponse() (response *DescribeCloudRunEventResponse) {
	response = &DescribeCloudRunEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询event列表
func (c *Client) DescribeCloudRunEvent(request *DescribeCloudRunEventRequest) (response *DescribeCloudRunEventResponse, err error) {
	if request == nil {
		request = NewDescribeCloudRunEventRequest()
	}
	response = NewDescribeCloudRunEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInstanceIdsRequest() (request *DescribeClusterInstanceIdsRequest) {
	request = &DescribeClusterInstanceIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstanceIds")
	return
}

func NewDescribeClusterInstanceIdsResponse() (response *DescribeClusterInstanceIdsResponse) {
	response = &DescribeClusterInstanceIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群节点ID列表【仅内部使用】
func (c *Client) DescribeClusterInstanceIds(request *DescribeClusterInstanceIdsRequest) (response *DescribeClusterInstanceIdsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInstanceIdsRequest()
	}
	response = NewDescribeClusterInstanceIdsResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterCIDRRequest() (request *AddClusterCIDRRequest) {
	request = &AddClusterCIDRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddClusterCIDR")
	return
}

func NewAddClusterCIDRResponse() (response *AddClusterCIDRResponse) {
	response = &AddClusterCIDRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给GR集群增加可用的ClusterCIDR（开白才能使用此功能，如需要请联系我们）
func (c *Client) AddClusterCIDR(request *AddClusterCIDRRequest) (response *AddClusterCIDRResponse, err error) {
	if request == nil {
		request = NewAddClusterCIDRRequest()
	}
	response = NewAddClusterCIDRResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceQuotaRequest() (request *DescribeResourceQuotaRequest) {
	request = &DescribeResourceQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeResourceQuota")
	return
}

func NewDescribeResourceQuotaResponse() (response *DescribeResourceQuotaResponse) {
	response = &DescribeResourceQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询虚拟集群所属资源配额信息
func (c *Client) DescribeResourceQuota(request *DescribeResourceQuotaRequest) (response *DescribeResourceQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeResourceQuotaRequest()
	}
	response = NewDescribeResourceQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCExternalClustersRequest() (request *DescribeTDCCExternalClustersRequest) {
	request = &DescribeTDCCExternalClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCExternalClusters")
	return
}

func NewDescribeTDCCExternalClustersResponse() (response *DescribeTDCCExternalClustersResponse) {
	response = &DescribeTDCCExternalClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询注册集群列表
func (c *Client) DescribeTDCCExternalClusters(request *DescribeTDCCExternalClustersRequest) (response *DescribeTDCCExternalClustersResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCExternalClustersRequest()
	}
	response = NewDescribeTDCCExternalClustersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTKEEdgeClusterRequest() (request *CreateTKEEdgeClusterRequest) {
	request = &CreateTKEEdgeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateTKEEdgeCluster")
	return
}

func NewCreateTKEEdgeClusterResponse() (response *CreateTKEEdgeClusterResponse) {
	response = &CreateTKEEdgeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建边缘计算集群
func (c *Client) CreateTKEEdgeCluster(request *CreateTKEEdgeClusterRequest) (response *CreateTKEEdgeClusterResponse, err error) {
	if request == nil {
		request = NewCreateTKEEdgeClusterRequest()
	}
	response = NewCreateTKEEdgeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePodsByIdRequest() (request *DeletePodsByIdRequest) {
	request = &DeletePodsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePodsById")
	return
}

func NewDeletePodsByIdResponse() (response *DeletePodsByIdResponse) {
	response = &DeletePodsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定eksId删除pod
func (c *Client) DeletePodsById(request *DeletePodsByIdRequest) (response *DeletePodsByIdResponse, err error) {
	if request == nil {
		request = NewDeletePodsByIdRequest()
	}
	response = NewDeletePodsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudRunPodRequest() (request *DescribeCloudRunPodRequest) {
	request = &DescribeCloudRunPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeCloudRunPod")
	return
}

func NewDescribeCloudRunPodResponse() (response *DescribeCloudRunPodResponse) {
	response = &DescribeCloudRunPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CloudRun Pod详情
func (c *Client) DescribeCloudRunPod(request *DescribeCloudRunPodRequest) (response *DescribeCloudRunPodResponse, err error) {
	if request == nil {
		request = NewDescribeCloudRunPodRequest()
	}
	response = NewDescribeCloudRunPodResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusConfigRequest() (request *ModifyPrometheusConfigRequest) {
	request = &ModifyPrometheusConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusConfig")
	return
}

func NewModifyPrometheusConfigResponse() (response *ModifyPrometheusConfigResponse) {
	response = &ModifyPrometheusConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群采集配置
func (c *Client) ModifyPrometheusConfig(request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusConfigRequest()
	}
	response = NewModifyPrometheusConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetVbcRouteRequest() (request *GetVbcRouteRequest) {
	request = &GetVbcRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetVbcRoute")
	return
}

func NewGetVbcRouteResponse() (response *GetVbcRouteResponse) {
	response = &GetVbcRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询tke集群cidr是否加入云联网
func (c *Client) GetVbcRoute(request *GetVbcRouteRequest) (response *GetVbcRouteResponse, err error) {
	if request == nil {
		request = NewGetVbcRouteRequest()
	}
	response = NewGetVbcRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirtualServicesRequest() (request *DescribeVirtualServicesRequest) {
	request = &DescribeVirtualServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeVirtualServices")
	return
}

func NewDescribeVirtualServicesResponse() (response *DescribeVirtualServicesResponse) {
	response = &DescribeVirtualServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询虚拟服务.
func (c *Client) DescribeVirtualServices(request *DescribeVirtualServicesRequest) (response *DescribeVirtualServicesResponse, err error) {
	if request == nil {
		request = NewDescribeVirtualServicesRequest()
	}
	response = NewDescribeVirtualServicesResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeClusterInstancesRequest() (request *UpgradeClusterInstancesRequest) {
	request = &UpgradeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterInstances")
	return
}

func NewUpgradeClusterInstancesResponse() (response *UpgradeClusterInstancesResponse) {
	response = &UpgradeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给集群的一批work节点进行升级
func (c *Client) UpgradeClusterInstances(request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewUpgradeClusterInstancesRequest()
	}
	response = NewUpgradeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeAppChartListRequest() (request *DescribeTKEEdgeAppChartListRequest) {
	request = &DescribeTKEEdgeAppChartListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeAppChartList")
	return
}

func NewDescribeTKEEdgeAppChartListResponse() (response *DescribeTKEEdgeAppChartListResponse) {
	response = &DescribeTKEEdgeAppChartListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TKEEdge支持的App列表
func (c *Client) DescribeTKEEdgeAppChartList(request *DescribeTKEEdgeAppChartListRequest) (response *DescribeTKEEdgeAppChartListResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeAppChartListRequest()
	}
	response = NewDescribeTKEEdgeAppChartListResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterLevelPriceRequest() (request *GetClusterLevelPriceRequest) {
	request = &GetClusterLevelPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetClusterLevelPrice")
	return
}

func NewGetClusterLevelPriceResponse() (response *GetClusterLevelPriceResponse) {
	response = &GetClusterLevelPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群规模价格
func (c *Client) GetClusterLevelPrice(request *GetClusterLevelPriceRequest) (response *GetClusterLevelPriceResponse, err error) {
	if request == nil {
		request = NewGetClusterLevelPriceRequest()
	}
	response = NewGetClusterLevelPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRouteTablesRequest() (request *DescribeClusterRouteTablesRequest) {
	request = &DescribeClusterRouteTablesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRouteTables")
	return
}

func NewDescribeClusterRouteTablesResponse() (response *DescribeClusterRouteTablesResponse) {
	response = &DescribeClusterRouteTablesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群路由表
func (c *Client) DescribeClusterRouteTables(request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRouteTablesRequest()
	}
	response = NewDescribeClusterRouteTablesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEtcdRegionsRequest() (request *DescribeEtcdRegionsRequest) {
	request = &DescribeEtcdRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEtcdRegions")
	return
}

func NewDescribeEtcdRegionsResponse() (response *DescribeEtcdRegionsResponse) {
	response = &DescribeEtcdRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看etcd支持地域
func (c *Client) DescribeEtcdRegions(request *DescribeEtcdRegionsRequest) (response *DescribeEtcdRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeEtcdRegionsRequest()
	}
	response = NewDescribeEtcdRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeGroupRequest() (request *DescribeNodeGroupRequest) {
	request = &DescribeNodeGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeNodeGroup")
	return
}

func NewDescribeNodeGroupResponse() (response *DescribeNodeGroupResponse) {
	response = &DescribeNodeGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取NodeGroup信息
func (c *Client) DescribeNodeGroup(request *DescribeNodeGroupRequest) (response *DescribeNodeGroupResponse, err error) {
	if request == nil {
		request = NewDescribeNodeGroupRequest()
	}
	response = NewDescribeNodeGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionStatRequest() (request *DescribeRegionStatRequest) {
	request = &DescribeRegionStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeRegionStat")
	return
}

func NewDescribeRegionStatResponse() (response *DescribeRegionStatResponse) {
	response = &DescribeRegionStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前地域统计数据，包括集群数、节点数、工作负载数
func (c *Client) DescribeRegionStat(request *DescribeRegionStatRequest) (response *DescribeRegionStatResponse, err error) {
	if request == nil {
		request = NewDescribeRegionStatRequest()
	}
	response = NewDescribeRegionStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksMetaFeatureProgressRequest() (request *DescribeEksMetaFeatureProgressRequest) {
	request = &DescribeEksMetaFeatureProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksMetaFeatureProgress")
	return
}

func NewDescribeEksMetaFeatureProgressResponse() (response *DescribeEksMetaFeatureProgressResponse) {
	response = &DescribeEksMetaFeatureProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询EKS集群开通跨租户弹性网卡开通状态接口
func (c *Client) DescribeEksMetaFeatureProgress(request *DescribeEksMetaFeatureProgressRequest) (response *DescribeEksMetaFeatureProgressResponse, err error) {
	if request == nil {
		request = NewDescribeEksMetaFeatureProgressRequest()
	}
	response = NewDescribeEksMetaFeatureProgressResponse()
	err = c.Send(request, response)
	return
}

func NewEnableEdgeEventPersistenceRequest() (request *EnableEdgeEventPersistenceRequest) {
	request = &EnableEdgeEventPersistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableEdgeEventPersistence")
	return
}

func NewEnableEdgeEventPersistenceResponse() (response *EnableEdgeEventPersistenceResponse) {
	response = &EnableEdgeEventPersistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启事件持久化功能
func (c *Client) EnableEdgeEventPersistence(request *EnableEdgeEventPersistenceRequest) (response *EnableEdgeEventPersistenceResponse, err error) {
	if request == nil {
		request = NewEnableEdgeEventPersistenceRequest()
	}
	response = NewEnableEdgeEventPersistenceResponse()
	err = c.Send(request, response)
	return
}

func NewListECMEKSPodsRequest() (request *ListECMEKSPodsRequest) {
	request = &ListECMEKSPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListECMEKSPods")
	return
}

func NewListECMEKSPodsResponse() (response *ListECMEKSPodsResponse) {
	response = &ListECMEKSPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ECM环境pod详细信息
func (c *Client) ListECMEKSPods(request *ListECMEKSPodsRequest) (response *ListECMEKSPodsResponse, err error) {
	if request == nil {
		request = NewListECMEKSPodsRequest()
	}
	response = NewListECMEKSPodsResponse()
	err = c.Send(request, response)
	return
}

func NewAcquireEKSClusterAdminRoleRequest() (request *AcquireEKSClusterAdminRoleRequest) {
	request = &AcquireEKSClusterAdminRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AcquireEKSClusterAdminRole")
	return
}

func NewAcquireEKSClusterAdminRoleResponse() (response *AcquireEKSClusterAdminRoleResponse) {
	response = &AcquireEKSClusterAdminRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过此接口，可以获取EKS集群的tke:admin的ClusterRole，即管理员角色，可以用于CAM侧高权限的用户，通过CAM策略给予子账户此接口权限，进而可以通过此接口直接获取到kubernetes集群内的管理员角色。
func (c *Client) AcquireEKSClusterAdminRole(request *AcquireEKSClusterAdminRoleRequest) (response *AcquireEKSClusterAdminRoleResponse, err error) {
	if request == nil {
		request = NewAcquireEKSClusterAdminRoleRequest()
	}
	response = NewAcquireEKSClusterAdminRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterStaticInstallerRequest() (request *DescribeClusterStaticInstallerRequest) {
	request = &DescribeClusterStaticInstallerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterStaticInstaller")
	return
}

func NewDescribeClusterStaticInstallerResponse() (response *DescribeClusterStaticInstallerResponse) {
	response = &DescribeClusterStaticInstallerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群节点静态安装脚本
func (c *Client) DescribeClusterStaticInstaller(request *DescribeClusterStaticInstallerRequest) (response *DescribeClusterStaticInstallerResponse, err error) {
	if request == nil {
		request = NewDescribeClusterStaticInstallerRequest()
	}
	response = NewDescribeClusterStaticInstallerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeClustersRequest() (request *DescribeTKEEdgeClustersRequest) {
	request = &DescribeTKEEdgeClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusters")
	return
}

func NewDescribeTKEEdgeClustersResponse() (response *DescribeTKEEdgeClustersResponse) {
	response = &DescribeTKEEdgeClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询边缘集群列表
func (c *Client) DescribeTKEEdgeClusters(request *DescribeTKEEdgeClustersRequest) (response *DescribeTKEEdgeClustersResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeClustersRequest()
	}
	response = NewDescribeTKEEdgeClustersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudRunClusterQuotaRequest() (request *ModifyCloudRunClusterQuotaRequest) {
	request = &ModifyCloudRunClusterQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyCloudRunClusterQuota")
	return
}

func NewModifyCloudRunClusterQuotaResponse() (response *ModifyCloudRunClusterQuotaResponse) {
	response = &ModifyCloudRunClusterQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改CloudRun集群配额
func (c *Client) ModifyCloudRunClusterQuota(request *ModifyCloudRunClusterQuotaRequest) (response *ModifyCloudRunClusterQuotaResponse, err error) {
	if request == nil {
		request = NewModifyCloudRunClusterQuotaRequest()
	}
	response = NewModifyCloudRunClusterQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewForwardTKEEdgeClusterRequestRequest() (request *ForwardTKEEdgeClusterRequestRequest) {
	request = &ForwardTKEEdgeClusterRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ForwardTKEEdgeClusterRequest")
	return
}

func NewForwardTKEEdgeClusterRequestResponse() (response *ForwardTKEEdgeClusterRequestResponse) {
	response = &ForwardTKEEdgeClusterRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询、新增、删除、编辑TKE边缘计算集群内资源
func (c *Client) ForwardTKEEdgeClusterRequest(request *ForwardTKEEdgeClusterRequestRequest) (response *ForwardTKEEdgeClusterRequestResponse, err error) {
	if request == nil {
		request = NewForwardTKEEdgeClusterRequestRequest()
	}
	response = NewForwardTKEEdgeClusterRequestResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrometheusTemplateSyncRequest() (request *DeletePrometheusTemplateSyncRequest) {
	request = &DeletePrometheusTemplateSyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusTemplateSync")
	return
}

func NewDeletePrometheusTemplateSyncResponse() (response *DeletePrometheusTemplateSyncResponse) {
	response = &DeletePrometheusTemplateSyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消模板同步，这将会删除目标中该模板所生产的配置
func (c *Client) DeletePrometheusTemplateSync(request *DeletePrometheusTemplateSyncRequest) (response *DeletePrometheusTemplateSyncResponse, err error) {
	if request == nil {
		request = NewDeletePrometheusTemplateSyncRequest()
	}
	response = NewDeletePrometheusTemplateSyncResponse()
	err = c.Send(request, response)
	return
}

func NewClearExpiredClusterAuthRequest() (request *ClearExpiredClusterAuthRequest) {
	request = &ClearExpiredClusterAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ClearExpiredClusterAuth")
	return
}

func NewClearExpiredClusterAuthResponse() (response *ClearExpiredClusterAuthResponse) {
	response = &ClearExpiredClusterAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键清除集群内过期的权限信息
func (c *Client) ClearExpiredClusterAuth(request *ClearExpiredClusterAuthRequest) (response *ClearExpiredClusterAuthResponse, err error) {
	if request == nil {
		request = NewClearExpiredClusterAuthRequest()
	}
	response = NewClearExpiredClusterAuthResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecretRequest() (request *CreateSecretRequest) {
	request = &CreateSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateSecret")
	return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
	response = &CreateSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
	if request == nil {
		request = NewCreateSecretRequest()
	}
	response = NewCreateSecretResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUnitClusterRequest() (request *UpdateUnitClusterRequest) {
	request = &UpdateUnitClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateUnitCluster")
	return
}

func NewUpdateUnitClusterResponse() (response *UpdateUnitClusterResponse) {
	response = &UpdateUnitClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改边缘集群中部署的k3s集群的属性，例如开启公网访问
func (c *Client) UpdateUnitCluster(request *UpdateUnitClusterRequest) (response *UpdateUnitClusterResponse, err error) {
	if request == nil {
		request = NewUpdateUnitClusterRequest()
	}
	response = NewUpdateUnitClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodRequest() (request *GetPodRequest) {
	request = &GetPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetPod")
	return
}

func NewGetPodResponse() (response *GetPodResponse) {
	response = &GetPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询pod信息
func (c *Client) GetPod(request *GetPodRequest) (response *GetPodResponse, err error) {
	if request == nil {
		request = NewGetPodRequest()
	}
	response = NewGetPodResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterEndpointVipRequest() (request *DeleteClusterEndpointVipRequest) {
	request = &DeleteClusterEndpointVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpointVip")
	return
}

func NewDeleteClusterEndpointVipResponse() (response *DeleteClusterEndpointVipResponse) {
	response = &DeleteClusterEndpointVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除托管集群外网访问端口（老的方式，仅支持托管集群外网端口）
func (c *Client) DeleteClusterEndpointVip(request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
	if request == nil {
		request = NewDeleteClusterEndpointVipRequest()
	}
	response = NewDeleteClusterEndpointVipResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterVirtualNodePoolRequest() (request *ModifyClusterVirtualNodePoolRequest) {
	request = &ModifyClusterVirtualNodePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterVirtualNodePool")
	return
}

func NewModifyClusterVirtualNodePoolResponse() (response *ModifyClusterVirtualNodePoolResponse) {
	response = &ModifyClusterVirtualNodePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改超级节点池
func (c *Client) ModifyClusterVirtualNodePool(request *ModifyClusterVirtualNodePoolRequest) (response *ModifyClusterVirtualNodePoolResponse, err error) {
	if request == nil {
		request = NewModifyClusterVirtualNodePoolRequest()
	}
	response = NewModifyClusterVirtualNodePoolResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRouteRequest() (request *CreateClusterRouteRequest) {
	request = &CreateClusterRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRoute")
	return
}

func NewCreateClusterRouteResponse() (response *CreateClusterRouteResponse) {
	response = &CreateClusterRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群路由
func (c *Client) CreateClusterRoute(request *CreateClusterRouteRequest) (response *CreateClusterRouteResponse, err error) {
	if request == nil {
		request = NewCreateClusterRouteRequest()
	}
	response = NewCreateClusterRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEtcdSnapshotRequest() (request *DeleteEtcdSnapshotRequest) {
	request = &DeleteEtcdSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEtcdSnapshot")
	return
}

func NewDeleteEtcdSnapshotResponse() (response *DeleteEtcdSnapshotResponse) {
	response = &DeleteEtcdSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除etcd快照
func (c *Client) DeleteEtcdSnapshot(request *DeleteEtcdSnapshotRequest) (response *DeleteEtcdSnapshotResponse, err error) {
	if request == nil {
		request = NewDeleteEtcdSnapshotRequest()
	}
	response = NewDeleteEtcdSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAuthenticationOptionsRequest() (request *DescribeClusterAuthenticationOptionsRequest) {
	request = &DescribeClusterAuthenticationOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAuthenticationOptions")
	return
}

func NewDescribeClusterAuthenticationOptionsResponse() (response *DescribeClusterAuthenticationOptionsResponse) {
	response = &DescribeClusterAuthenticationOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看集群认证配置
func (c *Client) DescribeClusterAuthenticationOptions(request *DescribeClusterAuthenticationOptionsRequest) (response *DescribeClusterAuthenticationOptionsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAuthenticationOptionsRequest()
	}
	response = NewDescribeClusterAuthenticationOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTKEEdgeExternalKubeconfigRequest() (request *DescribeTKEEdgeExternalKubeconfigRequest) {
	request = &DescribeTKEEdgeExternalKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeExternalKubeconfig")
	return
}

func NewDescribeTKEEdgeExternalKubeconfigResponse() (response *DescribeTKEEdgeExternalKubeconfigResponse) {
	response = &DescribeTKEEdgeExternalKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘计算外部访问的kubeconfig
func (c *Client) DescribeTKEEdgeExternalKubeconfig(request *DescribeTKEEdgeExternalKubeconfigRequest) (response *DescribeTKEEdgeExternalKubeconfigResponse, err error) {
	if request == nil {
		request = NewDescribeTKEEdgeExternalKubeconfigRequest()
	}
	response = NewDescribeTKEEdgeExternalKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetVbcInstanceRequest() (request *GetVbcInstanceRequest) {
	request = &GetVbcInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetVbcInstance")
	return
}

func NewGetVbcInstanceResponse() (response *GetVbcInstanceResponse) {
	response = &GetVbcInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询vpc是否加入云联网
func (c *Client) GetVbcInstance(request *GetVbcInstanceRequest) (response *GetVbcInstanceResponse, err error) {
	if request == nil {
		request = NewGetVbcInstanceRequest()
	}
	response = NewGetVbcInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudRunPodRequest() (request *ModifyCloudRunPodRequest) {
	request = &ModifyCloudRunPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyCloudRunPod")
	return
}

func NewModifyCloudRunPodResponse() (response *ModifyCloudRunPodResponse) {
	response = &ModifyCloudRunPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改CloudRun Pod规格信息
func (c *Client) ModifyCloudRunPod(request *ModifyCloudRunPodRequest) (response *ModifyCloudRunPodResponse, err error) {
	if request == nil {
		request = NewModifyCloudRunPodRequest()
	}
	response = NewModifyCloudRunPodResponse()
	err = c.Send(request, response)
	return
}

func NewGetTkeAppUpgradeInfoRequest() (request *GetTkeAppUpgradeInfoRequest) {
	request = &GetTkeAppUpgradeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GetTkeAppUpgradeInfo")
	return
}

func NewGetTkeAppUpgradeInfoResponse() (response *GetTkeAppUpgradeInfoResponse) {
	response = &GetTkeAppUpgradeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群下组件升级信息
func (c *Client) GetTkeAppUpgradeInfo(request *GetTkeAppUpgradeInfoRequest) (response *GetTkeAppUpgradeInfoResponse, err error) {
	if request == nil {
		request = NewGetTkeAppUpgradeInfoRequest()
	}
	response = NewGetTkeAppUpgradeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMigrationRequest() (request *DescribeClusterMigrationRequest) {
	request = &DescribeClusterMigrationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterMigration")
	return
}

func NewDescribeClusterMigrationResponse() (response *DescribeClusterMigrationResponse) {
	response = &DescribeClusterMigrationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TKE集群迁移状态
func (c *Client) DescribeClusterMigration(request *DescribeClusterMigrationRequest) (response *DescribeClusterMigrationResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMigrationRequest()
	}
	response = NewDescribeClusterMigrationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusGlobalNotificationRequest() (request *DescribePrometheusGlobalNotificationRequest) {
	request = &DescribePrometheusGlobalNotificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusGlobalNotification")
	return
}

func NewDescribePrometheusGlobalNotificationResponse() (response *DescribePrometheusGlobalNotificationResponse) {
	response = &DescribePrometheusGlobalNotificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全局告警通知渠道
func (c *Client) DescribePrometheusGlobalNotification(request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusGlobalNotificationRequest()
	}
	response = NewDescribePrometheusGlobalNotificationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusInstanceDetailRequest() (request *DescribePrometheusInstanceDetailRequest) {
	request = &DescribePrometheusInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstanceDetail")
	return
}

func NewDescribePrometheusInstanceDetailResponse() (response *DescribePrometheusInstanceDetailResponse) {
	response = &DescribePrometheusInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取TMP实例详情
func (c *Client) DescribePrometheusInstanceDetail(request *DescribePrometheusInstanceDetailRequest) (response *DescribePrometheusInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusInstanceDetailRequest()
	}
	response = NewDescribePrometheusInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEksAlarmPoliciesRequest() (request *DeleteEksAlarmPoliciesRequest) {
	request = &DeleteEksAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteEksAlarmPolicies")
	return
}

func NewDeleteEksAlarmPoliciesResponse() (response *DeleteEksAlarmPoliciesResponse) {
	response = &DeleteEksAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性集群告警策略，支持批量删除
func (c *Client) DeleteEksAlarmPolicies(request *DeleteEksAlarmPoliciesRequest) (response *DeleteEksAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteEksAlarmPoliciesRequest()
	}
	response = NewDeleteEksAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterReleaseHistoryRequest() (request *DescribeClusterReleaseHistoryRequest) {
	request = &DescribeClusterReleaseHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterReleaseHistory")
	return
}

func NewDescribeClusterReleaseHistoryResponse() (response *DescribeClusterReleaseHistoryResponse) {
	response = &DescribeClusterReleaseHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群在应用市场中某个已安装应用的版本历史
func (c *Client) DescribeClusterReleaseHistory(request *DescribeClusterReleaseHistoryRequest) (response *DescribeClusterReleaseHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeClusterReleaseHistoryRequest()
	}
	response = NewDescribeClusterReleaseHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVirtualServicesRequest() (request *DeleteVirtualServicesRequest) {
	request = &DeleteVirtualServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteVirtualServices")
	return
}

func NewDeleteVirtualServicesResponse() (response *DeleteVirtualServicesResponse) {
	response = &DeleteVirtualServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除虚拟服务.
func (c *Client) DeleteVirtualServices(request *DeleteVirtualServicesRequest) (response *DeleteVirtualServicesResponse, err error) {
	if request == nil {
		request = NewDeleteVirtualServicesRequest()
	}
	response = NewDeleteVirtualServicesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCloudRunHPARequest() (request *ModifyCloudRunHPARequest) {
	request = &ModifyCloudRunHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyCloudRunHPA")
	return
}

func NewModifyCloudRunHPAResponse() (response *ModifyCloudRunHPAResponse) {
	response = &ModifyCloudRunHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// cloudrun hpa
func (c *Client) ModifyCloudRunHPA(request *ModifyCloudRunHPARequest) (response *ModifyCloudRunHPAResponse, err error) {
	if request == nil {
		request = NewModifyCloudRunHPARequest()
	}
	response = NewModifyCloudRunHPAResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterAsGroupRequest() (request *CreateClusterAsGroupRequest) {
	request = &CreateClusterAsGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterAsGroup")
	return
}

func NewCreateClusterAsGroupResponse() (response *CreateClusterAsGroupResponse) {
	response = &CreateClusterAsGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为已经存在的集群创建伸缩组
func (c *Client) CreateClusterAsGroup(request *CreateClusterAsGroupRequest) (response *CreateClusterAsGroupResponse, err error) {
	if request == nil {
		request = NewCreateClusterAsGroupRequest()
	}
	response = NewCreateClusterAsGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterTagsRequest() (request *ModifyClusterTagsRequest) {
	request = &ModifyClusterTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterTags")
	return
}

func NewModifyClusterTagsResponse() (response *ModifyClusterTagsResponse) {
	response = &ModifyClusterTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群标签
func (c *Client) ModifyClusterTags(request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
	if request == nil {
		request = NewModifyClusterTagsRequest()
	}
	response = NewModifyClusterTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeRegionRequest() (request *DescribeEdgeRegionRequest) {
	request = &DescribeEdgeRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeRegion")
	return
}

func NewDescribeEdgeRegionResponse() (response *DescribeEdgeRegionResponse) {
	response = &DescribeEdgeRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘集群的区域列表
func (c *Client) DescribeEdgeRegion(request *DescribeEdgeRegionRequest) (response *DescribeEdgeRegionResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeRegionRequest()
	}
	response = NewDescribeEdgeRegionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGrayScaleDeploymentGridRequest() (request *UpdateGrayScaleDeploymentGridRequest) {
	request = &UpdateGrayScaleDeploymentGridRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateGrayScaleDeploymentGrid")
	return
}

func NewUpdateGrayScaleDeploymentGridResponse() (response *UpdateGrayScaleDeploymentGridResponse) {
	response = &UpdateGrayScaleDeploymentGridResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 灰度DeploymentGrid
func (c *Client) UpdateGrayScaleDeploymentGrid(request *UpdateGrayScaleDeploymentGridRequest) (response *UpdateGrayScaleDeploymentGridResponse, err error) {
	if request == nil {
		request = NewUpdateGrayScaleDeploymentGridRequest()
	}
	response = NewUpdateGrayScaleDeploymentGridResponse()
	err = c.Send(request, response)
	return
}

func NewCancelClusterReleaseRequest() (request *CancelClusterReleaseRequest) {
	request = &CancelClusterReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CancelClusterRelease")
	return
}

func NewCancelClusterReleaseResponse() (response *CancelClusterReleaseResponse) {
	response = &CancelClusterReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在应用市场中取消安装失败的应用
func (c *Client) CancelClusterRelease(request *CancelClusterReleaseRequest) (response *CancelClusterReleaseResponse, err error) {
	if request == nil {
		request = NewCancelClusterReleaseRequest()
	}
	response = NewCancelClusterReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterExtraArgsRequest() (request *DescribeClusterExtraArgsRequest) {
	request = &DescribeClusterExtraArgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterExtraArgs")
	return
}

func NewDescribeClusterExtraArgsResponse() (response *DescribeClusterExtraArgsResponse) {
	response = &DescribeClusterExtraArgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群自定义参数
func (c *Client) DescribeClusterExtraArgs(request *DescribeClusterExtraArgsRequest) (response *DescribeClusterExtraArgsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterExtraArgsRequest()
	}
	response = NewDescribeClusterExtraArgsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPrometheusInstanceRequest() (request *ModifyPrometheusInstanceRequest) {
	request = &ModifyPrometheusInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusInstance")
	return
}

func NewModifyPrometheusInstanceResponse() (response *ModifyPrometheusInstanceResponse) {
	response = &ModifyPrometheusInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例属性
func (c *Client) ModifyPrometheusInstance(request *ModifyPrometheusInstanceRequest) (response *ModifyPrometheusInstanceResponse, err error) {
	if request == nil {
		request = NewModifyPrometheusInstanceRequest()
	}
	response = NewModifyPrometheusInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterVirtualNodeRequest() (request *DescribeClusterVirtualNodeRequest) {
	request = &DescribeClusterVirtualNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNode")
	return
}

func NewDescribeClusterVirtualNodeResponse() (response *DescribeClusterVirtualNodeResponse) {
	response = &DescribeClusterVirtualNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看超级节点列表
func (c *Client) DescribeClusterVirtualNode(request *DescribeClusterVirtualNodeRequest) (response *DescribeClusterVirtualNodeResponse, err error) {
	if request == nil {
		request = NewDescribeClusterVirtualNodeRequest()
	}
	response = NewDescribeClusterVirtualNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteECMInstancesRequest() (request *DeleteECMInstancesRequest) {
	request = &DeleteECMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DeleteECMInstances")
	return
}

func NewDeleteECMInstancesResponse() (response *DeleteECMInstancesResponse) {
	response = &DeleteECMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ECM实例
func (c *Client) DeleteECMInstances(request *DeleteECMInstancesRequest) (response *DeleteECMInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteECMInstancesRequest()
	}
	response = NewDeleteECMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEdgeCVMInstancesRequest() (request *DescribeEdgeCVMInstancesRequest) {
	request = &DescribeEdgeCVMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeCVMInstances")
	return
}

func NewDescribeEdgeCVMInstancesResponse() (response *DescribeEdgeCVMInstancesResponse) {
	response = &DescribeEdgeCVMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘容器CVM实例相关信息
func (c *Client) DescribeEdgeCVMInstances(request *DescribeEdgeCVMInstancesRequest) (response *DescribeEdgeCVMInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeEdgeCVMInstancesRequest()
	}
	response = NewDescribeEdgeCVMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEdgeCVMInstancesRequest() (request *CreateEdgeCVMInstancesRequest) {
	request = &CreateEdgeCVMInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeCVMInstances")
	return
}

func NewCreateEdgeCVMInstancesResponse() (response *CreateEdgeCVMInstancesResponse) {
	response = &CreateEdgeCVMInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建边缘容器CVM机器
func (c *Client) CreateEdgeCVMInstances(request *CreateEdgeCVMInstancesRequest) (response *CreateEdgeCVMInstancesResponse, err error) {
	if request == nil {
		request = NewCreateEdgeCVMInstancesRequest()
	}
	response = NewCreateEdgeCVMInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTemplateSyncRequest() (request *DescribePrometheusTemplateSyncRequest) {
	request = &DescribePrometheusTemplateSyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTemplateSync")
	return
}

func NewDescribePrometheusTemplateSyncResponse() (response *DescribePrometheusTemplateSyncResponse) {
	response = &DescribePrometheusTemplateSyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模板同步信息
func (c *Client) DescribePrometheusTemplateSync(request *DescribePrometheusTemplateSyncRequest) (response *DescribePrometheusTemplateSyncResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTemplateSyncRequest()
	}
	response = NewDescribePrometheusTemplateSyncResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTDCCClusterMetricDataRequest() (request *DescribeTDCCClusterMetricDataRequest) {
	request = &DescribeTDCCClusterMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeTDCCClusterMetricData")
	return
}

func NewDescribeTDCCClusterMetricDataResponse() (response *DescribeTDCCClusterMetricDataResponse) {
	response = &DescribeTDCCClusterMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出注册集群Metric信息
func (c *Client) DescribeTDCCClusterMetricData(request *DescribeTDCCClusterMetricDataRequest) (response *DescribeTDCCClusterMetricDataResponse, err error) {
	if request == nil {
		request = NewDescribeTDCCClusterMetricDataRequest()
	}
	response = NewDescribeTDCCClusterMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
	request = &DescribeProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeProducts")
	return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
	response = &DescribeProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用市场中制品列表
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
	if request == nil {
		request = NewDescribeProductsRequest()
	}
	response = NewDescribeProductsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableMetaFeatureRequest() (request *EnableMetaFeatureRequest) {
	request = &EnableMetaFeatureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "EnableMetaFeature")
	return
}

func NewEnableMetaFeatureResponse() (response *EnableMetaFeatureResponse) {
	response = &EnableMetaFeatureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 跨租户相关接口，需要控制不能对外
func (c *Client) EnableMetaFeature(request *EnableMetaFeatureRequest) (response *EnableMetaFeatureResponse, err error) {
	if request == nil {
		request = NewEnableMetaFeatureRequest()
	}
	response = NewEnableMetaFeatureResponse()
	err = c.Send(request, response)
	return
}

func NewServiceMeshForwardRequestRequest() (request *ServiceMeshForwardRequestRequest) {
	request = &ServiceMeshForwardRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ServiceMeshForwardRequest")
	return
}

func NewServiceMeshForwardRequestResponse() (response *ServiceMeshForwardRequestResponse) {
	response = &ServiceMeshForwardRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务网格代理转发
func (c *Client) ServiceMeshForwardRequest(request *ServiceMeshForwardRequestRequest) (response *ServiceMeshForwardRequestResponse, err error) {
	if request == nil {
		request = NewServiceMeshForwardRequestRequest()
	}
	response = NewServiceMeshForwardRequestResponse()
	err = c.Send(request, response)
	return
}

func NewCheckIsPrometheusNewUserRequest() (request *CheckIsPrometheusNewUserRequest) {
	request = &CheckIsPrometheusNewUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CheckIsPrometheusNewUser")
	return
}

func NewCheckIsPrometheusNewUserResponse() (response *CheckIsPrometheusNewUserResponse) {
	response = &CheckIsPrometheusNewUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 判断用户是否为云原生监控新用户，即在任何地域下均未创建过监控实例的用户
func (c *Client) CheckIsPrometheusNewUser(request *CheckIsPrometheusNewUserRequest) (response *CheckIsPrometheusNewUserResponse, err error) {
	if request == nil {
		request = NewCheckIsPrometheusNewUserRequest()
	}
	response = NewCheckIsPrometheusNewUserResponse()
	err = c.Send(request, response)
	return
}

func NewGrantCodingClusterRoleBindingRequest() (request *GrantCodingClusterRoleBindingRequest) {
	request = &GrantCodingClusterRoleBindingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "GrantCodingClusterRoleBinding")
	return
}

func NewGrantCodingClusterRoleBindingResponse() (response *GrantCodingClusterRoleBindingResponse) {
	response = &GrantCodingClusterRoleBindingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 授予Coding的服务角色RBAC tke:admin权限
func (c *Client) GrantCodingClusterRoleBinding(request *GrantCodingClusterRoleBindingRequest) (response *GrantCodingClusterRoleBindingResponse, err error) {
	if request == nil {
		request = NewGrantCodingClusterRoleBindingRequest()
	}
	response = NewGrantCodingClusterRoleBindingResponse()
	err = c.Send(request, response)
	return
}

func NewListTDCCRegionRequest() (request *ListTDCCRegionRequest) {
	request = &ListTDCCRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListTDCCRegion")
	return
}

func NewListTDCCRegionResponse() (response *ListTDCCRegionResponse) {
	response = &ListTDCCRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TDCC可用地域
func (c *Client) ListTDCCRegion(request *ListTDCCRegionRequest) (response *ListTDCCRegionResponse, err error) {
	if request == nil {
		request = NewListTDCCRegionRequest()
	}
	response = NewListTDCCRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEksContainerInstanceLogRequest() (request *DescribeEksContainerInstanceLogRequest) {
	request = &DescribeEksContainerInstanceLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeEksContainerInstanceLog")
	return
}

func NewDescribeEksContainerInstanceLogResponse() (response *DescribeEksContainerInstanceLogResponse) {
	response = &DescribeEksContainerInstanceLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器实例中容器日志
func (c *Client) DescribeEksContainerInstanceLog(request *DescribeEksContainerInstanceLogRequest) (response *DescribeEksContainerInstanceLogResponse, err error) {
	if request == nil {
		request = NewDescribeEksContainerInstanceLogRequest()
	}
	response = NewDescribeEksContainerInstanceLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpenPolicyListRequest() (request *DescribeOpenPolicyListRequest) {
	request = &DescribeOpenPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeOpenPolicyList")
	return
}

func NewDescribeOpenPolicyListResponse() (response *DescribeOpenPolicyListResponse) {
	response = &DescribeOpenPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询opa策略列表
func (c *Client) DescribeOpenPolicyList(request *DescribeOpenPolicyListRequest) (response *DescribeOpenPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeOpenPolicyListRequest()
	}
	response = NewDescribeOpenPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusTempSyncRequest() (request *DescribePrometheusTempSyncRequest) {
	request = &DescribePrometheusTempSyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusTempSync")
	return
}

func NewDescribePrometheusTempSyncResponse() (response *DescribePrometheusTempSyncResponse) {
	response = &DescribePrometheusTempSyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模板关联实例信息，针对V2版本实例
func (c *Client) DescribePrometheusTempSync(request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusTempSyncRequest()
	}
	response = NewDescribePrometheusTempSyncResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeECMEKSClusterCredentialRequest() (request *DescribeECMEKSClusterCredentialRequest) {
	request = &DescribeECMEKSClusterCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeECMEKSClusterCredential")
	return
}

func NewDescribeECMEKSClusterCredentialResponse() (response *DescribeECMEKSClusterCredentialResponse) {
	response = &DescribeECMEKSClusterCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ECM弹性容器集群的接入认证信息
func (c *Client) DescribeECMEKSClusterCredential(request *DescribeECMEKSClusterCredentialRequest) (response *DescribeECMEKSClusterCredentialResponse, err error) {
	if request == nil {
		request = NewDescribeECMEKSClusterCredentialRequest()
	}
	response = NewDescribeECMEKSClusterCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitClusterKubeconfigRequest() (request *DescribeUnitClusterKubeconfigRequest) {
	request = &DescribeUnitClusterKubeconfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeUnitClusterKubeconfig")
	return
}

func NewDescribeUnitClusterKubeconfigResponse() (response *DescribeUnitClusterKubeconfigResponse) {
	response = &DescribeUnitClusterKubeconfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取边缘集群的边缘节点上部署的k3s集群的kubeconfig
func (c *Client) DescribeUnitClusterKubeconfig(request *DescribeUnitClusterKubeconfigRequest) (response *DescribeUnitClusterKubeconfigResponse, err error) {
	if request == nil {
		request = NewDescribeUnitClusterKubeconfigRequest()
	}
	response = NewDescribeUnitClusterKubeconfigResponse()
	err = c.Send(request, response)
	return
}

func NewDisableVpcCniNetworkTypeRequest() (request *DisableVpcCniNetworkTypeRequest) {
	request = &DisableVpcCniNetworkTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DisableVpcCniNetworkType")
	return
}

func NewDisableVpcCniNetworkTypeResponse() (response *DisableVpcCniNetworkTypeResponse) {
	response = &DisableVpcCniNetworkTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供给附加了VPC-CNI能力的Global-Route集群关闭VPC-CNI
func (c *Client) DisableVpcCniNetworkType(request *DisableVpcCniNetworkTypeRequest) (response *DisableVpcCniNetworkTypeResponse, err error) {
	if request == nil {
		request = NewDisableVpcCniNetworkTypeRequest()
	}
	response = NewDisableVpcCniNetworkTypeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterInstancesRequest() (request *CreateClusterInstancesRequest) {
	request = &CreateClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreateClusterInstances")
	return
}

func NewCreateClusterInstancesResponse() (response *CreateClusterInstancesResponse) {
	response = &CreateClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩展(新建)集群节点
func (c *Client) CreateClusterInstances(request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
	if request == nil {
		request = NewCreateClusterInstancesRequest()
	}
	response = NewCreateClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePrometheusAlertRuleRequest() (request *CreatePrometheusAlertRuleRequest) {
	request = &CreatePrometheusAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAlertRule")
	return
}

func NewCreatePrometheusAlertRuleResponse() (response *CreatePrometheusAlertRuleResponse) {
	response = &CreatePrometheusAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警规则
func (c *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
	if request == nil {
		request = NewCreatePrometheusAlertRuleRequest()
	}
	response = NewCreatePrometheusAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewRestoreEtcdInstanceRequest() (request *RestoreEtcdInstanceRequest) {
	request = &RestoreEtcdInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "RestoreEtcdInstance")
	return
}

func NewRestoreEtcdInstanceResponse() (response *RestoreEtcdInstanceResponse) {
	response = &RestoreEtcdInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复etcd实例
func (c *Client) RestoreEtcdInstance(request *RestoreEtcdInstanceRequest) (response *RestoreEtcdInstanceResponse, err error) {
	if request == nil {
		request = NewRestoreEtcdInstanceRequest()
	}
	response = NewRestoreEtcdInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrometheusDashboardRequest() (request *DescribePrometheusDashboardRequest) {
	request = &DescribePrometheusDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusDashboard")
	return
}

func NewDescribePrometheusDashboardResponse() (response *DescribePrometheusDashboardResponse) {
	response = &DescribePrometheusDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询grafana监控面板
func (c *Client) DescribePrometheusDashboard(request *DescribePrometheusDashboardRequest) (response *DescribePrometheusDashboardResponse, err error) {
	if request == nil {
		request = NewDescribePrometheusDashboardRequest()
	}
	response = NewDescribePrometheusDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirtualClustersRequest() (request *DescribeVirtualClustersRequest) {
	request = &DescribeVirtualClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeVirtualClusters")
	return
}

func NewDescribeVirtualClustersResponse() (response *DescribeVirtualClustersResponse) {
	response = &DescribeVirtualClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询虚拟集群.
func (c *Client) DescribeVirtualClusters(request *DescribeVirtualClustersRequest) (response *DescribeVirtualClustersResponse, err error) {
	if request == nil {
		request = NewDescribeVirtualClustersRequest()
	}
	response = NewDescribeVirtualClustersResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
	request = &AddClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "AddClusterInstances")
	return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
	response = &AddClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩展集群节点，API 3.0
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
	if request == nil {
		request = NewAddClusterInstancesRequest()
	}
	response = NewAddClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sWorkloadRequest() (request *DescribeK8sWorkloadRequest) {
	request = &DescribeK8sWorkloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeK8sWorkload")
	return
}

func NewDescribeK8sWorkloadResponse() (response *DescribeK8sWorkloadResponse) {
	response = &DescribeK8sWorkloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作负载的配置
func (c *Client) DescribeK8sWorkload(request *DescribeK8sWorkloadRequest) (response *DescribeK8sWorkloadResponse, err error) {
	if request == nil {
		request = NewDescribeK8sWorkloadRequest()
	}
	response = NewDescribeK8sWorkloadResponse()
	err = c.Send(request, response)
	return
}

func NewListECMEKSK8SVersionRequest() (request *ListECMEKSK8SVersionRequest) {
	request = &ListECMEKSK8SVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "ListECMEKSK8SVersion")
	return
}

func NewListECMEKSK8SVersionResponse() (response *ListECMEKSK8SVersionResponse) {
	response = &ListECMEKSK8SVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ECM EKS支持的k8s版本
func (c *Client) ListECMEKSK8SVersion(request *ListECMEKSK8SVersionRequest) (response *ListECMEKSK8SVersionResponse, err error) {
	if request == nil {
		request = NewListECMEKSK8SVersionRequest()
	}
	response = NewListECMEKSK8SVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersCostStatesRequest() (request *DescribeClustersCostStatesRequest) {
	request = &DescribeClustersCostStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "DescribeClustersCostStates")
	return
}

func NewDescribeClustersCostStatesResponse() (response *DescribeClustersCostStatesResponse) {
	response = &DescribeClustersCostStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群成本概览列表数据
func (c *Client) DescribeClustersCostStates(request *DescribeClustersCostStatesRequest) (response *DescribeClustersCostStatesResponse, err error) {
	if request == nil {
		request = NewDescribeClustersCostStatesRequest()
	}
	response = NewDescribeClustersCostStatesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEKSContainerInstanceRequest() (request *UpdateEKSContainerInstanceRequest) {
	request = &UpdateEKSContainerInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tke", APIVersion, "UpdateEKSContainerInstance")
	return
}

func NewUpdateEKSContainerInstanceResponse() (response *UpdateEKSContainerInstanceResponse) {
	response = &UpdateEKSContainerInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新容器实例
func (c *Client) UpdateEKSContainerInstance(request *UpdateEKSContainerInstanceRequest) (response *UpdateEKSContainerInstanceResponse, err error) {
	if request == nil {
		request = NewUpdateEKSContainerInstanceRequest()
	}
	response = NewUpdateEKSContainerInstanceResponse()
	err = c.Send(request, response)
	return
}

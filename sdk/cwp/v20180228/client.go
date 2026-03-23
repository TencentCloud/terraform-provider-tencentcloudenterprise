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

package v20180228

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-02-28"

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

func NewDescribeVulDefencePluginDetailRequest() (request *DescribeVulDefencePluginDetailRequest) {
	request = &DescribeVulDefencePluginDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginDetail")
	return
}

func NewDescribeVulDefencePluginDetailResponse() (response *DescribeVulDefencePluginDetailResponse) {
	response = &DescribeVulDefencePluginDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单台主机漏洞防御插件信息
func (c *Client) DescribeVulDefencePluginDetail(request *DescribeVulDefencePluginDetailRequest) (response *DescribeVulDefencePluginDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefencePluginDetailRequest()
	}
	response = NewDescribeVulDefencePluginDetailResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDetectionExcelRequest() (request *ExportVulDetectionExcelRequest) {
	request = &ExportVulDetectionExcelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDetectionExcel")
	return
}

func NewExportVulDetectionExcelResponse() (response *ExportVulDetectionExcelResponse) {
	response = &ExportVulDetectionExcelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出本次漏洞检测Excel
func (c *Client) ExportVulDetectionExcel(request *ExportVulDetectionExcelRequest) (response *ExportVulDetectionExcelResponse, err error) {
	if request == nil {
		request = NewExportVulDetectionExcelRequest()
	}
	response = NewExportVulDetectionExcelResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgentInstallTaskRequest() (request *CreateAgentInstallTaskRequest) {
	request = &CreateAgentInstallTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateAgentInstallTask")
	return
}

func NewCreateAgentInstallTaskResponse() (response *CreateAgentInstallTaskResponse) {
	response = &CreateAgentInstallTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建安装主机安全agent任务
func (c *Client) CreateAgentInstallTask(request *CreateAgentInstallTaskRequest) (response *CreateAgentInstallTaskResponse, err error) {
	if request == nil {
		request = NewCreateAgentInstallTaskRequest()
	}
	response = NewCreateAgentInstallTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulFixRequest() (request *CreateVulFixRequest) {
	request = &CreateVulFixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateVulFix")
	return
}

func NewCreateVulFixResponse() (response *CreateVulFixResponse) {
	response = &CreateVulFixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交漏洞修护
func (c *Client) CreateVulFix(request *CreateVulFixRequest) (response *CreateVulFixResponse, err error) {
	if request == nil {
		request = NewCreateVulFixRequest()
	}
	response = NewCreateVulFixResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineClearHistoryRequest() (request *DescribeMachineClearHistoryRequest) {
	request = &DescribeMachineClearHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineClearHistory")
	return
}

func NewDescribeMachineClearHistoryResponse() (response *DescribeMachineClearHistoryResponse) {
	response = &DescribeMachineClearHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机器清理历史记录
func (c *Client) DescribeMachineClearHistory(request *DescribeMachineClearHistoryRequest) (response *DescribeMachineClearHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeMachineClearHistoryRequest()
	}
	response = NewDescribeMachineClearHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenGeneralStatRequest() (request *DescribeScreenGeneralStatRequest) {
	request = &DescribeScreenGeneralStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenGeneralStat")
	return
}

func NewDescribeScreenGeneralStatResponse() (response *DescribeScreenGeneralStatResponse) {
	response = &DescribeScreenGeneralStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化获取主机相关统计
func (c *Client) DescribeScreenGeneralStat(request *DescribeScreenGeneralStatRequest) (response *DescribeScreenGeneralStatResponse, err error) {
	if request == nil {
		request = NewDescribeScreenGeneralStatRequest()
	}
	response = NewDescribeScreenGeneralStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperRuleCountRequest() (request *DescribeFileTamperRuleCountRequest) {
	request = &DescribeFileTamperRuleCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRuleCount")
	return
}

func NewDescribeFileTamperRuleCountResponse() (response *DescribeFileTamperRuleCountResponse) {
	response = &DescribeFileTamperRuleCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机关联文件监控规则数量
func (c *Client) DescribeFileTamperRuleCount(request *DescribeFileTamperRuleCountRequest) (response *DescribeFileTamperRuleCountResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperRuleCountRequest()
	}
	response = NewDescribeFileTamperRuleCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImportMachineInfoRequest() (request *DescribeImportMachineInfoRequest) {
	request = &DescribeImportMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeImportMachineInfo")
	return
}

func NewDescribeImportMachineInfoResponse() (response *DescribeImportMachineInfoResponse) {
	response = &DescribeImportMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询批量导入机器信息
func (c *Client) DescribeImportMachineInfo(request *DescribeImportMachineInfoRequest) (response *DescribeImportMachineInfoResponse, err error) {
	if request == nil {
		request = NewDescribeImportMachineInfoRequest()
	}
	response = NewDescribeImportMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBaselineStrategyRequest() (request *CreateBaselineStrategyRequest) {
	request = &CreateBaselineStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateBaselineStrategy")
	return
}

func NewCreateBaselineStrategyResponse() (response *CreateBaselineStrategyResponse) {
	response = &CreateBaselineStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据策略信息创建基线策略
func (c *Client) CreateBaselineStrategy(request *CreateBaselineStrategyRequest) (response *CreateBaselineStrategyResponse, err error) {
	if request == nil {
		request = NewCreateBaselineStrategyRequest()
	}
	response = NewCreateBaselineStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRansomDefenseEventsStatusRequest() (request *ModifyRansomDefenseEventsStatusRequest) {
	request = &ModifyRansomDefenseEventsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyRansomDefenseEventsStatus")
	return
}

func NewModifyRansomDefenseEventsStatusResponse() (response *ModifyRansomDefenseEventsStatusResponse) {
	response = &ModifyRansomDefenseEventsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防勒索事件状态
func (c *Client) ModifyRansomDefenseEventsStatus(request *ModifyRansomDefenseEventsStatusRequest) (response *ModifyRansomDefenseEventsStatusResponse, err error) {
	if request == nil {
		request = NewModifyRansomDefenseEventsStatusRequest()
	}
	response = NewModifyRansomDefenseEventsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUsersConfigRequest() (request *ModifyUsersConfigRequest) {
	request = &ModifyUsersConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyUsersConfig")
	return
}

func NewModifyUsersConfigResponse() (response *ModifyUsersConfigResponse) {
	response = &ModifyUsersConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户配置
func (c *Client) ModifyUsersConfig(request *ModifyUsersConfigRequest) (response *ModifyUsersConfigResponse, err error) {
	if request == nil {
		request = NewModifyUsersConfigRequest()
	}
	response = NewModifyUsersConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineListRequest() (request *ExportBaselineListRequest) {
	request = &ExportBaselineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineList")
	return
}

func NewExportBaselineListResponse() (response *ExportBaselineListResponse) {
	response = &ExportBaselineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线列表
func (c *Client) ExportBaselineList(request *ExportBaselineListRequest) (response *ExportBaselineListResponse, err error) {
	if request == nil {
		request = NewExportBaselineListRequest()
	}
	response = NewExportBaselineListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemListRequest() (request *DescribeBaselineItemListRequest) {
	request = &DescribeBaselineItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemList")
	return
}

func NewDescribeBaselineItemListResponse() (response *DescribeBaselineItemListResponse) {
	response = &DescribeBaselineItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线项检测结果列表
func (c *Client) DescribeBaselineItemList(request *DescribeBaselineItemListRequest) (response *DescribeBaselineItemListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemListRequest()
	}
	response = NewDescribeBaselineItemListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDiskListRequest() (request *DescribeAssetDiskListRequest) {
	request = &DescribeAssetDiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDiskList")
	return
}

func NewDescribeAssetDiskListResponse() (response *DescribeAssetDiskListResponse) {
	response = &DescribeAssetDiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机磁盘分区列表
func (c *Client) DescribeAssetDiskList(request *DescribeAssetDiskListRequest) (response *DescribeAssetDiskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDiskListRequest()
	}
	response = NewDescribeAssetDiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackVulTypeListRequest() (request *DescribeAttackVulTypeListRequest) {
	request = &DescribeAttackVulTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackVulTypeList")
	return
}

func NewDescribeAttackVulTypeListResponse() (response *DescribeAttackVulTypeListResponse) {
	response = &DescribeAttackVulTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络攻击威胁类型列表
func (c *Client) DescribeAttackVulTypeList(request *DescribeAttackVulTypeListRequest) (response *DescribeAttackVulTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeAttackVulTypeListRequest()
	}
	response = NewDescribeAttackVulTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineClearHistoryRequest() (request *DeleteMachineClearHistoryRequest) {
	request = &DeleteMachineClearHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachineClearHistory")
	return
}

func NewDeleteMachineClearHistoryResponse() (response *DeleteMachineClearHistoryResponse) {
	response = &DeleteMachineClearHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机器清理记录
func (c *Client) DeleteMachineClearHistory(request *DeleteMachineClearHistoryRequest) (response *DeleteMachineClearHistoryResponse, err error) {
	if request == nil {
		request = NewDeleteMachineClearHistoryRequest()
	}
	response = NewDeleteMachineClearHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaliciousRequestWhiteListRequest() (request *ModifyMaliciousRequestWhiteListRequest) {
	request = &ModifyMaliciousRequestWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaliciousRequestWhiteList")
	return
}

func NewModifyMaliciousRequestWhiteListResponse() (response *ModifyMaliciousRequestWhiteListResponse) {
	response = &ModifyMaliciousRequestWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新恶意请求白名单
func (c *Client) ModifyMaliciousRequestWhiteList(request *ModifyMaliciousRequestWhiteListRequest) (response *ModifyMaliciousRequestWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyMaliciousRequestWhiteListRequest()
	}
	response = NewModifyMaliciousRequestWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLicenseRecordAllRequest() (request *DeleteLicenseRecordAllRequest) {
	request = &DeleteLicenseRecordAllRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteLicenseRecordAll")
	return
}

func NewDeleteLicenseRecordAllResponse() (response *DeleteLicenseRecordAllResponse) {
	response = &DeleteLicenseRecordAllResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除授权全部记录
func (c *Client) DeleteLicenseRecordAll(request *DeleteLicenseRecordAllRequest) (response *DeleteLicenseRecordAllResponse, err error) {
	if request == nil {
		request = NewDeleteLicenseRecordAllRequest()
	}
	response = NewDeleteLicenseRecordAllResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineTagTopRequest() (request *DescribeAssetMachineTagTopRequest) {
	request = &DescribeAssetMachineTagTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineTagTop")
	return
}

func NewDescribeAssetMachineTagTopResponse() (response *DescribeAssetMachineTagTopResponse) {
	response = &DescribeAssetMachineTagTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机标签Top5
func (c *Client) DescribeAssetMachineTagTop(request *DescribeAssetMachineTagTopRequest) (response *DescribeAssetMachineTagTopResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineTagTopRequest()
	}
	response = NewDescribeAssetMachineTagTopResponse()
	err = c.Send(request, response)
	return
}

func NewExportIgnoreRuleEffectHostListRequest() (request *ExportIgnoreRuleEffectHostListRequest) {
	request = &ExportIgnoreRuleEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportIgnoreRuleEffectHostList")
	return
}

func NewExportIgnoreRuleEffectHostListResponse() (response *ExportIgnoreRuleEffectHostListResponse) {
	response = &ExportIgnoreRuleEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据检测项id导出忽略检测项影响主机列表
func (c *Client) ExportIgnoreRuleEffectHostList(request *ExportIgnoreRuleEffectHostListRequest) (response *ExportIgnoreRuleEffectHostListResponse, err error) {
	if request == nil {
		request = NewExportIgnoreRuleEffectHostListRequest()
	}
	response = NewExportIgnoreRuleEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewExportBruteAttacksRequest() (request *ExportBruteAttacksRequest) {
	request = &ExportBruteAttacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBruteAttacks")
	return
}

func NewExportBruteAttacksResponse() (response *ExportBruteAttacksResponse) {
	response = &ExportBruteAttacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ExportBruteAttacks) 用于导出密码破解记录成CSV文件。
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
	if request == nil {
		request = NewExportBruteAttacksRequest()
	}
	response = NewExportBruteAttacksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJavaMemShellsStatusRequest() (request *ModifyJavaMemShellsStatusRequest) {
	request = &ModifyJavaMemShellsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyJavaMemShellsStatus")
	return
}

func NewModifyJavaMemShellsStatusResponse() (response *ModifyJavaMemShellsStatusResponse) {
	response = &ModifyJavaMemShellsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改java内存马事件状态
func (c *Client) ModifyJavaMemShellsStatus(request *ModifyJavaMemShellsStatusRequest) (response *ModifyJavaMemShellsStatusResponse, err error) {
	if request == nil {
		request = NewModifyJavaMemShellsStatusRequest()
	}
	response = NewModifyJavaMemShellsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSearchTemplateRequest() (request *DeleteSearchTemplateRequest) {
	request = &DeleteSearchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteSearchTemplate")
	return
}

func NewDeleteSearchTemplateResponse() (response *DeleteSearchTemplateResponse) {
	response = &DeleteSearchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除检索模板
func (c *Client) DeleteSearchTemplate(request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSearchTemplateRequest()
	}
	response = NewDeleteSearchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventInfoRequest() (request *DescribeReverseShellEventInfoRequest) {
	request = &DescribeReverseShellEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellEventInfo")
	return
}

func NewDescribeReverseShellEventInfoResponse() (response *DescribeReverseShellEventInfoResponse) {
	response = &DescribeReverseShellEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 反弹shell信息详情
func (c *Client) DescribeReverseShellEventInfo(request *DescribeReverseShellEventInfoRequest) (response *DescribeReverseShellEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventInfoRequest()
	}
	response = NewDescribeReverseShellEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineLicenseDetailRequest() (request *DescribeMachineLicenseDetailRequest) {
	request = &DescribeMachineLicenseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineLicenseDetail")
	return
}

func NewDescribeMachineLicenseDetailResponse() (response *DescribeMachineLicenseDetailResponse) {
	response = &DescribeMachineLicenseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeMachineLicenseDetail)查询机器授权信息
func (c *Client) DescribeMachineLicenseDetail(request *DescribeMachineLicenseDetailRequest) (response *DescribeMachineLicenseDetailResponse, err error) {
	if request == nil {
		request = NewDescribeMachineLicenseDetailRequest()
	}
	response = NewDescribeMachineLicenseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineItemDetectListRequest() (request *ExportBaselineItemDetectListRequest) {
	request = &ExportBaselineItemDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineItemDetectList")
	return
}

func NewExportBaselineItemDetectListResponse() (response *ExportBaselineItemDetectListResponse) {
	response = &ExportBaselineItemDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线检测项
func (c *Client) ExportBaselineItemDetectList(request *ExportBaselineItemDetectListRequest) (response *ExportBaselineItemDetectListResponse, err error) {
	if request == nil {
		request = NewExportBaselineItemDetectListRequest()
	}
	response = NewExportBaselineItemDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseStrategyListRequest() (request *DescribeRansomDefenseStrategyListRequest) {
	request = &DescribeRansomDefenseStrategyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyList")
	return
}

func NewDescribeRansomDefenseStrategyListResponse() (response *DescribeRansomDefenseStrategyListResponse) {
	response = &DescribeRansomDefenseStrategyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防勒索策略列表
func (c *Client) DescribeRansomDefenseStrategyList(request *DescribeRansomDefenseStrategyListRequest) (response *DescribeRansomDefenseStrategyListResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseStrategyListRequest()
	}
	response = NewDescribeRansomDefenseStrategyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenceEventDetailRequest() (request *DescribeDefenceEventDetailRequest) {
	request = &DescribeDefenceEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeDefenceEventDetail")
	return
}

func NewDescribeDefenceEventDetailResponse() (response *DescribeDefenceEventDetailResponse) {
	response = &DescribeDefenceEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御事件详情
func (c *Client) DescribeDefenceEventDetail(request *DescribeDefenceEventDetailRequest) (response *DescribeDefenceEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeDefenceEventDetailRequest()
	}
	response = NewDescribeDefenceEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGeneralRequest() (request *DescribeMachineGeneralRequest) {
	request = &DescribeMachineGeneralRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineGeneral")
	return
}

func NewDescribeMachineGeneralResponse() (response *DescribeMachineGeneralResponse) {
	response = &DescribeMachineGeneralResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机概览信息
func (c *Client) DescribeMachineGeneral(request *DescribeMachineGeneralRequest) (response *DescribeMachineGeneralResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGeneralRequest()
	}
	response = NewDescribeMachineGeneralResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBashPolicyRequest() (request *ModifyBashPolicyRequest) {
	request = &ModifyBashPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBashPolicy")
	return
}

func NewModifyBashPolicyResponse() (response *ModifyBashPolicyResponse) {
	response = &ModifyBashPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改高危命令策略
func (c *Client) ModifyBashPolicy(request *ModifyBashPolicyRequest) (response *ModifyBashPolicyResponse, err error) {
	if request == nil {
		request = NewModifyBashPolicyRequest()
	}
	response = NewModifyBashPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulLevelCountRequest() (request *DescribeVulLevelCountRequest) {
	request = &DescribeVulLevelCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulLevelCount")
	return
}

func NewDescribeVulLevelCountResponse() (response *DescribeVulLevelCountResponse) {
	response = &DescribeVulLevelCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞数量等级分布统计
func (c *Client) DescribeVulLevelCount(request *DescribeVulLevelCountRequest) (response *DescribeVulLevelCountResponse, err error) {
	if request == nil {
		request = NewDescribeVulLevelCountRequest()
	}
	response = NewDescribeVulLevelCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackEventInfoRequest() (request *DescribeAttackEventInfoRequest) {
	request = &DescribeAttackEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackEventInfo")
	return
}

func NewDescribeAttackEventInfoResponse() (response *DescribeAttackEventInfoResponse) {
	response = &DescribeAttackEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeAttackEventInfo(request *DescribeAttackEventInfoRequest) (response *DescribeAttackEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAttackEventInfoRequest()
	}
	response = NewDescribeAttackEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellRulesRequest() (request *DescribeReverseShellRulesRequest) {
	request = &DescribeReverseShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellRules")
	return
}

func NewDescribeReverseShellRulesResponse() (response *DescribeReverseShellRulesResponse) {
	response = &DescribeReverseShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取反弹Shell规则列表
func (c *Client) DescribeReverseShellRules(request *DescribeReverseShellRulesRequest) (response *DescribeReverseShellRulesResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellRulesRequest()
	}
	response = NewDescribeReverseShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceSettingRequest() (request *ModifyVulDefenceSettingRequest) {
	request = &ModifyVulDefenceSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyVulDefenceSetting")
	return
}

func NewModifyVulDefenceSettingResponse() (response *ModifyVulDefenceSettingResponse) {
	response = &ModifyVulDefenceSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御插件设置
// 1）新增主机自动加入，scope为1，quuids为空
// 2）全量旗舰版不自动加入，scope为0，quuids为当前quuid列表，
// 3）给定quuid列表，scope为0，quuids为用户选择quuid
func (c *Client) ModifyVulDefenceSetting(request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceSettingRequest()
	}
	response = NewModifyVulDefenceSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaCustomPolicyListRequest() (request *DescribeMaCustomPolicyListRequest) {
	request = &DescribeMaCustomPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaCustomPolicyList")
	return
}

func NewDescribeMaCustomPolicyListResponse() (response *DescribeMaCustomPolicyListResponse) {
	response = &DescribeMaCustomPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警自定义白名单
func (c *Client) DescribeMaCustomPolicyList(request *DescribeMaCustomPolicyListRequest) (response *DescribeMaCustomPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeMaCustomPolicyListRequest()
	}
	response = NewDescribeMaCustomPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
	request = &DescribeOpenPortStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeOpenPortStatistics")
	return
}

func NewDescribeOpenPortStatisticsResponse() (response *DescribeOpenPortStatisticsResponse) {
	response = &DescribeOpenPortStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeOpenPortStatistics) 用于获取端口统计列表。
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeOpenPortStatisticsRequest()
	}
	response = NewDescribeOpenPortStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewEditBashRulesRequest() (request *EditBashRulesRequest) {
	request = &EditBashRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "EditBashRules")
	return
}

func NewEditBashRulesResponse() (response *EditBashRulesResponse) {
	response = &EditBashRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改高危命令规则
func (c *Client) EditBashRules(request *EditBashRulesRequest) (response *EditBashRulesResponse, err error) {
	if request == nil {
		request = NewEditBashRulesRequest()
	}
	response = NewEditBashRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanTaskDetailsRequest() (request *DescribeScanTaskDetailsRequest) {
	request = &DescribeScanTaskDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanTaskDetails")
	return
}

func NewDescribeScanTaskDetailsResponse() (response *DescribeScanTaskDetailsResponse) {
	response = &DescribeScanTaskDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeScanTaskDetails 查询扫描任务详情 , 可以查询扫描进度信息/异常;
func (c *Client) DescribeScanTaskDetails(request *DescribeScanTaskDetailsRequest) (response *DescribeScanTaskDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeScanTaskDetailsRequest()
	}
	response = NewDescribeScanTaskDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebHookReceiverUsageRequest() (request *DescribeWebHookReceiverUsageRequest) {
	request = &DescribeWebHookReceiverUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookReceiverUsage")
	return
}

func NewDescribeWebHookReceiverUsageResponse() (response *DescribeWebHookReceiverUsageResponse) {
	response = &DescribeWebHookReceiverUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定告警接收人的关联策略使用信息
func (c *Client) DescribeWebHookReceiverUsage(request *DescribeWebHookReceiverUsageRequest) (response *DescribeWebHookReceiverUsageResponse, err error) {
	if request == nil {
		request = NewDescribeWebHookReceiverUsageRequest()
	}
	response = NewDescribeWebHookReceiverUsageResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineFixListRequest() (request *ExportBaselineFixListRequest) {
	request = &ExportBaselineFixListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineFixList")
	return
}

func NewExportBaselineFixListResponse() (response *ExportBaselineFixListResponse) {
	response = &ExportBaselineFixListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出修复列表
func (c *Client) ExportBaselineFixList(request *ExportBaselineFixListRequest) (response *ExportBaselineFixListResponse, err error) {
	if request == nil {
		request = NewExportBaselineFixListRequest()
	}
	response = NewExportBaselineFixListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackStatisticsRequest() (request *DescribeAttackStatisticsRequest) {
	request = &DescribeAttackStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackStatistics")
	return
}

func NewDescribeAttackStatisticsResponse() (response *DescribeAttackStatisticsResponse) {
	response = &DescribeAttackStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeAttackStatistics(request *DescribeAttackStatisticsRequest) (response *DescribeAttackStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackStatisticsRequest()
	}
	response = NewDescribeAttackStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaWorkloadListRequest() (request *DescribeMaWorkloadListRequest) {
	request = &DescribeMaWorkloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaWorkloadList")
	return
}

func NewDescribeMaWorkloadListResponse() (response *DescribeMaWorkloadListResponse) {
	response = &DescribeMaWorkloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警负载列表
func (c *Client) DescribeMaWorkloadList(request *DescribeMaWorkloadListRequest) (response *DescribeMaWorkloadListResponse, err error) {
	if request == nil {
		request = NewDescribeMaWorkloadListRequest()
	}
	response = NewDescribeMaWorkloadListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenProtectionStatRequest() (request *DescribeScreenProtectionStatRequest) {
	request = &DescribeScreenProtectionStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenProtectionStat")
	return
}

func NewDescribeScreenProtectionStatResponse() (response *DescribeScreenProtectionStatResponse) {
	response = &DescribeScreenProtectionStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏获取安全防护状态
func (c *Client) DescribeScreenProtectionStat(request *DescribeScreenProtectionStatRequest) (response *DescribeScreenProtectionStatResponse, err error) {
	if request == nil {
		request = NewDescribeScreenProtectionStatRequest()
	}
	response = NewDescribeScreenProtectionStatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogClsDeliverSwitchRequest() (request *ModifyLogClsDeliverSwitchRequest) {
	request = &ModifyLogClsDeliverSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogClsDeliverSwitch")
	return
}

func NewModifyLogClsDeliverSwitchResponse() (response *ModifyLogClsDeliverSwitchResponse) {
	response = &ModifyLogClsDeliverSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改cls的投递开关
func (c *Client) ModifyLogClsDeliverSwitch(request *ModifyLogClsDeliverSwitchRequest) (response *ModifyLogClsDeliverSwitchResponse, err error) {
	if request == nil {
		request = NewModifyLogClsDeliverSwitchRequest()
	}
	response = NewModifyLogClsDeliverSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrivilegeRulesRequest() (request *DeletePrivilegeRulesRequest) {
	request = &DeletePrivilegeRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeletePrivilegeRules")
	return
}

func NewDeletePrivilegeRulesResponse() (response *DeletePrivilegeRulesResponse) {
	response = &DeletePrivilegeRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除本地提权规则
func (c *Client) DeletePrivilegeRules(request *DeletePrivilegeRulesRequest) (response *DeletePrivilegeRulesResponse, err error) {
	if request == nil {
		request = NewDeletePrivilegeRulesRequest()
	}
	response = NewDeletePrivilegeRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSafeInfoRequest() (request *DescribeSafeInfoRequest) {
	request = &DescribeSafeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSafeInfo")
	return
}

func NewDescribeSafeInfoResponse() (response *DescribeSafeInfoResponse) {
	response = &DescribeSafeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全通知信息
func (c *Client) DescribeSafeInfo(request *DescribeSafeInfoRequest) (response *DescribeSafeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSafeInfoRequest()
	}
	response = NewDescribeSafeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchBashRulesRequest() (request *SwitchBashRulesRequest) {
	request = &SwitchBashRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SwitchBashRules")
	return
}

func NewSwitchBashRulesResponse() (response *SwitchBashRulesResponse) {
	response = &SwitchBashRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换高危命令规则状态
func (c *Client) SwitchBashRules(request *SwitchBashRulesRequest) (response *SwitchBashRulesResponse, err error) {
	if request == nil {
		request = NewSwitchBashRulesRequest()
	}
	response = NewSwitchBashRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleRequest() (request *DescribeBaselineRuleRequest) {
	request = &DescribeBaselineRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRule")
	return
}

func NewDescribeBaselineRuleResponse() (response *DescribeBaselineRuleResponse) {
	response = &DescribeBaselineRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线id查询下属检测项信息
func (c *Client) DescribeBaselineRule(request *DescribeBaselineRuleRequest) (response *DescribeBaselineRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleRequest()
	}
	response = NewDescribeBaselineRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceOverviewRequest() (request *DescribeVulDefenceOverviewRequest) {
	request = &DescribeVulDefenceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceOverview")
	return
}

func NewDescribeVulDefenceOverviewResponse() (response *DescribeVulDefenceOverviewResponse) {
	response = &DescribeVulDefenceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御概览信息，包括事件趋势及插件开启情况
func (c *Client) DescribeVulDefenceOverview(request *DescribeVulDefenceOverviewRequest) (response *DescribeVulDefenceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceOverviewRequest()
	}
	response = NewDescribeVulDefenceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonthInspectionReportRequest() (request *DescribeMonthInspectionReportRequest) {
	request = &DescribeMonthInspectionReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMonthInspectionReport")
	return
}

func NewDescribeMonthInspectionReportResponse() (response *DescribeMonthInspectionReportResponse) {
	response = &DescribeMonthInspectionReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专家服务-安全管家月巡检报告下载
func (c *Client) DescribeMonthInspectionReport(request *DescribeMonthInspectionReportRequest) (response *DescribeMonthInspectionReportResponse, err error) {
	if request == nil {
		request = NewDescribeMonthInspectionReportRequest()
	}
	response = NewDescribeMonthInspectionReportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaWorkloadLabelRequest() (request *ModifyMaWorkloadLabelRequest) {
	request = &ModifyMaWorkloadLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaWorkloadLabel")
	return
}

func NewModifyMaWorkloadLabelResponse() (response *ModifyMaWorkloadLabelResponse) {
	response = &ModifyMaWorkloadLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改微告警负载标签
func (c *Client) ModifyMaWorkloadLabel(request *ModifyMaWorkloadLabelRequest) (response *ModifyMaWorkloadLabelResponse, err error) {
	if request == nil {
		request = NewModifyMaWorkloadLabelRequest()
	}
	response = NewModifyMaWorkloadLabelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceEventStatusRequest() (request *ModifyVulDefenceEventStatusRequest) {
	request = &ModifyVulDefenceEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyVulDefenceEventStatus")
	return
}

func NewModifyVulDefenceEventStatusResponse() (response *ModifyVulDefenceEventStatusResponse) {
	response = &ModifyVulDefenceEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御事件状态（修复漏洞通过其他接口实现）
func (c *Client) ModifyVulDefenceEventStatus(request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceEventStatusRequest()
	}
	response = NewModifyVulDefenceEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaWhitelistRequest() (request *ModifyMaWhitelistRequest) {
	request = &ModifyMaWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaWhitelist")
	return
}

func NewModifyMaWhitelistResponse() (response *ModifyMaWhitelistResponse) {
	response = &ModifyMaWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改微告警全局白名单
func (c *Client) ModifyMaWhitelist(request *ModifyMaWhitelistRequest) (response *ModifyMaWhitelistResponse, err error) {
	if request == nil {
		request = NewModifyMaWhitelistRequest()
	}
	response = NewModifyMaWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellRulesAggregationRequest() (request *ModifyReverseShellRulesAggregationRequest) {
	request = &ModifyReverseShellRulesAggregationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyReverseShellRulesAggregation")
	return
}

func NewModifyReverseShellRulesAggregationResponse() (response *ModifyReverseShellRulesAggregationResponse) {
	response = &ModifyReverseShellRulesAggregationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑反弹Shell规则（支持多服务器选择）-聚合版本-支持正则
func (c *Client) ModifyReverseShellRulesAggregation(request *ModifyReverseShellRulesAggregationRequest) (response *ModifyReverseShellRulesAggregationResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellRulesAggregationRequest()
	}
	response = NewModifyReverseShellRulesAggregationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAllJavaMemShellsRequest() (request *DeleteAllJavaMemShellsRequest) {
	request = &DeleteAllJavaMemShellsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteAllJavaMemShells")
	return
}

func NewDeleteAllJavaMemShellsResponse() (response *DeleteAllJavaMemShellsResponse) {
	response = &DeleteAllJavaMemShellsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除全部java内存马事件
func (c *Client) DeleteAllJavaMemShells(request *DeleteAllJavaMemShellsRequest) (response *DeleteAllJavaMemShellsResponse, err error) {
	if request == nil {
		request = NewDeleteAllJavaMemShellsRequest()
	}
	response = NewDeleteAllJavaMemShellsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyShellPolicyStatusRequest() (request *ModifyShellPolicyStatusRequest) {
	request = &ModifyShellPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyShellPolicyStatus")
	return
}

func NewModifyShellPolicyStatusResponse() (response *ModifyShellPolicyStatusResponse) {
	response = &ModifyShellPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑反弹shell策略开关
func (c *Client) ModifyShellPolicyStatus(request *ModifyShellPolicyStatusRequest) (response *ModifyShellPolicyStatusResponse, err error) {
	if request == nil {
		request = NewModifyShellPolicyStatusRequest()
	}
	response = NewModifyShellPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInquireRenewFlagRequest() (request *ModifyInquireRenewFlagRequest) {
	request = &ModifyInquireRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyInquireRenewFlag")
	return
}

func NewModifyInquireRenewFlagResponse() (response *ModifyInquireRenewFlagResponse) {
	response = &ModifyInquireRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置旗舰版试用订单到期新购状态
func (c *Client) ModifyInquireRenewFlag(request *ModifyInquireRenewFlagRequest) (response *ModifyInquireRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyInquireRenewFlagRequest()
	}
	response = NewModifyInquireRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineEffectHostListRequest() (request *ExportBaselineEffectHostListRequest) {
	request = &ExportBaselineEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineEffectHostList")
	return
}

func NewExportBaselineEffectHostListResponse() (response *ExportBaselineEffectHostListResponse) {
	response = &ExportBaselineEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线影响主机列表
func (c *Client) ExportBaselineEffectHostList(request *ExportBaselineEffectHostListRequest) (response *ExportBaselineEffectHostListResponse, err error) {
	if request == nil {
		request = NewExportBaselineEffectHostListRequest()
	}
	response = NewExportBaselineEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetsProtectionStatusRequest() (request *ModifyAssetsProtectionStatusRequest) {
	request = &ModifyAssetsProtectionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyAssetsProtectionStatus")
	return
}

func NewModifyAssetsProtectionStatusResponse() (response *ModifyAssetsProtectionStatusResponse) {
	response = &ModifyAssetsProtectionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改主机防护状态
func (c *Client) ModifyAssetsProtectionStatus(request *ModifyAssetsProtectionStatusRequest) (response *ModifyAssetsProtectionStatusResponse, err error) {
	if request == nil {
		request = NewModifyAssetsProtectionStatusRequest()
	}
	response = NewModifyAssetsProtectionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashEventsRequest() (request *DescribeBashEventsRequest) {
	request = &DescribeBashEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEvents")
	return
}

func NewDescribeBashEventsResponse() (response *DescribeBashEventsResponse) {
	response = &DescribeBashEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高危命令列表
func (c *Client) DescribeBashEvents(request *DescribeBashEventsRequest) (response *DescribeBashEventsResponse, err error) {
	if request == nil {
		request = NewDescribeBashEventsRequest()
	}
	response = NewDescribeBashEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostLoginListRequest() (request *DescribeHostLoginListRequest) {
	request = &DescribeHostLoginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeHostLoginList")
	return
}

func NewDescribeHostLoginListResponse() (response *DescribeHostLoginListResponse) {
	response = &DescribeHostLoginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取登录审计列表
func (c *Client) DescribeHostLoginList(request *DescribeHostLoginListRequest) (response *DescribeHostLoginListResponse, err error) {
	if request == nil {
		request = NewDescribeHostLoginListRequest()
	}
	response = NewDescribeHostLoginListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseEventsListRequest() (request *DescribeRansomDefenseEventsListRequest) {
	request = &DescribeRansomDefenseEventsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseEventsList")
	return
}

func NewDescribeRansomDefenseEventsListResponse() (response *DescribeRansomDefenseEventsListResponse) {
	response = &DescribeRansomDefenseEventsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防勒索事件列表
func (c *Client) DescribeRansomDefenseEventsList(request *DescribeRansomDefenseEventsListRequest) (response *DescribeRansomDefenseEventsListResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseEventsListRequest()
	}
	response = NewDescribeRansomDefenseEventsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStrategyExistRequest() (request *DescribeStrategyExistRequest) {
	request = &DescribeStrategyExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeStrategyExist")
	return
}

func NewDescribeStrategyExistResponse() (response *DescribeStrategyExistResponse) {
	response = &DescribeStrategyExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据策略名查询策略是否存在
func (c *Client) DescribeStrategyExist(request *DescribeStrategyExistRequest) (response *DescribeStrategyExistResponse, err error) {
	if request == nil {
		request = NewDescribeStrategyExistRequest()
	}
	response = NewDescribeStrategyExistResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetSystemPackageListRequest() (request *ExportAssetSystemPackageListRequest) {
	request = &ExportAssetSystemPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetSystemPackageList")
	return
}

func NewExportAssetSystemPackageListResponse() (response *ExportAssetSystemPackageListResponse) {
	response = &ExportAssetSystemPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理系统安装包列表
func (c *Client) ExportAssetSystemPackageList(request *ExportAssetSystemPackageListRequest) (response *ExportAssetSystemPackageListResponse, err error) {
	if request == nil {
		request = NewExportAssetSystemPackageListRequest()
	}
	response = NewExportAssetSystemPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineRiskCntRequest() (request *DescribeMachineRiskCntRequest) {
	request = &DescribeMachineRiskCntRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRiskCnt")
	return
}

func NewDescribeMachineRiskCntResponse() (response *DescribeMachineRiskCntResponse) {
	response = &DescribeMachineRiskCntResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机入侵检测事件统计
func (c *Client) DescribeMachineRiskCnt(request *DescribeMachineRiskCntRequest) (response *DescribeMachineRiskCntResponse, err error) {
	if request == nil {
		request = NewDescribeMachineRiskCntRequest()
	}
	response = NewDescribeMachineRiskCntResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaLabelRequest() (request *DescribeMaLabelRequest) {
	request = &DescribeMaLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaLabel")
	return
}

func NewDescribeMaLabelResponse() (response *DescribeMaLabelResponse) {
	response = &DescribeMaLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警标签信息
func (c *Client) DescribeMaLabel(request *DescribeMaLabelRequest) (response *DescribeMaLabelResponse, err error) {
	if request == nil {
		request = NewDescribeMaLabelRequest()
	}
	response = NewDescribeMaLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellSystemPolicyConfigRequest() (request *DescribeReverseShellSystemPolicyConfigRequest) {
	request = &DescribeReverseShellSystemPolicyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellSystemPolicyConfig")
	return
}

func NewDescribeReverseShellSystemPolicyConfigResponse() (response *DescribeReverseShellSystemPolicyConfigResponse) {
	response = &DescribeReverseShellSystemPolicyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell系统策略配置
func (c *Client) DescribeReverseShellSystemPolicyConfig(request *DescribeReverseShellSystemPolicyConfigRequest) (response *DescribeReverseShellSystemPolicyConfigResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellSystemPolicyConfigRequest()
	}
	response = NewDescribeReverseShellSystemPolicyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetAttackSettingRequest() (request *DescribeNetAttackSettingRequest) {
	request = &DescribeNetAttackSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeNetAttackSetting")
	return
}

func NewDescribeNetAttackSettingResponse() (response *DescribeNetAttackSettingResponse) {
	response = &DescribeNetAttackSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeNetAttackSetting(request *DescribeNetAttackSettingRequest) (response *DescribeNetAttackSettingResponse, err error) {
	if request == nil {
		request = NewDescribeNetAttackSettingRequest()
	}
	response = NewDescribeNetAttackSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBaselineRuleRequest() (request *DeleteBaselineRuleRequest) {
	request = &DeleteBaselineRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineRule")
	return
}

func NewDeleteBaselineRuleResponse() (response *DeleteBaselineRuleResponse) {
	response = &DeleteBaselineRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除基线规则
func (c *Client) DeleteBaselineRule(request *DeleteBaselineRuleRequest) (response *DeleteBaselineRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBaselineRuleRequest()
	}
	response = NewDeleteBaselineRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteScanTaskRequest() (request *DeleteScanTaskRequest) {
	request = &DeleteScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteScanTask")
	return
}

func NewDeleteScanTaskResponse() (response *DeleteScanTaskResponse) {
	response = &DeleteScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteScanTask 该接口可以对指定类型的扫描任务进行停止扫描;
func (c *Client) DeleteScanTask(request *DeleteScanTaskRequest) (response *DeleteScanTaskResponse, err error) {
	if request == nil {
		request = NewDeleteScanTaskRequest()
	}
	response = NewDeleteScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewEditReverseShellRulesRequest() (request *EditReverseShellRulesRequest) {
	request = &EditReverseShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "EditReverseShellRules")
	return
}

func NewEditReverseShellRulesResponse() (response *EditReverseShellRulesResponse) {
	response = &EditReverseShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑反弹Shell规则（支持多服务器选择）
func (c *Client) EditReverseShellRules(request *EditReverseShellRulesRequest) (response *EditReverseShellRulesResponse, err error) {
	if request == nil {
		request = NewEditReverseShellRulesRequest()
	}
	response = NewEditReverseShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewScanTaskAgainRequest() (request *ScanTaskAgainRequest) {
	request = &ScanTaskAgainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanTaskAgain")
	return
}

func NewScanTaskAgainResponse() (response *ScanTaskAgainResponse) {
	response = &ScanTaskAgainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ScanTaskAgain  重新开始扫描任务，可以指定机器
func (c *Client) ScanTaskAgain(request *ScanTaskAgainRequest) (response *ScanTaskAgainResponse, err error) {
	if request == nil {
		request = NewScanTaskAgainRequest()
	}
	response = NewScanTaskAgainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCanFixVulMachineRequest() (request *DescribeCanFixVulMachineRequest) {
	request = &DescribeCanFixVulMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeCanFixVulMachine")
	return
}

func NewDescribeCanFixVulMachineResponse() (response *DescribeCanFixVulMachineResponse) {
	response = &DescribeCanFixVulMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞修护-查询可修护主机信息
func (c *Client) DescribeCanFixVulMachine(request *DescribeCanFixVulMachineRequest) (response *DescribeCanFixVulMachineResponse, err error) {
	if request == nil {
		request = NewDescribeCanFixVulMachineRequest()
	}
	response = NewDescribeCanFixVulMachineResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEventAttackStatusRequest() (request *ModifyEventAttackStatusRequest) {
	request = &ModifyEventAttackStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyEventAttackStatus")
	return
}

func NewModifyEventAttackStatusResponse() (response *ModifyEventAttackStatusResponse) {
	response = &ModifyEventAttackStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) ModifyEventAttackStatus(request *ModifyEventAttackStatusRequest) (response *ModifyEventAttackStatusResponse, err error) {
	if request == nil {
		request = NewModifyEventAttackStatusRequest()
	}
	response = NewModifyEventAttackStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRetryCreateSnapshotRequest() (request *RetryCreateSnapshotRequest) {
	request = &RetryCreateSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RetryCreateSnapshot")
	return
}

func NewRetryCreateSnapshotResponse() (response *RetryCreateSnapshotResponse) {
	response = &RetryCreateSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 快照创建失败时可以重试创建快照并且自动进行漏洞修复
func (c *Client) RetryCreateSnapshot(request *RetryCreateSnapshotRequest) (response *RetryCreateSnapshotResponse, err error) {
	if request == nil {
		request = NewRetryCreateSnapshotRequest()
	}
	response = NewRetryCreateSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewEditPrivilegeRulesRequest() (request *EditPrivilegeRulesRequest) {
	request = &EditPrivilegeRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "EditPrivilegeRules")
	return
}

func NewEditPrivilegeRulesResponse() (response *EditPrivilegeRulesResponse) {
	response = &EditPrivilegeRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改本地提权规则（支持多服务器选择）
func (c *Client) EditPrivilegeRules(request *EditPrivilegeRulesRequest) (response *EditPrivilegeRulesResponse, err error) {
	if request == nil {
		request = NewEditPrivilegeRulesRequest()
	}
	response = NewEditPrivilegeRulesResponse()
	err = c.Send(request, response)
	return
}

func NewExportIgnoreBaselineRuleRequest() (request *ExportIgnoreBaselineRuleRequest) {
	request = &ExportIgnoreBaselineRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportIgnoreBaselineRule")
	return
}

func NewExportIgnoreBaselineRuleResponse() (response *ExportIgnoreBaselineRuleResponse) {
	response = &ExportIgnoreBaselineRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出已忽略基线检测项信息
func (c *Client) ExportIgnoreBaselineRule(request *ExportIgnoreBaselineRuleRequest) (response *ExportIgnoreBaselineRuleResponse, err error) {
	if request == nil {
		request = NewExportIgnoreBaselineRuleRequest()
	}
	response = NewExportIgnoreBaselineRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
	request = &DescribeAlarmAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAlarmAttribute")
	return
}

func NewDescribeAlarmAttributeResponse() (response *DescribeAlarmAttributeResponse) {
	response = &DescribeAlarmAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAlarmAttribute) 用于获取告警设置。
func (c *Client) DescribeAlarmAttribute(request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmAttributeRequest()
	}
	response = NewDescribeAlarmAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppProcessListRequest() (request *DescribeAssetAppProcessListRequest) {
	request = &DescribeAssetAppProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppProcessList")
	return
}

func NewDescribeAssetAppProcessListResponse() (response *DescribeAssetAppProcessListResponse) {
	response = &DescribeAssetAppProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取软件关联进程列表
func (c *Client) DescribeAssetAppProcessList(request *DescribeAssetAppProcessListRequest) (response *DescribeAssetAppProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppProcessListRequest()
	}
	response = NewDescribeAssetAppProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBanModeRequest() (request *DescribeBanModeRequest) {
	request = &DescribeBanModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanMode")
	return
}

func NewDescribeBanModeResponse() (response *DescribeBanModeResponse) {
	response = &DescribeBanModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取爆破阻断模式
func (c *Client) DescribeBanMode(request *DescribeBanModeRequest) (response *DescribeBanModeResponse, err error) {
	if request == nil {
		request = NewDescribeBanModeRequest()
	}
	response = NewDescribeBanModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
	request = &DescribeVulsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVuls")
	return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
	response = &DescribeVulsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeVuls) 用于获取漏洞列表数据。
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
	if request == nil {
		request = NewDescribeVulsRequest()
	}
	response = NewDescribeVulsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserInfoRequest() (request *DescribeAssetUserInfoRequest) {
	request = &DescribeAssetUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserInfo")
	return
}

func NewDescribeAssetUserInfoResponse() (response *DescribeAssetUserInfoResponse) {
	response = &DescribeAssetUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机账号详情
func (c *Client) DescribeAssetUserInfo(request *DescribeAssetUserInfoRequest) (response *DescribeAssetUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserInfoRequest()
	}
	response = NewDescribeAssetUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
	request = &DescribeLogHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogHistogram")
	return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
	response = &DescribeLogHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志直方图信息
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeLogHistogramRequest()
	}
	response = NewDescribeLogHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineListRequest() (request *DescribeMachineListRequest) {
	request = &DescribeMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineList")
	return
}

func NewDescribeMachineListResponse() (response *DescribeMachineListResponse) {
	response = &DescribeMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于网页防篡改获取区域主机列表。
func (c *Client) DescribeMachineList(request *DescribeMachineListRequest) (response *DescribeMachineListResponse, err error) {
	if request == nil {
		request = NewDescribeMachineListRequest()
	}
	response = NewDescribeMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectDirRelatedServerRequest() (request *DescribeProtectDirRelatedServerRequest) {
	request = &DescribeProtectDirRelatedServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProtectDirRelatedServer")
	return
}

func NewDescribeProtectDirRelatedServerResponse() (response *DescribeProtectDirRelatedServerResponse) {
	response = &DescribeProtectDirRelatedServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防护目录关联服务器列表信息
func (c *Client) DescribeProtectDirRelatedServer(request *DescribeProtectDirRelatedServerRequest) (response *DescribeProtectDirRelatedServerResponse, err error) {
	if request == nil {
		request = NewDescribeProtectDirRelatedServerRequest()
	}
	response = NewDescribeProtectDirRelatedServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
	request = &DescribeOverviewStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeOverviewStatistics")
	return
}

func NewDescribeOverviewStatisticsResponse() (response *DescribeOverviewStatisticsResponse) {
	response = &DescribeOverviewStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取概览统计数据。
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewStatisticsRequest()
	}
	response = NewDescribeOverviewStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemRiskTopRequest() (request *DescribeBaselineItemRiskTopRequest) {
	request = &DescribeBaselineItemRiskTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemRiskTop")
	return
}

func NewDescribeBaselineItemRiskTopResponse() (response *DescribeBaselineItemRiskTopResponse) {
	response = &DescribeBaselineItemRiskTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测项TOP5
func (c *Client) DescribeBaselineItemRiskTop(request *DescribeBaselineItemRiskTopRequest) (response *DescribeBaselineItemRiskTopResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemRiskTopRequest()
	}
	response = NewDescribeBaselineItemRiskTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineRegionListRequest() (request *DescribeMachineRegionListRequest) {
	request = &DescribeMachineRegionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRegionList")
	return
}

func NewDescribeMachineRegionListResponse() (response *DescribeMachineRegionListResponse) {
	response = &DescribeMachineRegionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机器地域列表
func (c *Client) DescribeMachineRegionList(request *DescribeMachineRegionListRequest) (response *DescribeMachineRegionListResponse, err error) {
	if request == nil {
		request = NewDescribeMachineRegionListRequest()
	}
	response = NewDescribeMachineRegionListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServersAndRiskAndFirstInfoRequest() (request *DescribeServersAndRiskAndFirstInfoRequest) {
	request = &DescribeServersAndRiskAndFirstInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeServersAndRiskAndFirstInfo")
	return
}

func NewDescribeServersAndRiskAndFirstInfoResponse() (response *DescribeServersAndRiskAndFirstInfoResponse) {
	response = &DescribeServersAndRiskAndFirstInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取待处理风险文件数+影响服务器数+是否试用检测+最近检测时间
func (c *Client) DescribeServersAndRiskAndFirstInfo(request *DescribeServersAndRiskAndFirstInfoRequest) (response *DescribeServersAndRiskAndFirstInfoResponse, err error) {
	if request == nil {
		request = NewDescribeServersAndRiskAndFirstInfoRequest()
	}
	response = NewDescribeServersAndRiskAndFirstInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeChangingAssetInfoRequest() (request *DescribeChangingAssetInfoRequest) {
	request = &DescribeChangingAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeChangingAssetInfo")
	return
}

func NewDescribeChangingAssetInfoResponse() (response *DescribeChangingAssetInfoResponse) {
	response = &DescribeChangingAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取修改中的主机列表
func (c *Client) DescribeChangingAssetInfo(request *DescribeChangingAssetInfoRequest) (response *DescribeChangingAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeChangingAssetInfoRequest()
	}
	response = NewDescribeChangingAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogExportsRequest() (request *DescribeLogExportsRequest) {
	request = &DescribeLogExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogExports")
	return
}

func NewDescribeLogExportsResponse() (response *DescribeLogExportsResponse) {
	response = &DescribeLogExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志下载任务列表
func (c *Client) DescribeLogExports(request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
	if request == nil {
		request = NewDescribeLogExportsRequest()
	}
	response = NewDescribeLogExportsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationPathListRequest() (request *DescribeAssetWebLocationPathListRequest) {
	request = &DescribeAssetWebLocationPathListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationPathList")
	return
}

func NewDescribeAssetWebLocationPathListResponse() (response *DescribeAssetWebLocationPathListResponse) {
	response = &DescribeAssetWebLocationPathListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点虚拟目录列表
func (c *Client) DescribeAssetWebLocationPathList(request *DescribeAssetWebLocationPathListRequest) (response *DescribeAssetWebLocationPathListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationPathListRequest()
	}
	response = NewDescribeAssetWebLocationPathListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortCountRequest() (request *DescribeAssetPortCountRequest) {
	request = &DescribeAssetPortCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPortCount")
	return
}

func NewDescribeAssetPortCountResponse() (response *DescribeAssetPortCountResponse) {
	response = &DescribeAssetPortCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有端口数量
func (c *Client) DescribeAssetPortCount(request *DescribeAssetPortCountRequest) (response *DescribeAssetPortCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortCountRequest()
	}
	response = NewDescribeAssetPortCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaSelfLearnPolicyListRequest() (request *ExportMaSelfLearnPolicyListRequest) {
	request = &ExportMaSelfLearnPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaSelfLearnPolicyList")
	return
}

func NewExportMaSelfLearnPolicyListResponse() (response *ExportMaSelfLearnPolicyListResponse) {
	response = &ExportMaSelfLearnPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警自学习子策略
func (c *Client) ExportMaSelfLearnPolicyList(request *ExportMaSelfLearnPolicyListRequest) (response *ExportMaSelfLearnPolicyListResponse, err error) {
	if request == nil {
		request = NewExportMaSelfLearnPolicyListRequest()
	}
	response = NewExportMaSelfLearnPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenBroadcastsRequest() (request *DescribeScreenBroadcastsRequest) {
	request = &DescribeScreenBroadcastsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenBroadcasts")
	return
}

func NewDescribeScreenBroadcastsResponse() (response *DescribeScreenBroadcastsResponse) {
	response = &DescribeScreenBroadcastsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化安全播报
func (c *Client) DescribeScreenBroadcasts(request *DescribeScreenBroadcastsRequest) (response *DescribeScreenBroadcastsResponse, err error) {
	if request == nil {
		request = NewDescribeScreenBroadcastsRequest()
	}
	response = NewDescribeScreenBroadcastsResponse()
	err = c.Send(request, response)
	return
}

func NewRansomDefenseRollbackRequest() (request *RansomDefenseRollbackRequest) {
	request = &RansomDefenseRollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RansomDefenseRollback")
	return
}

func NewRansomDefenseRollbackResponse() (response *RansomDefenseRollbackResponse) {
	response = &RansomDefenseRollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防勒索快照回滚
func (c *Client) RansomDefenseRollback(request *RansomDefenseRollbackRequest) (response *RansomDefenseRollbackResponse, err error) {
	if request == nil {
		request = NewRansomDefenseRollbackRequest()
	}
	response = NewRansomDefenseRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebHookPolicyRequest() (request *ModifyWebHookPolicyRequest) {
	request = &ModifyWebHookPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookPolicy")
	return
}

func NewModifyWebHookPolicyResponse() (response *ModifyWebHookPolicyResponse) {
	response = &ModifyWebHookPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改告警策略
func (c *Client) ModifyWebHookPolicy(request *ModifyWebHookPolicyRequest) (response *ModifyWebHookPolicyResponse, err error) {
	if request == nil {
		request = NewModifyWebHookPolicyRequest()
	}
	response = NewModifyWebHookPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetAttackWhiteListRequest() (request *ModifyNetAttackWhiteListRequest) {
	request = &ModifyNetAttackWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyNetAttackWhiteList")
	return
}

func NewModifyNetAttackWhiteListResponse() (response *ModifyNetAttackWhiteListResponse) {
	response = &ModifyNetAttackWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) ModifyNetAttackWhiteList(request *ModifyNetAttackWhiteListRequest) (response *ModifyNetAttackWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyNetAttackWhiteListRequest()
	}
	response = NewModifyNetAttackWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMemShellRulesRequest() (request *DeleteMemShellRulesRequest) {
	request = &DeleteMemShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMemShellRules")
	return
}

func NewDeleteMemShellRulesResponse() (response *DeleteMemShellRulesResponse) {
	response = &DeleteMemShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除内存马白名单规则
func (c *Client) DeleteMemShellRules(request *DeleteMemShellRulesRequest) (response *DeleteMemShellRulesResponse, err error) {
	if request == nil {
		request = NewDeleteMemShellRulesRequest()
	}
	response = NewDeleteMemShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineHostDetectListRequest() (request *DescribeBaselineHostDetectListRequest) {
	request = &DescribeBaselineHostDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostDetectList")
	return
}

func NewDescribeBaselineHostDetectListResponse() (response *DescribeBaselineHostDetectListResponse) {
	response = &DescribeBaselineHostDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测主机列表
func (c *Client) DescribeBaselineHostDetectList(request *DescribeBaselineHostDetectListRequest) (response *DescribeBaselineHostDetectListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineHostDetectListRequest()
	}
	response = NewDescribeBaselineHostDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMachineTagsRequest() (request *UpdateMachineTagsRequest) {
	request = &UpdateMachineTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "UpdateMachineTags")
	return
}

func NewUpdateMachineTagsResponse() (response *UpdateMachineTagsResponse) {
	response = &UpdateMachineTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联机器标签列表
func (c *Client) UpdateMachineTags(request *UpdateMachineTagsRequest) (response *UpdateMachineTagsResponse, err error) {
	if request == nil {
		request = NewUpdateMachineTagsRequest()
	}
	response = NewUpdateMachineTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppListRequest() (request *DescribeAssetWebAppListRequest) {
	request = &DescribeAssetWebAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppList")
	return
}

func NewDescribeAssetWebAppListResponse() (response *DescribeAssetWebAppListResponse) {
	response = &DescribeAssetWebAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理Web应用列表
func (c *Client) DescribeAssetWebAppList(request *DescribeAssetWebAppListRequest) (response *DescribeAssetWebAppListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppListRequest()
	}
	response = NewDescribeAssetWebAppListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllMachinesRequest() (request *DescribeAllMachinesRequest) {
	request = &DescribeAllMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAllMachines")
	return
}

func NewDescribeAllMachinesResponse() (response *DescribeAllMachinesResponse) {
	response = &DescribeAllMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAllMachines）按所有获取机器列表。
func (c *Client) DescribeAllMachines(request *DescribeAllMachinesRequest) (response *DescribeAllMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeAllMachinesRequest()
	}
	response = NewDescribeAllMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInstallCommandRequest() (request *DescribeAgentInstallCommandRequest) {
	request = &DescribeAgentInstallCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallCommand")
	return
}

func NewDescribeAgentInstallCommandResponse() (response *DescribeAgentInstallCommandResponse) {
	response = &DescribeAgentInstallCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取agent安装命令
func (c *Client) DescribeAgentInstallCommand(request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInstallCommandRequest()
	}
	response = NewDescribeAgentInstallCommandResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineListRequest() (request *DescribeAssetMachineListRequest) {
	request = &DescribeAssetMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineList")
	return
}

func NewDescribeAssetMachineListResponse() (response *DescribeAssetMachineListResponse) {
	response = &DescribeAssetMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产指纹页面的资源监控列表
func (c *Client) DescribeAssetMachineList(request *DescribeAssetMachineListRequest) (response *DescribeAssetMachineListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineListRequest()
	}
	response = NewDescribeAssetMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseListRequest() (request *DescribeLicenseListRequest) {
	request = &DescribeLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseList")
	return
}

func NewDescribeLicenseListResponse() (response *DescribeLicenseListResponse) {
	response = &DescribeLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户所有授权订单信息
func (c *Client) DescribeLicenseList(request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseListRequest()
	}
	response = NewDescribeLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMalwareScanTaskRequest() (request *DeleteMalwareScanTaskRequest) {
	request = &DeleteMalwareScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwareScanTask")
	return
}

func NewDeleteMalwareScanTaskResponse() (response *DeleteMalwareScanTaskResponse) {
	response = &DeleteMalwareScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵管理-终止扫描任务
func (c *Client) DeleteMalwareScanTask(request *DeleteMalwareScanTaskRequest) (response *DeleteMalwareScanTaskResponse, err error) {
	if request == nil {
		request = NewDeleteMalwareScanTaskRequest()
	}
	response = NewDeleteMalwareScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetScanTaskRequest() (request *CreateAssetScanTaskRequest) {
	request = &CreateAssetScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateAssetScanTask")
	return
}

func NewCreateAssetScanTaskResponse() (response *CreateAssetScanTaskResponse) {
	response = &CreateAssetScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资产扫描任务
func (c *Client) CreateAssetScanTask(request *CreateAssetScanTaskRequest) (response *CreateAssetScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateAssetScanTaskRequest()
	}
	response = NewCreateAssetScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortInfoListRequest() (request *DescribeAssetPortInfoListRequest) {
	request = &DescribeAssetPortInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPortInfoList")
	return
}

func NewDescribeAssetPortInfoListResponse() (response *DescribeAssetPortInfoListResponse) {
	response = &DescribeAssetPortInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理端口列表
func (c *Client) DescribeAssetPortInfoList(request *DescribeAssetPortInfoListRequest) (response *DescribeAssetPortInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortInfoListRequest()
	}
	response = NewDescribeAssetPortInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetRecentMachineInfoRequest() (request *ExportAssetRecentMachineInfoRequest) {
	request = &ExportAssetRecentMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetRecentMachineInfo")
	return
}

func NewExportAssetRecentMachineInfoResponse() (response *ExportAssetRecentMachineInfoResponse) {
	response = &ExportAssetRecentMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出主机最近趋势情况（最长最近90天）
func (c *Client) ExportAssetRecentMachineInfo(request *ExportAssetRecentMachineInfoRequest) (response *ExportAssetRecentMachineInfoResponse, err error) {
	if request == nil {
		request = NewExportAssetRecentMachineInfoRequest()
	}
	response = NewExportAssetRecentMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClientExceptionRequest() (request *DescribeClientExceptionRequest) {
	request = &DescribeClientExceptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeClientException")
	return
}

func NewDescribeClientExceptionResponse() (response *DescribeClientExceptionResponse) {
	response = &DescribeClientExceptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户端异常事件
func (c *Client) DescribeClientException(request *DescribeClientExceptionRequest) (response *DescribeClientExceptionResponse, err error) {
	if request == nil {
		request = NewDescribeClientExceptionRequest()
	}
	response = NewDescribeClientExceptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogKafkaDeliverInfoRequest() (request *DescribeLogKafkaDeliverInfoRequest) {
	request = &DescribeLogKafkaDeliverInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogKafkaDeliverInfo")
	return
}

func NewDescribeLogKafkaDeliverInfoResponse() (response *DescribeLogKafkaDeliverInfoResponse) {
	response = &DescribeLogKafkaDeliverInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取kafka投递信息
func (c *Client) DescribeLogKafkaDeliverInfo(request *DescribeLogKafkaDeliverInfoRequest) (response *DescribeLogKafkaDeliverInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLogKafkaDeliverInfoRequest()
	}
	response = NewDescribeLogKafkaDeliverInfoResponse()
	err = c.Send(request, response)
	return
}

func NewScanVulRequest() (request *ScanVulRequest) {
	request = &ScanVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanVul")
	return
}

func NewScanVulResponse() (response *ScanVulResponse) {
	response = &ScanVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键检测
func (c *Client) ScanVul(request *ScanVulRequest) (response *ScanVulResponse, err error) {
	if request == nil {
		request = NewScanVulRequest()
	}
	response = NewScanVulResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLoginPermissionPushRequest() (request *CreateLoginPermissionPushRequest) {
	request = &CreateLoginPermissionPushRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateLoginPermissionPush")
	return
}

func NewCreateLoginPermissionPushResponse() (response *CreateLoginPermissionPushResponse) {
	response = &CreateLoginPermissionPushResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下发登录许可
func (c *Client) CreateLoginPermissionPush(request *CreateLoginPermissionPushRequest) (response *CreateLoginPermissionPushResponse, err error) {
	if request == nil {
		request = NewCreateLoginPermissionPushRequest()
	}
	response = NewCreateLoginPermissionPushResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralStatRequest() (request *DescribeGeneralStatRequest) {
	request = &DescribeGeneralStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeGeneralStat")
	return
}

func NewDescribeGeneralStatResponse() (response *DescribeGeneralStatResponse) {
	response = &DescribeGeneralStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机相关统计
func (c *Client) DescribeGeneralStat(request *DescribeGeneralStatRequest) (response *DescribeGeneralStatResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralStatRequest()
	}
	response = NewDescribeGeneralStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseGeneralRequest() (request *DescribeLicenseGeneralRequest) {
	request = &DescribeLicenseGeneralRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseGeneral")
	return
}

func NewDescribeLicenseGeneralResponse() (response *DescribeLicenseGeneralResponse) {
	response = &DescribeLicenseGeneralResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 授权管理-授权概览信息
func (c *Client) DescribeLicenseGeneral(request *DescribeLicenseGeneralRequest) (response *DescribeLicenseGeneralResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseGeneralRequest()
	}
	response = NewDescribeLicenseGeneralResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetScanTaskDetailRequest() (request *DescribeAssetScanTaskDetailRequest) {
	request = &DescribeAssetScanTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetScanTaskDetail")
	return
}

func NewDescribeAssetScanTaskDetailResponse() (response *DescribeAssetScanTaskDetailResponse) {
	response = &DescribeAssetScanTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产扫描任务详情
func (c *Client) DescribeAssetScanTaskDetail(request *DescribeAssetScanTaskDetailRequest) (response *DescribeAssetScanTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetScanTaskDetailRequest()
	}
	response = NewDescribeAssetScanTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationCountRequest() (request *DescribeAssetWebLocationCountRequest) {
	request = &DescribeAssetWebLocationCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationCount")
	return
}

func NewDescribeAssetWebLocationCountResponse() (response *DescribeAssetWebLocationCountResponse) {
	response = &DescribeAssetWebLocationCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web站点数量
func (c *Client) DescribeAssetWebLocationCount(request *DescribeAssetWebLocationCountRequest) (response *DescribeAssetWebLocationCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationCountRequest()
	}
	response = NewDescribeAssetWebLocationCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrivilegeEventInfoRequest() (request *DescribePrivilegeEventInfoRequest) {
	request = &DescribePrivilegeEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeEventInfo")
	return
}

func NewDescribePrivilegeEventInfoResponse() (response *DescribePrivilegeEventInfoResponse) {
	response = &DescribePrivilegeEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本地提权信息详情
func (c *Client) DescribePrivilegeEventInfo(request *DescribePrivilegeEventInfoRequest) (response *DescribePrivilegeEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribePrivilegeEventInfoRequest()
	}
	response = NewDescribePrivilegeEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBaselinePolicyRequest() (request *DeleteBaselinePolicyRequest) {
	request = &DeleteBaselinePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselinePolicy")
	return
}

func NewDeleteBaselinePolicyResponse() (response *DeleteBaselinePolicyResponse) {
	response = &DeleteBaselinePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除基线策略配置
func (c *Client) DeleteBaselinePolicy(request *DeleteBaselinePolicyRequest) (response *DeleteBaselinePolicyResponse, err error) {
	if request == nil {
		request = NewDeleteBaselinePolicyRequest()
	}
	response = NewDeleteBaselinePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewExportFileTamperEventsRequest() (request *ExportFileTamperEventsRequest) {
	request = &ExportFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportFileTamperEvents")
	return
}

func NewExportFileTamperEventsResponse() (response *ExportFileTamperEventsResponse) {
	response = &ExportFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出核心文件事件
func (c *Client) ExportFileTamperEvents(request *ExportFileTamperEventsRequest) (response *ExportFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewExportFileTamperEventsRequest()
	}
	response = NewExportFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanScheduleRequest() (request *DescribeScanScheduleRequest) {
	request = &DescribeScanScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanSchedule")
	return
}

func NewDescribeScanScheduleResponse() (response *DescribeScanScheduleResponse) {
	response = &DescribeScanScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据taskid查询检测进度
func (c *Client) DescribeScanSchedule(request *DescribeScanScheduleRequest) (response *DescribeScanScheduleResponse, err error) {
	if request == nil {
		request = NewDescribeScanScheduleRequest()
	}
	response = NewDescribeScanScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaGlobalConfRequest() (request *DescribeMaGlobalConfRequest) {
	request = &DescribeMaGlobalConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaGlobalConf")
	return
}

func NewDescribeMaGlobalConfResponse() (response *DescribeMaGlobalConfResponse) {
	response = &DescribeMaGlobalConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警全局配置
func (c *Client) DescribeMaGlobalConf(request *DescribeMaGlobalConfRequest) (response *DescribeMaGlobalConfResponse, err error) {
	if request == nil {
		request = NewDescribeMaGlobalConfRequest()
	}
	response = NewDescribeMaGlobalConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStorageStatisticRequest() (request *DescribeLogStorageStatisticRequest) {
	request = &DescribeLogStorageStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageStatistic")
	return
}

func NewDescribeLogStorageStatisticResponse() (response *DescribeLogStorageStatisticResponse) {
	response = &DescribeLogStorageStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志检索容量使用统计
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
	if request == nil {
		request = NewDescribeLogStorageStatisticRequest()
	}
	response = NewDescribeLogStorageStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBaselineWeakPasswordRequest() (request *ModifyBaselineWeakPasswordRequest) {
	request = &ModifyBaselineWeakPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineWeakPassword")
	return
}

func NewModifyBaselineWeakPasswordResponse() (response *ModifyBaselineWeakPasswordResponse) {
	response = &ModifyBaselineWeakPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改或新增弱口令
func (c *Client) ModifyBaselineWeakPassword(request *ModifyBaselineWeakPasswordRequest) (response *ModifyBaselineWeakPasswordResponse, err error) {
	if request == nil {
		request = NewModifyBaselineWeakPasswordRequest()
	}
	response = NewModifyBaselineWeakPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineDefaultStrategyListRequest() (request *DescribeBaselineDefaultStrategyListRequest) {
	request = &DescribeBaselineDefaultStrategyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDefaultStrategyList")
	return
}

func NewDescribeBaselineDefaultStrategyListResponse() (response *DescribeBaselineDefaultStrategyListResponse) {
	response = &DescribeBaselineDefaultStrategyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询基线默认策略列表信息
func (c *Client) DescribeBaselineDefaultStrategyList(request *DescribeBaselineDefaultStrategyListRequest) (response *DescribeBaselineDefaultStrategyListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineDefaultStrategyListRequest()
	}
	response = NewDescribeBaselineDefaultStrategyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineDetectOverviewRequest() (request *DescribeBaselineDetectOverviewRequest) {
	request = &DescribeBaselineDetectOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetectOverview")
	return
}

func NewDescribeBaselineDetectOverviewResponse() (response *DescribeBaselineDetectOverviewResponse) {
	response = &DescribeBaselineDetectOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测概览
func (c *Client) DescribeBaselineDetectOverview(request *DescribeBaselineDetectOverviewRequest) (response *DescribeBaselineDetectOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineDetectOverviewRequest()
	}
	response = NewDescribeBaselineDetectOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineStrategyListRequest() (request *DescribeBaselineStrategyListRequest) {
	request = &DescribeBaselineStrategyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineStrategyList")
	return
}

func NewDescribeBaselineStrategyListResponse() (response *DescribeBaselineStrategyListResponse) {
	response = &DescribeBaselineStrategyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一个用户下的基线策略信息
func (c *Client) DescribeBaselineStrategyList(request *DescribeBaselineStrategyListRequest) (response *DescribeBaselineStrategyListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineStrategyListRequest()
	}
	response = NewDescribeBaselineStrategyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHotVulTopRequest() (request *DescribeHotVulTopRequest) {
	request = &DescribeHotVulTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeHotVulTop")
	return
}

func NewDescribeHotVulTopResponse() (response *DescribeHotVulTopResponse) {
	response = &DescribeHotVulTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全网热点漏洞
func (c *Client) DescribeHotVulTop(request *DescribeHotVulTopRequest) (response *DescribeHotVulTopResponse, err error) {
	if request == nil {
		request = NewDescribeHotVulTopRequest()
	}
	response = NewDescribeHotVulTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaWorkloadLabelRequest() (request *DescribeMaWorkloadLabelRequest) {
	request = &DescribeMaWorkloadLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaWorkloadLabel")
	return
}

func NewDescribeMaWorkloadLabelResponse() (response *DescribeMaWorkloadLabelResponse) {
	response = &DescribeMaWorkloadLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警负载标签
func (c *Client) DescribeMaWorkloadLabel(request *DescribeMaWorkloadLabelRequest) (response *DescribeMaWorkloadLabelResponse, err error) {
	if request == nil {
		request = NewDescribeMaWorkloadLabelRequest()
	}
	response = NewDescribeMaWorkloadLabelResponse()
	err = c.Send(request, response)
	return
}

func NewRetryVulFixRequest() (request *RetryVulFixRequest) {
	request = &RetryVulFixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RetryVulFix")
	return
}

func NewRetryVulFixResponse() (response *RetryVulFixResponse) {
	response = &RetryVulFixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修复失败时单独对某一个主机修复漏洞
func (c *Client) RetryVulFix(request *RetryVulFixRequest) (response *RetryVulFixResponse, err error) {
	if request == nil {
		request = NewRetryVulFixRequest()
	}
	response = NewRetryVulFixResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
	request = &DeleteMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwares")
	return
}

func NewDeleteMalwaresResponse() (response *DeleteMalwaresResponse) {
	response = &DeleteMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteMalwares) 用于删除木马记录。
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
	if request == nil {
		request = NewDeleteMalwaresRequest()
	}
	response = NewDeleteMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReverseShellSystemPolicyConfigRequest() (request *ModifyReverseShellSystemPolicyConfigRequest) {
	request = &ModifyReverseShellSystemPolicyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyReverseShellSystemPolicyConfig")
	return
}

func NewModifyReverseShellSystemPolicyConfigResponse() (response *ModifyReverseShellSystemPolicyConfigResponse) {
	response = &ModifyReverseShellSystemPolicyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改反弹shell系统策略配置
func (c *Client) ModifyReverseShellSystemPolicyConfig(request *ModifyReverseShellSystemPolicyConfigRequest) (response *ModifyReverseShellSystemPolicyConfigResponse, err error) {
	if request == nil {
		request = NewModifyReverseShellSystemPolicyConfigRequest()
	}
	response = NewModifyReverseShellSystemPolicyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLogKafkaConnectionStateRequest() (request *CheckLogKafkaConnectionStateRequest) {
	request = &CheckLogKafkaConnectionStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CheckLogKafkaConnectionState")
	return
}

func NewCheckLogKafkaConnectionStateResponse() (response *CheckLogKafkaConnectionStateResponse) {
	response = &CheckLogKafkaConnectionStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查日志投递kafka连通性
func (c *Client) CheckLogKafkaConnectionState(request *CheckLogKafkaConnectionStateRequest) (response *CheckLogKafkaConnectionStateResponse, err error) {
	if request == nil {
		request = NewCheckLogKafkaConnectionStateRequest()
	}
	response = NewCheckLogKafkaConnectionStateResponse()
	err = c.Send(request, response)
	return
}

func NewStopBaselineDetectRequest() (request *StopBaselineDetectRequest) {
	request = &StopBaselineDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "StopBaselineDetect")
	return
}

func NewStopBaselineDetectResponse() (response *StopBaselineDetectResponse) {
	response = &StopBaselineDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止基线检测
func (c *Client) StopBaselineDetect(request *StopBaselineDetectRequest) (response *StopBaselineDetectResponse, err error) {
	if request == nil {
		request = NewStopBaselineDetectRequest()
	}
	response = NewStopBaselineDetectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostTotalCountRequest() (request *DescribeAssetHostTotalCountRequest) {
	request = &DescribeAssetHostTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetHostTotalCount")
	return
}

func NewDescribeAssetHostTotalCountResponse() (response *DescribeAssetHostTotalCountResponse) {
	response = &DescribeAssetHostTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机所有资源数量
func (c *Client) DescribeAssetHostTotalCount(request *DescribeAssetHostTotalCountRequest) (response *DescribeAssetHostTotalCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostTotalCountRequest()
	}
	response = NewDescribeAssetHostTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulCountByDatesRequest() (request *DescribeVulCountByDatesRequest) {
	request = &DescribeVulCountByDatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulCountByDates")
	return
}

func NewDescribeVulCountByDatesResponse() (response *DescribeVulCountByDatesResponse) {
	response = &DescribeVulCountByDatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理模块，获取近日指定类型的漏洞数量和主机数量
func (c *Client) DescribeVulCountByDates(request *DescribeVulCountByDatesRequest) (response *DescribeVulCountByDatesResponse, err error) {
	if request == nil {
		request = NewDescribeVulCountByDatesRequest()
	}
	response = NewDescribeVulCountByDatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineInfoRequest() (request *DescribeAssetMachineInfoRequest) {
	request = &DescribeAssetMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineInfo")
	return
}

func NewDescribeAssetMachineInfoResponse() (response *DescribeAssetMachineInfoResponse) {
	response = &DescribeAssetMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机资源详情
func (c *Client) DescribeAssetMachineInfo(request *DescribeAssetMachineInfoRequest) (response *DescribeAssetMachineInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineInfoRequest()
	}
	response = NewDescribeAssetMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaConnEventTypeRequest() (request *ModifyMaConnEventTypeRequest) {
	request = &ModifyMaConnEventTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaConnEventType")
	return
}

func NewModifyMaConnEventTypeResponse() (response *ModifyMaConnEventTypeResponse) {
	response = &ModifyMaConnEventTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改微告警访问事件类型
func (c *Client) ModifyMaConnEventType(request *ModifyMaConnEventTypeRequest) (response *ModifyMaConnEventTypeResponse, err error) {
	if request == nil {
		request = NewModifyMaConnEventTypeRequest()
	}
	response = NewModifyMaConnEventTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebFrameListRequest() (request *DescribeAssetWebFrameListRequest) {
	request = &DescribeAssetWebFrameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebFrameList")
	return
}

func NewDescribeAssetWebFrameListResponse() (response *DescribeAssetWebFrameListResponse) {
	response = &DescribeAssetWebFrameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理Web框架列表
func (c *Client) DescribeAssetWebFrameList(request *DescribeAssetWebFrameListRequest) (response *DescribeAssetWebFrameListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebFrameListRequest()
	}
	response = NewDescribeAssetWebFrameListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetJarListRequest() (request *ExportAssetJarListRequest) {
	request = &ExportAssetJarListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetJarList")
	return
}

func NewExportAssetJarListResponse() (response *ExportAssetJarListResponse) {
	response = &ExportAssetJarListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出Jar包列表
func (c *Client) ExportAssetJarList(request *ExportAssetJarListRequest) (response *ExportAssetJarListResponse, err error) {
	if request == nil {
		request = NewExportAssetJarListRequest()
	}
	response = NewExportAssetJarListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaliciousRequestWhiteListRequest() (request *CreateMaliciousRequestWhiteListRequest) {
	request = &CreateMaliciousRequestWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateMaliciousRequestWhiteList")
	return
}

func NewCreateMaliciousRequestWhiteListResponse() (response *CreateMaliciousRequestWhiteListResponse) {
	response = &CreateMaliciousRequestWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加恶意请求白名单
func (c *Client) CreateMaliciousRequestWhiteList(request *CreateMaliciousRequestWhiteListRequest) (response *CreateMaliciousRequestWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateMaliciousRequestWhiteListRequest()
	}
	response = NewCreateMaliciousRequestWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetAttackWhiteListRequest() (request *CreateNetAttackWhiteListRequest) {
	request = &CreateNetAttackWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateNetAttackWhiteList")
	return
}

func NewCreateNetAttackWhiteListResponse() (response *CreateNetAttackWhiteListResponse) {
	response = &CreateNetAttackWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) CreateNetAttackWhiteList(request *CreateNetAttackWhiteListRequest) (response *CreateNetAttackWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateNetAttackWhiteListRequest()
	}
	response = NewCreateNetAttackWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemDetectListRequest() (request *DescribeBaselineItemDetectListRequest) {
	request = &DescribeBaselineItemDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemDetectList")
	return
}

func NewDescribeBaselineItemDetectListResponse() (response *DescribeBaselineItemDetectListResponse) {
	response = &DescribeBaselineItemDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测项的列表
func (c *Client) DescribeBaselineItemDetectList(request *DescribeBaselineItemDetectListRequest) (response *DescribeBaselineItemDetectListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemDetectListRequest()
	}
	response = NewDescribeBaselineItemDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningHostConfigRequest() (request *DescribeWarningHostConfigRequest) {
	request = &DescribeWarningHostConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWarningHostConfig")
	return
}

func NewDescribeWarningHostConfigResponse() (response *DescribeWarningHostConfigResponse) {
	response = &DescribeWarningHostConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警机器范围配置
func (c *Client) DescribeWarningHostConfig(request *DescribeWarningHostConfigRequest) (response *DescribeWarningHostConfigResponse, err error) {
	if request == nil {
		request = NewDescribeWarningHostConfigRequest()
	}
	response = NewDescribeWarningHostConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportJavaMemShellsRequest() (request *ExportJavaMemShellsRequest) {
	request = &ExportJavaMemShellsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportJavaMemShells")
	return
}

func NewExportJavaMemShellsResponse() (response *ExportJavaMemShellsResponse) {
	response = &ExportJavaMemShellsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出java内存马事件列表
func (c *Client) ExportJavaMemShells(request *ExportJavaMemShellsRequest) (response *ExportJavaMemShellsResponse, err error) {
	if request == nil {
		request = NewExportJavaMemShellsRequest()
	}
	response = NewExportJavaMemShellsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackTopRequest() (request *DescribeAttackTopRequest) {
	request = &DescribeAttackTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackTop")
	return
}

func NewDescribeAttackTopResponse() (response *DescribeAttackTopResponse) {
	response = &DescribeAttackTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeAttackTop(request *DescribeAttackTopRequest) (response *DescribeAttackTopResponse, err error) {
	if request == nil {
		request = NewDescribeAttackTopRequest()
	}
	response = NewDescribeAttackTopResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetCoreModuleListRequest() (request *ExportAssetCoreModuleListRequest) {
	request = &ExportAssetCoreModuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetCoreModuleList")
	return
}

func NewExportAssetCoreModuleListResponse() (response *ExportAssetCoreModuleListResponse) {
	response = &ExportAssetCoreModuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理内核模块列表
func (c *Client) ExportAssetCoreModuleList(request *ExportAssetCoreModuleListRequest) (response *ExportAssetCoreModuleListResponse, err error) {
	if request == nil {
		request = NewExportAssetCoreModuleListRequest()
	}
	response = NewExportAssetCoreModuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNonlocalLoginPlacesRequest() (request *DeleteNonlocalLoginPlacesRequest) {
	request = &DeleteNonlocalLoginPlacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteNonlocalLoginPlaces")
	return
}

func NewDeleteNonlocalLoginPlacesResponse() (response *DeleteNonlocalLoginPlacesResponse) {
	response = &DeleteNonlocalLoginPlacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteNonlocalLoginPlaces) 用于删除异地登录记录。
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
	if request == nil {
		request = NewDeleteNonlocalLoginPlacesRequest()
	}
	response = NewDeleteNonlocalLoginPlacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceInfoListRequest() (request *DescribeAssetWebServiceInfoListRequest) {
	request = &DescribeAssetWebServiceInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceInfoList")
	return
}

func NewDescribeAssetWebServiceInfoListResponse() (response *DescribeAssetWebServiceInfoListResponse) {
	response = &DescribeAssetWebServiceInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理Web服务列表
func (c *Client) DescribeAssetWebServiceInfoList(request *DescribeAssetWebServiceInfoListRequest) (response *DescribeAssetWebServiceInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceInfoListRequest()
	}
	response = NewDescribeAssetWebServiceInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityBroadcastInfoRequest() (request *DescribeSecurityBroadcastInfoRequest) {
	request = &DescribeSecurityBroadcastInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityBroadcastInfo")
	return
}

func NewDescribeSecurityBroadcastInfoResponse() (response *DescribeSecurityBroadcastInfoResponse) {
	response = &DescribeSecurityBroadcastInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全播报文章信息
func (c *Client) DescribeSecurityBroadcastInfo(request *DescribeSecurityBroadcastInfoRequest) (response *DescribeSecurityBroadcastInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityBroadcastInfoRequest()
	}
	response = NewDescribeSecurityBroadcastInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseBackupListRequest() (request *DescribeRansomDefenseBackupListRequest) {
	request = &DescribeRansomDefenseBackupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseBackupList")
	return
}

func NewDescribeRansomDefenseBackupListResponse() (response *DescribeRansomDefenseBackupListResponse) {
	response = &DescribeRansomDefenseBackupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机快照备份列表
func (c *Client) DescribeRansomDefenseBackupList(request *DescribeRansomDefenseBackupListRequest) (response *DescribeRansomDefenseBackupListResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseBackupListRequest()
	}
	response = NewDescribeRansomDefenseBackupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseBindListRequest() (request *DescribeLicenseBindListRequest) {
	request = &DescribeLicenseBindListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseBindList")
	return
}

func NewDescribeLicenseBindListResponse() (response *DescribeLicenseBindListResponse) {
	response = &DescribeLicenseBindListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口可以获取设置中心-授权管理,某个授权下已绑定的授权机器列表
func (c *Client) DescribeLicenseBindList(request *DescribeLicenseBindListRequest) (response *DescribeLicenseBindListResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseBindListRequest()
	}
	response = NewDescribeLicenseBindListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProVersionStatusRequest() (request *DescribeProVersionStatusRequest) {
	request = &DescribeProVersionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProVersionStatus")
	return
}

func NewDescribeProVersionStatusResponse() (response *DescribeProVersionStatusResponse) {
	response = &DescribeProVersionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取单台主机或所有主机是否开通专业版状态。
func (c *Client) DescribeProVersionStatus(request *DescribeProVersionStatusRequest) (response *DescribeProVersionStatusResponse, err error) {
	if request == nil {
		request = NewDescribeProVersionStatusRequest()
	}
	response = NewDescribeProVersionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityProtectionStatRequest() (request *DescribeSecurityProtectionStatRequest) {
	request = &DescribeSecurityProtectionStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityProtectionStat")
	return
}

func NewDescribeSecurityProtectionStatResponse() (response *DescribeSecurityProtectionStatResponse) {
	response = &DescribeSecurityProtectionStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全防护状态汇总
func (c *Client) DescribeSecurityProtectionStat(request *DescribeSecurityProtectionStatRequest) (response *DescribeSecurityProtectionStatResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityProtectionStatRequest()
	}
	response = NewDescribeSecurityProtectionStatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJavaMemShellPluginSwitchRequest() (request *ModifyJavaMemShellPluginSwitchRequest) {
	request = &ModifyJavaMemShellPluginSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyJavaMemShellPluginSwitch")
	return
}

func NewModifyJavaMemShellPluginSwitchResponse() (response *ModifyJavaMemShellPluginSwitchResponse) {
	response = &ModifyJavaMemShellPluginSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关java内存马插件
func (c *Client) ModifyJavaMemShellPluginSwitch(request *ModifyJavaMemShellPluginSwitchRequest) (response *ModifyJavaMemShellPluginSwitchResponse, err error) {
	if request == nil {
		request = NewModifyJavaMemShellPluginSwitchRequest()
	}
	response = NewModifyJavaMemShellPluginSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsPolicyListRequest() (request *DescribeRiskDnsPolicyListRequest) {
	request = &DescribeRiskDnsPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsPolicyList")
	return
}

func NewDescribeRiskDnsPolicyListResponse() (response *DescribeRiskDnsPolicyListResponse) {
	response = &DescribeRiskDnsPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取恶意请求策略列表
func (c *Client) DescribeRiskDnsPolicyList(request *DescribeRiskDnsPolicyListRequest) (response *DescribeRiskDnsPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsPolicyListRequest()
	}
	response = NewDescribeRiskDnsPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
	request = &DescribeSecurityDynamicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityDynamics")
	return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
	response = &DescribeSecurityDynamicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeSecurityDynamics) 用于获取安全事件动态消息数据。
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityDynamicsRequest()
	}
	response = NewDescribeSecurityDynamicsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBaselinePolicyStateRequest() (request *ModifyBaselinePolicyStateRequest) {
	request = &ModifyBaselinePolicyStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselinePolicyState")
	return
}

func NewModifyBaselinePolicyStateResponse() (response *ModifyBaselinePolicyStateResponse) {
	response = &ModifyBaselinePolicyStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改基线策略状态
func (c *Client) ModifyBaselinePolicyState(request *ModifyBaselinePolicyStateRequest) (response *ModifyBaselinePolicyStateResponse, err error) {
	if request == nil {
		request = NewModifyBaselinePolicyStateRequest()
	}
	response = NewModifyBaselinePolicyStateResponse()
	err = c.Send(request, response)
	return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
	request = &OpenProVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "OpenProVersion")
	return
}

func NewOpenProVersionResponse() (response *OpenProVersionResponse) {
	response = &OpenProVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (OpenProVersion) 用于开通专业版。
func (c *Client) OpenProVersion(request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
	if request == nil {
		request = NewOpenProVersionRequest()
	}
	response = NewOpenProVersionResponse()
	err = c.Send(request, response)
	return
}

func NewScanVulSettingRequest() (request *ScanVulSettingRequest) {
	request = &ScanVulSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanVulSetting")
	return
}

func NewScanVulSettingResponse() (response *ScanVulSettingResponse) {
	response = &ScanVulSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 定期扫描漏洞设置
func (c *Client) ScanVulSetting(request *ScanVulSettingRequest) (response *ScanVulSettingResponse, err error) {
	if request == nil {
		request = NewScanVulSettingRequest()
	}
	response = NewScanVulSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselinePolicyListRequest() (request *DescribeBaselinePolicyListRequest) {
	request = &DescribeBaselinePolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselinePolicyList")
	return
}

func NewDescribeBaselinePolicyListResponse() (response *DescribeBaselinePolicyListResponse) {
	response = &DescribeBaselinePolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线策略列表
func (c *Client) DescribeBaselinePolicyList(request *DescribeBaselinePolicyListRequest) (response *DescribeBaselinePolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselinePolicyListRequest()
	}
	response = NewDescribeBaselinePolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsEventListRequest() (request *DescribeRiskDnsEventListRequest) {
	request = &DescribeRiskDnsEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsEventList")
	return
}

func NewDescribeRiskDnsEventListResponse() (response *DescribeRiskDnsEventListResponse) {
	response = &DescribeRiskDnsEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取恶意请求事件列表
func (c *Client) DescribeRiskDnsEventList(request *DescribeRiskDnsEventListRequest) (response *DescribeRiskDnsEventListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsEventListRequest()
	}
	response = NewDescribeRiskDnsEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShellPolicyListRequest() (request *DescribeShellPolicyListRequest) {
	request = &DescribeShellPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeShellPolicyList")
	return
}

func NewDescribeShellPolicyListResponse() (response *DescribeShellPolicyListResponse) {
	response = &DescribeShellPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询反弹shell策略列表
func (c *Client) DescribeShellPolicyList(request *DescribeShellPolicyListRequest) (response *DescribeShellPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeShellPolicyListRequest()
	}
	response = NewDescribeShellPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewEditReverseShellRuleRequest() (request *EditReverseShellRuleRequest) {
	request = &EditReverseShellRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "EditReverseShellRule")
	return
}

func NewEditReverseShellRuleResponse() (response *EditReverseShellRuleResponse) {
	response = &EditReverseShellRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑反弹Shell规则
func (c *Client) EditReverseShellRule(request *EditReverseShellRuleRequest) (response *EditReverseShellRuleResponse, err error) {
	if request == nil {
		request = NewEditReverseShellRuleRequest()
	}
	response = NewEditReverseShellRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBaselinePolicyRequest() (request *ModifyBaselinePolicyRequest) {
	request = &ModifyBaselinePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselinePolicy")
	return
}

func NewModifyBaselinePolicyResponse() (response *ModifyBaselinePolicyResponse) {
	response = &ModifyBaselinePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改基线策略设置
func (c *Client) ModifyBaselinePolicy(request *ModifyBaselinePolicyRequest) (response *ModifyBaselinePolicyResponse, err error) {
	if request == nil {
		request = NewModifyBaselinePolicyRequest()
	}
	response = NewModifyBaselinePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaliciousRequestsRequest() (request *DeleteMaliciousRequestsRequest) {
	request = &DeleteMaliciousRequestsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaliciousRequests")
	return
}

func NewDeleteMaliciousRequestsResponse() (response *DeleteMaliciousRequestsResponse) {
	response = &DeleteMaliciousRequestsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteMaliciousRequests) 用于删除恶意请求记录。
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
	if request == nil {
		request = NewDeleteMaliciousRequestsRequest()
	}
	response = NewDeleteMaliciousRequestsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseInfoRequest() (request *DescribeAssetDatabaseInfoRequest) {
	request = &DescribeAssetDatabaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseInfo")
	return
}

func NewDescribeAssetDatabaseInfoResponse() (response *DescribeAssetDatabaseInfoResponse) {
	response = &DescribeAssetDatabaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理数据库详情
func (c *Client) DescribeAssetDatabaseInfo(request *DescribeAssetDatabaseInfoRequest) (response *DescribeAssetDatabaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseInfoRequest()
	}
	response = NewDescribeAssetDatabaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseCountRequest() (request *DescribeAssetDatabaseCountRequest) {
	request = &DescribeAssetDatabaseCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseCount")
	return
}

func NewDescribeAssetDatabaseCountResponse() (response *DescribeAssetDatabaseCountResponse) {
	response = &DescribeAssetDatabaseCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有数据库数量
func (c *Client) DescribeAssetDatabaseCount(request *DescribeAssetDatabaseCountRequest) (response *DescribeAssetDatabaseCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseCountRequest()
	}
	response = NewDescribeAssetDatabaseCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBanWhiteListRequest() (request *ModifyBanWhiteListRequest) {
	request = &ModifyBanWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanWhiteList")
	return
}

func NewModifyBanWhiteListResponse() (response *ModifyBanWhiteListResponse) {
	response = &ModifyBanWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改阻断白名单列表
func (c *Client) ModifyBanWhiteList(request *ModifyBanWhiteListRequest) (response *ModifyBanWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyBanWhiteListRequest()
	}
	response = NewModifyBanWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackLogInfoRequest() (request *DescribeAttackLogInfoRequest) {
	request = &DescribeAttackLogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackLogInfo")
	return
}

func NewDescribeAttackLogInfoResponse() (response *DescribeAttackLogInfoResponse) {
	response = &DescribeAttackLogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网络攻击日志详情
func (c *Client) DescribeAttackLogInfo(request *DescribeAttackLogInfoRequest) (response *DescribeAttackLogInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAttackLogInfoRequest()
	}
	response = NewDescribeAttackLogInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineWeakPasswordListRequest() (request *ExportBaselineWeakPasswordListRequest) {
	request = &ExportBaselineWeakPasswordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineWeakPasswordList")
	return
}

func NewExportBaselineWeakPasswordListResponse() (response *ExportBaselineWeakPasswordListResponse) {
	response = &ExportBaselineWeakPasswordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出弱口令配置列表
func (c *Client) ExportBaselineWeakPasswordList(request *ExportBaselineWeakPasswordListRequest) (response *ExportBaselineWeakPasswordListResponse, err error) {
	if request == nil {
		request = NewExportBaselineWeakPasswordListRequest()
	}
	response = NewExportBaselineWeakPasswordListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
	request = &DeleteReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteReverseShellEvents")
	return
}

func NewDeleteReverseShellEventsResponse() (response *DeleteReverseShellEventsResponse) {
	response = &DeleteReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Ids删除反弹Shell事件
func (c *Client) DeleteReverseShellEvents(request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellEventsRequest()
	}
	response = NewDeleteReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBanRegionsRequest() (request *DescribeBanRegionsRequest) {
	request = &DescribeBanRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanRegions")
	return
}

func NewDescribeBanRegionsResponse() (response *DescribeBanRegionsResponse) {
	response = &DescribeBanRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取阻断地域
func (c *Client) DescribeBanRegions(request *DescribeBanRegionsRequest) (response *DescribeBanRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeBanRegionsRequest()
	}
	response = NewDescribeBanRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTrendRequest() (request *DescribeVulTrendRequest) {
	request = &DescribeVulTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulTrend")
	return
}

func NewDescribeVulTrendResponse() (response *DescribeVulTrendResponse) {
	response = &DescribeVulTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞态势信息
func (c *Client) DescribeVulTrend(request *DescribeVulTrendRequest) (response *DescribeVulTrendResponse, err error) {
	if request == nil {
		request = NewDescribeVulTrendRequest()
	}
	response = NewDescribeVulTrendResponse()
	err = c.Send(request, response)
	return
}

func NewExportPrivilegeEventsRequest() (request *ExportPrivilegeEventsRequest) {
	request = &ExportPrivilegeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportPrivilegeEvents")
	return
}

func NewExportPrivilegeEventsResponse() (response *ExportPrivilegeEventsResponse) {
	response = &ExportPrivilegeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出本地提权事件
func (c *Client) ExportPrivilegeEvents(request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
	if request == nil {
		request = NewExportPrivilegeEventsRequest()
	}
	response = NewExportPrivilegeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
	request = &SeparateMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SeparateMalwares")
	return
}

func NewSeparateMalwaresResponse() (response *SeparateMalwaresResponse) {
	response = &SeparateMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SeparateMalwares）用于隔离木马。
func (c *Client) SeparateMalwares(request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
	if request == nil {
		request = NewSeparateMalwaresRequest()
	}
	response = NewSeparateMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
	request = &DescribeTagMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeTagMachines")
	return
}

func NewDescribeTagMachinesResponse() (response *DescribeTagMachinesResponse) {
	response = &DescribeTagMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定标签关联的服务器信息
func (c *Client) DescribeTagMachines(request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeTagMachinesRequest()
	}
	response = NewDescribeTagMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaLabelListRequest() (request *DescribeMaLabelListRequest) {
	request = &DescribeMaLabelListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaLabelList")
	return
}

func NewDescribeMaLabelListResponse() (response *DescribeMaLabelListResponse) {
	response = &DescribeMaLabelListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警标签列表
func (c *Client) DescribeMaLabelList(request *DescribeMaLabelListRequest) (response *DescribeMaLabelListResponse, err error) {
	if request == nil {
		request = NewDescribeMaLabelListRequest()
	}
	response = NewDescribeMaLabelListResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageExpireRequest() (request *SetLocalStorageExpireRequest) {
	request = &SetLocalStorageExpireRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SetLocalStorageExpire")
	return
}

func NewSetLocalStorageExpireResponse() (response *SetLocalStorageExpireResponse) {
	response = &SetLocalStorageExpireResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置本地存储过期时间
func (c *Client) SetLocalStorageExpire(request *SetLocalStorageExpireRequest) (response *SetLocalStorageExpireResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageExpireRequest()
	}
	response = NewSetLocalStorageExpireResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBashRuleParamsRequest() (request *CheckBashRuleParamsRequest) {
	request = &CheckBashRuleParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CheckBashRuleParams")
	return
}

func NewCheckBashRuleParamsResponse() (response *CheckBashRuleParamsResponse) {
	response = &CheckBashRuleParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验高危命令用户规则新增和编辑时的参数。
func (c *Client) CheckBashRuleParams(request *CheckBashRuleParamsRequest) (response *CheckBashRuleParamsResponse, err error) {
	if request == nil {
		request = NewCheckBashRuleParamsRequest()
	}
	response = NewCheckBashRuleParamsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogExportRequest() (request *DeleteLogExportRequest) {
	request = &DeleteLogExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteLogExport")
	return
}

func NewDeleteLogExportResponse() (response *DeleteLogExportResponse) {
	response = &DeleteLogExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志下载任务
func (c *Client) DeleteLogExport(request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
	if request == nil {
		request = NewDeleteLogExportRequest()
	}
	response = NewDeleteLogExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESHitsRequest() (request *DescribeESHitsRequest) {
	request = &DescribeESHitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeESHits")
	return
}

func NewDescribeESHitsResponse() (response *DescribeESHitsResponse) {
	response = &DescribeESHitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES查询文档列表
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
	if request == nil {
		request = NewDescribeESHitsRequest()
	}
	response = NewDescribeESHitsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBaselineStrategyRequest() (request *UpdateBaselineStrategyRequest) {
	request = &UpdateBaselineStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "UpdateBaselineStrategy")
	return
}

func NewUpdateBaselineStrategyResponse() (response *UpdateBaselineStrategyResponse) {
	response = &UpdateBaselineStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线策略id更新策略信息
func (c *Client) UpdateBaselineStrategy(request *UpdateBaselineStrategyRequest) (response *UpdateBaselineStrategyResponse, err error) {
	if request == nil {
		request = NewUpdateBaselineStrategyRequest()
	}
	response = NewUpdateBaselineStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSystemPackageListRequest() (request *DescribeAssetSystemPackageListRequest) {
	request = &DescribeAssetSystemPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetSystemPackageList")
	return
}

func NewDescribeAssetSystemPackageListResponse() (response *DescribeAssetSystemPackageListResponse) {
	response = &DescribeAssetSystemPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理系统安装包列表
func (c *Client) DescribeAssetSystemPackageList(request *DescribeAssetSystemPackageListRequest) (response *DescribeAssetSystemPackageListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSystemPackageListRequest()
	}
	response = NewDescribeAssetSystemPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
	request = &DescribeHistoryAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeHistoryAccounts")
	return
}

func NewDescribeHistoryAccountsResponse() (response *DescribeHistoryAccountsResponse) {
	response = &DescribeHistoryAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeHistoryAccounts) 用于获取帐号变更历史列表数据。
func (c *Client) DescribeHistoryAccounts(request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeHistoryAccountsRequest()
	}
	response = NewDescribeHistoryAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefenceEventRequest() (request *ExportVulDefenceEventRequest) {
	request = &ExportVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefenceEvent")
	return
}

func NewExportVulDefenceEventResponse() (response *ExportVulDefenceEventResponse) {
	response = &ExportVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞防御事件
func (c *Client) ExportVulDefenceEvent(request *ExportVulDefenceEventRequest) (response *ExportVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewExportVulDefenceEventRequest()
	}
	response = NewExportVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppCountRequest() (request *DescribeAssetAppCountRequest) {
	request = &DescribeAssetAppCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppCount")
	return
}

func NewDescribeAssetAppCountResponse() (response *DescribeAssetAppCountResponse) {
	response = &DescribeAssetAppCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有软件应用数量
func (c *Client) DescribeAssetAppCount(request *DescribeAssetAppCountRequest) (response *DescribeAssetAppCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppCountRequest()
	}
	response = NewDescribeAssetAppCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetLoadInfoRequest() (request *DescribeAssetLoadInfoRequest) {
	request = &DescribeAssetLoadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetLoadInfo")
	return
}

func NewDescribeAssetLoadInfoResponse() (response *DescribeAssetLoadInfoResponse) {
	response = &DescribeAssetLoadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统负载、内存使用率、硬盘使用率情况
func (c *Client) DescribeAssetLoadInfo(request *DescribeAssetLoadInfoRequest) (response *DescribeAssetLoadInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetLoadInfoRequest()
	}
	response = NewDescribeAssetLoadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseStrategyMachinesRequest() (request *DescribeRansomDefenseStrategyMachinesRequest) {
	request = &DescribeRansomDefenseStrategyMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyMachines")
	return
}

func NewDescribeRansomDefenseStrategyMachinesResponse() (response *DescribeRansomDefenseStrategyMachinesResponse) {
	response = &DescribeRansomDefenseStrategyMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询防勒索策略绑定机器列表
func (c *Client) DescribeRansomDefenseStrategyMachines(request *DescribeRansomDefenseStrategyMachinesRequest) (response *DescribeRansomDefenseStrategyMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseStrategyMachinesRequest()
	}
	response = NewDescribeRansomDefenseStrategyMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInstallTaskResultRequest() (request *DescribeAgentInstallTaskResultRequest) {
	request = &DescribeAgentInstallTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallTaskResult")
	return
}

func NewDescribeAgentInstallTaskResultResponse() (response *DescribeAgentInstallTaskResultResponse) {
	response = &DescribeAgentInstallTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ageng安装任务完成情况
func (c *Client) DescribeAgentInstallTaskResult(request *DescribeAgentInstallTaskResultRequest) (response *DescribeAgentInstallTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInstallTaskResultRequest()
	}
	response = NewDescribeAgentInstallTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenProtectionCntRequest() (request *DescribeScreenProtectionCntRequest) {
	request = &DescribeScreenProtectionCntRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenProtectionCnt")
	return
}

func NewDescribeScreenProtectionCntResponse() (response *DescribeScreenProtectionCntResponse) {
	response = &DescribeScreenProtectionCntResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化主机安全防护引擎介绍
func (c *Client) DescribeScreenProtectionCnt(request *DescribeScreenProtectionCntRequest) (response *DescribeScreenProtectionCntResponse, err error) {
	if request == nil {
		request = NewDescribeScreenProtectionCntRequest()
	}
	response = NewDescribeScreenProtectionCntResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulListRequest() (request *ExportVulListRequest) {
	request = &ExportVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulList")
	return
}

func NewExportVulListResponse() (response *ExportVulListResponse) {
	response = &ExportVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-导出漏洞列表
func (c *Client) ExportVulList(request *ExportVulListRequest) (response *ExportVulListResponse, err error) {
	if request == nil {
		request = NewExportVulListRequest()
	}
	response = NewExportVulListResponse()
	err = c.Send(request, response)
	return
}

func NewSetLocalStorageItemRequest() (request *SetLocalStorageItemRequest) {
	request = &SetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SetLocalStorageItem")
	return
}

func NewSetLocalStorageItemResponse() (response *SetLocalStorageItemResponse) {
	response = &SetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置本地存储数据
func (c *Client) SetLocalStorageItem(request *SetLocalStorageItemRequest) (response *SetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewSetLocalStorageItemRequest()
	}
	response = NewSetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
	request = &IgnoreImpactedHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "IgnoreImpactedHosts")
	return
}

func NewIgnoreImpactedHostsResponse() (response *IgnoreImpactedHostsResponse) {
	response = &IgnoreImpactedHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (IgnoreImpactedHosts) 用于忽略漏洞。
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
	if request == nil {
		request = NewIgnoreImpactedHostsRequest()
	}
	response = NewIgnoreImpactedHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperRuleInfoRequest() (request *DescribeFileTamperRuleInfoRequest) {
	request = &DescribeFileTamperRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRuleInfo")
	return
}

func NewDescribeFileTamperRuleInfoResponse() (response *DescribeFileTamperRuleInfoResponse) {
	response = &DescribeFileTamperRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个监控规则的详情
func (c *Client) DescribeFileTamperRuleInfo(request *DescribeFileTamperRuleInfoRequest) (response *DescribeFileTamperRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperRuleInfoRequest()
	}
	response = NewDescribeFileTamperRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeTags")
	return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有主机标签
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	if request == nil {
		request = NewDescribeTagsRequest()
	}
	response = NewDescribeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewExportTasksRequest() (request *ExportTasksRequest) {
	request = &ExportTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportTasks")
	return
}

func NewExportTasksResponse() (response *ExportTasksResponse) {
	response = &ExportTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于异步导出数据量大的日志文件
func (c *Client) ExportTasks(request *ExportTasksRequest) (response *ExportTasksResponse, err error) {
	if request == nil {
		request = NewExportTasksRequest()
	}
	response = NewExportTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchExportListRequest() (request *DescribeSearchExportListRequest) {
	request = &DescribeSearchExportListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchExportList")
	return
}

func NewDescribeSearchExportListResponse() (response *DescribeSearchExportListResponse) {
	response = &DescribeSearchExportListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出ES查询文档列表
func (c *Client) DescribeSearchExportList(request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
	if request == nil {
		request = NewDescribeSearchExportListRequest()
	}
	response = NewDescribeSearchExportListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTypesRequest() (request *DescribeAssetTypesRequest) {
	request = &DescribeAssetTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTypes")
	return
}

func NewDescribeAssetTypesResponse() (response *DescribeAssetTypesResponse) {
	response = &DescribeAssetTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产指纹类型列表
func (c *Client) DescribeAssetTypes(request *DescribeAssetTypesRequest) (response *DescribeAssetTypesResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTypesRequest()
	}
	response = NewDescribeAssetTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenEventsCntRequest() (request *DescribeScreenEventsCntRequest) {
	request = &DescribeScreenEventsCntRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenEventsCnt")
	return
}

func NewDescribeScreenEventsCntResponse() (response *DescribeScreenEventsCntResponse) {
	response = &DescribeScreenEventsCntResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化获取安全概览相关事件统计数据接口
func (c *Client) DescribeScreenEventsCnt(request *DescribeScreenEventsCntRequest) (response *DescribeScreenEventsCntResponse, err error) {
	if request == nil {
		request = NewDescribeScreenEventsCntRequest()
	}
	response = NewDescribeScreenEventsCntResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebHookReceiverRequest() (request *ModifyWebHookReceiverRequest) {
	request = &ModifyWebHookReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookReceiver")
	return
}

func NewModifyWebHookReceiverResponse() (response *ModifyWebHookReceiverResponse) {
	response = &ModifyWebHookReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或更新告警接收人
func (c *Client) ModifyWebHookReceiver(request *ModifyWebHookReceiverRequest) (response *ModifyWebHookReceiverResponse, err error) {
	if request == nil {
		request = NewModifyWebHookReceiverRequest()
	}
	response = NewModifyWebHookReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewScanVulAgainRequest() (request *ScanVulAgainRequest) {
	request = &ScanVulAgainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanVulAgain")
	return
}

func NewScanVulAgainResponse() (response *ScanVulAgainResponse) {
	response = &ScanVulAgainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-重新检测接口
func (c *Client) ScanVulAgain(request *ScanVulAgainRequest) (response *ScanVulAgainResponse, err error) {
	if request == nil {
		request = NewScanVulAgainRequest()
	}
	response = NewScanVulAgainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetCoreModuleListRequest() (request *DescribeAssetCoreModuleListRequest) {
	request = &DescribeAssetCoreModuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetCoreModuleList")
	return
}

func NewDescribeAssetCoreModuleListResponse() (response *DescribeAssetCoreModuleListResponse) {
	response = &DescribeAssetCoreModuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理内核模块列表
func (c *Client) DescribeAssetCoreModuleList(request *DescribeAssetCoreModuleListRequest) (response *DescribeAssetCoreModuleListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetCoreModuleListRequest()
	}
	response = NewDescribeAssetCoreModuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserKeyListRequest() (request *DescribeAssetUserKeyListRequest) {
	request = &DescribeAssetUserKeyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserKeyList")
	return
}

func NewDescribeAssetUserKeyListResponse() (response *DescribeAssetUserKeyListResponse) {
	response = &DescribeAssetUserKeyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机账号Key列表
func (c *Client) DescribeAssetUserKeyList(request *DescribeAssetUserKeyListRequest) (response *DescribeAssetUserKeyListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserKeyListRequest()
	}
	response = NewDescribeAssetUserKeyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackSourceEventsRequest() (request *DescribeAttackSourceEventsRequest) {
	request = &DescribeAttackSourceEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackSourceEvents")
	return
}

func NewDescribeAttackSourceEventsResponse() (response *DescribeAttackSourceEventsResponse) {
	response = &DescribeAttackSourceEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询攻击溯源事件
func (c *Client) DescribeAttackSourceEvents(request *DescribeAttackSourceEventsRequest) (response *DescribeAttackSourceEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackSourceEventsRequest()
	}
	response = NewDescribeAttackSourceEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportProtectDirListRequest() (request *ExportProtectDirListRequest) {
	request = &ExportProtectDirListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportProtectDirList")
	return
}

func NewExportProtectDirListResponse() (response *ExportProtectDirListResponse) {
	response = &ExportProtectDirListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出网页防篡改防护目录列表
func (c *Client) ExportProtectDirList(request *ExportProtectDirListRequest) (response *ExportProtectDirListResponse, err error) {
	if request == nil {
		request = NewExportProtectDirListRequest()
	}
	response = NewExportProtectDirListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebHookReceiverRequest() (request *DeleteWebHookReceiverRequest) {
	request = &DeleteWebHookReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookReceiver")
	return
}

func NewDeleteWebHookReceiverResponse() (response *DeleteWebHookReceiverResponse) {
	response = &DeleteWebHookReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警接收人
func (c *Client) DeleteWebHookReceiver(request *DeleteWebHookReceiverRequest) (response *DeleteWebHookReceiverResponse, err error) {
	if request == nil {
		request = NewDeleteWebHookReceiverRequest()
	}
	response = NewDeleteWebHookReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogIndexRequest() (request *DescribeLogIndexRequest) {
	request = &DescribeLogIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogIndex")
	return
}

func NewDescribeLogIndexResponse() (response *DescribeLogIndexResponse) {
	response = &DescribeLogIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询索引
func (c *Client) DescribeLogIndex(request *DescribeLogIndexRequest) (response *DescribeLogIndexResponse, err error) {
	if request == nil {
		request = NewDescribeLogIndexRequest()
	}
	response = NewDescribeLogIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineFileTamperRulesRequest() (request *DescribeMachineFileTamperRulesRequest) {
	request = &DescribeMachineFileTamperRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineFileTamperRules")
	return
}

func NewDescribeMachineFileTamperRulesResponse() (response *DescribeMachineFileTamperRulesResponse) {
	response = &DescribeMachineFileTamperRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机相关核心文件监控规则列 表
func (c *Client) DescribeMachineFileTamperRules(request *DescribeMachineFileTamperRulesRequest) (response *DescribeMachineFileTamperRulesResponse, err error) {
	if request == nil {
		request = NewDescribeMachineFileTamperRulesRequest()
	}
	response = NewDescribeMachineFileTamperRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareTimingScanSettingRequest() (request *DescribeMalwareTimingScanSettingRequest) {
	request = &DescribeMalwareTimingScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareTimingScanSetting")
	return
}

func NewDescribeMalwareTimingScanSettingResponse() (response *DescribeMalwareTimingScanSettingResponse) {
	response = &DescribeMalwareTimingScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询定时扫描配置
func (c *Client) DescribeMalwareTimingScanSetting(request *DescribeMalwareTimingScanSettingRequest) (response *DescribeMalwareTimingScanSettingResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareTimingScanSettingRequest()
	}
	response = NewDescribeMalwareTimingScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenMachineRegionsRequest() (request *DescribeScreenMachineRegionsRequest) {
	request = &DescribeScreenMachineRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenMachineRegions")
	return
}

func NewDescribeScreenMachineRegionsResponse() (response *DescribeScreenMachineRegionsResponse) {
	response = &DescribeScreenMachineRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化主机区域选项列表
func (c *Client) DescribeScreenMachineRegions(request *DescribeScreenMachineRegionsRequest) (response *DescribeScreenMachineRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeScreenMachineRegionsRequest()
	}
	response = NewDescribeScreenMachineRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectDirListRequest() (request *DescribeProtectDirListRequest) {
	request = &DescribeProtectDirListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProtectDirList")
	return
}

func NewDescribeProtectDirListResponse() (response *DescribeProtectDirListResponse) {
	response = &DescribeProtectDirListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡改防护目录列表
func (c *Client) DescribeProtectDirList(request *DescribeProtectDirListRequest) (response *DescribeProtectDirListResponse, err error) {
	if request == nil {
		request = NewDescribeProtectDirListRequest()
	}
	response = NewDescribeProtectDirListResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
	request = &ExportMaliciousRequestsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaliciousRequests")
	return
}

func NewExportMaliciousRequestsResponse() (response *ExportMaliciousRequestsResponse) {
	response = &ExportMaliciousRequestsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ExportMaliciousRequests) 用于导出下载恶意请求文件。
func (c *Client) ExportMaliciousRequests(request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
	if request == nil {
		request = NewExportMaliciousRequestsRequest()
	}
	response = NewExportMaliciousRequestsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebPageGeneralizeRequest() (request *DescribeWebPageGeneralizeRequest) {
	request = &DescribeWebPageGeneralizeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageGeneralize")
	return
}

func NewDescribeWebPageGeneralizeResponse() (response *DescribeWebPageGeneralizeResponse) {
	response = &DescribeWebPageGeneralizeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网站防篡改概览信息
func (c *Client) DescribeWebPageGeneralize(request *DescribeWebPageGeneralizeRequest) (response *DescribeWebPageGeneralizeResponse, err error) {
	if request == nil {
		request = NewDescribeWebPageGeneralizeRequest()
	}
	response = NewDescribeWebPageGeneralizeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMemShellRulesRequest() (request *ModifyMemShellRulesRequest) {
	request = &ModifyMemShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMemShellRules")
	return
}

func NewModifyMemShellRulesResponse() (response *ModifyMemShellRulesResponse) {
	response = &ModifyMemShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑或者创建java内存马白名单规则
func (c *Client) ModifyMemShellRules(request *ModifyMemShellRulesRequest) (response *ModifyMemShellRulesResponse, err error) {
	if request == nil {
		request = NewModifyMemShellRulesRequest()
	}
	response = NewModifyMemShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperRulesRequest() (request *DescribeFileTamperRulesRequest) {
	request = &DescribeFileTamperRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRules")
	return
}

func NewDescribeFileTamperRulesResponse() (response *DescribeFileTamperRulesResponse) {
	response = &DescribeFileTamperRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 核心文件监控规则列表
func (c *Client) DescribeFileTamperRules(request *DescribeFileTamperRulesRequest) (response *DescribeFileTamperRulesResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperRulesRequest()
	}
	response = NewDescribeFileTamperRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFileTamperEventsRequest() (request *ModifyFileTamperEventsRequest) {
	request = &ModifyFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperEvents")
	return
}

func NewModifyFileTamperEventsResponse() (response *ModifyFileTamperEventsResponse) {
	response = &ModifyFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 核心文件事件更新
func (c *Client) ModifyFileTamperEvents(request *ModifyFileTamperEventsRequest) (response *ModifyFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewModifyFileTamperEventsRequest()
	}
	response = NewModifyFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanMalwareScheduleRequest() (request *DescribeScanMalwareScheduleRequest) {
	request = &DescribeScanMalwareScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanMalwareSchedule")
	return
}

func NewDescribeScanMalwareScheduleResponse() (response *DescribeScanMalwareScheduleResponse) {
	response = &DescribeScanMalwareScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马扫描进度
func (c *Client) DescribeScanMalwareSchedule(request *DescribeScanMalwareScheduleRequest) (response *DescribeScanMalwareScheduleResponse, err error) {
	if request == nil {
		request = NewDescribeScanMalwareScheduleRequest()
	}
	response = NewDescribeScanMalwareScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMalwareWhiteListRequest() (request *ModifyMalwareWhiteListRequest) {
	request = &ModifyMalwareWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMalwareWhiteList")
	return
}

func NewModifyMalwareWhiteListResponse() (response *ModifyMalwareWhiteListResponse) {
	response = &ModifyMalwareWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑木马白名单
func (c *Client) ModifyMalwareWhiteList(request *ModifyMalwareWhiteListRequest) (response *ModifyMalwareWhiteListResponse, err error) {
	if request == nil {
		request = NewModifyMalwareWhiteListRequest()
	}
	response = NewModifyMalwareWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenAttackHotspotRequest() (request *DescribeScreenAttackHotspotRequest) {
	request = &DescribeScreenAttackHotspotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenAttackHotspot")
	return
}

func NewDescribeScreenAttackHotspotResponse() (response *DescribeScreenAttackHotspotResponse) {
	response = &DescribeScreenAttackHotspotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化获取全网攻击热点
func (c *Client) DescribeScreenAttackHotspot(request *DescribeScreenAttackHotspotRequest) (response *DescribeScreenAttackHotspotResponse, err error) {
	if request == nil {
		request = NewDescribeScreenAttackHotspotRequest()
	}
	response = NewDescribeScreenAttackHotspotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskDurationRequest() (request *DescribeTaskDurationRequest) {
	request = &DescribeTaskDurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeTaskDuration")
	return
}

func NewDescribeTaskDurationResponse() (response *DescribeTaskDurationResponse) {
	response = &DescribeTaskDurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务下发时长
func (c *Client) DescribeTaskDuration(request *DescribeTaskDurationRequest) (response *DescribeTaskDurationResponse, err error) {
	if request == nil {
		request = NewDescribeTaskDurationRequest()
	}
	response = NewDescribeTaskDurationResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetPlanTaskListRequest() (request *ExportAssetPlanTaskListRequest) {
	request = &ExportAssetPlanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetPlanTaskList")
	return
}

func NewExportAssetPlanTaskListResponse() (response *ExportAssetPlanTaskListResponse) {
	response = &ExportAssetPlanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理计划任务列表
func (c *Client) ExportAssetPlanTaskList(request *ExportAssetPlanTaskListRequest) (response *ExportAssetPlanTaskListResponse, err error) {
	if request == nil {
		request = NewExportAssetPlanTaskListRequest()
	}
	response = NewExportAssetPlanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebPageServiceInfoRequest() (request *DescribeWebPageServiceInfoRequest) {
	request = &DescribeWebPageServiceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageServiceInfo")
	return
}

func NewDescribeWebPageServiceInfoResponse() (response *DescribeWebPageServiceInfoResponse) {
	response = &DescribeWebPageServiceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网站防篡改-查询网页防篡改服务器购买信息及服务器信息
func (c *Client) DescribeWebPageServiceInfo(request *DescribeWebPageServiceInfoRequest) (response *DescribeWebPageServiceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWebPageServiceInfoRequest()
	}
	response = NewDescribeWebPageServiceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportRansomDefenseStrategyMachinesRequest() (request *ExportRansomDefenseStrategyMachinesRequest) {
	request = &ExportRansomDefenseStrategyMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseStrategyMachines")
	return
}

func NewExportRansomDefenseStrategyMachinesResponse() (response *ExportRansomDefenseStrategyMachinesResponse) {
	response = &ExportRansomDefenseStrategyMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出勒索防御策略绑定机器列表
func (c *Client) ExportRansomDefenseStrategyMachines(request *ExportRansomDefenseStrategyMachinesRequest) (response *ExportRansomDefenseStrategyMachinesResponse, err error) {
	if request == nil {
		request = NewExportRansomDefenseStrategyMachinesRequest()
	}
	response = NewExportRansomDefenseStrategyMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudVendorAccountAKListRequest() (request *DescribeCloudVendorAccountAKListRequest) {
	request = &DescribeCloudVendorAccountAKListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeCloudVendorAccountAKList")
	return
}

func NewDescribeCloudVendorAccountAKListResponse() (response *DescribeCloudVendorAccountAKListResponse) {
	response = &DescribeCloudVendorAccountAKListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户添加的云厂商AK列表
func (c *Client) DescribeCloudVendorAccountAKList(request *DescribeCloudVendorAccountAKListRequest) (response *DescribeCloudVendorAccountAKListResponse, err error) {
	if request == nil {
		request = NewDescribeCloudVendorAccountAKListRequest()
	}
	response = NewDescribeCloudVendorAccountAKListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaSelfLearnConfigRequest() (request *DescribeMaSelfLearnConfigRequest) {
	request = &DescribeMaSelfLearnConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaSelfLearnConfig")
	return
}

func NewDescribeMaSelfLearnConfigResponse() (response *DescribeMaSelfLearnConfigResponse) {
	response = &DescribeMaSelfLearnConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警自学习策略
func (c *Client) DescribeMaSelfLearnConfig(request *DescribeMaSelfLearnConfigRequest) (response *DescribeMaSelfLearnConfigResponse, err error) {
	if request == nil {
		request = NewDescribeMaSelfLearnConfigRequest()
	}
	response = NewDescribeMaSelfLearnConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetAttackWhiteListRequest() (request *DescribeNetAttackWhiteListRequest) {
	request = &DescribeNetAttackWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeNetAttackWhiteList")
	return
}

func NewDescribeNetAttackWhiteListResponse() (response *DescribeNetAttackWhiteListResponse) {
	response = &DescribeNetAttackWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeNetAttackWhiteList(request *DescribeNetAttackWhiteListRequest) (response *DescribeNetAttackWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeNetAttackWhiteListRequest()
	}
	response = NewDescribeNetAttackWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogKafkaDeliverTypeRequest() (request *ModifyLogKafkaDeliverTypeRequest) {
	request = &ModifyLogKafkaDeliverTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaDeliverType")
	return
}

func NewModifyLogKafkaDeliverTypeResponse() (response *ModifyLogKafkaDeliverTypeResponse) {
	response = &ModifyLogKafkaDeliverTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改指定日志类别投递配置、开关
func (c *Client) ModifyLogKafkaDeliverType(request *ModifyLogKafkaDeliverTypeRequest) (response *ModifyLogKafkaDeliverTypeResponse, err error) {
	if request == nil {
		request = NewModifyLogKafkaDeliverTypeRequest()
	}
	response = NewModifyLogKafkaDeliverTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductStatusRequest() (request *DescribeProductStatusRequest) {
	request = &DescribeProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProductStatus")
	return
}

func NewDescribeProductStatusResponse() (response *DescribeProductStatusResponse) {
	response = &DescribeProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品试用状态查询接口
func (c *Client) DescribeProductStatus(request *DescribeProductStatusRequest) (response *DescribeProductStatusResponse, err error) {
	if request == nil {
		request = NewDescribeProductStatusRequest()
	}
	response = NewDescribeProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaHostListRequest() (request *ExportMaHostListRequest) {
	request = &ExportMaHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaHostList")
	return
}

func NewExportMaHostListResponse() (response *ExportMaHostListResponse) {
	response = &ExportMaHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警主机列表
func (c *Client) ExportMaHostList(request *ExportMaHostListRequest) (response *ExportMaHostListResponse, err error) {
	if request == nil {
		request = NewExportMaHostListRequest()
	}
	response = NewExportMaHostListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
	request = &DeleteBruteAttacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBruteAttacks")
	return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
	response = &DeleteBruteAttacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DeleteBruteAttacks) 用于删除暴力破解记录。
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
	if request == nil {
		request = NewDeleteBruteAttacksRequest()
	}
	response = NewDeleteBruteAttacksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbTestStatusRequest() (request *DescribeAbTestStatusRequest) {
	request = &DescribeAbTestStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAbTestStatus")
	return
}

func NewDescribeAbTestStatusResponse() (response *DescribeAbTestStatusResponse) {
	response = &DescribeAbTestStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户当前灰度状态
func (c *Client) DescribeAbTestStatus(request *DescribeAbTestStatusRequest) (response *DescribeAbTestStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAbTestStatusRequest()
	}
	response = NewDescribeAbTestStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityBroadcastsRequest() (request *DescribeSecurityBroadcastsRequest) {
	request = &DescribeSecurityBroadcastsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityBroadcasts")
	return
}

func NewDescribeSecurityBroadcastsResponse() (response *DescribeSecurityBroadcastsResponse) {
	response = &DescribeSecurityBroadcastsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全播报列表页
func (c *Client) DescribeSecurityBroadcasts(request *DescribeSecurityBroadcastsRequest) (response *DescribeSecurityBroadcastsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityBroadcastsRequest()
	}
	response = NewDescribeSecurityBroadcastsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseWhiteConfigRequest() (request *DescribeLicenseWhiteConfigRequest) {
	request = &DescribeLicenseWhiteConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseWhiteConfig")
	return
}

func NewDescribeLicenseWhiteConfigResponse() (response *DescribeLicenseWhiteConfigResponse) {
	response = &DescribeLicenseWhiteConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权白名单的可用配置
func (c *Client) DescribeLicenseWhiteConfig(request *DescribeLicenseWhiteConfigRequest) (response *DescribeLicenseWhiteConfigResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseWhiteConfigRequest()
	}
	response = NewDescribeLicenseWhiteConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
	request = &DescribeComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeComponents")
	return
}

func NewDescribeComponentsResponse() (response *DescribeComponentsResponse) {
	response = &DescribeComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件列表
func (c *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
	if request == nil {
		request = NewDescribeComponentsRequest()
	}
	response = NewDescribeComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBashPoliciesRequest() (request *DeleteBashPoliciesRequest) {
	request = &DeleteBashPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashPolicies")
	return
}

func NewDeleteBashPoliciesResponse() (response *DeleteBashPoliciesResponse) {
	response = &DeleteBashPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高危命令策略
func (c *Client) DeleteBashPolicies(request *DeleteBashPoliciesRequest) (response *DeleteBashPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteBashPoliciesRequest()
	}
	response = NewDeleteBashPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsLogsetsRequest() (request *DescribeClsLogsetsRequest) {
	request = &DescribeClsLogsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeClsLogsets")
	return
}

func NewDescribeClsLogsetsResponse() (response *DescribeClsLogsetsResponse) {
	response = &DescribeClsLogsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cls日志集列表
func (c *Client) DescribeClsLogsets(request *DescribeClsLogsetsRequest) (response *DescribeClsLogsetsResponse, err error) {
	if request == nil {
		request = NewDescribeClsLogsetsRequest()
	}
	response = NewDescribeClsLogsetsResponse()
	err = c.Send(request, response)
	return
}

func NewExportBashEventsNewRequest() (request *ExportBashEventsNewRequest) {
	request = &ExportBashEventsNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBashEventsNew")
	return
}

func NewExportBashEventsNewResponse() (response *ExportBashEventsNewResponse) {
	response = &ExportBashEventsNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出高危命令事件(新)
func (c *Client) ExportBashEventsNew(request *ExportBashEventsNewRequest) (response *ExportBashEventsNewResponse, err error) {
	if request == nil {
		request = NewExportBashEventsNewRequest()
	}
	response = NewExportBashEventsNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanStateRequest() (request *DescribeScanStateRequest) {
	request = &DescribeScanStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanState")
	return
}

func NewDescribeScanStateResponse() (response *DescribeScanStateResponse) {
	response = &DescribeScanStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeScanState 该接口能查询对应模块正在进行的扫描任务状态
func (c *Client) DescribeScanState(request *DescribeScanStateRequest) (response *DescribeScanStateResponse, err error) {
	if request == nil {
		request = NewDescribeScanStateRequest()
	}
	response = NewDescribeScanStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppPluginListRequest() (request *DescribeAssetWebAppPluginListRequest) {
	request = &DescribeAssetWebAppPluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppPluginList")
	return
}

func NewDescribeAssetWebAppPluginListResponse() (response *DescribeAssetWebAppPluginListResponse) {
	response = &DescribeAssetWebAppPluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理Web应用插件列表
func (c *Client) DescribeAssetWebAppPluginList(request *DescribeAssetWebAppPluginListRequest) (response *DescribeAssetWebAppPluginListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppPluginListRequest()
	}
	response = NewDescribeAssetWebAppPluginListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseRollBackTaskListRequest() (request *DescribeRansomDefenseRollBackTaskListRequest) {
	request = &DescribeRansomDefenseRollBackTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseRollBackTaskList")
	return
}

func NewDescribeRansomDefenseRollBackTaskListResponse() (response *DescribeRansomDefenseRollBackTaskListResponse) {
	response = &DescribeRansomDefenseRollBackTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询回滚任务列表
func (c *Client) DescribeRansomDefenseRollBackTaskList(request *DescribeRansomDefenseRollBackTaskListRequest) (response *DescribeRansomDefenseRollBackTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseRollBackTaskListRequest()
	}
	response = NewDescribeRansomDefenseRollBackTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaLabelRequest() (request *ModifyMaLabelRequest) {
	request = &ModifyMaLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaLabel")
	return
}

func NewModifyMaLabelResponse() (response *ModifyMaLabelResponse) {
	response = &ModifyMaLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或更改标签信息
func (c *Client) ModifyMaLabel(request *ModifyMaLabelRequest) (response *ModifyMaLabelResponse, err error) {
	if request == nil {
		request = NewModifyMaLabelRequest()
	}
	response = NewModifyMaLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogClsDeliverInfoRequest() (request *DescribeLogClsDeliverInfoRequest) {
	request = &DescribeLogClsDeliverInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogClsDeliverInfo")
	return
}

func NewDescribeLogClsDeliverInfoResponse() (response *DescribeLogClsDeliverInfoResponse) {
	response = &DescribeLogClsDeliverInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回cls投递配置
func (c *Client) DescribeLogClsDeliverInfo(request *DescribeLogClsDeliverInfoRequest) (response *DescribeLogClsDeliverInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLogClsDeliverInfoRequest()
	}
	response = NewDescribeLogClsDeliverInfoResponse()
	err = c.Send(request, response)
	return
}

func NewStopNoticeBanTipsRequest() (request *StopNoticeBanTipsRequest) {
	request = &StopNoticeBanTipsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "StopNoticeBanTips")
	return
}

func NewStopNoticeBanTipsResponse() (response *StopNoticeBanTipsResponse) {
	response = &StopNoticeBanTipsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 不再提醒爆破阻断提示弹窗
func (c *Client) StopNoticeBanTips(request *StopNoticeBanTipsRequest) (response *StopNoticeBanTipsResponse, err error) {
	if request == nil {
		request = NewStopNoticeBanTipsRequest()
	}
	response = NewStopNoticeBanTipsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInjectRiskyServiceSwitchRequest() (request *ModifyInjectRiskyServiceSwitchRequest) {
	request = &ModifyInjectRiskyServiceSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyInjectRiskyServiceSwitch")
	return
}

func NewModifyInjectRiskyServiceSwitchResponse() (response *ModifyInjectRiskyServiceSwitchResponse) {
	response = &ModifyInjectRiskyServiceSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内存马和rasp是否注入风险服务开关
func (c *Client) ModifyInjectRiskyServiceSwitch(request *ModifyInjectRiskyServiceSwitchRequest) (response *ModifyInjectRiskyServiceSwitchResponse, err error) {
	if request == nil {
		request = NewModifyInjectRiskyServiceSwitchRequest()
	}
	response = NewModifyInjectRiskyServiceSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenDefenseTrendsRequest() (request *DescribeScreenDefenseTrendsRequest) {
	request = &DescribeScreenDefenseTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenDefenseTrends")
	return
}

func NewDescribeScreenDefenseTrendsResponse() (response *DescribeScreenDefenseTrendsResponse) {
	response = &DescribeScreenDefenseTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化防趋势接口
func (c *Client) DescribeScreenDefenseTrends(request *DescribeScreenDefenseTrendsRequest) (response *DescribeScreenDefenseTrendsResponse, err error) {
	if request == nil {
		request = NewDescribeScreenDefenseTrendsRequest()
	}
	response = NewDescribeScreenDefenseTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVersionStatisticsRequest() (request *DescribeVersionStatisticsRequest) {
	request = &DescribeVersionStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVersionStatistics")
	return
}

func NewDescribeVersionStatisticsResponse() (response *DescribeVersionStatisticsResponse) {
	response = &DescribeVersionStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于统计专业版和基础版机器数。
func (c *Client) DescribeVersionStatistics(request *DescribeVersionStatisticsRequest) (response *DescribeVersionStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeVersionStatisticsRequest()
	}
	response = NewDescribeVersionStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQuickProtectSettingRequest() (request *CreateQuickProtectSettingRequest) {
	request = &CreateQuickProtectSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateQuickProtectSetting")
	return
}

func NewCreateQuickProtectSettingResponse() (response *CreateQuickProtectSettingResponse) {
	response = &CreateQuickProtectSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键创建快捷防护设置
func (c *Client) CreateQuickProtectSetting(request *CreateQuickProtectSettingRequest) (response *CreateQuickProtectSettingResponse, err error) {
	if request == nil {
		request = NewCreateQuickProtectSettingRequest()
	}
	response = NewCreateQuickProtectSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaliciousRequestWhiteListRequest() (request *DeleteMaliciousRequestWhiteListRequest) {
	request = &DeleteMaliciousRequestWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaliciousRequestWhiteList")
	return
}

func NewDeleteMaliciousRequestWhiteListResponse() (response *DeleteMaliciousRequestWhiteListResponse) {
	response = &DeleteMaliciousRequestWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求白名单
func (c *Client) DeleteMaliciousRequestWhiteList(request *DeleteMaliciousRequestWhiteListRequest) (response *DeleteMaliciousRequestWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteMaliciousRequestWhiteListRequest()
	}
	response = NewDeleteMaliciousRequestWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebPageEventLogRequest() (request *DeleteWebPageEventLogRequest) {
	request = &DeleteWebPageEventLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebPageEventLog")
	return
}

func NewDeleteWebPageEventLogResponse() (response *DeleteWebPageEventLogResponse) {
	response = &DeleteWebPageEventLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网站防篡改-删除事件记录
func (c *Client) DeleteWebPageEventLog(request *DeleteWebPageEventLogRequest) (response *DeleteWebPageEventLogResponse, err error) {
	if request == nil {
		request = NewDeleteWebPageEventLogRequest()
	}
	response = NewDeleteWebPageEventLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineHostIgnoreListRequest() (request *DescribeBaselineHostIgnoreListRequest) {
	request = &DescribeBaselineHostIgnoreListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostIgnoreList")
	return
}

func NewDescribeBaselineHostIgnoreListResponse() (response *DescribeBaselineHostIgnoreListResponse) {
	response = &DescribeBaselineHostIgnoreListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取忽略规则主机列表
func (c *Client) DescribeBaselineHostIgnoreList(request *DescribeBaselineHostIgnoreListRequest) (response *DescribeBaselineHostIgnoreListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineHostIgnoreListRequest()
	}
	response = NewDescribeBaselineHostIgnoreListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashRulesRequest() (request *DescribeBashRulesRequest) {
	request = &DescribeBashRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashRules")
	return
}

func NewDescribeBashRulesResponse() (response *DescribeBashRulesResponse) {
	response = &DescribeBashRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高危命令规则列表
func (c *Client) DescribeBashRules(request *DescribeBashRulesRequest) (response *DescribeBashRulesResponse, err error) {
	if request == nil {
		request = NewDescribeBashRulesRequest()
	}
	response = NewDescribeBashRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEmergencyVulListRequest() (request *DescribeEmergencyVulListRequest) {
	request = &DescribeEmergencyVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeEmergencyVulList")
	return
}

func NewDescribeEmergencyVulListResponse() (response *DescribeEmergencyVulListResponse) {
	response = &DescribeEmergencyVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应急漏洞列表
func (c *Client) DescribeEmergencyVulList(request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
	if request == nil {
		request = NewDescribeEmergencyVulListRequest()
	}
	response = NewDescribeEmergencyVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningServiceIPListRequest() (request *DescribeWarningServiceIPListRequest) {
	request = &DescribeWarningServiceIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWarningServiceIPList")
	return
}

func NewDescribeWarningServiceIPListResponse() (response *DescribeWarningServiceIPListResponse) {
	response = &DescribeWarningServiceIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警服务IP名单
func (c *Client) DescribeWarningServiceIPList(request *DescribeWarningServiceIPListRequest) (response *DescribeWarningServiceIPListResponse, err error) {
	if request == nil {
		request = NewDescribeWarningServiceIPListRequest()
	}
	response = NewDescribeWarningServiceIPListResponse()
	err = c.Send(request, response)
	return
}

func NewClearLocalStorageRequest() (request *ClearLocalStorageRequest) {
	request = &ClearLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ClearLocalStorage")
	return
}

func NewClearLocalStorageResponse() (response *ClearLocalStorageResponse) {
	response = &ClearLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理本地存储数据
func (c *Client) ClearLocalStorage(request *ClearLocalStorageRequest) (response *ClearLocalStorageResponse, err error) {
	if request == nil {
		request = NewClearLocalStorageRequest()
	}
	response = NewClearLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
	request = &DeleteLoginWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteLoginWhiteList")
	return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
	response = &DeleteLoginWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除异地登录白名单规则。
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteLoginWhiteListRequest()
	}
	response = NewDeleteLoginWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationInfoRequest() (request *DescribeAssetWebLocationInfoRequest) {
	request = &DescribeAssetWebLocationInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationInfo")
	return
}

func NewDescribeAssetWebLocationInfoResponse() (response *DescribeAssetWebLocationInfoResponse) {
	response = &DescribeAssetWebLocationInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点详情
func (c *Client) DescribeAssetWebLocationInfo(request *DescribeAssetWebLocationInfoRequest) (response *DescribeAssetWebLocationInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationInfoRequest()
	}
	response = NewDescribeAssetWebLocationInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachineLoginTypeRequest() (request *ModifyMachineLoginTypeRequest) {
	request = &ModifyMachineLoginTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMachineLoginType")
	return
}

func NewModifyMachineLoginTypeResponse() (response *ModifyMachineLoginTypeResponse) {
	response = &ModifyMachineLoginTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改主机登录方式
func (c *Client) ModifyMachineLoginType(request *ModifyMachineLoginTypeRequest) (response *ModifyMachineLoginTypeResponse, err error) {
	if request == nil {
		request = NewModifyMachineLoginTypeRequest()
	}
	response = NewModifyMachineLoginTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackEventsRequest() (request *DescribeAttackEventsRequest) {
	request = &DescribeAttackEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackEvents")
	return
}

func NewDescribeAttackEventsResponse() (response *DescribeAttackEventsResponse) {
	response = &DescribeAttackEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeAttackEvents(request *DescribeAttackEventsRequest) (response *DescribeAttackEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackEventsRequest()
	}
	response = NewDescribeAttackEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportMachinesRequest() (request *DescribeExportMachinesRequest) {
	request = &DescribeExportMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeExportMachines")
	return
}

func NewDescribeExportMachinesResponse() (response *DescribeExportMachinesResponse) {
	response = &DescribeExportMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeExportMachines) 用于导出区域主机列表。
func (c *Client) DescribeExportMachines(request *DescribeExportMachinesRequest) (response *DescribeExportMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeExportMachinesRequest()
	}
	response = NewDescribeExportMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFirstScanBaselineRequest() (request *CheckFirstScanBaselineRequest) {
	request = &CheckFirstScanBaselineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CheckFirstScanBaseline")
	return
}

func NewCheckFirstScanBaselineResponse() (response *CheckFirstScanBaselineResponse) {
	response = &CheckFirstScanBaselineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询基线是否第一次检测
func (c *Client) CheckFirstScanBaseline(request *CheckFirstScanBaselineRequest) (response *CheckFirstScanBaselineResponse, err error) {
	if request == nil {
		request = NewCheckFirstScanBaselineRequest()
	}
	response = NewCheckFirstScanBaselineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVersionCompareChartRequest() (request *DescribeVersionCompareChartRequest) {
	request = &DescribeVersionCompareChartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVersionCompareChart")
	return
}

func NewDescribeVersionCompareChartResponse() (response *DescribeVersionCompareChartResponse) {
	response = &DescribeVersionCompareChartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取版本对比信息
func (c *Client) DescribeVersionCompareChart(request *DescribeVersionCompareChartRequest) (response *DescribeVersionCompareChartResponse, err error) {
	if request == nil {
		request = NewDescribeVersionCompareChartRequest()
	}
	response = NewDescribeVersionCompareChartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTotalCountRequest() (request *DescribeAssetTotalCountRequest) {
	request = &DescribeAssetTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTotalCount")
	return
}

func NewDescribeAssetTotalCountResponse() (response *DescribeAssetTotalCountResponse) {
	response = &DescribeAssetTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有资源数量：主机、账号、端口、进程、软件、数据库、Web应用、Web框架、Web服务、Web站点
func (c *Client) DescribeAssetTotalCount(request *DescribeAssetTotalCountRequest) (response *DescribeAssetTotalCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTotalCountRequest()
	}
	response = NewDescribeAssetTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBruteAttackRulesRequest() (request *DescribeBruteAttackRulesRequest) {
	request = &DescribeBruteAttackRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBruteAttackRules")
	return
}

func NewDescribeBruteAttackRulesResponse() (response *DescribeBruteAttackRulesResponse) {
	response = &DescribeBruteAttackRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取爆破破解规则
func (c *Client) DescribeBruteAttackRules(request *DescribeBruteAttackRulesRequest) (response *DescribeBruteAttackRulesResponse, err error) {
	if request == nil {
		request = NewDescribeBruteAttackRulesRequest()
	}
	response = NewDescribeBruteAttackRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaGlobalConfRequest() (request *ModifyMaGlobalConfRequest) {
	request = &ModifyMaGlobalConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaGlobalConf")
	return
}

func NewModifyMaGlobalConfResponse() (response *ModifyMaGlobalConfResponse) {
	response = &ModifyMaGlobalConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改微告警全局配置
func (c *Client) ModifyMaGlobalConf(request *ModifyMaGlobalConfRequest) (response *ModifyMaGlobalConfResponse, err error) {
	if request == nil {
		request = NewModifyMaGlobalConfRequest()
	}
	response = NewModifyMaGlobalConfResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRiskDnsPolicyRequest() (request *DeleteRiskDnsPolicyRequest) {
	request = &DeleteRiskDnsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteRiskDnsPolicy")
	return
}

func NewDeleteRiskDnsPolicyResponse() (response *DeleteRiskDnsPolicyResponse) {
	response = &DeleteRiskDnsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求策略
func (c *Client) DeleteRiskDnsPolicy(request *DeleteRiskDnsPolicyRequest) (response *DeleteRiskDnsPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteRiskDnsPolicyRequest()
	}
	response = NewDeleteRiskDnsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseListRequest() (request *DescribeAssetDatabaseListRequest) {
	request = &DescribeAssetDatabaseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseList")
	return
}

func NewDescribeAssetDatabaseListResponse() (response *DescribeAssetDatabaseListResponse) {
	response = &DescribeAssetDatabaseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理数据库列表
func (c *Client) DescribeAssetDatabaseList(request *DescribeAssetDatabaseListRequest) (response *DescribeAssetDatabaseListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseListRequest()
	}
	response = NewDescribeAssetDatabaseListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineDetailRequest() (request *DescribeBaselineDetailRequest) {
	request = &DescribeBaselineDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetail")
	return
}

func NewDescribeBaselineDetailResponse() (response *DescribeBaselineDetailResponse) {
	response = &DescribeBaselineDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线id查询基线详情接口
func (c *Client) DescribeBaselineDetail(request *DescribeBaselineDetailRequest) (response *DescribeBaselineDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineDetailRequest()
	}
	response = NewDescribeBaselineDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEmergencyVulScanRequest() (request *CreateEmergencyVulScanRequest) {
	request = &CreateEmergencyVulScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateEmergencyVulScan")
	return
}

func NewCreateEmergencyVulScanResponse() (response *CreateEmergencyVulScanResponse) {
	response = &CreateEmergencyVulScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应急漏洞扫描任务
func (c *Client) CreateEmergencyVulScan(request *CreateEmergencyVulScanRequest) (response *CreateEmergencyVulScanResponse, err error) {
	if request == nil {
		request = NewCreateEmergencyVulScanRequest()
	}
	response = NewCreateEmergencyVulScanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebPageProtectStatRequest() (request *DescribeWebPageProtectStatRequest) {
	request = &DescribeWebPageProtectStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageProtectStat")
	return
}

func NewDescribeWebPageProtectStatResponse() (response *DescribeWebPageProtectStatResponse) {
	response = &DescribeWebPageProtectStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网站防篡改-查询动态防护信息
func (c *Client) DescribeWebPageProtectStat(request *DescribeWebPageProtectStatRequest) (response *DescribeWebPageProtectStatResponse, err error) {
	if request == nil {
		request = NewDescribeWebPageProtectStatRequest()
	}
	response = NewDescribeWebPageProtectStatResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetDatabaseListRequest() (request *ExportAssetDatabaseListRequest) {
	request = &ExportAssetDatabaseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetDatabaseList")
	return
}

func NewExportAssetDatabaseListResponse() (response *ExportAssetDatabaseListResponse) {
	response = &ExportAssetDatabaseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理数据库列表
func (c *Client) ExportAssetDatabaseList(request *ExportAssetDatabaseListRequest) (response *ExportAssetDatabaseListResponse, err error) {
	if request == nil {
		request = NewExportAssetDatabaseListRequest()
	}
	response = NewExportAssetDatabaseListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWarningSettingRequest() (request *ModifyWarningSettingRequest) {
	request = &ModifyWarningSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWarningSetting")
	return
}

func NewModifyWarningSettingResponse() (response *ModifyWarningSettingResponse) {
	response = &ModifyWarningSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警设置
func (c *Client) ModifyWarningSetting(request *ModifyWarningSettingRequest) (response *ModifyWarningSettingResponse, err error) {
	if request == nil {
		request = NewModifyWarningSettingRequest()
	}
	response = NewModifyWarningSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLicenseRecordRequest() (request *DeleteLicenseRecordRequest) {
	request = &DeleteLicenseRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteLicenseRecord")
	return
}

func NewDeleteLicenseRecordResponse() (response *DeleteLicenseRecordResponse) {
	response = &DeleteLicenseRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对授权管理-订单列表内已过期的订单进行删除.(删除后的订单不在统计范畴内)
func (c *Client) DeleteLicenseRecord(request *DeleteLicenseRecordRequest) (response *DeleteLicenseRecordResponse, err error) {
	if request == nil {
		request = NewDeleteLicenseRecordRequest()
	}
	response = NewDeleteLicenseRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellRulesAggregationRequest() (request *DeleteReverseShellRulesAggregationRequest) {
	request = &DeleteReverseShellRulesAggregationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteReverseShellRulesAggregation")
	return
}

func NewDeleteReverseShellRulesAggregationResponse() (response *DeleteReverseShellRulesAggregationResponse) {
	response = &DeleteReverseShellRulesAggregationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除反弹Shell规则-聚合版本
func (c *Client) DeleteReverseShellRulesAggregation(request *DeleteReverseShellRulesAggregationRequest) (response *DeleteReverseShellRulesAggregationResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellRulesAggregationRequest()
	}
	response = NewDeleteReverseShellRulesAggregationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
	request = &DescribeReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellEvents")
	return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
	response = &DescribeReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取反弹Shell列表
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsRequest()
	}
	response = NewDescribeReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaCustomPolicyRequest() (request *DeleteMaCustomPolicyRequest) {
	request = &DeleteMaCustomPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaCustomPolicy")
	return
}

func NewDeleteMaCustomPolicyResponse() (response *DeleteMaCustomPolicyResponse) {
	response = &DeleteMaCustomPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微告警自定义白名单
func (c *Client) DeleteMaCustomPolicy(request *DeleteMaCustomPolicyRequest) (response *DeleteMaCustomPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteMaCustomPolicyRequest()
	}
	response = NewDeleteMaCustomPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserListRequest() (request *DescribeAssetUserListRequest) {
	request = &DescribeAssetUserListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserList")
	return
}

func NewDescribeAssetUserListResponse() (response *DescribeAssetUserListResponse) {
	response = &DescribeAssetUserListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号列表
func (c *Client) DescribeAssetUserList(request *DescribeAssetUserListRequest) (response *DescribeAssetUserListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserListRequest()
	}
	response = NewDescribeAssetUserListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgnoreRuleEffectHostListRequest() (request *DescribeIgnoreRuleEffectHostListRequest) {
	request = &DescribeIgnoreRuleEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreRuleEffectHostList")
	return
}

func NewDescribeIgnoreRuleEffectHostListResponse() (response *DescribeIgnoreRuleEffectHostListResponse) {
	response = &DescribeIgnoreRuleEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据检测项id与筛选条件查询忽略检测项影响主机列表信息
func (c *Client) DescribeIgnoreRuleEffectHostList(request *DescribeIgnoreRuleEffectHostListRequest) (response *DescribeIgnoreRuleEffectHostListResponse, err error) {
	if request == nil {
		request = NewDescribeIgnoreRuleEffectHostListRequest()
	}
	response = NewDescribeIgnoreRuleEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaPodListRequest() (request *ExportMaPodListRequest) {
	request = &ExportMaPodListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaPodList")
	return
}

func NewExportMaPodListResponse() (response *ExportMaPodListResponse) {
	response = &ExportMaPodListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警POD列表
func (c *Client) ExportMaPodList(request *ExportMaPodListRequest) (response *ExportMaPodListResponse, err error) {
	if request == nil {
		request = NewExportMaPodListRequest()
	}
	response = NewExportMaPodListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetPortInfoListRequest() (request *ExportAssetPortInfoListRequest) {
	request = &ExportAssetPortInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetPortInfoList")
	return
}

func NewExportAssetPortInfoListResponse() (response *ExportAssetPortInfoListResponse) {
	response = &ExportAssetPortInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理端口列表
func (c *Client) ExportAssetPortInfoList(request *ExportAssetPortInfoListRequest) (response *ExportAssetPortInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetPortInfoListRequest()
	}
	response = NewExportAssetPortInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetUserListRequest() (request *ExportAssetUserListRequest) {
	request = &ExportAssetUserListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetUserList")
	return
}

func NewExportAssetUserListResponse() (response *ExportAssetUserListResponse) {
	response = &ExportAssetUserListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出账号列表
func (c *Client) ExportAssetUserList(request *ExportAssetUserListRequest) (response *ExportAssetUserListResponse, err error) {
	if request == nil {
		request = NewExportAssetUserListRequest()
	}
	response = NewExportAssetUserListResponse()
	err = c.Send(request, response)
	return
}

func NewExportRansomDefenseEventsListRequest() (request *ExportRansomDefenseEventsListRequest) {
	request = &ExportRansomDefenseEventsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseEventsList")
	return
}

func NewExportRansomDefenseEventsListResponse() (response *ExportRansomDefenseEventsListResponse) {
	response = &ExportRansomDefenseEventsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出防勒索事件列表
func (c *Client) ExportRansomDefenseEventsList(request *ExportRansomDefenseEventsListRequest) (response *ExportRansomDefenseEventsListResponse, err error) {
	if request == nil {
		request = NewExportRansomDefenseEventsListRequest()
	}
	response = NewExportRansomDefenseEventsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackTrendsRequest() (request *DescribeAttackTrendsRequest) {
	request = &DescribeAttackTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackTrends")
	return
}

func NewDescribeAttackTrendsResponse() (response *DescribeAttackTrendsResponse) {
	response = &DescribeAttackTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DescribeAttackTrends(request *DescribeAttackTrendsRequest) (response *DescribeAttackTrendsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackTrendsRequest()
	}
	response = NewDescribeAttackTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseStateRequest() (request *DescribeRansomDefenseStateRequest) {
	request = &DescribeRansomDefenseStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseState")
	return
}

func NewDescribeRansomDefenseStateResponse() (response *DescribeRansomDefenseStateResponse) {
	response = &DescribeRansomDefenseStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户防勒索趋势
func (c *Client) DescribeRansomDefenseState(request *DescribeRansomDefenseStateRequest) (response *DescribeRansomDefenseStateResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseStateRequest()
	}
	response = NewDescribeRansomDefenseStateResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDetectionReportRequest() (request *ExportVulDetectionReportRequest) {
	request = &ExportVulDetectionReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDetectionReport")
	return
}

func NewExportVulDetectionReportResponse() (response *ExportVulDetectionReportResponse) {
	response = &ExportVulDetectionReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞检测报告。
func (c *Client) ExportVulDetectionReport(request *ExportVulDetectionReportRequest) (response *ExportVulDetectionReportResponse, err error) {
	if request == nil {
		request = NewExportVulDetectionReportRequest()
	}
	response = NewExportVulDetectionReportResponse()
	err = c.Send(request, response)
	return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
	request = &CloseProVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CloseProVersion")
	return
}

func NewCloseProVersionResponse() (response *CloseProVersionResponse) {
	response = &CloseProVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CloseProVersion) 已全面改为授权模式,该接口只能解绑授权.
func (c *Client) CloseProVersion(request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
	if request == nil {
		request = NewCloseProVersionRequest()
	}
	response = NewCloseProVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRiskDnsEventRequest() (request *DeleteRiskDnsEventRequest) {
	request = &DeleteRiskDnsEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteRiskDnsEvent")
	return
}

func NewDeleteRiskDnsEventResponse() (response *DeleteRiskDnsEventResponse) {
	response = &DeleteRiskDnsEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求事件
func (c *Client) DeleteRiskDnsEvent(request *DeleteRiskDnsEventRequest) (response *DeleteRiskDnsEventResponse, err error) {
	if request == nil {
		request = NewDeleteRiskDnsEventRequest()
	}
	response = NewDeleteRiskDnsEventResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetMachineDetailRequest() (request *ExportAssetMachineDetailRequest) {
	request = &ExportAssetMachineDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetMachineDetail")
	return
}

func NewExportAssetMachineDetailResponse() (response *ExportAssetMachineDetailResponse) {
	response = &ExportAssetMachineDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理主机资源详细信息
func (c *Client) ExportAssetMachineDetail(request *ExportAssetMachineDetailRequest) (response *ExportAssetMachineDetailResponse, err error) {
	if request == nil {
		request = NewExportAssetMachineDetailRequest()
	}
	response = NewExportAssetMachineDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMalwareWhiteListRequest() (request *DeleteMalwareWhiteListRequest) {
	request = &DeleteMalwareWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwareWhiteList")
	return
}

func NewDeleteMalwareWhiteListResponse() (response *DeleteMalwareWhiteListResponse) {
	response = &DeleteMalwareWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除木马白名单
func (c *Client) DeleteMalwareWhiteList(request *DeleteMalwareWhiteListRequest) (response *DeleteMalwareWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteMalwareWhiteListRequest()
	}
	response = NewDeleteMalwareWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHistoryServiceRequest() (request *DescribeHistoryServiceRequest) {
	request = &DescribeHistoryServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeHistoryService")
	return
}

func NewDescribeHistoryServiceResponse() (response *DescribeHistoryServiceResponse) {
	response = &DescribeHistoryServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志检索服务信息
func (c *Client) DescribeHistoryService(request *DescribeHistoryServiceRequest) (response *DescribeHistoryServiceResponse, err error) {
	if request == nil {
		request = NewDescribeHistoryServiceRequest()
	}
	response = NewDescribeHistoryServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTopRequest() (request *DescribeVulTopRequest) {
	request = &DescribeVulTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulTop")
	return
}

func NewDescribeVulTopResponse() (response *DescribeVulTopResponse) {
	response = &DescribeVulTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞top统计
func (c *Client) DescribeVulTop(request *DescribeVulTopRequest) (response *DescribeVulTopResponse, err error) {
	if request == nil {
		request = NewDescribeVulTopRequest()
	}
	response = NewDescribeVulTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
	request = &DescribeUsualLoginPlacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeUsualLoginPlaces")
	return
}

func NewDescribeUsualLoginPlacesResponse() (response *DescribeUsualLoginPlacesResponse) {
	response = &DescribeUsualLoginPlacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口（DescribeUsualLoginPlaces）用于查询常用登录地。
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
	if request == nil {
		request = NewDescribeUsualLoginPlacesRequest()
	}
	response = NewDescribeUsualLoginPlacesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSearchTemplateRequest() (request *CreateSearchTemplateRequest) {
	request = &CreateSearchTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateSearchTemplate")
	return
}

func NewCreateSearchTemplateResponse() (response *CreateSearchTemplateResponse) {
	response = &CreateSearchTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加检索模板
func (c *Client) CreateSearchTemplate(request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSearchTemplateRequest()
	}
	response = NewCreateSearchTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginWhiteInfoRequest() (request *ModifyLoginWhiteInfoRequest) {
	request = &ModifyLoginWhiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLoginWhiteInfo")
	return
}

func NewModifyLoginWhiteInfoResponse() (response *ModifyLoginWhiteInfoResponse) {
	response = &ModifyLoginWhiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新登录审计白名单信息
func (c *Client) ModifyLoginWhiteInfo(request *ModifyLoginWhiteInfoRequest) (response *ModifyLoginWhiteInfoResponse, err error) {
	if request == nil {
		request = NewModifyLoginWhiteInfoRequest()
	}
	response = NewModifyLoginWhiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaConnEventListRequest() (request *ExportMaConnEventListRequest) {
	request = &ExportMaConnEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaConnEventList")
	return
}

func NewExportMaConnEventListResponse() (response *ExportMaConnEventListResponse) {
	response = &ExportMaConnEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警访问事件
func (c *Client) ExportMaConnEventList(request *ExportMaConnEventListRequest) (response *ExportMaConnEventListResponse, err error) {
	if request == nil {
		request = NewExportMaConnEventListRequest()
	}
	response = NewExportMaConnEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineSnapshotRequest() (request *DescribeMachineSnapshotRequest) {
	request = &DescribeMachineSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineSnapshot")
	return
}

func NewDescribeMachineSnapshotResponse() (response *DescribeMachineSnapshotResponse) {
	response = &DescribeMachineSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞修护-查询主机创建的快照
func (c *Client) DescribeMachineSnapshot(request *DescribeMachineSnapshotRequest) (response *DescribeMachineSnapshotResponse, err error) {
	if request == nil {
		request = NewDescribeMachineSnapshotRequest()
	}
	response = NewDescribeMachineSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulLabelsRequest() (request *DescribeVulLabelsRequest) {
	request = &DescribeVulLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulLabels")
	return
}

func NewDescribeVulLabelsResponse() (response *DescribeVulLabelsResponse) {
	response = &DescribeVulLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户漏洞所有标签列表
func (c *Client) DescribeVulLabels(request *DescribeVulLabelsRequest) (response *DescribeVulLabelsResponse, err error) {
	if request == nil {
		request = NewDescribeVulLabelsRequest()
	}
	response = NewDescribeVulLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebAppListRequest() (request *ExportAssetWebAppListRequest) {
	request = &ExportAssetWebAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebAppList")
	return
}

func NewExportAssetWebAppListResponse() (response *ExportAssetWebAppListResponse) {
	response = &ExportAssetWebAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web应用列表
func (c *Client) ExportAssetWebAppList(request *ExportAssetWebAppListRequest) (response *ExportAssetWebAppListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebAppListRequest()
	}
	response = NewExportAssetWebAppListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogStorageConfigRequest() (request *ModifyLogStorageConfigRequest) {
	request = &ModifyLogStorageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogStorageConfig")
	return
}

func NewModifyLogStorageConfigResponse() (response *ModifyLogStorageConfigResponse) {
	response = &ModifyLogStorageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志存储配置
func (c *Client) ModifyLogStorageConfig(request *ModifyLogStorageConfigRequest) (response *ModifyLogStorageConfigResponse, err error) {
	if request == nil {
		request = NewModifyLogStorageConfigRequest()
	}
	response = NewModifyLogStorageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineHostRiskTopRequest() (request *DescribeBaselineHostRiskTopRequest) {
	request = &DescribeBaselineHostRiskTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostRiskTop")
	return
}

func NewDescribeBaselineHostRiskTopResponse() (response *DescribeBaselineHostRiskTopResponse) {
	response = &DescribeBaselineHostRiskTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线服务器风险TOP5
func (c *Client) DescribeBaselineHostRiskTop(request *DescribeBaselineHostRiskTopRequest) (response *DescribeBaselineHostRiskTopResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineHostRiskTopRequest()
	}
	response = NewDescribeBaselineHostRiskTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceEventRequest() (request *DescribeVulDefenceEventRequest) {
	request = &DescribeVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceEvent")
	return
}

func NewDescribeVulDefenceEventResponse() (response *DescribeVulDefenceEventResponse) {
	response = &DescribeVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御事件列表
func (c *Client) DescribeVulDefenceEvent(request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceEventRequest()
	}
	response = NewDescribeVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetAttackSettingRequest() (request *ModifyNetAttackSettingRequest) {
	request = &ModifyNetAttackSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyNetAttackSetting")
	return
}

func NewModifyNetAttackSettingResponse() (response *ModifyNetAttackSettingResponse) {
	response = &ModifyNetAttackSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) ModifyNetAttackSetting(request *ModifyNetAttackSettingRequest) (response *ModifyNetAttackSettingResponse, err error) {
	if request == nil {
		request = NewModifyNetAttackSettingRequest()
	}
	response = NewModifyNetAttackSettingResponse()
	err = c.Send(request, response)
	return
}

func NewExportBashPoliciesRequest() (request *ExportBashPoliciesRequest) {
	request = &ExportBashPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBashPolicies")
	return
}

func NewExportBashPoliciesResponse() (response *ExportBashPoliciesResponse) {
	response = &ExportBashPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出高危命令策略
func (c *Client) ExportBashPolicies(request *ExportBashPoliciesRequest) (response *ExportBashPoliciesResponse, err error) {
	if request == nil {
		request = NewExportBashPoliciesRequest()
	}
	response = NewExportBashPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetEnvListRequest() (request *DescribeAssetEnvListRequest) {
	request = &DescribeAssetEnvListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetEnvList")
	return
}

func NewDescribeAssetEnvListResponse() (response *DescribeAssetEnvListResponse) {
	response = &DescribeAssetEnvListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理环境变量列表
func (c *Client) DescribeAssetEnvList(request *DescribeAssetEnvListRequest) (response *DescribeAssetEnvListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetEnvListRequest()
	}
	response = NewDescribeAssetEnvListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMalwareTimingScanSettingsRequest() (request *ModifyMalwareTimingScanSettingsRequest) {
	request = &ModifyMalwareTimingScanSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMalwareTimingScanSettings")
	return
}

func NewModifyMalwareTimingScanSettingsResponse() (response *ModifyMalwareTimingScanSettingsResponse) {
	response = &ModifyMalwareTimingScanSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 定时扫描设置
func (c *Client) ModifyMalwareTimingScanSettings(request *ModifyMalwareTimingScanSettingsRequest) (response *ModifyMalwareTimingScanSettingsResponse, err error) {
	if request == nil {
		request = NewModifyMalwareTimingScanSettingsRequest()
	}
	response = NewModifyMalwareTimingScanSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskBatchStatusRequest() (request *DescribeRiskBatchStatusRequest) {
	request = &DescribeRiskBatchStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskBatchStatus")
	return
}

func NewDescribeRiskBatchStatusResponse() (response *DescribeRiskBatchStatusResponse) {
	response = &DescribeRiskBatchStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询入侵检测事件更新状态任务是否完成
func (c *Client) DescribeRiskBatchStatus(request *DescribeRiskBatchStatusRequest) (response *DescribeRiskBatchStatusResponse, err error) {
	if request == nil {
		request = NewDescribeRiskBatchStatusRequest()
	}
	response = NewDescribeRiskBatchStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulEffectModulesRequest() (request *DescribeVulEffectModulesRequest) {
	request = &DescribeVulEffectModulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEffectModules")
	return
}

func NewDescribeVulEffectModulesResponse() (response *DescribeVulEffectModulesResponse) {
	response = &DescribeVulEffectModulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞影响组件列表
func (c *Client) DescribeVulEffectModules(request *DescribeVulEffectModulesRequest) (response *DescribeVulEffectModulesResponse, err error) {
	if request == nil {
		request = NewDescribeVulEffectModulesRequest()
	}
	response = NewDescribeVulEffectModulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInstallationTokenRequest() (request *DescribeAgentInstallationTokenRequest) {
	request = &DescribeAgentInstallationTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallationToken")
	return
}

func NewDescribeAgentInstallationTokenResponse() (response *DescribeAgentInstallationTokenResponse) {
	response = &DescribeAgentInstallationTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 混合云安装agent token获取
func (c *Client) DescribeAgentInstallationToken(request *DescribeAgentInstallationTokenRequest) (response *DescribeAgentInstallationTokenResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInstallationTokenRequest()
	}
	response = NewDescribeAgentInstallationTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaHostListRequest() (request *DescribeMaHostListRequest) {
	request = &DescribeMaHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaHostList")
	return
}

func NewDescribeMaHostListResponse() (response *DescribeMaHostListResponse) {
	response = &DescribeMaHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警主机列表
func (c *Client) DescribeMaHostList(request *DescribeMaHostListRequest) (response *DescribeMaHostListResponse, err error) {
	if request == nil {
		request = NewDescribeMaHostListRequest()
	}
	response = NewDescribeMaHostListResponse()
	err = c.Send(request, response)
	return
}

func NewSyncMachinesRequest() (request *SyncMachinesRequest) {
	request = &SyncMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SyncMachines")
	return
}

func NewSyncMachinesResponse() (response *SyncMachinesResponse) {
	response = &SyncMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步机器信息
func (c *Client) SyncMachines(request *SyncMachinesRequest) (response *SyncMachinesResponse, err error) {
	if request == nil {
		request = NewSyncMachinesRequest()
	}
	response = NewSyncMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineLoginTypeRequest() (request *DescribeMachineLoginTypeRequest) {
	request = &DescribeMachineLoginTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineLoginType")
	return
}

func NewDescribeMachineLoginTypeResponse() (response *DescribeMachineLoginTypeResponse) {
	response = &DescribeMachineLoginTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机登录方式
func (c *Client) DescribeMachineLoginType(request *DescribeMachineLoginTypeRequest) (response *DescribeMachineLoginTypeResponse, err error) {
	if request == nil {
		request = NewDescribeMachineLoginTypeRequest()
	}
	response = NewDescribeMachineLoginTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellListRequest() (request *DescribeJavaMemShellListRequest) {
	request = &DescribeJavaMemShellListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellList")
	return
}

func NewDescribeJavaMemShellListResponse() (response *DescribeJavaMemShellListResponse) {
	response = &DescribeJavaMemShellListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java内存马事件列表
func (c *Client) DescribeJavaMemShellList(request *DescribeJavaMemShellListRequest) (response *DescribeJavaMemShellListResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellListRequest()
	}
	response = NewDescribeJavaMemShellListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleListRequest() (request *DescribeBaselineRuleListRequest) {
	request = &DescribeBaselineRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleList")
	return
}

func NewDescribeBaselineRuleListResponse() (response *DescribeBaselineRuleListResponse) {
	response = &DescribeBaselineRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线规则列表
func (c *Client) DescribeBaselineRuleList(request *DescribeBaselineRuleListRequest) (response *DescribeBaselineRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleListRequest()
	}
	response = NewDescribeBaselineRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenRiskAssetsTopRequest() (request *DescribeScreenRiskAssetsTopRequest) {
	request = &DescribeScreenRiskAssetsTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenRiskAssetsTop")
	return
}

func NewDescribeScreenRiskAssetsTopResponse() (response *DescribeScreenRiskAssetsTopResponse) {
	response = &DescribeScreenRiskAssetsTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化风险资产top5（今日），统计今日风险资产
func (c *Client) DescribeScreenRiskAssetsTop(request *DescribeScreenRiskAssetsTopRequest) (response *DescribeScreenRiskAssetsTopResponse, err error) {
	if request == nil {
		request = NewDescribeScreenRiskAssetsTopRequest()
	}
	response = NewDescribeScreenRiskAssetsTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESAggregationsRequest() (request *DescribeESAggregationsRequest) {
	request = &DescribeESAggregationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeESAggregations")
	return
}

func NewDescribeESAggregationsResponse() (response *DescribeESAggregationsResponse) {
	response = &DescribeESAggregationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES字段聚合结果
func (c *Client) DescribeESAggregations(request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
	if request == nil {
		request = NewDescribeESAggregationsRequest()
	}
	response = NewDescribeESAggregationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBanStatusRequest() (request *ModifyBanStatusRequest) {
	request = &ModifyBanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanStatus")
	return
}

func NewModifyBanStatusResponse() (response *ModifyBanStatusResponse) {
	response = &ModifyBanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置阻断开关状态
func (c *Client) ModifyBanStatus(request *ModifyBanStatusRequest) (response *ModifyBanStatusResponse, err error) {
	if request == nil {
		request = NewModifyBanStatusRequest()
	}
	response = NewModifyBanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackLogsRequest() (request *DescribeAttackLogsRequest) {
	request = &DescribeAttackLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackLogs")
	return
}

func NewDescribeAttackLogsResponse() (response *DescribeAttackLogsResponse) {
	response = &DescribeAttackLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按分页形式展示网络攻击日志列表
func (c *Client) DescribeAttackLogs(request *DescribeAttackLogsRequest) (response *DescribeAttackLogsResponse, err error) {
	if request == nil {
		request = NewDescribeAttackLogsRequest()
	}
	response = NewDescribeAttackLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgnoreHostAndItemConfigRequest() (request *DescribeIgnoreHostAndItemConfigRequest) {
	request = &DescribeIgnoreHostAndItemConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreHostAndItemConfig")
	return
}

func NewDescribeIgnoreHostAndItemConfigResponse() (response *DescribeIgnoreHostAndItemConfigResponse) {
	response = &DescribeIgnoreHostAndItemConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一键忽略受影响的检测项和主机信息
func (c *Client) DescribeIgnoreHostAndItemConfig(request *DescribeIgnoreHostAndItemConfigRequest) (response *DescribeIgnoreHostAndItemConfigResponse, err error) {
	if request == nil {
		request = NewDescribeIgnoreHostAndItemConfigRequest()
	}
	response = NewDescribeIgnoreHostAndItemConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseMachineStrategyInfoRequest() (request *DescribeRansomDefenseMachineStrategyInfoRequest) {
	request = &DescribeRansomDefenseMachineStrategyInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseMachineStrategyInfo")
	return
}

func NewDescribeRansomDefenseMachineStrategyInfoResponse() (response *DescribeRansomDefenseMachineStrategyInfoResponse) {
	response = &DescribeRansomDefenseMachineStrategyInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机绑定策略列表
func (c *Client) DescribeRansomDefenseMachineStrategyInfo(request *DescribeRansomDefenseMachineStrategyInfoRequest) (response *DescribeRansomDefenseMachineStrategyInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseMachineStrategyInfoRequest()
	}
	response = NewDescribeRansomDefenseMachineStrategyInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportRansomDefenseMachineListRequest() (request *ExportRansomDefenseMachineListRequest) {
	request = &ExportRansomDefenseMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseMachineList")
	return
}

func NewExportRansomDefenseMachineListResponse() (response *ExportRansomDefenseMachineListResponse) {
	response = &ExportRansomDefenseMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出备份详情列表
func (c *Client) ExportRansomDefenseMachineList(request *ExportRansomDefenseMachineListRequest) (response *ExportRansomDefenseMachineListResponse, err error) {
	if request == nil {
		request = NewExportRansomDefenseMachineListRequest()
	}
	response = NewExportRansomDefenseMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebPageProtectDirRequest() (request *ModifyWebPageProtectDirRequest) {
	request = &ModifyWebPageProtectDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectDir")
	return
}

func NewModifyWebPageProtectDirResponse() (response *ModifyWebPageProtectDirResponse) {
	response = &ModifyWebPageProtectDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建/修改网站防护目录
func (c *Client) ModifyWebPageProtectDir(request *ModifyWebPageProtectDirRequest) (response *ModifyWebPageProtectDirResponse, err error) {
	if request == nil {
		request = NewModifyWebPageProtectDirRequest()
	}
	response = NewModifyWebPageProtectDirResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeABTestConfigRequest() (request *DescribeABTestConfigRequest) {
	request = &DescribeABTestConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeABTestConfig")
	return
}

func NewDescribeABTestConfigResponse() (response *DescribeABTestConfigResponse) {
	response = &DescribeABTestConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户当前灰度配置
func (c *Client) DescribeABTestConfig(request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
	if request == nil {
		request = NewDescribeABTestConfigRequest()
	}
	response = NewDescribeABTestConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFastAnalysisRequest() (request *DescribeFastAnalysisRequest) {
	request = &DescribeFastAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFastAnalysis")
	return
}

func NewDescribeFastAnalysisResponse() (response *DescribeFastAnalysisResponse) {
	response = &DescribeFastAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志快速分析统计
func (c *Client) DescribeFastAnalysis(request *DescribeFastAnalysisRequest) (response *DescribeFastAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeFastAnalysisRequest()
	}
	response = NewDescribeFastAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebLocationListRequest() (request *ExportAssetWebLocationListRequest) {
	request = &ExportAssetWebLocationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebLocationList")
	return
}

func NewExportAssetWebLocationListResponse() (response *ExportAssetWebLocationListResponse) {
	response = &ExportAssetWebLocationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出Web站点列表
func (c *Client) ExportAssetWebLocationList(request *ExportAssetWebLocationListRequest) (response *ExportAssetWebLocationListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebLocationListRequest()
	}
	response = NewExportAssetWebLocationListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFileTamperRuleRequest() (request *ModifyFileTamperRuleRequest) {
	request = &ModifyFileTamperRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperRule")
	return
}

func NewModifyFileTamperRuleResponse() (response *ModifyFileTamperRuleResponse) {
	response = &ModifyFileTamperRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑、新增核心文件监控规则
func (c *Client) ModifyFileTamperRule(request *ModifyFileTamperRuleRequest) (response *ModifyFileTamperRuleResponse, err error) {
	if request == nil {
		request = NewModifyFileTamperRuleRequest()
	}
	response = NewModifyFileTamperRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScanMalwareSettingRequest() (request *CreateScanMalwareSettingRequest) {
	request = &CreateScanMalwareSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateScanMalwareSetting")
	return
}

func NewCreateScanMalwareSettingResponse() (response *CreateScanMalwareSettingResponse) {
	response = &CreateScanMalwareSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口可以对入侵检测-文件查杀扫描检测
func (c *Client) CreateScanMalwareSetting(request *CreateScanMalwareSettingRequest) (response *CreateScanMalwareSettingResponse, err error) {
	if request == nil {
		request = NewCreateScanMalwareSettingRequest()
	}
	response = NewCreateScanMalwareSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginWhiteHostListRequest() (request *DescribeLoginWhiteHostListRequest) {
	request = &DescribeLoginWhiteHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteHostList")
	return
}

func NewDescribeLoginWhiteHostListResponse() (response *DescribeLoginWhiteHostListResponse) {
	response = &DescribeLoginWhiteHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合并后白名单机器列表
func (c *Client) DescribeLoginWhiteHostList(request *DescribeLoginWhiteHostListRequest) (response *DescribeLoginWhiteHostListResponse, err error) {
	if request == nil {
		request = NewDescribeLoginWhiteHostListRequest()
	}
	response = NewDescribeLoginWhiteHostListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAttackEventsRequest() (request *ExportAttackEventsRequest) {
	request = &ExportAttackEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAttackEvents")
	return
}

func NewExportAttackEventsResponse() (response *ExportAttackEventsResponse) {
	response = &ExportAttackEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) ExportAttackEvents(request *ExportAttackEventsRequest) (response *ExportAttackEventsResponse, err error) {
	if request == nil {
		request = NewExportAttackEventsRequest()
	}
	response = NewExportAttackEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachineAutoClearConfigRequest() (request *ModifyMachineAutoClearConfigRequest) {
	request = &ModifyMachineAutoClearConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMachineAutoClearConfig")
	return
}

func NewModifyMachineAutoClearConfigResponse() (response *ModifyMachineAutoClearConfigResponse) {
	response = &ModifyMachineAutoClearConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机器清理配置
func (c *Client) ModifyMachineAutoClearConfig(request *ModifyMachineAutoClearConfigRequest) (response *ModifyMachineAutoClearConfigResponse, err error) {
	if request == nil {
		request = NewModifyMachineAutoClearConfigRequest()
	}
	response = NewModifyMachineAutoClearConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineLoginTypeStatusRequest() (request *DescribeMachineLoginTypeStatusRequest) {
	request = &DescribeMachineLoginTypeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineLoginTypeStatus")
	return
}

func NewDescribeMachineLoginTypeStatusResponse() (response *DescribeMachineLoginTypeStatusResponse) {
	response = &DescribeMachineLoginTypeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机登录方式状态
func (c *Client) DescribeMachineLoginTypeStatus(request *DescribeMachineLoginTypeStatusRequest) (response *DescribeMachineLoginTypeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeMachineLoginTypeStatusRequest()
	}
	response = NewDescribeMachineLoginTypeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaWhitelistStatusRequest() (request *ModifyMaWhitelistStatusRequest) {
	request = &ModifyMaWhitelistStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaWhitelistStatus")
	return
}

func NewModifyMaWhitelistStatusResponse() (response *ModifyMaWhitelistStatusResponse) {
	response = &ModifyMaWhitelistStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改微告警全局白名单状态
func (c *Client) ModifyMaWhitelistStatus(request *ModifyMaWhitelistStatusRequest) (response *ModifyMaWhitelistStatusResponse, err error) {
	if request == nil {
		request = NewModifyMaWhitelistStatusRequest()
	}
	response = NewModifyMaWhitelistStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
	request = &DescribeAccountStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAccountStatistics")
	return
}

func NewDescribeAccountStatisticsResponse() (response *DescribeAccountStatisticsResponse) {
	response = &DescribeAccountStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAccountStatistics) 用于获取帐号统计列表数据。
func (c *Client) DescribeAccountStatistics(request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAccountStatisticsRequest()
	}
	response = NewDescribeAccountStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewGetLocalStorageItemRequest() (request *GetLocalStorageItemRequest) {
	request = &GetLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "GetLocalStorageItem")
	return
}

func NewGetLocalStorageItemResponse() (response *GetLocalStorageItemResponse) {
	response = &GetLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地存储数据
func (c *Client) GetLocalStorageItem(request *GetLocalStorageItemRequest) (response *GetLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewGetLocalStorageItemRequest()
	}
	response = NewGetLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
	request = &DeleteMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachine")
	return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
	response = &DeleteMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteMachine）用于卸载云镜客户端。
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
	if request == nil {
		request = NewDeleteMachineRequest()
	}
	response = NewDeleteMachineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleIgnoreListRequest() (request *DescribeBaselineRuleIgnoreListRequest) {
	request = &DescribeBaselineRuleIgnoreListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleIgnoreList")
	return
}

func NewDescribeBaselineRuleIgnoreListResponse() (response *DescribeBaselineRuleIgnoreListResponse) {
	response = &DescribeBaselineRuleIgnoreListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线忽略规则列表
func (c *Client) DescribeBaselineRuleIgnoreList(request *DescribeBaselineRuleIgnoreListRequest) (response *DescribeBaselineRuleIgnoreListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleIgnoreListRequest()
	}
	response = NewDescribeBaselineRuleIgnoreListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaCustomPolicyStatusRequest() (request *ModifyMaCustomPolicyStatusRequest) {
	request = &ModifyMaCustomPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaCustomPolicyStatus")
	return
}

func NewModifyMaCustomPolicyStatusResponse() (response *ModifyMaCustomPolicyStatusResponse) {
	response = &ModifyMaCustomPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更微告警自定义白名单状态
func (c *Client) ModifyMaCustomPolicyStatus(request *ModifyMaCustomPolicyStatusRequest) (response *ModifyMaCustomPolicyStatusResponse, err error) {
	if request == nil {
		request = NewModifyMaCustomPolicyStatusRequest()
	}
	response = NewModifyMaCustomPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWarningListRequest() (request *DescribeWarningListRequest) {
	request = &DescribeWarningListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWarningList")
	return
}

func NewDescribeWarningListResponse() (response *DescribeWarningListResponse) {
	response = &DescribeWarningListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前用户告警列表
func (c *Client) DescribeWarningList(request *DescribeWarningListRequest) (response *DescribeWarningListResponse, err error) {
	if request == nil {
		request = NewDescribeWarningListRequest()
	}
	response = NewDescribeWarningListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetInitServiceListRequest() (request *ExportAssetInitServiceListRequest) {
	request = &ExportAssetInitServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetInitServiceList")
	return
}

func NewExportAssetInitServiceListResponse() (response *ExportAssetInitServiceListResponse) {
	response = &ExportAssetInitServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理启动服务列表
func (c *Client) ExportAssetInitServiceList(request *ExportAssetInitServiceListRequest) (response *ExportAssetInitServiceListResponse, err error) {
	if request == nil {
		request = NewExportAssetInitServiceListRequest()
	}
	response = NewExportAssetInitServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebHookPolicyStatusRequest() (request *ModifyWebHookPolicyStatusRequest) {
	request = &ModifyWebHookPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookPolicyStatus")
	return
}

func NewModifyWebHookPolicyStatusResponse() (response *ModifyWebHookPolicyStatusResponse) {
	response = &ModifyWebHookPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警策略开关
func (c *Client) ModifyWebHookPolicyStatus(request *ModifyWebHookPolicyStatusRequest) (response *ModifyWebHookPolicyStatusResponse, err error) {
	if request == nil {
		request = NewModifyWebHookPolicyStatusRequest()
	}
	response = NewModifyWebHookPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrderListRequest() (request *DescribeOrderListRequest) {
	request = &DescribeOrderListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeOrderList")
	return
}

func NewDescribeOrderListResponse() (response *DescribeOrderListResponse) {
	response = &DescribeOrderListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单列表
func (c *Client) DescribeOrderList(request *DescribeOrderListRequest) (response *DescribeOrderListResponse, err error) {
	if request == nil {
		request = NewDescribeOrderListRequest()
	}
	response = NewDescribeOrderListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineInfoRequest() (request *DescribeMachineInfoRequest) {
	request = &DescribeMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineInfo")
	return
}

func NewDescribeMachineInfoResponse() (response *DescribeMachineInfoResponse) {
	response = &DescribeMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeMachineInfo）用于获取机器详细信息。
func (c *Client) DescribeMachineInfo(request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
	if request == nil {
		request = NewDescribeMachineInfoRequest()
	}
	response = NewDescribeMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginWhiteListNewRequest() (request *DescribeLoginWhiteListNewRequest) {
	request = &DescribeLoginWhiteListNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteListNew")
	return
}

func NewDescribeLoginWhiteListNewResponse() (response *DescribeLoginWhiteListNewResponse) {
	response = &DescribeLoginWhiteListNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异地登录白名单列表支持新版筛选
func (c *Client) DescribeLoginWhiteListNew(request *DescribeLoginWhiteListNewRequest) (response *DescribeLoginWhiteListNewResponse, err error) {
	if request == nil {
		request = NewDescribeLoginWhiteListNewRequest()
	}
	response = NewDescribeLoginWhiteListNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceSettingRequest() (request *DescribeVulDefenceSettingRequest) {
	request = &DescribeVulDefenceSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceSetting")
	return
}

func NewDescribeVulDefenceSettingResponse() (response *DescribeVulDefenceSettingResponse) {
	response = &DescribeVulDefenceSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前漏洞防御插件设置
func (c *Client) DescribeVulDefenceSetting(request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceSettingRequest()
	}
	response = NewDescribeVulDefenceSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulHostCountScanTimeRequest() (request *DescribeVulHostCountScanTimeRequest) {
	request = &DescribeVulHostCountScanTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulHostCountScanTime")
	return
}

func NewDescribeVulHostCountScanTimeResponse() (response *DescribeVulHostCountScanTimeResponse) {
	response = &DescribeVulHostCountScanTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取待处理漏洞数+影响主机数
func (c *Client) DescribeVulHostCountScanTime(request *DescribeVulHostCountScanTimeRequest) (response *DescribeVulHostCountScanTimeResponse, err error) {
	if request == nil {
		request = NewDescribeVulHostCountScanTimeRequest()
	}
	response = NewDescribeVulHostCountScanTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanVulSettingRequest() (request *DescribeScanVulSettingRequest) {
	request = &DescribeScanVulSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanVulSetting")
	return
}

func NewDescribeScanVulSettingResponse() (response *DescribeScanVulSettingResponse) {
	response = &DescribeScanVulSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询定期检测的配置
func (c *Client) DescribeScanVulSetting(request *DescribeScanVulSettingRequest) (response *DescribeScanVulSettingResponse, err error) {
	if request == nil {
		request = NewDescribeScanVulSettingRequest()
	}
	response = NewDescribeScanVulSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerRelatedDirInfoRequest() (request *DescribeServerRelatedDirInfoRequest) {
	request = &DescribeServerRelatedDirInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeServerRelatedDirInfo")
	return
}

func NewDescribeServerRelatedDirInfoResponse() (response *DescribeServerRelatedDirInfoResponse) {
	response = &DescribeServerRelatedDirInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务区关联目录详情
func (c *Client) DescribeServerRelatedDirInfo(request *DescribeServerRelatedDirInfoRequest) (response *DescribeServerRelatedDirInfoResponse, err error) {
	if request == nil {
		request = NewDescribeServerRelatedDirInfoRequest()
	}
	response = NewDescribeServerRelatedDirInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFileTamperRuleStatusRequest() (request *ModifyFileTamperRuleStatusRequest) {
	request = &ModifyFileTamperRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperRuleStatus")
	return
}

func NewModifyFileTamperRuleStatusResponse() (response *ModifyFileTamperRuleStatusResponse) {
	response = &ModifyFileTamperRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 核心文件规则状态更新，支持批量删除 关闭
func (c *Client) ModifyFileTamperRuleStatus(request *ModifyFileTamperRuleStatusRequest) (response *ModifyFileTamperRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyFileTamperRuleStatusRequest()
	}
	response = NewModifyFileTamperRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProtectDirRequest() (request *DeleteProtectDirRequest) {
	request = &DeleteProtectDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteProtectDir")
	return
}

func NewDeleteProtectDirResponse() (response *DeleteProtectDirResponse) {
	response = &DeleteProtectDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除防护网站
func (c *Client) DeleteProtectDir(request *DeleteProtectDirRequest) (response *DeleteProtectDirResponse, err error) {
	if request == nil {
		request = NewDeleteProtectDirRequest()
	}
	response = NewDeleteProtectDirResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellRulesAggregationRequest() (request *DescribeReverseShellRulesAggregationRequest) {
	request = &DescribeReverseShellRulesAggregationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellRulesAggregation")
	return
}

func NewDescribeReverseShellRulesAggregationResponse() (response *DescribeReverseShellRulesAggregationResponse) {
	response = &DescribeReverseShellRulesAggregationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取反弹Shell规则列表-聚合版本
func (c *Client) DescribeReverseShellRulesAggregation(request *DescribeReverseShellRulesAggregationRequest) (response *DescribeReverseShellRulesAggregationResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellRulesAggregationRequest()
	}
	response = NewDescribeReverseShellRulesAggregationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceListRequest() (request *DescribeVulDefenceListRequest) {
	request = &DescribeVulDefenceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceList")
	return
}

func NewDescribeVulDefenceListResponse() (response *DescribeVulDefenceListResponse) {
	response = &DescribeVulDefenceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御列表
func (c *Client) DescribeVulDefenceList(request *DescribeVulDefenceListRequest) (response *DescribeVulDefenceListResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceListRequest()
	}
	response = NewDescribeVulDefenceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStorageConfigRequest() (request *DescribeLogStorageConfigRequest) {
	request = &DescribeLogStorageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageConfig")
	return
}

func NewDescribeLogStorageConfigResponse() (response *DescribeLogStorageConfigResponse) {
	response = &DescribeLogStorageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志存储配置
func (c *Client) DescribeLogStorageConfig(request *DescribeLogStorageConfigRequest) (response *DescribeLogStorageConfigResponse, err error) {
	if request == nil {
		request = NewDescribeLogStorageConfigRequest()
	}
	response = NewDescribeLogStorageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineInfoByHashKeyRequest() (request *DescribeMachineInfoByHashKeyRequest) {
	request = &DescribeMachineInfoByHashKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineInfoByHashKey")
	return
}

func NewDescribeMachineInfoByHashKeyResponse() (response *DescribeMachineInfoByHashKeyResponse) {
	response = &DescribeMachineInfoByHashKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器信息
func (c *Client) DescribeMachineInfoByHashKey(request *DescribeMachineInfoByHashKeyRequest) (response *DescribeMachineInfoByHashKeyResponse, err error) {
	if request == nil {
		request = NewDescribeMachineInfoByHashKeyRequest()
	}
	response = NewDescribeMachineInfoByHashKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineFixListRequest() (request *DescribeBaselineFixListRequest) {
	request = &DescribeBaselineFixListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineFixList")
	return
}

func NewDescribeBaselineFixListResponse() (response *DescribeBaselineFixListResponse) {
	response = &DescribeBaselineFixListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线修复列表
func (c *Client) DescribeBaselineFixList(request *DescribeBaselineFixListRequest) (response *DescribeBaselineFixListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineFixListRequest()
	}
	response = NewDescribeBaselineFixListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVdbAndPocInfoRequest() (request *DescribeVdbAndPocInfoRequest) {
	request = &DescribeVdbAndPocInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVdbAndPocInfo")
	return
}

func NewDescribeVdbAndPocInfoResponse() (response *DescribeVdbAndPocInfoResponse) {
	response = &DescribeVdbAndPocInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取病毒库及POC的更新信息
func (c *Client) DescribeVdbAndPocInfo(request *DescribeVdbAndPocInfoRequest) (response *DescribeVdbAndPocInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVdbAndPocInfoRequest()
	}
	response = NewDescribeVdbAndPocInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDisableSelfLearnRequest() (request *DisableSelfLearnRequest) {
	request = &DisableSelfLearnRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DisableSelfLearn")
	return
}

func NewDisableSelfLearnResponse() (response *DisableSelfLearnResponse) {
	response = &DisableSelfLearnResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭微告警自学习
func (c *Client) DisableSelfLearn(request *DisableSelfLearnRequest) (response *DisableSelfLearnResponse, err error) {
	if request == nil {
		request = NewDisableSelfLearnRequest()
	}
	response = NewDisableSelfLearnResponse()
	err = c.Send(request, response)
	return
}

func NewStopAssetScanRequest() (request *StopAssetScanRequest) {
	request = &StopAssetScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "StopAssetScan")
	return
}

func NewStopAssetScanResponse() (response *StopAssetScanResponse) {
	response = &StopAssetScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止资产扫描任务
func (c *Client) StopAssetScan(request *StopAssetScanRequest) (response *StopAssetScanResponse, err error) {
	if request == nil {
		request = NewStopAssetScanRequest()
	}
	response = NewStopAssetScanResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetScanRequest() (request *SyncAssetScanRequest) {
	request = &SyncAssetScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SyncAssetScan")
	return
}

func NewSyncAssetScanResponse() (response *SyncAssetScanResponse) {
	response = &SyncAssetScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步资产扫描信息
func (c *Client) SyncAssetScan(request *SyncAssetScanRequest) (response *SyncAssetScanResponse, err error) {
	if request == nil {
		request = NewSyncAssetScanRequest()
	}
	response = NewSyncAssetScanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulListRequest() (request *DescribeVulListRequest) {
	request = &DescribeVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulList")
	return
}

func NewDescribeVulListResponse() (response *DescribeVulListResponse) {
	response = &DescribeVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞列表数据
func (c *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
	if request == nil {
		request = NewDescribeVulListRequest()
	}
	response = NewDescribeVulListResponse()
	err = c.Send(request, response)
	return
}

func NewExportMalwaresRequest() (request *ExportMalwaresRequest) {
	request = &ExportMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMalwares")
	return
}

func NewExportMalwaresResponse() (response *ExportMalwaresResponse) {
	response = &ExportMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ExportMalwares) 用于导出木马记录CSV文件。
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
	if request == nil {
		request = NewExportMalwaresRequest()
	}
	response = NewExportMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineStrategyDetailRequest() (request *DescribeBaselineStrategyDetailRequest) {
	request = &DescribeBaselineStrategyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineStrategyDetail")
	return
}

func NewDescribeBaselineStrategyDetailResponse() (response *DescribeBaselineStrategyDetailResponse) {
	response = &DescribeBaselineStrategyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线策略id查询策略详情
func (c *Client) DescribeBaselineStrategyDetail(request *DescribeBaselineStrategyDetailRequest) (response *DescribeBaselineStrategyDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineStrategyDetailRequest()
	}
	response = NewDescribeBaselineStrategyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefencePluginExceptionCountRequest() (request *DescribeVulDefencePluginExceptionCountRequest) {
	request = &DescribeVulDefencePluginExceptionCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginExceptionCount")
	return
}

func NewDescribeVulDefencePluginExceptionCountResponse() (response *DescribeVulDefencePluginExceptionCountResponse) {
	response = &DescribeVulDefencePluginExceptionCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前异常插件数
func (c *Client) DescribeVulDefencePluginExceptionCount(request *DescribeVulDefencePluginExceptionCountRequest) (response *DescribeVulDefencePluginExceptionCountResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefencePluginExceptionCountRequest()
	}
	response = NewDescribeVulDefencePluginExceptionCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulEffectHostListRequest() (request *DescribeVulEffectHostListRequest) {
	request = &DescribeVulEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEffectHostList")
	return
}

func NewDescribeVulEffectHostListResponse() (response *DescribeVulEffectHostListResponse) {
	response = &DescribeVulEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞影响主机列表
func (c *Client) DescribeVulEffectHostList(request *DescribeVulEffectHostListRequest) (response *DescribeVulEffectHostListResponse, err error) {
	if request == nil {
		request = NewDescribeVulEffectHostListRequest()
	}
	response = NewDescribeVulEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetProcessInfoListRequest() (request *ExportAssetProcessInfoListRequest) {
	request = &ExportAssetProcessInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetProcessInfoList")
	return
}

func NewExportAssetProcessInfoListResponse() (response *ExportAssetProcessInfoListResponse) {
	response = &ExportAssetProcessInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理进程列表
func (c *Client) ExportAssetProcessInfoList(request *ExportAssetProcessInfoListRequest) (response *ExportAssetProcessInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetProcessInfoListRequest()
	}
	response = NewExportAssetProcessInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCloudVendorAccountAKRequest() (request *DeleteCloudVendorAccountAKRequest) {
	request = &DeleteCloudVendorAccountAKRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteCloudVendorAccountAK")
	return
}

func NewDeleteCloudVendorAccountAKResponse() (response *DeleteCloudVendorAccountAKResponse) {
	response = &DeleteCloudVendorAccountAKResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除AK记录
func (c *Client) DeleteCloudVendorAccountAK(request *DeleteCloudVendorAccountAKRequest) (response *DeleteCloudVendorAccountAKResponse, err error) {
	if request == nil {
		request = NewDeleteCloudVendorAccountAKRequest()
	}
	response = NewDeleteCloudVendorAccountAKResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineAnalysisDataRequest() (request *DescribeBaselineAnalysisDataRequest) {
	request = &DescribeBaselineAnalysisDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineAnalysisData")
	return
}

func NewDescribeBaselineAnalysisDataResponse() (response *DescribeBaselineAnalysisDataResponse) {
	response = &DescribeBaselineAnalysisDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线策略id查询基线策略数据概览统计
func (c *Client) DescribeBaselineAnalysisData(request *DescribeBaselineAnalysisDataRequest) (response *DescribeBaselineAnalysisDataResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineAnalysisDataRequest()
	}
	response = NewDescribeBaselineAnalysisDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsersConfigRequest() (request *DescribeUsersConfigRequest) {
	request = &DescribeUsersConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeUsersConfig")
	return
}

func NewDescribeUsersConfigResponse() (response *DescribeUsersConfigResponse) {
	response = &DescribeUsersConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户配置
func (c *Client) DescribeUsersConfig(request *DescribeUsersConfigRequest) (response *DescribeUsersConfigResponse, err error) {
	if request == nil {
		request = NewDescribeUsersConfigRequest()
	}
	response = NewDescribeUsersConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalWareListRequest() (request *DescribeMalWareListRequest) {
	request = &DescribeMalWareListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalWareList")
	return
}

func NewDescribeMalWareListResponse() (response *DescribeMalWareListResponse) {
	response = &DescribeMalWareListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵检测获取木马列表
func (c *Client) DescribeMalWareList(request *DescribeMalWareListRequest) (response *DescribeMalWareListResponse, err error) {
	if request == nil {
		request = NewDescribeMalWareListRequest()
	}
	response = NewDescribeMalWareListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseStrategyDetailRequest() (request *DescribeRansomDefenseStrategyDetailRequest) {
	request = &DescribeRansomDefenseStrategyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyDetail")
	return
}

func NewDescribeRansomDefenseStrategyDetailResponse() (response *DescribeRansomDefenseStrategyDetailResponse) {
	response = &DescribeRansomDefenseStrategyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取策略详情
func (c *Client) DescribeRansomDefenseStrategyDetail(request *DescribeRansomDefenseStrategyDetailRequest) (response *DescribeRansomDefenseStrategyDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseStrategyDetailRequest()
	}
	response = NewDescribeRansomDefenseStrategyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询价格
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	response = NewDescribePriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUndoVulCountsRequest() (request *DescribeUndoVulCountsRequest) {
	request = &DescribeUndoVulCountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeUndoVulCounts")
	return
}

func NewDescribeUndoVulCountsResponse() (response *DescribeUndoVulCountsResponse) {
	response = &DescribeUndoVulCountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞管理模块指定类型的待处理漏洞数、主机数和非专业版主机数量
func (c *Client) DescribeUndoVulCounts(request *DescribeUndoVulCountsRequest) (response *DescribeUndoVulCountsResponse, err error) {
	if request == nil {
		request = NewDescribeUndoVulCountsRequest()
	}
	response = NewDescribeUndoVulCountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetInitServiceListRequest() (request *DescribeAssetInitServiceListRequest) {
	request = &DescribeAssetInitServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetInitServiceList")
	return
}

func NewDescribeAssetInitServiceListResponse() (response *DescribeAssetInitServiceListResponse) {
	response = &DescribeAssetInitServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理启动服务列表
func (c *Client) DescribeAssetInitServiceList(request *DescribeAssetInitServiceListRequest) (response *DescribeAssetInitServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetInitServiceListRequest()
	}
	response = NewDescribeAssetInitServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBanWhiteListRequest() (request *DescribeBanWhiteListRequest) {
	request = &DescribeBanWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanWhiteList")
	return
}

func NewDescribeBanWhiteListResponse() (response *DescribeBanWhiteListResponse) {
	response = &DescribeBanWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取阻断白名单列表
func (c *Client) DescribeBanWhiteList(request *DescribeBanWhiteListRequest) (response *DescribeBanWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeBanWhiteListRequest()
	}
	response = NewDescribeBanWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrderRequest() (request *ModifyOrderRequest) {
	request = &ModifyOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyOrder")
	return
}

func NewModifyOrderResponse() (response *ModifyOrderResponse) {
	response = &ModifyOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑订单
func (c *Client) ModifyOrder(request *ModifyOrderRequest) (response *ModifyOrderResponse, err error) {
	if request == nil {
		request = NewModifyOrderRequest()
	}
	response = NewModifyOrderResponse()
	err = c.Send(request, response)
	return
}

func NewOpenProVersionQuickRequest() (request *OpenProVersionQuickRequest) {
	request = &OpenProVersionQuickRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "OpenProVersionQuick")
	return
}

func NewOpenProVersionQuickResponse() (response *OpenProVersionQuickResponse) {
	response = &OpenProVersionQuickResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (OpenProVersionQuick) 用于快速开通所有主机为专业版。
func (c *Client) OpenProVersionQuick(request *OpenProVersionQuickRequest) (response *OpenProVersionQuickResponse, err error) {
	if request == nil {
		request = NewOpenProVersionQuickRequest()
	}
	response = NewOpenProVersionQuickResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBashRulesRequest() (request *DeleteBashRulesRequest) {
	request = &DeleteBashRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashRules")
	return
}

func NewDeleteBashRulesResponse() (response *DeleteBashRulesResponse) {
	response = &DeleteBashRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高危命令规则
func (c *Client) DeleteBashRules(request *DeleteBashRulesRequest) (response *DeleteBashRulesResponse, err error) {
	if request == nil {
		request = NewDeleteBashRulesRequest()
	}
	response = NewDeleteBashRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEmergencyResponseListRequest() (request *DescribeEmergencyResponseListRequest) {
	request = &DescribeEmergencyResponseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeEmergencyResponseList")
	return
}

func NewDescribeEmergencyResponseListResponse() (response *DescribeEmergencyResponseListResponse) {
	response = &DescribeEmergencyResponseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专家服务-应急响应列表
func (c *Client) DescribeEmergencyResponseList(request *DescribeEmergencyResponseListRequest) (response *DescribeEmergencyResponseListResponse, err error) {
	if request == nil {
		request = NewDescribeEmergencyResponseListRequest()
	}
	response = NewDescribeEmergencyResponseListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrivilegeRulesRequest() (request *DescribePrivilegeRulesRequest) {
	request = &DescribePrivilegeRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeRules")
	return
}

func NewDescribePrivilegeRulesResponse() (response *DescribePrivilegeRulesResponse) {
	response = &DescribePrivilegeRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地提权规则列表
func (c *Client) DescribePrivilegeRules(request *DescribePrivilegeRulesRequest) (response *DescribePrivilegeRulesResponse, err error) {
	if request == nil {
		request = NewDescribePrivilegeRulesRequest()
	}
	response = NewDescribePrivilegeRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaSelfLearnConfigRequest() (request *ModifyMaSelfLearnConfigRequest) {
	request = &ModifyMaSelfLearnConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaSelfLearnConfig")
	return
}

func NewModifyMaSelfLearnConfigResponse() (response *ModifyMaSelfLearnConfigResponse) {
	response = &ModifyMaSelfLearnConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置自学习策略
func (c *Client) ModifyMaSelfLearnConfig(request *ModifyMaSelfLearnConfigRequest) (response *ModifyMaSelfLearnConfigResponse, err error) {
	if request == nil {
		request = NewModifyMaSelfLearnConfigRequest()
	}
	response = NewModifyMaSelfLearnConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachinesSimpleRequest() (request *DescribeMachinesSimpleRequest) {
	request = &DescribeMachinesSimpleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachinesSimple")
	return
}

func NewDescribeMachinesSimpleResponse() (response *DescribeMachinesSimpleResponse) {
	response = &DescribeMachinesSimpleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeMachinesSimple) 用于获取主机列表。
func (c *Client) DescribeMachinesSimple(request *DescribeMachinesSimpleRequest) (response *DescribeMachinesSimpleResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesSimpleRequest()
	}
	response = NewDescribeMachinesSimpleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppListRequest() (request *DescribeAssetAppListRequest) {
	request = &DescribeAssetAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppList")
	return
}

func NewDescribeAssetAppListResponse() (response *DescribeAssetAppListResponse) {
	response = &DescribeAssetAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用列表
func (c *Client) DescribeAssetAppList(request *DescribeAssetAppListRequest) (response *DescribeAssetAppListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppListRequest()
	}
	response = NewDescribeAssetAppListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareFileRequest() (request *DescribeMalwareFileRequest) {
	request = &DescribeMalwareFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareFile")
	return
}

func NewDescribeMalwareFileResponse() (response *DescribeMalwareFileResponse) {
	response = &DescribeMalwareFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取木马文件下载地址
func (c *Client) DescribeMalwareFile(request *DescribeMalwareFileRequest) (response *DescribeMalwareFileResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareFileRequest()
	}
	response = NewDescribeMalwareFileResponse()
	err = c.Send(request, response)
	return
}

func NewExportRansomDefenseStrategyListRequest() (request *ExportRansomDefenseStrategyListRequest) {
	request = &ExportRansomDefenseStrategyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseStrategyList")
	return
}

func NewExportRansomDefenseStrategyListResponse() (response *ExportRansomDefenseStrategyListResponse) {
	response = &ExportRansomDefenseStrategyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出防勒索策略列表
func (c *Client) ExportRansomDefenseStrategyList(request *ExportRansomDefenseStrategyListRequest) (response *ExportRansomDefenseStrategyListResponse, err error) {
	if request == nil {
		request = NewExportRansomDefenseStrategyListRequest()
	}
	response = NewExportRansomDefenseStrategyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttackSourceRequest() (request *DescribeAttackSourceRequest) {
	request = &DescribeAttackSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackSource")
	return
}

func NewDescribeAttackSourceResponse() (response *DescribeAttackSourceResponse) {
	response = &DescribeAttackSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询攻击溯源
func (c *Client) DescribeAttackSource(request *DescribeAttackSourceRequest) (response *DescribeAttackSourceResponse, err error) {
	if request == nil {
		request = NewDescribeAttackSourceRequest()
	}
	response = NewDescribeAttackSourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineScanScheduleRequest() (request *DescribeBaselineScanScheduleRequest) {
	request = &DescribeBaselineScanScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineScanSchedule")
	return
}

func NewDescribeBaselineScanScheduleResponse() (response *DescribeBaselineScanScheduleResponse) {
	response = &DescribeBaselineScanScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务id查询基线检测进度
func (c *Client) DescribeBaselineScanSchedule(request *DescribeBaselineScanScheduleRequest) (response *DescribeBaselineScanScheduleResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineScanScheduleRequest()
	}
	response = NewDescribeBaselineScanScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseRequest() (request *DescribeLicenseRequest) {
	request = &DescribeLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicense")
	return
}

func NewDescribeLicenseResponse() (response *DescribeLicenseResponse) {
	response = &DescribeLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权信息
func (c *Client) DescribeLicense(request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseRequest()
	}
	response = NewDescribeLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareInfoRequest() (request *DescribeMalwareInfoRequest) {
	request = &DescribeMalwareInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareInfo")
	return
}

func NewDescribeMalwareInfoResponse() (response *DescribeMalwareInfoResponse) {
	response = &DescribeMalwareInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看恶意文件详情
func (c *Client) DescribeMalwareInfo(request *DescribeMalwareInfoRequest) (response *DescribeMalwareInfoResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareInfoRequest()
	}
	response = NewDescribeMalwareInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetScanStatusRequest() (request *DescribeAssetScanStatusRequest) {
	request = &DescribeAssetScanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetScanStatus")
	return
}

func NewDescribeAssetScanStatusResponse() (response *DescribeAssetScanStatusResponse) {
	response = &DescribeAssetScanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产扫描状态
func (c *Client) DescribeAssetScanStatus(request *DescribeAssetScanStatusRequest) (response *DescribeAssetScanStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetScanStatusRequest()
	}
	response = NewDescribeAssetScanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineDetectListRequest() (request *DescribeBaselineDetectListRequest) {
	request = &DescribeBaselineDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetectList")
	return
}

func NewDescribeBaselineDetectListResponse() (response *DescribeBaselineDetectListResponse) {
	response = &DescribeBaselineDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测详情记录
func (c *Client) DescribeBaselineDetectList(request *DescribeBaselineDetectListRequest) (response *DescribeBaselineDetectListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineDetectListRequest()
	}
	response = NewDescribeBaselineDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewAddLoginWhiteListsRequest() (request *AddLoginWhiteListsRequest) {
	request = &AddLoginWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "AddLoginWhiteLists")
	return
}

func NewAddLoginWhiteListsResponse() (response *AddLoginWhiteListsResponse) {
	response = &AddLoginWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加异地登录白名单
func (c *Client) AddLoginWhiteLists(request *AddLoginWhiteListsRequest) (response *AddLoginWhiteListsResponse, err error) {
	if request == nil {
		request = NewAddLoginWhiteListsRequest()
	}
	response = NewAddLoginWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareRiskWarningRequest() (request *DescribeMalwareRiskWarningRequest) {
	request = &DescribeMalwareRiskWarningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareRiskWarning")
	return
}

func NewDescribeMalwareRiskWarningResponse() (response *DescribeMalwareRiskWarningResponse) {
	response = &DescribeMalwareRiskWarningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开入侵检测-恶意文件检测,弹出风险预警内容
func (c *Client) DescribeMalwareRiskWarning(request *DescribeMalwareRiskWarningRequest) (response *DescribeMalwareRiskWarningResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareRiskWarningRequest()
	}
	response = NewDescribeMalwareRiskWarningResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseTrendRequest() (request *DescribeRansomDefenseTrendRequest) {
	request = &DescribeRansomDefenseTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseTrend")
	return
}

func NewDescribeRansomDefenseTrendResponse() (response *DescribeRansomDefenseTrendResponse) {
	response = &DescribeRansomDefenseTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全网勒索态势
func (c *Client) DescribeRansomDefenseTrend(request *DescribeRansomDefenseTrendRequest) (response *DescribeRansomDefenseTrendResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseTrendRequest()
	}
	response = NewDescribeRansomDefenseTrendResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
	request = &ModifyAutoOpenProVersionConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyAutoOpenProVersionConfig")
	return
}

func NewModifyAutoOpenProVersionConfigResponse() (response *ModifyAutoOpenProVersionConfigResponse) {
	response = &ModifyAutoOpenProVersionConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于设置新增主机自动开通专业防护配置。
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
	if request == nil {
		request = NewModifyAutoOpenProVersionConfigRequest()
	}
	response = NewModifyAutoOpenProVersionConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaWhitelistRequest() (request *DeleteMaWhitelistRequest) {
	request = &DeleteMaWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaWhitelist")
	return
}

func NewDeleteMaWhitelistResponse() (response *DeleteMaWhitelistResponse) {
	response = &DeleteMaWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微告警全局白名单
func (c *Client) DeleteMaWhitelist(request *DeleteMaWhitelistRequest) (response *DeleteMaWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteMaWhitelistRequest()
	}
	response = NewDeleteMaWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaConnEventRequest() (request *DeleteMaConnEventRequest) {
	request = &DeleteMaConnEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaConnEvent")
	return
}

func NewDeleteMaConnEventResponse() (response *DeleteMaConnEventResponse) {
	response = &DeleteMaConnEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微告警访问事件
func (c *Client) DeleteMaConnEvent(request *DeleteMaConnEventRequest) (response *DeleteMaConnEventResponse, err error) {
	if request == nil {
		request = NewDeleteMaConnEventRequest()
	}
	response = NewDeleteMaConnEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInstallResultRequest() (request *DescribeAgentInstallResultRequest) {
	request = &DescribeAgentInstallResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallResult")
	return
}

func NewDescribeAgentInstallResultResponse() (response *DescribeAgentInstallResultResponse) {
	response = &DescribeAgentInstallResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器安装结果
func (c *Client) DescribeAgentInstallResult(request *DescribeAgentInstallResultRequest) (response *DescribeAgentInstallResultResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInstallResultRequest()
	}
	response = NewDescribeAgentInstallResultResponse()
	err = c.Send(request, response)
	return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
	request = &EditTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "EditTags")
	return
}

func NewEditTagsResponse() (response *EditTagsResponse) {
	response = &EditTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或编辑标签
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
	if request == nil {
		request = NewEditTagsRequest()
	}
	response = NewEditTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaPodListRequest() (request *DescribeMaPodListRequest) {
	request = &DescribeMaPodListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaPodList")
	return
}

func NewDescribeMaPodListResponse() (response *DescribeMaPodListResponse) {
	response = &DescribeMaPodListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警POD列表
func (c *Client) DescribeMaPodList(request *DescribeMaPodListRequest) (response *DescribeMaPodListResponse, err error) {
	if request == nil {
		request = NewDescribeMaPodListRequest()
	}
	response = NewDescribeMaPodListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRansomDefenseStrategyStatusRequest() (request *ModifyRansomDefenseStrategyStatusRequest) {
	request = &ModifyRansomDefenseStrategyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyRansomDefenseStrategyStatus")
	return
}

func NewModifyRansomDefenseStrategyStatusResponse() (response *ModifyRansomDefenseStrategyStatusResponse) {
	response = &ModifyRansomDefenseStrategyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改防勒索策略状态
func (c *Client) ModifyRansomDefenseStrategyStatus(request *ModifyRansomDefenseStrategyStatusRequest) (response *ModifyRansomDefenseStrategyStatusResponse, err error) {
	if request == nil {
		request = NewModifyRansomDefenseStrategyStatusRequest()
	}
	response = NewModifyRansomDefenseStrategyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceProcessListRequest() (request *DescribeAssetWebServiceProcessListRequest) {
	request = &DescribeAssetWebServiceProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceProcessList")
	return
}

func NewDescribeAssetWebServiceProcessListResponse() (response *DescribeAssetWebServiceProcessListResponse) {
	response = &DescribeAssetWebServiceProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web服务关联进程列表
func (c *Client) DescribeAssetWebServiceProcessList(request *DescribeAssetWebServiceProcessListRequest) (response *DescribeAssetWebServiceProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceProcessListRequest()
	}
	response = NewDescribeAssetWebServiceProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogDeliveryKafkaOptionsRequest() (request *DescribeLogDeliveryKafkaOptionsRequest) {
	request = &DescribeLogDeliveryKafkaOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogDeliveryKafkaOptions")
	return
}

func NewDescribeLogDeliveryKafkaOptionsResponse() (response *DescribeLogDeliveryKafkaOptionsResponse) {
	response = &DescribeLogDeliveryKafkaOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志投递kafka可选项列表
func (c *Client) DescribeLogDeliveryKafkaOptions(request *DescribeLogDeliveryKafkaOptionsRequest) (response *DescribeLogDeliveryKafkaOptionsResponse, err error) {
	if request == nil {
		request = NewDescribeLogDeliveryKafkaOptionsRequest()
	}
	response = NewDescribeLogDeliveryKafkaOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashPoliciesRequest() (request *DescribeBashPoliciesRequest) {
	request = &DescribeBashPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashPolicies")
	return
}

func NewDescribeBashPoliciesResponse() (response *DescribeBashPoliciesResponse) {
	response = &DescribeBashPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高危命令策略列表
func (c *Client) DescribeBashPolicies(request *DescribeBashPoliciesRequest) (response *DescribeBashPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeBashPoliciesRequest()
	}
	response = NewDescribeBashPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetAppListRequest() (request *ExportAssetAppListRequest) {
	request = &ExportAssetAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetAppList")
	return
}

func NewExportAssetAppListResponse() (response *ExportAssetAppListResponse) {
	response = &ExportAssetAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理应用列表
func (c *Client) ExportAssetAppList(request *ExportAssetAppListRequest) (response *ExportAssetAppListResponse, err error) {
	if request == nil {
		request = NewExportAssetAppListRequest()
	}
	response = NewExportAssetAppListResponse()
	err = c.Send(request, response)
	return
}

func NewExportNonlocalLoginPlacesRequest() (request *ExportNonlocalLoginPlacesRequest) {
	request = &ExportNonlocalLoginPlacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportNonlocalLoginPlaces")
	return
}

func NewExportNonlocalLoginPlacesResponse() (response *ExportNonlocalLoginPlacesResponse) {
	response = &ExportNonlocalLoginPlacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ExportNonlocalLoginPlaces) 用于导出异地登录事件记录CSV文件。
func (c *Client) ExportNonlocalLoginPlaces(request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
	if request == nil {
		request = NewExportNonlocalLoginPlacesRequest()
	}
	response = NewExportNonlocalLoginPlacesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWarningHostConfigRequest() (request *ModifyWarningHostConfigRequest) {
	request = &ModifyWarningHostConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWarningHostConfig")
	return
}

func NewModifyWarningHostConfigResponse() (response *ModifyWarningHostConfigResponse) {
	response = &ModifyWarningHostConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警机器范围配置
func (c *Client) ModifyWarningHostConfig(request *ModifyWarningHostConfigRequest) (response *ModifyWarningHostConfigResponse, err error) {
	if request == nil {
		request = NewModifyWarningHostConfigRequest()
	}
	response = NewModifyWarningHostConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPlanTaskListRequest() (request *DescribeAssetPlanTaskListRequest) {
	request = &DescribeAssetPlanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPlanTaskList")
	return
}

func NewDescribeAssetPlanTaskListResponse() (response *DescribeAssetPlanTaskListResponse) {
	response = &DescribeAssetPlanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理计划任务列表
func (c *Client) DescribeAssetPlanTaskList(request *DescribeAssetPlanTaskListRequest) (response *DescribeAssetPlanTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPlanTaskListRequest()
	}
	response = NewDescribeAssetPlanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebHookRuleRequest() (request *ModifyWebHookRuleRequest) {
	request = &ModifyWebHookRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookRule")
	return
}

func NewModifyWebHookRuleResponse() (response *ModifyWebHookRuleResponse) {
	response = &ModifyWebHookRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改企微机器人规则
func (c *Client) ModifyWebHookRule(request *ModifyWebHookRuleRequest) (response *ModifyWebHookRuleResponse, err error) {
	if request == nil {
		request = NewModifyWebHookRuleRequest()
	}
	response = NewModifyWebHookRuleResponse()
	err = c.Send(request, response)
	return
}

func NewScanAssetRequest() (request *ScanAssetRequest) {
	request = &ScanAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanAsset")
	return
}

func NewScanAssetResponse() (response *ScanAssetResponse) {
	response = &ScanAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产指纹启动扫描
func (c *Client) ScanAsset(request *ScanAssetRequest) (response *ScanAssetResponse, err error) {
	if request == nil {
		request = NewScanAssetRequest()
	}
	response = NewScanAssetResponse()
	err = c.Send(request, response)
	return
}

func NewExportLicenseDetailRequest() (request *ExportLicenseDetailRequest) {
	request = &ExportLicenseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportLicenseDetail")
	return
}

func NewExportLicenseDetailResponse() (response *ExportLicenseDetailResponse) {
	response = &ExportLicenseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出授权列表对应的绑定信息
func (c *Client) ExportLicenseDetail(request *ExportLicenseDetailRequest) (response *ExportLicenseDetailResponse, err error) {
	if request == nil {
		request = NewExportLicenseDetailRequest()
	}
	response = NewExportLicenseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBruteAttackListRequest() (request *DescribeBruteAttackListRequest) {
	request = &DescribeBruteAttackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBruteAttackList")
	return
}

func NewDescribeBruteAttackListResponse() (response *DescribeBruteAttackListResponse) {
	response = &DescribeBruteAttackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取密码破解列表
func (c *Client) DescribeBruteAttackList(request *DescribeBruteAttackListRequest) (response *DescribeBruteAttackListResponse, err error) {
	if request == nil {
		request = NewDescribeBruteAttackListRequest()
	}
	response = NewDescribeBruteAttackListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostInfoRequest() (request *DescribeHostInfoRequest) {
	request = &DescribeHostInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeHostInfo")
	return
}

func NewDescribeHostInfoResponse() (response *DescribeHostInfoResponse) {
	response = &DescribeHostInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主机信息与标签信息查询
func (c *Client) DescribeHostInfo(request *DescribeHostInfoRequest) (response *DescribeHostInfoResponse, err error) {
	if request == nil {
		request = NewDescribeHostInfoRequest()
	}
	response = NewDescribeHostInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginPermissionRequest() (request *DescribeLoginPermissionRequest) {
	request = &DescribeLoginPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginPermission")
	return
}

func NewDescribeLoginPermissionResponse() (response *DescribeLoginPermissionResponse) {
	response = &DescribeLoginPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户主机登录权限状态
func (c *Client) DescribeLoginPermission(request *DescribeLoginPermissionRequest) (response *DescribeLoginPermissionResponse, err error) {
	if request == nil {
		request = NewDescribeLoginPermissionRequest()
	}
	response = NewDescribeLoginPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
	request = &RecoverMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RecoverMalwares")
	return
}

func NewRecoverMalwaresResponse() (response *RecoverMalwaresResponse) {
	response = &RecoverMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RecoverMalwares）用于批量恢复已经被隔离的木马文件。
func (c *Client) RecoverMalwares(request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
	if request == nil {
		request = NewRecoverMalwaresRequest()
	}
	response = NewRecoverMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveLocalStorageItemRequest() (request *RemoveLocalStorageItemRequest) {
	request = &RemoveLocalStorageItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RemoveLocalStorageItem")
	return
}

func NewRemoveLocalStorageItemResponse() (response *RemoveLocalStorageItemResponse) {
	response = &RemoveLocalStorageItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除本地存储数据
func (c *Client) RemoveLocalStorageItem(request *RemoveLocalStorageItemRequest) (response *RemoveLocalStorageItemResponse, err error) {
	if request == nil {
		request = NewRemoveLocalStorageItemRequest()
	}
	response = NewRemoveLocalStorageItemResponse()
	err = c.Send(request, response)
	return
}

func NewScanBaselineRequest() (request *ScanBaselineRequest) {
	request = &ScanBaselineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ScanBaseline")
	return
}

func NewScanBaselineResponse() (response *ScanBaselineResponse) {
	response = &ScanBaselineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 基线检测与基线重新检测接口
func (c *Client) ScanBaseline(request *ScanBaselineRequest) (response *ScanBaselineResponse, err error) {
	if request == nil {
		request = NewScanBaselineRequest()
	}
	response = NewScanBaselineResponse()
	err = c.Send(request, response)
	return
}

func NewChangeStrategyEnableStatusRequest() (request *ChangeStrategyEnableStatusRequest) {
	request = &ChangeStrategyEnableStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ChangeStrategyEnableStatus")
	return
}

func NewChangeStrategyEnableStatusResponse() (response *ChangeStrategyEnableStatusResponse) {
	response = &ChangeStrategyEnableStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据策略id修改策略可用状态
func (c *Client) ChangeStrategyEnableStatus(request *ChangeStrategyEnableStatusRequest) (response *ChangeStrategyEnableStatusResponse, err error) {
	if request == nil {
		request = NewChangeStrategyEnableStatusRequest()
	}
	response = NewChangeStrategyEnableStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppCountRequest() (request *DescribeAssetWebAppCountRequest) {
	request = &DescribeAssetWebAppCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppCount")
	return
}

func NewDescribeAssetWebAppCountResponse() (response *DescribeAssetWebAppCountResponse) {
	response = &DescribeAssetWebAppCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web应用数量
func (c *Client) DescribeAssetWebAppCount(request *DescribeAssetWebAppCountRequest) (response *DescribeAssetWebAppCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppCountRequest()
	}
	response = NewDescribeAssetWebAppCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerDetailRequest() (request *DescribeCustomerDetailRequest) {
	request = &DescribeCustomerDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeCustomerDetail")
	return
}

func NewDescribeCustomerDetailResponse() (response *DescribeCustomerDetailResponse) {
	response = &DescribeCustomerDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询客户信息
func (c *Client) DescribeCustomerDetail(request *DescribeCustomerDetailRequest) (response *DescribeCustomerDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerDetailRequest()
	}
	response = NewDescribeCustomerDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsListRequest() (request *DescribeRiskDnsListRequest) {
	request = &DescribeRiskDnsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsList")
	return
}

func NewDescribeRiskDnsListResponse() (response *DescribeRiskDnsListResponse) {
	response = &DescribeRiskDnsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵检测，获取恶意请求列表
func (c *Client) DescribeRiskDnsList(request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsListRequest()
	}
	response = NewDescribeRiskDnsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectInstallCommandRequest() (request *DescribeDirectConnectInstallCommandRequest) {
	request = &DescribeDirectConnectInstallCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeDirectConnectInstallCommand")
	return
}

func NewDescribeDirectConnectInstallCommandResponse() (response *DescribeDirectConnectInstallCommandResponse) {
	response = &DescribeDirectConnectInstallCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取专线agent安装命令，包含token
func (c *Client) DescribeDirectConnectInstallCommand(request *DescribeDirectConnectInstallCommandRequest) (response *DescribeDirectConnectInstallCommandResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectInstallCommandRequest()
	}
	response = NewDescribeDirectConnectInstallCommandResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBanStatusRequest() (request *DescribeBanStatusRequest) {
	request = &DescribeBanStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanStatus")
	return
}

func NewDescribeBanStatusResponse() (response *DescribeBanStatusResponse) {
	response = &DescribeBanStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取阻断按钮状态
func (c *Client) DescribeBanStatus(request *DescribeBanStatusRequest) (response *DescribeBanStatusResponse, err error) {
	if request == nil {
		request = NewDescribeBanStatusRequest()
	}
	response = NewDescribeBanStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareWhiteListAffectListRequest() (request *DescribeMalwareWhiteListAffectListRequest) {
	request = &DescribeMalwareWhiteListAffectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareWhiteListAffectList")
	return
}

func NewDescribeMalwareWhiteListAffectListResponse() (response *DescribeMalwareWhiteListAffectListResponse) {
	response = &DescribeMalwareWhiteListAffectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取木马白名单受影响列表
func (c *Client) DescribeMalwareWhiteListAffectList(request *DescribeMalwareWhiteListAffectListRequest) (response *DescribeMalwareWhiteListAffectListResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareWhiteListAffectListRequest()
	}
	response = NewDescribeMalwareWhiteListAffectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsTopicsRequest() (request *DescribeClsTopicsRequest) {
	request = &DescribeClsTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeClsTopics")
	return
}

func NewDescribeClsTopicsResponse() (response *DescribeClsTopicsResponse) {
	response = &DescribeClsTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cls日志主题列表
func (c *Client) DescribeClsTopics(request *DescribeClsTopicsRequest) (response *DescribeClsTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeClsTopicsRequest()
	}
	response = NewDescribeClsTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrialReportRequest() (request *DescribeTrialReportRequest) {
	request = &DescribeTrialReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeTrialReport")
	return
}

func NewDescribeTrialReportResponse() (response *DescribeTrialReportResponse) {
	response = &DescribeTrialReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询试用报告
func (c *Client) DescribeTrialReport(request *DescribeTrialReportRequest) (response *DescribeTrialReportResponse, err error) {
	if request == nil {
		request = NewDescribeTrialReportRequest()
	}
	response = NewDescribeTrialReportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBaselineRuleRequest() (request *ModifyBaselineRuleRequest) {
	request = &ModifyBaselineRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineRule")
	return
}

func NewModifyBaselineRuleResponse() (response *ModifyBaselineRuleResponse) {
	response = &ModifyBaselineRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改基线检测规则
func (c *Client) ModifyBaselineRule(request *ModifyBaselineRuleRequest) (response *ModifyBaselineRuleResponse, err error) {
	if request == nil {
		request = NewModifyBaselineRuleRequest()
	}
	response = NewModifyBaselineRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOrderAttributeRequest() (request *ModifyOrderAttributeRequest) {
	request = &ModifyOrderAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyOrderAttribute")
	return
}

func NewModifyOrderAttributeResponse() (response *ModifyOrderAttributeResponse) {
	response = &ModifyOrderAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对订单属性编辑
func (c *Client) ModifyOrderAttribute(request *ModifyOrderAttributeRequest) (response *ModifyOrderAttributeResponse, err error) {
	if request == nil {
		request = NewModifyOrderAttributeRequest()
	}
	response = NewModifyOrderAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaLabelRequest() (request *DeleteMaLabelRequest) {
	request = &DeleteMaLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaLabel")
	return
}

func NewDeleteMaLabelResponse() (response *DeleteMaLabelResponse) {
	response = &DeleteMaLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微告警标签
func (c *Client) DeleteMaLabel(request *DeleteMaLabelRequest) (response *DeleteMaLabelResponse, err error) {
	if request == nil {
		request = NewDeleteMaLabelRequest()
	}
	response = NewDeleteMaLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulStoreListRequest() (request *DescribeVulStoreListRequest) {
	request = &DescribeVulStoreListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulStoreList")
	return
}

func NewDescribeVulStoreListResponse() (response *DescribeVulStoreListResponse) {
	response = &DescribeVulStoreListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞库列表
func (c *Client) DescribeVulStoreList(request *DescribeVulStoreListRequest) (response *DescribeVulStoreListResponse, err error) {
	if request == nil {
		request = NewDescribeVulStoreListRequest()
	}
	response = NewDescribeVulStoreListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineListRequest() (request *DescribeBaselineListRequest) {
	request = &DescribeBaselineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineList")
	return
}

func NewDescribeBaselineListResponse() (response *DescribeBaselineListResponse) {
	response = &DescribeBaselineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询基线列表信息
func (c *Client) DescribeBaselineList(request *DescribeBaselineListRequest) (response *DescribeBaselineListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineListRequest()
	}
	response = NewDescribeBaselineListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebPageProtectSettingRequest() (request *ModifyWebPageProtectSettingRequest) {
	request = &ModifyWebPageProtectSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectSetting")
	return
}

func NewModifyWebPageProtectSettingResponse() (response *ModifyWebPageProtectSettingResponse) {
	response = &ModifyWebPageProtectSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网站防护设置
func (c *Client) ModifyWebPageProtectSetting(request *ModifyWebPageProtectSettingRequest) (response *ModifyWebPageProtectSettingResponse, err error) {
	if request == nil {
		request = NewModifyWebPageProtectSettingRequest()
	}
	response = NewModifyWebPageProtectSettingResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveMachineRequest() (request *RemoveMachineRequest) {
	request = &RemoveMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "RemoveMachine")
	return
}

func NewRemoveMachineResponse() (response *RemoveMachineResponse) {
	response = &RemoveMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主机所有记录，目前只支持非云主机，且需要主机在离线状态
func (c *Client) RemoveMachine(request *RemoveMachineRequest) (response *RemoveMachineResponse, err error) {
	if request == nil {
		request = NewRemoveMachineRequest()
	}
	response = NewRemoveMachineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineBasicInfoRequest() (request *DescribeBaselineBasicInfoRequest) {
	request = &DescribeBaselineBasicInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineBasicInfo")
	return
}

func NewDescribeBaselineBasicInfoResponse() (response *DescribeBaselineBasicInfoResponse) {
	response = &DescribeBaselineBasicInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询基线基础信息列表
func (c *Client) DescribeBaselineBasicInfo(request *DescribeBaselineBasicInfoRequest) (response *DescribeBaselineBasicInfoResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineBasicInfoRequest()
	}
	response = NewDescribeBaselineBasicInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellRulesRequest() (request *DeleteReverseShellRulesRequest) {
	request = &DeleteReverseShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteReverseShellRules")
	return
}

func NewDeleteReverseShellRulesResponse() (response *DeleteReverseShellRulesResponse) {
	response = &DeleteReverseShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除反弹Shell规则
func (c *Client) DeleteReverseShellRules(request *DeleteReverseShellRulesRequest) (response *DeleteReverseShellRulesResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellRulesRequest()
	}
	response = NewDeleteReverseShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskProcessEventsRequest() (request *DescribeRiskProcessEventsRequest) {
	request = &DescribeRiskProcessEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskProcessEvents")
	return
}

func NewDescribeRiskProcessEventsResponse() (response *DescribeRiskProcessEventsResponse) {
	response = &DescribeRiskProcessEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异常进程列表
func (c *Client) DescribeRiskProcessEvents(request *DescribeRiskProcessEventsRequest) (response *DescribeRiskProcessEventsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskProcessEventsRequest()
	}
	response = NewDescribeRiskProcessEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportBashEventsRequest() (request *ExportBashEventsRequest) {
	request = &ExportBashEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBashEvents")
	return
}

func NewExportBashEventsResponse() (response *ExportBashEventsResponse) {
	response = &ExportBashEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出高危命令事件
func (c *Client) ExportBashEvents(request *ExportBashEventsRequest) (response *ExportBashEventsResponse, err error) {
	if request == nil {
		request = NewExportBashEventsRequest()
	}
	response = NewExportBashEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportRiskDnsPolicyListRequest() (request *ExportRiskDnsPolicyListRequest) {
	request = &ExportRiskDnsPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskDnsPolicyList")
	return
}

func NewExportRiskDnsPolicyListResponse() (response *ExportRiskDnsPolicyListResponse) {
	response = &ExportRiskDnsPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出恶意请求策略列表
func (c *Client) ExportRiskDnsPolicyList(request *ExportRiskDnsPolicyListRequest) (response *ExportRiskDnsPolicyListResponse, err error) {
	if request == nil {
		request = NewExportRiskDnsPolicyListRequest()
	}
	response = NewExportRiskDnsPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBanWhiteListRequest() (request *DeleteBanWhiteListRequest) {
	request = &DeleteBanWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBanWhiteList")
	return
}

func NewDeleteBanWhiteListResponse() (response *DeleteBanWhiteListResponse) {
	response = &DeleteBanWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除阻断白名单列表
func (c *Client) DeleteBanWhiteList(request *DeleteBanWhiteListRequest) (response *DeleteBanWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteBanWhiteListRequest()
	}
	response = NewDeleteBanWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewExportRansomDefenseBackupListRequest() (request *ExportRansomDefenseBackupListRequest) {
	request = &ExportRansomDefenseBackupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseBackupList")
	return
}

func NewExportRansomDefenseBackupListResponse() (response *ExportRansomDefenseBackupListResponse) {
	response = &ExportRansomDefenseBackupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出主机快照备份列表
func (c *Client) ExportRansomDefenseBackupList(request *ExportRansomDefenseBackupListRequest) (response *ExportRansomDefenseBackupListResponse, err error) {
	if request == nil {
		request = NewExportRansomDefenseBackupListRequest()
	}
	response = NewExportRansomDefenseBackupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsEventInfoRequest() (request *DescribeRiskDnsEventInfoRequest) {
	request = &DescribeRiskDnsEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsEventInfo")
	return
}

func NewDescribeRiskDnsEventInfoResponse() (response *DescribeRiskDnsEventInfoResponse) {
	response = &DescribeRiskDnsEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件详情
func (c *Client) DescribeRiskDnsEventInfo(request *DescribeRiskDnsEventInfoRequest) (response *DescribeRiskDnsEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsEventInfoRequest()
	}
	response = NewDescribeRiskDnsEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportWebPageEventListRequest() (request *ExportWebPageEventListRequest) {
	request = &ExportWebPageEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportWebPageEventList")
	return
}

func NewExportWebPageEventListResponse() (response *ExportWebPageEventListResponse) {
	response = &ExportWebPageEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出篡改事件列表
func (c *Client) ExportWebPageEventList(request *ExportWebPageEventListRequest) (response *ExportWebPageEventListResponse, err error) {
	if request == nil {
		request = NewExportWebPageEventListRequest()
	}
	response = NewExportWebPageEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseBindScheduleRequest() (request *DescribeLicenseBindScheduleRequest) {
	request = &DescribeLicenseBindScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseBindSchedule")
	return
}

func NewDescribeLicenseBindScheduleResponse() (response *DescribeLicenseBindScheduleResponse) {
	response = &DescribeLicenseBindScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权绑定任务的进度
func (c *Client) DescribeLicenseBindSchedule(request *DescribeLicenseBindScheduleRequest) (response *DescribeLicenseBindScheduleResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseBindScheduleRequest()
	}
	response = NewDescribeLicenseBindScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationListRequest() (request *DescribeAssetWebLocationListRequest) {
	request = &DescribeAssetWebLocationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationList")
	return
}

func NewDescribeAssetWebLocationListResponse() (response *DescribeAssetWebLocationListResponse) {
	response = &DescribeAssetWebLocationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点列表
func (c *Client) DescribeAssetWebLocationList(request *DescribeAssetWebLocationListRequest) (response *DescribeAssetWebLocationListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationListRequest()
	}
	response = NewDescribeAssetWebLocationListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgnoreBaselineRuleRequest() (request *DescribeIgnoreBaselineRuleRequest) {
	request = &DescribeIgnoreBaselineRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreBaselineRule")
	return
}

func NewDescribeIgnoreBaselineRuleResponse() (response *DescribeIgnoreBaselineRuleResponse) {
	response = &DescribeIgnoreBaselineRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询已经忽略的检测项信息
func (c *Client) DescribeIgnoreBaselineRule(request *DescribeIgnoreBaselineRuleRequest) (response *DescribeIgnoreBaselineRuleResponse, err error) {
	if request == nil {
		request = NewDescribeIgnoreBaselineRuleRequest()
	}
	response = NewDescribeIgnoreBaselineRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSaveOrUpdateWarningsRequest() (request *DescribeSaveOrUpdateWarningsRequest) {
	request = &DescribeSaveOrUpdateWarningsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSaveOrUpdateWarnings")
	return
}

func NewDescribeSaveOrUpdateWarningsResponse() (response *DescribeSaveOrUpdateWarningsResponse) {
	response = &DescribeSaveOrUpdateWarningsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新或者插入用户告警设置(该接口废弃,请调用 ModifyWarningSetting )
func (c *Client) DescribeSaveOrUpdateWarnings(request *DescribeSaveOrUpdateWarningsRequest) (response *DescribeSaveOrUpdateWarningsResponse, err error) {
	if request == nil {
		request = NewDescribeSaveOrUpdateWarningsRequest()
	}
	response = NewDescribeSaveOrUpdateWarningsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetCoreModuleInfoRequest() (request *DescribeAssetCoreModuleInfoRequest) {
	request = &DescribeAssetCoreModuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetCoreModuleInfo")
	return
}

func NewDescribeAssetCoreModuleInfoResponse() (response *DescribeAssetCoreModuleInfoResponse) {
	response = &DescribeAssetCoreModuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取内核模块详情
func (c *Client) DescribeAssetCoreModuleInfo(request *DescribeAssetCoreModuleInfoRequest) (response *DescribeAssetCoreModuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetCoreModuleInfoRequest()
	}
	response = NewDescribeAssetCoreModuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAESKeyRequest() (request *DescribeAESKeyRequest) {
	request = &DescribeAESKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAESKey")
	return
}

func NewDescribeAESKeyResponse() (response *DescribeAESKeyResponse) {
	response = &DescribeAESKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置的aeskey和aesiv
func (c *Client) DescribeAESKey(request *DescribeAESKeyRequest) (response *DescribeAESKeyResponse, err error) {
	if request == nil {
		request = NewDescribeAESKeyRequest()
	}
	response = NewDescribeAESKeyResponse()
	err = c.Send(request, response)
	return
}

func NewKeysLocalStorageRequest() (request *KeysLocalStorageRequest) {
	request = &KeysLocalStorageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "KeysLocalStorage")
	return
}

func NewKeysLocalStorageResponse() (response *KeysLocalStorageResponse) {
	response = &KeysLocalStorageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地存储键值列表
func (c *Client) KeysLocalStorage(request *KeysLocalStorageRequest) (response *KeysLocalStorageResponse, err error) {
	if request == nil {
		request = NewKeysLocalStorageRequest()
	}
	response = NewKeysLocalStorageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
	request = &ModifyProVersionRenewFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyProVersionRenewFlag")
	return
}

func NewModifyProVersionRenewFlagResponse() (response *ModifyProVersionRenewFlagResponse) {
	response = &ModifyProVersionRenewFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyProVersionRenewFlag) 用于修改专业版包年包月续费标识。
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
	if request == nil {
		request = NewModifyProVersionRenewFlagRequest()
	}
	response = NewModifyProVersionRenewFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebHookReceiverRequest() (request *DescribeWebHookReceiverRequest) {
	request = &DescribeWebHookReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookReceiver")
	return
}

func NewDescribeWebHookReceiverResponse() (response *DescribeWebHookReceiverResponse) {
	response = &DescribeWebHookReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警接收人列表
func (c *Client) DescribeWebHookReceiver(request *DescribeWebHookReceiverRequest) (response *DescribeWebHookReceiverResponse, err error) {
	if request == nil {
		request = NewDescribeWebHookReceiverRequest()
	}
	response = NewDescribeWebHookReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefenceListRequest() (request *ExportVulDefenceListRequest) {
	request = &ExportVulDefenceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefenceList")
	return
}

func NewExportVulDefenceListResponse() (response *ExportVulDefenceListResponse) {
	response = &ExportVulDefenceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞防御列表
func (c *Client) ExportVulDefenceList(request *ExportVulDefenceListRequest) (response *ExportVulDefenceListResponse, err error) {
	if request == nil {
		request = NewExportVulDefenceListRequest()
	}
	response = NewExportVulDefenceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineRegionsRequest() (request *DescribeMachineRegionsRequest) {
	request = &DescribeMachineRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRegions")
	return
}

func NewDescribeMachineRegionsResponse() (response *DescribeMachineRegionsResponse) {
	response = &DescribeMachineRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器地域列表
func (c *Client) DescribeMachineRegions(request *DescribeMachineRegionsRequest) (response *DescribeMachineRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeMachineRegionsRequest()
	}
	response = NewDescribeMachineRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetInfoRequest() (request *DescribeAssetInfoRequest) {
	request = &DescribeAssetInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetInfo")
	return
}

func NewDescribeAssetInfoResponse() (response *DescribeAssetInfoResponse) {
	response = &DescribeAssetInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产数量： 主机数、账号数、端口数、进程数、软件数、数据库数、Web应用数、Web框架数、Web服务数、Web站点数
func (c *Client) DescribeAssetInfo(request *DescribeAssetInfoRequest) (response *DescribeAssetInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetInfoRequest()
	}
	response = NewDescribeAssetInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
	request = &DescribeProVersionInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProVersionInfo")
	return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
	response = &DescribeProVersionInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取专业版概览信息。
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProVersionInfoRequest()
	}
	response = NewDescribeProVersionInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulHostTopRequest() (request *DescribeVulHostTopRequest) {
	request = &DescribeVulHostTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulHostTop")
	return
}

func NewDescribeVulHostTopResponse() (response *DescribeVulHostTopResponse) {
	response = &DescribeVulHostTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务器风险top列表
func (c *Client) DescribeVulHostTop(request *DescribeVulHostTopRequest) (response *DescribeVulHostTopResponse, err error) {
	if request == nil {
		request = NewDescribeVulHostTopRequest()
	}
	response = NewDescribeVulHostTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityEventsCntRequest() (request *DescribeSecurityEventsCntRequest) {
	request = &DescribeSecurityEventsCntRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityEventsCnt")
	return
}

func NewDescribeSecurityEventsCntResponse() (response *DescribeSecurityEventsCntResponse) {
	response = &DescribeSecurityEventsCntResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全概览相关事件统计数据接口
func (c *Client) DescribeSecurityEventsCnt(request *DescribeSecurityEventsCntRequest) (response *DescribeSecurityEventsCntResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityEventsCntRequest()
	}
	response = NewDescribeSecurityEventsCntResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLicenseOrderRequest() (request *CreateLicenseOrderRequest) {
	request = &CreateLicenseOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateLicenseOrder")
	return
}

func NewCreateLicenseOrderResponse() (response *CreateLicenseOrderResponse) {
	response = &CreateLicenseOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateLicenseOrder 该接口可以创建专业版/旗舰版订单
// 支持预付费后付费创建
// 后付费订单直接创建成功
// 预付费订单仅下单不支付,需要调用计费支付接口进行支付
func (c *Client) CreateLicenseOrder(request *CreateLicenseOrderRequest) (response *CreateLicenseOrderResponse, err error) {
	if request == nil {
		request = NewCreateLicenseOrderRequest()
	}
	response = NewCreateLicenseOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleDetectListRequest() (request *DescribeBaselineRuleDetectListRequest) {
	request = &DescribeBaselineRuleDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleDetectList")
	return
}

func NewDescribeBaselineRuleDetectListResponse() (response *DescribeBaselineRuleDetectListResponse) {
	response = &DescribeBaselineRuleDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线规则检测列表
func (c *Client) DescribeBaselineRuleDetectList(request *DescribeBaselineRuleDetectListRequest) (response *DescribeBaselineRuleDetectListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleDetectListRequest()
	}
	response = NewDescribeBaselineRuleDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogTypeRequest() (request *DescribeLogTypeRequest) {
	request = &DescribeLogTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogType")
	return
}

func NewDescribeLogTypeResponse() (response *DescribeLogTypeResponse) {
	response = &DescribeLogTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志分析功能-获取日志类型，使用该接口返回的结果暂时可过滤的日志类型
func (c *Client) DescribeLogType(request *DescribeLogTypeRequest) (response *DescribeLogTypeResponse, err error) {
	if request == nil {
		request = NewDescribeLogTypeRequest()
	}
	response = NewDescribeLogTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaConnEventListRequest() (request *DescribeMaConnEventListRequest) {
	request = &DescribeMaConnEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaConnEventList")
	return
}

func NewDescribeMaConnEventListResponse() (response *DescribeMaConnEventListResponse) {
	response = &DescribeMaConnEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警访问事件
func (c *Client) DescribeMaConnEventList(request *DescribeMaConnEventListRequest) (response *DescribeMaConnEventListResponse, err error) {
	if request == nil {
		request = NewDescribeMaConnEventListRequest()
	}
	response = NewDescribeMaConnEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebHookPolicyRequest() (request *DescribeWebHookPolicyRequest) {
	request = &DescribeWebHookPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookPolicy")
	return
}

func NewDescribeWebHookPolicyResponse() (response *DescribeWebHookPolicyResponse) {
	response = &DescribeWebHookPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警策略
func (c *Client) DescribeWebHookPolicy(request *DescribeWebHookPolicyRequest) (response *DescribeWebHookPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeWebHookPolicyRequest()
	}
	response = NewDescribeWebHookPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachineRemarkRequest() (request *ModifyMachineRemarkRequest) {
	request = &ModifyMachineRemarkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMachineRemark")
	return
}

func NewModifyMachineRemarkResponse() (response *ModifyMachineRemarkResponse) {
	response = &ModifyMachineRemarkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改主机备注信息
func (c *Client) ModifyMachineRemark(request *ModifyMachineRemarkRequest) (response *ModifyMachineRemarkResponse, err error) {
	if request == nil {
		request = NewModifyMachineRemarkRequest()
	}
	response = NewModifyMachineRemarkResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetMachineListRequest() (request *ExportAssetMachineListRequest) {
	request = &ExportAssetMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetMachineList")
	return
}

func NewExportAssetMachineListResponse() (response *ExportAssetMachineListResponse) {
	response = &ExportAssetMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资源监控列表
func (c *Client) ExportAssetMachineList(request *ExportAssetMachineListRequest) (response *ExportAssetMachineListResponse, err error) {
	if request == nil {
		request = NewExportAssetMachineListRequest()
	}
	response = NewExportAssetMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaCustomPolicyRequest() (request *ModifyMaCustomPolicyRequest) {
	request = &ModifyMaCustomPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaCustomPolicy")
	return
}

func NewModifyMaCustomPolicyResponse() (response *ModifyMaCustomPolicyResponse) {
	response = &ModifyMaCustomPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或者更改微告警自定义白名单
func (c *Client) ModifyMaCustomPolicy(request *ModifyMaCustomPolicyRequest) (response *ModifyMaCustomPolicyResponse, err error) {
	if request == nil {
		request = NewModifyMaCustomPolicyRequest()
	}
	response = NewModifyMaCustomPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePayPageBrowsingRecordsRequest() (request *CreatePayPageBrowsingRecordsRequest) {
	request = &CreatePayPageBrowsingRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreatePayPageBrowsingRecords")
	return
}

func NewCreatePayPageBrowsingRecordsResponse() (response *CreatePayPageBrowsingRecordsResponse) {
	response = &CreatePayPageBrowsingRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对访问并停留在云镜购买页，但最终未购买服务的用户信息进行记录。
func (c *Client) CreatePayPageBrowsingRecords(request *CreatePayPageBrowsingRecordsRequest) (response *CreatePayPageBrowsingRecordsResponse, err error) {
	if request == nil {
		request = NewCreatePayPageBrowsingRecordsRequest()
	}
	response = NewCreatePayPageBrowsingRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyOrderRequest() (request *DestroyOrderRequest) {
	request = &DestroyOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DestroyOrder")
	return
}

func NewDestroyOrderResponse() (response *DestroyOrderResponse) {
	response = &DestroyOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DestroyOrder  该接口可以对资源销毁.
func (c *Client) DestroyOrder(request *DestroyOrderRequest) (response *DestroyOrderResponse, err error) {
	if request == nil {
		request = NewDestroyOrderRequest()
	}
	response = NewDestroyOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellPluginListRequest() (request *DescribeJavaMemShellPluginListRequest) {
	request = &DescribeJavaMemShellPluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellPluginList")
	return
}

func NewDescribeJavaMemShellPluginListResponse() (response *DescribeJavaMemShellPluginListResponse) {
	response = &DescribeJavaMemShellPluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java内存马插件列表
func (c *Client) DescribeJavaMemShellPluginList(request *DescribeJavaMemShellPluginListRequest) (response *DescribeJavaMemShellPluginListResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellPluginListRequest()
	}
	response = NewDescribeJavaMemShellPluginListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBashPolicyStatusRequest() (request *ModifyBashPolicyStatusRequest) {
	request = &ModifyBashPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBashPolicyStatus")
	return
}

func NewModifyBashPolicyStatusResponse() (response *ModifyBashPolicyStatusResponse) {
	response = &ModifyBashPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换高危命令策略状态
func (c *Client) ModifyBashPolicyStatus(request *ModifyBashPolicyStatusRequest) (response *ModifyBashPolicyStatusResponse, err error) {
	if request == nil {
		request = NewModifyBashPolicyStatusRequest()
	}
	response = NewModifyBashPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLicenseUnBindsRequest() (request *ModifyLicenseUnBindsRequest) {
	request = &ModifyLicenseUnBindsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseUnBinds")
	return
}

func NewModifyLicenseUnBindsResponse() (response *ModifyLicenseUnBindsResponse) {
	response = &ModifyLicenseUnBindsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置中心-授权管理 对某个授权批量解绑机器
func (c *Client) ModifyLicenseUnBinds(request *ModifyLicenseUnBindsRequest) (response *ModifyLicenseUnBindsResponse, err error) {
	if request == nil {
		request = NewModifyLicenseUnBindsRequest()
	}
	response = NewModifyLicenseUnBindsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBaselineWeakPasswordRequest() (request *DeleteBaselineWeakPasswordRequest) {
	request = &DeleteBaselineWeakPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineWeakPassword")
	return
}

func NewDeleteBaselineWeakPasswordResponse() (response *DeleteBaselineWeakPasswordResponse) {
	response = &DeleteBaselineWeakPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除基线弱口令
func (c *Client) DeleteBaselineWeakPassword(request *DeleteBaselineWeakPasswordRequest) (response *DeleteBaselineWeakPasswordResponse, err error) {
	if request == nil {
		request = NewDeleteBaselineWeakPasswordRequest()
	}
	response = NewDeleteBaselineWeakPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBashEventsRequest() (request *DeleteBashEventsRequest) {
	request = &DeleteBashEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashEvents")
	return
}

func NewDeleteBashEventsResponse() (response *DeleteBashEventsResponse) {
	response = &DeleteBashEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Ids删除高危命令事件
func (c *Client) DeleteBashEvents(request *DeleteBashEventsRequest) (response *DeleteBashEventsResponse, err error) {
	if request == nil {
		request = NewDeleteBashEventsRequest()
	}
	response = NewDeleteBashEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessCountRequest() (request *DescribeAssetProcessCountRequest) {
	request = &DescribeAssetProcessCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetProcessCount")
	return
}

func NewDescribeAssetProcessCountResponse() (response *DescribeAssetProcessCountResponse) {
	response = &DescribeAssetProcessCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有进程数量
func (c *Client) DescribeAssetProcessCount(request *DescribeAssetProcessCountRequest) (response *DescribeAssetProcessCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessCountRequest()
	}
	response = NewDescribeAssetProcessCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportRiskProcessEventsRequest() (request *ExportRiskProcessEventsRequest) {
	request = &ExportRiskProcessEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskProcessEvents")
	return
}

func NewExportRiskProcessEventsResponse() (response *ExportRiskProcessEventsResponse) {
	response = &ExportRiskProcessEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出异常进程事件
func (c *Client) ExportRiskProcessEvents(request *ExportRiskProcessEventsRequest) (response *ExportRiskProcessEventsResponse, err error) {
	if request == nil {
		request = NewExportRiskProcessEventsRequest()
	}
	response = NewExportRiskProcessEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBaselineRuleIgnoreRequest() (request *DeleteBaselineRuleIgnoreRequest) {
	request = &DeleteBaselineRuleIgnoreRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineRuleIgnore")
	return
}

func NewDeleteBaselineRuleIgnoreResponse() (response *DeleteBaselineRuleIgnoreResponse) {
	response = &DeleteBaselineRuleIgnoreResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除基线忽略规则
func (c *Client) DeleteBaselineRuleIgnore(request *DeleteBaselineRuleIgnoreRequest) (response *DeleteBaselineRuleIgnoreResponse, err error) {
	if request == nil {
		request = NewDeleteBaselineRuleIgnoreRequest()
	}
	response = NewDeleteBaselineRuleIgnoreResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetAttackWhiteListRequest() (request *DeleteNetAttackWhiteListRequest) {
	request = &DeleteNetAttackWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteNetAttackWhiteList")
	return
}

func NewDeleteNetAttackWhiteListResponse() (response *DeleteNetAttackWhiteListResponse) {
	response = &DeleteNetAttackWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) DeleteNetAttackWhiteList(request *DeleteNetAttackWhiteListRequest) (response *DeleteNetAttackWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteNetAttackWhiteListRequest()
	}
	response = NewDeleteNetAttackWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebFrameCountRequest() (request *DescribeAssetWebFrameCountRequest) {
	request = &DescribeAssetWebFrameCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebFrameCount")
	return
}

func NewDescribeAssetWebFrameCountResponse() (response *DescribeAssetWebFrameCountResponse) {
	response = &DescribeAssetWebFrameCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web框架数量
func (c *Client) DescribeAssetWebFrameCount(request *DescribeAssetWebFrameCountRequest) (response *DescribeAssetWebFrameCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebFrameCountRequest()
	}
	response = NewDescribeAssetWebFrameCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenHostInvasionRequest() (request *DescribeScreenHostInvasionRequest) {
	request = &DescribeScreenHostInvasionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenHostInvasion")
	return
}

func NewDescribeScreenHostInvasionResponse() (response *DescribeScreenHostInvasionResponse) {
	response = &DescribeScreenHostInvasionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化主机入侵详情
func (c *Client) DescribeScreenHostInvasion(request *DescribeScreenHostInvasionRequest) (response *DescribeScreenHostInvasionResponse, err error) {
	if request == nil {
		request = NewDescribeScreenHostInvasionRequest()
	}
	response = NewDescribeScreenHostInvasionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskEventsStatusRequest() (request *ModifyRiskEventsStatusRequest) {
	request = &ModifyRiskEventsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskEventsStatus")
	return
}

func NewModifyRiskEventsStatusResponse() (response *ModifyRiskEventsStatusResponse) {
	response = &ModifyRiskEventsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵检测所有事件的状态，包括：文件查杀，异常登录，密码破解，高危命令，反弹shell，本地提取
func (c *Client) ModifyRiskEventsStatus(request *ModifyRiskEventsStatusRequest) (response *ModifyRiskEventsStatusResponse, err error) {
	if request == nil {
		request = NewModifyRiskEventsStatusRequest()
	}
	response = NewModifyRiskEventsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInjectRiskyServiceSwitchRequest() (request *DescribeInjectRiskyServiceSwitchRequest) {
	request = &DescribeInjectRiskyServiceSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeInjectRiskyServiceSwitch")
	return
}

func NewDescribeInjectRiskyServiceSwitchResponse() (response *DescribeInjectRiskyServiceSwitchResponse) {
	response = &DescribeInjectRiskyServiceSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java内存马和rasp是否注入风险服务配置
func (c *Client) DescribeInjectRiskyServiceSwitch(request *DescribeInjectRiskyServiceSwitchRequest) (response *DescribeInjectRiskyServiceSwitchResponse, err error) {
	if request == nil {
		request = NewDescribeInjectRiskyServiceSwitchRequest()
	}
	response = NewDescribeInjectRiskyServiceSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrivilegeEventsRequest() (request *DescribePrivilegeEventsRequest) {
	request = &DescribePrivilegeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeEvents")
	return
}

func NewDescribePrivilegeEventsResponse() (response *DescribePrivilegeEventsResponse) {
	response = &DescribePrivilegeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地提权事件列表
func (c *Client) DescribePrivilegeEvents(request *DescribePrivilegeEventsRequest) (response *DescribePrivilegeEventsResponse, err error) {
	if request == nil {
		request = NewDescribePrivilegeEventsRequest()
	}
	response = NewDescribePrivilegeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineHostTopRequest() (request *DescribeBaselineHostTopRequest) {
	request = &DescribeBaselineHostTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostTop")
	return
}

func NewDescribeBaselineHostTopResponse() (response *DescribeBaselineHostTopResponse) {
	response = &DescribeBaselineHostTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口返回TopN的风险服务器
func (c *Client) DescribeBaselineHostTop(request *DescribeBaselineHostTopRequest) (response *DescribeBaselineHostTopResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineHostTopRequest()
	}
	response = NewDescribeBaselineHostTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginWhiteCombinedListRequest() (request *DescribeLoginWhiteCombinedListRequest) {
	request = &DescribeLoginWhiteCombinedListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteCombinedList")
	return
}

func NewDescribeLoginWhiteCombinedListResponse() (response *DescribeLoginWhiteCombinedListResponse) {
	response = &DescribeLoginWhiteCombinedListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异地登录白名单合并后列表
func (c *Client) DescribeLoginWhiteCombinedList(request *DescribeLoginWhiteCombinedListRequest) (response *DescribeLoginWhiteCombinedListResponse, err error) {
	if request == nil {
		request = NewDescribeLoginWhiteCombinedListRequest()
	}
	response = NewDescribeLoginWhiteCombinedListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenEmergentMsgRequest() (request *DescribeScreenEmergentMsgRequest) {
	request = &DescribeScreenEmergentMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenEmergentMsg")
	return
}

func NewDescribeScreenEmergentMsgResponse() (response *DescribeScreenEmergentMsgResponse) {
	response = &DescribeScreenEmergentMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化紧急通知
func (c *Client) DescribeScreenEmergentMsg(request *DescribeScreenEmergentMsgRequest) (response *DescribeScreenEmergentMsgResponse, err error) {
	if request == nil {
		request = NewDescribeScreenEmergentMsgRequest()
	}
	response = NewDescribeScreenEmergentMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebPageEventListRequest() (request *DescribeWebPageEventListRequest) {
	request = &DescribeWebPageEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageEventList")
	return
}

func NewDescribeWebPageEventListResponse() (response *DescribeWebPageEventListResponse) {
	response = &DescribeWebPageEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询篡改事件列表
func (c *Client) DescribeWebPageEventList(request *DescribeWebPageEventListRequest) (response *DescribeWebPageEventListResponse, err error) {
	if request == nil {
		request = NewDescribeWebPageEventListRequest()
	}
	response = NewDescribeWebPageEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
	request = &DescribeLoginWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteList")
	return
}

func NewDescribeLoginWhiteListResponse() (response *DescribeLoginWhiteListResponse) {
	response = &DescribeLoginWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异地登录白名单列表
func (c *Client) DescribeLoginWhiteList(request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeLoginWhiteListRequest()
	}
	response = NewDescribeLoginWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCanNotSeparateMachineRequest() (request *DescribeCanNotSeparateMachineRequest) {
	request = &DescribeCanNotSeparateMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeCanNotSeparateMachine")
	return
}

func NewDescribeCanNotSeparateMachineResponse() (response *DescribeCanNotSeparateMachineResponse) {
	response = &DescribeCanNotSeparateMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取木马不可隔离的主机
func (c *Client) DescribeCanNotSeparateMachine(request *DescribeCanNotSeparateMachineRequest) (response *DescribeCanNotSeparateMachineResponse, err error) {
	if request == nil {
		request = NewDescribeCanNotSeparateMachineRequest()
	}
	response = NewDescribeCanNotSeparateMachineResponse()
	err = c.Send(request, response)
	return
}

func NewExportAttackLogsRequest() (request *ExportAttackLogsRequest) {
	request = &ExportAttackLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAttackLogs")
	return
}

func NewExportAttackLogsResponse() (response *ExportAttackLogsResponse) {
	response = &ExportAttackLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出网络攻击日志
func (c *Client) ExportAttackLogs(request *ExportAttackLogsRequest) (response *ExportAttackLogsResponse, err error) {
	if request == nil {
		request = NewExportAttackLogsRequest()
	}
	response = NewExportAttackLogsResponse()
	err = c.Send(request, response)
	return
}

func NewOpenProVersionActivityRequest() (request *OpenProVersionActivityRequest) {
	request = &OpenProVersionActivityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "OpenProVersionActivity")
	return
}

func NewOpenProVersionActivityResponse() (response *OpenProVersionActivityResponse) {
	response = &OpenProVersionActivityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 活动开通专业版
func (c *Client) OpenProVersionActivity(request *OpenProVersionActivityRequest) (response *OpenProVersionActivityResponse, err error) {
	if request == nil {
		request = NewOpenProVersionActivityRequest()
	}
	response = NewOpenProVersionActivityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
	request = &DescribeSecurityTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityTrends")
	return
}

func NewDescribeSecurityTrendsResponse() (response *DescribeSecurityTrendsResponse) {
	response = &DescribeSecurityTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeSecurityTrends) 用于获取安全事件统计数据。
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityTrendsRequest()
	}
	response = NewDescribeSecurityTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBanModeRequest() (request *ModifyBanModeRequest) {
	request = &ModifyBanModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanMode")
	return
}

func NewModifyBanModeResponse() (response *ModifyBanModeResponse) {
	response = &ModifyBanModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改爆破阻断模式
func (c *Client) ModifyBanMode(request *ModifyBanModeRequest) (response *ModifyBanModeResponse, err error) {
	if request == nil {
		request = NewModifyBanModeRequest()
	}
	response = NewModifyBanModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulInfoCvssRequest() (request *DescribeVulInfoCvssRequest) {
	request = &DescribeVulInfoCvssRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulInfoCvss")
	return
}

func NewDescribeVulInfoCvssResponse() (response *DescribeVulInfoCvssResponse) {
	response = &DescribeVulInfoCvssResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞详情，带CVSS版本
func (c *Client) DescribeVulInfoCvss(request *DescribeVulInfoCvssRequest) (response *DescribeVulInfoCvssResponse, err error) {
	if request == nil {
		request = NewDescribeVulInfoCvssRequest()
	}
	response = NewDescribeVulInfoCvssResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulOverviewRequest() (request *DescribeVulOverviewRequest) {
	request = &DescribeVulOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulOverview")
	return
}

func NewDescribeVulOverviewResponse() (response *DescribeVulOverviewResponse) {
	response = &DescribeVulOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞概览数据
func (c *Client) DescribeVulOverview(request *DescribeVulOverviewRequest) (response *DescribeVulOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeVulOverviewRequest()
	}
	response = NewDescribeVulOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewExportScanTaskDetailsRequest() (request *ExportScanTaskDetailsRequest) {
	request = &ExportScanTaskDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportScanTaskDetails")
	return
}

func NewExportScanTaskDetailsResponse() (response *ExportScanTaskDetailsResponse) {
	response = &ExportScanTaskDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务id导出指定扫描任务详情
func (c *Client) ExportScanTaskDetails(request *ExportScanTaskDetailsRequest) (response *ExportScanTaskDetailsResponse, err error) {
	if request == nil {
		request = NewExportScanTaskDetailsRequest()
	}
	response = NewExportScanTaskDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
	request = &TrustMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "TrustMalwares")
	return
}

func NewTrustMalwaresResponse() (response *TrustMalwaresResponse) {
	response = &TrustMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(TrustMalwares)将被识别木马文件设为信任。
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
	if request == nil {
		request = NewTrustMalwaresRequest()
	}
	response = NewTrustMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperEventsRequest() (request *DescribeFileTamperEventsRequest) {
	request = &DescribeFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperEvents")
	return
}

func NewDescribeFileTamperEventsResponse() (response *DescribeFileTamperEventsResponse) {
	response = &DescribeFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 核心文件监控事件列表
func (c *Client) DescribeFileTamperEvents(request *DescribeFileTamperEventsRequest) (response *DescribeFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperEventsRequest()
	}
	response = NewDescribeFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogClsDeliverInfoRequest() (request *DeleteLogClsDeliverInfoRequest) {
	request = &DeleteLogClsDeliverInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteLogClsDeliverInfo")
	return
}

func NewDeleteLogClsDeliverInfoResponse() (response *DeleteLogClsDeliverInfoResponse) {
	response = &DeleteLogClsDeliverInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志cls投递配置
func (c *Client) DeleteLogClsDeliverInfo(request *DeleteLogClsDeliverInfoRequest) (response *DeleteLogClsDeliverInfoResponse, err error) {
	if request == nil {
		request = NewDeleteLogClsDeliverInfoRequest()
	}
	response = NewDeleteLogClsDeliverInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
	request = &DescribeMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachines")
	return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
	response = &DescribeMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeMachines) 用于获取区域主机列表。
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesRequest()
	}
	response = NewDescribeMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStorageRecordRequest() (request *DescribeLogStorageRecordRequest) {
	request = &DescribeLogStorageRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageRecord")
	return
}

func NewDescribeLogStorageRecordResponse() (response *DescribeLogStorageRecordResponse) {
	response = &DescribeLogStorageRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志存储量记录
func (c *Client) DescribeLogStorageRecord(request *DescribeLogStorageRecordRequest) (response *DescribeLogStorageRecordResponse, err error) {
	if request == nil {
		request = NewDescribeLogStorageRecordRequest()
	}
	response = NewDescribeLogStorageRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGeneralNewRequest() (request *DescribeMachineGeneralNewRequest) {
	request = &DescribeMachineGeneralNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineGeneralNew")
	return
}

func NewDescribeMachineGeneralNewResponse() (response *DescribeMachineGeneralNewResponse) {
	response = &DescribeMachineGeneralNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机概览信息
func (c *Client) DescribeMachineGeneralNew(request *DescribeMachineGeneralNewRequest) (response *DescribeMachineGeneralNewResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGeneralNewRequest()
	}
	response = NewDescribeMachineGeneralNewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBanWhiteListRequest() (request *CreateBanWhiteListRequest) {
	request = &CreateBanWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateBanWhiteList")
	return
}

func NewCreateBanWhiteListResponse() (response *CreateBanWhiteListResponse) {
	response = &CreateBanWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加阻断白名单列表
func (c *Client) CreateBanWhiteList(request *CreateBanWhiteListRequest) (response *CreateBanWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateBanWhiteListRequest()
	}
	response = NewCreateBanWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineHostDetectListRequest() (request *ExportBaselineHostDetectListRequest) {
	request = &ExportBaselineHostDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineHostDetectList")
	return
}

func NewExportBaselineHostDetectListResponse() (response *ExportBaselineHostDetectListResponse) {
	response = &ExportBaselineHostDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线主机检测
func (c *Client) ExportBaselineHostDetectList(request *ExportBaselineHostDetectListRequest) (response *ExportBaselineHostDetectListResponse, err error) {
	if request == nil {
		request = NewExportBaselineHostDetectListRequest()
	}
	response = NewExportBaselineHostDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellPluginInfoRequest() (request *DescribeJavaMemShellPluginInfoRequest) {
	request = &DescribeJavaMemShellPluginInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellPluginInfo")
	return
}

func NewDescribeJavaMemShellPluginInfoResponse() (response *DescribeJavaMemShellPluginInfoResponse) {
	response = &DescribeJavaMemShellPluginInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询给定主机java内存马插件信息
func (c *Client) DescribeJavaMemShellPluginInfo(request *DescribeJavaMemShellPluginInfoRequest) (response *DescribeJavaMemShellPluginInfoResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellPluginInfoRequest()
	}
	response = NewDescribeJavaMemShellPluginInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMemShellRulesRequest() (request *DescribeMemShellRulesRequest) {
	request = &DescribeMemShellRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMemShellRules")
	return
}

func NewDescribeMemShellRulesResponse() (response *DescribeMemShellRulesResponse) {
	response = &DescribeMemShellRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取java内存马白名单列表
func (c *Client) DescribeMemShellRules(request *DescribeMemShellRulesRequest) (response *DescribeMemShellRulesResponse, err error) {
	if request == nil {
		request = NewDescribeMemShellRulesRequest()
	}
	response = NewDescribeMemShellRulesResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebFrameListRequest() (request *ExportAssetWebFrameListRequest) {
	request = &ExportAssetWebFrameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebFrameList")
	return
}

func NewExportAssetWebFrameListResponse() (response *ExportAssetWebFrameListResponse) {
	response = &ExportAssetWebFrameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web框架列表
func (c *Client) ExportAssetWebFrameList(request *ExportAssetWebFrameListRequest) (response *ExportAssetWebFrameListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebFrameListRequest()
	}
	response = NewExportAssetWebFrameListResponse()
	err = c.Send(request, response)
	return
}

func NewStartBaselineDetectRequest() (request *StartBaselineDetectRequest) {
	request = &StartBaselineDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "StartBaselineDetect")
	return
}

func NewStartBaselineDetectResponse() (response *StartBaselineDetectResponse) {
	response = &StartBaselineDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检测基线
func (c *Client) StartBaselineDetect(request *StartBaselineDetectRequest) (response *StartBaselineDetectResponse, err error) {
	if request == nil {
		request = NewStartBaselineDetectRequest()
	}
	response = NewStartBaselineDetectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashEventsInfoNewRequest() (request *DescribeBashEventsInfoNewRequest) {
	request = &DescribeBashEventsInfoNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsInfoNew")
	return
}

func NewDescribeBashEventsInfoNewResponse() (response *DescribeBashEventsInfoNewResponse) {
	response = &DescribeBashEventsInfoNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询高危命令事件详情(新)
func (c *Client) DescribeBashEventsInfoNew(request *DescribeBashEventsInfoNewRequest) (response *DescribeBashEventsInfoNewResponse, err error) {
	if request == nil {
		request = NewDescribeBashEventsInfoNewRequest()
	}
	response = NewDescribeBashEventsInfoNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventListRequest() (request *DescribeEventListRequest) {
	request = &DescribeEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeEventList")
	return
}

func NewDescribeEventListResponse() (response *DescribeEventListResponse) {
	response = &DescribeEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理模块，获取事件列表
func (c *Client) DescribeEventList(request *DescribeEventListRequest) (response *DescribeEventListResponse, err error) {
	if request == nil {
		request = NewDescribeEventListRequest()
	}
	response = NewDescribeEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineDefenseCntRequest() (request *DescribeMachineDefenseCntRequest) {
	request = &DescribeMachineDefenseCntRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineDefenseCnt")
	return
}

func NewDescribeMachineDefenseCntResponse() (response *DescribeMachineDefenseCntResponse) {
	response = &DescribeMachineDefenseCntResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机高级防御事件数统计
func (c *Client) DescribeMachineDefenseCnt(request *DescribeMachineDefenseCntRequest) (response *DescribeMachineDefenseCntResponse, err error) {
	if request == nil {
		request = NewDescribeMachineDefenseCntRequest()
	}
	response = NewDescribeMachineDefenseCntResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserCountRequest() (request *DescribeAssetUserCountRequest) {
	request = &DescribeAssetUserCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserCount")
	return
}

func NewDescribeAssetUserCountResponse() (response *DescribeAssetUserCountResponse) {
	response = &DescribeAssetUserCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有账号数量
func (c *Client) DescribeAssetUserCount(request *DescribeAssetUserCountRequest) (response *DescribeAssetUserCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserCountRequest()
	}
	response = NewDescribeAssetUserCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceSettingListRequest() (request *DescribeVulDefenceSettingListRequest) {
	request = &DescribeVulDefenceSettingListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceSettingList")
	return
}

func NewDescribeVulDefenceSettingListResponse() (response *DescribeVulDefenceSettingListResponse) {
	response = &DescribeVulDefenceSettingListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞防御设置列表
func (c *Client) DescribeVulDefenceSettingList(request *DescribeVulDefenceSettingListRequest) (response *DescribeVulDefenceSettingListResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceSettingListRequest()
	}
	response = NewDescribeVulDefenceSettingListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulFixStatusRequest() (request *DescribeVulFixStatusRequest) {
	request = &DescribeVulFixStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulFixStatus")
	return
}

func NewDescribeVulFixStatusResponse() (response *DescribeVulFixStatusResponse) {
	response = &DescribeVulFixStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞修护-查找主机漏洞修护进度
func (c *Client) DescribeVulFixStatus(request *DescribeVulFixStatusRequest) (response *DescribeVulFixStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVulFixStatusRequest()
	}
	response = NewDescribeVulFixStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProtectServerRequest() (request *CreateProtectServerRequest) {
	request = &CreateProtectServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateProtectServer")
	return
}

func NewCreateProtectServerResponse() (response *CreateProtectServerResponse) {
	response = &CreateProtectServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网站防护服务器
func (c *Client) CreateProtectServer(request *CreateProtectServerRequest) (response *CreateProtectServerResponse, err error) {
	if request == nil {
		request = NewCreateProtectServerRequest()
	}
	response = NewCreateProtectServerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttackLogsRequest() (request *DeleteAttackLogsRequest) {
	request = &DeleteAttackLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteAttackLogs")
	return
}

func NewDeleteAttackLogsResponse() (response *DeleteAttackLogsResponse) {
	response = &DeleteAttackLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络攻击日志
func (c *Client) DeleteAttackLogs(request *DeleteAttackLogsRequest) (response *DeleteAttackLogsResponse, err error) {
	if request == nil {
		request = NewDeleteAttackLogsRequest()
	}
	response = NewDeleteAttackLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceCountRequest() (request *DescribeAssetWebServiceCountRequest) {
	request = &DescribeAssetWebServiceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceCount")
	return
}

func NewDescribeAssetWebServiceCountResponse() (response *DescribeAssetWebServiceCountResponse) {
	response = &DescribeAssetWebServiceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web服务数量
func (c *Client) DescribeAssetWebServiceCount(request *DescribeAssetWebServiceCountRequest) (response *DescribeAssetWebServiceCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceCountRequest()
	}
	response = NewDescribeAssetWebServiceCountResponse()
	err = c.Send(request, response)
	return
}

func NewSetBashEventsStatusRequest() (request *SetBashEventsStatusRequest) {
	request = &SetBashEventsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SetBashEventsStatus")
	return
}

func NewSetBashEventsStatusResponse() (response *SetBashEventsStatusResponse) {
	response = &SetBashEventsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置高危命令事件状态
func (c *Client) SetBashEventsStatus(request *SetBashEventsStatusRequest) (response *SetBashEventsStatusResponse, err error) {
	if request == nil {
		request = NewSetBashEventsStatusRequest()
	}
	response = NewSetBashEventsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaSelfLearnPolicyListRequest() (request *DescribeMaSelfLearnPolicyListRequest) {
	request = &DescribeMaSelfLearnPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaSelfLearnPolicyList")
	return
}

func NewDescribeMaSelfLearnPolicyListResponse() (response *DescribeMaSelfLearnPolicyListResponse) {
	response = &DescribeMaSelfLearnPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警自学习子策略列表
func (c *Client) DescribeMaSelfLearnPolicyList(request *DescribeMaSelfLearnPolicyListRequest) (response *DescribeMaSelfLearnPolicyListResponse, err error) {
	if request == nil {
		request = NewDescribeMaSelfLearnPolicyListRequest()
	}
	response = NewDescribeMaSelfLearnPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
	request = &DescribeProcessStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeProcessStatistics")
	return
}

func NewDescribeProcessStatisticsResponse() (response *DescribeProcessStatisticsResponse) {
	response = &DescribeProcessStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeProcessStatistics) 用于获取进程统计列表数据。
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeProcessStatisticsRequest()
	}
	response = NewDescribeProcessStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewExportFileTamperRulesRequest() (request *ExportFileTamperRulesRequest) {
	request = &ExportFileTamperRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportFileTamperRules")
	return
}

func NewExportFileTamperRulesResponse() (response *ExportFileTamperRulesResponse) {
	response = &ExportFileTamperRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出核心文件监控规则
func (c *Client) ExportFileTamperRules(request *ExportFileTamperRulesRequest) (response *ExportFileTamperRulesResponse, err error) {
	if request == nil {
		request = NewExportFileTamperRulesRequest()
	}
	response = NewExportFileTamperRulesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
	request = &SearchLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SearchLog")
	return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
	response = &SearchLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
	if request == nil {
		request = NewSearchLogRequest()
	}
	response = NewSearchLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefencePluginStatusRequest() (request *DescribeVulDefencePluginStatusRequest) {
	request = &DescribeVulDefencePluginStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginStatus")
	return
}

func NewDescribeVulDefencePluginStatusResponse() (response *DescribeVulDefencePluginStatusResponse) {
	response = &DescribeVulDefencePluginStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各主机漏洞防御插件状态
func (c *Client) DescribeVulDefencePluginStatus(request *DescribeVulDefencePluginStatusRequest) (response *DescribeVulDefencePluginStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefencePluginStatusRequest()
	}
	response = NewDescribeVulDefencePluginStatusResponse()
	err = c.Send(request, response)
	return
}

func NewChangeRuleEventsIgnoreStatusRequest() (request *ChangeRuleEventsIgnoreStatusRequest) {
	request = &ChangeRuleEventsIgnoreStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ChangeRuleEventsIgnoreStatus")
	return
}

func NewChangeRuleEventsIgnoreStatusResponse() (response *ChangeRuleEventsIgnoreStatusResponse) {
	response = &ChangeRuleEventsIgnoreStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据检测项id或事件id批量忽略事件或取消忽略
func (c *Client) ChangeRuleEventsIgnoreStatus(request *ChangeRuleEventsIgnoreStatusRequest) (response *ChangeRuleEventsIgnoreStatusResponse, err error) {
	if request == nil {
		request = NewChangeRuleEventsIgnoreStatusRequest()
	}
	response = NewChangeRuleEventsIgnoreStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebPageProtectSwitchRequest() (request *ModifyWebPageProtectSwitchRequest) {
	request = &ModifyWebPageProtectSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectSwitch")
	return
}

func NewModifyWebPageProtectSwitchResponse() (response *ModifyWebPageProtectSwitchResponse) {
	response = &ModifyWebPageProtectSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网站防篡改防护设置开关
func (c *Client) ModifyWebPageProtectSwitch(request *ModifyWebPageProtectSwitchRequest) (response *ModifyWebPageProtectSwitchResponse, err error) {
	if request == nil {
		request = NewModifyWebPageProtectSwitchRequest()
	}
	response = NewModifyWebPageProtectSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaWorkloadLabelRequest() (request *CreateMaWorkloadLabelRequest) {
	request = &CreateMaWorkloadLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateMaWorkloadLabel")
	return
}

func NewCreateMaWorkloadLabelResponse() (response *CreateMaWorkloadLabelResponse) {
	response = &CreateMaWorkloadLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建微告警负载标签
func (c *Client) CreateMaWorkloadLabel(request *CreateMaWorkloadLabelRequest) (response *CreateMaWorkloadLabelResponse, err error) {
	if request == nil {
		request = NewCreateMaWorkloadLabelRequest()
	}
	response = NewCreateMaWorkloadLabelResponse()
	err = c.Send(request, response)
	return
}

func NewSyncBaselineDetectSummaryRequest() (request *SyncBaselineDetectSummaryRequest) {
	request = &SyncBaselineDetectSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "SyncBaselineDetectSummary")
	return
}

func NewSyncBaselineDetectSummaryResponse() (response *SyncBaselineDetectSummaryResponse) {
	response = &SyncBaselineDetectSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步基线检测进度概要
func (c *Client) SyncBaselineDetectSummary(request *SyncBaselineDetectSummaryRequest) (response *SyncBaselineDetectSummaryResponse, err error) {
	if request == nil {
		request = NewSyncBaselineDetectSummaryRequest()
	}
	response = NewSyncBaselineDetectSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceSwitchRequest() (request *ModifyVulDefenceSwitchRequest) {
	request = &ModifyVulDefenceSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyVulDefenceSwitch")
	return
}

func NewModifyVulDefenceSwitchResponse() (response *ModifyVulDefenceSwitchResponse) {
	response = &ModifyVulDefenceSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全中心漏洞防御、网络攻击开关修改
func (c *Client) ModifyVulDefenceSwitch(request *ModifyVulDefenceSwitchRequest) (response *ModifyVulDefenceSwitchResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceSwitchRequest()
	}
	response = NewModifyVulDefenceSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
	request = &UntrustMalwaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "UntrustMalwares")
	return
}

func NewUntrustMalwaresResponse() (response *UntrustMalwaresResponse) {
	response = &UntrustMalwaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UntrustMalwares）用于取消信任木马文件。
func (c *Client) UntrustMalwares(request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
	if request == nil {
		request = NewUntrustMalwaresRequest()
	}
	response = NewUntrustMalwaresResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetJarInfoRequest() (request *DescribeAssetJarInfoRequest) {
	request = &DescribeAssetJarInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetJarInfo")
	return
}

func NewDescribeAssetJarInfoResponse() (response *DescribeAssetJarInfoResponse) {
	response = &DescribeAssetJarInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Jar包详情
func (c *Client) DescribeAssetJarInfo(request *DescribeAssetJarInfoRequest) (response *DescribeAssetJarInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetJarInfoRequest()
	}
	response = NewDescribeAssetJarInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineOsListRequest() (request *DescribeMachineOsListRequest) {
	request = &DescribeMachineOsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineOsList")
	return
}

func NewDescribeMachineOsListResponse() (response *DescribeMachineOsListResponse) {
	response = &DescribeMachineOsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可筛选操作系统列表.
func (c *Client) DescribeMachineOsList(request *DescribeMachineOsListRequest) (response *DescribeMachineOsListResponse, err error) {
	if request == nil {
		request = NewDescribeMachineOsListRequest()
	}
	response = NewDescribeMachineOsListResponse()
	err = c.Send(request, response)
	return
}

func NewExportSecurityTrendsRequest() (request *ExportSecurityTrendsRequest) {
	request = &ExportSecurityTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportSecurityTrends")
	return
}

func NewExportSecurityTrendsResponse() (response *ExportSecurityTrendsResponse) {
	response = &ExportSecurityTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出风险趋势
func (c *Client) ExportSecurityTrends(request *ExportSecurityTrendsRequest) (response *ExportSecurityTrendsResponse, err error) {
	if request == nil {
		request = NewExportSecurityTrendsRequest()
	}
	response = NewExportSecurityTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebHookPolicyRequest() (request *DeleteWebHookPolicyRequest) {
	request = &DeleteWebHookPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookPolicy")
	return
}

func NewDeleteWebHookPolicyResponse() (response *DeleteWebHookPolicyResponse) {
	response = &DeleteWebHookPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警策略
func (c *Client) DeleteWebHookPolicy(request *DeleteWebHookPolicyRequest) (response *DeleteWebHookPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteWebHookPolicyRequest()
	}
	response = NewDeleteWebHookPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebHookRuleRequest() (request *DeleteWebHookRuleRequest) {
	request = &DeleteWebHookRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookRule")
	return
}

func NewDeleteWebHookRuleResponse() (response *DeleteWebHookRuleResponse) {
	response = &DeleteWebHookRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除企微机器人规则
func (c *Client) DeleteWebHookRule(request *DeleteWebHookRuleRequest) (response *DeleteWebHookRuleResponse, err error) {
	if request == nil {
		request = NewDeleteWebHookRuleRequest()
	}
	response = NewDeleteWebHookRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebHookRuleRequest() (request *DescribeWebHookRuleRequest) {
	request = &DescribeWebHookRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookRule")
	return
}

func NewDescribeWebHookRuleResponse() (response *DescribeWebHookRuleResponse) {
	response = &DescribeWebHookRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企微机器人规则详情
func (c *Client) DescribeWebHookRule(request *DescribeWebHookRuleRequest) (response *DescribeWebHookRuleResponse, err error) {
	if request == nil {
		request = NewDescribeWebHookRuleRequest()
	}
	response = NewDescribeWebHookRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBashPolicyParamsRequest() (request *CheckBashPolicyParamsRequest) {
	request = &CheckBashPolicyParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CheckBashPolicyParams")
	return
}

func NewCheckBashPolicyParamsResponse() (response *CheckBashPolicyParamsResponse) {
	response = &CheckBashPolicyParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验高危命令用户规则新增和编辑时的参数。
func (c *Client) CheckBashPolicyParams(request *CheckBashPolicyParamsRequest) (response *CheckBashPolicyParamsResponse, err error) {
	if request == nil {
		request = NewCheckBashPolicyParamsRequest()
	}
	response = NewCheckBashPolicyParamsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIncidentBacktrackingRequest() (request *CreateIncidentBacktrackingRequest) {
	request = &CreateIncidentBacktrackingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateIncidentBacktracking")
	return
}

func NewCreateIncidentBacktrackingResponse() (response *CreateIncidentBacktrackingResponse) {
	response = &CreateIncidentBacktrackingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发事件调查回溯
func (c *Client) CreateIncidentBacktracking(request *CreateIncidentBacktrackingRequest) (response *CreateIncidentBacktrackingResponse, err error) {
	if request == nil {
		request = NewCreateIncidentBacktrackingRequest()
	}
	response = NewCreateIncidentBacktrackingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
	request = &DescribeComponentStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeComponentStatistics")
	return
}

func NewDescribeComponentStatisticsResponse() (response *DescribeComponentStatisticsResponse) {
	response = &DescribeComponentStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeComponentStatistics) 用于获取组件统计列表数据。
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeComponentStatisticsRequest()
	}
	response = NewDescribeComponentStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRansomDefenseStrategyRequest() (request *CreateRansomDefenseStrategyRequest) {
	request = &CreateRansomDefenseStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateRansomDefenseStrategy")
	return
}

func NewCreateRansomDefenseStrategyResponse() (response *CreateRansomDefenseStrategyResponse) {
	response = &CreateRansomDefenseStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或修改防勒索策略
func (c *Client) CreateRansomDefenseStrategy(request *CreateRansomDefenseStrategyRequest) (response *CreateRansomDefenseStrategyResponse, err error) {
	if request == nil {
		request = NewCreateRansomDefenseStrategyRequest()
	}
	response = NewCreateRansomDefenseStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineDetailRequest() (request *DescribeAssetMachineDetailRequest) {
	request = &DescribeAssetMachineDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineDetail")
	return
}

func NewDescribeAssetMachineDetailResponse() (response *DescribeAssetMachineDetailResponse) {
	response = &DescribeAssetMachineDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理主机资源详细信息
func (c *Client) DescribeAssetMachineDetail(request *DescribeAssetMachineDetailRequest) (response *DescribeAssetMachineDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineDetailRequest()
	}
	response = NewDescribeAssetMachineDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellInfoRequest() (request *DescribeJavaMemShellInfoRequest) {
	request = &DescribeJavaMemShellInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellInfo")
	return
}

func NewDescribeJavaMemShellInfoResponse() (response *DescribeJavaMemShellInfoResponse) {
	response = &DescribeJavaMemShellInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java内存马事件详细信息
func (c *Client) DescribeJavaMemShellInfo(request *DescribeJavaMemShellInfoRequest) (response *DescribeJavaMemShellInfoResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellInfoRequest()
	}
	response = NewDescribeJavaMemShellInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebHookRulesRequest() (request *DescribeWebHookRulesRequest) {
	request = &DescribeWebHookRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookRules")
	return
}

func NewDescribeWebHookRulesResponse() (response *DescribeWebHookRulesResponse) {
	response = &DescribeWebHookRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企微机器人规则列表
func (c *Client) DescribeWebHookRules(request *DescribeWebHookRulesRequest) (response *DescribeWebHookRulesResponse, err error) {
	if request == nil {
		request = NewDescribeWebHookRulesRequest()
	}
	response = NewDescribeWebHookRulesResponse()
	err = c.Send(request, response)
	return
}

func NewFixBaselineDetectRequest() (request *FixBaselineDetectRequest) {
	request = &FixBaselineDetectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "FixBaselineDetect")
	return
}

func NewFixBaselineDetectResponse() (response *FixBaselineDetectResponse) {
	response = &FixBaselineDetectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修复基线检测
func (c *Client) FixBaselineDetect(request *FixBaselineDetectRequest) (response *FixBaselineDetectResponse, err error) {
	if request == nil {
		request = NewFixBaselineDetectRequest()
	}
	response = NewFixBaselineDetectResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLicenseBindsRequest() (request *ModifyLicenseBindsRequest) {
	request = &ModifyLicenseBindsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseBinds")
	return
}

func NewModifyLicenseBindsResponse() (response *ModifyLicenseBindsResponse) {
	response = &ModifyLicenseBindsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置中心-授权管理 对某个授权批量绑定机器
func (c *Client) ModifyLicenseBinds(request *ModifyLicenseBindsRequest) (response *ModifyLicenseBindsResponse, err error) {
	if request == nil {
		request = NewModifyLicenseBindsRequest()
	}
	response = NewModifyLicenseBindsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineTopRequest() (request *DescribeBaselineTopRequest) {
	request = &DescribeBaselineTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineTop")
	return
}

func NewDescribeBaselineTopResponse() (response *DescribeBaselineTopResponse) {
	response = &DescribeBaselineTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据策略id查询基线检测项TOP
func (c *Client) DescribeBaselineTop(request *DescribeBaselineTopRequest) (response *DescribeBaselineTopResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineTopRequest()
	}
	response = NewDescribeBaselineTopResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaSelfLearnPolicyRequest() (request *DeleteMaSelfLearnPolicyRequest) {
	request = &DeleteMaSelfLearnPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaSelfLearnPolicy")
	return
}

func NewDeleteMaSelfLearnPolicyResponse() (response *DeleteMaSelfLearnPolicyResponse) {
	response = &DeleteMaSelfLearnPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微告警自学习子策略
func (c *Client) DeleteMaSelfLearnPolicy(request *DeleteMaSelfLearnPolicyRequest) (response *DeleteMaSelfLearnPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteMaSelfLearnPolicyRequest()
	}
	response = NewDeleteMaSelfLearnPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareWhiteListRequest() (request *DescribeMalwareWhiteListRequest) {
	request = &DescribeMalwareWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareWhiteList")
	return
}

func NewDescribeMalwareWhiteListResponse() (response *DescribeMalwareWhiteListResponse) {
	response = &DescribeMalwareWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取木马白名单列表
func (c *Client) DescribeMalwareWhiteList(request *DescribeMalwareWhiteListRequest) (response *DescribeMalwareWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareWhiteListRequest()
	}
	response = NewDescribeMalwareWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaWhitelistRequest() (request *ExportMaWhitelistRequest) {
	request = &ExportMaWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaWhitelist")
	return
}

func NewExportMaWhitelistResponse() (response *ExportMaWhitelistResponse) {
	response = &ExportMaWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警全局白名单
func (c *Client) ExportMaWhitelist(request *ExportMaWhitelistRequest) (response *ExportMaWhitelistResponse, err error) {
	if request == nil {
		request = NewExportMaWhitelistRequest()
	}
	response = NewExportMaWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineTagRequest() (request *DeleteMachineTagRequest) {
	request = &DeleteMachineTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachineTag")
	return
}

func NewDeleteMachineTagResponse() (response *DeleteMachineTagResponse) {
	response = &DeleteMachineTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务器关联的标签
func (c *Client) DeleteMachineTag(request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
	if request == nil {
		request = NewDeleteMachineTagRequest()
	}
	response = NewDeleteMachineTagResponse()
	err = c.Send(request, response)
	return
}

func NewExportReverseShellEventsRequest() (request *ExportReverseShellEventsRequest) {
	request = &ExportReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportReverseShellEvents")
	return
}

func NewExportReverseShellEventsResponse() (response *ExportReverseShellEventsResponse) {
	response = &ExportReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出反弹Shell事件
func (c *Client) ExportReverseShellEvents(request *ExportReverseShellEventsRequest) (response *ExportReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewExportReverseShellEventsRequest()
	}
	response = NewExportReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginWhiteRecordRequest() (request *ModifyLoginWhiteRecordRequest) {
	request = &ModifyLoginWhiteRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLoginWhiteRecord")
	return
}

func NewModifyLoginWhiteRecordResponse() (response *ModifyLoginWhiteRecordResponse) {
	response = &ModifyLoginWhiteRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新合并后登录审计白名单信息（服务器列表数目应小于1000）
func (c *Client) ModifyLoginWhiteRecord(request *ModifyLoginWhiteRecordRequest) (response *ModifyLoginWhiteRecordResponse, err error) {
	if request == nil {
		request = NewModifyLoginWhiteRecordRequest()
	}
	response = NewModifyLoginWhiteRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaWhitelistRequest() (request *DescribeMaWhitelistRequest) {
	request = &DescribeMaWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaWhitelist")
	return
}

func NewDescribeMaWhitelistResponse() (response *DescribeMaWhitelistResponse) {
	response = &DescribeMaWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警全局白名单
func (c *Client) DescribeMaWhitelist(request *DescribeMaWhitelistRequest) (response *DescribeMaWhitelistResponse, err error) {
	if request == nil {
		request = NewDescribeMaWhitelistRequest()
	}
	response = NewDescribeMaWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulEmergentMsgRequest() (request *DescribeVulEmergentMsgRequest) {
	request = &DescribeVulEmergentMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEmergentMsg")
	return
}

func NewDescribeVulEmergentMsgResponse() (response *DescribeVulEmergentMsgResponse) {
	response = &DescribeVulEmergentMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞紧急通知
func (c *Client) DescribeVulEmergentMsg(request *DescribeVulEmergentMsgRequest) (response *DescribeVulEmergentMsgResponse, err error) {
	if request == nil {
		request = NewDescribeVulEmergentMsgRequest()
	}
	response = NewDescribeVulEmergentMsgResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineItemListRequest() (request *ExportBaselineItemListRequest) {
	request = &ExportBaselineItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineItemList")
	return
}

func NewExportBaselineItemListResponse() (response *ExportBaselineItemListResponse) {
	response = &ExportBaselineItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出检测项结果列表
func (c *Client) ExportBaselineItemList(request *ExportBaselineItemListRequest) (response *ExportBaselineItemListResponse, err error) {
	if request == nil {
		request = NewExportBaselineItemListRequest()
	}
	response = NewExportBaselineItemListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetJarListRequest() (request *DescribeAssetJarListRequest) {
	request = &DescribeAssetJarListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetJarList")
	return
}

func NewDescribeAssetJarListResponse() (response *DescribeAssetJarListResponse) {
	response = &DescribeAssetJarListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Jar包列表
func (c *Client) DescribeAssetJarList(request *DescribeAssetJarListRequest) (response *DescribeAssetJarListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetJarListRequest()
	}
	response = NewDescribeAssetJarListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetScanHostDetailRequest() (request *DescribeAssetScanHostDetailRequest) {
	request = &DescribeAssetScanHostDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetScanHostDetail")
	return
}

func NewDescribeAssetScanHostDetailResponse() (response *DescribeAssetScanHostDetailResponse) {
	response = &DescribeAssetScanHostDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产扫描单台主机的任务状态
func (c *Client) DescribeAssetScanHostDetail(request *DescribeAssetScanHostDetailRequest) (response *DescribeAssetScanHostDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetScanHostDetailRequest()
	}
	response = NewDescribeAssetScanHostDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRansomDefenseMachineListRequest() (request *DescribeRansomDefenseMachineListRequest) {
	request = &DescribeRansomDefenseMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseMachineList")
	return
}

func NewDescribeRansomDefenseMachineListResponse() (response *DescribeRansomDefenseMachineListResponse) {
	response = &DescribeRansomDefenseMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询备份详情列表
func (c *Client) DescribeRansomDefenseMachineList(request *DescribeRansomDefenseMachineListRequest) (response *DescribeRansomDefenseMachineListResponse, err error) {
	if request == nil {
		request = NewDescribeRansomDefenseMachineListRequest()
	}
	response = NewDescribeRansomDefenseMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulInfoRequest() (request *ExportVulInfoRequest) {
	request = &ExportVulInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulInfo")
	return
}

func NewExportVulInfoResponse() (response *ExportVulInfoResponse) {
	response = &ExportVulInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞信息，包括影响主机列表，组件信息
func (c *Client) ExportVulInfo(request *ExportVulInfoRequest) (response *ExportVulInfoResponse, err error) {
	if request == nil {
		request = NewExportVulInfoRequest()
	}
	response = NewExportVulInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScreenMachinesRequest() (request *DescribeScreenMachinesRequest) {
	request = &DescribeScreenMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenMachines")
	return
}

func NewDescribeScreenMachinesResponse() (response *DescribeScreenMachinesResponse) {
	response = &DescribeScreenMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 大屏可视化主机区域列表
func (c *Client) DescribeScreenMachines(request *DescribeScreenMachinesRequest) (response *DescribeScreenMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeScreenMachinesRequest()
	}
	response = NewDescribeScreenMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineDownloadListRequest() (request *DescribeBaselineDownloadListRequest) {
	request = &DescribeBaselineDownloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDownloadList")
	return
}

func NewDescribeBaselineDownloadListResponse() (response *DescribeBaselineDownloadListResponse) {
	response = &DescribeBaselineDownloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线下载列表
func (c *Client) DescribeBaselineDownloadList(request *DescribeBaselineDownloadListRequest) (response *DescribeBaselineDownloadListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineDownloadListRequest()
	}
	response = NewDescribeBaselineDownloadListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskDnsPolicyStatusRequest() (request *ModifyRiskDnsPolicyStatusRequest) {
	request = &ModifyRiskDnsPolicyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskDnsPolicyStatus")
	return
}

func NewModifyRiskDnsPolicyStatusResponse() (response *ModifyRiskDnsPolicyStatusResponse) {
	response = &ModifyRiskDnsPolicyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改恶意请求策略状态
func (c *Client) ModifyRiskDnsPolicyStatus(request *ModifyRiskDnsPolicyStatusRequest) (response *ModifyRiskDnsPolicyStatusResponse, err error) {
	if request == nil {
		request = NewModifyRiskDnsPolicyStatusRequest()
	}
	response = NewModifyRiskDnsPolicyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogExportRequest() (request *CreateLogExportRequest) {
	request = &CreateLogExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateLogExport")
	return
}

func NewCreateLogExportResponse() (response *CreateLogExportResponse) {
	response = &CreateLogExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志下载任务
func (c *Client) CreateLogExport(request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
	if request == nil {
		request = NewCreateLogExportRequest()
	}
	response = NewCreateLogExportResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrivilegeEventsRequest() (request *DeletePrivilegeEventsRequest) {
	request = &DeletePrivilegeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeletePrivilegeEvents")
	return
}

func NewDeletePrivilegeEventsResponse() (response *DeletePrivilegeEventsResponse) {
	response = &DeletePrivilegeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据Ids删除本地提权
func (c *Client) DeletePrivilegeEvents(request *DeletePrivilegeEventsRequest) (response *DeletePrivilegeEventsResponse, err error) {
	if request == nil {
		request = NewDeletePrivilegeEventsRequest()
	}
	response = NewDeletePrivilegeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogClsDeliverInfoRequest() (request *ModifyLogClsDeliverInfoRequest) {
	request = &ModifyLogClsDeliverInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogClsDeliverInfo")
	return
}

func NewModifyLogClsDeliverInfoResponse() (response *ModifyLogClsDeliverInfoResponse) {
	response = &ModifyLogClsDeliverInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志cls投递配置
func (c *Client) ModifyLogClsDeliverInfo(request *ModifyLogClsDeliverInfoRequest) (response *ModifyLogClsDeliverInfoResponse, err error) {
	if request == nil {
		request = NewModifyLogClsDeliverInfoRequest()
	}
	response = NewModifyLogClsDeliverInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSearchLogRequest() (request *CreateSearchLogRequest) {
	request = &CreateSearchLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateSearchLog")
	return
}

func NewCreateSearchLogResponse() (response *CreateSearchLogResponse) {
	response = &CreateSearchLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加历史搜索记录
func (c *Client) CreateSearchLog(request *CreateSearchLogRequest) (response *CreateSearchLogResponse, err error) {
	if request == nil {
		request = NewCreateSearchLogRequest()
	}
	response = NewCreateSearchLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
	request = &DeleteTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteTags")
	return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
	response = &DeleteTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除标签
func (c *Client) DeleteTags(request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
	if request == nil {
		request = NewDeleteTagsRequest()
	}
	response = NewDeleteTagsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLicenseOrderRequest() (request *ModifyLicenseOrderRequest) {
	request = &ModifyLicenseOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseOrder")
	return
}

func NewModifyLicenseOrderResponse() (response *ModifyLicenseOrderResponse) {
	response = &ModifyLicenseOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑授权订单
func (c *Client) ModifyLicenseOrder(request *ModifyLicenseOrderRequest) (response *ModifyLicenseOrderResponse, err error) {
	if request == nil {
		request = NewModifyLicenseOrderRequest()
	}
	response = NewModifyLicenseOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchLogsRequest() (request *DescribeSearchLogsRequest) {
	request = &DescribeSearchLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchLogs")
	return
}

func NewDescribeSearchLogsResponse() (response *DescribeSearchLogsResponse) {
	response = &DescribeSearchLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取历史搜索记录
func (c *Client) DescribeSearchLogs(request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSearchLogsRequest()
	}
	response = NewDescribeSearchLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecommendedProtectCpuRequest() (request *DescribeRecommendedProtectCpuRequest) {
	request = &DescribeRecommendedProtectCpuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRecommendedProtectCpu")
	return
}

func NewDescribeRecommendedProtectCpuResponse() (response *DescribeRecommendedProtectCpuResponse) {
	response = &DescribeRecommendedProtectCpuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询推荐购买防护核数
func (c *Client) DescribeRecommendedProtectCpu(request *DescribeRecommendedProtectCpuRequest) (response *DescribeRecommendedProtectCpuResponse, err error) {
	if request == nil {
		request = NewDescribeRecommendedProtectCpuRequest()
	}
	response = NewDescribeRecommendedProtectCpuResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaLabelSourceListRequest() (request *DescribeMaLabelSourceListRequest) {
	request = &DescribeMaLabelSourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaLabelSourceList")
	return
}

func NewDescribeMaLabelSourceListResponse() (response *DescribeMaLabelSourceListResponse) {
	response = &DescribeMaLabelSourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警标签来源列表
func (c *Client) DescribeMaLabelSourceList(request *DescribeMaLabelSourceListRequest) (response *DescribeMaLabelSourceListResponse, err error) {
	if request == nil {
		request = NewDescribeMaLabelSourceListRequest()
	}
	response = NewDescribeMaLabelSourceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemInfoRequest() (request *DescribeBaselineItemInfoRequest) {
	request = &DescribeBaselineItemInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemInfo")
	return
}

func NewDescribeBaselineItemInfoResponse() (response *DescribeBaselineItemInfoResponse) {
	response = &DescribeBaselineItemInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测项信息
func (c *Client) DescribeBaselineItemInfo(request *DescribeBaselineItemInfoRequest) (response *DescribeBaselineItemInfoResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemInfoRequest()
	}
	response = NewDescribeBaselineItemInfoResponse()
	err = c.Send(request, response)
	return
}

func NewTestWebHookRuleRequest() (request *TestWebHookRuleRequest) {
	request = &TestWebHookRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "TestWebHookRule")
	return
}

func NewTestWebHookRuleResponse() (response *TestWebHookRuleResponse) {
	response = &TestWebHookRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试企微机器人规则
func (c *Client) TestWebHookRule(request *TestWebHookRuleRequest) (response *TestWebHookRuleResponse, err error) {
	if request == nil {
		request = NewTestWebHookRuleRequest()
	}
	response = NewTestWebHookRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchTemplatesRequest() (request *DescribeSearchTemplatesRequest) {
	request = &DescribeSearchTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchTemplates")
	return
}

func NewDescribeSearchTemplatesResponse() (response *DescribeSearchTemplatesResponse) {
	response = &DescribeSearchTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取快速检索列表
func (c *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSearchTemplatesRequest()
	}
	response = NewDescribeSearchTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFileTamperRuleRequest() (request *CheckFileTamperRuleRequest) {
	request = &CheckFileTamperRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CheckFileTamperRule")
	return
}

func NewCheckFileTamperRuleResponse() (response *CheckFileTamperRuleResponse) {
	response = &CheckFileTamperRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检验核心文件监控前端新增和编辑时的规则参数。
func (c *Client) CheckFileTamperRule(request *CheckFileTamperRuleRequest) (response *CheckFileTamperRuleResponse, err error) {
	if request == nil {
		request = NewCheckFileTamperRuleRequest()
	}
	response = NewCheckFileTamperRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareRiskOverviewRequest() (request *DescribeMalwareRiskOverviewRequest) {
	request = &DescribeMalwareRiskOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareRiskOverview")
	return
}

func NewDescribeMalwareRiskOverviewResponse() (response *DescribeMalwareRiskOverviewResponse) {
	response = &DescribeMalwareRiskOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取文件查杀概览信息
func (c *Client) DescribeMalwareRiskOverview(request *DescribeMalwareRiskOverviewRequest) (response *DescribeMalwareRiskOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareRiskOverviewRequest()
	}
	response = NewDescribeMalwareRiskOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperEventRuleInfoRequest() (request *DescribeFileTamperEventRuleInfoRequest) {
	request = &DescribeFileTamperEventRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperEventRuleInfo")
	return
}

func NewDescribeFileTamperEventRuleInfoResponse() (response *DescribeFileTamperEventRuleInfoResponse) {
	response = &DescribeFileTamperEventRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看产生事件时规则详情接口
func (c *Client) DescribeFileTamperEventRuleInfo(request *DescribeFileTamperEventRuleInfoRequest) (response *DescribeFileTamperEventRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperEventRuleInfoRequest()
	}
	response = NewDescribeFileTamperEventRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessInfoListRequest() (request *DescribeAssetProcessInfoListRequest) {
	request = &DescribeAssetProcessInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetProcessInfoList")
	return
}

func NewDescribeAssetProcessInfoListResponse() (response *DescribeAssetProcessInfoListResponse) {
	response = &DescribeAssetProcessInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理进程列表
func (c *Client) DescribeAssetProcessInfoList(request *DescribeAssetProcessInfoListRequest) (response *DescribeAssetProcessInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessInfoListRequest()
	}
	response = NewDescribeAssetProcessInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewExportRiskDnsEventListRequest() (request *ExportRiskDnsEventListRequest) {
	request = &ExportRiskDnsEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskDnsEventList")
	return
}

func NewExportRiskDnsEventListResponse() (response *ExportRiskDnsEventListResponse) {
	response = &ExportRiskDnsEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出恶意请求事件列表
func (c *Client) ExportRiskDnsEventList(request *ExportRiskDnsEventListRequest) (response *ExportRiskDnsEventListResponse, err error) {
	if request == nil {
		request = NewExportRiskDnsEventListRequest()
	}
	response = NewExportRiskDnsEventListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMaWorkloadMonitorRequest() (request *ModifyMaWorkloadMonitorRequest) {
	request = &ModifyMaWorkloadMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaWorkloadMonitor")
	return
}

func NewModifyMaWorkloadMonitorResponse() (response *ModifyMaWorkloadMonitorResponse) {
	response = &ModifyMaWorkloadMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改微告警负载监控
func (c *Client) ModifyMaWorkloadMonitor(request *ModifyMaWorkloadMonitorRequest) (response *ModifyMaWorkloadMonitorResponse, err error) {
	if request == nil {
		request = NewModifyMaWorkloadMonitorRequest()
	}
	response = NewModifyMaWorkloadMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetRecentMachineInfoRequest() (request *DescribeAssetRecentMachineInfoRequest) {
	request = &DescribeAssetRecentMachineInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetRecentMachineInfo")
	return
}

func NewDescribeAssetRecentMachineInfoResponse() (response *DescribeAssetRecentMachineInfoResponse) {
	response = &DescribeAssetRecentMachineInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机最近趋势情况
func (c *Client) DescribeAssetRecentMachineInfo(request *DescribeAssetRecentMachineInfoRequest) (response *DescribeAssetRecentMachineInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetRecentMachineInfoRequest()
	}
	response = NewDescribeAssetRecentMachineInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineWeakPasswordListRequest() (request *DescribeBaselineWeakPasswordListRequest) {
	request = &DescribeBaselineWeakPasswordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineWeakPasswordList")
	return
}

func NewDescribeBaselineWeakPasswordListResponse() (response *DescribeBaselineWeakPasswordListResponse) {
	response = &DescribeBaselineWeakPasswordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线弱口令列表
func (c *Client) DescribeBaselineWeakPasswordList(request *DescribeBaselineWeakPasswordListRequest) (response *DescribeBaselineWeakPasswordListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineWeakPasswordListRequest()
	}
	response = NewDescribeBaselineWeakPasswordListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityEventStatRequest() (request *DescribeSecurityEventStatRequest) {
	request = &DescribeSecurityEventStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityEventStat")
	return
}

func NewDescribeSecurityEventStatResponse() (response *DescribeSecurityEventStatResponse) {
	response = &DescribeSecurityEventStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全事件统计
func (c *Client) DescribeSecurityEventStat(request *DescribeSecurityEventStatRequest) (response *DescribeSecurityEventStatResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityEventStatRequest()
	}
	response = NewDescribeSecurityEventStatResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebServiceInfoListRequest() (request *ExportAssetWebServiceInfoListRequest) {
	request = &ExportAssetWebServiceInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebServiceInfoList")
	return
}

func NewExportAssetWebServiceInfoListResponse() (response *ExportAssetWebServiceInfoListResponse) {
	response = &ExportAssetWebServiceInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web服务列表
func (c *Client) ExportAssetWebServiceInfoList(request *ExportAssetWebServiceInfoListRequest) (response *ExportAssetWebServiceInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebServiceInfoListRequest()
	}
	response = NewExportAssetWebServiceInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaSummaryRequest() (request *DescribeMaSummaryRequest) {
	request = &DescribeMaSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaSummary")
	return
}

func NewDescribeMaSummaryResponse() (response *DescribeMaSummaryResponse) {
	response = &DescribeMaSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微告警整体概要
func (c *Client) DescribeMaSummary(request *DescribeMaSummaryRequest) (response *DescribeMaSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeMaSummaryRequest()
	}
	response = NewDescribeMaSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMalwareWhiteListRequest() (request *CreateMalwareWhiteListRequest) {
	request = &CreateMalwareWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateMalwareWhiteList")
	return
}

func NewCreateMalwareWhiteListResponse() (response *CreateMalwareWhiteListResponse) {
	response = &CreateMalwareWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建木马白名单
func (c *Client) CreateMalwareWhiteList(request *CreateMalwareWhiteListRequest) (response *CreateMalwareWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateMalwareWhiteListRequest()
	}
	response = NewCreateMalwareWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineEffectHostListRequest() (request *DescribeBaselineEffectHostListRequest) {
	request = &DescribeBaselineEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineEffectHostList")
	return
}

func NewDescribeBaselineEffectHostListResponse() (response *DescribeBaselineEffectHostListResponse) {
	response = &DescribeBaselineEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线id查询基线影响主机列表
func (c *Client) DescribeBaselineEffectHostList(request *DescribeBaselineEffectHostListRequest) (response *DescribeBaselineEffectHostListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineEffectHostListRequest()
	}
	response = NewDescribeBaselineEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
	request = &DescribeIndexListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeIndexList")
	return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
	response = &DescribeIndexListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取索引列表
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
	if request == nil {
		request = NewDescribeIndexListRequest()
	}
	response = NewDescribeIndexListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulCveIdInfoRequest() (request *DescribeVulCveIdInfoRequest) {
	request = &DescribeVulCveIdInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulCveIdInfo")
	return
}

func NewDescribeVulCveIdInfoResponse() (response *DescribeVulCveIdInfoResponse) {
	response = &DescribeVulCveIdInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CveId查询漏洞详情
func (c *Client) DescribeVulCveIdInfo(request *DescribeVulCveIdInfoRequest) (response *DescribeVulCveIdInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVulCveIdInfoRequest()
	}
	response = NewDescribeVulCveIdInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleCategoryListRequest() (request *DescribeBaselineRuleCategoryListRequest) {
	request = &DescribeBaselineRuleCategoryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleCategoryList")
	return
}

func NewDescribeBaselineRuleCategoryListResponse() (response *DescribeBaselineRuleCategoryListResponse) {
	response = &DescribeBaselineRuleCategoryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线分类列表
func (c *Client) DescribeBaselineRuleCategoryList(request *DescribeBaselineRuleCategoryListRequest) (response *DescribeBaselineRuleCategoryListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleCategoryListRequest()
	}
	response = NewDescribeBaselineRuleCategoryListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicProxyInstallCommandRequest() (request *DescribePublicProxyInstallCommandRequest) {
	request = &DescribePublicProxyInstallCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribePublicProxyInstallCommand")
	return
}

func NewDescribePublicProxyInstallCommandResponse() (response *DescribePublicProxyInstallCommandResponse) {
	response = &DescribePublicProxyInstallCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公网接入代理安装命令
func (c *Client) DescribePublicProxyInstallCommand(request *DescribePublicProxyInstallCommandRequest) (response *DescribePublicProxyInstallCommandResponse, err error) {
	if request == nil {
		request = NewDescribePublicProxyInstallCommandRequest()
	}
	response = NewDescribePublicProxyInstallCommandResponse()
	err = c.Send(request, response)
	return
}

func NewExportMaCustomPolicyListRequest() (request *ExportMaCustomPolicyListRequest) {
	request = &ExportMaCustomPolicyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportMaCustomPolicyList")
	return
}

func NewExportMaCustomPolicyListResponse() (response *ExportMaCustomPolicyListResponse) {
	response = &ExportMaCustomPolicyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出微告警自定义白名单
func (c *Client) ExportMaCustomPolicyList(request *ExportMaCustomPolicyListRequest) (response *ExportMaCustomPolicyListResponse, err error) {
	if request == nil {
		request = NewExportMaCustomPolicyListRequest()
	}
	response = NewExportMaCustomPolicyListResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineRuleDetectListRequest() (request *ExportBaselineRuleDetectListRequest) {
	request = &ExportBaselineRuleDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineRuleDetectList")
	return
}

func NewExportBaselineRuleDetectListResponse() (response *ExportBaselineRuleDetectListResponse) {
	response = &ExportBaselineRuleDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线检测规则
func (c *Client) ExportBaselineRuleDetectList(request *ExportBaselineRuleDetectListRequest) (response *ExportBaselineRuleDetectListResponse, err error) {
	if request == nil {
		request = NewExportBaselineRuleDetectListRequest()
	}
	response = NewExportBaselineRuleDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsAccountStatusRequest() (request *DescribeClsAccountStatusRequest) {
	request = &DescribeClsAccountStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeClsAccountStatus")
	return
}

func NewDescribeClsAccountStatusResponse() (response *DescribeClsAccountStatusResponse) {
	response = &DescribeClsAccountStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cls账户状态
func (c *Client) DescribeClsAccountStatus(request *DescribeClsAccountStatusRequest) (response *DescribeClsAccountStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClsAccountStatusRequest()
	}
	response = NewDescribeClsAccountStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebHookRuleStatusRequest() (request *ModifyWebHookRuleStatusRequest) {
	request = &ModifyWebHookRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookRuleStatus")
	return
}

func NewModifyWebHookRuleStatusResponse() (response *ModifyWebHookRuleStatusResponse) {
	response = &ModifyWebHookRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改企微机器人规则状态
func (c *Client) ModifyWebHookRuleStatus(request *ModifyWebHookRuleStatusRequest) (response *ModifyWebHookRuleStatusResponse, err error) {
	if request == nil {
		request = NewModifyWebHookRuleStatusRequest()
	}
	response = NewModifyWebHookRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetEnvListRequest() (request *ExportAssetEnvListRequest) {
	request = &ExportAssetEnvListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetEnvList")
	return
}

func NewExportAssetEnvListResponse() (response *ExportAssetEnvListResponse) {
	response = &ExportAssetEnvListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理环境变量列表
func (c *Client) ExportAssetEnvList(request *ExportAssetEnvListRequest) (response *ExportAssetEnvListResponse, err error) {
	if request == nil {
		request = NewExportAssetEnvListRequest()
	}
	response = NewExportAssetEnvListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBaselineRuleIgnoreRequest() (request *ModifyBaselineRuleIgnoreRequest) {
	request = &ModifyBaselineRuleIgnoreRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineRuleIgnore")
	return
}

func NewModifyBaselineRuleIgnoreResponse() (response *ModifyBaselineRuleIgnoreResponse) {
	response = &ModifyBaselineRuleIgnoreResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改基线忽略规则
func (c *Client) ModifyBaselineRuleIgnore(request *ModifyBaselineRuleIgnoreRequest) (response *ModifyBaselineRuleIgnoreResponse, err error) {
	if request == nil {
		request = NewModifyBaselineRuleIgnoreRequest()
	}
	response = NewModifyBaselineRuleIgnoreResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemIgnoreListRequest() (request *DescribeBaselineItemIgnoreListRequest) {
	request = &DescribeBaselineItemIgnoreListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemIgnoreList")
	return
}

func NewDescribeBaselineItemIgnoreListResponse() (response *DescribeBaselineItemIgnoreListResponse) {
	response = &DescribeBaselineItemIgnoreListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取忽略规则项列表
func (c *Client) DescribeBaselineItemIgnoreList(request *DescribeBaselineItemIgnoreListRequest) (response *DescribeBaselineItemIgnoreListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemIgnoreListRequest()
	}
	response = NewDescribeBaselineItemIgnoreListResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulEffectHostListRequest() (request *ExportVulEffectHostListRequest) {
	request = &ExportVulEffectHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulEffectHostList")
	return
}

func NewExportVulEffectHostListResponse() (response *ExportVulEffectHostListResponse) {
	response = &ExportVulEffectHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞影响主机列表
func (c *Client) ExportVulEffectHostList(request *ExportVulEffectHostListRequest) (response *ExportVulEffectHostListResponse, err error) {
	if request == nil {
		request = NewExportVulEffectHostListRequest()
	}
	response = NewExportVulEffectHostListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBruteAttackRulesRequest() (request *ModifyBruteAttackRulesRequest) {
	request = &ModifyBruteAttackRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyBruteAttackRules")
	return
}

func NewModifyBruteAttackRulesResponse() (response *ModifyBruteAttackRulesResponse) {
	response = &ModifyBruteAttackRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改暴力破解规则
func (c *Client) ModifyBruteAttackRules(request *ModifyBruteAttackRulesRequest) (response *ModifyBruteAttackRulesResponse, err error) {
	if request == nil {
		request = NewModifyBruteAttackRulesRequest()
	}
	response = NewModifyBruteAttackRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogKafkaStateRequest() (request *ModifyLogKafkaStateRequest) {
	request = &ModifyLogKafkaStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaState")
	return
}

func NewModifyLogKafkaStateResponse() (response *ModifyLogKafkaStateResponse) {
	response = &ModifyLogKafkaStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志投递状态信息
func (c *Client) ModifyLogKafkaState(request *ModifyLogKafkaStateRequest) (response *ModifyLogKafkaStateResponse, err error) {
	if request == nil {
		request = NewModifyLogKafkaStateRequest()
	}
	response = NewModifyLogKafkaStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsInfoRequest() (request *DescribeRiskDnsInfoRequest) {
	request = &DescribeRiskDnsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsInfo")
	return
}

func NewDescribeRiskDnsInfoResponse() (response *DescribeRiskDnsInfoResponse) {
	response = &DescribeRiskDnsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求详情
func (c *Client) DescribeRiskDnsInfo(request *DescribeRiskDnsInfoRequest) (response *DescribeRiskDnsInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsInfoRequest()
	}
	response = NewDescribeRiskDnsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportJavaMemShellPluginsRequest() (request *ExportJavaMemShellPluginsRequest) {
	request = &ExportJavaMemShellPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportJavaMemShellPlugins")
	return
}

func NewExportJavaMemShellPluginsResponse() (response *ExportJavaMemShellPluginsResponse) {
	response = &ExportJavaMemShellPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出java内存马插件信息
func (c *Client) ExportJavaMemShellPlugins(request *ExportJavaMemShellPluginsRequest) (response *ExportJavaMemShellPluginsResponse, err error) {
	if request == nil {
		request = NewExportJavaMemShellPluginsRequest()
	}
	response = NewExportJavaMemShellPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefencePluginEventRequest() (request *ExportVulDefencePluginEventRequest) {
	request = &ExportVulDefencePluginEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefencePluginEvent")
	return
}

func NewExportVulDefencePluginEventResponse() (response *ExportVulDefencePluginEventResponse) {
	response = &ExportVulDefencePluginEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞防御插件事件
func (c *Client) ExportVulDefencePluginEvent(request *ExportVulDefencePluginEventRequest) (response *ExportVulDefencePluginEventResponse, err error) {
	if request == nil {
		request = NewExportVulDefencePluginEventRequest()
	}
	response = NewExportVulDefencePluginEventResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBaselineStrategyRequest() (request *DeleteBaselineStrategyRequest) {
	request = &DeleteBaselineStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineStrategy")
	return
}

func NewDeleteBaselineStrategyResponse() (response *DeleteBaselineStrategyResponse) {
	response = &DeleteBaselineStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据基线策略id删除策略
func (c *Client) DeleteBaselineStrategy(request *DeleteBaselineStrategyRequest) (response *DeleteBaselineStrategyResponse, err error) {
	if request == nil {
		request = NewDeleteBaselineStrategyRequest()
	}
	response = NewDeleteBaselineStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBuyBindTaskRequest() (request *CreateBuyBindTaskRequest) {
	request = &CreateBuyBindTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CreateBuyBindTask")
	return
}

func NewCreateBuyBindTaskResponse() (response *CreateBuyBindTaskResponse) {
	response = &CreateBuyBindTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新购授权自动绑定任务
func (c *Client) CreateBuyBindTask(request *CreateBuyBindTaskRequest) (response *CreateBuyBindTaskResponse, err error) {
	if request == nil {
		request = NewCreateBuyBindTaskRequest()
	}
	response = NewCreateBuyBindTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashEventsNewRequest() (request *DescribeBashEventsNewRequest) {
	request = &DescribeBashEventsNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsNew")
	return
}

func NewDescribeBashEventsNewResponse() (response *DescribeBashEventsNewResponse) {
	response = &DescribeBashEventsNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取高危命令列表(新)
func (c *Client) DescribeBashEventsNew(request *DescribeBashEventsNewRequest) (response *DescribeBashEventsNewResponse, err error) {
	if request == nil {
		request = NewDescribeBashEventsNewRequest()
	}
	response = NewDescribeBashEventsNewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRiskDnsPolicyRequest() (request *ModifyRiskDnsPolicyRequest) {
	request = &ModifyRiskDnsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskDnsPolicy")
	return
}

func NewModifyRiskDnsPolicyResponse() (response *ModifyRiskDnsPolicyResponse) {
	response = &ModifyRiskDnsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改恶意请求策略
func (c *Client) ModifyRiskDnsPolicy(request *ModifyRiskDnsPolicyRequest) (response *ModifyRiskDnsPolicyResponse, err error) {
	if request == nil {
		request = NewModifyRiskDnsPolicyRequest()
	}
	response = NewModifyRiskDnsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousRequestWhiteListRequest() (request *DescribeMaliciousRequestWhiteListRequest) {
	request = &DescribeMaliciousRequestWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaliciousRequestWhiteList")
	return
}

func NewDescribeMaliciousRequestWhiteListResponse() (response *DescribeMaliciousRequestWhiteListResponse) {
	response = &DescribeMaliciousRequestWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求白名单列表
func (c *Client) DescribeMaliciousRequestWhiteList(request *DescribeMaliciousRequestWhiteListRequest) (response *DescribeMaliciousRequestWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousRequestWhiteListRequest()
	}
	response = NewDescribeMaliciousRequestWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTypeTopRequest() (request *DescribeAssetTypeTopRequest) {
	request = &DescribeAssetTypeTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTypeTop")
	return
}

func NewDescribeAssetTypeTopResponse() (response *DescribeAssetTypeTopResponse) {
	response = &DescribeAssetTypeTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各种类型资源Top5
func (c *Client) DescribeAssetTypeTop(request *DescribeAssetTypeTopRequest) (response *DescribeAssetTypeTopResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTypeTopRequest()
	}
	response = NewDescribeAssetTypeTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashEventsInfoRequest() (request *DescribeBashEventsInfoRequest) {
	request = &DescribeBashEventsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsInfo")
	return
}

func NewDescribeBashEventsInfoResponse() (response *DescribeBashEventsInfoResponse) {
	response = &DescribeBashEventsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询高危命令事件详情
func (c *Client) DescribeBashEventsInfo(request *DescribeBashEventsInfoRequest) (response *DescribeBashEventsInfoResponse, err error) {
	if request == nil {
		request = NewDescribeBashEventsInfoRequest()
	}
	response = NewDescribeBashEventsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceOverviewCountRequest() (request *DescribeVulDefenceOverviewCountRequest) {
	request = &DescribeVulDefenceOverviewCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceOverviewCount")
	return
}

func NewDescribeVulDefenceOverviewCountResponse() (response *DescribeVulDefenceOverviewCountResponse) {
	response = &DescribeVulDefenceOverviewCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御策略和事件统计
func (c *Client) DescribeVulDefenceOverviewCount(request *DescribeVulDefenceOverviewCountRequest) (response *DescribeVulDefenceOverviewCountResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceOverviewCountRequest()
	}
	response = NewDescribeVulDefenceOverviewCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogKafkaAccessRequest() (request *ModifyLogKafkaAccessRequest) {
	request = &ModifyLogKafkaAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaAccess")
	return
}

func NewModifyLogKafkaAccessResponse() (response *ModifyLogKafkaAccessResponse) {
	response = &ModifyLogKafkaAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改日志投递kafka接入配置
func (c *Client) ModifyLogKafkaAccess(request *ModifyLogKafkaAccessRequest) (response *ModifyLogKafkaAccessResponse, err error) {
	if request == nil {
		request = NewModifyLogKafkaAccessRequest()
	}
	response = NewModifyLogKafkaAccessResponse()
	err = c.Send(request, response)
	return
}

func NewCancelIgnoreVulRequest() (request *CancelIgnoreVulRequest) {
	request = &CancelIgnoreVulRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cwp", APIVersion, "CancelIgnoreVul")
	return
}

func NewCancelIgnoreVulResponse() (response *CancelIgnoreVulResponse) {
	response = &CancelIgnoreVulResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消漏洞忽略
func (c *Client) CancelIgnoreVul(request *CancelIgnoreVulRequest) (response *CancelIgnoreVulResponse, err error) {
	if request == nil {
		request = NewCancelIgnoreVulRequest()
	}
	response = NewCancelIgnoreVulResponse()
	err = c.Send(request, response)
	return
}

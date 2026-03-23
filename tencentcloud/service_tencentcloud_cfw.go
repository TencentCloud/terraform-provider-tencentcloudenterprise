package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type CfwService struct {
	client *connectivity.TencentCloudClient
}

func (me *CfwService) DescribeNatFwInstancesInfoById(ctx context.Context, instanceId string) (
	natInsInfo *cfw.NatInstanceInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatFwInstancesInfoRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	filter := cfw.NatFwFilter{
		FilterType:    helper.String("NatinsId"),
		FilterContent: helper.String(instanceId),
	}

	request.Filter = []*cfw.NatFwFilter{&filter}
	request.Limit = helper.IntInt64(20)
	request.Offset = helper.IntInt64(0)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeNatFwInstancesInfo(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response.NatinsLst != nil && len(response.Response.NatinsLst) > 0 {
		natInsInfo = response.Response.NatinsLst[0]
	}

	return
}

func (me *CfwService) DescribeNatFwVpcDnsLstById(ctx context.Context, instanceId string) (
	vpcIds []string, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatFwVpcDnsLstRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.NatFwInsId = helper.String(instanceId)
	request.Limit = helper.IntInt64(100) // Set a reasonable limit
	request.Offset = helper.IntInt64(0)

	ratelimit.Check(request.GetAction())

	// Fetch all pages if needed
	vpcIds = make([]string, 0)
	for {
		response, e := me.client.UseCfwClient().DescribeNatFwVpcDnsLst(request)
		if e != nil {
			err = e
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		// Extract VPC IDs from VpcDnsInfo objects
		for _, vpcDns := range response.Response.VpcDnsSwitchLst {
			if vpcDns.VpcId != nil {
				vpcIds = append(vpcIds, *vpcDns.VpcId)
			}
		}

		// Check if there are more pages
		if response.Response.Total == nil {
			break
		}

		total := *response.Response.Total
		currentCount := int64(len(vpcIds))

		if currentCount >= total {
			// All records fetched
			break
		}

		// Fetch next page
		request.Offset = &currentCount
	}

	return
}

func (me *CfwService) DescribeCfwEipsById(ctx context.Context, instanceId string) (
	natGwIds []string, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeCfwEipsRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.CfwInstance = helper.String(instanceId)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeCfwEips(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	// Extract NAT Gateway IDs from NatFwEipsInfo objects
	natGwIds = make([]string, 0, len(response.Response.NatFwEipList))
	for _, eipInfo := range response.Response.NatFwEipList {
		if eipInfo.NatGatewayId != nil {
			natGwIds = append(natGwIds, *eipInfo.NatGatewayId)
		}
	}

	return
}

func (me *CfwService) ModifyNatInstance(ctx context.Context, instanceId, instanceName string) (
	err error) {

	logId := getLogId(ctx)
	request := cfw.NewModifyNatInstanceRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.NatInstanceId = helper.String(instanceId)
	request.InstanceName = helper.String(instanceName)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().ModifyNatInstance(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DeleteNatFwInstanceById(ctx context.Context, instanceId string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteNatFwInstanceRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.CfwInstance = helper.String(instanceId)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteNatFwInstance(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DescribeVpcFwGroupInstanceById(ctx context.Context, fwGroupId string) (
	vpcFwInstance *cfw.VpcFwGroupInfo, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeFwGroupInstanceInfoRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{fwGroupId}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntInt64(20)
	request.Offset = helper.IntInt64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeFwGroupInstanceInfo(request)
	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.VpcFwGroupLst) < 1 {
		return
	}

	vpcFwInstance = response.Response.VpcFwGroupLst[0]
	return
}

// DeleteCfwVpcInstanceById
func (me *CfwService) DeleteCfwVpcInstanceById(ctx context.Context, fwGroupId string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteVpcFwGroupRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	// First, get the firewall group info to retrieve instance IDs
	describeRequest := cfw.NewDescribeFwGroupInstanceInfoRequest()
	describeRequest.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{fwGroupId}),
			OperatorType: helper.Int64(1),
		},
	}
	describeRequest.Limit = helper.IntInt64(20)
	describeRequest.Offset = helper.IntInt64(0)

	ratelimit.Check(describeRequest.GetAction())

	describeResponse, err := me.client.UseCfwClient().DescribeFwGroupInstanceInfo(describeRequest)
	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, describeRequest.GetAction(), describeRequest.ToJsonString(), describeResponse.ToJsonString())

	// Build VpcFwInsList from the firewall instances
	vpcFwInsList := []*string{}
	if len(describeResponse.Response.VpcFwGroupLst) > 0 && describeResponse.Response.VpcFwGroupLst[0].FwInstanceLst != nil {
		for _, instance := range describeResponse.Response.VpcFwGroupLst[0].FwInstanceLst {
			if instance.FwInsId != nil {
				vpcFwInsList = append(vpcFwInsList, instance.FwInsId)
			}
		}
	}

	// Set delete parameters
	request.FwGroupId = helper.String(fwGroupId)
	request.VpcFwInsList = vpcFwInsList
	request.DeleteFwGroup = helper.Int64(1) // Delete entire firewall group

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteVpcFwGroup(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeNatFwPolicyById
func (me *CfwService) DescribeNatFwPolicyById(ctx context.Context, uuid string) (
	natPolicy *cfw.DescAcItem, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeNatAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeNatAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	natPolicy = response.Response.Data[0]

	return
}

// DeleteNatFwPolicyById ...
func (me *CfwService) DeleteNatFwPolicyById(ctx context.Context, uuid string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewRemoveNatAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()
	uuidInt, _ := strconv.ParseInt(uuid, 10, 64)

	request.RuleUuid = []*int64{&uuidInt}

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().RemoveNatAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeVpcFwPolicyById
func (me *CfwService) DescribeVpcFwPolicyById(ctx context.Context, uuid string) (
	vpcPolicy *cfw.VpcRuleItem, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeVpcAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeVpcAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	vpcPolicy = response.Response.Data[0]

	return
}

// DeleteVpcFwPolicyById
func (me *CfwService) DeleteVpcFwPolicyById(ctx context.Context, uuid string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewRemoveVpcAcRuleRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	uuidInt, _ := strconv.ParseInt(uuid, 10, 64)

	request.RuleUuids = []*int64{&uuidInt}

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().RemoveVpcAcRule(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeBlockIgnoreListById
func (me *CfwService) DescribeBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (
	blockIgnoreList *cfw.BlockIgnoreRule, err error) {

	logId := getLogId(ctx)
	request := cfw.NewDescribeBlockIgnoreListRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	var searchStr string
	if iP != "" {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, iP)
	} else {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, domain)
	}

	request.Limit = helper.IntInt64(20)
	request.Offset = helper.IntInt64(0)
	request.SearchValue = &searchStr
	request.Direction = &direction
	ruleTypeInt, _ := strconv.ParseUint(ruleType, 10, 64)
	request.RuleType = &ruleTypeInt
	request.By = helper.String("EndTime")
	request.Order = helper.String("desc")

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeBlockIgnoreList(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	blockIgnoreList = response.Response.Data[0]

	return
}

// DeleteBlockIgnoreListById ...
func (me *CfwService) DeleteBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (err error) {
	logId := getLogId(ctx)
	request := cfw.NewDeleteBlockIgnoreRuleListRequest()

	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	directionInt, _ := strconv.ParseInt(direction, 10, 64)
	if iP != "" {
		request.Rules = []*cfw.IocListData{
			{
				IP:        helper.String(iP),
				Direction: helper.Int64(directionInt),
			},
		}
	} else {
		request.Rules = []*cfw.IocListData{
			{
				Domain:    helper.String(domain),
				Direction: helper.Int64(directionInt),
			},
		}
	}

	ruleTypeInt, _ := strconv.ParseInt(ruleType, 10, 64)
	request.RuleType = helper.Int64(ruleTypeInt)

	ratelimit.Check(request.GetAction())

	_, err = me.client.UseCfwClient().DeleteBlockIgnoreRuleList(request)

	if err != nil {
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DescribeCfwEdgeFwSwitchesByFilter(ctx context.Context) (edgeFwSwitches []*cfw.EdgeIpInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cfw.NewDescribeFwEdgeIpsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset int64 = 0
		limit  int64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseCfwClient().DescribeFwEdgeIps(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Data) < 1 {
			break
		}

		edgeFwSwitches = append(edgeFwSwitches, response.Response.Data...)
		if len(response.Response.Data) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *CfwService) DescribeCfwNatFwSwitchesByFilter(ctx context.Context, param map[string]interface{}) (natFwSwitches []*cfw.NatSwitchListData, errRet error) {
	var (
		logId    = getLogId(ctx)
		request  = cfw.NewDescribeNatFwSwitchRequest()
		response = cfw.NewDescribeNatFwSwitchResponse()
		filters  []*cfw.CommonFilter
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "NatInsId" {
			filters = append(filters, &cfw.CommonFilter{
				Name:         helper.String("NatInsId"),
				OperatorType: helper.Int64(1),
				Values:       helper.Strings([]string{v.(string)}),
			})
		}

		if k == "Status" {
			filters = append(filters, &cfw.CommonFilter{
				Name:         helper.String("Status"),
				OperatorType: helper.Int64(1),
				Values:       helper.Strings([]string{helper.IntToStr(v.(int))}),
			})
		}

		if k == "Enable" {
			filters = append(filters, &cfw.CommonFilter{
				Name:         helper.String("Enable"),
				OperatorType: helper.Int64(1),
				Values:       helper.Strings([]string{helper.IntToStr(v.(int))}),
			})
		}
	}

	if len(filters) > 0 {
		request.Filters = filters
	}

	var (
		offset int64 = 0
		limit  int64 = 20
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.UseCfwClient().DescribeNatFwSwitch(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			if result == nil || result.Response == nil || result.Response.Data == nil {
				return resource.NonRetryableError(fmt.Errorf("Describe nat firewall switch failed, Response is nil."))
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if len(response.Response.Data) < 1 {
			break
		}

		natFwSwitches = append(natFwSwitches, response.Response.Data...)
		if len(response.Response.Data) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *CfwService) DescribeCfwVpcFwSwitchesByFilter(ctx context.Context, vpcInsId string) (vpcFirewallSwitch []*cfw.FwGroupSwitchShow, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeVpcFwGroupSwitchRequest()
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{vpcInsId}),
			OperatorType: helper.Int64(1),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseCfwClient().DescribeVpcFwGroupSwitch(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.SwitchList) < 1 {
			break
		}

		vpcFirewallSwitch = append(vpcFirewallSwitch, response.Response.SwitchList...)
		if len(response.Response.SwitchList) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

// DescribeCfwAddressTemplateById
func (me *CfwService) DescribeCfwAddressTemplateById(ctx context.Context, uuid string) (addressTemplate *cfw.TemplateListInfo, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeAddressTemplateListRequest()
	request.Uuid = &uuid

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeAddressTemplateList(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	addressTemplate = response.Response.Data[0]
	return
}

// DeleteCfwAddressTemplateById
func (me *CfwService) DeleteCfwAddressTemplateById(ctx context.Context, uuid string) (errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDeleteAddressTemplateRequest()
	request.Uuid = &uuid

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DeleteAddressTemplate(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

// DescribeCfwBlockIgnoreListById
func (me *CfwService) DescribeCfwBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (blockIgnoreRule *cfw.BlockIgnoreRule, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeBlockIgnoreListRequest()
	var searchStr string
	if iP != "" {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, iP)
	} else {
		searchStr = fmt.Sprintf(`{"domain":"%s"}`, domain)
	}

	request.Limit = helper.Int64(20)
	request.Offset = helper.Int64(0)
	request.SearchValue = &searchStr
	request.Direction = &direction
	ruleTypeInt, _ := strconv.ParseUint(ruleType, 10, 64)
	request.RuleType = &ruleTypeInt
	request.By = helper.String("EndTime")
	request.Order = helper.String("desc")

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeBlockIgnoreList(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	blockIgnoreRule = response.Response.Data[0]
	return
}

// DescribeCfwEdgeFirewallSwitchById
func (me *CfwService) DescribeCfwEdgeFirewallSwitchById(ctx context.Context, publicIp string) (edgeFirewallSwitch *cfw.EdgeIpInfo, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeFwEdgeIpsRequest()
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("PublicIp"),
			Values:       helper.Strings([]string{publicIp}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.Int64(10)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeFwEdgeIps(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	edgeFirewallSwitch = response.Response.Data[0]
	return
}

// DescribeCfwEdgePolicyById
func (me *CfwService) DescribeCfwEdgePolicyById(ctx context.Context, uuid string) (edgePolicy *cfw.DescAcItem, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeAclRuleRequest()
	request.Limit = helper.Uint64(20)
	request.Offset = helper.Uint64(0)
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeAclRule(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	edgePolicy = response.Response.Data[0]
	return
}

// DeleteCfwEdgePolicyById
func (me *CfwService) DeleteCfwEdgePolicyById(ctx context.Context, uuid string) (errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewRemoveAclRuleRequest()
	uuidInt, _ := strconv.ParseInt(uuid, 10, 64)
	request.RuleUuid = []*int64{&uuidInt}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseCfwClient().RemoveAclRule(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

// DescribeCfwNatFirewallFwSwitchById
func (me *CfwService) DescribeCfwNatFirewallFwSwitchById(ctx context.Context, natInsId, subnetId string) (natFirewallSwitch *cfw.NatSwitchListData, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeNatFwSwitchRequest()
	response := cfw.NewDescribeNatFwSwitchResponse()
	request.Offset = helper.Int64(0)
	request.Limit = helper.Int64(20)
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("NatInsId"),
			OperatorType: helper.Int64(1),
			Values:       helper.Strings([]string{natInsId}),
		},
		{
			Name:         helper.String("SubnetId"),
			OperatorType: helper.Int64(1),
			Values:       helper.Strings([]string{subnetId}),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	result, err := me.client.UseCfwClient().DescribeNatFwSwitch(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())

	if result == nil || result.Response == nil || result.Response.Data == nil {
		errRet = fmt.Errorf("Describe nat firewall switch failed, Response is nil.")
		return
	}

	response = result

	if len(response.Response.Data) < 1 {
		return
	}

	natFirewallSwitch = response.Response.Data[0]
	return
}

// DescribeCfwNatPolicyById
func (me *CfwService) DescribeCfwNatPolicyById(ctx context.Context, uuid string) (natPolicy *cfw.DescAcItem, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeNatAcRuleRequest()
	response := cfw.NewDescribeNatAcRuleResponse()
	request.Limit = helper.Uint64(20)
	request.Offset = helper.Uint64(0)
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("Id"),
			Values:       helper.Strings([]string{uuid}),
			OperatorType: helper.Int64(1),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	result, err := me.client.UseCfwClient().DescribeNatAcRule(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if result == nil || result.Response == nil || result.Response.Data == nil {
		errRet = fmt.Errorf("Describe nat ac rule failed, Response is nil.")
		return
	}

	response = result

	if len(response.Response.Data) < 1 {
		return
	}

	natPolicy = response.Response.Data[0]
	return
}

// DescribeCfwVpcFirewallSwitchById
func (me *CfwService) DescribeCfwVpcFirewallSwitchById(ctx context.Context, vpcInsId, switchId string) (vpcFirewallSwitch *cfw.FwGroupSwitchShow, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeVpcFwGroupSwitchRequest()
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("SwitchId"),
			Values:       helper.Strings([]string{switchId}),
			OperatorType: helper.Int64(1),
		},
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{vpcInsId}),
			OperatorType: helper.Int64(1),
		},
	}
	request.Limit = helper.Uint64(20)
	request.Offset = helper.Uint64(0)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeVpcFwGroupSwitch(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.SwitchList) < 1 {
		return
	}

	vpcFirewallSwitch = response.Response.SwitchList[0]
	return
}

// DeleteCfwBlockIgnoreListById wrapper for compatibility
func (me *CfwService) DeleteCfwBlockIgnoreListById(ctx context.Context, iP, domain, direction, ruleType string) (errRet error) {
	return me.DeleteBlockIgnoreListById(ctx, iP, domain, direction, ruleType)
}

// DeleteCfwNatPolicyById wrapper for compatibility
func (me *CfwService) DeleteCfwNatPolicyById(ctx context.Context, uuid string) (errRet error) {
	return me.DeleteNatFwPolicyById(ctx, uuid)
}

func (me *CfwService) DescribeSgRuleById(ctx context.Context, ruleUuid string) (ret *cfw.SecurityGroupListData, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeSecurityGroupListRequest()
	request.Limit = helper.Uint64(100)
	request.Offset = helper.Uint64(0)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCfwClient().DescribeSecurityGroupList(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response.Data != nil {
			for _, item := range response.Response.Data {
				if item.Uuid != nil && *item.Uuid == ruleUuid {
					ret = item
					return
				}
			}
		}

		if response.Response.Total == nil || uint64(len(response.Response.Data)) >= *response.Response.Total || len(response.Response.Data) == 0 {
			break
		}
		*request.Offset += *request.Limit
	}

	return
}

func (me *CfwService) DeleteSecurityGroupRule(ctx context.Context, id uint64, area string, direction uint64) (errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDeleteSecurityGroupRuleRequest()
	request.Id = &id
	request.Area = &area
	request.Direction = &direction

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseCfwClient().DeleteSecurityGroupRule(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n", logId, request.GetAction(), request.ToJsonString())

	return
}

func (me *CfwService) DescribeFwGroupInstanceInfoById(ctx context.Context, fwGroupId string) (vpcFwGroupInfo *cfw.VpcFwGroupInfo, errRet error) {
	logId := getLogId(ctx)

	request := cfw.NewDescribeFwGroupInstanceInfoRequest()
	request.Offset = helper.IntInt64(0)
	request.Limit = helper.IntInt64(10)
	request.Filters = []*cfw.CommonFilter{
		{
			Name:         helper.String("FwGroupId"),
			Values:       helper.Strings([]string{fwGroupId}),
			OperatorType: helper.Int64(1),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCfwClient().DescribeFwGroupInstanceInfo(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.VpcFwGroupLst) < 1 {
		return
	}

	vpcFwGroupInfo = response.Response.VpcFwGroupLst[0]
	return
}

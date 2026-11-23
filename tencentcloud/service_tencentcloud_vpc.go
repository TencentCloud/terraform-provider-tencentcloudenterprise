package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	"strconv"
	"strings"
	"sync"
	"terraform-provider-tencentcloudenterprise/sdk/common"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

var eipUnattachLocker = &sync.Mutex{}

/* For Adun Sake please DO NOT Declare the redundant Type STRUCT!! */
// VPC basic information
type VpcBasicInfo struct {
	vpcId                string
	vpcNumId             uint64
	name                 string
	cidr                 string
	isMulticast          bool
	isDefault            bool
	dnsServers           []string
	createTime           string
	tags                 []*vpc.Tag
	assistantCidrs       []string
	dockerAssistantCidrs []string
}

// subnet basic information
type VpcSubnetBasicInfo struct {
	vpcId            string
	subnetId         string
	routeTableId     string
	name             string
	cidr             string
	isMulticast      bool
	isDefault        bool
	zone             string
	availableIpCount int64
	createTime       string
}

type VpcSecurityGroupLiteRule struct {
	action                  string
	cidrIp                  string
	port                    string
	protocol                string
	addressId               string
	addressGroupId          string
	securityGroupId         string
	protocolTemplateId      string
	protocolTemplateGroupId string
}

func (rule VpcSecurityGroupLiteRule) String() string {

	var source string

	if rule.cidrIp != "" {
		source = rule.cidrIp
	}
	if rule.securityGroupId != "" {
		source = rule.securityGroupId
	}
	if rule.addressId != "" {
		source = rule.addressId
	}
	if rule.addressGroupId != "" {
		source = rule.addressGroupId
	}

	protocol := rule.protocol

	if protocol == "" && rule.protocolTemplateId != "" {
		protocol = rule.protocolTemplateId
	} else if protocol == "" && rule.protocolTemplateGroupId != "" {
		protocol = rule.protocolTemplateGroupId
	}

	return fmt.Sprintf("%s#%s#%s#%s", rule.action, source, rule.port, protocol)
}

type VpcService struct {
	client *connectivity.TencentCloudClient
}

// ///////common
func (me *VpcService) fillFilter(ins []*vpc.Filter, key, value string) (outs []*vpc.Filter) {
	if ins == nil {
		ins = make([]*vpc.Filter, 0, 2)
	}

	var filter = vpc.Filter{Name: &key, Values: []*string{&value}}
	ins = append(ins, &filter)
	outs = ins
	return
}

func (me *VpcService) DescribeVpcs(ctx context.Context,
	vpcId, name string,
	tags map[string]string,
	isDefaultPtr *bool,
	tagKey string,
	cidrBlock string) (infos []VpcBasicInfo, errRet error) {
	logId := getLogId(ctx)
	request := vpc.NewDescribeVpcsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	infos = make([]VpcBasicInfo, 0, 100)

	var (
		offset  = 0
		limit   = 100
		total   = -1
		hasVpc  = map[string]bool{}
		filters []*vpc.Filter
	)

	if vpcId != "" {
		filters = me.fillFilter(filters, "vpc-id", vpcId)
	}

	if name != "" {
		filters = me.fillFilter(filters, "vpc-name", name)
	}

	if tagKey != "" {
		filters = me.fillFilter(filters, "tag-key", tagKey)
	}

	if cidrBlock != "" {
		filters = me.fillFilter(filters, "cidr-block", cidrBlock)
	}

	if isDefaultPtr != nil {
		filters = me.fillFilter(filters, "is-default", map[bool]string{true: "true", false: "false"}[*isDefaultPtr])
	}

	for k, v := range tags {
		filters = me.fillFilter(filters, "tag:"+k, v)
	}

	if len(filters) > 0 {
		request.Filters = filters
	}

getMoreData:

	if total >= 0 {
		if offset >= total {
			return
		}
	}
	var strLimit = fmt.Sprintf("%d", limit)
	request.Limit = &strLimit

	var strOffset = fmt.Sprintf("%d", offset)
	request.Offset = &strOffset
	var response *vpc.DescribeVpcsResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseVpcClient().DescribeVpcs(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read vpc failed, reason: %v", logId, err)
		return nil, err
	}

	if total < 0 {
		total = int(*response.Response.TotalCount)
	}

	if len(response.Response.VpcSet) > 0 {
		offset += limit
	} else {
		// get empty VpcInfo, we're done
		return
	}
	for _, item := range response.Response.VpcSet {
		var basicInfo VpcBasicInfo
		basicInfo.cidr = *item.CidrBlock
		basicInfo.createTime = *item.CreatedTime
		basicInfo.dnsServers = make([]string, 0, len(item.DnsServerSet))

		for _, v := range item.DnsServerSet {
			basicInfo.dnsServers = append(basicInfo.dnsServers, *v)
		}
		basicInfo.isDefault = *item.IsDefault
		basicInfo.isMulticast = *item.EnableMulticast
		basicInfo.name = *item.VpcName
		basicInfo.vpcId = *item.VpcId
		if item.VpcNumId != nil {
			basicInfo.vpcNumId = *item.VpcNumId
		}

		if hasVpc[basicInfo.vpcId] {
			errRet = fmt.Errorf("get repeated vpc_id[%s] when doing DescribeVpcs", basicInfo.vpcId)
			return
		}
		hasVpc[basicInfo.vpcId] = true

		if len(item.AssistantCidrSet) > 0 {
			for i := range item.AssistantCidrSet {
				kind := item.AssistantCidrSet[i].AssistantType
				cidr := item.AssistantCidrSet[i].CidrBlock
				if kind != nil && *kind == 0 {
					basicInfo.assistantCidrs = append(basicInfo.assistantCidrs, *cidr)
				} else {
					basicInfo.dockerAssistantCidrs = append(basicInfo.dockerAssistantCidrs, *cidr)
				}
			}
		}

		if len(item.TagSet) > 0 {
			basicInfo.tags = item.TagSet
		}

		infos = append(infos, basicInfo)
	}
	goto getMoreData

}

func (me *VpcService) DescribeSubnets(ctx context.Context,
	subnetId,
	vpcId,
	subnetName,
	zone string,
	tags map[string]string,
	isDefaultPtr *bool,
	isRemoteVpcSNAT *bool,
	tagKey,
	cidrBlock string) (infos []VpcSubnetBasicInfo, errRet error) {

	logId := getLogId(ctx)
	request := vpc.NewDescribeSubnetsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		offset    = 0
		limit     = 100
		total     = -1
		hasSubnet = map[string]bool{}
		filters   []*vpc.Filter
	)

	// 注意： 不加这个filter， 接口无返回
	filters = append(filters, &vpc.Filter{
		Name: common.StringPtr("type"),
		Values: []*string{
			common.StringPtr("0"),
			common.StringPtr("9"),
			common.StringPtr("10"),
		},
	})

	if subnetId != "" {
		filters = me.fillFilter(filters, "subnet-id", subnetId)
	}
	if vpcId != "" {
		filters = me.fillFilter(filters, "vpc-id", vpcId)
	}
	if subnetName != "" {
		filters = me.fillFilter(filters, "subnet-name", subnetName)
	}
	if zone != "" {
		filters = me.fillFilter(filters, "zone", zone)
	}

	if isDefaultPtr != nil {
		filters = me.fillFilter(filters, "is-default", map[bool]string{true: "true", false: "false"}[*isDefaultPtr])
	}

	if isRemoteVpcSNAT != nil {
		filters = me.fillFilter(filters, "is-remote-vpc-snat", map[bool]string{true: "true", false: "false"}[*isRemoteVpcSNAT])
	}

	if tagKey != "" {
		filters = me.fillFilter(filters, "tag-key", tagKey)
	}
	if cidrBlock != "" {
		filters = me.fillFilter(filters, "cidr-block", cidrBlock)
	}

	for k, v := range tags {
		filters = me.fillFilter(filters, "tag:"+k, v)
	}

	if len(filters) > 0 {
		request.Filters = filters
	}

getMoreData:
	if total >= 0 {
		if offset >= total {
			return
		}
	}
	var strLimit = fmt.Sprintf("%d", limit)
	request.Limit = &strLimit

	var strOffset = fmt.Sprintf("%d", offset)
	request.Offset = &strOffset
	var response *vpc.DescribeSubnetsResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseVpcClient().DescribeSubnets(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read subnets failed, reason: %v", logId, err)
		return nil, err
	}
	log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if total < 0 {
		total = int(*response.Response.TotalCount)
	}

	if len(response.Response.SubnetSet) > 0 {
		offset += limit
	} else {
		// get empty subnet, we're done
		return
	}
	for _, item := range response.Response.SubnetSet {
		var basicInfo VpcSubnetBasicInfo

		basicInfo.cidr = *item.CidrBlock
		basicInfo.createTime = *item.CreatedTime
		basicInfo.vpcId = *item.VpcId
		basicInfo.subnetId = *item.SubnetId
		basicInfo.routeTableId = *item.RouteTableId

		basicInfo.name = *item.SubnetName
		basicInfo.isDefault = *item.IsDefault
		basicInfo.isMulticast = *item.EnableBroadcast

		basicInfo.zone = *item.Zone
		basicInfo.availableIpCount = int64(*item.AvailableIpAddressCount)

		if hasSubnet[basicInfo.subnetId] {
			errRet = fmt.Errorf("get repeated subnetId[%s] when doing DescribeSubnets", basicInfo.subnetId)
			return
		}
		hasSubnet[basicInfo.subnetId] = true
		infos = append(infos, basicInfo)
	}
	goto getMoreData
}

/*
EIP
*/
func (me *VpcService) DescribeEipById(ctx context.Context, eipId string) (eip *vpc.Address, errRet error) {
	logId := getLogId(ctx)
	request := vpc.NewDescribeAddressesRequest()
	request.AddressIds = []*string{&eipId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().DescribeAddresses(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.AddressSet) < 1 {
		return
	}
	eip = response.Response.AddressSet[0]
	return
}

func (me *VpcService) AttachEip(ctx context.Context, eipId, instanceId string) error {
	logId := getLogId(ctx)
	request := vpc.NewAssociateAddressRequest()
	request.AddressId = &eipId
	request.InstanceId = &instanceId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().AssociateAddress(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	// Wait for task completion if TaskId is returned
	if response.Response.TaskId == nil {
		return nil
	}
	taskId, err := strconv.ParseUint(*response.Response.TaskId, 10, 64)
	if err != nil {
		return nil
	}

	taskRequest := vpc.NewDescribeTaskResultRequest()
	taskRequest.TaskId = &taskId
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(taskRequest.GetAction())
		taskResponse, err := me.client.UseVpcClient().DescribeTaskResult(taskRequest)
		if err != nil {
			return retryError(err)
		}
		if taskResponse.Response.Result != nil && *taskResponse.Response.Result == EIP_TASK_STATUS_RUNNING {
			return resource.RetryableError(errors.New("eip attach task is running"))
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (me *VpcService) DescribeNatGatewayById(ctx context.Context, natGateWayId string) (natGateWay *vpc.NatGateway, errRet error) {
	logId := getLogId(ctx)
	request := vpc.NewDescribeNatGatewaysRequest()
	request.NatGatewayIds = []*string{&natGateWayId}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().DescribeNatGateways(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.NatGatewaySet) > 0 {
		natGateWay = response.Response.NatGatewaySet[0]
	}

	return
}

func (me *VpcService) DisassociateNatGatewayAddress(ctx context.Context, request *vpc.DisassociateNatGatewayAddressRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	// Check if Nat Gateway Ip still associate
	gateway, err := me.DescribeNatGatewayById(ctx, *request.NatGatewayId)

	if err != nil {
		errRet = err
		return
	}

	if gateway == nil || len(gateway.PublicIpAddressSet) == 0 {
		return
	}

	var gatewayAddresses []string
	var candidates []*string

	for i := range gateway.PublicIpAddressSet {
		addr := gateway.PublicIpAddressSet[i].PublicIpAddress
		gatewayAddresses = append(gatewayAddresses, *addr)
	}

	for i := range request.PublicIpAddresses {
		addr := request.PublicIpAddresses[i]
		if helper.StringsContain(gatewayAddresses, *addr) {
			candidates = append(candidates, addr)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	request.PublicIpAddresses = candidates

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().DisassociateNatGatewayAddress(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *VpcService) UnattachEip(ctx context.Context, eipId string) error {
	eipUnattachLocker.Lock()
	defer eipUnattachLocker.Unlock()

	logId := getLogId(ctx)
	eip, err := me.DescribeEipById(ctx, eipId)
	if err != nil {
		return err
	}
	if eip == nil || *eip.AddressStatus == EIP_STATUS_UNBIND {
		return nil
	}

	// DisassociateAddress Doesn't support Disassociate NAT Address
	if eip.InstanceId != nil && strings.HasPrefix(*eip.InstanceId, "nat-") {
		request := vpc.NewDisassociateNatGatewayAddressRequest()
		request.NatGatewayId = eip.InstanceId
		request.PublicIpAddresses = []*string{eip.AddressIp}
		err := me.DisassociateNatGatewayAddress(ctx, request)
		if err != nil {
			return err
		}

		outErr := resource.Retry(readRetryTimeout*3, func() *resource.RetryError {
			eip, err := me.DescribeEipById(ctx, eipId)
			if err != nil {
				return retryError(err)
			}
			if eip != nil && *eip.AddressStatus != EIP_STATUS_UNBIND {
				return resource.RetryableError(fmt.Errorf("eip is still %s", EIP_STATUS_UNBIND))
			}
			return nil
		})

		if outErr != nil {
			return outErr
		}
		return nil
	}

	// DisassociateAddress Doesn't support Disassociate HAVIP Address
	if (eip.BindResourceType != nil && *eip.BindResourceType == "BIND_HAVIP") ||
		(eip.InstanceId != nil && strings.HasPrefix(*eip.InstanceId, "havip-")) {
		request := vpc.NewHaVipDisassociateAddressIpRequest()
		request.HaVipId = eip.InstanceId
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseVpcClient().HaVipDisassociateAddressIp(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return err
		}

		outErr := resource.Retry(readRetryTimeout*3, func() *resource.RetryError {
			eip, err := me.DescribeEipById(ctx, eipId)
			if err != nil {
				return retryError(err)
			}
			if eip != nil && *eip.AddressStatus != EIP_STATUS_UNBIND {
				return resource.RetryableError(fmt.Errorf("eip is still %s", EIP_STATUS_UNBIND))
			}
			return nil
		})

		if outErr != nil {
			return outErr
		}
		return nil
	}

	request := vpc.NewDisassociateAddressRequest()
	request.AddressId = &eipId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().DisassociateAddress(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	if response.Response.TaskId == nil {
		return nil
	}
	taskId, err := strconv.ParseUint(*response.Response.TaskId, 10, 64)
	if err != nil {
		return nil
	}

	taskRequest := vpc.NewDescribeTaskResultRequest()
	taskRequest.TaskId = &taskId
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(taskRequest.GetAction())
		taskResponse, err := me.client.UseVpcClient().DescribeTaskResult(taskRequest)
		if err != nil {
			return retryError(err)
		}
		if taskResponse.Response.Result != nil && *taskResponse.Response.Result == EIP_TASK_STATUS_RUNNING {
			return resource.RetryableError(errors.New("eip task is running"))
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (me *VpcService) DescribeTaskResult(ctx context.Context, taskId *uint64) (result *vpc.DescribeTaskResultResponse, err error) {

	logId := getLogId(ctx)
	request := vpc.NewDescribeTaskResultRequest()
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail,reason[%s]", logId, request.GetAction(), err.Error())
		}
	}()
	request.TaskId = taskId
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseVpcClient().DescribeTaskResult(request)
		if err != nil {
			return retryError(err)
		}
		result = response
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (me *VpcService) DescribeAddressesByFilter(ctx context.Context, filters map[string]string) (addresses []*vpc.Address, errRet error) {
	logId := getLogId(ctx)
	request := vpc.NewDescribeAddressesRequest()

	if len(filters) > 0 {
		request.Filters = make([]*vpc.Filter, 0, len(filters))
		for k, v := range filters {
			filter := &vpc.Filter{
				Name:   helper.String(k),
				Values: []*string{helper.String(v)},
			}
			request.Filters = append(request.Filters, filter)
		}
	}

	var offset int64 = 0
	var pageSize int64 = 100
	addresses = make([]*vpc.Address, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseVpcClient().DescribeAddresses(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || len(response.Response.AddressSet) < 1 {
			break
		}
		addresses = append(addresses, response.Response.AddressSet...)
		if len(response.Response.AddressSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}
	return
}

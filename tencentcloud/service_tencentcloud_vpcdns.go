package tencentcloud

import (
	"context"
	"log"

	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
)

type VpcDnsService struct {
	client *connectivity.TencentCloudClient
}

func (me *VpcDnsService) DescribeVpcDnsZoneById(ctx context.Context, zoneId string) (
	zone *vpcdns.PrivateZone, errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDescribePrivateZoneRequest()
	request.ZoneId = helper.String(zoneId)

	var response *vpcdns.DescribePrivateZoneResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseVpcDnsClient().DescribePrivateZone(request)
		if err != nil {
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read PrivateDns zone failed, reason: %v", logId, err)
		return nil, err
	}

	if response == nil || response.Response == nil {
		return nil, errors.New("DescribePrivateZone response is nil")
	}

	return response.Response.PrivateZone, nil
}

// CreateVpcDnsForwardRule create vpc dns forward rule
func (me *VpcDnsService) CreateVpcDnsForwardRule(ctx context.Context, remark, domainId string, forwardAddress []string) (
	ruleId string, errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewCreateVpcDnsForwardRuleRequest()
	request.Remark = helper.String(remark)
	request.DomainIdList = []*string{helper.String(domainId)}
	request.ForwardAddress = helper.Strings(forwardAddress)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().CreateVpcDnsForwardRule(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}

	ruleId = *response.Response.RuleIdList[0]
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// DescribeVpcDnsForwardRuleById fetch all forward rules and match by ruleId
func (me *VpcDnsService) DescribeVpcDnsForwardRuleById(ctx context.Context, ruleId string) (
	forwardRule *vpcdns.VpcDnsForwardRuleDetail, errRet error) {
	logId := getLogId(ctx)
	var (
		offset   = 0
		limit    = 20
		ruleList []*vpcdns.VpcDnsForwardRuleDetail
	)
	for {
		request := vpcdns.NewDescribeVpcDnsForwardRuleRequest()
		request.Limit = helper.IntUint64(limit)
		request.Offset = helper.IntUint64(offset)
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseVpcDnsClient().DescribeVpcDnsForwardRule(request)
		if err != nil {
			errRet = errors.WithStack(err)
			return
		}
		ruleList = append(ruleList, response.Response.ForwardRuleList...)
		if len(response.Response.ForwardRuleList) < limit {
			break
		}
		offset += limit
	}
	for _, rule := range ruleList {
		if rule.RuleId != nil && *rule.RuleId == ruleId {
			forwardRule = rule
			break
		}
	}
	log.Printf("[DEBUG]%s api[%s] success, total rules [%d], target ruleId [%s], found [%v]\n",
		logId, "DescribeForwardRuleList", len(ruleList), ruleId, forwardRule != nil)
	return
}

// DescribeVpcDnsForwardRuleAll fetch all forward rules (for data source)
func (me *VpcDnsService) DescribeVpcDnsForwardRuleAll(ctx context.Context) (
	ruleList []*vpcdns.VpcDnsForwardRuleDetail, errRet error) {
	logId := getLogId(ctx)
	var (
		offset = 0
		limit  = 20
	)
	for {
		request := vpcdns.NewDescribeVpcDnsForwardRuleRequest()
		request.Limit = helper.IntUint64(limit)
		request.Offset = helper.IntUint64(offset)
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseVpcDnsClient().DescribeVpcDnsForwardRule(request)
		if err != nil {
			errRet = errors.WithStack(err)
			return
		}
		ruleList = append(ruleList, response.Response.ForwardRuleList...)
		if len(response.Response.ForwardRuleList) < limit {
			break
		}
		offset += limit
	}
	log.Printf("[DEBUG]%s api[%s] success, total rules [%d]\n",
		logId, "DescribeForwardRuleList", len(ruleList))
	return
}

func (me *VpcDnsService) ModifyVpcDnsForwardRule(ctx context.Context, ruleId string, remark string,
	forwardAddress []string) (errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewModifyVpcDnsForwardRuleRequest()
	request.RuleId = helper.String(ruleId)
	request.Remark = helper.String(remark)
	request.ForwardAddress = helper.Strings(forwardAddress)
	ratelimit.Check(request.GetAction())
	_, err := me.client.UseVpcDnsClient().ModifyVpcDnsForwardRule(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n",
		logId, request.GetAction(), request.ToJsonString())
	return
}

func (me *VpcDnsService) DeleteVpcDnsForwardRule(ctx context.Context, ruleId string) (errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDeleteVpcDnsForwardRuleRequest()
	request.RuleIdList = []*string{&ruleId}
	ratelimit.Check(request.GetAction())
	_, err := me.client.UseVpcDnsClient().DeleteVpcDnsForwardRule(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s]\n",
		logId, request.GetAction(), request.ToJsonString())
	return
}

// DescribeVpcDnsZoneList describe zone list with pagination
func (me *VpcDnsService) DescribeVpcDnsZoneList(ctx context.Context,
	filters []*vpcdns.Filter) (zones []*vpcdns.PrivateZone, errRet error) {
	logId := getLogId(ctx)
	var (
		limit  int64 = 20
		offset int64 = 0
	)
	for {
		request := vpcdns.NewDescribePrivateZoneListRequest()
		request.Limit = &limit
		request.Offset = &offset
		request.Filters = filters
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseVpcDnsClient().DescribePrivateZoneList(request)
		if err != nil {
			errRet = errors.WithStack(err)
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || response.Response.PrivateZoneSet == nil {
			break
		}
		zones = append(zones, response.Response.PrivateZoneSet...)
		if int64(len(response.Response.PrivateZoneSet)) < limit {
			break
		}
		offset += limit
	}
	return
}
func (me *VpcDnsService) DescribeVpcDnsZoneRecordByFilter(ctx context.Context, zoneId string,
	recordId string) (recordInfos []*vpcdns.PrivateZoneRecord, errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDescribePrivateZoneRecordListRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	var (
		limit  int64 = 20
		offset int64 = 0
		total  int64 = -1
	)
	request.ZoneId = &zoneId
	request.Filters = make([]*vpcdns.Filter, 0)

	if recordId != "" {
		filter := vpcdns.Filter{
			Name:   helper.String("RecordId"),
			Values: []*string{&recordId},
		}
		request.Filters = append(request.Filters, &filter)
	}

getMoreData:

	if total >= 0 {
		if offset >= total {
			return
		}
	}
	var response *vpcdns.DescribePrivateZoneRecordListResponse

	ratelimit.Check(request.GetAction())
	request.Limit = &limit
	request.Offset = &offset

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseVpcDnsClient().DescribePrivateZoneRecordList(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read vpcdns zone record failed, reason: %v", logId, err)
		return nil, err
	}
	if total < 0 {
		total = *response.Response.TotalCount
	}

	if len(response.Response.RecordSet) > 0 {
		offset = offset + limit
	} else {
		return
	}

	recordInfos = append(recordInfos, response.Response.RecordSet...)
	goto getMoreData
}

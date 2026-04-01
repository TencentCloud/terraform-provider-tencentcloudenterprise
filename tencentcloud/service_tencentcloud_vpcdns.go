package tencentcloud

import (
	"context"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
	"log"
)

type VpcDnsService struct {
	client *connectivity.TencentCloudClient
}

func (me *VpcDnsService) DescribeVpcDnsDomainById(ctx context.Context,
	domainId string) (clbInstance *vpcdns.DomainDetail,
	errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDescribeVpcDnsDomainListRequest()
	request.Filters = []*vpcdns.DomainListFilters{
		{
			Name: helper.String("domain-id"),
			Values: []*string{
				helper.String(domainId),
			},
		},
	}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().DescribeVpcDnsDomainList(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Domains) < 1 {
		return
	}
	clbInstance = response.Response.Domains[0]
	return
}

// ModifyVpcDnsDomain modify vpc dns dnsForWardStatus
func (me *VpcDnsService) ModifyVpcDnsDomain(ctx context.Context,
	domainId string, dnsForWardStatus string) (errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewModifyVpcDnsDomainRequest()
	request.DomainIds = helper.String(domainId)
	request.DnsForwardStatus = helper.String(dnsForWardStatus)
	response, err := me.client.UseVpcDnsClient().ModifyVpcDnsDomain(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

func (me *VpcDnsService) DescribeVpcDnsDomainByFilters(ctx context.Context,
	filters []*vpcdns.DomainListFilters) (clbInstance *vpcdns.DomainDetail, errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDescribeVpcDnsDomainListRequest()
	request.Filters = filters
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().DescribeVpcDnsDomainList(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Domains) < 1 {
		return
	}
	clbInstance = response.Response.Domains[0]
	return
}

func (me *VpcDnsService) DeleteVpcDnsDomain(ctx context.Context, domainId string) error {
	logId := getLogId(ctx)
	request := vpcdns.NewDeleteVpcDnsDomainRequest()
	request.DomainIds = helper.String(domainId)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().DeleteVpcDnsDomain(request)
	if err != nil {
		//if e, ok := err.(*sdkErrors.CloudSDKError); ok {
		//	if e.GetCode() == "InvalidParameter.LBIdNotFound" {
		//		return nil
		//	}
		//}
		return errors.WithStack(err)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *VpcDnsService) CreateVpcDnsDomain(ctx context.Context, domain string, tags map[string]string,
	dnsForwardStatus string) error {
	logId := getLogId(ctx)
	request := vpcdns.NewCreateVpcDnsDomainRequest()
	request.Domain = helper.String(domain)

	if len(tags) > 0 {
		for tagKey, tagValue := range tags {
			tag := vpcdns.Tag{
				Key:   helper.String(tagKey),
				Value: helper.String(tagValue),
			}
			request.Tags = append(request.Tags, &tag)
		}
	}
	request.DnsForwardStatus = helper.String(dnsForwardStatus)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().CreateVpcDnsDomain(request)
	if err != nil {
		//if e, ok := err.(*sdkErrors.CloudSDKError); ok {
		//	if e.GetCode() == "InvalidParameter.LBIdNotFound" {
		//		return nil
		//	}
		//}
		return errors.WithStack(err)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil

}

func (me *VpcDnsService) BindVpcDnsDomain(ctx context.Context, domainId uint64, vpcInfos []*vpcdns.VpcInfos) error {
	logId := getLogId(ctx)
	request := vpcdns.NewBindVpcDnsDomainRequest()
	request.DomainId = helper.Uint64(domainId)
	request.VpcInfos = vpcInfos
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().BindVpcDnsDomain(request)
	if err != nil {
		return errors.WithStack(err)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

// CreateVpcDnsRecord create vpc dns domain record
func (me *VpcDnsService) CreateVpcDnsRecord(ctx context.Context, request *vpcdns.CreateVpcDnsRecordRequest) (
	recordId int64, errRet error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().CreateVpcDnsRecord(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	recordId = *response.Response.Data.RecordId
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

// DescribeVpcDnsRecordByFilters describe vpc dns record by id
func (me *VpcDnsService) DescribeVpcDnsRecordByFilters(ctx context.Context, domainId int,
	filters []*vpcdns.RecordListFilters) (
	record *vpcdns.RecordDetail, errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDescribeVpcDnsRecordListRequest()
	request.Filters = filters
	request.DomainId = helper.Uint64(uint64(domainId))
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().DescribeVpcDnsRecordList(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	if len(response.Response.Records) < 1 {
		return
	}
	record = response.Response.Records[0]
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// ModifyVpcDnsRecord modify vpc dns record
func (me *VpcDnsService) ModifyVpcDnsRecord(ctx context.Context, request *vpcdns.ModifyVpcDnsRecordRequest) (
	errRet error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().ModifyVpcDnsRecord(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// DeleteVpcDnsRecord delete vpc dns record
func (me *VpcDnsService) DeleteVpcDnsRecord(ctx context.Context, domainId uint64, recordId string) (
	errRet error) {
	logId := getLogId(ctx)
	request := vpcdns.NewDeleteVpcDnsRecordRequest()
	request.DomainId = helper.Uint64(domainId)
	request.RecordIds = helper.String(recordId)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcDnsClient().DeleteVpcDnsRecord(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
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

// DescribeVpcDnsForwardRuleList 分页获取所有转发规则
func (me *VpcDnsService) DescribeVpcDnsForwardRuleList(ctx context.Context, filterMap map[string][]string) (
	forwardRule *vpcdns.VpcDnsForwardRuleDetail, errRet error) {
	logId := getLogId(ctx)
	var (
		offset   = 0
		limit    = 20
		ruleList []*vpcdns.VpcDnsForwardRuleDetail
	)
	for {
		request := vpcdns.NewDescribeVpcDnsForwardRuleRequest()
		request.Filters = []*vpcdns.ForwardRuleFilter{}
		for name, val := range filterMap {
			request.Filters = append(request.Filters, &vpcdns.ForwardRuleFilter{
				Name:   helper.String(name),
				Values: helper.Strings(val),
			})
		}
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
	if len(ruleList) > 0 {
		forwardRule = ruleList[0]
	}
	log.Printf("[DEBUG]%s api[%s] success, total rules [%d]\n",
		logId, "DescribeVpcDnsForwardRule", len(ruleList))
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

// DescribeVpcDnsZoneRecordByFilter describe zone records with pagination
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

package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-tencentcloudenterprise/sdk/common"
	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type NgwafService struct {
	client *TencentCloudClient
}

func (me *NgwafService) DescribeWafCustomRuleById(ctx context.Context, domain, ruleId string) (
	CustomRule *ngwaf.DescribeCustomRulesRspRuleListItem, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeCustomRuleListRequest()
	response := ngwaf.NewDescribeCustomRuleListResponse()
	request.Domain = &domain
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(20)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("RuleId"),
			Values:     common.StringPtrs([]string{ruleId}),
			ExactMatch: helper.Bool(true),
		},
	}

	if errRet != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), errRet.Error())
	}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeCustomRuleList(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.RuleList) < 1 {
		return
	}

	CustomRule = response.Response.RuleList[0]
	return
}

func (me *NgwafService) DeleteWafCustomRuleById(ctx context.Context, domain, ruleId string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteCustomRuleRequest()
	request.Domain = &domain
	request.RuleId = &ruleId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DeleteCustomRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *NgwafService) DescribeWafCustomWhiteRuleById(ctx context.Context, domain, ruleId string) (
	CustomWhiteRule *ngwaf.DescribeCustomRulesRspRuleListItem, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeCustomWhiteRuleRequest()
	response := ngwaf.NewDescribeCustomWhiteRuleResponse()
	request.Domain = &domain
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(20)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("RuleId"),
			Values:     common.StringPtrs([]string{ruleId}),
			ExactMatch: helper.Bool(true),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeCustomWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.RuleList == nil {
			return resource.NonRetryableError(fmt.Errorf("Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.RuleList) < 1 {
		return
	}

	CustomWhiteRule = response.Response.RuleList[0]
	return
}

func (me *NgwafService) DeleteWafCustomWhiteRuleById(ctx context.Context, domain, ruleId string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteCustomWhiteRuleRequest()
	request.Domain = &domain
	tmpRuleId, _ := strconv.ParseUint(ruleId, 10, 64)
	request.RuleId = &tmpRuleId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DeleteCustomWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *NgwafService) DescribeWafCiphersByFilter(ctx context.Context) (
	ciphers []*ngwaf.TLSCiphers, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeCiphersDetailRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeCiphersDetail(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.Ciphers) < 1 {
		return
	}

	ciphers = response.Response.Ciphers
	return
}

func (me *NgwafService) DescribeWafTlsVersionsByFilter(ctx context.Context) (
	tlsVersions []*ngwaf.TLSVersion, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeTlsVersionRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeTlsVersion(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.TLS) < 1 {
		return
	}

	tlsVersions = response.Response.TLS
	return
}

func (me *NgwafService) DescribeDomainsById(ctx context.Context, instanceID, domain string) (
	domainInfo *ngwaf.DomainInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeDomainsRequest()
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(20)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("InstanceId"),
			Values:     common.StringPtrs([]string{instanceID}),
			ExactMatch: helper.Bool(true),
		},
		{
			Name:       helper.String("Domain"),
			Values:     common.StringPtrs([]string{domain}),
			ExactMatch: helper.Bool(true),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeDomains(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Domains) < 1 {
		return
	}

	domainInfo = response.Response.Domains[0]
	return
}

func (me *NgwafService) DescribeWafClbDomainById(ctx context.Context, instanceID, domain, domainId string) (
	clbDomainInfo *ngwaf.ClbDomainsInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeDomainDetailsClbRequest()
	request.InstanceId = &instanceID
	request.Domain = &domain
	request.DomainId = &domainId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeDomainDetailsClb(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response.DomainsClbPartInfo == nil {
		return
	}

	clbDomainInfo = response.Response.DomainsClbPartInfo
	return
}

func (me *NgwafService) DeleteWafClbDomainById(ctx context.Context, instanceID, domain, domainId string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteHostRequest()
	request.HostsDel = []*ngwaf.HostDel{
		{
			Domain:     helper.String(domain),
			InstanceID: helper.String(instanceID),
			DomainId:   helper.String(domainId),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteHost(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafSaasDomainById(ctx context.Context, instanceID, domain, domainId string) (
	saasDomain *ngwaf.DomainsPartInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeDomainDetailsSaasRequest()
	request.InstanceId = &instanceID
	request.Domain = &domain
	request.DomainId = &domainId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeDomainDetailsSaas(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response.DomainsPartInfo == nil {
		return
	}

	saasDomain = response.Response.DomainsPartInfo
	return
}

func (me *NgwafService) DeleteWafSaasDomainById(ctx context.Context, instanceID, domain string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteSpartaProtectionRequest()
	request.InstanceID = helper.String(instanceID)
	request.Domains = common.StringPtrs([]string{domain})

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteSpartaProtection(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafDomainsByFilter(ctx context.Context, instanceID, domain string) (
	domains []*ngwaf.DomainInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeDomainsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	tmpFilter := []*ngwaf.FiltersItemNew{}
	if instanceID != "" {
		tmpFilter = append(tmpFilter, &ngwaf.FiltersItemNew{
			Name:       helper.String("InstanceId"),
			Values:     common.StringPtrs([]string{instanceID}),
			ExactMatch: helper.Bool(true),
		})
	}

	if domain != "" {
		tmpFilter = append(tmpFilter, &ngwaf.FiltersItemNew{
			Name:       helper.String("Domain"),
			Values:     common.StringPtrs([]string{domain}),
			ExactMatch: helper.Bool(true),
		})
	}

	request.Filters = tmpFilter

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.apiV3Conn.UseNgwafClient().DescribeDomains(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Domains) < 1 {
			break
		}

		domains = append(domains, response.Response.Domains...)
		if len(response.Response.Domains) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *NgwafService) DescribeWafFindDomainsByFilter(ctx context.Context, param map[string]interface{}) (
	findDomains []*ngwaf.FindAllDomainDetail, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeFindDomainListRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Key" {
			request.Key = v.(*string)
		}

		if k == "IsWafDomain" {
			request.IsWafDomain = v.(*string)
		}

		if k == "By" {
			request.By = v.(*string)
		}

		if k == "Order" {
			request.Order = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 1
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.apiV3Conn.UseNgwafClient().DescribeFindDomainList(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}

		findDomains = append(findDomains, response.Response.List...)
		if len(response.Response.List) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *NgwafService) DescribeWafPortsByFilter(ctx context.Context, param map[string]interface{}) (
	ports *ngwaf.DescribePortsResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribePortsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Edition" {
			request.Edition = v.(*string)
		}

		if k == "InstanceID" {
			request.InstanceID = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribePorts(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	ports = response
	return
}

func (me *NgwafService) DescribeWafUserDomainsByFilter(ctx context.Context) (
	userDomains []*ngwaf.UserDomainInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeUserDomainInfoRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeUserDomainInfo(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.UsersInfo) < 1 {
		return
	}

	userDomains = response.Response.UsersInfo
	return
}

func (me *NgwafService) DescribeWafInstances(ctx context.Context) (
	instances []*ngwaf.InstanceInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeInstancesRequest()
	request.Offset = helper.Uint64(1)
	request.Limit = helper.Uint64(100)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		offset uint64 = 1
		limit  uint64 = 100
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.apiV3Conn.UseNgwafClient().DescribeInstances(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			if result == nil || result.Response == nil {
				return resource.NonRetryableError(fmt.Errorf("DescribeInstances response is nil"))
			}

			if len(result.Response.Instances) > 0 {
				instances = append(instances, result.Response.Instances...)
			}

			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if len(instances) < int(offset+limit-1) {
			break
		}

		offset += limit
	}

	return
}

func (me *NgwafService) DescribeWafInstanceById(ctx context.Context, instanceId string) (
	instance *ngwaf.InstanceInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeInstancesRequest()
	response := ngwaf.NewDescribeInstancesResponse()
	request.Offset = helper.Uint64(1)
	request.Limit = helper.Uint64(20)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("InstanceId"),
			Values:     common.StringPtrs([]string{instanceId}),
			ExactMatch: helper.Bool(false),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeInstances(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.Instances) < 1 {
		return
	}

	instance = response.Response.Instances[0]
	return
}

func (me *NgwafService) DescribeWafInstanceWaitStatusById(ctx context.Context, instanceId string) error {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeInstancesRequest()
	request.Offset = helper.Uint64(1)
	request.Limit = helper.Uint64(20)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("InstanceId"),
			Values:     common.StringPtrs([]string{instanceId}),
			ExactMatch: helper.Bool(false),
		},
	}

	err := resource.Retry(readRetryTimeout*10, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeInstances(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("DescribeInstances response is nil"))
		}

		// Check if instance exists by Total count
		if result.Response.Total != nil && *result.Response.Total > 0 {
			// If Total > 0 but Instances array is empty, data not fully synced yet
			if len(result.Response.Instances) < 1 {
				return resource.RetryableError(fmt.Errorf("waf instance %s found (Total=%d) but details not ready yet", instanceId, *result.Response.Total))
			}

			// Instance found with details, check status
			instance := result.Response.Instances[0]
			if instance.Status != nil && *instance.Status == 0 {
				return nil
			}

			status := uint64(0)
			if instance.Status != nil {
				status = *instance.Status
			}

			return resource.RetryableError(fmt.Errorf("waf instance still running, status is %d", status))
		}

		// Total is 0 or nil, instance not found yet
		return resource.RetryableError(fmt.Errorf("waf instance %s not found yet", instanceId))
	})

	if err != nil {
		return err
	}

	return nil
}

func (me *NgwafService) DescribeNgwafAttackLogHistogramByFilter(ctx context.Context, param map[string]interface{}) (
	AttackLogHistogram *ngwaf.GetAttackHistogramResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewGetAttackHistogramRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Domain" {
			request.Domain = v.(*string)
		}

		if k == "StartTime" {
			request.StartTime = v.(*string)
		}

		if k == "EndTime" {
			request.EndTime = v.(*string)
		}

		if k == "QueryString" {
			request.QueryString = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().GetAttackHistogram(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	AttackLogHistogram = response
	return
}

func (me *NgwafService) DescribeWafAttackLogListByFilter(ctx context.Context, param map[string]interface{}) (
	AttackLogList []*ngwaf.AttackLogInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewSearchAttackLogRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Domain" {
			request.Domain = v.(*string)
		}

		if k == "StartTime" {
			request.StartTime = v.(*string)
		}

		if k == "EndTime" {
			request.EndTime = v.(*string)
		}

		if k == "Count" {
			request.Count = v.(*int64)
		}

		if k == "QueryString" {
			request.QueryString = v.(*string)
		}

		if k == "Sort" {
			request.Sort = v.(*string)
		}

		if k == "Page" {
			request.Page = v.(*int64)
		}
	}

	request.Context = helper.String("")

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().SearchAttackLog(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	AttackLogList = response.Response.Data
	return
}

func (me *NgwafService) DescribeNgwafAttackOverviewByFilter(ctx context.Context, param map[string]interface{}) (
	AttackOverview *ngwaf.DescribeAttackOverviewResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeAttackOverviewRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "FromTime" {
			request.FromTime = v.(*string)
		}

		if k == "ToTime" {
			request.ToTime = v.(*string)
		}

		if k == "Appid" {
			request.Appid = v.(*uint64)
		}

		if k == "Domain" {
			request.Domain = v.(*string)
		}

		if k == "Edition" {
			request.Edition = v.(*string)
		}

		if k == "InstanceID" {
			request.InstanceID = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeAttackOverview(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	AttackOverview = response
	return
}

func (me *NgwafService) DescribeWafAttackTotalCountByFilter(ctx context.Context, param map[string]interface{}) (
	AttackTotalCount *ngwaf.GetAttackTotalCountResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewGetAttackTotalCountRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "StartTime" {
			request.StartTime = v.(*string)
		}

		if k == "EndTime" {
			request.EndTime = v.(*string)
		}

		if k == "Domain" {
			request.Domain = v.(*string)
		}

		if k == "QueryString" {
			request.QueryString = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().GetAttackTotalCount(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	AttackTotalCount = response
	return
}

func (me *NgwafService) DescribeWafPeakPointsByFilter(ctx context.Context, param map[string]interface{}) (
	PeakPoints []*ngwaf.PeakPointsItem, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribePeakPointsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "FromTime" {
			request.FromTime = v.(*string)
		}

		if k == "ToTime" {
			request.ToTime = v.(*string)
		}

		if k == "Domain" {
			request.Domain = v.(*string)
		}

		if k == "Edition" {
			request.Edition = v.(*string)
		}

		if k == "InstanceID" {
			request.InstanceID = v.(*string)
		}

		if k == "MetricName" {
			request.MetricName = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribePeakPoints(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	PeakPoints = response.Response.Points
	return
}

func (me *NgwafService) DescribeWafAntiFakeById(ctx context.Context, id, domain string) (
	antiFake *ngwaf.CacheUrlItems, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeAntiFakeRulesRequest()
	request.Domain = &domain
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(10)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("RuleId"),
			Values:     common.StringPtrs([]string{id}),
			ExactMatch: helper.Bool(true),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeAntiFakeRules(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Data) < 1 {
		return
	}

	antiFake = response.Response.Data[0]
	return
}

func (me *NgwafService) DeleteWafAntiFakeById(ctx context.Context, id, domain string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteAntiFakeUrlRequest()
	idInt, _ := strconv.ParseUint(id, 10, 64)
	request.Id = helper.Uint64(idInt)
	request.Domain = helper.String(domain)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteAntiFakeUrl(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafAntiInfoLeakById(ctx context.Context, ruleId, domain string) (
	antiInfoLeak *ngwaf.DescribeAntiLeakageItem, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeAntiInfoLeakageRulesRequest()
	request.Domain = &domain
	request.Limit = helper.Uint64(10)
	request.Offset = helper.Uint64(0)
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("RuleId"),
			Values:     common.StringPtrs([]string{ruleId}),
			ExactMatch: helper.Bool(true),
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeAntiInfoLeakageRules(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.RuleList) < 1 {
		return
	}

	antiInfoLeak = response.Response.RuleList[0]
	return
}

func (me *NgwafService) DeleteWafAntiInfoLeakById(ctx context.Context, ruleId, domain string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteAntiInfoLeakRuleRequest()
	ruleIdInt, _ := strconv.ParseUint(ruleId, 10, 64)
	request.Domain = &domain
	request.RuleId = &ruleIdInt

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteAntiInfoLeakRule(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafInstanceQpsLimitByFilter(ctx context.Context, param map[string]interface{}) (
	instanceQpsLimit *ngwaf.QpsData, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewGetInstanceQpsLimitRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "InstanceId" {
			request.InstanceId = v.(*string)
		}

		if k == "Type" {
			request.Type = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().GetInstanceQpsLimit(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	instanceQpsLimit = response.Response.QpsData
	return
}

func (me *NgwafService) DescribeWafAutoDenyRulesById(ctx context.Context, domain string) (
	autoDenyRules *ngwaf.DescribeWafAutoDenyRulesResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeWafAutoDenyRulesRequest()
	request.Domain = &domain

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeWafAutoDenyRules(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	autoDenyRules = response
	return
}

func (me *NgwafService) DescribeWafModuleStatusById(ctx context.Context, domain string) (
	moduleStatus *ngwaf.DescribeModuleStatusResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeModuleStatusRequest()
	request.Domain = &domain

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeModuleStatus(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	moduleStatus = response
	return
}

func (me *NgwafService) DescribeSpartaProtectionInfoById(ctx context.Context, domain, edition string) (
	protectionInfo *ngwaf.DescribeSpartaProtectionInfoResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeSpartaProtectionInfoRequest()
	request.Domain = &domain
	request.Edition = &edition

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeSpartaProtectionInfo(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	protectionInfo = response
	return
}

func (me *NgwafService) DescribeWafWebShellById(ctx context.Context, domain string) (
	webShell *ngwaf.DescribeWebshellStatusResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeWebshellStatusRequest()
	request.Domain = &domain

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeWebshellStatus(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	webShell = response
	return
}

func (me *NgwafService) DescribeWafUserClbRegionsByFilter(ctx context.Context) (
	userClbRegions *ngwaf.DescribeUserClbWafRegionsResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = ngwaf.NewDescribeUserClbWafRegionsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeUserClbWafRegions(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	userClbRegions = response
	return
}

func (me *NgwafService) DescribeWafCcById(ctx context.Context, domain, ruleId string) (cc *ngwaf.CCRuleItems, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeCCRuleListRequest()
	response := ngwaf.NewDescribeCCRuleListResponse()
	request.Domain = &domain
	request.Filters = []*ngwaf.FiltersItemNew{
		{
			Name:       helper.String("RuleId"),
			Values:     common.StringPtrs([]string{ruleId}),
			ExactMatch: helper.Bool(true),
		},
	}
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(10)
	request.By = helper.String("ts_version")

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeCCRuleList(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.Data == nil {
			return resource.NonRetryableError(fmt.Errorf("Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.Data.Res) != 1 {
		return
	}

	cc = response.Response.Data.Res[0]
	return
}

func (me *NgwafService) DeleteWafCcById(ctx context.Context, domain, ruleId, name string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteCCRuleRequest()
	request.Domain = helper.String(domain)
	request.Name = helper.String(name)
	ruleIdInt, _ := strconv.ParseInt(ruleId, 10, 64)
	request.RuleId = common.Int64Ptr(ruleIdInt)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DeleteCCRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *NgwafService) DescribeWafCcAutoStatusById(ctx context.Context, domain string) (
	CcAutoStatus *ngwaf.DescribeCCAutoStatusResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeCCAutoStatusRequest()
	request.Domain = &domain

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeCCAutoStatus(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	CcAutoStatus = response
	return
}

func (me *NgwafService) DeleteWafCcAutoStatusById(ctx context.Context, domain, edition string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewUpsertCCAutoStatusRequest()
	request.Domain = &domain
	request.Edition = &edition
	request.Value = helper.Int64(0)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().UpsertCCAutoStatus(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafCcSessionById(ctx context.Context, domain, edition, sessionID string) (
	ccSession *ngwaf.SessionItem, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeSessionRequest()
	request.Domain = &domain
	request.Edition = &edition
	sessionIDInt, _ := strconv.ParseInt(sessionID, 10, 64)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeSession(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.Data.Res) < 1 {
		return
	}

	for _, item := range response.Response.Data.Res {
		if *item.SessionId == sessionIDInt {
			ccSession = item
			break
		}
	}

	return
}

func (me *NgwafService) DeleteWafCcSessionById(ctx context.Context, domain, edition, sessionID string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteSessionRequest()
	request.Domain = &domain
	request.Edition = &edition
	sessionIDInt, _ := strconv.ParseInt(sessionID, 10, 64)
	request.SessionID = &sessionIDInt

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteSession(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafIpAccessControlById(ctx context.Context, domain string) (
	ipAccessControlList []*ngwaf.IpAccessControlItem, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeIpAccessControlRequest()
	request.Domain = &domain
	request.Count = helper.Uint64(1)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)

	for {
		request.OffSet = &offset
		request.Limit = &limit
		response, err := me.client.apiV3Conn.UseNgwafClient().DescribeIpAccessControl(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Data.Res) < 1 {
			break
		}

		ipAccessControlList = append(ipAccessControlList, response.Response.Data.Res...)
		if len(response.Response.Data.Res) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *NgwafService) DeleteWafIpAccessControlByDiff(ctx context.Context, domain string, ids []string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteIpAccessControlRequest()
	request.Domain = &domain
	request.IsId = helper.Bool(true)
	request.Items = common.StringPtrs(ids)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteIpAccessControl(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DeleteWafIpAccessControlById(ctx context.Context, domain string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteIpAccessControlRequest()
	request.Domain = &domain
	request.Items = common.StringPtrs([]string{""})
	request.DeleteAll = helper.Bool(true)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DeleteIpAccessControl(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *NgwafService) DescribeWafIpAccessControlV2ById(ctx context.Context, domain string, ruleId string) (
	ret *ngwaf.DescribeIpAccessControlResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeIpAccessControlRequest()
	request.Domain = helper.String(domain)
	request.Count = helper.Uint64(1)
	request.RuleId = func() *uint64 { v, _ := strconv.ParseUint(ruleId, 10, 64); return &v }()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.apiV3Conn.UseNgwafClient().DescribeIpAccessControl(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	ret = response
	return
}

func (me *NgwafService) DescribeWafLogPostClsFlowById(ctx context.Context, logType int64) (
	ret *ngwaf.DescribePostCLSFlowsResponse, errRet error) {
	logId := getLogId(ctx)
	request := ngwaf.NewDescribePostCLSFlowsRequest()
	response := ngwaf.NewDescribePostCLSFlowsResponse()
	request.LogType = &logType

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribePostCLSFlows(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response
	return
}

func (me *NgwafService) DescribeWafLogPostCkafkaFlowById(ctx context.Context, logType int64) (
	ret *ngwaf.DescribePostCKafkaFlowsResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribePostCKafkaFlowsRequest()
	response := ngwaf.NewDescribePostCKafkaFlowsResponse()
	request.LogType = &logType

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribePostCKafkaFlows(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response
	return
}

func (me *NgwafService) DescribeWafDomainPostActionById(ctx context.Context, domain string) (
	domains []*ngwaf.DomainInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeDomainsRequest()
	response := ngwaf.NewDescribeDomainsResponse()
	tmpFilter := []*ngwaf.FiltersItemNew{}
	if domain != "" {
		tmpFilter = append(tmpFilter, &ngwaf.FiltersItemNew{
			Name:       helper.String("Domain"),
			Values:     common.StringPtrs([]string{domain}),
			ExactMatch: helper.Bool(true),
		})
	}

	request.Filters = tmpFilter

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.apiV3Conn.UseNgwafClient().DescribeDomains(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if response == nil || len(response.Response.Domains) < 1 {
			break
		}

		domains = append(domains, response.Response.Domains...)
		if len(response.Response.Domains) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *NgwafService) DescribeWafBotSceneStatusConfigById(ctx context.Context, domain string, sceneId string) (
	ret *ngwaf.BotSceneInfo, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeBotSceneListRequest()
	response := ngwaf.NewDescribeBotSceneListResponse()
	request.Domain = &domain
	request.SceneId = &sceneId
	// wait waf sdk update
	// request.BusinessType = common.StringPtrs([]string{"all"})
	request.BusinessType = helper.Strings([]string{"login", "seckill", "crawl", "scan", "key-protect",
		"click-farming", "junk-mail", "social-media", "auto-download", "custom"})

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		BotSceneList []*ngwaf.BotSceneInfo
		offset       int64 = 0
		limit        int64 = 20
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.apiV3Conn.UseNgwafClient().DescribeBotSceneList(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if response == nil || len(response.Response.BotSceneList) < 1 {
			break
		}

		BotSceneList = append(BotSceneList, response.Response.BotSceneList...)
		if len(response.Response.BotSceneList) < int(limit) {
			break
		}

		offset += limit
	}

	for _, item := range BotSceneList {
		if *item.SceneId == sceneId {
			ret = item
			return
		}
	}

	return
}

func (me *NgwafService) DescribeWafBotStatusConfigById(ctx context.Context, domain string) (
	ret *ngwaf.DescribeBotSceneOverviewResponse, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeBotSceneOverviewRequest()
	response := ngwaf.NewDescribeBotSceneOverviewResponse()
	request.Domain = helper.String(domain)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DescribeBotSceneOverview(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response
	return
}

func (me *NgwafService) DescribeWafBotSceneUCBRuleById(ctx context.Context, domain, sceneId, ruleId string) (
	ret *ngwaf.InOutputBotUCBRule, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeBotSceneUCBRuleRequest()
	response := ngwaf.NewDescribeBotSceneUCBRuleResponse()
	request.Domain = &domain
	request.SceneId = &sceneId
	request.RuleId = &ruleId
	request.Sort = helper.String("timestamp:-1")

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		BotSceneUCBRuleList []*ngwaf.InOutputBotUCBRule
		skip                uint64 = 0
		limit               uint64 = 20
	)

	for {
		request.Skip = &skip
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.apiV3Conn.UseNgwafClient().DescribeBotSceneUCBRule(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if response == nil || response.Response == nil || response.Response.Data == nil || len(response.Response.Data.Res) < 1 {
			break
		}

		BotSceneUCBRuleList = append(BotSceneUCBRuleList, response.Response.Data.Res...)
		if len(response.Response.Data.Res) < int(limit) {
			break
		}

		limit += skip
	}

	for _, item := range BotSceneUCBRuleList {
		if *item.SceneId == sceneId {
			ret = item
			return
		}
	}

	return
}

func (me *NgwafService) DeleteWafBotSceneUCBRuleById(ctx context.Context, domain, sceneId, ruleId string) (errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDeleteBotSceneUCBRuleRequest()
	request.Domain = helper.String(domain)
	request.SceneId = helper.String(sceneId)
	request.RuleId = helper.String(ruleId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.apiV3Conn.UseNgwafClient().DeleteBotSceneUCBRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *NgwafService) DescribeWafAttackWhiteRuleById(ctx context.Context, domain string, ruleId uint64) (
	ret *ngwaf.UserWhiteRule, errRet error) {
	logId := getLogId(ctx)

	request := ngwaf.NewDescribeAttackWhiteRuleRequest()
	response := ngwaf.NewDescribeAttackWhiteRuleResponse()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Domain = &domain

	var (
		offset uint64 = 0
		limit  uint64 = 20
		wrList []*ngwaf.UserWhiteRule
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.apiV3Conn.UseNgwafClient().DescribeAttackWhiteRule(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if response == nil || len(response.Response.List) < 1 {
			break
		}

		wrList = append(wrList, response.Response.List...)
		if len(response.Response.List) < int(limit) {
			break
		}

		offset += limit
	}

	for _, item := range wrList {
		if item.WhiteRuleId != nil && *item.WhiteRuleId == ruleId {
			ret = item
			break
		}
	}

	return
}
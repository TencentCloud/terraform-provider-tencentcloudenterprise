package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func NewBhService(client *connectivity.TencentCloudClient) BhService {
	return BhService{client: client}
}

type BhService struct {
	client *connectivity.TencentCloudClient
}

func (me *BhService) DescribeBhAccessWhiteListRuleById(ctx context.Context, ruleId string) (ret *v20191018.AccessWhiteListRule, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeAccessWhiteListRulesRequest()
	response := v20191018.NewDescribeAccessWhiteListRulesResponse()
	request.IdSet = []*uint64{helper.StrToUint64Point(ruleId)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeAccessWhiteListRules(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.AccessWhiteListRuleSet == nil || len(result.Response.AccessWhiteListRuleSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe access white list rules failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.AccessWhiteListRuleSet[0]
	return
}

func (me *BhService) DescribeBhAccessWhiteListConfigById(ctx context.Context) (ret *v20191018.DescribeAccessWhiteListRulesResponse, errRet error) {
	logId := getLogId(ctx)
	request := v20191018.NewDescribeAccessWhiteListRulesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeAccessWhiteListRules(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe access white list rules failed, Response is nil."))
		}

		ret = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *BhService) DescribeBhDeviceById(ctx context.Context, deviceId string) (ret *v20191018.Device, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeDevicesRequest()
	response := v20191018.NewDescribeDevicesResponse()
	request.IdSet = []*uint64{helper.StrToUint64Point(deviceId)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeDevices(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.DeviceSet == nil || len(result.Response.DeviceSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe device failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.DeviceSet[0]
	return
}

func (me *BhService) DescribeBhAssetSyncFlagConfigById(ctx context.Context) (ret *v20191018.AssetSyncFlags, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeAssetSyncFlagRequest()
	response := v20191018.NewDescribeAssetSyncFlagResponse()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeAssetSyncFlag(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.AssetSyncFlags == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe asset sync flag failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.AssetSyncFlags
	return
}

func (me *BhService) DescribeBhResourceById(ctx context.Context, resourceId string) (ret *v20191018.Resource, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeResourcesRequest()
	response := v20191018.NewDescribeResourcesResponse()
	request.ResourceIds = []*string{&resourceId}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeResources(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.ResourceSet == nil || len(result.Response.ResourceSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe resource failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.ResourceSet[0]
	return
}

func (me *BhService) DescribeBhReconnectionSettingConfigById(ctx context.Context) (ret *v20191018.SecuritySetting, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeSecuritySettingRequest()
	response := v20191018.NewDescribeSecuritySettingResponse()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeSecuritySetting(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.SecuritySetting == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe security setting failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.SecuritySetting
	return
}

func (me *BhService) DescribeBhDepartments(ctx context.Context) (ret *v20191018.Departments, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeDepartmentsRequest()
	response := v20191018.NewDescribeDepartmentsResponse()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeDepartments(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.Departments == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe departments failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.Departments
	return
}

func (me *BhService) DescribeBhUserById(ctx context.Context, userId string) (ret *v20191018.User, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeUsersRequest()
	response := v20191018.NewDescribeUsersResponse()
	request.IdSet = []*uint64{helper.StrToUint64Point(userId)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeUsers(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.UserSet == nil || len(result.Response.UserSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe users failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.UserSet[0]
	return
}

func (me *BhService) DescribeBhUserGroupById(ctx context.Context, userGroupId string) (ret *v20191018.Group, errRet error) {
	logId := getLogId(ctx)

	request := v20191018.NewDescribeUserGroupsRequest()
	response := v20191018.NewDescribeUserGroupsResponse()
	request.IdSet = []*uint64{helper.StrToUint64Point(userGroupId)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseBhsaasClient().DescribeUserGroups(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.GroupSet == nil || len(result.Response.GroupSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe user groups failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	ret = response.Response.GroupSet[0]
	return
}

func (me *BhService) DescribeBhDevicesByFilter(ctx context.Context, param map[string]interface{}) (ret []*v20191018.Device, errRet error) {
	var (
		logId    = getLogId(ctx)
		request  = v20191018.NewDescribeDevicesRequest()
		response = v20191018.NewDescribeDevicesResponse()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "IdSet" {
			request.IdSet = v.([]*uint64)
		}
		if k == "Name" {
			request.Name = v.(*string)
		}
		if k == "Ip" {
			request.Ip = v.(*string)
		}
		if k == "ApCodeSet" {
			request.ApCodeSet = v.([]*string)
		}
		if k == "Kind" {
			request.Kind = v.(*uint64)
		}
		if k == "AuthorizedUserIdSet" {
			request.AuthorizedUserIdSet = v.([]*uint64)
		}
		if k == "ResourceIdSet" {
			request.ResourceIdSet = v.([]*string)
		}
		if k == "KindSet" {
			request.KindSet = v.([]*uint64)
		}
		if k == "ManagedAccount" {
			request.ManagedAccount = v.(*string)
		}
		if k == "DepartmentId" {
			request.DepartmentId = v.(*string)
		}
		if k == "TagFilters" {
			request.TagFilters = v.([]*v20191018.TagFilter)
		}
		if k == "Filters" {
			request.Filters = v.([]*v20191018.Filter)
		}
	}

	var (
		offset uint64 = 0
		limit  uint64 = 100
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			result, e := me.client.UseBhsaasClient().DescribeDevices(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			if result == nil || result.Response == nil || result.Response.DeviceSet == nil {
				return resource.NonRetryableError(fmt.Errorf("Describe devices failed, Response is nil."))
			}

			response = result
			return nil
		})

		if err != nil {
			errRet = err
			return
		}

		if len(response.Response.DeviceSet) < 1 {
			break
		}

		ret = append(ret, response.Response.DeviceSet...)
		if len(response.Response.DeviceSet) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

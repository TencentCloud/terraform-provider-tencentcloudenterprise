package tencentcloud

import (
	"context"
	"errors"
	"log"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type TbaseService struct {
	client *connectivity.TencentCloudClient
}

// ModifyInstanceName modify instance name
func (me *TbaseService) ModifyInstanceName(ctx context.Context, instanceId, instanceName *string) (errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tbase.NewModifyInstanceNameRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "ModifyInstanceName", request.ToJsonString(), errRet.Error())
		}
	}()

	request.InstanceId = instanceId
	request.InstanceName = instanceName

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().ModifyInstanceName(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

// DescribePGInstances
func (me *TbaseService) DescribePGInstances(ctx context.Context, request *tbase.DescribePGInstancesRequest) (
	instances []*tbase.PGInstanceSet, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "DescribePGInstances", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().DescribePGInstances(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}

	if response == nil || response.Response == nil || len(response.Response.Items) < 1 {
		return
	}
	instances = append(instances, response.Response.Items...)

	return
}

// CreatePGInstanceHour create pg instance by hour
func (me *TbaseService) CreatePGInstanceHour(ctx context.Context, request *tbase.CreatePGInstanceHourRequest) (
	instanceId string, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "CreatePGInstanceHour", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().CreatePGInstanceHour(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	if response == nil || response.Response == nil || len(response.Response.InstanceIds) == 0 || response.Response.InstanceIds[0] == nil || *response.Response.InstanceIds[0] == "" {
		errRet = errors.New("create instance fail: empty instance id")
		return
	}
	instanceId = *response.Response.InstanceIds[0]

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// ResetAccountPassword reset account password
func (me *TbaseService) ResetAccountPassword(ctx context.Context, request *tbase.ResetAccountPasswordRequest) (
	errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "ResetAccountPassword", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().ResetAccountPassword(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// IsolatePGInstanceHour isolate pg instance by hour
func (me *TbaseService) IsolatePGInstanceHour(ctx context.Context, request *tbase.IsolatePGInstanceHourRequest) (
	errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "IsolatePGInstanceHour", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().IsolatePGInstanceHour(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// ReleaseIsolatedPGInstance permanently releases an isolated PG instance.
func (me *TbaseService) ReleaseIsolatedPGInstance(ctx context.Context, instanceId string) (errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tbase.NewReleaseIsolatedPGInstanceRequest()
	)
	request.InstanceId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "ReleaseIsolatedPGInstance", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().ReleaseIsolatedPGInstance(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// ModifyDBInstanceSecurityGroups modify db instance security groups
func (me *TbaseService) ModifyDBInstanceSecurityGroups(ctx context.Context, instanceId *string,
	securityGroupList []*string) (errRet error) {
	var (
		logId = getLogId(ctx)
	)
	request := tbase.NewModifyDBInstanceSecurityGroupsRequest()
	request.Product = helper.String(TBASE_PRODUCT)
	request.InstanceId = instanceId
	request.SecurityGroupIds = securityGroupList
	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, reason[%s]",
				logId, "ModifyDBInstanceSecurityGroups", errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().ModifyDBInstanceSecurityGroups(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response == nil {
		errRet = errors.New("response is nil")
		return
	}

	return
}

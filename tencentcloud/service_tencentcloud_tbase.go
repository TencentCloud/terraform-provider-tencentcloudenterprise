package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	//TDCPG_CLUSTER_FILTER_ID         = "ClusterId"
	//TDCPG_CLUSTER_FILTER_NAME       = "ClusterName"
	//TDCPG_CLUSTER_FILTER_PROJECT_ID = "ProjectId"
	//TDCPG_CLUSTER_FILTER_STATUS     = "Status"
	//TDCPG_CLUSTER_FILTER_PAY_MODE   = "PayMode"
	//
	TBASE_INSTANCE_FILTER_ID   = "InstanceId"
	TBASE_INSTANCE_FILTER_NAME = "InstanceName"
	//TDCPG_INSTANCE_FILTER_ENDPOINT_ID = "EndpointId"
	TBASE_INSTANCE_FILTER_STATUS = "Status"
	TBASE_INSTANCE_FILTER_TYPE   = "InstanceType"
)

type TbaseService struct {
	client *connectivity.TencentCloudClient
}

func (me *TbaseService) TerminateInstanceById(ctx context.Context, instanceId *string) (errRet error) {
	logId := getLogId(ctx)

	request := tbase.NewTerminateInstancesRequest()
	request.InstanceIds = []*string{instanceId}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "isolate tdcpg instance object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTbaseClient().TerminateInstances(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TbaseService) ReleaseIsolatedInstanceById(ctx context.Context, instanceId *string) (errRet error) {
	logId := getLogId(ctx)

	request := tbase.NewReleaseIsolatedInstanceRequest()
	request.InstanceId = instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "isolate tdcpg instance object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTbaseClient().ReleaseIsolatedInstance(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TbaseService) DescribeTbaseInstance(ctx context.Context,
	instanceId *string) (instances []*tbase.Instance, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, reason[%s]\n", logId, "DescribeTdcpgInstance", errRet.Error())
		}
	}()

	paramMap := make(map[string]interface{})
	paramMap["instance_ids"] = []*string{instanceId}

	result, err := me.DescribeTbaseInstancesByFilter(ctx, paramMap)
	if err != nil {
		errRet = err
		return
	}
	instances = result

	return
}

func (me *TbaseService) DeleteTbaseInstanceById(ctx context.Context, instanceId *string) (errRet error) {
	logId := getLogId(ctx)
	var status string
	if err := me.TerminateInstanceById(ctx, instanceId); err != nil {
		return err
	}

	// polling the instance's status to isolated
	err := resource.Retry(3*readRetryTimeout, func() *resource.RetryError {
		result, err := me.DescribeTbaseInstance(ctx, instanceId)
		if err != nil {
			return retryError(err)
		}
		if result != nil {
			status = *result[0].Status
			if status == "isolated" || status == "deleted" {
				return nil
			}

			if status == "isolating" {
				return resource.RetryableError(fmt.Errorf("instance status still on isolating, retry..."))
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITICAL]%s query tdcpg instance failed, reason:%+v", logId, err)
		return err
	}
	if status == "deleted" {
		// do not need delete
		return nil
	}

	err = me.ReleaseIsolatedInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	// wait the instance to be deleted
	err = resource.Retry(3*readRetryTimeout, func() *resource.RetryError {
		result, err := me.DescribeTbaseInstance(ctx, instanceId)
		if err != nil {
			return retryError(err)
		}
		if result != nil {
			status = *result[0].Status
			if status == "deleted" {
				return nil
			}
		}
		return nil
	})
	if err != nil {
		errRet = err
		return err
	}

	return
}

func (me *TbaseService) DescribeTbaseInstancesByFilter(ctx context.Context,
	param map[string]interface{}) (instances []*tbase.Instance, errRet error) {
	var (
		logId         = getLogId(ctx)
		request       = tbase.NewDescribeInstancesRequest()
		index         = 0
		offset  int64 = 0
		limit   int64 = 20
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "instance_ids" {
			var ids []*string
			ids = append(ids, v.([]*string)...)
			request.InstanceIds = ids
			continue
		}

		if k == "instance_name" {
			request.Filters[index] = &tbase.Filter{
				Name:   helper.String(TBASE_INSTANCE_FILTER_NAME),
				Values: []*string{v.(*string)},
			}
			index++
			continue
		}

		if k == "status" {
			request.Status = helper.String(v.(string))
			request.Filters[index] = &tbase.Filter{
				Name:   helper.String(TBASE_INSTANCE_FILTER_STATUS),
				Values: []*string{v.(*string)},
			}
			index++
			continue
		}

		if k == "instance_type" {
			request.Filters[index] = &tbase.Filter{
				Name:   helper.String(TBASE_INSTANCE_FILTER_TYPE),
				Values: []*string{v.(*string)},
			}
			index++
			continue
		}
	}
	ratelimit.Check(request.GetAction())

	request.Offset = helper.Int64(offset)
	request.Limit = helper.Int64(limit)

	response, err := me.client.UseTbaseClient().DescribeInstances(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.Items) < 1 {
		return
	}
	instances = append(instances, response.Response.Items...)

	return
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

// CreateInstanceHour create instance by hour
func (me *TbaseService) CreateInstanceHour(ctx context.Context, request *tbase.CreateInstanceHourRequest) (
	instanceId string, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "ModifyInstanceName", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().CreateInstanceHour(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	if response.Response == nil || len(response.Response.InstanceIds) < 1 {
		errRet = errors.New("create instance fail")
		return
	}

	instanceId = *response.Response.InstanceIds[0]

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

	if response == nil || len(response.Response.Items) < 1 {
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
	if response.Response.InstanceIds == nil {
		errRet = errors.New("create instance fail")
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

// CreatePGReadOnlyVip create pg read only vip
func (me *TbaseService) CreatePGReadOnlyVip(ctx context.Context, request *tbase.CreatePGReadOnlyVipRequest) (
	taskId int64, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "CreatePGReadOnlyVip", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().CreatePGReadOnlyVip(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	if response.Response == nil {
		errRet = errors.New("response is nil")
		return
	}
	taskId = *response.Response.TaskId

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// DescribeInstanceTaskStatus describe instance task status
func (me *TbaseService) DescribeInstanceTaskStatus(ctx context.Context, taskId int64) (
	status string, errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, reason[%s]\n", logId, "DescribeInstanceTaskStatus", errRet.Error())
		}
	}()

	request := tbase.NewDescribeInstanceTaskStatusRequest()
	request.TaskId = &taskId
	response, err := me.client.UseTbaseClient().DescribeInstanceTaskStatus(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response == nil {
		errRet = errors.New("response is nil")
		return
	}
	status = *response.Response.Status
	return
}

func (me *TbaseService) InstanceStateRefreshFunc(taskId int64, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := contextNil
		rets, err := me.DescribeInstanceTaskStatus(ctx, taskId)

		if err != nil {
			return nil, "", err
		}

		for _, state := range failStates {
			if rets == state {
				return state, state, nil
			}
		}

		return rets, rets, nil
	}
}

// DescribePGInstanceDetails describe pg instance details
func (me *TbaseService) DescribePGInstanceDetails(ctx context.Context, instanceId string) (
	response *tbase.DescribePGInstanceDetailsResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tbase.NewDescribePGInstanceDetailsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "DescribePGInstanceDetails", request.ToJsonString(), errRet.Error())
		}
	}()

	request.InstanceId = &instanceId

	ratelimit.Check(request.GetAction())

	res, err := me.client.UseTbaseClient().DescribePGInstanceDetails(request)
	if err != nil {
		log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}

	if res.Response == nil {
		errRet = errors.New("response is nil")
		return
	}

	response = res

	return
}

// ModifyVip modify vip
func (me *TbaseService) ModifyVip(ctx context.Context, request *tbase.ModifyVipRequest) (
	errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "ModifyVip", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().ModifyVip(request)
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

// DeletePGReadOnlyVip delete pg read only vip
func (me *TbaseService) DeletePGReadOnlyVip(ctx context.Context, request *tbase.DeletePGReadOnlyVipRequest) (
	errRet error) {
	var (
		logId = getLogId(ctx)
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, request body [%s], reason[%s]",
				logId, "DeletePGReadOnlyVip", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().DeletePGReadOnlyVip(request)
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

// DescribeDBParameters describe db parameters
func (me *TbaseService) DescribeDBParameters(ctx context.Context, instanceId *string,
	nodeTypes []*string) (response *tbase.DescribeDBParametersResponse, errRet error) {
	var (
		logId = getLogId(ctx)
	)
	request := tbase.NewDescribeDBParametersRequest()
	if nodeTypes != nil {
		request.NodeTypes = nodeTypes
	}
	request.InstanceId = instanceId
	defer func() {
		if errRet != nil {
			log.Printf("[CRITICAL]%s api[%s] fail, reason[%s]",
				logId, "DescribeDBParameters", errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTbaseClient().DescribeDBParameters(request)
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

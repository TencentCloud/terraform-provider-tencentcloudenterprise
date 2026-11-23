package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	mariadb "terraform-provider-tencentcloudenterprise/sdk/mariadb/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type MariadbService struct {
	client *connectivity.TencentCloudClient
}

func mariaDBInstanceNotFound(err error) bool {
	return isExpectError(err, []string{"InvalidParameter.InstanceNotFound"})
}

func (me *MariadbService) InitDbInstance(ctx context.Context, instanceId string, params []*mariadb.DBParamValue) (initRet bool, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = mariadb.NewDescribeDBInstancesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(8*readRetryTimeout, func() *resource.RetryError {
		dbInstance, errResp := me.DescribeMariadbDbInstance(ctx, instanceId)
		if errResp != nil {
			return retryError(errResp, InternalError)
		}
		if dbInstance == nil || dbInstance.Status == nil {
			return resource.RetryableError(fmt.Errorf("waiting for MariaDB instance %s to become visible", instanceId))
		}
		if *dbInstance.Status < 0 {
			return resource.NonRetryableError(fmt.Errorf("db instance init status is %v, operate failed", *dbInstance.Status))
		}
		if *dbInstance.Status == 2 {
			return nil
		}
		if *dbInstance.Status == 3 {
			iniRequest := mariadb.NewInitDBInstancesRequest()
			iniRequest.InstanceIds = []*string{&instanceId}
			iniRequest.Params = params
			initErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := me.client.UseMariadbClient().InitDBInstances(iniRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
						logId, iniRequest.GetAction(), iniRequest.ToJsonString(), result.ToJsonString())
				}
				return nil
			})
			if initErr != nil {
				return resource.NonRetryableError(fmt.Errorf("db instance init error %v, operate failed", initErr))
			}
			return resource.RetryableError(fmt.Errorf("db instance initializing, retry..."))
		}
		return resource.RetryableError(fmt.Errorf("db instance init status is %v, retry...", *dbInstance.Status))
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (me *MariadbService) DescribeMariadbDbInstancesByFilter(ctx context.Context, param map[string]interface{}) (dbInstances []*mariadb.DBInstance, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = mariadb.NewDescribeDBInstancesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "instance_ids" {
			request.InstanceIds = v.([]*string)
		}

		if k == "project_ids" {
			request.ProjectIds = v.([]*int64)
		}

		if k == "search_name" {
			request.SearchName = v.(*string)
		}

		if k == "vpc_id" {
			request.VpcId = v.(*string)
		}

		if k == "subnet_id" {
			request.SubnetId = v.(*string)
		}

		if k == "excluster_type" {
			request.IsFilterExcluster = helper.Bool(true)
			request.ExclusterType = v.(*int64)
		}

		if k == "excluster_ids" {
			request.ExclusterIds = v.([]*string)
		}
	}
	ratelimit.Check(request.GetAction())

	var offset int64 = 0
	var pageSize int64 = 20

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseMariadbClient().DescribeDBInstances(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Instances) < 1 {
			break
		}
		dbInstances = append(dbInstances, response.Response.Instances...)
		if len(response.Response.Instances) < int(pageSize) {
			break
		}
		offset += pageSize
	}
	return
}

func (me *MariadbService) DescribeMariadbDbInstance(ctx context.Context, instanceId string) (dbInstance *mariadb.DBInstance, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = mariadb.NewDescribeDBInstancesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()
	request.InstanceIds = []*string{&instanceId}

	response, err := me.client.UseMariadbClient().DescribeDBInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	if response == nil || len(response.Response.Instances) < 1 {
		return
	}
	dbInstance = response.Response.Instances[0]
	return
}

func (me *MariadbService) DescribeMariadbDbInstanceDetail(ctx context.Context, instanceId string) (dbInstanceDetail *mariadb.DescribeDBInstanceDetailResponseParams, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = mariadb.NewDescribeDBInstanceDetailRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	request.InstanceId = &instanceId

	response, err := me.client.UseMariadbClient().DescribeDBInstanceDetail(request)
	if err != nil {
		if mariaDBInstanceNotFound(err) {
			return nil, nil
		}
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil {
		return
	}

	dbInstanceDetail = response.Response
	return
}

func (me *MariadbService) DeleteMariadbHourDbInstanceById(ctx context.Context, instanceId string) (errRet error) {
	logId := getLogId(ctx)

	request := mariadb.NewDestroyHourDBInstanceRequest()

	request.InstanceId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseMariadbClient().DestroyHourDBInstance(request)
	if err != nil {
		if mariaDBInstanceNotFound(err) {
			return nil
		}
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return resource.Retry(8*readRetryTimeout, func() *resource.RetryError {
		detail, err := me.DescribeMariadbDbInstanceDetail(ctx, instanceId)
		if err != nil {
			return retryError(err)
		}
		if detail != nil {
			return resource.RetryableError(fmt.Errorf("waiting for MariaDB instance %s to be deleted", instanceId))
		}
		return nil
	})
}

func (me *MariadbService) DescribeFlowById(ctx context.Context, flowId int64) (flowParams *mariadb.DescribeFlowResponseParams, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = mariadb.NewDescribeFlowRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	request.FlowId = &flowId
	response, err := me.client.UseMariadbClient().DescribeFlow(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	flowParams = response.Response

	return
}

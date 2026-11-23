package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type CkafkaService struct {
	client *connectivity.TencentCloudClient
}

func (me *CkafkaService) CheckCkafkaInstanceReady(ctx context.Context,
	instanceId string) (has bool, ready bool, errRet error) {
	logId := getLogId(ctx)
	var (
		request  = ckafka.NewDescribeInstancesDetailRequest()
		response = ckafka.NewDescribeInstancesDetailResponse()
		info     *ckafka.InstanceDetail
	)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.InstanceId = &instanceId
	var notFound bool
	if err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, err := me.client.UseCkafkaClient().DescribeInstancesDetail(request)
		if err != nil {
			if isExpectError(err, []string{CkafkaInstanceNotFound}) {
				notFound = true
				return nil
			}
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read ckafka instance failed, reason: %v", logId, err)
		return false, false, err
	}
	if notFound || response == nil || response.Response == nil || response.Response.Result == nil || len(response.Response.Result.InstanceList) < 1 {
		return
	}
	info = response.Response.Result.InstanceList[0]
	if info == nil {
		return
	}
	if info.Status != nil {
		switch *info.Status {
		case -1:
			return true, false, fmt.Errorf("ckafka instance %s creation failed (status -1)", instanceId)
		case 2:
			// Status 2 indicates the instance has been deleted.
			return false, false, nil
		case 1:
			has = true
			ready = true
			return
		}
	}
	has = true
	return
}

func (me *CkafkaService) DescribeCkafkaInstanceById(ctx context.Context,
	instanceId string) (info *ckafka.InstanceDetail, has bool, errRet error) {
	logId := getLogId(ctx)
	var (
		request  = ckafka.NewDescribeInstancesDetailRequest()
		response = ckafka.NewDescribeInstancesDetailResponse()
	)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.InstanceId = &instanceId
	var notFound bool
	if err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, err := me.client.UseCkafkaClient().DescribeInstancesDetail(request)
		if err != nil {
			if isExpectError(err, []string{CkafkaInstanceNotFound}) {
				notFound = true
				return nil
			}
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read ckafka instance failed, reason: %v", logId, err)
		return nil, false, err
	}
	if notFound || response == nil || response.Response == nil || response.Response.Result == nil || len(response.Response.Result.InstanceList) < 1 {
		return
	}
	info = response.Response.Result.InstanceList[0]
	if info == nil || (info.Status != nil && *info.Status == 2) {
		return nil, false, nil
	}
	has = true
	return
}

func (me *CkafkaService) OperateStatusCheck(ctx context.Context, result *ckafka.JgwOperateResponse) (isSucceed bool) {
	logId := getLogId(ctx)
	if result == nil {
		log.Printf("[CRITAL]%s OperateStatusCheck fail, result is nil", logId)
		return false
	}

	if result != nil && *result.ReturnCode == "0" {
		return true
	} else {
		return false
	}
}

func (me *CkafkaService) ModifyPassword(ctx context.Context, instanceId, user, oldPasswd, newPasswd string) (errRet error) {
	logId := getLogId(ctx)
	request := ckafka.NewModifyPasswordRequest()
	request.InstanceId = &instanceId
	request.Name = &user
	request.Password = &oldPasswd
	request.PasswordNew = &newPasswd

	var response *ckafka.ModifyPasswordResponse
	var err error
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		response, err = me.client.UseCkafkaClient().ModifyPassword(request)
		if err != nil {
			return retryError(err)
		}
		return nil
	})

	if err != nil {
		return err
	}
	if response != nil && response.Response != nil && !me.OperateStatusCheck(ctx, response.Response.Result) {
		return fmt.Errorf("[CRITAL]%s api[%s] fail, request body [%s], reason[operate status check failed]", logId, request.GetAction(), request.ToJsonString())
	}
	return nil
}

func (me *CkafkaService) DescribeInstanceById(ctx context.Context, instanceId string) (instanceInfo *ckafka.InstanceAttributesResponse, has bool, errRet error) {
	logId := getLogId(ctx)

	request := ckafka.NewDescribeInstanceAttributesRequest()
	request.InstanceId = &instanceId
	var response *ckafka.DescribeInstanceAttributesResponse
	var err error
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err = me.client.UseCkafkaClient().DescribeInstanceAttributes(request)
		if err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == CkafkaInstanceNotFound || sdkErr.Code == CkafkaFailedOperation {
					return nil
				}
			}
			return retryError(err)
		}
		return nil
	})
	if err != nil {
		errRet = fmt.Errorf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]", logId, request.GetAction(), request.ToJsonString(), err.Error())
		return
	}

	if response != nil && response.Response != nil {
		if instanceInfo = response.Response.Result; instanceInfo != nil {
			has = true
			return
		}
	}

	has = false
	return
}

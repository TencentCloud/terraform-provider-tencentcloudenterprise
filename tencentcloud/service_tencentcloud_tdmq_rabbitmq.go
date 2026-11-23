package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func (me *TdmqService) DescribeTdmqRabbitmqVipInstanceByFilter(ctx context.Context, param map[string]interface{}) (rabbitmqVipInstance []*tdmq.RabbitMQVipInstance, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tdmq.NewDescribeRabbitMQVipInstancesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "filters" {
			request.Filters = v.([]*tdmq.Filter)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseTdmqClient().DescribeRabbitMQVipInstances(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || response.Response.TotalCount == nil {
			return nil, fmt.Errorf("RabbitMQ list response is missing TotalCount")
		}
		if *response.Response.TotalCount == 0 && len(response.Response.Instances) != 0 {
			return nil, fmt.Errorf("RabbitMQ list response contains instances with zero TotalCount")
		}
		if offset < *response.Response.TotalCount && len(response.Response.Instances) == 0 {
			return nil, fmt.Errorf("RabbitMQ list response is missing instances")
		}

		rabbitmqVipInstance = append(rabbitmqVipInstance, response.Response.Instances...)
		if len(response.Response.Instances) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *TdmqService) DescribeTdmqRabbitmqVipInstanceById(ctx context.Context, instanceId string) (rabbitmqVipInstance *tdmq.DescribeRabbitMQVipInstanceResponse, errRet error) {
	logId := getLogId(ctx)

	request := tdmq.NewDescribeRabbitMQVipInstanceRequest()
	request.ClusterId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DescribeRabbitMQVipInstance(request)
	if err != nil {
		if isExpectError(err, []string{"ResourceNotFound.Cluster"}) {
			return nil, nil
		}
		if sdkErr, ok := err.(*errors.CloudSDKError); ok && sdkErr.Code == "FailedOperation" && sdkErr.Message == "not query instance" {
			// Some versions return this generic failure for a missing instance.
			// Confirm absence before allowing Terraform to forget the resource.
			filter := &tdmq.Filter{Name: helper.String("instanceIds"), Values: []*string{&instanceId}}
			instances, listErr := me.DescribeTdmqRabbitmqVipInstanceByFilter(ctx, map[string]interface{}{"filters": []*tdmq.Filter{filter}})
			if listErr != nil {
				return nil, listErr
			}
			if len(instances) == 0 {
				return nil, nil
			}
		}
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	rabbitmqVipInstance = response
	return
}

func (me *TdmqService) DeleteTdmqRabbitmqVipInstanceById(ctx context.Context, instanceId string) (errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tdmq.NewDeleteRabbitMQVipInstanceRequest()
	)

	request.InstanceId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DeleteRabbitMQVipInstance(request)
	if err != nil {
		if isExpectError(err, []string{"ResourceNotFound.Cluster"}) {
			return nil
		}
		if sdkErr, ok := err.(*errors.CloudSDKError); ok && sdkErr.Code == "FailedOperation" && sdkErr.Message == "not query instance" {
			filter := &tdmq.Filter{Name: helper.String("instanceIds"), Values: []*string{&instanceId}}
			instances, listErr := me.DescribeTdmqRabbitmqVipInstanceByFilter(ctx, map[string]interface{}{"filters": []*tdmq.Filter{filter}})
			if listErr == nil && len(instances) == 0 {
				return nil
			}
		}
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

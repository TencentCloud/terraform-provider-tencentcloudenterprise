package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
	"log"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type ClbService struct {
	client *connectivity.TencentCloudClient
}

func (me *ClbService) DescribeLoadBalancerById(ctx context.Context, clbId string) (clbInstance *clb.LoadBalancer, errRet error) {
	logId := getLogId(ctx)
	request := clb.NewDescribeLoadBalancersRequest()
	request.LoadBalancerIds = []*string{&clbId}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseClbClient().DescribeLoadBalancers(request)
	if err != nil {
		errRet = errors.WithStack(err)
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.LoadBalancerSet) < 1 {
		return
	}
	clbInstance = response.Response.LoadBalancerSet[0]
	return
}

func (me *ClbService) DescribeLoadBalancerByFilter(ctx context.Context, params map[string]interface{}) (clbs []*clb.LoadBalancer, errRet error) {
	logId := getLogId(ctx)
	request := clb.NewDescribeLoadBalancersRequest()

	for k, v := range params {
		if k == "clb_id" {
			request.LoadBalancerIds = []*string{helper.String(v.(string))}
		}
		if k == "network_type" {
			request.LoadBalancerType = helper.String(v.(string))
		}
		if k == "clb_name" {
			request.LoadBalancerName = helper.String(v.(string))
		}
		if k == "project_id" {
			projectId := int64(v.(int))
			request.ProjectId = &projectId
		}
		if k == "master_zone" {
			request.MasterZone = helper.String(v.(string))
		}
	}

	offset := int64(0)
	pageSize := int64(CLB_PAGE_LIMIT)
	clbs = make([]*clb.LoadBalancer, 0)
	for {
		request.Offset = &(offset)
		request.Limit = &(pageSize)
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseClbClient().DescribeLoadBalancers(request)
		if err != nil {
			errRet = errors.WithStack(err)
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.LoadBalancerSet) < 1 {
			break
		}

		clbs = append(clbs, response.Response.LoadBalancerSet...)

		if int64(len(response.Response.LoadBalancerSet)) < pageSize {
			break
		}
		offset += pageSize
	}
	return
}

func (me *ClbService) DeleteLoadBalancerById(ctx context.Context, clbId string) error {
	logId := getLogId(ctx)
	request := clb.NewDeleteLoadBalancerRequest()
	request.LoadBalancerIds = []*string{&clbId}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseClbClient().DeleteLoadBalancer(request)
	if err != nil {
		if e, ok := err.(*sdkErrors.CloudSDKError); ok {
			if e.GetCode() == "InvalidParameter.LBIdNotFound" {
				return nil
			}
		}
		return errors.WithStack(err)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	requestId := *response.Response.RequestId
	retryErr := waitForTaskFinish(requestId, me.client.UseClbClient())
	if retryErr != nil {
		return retryErr
	}
	return nil
}

func waitForTaskFinish(requestId string, meta *clb.Client) (err error) {
	taskQueryRequest := clb.NewDescribeTaskStatusRequest()
	taskQueryRequest.TaskId = &requestId
	err = resource.Retry(4*readRetryTimeout, func() *resource.RetryError {
		taskResponse, e := meta.DescribeTaskStatus(taskQueryRequest)
		if e != nil {
			return resource.NonRetryableError(errors.WithStack(e))
		}
		if *taskResponse.Response.Status == int64(CLB_TASK_EXPANDING) {
			return resource.RetryableError(errors.WithStack(fmt.Errorf("CLB task status is %d(expanding), requestId is %s", *taskResponse.Response.Status, *taskResponse.Response.RequestId)))
		} else if *taskResponse.Response.Status == int64(CLB_TASK_FAIL) {
			return resource.NonRetryableError(errors.WithStack(fmt.Errorf("CLB task status is %d(failed), requestId is %s", *taskResponse.Response.Status, *taskResponse.Response.RequestId)))
		}
		return nil
	})
	return
}

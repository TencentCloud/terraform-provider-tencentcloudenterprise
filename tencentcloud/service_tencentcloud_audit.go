package tencentcloud

import (
	"context"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	audit "terraform-provider-tencentcloudenterprise/sdk/cloudaudit/v20190304"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type AuditService struct {
	client *connectivity.TencentCloudClient
}

func (me *AuditService) DescribeAuditTrackById(ctx context.Context, trackId string) (track *audit.DescribeAuditTrackResponse, errRet error) {
	logId := getLogId(ctx)

	request := audit.NewDescribeAuditTrackRequest()
	request.TrackId = helper.StrToInt64Point(trackId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseAuditClient().DescribeAuditTrack(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil {
		return
	}
	track = response
	return
}

func (me *AuditService) DeleteAuditTrackById(ctx context.Context, trackId string) (errRet error) {
	logId := getLogId(ctx)

	request := audit.NewDeleteAuditTrackRequest()

	request.TrackId = helper.StrToInt64Point(trackId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseAuditClient().DeleteAuditTrack(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

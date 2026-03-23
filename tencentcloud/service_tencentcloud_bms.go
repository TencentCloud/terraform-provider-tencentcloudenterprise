package tencentcloud

import (
	"context"
	"fmt"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"log"
)

type BmsService struct {
	client *connectivity.TencentCloudClient
}

func (me *BmsService) DescribeInstanceById(ctx context.Context, instanceId string) (instance *bms.Instance,
	errRet error) {
	logId := getLogId(ctx)
	request := bms.NewDescribeInstancesRequest()
	request.InstanceIds = []*string{&instanceId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().DescribeInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.InstanceSet) < 1 {
		return
	}
	instance = response.Response.InstanceSet[0]
	return
}

func (me *BmsService) DescribeInstanceByName(ctx context.Context, name string) (instance *bms.Instance,
	errRet error) {
	logId := getLogId(ctx)
	request := bms.NewDescribeInstancesRequest()
	request.Offset = helper.IntInt64(0)
	request.Limit = helper.IntInt64(20)
	request.Filters = []*bms.Filter{
		{
			Name: helper.String("instance-name"),
			Values: []*string{
				&name,
			},
		},
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().DescribeInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.InstanceSet) < 1 {
		return
	}
	instance = response.Response.InstanceSet[0]
	return
}

func (me *BmsService) DeleteInstance(ctx context.Context, instanceId string) error {
	logId := getLogId(ctx)
	request := bms.NewTerminateInstancesRequest()
	request.InstanceIds = []*string{&instanceId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().TerminateInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *BmsService) CreatePlacementGroup(ctx context.Context, placementName, placementType string) (errRet error) {
	logId := getLogId(ctx)
	request := bms.NewCreateDisasterRecoverGroupRequest()
	request.Name = &placementName
	request.Type = &placementType

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().CreateDisasterRecoverGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *BmsService) DescribePlacementGroupByFilter(ctx context.Context, id,
	name string) (placementGroups *bms.DisasterRecoverGroup, errRet error) {
	logId := getLogId(ctx)
	request := bms.NewDescribeDisasterRecoverGroupsRequest()
	if id != "" {
		request.GroupIds = []*string{&id}
	}
	if name != "" {
		request.Filters = append(request.Filters, &bms.Filter{
			Name:   helper.String("Name"),
			Values: []*string{&name},
		})
	}

	var offset int64 = 0
	var limit int64 = 200
	request.Offset = &offset
	request.Limit = &limit
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().DescribeDisasterRecoverGroups(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if *response.Response.TotalCount < 1 {
		errRet = fmt.Errorf("placement group not found")
		return
	}
	placementGroups = response.Response.GroupSet[0]
	return
}

func (me *BmsService) UpdateDisasterRecoverGroup(ctx context.Context, placementId, name string) error {
	logId := getLogId(ctx)
	request := bms.NewUpdateDisasterRecoverGroupRequest()
	request.GroupId = &placementId
	request.Name = &name

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().UpdateDisasterRecoverGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *BmsService) DeletePlacementGroup(ctx context.Context, placementId string) error {
	logId := getLogId(ctx)
	request := bms.NewDeleteDisasterRecoverGroupsRequest()
	request.GroupIds = []*string{&placementId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseBmsClient().DeleteDisasterRecoverGroups(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *BmsService) DescribeInstancesByFilter(ctx context.Context, params map[string]interface{}) (
	instances []*bms.Instance, errRet error) {
	logId := getLogId(ctx)
	request := bms.NewDescribeInstancesRequest()
	request.Filters = make([]*bms.Filter, 0, len(params))
	for k, v := range params {
		filter := &bms.Filter{
			Name: helper.String(k),
		}
		switch v := v.(type) {
		case string:
			filter.Values = []*string{helper.String(v)}
		case []*string:
			filter.Values = v
		}
		request.Filters = append(request.Filters, filter)
	}

	offset := 0
	pageSize := 100
	instances = make([]*bms.Instance, 0)
	for {
		request.Offset = helper.IntInt64(offset)
		request.Limit = helper.IntInt64(pageSize)
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseBmsClient().DescribeInstances(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.InstanceSet) < 1 {
			break
		}

		instances = append(instances, response.Response.InstanceSet...)

		if len(response.Response.InstanceSet) < pageSize {
			break
		}
		offset += pageSize
	}
	return
}

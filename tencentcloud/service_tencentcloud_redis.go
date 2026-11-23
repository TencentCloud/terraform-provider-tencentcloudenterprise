package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	"strconv"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"time"
)

type CreateInstanceReqOption func(request *redis.CreateInstancesRequest)

// WithResourcePool ...
func WithResourcePool(poolID string) CreateInstanceReqOption {
	return func(request *redis.CreateInstancesRequest) {
		request.ResourcePoolId = &poolID
	}
}

// WithPlatformProjectID ...
func WithPlatformProjectID(projectID string) CreateInstanceReqOption {
	return func(request *redis.CreateInstancesRequest) {
		request.PlatformProjectId = &projectID
	}
}

type RedisService struct {
	client  *connectivity.TencentCloudClient
	zoneMap map[int64]string
}

type TencentCloudRedisDetail struct {
	RedisId          string
	Name             string
	Zone             string
	ProjectId        int64
	TypeId           int64
	Type             string
	MemSize          int64
	Status           string
	VpcId            string
	SubnetId         string
	Ip               string
	Port             int64
	RedisShardNum    int64
	RedisReplicasNum int64
	CreateTime       string
	Tags             map[string]string
	BillingMode      string
	NodeInfo         []map[string]interface{}
}

func (me *RedisService) fullZoneId() (errRet error) {
	if me.zoneMap == nil {
		me.zoneMap = make(map[int64]string)
	}
	if len(me.zoneMap) != 0 {
		return
	}
	response, err := me.client.UseCvmClient().DescribeZones(cvm.NewDescribeZonesRequest())
	if err != nil {
		return err
	}
	for _, item := range response.Response.ZoneSet {
		if zoneId, err := strconv.ParseInt(*item.ZoneId, 10, 64); err != nil {
			return fmt.Errorf("[sdk]DescribeZones return ZoneId is not illegal,%s", *item.ZoneId)
		} else {
			me.zoneMap[zoneId] = *item.Zone
		}
	}

	return nil
}

func (me *RedisService) getZoneId(name string) (id int64, errRet error) {
	if errRet = me.fullZoneId(); errRet != nil {
		return
	}
	for key, value := range me.zoneMap {
		if value == name {
			id = key
			return
		}
	}
	errRet = fmt.Errorf("this redis zone %s not support yet", name)
	return
}

func (me *RedisService) getZoneName(id int64) (name string, errRet error) {
	if errRet = me.fullZoneId(); errRet != nil {
		return
	}
	name = me.zoneMap[id]
	if name == "" {
		errRet = fmt.Errorf("this redis zoneid %d not support yet", id)
	}
	return
}

func (me *RedisService) DescribeRedisZoneConfig(ctx context.Context) (sellConfigures []*redis.RegionConf, errRet error) {
	logId := getLogId(ctx)
	request := redis.NewDescribeProductInfoRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.BaseRequest, request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DescribeProductInfo(request)
	if err != nil {
		errRet = err
		return
	}
	sellConfigures = response.Response.RegionSet
	return
}

func (me *RedisService) DescribeInstances(ctx context.Context, zoneName, searchKey string,
	projectId, needLimit int64) (instances []TencentCloudRedisDetail, errRet error) {

	logId := getLogId(ctx)

	var zoneId int64 = -1

	if zoneName != "" {
		zoneId, errRet = me.getZoneId(zoneName)
		if errRet != nil {
			return
		}
	}

	listInitSize := map[bool]int64{true: 500, false: needLimit}[needLimit > 500 || needLimit < 1]
	instances = make([]TencentCloudRedisDetail, 0, listInitSize)

	request := redis.NewDescribeInstancesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var (
		limit, offset  uint64 = 2, 0
		leftNumber     int64
		leftNumberInit bool
	)

	request.Limit = &limit
	request.Offset = &offset

needMoreItems:
	if searchKey != "" {
		request.SearchKey = &searchKey
	}
	if projectId >= 0 {
		request.ProjectIds = []*int64{&projectId}
	}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DescribeInstances(request)
	if err != nil {
		errRet = err
		return
	}
	if !leftNumberInit {
		leftNumber = *response.Response.TotalCount
		leftNumberInit = true
	}
	leftNumber = leftNumber - int64(limit)
	offset = offset + limit
	for _, item := range response.Response.InstanceSet {
		if zoneId != -1 && *item.ZoneId != zoneId {
			continue
		}
		var instance TencentCloudRedisDetail
		instance.Type = REDIS_NAMES[*item.Type]
		if REDIS_STATUS[*item.Status] == "" {
			instance.Status = "unknown"
		} else {
			instance.Status = REDIS_STATUS[*item.Status]
		}

		name, err := me.getZoneName(*item.ZoneId)
		if err != nil {
			errRet = err
			return
		}

		instance.Zone = name
		instance.CreateTime = *item.Createtime
		instance.Ip = *item.WanIp
		instance.MemSize = *item.RedisShardSize
		instance.Name = *item.InstanceName
		instance.Port = *item.Port
		instance.ProjectId = *item.ProjectId
		instance.RedisId = *item.InstanceId
		instance.SubnetId = *item.UniqSubnetId
		instance.VpcId = *item.UniqVpcId
		instance.BillingMode = REDIS_CHARGE_TYPE_NAME[*item.BillingMode]

		instance.TypeId = *item.Type
		if item.RedisReplicasNum != nil {
			instance.RedisReplicasNum = *item.RedisReplicasNum
		}
		if item.RedisShardNum != nil {
			instance.RedisShardNum = *item.RedisShardNum
		}
		//
		//if item.NodeSet != nil {
		//	nodeInfos := make([]map[string]interface{}, 0, len(item.NodeSet))
		//	for i := range item.NodeSet {
		//		dMap := make(map[string]interface{})
		//		nodeInfo := item.NodeSet[i]
		//		if *nodeInfo.NodeType == 0 {
		//			dMap["master"] = true
		//		} else {
		//			dMap["master"] = false
		//		}
		//		dMap["id"] = *nodeInfo.NodeId
		//		dMap["zone_id"] = *nodeInfo.ZoneId
		//		nodeInfos = append(nodeInfos, dMap)
		//	}
		//	instance.NodeInfo = nodeInfos
		//}
		instance.Tags = make(map[string]string, len(item.InstanceTags))
		for _, tag := range item.InstanceTags {
			if tag.TagKey == nil {
				return nil, fmt.Errorf("[CRITAL]%s api[%s] redis instance tag key is nil", logId, request.GetAction())
			}
			if tag.TagValue == nil {
				return nil, fmt.Errorf("[CRITAL]%s api[%s] redis instance tag value is nil", logId, request.GetAction())
			}

			instance.Tags[*tag.TagKey] = *tag.TagValue
		}

		instances = append(instances, instance)

		if needLimit > 0 && int64(len(instances)) >= needLimit {
			return
		}
	}
	if leftNumber < 0 {
		return
	} else {
		goto needMoreItems
	}
}

func (me *RedisService) CreateInstances(ctx context.Context,
	zoneName string, typeId int64, password, vpcId, subnetId, redisName string,
	memSize, projectId, port int64,
	securityGroups []string,
	redisShardNum,
	redisReplicasNum int,
	chargeTypeID int64,
	chargePeriod uint64,
	noAuth bool,
	autoRenewFlag int,
	replicasReadonly bool,
	paramsTemplateId string,
	opts ...CreateInstanceReqOption,
) (instanceIds []*string, errRet error) {

	logId := getLogId(ctx)
	request := redis.NewCreateInstancesRequest()
	for _, opt := range opts {
		opt(request)
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	// zone
	var intZoneId int64
	intZoneId, errRet = me.getZoneId(zoneName)
	if errRet != nil {
		return
	}
	request.ZoneId = helper.Int64Uint64(intZoneId)
	request.TypeId = helper.Int64Uint64(typeId)

	// vpc
	if (vpcId == "" && subnetId != "") || (vpcId != "" && subnetId == "") {
		errRet = fmt.Errorf("redis need vpcId and subnetId both set or none")
		return
	}
	if vpcId != "" && subnetId != "" {
		request.VpcId = &vpcId
		request.SubnetId = &subnetId
	} else {
		if len(securityGroups) > 0 {
			errRet = fmt.Errorf("redis need empty security_groups if vpc_id and subnet_id is empty")
			return
		}
	}

	if projectId >= 0 {
		request.ProjectId = &projectId
	}

	var (
		vport           = uint64(port)
		umemSize        = uint64(memSize)
		goodsNum uint64 = 1
	)
	request.VPort = &vport
	request.MemSize = &umemSize
	request.BillingMode = &chargeTypeID
	request.GoodsNum = &goodsNum
	request.Period = &chargePeriod
	if redisShardNum > 0 {
		request.RedisShardNum = helper.IntInt64(redisShardNum)
	}
	if redisReplicasNum > 0 {
		request.RedisReplicasNum = helper.IntInt64(redisReplicasNum)
	}
	if redisName != "" {
		request.InstanceName = &redisName
	}

	request.Password = &password

	if len(securityGroups) > 0 {
		request.SecurityGroupIdList = make([]*string, 0, len(securityGroups))
		for v := range securityGroups {
			request.SecurityGroupIdList = append(request.SecurityGroupIdList, &securityGroups[v])
		}
	}
	// 专有云不支持此参数
	//if len(nodeInfo) > 0 {
	//	request.NodeSet = nodeInfo
	//}

	if noAuth {
		request.NoAuth = &noAuth
	}
	if chargeTypeID == REDIS_CHARGE_TYPE_ID[REDIS_CHARGE_TYPE_PREPAID] {
		request.AutoRenew = helper.IntUint64(autoRenewFlag)
	}
	if replicasReadonly {
		request.ReplicasReadonly = &replicasReadonly
	}

	// 专有云不支持此参数
	//if paramsTemplateId != "" {
	//	request.TemplateId = &paramsTemplateId
	//}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().CreateInstances(request)
	if err != nil {
		errRet = err
		return
	}
	log.Println(response.ToJsonString())
	instanceIds = response.Response.InstanceIds
	return
}

func (me *RedisService) CheckRedisOnlineOk(ctx context.Context, redisId string, retryTimeout time.Duration, waitForCreation bool) (has bool,
	online bool,
	info *redis.InstanceSet,
	errRet error) {

	logId := getLogId(ctx)

	request := redis.NewDescribeInstancesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.InstanceId = &redisId

	// Post https://cdb.tencentcloudapi.com/: always get "Gateway Time-out"
	var response *redis.DescribeInstancesResponse
	err := resource.Retry(retryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseRedisClient().DescribeInstances(request)
		if e != nil {
			log.Printf("[CRITAL]%s CheckRedisOnlineOk fail, reason:%s\n", logId, e.Error())
			return retryError(e)
		}
		response = result

		if len(response.Response.InstanceSet) == 0 {
			has = false
			if waitForCreation {
				return resource.RetryableError(fmt.Errorf("waiting for Redis instance %s to become visible", redisId))
			}
			return nil
		}

		info = response.Response.InstanceSet[0]
		has = true

		if *info.Status == REDIS_STATUS_ONLINE {
			online = true
			return nil
		}

		if *info.Status == REDIS_STATUS_INIT || *info.Status == REDIS_STATUS_PROCESSING {
			online = false
			return resource.RetryableError(fmt.Errorf("istance %s status is %d, retrying", redisId, *info.Status))
		}

		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *RedisService) CheckRedisUpdateOk(ctx context.Context, redisId string) (errRet error) {
	var startUpdate bool
	logId := getLogId(ctx)
	request := redis.NewDescribeInstancesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.InstanceId = &redisId
	errRet = resource.Retry(readRetryTimeout*20, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseRedisClient().DescribeInstances(request)
		if err != nil {
			return retryError(err)
		}
		if len(result.Response.InstanceSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("redis %s not exist", redisId))
		}
		info := result.Response.InstanceSet[0]
		if !startUpdate && *info.Status == REDIS_STATUS_ONLINE {
			return resource.RetryableError(fmt.Errorf("waiting for upgrade start"))
		}
		startUpdate = true
		if *info.Status == REDIS_STATUS_PROCESSING || *info.Status == REDIS_STATUS_INIT {
			return resource.RetryableError(fmt.Errorf("instance %s status is %d", redisId, *info.Status))
		}
		return nil
	})

	return
}

func (me *RedisService) DescribeInstanceDealDetail(ctx context.Context, dealId string) (done bool, redisId string, errRet error) {
	logId := getLogId(ctx)
	request := redis.NewDescribeInstanceDealDetailRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.DealIds = []*string{&dealId}

	// Post https://cdb.tencentcloudapi.com/: always get "Gateway Time-out"
	var response *redis.DescribeInstanceDealDetailResponse
	err := resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseRedisClient().DescribeInstanceDealDetail(request)
		if e != nil {
			log.Printf("[CRITAL]%s DescribeInstanceDealDetail fail, reason:%s\n", logId, e.Error())
			return retryError(e)
		}
		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.DealDetails) != 1 {
		errRet = fmt.Errorf("Redis api DescribeInstanceDealDetail one dealId[%s] return %d deal infos.",
			dealId, len(response.Response.DealDetails))
		return
	}

	dealDetail := response.Response.DealDetails[0]
	status := *dealDetail.Status

	if status == REDIS_ORDER_SUCCESS_DELIVERY {

		if len(dealDetail.InstanceIds) != 1 {
			errRet = fmt.Errorf("redis one dealid give %d redis id", len(dealDetail.InstanceIds))
			return
		}
		redisId = *dealDetail.InstanceIds[0]
		done = true
		return
	}
	if status < REDIS_ORDER_SUCCESS_DELIVERY || status == REDIS_ORDER_PAYMENT {
		return
	}
	errRet = fmt.Errorf("redis instance delivery failure, deal status is %d", status)
	return
}

func (me *RedisService) ModifyInstanceName(ctx context.Context, redisId string, name string) (errRet error) {
	logId := getLogId(ctx)
	request := redis.NewModifyInstanceRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	op := "rename"
	request.InstanceName = &name
	request.Operation = &op
	request.InstanceId = &redisId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().ModifyInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	}
	errRet = err
	return
}

func (me *RedisService) ModifyInstanceProjectId(ctx context.Context, redisId string, projectId int64) (errRet error) {
	logId := getLogId(ctx)
	request := redis.NewModifyInstanceRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	op := "modifyProject"
	request.ProjectId = &projectId
	request.Operation = &op
	request.InstanceId = &redisId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().ModifyInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	}
	errRet = err
	return
}

func (me *RedisService) DescribeInstanceSecurityGroup(ctx context.Context, redisId string) (sg []string, errRet error) {
	logId := getLogId(ctx)
	request := redis.NewDescribeInstanceSecurityGroupRequest()
	request.InstanceIds = []*string{&redisId}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DescribeInstanceSecurityGroup(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	}
	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.InstanceSecurityGroupsDetail) > 0 {
		for _, item := range response.Response.InstanceSecurityGroupsDetail {
			if *item.InstanceId == redisId {
				sg = make([]string, 0, len(item.SecurityGroupDetails))
				for _, v := range item.SecurityGroupDetails {
					sg = append(sg, *v.SecurityGroupId)
				}
				break
			}
		}
	}
	return
}

func (me *RedisService) ModifyDBInstanceSecurityGroups(ctx context.Context, product string, instanceId string, securityGroupIds []*string) (errRet error) {
	logId := getLogId(ctx)
	request := redis.NewModifyDBInstanceSecurityGroupsRequest()
	request.Product = &product
	request.InstanceId = &instanceId
	request.SecurityGroupIds = securityGroupIds

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().ModifyDBInstanceSecurityGroups(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	}
	errRet = err
	return
}

func (me *RedisService) DestroyPostpaidInstance(ctx context.Context, redisId string) (taskId int64, errRet error) {
	logId := getLogId(ctx)
	request := redis.NewDestroyPostpaidInstanceRequest()
	request.InstanceId = &redisId
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DestroyPostpaidInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	} else {
		errRet = err
		return
	}

	taskId = *response.Response.TaskId
	return
}

func (me *RedisService) CleanUpInstance(ctx context.Context, redisId string) (taskId int64, errRet error) {
	logId := getLogId(ctx)
	request := redis.NewCleanUpInstanceRequest()
	request.InstanceId = &redisId
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	// Cleaning up action for prepaid instances needs to retry.
	var response *redis.CleanUpInstanceResponse
	err := resource.Retry(6*writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseRedisClient().CleanUpInstance(request)
		if e != nil {
			log.Printf("[CRITAL]%s CleanUpInstance fail, reason:%s\n", logId, e.Error())
			return retryError(e)
		}
		response = result
		return nil
	})
	if err != nil {
		errRet = err
		return
	}
	taskId = *response.Response.TaskId
	return
}

func (me *RedisService) UpgradeInstance(ctx context.Context, redisId string, newMemSize, redisShardNum, redisReplicasNum int) (dealId string, errRet error) {
	logId := getLogId(ctx)

	request := redis.NewUpgradeInstanceRequest()
	request.InstanceId = &redisId
	if newMemSize > 0 {
		request.MemSize = helper.IntUint64(newMemSize)
	}
	if redisShardNum > 0 {
		request.RedisShardNum = helper.IntUint64(redisShardNum)
	}
	if redisReplicasNum != 0 {
		request.RedisReplicasNum = helper.IntUint64(redisReplicasNum)
	}
	//if len(nodeSet) > 0 {
	//	request.NodeSet = nodeSet
	//}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().UpgradeInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	} else {
		errRet = err
		return
	}

	dealId = *response.Response.DealId
	return
}

func (me *RedisService) DescribeTaskInfo(ctx context.Context, redisId string, taskId int64) (ok bool, errRet error) {
	logId := getLogId(ctx)
	var uintTaskId = uint64(taskId)
	request := redis.NewDescribeTaskInfoRequest()
	request.TaskId = &uintTaskId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DescribeTaskInfo(request)

	if err != nil {
		errRet = err
		return
	}
	if *response.Response.Status == REDIS_TASK_RUNNING || *response.Response.Status == REDIS_TASK_PREPARING {
		return
	}
	if *response.Response.Status == REDIS_TASK_SUCCEED {
		ok = true
		return
	}
	errRet = fmt.Errorf("redis task exe fail, task status is %s", *response.Response.Status)
	return
}

func (me *RedisService) ResetPassword(ctx context.Context, redisId string, newPassword string, noAuth bool) (taskId int64, errRet error) {
	logId := getLogId(ctx)

	request := redis.NewResetPasswordRequest()
	request.InstanceId = &redisId
	request.Password = &newPassword
	request.NoAuth = &noAuth
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseRedisClient().ResetPassword(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	} else {
		errRet = err
		return
	}

	taskId = *response.Response.TaskId
	return

}

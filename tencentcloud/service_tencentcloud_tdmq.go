package tencentcloud

import (
	"context"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

// basic information

type TdmqService struct {
	client *connectivity.TencentCloudClient
}

// ////////api
// tdmq instance

func (me *TdmqService) DescribeTdmqInstanceById(ctx context.Context,
	clusterId string) (info *tdmq.Cluster, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeClustersRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterIdList = []*string{&clusterId}

	var response *tdmq.DescribeClustersResponse

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeClusters(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq failed, reason: %v", logId, err)
		return nil, false, err
	}

	if len(response.Response.ClusterSet) < 1 {
		return
	}
	has = true
	info = response.Response.ClusterSet[0]
	return
}

func (me *TdmqService) ModifyTdmqInstanceAttribute(ctx context.Context, clusterId, clusterName string,
	remark string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewModifyClusterRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &clusterId
	request.ClusterName = &clusterName
	request.Remark = &remark

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().ModifyCluster(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s modify tdmq failed, reason: %v", logId, err)
		return err
	}
	return
}

func (me *TdmqService) DeleteTdmqInstance(ctx context.Context, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteClusterRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &clusterId
	response, err := me.client.UseTdmqClient().DeleteCluster(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return
}

// tdmq namespace
func (me *TdmqService) CreateTdmqNamespace(ctx context.Context, environName string, msgTtl uint64, clusterId string,
	remark string, retentionPolicy tdmq.RetentionPolicy) (environId string, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewCreateEnvironmentRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.EnvironmentId = &environName
	request.MsgTTL = &msgTtl
	request.ClusterId = &clusterId
	request.Remark = &remark
	request.RetentionPolicy = &retentionPolicy

	var response *tdmq.CreateEnvironmentResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().CreateEnvironment(request)
		if err != nil {
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s create tdmq namespace failed, reason: %v", logId, err)
		errRet = err
		return
	}
	environId = *response.Response.EnvironmentId
	return
}

func (me *TdmqService) DescribeTdmqNamespaceById(ctx context.Context,
	environId string, clusterId string) (info *tdmq.Environment, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeEnvironmentsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.EnvironmentId = &environId
	request.ClusterId = &clusterId

	var response *tdmq.DescribeEnvironmentsResponse

	if err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeEnvironments(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq failed, reason: %v", logId, err)
		return nil, false, err
	}

	if len(response.Response.EnvironmentSet) < 1 {
		return
	}
	has = true
	info = response.Response.EnvironmentSet[0]
	return
}

func (me *TdmqService) ModifyTdmqNamespaceAttribute(ctx context.Context, environId string, msgTtl uint64,
	remark string, clusterId string, retentionPolicy *tdmq.RetentionPolicy) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewModifyEnvironmentAttributesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentId = &environId
	request.MsgTTL = &msgTtl
	request.Remark = &remark
	request.ClusterId = &clusterId
	request.RetentionPolicy = retentionPolicy

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().ModifyEnvironmentAttributes(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s modify tdmq namespace failed, reason: %v", logId, err)
		return err
	}
	return
}

func (me *TdmqService) DeleteTdmqNamespace(ctx context.Context, environId string, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteEnvironmentsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentIds = []*string{&environId}
	request.ClusterId = &clusterId
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().DeleteEnvironments(request)
		if sdkErr, ok := err.(*errors.CloudSDKError); ok {
			if sdkErr.Code == VPCNotFound {
				return nil
			}
		}
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s delete tdmq namespace failed, reason: %v", logId, err)
		return err
	}
	return
}

// tdmq topic
func (me *TdmqService) CreateTdmqTopic(ctx context.Context, environId string, topicName string, partitions uint64,
	topicType int64, remark string, clusterId string, pulsarTopicType int64) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewCreateTopicRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.EnvironmentId = &environId
	request.TopicName = &topicName
	request.Partitions = &partitions
	if topicType != NoneTopicType {
		request.TopicType = common.Uint64Ptr(uint64(topicType))
	}
	request.Remark = &remark
	request.ClusterId = &clusterId
	if pulsarTopicType != NonePulsarTopicType {
		request.TopicType = helper.Int64Uint64(pulsarTopicType)
	}

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().CreateTopic(request)
		if err != nil {
			return retryError(err)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s create tdmq topic failed, reason: %v", logId, err)
		errRet = err
		return
	}
	return
}

func (me *TdmqService) DescribeTdmqTopicById(ctx context.Context,
	environId string, topicName string, clusterId string) (info *tdmq.Topic, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeTopicsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentId = &environId
	request.TopicName = &topicName
	request.ClusterId = &clusterId

	var response *tdmq.DescribeTopicsResponse

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeTopics(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq failed, reason: %v", logId, err)
		return nil, false, err
	}

	if len(response.Response.TopicSets) < 1 {
		return
	}
	has = true
	info = response.Response.TopicSets[0]
	return
}

func (me *TdmqService) ModifyTdmqTopicAttribute(ctx context.Context, environId string, topicName string,
	partitions uint64, remark string, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewModifyTopicRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentId = &environId
	request.TopicName = &topicName
	request.Partitions = &partitions
	request.Remark = &remark
	request.ClusterId = &clusterId

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().ModifyTopic(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s modify tdmq topic failed, reason: %v", logId, err)
		return err
	}
	return
}

func (me *TdmqService) DeleteTdmqTopic(ctx context.Context, environId string, topicName string, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteTopicsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	var (
		topicRecord tdmq.TopicRecord
	)
	topicRecord.TopicName = &topicName
	topicRecord.EnvironmentId = &environId
	request.TopicSets = []*tdmq.TopicRecord{&topicRecord}
	request.ClusterId = &clusterId

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().DeleteTopics(request)
		if err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s delete tdmq topic failed, reason: %v", logId, err)
		return err
	}
	return
}

// tdmq role
func (me *TdmqService) CreateTdmqRole(ctx context.Context, roleName string, clusterId string,
	remark string) (roleId string, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewCreateRoleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.RoleName = &roleName
	request.ClusterId = &clusterId
	request.Remark = &remark

	var response *tdmq.CreateRoleResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().CreateRole(request)
		if err != nil {
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s create tdmq topic failed, reason: %v", logId, err)
		errRet = err
		return
	}
	roleId = *response.Response.RoleName
	return
}

func (me *TdmqService) DescribeTdmqRoleById(ctx context.Context,
	roleName string, clusterId string) (info *tdmq.Role, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeRolesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.RoleName = &roleName
	request.ClusterId = &clusterId

	var response *tdmq.DescribeRolesResponse

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeRoles(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq role failed, reason: %v", logId, err)
		return nil, false, err
	}

	if len(response.Response.RoleSets) < 1 {
		return
	}
	has = true
	info = response.Response.RoleSets[0]
	return
}

func (me *TdmqService) ModifyTdmqRoleAttribute(ctx context.Context, roleName string, clusterId string,
	remark string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewModifyRoleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.RoleName = &roleName
	request.ClusterId = &clusterId
	request.Remark = &remark

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().ModifyRole(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s modify tdmq role failed, reason: %v", logId, err)
		return err
	}
	return
}

func (me *TdmqService) DeleteTdmqRole(ctx context.Context, roleName string, cluserId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteRolesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.RoleNames = []*string{&roleName}
	request.ClusterId = &cluserId

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().DeleteRoles(request)
		if err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s delete tdmq roles failed, reason: %v", logId, err)
		return err
	}
	return
}

// tdmq role
func (me *TdmqService) CreateTdmqNamespaceRoleAttachment(ctx context.Context, environId string,
	roleName string, permissions []*string, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewCreateEnvironmentRoleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.EnvironmentId = &environId
	request.RoleName = &roleName
	request.Permissions = permissions
	request.ClusterId = &clusterId

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().CreateEnvironmentRole(request)
		if err != nil {
			return retryError(err)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s create tdmq topic failed, reason: %v", logId, err)
		errRet = err
		return
	}
	return
}

func (me *TdmqService) DescribeTdmqNamespaceRoleAttachment(ctx context.Context,
	environId string, roleName string, clusterId string) (info *tdmq.EnvironmentRole, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeEnvironmentRolesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentId = &environId
	request.RoleName = &roleName
	request.ClusterId = &clusterId

	var response *tdmq.DescribeEnvironmentRolesResponse

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeEnvironmentRoles(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq environment role failed, reason: %v", logId, err)
		return nil, false, err
	}

	if len(response.Response.EnvironmentRoleSets) < 1 {
		return
	}
	has = true
	info = response.Response.EnvironmentRoleSets[0]
	return
}

func (me *TdmqService) ModifyTdmqNamespaceRoleAttachment(ctx context.Context,
	environId string, roleName string, permissions []*string, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewModifyEnvironmentRoleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.EnvironmentId = &environId
	request.RoleName = &roleName
	request.ClusterId = &clusterId
	request.Permissions = permissions

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().ModifyEnvironmentRole(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s modify tdmq environment role failed, reason: %v", logId, err)
		return err
	}
	return
}

func (me *TdmqService) DeleteTdmqNamespaceRoleAttachment(ctx context.Context, environId string,
	roleName string, cluserId string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteEnvironmentRolesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.EnvironmentId = &environId
	request.RoleNames = []*string{&roleName}
	request.ClusterId = &cluserId

	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().DeleteEnvironmentRoles(request)
		if err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s delete tdmq environments roles failed, reason: %v", logId, err)
		return err
	}
	return
}

// CreateRoute create route
func (me *TdmqService) CreateRoute(ctx context.Context, netType uint64, clusterId string, uniqueVpcId string,
	uniqueSubnetId, remark string) (routeId string, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewCreateRouteRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%+v], reason[%s]\n",
				logId, request.GetAction(), request, errRet.Error())
		}
	}()

	request.NetType = &netType
	request.ClusterId = &clusterId
	request.UniqueVpcId = &uniqueVpcId
	request.UniqueSubnetId = &uniqueSubnetId
	request.Remark = &remark

	var response *tdmq.CreateRouteResponse
	if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().CreateRoute(request)
		if err != nil {
			return retryError(err)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s create tdmq route failed, reason: %v", logId, err)
		errRet = err
		return
	}
	routeId = *response.Response.RouterId
	return
}

// DescribeTdmqRouteById describe route by id
func (me *TdmqService) DescribeTdmqRouteById(ctx context.Context, clusterId, routeId string) (
	info *tdmq.RouteRecord, has bool, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeRouteRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%+v], reason[%s]\n",
				logId, request.GetAction(), request, errRet.Error())
		}
	}()
	request.ClusterId = &clusterId

	var response *tdmq.DescribeRouteResponse

	if err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, err := me.client.UseTdmqClient().DescribeRoute(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		response = result
		return nil
	}); err != nil {
		log.Printf("[CRITAL]%s read tdmq route failed, reason: %v", logId, err)
		return nil, false, err
	}

	if *response.Response.TotalCount < 1 {
		return
	}
	for _, v := range response.Response.RouteRecords {
		if *v.RouterId == routeId {
			info = v
			has = true
			return
		}
	}
	return
}

// DeleteTdmqRoute delete route
func (me *TdmqService) DeleteTdmqRoute(ctx context.Context, request *tdmq.DeleteRouteRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%+v], reason[%s]\n",
				logId, request.GetAction(), request, errRet.Error())
		}
	}()

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		_, err := me.client.UseTdmqClient().DeleteRoute(request)
		if err != nil {
			return retryError(err)
		}
		return nil
	})
	if err != nil {
		errRet = err
		return
	}
	return
}

// DescribeEnvironments describe environments
func (me *TdmqService) DescribeEnvironments(ctx context.Context, request *tdmq.DescribeEnvironmentsRequest) (
	infos []*tdmq.Environment, errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%+v], reason[%s]\n",
				logId, request.GetAction(), request, errRet.Error())
		}
	}()

	result, err := me.client.UseTdmqClient().DescribeEnvironments(request)
	if err != nil {
		errRet = err
		return
	}
	infos = result.Response.EnvironmentSet
	return
}


func (me *TdmqService) DescribeTdmqRabbitmqNodeListByFilter(ctx context.Context, param map[string]interface{}) (rabbitmqNodeList []*tdmq.RabbitMQPrivateNode, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tdmq.NewDescribeRabbitMQNodeListRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "InstanceId" {
			request.InstanceId = v.(*string)
		}
		if k == "NodeName" {
			request.NodeName = v.(*string)
		}
		if k == "Filters" {
			request.Filters = v.([]*tdmq.Filter)
		}
		if k == "SortElement" {
			request.SortElement = v.(*string)
		}
		if k == "SortOrder" {
			request.SortOrder = v.(*string)
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
		response, err := me.client.UseTdmqClient().DescribeRabbitMQNodeList(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || *response.Response.TotalCount == 0 {
			break
		}

		rabbitmqNodeList = append(rabbitmqNodeList, response.Response.NodeList...)
		if len(response.Response.NodeList) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

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

		if response == nil || *response.Response.TotalCount == 0 {
			break
		}

		rabbitmqVipInstance = append(rabbitmqVipInstance, response.Response.Instances...)
		if len(response.Response.Instances) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *TdmqService) DescribeTdmqRabbitmqUserById(ctx context.Context, instanceId, user string) (rabbitmqUser *tdmq.RabbitMQUser, errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDescribeRabbitMQUserRequest()
	request.InstanceId = &instanceId
	if user != "" {
		request.User = &user
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DescribeRabbitMQUser(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.RabbitMQUserList) < 1 {
		return
	}

	rabbitmqUser = response.Response.RabbitMQUserList[0]
	return
}

func (me *TdmqService) DeleteTdmqRabbitmqUserById(ctx context.Context, instanceId, user string) (errRet error) {
	logId := getLogId(ctx)
	request := tdmq.NewDeleteRabbitMQUserRequest()
	request.InstanceId = &instanceId
	request.User = &user

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DeleteRabbitMQUser(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

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
	var iacExtInfo connectivity.IacExtInfo
	iacExtInfo.InstanceId = instanceId
	response, err := me.client.UseTdmqClient(iacExtInfo).DescribeRabbitMQVipInstance(request)
	if err != nil {
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
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TdmqService) DescribeTdmqRabbitmqVirtualHostById(ctx context.Context, instanceId, virtualHost string) (rabbitmqVirtualHost *tdmq.RabbitMQVirtualHostInfo, errRet error) {
	logId := getLogId(ctx)

	request := tdmq.NewDescribeRabbitMQVirtualHostRequest()
	request.InstanceId = &instanceId
	if virtualHost != "" {
		request.VirtualHost = &virtualHost
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DescribeRabbitMQVirtualHost(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.VirtualHostList) < 1 {
		return
	}

	rabbitmqVirtualHost = response.Response.VirtualHostList[0]
	return
}

func (me *TdmqService) DeleteTdmqRabbitmqVirtualHostById(ctx context.Context, instanceId, virtualHost string) (errRet error) {
	logId := getLogId(ctx)

	request := tdmq.NewDeleteRabbitMQVirtualHostRequest()
	request.InstanceId = &instanceId
	request.VirtualHost = &virtualHost

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTdmqClient().DeleteRabbitMQVirtualHost(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}
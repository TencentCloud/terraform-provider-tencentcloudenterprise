package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type TurbofsService struct {
	client *connectivity.TencentCloudClient
}

func (me *TurbofsService) DescribeFileSystem(ctx context.Context, fsId *string) (fs []*turbofs.FileSystemInfo, errRet error) {
	logId := getLogId(ctx)
	request := turbofs.NewDescribeCfsFileSystemsRequest()

	if fsId != nil {
		request.FileSystemId = fsId
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DescribeCfsFileSystems(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	fs = response.Response.FileSystems
	return
}

func (me *TurbofsService) DescribeMountTargets(ctx context.Context, fsId string) (targets []*turbofs.MountInfo, errRet error) {
	logId := getLogId(ctx)
	request := turbofs.NewDescribeMountTargetsRequest()
	request.FileSystemId = &fsId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DescribeMountTargets(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	targets = response.Response.MountTargets
	return
}

func (me *TurbofsService) ModifyFileSystemName(ctx context.Context, fsId, fsName string) error {
	logId := getLogId(ctx)
	request := turbofs.NewUpdateCfsFileSystemNameRequest()
	request.FileSystemId = &fsId
	request.FsName = &fsName

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().UpdateCfsFileSystemName(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) ModifyFileSystemPGroup(ctx context.Context, fsId, pGroupId string) error {
	logId := getLogId(ctx)
	request := turbofs.NewUpdateCfsFileSystemPGroupRequest()
	request.FileSystemId = &fsId
	request.PGroupId = &pGroupId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().UpdateCfsFileSystemPGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) ScaleUpFileSystem(ctx context.Context, fsId string, targetCapacity uint64) error {
	logId := getLogId(ctx)
	request := turbofs.NewScaleUpFileSystemRequest()
	request.FileSystemId = &fsId
	request.TargetCapacity = &targetCapacity

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().ScaleUpFileSystem(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) WaitForFileSystemAvailable(ctx context.Context, fsId string) error {
	logId := getLogId(ctx)
	
	return resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
		fileSystems, errRet := me.DescribeFileSystem(ctx, &fsId)
		if errRet != nil {
			log.Printf("[DEBUG]%s describe file system fail, reason[%s]\n", logId, errRet.Error())
			return retryError(errRet, InternalError)
		}
		if fileSystems == nil || len(fileSystems) < 1 {
			return resource.RetryableError(fmt.Errorf("file system %s not found", fsId))
		}

		fileSystem := fileSystems[0]
		if fileSystem.LifeCycleState == nil {
			return resource.RetryableError(fmt.Errorf("file system %s state is nil", fsId))
		}

		log.Printf("[DEBUG]%s file system %s current state: %s\n", logId, fsId, *fileSystem.LifeCycleState)

		if *fileSystem.LifeCycleState == TURBOFS_LIFECYCLE_STATE_AVAILABLE {
			return nil
		}
		if *fileSystem.LifeCycleState == TURBOFS_LIFECYCLE_STATE_CREATING {
			return resource.RetryableError(fmt.Errorf("file system %s is still creating", fsId))
		}
		if *fileSystem.LifeCycleState == TURBOFS_LIFECYCLE_STATE_EXPANDING {
			return resource.RetryableError(fmt.Errorf("file system %s is still expanding", fsId))
		}
		return resource.NonRetryableError(fmt.Errorf("file system %s creation failed, current state: %s", fsId, *fileSystem.LifeCycleState))
	})
}

func (me *TurbofsService) DeleteFileSystem(ctx context.Context, fsId string) error {
	logId := getLogId(ctx)
	request := turbofs.NewDeleteCfsFileSystemRequest()
	request.FileSystemId = &fsId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DeleteCfsFileSystem(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) CreateCfsPGroup(ctx context.Context, name, descInfo *string) (pgroupId *string, errRet error) {
	logId := getLogId(ctx)
	request := turbofs.NewCreateCfsPGroupRequest()
	request.Name = name
	request.DescInfo = descInfo

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().CreateCfsPGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		errRet = fmt.Errorf("create turbofs pgroup failed, response is nil.")
	}
	pgroupId = response.Response.PGroupId

	return
}

func (me *TurbofsService) DescribeCfsPGroup(ctx context.Context, id string) (pGroupList []*turbofs.PGroupInfo, errRet error) {
	logId := getLogId(ctx)
	request := turbofs.NewDescribeCfsPGroupsRequest()

	// 设置分页参数
	var (
		offset uint64 = 0
		limit  uint64 = 20
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for {
		request.Offset = &offset
		request.Limit = &limit

		ratelimit.Check(request.GetAction())
		response, err := me.client.UseTurbofsClient().DescribeCfsPGroups(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || response.Response.PGroupList == nil {
			break
		}

		// 在当前页结果中查找匹配的权限组
		for _, pGroup := range response.Response.PGroupList {
			if pGroup.PGroupId != nil && *pGroup.PGroupId == id {
				pGroupList = append(pGroupList, pGroup)
				return // 找到第一个匹配项后立即返回
			}
		}

		// 如果当前页结果数量小于limit，说明已经到最后一页
		if len(response.Response.PGroupList) < int(limit) {
			break
		}

		// 继续查询下一页
		offset += limit
	}

	return
}

func (me *TurbofsService) DeleteCfsPGroup(ctx context.Context, id *string) error {
	logId := getLogId(ctx)
	request := turbofs.NewDeleteCfsPGroupRequest()
	request.PGroupId = id
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DeleteCfsPGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) DescribeCfsRule(ctx context.Context, accessGroupId, accessRuleId string) (accessRules []*turbofs.PGroupRuleInfo, errRet error) {
	logId := getLogId(ctx)
	request := turbofs.NewDescribeCfsRulesRequest()
	request.PGroupId = &accessGroupId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DescribeCfsRules(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		if sdkErr, ok := err.(*errors.CloudSDKError); ok {
			if sdkErr.Code == CfsInvalidPgroup {
				return nil, nil
			}
		}
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	accessRules = make([]*turbofs.PGroupRuleInfo, 0)
	for _, accessRule := range response.Response.RuleList {
		if accessRuleId != "" && *accessRule.RuleId != accessRuleId {
			continue
		}
		accessRules = append(accessRules, accessRule)
	}
	return
}

func (me *TurbofsService) DeleteAccessRule(ctx context.Context, accessGroupId, accessRuleId string) error {
	logId := getLogId(ctx)
	request := turbofs.NewDeleteCfsRuleRequest()
	request.PGroupId = &accessGroupId
	request.RuleId = &accessRuleId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTurbofsClient().DeleteCfsRule(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *TurbofsService) DescribeCfsAutoSnapshotPolicyById(ctx context.Context,
	autoSnapshotPolicyId string) (autoSnapshotPolicy *turbofs.AutoSnapshotPolicyInfo, errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDescribeAutoSnapshotPoliciesRequest()
	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	instances := make([]*turbofs.AutoSnapshotPolicyInfo, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseTurbofsClient().DescribeAutoSnapshotPolicies(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.AutoSnapshotPolicies) < 1 {
			break
		}
		instances = append(instances, response.Response.AutoSnapshotPolicies...)
		if len(response.Response.AutoSnapshotPolicies) < int(limit) {
			break
		}

		offset += limit
	}

	if len(instances) < 1 {
		return
	}
	autoSnapshotPolicy = instances[0]
	return
}

func (me *TurbofsService) DeleteCfsAutoSnapshotPolicyById(ctx context.Context, autoSnapshotPolicyId string) (errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDeleteAutoSnapshotPolicyRequest()
	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().DeleteAutoSnapshotPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TurbofsService) DescribeCfsAutoSnapshotPolicyAttachmentById(ctx context.Context, autoSnapshotPolicyId string,
	fileSystemId string) (snapShotId, fsId *string, errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDescribeAutoSnapshotPoliciesRequest()
	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId,
				request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	instances := make([]*turbofs.AutoSnapshotPolicyInfo, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseTurbofsClient().DescribeAutoSnapshotPolicies(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.AutoSnapshotPolicies) < 1 {
			break
		}
		instances = append(instances, response.Response.AutoSnapshotPolicies...)
		if len(response.Response.AutoSnapshotPolicies) < int(limit) {
			break
		}

		offset += limit
	}

	if len(instances) < 1 {
		return
	}

	autoSnapshotPolicy := instances[0]

	for _, fsInfo := range autoSnapshotPolicy.FsInfo {
		if fsInfo.FileSystemId != nil && *fsInfo.FileSystemId == fileSystemId {
			snapShotId = autoSnapshotPolicy.AutoSnapshotPolicyId
			fsId = fsInfo.FileSystemId
			break
		}
	}

	return
}

func (me *TurbofsService) DeleteCfsAutoSnapshotPolicyAttachmentById(ctx context.Context, autoSnapshotPolicyId string, fileSystemIds string) (errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewUnbindAutoSnapshotPolicyRequest()
	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId
	request.FileSystemIds = &fileSystemIds

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().UnbindAutoSnapshotPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TurbofsService) DescribeCfsSnapshotById(ctx context.Context, jobId string) (snapshot *turbofs.SnapList,
	errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDescribeCfsSnapshotsRequest()
	request.JobId = &jobId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	instances := make([]*turbofs.SnapList, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseTurbofsClient().DescribeCfsSnaps(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Snaps) < 1 {
			break
		}
		instances = append(instances, response.Response.Snaps...)
		if len(response.Response.Snaps) < int(limit) {
			break
		}

		offset += limit
	}

	if len(instances) < 1 {
		return
	}
	snapshot = instances[0]
	return
}

func (me *TurbofsService) DeleteCfsSnapshotById(ctx context.Context, snapshotId string) (errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDeleteCfsSnapRequest()
	request.SnapshotId = &snapshotId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().DeleteCfsSnap(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TurbofsService) CfsSnapshotStateRefreshFunc(jobId string, failStates []string) resource.StateRefreshFunc {
	mp := make(map[uint64]string, 0)
	mp[1] = "CREATING"
	mp[2] = "available"
	return func() (interface{}, string, error) {
		ctx := contextNil

		object, err := me.DescribeCfsSnapshotById(ctx, jobId)

		if err != nil {
			return nil, "", err
		}

		return object, mp[*object.Status], nil
	}
}
func (me *TurbofsService) DescribeCfsMountTargetsById(ctx context.Context, fileSystemId string) (mountTargets []*turbofs.MountInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = turbofs.NewDescribeMountTargetsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.FileSystemId = helper.String(fileSystemId)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().DescribeMountTargets(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	mountTargets = response.Response.MountTargets

	return
}

func (me *TurbofsService) DescribeCfsFileSystemClientsById(ctx context.Context, fileSystemId string) (fileSystemClients []*turbofs.FileSystemClient, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = turbofs.NewDescribeCfsFileSystemClientsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.FileSystemId = helper.String(fileSystemId)

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().DescribeCfsFileSystemClients(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	fileSystemClients = response.Response.ClientList

	return
}

/*
func (me *TurbofsService) DescribeCfsUserQuotaById(ctx context.Context, fileSystemId string, userType string, userId string) (userQuota *turbofs.UserQuota, errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDescribeUserQuotaRequest()
	request.FileSystemId = &fileSystemId

	typeFilter := turbofs.Filter{
		Name:   helper.String("UserType"),
		Values: []*string{helper.String(userType)},
	}
	request.Filters = append(request.Filters, &typeFilter)

	idFilter := turbofs.Filter{
		Name:   helper.String("UserId"),
		Values: []*string{helper.String(userId)},
	}
	request.Filters = append(request.Filters, &idFilter)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	instances := make([]*turbofs.UserQuota, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseTurbofsClient().DescribeUserQuota(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.UserQuotaInfo) < 1 {
			break
		}
		instances = append(instances, response.Response.UserQuotaInfo...)
		if len(response.Response.UserQuotaInfo) < int(limit) {
			break
		}

		offset += limit
	}

	if len(instances) < 1 {
		return
	}
	userQuota = instances[0]
	return
}

*/

/*
func (me *TurbofsService) DeleteCfsUserQuotaById(ctx context.Context, fileSystemId string, userType string, userId string) (errRet error) {
	logId := getLogId(ctx)

	request := turbofs.NewDeleteUserQuotaRequest()
	request.FileSystemId = &fileSystemId
	request.UserType = &userType
	request.UserId = &userId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTurbofsClient().DeleteUserQuota(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

*/

func (me *TurbofsService) DescribeCfsAvailableZoneByFilter(ctx context.Context) (availableZone []*turbofs.AvailableRegion, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = turbofs.NewDescribeAvailableZoneInfoRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	response, err := me.client.UseTurbofsClient().DescribeAvailableZoneInfo(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	availableZone = response.Response.RegionZones

	return
}

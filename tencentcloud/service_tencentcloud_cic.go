package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type CicService struct {
	client *connectivity.TencentCloudClient
}

func (me *CicService) DescribeCicIdentityCenter(ctx context.Context) (ret *cic.DescribeIdentityCenterResponse, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewDescribeIdentityCenterRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().DescribeIdentityCenter(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	ret = response
	return
}

// DescribeCicExternalSamlIdentityProviderById retrieves the SAML service provider information for a given zone ID.
func (me *CicService) DescribeCicExternalSamlIdentityProviderById(ctx context.Context, zoneId string) (
	samlServiceProvider *cic.SAMLServiceProvider, errRet error) {

	logId := getLogId(ctx)
	request := cic.NewGetZoneSAMLServiceProviderInfoRequest()
	request.ZoneId = helper.String(zoneId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetZoneSAMLServiceProviderInfo(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), errors.WithStack(err))
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	samlServiceProvider = response.Response.SAMLServiceProvider

	return
}

// DescribeCicExternalSamlIdentityProviderConfigById  retrieves the SAML identity provider configuration for a given zone ID.
func (me *CicService) DescribeCicExternalSamlIdentityProviderConfigById(ctx context.Context, zoneId string) (
	ret *cic.SAMLIdentityProviderConfiguration, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewGetExternalSAMLIdentityProviderRequest()
	request.ZoneId = helper.String(zoneId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetExternalSAMLIdentityProvider(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	ret = response.Response.SAMLIdentityProviderConfiguration
	return
}

// DescribeCicUserById retrieves a user by its ID.
func (me *CicService) DescribeCicUserById(ctx context.Context, zoneId string, userId string) (
	ret *cic.UserInfo, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewGetUserRequest()
	request.UserId = helper.String(userId)
	request.ZoneId = helper.String(zoneId)
	response := cic.NewGetUserResponse()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseCicClient().GetUser(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		errRet = err
		log.Printf("[CRITAL]%s update identity center user failed, reason:%+v", logId, err)
		return
	}

	if response.Response == nil {
		return
	}

	ret = response.Response.UserInfo
	return
}

// DescribeCicRoleAssignmentById retrieves a role assignment by its ID.
func (me *CicService) DescribeCicRoleAssignmentById(ctx context.Context, roleAssignmentId string) (ret *cic.ListRoleAssignmentsResponseParams, errRet error) {
	logId := getLogId(ctx)

	idSplit := strings.Split(roleAssignmentId, FILED_SP)
	if len(idSplit) != 6 {
		errRet = fmt.Errorf("roleAssignmentId is broken,%s", roleAssignmentId)
		return
	}

	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]
	targetType := idSplit[2]
	targetUinString := idSplit[3]
	principalType := idSplit[4]
	principalId := idSplit[5]

	request := cic.NewListRoleAssignmentsRequest()
	request.ZoneId = helper.String(zoneId)
	request.RoleConfigurationId = helper.String(roleConfigurationId)
	request.TargetType = helper.String(targetType)
	targetUin, err := strconv.ParseInt(targetUinString, 10, 64)
	if err != nil {
		errRet = err
		return
	}
	request.TargetUin = helper.Int64(targetUin)
	request.PrincipalType = helper.String(principalType)
	request.PrincipalId = helper.String(principalId)
	request.MaxResults = helper.Int64(10)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().ListRoleAssignments(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	ret = response.Response
	return
}

// DescribeCicGroupById retrieves a group by its ID.
func (me *CicService) DescribeCicGroupById(ctx context.Context, zoneId string, groupId string) (ret *cic.GroupInfo, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewGetGroupRequest()
	request.ZoneId = helper.String(zoneId)
	request.GroupId = helper.String(groupId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetGroup(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	ret = response.Response.GroupInfo
	return
}

func (me *CicService) GetAssignmentTaskStatus(ctx context.Context, zoneId, taskId string) (taskStatus *cic.TaskStatus, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewGetTaskStatusRequest()
	request.ZoneId = helper.String(zoneId)
	request.TaskId = helper.String(taskId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, e := me.client.UseCicClient().GetTaskStatus(request)
	if e != nil {
		errRet = e
		return
	}

	if response.Response != nil {
		taskStatus = response.Response.TaskStatus
	}
	return

}

func (me *CicService) AssignmentTaskStatusStateRefreshFunc(zoneId, taskId string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := contextNil

		var object *cic.TaskStatus
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := me.GetAssignmentTaskStatus(ctx, zoneId, taskId)
			if e != nil {
				return retryError(e)
			}
			object = result
			return nil
		})
		if err != nil {
			return nil, "", err
		}

		return object, *object.Status, nil
	}
}

func (me *CicService) DescribeCicRoleConfigurationById(ctx context.Context, zoneId string,
	roleConfigurationId string) (ret *cic.RoleConfiguration, errRet error) {

	logId := getLogId(ctx)

	request := cic.NewGetRoleConfigurationRequest()
	request.ZoneId = helper.String(zoneId)
	request.RoleConfigurationId = helper.String(roleConfigurationId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetRoleConfiguration(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	ret = response.Response.RoleConfigurationInfo
	return
}

func (me *CicService) DescribeCicRoleConfigurationPermissionPolicyAttachmentById(
	ctx context.Context, zoneId, roleConfigurationId, rolePolicyType string) (
	ret *cic.ListPermissionPoliciesInRoleConfigurationResponseParams, errRet error) {

	logId := getLogId(ctx)

	request := cic.NewListPermissionPoliciesInRoleConfigurationRequest()
	request.ZoneId = helper.String(zoneId)
	request.RoleConfigurationId = helper.String(roleConfigurationId)
	request.RolePolicyType = helper.String(rolePolicyType)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().ListPermissionPoliciesInRoleConfiguration(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId,
		request.GetAction(), request.ToJsonString(), response.ToJsonString())

	ret = response.Response
	return
}

func (me *CicService) DescribeCicScimCredentialById(ctx context.Context, zoneId string, credentialId string) (
	ret *cic.SCIMCredential, errRet error) {

	logId := getLogId(ctx)

	request := cic.NewListSCIMCredentialsRequest()
	request.ZoneId = helper.String(zoneId)
	request.CredentialId = helper.String(credentialId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().ListSCIMCredentials(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.SCIMCredentials) < 1 {
		return
	}

	ret = response.Response.SCIMCredentials[0]
	return
}

func (me *CicService) DescribeCicScimSynchronizationStatusById(ctx context.Context, zoneId string) (ret *cic.GetSCIMSynchronizationStatusResponseParams, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewGetSCIMSynchronizationStatusRequest()
	request.ZoneId = helper.String(zoneId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetSCIMSynchronizationStatus(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	ret = response.Response
	return
}

func (me *CicService) DescribeCicUserGroupAttachmentById(ctx context.Context, zoneId, groupId, userId string) (
	joinedGroup *cic.JoinedGroups, errRet error) {
	logId := getLogId(ctx)

	request := cic.NewListJoinedGroupsForUserRequest()
	request.ZoneId = helper.String(zoneId)
	request.UserId = helper.String(userId)
	request.MaxResults = helper.Int64(100)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCicClient().ListJoinedGroupsForUser(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		for _, v := range response.Response.JoinedGroups {
			if *v.GroupId == groupId {
				joinedGroup = v
				return
			}
		}
		if len(response.Response.JoinedGroups) < int(*request.MaxResults) {
			break
		} else {
			request.NextToken = response.Response.NextToken
		}
	}
	return
}

func (me *CicService) DescribeCicUserSyncProvisioningById(ctx context.Context, zoneId, userProvisioningId string) (
	ret *cic.UserProvisioning, errRet error) {

	logId := getLogId(ctx)

	request := cic.NewGetUserSyncProvisioningRequest()
	request.ZoneId = helper.String(zoneId)
	request.UserProvisioningId = helper.String(userProvisioningId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCicClient().GetUserSyncProvisioning(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response == nil {
		return
	}

	ret = response.Response.UserProvisioning
	return
}

func (me *CicService) DescribeCicGroupsByFilter(ctx context.Context, param map[string]interface{}) (groups []*cic.GroupInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cic.NewListGroupsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "ZoneId" {
			request.ZoneId = v.(*string)
		}
		if k == "Filter" {
			request.Filter = v.(*string)
		}
		if k == "GroupType" {
			request.GroupType = v.(*string)
		}
		if k == "FilterUsers" {
			request.FilterUsers = v.([]*string)
		}
		if k == "SortField" {
			request.SortField = v.(*string)
		}
		if k == "SortType" {
			request.SortType = v.(*string)
		}
	}

	groups = make([]*cic.GroupInfo, 0)
	for {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCicClient().ListGroups(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil {
			return
		}

		groups = append(groups, response.Response.Groups...)

		if response.Response.IsTruncated != nil {
			if *response.Response.IsTruncated {
				request.NextToken = response.Response.NextToken
			} else {
				break
			}
		} else {
			errRet = fmt.Errorf("ListGroups IsTruncated is nil")
			return
		}
	}

	return
}

func (me *CicService) DescribeCicUsersByFilter(ctx context.Context, param map[string]interface{}) (users []*cic.UserInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cic.NewListUsersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "ZoneId" {
			request.ZoneId = v.(*string)
		}
		if k == "Filter" {
			request.Filter = v.(*string)
		}
		if k == "FilterGroups" {
			request.FilterGroups = v.([]*string)
		}
		if k == "UserStatus" {
			request.UserStatus = v.(*string)
		}
		if k == "UserType" {
			request.UserType = v.(*string)
		}
		if k == "SortField" {
			request.SortField = v.(*string)
		}
		if k == "SortType" {
			request.SortType = v.(*string)
		}
	}

	users = make([]*cic.UserInfo, 0)
	for {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCicClient().ListUsers(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil {
			return
		}

		users = append(users, response.Response.Users...)

		if response.Response.IsTruncated != nil {
			if *response.Response.IsTruncated {
				request.NextToken = response.Response.NextToken
			} else {
				break
			}
		} else {
			errRet = fmt.Errorf("ListUsers IsTruncated is nil")
			return
		}
	}

	return
}

func (me *CicService) DescribeCicRoleConfigurationsByFilter(ctx context.Context, param map[string]interface{}) (roleConfigurations []*cic.RoleConfiguration, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cic.NewListRoleConfigurationsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "ZoneId" {
			request.ZoneId = v.(*string)
		}
		if k == "Filter" {
			request.Filter = v.(*string)
		}
		if k == "FilterTargets" {
			request.FilterTargets = v.([]*int64)
		}
		if k == "PrincipalId" {
			request.PrincipalId = v.(*string)
		}
	}

	roleConfigurations = make([]*cic.RoleConfiguration, 0)
	for {
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCicClient().ListRoleConfigurations(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil {
			return
		}

		roleConfigurations = append(roleConfigurations, response.Response.RoleConfigurations...)

		if response.Response.IsTruncated != nil {
			if *response.Response.IsTruncated {
				request.NextToken = response.Response.NextToken
			} else {
				break
			}
		} else {
			errRet = fmt.Errorf("ListRoleConfigurations IsTruncated is nil")
			return
		}
	}

	return
}

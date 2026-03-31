package tencentcloud

import (
	"context"
	"log"
	"strconv"
	"strings"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/sdk/common"

	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type DasbService struct {
	client *connectivity.TencentCloudClient
}

func (me *DasbService) DescribeDasbAclById(ctx context.Context, aclId string) (acl *bhsaas.Acl, errRet error) {
	logId := getLogId(ctx)
	request := bhsaas.NewDescribeAclsRequest()
	aclIdInt, _ := strconv.ParseUint(aclId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(aclIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeAcls(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.AclSet) != 1 {
		return
	}

	acl = response.Response.AclSet[0]
	return
}

func (me *DasbService) DeleteDasbAclById(ctx context.Context, aclId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteAclsRequest()
	aclIdInt, _ := strconv.ParseUint(aclId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(aclIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteAcls(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbCmdTemplateById(ctx context.Context, templateId string) (CmdTemplate *bhsaas.CmdTemplate, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeCmdTemplatesRequest()
	templateIdInt, _ := strconv.ParseUint(templateId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(templateIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeCmdTemplates(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.CmdTemplateSet) != 1 {
		return
	}

	CmdTemplate = response.Response.CmdTemplateSet[0]
	return
}

func (me *DasbService) DeleteDasbCmdTemplateById(ctx context.Context, templateId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteCmdTemplatesRequest()
	templateIdInt, _ := strconv.ParseUint(templateId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(templateIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteCmdTemplates(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbDeviceGroupById(ctx context.Context, deviceGroupId string) (DeviceGroup *bhsaas.Group, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeDeviceGroupsRequest()
	deviceGroupIdInt, _ := strconv.ParseUint(deviceGroupId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceGroupIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeDeviceGroups(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.GroupSet) != 1 {
		return
	}

	DeviceGroup = response.Response.GroupSet[0]
	return
}

func (me *DasbService) DeleteDasbDeviceGroupById(ctx context.Context, deviceGroupId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteDeviceGroupsRequest()
	deviceGroupIdInt, _ := strconv.ParseUint(deviceGroupId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceGroupIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteDeviceGroups(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbUserById(ctx context.Context, userId string) (user *bhsaas.User, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeUsersRequest()
	userIdInt, _ := strconv.ParseUint(userId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(userIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeUsers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.UserSet) != 1 {
		return
	}

	user = response.Response.UserSet[0]
	return
}

func (me *DasbService) DeleteDasbUserById(ctx context.Context, userId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteUsersRequest()
	userIdInt, _ := strconv.ParseUint(userId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(userIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteUsers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbDeviceAccountById(ctx context.Context, deviceAccountId string) (DeviceAccount *bhsaas.DeviceAccount, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeDeviceAccountsRequest()
	deviceAccountIdInt, _ := strconv.ParseUint(deviceAccountId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceAccountIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeDeviceAccounts(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.DeviceAccountSet) != 1 {
		return
	}

	DeviceAccount = response.Response.DeviceAccountSet[0]
	return
}

func (me *DasbService) DeleteDasbDeviceAccountById(ctx context.Context, deviceAccountId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteDeviceAccountsRequest()
	deviceAccountIdInt, _ := strconv.ParseUint(deviceAccountId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceAccountIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteDeviceAccounts(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbDeviceGroupMembersById(ctx context.Context, deviceGroupId string) (DeviceGroupMembers []uint64, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeDeviceGroupMembersRequest()
	deviceGroupIdInt, _ := strconv.ParseUint(deviceGroupId, 10, 64)
	request.Id = &deviceGroupIdInt
	request.Bound = common.BoolPtr(true)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeDeviceGroupMembers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.DeviceSet) < 1 {
		return
	}

	for _, item := range response.Response.DeviceSet {
		if item.Id != nil {
			DeviceGroupMembers = append(DeviceGroupMembers, *item.Id)
		}
	}

	return
}

func (me *DasbService) DeleteDasbDeviceGroupMembersById(ctx context.Context, deviceGroupId, memberIdSetStr string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteDeviceGroupMembersRequest()
	deviceGroupIdInt, _ := strconv.ParseUint(deviceGroupId, 10, 64)
	request.Id = &deviceGroupIdInt
	memberIdSet := strings.Split(memberIdSetStr, COMMA_SP)
	tmpList := make([]*uint64, 0)
	for _, item := range memberIdSet {
		itemInt, _ := strconv.ParseUint(item, 10, 64)
		tmpList = append(tmpList, &itemInt)
	}

	request.MemberIdSet = tmpList

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteDeviceGroupMembers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbUserGroupMembersById(ctx context.Context, userGroupId string) (UserGroupMembers []uint64, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeUserGroupMembersRequest()
	userGroupIdInt, _ := strconv.ParseUint(userGroupId, 10, 64)
	request.Id = &userGroupIdInt
	request.Bound = common.BoolPtr(true)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeUserGroupMembers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.UserSet) < 1 {
		return
	}

	for _, item := range response.Response.UserSet {
		if item.Id != nil {
			UserGroupMembers = append(UserGroupMembers, *item.Id)
		}
	}

	return
}

func (me *DasbService) DeleteDasbUserGroupMembersById(ctx context.Context, userGroupId, memberIdSetStr string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteUserGroupMembersRequest()
	userGroupIdInt, _ := strconv.ParseUint(userGroupId, 10, 64)
	request.Id = &userGroupIdInt
	memberIdSet := strings.Split(memberIdSetStr, COMMA_SP)
	tmpList := make([]*uint64, 0)
	for _, item := range memberIdSet {
		itemInt, _ := strconv.ParseUint(item, 10, 64)
		tmpList = append(tmpList, &itemInt)
	}

	request.MemberIdSet = tmpList

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteUserGroupMembers(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbResourceById(ctx context.Context, resourceId string) (Resource *bhsaas.Resource, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeResourcesRequest()
	request.ResourceIds = common.StringPtrs([]string{resourceId})

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeResources(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.ResourceSet) != 1 {
		return
	}

	Resource = response.Response.ResourceSet[0]
	return
}

func (me *DasbService) DescribeDasbDeviceById(ctx context.Context, deviceId string) (device *bhsaas.Device, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeDevicesRequest()
	deviceIdInt, _ := strconv.ParseUint(deviceId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeDevices(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.DeviceSet) != 1 {
		return
	}

	device = response.Response.DeviceSet[0]
	return
}

func (me *DasbService) DescribeDasbDeviceByResourceId(ctx context.Context, resourceId string) (deviceSet []*bhsaas.Device, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeDevicesRequest()
	request.ResourceIdSet = common.StringPtrs([]string{resourceId})

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 100
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseBhsaasClient().DescribeDevices(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.DeviceSet) < 1 {
			break
		}

		deviceSet = append(deviceSet, response.Response.DeviceSet...)
		if len(response.Response.DeviceSet) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *DasbService) DeleteDasbDeviceById(ctx context.Context, deviceId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteDevicesRequest()
	deviceIdInt, _ := strconv.ParseUint(deviceId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteDevices(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DescribeDasbUserGroupById(ctx context.Context, userGroupId string) (UserGroup *bhsaas.Group, errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDescribeUserGroupsRequest()
	userGroupIdInt, _ := strconv.ParseUint(userGroupId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(userGroupIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DescribeUserGroups(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.GroupSet) != 1 {
		return
	}

	UserGroup = response.Response.GroupSet[0]
	return
}

func (me *DasbService) DeleteDasbUserGroupById(ctx context.Context, userGroupId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewDeleteUserGroupsRequest()
	userGroupIdInt, _ := strconv.ParseUint(userGroupId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(userGroupIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().DeleteUserGroups(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DeleteDasbBindDeviceAccountPrivateKeyById(ctx context.Context, deviceAccountId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewResetDeviceAccountPrivateKeyRequest()
	deviceAccountIdInt, _ := strconv.ParseUint(deviceAccountId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceAccountIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().ResetDeviceAccountPrivateKey(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DasbService) DeleteDasbBindDeviceAccountPasswordById(ctx context.Context, deviceAccountId string) (errRet error) {
	logId := getLogId(ctx)

	request := bhsaas.NewResetDeviceAccountPasswordRequest()
	deviceAccountIdInt, _ := strconv.ParseUint(deviceAccountId, 10, 64)
	request.IdSet = []*uint64{common.Uint64Ptr(deviceAccountIdInt)}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseBhsaasClient().ResetDeviceAccountPassword(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"log"
	"reflect"
	"strconv"
	"strings"

	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type CamService struct {
	client *connectivity.TencentCloudClient
}

// Document types for parsing PolicyDocument
type CamPrincipal struct {
	Qcs     []string `json:"qcs"`
	Service []string `json:"service"`
}

type CamStatement struct {
	Resource  interface{}  `json:"resource"`
	Action    interface{}  `json:"action"`
	Principal CamPrincipal `json:"principal"`
}

type CamDocument struct {
	Version   string         `json:"version"`
	Statement []CamStatement `json:"statement"`
}

func (me *CamService) DescribeRoleById(ctx context.Context, roleId string) (camInstance *cam.RoleInfo, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewDescribeRoleListRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM) //to save in extension
	result := make([]*cam.RoleInfo, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().DescribeRoleList(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}
		for _, role := range response.Response.List {
			if *role.RoleId == roleId {
				result = append(result, role)
			}
		}
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}

	if len(result) == 0 {
		return
	}
	camInstance = result[0]
	return
}

func (me *CamService) DescribeRolesByFilter(ctx context.Context, params map[string]interface{}) (roles []*cam.RoleInfo, errRet error) {
	logId := getLogId(ctx)
	//need travel
	request := cam.NewDescribeRoleListRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)
	roles = make([]*cam.RoleInfo, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().DescribeRoleList(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}

		for _, role := range response.Response.List {
			if params["role_id"] != nil {
				if *role.RoleId != params["role_id"].(string) {
					continue
				}
			}
			if params["name"] != nil {
				if *role.RoleName != params["name"].(string) {
					continue
				}
			}
			if params["description"] != nil {
				if *role.Description != params["description"].(string) {
					continue
				}
			}
			roles = append(roles, role)
		}
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	return
}

func (me *CamService) DeleteRoleById(ctx context.Context, roleId string) error {

	logId := getLogId(ctx)
	request := cam.NewDeleteRoleRequest()
	request.RoleId = &roleId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().DeleteRole(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CamService) DeleteRoleByName(ctx context.Context, roleName string) error {

	logId := getLogId(ctx)
	request := cam.NewDeleteRoleRequest()
	request.RoleName = &roleName
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().DeleteRole(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CamService) decodeCamPolicyAttachmentId(id string) (instanceId string, policyId64 uint64, errRet error) {
	items := strings.Split(id, "#")
	if len(items) != 2 {
		return instanceId, policyId64, fmt.Errorf(" id is not exist %s", id)
	}
	instanceId = items[0]
	policyId, e := strconv.Atoi(items[1])
	if e != nil {
		errRet = e
		return
	}
	policyId64 = uint64(policyId)
	return
}

func (me *CamService) DescribeRolePolicyAttachmentByName(ctx context.Context, roleName string, params map[string]interface{}) (policyOfRole *cam.AttachedPolicyOfRole, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewListAttachedRolePoliciesRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)
	result := make([]*cam.AttachedPolicyOfRole, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		request.RoleName = &roleName
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListAttachedRolePolicies(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "ResourceNotFound") || errCode == "InvalidParameter.RoleNotExist" {
					return
				}
			}
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}
		for _, policy := range response.Response.List {
			policyName, ok := params["policy_name"]
			if ok && *policy.PolicyName == policyName.(string) {
				result = append(result, policy)
			}
		}
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}

	if len(result) == 0 {
		return
	}
	policyOfRole = result[0]
	return
}

func (me *CamService) DescribeRolePolicyAttachmentById(ctx context.Context, rolePolicyAttachmentId string) (policyOfRole *cam.AttachedPolicyOfRole, errRet error) {
	logId := getLogId(ctx)
	roleId, policyId, e := me.decodeCamPolicyAttachmentId(rolePolicyAttachmentId)
	if e != nil {
		return nil, e
	}

	// Use ListEntitiesForPolicy to check if attachment exists
	entity, err := me.listEntitiesForPolicy(ctx, policyId, "Role", roleId)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return
	}

	// Build AttachedPolicyOfRole from entity
	policyOfRole = &cam.AttachedPolicyOfRole{
		PolicyId: &policyId,
	}

	// If entity has Name field, use it to populate PolicyName
	if entity.Name != nil {
		policyOfRole.PolicyName = entity.Name
	}

	// Get additional info from policy if needed
	policyInfo, err := me.DescribePolicyById(ctx, fmt.Sprintf("%d", policyId))
	if err != nil {
		log.Printf("[WARN]%s failed to get policy info for policy %d, reason: %s\n", logId, policyId, err.Error())
	} else if policyInfo != nil && policyInfo.Response != nil {
		if policyInfo.Response.PolicyName != nil {
			policyOfRole.PolicyName = policyInfo.Response.PolicyName
		}
		if policyInfo.Response.AddTime != nil {
			policyOfRole.AddTime = policyInfo.Response.AddTime
		}
		if policyInfo.Response.CreateMode != nil {
			policyOfRole.CreateMode = policyInfo.Response.CreateMode
		}
		// Note: GetPolicy returns Type (uint64), but AttachedPolicyOfRole expects PolicyType (string)
		// We'll convert Type to a string representation
		if policyInfo.Response.Type != nil {
			typeStr := fmt.Sprintf("%d", *policyInfo.Response.Type)
			policyOfRole.PolicyType = &typeStr
		}
	}

	return
}

func (me *CamService) DescribeRolePolicyAttachmentsByFilter(ctx context.Context, params map[string]interface{}) (policyOfRoles []*cam.AttachedPolicyOfRole, errRet error) {
	logId := getLogId(ctx)
	roleId := params["role_id"].(string)
	policyOfRoles = make([]*cam.AttachedPolicyOfRole, 0)

	// Get role info to get role name
	role, err := me.DescribeRoleById(ctx, roleId)
	if err != nil {
		return nil, err
	}
	if role == nil || role.RoleName == nil {
		return
	}
	roleName := role.RoleName

	// Use ListAttachedRolePolicies to list all policies attached to the role
	request := cam.NewListAttachedRolePoliciesRequest()
	request.RoleName = roleName
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)

	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListAttachedRolePolicies(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				if strings.Contains(errCode, "ResourceNotFound") || errCode == "InvalidParameter.RoleNotExist" {
					return
				}
			}
			return nil, err
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || len(response.Response.List) == 0 {
			break
		}

		// Apply filters if provided
		for _, policy := range response.Response.List {
			// Filter by policy_id if provided
			if params["policy_id"] != nil {
				if policy.PolicyId == nil || *policy.PolicyId != params["policy_id"].(uint64) {
					continue
				}
			}

			// Filter by create_mode if provided
			if params["create_mode"] != nil {
				if policy.CreateMode == nil || int(*policy.CreateMode) != params["create_mode"].(int) {
					continue
				}
			}

			policyOfRoles = append(policyOfRoles, policy)
		}

		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart++
	}

	return
}

func (me *CamService) DeleteRolePolicyAttachmentById(ctx context.Context, rolePolicyAttachmentId string) error {
	logId := getLogId(ctx)
	roleIdStr, policyId, e := me.decodeCamPolicyAttachmentId(rolePolicyAttachmentId)
	if e != nil {
		return e
	}

	roleIdUint, e := strconv.ParseUint(roleIdStr, 10, 64)
	if e != nil {
		return e
	}

	request := cam.NewDetachRolePoliciesRequest()
	request.RoleId = &roleIdUint
	request.PolicyId = []*uint64{&policyId}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().DetachRolePolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CamService) DescribeGroupPolicyAttachmentById(ctx context.Context, groupPolicyAttachmentId string) (policyResults *cam.AttachPolicyInfo, errRet error) {
	logId := getLogId(ctx)
	groupId, policyId, e := me.decodeCamPolicyAttachmentId(groupPolicyAttachmentId)
	if e != nil {
		errRet = e
		return
	}

	// Use ListEntitiesForPolicy to check if attachment exists
	entity, err := me.listEntitiesForPolicy(ctx, policyId, "Group", groupId)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return
	}

	// Build AttachPolicyInfo from entity
	policyResults = &cam.AttachPolicyInfo{
		PolicyId: &policyId,
	}

	// If entity has Name field, use it to populate PolicyName
	if entity.Name != nil {
		policyResults.PolicyName = entity.Name
	}

	// Get additional info from policy if needed
	policyInfo, err := me.DescribePolicyById(ctx, fmt.Sprintf("%d", policyId))
	if err != nil {
		log.Printf("[WARN]%s failed to get policy info for policy %d, reason: %s\n", logId, policyId, err.Error())
	} else if policyInfo != nil && policyInfo.Response != nil {
		if policyInfo.Response.PolicyName != nil {
			policyResults.PolicyName = policyInfo.Response.PolicyName
		}
		if policyInfo.Response.AddTime != nil {
			policyResults.AddTime = policyInfo.Response.AddTime
		}
		if policyInfo.Response.CreateMode != nil {
			policyResults.CreateMode = policyInfo.Response.CreateMode
		}
	}

	return
}

func (me *CamService) DescribeGroupPolicyAttachmentsByFilter(ctx context.Context, params map[string]interface{}) (policyResults []*cam.AttachPolicyInfo, errRet error) {
	logId := getLogId(ctx)
	groupId := params["group_id"].(string)
	groupIdInt, e := strconv.Atoi(groupId)
	if e != nil {
		errRet = e
		return
	}
	groupIdInt64 := uint64(groupIdInt)
	request := cam.NewListAttachedGroupPoliciesRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)
	policyResults = make([]*cam.AttachPolicyInfo, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		request.TargetGroupId = &groupIdInt64
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListAttachedGroupPolicies(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "ResourceNotFound") {
					return
				}
			}
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}

		for _, policy := range response.Response.List {
			if params["policy_id"] != nil {
				if *policy.PolicyId != params["policy_id"].(uint64) {
					continue
				}
			}
			if params["create_mode"] != nil {
				if int(*policy.CreateMode) != params["create_mode"].(int) {
					continue
				}
			}
			policyResults = append(policyResults, policy)
		}
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	return
}

func (me *CamService) DescribePolicyById(ctx context.Context, policyId string) (result *cam.GetPolicyResponse, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewGetPolicyRequest()
	policyIdInt, e := strconv.Atoi(policyId)
	if e != nil {
		errRet = e
		return
	}
	policyIdInt64 := uint64(policyIdInt)
	request.PolicyId = &policyIdInt64
	result, err := me.client.UseCamClient().GetPolicy(request)

	if err != nil {
		log.Printf("[CRITAL]%s read CAM policy failed, reason:%s\n", logId, err.Error())
		if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
			errCode := ee.GetCode()
			//check if read empty
			if strings.Contains(errCode, "ResourceNotFound") {
				return
			}
		}
		return nil, err
	} else {
		if result == nil || result.Response == nil || result.Response.PolicyName == nil {
			return
		}
	}

	return
}

func (me *CamService) DescribePoliciesByFilter(ctx context.Context, params map[string]interface{}) (policies []*cam.StrategyInfo, errRet error) {
	logId := getLogId(ctx)
	policyId := -1
	policyName := ""
	//notice this policy type is different from the policy attachment, this sdk returns int while the attachments returns string
	policyType := -1
	description := ""
	createMode := -1

	request := cam.NewListPoliciesRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)

	for k, v := range params {
		if k == "policy_id" {
			policyId = v.(int)
		}
		if k == "name" {
			policyName = v.(string)
		}
		if k == "type" {
			policyType = v.(int)
		}
		if k == "description" {
			description = v.(string)
		}
		if k == "create_mode" {
			createMode = v.(int)
		}
	}
	policies = make([]*cam.StrategyInfo, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListPolicies(request)
		if err != nil {
			log.Printf("[CRITAL]%s read CAM policy failed, reason:%s\n", logId, err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "ResourceNotFound") {
					return
				}
			}
			errRet = err
			return
		}
		for _, policy := range response.Response.List {
			if policyId != -1 {
				if int(*policy.PolicyId) != policyId {
					continue
				}
			}
			if policyName != "" {
				if *policy.PolicyName != policyName {
					continue
				}
			}
			if policyType != -1 {
				if int(*policy.Type) != policyType {
					continue
				}
			}
			if description != "" {
				if *policy.Description != description {
					continue
				}
			}
			if createMode != -1 {
				if int(*policy.CreateMode) != createMode {
					continue
				}
			}
			policies = append(policies, policy)
		}
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	return
}

func (me *CamService) DescribeGroupsByFilter(ctx context.Context, params map[string]interface{}) (groups []*cam.GroupInfo, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewListGroupsRequest()
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)
	groups = make([]*cam.GroupInfo, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListGroups(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "ResourceNotFound") {
					return
				}
			}
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.GroupInfo) < 1 {
			break
		}
		for _, group := range response.Response.GroupInfo {
			if params["group_id"] != nil {
				if int(*group.GroupId) != params["group_id"].(int) {
					continue
				}
			}
			if params["name"] != nil {
				if *group.GroupName != params["name"].(string) {
					continue
				}
			}
			if params["remark"] != nil {
				if group.Remark == nil || (group.Remark != nil && *group.Remark != params["remark"].(string)) {
					continue
				}
				log.Printf("in")
			}
			groups = append(groups, group)
		}
		if len(response.Response.GroupInfo) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	return
}

func (me *CamService) DescribeGroupMembershipById(ctx context.Context, groupId string) (members []*string, errRet error) {
	logId := getLogId(ctx)
	groupIdInt, e := strconv.Atoi(groupId)
	if e != nil {
		errRet = e
		return
	}
	groupIdInt64 := int64(groupIdInt)
	pageStart := int64(1)
	rp := int64(PAGE_ITEM)
	members = make([]*string, 0)
	request := cam.NewListUsersForGroupRequest()
	request.GroupId = &groupIdInt64
	for {
		request.Page = &pageStart
		request.Rp = &rp
		response, err := me.client.UseCamClient().ListUsersForGroup(request)
		if err != nil {
			log.Printf("[CRITAL]%s read CAM group membership failed, reason:%s\n", logId, err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "ResourceNotFound") {
					return
				}
			}
			errRet = err
			return
		}

		if response == nil || len(response.Response.UserInfo) < 1 {
			break
		}
		for _, member := range response.Response.UserInfo {

			members = append(members, member.Name)
		}
		if len(response.Response.UserInfo) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	return
}

func (me *CamService) PolicyDocumentForceCheck(document string) error {

	//Policy syntax allows multi formats, but returns with only one format. In this case, the user's input may be different from the output value. To avoid this, terraform must make sure the syntax of the input policy document consists with the syntax of final returned output
	var documentJson CamDocument
	err := json.Unmarshal([]byte(document), &documentJson)
	if err != nil {
		return err
	}
	for _, state := range documentJson.Statement {
		//multi value case in elemant `resource`, `action`: input:""/[""], output: [""]
		if state.Resource != nil {
			if reflect.TypeOf(state.Resource) == reflect.TypeOf("string") {
				return fmt.Errorf("The format of `resource` in policy document is invalid, its type must be array")
			}
		}

		if state.Action != nil {
			if reflect.TypeOf(state.Action) == reflect.TypeOf("string") {
				return fmt.Errorf("The format of `resource` in policy document is invalid, its type must be array")
			}

		}
		//multi value case in elemant `principal.qcs`:input :root/[uin of the user], output:[uin of the user]
		for _, qcs := range state.Principal.Qcs {
			if strings.Contains(qcs, "root") {
				return fmt.Errorf("`root` format is not supported, please replace it with uin")
			}
		}
	}
	return nil
}

func (me *CamService) DescribeCamServiceLinkedRole(ctx context.Context, roleId string) (serviceLinkedRole *cam.RoleInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cam.NewGetRoleRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()
	request.RoleId = &roleId

	response, err := me.client.UseCamClient().GetRole(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil {
		serviceLinkedRole = response.Response.RoleInfo
	}

	return
}

// User related methods

func (me *CamService) DescribeUserNameByUin(ctx context.Context, uin int64) (userName string, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewGetUserListByUinListRequest()
	request.UinList = []*int64{&uin}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().GetUserListByUinList(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return "", err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil {
		return "", nil
	}
	for _, user := range response.Response.UserList {
		if user.UserUin != nil && *user.UserUin == uint64(uin) && user.UserName != nil {
			return *user.UserName, nil
		}
	}

	return "", nil
}

func (me *CamService) DescribeUserById(ctx context.Context, userId string) (user *cam.SubAccountFilter, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewListSubAccountsRequest()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().ListSubAccounts(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
			errCode := ee.GetCode()
			if strings.Contains(errCode, "ResourceNotFound") {
				return nil, nil
			}
		}
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil || len(response.Response.UserInfo) < 1 {
		return nil, nil
	}

	for _, subAccount := range response.Response.UserInfo {
		if subAccount.Name != nil && *subAccount.Name == userId {
			return subAccount, nil
		}
	}

	// User not found
	return nil, nil
}

func (me *CamService) DescribePasswordRules(ctx context.Context) (rules *cam.PasswordRules, blacklist *string, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewGetPasswordRulesRequest()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().GetPasswordRules(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil {
		rules = response.Response.Rules
		blacklist = response.Response.BlackList
	}
	return
}

func (me *CamService) UpdatePasswordRules(ctx context.Context, rules *cam.PasswordRules, blacklist *string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewUpdatePasswordRulesRequest()
	request.Rules = rules
	request.BlackList = blacklist

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().UpdatePasswordRules(request)
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

func (me *CamService) DescribeLoginRules(ctx context.Context) (sessionDuration *int64, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewGetLoginRulesRequest()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().GetLoginRules(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil {
		sessionDuration = response.Response.SessionDuration
	}
	return
}

func (me *CamService) SetLoginRules(ctx context.Context, sessionDuration *int64) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewSetLoginRulesRequest()
	request.SessionDuration = sessionDuration

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().SetLoginRules(request)
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

func (me *CamService) CreateApiKey(ctx context.Context, apiUin uint64, customSecretId, customSecretKey *string) (key *cam.ApiKey, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewCreateApiKeyRequest()
	request.ApiUin = &apiUin
	if customSecretId != nil {
		request.CustomSecretId = customSecretId
	}
	if customSecretKey != nil {
		request.CustomSecretKey = customSecretKey
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().CreateApiKey(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil && len(response.Response.IdKeys) > 0 {
		key = response.Response.IdKeys[0]
	}
	return
}

func (me *CamService) QueryApiKey(ctx context.Context, targetUin uint64) (keys []*cam.ApiKey, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewQueryApiKeyRequest()
	request.TargetUin = &targetUin

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().QueryApiKey(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil {
		keys = response.Response.IdKeys
	}
	return
}

func (me *CamService) EnableApiKey(ctx context.Context, apiUin uint64, secretId string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewEnableApiKeyRequest()
	request.ApiUin = &apiUin
	request.ApiSecretId = &secretId

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCamClient().EnableApiKey(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	return
}

func (me *CamService) DisableApiKey(ctx context.Context, apiUin uint64, secretId string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewDisableApiKeyRequest()
	request.ApiUin = &apiUin
	request.ApiSecretId = &secretId

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCamClient().DisableApiKey(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	return
}

func (me *CamService) DeleteApiKey(ctx context.Context, apiUin uint64, secretId string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewDeleteApiKeyRequest()
	request.ApiUin = &apiUin
	request.ApiSecretId = &secretId

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCamClient().DeleteApiKey(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	return
}

func (me *CamService) DescribeUsersByFilter(ctx context.Context, params map[string]interface{}) (users []*cam.SubAccountFilter, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewListSubAccountsRequest()

	users = make([]*cam.SubAccountFilter, 0)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().ListSubAccounts(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
			errCode := ee.GetCode()
			if strings.Contains(errCode, "ResourceNotFound") {
				return
			}
		}
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil || len(response.Response.UserInfo) < 1 {
		return
	}

	// Filter results based on params
	for _, user := range response.Response.UserInfo {
		// Filter by name
		if params["name"] != nil {
			if user.Name == nil || *user.Name != params["name"].(string) {
				continue
			}
		}

		// Filter by remark
		if params["remark"] != nil {
			if user.Remark == nil || *user.Remark != params["remark"].(string) {
				continue
			}
		}

		// Filter by phone_num
		if params["phone_num"] != nil {
			if user.PhoneNum == nil || *user.PhoneNum != params["phone_num"].(string) {
				continue
			}
		}

		// Filter by country_code
		if params["country_code"] != nil {
			if user.CountryCode == nil || *user.CountryCode != params["country_code"].(string) {
				continue
			}
		}

		// Filter by email
		if params["email"] != nil {
			if user.Email == nil || *user.Email != params["email"].(string) {
				continue
			}
		}

		// Filter by console_login
		if params["console_login"] != nil {
			if user.ConsoleLogin == nil || int(*user.ConsoleLogin) != params["console_login"].(int) {
				continue
			}
		}

		// Filter by uin
		if params["uin"] != nil {
			if user.Uin == nil || int(*user.Uin) != params["uin"].(int) {
				continue
			}
		}

		// Filter by uid
		if params["uid"] != nil {
			if user.Uid == nil || int(*user.Uid) != params["uid"].(int) {
				continue
			}
		}

		users = append(users, user)
	}

	return
}

// Group related methods

func (me *CamService) DescribeGroupById(ctx context.Context, groupId string) (group *cam.GroupInfo, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewListGroupsRequest()

	groupIdInt, e := strconv.ParseUint(groupId, 10, 64)
	if e != nil {
		errRet = fmt.Errorf("group_id should be a valid number string")
		return
	}

	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)

	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListGroups(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				if strings.Contains(errCode, "ResourceNotFound") {
					return nil, nil
				}
			}
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || len(response.Response.GroupInfo) < 1 {
			break
		}

		for _, g := range response.Response.GroupInfo {
			if g.GroupId != nil && *g.GroupId == groupIdInt {
				return g, nil
			}
		}

		if len(response.Response.GroupInfo) < PAGE_ITEM {
			break
		}
		pageStart++
	}

	return nil, nil
}

func (me *CamService) DeleteRolePolicyAttachmentByName(ctx context.Context, roleName, policyName string) error {
	logId := getLogId(ctx)
	request := cam.NewDetachRolePoliciesRequest()
	request.RoleName = &roleName
	request.PolicyName = []*string{&policyName}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().DetachRolePolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CamService) listEntitiesForPolicy(ctx context.Context, policyId uint64, entityFilter string, matchIdentifier string) (entity *cam.AttachEntityOfPolicy, errRet error) {
	logId := getLogId(ctx)
	req := cam.NewListEntitiesForPolicyRequest()
	req.PolicyId = &policyId
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)
	req.Rp = &rp
	req.EntityFilter = &entityFilter

	for {
		req.Page = &pageStart
		ratelimit.Check(req.GetAction())
		resp, err := me.client.UseCamClient().ListEntitiesForPolicy(req)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, req.GetAction(), req.ToJsonString(), err.Error())
			if ee, ok := err.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				if strings.Contains(errCode, "ResourceNotFound") {
					return
				}
			}
			errRet = err
			return
		}
		if resp == nil || resp.Response == nil || len(resp.Response.List) == 0 {
			break
		}
		for _, item := range resp.Response.List {
			// Match by both Id and Name fields to handle different entity types
			// For User: Name is username, Id is Uid
			// For Group: Name is group name, Id is group ID
			// For Role: Name is role name, Id is role ID
			if matchIdentifier == "" {
				entity = item
				return
			}

			// Try to match by Name first (for User entities where we use username)
			if item.Name != nil && *item.Name == matchIdentifier {
				entity = item
				return
			}

			// Try to match by Id (for Group and Role entities where we use ID)
			if item.Id != nil && *item.Id == matchIdentifier {
				entity = item
				return
			}
		}
		if len(resp.Response.List) < int(PAGE_ITEM) {
			break
		}
		pageStart++
	}
	return
}

func (me *CamService) DescribeUserPolicyAttachmentById(ctx context.Context, userPolicyAttachmentId string) (policyResults *cam.AttachPolicyInfo, errRet error) {
	logId := getLogId(ctx)
	userId, policyId, e := me.decodeCamPolicyAttachmentId(userPolicyAttachmentId)
	if e != nil {
		return nil, e
	}

	// Use ListEntitiesForPolicy to check if attachment exists
	entity, err := me.listEntitiesForPolicy(ctx, policyId, "User", userId)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return
	}

	// Build AttachPolicyInfo from entity
	policyResults = &cam.AttachPolicyInfo{
		PolicyId: &policyId,
	}

	// If entity has Name field, use it to populate PolicyName
	if entity.Name != nil {
		policyResults.PolicyName = entity.Name
	}

	// Get additional info from policy if needed
	policyInfo, err := me.DescribePolicyById(ctx, fmt.Sprintf("%d", policyId))
	if err != nil {
		log.Printf("[WARN]%s failed to get policy info for policy %d, reason: %s\n", logId, policyId, err.Error())
	} else if policyInfo != nil && policyInfo.Response != nil {
		if policyInfo.Response.PolicyName != nil {
			policyResults.PolicyName = policyInfo.Response.PolicyName
		}
		if policyInfo.Response.AddTime != nil {
			policyResults.AddTime = policyInfo.Response.AddTime
		}
		if policyInfo.Response.CreateMode != nil {
			policyResults.CreateMode = policyInfo.Response.CreateMode
		}
	}

	return
}

func (me *CamService) DescribeUserPolicyAttachmentsByFilter(ctx context.Context, params map[string]interface{}) (policyResults []*cam.AttachPolicyInfo, errRet error) {
	logId := getLogId(ctx)
	userId := params["user_id"].(string)
	policyResults = make([]*cam.AttachPolicyInfo, 0)

	// Get user info to get UIN
	user, err := me.DescribeUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Uin == nil {
		return
	}
	targetUin := user.Uin

	// Use ListAttachedUserAllPolicies to list all policies attached to the user
	request := cam.NewListAttachedUserAllPoliciesRequest()
	request.TargetUin = targetUin
	// AttachType: 0=返回直接关联和随组关联策略，1=只返回直接关联策略，2=只返回随组关联策略
	attachType := uint64(0) // Get all policies (direct and group-attached)
	request.AttachType = &attachType
	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM)

	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListAttachedUserAllPolicies(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return nil, err
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || len(response.Response.PolicyList) == 0 {
			break
		}

		// Convert AttachedUserPolicy to AttachPolicyInfo and apply filters
		for _, userPolicy := range response.Response.PolicyList {
			// Filter by policy_id if provided
			if params["policy_id"] != nil {
				if userPolicy.PolicyId == nil || *userPolicy.PolicyId != params["policy_id"].(uint64) {
					continue
				}
			}

			// Filter by create_mode if provided
			if params["create_mode"] != nil {
				if userPolicy.CreateMode == nil || int(*userPolicy.CreateMode) != params["create_mode"].(int) {
					continue
				}
			}

			policyInfo := &cam.AttachPolicyInfo{
				PolicyId:   userPolicy.PolicyId,
				PolicyName: userPolicy.PolicyName,
				AddTime:    userPolicy.AddTime,
				CreateMode: userPolicy.CreateMode,
			}
			policyResults = append(policyResults, policyInfo)
		}

		if len(response.Response.PolicyList) < PAGE_ITEM {
			break
		}
		pageStart++
	}

	return
}

func (me *CamService) AddUserPolicyAttachment(ctx context.Context, userId string, policyId string) error {
	logId := getLogId(ctx)

	user, err := me.DescribeUserById(ctx, userId)
	if err != nil {
		return err
	}
	if user == nil || user.Uin == nil {
		return nil
	}
	uin := *user.Uin
	policyIdInt, e := strconv.Atoi(policyId)
	if e != nil {
		return e
	}
	policyIdInt64 := uint64(policyIdInt)

	req := cam.NewAttachUsersPolicyRequest()
	req.PolicyId = &policyIdInt64
	req.TargetUin = []*uint64{&uin}
	ratelimit.Check(req.GetAction())
	resp, err := me.client.UseCamClient().AttachUsersPolicy(req)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, req.GetAction(), req.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, req.GetAction(), req.ToJsonString(), resp.ToJsonString())
	return nil
}

func (me *CamService) DeleteUserPolicyAttachmentById(ctx context.Context, userPolicyAttachmentId string) error {
	logId := getLogId(ctx)
	userId, policyId, e := me.decodeCamPolicyAttachmentId(userPolicyAttachmentId)
	if e != nil {
		return e
	}
	user, err := me.DescribeUserById(ctx, userId)
	if err != nil {
		return err
	}
	if user == nil || user.Uin == nil {
		return nil
	}
	uin := *user.Uin

	req := cam.NewDetachUsersPolicyRequest()
	req.PolicyId = &policyId
	req.TargetUin = []*uint64{&uin}
	ratelimit.Check(req.GetAction())
	resp, err := me.client.UseCamClient().DetachUsersPolicy(req)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, req.GetAction(), req.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, req.GetAction(), req.ToJsonString(), resp.ToJsonString())
	return nil
}

func (me *CamService) AddGroupPolicyAttachment(ctx context.Context, groupId string, policyId string) error {
	logId := getLogId(ctx)

	groupIdInt, e := strconv.Atoi(groupId)
	if e != nil {
		return e
	}
	groupIdInt64 := uint64(groupIdInt)
	policyIdInt, ee := strconv.Atoi(policyId)
	if ee != nil {
		return ee
	}
	policyIdInt64 := uint64(policyIdInt)

	request := cam.NewAttachGroupPoliciesRequest()
	request.GroupId = &groupIdInt64
	request.PolicyId = []*uint64{&policyIdInt64}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().AttachGroupPolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CamService) DeleteGroupPolicyAttachmentById(ctx context.Context, groupPolicyAttachmentId string) error {
	logId := getLogId(ctx)
	groupId, policyId, e := me.decodeCamPolicyAttachmentId(groupPolicyAttachmentId)
	if e != nil {
		return e
	}
	groupIdInt, ee := strconv.Atoi(groupId)
	if ee != nil {
		return ee
	}
	groupIdInt64 := uint64(groupIdInt)

	request := cam.NewDetachGroupPoliciesRequest()
	request.GroupId = &groupIdInt64
	request.PolicyId = []*uint64{&policyId}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().DetachGroupPolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

//func (me *CamService) DescribeGroupPolicyAttachmentById(ctx context.Context, groupPolicyAttachmentId string) (policyResults *cam.AttachPolicyInfo, errRet error) {
//	groupId, policyId, e := me.decodeCamPolicyAttachmentId(groupPolicyAttachmentId)
//	if e != nil {
//		errRet = e
//		return
//	}
//	entity, err := me.listEntitiesForPolicy(ctx, policyId, "Group", groupId)
//	if err != nil {
//		return nil, err
//	}
//	if entity == nil {
//		return
//	}
//	policyResults = &cam.AttachPolicyInfo{PolicyId: &policyId}
//	return
//}

//func (me *CamService) DescribeGroupPolicyAttachmentsByFilter(ctx context.Context, params map[string]interface{}) (policyResults []*cam.AttachPolicyInfo, errRet error) {
//	groupId := params["group_id"].(string)
//	policyResults = make([]*cam.AttachPolicyInfo, 0)
//	if params["policy_id"] != nil {
//		policyId := params["policy_id"].(uint64)
//		entity, err := me.listEntitiesForPolicy(ctx, policyId, "Group", groupId)
//		if err != nil {
//			return nil, err
//		}
//		if entity != nil {
//			policyResults = append(policyResults, &cam.AttachPolicyInfo{PolicyId: &policyId})
//		}
//	}
//	return
//}

func (me *CamService) DescribeCamListEntitiesForPolicyByFilter(ctx context.Context, param map[string]interface{}) (ListEntitiesForPolicy []*cam.AttachEntityOfPolicy, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cam.NewListEntitiesForPolicyRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "PolicyId" {
			request.PolicyId = v.(*uint64)
		}
		if k == "Rp" {
			request.Rp = v.(*uint64)
		}
		if k == "EntityFilter" {
			val := v.(string)
			request.EntityFilter = &val
		}
	}

	ratelimit.Check(request.GetAction())

	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM) //to save in extension
	result := make([]*cam.AttachEntityOfPolicy, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListEntitiesForPolicy(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || len(response.Response.List) < 1 {
			break
		}
		result = append(result, response.Response.List...)
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}
	ListEntitiesForPolicy = result
	return
}

func (me *CamService) DescribeCamRoleDetailByFilter(ctx context.Context, param map[string]interface{}) (ret *cam.RoleInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cam.NewGetRoleRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "RoleId" {
			val := v.(string)
			request.RoleId = &val
		}
		if k == "RoleName" {
			val := v.(string)
			request.RoleName = &val
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCamClient().GetRole(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil {
		return
	}

	ret = response.Response.RoleInfo
	return
}

func (me *CamService) DescribeCamSubAccountsByFilter(ctx context.Context, param map[string]interface{}) (ret []*cam.SubAccountFilter, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cam.NewListSubAccountsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCamClient().ListSubAccounts(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil || len(response.Response.UserInfo) < 1 {
		return
	}

	ret = response.Response.UserInfo

	// Manual filter by UIN if FilterSubAccountUin is provided in params
	if v, ok := param["FilterSubAccountUin"]; ok {
		filterUins := v.([]*uint64)
		if len(filterUins) > 0 {
			filtered := make([]*cam.SubAccountFilter, 0)
			uinMap := make(map[uint64]bool)
			for _, uin := range filterUins {
				if uin != nil {
					uinMap[*uin] = true
				}
			}
			for _, user := range ret {
				if user.Uin != nil && uinMap[*user.Uin] {
					filtered = append(filtered, user)
				}
			}
			ret = filtered
		}
	}

	return
}

func (me *CamService) DescribeCamListAttachedUserPolicyByFilter(ctx context.Context, param map[string]interface{}) (ret []*cam.AttachedUserPolicy, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = cam.NewListAttachedUserAllPoliciesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "TargetUin" {
			request.TargetUin = v.(*uint64)
		}
		if k == "AttachType" {
			request.AttachType = v.(*uint64)
		}
		if k == "StrategyType" {
			request.StrategyType = v.(*uint64)
		}
		if k == "Keyword" {
			// Internal SDK uses Filter instead of Keyword
			val := v.(string)
			request.Filter = &val
		}
	}

	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM) //to save in extension
	result := make([]*cam.AttachedUserPolicy, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListAttachedUserAllPolicies(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || len(response.Response.PolicyList) < 1 {
			break
		}
		result = append(result, response.Response.PolicyList...)
		if len(response.Response.PolicyList) < PAGE_ITEM {
			break
		}

		pageStart += 1
	}
	ret = result
	return
}

func (me *CamService) DescribeCamMfaFlagById(ctx context.Context, id uint64) (loginFlag *account.SafeAuthFlag, actionFlag *account.SafeAuthFlag, errRet error) {
	logId := getLogId(ctx)

	request := account.NewGetSafeAuthConfigRequest()
	response := account.NewGetSafeAuthConfigResponse()
	request.UserUin = &id
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseAccountClient().GetSafeAuthConfig(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil || response.Response.Flag == nil {
		return
	}

	if response.Response.Flag.LoginFlag != nil {
		loginFlag = response.Response.Flag.LoginFlag
	}

	if response.Response.Flag.ActionFlag != nil {
		actionFlag = response.Response.Flag.ActionFlag
	}

	return
}

func (me *CamService) DeleteCamServiceLinkedRoleByName(ctx context.Context, roleName string) (errRet error) {
	logId := getLogId(ctx)

	request := cam.NewDeleteServiceLinkedRoleRequest()
	request.RoleName = &roleName

	errRet = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseCamClient().DeleteServiceLinkedRole(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})

	if errRet != nil {
		log.Printf("[CRITAL]%s delete CAM serviceLinkedRole failed, reason:%s\n", logId, errRet.Error())
	}

	return
}

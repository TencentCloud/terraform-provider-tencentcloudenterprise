package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type OrganizationService struct {
	client *connectivity.TencentCloudClient
}

func (me *OrganizationService) DescribeOrganizationOrgNode(ctx context.Context, nodeId string) (orgNode *organization.OrgNode, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = organization.NewDescribeOrganizationNodesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	var offset int64 = 0
	var pageSize int64 = 50
	instances := make([]*organization.OrgNode, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().DescribeOrganizationNodes(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		instances = append(instances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	if len(instances) < 1 {
		return
	}

	for _, instance := range instances {
		if helper.Int64ToStr(*instance.NodeId) == nodeId {
			orgNode = instance
		}
	}

	return
}

func (me *OrganizationService) DeleteOrganizationOrgNodeById(ctx context.Context, nodeId string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationNodesRequest()

	request.NodeId = []*int64{helper.Int64(helper.StrToInt64(nodeId))}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().DeleteOrganizationNodes(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMember(ctx context.Context, uin string) (orgMember *organization.OrgMember, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = organization.NewDescribeOrganizationMembersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var offset uint64 = 0
	var pageSize uint64 = 50
	instances := make([]*organization.OrgMember, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMembers(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		instances = append(instances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	if len(instances) < 1 {
		return
	}

	for _, instance := range instances {
		if helper.UInt64ToStr(*instance.MemberUin) == uin {
			orgMember = instance
		}
	}

	return

}

func (me *OrganizationService) DeleteOrganizationOrgMemberById(ctx context.Context, uin string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationMembersRequest()

	request.MemberUin = []*int64{helper.Int64(helper.StrToInt64(uin))}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().DeleteOrganizationMembers(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationPolicySubAccountAttachment(ctx context.Context, policyId, memberUin string) (policySubAccountAttachment *organization.OrgMemberAuthAccount, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = organization.NewDescribeOrganizationMemberAuthAccountsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()
	request.PolicyId = helper.StrToInt64Point(policyId)
	request.MemberUin = helper.StrToUint64Point(memberUin)
	request.Limit = helper.IntInt64(50)
	request.Offset = helper.IntInt64(0)

	response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberAuthAccounts(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	if len(response.Response.Items) < 1 {
		return
	}
	policySubAccountAttachment = response.Response.Items[0]
	return
}

func (me *OrganizationService) DeleteOrganizationPolicySubAccountAttachmentById(ctx context.Context, policyId, memberUin, orgSubAccountUin string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewCancelOrganizationMemberAuthAccountRequest()

	request.PolicyId = helper.StrToInt64Point(policyId)
	request.MemberUin = helper.StrToUint64Point(memberUin)
	request.OrgSubAccountUin = helper.StrToUint64Point(orgSubAccountUin)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().CancelOrganizationMemberAuthAccount(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrganizationById(ctx context.Context) (
	result *organization.DescribeOrganizationResponseParams, errRet error) {

	logId := getLogId(ctx)

	request := organization.NewDescribeOrganizationRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeOrganization(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	result = response.Response

	return
}

func (me *OrganizationService) UpdateOrganizationRootNodeName(ctx context.Context, orgId uint64, name string) (errRet error) {
	logId := getLogId(ctx)
	organizationResponseParams, err := me.DescribeOrganizationOrganizationById(ctx)
	if err != nil {
		return err
	}
	if organizationResponseParams == nil {
		return fmt.Errorf("organization is nil")
	}
	rootNodeId := organizationResponseParams.RootNodeId

	request := organization.NewUpdateOrganizationNodeRequest()
	request.NodeId = rootNodeId
	request.Name = &name

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := me.client.UseOrganizationClient().UpdateOrganizationNode(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update organization orgNode name failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

func (me *OrganizationService) DeleteOrganizationOrganizationById(ctx context.Context) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganization(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

// DescribeOrganizationOrgIdentityById retrieves an organization identity by its ID.
func (me *OrganizationService) DescribeOrganizationOrgIdentityById(ctx context.Context, identityId string) (orgIdentity *organization.OrgIdentity, errRet error) {
	logId := getLogId(ctx)

	request := organization.NewListOrganizationIdentityRequest()
	request.IdentityId = helper.StrToUint64Point(identityId)

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
	var tmp []*organization.OrgIdentity
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().ListOrganizationIdentity(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		tmp = append(tmp, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}
	}
	for _, item := range tmp {
		if *item.IdentityId == helper.StrToInt64(identityId) {
			orgIdentity = item
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgIdentityById(ctx context.Context, identityId string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationIdentityRequest()
	request.IdentityId = helper.StrToUint64Point(identityId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganizationIdentity(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyById(ctx context.Context, policyId, policyType string) (
	OrgManagePolicy *organization.DescribePolicyResponseParams, errRet error) {
	logId := getLogId(ctx)

	request := organization.NewListPoliciesRequest()
	request.PolicyType = helper.String(policyType)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM) //to save in extension
	result := make([]*organization.ListPolicyNode, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().ListPolicies(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}
		result = append(result, response.Response.List...)
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}

	for _, item := range result {
		if helper.UInt64ToStr(*item.PolicyId) == policyId {
			requestDescribe := organization.NewDescribePolicyRequest()
			requestDescribe.PolicyId = item.PolicyId
			requestDescribe.PolicyType = helper.String(policyType)
			responseDescribe, err := me.client.UseOrganizationClient().DescribePolicy(requestDescribe)
			if err != nil {
				errRet = err
				return
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), responseDescribe.ToJsonString())

			if responseDescribe == nil || responseDescribe.Response == nil {
				break
			}
			OrgManagePolicy = responseDescribe.Response
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyById(ctx context.Context, policyId, policyType string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeletePolicyRequest()
	request.PolicyId = helper.StrToUint64Point(policyId)
	request.Type = helper.String(policyType)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeletePolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyConfigById(ctx context.Context, organizationId string, policyType string) (OrgManagePolicyConfig *organization.DescribePolicyConfigResponseParams, errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDescribePolicyConfigRequest()
	request.OrganizationId = helper.StrToUint64Point(organizationId)
	request.Type = helper.IntUint64(ServiceControlPolicyCode)

	if policyType == TagPolicyType {
		request.Type = helper.IntUint64(TagPolicyCode)
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribePolicyConfig(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if *response.Response.Status == 1 {
		OrgManagePolicyConfig = response.Response
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyConfigById(ctx context.Context, organizationId string, policyType string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDisablePolicyTypeRequest()
	request.OrganizationId = helper.StrToUint64Point(organizationId)
	request.PolicyType = &policyType

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DisablePolicyType(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyTargetById(ctx context.Context, policyType string, policyId string, targetType string, targetId string) (OrgManagePolicyTarget *organization.ListTargetsForPolicyNode, errRet error) {
	logId := getLogId(ctx)

	request := organization.NewListTargetsForPolicyRequest()
	request.PolicyType = &policyType
	request.PolicyId = helper.StrToUint64Point(policyId)
	switch targetType {
	case TargetTypeNode:
		request.TargetType = helper.String(DescribeTargetTypeNode)
	case TargetTypeMember:
		request.TargetType = helper.String(DescribeTargetTypeMember)
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().ListTargetsForPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	for _, item := range response.Response.List {
		if item.Uin != nil && helper.UInt64ToStr(*item.Uin) == targetId {
			OrgManagePolicyTarget = item
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyTargetById(ctx context.Context, policyType string, policyId string, targetType string, targetId string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDetachPolicyRequest()
	request.Type = &policyType
	request.PolicyId = helper.StrToUint64Point(policyId)
	request.TargetType = &targetType
	request.TargetId = helper.StrToUint64Point(targetId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DetachPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMemberAuthIdentityById(ctx context.Context, memberUin int64) (identityIds []int64, errRet error) {
	logId := getLogId(ctx)
	request := organization.NewDescribeOrganizationMemberAuthIdentitiesRequest()
	request.MemberUin = &memberUin

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	var (
		offset int64 = 0
		limit  int64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberAuthIdentities(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if len(response.Response.Items) < 1 {
			return
		}

		for _, v := range response.Response.Items {
			if v.MemberUin != nil && *v.MemberUin == memberUin {
				if *v.IdentityId == 1 {
					continue
				}
				identityIds = append(identityIds, *v.IdentityId)
			}
		}
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *OrganizationService) DeleteOrganizationOrgMemberAuthIdentityById(ctx context.Context, memberUin string, identityId []string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationMemberAuthIdentityRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	for _, id := range identityId {
		log.Printf("[DEBUG]%s api[%s] delete identity, uin [%s], identityId [%s]\n", logId, request.GetAction(), memberUin, id)

		request.MemberUin = helper.StrToUint64Point(memberUin)
		request.IdentityId = helper.StrToUint64Point(id)

		response, err := me.client.UseOrganizationClient().DeleteOrganizationMemberAuthIdentity(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] delete identity success, request memberUin [%s], response body [%s]\n", logId, request.GetAction(), memberUin, response.ToJsonString())
	}

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMemberEmailById(ctx context.Context, memberUin int64, bindId uint64) (orgMemberEmail *organization.DescribeOrganizationMemberEmailBindResponseParams, errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDescribeOrganizationMemberEmailBindRequest()
	request.MemberUin = &memberUin

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberEmailBind(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil {
		return
	}
	if *response.Response.BindId != bindId {
		return
	}
	orgMemberEmail = response.Response
	return
}

func (me *OrganizationService) DeleteOrganizationOrgMemberPolicyAttachmentById(ctx context.Context, policyId string) (errRet error) {
	logId := getLogId(ctx)

	request := organization.NewDeleteOrganizationMembersPolicyRequest()
	request.PolicyId = helper.StrToUint64Point(policyId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganizationMembersPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationMembersByFilter(ctx context.Context, param map[string]interface{}) (members []*organization.OrgMember, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = organization.NewDescribeOrganizationMembersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Lang" {
			request.Lang = v.(*string)
		}
		if k == "SearchKey" {
			request.SearchKey = v.(*string)
		}
		if k == "AuthName" {
			request.AuthName = v.(*string)
		}
		if k == "Product" {
			request.Product = v.(*string)
		}
	}

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMembers(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || len(response.Response.Items) < 1 {
			break
		}
		members = append(members, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

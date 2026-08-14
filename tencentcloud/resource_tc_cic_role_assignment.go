/*
Provides a resource to create a organization cic_role_assignment

# Example Usage

```hcl

	resource "tencentcloudenterprise_cic_role_assignment" "cic_role_assignment" {
	  zone_id = "z-xxxxxx"
	  principal_id = "u-xxxxxx"
	  principal_type = "User"
	  target_uin = "xxxxxx"
	  target_type = "MemberUin"
	  role_configuration_id = "rc-xxxxxx"
	}

```

# Import

organization cic_role_assignment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_role_assignment.cic_role_assignment {zoneId}#{roleConfigurationId}#{targetType}#{targetUinString}#{principalType}#{principalId}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
)

// Serialize Create/Delete that share the same (zone, role configuration, target)
// so Terraform for_each parallelism cannot race List lag or last-deprovision.
var cicRoleAssignmentTargetLocks sync.Map // map[string]*sync.Mutex

func cicRoleAssignmentTargetUnlock(zoneId, roleConfigurationId, targetType string, targetUin int64) func() {
	key := strings.Join([]string{zoneId, roleConfigurationId, targetType, strconv.FormatInt(targetUin, 10)}, FILED_SP)
	v, _ := cicRoleAssignmentTargetLocks.LoadOrStore(key, &sync.Mutex{})
	mu := v.(*sync.Mutex)
	mu.Lock()
	return mu.Unlock
}

// isCicDismantleAlreadyGone reports whether a dismantle API/task failure means the
// deployment is already gone (typical under concurrent last-assignment destroy).
func isCicDismantleAlreadyGone(errMsg string) bool {
	return strings.Contains(errMsg, "RoleConfigurationProvisioningNotFound") ||
		strings.Contains(errMsg, "RecordNotFound") ||
		strings.Contains(errMsg, "DBOperationError")
}

// cicRoleAssignmentID is also the provider's import/state ID. Build it before
// calling CreateRoleAssignment so a retry can first reconcile an assignment
// that was created remotely even if its asynchronous task was reported as
// failed or Terraform was interrupted before state was written.
func cicRoleAssignmentID(zoneId, roleConfigurationId, targetType string, targetUin int64, principalType, principalId string) string {
	return strings.Join([]string{
		zoneId,
		roleConfigurationId,
		targetType,
		strconv.FormatInt(targetUin, 10),
		principalType,
		principalId,
	}, FILED_SP)
}

func cicRoleAssignmentFailureReason(reason *string) string {
	if reason == nil {
		return ""
	}
	return strings.TrimSpace(*reason)
}

// CIC has returned Failed with a nil/empty FailureReason while a newly-created
// organization member was still propagating to Identity Center. Such a task
// is safe to reconcile and retry; an explicit reason is a real API failure and
// must be surfaced to the user immediately.
func cicRoleAssignmentTaskNeedsRetry(status, reason string) bool {
	return status == TASK_STATUS_FAILED && strings.TrimSpace(reason) == ""
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_role_assignment", CNDescription{
		TerraformTypeCN: "身份中心角色分配",
		DescriptionCN:   "提供身份中心角色分配资源，用于将角色配置分配给用户或用户组。",
		AttributesCN: map[string]string{
			"zone_id":               "空间ID",
			"principal_id":          "委托人ID",
			"principal_type":        "委托人类型",
			"target_uin":            "目标账号UIN",
			"target_type":           "目标类型",
			"role_configuration_id": "角色配置ID",
			"create_time":           "创建时间",
		},
	})
}

func resourceTencentCloudCicRoleAssignment() *schema.Resource {
	return &schema.Resource{
		Description: "Provide identity center role assignment resources for allocating role configurations to users or user groups.",
		Create:      resourceTencentCloudCicRoleAssignmentCreate,
		Read:        resourceTencentCloudCicRoleAssignmentRead,
		Delete:      resourceTencentCloudCicRoleAssignmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID.",
			},
			"principal_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Identity ID for the CAM user synchronization. Valid values:\nWhen the PrincipalType value is Group, it is the CIC user group ID (g-********).\nWhen the PrincipalType value is User, it is the CIC user ID (u-********).",
			},
			"principal_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Identity type for the CAM user synchronization. Valid values:\n\nUser: indicates that the identity for the CAM user synchronization is a CIC user.\nGroup: indicates that the identity for the CAM user synchronization is a CIC user group.",
			},
			"target_uin": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "UIN of the synchronized target account of the Tencent Cloud Organization.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of the synchronized target account of the Tencent Cloud Organization. ManagerUin: admin account; MemberUin: member account.",
			},
			"role_configuration_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Permission configuration ID.",
			},
			"deprovision_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
				ForceNew:    true,
				Description: "Whether to dismantle the role configuration deployment on the target account after removing this authorization. Valid values: DeprovisionForLastRoleAssignmentOnAccount (dismantle only when no other user/group authorization remains on the same role configuration and target account), None (default, only remove this authorization and keep deployment).",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time.",
			},
			"role_configuration_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role configuration name.",
			},
			"target_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target name.",
			},
			"principal_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Principal name.",
			},
		},
	}
}

func resourceTencentCloudCicRoleAssignmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_assignment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}
	var (
		zoneId              string
		roleConfigurationId string
		targetType          string
		targetUin           int64
		principalType       string
		principalId         string
	)
	request := cic.NewCreateRoleAssignmentRequest()

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
		request.ZoneId = helper.String(zoneId)
	}

	roleAssignmentInfo := cic.RoleAssignmentInfo{}
	if v, ok := d.GetOk("principal_id"); ok {
		principalId = v.(string)
		roleAssignmentInfo.PrincipalId = helper.String(principalId)
	}
	if v, ok := d.GetOk("principal_type"); ok {
		principalType = v.(string)
		roleAssignmentInfo.PrincipalType = helper.String(principalType)
	}
	if v, ok := d.GetOk("target_uin"); ok {
		targetUin = int64(v.(int))
		roleAssignmentInfo.TargetUin = helper.Int64(targetUin)
	}
	if v, ok := d.GetOk("target_type"); ok {
		targetType = v.(string)
		roleAssignmentInfo.TargetType = helper.String(targetType)
	}
	if v, ok := d.GetOk("role_configuration_id"); ok {
		roleConfigurationId = v.(string)
		roleAssignmentInfo.RoleConfigurationId = helper.String(roleConfigurationId)
	}
	request.RoleAssignmentInfo = []*cic.RoleAssignmentInfo{&roleAssignmentInfo}
	assignmentID := cicRoleAssignmentID(zoneId, roleConfigurationId, targetType, targetUin, principalType, principalId)

	// Hold until post-create Read finishes so concurrent creates on the same
	// target do not race ListRoleAssignments / provision tasks.
	unlock := cicRoleAssignmentTargetUnlock(zoneId, roleConfigurationId, targetType, targetUin)
	defer unlock()

	// The organization API can return a new member UIN before CIC can accept
	// an authorization for that member. Retry the complete asynchronous
	// operation, but reconcile first so a successful remote operation is never
	// submitted a second time.
	err := resource.Retry(2*writeRetryTimeout, func() *resource.RetryError {
		existing, e := service.DescribeCicRoleAssignmentById(context.WithValue(context.Background(), logIdKey, logId), assignmentID)
		if e != nil {
			return retryError(e)
		}
		if existing != nil && len(existing.RoleAssignments) > 0 {
			d.SetId(assignmentID)
			log.Printf("[INFO]%s CIC role assignment already exists, adopting remote assignment %s", logId, assignmentID)
			return nil
		}

		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateRoleAssignment(request)
		if e != nil {
			// A previous attempt may have completed remotely while returning an
			// ambiguous duplicate/authorization error. Reconcile that case.
			if strings.Contains(e.Error(), "AuthorizationExist") {
				existing, listErr := service.DescribeCicRoleAssignmentById(context.WithValue(context.Background(), logIdKey, logId), assignmentID)
				if listErr != nil {
					return retryError(listErr)
				}
				if existing != nil && len(existing.RoleAssignments) > 0 {
					d.SetId(assignmentID)
					return nil
				}
				return resource.RetryableError(fmt.Errorf("CIC reports the role assignment already exists but it is not visible yet; retrying"))
			}
			return retryError(e)
		}

		if result == nil || result.Response == nil || len(result.Response.Tasks) == 0 || result.Response.Tasks[0] == nil {
			return resource.NonRetryableError(fmt.Errorf("create role assignment succeeded but no task returned"))
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())

		task := result.Response.Tasks[0]
		taskReason := cicRoleAssignmentFailureReason(task.FailureReason)
		if task.Status != nil && *task.Status == TASK_STATUS_FAILED {
			if !cicRoleAssignmentTaskNeedsRetry(*task.Status, taskReason) {
				return resource.NonRetryableError(fmt.Errorf("create role assignment task failed, failure reason:%s", taskReason))
			}

			// The task can be marked Failed without a reason while the remote
			// assignment is still becoming visible. Check before retrying Create.
			existing, listErr := service.DescribeCicRoleAssignmentById(context.WithValue(context.Background(), logIdKey, logId), assignmentID)
			if listErr != nil {
				return retryError(listErr)
			}
			if existing != nil && len(existing.RoleAssignments) > 0 {
				d.SetId(assignmentID)
				return nil
			}
			return resource.RetryableError(fmt.Errorf("create role assignment task failed without a reason; retrying after CIC propagation delay"))
		}

		if task.TaskId == nil {
			return resource.NonRetryableError(fmt.Errorf("create role assignment task id is nil"))
		}
		conf := BuildStateChangeConf([]string{}, []string{TASK_STATUS_SUCCESS, TASK_STATUS_FAILED},
			2*readRetryTimeout, time.Second, service.AssignmentTaskStatusStateRefreshFunc(zoneId, *task.TaskId, []string{}))
		object, waitErr := conf.WaitForState()
		if waitErr != nil {
			return resource.RetryableError(waitErr)
		}
		taskStatus, ok := object.(*cic.TaskStatus)
		if !ok || taskStatus == nil || taskStatus.Status == nil {
			return resource.RetryableError(fmt.Errorf("create role assignment task status is nil; retrying"))
		}
		statusReason := cicRoleAssignmentFailureReason(taskStatus.FailureReason)
		if *taskStatus.Status == TASK_STATUS_FAILED {
			if !cicRoleAssignmentTaskNeedsRetry(*taskStatus.Status, statusReason) {
				return resource.NonRetryableError(fmt.Errorf("create role assignment task failed, failure reason:%s", statusReason))
			}

			existing, listErr := service.DescribeCicRoleAssignmentById(context.WithValue(context.Background(), logIdKey, logId), assignmentID)
			if listErr != nil {
				return retryError(listErr)
			}
			if existing != nil && len(existing.RoleAssignments) > 0 {
				d.SetId(assignmentID)
				return nil
			}
			return resource.RetryableError(fmt.Errorf("create role assignment task failed without a reason; retrying after CIC propagation delay"))
		}
		if *taskStatus.Status != TASK_STATUS_SUCCESS {
			return resource.RetryableError(fmt.Errorf("create role assignment task status is %s; retrying", *taskStatus.Status))
		}

		d.SetId(assignmentID)
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center role assignment failed, reason:%+v", logId, err)
		return err
	}

	if d.Id() == "" {
		return fmt.Errorf("create role assignment succeeded but no id returned; cannot persist resource state")
	}

	return resourceTencentCloudCicRoleAssignmentRead(d, meta)
}

func resourceTencentCloudCicRoleAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_assignment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 6 {
		return fmt.Errorf("roleAssignmentId is broken,%s", d.Id())
	}
	zoneId := idSplit[0]

	// ListRoleAssignments can lag briefly after CreateRoleAssignment task Success,
	// especially under concurrent creates on the same target. Only retry empty
	// results for newly created resources; otherwise treat empty as deleted.
	var roleAssignment *cic.RoleAssignments
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCicRoleAssignmentById(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}
		if result == nil || len(result.RoleAssignments) == 0 {
			if d.IsNewResource() {
				return resource.RetryableError(fmt.Errorf("cic role assignment %s not found after create, retrying", d.Id()))
			}
			return nil
		}
		roleAssignment = result.RoleAssignments[0]
		return nil
	})
	if err != nil {
		return err
	}
	if roleAssignment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_role_assignment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("zone_id", zoneId)
	if roleAssignment.RoleConfigurationId != nil {
		_ = d.Set("role_configuration_id", roleAssignment.RoleConfigurationId)
	}
	if roleAssignment.RoleConfigurationName != nil {
		_ = d.Set("role_configuration_name", roleAssignment.RoleConfigurationName)
	}
	if roleAssignment.TargetUin != nil {
		_ = d.Set("target_uin", roleAssignment.TargetUin)
	}
	if roleAssignment.TargetType != nil {
		_ = d.Set("target_type", roleAssignment.TargetType)
	}
	if roleAssignment.PrincipalId != nil {
		_ = d.Set("principal_id", roleAssignment.PrincipalId)
	}
	if roleAssignment.PrincipalType != nil {
		_ = d.Set("principal_type", roleAssignment.PrincipalType)
	}
	if roleAssignment.PrincipalName != nil {
		_ = d.Set("principal_name", roleAssignment.PrincipalName)
	}
	if roleAssignment.TargetName != nil {
		_ = d.Set("target_name", roleAssignment.TargetName)
	}
	if roleAssignment.CreateTime != nil {
		_ = d.Set("create_time", roleAssignment.CreateTime)
	}
	if roleAssignment.UpdateTime != nil {
		_ = d.Set("update_time", roleAssignment.UpdateTime)
	}

	return nil
}

func resourceTencentCloudCicRoleAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_assignment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 6 {
		return fmt.Errorf("roleAssignmentId is broken,%s", d.Id())
	}

	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]
	targetType := idSplit[2]
	targetUinString := idSplit[3]
	principalType := idSplit[4]
	principalId := idSplit[5]

	var (
		deleteRoleAssignmentRequest  = cic.NewDeleteRoleAssignmentRequest()
		deleteRoleAssignmentResponse = cic.NewDeleteRoleAssignmentResponse()
	)
	deleteRoleAssignmentRequest.ZoneId = helper.String(zoneId)
	deleteRoleAssignmentRequest.RoleConfigurationId = helper.String(roleConfigurationId)
	deleteRoleAssignmentRequest.TargetType = helper.String(targetType)
	targetUin, err := strconv.ParseInt(targetUinString, 10, 64)
	if err != nil {
		return err
	}
	deleteRoleAssignmentRequest.TargetUin = helper.Int64(targetUin)
	deleteRoleAssignmentRequest.PrincipalType = helper.String(principalType)
	deleteRoleAssignmentRequest.PrincipalId = helper.String(principalId)

	// Serialize Create/Delete on the same target so concurrent for_each destroy
	// has a well-defined "last assignment" for deprovision / Dismantle.
	unlock := cicRoleAssignmentTargetUnlock(zoneId, roleConfigurationId, targetType, targetUin)
	defer unlock()

	deprovisionStrategy := "None"
	if v, ok := d.GetOk("deprovision_strategy"); ok {
		deprovisionStrategy = v.(string)
	}
	deleteRoleAssignmentRequest.DeprovisionStrategy = helper.String(deprovisionStrategy)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteRoleAssignment(deleteRoleAssignmentRequest)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, deleteRoleAssignmentRequest.GetAction(), deleteRoleAssignmentRequest.ToJsonString(), result.ToJsonString())
		}
		deleteRoleAssignmentResponse = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center role assignment failed, reason:%+v", logId, err)
		return err
	}

	if deleteRoleAssignmentResponse == nil || deleteRoleAssignmentResponse.Response == nil {
		return fmt.Errorf("delete role assignment response is nil")
	}
	if deleteRoleAssignmentResponse.Response.Task == nil {
		return fmt.Errorf("delete role assignment task is nil")
	}
	task := deleteRoleAssignmentResponse.Response.Task
	if task.Status != nil && *task.Status == TASK_STATUS_FAILED {
		if task.FailureReason != nil {
			return fmt.Errorf("delete role assignment failed, failure reason:%s", *task.FailureReason)
		}
		return fmt.Errorf("delete role assignment failed")
	}
	if task.TaskId == nil {
		return fmt.Errorf("delete role assignment task id is nil")
	}
	conf := BuildStateChangeConf([]string{}, []string{TASK_STATUS_SUCCESS, TASK_STATUS_FAILED}, 2*readRetryTimeout, time.Second, service.AssignmentTaskStatusStateRefreshFunc(zoneId, *task.TaskId, []string{}))
	if object, e := conf.WaitForState(); e != nil {
		return e
	} else {
		taskStatus := object.(*cic.TaskStatus)
		if taskStatus.Status != nil && *taskStatus.Status == TASK_STATUS_FAILED {
			return fmt.Errorf("delete role assignment failed")
		}
	}

	// None: only remove authorization (keep deployment).
	// DeprovisionForLast...: after delete, List remaining bindings on this target;
	// dismantle only when none remain. Target lock + list avoids concurrent races.
	if deprovisionStrategy != "DeprovisionForLastRoleAssignmentOnAccount" {
		return nil
	}

	ctx := context.WithValue(context.Background(), logIdKey, logId)
	var remaining int64
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		count, e := service.CountCicRoleAssignmentsOnTarget(ctx, zoneId, roleConfigurationId, targetType, targetUin)
		if e != nil {
			return retryError(e)
		}
		remaining = count
		return nil
	})
	if err != nil {
		return err
	}
	if remaining > 0 {
		log.Printf("[INFO]%s skip dismantle role configuration %s on target %d: %d authorization(s) remain",
			logId, roleConfigurationId, targetUin, remaining)
		return nil
	}

	dismantleRoleConfigurationRequest := cic.NewDismantleRoleConfigurationRequest()
	dismantleRoleConfigurationRequest.RoleConfigurationId = helper.String(roleConfigurationId)
	dismantleRoleConfigurationRequest.ZoneId = helper.String(zoneId)
	dismantleRoleConfigurationRequest.TargetType = helper.String(targetType)
	dismantleRoleConfigurationRequest.TargetUin = helper.Int64(targetUin)

	var dismantleRoleConfigurationResponse *cic.DismantleRoleConfigurationResponse
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DismantleRoleConfiguration(dismantleRoleConfigurationRequest)
		if e != nil {
			msg := e.Error()
			if strings.Contains(msg, "RoleConfigurationAuthorizationExist") {
				count, listErr := service.CountCicRoleAssignmentsOnTarget(ctx, zoneId, roleConfigurationId, targetType, targetUin)
				if listErr != nil {
					return retryError(listErr)
				}
				if count > 0 {
					log.Printf("[INFO]%s skip dismantle after AuthorizationExist: %d authorization(s) remain", logId, count)
					return nil
				}
				return resource.RetryableError(fmt.Errorf("dismantle raced with concurrent authorization change: %v", e))
			}
			if isCicDismantleAlreadyGone(msg) {
				log.Printf("[WARN]%s role configuration already dismantled: %v", logId, e)
				return nil
			}
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, dismantleRoleConfigurationRequest.GetAction(), dismantleRoleConfigurationRequest.ToJsonString(), result.ToJsonString())
		dismantleRoleConfigurationResponse = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s dismantle role configuration failed, reason:%+v", logId, err)
		return err
	}
	if dismantleRoleConfigurationResponse == nil || dismantleRoleConfigurationResponse.Response == nil ||
		dismantleRoleConfigurationResponse.Response.Task == nil {
		return nil
	}

	dismantleTask := dismantleRoleConfigurationResponse.Response.Task
	if dismantleTask.TaskStatus != nil && *dismantleTask.TaskStatus == TASK_STATUS_FAILED {
		return fmt.Errorf("dismantle role assignment task failed")
	}
	if dismantleTask.TaskId == nil {
		return fmt.Errorf("dismantle role assignment task id is nil")
	}
	conf = BuildStateChangeConf([]string{}, []string{TASK_STATUS_SUCCESS, TASK_STATUS_FAILED}, 2*readRetryTimeout, time.Second, service.AssignmentTaskStatusStateRefreshFunc(zoneId, *dismantleTask.TaskId, []string{}))
	if object, e := conf.WaitForState(); e != nil {
		return e
	} else {
		taskStatus := object.(*cic.TaskStatus)
		if taskStatus.Status != nil && *taskStatus.Status == TASK_STATUS_FAILED {
			reason := ""
			if taskStatus.FailureReason != nil {
				reason = *taskStatus.FailureReason
			}
			// Concurrent destroy may accept multiple Deprovision tasks; losers fail with RecordNotFound.
			if isCicDismantleAlreadyGone(reason) {
				log.Printf("[WARN]%s dismantle task failed but deployment already gone: %s", logId, reason)
				return nil
			}
			return fmt.Errorf("dismantle role assignment task failed")
		}
	}

	return nil
}

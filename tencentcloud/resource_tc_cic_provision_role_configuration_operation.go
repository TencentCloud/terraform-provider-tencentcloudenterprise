/*
Provides a CIC permission configuration deployment operation. With no target
arguments, the resource discovers the current identity center and deploys every
permission configuration whose status is DeployedRequired.

If automatic discovery finds no matching record, the operation succeeds and
exports `provisioned_count = 0` without submitting a deployment task.

`role_configuration_id`, `target_type`, and `target_uin` must either all be
specified or all be omitted. Because this is an operation resource, Terraform
runs it when the resource is created. Use `terraform apply -replace=...` to run
the same configuration again.

# Example Usage

```hcl

	# Deploy every role configuration whose status is DeployedRequired.
	resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "all_required" {}

```

# Deploy or redeploy one specific target

```hcl

	resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "target" {
	  role_configuration_id = "rc-xxxxxxxxxxxx"
	  target_type           = "MemberUin"
	  target_uin            = 100001234567
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_provision_role_configuration_operation", CNDescription{
		TerraformTypeCN: "身份中心权限配置部署操作",
		DescriptionCN:   "提供身份中心权限配置部署操作资源，用于将权限配置（重新）部署到成员账号。",
		AttributesCN: map[string]string{
			"zone_id":               "空间ID，不填写时自动读取身份中心空间",
			"role_configuration_id": "权限配置ID；与目标类型、目标UIN同时填写时执行定向部署",
			"target_type":           "目标账号类型；与权限配置ID、目标UIN同时填写时执行定向部署",
			"target_uin":            "目标账号UIN；与权限配置ID、目标类型同时填写时执行定向部署",
			"provisioned_count":     "本次部署的目标数量",
			"provisioned":           "本次部署的目标明细",
		},
	})
}

func resourceTencentCloudCicProvisionRoleConfigurationOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCicProvisionRoleConfigurationOperationCreate,
		Read:   resourceTencentCloudCicProvisionRoleConfigurationOperationRead,
		Delete: resourceTencentCloudCicProvisionRoleConfigurationOperationDelete,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
			for _, name := range []string{"role_configuration_id", "target_type", "target_uin"} {
				if !d.NewValueKnown(name) {
					return nil
				}
			}
			return validateCicProvisionRoleConfigurationOperationTarget(d)
		},
		Description: "Provide an identity center permission configuration provision operation resource. With no target arguments, it deploys every DeployedRequired record discovered from CIC. With a complete target, it deploys or redeploys that target.",
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Space ID. If omitted, the provider reads it from the current CIC identity center.",
			},
			"role_configuration_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Permission configuration ID for a targeted deployment. Must be specified together with target_type and target_uin.",
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ManagerUin",
					"MemberUin",
				}, false),
				Description: "Type of the synchronized target account for a targeted deployment. Valid values: ManagerUin and MemberUin. Must be specified together with role_configuration_id and target_uin.",
			},
			"target_uin": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "UIN of the target account for a targeted deployment. Must be specified together with role_configuration_id and target_type.",
			},
			"provisioned_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of targets provisioned by this operation.",
			},
			"provisioned": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details of targets provisioned by this operation.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"role_configuration_id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Permission configuration ID deployed by this operation.",
					},
					"role_configuration_name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Permission configuration name returned by CIC.",
					},
					"target_type": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Type of the target account deployed by this operation.",
					},
					"target_uin": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "UIN of the target account deployed by this operation.",
					},
					"target_name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Name of the target account returned by CIC.",
					},
					"deployment_status": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Deployment status after this operation completes.",
					},
				}},
			},
		},
	}
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	client := meta.(*TencentCloudClient)
	service := CicService{client: client.apiV3Conn}
	if err := validateCicProvisionRoleConfigurationOperationTarget(d); err != nil {
		return err
	}

	zoneId := ""
	if value, ok := d.GetOk("zone_id"); ok {
		zoneId = value.(string)
	} else {
		identityCenter, err := service.DescribeCicIdentityCenter(ctx)
		if err != nil {
			return fmt.Errorf("describe CIC identity center failed: %w", err)
		}
		if identityCenter == nil || identityCenter.Response == nil || identityCenter.Response.ZoneId == nil || *identityCenter.Response.ZoneId == "" {
			return fmt.Errorf("CIC identity center returned an empty space ID")
		}
		zoneId = *identityCenter.Response.ZoneId
		if err := d.Set("zone_id", zoneId); err != nil {
			return fmt.Errorf("set CIC space ID failed: %w", err)
		}
	}

	roleConfigurationId, roleConfigurationIdSet := getCicProvisionOperationString(d, "role_configuration_id")
	targetType, targetTypeSet := getCicProvisionOperationString(d, "target_type")
	targetUinValue, targetUinSet := d.GetOk("target_uin")
	explicitTarget := roleConfigurationIdSet && targetTypeSet && targetUinSet

	// An explicit target preserves the original single-operation behavior. With
	// no target arguments, mirror the CIC console's one-click flow and discover
	// only records whose status is DeployedRequired.
	targets := make([]*cic.RoleConfigurationProvisionings, 0)
	if explicitTarget {
		targets = append(targets, &cic.RoleConfigurationProvisionings{
			RoleConfigurationId: helper.String(roleConfigurationId),
			TargetType:          helper.String(targetType),
			TargetUin:           helper.IntInt64(targetUinValue.(int)),
		})
	} else {
		var provisionings []*cic.RoleConfigurationProvisionings
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := service.ListCicRoleConfigurationProvisionings(ctx, zoneId, "DeployedRequired", "", "", 0)
			if e != nil {
				return retryError(e)
			}
			provisionings = result
			return nil
		})
		if err != nil {
			return fmt.Errorf("list CIC role configuration provisionings with status %q failed: %w", "DeployedRequired", err)
		}
		targets = append(targets, provisionings...)
	}

	uniqueTargets := make([]*cic.RoleConfigurationProvisionings, 0, len(targets))
	seenTargets := make(map[string]struct{}, len(targets))
	for _, target := range targets {
		key := cicProvisioningTargetKey(target)
		if _, exists := seenTargets[key]; exists {
			continue
		}
		seenTargets[key] = struct{}{}
		uniqueTargets = append(uniqueTargets, target)
	}
	targets = uniqueTargets

	// The list endpoint normally returns a stable order, but sorting makes the
	// Terraform ID and state deterministic across refreshes and pages.
	sort.SliceStable(targets, func(i, j int) bool {
		return cicProvisioningTargetKey(targets[i]) < cicProvisioningTargetKey(targets[j])
	})

	provisioned := make([]map[string]interface{}, 0, len(targets))
	ids := []string{zoneId}
	for _, target := range targets {
		if target == nil || target.TargetUin == nil || *target.TargetUin <= 0 {
			return fmt.Errorf("CIC role configuration provisioning returned an invalid target UIN")
		}
		provisioningRoleConfigurationId := roleConfigurationId
		if target.RoleConfigurationId != nil && *target.RoleConfigurationId != "" {
			provisioningRoleConfigurationId = *target.RoleConfigurationId
		}
		provisioningTargetType := targetType
		if target.TargetType != nil && *target.TargetType != "" {
			provisioningTargetType = *target.TargetType
		}
		if provisioningRoleConfigurationId == "" || provisioningTargetType == "" {
			return fmt.Errorf("CIC role configuration provisioning returned an incomplete target")
		}
		if err := provisionCicRoleConfiguration(ctx, client, zoneId, provisioningRoleConfigurationId, provisioningTargetType, *target.TargetUin); err != nil {
			return fmt.Errorf("provision CIC role configuration %s to target %s#%d failed: %w", provisioningRoleConfigurationId, provisioningTargetType, *target.TargetUin, err)
		}

		entry := map[string]interface{}{
			"role_configuration_id": provisioningRoleConfigurationId,
			"target_type":           provisioningTargetType,
			"target_uin":            *target.TargetUin,
		}
		if target.RoleConfigurationName != nil {
			entry["role_configuration_name"] = *target.RoleConfigurationName
		}
		if target.TargetName != nil {
			entry["target_name"] = *target.TargetName
		}
		entry["deployment_status"] = "Deployed"
		provisioned = append(provisioned, entry)
		ids = append(ids, cicProvisioningTargetKey(target))
	}

	_ = d.Set("provisioned_count", len(provisioned))
	_ = d.Set("provisioned", provisioned)
	if explicitTarget {
		d.SetId(strings.Join([]string{zoneId, roleConfigurationId, targetType, helper.IntToStr(targetUinValue.(int))}, FILED_SP))
	} else {
		d.SetId(helper.DataResourceIdsHash(ids))
	}
	return resourceTencentCloudCicProvisionRoleConfigurationOperationRead(d, meta)
}

type cicProvisionOperationSchemaGetter interface {
	GetOk(string) (interface{}, bool)
}

func validateCicProvisionRoleConfigurationOperationTarget(d cicProvisionOperationSchemaGetter) error {
	_, roleConfigurationIdSet := getCicProvisionOperationString(d, "role_configuration_id")
	_, targetTypeSet := getCicProvisionOperationString(d, "target_type")
	_, targetUinSet := d.GetOk("target_uin")
	setCount := 0
	for _, set := range []bool{roleConfigurationIdSet, targetTypeSet, targetUinSet} {
		if set {
			setCount++
		}
	}
	if setCount != 0 && setCount != 3 {
		return fmt.Errorf("role_configuration_id, target_type, and target_uin must either all be specified for a targeted deployment or all be omitted for automatic DeployedRequired deployment")
	}
	return nil
}

func getCicProvisionOperationString(d cicProvisionOperationSchemaGetter, name string) (string, bool) {
	value, ok := d.GetOk(name)
	if !ok {
		return "", false
	}
	text, ok := value.(string)
	return text, ok && text != ""
}

// provisionCicRoleConfiguration submits one deployment operation and waits for
// its asynchronous CIC task to finish. Keeping this in one helper ensures the
// explicit-UIN and auto-discovery paths have identical retry and failure logic.
func provisionCicRoleConfiguration(ctx context.Context, client *TencentCloudClient, zoneId, roleConfigurationId, targetType string, targetUin int64) error {
	logId := getLogId(ctx)
	request := cic.NewProvisionRoleConfigurationRequest()
	request.ZoneId = helper.String(zoneId)
	request.RoleConfigurationId = helper.String(roleConfigurationId)
	request.TargetType = helper.String(targetType)
	request.TargetUin = helper.Int64(targetUin)
	var response *cic.ProvisionRoleConfigurationResponse

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := client.apiV3Conn.UseCicClient().ProvisionRoleConfiguration(request)
		if e != nil {
			return retryError(e)
		}
		if result == nil || result.Response == nil || result.Response.Task == nil || result.Response.Task.TaskId == nil {
			return resource.NonRetryableError(fmt.Errorf("CIC returned no provisioning task"))
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		return err
	}

	service := CicService{client: client.apiV3Conn}
	taskId := *response.Response.Task.TaskId
	conf := BuildStateChangeConf([]string{"InProgress", ""}, []string{"Success", "Failed"}, 600*time.Second, 5*time.Second,
		service.AssignmentTaskStatusStateRefreshFunc(zoneId, taskId, []string{"Failed"}))
	object, err := conf.WaitForStateContext(ctx)
	if err != nil {
		return err
	}
	taskStatus, ok := object.(*cic.TaskStatus)
	if !ok || taskStatus == nil || taskStatus.Status == nil {
		return fmt.Errorf("CIC provisioning task %s returned an empty status", taskId)
	}
	if *taskStatus.Status == "Failed" {
		reason := ""
		if taskStatus.FailureReason != nil {
			reason = *taskStatus.FailureReason
		}
		if reason != "" {
			return fmt.Errorf("CIC provisioning task %s failed: %s", taskId, reason)
		}
		return fmt.Errorf("CIC provisioning task %s failed", taskId)
	}
	return nil
}

func cicProvisioningTargetKey(target *cic.RoleConfigurationProvisionings) string {
	if target == nil {
		return "<nil>"
	}
	roleConfigurationId := ""
	if target.RoleConfigurationId != nil {
		roleConfigurationId = *target.RoleConfigurationId
	}
	targetType := ""
	if target.TargetType != nil {
		targetType = *target.TargetType
	}
	targetUin := int64(0)
	if target.TargetUin != nil {
		targetUin = *target.TargetUin
	}
	return strings.Join([]string{roleConfigurationId, targetType, helper.Int64ToStr(targetUin)}, FILED_SP)
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.read")()
	defer inconsistentCheck(d, meta)()
	return nil
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.delete")()
	defer inconsistentCheck(d, meta)()
	return nil
}

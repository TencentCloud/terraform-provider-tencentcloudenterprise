/*
Provides a resource to create an organization cic_user_sync_provisioning

Example Usage

```hcl
resource "tencentcloudenterprise_cic_user_sync_provisioning" "cic_user_sync_provisioning" {
  zone_id = "z-xxxxxx"
  description = "tf-test"
  deletion_strategy = "Keep"
  duplication_strategy = "TakeOver"
  principal_id = "u-xxxxxx"
  principal_type = "User"
  target_uin = "xxxxxx"
  target_type = "MemberUin"
}
```

Import

organization cic_user_sync_provisioning can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_user_sync_provisioning.cic_user_sync_provisioning ${zoneId}#${userProvisioningId}
```
 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_user_sync_provisioning", CNDescription{
		TerraformTypeCN: "身份中心用户同步供应",
		DescriptionCN:   "提供身份中心用户同步供应资源，用于管理用户同步。",
	})
}

func resourceTencentCloudCicUserSyncProvisioning() *schema.Resource {
	return &schema.Resource{
		Description: "Provide identity center user synchronization provisioning resources for managing user synchronization.",
		Create: resourceTencentCloudCicUserSyncProvisioningCreate,
		Read:   resourceTencentCloudCicUserSyncProvisioningRead,
		Update: resourceTencentCloudCicUserSyncProvisioningUpdate,
		Delete: resourceTencentCloudCicUserSyncProvisioningDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space ID.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description.",
			},
			"principal_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identity ID for the CAM user synchronization. Valid values:\nWhen the PrincipalType value is Group, it is the CIC user group ID (g-********).\nWhen the PrincipalType value is User, it is the CIC user ID (u-********).",
			},
			"principal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identity type for the CAM user synchronization. Valid values:\n\nUser: indicates that the identity for the CAM user synchronization is a CIC user.\nGroup: indicates that the identity for the CAM user synchronization is a CIC user group.",
			},
			"target_uin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "UIN of the synchronized target account of the Tencent Cloud Cic.",
			},
			"duplication_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Conflict policy. It indicates the handling policy for existence of a user with the same username when CIC users are synchronized to CAM. Valid values: KeepBoth: Keep both, that is, add the _cic suffix to the CIC user's username and then try to create a CAM user with the username when CIC users are synchronized to CAM and a user with the same username already exists in CAM; TakeOver: Replace, that is, directly replace the existing CAM user with the synchronized CIC user when CIC users are synchronized to CAM and a user with the same username already exists in CAM.",
			},
			"deletion_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Deletion policy. It indicates the handling policy for CAM users already synchronized when the CAM user synchronization is deleted. Valid values: Delete: Delete the CAM users already synchronized from CIC to CAM when the CAM user synchronization is deleted; Keep: Keep the CAM users already synchronized from CIC to CAM when the CAM user synchronization is deleted.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the synchronized target account of the Tencent Cloud Cic. ManagerUin: admin account; MemberUin: member account.",
			},
			"user_provisioning_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User provisioning id.",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Status of CAM user synchronization. Value:\n" +
					"	* Enabled: CAM user synchronization is enabled;\n" +
					"	* Disabled: CAM user synchronization is not enabled.",
			},
			"principal_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identity name of the CAM user synchronization. Value: When PrincipalType is Group, the value is the CIC user group name; When PrincipalType takes the value to User, the value is the CIC user name.",
			},
			"target_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Group account The name of the target account..",
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
		},
	}
}

func resourceTencentCloudCicUserSyncProvisioningCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_sync_provisioning.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId   string
		request  = cic.NewCreateUserSyncProvisioningRequest()
		response = cic.NewCreateUserSyncProvisioningResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
		request.ZoneId = helper.String(zoneId)
	}

	userSyncProvisioning := cic.UserSyncProvisioning{}
	if v, ok := d.GetOk("description"); ok {
		userSyncProvisioning.Description = helper.String(v.(string))
	}
	if v, ok := d.GetOk("principal_id"); ok {
		userSyncProvisioning.PrincipalId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("principal_type"); ok {
		userSyncProvisioning.PrincipalType = helper.String(v.(string))
	}
	if v, ok := d.GetOk("target_uin"); ok {
		userSyncProvisioning.TargetUin = helper.IntInt64(v.(int))
	}
	if v, ok := d.GetOk("duplication_strategy"); ok {
		userSyncProvisioning.DuplicationStrategy = helper.String(v.(string))
	}
	if v, ok := d.GetOk("deletion_strategy"); ok {
		userSyncProvisioning.DeletionStrategy = helper.String(v.(string))
	}
	if v, ok := d.GetOk("target_type"); ok {
		userSyncProvisioning.TargetType = helper.String(v.(string))
	}

	request.UserSyncProvisionings = []*cic.UserSyncProvisioning{&userSyncProvisioning}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateUserSyncProvisioning(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center user sync provisioning failed, reason:%+v", logId, err)
		return err
	}
	if len(response.Response.Tasks) > 0 {
		task := response.Response.Tasks[0]
		taskId := *task.TaskId
		userProvisioningId := *task.UserProvisioningId
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			request := cic.NewGetProvisioningTaskStatusRequest()
			request.ZoneId = helper.String(zoneId)
			request.TaskId = helper.String(taskId)
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().GetProvisioningTaskStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			}
			if response.Response.TaskStatus != nil {
				status := *response.Response.TaskStatus.Status
				if status == "Failed" {
					return resource.NonRetryableError(fmt.Errorf("task status is %s", status))
				}
				if status != "Success" {
					return resource.RetryableError(fmt.Errorf("task status is %s", status))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s create identity center user sync provisioning failed, reason:%+v", logId, err)
			return err
		}
		d.SetId(strings.Join([]string{zoneId, userProvisioningId}, FILED_SP))

	}

	return resourceTencentCloudCicUserSyncProvisioningRead(d, meta)
}

func resourceTencentCloudCicUserSyncProvisioningRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_sync_provisioning.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	zoneId := idSplit[0]
	userProvisioningId := idSplit[1]
	respData, err := service.DescribeCicUserSyncProvisioningById(ctx, zoneId, userProvisioningId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_user_sync_provisioning` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	_ = d.Set("zone_id", zoneId)
	_ = d.Set("user_provisioning_id", respData.UserProvisioningId)

	if respData.Description != nil {
		_ = d.Set("description", respData.Description)
	}

	if respData.Status != nil {
		_ = d.Set("status", respData.Status)
	}

	if respData.PrincipalId != nil {
		_ = d.Set("principal_id", respData.PrincipalId)
	}

	if respData.PrincipalName != nil {
		_ = d.Set("principal_name", respData.PrincipalName)
	}

	if respData.PrincipalType != nil {
		_ = d.Set("principal_type", respData.PrincipalType)
	}

	if respData.TargetUin != nil {
		_ = d.Set("target_uin", respData.TargetUin)
	}

	if respData.TargetName != nil {
		_ = d.Set("target_name", respData.TargetName)
	}

	if respData.DuplicationStrategy != nil {
		_ = d.Set("duplication_strategy", respData.DuplicationStrategy)
	}

	if respData.DeletionStrategy != nil {
		_ = d.Set("deletion_strategy", respData.DeletionStrategy)
	}

	if respData.CreateTime != nil {
		_ = d.Set("create_time", respData.CreateTime)
	}

	if respData.UpdateTime != nil {
		_ = d.Set("update_time", respData.UpdateTime)
	}

	if respData.TargetType != nil {
		_ = d.Set("target_type", respData.TargetType)
	}

	return nil
}

func resourceTencentCloudCicUserSyncProvisioningUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_sync_provisioning.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	immutableArgs := []string{"zone_id", "user_sync_provisionings"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	zoneId := idSplit[0]
	userProvisioningId := idSplit[1]

	needChange := false
	mutableArgs := []string{"description", "duplication_stateful", "deletion_strategy"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewUpdateUserSyncProvisioningRequest()

		request.ZoneId = helper.String(zoneId)

		request.UserProvisioningId = helper.String(userProvisioningId)

		if v, ok := d.GetOk("description"); ok {
			request.NewDescription = helper.String(v.(string))
		}

		if v, ok := d.GetOk("duplication_stateful"); ok {
			request.NewDuplicationStateful = helper.String(v.(string))
		}

		if v, ok := d.GetOk("deletion_strategy"); ok {
			request.NewDeletionStrategy = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateUserSyncProvisioning(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center user sync provisioning failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCicUserSyncProvisioningRead(d, meta)
}

func resourceTencentCloudCicUserSyncProvisioningDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_sync_provisioning.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	zoneId := idSplit[0]
	userProvisioningId := idSplit[1]

	var (
		request  = cic.NewDeleteUserSyncProvisioningRequest()
		response = cic.NewDeleteUserSyncProvisioningResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.UserProvisioningId = helper.String(userProvisioningId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteUserSyncProvisioning(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center user sync provisioning failed, reason:%+v", logId, err)
		return err
	}

	if response.Response != nil && response.Response.Tasks != nil && response.Response.Tasks.TaskId != nil {
		taskId := *response.Response.Tasks.TaskId
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			request := cic.NewGetProvisioningTaskStatusRequest()
			request.ZoneId = helper.String(zoneId)
			request.TaskId = helper.String(taskId)
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().GetProvisioningTaskStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			}
			if response.Response.TaskStatus != nil {
				status := *response.Response.TaskStatus.Status
				if status == "Failed" {
					return resource.NonRetryableError(fmt.Errorf("task status is %s", status))
				}
				if status != "Success" {
					return resource.RetryableError(fmt.Errorf("task status is %s", status))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s delete identity center user sync provisioning failed, reason:%+v", logId, err)
			return err
		}
	}

	return nil
}

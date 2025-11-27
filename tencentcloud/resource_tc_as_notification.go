/*
Provides a resource for an AS (Auto scaling) notification.

# Example Usage

```hcl

	resource "tencentcloudenterprise_as_notification" "as_notification" {
	  scaling_group_id            = "sg-12af45"
	  notification_types          = ["SCALE_OUT_FAILED", "SCALE_IN_SUCCESSFUL", "SCALE_IN_FAILED", "REPLACE_UNHEALTHY_INSTANCE_FAILED"]
	  notification_user_group_ids = ["76955"]
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	as "terraform-provider-tencentcloudenterprise/sdk/as/v20180419"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_as_notification", CNDescription{
		TerraformTypeCN: "创建通知",
		DescriptionCN:   "提供自动伸缩通知资源，用于创建和管理AS通知。",
		AttributesCN: map[string]string{
			"scaling_group_id":            "伸缩组ID",
			"notification_types":          "通知类型",
			"notification_user_group_ids": "通知用户组ID",
		},
	})
}
func resourceTencentCloudAsNotification() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource for an AS (Auto scaling) notification.",
		Create:      resourceTencentCloudAsNotificationCreate,
		Read:        resourceTencentCloudAsNotificationRead,
		Update:      resourceTencentCloudAsNotificationUpdate,
		Delete:      resourceTencentCloudAsNotificationDelete,

		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of a scaling group.",
			},
			"notification_types": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateAllowedStringValue(SCALING_GROUP_NOTIFICATION_TYPE),
				},
			},
			"notification_user_group_ids": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A group of user IDs to be notified.",
			},
		},
	}
}

func resourceTencentCloudAsNotificationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_notification.create")()

	logId := getLogId(contextNil)

	request := as.NewCreateNotificationConfigurationRequest()
	request.AutoScalingGroupId = helper.String(d.Get("scaling_group_id").(string))
	notificationTypes := d.Get("notification_types").([]interface{})
	request.NotificationTypes = make([]*string, 0, len(notificationTypes))
	for _, value := range notificationTypes {
		request.NotificationTypes = append(request.NotificationTypes, helper.String(value.(string)))
	}
	userGroupIds := d.Get("notification_user_group_ids").([]interface{})
	request.NotificationUserGroupIds = make([]*string, 0, len(userGroupIds))
	for _, value := range userGroupIds {
		request.NotificationUserGroupIds = append(request.NotificationUserGroupIds, helper.String(value.(string)))
	}

	response, err := meta.(*TencentCloudClient).apiV3Conn.UseAsClient().CreateNotificationConfiguration(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response.AutoScalingNotificationId == nil {
		return fmt.Errorf("scaling policy id is nil")
	}
	d.SetId(*response.Response.AutoScalingNotificationId)

	return resourceTencentCloudAsNotificationRead(d, meta)
}

func resourceTencentCloudAsNotificationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_notification.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	notificationId := d.Id()
	asService := AsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		notification, has, e := asService.DescribeNotificationById(ctx, notificationId)
		if e != nil {
			return retryError(e)
		}
		if has == 0 {
			d.SetId("")
			return nil
		}
		_ = d.Set("scaling_group_id", *notification.AutoScalingGroupId)
		_ = d.Set("notification_type", helper.StringsInterfaces(notification.NotificationTypes))
		_ = d.Set("notification_user_group_ids", helper.StringsInterfaces(notification.NotificationUserGroupIds))
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudAsNotificationUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_notification.update")()

	logId := getLogId(contextNil)

	request := as.NewModifyNotificationConfigurationRequest()
	notificationId := d.Id()
	request.AutoScalingNotificationId = &notificationId
	if d.HasChange("notification_type") {
		notificationTypes := d.Get("notification_types").([]interface{})
		request.NotificationTypes = make([]*string, 0, len(notificationTypes))
		for _, value := range notificationTypes {
			request.NotificationTypes = append(request.NotificationTypes, helper.String(value.(string)))
		}
	}
	if d.HasChange("notification_user_group_ids") {
		userGroupIds := d.Get("notification_user_group_ids").([]interface{})
		request.NotificationUserGroupIds = make([]*string, 0, len(userGroupIds))
		for _, value := range userGroupIds {
			request.NotificationUserGroupIds = append(request.NotificationUserGroupIds, helper.String(value.(string)))
		}
	}

	response, err := meta.(*TencentCloudClient).apiV3Conn.UseAsClient().ModifyNotificationConfiguration(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func resourceTencentCloudAsNotificationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_notification.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	notificationId := d.Id()
	asService := AsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := asService.DeleteNotification(ctx, notificationId)
	if err != nil {
		return err
	}

	return nil
}

/*
Provides a resource to create and manage TDMQ RabbitMQ user

Example Usage

### Create a RabbitMQ user with administrator role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "example" {
  instance_id = "amqp-xxxxxxxx"
  user        = "admin_user"
  password    = "AdminPassword123!"
  description = "Administrator user for RabbitMQ"
  tags        = ["administrator"]
}
```

### Create a RabbitMQ user with monitoring role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "monitoring" {
  instance_id = "amqp-xxxxxxxx"
  user        = "monitor_user"
  password    = "MonitorPass123!"
  description = "Monitoring user for RabbitMQ"
  tags        = ["monitoring"]
}
```

### Create a RabbitMQ user with management role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "management" {
  instance_id = "amqp-xxxxxxxx"
  user        = "mgmt_user"
  password    = "ManagementPass123!"
  description = "Management user for RabbitMQ"
  tags        = ["management"]
}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_user", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ用户",
		DescriptionCN:   "提供TDMQ RabbitMQ用户资源，用于创建和管理TDMQ RabbitMQ用户。",
		AttributesCN: map[string]string{
			"instance_id":     "RabbitMQ实例ID",
			"user":            "用户名,用于登录RabbitMQ服务",
			"password":        "密码,用于登录RabbitMQ服务",
			"description":     "用户描述信息",
			"tags":            "用户角色,可选值:administrator(管理员)、monitoring(监控)、policymaker(策略制定者)、management(管理)、none(无角色)",
		},
	})
}

func resourceTencentCloudTdmqRabbitmqUser() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTdmqRabbitmqUserCreate,
		Read:        resourceTencentCloudTdmqRabbitmqUserRead,
		Update:      resourceTencentCloudTdmqRabbitmqUserUpdate,
		Delete:      resourceTencentCloudTdmqRabbitmqUserDelete,
		Description: "Provides a resource to create and manage TDMQ RabbitMQ user",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "RabbitMQ cluster instance ID. The ID of the RabbitMQ instance where the user will be created.",
			},
			"user": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Username for RabbitMQ authentication. This will be used to log in to the RabbitMQ server and management console.",
			},
			"password": {
				Required:    true,
				Type:        schema.TypeString,
				Sensitive:   true,
				Description: "Password for RabbitMQ authentication. This will be used to log in to the RabbitMQ server and management console. The password is stored securely and will not be displayed in logs.",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description for the user. Provides additional information about the user's purpose or role.",
			},
			"tags": {
				Required:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "User tags. Determines the user's access permissions to RabbitMQ Management console. Valid values: `administrator` (full admin access, default), `monitoring` (read-only monitoring access), `policymaker` (can manage policies), `management` (can manage resources), `none` (no management console access).",
			},
		},
	}
}

func resourceTencentCloudTdmqRabbitmqUserCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_user.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = tdmq.NewCreateRabbitMQUserRequest()
		response   = tdmq.NewCreateRabbitMQUserResponse()
		instanceId string
		user       string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	if v, ok := d.GetOk("user"); ok {
		request.User = helper.String(v.(string))
	}

	if v, ok := d.GetOk("password"); ok {
		request.Password = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tags"); ok {
		request.Tags = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRabbitMQUser(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq rabbitmqUser failed, reason:%+v", logId, err)
		return err
	}

	user = *response.Response.User

	d.SetId(strings.Join([]string{instanceId, user}, FILED_SP))
	return resourceTencentCloudTdmqRabbitmqUserRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqUserRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_user.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	instanceId := idSplit[0]
	user := idSplit[1]

	rabbitmqUser, err := service.DescribeTdmqRabbitmqUserById(ctx, instanceId, user)
	if err != nil {
		return err
	}

	if rabbitmqUser == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TdmqRabbitmqUser` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if rabbitmqUser.InstanceId != nil {
		_ = d.Set("instance_id", rabbitmqUser.InstanceId)
	}

	if rabbitmqUser.User != nil {
		_ = d.Set("user", rabbitmqUser.User)
	}

	if rabbitmqUser.Password != nil {
		_ = d.Set("password", rabbitmqUser.Password)
	}

	if rabbitmqUser.Description != nil {
		_ = d.Set("description", rabbitmqUser.Description)
	}

	if rabbitmqUser.Tags != nil {
		_ = d.Set("tags", rabbitmqUser.Tags)
	}

	return nil
}

func resourceTencentCloudTdmqRabbitmqUserUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_user.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = tdmq.NewModifyRabbitMQUserRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	instanceId := idSplit[0]
	user := idSplit[1]

	immutableArgs := []string{"instance_id", "user", "password"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("description") || d.HasChange("tags") {
		request.InstanceId = &instanceId
		request.User = &user

		if v, ok := d.GetOk("password"); ok {
			request.Password = helper.String(v.(string))
		}

		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}

		if v, ok := d.GetOk("tags"); ok {
			request.Tags = helper.InterfacesStringsPoint(v.([]interface{}))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().ModifyRabbitMQUser(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update tdmq rabbitmqUser failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudTdmqRabbitmqUserRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqUserDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_user.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	instanceId := idSplit[0]
	user := idSplit[1]

	if err := service.DeleteTdmqRabbitmqUserById(ctx, instanceId, user); err != nil {
		return err
	}

	return nil
}

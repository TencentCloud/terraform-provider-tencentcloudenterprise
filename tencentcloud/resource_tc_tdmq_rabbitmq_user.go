package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_user", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ用户",
		DescriptionCN:   "提供TDMQ RabbitMQ用户资源，用于创建和管理TDMQ RabbitMQ用户。",
		AttributesCN: map[string]string{
			"instance_id":     "实例ID",
			"user":            "用户名",
			"password":        "密码",
			"description":     "用户描述",
			"tags":            "用户标签",
			"max_connections": "最大连接数",
			"max_channels":    "最大通道数",
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
				Description: "Cluster instance ID.",
			},
			"user": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Username, used when logging in.",
			},
			"password": {
				Required:    true,
				Type:        schema.TypeString,
				Sensitive:   true,
				Description: "Password, used when logging in.",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Describe.",
			},
			"tags": {
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.",
			},
			"max_connections": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The maximum number of connections for this user, if not filled in, there is no limit.",
			},
			"max_channels": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The maximum number of channels for this user, if not filled in, there is no limit.",
			},
		},
	}
}

func resourceTencentCloudTdmqRabbitmqUserCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_rabbitmq_user.create")()
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

	if v, ok := d.GetOkExists("max_connections"); ok {
		request.MaxConnections = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("max_channels"); ok {
		request.MaxChannels = helper.IntInt64(v.(int))
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
	defer logElapsed("resource.cloud_tdmq_rabbitmq_user.read")()
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

	if rabbitmqUser.MaxConnections != nil {
		_ = d.Set("max_connections", rabbitmqUser.MaxConnections)
	}

	if rabbitmqUser.MaxChannels != nil {
		_ = d.Set("max_channels", rabbitmqUser.MaxChannels)
	}

	return nil
}

func resourceTencentCloudTdmqRabbitmqUserUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_rabbitmq_user.update")()
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

	if d.HasChange("description") || d.HasChange("tags") || d.HasChange("max_connections") || d.HasChange("max_channels") {
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

		if v, ok := d.GetOkExists("max_connections"); ok {
			request.MaxConnections = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOkExists("max_channels"); ok {
			request.MaxChannels = helper.IntInt64(v.(int))
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
	defer logElapsed("resource.cloud_tdmq_rabbitmq_user.delete")()
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

/*
Provides a resource to create and manage TDMQ RabbitMQ virtual host

Example Usage

### Create a basic virtual host

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "example" {
  instance_id = "amqp-xxxxxxxx"
  virtual_host = "my_vhost"
  description  = "Virtual host for application 1"
}
```

### Create a virtual host with mirror queue policy

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "with_mirror" {
  instance_id = "amqp-xxxxxxxx"
  virtual_host = "mirror_vhost"
  description  = "Virtual host with mirror queue policy enabled"
  mirror_queue_policy_flag = true
}
```

### Create a virtual host without mirror queue policy

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "without_mirror" {
  instance_id = "amqp-xxxxxxxx"
  virtual_host = "no_mirror_vhost"
  description  = "Virtual host without mirror queue policy"
  mirror_queue_policy_flag = false
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
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_virtual_host", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ虚拟主机",
		DescriptionCN:   "提供TDMQ RabbitMQ虚拟主机资源，用于创建和管理TDMQ RabbitMQ虚拟主机。",
		AttributesCN: map[string]string{
			"instance_id":               "RabbitMQ实例ID",
			"virtual_host":              "vhost名称,用于隔离不同应用的消息",
			"description":               "vhost描述信息",
			"mirror_queue_policy_flag": "是否创建镜像队列策略,true为创建,false为不创建,默认true",
		},
	})
}

func resourceTencentCloudTdmqRabbitmqVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTdmqRabbitmqVirtualHostCreate,
		Read:        resourceTencentCloudTdmqRabbitmqVirtualHostRead,
		Update:      resourceTencentCloudTdmqRabbitmqVirtualHostUpdate,
		Delete:      resourceTencentCloudTdmqRabbitmqVirtualHostDelete,
		Description: "Provides a resource to create and manage TDMQ RabbitMQ virtual host",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "RabbitMQ cluster instance ID. The ID of the RabbitMQ instance where the virtual host will be created.",
			},
			"virtual_host": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Virtual host (vhost) name. Virtual hosts provide logical grouping and separation of resources (exchanges, queues, bindings) within a RabbitMQ instance, allowing multiple applications to share the same RabbitMQ instance securely.",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description for the virtual host. Provides additional information about the vhost's purpose or usage.",
			},
		"trace_flag": {
			Computed:    true,
			Type:        schema.TypeBool,
			Description: "Message tracing switch status (read-only).",
		},
			"mirror_queue_policy_flag": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Default:     true,
				Description: "Whether to create a mirror queue policy. When enabled (`true`), a mirror queue policy will be automatically created to replicate queues across cluster nodes for high availability. When disabled (`false`), no mirror queue policy is created. Default is `true` (enabled). Note: This can only be set during virtual host creation and cannot be modified afterwards.",
			},
		},
	}
}

func resourceTencentCloudTdmqRabbitmqVirtualHostCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_virtual_host.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		request     = tdmq.NewCreateRabbitMQVirtualHostRequest()
		response    = tdmq.NewCreateRabbitMQVirtualHostResponse()
		instanceId  string
		virtualHost string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	if v, ok := d.GetOk("virtual_host"); ok {
		request.VirtualHost = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	request.TraceFlag = helper.Bool(false)

	if v, ok := d.GetOkExists("mirror_queue_policy_flag"); ok {
		request.MirrorQueuePolicyFlag = helper.Bool(v.(bool))
	} else {
		request.MirrorQueuePolicyFlag = helper.Bool(true)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRabbitMQVirtualHost(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq rabbitmqVirtualHost failed, reason:%+v", logId, err)
		return err
	}

	virtualHost = *response.Response.VirtualHost
	d.SetId(strings.Join([]string{instanceId, virtualHost}, FILED_SP))

	return resourceTencentCloudTdmqRabbitmqVirtualHostRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqVirtualHostRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_virtual_host.read")()
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
	virtualHost := idSplit[1]

	rabbitmqVirtualHost, err := service.DescribeTdmqRabbitmqVirtualHostById(ctx, instanceId, virtualHost)
	if err != nil {
		return err
	}

	if rabbitmqVirtualHost == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TdmqRabbitmqVirtualHost` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if rabbitmqVirtualHost.InstanceId != nil {
		_ = d.Set("instance_id", rabbitmqVirtualHost.InstanceId)
	}

	if rabbitmqVirtualHost.VirtualHost != nil {
		_ = d.Set("virtual_host", rabbitmqVirtualHost.VirtualHost)
	}

	if rabbitmqVirtualHost.Description != nil {
		_ = d.Set("description", rabbitmqVirtualHost.Description)
	}

	if rabbitmqVirtualHost.TraceFlag != nil {
		_ = d.Set("trace_flag", rabbitmqVirtualHost.TraceFlag)
	}

	return nil
}

func resourceTencentCloudTdmqRabbitmqVirtualHostUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_virtual_host.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = tdmq.NewModifyRabbitMQVirtualHostRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	instanceId := idSplit[0]
	virtualHost := idSplit[1]

	immutableArgs := []string{"instance_id", "virtual_host", "mirror_queue_policy_flag"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("description") {
		request.InstanceId = &instanceId
		request.VirtualHost = &virtualHost

		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().ModifyRabbitMQVirtualHost(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update tdmq rabbitmqVirtualHost failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudTdmqRabbitmqVirtualHostRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqVirtualHostDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_virtual_host.delete")()
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
	virtualHost := idSplit[1]

	if err := service.DeleteTdmqRabbitmqVirtualHostById(ctx, instanceId, virtualHost); err != nil {
		return err
	}

	return nil
}

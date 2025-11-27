/*
Provides a resource to create a clear_instance_operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_redis_clear_instance_operation" "clear_instance_operation" {
	  instance_id = "crs-c1nl9rpv"
	  password 	  = ""
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_redis_clear_instance_operation", CNDescription{
		TerraformTypeCN: "redis®清空实例数据",
		DescriptionCN:   "提供Redis®清空实例数据资源，用于清空Redis®实例的所有数据。",
		AttributesCN: map[string]string{
			"instance_id": "实例id",
			"password":    "实例密码",
		},
	})
}
func resourceTencentCloudRedisClearInstanceOperation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a clear_instance_operation",
		Create:      resourceTencentCloudRedisClearInstanceOperationCreate,
		Read:        resourceTencentCloudRedisClearInstanceOperationRead,
		Delete:      resourceTencentCloudRedisClearInstanceOperationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"password": {
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
				Description: "Redis® instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).",
			},
		},
	}
}

func resourceTencentCloudRedisClearInstanceOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_clear_instance_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request    = redis.NewClearInstanceRequest()
		response   = redis.NewClearInstanceResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("password"); ok {
		request.Password = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseRedisClient().ClearInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate redis clearInstanceOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}
	taskId := *response.Response.TaskId
	err = resource.Retry(6*readRetryTimeout, func() *resource.RetryError {
		ok, err := service.DescribeTaskInfo(ctx, instanceId, taskId)
		if err != nil {
			if _, ok := err.(*sdkErrors.CloudSDKError); !ok {
				return resource.RetryableError(err)
			} else {
				return resource.NonRetryableError(err)
			}
		}
		if ok {
			return nil
		} else {
			return resource.RetryableError(fmt.Errorf("clear instance is processing"))
		}
	})

	if err != nil {
		log.Printf("[CRITAL]%s redis clear instance fail, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudRedisClearInstanceOperationRead(d, meta)
}

func resourceTencentCloudRedisClearInstanceOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_clear_instance_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudRedisClearInstanceOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_clear_instance_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

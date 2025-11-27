/*
Provides a resource to create a startup_instance_operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_redis_startup_instance_operation" "startup_instance_operation" {
	  instance_id = "crs-c1nl9rpv"
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
	registerResourceDescriptionProvider("tencentcloudenterprise_redis_startup_instance_operation", CNDescription{
		TerraformTypeCN: "redis®启动实例",
		DescriptionCN:   "提供Redis®启动实例资源，用于启动已停止的Redis®实例。",
		AttributesCN: map[string]string{
			"instance_id": "实例id",
		},
	})
}
func resourceTencentCloudRedisStartupInstanceOperation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a startup_instance_operation",
		Create:      resourceTencentCloudRedisStartupInstanceOperationCreate,
		Read:        resourceTencentCloudRedisStartupInstanceOperationRead,
		Delete:      resourceTencentCloudRedisStartupInstanceOperationDelete,
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
		},
	}
}

func resourceTencentCloudRedisStartupInstanceOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_startup_instance_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request    = redis.NewStartupInstanceRequest()
		response   = redis.NewStartupInstanceResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseRedisClient().StartupInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate redis startupInstance failed, reason:%+v", logId, err)
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
			return resource.RetryableError(fmt.Errorf("startup instance is processing"))
		}
	})

	if err != nil {
		log.Printf("[CRITAL]%s redis startup instance fail, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudRedisStartupInstanceOperationRead(d, meta)
}

func resourceTencentCloudRedisStartupInstanceOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_startup_instance_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudRedisStartupInstanceOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_startup_instance_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

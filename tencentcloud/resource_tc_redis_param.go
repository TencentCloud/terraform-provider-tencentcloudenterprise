/*
Provides a resource to create a param

# Example Usage

```hcl

	resource "tencentcloudenterprise_redis_param" "param" {
	    instance_id     = "crs-c1nl9rpv"
	    instance_params = {
	        "cluster-node-timeout"          = "15000"
	        "disable-command-list"          = "\"\""
	        "hash-max-ziplist-entries"      = "512"
	        "hash-max-ziplist-value"        = "64"
	        "hz"                            = "10"
	        "lazyfree-lazy-eviction"        = "yes"
	        "lazyfree-lazy-expire"          = "yes"
	        "lazyfree-lazy-server-del"      = "yes"
	        "maxmemory-policy"              = "noeviction"
	        "notify-keyspace-events"        = "\"\""
	        "proxy-slowlog-log-slower-than" = "500"
	        "replica-lazy-flush"            = "yes"
	        "sentineauth"                   = "no"
	        "set-max-intset-entries"        = "512"
	        "slowlog-log-slower-than"       = "10"
	        "timeout"                       = "31536000"
	        "zset-max-ziplist-entries"      = "128"
	        "zset-max-ziplist-value"        = "64"
	    }
	}

```

# Import

redis param can be imported using the instanceId, e.g.

```
terraform import cloud_redis_param.param crs-c1nl9rpv
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
	registerResourceDescriptionProvider("tencentcloudenterprise_redis_param", CNDescription{
		TerraformTypeCN: "redis®参数",
		DescriptionCN:   "提供Redis®参数资源，用于修改Redis®实例的配置参数。",
		AttributesCN: map[string]string{
			"instance_id":     "实例id",
			"instance_params": "实例参数",
		},
	})
}
func resourceTencentCloudRedisParam() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a param",
		Create:      resourceTencentCloudRedisParamCreate,
		Read:        resourceTencentCloudRedisParamRead,
		Update:      resourceTencentCloudRedisParamUpdate,
		Delete:      resourceTencentCloudRedisParamDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"instance_params": {
				Required:    true,
				Type:        schema.TypeMap,
				Description: "A list of parameters modified by the instance.",
			},
		},
	}
}

func resourceTencentCloudRedisParamCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_param.create")()
	defer inconsistentCheck(d, meta)()

	var (
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
	}

	d.SetId(instanceId)

	return resourceTencentCloudRedisParamUpdate(d, meta)
}

func resourceTencentCloudRedisParamRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_param.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	param, err := service.DescribeRedisParamById(ctx, instanceId)
	if err != nil {
		return err
	}

	if len(param) == 0 {
		d.SetId("")
		log.Printf("[WARN]%s resource `RedisParam` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)

	instanceParamsMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_params"); ok {
		for k := range v.(map[string]interface{}) {
			instanceParamsMap[k] = param[k]
		}
	} else {
		instanceParamsMap = param
	}
	_ = d.Set("instance_params", instanceParamsMap)

	return nil
}

func resourceTencentCloudRedisParamUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_param.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	request := redis.NewModifyInstanceParamsRequest()
	response := redis.NewModifyInstanceParamsResponse()

	instanceId := d.Id()
	request.InstanceId = &instanceId

	if v, ok := d.GetOk("instance_params"); ok {
		service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}
		param, err := service.DescribeRedisParamById(ctx, instanceId)
		if err != nil && len(param) == 0 {
			return fmt.Errorf("[ERROR] resource `RedisParam` [%s] not found, please check if it has been deleted.\n", d.Id())
		}
		for k, v := range v.(map[string]interface{}) {
			if value, ok := param[k]; ok {
				if value != v {
					instanceParam := redis.InstanceParam{}
					instanceParam.Key = helper.String(k)
					instanceParam.Value = helper.String(v.(string))
					request.InstanceParams = append(request.InstanceParams, &instanceParam)
				}
			} else {
				return fmt.Errorf("[ERROR] The parameter name [%v] does not exist, please check the parameter name.\n", k)
			}
		}
	}

	if len(request.InstanceParams) == 0 {
		return resourceTencentCloudRedisParamRead(d, meta)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseRedisClient().ModifyInstanceParams(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update redis param failed, reason:%+v", logId, err)
		return err
	}

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
			return resource.RetryableError(fmt.Errorf("change param is processing"))
		}
	})

	if err != nil {
		log.Printf("[CRITAL]%s redis change param fail, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudRedisParamRead(d, meta)
}

func resourceTencentCloudRedisParamDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_redis_param.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

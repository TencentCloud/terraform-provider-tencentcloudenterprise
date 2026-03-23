/*
Provides a resource to create a apm instance

Example Usage

```hcl
resource "tencentcloudenterprise_apm_instance" "instance" {
  name           = "terraform-test"
  description    = "for terraform test"
  trace_duration = 15
}

output "instance_id" {
  value = tencentcloudenterprise_apm_instance.instance.instance_id
}
```

Import

apm instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_apm_instance.instance instance_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	apm "terraform-provider-tencentcloudenterprise/sdk/apm/v20210622"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_apm_instance", CNDescription{
		TerraformTypeCN: "APM 业务系统",
		DescriptionCN:   "提供 APM（应用性能监控）业务系统资源，用于创建和管理 APM 实例。",
		AttributesCN: map[string]string{
			"name":           "业务系统名称,支持长度小于40的中文、英文、数字以及分隔符(\".\",\"_\",\"-\")",
			"description":    "业务系统描述信息,请在100个字符以内进行描述",
			"trace_duration": "Trace 数据保存时长,单位为天,取值范围:1-30,默认值:7",
			"instance_id":    "业务系统实例 ID",
		},
	})
}

func resourceTencentCloudApmInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudApmInstanceCreate,
		Read:   resourceTencentCloudApmInstanceRead,
		Update: resourceTencentCloudApmInstanceUpdate,
		Delete: resourceTencentCloudApmInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
		"name": {
			Required:    true,
			Type:        schema.TypeString,
			Description: "Name Of Instance.",
			ValidateFunc: validateApmInstanceName,
		},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description Of Instance.",
				ValidateFunc: validateStringLengthInRange(0, 100),
			},

			"trace_duration": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      7,
				Description:  "Duration of trace data retention in days. Valid values range from 1 to 30. Default is 7.",
				ValidateFunc: validateIntegerInRange(1, 30),
			},

			"instance_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},
		},
	}
}

func resourceTencentCloudApmInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_apm_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = apm.NewCreateApmInstanceRequest()
		response   = apm.NewCreateApmInstanceResponse()
		instanceId string
	)
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("trace_duration"); ok {
		request.TraceDuration = helper.IntInt64(v.(int))
	}

	// Set default values for fields not exposed in schema
	request.SpanDailyCounters = helper.IntUint64(0)
	request.PayMode = helper.IntInt64(0)
	request.Free = helper.IntInt64(0)
	request.Tags = []*apm.ApmTag{}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseApmClient().CreateApmInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create apm instance failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId
	d.SetId(instanceId)

	return resourceTencentCloudApmInstanceRead(d, meta)
}

func resourceTencentCloudApmInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_apm_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ApmService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	instance, err := service.DescribeApmInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if instance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ApmInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if instance.Name != nil {
		_ = d.Set("name", instance.Name)
	}

	if instance.Description != nil {
		_ = d.Set("description", instance.Description)
	}

	if instance.TraceDuration != nil {
		_ = d.Set("trace_duration", instance.TraceDuration)
	}

	if instance.InstanceId != nil {
		_ = d.Set("instance_id", instance.InstanceId)
	}

	return nil
}

func resourceTencentCloudApmInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_apm_instance.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := apm.NewModifyApmInstanceRequest()

	needChange := false

	instanceId := d.Id()

	request.InstanceId = &instanceId

	mutableArgs := []string{"name", "description", "trace_duration"}

	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {

		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}

		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("trace_duration"); ok {
			request.TraceDuration = helper.IntInt64(v.(int))
		}

		// Set default values for fields not exposed in schema
		request.SpanDailyCounters = helper.IntUint64(0)
		request.PayMode = helper.IntInt64(0)
		request.Free = helper.IntInt64(0)
		request.Tags = []*apm.ApmTag{}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseApmClient().ModifyApmInstance(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update apm instance failed, reason:%+v", logId, err)
			return err
		}

	}

	return resourceTencentCloudApmInstanceRead(d, meta)
}

func resourceTencentCloudApmInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_apm_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ApmService{client: meta.(*TencentCloudClient).apiV3Conn}
	instanceId := d.Id()

	if err := service.DeleteApmInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}

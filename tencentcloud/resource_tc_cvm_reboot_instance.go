/*
Provides a resource to create a cvm reboot_instance

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_reboot_instance" "reboot_instance" {
	  instance_id = "ins-xxxxx"
	  stop_type = "SOFT"
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_reboot_instance", CNDescription{
		TerraformTypeCN: "重启实例",
		DescriptionCN:   "提供云服务器重启实例资源，用于重启云服务器实例。",
		AttributesCN: map[string]string{
			"instance_id": "实例ID",
			"stop_type":   "关机类型。取值范围：SOFT：表示软关机HARD：表示硬关机SOFT_FIRST：表示优先软关机，失败再执行硬关机默认取值：SOFT",
		},
	})
}
func resourceTencentCloudCvmRebootInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cvm reboot_instance",
		Create:      resourceTencentCloudCvmRebootInstanceCreate,
		Read:        resourceTencentCloudCvmRebootInstanceRead,
		Delete:      resourceTencentCloudCvmRebootInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},

			"stop_type": {
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeString,
				Default:  CVM_REBOOT_TYPE_SOFT,
				Description: "Shutdown type. Valid values: `SOFT`: soft shutdown; `HARD`: hard shutdown; `SOFT_FIRST" +
					"`: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails. Default value: SOFT.",
				ValidateFunc: validateAllowedStringValue(CVM_REBOOT_TYPE),
			},
		},
	}
}

func resourceTencentCloudCvmRebootInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_reboot_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cvm.NewRebootInstancesRequest()
	instanceId := d.Get("instance_id").(string)
	request.InstanceIds = []*string{&instanceId}

	if v, ok := d.GetOk("stop_type"); ok {
		request.StopType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().RebootInstances(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate cvm rebootInstance failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	return resourceTencentCloudCvmRebootInstanceRead(d, meta)
}

func resourceTencentCloudCvmRebootInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_reboot_instance.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCvmRebootInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_reboot_instance.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

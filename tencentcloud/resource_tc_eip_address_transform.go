/*
Provides a resource to create a eip address_transform

# Example Usage

```hcl

	resource "tencentcloudenterprise_eip_address_transform" "address_transform" {
	  instance_id = ""
	}

```

# Import

eip address_transform can be imported using the id, e.g.

```
terraform import cloud_eip_address_transform.address_transform address_transform_id
```
*/
package tencentcloud

import (
	"log"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_eip_address_transform", CNDescription{
		TerraformTypeCN: "普通公网转换为弹性公网",
		DescriptionCN:   "提供普通公网IP转换为弹性公网IP资源，用于将普通公网IP转换为弹性公网IP。",
		AttributesCN: map[string]string{
			"instance_id": "待操作有普通公网 IP 的实例 ID",
		},
	})
}

func resourceTencentCloudEipAddressTransform() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a eip address_transform.",
		Create:      resourceTencentCloudEipAddressTransformCreate,
		Read:        resourceTencentCloudEipAddressTransformRead,
		Delete:      resourceTencentCloudEipAddressTransformDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The instance ID of a normal public network IP to be operated. eg:ins-23mk45jn.",
			},
		},
	}
}

func resourceTencentCloudEipAddressTransformCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_address_transform.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = vpc.NewTransformAddressRequest()
		response   = vpc.NewTransformAddressResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().TransformAddress(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate eip addressTransform failed, reason:%+v", logId, err)
		return err
	}

	taskId := *response.Response.TaskId
	d.SetId(instanceId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"SUCCESS"}, 1*readRetryTimeout, time.Second, service.VpcIpv6AddressStateRefreshFunc(helper.UInt64ToStr(taskId), []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	return resourceTencentCloudEipAddressTransformRead(d, meta)
}

func resourceTencentCloudEipAddressTransformRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_address_transform.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudEipAddressTransformDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_address_transform.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

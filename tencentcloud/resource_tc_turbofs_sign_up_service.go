/*
Provides a resource to create a turbofs sign up service

# Example Usage

```hcl
resource "tencentcloudenterprise_turbofs_sign_up_service" "sign_up" {}
```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_sign_up_service", CNDescription{
		TerraformTypeCN: "开通TurboFS服务",
		DescriptionCN:   "提供TurboFS服务开通资源，用于开通TurboFS服务。",
		AttributesCN: map[string]string{
			"cfs_service_status": "TurboFS服务状态，包含creating(创建中)和created(创建成功)",
		},
	})
}
func resourceTencentCloudTurbofsSignUpService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a TurboFS sign_up_service.",
		Create:      resourceTencentCloudTurbofsSignUpServiceCreate,
		Read:        resourceTencentCloudTurbofsSignUpServiceRead,
		Delete:      resourceTencentCloudTurbofsSignUpServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cfs_service_status": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Current status of the TurboFS service for this user. Valid values: creating (activating); created (activated).",
			},
		},
	}
}

func resourceTencentCloudTurbofsSignUpServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_turbofs_sign_up_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = turbofs.NewSignUpCfsServiceRequest()
		response             = turbofs.NewSignUpCfsServiceResponse()
		turbofsServiceStatus string
	)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().SignUpCfsService(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate turbofs signUpService failed, reason:%+v", logId, err)
		return nil
	}

	turbofsServiceStatus = *response.Response.CfsServiceStatus
	d.SetId(turbofsServiceStatus)

	return resourceTencentCloudTurbofsSignUpServiceRead(d, meta)
}

func resourceTencentCloudTurbofsSignUpServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_sign_up_service.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = turbofs.NewDescribeCfsServiceStatusRequest()
		response = turbofs.NewDescribeCfsServiceStatusResponse()
	)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().DescribeCfsServiceStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate turbofs signUpService failed, reason:%+v", logId, err)
		return nil
	}

	_ = d.Set("cfs_service_status", response.Response.CfsServiceStatus)

	return nil
}

func resourceTencentCloudTurbofsSignUpServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_sign_up_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

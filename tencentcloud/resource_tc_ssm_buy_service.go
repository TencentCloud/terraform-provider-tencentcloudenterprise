/*
Provides a resource to create a ssm buy_service

Example Usage

```hcl
resource "tencentcloudenterprise_ssm_buy_service" "buy_service" {}
```
*/
package tencentcloud

import (
	"log"
	"strings"

	ssm "terraform-provider-tencentcloudenterprise/sdk/ssm/v20190923"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ssm_buy_service", CNDescription{
		TerraformTypeCN: "开通SSM服务",
		DescriptionCN:   "提供SSM服务开通资源，用于开通凭据管理系统（SSM）服务。",
		AttributesCN: map[string]string{
			"service_enabled": "服务是否已开通，true表示已开通，false表示尚未开通",
			"invalid_type":    "服务不可用类型：0-未购买，1-正常，2-欠费停服，3-资源释放",
		},
	})
}
func resourceTencentCloudSsmBuyService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a SSM buy_service.",
		Create:      resourceTencentCloudSsmBuyServiceCreate,
		Read:        resourceTencentCloudSsmBuyServiceRead,
		Delete:      resourceTencentCloudSsmBuyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"service_enabled": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether the SSM service is enabled. true: enabled; false: not enabled.",
			},
			"invalid_type": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Service unavailable type. 0: not purchased; 1: normal; 2: arrears; 3: resource released.",
			},
		},
	}
}

func resourceTencentCloudSsmBuyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ssm_buy_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = ssm.NewBuyServiceRequest()
	)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseSsmClient().BuyService(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok {
				msg := sdkErr.GetMessage()
				if strings.Contains(msg, "-100") || strings.Contains(msg, "已开通") {
					log.Printf("[WARN]%s ssm buy service already activated, treat as success: %s", logId, e.Error())
					return nil
				}
			}
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate ssm buyService failed, reason:%+v", logId, err)
		return err
	}

	d.SetId("ssm_service")

	return resourceTencentCloudSsmBuyServiceRead(d, meta)
}

func resourceTencentCloudSsmBuyServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ssm_buy_service.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = ssm.NewGetServiceStatusRequest()
		response = ssm.NewGetServiceStatusResponse()
	)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseSsmClient().GetServiceStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate ssm getServiceStatus failed, reason:%+v", logId, err)
		return err
	}

	if response.Response.ServiceEnabled != nil {
		_ = d.Set("service_enabled", *response.Response.ServiceEnabled)
		if !*response.Response.ServiceEnabled {
			log.Printf("[WARN]%s resource `tencentcloudenterprise_ssm_buy_service` [%s] not found, SSM service is not enabled.\n", logId, d.Id())
			d.SetId("")
			return nil
		}
	}
	if response.Response.InvalidType != nil {
		_ = d.Set("invalid_type", *response.Response.InvalidType)
	}

	return nil
}

func resourceTencentCloudSsmBuyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ssm_buy_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

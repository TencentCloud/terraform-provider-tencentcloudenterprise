/*
Provides a resource to create a kms buy_service

Example Usage

```hcl
resource "tencentcloudenterprise_kms_buy_service" "buy_service" {}
```
*/
package tencentcloud

import (
	"log"
	"strings"

	kms "terraform-provider-tencentcloudenterprise/sdk/kms/v20190118"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_kms_buy_service", CNDescription{
		TerraformTypeCN: "开通KMS服务",
		DescriptionCN:   "提供KMS服务开通资源，用于开通密钥管理系统（KMS）服务。",
		AttributesCN: map[string]string{
			"service_enabled": "服务是否已开通，true表示已开通，false表示尚未开通",
			"invalid_type":    "服务不可用类型：0-未购买，1-正常，2-欠费停服，3-资源释放",
		},
	})
}
func resourceTencentCloudKmsBuyService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a KMS buy_service.",
		Create:      resourceTencentCloudKmsBuyServiceCreate,
		Read:        resourceTencentCloudKmsBuyServiceRead,
		Delete:      resourceTencentCloudKmsBuyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"service_enabled": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether the KMS service is enabled. true: enabled; false: not enabled.",
			},
			"invalid_type": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Service unavailable type. 0: not purchased; 1: normal; 2: arrears; 3: resource released.",
			},
		},
	}
}

func resourceTencentCloudKmsBuyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kms_buy_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = kms.NewBuyServiceRequest()
	)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseKmsClient().BuyService(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok {
				msg := sdkErr.GetMessage()
				if strings.Contains(msg, "-100") || strings.Contains(msg, "已开通") {
					log.Printf("[WARN]%s kms buy service already activated, treat as success: %s", logId, e.Error())
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
		log.Printf("[CRITAL]%s operate kms buyService failed, reason:%+v", logId, err)
		return err
	}

	d.SetId("kms_service")

	return resourceTencentCloudKmsBuyServiceRead(d, meta)
}

func resourceTencentCloudKmsBuyServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kms_buy_service.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = kms.NewGetServiceStatusRequest()
		response = kms.NewGetServiceStatusResponse()
	)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseKmsClient().GetServiceStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate kms getServiceStatus failed, reason:%+v", logId, err)
		return err
	}

	if response.Response.ServiceEnabled != nil {
		_ = d.Set("service_enabled", *response.Response.ServiceEnabled)
		if !*response.Response.ServiceEnabled {
			log.Printf("[WARN]%s resource `tencentcloudenterprise_kms_buy_service` [%s] not found, KMS service is not enabled.\n", logId, d.Id())
			d.SetId("")
			return nil
		}
	}
	if response.Response.InvalidType != nil {
		_ = d.Set("invalid_type", *response.Response.InvalidType)
	}

	return nil
}

func resourceTencentCloudKmsBuyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kms_buy_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

/*
Provides a resource to create a csp open_cos_billing

Example Usage

```hcl
resource "tencentcloudenterprise_csp_open_cos_billing" "example" {
  cos_region = "ap-guangzhou"
}
```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strings"

	cossdk "terraform-provider-tencentcloudenterprise/sdk/cos/v20200107"
	csp "terraform-provider-tencentcloudenterprise/sdk/csp/v20200107"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_csp_open_cos_billing", CNDescription{
		TerraformTypeCN: "开通COS计费",
		DescriptionCN:   "提供COS计费开通资源，用于开通对象存储（COS）的计费服务。",
		AttributesCN: map[string]string{
			"cos_region": "存储桶所在的COS地域",
		},
	})
}

func resourceTencentCloudCspOpenCosBilling() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a CSP open_cos_billing.",
		Create:      resourceTencentCloudCspOpenCosBillingCreate,
		Read:        resourceTencentCloudCspOpenCosBillingRead,
		Delete:      resourceTencentCloudCspOpenCosBillingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cos_region": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The COS region where the bucket resides.",
			},
		},
	}
}

func resourceTencentCloudCspOpenCosBillingCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_open_cos_billing.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := csp.NewOpenCosBillingRequest()
	if v, ok := d.GetOk("cos_region"); ok {
		request.CosRegion = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCspClient().OpenCosBilling(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok {
				msg := sdkErr.GetMessage()
				if strings.Contains(msg, "-100") || strings.Contains(msg, "save resource error") {
					log.Printf("[WARN]%s csp OpenCosBilling already activated, treat as success: %s", logId, e.Error())
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
		log.Printf("[CRITAL]%s operate csp OpenCosBilling failed, reason:%+v", logId, err)
		return err
	}

	d.SetId("csp_cos_billing")

	return resourceTencentCloudCspOpenCosBillingRead(d, meta)
}

func resourceTencentCloudCspOpenCosBillingRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_open_cos_billing.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cossdk.NewGetBillingTypeRequest()
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCosSdkClient().GetBillingType(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, response body [%s]\n", logId, request.GetAction(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("GetBillingType response is nil."))
		}

		if result.Response.IsBilling != nil && !*result.Response.IsBilling {
			log.Printf("[WARN]%s resource `tencentcloudenterprise_csp_open_cos_billing` [%s] not found, COS billing is not activated.\n", logId, d.Id())
			d.SetId("")
		}

		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate cos GetBillingType failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

func resourceTencentCloudCspOpenCosBillingDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_open_cos_billing.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

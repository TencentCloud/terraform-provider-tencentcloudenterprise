/*
Provide a resource to create a BH access white list config

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_access_white_list_config" "example" {
	  allow_any  = false
	  allow_auto = true
	}

```

# Import

BH access white list config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_access_white_list_config.example id
```
*/
package tencentcloud

import (
	"context"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_access_white_list_config", CNDescription{
		TerraformTypeCN: "BH 访问白名单配置",
		DescriptionCN:   "提供 BH 访问白名单配置资源，用于管理堡垒机访问白名单配置。",
		AttributesCN: map[string]string{
			"allow_any":  "是否放通所有来源IP",
			"allow_auto": "是否放通自动添加的IP",
		},
	})
}

func ResourceTencentCloudBhAccessWhiteListConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhAccessWhiteListConfigCreate,
		Read:   resourceTencentCloudBhAccessWhiteListConfigRead,
		Update: resourceTencentCloudBhAccessWhiteListConfigUpdate,
		Delete: resourceTencentCloudBhAccessWhiteListConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"allow_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "true: allow all source IPs; false: do not allow all source IPs.",
			},

			"allow_auto": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "true: allow automatically added IPs; false: do not allow automatically added IPs.",
			},
		},
	}
}

func resourceTencentCloudBhAccessWhiteListConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_access_white_list_config.create")()
	defer inconsistentCheck(d, meta)()

	d.SetId(helper.BuildToken())

	return resourceTencentCloudBhAccessWhiteListConfigUpdate(d, meta)
}

func resourceTencentCloudBhAccessWhiteListConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_access_white_list_config.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = BhService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	respData, err := service.DescribeBhAccessWhiteListConfigById(ctx)
	if err != nil {
		return err
	}

	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_bh_access_white_list_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if respData.Response.AllowAny != nil {
		_ = d.Set("allow_any", respData.Response.AllowAny)
	}

	if respData.Response.AllowAuto != nil {
		_ = d.Set("allow_auto", respData.Response.AllowAuto)
	}

	return nil
}

func resourceTencentCloudBhAccessWhiteListConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_access_white_list_config.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId = getLogId(contextNil)
	)

	if d.HasChange("allow_any") {
		request := bhsaas.NewModifyAccessWhiteListStatusRequest()
		if v, ok := d.GetOkExists("allow_any"); ok {
			request.AllowAny = helper.Bool(v.(bool))
		}

		reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyAccessWhiteListStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s update bh access white list config failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	if d.HasChange("allow_auto") {
		request := bhsaas.NewModifyAccessWhiteListAutoStatusRequest()
		if v, ok := d.GetOkExists("allow_auto"); ok {
			request.AllowAuto = helper.Bool(v.(bool))
		}

		reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyAccessWhiteListAutoStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s update bh access white list config failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	return resourceTencentCloudBhAccessWhiteListConfigRead(d, meta)
}

func resourceTencentCloudBhAccessWhiteListConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_access_white_list_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

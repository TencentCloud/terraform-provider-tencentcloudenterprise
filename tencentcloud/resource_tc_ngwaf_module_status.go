/*
Provides a resource to manage NGWAF module status for a domain.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_module_status" "example" {
  domain         = "example.com"
  web_security   = 1
  access_control = 1
  cc_protection  = 1
  api_protection = 0
  anti_tamper    = 0
  anti_leakage   = 0
}
```

Import

NGWAF module status can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_module_status.example example.com
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafModuleStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafModuleStatusCreate,
		Read:   resourceTencentCloudNgwafModuleStatusRead,
		Update: resourceTencentCloudNgwafModuleStatusUpdate,
		Delete: resourceTencentCloudNgwafModuleStatusDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"web_security": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "WEB security module status, 0:closed, 1:opened.",
			},
			"access_control": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "ACL module status, 0:closed, 1:opened.",
			},
			"cc_protection": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "CC module status, 0:closed, 1:opened.",
			},
			"api_protection": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "API security module status, 0:closed, 1:opened.",
			},
			"anti_tamper": {
				Optional:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "Anti tamper module status, 0:closed, 1:opened.",
			},
			"anti_leakage": {
				Optional:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "Anti leakage module status, 0:closed, 1:opened.",
			},
		},
	}
}

func resourceTencentCloudNgwafModuleStatusCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_module_status.create")()
	defer inconsistentCheck(d, meta)()

	var domain string
	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	d.SetId(domain)

	return resourceTencentCloudNgwafModuleStatusUpdate(d, meta)
}

func resourceTencentCloudNgwafModuleStatusRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_module_status.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		domain  = d.Id()
	)

	moduleStatus, err := service.DescribeWafModuleStatusById(ctx, domain)
	if err != nil {
		return err
	}

	if moduleStatus == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafModuleStatus` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)

	if moduleStatus.Response.WebSecurity != nil {
		_ = d.Set("web_security", moduleStatus.Response.WebSecurity)
	}

	if moduleStatus.Response.AccessControl != nil {
		_ = d.Set("access_control", moduleStatus.Response.AccessControl)
	}

	if moduleStatus.Response.CcProtection != nil {
		_ = d.Set("cc_protection", moduleStatus.Response.CcProtection)
	}

	if moduleStatus.Response.ApiProtection != nil {
		_ = d.Set("api_protection", moduleStatus.Response.ApiProtection)
	}

	if moduleStatus.Response.AntiTamper != nil {
		_ = d.Set("anti_tamper", moduleStatus.Response.AntiTamper)
	}

	if moduleStatus.Response.AntiLeakage != nil {
		_ = d.Set("anti_leakage", moduleStatus.Response.AntiLeakage)
	}

	return nil
}

func resourceTencentCloudNgwafModuleStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_module_status.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyModuleStatusRequest()
		domain  = d.Id()
	)

	request.Domain = &domain

	if v, ok := d.GetOkExists("web_security"); ok {
		request.WebSecurity = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("access_control"); ok {
		request.AccessControl = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("cc_protection"); ok {
		request.CcProtection = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("api_protection"); ok {
		request.ApiProtection = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("anti_tamper"); ok {
		request.AntiTamper = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("anti_leakage"); ok {
		request.AntiLeakage = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyModuleStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil {
			e = fmt.Errorf("waf moduleStatus version not exists")
			return resource.NonRetryableError(e)
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s modify waf moduleStatus failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafModuleStatusRead(d, meta)
}

func resourceTencentCloudNgwafModuleStatusDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_module_status.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

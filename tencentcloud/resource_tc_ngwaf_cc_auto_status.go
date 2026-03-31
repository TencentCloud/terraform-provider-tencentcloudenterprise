/*
Provides a resource to manage NGWAF CC auto status.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ngwaf_cc_auto_status" "example" {
	  domain  = "keep.qcloudwaf.com"
	  edition = "sparta-waf"
	}

```

# Import

NGWAF CC auto status can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_cc_auto_status.example keep.qcloudwaf.com#sparta-waf
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ngwaf_cc_auto_status", CNDescription{
		TerraformTypeCN: "NGWAF CC 自动状态配置",
		DescriptionCN:   "提供 NGWAF CC 自动状态配置资源，用于开启或关闭指定域名的 CC 自动防护开关。",
		AttributesCN: map[string]string{
			"domain":  "域名",
			"edition": "WAF 版本",
			"status":  "CC 自动防护开关状态，1 表示开启，0 表示关闭",
		},
	})
}

func resourceTencentCloudNgwafCcAutoStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafCcAutoStatusCreate,
		Read:   resourceTencentCloudNgwafCcAutoStatusRead,
		Delete: resourceTencentCloudNgwafCcAutoStatusDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain.",
			},
			"edition": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "CC auto status, 1 means enabled and 0 means disabled.",
			},
		},
	}
}

func resourceTencentCloudNgwafCcAutoStatusCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_auto_status.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewUpsertCCAutoStatusRequest()
		domain  string
		edition string
	)

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
		request.Domain = helper.String(domain)
	}

	if v, ok := d.GetOk("edition"); ok {
		edition = v.(string)
		request.Edition = helper.String(edition)
	}

	request.Value = helper.Int64(1)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().UpsertCCAutoStatus(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create ngwaf cc auto status failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{domain, edition}, FILED_SP))

	return resourceTencentCloudNgwafCcAutoStatusRead(d, meta)
}

func resourceTencentCloudNgwafCcAutoStatusRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_auto_status.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.Background(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	edition := idSplit[1]

	respData, err := service.DescribeWafCcAutoStatusById(ctx, domain)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ngwaf_cc_auto_status` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)
	_ = d.Set("edition", edition)

	if respData.Response != nil && respData.Response.AutoCCSwitch != nil {
		_ = d.Set("status", *respData.Response.AutoCCSwitch)
	}

	return nil
}

func resourceTencentCloudNgwafCcAutoStatusDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_auto_status.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.Background(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	edition := idSplit[1]

	if err := service.DeleteWafCcAutoStatusById(ctx, domain, edition); err != nil {
		return err
	}

	return nil
}

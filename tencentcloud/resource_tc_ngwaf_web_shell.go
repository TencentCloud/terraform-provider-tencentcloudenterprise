package tencentcloud

import (
	"context"
	"log"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafWebShell() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafWebShellCreate,
		Read:   resourceTencentCloudNgwafWebShellRead,
		Update: resourceTencentCloudNgwafWebShellUpdate,
		Delete: resourceTencentCloudNgwafWebShellDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain.",
			},
			"status": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateAllowedIntValue([]int{0, 1, 2}),
				Description:  "Webshell status, 1: open; 0: closed; 2: log.",
			},
		},
	}
}

func resourceTencentCloudNgwafWebShellCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_web_shell.create")()
	defer inconsistentCheck(d, meta)()

	var domain string

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	d.SetId(domain)

	return resourceTencentCloudNgwafWebShellUpdate(d, meta)
}

func resourceTencentCloudNgwafWebShellRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_web_shell.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		domain  = d.Id()
	)

	webShell, err := service.DescribeWafWebShellById(ctx, domain)
	if err != nil {
		return err
	}

	if webShell == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafWebShell` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if webShell.Response.Domain != nil {
		_ = d.Set("domain", webShell.Response.Domain)
	}

	if webShell.Response.Status != nil {
		_ = d.Set("status", webShell.Response.Status)
	}

	return nil
}

func resourceTencentCloudNgwafWebShellUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_web_shell.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyWebshellStatusRequest()
		domain  = d.Id()
	)

	webShellStatus := ngwaf.WebshellStatus{}
	webShellStatus.Domain = helper.String(domain)

	if v, ok := d.GetOkExists("status"); ok {
		webShellStatus.Status = helper.IntUint64(v.(int))
	}

	request.Webshell = &webShellStatus

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyWebshellStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf webShell failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafWebShellRead(d, meta)
}

func resourceTencentCloudNgwafWebShellDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_web_shell.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

/*
Provides a resource to create a NGWAF bot resource.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_bot_resource" "example" {
  instance_id = "waf-xxxxxxxx"
  region_id   = "1"
  pay_mode    = 1
  pid         = 1000
  project_id  = 0
  goods_num   = 1
}
```

Import

NGWAF bot resource can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_bot_resource.example waf-xxxxxxxx
```
*/
package tencentcloud

import (
	"fmt"
	"log"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafBotResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafBotResourceCreate,
		Read: resourceTencentCloudNgwafBotResourceRead,
		Delete: resourceTencentCloudNgwafBotResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "WAF instance ID.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Region ID.",
			},
			"pay_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Payment mode, 0: post-pay, 1: pre-pay.",
			},
			"pid": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Product ID.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Project ID.",
			},
			"goods_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				ForceNew:    true,
				Description: "Goods number.",
			},
		},
	}
}

func resourceTencentCloudNgwafBotResourceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_resource.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewWafCreateBotResourceAfterPayRequest()
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("region_id"); ok {
		request.RegionId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("pay_mode"); ok {
		request.PayMode = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("pid"); ok {
		request.Pid = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("goods_num"); ok {
		request.GoodsNum = helper.Int64(int64(v.(int)))
	}

	// Set default interface name for bot resource creation
	request.InterfaceName = helper.String("qcloud.Hour.create")

	var result *ngwaf.WafCreateBotResourceAfterPayResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().WafCreateBotResourceAfterPay(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", 
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			result = response
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf bot resource failed, reason:%+v", logId, err)
		return err
	}

	if result == nil || result.Response == nil {
		return fmt.Errorf("create waf bot resource response is nil")
	}

	// Use instance_id as the resource ID since the bot resource is tied to the instance
	instanceId := d.Get("instance_id").(string)
	d.SetId(instanceId)

	return nil
}

func resourceTencentCloudNgwafBotResourceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_resource.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudNgwafBotResourceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_resource.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewDeleteTceWafBotResourceRequest()
	)

	instanceId := d.Id()
	request.InstanceId = &instanceId

	if v, ok := d.GetOk("region_id"); ok {
		request.RegionId = helper.String(v.(string))
	}

	// Set default interface name for bot resource deletion
	request.InterfaceName = helper.String("qcloud.Hour.unblock")

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DeleteTceWafBotResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", 
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s delete waf bot resource failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
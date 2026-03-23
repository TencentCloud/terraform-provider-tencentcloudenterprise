/*
Provides a resource to create a NGWAF instance.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_instance" "clb_waf" {
  instance_name  = "production-clb-waf"
  region_id      = "50000001"
  pay_mode       = 0
  pid            = 11543
  type           = 1
  project_id     = 0
  goods_num      = 1
  bot_expand     = 0
  interface_name = "qcloud.Hour.create"
}
```

Import

NGWAF instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_instance.clb_waf waf-xxxxxxxx
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceTencentCloudNgwafInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafInstanceCreate,
		Read:   resourceTencentCloudNgwafInstanceRead,
		Update: resourceTencentCloudNgwafInstanceUpdate,
		Delete: resourceTencentCloudNgwafInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Instance name.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Region ID.",
			},
			"pay_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Payment mode, 0: post-pay, 1: pre-pay.",
			},
			"pid": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Product ID.",
			},
			"type": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Instance type.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
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
			"bot_expand": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				ForceNew:    true,
				Description: "Bot expansion flag.",
			},
			"interface_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Interface name for API request.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance unique ID.",
			},
			"edition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance edition.",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Region name.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Instance status.",
			},
			"attack_log_post": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Attack log post status.",
			},
			"valid_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Valid time.",
			},
			"begin_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Begin time.",
			},
			"domain_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Domain count.",
			},
			"main_domain_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Main domain count.",
			},
		},
	}
}

func resourceTencentCloudNgwafInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewWafCreateResourceAfterPayRequest()
	)

	if v, ok := d.GetOk("region_id"); ok {
		request.RegionId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("pay_mode"); ok {
		request.PayMode = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("pid"); ok {
		request.Pid = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok {
		request.Type = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("goods_num"); ok {
		request.GoodsNum = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOkExists("bot_expand"); ok {
		request.BotExpand = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("interface_name"); ok {
		request.InterfaceName = helper.String(v.(string))
	}

	var result *ngwaf.WafCreateResourceAfterPayResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().WafCreateResourceAfterPay(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			result = response
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf instance failed, reason:%+v", logId, err)
		return err
	}

	if result == nil || result.Response == nil || result.Response.InstanceId == nil {
		return fmt.Errorf("instance ID is nil in create response")
	}

	d.SetId(*result.Response.InstanceId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := &NgwafService{client: meta.(*TencentCloudClient)}

	if err := service.DescribeWafInstanceWaitStatusById(ctx, d.Id()); err != nil {
		log.Printf("[CRITAL]%s wait waf instance status failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafInstanceRead(d, meta)
}


func resourceTencentCloudNgwafInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	instanceId := d.Id()

	instanceInfo, err := service.DescribeWafInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if instanceInfo == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_ngwaf_instance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if instanceInfo.InstanceName != nil {
		_ = d.Set("instance_name", instanceInfo.InstanceName)
	}

	if instanceInfo.InstanceId != nil {
		_ = d.Set("instance_id", instanceInfo.InstanceId)
	}

	if instanceInfo.Edition != nil {
		_ = d.Set("edition", instanceInfo.Edition)
	}

	if instanceInfo.Region != nil {
		_ = d.Set("region", instanceInfo.Region)
	}

	if instanceInfo.RegionId != nil {
		_ = d.Set("region_id", fmt.Sprintf("%d", *instanceInfo.RegionId))
	}

	if instanceInfo.PayMode != nil {
		_ = d.Set("pay_mode", instanceInfo.PayMode)
	}

	if instanceInfo.Status != nil {
		_ = d.Set("status", instanceInfo.Status)
	}

	if instanceInfo.AttackLogPost != nil {
		_ = d.Set("attack_log_post", instanceInfo.AttackLogPost)
	}

	if instanceInfo.ValidTime != nil {
		_ = d.Set("valid_time", instanceInfo.ValidTime)
	}

	if instanceInfo.BeginTime != nil {
		_ = d.Set("begin_time", instanceInfo.BeginTime)
	}

	if instanceInfo.DomainCount != nil {
		_ = d.Set("domain_count", instanceInfo.DomainCount)
	}

	if instanceInfo.MainDomainCount != nil {
		_ = d.Set("main_domain_count", instanceInfo.MainDomainCount)
	}

	return nil
}

func resourceTencentCloudNgwafInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyInstanceNameRequest()
	)

	if !d.HasChange("instance_name") {
		return resourceTencentCloudNgwafInstanceRead(d, meta)
	}

	instanceId := d.Id()
	request.InstanceID = &instanceId

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok{
		instanceType := v.(int)
		if instanceType == 0 {
			request.Edition = helper.String(EDITION_SAAS)
		} else if instanceType == 1 {
			request.Edition = helper.String(EDITION_CLB)
		}

	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyInstanceName(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf instance name failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafInstanceRead(d, meta)
}

func resourceTencentCloudNgwafInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewDeleteInstanceRequest()
	)

	instanceId := d.Id()
	request.InstanceID = &instanceId

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DeleteInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s delete waf instance failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
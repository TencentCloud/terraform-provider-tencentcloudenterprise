/*
Provides a resource to manage NGWAF instance attack log post configuration.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ngwaf_instance_attack_log_post_config" "example" {
	  instance_id     = "waf_2kxtlbky11b4wcrb"
	  attack_log_post = 1
	}

```

# Import

NGWAF instance attack log post configuration can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example waf_2kxtlbky11b4wcrb
```
*/
package tencentcloud

import (
	"context"
	"log"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ngwaf_instance_attack_log_post_config", CNDescription{
		TerraformTypeCN: "NGWAF 实例攻击日志投递配置",
		DescriptionCN:   "提供 NGWAF 实例攻击日志投递配置资源，用于开启或关闭实例级攻击日志投递开关。",
		AttributesCN: map[string]string{
			"instance_id":     "WAF 实例 ID",
			"attack_log_post": "攻击日志投递开关，1 表示开启，0 表示关闭",
		},
	})
}

func resourceTencentCloudNgwafInstanceAttackLogPostConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafInstanceAttackLogPostConfigCreate,
		Read:   resourceTencentCloudNgwafInstanceAttackLogPostConfigRead,
		Update: resourceTencentCloudNgwafInstanceAttackLogPostConfigUpdate,
		Delete: resourceTencentCloudNgwafInstanceAttackLogPostConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Waf instance ID.",
			},
			"attack_log_post": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "Attack log delivery switch. 0 means disabled and 1 means enabled.",
			},
		},
	}
}

func resourceTencentCloudNgwafInstanceAttackLogPostConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance_attack_log_post_config.create")()
	defer inconsistentCheck(d, meta)()

	if v, ok := d.GetOk("instance_id"); ok {
		d.SetId(v.(string))
	}

	return resourceTencentCloudNgwafInstanceAttackLogPostConfigUpdate(d, meta)
}

func resourceTencentCloudNgwafInstanceAttackLogPostConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance_attack_log_post_config.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.Background(), logIdKey, logId)
		service    = &NgwafService{client: meta.(*TencentCloudClient)}
		instanceId = d.Id()
	)

	respData, err := service.DescribeWafInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ngwaf_instance_attack_log_post_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)

	if respData.AttackLogPost != nil {
		_ = d.Set("attack_log_post", *respData.AttackLogPost)
	}

	return nil
}

func resourceTencentCloudNgwafInstanceAttackLogPostConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance_attack_log_post_config.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = ngwaf.NewModifyInstanceAttackLogPostRequest()
		instanceId = d.Id()
	)

	request.InstanceId = &instanceId

	if v, ok := d.GetOkExists("attack_log_post"); ok {
		request.AttackLogPost = helper.IntInt64(v.(int))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyInstanceAttackLogPost(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if reqErr != nil {
		log.Printf("[CRITAL]%s update ngwaf instance attack log post config failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudNgwafInstanceAttackLogPostConfigRead(d, meta)
}

func resourceTencentCloudNgwafInstanceAttackLogPostConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_instance_attack_log_post_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

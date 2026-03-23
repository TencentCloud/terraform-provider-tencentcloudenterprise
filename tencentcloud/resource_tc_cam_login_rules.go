/*
Provides a resource to manage CAM login rules.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_login_rules" "foo" {
	  session_duration = 1440
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_login_rules", CNDescription{
		TerraformTypeCN: "CAM登录规则",
		DescriptionCN:   "提供 CAM 登录规则资源，用于设置用户的登录会话时长等规则。",
		AttributesCN: map[string]string{
			"session_duration": "登录会话时长 (分钟)",
		},
	})
}

func resourceTencentCloudCamLoginRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamLoginRulesCreate,
		Read:   resourceTencentCloudCamLoginRulesRead,
		Update: resourceTencentCloudCamLoginRulesUpdate,
		Delete: resourceTencentCloudCamLoginRulesDelete,

		Schema: map[string]*schema.Schema{
			"session_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Login session duration in minutes.",
			},
		},
	}
}

func resourceTencentCloudCamLoginRulesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_login_rules.create")()
	d.SetId("cam-login-rules")

	if err := resourceTencentCloudCamLoginRulesApply(d, meta); err != nil {
		return err
	}
	return resourceTencentCloudCamLoginRulesRead(d, meta)
}

func resourceTencentCloudCamLoginRulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_login_rules.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}
	sessionDuration, err := camService.DescribeLoginRules(ctx)
	if err != nil {
		return err
	}

	if sessionDuration == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("session_duration", int(*sessionDuration))

	return nil
}

func resourceTencentCloudCamLoginRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_login_rules.update")()
	defer inconsistentCheck(d, meta)()

	if err := resourceTencentCloudCamLoginRulesApply(d, meta); err != nil {
		return err
	}
	return resourceTencentCloudCamLoginRulesRead(d, meta)
}

func resourceTencentCloudCamLoginRulesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_login_rules.delete")()
	defer inconsistentCheck(d, meta)()
	// No delete API; keep state only.
	return nil
}

func resourceTencentCloudCamLoginRulesApply(d *schema.ResourceData, meta interface{}) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	sessionDuration := expandCamLoginRules(d)
	if sessionDuration == nil {
		return nil
	}

	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if e := camService.SetLoginRules(ctx, sessionDuration); e != nil {
			log.Printf("[CRITAL]%s set login rules fail, reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update CAM login rules failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

func expandCamLoginRules(d *schema.ResourceData) (sessionDuration *int64) {
	if v, ok := d.GetOk("session_duration"); ok {
		val := int64(v.(int))
		sessionDuration = &val
	}
	return
}

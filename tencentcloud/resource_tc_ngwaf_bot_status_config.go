/*
Provides a resource to manage the bot status configuration for a NGWAF protected domain.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_bot_status_config" "example" {
  instance_id = "waf-xxxxxxxx"
  domain      = "example.com"
  status      = "1"
}
```

Import

NGWAF bot status config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_bot_status_config.example waf-xxxxxxxx#example.com
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

func resourceTencentCloudNgwafBotStatusConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafBotStatusConfigCreate,
		Read:   resourceTencentCloudNgwafBotStatusConfigRead,
		Update: resourceTencentCloudNgwafBotStatusConfigUpdate,
		Delete: resourceTencentCloudNgwafBotStatusConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID.",
			},

			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain.",
			},

			"status": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bot status. 1 - enable; 0 - disable.",
			},

			"scene_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Scene total count.",
			},

			"valid_scene_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of effective scenarios.",
			},

			"current_global_scene": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The currently enabled scenario with a global matching range and the highest priority.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scene_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Scene ID.",
						},
						"scene_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Scene name.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Priority.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update time.",
						},
					},
				},
			},

			"custom_rule_nums": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of custom rules, excluding BOT whitelist.",
			},
		},
	}
}

func resourceTencentCloudNgwafBotStatusConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_status_config.create")()
	defer inconsistentCheck(d, meta)()

	var (
		instanceId string
		domain     string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
	}

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	d.SetId(strings.Join([]string{instanceId, domain}, FILED_SP))

	return resourceTencentCloudNgwafBotStatusConfigUpdate(d, meta)
}

func resourceTencentCloudNgwafBotStatusConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_status_config.read")()
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

	instanceId := idSplit[0]
	domain := idSplit[1]

	respData, err := service.DescribeWafBotStatusConfigById(ctx, domain)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `waf_bot_status_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)
	_ = d.Set("instance_id", instanceId)

	if respData.Response != nil && respData.Response.Status != nil {
		if *respData.Response.Status {
			_ = d.Set("status", "1")
		} else {
			_ = d.Set("status", "0")
		}
	}

	if respData.Response != nil && respData.Response.SceneCount != nil {
		_ = d.Set("scene_count", respData.Response.SceneCount)
	}

	if respData.Response != nil && respData.Response.ValidSceneCount != nil {
		_ = d.Set("valid_scene_count", respData.Response.ValidSceneCount)
	}

	if respData.Response != nil && respData.Response.CurrentGlobalScene != nil {
		tmpList := make([]map[string]interface{}, 0, 1)
		dMap := make(map[string]interface{})
		if respData.Response.CurrentGlobalScene.SceneId != nil {
			dMap["scene_id"] = respData.Response.CurrentGlobalScene.SceneId
		}

		if respData.Response.CurrentGlobalScene.SceneName != nil {
			dMap["scene_name"] = respData.Response.CurrentGlobalScene.SceneName
		}

		if respData.Response.CurrentGlobalScene.Priority != nil {
			dMap["priority"] = respData.Response.CurrentGlobalScene.Priority
		}

		if respData.Response.CurrentGlobalScene.UpdateTime != nil {
			dMap["update_time"] = respData.Response.CurrentGlobalScene.UpdateTime
		}

		tmpList = append(tmpList, dMap)
		_ = d.Set("current_global_scene", tmpList)
	}

	if respData.Response != nil && respData.Response.CustomRuleNums != nil {
		_ = d.Set("custom_rule_nums", respData.Response.CustomRuleNums)
	}

	return nil
}

func resourceTencentCloudNgwafBotStatusConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_status_config.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyBotStatusRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	instanceId := idSplit[0]
	domain := idSplit[1]

	if v, ok := d.GetOk("status"); ok {
		request.Status = helper.String(v.(string))
	}

	request.InstanceID = &instanceId
	request.Domain = &domain
	request.Category = helper.String("bot")
	request.IsVersionFour = helper.Bool(true)
	request.BotVersion = helper.String("4.1.0")
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyBotStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s update waf bot status config failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudNgwafBotStatusConfigRead(d, meta)
}

func resourceTencentCloudNgwafBotStatusConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_status_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

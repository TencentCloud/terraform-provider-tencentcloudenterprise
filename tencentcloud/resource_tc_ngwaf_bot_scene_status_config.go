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

func resourceTencentCloudNgwafBotSceneStatusConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafBotSceneStatusConfigCreate,
		Read:   resourceTencentCloudNgwafBotSceneStatusConfigRead,
		Update: resourceTencentCloudNgwafBotSceneStatusConfigUpdate,
		Delete: resourceTencentCloudNgwafBotSceneStatusConfigDelete,
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

			"scene_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Scene ID.",
			},

			"status": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Bot status. true - enable; false - disable.",
			},

			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Scene type, default: Default scenario, custom: Non default scenario.",
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
		},
	}
}

func resourceTencentCloudNgwafBotSceneStatusConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_scene_status_config.create")()
	defer inconsistentCheck(d, meta)()

	var (
		domain  string
		sceneId string
	)

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	if v, ok := d.GetOk("scene_id"); ok {
		sceneId = v.(string)
	}

	d.SetId(strings.Join([]string{domain, sceneId}, FILED_SP))

	return resourceTencentCloudNgwafBotSceneStatusConfigUpdate(d, meta)
}

func resourceTencentCloudNgwafBotSceneStatusConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_scene_status_config.read")()
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
	sceneId := idSplit[1]

	respData, err := service.DescribeWafBotSceneStatusConfigById(ctx, domain, sceneId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `waf_bot_scene_status_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)
	_ = d.Set("scene_id", sceneId)

	if respData.SceneStatus != nil {
		_ = d.Set("status", respData.SceneStatus)
	}

	if respData.Type != nil {
		_ = d.Set("type", respData.Type)
	}

	if respData.SceneName != nil {
		_ = d.Set("scene_name", respData.SceneName)
	}

	if respData.Priority != nil {
		_ = d.Set("priority", respData.Priority)
	}

	return nil
}

func resourceTencentCloudNgwafBotSceneStatusConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_scene_status_config.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyBotSceneStatusRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	sceneId := idSplit[1]

	if v, ok := d.GetOkExists("status"); ok {
		request.Status = helper.Bool(v.(bool))
	}

	request.Domain = &domain
	request.SceneId = &sceneId
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyBotSceneStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s update waf bot scene status config failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudNgwafBotSceneStatusConfigRead(d, meta)
}

func resourceTencentCloudNgwafBotSceneStatusConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_bot_scene_status_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

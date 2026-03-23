/*
Provides a resource to manage CAM API keys.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_api_key" "foo" {
	  api_uin = "100000000001"
	  status  = "enabled"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_api_key", CNDescription{
		TerraformTypeCN: "CAM API密钥",
		DescriptionCN:   "提供 CAM API 密钥资源，用于管理子用户的 API 访问密钥。",
		AttributesCN: map[string]string{
			"api_uin":           "拥有该密钥的 API UIN",
			"custom_secret_id":  "自定义 SecretId",
			"custom_secret_key": "自定义 SecretKey",
			"status":            "密钥状态: enabled (启用) 或 disabled (禁用)",
			"secret_id":         "密钥 ID",
			"secret_key":        "密钥 Key",
			"create_time":       "创建时间 (Unix 时间戳)",
			"source":            "密钥来源",
		},
	})
}

const (
	camAPIKeyEnabled  = "enabled"
	camAPIKeyDisabled = "disabled"
)

var camAPIKeyStatusStr2Int = map[string]uint64{
	camAPIKeyEnabled:  2,
	camAPIKeyDisabled: 3,
}

func camAPIKeyInt2Str(v *uint64) string {
	if v == nil {
		return ""
	}
	for k, val := range camAPIKeyStatusStr2Int {
		if val == *v {
			return k
		}
	}
	return strconv.FormatUint(*v, 10)
}

func resourceTencentCloudCamApiKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamApiKeyCreate,
		Read:   resourceTencentCloudCamApiKeyRead,
		Update: resourceTencentCloudCamApiKeyUpdate,
		Delete: resourceTencentCloudCamApiKeyDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// Import ID format: <api_uin>#<secret_id>
				parts := strings.Split(d.Id(), "#")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid import id format, expected <api_uin>#<secret_id>")
				}
				if err := d.Set("api_uin", parts[0]); err != nil {
					return nil, err
				}
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"api_uin": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "API UIN owning the key.",
			},
			"custom_secret_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Custom SecretId to create (optional).",
			},
			"custom_secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				ForceNew:    true,
				Description: "Custom SecretKey to create (optional, required if custom_secret_id is set).",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      camAPIKeyEnabled,
				ValidateFunc: validation.StringInSlice([]string{camAPIKeyEnabled, camAPIKeyDisabled}, false),
				Description:  "Key status: enabled or disabled.",
			},
			"secret_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secret ID of the API key.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Secret key material.",
			},
			"create_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Create time (unix timestamp).",
			},
			"source": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Source field from CAM API key.",
			},
		},
	}
}

func resourceTencentCloudCamApiKeyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_api_key.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	apiUinStr := d.Get("api_uin").(string)
	apiUin, err := strconv.ParseUint(apiUinStr, 10, 64)
	if err != nil {
		return fmt.Errorf("api_uin must be uint64: %v", err)
	}

	var customSecretId, customSecretKey *string
	if v, ok := d.GetOk("custom_secret_id"); ok {
		csid := v.(string)
		customSecretId = &csid
	}
	if v, ok := d.GetOk("custom_secret_key"); ok {
		csk := v.(string)
		customSecretKey = &csk
	}

	var created *cam.ApiKey
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		key, e := camService.CreateApiKey(ctx, apiUin, customSecretId, customSecretKey)
		if e != nil {
			log.Printf("[CRITAL]%s create api key fail, reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		created = key
		return nil
	})
	if err != nil {
		return err
	}

	if created == nil || created.SecretId == nil {
		return fmt.Errorf("create api key got empty secret_id")
	}
	d.SetId(*created.SecretId)
	_ = d.Set("secret_id", created.SecretId)
	if created.SecretKey != nil {
		_ = d.Set("secret_key", created.SecretKey)
	}

	// set desired status
	desiredStatus := d.Get("status").(string)
	if desiredStatus == camAPIKeyDisabled {
		if err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			if e := camService.DisableApiKey(ctx, apiUin, *created.SecretId); e != nil {
				return retryError(e)
			}
			return nil
		}); err != nil {
			return err
		}
	}

	return resourceTencentCloudCamApiKeyRead(d, meta)
}

func resourceTencentCloudCamApiKeyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_api_key.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	apiUinStr := d.Get("api_uin").(string)
	apiUin, err := strconv.ParseUint(apiUinStr, 10, 64)
	if err != nil {
		return fmt.Errorf("api_uin must be uint64: %v", err)
	}

	var keys []*cam.ApiKey
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		list, e := camService.QueryApiKey(ctx, apiUin)
		if e != nil {
			return retryError(e)
		}
		keys = list
		return nil
	})
	if err != nil {
		return err
	}

	var found *cam.ApiKey
	for _, k := range keys {
		if k.SecretId != nil && *k.SecretId == d.Id() {
			found = k
			break
		}
	}

	if found == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("secret_id", found.SecretId)
	if found.SecretKey != nil {
		_ = d.Set("secret_key", found.SecretKey)
	}
	if found.CreateTime != nil {
		_ = d.Set("create_time", int(*found.CreateTime))
	}
	if found.Source != nil {
		_ = d.Set("source", int(*found.Source))
	}
	_ = d.Set("status", camAPIKeyInt2Str(found.Status))

	return nil
}

func resourceTencentCloudCamApiKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_api_key.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	apiUinStr := d.Get("api_uin").(string)
	apiUin, err := strconv.ParseUint(apiUinStr, 10, 64)
	if err != nil {
		return fmt.Errorf("api_uin must be uint64: %v", err)
	}

	secretId := d.Id()
	if d.HasChange("status") {
		desired := d.Get("status").(string)
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			var e error
			if desired == camAPIKeyDisabled {
				e = camService.DisableApiKey(ctx, apiUin, secretId)
			} else {
				e = camService.EnableApiKey(ctx, apiUin, secretId)
			}
			if e != nil {
				return retryError(e)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudCamApiKeyRead(d, meta)
}

func resourceTencentCloudCamApiKeyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_api_key.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	apiUinStr := d.Get("api_uin").(string)
	apiUin, err := strconv.ParseUint(apiUinStr, 10, 64)
	if err != nil {
		return fmt.Errorf("api_uin must be uint64: %v", err)
	}

	secretId := d.Id()

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if e := camService.DeleteApiKey(ctx, apiUin, secretId); e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM api key failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

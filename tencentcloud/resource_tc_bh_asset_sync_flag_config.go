/*
Provide a resource to create a BH asset sync flag config

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_asset_sync_flag_config" "example" {
	  auto_sync = true
	}

```

# Import

BH asset sync flag config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_asset_sync_flag_config.example id
```
*/
package tencentcloud

import (
	"context"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_asset_sync_flag_config", CNDescription{
		TerraformTypeCN: "BH 资产自动同步配置",
		DescriptionCN:   "提供 BH 资产自动同步配置资源，用于管理堡垒机资产自动同步开关。",
		AttributesCN: map[string]string{
			"auto_sync":    "是否开启资产自动同步",
			"role_granted": "角色是否已授权",
		},
	})
}

func ResourceTencentCloudBhAssetSyncFlagConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhAssetSyncFlagConfigCreate,
		Read:   resourceTencentCloudBhAssetSyncFlagConfigRead,
		Update: resourceTencentCloudBhAssetSyncFlagConfigUpdate,
		Delete: resourceTencentCloudBhAssetSyncFlagConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"auto_sync": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable asset auto-sync, false - disabled, true - enabled.",
			},

			// computed
			"role_granted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the role has been authorized, false - not authorized, true - authorized.",
			},
		},
	}
}

func resourceTencentCloudBhAssetSyncFlagConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_flag_config.create")()
	defer inconsistentCheck(d, meta)()

	d.SetId(helper.BuildToken())

	return resourceTencentCloudBhAssetSyncFlagConfigUpdate(d, meta)
}

func resourceTencentCloudBhAssetSyncFlagConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_flag_config.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = BhService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	respData, err := service.DescribeBhAssetSyncFlagConfigById(ctx)
	if err != nil {
		return err
	}

	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_bh_asset_sync_flag_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if respData.AutoSync != nil {
		_ = d.Set("auto_sync", respData.AutoSync)
	}

	if respData.RoleGranted != nil {
		_ = d.Set("role_granted", respData.RoleGranted)
	}

	return nil
}

func resourceTencentCloudBhAssetSyncFlagConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_flag_config.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = bhsaas.NewModifyAssetSyncFlagRequest()
	)

	if v, ok := d.GetOkExists("auto_sync"); ok {
		request.AutoSync = helper.Bool(v.(bool))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyAssetSyncFlag(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s update bh asset sync flag config failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudBhAssetSyncFlagConfigRead(d, meta)
}

func resourceTencentCloudBhAssetSyncFlagConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_flag_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

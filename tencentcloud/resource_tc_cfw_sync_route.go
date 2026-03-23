/*
Provides a resource to trigger cloud firewall (cfw) route synchronization.

Example Usage

```hcl

resource "tencentcloudenterprise_cfw_sync_route" "example" {
  sync_type = "Route"
  fw_type   = "nat"
}

```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strconv"
	"time"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudCfwSyncRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCfwSyncRouteCreate,
		Read:   resourceTencentCloudCfwSyncRouteRead,
		Delete: resourceTencentCloudCfwSyncRouteDelete,

		Schema: map[string]*schema.Schema{
			"sync_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Synchronization operation type: Route, synchronize firewall routing.",
			},
			"fw_type": {
				Optional:     true,
				ForceNew:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(FW_TYPE),
				Description:  "Firewall type; nat: nat firewall; ew: inter-vpc firewall.",
			},
		},
	}
}

func resourceTencentCloudCfwSyncRouteCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sync_route.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		request       = cfw.NewSyncFwOperateRequest()
		statusRequest = cfw.NewDescribeFwSyncStatusRequest()
	)

	if v, ok := d.GetOk("sync_type"); ok {
		request.SyncType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("fw_type"); ok {
		request.FwType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().SyncFwOperate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s operate cfw syncRoute failed, reason:%+v", logId, err)
		return err
	}

	// wait
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().DescribeFwSyncStatus(statusRequest)
		if e != nil {
			return retryError(e)
		}

		if *result.Response.SyncStatus == 0 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("The fw sync status is %d.", *result.Response.SyncStatus))
	})

	if err != nil {
		log.Printf("[CRITAL]%s operate cfw syncAsset status failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceTencentCloudCfwSyncRouteRead(d, meta)
}

func resourceTencentCloudCfwSyncRouteRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sync_route.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCfwSyncRouteDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sync_route.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_sync_route", CNDescription{
		TerraformTypeCN: "路由同步",
		DescriptionCN:   "提供云防火墙路由同步资源，用于触发云防火墙路由同步操作。",
		AttributesCN: map[string]string{
			"sync_type": "同步类型",
			"fw_type":   "防火墙类型",
		},
	})
}

/*
Provides a resource to manage identity center scim synchronization status

Example Usage

```hcl
resource "tencentcloudenterprise_cic_scim_synchronization_status" "cic_scim_synchronization_status" {
  zone_id = "z-xxxxxx"
  scim_synchronization_status = "Enabled"
}
```

Import

organization cic_scim_synchronization_status can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_scim_synchronization_status.cic_scim_synchronization_status ${zone_id}
```

 */
package tencentcloud

import (
	"context"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_scim_synchronization_status", CNDescription{
		TerraformTypeCN: "身份中心SCIM同步状态",
		DescriptionCN:   "提供身份中心SCIM同步状态资源，用于管理SCIM同步功能的启用状态。",
		AttributesCN: map[string]string{
			"zone_id":                     "空间ID",
			"scim_synchronization_status": "SCIM同步状态",
		},
	})
}

func resourceTencentCloudCicScimSynchronizationStatus() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage identity center scim synchronization status",
		Create: resourceTencentCloudCicScimSynchronizationStatusCreate,
		Read:   resourceTencentCloudCicScimSynchronizationStatusRead,
		Update: resourceTencentCloudCicScimSynchronizationStatusUpdate,
		Delete: resourceTencentCloudCicScimSynchronizationStatusDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID. z-prefix starts with 12 random digits/lowercase letters.",
			},

			"scim_synchronization_status": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SCIM synchronization status. Enabled-enabled. Disabled-disables.",
			},
		},
	}
}

func resourceTencentCloudCicScimSynchronizationStatusCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_synchronization_status.create")()
	defer inconsistentCheck(d, meta)()

	var (
		zoneId string
	)
	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	d.SetId(zoneId)

	return resourceTencentCloudCicScimSynchronizationStatusUpdate(d, meta)
}

func resourceTencentCloudCicScimSynchronizationStatusRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_synchronization_status.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	zoneId := d.Id()

	_ = d.Set("zone_id", zoneId)

	respData, err := service.DescribeCicScimSynchronizationStatusById(ctx, zoneId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_scim_synchronization_status` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.SCIMSynchronizationStatus != nil {
		_ = d.Set("scim_synchronization_status", respData.SCIMSynchronizationStatus)
	}

	return nil
}

func resourceTencentCloudCicScimSynchronizationStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_synchronization_status.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	zoneId := d.Id()

	needChange := false
	mutableArgs := []string{"scim_synchronization_status"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewUpdateSCIMSynchronizationStatusRequest()

		request.ZoneId = helper.String(zoneId)

		if v, ok := d.GetOk("scim_synchronization_status"); ok {
			request.SCIMSynchronizationStatus = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateSCIMSynchronizationStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center scim synchronization status failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCicScimSynchronizationStatusRead(d, meta)
}

func resourceTencentCloudCicScimSynchronizationStatusDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_synchronization_status.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

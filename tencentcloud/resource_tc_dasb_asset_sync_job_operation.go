/*
Provide a resource to create a DASB asset sync job operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_asset_sync_job_operation" "example" {
	  category = 1
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strconv"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_asset_sync_job_operation", CNDescription{
		TerraformTypeCN: "DASB 资产同步任务",
		DescriptionCN:   "提供 DASB 资产同步任务资源，用于创建和执行资产同步任务。",
		AttributesCN: map[string]string{
			"category": "同步资产类别",
		},
	})
}

func ResourceTencentCloudDasbAssetSyncJobOperationOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbAssetSyncJobOperationCreate,
		Read:   resourceTencentCloudDasbAssetSyncJobOperationRead,
		Delete: resourceTencentCloudDasbAssetSyncJobOperationDelete,

		Schema: map[string]*schema.Schema{
			"category": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Synchronize asset categories, 1- Host assets, 2- Database assets.",
			},
		},
	}
}

func resourceTencentCloudDasbAssetSyncJobOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_asset_sync_job_operation.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewCreateAssetSyncJobRequest()
		waitReq  = bhsaas.NewDescribeAssetSyncStatusRequest()
		category string
	)

	if v, ok := d.GetOkExists("category"); ok {
		request.Category = helper.IntUint64(v.(int))
		waitReq.Category = helper.IntUint64(v.(int))
		category = strconv.Itoa(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateAssetSyncJob(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb AssetSyncJob failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(category)

	// wait
	err = resource.Retry(4*writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeAssetSyncStatus(waitReq)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.Status == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe dasb AssetSyncJob failed, Response is nil."))
		}

		if result.Response.Status.InProcess == nil {
			return resource.NonRetryableError(fmt.Errorf("InProcess is nil."))
		}

		if !*result.Response.Status.InProcess {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("Dasb asset sync job is still running..."))
	})

	if err != nil {
		log.Printf("[CRITAL]%s describe dasb AssetSyncJob failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbAssetSyncJobOperationRead(d, meta)
}

func resourceTencentCloudDasbAssetSyncJobOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_asset_sync_job_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudDasbAssetSyncJobOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_asset_sync_job_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

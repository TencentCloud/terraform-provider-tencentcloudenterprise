/*
Provide a resource to create a BH asset sync job operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_asset_sync_job_operation" "example" {
	  category = 1
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_asset_sync_job_operation", CNDescription{
		TerraformTypeCN: "BH 资产同步任务",
		DescriptionCN:   "提供 BH 资产同步任务资源，用于触发堡垒机资产同步操作。",
		AttributesCN: map[string]string{
			"category": "资产同步类别",
		},
	})
}

func ResourceTencentCloudBhAssetSyncJobOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhAssetSyncJobOperationCreate,
		Read:   resourceTencentCloudBhAssetSyncJobOperationRead,
		Delete: resourceTencentCloudBhAssetSyncJobOperationDelete,
		Schema: map[string]*schema.Schema{
			"category": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedIntValue([]int{1, 2, 3}),
				Description:  "Asset synchronization category. 1 - host assets, 2 - database assets, 3 - Container assets.",
			},
		},
	}
}

func resourceTencentCloudBhAssetSyncJobOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_job_operation.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewCreateAssetSyncJobRequest()
		category string
	)

	if v, ok := d.GetOkExists("category"); ok {
		request.Category = helper.IntUint64(v.(int))
		category = helper.IntToStr(v.(int))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateAssetSyncJob(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create bh asset sync job operation failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	d.SetId(category)

	// wait
	waitReq := bhsaas.NewDescribeAssetSyncStatusRequest()
	if v, ok := d.GetOkExists("category"); ok {
		waitReq.Category = helper.IntUint64(v.(int))
	}

	reqErr = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeAssetSyncStatus(waitReq)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.Status == nil || result.Response.Status.InProcess == nil {
			return resource.NonRetryableError(fmt.Errorf("Describe bh asset sync status failed, Response is nil."))
		}

		if !*result.Response.Status.InProcess {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("Asset sync status is still running..."))
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s describe bh asset sync status failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudBhAssetSyncJobOperationRead(d, meta)
}

func resourceTencentCloudBhAssetSyncJobOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_job_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudBhAssetSyncJobOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_asset_sync_job_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

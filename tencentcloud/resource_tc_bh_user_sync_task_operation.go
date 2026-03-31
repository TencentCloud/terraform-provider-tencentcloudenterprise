/*
Provide a resource to create a BH user sync task operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_user_sync_task_operation" "example" {
	  user_kind = 1
	}

```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_user_sync_task_operation", CNDescription{
		TerraformTypeCN: "BH 用户同步任务",
		DescriptionCN:   "提供 BH 用户同步任务资源，用于触发堡垒机用户同步操作。",
		AttributesCN: map[string]string{
			"user_kind": "同步用户类型",
		},
	})
}

func ResourceTencentCloudBhUserSyncTaskOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhUserSyncTaskOperationCreate,
		Read:   resourceTencentCloudBhUserSyncTaskOperationRead,
		Delete: resourceTencentCloudBhUserSyncTaskOperationDelete,
		Schema: map[string]*schema.Schema{
			"user_kind": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Synchronized user type, 1-synchronize IOA users.",
			},
		},
	}
}

func resourceTencentCloudBhUserSyncTaskOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_sync_task_operation.create")()
	defer inconsistentCheck(d, meta)()

	var (
		userKind string
	)

	if v, ok := d.GetOkExists("user_kind"); ok {
		userKind = helper.IntToStr(v.(int))
	}

	d.SetId(userKind)

	return resourceTencentCloudBhUserSyncTaskOperationRead(d, meta)
}

func resourceTencentCloudBhUserSyncTaskOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_sync_task_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudBhUserSyncTaskOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_sync_task_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

/*
Provides a resource to create a cbs snapshot_share_permission

# Example Usage

```hcl

	resource "tencentcloudenterprise_cbs_snapshot_share_permission" "snapshot_share_permission" {
	  account_ids = ["1xxxxxx", "2xxxxxx"]
	  snapshot_id = "snap-xxxxxx"
	}

```

# Import

cbs snapshot_share_permission can be imported using the id, e.g.

```
terraform import cloud_cbs_snapshot_share_permission.snapshot_share_permission snap-xxxxxx
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cbs_snapshot_share_permission", CNDescription{
		TerraformTypeCN: "快照分享",
		DescriptionCN:   "提供CBS快照分享权限资源，用于创建和管理快照分享权限。",
		AttributesCN: map[string]string{
			"account_ids": "账户ID列表",
			"snapshot_id": "快照ID",
		},
	})

}
func resourceTencentCloudCbsSnapshotSharePermission() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cbs snapshot_share_permission",
		Create:      resourceTencentCloudCbsSnapshotSharePermissionCreate,
		Read:        resourceTencentCloudCbsSnapshotSharePermissionRead,
		Update:      resourceTencentCloudCbsSnapshotSharePermissionUpdate,
		Delete:      resourceTencentCloudCbsSnapshotSharePermissionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"account_ids": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of account IDs with which a snapshot is shared. For the format of array-type parameters, see[API Introduction](https://cloud.tencent.com/document/api/213/568). You can find the account ID in[Account Information](https://console.cloud.tencent.com/developer).",
			},
			"snapshot_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of the snapshot to be queried. You can obtain this by using [DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647).",
			},
		},
	}
}

func resourceTencentCloudCbsSnapshotSharePermissionCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_share_permission.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var snapshotId string
	var accountIdsSet []interface{}
	log.Printf("11111111111%v", d)
	if v, ok := d.GetOk("account_ids"); ok {
		accountIdsSet = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("snapshot_id"); ok {
		snapshotId = v.(string)
	}

	service := CbsService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := service.ModifySnapshotsSharePermission(ctx, snapshotId, SNAPSHOT_SHARE_PERMISSION_SHARE, helper.InterfacesStringsPoint(accountIdsSet))
	if err != nil {
		return err
	}
	d.SetId(snapshotId)

	return resourceTencentCloudCbsSnapshotSharePermissionRead(d, meta)
}

func resourceTencentCloudCbsSnapshotSharePermissionRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_share_permission.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CbsService{client: meta.(*TencentCloudClient).apiV3Conn}

	snapshotId := d.Id()

	snapshotSharePermissions, err := service.DescribeCbsSnapshotSharePermissionById(ctx, snapshotId)
	if err != nil {
		return err
	}

	accountIds := make([]*string, 0)
	for _, snapshotSharePermission := range snapshotSharePermissions {
		accountIds = append(accountIds, helper.Int64ToStrPoint(*snapshotSharePermission.AccountId))
	}

	_ = d.Set("account_ids", accountIds)
	_ = d.Set("snapshot_id", snapshotId)
	return nil
}

func resourceTencentCloudCbsSnapshotSharePermissionUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_share_permission.update")()
	defer inconsistentCheck(d, meta)()

	snapshotId := d.Id()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CbsService{client: meta.(*TencentCloudClient).apiV3Conn}
	if d.HasChange("account_ids") {
		old, new := d.GetChange("account_ids")
		oldSet := old.(*schema.Set)
		newSet := new.(*schema.Set)
		add := newSet.Difference(oldSet).List()
		remove := oldSet.Difference(newSet).List()
		if len(add) > 0 {
			addError := service.ModifySnapshotsSharePermission(ctx, snapshotId, SNAPSHOT_SHARE_PERMISSION_SHARE, helper.InterfacesStringsPoint(add))
			if addError != nil {
				return addError
			}
		}
		if len(remove) > 0 {
			removeError := service.ModifySnapshotsSharePermission(ctx, snapshotId, SNAPSHOT_SHARE_PERMISSION_CANCEL, helper.InterfacesStringsPoint(remove))
			if removeError != nil {
				return removeError
			}
		}
	}

	return resourceTencentCloudCbsSnapshotSharePermissionRead(d, meta)
}

func resourceTencentCloudCbsSnapshotSharePermissionDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_share_permission.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CbsService{client: meta.(*TencentCloudClient).apiV3Conn}
	snapshotId := d.Id()
	snapshotSharePermissions, err := service.DescribeCbsSnapshotSharePermissionById(ctx, snapshotId)
	if err != nil {
		return err
	}

	accountIds := make([]*string, 0)
	for _, snapshotSharePermission := range snapshotSharePermissions {
		accountIds = append(accountIds, helper.Int64ToStrPoint(*snapshotSharePermission.AccountId))
	}
	if err := service.ModifySnapshotsSharePermission(ctx, snapshotId, SNAPSHOT_SHARE_PERMISSION_CANCEL, accountIds); err != nil {
		return err
	}

	return nil
}

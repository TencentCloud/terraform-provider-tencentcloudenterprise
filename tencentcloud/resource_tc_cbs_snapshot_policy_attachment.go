/*
Provides a CBS snapshot policy attachment resource.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cbs_snapshot_policy_attachment" "foo" {
	  storage_id         = cloud_cbs_storage.foo.id
	  snapshot_policy_id = cloud_cbs_snapshot_policy.policy.id
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cbs "terraform-provider-tencentcloudenterprise/sdk/cbs/v20170312"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cbs_snapshot_policy_attachment", CNDescription{
		TerraformTypeCN: "定期快照策略关联云硬盘",
		DescriptionCN:   "提供CBS快照策略关联资源，用于将快照策略关联到云硬盘。",
		AttributesCN: map[string]string{
			"storage_id":         "云硬盘ID",
			"snapshot_policy_id": "快照策略ID",
		},
	})
}

func resourceTencentCloudCbsSnapshotPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a CBS snapshot policy attachment resource.",
		Create:      resourceTencentCloudCbsSnapshotPolicyAttachmentCreate,
		Read:        resourceTencentCloudCbsSnapshotPolicyAttachmentRead,
		Delete:      resourceTencentCloudCbsSnapshotPolicyAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"storage_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CBS.",
			},
			"snapshot_policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CBS snapshot policy.",
			},
		},
	}
}

func resourceTencentCloudCbsSnapshotPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_policy_attachment.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	storageId := d.Get("storage_id").(string)
	policyId := d.Get("snapshot_policy_id").(string)
	cbsService := CbsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := cbsService.AttachSnapshotPolicy(ctx, storageId, policyId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s cbs storage policy attach failed, reason:%s\n ", logId, err.Error())
		return err
	}

	d.SetId(storageId + FILED_SP + policyId)
	return resourceTencentCloudCbsSnapshotPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCbsSnapshotPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	idSplit := strings.Split(id, FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("tencentcloudenterprise_cbs_snapshot_policy_attachment id is illegal: %s", id)
	}
	storageId := idSplit[0]
	policyId := idSplit[1]
	cbsService := CbsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var policy *cbs.AutoSnapshotPolicy
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		policy, errRet = cbsService.DescribeAttachedSnapshotPolicy(ctx, storageId, policyId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s cbs storage policy attach failed, reason:%s\n ", logId, err.Error())
		return err
	}
	if policy == nil {
		d.SetId("")
		return nil
	}
	_ = d.Set("storage_id", storageId)
	_ = d.Set("snapshot_policy_id", policyId)

	return nil
}

func resourceTencentCloudCbsSnapshotPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cbs_snapshot_policy_attachment.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	idSplit := strings.Split(id, FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("tencentcloudenterprise_cbs_snapshot_policy_attachment id is illegal: %s", id)
	}
	storageId := idSplit[0]
	policyId := idSplit[1]
	cbsService := CbsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := cbsService.UnattachSnapshotPolicy(ctx, storageId, policyId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s cbs storage policy unattach failed, reason:%s\n ", logId, err.Error())
		return err
	}

	return nil
}

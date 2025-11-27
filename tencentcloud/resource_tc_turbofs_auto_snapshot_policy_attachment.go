/*
Provides a resource to create a turbofs auto_snapshot_policy_attachment

# Example Usage

```hcl

	resource "tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment" "auto_snapshot_policy_attachment" {
	  auto_snapshot_policy_id = "asp-basic"
	  file_system_id          = "turbofs-4xzkct19"
	}

```

# Import

turbofs auto_snapshot_policy_attachment can be imported using the id, e.g.

```
terraform import cloud_turbofs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment auto_snapshot_policy_id#file_system_ids
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
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment", CNDescription{
		TerraformTypeCN: "自动快照策略绑定",
		DescriptionCN:   "提供TurboFS自动快照策略绑定资源，用于将自动快照策略绑定到TurboFS。",
		AttributesCN: map[string]string{
			"auto_snapshot_policy_id": "自动快照策略ID",
			"file_system_id":          "文件系统ID",
		},
	})
}
func resourceTencentCloudTurbofsAutoSnapshotPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a TurboFS auto snapshot policy attachment.",
		Create:      resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentCreate,
		Read:        resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentRead,
		Delete:      resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"auto_snapshot_policy_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of the snapshot to bound.",
			},

			"file_system_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of the file systems to bound.",
			},
		},
	}
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = turbofs.NewBindAutoSnapshotPolicyRequest()
		autoSnapshotPolicyId string
		fileSystemIds        string
	)
	if v, ok := d.GetOk("auto_snapshot_policy_id"); ok {
		autoSnapshotPolicyId = v.(string)
		request.AutoSnapshotPolicyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("file_system_id"); ok {
		fileSystemIds = v.(string)
		request.FileSystemIds = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().BindAutoSnapshotPolicy(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create turbofs autoSnapshotPolicyAttachment failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(autoSnapshotPolicyId + FILED_SP + fileSystemIds)

	return resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TurbofsService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	autoSnapshotPolicyId := idSplit[0]
	fileSystemIds := idSplit[1]

	snapShotId, fsId, err := service.DescribeCfsAutoSnapshotPolicyAttachmentById(ctx, autoSnapshotPolicyId, fileSystemIds)
	if err != nil {
		return err
	}

	if snapShotId != nil {
		_ = d.Set("auto_snapshot_policy_id", *snapShotId)
	}

	if fsId != nil {
		_ = d.Set("file_system_ids", *fsId)
	}

	return nil
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TurbofsService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	autoSnapshotPolicyId := idSplit[0]
	fileSystemIds := idSplit[1]

	if err := service.DeleteCfsAutoSnapshotPolicyAttachmentById(ctx, autoSnapshotPolicyId, fileSystemIds); err != nil {
		return err
	}

	return nil
}

/*
Provides a resource to create a brc disk(cbs) backup action

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_backup_disk" "example" {
	  disk_id      = "disk-ewei0a2q"
	  backup_name  = "my-backup"
	  deadline     = "2025-07-05 19:03:29"
	  backup_class = "INC"
	}

```

# Import

brc backup_disk can be imported using the id, e.g.

```
terraform import cloud_brc_backup_disk.example backup_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_backup_disk", CNDescription{
		TerraformTypeCN: "CBS云硬盘备份",
		DescriptionCN:   "提供BRC磁盘备份资源，用于创建CBS云硬盘备份。",
		AttributesCN: map[string]string{
			"disk_id":      "云硬盘ID",
			"backup_name":  "备份名称",
			"deadline":     "指定备份到期时间，如果未传入该参数，默认为永久保留。示例：\"2024-04-26 10:00:00\"",
			"backup_class": "指定做全量或增量备份。取值范围: FULL：全量备份, INC：增量备份。",
			"create_speed": "创建备份的带宽上限，范围：[0, 100]",
			"need_archive": "标识该磁盘备份将用于归档。",
			"backup_id":    "备份ID",
		},
	})
}

func resourceTencentCloudBrcBackupDisk() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create CBS backup resource.",
		Create:      resourceTencentCloudBrcBackupDiskCreate,
		Read:        resourceTencentCloudBrcBackupDiskRead,
		Delete:      resourceTencentCloudBrcBackupDiskDelete,

		Schema: map[string]*schema.Schema{
			"disk_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the CBS disk to be backed up.",
			},
			"backup_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the backup.",
			},
			"deadline": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The deadline for the backup. If not set, the backup will be permanent. standard time format: 2024-04-26 10:00:00",
			},
			"backup_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "FULL",
				Description: "The class of the backup (e.g., 'FULL' or 'INC'). default: FULL",
			},
			"create_speed": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "The bandwidth limit for creating backup, range: [0, 100].",
			},
			"need_archive": {
				Type:        schema.TypeBool,
				ForceNew:    true,
				Optional:    true,
				Description: "Indicates whether this disk backup will be used for archiving.",
			},
			// Computed fields
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the backup.",
			},
		},
	}
}

func resourceTencentCloudBrcBackupDiskCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_disk.create")()
	logId := getLogId(contextNil)

	request := brc.NewCreateBackupRequest()

	request.DiskId = helper.String(d.Get("disk_id").(string))

	if v, ok := d.GetOk("backup_name"); ok {
		request.BackupName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("deadline"); ok {
		layout := "2006-01-02 15:04:05"
		deadline, err := time.Parse(layout, v.(string))
		if err != nil {
			return fmt.Errorf("failed to parse deadline: %s", err)
		}
		request.Deadline = &deadline
	}

	if v, ok := d.GetOk("backup_class"); ok {
		request.BackupClass = helper.String(v.(string))
	}

	if v, ok := d.GetOk("create_speed"); ok {
		request.CreateSpeed = helper.Uint64(uint64(v.(int)))
	}

	if v, ok := d.GetOkExists("need_archive"); ok {
		request.NeedArchive = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseBrcClient().CreateBackup(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		d.SetId(*response.Response.BackupId)
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudBrcBackupDiskRead(d, meta)
}

func resourceTencentCloudBrcBackupDiskRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_disk.read")()
	defer inconsistentCheck(d, meta)()
	//logId := getLogId(contextNil)
	//ctx := context.WithValue(context.Background(), logIdKey, logId)
	//
	//brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	//backupId := d.Id()
	//var backup *brc.Backup
	//
	//err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
	//	backupData, e := brcService.DescribeBackupById(ctx, backupId)
	//	if e != nil {
	//		return retryError(e)
	//	}
	//	backup = backupData
	//	return nil
	//})
	//
	//if err != nil {
	//	return err
	//}
	//
	//if backup == nil {
	//	d.SetId("")
	//	return nil
	//}

	// Set all attributes
	//if backup.BackupId != nil {
	//	_ = d.Set("backup_id", backup.BackupId)
	//}
	//
	//if backup.DiskId != nil {
	//	_ = d.Set("disk_id", backup.DiskId)
	//}
	//
	//if backup.BackupName != nil {
	//	_ = d.Set("backup_name", backup.BackupName)
	//}
	//
	//if backup.BackupClass != nil {
	//	_ = d.Set("backup_class", backup.BackupClass)
	//}
	//
	//if backup.CreateSpeed != nil {
	//	_ = d.Set("create_speed", backup.CreateSpeed)
	//}
	//
	//if backup.NeedArchive != nil {
	//	_ = d.Set("need_archive", backup.NeedArchive)
	//}
	//
	return nil
}

func resourceTencentCloudBrcBackupDiskDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_disk.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	backupId := d.Id()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := brcService.DeleteBackupById(ctx, backupId)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, "DeleteBackup", e.Error())
			return retryError(e)
		}
		return nil
	})

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

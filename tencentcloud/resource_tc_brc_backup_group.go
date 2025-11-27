/*
Provides a resource to create a brc group(cvm instance disks) backup

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_backup_group" "example" {
	  disk_ids          = ["disk-23h487qfj", "disk-da73ha9dj"]
	  backup_group_name = "basic-backup-group"
	  deadline          = "2024-12-31 23:59:59"
	  backup_class      = "FULL"
	  create_speed      = 50
	  need_archive      = false
	}

```

# Import

brc backup_group can be imported using the id, e.g.

```
terraform import cloud_brc_backup_group.example backup_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_backup_group", CNDescription{
		TerraformTypeCN: "CVM云服务器的硬盘备份组",
		DescriptionCN:   "提供BRC备份组资源，用于创建CVM实例磁盘的备份组。",
		AttributesCN: map[string]string{
			"disk_ids":          "云服务器的云硬盘ID列表",
			"backup_group_name": "备份组名称",
			"deadline":          "指定备份组到期时间，如果未传入该参数，默认为永久保留。示例：\"2024-04-26 10:00:00\"",
			"create_speed":      "创建备份的带宽上限，范围：[0, 100]",
			"need_archive":      "是否需要归档，true：需要归档，false：不需要归档",
			"backup_class":      "指定做全量或增量备份。取值范围: FULL：全量备份, INC：增量备份。",
			"backup_group_id":   "备份组ID",
		},
	})
}

func resourceTencentCloudBrcBackupGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create CVM disks backup group resource.",
		Create:      resourceTencentCloudBrcBackupGroupCreate,
		Read:        resourceTencentCloudBrcBackupGroupRead,
		Delete:      resourceTencentCloudBrcBackupGroupDelete,

		Schema: map[string]*schema.Schema{
			"disk_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "CBS disk IDs of the CVM instance.",
			},
			"backup_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Backup group name.",
			},
			"deadline": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Deadline of the backup group. If not set, the backup group will be permanent. Standard time format: 2024-04-26 10:00:00.",
			},
			"create_speed": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Bandwidth limit for creating backup. Range: [0, 100].",
			},
			"need_archive": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether archiving is needed. true: need archive, false: don't need archive. Default: false.",
			},
			"backup_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "FULL",
				Description: "Backup class. FULL: full backup, INC: incremental backup. Default: FULL.",
			},
			"backup_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the CVM backup group.",
			},
		},
	}
}

func resourceTencentCloudBrcBackupGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_group.create")()
	logId := getLogId(contextNil)

	request := brc.NewCreateBackupCvmResourceRequest()

	var diskIds []*string
	for _, diskId := range d.Get("disk_ids").([]interface{}) {
		diskIds = append(diskIds, helper.String(diskId.(string)))
	}
	request.DiskIds = diskIds

	if v, ok := d.GetOk("backup_group_name"); ok {
		request.BackupGroupName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("deadline"); ok {
		layout := "2006-01-02 15:04:05"
		deadline, err := time.Parse(layout, v.(string))
		if err != nil {
			return fmt.Errorf("failed to parse deadline: %s", err)
		}
		request.Deadline = &deadline
	}

	if v, ok := d.GetOk("create_speed"); ok {
		request.CreateSpeed = helper.Uint64(uint64(v.(int)))
	}

	if v, ok := d.GetOk("need_archive"); ok {
		request.NeedArchive = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("backup_class"); ok {
		request.BackupClass = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseBrcClient().CreateBackupCvmResource(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		d.SetId(*response.Response.BackupGroupId)
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudBrcBackupGroupRead(d, meta)
}

func resourceTencentCloudBrcBackupGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_group.read")()
	defer inconsistentCheck(d, meta)()
	//logId := getLogId(contextNil)
	//ctx := context.WithValue(context.Background(), logIdKey, logId)
	//
	//brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	//backupGroupId := d.Id()
	//var backupGroup *brc.GroupBackup
	//
	//err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
	//	backup, e := brcService.DescribeBackupCvmResourceById(ctx, backupGroupId)
	//	if e != nil {
	//		return retryError(e)
	//	}
	//	backupGroup = backup
	//	return nil
	//})
	//
	//if err != nil {
	//	return err
	//}
	//
	//if backupGroup == nil {
	//	d.SetId("")
	//	return nil
	//}
	//
	//d.Set("disk_ids", backupGroup.DiskIds)
	//_ = d.Set("backup_name", backupGroup.BackupName)
	//_ = d.Set("backup_class", backupGroup.BackupClass)
	//_ = d.Set("deadline", backupGroup.Deadline)
	return nil
}

func resourceTencentCloudBrcBackupGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_group.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	backupGroupId := d.Id()

	var onlyDismiss *bool
	if v, ok := d.GetOk("only_dismiss"); ok {
		onlyDismiss = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := brcService.DeleteBackupCvmResourceById(ctx, backupGroupId, onlyDismiss)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, "DeleteBackupCvmResource", e.Error())
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

/*
Provides a resource to create a brc cfs backup

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_backup_cfs" "example" {
	  file_system_id = "cfs-bo74chkp"
	  backup_name    = "test"
	  deadline       = "2025-07-06 03:06:09"
	  backup_class   = "INC"
	}

```

# Import

brc backup_cfs can be imported using the id, e.g.

```
terraform import cloud_brc_backup_cfs.example backup_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_backup_cfs", CNDescription{
		TerraformTypeCN: "CFS文件系统备份",
		DescriptionCN:   "提供BRC文件系统备份资源，用于创建CFS文件系统备份。",
		AttributesCN: map[string]string{
			"file_system_id": "文件系统ID",
			"backup_name":    "备份名称",
			"deadline":       "指定备份到期时间，如果未传入该参数，默认为永久保留。示例：\"2024-04-26 10:00:00\"",
			"backup_class":   "指定做全量或增量备份。取值范围: FULL：全量备份, INC：增量备份。",
			"create_speed":   "创建备份的带宽上限，范围：[0, 100]",
			"need_archive":   "标识该文件系统备份将用于归档。",
			"backup_id":      "备份ID",
			//"source_remote_cloud": "生产源云名称",
			//"source_remote_region":"生产源地域名称",
			//"destination_remote_cloud":"跨云目标云名称",
			//"destination_remote_region":"跨云目标地域名称",
			//"remote_app_id":"跨云目的账号ID",
			//"cross_region_name":"跨地域信息",
			//"is_isolated_storage":"是否为独立存储",
		},
	})
}

func resourceTencentCloudBrcBackupCfs() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create CFS backup resource.",
		Create:      resourceTencentCloudBrcBackupCfsCreate,
		Read:        resourceTencentCloudBrcBackupCfsRead,
		Delete:      resourceTencentCloudBrcBackupCfsDelete,

		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the CFS file system to be backed up.",
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
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the backup.",
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
				Description: "Indicates whether this file system backup will be used for archiving.",
			},
			//"remote_info": {
			//	Type:        schema.TypeList,
			//	ForceNew:    true,
			//	Optional:    true,
			//	Description: "Cross cloud/region information for backup.",
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"source_remote_cloud": {
			//				Type:        schema.TypeString,
			//				Optional:    true,
			//				Description: "The name of the source cloud for production.",
			//			},
			//			"source_remote_region": {
			//				Type:        schema.TypeString,
			//				Optional:    true,
			//				Description: "The name of the source region for production.",
			//			},
			//			"destination_remote_cloud": {
			//				Type:        schema.TypeString,
			//				Optional:    true,
			//				Description: "The name of the target cloud for cross-cloud backup.",
			//			},
			//			"destination_remote_region": {
			//				Type:        schema.TypeString,
			//				Optional:    true,
			//				Description: "The name of the target region for cross-cloud backup.",
			//			},
			//			"remote_app_id": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Description: "The account ID for the cross-cloud destination.",
			//			},
			//			"cross_region_name": {
			//				Type:        schema.TypeString,
			//				Optional:    true,
			//				Description: "The cross-region information.",
			//			},
			//			"is_isolated_storage": {
			//				Type:        schema.TypeBool,
			//				Optional:    true,
			//				Description: "Indicates whether to use isolated storage.",
			//			},
			//		},
			//	},
			//},
		},
	}
}

func resourceTencentCloudBrcBackupCfsCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_cfs.create")()
	logId := getLogId(contextNil)

	request := brc.NewCreateCfsBackupRequest()

	request.FileSystemId = helper.String(d.Get("file_system_id").(string))

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

	//if v, ok := d.GetOk("remote_info"); ok {
	//	remoteInfoList := v.([]interface{})
	//	if len(remoteInfoList) > 0 {
	//		remoteInfoMap := remoteInfoList[0].(map[string]interface{})
	//		remoteInfoRequest := brc.RemoteStorageInfo{}
	//
	//		if sourceCloud, ok := remoteInfoMap["source_remote_cloud"].(string); ok && sourceCloud != "" {
	//			remoteInfoRequest.SourceRemoteCloud = helper.String(sourceCloud)
	//		}
	//
	//		if sourceRegion, ok := remoteInfoMap["source_remote_region"].(string); ok && sourceRegion != "" {
	//			remoteInfoRequest.SourceRemoteRegion = helper.String(sourceRegion)
	//		}
	//
	//		if destCloud, ok := remoteInfoMap["destination_remote_cloud"].(string); ok && destCloud != "" {
	//			remoteInfoRequest.DestinationRemoteCloud = helper.String(destCloud)
	//		}
	//
	//		if destRegion, ok := remoteInfoMap["destination_remote_region"].(string); ok && destRegion != "" {
	//			remoteInfoRequest.DestinationRemoteRegion = helper.String(destRegion)
	//		}
	//
	//		if appId, ok := remoteInfoMap["remote_app_id"].(int); ok {
	//			remoteInfoRequest.RemoteAppId = helper.Uint64(uint64(appId))
	//		}
	//
	//		if crossRegion, ok := remoteInfoMap["cross_region_name"].(string); ok && crossRegion != "" {
	//			remoteInfoRequest.CrossRegionName = helper.String(crossRegion)
	//		}
	//
	//		if isolatedStorage, ok := remoteInfoMap["is_isolated_storage"].(bool); ok {
	//			remoteInfoRequest.IsIsolatedStorage = helper.Bool(isolatedStorage)
	//		}
	//
	//		request.RemoteInfo = &remoteInfoRequest
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseBrcClient().CreateCfsBackup(request)
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

	return resourceTencentCloudBrcBackupCfsRead(d, meta)
}

func resourceTencentCloudBrcBackupCfsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_cfs.read")()
	defer inconsistentCheck(d, meta)()
	//logId := getLogId(contextNil)
	//ctx := context.WithValue(context.Background(), logIdKey, logId)
	//
	//brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	//backupId := d.Id()
	//var backup *brc.CfsBackup
	//
	//err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
	//	backupData, e := brcService.DescribeCfsBackupById(ctx, backupId)
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
	//
	//if backup.BackupId != nil {
	//	_ = d.Set("backup_id", backup.BackupId)
	//}
	//
	//if backup.FileSystemId != nil {
	//	_ = d.Set("file_system_id", backup.FileSystemId)
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
	//if backup.DeadlineTime != nil {
	//	_ = d.Set("deadline", backup.DeadlineTime)
	//}
	//
	//if backup.CreateSpeed != nil {
	//	_ = d.Set("create_speed", backup.CreateSpeed)
	//}
	//
	//if backup.NeedArchive != nil {
	//	_ = d.Set("need_archive", backup.NeedArchive)
	//}

	//if backup.RemoteInfo != nil {
	//	remoteInfo := make(map[string]interface{})
	//
	//	if backup.RemoteInfo.SourceRemoteCloud != nil {
	//		remoteInfo["source_remote_cloud"] = *backup.RemoteInfo.SourceRemoteCloud
	//	}
	//
	//	if backup.RemoteInfo.SourceRemoteRegion != nil {
	//		remoteInfo["source_remote_region"] = *backup.RemoteInfo.SourceRemoteRegion
	//	}
	//
	//	if backup.RemoteInfo.DestinationRemoteCloud != nil {
	//		remoteInfo["destination_remote_cloud"] = *backup.RemoteInfo.DestinationRemoteCloud
	//	}
	//
	//	if backup.RemoteInfo.DestinationRemoteRegion != nil {
	//		remoteInfo["destination_remote_region"] = *backup.RemoteInfo.DestinationRemoteRegion
	//	}
	//
	//	if backup.RemoteInfo.RemoteAppId != nil {
	//		remoteInfo["remote_app_id"] = *backup.RemoteInfo.RemoteAppId
	//	}
	//
	//	if backup.RemoteInfo.CrossRegionName != nil {
	//		remoteInfo["cross_region_name"] = *backup.RemoteInfo.CrossRegionName
	//	}
	//
	//	if backup.RemoteInfo.IsIsolatedStorage != nil {
	//		remoteInfo["is_isolated_storage"] = *backup.RemoteInfo.IsIsolatedStorage
	//	}
	//
	//	_ = d.Set("remote_info", []interface{}{remoteInfo})
	//}

	return nil
}

func resourceTencentCloudBrcBackupCfsDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_cfs.delete")()
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

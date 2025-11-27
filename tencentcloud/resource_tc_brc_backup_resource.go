/*
Provides a resource to create a brc resource backup

# Example Usage

```hcl
# cos backup

	resource "tencentcloudenterprise_brc_backup_resource" "cos_backup" {
	  resource_type = "COS"
	  backup_name   = "my-cos-backup"
	  bucket_detail {
	    cos_region = "chongqing"
	    prefix = "/"
	    resource_id = "test-456-1255000252"
	  }
	}

# mysql backup

	resource "tencentcloudenterprise_brc_backup_resource" "mysql_backup" {
	  resource_type = "MySQL_MariaDB"
	  backup_name   = "my-mysql-backup"
	  resource_id = "tdsql-baxmqi05"
	}

```

# Import

brc backup_resource can be imported using the id, e.g.

```
terraform import cloud_brc_backup_resource.example backup_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_backup_resource", CNDescription{
		TerraformTypeCN: "资源备份",
		DescriptionCN:   "提供BRC资源备份资源，用于创建和管理COS、CSP和数据库资源的备份。",
		AttributesCN: map[string]string{
			"resource_type":           "创建备份的资源类型。取值范围：TDSQL_MySQL：TDSQL_MySQL版, MySQL_MariaDB：关系型数据库(MySQL_MariaDB), COS：对象存储COS。",
			"backup_name":             "备份名称",
			"resource_id":             "创建备份的资源ID",
			"cos_region":              "COS地域",
			"prefix":                  "COS前缀",
			"create_speed":            "创建备份的带宽上限，范围：[0, 100]",
			"deadline":                "指定备份组到期时间，如果未传入该参数，默认为永久保留。示例：\"2024-04-26 10:00:00\"",
			"bucket_detail":           "创建COS备份时桶/前缀/桶所在地域映射关系，COS备份必传参数。",
			"delete_retreated_backup": "是否删除回档的备份数据，用于删除资源时使用",
			"backup_id":               "备份ID",
		},
	})
}

func resourceTencentCloudBrcBackupResource() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudBrcBackupResourceCreate,
		Read:        resourceTencentCloudBrcBackupResourceRead,
		Delete:      resourceTencentCloudBrcBackupResourceDelete,
		Description: "Provides a resource to create and manage BRC backup resources for COS, CSP, and database backups.",

		Schema: map[string]*schema.Schema{
			"resource_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Resource type for backup creation. Valid values: COS, CSP, MySQL_MariaDB, TDSQL_MySQL. Make sure the resource is deployed in BRC.",
				//ValidateFunc: validateAllowedStringValue(CreateResourceTypes),
			},
			"backup_name": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Name of the backup.",
			},
			"resource_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of the resource to be backed up.",
			},
			"deadline": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Backup expiration time. If not specified, the backup will be kept permanently. Example format: '2024-04-26 10:00:00'.",
			},
			"create_speed": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "Bandwidth limit for backup creation, range: [0, 100] Mbps.",
			},
			"bucket_detail": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Bucket details mapping for COS backup. Required for COS resource type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Required:    true,
							ForceNew:    true,
							Type:        schema.TypeString,
							Description: "COS bucket ID.",
						},
						"cos_region": {
							Required:    true,
							ForceNew:    true,
							Type:        schema.TypeString,
							Description: "COS bucket region.",
						},
						"prefix": {
							Optional:    true,
							ForceNew:    true,
							Type:        schema.TypeString,
							Description: "COS bucket prefix path.",
						},
					},
				},
			},
			"delete_retreated_backup": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether to delete the rolled-back backup data, which is used when deleting resources.",
			},
			// Computed fields
			"backup_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "ID of the created backup.",
			},
		},
	}
}

func resourceTencentCloudBrcBackupResourceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_resource.create")()

	var (
		logId        = getLogId(contextNil)
		ctx          = context.WithValue(context.TODO(), logIdKey, logId)
		brcService   = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceType *string
		resourceId   *string
		backupName   *string
		deadline     *time.Time
		createSpeed  *uint64
		bucketDetail *brc.BucketDetail
	)

	if v, ok := d.GetOk("resource_type"); ok {
		resourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_id"); ok {
		resourceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("backup_name"); ok {
		backupName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("deadline"); ok {
		layout := "2006-01-02 15:04:05"
		deadlineTime, err := time.Parse(layout, v.(string))
		if err != nil {
			return fmt.Errorf("failed to parse deadline: %s", err)
		}
		deadline = &deadlineTime
	}

	if v, ok := d.GetOk("create_speed"); ok {
		createSpeed = helper.Uint64(uint64(v.(int)))
	}

	// 对于对象存储类型，bucket_detail是必需的
	if v, ok := d.GetOk("bucket_detail"); ok {
		bucketDetail = &brc.BucketDetail{}
		bucketDetails := v.([]interface{})
		if len(bucketDetails) > 0 {
			detail := bucketDetails[0].(map[string]interface{})
			if v, ok := detail["cos_region"]; ok {
				bucketDetail.CosRegion = helper.String(v.(string))
			}
			if v, ok := detail["prefix"]; ok {
				bucketDetail.Prefix = helper.String(v.(string))
			}
			if v, ok := detail["resource_id"]; ok {
				bucketDetail.ResourceId = helper.String(v.(string))
			}
		}
	}

	backupId, err := brcService.CreateResourceBackup(ctx, resourceType, resourceId, backupName, deadline, createSpeed, bucketDetail)

	if err != nil {
		return err
	}

	d.SetId(backupId)
	return resourceTencentCloudBrcBackupResourceRead(d, meta)
}

func resourceTencentCloudBrcBackupResourceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_resource.read")()

	return nil
}

func resourceTencentCloudBrcBackupResourceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_backup_resource.delete")()

	var (
		logId                 = getLogId(contextNil)
		ctx                   = context.WithValue(context.TODO(), logIdKey, logId)
		brcService            = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		backupId              = helper.String(d.Id())
		deleteRetreatedBackup *bool
	)

	if v, ok := d.GetOk("delete_retreated_backup"); ok {
		deleteRetreatedBackup = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := brcService.DeleteResourceBackups(ctx, []*string{backupId}, deleteRetreatedBackup)
		if err != nil {
			return retryError(err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

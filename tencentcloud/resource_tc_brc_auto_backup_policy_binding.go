/*
Provides a resource to create a brc auto_snapshot_policy binding

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_auto_backup_policy_binding" "example" {
	  auto_backup_policy_id = cloud_brc_autobackup_policy.example.id
	  instance_ids          = ["ins-21ahx7qj"]
	  resource_type         = "INSTANCE"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_auto_backup_policy_binding", CNDescription{
		TerraformTypeCN: "自动备份策略绑定",
		DescriptionCN:   "提供BRC自动备份策略绑定资源，用于将自动备份策略绑定到指定实例。",
		AttributesCN: map[string]string{
			"auto_backup_policy_id": "自动备份策略ID",
			"instance_ids":          "实例ID列表",
			"resource_type":         "资源类型",
			"disk_ids":              "要绑定的云硬盘ID列表。",
			"file_system_ids":       "要绑定的文件系统ID列表。",
			"resource_ids":          "要绑定的实例资源ID列表。",
			"bucket_details":        "需要绑定的备份源桶/前缀/桶所属地域对应关系映射，用于COS定期备份",
		},
	})
}

func resourceTencentCloudBrcAutoBackupPolicyBinding() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudBrcAutoBackupPolicyBindingCreate,
		Read:        resourceTencentCloudBrcAutoBackupPolicyBindingRead,
		Delete:      resourceTencentCloudBrcAutoBackupPolicyBindingDelete,
		Description: "Provides a resource to create and manage BRC auto backup policy binding",

		Schema: map[string]*schema.Schema{
			"auto_backup_policy_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Auto backup policy ID.",
			},
			"instance_ids": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "Instance IDs to bind with the auto backup policy.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disk_ids": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "Disk IDs to bind with the auto backup policy.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"file_system_ids": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "File system IDs to bind with the auto backup policy.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_ids": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "Resource IDs to bind with the auto backup policy.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_type": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Resource type. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.",
				//ValidateFunc: validateAllowedStringValue(BackupResouceTypes),
			},
			"bucket_details": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "The mapping of corresponding relationships between the backup source bucket to be bound, prefix, and the region where the bucket belongs, which is used for regular COS backup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "cos bucket backup info, required for cos backup.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "cos bucket backup prefix, required for cos backup.",
						},
						"cos_region": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "cos bucket backup region, required for cos backup.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudBrcAutoBackupPolicyBindingCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy_binding.create")()

	var (
		logId              = getLogId(contextNil)
		ctx                = context.WithValue(context.TODO(), logIdKey, logId)
		brcService         = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		autoBackupPolicyId *string
		resourceType       *string
		instanceIds        []*string
		diskIds            []*string
		fileSystemIds      []*string
		resourceIds        []*string
		bucketDetails      []*brc.BucketDetail
	)

	if v, ok := d.GetOk("auto_backup_policy_id"); ok {
		autoBackupPolicyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_type"); ok {
		resourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_ids"); ok {
		for _, id := range v.([]interface{}) {
			instanceIds = append(instanceIds, helper.String(id.(string)))
		}
	}

	// 处理 disk_ids
	if v, ok := d.GetOk("disk_ids"); ok {
		for _, id := range v.([]interface{}) {
			diskIds = append(diskIds, helper.String(id.(string)))
		}
	}

	// 处理 file_system_ids
	if v, ok := d.GetOk("file_system_ids"); ok {
		for _, id := range v.([]interface{}) {
			fileSystemIds = append(fileSystemIds, helper.String(id.(string)))
		}
	}

	// 处理 resource_ids
	if v, ok := d.GetOk("resource_ids"); ok {
		for _, id := range v.([]interface{}) {
			resourceIds = append(resourceIds, helper.String(id.(string)))
		}
	}

	// 处理 bucket_details
	if v, ok := d.GetOk("bucket_details"); ok {
		bucketDetailsList := v.([]interface{})
		for _, item := range bucketDetailsList {
			detail := item.(map[string]interface{})
			bucketDetail := &brc.BucketDetail{}

			if resourceId, ok := detail["resource_id"]; ok {
				bucketDetail.ResourceId = &[]string{resourceId.(string)}[0]
			}
			if prefix, ok := detail["prefix"]; ok {
				bucketDetail.Prefix = &[]string{prefix.(string)}[0]
			}
			if cosRegion, ok := detail["cos_region"]; ok {
				bucketDetail.CosRegion = &[]string{cosRegion.(string)}[0]
			}

			bucketDetails = append(bucketDetails, bucketDetail)
		}
	}

	err := brcService.BindAutoBackupPolicy(ctx, instanceIds, diskIds, fileSystemIds, resourceIds, bucketDetails, autoBackupPolicyId, resourceType)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%s", *autoBackupPolicyId, uuid.New().String()))

	return nil
}

func resourceTencentCloudBrcAutoBackupPolicyBindingRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy_binding.read")()
	return nil
}

func resourceTencentCloudBrcAutoBackupPolicyBindingDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy_binding.delete")()

	var (
		logId              = getLogId(contextNil)
		ctx                = context.WithValue(context.TODO(), logIdKey, logId)
		brcService         = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		autoBackupPolicyId *string
		resourceType       *string
		instanceIds        []*string
		diskIds            []*string
		fileSystemIds      []*string
		resourceIds        []*string
		bucketDetails      []*brc.BucketDetail
	)

	if v, ok := d.GetOk("auto_backup_policy_id"); ok {
		autoBackupPolicyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_type"); ok {
		resourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_ids"); ok {
		for _, id := range v.([]interface{}) {
			instanceIds = append(instanceIds, helper.String(id.(string)))
		}
	}

	// 处理 disk_ids
	if v, ok := d.GetOk("disk_ids"); ok {
		for _, id := range v.([]interface{}) {
			diskIds = append(diskIds, helper.String(id.(string)))
		}
	}

	// 处理 file_system_ids
	if v, ok := d.GetOk("file_system_ids"); ok {
		for _, id := range v.([]interface{}) {
			fileSystemIds = append(fileSystemIds, helper.String(id.(string)))
		}
	}

	// 处理 resource_ids
	if v, ok := d.GetOk("resource_ids"); ok {
		for _, id := range v.([]interface{}) {
			resourceIds = append(resourceIds, helper.String(id.(string)))
		}
	}

	// 处理 bucket_details
	if v, ok := d.GetOk("bucket_details"); ok {
		bucketDetailsList := v.([]interface{})
		for _, item := range bucketDetailsList {
			detail := item.(map[string]interface{})
			bucketDetail := &brc.BucketDetail{}

			if resourceId, ok := detail["resource_id"]; ok {
				bucketDetail.ResourceId = &[]string{resourceId.(string)}[0]
			}
			if prefix, ok := detail["prefix"]; ok {
				bucketDetail.Prefix = &[]string{prefix.(string)}[0]
			}
			if cosRegion, ok := detail["cos_region"]; ok {
				bucketDetail.CosRegion = &[]string{cosRegion.(string)}[0]
			}

			bucketDetails = append(bucketDetails, bucketDetail)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := brcService.UnbindAutoBackupPolicy(ctx, instanceIds, diskIds, fileSystemIds, resourceIds, bucketDetails, autoBackupPolicyId, resourceType)
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

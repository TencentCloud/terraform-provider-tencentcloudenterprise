/*
Provides a resource to activate a brc backup service

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_activate_backup_service" "activate_backup_service" {
	  resource_type = "COS"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_activate_backup_service", CNDescription{
		TerraformTypeCN: "开通BRC服务",
		DescriptionCN:   "提供BRC备份服务开通资源，用于激活指定资源类型的备份服务。",
		AttributesCN: map[string]string{
			"resource_type":             "要开通备份服务的产品类型。取值范围：DISK：云硬盘,CFS：文件存储,CVM：云服务器,TDSQL_MySQL：TDSQL_MySQL版,MySQL_MariaDB：关系型数据库(MySQL_MariaDB),COS：对象存储。",
			"brc_backup_service_status": "BRC服务状态",
		},
	})
}
func resourceTencentCloudBrcActivateBackupService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a brc activate_backup_service.",
		Create:      resourceTencentCloudBrcActivateBackupServiceCreate,
		Read:        resourceTencentCloudBrcActivateBackupServiceRead,
		Delete:      resourceTencentCloudBrcActivateBackupServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Required: true,
				ForceNew: true,
				//ValidateFunc: validateAllowedStringValue(BackupResouceTypes),
				Type:        schema.TypeString,
				Description: "The resource type to be backed up. Valid values: 'CFS','COS', 'CSP', 'DISK', 'INSTANCE', 'MySQL_MariaDB', 'TDSQL_MySQL'",
			},
			"brc_backup_service_is_open": {
				Computed:    true,
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Current status of the BRC backup service for this user. Valid values: true, false.",
			},
		},
	}
}

func resourceTencentCloudBrcActivateBackupServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_activate_brc_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	resourceType := d.Get("resource_type").(string)
	request := brc.NewActivateBrcServiceRequest()
	request.ResourceType = helper.String(resourceType)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBrcClient().ActivateBackupService(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate brc activateBackupService failed, reason:%+v", logId, err)
		return nil
	}
	d.SetId(resourceType)

	return resourceTencentCloudBrcActivateBackupServiceRead(d, meta)
}

func resourceTencentCloudBrcActivateBackupServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_activate_brc_service.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	brcService := BRCService{meta.(*TencentCloudClient).apiV3Conn}
	resourceType := d.Id()
	isOpen := false
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := brcService.DescribeResourceBackupStatusByResourceType(ctx, resourceType)
		if e != nil {
			return retryError(e)
		}
		if result == nil {
			d.SetId("")
			return resource.NonRetryableError(fmt.Errorf("brc backup service for resource type %s not found", resourceType))
		}
		isOpen = *result.IsOpen
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate brc activate BrcService failed, reason:%+v", logId, err)
		return nil
	}

	_ = d.Set("brc_backup_service_is_open", isOpen)

	return nil
}

func resourceTencentCloudBrcActivateBackupServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_activate_brc_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

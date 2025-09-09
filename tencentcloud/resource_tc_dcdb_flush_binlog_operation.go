/*
Provides a resource to create a dcdb flush_binlog_operation

# Example Usage

```hcl

	resource "cloud_dcdb_flush_binlog_operation" "flush_operation" {
	  instance_id = local.dcdb_id
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("cloud_dcdb_flush_binlog_operation", CNDescription{
		TerraformTypeCN: "刷新binlog",
		DescriptionCN:   "提供DCDB刷新binlog操作资源，用于刷新DCDB实例的binlog。",
		AttributesCN: map[string]string{
			"instance_id": "实例ID",
		},
	})
}

func resourceTencentCloudDcdbFlushBinlogOperation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb flush_binlog_operation",
		Create:      resourceTencentCloudDcdbFlushBinlogOperationCreate,
		Read:        resourceTencentCloudDcdbFlushBinlogOperationRead,
		Delete:      resourceTencentCloudDcdbFlushBinlogOperationDelete,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},
		},
	}
}

func resourceTencentCloudDcdbFlushBinlogOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_flush_binlog_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = dcdb.NewFlushBinlogRequest()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().FlushBinlog(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate dcdb flushBinlogOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	return resourceTencentCloudDcdbFlushBinlogOperationRead(d, meta)
}

func resourceTencentCloudDcdbFlushBinlogOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_flush_binlog_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudDcdbFlushBinlogOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_flush_binlog_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

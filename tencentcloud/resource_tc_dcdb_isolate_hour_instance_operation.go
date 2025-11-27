/*
Provides a resource to create a dcdb isolate_hour_instance_operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_isolate_hour_instance_operation" "isolate_hour_instance_operation" {
	  instance_ids = local.dcdb_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_isolate_hour_instance_operation", CNDescription{
		TerraformTypeCN: "数据库实例隔离时间操作",
		DescriptionCN:   "提供DCDB实例隔离操作资源，用于隔离按小时计费的DCDB实例。",
		AttributesCN: map[string]string{
			"instance_id": "实例ID列表",
		},
	})
}

func resourceTencentCloudDcdbIsolateHourInstanceOperation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb isolate_hour_instance_operation",
		Create:      resourceTencentCloudDcdbIsolateHourInstanceOperationCreate,
		Read:        resourceTencentCloudDcdbIsolateHourInstanceOperationRead,
		Delete:      resourceTencentCloudDcdbIsolateHourInstanceOperationDelete,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance ID list.",
			},
		},
	}
}

func resourceTencentCloudDcdbIsolateHourInstanceOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_isolate_hour_instance_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = dcdb.NewIsolateHourDCDBInstanceRequest()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceIds = []*string{helper.String(instanceId)}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().IsolateHourDCDBInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate dcdb isolateHourInstanceOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	return resourceTencentCloudDcdbIsolateHourInstanceOperationRead(d, meta)
}

func resourceTencentCloudDcdbIsolateHourInstanceOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_isolate_hour_instance_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudDcdbIsolateHourInstanceOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_isolate_hour_instance_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

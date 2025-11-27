/*
Provides a resource to create a dcdb switch_db_instance_ha_operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_switch_db_instance_ha_operation" "switch_operation" {
	  instance_id = local.dcdb_id
	  zone = "ap-guangzhou-4" //3 to 4
	}

```
*/
package tencentcloud

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_switch_db_instance_ha_operation", CNDescription{
		TerraformTypeCN: "DCN容灾主备切换",
		DescriptionCN:   "提供DCDB容灾主备切换操作资源，用于执行DCDB实例的主备切换操作。",
		AttributesCN: map[string]string{
			"instance_id": "tdsqlshard-ow728lmc格式的实例ID",
			"zone":        "目标AZ中延迟最低的节点将自动升级为主节点",
		},
	})
}

func resourceTencentCloudDcdbSwitchDbInstanceHaOperation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb switch_db_instance_ha_operation",
		Create:      resourceTencentCloudDcdbSwitchDbInstanceHaOperationCreate,
		Read:        resourceTencentCloudDcdbSwitchDbInstanceHaOperationRead,
		Delete:      resourceTencentCloudDcdbSwitchDbInstanceHaOperationDelete,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance ID in the format of tdsqlshard-ow728lmc.",
			},

			"zone": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.",
			},
		},
	}
}

func resourceTencentCloudDcdbSwitchDbInstanceHaOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_switch_db_instance_ha_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = dcdb.NewSwitchDcnJobRequest()
		service    = DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId string
		flowId     *int64
	)
	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	//if v, ok := d.GetOk("zone"); ok {
	//	request.Zone = helper.String(v.(string))
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().SwitchDcnJob(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		flowId = result.Response.FlowId
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate dcdb switchDbInstanceHaOperation failed, reason:%+v", logId, err)
		return err
	}

	if flowId != nil {
		// need to wait init operation success
		// 0:success; 1:failed, 2:running
		conf := BuildStateChangeConf([]string{}, []string{"0"}, 3*readRetryTimeout, time.Second, service.DcdbDbInstanceStateRefreshFunc(flowId, []string{}))
		if _, e := conf.WaitForState(); e != nil {
			return e
		}
	}

	d.SetId(instanceId)

	return resourceTencentCloudDcdbSwitchDbInstanceHaOperationRead(d, meta)
}

func resourceTencentCloudDcdbSwitchDbInstanceHaOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_switch_db_instance_ha_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudDcdbSwitchDbInstanceHaOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_switch_db_instance_ha_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

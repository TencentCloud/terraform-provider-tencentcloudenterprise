/*
Provides a resource to accept cross-account CCN attachment requests.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_instances_accept_attach" "example" {
	  ccn_id = "ccn-39lqkygf"

	  instances {
	    instance_id     = "vpc-j9yhbzpn"
	    instance_region = "ap-guangzhou"
	    instance_type   = "VPC"
	  }
	}

```

# Import

CCN instance accept-attach actions can be imported using the CCN ID, e.g.

```bash
terraform import tencentcloudenterprise_ccn_instances_accept_attach.example ccn-39lqkygf
```
*/
package tencentcloud

import (
	"fmt"
	"log"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_instances_accept_attach", CNDescription{
		TerraformTypeCN: "云联网接受跨账号关联",
		DescriptionCN:   "用于云联网所有者接受跨账号网络实例关联申请。",
		AttributesCN: map[string]string{
			"ccn_id":          "云联网实例 ID",
			"instances":       "待接受的关联实例列表",
			"instance_id":     "关联实例 ID",
			"instance_region": "关联实例所属地域",
			"instance_type":   "关联实例类型",
		},
	})
}

func resourceTencentCloudCcnInstancesAcceptAttach() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to accept cross-account CCN attachment requests.",
		Create:      resourceTencentCloudCcnInstancesAcceptAttachCreate,
		Read:        resourceTencentCloudCcnInstancesAcceptAttachRead,
		Delete:      resourceTencentCloudCcnInstancesAcceptAttachDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "CCN instance ID.",
			},
			"instances": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "List of attachment instances to accept.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Attachment instance ID.",
						},
						"instance_region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Region of the attachment instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Attachment instance type. Valid values: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCcnInstancesAcceptAttachCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_accept_attach.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = ccn.NewAcceptAttachCcnInstancesRequest()
		ccnID   string
	)
	if v, ok := d.GetOk("ccn_id"); ok {
		ccnID = v.(string)
		request.CcnId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instances"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			ccnInstance := ccn.CcnInstance{}
			if v, ok := dMap["instance_id"]; ok {
				ccnInstance.InstanceId = helper.String(v.(string))
			}
			if v, ok := dMap["instance_region"]; ok {
				ccnInstance.InstanceRegion = helper.String(v.(string))
			}
			if v, ok := dMap["instance_type"]; ok {
				ccnInstance.InstanceType = helper.String(v.(string))
			}
			request.Instances = append(request.Instances, &ccnInstance)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().AcceptAttachCcnInstances(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return fmt.Errorf("[CRITAL]%s operate ccn ccnInstancesAcceptAttach failed, reason:%+v", logId, err)
	}

	d.SetId(ccnID)
	return resourceTencentCloudCcnInstancesAcceptAttachRead(d, meta)
}

func resourceTencentCloudCcnInstancesAcceptAttachRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_accept_attach.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCcnInstancesAcceptAttachDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_accept_attach.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

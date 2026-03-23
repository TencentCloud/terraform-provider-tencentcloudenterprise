/*
Provides a resource to reject cross-account CCN attachment requests.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_instances_reject_attach" "example" {
	  ccn_id = "ccn-39lqkygf"

	  instances {
	    instance_id     = "vpc-j9yhbzpn"
	    instance_region = "ap-guangzhou"
	    instance_type   = "VPC"
	  }
	}

```

# Import

CCN instance reject-attach actions can be imported using the CCN ID, e.g.

```bash
terraform import tencentcloudenterprise_ccn_instances_reject_attach.example ccn-39lqkygf
```
*/
package tencentcloud

import (
	"fmt"
	"log"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_instances_reject_attach", CNDescription{
		TerraformTypeCN: "云联网拒绝跨账号关联",
		DescriptionCN:   "用于云联网所有者拒绝跨账号网络实例关联申请。",
		AttributesCN: map[string]string{
			"ccn_id":          "云联网实例 ID",
			"instances":       "待拒绝的关联实例列表",
			"instance_id":     "关联实例 ID",
			"instance_region": "关联实例所属地域",
			"instance_type":   "关联实例类型",
		},
	})
}

func resourceTencentCloudCcnInstancesRejectAttach() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to reject cross-account CCN attachment requests.",
		Create:      resourceTencentCloudCcnInstancesRejectAttachCreate,
		Read:        resourceTencentCloudCcnInstancesRejectAttachRead,
		Delete:      resourceTencentCloudCcnInstancesRejectAttachDelete,
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
				Description: "List of attachment instances to reject.",
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

func resourceTencentCloudCcnInstancesRejectAttachCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reject_attach.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = vpc.NewRejectAttachCcnInstancesRequest()
		ccnID   string
	)
	if v, ok := d.GetOk("ccn_id"); ok {
		ccnID = v.(string)
		request.CcnId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instances"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			ccnInstance := vpc.CcnInstance{}
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
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().RejectAttachCcnInstances(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return fmt.Errorf("[CRITAL]%s operate vpc ccnInstancesRejectAttach failed, reason:%+v", logId, err)
	}

	d.SetId(ccnID)
	return resourceTencentCloudCcnInstancesRejectAttachRead(d, meta)
}

func resourceTencentCloudCcnInstancesRejectAttachRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reject_attach.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCcnInstancesRejectAttachDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reject_attach.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

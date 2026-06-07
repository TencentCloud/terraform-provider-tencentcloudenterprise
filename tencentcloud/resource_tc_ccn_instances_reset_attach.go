/*
Provides a resource to reset expired cross-account CCN attachment requests.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_instances_reset_attach" "example" {
	  ccn_id  = "ccn-39lqkygf"
	  ccn_uin = "100022975249"

	  instances {
	    instance_id     = "vpc-j9yhbzpn"
	    instance_region = "ap-guangzhou"
	    instance_type   = "VPC"
	  }
	}

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
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_instances_reset_attach", CNDescription{
		TerraformTypeCN: "云联网重置跨账号关联申请",
		DescriptionCN:   "用于在跨账号云联网关联申请过期后重新发起关联申请。",
		AttributesCN: map[string]string{
			"ccn_id":          "云联网实例 ID",
			"ccn_uin":         "云联网所属账号 UIN",
			"instances":       "待重置的关联实例列表",
			"instance_id":     "关联实例 ID",
			"instance_region": "关联实例所属地域",
			"instance_type":   "关联实例类型",
		},
	})
}

func resourceTencentCloudCcnInstancesResetAttach() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to reset expired cross-account CCN attachment requests.",
		Create:      resourceTencentCloudCcnInstancesResetAttachCreate,
		Read:        resourceTencentCloudCcnInstancesResetAttachRead,
		Delete:      resourceTencentCloudCcnInstancesResetAttachDelete,
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "CCN instance ID.",
			},
			"ccn_uin": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Root account UIN of the target CCN.",
			},
			"instances": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "List of attachment instances to reset.",
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

func resourceTencentCloudCcnInstancesResetAttachCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reset_attach.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = ccn.NewResetAttachCcnInstancesRequest()
		ccnID   string
	)
	if v, ok := d.GetOk("ccn_id"); ok {
		ccnID = v.(string)
		request.CcnId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ccn_uin"); ok {
		request.CcnUin = helper.String(v.(string))
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
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().ResetAttachCcnInstances(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return fmt.Errorf("[CRITAL]%s operate ccn ccnInstancesResetAttach failed, reason:%+v", logId, err)
	}

	d.SetId(ccnID)
	return resourceTencentCloudCcnInstancesResetAttachRead(d, meta)
}

func resourceTencentCloudCcnInstancesResetAttachRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reset_attach.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCcnInstancesResetAttachDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_instances_reset_attach.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

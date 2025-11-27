/*
Provides a resource to create a as protect_instances

# Example Usage

```hcl

	resource "tencentcloudenterprise_as_protect_instances" "protect_instances" {
	  auto_scaling_group_id = cloud_as_scaling_group.scaling_group.id
	  instance_ids = ["ins-xxxxx"]
	  protected_from_scale_in = true
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	as "terraform-provider-tencentcloudenterprise/sdk/as/v20180419"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_as_protect_instances", CNDescription{
		TerraformTypeCN: "自动伸缩组保护实例",
		DescriptionCN:   "提供自动伸缩保护实例资源，用于创建和管理AS保护实例。",
		AttributesCN: map[string]string{
			"auto_scaling_group_id":   "伸缩组ID",
			"instance_ids":            "实例ID列表",
			"protected_from_scale_in": "是否保护实例",
		},
	})
}
func resourceTencentCloudAsProtectInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a as protect_instances",
		Create:      resourceTencentCloudAsProtectInstancesCreate,
		Read:        resourceTencentCloudAsProtectInstancesRead,
		Delete:      resourceTencentCloudAsProtectInstancesDelete,
		Schema: map[string]*schema.Schema{
			"auto_scaling_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Launch configuration ID.",
			},

			"instance_ids": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of cvm instances to remove.",
			},

			"protected_from_scale_in": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Description: "If instances need protect.",
			},
		},
	}
}

func resourceTencentCloudAsProtectInstancesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_as_protect_instances.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = as.NewSetInstancesProtectionRequest()
	)
	if v, ok := d.GetOk("auto_scaling_group_id"); ok {
		request.AutoScalingGroupId = helper.String(v.(string))
	}
	ids := make([]string, 0)
	if v, ok := d.GetOk("instance_ids"); ok {
		instanceIdsSet := v.(*schema.Set).List()
		for i := range instanceIdsSet {
			instanceIds := instanceIdsSet[i].(string)
			ids = append(ids, instanceIds)
			request.InstanceIds = append(request.InstanceIds, &instanceIds)
		}
	}

	if v, _ := d.GetOk("protected_from_scale_in"); v != nil {
		request.ProtectedFromScaleIn = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseAsClient().SetInstancesProtection(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate as protectInstances failed, reason:%+v", logId, err)
		return nil
	}

	// 需要 setId，可以通过InstancesId作为ID
	d.SetId(helper.DataResourceIdsHash(ids))

	return resourceTencentCloudAsProtectInstancesRead(d, meta)
}

func resourceTencentCloudAsProtectInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_protect_instances.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudAsProtectInstancesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_protect_instances.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

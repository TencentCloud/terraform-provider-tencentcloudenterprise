/*
Provides a resource to create a as remove_instances

# Example Usage

```hcl

	resource "tencentcloudenterprise_as_remove_instances" "remove_instances" {
	  auto_scaling_group_id = cloud_as_scaling_group.scaling_group.id
	  instance_ids = ["ins-xxxxxx"]
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
	registerResourceDescriptionProvider("tencentcloudenterprise_as_remove_instances", CNDescription{
		TerraformTypeCN: "自动伸缩组移除实例",
		DescriptionCN:   "提供自动伸缩移除实例资源，用于创建和管理AS移除实例。",
		AttributesCN: map[string]string{
			"auto_scaling_group_id": "伸缩组ID",
			"instance_ids":          "实例ID列表",
		},
	})
}
func resourceTencentCloudAsRemoveInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a as remove_instances",
		Create:      resourceTencentCloudAsRemoveInstancesCreate,
		Read:        resourceTencentCloudAsRemoveInstancesRead,
		Delete:      resourceTencentCloudAsRemoveInstancesDelete,
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
		},
	}
}

func resourceTencentCloudAsRemoveInstancesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_as_remove_instances.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = as.NewDetachInstancesRequest()
		response   = as.NewDetachInstancesResponse()
		activityId string
	)
	if v, ok := d.GetOk("auto_scaling_group_id"); ok {
		request.AutoScalingGroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_ids"); ok {
		instanceIdsSet := v.(*schema.Set).List()
		for i := range instanceIdsSet {
			instanceIds := instanceIdsSet[i].(string)
			request.InstanceIds = append(request.InstanceIds, &instanceIds)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseAsClient().DetachInstances(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate as removeInstances failed, reason:%+v", logId, err)
		return nil
	}

	activityId = *response.Response.ActivityId
	d.SetId(activityId)

	return resourceTencentCloudAsRemoveInstancesRead(d, meta)
}

func resourceTencentCloudAsRemoveInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_remove_instances.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudAsRemoveInstancesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_as_remove_instances.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

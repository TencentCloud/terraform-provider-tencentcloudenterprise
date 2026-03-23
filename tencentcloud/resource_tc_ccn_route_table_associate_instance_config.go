/*
Provides a resource to manage the instance associations of a CCN route table.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_route_table_associate_instance_config" "example" {
	  ccn_id         = "ccn-cwm743gl"
	  route_table_id = "ccnrtb-1ydgdxt1"

	  instances {
	    instance_id   = "vpc-ayl8ggap"
	    instance_type = "VPC"
	  }
	}

```

# Import

CCN route table associate instance config can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table_associate_instance_config.example ccn-cwm743gl#ccnrtb-1ydgdxt1
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_route_table_associate_instance_config", CNDescription{
		TerraformTypeCN: "云联网路由表关联实例配置",
		DescriptionCN:   "用于管理云联网路由表关联的实例集合。",
		AttributesCN: map[string]string{
			"ccn_id":         "云联网实例 ID",
			"route_table_id": "云联网路由表 ID",
			"instances":      "实例列表",
			"instance_id":    "实例 ID",
			"instance_type":  "实例类型",
		},
	})
}

func resourceTencentCloudCcnRouteTableAssociateInstanceConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage the instance associations of a CCN route table.",
		Create:      resourceTencentCloudCcnRouteTableAssociateInstanceConfigCreate,
		Read:        resourceTencentCloudCcnRouteTableAssociateInstanceConfigRead,
		Update:      resourceTencentCloudCcnRouteTableAssociateInstanceConfigUpdate,
		Delete:      resourceTencentCloudCcnRouteTableAssociateInstanceConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the CCN.",
			},
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CCN route table ID.",
			},
			"instances": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Associated instance list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Instance ID.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Instance type. Valid values include `VPC`, `DIRECTCONNECT`, `BMVPC`, `EDGE`, `EDGE_TUNNEL`, `EDGE_VPNGW`, `VPNGW`.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCcnRouteTableAssociateInstanceConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_associate_instance_config.create")()
	defer inconsistentCheck(d, meta)()

	ccnId := d.Get("ccn_id").(string)
	routeTableId := d.Get("route_table_id").(string)
	d.SetId(strings.Join([]string{ccnId, routeTableId}, FILED_SP))

	return resourceTencentCloudCcnRouteTableAssociateInstanceConfigUpdate(d, meta)
}

func resourceTencentCloudCcnRouteTableAssociateInstanceConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_associate_instance_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	ccnId := items[0]
	routeTableId := items[1]

	instanceBindList, err := service.DescribeRouteTableAssociatedInstancesById(ctx, ccnId, routeTableId)
	if err != nil {
		return err
	}
	if len(instanceBindList) == 0 {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_ccn_route_table_associate_instance_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("ccn_id", ccnId)
	_ = d.Set("route_table_id", routeTableId)

	result := make([]map[string]interface{}, 0, len(instanceBindList))
	for _, instanceBind := range instanceBindList {
		result = append(result, map[string]interface{}{
			"instance_id":   instanceBind.instanceId,
			"instance_type": instanceBind.instanceType,
		})
	}
	_ = d.Set("instances", result)

	return nil
}

func resourceTencentCloudCcnRouteTableAssociateInstanceConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_associate_instance_config.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	request := ccn.NewAssociateInstancesToCcnRouteTableRequest()
	request.CcnId = helper.String(items[0])
	request.RouteTableId = helper.String(items[1])

	for _, item := range d.Get("instances").(*schema.Set).List() {
		instanceMap := item.(map[string]interface{})
		request.Instances = append(request.Instances, &ccn.CcnInstanceWithoutRegion{
			InstanceId:   helper.String(instanceMap["instance_id"].(string)),
			InstanceType: helper.String(instanceMap["instance_type"].(string)),
		})
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().AssociateInstancesToCcnRouteTable(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update AssociateInstancesToCcnRouteTable failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCcnRouteTableAssociateInstanceConfigRead(d, meta)
}

func resourceTencentCloudCcnRouteTableAssociateInstanceConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_associate_instance_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

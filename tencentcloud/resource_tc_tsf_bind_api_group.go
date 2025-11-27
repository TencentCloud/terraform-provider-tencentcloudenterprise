/*
Provides a resource to create a tsf bind_api_group

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_bind_api_group" "bind_api_group" {
	  gateway_deploy_group_id = "group-vzd97zpy"
	  group_id = "grp-qp0rj3zi"
	}

```

# Import

tsf bind_api_group can be imported using the id, e.g.

```
terraform import cloud_tsf_bind_api_group.bind_api_group bind_api_group_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_bind_api_group", CNDescription{
		TerraformTypeCN: "API 分组 绑定网关部署组",
		DescriptionCN:   "提供TSF API分组 绑定网关部署组资源，用于创建和管理API分组 绑定网关部署组。",
		AttributesCN: map[string]string{
			"gateway_deploy_group_id": "网关部署组ID",
			"group_id":                "分组ID",
		},
	})
}

func resourceTencentCloudTsfBindApiGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf bind_api_group",
		Create:      resourceTencentCloudTsfBindApiGroupCreate,
		Read:        resourceTencentCloudTsfBindApiGroupRead,
		Delete:      resourceTencentCloudTsfBindApiGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gateway_deploy_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Gateway group id.",
			},

			"group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Group id.",
			},
		},
	}
}

func resourceTencentCloudTsfBindApiGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_bind_api_group.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = tsf.NewBindApiGroupRequest()
		groupId              string
		gatewayDeployGroupId string
	)
	if v, ok := d.GetOk("gateway_deploy_group_id"); ok {
		gatewayDeployGroupId = v.(string)
	}

	if v, ok := d.GetOk("group_id"); ok {
		groupId = v.(string)
	}
	request.GroupGatewayList = []*tsf.GatewayGroupIds{
		{
			GatewayDeployGroupId: &gatewayDeployGroupId,
			GroupId:              &groupId,
		},
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().BindApiGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf bindApiGroup failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{groupId, gatewayDeployGroupId}, FILED_SP))

	return resourceTencentCloudTsfBindApiGroupRead(d, meta)
}

func resourceTencentCloudTsfBindApiGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_bind_api_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	groupId := idSplit[0]
	gatewayDeployGroupId := idSplit[1]

	bindApiGroup, err := service.DescribeTsfBindApiGroupById(ctx, groupId, gatewayDeployGroupId)
	if err != nil {
		return err
	}

	if bindApiGroup == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfBindApiGroup` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("gateway_deploy_group_id", gatewayDeployGroupId)
	_ = d.Set("group_id", groupId)

	return nil
}

func resourceTencentCloudTsfBindApiGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_bind_api_group.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	groupId := idSplit[0]
	gatewayDeployGroupId := idSplit[1]

	if err := service.DeleteTsfBindApiGroupById(ctx, groupId, gatewayDeployGroupId); err != nil {
		return err
	}

	return nil
}

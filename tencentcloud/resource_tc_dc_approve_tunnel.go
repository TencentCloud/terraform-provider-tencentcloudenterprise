/*
Provides a resource to approve a dc tunnel

Example Usage

```hcl
resource "tencentcloudenterprise_dc_approve_tunnel" "example {
	direct_connect_tunnel_name = "test-tunnel"
	direct_connect_tunnel_name = "dcx-ua8ej92a"
	approved				   = true
	comments				   = "ok"
}
```

Import

approve_tunnel can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_dc_approve_tunnel.example
```
*/
package tencentcloud

import (
	dc "terraform-provider-tencentcloudenterprise/sdk/dc/v20180410"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dc_approve_tunnel", CNDescription{
		TerraformTypeCN: "审批DC接入请求",
		AttributesCN: map[string]string{
			"direct_connect_tunnel_name":       "专线通道名称",
			"direct_connect_tunnel_id": 		"专线通道ID",
			"approved":							"审批结果, True: 同意， False: 不同意",
			"comments":							"备注",
		},
	})
}
func resourceTencentCloudDcApproveTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dc instance",
		Create:      resourceTencentCloudDcApproveTunnelCreate,
		Read:        resourceTencentCloudDcApproveTunnelRead,
		//Update:      resourceTencentCloudDcApproveTunnelUpdate,
		Delete:      resourceTencentCloudDcApproveTunnelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"direct_connect_tunnel_name": {
				Required:    true,
				ForceNew:	 true,
				Type:        schema.TypeString,
				Description: "Connection name.",
			},
			"direct_connect_tunnel_id": {
				Required:    true,
				ForceNew:	 true,
				Type:        schema.TypeString,
				Description: "ID of direct connect tunnel.",
			},
			"approved":{
				Required:    true,
				ForceNew:	 true,
				Type:        schema.TypeBool,
				Description: "Approval result for the resource: `true` (approved), `false` (disapproved).",
			},
			"comments": {
				Required:    true,
				ForceNew:	 true,
				Type:        schema.TypeString,
				Description: "Comment of the approval.",
			},
		},
	}
}

func resourceTencentCloudDcApproveTunnelCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_approve_tunnel.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = dc.NewApproveDirectConnectTunnelRequest()
		response = dc.NewApproveDirectConnectTunnelResponse()
	)
	if v, ok := d.GetOk("direct_connect_tunnel_name"); ok {
		request.DirectConnectTunnelName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("direct_connect_tunnel_id"); ok {
		request.DirectConnectTunnelId = helper.String(v.(string))
	}

	request.Approved = helper.Bool(d.Get("approved").(bool))

	if v, ok := d.GetOk("comments"); ok {
		request.Comments = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcClient().ApproveDirectConnectTunnel(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		requestId := ""
		if response != nil && response.Response != nil && response.Response.RequestId != nil {
			requestId = *response.Response.RequestId
		}
		log.Printf("[CRITAL]%s approve dc tunnel failed, reason:%+v, requestId: %s", logId, err, requestId)
		return err
	}

	// 设置资源ID为tunnel ID，用于Terraform状态跟踪
	if tunnelId, ok := d.GetOk("direct_connect_tunnel_id"); ok {
		d.SetId(tunnelId.(string))
	}

	return resourceTencentCloudDcApproveTunnelRead(d, meta)
}

func resourceTencentCloudDcApproveTunnelRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_approve_tunnel.read")()
	defer inconsistentCheck(d, meta)()

	// 审批操作是一次性的，无需读取状态
	// 如果资源ID存在，说明审批已完成
	if d.Id() == "" {
		return nil
	}

	return nil
}

//func resourceTencentCloudDcApproveTunnelUpdate(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("resource.tencentcloudenterprise_dc_approve_tunnel.update")()
//	defer inconsistentCheck(d, meta)()
//
//	return resourceTencentCloudDcInstanceRead(d, meta)
//}

func resourceTencentCloudDcApproveTunnelDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_approve_tunnel.delete")()
	defer inconsistentCheck(d, meta)()

	// 审批操作无法撤销，只需清除Terraform状态
	d.SetId("")
	return nil
}

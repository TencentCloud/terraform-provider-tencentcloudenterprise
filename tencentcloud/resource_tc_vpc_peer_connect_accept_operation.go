/*
Provides a resource to accept a same-region cross-account VPC peering connection.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_peer_connect_accept_operation" "example" {
	  peering_connection_id = "pcx-1asg3t63"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_peer_connect_accept_operation", CNDescription{
		TerraformTypeCN: "VPC同地域跨账号对等连接接受操作",
		DescriptionCN:   "提供VPC同地域跨账号对等连接接受操作资源，用于接受对端账号发起的同地域对等连接请求。",
		AttributesCN: map[string]string{
			"peering_connection_id": "对等连接唯一ID",
		},
	})
}

func resourceTencentCloudVpcPeerConnectAcceptOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudVpcPeerConnectAcceptOperationCreate,
		Read:        resourceTencentCloudVpcPeerConnectAcceptOperationRead,
		Delete:      resourceTencentCloudVpcPeerConnectAcceptOperationDelete,
		Description: "Provides a resource to accept a same-region cross-account VPC peering connection",
		Schema: map[string]*schema.Schema{
			"peering_connection_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The unique ID of the peering connection.",
			},
		},
	}
}

func resourceTencentCloudVpcPeerConnectAcceptOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencenttencentcloudenterprise_vpc_peer_connect_accept_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewAcceptVpcPeeringConnectionRequest()

	peeringConnectionId := d.Get("peering_connection_id").(string)
	request.PeeringConnectionId = helper.String(peeringConnectionId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AcceptVpcPeeringConnection(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s accept vpc PeerConnectOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(peeringConnectionId)

	return resourceTencentCloudVpcPeerConnectAcceptOperationRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectAcceptOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencenttencentcloudenterprise_vpc_peer_connect_accept_operation.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	peeringConnectionId := d.Id()

	peerConnection, err := service.DescribeVpcPeerConnectManagerById(ctx, peeringConnectionId)
	if err != nil {
		return err
	}

	if peerConnection == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcPeerConnectAcceptOperation` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	return nil
}

func resourceTencentCloudVpcPeerConnectAcceptOperationDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

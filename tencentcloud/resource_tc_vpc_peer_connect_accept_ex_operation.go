/*
Provides a resource to accept a cross-region cross-account VPC peering connection.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_peer_connect_accept_ex_operation" "example" {
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
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_peer_connect_accept_ex_operation", CNDescription{
		TerraformTypeCN: "VPC跨地域跨账号对等连接接受操作",
		DescriptionCN:   "提供VPC跨地域跨账号对等连接接受操作资源，用于接受对端账号发起的跨地域对等连接请求。",
		AttributesCN: map[string]string{
			"peering_connection_id": "对等连接唯一ID",
		},
	})
}

func resourceTencentCloudVpcPeerConnectAcceptExOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudVpcPeerConnectAcceptExOperationCreate,
		Read:        resourceTencentCloudVpcPeerConnectAcceptExOperationRead,
		Delete:      resourceTencentCloudVpcPeerConnectAcceptExOperationDelete,
		Description: "Provides a resource to accept a cross-region cross-account VPC peering connection",
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

func resourceTencentCloudVpcPeerConnectAcceptExOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencenttencentcloudenterprise_vpc_peer_connect_accept_ex_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewAcceptVpcPeeringConnectionExRequest()

	peeringConnectionId := d.Get("peering_connection_id").(string)
	request.PeeringConnectionId = helper.String(peeringConnectionId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AcceptVpcPeeringConnectionEx(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s accept vpc PeerConnectExOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(peeringConnectionId)

	return resourceTencentCloudVpcPeerConnectAcceptExOperationRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectAcceptExOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencenttencentcloudenterprise_vpc_peer_connect_accept_ex_operation.read")()
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
		log.Printf("[WARN]%s resource `VpcPeerConnectAcceptExOperation` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	return nil
}

func resourceTencentCloudVpcPeerConnectAcceptExOperationDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

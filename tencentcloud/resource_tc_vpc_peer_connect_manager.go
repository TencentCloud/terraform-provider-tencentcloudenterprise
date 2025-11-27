/*
Provides a resource to create and manage a VPC peering connection.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_peer_connect_manager" "foo" {
	  vpc_id                  = "vpc-4xxr2cy7"
	  peering_connection_name = "test_peer_connection"
	  peer_vpc_id             = "vpc-5ggr3dx8"
	  peer_uin                = "100001234567"
	  peer_region             = "ap-beijing"
	}

```

# Import

VPC peering connection can be imported using the id, e.g.

```
$ terraform import cloud_vpc_peer_connect_manager.foo pcx-1asg3t63
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_peer_connect_manager", CNDescription{
		TerraformTypeCN: "VPC对等连接管理",
		DescriptionCN:   "提供VPC对等连接管理资源，用于创建和管理VPC对等连接。",
		AttributesCN: map[string]string{
			"vpc_id":                  "本端VPC ID",
			"peering_connection_name": "对等连接名称",
			"peer_vpc_id":             "对端VPC ID",
			"peer_uin":                "对端用户UIN",
			"peer_region":             "对端地域",
		},
	})
}

func resourceTencentCloudVpcPeerConnectManager() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudVpcPeerConnectManagerCreate,
		Read:        resourceTencentCloudVpcPeerConnectManagerRead,
		Update:      resourceTencentCloudVpcPeerConnectManagerUpdate,
		Delete:      resourceTencentCloudVpcPeerConnectManagerDelete,
		Description: "Provides a resource to create and manage VPC peering connection",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "The unique ID of the local VPC.",
			},

			"peering_connection_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Peer connection name.",
			},

			"peer_vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "The unique ID of the peer VPC.",
			},

			"peer_uin": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Peer user UIN.",
			},
			"peer_region": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Peer region.",
			},
		},
	}
}

func resourceTencentCloudVpcPeerConnectManagerCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_manager.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request             = vpc.NewCreateVpcPeeringConnectionRequest()
		response            = vpc.NewCreateVpcPeeringConnectionResponse()
		peeringConnectionId string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		request.SourceVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peering_connection_name"); ok {
		request.PeeringConnectionName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_vpc_id"); ok {
		request.DestinationVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_uin"); ok {
		request.DestinationUin = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_region"); ok {
		request.DestinationRegion = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CreateVpcPeeringConnection(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc PeerConnectManager failed, reason:%+v", logId, err)
		return err
	}

	peeringConnectionId = *response.Response.PeeringConnectionId
	d.SetId(peeringConnectionId)

	return resourceTencentCloudVpcPeerConnectManagerRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectManagerRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_manager.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	peeringConnectionId := d.Id()

	PeerConnectManager, err := service.DescribeVpcPeerConnectManagerById(ctx, peeringConnectionId)
	if err != nil {
		return err
	}

	if PeerConnectManager == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcPeerConnectManager` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if PeerConnectManager.VpcId != nil {
		_ = d.Set("vpc_id", PeerConnectManager.VpcId)
	}

	if PeerConnectManager.PeeringConnectionName != nil {
		_ = d.Set("peering_connection_name", PeerConnectManager.PeeringConnectionName)
	}

	if PeerConnectManager.PeerUin != nil {
		_ = d.Set("peer_uin", helper.UInt64ToStr(*PeerConnectManager.PeerUin))
	}

	// if PeerConnectManager.Bandwidth != nil {
	// 	_ = d.Set("bandwidth", PeerConnectManager.Bandwidth)
	// }

	// if PeerConnectManager.Type != nil {
	// 	_ = d.Set("type", PeerConnectManager.Type)
	// }

	// if PeerConnectManager.ChargeType != nil {
	// 	_ = d.Set("charge_type", PeerConnectManager.ChargeType)
	// }

	return nil
}

func resourceTencentCloudVpcPeerConnectManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_manager.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewModifyVpcPeeringConnectionRequest()

	peeringConnectionId := d.Id()

	request.PeeringConnectionId = &peeringConnectionId

	immutableArgs := []string{"source_vpc_id", "destination_vpc_id", "destination_uin", "destination_region", "type", "qos_level"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("peering_connection_name") {
		if v, ok := d.GetOk("peering_connection_name"); ok {
			request.PeeringConnectionName = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyVpcPeeringConnection(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update vpc PeerConnectManager failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudVpcPeerConnectManagerRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectManagerDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_manager.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	peeringConnectionId := d.Id()

	if err := service.DeleteVpcPeerConnectManagerById(ctx, peeringConnectionId); err != nil {
		return err
	}

	return nil
}

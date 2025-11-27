/*
Provides a resource to create and manage a VPC peering connection with extended features.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_peer_connect_ex_manager" "example" {
	  vpc_id                  = "vpc-45dvaaw9"
	  peering_connection_name = "test-peer-connection-ex"
	  peer_vpc_id             = "vpc-44fnavba"
	  peer_uin                = "110000053176"
	  peer_region             = "ap-shenzhen-region"
	  bandwidth               = 500
	}

```

# Import

VPC peering connection ex can be imported using the id, e.g.

```
$ terraform import cloud_vpc_peer_connect_ex_manager.example pcx-1asg3t63
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"log"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_peer_connect_ex_manager", CNDescription{
		TerraformTypeCN: "VPC对等连接跨地域管理",
		DescriptionCN:   "提供VPC对等连接扩展管理资源，用于创建和管理跨地域的VPC对等连接。",
		AttributesCN: map[string]string{
			"vpc_id":                  "本端VPC ID",
			"peering_connection_name": "对等连接名称",
			"peer_vpc_id":             "对端VPC ID",
			"peer_uin":                "对端用户UIN",
			"peer_region":             "对端地域",
			"bandwidth":               "对等连接带宽值",
		},
	})
}

func resourceTencentCloudVpcPeerConnectExManager() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudVpcPeerConnectExManagerCreate,
		Read:        resourceTencentCloudVpcPeerConnectExManagerRead,
		Update:      resourceTencentCloudVpcPeerConnectExManagerUpdate,
		Delete:      resourceTencentCloudVpcPeerConnectExManagerDelete,
		Description: "Provides a resource to create and manage VPC peering connection with extended features",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
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
				Description: "The unique ID of the peer VPC.",
			},
			"peer_uin": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Peer user UIN.",
			},
			"peer_region": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Peer region.",
			},
			"bandwidth": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Peering connection bandwidth value.",
			},
		},
	}
}

func resourceTencentCloudVpcPeerConnectExManagerCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_ex_manager.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request             = vpc.NewCreateVpcPeeringConnectionExRequest()
		peeringConnectionId string
	)

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peering_connection_name"); ok {
		request.PeeringConnectionName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_vpc_id"); ok {
		request.PeerVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_uin"); ok {
		request.PeerUin = helper.String(v.(string))
	}

	if v, ok := d.GetOk("peer_region"); ok {
		request.PeerRegion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bandwidth"); ok {
		request.Bandwidth = helper.IntInt64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CreateVpcPeeringConnectionEx(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc PeerConnectExManager failed, reason:%+v", logId, err)
		return err
	}

	peeringConnectionName := d.Get("peering_connection_name").(string)
	vpcId := d.Get("vpc_id").(string)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		describeRequest := vpc.NewDescribeVpcPeeringConnectionsRequest()
		describeResponse, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DescribeVpcPeeringConnections(describeRequest)
		if e != nil {
			return retryError(e)
		}

		for _, conn := range describeResponse.Response.PeerConnectionSet {
			if conn.PeeringConnectionName != nil && *conn.PeeringConnectionName == peeringConnectionName &&
				conn.VpcId != nil && *conn.VpcId == vpcId {
				if conn.PeeringConnectionId != nil {
					peeringConnectionId = *conn.PeeringConnectionId
					log.Printf("[DEBUG]%s Found created peering connection: %s", logId, peeringConnectionId)
					return nil
				}
			}
		}

		return resource.RetryableError(fmt.Errorf("peering connection not found yet, retrying"))
	})

	if err != nil {
		log.Printf("[CRITAL]%s find created vpc PeerConnectExManager failed, reason:%+v", logId, err)
		return err
	}

	if peeringConnectionId == "" {
		return fmt.Errorf("create vpc peering connection ex succeeded but could not find the created connection")
	}

	d.SetId(peeringConnectionId)

	return resourceTencentCloudVpcPeerConnectExManagerRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectExManagerRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_ex_manager.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	peeringConnectionId := d.Id()

	PeerConnectExManager, err := service.DescribeVpcPeerConnectManagerById(ctx, peeringConnectionId)
	if err != nil {
		return err
	}

	if PeerConnectExManager == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcPeerConnectExManager` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if PeerConnectExManager.VpcId != nil {
		_ = d.Set("vpc_id", PeerConnectExManager.VpcId)
	}

	if PeerConnectExManager.PeeringConnectionName != nil {
		_ = d.Set("peering_connection_name", PeerConnectExManager.PeeringConnectionName)
	}

	if PeerConnectExManager.PeerVpcId != nil {
		_ = d.Set("peer_vpc_id", PeerConnectExManager.PeerVpcId)
	}

	if PeerConnectExManager.PeerUin != nil {
		_ = d.Set("peer_uin", helper.UInt64ToStr(*PeerConnectExManager.PeerUin))
	}

	if PeerConnectExManager.DstRegion != nil {
		_ = d.Set("peer_region", PeerConnectExManager.DstRegion)
	}

	if PeerConnectExManager.Bandwidth != nil {
		_ = d.Set("bandwidth", PeerConnectExManager.Bandwidth)
	}

	return nil
}

func resourceTencentCloudVpcPeerConnectExManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_ex_manager.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewModifyVpcPeeringConnectionExRequest()

	peeringConnectionId := d.Id()

	request.PeeringConnectionId = &peeringConnectionId

	immutableArgs := []string{"vpc_id", "peer_vpc_id", "peer_uin", "peer_region"}

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

	if d.HasChange("bandwidth") {
		if v, ok := d.GetOk("bandwidth"); ok {
			request.Bandwidth = helper.IntInt64(v.(int))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyVpcPeeringConnectionEx(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update vpc PeerConnectExManager failed, reason:%+v", logId, err)
		return err
	}

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		peerConnection, e := service.DescribeVpcPeerConnectManagerById(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}

		if peerConnection == nil {
			return resource.NonRetryableError(fmt.Errorf("peering connection not found"))
		}

		if d.HasChange("peering_connection_name") {
			expectedName := d.Get("peering_connection_name").(string)
			if peerConnection.PeeringConnectionName == nil || *peerConnection.PeeringConnectionName != expectedName {
				return resource.RetryableError(fmt.Errorf("peering connection name not updated yet"))
			}
		}

		if d.HasChange("bandwidth") {
			expectedBandwidth := int64(d.Get("bandwidth").(int))
			if peerConnection.Bandwidth == nil || *peerConnection.Bandwidth != expectedBandwidth {
				return resource.RetryableError(fmt.Errorf("peering connection bandwidth not updated yet"))
			}
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s waiting for vpc PeerConnectExManager update to take effect failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudVpcPeerConnectExManagerRead(d, meta)
}

func resourceTencentCloudVpcPeerConnectExManagerDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_peer_connect_ex_manager.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	peeringConnectionId := d.Id()

	if err := service.DeleteVpcPeerConnectExManagerById(ctx, peeringConnectionId); err != nil {
		return err
	}

	return nil
}
